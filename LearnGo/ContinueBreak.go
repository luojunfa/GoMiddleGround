package main
func main()  {
L1:
	for i:=0;i<=5;i++{
L2:
		for j:=0;j<3;j++{
			if j>2 {continue L2}//j>2时，跳出小循环，继续下一次大循环，go语言中，continue只能够用于for循环
			if i>1{break L1}//i>1时，跳出变量定义的作用域循环，break可以用于多中，比如for，switch，select，
			print(i,":",j," ")
		}		
		println()
	}
}

