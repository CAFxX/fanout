package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 string }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint; F1 int32; F2 int64 }
type GenS9 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS10 struct { F0 uint16; F1 bool; F2 bool }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int }
func genFn0(p0 int32, p1 uint8) int {
    p1 = 45
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int16 = 0
        if (((-(p0)) > (-(p0))) && true) {
            for fi9 := 0; fi9 < 10; fi9++ {
                fmt.Println(v8, p0)
                _ = fi9
            }
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 uint = 1
                _ = fi10
                _ = v11
            }
        } else {
            var v12 int8 = 127
            v13 := []int{(fi7 << ((2) & 63)), 32768, fi7, (fi7 - 108)}
            _ = v12
            _ = v13
        }
        var v14 float32 = ((-2.25 / 1.5) + 3.14159)
        _ = fi7
        _ = v14
        _ = v8
    }
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float64 = (-2.25 / 1e10)
    var v3 uint = 330
    var v4 int = 128
    v5 := []string{}
    var v6 string = "foo bar"
    fmt.Println(v4)
    switch v6 {
    case "hello":
        fmt.Println(v2)
    case "\n\t":
        for fi15 := 0; fi15 < 17; fi15++ {
            v6 = v6
            v3 = v3
            for fi16 := 0; fi16 < 18; fi16++ {
                fmt.Println(v6)
                _ = fi16
            }
            _ = fi15
        }
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 19; fi18++ {
            v1 = v1
            var v19 bool = (-2147483648 < 85)
            _ = fi18
            _ = v19
        }
        fmt.Println(v4, v2, v1)
        switch fi17 {
        case 2147483647:
            v20 := GenS1{F0: (32767 < 2), F1: "", F2: ""}
            _ = v20
        default:
            var v21 int32 = 53
            _ = v21
        }
        _ = fi17
    }
    fmt.Println(v1, v1, v6)
    for fi22 := 0; fi22 < 20; fi22++ {
        switch (v4 | (fi22 + genFn0(-2147483648, 179))) {
        case -199:
            var v23 uint = (^(v3))
            _ = v23
        case -1:
            fmt.Println(v1)
        case -9223372036854775808:
            switch fi22 {
            case 4294967295:
                switch (v4 / ((fi22) | 1)) {
                case -1:
                    v24 := GenS1{F0: (-2147483648 >= -1), F1: v6, F2: ((v6 + ((v6 + (v6 + v6)) + ("\n\t" + v6))) + "a")}
                    _ = v24
                }
            case -1:
                if false {
                    if false {
                        for fi25 := 0; fi25 < 10; fi25++ {
                            v2 -= float64(1e10)
                            if false {
                                v2 += v2
                                if (-1 < 42) {
                                    if ("" == v6) {
                                        v1 = -860
                                    }
                                    switch v6 {
                                    case "hello":
                                        v4 = v4
                                    case "foo bar":
                                        if (9 > 280) {
                                            v1 += -9223372036854775808
                                            if true {
                                                v26 := GenS1{F0: ((2147483647 < 0) && (v3 > (^(v3)))), F1: (((v6 + v6) + "hello") + v6), F2: "hello"}
                                                if true {
                                                    if false {
                                                        v27 := GenS1{F0: false, F1: "a", F2: "foo bar"}
                                                        _ = v27
                                                    } else {
                                                        v28 := GenS2{F0: 283, F1: 127}
                                                        var v29 uint = 103
                                                        _ = v28
                                                        _ = v29
                                                    }
                                                    v30 := GenS2{F0: 1, F1: 52}
                                                    if false {
                                                        switch 0 {
                                                        case 4294967295:
                                                            switch v6 {
                                                            case "hello":
                                                                v1 = (v1 - v1)
                                                            case "":
                                                                v3 -= v3
                                                            case "\n\t":
                                                                fmt.Println(v3, v6)
                                                            default:
                                                                var v31 float32 = -2.25
                                                                _ = v31
                                                            }
                                                        }
                                                        v32 := GenS1{F0: (uint64(v3) > 1), F1: v6, F2: ""}
                                                        if false {
                                                            fmt.Println(v4, v6)
                                                            _ = len(v5)
                                                        }
                                                        _ = v32
                                                    }
                                                    fmt.Println(v4)
                                                    var v33 float32 = (float32((v2 / float64(0.1))) / 0.1)
                                                    _ = v30
                                                    _ = v33
                                                }
                                                _ = len(v5)
                                                _ = v26
                                            } else {
                                                var v34 int = (v4 / ((fi22) | 1))
                                                var v35 string = (v6 + v6)
                                                if (v1 >= v1) {
                                                    var v36 int64 = v1
                                                    v35 = v35
                                                    _ = v36
                                                } else {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    if (0 < uint64(v3)) {
                                                        v6 = v6
                                                        if (0 <= 1) {
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = v34
                                                _ = v35
                                            }
                                        } else {
                                            var v37 string = v6
                                            if (!false) {
                                                if (-400 <= 23) {
                                                    v38 := GenS3{F0: (v1 / ((v1) | 1)), F1: 519, F2: 1}
                                                    var v39 float64 = float64(1e10)
                                                    v39 -= (0.1 / (1.5 - (v39 - (v39 - 0.1))))
                                                    _ = v38
                                                    _ = v39
                                                } else {
                                                    switch fi25 {
                                                    case 32767:
                                                        v37 = "foo bar"
                                                    case -2147483648:
                                                        fmt.Println(fi22, v4, v37)
                                                    default:
                                                        var v40 int16 = -276
                                                        _ = v40
                                                    }
                                                    var v41 uint16 = 1
                                                    var v42 int = fi22
                                                    _ = v41
                                                    _ = v42
                                                }
                                            } else {
                                                if (false || true) {
                                                    v43 := GenS1{F0: false, F1: v37, F2: v6}
                                                    var v44 float32 = 1e10
                                                    _ = v43
                                                    _ = v44
                                                } else {
                                                    v45 := GenS4{F0: 9223372036854775807}
                                                    _ = v45
                                                }
                                                v37 = "foo bar"
                                                v2 = v2
                                            }
                                            _ = v37
                                        }
                                    default:
                                        var v46 int64 = (v1 + (v1 | (v1 << ((425) & 63))))
                                        _ = v46
                                    }
                                } else {
                                    v47 := GenS3{F0: -9223372036854775808, F1: 2, F2: 1}
                                    _ = v47
                                }
                                var v48 uint64 = 0
                                if (true && true) {
                                    if true {
                                        _ = len(v5)
                                    } else {
                                        v4 = int(v1)
                                        fmt.Println(v3, v4, v4)
                                    }
                                } else {
                                    v1 = v1
                                    var v49 int8 = -128
                                    v5 = append(v5, ("\n\t" + v6))
                                    _ = v49
                                }
                                _ = v48
                            }
                            if (18446744073709551615 >= 32) {
                                var v50 bool = true
                                _ = len(v5)
                                _ = v50
                            }
                            _ = fi25
                        }
                        if true {
                            if len(v5) > 0 {
                                v5[0] = v6
                            }
                        }
                    } else {
                        for fi51 := 0; fi51 < 20; fi51++ {
                            if true {
                                v5 = append(v5, "foo bar")
                            } else {
                                if false {
                                    var v52 string = v6
                                    fmt.Println(v52, v4)
                                    switch ((v52 + v52) + v6) {
                                    case "\n\t":
                                        var v53 uint64 = 281
                                        _ = v53
                                    case "hello":
                                        _ = len(v5)
                                    case "a":
                                        fmt.Println(fi22, fi22)
                                    default:
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                    if ((("\n\t" + v52) > "") && true) {
                                        v54 := GenS3{F0: (v1 & v1), F1: 2147483647, F2: 477}
                                        switch genFn0(320, 128) {
                                        case 1:
                                            if true {
                                                fmt.Println(v2, v3, v1)
                                            }
                                        default:
                                            if true {
                                                v2 += v2
                                                var v55 int32 = 0
                                                v56 := GenS5{F0: (fi22 - (v4 >> ((47) & 63))), F1: v3}
                                                _ = v55
                                                _ = v56
                                            }
                                        }
                                        _ = v54
                                    } else {
                                        switch v4 {
                                        case 42:
                                            _ = len(v5)
                                        case -196:
                                            var v57 string = "a"
                                            _ = v57
                                        case 9223372036854775807:
                                            v58 := GenS3{F0: (v1 ^ v1), F1: 2147483648, F2: 0}
                                            _ = v58
                                        }
                                        fmt.Println(v2, fi51)
                                        var v59 float32 = 0.1
                                        var v60 int = fi22
                                        var v61 int64 = (-(v1))
                                        switch (((("a" + v52) + (((v6 + v6) + "hello") + "a")) + ((((v6 + v52) + ("a" + v6)) + "\n\t") + v52)) + "foo bar") {
                                        case "\n\t":
                                            var v62 uint64 = 843
                                            _ = v62
                                        }
                                        _ = v59
                                        _ = v60
                                        _ = v61
                                    }
                                    var v63 float32 = -2.25
                                    _ = v52
                                    _ = v63
                                }
                                if false {
                                    fmt.Println(v6, v4)
                                    fmt.Println(v4)
                                    if false {
                                        switch "\n\t" {
                                        case "hello":
                                            var v64 bool = false
                                            _ = v64
                                        case "foo bar":
                                            var v65 float64 = v2
                                            _ = v65
                                        case "a":
                                            if (v3 != v3) {
                                                _ = len(v5)
                                                v1 += (v1 >> ((1) & 63))
                                                var v66 float32 = 0.0
                                                var v67 uint32 = 2147483647
                                                fmt.Println(v1, fi51)
                                                _ = v66
                                                _ = v67
                                            }
                                        default:
                                            if true {
                                                v2 = v2
                                                v68 := GenS4{F0: 1}
                                                var v69 int8 = 42
                                                _ = v68
                                                _ = v69
                                            } else {
                                                v1 = (v1 | (v1 ^ (^(v1))))
                                                v4 = (-(fi22))
                                                v4 = fi51
                                            }
                                        }
                                        if (149 < 32767) {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            v1 -= (v1 % ((499) | 1))
                                        } else {
                                            v70 := GenS5{F0: (v4 << ((v3) & 63)), F1: v3}
                                            var v71 uint16 = 1
                                            _ = v70
                                            _ = v71
                                        }
                                    }
                                }
                                if len(v5) > 0 {
                                    v5[0] = v6
                                }
                            }
                            var v72 float64 = (v2 / float64(1.5))
                            if (false && true) {
                                fmt.Println(v72)
                            } else {
                                v2 += v2
                                switch 32768 {
                                case -9223372036854775808:
                                    v3 += (v3 % ((65535) | 1))
                                case 32767:
                                    fmt.Println(v4, v72, v3)
                                }
                                if (!true) {
                                    if (127 == 255) {
                                        v73 := GenS3{F0: v1, F1: 42, F2: 42}
                                        _ = v73
                                    }
                                    var v74 uint64 = 4294967295
                                    fmt.Println(v3, v6, v6)
                                    _ = v74
                                }
                            }
                            _ = fi51
                            _ = v72
                        }
                        if true {
                            v75 := GenS5{F0: genFn0(-42, 76), F1: (^(v3))}
                            switch (fi22 | v4) {
                            case 331:
                                _ = len(v5)
                            case 1:
                                v1 -= (v1 * v1)
                            case -1:
                                v76 := GenS6{F0: v6, F1: 8}
                                _ = v76
                            }
                            _ = v75
                        } else {
                            if false {
                                if false {
                                    v3 += 0
                                } else {
                                    if (!(true && true)) {
                                        switch 32767 {
                                        case 9223372036854775807:
                                            switch 32768 {
                                            case 32768:
                                                for fi77 := 0; fi77 < 3; fi77++ {
                                                    v78 := GenS1{F0: (true && true), F1: (((("a" + (v6 + v6)) + (v6 + "a")) + v6) + (("hello" + "") + "hello")), F2: v6}
                                                    v5 = append(v5, v6)
                                                    v1 -= -9223372036854775808
                                                    v79 := GenS3{F0: (v1 * (v1 << ((v3) & 63))), F1: 0, F2: 2147483648}
                                                    if true {
                                                        var v80 float64 = (3.14159 * 3.14159)
                                                        _ = v80
                                                    }
                                                    fmt.Println(v2)
                                                    _ = fi77
                                                    _ = v78
                                                    _ = v79
                                                }
                                            case -9223372036854775808:
                                                v4 += -2147483648
                                            case 288:
                                                switch v6 {
                                                case "foo bar":
                                                    if len(v5) > 0 {
                                                        v5[0] = v6
                                                    }
                                                case "\n\t":
                                                    switch v4 {
                                                    case 795:
                                                        var v81 uint16 = 255
                                                        _ = v81
                                                    case 344:
                                                        switch v6 {
                                                        case "hello":
                                                            if len(v5) > 0 {
                                                                v5[0] = v6
                                                            }
                                                        case "a":
                                                            _ = len(v5)
                                                        case "":
                                                            if (false || (int8(v3) != 41)) {
                                                                _ = len(v5)
                                                                v82 := GenS7{F0: 2}
                                                                _ = v82
                                                            } else {
                                                                v5 = append(v5, "\n\t")
                                                                v1 = v1
                                                                if (!false) {
                                                                    _ = len(v5)
                                                                }
                                                            }
                                                        }
                                                    case 126:
                                                        v3 = v3
                                                    }
                                                case "hello":
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                default:
                                                    var v83 uint64 = 41
                                                    _ = v83
                                                }
                                            }
                                        case 1:
                                            switch genFn0(2147483647, 128) {
                                            case 4294967295:
                                                v4 += 1
                                            case 2147483647:
                                                v1 -= v1
                                            default:
                                                for fi84 := 0; fi84 < 19; fi84++ {
                                                    v85 := GenS6{F0: (v6 + v6), F1: int8(fi22)}
                                                    if len(v5) > 0 {
                                                        v5[0] = (v6 + ((("\n\t" + v6) + v6) + ""))
                                                    }
                                                    if true {
                                                        v86 := GenS3{F0: (-(v1)), F1: 510, F2: 1}
                                                        v2 += v2
                                                        _ = v86
                                                    } else {
                                                        var v87 uint8 = 222
                                                        fmt.Println(v4)
                                                        _ = v87
                                                    }
                                                    fmt.Println(v4)
                                                    if true {
                                                        v88 := GenS7{F0: 2147483648}
                                                        _ = v88
                                                    } else {
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        var v89 float32 = 1.5
                                                        _ = v89
                                                    }
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        switch (fi84 % ((fi84) | 1)) {
                                                        case -9223372036854775808:
                                                            if false {
                                                                v1 -= 761
                                                                v5 = append(v5, "")
                                                                fmt.Println(fi84, fi22)
                                                                v90 := GenS8{F0: (v3 ^ (v3 << ((783) & 63))), F1: 1, F2: 482}
                                                                v91 := GenS3{F0: (v1 ^ 1), F1: 2, F2: 4294967295}
                                                                _ = v90
                                                                _ = v91
                                                            }
                                                        case 1:
                                                            fmt.Println(fi22, v3)
                                                        }
                                                    case "a":
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    case "hello":
                                                        switch "\n\t" {
                                                        case "a":
                                                            var v92 uint16 = 749
                                                            _ = v92
                                                        default:
                                                            _ = len(v5)
                                                        }
                                                    default:
                                                        v2 += 3.14159
                                                    }
                                                    _ = fi84
                                                    _ = v85
                                                }
                                            }
                                        default:
                                            for fi93 := 0; fi93 < 20; fi93++ {
                                                v5 = append(v5, v6)
                                                _ = fi93
                                            }
                                        }
                                    } else {
                                        v94 := GenS9{F0: 127, F1: 42, F2: 18446744073709551615}
                                        _ = v94
                                    }
                                }
                                if ((971 > uint32(v3)) && true) {
                                    var v95 uint = (v3 + 65535)
                                    var v96 uint = v95
                                    v96 -= v3
                                    _ = v95
                                    _ = v96
                                } else {
                                    v2 = ((v2 + (v2 / 1.0)) * 0.1)
                                }
                                var v97 bool = false
                                _ = v97
                            } else {
                                for fi98 := 0; fi98 < 18; fi98++ {
                                    var v99 int = fi22
                                    v2 += v2
                                    _ = fi98
                                    _ = v99
                                }
                                _ = len(v5)
                                var v100 uint16 = 65535
                                _ = v100
                            }
                            v6 = v6
                            _ = len(v5)
                        }
                        v101 := GenS1{F0: (!(true && ((1 < 42) || false))), F1: "", F2: (("\n\t" + "foo bar") + "a")}
                        _ = v101
                    }
                    var v102 uint8 = 127
                    if false {
                        switch "\n\t" {
                        case "a":
                            if len(v5) > 0 {
                                v5[0] = ""
                            }
                        default:
                            var v103 uint = (^(v3))
                            _ = v103
                        }
                        switch v6 {
                        case "":
                            var v104 uint8 = (v102 ^ (v102 ^ (v102 + (v102 | (^(v102))))))
                            _ = v104
                        }
                    } else {
                        v6 = (v6 + v6)
                    }
                    _ = v102
                }
            default:
                if false {
                    var v105 bool = false
                    v106 := GenS2{F0: 578, F1: 109}
                    v4 -= int(v3)
                    _ = v105
                    _ = v106
                } else {
                    var v107 float32 = ((0.1 - 0.0) - (-2.25 - 3.14159))
                    var v108 int = fi22
                    _ = v107
                    _ = v108
                }
            }
        }
        fmt.Println(v3, v3)
        var v109 int64 = (^(v1))
        _ = fi22
        _ = v109
    }
    switch v4 {
    case 127:
        var v110 int16 = 1
        _ = v110
    }
    v6 = (v6 + "foo bar")
    v3 = 27
    v1 -= v1
    switch (v4 * (^(v4))) {
    case -752:
        var v111 string = v6
        _ = v111
    case -734:
        v112 := GenS9{F0: int8(v4), F1: 359, F2: 1}
        _ = v112
    default:
        _ = len(v5)
    }
    for fi113 := 0; fi113 < 4; fi113++ {
        v1 += v1
        for fi114 := 0; fi114 < 13; fi114++ {
            v115 := GenS2{F0: 124, F1: 42}
            _ = fi114
            _ = v115
        }
        _ = fi113
    }
    fmt.Println(v2)
    if false {
        var v116 uint32 = 2
        v117 := GenS8{F0: (^(v3)), F1: 1, F2: 0}
        _ = v116
        _ = v117
    } else {
        var v118 uint16 = uint16(v3)
        if false {
            for fi119 := 0; fi119 < 5; fi119++ {
                if (!false) {
                    v120 := GenS10{F0: (v118 - (v118 / ((65535) | 1))), F1: ((v3 % (((v3 >> ((2) & 63))) | 1)) != (v3 - 42)), F2: (v2 < 0.1)}
                    _ = v120
                } else {
                    var v121 int32 = -2147483648
                    switch 9223372036854775807 {
                    case 32768:
                        var v122 bool = false
                        _ = v122
                    case -763:
                        var v123 int = fi119
                        _ = v123
                    case 4294967295:
                        fmt.Println(v2, v118, v6)
                    }
                    _ = v121
                }
                v124 := GenS8{F0: (v3 >> ((v3) & 63)), F1: -53, F2: 524}
                _ = fi119
                _ = v124
            }
            var v125 uint16 = 32767
            if (v1 >= v1) {
                if (813 >= 2147483647) {
                    if (v6 < "\n\t") {
                        v6 = "hello"
                        for fi126 := 0; fi126 < 19; fi126++ {
                            v127 := GenS9{F0: -128, F1: v125, F2: 4294967295}
                            if true {
                                v4 -= v4
                                v5 = append(v5, (v6 + v6))
                            }
                            v128 := GenS11{F0: -54}
                            _ = fi126
                            _ = v127
                            _ = v128
                        }
                        for fi129 := 0; fi129 < 14; fi129++ {
                            var v130 float64 = 1.5
                            v4 = (v4 & (v4 >> (((v3 << ((uint(v125)) & 63))) & 63)))
                            _ = fi129
                            _ = v130
                        }
                    } else {
                        for fi131 := 0; fi131 < 5; fi131++ {
                            v3 = 2
                            v5 = append(v5, v6)
                            var v132 int32 = -1
                            _ = fi131
                            _ = v132
                        }
                    }
                    if (1 < 785) {
                        v5 = append(v5, v6)
                    } else {
                        fmt.Println(v4, v1)
                        v133 := GenS7{F0: 42}
                        v134 := GenS12{F0: -1}
                        _ = v133
                        _ = v134
                    }
                    for fi135 := 0; fi135 < 19; fi135++ {
                        if (1e10 == -2.25) {
                            if false {
                                _ = len(v5)
                                v2 -= v2
                                if false {
                                    fmt.Println(v125)
                                    for fi136 := 0; fi136 < 12; fi136++ {
                                        if false {
                                            var v137 string = v6
                                            var v138 uint16 = 65535
                                            switch fi135 {
                                            case 1:
                                                if true {
                                                    fmt.Println(v137, v1)
                                                    if len(v5) > 0 {
                                                        v5[0] = v6
                                                    }
                                                    v5 = append(v5, "")
                                                }
                                            case 32767:
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            case 371:
                                                fmt.Println(fi135, v137)
                                            default:
                                                v5 = append(v5, v137)
                                            }
                                            _ = v137
                                            _ = v138
                                        }
                                        if (!false) {
                                            v139 := GenS10{F0: v125, F1: false, F2: ((^(v125)) != 32767)}
                                            v140 := GenS6{F0: v6, F1: 0}
                                            var v141 int16 = -32768
                                            _ = v139
                                            _ = v140
                                            _ = v141
                                        }
                                        _ = fi136
                                    }
                                    _ = len(v5)
                                }
                            }
                            fmt.Println(v1)
                            v3 = v3
                        }
                        switch "\n\t" {
                        case "":
                            if true {
                                _ = len(v5)
                                for fi142 := 0; fi142 < 7; fi142++ {
                                    v143 := GenS9{F0: 42, F1: v118, F2: 42}
                                    fmt.Println(v1)
                                    var v144 int = (fi135 % (((-(v4))) | 1))
                                    var v145 uint8 = uint8(v125)
                                    var v146 uint64 = 1
                                    _ = fi142
                                    _ = v143
                                    _ = v144
                                    _ = v145
                                    _ = v146
                                }
                            } else {
                                v118 -= (v118 * 1)
                                for fi147 := 0; fi147 < 18; fi147++ {
                                    var v148 int16 = 1000
                                    _ = fi147
                                    _ = v148
                                }
                            }
                        case "a":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case "\n\t":
                            v6 = v6
                        default:
                            if ((v4 / (((fi135 >> ((v3) & 63))) | 1)) == (v4 ^ (-(v4)))) {
                                var v149 uint16 = 32767
                                if false {
                                    var v150 int8 = 79
                                    if false {
                                        v125 = v125
                                    }
                                    if true {
                                        if (0.0 > 0.1) {
                                            if true {
                                                var v151 uint64 = 991
                                                var v152 int8 = v150
                                                _ = v151
                                                _ = v152
                                            }
                                        }
                                        var v153 uint = (v3 + v3)
                                        _ = v153
                                    }
                                    _ = v150
                                } else {
                                    v6 = "foo bar"
                                }
                                _ = v149
                            } else {
                                if true {
                                    v6 = ((v6 + (v6 + "\n\t")) + ((v6 + v6) + v6))
                                    var v154 uint8 = 127
                                    for fi155 := 0; fi155 < 11; fi155++ {
                                        switch -607 {
                                        case -803:
                                            if ((0.0 / (v2 - 3.14159)) != v2) {
                                                v1 -= 1
                                            } else {
                                                _ = len(v5)
                                            }
                                        default:
                                            var v156 int64 = (v1 / ((v1) | 1))
                                            _ = v156
                                        }
                                        if ((-9223372036854775808 != v1) && ("a" > v6)) {
                                            v154 = 13
                                            v3 = uint(v118)
                                        } else {
                                            v6 = "\n\t"
                                            v157 := GenS13{F0: genFn0(2147483647, v154)}
                                            _ = v157
                                        }
                                        _ = fi155
                                    }
                                    _ = v154
                                }
                            }
                        }
                        v2 -= v2
                        _ = fi135
                    }
                } else {
                    for fi158 := 0; fi158 < 14; fi158++ {
                        v159 := GenS10{F0: (v118 >> ((255) & 63)), F1: (-1 != 0), F2: true}
                        _ = fi158
                        _ = v159
                    }
                    for fi160 := 0; fi160 < 15; fi160++ {
                        v4 -= int(v1)
                        _ = fi160
                    }
                }
            }
            _ = v125
        }
        var v161 string = v6
        _ = v118
        _ = v161
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
