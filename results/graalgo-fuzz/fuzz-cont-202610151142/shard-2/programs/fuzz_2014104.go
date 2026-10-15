package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 bool; F1 int64; F2 string }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint16; F1 uint16 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int16; F1 bool; F2 string }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int; F1 int32 }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 uint32 }
type GenS19 struct { F0 int16; F1 uint16 }
type GenS20 struct { F0 int16 }
type GenS21 struct { F0 int8 }
type GenS22 struct { F0 int; F1 int64; F2 uint32 }
type GenS23 struct { F0 uint }
func genFn0(p0 uint64) int16 {
    var v7 float64 = (3.14159 * (((1.5 + 1.5) - 0.1) + 1e10))
    var v8 bool = (true && false)
    v9 := GenS1{F0: (42 == 2), F1: 42}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -790
}
func genFn1(p0 int64) string {
    if true {
        fmt.Println(p0)
        fmt.Println(p0)
    } else {
        if (0 <= 63) {
            if false {
                var v10 uint = 42
                _ = v10
            } else {
                p0 = p0
                if (-2.25 == float32(0.0)) {
                    fmt.Println(p0)
                } else {
                    for fi11 := 0; fi11 < 8; fi11++ {
                        p0 += p0
                        v12 := GenS1{F0: (-490 <= 42), F1: 2}
                        _ = fi11
                        _ = v12
                    }
                    v13 := GenS1{F0: true, F1: 581}
                    _ = v13
                }
                v14 := GenS1{F0: (4294967295 < 1), F1: 18446744073709551615}
                _ = v14
            }
            var v15 uint32 = 2
            var v16 bool = (255 == 42)
            _ = v15
            _ = v16
        } else {
            p0 -= 4
        }
    }
    v17 := []int{128}
    _ = len(v17)
    if true {
        switch 725 {
        case 127:
            _ = len(v17)
        }
    } else {
        for fi18 := 0; fi18 < 19; fi18++ {
            if ((!true) && false) {
                switch (("" + "foo bar") + "foo bar") {
                case "hello":
                    var v19 int32 = 2147483647
                    _ = v19
                case "foo bar":
                    fmt.Println(fi18)
                default:
                    switch (("hello" + "a") + (("a" + "hello") + "")) {
                    case "foo bar":
                        if true {
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                            fmt.Println(p0, fi18)
                        }
                    case "\n\t":
                        p0 = -1
                    default:
                        p0 = (p0 - p0)
                    }
                }
                for fi20 := 0; fi20 < 8; fi20++ {
                    if (true || (true || true)) {
                        var v21 int64 = p0
                        var v22 float64 = 1.5
                        if (!true) {
                            v23 := GenS1{F0: false, F1: 2}
                            var v24 uint8 = 30
                            var v25 int32 = 42
                            _ = v23
                            _ = v24
                            _ = v25
                        } else {
                            switch fi20 {
                            case -351:
                                fmt.Println(fi18, fi20, v21)
                            case 1:
                                switch -25 {
                                case 42:
                                    fmt.Println(v21, fi18)
                                case -9223372036854775808:
                                    if false {
                                        _ = len(v17)
                                        switch fi20 {
                                        case 844:
                                            fmt.Println(fi18, v22, fi20)
                                        case 275:
                                            v26 := GenS1{F0: false, F1: 4294967295}
                                            _ = v26
                                        case 859:
                                            fmt.Println(fi20, v22, fi20)
                                        default:
                                            var v27 float32 = 0.0
                                            _ = v27
                                        }
                                        var v28 uint = 1
                                        var v29 int8 = 42
                                        if false {
                                            p0 = (v21 | v21)
                                        } else {
                                            switch "foo bar" {
                                            case "foo bar":
                                                if ((!(3.14159 > 0.0)) && true) {
                                                    var v30 uint32 = 0
                                                    var v31 bool = false
                                                    v32 := GenS2{F0: genFn0(4294967295)}
                                                    v33 := GenS3{F0: (v28 + (v28 + (v28 & 1)))}
                                                    _ = v30
                                                    _ = v31
                                                    _ = v32
                                                    _ = v33
                                                } else {
                                                    switch "a" {
                                                    case "":
                                                        v29 += v29
                                                    }
                                                }
                                            case "a":
                                                v34 := GenS3{F0: v28}
                                                _ = v34
                                            case "hello":
                                                v29 = (v29 & (v29 * v29))
                                            }
                                            switch ("\n\t" + ("" + "")) {
                                            case "a":
                                                v28 += v28
                                            case "foo bar":
                                                v35 := GenS4{F0: "\n\t", F1: "a"}
                                                _ = v35
                                            }
                                        }
                                        v17 = append(v17, (fi18 & fi18))
                                        var v36 int8 = v29
                                        _ = v28
                                        _ = v29
                                        _ = v36
                                    }
                                default:
                                    switch fi20 {
                                    case 9223372036854775807:
                                        var v37 int32 = 42
                                        _ = v37
                                    case -373:
                                        if (-798 == 859) {
                                            v17 = append(v17, fi18)
                                        } else {
                                            if true {
                                                if false {
                                                    _ = len(v17)
                                                    if (2147483647 != 0) {
                                                        v38 := GenS4{F0: "foo bar", F1: ""}
                                                        if len(v17) > 0 {
                                                            v17[0] = fi20
                                                        }
                                                        _ = v38
                                                    }
                                                }
                                                var v39 uint16 = 255
                                                _ = v39
                                            }
                                            v22 -= float64(0.1)
                                            p0 = p0
                                        }
                                    }
                                }
                            default:
                                v17 = append(v17, fi20)
                            }
                            p0 -= p0
                            if false {
                                switch (^(fi20)) {
                                case 42:
                                    if false {
                                        v21 += (p0 | p0)
                                        if len(v17) >= 2 {
                                            v17 = v17[1:len(v17)]
                                        }
                                        fmt.Println(p0, v21)
                                    } else {
                                        var v40 string = (("foo bar" + "\n\t") + "a")
                                        if true {
                                            _ = len(v17)
                                        } else {
                                            if false {
                                                var v41 string = v40
                                                _ = v41
                                            }
                                            fmt.Println(fi20, fi20, v22)
                                            if false {
                                                v42 := GenS4{F0: ("a" + (v40 + "a")), F1: (((((v40 + "") + v40) + "") + v40) + v40)}
                                                _ = v42
                                            }
                                        }
                                        _ = v40
                                    }
                                }
                            } else {
                                switch (fi20 << ((42) & 63)) {
                                case 9223372036854775807:
                                    if false {
                                        v43 := GenS4{F0: "a", F1: ""}
                                        v17 = append(v17, fi20)
                                        _ = v43
                                    } else {
                                        if false {
                                            var v44 float32 = 0.0
                                            p0 = v21
                                            if true {
                                                v45 := GenS5{F0: true, F1: v21, F2: "\n\t"}
                                                _ = v45
                                            } else {
                                                v44 = (float32(v22) - -2.25)
                                                var v46 uint64 = 4294967295
                                                fmt.Println(p0)
                                                var v47 int32 = 488
                                                v44 += 0.1
                                                _ = v46
                                                _ = v47
                                            }
                                            if false {
                                                fmt.Println(v21, v21)
                                                var v48 uint32 = 2
                                                _ = v48
                                            }
                                            v44 -= (((3.14159 - (0.0 - 3.14159)) * v44) + (v44 + 1e10))
                                            var v49 uint64 = 1
                                            var v50 int8 = 100
                                            _ = v44
                                            _ = v49
                                            _ = v50
                                        } else {
                                            switch (fi18 + (fi18 + 377)) {
                                            case 32768:
                                                p0 = v21
                                            case 1:
                                                var v51 int64 = (p0 >> ((255) & 63))
                                                _ = v51
                                            default:
                                                if false {
                                                    if (255 <= 65535) {
                                                        fmt.Println(v21)
                                                        var v52 uint8 = 54
                                                        if false {
                                                            if (2 > 339) {
                                                                switch fi18 {
                                                                case -9223372036854775808:
                                                                    if (0.0 <= 0.0) {
                                                                        v21 = (-(v21))
                                                                        v21 -= p0
                                                                    }
                                                                default:
                                                                    var v53 uint64 = 2
                                                                    _ = v53
                                                                }
                                                            }
                                                            if len(v17) >= 2 {
                                                                v17 = v17[1:len(v17)]
                                                            }
                                                            v17 = append(v17, fi20)
                                                        }
                                                        _ = v52
                                                    } else {
                                                        if (18446744073709551615 >= 4294967295) {
                                                            v21 = (p0 % ((-9223372036854775808) | 1))
                                                        } else {
                                                            var v54 bool = false
                                                            v54 = (true && v54)
                                                            v21 = v21
                                                            _ = v54
                                                        }
                                                        switch (("hello" + "hello") + "a") {
                                                        case "a":
                                                            v55 := GenS4{F0: "foo bar", F1: "foo bar"}
                                                            _ = v55
                                                        case "foo bar":
                                                            switch (fi18 ^ fi20) {
                                                            case -592:
                                                                v21 = v21
                                                            case 2147483647:
                                                                if true {
                                                                    v21 -= p0
                                                                }
                                                            case -1:
                                                                v56 := GenS5{F0: false, F1: (-(p0)), F2: ""}
                                                                _ = v56
                                                            }
                                                        case "hello":
                                                            v21 = 0
                                                        }
                                                        if true {
                                                            _ = len(v17)
                                                            var v57 bool = false
                                                            if (!v57) {
                                                                var v58 int32 = 0
                                                                _ = v58
                                                            }
                                                            _ = v57
                                                        }
                                                        p0 -= 0
                                                    }
                                                    var v59 bool = false
                                                    var v60 float64 = (v22 - v22)
                                                    _ = v59
                                                    _ = v60
                                                }
                                            }
                                        }
                                        var v61 string = "\n\t"
                                        if (!(false || true)) {
                                            fmt.Println(v61, p0, fi18)
                                            fmt.Println(v21, v22, fi20)
                                        }
                                        _ = v61
                                    }
                                case 4294967295:
                                    v62 := GenS5{F0: ((fi20 / ((9223372036854775807) | 1)) > fi20), F1: (v21 % ((1) | 1)), F2: ("a" + "a")}
                                    _ = v62
                                }
                            }
                        }
                        v17 = append(v17, fi18)
                        fmt.Println(v21)
                        _ = v21
                        _ = v22
                    }
                    _ = fi20
                }
                if false {
                    if (false && true) {
                        p0 -= (p0 | (p0 & (p0 - p0)))
                        switch ("\n\t" + "\n\t") {
                        case "foo bar":
                            v63 := GenS4{F0: "hello", F1: ""}
                            _ = v63
                        case "":
                            for fi64 := 0; fi64 < 18; fi64++ {
                                var v65 string = "foo bar"
                                if false {
                                    if true {
                                        p0 = -469
                                        v17 = append(v17, fi64)
                                        p0 = (p0 << ((1) & 63))
                                    }
                                    var v66 uint8 = 106
                                    p0 -= p0
                                    fmt.Println(v66, v65, p0)
                                    v65 = v65
                                    var v67 uint8 = 206
                                    switch fi18 {
                                    case 0:
                                        var v68 int = fi18
                                        _ = v68
                                    case 1:
                                        switch v65 {
                                        case "a":
                                            var v69 uint16 = 65535
                                            _ = v69
                                        }
                                    default:
                                        var v70 uint16 = 32767
                                        _ = v70
                                    }
                                    var v71 uint = 255
                                    _ = v66
                                    _ = v67
                                    _ = v71
                                }
                                var v72 uint64 = 4294967295
                                _ = fi64
                                _ = v65
                                _ = v72
                            }
                        }
                    } else {
                        for fi73 := 0; fi73 < 17; fi73++ {
                            switch (-(fi18)) {
                            case -387:
                                _ = len(v17)
                            case 32767:
                                p0 = 770
                                if (!((true && true) || ((fi73 / ((-2147483648) | 1)) < fi18))) {
                                    p0 = p0
                                }
                                var v74 bool = false
                                fmt.Println(v74)
                                var v75 uint64 = 4294967295
                                _ = v74
                                _ = v75
                            case 370:
                                v76 := GenS6{F0: (127 >= 26)}
                                _ = v76
                            default:
                                switch "\n\t" {
                                case "\n\t":
                                    switch "\n\t" {
                                    case "hello":
                                        var v77 int64 = (p0 ^ -765)
                                        var v78 bool = (1e10 != (-2.25 * 1e10))
                                        switch ("foo bar" + (("a" + (("" + "foo bar") + "\n\t")) + "hello")) {
                                        case "foo bar":
                                            var v79 uint = 42
                                            _ = v79
                                        case "a":
                                            if len(v17) > 0 {
                                                v17[0] = fi18
                                            }
                                        case "":
                                            v78 = false
                                        default:
                                            fmt.Println(v77, v78, fi18)
                                        }
                                        _ = v77
                                        _ = v78
                                    case "\n\t":
                                        v17 = append(v17, fi73)
                                        p0 -= p0
                                    case "a":
                                        if (65535 <= 65535) {
                                            if (1 < 1) {
                                                p0 -= (p0 & (p0 ^ p0))
                                                if (((false || (!((2 == 2) || (!false)))) || (!((!(-430 >= 2147483647)) || true))) || (353 > -807)) {
                                                    var v80 bool = true
                                                    switch "\n\t" {
                                                    case "a":
                                                        fmt.Println(p0, p0, v80)
                                                    case "foo bar":
                                                        var v81 string = ("foo bar" + ("\n\t" + "hello"))
                                                        _ = v81
                                                    }
                                                    switch "" {
                                                    case "foo bar":
                                                        var v82 int16 = 0
                                                        _ = v82
                                                    case "a":
                                                        v83 := GenS7{F0: 1}
                                                        _ = v83
                                                    case "\n\t":
                                                        var v84 int16 = 42
                                                        _ = v84
                                                    }
                                                    var v85 int = fi73
                                                    var v86 int16 = 1
                                                    _ = v80
                                                    _ = v85
                                                    _ = v86
                                                } else {
                                                    var v87 uint = 255
                                                    _ = v87
                                                }
                                            }
                                            _ = len(v17)
                                            var v88 float64 = (1e10 / (0.1 - 0.0))
                                            _ = v88
                                        }
                                        if true {
                                            if len(v17) > 0 {
                                                v17[0] = fi18
                                            }
                                            if true {
                                                v89 := GenS1{F0: false, F1: 255}
                                                if len(v17) > 0 {
                                                    v17[0] = fi73
                                                }
                                                _ = v89
                                            }
                                            if true {
                                                if len(v17) > 0 {
                                                    v17[0] = (^(fi18))
                                                }
                                                if (!(-32768 < 902)) {
                                                    if ((1e10 * float64(3.14159)) != 3.14159) {
                                                        switch fi18 {
                                                        case 4294967295:
                                                            if len(v17) >= 2 {
                                                                v17 = v17[1:len(v17)]
                                                            }
                                                        case 926:
                                                            if true {
                                                                var v90 uint = 1
                                                                var v91 uint = 31
                                                                if (32767 >= 0) {
                                                                    var v92 int64 = (p0 - 0)
                                                                    var v93 int16 = 823
                                                                    p0 = 9223372036854775807
                                                                    _ = v92
                                                                    _ = v93
                                                                } else {
                                                                    fmt.Println(v90, p0, fi73)
                                                                    var v94 int16 = 379
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        var v95 int64 = 0
                                                                        _ = v95
                                                                    case "foo bar":
                                                                        var v96 uint = v90
                                                                        _ = v96
                                                                    }
                                                                    _ = v94
                                                                }
                                                                _ = v90
                                                                _ = v91
                                                            } else {
                                                                if true {
                                                                    v97 := GenS2{F0: 32767}
                                                                    if (1.5 == 0.1) {
                                                                        var v98 uint16 = 65535
                                                                        switch (^(fi18)) {
                                                                        case 1:
                                                                            p0 = p0
                                                                        default:
                                                                            var v99 string = (("\n\t" + ("hello" + "\n\t")) + "")
                                                                            v99 = v99
                                                                            _ = v99
                                                                        }
                                                                        var v100 uint32 = 691
                                                                        fmt.Println(v100, fi18, p0)
                                                                        _ = v100
                                                                        _ = v98
                                                                    }
                                                                    _ = v97
                                                                }
                                                                if len(v17) >= 2 {
                                                                    v17 = v17[1:len(v17)]
                                                                }
                                                                p0 -= 1
                                                            }
                                                        default:
                                                            var v101 int16 = -218
                                                            _ = v101
                                                        }
                                                        if (((p0 - (p0 << ((787) & 63))) > (p0 + 9223372036854775807)) || (205 == p0)) {
                                                            var v102 float64 = (0.0 / (1.5 + 0.0))
                                                            if false {
                                                                v103 := GenS5{F0: (-1 <= -1), F1: -772, F2: (("foo bar" + "hello") + "\n\t")}
                                                                if len(v17) > 0 {
                                                                    v17[0] = fi18
                                                                }
                                                                _ = v103
                                                            }
                                                            _ = v102
                                                        }
                                                    }
                                                    if ((59 == 224) || (90 > 489)) {
                                                        p0 = p0
                                                        var v104 uint64 = 2
                                                        v105 := GenS1{F0: ((0 >= 100) && (!true)), F1: 833}
                                                        _ = v104
                                                        _ = v105
                                                    } else {
                                                        fmt.Println(fi73)
                                                    }
                                                    var v106 uint32 = 65535
                                                    _ = v106
                                                }
                                                if ("a" >= "\n\t") {
                                                    var v107 uint16 = 728
                                                    _ = len(v17)
                                                    _ = v107
                                                }
                                                if false {
                                                    if false {
                                                        var v108 string = "hello"
                                                        _ = v108
                                                    } else {
                                                        p0 -= 1
                                                        if false {
                                                            v109 := GenS8{F0: 648, F1: 2}
                                                            _ = v109
                                                        } else {
                                                            if true {
                                                                if (1 != int16(fi18)) {
                                                                    var v110 int32 = -1
                                                                    var v111 float32 = 0.0
                                                                    _ = v110
                                                                    _ = v111
                                                                } else {
                                                                    v112 := GenS4{F0: "", F1: ("\n\t" + "a")}
                                                                    var v113 string = (("foo bar" + "") + "\n\t")
                                                                    var v114 int16 = 0
                                                                    _ = v112
                                                                    _ = v113
                                                                    _ = v114
                                                                }
                                                            } else {
                                                                fmt.Println(fi18)
                                                                p0 = (p0 * (^(p0)))
                                                                var v115 uint32 = 0
                                                                fmt.Println(fi18, fi73, fi73)
                                                                _ = v115
                                                            }
                                                            var v116 uint = 629
                                                            v117 := GenS9{F0: 1}
                                                            if ((fi18 | (fi73 ^ (fi73 >> ((v116) & 63)))) < (fi73 - -690)) {
                                                                fmt.Println(fi18, fi18)
                                                            } else {
                                                                fmt.Println(fi73)
                                                                if (false && false) {
                                                                    v116 = (v116 | v116)
                                                                    switch ("a" + "\n\t") {
                                                                    case "\n\t":
                                                                        p0 += (p0 | -432)
                                                                    case "foo bar":
                                                                        v17 = append(v17, -1)
                                                                    case "hello":
                                                                        switch 42 {
                                                                        case 32767:
                                                                            var v118 uint32 = 2147483647
                                                                            _ = v118
                                                                        case -1:
                                                                            v116 = 0
                                                                        case 9223372036854775807:
                                                                            var v119 uint64 = 1
                                                                            _ = v119
                                                                        }
                                                                        v116 = v116
                                                                    default:
                                                                        v116 = v116
                                                                    }
                                                                } else {
                                                                    var v120 int = fi18
                                                                    _ = v120
                                                                }
                                                                var v121 uint = (^(v116))
                                                                _ = v121
                                                            }
                                                            _ = len(v17)
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                var v122 uint64 = 0
                                                                _ = v122
                                                            case "":
                                                                if false {
                                                                    fmt.Println(v116, fi73)
                                                                } else {
                                                                    v123 := GenS6{F0: ((1 >= 4294967295) && (!true))}
                                                                    var v124 uint16 = 401
                                                                    p0 -= 0
                                                                    v125 := GenS10{F0: 42, F1: true, F2: ("hello" + "foo bar")}
                                                                    _ = v123
                                                                    _ = v124
                                                                    _ = v125
                                                                }
                                                            case "a":
                                                                if len(v17) > 0 {
                                                                    v17[0] = (fi18 << ((v116) & 63))
                                                                }
                                                            default:
                                                                fmt.Println(fi18)
                                                            }
                                                            p0 = p0
                                                            if len(v17) > 0 {
                                                                v17[0] = fi18
                                                            }
                                                            _ = v116
                                                            _ = v117
                                                        }
                                                        p0 += p0
                                                        if true {
                                                            if true {
                                                                var v126 uint = 707
                                                                if false {
                                                                    if false {
                                                                        v126 = 2
                                                                        var v127 string = "hello"
                                                                        _ = v127
                                                                    } else {
                                                                        v128 := GenS11{F0: 161}
                                                                        v129 := GenS5{F0: (1 < 958), F1: (p0 << (((v126 >> ((961) & 63))) & 63)), F2: ("foo bar" + (("foo bar" + "foo bar") + ("a" + "hello")))}
                                                                        if (-1 <= 100) {
                                                                            var v130 uint16 = 2
                                                                            if (0.0 <= (0.1 / (1.5 - 3.14159))) {
                                                                                var v131 uint8 = 148
                                                                                if (4294967295 > 4294967295) {
                                                                                    var v132 float64 = ((float64(-2.25) - 0.0) + float64(1.5))
                                                                                    _ = v132
                                                                                }
                                                                                _ = v131
                                                                            }
                                                                            v126 = 517
                                                                            var v133 uint16 = 466
                                                                            p0 = p0
                                                                            _ = v130
                                                                            _ = v133
                                                                        } else {
                                                                            if ((73 < 0) || (0 >= 1)) {
                                                                                v126 -= (v126 - v126)
                                                                                if len(v17) >= 2 {
                                                                                    v17 = v17[1:len(v17)]
                                                                                }
                                                                            } else {
                                                                                v126 += v126
                                                                                var v134 int8 = 0
                                                                                _ = v134
                                                                            }
                                                                            v126 = (v126 | (v126 >> ((v126) & 63)))
                                                                        }
                                                                        var v135 string = ""
                                                                        _ = v128
                                                                        _ = v129
                                                                        _ = v135
                                                                    }
                                                                    v126 = 42
                                                                }
                                                                if (false || true) {
                                                                    var v136 uint16 = 2
                                                                    _ = v136
                                                                }
                                                                _ = v126
                                                            } else {
                                                                if (2 < 789) {
                                                                    p0 = (p0 & p0)
                                                                    if len(v17) > 0 {
                                                                        v17[0] = 128
                                                                    }
                                                                    var v137 string = ("\n\t" + ("\n\t" + ("" + "foo bar")))
                                                                    _ = v137
                                                                } else {
                                                                    v138 := GenS12{F0: 0, F1: 127, F2: genFn0(18446744073709551615)}
                                                                    p0 = p0
                                                                    _ = v138
                                                                }
                                                                var v139 string = ("" + "foo bar")
                                                                v140 := GenS13{F0: ((v139 + v139) + v139)}
                                                                if true {
                                                                    v139 = (v139 + v139)
                                                                    var v141 float32 = 0.0
                                                                    _ = v141
                                                                }
                                                                if ((true || (18446744073709551615 >= 1)) && (484 < 0)) {
                                                                    if false {
                                                                        v142 := GenS3{F0: 129}
                                                                        if (48 < 0) {
                                                                            _ = len(v17)
                                                                            var v143 bool = (432 != 1)
                                                                            if len(v17) >= 2 {
                                                                                v17 = v17[1:len(v17)]
                                                                            }
                                                                            v139 = v139
                                                                            _ = v143
                                                                        }
                                                                        var v144 bool = (((((fi18 + fi18) > 389) && false) && true) || false)
                                                                        if ((0.0 + ((0.1 / 1.0) / 1.0)) >= 3.14159) {
                                                                            var v145 string = ("foo bar" + v139)
                                                                            _ = v145
                                                                        } else {
                                                                            v144 = v144
                                                                            var v146 string = ((v139 + (((v139 + "foo bar") + ("hello" + v139)) + v139)) + "")
                                                                            var v147 int64 = p0
                                                                            _ = v146
                                                                            _ = v147
                                                                        }
                                                                        var v148 bool = v144
                                                                        _ = v142
                                                                        _ = v144
                                                                        _ = v148
                                                                    }
                                                                    if false {
                                                                        var v149 uint64 = 668
                                                                        switch "a" {
                                                                        case "hello":
                                                                            v17 = append(v17, -1)
                                                                        case "foo bar":
                                                                            v150 := GenS12{F0: 42, F1: 1, F2: 32767}
                                                                            _ = v150
                                                                        }
                                                                        _ = v149
                                                                    }
                                                                    var v151 uint64 = 509
                                                                    _ = v151
                                                                }
                                                                var v152 uint8 = 255
                                                                _ = v139
                                                                _ = v140
                                                                _ = v152
                                                            }
                                                            p0 += p0
                                                            p0 -= (p0 + (p0 & 42))
                                                        } else {
                                                            switch (fi18 & (fi73 % (((^(fi18))) | 1))) {
                                                            case 32768:
                                                                if true {
                                                                    p0 = p0
                                                                    var v153 uint64 = 42
                                                                    if ((138 < 1) || false) {
                                                                        fmt.Println(fi18)
                                                                        var v154 int32 = int32(fi73)
                                                                        p0 = p0
                                                                        _ = v154
                                                                    } else {
                                                                        var v155 uint64 = (^(v153))
                                                                        _ = v155
                                                                    }
                                                                    _ = v153
                                                                } else {
                                                                    switch fi18 {
                                                                    case 626:
                                                                        p0 = p0
                                                                    case 9223372036854775807:
                                                                        if len(v17) > 0 {
                                                                            v17[0] = -2147483648
                                                                        }
                                                                    }
                                                                }
                                                            case 9223372036854775807:
                                                                v156 := GenS5{F0: (4294967295 == 4294967295), F1: (p0 & (p0 / (((p0 << ((4294967295) & 63))) | 1))), F2: "a"}
                                                                _ = v156
                                                            case 315:
                                                                switch fi18 {
                                                                case 835:
                                                                    var v157 int8 = -1
                                                                    _ = v157
                                                                case -816:
                                                                    if (!true) {
                                                                        v158 := GenS14{F0: 255}
                                                                        if false {
                                                                            var v159 int = fi73
                                                                            if len(v17) > 0 {
                                                                                v17[0] = 9223372036854775807
                                                                            }
                                                                            v160 := GenS13{F0: ""}
                                                                            _ = v159
                                                                            _ = v160
                                                                        }
                                                                        _ = v158
                                                                    }
                                                                }
                                                            }
                                                            v161 := GenS6{F0: ("a" == ("foo bar" + ("hello" + ("" + ""))))}
                                                            p0 = (p0 % (((p0 >> ((312) & 63))) | 1))
                                                            _ = v161
                                                        }
                                                    }
                                                    v162 := GenS15{F0: fi73, F1: 1}
                                                    _ = v162
                                                } else {
                                                    var v163 int = (^(fi73))
                                                    _ = v163
                                                }
                                                fmt.Println(p0)
                                                if true {
                                                    if false {
                                                        v164 := GenS7{F0: 42}
                                                        if (-1 <= int8(fi73)) {
                                                            p0 += p0
                                                            p0 = p0
                                                            var v165 uint16 = 2
                                                            _ = v165
                                                        }
                                                        v17 = append(v17, (fi18 >> ((813) & 63)))
                                                        v166 := GenS8{F0: 255, F1: 1}
                                                        _ = v164
                                                        _ = v166
                                                    }
                                                }
                                            } else {
                                                v167 := GenS16{F0: 2}
                                                p0 = p0
                                                v168 := GenS17{F0: 65535}
                                                if len(v17) >= 2 {
                                                    v17 = v17[1:len(v17)]
                                                }
                                                fmt.Println(fi73, p0, fi18)
                                                _ = v167
                                                _ = v168
                                            }
                                        }
                                    }
                                default:
                                    var v169 int32 = -837
                                    _ = v169
                                }
                            }
                            switch "" {
                            case "hello":
                                p0 -= p0
                            default:
                                v170 := GenS15{F0: fi18, F1: -1}
                                _ = v170
                            }
                            switch ("a" + (("\n\t" + (("foo bar" + "foo bar") + "a")) + "hello")) {
                            case "\n\t":
                                var v171 int32 = 42
                                _ = v171
                            default:
                                var v172 int16 = -32768
                                _ = v172
                            }
                            _ = fi73
                        }
                        var v173 uint64 = 507
                        _ = v173
                    }
                    if false {
                        p0 = p0
                        var v174 bool = true
                        v175 := GenS10{F0: -536, F1: v174, F2: "hello"}
                        _ = v174
                        _ = v175
                    } else {
                        for fi176 := 0; fi176 < 2; fi176++ {
                            var v177 int64 = 1
                            _ = len(v17)
                            switch (fi176 * fi18) {
                            case 9223372036854775807:
                                var v178 uint8 = 22
                                _ = v178
                            case 2147483647:
                                if true {
                                    _ = len(v17)
                                    v179 := GenS7{F0: 42}
                                    _ = v179
                                } else {
                                    v180 := GenS16{F0: 18446744073709551615}
                                    if false {
                                        var v181 float32 = (((-2.25 + float32(3.14159)) / (-2.25 / -2.25)) * (((((3.14159 / -2.25) + 0.0) * 0.1) - float32(float64(1.5))) / (0.1 - 3.14159)))
                                        p0 = -1
                                        _ = v181
                                    } else {
                                        if len(v17) > 0 {
                                            v17[0] = int(p0)
                                        }
                                    }
                                    if true {
                                        fmt.Println(p0)
                                    } else {
                                        switch 0 {
                                        case 0:
                                            var v182 float64 = -2.25
                                            var v183 float64 = (1e10 * (v182 / 3.14159))
                                            var v184 string = "foo bar"
                                            _ = v182
                                            _ = v183
                                            _ = v184
                                        case -1:
                                            fmt.Println(fi18, fi18)
                                        case 127:
                                            p0 -= 0
                                        }
                                    }
                                    _ = v180
                                }
                            }
                            p0 -= v177
                            _ = fi176
                            _ = v177
                        }
                        p0 = p0
                        for fi185 := 0; fi185 < 5; fi185++ {
                            var v186 uint8 = 126
                            _ = fi185
                            _ = v186
                        }
                    }
                } else {
                    var v187 uint = 575
                    if false {
                        v17 = append(v17, -9223372036854775808)
                    } else {
                        if true {
                            switch fi18 {
                            case -387:
                                v188 := GenS14{F0: 2}
                                _ = v188
                            default:
                                if (false || (4294967295 >= 4294967295)) {
                                    if ((!false) || true) {
                                        v189 := GenS18{F0: 480}
                                        _ = v189
                                    }
                                    var v190 string = ("foo bar" + "a")
                                    _ = v190
                                } else {
                                    v187 += 255
                                    _ = len(v17)
                                }
                            }
                            v191 := GenS19{F0: 0, F1: 42}
                            v192 := GenS20{F0: 0}
                            _ = v191
                            _ = v192
                        }
                        v193 := GenS13{F0: (("foo bar" + ("foo bar" + "\n\t")) + ("hello" + "\n\t"))}
                        var v194 int8 = 0
                        _ = v193
                        _ = v194
                    }
                    _ = v187
                }
            }
            v17 = append(v17, (-(fi18)))
            fmt.Println(p0, p0)
            _ = fi18
        }
    }
    _ = p0
    _ = v17
    return ((("\n\t" + "a") + ("foo bar" + "a")) + (("hello" + (("" + "foo bar") + (("" + "hello") + "hello"))) + ("a" + "a")))
}

func main() {
    v1 := []string{"a", "a", "foo bar", ("a" + ("\n\t" + "\n\t"))}
    var v2 bool = false
    v3 := []int{-1, 127, 0}
    var v4 int64 = 0
    v5 := make(map[uint32]string)
    var v6 int = -683
    for fi195 := 0; fi195 < 15; fi195++ {
        if ((("a" + "\n\t") + "\n\t") < "foo bar") {
            v3 = append(v3, (^(fi195)))
            switch 4294967295 {
            case 1:
                fmt.Println(v2, v6)
            case 42:
                var v196 int = fi195
                _ = v196
            case 114:
                if v2 {
                    if true {
                        var v197 uint8 = 255
                        _ = v197
                    } else {
                        if (("foo bar" + "foo bar") >= ("" + ("hello" + ("hello" + "\n\t")))) {
                            fmt.Println(v6)
                            fmt.Println(v6, v6, v2)
                        }
                        switch "\n\t" {
                        case "":
                            var v198 int16 = genFn0(18446744073709551615)
                            _ = v198
                        case "hello":
                            var v199 int = (^(fi195))
                            _ = v199
                        default:
                            v5[614] = ("\n\t" + "foo bar")
                        }
                        v6 = (v6 * -2147483648)
                    }
                    var v200 float64 = -2.25
                    v201 := GenS16{F0: 233}
                    _ = v200
                    _ = v201
                } else {
                    var v202 int16 = -32768
                    _ = v202
                }
            default:
                if v2 {
                    v2 = (492 <= (v4 / (((v4 + v4)) | 1)))
                    v4 = v4
                    _ = len(v5)
                } else {
                    fmt.Println(v4)
                    v203 := GenS19{F0: 1, F1: 1}
                    _ = v203
                }
            }
        } else {
            switch fi195 {
            case 9223372036854775807:
                for fi204 := 0; fi204 < 12; fi204++ {
                    if (-2.25 > 0.0) {
                        if (v2 || v2) {
                            var v205 float32 = (0.0 - (3.14159 / 1.5))
                            _ = v205
                        }
                        v6 = (-(fi195))
                        v206 := GenS3{F0: 4294967295}
                        _ = v206
                    }
                    _ = fi204
                }
            default:
                var v207 bool = (1 > 1)
                _ = v207
            }
        }
        fmt.Println(v2, fi195, fi195)
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        _ = fi195
    }
    switch ("" + "") {
    case "hello":
        v4 = (-(v4))
    case "":
        switch "a" {
        case "\n\t":
            v5[42] = "a"
        case "foo bar":
            _ = len(v3)
        default:
            for fi208 := 0; fi208 < 16; fi208++ {
                v209 := GenS15{F0: v6, F1: -671}
                switch (fi208 % (((v6 << ((255) & 63))) | 1)) {
                case -9223372036854775808:
                    var v210 uint64 = 1
                    _ = v210
                case 127:
                    v6 += 32768
                }
                for fi211 := 0; fi211 < 5; fi211++ {
                    v212 := GenS21{F0: -1}
                    var v213 uint = 65535
                    _ = fi211
                    _ = v212
                    _ = v213
                }
                _ = fi208
                _ = v209
            }
        }
    case "a":
        v214 := GenS22{F0: v6, F1: v4, F2: 42}
        _ = v214
    default:
        v215 := GenS14{F0: 338}
        _ = v215
    }
    v2 = true
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch (("\n\t" + "foo bar") + "\n\t") {
    case "a":
        var v216 uint32 = 4294967295
        _ = v216
    case "foo bar":
        if v2 {
            v1 = append(v1, "\n\t")
            v6 -= 42
            for fi217 := 0; fi217 < 6; fi217++ {
                if v2 {
                    for fi218 := 0; fi218 < 9; fi218++ {
                        if ((^(fi218)) >= -1) {
                            v6 = 32768
                            if (65535 >= 1) {
                                v4 -= v4
                            }
                            var v219 uint32 = 350
                            _ = v219
                        } else {
                            v6 -= fi218
                        }
                        var v220 int32 = -2147483648
                        _ = fi218
                        _ = v220
                    }
                    var v221 int64 = 0
                    _ = v221
                } else {
                    v222 := v5[4294967295]
                    if v2 {
                        for fi223 := 0; fi223 < 19; fi223++ {
                            v2 = (v2 || (9223372036854775807 == (v6 + 2147483647)))
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi223
                        }
                        var v224 int = fi217
                        _ = v224
                    } else {
                        _ = len(v5)
                    }
                    _ = v222
                }
                _ = fi217
            }
        }
    }
    if v2 {
        var v225 int32 = -1
        _ = v225
    } else {
        v2 = true
        v226 := GenS21{F0: 6}
        _ = v226
    }
    v2 = (2 < 2)
    switch "hello" {
    case "":
        v227 := GenS23{F0: 692}
        _ = v227
    }
    fmt.Println(v4, v4, v6)
    if v2 {
        delete(v5, 0)
        v1 = append(v1, (("" + "hello") + ("" + "a")))
    }
    delete(v5, 2147483648)
    for fi228 := 0; fi228 < 9; fi228++ {
        v4 += v4
        _ = fi228
    }
    _ = len(v5)
    if (genFn0(857) == -32768) {
        var v229 uint8 = 121
        _ = v229
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
