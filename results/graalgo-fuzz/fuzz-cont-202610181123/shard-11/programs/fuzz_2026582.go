package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint16 = 372
    var v2 int8 = 0
    var v3 int16 = -1
    var v4 int = -266
    var v5 float32 = 1e10
    var v6 uint64 = 488
    var v7 int32 = int32(v4)
    v4 = -759
    v1 += v1
    fmt.Println(v5, v6, v2)
    v1 -= v1
    v4 -= v4
    v8 := []int{(v4 & -444), v4}
    if (65535 <= 74) {
        var v9 float64 = 0.0
        if false {
            v10 := GenS1{F0: (v4 * (v4 >> ((504) & 63))), F1: v1}
            for fi11 := 0; fi11 < 3; fi11++ {
                v12 := GenS1{F0: (fi11 + -594), F1: v1}
                _ = fi11
                _ = v12
            }
            v6 += v6
            _ = v10
        } else {
            var v13 int = (v4 | v4)
            _ = v13
        }
        _ = v9
    } else {
        v3 = (v3 | (v3 + (v3 >> ((42) & 63))))
        v14 := GenS2{F0: 887}
        fmt.Println(v2, v1, v5)
        _ = v14
    }
    var v15 uint = 385
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v4 = v4
    var v16 int16 = int16(v15)
    v6 -= (v6 - 0)
    var v17 int = v4
    if true {
        v17 -= (-(v4))
        v5 -= 0.0
        var v18 uint8 = 98
        _ = v18
    }
    var v19 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
