package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 int64; F2 int16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS9 struct { F0 uint8; F1 uint8; F2 bool }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint; F1 string }
type GenS13 struct { F0 int32; F1 bool; F2 uint32 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS17 struct { F0 bool; F1 int16; F2 int64 }
type GenS18 struct { F0 string; F1 int }
type GenS19 struct { F0 int; F1 int; F2 int16 }
type GenS20 struct { F0 uint8; F1 uint16 }
type GenS21 struct { F0 int32; F1 int }
func genFn0(p0 int8, p1 int32) uint {
    v6 := GenS1{F0: -1}
    if (1000 > 0) {
        if true {
            v7 := GenS1{F0: 0}
            for fi8 := 0; fi8 < 8; fi8++ {
                if (3.14159 != 0.1) {
                    var v9 string = ("\n\t" + ("foo bar" + ""))
                    v10 := []int{(fi8 + fi8), fi8, 128}
                    _ = v10
                    _ = v9
                }
                _ = fi8
            }
            fmt.Println(p1, p1)
            _ = v7
        }
        v11 := []int{32768, 981, -228}
        for fi12 := 0; fi12 < 3; fi12++ {
            if ((!true) && true) {
                v13 := GenS2{F0: 299, F1: 379, F2: 32767}
                var v14 uint = 1
                _ = v13
                _ = v14
            } else {
                var v15 uint8 = 255
                _ = v15
            }
            if (255 == 42) {
                var v16 uint16 = 1
                p1 += (-(p1))
                _ = v16
            } else {
                switch "a" {
                case "":
                    p1 = -1
                }
                var v17 uint = 42
                _ = v17
            }
            _ = fi12
        }
        _ = v11
    }
    switch "hello" {
    case "":
        v18 := []int{526, int(p0)}
        _ = v18
    case "hello":
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 uint16 = 1
            if false {
                switch "a" {
                case "a":
                    p1 = p1
                case "hello":
                    var v21 float64 = float64(-2.25)
                    _ = v21
                case "\n\t":
                    v22 := []int{783, 128}
                    _ = v22
                default:
                    v23 := GenS1{F0: -1}
                    _ = v23
                }
            } else {
                v20 += v20
                var v24 int8 = 42
                _ = v24
            }
            switch "a" {
            case "":
                if (((true || false) && (("a" + "foo bar") < ("foo bar" + ("hello" + "")))) || ((!(true || ((0.1 < 1e10) && (p1 == 113)))) && true)) {
                    v25 := GenS1{F0: 42}
                    v26 := []int{}
                    if false {
                        for fi27 := 0; fi27 < 5; fi27++ {
                            var v28 uint = 18446744073709551615
                            var v29 float64 = (-2.25 - -2.25)
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    } else {
                        if true {
                            if (507 >= 4294967295) {
                                p1 -= p1
                                var v30 uint = 1
                                _ = v30
                            } else {
                                fmt.Println(p1, p0)
                                p0 = p0
                                var v31 float32 = 1e10
                                _ = v31
                            }
                            for fi32 := 0; fi32 < 2; fi32++ {
                                switch (fi19 << ((65535) & 63)) {
                                case 32767:
                                    var v33 uint16 = 65535
                                    _ = v33
                                }
                                if ((!((p1 * p1) > p1)) && (true && false)) {
                                    p0 -= -128
                                } else {
                                    if (!false) {
                                        v34 := GenS2{F0: 2147483648, F1: 749, F2: 1}
                                        var v35 int64 = -9223372036854775808
                                        switch "" {
                                        case "a":
                                            v26 = append(v26, fi19)
                                        }
                                        _ = v34
                                        _ = v35
                                    } else {
                                        var v36 float32 = 0.0
                                        _ = v36
                                    }
                                }
                                fmt.Println(fi32)
                                _ = fi32
                            }
                            fmt.Println(p1, p0)
                        } else {
                            var v37 uint = 601
                            if (!true) {
                                v20 = (v20 - (v20 - uint16(v37)))
                                v38 := GenS1{F0: 0}
                                for fi39 := 0; fi39 < 17; fi39++ {
                                    var v40 int64 = 42
                                    _ = fi39
                                    _ = v40
                                }
                                _ = v38
                            }
                            p0 += (p0 >> (((v37 - 42)) & 63))
                            _ = v37
                        }
                    }
                    _ = v25
                    _ = v26
                } else {
                    v41 := GenS2{F0: 1, F1: 0, F2: v20}
                    var v42 string = "hello"
                    _ = v41
                    _ = v42
                }
            case "foo bar":
                v20 = v20
            case "a":
                p0 += -85
            default:
                var v43 int64 = -982
                _ = v43
            }
            _ = fi19
            _ = v20
        }
    case "\n\t":
        p1 += 1
    default:
        switch "" {
        case "a":
            p1 -= 2147483647
        case "foo bar":
            if true {
                for fi44 := 0; fi44 < 10; fi44++ {
                    switch (fi44 / (((^(fi44))) | 1)) {
                    case 42:
                        v45 := GenS3{F0: 0}
                        _ = v45
                    case 4294967295:
                        p0 -= p0
                    case 32768:
                        fmt.Println(p0, p0, p1)
                    default:
                        p1 -= (p1 + 112)
                    }
                    _ = fi44
                }
                v46 := []int{}
                p0 += (p0 & (p0 + (-(p0))))
                _ = v46
            } else {
                v47 := GenS4{F0: 1}
                for fi48 := 0; fi48 < 8; fi48++ {
                    fmt.Println(fi48, p1, fi48)
                    _ = fi48
                }
                if true {
                    v49 := GenS2{F0: 884, F1: 32767, F2: 200}
                    _ = v49
                }
                _ = v47
            }
        case "\n\t":
            v50 := GenS5{F0: -2147483648, F1: 269, F2: 652}
            _ = v50
        }
    }
    switch "hello" {
    case "a":
        var v51 float32 = float32(float64((3.14159 + (0.1 + 0.0))))
        _ = v51
    case "hello":
        v52 := []int{42, 1, -2147483648}
        _ = v52
    }
    _ = p0
    _ = p1
    _ = v6
    return 0
}
func genFn1(p0 uint, p1 int64, p2 int8) int8 {
    fmt.Println(p0)
    var v53 uint = (p0 / ((p0) | 1))
    if (1 != -694) {
        switch "" {
        case "a":
            if false {
                if true {
                    for fi54 := 0; fi54 < 11; fi54++ {
                        v55 := GenS6{F0: (p2 << (((p0 | (p0 / (((p0 ^ 255)) | 1)))) & 63))}
                        _ = fi54
                        _ = v55
                    }
                    v56 := []int{42, 2147483647}
                    p0 = p0
                    _ = v56
                }
                for fi57 := 0; fi57 < 3; fi57++ {
                    if (0.1 > -2.25) {
                        v53 = 65535
                    } else {
                        var v58 bool = true
                        v59 := GenS4{F0: 1}
                        _ = v58
                        _ = v59
                    }
                    _ = fi57
                }
                if true {
                    p0 = p0
                    switch "\n\t" {
                    case "":
                        p1 += 42
                    default:
                        p0 += 255
                    }
                    var v60 int32 = 42
                    _ = v60
                }
            }
        case "hello":
            v53 = v53
        case "foo bar":
            v61 := GenS7{F0: -32768, F1: int8(p1)}
            _ = v61
        default:
            switch 4294967295 {
            case -1:
                if ((true && (-32768 > 32767)) && true) {
                    var v62 float32 = -2.25
                    _ = v62
                }
            case 0:
                for fi63 := 0; fi63 < 16; fi63++ {
                    p2 = 1
                    for fi64 := 0; fi64 < 3; fi64++ {
                        v65 := []int{(fi63 & (fi64 + fi63))}
                        _ = fi64
                        _ = v65
                    }
                    p1 = (p1 ^ 1)
                    _ = fi63
                }
            default:
                for fi66 := 0; fi66 < 6; fi66++ {
                    for fi67 := 0; fi67 < 19; fi67++ {
                        v68 := []int{fi67, fi66, fi66}
                        var v69 int32 = -2147483648
                        var v70 uint64 = 201
                        _ = fi67
                        _ = v68
                        _ = v69
                        _ = v70
                    }
                    _ = fi66
                }
            }
        }
        var v71 bool = true
        _ = v71
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v53
    return p2
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint32 = 945
    var v3 int64 = 42
    var v4 int32 = 914
    v5 := make(map[uint32]string)
    if (127 > -2147483648) {
        v4 += -780
    } else {
        fmt.Println(v1)
        if (!true) {
            fmt.Println(v1, v4, v1)
            v72 := []int{127, 477, 1, 127}
            switch 488 {
            case 128:
                for fi73 := 0; fi73 < 14; fi73++ {
                    v74 := GenS8{F0: -128, F1: 2, F2: v1}
                    v75 := GenS4{F0: 65535}
                    var v76 string = "hello"
                    _ = fi73
                    _ = v74
                    _ = v75
                    _ = v76
                }
            case 9223372036854775807:
                for fi77 := 0; fi77 < 3; fi77++ {
                    v78 := GenS9{F0: 178, F1: 127, F2: false}
                    switch "" {
                    case "\n\t":
                        if false {
                            fmt.Println(v3, v4)
                            _ = len(v5)
                        } else {
                            if false {
                                if true {
                                    v1 -= (v1 ^ v1)
                                    v79 := GenS3{F0: 127}
                                    var v80 int64 = 615
                                    _ = v79
                                    _ = v80
                                }
                            }
                        }
                    case "foo bar":
                        for fi81 := 0; fi81 < 12; fi81++ {
                            if ((^(v1)) != v1) {
                                fmt.Println(v3, v4)
                                switch (fi77 + fi77) {
                                case 442:
                                    var v82 float64 = 1e10
                                    _ = v82
                                case -9223372036854775808:
                                    v3 += v3
                                default:
                                    var v83 int16 = int16(v2)
                                    _ = v83
                                }
                            } else {
                                fmt.Println(v2, v3)
                            }
                            v84 := GenS10{F0: 65535}
                            _ = fi81
                            _ = v84
                        }
                    default:
                        switch fi77 {
                        case 42:
                            _ = len(v5)
                        case -534:
                            var v85 float64 = 0.1
                            _ = v85
                        default:
                            _ = len(v72)
                        }
                    }
                    if false {
                        if (v1 <= (v1 | v1)) {
                            v3 += v3
                            var v86 uint8 = 1
                            var v87 uint8 = (v86 ^ v86)
                            _ = v86
                            _ = v87
                        } else {
                            switch fi77 {
                            case 0:
                                var v88 int16 = 72
                                _ = v88
                            case 9223372036854775807:
                                _ = len(v5)
                            case 4294967295:
                                if (1e10 <= 3.14159) {
                                    if true {
                                        switch (("" + "\n\t") + ("a" + "")) {
                                        case "a":
                                            var v89 uint64 = 1
                                            _ = v89
                                        case "":
                                            v3 += v3
                                        case "\n\t":
                                            var v90 uint16 = 1
                                            _ = v90
                                        }
                                        switch fi77 {
                                        case 32767:
                                            var v91 int16 = -1
                                            _ = v91
                                        case -89:
                                            switch fi77 {
                                            case 1:
                                                v4 -= int32(v3)
                                            case 9223372036854775807:
                                                v4 = (^(v4))
                                            default:
                                                switch fi77 {
                                                case 1:
                                                    v92 := GenS5{F0: (fi77 | (fi77 & fi77)), F1: -152, F2: 1000}
                                                    _ = v92
                                                default:
                                                    var v93 uint8 = 1
                                                    _ = v93
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                            switch "hello" {
                            case "\n\t":
                                switch "hello" {
                                case "":
                                    if (false || true) {
                                        for fi94 := 0; fi94 < 2; fi94++ {
                                            v72 = append(v72, fi77)
                                            v95 := GenS11{F0: 0}
                                            switch "hello" {
                                            case "foo bar":
                                                v1 = (v1 * (v1 + v1))
                                            case "a":
                                                if ((^(v3)) != int64(v1)) {
                                                    switch "hello" {
                                                    case "a":
                                                        if (-128 > genFn1(42, 797, genFn1(2, v3, 42))) {
                                                            v3 += -9223372036854775808
                                                            var v96 uint32 = 253
                                                            switch ("" + (("\n\t" + "") + "a")) {
                                                            case "a":
                                                                fmt.Println(v1, v3, v3)
                                                                fmt.Println(v4, fi94, v2)
                                                            case "\n\t":
                                                                var v97 int8 = 100
                                                                _ = v97
                                                            case "hello":
                                                                var v98 uint16 = 458
                                                                _ = v98
                                                            }
                                                            _ = v96
                                                        }
                                                    case "\n\t":
                                                        var v99 string = "hello"
                                                        _ = v99
                                                    default:
                                                        switch ("\n\t" + ((("\n\t" + ("hello" + (("" + "foo bar") + ""))) + "a") + "a")) {
                                                        case "foo bar":
                                                            v4 = v4
                                                        default:
                                                            switch "foo bar" {
                                                            case "":
                                                                var v100 int16 = 0
                                                                _ = v100
                                                            case "hello":
                                                                var v101 float32 = (1e10 - ((0.1 - -2.25) - ((0.0 - 0.0) - (3.14159 / 1.5))))
                                                                _ = v101
                                                            case "a":
                                                                v4 -= v4
                                                            }
                                                        }
                                                    }
                                                    v4 = v4
                                                    v2 = (v2 % ((v2) | 1))
                                                } else {
                                                    v3 -= (v3 + (v3 << ((genFn0(0, (-(v4)))) & 63)))
                                                    if (1.5 > 1e10) {
                                                        v3 = 0
                                                        v4 = v4
                                                    } else {
                                                        switch ("hello" + "") {
                                                        case "a":
                                                            _ = len(v5)
                                                        case "hello":
                                                            fmt.Println(v3, fi77, v3)
                                                        case "foo bar":
                                                            v1 += v1
                                                        }
                                                        var v102 uint = genFn0(-30, (v4 >> ((4294967295) & 63)))
                                                        _ = v102
                                                    }
                                                    switch ("hello" + "") {
                                                    case "foo bar":
                                                        if false {
                                                            var v103 string = "hello"
                                                            var v104 uint64 = v1
                                                            _ = len(v72)
                                                            _ = v103
                                                            _ = v104
                                                        }
                                                    case "\n\t":
                                                        if ("\n\t" <= "hello") {
                                                            v2 = 2147483647
                                                            var v105 int32 = v4
                                                            _ = v105
                                                        }
                                                    default:
                                                        if false {
                                                            v5[(^(v2))] = ""
                                                            v106 := GenS7{F0: 455, F1: -1}
                                                            var v107 float64 = (0.1 - (1e10 - 0.0))
                                                            _ = v106
                                                            _ = v107
                                                        } else {
                                                            fmt.Println(v2, fi77, v2)
                                                        }
                                                    }
                                                }
                                            case "\n\t":
                                                if true {
                                                    var v108 int32 = 0
                                                    if len(v72) >= 2 {
                                                        v72 = v72[1:len(v72)]
                                                    }
                                                    v109 := GenS8{F0: 1, F1: 42, F2: 598}
                                                    if (v4 <= v4) {
                                                        v108 += v4
                                                        switch 127 {
                                                        case 128:
                                                            if (((v1 >> ((794) & 63)) >= 2) || (v108 >= 925)) {
                                                                v110 := GenS12{F0: 112, F1: (("" + "foo bar") + "a")}
                                                                v111 := GenS8{F0: 1, F1: 32767, F2: (v1 & (v1 - v1))}
                                                                v112 := GenS4{F0: 729}
                                                                _ = v110
                                                                _ = v111
                                                                _ = v112
                                                            } else {
                                                                v2 += v2
                                                            }
                                                            if (!(32767 < int(v3))) {
                                                                v113 := GenS5{F0: 127, F1: 1, F2: 1000}
                                                                var v114 uint64 = v1
                                                                v114 = v1
                                                                _ = v113
                                                                _ = v114
                                                            }
                                                            v115 := GenS13{F0: (^(v108)), F1: true, F2: v2}
                                                            _ = v115
                                                        case -9223372036854775808:
                                                            var v116 uint16 = 579
                                                            _ = v116
                                                        }
                                                    }
                                                    v4 = (v108 & v108)
                                                    _ = v108
                                                    _ = v109
                                                }
                                            default:
                                                var v117 int32 = 670
                                                fmt.Println(v2, fi77, fi94)
                                                var v118 uint8 = 128
                                                _ = v117
                                                _ = v118
                                            }
                                            _ = fi94
                                            _ = v95
                                        }
                                        switch fi77 {
                                        case 128:
                                            v119 := v5[v2]
                                            _ = v119
                                        case 2147483647:
                                            switch (fi77 % ((-399) | 1)) {
                                            case 896:
                                                v120 := GenS3{F0: 128}
                                                _ = v120
                                            default:
                                                switch "\n\t" {
                                                case "a":
                                                    var v121 float32 = 0.1
                                                    _ = v121
                                                default:
                                                    v5[v2] = "a"
                                                }
                                            }
                                        }
                                    }
                                case "\n\t":
                                    if ("\n\t" == (("foo bar" + ((("foo bar" + "a") + "a") + "hello")) + "\n\t")) {
                                        var v122 string = ""
                                        _ = v122
                                    }
                                default:
                                    for fi123 := 0; fi123 < 13; fi123++ {
                                        var v124 int32 = -2147483648
                                        v125 := GenS2{F0: (v2 >> ((18446744073709551615) & 63)), F1: int16(v4), F2: 57}
                                        _ = fi123
                                        _ = v124
                                        _ = v125
                                    }
                                }
                            case "":
                                if (v1 <= 0) {
                                    v1 = v1
                                }
                            case "hello":
                                for fi126 := 0; fi126 < 1; fi126++ {
                                    v127 := v5[v2]
                                    fmt.Println(v127)
                                    var v128 bool = true
                                    _ = fi126
                                    _ = v127
                                    _ = v128
                                }
                            default:
                                var v129 uint32 = v2
                                _ = v129
                            }
                            switch "hello" {
                            case "":
                                var v130 float32 = -2.25
                                _ = v130
                            case "hello":
                                v1 = v1
                            case "a":
                                v1 += v1
                            default:
                                if false {
                                    v3 = (v3 + (v3 | (v3 | (v3 / ((v3) | 1)))))
                                    var v131 uint8 = 255
                                    _ = v131
                                } else {
                                    if ((false && (v4 != v4)) || (!false)) {
                                        v132 := GenS2{F0: v2, F1: -32768, F2: 1}
                                        if (127 < 1) {
                                            var v133 uint32 = v2
                                            _ = v133
                                        }
                                        v134 := GenS13{F0: v4, F1: true, F2: (v2 - (v2 % ((v2) | 1)))}
                                        _ = v132
                                        _ = v134
                                    }
                                    v4 = v4
                                }
                            }
                        }
                        v135 := GenS4{F0: 302}
                        if (((v2 - (v2 - v2)) != v2) || false) {
                            if (-1 <= 1000) {
                                for fi136 := 0; fi136 < 7; fi136++ {
                                    var v137 int64 = v3
                                    v3 -= v3
                                    var v138 uint = uint(v2)
                                    v138 -= v138
                                    _ = fi136
                                    _ = v137
                                    _ = v138
                                }
                                v2 += v2
                            }
                        } else {
                            v2 -= 244
                        }
                        _ = v135
                    }
                    _ = fi77
                    _ = v78
                }
            }
            _ = v72
        }
    }
    var v139 uint8 = 255
    var v140 uint16 = 0
    v2 += v2
    fmt.Println(v140)
    v4 = (v4 + int32(v139))
    var v141 int8 = 1
    switch ("hello" + ("\n\t" + ("\n\t" + "foo bar"))) {
    case "hello":
        v1 = v1
    case "\n\t":
        switch "" {
        case "a":
            v142 := []int{int(v3), 932}
            _ = v142
        case "\n\t":
            v143 := GenS6{F0: (v141 | int8(v140))}
            _ = v143
        default:
            for fi144 := 0; fi144 < 2; fi144++ {
                v139 += (v139 * 2)
                for fi145 := 0; fi145 < 10; fi145++ {
                    if (!(!(-2.25 < -2.25))) {
                        if true {
                            var v146 string = "a"
                            _ = v146
                        }
                        switch "\n\t" {
                        case "\n\t":
                            v147 := []int{fi144, -1}
                            _ = v147
                        default:
                            var v148 uint = 65535
                            switch fi144 {
                            case -282:
                                v149 := GenS14{F0: v148}
                                v150 := GenS2{F0: (v2 * 42), F1: -817, F2: uint16(v2)}
                                if false {
                                    v4 = (v4 / ((int32(v3)) | 1))
                                    v5[2147483647] = "foo bar"
                                    v151 := GenS7{F0: -570, F1: -1}
                                    v148 = genFn0((v141 / (((v141 + (-(v141)))) | 1)), v4)
                                    if true {
                                        var v152 float32 = 1.5
                                        _ = v152
                                    }
                                    _ = v151
                                }
                                _ = v149
                                _ = v150
                            }
                            switch "\n\t" {
                            case "hello":
                                v153 := GenS13{F0: (v4 % ((-1) | 1)), F1: (true || ((true && ((v139 >> (((v148 % ((uint(v2)) | 1))) & 63)) == 255)) && (!false))), F2: 849}
                                _ = v153
                            case "a":
                                v139 += 42
                            case "":
                                v154 := GenS15{F0: v139}
                                _ = v154
                            }
                            v1 -= (v1 * v1)
                            v141 -= (v141 % ((genFn1(uint(v2), 1, (v141 >> (((v148 / ((440) | 1))) & 63)))) | 1))
                            v141 += 100
                            _ = v148
                        }
                        v155 := v5[0]
                        _ = v155
                    }
                    if (!(-1 > 159)) {
                        v156 := GenS16{F0: (^(v140)), F1: (v4 >> ((683) & 63)), F2: 127}
                        var v157 bool = (false && true)
                        v158 := []int{fi145, fi144, fi144, 1}
                        _ = v156
                        _ = v157
                        _ = v158
                    }
                    _ = fi145
                }
                _ = fi144
            }
        }
    default:
        fmt.Println(v1)
    }
    if (42 >= 0) {
        for fi159 := 0; fi159 < 9; fi159++ {
            v141 = (v141 + v141)
            _ = fi159
        }
        if ((true && ("hello" != (((("hello" + "hello") + "a") + ((("" + (("foo bar" + ("foo bar" + "a")) + "\n\t")) + "hello") + ((("foo bar" + "") + "hello") + ""))) + "hello"))) && (!true)) {
            v160 := []int{42, 554}
            fmt.Println(v139, v141, v4)
            v161 := GenS17{F0: (!false), F1: 32767, F2: v3}
            _ = v160
            _ = v161
        }
        if (32767 > int(v4)) {
            switch -9223372036854775808 {
            case 128:
                delete(v5, v2)
            }
            switch 32767 {
            case 1:
                v162 := GenS1{F0: (v3 / ((v3) | 1))}
                _ = v162
            case 123:
                v163 := GenS14{F0: 18446744073709551615}
                _ = v163
            case 4294967295:
                switch ("hello" + "foo bar") {
                case "a":
                    if false {
                        if false {
                            for fi164 := 0; fi164 < 4; fi164++ {
                                var v165 float32 = 0.0
                                v141 = v141
                                for fi166 := 0; fi166 < 10; fi166++ {
                                    v167 := []int{fi166, fi166}
                                    fmt.Println(v141, v141, v2)
                                    if false {
                                        v141 = v141
                                        if false {
                                            switch fi164 {
                                            case 42:
                                                v140 = 2
                                            case -1:
                                                if len(v167) > 0 {
                                                    v167[0] = 32768
                                                }
                                            }
                                            v140 += (v140 >> ((18446744073709551615) & 63))
                                        }
                                        switch (^(fi166)) {
                                        case 480:
                                            var v168 int = (fi166 - (fi164 & (fi166 >> ((763) & 63))))
                                            _ = v168
                                        default:
                                            fmt.Println(v3)
                                        }
                                    } else {
                                        delete(v5, (v2 * 2147483647))
                                        var v169 int8 = v141
                                        _ = v169
                                    }
                                    _ = fi166
                                    _ = v167
                                }
                                _ = fi164
                                _ = v165
                            }
                            for fi170 := 0; fi170 < 10; fi170++ {
                                v3 = v3
                                _ = fi170
                            }
                        } else {
                            v171 := GenS2{F0: 1, F1: -1, F2: v140}
                            switch "foo bar" {
                            case "foo bar":
                                switch 4294967295 {
                                case 42:
                                    v172 := GenS2{F0: v2, F1: 42, F2: 42}
                                    _ = v172
                                }
                            case "a":
                                v1 = 42
                            default:
                                v1 -= uint64(v139)
                            }
                            if false {
                                var v173 int64 = v3
                                _ = v173
                            } else {
                                v139 += v139
                                if false {
                                    var v174 uint64 = (v1 ^ (^(v1)))
                                    _ = v174
                                } else {
                                    switch 9223372036854775807 {
                                    case 32768:
                                        v4 = (v4 % (((v4 % (((^(v4))) | 1))) | 1))
                                    }
                                }
                                var v175 float32 = (1.5 / (-2.25 * float32(-2.25)))
                                _ = v175
                            }
                            _ = v171
                        }
                        if (true || true) {
                            var v176 string = ""
                            var v177 uint = 547
                            _ = v176
                            _ = v177
                        } else {
                            var v178 int8 = (v141 % ((v141) | 1))
                            if (true || true) {
                                if (-98 > int8(v1)) {
                                    v5[(v2 >> ((18446744073709551615) & 63))] = ("\n\t" + "hello")
                                    var v179 int64 = (v3 & (-(v3)))
                                    v180 := v5[(v2 + 4294967295)]
                                    _ = v179
                                    _ = v180
                                }
                                if (!false) {
                                    for fi181 := 0; fi181 < 11; fi181++ {
                                        v182 := []int{fi181, fi181, fi181}
                                        _ = fi181
                                        _ = v182
                                    }
                                    if false {
                                        switch (("hello" + "hello") + "") {
                                        case "":
                                            switch "foo bar" {
                                            case "a":
                                                fmt.Println(v3, v2, v178)
                                            case "foo bar":
                                                v183 := GenS18{F0: "hello", F1: 207}
                                                _ = v183
                                            case "hello":
                                                v184 := GenS13{F0: (v4 >> ((genFn0(100, int32(v141))) & 63)), F1: true, F2: (v2 >> ((311) & 63))}
                                                _ = v184
                                            }
                                        case "foo bar":
                                            v5[217] = "\n\t"
                                        }
                                        if false {
                                            switch "a" {
                                            case "hello":
                                                if false {
                                                    v185 := GenS9{F0: (v139 | (^(v139))), F1: v139, F2: (105 < genFn1(1, (v3 % ((v3) | 1)), genFn1(4294967295, int64(v178), (v141 ^ (v178 << ((genFn0((v141 / (((^(v178))) | 1)), 2147483647)) & 63))))))}
                                                    v4 += (v4 - int32(v140))
                                                    _ = v185
                                                } else {
                                                    switch -1 {
                                                    case -666:
                                                        v139 += 1
                                                    case 67:
                                                        if true {
                                                            switch "hello" {
                                                            case "":
                                                                switch -2147483648 {
                                                                case 959:
                                                                    v186 := []int{690, 716, 4294967295}
                                                                    _ = v186
                                                                case 952:
                                                                    switch -1 {
                                                                    case -515:
                                                                        var v187 int = -2147483648
                                                                        _ = v187
                                                                    case 13:
                                                                        delete(v5, v2)
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                var v188 uint64 = (v1 << ((65535) & 63))
                                                                _ = v188
                                                            default:
                                                                v189 := []int{-675}
                                                                _ = v189
                                                            }
                                                            delete(v5, (v2 * v2))
                                                        } else {
                                                            for fi190 := 0; fi190 < 9; fi190++ {
                                                                switch (("a" + "hello") + ("\n\t" + "a")) {
                                                                case "":
                                                                    if ((false || (32767 <= 483)) || (42 <= 18446744073709551615)) {
                                                                        fmt.Println(v140)
                                                                        var v191 string = ""
                                                                        delete(v5, v2)
                                                                        _ = v191
                                                                    } else {
                                                                        var v192 uint32 = (v2 & v2)
                                                                        var v193 int = fi190
                                                                        if ((v140 > v140) || (0 == (v139 / (((v139 / ((v139) | 1))) | 1)))) {
                                                                            switch "foo bar" {
                                                                            case "foo bar":
                                                                                if (false && true) {
                                                                                    for fi194 := 0; fi194 < 6; fi194++ {
                                                                                        fmt.Println(v192, v178, v2)
                                                                                        var v195 int16 = 482
                                                                                        v196 := []int{fi190, fi190, -655, (fi194 << ((19) & 63))}
                                                                                        _ = fi194
                                                                                        _ = v195
                                                                                        _ = v196
                                                                                    }
                                                                                    v4 = (v4 + (v4 << ((genFn0(v141, (v4 - (v4 + int32(v139))))) & 63)))
                                                                                } else {
                                                                                    for fi197 := 0; fi197 < 5; fi197++ {
                                                                                        fmt.Println(v139)
                                                                                        var v198 uint8 = uint8(v140)
                                                                                        fmt.Println(fi197)
                                                                                        _ = fi197
                                                                                        _ = v198
                                                                                    }
                                                                                }
                                                                            default:
                                                                                if true {
                                                                                    if ((false || false) || ((!((0.1 * 1.5) <= (-2.25 + 0.1))) || ((v4 ^ (v4 + (-(v4)))) <= (v4 >> ((42) & 63))))) {
                                                                                        var v199 uint = 743
                                                                                        switch 2147483647 {
                                                                                        case 4294967295:
                                                                                            switch ("a" + "a") {
                                                                                            case "hello":
                                                                                                v178 -= (-(v141))
                                                                                            case "a":
                                                                                                switch (^(v193)) {
                                                                                                case 42:
                                                                                                    v178 += v178
                                                                                                default:
                                                                                                    for fi200 := 0; fi200 < 6; fi200++ {
                                                                                                        _ = len(v5)
                                                                                                        if true {
                                                                                                            if false {
                                                                                                                if true {
                                                                                                                    switch (v193 - (v193 >> (((v199 ^ (^(v199)))) & 63))) {
                                                                                                                    case 0:
                                                                                                                        switch 530 {
                                                                                                                        case 500:
                                                                                                                            var v201 int64 = (v3 >> (((v199 << ((genFn0(127, 42)) & 63))) & 63))
                                                                                                                            _ = v201
                                                                                                                        case 9223372036854775807:
                                                                                                                            var v202 bool = (((0.0 + 0.1) <= (0.0 - -2.25)) || false)
                                                                                                                            _ = v202
                                                                                                                        case -16:
                                                                                                                            v203 := GenS17{F0: (!(false && false)), F1: 625, F2: (v3 - (v3 | -652))}
                                                                                                                            _ = v203
                                                                                                                        default:
                                                                                                                            v204 := GenS16{F0: v140, F1: (-(v4)), F2: uint8(v1)}
                                                                                                                            _ = v204
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v205 := v5[(v2 / ((v192) | 1))]
                                                                                                                    _ = v205
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        _ = fi200
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        case 128:
                                                                                            v3 -= (^(v3))
                                                                                        case 1:
                                                                                            for fi206 := 0; fi206 < 4; fi206++ {
                                                                                                if true {
                                                                                                    v207 := v5[2147483647]
                                                                                                    v178 -= -128
                                                                                                    v208 := GenS19{F0: (-(fi190)), F1: (fi190 ^ -631), F2: 944}
                                                                                                    _ = v207
                                                                                                    _ = v208
                                                                                                } else {
                                                                                                    v209 := v5[42]
                                                                                                    v178 = 22
                                                                                                    _ = v209
                                                                                                }
                                                                                                if ((v141 % (((v178 | genFn1((v199 * 141), v3, (v141 | (v178 - genFn1((v199 - (v199 / (((v199 / ((42) | 1))) | 1))), v3, v178)))))) | 1)) >= v141) {
                                                                                                    v193 = (^(v193))
                                                                                                    switch (v193 >> (((^(v199))) & 63)) {
                                                                                                    case 739:
                                                                                                        switch (fi206 >> ((uint(v1)) & 63)) {
                                                                                                        case 42:
                                                                                                            v3 = (v3 | v3)
                                                                                                        case 32768:
                                                                                                            var v210 uint32 = (v192 << (((v199 & v199)) & 63))
                                                                                                            _ = v210
                                                                                                        default:
                                                                                                            v199 += uint(v192)
                                                                                                        }
                                                                                                    default:
                                                                                                        var v211 uint8 = (v139 / (((v139 & (v139 | (^(v139))))) | 1))
                                                                                                        _ = v211
                                                                                                    }
                                                                                                    if true {
                                                                                                        if true {
                                                                                                            if false {
                                                                                                                if false {
                                                                                                                    v1 = uint64(v2)
                                                                                                                    var v212 uint16 = v140
                                                                                                                    _ = v212
                                                                                                                } else {
                                                                                                                    fmt.Println(fi206, v199)
                                                                                                                }
                                                                                                                var v213 uint32 = v2
                                                                                                                _ = v213
                                                                                                            }
                                                                                                        } else {
                                                                                                            v140 = 42
                                                                                                            switch "foo bar" {
                                                                                                            case "foo bar":
                                                                                                                if (true || ((18446744073709551615 <= v199) || false)) {
                                                                                                                    var v214 bool = true
                                                                                                                    _ = v214
                                                                                                                } else {
                                                                                                                    if ((v141 + (v141 ^ (v178 + v141))) <= (v141 + genFn1((^(v199)), 9223372036854775807, -1))) {
                                                                                                                        fmt.Println(v140, v4, v199)
                                                                                                                    }
                                                                                                                    v215 := GenS9{F0: 128, F1: (v139 >> ((v199) & 63)), F2: true}
                                                                                                                    _ = v215
                                                                                                                }
                                                                                                            case "":
                                                                                                                fmt.Println(v3, v4, v192)
                                                                                                            }
                                                                                                            v216 := []int{fi206, -550, (fi190 / (((v193 ^ fi190)) | 1)), v193}
                                                                                                            fmt.Println(v140, v141, v199)
                                                                                                            _ = v216
                                                                                                        }
                                                                                                        var v217 int = (fi206 / ((fi206) | 1))
                                                                                                        _ = v217
                                                                                                    }
                                                                                                }
                                                                                                _ = fi206
                                                                                            }
                                                                                        default:
                                                                                            for fi218 := 0; fi218 < 5; fi218++ {
                                                                                                switch 4294967295 {
                                                                                                case -1:
                                                                                                    if true {
                                                                                                        var v219 float64 = 0.1
                                                                                                        var v220 int8 = (v178 - (-(v141)))
                                                                                                        _ = v219
                                                                                                        _ = v220
                                                                                                    }
                                                                                                }
                                                                                                switch v193 {
                                                                                                case 9223372036854775807:
                                                                                                    var v221 float64 = 1e10
                                                                                                    _ = v221
                                                                                                default:
                                                                                                    switch v193 {
                                                                                                    case 68:
                                                                                                        v141 = 89
                                                                                                    case -1:
                                                                                                        if false {
                                                                                                            v5[(v192 & (v192 / ((v192) | 1)))] = "\n\t"
                                                                                                        }
                                                                                                    case 42:
                                                                                                        v4 = (v4 + -108)
                                                                                                    }
                                                                                                }
                                                                                                var v222 float32 = ((0.0 * 3.14159) - 1e10)
                                                                                                _ = fi218
                                                                                                _ = v222
                                                                                            }
                                                                                        }
                                                                                        _ = v199
                                                                                    }
                                                                                    v223 := GenS20{F0: v139, F1: 65535}
                                                                                    v139 = (v139 % ((v139) | 1))
                                                                                    _ = v223
                                                                                }
                                                                            }
                                                                            var v224 int16 = -32768
                                                                            v2 -= (v192 << ((4294967295) & 63))
                                                                            _ = v224
                                                                        }
                                                                        _ = v192
                                                                        _ = v193
                                                                    }
                                                                case "foo bar":
                                                                    v178 -= v178
                                                                case "\n\t":
                                                                    var v225 string = ""
                                                                    _ = v225
                                                                }
                                                                _ = fi190
                                                            }
                                                            for fi226 := 0; fi226 < 19; fi226++ {
                                                                var v227 float64 = (1.5 + (1e10 / 1.0))
                                                                var v228 uint64 = v1
                                                                _ = fi226
                                                                _ = v227
                                                                _ = v228
                                                            }
                                                        }
                                                    }
                                                    switch int(v1) {
                                                    case 631:
                                                        v140 = 42
                                                    case 128:
                                                        v141 -= genFn1(255, (v3 & (^(v3))), v141)
                                                    }
                                                }
                                            case "":
                                                v229 := GenS4{F0: (v140 + 769)}
                                                _ = v229
                                            case "\n\t":
                                                v230 := GenS9{F0: v139, F1: 2, F2: (!false)}
                                                _ = v230
                                            }
                                        }
                                    } else {
                                        if ((v139 << ((2) & 63)) >= 160) {
                                            fmt.Println(v2, v2)
                                            switch 1 {
                                            case -927:
                                                v178 -= v178
                                            case 1:
                                                v231 := []int{int(v140), 2147483647, -800}
                                                _ = v231
                                            }
                                        }
                                        v232 := []int{-9223372036854775808, -2147483648, -101, -9223372036854775808}
                                        v139 = 127
                                        _ = v232
                                    }
                                }
                                if false {
                                    if false {
                                        v233 := GenS10{F0: (v2 % (((^(v2))) | 1))}
                                        _ = v233
                                    } else {
                                        var v234 int = -171
                                        v235 := []int{(v234 + (v234 & v234)), 0}
                                        var v236 float64 = (0.1 + 0.0)
                                        _ = v234
                                        _ = v235
                                        _ = v236
                                    }
                                    var v237 uint32 = (v2 | v2)
                                    _ = v237
                                } else {
                                    if (42 < v1) {
                                        v238 := GenS8{F0: v141, F1: v140, F2: (v1 / ((v1) | 1))}
                                        _ = v238
                                    }
                                }
                            }
                            _ = v178
                        }
                    }
                default:
                    v139 -= uint8(v2)
                }
            default:
                v1 = 0
            }
        } else {
            if true {
                delete(v5, (v2 - 295))
                if (1.5 != 3.14159) {
                    if (!(v3 >= v3)) {
                        fmt.Println(v3, v141, v2)
                        if (1e10 < 0.1) {
                            for fi239 := 0; fi239 < 3; fi239++ {
                                v5[65535] = "a"
                                _ = fi239
                            }
                            fmt.Println(v1, v3, v140)
                        }
                    } else {
                        v240 := GenS21{F0: int32(v1), F1: 2147483647}
                        for fi241 := 0; fi241 < 5; fi241++ {
                            v242 := v5[uint32(v1)]
                            var v243 int8 = (v141 >> ((65535) & 63))
                            v139 += 255
                            _ = fi241
                            _ = v242
                            _ = v243
                        }
                        switch -351 {
                        case 32767:
                            if true {
                                for fi244 := 0; fi244 < 19; fi244++ {
                                    for fi245 := 0; fi245 < 13; fi245++ {
                                        v4 -= v4
                                        fmt.Println(v141, v141, fi244)
                                        _ = fi245
                                    }
                                    var v246 int = -9223372036854775808
                                    v246 += fi244
                                    _ = fi244
                                    _ = v246
                                }
                                v140 -= (v140 | 2)
                            }
                        default:
                            v141 = v141
                        }
                        _ = v240
                    }
                    var v247 int32 = v4
                    _ = v247
                }
            }
        }
    } else {
        v248 := GenS9{F0: 250, F1: uint8(v1), F2: false}
        v141 = (v141 + v141)
        _ = v248
    }
    fmt.Println("v1:", v1)
    fmt.Println("v139:", v139)
    fmt.Println("v140:", v140)
    fmt.Println("v141:", v141)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v139
    _ = v140
    _ = v141
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
