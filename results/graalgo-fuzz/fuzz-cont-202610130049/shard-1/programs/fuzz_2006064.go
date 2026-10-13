package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint8 }
func genFn0(p0 uint8, p1 int8) bool {
    p0 -= p0
    p0 += p0
    _ = p0
    _ = p1
    return true
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 42
    var v3 int8 = -106
    v4 := []string{"hello", "\n\t"}
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: (v3 * v3)}
        v2 -= (v2 << ((735) & 63))
        _ = fi5
        _ = v6
    }
    if (true || true) {
        if (uint64(v2) != 42) {
            v2 = (v2 | 39)
            v7 := GenS2{F0: 229, F1: 127, F2: 0}
            _ = v7
        }
    }
    if (0 == 42) {
        v3 = (v3 / (((v3 & 86)) | 1))
        if true {
            var v8 int8 = -59
            _ = v8
        }
    } else {
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := v1[(-(fi9))]
            v11 := GenS3{F0: -1, F1: 255, F2: v2}
            _ = fi9
            _ = v10
            _ = v11
        }
        v3 = (v3 / (((-(v3))) | 1))
        v3 = (v3 * (v3 >> ((414) & 63)))
    }
    var v12 uint8 = 77
    if true {
        v2 = v12
        if false {
            var v13 uint = 1
            _ = v13
        }
        v12 = (v12 << ((0) & 63))
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        fmt.Println(v3)
        _ = fi14
    }
    v2 += (v12 | (v2 + v12))
    for fi15 := 0; fi15 < 4; fi15++ {
        fmt.Println(v12, v12, v2)
        _ = fi15
    }
    var v16 bool = (true && false)
    v17 := GenS1{F0: v3}
    v3 += (v3 * v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
