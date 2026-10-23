package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint }

func main() {
    var v1 uint64 = 4294967295
    var v2 int64 = -1
    var v3 uint32 = 0
    var v4 uint32 = v3
    v5 := make(map[int]bool)
    v6 := []int{-1, 32767, 128}
    var v7 bool = false
    v3 = (v3 * 1)
    var v8 float32 = ((0.1 * 1.5) / ((0.1 + 1.5) - 1e10))
    var v9 int64 = -1
    v10 := GenS1{F0: (v2 - v9), F1: 1}
    var v11 int8 = -125
    v4 = v3
    for fi12 := 0; fi12 < 11; fi12++ {
        v6 = append(v6, (fi12 & (fi12 ^ fi12)))
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 uint32 = (v3 & (v4 * (v3 / (((v3 << ((42) & 63))) | 1))))
            v5[128] = ((v2 * (-(v2))) > v9)
            v15 := v5[(-(fi13))]
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    if (!(v11 < v11)) {
        if v7 {
            if false {
                v5[1] = v7
                if v7 {
                    _ = len(v6)
                }
            }
            var v16 uint8 = 42
            _ = v16
        }
        fmt.Println(v11)
    }
    if true {
        v17 := GenS2{F0: v3, F1: 42, F2: v1}
        v9 = (-(v2))
        _ = v17
    } else {
        v1 = (v1 + 4294967295)
        delete(v5, -9223372036854775808)
    }
    var v18 int16 = 1
    v19 := GenS3{F0: 2}
    v11 += -1
    v18 += (v18 % ((-58) | 1))
    for fi20 := 0; fi20 < 20; fi20++ {
        v7 = v7
        _ = fi20
    }
    v8 = float32((((1.5 + (0.1 * (3.14159 - (1e10 * (((3.14159 + 1.5) - float64(((float32(1.5) * 3.14159) / float32(3.14159)))) - 1.5))))) / 1.0) * float64(((v8 - ((((-2.25 / 1.0) * v8) / v8) - ((v8 * 1.5) - (v8 + 1e10)))) / 3.14159))))
    var v21 uint32 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
