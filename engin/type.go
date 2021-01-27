package engin

type (
	Article struct {
		//Title   string
		//Content string
		Request []Request
		Item    []interface{} // title content
	}

	Request struct {
		Url       string                               //根据id拼接
		ParseFunc func(content []byte) (Article, bool) // bool 判断有没有必要加入到Request中进入下一阶段的爬虫
	}
)
