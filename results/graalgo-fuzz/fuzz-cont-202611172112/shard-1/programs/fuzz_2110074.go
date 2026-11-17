package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int8 }
type GenS2 struct { F0 bool; F1 uint8; F2 bool }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 bool; F2 uint8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint; F1 int16 }
type GenS9 struct { F0 int8; F1 uint32; F2 int }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint8; F1 uint64 }
type GenS12 struct { F0 uint; F1 int; F2 int }
type GenS13 struct { F0 uint32; F1 uint8; F2 uint8 }
func genFn0(p0 string, p1 int32, p2 int32) string {
    var v4 bool = ((!false) && false)
    for fi5 := 0; fi5 < 12; fi5++ {
        if true {
            var v6 int32 = (p2 + p1)
            _ = v6
        }
        v7 := []int{(-(fi5)), fi5, fi5}
        _ = fi5
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p0 + "foo bar")
}
func genFn1(p0 uint64, p1 uint16, p2 uint) int16 {
    for fi8 := 0; fi8 < 16; fi8++ {
        p0 = p0
        if (-9223372036854775808 > 922) {
            var v9 uint32 = 42
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 int8 = int8(fi10)
                v12 := GenS1{F0: (false || ((^(fi8)) < -691)), F1: 42, F2: (v11 / ((int8(v9)) | 1))}
                var v13 int64 = 9223372036854775807
                v14 := GenS1{F0: (!false), F1: 1000, F2: 81}
                p2 += p2
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
                _ = v14
            }
            _ = v9
        }
        v15 := []int{(fi8 / (((fi8 ^ 128)) | 1))}
        _ = fi8
        _ = v15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        p0 = p0
        v17 := []int{fi16, int(p0), -661, (fi16 >> ((p2) & 63))}
        v18 := GenS2{F0: false, F1: 2, F2: (929 != (p2 | 42))}
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := []int{}
    var v20 bool = false
    for fi21 := 0; fi21 < 11; fi21++ {
        if (p1 <= p1) {
            v22 := GenS3{F0: 42, F1: 190}
            _ = len(v19)
            _ = v22
        }
        if (100 < -32) {
            if v20 {
                if v20 {
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                } else {
                    if true {
                        v20 = (v20 || (false && false))
                        _ = len(v19)
                        p0 += uint64(p2)
                    } else {
                        var v23 float64 = 0.1
                        if false {
                            if false {
                                for fi24 := 0; fi24 < 2; fi24++ {
                                    if len(v19) > 0 {
                                        v19[0] = fi21
                                    }
                                    var v25 int64 = -516
                                    var v26 uint32 = 2
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                }
                                var v27 uint8 = 218
                                v28 := GenS4{F0: fi21, F1: p0}
                                _ = v27
                                _ = v28
                            } else {
                                v20 = true
                                if false {
                                    switch ("hello" + (("a" + "hello") + ("foo bar" + "\n\t"))) {
                                    case "\n\t":
                                        if true {
                                            v29 := GenS2{F0: (false && (v20 || false)), F1: 2, F2: (!(false && v20))}
                                            v30 := GenS4{F0: (-(fi21)), F1: 0}
                                            var v31 uint16 = (p1 - p1)
                                            _ = v29
                                            _ = v30
                                            _ = v31
                                        }
                                    case "foo bar":
                                        switch "" {
                                        case "a":
                                            fmt.Println(p1, p0)
                                        case "hello":
                                            var v32 uint64 = p0
                                            _ = v32
                                        case "\n\t":
                                            _ = len(v19)
                                        default:
                                            for fi33 := 0; fi33 < 9; fi33++ {
                                                var v34 bool = (-1 < 42)
                                                var v35 float64 = v23
                                                p1 -= p1
                                                _ = fi33
                                                _ = v34
                                                _ = v35
                                            }
                                        }
                                    case "hello":
                                        p2 = p2
                                    default:
                                        switch (fi21 & 127) {
                                        case 2147483647:
                                            for fi36 := 0; fi36 < 13; fi36++ {
                                                v20 = false
                                                v19 = append(v19, 156)
                                                v37 := GenS3{F0: int32(p0), F1: 255}
                                                if v20 {
                                                    _ = len(v19)
                                                    var v38 int16 = -603
                                                    fmt.Println(fi21, fi36, v23)
                                                    _ = v38
                                                } else {
                                                    if false {
                                                        p1 -= p1
                                                        v39 := GenS2{F0: (fi21 >= (fi36 / (((fi21 % (((fi21 * 32768)) | 1))) | 1))), F1: 0, F2: v20}
                                                        _ = v39
                                                    }
                                                    var v40 string = "\n\t"
                                                    v19 = append(v19, 827)
                                                    _ = v40
                                                }
                                                if len(v19) >= 2 {
                                                    v19 = v19[1:len(v19)]
                                                }
                                                _ = fi36
                                                _ = v37
                                            }
                                        case 32768:
                                            v19 = append(v19, (-(fi21)))
                                        case -2147483648:
                                            _ = len(v19)
                                        default:
                                            for fi41 := 0; fi41 < 5; fi41++ {
                                                _ = len(v19)
                                                var v42 uint32 = 190
                                                _ = fi41
                                                _ = v42
                                            }
                                        }
                                    }
                                    v20 = (true && (1e10 >= -2.25))
                                } else {
                                    var v43 uint = (p2 & 42)
                                    _ = v43
                                }
                            }
                        }
                        _ = v23
                    }
                    if false {
                        v44 := GenS5{F0: fi21}
                        switch 4294967295 {
                        case 561:
                            for fi45 := 0; fi45 < 17; fi45++ {
                                if len(v19) >= 2 {
                                    v19 = v19[1:len(v19)]
                                }
                                var v46 string = "hello"
                                var v47 uint64 = p0
                                var v48 uint8 = 16
                                if true {
                                    var v49 uint8 = v48
                                    if (!(true || (42 == 480))) {
                                        p2 = (p2 ^ (p2 >> (((p2 % ((18446744073709551615) | 1))) & 63)))
                                    } else {
                                        p2 += (p2 << (((p2 & p2)) & 63))
                                    }
                                    _ = v49
                                }
                                v47 = 4294967295
                                var v50 uint16 = p1
                                _ = fi45
                                _ = v46
                                _ = v47
                                _ = v48
                                _ = v50
                            }
                        case 329:
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                        case 32768:
                            fmt.Println(p0, p2)
                        default:
                            var v51 string = "foo bar"
                            _ = v51
                        }
                        var v52 int64 = 210
                        _ = v44
                        _ = v52
                    } else {
                        for fi53 := 0; fi53 < 1; fi53++ {
                            switch ("foo bar" + "foo bar") {
                            case "a":
                                if v20 {
                                    var v54 bool = ((v20 || (0 <= -903)) && (2147483648 <= 1))
                                    v19 = append(v19, (fi53 % ((fi21) | 1)))
                                    _ = v54
                                } else {
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                    v55 := GenS1{F0: v20, F1: -731, F2: -1}
                                    _ = v55
                                }
                            case "hello":
                                if (false || ((int16(p0) != 653) && ((v20 && v20) && (127 == -128)))) {
                                    if (v20 || (true && false)) {
                                        v56 := GenS5{F0: (-(fi21))}
                                        _ = v56
                                    }
                                    v57 := GenS4{F0: 0, F1: 2}
                                    _ = v57
                                }
                            case "\n\t":
                                p1 = 2
                            default:
                                p1 = (^(p1))
                            }
                            _ = fi53
                        }
                        switch (((("a" + "hello") + "foo bar") + "") + ("\n\t" + "a")) {
                        case "\n\t":
                            v58 := GenS2{F0: v20, F1: 1, F2: (!(!((((!(fi21 == fi21)) && true) || (!(!false))) && (p0 >= p0))))}
                            _ = v58
                        case "hello":
                            if v20 {
                                switch ("hello" + "foo bar") {
                                case "hello":
                                    p0 += (p0 % ((p0) | 1))
                                case "foo bar":
                                    p2 = p2
                                case "a":
                                    if len(v19) > 0 {
                                        v19[0] = fi21
                                    }
                                }
                                v20 = v20
                                v19 = append(v19, (fi21 & -1))
                            }
                        default:
                            var v59 string = ("foo bar" + "\n\t")
                            _ = v59
                        }
                        var v60 uint8 = 42
                        _ = v60
                    }
                    if (!(v20 && (((2 < 792) || (!(uint32(p1) < 2147483647))) || ((p1 << ((p2) & 63)) <= p1)))) {
                        switch ("" + "hello") {
                        case "a":
                            if len(v19) > 0 {
                                v19[0] = 128
                            }
                        case "\n\t":
                            if true {
                                p2 = p2
                                v20 = ((fi21 ^ fi21) <= fi21)
                                v19 = append(v19, 9223372036854775807)
                            } else {
                                for fi61 := 0; fi61 < 15; fi61++ {
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                    _ = fi61
                                }
                                fmt.Println(p0, p2, v20)
                                for fi62 := 0; fi62 < 20; fi62++ {
                                    _ = len(v19)
                                    _ = fi62
                                }
                            }
                        case "":
                            v63 := GenS6{F0: "a", F1: v20, F2: 127}
                            _ = v63
                        }
                    }
                }
            }
            v19 = append(v19, 1)
        } else {
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            if (10 != 0) {
                switch (-(fi21)) {
                case 32767:
                    if v20 {
                        p0 += (p0 / ((p0) | 1))
                        fmt.Println(p0, p0, p0)
                        var v64 int64 = 1
                        _ = v64
                    } else {
                        var v65 string = "foo bar"
                        v66 := GenS7{F0: 301}
                        if len(v19) > 0 {
                            v19[0] = (fi21 & fi21)
                        }
                        _ = v65
                        _ = v66
                    }
                case 4294967295:
                    v19 = append(v19, -1)
                case -545:
                    switch "" {
                    case "a":
                        for fi67 := 0; fi67 < 2; fi67++ {
                            if (!(p2 != p2)) {
                                p1 += p1
                            } else {
                                fmt.Println(v20)
                                var v68 int = -2147483648
                                v20 = ((false && (-2.25 <= 0.0)) && v20)
                                _ = v68
                            }
                            v69 := GenS4{F0: (fi67 * fi67), F1: p0}
                            _ = fi67
                            _ = v69
                        }
                    }
                default:
                    v20 = (127 >= -40)
                }
                v70 := GenS7{F0: 1}
                _ = v70
            } else {
                v20 = v20
                p1 = p1
            }
            p0 += (p0 ^ (p0 | p0))
        }
        if ((true && v20) || v20) {
            v71 := GenS2{F0: v20, F1: 229, F2: (-2.25 == (0.1 + 0.1))}
            _ = v71
        } else {
            v72 := GenS7{F0: 1}
            _ = v72
        }
        _ = fi21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return int16(p0)
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[string]string)
    v3 := make(map[string]int)
    if true {
        v1 = (v1 * (v1 - v1))
        switch -1 {
        case -1:
            fmt.Println(v1)
        case 416:
            if true {
                for fi73 := 0; fi73 < 11; fi73++ {
                    v1 = ((v1 - v1) - ((1e10 + v1) / (3.14159 * float32((1e10 * -2.25)))))
                    _ = fi73
                }
            } else {
                fmt.Println(v1)
                var v74 uint32 = 2147483648
                if true {
                    var v75 int64 = 42
                    if (0 == genFn1(193, 255, 444)) {
                        if false {
                            v76 := GenS6{F0: "", F1: (65535 <= 2), F2: 2}
                            _ = v76
                        }
                        v1 = v1
                        for fi77 := 0; fi77 < 15; fi77++ {
                            switch "a" {
                            case "\n\t":
                                if (42 == -128) {
                                    if false {
                                        if false {
                                            delete(v3, (("" + (("a" + ("" + "a")) + "a")) + "hello"))
                                            for fi78 := 0; fi78 < 17; fi78++ {
                                                switch (fi78 << ((0) & 63)) {
                                                case 127:
                                                    v79 := GenS5{F0: 32767}
                                                    _ = v79
                                                case 9223372036854775807:
                                                    var v80 int64 = v75
                                                    _ = v80
                                                case -719:
                                                    v1 = float32(3.14159)
                                                }
                                                _ = fi78
                                            }
                                            v74 -= (v74 | v74)
                                        } else {
                                            delete(v2, ("\n\t" + ("\n\t" + "a")))
                                        }
                                    } else {
                                        v75 -= -1
                                        for fi81 := 0; fi81 < 12; fi81++ {
                                            if (!(true || (!true))) {
                                                var v82 float32 = v1
                                                v83 := GenS8{F0: 65535, F1: -95}
                                                switch 823 {
                                                case -3:
                                                    if true {
                                                        var v84 int8 = -1
                                                        fmt.Println(v75, fi81)
                                                        v85 := []int{}
                                                        var v86 uint8 = 1
                                                        _ = v84
                                                        _ = v85
                                                        _ = v86
                                                    } else {
                                                        v87 := GenS4{F0: fi77, F1: 18446744073709551615}
                                                        switch "hello" {
                                                        case "a":
                                                            var v88 int32 = int32(fi81)
                                                            _ = v88
                                                        case "\n\t":
                                                            switch ((("\n\t" + ("a" + "hello")) + ("hello" + ("\n\t" + "foo bar"))) + "\n\t") {
                                                            case "\n\t":
                                                                v1 -= v82
                                                            default:
                                                                if (true && true) {
                                                                    switch (("a" + ("\n\t" + ("foo bar" + ""))) + "foo bar") {
                                                                    case "foo bar":
                                                                        v82 = v1
                                                                    default:
                                                                        switch "" {
                                                                        case "":
                                                                            v1 -= v82
                                                                        case "hello":
                                                                            var v89 int64 = (v75 | v75)
                                                                            _ = v89
                                                                        }
                                                                    }
                                                                    v1 = 1.5
                                                                    var v90 bool = (1 == 0)
                                                                    v90 = v90
                                                                    var v91 int32 = 2147483647
                                                                    v92 := GenS8{F0: 255, F1: -14}
                                                                    _ = v90
                                                                    _ = v91
                                                                    _ = v92
                                                                }
                                                            }
                                                        }
                                                        v93 := []int{(fi77 - (fi81 - fi77))}
                                                        _ = v87
                                                        _ = v93
                                                    }
                                                    var v94 int = 42
                                                    _ = v94
                                                case -584:
                                                    switch ("\n\t" + "") {
                                                    case "a":
                                                        var v95 uint16 = 255
                                                        _ = v95
                                                    case "":
                                                        v96 := []int{fi81, -9223372036854775808, (fi81 - fi81)}
                                                        _ = v96
                                                    case "foo bar":
                                                        var v97 uint32 = v74
                                                        _ = v97
                                                    default:
                                                        var v98 uint32 = (v74 % ((v74) | 1))
                                                        _ = v98
                                                    }
                                                }
                                                switch "" {
                                                case "a":
                                                    v82 += 0.0
                                                default:
                                                    switch (("" + (("foo bar" + ("" + ("hello" + ""))) + (("a" + "") + ""))) + "hello") {
                                                    case "\n\t":
                                                        var v99 uint = 247
                                                        _ = v99
                                                    default:
                                                        v100 := []int{}
                                                        _ = v100
                                                    }
                                                }
                                                var v101 float32 = v1
                                                v75 -= -370
                                                _ = v101
                                                _ = v82
                                                _ = v83
                                            }
                                            v102 := GenS9{F0: -1, F1: v74, F2: (^(fi81))}
                                            _ = fi81
                                            _ = v102
                                        }
                                    }
                                    if ((float64(v1) * 0.1) == (1e10 - ((1.5 * 1e10) * (0.0 / 1.0)))) {
                                        var v103 int64 = v75
                                        _ = v103
                                    } else {
                                        v104 := GenS3{F0: -1, F1: 255}
                                        _ = v104
                                    }
                                    v75 += (v75 * (v75 << ((255) & 63)))
                                } else {
                                    if true {
                                        var v105 bool = (128 < 107)
                                        var v106 bool = ((v105 || v105) && (42 != 0))
                                        _ = v105
                                        _ = v106
                                    }
                                    delete(v2, ("hello" + "a"))
                                    for fi107 := 0; fi107 < 18; fi107++ {
                                        v108 := GenS10{F0: int32(v74)}
                                        _ = fi107
                                        _ = v108
                                    }
                                }
                            case "a":
                                v74 -= 2147483648
                            }
                            if false {
                                switch (fi77 << ((179) & 63)) {
                                case 128:
                                    v109 := GenS9{F0: 100, F1: v74, F2: (^(fi77))}
                                    _ = v109
                                }
                                var v110 uint = 501
                                _ = len(v3)
                                _ = v110
                            }
                            _ = fi77
                        }
                    } else {
                        var v111 float64 = float64((v1 + float32(3.14159)))
                        _ = v111
                    }
                    _ = v75
                }
                _ = v74
            }
        case 2147483647:
            for fi112 := 0; fi112 < 4; fi112++ {
                var v113 uint16 = 1
                v113 = v113
                _ = fi112
                _ = v113
            }
        default:
            v1 += 0.1
        }
        for fi114 := 0; fi114 < 1; fi114++ {
            var v115 bool = (370 == -37)
            if false {
                v116 := GenS2{F0: (v115 && (0 < 4294967295)), F1: 127, F2: true}
                v1 -= 0.0
                _ = v116
            }
            _ = fi114
            _ = v115
        }
    }
    switch 42 {
    case 32767:
        v1 = 1.5
    case 9223372036854775807:
        v117 := []int{9223372036854775807}
        _ = v117
    }
    v118 := []int{-561, -826, 0}
    fmt.Println(v1)
    if (!(869 != 42)) {
        for fi119 := 0; fi119 < 6; fi119++ {
            var v120 uint16 = 42
            v2[("" + "hello")] = "\n\t"
            if (70 > int64(v120)) {
                switch "hello" {
                case "":
                    for fi121 := 0; fi121 < 9; fi121++ {
                        if true {
                            var v122 uint8 = 128
                            _ = v122
                        } else {
                            v1 = (0.1 * (v1 / (3.14159 * v1)))
                        }
                        if false {
                            v1 = v1
                            v120 += 2
                            v123 := GenS3{F0: 42, F1: 128}
                            _ = v123
                        }
                        var v124 float32 = v1
                        if len(v118) >= 2 {
                            v118 = v118[1:len(v118)]
                        }
                        _ = fi121
                        _ = v124
                    }
                }
                if len(v118) >= 2 {
                    v118 = v118[1:len(v118)]
                }
                switch (fi119 + fi119) {
                case -2147483648:
                    var v125 float32 = (0.0 + v1)
                    _ = v125
                }
            } else {
                fmt.Println(v120, v120, v120)
                if (3.14159 == v1) {
                    if len(v118) >= 2 {
                        v118 = v118[1:len(v118)]
                    }
                    v120 = v120
                }
            }
            _ = fi119
            _ = v120
        }
        v1 -= ((3.14159 / v1) * v1)
        if (true && ((!(!false)) || (!true))) {
            if true {
                if ((!false) || (-340 != 191)) {
                    for fi126 := 0; fi126 < 4; fi126++ {
                        var v127 int = fi126
                        switch v127 {
                        case -2147483648:
                            var v128 uint32 = 137
                            _ = v128
                        case 795:
                            switch 1 {
                            case 9223372036854775807:
                                v1 -= v1
                            case 32768:
                                for fi129 := 0; fi129 < 8; fi129++ {
                                    if false {
                                        var v130 uint = 42
                                        var v131 uint = v130
                                        v131 -= v131
                                        v127 -= fi129
                                        _ = v130
                                        _ = v131
                                    } else {
                                        fmt.Println(v127, v127, v127)
                                        v127 = 32768
                                    }
                                    v132 := GenS2{F0: (false && (1e10 >= 1e10)), F1: 145, F2: true}
                                    _ = fi129
                                    _ = v132
                                }
                            default:
                                switch v127 {
                                case 316:
                                    switch ("foo bar" + "\n\t") {
                                    case "":
                                        v1 = (v1 * v1)
                                    case "a":
                                        v127 -= -858
                                    default:
                                        v133 := GenS2{F0: false, F1: 23, F2: true}
                                        _ = v133
                                    }
                                default:
                                    var v134 int = v127
                                    _ = v134
                                }
                            }
                        case 490:
                            var v135 int16 = -1
                            _ = v135
                        default:
                            fmt.Println(v1, fi126, v127)
                        }
                        switch v127 {
                        case -1:
                            switch (("hello" + "\n\t") + "foo bar") {
                            case "":
                                switch (^(v127)) {
                                case -1:
                                    fmt.Println(v1, v1, v127)
                                default:
                                    var v136 int64 = 306
                                    _ = v136
                                }
                            case "hello":
                                v137 := GenS11{F0: 41, F1: 535}
                                _ = v137
                            case "foo bar":
                                if (float32(((0.1 / 1e10) + 1e10)) == float32(((3.14159 - 3.14159) - 1e10))) {
                                    v138 := GenS9{F0: -128, F1: 30, F2: 42}
                                    v139 := GenS12{F0: 1, F1: (^(fi126)), F2: -2147483648}
                                    _ = v138
                                    _ = v139
                                } else {
                                    fmt.Println(v1, v1)
                                    for fi140 := 0; fi140 < 11; fi140++ {
                                        var v141 uint32 = 482
                                        var v142 float64 = ((1.5 / 1.0) - (1e10 + -2.25))
                                        _ = fi140
                                        _ = v141
                                        _ = v142
                                    }
                                }
                            }
                        case 86:
                            fmt.Println(v1)
                        case -2147483648:
                            var v143 string = "foo bar"
                            _ = v143
                        default:
                            v127 -= fi126
                        }
                        _ = fi126
                        _ = v127
                    }
                    if true {
                        v144 := GenS9{F0: 0, F1: 2, F2: 32767}
                        var v145 uint8 = 42
                        _ = v144
                        _ = v145
                    }
                    v1 += v1
                } else {
                    var v146 int64 = 0
                    _ = v146
                }
                var v147 uint = 1
                _ = len(v118)
                _ = v147
            } else {
                v1 -= 3.14159
            }
            v118 = append(v118, 144)
            var v148 int64 = 42
            _ = v148
        } else {
            for fi149 := 0; fi149 < 9; fi149++ {
                v150 := v3["hello"]
                v150 -= (^(fi149))
                _ = fi149
                _ = v150
            }
            if (42 != -946) {
                v151 := GenS13{F0: 611, F1: 1, F2: 0}
                _ = v151
            } else {
                var v152 int64 = 0
                v1 -= ((((float32(-2.25) - v1) * v1) / float32(((0.1 - 0.1) - 3.14159))) - v1)
                v153 := GenS8{F0: 0, F1: genFn1(0, 210, 42)}
                _ = v152
                _ = v153
            }
        }
    }
    v1 = (v1 + ((1.5 / v1) + ((v1 / (1e10 + ((0.1 * 1e10) / 1.5))) / float32(3.14159))))
    if (1 <= -2147483648) {
        switch "foo bar" {
        case "a":
            var v154 uint16 = 42
            _ = v154
        }
    } else {
        var v155 float32 = float32(1e10)
        fmt.Println(v155)
        var v156 int8 = -128
        _ = v155
        _ = v156
    }
    v1 -= ((float32(3.14159) + v1) / (0.1 * float32(1e10)))
    var v157 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v118:", len(v118), v118)
    fmt.Println("v157:", v157)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v118
    _ = v157
    _ = v2
    _ = v3
}
