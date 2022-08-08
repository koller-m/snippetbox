package models

import (
	"database/sql"
	"errors"
	"time"
)

// Define Snippet type to hold data for an individual snippet
// The fields of the struct correspond with the MySQL table
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Define SnippetModel type which wraps sql.DB
type SnippetModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	// Write the SQL statement to be executed
	stmt := `INSERT INTO snippets (title, content, created, expires) 
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// Use Exec() method to execute the statement
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// Use LastInsertId() to get ID of our newly inserted record in the snippets table
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// The returned ID is int64
	// Convert to int
	return int(id), nil
}

// This will return a specific snippet based on ID
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets 
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// Use QueryRow() to execute SQL statement
	// Uses the id variable as the ? placeholder param
	// Returns a pointer to sql.Row
	row := m.DB.QueryRow(stmt, id)

	// Init a pointer to a new zeroed Snippet struct
	s := &Snippet{}

	// Use row.Scan() to copy values from sql.Row to Snippet struct
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

// This will return the 10 most recent snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets 
	WHERE expires > UTC_TIMESTAMP() ORDER BY id LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	// Defer rows.Close() to ensure result set is properly closed before
	// Latest() returns
	defer rows.Close()

	// Init an empty slice to hold the Snippet structs
	snippets := []*Snippet{}

	// Use rows.Next() to iterate through the rows in the result set
	for rows.Next() {
		// Create a pointer to a new zeroed Snippet struct
		s := &Snippet{}
		// Use rows.Scan() to copy values from each field in the row
		// To the new Snippet object
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		// Append to the slice
		snippets = append(snippets, s)
	}

	// Call rows.Err() to retrieve any errors during the iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}
	// Otherwise, everything went OK
	return snippets, nil
}
