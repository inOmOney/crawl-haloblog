package write

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//需要保存在那个文件夹下面
const filepath = "I:\\DESKTOP\\blog"

//在halo上图片存储的路径
const oldImgPath = "http://www.zhb.cool/upload/"

//需要转移到的自己图床，提前去halo部署的服务器把图片拉取下来,按原路径上床到图床
const newImgPath = "https://gitee.com/BzmAi/picture-bed/raw/master/"

func WriteFile(content []string) {
	replaceContent := strings.Replace(content[2], oldImgPath, newImgPath, -1)
	replaceContent = strings.Replace(replaceContent, "http://www.theanything.top/upload/", "https://gitee.com/BzmAi/picture-bed/raw/master/", -1)
	var filename string
	if content[0] != "" {
		filename = filepath + "\\" + content[0]
		if !Exists(filename) {
			os.Mkdir(filename, os.ModePerm)
		}
	} else {
		filename = filepath
	}
	filename = filename + "\\" + content[1] + ".md"

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("%s 文件打开失败\n", content[0])
	}
	//及时关闭file句柄
	defer file.Close()

	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	_, err = write.WriteString(replaceContent)
	if err != nil {
		fmt.Println(err)
	}
	write.Flush()
	if err != nil {
		fmt.Printf("%s 写入失败\n", content[1])
	} else {
		fmt.Printf("%s 写入成功\n", content[1])
	}
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
