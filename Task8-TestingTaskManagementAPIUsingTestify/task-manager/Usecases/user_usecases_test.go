package usecases

import (
	domain "TaskManager/task-manager/Domain"
	"TaskManager/task-manager/Infrastructure"
	"TaskManager/task-manager/mocks"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecaseSuite struct {
	suite.Suite
	userRepository *mocks.UserRepository
	password       *mocks.PasswordInterface
	userUsecase    UserUsecase
}

func (suite *UserUsecaseSuite) SetupSuite() {
	userRepository := new(mocks.UserRepository)
	passwordMock := new(mocks.PasswordInterface)
	password := Infrastructure.NewPassword()
	usecase := NewUserUsecase(userRepository, *password)

	suite.userRepository = userRepository
	suite.userUsecase = *usecase
	suite.password = passwordMock
}

func (suite *UserUsecaseSuite) SetupTest() {
	password := Infrastructure.NewPassword()
	suite.userRepository = new(mocks.UserRepository)
	suite.userUsecase = *NewUserUsecase(suite.userRepository, *password)
}

func (suite *UserUsecaseSuite) TestRegisterUser_positive() {
	// Prepare test data
	user := domain.User{
		ID:       primitive.NewObjectID(),
		UserName: "x",
		Password: "x",
		Role:     "user",
	}

	// Set up mock expectations
	suite.userRepository.On("GetUserByUserName", user.UserName).Return(domain.User{}, errors.New("error user not found"))
	suite.userRepository.On("CreateUser", user).Return(user, nil)

	// Call the method under test
	returnedUser, err := suite.userUsecase.RegisterUser(user)

	// Assertions
	suite.Nil(err, "should be nil")
	suite.Equal(user.ID, returnedUser.ID, "User ID should match")
	suite.Equal(user.UserName, returnedUser.UserName, "UserName should match")

	// Verify all mocks expectations were met
	suite.userRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestRegisterUser_UsernameAlreadyExists() {
	user := domain.User{
		ID:       primitive.NewObjectID(),
		UserName: "existingUser",
		Password: "password",
		Role:     "user",
	}

	// Set up mock expectations
	suite.userRepository.On("GetUserByUserName", user.UserName).Return(user, nil)

	// Call the method under test
	returnedUser, err := suite.userUsecase.RegisterUser(user)

	// Assertions
	suite.Error(err, "should return error when username already exists")
	suite.EqualError(err, "username is taken", "error message should match")
	suite.Empty(returnedUser.ID, "User ID should be empty")

	// Verify all mocks expectations were met
	suite.userRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestRegisterAdmin_positive() {
	// Prepare test data
	user := domain.User{
		ID:       primitive.NewObjectID(),
		UserName: "x",
		Password: "x",
	}

	// Set up mock expectations
	suite.userRepository.On("GetUserByUserName", user.UserName).Return(domain.User{}, errors.New("error user not found"))
	suite.userRepository.On("CreateAdmin", user).Return(user, nil)

	// Call the method under test
	returnedUser, err := suite.userUsecase.RegisterAdmin(user)

	// Assertions
	suite.Nil(err, "should be nil")
	suite.Equal(user.ID, returnedUser.ID, "User ID should match")
	suite.Equal(user.UserName, returnedUser.UserName, "UserName should match")

	// Verify all mocks expectations were met
	suite.userRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestRegisterAdmin_UsernameAlreadyExists() {
	user := domain.User{
		ID:       primitive.NewObjectID(),
		UserName: "existingUser",
		Password: "password",
		Role:     "user",
	}

	// Set up mock expectations
	suite.userRepository.On("GetUserByUserName", user.UserName).Return(user, nil)

	// Call the method under test
	returnedUser, err := suite.userUsecase.RegisterAdmin(user)

	// Assertions
	suite.Error(err, "should return error when username already exists")
	suite.EqualError(err, "username is taken", "error message should match")
	suite.Empty(returnedUser.ID, "User ID should be empty")

	// Verify all mocks expectations were met
	suite.userRepository.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestLoginUser_positive() {
	user := domain.User{
		UserName: "d",
		Password: "p",
	}
	suite.userRepository.On("GetUserByUserName", "d").Return(user, nil)
	suite.password.On("ValidatePasswordHash", user.Password, user.Password).Return(nil)
	res, err := suite.userUsecase.LoginUser(user)
	fmt.Println(res)
	suite.Error(err, "should not be error")
}

func (suite *UserUsecaseSuite) TestLoginUser_EmptyUsername() {
	user := domain.User{
		UserName: "",
		Password: "p",
	}
	suite.userRepository.On("GetUserByUserName", "").Return(domain.User{}, errors.New("username should not be empty"))
	res, err := suite.userUsecase.LoginUser(user)
	fmt.Println(res)
	suite.Error(err, "should return error for empty username")
}

func (suite *UserUsecaseSuite) TestLoginUser_UserNotFound() {
	user := domain.User{
		UserName: "d",
		Password: "p",
	}
	suite.userRepository.On("GetUserByUserName", "d").Return(domain.User{}, errors.New("could not find user"))
	res, err := suite.userUsecase.LoginUser(user)
	fmt.Println(res)
	suite.Error(err, "should return error when user is not found")
}

func (suite *UserUsecaseSuite) TestLoginUser_PasswordMismatch() {
	user := domain.User{
		UserName: "d",
		Password: "wrong_password",
	}
	suite.userRepository.On("GetUserByUserName", "d").Return(user, nil)
	suite.password.On("ValidatePasswordHash", user.Password, "wrong_password").Return(errors.New("invalid username or password"))
	res, err := suite.userUsecase.LoginUser(user)
	fmt.Println(res)
	suite.Error(err, "should return error for password mismatch")
}

func TestUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
}
