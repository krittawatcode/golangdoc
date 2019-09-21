package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	resp, err := http.Get("http://jsonplaceholder.typicode.com/todos")
	if err != nil {
		return
	}
	todoDecoder := json.NewDecoder(resp.Body)
	todos := []Todo{}
	todoDecoder.Decode(&todos)
	resp.Body.Close()
	// todoEncoder := json.NewEncoder(os.Stdout)
	// todoEncoder.Encode(todos)

	indexTemplateBytes, err := ioutil.ReadFile("D:/Krittawat/Learn/GoWorkspace/src/htmlTemplate/html_template.go")
	if err != nil {
		return
	}

	indexTemplate := template.Must(template.New("index").Parse(string(indexTemplateBytes)))
	indexTemplate.Execute(os.Stdout, todos)

}
