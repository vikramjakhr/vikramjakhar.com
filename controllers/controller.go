package controllers

import (
	"net/http"
	"html/template"
	"strings"
	"github.com/choudhary92/vikramjakhar.com/data"
	"fmt"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"os"
)

func init() {
	registerHandlers()
}

var pagination int64 = 5

func Root(w http.ResponseWriter, r *http.Request) {
	Pagination(w, r)
}

type BlogMetadata struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	Title            string `json:"title"`
	TmplFileName     string `json:"tmplFileName"`
	Author           string `json:"author"`
	PublishDate      int64 `json:"publishDate"`
	Tags             string `json:"tags"`
	CommentCount     int64 `json:"commentCount"`
	ImageUrl         string `json:"imageUrl"`
	CommentsFileName string `json:"commentsFileName"`
	Summary          string `json:"summary"`
}

func Pagination(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s := strings.Split(r.RequestURI, "/")
		page, err := strconv.ParseInt(s[len(s) - 1], 10, 0)
		if err != nil {
			page = 1
		}
		var fileName string
		var t, max int64 = 1, data.BlogMetaInfo.TotalBlogs
		if page > 0 && (page * pagination) <= data.BlogMetaInfo.TotalBlogs {
			t = (page * pagination) / 20
			fileName = "blog-" + strconv.FormatInt(t * 20 + 1, 10) + "-to-" + strconv.FormatInt(t * 20 + 20, 10) + ".json"
			max = pagination
		} else {
			page = 1
			t = 1
			max = data.BlogMetaInfo.TotalBlogs
			fileName = "blog-1-to-20.json"
		}

		_, err = os.Stat("data/blogs/" + fileName)
		if err != nil {
			NotFound(w, r)
			return
		}
		blogsData, err := ioutil.ReadFile("data/blogs/" + fileName)
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to find data"), http.StatusInternalServerError)
			return
		}
		var blogsMetadata []BlogMetadata

		if err := json.Unmarshal(blogsData, &blogsMetadata); err != nil {
			http.Error(w, fmt.Sprintf("Unable to find blogs data"), http.StatusInternalServerError)
			panic(err)
			return
		}

		a := (page * pagination) - (t * 20)
		if a < 1 {
			a = 1
		}
		templateInfo := data.TemplateInfo{Title:"Vikram Jakhar", Name:"summary-tmpl", Path:"views/posts/summary-tmpl.html"}
		renderMainLayout(w, r, templateInfo, templateInfo.Title, blogsMetadata[a - 1:a - 1 + max])
	default:
		http.Error(w, "Method not Allowd", http.StatusMethodNotAllowed)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	templateInfo := data.TemplateInfo{Title:"About | Vikram Jakhar", Name:"about-tmpl", Path:"views/about/about-tmpl.html"}
	renderLayoutWithoutSidebar(w, r, templateInfo, templateInfo.Title, nil)
}

func Golang(w http.ResponseWriter, r *http.Request) {
	Pagination(w, r)
}

func Post(w http.ResponseWriter, r *http.Request) {
	s := strings.Split(r.RequestURI, "/")
	tmplName := strings.TrimSpace(s[2])
	m := strings.TrimSpace(data.BlogMetaInfo.Meta[tmplName])
	if strings.EqualFold(m, "") {
		NotFound(w, r)
		return
	}
	blogInfo := strings.Split(m, ",")
	pos, err := strconv.ParseInt(blogInfo[0], 10, 0)
	blogMeta, err := GetPostInfo(tmplName, blogInfo[1], pos)
	if err != nil {
		NotFound(w, r)
		return
	}
	post, err := execTmplToString(blogMeta.Name, "views/posts/" + blogMeta.TmplFileName, blogMeta)
	if err != nil {
		NotFound(w, r)
		return
	}
	tmplInfo := data.TemplateInfo{Title:blogMeta.Title, Name:"post-tmpl", Path:"views/posts/post-tmpl.html"}
	renderMainLayout(w, r, tmplInfo, blogMeta.Title, template.HTML(post))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("views/err/404.html")
	tmpl.Execute(w, nil)
}

func InternalError(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("views/err/500.html")
	tmpl.Execute(w, nil)
}

func GetPostInfo(pName, fileName string, pos int64) (BlogMetadata, error) {
	blogMeta := BlogMetadata{}
	_, err := os.Stat("data/blogs/" + fileName)
	if err != nil {
		return blogMeta, err
	}
	blogsData, err := ioutil.ReadFile("data/blogs/" + fileName)
	if err != nil {
		return blogMeta, err
	}
	var blogsMetadata []BlogMetadata
	if err := json.Unmarshal(blogsData, &blogsMetadata); err != nil {
		return blogMeta, err
	}
	//TODO : pointing to same slice, use copy
	blogMeta = blogsMetadata[pos]
	return blogMeta, nil
}