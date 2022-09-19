package main

import (
	"flag"
	"fmt"
	//"net/http"
	"os"
)

func main() {
	name := flag.String("name", "zhangsan", "test param: ")
	flag.Parse()

	fmt.Println("os args is:", os.Args)
	fmt.Println("input param is:", *name)

	fullStrings := fmt.Sprintf("hello %s from go\n", *name)
	fmt.Println(fullStrings)
}

func passValue(a, b int) (x, y int) {
	x = a
	y = b

	return x, y
}
