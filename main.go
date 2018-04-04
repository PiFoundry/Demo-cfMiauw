package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type tContent struct {
	Title         string
	Head          string
	Message       string
	InstanceID    string
	InstanceIndex string
	CpuModel      string
}

var c tContent

func getCpuModel() string {
	cpuInfoBytes, _ := ioutil.ReadFile("/proc/cpuinfo")
	cpuInfo := fmt.Sprintf("%s", cpuInfoBytes)
	cpuInfoLines := strings.Split(cpuInfo, "\n")
	for _, cpuInfoLine := range cpuInfoLines {
		if strings.HasPrefix(cpuInfoLine, "model name") {
			return cpuInfoLine[strings.Index(cpuInfoLine, ":")+1:]
		}
	}
	return ""
}

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
	c.Title = "cfMiauw - CatFoundry Demo"
	c.Head = "CF Summit Boston - Running at (small) scale!"
	c.InstanceIndex = os.Getenv("CF_INSTANCE_INDEX")
	c.CpuModel = getCpuModel()

	router := mux.NewRouter()
	router.HandleFunc("/", vMiauw)
	http.Handle("/", router)
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
	if err != nil {
		fmt.Println(err)
	}
}
