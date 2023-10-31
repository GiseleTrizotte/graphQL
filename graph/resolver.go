package graph

import "example/internal/database"

// Injetamos a dependencia do resolver
// CategoryDB = recebe o arquivo *database.Category que contem a implementacao
// de uma categoria no banco de dados
type Resolver struct {
	CategoryDB *database.Category
}
