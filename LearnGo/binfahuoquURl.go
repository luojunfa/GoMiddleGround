package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type results struct {
	url    string
	result string
}//创建results结构体

func main() {

	start := time.Now()
	ch := make(chan string)//??
	result := make(map[string]string)
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		go fetch(url, ch)
		go fetch(url, ch)
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		for i := 0; i < 3; i++ {
			fmt.Println(<-ch)
		}
	}

	for k, v := range result {
		fmt.Printf("%s : %s\n", k, v)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
