package data

func init() {
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

func initBlogMetadata() {
	metadata := make(map[string]string)
	metadata["introduction-to-go"] = "0,blog-1-to-20.json"
	metadata["arrays-in-go"] = "1,blog-1-to-20.json"
	metadata["working-with-slices-in-go"] = "2,blog-1-to-20.json"
	BlogMetaInfo = BlogMetadata{(int64(len(metadata))), metadata}
}