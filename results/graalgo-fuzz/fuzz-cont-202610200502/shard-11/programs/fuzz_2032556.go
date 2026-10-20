package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int32; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint16; F1 string; F2 uint }
type GenS4 struct { F0 int; F1 uint8; F2 int64 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 uint8; F1 uint }

func main() {
    var v1 uint64 = 2
    var v2 uint8 = 128
    var v3 int32 = -1
    v4 := make(map[int]uint32)
    var v5 int8 = -128
    switch int(v2) {
    case -9223372036854775808:
        v6 := GenS1{F0: -9223372036854775808, F1: false}
        _ = v6
    case -509:
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 uint16 = uint16(v1)
            if ((true || false) || ((fi7 <= fi7) && (1 >= (v5 << ((233) & 63))))) {
                v9 := GenS2{F0: (v3 ^ (v3 >> ((65535) & 63))), F1: 0, F2: 1}
                for fi10 := 0; fi10 < 19; fi10++ {
                    switch fi7 {
                    case 9223372036854775807:
                        v2 += (v2 ^ 32)
                    case 127:
                        var v11 uint64 = (v1 / (((v1 ^ 0)) | 1))
                        _ = v11
                    }
                    v12 := GenS3{F0: v8, F1: "\n\t", F2: 255}
                    _ = fi10
                    _ = v12
                }
                for fi13 := 0; fi13 < 18; fi13++ {
                    fmt.Println(v5)
                    v14 := GenS2{F0: v3, F1: -455, F2: 2}
                    switch "hello" {
                    case "hello":
                        switch fi7 {
                        case 32768:
                            fmt.Println(v8, fi7, v1)
                        }
                    case "":
                        v15 := []int{fi7, (fi7 + int(v5))}
                        _ = v15
                    case "\n\t":
                        if ((36 > (^(v2))) || (-860 >= 1)) {
                            v8 = v8
                            var v16 uint64 = (v1 << ((2) & 63))
                            var v17 bool = (279 >= 454)
                            _ = v16
                            _ = v17
                        }
                    default:
                        v18 := GenS1{F0: (fi7 ^ fi13), F1: true}
                        fmt.Println(v3, fi7, v8)
                        _ = len(v4)
                        _ = v18
                    }
                    _ = fi13
                    _ = v14
                }
                _ = v9
            } else {
                for fi19 := 0; fi19 < 7; fi19++ {
                    v8 -= (v8 >> ((898) & 63))
                    _ = fi19
                }
                fmt.Println(v5)
            }
            for fi20 := 0; fi20 < 19; fi20++ {
                v21 := []int{(fi7 / (((fi20 & fi7)) | 1)), (fi7 ^ (-(fi7)))}
                v3 -= (^(v3))
                v5 = v5
                _ = fi20
                _ = v21
            }
            _ = fi7
            _ = v8
        }
    case -2147483648:
        switch "hello" {
        case "foo bar":
            v22 := GenS4{F0: -9223372036854775808, F1: (v2 | (v2 ^ (v2 << ((345) & 63)))), F2: 0}
            _ = v22
        case "a":
            v23 := []int{128, 9223372036854775807, 718}
            _ = v23
        case "\n\t":
            v24 := GenS4{F0: -9223372036854775808, F1: (v2 >> ((255) & 63)), F2: 781}
            _ = v24
        default:
            v5 -= 127
        }
    }
    v1 -= v1
    v1 -= v1
    v5 = (v5 * (v5 | v5))
    switch 526 {
    case 128:
        v25 := GenS1{F0: 1, F1: (!true)}
        _ = v25
    }
    v2 = 80
    var v26 float32 = -2.25
    v27 := GenS5{F0: 207, F1: -865}
    if (553 == -32768) {
        v1 = 0
        var v28 string = "hello"
        _ = v28
    } else {
        fmt.Println(v1, v3, v1)
        v29 := GenS6{F0: (^(v2)), F1: 12}
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
