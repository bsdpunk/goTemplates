package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type jsonStruct struct {
	Category  Category `json:"category"`
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	PhotoUrls []string `json:"photoUrls"`
	Status    string   `json:"status"`
	Tags      Category `json:"tags"`
}
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Tags []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	jsonStr := []byte(`{  "id": 0,  "category": {    "id": 0,    "name": "string"  },  "name": "doggie",  "photoUrls": [    "string"  ],  "tags": [    {      "id": 0,      "name": "string"    }  ],  "status": "available"}`)
	//jsonStruct{
	//	Category: Category{
	//		ID:   0,
	//		Name: "doggie"},
	//	ID:        0,
	//	Name:      "string",
	//	PhotoUrls: make([]string, 5, 6),
	//	Status:    "available",
	//	Tags: Category{
	//		ID:   0,
	//		Name: "doggie"},
	//}

	url := "http://petstore.swagger.io/v2/pet"
	fmt.Println("URL:>", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
