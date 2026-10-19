package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint; F1 int64; F2 bool }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 int64 }
func genFn0(p0 int64, p1 int8, p2 bool) uint32 {
    v6 := GenS1{F0: 1, F1: 2}
    if (true || (32767 < 1000)) {
        if (p2 || true) {
            switch 42 {
            case 2147483647:
                var v7 int32 = -467
                _ = v7
            }
        } else {
            v8 := []int{32767, -2147483648, 32768}
            var v9 int8 = p1
            for fi10 := 0; fi10 < 15; fi10++ {
                for fi11 := 0; fi11 < 13; fi11++ {
                    var v12 uint64 = 0
                    if true {
                        var v13 int8 = (p1 | 100)
                        _ = v13
                    } else {
                        var v14 uint16 = 42
                        v9 = v9
                        _ = v14
                    }
                    var v15 int64 = (p0 ^ 42)
                    _ = fi11
                    _ = v12
                    _ = v15
                }
                _ = fi10
            }
            _ = v8
            _ = v9
        }
    } else {
        switch "a" {
        case "foo bar":
            switch 912 {
            case -423:
                for fi16 := 0; fi16 < 15; fi16++ {
                    p0 = (p0 >> ((1) & 63))
                    _ = fi16
                }
            case -96:
                fmt.Println(p2)
            case 863:
                var v17 uint16 = 1
                _ = v17
            default:
                var v18 int = -603
                _ = v18
            }
        case "hello":
            if p2 {
                v19 := []int{-1}
                var v20 bool = (false && (p1 == p1))
                for fi21 := 0; fi21 < 9; fi21++ {
                    var v22 int32 = 2147483647
                    for fi23 := 0; fi23 < 16; fi23++ {
                        v24 := GenS1{F0: 1, F1: 2}
                        _ = fi23
                        _ = v24
                    }
                    _ = fi21
                    _ = v22
                }
                _ = v19
                _ = v20
            } else {
                var v25 uint64 = 159
                _ = v25
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}

func main() {
    var v1 uint32 = 2147483647
    v2 := []int{}
    v3 := make(map[uint32]int)
    var v4 bool = (2147483647 >= -2147483648)
    var v5 uint = 65535
    v1 += (v1 % ((v1) | 1))
    var v26 int = 907
    v27 := v3[(v1 / (((v1 << ((v5) & 63))) | 1))]
    v28 := GenS2{F0: (v1 - (v1 / (((v1 << (((^(v5))) & 63))) | 1)))}
    for fi29 := 0; fi29 < 11; fi29++ {
        v26 = v27
        switch ("" + ("hello" + "hello")) {
        case "hello":
            switch (v26 & (v26 | (fi29 >> (((v5 << (((v5 * v5)) & 63))) & 63)))) {
            case 0:
                fmt.Println(v1, v27, fi29)
            case 127:
                switch (v26 & -2147483648) {
                case 863:
                    for fi30 := 0; fi30 < 17; fi30++ {
                        var v31 string = "\n\t"
                        v26 -= fi29
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v32 := GenS1{F0: -1, F1: 213}
                        switch v31 {
                        case "hello":
                            v4 = v4
                        }
                        v33 := GenS1{F0: -32768, F1: 2}
                        _ = fi30
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                case 627:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            case 690:
                var v34 float32 = float32((-2.25 * 0.1))
                _ = v34
            }
        case "":
            v27 += fi29
        case "a":
            for fi35 := 0; fi35 < 6; fi35++ {
                var v36 float32 = (float32(1.5) - ((3.14159 / 1e10) + (-2.25 + -2.25)))
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi35
                _ = v36
            }
        default:
            if v4 {
                var v37 bool = (v4 && v4)
                if (2 == 255) {
                    if (v4 || (false || (((v1 + (v1 % (((^(v1))) | 1))) < genFn0(-1, 127, (fi29 == 9223372036854775807))) && (-128 != -73)))) {
                        if v37 {
                            var v38 uint = (v5 & 255)
                            v3[2147483648] = 32767
                            _ = v38
                        }
                    } else {
                        fmt.Println(v37)
                    }
                    var v39 uint64 = 481
                    fmt.Println(v5)
                    _ = v39
                }
                _ = v37
            } else {
                v27 = 0
            }
        }
        for fi40 := 0; fi40 < 15; fi40++ {
            switch 1 {
            case -1:
                var v41 int32 = -1
                _ = v41
            case 1:
                switch v27 {
                case 492:
                    var v42 uint8 = 128
                    _ = v42
                case -9223372036854775808:
                    v27 = fi29
                case 2147483647:
                    fmt.Println(v5)
                default:
                    if false {
                        if (v4 && (!(255 == uint16(v1)))) {
                            v43 := GenS2{F0: (^(v1))}
                            fmt.Println(v4)
                            var v44 uint32 = v1
                            _ = v43
                            _ = v44
                        }
                    } else {
                        if (1 < 0) {
                            var v45 int = fi29
                            delete(v3, v1)
                            _ = v45
                        }
                    }
                }
            }
            delete(v3, v1)
            var v46 string = "a"
            _ = fi40
            _ = v46
        }
        _ = fi29
    }
    v4 = (v5 != v5)
    v3[(v1 * (v1 & (^(v1))))] = v27
    v47 := GenS1{F0: -32768, F1: 255}
    if ((v27 | v27) <= v27) {
        switch "hello" {
        case "foo bar":
            var v48 int8 = 100
            _ = v48
        default:
            switch "\n\t" {
            case "\n\t":
                _ = len(v3)
            case "":
                if false {
                    fmt.Println(v5, v1)
                    if false {
                        switch "foo bar" {
                        case "":
                            v2 = append(v2, -341)
                        default:
                            for fi49 := 0; fi49 < 12; fi49++ {
                                var v50 int32 = int32(v5)
                                v1 = 2147483647
                                _ = fi49
                                _ = v50
                            }
                        }
                    } else {
                        v4 = (42 == 175)
                        v51 := GenS3{F0: 127, F1: (v5 ^ v5)}
                        if ((1 == 77) && ((225 <= 1) || (65535 >= 0))) {
                            v52 := GenS4{F0: v5, F1: -539, F2: ((v4 && false) && ((v1 + genFn0(42, 18, v4)) < (^(v1))))}
                            switch (v27 ^ v27) {
                            case 127:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case -758:
                                _ = len(v2)
                            }
                            var v53 float64 = 0.0
                            _ = v52
                            _ = v53
                        } else {
                            for fi54 := 0; fi54 < 14; fi54++ {
                                fmt.Println(fi54)
                                v2 = append(v2, (v26 << ((667) & 63)))
                                v55 := GenS1{F0: 0, F1: 210}
                                _ = fi54
                                _ = v55
                            }
                            v56 := GenS5{F0: v5, F1: 159}
                            var v57 bool = v4
                            _ = v56
                            _ = v57
                        }
                        _ = v51
                    }
                    v26 += v26
                } else {
                    if (v4 && (-430 <= 0)) {
                        v58 := GenS6{F0: -1}
                        _ = v58
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v59 int32 = 0
                        _ = v59
                    }
                }
            }
        }
    } else {
        var v60 string = "hello"
        _ = v60
    }
    fmt.Println(v26)
    v4 = v4
    fmt.Println(v27, v26)
    fmt.Println(v27)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
