package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func Json() {
	// Parse

	doc := `
	{
		"name": "maria",
		"age": 10
	}`

	var data map[string]interface{}
	json.Unmarshal([]byte(doc), &data)

	fmt.Println(data["name"], data["age"])

	// Create JSON

	data = make(map[string]interface{})

	data["name"] = "karu"
	data["age"] = 10
	data["school"] = "CAU"

	// d, _ := json.Marshal(data)
	d, _ := json.MarshalIndent(data, "", "	")
	fmt.Println(string(d))
}

type Author struct {
	Name  string `json:"name`
	Email string `json:"email` // 이렇게 하면 소문자로 저장
}

type Comment struct {
	Id      uint64
	Author  Author
	Content string
}

type Article struct {
	Id         uint64
	Title      string
	Author     Author
	Content    string
	Recommends []string
	Comments   []Comment
}

func Json2() {
	doc := `
	[{
		"Id": 1,
		"Title": "Hello, world!",
		"Author": {
			"Name": "Maria",
			"Email": "maria@rolling.ress"
		},
		"Content": "Hello!",
		"Recommends": [
			"John",
			"Andrew"
		],
		"Comments": [{
			"id": 1,
			"Author": {
				"Name": "Andrew",
				"Email": "andrew@rolling.ress"
			},
			"Content": "Hello Maria"
		}]
	}]`

	var data []Article
	json.Unmarshal([]byte(doc), &data)
	fmt.Println(data)

	// TO JSON

	// 구조체 슬라이스에 값 할당 후
	d, _ := json.Marshal(data)

	os.WriteFile("test.json", d, os.FileMode(0644))
}

func main() {
	Json()
	Json2()
}