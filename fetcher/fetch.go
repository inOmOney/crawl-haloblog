package fetcher

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

//实现浏览器抓包拿到下面的数据
const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"
const Authorization = "xxxxxxxxxxxxxxx" // 这个经常变
const cookie = "xxxxxxxxxxxxxxxxxxxxx"

func Fetch(url string) ([]byte, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Cookie", cookie)
	request.Header.Add("User-Agent", userAgent)
	request.Header.Add("Admin-Authorization", Authorization)

	resp, err := client.Do(request)
	reader := bufio.NewReader(resp.Body)
	determinEncoding := DeterminEncoding(reader)
	utf8content := transform.NewReader(reader, determinEncoding.NewDecoder())

	return ioutil.ReadAll(utf8content)
}

/*
	识别文本的编码方式
*/
func DeterminEncoding(r *bufio.Reader) encoding.Encoding {
	peek, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(peek, "")
	return e
}
