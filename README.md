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
parsley
├── api
│   └── openapi.yaml
├── backend
│   ├── cmd
│   │   └── backend
│   │       └── main.go
│   ├── go.mod
│   ├── go.sum
│   └── internal
│       ├── api
│       │   ├── api.gen.go
│       │   ├── config.yaml
│       │   └── handler.go
│       ├── config
│       │   └── config.go
│       ├── domain
│       │   └── models
│       │       ├── config.yaml
│       │       └── models.gen.go
│       └── repository
│           ├── db.go
│           ├── user_repo_test.go
│           └── user_repo.go
├── bff
│   ├── package-lock.json
│   ├── package.json
│   └── src
│       └── types
│           └── api.d.ts
├── LICENSE
├── README.md
├── scripts
├── Taskfile.yaml
└── web
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
