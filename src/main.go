package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Routes struct {
	Routes []Route `json:"routes"`
}

type Route struct {
	Path       string `json:"path"`
	Controller string `json:"controller"`
	Method     string `json:"method"`
	Auth       bool   `json:"auth"`
	Parameter  bool   `json:"parameter"`
}

type Databases struct {
	Databases []Database `json:"databases"`
}

type Database struct {
	Type     string `json:"type"`
	Driver   string `json:"driver"`
	DbName   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

func createPath(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
}

func createFile(file string, content []byte) {
	d := content
	ioutil.WriteFile(file, d, 0755)
}

func readFile(file string) ([]byte, error) {
	jsonFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}

	output, err := ioutil.ReadAll(jsonFile)

	defer jsonFile.Close()

	return output, err
}

func init() {
	createPath("build")
	createPath("build/models")
	createPath("build/handlers")
	createPath("build/database")
	createPath("build/controllers")
	createFile("build/main.go", []byte(""))
}

func main() {
	configs, err := readFile("config.json")

	if err != nil {
		fmt.Println(err)
	}

	var database Databases
	var routes Routes

	json.Unmarshal(configs, &database)
	json.Unmarshal(configs, &routes)

	fmt.Println(database)
	fmt.Println(routes)

	// fmt.Println(string(configs))

	// for i := 0; i < len(routes.Routes); i++ {
	// 	fmt.Println("Path: " + routes.Routes[i].Path)
	// 	fmt.Println("Controller: " + routes.Routes[i].Controller)
	// 	fmt.Println("Method: " + routes.Routes[i].Method)
	// }
}
