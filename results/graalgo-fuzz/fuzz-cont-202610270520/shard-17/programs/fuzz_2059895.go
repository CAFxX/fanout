package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS3 struct { F0 string; F1 string; F2 int64 }
type GenS4 struct { F0 int64; F1 uint64; F2 int32 }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 int8; F1 int16; F2 int }
type GenS8 struct { F0 uint8; F1 int }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 bool; F1 bool; F2 int8 }
type GenS12 struct { F0 int16; F1 bool; F2 string }
type GenS13 struct { F0 uint32; F1 uint8 }
type GenS14 struct { F0 bool; F1 uint32 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 uint32; F1 uint8 }
type GenS19 struct { F0 int16; F1 int32; F2 int32 }
type GenS20 struct { F0 int8; F1 int64; F2 uint64 }
type GenS21 struct { F0 uint64; F1 uint }
func genFn0(p0 bool) bool {
    fmt.Println(p0)
    v4 := []int{42}
    switch "" {
    case "hello":
        fmt.Println(p0)
    case "\n\t":
        fmt.Println(p0)
    case "":
        switch ("\n\t" + ("" + ("hello" + "hello"))) {
        case "":
            p0 = (965 >= 42)
        }
    }
    var v5 float64 = 0.0
    v5 = (-2.25 * ((0.1 + v5) * (0.0 + (3.14159 / 3.14159))))
    _ = p0
    _ = v4
    _ = v5
    return true
}
func genFn1(p0 uint8) int16 {
    var v6 int8 = -1
    for fi7 := 0; fi7 < 18; fi7++ {
        if (!((2 > 902) || false)) {
            var v8 uint16 = 42
            if (true || false) {
                v9 := []int{109}
                p0 = 1
                _ = v9
            } else {
                for fi10 := 0; fi10 < 8; fi10++ {
                    v11 := GenS1{F0: 128}
                    v12 := []int{-1, (^(fi10)), fi10}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
            if (599 > 166) {
                v13 := []int{}
                _ = v13
            }
            _ = v8
        } else {
            v14 := []int{(fi7 << ((2) & 63)), fi7, 1, (-(fi7))}
            p0 = p0
            _ = v14
        }
        var v15 int64 = -1
        switch (fi7 ^ int(v6)) {
        case 127:
            var v16 int = 2147483647
            _ = v16
        default:
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 float32 = 1.5
                v19 := GenS2{F0: 191, F1: -1, F2: 127}
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        _ = fi7
        _ = v15
    }
    if false {
        for fi20 := 0; fi20 < 16; fi20++ {
            var v21 int64 = -1
            _ = fi20
            _ = v21
        }
        if false {
            var v22 string = ""
            if false {
                if false {
                    var v23 int32 = -905
                    _ = v23
                } else {
                    var v24 uint32 = 2147483648
                    fmt.Println(v22)
                    for fi25 := 0; fi25 < 5; fi25++ {
                        if (0.0 == 0.0) {
                            if (!(true || (127 >= -1))) {
                                if true {
                                    fmt.Println(v22)
                                    for fi26 := 0; fi26 < 12; fi26++ {
                                        v24 = v24
                                        if false {
                                            var v27 int64 = 272
                                            v28 := []int{-9223372036854775808, fi26, (fi26 / ((32768) | 1))}
                                            fmt.Println(v22, v27, v6)
                                            _ = v27
                                            _ = v28
                                        } else {
                                            v29 := GenS1{F0: 128}
                                            _ = v29
                                        }
                                        _ = fi26
                                    }
                                }
                                v30 := GenS3{F0: (("hello" + v22) + ("a" + ("a" + "a"))), F1: v22, F2: 1}
                                v31 := GenS4{F0: -67, F1: 97, F2: -1}
                                _ = v30
                                _ = v31
                            } else {
                                switch fi25 {
                                case -9223372036854775808:
                                    if ((v22 + "foo bar") != ("\n\t" + v22)) {
                                        v32 := GenS4{F0: 1, F1: 18446744073709551615, F2: 0}
                                        v33 := []int{1, fi25, fi25}
                                        _ = v32
                                        _ = v33
                                    } else {
                                        v34 := []int{}
                                        if len(v34) >= 2 {
                                            v34 = v34[1:len(v34)]
                                        }
                                        v34 = append(v34, fi25)
                                        _ = v34
                                    }
                                case 1:
                                    if (((v22 + v22) <= "a") || false) {
                                        var v35 uint64 = 1
                                        _ = v35
                                    }
                                case -8:
                                    v6 += (v6 << ((526) & 63))
                                default:
                                    v24 = v24
                                }
                                fmt.Println(fi25, v22, p0)
                                v36 := GenS5{F0: (fi25 / ((2147483647) | 1)), F1: int16(v24)}
                                _ = v36
                            }
                            if (!true) {
                                if ((!true) || (1 != 18446744073709551615)) {
                                    v22 = ""
                                }
                                fmt.Println(fi25, v6, p0)
                            } else {
                                var v37 uint64 = 18446744073709551615
                                v6 += (v6 - (v6 << ((255) & 63)))
                                if (true && ((v6 >> ((1) & 63)) <= (v6 - (^(v6))))) {
                                    v6 -= v6
                                } else {
                                    v22 = v22
                                    v22 = v22
                                }
                                _ = v37
                            }
                            var v38 int64 = 42
                            _ = v38
                        } else {
                            for fi39 := 0; fi39 < 14; fi39++ {
                                v40 := GenS1{F0: 62}
                                if ("foo bar" <= ("foo bar" + "a")) {
                                    v22 = "foo bar"
                                    v22 = ((v22 + v22) + "foo bar")
                                }
                                v41 := []int{(fi25 % ((fi25) | 1)), (fi39 << ((685) & 63)), (fi25 * fi39), fi25}
                                _ = fi39
                                _ = v40
                                _ = v41
                            }
                            v42 := []int{fi25}
                            _ = v42
                        }
                        if (0 >= 605) {
                            for fi43 := 0; fi43 < 7; fi43++ {
                                var v44 uint64 = 0
                                _ = fi43
                                _ = v44
                            }
                            v45 := []int{-1, (fi25 + fi25)}
                            _ = v45
                        }
                        var v46 int8 = v6
                        _ = fi25
                        _ = v46
                    }
                    _ = v24
                }
                switch "hello" {
                case "\n\t":
                    fmt.Println(v6)
                case "foo bar":
                    v47 := []int{42, -9223372036854775808, 9223372036854775807, 327}
                    _ = v47
                default:
                    switch (v22 + (v22 + v22)) {
                    case "":
                        fmt.Println(v6, v22, v22)
                    case "a":
                        v48 := GenS6{F0: 42, F1: 0}
                        _ = v48
                    default:
                        var v49 uint = 42
                        _ = v49
                    }
                }
                for fi50 := 0; fi50 < 14; fi50++ {
                    if (true && (!(!(!false)))) {
                        var v51 int64 = 9223372036854775807
                        v51 = v51
                        var v52 uint = 0
                        _ = v51
                        _ = v52
                    }
                    v53 := GenS1{F0: (p0 * 42)}
                    _ = fi50
                    _ = v53
                }
            }
            p0 += (p0 - p0)
            _ = v22
        } else {
            switch 0 {
            case 127:
                var v54 int8 = v6
                _ = v54
            case 4294967295:
                var v55 uint16 = 255
                _ = v55
            }
            p0 -= 128
            if (!(-9223372036854775808 < 9223372036854775807)) {
                v6 = v6
                if true {
                    v56 := []int{1, -1}
                    v6 = v6
                    _ = v56
                } else {
                    var v57 int8 = (v6 / ((-1) | 1))
                    var v58 int32 = -911
                    _ = v57
                    _ = v58
                }
                if (true || (-2147483648 > -639)) {
                    v6 = (v6 / ((v6) | 1))
                }
            } else {
                fmt.Println(p0)
            }
        }
        var v59 float32 = 0.1
        _ = v59
    }
    if (false && false) {
        var v60 int = -2147483648
        for fi61 := 0; fi61 < 10; fi61++ {
            v6 = (v6 / (((v6 << ((65535) & 63))) | 1))
            _ = fi61
        }
        _ = v60
    } else {
        v62 := []int{1, -1}
        _ = v62
    }
    _ = p0
    _ = v6
    return 261
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 0
    var v3 uint8 = 2
    fmt.Println(v3, v3)
    v3 = 93
    switch "" {
    case "foo bar":
        var v63 bool = false
        _ = v63
    default:
        v2 += 128
    }
    fmt.Println(v2)
    v64 := GenS5{F0: 9223372036854775807, F1: 32767}
    v65 := GenS3{F0: (((("\n\t" + "a") + "a") + "hello") + "a"), F1: "hello", F2: -9223372036854775808}
    v66 := GenS7{F0: 100, F1: 705, F2: -347}
    v3 = (v2 / (((v2 * (v3 % ((128) | 1)))) | 1))
    fmt.Println(v3)
    v67 := GenS8{F0: (v2 * v2), F1: -967}
    if (18446744073709551615 > 4294967295) {
        var v68 uint = 255
        v69 := GenS8{F0: 239, F1: 32768}
        switch "a" {
        case "":
            v70 := v1[-365]
            _ = v70
        case "hello":
            if ("a" >= "foo bar") {
                v71 := GenS4{F0: 0, F1: 192, F2: 42}
                v3 += (v3 << (((v68 << ((v68) & 63))) & 63))
                _ = v71
            } else {
                if true {
                    for fi72 := 0; fi72 < 20; fi72++ {
                        v2 += v3
                        var v73 int8 = 100
                        var v74 int64 = -359
                        _ = fi72
                        _ = v73
                        _ = v74
                    }
                    switch -822 {
                    case 32768:
                        switch 472 {
                        case 128:
                            v75 := []int{-2147483648, 42, -761}
                            _ = v75
                        case 372:
                            v76 := GenS9{F0: 2147483648}
                            _ = v76
                        case 127:
                            if ((false || ((!(895 <= 32767)) && ((("foo bar" + "\n\t") <= "a") || false))) && false) {
                                fmt.Println(v68, v3, v68)
                                if (!(!true)) {
                                    var v77 uint = (v68 << ((42) & 63))
                                    v68 -= v68
                                    _ = v77
                                }
                            } else {
                                var v78 uint64 = 18446744073709551615
                                v2 += (^(v2))
                                _ = v78
                            }
                        default:
                            v3 = 127
                        }
                    case 620:
                        v68 = (^(v68))
                    }
                }
                if false {
                    v79 := GenS10{F0: 128}
                    if true {
                        v68 = (v68 + (v68 / ((4294967295) | 1)))
                        var v80 int8 = 1
                        v81 := GenS11{F0: false, F1: (!false), F2: -128}
                        _ = v80
                        _ = v81
                    }
                    switch 127 {
                    case -9223372036854775808:
                        v82 := GenS12{F0: 32767, F1: (190 != 146), F2: ("a" + "foo bar")}
                        _ = v82
                    case -143:
                        v68 = v68
                    }
                    _ = v79
                }
            }
        case "\n\t":
            if (true && (3.14159 > -2.25)) {
                var v83 int8 = 100
                var v84 int8 = (v83 - (v83 + (v83 / ((v83) | 1))))
                for fi85 := 0; fi85 < 3; fi85++ {
                    switch fi85 {
                    case -2147483648:
                        switch "foo bar" {
                        case "\n\t":
                            fmt.Println(v3, v83)
                        case "":
                            if (!(v83 <= (-(v84)))) {
                                for fi86 := 0; fi86 < 20; fi86++ {
                                    v87 := GenS6{F0: 1, F1: (-(fi86))}
                                    var v88 int = (fi86 + (^(fi85)))
                                    _ = fi86
                                    _ = v87
                                    _ = v88
                                }
                                fmt.Println(v83)
                            } else {
                                for fi89 := 0; fi89 < 5; fi89++ {
                                    var v90 int = -613
                                    _ = fi89
                                    _ = v90
                                }
                                fmt.Println(v83)
                                var v91 uint32 = 42
                                _ = v91
                            }
                        case "hello":
                            switch fi85 {
                            case 4294967295:
                                v68 = 255
                            case -1:
                                switch "hello" {
                                case "hello":
                                    v68 = v68
                                case "":
                                    v92 := []int{9223372036854775807, 2147483647}
                                    _ = v92
                                }
                            case 951:
                                v93 := GenS4{F0: -9223372036854775808, F1: 808, F2: 738}
                                _ = v93
                            }
                        default:
                            var v94 float32 = (3.14159 / (3.14159 - 1.5))
                            _ = v94
                        }
                    case -9223372036854775808:
                        if ((fi85 >= fi85) && ("\n\t" == ("\n\t" + "\n\t"))) {
                            v95 := GenS7{F0: 42, F1: -1, F2: (fi85 % ((0) | 1))}
                            v1[(^(fi85))] = 2
                            _ = v95
                        } else {
                            for fi96 := 0; fi96 < 8; fi96++ {
                                if false {
                                    v2 = (v3 * 53)
                                    var v97 int32 = 0
                                    _ = v97
                                }
                                _ = fi96
                            }
                        }
                    case 842:
                        fmt.Println(v68, v84, v84)
                    default:
                        var v98 int = 32767
                        _ = v98
                    }
                    v68 = (v68 << ((564) & 63))
                    _ = fi85
                }
                _ = v83
                _ = v84
            } else {
                v99 := GenS2{F0: 2, F1: 0, F2: uint8(v68)}
                if ((!false) && false) {
                    if true {
                        var v100 int32 = -1
                        switch "hello" {
                        case "foo bar":
                            if true {
                                if (2 <= 377) {
                                    var v101 float32 = (3.14159 + float32((0.1 + 1.5)))
                                    if false {
                                        v100 = v100
                                        fmt.Println(v100, v2)
                                    }
                                    _ = v101
                                } else {
                                    v102 := GenS4{F0: 42, F1: 42, F2: 1}
                                    var v103 uint32 = 4294967295
                                    _ = v102
                                    _ = v103
                                }
                                var v104 int64 = -591
                                _ = v104
                            } else {
                                if ("\n\t" <= "hello") {
                                    var v105 int16 = 39
                                    v106 := GenS13{F0: 398, F1: (v3 + (v2 * 255))}
                                    for fi107 := 0; fi107 < 12; fi107++ {
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v108 uint16 = 42
                                            _ = v108
                                        case "":
                                            if (1 <= (v105 << (((v68 % ((v68) | 1))) & 63))) {
                                                switch "hello" {
                                                case "\n\t":
                                                    v109 := GenS4{F0: 0, F1: 0, F2: (v100 - (v100 - (v100 * v100)))}
                                                    _ = v109
                                                case "hello":
                                                    if false {
                                                        v110 := []int{(fi107 | fi107), (fi107 | (fi107 % ((fi107) | 1))), (fi107 - (^(fi107))), (-(fi107))}
                                                        var v111 float32 = float32(0.0)
                                                        var v112 int32 = (v100 | 0)
                                                        _ = v110
                                                        _ = v111
                                                        _ = v112
                                                    } else {
                                                        v3 = 99
                                                    }
                                                }
                                                v113 := GenS14{F0: ((!(false && (false && (!true)))) && (false || true)), F1: 309}
                                                var v114 uint16 = 255
                                                _ = v113
                                                _ = v114
                                            }
                                        }
                                        var v115 float64 = 1e10
                                        v68 = v68
                                        _ = fi107
                                        _ = v115
                                    }
                                    _ = v105
                                    _ = v106
                                } else {
                                    for fi116 := 0; fi116 < 14; fi116++ {
                                        if false {
                                            v3 -= v3
                                            v68 = (v68 - 1)
                                            v1[int(v3)] = 2147483647
                                        } else {
                                            v100 = v100
                                            for fi117 := 0; fi117 < 1; fi117++ {
                                                var v118 float64 = 3.14159
                                                _ = fi117
                                                _ = v118
                                            }
                                            for fi119 := 0; fi119 < 9; fi119++ {
                                                v120 := []int{fi119, 32768}
                                                v3 = 2
                                                v121 := GenS3{F0: "hello", F1: "\n\t", F2: 42}
                                                var v122 uint = v68
                                                _ = fi119
                                                _ = v120
                                                _ = v121
                                                _ = v122
                                            }
                                        }
                                        switch ("hello" + ("a" + "foo bar")) {
                                        case "a":
                                            v123 := v1[(fi116 ^ (^(fi116)))]
                                            _ = v123
                                        case "hello":
                                            switch "\n\t" {
                                            case "":
                                                v124 := GenS15{F0: 42}
                                                _ = v124
                                            }
                                        case "":
                                            var v125 int32 = v100
                                            _ = v125
                                        default:
                                            for fi126 := 0; fi126 < 4; fi126++ {
                                                if (!((v3 + 178) == (v3 * v2))) {
                                                    v127 := []int{(^(fi126)), (fi126 ^ (fi116 * 9223372036854775807)), (fi126 - fi126)}
                                                    v128 := GenS16{F0: 65535, F1: -9223372036854775808, F2: 379}
                                                    switch (("" + ("foo bar" + "\n\t")) + "") {
                                                    case "a":
                                                        v68 -= (v68 & v68)
                                                    case "foo bar":
                                                        if len(v127) >= 2 {
                                                            v127 = v127[1:len(v127)]
                                                        }
                                                    default:
                                                        v129 := GenS17{F0: v100}
                                                        _ = v129
                                                    }
                                                    _ = v127
                                                    _ = v128
                                                }
                                                if false {
                                                    var v130 uint8 = 255
                                                    var v131 uint32 = 65535
                                                    v132 := []int{(fi116 & (^(fi126))), (fi126 - 4294967295)}
                                                    v133 := GenS10{F0: (fi126 % (((fi126 ^ fi116)) | 1))}
                                                    var v134 float64 = (((3.14159 - float64(1.5)) * 1.5) - (0.0 - 1e10))
                                                    _ = v130
                                                    _ = v131
                                                    _ = v132
                                                    _ = v133
                                                    _ = v134
                                                } else {
                                                    if true {
                                                        var v135 uint16 = 247
                                                        switch 127 {
                                                        case -2147483648:
                                                            var v136 int = fi126
                                                            _ = v136
                                                        case -1:
                                                            switch (fi116 | fi116) {
                                                            case 4294967295:
                                                                if false {
                                                                    v137 := []int{(fi116 | fi126)}
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        if len(v137) >= 2 {
                                                                            v137 = v137[1:len(v137)]
                                                                        }
                                                                    case "\n\t":
                                                                        v2 += (v2 * v3)
                                                                    case "foo bar":
                                                                        v100 = int32(fi126)
                                                                    default:
                                                                        var v138 float32 = (((3.14159 - (0.0 / float32(1.5))) + 0.0) - (1e10 / float32(3.14159)))
                                                                        _ = v138
                                                                    }
                                                                    if true {
                                                                        var v139 uint8 = 178
                                                                        if true {
                                                                            var v140 float64 = float64(1.5)
                                                                            v141 := GenS18{F0: 65535, F1: (v3 & (v139 % ((128) | 1)))}
                                                                            _ = v140
                                                                            _ = v141
                                                                        }
                                                                        _ = v139
                                                                    } else {
                                                                        fmt.Println(v3)
                                                                        var v142 float64 = ((((((1e10 / 3.14159) - -2.25) * 0.1) - float64(-2.25)) + (1e10 + float64(0.1))) * (-2.25 + -2.25))
                                                                        _ = v142
                                                                    }
                                                                    var v143 float64 = -2.25
                                                                    var v144 uint32 = 2147483647
                                                                    if false {
                                                                        v68 = (v68 >> (((v68 << ((v68) & 63))) & 63))
                                                                    }
                                                                    if (!(v143 > v143)) {
                                                                        if false {
                                                                            v144 = v144
                                                                            var v145 uint = (v68 * (v68 - (v68 * (v68 / ((v68) | 1)))))
                                                                            if len(v137) > 0 {
                                                                                v137[0] = fi116
                                                                            }
                                                                            _ = v145
                                                                        } else {
                                                                            if (((!(0 < 248)) || (!((false || false) && ((false && (-128 == -6)) || false)))) && (42 < (^(v68)))) {
                                                                                var v146 int16 = genFn1((v3 | v3))
                                                                                fmt.Println(v2, v135, v100)
                                                                                var v147 uint32 = v144
                                                                                _ = v146
                                                                                _ = v147
                                                                            }
                                                                            var v148 int8 = -1
                                                                            _ = v148
                                                                        }
                                                                        var v149 float32 = 1e10
                                                                        if true {
                                                                            var v150 int64 = 42
                                                                            var v151 uint = (v68 - v68)
                                                                            _ = v150
                                                                            _ = v151
                                                                        }
                                                                        var v152 uint32 = v144
                                                                        _ = v149
                                                                        _ = v152
                                                                    }
                                                                    _ = v137
                                                                    _ = v143
                                                                    _ = v144
                                                                }
                                                            case 10:
                                                                var v153 uint8 = (v3 + v2)
                                                                _ = v153
                                                            case -544:
                                                                v154 := GenS18{F0: 42, F1: (v2 | (v3 - 44))}
                                                                if (((v100 >> ((v68) & 63)) > v100) || true) {
                                                                    if true {
                                                                        var v155 uint16 = v135
                                                                        v135 = v135
                                                                        _ = v155
                                                                    }
                                                                    var v156 bool = (true || true)
                                                                    var v157 uint16 = (^(v135))
                                                                    var v158 bool = (1 < -774)
                                                                    _ = v156
                                                                    _ = v157
                                                                    _ = v158
                                                                } else {
                                                                    var v159 uint64 = 959
                                                                    _ = v159
                                                                }
                                                                _ = v154
                                                            default:
                                                                v135 = (v135 * v135)
                                                            }
                                                            v135 -= v135
                                                            var v160 int64 = -1
                                                            _ = v160
                                                        default:
                                                            var v161 uint16 = (v135 * v135)
                                                            v100 -= (v100 / (((v100 - (v100 * (v100 - v100)))) | 1))
                                                            _ = v161
                                                        }
                                                        delete(v1, -2147483648)
                                                        _ = v135
                                                    } else {
                                                        switch 2147483647 {
                                                        case 144:
                                                            v68 = (v68 << ((v68) & 63))
                                                        case 128:
                                                            var v162 uint16 = 32767
                                                            _ = v162
                                                        case 9223372036854775807:
                                                            v163 := GenS9{F0: 1}
                                                            _ = v163
                                                        default:
                                                            var v164 int8 = -1
                                                            var v165 float64 = (float64(-2.25) / 1.5)
                                                            v166 := v1[4294967295]
                                                            fmt.Println(v100, v165)
                                                            if (!(2 <= uint(v2))) {
                                                                switch ("foo bar" + "") {
                                                                case "\n\t":
                                                                    v166 += 458
                                                                case "hello":
                                                                    v167 := GenS12{F0: 32767, F1: ((fi116 != fi126) || false), F2: ("\n\t" + "a")}
                                                                    _ = v167
                                                                case "":
                                                                    if true {
                                                                        var v168 uint64 = 4294967295
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            fmt.Println(v168, v166)
                                                                        case "\n\t":
                                                                            v68 = 892
                                                                        }
                                                                        v1[int(v68)] = (v166 >> ((2) & 63))
                                                                        _ = v168
                                                                    } else {
                                                                        v2 = v3
                                                                    }
                                                                }
                                                                if false {
                                                                    v169 := GenS19{F0: int16(fi126), F1: v100, F2: (v100 % (((^(v100))) | 1))}
                                                                    v170 := GenS20{F0: (v164 >> ((v68) & 63)), F1: -9223372036854775808, F2: 1}
                                                                    _ = v169
                                                                    _ = v170
                                                                }
                                                            }
                                                            _ = v164
                                                            _ = v165
                                                            _ = v166
                                                        }
                                                        v171 := GenS20{F0: 127, F1: -9223372036854775808, F2: 602}
                                                        fmt.Println(v3, v100, v2)
                                                        _ = v171
                                                    }
                                                }
                                                _ = fi126
                                            }
                                        }
                                        var v172 uint8 = (v2 * 1)
                                        _ = fi116
                                        _ = v172
                                    }
                                }
                                if (false || (((1.5 + 3.14159) + (-2.25 - (((3.14159 / 1.0) + 0.0) * -2.25))) <= 1.5)) {
                                    v2 += v2
                                    fmt.Println(v3, v2, v2)
                                }
                            }
                        case "hello":
                            var v173 float64 = (0.1 + (0.0 * -2.25))
                            _ = v173
                        }
                        v68 -= (v68 >> ((375) & 63))
                        _ = v100
                    } else {
                        switch "hello" {
                        case "hello":
                            if ((-1 <= -258) && (0 >= 1)) {
                                var v174 uint8 = v2
                                fmt.Println(v68)
                                _ = v174
                            } else {
                                var v175 bool = true
                                v1[1] = 517
                                fmt.Println(v2)
                                _ = v175
                            }
                        case "a":
                            v68 -= 255
                        default:
                            var v176 float64 = 1.5
                            _ = v176
                        }
                        v177 := []int{0, 0, 32768, 32767}
                        _ = v177
                    }
                    var v178 int = -2147483648
                    _ = v178
                } else {
                    fmt.Println(v3, v3, v3)
                }
                _ = v99
            }
        default:
            v179 := GenS2{F0: 194, F1: 42, F2: v2}
            _ = v179
        }
        _ = v68
        _ = v69
    }
    v180 := []int{32767, 4294967295}
    var v181 uint64 = 4294967295
    var v182 int = 1
    v183 := GenS21{F0: (v181 + (^(v181))), F1: uint(v181)}
    fmt.Println(v182, v181, v2)
    switch 128 {
    case -9223372036854775808:
        for fi184 := 0; fi184 < 5; fi184++ {
            var v185 float64 = float64(((3.14159 / 1.5) - 0.0))
            delete(v1, fi184)
            _ = fi184
            _ = v185
        }
    }
    var v186 float64 = float64((1e10 * (-2.25 / 1.5)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v180:", len(v180), v180)
    fmt.Println("v181:", v181)
    fmt.Println("v182:", v182)
    fmt.Println("v186:", v186)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v180
    _ = v181
    _ = v182
    _ = v183
    _ = v186
    _ = v2
    _ = v3
    _ = v64
    _ = v65
    _ = v66
    _ = v67
}
