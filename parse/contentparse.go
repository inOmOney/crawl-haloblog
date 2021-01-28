package parse

import (
	"crawl-blog/engin"
	"crawl-blog/write"
	jsoniter "github.com/json-iterator/go"
)

//content: 拿到每个文章的origin 、title
//function: 写入文件
func Content(content []byte) (engin.Article, bool) {

	list := jsoniter.Get(content, "data")
	needcontent := []byte(list.ToString())

	title := jsoniter.Get(needcontent, "title").ToString()
	originalContent := jsoniter.Get(needcontent, "originalContent").ToString()

	jsoniter.Get(content, "data")
	//originalContent := data["originalContent"]
	//title := data["title"]
	msgs := []string{
		title,
		originalContent,
	}
	write.WriteFile(msgs)

	return engin.Article{}, false
}
