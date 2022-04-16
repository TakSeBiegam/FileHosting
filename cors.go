package main

import (
	"encoding/json"
	"net/http"
)

type Article struct {
	Id      string
	Title   string
	Desc    string
	Content string
}

var Articles []Article

func main() {
	Articles = []Article{
		{Id: "1", Title: "First article", Desc: "Title of this fine article", Content: "Content for this fine article"},
		{Id: "2", Title: "Second article", Desc: "Title of this majestic article", Content: "Content for this majestic article"},
	}
	http.HandleFunc("/articles", foo)
	http.ListenAndServe(":8002", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	js, err := json.Marshal(Articles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
