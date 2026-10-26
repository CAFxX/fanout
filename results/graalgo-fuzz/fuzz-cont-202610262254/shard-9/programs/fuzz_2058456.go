package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint32 }

func main() {
    v1 := make(map[string]int)
    var v2 float64 = float64((1e10 * (1.5 * (0.0 + (float32(-2.25) / 1.0)))))
    var v3 uint8 = 1
    v4 := GenS1{F0: 1, F1: 135, F2: 2}
    v5 := []int{-880, 2147483647}
    fmt.Println(v2, v3)
    var v6 uint8 = (v3 >> ((65535) & 63))
    v2 = 0.1
    var v7 uint32 = 2147483647
    if true {
        v6 = v6
        var v8 int32 = -205
        v9 := GenS1{F0: 100, F1: 18446744073709551615, F2: (^(v7))}
        _ = v8
        _ = v9
    } else {
        v10 := GenS1{F0: 0, F1: 18446744073709551615, F2: uint32(v3)}
        v6 -= 128
        if len(v5) > 0 {
            v5[0] = 128
        }
        _ = v10
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v11 float64 = v2
    var v12 int8 = 105
    if len(v5) > 0 {
        v5[0] = -1
    }
    var v13 int32 = -800
    v5 = append(v5, 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
