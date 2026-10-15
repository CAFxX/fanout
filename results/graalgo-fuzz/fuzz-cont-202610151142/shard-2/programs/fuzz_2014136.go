package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int32) uint8 {
    v5 := GenS1{F0: 0, F1: (p0 << ((42) & 63)), F2: 32767}
    p0 -= -1
    v6 := GenS2{F0: 2147483647}
    p0 -= p0
    if false {
        if true {
            fmt.Println(p0)
            v7 := []int{1, 554, -854}
            var v8 float64 = 0.0
            _ = v7
            _ = v8
        }
    }
    _ = p0
    _ = v5
    _ = v6
    return 166
}

func main() {
    var v1 uint = 0
    var v2 uint16 = 328
    var v3 float32 = (float32(3.14159) * 1.5)
    var v4 int16 = 0
    v9 := []int{-272, 32768, 1}
    if true {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        var v10 int64 = 0
        if (-1 > 4294967295) {
            for fi11 := 0; fi11 < 17; fi11++ {
                fmt.Println(v4, v10)
                fmt.Println(v1, v2, fi11)
                _ = fi11
            }
        }
        _ = v10
    }
    v2 = v2
    v4 = (v4 + 730)
    v1 += v1
    for fi12 := 0; fi12 < 17; fi12++ {
        v2 = v2
        fmt.Println(v2)
        _ = fi12
    }
    var v13 uint8 = genFn0(679)
    v14 := GenS2{F0: 2}
    var v15 float64 = 0.0
    var v16 int32 = -1
    var v17 float64 = v15
    for fi18 := 0; fi18 < 10; fi18++ {
        fmt.Println(v13, v3)
        var v19 uint16 = (v2 / ((390) | 1))
        _ = fi18
        _ = v19
    }
    switch -403 {
    case 32768:
        if true {
            var v20 uint = v1
            _ = len(v9)
            var v21 int64 = -1
            _ = v20
            _ = v21
        }
    case 128:
        var v22 int16 = v4
        _ = v22
    default:
        v9 = append(v9, -10)
    }
    fmt.Println(v17)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
