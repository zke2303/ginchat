package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nanfeng/ginchat/internal/model"
	"github.com/nanfeng/ginchat/internal/model/request"
	"github.com/nanfeng/ginchat/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) CreateUser(req *request.CreateUserRequest) (uuid.UUID, error) {
	var user model.User
	// 1.生成id
	uid, err := uuid.NewV7()
	if err != nil {
		return uuid.Nil, errors.New("id生成错误")
	}
	user.ID = uid
	// 2.加密密码
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, errors.New("密码加密错误")
	}
	user.Password = string(password)

	// 3.调用 repository 层方法
	id, err := svc.repo.CreateUser(&user)
	if err != nil {
		return uuid.Nil, err
	}
	// 4.返回结果
	return id, nil
}
