package main

import (
    "fmt"
    "net/http"
)

func main() {
    // 注册处理函数
    http.HandleFunc("/", helloHandler)

    // 启动服务器在 8080 端口
    fmt.Println("Server starting on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Server error: %v\n", err)
    }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
}