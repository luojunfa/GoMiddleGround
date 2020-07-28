package main
func main()  {
	//Map+Struct
	m:=map[int] struct{
		name string
		age int
	}{//前面的数字作为索引
		1:{"jack",22},
		2:{"lucy",23},
	}
	print(m[1].name," ",m[1].age)
	//2
	m01:=map[string]int{
		"a": 1}//声明+初始化
	if v,ok:=m01["a"];ok{
		print(ok,v)//ok是bool型，v是Int
	}
}

