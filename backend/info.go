package main

import (
	"do-tf-stack-builder/utils"
	"encoding/json"

	"log"
	"net/http"
)

type VersionResponse struct {
	Version string `json:"version"`
}

type HealthResponse struct {
	Status int `json:"status"`
}

func Version(w http.ResponseWriter, r *http.Request) {
	response := VersionResponse{
		Version: version,
	}

	_, err := json.MarshalIndent(&response, "", " ")
	if err != nil {
		log.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}

func Health(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status: http.StatusOK,
	}

	_, err := json.MarshalIndent(&response, "", " ")
	if err != nil {
		log.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}
