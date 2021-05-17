package main

import (
	cmd "awesomeMultiThreading/cmd"
	"fmt"
)

func main() {
	fmt.Println("===================================================")
	URL := "http://me.utm.md/"

	fmt.Println("====================Get Request====================")
	html := cmd.GetHtml(URL)
	//fmt.Printf("%s\n", html)
	content := string(html)
	//content = strings.ReplaceAll(content, "\n", "")
	imageLInks := cmd.FindImages(content)
	//fmt.Println(imageLInks)
	cmd.DownloadImages(imageLInks)
}
