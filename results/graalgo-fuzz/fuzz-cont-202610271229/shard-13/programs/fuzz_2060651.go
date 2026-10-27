package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint8; F1 int32; F2 int16 }
type GenS10 struct { F0 int32; F1 bool; F2 uint16 }
type GenS11 struct { F0 uint64; F1 uint16 }
type GenS12 struct { F0 int32; F1 uint16 }
type GenS13 struct { F0 int; F1 uint16; F2 bool }
type GenS14 struct { F0 uint8; F1 string; F2 int64 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 uint32; F1 string }
type GenS17 struct { F0 int8; F1 int8; F2 uint }
type GenS18 struct { F0 bool; F1 uint }
type GenS19 struct { F0 int64 }
type GenS20 struct { F0 bool; F1 uint8; F2 int16 }
type GenS21 struct { F0 bool; F1 string; F2 uint64 }
type GenS22 struct { F0 int32 }
func genFn0() uint {
    v7 := GenS1{F0: 499, F1: 4294967295}
    var v8 float32 = 0.1
    var v9 int64 = 9223372036854775807
    if (!(false || (42 >= 1))) {
        if (418 == 803) {
            var v10 uint32 = 2
            _ = v10
        } else {
            v11 := GenS2{F0: 2, F1: 1}
            _ = v11
        }
    } else {
        if (false || true) {
            v8 = 1.5
        } else {
            v12 := []int{0}
            if ((4294967295 <= 4294967295) && true) {
                v9 += (v9 & (^(v9)))
            } else {
                switch 128 {
                case -733:
                    var v13 int64 = (^(v9))
                    _ = v13
                }
                fmt.Println(v9, v9)
            }
            _ = v12
        }
        if ("" <= "foo bar") {
            v8 -= (0.0 / ((((v8 + (v8 - 0.0)) + (0.0 / 1.5)) / v8) * ((3.14159 + v8) + 3.14159)))
        } else {
            switch ("foo bar" + "a") {
            case "foo bar":
                v14 := []int{-380, 128, -842, int(v9)}
                _ = v14
            }
        }
    }
    _ = v7
    _ = v8
    _ = v9
    return 255
}

func main() {
    var v1 bool = (false && true)
    var v2 float64 = -2.25
    var v3 uint16 = 32767
    var v4 float64 = float64(-2.25)
    v5 := make(map[uint32]uint32)
    var v6 bool = (v1 && v1)
    v2 -= v2
    v1 = v6
    v4 += v2
    if v6 {
        v15 := GenS3{F0: (v3 * v3)}
        if v1 {
            switch -2147483648 {
            case 9223372036854775807:
                var v16 int16 = 738
                _ = v16
            default:
                v1 = (-32768 <= -447)
            }
        } else {
            switch 9223372036854775807 {
            case -9223372036854775808:
                if (15 >= 127) {
                    if (false && (!(v3 >= (v3 << ((445) & 63))))) {
                        fmt.Println(v1, v3, v6)
                        _ = len(v5)
                        var v17 bool = v6
                        _ = v17
                    }
                    v6 = v1
                    var v18 bool = v6
                    _ = v18
                }
            case -1:
                v3 = (v3 / (((v3 & (v3 % (((v3 - (^(v3)))) | 1)))) | 1))
            case 1:
                v19 := GenS1{F0: 198, F1: 42}
                _ = v19
            default:
                if (false || v6) {
                    if true {
                        for fi20 := 0; fi20 < 3; fi20++ {
                            var v21 bool = v6
                            for fi22 := 0; fi22 < 8; fi22++ {
                                fmt.Println(v4, v2, v2)
                                _ = fi22
                            }
                            _ = fi20
                            _ = v21
                        }
                        switch 4294967295 {
                        case 888:
                            v23 := v5[uint32(v3)]
                            _ = v23
                        case 128:
                            if (!(!v1)) {
                                v3 = 534
                                for fi24 := 0; fi24 < 18; fi24++ {
                                    if v6 {
                                        for fi25 := 0; fi25 < 10; fi25++ {
                                            if false {
                                                delete(v5, 765)
                                                fmt.Println(v4, v6)
                                                var v26 int = 128
                                                _ = v26
                                            } else {
                                                v27 := GenS4{F0: 42}
                                                _ = v27
                                            }
                                            v28 := GenS1{F0: -355, F1: 0}
                                            if true {
                                                var v29 uint = 4294967295
                                                _ = v29
                                            }
                                            var v30 float64 = v2
                                            v4 += v4
                                            v3 += (v3 >> ((295) & 63))
                                            fmt.Println(fi24, fi25, v6)
                                            _ = fi25
                                            _ = v28
                                            _ = v30
                                        }
                                    } else {
                                        for fi31 := 0; fi31 < 19; fi31++ {
                                            switch fi24 {
                                            case -1:
                                                if (474 <= 777) {
                                                    delete(v5, 65535)
                                                    v6 = (0 < -1)
                                                    switch "" {
                                                    case "foo bar":
                                                        v32 := []int{fi31, (fi31 | (fi24 * (fi24 * (fi31 << ((255) & 63)))))}
                                                        _ = v32
                                                    case "a":
                                                        var v33 float64 = (3.14159 * 3.14159)
                                                        _ = v33
                                                    case "":
                                                        v3 += v3
                                                    }
                                                }
                                            default:
                                                var v34 int16 = 42
                                                _ = v34
                                            }
                                            fmt.Println(fi24, v2, v2)
                                            _ = fi31
                                        }
                                    }
                                    var v35 string = ("hello" + "a")
                                    _ = fi24
                                    _ = v35
                                }
                            } else {
                                var v36 int = 127
                                var v37 uint = 861
                                delete(v5, 640)
                                _ = v36
                                _ = v37
                            }
                        default:
                            if (true || (v1 && (!((!(0 != 811)) && v1)))) {
                                var v38 uint16 = (v3 * 2)
                                for fi39 := 0; fi39 < 16; fi39++ {
                                    switch fi39 {
                                    case 1:
                                        for fi40 := 0; fi40 < 17; fi40++ {
                                            v41 := GenS5{F0: -704}
                                            if (("\n\t" > "a") && v6) {
                                                v42 := GenS5{F0: 1}
                                                v4 = v4
                                                v43 := GenS6{F0: 42, F1: 4294967295, F2: 42}
                                                fmt.Println(v4)
                                                if v6 {
                                                    if ((-259 <= 0) || v6) {
                                                        var v44 int = -1
                                                        v38 = v38
                                                        _ = v44
                                                    }
                                                    if (v6 && v1) {
                                                        var v45 bool = false
                                                        if v1 {
                                                            switch "" {
                                                            case "foo bar":
                                                                v46 := GenS7{F0: ((false && v45) && ((1 >= 764) || v6))}
                                                                _ = v46
                                                            case "\n\t":
                                                                var v47 int32 = 42
                                                                _ = v47
                                                            case "a":
                                                                var v48 uint = 619
                                                                _ = v48
                                                            }
                                                        } else {
                                                            if v45 {
                                                                if ((v3 & 0) >= (v3 ^ v3)) {
                                                                    var v49 uint64 = 4294967295
                                                                    var v50 uint32 = 2147483647
                                                                    switch ((("" + "foo bar") + "") + ("" + "\n\t")) {
                                                                    case "a":
                                                                        if false {
                                                                            v6 = v1
                                                                            fmt.Println(v6)
                                                                        }
                                                                    }
                                                                    _ = v49
                                                                    _ = v50
                                                                } else {
                                                                    v45 = (1000 >= 1000)
                                                                }
                                                                switch ("foo bar" + "\n\t") {
                                                                case "foo bar":
                                                                    v38 = 0
                                                                default:
                                                                    v38 = (^(v3))
                                                                }
                                                            }
                                                            v51 := v5[792]
                                                            var v52 int16 = 0
                                                            _ = v51
                                                            _ = v52
                                                        }
                                                        _ = v45
                                                    }
                                                    var v53 int8 = 42
                                                    var v54 int32 = 709
                                                    _ = v53
                                                    _ = v54
                                                } else {
                                                    switch 32767 {
                                                    case 32767:
                                                        v55 := GenS2{F0: 255, F1: 1}
                                                        _ = v55
                                                    case 127:
                                                        if true {
                                                            v56 := GenS8{F0: 1}
                                                            v57 := GenS4{F0: 1000}
                                                            var v58 int64 = 1
                                                            _ = v56
                                                            _ = v57
                                                            _ = v58
                                                        } else {
                                                            if ("" < ("a" + "foo bar")) {
                                                                v59 := []int{fi39, (fi40 & (fi40 + fi39)), (fi40 / ((4294967295) | 1))}
                                                                _ = len(v59)
                                                                _ = v59
                                                            }
                                                            v2 -= float64(1e10)
                                                            if (63 == 42) {
                                                                var v60 float32 = 0.1
                                                                _ = v60
                                                            }
                                                        }
                                                    case 32768:
                                                        if v6 {
                                                            switch 127 {
                                                            case 105:
                                                                v38 += (v3 + 479)
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    v61 := []int{32767, (fi40 + fi40), (fi40 / (((^(fi39))) | 1)), fi39}
                                                                    _ = v61
                                                                }
                                                            default:
                                                                if v1 {
                                                                    v1 = true
                                                                }
                                                            }
                                                        }
                                                    }
                                                    v1 = false
                                                    v38 = 0
                                                }
                                                if v6 {
                                                    if (1000 < 1000) {
                                                        if false {
                                                            if true {
                                                                var v62 float64 = (((-2.25 + v4) / v2) * (v2 / (-2.25 + v2)))
                                                                v38 = 1
                                                                v1 = (true && (32767 <= -1))
                                                                _ = v62
                                                            }
                                                            if (!(((v3 << ((428) & 63)) <= v38) || true)) {
                                                                v63 := []int{fi40}
                                                                var v64 uint = 2
                                                                var v65 uint16 = (v38 + 65535)
                                                                _ = v63
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                            v3 += (v38 ^ v3)
                                                            v6 = true
                                                            fmt.Println(v2, v3)
                                                            var v66 int = fi39
                                                            var v67 uint = 0
                                                            var v68 int8 = 42
                                                            v69 := GenS6{F0: 255, F1: v67, F2: 128}
                                                            _ = v66
                                                            _ = v67
                                                            _ = v68
                                                            _ = v69
                                                        } else {
                                                            if (-894 <= 816) {
                                                                switch (^(fi39)) {
                                                                case 840:
                                                                    var v70 uint32 = 1
                                                                    _ = v70
                                                                case 2147483647:
                                                                    if (v1 && ((false || false) && (true || true))) {
                                                                        v38 += (^(v3))
                                                                        var v71 uint16 = v38
                                                                        _ = v71
                                                                    } else {
                                                                        var v72 uint8 = 0
                                                                        if ((v72 * v72) < 128) {
                                                                            v6 = false
                                                                            v73 := []int{42}
                                                                            switch "\n\t" {
                                                                            case "a":
                                                                                fmt.Println(v6, v38, v38)
                                                                            case "hello":
                                                                                if (!v1) {
                                                                                    v74 := GenS9{F0: (v72 / (((^(v72))) | 1)), F1: 42, F2: 677}
                                                                                    if true {
                                                                                        if (fi39 <= (fi40 ^ (fi39 * (fi40 & -9223372036854775808)))) {
                                                                                            delete(v5, 2)
                                                                                        } else {
                                                                                            fmt.Println(v4, v6, v6)
                                                                                            v72 = 244
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(v2, v38)
                                                                                        fmt.Println(v3, v2)
                                                                                        switch -949 {
                                                                                        case 481:
                                                                                            v75 := GenS10{F0: -1, F1: v6, F2: v38}
                                                                                            _ = v75
                                                                                        case 836:
                                                                                            var v76 string = ("hello" + "foo bar")
                                                                                            var v77 string = ((v76 + v76) + "a")
                                                                                            var v78 string = ("hello" + "hello")
                                                                                            _ = v76
                                                                                            _ = v77
                                                                                            _ = v78
                                                                                        default:
                                                                                            var v79 string = "\n\t"
                                                                                            _ = v79
                                                                                        }
                                                                                    }
                                                                                    v38 -= v38
                                                                                    _ = v74
                                                                                }
                                                                            case "\n\t":
                                                                                delete(v5, 2)
                                                                            default:
                                                                                var v80 uint32 = 65535
                                                                                _ = v80
                                                                            }
                                                                            _ = v73
                                                                        } else {
                                                                            v81 := GenS11{F0: 4294967295, F1: v3}
                                                                            var v82 bool = v6
                                                                            switch fi40 {
                                                                            case 449:
                                                                                var v83 uint64 = 4294967295
                                                                                _ = v83
                                                                            case 4294967295:
                                                                                v72 = (v72 * 1)
                                                                            case -951:
                                                                                switch "foo bar" {
                                                                                case "foo bar":
                                                                                    v1 = (!false)
                                                                                }
                                                                            default:
                                                                                v82 = v6
                                                                            }
                                                                            _ = v81
                                                                            _ = v82
                                                                        }
                                                                        _ = v72
                                                                    }
                                                                case 9223372036854775807:
                                                                    switch 127 {
                                                                    case 0:
                                                                        switch fi40 {
                                                                        case -9223372036854775808:
                                                                            v2 -= 1e10
                                                                        case 32768:
                                                                            v84 := GenS12{F0: 42, F1: (v38 >> ((659) & 63))}
                                                                            _ = v84
                                                                        case 2147483647:
                                                                            v3 += v3
                                                                        default:
                                                                            v4 += float64(0.1)
                                                                        }
                                                                    case 390:
                                                                        delete(v5, 551)
                                                                    case 9223372036854775807:
                                                                        var v85 uint16 = 405
                                                                        _ = v85
                                                                    }
                                                                }
                                                                if (0 <= -800) {
                                                                    if false {
                                                                        if (65535 > 65535) {
                                                                            if v1 {
                                                                                var v86 uint8 = 42
                                                                                switch (fi40 & (^(fi39))) {
                                                                                case -1:
                                                                                    v87 := GenS5{F0: -373}
                                                                                    _ = v87
                                                                                case 1:
                                                                                    v88 := GenS13{F0: fi40, F1: (v3 << ((975) & 63)), F2: false}
                                                                                    _ = v88
                                                                                case -240:
                                                                                    var v89 uint32 = 4294967295
                                                                                    _ = v89
                                                                                }
                                                                                _ = v86
                                                                            } else {
                                                                                v90 := GenS1{F0: 1, F1: 1}
                                                                                switch "\n\t" {
                                                                                case "a":
                                                                                    if v1 {
                                                                                        v91 := GenS13{F0: (fi40 | fi40), F1: 195, F2: (fi39 == 0)}
                                                                                        v2 += v4
                                                                                        _ = v91
                                                                                    } else {
                                                                                        if true {
                                                                                            v4 = 0.1
                                                                                            v4 += v2
                                                                                            var v92 float32 = 0.0
                                                                                            _ = v92
                                                                                        }
                                                                                        v93 := GenS14{F0: 2, F1: "", F2: -1}
                                                                                        _ = v93
                                                                                    }
                                                                                case "foo bar":
                                                                                    fmt.Println(v3)
                                                                                case "\n\t":
                                                                                    switch "foo bar" {
                                                                                    case "hello":
                                                                                        fmt.Println(v38, v3)
                                                                                        v94 := []int{(fi39 + (fi40 + fi39))}
                                                                                        v95 := GenS12{F0: -2147483648, F1: 65535}
                                                                                        _ = v94
                                                                                        _ = v95
                                                                                    default:
                                                                                        v96 := GenS10{F0: -2147483648, F1: true, F2: v38}
                                                                                        _ = v96
                                                                                    }
                                                                                }
                                                                                fmt.Println(v2)
                                                                                _ = v90
                                                                            }
                                                                            fmt.Println(v38, v1, v6)
                                                                            if ((v1 || true) && (!v1)) {
                                                                                if (!(-2147483648 >= 2147483647)) {
                                                                                    fmt.Println(fi40, v3, v4)
                                                                                    if v6 {
                                                                                        var v97 int32 = -1
                                                                                        var v98 int64 = 42
                                                                                        _ = v97
                                                                                        _ = v98
                                                                                    }
                                                                                    v3 = 32767
                                                                                    if true {
                                                                                        v38 += v38
                                                                                        var v99 float64 = v4
                                                                                        _ = v99
                                                                                    } else {
                                                                                        var v100 int = (fi39 + fi39)
                                                                                        var v101 uint8 = 127
                                                                                        _ = v100
                                                                                        _ = v101
                                                                                    }
                                                                                    v2 = float64(float32(v4))
                                                                                } else {
                                                                                    v6 = v6
                                                                                }
                                                                                v102 := GenS13{F0: fi39, F1: 255, F2: true}
                                                                                switch -1 {
                                                                                case 32768:
                                                                                    v103 := GenS1{F0: 2147483647, F1: 65535}
                                                                                    _ = v103
                                                                                default:
                                                                                    v104 := GenS15{F0: 65535}
                                                                                    _ = v104
                                                                                }
                                                                                _ = v102
                                                                            } else {
                                                                                v4 = (v2 - v2)
                                                                                v105 := GenS6{F0: 127, F1: 42, F2: 1}
                                                                                _ = v105
                                                                            }
                                                                        }
                                                                        var v106 uint32 = 65535
                                                                        if (false || true) {
                                                                            v106 -= 2147483648
                                                                            v3 += v3
                                                                        }
                                                                        v5[(v106 << ((87) & 63))] = (v106 / (((v106 << ((112) & 63))) | 1))
                                                                        _ = v106
                                                                    } else {
                                                                        fmt.Println(v6, v3)
                                                                        var v107 float32 = 0.0
                                                                        if (true || (float32(v2) < v107)) {
                                                                            if ((v4 / (0.1 - v2)) != v2) {
                                                                                var v108 int8 = 127
                                                                                _ = v108
                                                                            } else {
                                                                                var v109 int64 = 1
                                                                                var v110 float32 = v107
                                                                                var v111 uint8 = 1
                                                                                v107 = 3.14159
                                                                                v112 := []int{(fi40 - (fi40 << ((65535) & 63))), 234, fi39}
                                                                                _ = v109
                                                                                _ = v110
                                                                                _ = v111
                                                                                _ = v112
                                                                            }
                                                                        } else {
                                                                            v38 += 541
                                                                            var v113 string = "foo bar"
                                                                            v3 -= (^(v38))
                                                                            _ = v113
                                                                        }
                                                                        v114 := GenS16{F0: 65535, F1: ("a" + ("foo bar" + ("" + "foo bar")))}
                                                                        v115 := []int{}
                                                                        _ = v107
                                                                        _ = v114
                                                                        _ = v115
                                                                    }
                                                                }
                                                                fmt.Println(fi40, v38, v2)
                                                            }
                                                            if v6 {
                                                                v6 = false
                                                                v116 := GenS7{F0: (42 > 0)}
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    switch (fi40 + 0) {
                                                                    case -401:
                                                                        var v117 int32 = -1
                                                                        _ = v117
                                                                    case -2147483648:
                                                                        if ((("hello" + ("a" + "a")) + "foo bar") == ("" + "a")) {
                                                                            if v1 {
                                                                                var v118 int16 = 0
                                                                                _ = v118
                                                                            } else {
                                                                                if (!(3.14159 == (float32(v4) - (1.5 - 0.0)))) {
                                                                                    var v119 uint64 = 0
                                                                                    fmt.Println(v2)
                                                                                    var v120 int64 = 413
                                                                                    _ = v119
                                                                                    _ = v120
                                                                                } else {
                                                                                    fmt.Println(v4, v2, fi39)
                                                                                    v2 = v2
                                                                                }
                                                                            }
                                                                            v5[972] = 65535
                                                                        } else {
                                                                            v121 := GenS17{F0: -1, F1: 96, F2: 0}
                                                                            _ = v121
                                                                        }
                                                                    case 128:
                                                                        var v122 int64 = 0
                                                                        _ = v122
                                                                    default:
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            var v123 int32 = 2147483647
                                                                            _ = v123
                                                                        }
                                                                    }
                                                                    v124 := GenS18{F0: v1, F1: 2}
                                                                    fmt.Println(v2, v1)
                                                                    v2 = v4
                                                                    v5[354] = 0
                                                                    _ = v124
                                                                default:
                                                                    delete(v5, 2)
                                                                    if (!false) {
                                                                        v125 := GenS19{F0: 642}
                                                                        var v126 int32 = -1
                                                                        if v6 {
                                                                            switch 358 {
                                                                            case 128:
                                                                                delete(v5, 799)
                                                                                v1 = v6
                                                                            case 1:
                                                                                var v127 int16 = 42
                                                                                var v128 uint8 = 2
                                                                                _ = v127
                                                                                _ = v128
                                                                            default:
                                                                                var v129 uint32 = 65535
                                                                                _ = v129
                                                                            }
                                                                        }
                                                                        v130 := []int{int(v126), (-(fi40)), fi40}
                                                                        if len(v130) >= 2 {
                                                                            v130 = v130[1:len(v130)]
                                                                        }
                                                                        v38 += v38
                                                                        v131 := v5[50]
                                                                        _ = v125
                                                                        _ = v126
                                                                        _ = v130
                                                                        _ = v131
                                                                    } else {
                                                                        v5[904] = 2147483647
                                                                        var v132 uint32 = 521
                                                                        if v1 {
                                                                            fmt.Println(v2, v132)
                                                                        } else {
                                                                            v1 = (32767 == fi40)
                                                                        }
                                                                        _ = v132
                                                                    }
                                                                }
                                                                _ = v116
                                                            }
                                                        }
                                                        if ((42 <= 505) && ((v38 % ((v3) | 1)) == 472)) {
                                                            v4 += (float64(0.0) - (3.14159 * v4))
                                                            v3 = (v38 ^ (v38 >> ((265) & 63)))
                                                        } else {
                                                            v133 := []int{fi40, (fi40 << ((4294967295) & 63))}
                                                            _ = v133
                                                        }
                                                    }
                                                    if v1 {
                                                        v2 -= (((v4 / v2) * (0.0 - v2)) / float64(1.5))
                                                        v134 := GenS5{F0: 0}
                                                        _ = v134
                                                    }
                                                } else {
                                                    v135 := GenS16{F0: 34, F1: ("\n\t" + "\n\t")}
                                                    _ = v135
                                                }
                                                v136 := GenS20{F0: (127 != 1), F1: 0, F2: -32768}
                                                _ = v136
                                                _ = v42
                                                _ = v43
                                            } else {
                                                v3 = v38
                                                delete(v5, 2)
                                            }
                                            if v6 {
                                                v2 -= (v4 * v4)
                                                v4 = (v2 + v2)
                                            } else {
                                                var v137 uint64 = 42
                                                _ = v137
                                            }
                                            _ = fi40
                                            _ = v41
                                        }
                                    default:
                                        v4 = ((v2 / ((v2 + ((1e10 / v4) / 1.5)) / v4)) - v4)
                                    }
                                    fmt.Println(v4, v1)
                                    _ = len(v5)
                                    _ = fi39
                                }
                                v138 := GenS10{F0: 0, F1: v1, F2: v3}
                                _ = v138
                                _ = v38
                            } else {
                                fmt.Println(v4, v4, v3)
                            }
                        }
                    }
                    v139 := GenS21{F0: (!(-304 != 895)), F1: (("a" + "\n\t") + ("hello" + ("\n\t" + (("foo bar" + "foo bar") + "a")))), F2: 4294967295}
                    var v140 float64 = 0.0
                    _ = v139
                    _ = v140
                }
            }
        }
        var v141 int8 = 42
        _ = v141
        _ = v15
    } else {
        var v142 bool = (42 >= 42)
        _ = v142
    }
    if (v1 && true) {
        v143 := GenS8{F0: 598}
        v1 = (4294967295 > 298)
        for fi144 := 0; fi144 < 14; fi144++ {
            v2 = v4
            _ = fi144
        }
        _ = v143
    } else {
        v6 = ("a" <= "")
        v145 := []int{32768, 2147483647}
        for fi146 := 0; fi146 < 10; fi146++ {
            _ = len(v5)
            v2 = (float64((-2.25 + ((0.1 * (1.5 - (float32((0.0 + 0.1)) / 0.1))) * 0.0))) - v2)
            _ = fi146
        }
        _ = v145
    }
    var v147 uint16 = (v3 >> ((780) & 63))
    var v148 float32 = ((-2.25 - 3.14159) * float32(float64(0.0)))
    fmt.Println(v3)
    var v149 uint64 = 18446744073709551615
    var v150 float32 = v148
    fmt.Println(v150)
    v147 = (v3 % ((v3) | 1))
    switch 32767 {
    case 42:
        v147 -= 65535
    case 987:
        var v151 uint16 = (v3 >> ((genFn0()) & 63))
        _ = v151
    case 2147483647:
        var v152 int8 = 127
        _ = v152
    default:
        switch 718 {
        case 32767:
            v153 := GenS14{F0: 37, F1: "hello", F2: 730}
            _ = v153
        default:
            fmt.Println(v147, v6)
        }
    }
    fmt.Println(v3, v1)
    if ((-13 >= 721) || true) {
        fmt.Println(v149, v6, v148)
        var v154 uint8 = 211
        switch ("" + ("\n\t" + ("foo bar" + ""))) {
        case "a":
            var v155 uint32 = 65535
            _ = v155
        case "hello":
            v1 = false
        default:
            v156 := GenS22{F0: 2147483647}
            _ = v156
        }
        _ = v154
    } else {
        if ((393 == 42) || v1) {
            var v157 bool = v1
            var v158 uint8 = 0
            _ = v157
            _ = v158
        } else {
            var v159 int64 = -1
            _ = v159
        }
        for fi160 := 0; fi160 < 2; fi160++ {
            fmt.Println(v4, fi160)
            _ = fi160
        }
        v161 := []int{1, -609, 2147483647, 1}
        _ = v161
    }
    switch 2147483647 {
    case -9223372036854775808:
        var v162 int32 = 42
        _ = v162
    case -269:
        v147 += (v147 >> ((255) & 63))
    default:
        delete(v5, 890)
    }
    if (true || (!true)) {
        var v163 int = 32767
        v1 = v6
        delete(v5, 2147483648)
        _ = v163
    }
    var v164 uint = genFn0()
    fmt.Println("v1:", v1)
    fmt.Println("v147:", v147)
    fmt.Println("v148:", v148)
    fmt.Println("v149:", v149)
    fmt.Println("v150:", v150)
    fmt.Println("v164:", v164)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v147
    _ = v148
    _ = v149
    _ = v150
    _ = v164
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
