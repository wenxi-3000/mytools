package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./ip-port.txt")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	ipPortList := strings.Split(string(content), "\n")

	for _, item := range ipPortList {
		httpx(item)
	}

}

func httpx(host string) {
	url := "http://" + host
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(body), "swagger-ui") {
		fmt.Println(host)
	}
}

// func GetFileContent(filename string) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }
