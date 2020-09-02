package main

import (
	"encoding/json"
	"fmt"
	"music-cli/src/lib/models/song"
	"net/http"
	"net/url"
	"flag"
)

func main() {

	searchTerm := flag.String("search", "", "Song to search for")

	flag.Parse();

	fmt.Printf("---\nSearching for: %s\n---\n", *searchTerm);

	searchTermQuery := url.QueryEscape(*searchTerm);

	url:= "https://orion-server.herokuapp.com/api/search?searchTerm=\""+searchTermQuery+"\"";
	

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
