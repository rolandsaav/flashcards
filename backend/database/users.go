package database

import "fmt"

type User struct {
	Id       int64
	Username string
	Hashed   string
	Salt     string
}

func (db *Database) CreateUser(user User) (*User, error) {
	_, err := db.DB.Exec(
		"INSERT INTO users (username, hashed, salt) VALUES (?, ?, ?)",
		user.Username,
		user.Hashed,
		user.Salt,
	)

	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("Create user: %v", err)
	}

	return &user, nil
}

func (db *Database) IsUniqueUsername(username string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE username = ?`

	// Execute the query and scan the result into count
	err := db.DB.QueryRow(query, username).Scan(&count)
	if err != nil {
		return false, err
	}

	// If count is 0, the username is unique
	return count == 0, nil
}

func (db *Database) GetUserFromUsername(username string) (*User, error) {
	result := db.DB.QueryRow(
		"SELECT * FROM users WHERE username = ?",
		username,
	)

	var user User
	err := result.Scan(&user.Id, &user.Username, &user.Hashed, &user.Salt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
