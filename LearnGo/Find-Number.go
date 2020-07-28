package main
import (
	// "sync"
// "bufio"
// "os"
"time"
"fmt"
"context"
// "strconv"
)
func main()  {
	number:=[]int{10,23,45,78,45,20,30,45,61,22,335,96,23,32,178,421,962,852,231,546,126,462}
	print("请输入你要找的数字")
	var num int
    _,err := fmt.Scanf("%d",&num)
	ctx:=context.Background()
	ResultChan:=make(chan bool)
	// wg:=new(sync.WaitGroup)
    // wg.Add(1)//通知程序有一个等待执行完成的任务
    // wg.Done()//表示正在等待的程序已执行完成
	// wg.Wait()//阻碍当前的程序直到等待的程序执行完成
	timer :=time.NewTimer(time.Second*5)
    for i := 0; i <len(number); i++ {
	   search(ctx,number,num,ResultChan)
    }
	select {
	case <-timer.C:
		fmt.Println("Time! Not,found!",err)
		// cancel()
	case <-ResultChan:
		fmt.Println("Found,it!")
		// cancel()
	}
	

}
func  search(ctx context.Context,data []int, target int, resultChan chan bool)  {
	for _,num:=range data{
		select {
			case <- ctx.Done()://链接断开，任务取消
				fmt.Printf( "Task cancelded! \n")
				return
			default:
			}
		time.Sleep(time.Millisecond * 1500)//耗时
		if num==target{
			resultChan<-true
		}
	}
}
