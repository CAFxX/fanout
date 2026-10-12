package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int8; F1 uint64; F2 uint }
type GenS3 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS4 struct { F0 int8; F1 int8; F2 string }
type GenS5 struct { F0 int64; F1 uint8; F2 bool }
type GenS6 struct { F0 bool; F1 string }
type GenS7 struct { F0 uint; F1 string }
type GenS8 struct { F0 int; F1 uint64 }
type GenS9 struct { F0 int16; F1 uint8 }
func genFn0(p0 int16, p1 uint8, p2 uint16) int64 {
    var v7 uint8 = 127
    for fi8 := 0; fi8 < 12; fi8++ {
        if (288 > uint(p2)) {
            if false {
                v9 := []int{fi8, fi8, fi8}
                v9 = append(v9, (fi8 | (-(fi8))))
                _ = v9
            } else {
                v10 := GenS1{F0: 564, F1: (((("a" + ("foo bar" + "a")) + ("foo bar" + ("hello" + "hello"))) + "hello") + ("hello" + "\n\t"))}
                if true {
                    var v11 string = ""
                    switch 4294967295 {
                    case 0:
                        if true {
                            if (true || true) {
                                v11 = v11
                                var v12 int8 = -128
                                _ = v12
                            }
                            v13 := []int{-442}
                            _ = v13
                        } else {
                            v14 := []int{}
                            for fi15 := 0; fi15 < 14; fi15++ {
                                v16 := GenS2{F0: -124, F1: 42, F2: 42}
                                v17 := GenS3{F0: 819, F1: p2, F2: 837}
                                var v18 int = (fi8 >> ((uint(p1)) & 63))
                                if len(v14) > 0 {
                                    v14[0] = fi8
                                }
                                _ = fi15
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                            var v19 uint32 = uint32(p1)
                            _ = v14
                            _ = v19
                        }
                    default:
                        v20 := []int{}
                        _ = v20
                    }
                    _ = v11
                } else {
                    if ((^(p1)) > (^(v7))) {
                        for fi21 := 0; fi21 < 3; fi21++ {
                            v22 := []int{(fi8 | (^(fi21)))}
                            var v23 uint = 4294967295
                            switch (fi21 ^ int(p0)) {
                            case 32768:
                                switch fi8 {
                                case -1:
                                    v23 = (v23 & (^(v23)))
                                }
                            case 32767:
                                var v24 bool = false
                                _ = v24
                            case 1:
                                fmt.Println(p0, fi21)
                            default:
                                v25 := GenS3{F0: 42, F1: (p2 >> (((v23 - (^(v23)))) & 63)), F2: 42}
                                _ = v25
                            }
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                        switch "foo bar" {
                        case "foo bar":
                            switch (-(fi8)) {
                            case 779:
                                v26 := []int{(fi8 ^ 2147483647)}
                                _ = v26
                            case -302:
                                var v27 int32 = -566
                                _ = v27
                            }
                        case "":
                            switch (-(fi8)) {
                            case -473:
                                v28 := GenS1{F0: p2, F1: "hello"}
                                _ = v28
                            case 32768:
                                v29 := []int{fi8, (fi8 + fi8), fi8, (fi8 / (((fi8 & 128)) | 1))}
                                _ = v29
                            case 548:
                                v30 := []int{2147483647, (fi8 / ((127) | 1))}
                                _ = v30
                            }
                        case "hello":
                            v31 := GenS4{F0: -1, F1: 0, F2: "foo bar"}
                            _ = v31
                        }
                    } else {
                        for fi32 := 0; fi32 < 17; fi32++ {
                            var v33 uint8 = 155
                            v34 := GenS5{F0: 42, F1: 0, F2: (!((true || (false || (p0 > 1))) || true))}
                            var v35 int32 = 0
                            _ = fi32
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                        switch (fi8 + fi8) {
                        case -952:
                            v36 := GenS6{F0: (0 <= -654), F1: ((("a" + ((("hello" + ("hello" + "foo bar")) + ((("foo bar" + "foo bar") + ("a" + "")) + "\n\t")) + "a")) + "hello") + ("a" + "hello"))}
                            _ = v36
                        case 32768:
                            fmt.Println(p1, p0, fi8)
                        case 42:
                            p1 = 174
                        }
                    }
                }
                _ = v10
            }
            fmt.Println(p2, p2)
            v37 := []int{fi8, (fi8 + (fi8 / ((9223372036854775807) | 1))), fi8, (fi8 / (((fi8 / (((fi8 / (((fi8 + (fi8 % (((fi8 - (fi8 ^ (fi8 * fi8)))) | 1)))) | 1))) | 1))) | 1))}
            _ = v37
        } else {
            v38 := []int{-1}
            v39 := GenS3{F0: 432, F1: 2, F2: 42}
            if len(v38) >= 2 {
                v38 = v38[1:len(v38)]
            }
            _ = v38
            _ = v39
        }
        _ = fi8
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    var v1 int16 = -32768
    var v2 int16 = (v1 | v1)
    var v3 float64 = 1e10
    var v4 int64 = 692
    v5 := []uint32{2147483647, 2147483648, 0}
    var v6 uint8 = 42
    if false {
        v40 := GenS2{F0: 42, F1: 714, F2: 1}
        if len(v5) > 0 {
            v5[0] = 42
        }
        v3 = (1.5 / 1.0)
        _ = v40
    }
    var v41 uint32 = 4294967295
    if true {
        for fi42 := 0; fi42 < 4; fi42++ {
            fmt.Println(v4)
            _ = fi42
        }
    } else {
        v2 += -510
        if true {
            v6 += 127
        }
        if false {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = len(v5)
            _ = len(v5)
        } else {
            switch "a" {
            case "foo bar":
                var v43 float64 = (v3 + (((v3 * v3) * ((v3 + 0.0) / v3)) - 1e10))
                _ = v43
            case "":
                fmt.Println(v2, v41)
            case "a":
                if (true || false) {
                    switch 516 {
                    case 497:
                        if ((v6 | 128) >= v6) {
                            for fi44 := 0; fi44 < 10; fi44++ {
                                if true {
                                    v2 = -736
                                    for fi45 := 0; fi45 < 17; fi45++ {
                                        if (float32((v3 / float64(1.5))) != float32(((((float64(0.0) - float64(1.5)) / 3.14159) + 3.14159) * (1e10 - 3.14159)))) {
                                            fmt.Println(v41)
                                        }
                                        v46 := GenS5{F0: 42, F1: v6, F2: (true && (42 != uint(v41)))}
                                        _ = fi45
                                        _ = v46
                                    }
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                if (1 > -391) {
                                    for fi47 := 0; fi47 < 8; fi47++ {
                                        v41 -= v41
                                        _ = fi47
                                    }
                                    for fi48 := 0; fi48 < 12; fi48++ {
                                        var v49 uint16 = 383
                                        var v50 float32 = 1.5
                                        _ = fi48
                                        _ = v49
                                        _ = v50
                                    }
                                } else {
                                    switch fi44 {
                                    case 144:
                                        _ = len(v5)
                                    case -2147483648:
                                        if len(v5) > 0 {
                                            v5[0] = v41
                                        }
                                    case 2147483647:
                                        v4 = (v4 - v4)
                                    }
                                }
                                _ = len(v5)
                                _ = fi44
                            }
                        } else {
                            for fi51 := 0; fi51 < 13; fi51++ {
                                fmt.Println(fi51, v1, fi51)
                                _ = fi51
                            }
                            var v52 int16 = 270
                            _ = v52
                        }
                    default:
                        v2 = 0
                    }
                    if false {
                        var v53 string = ""
                        _ = v53
                    }
                } else {
                    for fi54 := 0; fi54 < 9; fi54++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi54
                    }
                    v4 += (^(v4))
                }
            default:
                if (100 > 42) {
                    for fi55 := 0; fi55 < 15; fi55++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        if (fi55 == (fi55 & fi55)) {
                            if ((v1 << ((uint(v6)) & 63)) <= int16(v4)) {
                                var v56 bool = false
                                var v57 int64 = -34
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = v56
                                _ = v57
                            }
                        } else {
                            switch "\n\t" {
                            case "":
                                if (true || (fi55 > fi55)) {
                                    v2 = (v2 - 42)
                                    v2 = (v2 | v2)
                                    v1 = int16(v6)
                                } else {
                                    v6 -= uint8(v41)
                                    if ("foo bar" < (((("" + "\n\t") + (("hello" + "") + ("foo bar" + "a"))) + ("a" + "foo bar")) + "\n\t")) {
                                        var v58 int16 = (v1 & v1)
                                        _ = v58
                                    } else {
                                        _ = len(v5)
                                    }
                                    if ((^(v41)) == v41) {
                                        v4 += v4
                                    }
                                }
                            }
                        }
                        _ = fi55
                    }
                } else {
                    v3 = v3
                    if (float32(v3) <= 0.1) {
                        v41 += 4294967295
                        if true {
                            var v59 bool = (!(v41 == 2147483647))
                            v60 := GenS7{F0: 18446744073709551615, F1: "\n\t"}
                            var v61 int64 = (v4 / (((-(v4))) | 1))
                            _ = v59
                            _ = v60
                            _ = v61
                        } else {
                            v62 := GenS8{F0: 402, F1: 18446744073709551615}
                            fmt.Println(v6, v3)
                            _ = v62
                        }
                    }
                }
            }
        }
    }
    var v63 uint64 = 4294967295
    var v64 uint = 4294967295
    for fi65 := 0; fi65 < 3; fi65++ {
        v41 = (v41 ^ 1)
        for fi66 := 0; fi66 < 2; fi66++ {
            var v67 int16 = (v2 ^ v2)
            v68 := GenS5{F0: genFn0(v67, 226, 42), F1: (^(v6)), F2: ((v67 / (((v67 / ((848) | 1))) | 1)) != -1)}
            _ = fi66
            _ = v67
            _ = v68
        }
        _ = fi65
    }
    switch ("a" + "\n\t") {
    case "foo bar":
        if ((v64 < v64) || false) {
            v69 := GenS9{F0: (v1 % (((^(v1))) | 1)), F1: (v6 % ((0) | 1))}
            var v70 int8 = 1
            _ = v69
            _ = v70
        }
    default:
        v3 = v3
    }
    fmt.Println(v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v63
    _ = v64
}
