#!/bin/bash
docker-compose up -d

# コンテナ立ち上げ直後はmysqlにアクセスできなかったので数秒待機
sleep 10

# テーブル作成
# docker-compose run --rm dev mysqldef -uroot -ppassword -hdb go_stock_analysis < db/schema.sql
docker-compose exec dev bash -c "mysqldef -uroot -ppassword -hdb go_stock_analysis < db/schema.sql"

# データ投入
docker-compose exec db bash -c "mysql -uroot -ppassword go_stock_analysis < init.sql"