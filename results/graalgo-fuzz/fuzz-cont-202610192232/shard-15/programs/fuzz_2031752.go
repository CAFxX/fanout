package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int }
type GenS2 struct { F0 uint16; F1 int32 }
func genFn0(p0 int64, p1 int, p2 bool) uint16 {
    p1 = p1
    var v4 int64 = 171
    p0 += p0
    var v5 uint32 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 75
}
func genFn1(p0 uint, p1 int) int8 {
    fmt.Println(p0, p1)
    p1 = p1
    p1 = p1
    fmt.Println(p0)
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 100
}

func main() {
    var v1 uint32 = 868
    var v2 int = 4294967295
    v3 := []int{(v2 << ((0) & 63)), v2, v2, (v2 & 128)}
    var v6 int32 = -2147483648
    v1 -= v1
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 uint = 495
        _ = fi7
        _ = v8
    }
    _ = len(v3)
    v9 := GenS1{F0: 0, F1: -266, F2: -359}
    if false {
        var v10 int32 = v6
        _ = v10
    }
    var v11 float64 = 0.1
    _ = len(v3)
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 int16 = 1000
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 1; fi15++ {
            fmt.Println(fi12, v11, v11)
            var v16 uint32 = v1
            _ = len(v3)
            _ = fi15
            _ = v16
        }
        v17 := GenS2{F0: 261, F1: v6}
        _ = fi12
        _ = v17
    }
    switch (v2 | -617) {
    case 1:
        var v18 int = -2147483648
        _ = v18
    default:
        v6 += v6
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        fmt.Println(v11)
        var v20 int64 = -590
        var v21 float32 = 0.0
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
