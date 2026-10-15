package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }

func main() {
    var v1 float64 = 0.0
    var v2 uint16 = 1
    v3 := []int{-238}
    fmt.Println(v1, v1)
    v1 += v1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi4 := 0; fi4 < 17; fi4++ {
        if len(v3) > 0 {
            v3[0] = fi4
        }
        v2 += v2
        _ = fi4
    }
    if len(v3) > 0 {
        v3[0] = -240
    }
    v5 := GenS1{F0: -1, F1: 2}
    v1 -= (0.1 / float64(3.14159))
    v1 = (3.14159 * 3.14159)
    if (!(!true)) {
        v1 -= (v1 * (v1 - v1))
        if true {
            for fi6 := 0; fi6 < 10; fi6++ {
                var v7 float32 = 3.14159
                v1 -= float64(((-2.25 + v7) - ((float32(v1) + 1.5) + (v7 * (v7 / (v7 + v7))))))
                fmt.Println(fi6, v2)
                _ = fi6
                _ = v7
            }
        }
        var v8 uint32 = 4294967295
        _ = v8
    } else {
        v1 = (((float64(3.14159) * v1) + 1.5) / v1)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v9 int16 = -877
        _ = v9
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
}
