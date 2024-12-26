package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"newsApi/controllers"
	"newsApi/entitiy"

	"github.com/gorilla/mux"
)

func AddNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var news entitiy.NewsModels
	_ = json.NewDecoder(r.Body).Decode(&news)
	response := controllers.AddNews(news)
	fmt.Print(news)
	json.NewEncoder(w).Encode(response)
}

func DeleteNewsById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	response := controllers.DeleteNewsById(params["newsId"])
	json.NewEncoder(w).Encode(response)
}

func GetAllNews(w http.ResponseWriter, R *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allNews := controllers.GetAllNews()
	json.NewEncoder(w).Encode(allNews)
}

func DeleteAllNews(w http.ResponseWriter, R *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	controllers.DeleteAllNews()
}

func GetNewsById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params["newsId"])
	news := controllers.GetNewsById(params["newsId"])
	json.NewEncoder(w).Encode(news)
}
