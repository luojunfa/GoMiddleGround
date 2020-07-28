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
	size:=3//切片大小为三
	var num int//要查找的整形数
    _,err := fmt.Scanf("%d",&num)
	ctx:=context.Background()
	ResultChan:=make(chan bool)
	// wg:=new(sync.WaitGroup)
    // wg.Add(1)//通知程序有一个等待执行完成的任务
    // wg.Done()//表示正在等待的程序已执行完成
	// wg.Wait()//阻碍当前的程序直到等待的程序执行完成
	timer :=time.NewTimer(time.Second*5)
    for i := 0; i <len(number); i += size {
        end := i + size
        if end >= len(number) {
            end = len(number) - 1
        }
        go SearchTarget(ctx, number[i:end], num, ResultChan)//每次传入大小为三的切片
    }
    //做监控：时间耗尽或者说找到目标
	select {
	case <-timer.C://时间耗尽
		fmt.Println("Time! Not,found!",err)
	case <-ResultChan://通道里面的值为true时，即为找到了结果
        fmt.Println("Found,it!")
    // default:跳出监控，如果需要实现case里面的某一个条件下的语句必被执行，就不要加入default
	}
	

}
func SearchTarget(ctx context.Context, data []int, target int, resultChan chan bool) {
    for _, v := range data {//不需要的值用“_”做忽略
        select {
        case <- ctx.Done()://链接断开，任务取消
            fmt.Printf( "Task cancelded! \n")
            return
        default://多次执行，不存在链接断开的情况那么就脱离
        }
       
        fmt.Printf( "v: %d \n", v)
        time.Sleep(time.Millisecond * 1500)//耗时
        if target == v {
            resultChan <- true//找到了通道值为true
            return
        }
    }

}
