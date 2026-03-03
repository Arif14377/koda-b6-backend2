package users

import "github.com/arif14377/koda-b6-backend2/internal/models"

// var DataUsers []models.User

type Repository struct {
	db *[]models.User
}

// buat constructor - mekanisme untuk membuat/inisialisasi object/struct
func NewRepository(db *[]models.User) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAllUser() *[]models.User {
	return r.db
}
