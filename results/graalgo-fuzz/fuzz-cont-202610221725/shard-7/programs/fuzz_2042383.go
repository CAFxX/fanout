package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int16 }

func main() {
    var v1 int = 127
    v2 := []int{(v1 % ((v1) | 1)), v1, v1}
    var v3 float32 = (0.1 / 1.0)
    var v4 int64 = 42
    var v5 int64 = (v4 % ((500) | 1))
    v3 += v3
    v3 = ((0.0 + ((0.1 / 1.0) * (v3 / v3))) - v3)
    v1 += 31
    var v6 int16 = 1000
    if true {
        var v7 float32 = (v3 * float32(-2.25))
        v7 -= 1e10
        _ = v7
    } else {
        fmt.Println(v5, v6, v6)
    }
    switch (v1 / (((v1 << ((18446744073709551615) & 63))) | 1)) {
    case 2147483647:
        if (42 >= 193) {
            fmt.Println(v6, v5)
            if true {
                v8 := GenS1{F0: 42, F1: v6, F2: -1}
                _ = len(v2)
                _ = v8
            } else {
                for fi9 := 0; fi9 < 13; fi9++ {
                    for fi10 := 0; fi10 < 15; fi10++ {
                        var v11 float32 = (((v3 * v3) + float32(float64(v3))) * ((((v3 - 0.0) / v3) / (v3 * (float32(1.5) / v3))) / (-2.25 + v3)))
                        _ = fi10
                        _ = v11
                    }
                    v3 += (-2.25 - (0.1 - float32(0.0)))
                    _ = fi9
                }
                var v12 int8 = 0
                v5 += 571
                _ = v12
            }
            var v13 float32 = v3
            _ = v13
        }
    case 978:
        if len(v2) > 0 {
            v2[0] = v1
        }
    }
    var v14 float64 = -2.25
    var v15 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
