package main

import "fmt"

func main(){
	ch1 := make(chan string)
	go SendData(ch1)

	data := <- ch1 //阻塞等待子线程执行
	fmt.Println("主线程",data)

	ch1 <- "main"
	fmt.Println("over")


}

func SendData(ch1 chan string)  {
	ch1 <- "子线程执行"
	fmt.Println("子")
	data := <- ch1
	fmt.Println("子线程",data)

}