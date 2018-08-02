package route

import (
    . "../http"
)

func image(r Request) []byte {
    p := r.Query["path"]
    m := ResponseFile(p)
    return m
}

func index(r Request) []byte {
    name := "index.html"
    m := ResponseFile(name)
    return m
}

func doge(r Request) []byte {
    name := "doge.html"
    m := ResponseFile(name)
    return m
}

// 这是一个 map, value 是一个函数
var RouteIndex = map[string]func(Request) []byte{
    "/doge": doge,
    "/":     index,
    "/img":  image,
}
