package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int64 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int16; F1 int16 }
func genFn0(p0 int8, p1 bool) uint32 {
    var v4 int32 = -784
    p1 = p1
    v5 := GenS1{F0: 2, F1: 0, F2: -1}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = float32(0.0)
    var v3 uint64 = 2
    delete(v1, genFn0(-43, (false && (!false))))
    v6 := GenS1{F0: 0, F1: 981, F2: 42}
    var v7 uint32 = 898
    v8 := GenS1{F0: 255, F1: int64(v7), F2: -1}
    var v9 bool = false
    var v10 uint16 = 70
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 uint64 = (v3 - (v3 | v3))
        var v13 int = fi11
        v3 += (v3 - (v3 / ((v3) | 1)))
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int8 = 108
    for fi15 := 0; fi15 < 9; fi15++ {
        v3 = (v3 & (v3 - 1))
        v16 := v1[v7]
        v10 -= 1
        _ = fi15
        _ = v16
    }
    switch 4294967295 {
    case 744:
        fmt.Println(v14, v2, v9)
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 int32 = 1
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    var v20 uint = 870
    if v9 {
        var v21 int32 = 0
        _ = v21
    }
    switch "foo bar" {
    case "a":
        fmt.Println(v2, v7, v9)
    case "":
        v22 := GenS2{F0: uint32(v10), F1: true}
        _ = v22
    default:
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 uint8 = 2
            v25 := GenS2{F0: v7, F1: (!v9)}
            for fi26 := 0; fi26 < 19; fi26++ {
                v24 = 12
                _ = fi26
            }
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    switch 32767 {
    case 42:
        v14 -= (-(v14))
    case -130:
        switch -920 {
        case -2147483648:
            v14 = v14
        default:
            if true {
                v20 = (v20 << (((v20 + 897)) & 63))
            }
        }
    case 32767:
        v20 -= v20
    }
    v27 := []int{443, -576}
    switch "foo bar" {
    case "foo bar":
        var v28 uint16 = v10
        _ = v28
    case "a":
        v29 := GenS3{F0: -32768, F1: -135}
        _ = v29
    case "":
        var v30 uint = (v20 << (((^(v20))) & 63))
        _ = v30
    default:
        _ = len(v27)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
