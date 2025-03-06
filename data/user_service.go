package data

import (
	"TM/models"
	"context"
	"fmt"
	"errors"
	// "go/token"
	// "net/http"
	"os"
	"time"

	// "fmt"
	// "time"

	// "log"

	"github.com/dgrijalva/jwt-go"
	// "github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserManagment interface {
	RegisterUser(newUser models.User) error 
	LoginUser(email , userPassword string) (string,error)
}

type UserService struct {
	collection *mongo.Collection
}

func NewUserService (client *mongo.Client,dbName,collectionName string) *UserService {
	return &UserService{
		collection : client.Database(dbName).Collection(collectionName),
	}
}

func CreateToken(user models.User) (string, error) {
	jwtKeyHex, ok := os.LookupEnv("Jwt_secret")
	if !ok {
		return "", errors.New("Jwt_secret not found")
	}

	jwtKey := []byte(jwtKeyHex)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	JwtToken, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return JwtToken, nil
}






func (us *UserService) RegisterUser(newUser models.User) error {
	var existingUser models.User
	err := us.collection.FindOne(context.Background(), bson.M{"email": newUser.Email}).Decode(&existingUser)
	if err == nil {
		return errors.New("user exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("internal server error: %v", err)
	}

	newUser.Password = string(hashedPassword)
	_, err = us.collection.InsertOne(context.Background(), newUser)
	if err != nil {
		return fmt.Errorf("internal error to upload to database: %v", err)
	}

	return nil
}


func (us *UserService) LoginUser(email, password string) (string, error) {
	var result models.User
	err := us.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&result)
	if err != nil {
		return "", fmt.Errorf("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("incorrect username or password")
	}

	token, err := CreateToken(result)
	if err != nil {
		return "", fmt.Errorf("internal server error")
	}

	return token, nil
}