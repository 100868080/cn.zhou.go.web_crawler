package cn_zhou_tools

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type Task struct {
	Url string
	Filename string
	TopicSelector string
	TextSelector string
	UrlsSelector string
}

func (t Task)Execute(){
	//第一次爬取
	resp:=spider(t.Url)
	topic,text,nextUrl:=parseDoc(resp,t.TopicSelector,
		t.TextSelector,t.UrlsSelector)
	writeData(t.Filename,topic,text)
	//重复爬取
 	for nextUrl!=nil  {
		if nextUrl==nil {
			fmt.Println("爬取完毕")
			break
		}
		resp:=spider(nextUrl[3])
		topic,text,nextUrl=parseDoc(resp,t.TopicSelector,
			t.TextSelector,t.UrlsSelector)
		writeData(t.Filename,topic,text)
		fmt.Println(nextUrl[3],"ok")
	}
}

//根据Url　发送请求，返回响应体
func spider(url string) http.Response{
	client:=http.Client{} //创建客户端
	//创建请求
	get,err1:=http.NewRequest("GET",url,nil)
	Export{}.PrintMoreError(err1,"创建请求环节出问题了")
	//添加请求参数
	get.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36" +
		" (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36")
	//发送请求
	resp,err2:=client.Do(get)
	Export{}.PrintMoreError(err2,"发送请求环节出问题了")

	return  *resp

}

//根据得到的响应体，解析目标数据并返回
func parseDoc(resp http.Response,topicSelector,textSelector,
	urlsSelector string)(topic,text string,urls[] string){
	//从响应体中解析数据
	doc,err1:=goquery.NewDocumentFromReader(resp.Body)
	resp.Body.Close()
	Export{}.PrintError(err1)
	topic=doc.Find(topicSelector).Text() //标题
	text=doc.Find(textSelector).Text()  //文本内容

	urls=make([] string,10) //下一页的url
	 //queue:=3
 	doc.Find(urlsSelector).Each(func(i int, selection *goquery.Selection) {
		url,_:=selection.Attr("href")
		urls[i]=url
	})
 	return topic,text,urls
}

//把爬取到的目标数据写入文件中
func writeData(filename,topic,text string){
	f:=FileUtil{filename}.openAddition()
	f.WriteString(topic)
	f.WriteString(text)
	f.WriteString("\n\n")
	defer f.Close()
}



