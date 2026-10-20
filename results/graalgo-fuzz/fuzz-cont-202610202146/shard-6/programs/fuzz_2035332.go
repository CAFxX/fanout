package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]int)
    v2 := []string{(("hello" + ("hello" + "\n\t")) + "\n\t"), ""}
    v3 := []string{"hello", ""}
    var v4 int64 = 0
    v4 -= (^(v4))
    var v5 int16 = 1000
    fmt.Println(v4, v5)
    var v6 bool = (true || false)
    v6 = (2147483647 == -208)
    fmt.Println(v6)
    fmt.Println(v6, v5, v6)
    fmt.Println(v6, v6, v6)
    fmt.Println(v5, v4)
    var v7 int8 = 100
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
