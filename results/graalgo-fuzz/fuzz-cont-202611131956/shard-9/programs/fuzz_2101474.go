package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 string; F1 int64; F2 int8 }
func genFn0(p0 uint8, p1 uint, p2 uint64) int {
    v7 := GenS1{F0: -1, F1: (p1 & uint(p0))}
    v8 := []int{70, 32767, 120}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]uint32)
    v3 := []string{"hello", "", "a"}
    v4 := make(map[uint32]uint32)
    var v5 float32 = 0.1
    var v6 uint8 = 127
    switch 868 {
    case 4294967295:
        fmt.Println(v5, v6)
    case 964:
        _ = len(v3)
    case 32768:
        v9 := GenS1{F0: -1, F1: 46}
        _ = v9
    }
    v10 := GenS1{F0: 9223372036854775807, F1: 4294967295}
    v5 += (v5 + (v5 - v5))
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 uint8 = 133
        var v13 int64 = -9223372036854775808
        var v14 int32 = 86
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        fmt.Println(fi15, v5)
        _ = fi15
    }
    _ = len(v2)
    switch genFn0((v6 + v6), 42, 0) {
    case 128:
        var v16 uint = 470
        _ = v16
    case -1:
        v17 := GenS2{F0: "foo bar", F1: 42, F2: 127}
        _ = v17
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 uint64 = 0
        v5 -= v5
        _ = fi18
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
