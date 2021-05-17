package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetHtml(url string) []byte {
	fmt.Println(url)
	responce, errors := http.Get(url)
	if errors != nil {
		panic(errors)
	}
	defer responce.Body.Close()
	content, errors := ioutil.ReadAll(responce.Body)
	if errors != nil {
		panic(errors)
	}
	//fmt.Printf("%s\n", content)
	return content
}
