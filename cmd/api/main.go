package main

import (
	"jwt-api/internernal/database"
	"jwt-api/internernal/http"
	"jwt-api/internernal/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	db := database.NewDatabase("localhost", 5432, "postgres", "postgres", "postgres")
	conn, err := db.Connect()
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(); err != nil {
		panic(err)
	}

	defer conn.Close()

	userRepository := repository.NewUser(conn)
	handler := http.NewHandler(userRepository)

	router := gin.Default()
	http.SetUpRoutesUsers(router, handler)

	router.Run(":8090")

}
