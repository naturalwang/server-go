package main

import (
    "fmt"
    "net"
    "os"
    . "./http"
    . "./route"
    // . "./utils"
)

// TODO: 将 value 函数转换为通用函数
// value 目前为 参数是 Request 实例, 返回值为 []byte
type typeMapRoute map[string]func(Request) []byte

func addRoutes(m1 *typeMapRoute, m2 typeMapRoute) {
    for k, v := range m2 {
        (*m1)[k] = v
    }
}

func handleClient(conn net.Conn) {
    // 分配内存
    // make 可以为 slice map channel 分配内存并初始化
    // make 返回值类型为参数的 type
    // 而 new 方法返回 *type
    bufferSize := 1024
    request := make([]byte, bufferSize)
    // return 时做的三件事:
    // 1. 给返回值赋值
    // 2. 调用 defer 表达式
    // 3. 将返回值返回给调用函数
    defer conn.Close()
    // Read 方法从 Conn 中读取数据
    // TODO: 这里只读取了 1024 个字节
    num, err := conn.Read(request)
    checkError(err)
    // 类型强制转换
    raw := string(request[:num])
    fmt.Println("raw", raw)
    // 生成类实例
    r := Request{}
    r.Init(raw)
    path := r.Path
    fmt.Println("path", path)
    response := responseForPath(path, r)
    conn.Write(response)
}

func responseForPath(path string, r Request) []byte {
    // 生成类实例, 获取到的是引用
    m := typeMapRoute{}
    // 传递 map 类型要传递引用
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
        // Fprintf 将其余参数的内容写入第一个参数里(io.Writer)
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        // 退出程序, 0 代表成功, 其余为 error
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
        if err != nil {
            continue
        }
        // 创建新线程执行 handleClient 函数
        go handleClient(conn)
    }
}
