package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 bool; F1 int; F2 uint8 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 string; F1 uint8 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 string; F1 int }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 int16; F1 int16 }
type GenS16 struct { F0 uint16; F1 uint64 }
type GenS17 struct { F0 int64 }
type GenS18 struct { F0 bool; F1 int64 }
type GenS19 struct { F0 uint; F1 int }
type GenS20 struct { F0 int16 }
type GenS21 struct { F0 int8; F1 uint8 }
type GenS22 struct { F0 uint64; F1 int16 }
type GenS23 struct { F0 int8; F1 uint16 }
type GenS24 struct { F0 uint16; F1 uint; F2 int64 }
type GenS25 struct { F0 int64; F1 int64 }
type GenS26 struct { F0 int64; F1 int }
type GenS27 struct { F0 int; F1 int16 }
func genFn0(p0 int64, p1 int8, p2 int64) string {
    var v4 float32 = (3.14159 + 1e10)
    v5 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return ("" + "")
}
func genFn1(p0 uint16) uint8 {
    if false {
        if (-116 > -9223372036854775808) {
            v6 := []int{}
            _ = v6
        }
        fmt.Println(p0)
    }
    if (false && (32767 <= 891)) {
        var v7 int32 = 0
        if (!true) {
            v8 := []int{-2147483648}
            for fi9 := 0; fi9 < 16; fi9++ {
                v7 -= v7
                _ = len(v8)
                _ = fi9
            }
            _ = v8
        } else {
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 int64 = -787
                _ = fi10
                _ = v11
            }
            fmt.Println(v7)
            switch "hello" {
            case "foo bar":
                switch "\n\t" {
                case "hello":
                    v7 += (v7 & 603)
                case "a":
                    if (true || (v7 != v7)) {
                        p0 = p0
                        v12 := []int{42}
                        if true {
                            v12 = append(v12, 2147483647)
                        }
                        _ = v12
                    }
                case "":
                    v13 := GenS1{F0: 134}
                    _ = v13
                default:
                    v14 := GenS1{F0: 42}
                    _ = v14
                }
            default:
                switch "foo bar" {
                case "":
                    fmt.Println(p0)
                case "\n\t":
                    v15 := []int{}
                    _ = v15
                }
            }
        }
        for fi16 := 0; fi16 < 11; fi16++ {
            switch "foo bar" {
            case "foo bar":
                switch (fi16 << ((42) & 63)) {
                case 337:
                    v17 := GenS2{F0: (!true)}
                    _ = v17
                case 383:
                    if false {
                        if true {
                            switch int(v7) {
                            case 1:
                                for fi18 := 0; fi18 < 4; fi18++ {
                                    if (!true) {
                                        v19 := GenS3{F0: ((false || (true || false)) || false)}
                                        switch "foo bar" {
                                        case "":
                                            p0 = 824
                                        case "foo bar":
                                            fmt.Println(v7, p0)
                                        case "hello":
                                            v20 := GenS4{F0: -882}
                                            _ = v20
                                        default:
                                            if (!false) {
                                                v21 := []int{42, fi16}
                                                _ = v21
                                            } else {
                                                switch fi16 {
                                                case 2147483647:
                                                    var v22 int8 = -128
                                                    _ = v22
                                                case -319:
                                                    switch fi16 {
                                                    case 20:
                                                        switch (fi18 / (((fi18 * fi18)) | 1)) {
                                                        case 2147483647:
                                                            var v23 string = "hello"
                                                            _ = v23
                                                        case -9223372036854775808:
                                                            if false {
                                                                if true {
                                                                    p0 = 0
                                                                    v24 := []int{(fi16 - fi16), fi18}
                                                                    _ = v24
                                                                }
                                                            }
                                                            if false {
                                                                fmt.Println(v7, v7)
                                                            } else {
                                                                p0 = (p0 & p0)
                                                            }
                                                        case 32767:
                                                            if true {
                                                                if true {
                                                                    var v25 uint = 42
                                                                    var v26 int64 = 42
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        if true {
                                                                            if ((1 <= 0) && (false || (false && true))) {
                                                                                if (1 > int8(fi18)) {
                                                                                    var v27 uint8 = 0
                                                                                    v26 += (v26 | -989)
                                                                                    v25 = v25
                                                                                    _ = v27
                                                                                }
                                                                            }
                                                                            var v28 bool = true
                                                                            if (2147483647 > 0) {
                                                                                v7 += v7
                                                                                v29 := []int{fi16, fi18, (fi16 & -322)}
                                                                                _ = v29
                                                                            }
                                                                            _ = v28
                                                                        }
                                                                    }
                                                                    _ = v25
                                                                    _ = v26
                                                                }
                                                            }
                                                        }
                                                    }
                                                default:
                                                    if ((((9223372036854775807 > 42) || (!(("" + "\n\t") > "hello"))) && true) || (!true)) {
                                                        var v30 int32 = (v7 + (v7 / (((v7 ^ (v7 - v7))) | 1)))
                                                        var v31 uint = 42
                                                        switch fi16 {
                                                        case 9223372036854775807:
                                                            var v32 string = "a"
                                                            _ = v32
                                                        default:
                                                            v7 = 1
                                                            if false {
                                                                if (148 == 1) {
                                                                    v33 := []int{fi16, fi18, fi16, (^(fi18))}
                                                                    var v34 string = ("" + "hello")
                                                                    _ = v33
                                                                    _ = v34
                                                                } else {
                                                                    if (true && (true || (696 != 664))) {
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            if true {
                                                                                switch fi16 {
                                                                                case -9223372036854775808:
                                                                                    var v35 uint32 = 65535
                                                                                    _ = v35
                                                                                default:
                                                                                    switch (fi16 >> (((^(v31))) & 63)) {
                                                                                    case 4294967295:
                                                                                        var v36 int = fi16
                                                                                        _ = v36
                                                                                    case 9223372036854775807:
                                                                                        v37 := GenS2{F0: (!(false || (true && false)))}
                                                                                        _ = v37
                                                                                    default:
                                                                                        fmt.Println(v30, p0)
                                                                                    }
                                                                                }
                                                                                v30 = (v30 % ((v7) | 1))
                                                                            }
                                                                        case "hello":
                                                                            if false {
                                                                                var v38 uint16 = p0
                                                                                _ = v38
                                                                            }
                                                                        case "":
                                                                            switch ("" + ("foo bar" + (("hello" + "foo bar") + "hello"))) {
                                                                            case "\n\t":
                                                                                var v39 float32 = 1.5
                                                                                _ = v39
                                                                            default:
                                                                                v40 := GenS5{F0: fi18}
                                                                                _ = v40
                                                                            }
                                                                        }
                                                                        if (fi18 != fi16) {
                                                                            fmt.Println(v30, fi16)
                                                                            fmt.Println(fi16, p0, fi16)
                                                                            var v41 int8 = 0
                                                                            _ = v41
                                                                        } else {
                                                                            var v42 int16 = -32768
                                                                            p0 -= (p0 & 1)
                                                                            _ = v42
                                                                        }
                                                                    } else {
                                                                        var v43 float32 = 0.0
                                                                        var v44 int32 = -2147483648
                                                                        _ = v43
                                                                        _ = v44
                                                                    }
                                                                    if (1 < -269) {
                                                                        if (p0 <= p0) {
                                                                            v30 += v7
                                                                            v45 := []int{(fi16 * -6)}
                                                                            v46 := GenS3{F0: true}
                                                                            _ = v45
                                                                            _ = v46
                                                                        }
                                                                    } else {
                                                                        var v47 bool = false
                                                                        v48 := GenS6{F0: 42, F1: 172}
                                                                        var v49 float32 = 0.1
                                                                        _ = v47
                                                                        _ = v48
                                                                        _ = v49
                                                                    }
                                                                    v50 := []int{(^(fi16)), (fi18 / ((-2147483648) | 1))}
                                                                    var v51 uint64 = 334
                                                                    v52 := GenS7{F0: ("a" + "\n\t"), F1: (v31 != (v31 / (((v31 >> (((^(v31))) & 63))) | 1)))}
                                                                    _ = v50
                                                                    _ = v51
                                                                    _ = v52
                                                                }
                                                                v7 -= 355
                                                            }
                                                            var v53 float64 = ((3.14159 * -2.25) - (0.0 / 3.14159))
                                                            _ = v53
                                                        }
                                                        if ((!(4294967295 < 538)) || false) {
                                                            var v54 uint16 = (p0 - 42)
                                                            v55 := []int{fi18}
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                        fmt.Println(v7, v30)
                                                        v56 := GenS2{F0: true}
                                                        if (float64(3.14159) >= 0.0) {
                                                            switch "a" {
                                                            case "hello":
                                                                v7 += v30
                                                            }
                                                        } else {
                                                            v57 := GenS8{F0: 8}
                                                            if (32767 >= 127) {
                                                                var v58 bool = (!true)
                                                                if v58 {
                                                                    var v59 int8 = 42
                                                                    v31 -= (^(v31))
                                                                    v60 := []int{fi18}
                                                                    v61 := GenS9{F0: (v58 && v58), F1: (fi16 << ((18446744073709551615) & 63)), F2: 0}
                                                                    _ = v59
                                                                    _ = v60
                                                                    _ = v61
                                                                }
                                                                v62 := []int{}
                                                                fmt.Println(v30, fi16)
                                                                var v63 uint16 = (^(p0))
                                                                switch fi18 {
                                                                case -506:
                                                                    var v64 int = 915
                                                                    _ = v64
                                                                case 128:
                                                                    v62 = append(v62, (fi18 | (fi16 ^ fi18)))
                                                                default:
                                                                    switch (-(fi16)) {
                                                                    case 0:
                                                                        var v65 int64 = -403
                                                                        _ = v65
                                                                    }
                                                                }
                                                                v66 := GenS4{F0: 9223372036854775807}
                                                                v67 := GenS10{F0: v30}
                                                                if true {
                                                                    v68 := GenS8{F0: 252}
                                                                    var v69 int64 = 1
                                                                    if true {
                                                                        v62 = append(v62, fi18)
                                                                        v70 := GenS9{F0: (0 == 704), F1: (fi18 + fi16), F2: 0}
                                                                        switch ("a" + "") {
                                                                        case "\n\t":
                                                                            var v71 int = fi18
                                                                            _ = v71
                                                                        case "":
                                                                            v30 = v7
                                                                        case "a":
                                                                            v63 += v63
                                                                        default:
                                                                            v63 -= 537
                                                                        }
                                                                        _ = v70
                                                                    } else {
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            v72 := GenS11{F0: (("\n\t" + "a") + "\n\t"), F1: 255}
                                                                            _ = v72
                                                                        case "\n\t":
                                                                            v63 -= (p0 | (v63 % ((p0) | 1)))
                                                                        case "hello":
                                                                            if v58 {
                                                                                var v73 int = fi18
                                                                                if false {
                                                                                    if v58 {
                                                                                        _ = len(v62)
                                                                                    } else {
                                                                                        if (v69 >= v69) {
                                                                                            if len(v62) >= 2 {
                                                                                                v62 = v62[1:len(v62)]
                                                                                            }
                                                                                            if (!true) {
                                                                                                if false {
                                                                                                    _ = len(v62)
                                                                                                    if ((!(0 == -436)) && v58) {
                                                                                                        v31 = (^(v31))
                                                                                                        var v74 uint = (v31 / ((18446744073709551615) | 1))
                                                                                                        p0 = (^(v63))
                                                                                                        _ = v74
                                                                                                    }
                                                                                                    var v75 int8 = 1
                                                                                                    _ = v75
                                                                                                } else {
                                                                                                    switch fi18 {
                                                                                                    case -2147483648:
                                                                                                        v7 -= -1
                                                                                                    case -286:
                                                                                                        v76 := GenS7{F0: ("foo bar" + ""), F1: v58}
                                                                                                        _ = v76
                                                                                                    case 9223372036854775807:
                                                                                                        v77 := GenS8{F0: 255}
                                                                                                        _ = v77
                                                                                                    }
                                                                                                    v30 = 42
                                                                                                    v58 = v58
                                                                                                }
                                                                                                if len(v62) >= 2 {
                                                                                                    v62 = v62[1:len(v62)]
                                                                                                }
                                                                                            }
                                                                                            if v58 {
                                                                                                switch ((("a" + "hello") + ("\n\t" + ("a" + "hello"))) + "") {
                                                                                                case "foo bar":
                                                                                                    v58 = false
                                                                                                }
                                                                                                fmt.Println(fi18, fi18, fi16)
                                                                                            }
                                                                                            fmt.Println(v73, v63)
                                                                                            if len(v62) > 0 {
                                                                                                v62[0] = (v73 | (-(fi16)))
                                                                                            }
                                                                                            if true {
                                                                                                _ = len(v62)
                                                                                            } else {
                                                                                                var v78 bool = false
                                                                                                if len(v62) > 0 {
                                                                                                    v62[0] = v73
                                                                                                }
                                                                                                _ = v78
                                                                                            }
                                                                                        } else {
                                                                                            if len(v62) > 0 {
                                                                                                v62[0] = fi16
                                                                                            }
                                                                                            switch v73 {
                                                                                            case 32768:
                                                                                                v31 = (v31 | v31)
                                                                                            default:
                                                                                                if true {
                                                                                                    if v58 {
                                                                                                        _ = len(v62)
                                                                                                    }
                                                                                                    v73 -= fi18
                                                                                                    v73 = 2147483647
                                                                                                }
                                                                                            }
                                                                                            var v79 uint32 = 2147483648
                                                                                            _ = v79
                                                                                        }
                                                                                        v80 := GenS4{F0: (-(v69))}
                                                                                        v81 := GenS12{F0: 106}
                                                                                        if ((^(v31)) >= v31) {
                                                                                            var v82 bool = false
                                                                                            switch fi18 {
                                                                                            case -430:
                                                                                                var v83 int8 = 62
                                                                                                _ = v83
                                                                                            case -2147483648:
                                                                                                switch "foo bar" {
                                                                                                case "hello":
                                                                                                    v7 += (v7 | v7)
                                                                                                case "\n\t":
                                                                                                    v62 = append(v62, (fi16 / ((v73) | 1)))
                                                                                                case "a":
                                                                                                    var v84 uint = v31
                                                                                                    _ = v84
                                                                                                default:
                                                                                                    fmt.Println(v58)
                                                                                                }
                                                                                            case 32767:
                                                                                                var v85 bool = (float32((1e10 * 1.5)) != 1e10)
                                                                                                _ = v85
                                                                                            default:
                                                                                                if (42 == 0) {
                                                                                                    if v82 {
                                                                                                        var v86 bool = (!false)
                                                                                                        _ = v86
                                                                                                    } else {
                                                                                                        var v87 uint32 = 0
                                                                                                        if false {
                                                                                                            v88 := GenS11{F0: "", F1: 212}
                                                                                                            var v89 uint8 = 128
                                                                                                            v90 := GenS13{F0: "a", F1: fi16}
                                                                                                            _ = v88
                                                                                                            _ = v89
                                                                                                            _ = v90
                                                                                                        } else {
                                                                                                            var v91 int = (v73 << ((uint(v87)) & 63))
                                                                                                            _ = v91
                                                                                                        }
                                                                                                        _ = v87
                                                                                                    }
                                                                                                    v7 += (v30 * (v30 * 2147483647))
                                                                                                    var v92 int8 = 127
                                                                                                    _ = v92
                                                                                                } else {
                                                                                                    v93 := GenS1{F0: 128}
                                                                                                    v94 := GenS9{F0: true, F1: (fi18 | v73), F2: 92}
                                                                                                    if false {
                                                                                                        switch "a" {
                                                                                                        case "a":
                                                                                                            v62 = append(v62, (v73 ^ 2147483647))
                                                                                                        }
                                                                                                        v7 += (v30 ^ -2147483648)
                                                                                                    } else {
                                                                                                        var v95 bool = true
                                                                                                        var v96 int32 = (v30 & (v7 % ((-417) | 1)))
                                                                                                        var v97 int8 = 1
                                                                                                        _ = v95
                                                                                                        _ = v96
                                                                                                        _ = v97
                                                                                                    }
                                                                                                    _ = v93
                                                                                                    _ = v94
                                                                                                }
                                                                                            }
                                                                                            _ = v82
                                                                                        } else {
                                                                                            _ = len(v62)
                                                                                            v98 := GenS14{F0: 2147483647}
                                                                                            _ = v98
                                                                                        }
                                                                                        _ = v80
                                                                                        _ = v81
                                                                                    }
                                                                                    var v99 bool = (2147483648 < 4294967295)
                                                                                    var v100 float32 = float32(3.14159)
                                                                                    _ = v100
                                                                                    _ = v99
                                                                                } else {
                                                                                    fmt.Println(v73, v31, v73)
                                                                                }
                                                                                v101 := GenS5{F0: (v73 & fi18)}
                                                                                var v102 uint32 = 65535
                                                                                _ = v101
                                                                                _ = v102
                                                                                _ = v73
                                                                            } else {
                                                                                var v103 string = (((("foo bar" + "\n\t") + ("" + "hello")) + ("foo bar" + "a")) + ("" + ("foo bar" + "")))
                                                                                fmt.Println(v69, v30, v103)
                                                                                v69 += (v69 << ((25) & 63))
                                                                                _ = v103
                                                                            }
                                                                        }
                                                                        v62 = append(v62, fi16)
                                                                        v30 -= (v7 % ((v7) | 1))
                                                                    }
                                                                    v69 = -981
                                                                    _ = v68
                                                                    _ = v69
                                                                } else {
                                                                    v31 -= v31
                                                                    if v58 {
                                                                        fmt.Println(v58, fi18, fi16)
                                                                        var v104 uint64 = 0
                                                                        _ = v104
                                                                    }
                                                                }
                                                                _ = v58
                                                                _ = v62
                                                                _ = v63
                                                                _ = v66
                                                                _ = v67
                                                            } else {
                                                                v105 := []int{fi16, 4294967295}
                                                                _ = v105
                                                            }
                                                            _ = v57
                                                        }
                                                        v106 := GenS2{F0: ((fi18 >> ((30) & 63)) > fi18)}
                                                        _ = v106
                                                        _ = v30
                                                        _ = v31
                                                        _ = v56
                                                    }
                                                }
                                                if true {
                                                    v107 := GenS15{F0: 32767, F1: 1000}
                                                    _ = v107
                                                }
                                                p0 = 0
                                                var v108 string = ("foo bar" + "\n\t")
                                                _ = v108
                                            }
                                        }
                                        _ = v19
                                    }
                                    if (!(!(p0 == 1))) {
                                        if (!((128 <= p0) && true)) {
                                            if (v7 != -2147483648) {
                                                v109 := []int{0, 32768, (fi16 >> ((18446744073709551615) & 63))}
                                                v7 = v7
                                                v110 := GenS11{F0: "", F1: 1}
                                                switch fi16 {
                                                case 2147483647:
                                                    var v111 uint = 494
                                                    var v112 bool = (!((!false) && (true && false)))
                                                    _ = v111
                                                    _ = v112
                                                case 128:
                                                    if (!true) {
                                                        if len(v109) >= 2 {
                                                            v109 = v109[1:len(v109)]
                                                        }
                                                    } else {
                                                        if len(v109) >= 2 {
                                                            v109 = v109[1:len(v109)]
                                                        }
                                                    }
                                                case 0:
                                                    var v113 int64 = -1
                                                    _ = v113
                                                default:
                                                    switch (fi18 ^ (^(fi16))) {
                                                    case 131:
                                                        if true {
                                                            if false {
                                                                v7 = int32(p0)
                                                                switch "" {
                                                                case "foo bar":
                                                                    v114 := GenS4{F0: 1}
                                                                    p0 -= 65535
                                                                    _ = v114
                                                                case "\n\t":
                                                                    v115 := GenS12{F0: 2147483648}
                                                                    _ = v115
                                                                default:
                                                                    var v116 int16 = -32768
                                                                    _ = v116
                                                                }
                                                                switch (("a" + (("hello" + "\n\t") + "a")) + (("foo bar" + ("a" + "hello")) + (("hello" + "foo bar") + "foo bar"))) {
                                                                case "foo bar":
                                                                    if ((int64(fi16) > int64(fi18)) && true) {
                                                                        v7 = -1
                                                                        _ = len(v109)
                                                                        if ((1 != 310) || (2 <= 565)) {
                                                                            var v117 uint64 = 121
                                                                            v117 += v117
                                                                            v7 -= v7
                                                                            _ = v117
                                                                        }
                                                                    }
                                                                case "":
                                                                    var v118 int16 = 42
                                                                    _ = v118
                                                                case "\n\t":
                                                                    var v119 int64 = 9223372036854775807
                                                                    _ = v119
                                                                default:
                                                                    p0 = p0
                                                                }
                                                                fmt.Println(fi18)
                                                                v120 := GenS1{F0: 128}
                                                                fmt.Println(fi18, v7, fi16)
                                                                v7 += int32(p0)
                                                                _ = v120
                                                            } else {
                                                                fmt.Println(fi16, fi18, v7)
                                                                if true {
                                                                    switch (("hello" + ("a" + ("hello" + "\n\t"))) + "") {
                                                                    case "hello":
                                                                        if (true && false) {
                                                                            if (true && true) {
                                                                                v7 -= v7
                                                                                var v121 bool = (!(42 >= 769))
                                                                                if v121 {
                                                                                    if (42 <= 1) {
                                                                                        if len(v109) > 0 {
                                                                                            v109[0] = int(p0)
                                                                                        }
                                                                                    } else {
                                                                                        switch (fi18 & (fi16 % (((fi16 + 128)) | 1))) {
                                                                                        case 426:
                                                                                            var v122 uint64 = 1
                                                                                            _ = v122
                                                                                        case 207:
                                                                                            if len(v109) >= 2 {
                                                                                                v109 = v109[1:len(v109)]
                                                                                            }
                                                                                        case -631:
                                                                                            fmt.Println(p0, v7)
                                                                                        default:
                                                                                            switch "foo bar" {
                                                                                            case "hello":
                                                                                                var v123 string = "\n\t"
                                                                                                _ = v123
                                                                                            }
                                                                                            v109 = append(v109, fi18)
                                                                                            if true {
                                                                                                if len(v109) > 0 {
                                                                                                    v109[0] = fi16
                                                                                                }
                                                                                            } else {
                                                                                                v124 := GenS7{F0: (("hello" + "\n\t") + (("\n\t" + ("foo bar" + "a")) + "foo bar")), F1: ((^(fi18)) < (fi18 ^ fi16))}
                                                                                                v125 := GenS10{F0: (v7 - (v7 + (v7 ^ 180)))}
                                                                                                _ = v124
                                                                                                _ = v125
                                                                                            }
                                                                                        }
                                                                                        v7 += (^(v7))
                                                                                        v121 = true
                                                                                        if false {
                                                                                            v7 += v7
                                                                                            if false {
                                                                                                if (544 > 4294967295) {
                                                                                                    switch -377 {
                                                                                                    case 789:
                                                                                                        var v126 float32 = 3.14159
                                                                                                        _ = v126
                                                                                                    case -890:
                                                                                                        var v127 uint16 = 956
                                                                                                        _ = v127
                                                                                                    default:
                                                                                                        var v128 int16 = 314
                                                                                                        _ = v128
                                                                                                    }
                                                                                                } else {
                                                                                                    switch ("foo bar" + "hello") {
                                                                                                    case "foo bar":
                                                                                                        p0 -= p0
                                                                                                    case "":
                                                                                                        v7 = (v7 - (v7 | v7))
                                                                                                    default:
                                                                                                        v7 = -849
                                                                                                        if len(v109) > 0 {
                                                                                                            v109[0] = (^(fi16))
                                                                                                        }
                                                                                                        v121 = (36 >= -9223372036854775808)
                                                                                                    }
                                                                                                    fmt.Println(v7, v121)
                                                                                                }
                                                                                                var v129 uint = 4294967295
                                                                                                if len(v109) > 0 {
                                                                                                    v109[0] = fi16
                                                                                                }
                                                                                                _ = v129
                                                                                            } else {
                                                                                                var v130 uint32 = 2147483647
                                                                                                v109 = append(v109, fi18)
                                                                                                var v131 int16 = 1
                                                                                                _ = v130
                                                                                                _ = v131
                                                                                            }
                                                                                            var v132 int64 = -1
                                                                                            if v121 {
                                                                                                v133 := GenS16{F0: (^(p0)), F1: 42}
                                                                                                if false {
                                                                                                    if (!(((-110 >= 1) && (411 > 18446744073709551615)) || ((!(4294967295 <= 42)) && v121))) {
                                                                                                        p0 = 255
                                                                                                    } else {
                                                                                                        v134 := GenS17{F0: v132}
                                                                                                        fmt.Println(v121, v132, v132)
                                                                                                        _ = v134
                                                                                                    }
                                                                                                    fmt.Println(fi16)
                                                                                                    v132 = (v132 << ((255) & 63))
                                                                                                    v109 = append(v109, (fi16 ^ (^(fi18))))
                                                                                                    var v135 uint8 = 1
                                                                                                    _ = v135
                                                                                                } else {
                                                                                                    var v136 float64 = -2.25
                                                                                                    v7 -= 20
                                                                                                    fmt.Println(fi16, v7, p0)
                                                                                                    _ = v136
                                                                                                }
                                                                                                _ = v133
                                                                                            }
                                                                                            _ = v132
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    var v137 uint8 = 164
                                                                                    _ = v137
                                                                                }
                                                                                _ = v121
                                                                            } else {
                                                                                fmt.Println(p0, v7)
                                                                                if (true && (1 < 561)) {
                                                                                    p0 = p0
                                                                                    v138 := GenS13{F0: "hello", F1: fi18}
                                                                                    switch ("a" + ("foo bar" + ("" + (("\n\t" + (("\n\t" + "foo bar") + "foo bar")) + "\n\t")))) {
                                                                                    case "":
                                                                                        if (-128 != 100) {
                                                                                            if len(v109) >= 2 {
                                                                                                v109 = v109[1:len(v109)]
                                                                                            }
                                                                                        } else {
                                                                                            if len(v109) > 0 {
                                                                                                v109[0] = -226
                                                                                            }
                                                                                        }
                                                                                    case "foo bar":
                                                                                        if (false && false) {
                                                                                            v139 := GenS16{F0: 1, F1: 82}
                                                                                            v140 := GenS2{F0: true}
                                                                                            _ = v139
                                                                                            _ = v140
                                                                                        }
                                                                                    case "a":
                                                                                        var v141 uint8 = 255
                                                                                        v142 := GenS2{F0: ((false || (true && ((fi16 % ((fi16) | 1)) == (-(fi16))))) && false)}
                                                                                        _ = v141
                                                                                        _ = v142
                                                                                    default:
                                                                                        v7 = (v7 << ((4294967295) & 63))
                                                                                        v7 += v7
                                                                                        fmt.Println(p0, fi16)
                                                                                        var v143 float64 = -2.25
                                                                                        _ = v143
                                                                                    }
                                                                                    _ = v138
                                                                                } else {
                                                                                    if true {
                                                                                        v144 := GenS13{F0: "\n\t", F1: -234}
                                                                                        _ = v144
                                                                                    }
                                                                                }
                                                                            }
                                                                            if len(v109) >= 2 {
                                                                                v109 = v109[1:len(v109)]
                                                                            }
                                                                            p0 -= (p0 % (((p0 + (p0 << ((1) & 63)))) | 1))
                                                                        } else {
                                                                            v7 -= (v7 ^ (v7 & v7))
                                                                            if len(v109) > 0 {
                                                                                v109[0] = (^(fi16))
                                                                            }
                                                                            p0 = p0
                                                                        }
                                                                    default:
                                                                        if false {
                                                                            var v145 uint32 = 792
                                                                            _ = v145
                                                                        } else {
                                                                            switch fi18 {
                                                                            case 252:
                                                                                p0 = (^(p0))
                                                                            case 9223372036854775807:
                                                                                var v146 uint64 = 18446744073709551615
                                                                                _ = v146
                                                                            default:
                                                                                var v147 uint8 = 71
                                                                                _ = v147
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                var v148 uint32 = 65535
                                                                _ = v148
                                                            }
                                                            p0 -= (^(p0))
                                                            if true {
                                                                switch -17 {
                                                                case 9223372036854775807:
                                                                    var v149 string = ("a" + "hello")
                                                                    _ = v149
                                                                case -1:
                                                                    v109 = append(v109, fi16)
                                                                case -2147483648:
                                                                    v150 := GenS18{F0: false, F1: -9223372036854775808}
                                                                    _ = v150
                                                                default:
                                                                    v151 := GenS13{F0: ("" + "a"), F1: fi18}
                                                                    _ = v151
                                                                }
                                                                if len(v109) >= 2 {
                                                                    v109 = v109[1:len(v109)]
                                                                }
                                                            }
                                                            var v152 uint64 = 2
                                                            fmt.Println(p0, fi18)
                                                            _ = v152
                                                        } else {
                                                            var v153 float64 = (3.14159 + 3.14159)
                                                            _ = v153
                                                        }
                                                    case 9223372036854775807:
                                                        var v154 int8 = 0
                                                        _ = v154
                                                    }
                                                }
                                                fmt.Println(p0, p0)
                                                _ = v109
                                                _ = v110
                                            } else {
                                                fmt.Println(v7)
                                                v155 := []int{fi18, 4294967295, fi18}
                                                _ = v155
                                            }
                                            v156 := []int{}
                                            var v157 uint = 0
                                            _ = v156
                                            _ = v157
                                        } else {
                                            v7 += v7
                                        }
                                        v158 := []int{(fi16 ^ (fi16 ^ -668)), fi18, 2147483647, (fi16 >> ((42) & 63))}
                                        _ = v158
                                    } else {
                                        switch -878 {
                                        case 32768:
                                            var v159 int8 = 0
                                            _ = v159
                                        case -1:
                                            v160 := []int{(fi18 % (((-(fi16))) | 1))}
                                            _ = v160
                                        case -143:
                                            fmt.Println(fi16)
                                        }
                                        v7 -= (v7 % (((v7 >> ((65535) & 63))) | 1))
                                        var v161 float32 = -2.25
                                        _ = v161
                                    }
                                    v162 := GenS17{F0: 42}
                                    _ = fi18
                                    _ = v162
                                }
                            default:
                                var v163 bool = true
                                _ = v163
                            }
                        }
                        v7 -= 2147483647
                        p0 = 1
                    }
                case 9223372036854775807:
                    for fi164 := 0; fi164 < 12; fi164++ {
                        v7 = -2147483648
                        v165 := []int{(fi16 << ((2) & 63)), (fi16 - fi16), fi16, fi164}
                        if false {
                            v165 = append(v165, (fi16 & (fi16 >> ((0) & 63))))
                            if (false || true) {
                                if len(v165) > 0 {
                                    v165[0] = (fi164 % ((9223372036854775807) | 1))
                                }
                                switch ("a" + ("\n\t" + "foo bar")) {
                                case "a":
                                    fmt.Println(p0, fi164, fi164)
                                case "":
                                    switch "foo bar" {
                                    case "foo bar":
                                        v166 := GenS2{F0: (42 <= -99)}
                                        _ = v166
                                    case "\n\t":
                                        v167 := GenS1{F0: 128}
                                        _ = v167
                                    }
                                default:
                                    var v168 uint64 = 18446744073709551615
                                    _ = v168
                                }
                                if (710 <= p0) {
                                    var v169 uint64 = 0
                                    v7 += 2147483647
                                    _ = v169
                                }
                            } else {
                                _ = len(v165)
                            }
                            var v170 uint32 = 962
                            if (v170 == (^(v170))) {
                                v171 := GenS19{F0: 521, F1: 0}
                                v170 -= 492
                                _ = v171
                            } else {
                                var v172 string = "foo bar"
                                _ = v172
                            }
                            _ = v170
                        }
                        _ = fi164
                        _ = v165
                    }
                }
            }
            _ = fi16
        }
        _ = v7
    } else {
        var v173 int = 32767
        v174 := GenS3{F0: (-64 == -2147483648)}
        var v175 bool = false
        _ = v173
        _ = v174
        _ = v175
    }
    v176 := GenS20{F0: 0}
    fmt.Println(p0)
    _ = p0
    _ = v176
    return 202
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{}
    v3 := make(map[string]string)
    if (-55 != 0) {
        var v177 int32 = -2147483648
        _ = v177
    } else {
        _ = len(v1)
        v178 := GenS21{F0: 42, F1: 96}
        var v179 uint64 = 0
        _ = v178
        _ = v179
    }
    if false {
        var v180 int64 = 42
        _ = v180
    } else {
        delete(v1, 4294967295)
        var v181 int32 = 2147483647
        v182 := GenS2{F0: true}
        _ = v181
        _ = v182
    }
    var v183 uint = 4294967295
    if (((v183 & v183) <= v183) || (-1 < 78)) {
        var v184 uint64 = 627
        v184 -= (v184 & v184)
        v185 := GenS2{F0: (-22 < -128)}
        _ = v184
        _ = v185
    }
    for fi186 := 0; fi186 < 4; fi186++ {
        switch fi186 {
        case -1:
            for fi187 := 0; fi187 < 15; fi187++ {
                v183 = (^(v183))
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi187
            }
        case 42:
            v183 -= v183
        case 342:
            v188 := GenS21{F0: -31, F1: 193}
            _ = v188
        }
        _ = fi186
    }
    fmt.Println(v183)
    if (false || (!true)) {
        v1[0] = ("" + "a")
        if len(v2) > 0 {
            v2[0] = ("\n\t" + ("" + ("a" + "foo bar")))
        }
        var v189 uint16 = 32767
        _ = v189
    } else {
        v183 = v183
        if false {
            v183 = v183
        } else {
            if false {
                for fi190 := 0; fi190 < 17; fi190++ {
                    fmt.Println(v183, v183)
                    _ = fi190
                }
                v191 := v1[32767]
                v183 = 404
                _ = v191
            } else {
                switch 127 {
                case 127:
                    switch 2147483647 {
                    case -9223372036854775808:
                        if false {
                            v192 := GenS22{F0: 783, F1: 940}
                            v193 := GenS23{F0: -1, F1: 32767}
                            switch "" {
                            case "\n\t":
                                if true {
                                    if (44 == 195) {
                                        if true {
                                            fmt.Println(v183)
                                        } else {
                                            v183 = (v183 >> (((v183 | v183)) & 63))
                                        }
                                        v183 = (v183 >> ((v183) & 63))
                                        v3[("foo bar" + "\n\t")] = ("a" + (("foo bar" + (("\n\t" + ("hello" + ("hello" + "hello"))) + ("\n\t" + "foo bar"))) + "a"))
                                    }
                                    v194 := GenS24{F0: 0, F1: v183, F2: -1}
                                    switch "\n\t" {
                                    case "a":
                                        var v195 uint32 = 254
                                        _ = v195
                                    case "hello":
                                        switch -1 {
                                        case 958:
                                            v183 -= (v183 * (^(v183)))
                                        case -175:
                                            if true {
                                                fmt.Println(v183)
                                                var v196 int = -9223372036854775808
                                                _ = v196
                                            }
                                        case 9223372036854775807:
                                            for fi197 := 0; fi197 < 6; fi197++ {
                                                var v198 int64 = -141
                                                _ = fi197
                                                _ = v198
                                            }
                                        }
                                    }
                                    _ = v194
                                } else {
                                    switch 32767 {
                                    case 143:
                                        v2 = append(v2, "hello")
                                    }
                                    switch 349 {
                                    case 127:
                                        delete(v1, 725)
                                    case -906:
                                        if true {
                                            v199 := v1[-280]
                                            _ = v199
                                        }
                                    }
                                    v200 := GenS20{F0: 32767}
                                    _ = v200
                                }
                            case "hello":
                                v201 := GenS3{F0: false}
                                _ = v201
                            default:
                                if ((v183 * v183) != v183) {
                                    v183 = v183
                                    v183 -= 906
                                    var v202 string = ""
                                    _ = v202
                                } else {
                                    v203 := GenS2{F0: false}
                                    var v204 int8 = -1
                                    switch ("\n\t" + "foo bar") {
                                    case "\n\t":
                                        for fi205 := 0; fi205 < 6; fi205++ {
                                            var v206 int64 = 0
                                            if (fi205 <= 128) {
                                                if false {
                                                    v183 = 2
                                                    v207 := GenS6{F0: 0, F1: 283}
                                                    var v208 bool = (fi205 > -483)
                                                    _ = v207
                                                    _ = v208
                                                }
                                                fmt.Println(v183)
                                            }
                                            _ = fi205
                                            _ = v206
                                        }
                                    }
                                    _ = v203
                                    _ = v204
                                }
                            }
                            _ = v192
                            _ = v193
                        } else {
                            v209 := GenS25{F0: 1, F1: 1}
                            v210 := GenS26{F0: 0, F1: -1}
                            _ = v209
                            _ = v210
                        }
                    default:
                        if (0 == 965) {
                            v211 := GenS27{F0: 4294967295, F1: 42}
                            for fi212 := 0; fi212 < 18; fi212++ {
                                v183 -= v183
                                _ = fi212
                            }
                            _ = v211
                        } else {
                            for fi213 := 0; fi213 < 8; fi213++ {
                                v183 = v183
                                _ = fi213
                            }
                            var v214 bool = ((!true) && (210 != 32767))
                            _ = v214
                        }
                    }
                case -2147483648:
                    v215 := v1[-539]
                    _ = v215
                case 2147483647:
                    for fi216 := 0; fi216 < 8; fi216++ {
                        fmt.Println(fi216)
                        v183 += v183
                        _ = fi216
                    }
                }
            }
            v183 = (^(v183))
        }
    }
    var v217 float32 = 1.5
    var v218 int16 = 42
    v219 := GenS10{F0: -1}
    if false {
        fmt.Println(v218)
        fmt.Println(v218)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v183:", v183)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v217:", v217)
    fmt.Println("v218:", v218)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v183
    _ = v2
    _ = v217
    _ = v218
    _ = v219
    _ = v3
}
