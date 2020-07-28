package main
// import "fmt"
// import "os"
import "sync"
func main()  {
//1，Channel test1 普通channel
	// data:=make(chan int)//定义长度之打印顺序发生变化
	// exit:=make(chan bool)
	// go func ()  {
	// 	for d:=range data {
	// 		fmt.Println(d)
	// 	}
	// 	fmt.Println("recv over.")
	// 	exit<-true
	// }()
	// data<-1
	// data<-2
	// data<-3
	// close(data)
	// fmt.Println("send over.")
	// <-exit


//2，Channel test2 通道阻塞问题
//在执行时，类似生产者和消费者的关系，其中gofunc是消费者
//因为程序是顺序，当消费协程运行之前的生产数量超过3，就会出现阻塞死锁
	// data:=make(chan int, 3)//大小为3，参考带宽，就是说一次能写入的最大数量是三，但是，通道能容纳的整形数不只是三个
	// exit:=make(chan bool)
	// go func ()  {//在生产写入之前应该先去定义好消费协程，然后在写生产
	// 	for d:=range data{
	// 		 println(d)
	// 	}
	// 	println("recv over.")
	// 	exit<-true
	// }()
	// data<-1//向通道发送数据，类似生产
	// data<-2
	// data<-3
	// data<-4
	// data<-5
	// close(data)
	// println("send over.")
	// <-exit

//3，单向通道
// c:=make(chan int, 3)
// var send chan<-int=c//只发送
// var recv <-chan int=c//只接收
// send <-1
// <-recv

//4，select 语句同时处理多个channel
// a,b:=make(chan int,3),make(chan int)
// go func ()  {
// 	v,ok,s:=0,false,""
// 	for{
// 		select{
// 		case v,ok=<-a:s="a"//在做case判断时，随机选择可用通道接收数据
// 		case v,ok=<-b:s="b"
// 		}
// 	if ok{//当通道内接收到值，ok为true，否则为false
// 		println(s,v)
// 	}else{
// 		os.Exit(0)//ok为false，跳出
// 	}
// 	}
// }()
// for i :=0; i< 5; i++ {
// 	select{
// 	case a<-i:
// 	case b<-i:
// 	}
// }
// close(a)//关闭通道

// select{}//阻塞？？？
//5，使用channel实现信号量
wg:=sync.WaitGroup{}
wg.Add(3)
sem:=make(chan int, 1)
for i:=0; i<3;i++{
	go func (id int)  {
		// defer wg.Done()
		sem<-i
		for j:= 0; j < 3; j++ {
			println(id,j)
		}
		<-sem//对通道做里面的值做处理，防止阻塞
		wg.Done()
	}(i)
}
wg.Wait()//等待操作的完成
}
//将通道作为返回值                                                                          
