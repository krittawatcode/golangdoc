package main

import "net/http"

type Photos []struct {
	AlbumID	int	`json:"albumId"`
	ID	int	`json:"id"`
	Title	string	`json:"title"`
	URL	string	`json:"url"`
	ThumbnailURL	string	`json:"thumbnailUrl"`
  }
  

func main() {
	?? := getJson("https://jsonplaceholder.typicode.com/photos")
}

func getJson(url string) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
}
