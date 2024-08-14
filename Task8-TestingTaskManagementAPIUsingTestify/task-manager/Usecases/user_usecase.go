package usecases

import (
	domain "TaskManager/task-manager/Domain"
	Infrastructure "TaskManager/task-manager/Infrastructure"
	"errors"
	"fmt"
	"time"
)

type UserUsecase struct {
	UserRepository  domain.UserRepository
	PasswordService Infrastructure.PasswordService
}

func NewUserUsecase(repository domain.UserRepository, password Infrastructure.PasswordService) *UserUsecase {
	return &UserUsecase{
		UserRepository:  repository,
		PasswordService: password,
	}
}

func (u *UserUsecase) RegisterUser(user domain.User) (domain.User, error) {
	if user.UserName == "" || user.Password == "" {
		return domain.User{}, errors.New("username and password can not be empty")
	}

	_, err := u.UserRepository.GetUserByUserName(user.UserName)

	if err == nil {
		return domain.User{}, errors.New("username is taken")
	} else {
		user, err := u.UserRepository.CreateUser(user)
		if err == nil {
			return user, nil
		} else {
			return domain.User{}, err
		}
	}
}

func (u *UserUsecase) RegisterAdmin(user domain.User) (domain.User, error) {
	if user.UserName == "" || user.Password == "" {
		return domain.User{}, errors.New("username and password can not be empty")
	}

	_, err := u.UserRepository.GetUserByUserName(user.UserName)

	if err == nil {
		return domain.User{}, errors.New("username is taken")
	} else {
		user, err := u.UserRepository.CreateAdmin(user)
		if err == nil {
			return user, nil
		} else {
			return domain.User{}, err
		}
	}
}

func (u *UserUsecase) LoginUser(user domain.User) (string, error) {
	service := Infrastructure.Service{}
	if user.UserName == "" || user.Password == "" {
		return "", errors.New("username and password can not be empty")
	}

	existingUser, err := u.UserRepository.GetUserByUserName(user.UserName)

	if err == nil {
		// set a jwt service
		if u.PasswordService.ValidatePasswordHash(existingUser.Password, user.Password) != nil {
			return "", errors.New("invalid username or password")
		}
		fmt.Println("existing user", existingUser.ID)
		token, err := service.GenerateJwt(existingUser, time.Hour*2)
		if err == nil {
			return token, nil
		} else {
			return "", errors.New("could not generate a Jwt token")
		}
	} else {
		return "", errors.New("username or password does not exist")
	}
}
