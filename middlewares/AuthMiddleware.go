package middlewares

import (
	"TPFINAL-GINCITO/clients"
	"TPFINAL-GINCITO/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	authClient clients.AuthClientInterface
}

func NewAuthMiddleware(authClient clients.AuthClientInterface) *AuthMiddleware {
	return &AuthMiddleware{
		authClient: authClient,
	}
}

func (auth *AuthMiddleware) ValidateToken(c *gin.Context) {

	authToken := c.GetHeader("Authorization")

	if authToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token no encontrado"})
		return
	}
	user, err := auth.authClient.GetUserInfo(authToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	utils.SetUserInContext(c, user)

	c.Next()
}
