package graph

import "example/internal/database"

type Resolver struct {
	CategoryDB *database.Category
	CourseDB   *database.Course
}
