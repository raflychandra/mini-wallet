package entity

import "time"

var Bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

type Wallet struct {
	Id         string     `json:"id"`
	OwnedBy    string     `json:"owned_by"`
	Status     string     `json:"status"`
	EnabledAt  *time.Time `json:"enabled_at,omitempty"`
	DisabledAt *time.Time `json:"disabled_at,omitempty"`
	Balance    int        `json:"balance"`
}

type Deposits struct {
	Id          string    `json:"id"`
	DepositedBy string    `json:"deposited_by"`
	Status      string    `json:"status"`
	DepositedAt time.Time `json:"deposited_at"`
	Amount      int       `json:"amount"`
	ReferenceId string    `json:"reference_id"`
}

type WithDrawls struct {
	Id          string    `json:"id"`
	WithdrawnBy string    `json:"withdrawn_by"`
	Status      string    `json:"status"`
	WithdrawnAt time.Time `json:"withdrawn_at"`
	Amount      string    `json:"amount"`
	ReferenceId string    `json:"reference_id"`
}
