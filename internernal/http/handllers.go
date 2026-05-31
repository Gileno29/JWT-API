package http

import (
	"jwt-api/internernal/models"
	"jwt-api/internernal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handller struct {
	UserService *service.UserService
}

func NewHandler(userService *service.UserService) *Handller {
	return &Handller{UserService: userService}
}

func (h *Handller) getUsers(c *gin.Context) {
	users, err := h.UserService.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if len(users) == 0 {
		c.JSON(200, gin.H{"message": "Nenhum usuário encontrado"})
		return
	}
	c.JSON(200, users)

}

func (h *Handller) postUsers(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = h.UserService.Register(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, user)

}

func (h *Handller) getUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	user, err := h.UserService.GetUser(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)

}

func (h *Handller) putUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user.ID = id
	err = h.UserService.UpdateUser(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)

}

func (h *Handller) deleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var user *models.User
	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user.ID = id
	err = h.UserService.DeleteUser(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)

}

func (h *Handller) login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := h.UserService.Login(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Login realizado com sucesso", "token": token})

}

func (h *Handller) register(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = h.UserService.Register(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"status": "success", "user": user.Name})

}

// func (h *Handller) verifyEmail(c *gin.Context) {
// 	var user *models.User
// 	err := c.ShouldBindJSON(&user)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}
// 	user, err = h.UserService.GetUserByEmail(user.Email)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, user)

// }
