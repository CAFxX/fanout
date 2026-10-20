package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
func genFn0(p0 uint, p1 int32, p2 uint16) uint8 {
    v5 := []int{122, 921}
    var v6 bool = false
    fmt.Println(p0, v6, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 164
}
func genFn1(p0 bool, p1 uint16) int16 {
    p1 = (p1 * 42)
    fmt.Println(p1, p1)
    p1 += (^(p1))
    _ = p0
    _ = p1
    return int16(p1)
}

func main() {
    v1 := []int{-509, -1}
    var v2 uint = 962
    v3 := []int{80, 499, -9223372036854775808}
    var v4 int16 = 1000
    v2 = v2
    v2 -= 4294967295
    v7 := GenS1{F0: v4, F1: 229}
    if len(v3) > 0 {
        v3[0] = 32767
    }
    var v8 int16 = v4
    fmt.Println(v2, v8)
    if (118 > 100) {
        var v9 string = "\n\t"
        switch -727 {
        case -303:
            fmt.Println(v2)
        default:
            v10 := GenS1{F0: genFn1((true || (2 > 124)), 1), F1: -223}
            _ = v10
        }
        var v11 float32 = -2.25
        _ = v11
        _ = v9
    } else {
        fmt.Println(v4, v4, v4)
        var v12 float64 = (1.5 + 0.1)
        fmt.Println(v2, v8)
        _ = v12
    }
    v4 += v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
