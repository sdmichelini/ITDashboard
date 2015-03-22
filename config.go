package main

import (
	"encoding/json"
	"os"
)


type Configuration struct{
	//Username to connect to the Database
	Username string `json:"db_username"`
	//DB Password
	Password string `json:"db_password"`
	//Name of the Database
	Name string `json:"db_name"`

}
//Change to your congi file
//See config_example.json for an example
const CONFIG_FILE = "config.json"
/*!
	Load's the Current Configuration
*/
func GetRunningConfiguration() (Configuration, error){
	file, err := os.Open(CONFIG_FILE)
	//Error if we can't open the file
	if err != nil{
		return  Configuration{},err
	}

	//Close when we are done with it
	defer file.Close()

	//Our JSON decoder
	decoder := json.NewDecoder(file)

	//Holder for the Config
	config := Configuration{}

	//Try to Decode JSON
	err = decoder.Decode(&config)

	//Error if we can't decode the JSON
	if err != nil{
		return Configuration{}, err
	}

	return config, nil

}