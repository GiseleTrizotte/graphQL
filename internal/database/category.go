package database

import (
	"database/sql"

	"github.com/google/uuid"
)

// struct = estrutura de dados
type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

// A funcao foi criada para toda vez que criar uma categoria eu poder
// passar minha conexao com o banco de dados
func NewCategory(db *sql.DB) *Category {
	// Vai retonar um apontamento na memoria da minha aplicacao onde
	// estara o valor de categoria
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VAUES ($1, $2, $3)",
		id, name, description)

	if err != nil {
		// Retorna uma ategoria vazia e o erro
		return Category{}, err
	}

	// Retorna a categoria e o erro vazio
	return Category{ID: id, Name: name, Description: description}, nil
}
