package controllers

import (
	"net/http"
	"html/template"
	"bytes"
	"github.com/choudhary92/vikramjakhar.com/data"
	"log"
	"os"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

const (
	mainLayout string = "main-layout"
	layoutWithoutSidebar string = "layout-without-sidebar"
	postsPath string = "views/blogs/"
	aboutPath string = "views/about/"
)

var parsedMainLayout *template.Template
var parsedLayoutWithoutSidebar *template.Template

type TmplContext struct {
	Content template.HTML
	Title   string
}

type Data struct {
	Data interface{}
	Next int64
	Prev int64
}

func registerHandlers() {
	http.Handle("/page/", HttpFilter(http.HandlerFunc(Pagination)))
	http.Handle("/about/", HttpFilter(http.HandlerFunc(About)))
	http.Handle("/golang/", HttpFilter(http.HandlerFunc(Golang)))
	http.Handle("/post", HttpFilter(http.HandlerFunc(Root)))
	http.Handle("/post/", HttpFilter(http.HandlerFunc(Post)))
	http.Handle("/404/", HttpFilter(http.HandlerFunc(NotFound)))
	http.Handle("/assets/", http.FileServer(http.Dir("views")))
}

func parseMainLayout() {
	var err error
	parsedMainLayout, err = template.ParseFiles("views/layout/main-layout.html", "views/templates/login-signup-tmpl.html",
		"views/templates/footer.html", "views/templates/header.html", "views/templates/reside-bar-tmpl.html",
		"views/templates/reside-bar-search-tmpl.html", "views/templates/reside-bar-social-tmpl.html",
		"views/templates/reside-bar-latest-tmpl.html", "views/templates/reside-bar-widget-tmpl.html")
	if err != nil {
		log.Fatal("Error while parsing main-layout")
	}
}

func parseLayoutWithoutSidebar() {
	var err error
	parsedLayoutWithoutSidebar, err = template.ParseFiles("views/layout/layout-without-sidebar.html", "views/templates/login-signup-tmpl.html",
		"views/templates/footer.html", "views/templates/header.html")
	if err != nil {
		log.Fatal("Error while parsing main-layout")
	}
}

func execTmplToString(tmplName, path string, data interface{}) (string, error) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return "", err
	}
	buff := bytes.Buffer{}
	err = tmpl.ExecuteTemplate(&buff, tmplName, data)
	if err != nil {
		log.Println(err)
	}
	return buff.String(), nil
}

func executeTemplate(w http.ResponseWriter, tmpl *template.Template, name string, data interface{}) {
	tmpl.ExecuteTemplate(w, name, data)
}

func renderMainLayout(w http.ResponseWriter, r *http.Request, t data.TemplateInfo, title string, data interface{}) {
	tmplString, _ := execTmplToString(t.Name, t.Path, data)
	executeTemplate(w, parsedMainLayout, mainLayout, TmplContext{Content: template.HTML(tmplString), Title:title})
}

func renderLayoutWithoutSidebar(w http.ResponseWriter, r *http.Request, t data.TemplateInfo, title string, data interface{}) {
	tmplString, _ := execTmplToString(t.Name, t.Path, data)
	executeTemplate(w, parsedLayoutWithoutSidebar, layoutWithoutSidebar, TmplContext{Content: template.HTML(tmplString), Title:title})
}

type BlogMetadata struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	Title            string `json:"title"`
	TmplFileName     string `json:"tmplFileName"`
	Author           string `json:"author"`
	PublishDate      string `json:"publishDate"`
	Tags             string `json:"tags"`
	CommentCount     int64 `json:"commentCount"`
	ImageUrl         string `json:"imageUrl"`
	CommentsFileName string `json:"commentsFileName"`
	Summary          string `json:"summary"`
}

func GetPostInfo(pName, fileName string, pos int64) (BlogMetadata, error) {
	blogMeta := BlogMetadata{}
	blogs, err := GetAllPostsOfPage(fileName)
	if err != nil {
		return blogMeta, err
	}
	blogMeta = blogs[pos]
	return blogMeta, nil
}

func GetAllPostsOfPage(fileName string) ([]BlogMetadata, error) {
	var blogsMetadata []BlogMetadata
	_, err := os.Stat("data/blogs/" + fileName)
	if err != nil {
		return blogsMetadata, err
	}
	blogsData, err := ioutil.ReadFile("data/blogs/" + fileName)
	if err != nil {
		return blogsMetadata, err
	}
	if err := json.Unmarshal(blogsData, &blogsMetadata); err != nil {
		return blogsMetadata, err
	}
	return blogsMetadata, nil
}

func GetFileName(page int64) (string, int64, int64) {
	var fileName string
	var t, max int64 = 1, data.BlogMetaInfo.TotalBlogs
	if page > 0 && pagination <= data.BlogMetaInfo.TotalBlogs {
		t = (page * pagination) / 20
		fileName = "blog-" + strconv.FormatInt(t * 20 + 1, 10) + "-to-" + strconv.FormatInt(t * 20 + 20, 10) + ".json"
		max = pagination
	} else {
		page = 1
		t = 1
		max = data.BlogMetaInfo.TotalBlogs
		fileName = "blog-1-to-20.json"
	}
	return fileName, max, t
}