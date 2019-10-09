package resource

type Resource struct {
	Document *Document `json:"document"`
	Meta     *Meta     `json:"meta"`
}

// type raw struct {
// 	meta meta `json:"meta"`
// 	document []byte `json:"document"`
// }

type Meta struct {
	HTTPStatus int    `json:"http_status"`
	URL        string `json:"url"`
}

type Document struct {
	Title     string   `json:"title"`
	Content   []string `json:"content"`
	WordCount int
	Topics    []string
}
