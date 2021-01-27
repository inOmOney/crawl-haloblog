package write

import (
	"bufio"
	"fmt"
	"os"
)

func WriteFile(content string) error {
	filePath := "e:/golang.md"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	_, err = write.WriteString(content)
	return err
}
