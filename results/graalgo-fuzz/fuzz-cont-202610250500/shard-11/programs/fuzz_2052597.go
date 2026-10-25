package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }
func genFn0(p0 uint, p1 int16, p2 uint32) uint {
    var v4 uint32 = (p2 + (p2 % ((42) | 1)))
    if true {
        v4 = 1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p0 ^ p0)
}

func main() {
    var v1 int = 127
    var v2 bool = true
    var v3 uint8 = 244
    v1 = (v1 - v1)
    var v5 uint8 = v3
    if v2 {
        v6 := []int{v1, (v1 ^ v1)}
        if len(v6) > 0 {
            v6[0] = (v1 / (((v1 & v1)) | 1))
        }
        _ = v6
    }
    v7 := []int{v1, (v1 >> ((65535) & 63))}
    var v8 uint64 = 302
    switch (-(v1)) {
    case 1:
        for fi9 := 0; fi9 < 14; fi9++ {
            switch v1 {
            case 129:
                v10 := GenS1{F0: 42, F1: (v8 + (^(v8)))}
                _ = v10
            case 2147483647:
                v5 = v3
            }
            if v2 {
                for fi11 := 0; fi11 < 12; fi11++ {
                    var v12 uint = 608
                    v1 = v1
                    _ = fi11
                    _ = v12
                }
                fmt.Println(v1, v3, v3)
                v7 = append(v7, 0)
            } else {
                fmt.Println(v1, v8, v5)
                v7 = append(v7, fi9)
                v13 := GenS2{F0: v3}
                _ = v13
            }
            _ = fi9
        }
    default:
        for fi14 := 0; fi14 < 4; fi14++ {
            if (v2 || (0.1 == 1e10)) {
                v15 := GenS2{F0: v5}
                switch (fi14 + v1) {
                case -486:
                    _ = len(v7)
                case 1:
                    if (!v2) {
                        var v16 float32 = (-2.25 + 0.1)
                        _ = v16
                    } else {
                        v7 = append(v7, v1)
                    }
                case -524:
                    v17 := GenS3{F0: uint(v5)}
                    _ = v17
                default:
                    if (-115 >= 90) {
                        _ = len(v7)
                        fmt.Println(v2, fi14)
                    }
                }
                v7 = append(v7, (fi14 - (-(fi14))))
                _ = v15
            } else {
                for fi18 := 0; fi18 < 19; fi18++ {
                    switch "foo bar" {
                    case "hello":
                        var v19 uint8 = v5
                        _ = v19
                    case "":
                        switch (v1 ^ 1) {
                        case 9223372036854775807:
                            fmt.Println(fi18, v8, v5)
                        case -137:
                            v5 -= v3
                        case -2147483648:
                            _ = len(v7)
                        default:
                            v8 = (v8 - (^(v8)))
                        }
                    case "a":
                        fmt.Println(v5, v3)
                    default:
                        if ((int64(fi18) >= 9223372036854775807) || (v2 || v2)) {
                            v5 = (v5 - (v5 | 128))
                        }
                    }
                    _ = fi18
                }
                v2 = (-64 != 1)
                fmt.Println(fi14, v5)
            }
            _ = fi14
        }
    }
    v5 -= (v5 / ((v5) | 1))
    v1 = (-(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v7
    _ = v8
}
