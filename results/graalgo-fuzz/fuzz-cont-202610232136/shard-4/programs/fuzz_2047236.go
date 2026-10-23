package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int16; F2 int }
type GenS3 struct { F0 int16; F1 int64; F2 int16 }
type GenS4 struct { F0 int8; F1 int; F2 uint32 }

func main() {
    var v1 int = 127
    v2 := []uint32{4294967295, 4294967295, 2}
    v3 := make(map[int]bool)
    var v4 string = "foo bar"
    v5 := make(map[uint32]bool)
    v6 := GenS1{F0: 1}
    v7 := GenS2{F0: false, F1: 105, F2: 32767}
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(v4, v1, v4)
        _ = fi8
    }
    if true {
        fmt.Println(v1, v1)
        switch "foo bar" {
        case "a":
            fmt.Println(v1, v4)
        default:
            var v9 uint = 4294967295
            _ = v9
        }
    } else {
        v10 := GenS2{F0: ((v4 + (v4 + "a")) >= v4), F1: -501, F2: (v1 & v1)}
        _ = v10
    }
    v1 -= (^(v1))
    switch 2147483647 {
    case 85:
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := GenS1{F0: 0}
            for fi13 := 0; fi13 < 11; fi13++ {
                if (false && true) {
                    var v14 int = 840
                    var v15 int16 = -32768
                    _ = v14
                    _ = v15
                } else {
                    if len(v2) > 0 {
                        v2[0] = 535
                    }
                    var v16 bool = (32767 < 59)
                    _ = v16
                }
                var v17 float32 = 0.0
                _ = fi13
                _ = v17
            }
            _ = fi11
            _ = v12
        }
    case 1:
        v1 -= (v1 - (v1 >> ((943) & 63)))
    case -199:
        v18 := GenS2{F0: false, F1: -470, F2: v1}
        _ = v18
    default:
        if true {
            var v19 int16 = 1
            for fi20 := 0; fi20 < 6; fi20++ {
                v4 = v4
                v4 = v4
                _ = fi20
            }
            _ = v19
        } else {
            var v21 uint = 1
            _ = len(v2)
            v22 := GenS3{F0: -2, F1: 42, F2: int16(v1)}
            _ = v21
            _ = v22
        }
    }
    var v23 int32 = -1
    for fi24 := 0; fi24 < 13; fi24++ {
        var v25 int64 = 332
        switch -1 {
        case -975:
            v26 := GenS4{F0: 0, F1: (fi24 * fi24), F2: 801}
            _ = v26
        case -2147483648:
            var v27 int16 = 42
            _ = v27
        default:
            fmt.Println(v25, v25, v1)
        }
        var v28 uint16 = 0
        _ = fi24
        _ = v25
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
