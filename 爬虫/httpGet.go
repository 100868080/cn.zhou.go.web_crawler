package main

import (
	"fmt"
	"net/http"
	"os"
)

func printError(err error,msg string){
	if err!=nil {
		fmt.Print(msg)
		panic(err)
	}
}

func main() {
	//  https://www.jjshu.net/3/3022/11405055.html
	url :="https://www.jjshu.net/3/3022/11405055.html"
	response,err1:=http.Get(url)
	printError(err1,"get请求错误")
	//http.Client{}
	buf:=make([]byte,4*1024)
	n,err:=response.Body.Read(buf)

	defer response.Body.Close()

	printError(err,"读取ｂｏｄｙ内容出错")

	fmt.Println(string(buf[:n]))
	text:=string(buf[:n])

	file,err:=os.Create("/home/zhou/a.txt")
	printError(err,"写文件错误")
	file.WriteString(text)
}
