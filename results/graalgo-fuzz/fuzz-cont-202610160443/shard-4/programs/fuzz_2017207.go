package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint64; F1 uint }
func genFn0(p0 int64, p1 uint64) int8 {
    switch "a" {
    case "a":
        v7 := []int{0, 32768, 2147483647}
        _ = v7
    case "hello":
        p0 = p0
    default:
        p1 += (^(p1))
    }
    p0 = -1
    v8 := []int{32767, 9223372036854775807, -9223372036854775808, 0}
    for fi9 := 0; fi9 < 12; fi9++ {
        fmt.Println(fi9)
        var v10 uint8 = 0
        fmt.Println(fi9)
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v8
    return 127
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 776
    var v3 string = "\n\t"
    var v4 int8 = -1
    var v5 int16 = 458
    var v6 uint = 853
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 uint64 = (v2 | (v2 | 18446744073709551615))
        _ = fi11
        _ = v12
    }
    var v13 uint8 = 69
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 int32 = 2147483647
        if true {
            v16 := []int{(-(fi14))}
            _ = v16
        }
        _ = fi14
        _ = v15
    }
    var v17 uint8 = 2
    switch v3 {
    case "":
        v13 -= (v13 / ((v13) | 1))
    case "foo bar":
        var v18 float32 = (0.0 * 1.5)
        _ = v18
    }
    v19 := []int{788, 2147483647}
    if len(v19) > 0 {
        v19[0] = 1
    }
    var v20 int32 = 0
    v21 := GenS1{F0: 860, F1: false}
    v3 = "foo bar"
    fmt.Println(v3, v20)
    var v22 int = -1
    v17 = 42
    for fi23 := 0; fi23 < 12; fi23++ {
        v24 := GenS2{F0: v2, F1: (v6 * v6)}
        _ = fi23
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
