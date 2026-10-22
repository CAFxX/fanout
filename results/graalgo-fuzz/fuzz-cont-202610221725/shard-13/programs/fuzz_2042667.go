package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int64; F1 int32; F2 string }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int16; F1 bool; F2 int32 }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 uint8; F1 int8 }
type GenS8 struct { F0 uint64; F1 int32 }
type GenS9 struct { F0 int; F1 int64 }
type GenS10 struct { F0 int16; F1 int8; F2 int32 }
type GenS11 struct { F0 string; F1 uint16 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int64; F1 uint }
type GenS14 struct { F0 int16; F1 uint32 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 uint16; F1 uint64 }
type GenS17 struct { F0 string }
func genFn0() int16 {
    var v6 uint32 = 65535
    v6 -= v6
    var v7 int = -1
    var v8 bool = false
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 bool = (fi9 != v7)
        if ((0.0 + 1.5) >= 1e10) {
            v11 := []int{v7}
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := GenS1{F0: 127, F1: 1}
                v8 = (v10 || (-1 < -32768))
                if v8 {
                    switch fi12 {
                    case 42:
                        v6 = (v6 >> ((681) & 63))
                        v6 -= v6
                    case -2147483648:
                        fmt.Println(fi9, v10)
                    }
                    if (!v8) {
                        v8 = (2147483647 >= 584)
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        v14 := GenS1{F0: 127, F1: 1}
                        _ = v14
                    } else {
                        v10 = (!(-198 < 2147483647))
                        if v8 {
                            fmt.Println(v6, v6)
                        }
                        v15 := GenS1{F0: -1, F1: 127}
                        _ = v15
                    }
                    v16 := GenS1{F0: -128, F1: 81}
                    if v10 {
                        var v17 float64 = (3.14159 - -2.25)
                        v8 = (2 > 2)
                        _ = v17
                    }
                    v18 := GenS1{F0: -128, F1: 127}
                    var v19 uint32 = (^(v6))
                    _ = v16
                    _ = v18
                    _ = v19
                } else {
                    if v8 {
                        switch ("\n\t" + "") {
                        case "foo bar":
                            var v20 int8 = 1
                            _ = v20
                        case "":
                            if true {
                                v21 := GenS1{F0: 63, F1: 100}
                                if v8 {
                                    var v22 int64 = 9223372036854775807
                                    _ = v22
                                } else {
                                    fmt.Println(v8, v10, fi12)
                                    var v23 string = "a"
                                    _ = v23
                                }
                                v10 = true
                                _ = v21
                            }
                            v10 = v8
                            v24 := GenS2{F0: -1, F1: 2147483647, F2: "foo bar"}
                            var v25 int32 = 960
                            v26 := GenS3{F0: 482, F1: v6}
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        v6 = 602
                    }
                }
                v6 = 229
                fmt.Println(v7, v8)
                _ = fi12
                _ = v13
            }
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1() uint16 {
    if false {
        if false {
            for fi27 := 0; fi27 < 17; fi27++ {
                var v28 uint16 = 2
                fmt.Println(fi27, fi27)
                v28 -= (^(v28))
                _ = fi27
                _ = v28
            }
        } else {
            if false {
                var v29 int32 = 0
                v30 := []int{42, 535, -131, 0}
                _ = v29
                _ = v30
            } else {
                if (false && (((("" + "hello") + "\n\t") + "a") <= "hello")) {
                    if true {
                        var v31 uint32 = 2147483647
                        var v32 uint32 = 395
                        _ = v31
                        _ = v32
                    } else {
                        v33 := []int{32767}
                        _ = v33
                    }
                    v34 := []int{0, -269}
                    _ = v34
                }
                var v35 int8 = 0
                _ = v35
            }
            var v36 uint = 255
            _ = v36
        }
    }
    v37 := GenS3{F0: 2, F1: 0}
    _ = v37
    return 2
}

func main() {
    var v1 float32 = 0.1
    v2 := []uint32{0}
    var v3 int8 = -14
    var v4 bool = false
    var v5 int32 = 42
    if v4 {
        switch 1 {
        case 1:
            v38 := GenS4{F0: 233, F1: v5}
            _ = v38
        case 128:
            for fi39 := 0; fi39 < 1; fi39++ {
                v4 = v4
                _ = len(v2)
                var v40 uint = 1
                _ = fi39
                _ = v40
            }
        case -2147483648:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    switch -9223372036854775808 {
    case 4294967295:
        switch ("hello" + "") {
        case "a":
            v41 := GenS2{F0: -916, F1: v5, F2: "hello"}
            _ = v41
        case "":
            if v4 {
                switch "\n\t" {
                case "":
                    var v42 bool = v4
                    _ = v42
                case "foo bar":
                    v43 := GenS5{F0: -908, F1: (false && false), F2: (v5 + -1)}
                    _ = v43
                case "a":
                    for fi44 := 0; fi44 < 4; fi44++ {
                        for fi45 := 0; fi45 < 12; fi45++ {
                            v46 := GenS4{F0: 109, F1: (v5 / ((179) | 1))}
                            _ = fi45
                            _ = v46
                        }
                        v4 = (42 <= 1)
                        if (v4 && v4) {
                            switch (fi44 | (fi44 | 331)) {
                            case 0:
                                for fi47 := 0; fi47 < 10; fi47++ {
                                    var v48 uint = 255
                                    if (318 < uint32(v48)) {
                                        if (fi47 == fi44) {
                                            v49 := GenS2{F0: -9223372036854775808, F1: (v5 >> (((v48 + (v48 << ((0) & 63)))) & 63)), F2: ""}
                                            if ((v5 + 1) <= (v5 + (v5 >> ((v48) & 63)))) {
                                                v50 := GenS6{F0: "\n\t", F1: (v48 % ((v48) | 1))}
                                                if true {
                                                    if (genFn0() > -1) {
                                                        v4 = (65535 > 984)
                                                        _ = len(v2)
                                                    } else {
                                                        if len(v2) > 0 {
                                                            v2[0] = 956
                                                        }
                                                        v5 = v5
                                                        var v51 uint64 = 429
                                                        _ = v51
                                                    }
                                                    _ = len(v2)
                                                    v2 = append(v2, 1)
                                                }
                                                if (0 < 255) {
                                                    switch "" {
                                                    case "foo bar":
                                                        if (-1 >= 42) {
                                                            v48 = 255
                                                            var v52 int64 = 0
                                                            var v53 int16 = -226
                                                            if true {
                                                                v53 += (v53 + (v53 | 42))
                                                                if ((-(v3)) >= v3) {
                                                                    v52 = v52
                                                                } else {
                                                                    var v54 int8 = (v3 | (v3 | v3))
                                                                    v48 = 4294967295
                                                                    _ = v54
                                                                }
                                                                var v55 uint = 618
                                                                _ = v55
                                                            } else {
                                                                v5 = v5
                                                            }
                                                            v52 += (-(v52))
                                                            _ = v52
                                                            _ = v53
                                                        } else {
                                                            var v56 uint16 = 275
                                                            _ = len(v2)
                                                            if v4 {
                                                                var v57 float32 = 3.14159
                                                                fmt.Println(v4)
                                                                if ((-2.25 * 1e10) == 1.5) {
                                                                    v58 := GenS4{F0: 128, F1: (v5 & -1)}
                                                                    _ = v58
                                                                } else {
                                                                    v1 = ((v57 - v1) / v1)
                                                                    v59 := GenS7{F0: 128, F1: -1}
                                                                    v4 = true
                                                                    _ = v59
                                                                }
                                                                _ = v57
                                                            }
                                                            _ = v56
                                                        }
                                                    default:
                                                        if (42 <= 2) {
                                                            var v60 int8 = (v3 / ((1) | 1))
                                                            v48 = v48
                                                            var v61 uint = 4294967295
                                                            fmt.Println(fi44, v3, fi44)
                                                            v62 := GenS5{F0: -32768, F1: true, F2: v5}
                                                            _ = v60
                                                            _ = v61
                                                            _ = v62
                                                        } else {
                                                            if len(v2) > 0 {
                                                                v2[0] = 281
                                                            }
                                                            var v63 string = "\n\t"
                                                            var v64 uint = v48
                                                            _ = v63
                                                            _ = v64
                                                        }
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                fmt.Println(v4)
                                                _ = v50
                                            } else {
                                                if (70 <= int16(v5)) {
                                                    _ = len(v2)
                                                    v3 = 100
                                                }
                                                var v65 int32 = (v5 / (((v5 % ((v5) | 1))) | 1))
                                                _ = v65
                                            }
                                            _ = v49
                                        } else {
                                            _ = len(v2)
                                            if v4 {
                                                if (0.1 >= float64(0.1)) {
                                                    if false {
                                                        v1 += float32(0.0)
                                                        var v66 uint64 = 564
                                                        if len(v2) > 0 {
                                                            v2[0] = 187
                                                        }
                                                        _ = v66
                                                    } else {
                                                        v67 := GenS8{F0: 655, F1: v5}
                                                        _ = v67
                                                    }
                                                } else {
                                                    var v68 float64 = float64(v1)
                                                    if (!(!v4)) {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    } else {
                                                        switch "" {
                                                        case "\n\t":
                                                            var v69 uint16 = 42
                                                            _ = v69
                                                        default:
                                                            v48 = 13
                                                        }
                                                    }
                                                    v70 := GenS7{F0: 1, F1: (v3 - 110)}
                                                    _ = v68
                                                    _ = v70
                                                }
                                                var v71 int = fi47
                                                v4 = ((1.5 + (((v1 * v1) / 3.14159) / 1.0)) == 1.5)
                                                _ = v71
                                            }
                                        }
                                    } else {
                                        v3 = v3
                                        v5 += v5
                                        v4 = v4
                                        v5 -= (-(v5))
                                    }
                                    var v72 int64 = 9223372036854775807
                                    switch "hello" {
                                    case "hello":
                                        switch "" {
                                        case "hello":
                                            switch "hello" {
                                            case "a":
                                                if v4 {
                                                    switch fi47 {
                                                    case 0:
                                                        var v73 bool = ((("hello" + "") + "hello") > "\n\t")
                                                        _ = v73
                                                    }
                                                    switch fi44 {
                                                    case 362:
                                                        var v74 string = ("\n\t" + "foo bar")
                                                        _ = v74
                                                    case -950:
                                                        if false {
                                                            fmt.Println(v3, v72)
                                                            var v75 uint16 = 841
                                                            _ = v75
                                                        }
                                                    default:
                                                        v3 = v3
                                                    }
                                                    var v76 int = (fi44 + 2147483647)
                                                    if v4 {
                                                        v72 = v72
                                                    } else {
                                                        switch ("hello" + ("hello" + "a")) {
                                                        case "\n\t":
                                                            var v77 bool = false
                                                            var v78 uint64 = 42
                                                            switch (fi47 - fi47) {
                                                            case 22:
                                                                if (v77 && v77) {
                                                                    v78 = (v78 >> (((v48 - (v48 >> (((v48 - 18446744073709551615)) & 63)))) & 63))
                                                                    var v79 uint8 = 127
                                                                    if v4 {
                                                                        v72 -= (v72 << ((65535) & 63))
                                                                        var v80 uint = (v48 | (^(v48)))
                                                                        if v77 {
                                                                            v79 = (v79 + v79)
                                                                        } else {
                                                                            if true {
                                                                                v76 -= (fi47 + 1)
                                                                                fmt.Println(v78, v3, v4)
                                                                                switch "foo bar" {
                                                                                case "a":
                                                                                    v72 -= -1
                                                                                case "":
                                                                                    if ("" <= "\n\t") {
                                                                                        v81 := GenS6{F0: "a", F1: v48}
                                                                                        var v82 uint64 = (v78 ^ (v78 - v78))
                                                                                        _ = v81
                                                                                        _ = v82
                                                                                    }
                                                                                case "hello":
                                                                                    var v83 int8 = (v3 - int8(v79))
                                                                                    _ = v83
                                                                                }
                                                                                if ((v79 / ((v79) | 1)) < (v79 >> (((v48 / ((v80) | 1))) & 63))) {
                                                                                    var v84 uint64 = v78
                                                                                    if (v3 > (v3 / ((42) | 1))) {
                                                                                        if v77 {
                                                                                            if v77 {
                                                                                                var v85 float32 = v1
                                                                                                fmt.Println(v80, v3, v84)
                                                                                                v86 := GenS9{F0: -39, F1: -1}
                                                                                                _ = v85
                                                                                                _ = v86
                                                                                            }
                                                                                            var v87 int = fi44
                                                                                            _ = v87
                                                                                        }
                                                                                    }
                                                                                    _ = v84
                                                                                }
                                                                                var v88 uint16 = 32767
                                                                                _ = v88
                                                                            } else {
                                                                                switch 42 {
                                                                                case -102:
                                                                                    v89 := GenS5{F0: -918, F1: (v5 < v5), F2: (v5 + (-(v5)))}
                                                                                    _ = v89
                                                                                case 32768:
                                                                                    v90 := GenS7{F0: uint8(v80), F1: -44}
                                                                                    _ = v90
                                                                                case 630:
                                                                                    var v91 uint = (v80 / ((v48) | 1))
                                                                                    _ = v91
                                                                                default:
                                                                                    fmt.Println(v77, v76)
                                                                                }
                                                                            }
                                                                            fmt.Println(v48)
                                                                        }
                                                                        _ = v80
                                                                    }
                                                                    _ = v79
                                                                }
                                                            case 4294967295:
                                                                if true {
                                                                    if (v78 <= v78) {
                                                                        var v92 string = "a"
                                                                        _ = v92
                                                                    }
                                                                    v93 := GenS9{F0: 481, F1: v72}
                                                                    _ = v93
                                                                }
                                                            default:
                                                                var v94 uint = (v48 % ((v48) | 1))
                                                                _ = v94
                                                            }
                                                            _ = v77
                                                            _ = v78
                                                        case "":
                                                            v48 += v48
                                                        case "foo bar":
                                                            var v95 int16 = -1
                                                            _ = v95
                                                        default:
                                                            var v96 float32 = ((v1 + v1) * v1)
                                                            _ = v96
                                                        }
                                                    }
                                                    var v97 uint64 = 4294967295
                                                    var v98 float32 = v1
                                                    if true {
                                                        var v99 string = "foo bar"
                                                        _ = v99
                                                    }
                                                    var v100 int8 = (v3 >> ((v48) & 63))
                                                    var v101 bool = v4
                                                    _ = v100
                                                    _ = v101
                                                    _ = v76
                                                    _ = v97
                                                    _ = v98
                                                } else {
                                                    var v102 int32 = v5
                                                    _ = v102
                                                }
                                            case "foo bar":
                                                fmt.Println(v5)
                                                switch (fi44 & (fi47 & (fi44 / (((fi44 / (((^(fi44))) | 1))) | 1)))) {
                                                case -9223372036854775808:
                                                    if v4 {
                                                        if v4 {
                                                            v2 = append(v2, 1)
                                                        } else {
                                                            v103 := GenS4{F0: 24, F1: (v5 | v5)}
                                                            v4 = (v4 || true)
                                                            _ = v103
                                                        }
                                                        switch (("\n\t" + ("\n\t" + "a")) + "") {
                                                        case "\n\t":
                                                            v104 := GenS10{F0: 1000, F1: (-(v3)), F2: int32(fi47)}
                                                            _ = v104
                                                        case "":
                                                            v105 := GenS5{F0: 274, F1: (966 == 42), F2: v5}
                                                            _ = v105
                                                        }
                                                    } else {
                                                        v1 = 1e10
                                                        v106 := GenS9{F0: (fi44 ^ (-(fi44))), F1: (v72 | (v72 ^ v72))}
                                                        _ = v106
                                                    }
                                                    v107 := GenS5{F0: 41, F1: (!false), F2: int32(v72)}
                                                    v108 := GenS11{F0: ("" + ("\n\t" + "foo bar")), F1: 2}
                                                    _ = v107
                                                    _ = v108
                                                case 2147483647:
                                                    if (v4 && false) {
                                                        switch fi44 {
                                                        case 32767:
                                                            if v4 {
                                                                v48 += (v48 * 353)
                                                                fmt.Println(v1, v4)
                                                                if (222 < (v5 << ((v48) & 63))) {
                                                                    fmt.Println(v3, fi44)
                                                                    v109 := GenS12{F0: 18446744073709551615}
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        if (!(2 != 255)) {
                                                                            fmt.Println(v4)
                                                                            if (!v4) {
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                _ = len(v2)
                                                                            } else {
                                                                                v3 = v3
                                                                                var v110 uint8 = 1
                                                                                _ = len(v2)
                                                                                _ = v110
                                                                            }
                                                                            fmt.Println(v72)
                                                                        }
                                                                        if (!("foo bar" != "")) {
                                                                            switch -9223372036854775808 {
                                                                            case -2147483648:
                                                                                var v111 int32 = v5
                                                                                _ = v111
                                                                            }
                                                                        } else {
                                                                            v5 -= 2147483647
                                                                            var v112 float64 = 1e10
                                                                            _ = v112
                                                                        }
                                                                        fmt.Println(v1)
                                                                        if (1000 >= 588) {
                                                                            v113 := GenS13{F0: v72, F1: 389}
                                                                            _ = v113
                                                                        } else {
                                                                            switch "hello" {
                                                                            case "":
                                                                                fmt.Println(v5)
                                                                            case "a":
                                                                                fmt.Println(v72, v1)
                                                                                fmt.Println(v48, v1, v4)
                                                                                var v114 float64 = 0.0
                                                                                _ = v114
                                                                            default:
                                                                                var v115 string = ("\n\t" + ("a" + ("\n\t" + "")))
                                                                                _ = v115
                                                                            }
                                                                            var v116 int8 = v3
                                                                            v117 := GenS10{F0: 0, F1: (v116 ^ v116), F2: v5}
                                                                            _ = v116
                                                                            _ = v117
                                                                        }
                                                                        fmt.Println(v48, v3, fi44)
                                                                    }
                                                                    var v118 float32 = (0.0 * (((v1 - (v1 / v1)) / -2.25) * (((v1 * float32(0.0)) * (((((float32(((0.0 * 1.5) / 0.1)) / float32(0.1)) - v1) + float32(1.5)) + float32(((0.0 * 1e10) / float64((3.14159 - v1))))) + v1)) + float32(0.0))))
                                                                    if (v48 < (v48 & v48)) {
                                                                        v1 = v118
                                                                    }
                                                                    var v119 uint16 = 42
                                                                    _ = v109
                                                                    _ = v118
                                                                    _ = v119
                                                                }
                                                                if ((float32(1e10) <= v1) || (586 == 19)) {
                                                                    v1 += v1
                                                                    v3 -= v3
                                                                }
                                                            }
                                                        case 2147483647:
                                                            switch (fi47 + fi47) {
                                                            case -1:
                                                                _ = len(v2)
                                                            case -9223372036854775808:
                                                                v3 -= v3
                                                            case 0:
                                                                fmt.Println(v1, v48, v3)
                                                            default:
                                                                _ = len(v2)
                                                            }
                                                        case -2147483648:
                                                            var v120 uint64 = 336
                                                            var v121 int = 32768
                                                            v122 := GenS4{F0: 126, F1: v5}
                                                            _ = v120
                                                            _ = v121
                                                            _ = v122
                                                        default:
                                                            switch "foo bar" {
                                                            case "":
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    v123 := GenS13{F0: (-(v72)), F1: v48}
                                                                    v124 := GenS13{F0: 1, F1: v48}
                                                                    _ = v123
                                                                    _ = v124
                                                                case "hello":
                                                                    v72 = 42
                                                                }
                                                            case "foo bar":
                                                                if v4 {
                                                                    v125 := GenS11{F0: ("foo bar" + "\n\t"), F1: 32767}
                                                                    v48 -= (v48 * (^(v48)))
                                                                    _ = v125
                                                                }
                                                            case "hello":
                                                                if ((0 < 62) || false) {
                                                                    v3 = -128
                                                                    v126 := GenS7{F0: 1, F1: v3}
                                                                    var v127 int64 = v72
                                                                    v128 := GenS4{F0: 255, F1: 2147483647}
                                                                    _ = v126
                                                                    _ = v127
                                                                    _ = v128
                                                                }
                                                            }
                                                        }
                                                    }
                                                case 657:
                                                    v48 -= (v48 >> (((v48 ^ v48)) & 63))
                                                    v4 = v4
                                                    if (-565 > 1) {
                                                        fmt.Println(v5, v5, v4)
                                                    } else {
                                                        switch -1 {
                                                        case 32767:
                                                            var v129 uint = (v48 % ((209) | 1))
                                                            _ = v129
                                                        case -915:
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        default:
                                                            v130 := GenS8{F0: 365, F1: v5}
                                                            _ = v130
                                                        }
                                                        v4 = v4
                                                        v131 := GenS6{F0: "foo bar", F1: v48}
                                                        _ = v131
                                                    }
                                                    var v132 bool = (v4 || (("" + "hello") == "a"))
                                                    var v133 int = (fi47 / (((fi47 + -245)) | 1))
                                                    v3 -= (v3 + v3)
                                                    if (!v4) {
                                                        switch ("hello" + (("\n\t" + "a") + "a")) {
                                                        case "\n\t":
                                                            var v134 float32 = 3.14159
                                                            _ = v134
                                                        case "hello":
                                                            switch "" {
                                                            case "a":
                                                                v135 := GenS10{F0: genFn0(), F1: v3, F2: (v5 & v5)}
                                                                _ = v135
                                                            case "\n\t":
                                                                fmt.Println(v48, v48, v4)
                                                            default:
                                                                v133 = (fi47 >> (((v48 << ((65535) & 63))) & 63))
                                                                v3 = (v3 / (((v3 - v3)) | 1))
                                                                switch "a" {
                                                                case "":
                                                                    v3 = (v3 ^ v3)
                                                                }
                                                            }
                                                        }
                                                    }
                                                    var v136 string = ""
                                                    _ = v132
                                                    _ = v133
                                                    _ = v136
                                                default:
                                                    if v4 {
                                                        var v137 uint16 = 0
                                                        _ = v137
                                                    } else {
                                                        fmt.Println(v1, v4, v1)
                                                        switch fi44 {
                                                        case 2147483647:
                                                            var v138 uint8 = 127
                                                            _ = v138
                                                        case 127:
                                                            v48 = v48
                                                        case -9223372036854775808:
                                                            if v4 {
                                                                var v139 uint32 = 2147483647
                                                                _ = v139
                                                            } else {
                                                                v3 = v3
                                                                fmt.Println(v4)
                                                            }
                                                        }
                                                        v140 := GenS10{F0: -512, F1: v3, F2: v5}
                                                        _ = v140
                                                    }
                                                }
                                            case "hello":
                                                var v141 bool = false
                                                _ = v141
                                            }
                                        case "a":
                                            fmt.Println(v1, v3)
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v48 = v48
                                            v48 = 2
                                            v1 += 0.1
                                            v2 = append(v2, 65535)
                                        case "\n\t":
                                            fmt.Println(v72, fi44)
                                        default:
                                            v1 = 0.1
                                        }
                                    case "a":
                                        switch "\n\t" {
                                        case "\n\t":
                                            v142 := GenS10{F0: -32768, F1: (v3 & 100), F2: 42}
                                            _ = v142
                                        case "":
                                            var v143 bool = true
                                            _ = v143
                                        case "hello":
                                            var v144 bool = false
                                            _ = v144
                                        default:
                                            fmt.Println(v1, v48)
                                        }
                                        if v4 {
                                            v48 += v48
                                        } else {
                                            var v145 float64 = -2.25
                                            v146 := GenS12{F0: 42}
                                            _ = v145
                                            _ = v146
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case "foo bar":
                                        switch "a" {
                                        case "foo bar":
                                            switch "\n\t" {
                                            case "hello":
                                                switch ("hello" + "\n\t") {
                                                case "\n\t":
                                                    fmt.Println(v3, v48, v5)
                                                    if v4 {
                                                        v147 := GenS12{F0: 4294967295}
                                                        _ = v147
                                                    } else {
                                                        if v4 {
                                                            if ((42 < 2) && (!v4)) {
                                                                var v148 float64 = float64((float32(0.1) - 1e10))
                                                                var v149 int16 = genFn0()
                                                                _ = v148
                                                                _ = v149
                                                            }
                                                            var v150 int8 = (v3 & v3)
                                                            _ = v150
                                                        }
                                                        fmt.Println(fi44)
                                                        v1 -= 1e10
                                                        v151 := GenS8{F0: 18446744073709551615, F1: 0}
                                                        v5 -= (v5 << ((v48) & 63))
                                                        if (v4 || (0.0 != 0.0)) {
                                                            var v152 int32 = (v5 * v5)
                                                            v2 = append(v2, 625)
                                                            _ = v152
                                                        } else {
                                                            var v153 float32 = float32((0.1 / 1.0))
                                                            _ = v153
                                                        }
                                                        var v154 int16 = genFn0()
                                                        _ = v151
                                                        _ = v154
                                                    }
                                                case "a":
                                                    fmt.Println(v4)
                                                }
                                            case "\n\t":
                                                fmt.Println(fi44, v48)
                                            case "a":
                                                v3 = (-(v3))
                                            }
                                        }
                                    default:
                                        var v155 float32 = v1
                                        _ = v155
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = fi47
                                    _ = v48
                                    _ = v72
                                }
                            default:
                                v5 = v5
                            }
                            v5 = (v5 << ((938) & 63))
                            fmt.Println(fi44, v1)
                        } else {
                            v1 = (float32(((1e10 / 1.0) / ((float64(1.5) * float64((float32(0.1) - 1e10))) - -2.25))) / (float32(0.1) - ((v1 * float32(0.0)) + 0.1)))
                            switch ("\n\t" + "") {
                            case "a":
                                fmt.Println(v5, v1)
                            }
                            if false {
                                switch ("\n\t" + "foo bar") {
                                case "foo bar":
                                    v3 -= v3
                                case "\n\t":
                                    if v4 {
                                        switch (fi44 * fi44) {
                                        case -598:
                                            v156 := GenS14{F0: genFn0(), F1: 2147483647}
                                            _ = v156
                                        case 886:
                                            v1 += float32((1.5 * 1.5))
                                        case 407:
                                            fmt.Println(v3)
                                        default:
                                            if (genFn0() > 32767) {
                                                var v157 uint16 = 1
                                                var v158 int = (fi44 % ((fi44) | 1))
                                                _ = len(v2)
                                                _ = v157
                                                _ = v158
                                            } else {
                                                v3 = v3
                                            }
                                        }
                                        v1 += 3.14159
                                    } else {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v3 -= v3
                                    }
                                case "hello":
                                    _ = len(v2)
                                default:
                                    v3 += 32
                                }
                                for fi159 := 0; fi159 < 5; fi159++ {
                                    v160 := GenS3{F0: 1, F1: 839}
                                    v3 += (v3 & v3)
                                    var v161 int16 = -32768
                                    v162 := GenS15{F0: v3}
                                    _ = fi159
                                    _ = v160
                                    _ = v161
                                    _ = v162
                                }
                                var v163 int64 = 1
                                _ = v163
                            }
                        }
                        _ = fi44
                    }
                }
                fmt.Println(v3, v4, v1)
            }
        case "foo bar":
            for fi164 := 0; fi164 < 20; fi164++ {
                v5 = 1
                var v165 string = "\n\t"
                _ = fi164
                _ = v165
            }
        default:
            var v166 int8 = (v3 * (v3 - v3))
            _ = v166
        }
    case 42:
        switch "\n\t" {
        case "":
            switch 9223372036854775807 {
            case 4294967295:
                for fi167 := 0; fi167 < 11; fi167++ {
                    v3 -= (v3 * 1)
                    v168 := GenS12{F0: 482}
                    _ = fi167
                    _ = v168
                }
            case 32767:
                fmt.Println(v5)
            }
        case "hello":
            for fi169 := 0; fi169 < 18; fi169++ {
                v3 -= (v3 - -1)
                fmt.Println(v5, v3)
                _ = fi169
            }
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    default:
        v3 = v3
    }
    var v170 uint32 = 4294967295
    var v171 uint32 = 42
    switch 42 {
    case -1:
        fmt.Println(v171, v3)
    case 4294967295:
        v172 := GenS4{F0: 127, F1: -2147483648}
        _ = v172
    case -137:
        if len(v2) > 0 {
            v2[0] = (^(v170))
        }
    }
    for fi173 := 0; fi173 < 19; fi173++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi173
    }
    for fi174 := 0; fi174 < 14; fi174++ {
        if v4 {
            v175 := GenS16{F0: 817, F1: 42}
            v171 = (v171 / ((2) | 1))
            _ = v175
        } else {
            if true {
                v1 = v1
                for fi176 := 0; fi176 < 2; fi176++ {
                    v3 += v3
                    var v177 uint8 = 42
                    v2 = append(v2, 1)
                    _ = fi176
                    _ = v177
                }
            } else {
                v178 := GenS16{F0: 2, F1: 985}
                _ = v178
            }
            for fi179 := 0; fi179 < 7; fi179++ {
                switch (fi179 >> ((0) & 63)) {
                case 1:
                    fmt.Println(v1, v5)
                case 0:
                    _ = len(v2)
                case -9223372036854775808:
                    _ = len(v2)
                }
                v4 = ("\n\t" != "\n\t")
                var v180 bool = true
                _ = fi179
                _ = v180
            }
            fmt.Println(v5)
        }
        _ = fi174
    }
    for fi181 := 0; fi181 < 7; fi181++ {
        v5 = (v5 & v5)
        v182 := GenS14{F0: -1, F1: 854}
        v5 = v5
        _ = fi181
        _ = v182
    }
    for fi183 := 0; fi183 < 2; fi183++ {
        v184 := GenS17{F0: "foo bar"}
        _ = fi183
        _ = v184
    }
    if ((255 <= 32767) && (160 >= 127)) {
        v170 -= (v170 % ((60) | 1))
        var v185 int16 = 0
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v185
    }
    switch 2147483647 {
    case -1:
        v171 += (^(v170))
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v170:", v170)
    fmt.Println("v171:", v171)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v170
    _ = v171
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
