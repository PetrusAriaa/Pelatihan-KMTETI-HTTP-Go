package handler

import "net/http"

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from product"))
}
