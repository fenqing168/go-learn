package main

import (
	"fmt"
)

func main() {

	/*
		1. var var1 string; var <变量名> <类型>
		2. var var2 string = "string2"; var <变量名> <类型> = <值>
		3. var var2 = "string2"; var <变量名> = <值>，根据值进行推断类型
		4. var (
			<变量名> <类型>
			...
		)
		5. <变量名> := <值>
		6. <变量名>... := <值>...
		7. 指针变量 <变量名> = $<变量名>
		8. *指针变量，取出指针对应的值，使用new函数创建的对象，返回的是变量地址
	*/
	a, _ := GetData()
	_, b := GetData()
	fmt.Println(a, b)
}

func GetData() (int, int) {
	return 100, 200
}
