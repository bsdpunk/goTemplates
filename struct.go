package main

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
