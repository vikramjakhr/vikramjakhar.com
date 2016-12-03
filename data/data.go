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
	metadata["go-types"] = "3,blog-1-to-20.json"
	metadata["map-in-go"] = "4,blog-1-to-20.json"
	metadata["working-with-date-time-in-go"] = "5,blog-1-to-20.json"
	metadata["sort-in-golang"] = "6,blog-1-to-20.json"
	metadata["string-joiner-in-java8"] = "7,blog-1-to-20.json"
	metadata["ssl-certificate-installation-in-nginx"] = "8,blog-1-to-20.json"
	metadata["spring-boot-actuator"] = "9,blog-1-to-20.json"
	BlogMetaInfo = BlogMetadata{(int64(len(metadata))), metadata}
}