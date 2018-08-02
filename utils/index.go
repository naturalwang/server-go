package utils

import (
    "fmt"
)

/*
Go 语言使用名称首字母大小写来判断一个对象（全局变量、全局常量、类型、结构字段、函数、方法）的访问权限，对于包而言同样如此。包中成员名称首字母大小写决定了该成员的访问权限。首字母大写，可被包外访问，即为 public（公开的）；首字母小写，则仅包内成员可以访问，即为 internal（内部的）
 */

func Log(a ...interface{}) {
    fmt.Println(a)
}

// ------- 错误, Why? NotNil 函数错误: invalid memory address or nil pointer dereference
func NotNil(a ...interface{}) bool {
    return a != nil
}
