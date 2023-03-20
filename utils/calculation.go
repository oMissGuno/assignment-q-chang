package utils

import (
	"assignment/entity"
)

func Sum(array *[]entity.AmountOfBankNote) float32 {
	var result float32
	for _, v := range *array {
		result += *v.BankNote * float32(*v.Amount)
	}
	return result
}

func CountBankNote(bank float32, amount int, changeTotal float32) (float32, int, float32) {
	countBank := 0
	for i := amount; i > 0; i-- {
		if changeTotal < bank {
			return bank, countBank, changeTotal
		}
		changeTotal = changeTotal - bank
		countBank++
	}
	return bank, countBank, changeTotal
}
