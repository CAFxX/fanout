package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 string; F2 int16 }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint; F1 uint64 }
type GenS9 struct { F0 uint32; F1 bool; F2 uint }
type GenS10 struct { F0 uint; F1 uint8; F2 bool }
type GenS11 struct { F0 uint32; F1 int16 }
func genFn0(p0 int, p1 int8, p2 uint16) int16 {
    fmt.Println(p0, p1, p2)
    v6 := GenS1{F0: 42, F1: 4294967295}
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            v9 := GenS1{F0: -2147483648, F1: 255}
            switch 128 {
            case -429:
                var v10 float64 = 0.0
                _ = v10
            case -1:
                var v11 int16 = 69
                _ = v11
            case 790:
                fmt.Println(p1, p2)
            }
            p0 += fi8
            if (-2147483648 == (fi7 & (^(p0)))) {
                var v12 uint8 = 69
                var v13 bool = (!(-1 <= 42))
                v14 := GenS1{F0: 710, F1: 1}
                _ = v12
                _ = v13
                _ = v14
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v15 int16 = 1000
    if false {
        switch ("a" + "") {
        case "a":
            if true {
                v16 := []int{(p0 * (p0 / ((p0) | 1)))}
                for fi17 := 0; fi17 < 16; fi17++ {
                    switch (fi17 + fi17) {
                    case -630:
                        for fi18 := 0; fi18 < 3; fi18++ {
                            if false {
                                if false {
                                    p0 += p0
                                    v15 -= (v15 - v15)
                                    if (255 == 198) {
                                        if false {
                                            p0 -= 32767
                                            var v19 float32 = float32(-2.25)
                                            _ = len(v16)
                                            _ = v19
                                        } else {
                                            var v20 int32 = 42
                                            fmt.Println(v15)
                                            v15 = v15
                                            _ = v20
                                        }
                                        var v21 int32 = 0
                                        v15 = v15
                                        _ = v21
                                    }
                                }
                                if false {
                                    if (521 != 2) {
                                        if (false || (p2 != 32767)) {
                                            switch (fi17 + fi17) {
                                            case 4294967295:
                                                if len(v16) > 0 {
                                                    v16[0] = fi17
                                                }
                                            }
                                            var v22 string = "foo bar"
                                            v23 := GenS2{F0: 18446744073709551615}
                                            _ = len(v16)
                                            p2 -= p2
                                            _ = v22
                                            _ = v23
                                        } else {
                                            p2 = p2
                                            p2 -= p2
                                            var v24 uint64 = 0
                                            _ = v24
                                        }
                                        p1 += 1
                                        fmt.Println(p2)
                                    }
                                }
                                _ = len(v16)
                            } else {
                                var v25 int32 = 119
                                p0 = fi18
                                _ = v25
                            }
                            v26 := GenS3{F0: p1, F1: "foo bar", F2: -1}
                            if false {
                                p0 = (p0 >> ((uint(p2)) & 63))
                                p1 = (-(p1))
                                var v27 int8 = (p1 + p1)
                                _ = v27
                            } else {
                                fmt.Println(p0, v15)
                                p1 = p1
                                v16 = append(v16, (^(p0)))
                            }
                            _ = fi18
                            _ = v26
                        }
                    default:
                        if ((p2 + p2) > 2) {
                            v15 -= 32767
                            switch 0 {
                            case 1:
                                if len(v16) > 0 {
                                    v16[0] = 187
                                }
                            case 128:
                                switch p0 {
                                case 399:
                                    v16 = append(v16, (p0 & (fi17 % ((p0) | 1))))
                                }
                            case -9223372036854775808:
                                _ = len(v16)
                            }
                        } else {
                            var v28 uint16 = p2
                            fmt.Println(v15)
                            v15 -= (v15 ^ v15)
                            _ = v28
                        }
                    }
                    var v29 uint32 = 585
                    if ((1 == -363) || true) {
                        for fi30 := 0; fi30 < 18; fi30++ {
                            var v31 uint8 = 1
                            v32 := GenS3{F0: int8(v15), F1: "a", F2: (v15 & 1)}
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                    } else {
                        for fi33 := 0; fi33 < 20; fi33++ {
                            if len(v16) > 0 {
                                v16[0] = (fi17 >> ((729) & 63))
                            }
                            var v34 int64 = 42
                            _ = fi33
                            _ = v34
                        }
                        v15 = (v15 - (v15 + (-(v15))))
                    }
                    _ = fi17
                    _ = v29
                }
                _ = v16
            }
        default:
            v35 := GenS2{F0: 1}
            _ = v35
        }
        fmt.Println(v15, p0, p1)
        var v36 int = p0
        _ = v36
    } else {
        v37 := GenS1{F0: -1, F1: 255}
        v38 := []int{p0}
        _ = v37
        _ = v38
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v6
    return (v15 / ((-1) | 1))
}
func genFn1(p0 int32, p1 uint) uint8 {
    for fi39 := 0; fi39 < 5; fi39++ {
        switch 0 {
        case 32767:
            v40 := GenS2{F0: 544}
            _ = v40
        default:
            switch ("" + "a") {
            case "":
                v41 := GenS2{F0: 0}
                _ = v41
            default:
                v42 := GenS2{F0: 260}
                _ = v42
            }
        }
        switch fi39 {
        case 9223372036854775807:
            var v43 int8 = 127
            _ = v43
        case 954:
            for fi44 := 0; fi44 < 2; fi44++ {
                v45 := GenS2{F0: 18446744073709551615}
                p0 -= (p0 | p0)
                v46 := []int{(fi44 - (fi44 & 9223372036854775807)), fi39}
                switch "" {
                case "hello":
                    p1 -= (p1 * 255)
                    p1 -= (p1 >> ((0) & 63))
                case "":
                    var v47 int16 = -223
                    _ = v47
                case "\n\t":
                    p0 -= p0
                default:
                    switch -9223372036854775808 {
                    case 2147483647:
                        p0 -= (-(p0))
                    case -1:
                        switch fi44 {
                        case 4294967295:
                            var v48 int = 42
                            _ = v48
                        case 42:
                            if (127 <= 97) {
                                fmt.Println(p1)
                                p0 -= (-(p0))
                            }
                        default:
                            p0 = p0
                        }
                    default:
                        if (p1 == 65535) {
                            fmt.Println(p1)
                            v49 := GenS4{F0: 255, F1: 127}
                            if (!(!(1e10 <= 1.5))) {
                                switch "\n\t" {
                                case "foo bar":
                                    p0 = (p0 * p0)
                                case "a":
                                    var v50 int8 = -17
                                    p1 = p1
                                    _ = v50
                                default:
                                    v51 := GenS5{F0: 78, F1: 2}
                                    _ = v51
                                }
                            } else {
                                switch "\n\t" {
                                case "\n\t":
                                    var v52 uint32 = 1
                                    _ = v52
                                }
                            }
                            _ = v49
                        } else {
                            if (false && false) {
                                v46 = append(v46, 32768)
                                if (176 < 48) {
                                    if true {
                                        if false {
                                            var v53 bool = true
                                            var v54 string = ("a" + "hello")
                                            v55 := GenS6{F0: 16, F1: -1, F2: 1}
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        } else {
                                            if (true && true) {
                                                switch ("foo bar" + "") {
                                                case "\n\t":
                                                    switch "" {
                                                    case "hello":
                                                        var v56 uint8 = 2
                                                        _ = v56
                                                    default:
                                                        v57 := GenS7{F0: 42}
                                                        _ = v57
                                                    }
                                                case "foo bar":
                                                    var v58 bool = false
                                                    _ = v58
                                                case "a":
                                                    p0 = p0
                                                }
                                                var v59 float64 = 0.1
                                                _ = v59
                                            }
                                            if true {
                                                v60 := GenS4{F0: 2, F1: 57}
                                                p1 += (p1 | p1)
                                                _ = v60
                                            }
                                            if len(v46) > 0 {
                                                v46[0] = (fi44 - fi39)
                                            }
                                            var v61 string = ""
                                            var v62 int16 = int16(p0)
                                            var v63 int8 = 70
                                            switch v61 {
                                            case "a":
                                                var v64 int = (^(fi39))
                                                _ = v64
                                            case "":
                                                switch (fi39 ^ int(v62)) {
                                                case -57:
                                                    _ = len(v46)
                                                case -207:
                                                    switch ((v61 + v61) + ((v61 + v61) + v61)) {
                                                    case "":
                                                        _ = len(v46)
                                                        if (0.0 >= -2.25) {
                                                            v65 := GenS1{F0: (p0 << ((18446744073709551615) & 63)), F1: p1}
                                                            _ = v65
                                                        }
                                                        if false {
                                                            if (false && (!false)) {
                                                                var v66 uint64 = 4294967295
                                                                _ = v66
                                                            } else {
                                                                v67 := GenS8{F0: (^(p1)), F1: 360}
                                                                _ = v67
                                                            }
                                                            v68 := GenS9{F0: 0, F1: true, F2: (^(p1))}
                                                            _ = v68
                                                        }
                                                    default:
                                                        switch (fi44 * fi44) {
                                                        case 127:
                                                            var v69 int8 = 42
                                                            _ = v69
                                                        case 32767:
                                                            v46 = append(v46, fi44)
                                                        case -481:
                                                            var v70 bool = (((^(p1)) == 255) && false)
                                                            _ = v70
                                                        }
                                                    }
                                                case 127:
                                                    switch v61 {
                                                    case "foo bar":
                                                        var v71 uint = (p1 % ((0) | 1))
                                                        _ = v71
                                                    }
                                                }
                                            }
                                            _ = v61
                                            _ = v62
                                            _ = v63
                                        }
                                        v72 := GenS6{F0: 53, F1: 1, F2: 0}
                                        fmt.Println(p0)
                                        _ = v72
                                    } else {
                                        if (!false) {
                                            switch "hello" {
                                            case "\n\t":
                                                v73 := GenS7{F0: 0}
                                                _ = v73
                                            case "":
                                                p0 = p0
                                            }
                                            p0 = p0
                                        } else {
                                            if false {
                                                switch fi44 {
                                                case 9223372036854775807:
                                                    fmt.Println(p0)
                                                case 0:
                                                    v74 := GenS6{F0: 128, F1: -9223372036854775808, F2: 140}
                                                    _ = v74
                                                case 32768:
                                                    var v75 uint32 = 42
                                                    _ = v75
                                                }
                                            } else {
                                                p0 = p0
                                                switch (fi44 & fi44) {
                                                case 939:
                                                    v76 := GenS3{F0: -38, F1: "a", F2: 0}
                                                    _ = v76
                                                case -341:
                                                    p0 += p0
                                                case 127:
                                                    fmt.Println(p0, p1, p0)
                                                default:
                                                    p1 += (p1 >> ((1) & 63))
                                                }
                                                var v77 int64 = -9223372036854775808
                                                var v78 string = "foo bar"
                                                var v79 int32 = (^(p0))
                                                _ = v77
                                                _ = v78
                                                _ = v79
                                            }
                                        }
                                        fmt.Println(fi39, fi44, fi44)
                                        p0 = p0
                                    }
                                    if len(v46) > 0 {
                                        v46[0] = int(p1)
                                    }
                                }
                            } else {
                                fmt.Println(fi39)
                                v80 := GenS5{F0: 65535, F1: 0}
                                _ = v80
                            }
                            var v81 bool = (true || (-71 >= -1))
                            _ = v81
                        }
                    }
                }
                if true {
                    if false {
                        var v82 int = (fi44 % ((1) | 1))
                        p0 = (p0 >> ((p1) & 63))
                        fmt.Println(fi44)
                        _ = v82
                    } else {
                        p0 = (p0 >> ((18446744073709551615) & 63))
                    }
                }
                if true {
                    p1 -= 1
                    var v83 bool = ((p0 | (-(p0))) < p0)
                    if len(v46) > 0 {
                        v46[0] = 32768
                    }
                    fmt.Println(fi44, fi44)
                    _ = len(v46)
                    _ = v83
                }
                _ = fi44
                _ = v45
                _ = v46
            }
        case 32768:
            v84 := GenS4{F0: 128, F1: 1}
            _ = v84
        default:
            var v85 float32 = (-2.25 * 0.0)
            _ = v85
        }
        _ = fi39
    }
    var v86 int = 1
    if ((-2.25 * ((3.14159 / -2.25) - 1.5)) > 0.1) {
        var v87 float32 = -2.25
        for fi88 := 0; fi88 < 2; fi88++ {
            var v89 uint32 = 0
            v90 := GenS7{F0: fi88}
            var v91 uint8 = 255
            _ = fi88
            _ = v89
            _ = v90
            _ = v91
        }
        _ = v87
    } else {
        p1 = 42
        switch (v86 - v86) {
        case 42:
            if (false || (!(2 == 679))) {
                var v92 float64 = 1e10
                _ = v92
            } else {
                var v93 int32 = p0
                v94 := GenS7{F0: -1}
                _ = v93
                _ = v94
            }
        case 32767:
            v95 := []int{}
            _ = v95
        case -764:
            if (!false) {
                p1 += 18446744073709551615
                v96 := []int{v86, (v86 % ((v86) | 1)), (-(v86)), (v86 << (((^(p1))) & 63))}
                p1 += (p1 + p1)
                _ = v96
            }
        }
        v97 := GenS10{F0: (p1 & (^(p1))), F1: 127, F2: true}
        _ = v97
    }
    for fi98 := 0; fi98 < 13; fi98++ {
        if true {
            for fi99 := 0; fi99 < 18; fi99++ {
                v100 := []int{}
                _ = fi99
                _ = v100
            }
            var v101 int32 = p0
            v86 = 127
            _ = v101
        }
        _ = fi98
    }
    _ = p0
    _ = p1
    _ = v86
    return 128
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 557
    var v3 int8 = -1
    var v4 float64 = (0.0 * -2.25)
    var v5 int = 2147483647
    v3 = v3
    v102 := []int{(v5 * v5), v5, (-(v5))}
    if (0 <= 255) {
        var v103 bool = ((0 < 647) || (-794 < 0))
        var v104 uint8 = 1
        _ = v103
        _ = v104
    } else {
        for fi105 := 0; fi105 < 14; fi105++ {
            switch ("hello" + "foo bar") {
            case "\n\t":
                for fi106 := 0; fi106 < 2; fi106++ {
                    v107 := GenS10{F0: 760, F1: 163, F2: true}
                    _ = fi106
                    _ = v107
                }
            case "a":
                if len(v102) > 0 {
                    v102[0] = v5
                }
            default:
                var v108 uint = 704
                _ = v108
            }
            delete(v1, 570)
            v5 = -149
            _ = fi105
        }
    }
    var v109 float64 = (((v4 / 1.0) * v4) / v4)
    fmt.Println(v5, v3, v2)
    if false {
        switch "hello" {
        case "foo bar":
            v110 := GenS5{F0: uint32(v2), F1: 0}
            _ = v110
        case "a":
            var v111 float64 = ((v4 - 0.1) * 1.5)
            _ = v111
        }
    } else {
        var v112 uint8 = 1
        v2 = (^(v2))
        _ = v112
    }
    for fi113 := 0; fi113 < 11; fi113++ {
        var v114 int8 = (v3 % ((35) | 1))
        if (-1 >= v114) {
            if (!true) {
                if false {
                    fmt.Println(v109)
                } else {
                    for fi115 := 0; fi115 < 6; fi115++ {
                        if true {
                            fmt.Println(fi113)
                            v109 -= 1e10
                        } else {
                            switch fi115 {
                            case 9223372036854775807:
                                delete(v1, 42)
                            case 32767:
                                fmt.Println(v3)
                            case -321:
                                fmt.Println(v2)
                            default:
                                var v116 uint8 = 1
                                _ = v116
                            }
                            fmt.Println(v3, v114)
                            v5 = fi113
                        }
                        fmt.Println(v109, v5, fi115)
                        _ = fi115
                    }
                    switch ("foo bar" + ((("a" + "") + "a") + ("hello" + ""))) {
                    case "foo bar":
                        _ = len(v102)
                    case "hello":
                        var v117 string = "a"
                        _ = v117
                    }
                    v118 := GenS11{F0: 65535, F1: -1}
                    _ = v118
                }
            }
        }
        v2 = v2
        _ = fi113
        _ = v114
    }
    fmt.Println(v109, v5, v109)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v102:", len(v102), v102)
    fmt.Println("v109:", v109)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v102
    _ = v109
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
