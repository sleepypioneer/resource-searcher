package analyse

var (
	// keywords    = map[string]struct{}{"python": struct{}{}, "webservice": struct{}{}, "server": struct{}{}, "Django": struct{}{}, "list": struct{}{}}
	keywords = []string{"python", "webservice", "server", "Django", "list"}
)

// taken from https://www.golangprograms.com/remove-duplicate-values-from-slice.html
func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// FindTopics retuns a list of topics which the article includes
func FindTopics(words []string) (topics []string) {
	uniqueWords := unique(words)

	mb := make(map[string]struct{}, len(uniqueWords))
	for _, x := range uniqueWords {
		mb[x] = struct{}{}
	}

	for _, keyword := range keywords {
		if _, found := mb[keyword]; found {
			topics = append(topics, keyword)
		}
	}
	return topics
}
