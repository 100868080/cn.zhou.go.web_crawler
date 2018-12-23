package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	. "os"
	"regexp"
)

//使用 httpClient 可以添加参数，比如user-agent
//如果只是比较简单的访问使用 httpGet 就可以了，但复杂的请求就只有 httpClient才达到目标了
func printError2(err error,msg string){
	if err!=nil {
		fmt.Println(msg)
		panic(err)
	}
}

func doWork(url string)(result string){
	//创建ｈｔｔｐ客户端
	client:=http.Client{}

	//创建get(或者post)请求
	request,err1:=http.NewRequest("get",url,nil)
	printError2(err1,"ｇｅｔ请求错误")  //处理请求错误

	//设置请求参数
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36" +
		" (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36")
	request.Header.Set("accept-language","zh-CN,zh;q=0.9")

	//发送请求,获取响应
	response, err2 := client.Do(request)
	printError2(err2,"ｇｅｔ请求错误")  //处理错误

		n,err3:=ioutil.ReadAll(response.Body)
		printError2(err3,"Read读取错误")  //处理请错误
		defer response.Body.Close()

	 reg:=regexp.MustCompile(`<div id="content">[{Han}*]</div>`)
		// printError2(err,"regexp错误")

		text :=reg.FindAllSubmatch(n,-1)
	 fmt.Println("reg:",text)

	//return text
	return string(n)

}

func main() {
	url :="https://www.jjshu.net/3/3022/11405055.html"
	//url:="https://www.biquge.cc/html/318/318239/1608282.html"
	text:=doWork(url)

	filename:="/home/zhou/a.txt"
	//创建文件
	file, err4 := Create(filename)
	printError2(err4,"文件ｃｒｅａｔｅ错误")  //处理请错误

	//写入文件
	file.WriteString(text)


	//ioutil.WriteFile(filename,text,nil)

	fmt.Print("ok")


}
