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

.
├── api/
│ ├── cmd/
│ ├── internal/
│ │ ├── shared/
│ │ └── todo/
│ └── docker-compose.yml
└── web/
├── src/
├── components/
└── pages/


## セットアップ

### APIの起動
cd api
docker-compose up -d

### WEBの起動
cd web
npm install
npm run dev


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

