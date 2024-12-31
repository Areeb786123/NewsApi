package routes

import (
	"newsApi/services"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("api/loginUser", services.AddNews)
	router.HandleFunc("/newsApi/getAllNews", services.GetAllNews)
	router.HandleFunc("/newsApi/getNews/{newsId}", services.GetNewsById)
	router.HandleFunc("/newsApi/postNews", services.AddNews).Methods("POST")
	router.HandleFunc("/newsApi/deleteById/{newsId}", services.DeleteNewsById).Methods("DELETE")
	router.HandleFunc("/newsApi/validToken/{token}", services.ValidToken)
	return router
}
