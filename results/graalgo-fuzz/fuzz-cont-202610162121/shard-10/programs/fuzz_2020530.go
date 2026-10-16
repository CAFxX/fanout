package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int64 }
type GenS2 struct { F0 int64; F1 int8; F2 uint64 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint }
func genFn0(p0 int) int8 {
    v6 := GenS1{F0: 0, F1: 107, F2: 9223372036854775807}
    var v7 int16 = 139
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 float64 = 1e10
        var v10 uint32 = 111
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: -1, F1: 127, F2: 529}
    _ = p0
    _ = v11
    _ = v6
    _ = v7
    return -59
}
func genFn1(p0 bool) uint32 {
    p0 = (9223372036854775807 > 9223372036854775807)
    v12 := GenS1{F0: 100, F1: 554, F2: 42}
    var v13 uint64 = 2
    fmt.Println(p0)
    p0 = (p0 || p0)
    _ = p0
    _ = v12
    _ = v13
    return 80
}

func main() {
    var v1 bool = (100 != -41)
    var v2 int64 = -731
    var v3 float32 = 0.1
    var v4 uint8 = 150
    v5 := make(map[uint32]bool)
    v14 := GenS2{F0: (v2 * (v2 / ((v2) | 1))), F1: 0, F2: 552}
    fmt.Println(v3, v4, v2)
    var v15 float64 = 0.0
    v15 += float64(0.0)
    var v16 int16 = 32767
    if (v3 != (1.5 + v3)) {
        v17 := []int{905, 128, 128}
        var v18 int64 = v2
        _ = v17
        _ = v18
    }
    v3 = v3
    fmt.Println(v15, v16)
    v19 := GenS1{F0: 127, F1: 18446744073709551615, F2: v2}
    fmt.Println(v3, v1, v3)
    var v20 uint64 = 18446744073709551615
    var v21 uint32 = 680
    var v22 int16 = v16
    v23 := GenS3{F0: -1, F1: 2}
    var v24 int32 = -1
    if false {
        if (100 < 42) {
            v3 += float32(v15)
            v25 := GenS3{F0: v24, F1: (^(v4))}
            _ = v25
        }
        v16 -= (v16 & 1000)
    }
    v26 := GenS4{F0: 712}
    for fi27 := 0; fi27 < 2; fi27++ {
        v22 += (v22 ^ (v16 ^ 32767))
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
