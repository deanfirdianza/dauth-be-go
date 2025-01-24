package middlewares

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user repositories and drivers

		// Get user object

		// ...handle auth logic...
		c.Set("user", "example_user")
		c.Set("email", "example@gmail.com")
		c.Next()
	}
}
