package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 42
    v3 := make(map[int]uint32)
    var v4 int = 127
    var v5 int32 = 1
    var v6 uint16 = 0
    switch "foo bar" {
    case "":
        if false {
            for fi7 := 0; fi7 < 1; fi7++ {
                for fi8 := 0; fi8 < 1; fi8++ {
                    var v9 int32 = (-(v5))
                    var v10 uint8 = uint8(v6)
                    v11 := []int{int(v5), (v4 ^ (v4 ^ v4)), (fi8 * (fi7 << ((318) & 63))), (v4 | fi7)}
                    if true {
                        if ((float64(1e10) + 1e10) != -2.25) {
                            var v12 bool = (v9 <= (v9 + v5))
                            _ = v12
                        }
                        v2 = (v2 >> ((1) & 63))
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        v10 = 128
                        v13 := GenS1{F0: uint32(v10), F1: fi8}
                        _ = v13
                    }
                    v14 := v1[-115]
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v14
                    _ = v9
                }
                _ = fi7
            }
            var v15 uint8 = 128
            switch 645 {
            case -314:
                var v16 int64 = 587
                _ = v16
            case 1:
                fmt.Println(v2, v5, v5)
            default:
                var v17 bool = true
                _ = v17
            }
            _ = v15
        }
    case "a":
        v6 += (v6 << ((687) & 63))
    default:
        v4 = v4
    }
    switch "foo bar" {
    case "":
        v18 := GenS2{F0: -83, F1: -32768}
        _ = v18
    case "\n\t":
        for fi19 := 0; fi19 < 13; fi19++ {
            for fi20 := 0; fi20 < 16; fi20++ {
                v4 += fi19
                _ = fi20
            }
            _ = fi19
        }
    case "a":
        var v21 uint64 = v2
        _ = v21
    default:
        v22 := GenS1{F0: 0, F1: v4}
        _ = v22
    }
    v2 = v2
    v23 := GenS2{F0: v5, F1: 874}
    var v24 uint64 = (v2 * 88)
    for fi25 := 0; fi25 < 16; fi25++ {
        v24 -= v24
        v26 := v1[-683]
        var v27 uint32 = 42
        _ = fi25
        _ = v26
        _ = v27
    }
    if false {
        v1[v4] = ""
    }
    v28 := GenS3{F0: 2}
    v5 = (v5 >> ((18446744073709551615) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
