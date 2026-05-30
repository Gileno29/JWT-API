package http

import (
	"jwt-api/internernal/midware"

	"github.com/gin-gonic/gin"
)

// func CreateRouter() *gin.Engine {
// 	router := gin.Default()

// 	return router
// }

func SetUpRoutesUsers(r *gin.Engine, h *Handller) {
	v1 := r.Group("/api/v1", midware.AuthMiddleware())
	v1.GET("/users", h.getUsers)
	v1.POST("/users", h.postUsers)
	v1.GET("/users/:id", h.getUser)
	v1.PUT("/users/:id", h.putUser)
	v1.DELETE("/users/:id", h.deleteUser)

	public := r.Group("/public")
	public.POST("/login", h.login)
	public.POST("/register", h.register)
	public.POST("email", h.verifyEmail)

}
