package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @apiDefine CustomError
 *
 * @apiErrorExample 500 Server Error
 *     HTTP/1.1 500 Internal Server Error
 *     {
 *        "error" : "Not Found"
 *     }
 *
 */

// ICustomError define un error con Code y Error
type ICustomError interface {
	Code() int
	Error() string
}

// ErrorHandler a middleware to handle errors
func ErrorHandler(c *gin.Context) {
	c.Next()

	handleErrorIfNeeded(c)
}

func handleErrorIfNeeded(c *gin.Context) {
	err := c.Errors.Last()
	if err == nil {
		return
	}

	switch value := err.Err.(type) {
	case ICustomError:
		c.JSON(value.Code(),
			gin.H{
				"error": value.Error(),
			})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
}
