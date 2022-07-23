package service

import (
	walletHttp "go-mini-wallet/app/modules/wallet/http"
	walletUseCase "go-mini-wallet/app/modules/wallet/usecase"

	"github.com/sirupsen/logrus"
	"go-mini-wallet/config"
	"go-mini-wallet/message"
	"os"
)

type HandlerSetup struct {
	Env             *config.Config
	EnvironmentName *string
	Http            *bool
	WalletUseCase   walletUseCase.InterfaceUseCase
	WalletHttp      walletHttp.InterfaceHttp
}

func MakeHandler(Environment *string, Http *bool) *HandlerSetup {
	loadConfig := config.HandlerLoadConfig{
		Env: *Environment,
	}

	env, err := loadConfig.LoadConfig()
	if err != nil {
		message.Log(logrus.ErrorLevel, err.Error(), "SETUP CONFIG")
		os.Exit(1)
	}

	walletUseCaseClient := walletUseCase.NewUseCaseWallet(env)
	walletHttpClient := walletHttp.NewHttpWallet(env, walletUseCaseClient)

	return &HandlerSetup{
		Env:             env,
		EnvironmentName: Environment,
		Http:            Http,
		WalletUseCase:   walletUseCaseClient,
		WalletHttp:      walletHttpClient,
	}
}
