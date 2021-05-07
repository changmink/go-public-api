package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Data          Data   `json:"data"`
}

type Data struct {
	MovieCount int     `json:"movie_count"`
	Limit      int     `json:"limit"`
	PageNumber int     `json:"page_number"`
	Movies     []Movie `json:"movies"`
}

type Movie struct {
	Id    int    `json:"id"`
	Url   string `json:"url"`
	Title string `json:"title"`
	Image string `json:"small_cover_image"`
}

func main() {
	url := "https://yts.mx/api/v2/list_movies.json?sort_by=like_count&order_by=desc&limit=5"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("http call error!")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Get Response Error")
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal("body parse Error")
	}
	fmt.Println(response.Data)
}
