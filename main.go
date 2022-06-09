package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	Id   int    `json:"id"`
	Name string `json:name`
}

var listOfTodos []Todo

func init() {
	listOfTodosJSON := `
		[
			{
				"id": 1,
				"name": "Buy Groceries"
			},
			{
				"id": 2,
				"name": "Go to the piano class"
			}
		]
	`

	err := json.Unmarshal([]byte(listOfTodosJSON), &listOfTodos)

	if err != nil {
		log.Fatal(err)
	}
}

func todosHandler(w http.ResponseWriter, req *http.Request) {
	listOfTodosJSON, err := json.Marshal(listOfTodos)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(listOfTodosJSON)
}

func main() {
	fmt.Println("Server is running...")

	http.HandleFunc("/todos", todosHandler)

	err := http.ListenAndServe("localhost:4000", nil)

	if err != nil {
		fmt.Println("Something wrong happened")
		log.Fatal(err)
	}
}
