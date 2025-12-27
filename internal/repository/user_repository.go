package repository

import (
	"github.com/google/uuid"
	"github.com/nanfeng/ginchat/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db gorm.DB
}

func NewUserRepository(db gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) CreateUser(user *model.User) (uuid.UUID, error) {
	if err := repo.db.Create(user).Error; err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}
