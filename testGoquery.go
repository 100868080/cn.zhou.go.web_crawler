package main

//reflect.TypeOf(url)　根据数据获取数据的类型
// fmt.Println("urls:",reflect.TypeOf(urls))
import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
)


func printErr(err error){
	if(err!=nil){
		panic(err)
	}
}

func spider(url string)(resp2 io.Reader){
	client:=http.Client{}  //创建客户端
	get,err1:=http.NewRequest("GET",url,nil)//创建请求

	//设置请求参数
	//get.Header.Set()=get.Header.Add()
	get.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36" +
		" (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36")
	printErr(err1)  //检查错误
	resp,err2:=client.Do(get)   //使用client发送请求，并且获取响应体
	printErr(err2) //检查错误

	//如果响应体数据过大必须使用ioutil.ReadAll(resp)
	//ioutil.ReadAll(resp)返回的是一个字节数组的长度
	// 否则读取到的数据会丢失
	//resp.Body,是响应体的的正文部分
	//n,err:=ioutil.ReadAll(resp.Body)
	//printErr(err)
	//defer  resp.Body.Close()

	//fmt.Println(string(n))

	//return string(n)
	return resp.Body
}


func parseDoc(html io.Reader) (topic,text string,url[]  string ) {
	//document := goquery.Document{}
	doc,err:=goquery.NewDocumentFromReader(html)
	printErr(err)

	//s := doc.Selection
	topic= doc.Find("div.bookname h1").Text() //书名
	text =doc.Find("div#content").Text()       //正文

	url=make([]string,10)
	//urls,b:=doc.Find("div.bottem a").Attr("href")	//urls
	doc.Find("div.bottem a").Each(func(i int, u *goquery.Selection) {
		 urls,_:=u.Attr("href")
		 url[i]=urls
		//fmt.Println("url[]:",url[3] ,i)
	})

	return topic,text,url
}

//把爬取到的内容写入文件
func writeFile(topic,text string){

	//f,err:=os.Open("/home/zhou/科技霸权.txt")
	//以写跟追加的方式打开文件
	f, err := os.OpenFile("/home/zhou/科技霸权.docx",
		os.O_WRONLY|os.O_APPEND, 0666)
	if(err!=nil){
		f,err=os.Create("/home/zhou/科技霸权.docx")
	}

	f.WriteString(topic)
	// f.WriteString("\n")
	f.WriteString(text)
	f.WriteString("\n")
	f.WriteString("\n")
	defer f.Close()
}

func doWork(url string){

 	html:= spider(url)
	topic,text,nextUrl:=parseDoc(html)
	writeFile(topic,text)

	for  nextUrl!=nil {
		if nextUrl==nil{
			fmt.Println("最后一页url:",nextUrl[3],"ok")
			break;
		}
		html= spider(nextUrl[3])
		topic,text,nextUrl=parseDoc(html)
		writeFile(topic,text)
		fmt.Println("下一页url:",nextUrl[3],"ok")

	}

}

var url="https://www.biquge.info/11_11507/8003059.html"
func main() {
	doWork(url)
}
