package main

import (
    "fmt"
    "net/http"
)

func main(){
    client := &http.Client{}
    request,err := http.NewRequest("GET","http://www.baidu.com",nil)
    if err != nil{
        fmt.Println(err)
    }
    response,err := client.Do(request)
    fmt.Println(response.StatusCode)
    fmt.Println("Hello, World!")
}

