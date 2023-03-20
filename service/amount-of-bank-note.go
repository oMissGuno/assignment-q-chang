package service

import (
	"assignment/model"
	"assignment/repository"
	"assignment/utils"
	"errors"
)

type AmountOfBankNoteService interface {
	AmountOfBankNoteProcess(paymentAmount float32, productPrice float32) (*[]model.AmountOfBankNote, error)
}

type AmountOfBankNoteServiceImplement struct {
	Repo repository.AmountOfBankNoteRepository
}

func (d AmountOfBankNoteServiceImplement) AmountOfBankNoteProcess(paymentAmount float32, productPrice float32) (*[]model.AmountOfBankNote, error) {
	if paymentAmount < productPrice {
		return nil, errors.New("paymentAmount < productPrice")
	}
	totalAmount := paymentAmount - productPrice
	dataList, err := d.Repo.FindAll()
	if err != nil {
		return nil, err
	}
	changeTotal := utils.Sum(dataList)
	if totalAmount > changeTotal {
		return nil, errors.New("totalAmount > changeTotal")
	}
	responses := &[]model.AmountOfBankNote{}
	for _, data := range *dataList {
		temp := &model.AmountOfBankNote{}
		bank, amount, change := utils.CountBankNote(*data.BankNote, *data.Amount, totalAmount)
		temp.Amount = &amount
		temp.BankNote = &bank
		totalAmount = change
		*responses = append(*responses, *temp)
	}
	return responses, nil
}
