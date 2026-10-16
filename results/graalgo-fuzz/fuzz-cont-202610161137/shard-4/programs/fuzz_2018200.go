package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 bool }

func main() {
    var v1 int8 = -128
    var v2 uint = 1
    v3 := []int{-9223372036854775808, 2147483647, 42}
    v4 := make(map[int]int)
    v5 := []int{128, 1, 815}
    var v6 float64 = 1.5
    v2 = v2
    switch -1 {
    case 9223372036854775807:
        v7 := GenS1{F0: "hello", F1: 42}
        _ = v7
    default:
        switch 0 {
        case 4294967295:
            v6 -= 0.0
        case -714:
            var v8 int8 = v1
            _ = v8
        }
    }
    var v9 float32 = 1.5
    var v10 uint16 = 255
    switch -2147483648 {
    case -9223372036854775808:
        var v11 int = 127
        _ = v11
    case -746:
        v2 += (v2 ^ (^(v2)))
    case 42:
        var v12 int64 = 880
        _ = v12
    }
    fmt.Println(v1)
    if ((v6 < (3.14159 / ((v6 + 0.0) / -2.25))) && (-858 >= -32768)) {
        var v13 uint32 = 507
        _ = v13
    }
    var v14 int = 278
    if (42 >= v2) {
        fmt.Println(v14, v6, v1)
    } else {
        var v15 string = ("a" + "a")
        _ = v15
    }
    _ = len(v4)
    delete(v4, (v14 % ((v14) | 1)))
    v10 = v10
    v16 := GenS2{F0: true}
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
