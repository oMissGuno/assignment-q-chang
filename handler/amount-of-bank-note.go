package handler

import (
	"assignment/model"
	"assignment/service"
	"github.com/labstack/echo/v4"
)

type AmountOfBankNoteHandler interface {
	CalculateChange(c echo.Context) (*[]model.AmountOfBankNote, error)
}

type AmountOfBankNoteHandlerImplement struct {
	Service service.AmountOfBankNoteService
}

func (d AmountOfBankNoteHandlerImplement) CalculateChange(c echo.Context) (*[]model.AmountOfBankNote, error) {
	var request *model.RequestCalculateChange
	c.Bind(&request)
	resp, err := d.Service.AmountOfBankNoteProcess(*request.PaymentAmount, *request.ProductPrice)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
