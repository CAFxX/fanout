package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 bool; F1 bool; F2 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 uint32 }

func main() {
    var v1 int32 = 438
    var v2 int16 = 1
    var v3 int64 = -1
    v4 := make(map[uint32]bool)
    v5 := make(map[int]string)
    v3 -= (-(v3))
    var v6 uint = 1
    var v7 uint16 = 115
    if (836 != 0) {
        v8 := GenS1{F0: 127, F1: 919}
        v2 += 1
        _ = v8
    }
    var v9 uint32 = 4294967295
    var v10 int8 = 99
    v9 -= (v9 * (v9 * 65535))
    if (true && (32768 < 2147483647)) {
        var v11 int64 = v3
        v7 = (^(v7))
        var v12 int = 32767
        _ = v11
        _ = v12
    } else {
        v13 := GenS2{F0: (!false), F1: false, F2: int8(v6)}
        v14 := v4[(v9 - (v9 & (v9 >> (((v6 + v6)) & 63))))]
        v3 = v3
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 127, F1: 4294967295}
    var v16 int16 = v2
    if (!(-626 <= -2147483648)) {
        if true {
            var v17 int8 = v10
            _ = v17
        }
        switch 32767 {
        case -492:
            v2 = v2
        case 594:
            v18 := GenS3{F0: (v10 ^ (^(v10)))}
            _ = v18
        case -1:
            v6 -= v6
        }
    }
    v2 -= v2
    v9 = v9
    v19 := v4[v9]
    v20 := GenS4{F0: (v6 | v6), F1: (v9 + v9)}
    v1 += (v1 * v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
