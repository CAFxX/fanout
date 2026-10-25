package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint }
func genFn0(p0 int64, p1 uint16) bool {
    fmt.Println(p0, p0)
    v6 := []int{128, 128}
    p1 -= (p1 >> ((1) & 63))
    var v7 uint16 = p1
    var v8 int16 = 0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return true
}
func genFn1() uint8 {
    v9 := []int{322, 1}
    var v10 int64 = 1
    v10 = -536
    _ = v10
    _ = v9
    return 2
}

func main() {
    var v1 int = 128
    var v2 uint64 = 677
    var v3 bool = false
    var v4 float32 = -2.25
    var v5 uint16 = 255
    v2 -= v2
    fmt.Println(v5)
    v11 := []int{}
    if len(v11) > 0 {
        v11[0] = (v1 % ((v1) | 1))
    }
    if v3 {
        v12 := GenS1{F0: v3}
        _ = v12
    }
    var v13 uint64 = (^(v2))
    fmt.Println(v1, v1, v5)
    if (v3 && v3) {
        if ((v1 ^ 0) != 331) {
            v14 := GenS2{F0: 2147483648, F1: (v2 | v13), F2: 971}
            v15 := GenS1{F0: (!v3)}
            _ = v14
            _ = v15
        }
    }
    v1 = (v1 ^ (v1 / (((v1 / ((-1) | 1))) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
