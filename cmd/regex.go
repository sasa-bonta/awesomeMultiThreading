package cmd

import (
	"fmt"
	"regexp"
	"strings"
)

func FindImages(html string) []string {
	//fmt.Printf(html)

	re, _ := regexp.Compile("[^\"']*\\.(?:png|jpg|gif)")
	imgs := re.FindAllString(html, -1)

	//fmt.Println(imgs)
	imgsLen := len(imgs)
	fmt.Println(imgsLen)

	var links = make([]string, imgsLen)

	for i, el := range imgs {
		if strings.HasPrefix(el, "http://") {
			links[i] = el
		} else {
			el = "http://me.utm.md/" + el
			links[i] = el
		}
	}

	//for _, el := range links{
	//	fmt.Println(el)
	//}

	return links
}
