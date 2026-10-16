package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint32; F1 int64; F2 int }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 int64; F1 uint32; F2 bool }
type GenS6 struct { F0 int16; F1 int8; F2 uint }

func main() {
    var v1 uint = 4294967295
    v2 := []uint32{2147483648}
    v3 := make(map[uint32]bool)
    v1 = (v1 | (v1 & (v1 - (v1 | (^(v1))))))
    if false {
        var v4 string = "\n\t"
        v5 := GenS1{F0: 0, F1: (v1 >> (((^(v1))) & 63))}
        for fi6 := 0; fi6 < 1; fi6++ {
            for fi7 := 0; fi7 < 12; fi7++ {
                delete(v3, 134)
                fmt.Println(v1)
                _ = fi7
            }
            _ = fi6
        }
        _ = v4
        _ = v5
    } else {
        for fi8 := 0; fi8 < 19; fi8++ {
            for fi9 := 0; fi9 < 9; fi9++ {
                var v10 uint8 = 128
                v10 = v10
                _ = fi9
                _ = v10
            }
            v11 := GenS1{F0: -212, F1: 18446744073709551615}
            _ = fi8
            _ = v11
        }
    }
    if (!false) {
        for fi12 := 0; fi12 < 12; fi12++ {
            v1 -= v1
            for fi13 := 0; fi13 < 19; fi13++ {
                v1 = (v1 % (((v1 & (^(v1)))) | 1))
                _ = len(v3)
                switch fi13 {
                case -1:
                    fmt.Println(fi12, v1, fi13)
                case -9223372036854775808:
                    v1 -= 0
                default:
                    var v14 float64 = 1.5
                    v15 := GenS2{F0: int(v1), F1: 683}
                    _ = v14
                    _ = v15
                }
                _ = fi13
            }
            v1 += (v1 | v1)
            _ = fi12
        }
        var v16 uint64 = 18446744073709551615
        _ = v16
    }
    v1 -= v1
    v1 += 255
    if ((42 != 584) && true) {
        fmt.Println(v1)
    } else {
        for fi17 := 0; fi17 < 14; fi17++ {
            v1 -= 1
            v18 := GenS3{F0: 2, F1: 9223372036854775807, F2: 4294967295}
            _ = fi17
            _ = v18
        }
    }
    var v19 uint64 = 800
    v20 := v3[4294967295]
    var v21 int = -569
    for fi22 := 0; fi22 < 4; fi22++ {
        for fi23 := 0; fi23 < 3; fi23++ {
            var v24 int64 = 0
            v25 := GenS4{F0: 175, F1: 127}
            switch fi22 {
            case 2147483647:
                if (663 == 485) {
                    v2 = append(v2, 2147483648)
                    switch 103 {
                    case 0:
                        v26 := GenS5{F0: v24, F1: 767, F2: v20}
                        _ = v26
                    }
                }
            case -2147483648:
                var v27 bool = v20
                _ = v27
            default:
                if v20 {
                    switch "" {
                    case "a":
                        var v28 string = (("a" + (("foo bar" + "foo bar") + "")) + "")
                        _ = v28
                    }
                }
            }
            var v29 bool = false
            v19 = (v19 + (v19 / ((v19) | 1)))
            _ = fi23
            _ = v24
            _ = v25
            _ = v29
        }
        var v30 float32 = 0.1
        _ = fi22
        _ = v30
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        var v32 int32 = 519
        _ = fi31
        _ = v32
    }
    switch 161 {
    case -578:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        v33 := GenS6{F0: 42, F1: 0, F2: v1}
        _ = v33
    }
    v21 += v21
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
