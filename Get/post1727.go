package main

import (
    "fmt"
    "net/http"
    "strings"
    "io/ioutil"
)

func main(){
    resp,err := http.Post("http://www.baidu.com","application/x-www-form-urlencoded",strings.NewReader("user=admin&pass=admin"))
    if err != nil{
        fmt.Println(err)
    }

    defer resp.Body.Close()
    body,err := ioutil.ReadAll(resp.Body)
    if err != nil{
        fmt.Println(err)
    }
    fmt.Println(string(body))
    fmt.Println("Hello, World!")
}

