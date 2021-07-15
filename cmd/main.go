package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 3000, "")

	flag.Parse()

	fs := http.FileServer(http.Dir("/static"))
	http.Handle("/api/fs", fs)

	log.Println("listening on", *port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), fs)
	if err != nil {
		log.Fatal(err)
	}
}