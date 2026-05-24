package http

import (
	"github.com/gin-gonic/gin"
)

// func CreateRouter() *gin.Engine {
// 	router := gin.Default()

// 	return router
// }

func SetUpRoutesUsers(r *gin.Engine, h *Handller) *gin.RouterGroup {
	v1 := r.Group("/api/v1")
	v1.GET("/users", h.getUsers)
	v1.POST("/users", h.postUsers)
	v1.GET("/users/:id", h.getUser)
	v1.PUT("/users/:id", h.putUser)
	v1.DELETE("/users/:id", h.deleteUser)

	v1.POST("/login", h.login)
	v1.POST("/register", h.register)

	v1.POST("email", h.verifyEmail)

	return v1

}
