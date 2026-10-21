package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16; F1 uint16 }

func main() {
    v1 := []uint32{336, 42, 585}
    var v2 int32 = 0
    var v3 uint32 = 129
    var v4 uint = 18446744073709551615
    var v5 int16 = 0
    v6 := make(map[int]int)
    v3 = (v3 - (v3 % ((v3) | 1)))
    var v7 uint = v4
    v3 = v3
    if (!(!true)) {
        v3 = (v3 * v3)
        v6[1] = -18
    }
    var v8 int64 = -9223372036854775808
    if len(v1) > 0 {
        v1[0] = (^(v3))
    }
    if false {
        if false {
            delete(v6, int(v2))
            switch int(v5) {
            case 128:
                var v9 uint8 = 1
                _ = v9
            case -654:
                v10 := GenS1{F0: (-(v8)), F1: (v4 % (((^(v4))) | 1))}
                _ = v10
            default:
                if (253 == 127) {
                    for fi11 := 0; fi11 < 17; fi11++ {
                        var v12 uint8 = 189
                        var v13 float32 = ((0.1 / (0.1 * float32((0.1 / (1.5 + -2.25))))) - ((0.1 - float32(((float64((-2.25 * (3.14159 / 0.1))) - ((0.1 - 0.0) / 0.1)) / 1.0))) + 3.14159))
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    v14 := GenS2{F0: 218, F1: false}
                    _ = v14
                }
            }
            v2 = (v2 - v2)
        }
        for fi15 := 0; fi15 < 12; fi15++ {
            switch "\n\t" {
            case "foo bar":
                var v16 bool = (v4 != v7)
                _ = v16
            case "a":
                v2 += 660
            default:
                v1 = append(v1, v3)
            }
            if (false || (!false)) {
                if false {
                    fmt.Println(v5)
                    v3 += (v3 >> ((2) & 63))
                    for fi17 := 0; fi17 < 12; fi17++ {
                        fmt.Println(v7, v3)
                        _ = fi17
                    }
                } else {
                    v5 -= (-(v5))
                }
                v18 := GenS1{F0: (^(v8)), F1: v4}
                v19 := GenS3{F0: (v5 << (((v7 | (v4 << (((v7 % (((v7 | v7)) | 1))) & 63)))) & 63))}
                _ = v18
                _ = v19
            }
            v5 = -369
            _ = fi15
        }
    }
    var v20 bool = true
    for fi21 := 0; fi21 < 8; fi21++ {
        v5 -= 1000
        fmt.Println(v5, v8, v3)
        v22 := GenS4{F0: "a"}
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        v8 += v8
        var v24 float32 = 1e10
        v25 := GenS4{F0: ""}
        _ = fi23
        _ = v24
        _ = v25
    }
    v4 -= (^(v7))
    if len(v1) > 0 {
        v1[0] = v3
    }
    v1 = append(v1, (v3 >> ((307) & 63)))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (v3 < 65535) {
        switch "\n\t" {
        case "hello":
            switch "" {
            case "a":
                delete(v6, -735)
            default:
                var v26 float64 = (float64(-2.25) + 1.5)
                _ = v26
            }
        default:
            v27 := GenS5{F0: (v5 % ((1000) | 1)), F1: 0}
            _ = v27
        }
    } else {
        var v28 uint = v4
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
