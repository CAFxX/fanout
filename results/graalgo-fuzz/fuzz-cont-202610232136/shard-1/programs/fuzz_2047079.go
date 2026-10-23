package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int; F1 int64; F2 int32 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS6 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int16; F1 bool; F2 uint32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS12 struct { F0 uint64; F1 bool; F2 int64 }
func genFn0() string {
    var v5 uint16 = 1
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS1{F0: "\n\t", F1: 161, F2: uint8(v5)}
        fmt.Println(v5)
        _ = fi6
        _ = v7
    }
    if false {
        v8 := []int{2147483647, 1, 881, 820}
        v9 := GenS1{F0: ("foo bar" + "foo bar"), F1: 1, F2: 0}
        var v10 float32 = (0.1 - 0.0)
        _ = v10
        _ = v8
        _ = v9
    } else {
        v5 += (v5 / (((v5 >> ((637) & 63))) | 1))
        switch 2147483647 {
        case 0:
            for fi11 := 0; fi11 < 1; fi11++ {
                if true {
                    if ((1 <= 42) || (-1 < -516)) {
                        v12 := GenS2{F0: 0, F1: 0}
                        _ = v12
                    } else {
                        v5 += 32767
                        var v13 uint8 = 42
                        if (!((!true) && (0 <= -1))) {
                            if false {
                                var v14 uint = 0
                                for fi15 := 0; fi15 < 11; fi15++ {
                                    v16 := []int{fi15, 32767, (fi11 >> ((v14) & 63)), (fi15 << (((v14 + (^(v14)))) & 63))}
                                    _ = fi15
                                    _ = v16
                                }
                                var v17 uint8 = (^(v13))
                                _ = v14
                                _ = v17
                            } else {
                                v13 -= (v13 - (^(v13)))
                                switch ("" + "\n\t") {
                                case "\n\t":
                                    if false {
                                        var v18 string = "foo bar"
                                        _ = v18
                                    } else {
                                        v19 := []int{fi11, (fi11 & (fi11 - (fi11 << ((1) & 63)))), 127, fi11}
                                        var v20 int = fi11
                                        _ = v19
                                        _ = v20
                                    }
                                case "hello":
                                    var v21 uint = 353
                                    _ = v21
                                }
                                for fi22 := 0; fi22 < 13; fi22++ {
                                    var v23 uint8 = (v13 % (((v13 & v13)) | 1))
                                    _ = fi22
                                    _ = v23
                                }
                            }
                        }
                        _ = v13
                    }
                }
                if (!(((2 != 42) || false) || (-445 == -707))) {
                    v24 := []int{(-(fi11)), (fi11 & (fi11 & -742)), 532}
                    if false {
                        if ("\n\t" >= "") {
                            v25 := GenS3{F0: 868, F1: -1, F2: 42}
                            v26 := GenS4{F0: -352, F1: (fi11 & fi11)}
                            _ = v25
                            _ = v26
                        }
                    } else {
                        v5 = (v5 * (v5 << ((2) & 63)))
                    }
                    _ = v24
                }
                _ = fi11
            }
        case 32767:
            fmt.Println(v5)
        case 2147483647:
            for fi27 := 0; fi27 < 4; fi27++ {
                v28 := []int{345, fi27}
                v29 := GenS3{F0: (fi27 & fi27), F1: 1, F2: 42}
                var v30 int8 = 38
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
        }
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        fmt.Println(fi31, fi31)
        _ = fi31
    }
    for fi32 := 0; fi32 < 2; fi32++ {
        fmt.Println(v5, fi32)
        for fi33 := 0; fi33 < 5; fi33++ {
            v5 += v5
            fmt.Println(v5)
            _ = fi33
        }
        _ = fi32
    }
    _ = v5
    return ""
}

func main() {
    var v1 float32 = (((1.5 + 3.14159) / 1e10) * 3.14159)
    var v2 int8 = 75
    v3 := make(map[string]string)
    var v4 string = "foo bar"
    switch ("hello" + ("a" + v4)) {
    case "":
        var v34 string = v4
        _ = v34
    case "\n\t":
        if (611 > 0) {
            var v35 bool = ((v1 * (3.14159 - ((-2.25 + ((v1 / (0.0 + (v1 - v1))) / v1)) - (v1 - (v1 / (v1 + v1)))))) > -2.25)
            if v35 {
                var v36 string = (v4 + ((v4 + v4) + ((v4 + v4) + v4)))
                _ = v36
            }
            var v37 int64 = 0
            _ = v35
            _ = v37
        }
    }
    if (!false) {
        v38 := []int{2147483647, 128, 32768}
        if true {
            var v39 int = -941
            _ = v39
        } else {
            if false {
                for fi40 := 0; fi40 < 11; fi40++ {
                    v4 = ""
                    _ = fi40
                }
                v41 := GenS3{F0: 150, F1: 42, F2: 1}
                _ = v41
            } else {
                v42 := GenS5{F0: 9223372036854775807, F1: 1, F2: 371}
                _ = v42
            }
        }
        _ = v38
    } else {
        var v43 bool = true
        var v44 uint = 0
        v45 := []int{42, -832, 9223372036854775807, 128}
        _ = v43
        _ = v44
        _ = v45
    }
    if false {
        v4 = (v4 + ((("hello" + "\n\t") + "") + "foo bar"))
        v46 := GenS3{F0: 32768, F1: 0, F2: 706}
        _ = v46
    } else {
        if (true || (127 == 0)) {
            var v47 int32 = 2147483647
            switch -1 {
            case 0:
                switch ("\n\t" + ((v4 + "foo bar") + v4)) {
                case "hello":
                    var v48 uint = 4294967295
                    _ = v48
                case "\n\t":
                    v49 := []int{42}
                    _ = v49
                case "":
                    var v50 float32 = 1.5
                    _ = v50
                }
            case 127:
                v51 := []int{-932}
                _ = v51
            default:
                v4 = (v4 + ("hello" + v4))
            }
            v3[v4] = v4
            _ = v47
        } else {
            v4 = v4
            v1 -= ((v1 * v1) + (((((-2.25 - v1) * v1) + 0.1) - ((0.1 - ((0.0 - (((v1 / (0.1 + float32(-2.25))) - (v1 - v1)) - v1)) + ((((float32(1e10) + -2.25) / 0.1) / v1) + 1.5))) - -2.25)) + ((float32(3.14159) * 3.14159) - v1)))
        }
        switch (v4 + v4) {
        case "foo bar":
            if false {
                v2 = (-(v2))
            } else {
                _ = len(v3)
                v1 = 0.0
                if (2147483647 > 4294967295) {
                    switch 42 {
                    case 32768:
                        var v52 uint16 = 65535
                        _ = v52
                    case 42:
                        var v53 bool = (!false)
                        _ = v53
                    default:
                        v54 := []int{32768, 1}
                        _ = v54
                    }
                } else {
                    switch v4 {
                    case "":
                        v4 = "a"
                    }
                    var v55 int32 = 0
                    _ = v55
                }
            }
        case "\n\t":
            var v56 int16 = 127
            _ = v56
        case "hello":
            if false {
                var v57 float64 = 1.5
                for fi58 := 0; fi58 < 6; fi58++ {
                    fmt.Println(v2, v2)
                    _ = fi58
                }
                for fi59 := 0; fi59 < 1; fi59++ {
                    if false {
                        v57 += (float64(v1) / v57)
                    } else {
                        for fi60 := 0; fi60 < 10; fi60++ {
                            switch (("\n\t" + v4) + "a") {
                            case "hello":
                                if false {
                                    v61 := GenS6{F0: -775, F1: 376, F2: 65535}
                                    var v62 uint32 = 1
                                    _ = v61
                                    _ = v62
                                } else {
                                    switch v4 {
                                    case "a":
                                        var v63 uint64 = 877
                                        _ = v63
                                    case "\n\t":
                                        var v64 float64 = -2.25
                                        _ = v64
                                    case "":
                                        v65 := GenS4{F0: 1, F1: 42}
                                        if (!true) {
                                            v57 = (float64(3.14159) / (v57 - float64(0.1)))
                                            switch "\n\t" {
                                            case "a":
                                                var v66 uint8 = 227
                                                _ = v66
                                            case "":
                                                v67 := GenS6{F0: -1, F1: 764, F2: 2147483647}
                                                _ = v67
                                            default:
                                                var v68 int32 = 1
                                                _ = v68
                                            }
                                        } else {
                                            switch fi60 {
                                            case 897:
                                                v4 = "hello"
                                            default:
                                                v69 := GenS4{F0: 1000, F1: -313}
                                                var v70 int64 = -712
                                                _ = v69
                                                _ = v70
                                            }
                                        }
                                        v71 := GenS7{F0: 1}
                                        v4 = ""
                                        _ = v65
                                        _ = v71
                                    }
                                    if false {
                                        v72 := GenS5{F0: 1, F1: 255, F2: 0}
                                        var v73 int16 = 1
                                        var v74 bool = true
                                        _ = v72
                                        _ = v73
                                        _ = v74
                                    } else {
                                        var v75 int = (fi60 * (fi60 % (((fi60 ^ fi59)) | 1)))
                                        fmt.Println(v2, fi59)
                                        _ = v75
                                    }
                                }
                            case "a":
                                if true {
                                    v76 := GenS5{F0: -59, F1: 0, F2: 883}
                                    v77 := v3[v4]
                                    var v78 int16 = 1000
                                    _ = v76
                                    _ = v77
                                    _ = v78
                                }
                            default:
                                v79 := GenS8{F0: -1, F1: ("hello" > "a"), F2: 65535}
                                _ = v79
                            }
                            _ = fi60
                        }
                        if (!false) {
                            for fi80 := 0; fi80 < 13; fi80++ {
                                fmt.Println(v4)
                                if true {
                                    switch v4 {
                                    case "foo bar":
                                        var v81 int = (fi80 << ((96) & 63))
                                        _ = v81
                                    case "\n\t":
                                        switch v4 {
                                        case "":
                                            var v82 float32 = v1
                                            _ = v82
                                        case "\n\t":
                                            v2 = v2
                                        default:
                                            v1 = v1
                                        }
                                    case "a":
                                        v83 := GenS8{F0: 42, F1: true, F2: 0}
                                        _ = v83
                                    default:
                                        var v84 int8 = v2
                                        _ = v84
                                    }
                                    var v85 string = ""
                                    switch ("" + v4) {
                                    case "foo bar":
                                        var v86 uint16 = 32767
                                        _ = v86
                                    case "\n\t":
                                        v57 = (v57 * v57)
                                        var v87 float64 = 1.5
                                        _ = v87
                                    case "":
                                        if (0 == 1) {
                                            var v88 uint16 = 0
                                            var v89 int16 = 0
                                            var v90 int = fi80
                                            var v91 uint32 = 2147483648
                                            _ = v88
                                            _ = v89
                                            _ = v90
                                            _ = v91
                                        } else {
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v92 string = "hello"
                                                _ = v92
                                            }
                                        }
                                        var v93 uint = 2
                                        var v94 int16 = 42
                                        _ = v93
                                        _ = v94
                                    default:
                                        v95 := []int{fi59}
                                        _ = v95
                                    }
                                    _ = v85
                                } else {
                                    v2 += 127
                                    fmt.Println(v2, v1)
                                    var v96 int = (fi59 ^ 1)
                                    _ = v96
                                }
                                _ = fi80
                            }
                            v2 = v2
                            v57 -= float64(v1)
                        }
                        switch fi59 {
                        case 127:
                            var v97 uint64 = 2
                            _ = v97
                        default:
                            v2 = v2
                        }
                    }
                    fmt.Println(v57, v57)
                    _ = fi59
                }
                _ = v57
            } else {
                if false {
                    switch 127 {
                    case -510:
                        fmt.Println(v4, v4)
                    case 301:
                        fmt.Println(v4, v2)
                    default:
                        for fi98 := 0; fi98 < 10; fi98++ {
                            var v99 uint = 877
                            _ = fi98
                            _ = v99
                        }
                    }
                }
                v4 = "\n\t"
            }
        default:
            for fi100 := 0; fi100 < 2; fi100++ {
                v101 := GenS7{F0: 2}
                var v102 string = "a"
                _ = fi100
                _ = v101
                _ = v102
            }
        }
        v103 := GenS8{F0: 308, F1: (("hello" > v4) || false), F2: 2147483648}
        _ = v103
    }
    if (!(!true)) {
        v104 := GenS5{F0: 1, F1: 65535, F2: 42}
        for fi105 := 0; fi105 < 7; fi105++ {
            for fi106 := 0; fi106 < 20; fi106++ {
                fmt.Println(v2, v4)
                delete(v3, v4)
                var v107 int32 = 2147483647
                v108 := []int{}
                _ = fi106
                _ = v107
                _ = v108
            }
            _ = fi105
        }
        if (true && true) {
            var v109 float32 = float32(1e10)
            for fi110 := 0; fi110 < 4; fi110++ {
                if false {
                    switch fi110 {
                    case 4294967295:
                        if (false && (1 <= 2)) {
                            v111 := v3[(("" + ("foo bar" + "a")) + (v4 + "\n\t"))]
                            _ = v111
                        }
                    case -544:
                        var v112 int16 = 1000
                        _ = v112
                    case -782:
                        switch (^(fi110)) {
                        case -9223372036854775808:
                            if (842 <= 0) {
                                if (32767 <= 245) {
                                    v113 := GenS2{F0: (v2 & (-(v2))), F1: 18446744073709551615}
                                    v1 += float32(((1e10 / 1.0) * -2.25))
                                    _ = v113
                                } else {
                                    if (0.1 <= v109) {
                                        for fi114 := 0; fi114 < 1; fi114++ {
                                            var v115 float32 = (v109 + 1.5)
                                            var v116 int32 = -2147483648
                                            _ = fi114
                                            _ = v115
                                            _ = v116
                                        }
                                        for fi117 := 0; fi117 < 10; fi117++ {
                                            fmt.Println(v4)
                                            _ = fi117
                                        }
                                        if false {
                                            switch -342 {
                                            case 32767:
                                                fmt.Println(fi110, v2)
                                            default:
                                                var v118 int = (fi110 | 4294967295)
                                                _ = v118
                                            }
                                            var v119 uint64 = 923
                                            fmt.Println(v109, v1)
                                            _ = v119
                                        } else {
                                            var v120 string = v4
                                            v4 = (v120 + v4)
                                            _ = v120
                                        }
                                    }
                                }
                                if false {
                                    if (!(!(v2 <= (v2 | (v2 * v2))))) {
                                        var v121 uint16 = 1
                                        for fi122 := 0; fi122 < 9; fi122++ {
                                            var v123 int8 = (v2 - (v2 & (v2 & v2)))
                                            _ = fi122
                                            _ = v123
                                        }
                                        switch ("" + v4) {
                                        case "":
                                            v2 = 100
                                        default:
                                            if true {
                                                v124 := GenS1{F0: "", F1: 1, F2: 70}
                                                v3["hello"] = v4
                                                if true {
                                                    for fi125 := 0; fi125 < 20; fi125++ {
                                                        fmt.Println(v2)
                                                        v126 := GenS8{F0: 1, F1: true, F2: 278}
                                                        _ = fi125
                                                        _ = v126
                                                    }
                                                    v127 := v3[(("\n\t" + v4) + v4)]
                                                    v128 := v3[v4]
                                                    _ = v127
                                                    _ = v128
                                                } else {
                                                    var v129 int8 = 59
                                                    if false {
                                                        var v130 float32 = (3.14159 / v1)
                                                        _ = v130
                                                    } else {
                                                        var v131 uint32 = 4294967295
                                                        fmt.Println(v129)
                                                        fmt.Println(v1, fi110, v109)
                                                        _ = v131
                                                    }
                                                    _ = v129
                                                }
                                                _ = v124
                                            } else {
                                                v132 := GenS8{F0: -1, F1: true, F2: 685}
                                                _ = v132
                                            }
                                        }
                                        _ = v121
                                    }
                                    for fi133 := 0; fi133 < 7; fi133++ {
                                        var v134 int = 9223372036854775807
                                        var v135 bool = (621 != 1)
                                        if (((fi133 / ((fi133) | 1)) > v134) || (((!v135) && ((false && ((^(v2)) > v2)) || (!false))) && (!(2 == 255)))) {
                                            v2 = (v2 / (((v2 * (v2 >> ((0) & 63)))) | 1))
                                            v3[""] = (("a" + (v4 + v4)) + v4)
                                            switch (v4 + v4) {
                                            case "":
                                                var v136 int16 = -1
                                                _ = v136
                                            default:
                                                v1 = (1.5 / 1.0)
                                            }
                                            var v137 float64 = 0.0
                                            v137 = (1.5 - float64(1e10))
                                            _ = v137
                                        } else {
                                            var v138 int8 = v2
                                            _ = v138
                                        }
                                        v139 := GenS9{F0: -2147483648}
                                        fmt.Println(v134, fi133)
                                        _ = fi133
                                        _ = v134
                                        _ = v135
                                        _ = v139
                                    }
                                    fmt.Println(v1)
                                } else {
                                    v109 += ((v1 * v109) - (3.14159 + 1e10))
                                    if false {
                                        if false {
                                            v140 := []int{fi110, fi110, (^(fi110)), (fi110 - 823)}
                                            fmt.Println(v109, v4)
                                            _ = v140
                                        }
                                        v141 := GenS3{F0: -1, F1: 9223372036854775807, F2: -2147483648}
                                        var v142 uint64 = 533
                                        _ = v141
                                        _ = v142
                                    } else {
                                        v143 := []int{fi110}
                                        switch v4 {
                                        case "\n\t":
                                            delete(v3, "foo bar")
                                        case "":
                                            var v144 float64 = 3.14159
                                            _ = v144
                                        default:
                                            v145 := GenS7{F0: 103}
                                            _ = v145
                                        }
                                        v146 := GenS7{F0: 255}
                                        _ = v143
                                        _ = v146
                                    }
                                }
                            }
                        case 156:
                            _ = len(v3)
                        case -1:
                            switch ("foo bar" + v4) {
                            case "hello":
                                v109 -= (v1 * (1.5 + float32((-2.25 / 1.5))))
                            }
                        }
                    }
                } else {
                    for fi147 := 0; fi147 < 7; fi147++ {
                        var v148 int64 = -1
                        if false {
                            if true {
                                switch fi110 {
                                case 42:
                                    v149 := []int{128, (fi147 - 42), -1}
                                    _ = v149
                                default:
                                    if (((1.5 + (((v1 - float32((-2.25 + (3.14159 / 1.5)))) * (v1 - 1e10)) + 0.0)) + 3.14159) > v109) {
                                        v148 += (v148 + (v148 / (((v148 / ((42) | 1))) | 1)))
                                        if (!false) {
                                            if ((fi147 * 42) < fi147) {
                                                v150 := GenS6{F0: v148, F1: 2147483647, F2: 0}
                                                var v151 int16 = 1
                                                _ = v150
                                                _ = v151
                                            }
                                            switch "" {
                                            case "hello":
                                                if (!true) {
                                                    switch 127 {
                                                    case 4294967295:
                                                        v152 := GenS5{F0: (v148 >> ((255) & 63)), F1: 255, F2: 18446744073709551615}
                                                        _ = v152
                                                    case 287:
                                                        fmt.Println(v4, v2, v4)
                                                        v2 = v2
                                                        v2 = v2
                                                        var v153 int = fi110
                                                        _ = v153
                                                    }
                                                    v154 := GenS10{F0: -9}
                                                    _ = v154
                                                }
                                            case "foo bar":
                                                fmt.Println(fi147, v4)
                                                _ = len(v3)
                                            }
                                            v148 = (v148 << ((310) & 63))
                                            v155 := []int{fi110}
                                            if false {
                                                v156 := GenS11{F0: 127, F1: 98, F2: 4294967295}
                                                var v157 int16 = 42
                                                switch (fi110 | (fi147 << ((1) & 63))) {
                                                case 128:
                                                    if (!(-771 <= 0)) {
                                                        v158 := v3[v4]
                                                        switch -963 {
                                                        case 802:
                                                            v159 := GenS3{F0: fi147, F1: (-(v148)), F2: -1}
                                                            _ = v159
                                                        case 2147483647:
                                                            v160 := GenS6{F0: v148, F1: 2147483647, F2: 139}
                                                            _ = v160
                                                        }
                                                        _ = v158
                                                    } else {
                                                        v2 = v2
                                                        v161 := GenS4{F0: 296, F1: -2147483648}
                                                        v1 = (3.14159 + v109)
                                                        if (42 < 1) {
                                                            fmt.Println(v109, v4, fi147)
                                                        }
                                                        if (147 <= 184) {
                                                            if len(v155) >= 2 {
                                                                v155 = v155[1:len(v155)]
                                                            }
                                                            var v162 float64 = 0.0
                                                            v163 := GenS12{F0: 1, F1: false, F2: v148}
                                                            _ = v162
                                                            _ = v163
                                                        }
                                                        if len(v155) >= 2 {
                                                            v155 = v155[1:len(v155)]
                                                        }
                                                        _ = v161
                                                    }
                                                    switch v4 {
                                                    case "foo bar":
                                                        var v164 uint32 = 2
                                                        _ = v164
                                                    default:
                                                        fmt.Println(v1, v157, v157)
                                                    }
                                                case 9223372036854775807:
                                                    var v165 int32 = 2147483647
                                                    _ = v165
                                                default:
                                                    _ = len(v3)
                                                }
                                                _ = v156
                                                _ = v157
                                            }
                                            var v166 int16 = -1
                                            _ = v155
                                            _ = v166
                                        }
                                        v167 := GenS1{F0: (v4 + v4), F1: 1, F2: 127}
                                        _ = v167
                                    }
                                }
                                fmt.Println(v1)
                            }
                        }
                        _ = fi147
                        _ = v148
                    }
                }
                v168 := GenS5{F0: -9223372036854775808, F1: 42, F2: 4294967295}
                var v169 float32 = v109
                _ = fi110
                _ = v168
                _ = v169
            }
            var v170 float32 = v109
            _ = v109
            _ = v170
        }
        _ = v104
    }
    v171 := []int{}
    v172 := v3["a"]
    for fi173 := 0; fi173 < 3; fi173++ {
        for fi174 := 0; fi174 < 15; fi174++ {
            fmt.Println(v1, v4)
            var v175 uint = 2
            var v176 int8 = -53
            _ = fi174
            _ = v175
            _ = v176
        }
        v171 = append(v171, fi173)
        _ = fi173
    }
    var v177 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v171:", len(v171), v171)
    fmt.Println("v172:", v172)
    fmt.Println("v177:", v177)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v171
    _ = v172
    _ = v177
    _ = v2
    _ = v3
    _ = v4
}
