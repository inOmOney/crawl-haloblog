package parse

import (
	"crawl-blog/engin"
	jsoniter "github.com/json-iterator/go"
)

//content: 拿到每个文章的origin 、title
//function: 写入文件
func Content(content []byte) (engin.Article, bool) {

	data := make(map[string]interface{})
	err := jsoniter.Unmarshal(content, &data)
	if err != nil {
		panic(err)
	}

	//originalContent := data["originalContent"]
	//title := data["title"]

	return engin.Article{}, false
}
