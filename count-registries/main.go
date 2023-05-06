package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Object struct {
	Id        string  `json:"ID"`
	Date      int64   `json:"Date"`
	Region    string  `json:"Region"`
	City      string  `json:"City"`
	Category  string  `json:"Category"`
	Product   string  `json:"Product"`
	Qty       int64   `json:"Qty"`
	UnitPrice float64 `json:"UnitPrice"`
}

func main() {

	file, err := os.Open("/app/data/OUT/output.json")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	var objects []Object

	err = json.Unmarshal(bytes, &objects)

	if err != nil {
		panic(err)
	}

	fmt.Println("len(objects)", len(objects))

}
