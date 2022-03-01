package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
		fmt.Fprint(rw, "hello world")
	})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalln("start err", err)
	}

}
