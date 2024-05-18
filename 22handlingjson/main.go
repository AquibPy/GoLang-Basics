package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Handling JSON")

	EncodeJson()
	// responseJSon()
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func EncodeJson() {
	fmt.Println("Convert Data into Valid JSON")
	courses := []course{
		{
			"Data Science",
			299,
			"deeplearning.ai",
			"root123",
			[]string{"python", "ML"},
		},
		{
			"Web Development",
			199,
			"Udemy",
			"webdev456",
			[]string{"HTML", "CSS", "JavaScript"},
		},
		{
			"Machine Learning",
			349,
			"Coursera",
			"ml789",
			nil,
		},
	}

	// package this data as json data
	finalJson, err := json.Marshal(courses)
	checkNilErr(err)
	fmt.Printf("%s\n", finalJson)

	finalJsonNew, err := json.MarshalIndent(courses, "", "\t")
	checkNilErr(err)
	fmt.Printf("%s\n", finalJsonNew)

}

func responseJSon() {
	const myurl = "https://llm-pgc4.onrender.com/blog_generator"

	// Form data
	data := url.Values{}
	data.Add("topic", "Ai as future")

	resp, err := http.PostForm(myurl, data)
	checkNilErr(err)
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	checkNilErr(err)

	// Parse the JSON response
	var jsonResponse map[string]interface{}
	err = json.Unmarshal(content, &jsonResponse)
	checkNilErr(err)

	// Format JSON for printing
	jsonFormatted, err := json.MarshalIndent(jsonResponse, "", "  ")
	checkNilErr(err)

	fmt.Println(string(jsonFormatted))
}
