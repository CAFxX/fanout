package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int; F1 int8; F2 uint16 }
type GenS3 struct { F0 int }

func main() {
    var v1 uint = 65535
    v2 := make(map[string]int)
    v3 := make(map[string]bool)
    fmt.Println(v1)
    var v4 int8 = -85
    var v5 int64 = 0
    var v6 int32 = 450
    if true {
        v4 = v4
        v7 := GenS1{F0: v6, F1: 0, F2: 1}
        _ = v7
    } else {
        v1 = v1
        v8 := GenS1{F0: -2147483648, F1: 65535, F2: 2}
        var v9 int64 = -9223372036854775808
        _ = v8
        _ = v9
    }
    fmt.Println(v6, v1, v1)
    switch -135 {
    case 0:
        var v10 bool = false
        _ = v10
    case 32768:
        fmt.Println(v4, v6)
    }
    var v11 uint16 = 255
    if (v6 >= (v6 << (((v1 << ((v1) & 63))) & 63))) {
        v12 := GenS2{F0: 2147483647, F1: v4, F2: (v11 & (v11 % ((v11) | 1)))}
        v13 := GenS3{F0: 127}
        _ = v12
        _ = v13
    }
    v2["hello"] = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
