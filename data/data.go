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
	BlogMetaInfo = BlogMetadata{(int64(len(metadata))), metadata}
}

/*
{
"id": 1472765529445,
"name": "concurrency-in-go",
"title": "Concurrency in Go",
"tmplFileName": "concurrency-in-go.html",
"author": "Vikram",
"publishDate": 21315465,
"tags": "Golang",
"commentCount": 0,
"imageUrl": "concurrency-in-go.png",
"commentsFileName": "concurrency-in-go.json",
"summary": "A slice is a data structure that is very similar to an array, but without any specified length. Slices are abstraction built on top of an array type that provides a more convenient way of working with collections"
},*/
