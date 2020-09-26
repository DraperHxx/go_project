package main

import (
    "fmt"
    "net/http"
    "strconv"
)
func main(){
    client := &http.Client{}
    request,err := http.NewRequest("GET","http://www.baidu.com",nil)
    if err != nil {
        fmt.Println(err)
    }
    //使用http.Cookie初始化一个cookie键值对
    cookie := &http.Cookie{Name:"userId",Value:strconv.Itoa(123456)}

    //使用前面构建的request方法AddCookie往请求中添加Cookie
    request.AddCookie(cookie)
    request.AddCookie(&http.Cookie{Name:"session",Value:"YWRtaW4="})
    respones,err := client.Do(request)
    fmt.Println(respones.Request.Cookie)

    fmt.Println("Hello, World!")
}

