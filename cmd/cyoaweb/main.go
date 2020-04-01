package main

import (
	"flag"
	"fmt"
	"github.com/alchermd/cyoa"
	"os"
)

func main() {
	fileName := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonToStory(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
