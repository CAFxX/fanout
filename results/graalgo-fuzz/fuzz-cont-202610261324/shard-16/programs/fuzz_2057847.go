package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint8 }
func genFn0() int16 {
    var v4 string = ""
    v5 := []int{-129, -2147483648}
    _ = v4
    _ = v5
    return 10
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 2147483647
    var v3 int64 = -9223372036854775808
    v6 := []int{}
    v6 = append(v6, 4294967295)
    var v7 int16 = genFn0()
    var v8 float32 = 0.1
    if ((v7 ^ v7) <= 1) {
        _ = len(v1)
        fmt.Println(v7, v3, v8)
        v8 -= v8
    } else {
        var v9 int = -602
        fmt.Println(v2, v8, v8)
        v10 := GenS1{F0: 152, F1: (42 != 2), F2: 1}
        _ = v10
        _ = v9
    }
    var v11 int = 982
    v12 := GenS1{F0: 0, F1: (true || (1 > 1)), F2: 42}
    v1[1] = 0
    fmt.Println(v8)
    v3 = v3
    v3 += (v3 - v3)
    v13 := GenS1{F0: 190, F1: true, F2: 215}
    var v14 float64 = 3.14159
    v15 := GenS1{F0: 117, F1: true, F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
