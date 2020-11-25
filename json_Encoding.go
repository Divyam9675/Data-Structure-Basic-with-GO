package main

import (
	"encoding/json"
	"fmt"
)

type First struct {
	Name    string `json:"name"`
	Status string `json:"status"`
}

func main() {
	top := &First{
		Name:     "Divyam",
		Status: "Tense about Future",
	}

	data, _ := json.Marshal(top)

	fmt.Println(string(data))
}