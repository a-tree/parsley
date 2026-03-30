# サンプル用CRUD Webアプリ「Parsley(パセリ)」

## データ
ユーザー情報
```
{
	"Name": "名前文字列"
	"Email": "E mail文字列"
}
```

## パセリでできること
ユーザー情報の
- 登録
- 確認(取得)
- 登録内容の変更
- 削除
を Webで行う

## ディレクトリ構成
```
parsley/
  api/
    openapi.yaml          # 全ての設計の基盤 バックエンドとBFFの両方から参照
  backend/                # Go + Echo
    internal/api/         # 自動生成コードの出力先
    oapi-codegen.cfg.yaml # backend専用の生成設定 バックエンドの実装に近い場所に置くのが一般的
    main.go
  bff/                    # TS + Node.js + express
    src/types/            # openapi-typescriptの出力先
    src/index.ts
  web/                    # HTMX / Static contents
  scripts/                # (オプション) 補助スクリプト
  Taskfile.yaml            # ルートに置くことで全体を制御 task gen コマンド一つでbackendとbffの両方のコード生成を走らせるため
  docker-compose.yaml
```

## 構成
- フロント
  - HTMX
- リバースプロキシ&staticコンテンツ配信
  - nginx dockerコンテナ
  - port:80
- BFF
  - TypeScript + Node.js + express
- APIサーバー
  - Go + Echo
  - port:8080
- DBMS
  - SQLite3 or PostgresSQL
