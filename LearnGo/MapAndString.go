package main
import "fmt"
import "sort"
func main()  {
	//map-ages
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
		"john":44,
		"jack":55,
		"lucy":23,
	}
	//struct-employee
	type employee struct{
		ID int
		Name string
		Salary float32
		Add string
	}
	// Go语言初始化结构体的两种方式：
	// 方式1：
	// var employee01 employee
	// employee01.ID=12
	// employee01.Name="jack"
	// employee01.Salary=10000.0
	// employee01.Add="jiangxi"
	// 方式2：
	var employee01 employee=employee{12,"jack",10000.0,"jiangxi"}
	// 列表
	var names []string
	for name := range ages {//range方法会获取map的索引也就是Key值
		names = append(names, name)//从尾部添加新的，将ages的索引（key）扩展的列表names
		fmt.Print(name,"\n")//每次的结果否是不一样的，map的遍历是无序的
	}
	sort.Strings(names)//转为String
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])//通过列表的有序来实现map遍历的有序化
	}
	// for _,name:=range ages{
	// 	name01:=string(name)
	// 	fmt.Print(ages[name01])
	// }
	fmt.Print(employee01)
	}

