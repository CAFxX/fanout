package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint) int {
    p0 -= p0
    var v7 uint64 = 457
    _ = p0
    _ = v7
    return 395
}
func genFn1(p0 int32, p1 uint64) int8 {
    if false {
        p0 -= 399
    }
    p0 -= (^(p0))
    _ = p0
    _ = p1
    return -1
}

func main() {
    v1 := []uint32{2147483647, 313}
    v2 := []int{-9223372036854775808, 127}
    var v3 bool = true
    var v4 uint32 = 65535
    var v5 uint64 = 0
    var v6 int64 = -722
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v8 := GenS1{F0: -54, F1: 811}
    var v9 int = 4294967295
    v5 += 177
    if len(v2) > 0 {
        v2[0] = 9223372036854775807
    }
    var v10 float64 = -2.25
    fmt.Println(v9)
    if (!(false || (!true))) {
        fmt.Println(v3)
        fmt.Println(v4, v5, v9)
        var v11 float32 = 1e10
        _ = v11
    }
    v12 := GenS2{F0: 1, F1: 255}
    if len(v2) > 0 {
        v2[0] = genFn0(4294967295)
    }
    var v13 uint64 = v5
    if true {
        v9 += (^(v9))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
