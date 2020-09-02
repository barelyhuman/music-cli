package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/barelyhuman/music-cli/lib/models/song"
)

func main() {

	searchTerm := flag.String("s", "", "Song to search for (Required)")

	flag.Parse()

	if len(*searchTerm) <= 0 {
		flag.Usage()
		os.Exit(2)
	}

	fmt.Printf("---\nSearching for: %s\n---\n", *searchTerm)

	searchTermQuery := url.QueryEscape(*searchTerm)

	url := "https://orion-server.herokuapp.com/api/search?searchTerm=\"" + searchTermQuery + "\""

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	var songs []song.Song
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&songs)
	for _, song := range songs {
		fmt.Printf("%v\n", song.Title)
	}
}
