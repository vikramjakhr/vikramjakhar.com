package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"bitbucket.org/vikramjakhr/vikramjakhar.com/controllers"
	"bitbucket.org/vikramjakhr/vikramjakhar.com/config"
)

const (
	defaultPort string = "9000"
)

func main() {
	config.InitDB("mysql", "root:xxxxx@/vikramjakhar?charset=utf8")
	http.Handle("/", controllers.HttpFilter(http.HandlerFunc(controllers.Root)))
	port := port(os.Args)
	log.Println("Server started: http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

func port(args []string) string {
	port := defaultPort
	if len(args) > 1 && args[1] != "" {
		port = args[1]
		_, err := strconv.ParseUint(port, 10, 0)
		if err != nil {
			port = defaultPort
			log.Println("Incorrect port number. Falling back to default : ", defaultPort)
		}
	}
	return port
}
