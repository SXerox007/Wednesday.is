package main

import (
	"encoding/json"
	"log"
	"net/http"

	"Wednesday.is/base/router"
	"Wednesday.is/base/router/server"

	env "Wednesday.is/base/environment"
)

// common response
type CommonResponse struct {
	Success bool        `param:"success" json:"success"`
	Message string      `param:"message" json:"message"`
	Code    int         `param:"code" json:"code"`
	Data    interface{} `param:"data" json:"data"`
}

// init
func Init() {
	environment := env.GetEnv()
	port := env.GetPort()
	setupRouter(environment, port)
}

func main() {
	Init()
}

func setupRouter(env, port string) {
	//initilize the router
	router.InitRouter()
	wed := router.SubRouter("/wednesday/{version}/")
	wed.HandleFunc("/roman", RomanToInterger()).Methods("POST")

	log.Println("Server serve at: ", "localhost:"+port)
	server.StartServer(port)
}

// respondWithJSON - Serializes the payload to JSON and writes to ResponseWriter.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("X-XSS-Protection", "1; mode=block")
	w.WriteHeader(code)
	w.Write(response)
}
