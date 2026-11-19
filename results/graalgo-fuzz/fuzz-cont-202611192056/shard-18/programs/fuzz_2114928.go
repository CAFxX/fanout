package main

import (
    "fmt"
)


func main() {
    var v1 uint32 = 0
    v2 := make(map[string]bool)
    var v3 int = 2147483647
    v4 := v2["\n\t"]
    var v5 int8 = 100
    v4 = false
    fmt.Println(v3, v4)
    fmt.Println(v3)
    v4 = false
    var v6 string = ("foo bar" + "a")
    var v7 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
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
