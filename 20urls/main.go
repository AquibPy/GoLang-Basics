package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.york.ac.uk/teaching/cws/wws/webpage1.html?course=datascience&paymentid=4545sds"

func main() {
	fmt.Println("Handling Urls in Go Lang")
	fmt.Println(myurl)
	result, _ := url.Parse((myurl))

	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Port: ", result.Port())
	fmt.Println("RawQuery: ", result.RawQuery)

	qparams := result.Query()

	fmt.Println("qparams:", qparams)
	fmt.Println(qparams["course"])

	for _, val := range qparams {
		fmt.Println("Params is ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=aquib",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}

/**
output will be

Handling Urls in Go Lang
https://www.york.ac.uk/teaching/cws/wws/webpage1.html?course=datascience&paymentid=4545sds
Scheme:  https
Host:  www.york.ac.uk
Path:  /teaching/cws/wws/webpage1.html
Port:
RawQuery:  course=datascience&paymentid=4545sds
qparams: map[course:[datascience] paymentid:[4545sds]]
[datascience]
Params is  [datascience]
Params is  [4545sds]
https://lco.dev/tutcss

**/
