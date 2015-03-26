package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	
)

var services System
var config Configuration

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

func AccessTokenGetHandler(w http.ResponseWriter, r *http.Request){
	b, err := CreateSessionId()

	if err != nil{
		http.Error(w, "err", 500)
		return
	}

	u := AccessCode{AccessToken: b, Expires: time.Now().UTC().Add(24 * time.Hour) }

	buf, err := json.Marshal(u)

	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	fmt.Fprintf(w, "%s",buf)

}

func routeHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler).Methods("GET")
	router.HandleFunc("/services", GlobalServiceHandler).Methods("GET")
	router.HandleFunc("/register", RegisterGetHandler).Methods("GET")
	router.HandleFunc("/auth", AccessTokenGetHandler).Methods("GET")

	return router
}

func main() {
	con,err := GetRunningConfiguration()
	if err != nil {
		fmt.Println("Error Loading Configuration at: %s",CONFIG_FILE)
	}else{
		fmt.Println("Loaded Config.")
	}

	config = con



	fmt.Println(config.Username)
	db, err2 := sql.Open("mysql",fmt.Sprintf("%s:%s@/%s",config.Username,config.Password,config.Name))
	defer db.Close()
	if err2 != nil {
		fmt.Println("Error: Couldn't Open Database")
	}else{
		_, err2 := AuthenticateUser("admin","1234",config,db)

		if err2!=nil{
			fmt.Println("Authenication Issue")
		}
	}

	



	services = CreateSystem()

	fmt.Println("Started Webserver")
	fmt.Println("Services: ", len(services.Services))
	//Handle Static Routing for CSS and JS
	http.Handle("/static/", http.StripPrefix("/static/",http.FileServer(http.Dir("static/"))))

	http.Handle("/", routeHandler())
	http.ListenAndServe(":8080",nil)
}