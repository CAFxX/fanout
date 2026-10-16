package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -3
    var v3 bool = false
    var v4 int8 = -128
    var v5 int8 = v4
    var v6 string = "hello"
    v7 := v1[(v6 + v6)]
    v8 := []int{0}
    v9 := GenS1{F0: 42}
    var v10 bool = (!(-521 <= -2147483648))
    fmt.Println(v3, v5, v7)
    var v11 int = 128
    if v3 {
        _ = len(v8)
        var v12 int8 = -1
        _ = v12
    }
    var v13 int8 = (v5 / ((v5) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
