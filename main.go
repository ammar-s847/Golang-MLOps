package main

import (
	// stl modules
	"fmt"
	"log"
	"net/http"
	// external modules
	"github.com/go-chi/chi/v5"
	// local modules
	"golang-mlops/api"
)

func GetModels(w http.ResponseWriter, r *http.Request) {
	//_, err := w.Write([]byte("Hello World"))
	_, err := fmt.Fprintf(w, "Testing Get Models")
	if err != nil {
		log.Println(err)
	}
}

func GetModelMetadata(w http.ResponseWriter, r *http.Request) {
	modelID := chi.URLParam(r, "model_id")
	_, err := w.Write([]byte("Get metadata for Model with ID: " + modelID))
	if err != nil {
		log.Println(err)
	}
}

func UploadModel(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Testing Upload Model")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	router := chi.NewRouter()

	router.Get("/model", GetModels)
	router.Get("/model/{model_id}", GetModelMetadata)
	router.Post("/model", UploadModel)

	print("Starting API")
	api.RunRouter()

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Println(err)
	}
}
