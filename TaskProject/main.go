package main

import(
	"TaskProject/models"
	"TaskProject/routes"
)

func main()  {

	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)
	r.Run()
	
}