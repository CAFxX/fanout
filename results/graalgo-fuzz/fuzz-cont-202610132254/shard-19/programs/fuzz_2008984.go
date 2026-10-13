package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
func genFn0() uint64 {
    v5 := []int{}
    switch 0 {
    case 874:
        var v6 int64 = 994
        _ = v6
    default:
        var v7 int = -2147483648
        _ = v7
    }
    var v8 uint64 = 782
    var v9 int8 = 1
    var v10 int64 = 129
    _ = v10
    _ = v5
    _ = v8
    _ = v9
    return v8
}
func genFn1(p0 uint8, p1 uint8) uint64 {
    switch "hello" {
    case "hello":
        switch 42 {
        case 127:
            v11 := []int{362}
            _ = v11
        default:
            p0 = (p0 - 42)
        }
    case "\n\t":
        p0 -= (p0 << ((255) & 63))
    }
    var v12 int = 1
    v13 := GenS1{F0: 18446744073709551615, F1: int8(p0)}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 2
}

func main() {
    var v1 float32 = (1.5 * (1.5 / ((1e10 - 1.5) - (1e10 / (((-2.25 - 0.1) + 3.14159) + 0.1)))))
    v2 := make(map[int]int)
    var v3 int = 42
    v4 := make(map[uint32]string)
    v1 -= 1.5
    switch "hello" {
    case "\n\t":
        fmt.Println(v1)
    case "":
        fmt.Println(v3, v3)
    }
    var v14 int8 = -69
    var v15 int16 = -1
    var v16 int = (v3 + (v3 * v3))
    var v17 uint64 = 2
    var v18 int = (v3 ^ (v3 << ((0) & 63)))
    v19 := []int{(v16 << ((1) & 63)), int(v17)}
    v1 = v1
    v14 -= (v14 % ((v14) | 1))
    for fi20 := 0; fi20 < 3; fi20++ {
        v16 -= (v18 << ((255) & 63))
        v14 -= (v14 % (((v14 % ((v14) | 1))) | 1))
        v18 += (v16 / ((32767) | 1))
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
