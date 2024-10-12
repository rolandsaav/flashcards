package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type Flashcard struct {
	ID         int64  `json:"id"`
	OwnerID    int64  `json:"owner_id"`
	Term       string `json:"term"`
	Definition string `json:"definition"`
}

func (flashcard Flashcard) String() string {
	return fmt.Sprintf(
		"ID: %d\nOwner: %d\nTerm: %s\nDefinition: %s",
		flashcard.ID,
		flashcard.OwnerID,
		flashcard.Term,
		flashcard.Definition,
	)
}

func ConnectToDB(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		return nil, fmt.Errorf("Connect to database: %v", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("What the fuck 2?")
		log.Fatal(pingErr)
		return nil, fmt.Errorf("Connect to database: %v", pingErr)
	}
	return db, nil
}

func CreateFlashcard(db *sql.DB, flashcard Flashcard) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO flashcard (term, definition, ownerId) VALUES (?, ?, ?)",
		flashcard.Term,
		flashcard.Definition,
		flashcard.OwnerID,
	)

	if err != nil {
		return 0, fmt.Errorf("Add data: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Add data: %v", err)
	}

	return id, nil
}

func GetFlashcards(db *sql.DB) ([]Flashcard, error) {
	var flashcards []Flashcard

	rows, err := db.Query("SELECT * FROM flashcard")

	if err != nil {
		return nil, fmt.Errorf("Get all flashcards: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var card Flashcard
		if err := rows.Scan(&card.ID, &card.OwnerID, &card.Term, &card.Definition); err != nil {
			return nil, fmt.Errorf("Get all flashcards: %v", err)
		}
		flashcards = append(flashcards, card)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Get all flashcards: %v", err)
	}

	return flashcards, nil
}

func GetFlashcardsByOwner(db *sql.DB, ownerId int64) ([]Flashcard, error) {
	var flashcards []Flashcard

	rows, err := db.Query("SELECT * FROM flashcard WHERE ownerId = ?", ownerId)

	if err != nil {
		return nil, fmt.Errorf("Get flashcards by ownerId: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var card Flashcard
		if err := rows.Scan(&card.ID, &card.OwnerID, &card.Term, &card.Definition); err != nil {
			return nil, fmt.Errorf("Get flashcards by ownerId: %v", err)
		}
		flashcards = append(flashcards, card)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Get all flashcards: %v", err)
	}

	return flashcards, nil
}
