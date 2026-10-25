package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 float64 = (float64(((1e10 + 3.14159) / ((1.5 - 3.14159) / -2.25))) - 3.14159)
    var v2 int16 = 660
    var v3 int = 128
    for fi4 := 0; fi4 < 8; fi4++ {
        switch (^(fi4)) {
        case -91:
            var v5 int = fi4
            _ = v5
        default:
            if (v1 > 1.5) {
                var v6 int64 = 9223372036854775807
                _ = v6
            }
        }
        _ = fi4
    }
    var v7 uint32 = 0
    v8 := GenS1{F0: 1, F1: 2}
    var v9 int64 = 9223372036854775807
    v10 := GenS1{F0: 1, F1: 32767}
    if true {
        if (!true) {
            if (v2 <= v2) {
                v11 := GenS1{F0: 32767, F1: 255}
                if (!true) {
                    v3 -= 2147483647
                    v12 := []int{(v3 + (v3 + (-(v3)))), (v3 ^ -9223372036854775808), (v3 / ((-9223372036854775808) | 1))}
                    _ = v12
                }
                var v13 int64 = v9
                _ = v11
                _ = v13
            } else {
                var v14 uint16 = 32767
                var v15 int8 = 127
                _ = v14
                _ = v15
            }
            var v16 bool = true
            _ = v16
        } else {
            var v17 float32 = 1e10
            _ = v17
        }
        var v18 int64 = 189
        var v19 uint64 = 42
        _ = v18
        _ = v19
    } else {
        v9 += 0
        var v20 int16 = v2
        _ = v20
    }
    fmt.Println(v9, v7)
    var v21 int = v3
    v22 := []int{(v21 * 2147483647), 733}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
