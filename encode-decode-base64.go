package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var URL = "https://exampple.com/"
	var encodedStr = base64.URLEncoding.EncodeToString([]byte(URL))
	fmt.Println(encodedStr)

	var decodedByte, _ = base64.URLEncoding.DecodeString(encodedStr)
	var decodedStr = string(decodedByte)
	fmt.Println(decodedStr)
}
