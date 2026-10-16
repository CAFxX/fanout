package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int64 }
func genFn0() uint {
    var v5 float32 = (float32((1.5 * (0.1 + (((1e10 - float64(float32(-2.25))) - 0.0) / 1.5)))) / -2.25)
    var v6 uint = 65535
    v7 := GenS1{F0: 42}
    var v8 int16 = 42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (v6 | (v6 / ((v6) | 1)))
}

func main() {
    v1 := []int{42, -54}
    var v2 uint16 = 1
    v3 := make(map[uint32]uint32)
    var v4 int16 = 32767
    v2 = v2
    fmt.Println(v4, v4)
    var v9 int8 = 14
    for fi10 := 0; fi10 < 8; fi10++ {
        if false {
            fmt.Println(v2, v9, v4)
            v9 -= v9
            var v11 int64 = -1
            _ = v11
        }
        var v12 uint32 = 2147483648
        _ = fi10
        _ = v12
    }
    if ((v9 << ((1) & 63)) == v9) {
        v13 := GenS2{F0: v9, F1: -9223372036854775808}
        v1 = append(v1, 145)
        v4 = v4
        _ = v13
    } else {
        var v14 uint16 = v2
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v14
    }
    v9 -= v9
    v9 = 127
    fmt.Println(v4, v2)
    var v15 int16 = (v4 * (^(v4)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
