package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 string; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 int32; F2 int16 }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint32; F1 int8 }
type GenS9 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS10 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint32; F1 uint8 }
type GenS13 struct { F0 uint; F1 int }
func genFn0(p0 uint8, p1 uint32) int64 {
    v5 := GenS1{F0: p1}
    p0 = p0
    _ = p0
    _ = p1
    _ = v5
    return 1
}
func genFn1(p0 int64, p1 bool) uint64 {
    var v6 int16 = 763
    if p1 {
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 int32 = 0
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 int16 = 42
            _ = fi9
            _ = v10
        }
    }
    var v11 uint16 = 1
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    return 177
}

func main() {
    v1 := []uint32{}
    var v2 bool = (!true)
    v3 := []uint32{}
    var v4 uint = 65535
    switch "foo bar" {
    case "hello":
        v2 = (32767 < 32767)
    default:
        fmt.Println(v2, v4)
    }
    v1 = append(v1, 1)
    switch "hello" {
    case "\n\t":
        switch ("foo bar" + ("\n\t" + "hello")) {
        case "hello":
            var v12 bool = v2
            _ = v12
        case "foo bar":
            _ = len(v3)
        case "":
            for fi13 := 0; fi13 < 5; fi13++ {
                for fi14 := 0; fi14 < 17; fi14++ {
                    v15 := GenS2{F0: 127, F1: "\n\t", F2: 64}
                    v4 += 694
                    if (v2 && (255 <= 2)) {
                        if v2 {
                            v16 := GenS3{F0: 13}
                            _ = v16
                        } else {
                            switch (^(fi14)) {
                            case 2147483647:
                                var v17 int = fi14
                                _ = v17
                            default:
                                if (42 != (v4 / ((v4) | 1))) {
                                    switch "" {
                                    case "":
                                        v4 -= v4
                                    }
                                } else {
                                    v2 = v2
                                    v4 = 205
                                }
                            }
                            fmt.Println(fi13, v4, fi13)
                            if len(v3) > 0 {
                                v3[0] = 2
                            }
                        }
                        fmt.Println(fi14, fi13, v2)
                    }
                    v18 := GenS4{F0: 0, F1: -2147483648, F2: 1}
                    var v19 uint8 = 2
                    _ = fi14
                    _ = v15
                    _ = v18
                    _ = v19
                }
                _ = fi13
            }
        default:
            v20 := GenS5{F0: 42, F1: 42}
            _ = v20
        }
    case "":
        for fi21 := 0; fi21 < 4; fi21++ {
            fmt.Println(v4, fi21)
            switch "a" {
            case "a":
                var v22 int8 = 53
                _ = v22
            case "\n\t":
                v23 := GenS6{F0: 2, F1: 32767}
                _ = v23
            case "foo bar":
                for fi24 := 0; fi24 < 8; fi24++ {
                    fmt.Println(fi24, v2, fi21)
                    v4 = (v4 ^ v4)
                    v25 := GenS5{F0: 2, F1: 0}
                    _ = fi24
                    _ = v25
                }
            }
            switch "foo bar" {
            case "foo bar":
                for fi26 := 0; fi26 < 6; fi26++ {
                    switch ((("a" + "a") + ("\n\t" + "a")) + "foo bar") {
                    case "":
                        fmt.Println(fi26, fi21, v4)
                    case "hello":
                        var v27 uint8 = 42
                        _ = v27
                    }
                    if v2 {
                        v28 := GenS3{F0: 0}
                        v29 := GenS6{F0: 4294967295, F1: 125}
                        v3 = append(v3, 2)
                        _ = v28
                        _ = v29
                    }
                    _ = fi26
                }
            case "hello":
                var v30 uint = (v4 - v4)
                _ = v30
            }
            _ = fi21
        }
    case "a":
        var v31 string = "a"
        _ = v31
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        var v33 uint = v4
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 15; fi34++ {
        fmt.Println(fi34, fi34)
        var v35 float64 = 3.14159
        _ = fi34
        _ = v35
    }
    switch 42 {
    case 32768:
        var v36 int8 = -1
        _ = v36
    case -316:
        v4 = v4
    default:
        if true {
            v37 := GenS2{F0: 1, F1: "", F2: -813}
            if len(v1) > 0 {
                v1[0] = 2147483648
            }
            _ = v37
        } else {
            for fi38 := 0; fi38 < 14; fi38++ {
                var v39 int64 = genFn0(255, 328)
                _ = fi38
                _ = v39
            }
            v40 := GenS7{F0: -1}
            _ = v40
        }
    }
    v4 = v4
    if v2 {
        switch ("a" + (("\n\t" + "foo bar") + "hello")) {
        case "":
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
        default:
            v41 := GenS2{F0: -30, F1: "hello", F2: 9223372036854775807}
            _ = v41
        }
        var v42 uint32 = 835
        _ = v42
    } else {
        v43 := GenS4{F0: -520, F1: 0, F2: -32768}
        var v44 uint = (v4 ^ v4)
        var v45 uint8 = 255
        _ = v43
        _ = v44
        _ = v45
    }
    for fi46 := 0; fi46 < 12; fi46++ {
        if (1.5 >= 3.14159) {
            fmt.Println(fi46, fi46)
        }
        var v47 int8 = 0
        if false {
            v47 = v47
            for fi48 := 0; fi48 < 3; fi48++ {
                if v2 {
                    v49 := GenS1{F0: 4294967295}
                    if (1 <= 190) {
                        _ = len(v3)
                    }
                    if true {
                        var v50 int8 = v47
                        var v51 string = "a"
                        var v52 int64 = -497
                        _ = v50
                        _ = v51
                        _ = v52
                    } else {
                        if false {
                            v53 := GenS8{F0: 1, F1: (v47 / (((v47 + (^(v47)))) | 1))}
                            v2 = v2
                            v2 = true
                            _ = v53
                        } else {
                            switch (((("a" + "\n\t") + "hello") + "a") + "a") {
                            case "\n\t":
                                var v54 uint32 = 65535
                                _ = v54
                            default:
                                _ = len(v3)
                            }
                            var v55 uint16 = 65535
                            _ = v55
                        }
                    }
                    _ = v49
                }
                v56 := GenS3{F0: 2147483648}
                _ = fi48
                _ = v56
            }
            v2 = false
        } else {
            v57 := GenS4{F0: int(v47), F1: -442, F2: -65}
            for fi58 := 0; fi58 < 5; fi58++ {
                if len(v1) > 0 {
                    v1[0] = 65535
                }
                fmt.Println(v4, fi46, fi58)
                var v59 int = (fi46 % ((fi58) | 1))
                v60 := GenS2{F0: v47, F1: ("a" + "hello"), F2: -957}
                _ = len(v3)
                _ = fi58
                _ = v59
                _ = v60
            }
            _ = v57
        }
        _ = fi46
        _ = v47
    }
    if false {
        for fi61 := 0; fi61 < 2; fi61++ {
            var v62 float64 = -2.25
            _ = fi61
            _ = v62
        }
        var v63 uint32 = 2147483647
        switch 9223372036854775807 {
        case 746:
            var v64 int16 = 0
            _ = v64
        case -733:
            v65 := GenS8{F0: v63, F1: 107}
            _ = v65
        }
        _ = v63
    } else {
        for fi66 := 0; fi66 < 3; fi66++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if ((((!v2) || ((0.0 / (float32(1.5) * (((0.0 + (3.14159 - -2.25)) * 1e10) - 3.14159))) <= float32(float64(0.0)))) || (4294967295 != 470)) || (false && ((-2.25 == 1.5) || v2))) {
                v2 = (uint64(v4) == 2)
                v67 := GenS9{F0: 255, F1: -1, F2: 65535}
                _ = v67
            }
            _ = fi66
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi68 := 0; fi68 < 5; fi68++ {
            var v69 int8 = 0
            _ = fi68
            _ = v69
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v3) > 0 {
        v3[0] = 0
    }
    fmt.Println(v4)
    if (!(!v2)) {
        if v2 {
            switch 128 {
            case -503:
                if v2 {
                    var v70 uint64 = genFn1(9223372036854775807, false)
                    var v71 float32 = 0.1
                    _ = v70
                    _ = v71
                } else {
                    for fi72 := 0; fi72 < 11; fi72++ {
                        switch (^(fi72)) {
                        case 32768:
                            for fi73 := 0; fi73 < 5; fi73++ {
                                var v74 int64 = 829
                                if len(v1) > 0 {
                                    v1[0] = 42
                                }
                                _ = fi73
                                _ = v74
                            }
                        case 0:
                            if v2 {
                                fmt.Println(v2, fi72, v4)
                                var v75 uint32 = 0
                                var v76 uint16 = 65535
                                _ = v75
                                _ = v76
                            } else {
                                switch "" {
                                case "foo bar":
                                    if true {
                                        var v77 int16 = -32768
                                        switch 0 {
                                        case 137:
                                            if true {
                                                if ("\n\t" <= "\n\t") {
                                                    switch ("" + ((("foo bar" + "\n\t") + "foo bar") + ("" + "hello"))) {
                                                    case "a":
                                                        var v78 float32 = (1e10 + -2.25)
                                                        _ = v78
                                                    }
                                                    v79 := GenS10{F0: 24, F1: v77, F2: 871}
                                                    _ = v79
                                                }
                                                var v80 float64 = 0.1
                                                for fi81 := 0; fi81 < 17; fi81++ {
                                                    v82 := GenS11{F0: 34}
                                                    v2 = v2
                                                    v4 = 414
                                                    _ = fi81
                                                    _ = v82
                                                }
                                                _ = v80
                                            }
                                        default:
                                            v77 -= 0
                                        }
                                        _ = v77
                                    }
                                case "\n\t":
                                    var v83 int16 = -550
                                    _ = v83
                                case "a":
                                    for fi84 := 0; fi84 < 19; fi84++ {
                                        v4 -= v4
                                        _ = fi84
                                    }
                                default:
                                    for fi85 := 0; fi85 < 16; fi85++ {
                                        v86 := GenS1{F0: 2147483648}
                                        _ = fi85
                                        _ = v86
                                    }
                                }
                            }
                        case 649:
                            fmt.Println(fi72, fi72, v2)
                        }
                        fmt.Println(v2)
                        _ = fi72
                    }
                    if len(v3) > 0 {
                        v3[0] = 187
                    }
                }
            case 127:
                var v87 uint64 = 2
                _ = v87
            case 4294967295:
                if (1.5 < 1e10) {
                    fmt.Println(v2)
                } else {
                    v2 = ((42 == 0) && v2)
                    var v88 int = 2147483647
                    v89 := GenS12{F0: 65535, F1: 127}
                    _ = v88
                    _ = v89
                }
            }
            if v2 {
                switch "\n\t" {
                case "hello":
                    v2 = v2
                default:
                    v3 = append(v3, 1)
                }
            }
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v90 uint32 = 0
    for fi91 := 0; fi91 < 1; fi91++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi92 := 0; fi92 < 3; fi92++ {
            if len(v3) > 0 {
                v3[0] = (v90 & v90)
            }
            var v93 int16 = 32767
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi92
            _ = v93
        }
        switch "" {
        case "a":
            for fi94 := 0; fi94 < 4; fi94++ {
                v95 := GenS3{F0: 4294967295}
                v96 := GenS9{F0: 0, F1: -1, F2: (v90 - v90)}
                var v97 string = "a"
                fmt.Println(fi91, v90, fi91)
                fmt.Println(fi91, v2)
                _ = fi94
                _ = v95
                _ = v96
                _ = v97
            }
        case "":
            for fi98 := 0; fi98 < 6; fi98++ {
                var v99 int16 = -257
                _ = fi98
                _ = v99
            }
        default:
            for fi100 := 0; fi100 < 14; fi100++ {
                switch fi100 {
                case 1:
                    v101 := GenS13{F0: (v4 | v4), F1: int(v90)}
                    _ = v101
                default:
                    if v2 {
                        v90 = (v90 * v90)
                        var v102 uint8 = 39
                        _ = v102
                    } else {
                        v2 = true
                    }
                }
                v2 = true
                _ = fi100
            }
        }
        _ = fi91
    }
    for fi103 := 0; fi103 < 5; fi103++ {
        for fi104 := 0; fi104 < 3; fi104++ {
            v2 = v2
            if len(v3) > 0 {
                v3[0] = v90
            }
            _ = fi104
        }
        v4 -= (v4 & (v4 * (v4 | 337)))
        _ = fi103
    }
    v105 := GenS10{F0: 255, F1: -453, F2: (v90 << ((v4) & 63))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v105
    _ = v2
    _ = v3
    _ = v4
    _ = v90
}
