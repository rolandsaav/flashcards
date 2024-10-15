package database

import "fmt"
import "database/sql"
import "log"
import "github.com/go-sql-driver/mysql"

type Database struct {
	DB *sql.DB
}

func ConnectToDB(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		return nil, fmt.Errorf("Connect to database: %v", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
		return nil, fmt.Errorf("Connect to database: %v", pingErr)
	}
	return db, nil
}
