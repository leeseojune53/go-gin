package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"ginjwt/controllers"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/auth", controllers.Login)
	r.POST("/register", controllers.CreateUser)

	return r
}