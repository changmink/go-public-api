package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://yts.mx/api/v2/list_movies.json?sort_by=like_count&order_by=desc&limit=5"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("http call error!")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("body parse Error")
	}

	fmt.Println(string(body))
}
