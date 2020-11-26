//1.Encoding parsing
package main
import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	queryStr := "name=Divyam%20&Status=%2Bemployee&phone=%2B9525252525"
	params, err := url.ParseQuery(queryStr)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Query Params: ")
	for key, value := range params {
		fmt.Printf("  %v = %v\n", key, value)
	}
}