package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 float32 = 0.1
    var v2 bool = false
    v3 := make(map[uint32]int)
    v4 := make(map[int]bool)
    v5 := []int{9223372036854775807, -426, 2147483647}
    switch 286 {
    case 0:
        fmt.Println(v1, v1)
    case 663:
        v6 := GenS1{F0: 677}
        _ = v6
    default:
        if v2 {
            var v7 float32 = (float32(0.1) + v1)
            var v8 uint16 = 130
            _ = v7
            _ = v8
        }
    }
    var v9 uint64 = 887
    v3[2147483647] = -1
    v9 = (v9 / ((v9) | 1))
    v10 := GenS2{F0: ("a" + ("hello" + "hello")), F1: 70}
    if v2 {
        if v2 {
            fmt.Println(v9)
        } else {
            v2 = (2 < 2)
        }
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v1 -= v1
    }
    v2 = (255 != 42)
    v5 = append(v5, -9223372036854775808)
    v11 := GenS1{F0: v9}
    v12 := GenS1{F0: (^(v9))}
    v13 := GenS3{F0: 145}
    var v14 bool = true
    for fi15 := 0; fi15 < 1; fi15++ {
        v16 := GenS3{F0: 754}
        _ = fi15
        _ = v16
    }
    v14 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
