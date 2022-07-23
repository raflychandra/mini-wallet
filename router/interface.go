package router

import (
	"github.com/gin-gonic/gin"
	"go-mini-wallet/service"
)

type HandlerRouter struct {
	Setup *service.HandlerSetup
}

func NewHandlerRouter(setup *service.HandlerSetup) InterfaceRouter {
	return &HandlerRouter{
		Setup: setup,
	}
}

type InterfaceRouter interface {
	CORSMiddleware() gin.HandlerFunc
	ListRouter() *gin.Engine
	Middleware() gin.HandlerFunc
}
