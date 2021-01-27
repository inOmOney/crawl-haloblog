package fetcher

import (
	"bufio"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Fetch(url string) ([]byte, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Cookie", "UM_distinctid=175010c8d1aef-0f10e067d4d9bc-333376b-1fa400-175010c8d1b380; CNZZDATA1279079324=886727204-1602039469-null%7C1602039469; Hm_lvt_19759d82d10c4a2ed8c4cec53706ae13=1609407328,1609853812,1611715766; Hm_lvt_19759d82d10c4a2ed8c4cec53706ae13=1609407328,1609853812,1611646282,1611737974; Hm_lpvt_19759d82d10c4a2ed8c4cec53706ae13=1611737974; Hm_lpvt_19759d82d10c4a2ed8c4cec53706ae13=1611738241")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
	request.Header.Add("Admin-Authorization", "d17f5b32f8ed42cb8c04237e78992be9")

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

//方法1：获取指定字段的值(参数1：字段路径，参数2：解析后的数据)
func GetJsonFieldValue(path []string, data map[string]interface{}) (interface{}, bool) {
	if v, ok := data[path[0]]; ok == true {
		if len(path) == 1 {
			return v, true
		} else {
			value, ok := GetJsonFieldValue(path[1:], v.(map[string]interface{}))
			return value, ok
		}
	} else {
		return nil, false
	}
}

//方法2：获取指定字段的值(参数1：字段路径，参数2：原始json数据)
func GetFieldFromJson(path []string, value []byte) string {
	var temp jsoniter.Any
	for i, v := range path {
		if i == 0 {
			temp = jsoniter.Get(value, v)
			if temp == nil {
				return ""
			}
		} else {
			temp = temp.Get(v)
			if temp == nil {
				return ""
			}
		}
	}

	switch temp.ValueType() {
	case jsoniter.InvalidValue, jsoniter.NilValue, jsoniter.BoolValue, jsoniter.ArrayValue, jsoniter.ObjectValue:
		return ""
	case jsoniter.StringValue:
		return temp.ToString()
	case jsoniter.NumberValue:
		return strconv.Itoa(temp.ToInt())
	}
	return ""
}
