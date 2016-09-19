package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	jsonData := `{"name": "Luiz Alexandre", "age": "22"}`

	var obj map[string]string
	err := json.Unmarshal([]byte(jsonData), &obj)

	if err != nil {
		panic(err)
	}
	fmt.Println(obj)
	fmt.Println(jsonData)
}
