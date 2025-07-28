package database

import (
	"database/sql"

	"github.com/google/uuid"
)

func NewCourse(db *sql.DB) *Course {
	return &Course{
		db: db,
	}
}

type Course struct {
	db          *sql.DB
	ID          string
	Title       string
	Description *string
	CategoryID  string
}

func (c *Course) Create(title string, description string, categoryID string) (*Course, error) {
	id := uuid.New().String()
	query := "INSERT INTO courses (id, title, description, category_id) VALUES (?, ?, ?, ?)"
	_, err := c.db.Exec(query, id, title, description, categoryID)
	if err != nil {
		return nil, err
	}

	return &Course{
		ID:          id,
		Title:       title,
		Description: &description,
		CategoryID:  categoryID,
	}, nil
}

func (c *Course) FindAll() ([]*Course, error) {
	rows, err := c.db.Query("SELECT id, title, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*Course
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.Title, &course.Description, &course.CategoryID); err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return courses, nil
}
