package handlers

import (
	" hery-ciaputra/assignment-05-golang-backend/dto"
	" hery-ciaputra/assignment-05-golang-backend/httperror"
	" hery-ciaputra/assignment-05-golang-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) GetWallet(c *gin.Context) {
	user, _ := c.Get("user")
	ur := user.(models.User)

	wallet, err := h.walletService.GetWallet(ur.ID)
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, wallet)
}

func (h Handler) TopUpWallet(c *gin.Context) {
	user, _ := c.Get("user")
	ur := user.(models.User)
	payload, _ := c.Get("payload")
	topUpReq := payload.(*dto.TopUpReq)

	topUp, err := h.walletService.TopUpWallet(ur.ID, topUpReq)
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, topUp)
}

func (h Handler) TransferWallet(c *gin.Context) {
	user, _ := c.Get("user")
	ur := user.(models.User)
	payload, _ := c.Get("payload")
	transferReq := payload.(*dto.TransferReq)

	transfer, err := h.walletService.TransferWallet(ur.ID, transferReq)
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, transfer)
}

func (h Handler) GetTransactions(c *gin.Context) {
	user, _ := c.Get("user")
	search := c.Query("s")
	sortBy := c.Query("sortBy")
	sort := c.Query("sort")
	limit := c.Query("limit")
	transactionsFilter := &models.Filter{
		Search: search,
		SortBy: sortBy,
		Sort:   sort,
		Limit:  limit,
	}

	ur := user.(models.User)

	transactions, err := h.walletService.GetTransactions(ur.ID, transactionsFilter)
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, transactions)
}
