package handlers

import (
	"fmt"
	"mp-service/internal/models"
	auth "mp-service/internal/service"
	"mp-service/internal/utils"
	"net/http"
	"regexp"
	"time"
	"unicode"

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

	tokenHeader := c.GetHeader("Authorization")
	if tokenHeader != "" {
		isExpired, err := utils.IsTokenExpired(tokenHeader)
		if err == nil && !isExpired {
			_, err := utils.GetSubFromToken(tokenHeader)
			if err == nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": fmt.Sprintf("you are already authorized"),
				})
				return
			}
		}
	}

	authReq := models.AuthUserReq{}
	if err := c.ShouldBindJSON(&authReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong input data"})
		return
	}
	//передаем в сервис полученные данные, там проверяем соответствие данным из бд
	err := ah.service.AuthorizeUser(authReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Генерим сам токен и возвращаем сам токен ответом
	payload := jwt.MapClaims{
		"sub": authReq.Email,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString([]byte("jwtSecret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating jwt token"})
		return
	}

	c.Header("Authorization", "Bearer "+t)
	c.JSON(http.StatusOK, gin.H{"message": "you succesfully authorize!", "token": t})
}

func (ah *AuthHandler) Register(c *gin.Context) {
	regReq := models.AuthUserReq{}
	err := c.ShouldBindJSON(&regReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//проверка на валидность введенных данных
	if !isEmailValid(regReq.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrectly entered email address"})
		return
	}
	if !isPasswordValid(regReq.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the password is too simple it must contain at least 6 characters and 1 digit"})
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
	if len(c.GetHeader("Authorization")) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "you are not authorize"})
		return
	}

	c.Header("Authorization", "")
	c.JSON(http.StatusOK, gin.H{"message": "you logout succesfully!"})
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

func isPasswordValid(p string) bool {
	if len(p) < 6 {
		return false
	}
	digitFlag := false
	chFlag := false
	for _, ch := range p {
		if unicode.IsDigit(ch) {
			digitFlag = true
		}
		if unicode.IsLower(ch) || unicode.IsUpper(ch) {
			chFlag = true
		}
	}
	return digitFlag && chFlag
}
