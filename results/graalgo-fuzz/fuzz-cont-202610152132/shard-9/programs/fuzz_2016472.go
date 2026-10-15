package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
func genFn0(p0 string, p1 uint8) int16 {
    p1 = p1
    switch 0 {
    case 127:
        p1 -= p1
    default:
        v6 := []int{9223372036854775807, 128}
        _ = v6
    }
    _ = p0
    _ = p1
    return 1000
}
func genFn1(p0 int64, p1 uint64) int {
    var v7 int32 = 128
    v8 := []int{200, 128, 9223372036854775807, 0}
    var v9 float32 = -2.25
    v10 := GenS1{F0: (p1 + (p1 / ((0) | 1))), F1: 0}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    var v1 string = "hello"
    v2 := make(map[int]string)
    v3 := make(map[uint32]bool)
    var v4 int = -894
    var v5 uint64 = 42
    if false {
        if (1 <= (v5 + (v5 % ((v5) | 1)))) {
            v11 := GenS1{F0: (^(v5)), F1: 2}
            _ = v11
        }
        var v12 int32 = 191
        _ = v12
    } else {
        var v13 int64 = 42
        _ = v13
    }
    var v14 int = (^(v4))
    var v15 float64 = 3.14159
    var v16 float32 = -2.25
    v17 := GenS1{F0: v5, F1: 1}
    v5 = (v5 << ((65535) & 63))
    var v18 int16 = 1000
    switch (v14 / (((v4 % (((v4 << ((0) & 63))) | 1))) | 1)) {
    case 685:
        var v19 uint = 1
        _ = v19
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        var v21 float64 = v15
        v1 = "\n\t"
        _ = fi20
        _ = v21
    }
    v4 -= v4
    fmt.Println(v18)
    var v22 uint64 = (v5 - (v5 ^ 42))
    var v23 uint8 = 127
    var v24 int16 = (v18 | -1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
