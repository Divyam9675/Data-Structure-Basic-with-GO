package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	
	Second := map[string]interface{}{
		"Degree": map[string]string{
			"Bachlors": "B.C.A.",
			"Masters":  "M.C.A.",
		},
		"Name": "Divyam",
		"Status": "Fear of Future",
	}

	data, _ := json.Marshal(Second)
	fmt.Println(string(data))
}