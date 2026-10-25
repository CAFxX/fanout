package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
func genFn0(p0 uint32) int8 {
    for fi5 := 0; fi5 < 7; fi5++ {
        fmt.Println(fi5, fi5)
        var v6 int64 = 42
        v7 := GenS1{F0: (^(v6)), F1: p0}
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        fmt.Println(fi8, p0)
        p0 += 1
        v9 := []int{fi8}
        _ = fi8
        _ = v9
    }
    v10 := []int{-1, 9223372036854775807, -2, 9223372036854775807}
    var v11 int = -1
    v12 := GenS1{F0: 0, F1: 68}
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return 22
}

func main() {
    var v1 uint16 = 32767
    v2 := []int{-1}
    var v3 uint64 = 2
    v4 := make(map[string]uint32)
    if true {
        fmt.Println(v1, v1)
    }
    var v13 string = "a"
    switch 2147483647 {
    case 859:
        var v14 int64 = -1
        _ = v14
    default:
        var v15 float64 = (((0.1 * 3.14159) - float64((1.5 * 0.0))) / ((1e10 - 0.0) * (-2.25 / -2.25)))
        _ = v15
    }
    v1 -= (v1 << ((941) & 63))
    v13 = ""
    fmt.Println(v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        var v16 float32 = (1.5 * 1e10)
        var v17 int = 32767
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: 0, F1: 564}
    if (!(32767 <= 691)) {
        v19 := v4[(("" + "a") + v13)]
        var v20 uint64 = 18446744073709551615
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
