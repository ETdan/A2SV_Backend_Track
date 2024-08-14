package repositories

import (
	db "TaskManager/task-manager/DB"
	domain "TaskManager/task-manager/Domain"
	Infrastructure "TaskManager/task-manager/Infrastructure"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepositorySuite struct {
	suite.Suite
	repository *UserRepository
}

func (suite *UserRepositorySuite) SetupSuite() {
	if err := godotenv.Load("..\\Delivery\\.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	database := db.Database{Url: os.Getenv("DATABASE_URL")}

	if err := database.Connect(os.Getenv("DATABASE"), os.Getenv("USER_COLLECTION"), os.Getenv("TASK_COLLECTION")); err != nil {
		log.Fatal("could not connect with DB", err)
	}

	password := Infrastructure.NewPassword()

	repository := NewUserRepository(*database.Database, *database.UserCollection, *password)
	suite.repository = repository
}
func (suite *UserRepositorySuite) TestGetUserByUserName() {
	user, err := suite.repository.GetUserByUserName("d")
	suite.NoError(err, "should not return error when getting user by username")
	suite.NotNil(user, "Expected User to be non-nil")
	suite.NotEmpty(user.ID, "User ID should not be empty")
	suite.NotEmpty(user.Password, "User Password should not be empty")
	suite.Contains(user.Role, "user", "Expected role to contain a certain text")
	suite.IsType(primitive.ObjectID{}, user.ID, "User ID should be of type string")
	suite.IsType("", user.UserName, "User Title should be of type string")
}

func (suite *UserRepositorySuite) TestGetUserByID() {
	id, er := primitive.ObjectIDFromHex("")
	suite.NoError(er, "should not return error when converting string to objectid")
	user, err := suite.repository.GetUserByID(id)
	suite.NoError(err, "should not return error when getting user by username")
	suite.NotNil(user, "Expected User to be non-nil")
	suite.NotEmpty(user.ID, "User ID should not be empty")
	suite.NotEmpty(user.Password, "User Password should not be empty")
	suite.Contains(user.Role, "user", "Expected role to contain a certain text")
	suite.IsType(primitive.ObjectID{}, user.ID, "User ID should be of type string")
	suite.IsType("", user.UserName, "User Title should be of type string")
}

func (suite *UserRepositorySuite) TestCreateUser() {
	user, err := suite.repository.CreateUser(domain.User{
		ID:       primitive.NewObjectID(),
		UserName: "johndoe",
		Password: "$2a$12$KIXP.cTqX9FZIdaP/ujIZsEjJRSBguW53k4VjPqD5dJOsJbKQ0P4G",
		Role:     "user",
	})
	suite.NoError(err, "should not return error when getting user by username")
	suite.NotNil(user, "Expected User to be non-nil")
	suite.NotEmpty(user.ID, "User ID should not be empty")
	suite.NotEmpty(user.Password, "User Password should not be empty")
	suite.NotEmpty(user.UserName, "User Username should not be empty")
	suite.Contains(user.Role, "user", "Expected role to contain a certain text")
	suite.IsType(primitive.ObjectID{}, user.ID, "User ID should be of type string")
	suite.IsType("", user.UserName, "User Title should be of type string")

}

func (suite *UserRepositorySuite) TestCreateAdmin() {
	user, err := suite.repository.CreateAdmin(domain.User{
		ID:       primitive.NewObjectID(),
		UserName: "johndoe",
		Password: "$2a$12$KIXP.cTqX9FZIdaP/ujIZsEjJRSBguW53k4VjPqD5dJOsJbKQ0P4G",
		Role:     "admin",
	})
	suite.NoError(err, "should not return error when getting user by username")
	suite.NotNil(user, "Expected User to be non-nil")
	suite.NotEmpty(user.ID, "User ID should not be empty")
	suite.NotEmpty(user.Password, "User Password should not be empty")
	suite.NotEmpty(user.UserName, "User Username should not be empty")
	suite.Contains(user.Role, "admin", "Expected role to contain a certain text")
	suite.IsType(primitive.ObjectID{}, user.ID, "User ID should be of type string")
	suite.IsType("", user.UserName, "User Title should be of type string")
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}
