package main

import (
	"encoding/json"
	"io/ioutil"
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
	http.HandleFunc("/movies", func(w http.ResponseWriter, req *http.Request) {
		url := "https://yts.mx/api/v2/list_movies.json?sort_by=like_count&order_by=desc&limit=5"
		resp, err := http.Get(url)
		if err != nil {
			msg := "http call error!"
			w.Write([]byte(msg))
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			msg := "Get Response Error"
			w.Write([]byte(msg))
			return
		}

		var response Response
		if err := json.Unmarshal(body, &response); err != nil {
			msg := "body parse Error"
			w.Write([]byte(msg))
			return
		}

		data, err := json.Marshal(response.Data)
		if err != nil {
			msg := "convert response Error"
			w.Write([]byte(msg))
			return
		}

		w.Write(data)
	})

	http.ListenAndServe(":5000", nil)
}
