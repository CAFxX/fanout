package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 int; F2 int8 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 string = "a"
    var v3 float32 = (-2.25 * 1e10)
    var v4 bool = false
    v5 := GenS1{F0: (^(v1))}
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS2{F0: 696, F1: (fi6 - fi6), F2: 1}
        v2 = v2
        _ = fi6
        _ = v7
    }
    var v8 int8 = -26
    fmt.Println(v3, v4, v8)
    v3 += 0.0
    for fi9 := 0; fi9 < 3; fi9++ {
        v1 = v1
        if v4 {
            v10 := []int{(fi9 / ((-2147483648) | 1))}
            if len(v10) > 0 {
                v10[0] = 127
            }
            _ = v10
        }
        var v11 uint32 = 0
        _ = fi9
        _ = v11
    }
    v3 = (v3 - ((1e10 + (float32((0.0 * 1.5)) + (-2.25 - (v3 - (v3 * v3))))) / -2.25))
    var v12 int32 = 800
    if ((0 <= int16(v1)) && (18446744073709551615 <= (v1 * 18446744073709551615))) {
        var v13 int = -9223372036854775808
        v14 := []int{}
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = v13
        _ = v14
    }
    v12 = v12
    v3 = v3
    for fi15 := 0; fi15 < 14; fi15++ {
        v8 = (v8 & (v8 * (v8 >> ((0) & 63))))
        var v16 uint32 = 42
        _ = fi15
        _ = v16
    }
    v17 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
