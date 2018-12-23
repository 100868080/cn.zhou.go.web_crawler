package cn_zhou_tools

import (
 	"os"
)

type FileUtil struct {
	Filename string
	//file
}

//
//type  addition bool
//
//type file interface {
//	openNew()(file *os.File)
//	openAddition()(f *os.File)
//}
//
//func Open(f file)(file *os.File){
//
//	return  f.openNew() || f.openAddition()
//
//}


//打开一个文件若该文件不存在则创建一个新文件并打开
func (f FileUtil)openNew()(file *os.File){
	file,err:= os.Open(f.Filename)
	if err !=nil {
		file,err=os.Create(f.Filename)
		Export{}.PrintError(err)
 	}
	return file
 }

//以写跟追加的方式打开文件
func (fu FileUtil)openAddition() *os.File{
	f, err := os.OpenFile( fu.Filename, os.O_WRONLY|os.O_APPEND, 0666)
	if(err!=nil){
		f,err=os.Create(fu.Filename)
		Export{}.PrintError(err)
	}
	return f
}