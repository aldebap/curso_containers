////////////////////////////////////////////////////////////////////////////////
//	main.go  -  Dec/3/2023  -  aldebap
//
//	Rest API example
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type statusResponse struct {
	Status string `json:"status"`
}

// get status handler
func GetStatus(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	fill response payload
	var responseData = statusResponse{}

	responseData.Status = "Server up and running"

	httpResponse.Header().Add("Content-Type", "application/json")
	httpResponse.WriteHeader(http.StatusOK)
	json.NewEncoder(httpResponse).Encode(responseData)
}

func main() {

	//	splash screen
	fmt.Printf(">>> RestAPI\n\n")

	//	get configuration from environment
	var servicePort string

	servicePort = os.Getenv("SERVICEPORT")

	if len(servicePort) == 0 {
		servicePort = ":8080"
	} else if servicePort[0] != ':' {
		servicePort = ":" + servicePort
	}

	//	start the Web Server
	restRouter := mux.NewRouter()

	restRouter.HandleFunc("/status", GetStatus).Methods(http.MethodGet)

	http.Handle("/", restRouter)

	//start and listen to requests
	fmt.Printf("Listening port %s\n", servicePort)

	log.Panic(http.ListenAndServe(servicePort, restRouter))
}
