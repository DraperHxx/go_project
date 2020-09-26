package main

import (
    "fmt"
    "net/http"
    "strings"
    "io/ioutil"
)

func main(){
    client := &http.Client{}
    body := ioutil.NopCloser(strings.NewReader("user=admin&pass=admin"))
    req,err := http.NewRequest("POST","http://www.baidu.com",body)
    if err != nil{
        fmt.Println(err)
    }
    req_body,err := ioutil.ReadAll(req.Body)
    fmt.Println(string(req_body))


    req.Header.Set("Content-Type","application/x-www-form-urlencoded")
    resp,err := client.Do(req)
    if err != nil{
        fmt.Println(err)
    }
    //使用Do方法后，在次打印Body已经被关闭，readall会出错
    req_body,err = ioutil.ReadAll(resp.Request.Body)
    if err != nil{
        fmt.Println(err)
    }
    fmt.Println(string(req_body))
    fmt.Println("Hello, World!")
}

