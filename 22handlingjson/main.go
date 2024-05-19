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

	// EncodeJson()
	// responseJSon()
	DecodeJson()
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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
        "coursename": "Web Development",
        "Price": 199,
        "website": "Udemy",
        "tags": ["HTML","CSS","JavaScript"]}
	`)

	var newCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &newCourse)
		fmt.Printf("%#v\n", newCourse)
	} else {
		fmt.Println("JSON was not Valid !!!!")
	}

	// some cases where you want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v\n", k, v)
	}

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
