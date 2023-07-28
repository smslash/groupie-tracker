package pkg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	artistsURL  = "https://groupietrackers.herokuapp.com/api/artists"
	relationURL = "https://groupietrackers.herokuapp.com/api/relation"
	apiKey      = "AIzaSyDxzRzRqp3rv77CX4ZTK4pzw37uRQ-WVwk"
)

func Parse() {
	err := json.Unmarshal(ParseJSON(artistsURL), &Bands)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	var relation Relation
	err = json.Unmarshal(ParseJSON(relationURL), &relation)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	for i := 0; i < len(Bands); i++ {
		Bands[i].Relation = relation.Index[i]
	}
}

func ParseJSON(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return responseBody
}
