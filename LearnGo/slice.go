package main
import "fmt"
func main()  {
	//例1
	data:=[...]int{0,1,2,3,4,5,6}
	slic:=data[0:5:6]
	print(slic[0]," ",slic[1]," ",slic[2]," ",slic[3]," ",slic[4],"\n")
	//例2
	data1:=[...]int{0,1,2,3,4,5,6,7,8,9}
	s:=data1[8:]//从位置8开始切到最后
	s1:=data1[:5]//从0开始切到5，左闭右开
	print(s[0]," ",s[1]," ",s1[0]," ",s1[1]," ",s1[2]," ",s1[3]," ",s1[4],"\n")
	copy(s1,s)//copy，短的替换长的切片的相应位置
	fmt.Print(s,"\n")
	fmt.Print(s1,"\n")
	fmt.Print(data1,"\n")//copy 会对最底层的data1产生影响
	//fmt.Print和print的输出差异。前者可以通过数组名直接打印整个数组的内容，后者打印则需要具体的index打印具体的值
}
