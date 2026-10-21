package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }

func main() {
    var v1 int32 = -2147483648
    var v2 int8 = 100
    var v3 int = -474
    v4 := []int{-2147483648, (-(v3)), (^(v3))}
    v5 := GenS1{F0: 42, F1: 2}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS1{F0: (v1 | v1), F1: 42}
        v8 := GenS1{F0: (v1 + (v1 + v1)), F1: 4294967295}
        if (int32(v2) <= 164) {
            fmt.Println(fi6)
            switch -575 {
            case -1:
                fmt.Println(v1, v3)
            default:
                v9 := GenS1{F0: (v1 * v1), F1: 4294967295}
                _ = v9
            }
            v2 += (v2 - (v2 ^ v2))
        } else {
            var v10 int16 = 32767
            _ = v10
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    v3 = (v3 % (((v3 >> ((65535) & 63))) | 1))
    var v11 float32 = (3.14159 + (0.0 + 3.14159))
    v3 -= (v3 - (v3 / ((v3) | 1)))
    var v12 uint = 1
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
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
