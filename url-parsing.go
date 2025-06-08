package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlSrting := "http://example.com:80/hello?name=rizky&age=24"
	var url, err = url.Parse(urlSrting)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	name := url.Query()["name"][0]
	age := url.Query()["age"][0]

	fmt.Printf("url: %s\n", urlSrting)

	fmt.Printf("protocol: %s\n", url.Scheme) // http
	fmt.Printf("host: %s\n", url.Host)       //example.com:80
	fmt.Printf("path: %s\n", url.Path)       //hello
	fmt.Printf("name: %s, age: %s\n", name, age)
}
