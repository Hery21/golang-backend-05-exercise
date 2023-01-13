package middlewares

import (
	" hery-ciaputra/assignment-05-golang-backend/httperror"
	"github.com/gin-gonic/gin"
)

type ModelCreator func() any

func RequestValidator(creator ModelCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		model := creator()
		c.Set("payload", model)

		if err := c.ShouldBindJSON(&model); err != nil {
			badRequestError := httperror.BadRequestError(err.Error(), "BAD_REQUEST")
			c.AbortWithStatusJSON(badRequestError.StatusCode, badRequestError)
		}

		c.Set("payload", model)
	}
}
