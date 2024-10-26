package main

import (
	"fmt"
	"net/http"

	handler "github.com/PetrusAriaa/go-backend-pelatihan-kmteti/api"
)

func main() {
	h := http.NewServeMux()

	s := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}

	h.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	h.HandleFunc("/api/product", handler.ProductHandler)

	fmt.Println("HTTP Server running on port 8080")
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
