package repository

import (
	"github.com/google/uuid"
	"github.com/nanfeng/ginchat/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
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

func (repo *UserRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
