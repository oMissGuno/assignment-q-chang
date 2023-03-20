package model

type AmountOfBankNote struct {
	BankNote *float32 `json:"bankNote"`
	Amount   *int     `json:"amount"`
}

type RequestCalculateChange struct {
	PaymentAmount *float32 `json:"paymentAmount"`
	ProductPrice  *float32 `json:"productPrice"`
}
