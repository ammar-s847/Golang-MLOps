//package main
//
//import (
//	"fmt"
//	"github.com/go-chi/chi/v5"
//	"log"
//	"net/http"
//)
//
//func GetModel(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Testing GET")
//}
//
//func UploadModel(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Testing POST")
//}
//
//func UpdateModel(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Testing UPDATE")
//}
//
//func DeleteModel(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Testing DELETE")
//}
//
//func handleRequests() {
//	http.HandleFunc("/", GetModel)
//	log.Fatal(http.ListenAndServe(":8081", nil))
//
//}
//
//func main() {
//	handleRequests()
//	print("Hello")
//}

package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func GetModels(w http.ResponseWriter, r *http.Request) {
	//_, err := w.Write([]byte("Hello World"))
	_, err := fmt.Fprintf(w, "Testing Get Models")
	if err != nil {
		log.Println(err)
	}
}

func GetModelMetadata(w http.ResponseWriter, r *http.Request) {
	modelID := chi.URLParam(r, "model_id") // 👈 getting path param
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

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Println(err)
	}
}
