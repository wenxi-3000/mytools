package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//文件处理
	file, err := os.Open("./ip.txt")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	ipPortList := strings.Split(string(content), "\n")

	//设置http请求的超时时间
	client := &http.Client{
		Timeout: time.Second * 3,
	}

	//结果
	var results []string

	//循环请求所有的ip:port
	for _, item := range ipPortList {
		url := "http://" + item + "/testx"
		fmt.Println(url)
		//resp, err := http.Get(url)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println(err)
		}

		req.Header.Add("Connection", "close")
		req.Close = true

		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			continue
		}
		defer resp.Body.Close()

		if resp != nil {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
			}
			if strings.Contains(string(body), "swagger-ui") {
				results = append(results, item)
			}
		}
	}

	resultFile, err := os.OpenFile("./output.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(nil)
		os.Exit(1)
	}
	defer resultFile.Close()

	writer := bufio.NewWriter(resultFile)
	for _, item := range results {
		writer.WriteString(item + "\n")
	}
	writer.Flush()

}
