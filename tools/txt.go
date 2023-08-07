package tools

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteTxt(filePath string, content []string) {
	fmt.Println(len(content))
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for _, c := range content {
		write.WriteString(c + "\n")
	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}

func ReadTxt(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return nil
	}
	defer file.Close() //关闭文本流
	reader := bufio.NewReader(file)
	data := make([]string, 0)
	for {
		// 读取文本数据
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		data = append(data, str)
	}
	return data
}
