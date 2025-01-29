package middlewares

import (
	"net/http"

	"github.com/deanfirdianza/dauth-be-go/app/env"
	repository "github.com/deanfirdianza/dauth-be-go/modules/auths/v1/repositories"
	service "github.com/deanfirdianza/dauth-be-go/modules/auths/v1/services"
	userRepository "github.com/deanfirdianza/dauth-be-go/modules/users/v1/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Auth(conf env.Conf, DBSqlx *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get database connection

		// Get user repositories and drivers
		AuthRepo := repository.NewAuthRepository(DBSqlx)
		UserRepo := userRepository.NewUserRepository(DBSqlx)
		AuthService := service.NewAuthService(conf.App.Secret_key, AuthRepo, UserRepo)

		tokenString, err := c.Cookie("DAT")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Access token not found"})
			c.Abort()
			return
		}
		// Get user object
		claims, err := AuthService.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		uid, ok := claims["uid"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
			c.Abort()
			return
		}

		// ...handle auth logic...
		c.Set("user_id", uid)
		c.Next()
	}
}
