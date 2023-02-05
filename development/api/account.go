package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required , oneof = USD EUR"` //binding is validate field want and oneof is check scope data want or out of
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if errBind := ctx.ShouldBindJSON(&req); errBind != nil {
		//client provide invalid data
		ctx.JSON(http.StatusBadRequest, errorResponse(errBind))
		return

	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, errCreate := server.store.CreateAccount(ctx, arg)
	if errCreate != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errCreate))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
