package database

import (
	"database/sql"

	"github.com/google/uuid"
)

func NewCategory(db *sql.DB) *Category {
	return &Category{
		db: db,
	}
}

type Category struct {
	db          *sql.DB
	id          string
	name        string
	description string
}

func (c *Category) GetID() string {
	return c.id
}

func (c *Category) GetName() string {
	return c.name
}

func (c *Category) GetDescription() *string {
	return &c.description
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()

	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES (?, ?, ?)", id, name, description)
	if err != nil {
		return Category{}, err
	}

	return Category{
		id:          id,
		name:        name,
		description: description,
	}, nil
}
