# go_stock_analysis

## 開発手法
[ドメイン駆動設計](https://codezine.jp/article/detail/11968)

## ディレクトリ構成

```text
.
├── apitemplates/        // APIコード自動生成用のテンプレート
├── db/schema.sql        // DBスキーマSQL
├── message/             // メッセージ文面
├── mysql/               // 開発環境用Docker
├── domain/              // DDD domain層
├── infra/               // DDD infra層
├── usecase/             // DDD usecase層
├── server/              // DDD application層
├── Dockerfile           // 開発環境用Docker
├── README.md
├── api.yaml             // API定義書
├── docker-compose.yml   // 開発環境用Docker
├── go.mod               // 依存管理ファイル
├── go.sum               // 依存管理ファイル(lock)
├── makefile             // make
├── sqlboiler.toml       // sql-boiler用
└── start.sh             // コンテナ自動立ち上げ
```

## ライブラリ

| ライブラリ名 | Note | コマンド |
| ------------ | ---- | -------- |
| [chi](https://github.com/go-chi/chi)                          | Router         | - |
| [SQLBoiler](https://github.com/volatiletech/sqlboiler)        | ORM作成ツール  | sqlboiler |
| [sqldef](https://github.com/k0kubun/sqldef)                   | DBスキーマ同期 | mysqldef |


## ローカル開発環境セットアップ方法

1. 株取得用にAPIキーを取得する(https://www.alphavantage.co/support/#api-key)
  - `docker-compose.ymlのAPI_KEYを取得したAPIキーに置き換える`

2. Docker compose実行環境を用意する
  - `$ docker-compose up -d`
  - `$ docker-compose exec dev bash`
  - `/app$ mysqldef -uroot -ppassword -hdb go_stock_analysis < db/schema.sql`
3. DB初期データを反映する
  - dbコンテナに入り、sqlファイルを実行
    - `$ docker-compose exec db bash`
    - `/$ mysql -uroot -ppassword go_stock_analysis < init.sql`
4. サーバー立ち上げ
  - `/app$ make dev`
5. サーバが立ち上がる:tada:

※上記1〜3は`$ sh start.sh`でも実行できます

## インフラコード管理
[terraform_ecs](https://github.com/takehiro-1029/terraform_ecs)
