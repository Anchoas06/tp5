package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "https://www.fruityvice.com/api/fruit/all"
        var rates []map[string]interface{}

	spaceClient := http.Client{
		Timeout: time.Second * 10, // Timeout after 10 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal([]byte(body), &rates)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	//printing decoded array values one by one
	for _, rate := range rates {
         fmt.Println("Code:",  rate["name"], "Rate:", rate["nutritions"])
	}
}