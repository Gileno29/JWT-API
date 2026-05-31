package main

import (
	"jwt-api/internernal/database"
	"jwt-api/internernal/http"
	"jwt-api/internernal/repository"
	"jwt-api/internernal/service"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}
	db := database.NewDatabase(os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	conn, err := db.Connect()
	if err != nil {
		panic("erro ao conectar no banco de dados: " + err.Error())
	}

	defer conn.Close()

	userRepository := repository.NewUser(conn)
	userService := service.NewUserService(userRepository)
	handler := http.NewHandler(userService)

	router := gin.Default()
	http.SetUpRoutesUsers(router, handler)

	router.Run(":8090")

}
