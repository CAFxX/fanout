package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int16; F2 int32 }
type GenS3 struct { F0 int; F1 int; F2 uint }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 int; F1 int64; F2 uint }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 string; F1 int; F2 uint64 }
type GenS10 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 bool }
func genFn0(p0 bool, p1 string) int32 {
    for fi6 := 0; fi6 < 2; fi6++ {
        p0 = true
        switch p1 {
        case "hello":
            var v7 bool = p0
            _ = v7
        }
        p0 = ((32768 < (fi6 % (((fi6 >> ((18446744073709551615) & 63))) | 1))) && false)
        _ = fi6
    }
    switch 923 {
    case -356:
        var v8 uint64 = 528
        _ = v8
    case 32768:
        v9 := GenS1{F0: 1}
        _ = v9
    case -508:
        switch -694 {
        case 32767:
            fmt.Println(p1, p0)
        case 42:
            switch ((p1 + "a") + "\n\t") {
            case "foo bar":
                v10 := []int{32768, 128, 128}
                _ = v10
            case "a":
                var v11 int = 1
                _ = v11
            case "hello":
                if p0 {
                    var v12 uint64 = 42
                    _ = v12
                } else {
                    v13 := []int{42, 122, 280}
                    _ = v13
                }
            }
        }
    default:
        v14 := []int{-1}
        _ = v14
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 bool = (100 < 100)
            v18 := []int{fi16, fi15, 127, (fi16 & 9223372036854775807)}
            fmt.Println(p0)
            _ = fi16
            _ = v17
            _ = v18
        }
        fmt.Println(p0, fi15, fi15)
        _ = fi15
    }
    v19 := GenS2{F0: 1, F1: 1, F2: 1}
    var v20 int32 = 2147483647
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return v20
}
func genFn1(p0 uint16, p1 uint16, p2 uint16) uint32 {
    if (1 == 1) {
        v21 := GenS1{F0: 0}
        v22 := GenS3{F0: 0, F1: 1, F2: 12}
        switch "" {
        case "\n\t":
            p2 = 32767
        default:
            var v23 float32 = 3.14159
            _ = v23
        }
        _ = v21
        _ = v22
    } else {
        for fi24 := 0; fi24 < 4; fi24++ {
            var v25 uint16 = 42
            if false {
                if true {
                    v26 := []int{}
                    var v27 int8 = 0
                    v28 := GenS1{F0: -18}
                    _ = v26
                    _ = v27
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 19; fi29++ {
                        if false {
                            fmt.Println(p2)
                            if (255 < 145) {
                                var v30 int = (-(fi29))
                                var v31 uint8 = 255
                                v25 -= 65535
                                _ = v30
                                _ = v31
                            } else {
                                var v32 float64 = (float64(1e10) / (3.14159 + -2.25))
                                if ("hello" <= "\n\t") {
                                    if (!(!false)) {
                                        fmt.Println(p2, p0, v25)
                                        v33 := []int{fi29, fi24}
                                        _ = v33
                                    } else {
                                        v32 += (v32 / float64(0.1))
                                        var v34 int64 = -746
                                        var v35 string = ""
                                        if false {
                                            if true {
                                                var v36 uint = 213
                                                v32 += float64(0.1)
                                                _ = v36
                                            }
                                        } else {
                                            v37 := GenS4{F0: -1, F1: (!(false || (42 > genFn0((!(false && (785 != 0))), v35))))}
                                            v38 := []int{128, fi29, fi24}
                                            if len(v38) > 0 {
                                                v38[0] = fi24
                                            }
                                            _ = v37
                                            _ = v38
                                        }
                                        _ = v34
                                        _ = v35
                                    }
                                    var v39 uint16 = (p2 * v25)
                                    v40 := []int{fi24, 9223372036854775807}
                                    _ = v39
                                    _ = v40
                                }
                                var v41 int16 = -368
                                _ = v32
                                _ = v41
                            }
                            switch "" {
                            case "foo bar":
                                var v42 bool = (-1 <= 648)
                                _ = v42
                            case "hello":
                                p2 = 42
                            case "\n\t":
                                if false {
                                    p0 = (p2 | (v25 >> ((65535) & 63)))
                                    p0 = (p1 << ((65535) & 63))
                                    v43 := GenS5{F0: "hello", F1: false}
                                    _ = v43
                                }
                            }
                        } else {
                            v44 := GenS2{F0: 2, F1: 1000, F2: 0}
                            v25 = v25
                            _ = v44
                        }
                        _ = fi29
                    }
                    p0 -= p2
                }
                var v45 int8 = -24
                _ = v45
            } else {
                fmt.Println(p1, fi24)
                v46 := []int{(fi24 % ((fi24) | 1)), (fi24 + fi24)}
                if len(v46) >= 2 {
                    v46 = v46[1:len(v46)]
                }
                _ = v46
            }
            v47 := []int{(fi24 & 0)}
            _ = fi24
            _ = v25
            _ = v47
        }
        switch 976 {
        case 32768:
            v48 := []int{-769, 2147483647}
            _ = v48
        case -9223372036854775808:
            v49 := GenS4{F0: -924, F1: (45 == 1)}
            _ = v49
        case 165:
            var v50 int16 = 0
            _ = v50
        }
    }
    v51 := GenS6{F0: 127, F1: 2147483647}
    for fi52 := 0; fi52 < 6; fi52++ {
        for fi53 := 0; fi53 < 17; fi53++ {
            v54 := []int{(-(fi52)), fi52, (fi53 << ((241) & 63)), (fi52 / (((fi53 >> ((4294967295) & 63))) | 1))}
            fmt.Println(p1, p1)
            _ = fi53
            _ = v54
        }
        v55 := GenS7{F0: (fi52 & fi52), F1: -74, F2: 302}
        _ = fi52
        _ = v55
    }
    for fi56 := 0; fi56 < 1; fi56++ {
        v57 := GenS5{F0: ("" + "a"), F1: (!false)}
        v58 := []int{}
        fmt.Println(p1, p2)
        _ = fi56
        _ = v57
        _ = v58
    }
    p2 = 214
    _ = p0
    _ = p1
    _ = p2
    _ = v51
    return 2
}

func main() {
    var v1 int32 = 521
    v2 := []int{2147483647, 2147483647}
    var v3 int = 9223372036854775807
    v4 := make(map[string]bool)
    v5 := []int{(v3 % ((v3) | 1)), v3}
    v1 += (v1 << ((922) & 63))
    for fi59 := 0; fi59 < 7; fi59++ {
        fmt.Println(v3)
        _ = fi59
    }
    var v60 uint64 = 236
    var v61 int64 = 14
    v61 = (v61 >> ((4294967295) & 63))
    var v62 int8 = 0
    for fi63 := 0; fi63 < 2; fi63++ {
        var v64 uint64 = v60
        _ = fi63
        _ = v64
    }
    for fi65 := 0; fi65 < 12; fi65++ {
        var v66 uint16 = 0
        for fi67 := 0; fi67 < 7; fi67++ {
            if false {
                fmt.Println(v66)
                switch 89 {
                case 32768:
                    v2 = append(v2, v3)
                    var v68 uint = 908
                    _ = v68
                case -564:
                    v3 -= fi65
                case 816:
                    v69 := GenS2{F0: v60, F1: 32767, F2: v1}
                    _ = v69
                default:
                    switch "\n\t" {
                    case "hello":
                        if (v66 >= (^(v66))) {
                            switch fi65 {
                            case -748:
                                v62 = v62
                            case -1:
                                v62 += (v62 | (v62 + (v62 << ((18446744073709551615) & 63))))
                            case 354:
                                switch v3 {
                                case -9223372036854775808:
                                    switch fi67 {
                                    case -9223372036854775808:
                                        v5 = append(v5, fi65)
                                    case 128:
                                        if false {
                                            var v70 int = fi67
                                            _ = v70
                                        } else {
                                            v60 -= (v60 / (((v60 >> ((4294967295) & 63))) | 1))
                                        }
                                    }
                                    delete(v4, (("" + ("hello" + "foo bar")) + "hello"))
                                case -1:
                                    if (fi67 >= (^(v3))) {
                                        if false {
                                            v62 = (v62 - 11)
                                        }
                                    }
                                default:
                                    if ("\n\t" > ("" + "foo bar")) {
                                        v71 := GenS8{F0: v60}
                                        v61 = int64(v60)
                                        if (((1e10 / (float64(-2.25) + 3.14159)) < 0.0) && (!(969 >= 4294967295))) {
                                            v72 := v4[("foo bar" + "")]
                                            switch fi67 {
                                            case -2147483648:
                                                fmt.Println(v3, v60, fi67)
                                            case 339:
                                                v3 += (^(v3))
                                            default:
                                                if ((v72 || (!true)) && (1 > (v1 >> ((744) & 63)))) {
                                                    fmt.Println(v3, v3)
                                                }
                                            }
                                            _ = v72
                                        }
                                        _ = v71
                                    } else {
                                        v3 -= 9223372036854775807
                                        v73 := GenS9{F0: ("foo bar" + "foo bar"), F1: (v3 / ((-2147483648) | 1)), F2: (v60 | (v60 & 427))}
                                        var v74 float64 = (1.5 * 3.14159)
                                        _ = v73
                                        _ = v74
                                    }
                                }
                                if ((((v61 + 1) <= v61) && (128 >= -2147483648)) && (genFn0(true, ("a" + "hello")) > int32(v3))) {
                                    if ((v60 > (v60 << ((166) & 63))) && (255 < 4294967295)) {
                                        if (((!(((true || (((!false) || ((255 >= 650) || false)) && (v60 < 42))) || (255 != 42)) || (42 == 118))) && (77 == 128)) && false) {
                                            switch "" {
                                            case "\n\t":
                                                if (-32768 != 1000) {
                                                    var v75 bool = false
                                                    _ = v75
                                                } else {
                                                    var v76 int32 = 16
                                                    v76 = -2147483648
                                                    switch "a" {
                                                    case "\n\t":
                                                        v77 := GenS10{F0: (false || (((false || ((v66 >> ((2) & 63)) >= v66)) || (v62 < 127)) && (65535 > 4294967295))), F1: 186, F2: 8}
                                                        _ = v77
                                                    case "hello":
                                                        v78 := GenS3{F0: -5, F1: fi67, F2: 1}
                                                        _ = v78
                                                    case "foo bar":
                                                        fmt.Println(fi67, v3, v66)
                                                    default:
                                                        var v79 uint32 = 4294967295
                                                        _ = v79
                                                    }
                                                    var v80 uint64 = 4294967295
                                                    _ = v76
                                                    _ = v80
                                                }
                                            case "hello":
                                                v81 := GenS6{F0: 42, F1: (fi65 * fi67)}
                                                _ = v81
                                            case "foo bar":
                                                switch "" {
                                                case "foo bar":
                                                    v1 = (-(v1))
                                                default:
                                                    v4["hello"] = false
                                                }
                                            default:
                                                var v82 int64 = v61
                                                _ = v82
                                            }
                                            if (912 < 293) {
                                                if (true && (!((830 > 860) || (!((1.5 / 1.0) <= ((0.1 * float64(-2.25)) + (1.5 + 0.0))))))) {
                                                    _ = len(v5)
                                                    if len(v5) > 0 {
                                                        v5[0] = (fi65 + -1)
                                                    }
                                                } else {
                                                    v3 = (fi65 - (fi67 / ((-1) | 1)))
                                                }
                                            }
                                        }
                                        if (true || (!true)) {
                                            fmt.Println(v66, v61, v3)
                                            var v83 int8 = (v62 + (v62 - (v62 / (((v62 * -128)) | 1))))
                                            v84 := GenS1{F0: (v83 + -121)}
                                            _ = v83
                                            _ = v84
                                        }
                                        var v85 string = ""
                                        _ = v85
                                    }
                                }
                                if true {
                                    if true {
                                        if (v3 != -2147483648) {
                                            _ = len(v5)
                                            fmt.Println(v66)
                                        } else {
                                            if (false || ((fi65 != (fi67 + v3)) || false)) {
                                                v86 := GenS1{F0: (v62 & -1)}
                                                switch (^(fi65)) {
                                                case 42:
                                                    switch (fi65 >> ((2) & 63)) {
                                                    case -333:
                                                        v1 = (v1 >> ((2) & 63))
                                                    default:
                                                        var v87 int32 = v1
                                                        _ = v87
                                                    }
                                                default:
                                                    v88 := GenS11{F0: v66}
                                                    _ = v88
                                                }
                                                var v89 int32 = v1
                                                _ = v86
                                                _ = v89
                                            } else {
                                                if false {
                                                    v90 := GenS12{F0: (!(!(false || (0 == 32768))))}
                                                    _ = v90
                                                } else {
                                                    if ((2 <= 0) || true) {
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            v91 := GenS12{F0: false}
                                                            _ = v91
                                                        default:
                                                            var v92 uint16 = v66
                                                            _ = v92
                                                        }
                                                        fmt.Println(fi65, v62, v62)
                                                        v93 := GenS13{F0: -9223372036854775808}
                                                        v62 += (v62 % ((v62) | 1))
                                                        _ = v93
                                                    }
                                                }
                                                if true {
                                                    v2 = append(v2, fi65)
                                                }
                                                var v94 float64 = (float64((1e10 + (0.0 / 1e10))) * 3.14159)
                                                _ = v94
                                            }
                                            v1 += (^(v1))
                                        }
                                        fmt.Println(v60, v3, v61)
                                        var v95 uint = 255
                                        v96 := GenS14{F0: (v95 - (v95 / ((18446744073709551615) | 1)))}
                                        _ = v95
                                        _ = v96
                                    } else {
                                        _ = len(v2)
                                        if (!false) {
                                            if true {
                                                v3 -= -1
                                                var v97 int64 = (v61 | -9223372036854775808)
                                                _ = v97
                                            } else {
                                                v98 := GenS12{F0: (("" + "a") < "foo bar")}
                                                v99 := GenS5{F0: (("hello" + "") + "hello"), F1: ((!false) && true)}
                                                var v100 uint8 = 0
                                                _ = v100
                                                _ = v98
                                                _ = v99
                                            }
                                            fmt.Println(v62, v1)
                                            v101 := GenS2{F0: v60, F1: 42, F2: genFn0(true, "hello")}
                                            _ = v101
                                        }
                                        var v102 int16 = -1
                                        _ = v102
                                    }
                                    v103 := GenS11{F0: (v66 | v66)}
                                    _ = v103
                                } else {
                                    v104 := GenS4{F0: -1, F1: false}
                                    var v105 int32 = genFn0(((!true) && true), "foo bar")
                                    if true {
                                        if (true && (255 < 1)) {
                                            v4[("hello" + "foo bar")] = true
                                            var v106 int16 = -1
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            _ = v106
                                        } else {
                                            v3 = (-(fi65))
                                            var v107 int16 = int16(v105)
                                            _ = v107
                                        }
                                        fmt.Println(v62)
                                    } else {
                                        delete(v4, "\n\t")
                                        v108 := GenS15{F0: true}
                                        fmt.Println(v60, v105, v1)
                                        _ = v108
                                    }
                                    _ = v104
                                    _ = v105
                                }
                            }
                        } else {
                            if true {
                                if true {
                                    var v109 int64 = v61
                                    if (float32((3.14159 / 3.14159)) != 1.5) {
                                        if ((1.5 < float64(float32(1.5))) || ((((("hello" + "") + "") + "a") > ("hello" + "a")) && true)) {
                                            switch (fi65 ^ int(v66)) {
                                            case 975:
                                                var v110 float64 = 0.1
                                                _ = v110
                                            case 128:
                                                v3 = fi65
                                            case 127:
                                                var v111 uint8 = 128
                                                _ = v111
                                            }
                                            v1 -= -1
                                            _ = len(v5)
                                        } else {
                                            _ = len(v4)
                                            v60 += v60
                                            switch v3 {
                                            case -1:
                                                v62 += (v62 % (((v62 << ((42) & 63))) | 1))
                                            case -360:
                                                var v112 float32 = ((1.5 * (3.14159 + (float32(3.14159) + ((float32(1e10) + float32(3.14159)) + 1.5)))) / 1e10)
                                                _ = v112
                                            case -744:
                                                v60 += (v60 ^ (v60 ^ v60))
                                            }
                                            var v113 int32 = (v1 & (v1 * v1))
                                            _ = v113
                                        }
                                        _ = len(v4)
                                    } else {
                                        fmt.Println(v1, v61, v66)
                                    }
                                    _ = v109
                                } else {
                                    var v114 uint = 2
                                    if true {
                                        var v115 uint = v114
                                        _ = v115
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = (fi67 ^ 127)
                                        }
                                        fmt.Println(v60)
                                        switch 128 {
                                        case 0:
                                            switch ("a" + "hello") {
                                            case "a":
                                                delete(v4, "")
                                            }
                                        }
                                    }
                                    delete(v4, "foo bar")
                                    v5 = append(v5, (fi65 / ((2147483647) | 1)))
                                    _ = v114
                                }
                                switch ("hello" + "a") {
                                case "":
                                    switch "foo bar" {
                                    case "a":
                                        v66 -= (^(v66))
                                    case "\n\t":
                                        if (true && true) {
                                            v61 += (v61 | 0)
                                            if true {
                                                if (175 <= 154) {
                                                    v116 := GenS4{F0: 42, F1: false}
                                                    v3 += 32768
                                                    switch "" {
                                                    case "a":
                                                        fmt.Println(fi65, v3)
                                                    case "hello":
                                                        if (-2.25 > 0.1) {
                                                            fmt.Println(v3, v60, v61)
                                                        } else {
                                                            if ((-855 > -1) && (0.1 >= 1.5)) {
                                                                fmt.Println(v62, v1)
                                                                v117 := GenS7{F0: (fi65 >> ((803) & 63)), F1: (v61 + v61), F2: 65535}
                                                                _ = v117
                                                            } else {
                                                                fmt.Println(fi67, v1)
                                                            }
                                                            delete(v4, "")
                                                            var v118 float64 = -2.25
                                                            _ = v118
                                                        }
                                                    }
                                                    _ = v116
                                                } else {
                                                    var v119 uint32 = 4294967295
                                                    fmt.Println(v60, v61, v61)
                                                    var v120 uint64 = (^(v60))
                                                    if true {
                                                        fmt.Println(v62, v119)
                                                        var v121 bool = (false && false)
                                                        v121 = (!(v121 || (v121 && false)))
                                                        _ = v121
                                                    } else {
                                                        var v122 int32 = v1
                                                        _ = v122
                                                    }
                                                    v119 = (v119 % ((genFn1(255, (^(v66)), (v66 * v66))) | 1))
                                                    fmt.Println(v60, v60)
                                                    _ = v119
                                                    _ = v120
                                                }
                                                v123 := GenS8{F0: (^(v60))}
                                                _ = v123
                                            } else {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                            delete(v4, ((("hello" + (("" + "") + "\n\t")) + "foo bar") + ("foo bar" + ("\n\t" + ((("hello" + ("" + "hello")) + "hello") + "\n\t")))))
                                        }
                                    case "hello":
                                        if true {
                                            v61 -= v61
                                        }
                                    default:
                                        v1 += -2147483648
                                    }
                                case "foo bar":
                                    var v124 int8 = int8(v3)
                                    _ = v124
                                case "\n\t":
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                default:
                                    v61 = (-(v61))
                                }
                            }
                            v4[""] = false
                        }
                    default:
                        if true {
                            if (3.14159 != (1.5 * float64(0.1))) {
                                var v125 float32 = float32(-2.25)
                                var v126 int8 = 100
                                var v127 float64 = -2.25
                                _ = v125
                                _ = v126
                                _ = v127
                            } else {
                                switch ("a" + "a") {
                                case "a":
                                    v3 += 128
                                case "foo bar":
                                    var v128 uint16 = 65535
                                    _ = v128
                                case "\n\t":
                                    if (0.0 > 1.5) {
                                        switch "a" {
                                        case "\n\t":
                                            if false {
                                                var v129 float32 = 0.0
                                                v130 := GenS6{F0: 191, F1: fi67}
                                                _ = v129
                                                _ = v130
                                            }
                                        }
                                        v1 = (-(v1))
                                    } else {
                                        if true {
                                            if ((^(v60)) != 343) {
                                                v131 := GenS1{F0: v62}
                                                _ = v131
                                            } else {
                                                var v132 float64 = float64(-2.25)
                                                var v133 bool = false
                                                v134 := GenS14{F0: uint(v60)}
                                                _ = v132
                                                _ = v133
                                                _ = v134
                                            }
                                            v135 := GenS5{F0: "", F1: true}
                                            v3 += (fi65 - int(v60))
                                            _ = v135
                                        } else {
                                            v61 = v61
                                        }
                                        fmt.Println(v66, v60)
                                        if (!(v1 > v1)) {
                                            v136 := v4[(("\n\t" + "hello") + "\n\t")]
                                            v137 := GenS12{F0: v136}
                                            fmt.Println(v60, v136, v136)
                                            _ = v136
                                            _ = v137
                                        } else {
                                            _ = len(v5)
                                            if false {
                                                switch "hello" {
                                                case "\n\t":
                                                    var v138 int = -1
                                                    _ = v138
                                                case "a":
                                                    v3 = fi65
                                                }
                                                v4[("foo bar" + "\n\t")] = true
                                                var v139 uint32 = 469
                                                _ = v139
                                            } else {
                                                if (true || true) {
                                                    switch "hello" {
                                                    case "foo bar":
                                                        v140 := GenS6{F0: 42, F1: -527}
                                                        _ = v140
                                                    case "a":
                                                        var v141 uint = 18446744073709551615
                                                        _ = v141
                                                    }
                                                    switch fi67 {
                                                    case 1:
                                                        var v142 uint64 = 0
                                                        _ = v142
                                                    case 9223372036854775807:
                                                        v1 -= v1
                                                    }
                                                    v143 := GenS11{F0: (v66 * v66)}
                                                    _ = v143
                                                } else {
                                                    v61 = v61
                                                }
                                                v60 -= 2
                                                if (!(42 == 128)) {
                                                    if false {
                                                        if len(v2) > 0 {
                                                            v2[0] = -1
                                                        }
                                                        v60 -= (v60 - 42)
                                                    } else {
                                                        v61 = 184
                                                        v2 = append(v2, (v3 ^ (v3 * fi65)))
                                                        var v144 uint64 = 2
                                                        _ = v144
                                                    }
                                                }
                                            }
                                        }
                                    }
                                default:
                                    switch v3 {
                                    case -171:
                                        var v145 uint64 = (^(v60))
                                        _ = v145
                                    case 32767:
                                        var v146 uint32 = 492
                                        _ = v146
                                    }
                                }
                                if false {
                                    v60 = (^(v60))
                                    v147 := GenS14{F0: 882}
                                    v66 = (v66 - 42)
                                    _ = v147
                                }
                                var v148 int = 4294967295
                                _ = v148
                            }
                        }
                    }
                }
                v5 = append(v5, (fi65 & (fi67 * v3)))
                v149 := GenS12{F0: true}
                _ = v149
            } else {
                v1 = (-(v1))
                var v150 int8 = int8(v61)
                v150 -= int8(v66)
                _ = v150
            }
            v1 = (v1 >> ((457) & 63))
            _ = fi67
        }
        _ = fi65
        _ = v66
    }
    if true {
        for fi151 := 0; fi151 < 17; fi151++ {
            if true {
                _ = len(v2)
                v3 -= 9223372036854775807
            } else {
                for fi152 := 0; fi152 < 12; fi152++ {
                    if ("hello" >= "foo bar") {
                        v61 += v61
                        v3 -= v3
                    }
                    var v153 int16 = 42
                    v62 = v62
                    _ = fi152
                    _ = v153
                }
                v154 := GenS11{F0: 42}
                v60 = 4294967295
                _ = v154
            }
            _ = fi151
        }
        switch (-(v3)) {
        case 405:
            var v155 float32 = (-2.25 / 1.0)
            _ = v155
        default:
            var v156 int64 = (v61 >> ((65535) & 63))
            _ = v156
        }
        var v157 uint16 = 0
        _ = v157
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (0.0 <= 1.5) {
            fmt.Println(v1, v62)
            v4[("\n\t" + ("" + ("hello" + ("a" + ""))))] = false
        }
        fmt.Println(v3, v1, v62)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v61
    _ = v62
}
