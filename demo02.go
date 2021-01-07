package main

import (
	"fmt"
)

func main() {

	// 变量声明
	// 1. var var1 string; var <变量名> <类型>
	// 2. var var2 string = "string2"; var <变量名> <类型> = <值>
	// 3. var var2 = "string2"; var <变量名> = <值>，根据值进行推断类型
	/*
		4. var (
			<变量名> <类型>
			...
		)
	*/
	/*
		5. <变量名> := <值>
	*/
	/*
		6. <变量名>... := <值>...
	*/
	var5, var6 := "test", "test1"
	fmt.Println(var5, var6)
}
