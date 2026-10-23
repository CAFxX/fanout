package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 string }
type GenS2 struct { F0 int16; F1 uint; F2 int32 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 bool; F2 string }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 int32; F1 uint8; F2 string }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint; F1 uint64 }
type GenS10 struct { F0 int16; F1 int16 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint8; F1 int; F2 int }
type GenS14 struct { F0 uint32; F1 uint16; F2 int }
type GenS15 struct { F0 int32; F1 int32; F2 string }
type GenS16 struct { F0 uint8; F1 uint }
type GenS17 struct { F0 string; F1 int16; F2 int32 }
type GenS18 struct { F0 int64; F1 int32; F2 uint32 }
type GenS19 struct { F0 int; F1 int64 }
type GenS20 struct { F0 int8; F1 uint64 }
type GenS21 struct { F0 int64; F1 uint32 }
type GenS22 struct { F0 int8; F1 uint8 }
func genFn0(p0 int64, p1 string) int16 {
    for fi6 := 0; fi6 < 6; fi6++ {
        switch (^(fi6)) {
        case 127:
            v7 := GenS1{F0: 2147483648, F1: -2147483648, F2: (p1 + p1)}
            _ = v7
        case -1:
            if false {
                var v8 uint32 = 701
                switch ("a" + p1) {
                case "foo bar":
                    v9 := GenS1{F0: (v8 >> ((18446744073709551615) & 63)), F1: -969, F2: p1}
                    _ = v9
                }
                _ = v8
            }
        }
        if false {
            for fi10 := 0; fi10 < 16; fi10++ {
                v11 := GenS1{F0: 463, F1: -102, F2: ("a" + "")}
                switch "" {
                case "\n\t":
                    p0 = 42
                case "foo bar":
                    if false {
                        v12 := GenS2{F0: 973, F1: 2, F2: -2147483648}
                        _ = v12
                    } else {
                        p1 = (p1 + p1)
                        fmt.Println(fi10, p1)
                        var v13 float32 = 3.14159
                        _ = v13
                    }
                }
                _ = fi10
                _ = v11
            }
            if false {
                if true {
                    v14 := []int{-132}
                    _ = v14
                }
            } else {
                if false {
                    if (false || true) {
                        for fi15 := 0; fi15 < 17; fi15++ {
                            var v16 bool = (450 == 42)
                            _ = fi15
                            _ = v16
                        }
                        var v17 uint8 = 15
                        p1 = (p1 + (p1 + "foo bar"))
                        _ = v17
                    }
                }
                for fi18 := 0; fi18 < 1; fi18++ {
                    p0 = (p0 ^ p0)
                    _ = fi18
                }
            }
            fmt.Println(p1, p1, p1)
        }
        if false {
            var v19 int16 = 1
            var v20 string = "a"
            _ = v19
            _ = v20
        } else {
            v21 := GenS3{F0: -984, F1: 65535, F2: 4294967295}
            _ = v21
        }
        _ = fi6
    }
    fmt.Println(p1, p0)
    for fi22 := 0; fi22 < 15; fi22++ {
        v23 := GenS4{F0: 42}
        var v24 uint = 255
        switch p1 {
        case "foo bar":
            v24 += v24
        case "":
            p1 = "foo bar"
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 int = 32767
    for fi26 := 0; fi26 < 15; fi26++ {
        v27 := GenS2{F0: -583, F1: 65535, F2: 0}
        switch "a" {
        case "a":
            switch p1 {
            case "foo bar":
                var v28 int16 = -32768
                _ = v28
            case "hello":
                if (!(0 == 18446744073709551615)) {
                    var v29 int32 = -941
                    if (true || (false && (1e10 >= 0.1))) {
                        fmt.Println(p0, p0, p0)
                    } else {
                        v29 = -990
                    }
                    _ = v29
                } else {
                    var v30 int16 = 1
                    _ = v30
                }
            case "":
                var v31 uint64 = 895
                _ = v31
            default:
                v32 := []int{32768}
                _ = v32
            }
        case "\n\t":
            if true {
                switch p1 {
                case "\n\t":
                    v33 := []int{v25}
                    _ = v33
                }
                if ((float64(0.1) + 0.1) > ((-2.25 * (((float64(1.5) / (-2.25 - ((((((1.5 + (1.5 - float64(0.0))) * (-2.25 - 3.14159)) - (-2.25 - (0.0 * (((float64(1.5) * 0.1) + float64(1e10)) + (-2.25 + ((3.14159 * 3.14159) * 0.0)))))) + 1e10) + 0.1) - -2.25))) / 3.14159) / (((-2.25 + 3.14159) + -2.25) + 0.1))) / 1.5)) {
                    for fi34 := 0; fi34 < 8; fi34++ {
                        switch p1 {
                        case "foo bar":
                            p1 = ("foo bar" + p1)
                        default:
                            v35 := GenS5{F0: v25, F1: (852 == 1), F2: "a"}
                            _ = v35
                        }
                        v36 := GenS1{F0: 0, F1: 42, F2: p1}
                        _ = fi34
                        _ = v36
                    }
                }
                v37 := GenS4{F0: 255}
                _ = v37
            } else {
                if true {
                    var v38 int = 32767
                    _ = v38
                }
                for fi39 := 0; fi39 < 14; fi39++ {
                    switch (v25 + (v25 + -7)) {
                    case -1:
                        v40 := GenS4{F0: 62}
                        _ = v40
                    case 2147483647:
                        if true {
                            var v41 uint64 = 18446744073709551615
                            if (true && ((p0 << ((42) & 63)) <= (p0 * 42))) {
                                fmt.Println(fi26, fi39, v25)
                                switch p1 {
                                case "":
                                    if true {
                                        p0 = (p0 ^ p0)
                                        v42 := []int{fi26, -9223372036854775808, fi39, (-(fi39))}
                                        fmt.Println(v25, p0, v25)
                                        _ = v42
                                    } else {
                                        fmt.Println(v41, p0)
                                    }
                                default:
                                    v43 := []int{(-(v25)), v25, (^(fi39)), (^(fi39))}
                                    _ = v43
                                }
                                fmt.Println(fi26)
                            } else {
                                switch 4294967295 {
                                case -419:
                                    p1 = ("a" + "\n\t")
                                case 32767:
                                    var v44 uint = 653
                                    _ = v44
                                default:
                                    p0 = -9223372036854775808
                                }
                                v45 := GenS5{F0: 32767, F1: false, F2: p1}
                                v46 := GenS6{F0: 42, F1: (p0 | (p0 + (p0 & p0)))}
                                _ = v45
                                _ = v46
                            }
                            fmt.Println(v41)
                            _ = v41
                        } else {
                            if (463 > 316) {
                                if (3.14159 <= 1e10) {
                                    switch fi39 {
                                    case 128:
                                        if true {
                                            if (18446744073709551615 != 1) {
                                                v47 := GenS7{F0: 0, F1: 128, F2: p1}
                                                _ = v47
                                            } else {
                                                switch ("a" + ("foo bar" + "a")) {
                                                case "\n\t":
                                                    if false {
                                                        var v48 float32 = 1e10
                                                        v25 -= 2147483647
                                                        _ = v48
                                                    }
                                                    p1 = (p1 + p1)
                                                default:
                                                    v25 -= (fi26 >> ((42) & 63))
                                                }
                                                v49 := GenS8{F0: fi26}
                                                switch (v25 ^ 240) {
                                                case -1:
                                                    var v50 uint8 = 125
                                                    _ = v50
                                                case 32768:
                                                    p0 += p0
                                                default:
                                                    v51 := GenS7{F0: -502, F1: 119, F2: p1}
                                                    _ = v51
                                                }
                                                p1 = (p1 + "\n\t")
                                                p0 -= p0
                                                switch "hello" {
                                                case "":
                                                    p0 = (p0 & (p0 * p0))
                                                case "\n\t":
                                                    if (true || (-2.25 < (0.0 / -2.25))) {
                                                        fmt.Println(fi39, v25, v25)
                                                        var v52 string = p1
                                                        v53 := []int{fi26, fi26}
                                                        _ = v52
                                                        _ = v53
                                                    } else {
                                                        var v54 int = 127
                                                        if (p0 != p0) {
                                                            fmt.Println(v25, p1, p1)
                                                            v55 := []int{(v25 << ((739) & 63)), fi26, (v54 << ((42) & 63))}
                                                            _ = v55
                                                        } else {
                                                            switch "\n\t" {
                                                            case "":
                                                                if false {
                                                                    if (282 > 42) {
                                                                        if false {
                                                                            v56 := []int{}
                                                                            if ((!((fi26 >> ((0) & 63)) == (v54 + fi26))) || false) {
                                                                                if true {
                                                                                    var v57 bool = false
                                                                                    _ = v57
                                                                                }
                                                                            } else {
                                                                                switch p1 {
                                                                                case "hello":
                                                                                    v25 = v54
                                                                                    var v58 float64 = (((((1e10 / -2.25) * (-2.25 * 0.0)) + 3.14159) + ((3.14159 - (1.5 * (float64(0.0) * (1.5 * 0.1)))) + (3.14159 * 1.5))) / 1e10)
                                                                                    _ = v58
                                                                                default:
                                                                                    var v59 int64 = (-(p0))
                                                                                    _ = v59
                                                                                }
                                                                            }
                                                                            p1 = p1
                                                                            _ = v56
                                                                        }
                                                                        v54 = (fi26 / (((fi39 + (fi26 ^ fi39))) | 1))
                                                                        if true {
                                                                            fmt.Println(p1, fi26)
                                                                            fmt.Println(fi26, v25, fi26)
                                                                        }
                                                                        v60 := []int{}
                                                                        _ = v60
                                                                    } else {
                                                                        if ("a" < p1) {
                                                                            switch 32768 {
                                                                            case -9223372036854775808:
                                                                                v54 = (fi26 >> ((2) & 63))
                                                                            default:
                                                                                if (1 < 138) {
                                                                                    v61 := GenS2{F0: 42, F1: 18446744073709551615, F2: -267}
                                                                                    if (-1 <= 1) {
                                                                                        switch v25 {
                                                                                        case 42:
                                                                                            var v62 uint32 = 979
                                                                                            _ = v62
                                                                                        case 0:
                                                                                            var v63 uint16 = 255
                                                                                            _ = v63
                                                                                        }
                                                                                        var v64 uint16 = 32767
                                                                                        p0 += p0
                                                                                        v65 := GenS9{F0: 18446744073709551615, F1: 4294967295}
                                                                                        p1 = (p1 + "")
                                                                                        _ = v64
                                                                                        _ = v65
                                                                                    } else {
                                                                                        v66 := []int{(-(v25)), v54, fi39}
                                                                                        _ = v66
                                                                                    }
                                                                                    _ = v61
                                                                                } else {
                                                                                    var v67 uint32 = 886
                                                                                    if false {
                                                                                        var v68 uint8 = 127
                                                                                        fmt.Println(v68, fi39, v68)
                                                                                        _ = v68
                                                                                    } else {
                                                                                        if true {
                                                                                            fmt.Println(fi26)
                                                                                            switch p1 {
                                                                                            case "hello":
                                                                                                v69 := []int{4294967295, -548}
                                                                                                _ = v69
                                                                                            case "\n\t":
                                                                                                v67 += 612
                                                                                            }
                                                                                        } else {
                                                                                            v70 := []int{0, -193, v25, v54}
                                                                                            switch (fi26 >> ((255) & 63)) {
                                                                                            case -845:
                                                                                                p1 = ((p1 + (((p1 + ((p1 + p1) + ("\n\t" + p1))) + ((p1 + ((p1 + p1) + p1)) + (((p1 + p1) + (p1 + "\n\t")) + p1))) + p1)) + p1)
                                                                                            case -1:
                                                                                                _ = len(v70)
                                                                                            case 748:
                                                                                                var v71 float64 = -2.25
                                                                                                _ = v71
                                                                                            default:
                                                                                                var v72 uint64 = 0
                                                                                                _ = v72
                                                                                            }
                                                                                            v70 = append(v70, 32767)
                                                                                            var v73 string = "\n\t"
                                                                                            _ = v70
                                                                                            _ = v73
                                                                                        }
                                                                                        var v74 string = p1
                                                                                        v75 := GenS4{F0: 255}
                                                                                        _ = v74
                                                                                        _ = v75
                                                                                    }
                                                                                    _ = v67
                                                                                }
                                                                            }
                                                                            v76 := []int{(^(v25)), (v25 << ((4294967295) & 63)), (fi39 * 561)}
                                                                            v76 = append(v76, fi26)
                                                                            if (false && (65535 == 0)) {
                                                                                v25 = (^(v54))
                                                                            } else {
                                                                                switch (fi26 >> ((255) & 63)) {
                                                                                case 0:
                                                                                    var v77 int16 = 42
                                                                                    _ = v77
                                                                                case 698:
                                                                                    if len(v76) >= 2 {
                                                                                        v76 = v76[1:len(v76)]
                                                                                    }
                                                                                case 32767:
                                                                                    var v78 bool = true
                                                                                    _ = v78
                                                                                }
                                                                                fmt.Println(v25, v54, v54)
                                                                            }
                                                                            _ = v76
                                                                        }
                                                                        var v79 uint64 = 784
                                                                        _ = v79
                                                                    }
                                                                }
                                                            case "a":
                                                                fmt.Println(p0)
                                                            }
                                                            v25 = fi39
                                                            fmt.Println(p0, v54, v25)
                                                        }
                                                        _ = v54
                                                    }
                                                default:
                                                    v80 := GenS7{F0: 42, F1: 236, F2: ("foo bar" + (("" + "foo bar") + ((("" + p1) + "") + ((p1 + p1) + p1))))}
                                                    _ = v80
                                                }
                                                p0 = (^(p0))
                                                p0 = p0
                                                v81 := []int{32767}
                                                var v82 uint = 255
                                                if (12 > 42) {
                                                    if ((-2.25 - (1.5 - (0.1 / 1e10))) != 0.1) {
                                                        var v83 uint = v82
                                                        p0 -= (p0 << ((v83) & 63))
                                                        if false {
                                                            v82 -= (^(v82))
                                                            fmt.Println(p0, p1, fi39)
                                                        } else {
                                                            var v84 int8 = 127
                                                            v82 = 255
                                                            fmt.Println(p0)
                                                            p0 -= p0
                                                            _ = v84
                                                        }
                                                        if (123 > 32767) {
                                                            p1 = ("hello" + ("foo bar" + ("foo bar" + "a")))
                                                            if true {
                                                                if false {
                                                                    fmt.Println(p1, fi26)
                                                                    if (((p0 & (-(p0))) <= 0) && (1e10 <= 0.0)) {
                                                                        var v85 string = p1
                                                                        _ = v85
                                                                    }
                                                                    v83 = v82
                                                                }
                                                                if true {
                                                                    _ = len(v81)
                                                                    if true {
                                                                        var v86 uint8 = 146
                                                                        var v87 uint16 = 87
                                                                        _ = v86
                                                                        _ = v87
                                                                    } else {
                                                                        p0 += p0
                                                                    }
                                                                } else {
                                                                    var v88 uint16 = 613
                                                                    _ = v88
                                                                }
                                                                v89 := GenS5{F0: -572, F1: true, F2: p1}
                                                                _ = v89
                                                            }
                                                            p1 = ((p1 + (p1 + (p1 + p1))) + p1)
                                                        } else {
                                                            fmt.Println(v83, v82, v82)
                                                            v25 -= v25
                                                        }
                                                        _ = v83
                                                    }
                                                    if len(v81) >= 2 {
                                                        v81 = v81[1:len(v81)]
                                                    }
                                                    fmt.Println(fi39)
                                                } else {
                                                    switch 2147483647 {
                                                    case 128:
                                                        var v90 int64 = 682
                                                        _ = v90
                                                    case 4294967295:
                                                        if false {
                                                            var v91 int16 = 445
                                                            _ = v91
                                                        }
                                                        var v92 uint16 = 2
                                                        fmt.Println(fi39)
                                                        v82 = v82
                                                        if false {
                                                            v25 -= (fi39 | v25)
                                                            switch fi26 {
                                                            case 0:
                                                                if len(v81) > 0 {
                                                                    v81[0] = fi26
                                                                }
                                                            case 793:
                                                                if false {
                                                                    var v93 bool = ("a" == p1)
                                                                    var v94 float32 = 1e10
                                                                    _ = v93
                                                                    _ = v94
                                                                }
                                                            case 9223372036854775807:
                                                                if false {
                                                                    if (!((false || (("hello" != "") || true)) && ((((((((1.5 + 1.5) * -2.25) - float64(((-2.25 - (0.1 * (1.5 - (0.1 / 1.5)))) - 0.0))) / (((1e10 * (float64(((1e10 * float32((0.1 - 1.5))) * 0.1)) + 3.14159)) / 1.0) / 3.14159)) - 1.5) - -2.25) == (3.14159 - 0.0)) || (v92 != (^(v92)))))) {
                                                                        var v95 int64 = (^(p0))
                                                                        p0 = (-(p0))
                                                                        _ = v95
                                                                    }
                                                                    p0 = (p0 / ((int64(v82)) | 1))
                                                                    v25 += fi39
                                                                } else {
                                                                    if false {
                                                                        if len(v81) >= 2 {
                                                                            v81 = v81[1:len(v81)]
                                                                        }
                                                                        fmt.Println(fi39)
                                                                    }
                                                                    p0 = p0
                                                                    if true {
                                                                        v25 -= v25
                                                                        var v96 int8 = 100
                                                                        if (!false) {
                                                                            fmt.Println(fi39, p0, v92)
                                                                        } else {
                                                                            _ = len(v81)
                                                                            v97 := GenS3{F0: (p0 ^ 0), F1: v92, F2: 128}
                                                                            _ = v97
                                                                        }
                                                                        if false {
                                                                            v98 := GenS8{F0: v25}
                                                                            if len(v81) >= 2 {
                                                                                v81 = v81[1:len(v81)]
                                                                            }
                                                                            _ = v98
                                                                        }
                                                                        _ = v96
                                                                    }
                                                                }
                                                            }
                                                            var v99 string = ""
                                                            _ = v99
                                                        } else {
                                                            v100 := GenS10{F0: 1000, F1: 0}
                                                            _ = v100
                                                        }
                                                        v25 += 4294967295
                                                        v82 = 603
                                                        fmt.Println(v82)
                                                        var v101 uint64 = 2
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            p0 += p0
                                                        default:
                                                            fmt.Println(v92, v101, v25)
                                                        }
                                                        _ = v101
                                                        _ = v92
                                                    default:
                                                        switch "" {
                                                        case "a":
                                                            v102 := GenS11{F0: v25}
                                                            p1 = (("foo bar" + "a") + "\n\t")
                                                            fmt.Println(v82, v25)
                                                            var v103 int64 = p0
                                                            _ = v102
                                                            _ = v103
                                                        }
                                                        p1 = ("\n\t" + p1)
                                                        var v104 uint64 = 0
                                                        _ = v104
                                                    }
                                                    var v105 float64 = -2.25
                                                    _ = v105
                                                }
                                                v106 := GenS12{F0: (fi26 * fi26)}
                                                v107 := GenS2{F0: -344, F1: 0, F2: -1}
                                                var v108 uint32 = 2
                                                var v109 uint8 = 2
                                                _ = v106
                                                _ = v107
                                                _ = v108
                                                _ = v109
                                                _ = v49
                                                _ = v81
                                                _ = v82
                                            }
                                            var v110 uint32 = 65535
                                            v111 := GenS13{F0: 57, F1: 1, F2: 749}
                                            _ = v110
                                            _ = v111
                                        }
                                    case 81:
                                        var v112 bool = false
                                        _ = v112
                                    case 4294967295:
                                        v113 := []int{(fi26 << ((823) & 63)), 32767}
                                        _ = v113
                                    default:
                                        var v114 uint8 = 128
                                        _ = v114
                                    }
                                    var v115 uint32 = 42
                                    v116 := GenS2{F0: 32767, F1: 933, F2: 1}
                                    _ = v115
                                    _ = v116
                                } else {
                                    v25 = -670
                                    if true {
                                        v117 := GenS2{F0: 1, F1: 1, F2: -1}
                                        if true {
                                            p0 -= (p0 | (p0 & p0))
                                            var v118 uint8 = 2
                                            if true {
                                                v119 := GenS14{F0: 42, F1: 42, F2: v25}
                                                v120 := []int{fi26, 42, (v25 & (v25 % ((402) | 1)))}
                                                _ = v119
                                                _ = v120
                                            }
                                            _ = v118
                                        }
                                        if (1 > 65535) {
                                            v121 := GenS10{F0: 0, F1: 42}
                                            if false {
                                                v122 := []int{(v25 + -9223372036854775808)}
                                                switch (fi39 / ((v25) | 1)) {
                                                case 9223372036854775807:
                                                    p0 = (^(p0))
                                                case -1:
                                                    fmt.Println(fi26, v25, p1)
                                                }
                                                _ = v122
                                            }
                                            _ = v121
                                        } else {
                                            var v123 int16 = 995
                                            switch "hello" {
                                            case "foo bar":
                                                var v124 int = 128
                                                _ = v124
                                            case "hello":
                                                v25 += (v25 % ((fi39) | 1))
                                            default:
                                                var v125 uint64 = 1
                                                _ = v125
                                            }
                                            v25 += v25
                                            _ = v123
                                        }
                                        _ = v117
                                    } else {
                                        v25 = fi39
                                        fmt.Println(fi26)
                                    }
                                    if true {
                                        switch v25 {
                                        case 0:
                                            v126 := GenS4{F0: 228}
                                            v127 := GenS7{F0: -923, F1: 0, F2: p1}
                                            v25 -= (v25 << ((65535) & 63))
                                            _ = v126
                                            _ = v127
                                        }
                                        v128 := GenS10{F0: 42, F1: 161}
                                        _ = v128
                                    }
                                }
                            } else {
                                if true {
                                    p0 -= (p0 - -715)
                                } else {
                                    var v129 string = "\n\t"
                                    v130 := GenS13{F0: 128, F1: (fi26 - v25), F2: -227}
                                    _ = v129
                                    _ = v130
                                }
                            }
                            p1 = p1
                            var v131 uint8 = 255
                            var v132 bool = true
                            p1 = p1
                            _ = v131
                            _ = v132
                        }
                    }
                    if true {
                        fmt.Println(p1)
                        p1 = "a"
                        v133 := []int{-36}
                        switch "" {
                        case "\n\t":
                            var v134 int32 = -526
                            _ = v134
                        default:
                            v133 = append(v133, v25)
                        }
                        if ((!false) || false) {
                            var v135 float32 = -2.25
                            if (209 <= 128) {
                                v135 += v135
                                _ = len(v133)
                            } else {
                                switch "a" {
                                case "hello":
                                    v136 := GenS15{F0: -1, F1: 504, F2: (p1 + p1)}
                                    _ = v136
                                case "foo bar":
                                    var v137 int32 = 2147483647
                                    _ = v137
                                case "":
                                    fmt.Println(v135)
                                    if (!false) {
                                        switch p1 {
                                        case "\n\t":
                                            fmt.Println(fi26, fi39, fi39)
                                        case "hello":
                                            v25 -= fi39
                                        case "foo bar":
                                            p1 = ((p1 + ("" + "")) + p1)
                                        default:
                                            fmt.Println(p0, p0)
                                        }
                                        _ = len(v133)
                                        fmt.Println(v135)
                                    }
                                    var v138 int64 = p0
                                    _ = v138
                                }
                            }
                            v25 -= (v25 >> ((2) & 63))
                            _ = v135
                        }
                        _ = v133
                    } else {
                        v139 := GenS12{F0: fi26}
                        switch p1 {
                        case "\n\t":
                            p1 = (p1 + "a")
                        case "foo bar":
                            var v140 int = (fi39 + (v25 + fi39))
                            p1 = p1
                            if (!(true || (489 != 2))) {
                                var v141 int32 = -450
                                _ = v141
                            }
                            _ = v140
                        case "hello":
                            switch (p1 + p1) {
                            case "\n\t":
                                v142 := GenS2{F0: 1, F1: 65535, F2: 42}
                                _ = v142
                            case "a":
                                p1 = p1
                            case "":
                                var v143 int64 = 312
                                p1 = ""
                                _ = v143
                            default:
                                switch "foo bar" {
                                case "hello":
                                    if ((1e10 + float32(3.14159)) >= (0.1 / 1.0)) {
                                        fmt.Println(p1, p1)
                                    }
                                case "foo bar":
                                    v144 := GenS1{F0: 42, F1: -1, F2: (p1 + (p1 + p1))}
                                    _ = v144
                                }
                            }
                        }
                        v145 := GenS16{F0: 2, F1: 1}
                        _ = v139
                        _ = v145
                    }
                    switch v25 {
                    case 42:
                        p1 = p1
                    case -9223372036854775808:
                        if (((0.1 / float32(-2.25)) / 1.5) < -2.25) {
                            v146 := []int{fi39}
                            _ = v146
                        } else {
                            fmt.Println(v25, fi39)
                            var v147 int32 = 1
                            _ = v147
                        }
                    }
                    v148 := []int{(fi39 / ((127) | 1)), v25}
                    _ = fi39
                    _ = v148
                }
            }
        default:
            p0 = (p0 % ((-9223372036854775808) | 1))
        }
        _ = fi26
        _ = v27
    }
    _ = p0
    _ = p1
    _ = v25
    return 1
}
func genFn1(p0 uint32) uint32 {
    v149 := GenS17{F0: "hello", F1: 490, F2: 0}
    v150 := GenS10{F0: 32767, F1: 1000}
    if (true && false) {
        v151 := GenS18{F0: -9223372036854775808, F1: 42, F2: p0}
        _ = v151
    } else {
        var v152 float64 = 0.1
        var v153 uint8 = 196
        var v154 float32 = 1.5
        _ = v152
        _ = v153
        _ = v154
    }
    _ = p0
    _ = v149
    _ = v150
    return p0
}

func main() {
    var v1 int64 = 1
    var v2 uint8 = 2
    var v3 int64 = -1
    v4 := []string{"a", "\n\t", ((((("foo bar" + "a") + "") + "foo bar") + ((("\n\t" + "foo bar") + "foo bar") + "")) + "a")}
    v5 := []int{}
    fmt.Println(v2, v1, v1)
    var v155 int16 = 42
    _ = len(v4)
    v156 := GenS2{F0: (v155 >> ((255) & 63)), F1: 2, F2: 1}
    v157 := GenS4{F0: 42}
    var v158 uint8 = v2
    if (false && (!(1 != 128))) {
        for fi159 := 0; fi159 < 2; fi159++ {
            var v160 uint32 = genFn1(4294967295)
            if false {
                v161 := GenS1{F0: 568, F1: 489, F2: "\n\t"}
                _ = v161
            } else {
                if false {
                    v1 = (v3 | 0)
                    fmt.Println(v1)
                    for fi162 := 0; fi162 < 8; fi162++ {
                        switch fi159 {
                        case -349:
                            var v163 float32 = ((1e10 * (((float32(((float64(float32(1e10)) + ((3.14159 - ((-2.25 / 0.1) * (3.14159 / -2.25))) * 1e10)) / 0.1)) + 3.14159) * 0.0) * 1e10)) + (3.14159 * ((0.1 * ((0.0 / 1.0) * 0.1)) * 1.5)))
                            _ = v163
                        }
                        v164 := GenS16{F0: (v158 * 136), F1: 42}
                        var v165 uint64 = 0
                        if true {
                            switch "" {
                            case "\n\t":
                                v166 := GenS18{F0: (v3 | 0), F1: int32(v165), F2: (v160 + (v160 * (v160 | v160)))}
                                _ = v166
                            default:
                                _ = len(v4)
                            }
                        }
                        _ = fi162
                        _ = v164
                        _ = v165
                    }
                } else {
                    var v167 uint16 = 2
                    if (3.14159 == ((float64(0.1) + 3.14159) / (((1.5 * 0.0) / 1e10) - ((0.1 / 1.0) + float64(1.5))))) {
                        v155 = (v155 >> ((944) & 63))
                        v168 := GenS3{F0: v3, F1: 762, F2: 1}
                        _ = v168
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    var v169 int = (fi159 & fi159)
                    _ = v167
                    _ = v169
                }
                v155 = v155
                if (1 == 730) {
                    v5 = append(v5, fi159)
                }
            }
            if true {
                var v170 int = (fi159 | fi159)
                var v171 bool = false
                var v172 int64 = (v1 & int64(v160))
                _ = v170
                _ = v171
                _ = v172
            } else {
                v173 := GenS12{F0: (-(fi159))}
                _ = v173
            }
            _ = fi159
            _ = v160
        }
        if false {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    v174 := GenS4{F0: (v2 - 224)}
    v3 -= v1
    v3 += (v1 | -759)
    v158 += v158
    var v175 uint32 = 903
    if false {
        v3 -= -1
        for fi176 := 0; fi176 < 17; fi176++ {
            v1 += 0
            if true {
                var v177 int32 = int32(v155)
                for fi178 := 0; fi178 < 8; fi178++ {
                    v3 -= (v1 * -1)
                    var v179 int = (^(fi178))
                    _ = fi178
                    _ = v179
                }
                var v180 uint16 = 1
                _ = v177
                _ = v180
            } else {
                v181 := GenS19{F0: (^(fi176)), F1: (-(v1))}
                _ = v181
            }
            _ = fi176
        }
        var v182 int = 743
        _ = v182
    } else {
        for fi183 := 0; fi183 < 16; fi183++ {
            for fi184 := 0; fi184 < 5; fi184++ {
                if (((^(fi183)) < fi183) || false) {
                    var v185 int8 = 23
                    v1 += (v3 / (((v1 ^ v1)) | 1))
                    _ = v185
                }
                var v186 int32 = 2147483647
                if (true || true) {
                    if (1 < 2) {
                        v155 = (-(v155))
                    } else {
                        v175 = v175
                        var v187 bool = false
                        _ = v187
                    }
                    v188 := GenS20{F0: 21, F1: 2}
                    v189 := GenS21{F0: -9223372036854775808, F1: v175}
                    _ = v188
                    _ = v189
                } else {
                    v175 = 2147483648
                }
                _ = fi184
                _ = v186
            }
            v190 := GenS16{F0: v2, F1: 18446744073709551615}
            _ = fi183
            _ = v190
        }
        fmt.Println(v175, v2, v1)
    }
    v191 := GenS22{F0: 0, F1: (^(v2))}
    fmt.Println("v1:", v1)
    fmt.Println("v155:", v155)
    fmt.Println("v158:", v158)
    fmt.Println("v175:", v175)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v155
    _ = v156
    _ = v157
    _ = v158
    _ = v174
    _ = v175
    _ = v191
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
