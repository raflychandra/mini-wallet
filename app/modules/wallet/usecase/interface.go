package usecase

import (
	"go-mini-wallet/app/modules/wallet/entity"
	"go-mini-wallet/config"
)

type WalletUseCase struct {
	Config *config.Config
}

func NewUseCaseWallet(e *config.Config) InterfaceUseCase {
	return &WalletUseCase{
		Config: e,
	}
}

type InterfaceUseCase interface {
	InitAccount(customerId string) (string, error)
	ValidateUser(token string) error
	EnableWalled() (entity.Wallet, error)
	ViewMyBalance() (*entity.Wallet, error)
	AddVirtualMoneyToMyWallet(amount, referenceId string) (*entity.Deposits, error)
	UseVirtualMoneyToMyWallet(amount, referenceId string) (*entity.WithDrawls, error)
	DisableWallet(disabled bool) (*entity.Wallet, error)
	Encrypt(str, passphrase string) (string, error)
	Decrypt(str, passphrase string) (string, error)
}
