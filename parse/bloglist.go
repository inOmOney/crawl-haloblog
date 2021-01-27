package parse

import (
	"crawl-blog/engin"
	jsoniter "github.com/json-iterator/go"
)

// content: 从主页获取的json (所有的文章id和title)
// 主要功能: 将所有的文章id对应的URL返回到新的Request中
func ParseList(content []byte) (engin.Article, bool) {
	list := jsoniter.Get(content, "data", "content")
	//获取第一个
	size := list.Size()
	needcontent := []byte(list.ToString())
	article := engin.Article{}

	for i := 0; i < size; i++ { //将Url添加完成
		url := "http://www.zhb.cool/api/admin/posts/" + jsoniter.Get(needcontent, i, "id").ToString()
		article.Request = append(article.Request, engin.Request{
			Url:       url,
			ParseFunc: Content,
		})
	}
	return article, true
}
