package main

import (
    "fmt"
    "net"
    "os"
    . "./http"
    . "./route"
    . "./utils"
)

// todo, 将 value 函数转换为通用函数
type typeMapRoute map[string]func(Request) []byte

func addRoutes(m1 *typeMapRoute, m2 typeMapRoute) {
    for k, v := range m2 {
        (*m1)[k] = v
    }
}

func handleClient(conn net.Conn) {
    request := make([]byte, 1024)
    defer conn.Close()
    num, err := conn.Read(request)
    checkError(err)
    raw := string(request[:num])
    r := Request{}
    r.Init(raw)
    path := r.Path
    fmt.Println("path", path)
    s := responseForPath(path, r)
    conn.Write(s)
    request = make([]byte, 1024)
}

func responseForPath(path string, r Request) []byte {
    m := typeMapRoute{}
    addRoutes(&m, RouteIndex)
    fn, ok := m[path]
    var s []byte
    if !ok {
        s = ResponseError("404")
    } else {
        s = fn(r)
    }
    return s
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}

func main() {
    // 设置 host port
    host := "localhost"
    port := "3000"
    address := net.JoinHostPort(host, port)
    // 解析地址
    // ResolveTCPAddr 将 addr 作为 TCP 地址解析并返回
    // net 参数是 "tcp4"、"tcp6"、"tcp" 中的任意一个，分别表示 TCPv4、TCPv6 或者任意
    tcpAddress, err := net.ResolveTCPAddr("tcp", address)
    checkError(err)
    // 创建 tcpListener
    // ListenTCP 在本地 TCP 地址 laddr 上声明并返回一个 * TCPListener
    listener, err := net.ListenTCP("tcp", tcpAddress)
    checkError(err)
    // 无限循环监听端口
    for {
        conn, err := listener.Accept()
        // 不处理错误
        if notNil(err) {
            continue
        }
        go handleClient(conn)
    }
}
