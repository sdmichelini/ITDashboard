package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
)

type System struct{
	AccessPoints []ItService
}

var services System

func IndexHandler(w http.ResponseWriter, r *http.Request){
	//Load the Templates
	t, err := template.ParseFiles("templates/index.html")

	//Error Loading Template
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}	
	t.Execute(w, services)
}

func routeHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler).Methods("GET")
	return router
}

func main() {
	services = System{AccessPoints:CreateAccessPoints()}
	fmt.Println("Started Webserver")
	fmt.Println("Access Points: ", len(services.AccessPoints))
	//Handle Static Routing for CSS and JS
	http.Handle("/static/", http.StripPrefix("/static/",http.FileServer(http.Dir("static/"))))

	http.Handle("/", routeHandler())
	http.ListenAndServe(":8080",nil)
}