package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint; F1 int64; F2 uint64 }

func main() {
    v1 := []int{}
    v2 := []uint32{630}
    v3 := []int{32768, 2147483647, 2147483647, 684}
    var v4 int = 515
    var v5 uint32 = 2
    for fi6 := 0; fi6 < 11; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            v5 = (v5 >> ((4294967295) & 63))
            if (128 < 197) {
                v8 := GenS1{F0: fi7, F1: 493}
                v9 := GenS1{F0: -725, F1: -1}
                if false {
                    var v10 float64 = 0.1
                    _ = v10
                } else {
                    switch fi7 {
                    case 128:
                        v5 = 172
                    default:
                        v11 := GenS2{F0: 42, F1: 641}
                        _ = v11
                    }
                }
                _ = v8
                _ = v9
            }
            v12 := GenS3{F0: 0, F1: -1, F2: 2}
            v3 = append(v3, (fi7 + fi6))
            v4 += v4
            _ = fi7
            _ = v12
        }
        _ = fi6
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        v5 = 910
        if true {
            v5 = v5
        }
        var v14 uint32 = 42
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: v4, F1: 947}
    v16 := GenS3{F0: 541, F1: -787, F2: 0}
    v5 = v5
    v5 = v5
    if ((!false) && false) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v5)
    }
    v4 = (^(v4))
    v4 = (v4 ^ (v4 % ((157) | 1)))
    if len(v3) > 0 {
        v3[0] = (v4 << ((2) & 63))
    }
    v5 = (v5 + 928)
    v4 -= (v4 - v4)
    v4 += (v4 & 32767)
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 uint8 = 127
        _ = fi17
        _ = v18
    }
    if true {
        var v19 float32 = float32(-2.25)
        v4 -= -668
        _ = v19
    }
    if true {
        v4 -= (v4 & (v4 ^ (v4 << ((228) & 63))))
        for fi20 := 0; fi20 < 17; fi20++ {
            for fi21 := 0; fi21 < 19; fi21++ {
                var v22 int8 = 121
                v23 := GenS1{F0: 9223372036854775807, F1: 9223372036854775807}
                switch "a" {
                case "a":
                    var v24 float32 = 0.1
                    _ = v24
                default:
                    v22 += v22
                }
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = fi20
        }
        if true {
            if (!false) {
                switch "\n\t" {
                case "foo bar":
                    for fi25 := 0; fi25 < 13; fi25++ {
                        for fi26 := 0; fi26 < 7; fi26++ {
                            switch "a" {
                            case "hello":
                                v27 := GenS1{F0: v4, F1: -377}
                                _ = v27
                            case "":
                                var v28 int64 = -1
                                _ = v28
                            case "a":
                                if true {
                                    if (!false) {
                                        var v29 int64 = -830
                                        _ = v29
                                    }
                                }
                            default:
                                if len(v1) > 0 {
                                    v1[0] = fi26
                                }
                            }
                            v4 += fi26
                            v5 += v5
                            _ = fi26
                        }
                        _ = fi25
                    }
                default:
                    v30 := GenS3{F0: 255, F1: 42, F2: 4294967295}
                    _ = v30
                }
                v3 = append(v3, (v4 + (-(v4))))
                fmt.Println(v5, v5)
            }
        }
    } else {
        if (int32(v5) == -2147483648) {
            var v31 uint16 = 32767
            _ = v31
        } else {
            if len(v2) > 0 {
                v2[0] = v5
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v32 := GenS1{F0: v4, F1: int64(v4)}
            _ = v32
        }
        if len(v2) > 0 {
            v2[0] = v5
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
