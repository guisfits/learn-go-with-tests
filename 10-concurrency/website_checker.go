package concurrency

type WebsiteChecker func(string) bool

type Result struct {
	string
	bool
}

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan Result)

	for _, url := range urls {
		go func(u string) {
			resultsChannel <- Result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <- resultsChannel
		results[r.string] = r.bool
	}

	return results
}

