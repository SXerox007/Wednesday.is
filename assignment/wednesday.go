package main

import (
	"net/http"

	"Wednesday.is/base/utils"
)

//X -> 10
//XI -> 11
//I -> 1
// roman to int

type RomanToIntergerRequest struct {
	Value string `json:"value"`
}

func RomanToInterger() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := &RomanToIntergerRequest{}
		utils.ParseBody(r, body)
		if body.Value == "" {
			respondWithJSON(w, http.StatusBadRequest, CommonResponse{false, "Value Field is missing.", http.StatusBadRequest, nil})
		}
		if val := utils.RomanToIntergerConverter(body.Value); val == -1 {
			// error
			respondWithJSON(w, http.StatusBadRequest, CommonResponse{false, "Invalid Entry.", http.StatusBadRequest, nil})
		} else {
			// success
			respondWithJSON(w, http.StatusOK, CommonResponse{true, "Success", http.StatusOK, val})
		}
		return
	}
}
