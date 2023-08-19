package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    response, err := http.Get("https://httpbin.org/ip")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Public IP:", string(body))
}
