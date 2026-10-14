package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int16; F2 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint; F1 uint64 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int16; F1 uint32 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 int16; F1 int8 }
type GenS17 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS18 struct { F0 int32; F1 uint8 }
type GenS19 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS20 struct { F0 uint16 }
type GenS21 struct { F0 int8 }
type GenS22 struct { F0 int }
type GenS23 struct { F0 uint16; F1 string; F2 uint16 }
type GenS24 struct { F0 uint8; F1 bool }
type GenS25 struct { F0 uint16; F1 uint8 }
type GenS26 struct { F0 uint16; F1 bool }
type GenS27 struct { F0 int8; F1 int8; F2 string }
type GenS28 struct { F0 uint16; F1 int64; F2 uint8 }
func genFn0(p0 bool, p1 int64, p2 uint) uint32 {
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := []int{fi7, fi7}
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        switch "a" {
        case "":
            for fi9 := 0; fi9 < 11; fi9++ {
                _ = len(v8)
                if true {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    switch (((("foo bar" + "") + (("" + "foo bar") + "hello")) + "foo bar") + ("hello" + "")) {
                    case "":
                        switch "" {
                        case "":
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        case "\n\t":
                            fmt.Println(fi7, fi9, fi9)
                        case "a":
                            switch fi9 {
                            case 127:
                                if len(v8) > 0 {
                                    v8[0] = 157
                                }
                            case -896:
                                var v10 uint16 = 124
                                _ = v10
                            default:
                                fmt.Println(p0)
                            }
                        default:
                            p1 = int64(fi7)
                        }
                    case "\n\t":
                        var v11 float32 = (0.1 / (float32(1e10) - 0.1))
                        _ = v11
                    case "a":
                        p2 -= p2
                    default:
                        var v12 bool = p0
                        _ = v12
                    }
                    switch (-(fi9)) {
                    case 32768:
                        v8 = append(v8, 1)
                    case 2147483647:
                        p1 = (p1 | p1)
                        p2 = (p2 << (((p2 - p2)) & 63))
                        p2 = (^(p2))
                    case -347:
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                    default:
                        switch fi7 {
                        case 0:
                            v13 := GenS1{F0: (^(p2)), F1: ("foo bar" + "hello"), F2: 1}
                            _ = v13
                        case 656:
                            fmt.Println(fi9)
                        }
                    }
                } else {
                    p0 = p0
                }
                _ = len(v8)
                _ = fi9
            }
        case "foo bar":
            p2 += p2
        case "a":
            v14 := GenS2{F0: -9223372036854775808, F1: 65535}
            _ = v14
        }
        _ = fi7
        _ = v8
    }
    v15 := []int{32768}
    if len(v15) > 0 {
        v15[0] = 128
    }
    var v16 int8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 2147483647
}
func genFn1(p0 string, p1 int8) string {
    var v17 int = -654
    if true {
        for fi18 := 0; fi18 < 17; fi18++ {
            fmt.Println(p0, fi18, p0)
            if true {
                p0 = ""
                var v19 int32 = 2147483647
                _ = v19
            } else {
                var v20 uint32 = 2147483647
                v21 := GenS2{F0: 1, F1: 0}
                _ = v20
                _ = v21
            }
            var v22 bool = false
            _ = fi18
            _ = v22
        }
        for fi23 := 0; fi23 < 13; fi23++ {
            switch p0 {
            case "":
                var v24 bool = true
                _ = v24
            default:
                var v25 string = "\n\t"
                _ = v25
            }
            for fi26 := 0; fi26 < 19; fi26++ {
                v27 := []int{(v17 - 1), fi26}
                switch (p0 + ((p0 + "") + p0)) {
                case "hello":
                    if len(v27) > 0 {
                        v27[0] = v17
                    }
                    p1 = p1
                    p1 -= -97
                case "foo bar":
                    if len(v27) >= 2 {
                        v27 = v27[1:len(v27)]
                    }
                case "\n\t":
                    p1 = (p1 * 1)
                    fmt.Println(p0, p0)
                }
                _ = fi26
                _ = v27
            }
            v28 := []int{fi23, (fi23 - fi23), 4294967295}
            _ = fi23
            _ = v28
        }
        p1 += (p1 % ((p1) | 1))
    }
    var v29 uint32 = 65535
    _ = p0
    _ = p1
    _ = v17
    _ = v29
    return p0
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]uint32)
    v4 := []int{875, 888, 127}
    v5 := make(map[int]string)
    var v6 int16 = -1
    var v30 float32 = 1.5
    for fi31 := 0; fi31 < 18; fi31++ {
        v32 := GenS3{F0: 0, F1: -1, F2: true}
        _ = len(v1)
        _ = fi31
        _ = v32
    }
    v33 := GenS3{F0: 1, F1: (v6 * v6), F2: ((("" + "") == "foo bar") && (1 == 1))}
    if (37 > 42) {
        v6 += (^(v6))
        var v34 int8 = 100
        _ = v34
    } else {
        var v35 uint16 = 3
        v30 += 3.14159
        switch 42 {
        case 58:
            if (1000 < v6) {
                var v36 uint = 65535
                v6 += v6
                _ = v36
            } else {
                v35 += v35
                for fi37 := 0; fi37 < 19; fi37++ {
                    var v38 int64 = 153
                    var v39 float64 = 0.1
                    var v40 float64 = v39
                    _ = fi37
                    _ = v38
                    _ = v39
                    _ = v40
                }
                v4 = append(v4, 42)
            }
        case 4294967295:
            var v41 uint8 = 1
            _ = v41
        }
        _ = v35
    }
    if true {
        if true {
            if (!false) {
                if false {
                    v2[-9223372036854775808] = 411
                    for fi42 := 0; fi42 < 18; fi42++ {
                        v43 := GenS3{F0: 1, F1: -1, F2: true}
                        _ = fi42
                        _ = v43
                    }
                    var v44 uint16 = 322
                    _ = v44
                } else {
                    if true {
                        var v45 uint32 = 214
                        v46 := v2[32768]
                        v47 := GenS3{F0: 935, F1: 1, F2: ((v6 / ((v6) | 1)) < v6)}
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                }
                fmt.Println(v6, v6)
                for fi48 := 0; fi48 < 9; fi48++ {
                    fmt.Println(fi48, v30)
                    if true {
                        if (135 >= 826) {
                            switch (fi48 >> ((381) & 63)) {
                            case -240:
                                if true {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    for fi49 := 0; fi49 < 11; fi49++ {
                                        fmt.Println(fi49, v30)
                                        fmt.Println(fi48, v30, v30)
                                        _ = fi49
                                    }
                                }
                            case -607:
                                for fi50 := 0; fi50 < 10; fi50++ {
                                    v51 := GenS4{F0: (true || false)}
                                    _ = fi50
                                    _ = v51
                                }
                            default:
                                _ = len(v3)
                            }
                            v30 -= (3.14159 - (float32(((0.0 / (1.5 + (0.0 * ((3.14159 * (((1e10 * -2.25) - (float64(1e10) - -2.25)) + (0.0 / 0.1))) / 1e10)))) / 3.14159)) + ((v30 + v30) / v30)))
                            for fi52 := 0; fi52 < 5; fi52++ {
                                v30 = v30
                                _ = fi52
                            }
                        } else {
                            v2[(fi48 & fi48)] = 2
                            var v53 int8 = 78
                            if true {
                                for fi54 := 0; fi54 < 15; fi54++ {
                                    switch "\n\t" {
                                    case "\n\t":
                                        fmt.Println(fi48, v30)
                                    case "foo bar":
                                        var v55 int16 = (v6 >> ((216) & 63))
                                        _ = v55
                                    default:
                                        if false {
                                            v6 -= v6
                                            var v56 int32 = -1
                                            _ = v56
                                        } else {
                                            fmt.Println(v53, fi48, fi48)
                                            if ((false || true) && false) {
                                                var v57 int16 = v6
                                                if (true && false) {
                                                    var v58 uint = 725
                                                    _ = v58
                                                } else {
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        delete(v2, -325)
                                                        fmt.Println(v6, v57)
                                                        var v59 uint32 = 630
                                                        _ = v59
                                                    default:
                                                        var v60 float32 = -2.25
                                                        _ = v60
                                                    }
                                                    var v61 uint = 245
                                                    _ = v61
                                                }
                                                switch "a" {
                                                case "":
                                                    var v62 uint32 = 0
                                                    _ = v62
                                                case "foo bar":
                                                    var v63 int8 = (v53 >> ((4294967295) & 63))
                                                    _ = v63
                                                case "\n\t":
                                                    fmt.Println(fi48)
                                                default:
                                                    v64 := GenS5{F0: 508, F1: 18446744073709551615}
                                                    _ = v64
                                                }
                                                _ = v57
                                            } else {
                                                fmt.Println(fi54, fi54)
                                                v6 = 32767
                                            }
                                            var v65 bool = true
                                            _ = v65
                                        }
                                    }
                                    v66 := GenS6{F0: (v6 - 1000)}
                                    v67 := GenS7{F0: false, F1: 2}
                                    if (!true) {
                                        v53 = 100
                                    } else {
                                        switch fi54 {
                                        case -1:
                                            _ = len(v2)
                                        case 32768:
                                            v30 = (3.14159 + float32(0.0))
                                        case 42:
                                            v68 := GenS8{F0: "\n\t"}
                                            _ = v68
                                        }
                                        v69 := GenS9{F0: 627, F1: 18446744073709551615}
                                        _ = v69
                                    }
                                    delete(v1, ("\n\t" + "foo bar"))
                                    if true {
                                        _ = len(v1)
                                    } else {
                                        var v70 uint32 = 0
                                        v70 -= v70
                                        _ = v70
                                    }
                                    delete(v1, "foo bar")
                                    _ = fi54
                                    _ = v66
                                    _ = v67
                                }
                            } else {
                                _ = len(v3)
                            }
                            _ = v53
                        }
                        v3[genFn0(false, 1, 1)] = 561
                    } else {
                        switch ("foo bar" + ("\n\t" + ("foo bar" + ""))) {
                        case "":
                            _ = len(v4)
                        case "\n\t":
                            fmt.Println(fi48, v30, v6)
                        case "hello":
                            var v71 int16 = (v6 + v6)
                            _ = v71
                        default:
                            if ((v30 / (float32((0.0 + (1e10 / ((-2.25 * (1e10 - 0.1)) + (float64(0.0) - (float64(v30) / 1.0)))))) + (1e10 * v30))) == (float32(((0.1 + 1e10) - 0.0)) / (v30 - 3.14159))) {
                                var v72 int32 = 767
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v72
                            }
                        }
                    }
                    _ = fi48
                }
            }
            fmt.Println(v30, v6)
            if true {
                if true {
                    if ((42 != 2) || true) {
                        switch "a" {
                        case "foo bar":
                            switch -9223372036854775808 {
                            case -2147483648:
                                if (float64(3.14159) == 0.0) {
                                    v30 = 0.0
                                    if true {
                                        v30 += v30
                                    } else {
                                        if len(v4) > 0 {
                                            v4[0] = -182
                                        }
                                        for fi73 := 0; fi73 < 10; fi73++ {
                                            var v74 int8 = 100
                                            switch "hello" {
                                            case "\n\t":
                                                var v75 int16 = v6
                                                _ = v75
                                            }
                                            for fi76 := 0; fi76 < 10; fi76++ {
                                                v74 = v74
                                                if false {
                                                    v74 = 127
                                                }
                                                var v77 bool = true
                                                if ((v30 + (float32(-2.25) / v30)) > v30) {
                                                    v78 := GenS7{F0: (277 == 908), F1: 2}
                                                    _ = v78
                                                }
                                                _ = fi76
                                                _ = v77
                                            }
                                            _ = fi73
                                            _ = v74
                                        }
                                        if true {
                                            v79 := GenS10{F0: 127}
                                            fmt.Println(v30, v30)
                                            _ = v79
                                        }
                                    }
                                } else {
                                    v80 := GenS11{F0: 1}
                                    v30 += ((3.14159 / ((v30 * v30) + v30)) + v30)
                                    _ = v80
                                }
                            case -9223372036854775808:
                                fmt.Println(v6, v30)
                            }
                        case "":
                            var v81 int64 = 318
                            _ = v81
                        case "hello":
                            switch 128 {
                            case 32768:
                                v6 = v6
                            case 109:
                                v6 -= (v6 ^ (v6 << ((609) & 63)))
                            case 32767:
                                switch ((((("hello" + "hello") + (("a" + (((("foo bar" + "\n\t") + "hello") + "\n\t") + (("a" + "hello") + "\n\t"))) + "a")) + "a") + "a") + ("\n\t" + "hello")) {
                                case "\n\t":
                                    var v82 int8 = 127
                                    _ = v82
                                case "":
                                    if false {
                                        for fi83 := 0; fi83 < 7; fi83++ {
                                            var v84 int16 = 1000
                                            _ = fi83
                                            _ = v84
                                        }
                                        _ = len(v1)
                                        if (9223372036854775807 < 0) {
                                            v6 += (^(v6))
                                            var v85 uint = 2
                                            _ = v85
                                        } else {
                                            for fi86 := 0; fi86 < 18; fi86++ {
                                                switch fi86 {
                                                case 42:
                                                    v6 = (v6 | 0)
                                                case -1:
                                                    switch (-(fi86)) {
                                                    case 249:
                                                        v6 = v6
                                                    case -331:
                                                        v87 := GenS10{F0: 100}
                                                        _ = v87
                                                    default:
                                                        var v88 int16 = -438
                                                        _ = v88
                                                    }
                                                }
                                                v6 += v6
                                                var v89 int8 = -128
                                                _ = fi86
                                                _ = v89
                                            }
                                            fmt.Println(v30, v30)
                                            v90 := GenS5{F0: 80, F1: 4294967295}
                                            _ = v90
                                        }
                                    }
                                }
                            default:
                                if false {
                                    for fi91 := 0; fi91 < 3; fi91++ {
                                        v92 := GenS12{F0: fi91}
                                        var v93 int64 = 299
                                        _ = fi91
                                        _ = v92
                                        _ = v93
                                    }
                                    switch "hello" {
                                    case "":
                                        for fi94 := 0; fi94 < 7; fi94++ {
                                            if (0.0 <= (float64(v30) - 0.0)) {
                                                v30 = -2.25
                                                var v95 int = (fi94 & 85)
                                                _ = v95
                                            }
                                            if false {
                                                var v96 float32 = v30
                                                var v97 string = ((("foo bar" + "\n\t") + "foo bar") + "hello")
                                                _ = v96
                                                _ = v97
                                            } else {
                                                _ = len(v4)
                                            }
                                            _ = fi94
                                        }
                                    case "hello":
                                        for fi98 := 0; fi98 < 8; fi98++ {
                                            switch "hello" {
                                            case "foo bar":
                                                var v99 uint = 18446744073709551615
                                                _ = v99
                                            case "":
                                                v100 := v3[2]
                                                _ = v100
                                            default:
                                                var v101 uint = 0
                                                _ = v101
                                            }
                                            _ = len(v4)
                                            for fi102 := 0; fi102 < 5; fi102++ {
                                                _ = len(v4)
                                                var v103 int = -79
                                                _ = fi102
                                                _ = v103
                                            }
                                            _ = fi98
                                        }
                                    case "\n\t":
                                        v104 := GenS1{F0: 786, F1: "foo bar", F2: 65535}
                                        _ = v104
                                    }
                                } else {
                                    var v105 string = "foo bar"
                                    fmt.Println(v105)
                                    if true {
                                        v106 := GenS1{F0: 1, F1: "", F2: 221}
                                        v6 -= v6
                                        _ = v106
                                    }
                                    _ = v105
                                }
                            }
                        default:
                            v107 := GenS13{F0: (v6 << ((18446744073709551615) & 63)), F1: 0}
                            _ = v107
                        }
                        v108 := GenS13{F0: v6, F1: 2147483647}
                        if false {
                            if true {
                                var v109 int32 = -1
                                if true {
                                    var v110 uint = 42
                                    if true {
                                        if true {
                                            if (9223372036854775807 > -968) {
                                                if true {
                                                    for fi111 := 0; fi111 < 6; fi111++ {
                                                        for fi112 := 0; fi112 < 4; fi112++ {
                                                            fmt.Println(fi112)
                                                            fmt.Println(v6, v110, v109)
                                                            if len(v4) > 0 {
                                                                v4[0] = fi111
                                                            }
                                                            _ = fi112
                                                        }
                                                        v109 -= int32(fi111)
                                                        var v113 int8 = -128
                                                        _ = fi111
                                                        _ = v113
                                                    }
                                                    var v114 uint8 = 42
                                                    var v115 uint = 98
                                                    _ = v114
                                                    _ = v115
                                                } else {
                                                    v116 := v2[-2147483648]
                                                    v117 := v1["foo bar"]
                                                    for fi118 := 0; fi118 < 4; fi118++ {
                                                        if true {
                                                            _ = len(v5)
                                                            v119 := GenS14{F0: 1}
                                                            _ = v119
                                                        }
                                                        switch fi118 {
                                                        case 0:
                                                            _ = len(v4)
                                                        case -2147483648:
                                                            var v120 uint32 = v116
                                                            _ = v120
                                                        default:
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                        fmt.Println(fi118)
                                                        _ = fi118
                                                    }
                                                    _ = v116
                                                    _ = v117
                                                }
                                                v110 += (v110 * 0)
                                                v121 := GenS15{F0: 255}
                                                _ = v121
                                            } else {
                                                var v122 bool = ((!(0 == 918)) && true)
                                                v6 = 351
                                                var v123 uint = 255
                                                _ = v122
                                                _ = v123
                                            }
                                            v30 -= v30
                                        } else {
                                            v109 -= v109
                                        }
                                    } else {
                                        var v124 int32 = v109
                                        v125 := GenS2{F0: 478, F1: 577}
                                        _ = v124
                                        _ = v125
                                    }
                                    _ = v110
                                }
                                switch 127 {
                                case 127:
                                    for fi126 := 0; fi126 < 7; fi126++ {
                                        switch "" {
                                        case "a":
                                            switch fi126 {
                                            case 42:
                                                if true {
                                                    switch fi126 {
                                                    case -9223372036854775808:
                                                        var v127 int64 = 1
                                                        _ = v127
                                                    default:
                                                        v109 = (-(v109))
                                                    }
                                                    switch "hello" {
                                                    case "\n\t":
                                                        if (!false) {
                                                            _ = len(v5)
                                                            for fi128 := 0; fi128 < 6; fi128++ {
                                                                if (!(false || false)) {
                                                                    if (0.1 < v30) {
                                                                        fmt.Println(fi128, v109)
                                                                        var v129 float32 = float32(float64(v30))
                                                                        _ = v129
                                                                    }
                                                                } else {
                                                                    v109 -= v109
                                                                    if true {
                                                                        var v130 float32 = ((1e10 - float32((0.1 * -2.25))) - (float32(float64(v30)) * v30))
                                                                        _ = v130
                                                                    } else {
                                                                        v109 += (v109 & (v109 * v109))
                                                                        if false {
                                                                            fmt.Println(fi128, v30)
                                                                            var v131 uint8 = 0
                                                                            _ = v131
                                                                        }
                                                                        v109 += 874
                                                                        if (783 > 1) {
                                                                            v132 := GenS16{F0: (v6 >> ((658) & 63)), F1: -1}
                                                                            var v133 int16 = 1
                                                                            v6 = v133
                                                                            _ = v132
                                                                            _ = v133
                                                                        }
                                                                        if len(v4) > 0 {
                                                                            v4[0] = -1
                                                                        }
                                                                    }
                                                                    if false {
                                                                        v109 = 1
                                                                        var v134 float64 = ((0.0 - 1e10) / 1.5)
                                                                        _ = v134
                                                                    } else {
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                    }
                                                                }
                                                                _ = len(v1)
                                                                _ = fi128
                                                            }
                                                            v109 = v109
                                                        }
                                                    case "a":
                                                        for fi135 := 0; fi135 < 1; fi135++ {
                                                            v136 := GenS11{F0: (^(v109))}
                                                            if true {
                                                                if false {
                                                                    fmt.Println(fi135, fi126, fi135)
                                                                } else {
                                                                    switch (fi126 / (((fi126 - fi126)) | 1)) {
                                                                    case 1:
                                                                        if true {
                                                                            v6 += v6
                                                                        } else {
                                                                            var v137 bool = true
                                                                            v138 := v2[fi135]
                                                                            var v139 bool = v137
                                                                            v140 := GenS5{F0: 420, F1: 697}
                                                                            _ = v137
                                                                            _ = v138
                                                                            _ = v139
                                                                            _ = v140
                                                                        }
                                                                    default:
                                                                        if true {
                                                                            v109 = (v109 | v109)
                                                                            switch fi126 {
                                                                            case 9223372036854775807:
                                                                                _ = len(v4)
                                                                            }
                                                                        }
                                                                        var v141 int = (fi126 >> ((65535) & 63))
                                                                        switch fi135 {
                                                                        case 9223372036854775807:
                                                                            var v142 float64 = ((0.1 - (float64(float32(-2.25)) - (float64(0.0) / ((1e10 * (-2.25 - 0.0)) - 3.14159)))) - (0.1 / -2.25))
                                                                            _ = v142
                                                                        }
                                                                        _ = v141
                                                                    }
                                                                    fmt.Println(v30, fi126)
                                                                }
                                                            } else {
                                                                fmt.Println(v6, v6)
                                                                fmt.Println(v30, v6)
                                                            }
                                                            switch fi135 {
                                                            case 196:
                                                                if ("hello" < "\n\t") {
                                                                    if (false && (!false)) {
                                                                        v109 = 1
                                                                    }
                                                                    if false {
                                                                        var v143 int = fi135
                                                                        _ = v143
                                                                    } else {
                                                                        v144 := GenS17{F0: (v109 << ((18446744073709551615) & 63)), F1: 1, F2: 470}
                                                                        _ = v144
                                                                    }
                                                                    v30 -= (-2.25 / 1.0)
                                                                } else {
                                                                    var v145 bool = (v6 <= v6)
                                                                    _ = v145
                                                                }
                                                            case 872:
                                                                v1["\n\t"] = (genFn0(true, 0, 612) >= 721)
                                                            case 2147483647:
                                                                switch "" {
                                                                case "hello":
                                                                    switch "a" {
                                                                    case "a":
                                                                        var v146 float64 = 0.1
                                                                        _ = v146
                                                                    case "hello":
                                                                        if len(v4) > 0 {
                                                                            v4[0] = fi135
                                                                        }
                                                                    }
                                                                case "a":
                                                                    var v147 uint64 = 42
                                                                    _ = v147
                                                                }
                                                            default:
                                                                var v148 float32 = 1.5
                                                                _ = v148
                                                            }
                                                            if len(v4) > 0 {
                                                                v4[0] = (fi135 >> ((1) & 63))
                                                            }
                                                            switch fi135 {
                                                            case 32768:
                                                                if true {
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    v149 := GenS18{F0: v109, F1: 117}
                                                                    var v150 bool = false
                                                                    var v151 uint32 = 702
                                                                    _ = v149
                                                                    _ = v150
                                                                    _ = v151
                                                                } else {
                                                                    var v152 bool = (true && true)
                                                                    v153 := GenS9{F0: 304, F1: 18446744073709551615}
                                                                    if (!v152) {
                                                                        switch fi126 {
                                                                        case -2147483648:
                                                                            var v154 uint32 = 189
                                                                            _ = v154
                                                                        case -258:
                                                                            fmt.Println(fi126, v109)
                                                                        case -1:
                                                                            v155 := GenS1{F0: 18446744073709551615, F1: "", F2: 255}
                                                                            _ = v155
                                                                        }
                                                                    } else {
                                                                        var v156 int64 = 42
                                                                        fmt.Println(v30, v156)
                                                                        v157 := GenS19{F0: 120, F1: 128, F2: v109}
                                                                        _ = v156
                                                                        _ = v157
                                                                    }
                                                                    _ = v152
                                                                    _ = v153
                                                                }
                                                            case 4294967295:
                                                                v109 += v109
                                                            default:
                                                                var v158 int32 = (v109 >> ((0) & 63))
                                                                _ = v158
                                                            }
                                                            var v159 bool = false
                                                            var v160 uint16 = 658
                                                            _ = fi135
                                                            _ = v136
                                                            _ = v159
                                                            _ = v160
                                                        }
                                                    }
                                                    var v161 bool = true
                                                    _ = v161
                                                } else {
                                                    v162 := GenS9{F0: 2, F1: 0}
                                                    var v163 uint64 = 0
                                                    if (false && true) {
                                                        switch fi126 {
                                                        case -773:
                                                            v164 := GenS12{F0: 32768}
                                                            _ = v164
                                                        case 9223372036854775807:
                                                            if (true && true) {
                                                                var v165 int64 = 42
                                                                _ = v165
                                                            } else {
                                                                var v166 int32 = v109
                                                                _ = v166
                                                            }
                                                        case 2147483647:
                                                            var v167 uint32 = 635
                                                            _ = v167
                                                        default:
                                                            fmt.Println(v30)
                                                        }
                                                    } else {
                                                        for fi168 := 0; fi168 < 19; fi168++ {
                                                            v109 -= 1
                                                            _ = fi168
                                                        }
                                                        v163 = v163
                                                        var v169 float32 = v30
                                                        _ = v169
                                                    }
                                                    _ = v162
                                                    _ = v163
                                                }
                                            case 865:
                                                var v170 uint16 = 65535
                                                _ = v170
                                            case 2147483647:
                                                if false {
                                                    v30 += 1e10
                                                    if (float32(float64(v30)) != (3.14159 / 0.1)) {
                                                        v171 := GenS20{F0: 255}
                                                        _ = v171
                                                    }
                                                    var v172 int16 = 1000
                                                    _ = v172
                                                }
                                            default:
                                                for fi173 := 0; fi173 < 18; fi173++ {
                                                    v4 = append(v4, int(v6))
                                                    v30 = v30
                                                    _ = fi173
                                                }
                                            }
                                        case "":
                                            if (2 > 2) {
                                                for fi174 := 0; fi174 < 13; fi174++ {
                                                    var v175 uint64 = 1
                                                    _ = fi174
                                                    _ = v175
                                                }
                                                v176 := GenS21{F0: 66}
                                                _ = v176
                                            } else {
                                                for fi177 := 0; fi177 < 12; fi177++ {
                                                    v6 -= (v6 / (((^(v6))) | 1))
                                                    _ = fi177
                                                }
                                                if (!(-684 > -2147483648)) {
                                                    if false {
                                                        if (!(true || false)) {
                                                            v178 := GenS1{F0: 909, F1: "", F2: 0}
                                                            _ = v178
                                                        } else {
                                                            var v179 int8 = 75
                                                            _ = len(v2)
                                                            fmt.Println(v6, v30)
                                                            _ = v179
                                                        }
                                                        fmt.Println(v6)
                                                        delete(v1, ((((("foo bar" + "foo bar") + "\n\t") + (("\n\t" + "") + "a")) + (("" + ("a" + "\n\t")) + "a")) + ""))
                                                    } else {
                                                        for fi180 := 0; fi180 < 15; fi180++ {
                                                            var v181 int8 = 42
                                                            if true {
                                                                v182 := GenS22{F0: fi126}
                                                                _ = v182
                                                            }
                                                            _ = fi180
                                                            _ = v181
                                                        }
                                                        var v183 uint8 = 0
                                                        _ = len(v4)
                                                        _ = v183
                                                    }
                                                }
                                                switch "a" {
                                                case "":
                                                    fmt.Println(v109, fi126, v109)
                                                case "a":
                                                    v184 := GenS23{F0: 982, F1: (("" + "") + "a"), F2: 198}
                                                    _ = v184
                                                default:
                                                    if true {
                                                        if (false && true) {
                                                            switch "hello" {
                                                            case "hello":
                                                                if true {
                                                                    if (100 <= 0) {
                                                                        fmt.Println(v6, fi126, fi126)
                                                                        v30 = float32(-2.25)
                                                                    }
                                                                    v185 := GenS1{F0: 4294967295, F1: "foo bar", F2: 66}
                                                                    fmt.Println(fi126, fi126)
                                                                    _ = v185
                                                                } else {
                                                                    v186 := v2[(fi126 & fi126)]
                                                                    _ = v186
                                                                }
                                                            case "foo bar":
                                                                var v187 int = (^(fi126))
                                                                _ = v187
                                                            case "a":
                                                                if (128 > 122) {
                                                                    var v188 int32 = (v109 * (v109 | 2147483647))
                                                                    fmt.Println(v109)
                                                                    _ = v188
                                                                } else {
                                                                    for fi189 := 0; fi189 < 20; fi189++ {
                                                                        fmt.Println(v109, fi126)
                                                                        v30 += v30
                                                                        fmt.Println(v109, v6, v109)
                                                                        _ = fi189
                                                                    }
                                                                    fmt.Println(v30, v30)
                                                                }
                                                            }
                                                        }
                                                        v5[fi126] = ("a" + "")
                                                        for fi190 := 0; fi190 < 15; fi190++ {
                                                            var v191 int64 = 1
                                                            v191 = 26
                                                            delete(v1, "")
                                                            v191 += 558
                                                            _ = fi190
                                                            _ = v191
                                                        }
                                                    } else {
                                                        var v192 int = 0
                                                        v6 += v6
                                                        _ = v192
                                                    }
                                                }
                                            }
                                        default:
                                            if len(v4) > 0 {
                                                v4[0] = (fi126 ^ (fi126 % ((2147483647) | 1)))
                                            }
                                        }
                                        v193 := GenS10{F0: 3}
                                        if true {
                                            v194 := GenS24{F0: 252, F1: (0 > 901)}
                                            var v195 int8 = -113
                                            fmt.Println(v195)
                                            _ = v194
                                            _ = v195
                                        }
                                        _ = fi126
                                        _ = v193
                                    }
                                }
                                _ = v109
                            } else {
                                v30 = v30
                            }
                            v2[-314] = 210
                        }
                        _ = v108
                    } else {
                        var v196 float64 = 3.14159
                        var v197 uint = 815
                        _ = v196
                        _ = v197
                    }
                    var v198 int = 127
                    v198 = (v198 ^ (v198 + v198))
                    _ = v198
                }
                v199 := GenS25{F0: 99, F1: 93}
                switch -9223372036854775808 {
                case 621:
                    switch "" {
                    case "":
                        _ = len(v1)
                    }
                }
                _ = v199
            }
        }
    } else {
        var v200 int8 = 0
        _ = v200
    }
    v30 = v30
    _ = len(v4)
    v201 := GenS26{F0: 519, F1: true}
    v6 -= (v6 + v6)
    var v202 int64 = 9223372036854775807
    v203 := GenS27{F0: 100, F1: -1, F2: "hello"}
    switch -9223372036854775808 {
    case 42:
        var v204 int8 = 114
        _ = v204
    case -2147483648:
        if (255 != 421) {
            v30 -= v30
            v202 = (v202 - (v202 << ((882) & 63)))
        }
    }
    if true {
        switch -136 {
        case 32768:
            var v205 uint8 = 29
            _ = v205
        case -9223372036854775808:
            v206 := GenS28{F0: 793, F1: v202, F2: 0}
            _ = v206
        }
        v207 := GenS12{F0: 4294967295}
        fmt.Println(v30, v202, v202)
        _ = v207
    } else {
        if false {
            for fi208 := 0; fi208 < 10; fi208++ {
                for fi209 := 0; fi209 < 5; fi209++ {
                    v202 = v202
                    var v210 bool = (!false)
                    _ = fi209
                    _ = v210
                }
                v202 -= v202
                _ = fi208
            }
            v211 := GenS7{F0: ((v6 & (v6 | 1)) < -32768), F1: 0}
            var v212 int32 = -2147483648
            _ = v211
            _ = v212
        }
        v213 := GenS22{F0: 0}
        var v214 uint8 = 0
        _ = v213
        _ = v214
    }
    fmt.Println(v30, v6)
    if true {
        for fi215 := 0; fi215 < 6; fi215++ {
            v202 += 1
            var v216 uint8 = 1
            var v217 int64 = (-(v202))
            _ = fi215
            _ = v216
            _ = v217
        }
        if (true || (!true)) {
            for fi218 := 0; fi218 < 6; fi218++ {
                v202 = (v202 | (v202 - v202))
                var v219 float64 = 3.14159
                switch fi218 {
                case -9223372036854775808:
                    v220 := v5[fi218]
                    _ = v220
                default:
                    _ = len(v2)
                }
                _ = fi218
                _ = v219
            }
            _ = len(v4)
        }
        for fi221 := 0; fi221 < 20; fi221++ {
            var v222 int = fi221
            _ = fi221
            _ = v222
        }
    } else {
        var v223 int8 = 1
        _ = v223
    }
    if (0 > 65535) {
        if (962 >= -2147483648) {
            if len(v4) > 0 {
                v4[0] = 42
            }
            _ = len(v1)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v202:", v202)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v201
    _ = v202
    _ = v203
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
