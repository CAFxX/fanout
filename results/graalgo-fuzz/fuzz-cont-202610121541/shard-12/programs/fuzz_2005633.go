package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 int64; F2 int32 }
type GenS5 struct { F0 uint64; F1 uint32; F2 bool }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint; F1 string; F2 int8 }
type GenS8 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int8; F1 int32; F2 int16 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 uint8; F1 bool }
type GenS14 struct { F0 uint8; F1 uint; F2 int16 }
type GenS15 struct { F0 int32; F1 int32 }
type GenS16 struct { F0 int32; F1 int8; F2 uint64 }
type GenS17 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS18 struct { F0 uint8; F1 uint32 }
type GenS19 struct { F0 uint; F1 int64; F2 int64 }
func genFn0(p0 uint32, p1 bool, p2 int32) int {
    var v4 string = (("hello" + "") + ("foo bar" + "\n\t"))
    v5 := GenS1{F0: p0, F1: 32767, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 422
}
func genFn1(p0 int, p1 uint32) uint64 {
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 float32 = ((float32(1.5) * float32((1e10 * 1.5))) - -2.25)
        if false {
            if (256 <= 42) {
                for fi8 := 0; fi8 < 2; fi8++ {
                    if true {
                        v7 = float32(1e10)
                        var v9 int64 = 289
                        v9 += (-(v9))
                        _ = v9
                    }
                    _ = fi8
                }
                for fi10 := 0; fi10 < 17; fi10++ {
                    switch "a" {
                    case "foo bar":
                        if true {
                            p1 = p1
                            var v11 uint16 = 900
                            if false {
                                if false {
                                    var v12 uint8 = 94
                                    if true {
                                        v13 := GenS1{F0: p1, F1: (v11 * (v11 + (v11 | (v11 % ((32767) | 1))))), F2: 139}
                                        _ = v13
                                    } else {
                                        var v14 int16 = 559
                                        var v15 string = "foo bar"
                                        _ = v14
                                        _ = v15
                                    }
                                    _ = v12
                                }
                                var v16 int16 = -32768
                                _ = v16
                            } else {
                                v17 := GenS2{F0: v11, F1: true}
                                _ = v17
                            }
                            _ = v11
                        } else {
                            var v18 int16 = -496
                            _ = v18
                        }
                    case "":
                        var v19 int = fi6
                        var v20 int32 = 2147483647
                        v19 = v19
                        _ = v19
                        _ = v20
                    case "\n\t":
                        var v21 int32 = -788
                        _ = v21
                    default:
                        v22 := []int{genFn0(p1, (false || (((!false) || true) && (false || true))), 942)}
                        _ = v22
                    }
                    v23 := []int{}
                    _ = fi10
                    _ = v23
                }
            } else {
                var v24 int32 = 42
                _ = v24
            }
            if true {
                v25 := GenS3{F0: 885}
                if ((18446744073709551615 < 1) || true) {
                    for fi26 := 0; fi26 < 3; fi26++ {
                        var v27 uint = 4294967295
                        var v28 int = (p0 >> (((v27 / (((v27 % ((v27) | 1))) | 1))) & 63))
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                } else {
                    v29 := GenS1{F0: p1, F1: 42, F2: 42}
                    p0 = fi6
                    switch fi6 {
                    case 916:
                        if (4294967295 > 856) {
                            for fi30 := 0; fi30 < 19; fi30++ {
                                switch "\n\t" {
                                case "foo bar":
                                    p0 = 526
                                case "hello":
                                    if true {
                                        v31 := GenS3{F0: 0}
                                        _ = v31
                                    }
                                case "":
                                    if ((42 <= 2) || (true || true)) {
                                        if (32767 > 1) {
                                            var v32 string = "hello"
                                            v33 := GenS4{F0: 216, F1: 42, F2: 402}
                                            var v34 uint8 = uint8(p1)
                                            p0 -= (fi6 << ((694) & 63))
                                            _ = v32
                                            _ = v33
                                            _ = v34
                                        }
                                        v35 := []int{}
                                        _ = v35
                                    } else {
                                        if true {
                                            if (255 != 4294967295) {
                                                var v36 int16 = 176
                                                _ = v36
                                            } else {
                                                p1 = 67
                                                var v37 uint64 = 2
                                                _ = v37
                                            }
                                            v38 := []int{p0, p0, (^(fi6))}
                                            var v39 float32 = -2.25
                                            var v40 uint64 = 18446744073709551615
                                            switch "hello" {
                                            case "a":
                                                p1 += (p1 + p1)
                                            case "\n\t":
                                                var v41 uint64 = v40
                                                _ = v41
                                            case "":
                                                var v42 int = fi30
                                                _ = v42
                                            }
                                            _ = v38
                                            _ = v39
                                            _ = v40
                                        }
                                    }
                                }
                                switch "foo bar" {
                                case "\n\t":
                                    var v43 bool = (42 == -2147483648)
                                    _ = v43
                                }
                                var v44 bool = (!(("a" == "hello") && ((100 < -40) || (true && false))))
                                if v44 {
                                    v45 := GenS5{F0: 4294967295, F1: (^(p1)), F2: (!v44)}
                                    var v46 uint64 = 619
                                    if v44 {
                                        v46 += v46
                                    }
                                    if (v7 >= 3.14159) {
                                        switch "hello" {
                                        case "hello":
                                            var v47 int8 = 0
                                            _ = v47
                                        case "\n\t":
                                            switch (p0 * fi30) {
                                            case -9223372036854775808:
                                                switch "a" {
                                                case "\n\t":
                                                    var v48 bool = v44
                                                    _ = v48
                                                }
                                            case -2147483648:
                                                v7 += ((v7 * 3.14159) * 0.1)
                                            }
                                        case "":
                                            var v49 uint16 = 20
                                            _ = v49
                                        }
                                        if (-104 < -128) {
                                            if (-9223372036854775808 <= 593) {
                                                v50 := GenS6{F0: 2147483648}
                                                if v44 {
                                                    fmt.Println(fi30, v46, v7)
                                                } else {
                                                    if false {
                                                        if v44 {
                                                            switch (-(fi30)) {
                                                            case 9223372036854775807:
                                                                fmt.Println(v46, v7, v46)
                                                            case -1:
                                                                v51 := []int{p0, (p0 / (((p0 + (fi6 % (((fi6 >> ((255) & 63))) | 1)))) | 1))}
                                                                _ = v51
                                                            case 127:
                                                                var v52 uint64 = 4294967295
                                                                _ = v52
                                                            }
                                                        }
                                                        p0 -= p0
                                                        if (v44 || v44) {
                                                            p0 += fi30
                                                        } else {
                                                            switch (p0 / ((fi30) | 1)) {
                                                            case -9223372036854775808:
                                                                if v44 {
                                                                    if false {
                                                                        if (v44 || v44) {
                                                                            v53 := []int{p0, fi30, 199, -643}
                                                                            _ = v53
                                                                        } else {
                                                                            fmt.Println(v7, fi6)
                                                                            p0 += genFn0((p1 ^ 4294967295), v44, -1)
                                                                            fmt.Println(fi6, fi6, fi30)
                                                                            v54 := GenS7{F0: 42, F1: "\n\t", F2: -82}
                                                                            switch p0 {
                                                                            case 660:
                                                                                var v55 float64 = 1.5
                                                                                _ = v55
                                                                            default:
                                                                                v56 := []int{(^(fi6)), fi6}
                                                                                _ = v56
                                                                            }
                                                                            _ = v54
                                                                        }
                                                                    }
                                                                    if true {
                                                                        var v57 uint32 = p1
                                                                        _ = v57
                                                                    }
                                                                }
                                                            }
                                                            var v58 int64 = int64(v46)
                                                            _ = v58
                                                        }
                                                    }
                                                    p1 = 4294967295
                                                    var v59 uint = 255
                                                    if true {
                                                        var v60 int16 = 0
                                                        _ = v60
                                                    }
                                                    v61 := GenS8{F0: p1, F1: 32767, F2: 1000}
                                                    _ = v59
                                                    _ = v61
                                                }
                                                if v44 {
                                                    v62 := GenS1{F0: p1, F1: 42, F2: v46}
                                                    _ = v62
                                                }
                                                _ = v50
                                            } else {
                                                v7 = -2.25
                                                var v63 int32 = 0
                                                var v64 uint8 = 2
                                                fmt.Println(v64, v44, fi30)
                                                _ = v63
                                                _ = v64
                                            }
                                            if (128 > 0) {
                                                switch "foo bar" {
                                                case "hello":
                                                    v65 := []int{-707, fi30}
                                                    _ = v65
                                                case "\n\t":
                                                    if v44 {
                                                        p1 = (p1 >> ((uint(v46)) & 63))
                                                    }
                                                }
                                            } else {
                                                if v44 {
                                                    v44 = ((p1 < (p1 - p1)) || v44)
                                                    var v66 float64 = 0.1
                                                    var v67 bool = ((1 >= 0) && (42 > 9223372036854775807))
                                                    _ = v66
                                                    _ = v67
                                                } else {
                                                    if v44 {
                                                        if (false || v44) {
                                                            var v68 uint8 = 255
                                                            _ = v68
                                                        } else {
                                                            switch "a" {
                                                            case "a":
                                                                var v69 uint = 42
                                                                _ = v69
                                                            case "hello":
                                                                if (9223372036854775807 != 0) {
                                                                    fmt.Println(v44)
                                                                    var v70 int32 = -1
                                                                    _ = v70
                                                                } else {
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        if v44 {
                                                                            if (2 > 255) {
                                                                                fmt.Println(v46, fi6, fi6)
                                                                                fmt.Println(v44, fi30, p0)
                                                                                v71 := GenS9{F0: -101}
                                                                                _ = v71
                                                                            } else {
                                                                                v72 := GenS4{F0: (v46 << ((255) & 63)), F1: 682, F2: 0}
                                                                                v73 := []int{fi6, -1, fi30}
                                                                                v7 = v7
                                                                                _ = v72
                                                                                _ = v73
                                                                            }
                                                                            p0 = p0
                                                                            if false {
                                                                                fmt.Println(v44, fi6)
                                                                            }
                                                                        }
                                                                    case "":
                                                                        v74 := GenS10{F0: 255}
                                                                        switch "" {
                                                                        case "":
                                                                            v44 = (131 < (v46 - v46))
                                                                        case "hello":
                                                                            p0 = (fi6 % ((fi30) | 1))
                                                                        }
                                                                        v75 := GenS6{F0: (^(p1))}
                                                                        _ = v74
                                                                        _ = v75
                                                                    default:
                                                                        var v76 uint8 = 2
                                                                        _ = v76
                                                                    }
                                                                    v77 := GenS2{F0: 0, F1: (!(0 > 100))}
                                                                    _ = v77
                                                                }
                                                            case "":
                                                                var v78 uint64 = (v46 | (v46 % ((v46) | 1)))
                                                                _ = v78
                                                            }
                                                        }
                                                    }
                                                    v79 := []int{(fi6 & 265), (fi30 & p0), 127, genFn0(p1, v44, -2147483648)}
                                                    if v44 {
                                                        fmt.Println(v7, fi30, v46)
                                                    } else {
                                                        if (false || (42 < -1)) {
                                                            v79 = append(v79, -409)
                                                            v46 = 433
                                                        }
                                                    }
                                                    _ = v79
                                                }
                                            }
                                        }
                                        if v44 {
                                            var v80 int8 = 1
                                            if (v80 <= (-(v80))) {
                                                if (42 >= 0) {
                                                    p0 = p0
                                                } else {
                                                    switch ((("hello" + ((("\n\t" + "") + "\n\t") + (("a" + "") + "hello"))) + ("foo bar" + "\n\t")) + "hello") {
                                                    case "a":
                                                        switch -453 {
                                                        case -558:
                                                            p0 -= -681
                                                        case 32767:
                                                            var v81 float64 = 0.1
                                                            _ = v81
                                                        }
                                                    case "foo bar":
                                                        v82 := []int{466, (fi6 | fi6), (fi30 + fi30)}
                                                        _ = v82
                                                    default:
                                                        if (233 > 1000) {
                                                            v80 = (v80 ^ v80)
                                                            var v83 float64 = 0.1
                                                            var v84 int8 = (v80 << ((0) & 63))
                                                            var v85 uint8 = 156
                                                            var v86 uint = 0
                                                            if (0 == int16(v86)) {
                                                                var v87 uint32 = p1
                                                                switch (fi30 % ((fi30) | 1)) {
                                                                case 42:
                                                                    v88 := GenS11{F0: (v84 % ((v80) | 1)), F1: 1, F2: 32767}
                                                                    _ = v88
                                                                case -2147483648:
                                                                    v85 += (v85 ^ v85)
                                                                default:
                                                                    v89 := GenS3{F0: 9223372036854775807}
                                                                    _ = v89
                                                                }
                                                                if (v44 || (65535 < (v87 * 65535))) {
                                                                    var v90 float32 = ((v7 + ((v7 - 0.1) * float32(v83))) / 1e10)
                                                                    _ = v90
                                                                } else {
                                                                    v91 := []int{(fi6 % ((fi6) | 1)), (fi30 ^ fi6)}
                                                                    if (!(32767 > 32767)) {
                                                                        switch "\n\t" {
                                                                        case "":
                                                                            if len(v91) > 0 {
                                                                                v91[0] = genFn0((v87 - 500), v44, 42)
                                                                            }
                                                                        case "a":
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                p1 = p1
                                                                            case "a":
                                                                                var v92 string = ("foo bar" + "")
                                                                                _ = v92
                                                                            case "hello":
                                                                                v91 = append(v91, (^(fi6)))
                                                                            }
                                                                        case "hello":
                                                                            var v93 uint = (^(v86))
                                                                            _ = v93
                                                                        default:
                                                                            switch p0 {
                                                                            case 127:
                                                                                if ((^(p1)) > 0) {
                                                                                    if ((-(p0)) != (fi30 >> ((v86) & 63))) {
                                                                                        v94 := GenS12{F0: -2147483648}
                                                                                        fmt.Println(v85)
                                                                                        v80 += v80
                                                                                        _ = v94
                                                                                    }
                                                                                    v7 -= float32(v83)
                                                                                    var v95 int16 = 0
                                                                                    var v96 float32 = (1e10 - -2.25)
                                                                                    _ = v95
                                                                                    _ = v96
                                                                                } else {
                                                                                    v44 = false
                                                                                    if (((v85 * 161) >= (^(v85))) && ((v44 && ((true || true) || v44)) || true)) {
                                                                                        _ = len(v91)
                                                                                        if v44 {
                                                                                            v44 = (!(!(v44 || true)))
                                                                                            v83 -= v83
                                                                                            v97 := GenS13{F0: uint8(v46), F1: (2147483647 < int32(fi30))}
                                                                                            _ = v97
                                                                                        } else {
                                                                                            v83 = v83
                                                                                            v80 += (v80 >> ((v86) & 63))
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case 2147483647:
                                                                                switch "\n\t" {
                                                                                case "a":
                                                                                    var v98 uint32 = (v87 & 1)
                                                                                    _ = v98
                                                                                case "foo bar":
                                                                                    switch "foo bar" {
                                                                                    case "hello":
                                                                                        var v99 int32 = -667
                                                                                        _ = v99
                                                                                    case "":
                                                                                        switch ("foo bar" + "hello") {
                                                                                        case "\n\t":
                                                                                            v83 -= 1.5
                                                                                        case "":
                                                                                            if len(v91) >= 2 {
                                                                                                v91 = v91[1:len(v91)]
                                                                                            }
                                                                                        default:
                                                                                            v46 += (v46 << ((1) & 63))
                                                                                        }
                                                                                    }
                                                                                case "hello":
                                                                                    if v44 {
                                                                                        var v100 uint64 = (v46 & uint64(v85))
                                                                                        v100 += (v46 - v100)
                                                                                        var v101 float32 = v7
                                                                                        _ = v100
                                                                                        _ = v101
                                                                                    } else {
                                                                                        var v102 int32 = 2147483647
                                                                                        if (v85 > v85) {
                                                                                            var v103 uint8 = 128
                                                                                            v7 = ((v7 / 1.0) + ((-2.25 / v7) - v7))
                                                                                            _ = v103
                                                                                        }
                                                                                        var v104 uint = v86
                                                                                        _ = v102
                                                                                        _ = v104
                                                                                    }
                                                                                default:
                                                                                    var v105 string = "foo bar"
                                                                                    _ = v105
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if v44 {
                                                                            p1 = p1
                                                                        } else {
                                                                            var v106 int = p0
                                                                            _ = v106
                                                                        }
                                                                    }
                                                                    _ = v91
                                                                }
                                                                _ = v87
                                                            } else {
                                                                fmt.Println(p1, fi30)
                                                            }
                                                            _ = v83
                                                            _ = v84
                                                            _ = v85
                                                            _ = v86
                                                        } else {
                                                            var v107 int8 = -7
                                                            p0 = genFn0(686, (567 < 255), 42)
                                                            if (!(!v44)) {
                                                                var v108 bool = (!("hello" <= "foo bar"))
                                                                _ = v108
                                                            } else {
                                                                v109 := GenS9{F0: (v80 / ((v107) | 1))}
                                                                var v110 int32 = -851
                                                                _ = v109
                                                                _ = v110
                                                            }
                                                            _ = v107
                                                        }
                                                    }
                                                }
                                                fmt.Println(v46, v80)
                                                fmt.Println(fi30)
                                            }
                                            if ((^(v46)) >= (^(v46))) {
                                                v7 += (1e10 - v7)
                                            } else {
                                                switch (fi30 | fi6) {
                                                case 9223372036854775807:
                                                    v46 = (v46 ^ (v46 % ((2) | 1)))
                                                default:
                                                    v111 := GenS5{F0: (v46 * v46), F1: (p1 ^ (p1 << ((65535) & 63))), F2: true}
                                                    _ = v111
                                                }
                                                var v112 string = "foo bar"
                                                v7 = float32(0.1)
                                                _ = v112
                                            }
                                            _ = v80
                                        } else {
                                            v44 = v44
                                            if false {
                                                var v113 int8 = 120
                                                _ = v113
                                            }
                                        }
                                    } else {
                                        if ((9223372036854775807 < 1) || v44) {
                                            v114 := GenS14{F0: 140, F1: 255, F2: 694}
                                            v115 := GenS15{F0: int32(p1), F1: int32(fi6)}
                                            if (!("foo bar" > "foo bar")) {
                                                v7 = (float32(((1e10 + (((float64(v7) - -2.25) - 3.14159) - float64(v7))) / (-2.25 + -2.25))) / ((((v7 * 3.14159) - v7) * v7) * 3.14159))
                                            }
                                            _ = v114
                                            _ = v115
                                        } else {
                                            if (42 > -346) {
                                                var v116 uint8 = 2
                                                v117 := []int{fi6, (p0 * fi6), 992, genFn0((^(p1)), ((v44 && v44) && v44), 398)}
                                                _ = v116
                                                _ = v117
                                            }
                                        }
                                        p1 = 913
                                    }
                                    _ = v45
                                    _ = v46
                                }
                                var v118 bool = (p1 > (p1 - 4294967295))
                                _ = fi30
                                _ = v118
                                _ = v44
                            }
                            if true {
                                var v119 int64 = -162
                                switch "a" {
                                case "\n\t":
                                    p1 = 4294967295
                                }
                                _ = v119
                            } else {
                                p0 = p0
                            }
                            for fi120 := 0; fi120 < 11; fi120++ {
                                var v121 uint64 = 448
                                var v122 uint64 = (v121 * 0)
                                _ = fi120
                                _ = v121
                                _ = v122
                            }
                        } else {
                            for fi123 := 0; fi123 < 1; fi123++ {
                                var v124 bool = false
                                v125 := GenS14{F0: 1, F1: 215, F2: -809}
                                v126 := GenS7{F0: 2, F1: ((("hello" + "\n\t") + (("" + (("\n\t" + "hello") + "hello")) + "\n\t")) + "foo bar"), F2: 127}
                                v127 := GenS16{F0: 2147483647, F1: -1, F2: 1}
                                var v128 uint16 = 0
                                var v129 string = ("foo bar" + "hello")
                                v128 = 150
                                _ = fi123
                                _ = v124
                                _ = v125
                                _ = v126
                                _ = v127
                                _ = v128
                                _ = v129
                            }
                        }
                    }
                    _ = v29
                }
                _ = v25
            } else {
                if false {
                    v130 := GenS17{F0: (-1 >= 32767), F1: 550, F2: 54}
                    var v131 float64 = (3.14159 - float64(3.14159))
                    _ = v130
                    _ = v131
                } else {
                    fmt.Println(p0, p1, v7)
                    p1 = 2147483647
                    v132 := []int{(fi6 & -9223372036854775808), genFn0((p1 >> ((348) & 63)), (!false), -1)}
                    _ = v132
                }
                switch "foo bar" {
                case "a":
                    v133 := GenS9{F0: 1}
                    _ = v133
                case "foo bar":
                    if (-1 >= 32767) {
                        var v134 uint8 = 42
                        var v135 uint8 = (v134 % ((v134) | 1))
                        _ = v134
                        _ = v135
                    }
                case "hello":
                    v136 := GenS18{F0: 42, F1: p1}
                    _ = v136
                }
                p0 = fi6
            }
        } else {
            v137 := []int{(fi6 | p0), fi6, 128}
            var v138 uint = 65535
            _ = v137
            _ = v138
        }
        _ = fi6
        _ = v7
    }
    v139 := []int{(^(p0))}
    v140 := GenS17{F0: (51 != 127), F1: 1, F2: 0}
    _ = p0
    _ = p1
    _ = v139
    _ = v140
    return 463
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{(("" + "") + (("foo bar" + "a") + "hello"))}
    var v3 uint = 2
    v141 := GenS19{F0: (^(v3)), F1: -207, F2: -982}
    var v142 uint32 = 0
    v3 = (^(v3))
    v142 = 590
    if len(v2) > 0 {
        v2[0] = (("a" + "") + "\n\t")
    }
    v3 -= v3
    if (128 > 1) {
        v142 = 2147483647
    } else {
        v142 = v142
        var v143 int32 = -234
        _ = v143
    }
    for fi144 := 0; fi144 < 15; fi144++ {
        if true {
            v3 = v3
            delete(v1, "")
        } else {
            v142 -= (v142 / ((v142) | 1))
        }
        _ = fi144
    }
    if true {
        var v145 bool = false
        v146 := GenS1{F0: (v142 - 2), F1: 42, F2: 2}
        _ = v145
        _ = v146
    }
    if (0 > 1) {
        if true {
            fmt.Println(v142, v3)
            fmt.Println(v142)
        } else {
            v1[("a" + "foo bar")] = (v142 - 373)
        }
        fmt.Println(v142, v3)
        var v147 int64 = -50
        _ = v147
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v142:", v142)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v141
    _ = v142
    _ = v2
    _ = v3
}
