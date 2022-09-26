package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("./ip.txt")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	ipPortList := strings.Split(string(content), "\n")

	for _, item := range ipPortList {
		url := "http://" + item
		//resp, err := http.Get(url)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		cli := http.Client{
			Timeout: time.Millisecond * 3,
		}
		resp, err := cli.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if strings.Contains(string(body), "swagger-ui") {
			fmt.Println(item)
		}
	}

}
