package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 bool }
type GenS2 struct { F0 int16; F1 int64; F2 int64 }
func genFn0(p0 int64, p1 uint, p2 uint64) uint {
    var v5 bool = (67 >= 255)
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p1
}
func genFn1(p0 int32, p1 uint64) int {
    p0 = (-(p0))
    fmt.Println(p0, p1)
    v6 := GenS1{F0: 272, F1: 42, F2: true}
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := []int{703}
        _ = len(v8)
        _ = fi7
        _ = v8
    }
    if (true && true) {
        for fi9 := 0; fi9 < 10; fi9++ {
            fmt.Println(p1)
            p1 -= 2
            _ = fi9
        }
        for fi10 := 0; fi10 < 15; fi10++ {
            fmt.Println(p0, fi10)
            if false {
                fmt.Println(p0, fi10, p0)
            } else {
                v11 := GenS2{F0: 32767, F1: 1, F2: 1}
                fmt.Println(p0)
                p0 -= (p0 / ((p0) | 1))
                _ = v11
            }
            p1 -= (^(p1))
            _ = fi10
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return -1
}

func main() {
    var v1 int8 = 1
    var v2 string = "a"
    var v3 int64 = 0
    v4 := make(map[int]string)
    v3 = 42
    if (false || (1 > 1)) {
        var v12 uint16 = 1
        _ = v12
    }
    fmt.Println(v2)
    switch "" {
    case "\n\t":
        var v13 uint8 = 0
        _ = v13
    default:
        v14 := v4[42]
        _ = v14
    }
    fmt.Println(v3, v3)
    fmt.Println(v2, v2, v2)
    v1 -= 7
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
