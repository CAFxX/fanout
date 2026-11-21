package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 bool; F1 uint32; F2 int8 }
func genFn0() uint {
    var v7 int = 128
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := []int{1, v7, 32768}
        var v10 int64 = 504
        _ = fi8
        _ = v10
        _ = v9
    }
    switch 32767 {
    case 42:
        v11 := []int{v7, (^(v7)), 9223372036854775807, (v7 & (v7 & (-(v7))))}
        _ = v11
    default:
        var v12 float64 = 1e10
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        if (float64((float32(((1e10 * 0.1) * (0.1 / 3.14159))) - -2.25)) <= (1e10 + 3.14159)) {
            v14 := GenS1{F0: 0, F1: 4294967295}
            _ = v14
        } else {
            v7 = fi13
            var v15 int64 = 1
            _ = v15
        }
        switch v7 {
        case 42:
            v16 := GenS1{F0: 9223372036854775807, F1: 565}
            _ = v16
        case 2147483647:
            fmt.Println(fi13)
        case -9223372036854775808:
            var v17 uint64 = 42
            _ = v17
        default:
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 uint16 = 329
                _ = fi18
                _ = v19
            }
        }
        for fi20 := 0; fi20 < 16; fi20++ {
            v7 = (fi13 << ((2) & 63))
            var v21 int = (fi13 + (fi20 ^ -497))
            _ = fi20
            _ = v21
        }
        _ = fi13
    }
    _ = v7
    return 255
}

func main() {
    v1 := []string{"a", ("" + "\n\t")}
    var v2 uint64 = 18446744073709551615
    var v3 bool = ((591 >= 65535) || true)
    var v4 uint32 = 2147483648
    var v5 uint32 = v4
    var v6 int8 = 100
    if (109 <= -1) {
        fmt.Println(v6)
        v5 -= (^(v4))
        var v22 uint32 = v4
        _ = v22
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 int16 = 0
        v5 += v4
        _ = fi23
        _ = v24
    }
    var v25 bool = ((65535 == genFn0()) || ((1 >= -9223372036854775808) || (!true)))
    var v26 int8 = 0
    v6 = v26
    v27 := GenS1{F0: 773, F1: (v4 + (v4 - 42))}
    if v3 {
        v28 := GenS1{F0: -9223372036854775808, F1: v4}
        v4 = (^(v4))
        _ = v28
    }
    v6 -= -46
    v3 = true
    var v29 int = 128
    if v3 {
        v30 := GenS2{F0: -854}
        if (42 != 162) {
            v31 := GenS2{F0: (v29 / (((v29 ^ (v29 | v29))) | 1))}
            _ = v31
        }
        for fi32 := 0; fi32 < 11; fi32++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v33 int8 = (^(v6))
            _ = fi32
            _ = v33
        }
        _ = v30
    }
    v34 := GenS3{F0: (true && false), F1: (v2 + 350)}
    for fi35 := 0; fi35 < 19; fi35++ {
        var v36 uint64 = 0
        var v37 uint64 = v36
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 uint16 = 65535
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v6 -= -118
    var v39 bool = v3
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v40 := GenS4{F0: v25, F1: (v4 + v4), F2: 42}
        _ = v40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
