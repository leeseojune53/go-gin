package controllers

import (
	"ginjwt/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	jwt "github.com/appleboy/gin-jwt/v2"
	"net/http"
	"os"
	"time"
)

type CreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context)  {
	var input LoginUser
	var user models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid json provided")
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}

	if !(input.Password == user.Password) {
		c.JSON(http.StatusBadRequest, "Password did not match")
		return
	}

	c
}

func CreateToken(userid uint64) (string, error) {
	var err error

	os.Setenv("ACCESS_SECRET", "testAccessKey")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true;
	atClaims["user_id"] = userid;
	atClaims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	at := jwt.New(&jwt.GinJWTMiddleware{
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			return atClaims{}
		},
	})
}