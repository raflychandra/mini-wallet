package http

import (
	walletUseCase "go-mini-wallet/app/modules/wallet/usecase"

	"github.com/gin-gonic/gin"
	"go-mini-wallet/config"
)

type WalletHttp struct {
	Config        *config.Config
	WalletUseCase walletUseCase.InterfaceUseCase
}

func NewHttpWallet(e *config.Config, walletUseCaseClient walletUseCase.InterfaceUseCase) InterfaceHttp {
	return &WalletHttp{
		Config:        e,
		WalletUseCase: walletUseCaseClient,
	}
}

type InterfaceHttp interface {
	InitAccount(g *gin.Context)
	EnableMyWallet(g *gin.Context)
	ViewMyBalance(g *gin.Context)
	AddVirtualMoneyToMyWallet(g *gin.Context)
	UseVirtualMoneyToMyWallet(g *gin.Context)
	DisableMyWallet(g *gin.Context)
}
