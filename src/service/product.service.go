package service

import (
	"context"
	"errors"
	"log"

	"github.com/PetrusAriaa/go-backend-pelatihan-kmteti/src/db"
	"github.com/PetrusAriaa/go-backend-pelatihan-kmteti/src/model"
	"go.mongodb.org/mongo-driver/bson"
)

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ProductResponse struct {
	Data []*Product `json:"data"`
}

func GetAllProduct() (*ProductResponse, error) {
	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}

	coll := db.MongoDB.Collection("product")
	cur, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}

	var prodList []*Product

	for cur.Next(context.TODO()) {
		var prod model.Product
		cur.Decode(&prod)
		prodList = append(prodList, &Product{
			Name:  prod.Name,
			Price: prod.Price,
		})
	}
	return &ProductResponse{
		Data: prodList,
	}, nil
}
