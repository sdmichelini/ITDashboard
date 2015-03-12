package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var services System

func IndexHandler(w http.ResponseWriter, r *http.Request){
	//Load the Static HTML files
	file, err := ioutil.ReadFile("html/index.html")

	//Error Loading Template
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}	
	
	fmt.Fprintf(w, "%s",file)
}

func GlobalServiceHandler(w http.ResponseWriter, r *http.Request){
	buf, err := json.Marshal(services)

	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	fmt.Fprintf(w, "%s",buf)
}

func RegisterGetHandler(w http.ResponseWriter, r *http.Request){
	//Load the Static HTML files
	file, err := ioutil.ReadFile("html/register.html")

	//Error Loading Template
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}	
	
	fmt.Fprintf(w, "%s",file)
}

func routeHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler).Methods("GET")
	router.HandleFunc("/services", GlobalServiceHandler).Methods("GET")
	router.HandleFunc("/register", RegisterGetHandler).Methods("GET")
	return router
}

func main() {
	services = CreateSystem()

	fmt.Println("Started Webserver")
	fmt.Println("Services: ", len(services.Services))
	//Handle Static Routing for CSS and JS
	http.Handle("/static/", http.StripPrefix("/static/",http.FileServer(http.Dir("static/"))))

	http.Handle("/", routeHandler())
	http.ListenAndServe(":8080",nil)
}