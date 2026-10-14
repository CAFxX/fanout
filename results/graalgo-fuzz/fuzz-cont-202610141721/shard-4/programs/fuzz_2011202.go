package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int32; F1 int; F2 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32; F1 uint64; F2 int }
type GenS6 struct { F0 uint8; F1 int64; F2 uint64 }
func genFn0(p0 int64, p1 uint32) int32 {
    var v4 int16 = -32768
    switch 2147483647 {
    case 127:
        v5 := GenS1{F0: -615, F1: p0}
        _ = v5
    }
    v4 = v4
    if (!false) {
        var v6 string = "hello"
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return 2147483647
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 0
    var v3 uint8 = 0
    v3 -= 255
    v7 := []int{1}
    v8 := GenS2{F0: (224 <= -9223372036854775808), F1: false}
    v2 += v2
    var v9 int16 = 1
    var v10 int64 = 0
    switch 32767 {
    case 1:
        for fi11 := 0; fi11 < 20; fi11++ {
            if true {
                v2 += v2
                if false {
                    for fi12 := 0; fi12 < 16; fi12++ {
                        v2 += 42
                        v13 := GenS3{F0: int32(v3), F1: (fi11 ^ (^(fi11))), F2: 1}
                        _ = len(v7)
                        _ = fi12
                        _ = v13
                    }
                    v14 := GenS4{F0: 0}
                    var v15 float32 = -2.25
                    _ = v14
                    _ = v15
                }
                for fi16 := 0; fi16 < 1; fi16++ {
                    v9 -= (v9 | (v9 & (v9 >> ((uint(v3)) & 63))))
                    v3 += v3
                    _ = fi16
                }
            }
            if (!((!(393 != 42)) || (false || (false && false)))) {
                v17 := GenS5{F0: 696, F1: v2, F2: (^(fi11))}
                if true {
                    for fi18 := 0; fi18 < 3; fi18++ {
                        switch "a" {
                        case "foo bar":
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        case "\n\t":
                            v2 = v2
                        default:
                            v10 = (v10 * (v10 * (v10 & (v10 & 485))))
                        }
                        fmt.Println(fi18, fi11, v10)
                        _ = len(v7)
                        _ = fi18
                    }
                    v3 += v3
                }
                _ = v17
            } else {
                if false {
                    v19 := v1[308]
                    v10 = v10
                    var v20 int64 = (v10 + (v10 % ((-96) | 1)))
                    _ = v19
                    _ = v20
                }
            }
            _ = fi11
        }
    default:
        var v21 float64 = float64(((1.5 * 1.5) + (1.5 - 1.5)))
        _ = v21
    }
    fmt.Println(v9, v9)
    switch 0 {
    case 127:
        for fi22 := 0; fi22 < 14; fi22++ {
            v2 = (^(v2))
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi22
        }
    case 42:
        var v23 int = 32768
        _ = v23
    }
    v24 := GenS6{F0: v3, F1: (v10 - (v10 / (((v10 | v10)) | 1))), F2: v2}
    var v25 int32 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
