package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int; F2 int }
type GenS3 struct { F0 int16; F1 uint; F2 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint32; F1 uint }
type GenS8 struct { F0 int32; F1 uint8 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint64; F1 int64 }
type GenS11 struct { F0 uint32; F1 int64 }
type GenS12 struct { F0 uint; F1 uint32; F2 string }
type GenS13 struct { F0 int16 }
func genFn0() uint64 {
    if false {
        if true {
            var v6 int = 32767
            switch "a" {
            case "\n\t":
                var v7 uint32 = 2147483648
                _ = v7
            }
            if (!false) {
                for fi8 := 0; fi8 < 7; fi8++ {
                    v9 := []int{}
                    _ = fi8
                    _ = v9
                }
                var v10 int8 = -87
                _ = v10
            }
            _ = v6
        } else {
            var v11 int64 = 9223372036854775807
            v12 := []int{}
            _ = v11
            _ = v12
        }
        var v13 uint32 = 65535
        if true {
            v13 -= v13
        } else {
            switch "" {
            case "foo bar":
                v14 := []int{722, -377}
                _ = v14
            default:
                if (255 <= 1) {
                    v15 := GenS1{F0: 32767}
                    _ = v15
                } else {
                    var v16 int32 = 65
                    _ = v16
                }
            }
            v17 := []int{-852, 42}
            _ = v17
        }
        _ = v13
    }
    var v18 int16 = 32767
    _ = v18
    return 1
}
func genFn1(p0 uint, p1 uint16) int16 {
    for fi19 := 0; fi19 < 17; fi19++ {
        p0 = p0
        switch "\n\t" {
        case "\n\t":
            v20 := GenS2{F0: 127, F1: fi19, F2: fi19}
            _ = v20
        case "hello":
            v21 := []int{(fi19 % ((fi19) | 1))}
            _ = v21
        case "foo bar":
            if true {
                var v22 int = 127
                _ = v22
            } else {
                switch ("a" + "") {
                case "":
                    v23 := []int{(^(fi19)), 32768, fi19}
                    _ = v23
                case "\n\t":
                    var v24 int32 = 1
                    _ = v24
                case "hello":
                    if false {
                        if (!(648 != genFn0())) {
                            v25 := []int{fi19, (fi19 >> (((p0 | (p0 ^ (p0 + uint(p1))))) & 63)), fi19}
                            _ = v25
                        } else {
                            for fi26 := 0; fi26 < 7; fi26++ {
                                switch "" {
                                case "foo bar":
                                    p0 += p0
                                case "\n\t":
                                    v27 := []int{fi19, int(p1), fi26, fi26}
                                    _ = v27
                                case "a":
                                    p1 -= 759
                                    v28 := []int{fi19, (fi26 >> ((p0) & 63))}
                                    _ = v28
                                }
                                _ = fi26
                            }
                        }
                        v29 := []int{fi19, (^(fi19)), (fi19 | (fi19 % ((fi19) | 1))), 9223372036854775807}
                        v30 := GenS3{F0: -32768, F1: (p0 >> ((8) & 63)), F2: -271}
                        _ = v29
                        _ = v30
                    } else {
                        if true {
                            switch (^(fi19)) {
                            case 127:
                                var v31 int8 = 127
                                _ = v31
                            }
                        } else {
                            fmt.Println(p0)
                        }
                    }
                }
            }
        }
        if ((("a" + "foo bar") + "a") != ("hello" + "")) {
            var v32 int8 = -1
            for fi33 := 0; fi33 < 6; fi33++ {
                p0 = (p0 - p0)
                _ = fi33
            }
            _ = v32
        } else {
            for fi34 := 0; fi34 < 3; fi34++ {
                var v35 float64 = float64(-2.25)
                _ = fi34
                _ = v35
            }
            fmt.Println(p1, p1)
            p1 = (p1 % ((255) | 1))
        }
        _ = fi19
    }
    p0 -= (p0 << (((p0 >> ((1) & 63))) & 63))
    switch "foo bar" {
    case "a":
        v36 := []int{42, 32768, 32767}
        _ = v36
    case "\n\t":
        switch 32767 {
        case 9223372036854775807:
            for fi37 := 0; fi37 < 17; fi37++ {
                var v38 uint32 = 2
                _ = fi37
                _ = v38
            }
        default:
            v39 := []int{9223372036854775807, -9223372036854775808}
            _ = v39
        }
    default:
        p1 -= p1
    }
    _ = p0
    _ = p1
    return -32768
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[int]bool)
    var v3 string = (("a" + "a") + (("foo bar" + "hello") + "a"))
    var v4 int64 = 1
    v5 := make(map[uint32]bool)
    fmt.Println(v3, v1)
    v40 := v2[9223372036854775807]
    for fi41 := 0; fi41 < 9; fi41++ {
        var v42 bool = v40
        var v43 uint = 1
        fmt.Println(v42, fi41, v43)
        _ = fi41
        _ = v42
        _ = v43
    }
    if (255 >= 2) {
        fmt.Println(v3)
    }
    if false {
        v1 += 42
    }
    v44 := GenS4{F0: 42}
    fmt.Println(v40)
    var v45 int8 = 42
    v46 := GenS1{F0: -9223372036854775808}
    if (-788 == 1000) {
        for fi47 := 0; fi47 < 20; fi47++ {
            if v40 {
                switch ((v3 + "a") + (((v3 + "a") + v3) + "")) {
                case "hello":
                    v48 := GenS5{F0: v1, F1: 18446744073709551615}
                    _ = v48
                case "a":
                    if (v45 < 42) {
                        fmt.Println(v3, v40)
                    } else {
                        v40 = ((!(float64(1e10) != 1e10)) || ((3.14159 - ((1e10 - 0.0) / -2.25)) >= float32(((0.0 / (1e10 / 3.14159)) * 0.1))))
                        v49 := GenS4{F0: 962}
                        var v50 uint32 = 0
                        _ = v49
                        _ = v50
                    }
                case "\n\t":
                    switch v3 {
                    case "foo bar":
                        var v51 int16 = -83
                        _ = v51
                    case "a":
                        var v52 bool = false
                        _ = v52
                    case "\n\t":
                        if v40 {
                            var v53 float64 = float64(((1e10 + float32(((0.1 * -2.25) * 0.1))) * float32(0.0)))
                            v54 := []int{fi47}
                            switch -771 {
                            case -1:
                                v4 -= (v4 ^ (v4 % ((v4) | 1)))
                            default:
                                fmt.Println(fi47, fi47, v53)
                            }
                            _ = v53
                            _ = v54
                        } else {
                            v55 := GenS6{F0: 208}
                            _ = v55
                        }
                    default:
                        for fi56 := 0; fi56 < 14; fi56++ {
                            fmt.Println(fi47, v1, v40)
                            delete(v2, (fi47 | fi47))
                            var v57 int16 = 1
                            _ = fi56
                            _ = v57
                        }
                    }
                default:
                    if (float32(float64(float32(-2.25))) < 1.5) {
                        for fi58 := 0; fi58 < 4; fi58++ {
                            v1 -= 87
                            _ = fi58
                        }
                        for fi59 := 0; fi59 < 6; fi59++ {
                            v1 = 42
                            var v60 int8 = 42
                            _ = fi59
                            _ = v60
                        }
                        var v61 int32 = -2147483648
                        _ = v61
                    }
                }
            } else {
                v3 = (v3 + ("foo bar" + (v3 + "")))
                v62 := []int{}
                var v63 uint32 = v1
                _ = v62
                _ = v63
            }
            _ = fi47
        }
    } else {
        var v64 int32 = 42
        _ = v64
    }
    for fi65 := 0; fi65 < 7; fi65++ {
        v1 = v1
        var v66 uint32 = v1
        _ = fi65
        _ = v66
    }
    var v67 uint32 = (^(v1))
    v1 = (v1 << ((1) & 63))
    v67 = (v1 & 1)
    if (!((!(v4 > (v4 % ((v4) | 1)))) && v40)) {
        if v40 {
            v4 = v4
        }
        for fi68 := 0; fi68 < 2; fi68++ {
            if (!(v1 != v1)) {
                fmt.Println(v40, v3)
                var v69 int64 = (v4 | (v4 % ((v4) | 1)))
                switch fi68 {
                case 127:
                    v3 = ("\n\t" + "a")
                case -1:
                    v70 := []int{(^(fi68))}
                    _ = v70
                }
                _ = v69
            } else {
                v71 := []int{fi68}
                if (18446744073709551615 <= 65535) {
                    fmt.Println(v67)
                    for fi72 := 0; fi72 < 5; fi72++ {
                        v45 += (-(v45))
                        _ = fi72
                    }
                } else {
                    v40 = v40
                    var v73 int = (fi68 - fi68)
                    _ = v73
                }
                _ = v71
            }
            _ = fi68
        }
        v4 += v4
    } else {
        var v74 bool = false
        switch "a" {
        case "\n\t":
            v67 = 845
        case "hello":
            switch 365 {
            case 1:
                for fi75 := 0; fi75 < 2; fi75++ {
                    for fi76 := 0; fi76 < 5; fi76++ {
                        if (255 >= 32767) {
                            switch fi76 {
                            case 624:
                                v5[133] = v74
                            default:
                                var v77 int64 = v4
                                v74 = v74
                                _ = v77
                            }
                            if (v74 || false) {
                                v67 -= (v67 / ((v67) | 1))
                                v78 := []int{(fi76 % ((fi76) | 1)), -2147483648, fi75, (fi76 & fi76)}
                                var v79 float32 = (0.1 - (((0.1 / 1.0) + 0.0) - 1e10))
                                _ = v78
                                _ = v79
                            }
                            v45 -= 100
                            fmt.Println(v45)
                            v74 = v40
                        } else {
                            v67 += 281
                        }
                        v80 := GenS7{F0: v1, F1: 498}
                        if v40 {
                            var v81 int8 = (v45 + 100)
                            v74 = v74
                            _ = v81
                        } else {
                            if true {
                                var v82 uint64 = 1
                                _ = v82
                            } else {
                                if v74 {
                                    switch ("foo bar" + "a") {
                                    case "hello":
                                        if v74 {
                                            if v74 {
                                                v67 = v67
                                                v1 -= v67
                                                v83 := GenS1{F0: (fi76 * fi75)}
                                                _ = v83
                                            }
                                            v40 = v40
                                            fmt.Println(v1, v1)
                                        } else {
                                            var v84 bool = v40
                                            delete(v5, (v1 % ((1) | 1)))
                                            _ = v84
                                        }
                                    case "\n\t":
                                        v1 += (v67 ^ v1)
                                    case "":
                                        v85 := GenS2{F0: 1, F1: (fi75 - 9223372036854775807), F2: fi75}
                                        _ = v85
                                    default:
                                        if v40 {
                                            v86 := v5[v67]
                                            _ = v86
                                        }
                                    }
                                    var v87 int16 = genFn1(4294967295, 32767)
                                    _ = v87
                                } else {
                                    fmt.Println(fi75, fi75, v4)
                                }
                            }
                            v88 := GenS1{F0: fi76}
                            v89 := GenS2{F0: (v45 + v45), F1: 4294967295, F2: fi76}
                            if v40 {
                                if v40 {
                                    delete(v5, v67)
                                    fmt.Println(v1, v1, fi76)
                                } else {
                                    v40 = (848 >= int32(fi75))
                                }
                                var v90 int8 = v45
                                v91 := GenS5{F0: (v67 % (((^(v67))) | 1)), F1: 18446744073709551615}
                                if ((v4 | (^(v4))) >= -9223372036854775808) {
                                    fmt.Println(v74, v67)
                                    switch 42 {
                                    case 127:
                                        switch v3 {
                                        case "a":
                                            var v92 uint8 = 109
                                            _ = v92
                                        case "":
                                            if v40 {
                                                v93 := []int{(-(fi76))}
                                                _ = v93
                                            }
                                        default:
                                            if false {
                                                var v94 int32 = 1
                                                _ = v94
                                            }
                                        }
                                    case 42:
                                        switch fi76 {
                                        case 32768:
                                            var v95 bool = (18446744073709551615 > 436)
                                            _ = v95
                                        case -262:
                                            v96 := []int{fi76, int(v4)}
                                            _ = v96
                                        case 9223372036854775807:
                                            if v40 {
                                                var v97 int8 = (-(v90))
                                                _ = v97
                                            }
                                        }
                                    case 1:
                                        var v98 uint8 = 127
                                        _ = v98
                                    }
                                } else {
                                    v99 := GenS5{F0: v1, F1: uint64(v67)}
                                    v100 := GenS4{F0: 123}
                                    if ((1e10 - 0.0) <= ((-2.25 / 0.1) - (3.14159 * (1.5 / ((float64((3.14159 * 3.14159)) - 1.5) * float64((1.5 + ((((3.14159 * (-2.25 * float32((float64(0.0) - -2.25)))) - -2.25) + ((((1e10 / 0.1) * ((-2.25 + 3.14159) + 1e10)) / 1.0) - 1.5)) - 0.0)))))))) {
                                        if false {
                                            v101 := []int{9223372036854775807, (fi75 & 0)}
                                            v3 = v3
                                            v102 := GenS1{F0: fi76}
                                            var v103 uint8 = 128
                                            _ = v101
                                            _ = v102
                                            _ = v103
                                        } else {
                                            v104 := []int{382, fi76, fi75}
                                            _ = v104
                                        }
                                    }
                                    if (787 >= -1) {
                                        fmt.Println(v4, v45)
                                        switch (fi76 + fi75) {
                                        case 127:
                                            var v105 uint32 = v1
                                            _ = v105
                                        case 4294967295:
                                            var v106 string = "hello"
                                            _ = v106
                                        case -2147483648:
                                            switch v3 {
                                            case "foo bar":
                                                switch 2147483647 {
                                                case 4294967295:
                                                    if (!true) {
                                                        v40 = ((int64(v90) <= (v4 * v4)) && (4294967295 > 197))
                                                        var v107 uint64 = 1
                                                        v67 += v67
                                                        var v108 string = v3
                                                        _ = v107
                                                        _ = v108
                                                    }
                                                default:
                                                    fmt.Println(v74, v67)
                                                }
                                            case "":
                                                switch (v3 + (v3 + v3)) {
                                                case "\n\t":
                                                    if false {
                                                        switch (fi76 | (fi75 / ((fi76) | 1))) {
                                                        case 9223372036854775807:
                                                            switch (fi76 - fi75) {
                                                            case 68:
                                                                switch v3 {
                                                                case "":
                                                                    v40 = true
                                                                }
                                                            default:
                                                                var v109 uint16 = 2
                                                                _ = v109
                                                            }
                                                        case -1:
                                                            if false {
                                                                v110 := v2[0]
                                                                _ = v110
                                                            } else {
                                                                v111 := GenS8{F0: -2147483648, F1: 0}
                                                                v45 = v45
                                                                v3 = (v3 + (v3 + v3))
                                                                _ = v111
                                                            }
                                                            fmt.Println(v4)
                                                            v67 += 233
                                                        }
                                                        v112 := GenS6{F0: 224}
                                                        var v113 uint16 = 1
                                                        _ = v112
                                                        _ = v113
                                                    }
                                                default:
                                                    switch v3 {
                                                    case "foo bar":
                                                        switch fi76 {
                                                        case -634:
                                                            v45 = 127
                                                        case 128:
                                                            v114 := GenS9{F0: 239}
                                                            _ = v114
                                                        }
                                                    default:
                                                        v115 := GenS4{F0: 42}
                                                        _ = v115
                                                    }
                                                }
                                            case "hello":
                                                if false {
                                                    switch 127 {
                                                    case -373:
                                                        var v116 uint64 = 0
                                                        _ = v116
                                                    case -140:
                                                        v117 := GenS10{F0: 18446744073709551615, F1: (v4 * v4)}
                                                        _ = v117
                                                    default:
                                                        v118 := GenS8{F0: 692, F1: 1}
                                                        _ = v118
                                                    }
                                                }
                                            }
                                        default:
                                            var v119 int32 = -340
                                            _ = v119
                                        }
                                        var v120 uint16 = 623
                                        _ = v120
                                    }
                                    var v121 int8 = 127
                                    _ = v100
                                    _ = v121
                                    _ = v99
                                }
                                _ = v90
                                _ = v91
                            }
                            switch (v3 + "foo bar") {
                            case "foo bar":
                                fmt.Println(v1, v3, fi75)
                            case "\n\t":
                                var v122 int64 = v4
                                _ = v122
                            case "":
                                if v40 {
                                    v1 = (v1 * 199)
                                }
                            }
                            _ = v88
                            _ = v89
                        }
                        _ = fi76
                        _ = v80
                    }
                    for fi123 := 0; fi123 < 12; fi123++ {
                        if true {
                            switch fi75 {
                            case 128:
                                fmt.Println(fi123)
                            case 32767:
                                switch fi75 {
                                case 128:
                                    _ = len(v5)
                                case 1:
                                    fmt.Println(v3, v45, fi123)
                                case 0:
                                    switch (fi75 + (fi75 / ((fi123) | 1))) {
                                    case -2147483648:
                                        v124 := v5[(v1 - 0)]
                                        _ = v124
                                    }
                                }
                            case 654:
                                if (992 <= 42) {
                                    if (42 >= 1) {
                                        if true {
                                            switch 401 {
                                            case 443:
                                                var v125 uint16 = 0
                                                _ = v125
                                            case 42:
                                                var v126 uint32 = v67
                                                _ = v126
                                            case 461:
                                                v40 = (!v40)
                                            default:
                                                var v127 int64 = (^(v4))
                                                _ = v127
                                            }
                                            switch (fi123 >> ((1) & 63)) {
                                            case 32767:
                                                fmt.Println(v74, v4, v40)
                                                if (false && v40) {
                                                    if v40 {
                                                        v128 := GenS9{F0: 255}
                                                        fmt.Println(v40)
                                                        _ = v128
                                                    }
                                                } else {
                                                    v129 := GenS11{F0: v67, F1: (-(v4))}
                                                    var v130 bool = v74
                                                    _ = v129
                                                    _ = v130
                                                }
                                            case -1:
                                                delete(v2, 42)
                                            case 1:
                                                delete(v2, 32768)
                                                if true {
                                                    v45 += (v45 / ((v45) | 1))
                                                }
                                            default:
                                                fmt.Println(v1)
                                            }
                                            var v131 int = 42
                                            _ = v131
                                        }
                                        switch 1 {
                                        case -9223372036854775808:
                                            var v132 uint = 255
                                            _ = v132
                                        case 32767:
                                            var v133 uint16 = 2
                                            _ = v133
                                        default:
                                            var v134 int16 = 32767
                                            _ = v134
                                        }
                                    }
                                    if true {
                                        v135 := []int{fi75}
                                        _ = v135
                                    } else {
                                        v136 := GenS12{F0: 42, F1: 2, F2: v3}
                                        _ = v136
                                    }
                                } else {
                                    v137 := []int{(-(fi75)), (fi123 * (fi123 & fi123)), (fi75 - 4294967295)}
                                    v40 = v40
                                    switch (fi123 >> ((1) & 63)) {
                                    case -869:
                                        if v40 {
                                            v138 := GenS3{F0: genFn1(2, 426), F1: 737, F2: 2147483647}
                                            if v40 {
                                                var v139 int64 = (v4 | -628)
                                                _ = v139
                                            }
                                            fmt.Println(v67)
                                            var v140 int8 = (v45 << ((568) & 63))
                                            if len(v137) > 0 {
                                                v137[0] = fi123
                                            }
                                            _ = v138
                                            _ = v140
                                        }
                                    case 0:
                                        v4 = 9223372036854775807
                                    default:
                                        if ((v3 + ("\n\t" + (v3 + v3))) >= (v3 + ((v3 + "foo bar") + "a"))) {
                                            if v40 {
                                                if (4294967295 >= 42) {
                                                    switch 4294967295 {
                                                    case 1:
                                                        if true {
                                                            v4 += int64(v67)
                                                            var v141 int = (fi75 >> ((0) & 63))
                                                            if false {
                                                                v67 += v1
                                                                var v142 float64 = 0.1
                                                                var v143 string = v3
                                                                _ = v142
                                                                _ = v143
                                                            }
                                                            _ = v141
                                                        } else {
                                                            v144 := GenS11{F0: (v67 / ((2147483648) | 1)), F1: v4}
                                                            _ = v144
                                                        }
                                                    case -976:
                                                        v145 := GenS7{F0: v1, F1: 520}
                                                        _ = v145
                                                    }
                                                }
                                                if false {
                                                    v146 := GenS7{F0: (v67 & v67), F1: 18446744073709551615}
                                                    v147 := GenS8{F0: 2147483647, F1: 230}
                                                    _ = v146
                                                    _ = v147
                                                }
                                                v148 := GenS1{F0: (-(fi123))}
                                                _ = v148
                                            } else {
                                                if v74 {
                                                    switch 4294967295 {
                                                    case 2147483647:
                                                        v45 = (v45 ^ (v45 * v45))
                                                    case -664:
                                                        if len(v137) > 0 {
                                                            v137[0] = (fi75 << ((4294967295) & 63))
                                                        }
                                                    }
                                                } else {
                                                    var v149 int16 = 668
                                                    if true {
                                                        var v150 int64 = 0
                                                        v151 := GenS9{F0: 236}
                                                        _ = v150
                                                        _ = v151
                                                    } else {
                                                        var v152 bool = (false && v74)
                                                        _ = v152
                                                    }
                                                    var v153 uint32 = (^(v1))
                                                    var v154 float32 = ((-2.25 / (-2.25 + (3.14159 * 3.14159))) - 1e10)
                                                    _ = v149
                                                    _ = v153
                                                    _ = v154
                                                }
                                            }
                                            if v74 {
                                                if (int64(v67) < v4) {
                                                    if v74 {
                                                        v155 := GenS5{F0: v67, F1: genFn0()}
                                                        switch v3 {
                                                        case "foo bar":
                                                            if v74 {
                                                                v156 := GenS4{F0: 4294967295}
                                                                var v157 int32 = 1
                                                                var v158 int16 = int16(v157)
                                                                if v40 {
                                                                    v159 := GenS13{F0: (v158 + (v158 + v158))}
                                                                    var v160 int32 = (v157 - (v157 | (v157 >> ((18446744073709551615) & 63))))
                                                                    _ = v159
                                                                    _ = v160
                                                                }
                                                                switch -2147483648 {
                                                                case -9223372036854775808:
                                                                    v45 = 0
                                                                    v74 = (0 < (v4 & 42))
                                                                    fmt.Println(v45)
                                                                }
                                                                _ = v156
                                                                _ = v157
                                                                _ = v158
                                                            }
                                                        default:
                                                            switch (fi123 * fi123) {
                                                            case 32768:
                                                                switch v3 {
                                                                case "hello":
                                                                    var v161 int16 = 0
                                                                    _ = v161
                                                                case "":
                                                                    switch (fi75 | 2147483647) {
                                                                    case 2147483647:
                                                                        v4 += 1
                                                                    }
                                                                case "foo bar":
                                                                    switch v3 {
                                                                    case "":
                                                                        var v162 uint32 = v1
                                                                        _ = v162
                                                                    }
                                                                }
                                                            case -1:
                                                                fmt.Println(fi123)
                                                            case -9223372036854775808:
                                                                var v163 uint = 4294967295
                                                                var v164 int32 = 2147483647
                                                                var v165 int64 = v4
                                                                _ = v163
                                                                _ = v164
                                                                _ = v165
                                                            }
                                                        }
                                                        _ = v155
                                                    }
                                                    var v166 int32 = 281
                                                    if len(v137) > 0 {
                                                        v137[0] = 0
                                                    }
                                                    _ = v166
                                                }
                                            }
                                        }
                                    }
                                    _ = v137
                                }
                            default:
                                fmt.Println(v3)
                            }
                            if (false && true) {
                                _ = len(v5)
                            }
                            switch fi75 {
                            case 2147483647:
                                v1 = v1
                            }
                        }
                        var v167 uint = 18446744073709551615
                        v168 := GenS2{F0: v45, F1: (^(fi123)), F2: (-(fi75))}
                        _ = fi123
                        _ = v167
                        _ = v168
                    }
                    for fi169 := 0; fi169 < 8; fi169++ {
                        v170 := GenS5{F0: 2147483647, F1: 42}
                        _ = fi169
                        _ = v170
                    }
                    _ = fi75
                }
            case 266:
                if (v40 || (v4 > v4)) {
                    var v171 int32 = 1
                    _ = v171
                } else {
                    for fi172 := 0; fi172 < 13; fi172++ {
                        if v40 {
                            v3 = v3
                            var v173 float64 = 3.14159
                            var v174 uint = 63
                            _ = v173
                            _ = v174
                        }
                        _ = fi172
                    }
                }
            default:
                _ = len(v2)
            }
        case "":
            v45 = (v45 ^ v45)
        }
        _ = v74
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v67
}
