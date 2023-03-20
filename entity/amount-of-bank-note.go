package entity

type AmountOfBankNote struct {
	BankNote *float32 `bson:"bankNote"`
	Amount   *int     `bson:"amount"`
}
