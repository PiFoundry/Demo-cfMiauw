package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/vchrisr/vMiauw/config"

	"github.com/gorilla/mux"
)

type tContent struct {
	Title         string
	Head          string
	Message       string
	InstanceID    string
	InstanceIndex string
}

var c tContent

func vMiauw(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("vmtemplate").ParseFiles("index.tmpl")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "index.tmpl", c)
}

func main() {
	c.Message = "This CatFoundry Demo instance is running on:"
	c.InstanceID = os.Getenv("CF_INSTANCE_GUID")
	c.Title = "vMiauw - CatFoundry Demo"
	c.Head = "Hello RWS!"
	c.InstanceIndex = os.Getenv("CF_INSTANCE_INDEX")

	router := mux.NewRouter()
	router.HandleFunc("/", vMiauw)
	http.Handle("/", router)
	err := http.ListenAndServe(":"+config.GetPort(), nil)
	if err != nil {
		fmt.Println(err)
	}
}
