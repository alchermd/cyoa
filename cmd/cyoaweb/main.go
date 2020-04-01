package main

import (
	"flag"
	"fmt"
	"github.com/alchermd/cyoa"
	"net/http"
	"os"
)

func main() {
	fileName := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	port := flag.Int("port", 8000, "the PORT to run the web server on")
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonToStory(f)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", cyoa.NewHandler(story))
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
