package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int32; F2 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 string; F1 int64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int; F1 int64 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int64; F1 int16; F2 int64 }
type GenS13 struct { F0 uint; F1 uint16 }
type GenS14 struct { F0 int32; F1 uint16 }
type GenS15 struct { F0 uint; F1 bool }
type GenS16 struct { F0 int64; F1 string }
type GenS17 struct { F0 int }
type GenS18 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint64, p1 bool) int32 {
    var v5 int8 = -1
    var v6 int32 = -1
    var v7 int16 = 0
    if (!(!(false || (0 != int64(v5))))) {
        switch "hello" {
        case "\n\t":
            v6 = v6
        case "a":
            v8 := []int{4294967295}
            _ = v8
        case "":
            fmt.Println(v5, v7)
        default:
            if p1 {
                for fi9 := 0; fi9 < 2; fi9++ {
                    switch 42 {
                    case -9223372036854775808:
                        if (v6 >= v6) {
                            v6 -= v6
                            for fi10 := 0; fi10 < 7; fi10++ {
                                var v11 int32 = (v6 % (((^(v6))) | 1))
                                v6 = (v11 % ((256) | 1))
                                var v12 bool = p1
                                _ = fi10
                                _ = v11
                                _ = v12
                            }
                            v6 = 0
                        }
                    case 127:
                        if (true && (42 > -1)) {
                            switch (("\n\t" + "") + "foo bar") {
                            case "":
                                v13 := []int{(fi9 / ((fi9) | 1)), fi9, fi9}
                                _ = v13
                            case "a":
                                var v14 float32 = 0.1
                                _ = v14
                            default:
                                for fi15 := 0; fi15 < 18; fi15++ {
                                    v16 := []int{}
                                    _ = fi15
                                    _ = v16
                                }
                            }
                            v17 := GenS1{F0: 2}
                            fmt.Println(fi9)
                            _ = v17
                        }
                    default:
                        if (false || ((v6 <= (v6 ^ (v6 + -348))) && ((p0 % ((p0) | 1)) > p0))) {
                            v5 = (v5 / (((-(v5))) | 1))
                            v18 := []int{(fi9 | -882), fi9}
                            if p1 {
                                p1 = p1
                                var v19 uint = 646
                                _ = v19
                            }
                            _ = v18
                        }
                    }
                    fmt.Println(v5, v6, v7)
                    _ = fi9
                }
                v20 := GenS1{F0: 2}
                var v21 string = "hello"
                _ = v20
                _ = v21
            }
        }
        for fi22 := 0; fi22 < 10; fi22++ {
            fmt.Println(fi22)
            v5 += v5
            _ = fi22
        }
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return (-(v6))
}
func genFn1(p0 int, p1 uint, p2 bool) int64 {
    if (((((0.1 - (0.0 / ((1e10 * 0.0) + float64((-2.25 * 3.14159))))) - 1e10) * 3.14159) >= (3.14159 - float64(-2.25))) && (-545 == 2147483647)) {
        switch (-(p0)) {
        case -1:
            for fi23 := 0; fi23 < 14; fi23++ {
                var v24 uint32 = 914
                v25 := []int{32767, p0, -2147483648, fi23}
                if p2 {
                    var v26 uint32 = v24
                    _ = v26
                } else {
                    p2 = p2
                    v27 := GenS2{F0: (p1 & (^(p1)))}
                    _ = v27
                }
                _ = fi23
                _ = v24
                _ = v25
            }
        }
    }
    v28 := GenS3{F0: 0, F1: -2147483648, F2: 18446744073709551615}
    v29 := []int{-2147483648, (p0 >> ((p1) & 63))}
    v29 = append(v29, (p0 | (p0 | p0)))
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    _ = v29
    return -1
}

func main() {
    v1 := []int{0, 4294967295, -1, 42}
    v2 := []string{"a", ""}
    v3 := []uint32{2147483647, 871, 42}
    var v4 float64 = 0.1
    var v30 uint32 = 1
    v30 += v30
    for fi31 := 0; fi31 < 6; fi31++ {
        if (9223372036854775807 < 9223372036854775807) {
            fmt.Println(v30)
        }
        if (false || true) {
            for fi32 := 0; fi32 < 11; fi32++ {
                _ = len(v3)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v33 := GenS4{F0: (fi31 / ((fi32) | 1))}
                _ = fi32
                _ = v33
            }
            v3 = append(v3, v30)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi31
    }
    var v34 uint = 65535
    v30 += 0
    v35 := GenS2{F0: v34}
    switch -467 {
    case 0:
        v2 = append(v2, "hello")
    case 42:
        v4 -= 1e10
    case 32768:
        var v36 bool = true
        _ = v36
    default:
        if true {
            var v37 uint32 = v30
            switch "hello" {
            case "hello":
                v4 = float64((0.0 + 0.0))
            case "\n\t":
                if (484 <= (v34 | v34)) {
                    for fi38 := 0; fi38 < 6; fi38++ {
                        switch (("a" + "a") + "\n\t") {
                        case "a":
                            for fi39 := 0; fi39 < 15; fi39++ {
                                var v40 int8 = -69
                                _ = fi39
                                _ = v40
                            }
                        case "\n\t":
                            var v41 float32 = 1.5
                            _ = v41
                        default:
                            for fi42 := 0; fi42 < 12; fi42++ {
                                var v43 float64 = (v4 / 1.5)
                                var v44 int64 = -9223372036854775808
                                switch "\n\t" {
                                case "foo bar":
                                    if (true || true) {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v45 int = -758
                                        _ = v45
                                    } else {
                                        var v46 uint64 = 705
                                        _ = len(v3)
                                        v34 = (v34 / ((255) | 1))
                                        _ = v46
                                    }
                                    v4 = 1e10
                                    if (2147483647 <= 2147483647) {
                                        v30 = (v30 % ((v37) | 1))
                                        v47 := GenS5{F0: (v30 / (((^(v30))) | 1))}
                                        _ = v47
                                    }
                                default:
                                    v4 = (v43 + ((v43 / (v4 * 1.5)) + v4))
                                }
                                _ = fi42
                                _ = v43
                                _ = v44
                            }
                        }
                        switch "a" {
                        case "hello":
                            if len(v2) > 0 {
                                v2[0] = "\n\t"
                            }
                        case "a":
                            var v48 uint32 = (^(v30))
                            _ = v48
                        case "":
                            var v49 bool = false
                            _ = v49
                        }
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                        _ = fi38
                    }
                    v2 = append(v2, "a")
                    v37 -= v30
                }
            case "a":
                if ((!(!false)) || (32767 < 42)) {
                    for fi50 := 0; fi50 < 17; fi50++ {
                        switch "\n\t" {
                        case "hello":
                            var v51 uint32 = v30
                            _ = v51
                        }
                        var v52 int16 = -319
                        _ = fi50
                        _ = v52
                    }
                }
            default:
                if true {
                    var v53 bool = true
                    _ = v53
                } else {
                    v34 -= v34
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v4 -= v4
                }
            }
            if (2 > 78) {
                v54 := GenS4{F0: 128}
                _ = v54
            } else {
                var v55 float64 = (v4 - v4)
                _ = v55
            }
            _ = v37
        } else {
            var v56 uint16 = 65535
            if (!false) {
                var v57 uint64 = 2
                v3 = append(v3, v30)
                _ = v57
            } else {
                v1 = append(v1, 4294967295)
            }
            _ = v56
        }
    }
    v34 += (v34 << ((2) & 63))
    fmt.Println(v4, v34, v30)
    v58 := GenS1{F0: 0}
    fmt.Println(v34, v34, v30)
    if len(v2) > 0 {
        v2[0] = ""
    }
    _ = len(v3)
    var v59 uint64 = 18446744073709551615
    if true {
        v60 := GenS2{F0: v34}
        switch 32767 {
        case 926:
            v30 = (^(v30))
        case -650:
            if false {
                v61 := GenS6{F0: 1, F1: v59}
                _ = v61
            } else {
                _ = len(v3)
                switch 32767 {
                case 1:
                    for fi62 := 0; fi62 < 15; fi62++ {
                        v1 = append(v1, 1)
                        _ = fi62
                    }
                case 127:
                    for fi63 := 0; fi63 < 15; fi63++ {
                        var v64 int16 = -32768
                        _ = fi63
                        _ = v64
                    }
                }
            }
        case 32768:
            if false {
                fmt.Println(v34, v34)
                if (-1 != -980) {
                    if (true && false) {
                        var v65 int32 = -1
                        if false {
                            if (v34 <= 1) {
                                v30 = 2147483647
                                for fi66 := 0; fi66 < 4; fi66++ {
                                    if len(v3) > 0 {
                                        v3[0] = (v30 - (^(v30)))
                                    }
                                    fmt.Println(v65, fi66)
                                    fmt.Println(v4, v59, fi66)
                                    _ = fi66
                                }
                                fmt.Println(v4, v4, v30)
                            } else {
                                if (uint64(v30) >= v59) {
                                    if true {
                                        var v67 uint64 = (v59 - (^(v59)))
                                        _ = v67
                                    }
                                    for fi68 := 0; fi68 < 13; fi68++ {
                                        for fi69 := 0; fi69 < 6; fi69++ {
                                            v65 = (-(v65))
                                            var v70 int = fi68
                                            v34 += (v34 * uint(v59))
                                            _ = fi69
                                            _ = v70
                                        }
                                        v59 = (^(v59))
                                        _ = fi68
                                    }
                                    v59 = (v59 | 0)
                                }
                                fmt.Println(v4, v59)
                                var v71 float32 = 0.1
                                _ = v71
                            }
                            v72 := GenS4{F0: -2147483648}
                            _ = v72
                        } else {
                            fmt.Println(v34, v65)
                            switch -1 {
                            case -9223372036854775808:
                                if true {
                                    var v73 float64 = (v4 / (v4 - (float64(-2.25) - -2.25)))
                                    _ = v73
                                }
                            case 9223372036854775807:
                                v74 := GenS7{F0: (("" + "a") + ((("foo bar" + "") + (("\n\t" + "foo bar") + "\n\t")) + "foo bar")), F1: -648}
                                _ = v74
                            case 94:
                                v75 := GenS8{F0: (("hello" + "foo bar") + "hello")}
                                _ = v75
                            }
                            v65 += v65
                        }
                        switch "hello" {
                        case "foo bar":
                            v76 := GenS8{F0: (("" + ("foo bar" + "hello")) + "a")}
                            _ = v76
                        case "":
                            if false {
                                var v77 uint64 = 18446744073709551615
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v78 bool = (!true)
                                _ = v77
                                _ = v78
                            } else {
                                if (32767 != int16(v65)) {
                                    switch -2147483648 {
                                    case 2147483647:
                                        var v79 uint64 = (v59 % (((v59 + (v59 - 39))) | 1))
                                        _ = v79
                                    case -32:
                                        var v80 int8 = 127
                                        _ = v80
                                    default:
                                        v81 := GenS1{F0: 65535}
                                        _ = v81
                                    }
                                    var v82 uint8 = 2
                                    _ = v82
                                } else {
                                    var v83 int32 = (v65 % ((-1) | 1))
                                    v30 = v30
                                    switch "\n\t" {
                                    case "":
                                        if false {
                                            v84 := GenS9{F0: int(v30), F1: 0}
                                            v3 = append(v3, v30)
                                            switch 32767 {
                                            case -9223372036854775808:
                                                v85 := GenS10{F0: -77}
                                                _ = v85
                                            case -794:
                                                fmt.Println(v65, v30, v30)
                                            default:
                                                switch ("\n\t" + "\n\t") {
                                                case "":
                                                    for fi86 := 0; fi86 < 14; fi86++ {
                                                        v87 := GenS11{F0: (fi86 + (-(fi86)))}
                                                        _ = fi86
                                                        _ = v87
                                                    }
                                                default:
                                                    var v88 int16 = 0
                                                    _ = v88
                                                }
                                            }
                                            _ = v84
                                        }
                                    case "\n\t":
                                        if false {
                                            _ = len(v3)
                                            v4 -= (v4 / 0.1)
                                        } else {
                                            v65 = v65
                                            if len(v1) > 0 {
                                                v1[0] = -996
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                    case "foo bar":
                                        if (768 <= 937) {
                                            switch "\n\t" {
                                            case "":
                                                var v89 int64 = 9223372036854775807
                                                _ = v89
                                            default:
                                                for fi90 := 0; fi90 < 16; fi90++ {
                                                    for fi91 := 0; fi91 < 5; fi91++ {
                                                        var v92 int64 = genFn1(fi91, v34, (1 == 42))
                                                        var v93 float32 = float32(v4)
                                                        var v94 uint32 = v30
                                                        _ = fi91
                                                        _ = v92
                                                        _ = v93
                                                        _ = v94
                                                    }
                                                    _ = fi90
                                                }
                                            }
                                        } else {
                                            v1 = append(v1, -2147483648)
                                            if ((0.1 + 1e10) == (0.1 / -2.25)) {
                                                if true {
                                                    var v95 float64 = v4
                                                    _ = v95
                                                } else {
                                                    var v96 uint16 = 1
                                                    var v97 float64 = v4
                                                    _ = v96
                                                    _ = v97
                                                }
                                                v4 += ((v4 - (float64(3.14159) * 0.1)) + ((v4 / (v4 - ((v4 / v4) - v4))) + v4))
                                            }
                                            for fi98 := 0; fi98 < 1; fi98++ {
                                                var v99 uint64 = (v59 >> (((^(v34))) & 63))
                                                _ = fi98
                                                _ = v99
                                            }
                                        }
                                    default:
                                        v59 += (v59 * (^(v59)))
                                    }
                                    _ = v83
                                }
                                v59 -= v59
                            }
                        case "\n\t":
                            var v100 uint32 = (v30 % (((v30 | (v30 * uint32(v59)))) | 1))
                            _ = v100
                        default:
                            if (uint32(v34) <= v30) {
                                for fi101 := 0; fi101 < 10; fi101++ {
                                    switch (-(fi101)) {
                                    case 1:
                                        for fi102 := 0; fi102 < 6; fi102++ {
                                            fmt.Println(v4, v65)
                                            v59 = 2
                                            v103 := GenS12{F0: 9223372036854775807, F1: 42, F2: 9223372036854775807}
                                            _ = fi102
                                            _ = v103
                                        }
                                    case -2147483648:
                                        v59 += 4294967295
                                    case 127:
                                        var v104 int8 = 100
                                        _ = v104
                                    }
                                    for fi105 := 0; fi105 < 4; fi105++ {
                                        _ = len(v3)
                                        var v106 uint = 42
                                        _ = fi105
                                        _ = v106
                                    }
                                    var v107 int32 = -2147483648
                                    _ = fi101
                                    _ = v107
                                }
                                var v108 uint = 18446744073709551615
                                _ = v108
                            }
                        }
                        _ = v65
                    } else {
                        if true {
                            v109 := GenS13{F0: v34, F1: 673}
                            _ = v109
                        }
                        v30 += 1
                    }
                    v110 := GenS9{F0: -95, F1: genFn1(-745, 4294967295, false)}
                    _ = v110
                }
            } else {
                for fi111 := 0; fi111 < 16; fi111++ {
                    switch fi111 {
                    case -743:
                        v4 = 1e10
                    default:
                        var v112 uint64 = (v59 / ((v59) | 1))
                        _ = v112
                    }
                    _ = fi111
                }
                switch "a" {
                case "a":
                    for fi113 := 0; fi113 < 7; fi113++ {
                        for fi114 := 0; fi114 < 9; fi114++ {
                            v59 += (v59 + (v59 / ((uint64(v34)) | 1)))
                            _ = fi114
                        }
                        v115 := GenS7{F0: ("a" + (("hello" + ("" + "")) + "")), F1: -1}
                        _ = fi113
                        _ = v115
                    }
                case "hello":
                    v116 := GenS14{F0: -118, F1: 0}
                    _ = v116
                }
                for fi117 := 0; fi117 < 16; fi117++ {
                    var v118 uint = 2
                    if false {
                        if false {
                            for fi119 := 0; fi119 < 13; fi119++ {
                                if false {
                                    fmt.Println(v30, fi119)
                                    v34 -= (v118 ^ 2)
                                    if (true || (!((genFn1((fi117 / (((fi117 << ((v118) & 63))) | 1)), v34, true) >= -9223372036854775808) && true))) {
                                        var v120 uint16 = 255
                                        if ("foo bar" == "a") {
                                            var v121 uint16 = 657
                                            v2 = append(v2, "foo bar")
                                            fmt.Println(v121)
                                            _ = v121
                                        }
                                        v122 := GenS13{F0: v34, F1: 375}
                                        if true {
                                            if ((fi119 & 4294967295) <= (fi117 | int(v120))) {
                                                v59 = 18446744073709551615
                                            } else {
                                                var v123 bool = true
                                                fmt.Println(v34, v123, v30)
                                                _ = v123
                                            }
                                            if (v118 >= (v34 % ((65535) | 1))) {
                                                fmt.Println(v59, v118, v120)
                                                _ = len(v1)
                                                if true {
                                                    var v124 uint = 195
                                                    if len(v2) > 0 {
                                                        v2[0] = ("foo bar" + ("" + ("a" + "hello")))
                                                    }
                                                    _ = v124
                                                } else {
                                                    v59 -= (v59 * v59)
                                                    var v125 int32 = -2147483648
                                                    var v126 uint64 = 1
                                                    _ = v125
                                                    _ = v126
                                                }
                                            } else {
                                                v59 = 1
                                            }
                                        } else {
                                            v127 := GenS9{F0: (fi117 * (fi117 | (fi119 & fi117))), F1: 915}
                                            var v128 uint = (v118 + 42)
                                            _ = v127
                                            _ = v128
                                        }
                                        _ = v120
                                        _ = v122
                                    } else {
                                        switch (fi117 | -225) {
                                        case -60:
                                            v3 = append(v3, uint32(v59))
                                        default:
                                            v129 := GenS2{F0: v118}
                                            _ = v129
                                        }
                                    }
                                }
                                fmt.Println(v118, fi117, v59)
                                v130 := GenS2{F0: v34}
                                v30 = v30
                                v59 -= 4294967295
                                _ = fi119
                                _ = v130
                            }
                        } else {
                            var v131 uint = v34
                            _ = len(v1)
                            _ = v131
                        }
                        v132 := GenS15{F0: 691, F1: ((v59 % ((1) | 1)) <= v59)}
                        _ = v132
                    }
                    var v133 uint8 = 128
                    _ = fi117
                    _ = v118
                    _ = v133
                }
            }
        }
        v59 = v59
        _ = v60
    } else {
        var v134 uint32 = (v30 << ((0) & 63))
        var v135 int16 = 1
        v59 = v59
        _ = v134
        _ = v135
    }
    for fi136 := 0; fi136 < 9; fi136++ {
        v137 := GenS16{F0: 1, F1: ""}
        if (!((2 <= v59) || false)) {
            _ = len(v3)
        } else {
            for fi138 := 0; fi138 < 14; fi138++ {
                v139 := GenS11{F0: -885}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi138
                _ = v139
            }
            switch (fi136 ^ fi136) {
            case 4294967295:
                switch (fi136 + (fi136 & -78)) {
                case 128:
                    if false {
                        fmt.Println(v59, v30, v34)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        switch (fi136 * -9223372036854775808) {
                        case 884:
                            if (2147483647 != -224) {
                                v140 := GenS17{F0: (fi136 & (fi136 / ((127) | 1)))}
                                v34 -= 42
                                if (false || ((!true) || false)) {
                                    v141 := GenS15{F0: 18446744073709551615, F1: (v4 == (v4 * v4))}
                                    _ = v141
                                } else {
                                    var v142 uint64 = uint64(v34)
                                    _ = v142
                                }
                                _ = v140
                            } else {
                                var v143 int16 = -226
                                v144 := GenS8{F0: ""}
                                var v145 int8 = 7
                                _ = v143
                                _ = v144
                                _ = v145
                            }
                        case 457:
                            v30 -= 1
                        }
                        switch (-(fi136)) {
                        case -31:
                            var v146 uint = (^(v34))
                            _ = v146
                        default:
                            if false {
                                switch (fi136 / (((^(fi136))) | 1)) {
                                case 32768:
                                    _ = len(v3)
                                default:
                                    v30 = 321
                                }
                                var v147 uint8 = uint8(v30)
                                _ = v147
                            }
                        }
                        v59 = (v59 - v59)
                    }
                }
            case 128:
                if true {
                    var v148 uint32 = (v30 * (^(v30)))
                    _ = v148
                } else {
                    v149 := GenS13{F0: 255, F1: 496}
                    v59 = v59
                    _ = v149
                }
            case -9223372036854775808:
                v150 := GenS18{F0: 177, F1: 575}
                _ = v150
            }
        }
        _ = fi136
        _ = v137
    }
    v59 = v59
    if (((0.1 * -2.25) > ((((-2.25 * 1.5) / (1.5 - 1e10)) + 0.0) * 0.1)) || false) {
        var v151 uint = v34
        fmt.Println(v30)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v151
    } else {
        var v152 uint64 = 4294967295
        _ = v152
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v4
    _ = v58
    _ = v59
}
