package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"syscall"
)

func DownloadImages(links []string) {

	//fanOut() {
		//Imagine you are the manager again but this time you hire a team of employees.
		//You have an individual task you want each employee to perform.
		//As each individual employee finishes their task,
		//they need to provide you with a paper report that must be placed in your BOX on your desk. }

		emps := len(links)
		ch := make(chan string, emps)

		for _, el := range links{
			go func(link string) {
				fileName := "/home/abonta/goProjects/src/awesomeMultiThreading/images/" + path.Base(link)
				//fmt.Println(i, fileName)
				err := downloadFile(link, fileName)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Print(syscall.Gettid(), " : ")
				ch <- fmt.Sprintf("File %s download \n", path.Base(link))
			} (el)
		}

		for emps > 0 {
			p := <-ch
			fmt.Println(p)
			emps--
	}
}

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}




