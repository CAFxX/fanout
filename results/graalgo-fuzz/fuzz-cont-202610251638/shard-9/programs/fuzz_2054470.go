package main

import (
    "fmt"
)


func main() {
    var v1 uint = 255
    var v2 int64 = 0
    v3 := []int{4294967295, -1, 32767}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v4 bool = ((v1 ^ v1) != (v1 + (v1 & v1)))
    var v5 int = -1
    fmt.Println(v5, v2, v5)
    var v6 uint16 = 32767
    v2 = v2
    var v7 int8 = -128
    var v8 int32 = 1
    var v9 float32 = (1e10 + 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
