package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 string; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int }

func main() {
    var v1 uint32 = 42
    var v2 uint = 255
    v3 := make(map[int]int)
    var v4 int8 = -128
    v5 := []uint32{(^(v1)), (v1 ^ v1), v1, 826}
    var v6 uint8 = 127
    if (!false) {
        v6 -= 127
        v3[-702] = 32767
    }
    if false {
        if (!true) {
            var v7 int16 = -445
            _ = v7
        } else {
            _ = len(v5)
            var v8 uint = (v2 | v2)
            for fi9 := 0; fi9 < 12; fi9++ {
                switch (^(fi9)) {
                case 128:
                    v8 -= (v2 & v2)
                case 42:
                    for fi10 := 0; fi10 < 19; fi10++ {
                        var v11 float64 = (float64(0.0) - float64(1.5))
                        _ = fi10
                        _ = v11
                    }
                default:
                    var v12 uint64 = uint64(v6)
                    _ = v12
                }
                var v13 uint8 = v6
                _ = fi9
                _ = v13
            }
            _ = v8
        }
    }
    v4 = v4
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := GenS1{F0: fi14, F1: (v1 << ((0) & 63))}
        v4 = 77
        v4 -= (v4 << ((431) & 63))
        _ = fi14
        _ = v15
    }
    v2 = (v2 << ((v2) & 63))
    _ = len(v3)
    v2 = 255
    fmt.Println(v4, v6)
    if (!(18446744073709551615 == 941)) {
        if (501 > 0) {
            for fi16 := 0; fi16 < 16; fi16++ {
                var v17 uint32 = (v1 >> ((v2) & 63))
                switch fi16 {
                case 0:
                    switch (fi16 % ((fi16) | 1)) {
                    case -546:
                        v3[fi16] = fi16
                    case -1:
                        v1 -= v17
                    case 574:
                        if false {
                            fmt.Println(v17, v2, v17)
                            var v18 uint8 = v6
                            v6 = v6
                            _ = v18
                        }
                    default:
                        v1 = 1
                    }
                case 42:
                    v1 = (^(v17))
                case 2147483647:
                    var v19 float64 = 1.5
                    _ = v19
                default:
                    if (65535 == (v2 ^ uint(v6))) {
                        var v20 float32 = ((0.0 * -2.25) - 0.0)
                        v21 := GenS2{F0: "\n\t", F1: 2, F2: 255}
                        switch "\n\t" {
                        case "hello":
                            var v22 uint32 = (v1 - 2)
                            _ = v22
                        case "\n\t":
                            delete(v3, fi16)
                        default:
                            var v23 uint = v2
                            _ = v23
                        }
                        _ = v20
                        _ = v21
                    } else {
                        if false {
                            _ = len(v3)
                            v24 := GenS2{F0: "", F1: 1, F2: v6}
                            _ = v24
                        }
                        var v25 uint16 = 42
                        _ = v25
                    }
                }
                _ = fi16
                _ = v17
            }
            v1 += v1
            v26 := GenS3{F0: -2147483648}
            _ = v26
        }
    }
    switch -381 {
    case -88:
        fmt.Println(v4, v1)
    case 2147483647:
        var v27 int16 = 1
        _ = v27
    case 998:
        v28 := v3[463]
        _ = v28
    default:
        v6 = 246
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        v1 += v1
        _ = fi29
    }
    v4 -= (v4 + (v4 + (v4 * (v4 + (v4 << (((v2 % (((v2 / (((v2 % ((249) | 1))) | 1))) | 1))) & 63))))))
    var v30 uint = (v2 / ((v2) | 1))
    fmt.Println(v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
