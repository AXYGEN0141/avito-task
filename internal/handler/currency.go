package handler

import (
	"avito-task/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCurrency adds new currency to account.
func (h *Handler) CreateCurrency(c *gin.Context) {

	var (
		err     error
		account model.Account
	)

	if err = c.ShouldBindJSON(&account); err != nil {
		responseWithStatus(c, http.StatusBadRequest, nil, "Error", err.Error())
		return
	}

	err = h.currencyService.Create(account.Currency)
	if err != nil {
		responseWithStatus(c, http.StatusInternalServerError, nil, "Error", err.Error())
		return
	}
	responseWithStatus(c, http.StatusOK, account.Currency, "Success", "create currency")
}
