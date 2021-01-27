package main

import (
	"crawl-blog/engin"
	"crawl-blog/parse"
)

func main() {
	engin.Run(engin.Request{
		Url:       "http://www.zhb.cool/api/admin/posts?page=0&size=100",
		ParseFunc: parse.ParseList,
	})
}
