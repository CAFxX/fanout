package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int64 }
func genFn0(p0 bool) int32 {
    p0 = (!(466 < 42))
    var v4 int32 = 1
    _ = p0
    _ = v4
    return v4
}
func genFn1(p0 string) int8 {
    p0 = p0
    v5 := GenS1{F0: 1, F1: 0}
    _ = p0
    _ = v5
    return 42
}

func main() {
    var v1 int32 = 42
    v2 := []int{}
    var v3 bool = false
    var v6 int32 = 0
    if false {
        v2 = append(v2, 4294967295)
    }
    _ = len(v2)
    if len(v2) > 0 {
        v2[0] = 572
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 9223372036854775807 {
    case -297:
        if false {
            v6 += v6
            v3 = v3
        }
    case 128:
        v7 := GenS1{F0: 1, F1: 196}
        _ = v7
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int8 = 36
        var v10 int64 = 1
        if (!(((-2.25 * 1e10) + (3.14159 / 1e10)) != (float64(1e10) * (((0.1 / 0.1) / float64((0.1 - (0.1 + float32(1e10))))) - 1e10)))) {
            v3 = ((^(v6)) == 42)
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: -9223372036854775808}
    if false {
        _ = len(v2)
        _ = len(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v6
}
