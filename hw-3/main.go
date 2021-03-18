package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserId   int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {

	data := []byte(`
		{
			"id": 1,
			"name": "John",
			"password": "123"
		}
	`)

	user := User{}

	err := json.Unmarshal(data, &user)

	if err != nil {
		fmt.Printf("Eror encoding JSON")
	}

	fmt.Printf("%v", user)
}
