package handlers

import (
	"fmt"
	"mp-service/internal/auth"
	"mp-service/internal/models"
	"net/http"
	"regexp"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *auth.AuthService
}

func NewAuthHandler(service *auth.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (ah *AuthHandler) Login(c *gin.Context) {

}

func (ah *AuthHandler) Register(c *gin.Context) {
	regReq := models.RegistrationUserReq{}
	err := c.ShouldBindJSON(&regReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//TODO добавить проверку валидности пароля

	//проверка на валидность введенных данных
	if !isEmailValid(regReq.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrectly entered email address"})
		return
	}
	//проверка зарегистрирован ли пользователь с таким емайл уже?
	_, err = ah.service.FindUserByEmail(regReq.Email)
	if err != nil {
		fmt.Println(err.Error())
		if err.Error() != "user with that email not found" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user with that email already exist"})
			return
		}
	}

	err = ah.service.CreateUser(regReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//generate token
	payload := jwt.MapClaims{
		"sub": regReq.Email,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString([]byte("jwtSecret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating jwt token"})
		return
	}

	c.Header("Authorization", "Bearer "+t)
	c.JSON(http.StatusOK, gin.H{"message": "user was succesfully created!", "user_info": regReq})

}

func (ah *AuthHandler) Logout(c *gin.Context) {

}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}
