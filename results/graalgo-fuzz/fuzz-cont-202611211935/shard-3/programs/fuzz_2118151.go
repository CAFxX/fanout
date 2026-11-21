package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int8; F2 int16 }
func genFn0(p0 int64, p1 uint8, p2 int8) int32 {
    var v7 float64 = 3.14159
    for fi8 := 0; fi8 < 14; fi8++ {
        fmt.Println(v7)
        _ = fi8
    }
    v7 = ((v7 / v7) + 3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[uint32]uint32)
    var v3 float64 = (0.0 + (0.1 / 1.5))
    v4 := []int{-9223372036854775808}
    var v5 int64 = -538
    v6 := make(map[uint32]int)
    var v9 bool = (float32(3.14159) <= (1.5 * 0.1))
    var v10 int8 = -54
    if v9 {
        v11 := GenS1{F0: (v10 < v10)}
        var v12 int16 = 1
        if v9 {
            v3 -= (0.0 - v3)
            fmt.Println(v3, v12)
            v13 := GenS1{F0: false}
            _ = v13
        }
        _ = v11
        _ = v12
    }
    v14 := GenS1{F0: v9}
    var v15 bool = (!(!false))
    v10 += (v10 + 0)
    for fi16 := 0; fi16 < 17; fi16++ {
        if v9 {
            var v17 int64 = 1
            for fi18 := 0; fi18 < 9; fi18++ {
                var v19 uint = 18446744073709551615
                _ = fi18
                _ = v19
            }
            _ = v17
        }
        v10 += (v10 / (((v10 - v10)) | 1))
        _ = fi16
    }
    var v20 string = "a"
    for fi21 := 0; fi21 < 4; fi21++ {
        v10 = v10
        for fi22 := 0; fi22 < 1; fi22++ {
            fmt.Println(v15, v10)
            v5 -= v5
            _ = fi22
        }
        _ = fi21
    }
    var v23 int32 = v1
    v3 -= ((-2.25 / (v3 * v3)) + (float64((-2.25 - 1e10)) / v3))
    var v24 uint = 699
    var v25 uint32 = 2147483648
    for fi26 := 0; fi26 < 7; fi26++ {
        var v27 int32 = 312
        v28 := GenS2{F0: fi26, F1: (v10 * 42), F2: 1000}
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
