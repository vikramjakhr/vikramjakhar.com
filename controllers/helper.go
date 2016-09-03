package controllers

import (
	"net/http"
	"html/template"
	"bytes"
	"github.com/choudhary92/vikramjakhar.com/data"
	"log"
)

const (
	mainLayout string = "main-layout"
	layoutWithoutSidebar string = "layout-without-sidebar"
)

type TmplContext struct {
	Content template.HTML
	Title   string
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

func parseMainLayout() (*template.Template, error) {
	tmpl, err := template.ParseFiles("views/layout/main-layout.html", "views/templates/login-signup-tmpl.html",
		"views/templates/footer.html", "views/templates/header.html", "views/templates/reside-bar-tmpl.html",
		"views/templates/reside-bar-search-tmpl.html", "views/templates/reside-bar-social-tmpl.html",
		"views/templates/reside-bar-latest-tmpl.html", "views/templates/reside-bar-widget-tmpl.html")
	return tmpl, err
}

func parseLayoutWithoutSidebar() (*template.Template, error) {
	tmpl, err := template.ParseFiles("views/layout/layout-without-sidebar.html", "views/templates/login-signup-tmpl.html",
		"views/templates/footer.html", "views/templates/header.html")
	return tmpl, err
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
	tmpl, err := parseMainLayout()
	if err != nil {
		InternalError(w, r)
		return
	}
	executeTemplate(w, tmpl, mainLayout, TmplContext{Content: template.HTML(tmplString), Title:title})
}

func renderLayoutWithoutSidebar(w http.ResponseWriter, r *http.Request, t data.TemplateInfo, title string, data interface{}) {
	tmplString, _ := execTmplToString(t.Name, t.Path, data)
	tmpl, err := parseLayoutWithoutSidebar()
	if err != nil {
		InternalError(w, r)
		return
	}
	executeTemplate(w, tmpl, layoutWithoutSidebar, TmplContext{Content: template.HTML(tmplString), Title:title})
}