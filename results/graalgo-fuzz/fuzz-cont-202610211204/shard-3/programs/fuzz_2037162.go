package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int, p1 bool) int64 {
    v4 := GenS1{F0: (374 < 0)}
    p0 -= p0
    p0 = -1
    v5 := []int{p0, p0, -943}
    if len(v5) > 0 {
        v5[0] = p0
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -804
}
func genFn1() uint64 {
    var v6 uint64 = 1
    v6 = 4294967295
    var v7 uint32 = 0
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 float32 = (0.0 + 0.0)
    var v2 int16 = 42
    var v3 float64 = 1.5
    v8 := []int{1, 2147483647}
    v8 = append(v8, 0)
    v9 := GenS1{F0: (-1 == 42)}
    v1 = (-2.25 / ((v1 - 0.1) / ((float32((v3 - 1e10)) / 1.5) + v1)))
    for fi10 := 0; fi10 < 11; fi10++ {
        v1 = (v1 + v1)
        var v11 int16 = v2
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := GenS1{F0: true}
        _ = fi12
        _ = v13
    }
    fmt.Println(v1, v2)
    fmt.Println(v1, v3)
    var v14 uint64 = 93
    v2 = -932
    if (!true) {
        v8 = append(v8, -2147483648)
        v2 -= v2
    } else {
        fmt.Println(v3, v14)
        v3 -= (v3 * (1e10 + v3))
        v8 = append(v8, 98)
    }
    if (v1 <= v1) {
        v2 = (v2 | 1000)
        var v15 uint16 = 46
        v16 := GenS2{F0: (v14 >> ((165) & 63)), F1: 1, F2: 0}
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        for fi18 := 0; fi18 < 13; fi18++ {
            v3 = (1.5 / v3)
            _ = fi18
        }
        v19 := GenS3{F0: 1}
        _ = fi17
        _ = v19
    }
    v8 = append(v8, 42)
    var v20 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v8
    _ = v9
}
