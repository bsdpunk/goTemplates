package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://api.datamuse.com/words?rel_rhy=hammer"
	fmt.Println("URL:>", url)

	req, err := http.Get(url) //bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	fmt.Println("reqonse Status:", req.Status)
	fmt.Println("reqonse Headers:", req.Header)
	fmt.Println("reqonse Body:", string(body))
}
