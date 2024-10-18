package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock uint8  `json:"stock"`
}

var ProdList []*Product = []*Product{
	{
		Id:    1,
		Name:  "Tepung",
		Price: 15000,
		Stock: 32,
	},
	{
		Id:    2,
		Name:  "Kecap",
		Price: 7500,
		Stock: 12,
	},
}

type ProductRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock uint8  `json:"stock"`
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ProdList)
		return

	case "POST":
		var data ProductRequest

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		p := Product{
			Id:    int(rand.Uint32()),
			Name:  data.Name,
			Price: data.Price,
			Stock: data.Stock,
		}

		ProdList = append(ProdList, &p)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("product added successfully"))

		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
}
