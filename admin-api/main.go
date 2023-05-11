package main

import (
	"admin-api/config"
	"admin-api/controllers"
	"admin-api/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	config.EnvLoad()
}

func main() {
	con := config.ConnectDB()
	app := gin.Default()
	e := controllers.NewEmployeeConroller(con)

	app.POST("/login", e.Login)
	app.Use(middleware.Authenticate())
	employee := app.Group("/employee")
	employee.GET("/all", e.GetAllEmployee)
	employee.GET("/:id", e.GetEmployeeByID)
	employee.POST("/create", e.CreateEmployee)
	employee.POST("/update/:id", e.UpdateEmployee)
	employee.POST("/delete/:id", e.RemoveEmployee)

	app.Run(":" + os.Getenv("PORT"))
}
