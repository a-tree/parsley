package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Database struct {
		Name string `toml:"name" env:"NAME,required"`
		Host string `toml:"host" env:"HOST"`
		Port int    `toml:"port" env:"PORT"`
		User string `env:"USER"`
		Pass string `env:"PASSWORD"`
	} `toml:"database" envPrefix:"DB_"`

	Api struct {
		Port int `toml:"port" env:"PORT"`
	} `toml:"api" envPrefix:"API_"`
}

func NewConfig() (*Config, error) {

	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "develop"
	}

	path := ConfigPath("./config/config.toml")
	newConfig, err := ParseToml(path)
	envFile := fmt.Sprintf(".env.%s", appEnv)

	// ファイルが存在する場合のみ読み込む（本番環境等でファイルがないケースを許容）
	_, err = os.Stat(envFile)
	if err == nil {
		err = godotenv.Load(envFile)
		if err != nil {
			return nil, fmt.Errorf("failed to load %s: %w", envFile, err)
		}
	}

	err = env.Parse(newConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &newConfig, nil
}

func ParseToml(path string) (Config, error) {
	config := defaultConfig()
	f, err := os.Open(path)
	if err != nil {
		return config, err
	}
	defer f.Close()
	if _, err := toml.NewDecoder(f).Decode(&config); err != nil {
		return config, fmt.Errorf("[setting] 設定デコードエラー: %v", err)
	}
	return config, nil
}

func defaultConfig() Config {
	defaultConfig := Config{}

	defaultConfig.Database.Name = "dev"
	defaultConfig.Database.Host = "localhost"
	defaultConfig.Database.Port = 5432
	defaultConfig.Api.Port = 8080

	return defaultConfig
}

func ConfigPath(defaultPath string) string {

	var settingFilePath string

	// --setting-file フラグを定義
	flag.StringVar(
		&settingFilePath,
		"setting-file",
		defaultPath,
		"設定ファイルのパスを指定します", // 使用法のメッセージ
	)

	// -f フラグを定義（--setting-file のショートハンドとして同じ変数に紐付ける）
	flag.StringVar(
		&settingFilePath,
		"f",
		defaultPath,
		"設定ファイルのパスを指定します (ショートハンド)", // 使用法のメッセージ
	)

	// コマンドオプションが -h --help の時に Usageを表示
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "使い方: %s [options]\n", os.Args)
		fmt.Fprintln(os.Stderr, "オプション:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.Parse()

	return settingFilePath
}
