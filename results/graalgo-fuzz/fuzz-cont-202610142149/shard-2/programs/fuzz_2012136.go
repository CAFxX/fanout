package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint }
type GenS2 struct { F0 int64; F1 int; F2 int8 }
func genFn0(p0 uint32, p1 uint16, p2 uint) uint32 {
    p1 += 306
    var v4 float64 = 3.14159
    p0 -= (p0 / ((p0) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p0
}
func genFn1(p0 string, p1 uint16, p2 uint8) int16 {
    v5 := []int{1, 319, -1}
    if (1 > 762) {
        v5 = append(v5, -2147483648)
        fmt.Println(p0, p2, p2)
        var v6 uint = 42
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 32767
}

func main() {
    v1 := []uint32{1, 258, 1, 163}
    var v2 float64 = 0.0
    var v3 int16 = 382
    var v7 int32 = -1
    if true {
        _ = len(v1)
        fmt.Println(v7, v3)
        v8 := GenS1{F0: 700, F1: 182, F2: 761}
        _ = v8
    }
    if true {
        if true {
            var v9 int16 = v3
            _ = v9
        }
    }
    if len(v1) > 0 {
        v1[0] = genFn0(2147483647, 65535, 32)
    }
    v7 = (^(v7))
    var v10 uint = 4294967295
    v2 += v2
    var v11 int = int(v7)
    fmt.Println(v11, v7, v2)
    v12 := GenS2{F0: 42, F1: (v11 * v11), F2: 7}
    fmt.Println(v7, v11)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v7
}
