package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 uint; F1 int64; F2 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16; F1 string }
func genFn0() int8 {
    var v5 int = -397
    var v6 int8 = -63
    var v7 int32 = 42
    _ = v5
    _ = v6
    _ = v7
    return v6
}
func genFn1(p0 uint64, p1 uint8, p2 int) int {
    if true {
        p0 = p0
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := []int{fi8, fi8}
            _ = fi8
            _ = v9
        }
    } else {
        if true {
            p1 = (p1 + p1)
        } else {
            var v10 bool = true
            v11 := GenS1{F0: v10}
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 9; fi12++ {
            p2 -= (fi12 / ((p2) | 1))
            _ = fi12
        }
    }
    v13 := GenS1{F0: false}
    v14 := []int{p2}
    v15 := GenS2{F0: 118, F1: -1}
    v16 := GenS2{F0: 30, F1: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    return p2
}

func main() {
    var v1 uint = 2
    var v2 uint64 = 2
    v3 := []int{2147483647, -718}
    var v4 bool = (127 < 128)
    fmt.Println(v4)
    v3 = append(v3, 311)
    if (238 <= uint8(v2)) {
        switch genFn1((^(v2)), 42, -2147483648) {
        case 819:
            v4 = (127 >= 127)
        case 2147483647:
            v17 := GenS3{F0: (v1 - (v1 << (((v1 & v1)) & 63))), F1: -1, F2: -2147483648}
            _ = v17
        case -9223372036854775808:
            var v18 int64 = -1
            _ = v18
        }
        var v19 uint16 = 42
        _ = v19
    } else {
        v20 := GenS4{F0: (v2 >= 0)}
        _ = v20
    }
    v21 := GenS2{F0: 42, F1: 32767}
    var v22 int = -163
    v1 += (v1 >> (((v1 - v1)) & 63))
    v23 := GenS4{F0: (!((v1 * (v1 << (((^(v1))) & 63))) == 2))}
    if ((1000 != -563) && v4) {
        for fi24 := 0; fi24 < 2; fi24++ {
            for fi25 := 0; fi25 < 11; fi25++ {
                var v26 uint32 = 0
                var v27 uint16 = 0
                _ = fi25
                _ = v26
                _ = v27
            }
            v2 = (v2 & (v2 * v2))
            _ = fi24
        }
        v4 = (42 <= 2)
    } else {
        v28 := GenS2{F0: 127, F1: -400}
        switch (^(v22)) {
        case -195:
            fmt.Println(v4, v2, v4)
        default:
            _ = len(v3)
        }
        _ = v28
    }
    for fi29 := 0; fi29 < 20; fi29++ {
        if (!false) {
            v30 := GenS1{F0: v4}
            var v31 float64 = (1e10 / (1.5 / 1e10))
            var v32 uint64 = (^(v2))
            _ = v30
            _ = v31
            _ = v32
        }
        v4 = (!(((0.1 > 0.0) || ((false && v4) && ((v1 / ((v1) | 1)) >= (v1 & (v1 | 348))))) && v4))
        var v33 int = v22
        _ = fi29
        _ = v33
    }
    for fi34 := 0; fi34 < 9; fi34++ {
        if ((0.1 + -2.25) == (1e10 - 0.1)) {
            var v35 int = (v22 - 254)
            _ = v35
        }
        _ = fi34
    }
    fmt.Println(v1, v2)
    v4 = v4
    v22 += v22
    if (v4 && v4) {
        v36 := GenS5{F0: 1, F1: ""}
        v37 := GenS3{F0: (^(v1)), F1: -9223372036854775808, F2: 2147483647}
        _ = v36
        _ = v37
    }
    var v38 int64 = -9223372036854775808
    switch v22 {
    case -9223372036854775808:
        v1 = 200
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v38
    _ = v4
}
