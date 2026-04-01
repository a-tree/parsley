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
- 確認
  - 全件
  - 1件 (未実装)
- 登録内容の変更 (未実装)
- 削除 (未実装)
を Webで行う

## 構成
- フロント
  - HTMX
- リバースプロキシ&staticコンテンツ配信
  - nginx dockerコンテナ
  - port:80
  - `/user` へのリクエストを BFFへ転送する
- BFF
  - TypeScript + Node.js + express
  - port:3000
- APIサーバー
  - Go + Echo
  - port:8080
- DBMS
  - SQLite3 or PostgresSQL

## 利用するツール類
- openapi-codegen
- openapi-typescript
- go-task

## ディレクトリ構成
```
parsley
├── api
│   └── openapi.yaml
├── backend
│   ├── cmd
│   │   └── backend
│   │       └── main.go
│   ├── go.mod
│   ├── go.sum
│   ├── internal
│   │   ├── api
│   │   │   ├── api.gen.go
│   │   │   ├── config.yaml
│   │   │   └── handler.go
│   │   ├── config
│   │   │   └── config.go
│   │   ├── domain
│   │   │   └── models
│   │   │       ├── config.yaml
│   │   │       └── models.gen.go
│   │   └── repository
│   │       ├── db.go
│   │       ├── user_repo_test.go
│   │       └── user_repo.go
│   └── Taskfile.yaml
├── bff
│   ├── package-lock.json
│   ├── package.json
│   ├── src
│   │   ├── index.ts
│   │   ├── lib
│   │   │   └── api.ts
│   │   ├── routes
│   │   │   └── users.ts
│   │   └── types
│   │       └── api.d.ts
│   ├── Taskfile.yaml
│   └── tsconfig.json
├── bin
├── config
│   ├── config.toml
│   └── nginx.conf
├── db
│   └── dev.db (APIが作成する)
├── docker-compose.yaml
├── LICENSE
├── README.md
├── scripts
├── Taskfile.yaml
└── www
    └── html
        ├── docs.html
        └── index.html
```

