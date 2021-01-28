package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"ginjwt/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(c *gin.Context)  {
	var input User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid json provided")
		return
	}

	user := models.User{Username: input.Username, Password: hashing(input.Password)}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&user)

	c.JSON(http.StatusOK, "Create user Success")
}

func hashing(data string) string {
	hash := sha256.New()

	hash.Write([]byte(data))

	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)

	return mdStr
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

	token, err := CreateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, token)

}

func CreateToken(userid uint64) (string, error) {
	var err error

	os.Setenv("ACCESS_SECRET", "testAccessKey")

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return "", err
	}
	return token, nil
}