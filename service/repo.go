package service

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

var r *Repository

func NewRepository(db *gorm.DB) *Repository {
	if r == nil {
		r = &Repository{db: db}
	}

	return r
}