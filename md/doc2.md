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
5. 指针变量
    1. 指针变量 <变量名> = $<变量名>
        ```go
            var5 := "test"
            var6 := $var5
        ```
    2. new函数创建的对象，返回的是变量地址
        ```go
            var5 := new(int)
        ```
   3. *指针变量，取出指针对应的值，使用new函数创建的对象，返回的是变量地址
       ```go
           var5 := new(int)
           var6 := *var5
       ```
6. 匿名变量
    1. 不会占用内存，不需要思考命名，多次声明不会出现问题，一般用于方法多返回值时占位，以便于取到真正要的值
       ```go
           func GetData() (int, int) {
               return 100, 200
           }
           func main(){
               a, _ := GetData()
               _, b := GetData()
               fmt.Println(a, b)
           }
       ```