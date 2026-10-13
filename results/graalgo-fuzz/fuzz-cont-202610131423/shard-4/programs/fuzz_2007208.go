package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int }

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 1e10
    var v3 string = "hello"
    var v4 uint8 = 128
    var v5 int32 = 0
    var v6 uint64 = 2
    var v7 int64 = 1
    v4 += v4
    var v8 uint = 1
    var v9 uint16 = 2
    fmt.Println(v5)
    var v10 bool = true
    switch -1 {
    case -55:
        v6 -= (v6 * v6)
    case 4294967295:
        var v11 string = v3
        _ = v11
    default:
        v12 := GenS1{F0: (v3 == "\n\t"), F1: 0, F2: 42}
        _ = v12
    }
    v10 = false
    var v13 int8 = 127
    var v14 int64 = (v7 & (v7 + (v7 + v7)))
    v15 := GenS1{F0: (-486 >= 42), F1: 2147483647, F2: -127}
    fmt.Println(v4)
    var v16 string = "hello"
    v6 = (v6 | 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
