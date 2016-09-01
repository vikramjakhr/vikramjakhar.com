package controllers

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"strings"
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

type commentResponse struct {
	Comments     []comment `json:"comments"`
	CommentCount int `json:"commentCount"`
}

func fetchComments(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		postName := strings.TrimSpace(string(body))

		fileName := "data/comments/" + postName + ".json"
		_, err = os.Stat(fileName)
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to find comments data"), http.StatusInternalServerError)
			return
		}

		commentData, err := ioutil.ReadFile(fileName)
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to find comments data"), http.StatusInternalServerError)
			return
		}

		var comments []comment

		if err := json.Unmarshal(commentData, &comments); err != nil {
			http.Error(w, fmt.Sprintf("Unable to find comments data"), http.StatusInternalServerError)
			return
		}

		commentRes := commentResponse{Comments:comments, CommentCount:len(comments)}
		res, err := json.Marshal(commentRes)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(res)
	default:
		http.Error(w, fmt.Sprintf("Unsupported Method : %s", req.Method), http.StatusMethodNotAllowed)
	}
}

type CommentReqData struct {
	Author   string `json:"author"`
	ID       int64 `json:"id"`
	Text     string `json:"text"`
	PostName string `json:"postName"`
}

func postComment(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		decoder := json.NewDecoder(req.Body)
		var cmntData CommentReqData
		err := decoder.Decode(&cmntData)
		if err != nil {
			panic(err)
		}

		fileName := "data/comments/" + strings.TrimSpace(cmntData.PostName) + ".json"
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
		var comments []comment

		if err := json.Unmarshal(commentData, &comments); err != nil {
			http.Error(w, fmt.Sprintf("Unable to find comments data"), http.StatusInternalServerError)
			return
		}

		comments = append(comments, comment{ID:cmntData.ID, Author:cmntData.Author, Text:cmntData.Text})

		commentData, err = json.MarshalIndent(comments, "", "    ")
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to marshal comments to json: %s", err), http.StatusInternalServerError)
			return
		}

		err = ioutil.WriteFile(fileName, commentData, fi.Mode())
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to write comments"), http.StatusInternalServerError)
			return
		}

		commentRes := commentResponse{Comments:comments, CommentCount:len(comments)}
		res, err := json.Marshal(commentRes)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(res)
	default:
		http.Error(w, fmt.Sprintf("Unsupported Method : %s", req.Method), http.StatusMethodNotAllowed)
	}
}