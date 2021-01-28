package main

import (
	"crawl-blog/engin"
	"crawl-blog/parse"
)

const seed = "http://www.zhb.cool/api/admin/posts?page=0&size=100"

func main() {
	engin.Run(engin.Request{
		Url:       seed,
		ParseFunc: parse.ParseList,
	})
}
