package data

func init() {
	initTemplateInfo()
	initBlogMetadata()
}

type TemplateInfo struct {
	Title string
	Name  string
	Path  string
}

type BlogMetadata struct {
	TotalBlogs int64
	Meta       map[string]string
}

var BlogMetaInfo BlogMetadata

/*func initTemplateInfo() {
	TemplateInfoMap = make(map[string]TemplateInfo)
	TemplateInfoMap["introduction-to-go"] = TemplateInfo{Title:"Introduction to Golang",
		Name:"introduction-to-go", Path:"views/posts/introduction-to-go.html"}
	TemplateInfoMap["arrays-in-go"] = TemplateInfo{Title:"Arrays In Golang",
		Name:"arrays-in-go", Path:"views/posts/arrays-in-go.html"}
	TemplateInfoMap["working-with-slices-in-go"] = TemplateInfo{Title:"Working with Slices in Golang",
		Name:"working-with-slices-in-go", Path:"views/posts/working-with-slices-in-go.html"}
	TemplateInfoMap["golang"] = TemplateInfo{Title:"Golang", Name:"summary-tmpl",
		Path:"views/posts/summary-tmpl.html"}
	TemplateInfoMap["post"] = TemplateInfo{Title:"Post", Name:"post-tmpl",
		Path:"views/posts/post-tmpl.html"}
	TemplateInfoMap["about"] = TemplateInfo{Title:"About", Name:"about-tmpl",
		Path:"views/about/about-tmpl.html"}
	TemplateInfoMap["home"] = TemplateInfo{Title:"Vikram Jakhar", Name:"summary-tmpl",
		Path:"views/posts/summary-tmpl.html"}
	TemplateInfoMap["register"] = TemplateInfo{Title:"Register", Name:"register-tmpl",
		Path:"views/user/register-tmpl.html"}
}*/

func initBlogMetadata() {
	metadata := make(map[string]string)
	metadata["introduction-to-go"] = "1,blog-1-to-20.json"
	metadata["arrays-in-go"] = "2,blog-1-to-20.json"
	metadata["working-with-slices-in-go"] = "3,blog-1-to-20.json"
	BlogMetaInfo = BlogMetadata{len(metadata), metadata}
}