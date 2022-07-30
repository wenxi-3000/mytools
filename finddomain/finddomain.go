package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
)

var (
	target string

	concurrency int
	results     map[string]struct{}

	rwmutex sync.RWMutex
)

func main() {
	flag.StringVar(&target, "t", "", "Specify target to clean")
	flag.IntVar(&concurrency, "c", 20, "Set the concurrency level")
	flag.Parse()
	results = make(map[string]struct{})
	var wg sync.WaitGroup
	jobs := make(chan string, concurrency)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				worker(job)
			}
		}()
	}

	sc := bufio.NewScanner(os.Stdin)
	go func() {
		for sc.Scan() {
			url := strings.TrimSpace(sc.Text())
			if err := sc.Err(); err == nil && url != "" {
				jobs <- url
			}
		}
		close(jobs)
	}()
	wg.Wait()

	for item := range results {
		fmt.Println(item)
	}
}

func worker(ur string) {
	u, err := url.Parse(ur)
	if err != nil {
		log.Fatal(err)
	}
	domain := u.Hostname()
	rwmutex.Lock()
	results[domain] = struct{}{}
	rwmutex.Unlock()
}

//匹配域名
func getHost(source string, domain string) []string {
	//strings.Replace(domain,'.', '`\.`')
	reg := `(?:[a-z0-9](?:[a-z0-9\-]{0,61}[a-z0-9])?\.){0,}` + domain
	//results_domains = re.findall(regexp, str(source), re.I)
	var linkFinderRegex = regexp.MustCompile(reg)
	matchs := linkFinderRegex.FindAllString(source, -1)
	fmt.Println(matchs)
	return matchs
}
