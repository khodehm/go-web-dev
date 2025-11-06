package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	ID     string
	Name   string
	Age    string
	Gender string
}

type Users map[string]User

// StoreUser stores users to  data.json file
func StoreUser(m map[string]User) {
	f, err := os.Create("data.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(&m)
}

// LoadUsers loads users form data.json file
func LoadUsers() map[string]User {
	u := make(map[string]User)
	f, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return u
	}
	err = json.NewDecoder(f).Decode(&u)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	return u

}
