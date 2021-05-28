package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	//--Цикл запуска горутин - выполнится параллельно
	for _, url := range urls {
		go func(url string) {
			resultChannel <- result{url, wc(url)}
		}(url)
	}
	//-

	//--Цикл приёма результата - выполнится последовательно
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	//-

	return results
}
