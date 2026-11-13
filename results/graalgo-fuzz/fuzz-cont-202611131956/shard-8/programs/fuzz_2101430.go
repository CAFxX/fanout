package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint64 }

func main() {
    var v1 uint32 = 1
    var v2 int64 = 9223372036854775807
    var v3 float32 = (3.14159 * (1e10 * 3.14159))
    v4 := []string{""}
    v5 := GenS1{F0: 2147483647, F1: 127}
    v2 -= v2
    for fi6 := 0; fi6 < 16; fi6++ {
        if (!((v2 % ((v2) | 1)) >= v2)) {
            if true {
                _ = len(v4)
            } else {
                var v7 uint8 = uint8(v1)
                fmt.Println(v1, fi6, v7)
                _ = v7
            }
        } else {
            v2 -= (^(v2))
        }
        _ = fi6
    }
    fmt.Println(v2, v2)
    if false {
        switch 42 {
        case -818:
            fmt.Println(v3, v2, v1)
        default:
            v8 := GenS1{F0: 0, F1: 97}
            _ = v8
        }
        var v9 float32 = (v3 - 3.14159)
        v10 := GenS2{F0: -32768, F1: v1, F2: 1}
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 2147483647, F1: 44}
    v3 += (v3 - (v3 / (v3 * v3)))
    var v12 uint16 = 268
    v2 = -9223372036854775808
    v12 -= (v12 / ((32767) | 1))
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
