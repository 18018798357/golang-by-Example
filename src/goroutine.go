package main

import (
	"fmt"
	"time"
)


func main(){
	ch := make(chan string)


	go getData(ch)
	go senData(ch)
	time.Sleep(9)
}


func senData(ch chan string){
	ch <- "Washington"
	ch <- "Tripoil"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}


func getData(ch chan string){
	var input string
	var input2 string
		input = <- ch
	input2 = <- ch
		fmt.Printf("%s\n",input)
		fmt.Printf("%s\n",input2)

}