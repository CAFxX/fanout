package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint16 }

func main() {
    var v1 uint = 0
    var v2 uint16 = 0
    var v3 int32 = 42
    v4 := make(map[uint32]uint32)
    v1 = 1
    if true {
        v5 := []int{4294967295}
        _ = v5
    } else {
        var v6 int16 = 1000
        v7 := GenS1{F0: v2, F1: v1, F2: v2}
        _ = v6
        _ = v7
    }
    fmt.Println(v2, v1)
    _ = len(v4)
    var v8 bool = false
    v9 := v4[978]
    var v10 bool = (1 != 787)
    var v11 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
