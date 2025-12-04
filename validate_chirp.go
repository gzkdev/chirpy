package main

import (
	"encoding/json"
	"net/http"
)

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type Chirp struct {
		Body string `json:"body"`
	}

	type Error struct {
		Error string `json:"error"`
	}

	type Response struct {
		Valid bool `json:"valid"`
	}

	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	chirp := Chirp{}
	err := decoder.Decode(&chirp)

	generalError := Error{
		Error: "something went wrong",
	}

	if err != nil {
		w.WriteHeader(500)

		data, err := json.Marshal(generalError)
		if err != nil {
			return
		}
		w.Write(data)
		w.WriteHeader(500)
	}

	if len(chirp.Body) > 140 {
		w.WriteHeader(400)
		error := Error{
			Error: "Chirp is too long",
		}
		data, err := json.Marshal(error)
		if err != nil {
			return
		}
		w.Write(data)
	}

	w.WriteHeader(200)
	response := Response{
		Valid: true,
	}

	data, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(data)
}
