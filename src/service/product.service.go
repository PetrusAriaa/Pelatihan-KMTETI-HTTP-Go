package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"

	"github.com/PetrusAriaa/go-backend-pelatihan-kmteti/src/db"
	"github.com/PetrusAriaa/go-backend-pelatihan-kmteti/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ProductResponse struct {
	Data []*Product `json:"data"`
}

type ProductRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func GetAllProduct() (*ProductResponse, error) {
	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

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

func CreateProduct(req io.Reader) error {
	var prodReq ProductRequest
	err := json.NewDecoder(req).Decode(&prodReq)
	if err != nil {
		return errors.New("bad request")
	}

	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("product")
	_, err = coll.InsertOne(context.TODO(), model.Product{
		ID:    primitive.NewObjectID(),
		Name:  prodReq.Name,
		Price: prodReq.Price,
		Stock: 0,
	})
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}
	return nil
}
