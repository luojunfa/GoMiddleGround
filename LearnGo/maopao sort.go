package main
import "fmt"
// 冒泡排序
func  maopaosort(value []int)  {
	for i:=0;i<len(value)-1;i++{
		for j :=i+1; j < len(value); j++ {
			if value[i]>value[j]{
				// a:=value[i]
				// value[i]:=value[j]
				value[i],value[j]=value[j],value[i]//go直接做值的对换
				// value[j]:=a
			}
		}
	}
	fmt.Println("冒泡排序：",value)//打印结果
}
// 选择排序
func selectsort(value []int)  {
	for i := 0; i < len(value); i++ {
		for j := 0; j < i; j++ {
			if value[i] <value[j]{
				value[j],value[i]=value[i],value[j]
			}
		}
	}
	fmt.Println("选择排序：",value)
}
func main()  {
	value :=[]int{4,6,3,8,7,5,12,17,11}
	maopaosort(value)//冒泡排序
	selectsort(value)
}

