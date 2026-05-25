package midware

import (
	"jwt-api/internernal/auth"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt := auth.NewJWT([]byte(os.Getenv("SECRET_KEY")))

		// Pega o cabeçalho Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			return
		}

		// Divide o cabeçalho para pegar apenas o token (formato: Bearer <token>)
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
			return
		}

		tokenString := parts[1]

		// Valida e faz o parsing do token
		_, err := jwt.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido ou expirado"})
			return
		}

		// Se o token for válido, passa a requisição para o próximo handler
		c.Next()
	}
}
