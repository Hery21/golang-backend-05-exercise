package handlers

import (
	"net/http"

	" hery-ciaputra/assignment-05-golang-backend/dto"
	" hery-ciaputra/assignment-05-golang-backend/httperror"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignIn(c *gin.Context) {
	value, _ := c.Get("payload")
	signInReq := value.(*dto.SignInReq)

	result, err := h.authService.SignIn(signInReq)
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, result)
}
