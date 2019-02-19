package main

import (
	"fmt"
	"github.com/hello-world-go/pkg/util"
)

func main() {

	//var message string = "Hello world" //static typing
	//message :="hello world" // dynamic typing

	//m, v := getMessage()
	//m, _ := util.GetMessage()

	//fmt.Printf("%v", m)
	fmt.Printf("%v", util.HelloworldMessage)

	if m, v, error :=util.CreateMessage(2); error != nil{
		fmt.Printf("%v", error)
	}else{
		fmt.Printf("%v %v\n", m, v)
	}
	//println("Hello world!!!")
}

//
//func getMessage() string{
//	message := "hello world"
//
//	return message
//}

//func getMessage() (string, int) {
//
//	message := "Hello world"
//
//	number := 1
//	return message, number
//
//}
