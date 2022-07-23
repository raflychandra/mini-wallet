package router

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-mini-wallet/message"
	"go-mini-wallet/utils"
	"net/http"
	"strings"
)

func (hr *HandlerRouter) ListRouter() *gin.Engine {
	g := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	g.Use(hr.CORSMiddleware())

	apiV1 := g.Group("/api/v1")
	apiV1.POST("/init", hr.Setup.WalletHttp.InitAccount)

	walletApi := apiV1.Group("/wallet", hr.Middleware())
	walletApi.POST("", hr.Setup.WalletHttp.EnableMyWallet)
	walletApi.GET("", hr.Setup.WalletHttp.ViewMyBalance)
	walletApi.POST("/deposits", hr.Setup.WalletHttp.AddVirtualMoneyToMyWallet)
	walletApi.POST("/withdrawals", hr.Setup.WalletHttp.UseVirtualMoneyToMyWallet)
	walletApi.PATCH("", hr.Setup.WalletHttp.DisableMyWallet)
	return g
}

func (hr *HandlerRouter) CORSMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		g.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		g.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		g.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		g.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if g.Request.Method == "OPTIONS" {
			g.AbortWithStatus(204)
			return
		}

		correlationID := g.Request.Header.Get("CorrelationID")
		if strings.TrimSpace(correlationID) == "" {

			g.Request.Header.Add("CorrelationID", utils.UUID())
		}
		g.Next()
	}
}

func (hr *HandlerRouter) Middleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		auth := g.Request.Header.Get("Authorization")
		arr := strings.Split(auth, " ")
		if arr[0] != "Token" {
			g.JSON(http.StatusUnauthorized, message.RenderResponse(nil, errors.New("required Authorization").Error()))
			g.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		err := hr.Setup.WalletUseCase.ValidateUser(arr[1])
		if err != nil {
			g.JSON(http.StatusUnauthorized, message.RenderResponse(nil, errors.New("customer not initialization").Error()))
			g.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		g.Next()
	}
}
