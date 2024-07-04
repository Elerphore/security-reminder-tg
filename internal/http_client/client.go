package http_client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	client = &http.Client{
		Timeout: time.Second * 2,
	}
)

func GetJoke() Joke {
	resp, err := client.Get("https://v2.jokeapi.dev/joke/Any?blacklistFlags=nsfw,religious,political,racist,sexist,explicit&type=single")

	if err != nil {
		log.Default().Println(err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	joke := Joke{}

	jsonErr := json.Unmarshal(body, &joke)

	if jsonErr != nil {
		log.Fatal(readErr)
	}

	return joke
}
