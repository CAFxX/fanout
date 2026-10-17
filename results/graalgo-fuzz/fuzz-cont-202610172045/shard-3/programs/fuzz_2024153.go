package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 float32 = 0.0
    var v3 int16 = 42
    if true {
        var v4 uint32 = 560
        v5 := GenS1{F0: 127, F1: 1, F2: 9223372036854775807}
        _ = v4
        _ = v5
    } else {
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := []int{}
            v3 -= (v3 * int16(fi6))
            var v8 uint = 2
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    v9 := []int{-1, -144, 995, 32767}
    var v10 float64 = float64(float32(1.5))
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := v1["a"]
        fmt.Println(v10, v2)
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        if ((false || false) && (295 < 18446744073709551615)) {
            var v14 int64 = 9223372036854775807
            var v15 uint32 = 65535
            _ = v14
            _ = v15
        } else {
            v10 -= ((v10 + (v10 * v10)) + 3.14159)
        }
        var v16 uint = 707
        _ = fi13
        _ = v16
    }
    if (v10 == float64(float32(float64(v2)))) {
        var v17 int64 = 0
        if len(v9) > 0 {
            v9[0] = -9223372036854775808
        }
        _ = v17
    }
    if len(v9) > 0 {
        v9[0] = 32767
    }
    var v18 uint = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
