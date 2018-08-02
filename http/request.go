package http

import (
    "strings"
    "fmt"
)

// 定义 Request, 相当于类
// 方法和属性的首字母都大写, 表明作用域为 Public
type Request struct {
    // 原始值
    Raw     string
    //
    Method  string
    Path    string
    // dict
    Query   map[string]string
    Body    string
    Headers map[string]string
    Cookies map[string]string
}

// 给 *Request 挂方法, 而不是 Request 
func (r *Request) Init(raw string) {
    r.Raw = raw
    // first line
    line := strings.Split(raw, "\r\n")[0]
    // FIXME: go 里面好像没有解构赋值
    es := strings.Split(line, " ")
    method, path := es[0], es[1]
    r.Method = method
    r.Body = strings.Split(raw, "\r\n\r\n")[1]
    // map 为字典(go 里叫动态增长的关联数组), map[keyType]ValueType
    // make map 的时候通常省略 size, size 会随着数据量的增长动态扩容
    r.Headers = make(map[string]string)
    r.Query = make(map[string]string)
    // 调用挂载的其他方法
    r.AddCookies()
    r.AddHeaders()
    r.ParseQuery(path)
}

func (r *Request) AddCookies() {
    // 挂载的方法可以取到 r, 类似于 this
    // TODO: key 不存在时会怎样?
    // 由于定义的 map 为 map[string]string, 所以返回了空字符串
    e := r.Headers["Cookie"]
    // test := r.Headers["hahaha"]
    // fmt.Println("debug test", "(", test, ")")
    fmt.Println("debug cookies", e)
    kvs := strings.Split(e, ": ")
    for _, s := range kvs {
        if strings.Contains(s, "=") {
            kv := strings.Split(s, "=")
            k, v := kv[0], kv[1]
            r.Cookies[k] = v
        }
    }
}

func (r *Request) AddHeaders() {
    raw := r.Raw
    s := strings.Split(raw, "\r\n\r\n")[0]
    hs := strings.Split(s, "\r\n")[1:]
    for _, s := range hs {
        kv := strings.Split(s, ": ")
        k, v := kv[0], kv[1]
        r.Headers[k] = v
    }
}

func (r *Request) ParseQuery(query string) {
    if strings.Contains(query, "?") {
        es := strings.Split(query, "?")
        r.Path = es[0]
        q := es[1]
        ms := strings.Split(q, "&")
        for _, m := range ms {
            kv := strings.Split(m, "=")
            k, v := kv[0], kv[1]
            r.Query[k] = v
        }

    } else {

        r.Path = query
    }
}
