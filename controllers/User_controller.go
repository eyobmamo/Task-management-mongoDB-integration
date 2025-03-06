package controllers

import (

	"TM/data"
	"TM/models"
	"net/http"
	// "strconv"
	"log"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

)
type UserController struct {
	UserService data.UserManagment
}

func NewUserController(UserService data.UserManagment) *UserController {
	return &UserController{
		UserService:  UserService,
	}
}

func (uc *UserController)RegisterUserController(c *gin.Context){
	var newUser models.User
	
	err := c.ShouldBindJSON(&newUser)
	log.Printf("New user payload: %+v", newUser)
	newUser.Role ="User"
	newUser.ID = primitive.NewObjectID()
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"invalid payload"})
		return
	}

	err = uc.UserService.RegisterUser(newUser) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated,gin.H{"message" : "User Registered Successfully"})

}


func (uc *UserController)LoginUserController(c *gin.Context){
	var loginDetails struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := c.ShouldBindJSON(&loginDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid payload"})
		return
	}

	if loginDetails.Email == "" || loginDetails.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "email or password cannot be empty"})
		return
	}

	token, err := uc.UserService.LoginUser(loginDetails.Email, loginDetails.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "message": "login successful"})
}