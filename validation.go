package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

var badWords = map[string]struct{}{
	"kerfuffle": {},
	"sharbert":  {},
	"fornax":    {},
}

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type Chirp struct {
		Body string `json:"body"`
	}

	type Response struct {
		CleanedBody string `json:"cleaned_body"`
	}

	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	chirp := Chirp{}
	err := decoder.Decode(&chirp)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(chirp.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	respondWithJSON(w, http.StatusOK, Response{CleanedBody: getCleanedChirp(chirp.Body, badWords)})
}

func getCleanedChirp(chirp string, badWords map[string]struct{}) string {
	words := strings.Split(chirp, " ")

	for i, word := range words {
		smallWord := strings.ToLower(word)
		if _, ok := badWords[smallWord]; ok {
			words[i] = "****"
		}
	}

	return strings.Join(words, " ")
}
