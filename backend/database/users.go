package database

import "database/sql"

type User struct {
	id       int64
	username string
	hashed   string
	salt     string
}

type UserDB struct {
	DB *sql.DB
}
