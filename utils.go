package utils

import (
    "fmt"
)


func log(a ...interface{}) {
    fmt.Println(a)
}

func notNil(a ...interface{}) {
    return a != nil
}
