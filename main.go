package main

import (
	"fmt"
	"net/http"
	dbconnection "newsApi/dbConnection"
	"newsApi/routes"
)

func main() {
	dbconnection.ConnectDB()
	r := routes.Router()
	fmt.Println("servering to port 4000")
	err := http.ListenAndServe(":4000", r)
	if err != nil {
		panic("some error occur while connecting to servee")
	}
	fmt.Println("servering to port 44545")

}
