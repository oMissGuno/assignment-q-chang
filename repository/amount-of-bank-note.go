package repository

import (
	"assignment/entity"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type AmountOfBankNoteRepository interface {
	FindAll() (*[]entity.AmountOfBankNote, error)
}

type AmountOfBankNoteRepositoryImplement struct {
	DB *mongo.Database
}

func (c AmountOfBankNoteRepositoryImplement) FindAll() (*[]entity.AmountOfBankNote, error) {
	AmountOfBankNoteList := &[]entity.AmountOfBankNote{}
	collection := c.DB.Collection(os.Getenv("DB_COLLECTION"))
	result, err := collection.Find(context.TODO(), bson.D{})
	for result.Next(context.TODO()) {
		AmountOfBankNote := &entity.AmountOfBankNote{}
		if err := result.Decode(&AmountOfBankNote); err != nil {
			return nil, errors.New("decode fail")
		}
		*AmountOfBankNoteList = append(*AmountOfBankNoteList, *AmountOfBankNote)
	}
	if err := result.Err(); err != nil {
		return nil, errors.New("decode fail")
	}
	return AmountOfBankNoteList, err
}
