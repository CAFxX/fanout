package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int64; F2 string }
type GenS3 struct { F0 uint8; F1 uint16; F2 string }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 int32; F1 int64; F2 int16 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 int, p1 uint16) uint8 {
    v7 := []int{}
    if len(v7) > 0 {
        v7[0] = 32768
    }
    if (0 == 0) {
        fmt.Println(p1, p1)
    } else {
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := GenS1{F0: 583}
            switch "" {
            case "\n\t":
                p1 = (p1 - p1)
            }
            for fi10 := 0; fi10 < 16; fi10++ {
                v7 = append(v7, -591)
                v11 := GenS1{F0: 42}
                _ = fi10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
        p1 -= p1
        v12 := GenS1{F0: 4294967295}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    return 128
}
func genFn1(p0 int16, p1 bool) uint16 {
    var v13 uint = 65535
    v14 := []int{-303, -40, -335, 49}
    if len(v14) > 0 {
        v14[0] = 42
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 326
}

func main() {
    v1 := []string{}
    var v2 uint16 = 42
    var v3 uint8 = 128
    v4 := make(map[int]uint32)
    var v5 float32 = -2.25
    v6 := []int{1, -2147483648}
    _ = len(v4)
    delete(v4, 9223372036854775807)
    v2 = (^(v2))
    v3 = (v3 / ((uint8(v2)) | 1))
    v3 = genFn0(97, 255)
    v3 -= v3
    var v15 uint32 = 0
    delete(v4, 42)
    v2 += v2
    for fi16 := 0; fi16 < 14; fi16++ {
        switch fi16 {
        case 32767:
            var v17 bool = false
            _ = v17
        case -877:
            if true {
                fmt.Println(v15, fi16)
                var v18 bool = (v2 != 768)
                _ = v18
            } else {
                fmt.Println(fi16)
                switch "\n\t" {
                case "hello":
                    fmt.Println(fi16, v15, v15)
                default:
                    fmt.Println(v2, v5, v2)
                }
                var v19 int64 = 917
                _ = v19
            }
        }
        for fi20 := 0; fi20 < 16; fi20++ {
            switch fi16 {
            case 32768:
                var v21 int64 = -793
                _ = v21
            case 2147483647:
                v3 = v3
                var v22 int = fi16
                _ = v22
            case 32767:
                if true {
                    var v23 float32 = float32(0.0)
                    v2 += (v2 * 205)
                    var v24 uint8 = v3
                    v25 := GenS1{F0: 42}
                    _ = v23
                    _ = v24
                    _ = v25
                }
            default:
                var v26 uint32 = 2147483648
                _ = v26
            }
            switch "\n\t" {
            case "a":
                v27 := GenS1{F0: 843}
                _ = v27
            case "\n\t":
                v2 = (v2 + genFn1(32767, false))
            case "hello":
                fmt.Println(v15, v5)
            default:
                v28 := GenS2{F0: v3, F1: 316, F2: "hello"}
                _ = v28
            }
            v3 += (v3 + (v3 * v3))
            _ = fi20
        }
        switch (fi16 >> ((255) & 63)) {
        case 32768:
            switch (fi16 / (((^(fi16))) | 1)) {
            case 42:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        case 42:
            v5 = v5
        }
        _ = fi16
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        v3 -= v3
        _ = fi29
    }
    var v30 uint8 = (v3 | (v3 & v3))
    if false {
        v31 := GenS1{F0: 18446744073709551615}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v15 -= (^(v15))
        _ = v31
    } else {
        v32 := v4[4294967295]
        for fi33 := 0; fi33 < 17; fi33++ {
            fmt.Println(v3)
            switch "foo bar" {
            case "":
                v1 = append(v1, (("foo bar" + ("" + "\n\t")) + "\n\t"))
            default:
                for fi34 := 0; fi34 < 6; fi34++ {
                    v5 += v5
                    switch -9223372036854775808 {
                    case 0:
                        v35 := GenS3{F0: 42, F1: (v2 << ((1) & 63)), F2: "a"}
                        _ = v35
                    case -1:
                        if len(v1) > 0 {
                            v1[0] = (("a" + ("\n\t" + "foo bar")) + ("" + "\n\t"))
                        }
                    case -9223372036854775808:
                        var v36 int64 = -93
                        _ = v36
                    }
                    var v37 uint32 = v32
                    _ = fi34
                    _ = v37
                }
            }
            _ = fi33
        }
        if false {
            switch "foo bar" {
            case "a":
                var v38 bool = (51 > 42)
                _ = v38
            case "hello":
                if true {
                    if false {
                        v39 := GenS1{F0: 0}
                        _ = v39
                    } else {
                        var v40 string = "\n\t"
                        var v41 uint64 = 0
                        if true {
                            fmt.Println(v15, v32)
                            v42 := GenS2{F0: 42, F1: int64(v41), F2: v40}
                            _ = v42
                        } else {
                            if false {
                                v41 += (v41 << ((65535) & 63))
                                if false {
                                    v43 := GenS2{F0: 127, F1: -9223372036854775808, F2: ((v40 + v40) + ("\n\t" + ""))}
                                    if (false || (-989 < 658)) {
                                        fmt.Println(v5, v5)
                                        fmt.Println(v40)
                                        switch ((v40 + (v40 + "")) + (v40 + "foo bar")) {
                                        case "":
                                            v44 := GenS4{F0: 746, F1: true}
                                            _ = v44
                                        case "a":
                                            var v45 int = 32768
                                            _ = v45
                                        default:
                                            v3 += v30
                                        }
                                    } else {
                                        if (!(2 == 18446744073709551615)) {
                                            v5 += (v5 - (v5 / ((float32(3.14159) * ((v5 - (v5 / 1.0)) + ((0.0 / 1.5) / (v5 * v5)))) - v5)))
                                            delete(v4, 72)
                                        } else {
                                            switch 128 {
                                            case -2147483648:
                                                if false {
                                                    v46 := GenS5{F0: -304, F1: 1, F2: -571}
                                                    v3 += (^(v30))
                                                    _ = len(v1)
                                                    _ = v46
                                                }
                                            case 1:
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            case 127:
                                                v1 = append(v1, v40)
                                            default:
                                                _ = len(v4)
                                            }
                                        }
                                        switch 772 {
                                        case 869:
                                            switch -288 {
                                            case 2147483647:
                                                var v47 float32 = -2.25
                                                _ = v47
                                            case 4294967295:
                                                if false {
                                                    for fi48 := 0; fi48 < 10; fi48++ {
                                                        fmt.Println(v15, v32)
                                                        _ = fi48
                                                    }
                                                    fmt.Println(v40)
                                                }
                                            }
                                        default:
                                            switch "hello" {
                                            case "hello":
                                                fmt.Println(v32, v41)
                                            case "a":
                                                v32 -= 42
                                            case "\n\t":
                                                v3 += (v30 - v30)
                                            default:
                                                v2 = v2
                                            }
                                        }
                                        var v49 uint32 = (v32 ^ (^(v32)))
                                        _ = v49
                                    }
                                    _ = v43
                                } else {
                                    if ((v15 & 145) == 986) {
                                        var v50 uint16 = genFn1(-958, (!(!((genFn0(2147483647, v2) < 42) || (319 >= -1)))))
                                        v51 := GenS6{F0: 65535}
                                        _ = v50
                                        _ = v51
                                    } else {
                                        var v52 uint8 = (^(v30))
                                        _ = v52
                                    }
                                }
                            } else {
                                if true {
                                    v40 = ""
                                } else {
                                    v5 += v5
                                    v4[32767] = 2147483648
                                    var v53 uint64 = v41
                                    _ = v53
                                }
                            }
                        }
                        _ = v40
                        _ = v41
                    }
                    for fi54 := 0; fi54 < 20; fi54++ {
                        if false {
                            var v55 float32 = ((v5 / v5) + v5)
                            for fi56 := 0; fi56 < 16; fi56++ {
                                delete(v4, fi56)
                                _ = len(v1)
                                if (float32((0.0 / (0.1 / 1.0))) < float32(0.0)) {
                                    if (32768 < (fi56 | fi54)) {
                                        v2 += uint16(v3)
                                        _ = len(v1)
                                    }
                                    v2 += 65535
                                }
                                _ = fi56
                            }
                            _ = v55
                        }
                        if true {
                            fmt.Println(v2, v3)
                            if ((!false) && (false || (0.0 >= 0.0))) {
                                switch "foo bar" {
                                case "a":
                                    v6 = append(v6, -439)
                                case "\n\t":
                                    switch (fi54 - fi54) {
                                    case 32768:
                                        if (false || (!false)) {
                                            for fi57 := 0; fi57 < 1; fi57++ {
                                                if (276 >= 42) {
                                                    var v58 uint = 2
                                                    _ = v58
                                                }
                                                _ = fi57
                                            }
                                            v59 := GenS1{F0: 127}
                                            _ = v59
                                        } else {
                                            for fi60 := 0; fi60 < 12; fi60++ {
                                                var v61 int = (fi60 + fi60)
                                                fmt.Println(v5)
                                                _ = fi60
                                                _ = v61
                                            }
                                        }
                                    case 453:
                                        v62 := GenS4{F0: fi54, F1: (uint(v2) <= 134)}
                                        _ = v62
                                    default:
                                        var v63 int16 = -32768
                                        _ = v63
                                    }
                                case "":
                                    var v64 uint64 = 983
                                    _ = v64
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                                var v65 uint16 = v2
                                _ = v65
                            } else {
                                switch ("foo bar" + "\n\t") {
                                case "":
                                    var v66 uint32 = v32
                                    _ = v66
                                case "\n\t":
                                    for fi67 := 0; fi67 < 15; fi67++ {
                                        var v68 string = ""
                                        v3 += v3
                                        _ = fi67
                                        _ = v68
                                    }
                                case "hello":
                                    if (!(42 > 1000)) {
                                        if (!(18446744073709551615 < 180)) {
                                            v2 = (v2 * 0)
                                            if true {
                                                v32 = v32
                                                var v69 int64 = -597
                                                _ = v69
                                            }
                                            for fi70 := 0; fi70 < 4; fi70++ {
                                                fmt.Println(v32, fi70)
                                                if ("" < "\n\t") {
                                                    v30 = v3
                                                    v71 := v4[(-(fi54))]
                                                    if true {
                                                        v6 = append(v6, fi54)
                                                        delete(v4, fi70)
                                                    }
                                                    _ = v71
                                                }
                                                v32 += (^(v15))
                                                _ = fi70
                                            }
                                        }
                                        var v72 int32 = 42
                                        _ = v72
                                    }
                                }
                                if false {
                                    switch int(v3) {
                                    case -1:
                                        _ = len(v6)
                                    }
                                    delete(v4, int(v32))
                                } else {
                                    for fi73 := 0; fi73 < 17; fi73++ {
                                        v74 := v4[fi54]
                                        var v75 int32 = 42
                                        _ = fi73
                                        _ = v74
                                        _ = v75
                                    }
                                    fmt.Println(v30)
                                }
                                fmt.Println(v2, v30)
                            }
                        }
                        var v76 uint64 = 1
                        _ = fi54
                        _ = v76
                    }
                    v15 = (^(v15))
                }
            case "":
                var v77 uint = 4294967295
                _ = v77
            }
            if true {
                var v78 int64 = 42
                var v79 int16 = -1
                v30 = (v30 ^ genFn0(-1, 1))
                _ = v78
                _ = v79
            }
            var v80 float32 = 1.5
            _ = v80
        } else {
            if (881 >= 42) {
                if false {
                    _ = len(v6)
                    _ = len(v4)
                    v2 = genFn1(0, false)
                }
                fmt.Println(v15, v2, v15)
            } else {
                switch ("" + "\n\t") {
                case "foo bar":
                    var v81 int32 = -2147483648
                    _ = v81
                case "\n\t":
                    var v82 bool = (false && (uint8(v32) < (^(v30))))
                    _ = v82
                default:
                    v2 = (v2 << ((65535) & 63))
                }
            }
        }
        _ = v32
    }
    if (0 != 926) {
        fmt.Println(v15, v3)
        v15 -= 2147483647
    } else {
        var v83 float64 = (-2.25 / (3.14159 / 0.1))
        if true {
            v83 = (v83 * (v83 - (((1.5 - (0.0 / float64((v5 * 1e10)))) / float64(v5)) / v83)))
            var v84 int16 = 509
            _ = v84
        }
        if false {
            v2 += genFn1(1000, (false || true))
            var v85 int64 = 0
            v30 = 246
            _ = v85
        }
        _ = v83
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
