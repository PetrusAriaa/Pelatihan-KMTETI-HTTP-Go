package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/PetrusAriaa/go-backend-pelatihan-kmteti/src/db"
	"github.com/PetrusAriaa/go-backend-pelatihan-kmteti/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock uint8  `json:"stock"`
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		db, err := db.DBConnection()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		coll := db.MongoDB.Collection("product")
		cur, err := coll.Find(context.TODO(), bson.D{})
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		var prodList []*model.Product

		for cur.Next(context.TODO()) {
			var prod model.Product
			cur.Decode(&prod)
			prodList = append(prodList, &prod)
		}

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(prodList)
		return

	case "POST":
		var data ProductRequest

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		db, err := db.DBConnection()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		coll := db.MongoDB.Collection("product")

		_, err = coll.InsertOne(context.TODO(), model.Product{
			ID:    primitive.NewObjectID(),
			Name:  data.Name,
			Price: data.Price,
			Stock: int(data.Stock),
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Data added successfully"))

		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
}
