package main

import (
	"database/sql"
	"fmt"
	"go_stock_analysis/registry"
	"go_stock_analysis/server"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func run() error {

	dsn := mysqlDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	defer db.Close()

	r := registry.NewRegistry(db)
	s := server.NewServer(r, true)

	http.ListenAndServe(":8080", s)

	return fmt.Errorf("dd")
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// mysqlDSN MySQLの接続に必要なdata source nameを返す
func mysqlDSN() string {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbName)
}
