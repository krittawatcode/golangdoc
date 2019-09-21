package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed"`
}

func main() {

	resp, err := http.Get("http://jsonplaceholder.typicode.com/todos")
	fmt.Println(err)
	fmt.Println(resp)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}

	dataStruct := []Todo{}
	v := &dataStruct
	json.Unmarshal(body, v)
	// json.Unmarshal([]byte(body), v)
	// dataStruct[0].Completed = true
	// dataStruct[1].Completed = true

	// dataStruct --> std.Output
	// func Marshal(v interface{}) ([]byte, error) {...}

	/* ---- Part I ---- */
	/*
		result, err := json.Marshal(dataStruct)
		if err != nil {
			return
		}
		fmt.Println(string(result))
	*/

	/* ---- Part II : MarshalIndent ---- */

	// func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {...}

	result, err := json.MarshalIndent(dataStruct, "", "    ") // for make readable json format
	if err != nil {
		return
	}

	fmt.Println(string(result))
}
