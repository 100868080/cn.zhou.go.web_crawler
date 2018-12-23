// 爬虫.go
package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
)

func printError(err error){
	if(err!=nil){
		panic(err)
	}
}

func s(url string)(nextUrl[] string)  {
	client := http.Client{}                       //创建客户端
	get, err1 := http.NewRequest("GET", url, nil) //创建请求
	//设置请求参数
	//get.Header.Set()=get.Header.Add()
	get.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36"+
		" (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36")
	printError(err1)               //检查错误
	resp, err2 := client.Do(get) //使用client发送请求，并且获取响应体
	printError(err2)               //检查错误
	defer resp.Body.Close()
	//document := goquery.Document{}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	printError(err)

	//s := doc.Selection
	topic :=doc.Find("div.bookname h1").Text() //书名
	text :=doc.Find("div#content").Text()      //正文

	u1 := make([]string,10)
	//urls,b:=doc.Find("div.bottem a").Attr("href")	//urls
	doc.Find("div.bottem a").Each(func(i int, u *goquery.Selection) {
		u2,_:= u.Attr("href")
		u1[i] = u2
		//fmt.Println("url[]:",url[3] ,i)
	})

	f, err := os.OpenFile("/home/zhou/科技霸权.docx",
		os.O_WRONLY|os.O_APPEND, 0666)
	if (err != nil) {
		f, err = os.Create("/home/zhou/科技霸权.docx")
	}

	f.WriteString(topic)
	// f.WriteString("\n")
	f.WriteString(text)
	f.WriteString("\n")
	f.WriteString("\n")

//	defer f.Close()

	return u1
}

func do(url string) {
	nextUrl:=s(url)
	for nextUrl != nil {

		nextUrl=s(nextUrl[3])

		if nextUrl == nil {
			fmt.Println("最后一页url:", nextUrl[3], "ok")
			break;
		}

		fmt.Println("下一页url:", nextUrl[3], "ok")
	}
}

func abc()  {
	fmt.Println("abc")
}

func main() {
	url:="https://www.biquge.info/11_11507/8003059.html"
	   do(url)

}
