package main

import (
	"fmt"
)

func main() {

	fmt.Println("hello word!你好，世界！")

	for i := 0; i < 10; i++ {
		v := i
		fmt.Println(&v)
	}

}

const PI = 3.14

var name = "gopher"

type newType int

type gopher struct{}

type golang interface{}
