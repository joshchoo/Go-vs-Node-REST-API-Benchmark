package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type numbers struct {
	A int `json:"a"`
	B int `json:"b"`
}

type result struct {
	Result int `json:"result"`
}

func main() {
	r := chi.NewRouter()
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var nums numbers
		json.NewDecoder(r.Body).Decode(&nums)
		sum := nums.A + nums.B
		res := result{
			Result: sum,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	})

	http.ListenAndServe(":8080", r)
}
