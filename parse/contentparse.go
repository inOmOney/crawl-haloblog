package parse

import (
	"crawl-blog/engin"
	"crawl-blog/write"
	jsoniter "github.com/json-iterator/go"
)

//content: 拿到每个文章的origin 、title
//function: 写入文件
func Content(content []byte) (engin.Article, bool) {

	json1 := jsoniter.Get(content, "data")
	title_content := []byte(json1.ToString())

	title := jsoniter.Get(title_content, "title").ToString()
	originalContent := jsoniter.Get(title_content, "originalContent").ToString()

	list := jsoniter.Get(title_content, "categories")
	needcontent := []byte(list.ToString())
	categories := jsoniter.Get(needcontent, 0, "name").ToString()
	if categories == "默认" {
		categories = ""
	}

	msgs := []string{
		categories,
		title,
		originalContent,
	}

	write.WriteFile(msgs)

	return engin.Article{}, false
}
