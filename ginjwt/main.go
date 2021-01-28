package main

import (
	"ginjwt/models"
	"ginjwt/routes"
)

func main()  {
	db := models.SetupDB()
	db.AutoMigrate(&models.User{})

	r := routes.SetupRoutes(db)

	r.Run()
}
