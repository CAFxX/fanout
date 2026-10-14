package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS2 struct { F0 string; F1 uint64; F2 uint }
func genFn0(p0 uint64, p1 int8, p2 int16) uint8 {
    var v5 string = "hello"
    var v6 int16 = p2
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := []int{}
            var v10 uint16 = 528
            fmt.Println(p2)
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    v11 := GenS1{F0: 0, F1: 2, F2: 156}
    v12 := GenS1{F0: 2, F1: 215, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v5
    _ = v6
    return 255
}
func genFn1(p0 uint64) uint8 {
    p0 += p0
    for fi13 := 0; fi13 < 12; fi13++ {
        v14 := GenS1{F0: 288, F1: genFn0((p0 >> ((961) & 63)), 0, int16(p0)), F2: 2}
        var v15 float64 = 0.1
        v16 := GenS2{F0: "hello", F1: (^(p0)), F2: 593}
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 int32 = -4
    var v18 uint64 = 907
    var v19 int8 = 1
    _ = p0
    _ = v17
    _ = v18
    _ = v19
    return 90
}

func main() {
    var v1 int = 2147483647
    var v2 float32 = -2.25
    var v3 bool = (!true)
    var v4 uint16 = 656
    v3 = v3
    var v20 int = v1
    v21 := GenS1{F0: 1, F1: 1, F2: (v4 + (v4 ^ 834))}
    v22 := []int{(-(v20)), (v20 << ((2) & 63))}
    var v23 string = "hello"
    v2 += -2.25
    for fi24 := 0; fi24 < 7; fi24++ {
        _ = len(v22)
        _ = fi24
    }
    var v25 int64 = 42
    if v3 {
        v3 = (float64(v2) <= float64(float32(-2.25)))
        fmt.Println(v1)
        for fi26 := 0; fi26 < 19; fi26++ {
            for fi27 := 0; fi27 < 3; fi27++ {
                v28 := GenS2{F0: ("" + "foo bar"), F1: 2, F2: 0}
                var v29 uint32 = 2147483647
                _ = fi27
                _ = v28
                _ = v29
            }
            _ = fi26
        }
    } else {
        v2 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
}
