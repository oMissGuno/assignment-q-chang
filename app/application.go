package app

import (
	"assignment/handler"
	"assignment/repository"
	"assignment/service"
	"go.mongodb.org/mongo-driver/mongo"
)

// App is a dependency container for the api
type App struct {
	Handler Handler
}

type Handler struct {
	//User            handler.UserHandler
	AmountOfBankNote handler.AmountOfBankNoteHandler
}

func InitializeApp(db *mongo.Database) App {
	return App{
		Handler: Handler{
			AmountOfBankNote: handler.AmountOfBankNoteHandlerImplement{
				Service: service.AmountOfBankNoteServiceImplement{Repo: repository.AmountOfBankNoteRepositoryImplement{DB: db}},
			},
		},
	}
}
