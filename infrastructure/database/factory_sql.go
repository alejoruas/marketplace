package database

import "marketplace/adapter/repository"

func NewDBSQL() (repository.SQL, error) {
	return NewPostgressql(CreatePostgresConfig())
}
