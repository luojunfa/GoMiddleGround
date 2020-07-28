package main
import "fmt"
func main()  {
	a:=[]int{1,2,3}
	print(len(a),"\n")
	x:=1
	p:=&x
	fmt.Print(p,"\n")
	*p=2
	fmt.Print(x)
}