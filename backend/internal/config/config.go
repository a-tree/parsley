package config

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

type Config struct {
	Database struct {
		Name string `mapstructure:"name"`
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
		User string `mapstructure:"user"`
		Pass string `mapstructure:"password"`
	} `mapstructure:"database"`

	Api struct {
		Port int    `mapstructure:"port"`
		Host string `mapstructure:"host"`
	} `mapstructure:"api"`
}

func NewConfig() (*Config, error) {
	v := viper.New()

	// 1. デフォルト値の設定
	v.SetDefault("database.name", "dev")
	v.SetDefault("api.port", 8080)
	v.SetDefault("api.host", "localhost")

	// 2. config.toml の読み込み
	configPath := configPath()
	v.SetConfigFile(configPath)
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("[NewConfig] ReadInConfig Error: %v\n", err)
	}

	// 3.APP_ENVに合わせた .env ファイルの読み込み (ファイルがあれば読み込む)
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "develop"
	}
	envFile := fmt.Sprintf("../.env.%s", appEnv)
	// .env は環境変数形式で記述するため godotenv を用いる
	_ = gotenv.Load(envFile)
	// Viper の .env サポートを利用する場合は database.password=".." のように記述する必要がある

	// 4. システム環境変数の読み込みとマッピング設定
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // Database.Name -> DATABASE_NAME
	v.AutomaticEnv()

	// 個別のプレフィックス対応 (DB_USER -> database.user)
	v.BindEnv("database.name", "DB_NAME")
	v.BindEnv("database.host", "DB_HOST")
	v.BindEnv("database.port", "DB_PORT")
	v.BindEnv("database.user", "DB_USER")
	v.BindEnv("database.password", "DB_PASSWORD")
	v.BindEnv("api.port", "API_PORT")
	v.BindEnv("api.host", "API_HOST")

	// 5. 構造体へデコード
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	// 必須チェック
	if cfg.Database.Name == "" {
		return nil, fmt.Errorf("DB_NAME is required")
	}

	return &cfg, nil
}

func configPath() string {
	var path string
	flag.StringVar(&path, "config-file", "../config/config.toml", "設定ファイルのパス")
	flag.StringVar(&path, "f", "../config/config.toml", "設定ファイルのパス (short)")
	flag.Parse()
	return path
}
