package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)
import "strconv"

func httpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1 //将封装函数内部的错误，传出给调用者
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	//读数据（循环），传出给调用者
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取网页已完成")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2 //将封装函数内部的错误，传出给调用者
			return
		}
		//累加每一次循环读到的buf数据，存到result中，一次性返回
		result += string(buf[:n])
	}
	return
}

//爬取页面操作
func working(start, end int) {
	fmt.Print("正在爬取第%d页到%d页......\n", start, end)
	//循环爬取每一页的数据
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		result, err := httpGet(url)
		if err != nil {
			fmt.Println("httpGet err:", err)
			continue
		}
		fmt.Println("result=", result)
		//将读到的网页数据，保存成文件
		f, err := os.Create("第 " + strconv.Itoa(i) + "页 " + ".html")
		if err != nil {
			fmt.Println("Create err:", err)
			continue
		}
		f.WriteString(result)
		f.Close()  //保存完一个文件，关闭一个文件
	}
}

func main() {
	//指定爬取起始 终止
	var start, end int
	fmt.Print("请输入爬取的起始页(>=1):")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的终止页(>=1):")
	fmt.Scan(&end)

	working(start, end)
}
