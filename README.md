# TODOアプリ

シンプルで使いやすいTODOタスク管理アプリケーション

## 機能

- タスクの検索
- タスクの登録
- タスクの更新
- タスクの削除

## 技術スタック

### バックエンド (API)

- Go
- PostgreSQL
- Docker

### フロントエンド (WEB)

- TypeScript
- React
- Next.js

## プロジェクト構成
<pre>
.
├── api/
│   ├── Dockerfile
│   ├── docker-compose.yml
│   ├── docs
│   │   ├── docs.go
│   │   ├── swagger.json
│   │   └── swagger.yaml
│   ├── go.mod
│   ├── go.sum
│   ├── internal
│   │   ├── domain
│   │   ├── infrastructure
│   │   ├── interface_adapter
│   │   ├── shared
│   │   └── usecase
│   ├── main.go
│   └── migrations
└── web/
    ├── src/
    ├── components/
    └── pages/
</pre>

## セットアップ

### APIの起動
<pre>
cd api
docker-compose up -d
</pre>

### WEBの起動
<pre>
cd web
npm install
npm run dev
</pre>

## API エンドポイント

| メソッド | エンドポイント | 説明 |
|----------|----------------|------|
| GET      | /api/todos     | タスク一覧の取得 |
| POST     | /api/todos     | タスクの作成 |
| PUT      | /api/todos/:id | タスクの更新 |
| DELETE   | /api/todos/:id | タスクの削除 |

## 開発者向け情報

- APIドキュメント: `http://localhost:8080/swagger/index.html`
- フロントエンド開発サーバー: `http://localhost:3000`

