package data

func init() {
	initTemplateInfo()
	initSummaryData()
}

type TemplateInfo struct {
	Title string
	Name  string
	Path  string
}

var TemplateInfoMap map[string]TemplateInfo

func initTemplateInfo() {
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
}

type SummaryData struct {
	Title            string
	Name            string
	URI              string
	Author           string
	AuthorProfileURI string
	DatePublished    string
	Tag              string
	CmntCount        int
	ImageURI         string
	ID               string
	Summary          string
}

var Summary map[string]SummaryData

func initSummaryData() {
	Summary = map[string]SummaryData{
		"introduction-to-go":introToGo(),
		"arrays-in-go":arraysInGo(),
		"working-with-slices-in-go":workingWithSlicesInGo(),
	}
}

func introToGo() SummaryData {
	return SummaryData{
		Title:TemplateInfoMap["introduction-to-go"].Title,
		Name:TemplateInfoMap["introduction-to-go"].Name,
		URI:"/post/introduction-to-go",
		Author:"Vikram",
		AuthorProfileURI:"/about",
		DatePublished:"9:40",
		Tag:"Go",
		CmntCount:0,
		ImageURI:"images/introduction-to-go.jpg",
		ID:"summary-introduction-to-go",
		Summary:`As in modern era of rapid development, developers have to make an
		uncomfortable choice between rapid development and high performance
		while choosing a language for their projects. Languages like C
		and C++ offer fast execution, whereas languages like Python and Ruby
		offers rapid development.
		But the Go team went beyond to solve this problem of developers
		and offers a high-performance language with features that make
		development fast.`,
	}
}

func arraysInGo() SummaryData {
	return SummaryData{
		Title:TemplateInfoMap["arrays-in-go"].Title,
		Name:TemplateInfoMap["arrays-in-go"].Name,
		URI:"/post/arrays-in-go",
		Author:"Vikram",
		AuthorProfileURI:"/about",
		DatePublished:"01:40",
		Tag:"Go",
		CmntCount:0,
		ImageURI:"images/arrays-in-go.png",
		ID:"summary-arrays-in-go",
		Summary:`Go provides three types of data structures to manage data
                collections: arrays, slices, and maps. An array is a fixed-length data type that contains the
                sequence of elements of a single type. An array is
                declared by specifying the data type and the length.`,
	}
}

func workingWithSlicesInGo() SummaryData {
	return SummaryData{
		Title:TemplateInfoMap["working-with-slices-in-go"].Title,
		Name:TemplateInfoMap["working-with-slices-in-go"].Name,
		URI:"/post/working-with-slices-in-go",
		Author:"Vikram",
		AuthorProfileURI:"/about",
		DatePublished:"03:40",
		Tag:"Go",
		CmntCount:0,
		ImageURI:"images/working-with-slices-in-go.png",
		ID:"summary-working-with-slices-in-go",
		Summary:`A slice is a data structure that is very similar to an array, but without any specified length.
                Slices are abstraction built on top of an array type that provides a more convenient way of working with collections.
                Unlike regular arrays where array size remains fixed, slices are dynamic in nature. Which means the length of
                slices can be changed at a later stage as size of data increases or shrinks.`,
	}
}