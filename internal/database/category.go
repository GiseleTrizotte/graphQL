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
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)",
		id, name, description)

	if err != nil {
		// Retorna uma ategoria vazia e o erro
		return Category{}, err
	}

	// Retorna a categoria e o erro vazio
	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM  categories")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Cria um objeto de categoria em branco
	categories := []Category{}

	for rows.Next() {
		var id, name, description string

		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}

		categories = append(categories, Category{ID: id, Name: name, Description: description})

	}
	return categories, nil
}

func (c *Category) FindByCourseID(courseID string) (Category, error) {
	var id, name, description string
	err := c.db.QueryRow("SELECT c.id, c.name, c.description FROM categories c JOIN courses co ON c.id = co.category_id WHERE co.id = $1", courseID).
		Scan(&id, &name, &description)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil
}
