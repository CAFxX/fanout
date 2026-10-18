package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 468
    var v5 int16 = -32768
    var v6 uint64 = 1
    v4 += (v4 & (v4 + (v4 >> ((4294967295) & 63))))
    switch "hello" {
    case "":
        _ = len(v2)
    case "\n\t":
        var v7 string = "a"
        _ = v7
    case "hello":
        var v8 int8 = -128
        _ = v8
    default:
        v5 += v5
    }
    v5 = (^(v5))
    fmt.Println(v6, v6, v5)
    var v9 float64 = 0.0
    var v10 bool = (-2147483648 == 2147483647)
    v11 := v1["\n\t"]
    fmt.Println(v6, v4, v10)
    if true {
        fmt.Println(v9)
        var v12 uint64 = 0
        _ = v12
    } else {
        fmt.Println(v10, v4, v10)
    }
    var v13 float64 = v9
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
