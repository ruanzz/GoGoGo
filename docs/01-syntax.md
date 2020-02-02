## 变量定义
 1. 使用关键字var定义变量 var variable_name type
 比如 var name string
 
 2. 类型推断
  比如 var a,b = 1,"11"
  或者 
  ```bash
  var(
      a = 1
      b = "11"
  )
  ```
  
 3. := 定义变量,只能在函数内定义
   比如 lang := "Go"
   
## 变量类型
 - bool string 
 - (u)int (u)int8 (u)int16 (u)int32 (u)int64 uintptr
 - byte rune
 - float32 float64 complex64 complex128
 
## 强制类型转换
 Go语言中类型转换是强制的，没有隐性类型转换，直接转换比如 float64(3)
 
## 常量
  使用const关键字类定义 比如 const a = 1
  
  借助const来定义枚举类

## 分支
  condition表达式没有小括号
  
## 循环
  同样没有小括号

## 指针
  指针不能运算

## range
  范围
  
## make
  生成对象
  
## 数组
  []在类型前面，比如var arr [5]int,取数组长度len(arr),容量cap(arr),值得注意的是数组在go中是值类型
  
## 切片
  是array的一个view,更改slice会同步修改原数组,添加元素append(slice,value)
 
## Map
 Go语言中提供了Map这种类型，其他的语言都是通过语言写出来的一种数据结构，定义格式: map[K]V
 比如：
 ```gos
  m := map[string]string{
      		"key1":"value1",
      		"key2":"value2",
      	}
 ```
 
   
