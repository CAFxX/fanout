package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint }
type GenS2 struct { F0 uint32; F1 int; F2 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 int8; F2 int32 }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 bool; F1 int8; F2 uint }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int8; F1 bool }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int64; F1 uint }
type GenS12 struct { F0 uint64; F1 string; F2 bool }
type GenS13 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS14 struct { F0 uint32; F1 uint32 }
type GenS15 struct { F0 string; F1 int; F2 uint8 }
type GenS16 struct { F0 int; F1 bool }

func main() {
    var v1 bool = (!true)
    var v2 uint32 = 736
    var v3 uint = 18446744073709551615
    var v4 uint = v3
    switch 32768 {
    case 1:
        v5 := GenS1{F0: uint16(v4), F1: 65535, F2: 655}
        _ = v5
    case 127:
        v6 := []int{434, int(v3), 127, 127}
        _ = v6
    case -616:
        var v7 uint8 = 127
        _ = v7
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        fmt.Println(v4, v1, v4)
        var v9 bool = (v2 > 2147483648)
        switch (("\n\t" + ("" + "")) + "foo bar") {
        case "foo bar":
            v10 := GenS2{F0: 1, F1: -2147483648, F2: -1}
            _ = v10
        case "a":
            v11 := GenS3{F0: 128}
            _ = v11
        default:
            switch (fi8 << (((^(v3))) & 63)) {
            case 9223372036854775807:
                fmt.Println(v3, v3, v3)
            case 2147483647:
                v1 = v9
            }
        }
        _ = fi8
        _ = v9
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 int64 = 307
        v2 += (v2 >> (((v3 / ((v4) | 1))) & 63))
        _ = fi12
        _ = v13
    }
    var v14 float64 = (-2.25 - (-2.25 * -2.25))
    var v15 int = int(v2)
    v15 = int(v2)
    if (("\n\t" + "\n\t") > "a") {
        if (v1 && (223 > 128)) {
            if (!((v14 - -2.25) > 1e10)) {
                v16 := GenS1{F0: 1, F1: 32767, F2: (^(v3))}
                _ = v16
            } else {
                fmt.Println(v2, v14, v4)
                var v17 uint16 = 65535
                _ = v17
            }
            var v18 int64 = 0
            v19 := GenS1{F0: 32767, F1: 1, F2: (^(v4))}
            _ = v18
            _ = v19
        } else {
            var v20 string = (("foo bar" + "a") + "\n\t")
            for fi21 := 0; fi21 < 6; fi21++ {
                switch (^(v15)) {
                case 127:
                    v22 := []int{v15}
                    _ = v22
                default:
                    v23 := GenS2{F0: (v2 % (((^(v2))) | 1)), F1: (-(v15)), F2: -32768}
                    _ = v23
                }
                var v24 int8 = 0
                _ = fi21
                _ = v24
            }
            _ = v20
        }
        switch "\n\t" {
        case "foo bar":
            v25 := GenS1{F0: uint16(v2), F1: 45, F2: v3}
            _ = v25
        case "hello":
            switch "" {
            case "\n\t":
                if (42 >= 2147483647) {
                    v4 += (^(v3))
                    if (false || v1) {
                        switch (^(v15)) {
                        case 0:
                            for fi26 := 0; fi26 < 18; fi26++ {
                                for fi27 := 0; fi27 < 4; fi27++ {
                                    v28 := []int{v15, (fi26 / ((-783) | 1)), 32767, int(v2)}
                                    switch "a" {
                                    case "foo bar":
                                        fmt.Println(v15, v14)
                                        if v1 {
                                            if v1 {
                                                var v29 int32 = 2147483647
                                                if len(v28) >= 2 {
                                                    v28 = v28[1:len(v28)]
                                                }
                                                if len(v28) >= 2 {
                                                    v28 = v28[1:len(v28)]
                                                }
                                                _ = v29
                                            } else {
                                                if ("foo bar" >= ("foo bar" + ("\n\t" + "a"))) {
                                                    var v30 string = ("a" + "hello")
                                                    _ = v30
                                                } else {
                                                    if ((!((v15 % ((-797) | 1)) == fi27)) || (0 < 100)) {
                                                        if len(v28) > 0 {
                                                            v28[0] = v15
                                                        }
                                                        v31 := GenS3{F0: 4294967295}
                                                        if false {
                                                            if v1 {
                                                                var v32 uint16 = uint16(v3)
                                                                if (v1 || (!v1)) {
                                                                    var v33 int = 195
                                                                    _ = v33
                                                                } else {
                                                                    v32 += v32
                                                                    fmt.Println(fi26, fi26)
                                                                }
                                                                v32 += (v32 ^ v32)
                                                                _ = v32
                                                            } else {
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    var v34 uint = (v3 - (v4 * 18446744073709551615))
                                                                    _ = v34
                                                                case "a":
                                                                    if len(v28) >= 2 {
                                                                        v28 = v28[1:len(v28)]
                                                                    }
                                                                default:
                                                                    var v35 float32 = 1e10
                                                                    _ = v35
                                                                }
                                                                v36 := GenS3{F0: fi26}
                                                                var v37 float32 = (1.5 + float32(v14))
                                                                _ = v36
                                                                _ = v37
                                                            }
                                                            v14 += (-2.25 / float64(1.5))
                                                            v1 = false
                                                        }
                                                        _ = v31
                                                    } else {
                                                        v3 += (v4 / (((^(v3))) | 1))
                                                        var v38 int8 = 0
                                                        v2 += v2
                                                        _ = v38
                                                    }
                                                    if (!v1) {
                                                        if len(v28) >= 2 {
                                                            v28 = v28[1:len(v28)]
                                                        }
                                                    }
                                                }
                                                fmt.Println(v2, v2, v1)
                                                var v39 uint = 1
                                                var v40 float64 = (v14 * (v14 / v14))
                                                _ = v39
                                                _ = v40
                                            }
                                        }
                                    case "a":
                                        fmt.Println(v15, v2, v14)
                                    case "\n\t":
                                        switch "\n\t" {
                                        case "a":
                                            var v41 int8 = 102
                                            _ = v41
                                        case "foo bar":
                                            fmt.Println(v1, v15)
                                        case "hello":
                                            switch fi27 {
                                            case 42:
                                                var v42 float32 = (0.0 - 1e10)
                                                _ = v42
                                            default:
                                                var v43 int32 = 2147483647
                                                _ = v43
                                            }
                                        default:
                                            if false {
                                                var v44 uint16 = 42
                                                _ = v44
                                            } else {
                                                v4 = 66
                                                var v45 int16 = -1
                                                fmt.Println(v15, v2, v14)
                                                if len(v28) >= 2 {
                                                    v28 = v28[1:len(v28)]
                                                }
                                                _ = len(v28)
                                                v4 -= v3
                                                _ = v45
                                            }
                                        }
                                    default:
                                        fmt.Println(v14, v15, v1)
                                    }
                                    v3 -= (^(v4))
                                    _ = fi27
                                    _ = v28
                                }
                                v46 := []int{}
                                _ = fi26
                                _ = v46
                            }
                        }
                    }
                    if v1 {
                        switch ("foo bar" + "") {
                        case "":
                            v47 := GenS4{F0: "foo bar", F1: 42, F2: -682}
                            _ = v47
                        case "foo bar":
                            var v48 float32 = (1.5 * -2.25)
                            _ = v48
                        default:
                            var v49 int = (v15 - v15)
                            _ = v49
                        }
                        if (-80 < 0) {
                            for fi50 := 0; fi50 < 8; fi50++ {
                                fmt.Println(v2)
                                _ = fi50
                            }
                        }
                        for fi51 := 0; fi51 < 14; fi51++ {
                            var v52 int16 = 1000
                            var v53 int8 = -128
                            _ = fi51
                            _ = v52
                            _ = v53
                        }
                    }
                }
            case "hello":
                for fi54 := 0; fi54 < 4; fi54++ {
                    v2 = (v2 ^ (v2 + v2))
                    v1 = v1
                    var v55 uint = 512
                    _ = fi54
                    _ = v55
                }
            default:
                v3 -= (v4 | (v3 & v4))
            }
        default:
            if (1 < -1) {
                switch "hello" {
                case "a":
                    for fi56 := 0; fi56 < 14; fi56++ {
                        v14 += (v14 / 1.5)
                        fmt.Println(v4, fi56, v15)
                        _ = fi56
                    }
                }
                v57 := []int{v15, v15}
                _ = v57
            } else {
                v58 := []int{v15, v15, v15, v15}
                _ = v58
            }
        }
        switch "a" {
        case "\n\t":
            if v1 {
                for fi59 := 0; fi59 < 10; fi59++ {
                    switch "\n\t" {
                    case "\n\t":
                        var v60 uint8 = 167
                        _ = v60
                    case "foo bar":
                        if (false || (0.0 >= float32(1e10))) {
                            v61 := []int{(v15 & fi59), (fi59 >> (((^(v4))) & 63)), int(v2)}
                            _ = v61
                        }
                    default:
                        v62 := GenS5{F0: v1, F1: (v4 << (((v3 >> ((18446744073709551615) & 63))) & 63))}
                        _ = v62
                    }
                    v63 := GenS3{F0: (-(fi59))}
                    v64 := GenS5{F0: ((fi59 + v15) <= fi59), F1: v4}
                    _ = fi59
                    _ = v63
                    _ = v64
                }
                v1 = v1
            } else {
                switch v15 {
                case 50:
                    v65 := []int{v15, 4294967295, 128, 368}
                    _ = v65
                case -9223372036854775808:
                    for fi66 := 0; fi66 < 12; fi66++ {
                        for fi67 := 0; fi67 < 19; fi67++ {
                            v68 := GenS4{F0: "", F1: 100, F2: -192}
                            _ = fi67
                            _ = v68
                        }
                        switch v15 {
                        case -335:
                            for fi69 := 0; fi69 < 19; fi69++ {
                                v70 := []int{(fi66 & (fi66 % ((fi69) | 1))), (v15 / ((fi66) | 1)), 127}
                                _ = fi69
                                _ = v70
                            }
                        default:
                            v14 = (v14 + v14)
                        }
                        _ = fi66
                    }
                default:
                    for fi71 := 0; fi71 < 8; fi71++ {
                        if (true || v1) {
                            var v72 uint32 = (v2 << ((3) & 63))
                            if (!(v1 && false)) {
                                if v1 {
                                    for fi73 := 0; fi73 < 11; fi73++ {
                                        var v74 float64 = v14
                                        _ = fi73
                                        _ = v74
                                    }
                                    v2 = (v2 | (v2 / ((v2) | 1)))
                                    if v1 {
                                        for fi75 := 0; fi75 < 3; fi75++ {
                                            v76 := []int{(v15 % (((fi75 * 32768)) | 1)), (v15 / (((fi71 ^ (v15 & (v15 << (((v3 - v4)) & 63))))) | 1))}
                                            v77 := GenS4{F0: ("a" + ("" + "")), F1: 1, F2: 2147483647}
                                            _ = fi75
                                            _ = v76
                                            _ = v77
                                        }
                                        v2 += (^(v72))
                                        v78 := []int{128, (v15 % (((fi71 * fi71)) | 1))}
                                        _ = v78
                                    } else {
                                        for fi79 := 0; fi79 < 19; fi79++ {
                                            v1 = (100 >= 0)
                                            v15 += (fi79 >> (((v4 / ((999) | 1))) & 63))
                                            v80 := GenS6{F0: v1, F1: 1, F2: (v3 & 42)}
                                            _ = fi79
                                            _ = v80
                                        }
                                        v81 := []int{}
                                        _ = v81
                                    }
                                }
                            } else {
                                v82 := GenS7{F0: 797, F1: ""}
                                fmt.Println(v14)
                                v1 = ((fi71 | fi71) <= (v15 | 42))
                                _ = v82
                            }
                            var v83 int64 = 0
                            _ = v72
                            _ = v83
                        }
                        _ = fi71
                    }
                }
                var v84 int64 = 9223372036854775807
                for fi85 := 0; fi85 < 12; fi85++ {
                    v86 := GenS7{F0: 2147483647, F1: ""}
                    fmt.Println(v84, fi85, v2)
                    _ = fi85
                    _ = v86
                }
                _ = v84
            }
        case "":
            v2 = uint32(v3)
        }
    } else {
        v2 = (v2 / ((v2) | 1))
        v87 := GenS8{F0: (v3 ^ 255)}
        v15 += (v15 - v15)
        _ = v87
    }
    var v88 uint64 = 42
    if (true && true) {
        var v89 int8 = -1
        _ = v89
    }
    switch ("" + "hello") {
    case "a":
        var v90 int = (v15 + 4294967295)
        _ = v90
    case "hello":
        if false {
            if true {
                switch ("foo bar" + "hello") {
                case "hello":
                    var v91 int = v15
                    _ = v91
                }
                if v1 {
                    var v92 uint8 = 1
                    var v93 uint16 = 0
                    v94 := []int{v15, v15, v15}
                    _ = v92
                    _ = v93
                    _ = v94
                } else {
                    var v95 int8 = -128
                    var v96 int16 = 1000
                    _ = v95
                    _ = v96
                }
                fmt.Println(v2, v4)
            }
            v4 += v4
            var v97 int16 = 809
            _ = v97
        }
    }
    if v1 {
        if false {
            if (!v1) {
                var v98 string = "foo bar"
                if (false && v1) {
                    for fi99 := 0; fi99 < 8; fi99++ {
                        switch v98 {
                        case "foo bar":
                            fmt.Println(v88, v3, v3)
                        }
                        _ = fi99
                    }
                    v4 += (^(v4))
                    for fi100 := 0; fi100 < 19; fi100++ {
                        switch "" {
                        case "hello":
                            v101 := GenS9{F0: -31, F1: (0.1 < 0.0)}
                            _ = v101
                        case "foo bar":
                            v4 = v3
                        case "":
                            var v102 int16 = 42
                            _ = v102
                        default:
                            var v103 bool = true
                            _ = v103
                        }
                        var v104 uint8 = 2
                        _ = fi100
                        _ = v104
                    }
                }
                v88 = (v88 & 42)
                _ = v98
            }
            v105 := GenS2{F0: 1, F1: -2147483648, F2: 0}
            v106 := []int{}
            _ = v105
            _ = v106
        } else {
            v107 := []int{}
            _ = v107
        }
        v108 := []int{32768, v15, v15}
        _ = v108
    }
    v14 -= v14
    v109 := []int{(v15 | 4294967295), (-(v15))}
    var v110 uint32 = (v2 / ((v2) | 1))
    for fi111 := 0; fi111 < 10; fi111++ {
        v110 -= v110
        var v112 int64 = 0
        _ = fi111
        _ = v112
    }
    for fi113 := 0; fi113 < 11; fi113++ {
        var v114 float32 = 1.5
        if (1 >= 1000) {
            var v115 uint = v4
            _ = v115
        } else {
            switch (("hello" + "hello") + (("foo bar" + "\n\t") + "a")) {
            case "\n\t":
                v88 = 176
            case "a":
                var v116 int = -149
                _ = v116
            }
        }
        switch (v15 >> ((v3) & 63)) {
        case -917:
            if v1 {
                var v117 uint64 = 213
                _ = v117
            }
        default:
            var v118 uint16 = 666
            _ = v118
        }
        _ = fi113
        _ = v114
    }
    switch 1 {
    case 4294967295:
        var v119 uint32 = v110
        _ = v119
    }
    if ((0 <= 756) || true) {
        if (!((v14 / (v14 - ((float64(((((float32(v14) * (0.0 - -2.25)) - 1.5) - 1.5) / -2.25)) + 1.5) * 3.14159))) <= 0.1)) {
            v2 = (v2 % ((v110) | 1))
            v120 := GenS10{F0: 0}
            v109 = append(v109, v15)
            _ = v120
        } else {
            if true {
                v121 := GenS8{F0: (v4 / ((255) | 1))}
                fmt.Println(v110, v110, v3)
                fmt.Println(v4, v4, v88)
                _ = v121
            }
            for fi122 := 0; fi122 < 7; fi122++ {
                var v123 uint = (v3 % ((v4) | 1))
                if v1 {
                    if (((v2 < (v2 / (((v2 + (v2 - (v110 & v110)))) | 1))) && (!false)) && true) {
                        v2 += v110
                        switch v15 {
                        case 9223372036854775807:
                            fmt.Println(v4, fi122)
                        case 42:
                            fmt.Println(v110, v110)
                        default:
                            _ = len(v109)
                        }
                        v124 := GenS11{F0: -1, F1: v4}
                        _ = v124
                    } else {
                        for fi125 := 0; fi125 < 17; fi125++ {
                            switch "a" {
                            case "a":
                                var v126 uint = v3
                                switch "hello" {
                                case "":
                                    if ((v88 < (v88 % ((v88) | 1))) && (516 >= 2)) {
                                        v127 := GenS12{F0: uint64(v2), F1: ("foo bar" + "hello"), F2: v1}
                                        var v128 string = "foo bar"
                                        _ = v127
                                        _ = v128
                                    }
                                default:
                                    switch (v15 & fi122) {
                                    case -9223372036854775808:
                                        v129 := GenS2{F0: uint32(v123), F1: 9223372036854775807, F2: 949}
                                        _ = v129
                                    default:
                                        v130 := GenS5{F0: false, F1: v126}
                                        _ = v130
                                    }
                                }
                                _ = v126
                            case "":
                                v131 := GenS13{F0: 300, F1: 32767, F2: 193}
                                _ = v131
                            case "\n\t":
                                v123 = 924
                                if len(v109) > 0 {
                                    v109[0] = (fi125 ^ (-(fi122)))
                                }
                                v88 = uint64(v110)
                            default:
                                if v1 {
                                    var v132 int16 = -938
                                    var v133 string = "a"
                                    v134 := GenS8{F0: v3}
                                    _ = v132
                                    _ = v133
                                    _ = v134
                                }
                                if len(v109) >= 2 {
                                    v109 = v109[1:len(v109)]
                                }
                                var v135 string = "foo bar"
                                _ = v135
                            }
                            if true {
                                var v136 float64 = v14
                                _ = v136
                            } else {
                                switch (fi125 % ((fi125) | 1)) {
                                case -174:
                                    if true {
                                        var v137 int = fi125
                                        _ = v137
                                    } else {
                                        fmt.Println(v123)
                                    }
                                }
                                v2 += (v110 & v110)
                            }
                            _ = len(v109)
                            _ = fi125
                        }
                        for fi138 := 0; fi138 < 6; fi138++ {
                            if v1 {
                                v139 := GenS10{F0: 128}
                                _ = v139
                            } else {
                                if v1 {
                                    v140 := GenS11{F0: 9223372036854775807, F1: v3}
                                    v88 += v88
                                    v141 := GenS11{F0: -435, F1: (v3 | (v123 << (((v4 % ((v3) | 1))) & 63)))}
                                    _ = v140
                                    _ = v141
                                } else {
                                    _ = len(v109)
                                    var v142 uint8 = 138
                                    fmt.Println(fi138, fi122)
                                    fmt.Println(v142)
                                    v143 := GenS14{F0: (v2 & (v110 - 526)), F1: (^(v2))}
                                    _ = v142
                                    _ = v143
                                }
                                var v144 bool = true
                                switch 32767 {
                                case -1:
                                    var v145 string = ("a" + (("a" + ("foo bar" + (((("foo bar" + ("foo bar" + "\n\t")) + ("\n\t" + ((("foo bar" + "a") + ("foo bar" + ((("foo bar" + ((("a" + ("foo bar" + ((("hello" + "\n\t") + "a") + "\n\t"))) + ("" + "\n\t")) + ("a" + "hello"))) + "hello") + ("" + "foo bar")))) + "\n\t"))) + ("\n\t" + "")) + ("a" + ("a" + ""))))) + ("hello" + ("foo bar" + ("" + "foo bar")))))
                                    _ = v145
                                case -817:
                                    fmt.Println(v123)
                                case 42:
                                    if ((-9223372036854775808 == v15) && v144) {
                                        v146 := GenS7{F0: 2147483647, F1: ("foo bar" + ("\n\t" + "\n\t"))}
                                        _ = v146
                                    } else {
                                        var v147 float32 = (0.0 - 3.14159)
                                        var v148 float64 = v14
                                        _ = v147
                                        _ = v148
                                    }
                                }
                                _ = v144
                            }
                            if v1 {
                                v14 = 3.14159
                                var v149 uint8 = uint8(v3)
                                if true {
                                    if (v14 != v14) {
                                        if v1 {
                                            if (((false && (!v1)) && (((0.1 * 0.1) / float32(v14)) <= 0.1)) || v1) {
                                                v1 = v1
                                                fmt.Println(v149, v123)
                                            } else {
                                                if v1 {
                                                    switch (("\n\t" + "foo bar") + (("" + (("" + "a") + "foo bar")) + ("a" + "hello"))) {
                                                    case "hello":
                                                        v150 := GenS13{F0: 1, F1: 2, F2: v149}
                                                        _ = v150
                                                    case "foo bar":
                                                        var v151 int16 = 42
                                                        _ = v151
                                                    case "\n\t":
                                                        _ = len(v109)
                                                    default:
                                                        fmt.Println(v15)
                                                    }
                                                    v109 = append(v109, v15)
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v152 uint16 = 42
                                                        _ = v152
                                                    case "foo bar":
                                                        var v153 int8 = 42
                                                        _ = v153
                                                    default:
                                                        v154 := GenS2{F0: (v2 ^ 2147483647), F1: (fi138 + fi122), F2: 118}
                                                        _ = v154
                                                    }
                                                    v155 := GenS6{F0: (!(true && (-68 >= -128))), F1: 42, F2: (v123 << ((v3) & 63))}
                                                    var v156 uint = (v123 & 18446744073709551615)
                                                    _ = v155
                                                    _ = v156
                                                } else {
                                                    v157 := GenS15{F0: ("\n\t" + "foo bar"), F1: (-(fi138)), F2: v149}
                                                    if len(v109) >= 2 {
                                                        v109 = v109[1:len(v109)]
                                                    }
                                                    _ = v157
                                                }
                                                var v158 int64 = 0
                                                if len(v109) >= 2 {
                                                    v109 = v109[1:len(v109)]
                                                }
                                                v159 := GenS1{F0: 2, F1: 2, F2: v3}
                                                _ = v158
                                                _ = v159
                                            }
                                        } else {
                                            var v160 uint16 = 284
                                            fmt.Println(v123, v110, v149)
                                            var v161 int16 = -857
                                            _ = v160
                                            _ = v161
                                        }
                                        _ = len(v109)
                                    }
                                    var v162 uint64 = 1
                                    _ = v162
                                }
                                _ = v149
                            }
                            switch int(v3) {
                            case -1:
                                var v163 int8 = 100
                                _ = v163
                            case 32767:
                                v110 = uint32(v3)
                            }
                            _ = fi138
                        }
                        for fi164 := 0; fi164 < 18; fi164++ {
                            var v165 int64 = 1
                            _ = fi164
                            _ = v165
                        }
                    }
                    v166 := GenS16{F0: (fi122 >> (((v123 % ((v3) | 1))) & 63)), F1: (v1 && ((0.0 * v14) != 0.1))}
                    _ = v166
                } else {
                    var v167 float64 = v14
                    _ = v167
                }
                _ = fi122
                _ = v123
            }
        }
        if len(v109) > 0 {
            v109[0] = v15
        }
        v4 += v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v109:", len(v109), v109)
    fmt.Println("v110:", v110)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v109
    _ = v110
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v88
}
