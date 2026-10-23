package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int8 }

func main() {
    v1 := []int{0, -288, 0}
    var v2 uint16 = 478
    var v3 bool = true
    v1 = append(v1, 936)
    if ((3.14159 * 1e10) == 3.14159) {
        if v3 {
            v3 = (v3 && (!v3))
            var v4 uint = 348
            _ = v4
        }
        var v5 uint32 = 609
        _ = v5
    } else {
        for fi6 := 0; fi6 < 3; fi6++ {
            if len(v1) > 0 {
                v1[0] = (fi6 / ((fi6) | 1))
            }
            v3 = v3
            _ = fi6
        }
    }
    v3 = v3
    var v7 uint16 = (v2 / (((v2 * (^(v2)))) | 1))
    v8 := GenS1{F0: -751, F1: v7}
    v9 := GenS2{F0: 127, F1: 5}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = (v7 >> ((1) & 63))
    var v10 uint = 183
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
