# go_stock_analysis

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
