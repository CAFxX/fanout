package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 uint64; F2 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int64; F1 int8 }
type GenS9 struct { F0 int32; F1 int64; F2 int }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS12 struct { F0 string; F1 uint8 }
type GenS13 struct { F0 uint16; F1 uint32 }
type GenS14 struct { F0 int; F1 int16 }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 uint8; F1 int16; F2 uint }
type GenS20 struct { F0 uint }
type GenS21 struct { F0 int32; F1 bool }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 int64 }
func genFn0(p0 int) uint64 {
    v5 := GenS1{F0: 42}
    v6 := GenS1{F0: 0}
    p0 += 9223372036854775807
    _ = p0
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int8, p1 uint8, p2 int16) int16 {
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            if false {
                switch 4294967295 {
                case -9223372036854775808:
                    v9 := []int{fi8, (fi7 >> ((255) & 63)), (fi8 ^ -1), fi7}
                    _ = v9
                case 825:
                    var v10 float64 = (1e10 / (3.14159 / (-2.25 + 1e10)))
                    _ = v10
                }
                v11 := []int{}
                _ = v11
            } else {
                v12 := []int{fi8, fi7, (^(fi7)), 989}
                _ = v12
            }
            v13 := GenS1{F0: (p1 + 245)}
            fmt.Println(fi7)
            _ = fi8
            _ = v13
        }
        switch 0 {
        case -1:
            v14 := GenS2{F0: (p0 ^ 39)}
            _ = v14
        }
        _ = fi7
    }
    if (!((-2.25 >= (0.0 + (0.1 - (0.1 * 3.14159)))) || true)) {
        v15 := GenS2{F0: p0}
        if true {
            v16 := []int{231}
            _ = v16
        }
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 uint8 = 125
    var v2 int8 = 100
    v3 := []string{("foo bar" + "a")}
    v4 := []int{-2147483648, 2147483647}
    if false {
        var v17 uint8 = 38
        fmt.Println(v1, v1)
        _ = v17
    } else {
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 uint64 = 1
            _ = fi18
            _ = v19
        }
        fmt.Println(v2, v2)
        var v20 float32 = (3.14159 + -2.25)
        _ = v20
    }
    var v21 float32 = (((-2.25 + (0.0 - (0.1 - 0.1))) / 1e10) + -2.25)
    v22 := GenS3{F0: 42, F1: uint64(v1), F2: true}
    if (1 < 42) {
        fmt.Println(v21, v21)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        switch -9223372036854775808 {
        case 4294967295:
            if (((0.0 + ((3.14159 - 0.0) * 0.1)) < ((1.5 / 1.0) - ((1e10 * (1e10 * 1.5)) / (1e10 + -2.25)))) && (v21 <= (0.0 * float32(0.1)))) {
                if true {
                    var v23 int32 = -1
                    _ = v23
                }
            } else {
                v21 += (float32(1e10) - v21)
                if (!false) {
                    var v24 int = 127
                    for fi25 := 0; fi25 < 8; fi25++ {
                        var v26 uint16 = 1
                        fmt.Println(v1, v2, v2)
                        for fi27 := 0; fi27 < 8; fi27++ {
                            if (((v26 & v26) != v26) && (("foo bar" + "hello") <= "\n\t")) {
                                v28 := GenS4{F0: (fi25 + (fi25 & 0))}
                                var v29 float64 = 3.14159
                                var v30 int8 = v2
                                v29 += v29
                                _ = v28
                                _ = v29
                                _ = v30
                            } else {
                                fmt.Println(fi27)
                                switch "foo bar" {
                                case "a":
                                    var v31 int64 = 42
                                    _ = v31
                                default:
                                    fmt.Println(v2, fi25, fi25)
                                }
                                if len(v4) > 0 {
                                    v4[0] = fi25
                                }
                                fmt.Println(v26)
                                var v32 uint64 = 1
                                _ = v32
                            }
                            v2 = (v2 << ((42) & 63))
                            v24 = fi25
                            _ = fi27
                        }
                        _ = fi25
                        _ = v26
                    }
                    v24 = 1
                    _ = v24
                }
                v2 = v2
            }
        case 32767:
            switch 0 {
            case 9223372036854775807:
                fmt.Println(v21, v2)
            case 0:
                v1 = v1
            }
        default:
            switch 32768 {
            case 158:
                v33 := GenS4{F0: -164}
                _ = v33
            case -2147483648:
                if (-1 != 910) {
                    fmt.Println(v1, v21)
                    fmt.Println(v21)
                    if ("hello" > "\n\t") {
                        var v34 uint32 = 2147483647
                        v21 = v21
                        v1 += (v1 + (v1 * (^(v1))))
                        _ = v34
                    }
                } else {
                    if false {
                        v2 -= v2
                    }
                }
            }
        }
    }
    for fi35 := 0; fi35 < 20; fi35++ {
        for fi36 := 0; fi36 < 19; fi36++ {
            switch "a" {
            case "":
                v2 = v2
            case "a":
                switch "a" {
                case "\n\t":
                    var v37 bool = true
                    _ = v37
                case "hello":
                    if false {
                        if (-128 != v2) {
                            v38 := GenS3{F0: -2147483648, F1: 1, F2: true}
                            var v39 uint16 = 65535
                            var v40 uint16 = (^(v39))
                            v1 += (v1 - (^(v1)))
                            v1 = (v1 << ((2) & 63))
                            _ = v38
                            _ = v39
                            _ = v40
                        } else {
                            if false {
                                switch ("hello" + "hello") {
                                case "foo bar":
                                    if true {
                                        if false {
                                            if (!(false || true)) {
                                                v41 := GenS5{F0: "\n\t"}
                                                var v42 float64 = 1.5
                                                v1 += v1
                                                _ = v41
                                                _ = v42
                                            }
                                            v1 -= v1
                                            var v43 int8 = v2
                                            if (v21 >= v21) {
                                                if true {
                                                    v44 := GenS6{F0: 0, F1: "hello"}
                                                    var v45 float64 = 3.14159
                                                    _ = v44
                                                    _ = v45
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                fmt.Println(fi35, fi36)
                                                if false {
                                                    v43 = (^(v2))
                                                } else {
                                                    if (-2.25 != -2.25) {
                                                        var v46 int32 = -2147483648
                                                        _ = v46
                                                    }
                                                    fmt.Println(v43, fi36, v21)
                                                }
                                            }
                                            _ = v43
                                        }
                                    }
                                case "\n\t":
                                    var v47 float64 = ((1.5 + ((-2.25 * (1.5 + (0.1 * 1.5))) / (1e10 / -2.25))) + -2.25)
                                    _ = v47
                                }
                                var v48 uint = 1
                                _ = v48
                            }
                            fmt.Println(v21, v2, v1)
                            if (1e10 != (0.0 - ((float64(3.14159) - 0.0) * 3.14159))) {
                                var v49 uint8 = 217
                                _ = v49
                            } else {
                                var v50 int64 = 490
                                v50 += (v50 + -55)
                                if true {
                                    switch ("\n\t" + ("a" + ((("a" + ("\n\t" + ("\n\t" + "hello"))) + "\n\t") + (("" + "hello") + "a")))) {
                                    case "hello":
                                        v51 := GenS7{F0: false}
                                        _ = v51
                                    case "\n\t":
                                        switch fi36 {
                                        case 42:
                                            if false {
                                                fmt.Println(v50, v21)
                                            } else {
                                                if (false && (true && (913 == 579))) {
                                                    v52 := GenS8{F0: v50, F1: 0}
                                                    _ = v52
                                                } else {
                                                    v2 -= v2
                                                }
                                                v3 = append(v3, "\n\t")
                                                v2 = (v2 ^ v2)
                                            }
                                        }
                                    case "a":
                                        v53 := GenS9{F0: 1, F1: v50, F2: fi35}
                                        _ = v53
                                    }
                                } else {
                                    if false {
                                        switch ("" + "foo bar") {
                                        case "a":
                                            v1 = (v1 / (((v1 - v1)) | 1))
                                        case "":
                                            var v54 float64 = ((3.14159 - 3.14159) - float64(v21))
                                            _ = v54
                                        }
                                        if true {
                                            var v55 int16 = -32768
                                            if len(v4) > 0 {
                                                v4[0] = fi35
                                            }
                                            v1 += v1
                                            _ = v55
                                        }
                                        v2 = (v2 * (v2 + v2))
                                        var v56 float32 = v21
                                        var v57 uint32 = 0
                                        _ = v56
                                        _ = v57
                                    }
                                }
                                if false {
                                    switch "hello" {
                                    case "a":
                                        switch "hello" {
                                        case "a":
                                            var v58 uint64 = 446
                                            _ = v58
                                        default:
                                            var v59 uint64 = 496
                                            _ = v59
                                        }
                                    }
                                } else {
                                    if true {
                                        fmt.Println(v50)
                                        if len(v3) > 0 {
                                            v3[0] = (("hello" + ("\n\t" + ((("" + "foo bar") + "\n\t") + (((("foo bar" + "a") + ((("foo bar" + "\n\t") + "\n\t") + ("hello" + "\n\t"))) + "foo bar") + ("\n\t" + "\n\t"))))) + "")
                                        }
                                    }
                                    switch ("foo bar" + "a") {
                                    case "foo bar":
                                        v60 := GenS6{F0: 471, F1: "foo bar"}
                                        _ = v60
                                    case "":
                                        v50 = v50
                                    case "\n\t":
                                        var v61 int64 = 42
                                        _ = v61
                                    }
                                    v50 -= v50
                                    if false {
                                        switch "" {
                                        case "\n\t":
                                            fmt.Println(fi36, fi36, fi36)
                                        case "hello":
                                            var v62 uint16 = 42
                                            _ = v62
                                        case "":
                                            v4 = append(v4, fi35)
                                        default:
                                            var v63 uint64 = 981
                                            _ = v63
                                        }
                                    } else {
                                        if false {
                                            v64 := GenS6{F0: 18446744073709551615, F1: "\n\t"}
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = v64
                                        }
                                        if ((v50 != v50) && (927 == 207)) {
                                            if len(v3) > 0 {
                                                v3[0] = "\n\t"
                                            }
                                            fmt.Println(v21)
                                        }
                                        var v65 int = (fi35 | 871)
                                        _ = v65
                                    }
                                    if false {
                                        if true {
                                            var v66 int32 = 42
                                            if (65535 >= 2147483647) {
                                                v67 := GenS7{F0: true}
                                                if len(v3) > 0 {
                                                    v3[0] = ""
                                                }
                                                _ = v67
                                            } else {
                                                var v68 uint64 = 1
                                                fmt.Println(v21, v50)
                                                var v69 bool = true
                                                _ = v68
                                                _ = v69
                                            }
                                            _ = v66
                                        } else {
                                            var v70 uint = 464
                                            v21 -= v21
                                            v71 := GenS10{F0: ((float64(v21) / -2.25) < float64((v21 - 0.1)))}
                                            _ = v70
                                            _ = v71
                                        }
                                        if false {
                                            switch "" {
                                            case "":
                                                _ = len(v4)
                                            case "hello":
                                                var v72 float32 = (-2.25 / (1e10 - 3.14159))
                                                _ = v72
                                            }
                                            fmt.Println(fi36, v21)
                                            var v73 bool = (true && false)
                                            v2 = -107
                                            v21 = -2.25
                                            _ = v73
                                        } else {
                                            v21 = v21
                                            switch "\n\t" {
                                            case "foo bar":
                                                if (false || (!(!(855 < 894)))) {
                                                    if (13 != 2147483647) {
                                                        if false {
                                                            fmt.Println(fi36, v2)
                                                        } else {
                                                            v74 := GenS11{F0: v2, F1: 42, F2: 42}
                                                            if false {
                                                                v1 = (v1 / (((v1 / ((v1) | 1))) | 1))
                                                                v1 = v1
                                                            }
                                                            v75 := GenS12{F0: "", F1: v1}
                                                            var v76 uint = 257
                                                            var v77 uint16 = 1
                                                            var v78 uint8 = v1
                                                            switch fi35 {
                                                            case 947:
                                                                switch fi35 {
                                                                case -1:
                                                                    if (true || ((1.5 + (-2.25 / 1.0)) == float64(v21))) {
                                                                        fmt.Println(v77, v21)
                                                                        var v79 uint32 = 126
                                                                        var v80 uint64 = 0
                                                                        _ = v79
                                                                        _ = v80
                                                                    } else {
                                                                        v21 = v21
                                                                        v1 = (v78 + (v78 * v1))
                                                                        v81 := GenS13{F0: v77, F1: 4294967295}
                                                                        _ = v81
                                                                    }
                                                                default:
                                                                    if true {
                                                                        switch "hello" {
                                                                        case "\n\t":
                                                                            v21 += 3.14159
                                                                        case "foo bar":
                                                                            var v82 int8 = (v2 | 1)
                                                                            _ = v82
                                                                        }
                                                                        var v83 uint64 = 42
                                                                        var v84 int16 = 1
                                                                        _ = v83
                                                                        _ = v84
                                                                    }
                                                                }
                                                                v2 = v2
                                                            case -9223372036854775808:
                                                                var v85 uint64 = genFn0((fi35 >> (((v76 / ((v76) | 1))) & 63)))
                                                                _ = v85
                                                            case -655:
                                                                var v86 uint = v76
                                                                _ = v86
                                                            default:
                                                                var v87 int64 = -1
                                                                _ = v87
                                                            }
                                                            var v88 uint32 = uint32(v76)
                                                            _ = v74
                                                            _ = v75
                                                            _ = v76
                                                            _ = v77
                                                            _ = v78
                                                            _ = v88
                                                        }
                                                    } else {
                                                        v89 := GenS14{F0: (fi36 ^ fi35), F1: -32768}
                                                        fmt.Println(fi36, v1)
                                                        _ = v89
                                                    }
                                                    var v90 uint16 = 0
                                                    if true {
                                                        var v91 string = ""
                                                        _ = v91
                                                    }
                                                    _ = v90
                                                }
                                            default:
                                                v92 := GenS4{F0: 4294967295}
                                                switch "hello" {
                                                case "foo bar":
                                                    v21 = v21
                                                default:
                                                    v50 += 1
                                                }
                                                v93 := GenS5{F0: "\n\t"}
                                                _ = v92
                                                _ = v93
                                            }
                                            v50 -= (v50 - -9223372036854775808)
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                }
                                _ = v50
                            }
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v94 int8 = (v2 << ((65535) & 63))
                        _ = v94
                    }
                case "":
                    v95 := GenS12{F0: "hello", F1: 127}
                    v1 -= 2
                    _ = v95
                default:
                    fmt.Println(v1, fi35, fi36)
                }
            }
            switch (fi36 * 32767) {
            case -19:
                var v96 float64 = -2.25
                _ = v96
            case -469:
                fmt.Println(v2)
                var v97 int16 = -515
                var v98 float32 = ((float32(1e10) + v21) + (v21 + v21))
                _ = v97
                _ = v98
            }
            if true {
                v1 += (v1 / ((v1) | 1))
                var v99 int = fi35
                if len(v3) > 0 {
                    v3[0] = ("a" + ("" + "\n\t"))
                }
                _ = v99
            }
            _ = fi36
        }
        v100 := GenS6{F0: 0, F1: "hello"}
        switch ("\n\t" + "\n\t") {
        case "hello":
            var v101 uint16 = 32767
            _ = v101
        case "\n\t":
            v21 -= ((float32((float64(1.5) - 0.1)) - ((3.14159 - v21) * (3.14159 - (((0.1 - ((v21 - (v21 + 3.14159)) * (v21 + (v21 / (v21 / v21))))) - v21) + v21)))) + v21)
        case "":
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if (4294967295 == 65535) {
                    var v102 int32 = 2147483647
                    switch ("" + ("" + "foo bar")) {
                    case "foo bar":
                        switch ("\n\t" + "foo bar") {
                        case "a":
                            var v103 uint16 = 810
                            _ = v103
                        case "foo bar":
                            var v104 uint8 = v1
                            _ = v104
                        }
                    }
                    _ = v102
                }
            } else {
                if (!false) {
                    if (((0.0 + -2.25) * 3.14159) != 3.14159) {
                        var v105 int32 = 1
                        _ = v105
                    }
                    for fi106 := 0; fi106 < 9; fi106++ {
                        switch (fi35 ^ fi106) {
                        case -9223372036854775808:
                            switch (-(fi35)) {
                            case -136:
                                v1 = v1
                            }
                        case 4294967295:
                            v1 = (v1 | 2)
                        case 32767:
                            if false {
                                v2 += v2
                                v21 += v21
                                if (-9223372036854775808 != 9223372036854775807) {
                                    if true {
                                        v1 -= (^(v1))
                                        if ((("\n\t" + "hello") == "hello") || ((^(v1)) >= (v1 ^ v1))) {
                                            var v107 int8 = (^(v2))
                                            var v108 uint16 = 223
                                            _ = v107
                                            _ = v108
                                        }
                                    }
                                }
                                v109 := GenS7{F0: false}
                                _ = v109
                            }
                        default:
                            if (v21 < v21) {
                                v1 = v1
                            }
                        }
                        v1 += (v1 >> ((1) & 63))
                        v110 := GenS7{F0: true}
                        var v111 uint64 = 2
                        _ = fi106
                        _ = v110
                        _ = v111
                    }
                }
                if len(v4) > 0 {
                    v4[0] = (-(fi35))
                }
            }
        default:
            if false {
                v1 += (^(v1))
                var v112 uint64 = 2
                _ = v112
            }
        }
        _ = fi35
        _ = v100
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    v1 = v1
    var v113 int32 = 731
    var v114 uint8 = v1
    for fi115 := 0; fi115 < 5; fi115++ {
        v116 := GenS15{F0: (v113 * v113)}
        _ = fi115
        _ = v116
    }
    var v117 float32 = 0.1
    switch "a" {
    case "\n\t":
        var v118 uint64 = genFn0(int(v1))
        _ = v118
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = ("hello" + "\n\t")
        }
    default:
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    }
    v2 = (v2 & 1)
    v2 -= v2
    if false {
        for fi119 := 0; fi119 < 12; fi119++ {
            var v120 bool = true
            v2 += 127
            _ = fi119
            _ = v120
        }
        if (-1 < 109) {
            v121 := GenS8{F0: 42, F1: 0}
            if (true || true) {
                fmt.Println(v21, v114, v117)
            }
            if (2 < 61) {
                if (!true) {
                    v2 -= (v2 + v2)
                    fmt.Println(v117, v21, v21)
                } else {
                    for fi122 := 0; fi122 < 11; fi122++ {
                        var v123 float32 = 0.1
                        for fi124 := 0; fi124 < 5; fi124++ {
                            var v125 uint32 = 602
                            _ = fi124
                            _ = v125
                        }
                        _ = fi122
                        _ = v123
                    }
                    v126 := GenS16{F0: 197}
                    v113 = 797
                    _ = v126
                }
                _ = len(v3)
                v21 += 0.0
            }
            _ = v121
        } else {
            v117 = v117
        }
        v127 := GenS8{F0: 9223372036854775807, F1: int8(v113)}
        _ = v127
    } else {
        for fi128 := 0; fi128 < 12; fi128++ {
            for fi129 := 0; fi129 < 4; fi129++ {
                v130 := GenS12{F0: ("hello" + "a"), F1: v1}
                _ = fi129
                _ = v130
            }
            for fi131 := 0; fi131 < 13; fi131++ {
                switch (-(fi128)) {
                case 128:
                    v4 = append(v4, (-(fi131)))
                case -1:
                    v4 = append(v4, 32767)
                case 32768:
                    var v132 int8 = (v2 % (((-(v2))) | 1))
                    _ = v132
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                switch ("\n\t" + "hello") {
                case "":
                    var v133 float32 = v117
                    _ = v133
                default:
                    fmt.Println(v117, fi128)
                }
                _ = len(v4)
                _ = fi131
            }
            _ = fi128
        }
        var v134 float64 = (1.5 / 1.0)
        v2 = 0
        _ = v134
    }
    switch ("\n\t" + "") {
    case "a":
        if false {
            switch "a" {
            case "":
                if false {
                    switch 473 {
                    case 914:
                        if len(v4) > 0 {
                            v4[0] = 4294967295
                        }
                    case 34:
                        if (true && false) {
                            v113 = v113
                        } else {
                            var v135 uint64 = 745
                            v136 := GenS11{F0: 127, F1: 4294967295, F2: 32767}
                            _ = v135
                            _ = v136
                        }
                    default:
                        var v137 uint64 = 4294967295
                        _ = v137
                    }
                }
            case "foo bar":
                for fi138 := 0; fi138 < 4; fi138++ {
                    for fi139 := 0; fi139 < 11; fi139++ {
                        v140 := GenS17{F0: false}
                        v3 = append(v3, "")
                        _ = fi139
                        _ = v140
                    }
                    _ = fi138
                }
            case "a":
                v21 = v21
            }
            v141 := GenS2{F0: (v2 * -128)}
            _ = v141
        }
    case "hello":
        var v142 float32 = -2.25
        _ = v142
    case "foo bar":
        if ("" != "\n\t") {
            v143 := GenS18{F0: false}
            if (0 == 4294967295) {
                v2 += (v2 * v2)
            }
            _ = v143
        } else {
            for fi144 := 0; fi144 < 3; fi144++ {
                var v145 int32 = 2147483647
                for fi146 := 0; fi146 < 4; fi146++ {
                    if (1 > 1) {
                        if false {
                            switch fi146 {
                            case -499:
                                switch "a" {
                                case "":
                                    v4 = append(v4, (fi146 - (fi146 / ((fi144) | 1))))
                                case "a":
                                    v147 := GenS5{F0: "a"}
                                    _ = v147
                                case "foo bar":
                                    var v148 int32 = 1
                                    _ = v148
                                default:
                                    switch ("foo bar" + (("a" + "\n\t") + "hello")) {
                                    case "":
                                        var v149 uint64 = 1
                                        _ = v149
                                    }
                                }
                            case 4294967295:
                                var v150 string = "a"
                                _ = v150
                            }
                        } else {
                            v117 -= v21
                            fmt.Println(v114, fi144)
                        }
                        v151 := GenS5{F0: (("hello" + "hello") + ("foo bar" + ("foo bar" + "a")))}
                        if (1 >= 42) {
                            _ = len(v3)
                        } else {
                            if (((v2 % ((42) | 1)) >= -128) || (464 != 2147483647)) {
                                var v152 int64 = -1
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v153 := GenS19{F0: 42, F1: 43, F2: 18446744073709551615}
                                _ = v152
                                _ = v153
                            }
                            switch fi144 {
                            case 46:
                                var v154 uint64 = 425
                                var v155 uint = 782
                                _ = v154
                                _ = v155
                            case -1:
                                switch "hello" {
                                case "a":
                                    if false {
                                        if false {
                                            v2 -= v2
                                            fmt.Println(v145, fi146, v21)
                                            v114 = 255
                                        } else {
                                            var v156 uint = 896
                                            _ = v156
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = "foo bar"
                                        }
                                        var v157 uint16 = 268
                                        _ = v157
                                    }
                                default:
                                    switch "" {
                                    case "":
                                        var v158 uint16 = 1
                                        _ = v158
                                    case "hello":
                                        v21 = v117
                                    case "foo bar":
                                        v1 = v1
                                    }
                                }
                            case 495:
                                var v159 int32 = v113
                                _ = v159
                            default:
                                v160 := GenS15{F0: -867}
                                _ = v160
                            }
                        }
                        _ = v151
                    } else {
                        var v161 uint32 = 2
                        _ = v161
                    }
                    var v162 bool = true
                    _ = fi146
                    _ = v162
                }
                _ = fi144
                _ = v145
            }
            switch 4294967295 {
            case -52:
                var v163 int32 = 42
                _ = v163
            case -9223372036854775808:
                v1 -= v114
            case 128:
                v117 += 0.0
            default:
                var v164 int8 = -73
                _ = v164
            }
        }
    default:
        switch "a" {
        case "foo bar":
            switch -9223372036854775808 {
            case -2147483648:
                v3 = append(v3, "hello")
            case 42:
                var v165 int16 = 32767
                _ = v165
            case 32768:
                v166 := GenS20{F0: 216}
                _ = v166
            default:
                if (v113 <= v113) {
                    switch ("\n\t" + "foo bar") {
                    case "foo bar":
                        v117 = ((v117 / float32(1.5)) + v117)
                    case "":
                        for fi167 := 0; fi167 < 13; fi167++ {
                            _ = len(v3)
                            v168 := GenS21{F0: v113, F1: true}
                            var v169 uint16 = 0
                            _ = fi167
                            _ = v168
                            _ = v169
                        }
                    case "a":
                        switch 127 {
                        case -570:
                            for fi170 := 0; fi170 < 20; fi170++ {
                                var v171 int = 460
                                var v172 uint16 = uint16(v114)
                                _ = fi170
                                _ = v171
                                _ = v172
                            }
                        default:
                            fmt.Println(v1, v1, v2)
                        }
                    default:
                        for fi173 := 0; fi173 < 15; fi173++ {
                            if (723 <= 461) {
                                for fi174 := 0; fi174 < 7; fi174++ {
                                    v175 := GenS22{F0: (((v21 / (v117 - v21)) - 1.5) <= (v117 - 0.0))}
                                    v113 += v113
                                    _ = fi174
                                    _ = v175
                                }
                                v2 = (-(v2))
                            } else {
                                fmt.Println(v1, v117, v117)
                                switch (fi173 >> ((4294967295) & 63)) {
                                case -2147483648:
                                    var v176 string = ""
                                    _ = v176
                                case 42:
                                    if true {
                                        v177 := GenS1{F0: v114}
                                        var v178 int32 = v113
                                        if (42 <= 832) {
                                            v179 := GenS9{F0: (v178 + (v113 | v178)), F1: -1, F2: fi173}
                                            if false {
                                                v113 = -2147483648
                                            }
                                            _ = v179
                                        }
                                        _ = v177
                                        _ = v178
                                    } else {
                                        var v180 uint64 = 2
                                        _ = len(v3)
                                        _ = v180
                                    }
                                case 32768:
                                    for fi181 := 0; fi181 < 15; fi181++ {
                                        if true {
                                            var v182 int16 = 912
                                            _ = v182
                                        } else {
                                            fmt.Println(v113, v1, v114)
                                            v183 := GenS19{F0: (v1 ^ (v1 + (v114 << ((878) & 63)))), F1: -32768, F2: 378}
                                            var v184 int32 = 354
                                            _ = v183
                                            _ = v184
                                        }
                                        var v185 int8 = 127
                                        if len(v3) > 0 {
                                            v3[0] = "foo bar"
                                        }
                                        _ = fi181
                                        _ = v185
                                    }
                                }
                            }
                            v186 := GenS8{F0: -1, F1: (v2 ^ v2)}
                            _ = fi173
                            _ = v186
                        }
                    }
                    var v187 uint64 = 924
                    for fi188 := 0; fi188 < 7; fi188++ {
                        switch (-(fi188)) {
                        case 9223372036854775807:
                            if (false && false) {
                                switch "\n\t" {
                                case "a":
                                    if false {
                                        var v189 uint = 42
                                        var v190 int16 = 0
                                        _ = v189
                                        _ = v190
                                    }
                                case "\n\t":
                                    v187 = v187
                                }
                                v1 = v114
                                var v191 int32 = int32(v2)
                                _ = v191
                            } else {
                                fmt.Println(v1, v187)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        default:
                            v113 += v113
                        }
                        _ = fi188
                    }
                    _ = v187
                } else {
                    if true {
                        if true {
                            for fi192 := 0; fi192 < 5; fi192++ {
                                switch "hello" {
                                case "":
                                    switch 0 {
                                    case 98:
                                        var v193 uint8 = (v1 >> ((928) & 63))
                                        _ = v193
                                    default:
                                        _ = len(v4)
                                    }
                                default:
                                    switch fi192 {
                                    case 4294967295:
                                        var v194 string = ""
                                        _ = v194
                                    case 1:
                                        for fi195 := 0; fi195 < 3; fi195++ {
                                            var v196 int = (fi195 | fi192)
                                            v197 := GenS23{F0: 42}
                                            _ = fi195
                                            _ = v196
                                            _ = v197
                                        }
                                    default:
                                        v198 := GenS4{F0: 42}
                                        _ = v198
                                    }
                                }
                                fmt.Println(v117, fi192, v114)
                                var v199 uint16 = 0
                                _ = fi192
                                _ = v199
                            }
                        }
                        v200 := GenS1{F0: (^(v114))}
                        _ = v200
                    }
                }
            }
        case "a":
            fmt.Println(v113)
        case "":
            var v201 int8 = v2
            _ = v201
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("v114:", v114)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v113
    _ = v114
    _ = v117
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
