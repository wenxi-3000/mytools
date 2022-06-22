package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var inputFile = flag.String("f", "input.txt", "输入的文件")
var outputFile = flag.String("o", "output.txt", "输出的文件")

func main() {

	flag.Parse()

	//获取命令行参数
	inputFile := *inputFile
	outputFile := *outputFile

	//读入数据

	inputList := inputWork(inputFile)

	results := removeDuplicateElement(inputList)

	outputWork(outputFile, results)

}

func removeDuplicateElement(languages []string) []string {
	result := make([]string, 0, len(languages))
	temp := map[string]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

//输出
func outputWork(path string, results []string) {
	//创建文件句柄
	foutput, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer foutput.Close()
	w := bufio.NewWriter(foutput)
	defer w.Flush()

	for _, item := range results {
		//fmt.Println("ip:", ip)
		fmt.Println(item)
		w.WriteString(item + "\n")
	}
}

//从文件中读入数据，写入到workchan
func inputWork(path string) []string {
	var inputString []string
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputString = append(inputString, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return inputString

}
