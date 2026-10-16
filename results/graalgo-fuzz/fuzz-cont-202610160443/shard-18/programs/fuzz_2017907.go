package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int; F2 int8 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 uint8 = 233
    v2 := make(map[string]int)
    v3 := []int{2147483647}
    v4 := make(map[uint32]int)
    var v5 float64 = 1e10
    v1 -= (v1 % (((v1 >> ((4294967295) & 63))) | 1))
    v6 := GenS1{F0: 2, F1: 42}
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS2{F0: 104, F1: fi7, F2: 42}
        _ = fi7
        _ = v8
    }
    fmt.Println(v1, v1)
    v1 += 109
    var v9 bool = (2147483648 > 65535)
    v5 += -2.25
    v3 = append(v3, 9223372036854775807)
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 int8 = -128
        _ = fi10
        _ = v11
    }
    if v9 {
        v12 := GenS2{F0: 0, F1: 1, F2: -31}
        _ = v12
    }
    fmt.Println(v1)
    if (v9 || v9) {
        fmt.Println(v5, v1, v5)
        if len(v3) > 0 {
            v3[0] = 417
        }
    } else {
        var v13 uint64 = 1
        if (-2147483648 <= -1) {
            v13 = (^(v13))
            if false {
                v13 += (v13 * (^(v13)))
                v14 := GenS3{F0: -128}
                _ = v14
            } else {
                v1 = 2
            }
        }
        _ = v13
    }
    delete(v2, "hello")
    var v15 int64 = 909
    fmt.Println(v9, v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
