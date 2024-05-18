package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web Verb")

	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func PerformGetRequest() {
	const myurl = "https://www.york.ac.uk/teaching/cws/wws/webpage1.html"

	response, err := http.Get(myurl)

	checkNilErr(err)
	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length is ", response.ContentLength)

	content, err := io.ReadAll(response.Body)
	// fmt.Println(string(content))
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}

func PerformPostRequest() {
	const url = "https://redis-4eqp.onrender.com/refresh-token"

	payload := `"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJtb2hkYXF1aWIzMUBnbWFpbC5jb20iLCJleHAiOjE3MTYwMzM1NzF9.3-1ZKkvVPz_1DtGMxG2gjHbrNMfGDUnEBtLdvJ-mueI"`

	payloadReader := strings.NewReader(payload)

	resp, err := http.Post(url, "application/json", payloadReader)
	checkNilErr(err)
	defer resp.Body.Close()
	fmt.Println("Status Code:", resp.StatusCode)

	/** method 1
	content, err := io.ReadAll(resp.Body)
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())
	**/

	// Method 2
	var responseBuffer bytes.Buffer
	_, err = responseBuffer.ReadFrom(resp.Body)
	checkNilErr(err)
	fmt.Println("Response Body:", responseBuffer.String())

}

func PerformPostFormRequest() {
	const myurl = "https://llm-pgc4.onrender.com/blog_generator"

	//form data
	data := url.Values{}
	data.Add("topic", "Go Lang vs Python in Backend Development")

	resp, err := http.PostForm(myurl, data)
	checkNilErr(err)

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)

	fmt.Println(string(content))

}
