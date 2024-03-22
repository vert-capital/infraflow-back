package middleware

import (
	usecase_user "app/usecase/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticatedMiddleware(usercase usecase_user.IUsecaseUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get bearer token from header
		token := c.Request.Header.Get("Authorization")

		user, err := usercase.GetUserByToken(token)

		// check if token is valid
		if err == nil {

			// set user to context
			c.Set("user", *user)

			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
		}
	}
}

func AdminMiddleware(usercase usecase_user.IUsecaseUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get bearer token from header
		token := c.Request.Header.Get("Authorization")

		// token := strings.Split(bearerToken, " ")[1]

		user, err := usercase.GetUserByToken(token)

		// check if token is valid
		if err == nil {

			if !user.IsAdmin {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
				return
			}

			// set user to context
			c.Set("user", *user)

			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
		}
	}
}
