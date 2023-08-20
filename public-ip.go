package main

import (
        "github.com/gin-gonic/gin"
        "net/http"
        "fmt"
        "io/ioutil"
)
func main() {
        r := gin.Default()

        r.GET("/", func(c *gin.Context) {
                c.String(http.StatusOK, "hello world1")
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

        })

        r.Run()
        }
