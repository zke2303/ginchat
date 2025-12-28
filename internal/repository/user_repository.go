package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nanfeng/ginchat/internal/model"
	"github.com/nanfeng/ginchat/internal/pkg/xerr"
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

// GetById 根据用户id查询用户信息
func (repo UserRepository) GetById(id string) (*model.User, error) {
	// 1.定义变量，接收查询结果
	var user model.User
	if err := repo.db.First(&user, "id = ?", id).Error; err != nil {
		// 判断异常类型
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerr.Wrap(
				xerr.CodeNotFound,
				"Not foud",
				err,
			)
		}

		return nil, err
	}

	// 2.查询成功,返回查询到的信息
	return &user, nil
}

// Delete
func (repo *UserRepository) Delete(id string) error {
	// 1.执行Sql操作
	result := repo.db.Delete(&model.User{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return xerr.New(xerr.CodeNotFound, "This user is not exists")
	}

	return nil
}

// Update
func (repo *UserRepository) Update(id string, maps *map[string]any) error {
	result := repo.db.Model(&model.User{}).Where("id = ?", id).Updates(maps)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return xerr.Wrap(xerr.CodeNotFound, "Not Found", result.Error)
	}

	return nil
}
