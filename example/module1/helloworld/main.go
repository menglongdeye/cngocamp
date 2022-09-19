package main

import (
	"fmt"
	"os"
	//"log"
	//"net/http"
	//"reflect"
	//"net/http"
)

//type MyType struct {
//	Name string `json:"name"`
//}
//
//func main() {
//	mt := MyType{Name: "test"}
//	myType := reflect.TypeOf(mt)
//	name := myType.Field(0)
//	tag := name.Tag.Get("json")
//	println(tag)
//}

func main() {
	fmt.Println("Hello World")
	fmt.Println("os args is: ", os.Args)
	fmt.Println("Hello World")

	// （1）Print-format 错误，检查类型不匹配的print
	str := "hello world!"
	kk := 10
	fmt.Printf("%d\n", str)
	fmt.Printf("%d\n", kk)

	//（2）Boolean 错误，检查一直为 true、false 或者冗余的表达式
	i := 1
	fmt.Println(i != 0 || i != 1)

	//（3）Range 循环，比如如下代码主协程会先退出，go routine无法被执行
	words := []string{"foo", "bar", "baz"}
	for _, word := range words {
		go func() {
			fmt.Println(word)
		}()
	}

	//其他错误，比如变量自赋值，error 检查滞后等
	//res, err := http.Get("https://www.spreadsheetdb.io/")
	//defer res.Body.Close()
	//if err != ni{
	//	log.Fatal(err)
	//}

	// Service Type string describes ingress methods for a service
	type ServiceType string
	const (
		ServiceTypeClusterIP    ServiceType = "ClusterIP"
		ServiceTypeNodePort     ServiceType = "NodePort"
		ServiceTypeLoadBalancer ServiceType = "LoadBalancer"
		ServiceTypeExternalName ServiceType = "ExternalName"
	)

}
