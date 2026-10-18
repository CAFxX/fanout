package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }

func main() {
    var v1 bool = true
    v2 := []uint32{2147483647, 2, 2147483647}
    var v3 float32 = 0.0
    var v4 bool = (!true)
    var v5 uint = 952
    var v6 int32 = 42
    if ((170 == 36) || v1) {
        if ((1 < (v6 * 0)) && (2147483648 <= 4294967295)) {
            var v7 int8 = 127
            _ = v7
        }
    }
    _ = len(v2)
    var v8 float64 = ((float64(v3) + 3.14159) / 1.0)
    for fi9 := 0; fi9 < 7; fi9++ {
        if (1 != 255) {
            fmt.Println(v4, v1, v3)
            var v10 uint32 = 4294967295
            var v11 uint16 = 516
            _ = v10
            _ = v11
        } else {
            v12 := GenS1{F0: 9223372036854775807, F1: (v5 % (((v5 / (((v5 << ((452) & 63))) | 1))) | 1))}
            _ = v12
        }
        _ = fi9
    }
    var v13 bool = (42 <= v6)
    if len(v2) > 0 {
        v2[0] = 0
    }
    v2 = append(v2, 2147483648)
    v8 = v8
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
