package engin

import "crawl-blog/fetcher"

func Run(seeds ...Request) {
	var requests []Request
	for _, e := range seeds {
		requests = append(requests, e)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		json, err := fetcher.Fetch(r.Url)
		if err != nil {
			panic(err)
		}
		article, bool := r.ParseFunc(json)
		if bool {
			requests = append(requests, article.Request...)
		}
	}
}
