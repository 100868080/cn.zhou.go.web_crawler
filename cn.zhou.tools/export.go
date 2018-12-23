package cn_zhou_tools

type Export struct {
	print
}

type print interface {
	PrintError(err error)
	PrintMoreError(err error,message string)
}


func (p Export) PrintError(err error){
	if  err!=nil {
		panic(err)
	}
}
//err error,message string
func (p Export)  PrintMoreError(err error,message string){
	if  err!=nil {
		println( message)
		panic(err)
	}
}
