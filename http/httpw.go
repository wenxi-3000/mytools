package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	//文件处理
	file, err := os.Open("./ip.txt")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	ipPortList := strings.Split(string(content), "\n")

	//http机制
	timeout := time.Duration(10000 * 1000000)
	var tr = &http.Transport{
		MaxIdleConns:      30,
		IdleConnTimeout:   time.Second,
		DisableKeepAlives: true,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: time.Second,
		}).DialContext,
	}

	re := func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	client := &http.Client{
		Transport:     tr,
		CheckRedirect: re,
		Timeout:       timeout,
	}

	for _, item := range ipPortList {
		url := "http://" + item
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
		}
		defer resp.Body.Close()

		if resp != nil {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
			}
			if strings.Contains(string(body), "swagger-ui") {
				fmt.Println(item)
			}
		}
	}
}
