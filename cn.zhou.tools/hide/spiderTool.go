package hide

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type Do struct {  //结构体里面的字段也需要首字母大写才能被其他包调用
	Filename string
	InitUrl string
	UrlsSelector string
	TopicSelector string
	TextSelector string
}

func (do Do) DoWork(){
	var resp http.Response
	var capital,text,url string

	resp=spider(do.InitUrl)
	capital,text,url=parseDoc(resp,do.TopicSelector,
		do.TextSelector,do.UrlsSelector)
	writeData(do.Filename,capital,text )

	i:=3  //url数组的下一个url的索引
	for url!=nil{
		resp=spider(url[i])
		capital,text,url=parseDoc(resp,do.TopicSelector,
			do.TextSelector,do.UrlsSelector)
		writeData(do.Filename,capital,text )
		fmt.Println(url[i],"ok")
	}
}


//发送请求，获取响应
func spider(url string)(resp  http.Response){

	client :=http.Client{} //创建client
	get,err1:=http.NewRequest("GET",url,nil) //创建请求
	Export{}.PrintError(err1) //处理错误
	//添加请求参数
	get.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36" +
		" (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36")
	//发送请求
	resp, err2 := client.Do(get)

 	Export {}.PrintMoreError(err2,"response.错误")

	//defer resp.Body.Close()

	return resp
}

//解析响应体,返回写数据
func parseDoc(resp http.Response,topicSelector,textSelector,urlsSelector string)(
	capital,text,url[] string) {
	doc,err1:=goquery.NewDocumentFromReader(resp.Body)
	//defer resp.Body.Close()
	Export{}.PrintError(err1) //处理错误

	capital=doc.Find(topicSelector).Text() //书名
	text=doc.Find(textSelector).Text()
	url=make([]string,10)
	//finallyUrl:=make([]string,10)
	doc.Find(urlsSelector).Each(func(i int, selection *goquery.Selection) {
		url[i],_=selection.Attr("href")  //将遍历出来的url装入切片
	})
	return capital,text,url
}


func writeData(filename,capital,text string)  {
	f:=OpenAddition(filename)
	f.WriteString(capital)
	f.WriteString(text)
	f.WriteString("\n\n")
	//defer f.Close()
}