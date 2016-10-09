package controllers

import (
	"net/http"
	"html/template"
	"strings"
	"bitbucket.org/vikramjakhr/vikramjakhar.com/data"
	"strconv"
	"sort"
)

func init() {
	registerHandlers()
	parseMainLayout()
	parseLayoutWithoutSidebar()
}

var pagination int64 = 5

func Root(w http.ResponseWriter, r *http.Request) {
	Pagination(w, r)
}

func Pagination(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s := strings.Split(r.RequestURI, "/")
		page, _ := strconv.ParseInt(s[len(s) - 1], 10, 0)
		fileName, _, _ := GetFileName(page)
		blogs, err := GetAllPostsOfPage(fileName)
		if err != nil {
			NotFound(w, r)
			return
		}
		sort.Sort(Blogs(blogs))
		start := page * pagination
		end := start + pagination

		if start >= data.BlogMetaInfo.TotalBlogs {
			start = 0;
			end = pagination;
		}
		if end > data.BlogMetaInfo.TotalBlogs {
			end = data.BlogMetaInfo.TotalBlogs
		}
		next := page + 1
		prev := page - 1
		if (next * pagination) >= data.BlogMetaInfo.TotalBlogs {
			next = -1
		}
		templateInfo := data.TemplateInfo{Title:"Vikram Jakhar", Name:"summary-tmpl", Path:postsPath + "summary-tmpl.html"}
		data := Data{Data:blogs[start:end], Next:next, Prev:prev}
		renderMainLayout(w, r, templateInfo, templateInfo.Title, data)
	default:
		http.Error(w, "Method not Allowd", http.StatusMethodNotAllowed)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	templateInfo := data.TemplateInfo{Title:"About | Vikram Jakhar", Name:"about-tmpl", Path:aboutPath + "about-tmpl.html"}
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
	post, err := execTmplToString(blogMeta.Name, postsPath + blogMeta.TmplFileName, blogMeta)
	if err != nil {
		NotFound(w, r)
		return
	}
	tmplInfo := data.TemplateInfo{Title:blogMeta.Title, Name:"post-tmpl", Path:postsPath + "post-tmpl.html"}
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