package config

import (
	"os"
	"strconv"

	"github.com/cloudfoundry-community/go-cfenv"
)

//GetPort : Get TCP PORT from CF env
func GetPort() string {
	var portNumber int

	cfEnv, err := cfenv.Current()
	if err != nil {
		portNumber, _ = strconv.Atoi(os.Getenv("PORT"))
		if portNumber == 0 {
			portNumber = 3000
		}
	} else {
		portNumber = cfEnv.Port
	}

	return strconv.Itoa(portNumber)
}

//GetAPICreds : Get api creds from OS env
func GetAPICreds() (string, string) {
	apiuser := os.Getenv("APIUSER")
	if apiuser == "" {
		apiuser = "api"
	}

	apipass := os.Getenv("APIPASS")
	if apipass == "" {
		apipass = "api"
	}

	return apiuser, apipass
}
