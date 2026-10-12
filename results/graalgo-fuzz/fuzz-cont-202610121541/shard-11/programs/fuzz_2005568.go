package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int8; F1 uint64; F2 int32 }
type GenS5 struct { F0 uint16; F1 int8; F2 int64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS9 struct { F0 int8; F1 uint8 }
type GenS10 struct { F0 int8; F1 uint64 }
type GenS11 struct { F0 uint8; F1 int8; F2 int8 }
type GenS12 struct { F0 int16; F1 int32 }
type GenS13 struct { F0 bool; F1 int; F2 string }
func genFn0(p0 int32) uint64 {
    v6 := GenS1{F0: 546}
    v7 := []int{-9223372036854775808, 2147483647}
    var v8 int = 128
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 748
}
func genFn1(p0 uint8, p1 uint8) uint {
    if false {
        var v9 int8 = 62
        _ = v9
    }
    switch "foo bar" {
    case "\n\t":
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 float64 = 1e10
            _ = fi10
            _ = v11
        }
    case "a":
        if (153 > -478) {
            if (9223372036854775807 != 101) {
                if false {
                    var v12 int32 = 42
                    v13 := GenS2{F0: 42}
                    for fi14 := 0; fi14 < 19; fi14++ {
                        if false {
                            fmt.Println(v12, p0)
                            switch "hello" {
                            case "a":
                                var v15 int = fi14
                                _ = v15
                            case "\n\t":
                                for fi16 := 0; fi16 < 2; fi16++ {
                                    v12 = v12
                                    _ = fi16
                                }
                            default:
                                if true {
                                    for fi17 := 0; fi17 < 19; fi17++ {
                                        v12 += (v12 | (v12 % (((v12 % ((v12) | 1))) | 1)))
                                        _ = fi17
                                    }
                                    if false {
                                        p0 -= p0
                                        v18 := GenS3{F0: ((0.1 * float32(0.0)) != (-2.25 / float32(1.5))), F1: (^(v12))}
                                        for fi19 := 0; fi19 < 2; fi19++ {
                                            if false {
                                                v20 := []int{4294967295, (fi19 & fi14), fi19, fi14}
                                                _ = v20
                                            } else {
                                                v12 = (v12 ^ v12)
                                            }
                                            p1 += p1
                                            v21 := []int{(fi14 % (((fi19 % ((fi14) | 1))) | 1)), fi14, -2147483648}
                                            if true {
                                                switch "hello" {
                                                case "hello":
                                                    var v22 uint = 1000
                                                    _ = v22
                                                case "\n\t":
                                                    p1 += p1
                                                }
                                                v12 -= v12
                                            } else {
                                                v23 := GenS3{F0: ((false || (p0 < p0)) && true), F1: (^(v12))}
                                                switch "hello" {
                                                case "hello":
                                                    fmt.Println(p0, fi14, v12)
                                                case "foo bar":
                                                    switch (fi19 / ((fi19) | 1)) {
                                                    case 127:
                                                        fmt.Println(v12)
                                                    case 9223372036854775807:
                                                        v12 = 2147483647
                                                    default:
                                                        p0 = p0
                                                    }
                                                case "a":
                                                    switch (("a" + "\n\t") + ("" + "")) {
                                                    case "hello":
                                                        v24 := GenS2{F0: 567}
                                                        _ = v24
                                                    case "":
                                                        var v25 float64 = 1.5
                                                        _ = v25
                                                    }
                                                }
                                                if (!((!(1e10 >= float32(0.0))) && ((0.0 <= (0.0 / 0.1)) || false))) {
                                                    if false {
                                                        v12 = v12
                                                        v26 := GenS4{F0: 24, F1: 0, F2: v12}
                                                        if true {
                                                            switch -270 {
                                                            case 4294967295:
                                                                switch "\n\t" {
                                                                case "a":
                                                                    var v27 int = fi14
                                                                    _ = v27
                                                                case "foo bar":
                                                                    var v28 string = ((("\n\t" + "hello") + ("foo bar" + (("hello" + "") + ("foo bar" + ("" + "hello"))))) + ("\n\t" + ""))
                                                                    _ = v28
                                                                default:
                                                                    v21 = append(v21, fi14)
                                                                }
                                                            }
                                                            p1 = p0
                                                        }
                                                        _ = v26
                                                    } else {
                                                        fmt.Println(p0, p1, p0)
                                                        if false {
                                                            switch "hello" {
                                                            case "\n\t":
                                                                v21 = append(v21, fi14)
                                                            case "":
                                                                v12 = v12
                                                                v12 -= (-(v12))
                                                                var v29 int32 = v12
                                                                _ = v29
                                                            }
                                                            if true {
                                                                var v30 int64 = 755
                                                                v12 += v12
                                                                switch 32767 {
                                                                case 32767:
                                                                    if len(v21) > 0 {
                                                                        v21[0] = int(p1)
                                                                    }
                                                                case 377:
                                                                    var v31 int8 = 127
                                                                    _ = v31
                                                                case -964:
                                                                    p0 -= p0
                                                                default:
                                                                    if (uint(p1) >= 319) {
                                                                        var v32 uint64 = 2
                                                                        var v33 bool = (42 > 437)
                                                                        _ = v32
                                                                        _ = v33
                                                                    }
                                                                }
                                                                _ = v30
                                                            } else {
                                                                fmt.Println(fi19, p0)
                                                            }
                                                            switch (fi14 & fi14) {
                                                            case -1:
                                                                var v34 int8 = 127
                                                                _ = v34
                                                            case 617:
                                                                switch "a" {
                                                                case "\n\t":
                                                                    var v35 float64 = -2.25
                                                                    _ = v35
                                                                case "foo bar":
                                                                    v36 := GenS5{F0: 255, F1: 0, F2: -257}
                                                                    _ = v36
                                                                default:
                                                                    switch "a" {
                                                                    case "":
                                                                        fmt.Println(p1)
                                                                    case "hello":
                                                                        if len(v21) > 0 {
                                                                            v21[0] = 42
                                                                        }
                                                                    }
                                                                }
                                                            default:
                                                                if true {
                                                                    if len(v21) > 0 {
                                                                        v21[0] = fi14
                                                                    }
                                                                    switch (fi14 / ((fi19) | 1)) {
                                                                    case 4294967295:
                                                                        v21 = append(v21, (-(fi14)))
                                                                    case 920:
                                                                        if true {
                                                                            var v37 int64 = 9223372036854775807
                                                                            var v38 uint8 = (p0 << ((0) & 63))
                                                                            p0 = p1
                                                                            fmt.Println(p0, fi19)
                                                                            var v39 int16 = -699
                                                                            _ = v37
                                                                            _ = v38
                                                                            _ = v39
                                                                        }
                                                                    default:
                                                                        v40 := GenS4{F0: -128, F1: 18446744073709551615, F2: 42}
                                                                        v41 := GenS6{F0: 110}
                                                                        _ = v40
                                                                        _ = v41
                                                                    }
                                                                    p1 = 0
                                                                    if (((("" + ("hello" + "a")) + "") >= "a") && true) {
                                                                        v42 := GenS7{F0: (v12 | v12)}
                                                                        v21 = append(v21, fi14)
                                                                        _ = v42
                                                                    } else {
                                                                        fmt.Println(v12, p0)
                                                                    }
                                                                }
                                                            }
                                                            var v43 uint8 = p0
                                                            _ = len(v21)
                                                            if ((-2.25 / ((3.14159 * (((0.0 * float32((float64((1.5 / ((1e10 - -2.25) - float32(3.14159)))) * -2.25))) / 1.5) - (3.14159 * 3.14159))) / 3.14159)) > ((((3.14159 + 3.14159) - 0.1) / -2.25) + 1.5)) {
                                                                var v44 uint32 = 2
                                                                _ = v44
                                                            }
                                                            _ = v43
                                                        }
                                                    }
                                                    v45 := GenS2{F0: 209}
                                                    fmt.Println(fi19, p0)
                                                    _ = v45
                                                } else {
                                                    if false {
                                                        if ((-(v12)) <= (v12 >> ((990) & 63))) {
                                                            var v46 int8 = -42
                                                            fmt.Println(fi14)
                                                            fmt.Println(p0, fi19, p0)
                                                            _ = v46
                                                        } else {
                                                            if true {
                                                                var v47 uint8 = 121
                                                                var v48 int32 = (-(v12))
                                                                _ = v47
                                                                _ = v48
                                                            } else {
                                                                var v49 uint64 = genFn0(v12)
                                                                _ = v49
                                                            }
                                                            switch ((((("hello" + "\n\t") + "") + ((("a" + "a") + ("foo bar" + ("" + ("a" + ("foo bar" + "a"))))) + "")) + "a") + "hello") {
                                                            case "foo bar":
                                                                if (!(42 < 2147483648)) {
                                                                    p0 = p0
                                                                    p1 -= 0
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        var v50 uint64 = 4294967295
                                                        p1 += p0
                                                        _ = v50
                                                    }
                                                    if len(v21) >= 2 {
                                                        v21 = v21[1:len(v21)]
                                                    }
                                                }
                                                _ = v23
                                            }
                                            _ = fi19
                                            _ = v21
                                        }
                                        _ = v18
                                    } else {
                                        p0 += 127
                                    }
                                    p1 -= p0
                                } else {
                                    for fi51 := 0; fi51 < 6; fi51++ {
                                        var v52 uint8 = (^(p0))
                                        var v53 uint16 = 951
                                        var v54 string = "\n\t"
                                        _ = fi51
                                        _ = v52
                                        _ = v53
                                        _ = v54
                                    }
                                    fmt.Println(p1)
                                    var v55 uint = 18446744073709551615
                                    _ = v55
                                }
                            }
                            var v56 uint = 42
                            _ = v56
                        }
                        _ = fi14
                    }
                    _ = v12
                    _ = v13
                }
                p0 = 24
                v57 := GenS8{F0: 42, F1: 2, F2: -892}
                _ = v57
            }
            p0 = (p1 >> ((255) & 63))
            var v58 string = ("a" + ("hello" + "\n\t"))
            _ = v58
        } else {
            for fi59 := 0; fi59 < 18; fi59++ {
                fmt.Println(p1, p1)
                _ = fi59
            }
            var v60 int32 = 42
            _ = v60
        }
    case "":
        var v61 string = ""
        _ = v61
    default:
        v62 := GenS9{F0: -128, F1: p1}
        _ = v62
    }
    _ = p0
    _ = p1
    return 595
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[int]uint32)
    var v3 int64 = 42
    v4 := make(map[uint32]int)
    var v5 float32 = 0.1
    var v63 string = "foo bar"
    switch v63 {
    case "foo bar":
        if true {
            v64 := GenS10{F0: -15, F1: genFn0(42)}
            for fi65 := 0; fi65 < 15; fi65++ {
                v66 := []int{32768}
                _ = fi65
                _ = v66
            }
            v67 := GenS11{F0: 163, F1: 125, F2: 42}
            _ = v64
            _ = v67
        }
    case "a":
        for fi68 := 0; fi68 < 15; fi68++ {
            v1 -= (v1 + 0)
            v1 = 42
            _ = fi68
        }
    case "\n\t":
        var v69 string = "a"
        _ = v69
    default:
        v3 += (v3 - (v3 * (v3 & v3)))
    }
    if false {
        if (0.0 != 1e10) {
            switch 4294967295 {
            case -29:
                fmt.Println(v1)
            case 9223372036854775807:
                switch 9223372036854775807 {
                case 128:
                    v70 := GenS8{F0: v1, F1: (v1 - (v1 - (v1 - v1))), F2: -716}
                    _ = v70
                case 9223372036854775807:
                    v1 = (v1 | v1)
                }
            }
            v4[2] = 32767
            v1 -= 2
        } else {
            v71 := GenS12{F0: -471, F1: -2147483648}
            var v72 uint64 = uint64(v1)
            fmt.Println(v1, v72, v63)
            _ = v71
            _ = v72
        }
    }
    fmt.Println(v3, v1)
    var v73 string = v63
    v5 -= 0.1
    var v74 int = -1
    switch 904 {
    case 42:
        if true {
            v74 = (^(v74))
        } else {
            var v75 string = v73
            for fi76 := 0; fi76 < 6; fi76++ {
                for fi77 := 0; fi77 < 11; fi77++ {
                    fmt.Println(fi77, fi76)
                    var v78 int = fi77
                    _ = fi77
                    _ = v78
                }
                v5 = (v5 * (float32(3.14159) + v5))
                _ = fi76
            }
            fmt.Println(v3)
            _ = v75
        }
    case 0:
        var v79 uint64 = 42
        _ = v79
    case 1:
        if false {
            v80 := GenS8{F0: (v1 - (v1 * 0)), F1: (^(v1)), F2: 32767}
            for fi81 := 0; fi81 < 16; fi81++ {
                if false {
                    v3 = -1
                    if (42 <= 2147483648) {
                        switch fi81 {
                        case 736:
                            var v82 float32 = -2.25
                            _ = v82
                        case 461:
                            v74 += fi81
                        case 32767:
                            v83 := []int{(v74 + 32767), -9223372036854775808, fi81}
                            _ = v83
                        default:
                            var v84 float32 = float32(1.5)
                            _ = v84
                        }
                    } else {
                        if true {
                            fmt.Println(fi81, v73, v3)
                            if false {
                                for fi85 := 0; fi85 < 11; fi85++ {
                                    fmt.Println(v1, v1, fi81)
                                    if (-109 < -1) {
                                        if true {
                                            v74 -= fi85
                                            v86 := v2[(fi85 / ((v74) | 1))]
                                            fmt.Println(v73, fi85)
                                            _ = v86
                                        }
                                    }
                                    v74 -= (v74 - -920)
                                    v4[2147483647] = (v74 / ((32767) | 1))
                                    v87 := GenS10{F0: 10, F1: 26}
                                    _ = fi85
                                    _ = v87
                                }
                                switch "\n\t" {
                                case "hello":
                                    for fi88 := 0; fi88 < 4; fi88++ {
                                        v5 += 1e10
                                        var v89 uint16 = (^(v1))
                                        _ = fi88
                                        _ = v89
                                    }
                                case "a":
                                    var v90 int32 = -1
                                    _ = v90
                                case "":
                                    fmt.Println(v73)
                                }
                                if (false || false) {
                                    var v91 uint8 = 8
                                    _ = v91
                                }
                            } else {
                                fmt.Println(v63)
                            }
                        } else {
                            for fi92 := 0; fi92 < 12; fi92++ {
                                v3 += v3
                                v93 := GenS12{F0: 32767, F1: -377}
                                v94 := GenS4{F0: -128, F1: genFn0(2147483647), F2: 0}
                                fmt.Println(v74)
                                if true {
                                    v74 += v74
                                } else {
                                    v3 = v3
                                }
                                _ = fi92
                                _ = v93
                                _ = v94
                            }
                            fmt.Println(fi81, v73)
                            var v95 int16 = -32768
                            _ = v95
                        }
                        v96 := []int{(^(v74))}
                        var v97 uint8 = 127
                        _ = v96
                        _ = v97
                    }
                }
                for fi98 := 0; fi98 < 20; fi98++ {
                    switch (v63 + ((((("foo bar" + "") + v63) + v63) + v73) + v73)) {
                    case "a":
                        fmt.Println(v1, v5, fi81)
                    case "\n\t":
                        v2[fi81] = 301
                    }
                    switch fi81 {
                    case 133:
                        v2[fi81] = 4294967295
                    }
                    var v99 uint8 = 200
                    _ = fi98
                    _ = v99
                }
                _ = fi81
            }
            var v100 uint16 = (v1 << ((255) & 63))
            _ = v100
            _ = v80
        }
    default:
        fmt.Println(v73)
    }
    v101 := GenS8{F0: (v1 / (((v1 - v1)) | 1)), F1: (v1 + v1), F2: 42}
    v102 := []int{-9223372036854775808, -1, (v74 - 1)}
    for fi103 := 0; fi103 < 6; fi103++ {
        for fi104 := 0; fi104 < 6; fi104++ {
            v74 = 0
            if (!true) {
                v105 := GenS7{F0: -2147483648}
                var v106 int = 9223372036854775807
                _ = v105
                _ = v106
            }
            fmt.Println(v1, v63)
            _ = fi104
        }
        switch (("" + v63) + "foo bar") {
        case "\n\t":
            var v107 int = (fi103 >> ((699) & 63))
            _ = v107
        case "hello":
            var v108 float32 = 0.1
            _ = v108
        }
        _ = fi103
    }
    for fi109 := 0; fi109 < 12; fi109++ {
        v1 = 997
        for fi110 := 0; fi110 < 15; fi110++ {
            _ = len(v102)
            v73 = (((v73 + (((v73 + "") + (v63 + v63)) + "foo bar")) + v63) + (v73 + "hello"))
            _ = fi110
        }
        v3 -= -294
        _ = fi109
    }
    var v111 int16 = -32768
    v112 := GenS13{F0: ((4294967295 != 285) && (!true)), F1: v74, F2: (v63 + ("\n\t" + (((v63 + v63) + v73) + (v63 + (v63 + v73)))))}
    _ = len(v4)
    v113 := GenS10{F0: -1, F1: 214}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v102:", len(v102), v102)
    fmt.Println("v111:", v111)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v101
    _ = v102
    _ = v111
    _ = v112
    _ = v113
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v63
    _ = v73
    _ = v74
}
