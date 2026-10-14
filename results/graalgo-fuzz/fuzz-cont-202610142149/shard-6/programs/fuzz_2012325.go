package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int8; F1 int64; F2 int8 }
type GenS3 struct { F0 string; F1 int64; F2 int }
func genFn0(p0 uint8, p1 uint64, p2 int64) int8 {
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: 255, F1: p0, F2: 2}
        p2 -= (p2 & -1)
        p2 -= (p2 + p2)
        _ = fi6
        _ = v7
    }
    fmt.Println(p2, p1, p2)
    v8 := GenS1{F0: 42, F1: (p0 + (p0 << ((65535) & 63))), F2: p1}
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(p2)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 42
}

func main() {
    var v1 int8 = 1
    v2 := make(map[uint32]int)
    v3 := []uint32{2147483648, 0, 168}
    v4 := make(map[uint32]string)
    var v5 uint32 = 2
    var v10 bool = true
    v11 := GenS2{F0: int8(v5), F1: -9223372036854775808, F2: v1}
    v12 := GenS3{F0: ("hello" + "hello"), F1: -422, F2: 1}
    switch "" {
    case "\n\t":
        var v13 int32 = -1
        _ = v13
    }
    var v14 int32 = 2147483647
    _ = len(v3)
    var v15 int8 = (v1 | v1)
    var v16 uint = 4294967295
    v10 = ((v14 - (v14 % ((v14) | 1))) < (-(v14)))
    for fi17 := 0; fi17 < 5; fi17++ {
        if true {
            var v18 uint8 = 1
            var v19 float32 = ((0.0 * 1e10) + 1.5)
            v3 = append(v3, (v5 - (v5 * (v5 << ((255) & 63)))))
            _ = v18
            _ = v19
        } else {
            v20 := GenS2{F0: v1, F1: 344, F2: v15}
            _ = v20
        }
        v15 -= v1
        v14 += 0
        _ = fi17
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        var v22 int16 = 1000
        var v23 uint64 = 4294967295
        _ = fi21
        _ = v22
        _ = v23
    }
    if v10 {
        if (!true) {
            for fi24 := 0; fi24 < 6; fi24++ {
                v16 = (v16 >> ((v16) & 63))
                v10 = (v10 || v10)
                _ = fi24
            }
            v25 := GenS3{F0: "", F1: 480, F2: 128}
            var v26 uint64 = 452
            _ = v25
            _ = v26
        }
    } else {
        v16 = (v16 >> ((v16) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
