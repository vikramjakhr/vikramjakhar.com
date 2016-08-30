package controllers

import (
	"net/http"
	"html/template"
	"strings"
	"github.com/vikramjakhr/vikramjakhar.com/data"
)

func init() {
	registerHandlers()
}

func Root(w http.ResponseWriter, r *http.Request) {
	templateInfo := data.TemplateInfoMap["home"]
	renderMainLayout(w, r, templateInfo, templateInfo.Title, data.Summary)
}

func About(w http.ResponseWriter, r *http.Request) {
	templateInfo := data.TemplateInfoMap["about"]
	renderLayoutWithoutSidebar(w, r, templateInfo, templateInfo.Title, nil)
}

func Golang(w http.ResponseWriter, r *http.Request) {
	templateInfo := data.TemplateInfoMap["golang"]
	renderMainLayout(w, r, templateInfo, templateInfo.Title, data.Summary)
}

func Post(w http.ResponseWriter, r *http.Request) {
	s := strings.Split(r.RequestURI, "/")
	tmplName := strings.TrimSpace(s[2])
	templateInfo := data.TemplateInfoMap[tmplName]
	post, err := execTmplToString(templateInfo.Name, templateInfo.Path, data.Summary[tmplName])
	if err != nil {
		NotFound(w, r)
		return
	}
	renderMainLayout(w, r, data.TemplateInfoMap["post"], templateInfo.Title, template.HTML(post))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("views/err/404.html")
	tmpl.Execute(w, nil)
}

func InternalError(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("views/err/500.html")
	tmpl.Execute(w, nil)
}