package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMYSqlStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("[ERRPR in pakage DB]:", err)
	}

	return db, nil
}

func InitDataBase(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	log.Println("DB SUCCESSFULLY CONNECTED....")
	return nil
}
