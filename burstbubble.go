package main

import "fmt"
import "net/http"
import "io/ioutil"

func main() {
	fmt.Printf("Time to kick ass and bust bubbles ... and I'm all out of ass\n")
	resp, err := http.Get("http://google.com/")
	if err != nil {
		fmt.Printf("Error occurred")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	
	fmt.Printf(string(body))
}
