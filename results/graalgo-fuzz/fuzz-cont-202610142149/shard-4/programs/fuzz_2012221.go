package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint64; F1 uint }
type GenS11 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS12 struct { F0 uint16; F1 uint16 }
type GenS13 struct { F0 int16; F1 int64; F2 string }
type GenS14 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 uint32; F1 uint; F2 uint }
type GenS18 struct { F0 bool; F1 uint }
type GenS19 struct { F0 uint32 }
type GenS20 struct { F0 int16; F1 uint16 }
type GenS21 struct { F0 uint64; F1 int8 }
type GenS22 struct { F0 uint64; F1 int; F2 int8 }
func genFn0(p0 string, p1 uint16, p2 int8) uint8 {
    var v6 int16 = 32767
    switch ((("" + "foo bar") + ("\n\t" + (((p0 + "foo bar") + "\n\t") + ""))) + p0) {
    case "foo bar":
        p0 = ((("foo bar" + "\n\t") + (p0 + "a")) + (("a" + "") + (p0 + p0)))
    case "a":
        var v7 int = 127
        _ = v7
    case "":
        for fi8 := 0; fi8 < 9; fi8++ {
            p2 = (p2 - 100)
            v6 -= v6
            v9 := GenS1{F0: p0, F1: (^(p2)), F2: (v6 ^ 393)}
            _ = fi8
            _ = v9
        }
    }
    var v10 float32 = ((1.5 / 3.14159) - -2.25)
    switch 127 {
    case 2147483647:
        fmt.Println(v10)
    case 0:
        var v11 uint16 = (^(p1))
        _ = v11
    case 4294967295:
        v12 := []int{988, 4294967295, -826, 128}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    return 255
}
func genFn1(p0 int64, p1 int64, p2 int8) uint {
    p0 = p0
    var v13 int16 = -1
    v14 := []int{32767}
    for fi15 := 0; fi15 < 4; fi15++ {
        if true {
            for fi16 := 0; fi16 < 20; fi16++ {
                v17 := GenS2{F0: 292, F1: 65535}
                var v18 uint64 = 0
                _ = fi16
                _ = v17
                _ = v18
            }
            v19 := GenS3{F0: -2147483648, F1: 42}
            fmt.Println(fi15, fi15, p2)
            _ = v19
        } else {
            v20 := GenS4{F0: 1}
            fmt.Println(p2, fi15, p1)
            _ = v20
        }
        p1 = p0
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 185
}

func main() {
    var v1 string = ("" + "a")
    var v2 uint32 = 65535
    var v3 uint64 = 42
    var v4 uint16 = 42
    var v5 float32 = 1e10
    var v21 int8 = -1
    v5 = v5
    if (v21 >= int8(v4)) {
        v2 -= v2
        switch (("hello" + v1) + v1) {
        case "foo bar":
            v5 -= v5
        case "":
            v22 := []int{}
            _ = v22
        }
    }
    v2 -= 65535
    var v23 uint32 = v2
    v24 := []int{}
    _ = len(v24)
    if (91 != uint8(v4)) {
        _ = len(v24)
        switch 4294967295 {
        case -690:
            var v25 float64 = -2.25
            _ = v25
        case -9223372036854775808:
            for fi26 := 0; fi26 < 7; fi26++ {
                switch (fi26 | 877) {
                case 1:
                    if true {
                        var v27 string = ((v1 + (v1 + "")) + v1)
                        v28 := GenS1{F0: "a", F1: -128, F2: 572}
                        var v29 int = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    } else {
                        var v30 bool = false
                        _ = v30
                    }
                case 42:
                    v31 := GenS5{F0: v3, F1: 9223372036854775807, F2: 1}
                    _ = v31
                default:
                    for fi32 := 0; fi32 < 4; fi32++ {
                        v21 -= -1
                        if (!false) {
                            var v33 uint8 = 42
                            if (v21 != v21) {
                                v34 := GenS1{F0: v1, F1: -1, F2: 1000}
                                var v35 string = (v1 + v1)
                                _ = v34
                                _ = v35
                            }
                            v23 -= (v2 - 42)
                            var v36 int16 = 137
                            v37 := GenS4{F0: v23}
                            _ = v33
                            _ = v36
                            _ = v37
                        } else {
                            fmt.Println(v21, v21)
                            v38 := GenS1{F0: ("\n\t" + v1), F1: (v21 - (-(v21))), F2: 290}
                            _ = v38
                        }
                        _ = fi32
                    }
                }
                var v39 int64 = -1
                _ = fi26
                _ = v39
            }
        case 1:
            var v40 uint8 = 127
            _ = v40
        default:
            fmt.Println(v2, v2, v5)
        }
        if ((v4 + v4) < 832) {
            switch -106 {
            case 32768:
                for fi41 := 0; fi41 < 5; fi41++ {
                    for fi42 := 0; fi42 < 11; fi42++ {
                        if (18446744073709551615 > 4294967295) {
                            if ((^(v3)) == (v3 << ((18446744073709551615) & 63))) {
                                if (390 <= 188) {
                                    _ = len(v24)
                                    if len(v24) >= 2 {
                                        v24 = v24[1:len(v24)]
                                    }
                                }
                                var v43 uint = 42
                                v44 := GenS3{F0: int32(v21), F1: 0}
                                switch v1 {
                                case "a":
                                    v45 := GenS4{F0: v2}
                                    _ = v45
                                case "foo bar":
                                    if false {
                                        v23 = uint32(v4)
                                        v21 = (v21 / ((1) | 1))
                                        v46 := GenS2{F0: 725, F1: (v23 * (v2 ^ 1))}
                                        var v47 int16 = -690
                                        v48 := GenS6{F0: fi41}
                                        switch fi42 {
                                        case 884:
                                            var v49 float32 = v5
                                            _ = v49
                                        case -9223372036854775808:
                                            switch "a" {
                                            case "a":
                                                fmt.Println(fi41, v21)
                                            case "\n\t":
                                                switch -1 {
                                                case 4294967295:
                                                    var v50 uint16 = (v4 << ((uint(v3)) & 63))
                                                    _ = v50
                                                case 128:
                                                    var v51 float64 = 0.1
                                                    _ = v51
                                                case 0:
                                                    if false {
                                                        var v52 float64 = ((1.5 / 1.5) + 1.5)
                                                        v53 := GenS2{F0: 2147483648, F1: (v23 - v2)}
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                }
                                            default:
                                                _ = len(v24)
                                            }
                                        }
                                        _ = v46
                                        _ = v47
                                        _ = v48
                                    }
                                case "":
                                    fmt.Println(v21, v43, v3)
                                default:
                                    v54 := GenS5{F0: v3, F1: -9223372036854775808, F2: (v3 % ((v3) | 1))}
                                    _ = v54
                                }
                                _ = v43
                                _ = v44
                            }
                        }
                        var v55 string = "a"
                        _ = fi42
                        _ = v55
                    }
                    _ = fi41
                }
            case 32767:
                var v56 uint64 = v3
                _ = v56
            case 2147483647:
                if (-367 == -1) {
                    for fi57 := 0; fi57 < 5; fi57++ {
                        var v58 bool = ((fi57 | 32768) >= fi57)
                        var v59 uint64 = 42
                        var v60 float32 = v5
                        _ = fi57
                        _ = v58
                        _ = v59
                        _ = v60
                    }
                    if (!(!false)) {
                        v61 := GenS7{F0: ("" + ("" + "\n\t")), F1: -1}
                        _ = v61
                    }
                    v5 -= 1e10
                }
            }
            var v62 float64 = 0.1
            var v63 uint16 = uint16(v2)
            _ = v62
            _ = v63
        }
    } else {
        var v64 uint8 = 42
        if false {
            v65 := GenS6{F0: -689}
            var v66 int64 = -200
            fmt.Println(v4, v66, v3)
            _ = v65
            _ = v66
        } else {
            fmt.Println(v2)
        }
        v21 = v21
        _ = v64
    }
    switch v1 {
    case "foo bar":
        v67 := GenS4{F0: (v23 - (v23 << ((2) & 63)))}
        _ = v67
    case "":
        if true {
            switch v1 {
            case "hello":
                v68 := GenS8{F0: v2, F1: 1, F2: (v4 * (^(v4)))}
                _ = v68
            case "foo bar":
                v2 += v23
            }
            _ = len(v24)
            fmt.Println(v23)
        } else {
            fmt.Println(v23, v3, v2)
            var v69 uint16 = (v4 + (v4 >> ((155) & 63)))
            _ = v69
        }
    case "\n\t":
        var v70 bool = (1 != 1)
        _ = v70
    }
    if true {
        v24 = append(v24, 128)
        fmt.Println(v3, v4, v4)
        if false {
            v2 += 42
            if true {
                v3 = 1
                switch 9223372036854775807 {
                case 1:
                    if false {
                        var v71 float64 = (1e10 - -2.25)
                        _ = v71
                    } else {
                        for fi72 := 0; fi72 < 12; fi72++ {
                            v73 := GenS9{F0: v21}
                            for fi74 := 0; fi74 < 6; fi74++ {
                                var v75 int32 = 2147483647
                                if len(v24) >= 2 {
                                    v24 = v24[1:len(v24)]
                                }
                                switch fi72 {
                                case 0:
                                    var v76 int32 = (v75 & int32(v23))
                                    _ = v76
                                }
                                _ = fi74
                                _ = v75
                            }
                            if false {
                                v77 := GenS7{F0: "\n\t", F1: -618}
                                var v78 bool = (((^(v3)) <= (v3 / ((0) | 1))) || (int64(v3) <= -9223372036854775808))
                                _ = v77
                                _ = v78
                            }
                            _ = fi72
                            _ = v73
                        }
                        if false {
                            v5 += v5
                        } else {
                            if false {
                                if len(v24) > 0 {
                                    v24[0] = 42
                                }
                                if ((^(v4)) < 2) {
                                    if ((v3 % (((v3 / ((v3) | 1))) | 1)) > (v3 << ((255) & 63))) {
                                        v3 -= 620
                                        var v79 uint32 = v2
                                        _ = v79
                                    }
                                } else {
                                    for fi80 := 0; fi80 < 17; fi80++ {
                                        if true {
                                            for fi81 := 0; fi81 < 1; fi81++ {
                                                switch 128 {
                                                case 42:
                                                    var v82 int = fi81
                                                    _ = v82
                                                case -1:
                                                    v3 += 4294967295
                                                case 128:
                                                    if (false || ((v21 - -128) == v21)) {
                                                        v2 = (v23 ^ (v23 - (v23 & (^(v23)))))
                                                    } else {
                                                        if true {
                                                            var v83 int32 = int32(v21)
                                                            v3 = v3
                                                            _ = len(v24)
                                                            _ = v83
                                                        } else {
                                                            var v84 bool = (!false)
                                                            var v85 int8 = 1
                                                            _ = v84
                                                            _ = v85
                                                        }
                                                        v86 := GenS10{F0: v3, F1: 255}
                                                        _ = v86
                                                    }
                                                default:
                                                    switch "" {
                                                    case "foo bar":
                                                        fmt.Println(v1, fi81)
                                                    default:
                                                        var v87 uint16 = v4
                                                        _ = v87
                                                    }
                                                }
                                                var v88 bool = false
                                                switch "\n\t" {
                                                case "hello":
                                                    v2 = v2
                                                }
                                                _ = fi81
                                                _ = v88
                                            }
                                            var v89 uint16 = 1
                                            _ = v89
                                        }
                                        if true {
                                            v2 -= v2
                                            var v90 uint32 = 2147483647
                                            _ = v90
                                        }
                                        _ = fi80
                                    }
                                }
                            } else {
                                v23 -= (v23 - v2)
                            }
                            switch (v1 + "hello") {
                            case "foo bar":
                                for fi91 := 0; fi91 < 20; fi91++ {
                                    v92 := GenS11{F0: v21, F1: (v2 | 692), F2: 41}
                                    _ = fi91
                                    _ = v92
                                }
                            }
                        }
                    }
                case 985:
                    if (!(!((v3 | (v3 >> ((283) & 63))) == (^(v3))))) {
                        for fi93 := 0; fi93 < 8; fi93++ {
                            var v94 int64 = 42
                            for fi95 := 0; fi95 < 19; fi95++ {
                                v24 = append(v24, (fi95 ^ (-(fi93))))
                                v96 := GenS12{F0: (v4 >> ((2) & 63)), F1: v4}
                                var v97 int = (-(fi93))
                                _ = fi95
                                _ = v96
                                _ = v97
                            }
                            v1 = v1
                            _ = fi93
                            _ = v94
                        }
                        switch 1 {
                        case 809:
                            for fi98 := 0; fi98 < 7; fi98++ {
                                var v99 uint = 400
                                fmt.Println(v21, v3)
                                _ = fi98
                                _ = v99
                            }
                        case 42:
                            var v100 string = ""
                            _ = v100
                        default:
                            v101 := GenS13{F0: 452, F1: 9223372036854775807, F2: ""}
                            _ = v101
                        }
                        for fi102 := 0; fi102 < 14; fi102++ {
                            v4 += (v4 % ((32767) | 1))
                            if (fi102 > fi102) {
                                switch fi102 {
                                case 0:
                                    switch (v1 + ("hello" + ("" + ((v1 + v1) + "foo bar")))) {
                                    case "foo bar":
                                        _ = len(v24)
                                    case "hello":
                                        v2 += v2
                                    case "\n\t":
                                        var v103 int = (fi102 * 127)
                                        _ = v103
                                    default:
                                        var v104 string = ((v1 + v1) + "foo bar")
                                        _ = v104
                                    }
                                case 2147483647:
                                    if (false && (v1 <= v1)) {
                                        if len(v24) > 0 {
                                            v24[0] = 911
                                        }
                                        v105 := GenS11{F0: (v21 | (^(v21))), F1: v2, F2: 0}
                                        _ = v105
                                    }
                                case 1:
                                    for fi106 := 0; fi106 < 15; fi106++ {
                                        if false {
                                            if false {
                                                if len(v24) >= 2 {
                                                    v24 = v24[1:len(v24)]
                                                }
                                            }
                                            v23 = (v23 | v2)
                                            if len(v24) >= 2 {
                                                v24 = v24[1:len(v24)]
                                            }
                                        } else {
                                            if (!true) {
                                                v5 = (v5 - v5)
                                            } else {
                                                v5 = v5
                                            }
                                        }
                                        _ = fi106
                                    }
                                default:
                                    var v107 int16 = 883
                                    _ = v107
                                }
                                var v108 int32 = 42
                                _ = v108
                            } else {
                                for fi109 := 0; fi109 < 5; fi109++ {
                                    switch 4294967295 {
                                    case -1:
                                        var v110 string = v1
                                        _ = v110
                                    case -936:
                                        v111 := GenS3{F0: -2147483648, F1: -377}
                                        _ = v111
                                    case 42:
                                        switch "hello" {
                                        case "hello":
                                            switch fi102 {
                                            case 939:
                                                var v112 int32 = int32(v3)
                                                _ = v112
                                            default:
                                                v1 = v1
                                            }
                                        case "":
                                            if ((fi102 ^ 4294967295) < (^(fi109))) {
                                                v113 := GenS13{F0: -1, F1: 1, F2: ("foo bar" + (v1 + ((("hello" + v1) + v1) + (v1 + (("foo bar" + v1) + "foo bar")))))}
                                                v114 := GenS3{F0: -1, F1: 656}
                                                _ = v113
                                                _ = v114
                                            }
                                        default:
                                            var v115 uint16 = 0
                                            _ = v115
                                        }
                                    default:
                                        fmt.Println(fi109, v5)
                                    }
                                    _ = fi109
                                }
                            }
                            var v116 uint64 = v3
                            _ = fi102
                            _ = v116
                        }
                    } else {
                        switch "hello" {
                        case "":
                            if true {
                                for fi117 := 0; fi117 < 20; fi117++ {
                                    _ = len(v24)
                                    if (!(3.14159 <= v5)) {
                                        var v118 uint64 = (v3 << ((65535) & 63))
                                        _ = v118
                                    } else {
                                        var v119 float32 = v5
                                        fmt.Println(fi117, v5, v21)
                                        _ = v119
                                    }
                                    _ = fi117
                                }
                                if len(v24) > 0 {
                                    v24[0] = -2147483648
                                }
                                v4 += 282
                            } else {
                                var v120 int8 = (v21 & 1)
                                switch 9223372036854775807 {
                                case 885:
                                    if false {
                                        if len(v24) > 0 {
                                            v24[0] = 32767
                                        }
                                    } else {
                                        for fi121 := 0; fi121 < 2; fi121++ {
                                            var v122 float64 = float64(v5)
                                            switch "hello" {
                                            case "a":
                                                var v123 bool = (!false)
                                                _ = v123
                                            }
                                            var v124 int16 = -1
                                            _ = fi121
                                            _ = v122
                                            _ = v124
                                        }
                                    }
                                case 4294967295:
                                    v5 -= (v5 - 1e10)
                                default:
                                    var v125 int32 = 890
                                    _ = v125
                                }
                                var v126 uint8 = 210
                                _ = v120
                                _ = v126
                            }
                        default:
                            v2 = v23
                        }
                        var v127 uint32 = v23
                        v1 = "a"
                        _ = v127
                    }
                }
            }
            var v128 uint32 = (v23 | v2)
            _ = v128
        } else {
            if true {
                switch 128 {
                case 127:
                    if true {
                        fmt.Println(v5, v21)
                        fmt.Println(v5, v3, v23)
                        switch 856 {
                        case 128:
                            var v129 string = (v1 + v1)
                            _ = v129
                        case 9223372036854775807:
                            v24 = append(v24, 127)
                        case 32767:
                            v2 = (v2 / ((42) | 1))
                        default:
                            if len(v24) > 0 {
                                v24[0] = 63
                            }
                        }
                    } else {
                        var v130 int8 = (v21 * (-(v21)))
                        var v131 uint32 = uint32(v4)
                        var v132 bool = false
                        _ = v130
                        _ = v131
                        _ = v132
                    }
                case -16:
                    var v133 int32 = 0
                    _ = v133
                default:
                    if len(v24) >= 2 {
                        v24 = v24[1:len(v24)]
                    }
                }
                switch 1 {
                case 6:
                    if (float64(v5) > (-2.25 / 1e10)) {
                        var v134 int16 = 0
                        _ = v134
                    } else {
                        v135 := GenS14{F0: (^(v3)), F1: (v4 & v4), F2: 65535}
                        switch 874 {
                        case 1:
                            if false {
                                if false {
                                    v21 -= 42
                                    switch (v1 + v1) {
                                    case "a":
                                        for fi136 := 0; fi136 < 5; fi136++ {
                                            var v137 int64 = 9223372036854775807
                                            for fi138 := 0; fi138 < 14; fi138++ {
                                                if len(v24) >= 2 {
                                                    v24 = v24[1:len(v24)]
                                                }
                                                _ = fi138
                                            }
                                            _ = fi136
                                            _ = v137
                                        }
                                    case "":
                                        var v139 bool = ((1e10 + -2.25) >= (((v5 * 1e10) * v5) + v5))
                                        _ = v139
                                    default:
                                        v1 = (((("\n\t" + v1) + ("" + v1)) + ("a" + v1)) + "")
                                    }
                                    if ((!true) && true) {
                                        for fi140 := 0; fi140 < 9; fi140++ {
                                            if ((true && false) || false) {
                                                if true {
                                                    for fi141 := 0; fi141 < 15; fi141++ {
                                                        if true {
                                                            _ = len(v24)
                                                            v142 := GenS15{F0: -895}
                                                            switch "" {
                                                            case "":
                                                                var v143 uint32 = (v23 % ((2) | 1))
                                                                _ = v143
                                                            case "a":
                                                                _ = len(v24)
                                                            case "\n\t":
                                                                fmt.Println(v1, fi140, fi140)
                                                            default:
                                                                if len(v24) >= 2 {
                                                                    v24 = v24[1:len(v24)]
                                                                }
                                                            }
                                                            _ = v142
                                                        }
                                                        v5 -= 3.14159
                                                        var v144 int8 = v21
                                                        v24 = append(v24, 32767)
                                                        v145 := GenS3{F0: 0, F1: 57}
                                                        _ = fi141
                                                        _ = v144
                                                        _ = v145
                                                    }
                                                }
                                                for fi146 := 0; fi146 < 5; fi146++ {
                                                    v1 = v1
                                                    var v147 int = fi140
                                                    _ = fi146
                                                    _ = v147
                                                }
                                            } else {
                                                _ = len(v24)
                                                if false {
                                                    fmt.Println(v23)
                                                    v148 := GenS12{F0: (v4 / ((986) | 1)), F1: (v4 << ((0) & 63))}
                                                    _ = v148
                                                } else {
                                                    fmt.Println(v23, fi140, v23)
                                                    var v149 string = v1
                                                    _ = v149
                                                }
                                            }
                                            fmt.Println(v2, fi140)
                                            switch v1 {
                                            case "":
                                                v2 -= 1
                                            case "\n\t":
                                                v5 -= ((v5 + v5) / v5)
                                            }
                                            _ = fi140
                                        }
                                        var v150 bool = true
                                        if v150 {
                                            v150 = v150
                                            fmt.Println(v5, v21, v1)
                                        }
                                        _ = v150
                                    }
                                }
                                if (!false) {
                                    var v151 string = v1
                                    _ = v151
                                } else {
                                    if (2 < 2) {
                                        if len(v24) > 0 {
                                            v24[0] = 32768
                                        }
                                        var v152 bool = (uint(v3) < 4294967295)
                                        _ = v152
                                    }
                                }
                                if len(v24) > 0 {
                                    v24[0] = 37
                                }
                            } else {
                                v153 := GenS5{F0: (v3 % (((v3 & v3)) | 1)), F1: -904, F2: (^(v3))}
                                _ = v153
                            }
                        case -218:
                            fmt.Println(v1, v3)
                        case 73:
                            if len(v24) > 0 {
                                v24[0] = 4294967295
                            }
                        default:
                            for fi154 := 0; fi154 < 17; fi154++ {
                                if false {
                                    switch 2147483647 {
                                    case -9223372036854775808:
                                        if ((1.5 + ((v5 + (1e10 * (0.1 + 1.5))) + 1e10)) < 1e10) {
                                            switch "\n\t" {
                                            case "foo bar":
                                                for fi155 := 0; fi155 < 2; fi155++ {
                                                    switch v1 {
                                                    case "foo bar":
                                                        v156 := GenS7{F0: v1, F1: 0}
                                                        _ = v156
                                                    case "\n\t":
                                                        if true {
                                                            if len(v24) > 0 {
                                                                v24[0] = fi155
                                                            }
                                                            v2 += v23
                                                        }
                                                    case "hello":
                                                        fmt.Println(v23, v21)
                                                        fmt.Println(v1, v1)
                                                        var v157 int = -781
                                                        _ = v157
                                                    default:
                                                        var v158 float64 = 1e10
                                                        v159 := GenS3{F0: 147, F1: 433}
                                                        var v160 int = (fi155 + (-(fi154)))
                                                        _ = v158
                                                        _ = v159
                                                        _ = v160
                                                    }
                                                    _ = fi155
                                                }
                                            case "\n\t":
                                                for fi161 := 0; fi161 < 3; fi161++ {
                                                    v2 = (v23 << ((18446744073709551615) & 63))
                                                    _ = fi161
                                                }
                                            case "hello":
                                                if true {
                                                    if (!true) {
                                                        v162 := GenS11{F0: (^(v21)), F1: v23, F2: 42}
                                                        _ = v162
                                                    }
                                                    fmt.Println(v2)
                                                }
                                            default:
                                                for fi163 := 0; fi163 < 18; fi163++ {
                                                    if false {
                                                        v164 := GenS5{F0: (v3 % ((v3) | 1)), F1: 9223372036854775807, F2: v3}
                                                        var v165 uint64 = (v3 << ((298) & 63))
                                                        var v166 int16 = 42
                                                        v166 -= int16(fi154)
                                                        var v167 float32 = (float32(3.14159) + (1.5 / v5))
                                                        _ = v164
                                                        _ = v165
                                                        _ = v166
                                                        _ = v167
                                                    } else {
                                                        if false {
                                                            fmt.Println(fi163, fi163)
                                                            _ = len(v24)
                                                            var v168 uint32 = v23
                                                            _ = v168
                                                        }
                                                        var v169 int = 0
                                                        v170 := GenS16{F0: 255}
                                                        _ = v169
                                                        _ = v170
                                                    }
                                                    v21 = (v21 | (v21 - int8(v4)))
                                                    v23 = v2
                                                    v23 = (v2 - v2)
                                                    if false {
                                                        if len(v24) > 0 {
                                                            v24[0] = fi163
                                                        }
                                                        v5 -= (v5 / v5)
                                                        if false {
                                                            if len(v24) > 0 {
                                                                v24[0] = fi163
                                                            }
                                                            v23 = (v23 - (v23 ^ v2))
                                                            v171 := GenS1{F0: v1, F1: v21, F2: 96}
                                                            switch (fi154 / ((42) | 1)) {
                                                            case -735:
                                                                var v172 uint16 = v4
                                                                _ = v172
                                                            default:
                                                                _ = len(v24)
                                                            }
                                                            if false {
                                                                _ = len(v24)
                                                                v2 = (v2 * (^(v2)))
                                                                fmt.Println(v1, v1)
                                                                v173 := GenS6{F0: (fi163 / ((fi154) | 1))}
                                                                _ = v173
                                                            } else {
                                                                var v174 string = (v1 + "foo bar")
                                                                v175 := GenS11{F0: -1, F1: v23, F2: 1}
                                                                _ = v174
                                                                _ = v175
                                                            }
                                                            var v176 int = (fi163 << ((1) & 63))
                                                            if (true || true) {
                                                                if len(v24) > 0 {
                                                                    v24[0] = (^(fi154))
                                                                }
                                                                fmt.Println(v21, v21)
                                                            } else {
                                                                switch ("" + v1) {
                                                                case "hello":
                                                                    switch fi163 {
                                                                    case 4294967295:
                                                                        v3 = v3
                                                                    }
                                                                case "a":
                                                                    var v177 int16 = 661
                                                                    _ = v177
                                                                }
                                                                switch (v176 ^ fi163) {
                                                                case 334:
                                                                    switch ("hello" + v1) {
                                                                    case "hello":
                                                                        v178 := GenS3{F0: 873, F1: -2147483648}
                                                                        _ = v178
                                                                    case "foo bar":
                                                                        v2 += v23
                                                                    }
                                                                default:
                                                                    v176 -= (^(v176))
                                                                    switch v1 {
                                                                    case "a":
                                                                        var v179 int8 = 42
                                                                        _ = v179
                                                                    case "foo bar":
                                                                        v1 = (("\n\t" + ("hello" + "\n\t")) + (v1 + (v1 + (v1 + ("hello" + "hello")))))
                                                                    default:
                                                                        if true {
                                                                            switch v176 {
                                                                            case 32768:
                                                                                if (!(fi154 < 42)) {
                                                                                    var v180 uint64 = 0
                                                                                    var v181 int16 = 42
                                                                                    v2 = 4294967295
                                                                                    _ = v180
                                                                                    _ = v181
                                                                                } else {
                                                                                    switch 605 {
                                                                                    case 42:
                                                                                        var v182 uint16 = 255
                                                                                        _ = len(v24)
                                                                                        _ = v182
                                                                                    case -2147483648:
                                                                                        v176 -= v176
                                                                                    default:
                                                                                        if len(v24) > 0 {
                                                                                            v24[0] = 2147483647
                                                                                        }
                                                                                        if false {
                                                                                            v24 = append(v24, (fi154 >> ((42) & 63)))
                                                                                            _ = len(v24)
                                                                                        } else {
                                                                                            switch v1 {
                                                                                            case "hello":
                                                                                                switch v1 {
                                                                                                case "foo bar":
                                                                                                    var v183 int32 = -1
                                                                                                    _ = v183
                                                                                                case "":
                                                                                                    if true {
                                                                                                        var v184 bool = true
                                                                                                        _ = v184
                                                                                                    } else {
                                                                                                        var v185 float32 = v5
                                                                                                        _ = v185
                                                                                                    }
                                                                                                    v4 = (v4 & 1)
                                                                                                }
                                                                                            }
                                                                                            if len(v24) >= 2 {
                                                                                                v24 = v24[1:len(v24)]
                                                                                            }
                                                                                            fmt.Println(v3)
                                                                                        }
                                                                                    }
                                                                                    var v186 int64 = 0
                                                                                    _ = v186
                                                                                }
                                                                            case -699:
                                                                                if (!false) {
                                                                                    if len(v24) > 0 {
                                                                                        v24[0] = (v176 - (v176 / ((fi154) | 1)))
                                                                                    }
                                                                                    if true {
                                                                                        v21 = -67
                                                                                    } else {
                                                                                        var v187 string = ("foo bar" + v1)
                                                                                        if len(v24) >= 2 {
                                                                                            v24 = v24[1:len(v24)]
                                                                                        }
                                                                                        v188 := GenS17{F0: (^(v2)), F1: 802, F2: genFn1(0, 9223372036854775807, (v21 * (-(v21))))}
                                                                                        _ = v187
                                                                                        _ = v188
                                                                                    }
                                                                                    switch (("hello" + (v1 + v1)) + v1) {
                                                                                    case "\n\t":
                                                                                        fmt.Println(v176)
                                                                                    case "hello":
                                                                                        var v189 int16 = -32768
                                                                                        var v190 string = ("\n\t" + "hello")
                                                                                        _ = v189
                                                                                        _ = v190
                                                                                    }
                                                                                }
                                                                            case 42:
                                                                                var v191 string = v1
                                                                                _ = v191
                                                                            }
                                                                            var v192 bool = ((false || (!(false && false))) || false)
                                                                            _ = len(v24)
                                                                            fmt.Println(v21)
                                                                            if len(v24) >= 2 {
                                                                                v24 = v24[1:len(v24)]
                                                                            }
                                                                            if false {
                                                                                if v192 {
                                                                                    v193 := GenS7{F0: ("hello" + ((((("" + ((v1 + ("\n\t" + (v1 + v1))) + v1)) + v1) + (("hello" + "\n\t") + ("" + ((v1 + "a") + (v1 + (v1 + v1)))))) + (v1 + "hello")) + "")), F1: 0}
                                                                                    v194 := GenS13{F0: -136, F1: 1, F2: (v1 + (((v1 + v1) + ("\n\t" + ("a" + v1))) + "a"))}
                                                                                    switch ((v1 + ("a" + "\n\t")) + ("foo bar" + (((v1 + "hello") + "foo bar") + ("foo bar" + v1)))) {
                                                                                    case "\n\t":
                                                                                        if false {
                                                                                            var v195 uint8 = uint8(v4)
                                                                                            var v196 int64 = -849
                                                                                            _ = v195
                                                                                            _ = v196
                                                                                        }
                                                                                    }
                                                                                    if len(v24) > 0 {
                                                                                        v24[0] = v176
                                                                                    }
                                                                                    v197 := GenS17{F0: uint32(v3), F1: 310, F2: genFn1(int64(fi163), 391, (v21 >> ((65535) & 63)))}
                                                                                    _ = v193
                                                                                    _ = v194
                                                                                    _ = v197
                                                                                } else {
                                                                                    v1 = (("a" + v1) + v1)
                                                                                    fmt.Println(v1, v176, v3)
                                                                                    var v198 int64 = 9223372036854775807
                                                                                    _ = v198
                                                                                }
                                                                                fmt.Println(v176)
                                                                            }
                                                                            switch ((v1 + "hello") + v1) {
                                                                            case "a":
                                                                                fmt.Println(v3)
                                                                            case "hello":
                                                                                v1 = v1
                                                                            default:
                                                                                fmt.Println(v1, fi163)
                                                                            }
                                                                            _ = v192
                                                                        }
                                                                    }
                                                                }
                                                                v199 := GenS10{F0: v3, F1: 1}
                                                                _ = v199
                                                            }
                                                            _ = v171
                                                            _ = v176
                                                        } else {
                                                            v1 = v1
                                                        }
                                                        var v200 float64 = (-2.25 * 0.1)
                                                        v2 -= (v2 >> ((uint(v23)) & 63))
                                                        if (!false) {
                                                            switch v1 {
                                                            case "a":
                                                                if true {
                                                                    var v201 uint16 = v4
                                                                    _ = v201
                                                                }
                                                            case "hello":
                                                                fmt.Println(v5, v2)
                                                            }
                                                            v202 := GenS18{F0: (v200 > 1.5), F1: 143}
                                                            v203 := GenS19{F0: v23}
                                                            _ = v202
                                                            _ = v203
                                                        }
                                                        var v204 uint64 = (v3 / (((^(v3))) | 1))
                                                        _ = v200
                                                        _ = v204
                                                    } else {
                                                        v205 := GenS2{F0: (^(v2)), F1: 42}
                                                        _ = v205
                                                    }
                                                    v3 += v3
                                                    _ = fi163
                                                }
                                            }
                                        } else {
                                            v206 := GenS20{F0: 1000, F1: 65535}
                                            var v207 bool = (2 != (v3 & uint64(v4)))
                                            _ = v206
                                            _ = v207
                                        }
                                    case 575:
                                        for fi208 := 0; fi208 < 8; fi208++ {
                                            v5 += (v5 * (v5 + v5))
                                            _ = fi208
                                        }
                                    case -438:
                                        switch v1 {
                                        case "\n\t":
                                            var v209 float64 = (1.5 - 0.1)
                                            _ = v209
                                        case "a":
                                            var v210 int64 = 0
                                            _ = v210
                                        case "":
                                            switch fi154 {
                                            case 127:
                                                var v211 float64 = 1.5
                                                _ = v211
                                            case 9223372036854775807:
                                                for fi212 := 0; fi212 < 6; fi212++ {
                                                    if ((0.1 >= 1e10) || (!false)) {
                                                        switch v1 {
                                                        case "hello":
                                                            v213 := GenS21{F0: (^(v3)), F1: 118}
                                                            if (true || true) {
                                                                var v214 float64 = 1e10
                                                                var v215 uint = 42
                                                                fmt.Println(v23)
                                                                _ = v214
                                                                _ = v215
                                                            } else {
                                                                fmt.Println(v23, v21, v2)
                                                            }
                                                            var v216 int = (fi154 & (fi212 - (fi154 >> ((2) & 63))))
                                                            _ = v213
                                                            _ = v216
                                                        case "a":
                                                            if (true && (!((uint(v23) > 625) || (!true)))) {
                                                                v3 = v3
                                                                var v217 uint32 = (v23 | (v23 & (^(v2))))
                                                                if len(v24) >= 2 {
                                                                    v24 = v24[1:len(v24)]
                                                                }
                                                                v4 += (v4 ^ (^(v4)))
                                                                switch v1 {
                                                                case "hello":
                                                                    var v218 float32 = ((0.0 - ((((v5 * v5) - (v5 + v5)) * (v5 / v5)) * 1e10)) + v5)
                                                                    _ = v218
                                                                }
                                                                v24 = append(v24, fi212)
                                                                _ = v217
                                                            } else {
                                                                var v219 int32 = 2147483647
                                                                _ = v219
                                                            }
                                                        default:
                                                            switch (v1 + "hello") {
                                                            case "\n\t":
                                                                v220 := GenS13{F0: -61, F1: -1, F2: (v1 + v1)}
                                                                _ = v220
                                                            default:
                                                                v1 = v1
                                                            }
                                                        }
                                                    } else {
                                                        var v221 int64 = 42
                                                        fmt.Println(v4)
                                                        _ = v221
                                                    }
                                                    _ = fi212
                                                }
                                            case -414:
                                                if len(v24) > 0 {
                                                    v24[0] = fi154
                                                }
                                            default:
                                                for fi222 := 0; fi222 < 10; fi222++ {
                                                    fmt.Println(v23)
                                                    v21 = v21
                                                    fmt.Println(v2, v5, v3)
                                                    _ = fi222
                                                }
                                            }
                                        }
                                    }
                                    v21 = v21
                                    switch "foo bar" {
                                    case "":
                                        v23 += (v2 + v2)
                                    case "foo bar":
                                        fmt.Println(v2, v2, v4)
                                    case "a":
                                        var v223 uint16 = v4
                                        _ = v223
                                    }
                                }
                                switch v1 {
                                case "a":
                                    _ = len(v24)
                                case "hello":
                                    var v224 float32 = 1.5
                                    _ = v224
                                case "\n\t":
                                    if ((((uint32(v3) > v23) && true) && false) || false) {
                                        v2 = v23
                                        v225 := GenS1{F0: "foo bar", F1: (v21 ^ int8(v3)), F2: 1}
                                        _ = v225
                                    }
                                default:
                                    for fi226 := 0; fi226 < 9; fi226++ {
                                        fmt.Println(v4, fi226)
                                        _ = fi226
                                    }
                                }
                                _ = fi154
                            }
                        }
                        v227 := GenS22{F0: v3, F1: 4294967295, F2: (v21 & (v21 * v21))}
                        _ = v135
                        _ = v227
                    }
                case -2147483648:
                    fmt.Println(v2)
                case 9223372036854775807:
                    v5 -= v5
                }
            }
            v1 = "a"
            var v228 float32 = ((1.5 * (v5 + (-2.25 / v5))) - ((-2.25 - (float32(3.14159) - float32(3.14159))) + v5))
            _ = v228
        }
    }
    var v229 uint16 = v4
    v21 = (-(v21))
    var v230 uint8 = 42
    var v231 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v229:", v229)
    fmt.Println("v23:", v23)
    fmt.Println("v230:", v230)
    fmt.Println("v231:", v231)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v229
    _ = v23
    _ = v230
    _ = v231
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
