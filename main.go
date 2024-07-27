package main

import (
	"chatroom/internal/controller"
	"chatroom/internal/core/config"
	"chatroom/internal/core/server"
	"chatroom/internal/core/service"
	infraConf "chatroom/internal/infra/config"
	"chatroom/internal/infra/repository"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("****chat room api*****")
	// Create a new instance of the Gin router
	instance := gin.New()
	instance.Use(gin.Recovery())

	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	url := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/chatroom_db?charset=utf8mb4&parseTime=true&loc=UTC&tls=false&readTimeout=3s&writeTimeout=3s&timeout=3s&clientFoundRows=true", username, password)

	// Initialize the database connection
	db, err := repository.NewDB(
		infraConf.DatabaseConfig{
			Driver:                  "mysql",
			Url:                     url,
			ConnMaxLifetimeInMinute: 3,
			MaxOpenConns:            10,
			MaxIdleConns:            1,
		},
	)
	if err != nil {
		log.Fatalf("failed to new database err=%s\n", err.Error())
	}

	// Create the ChatroomRepository
	chatroomRepo := repository.NewChatroomRepository(db)

	// Create the ChatroomService
	chatroomService := service.NewChatroomService(chatroomRepo)

	// Create the chatroomController
	chatroomController := controller.NewChatroomController(instance, chatroomService)

	// Initialize the routes for ChatroomController
	chatroomController.InitRouter()

	// Create the HTTP server
	httpServer := server.NewHttpServer(
		instance,
		config.HttpServerConfig{
			Port: 2009,
		},
	)

	// Start the HTTP server
	httpServer.Start()

}
