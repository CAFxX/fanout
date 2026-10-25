package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int16; F1 int32; F2 int32 }
type GenS3 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int32; F1 int64; F2 int8 }
type GenS8 struct { F0 int64; F1 uint64; F2 int16 }
type GenS9 struct { F0 uint16; F1 int64 }
type GenS10 struct { F0 int16; F1 uint8; F2 int64 }
type GenS11 struct { F0 uint64; F1 uint; F2 int }
type GenS12 struct { F0 int64; F1 int }
func genFn0() int {
    v6 := GenS1{F0: ((2147483648 < 105) || (true || false)), F1: 42, F2: 227}
    v7 := []int{-1, -665, 54}
    for fi8 := 0; fi8 < 11; fi8++ {
        if (true || (9223372036854775807 < 42)) {
            if false {
                for fi9 := 0; fi9 < 17; fi9++ {
                    fmt.Println(fi8, fi9)
                    var v10 float32 = (1e10 - (1e10 / 1.0))
                    _ = fi9
                    _ = v10
                }
            } else {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                v11 := GenS1{F0: false, F1: 42, F2: 0}
                _ = v11
            }
            v12 := GenS2{F0: -294, F1: -1, F2: 0}
            v7 = append(v7, fi8)
            _ = v12
        } else {
            fmt.Println(fi8)
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 uint64 = 594
                _ = fi13
                _ = v14
            }
        }
        _ = fi8
    }
    if len(v7) > 0 {
        v7[0] = 2147483647
    }
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint8 = 127
    var v2 uint = 2
    var v3 int64 = 1
    v4 := make(map[string]bool)
    v5 := make(map[string]int)
    var v15 uint = (v2 * 675)
    if (!(true || ((!true) || true))) {
        v3 = v3
        v16 := GenS1{F0: false, F1: v1, F2: uint16(v1)}
        _ = v16
    }
    if ((int16(v3) < 0) && ((!true) || ((^(v15)) < 18446744073709551615))) {
        switch 0 {
        case -9:
            for fi17 := 0; fi17 < 15; fi17++ {
                if false {
                    v1 += (v1 | (v1 - (v1 & v1)))
                    if false {
                        var v18 int8 = 42
                        _ = v18
                    } else {
                        var v19 int16 = 607
                        _ = v19
                    }
                    v20 := GenS3{F0: 2147483648, F1: 2, F2: 1}
                    _ = v20
                } else {
                    if (v3 < -1) {
                        v21 := []int{}
                        _ = v21
                    } else {
                        if false {
                            var v22 int = fi17
                            _ = v22
                        } else {
                            v3 -= v3
                        }
                    }
                    var v23 uint64 = 18446744073709551615
                    switch "foo bar" {
                    case "a":
                        var v24 int8 = -128
                        _ = v24
                    case "hello":
                        v1 = 31
                    case "foo bar":
                        var v25 uint8 = v1
                        _ = v25
                    }
                    _ = v23
                }
                _ = fi17
            }
        case 9223372036854775807:
            fmt.Println(v3, v1, v15)
        case -235:
            for fi26 := 0; fi26 < 2; fi26++ {
                for fi27 := 0; fi27 < 20; fi27++ {
                    fmt.Println(fi27, v3, v15)
                    fmt.Println(v15, v15, fi27)
                    var v28 uint8 = v1
                    _ = fi27
                    _ = v28
                }
                if true {
                    v29 := GenS3{F0: 2147483648, F1: 989, F2: -2147483648}
                    if false {
                        for fi30 := 0; fi30 < 16; fi30++ {
                            var v31 float64 = -2.25
                            v32 := GenS3{F0: 0, F1: 42, F2: 682}
                            v33 := GenS1{F0: false, F1: v1, F2: 255}
                            _ = fi30
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                        var v34 int64 = (v3 >> ((2) & 63))
                        var v35 uint16 = 255
                        _ = v34
                        _ = v35
                    }
                    _ = v29
                }
                var v36 uint32 = 2
                _ = fi26
                _ = v36
            }
        }
    } else {
        v37 := []int{genFn0()}
        if (0.1 != 1.5) {
            var v38 bool = false
            _ = v38
        }
        if len(v37) > 0 {
            v37[0] = -1
        }
        _ = v37
    }
    v1 = 127
    for fi39 := 0; fi39 < 9; fi39++ {
        switch "" {
        case "":
            v2 += 198
        case "\n\t":
            v1 -= v1
        case "hello":
            v40 := []int{}
            _ = v40
        }
        _ = fi39
    }
    var v41 uint32 = 4294967295
    v42 := GenS4{F0: -2147483648, F1: -81}
    if (!((!true) || true)) {
        fmt.Println(v2)
    } else {
        v43 := []int{128, 2147483647, 0, 519}
        _ = v43
    }
    fmt.Println(v2, v41, v1)
    v44 := GenS5{F0: ("" + ""), F1: 0}
    for fi45 := 0; fi45 < 1; fi45++ {
        v46 := []int{}
        var v47 float64 = 1e10
        if len(v46) > 0 {
            v46[0] = genFn0()
        }
        _ = fi45
        _ = v46
        _ = v47
    }
    v3 = v3
    v3 -= -526
    if false {
        v48 := []int{0, 4294967295, -2147483648}
        _ = v48
    }
    if true {
        switch 187 {
        case 9223372036854775807:
            fmt.Println(v41, v15, v3)
        }
        for fi49 := 0; fi49 < 16; fi49++ {
            for fi50 := 0; fi50 < 18; fi50++ {
                switch fi50 {
                case -1:
                    switch "\n\t" {
                    case "\n\t":
                        if (885 == 374) {
                            v51 := GenS5{F0: "a", F1: (v41 - v41)}
                            _ = v51
                        }
                    default:
                        fmt.Println(v2, v2)
                    }
                case 32768:
                    if (((3.14159 - 3.14159) == 0.1) || false) {
                        var v52 uint8 = (v1 & 0)
                        fmt.Println(v15, v3, fi50)
                        v1 = uint8(v15)
                        switch genFn0() {
                        case 42:
                            switch "\n\t" {
                            case "a":
                                v15 -= v2
                            case "":
                                fmt.Println(fi49, v52, v15)
                            default:
                                if (true || (32767 == uint16(v1))) {
                                    v1 += (v52 / (((v1 ^ uint8(v15))) | 1))
                                    var v53 int = (fi49 - (fi50 % ((fi49) | 1)))
                                    var v54 int8 = 100
                                    _ = v53
                                    _ = v54
                                } else {
                                    v52 = v52
                                    _ = len(v5)
                                    if false {
                                        var v55 int16 = 42
                                        if (false || (genFn0() != genFn0())) {
                                            switch -2147483648 {
                                            case 2147483647:
                                                var v56 float32 = 3.14159
                                                _ = v56
                                            case 42:
                                                if ((-1 == 25) && (0.1 != 1e10)) {
                                                    var v57 int32 = -1
                                                    v52 -= v1
                                                    _ = v57
                                                }
                                            case 128:
                                                if false {
                                                    switch ("a" + "hello") {
                                                    case "\n\t":
                                                        v2 += v15
                                                    case "hello":
                                                        v41 = v41
                                                    default:
                                                        switch fi50 {
                                                        case 127:
                                                            var v58 float32 = 0.0
                                                            _ = v58
                                                        case 42:
                                                            v3 -= 1
                                                        default:
                                                            if ((v41 / ((705) | 1)) == (v41 & uint32(v1))) {
                                                                v41 = (^(v41))
                                                                var v59 int32 = 1
                                                                _ = v59
                                                            } else {
                                                                var v60 int = (fi49 & fi50)
                                                                fmt.Println(fi50, v60)
                                                                var v61 int16 = v55
                                                                _ = v60
                                                                _ = v61
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(v3, v1, v55)
                                                    v62 := GenS6{F0: 42}
                                                    _ = v62
                                                }
                                            }
                                            fmt.Println(fi49, v15)
                                            v52 += v1
                                        } else {
                                            v41 = 65535
                                            v41 = 122
                                            v52 -= v52
                                        }
                                        _ = v55
                                    }
                                }
                            }
                        case -9223372036854775808:
                            var v63 uint16 = uint16(v15)
                            _ = v63
                        case -16:
                            switch "" {
                            case "":
                                if false {
                                    switch fi50 {
                                    case 32767:
                                        if (2 > 823) {
                                            v2 -= v15
                                            switch "a" {
                                            case "foo bar":
                                                v2 += 1
                                            case "a":
                                                if (2 >= 1) {
                                                    if (v1 >= (v1 ^ v52)) {
                                                        v64 := GenS7{F0: 1, F1: 959, F2: 78}
                                                        _ = v64
                                                    }
                                                }
                                            case "":
                                                var v65 int = (fi49 << ((v15) & 63))
                                                _ = v65
                                            default:
                                                v66 := GenS2{F0: -32768, F1: -1, F2: -1}
                                                _ = v66
                                            }
                                            if (((1e10 != 1e10) && false) && false) {
                                                v15 -= (v2 & (v2 + (v2 & (^(v15)))))
                                                var v67 int8 = 100
                                                var v68 int32 = -1
                                                _ = v67
                                                _ = v68
                                            } else {
                                                if (float64(float32(-2.25)) > 0.1) {
                                                    if ((^(v3)) > 963) {
                                                        v69 := []int{genFn0(), genFn0(), fi50}
                                                        fmt.Println(v15, v2)
                                                        _ = v69
                                                    }
                                                    var v70 int8 = 100
                                                    v71 := GenS2{F0: 0, F1: int32(v15), F2: -1}
                                                    _ = v70
                                                    _ = v71
                                                } else {
                                                    v72 := GenS7{F0: -398, F1: v3, F2: 57}
                                                    _ = v72
                                                }
                                                if false {
                                                    v15 = v15
                                                }
                                            }
                                            v73 := []int{(^(fi49)), (-(fi50))}
                                            _ = v73
                                        } else {
                                            var v74 string = "a"
                                            if false {
                                                var v75 uint64 = 0
                                                _ = len(v4)
                                                _ = v75
                                            }
                                            _ = v74
                                        }
                                    }
                                    v76 := GenS8{F0: 0, F1: 4294967295, F2: 0}
                                    v77 := GenS7{F0: -2147483648, F1: 9223372036854775807, F2: 100}
                                    _ = v76
                                    _ = v77
                                }
                            case "\n\t":
                                v78 := []int{}
                                _ = v78
                            case "hello":
                                fmt.Println(v41, fi49, v52)
                            }
                            fmt.Println(v52)
                            switch "foo bar" {
                            case "\n\t":
                                if false {
                                    v52 = (v1 ^ 1)
                                    fmt.Println(v41)
                                    v2 = (^(v2))
                                }
                            }
                            var v79 int16 = -32768
                            _ = v79
                        }
                        _ = v52
                    }
                default:
                    var v80 float32 = ((1.5 / 0.1) * (0.1 / ((3.14159 + 1e10) - (3.14159 * 0.1))))
                    _ = v80
                }
                v2 += v2
                if false {
                    v2 += 0
                    v81 := []int{fi50, (fi50 | (fi49 & (fi49 % ((fi49) | 1)))), 559, fi50}
                    _ = v81
                } else {
                    if (-360 <= -1) {
                        v82 := GenS1{F0: (true || (false && (!(false || (!((v1 < (^(v1))) && (-1 == 0))))))), F1: (v1 ^ (v1 & (v1 + (^(v1))))), F2: uint16(v15)}
                        v83 := []int{(^(fi50)), genFn0(), (^(fi50))}
                        _ = v82
                        _ = v83
                    }
                    if true {
                        if ((v3 * v3) <= v3) {
                            var v84 int8 = 100
                            v85 := []int{9223372036854775807, fi50, fi50, (-(fi50))}
                            fmt.Println(v3, v15, v2)
                            _ = v84
                            _ = v85
                        } else {
                            v86 := GenS8{F0: -1, F1: 0, F2: -1}
                            switch ("a" + "a") {
                            case "\n\t":
                                if true {
                                    switch 32768 {
                                    case -2147483648:
                                        v1 += v1
                                    default:
                                        var v87 string = (("" + "foo bar") + ("foo bar" + "\n\t"))
                                        _ = v87
                                    }
                                    v2 -= 65535
                                    var v88 int8 = 80
                                    v89 := []int{fi50, 223, fi50}
                                    switch (fi49 / ((fi50) | 1)) {
                                    case -1:
                                        if (false && (!(true && true))) {
                                            v90 := GenS8{F0: (v3 | -576), F1: 2, F2: -1}
                                            var v91 int64 = (^(v3))
                                            v88 -= (v88 * 100)
                                            _ = v90
                                            _ = v91
                                        }
                                    case 62:
                                        var v92 uint64 = 0
                                        _ = v92
                                    case 509:
                                        v41 = (v41 * (v41 % ((2) | 1)))
                                        var v93 uint8 = (v1 | (v1 | (v1 % ((42) | 1))))
                                        v89 = append(v89, (fi49 / (((fi49 - 128)) | 1)))
                                        switch "hello" {
                                        case "foo bar":
                                            var v94 uint = v2
                                            _ = v94
                                        default:
                                            v95 := GenS9{F0: 2, F1: (v3 | v3)}
                                            _ = v95
                                        }
                                        _ = len(v89)
                                        v96 := v4[(("a" + "hello") + "\n\t")]
                                        _ = v93
                                        _ = v96
                                    default:
                                        v97 := GenS10{F0: 15, F1: v1, F2: (v3 & v3)}
                                        _ = v97
                                    }
                                    _ = v88
                                    _ = v89
                                } else {
                                    v98 := []int{fi50, genFn0()}
                                    v3 -= v3
                                    _ = v98
                                }
                            }
                            _ = v86
                        }
                        if ((float32((0.0 - 0.0)) - ((1e10 - (((0.0 / 0.1) - 0.1) - 0.1)) + float32(1e10))) >= 3.14159) {
                            if (!(false || true)) {
                                fmt.Println(v3, v15)
                                if ((v2 != v15) || ((v3 != 1) && false)) {
                                    var v99 float64 = (1e10 / 1.0)
                                    v100 := GenS11{F0: 4294967295, F1: v15, F2: (fi50 + (-(fi50)))}
                                    if (!false) {
                                        v101 := []int{(fi49 * 32768), (fi50 << (((^(v2))) & 63)), (-(fi50))}
                                        fmt.Println(fi50, fi50)
                                        if true {
                                            v1 = (v1 | (^(v1)))
                                            v4["a"] = (true || false)
                                            v41 += v41
                                        } else {
                                            switch ("" + "a") {
                                            case "a":
                                                fmt.Println(fi50)
                                            case "foo bar":
                                                _ = len(v5)
                                            case "\n\t":
                                                v41 = v41
                                            default:
                                                switch (fi49 & (-(fi50))) {
                                                case -44:
                                                    v15 = 1
                                                    v99 = (1.5 + (v99 + (v99 - v99)))
                                                case -334:
                                                    var v102 int16 = 32767
                                                    _ = v102
                                                }
                                            }
                                            v1 -= (v1 % (((^(v1))) | 1))
                                        }
                                        _ = v101
                                    }
                                    _ = v100
                                    _ = v99
                                } else {
                                    v103 := GenS9{F0: 2, F1: 343}
                                    v104 := v5["foo bar"]
                                    if false {
                                        v2 = 574
                                        v105 := []int{fi50, -2147483648, (v104 - v104)}
                                        var v106 uint = 348
                                        v15 += (v106 ^ 518)
                                        var v107 uint32 = (v41 ^ v41)
                                        _ = v105
                                        _ = v106
                                        _ = v107
                                    } else {
                                        var v108 int16 = 32767
                                        _ = v108
                                    }
                                    _ = v103
                                    _ = v104
                                }
                            }
                            v109 := v5["\n\t"]
                            if (true && true) {
                                v41 -= 4294967295
                                fmt.Println(v15, v1, v15)
                                v110 := []int{(v109 - 20), fi49, fi50, -658}
                                var v111 bool = ((false || (128 == v1)) && (38 == -364))
                                _ = v110
                                _ = v111
                            } else {
                                v41 = (^(v41))
                            }
                            _ = v109
                        } else {
                            switch "" {
                            case "":
                                delete(v4, "hello")
                                var v112 int64 = 1
                                _ = v112
                            case "foo bar":
                                var v113 string = "a"
                                _ = v113
                            }
                            var v114 int16 = 32767
                            v115 := []int{127, (fi49 ^ -2147483648)}
                            var v116 int32 = 333
                            v2 = v15
                            var v117 float64 = float64(-2.25)
                            fmt.Println(v3)
                            switch genFn0() {
                            case 2147483647:
                                var v118 int32 = 2147483647
                                _ = v118
                            case 674:
                                switch int(v41) {
                                case -546:
                                    v2 = 939
                                }
                            }
                            _ = v114
                            _ = v115
                            _ = v116
                            _ = v117
                        }
                        var v119 uint16 = 255
                        _ = v119
                    } else {
                        switch 127 {
                        case -951:
                            var v120 uint64 = 152
                            _ = v120
                        default:
                            v41 = v41
                        }
                        switch "hello" {
                        case "\n\t":
                            if true {
                                _ = len(v5)
                                if false {
                                    if (uint(v41) == 421) {
                                        delete(v5, "hello")
                                        v3 += v3
                                        v15 = 18446744073709551615
                                    }
                                    var v121 int = (fi49 << (((v2 - v15)) & 63))
                                    v122 := GenS7{F0: 2147483647, F1: 0, F2: -88}
                                    _ = v121
                                    _ = v122
                                }
                                v41 = 0
                            }
                        case "hello":
                            if (2 > 340) {
                                v15 -= v15
                                v123 := GenS6{F0: (fi49 * (^(fi50)))}
                                _ = v123
                            } else {
                                v2 = (v2 & v15)
                            }
                        }
                    }
                }
                _ = fi50
            }
            v124 := []int{(fi49 & genFn0())}
            switch "" {
            case "a":
                switch ("a" + (("foo bar" + "a") + "")) {
                case "a":
                    v2 = 2
                case "":
                    fmt.Println(v1)
                default:
                    var v125 uint32 = (v41 << (((v2 >> (((^(v15))) & 63))) & 63))
                    _ = v125
                }
            case "hello":
                v1 -= (v1 * (v1 % ((uint8(v15)) | 1)))
            case "":
                v126 := GenS12{F0: (v3 * (v3 & 42)), F1: (fi49 << ((v2) & 63))}
                _ = v126
            }
            _ = fi49
            _ = v124
        }
    }
    var v127 int64 = (^(v3))
    if true {
        v128 := GenS11{F0: 713, F1: (v2 << (((^(v15))) & 63)), F2: 232}
        v1 += uint8(v15)
        if false {
            var v129 float64 = -2.25
            fmt.Println(v127, v1, v1)
            switch (("a" + "") + ("a" + "a")) {
            case "":
                var v130 string = "foo bar"
                _ = v130
            case "hello":
                v129 -= ((-2.25 - (v129 + v129)) * 3.14159)
            }
            _ = v129
        }
        _ = v128
    }
    v131 := v5[(("foo bar" + ("\n\t" + (("a" + "a") + "\n\t"))) + "hello")]
    fmt.Println("v1:", v1)
    fmt.Println("v127:", v127)
    fmt.Println("v131:", v131)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v127
    _ = v131
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v44
    _ = v5
}
