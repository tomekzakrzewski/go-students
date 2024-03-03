package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tomekzakrzewski/go-students/db"
	"github.com/tomekzakrzewski/go-students/handlers"
	"github.com/tomekzakrzewski/go-students/middleware"

	_ "github.com/tomekzakrzewski/go-students/docs"
)

//	@title			Student CRUD
//	@version		1.0
//	@description	Student CRUD

// @host		localhost:8080
// @BasePath	/api
func main() {
	database, err := db.DBconnect()
	if err != nil {
		panic("Failed to connect to database")
	}
	sqliteDB, err := database.DB()
	if err != nil {
		panic("Failed to connect to database")
	}
	defer sqliteDB.Close()

	var (
		studentStore   = db.NewStudentStore(database)
		studentHandler = handlers.NewStudentHandler(studentStore)
		r              = gin.Default()
		api            = r.Group("/api")
		listenAddr     = os.Getenv("LISTEN_ADDR")
	)

	api.Use(middleware.ErrorHandler())
	api.POST("/students", studentHandler.HandleAddStudent)
	api.GET("/students", studentHandler.HandleGetStudents)
	api.GET("/students/:id", studentHandler.HandleGetStudent)
	api.PUT("/students/:id", studentHandler.HandleUpdateStudent)
	api.DELETE("/students/:id", studentHandler.HandleDeleteStudent)

	// swaggger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(listenAddr)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
