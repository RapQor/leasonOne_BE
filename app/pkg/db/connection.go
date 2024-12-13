package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {
	var err error
	dsn := "root:apayak@tcp(mysql_container:3306)/leasonOne"
	log.Println(dsn, "ini database")
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Gagal menghubungkan ke database: %v", err)
	}

	fmt.Println("Berhasil terhubung ke database!")
}
