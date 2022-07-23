package entity

var (
	CustomerIdGlobal = ""
	WalletGlobal     = Wallet{
		Id:         "",
		OwnedBy:    "",
		Status:     "disabled",
		EnabledAt:  nil,
		DisabledAt: nil,
		Balance:    0,
	}
)
