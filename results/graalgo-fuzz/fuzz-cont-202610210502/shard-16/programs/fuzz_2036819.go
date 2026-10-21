package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 bool }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint64, p1 uint16, p2 uint8) bool {
    p1 -= uint16(p0)
    v4 := GenS1{F0: 310, F1: 791, F2: true}
    var v5 uint16 = p1
    var v6 int64 = 42
    v5 = v5
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return true
}

func main() {
    v1 := []int{9223372036854775807}
    var v2 float64 = float64((-2.25 * (((0.0 + (1e10 * (0.0 - (float32(1e10) / (0.1 - 0.0))))) - (float32(float64(float32(float64(3.14159)))) + (1.5 / 1.0))) / 1.5)))
    var v3 int16 = 1000
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(v2)
        _ = fi7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        v2 += ((float64((((0.0 / (3.14159 / 1.0)) / 1.5) / 1.0)) / (v2 - v2)) + 0.0)
        v9 := GenS1{F0: v3, F1: 4294967295, F2: ((440 == -1) && (3.14159 >= -2.25))}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 int32 = -470
        if false {
            for fi12 := 0; fi12 < 20; fi12++ {
                fmt.Println(fi10, fi10)
                v11 = v11
                var v13 int32 = -2147483648
                var v14 uint8 = 82
                _ = fi12
                _ = v13
                _ = v14
            }
            fmt.Println(v3, v11, fi10)
        }
        _ = fi10
        _ = v11
    }
    var v15 bool = true
    for fi16 := 0; fi16 < 20; fi16++ {
        v17 := GenS2{F0: 4294967295}
        fmt.Println(fi16)
        var v18 int = (fi16 / ((fi16) | 1))
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 uint64 = 917
    var v20 int = 490
    var v21 uint = 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
