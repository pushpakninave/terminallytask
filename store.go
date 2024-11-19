package main

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Note struct {
	ID    int64
	Title string
	Body  string
}

type Store struct {
	conn *sql.DB
}

func (s *Store) Init() error {
	var err error
	s.conn, err = sql.Open("sqlite3", "./notes.db")
	if err != nil {
		return err
	}

	createTableStatement := `CREATE TABLE IF NOT EXISTS notes(
		id integer not null primary key, 
		title text not null,
		body text not null
	);`

	if _, err = s.conn.Exec(createTableStatement); err != nil {
		return err
	}

	return nil
}

func (s *Store) GetNotes() ([]Note, error) {
	rows, err := s.conn.Query("SELECT * FROM notes")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	notes := []Note{}

	for rows.Next() {
		var note Note
		rows.Scan(&note.ID, &note.Title, &note.Body)
		notes = append(notes, note)
	}

	return notes, nil
}

func (s *Store) DeleteNote(noteId int64) error {
	deleteQuery := `DELETE FROM notes WHERE id = ?`
	if _, err := s.conn.Exec(deleteQuery, noteId); err != nil {
		return err
	}
	return nil
}

func (s *Store) SaveNote(note Note) error {
	if note.ID == 0 {
		note.ID = time.Now().UTC().UnixNano()
	}
	upsertQuery := `INSERT INTO notes (id, title, body)
	VALUES (?,?,?)
	ON CONFLICT(id) DO UPDATE SET 
	title=excluded.title, 
	body=excluded.body;
	`
	if _, err := s.conn.Exec(upsertQuery, note.ID, note.Title, note.Body); err != nil {
		return err
	}

	return nil
}
