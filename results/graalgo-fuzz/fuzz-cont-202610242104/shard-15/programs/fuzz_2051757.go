package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 string; F2 uint }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 string; F1 int }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int16; F1 int8; F2 string }
type GenS11 struct { F0 int; F1 int8; F2 int64 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int32; F1 uint32 }
type GenS14 struct { F0 int64; F1 int16 }
type GenS15 struct { F0 uint8; F1 uint32 }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 int16; F1 uint64 }
func genFn0() string {
    var v6 int = 0
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: ((false && true) && false)}
        var v9 uint16 = 0
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: 1, F1: "a", F2: 2}
    if true {
        v11 := GenS1{F0: true}
        _ = v11
    }
    _ = v10
    _ = v6
    return "\n\t"
}
func genFn1() int8 {
    v12 := []int{127, -2147483648}
    var v13 int32 = 878
    for fi14 := 0; fi14 < 17; fi14++ {
        switch (fi14 >> ((517) & 63)) {
        case -2147483648:
            v15 := GenS3{F0: 2, F1: "\n\t"}
            _ = v15
        case 0:
            v16 := GenS2{F0: -1, F1: "hello", F2: 65535}
            _ = v16
        case 1:
            v17 := GenS4{F0: -28}
            _ = v17
        }
        if false {
            switch (^(fi14)) {
            case -187:
                v12 = append(v12, -2147483648)
            case 0:
                v13 -= (-(v13))
            default:
                var v18 int32 = v13
                _ = v18
            }
            switch (fi14 | fi14) {
            case -1:
                if (2147483648 >= 2147483648) {
                    v13 = (^(v13))
                    v12 = append(v12, fi14)
                    if (((!false) || true) || false) {
                        fmt.Println(fi14, fi14)
                    } else {
                        if (0 == 255) {
                            v13 = (-(v13))
                            switch ("\n\t" + ("a" + "")) {
                            case "\n\t":
                                var v19 string = "foo bar"
                                _ = v19
                            case "":
                                switch fi14 {
                                case 0:
                                    for fi20 := 0; fi20 < 7; fi20++ {
                                        if true {
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                            if (!(2147483648 != 285)) {
                                                if true {
                                                    v12 = append(v12, fi14)
                                                }
                                                v21 := GenS5{F0: -559}
                                                _ = v21
                                            }
                                            v13 = -2147483648
                                        } else {
                                            switch "hello" {
                                            case "foo bar":
                                                if len(v12) > 0 {
                                                    v12[0] = 933
                                                }
                                            default:
                                                fmt.Println(fi14)
                                            }
                                            var v22 uint = 255
                                            var v23 float64 = 1e10
                                            _ = v22
                                            _ = v23
                                        }
                                        _ = fi20
                                    }
                                default:
                                    v13 += 0
                                }
                            default:
                                var v24 string = "foo bar"
                                _ = v24
                            }
                            var v25 uint32 = 0
                            _ = v25
                        }
                    }
                }
            case 127:
                var v26 uint = 18446744073709551615
                _ = v26
            }
        } else {
            var v27 uint64 = 0
            var v28 string = "a"
            switch v28 {
            case "foo bar":
                var v29 string = "hello"
                _ = v29
            }
            _ = v27
            _ = v28
        }
        _ = fi14
    }
    switch 4294967295 {
    case -346:
        switch "a" {
        case "foo bar":
            fmt.Println(v13)
        case "":
            switch -122 {
            case -275:
                fmt.Println(v13)
            }
        default:
            switch -2147483648 {
            case 127:
                v13 -= v13
            case 709:
                if ((3.14159 * 1e10) < 1.5) {
                    if (4294967295 > 0) {
                        if (false && true) {
                            var v30 int32 = (v13 << ((0) & 63))
                            for fi31 := 0; fi31 < 4; fi31++ {
                                for fi32 := 0; fi32 < 7; fi32++ {
                                    switch (fi32 / (((fi32 >> ((0) & 63))) | 1)) {
                                    case 32767:
                                        if true {
                                            v12 = append(v12, fi32)
                                            var v33 int16 = int16(fi31)
                                            v34 := GenS6{F0: ("foo bar" + ("foo bar" + ("" + "foo bar"))), F1: 2}
                                            if false {
                                                var v35 float64 = 0.1
                                                v36 := GenS5{F0: 768}
                                                var v37 float64 = v35
                                                _ = v35
                                                _ = v36
                                                _ = v37
                                            }
                                            fmt.Println(fi31, fi31, fi31)
                                            _ = v33
                                            _ = v34
                                        } else {
                                            fmt.Println(fi32, fi32)
                                            var v38 string = "hello"
                                            _ = v38
                                        }
                                    default:
                                        v12 = append(v12, (fi31 << ((958) & 63)))
                                    }
                                    _ = fi32
                                }
                                v12 = append(v12, fi31)
                                fmt.Println(fi31, v13)
                                _ = fi31
                            }
                            _ = v30
                        }
                        switch ("hello" + "\n\t") {
                        case "foo bar":
                            for fi39 := 0; fi39 < 14; fi39++ {
                                switch 0 {
                                case -1:
                                    if false {
                                        for fi40 := 0; fi40 < 5; fi40++ {
                                            var v41 float32 = 1e10
                                            switch ("foo bar" + "") {
                                            case "\n\t":
                                                fmt.Println(fi40)
                                            }
                                            _ = fi40
                                            _ = v41
                                        }
                                        switch (-(fi39)) {
                                        case 32767:
                                            switch fi39 {
                                            case -9223372036854775808:
                                                fmt.Println(v13)
                                            case 128:
                                                var v42 float64 = 0.0
                                                _ = v42
                                            }
                                        case 9223372036854775807:
                                            v43 := GenS7{F0: "\n\t", F1: fi39}
                                            _ = v43
                                        case 32768:
                                            var v44 float32 = float32(-2.25)
                                            _ = v44
                                        default:
                                            v45 := GenS6{F0: "", F1: 2}
                                            _ = v45
                                        }
                                        v46 := GenS8{F0: fi39}
                                        _ = v46
                                    } else {
                                        switch "hello" {
                                        case "hello":
                                            for fi47 := 0; fi47 < 10; fi47++ {
                                                fmt.Println(v13)
                                                _ = fi47
                                            }
                                        case "foo bar":
                                            var v48 int64 = 9223372036854775807
                                            _ = v48
                                        default:
                                            fmt.Println(fi39)
                                        }
                                        v13 += v13
                                        if false {
                                            v13 = (v13 >> ((279) & 63))
                                            _ = len(v12)
                                            v13 = 302
                                        }
                                    }
                                case -733:
                                    if (float64((1.5 + 1.5)) <= float64(0.1)) {
                                        if false {
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                            if true {
                                                var v49 uint8 = 1
                                                _ = v49
                                            } else {
                                                if false {
                                                    var v50 uint64 = 0
                                                    _ = v50
                                                } else {
                                                    if len(v12) > 0 {
                                                        v12[0] = (fi39 >> ((65535) & 63))
                                                    }
                                                }
                                                v12 = append(v12, (-(fi39)))
                                                v51 := GenS9{F0: 1}
                                                _ = v51
                                            }
                                            v13 -= (v13 * int32(fi39))
                                        }
                                    }
                                case 32767:
                                    for fi52 := 0; fi52 < 4; fi52++ {
                                        v13 -= v13
                                        v53 := GenS9{F0: 451}
                                        _ = fi52
                                        _ = v53
                                    }
                                }
                                for fi54 := 0; fi54 < 6; fi54++ {
                                    if (723 > 287) {
                                        var v55 float64 = -2.25
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                        _ = v55
                                    } else {
                                        var v56 uint16 = 32767
                                        switch fi54 {
                                        case 2147483647:
                                            v13 = v13
                                        case -923:
                                            if true {
                                                switch 42 {
                                                case 630:
                                                    if len(v12) >= 2 {
                                                        v12 = v12[1:len(v12)]
                                                    }
                                                case 127:
                                                    switch (("" + ("a" + "a")) + "foo bar") {
                                                    case "":
                                                        var v57 float64 = 1e10
                                                        _ = v57
                                                    }
                                                case 258:
                                                    fmt.Println(v56)
                                                default:
                                                    fmt.Println(fi39)
                                                    v13 += (v13 ^ v13)
                                                    if (2 >= 485) {
                                                        v12 = append(v12, 467)
                                                        if len(v12) >= 2 {
                                                            v12 = v12[1:len(v12)]
                                                        }
                                                    } else {
                                                        if len(v12) > 0 {
                                                            v12[0] = (-(fi54))
                                                        }
                                                        if true {
                                                            fmt.Println(v56, v56)
                                                            if len(v12) >= 2 {
                                                                v12 = v12[1:len(v12)]
                                                            }
                                                        } else {
                                                            v12 = append(v12, fi54)
                                                            v12 = append(v12, 128)
                                                            v56 += (v56 & v56)
                                                        }
                                                        fmt.Println(fi39, fi54, v56)
                                                    }
                                                }
                                                v58 := GenS6{F0: "a", F1: 18446744073709551615}
                                                _ = v58
                                            }
                                        default:
                                            v56 -= 726
                                        }
                                        _ = v56
                                    }
                                    v13 = 984
                                    v13 = (v13 - (-(v13)))
                                    if ((690 == 255) || (!((!true) && true))) {
                                        v59 := GenS1{F0: (true && (("\n\t" + ("hello" + "a")) > ("" + ("" + "foo bar"))))}
                                        _ = v59
                                    }
                                    _ = fi54
                                }
                                var v60 int64 = 1
                                _ = fi39
                                _ = v60
                            }
                        case "\n\t":
                            v61 := GenS5{F0: -949}
                            _ = v61
                        case "":
                            _ = len(v12)
                        }
                    }
                    fmt.Println(v13)
                } else {
                    v13 -= 1
                    var v62 bool = false
                    for fi63 := 0; fi63 < 11; fi63++ {
                        fmt.Println(v13, v62, v13)
                        fmt.Println(v13, v62)
                        var v64 uint64 = 18446744073709551615
                        _ = fi63
                        _ = v64
                    }
                    _ = v62
                }
            default:
                v13 += (v13 << ((1) & 63))
            }
        }
    case -410:
        v13 -= (v13 % ((v13) | 1))
    default:
        _ = len(v12)
    }
    switch ("foo bar" + ("hello" + ("foo bar" + "a"))) {
    case "a":
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    default:
        v13 += 42
    }
    _ = v12
    _ = v13
    return -128
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]bool)
    var v3 float32 = 0.0
    v4 := make(map[int]int)
    v5 := make(map[string]int)
    switch 128 {
    case 42:
        var v65 int8 = -1
        _ = v65
    case -1:
        v3 -= 1.5
    }
    switch 127 {
    case 42:
        v3 = (v3 * 0.0)
    case -948:
        for fi66 := 0; fi66 < 9; fi66++ {
            fmt.Println(fi66, v3)
            _ = fi66
        }
    }
    var v67 int16 = -990
    var v68 int = 128
    v68 += (v68 ^ v68)
    v69 := GenS9{F0: 2}
    if ("\n\t" != "foo bar") {
        var v70 uint16 = 32767
        if false {
            switch "\n\t" {
            case "":
                for fi71 := 0; fi71 < 14; fi71++ {
                    var v72 uint8 = 0
                    _ = fi71
                    _ = v72
                }
            }
            for fi73 := 0; fi73 < 14; fi73++ {
                v3 = (0.1 - (v3 - ((v3 / v3) / ((1e10 - -2.25) / 3.14159))))
                _ = fi73
            }
        }
        _ = v70
    }
    if (-2.25 > v3) {
        if false {
            if false {
                v3 = 1e10
            } else {
                v74 := v1[(v68 & (v68 * (v68 | v68)))]
                for fi75 := 0; fi75 < 14; fi75++ {
                    if (262 <= 42) {
                        v76 := []int{-9223372036854775808, fi75, (fi75 | (fi75 | fi75)), v68}
                        for fi77 := 0; fi77 < 5; fi77++ {
                            switch v74 {
                            case "":
                                var v78 int16 = (v67 & (^(v67)))
                                if true {
                                    delete(v5, "hello")
                                }
                                var v79 int16 = (v78 / (((v67 / (((v67 / ((1) | 1))) | 1))) | 1))
                                _ = v78
                                _ = v79
                            }
                            _ = len(v76)
                            var v80 uint64 = 42
                            _ = fi77
                            _ = v80
                        }
                        _ = v76
                    } else {
                        v68 = -9223372036854775808
                        if false {
                            v67 += (^(v67))
                            for fi81 := 0; fi81 < 20; fi81++ {
                                v82 := GenS10{F0: v67, F1: 1, F2: (v74 + v74)}
                                _ = fi81
                                _ = v82
                            }
                        } else {
                            fmt.Println(v68)
                            v83 := v2[0]
                            var v84 int8 = 0
                            _ = v83
                            _ = v84
                        }
                    }
                    v68 += (-(v68))
                    _ = fi75
                }
                if false {
                    v85 := GenS5{F0: 9223372036854775807}
                    var v86 int64 = -9223372036854775808
                    for fi87 := 0; fi87 < 15; fi87++ {
                        var v88 int32 = -2147483648
                        for fi89 := 0; fi89 < 13; fi89++ {
                            if (float32(3.14159) == float32((0.0 * 3.14159))) {
                                var v90 string = ""
                                v91 := GenS11{F0: -2147483648, F1: -107, F2: v86}
                                var v92 uint64 = 1
                                _ = v90
                                _ = v91
                                _ = v92
                            } else {
                                var v93 uint16 = 2
                                v94 := GenS10{F0: (v67 & 0), F1: 1, F2: ("foo bar" + v74)}
                                if (929 == 542) {
                                    if (550 > 461) {
                                        switch (v68 % (((fi89 - 2147483647)) | 1)) {
                                        case 858:
                                            if true {
                                                delete(v1, 369)
                                                v67 = v67
                                            } else {
                                                fmt.Println(v67)
                                            }
                                        }
                                    } else {
                                        v95 := GenS9{F0: 808}
                                        fmt.Println(v67, v74, v74)
                                        v96 := GenS12{F0: 26}
                                        _ = v95
                                        _ = v96
                                    }
                                    v97 := GenS2{F0: v86, F1: v74, F2: 18446744073709551615}
                                    if (((3.14159 / v3) * 0.0) > v3) {
                                        var v98 float32 = v3
                                        if false {
                                            var v99 uint32 = 2147483647
                                            fmt.Println(fi89, v99, v86)
                                            var v100 float64 = 1e10
                                            v99 += (v99 & v99)
                                            _ = v100
                                            _ = v99
                                        } else {
                                            v3 -= 3.14159
                                            if (v88 < v88) {
                                                v3 += (0.1 + (0.1 / ((v3 * ((v98 / (-2.25 - (((1e10 * v3) + v3) / (v98 - (v3 - (float32(1.5) + float32((0.0 / (0.1 / 1.0))))))))) + 0.1)) / 1.0)))
                                                v101 := GenS13{F0: -2147483648, F1: 42}
                                                v93 += (v93 >> ((0) & 63))
                                                v102 := []int{32767, v68, fi89}
                                                _ = v101
                                                _ = v102
                                            }
                                        }
                                        v103 := []int{fi89, 32768, v68, (fi89 % (((fi89 / ((v68) | 1))) | 1))}
                                        if len(v103) >= 2 {
                                            v103 = v103[1:len(v103)]
                                        }
                                        _ = len(v4)
                                        _ = v103
                                        _ = v98
                                    } else {
                                        switch (fi87 * (fi89 >> ((685) & 63))) {
                                        case 42:
                                            if (4294967295 == 731) {
                                                var v104 uint64 = 2
                                                if (!(-2.25 > 1.5)) {
                                                    if false {
                                                        fmt.Println(v88, v93, v88)
                                                        if false {
                                                            var v105 int32 = v88
                                                            _ = v105
                                                        }
                                                        v106 := GenS8{F0: (fi89 & 42)}
                                                        _ = v106
                                                    }
                                                    v107 := v2[1]
                                                    fmt.Println(v67, v68)
                                                    _ = v107
                                                }
                                                if false {
                                                    v104 = 1
                                                    v93 -= v93
                                                    delete(v1, (fi89 ^ fi89))
                                                    v108 := []int{(fi87 | int(v67)), v68}
                                                    v93 = v93
                                                    _ = v108
                                                } else {
                                                    v93 = (v93 + 255)
                                                    switch (v74 + v74) {
                                                    case "\n\t":
                                                        if true {
                                                            var v109 float64 = -2.25
                                                            switch 13 {
                                                            case 287:
                                                                v74 = "\n\t"
                                                            }
                                                            fmt.Println(v3)
                                                            _ = v109
                                                        } else {
                                                            v110 := GenS5{F0: (-(v86))}
                                                            var v111 uint64 = (v104 * v104)
                                                            if false {
                                                                v67 += 42
                                                                v112 := GenS2{F0: (v86 + (v86 / (((v86 + v86)) | 1))), F1: v74, F2: 672}
                                                                _ = v112
                                                            } else {
                                                                var v113 bool = (uint8(v93) <= 42)
                                                                if v113 {
                                                                    v114 := GenS14{F0: (-(v86)), F1: 267}
                                                                    _ = v114
                                                                } else {
                                                                    v115 := []int{}
                                                                    var v116 float32 = (float32(3.14159) * (float32((0.0 / float64((((v3 * v3) - v3) * v3)))) / v3))
                                                                    var v117 uint = 4294967295
                                                                    v118 := GenS10{F0: 42, F1: 38, F2: ""}
                                                                    _ = v115
                                                                    _ = v116
                                                                    _ = v117
                                                                    _ = v118
                                                                }
                                                                v111 = (v111 >> ((238) & 63))
                                                                _ = v113
                                                            }
                                                            _ = v110
                                                            _ = v111
                                                        }
                                                    }
                                                    delete(v1, 128)
                                                }
                                                _ = v104
                                            } else {
                                                v119 := GenS7{F0: v74, F1: 0}
                                                var v120 uint16 = v93
                                                var v121 uint32 = 739
                                                v120 -= v93
                                                _ = v119
                                                _ = v120
                                                _ = v121
                                            }
                                        }
                                        v3 += 0.0
                                        switch ((((v74 + "\n\t") + v74) + "a") + ("foo bar" + (v74 + v74))) {
                                        case "a":
                                            delete(v1, fi89)
                                        case "\n\t":
                                            var v122 uint32 = 1
                                            _ = v122
                                        case "":
                                            v123 := GenS8{F0: (v68 / (((-(fi87))) | 1))}
                                            _ = v123
                                        }
                                    }
                                    v124 := []int{fi89, (fi89 | 42), -9223372036854775808}
                                    v88 -= -604
                                    _ = v124
                                    _ = v97
                                } else {
                                    var v125 int8 = -1
                                    _ = v125
                                }
                                _ = v93
                                _ = v94
                            }
                            _ = fi89
                        }
                        if (3.14159 <= -2.25) {
                            switch fi87 {
                            case -2147483648:
                                if (((^(fi87)) != -1) && false) {
                                    v86 = v86
                                } else {
                                    var v126 uint = 346
                                    for fi127 := 0; fi127 < 19; fi127++ {
                                        v128 := GenS12{F0: 42}
                                        var v129 uint8 = 128
                                        _ = fi127
                                        _ = v128
                                        _ = v129
                                    }
                                    for fi130 := 0; fi130 < 14; fi130++ {
                                        switch (v74 + "\n\t") {
                                        case "":
                                            v131 := GenS10{F0: v67, F1: -59, F2: (v74 + v74)}
                                            _ = v131
                                        case "foo bar":
                                            v68 = fi87
                                            _ = len(v1)
                                            switch (("hello" + v74) + "\n\t") {
                                            case "hello":
                                                if true {
                                                    fmt.Println(v68, v86, fi87)
                                                }
                                            case "\n\t":
                                                var v132 bool = (!((0.0 > -2.25) && (false && (!((1.5 <= (float64(v3) / float64((v3 + (((0.0 - v3) / v3) / 1.0))))) || true)))))
                                                _ = v132
                                            case "":
                                                v67 += v67
                                            }
                                        case "\n\t":
                                            v133 := GenS15{F0: 127, F1: 42}
                                            v68 = (fi130 << ((v126) & 63))
                                            _ = v133
                                        default:
                                            var v134 string = ("\n\t" + v74)
                                            _ = v134
                                        }
                                        if false {
                                            v135 := GenS8{F0: 2147483647}
                                            _ = v135
                                        } else {
                                            var v136 string = ("foo bar" + (v74 + v74))
                                            switch v68 {
                                            case 0:
                                                var v137 int = (v68 << ((v126) & 63))
                                                _ = v137
                                            case -203:
                                                if true {
                                                    var v138 int16 = (v67 & (v67 % ((v67) | 1)))
                                                    _ = v138
                                                } else {
                                                    var v139 float32 = v3
                                                    var v140 int64 = v86
                                                    v141 := GenS4{F0: -1}
                                                    v86 += (v140 + v86)
                                                    _ = v139
                                                    _ = v140
                                                    _ = v141
                                                }
                                            case 42:
                                                _ = len(v2)
                                            default:
                                                var v142 int32 = v88
                                                _ = v142
                                            }
                                            if (!(65535 != 520)) {
                                                fmt.Println(fi87, fi87, v126)
                                            } else {
                                                switch -2147483648 {
                                                case 232:
                                                    if (32767 >= (^(v67))) {
                                                        v143 := GenS16{F0: (^(v126))}
                                                        var v144 int8 = 100
                                                        _ = v143
                                                        _ = v144
                                                    } else {
                                                        v145 := []int{(fi130 % ((32768) | 1))}
                                                        var v146 uint64 = 2
                                                        var v147 float32 = -2.25
                                                        _ = v145
                                                        _ = v146
                                                        _ = v147
                                                    }
                                                case 9223372036854775807:
                                                    v148 := GenS4{F0: -49}
                                                    _ = v148
                                                default:
                                                    fmt.Println(v88, v126)
                                                }
                                                v3 = v3
                                                var v149 string = v136
                                                _ = v149
                                            }
                                            _ = v136
                                        }
                                        _ = fi130
                                    }
                                    _ = v126
                                }
                            case 42:
                                var v150 uint64 = 2
                                _ = v150
                            case -249:
                                if true {
                                    fmt.Println(fi87)
                                    fmt.Println(fi87)
                                    if true {
                                        switch (fi87 ^ (v68 + fi87)) {
                                        case -117:
                                            v151 := GenS17{F0: v67, F1: 1}
                                            _ = v151
                                        }
                                        v152 := v5[(v74 + v74)]
                                        _ = v152
                                    }
                                }
                            }
                            fmt.Println(v67)
                            var v153 float32 = 1.5
                            _ = v153
                        } else {
                            fmt.Println(v3, v68)
                            v3 -= v3
                        }
                        _ = fi87
                        _ = v88
                    }
                    _ = v85
                    _ = v86
                } else {
                    v68 -= 345
                    v154 := GenS1{F0: (true && true)}
                    var v155 uint8 = 255
                    _ = v154
                    _ = v155
                }
                _ = v74
            }
            for fi156 := 0; fi156 < 13; fi156++ {
                v3 -= (3.14159 / v3)
                for fi157 := 0; fi157 < 5; fi157++ {
                    v158 := []int{(fi157 & fi156), fi157, 127}
                    if len(v158) >= 2 {
                        v158 = v158[1:len(v158)]
                    }
                    _ = fi157
                    _ = v158
                }
                _ = fi156
            }
            var v159 int16 = (v67 & v67)
            _ = v159
        }
    } else {
        fmt.Println(v68, v68)
        v3 += v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v67
    _ = v68
    _ = v69
}
