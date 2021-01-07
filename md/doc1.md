[目录](../README.md)
1. 一行代码声明
    1. var <变量名> <类型>
       ```go
          var var2 string
       ```
    2. var <变量名> <类型> = <值>
       ```go
          var var2 string = "string2"
       ```
    3. var <变量名> = <值>，根据值进行推断类型
        ```go
           var var2 = "string2"
        ```
2. 多个变量一起声明
    1. var (
           <变量名> <类型>
           ...
       )
    ```go
       var (
       	var3 string
            var4 string
       )
    ```
3. 声明和初始化一个变量
    1. 使用 := （推导声明写法或者短类型声明法：编译器会自动根据右值类型推断出左值的对应类型。），可以声明一个变量，并对其进行（显式）初始化。
    ```go
        var5 := "test"
    ```
4. 声明和初始化多个变量
    1. <变量名>... := <值>...
    ```go
        var5, var6 := "test", "test1"
    ```