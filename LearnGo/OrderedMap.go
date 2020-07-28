package main
import "sort" 
      
func main()  {
	map01:=map[string]int{
		"student01" : 90,
		"student02" : 95,
		"student03" : 97,
		"student04" : 66,
		"student05"	: 69,
	}
	var student []string
	print("有序化前：\n")
	for stu:=range map01{
		student=append(student,stu)
		stuC:=string(stu)//类型转换
		print(stuC," ",":"," ",map01[stuC],"\n")
	}
	sort.Strings(student)
	// fmt.Print(map01)
	print("有序化后：\n")
	for _,stu:=range student{
		print(stu," ",":"," ",map01[stu],"\n")
	}
}

