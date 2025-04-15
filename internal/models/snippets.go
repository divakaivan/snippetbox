package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `
		insert into snippets (title, content, created, expires)
		values(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))
	`
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `
		select id, title, content, created, expires from snippets
		where expires > UTC_TIMESTAMP() and id = ?
	`

	row := m.DB.QueryRow(stmt, id)

	s := &Snippet{}

	// Scan auto converts SQL types to Go types
	// char, varchar, text -> string
	// boolean -> bool
	// int -> int; bigint -> int64
	// decimal, numeric -> float
	// time, date, timestamp -> time
	// we pass parseTime=true in the sql string in main to force it to convert time and date fields to time.Time
	// otherwise it's []byte objects
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}

	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {

	stmt := `
		select id, title, content, created, expires from snippets
		where expires > UTC_TIMESTAMP() order by id desc limit 10
	`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close() // important!!!

	snippets := []*Snippet{}

	for rows.Next() {
		s := &Snippet{}
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	// when the rows.Next() loop finishes call rows.Error to get any encountered error during the iter
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
