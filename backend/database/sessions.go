package database

import (
	"fmt"
	"time"
)

type Session struct {
	Id         int64
	UserId     int64
	Token      string
	Expiration time.Time
	Created    time.Time
	Expired    bool
}

func (db *Database) CreateSession(session Session) (*Session, error) {
	_, err := db.DB.Exec(
		"INSERT INTO sessions (user_id, token, expiration, created_at, expired) VALUES (?, ?, ?, ?, ?)",
		session.UserId,
		session.Token,
		session.Expiration,
		session.Created,
		false,
	)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &session, nil
}

func (db *Database) GetSessionByToken(token string) (*Session, error) {
	result := db.DB.QueryRow(
		"SELECT * FROM sessions WHERE token = ?",
		token,
	)

	var session Session
	err := result.Scan(&session.Id, &session.UserId, &session.Token, &session.Expiration, &session.Created, &session.Expired)

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (db *Database) UpdateSessionToken(session Session) error {
	_, err := db.DB.Exec("UPDATE sessions SET expiration = ? WHERE id = ?",
		time.Now().Add(time.Minute*30),
		session.Id)

	if err != nil {
		return err
	}

	return nil
}

func (db *Database) InvalidateSessionToken(token string) error {
	_, err := db.DB.Exec("UPDATE sessions SET expired = ? WHERE token = ?",
		true,
		token)

	if err != nil {
		return err
	}

	return nil
}
