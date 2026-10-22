package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 uint16; F2 string }
func genFn0(p0 string, p1 uint16, p2 uint) uint16 {
    v5 := []int{9223372036854775807}
    v6 := GenS1{F0: 42, F1: -1, F2: 2147483647}
    var v7 int64 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return p1
}

func main() {
    var v1 float64 = (0.0 - (-2.25 + 1e10))
    var v2 uint8 = 255
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]int)
    v2 = 238
    var v8 float32 = float32(float64(1.5))
    v9 := []int{128, 393}
    if false {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        var v10 bool = true
        _ = v10
    } else {
        var v11 uint8 = v2
        _ = v11
    }
    v12 := v4[4294967295]
    fmt.Println(v2, v12)
    var v13 int64 = 42
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 string = ("\n\t" + "foo bar")
        var v16 uint16 = 0
        fmt.Println(v2, v15)
        _ = fi14
        _ = v15
        _ = v16
    }
    if (!("" != "\n\t")) {
        for fi17 := 0; fi17 < 7; fi17++ {
            for fi18 := 0; fi18 < 12; fi18++ {
                if ((uint16(v2) != 42) && (true || false)) {
                    var v19 float32 = float32(v1)
                    _ = v19
                }
                v20 := GenS2{F0: "\n\t", F1: 42}
                _ = fi18
                _ = v20
            }
            _ = fi17
        }
        var v21 int64 = (v13 >> ((457) & 63))
        for fi22 := 0; fi22 < 1; fi22++ {
            fmt.Println(v13, v21, fi22)
            _ = fi22
        }
        _ = v21
    }
    var v23 uint8 = (v2 / (((v2 | v2)) | 1))
    var v24 float64 = ((v1 / (0.1 * 1.5)) + float64(((v8 / v8) + v8)))
    for fi25 := 0; fi25 < 19; fi25++ {
        switch v12 {
        case 2147483647:
            switch 128 {
            case 4294967295:
                v8 = float32(v1)
            case -2147483648:
                for fi26 := 0; fi26 < 11; fi26++ {
                    switch v12 {
                    case -353:
                        v27 := GenS3{F0: v2}
                        _ = v27
                    case -1:
                        v8 = 0.1
                    default:
                        fmt.Println(v2)
                    }
                    fmt.Println(v23, fi25)
                    _ = fi26
                }
            }
        default:
            for fi28 := 0; fi28 < 15; fi28++ {
                var v29 int = fi25
                v8 += 1.5
                _ = fi28
                _ = v29
            }
        }
        v30 := GenS4{F0: 65535, F1: 32767, F2: "\n\t"}
        _ = fi25
        _ = v30
    }
    if false {
        fmt.Println(v12, v8)
    }
    v13 = (v13 * v13)
    var v31 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v8
    _ = v9
}
