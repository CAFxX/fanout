package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 bool; F1 int32 }
func genFn0(p0 uint8, p1 bool, p2 int8) string {
    p1 = (323 != 106)
    var v7 float32 = (float32(-2.25) + (1e10 - 3.14159))
    v7 = 1e10
    switch 934 {
    case 32767:
        v8 := GenS1{F0: 18446744073709551615, F1: p2}
        _ = v8
    case 2147483647:
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 uint16 = 42
            for fi11 := 0; fi11 < 10; fi11++ {
                var v12 int32 = -199
                _ = fi11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
    case 0:
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := []int{(fi13 & (fi13 * 32767)), 127}
            _ = fi13
            _ = v14
        }
    }
    if p1 {
        var v15 int8 = p2
        v7 = (float32(1.5) / (float32(1e10) / 1.0))
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ""
}
func genFn1(p0 uint) int {
    v16 := GenS2{F0: false, F1: 42}
    v17 := []int{90, 0, 4294967295, 1}
    if ((0.0 / 1.0) < (0.1 / 0.1)) {
        p0 += p0
        fmt.Println(p0)
        v18 := GenS1{F0: p0, F1: 1}
        _ = v18
    }
    p0 = p0
    if (true || (0 != 42)) {
        p0 -= (p0 + p0)
    } else {
        if false {
            fmt.Println(p0)
        }
    }
    _ = p0
    _ = v16
    _ = v17
    return 127
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float32 = 1e10
    var v3 bool = (-1 >= 2147483647)
    v4 := []int{2147483647, -1, -1}
    v5 := make(map[string]string)
    v6 := []string{""}
    for fi19 := 0; fi19 < 3; fi19++ {
        v1 = v1
        var v20 int32 = 2147483647
        _ = fi19
        _ = v20
    }
    v1 = v1
    var v21 int16 = -1
    if len(v4) > 0 {
        v4[0] = 379
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 uint64 = 144
        v21 = 190
        var v24 float64 = 0.0
        _ = fi22
        _ = v23
        _ = v24
    }
    v2 = (v2 - 0.0)
    v25 := GenS2{F0: (0 != 0), F1: 76}
    v2 -= (v2 + ((1.5 - v2) * -2.25))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
