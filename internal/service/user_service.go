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
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) CreateUser(req *request.CreateUserRequest) (uuid.UUID, error) {
	// 1.校验 username和email是否重复
	if _, err := svc.repo.GetByUsername(req.Username); err == nil {
		return uuid.Nil, errors.New("The username has Already been exists!")
	}

	if _, err := svc.repo.GetByEmail(req.Email); err == nil {
		return uuid.Nil, errors.New("The email has Already been exists!")
	}

	var user model.User
	// 2.生成id
	uid, err := uuid.NewV7()
	if err != nil {
		return uuid.Nil, errors.New("id生成错误")
	}
	user.ID = uid
	// 3.加密密码
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, errors.New("密码加密错误")
	}
	user.Password = string(password)

	// 4.赋值其他属性
	user.Username = req.Username
	user.Email = req.Email
	user.Age = req.Age
	user.Gender = req.Gender

	// 5.repository 层方法
	id, err := svc.repo.CreateUser(&user)
	if err != nil {
		return uuid.Nil, err
	}

	// 6.返回结果
	return id, nil
}
