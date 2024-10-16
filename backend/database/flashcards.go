package database

import (
	"fmt"
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

func (db *Database) CreateFlashcard(flashcard Flashcard) (*Flashcard, error) {
	result, err := db.DB.Exec(
		"INSERT INTO flashcards (term, definition, ownerId) VALUES (?, ?, ?)",
		flashcard.Term,
		flashcard.Definition,
		flashcard.OwnerID,
	)

	if err != nil {
		return nil, fmt.Errorf("Add data: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("Add data: %v", err)
	}

	newFlashcard := Flashcard{
		ID:         id,
		OwnerID:    flashcard.OwnerID,
		Term:       flashcard.Term,
		Definition: flashcard.Definition,
	}

	return &newFlashcard, nil
}

func (db *Database) GetFlashcards() ([]Flashcard, error) {
	var flashcards []Flashcard

	rows, err := db.DB.Query("SELECT * FROM flashcards")

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

func (db *Database) GetFlashcardsByOwner(ownerId int64) ([]Flashcard, error) {
	var flashcards []Flashcard

	rows, err := db.DB.Query("SELECT * FROM flashcards WHERE ownerId = ?", ownerId)

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

type NoFlashcardError struct {
	ID int64
}

func (e *NoFlashcardError) Error() string {
	return fmt.Sprintf("No flashcard found in database with id: %d", e.ID)
}

func (db *Database) UpdateFlashcard(flashcard Flashcard) (*Flashcard, error) {
	result, err := db.DB.Exec("UPDATE flashcards SET term = ?, definition = ? WHERE id = ?",
		flashcard.Term,
		flashcard.Definition,
		flashcard.ID,
	)

	if err != nil {
		return nil, fmt.Errorf("Update flashcard: %v", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return nil, fmt.Errorf("Update flashcard, rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return nil, &NoFlashcardError{ID: flashcard.ID}
	}

	resultFlashcard := Flashcard{
		ID:         flashcard.ID,
		OwnerID:    flashcard.OwnerID,
		Term:       flashcard.Term,
		Definition: flashcard.Definition,
	}

	return &resultFlashcard, nil
}

func (db *Database) DeleteFlashcard(flashcardId int64) (bool, error) {
	result, err := db.DB.Exec("DELETE FROM flashcards WHERE id = ?", flashcardId)

	if err != nil {
		return false, fmt.Errorf("Delete flashcard: %v", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return false, fmt.Errorf("Delete flashcard, rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return false, &NoFlashcardError{ID: flashcardId}
	}

	return true, nil
}
