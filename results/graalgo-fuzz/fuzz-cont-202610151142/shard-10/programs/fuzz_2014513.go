package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
func genFn0(p0 string, p1 uint8) int64 {
    fmt.Println(p1, p1)
    v6 := GenS1{F0: 4294967295, F1: 0}
    var v7 uint64 = 42
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 9223372036854775807
}

func main() {
    v1 := []int{127, -566}
    v2 := []uint32{637}
    var v3 float64 = -2.25
    var v4 int64 = 42
    v5 := make(map[int]string)
    fmt.Println(v4, v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch ("" + "") {
    case "\n\t":
        switch 128 {
        case -679:
            v8 := v5[1]
            _ = v8
        case 127:
            v4 -= 42
        case -9223372036854775808:
            v4 += v4
        }
    }
    if len(v1) > 0 {
        v1[0] = 134
    }
    v3 = 0.1
    v4 -= (-(v4))
    fmt.Println(v4, v4)
    v9 := GenS1{F0: 42, F1: 7}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
