package controllers

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"io"
	"bytes"
	"time"
)

func init() {
	http.HandleFunc("/api/comments/fetch", fetchComments)
	http.HandleFunc("/api/comments/post", postComment)
}

type comment struct {
	ID     int64  `json:"id"`
	Author string `json:"author"`
	Text   string `json:"text"`
}

func postComment(w http.ResponseWriter, req *http.Request) {

	fileName := "data/comments/comments-introduction-to-go.json"
	fi, err := os.Stat(fileName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to find comments data"), http.StatusInternalServerError)
		return
	}

	commentData, err := ioutil.ReadFile(fileName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to find comments data"), http.StatusInternalServerError)
		return
	}

	switch req.Method {
	case "POST":
		var comments []comment

		if err := json.Unmarshal(commentData, &comments); err != nil {
			http.Error(w, fmt.Sprintf("Unable to find comments data"), http.StatusInternalServerError)
			return
		}

		comments = append(comments, comment{ID:time.Now().Unix() / 10000, Author:"guest", Text:"Hellooo : " + time.Now().String()})

		commentData, err = json.MarshalIndent(comments, "", "    ")
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to marshal comments to json: %s", err), http.StatusInternalServerError)
			return
		}


		err := ioutil.WriteFile(fileName, commentData, fi.Mode())
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to write comments"), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		io.Copy(w, bytes.NewReader(commentData))
	default:
		http.Error(w, fmt.Sprintf("Unsupported Method : %s", req.Method), http.StatusMethodNotAllowed)
	}
}

func fetchComments(w http.ResponseWriter, req *http.Request) {
	fileName := "data/comments/comments-introduction-to-go.json"
	_, err := os.Stat(fileName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to find comments data"), http.StatusInternalServerError)
		return
	}

	commentData, err := ioutil.ReadFile(fileName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to find comments data"), http.StatusInternalServerError)
		return
	}

	switch req.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		io.Copy(w, bytes.NewReader(commentData))
	default:
		http.Error(w, fmt.Sprintf("Unsupported Method : %s", req.Method), http.StatusMethodNotAllowed)
	}
}