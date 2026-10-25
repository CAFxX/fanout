package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint) int {
    fmt.Println(p0)
    v7 := []int{599, 128, -462}
    p0 = p0
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 uint16 = uint16(p0)
        _ = fi8
        _ = v9
    }
    var v10 string = "\n\t"
    _ = p0
    _ = v10
    _ = v7
    return int(p0)
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{-2147483648}
    var v3 bool = (0.1 <= 0.0)
    var v4 int64 = -334
    var v5 int64 = (-(v4))
    var v6 int16 = -486
    switch "\n\t" {
    case "":
        fmt.Println(v4, v6)
    case "foo bar":
        _ = len(v1)
    case "hello":
        if (1e10 >= 0.0) {
            v1[-2147483648] = false
        } else {
            fmt.Println(v4)
        }
    default:
        var v11 int = -9223372036854775808
        _ = v11
    }
    v6 = v6
    v12 := GenS1{F0: 879, F1: -2147483648}
    v3 = v3
    for fi13 := 0; fi13 < 14; fi13++ {
        _ = len(v2)
        _ = fi13
    }
    fmt.Println(v6, v6)
    var v14 uint64 = 2
    var v15 bool = v3
    var v16 string = ("a" + "a")
    v5 = (v5 ^ -390)
    switch 0 {
    case -9223372036854775808:
        delete(v1, genFn0(0))
    }
    var v17 string = v16
    v18 := GenS2{F0: 2, F1: v6}
    v19 := GenS2{F0: 65535, F1: (v6 + v6)}
    var v20 int8 = -128
    v21 := v1[-2147483648]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
