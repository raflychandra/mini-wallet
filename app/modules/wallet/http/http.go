package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-mini-wallet/message"
	"go-mini-wallet/utils"
	"net/http"
	"strconv"
)

func (wh *WalletHttp) InitAccount(g *gin.Context) {
	logCtx := fmt.Sprintf("%T.InitAccount", *wh)

	customerId := g.PostForm("customer_xid")

	if !utils.IsValidUUID(customerId) {
		message.Log(logrus.ErrorLevel, "customer id must be uuid type", logCtx, g.Request.Header.Get("CorrelationID"))
		g.JSON(http.StatusBadRequest, message.RenderResponse(nil, "customer id must be uuid type"))
		return
	}

	token, err := wh.WalletUseCase.InitAccount(customerId)
	if err != nil {
		message.Log(logrus.ErrorLevel, err.Error(), logCtx, g.Request.Header.Get("CorrelationID"))
		g.JSON(http.StatusInternalServerError, message.RenderResponse(nil, err.Error()))
		return
	}

	message.Log(logrus.InfoLevel, "success", logCtx, g.Request.Header.Get("CorrelationID"))
	g.JSON(http.StatusOK, message.RenderResponse(
		map[string]interface{}{
			"token": token,
		},
		"success"),
	)
}

func (wh *WalletHttp) EnableMyWallet(g *gin.Context) {
	logCtx := fmt.Sprintf("%T.EnableMyWallet", *wh)

	wallet, err := wh.WalletUseCase.EnableWalled()
	if err != nil {
		message.Log(logrus.ErrorLevel, err.Error(), logCtx, g.Request.Header.Get("CorrelationID"))
		g.JSON(http.StatusInternalServerError, message.RenderResponse(nil, err.Error()))
		return
	}

	message.Log(logrus.InfoLevel, "success", logCtx, g.Request.Header.Get("CorrelationID"))
	g.JSON(http.StatusOK, message.RenderResponse(
		map[string]interface{}{
			"wallet": wallet,
		},
		"success"),
	)
}

func (wh *WalletHttp) ViewMyBalance(g *gin.Context) {
	logCtx := fmt.Sprintf("%T.ViewMyBalance", *wh)

	wallet, err := wh.WalletUseCase.ViewMyBalance()
	if err != nil {
		message.Log(logrus.ErrorLevel, err.Error(), logCtx, g.Request.Header.Get("CorrelationID"))
		g.JSON(http.StatusInternalServerError, message.RenderResponse(nil, err.Error()))
		return
	}

	message.Log(logrus.InfoLevel, "success", logCtx, g.Request.Header.Get("CorrelationID"))
	g.JSON(http.StatusOK, message.RenderResponse(
		map[string]interface{}{
			"wallet": wallet,
		},
		"success"),
	)
}

func (wh *WalletHttp) AddVirtualMoneyToMyWallet(g *gin.Context) {
	logCtx := fmt.Sprintf("%T.AddVirtualMoneyToMyWallet", *wh)

	amount := g.PostForm("amount")
	referenceId := g.PostForm("reference_id")
	wallet, err := wh.WalletUseCase.AddVirtualMoneyToMyWallet(amount, referenceId)
	if err != nil {
		message.Log(logrus.ErrorLevel, err.Error(), logCtx, g.Request.Header.Get("CorrelationID"))
		g.JSON(http.StatusInternalServerError, message.RenderResponse(nil, err.Error()))
		return
	}

	message.Log(logrus.InfoLevel, "success", logCtx, g.Request.Header.Get("CorrelationID"))
	g.JSON(http.StatusOK, message.RenderResponse(
		map[string]interface{}{
			"wallet": wallet,
		},
		"success"),
	)
}

func (wh *WalletHttp) UseVirtualMoneyToMyWallet(g *gin.Context) {
	logCtx := fmt.Sprintf("%T.AddVirtualMoneyToMyWallet", *wh)

	amount := g.PostForm("amount")
	referenceId := g.PostForm("reference_id")
	wallet, err := wh.WalletUseCase.UseVirtualMoneyToMyWallet(amount, referenceId)
	if err != nil {
		message.Log(logrus.ErrorLevel, err.Error(), logCtx, g.Request.Header.Get("CorrelationID"))
		g.JSON(http.StatusInternalServerError, message.RenderResponse(nil, err.Error()))
		return
	}

	message.Log(logrus.InfoLevel, "success", logCtx, g.Request.Header.Get("CorrelationID"))
	g.JSON(http.StatusOK, message.RenderResponse(
		map[string]interface{}{
			"wallet": wallet,
		},
		"success"),
	)
}

func (wh *WalletHttp) DisableMyWallet(g *gin.Context) {
	logCtx := fmt.Sprintf("%T.DisableMyWallet", *wh)

	disabled := g.PostForm("is_disabled")
	boolValue, err := strconv.ParseBool(disabled)
	if err != nil {
		message.Log(logrus.ErrorLevel, err.Error(), logCtx, g.Request.Header.Get("CorrelationID"))
		g.JSON(http.StatusBadRequest, message.RenderResponse(nil, err.Error()))
		return
	}

	wallet, err := wh.WalletUseCase.DisableWallet(boolValue)
	if err != nil {
		message.Log(logrus.ErrorLevel, err.Error(), logCtx, g.Request.Header.Get("CorrelationID"))
		g.JSON(http.StatusInternalServerError, message.RenderResponse(nil, err.Error()))
		return
	}

	message.Log(logrus.InfoLevel, "success", logCtx, g.Request.Header.Get("CorrelationID"))
	g.JSON(http.StatusOK, message.RenderResponse(
		map[string]interface{}{
			"wallet": wallet,
		},
		"success"),
	)
}
