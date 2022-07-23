package usecase

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"go-mini-wallet/app/modules/wallet/entity"
	"go-mini-wallet/utils"
	"strconv"
	"time"
)

func (wu *WalletUseCase) InitAccount(customerId string) (string, error) {
	token, err := wu.Encrypt(customerId, wu.Config.Wallet.KeySignature)
	if err != nil {
		return "", err
	}
	entity.CustomerIdGlobal = customerId
	return token, nil
}

func (wu *WalletUseCase) ValidateUser(token string) error {
	decryptStr, err := wu.Decrypt(token, wu.Config.Wallet.KeySignature)
	if err != nil {
		return err
	}

	if decryptStr != entity.CustomerIdGlobal {
		return errors.New("customer not found")
	}
	return nil
}

func (wu *WalletUseCase) EnableWalled() (entity.Wallet, error) {
	if entity.WalletGlobal.Status == "enabled" {
		return entity.WalletGlobal, errors.New("wallet already enable")
	}

	timeNow := time.Now()
	entity.WalletGlobal = entity.Wallet{
		Id:        entity.CustomerIdGlobal,
		OwnedBy:   utils.UUID(),
		Status:    "enabled",
		EnabledAt: &timeNow,
		Balance:   0,
	}
	return entity.WalletGlobal, nil
}

func (wu *WalletUseCase) ViewMyBalance() (*entity.Wallet, error) {
	if entity.WalletGlobal.Status == "disabled" {
		return nil, errors.New("wallet still disabled")
	}
	entity.WalletGlobal.Id = entity.CustomerIdGlobal
	entity.WalletGlobal.OwnedBy = utils.UUID()
	return &entity.WalletGlobal, nil
}

func (wu *WalletUseCase) AddVirtualMoneyToMyWallet(amount, referenceId string) (*entity.Deposits, error) {
	if entity.WalletGlobal.Status == "disabled" {
		return nil, errors.New("wallet still disabled")
	}

	amountInt, err := strconv.Atoi(amount)
	if err != nil {
		return nil, err
	}

	var deposits = entity.Deposits{
		Id:          entity.CustomerIdGlobal,
		DepositedBy: utils.UUID(),
		Status:      "success",
		DepositedAt: time.Now(),
		Amount:      amountInt,
		ReferenceId: referenceId,
	}
	entity.WalletGlobal.Balance = entity.WalletGlobal.Balance + amountInt
	return &deposits, nil
}

func (wu *WalletUseCase) UseVirtualMoneyToMyWallet(amount, referenceId string) (*entity.WithDrawls, error) {
	if entity.WalletGlobal.Status == "disabled" {
		return nil, errors.New("wallet still disabled")
	}

	amountInt, err := strconv.Atoi(amount)
	if err != nil {
		return nil, err
	}

	if amountInt > entity.WalletGlobal.Balance {
		return nil, errors.New("insufficient balance")
	}

	var withDrawls = entity.WithDrawls{
		Id:          entity.CustomerIdGlobal,
		WithdrawnBy: utils.UUID(),
		Status:      "success",
		WithdrawnAt: time.Now(),
		Amount:      amount,
		ReferenceId: referenceId,
	}

	entity.WalletGlobal.Balance = entity.WalletGlobal.Balance - amountInt
	return &withDrawls, nil
}

func (wu *WalletUseCase) DisableWallet(disabled bool) (*entity.Wallet, error) {
	if entity.WalletGlobal.Status == "disabled" {
		return nil, errors.New("wallet already disable")
	}

	if disabled {
		timeNow := time.Now()
		entity.WalletGlobal = entity.Wallet{
			Id:         entity.CustomerIdGlobal,
			OwnedBy:    utils.UUID(),
			Status:     "disabled",
			DisabledAt: &timeNow,
			Balance:    0,
		}
		return &entity.WalletGlobal, nil
	}

	return nil, nil
}

func (wu *WalletUseCase) Encrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, entity.Bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return utils.Encode(cipherText), nil
}

func (wu *WalletUseCase) Decrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}

	cipherText, err := utils.Decode(text)
	if err != nil {
		return "", err
	}

	cfb := cipher.NewCFBDecrypter(block, entity.Bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
