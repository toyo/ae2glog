package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/toyo/ae2glog"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		pl := ae2glog.NewLog(req)
		ae2glog.Infof(pl, "Test1!!")
		fmt.Fprintf(w, "Hello World")

	})

	port := os.Getenv("PORT")
	if port == "" {
		log.Print("NO PORT VARIABLE")
		port = "8080"
	}
	log.Printf("Listen on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
