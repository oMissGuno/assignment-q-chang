package main

import (
	"assignment/app"
	"assignment/database"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	e := echo.New()
	_, err := ioutil.ReadFile(".env")
	if err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println("Error loading .env file: ", err)
		}
	}
	mongodb, errConnectDb := database.OpenDB(context.Background(), os.Getenv("DB_URI"), os.Getenv("DB_NAME"))
	if errConnectDb != nil {
		fmt.Println("Can not connect DB", errConnectDb)
	}
	app := app.InitializeApp(mongodb)
	e.POST("/calculate/change", func(c echo.Context) error {
		resp, err := app.Handler.AmountOfBankNote.CalculateChange(c)
		if err != nil {
			return c.JSONPretty(http.StatusBadRequest, err.Error(), "  ")
		}
		return c.JSONPretty(http.StatusOK, resp, "  ")
	},
	)
	e.Logger.Fatal(e.Start(":8080"))
}
