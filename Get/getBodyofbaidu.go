package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main(){
    client := &http.Client{}
    request,err := http.NewRequest("GET","http://www.baidu.com",nil)
    if err != nil {
        fmt.Println(err)
    }
    response,err := client.Do(request)
    res,err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(res))
    fmt.Println("Hello, World!")
}

