package main

import (
	"encoding/json"
	"fmt"
	"rest.http/v/domain"
)

func main() {

	user := domain.User{
		Name:      "mihai",
		Mail:      "mihai@gmail.com",
		Age:       22,
		Interests: []string{"manele", "bani", "femei"},
	}

	fmt.Printf("User %v\n", user)
	res, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("Failed to marshal, err=%v\n", err)
	}

	fmt.Printf("%s\n", string(res))

	data := `{
		"name": "mihai",
		"mail": "mihai@gmail.com",
		"age": 22,
		"interests": [
			"manele", "bani", "femei"
		]
	}`

	var v map[string]interface{}
	err = json.Unmarshal([]byte(data), &v)

	if err != nil {
		fmt.Printf("Failed to unmarshal, err=%v\n", err)
		return
	}

	fmt.Printf("%v\n", v)

	// sau

	var vv domain.User
	err2 := json.Unmarshal([]byte(data), &vv)
	if err2 != nil {
		fmt.Printf("Failed to unmarshal2, err=%v\n", err2)
		return
	}

	fmt.Printf("%v\n", vv.Name)

}
