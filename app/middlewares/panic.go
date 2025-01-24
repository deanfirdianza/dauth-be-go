package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PanicRecovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			var err error
			// Assert the recovered value as an error
			if e, ok := r.(error); ok {
				err = e
			} else {
				// If it's not an error, create a new error
				err = fmt.Errorf("%v", r)
			}

			c.JSON(http.StatusInternalServerError, err.Error())
			c.Abort()
		}
	}()
	c.Next()
}
