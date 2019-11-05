package analyse

var (
	// keywords    = map[string]struct{}{"python": struct{}{}, "webservice": struct{}{}, "server": struct{}{}, "Django": struct{}{}, "list": struct{}{}}
	keywords = []string{"python", "webservice", "server", "Django", "list"}
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// FindTopics retuns a list of topics which the article includes
func FindTopics(words map[string]int) (topics []string) {
	for _, keyword := range keywords {
		if _, found := words[keyword]; found {
			topics = append(topics, keyword)
		}
	}
	return topics
}
