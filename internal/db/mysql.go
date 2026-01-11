package db

import (
	"auth-service/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnecMysql(cfg *config.ConfigMySQL) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&tls=false",
		cfg.MySqlUser,
		cfg.MySqlPassword,
		cfg.MySqlHost,
		cfg.MySqlPort,
		cfg.MySqlDB,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("connected to MySQL!")
	return db, err
}
