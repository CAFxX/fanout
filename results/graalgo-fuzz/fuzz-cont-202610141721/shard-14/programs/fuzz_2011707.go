package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int64; F1 uint16 }
func genFn0(p0 uint8) bool {
    v5 := []int{-1, 2147483647, -996}
    fmt.Println(p0)
    p0 = p0
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 int64 = 0
            v9 := GenS1{F0: -281, F1: 1000}
            if true {
                var v10 uint8 = p0
                v10 -= (v10 ^ (p0 | 10))
                _ = v10
            } else {
                v11 := GenS2{F0: 42, F1: 0}
                _ = v11
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        fmt.Println(fi6, fi6)
        var v12 int64 = -1
        _ = fi6
        _ = v12
    }
    if true {
        fmt.Println(p0)
    }
    _ = p0
    _ = v5
    return false
}
func genFn1() uint8 {
    var v13 float64 = float64(0.0)
    v14 := GenS2{F0: 220, F1: 2}
    fmt.Println(v13)
    v13 = v13
    _ = v13
    _ = v14
    return 2
}

func main() {
    var v1 uint = 912
    var v2 uint8 = 4
    var v3 int = 127
    var v4 uint64 = 76
    var v15 uint16 = 0
    var v16 uint32 = 384
    v16 = (v16 % ((v16) | 1))
    v17 := GenS3{F0: 42, F1: v15}
    if false {
        v16 = (v16 ^ (v16 * (^(v16))))
        var v18 int8 = 0
        var v19 float32 = 1e10
        _ = v18
        _ = v19
    } else {
        for fi20 := 0; fi20 < 11; fi20++ {
            v15 += v15
            _ = fi20
        }
        v21 := []int{v3, (v3 >> ((v1) & 63))}
        _ = v21
    }
    v3 += (v3 * (v3 / (((v3 & (v3 % ((32767) | 1)))) | 1)))
    for fi22 := 0; fi22 < 17; fi22++ {
        v1 = 18446744073709551615
        _ = fi22
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        fmt.Println(v4)
        fmt.Println(v1, v15)
        _ = fi23
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        v3 += (v3 ^ 42)
        var v25 bool = false
        _ = fi24
        _ = v25
    }
    v4 = (v4 ^ 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
