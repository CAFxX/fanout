package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 string; F1 uint32; F2 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint; F1 uint16 }
type GenS8 struct { F0 int32; F1 int16; F2 uint32 }
type GenS9 struct { F0 uint16; F1 string; F2 int64 }
type GenS10 struct { F0 uint8; F1 uint16 }
type GenS11 struct { F0 bool; F1 bool; F2 uint32 }
type GenS12 struct { F0 bool; F1 int; F2 uint32 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int32; F1 uint64 }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint16; F1 int64 }
func genFn0(p0 int16) uint64 {
    v4 := GenS1{F0: 255, F1: 911, F2: 42}
    p0 = (p0 >> ((988) & 63))
    v5 := []int{9223372036854775807, -2147483648}
    var v6 int32 = -1
    if len(v5) > 0 {
        v5[0] = 1
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 28
}

func main() {
    v1 := []int{}
    var v2 float32 = 1e10
    v3 := make(map[uint32]uint32)
    switch "" {
    case "foo bar":
        for fi7 := 0; fi7 < 16; fi7++ {
            v8 := GenS1{F0: 2, F1: 2, F2: 255}
            _ = fi7
            _ = v8
        }
    default:
        if (false || (true && (0 <= 0))) {
            if true {
                v1 = append(v1, 128)
                var v9 uint = 2
                if false {
                    fmt.Println(v2, v2)
                    var v10 uint32 = 239
                    var v11 int = 1
                    _ = v10
                    _ = v11
                } else {
                    v12 := v3[65535]
                    _ = v12
                }
                _ = v9
            }
            var v13 uint8 = 217
            v2 -= ((-2.25 - (((float32((((1e10 / 0.1) * (float64(v2) * (3.14159 + 1.5))) + -2.25)) * v2) / v2) - (v2 + ((v2 + float32(1e10)) + 1e10)))) / 3.14159)
            _ = v13
        }
    }
    v14 := GenS2{F0: 992}
    v2 = v2
    fmt.Println(v2)
    v2 = v2
    if (-1 > -9223372036854775808) {
        v2 += v2
    }
    var v15 string = "a"
    var v16 int8 = 100
    var v17 int = 2147483647
    if ((0 != 27) || (-2.25 <= -2.25)) {
        v18 := GenS2{F0: 0}
        var v19 int16 = 1000
        if ((!true) && true) {
            if true {
                var v20 uint = 18446744073709551615
                if true {
                    if (!false) {
                        var v21 int16 = 32767
                        v21 += v21
                        _ = v21
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v17 -= (-(v17))
                    }
                    if false {
                        v19 -= (v19 & v19)
                        var v22 int = v17
                        switch (("a" + "") + (v15 + "\n\t")) {
                        case "":
                            var v23 int32 = -2147483648
                            _ = v23
                        case "a":
                            var v24 int = (v22 << ((350) & 63))
                            _ = v24
                        case "foo bar":
                            v25 := GenS3{F0: 42, F1: 2147483648}
                            _ = v25
                        default:
                            v26 := GenS3{F0: 42, F1: 1}
                            _ = v26
                        }
                        _ = v22
                    }
                    v3[42] = 4294967295
                }
                if ((v20 & v20) != v20) {
                    var v27 int16 = v19
                    _ = v27
                } else {
                    if ((1 >= 65535) || (v19 > (v19 ^ v19))) {
                        for fi28 := 0; fi28 < 2; fi28++ {
                            if (false || false) {
                                v3[159] = 2147483648
                                v15 = ("hello" + "\n\t")
                                var v29 int8 = 61
                                _ = v29
                            }
                            v30 := GenS4{F0: ("hello" + "foo bar"), F1: 2147483648, F2: -1}
                            if (true && false) {
                                fmt.Println(v16)
                            } else {
                                switch v15 {
                                case "hello":
                                    v1 = append(v1, fi28)
                                default:
                                    switch 1 {
                                    case 134:
                                        v20 += 4294967295
                                    }
                                }
                                for fi31 := 0; fi31 < 17; fi31++ {
                                    delete(v3, 65535)
                                    _ = fi31
                                }
                            }
                            _ = fi28
                            _ = v30
                        }
                    } else {
                        if true {
                            for fi32 := 0; fi32 < 18; fi32++ {
                                var v33 uint64 = 2
                                v34 := GenS4{F0: v15, F1: 374, F2: 1}
                                _ = fi32
                                _ = v33
                                _ = v34
                            }
                        } else {
                            var v35 uint = v20
                            _ = v35
                        }
                        v3[uint32(v20)] = 2
                    }
                    switch "a" {
                    case "a":
                        fmt.Println(v15)
                    }
                    _ = len(v3)
                }
                _ = v20
            }
            if (42 >= 235) {
                for fi36 := 0; fi36 < 7; fi36++ {
                    fmt.Println(v19)
                    var v37 string = v15
                    var v38 int32 = -930
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                v39 := GenS5{F0: 18446744073709551615}
                _ = v39
            }
            fmt.Println(v17, v16, v15)
        } else {
            v2 -= 1.5
            v19 = (v19 + (v19 % (((^(v19))) | 1)))
        }
        _ = v18
        _ = v19
    } else {
        switch v15 {
        case "foo bar":
            _ = len(v3)
        case "":
            _ = len(v3)
        case "\n\t":
            var v40 int64 = -179
            _ = v40
        default:
            for fi41 := 0; fi41 < 20; fi41++ {
                if true {
                    v2 = v2
                }
                switch ("\n\t" + (v15 + v15)) {
                case "foo bar":
                    for fi42 := 0; fi42 < 11; fi42++ {
                        var v43 float32 = (v2 / 1e10)
                        v17 = v17
                        _ = fi42
                        _ = v43
                    }
                }
                _ = fi41
            }
        }
    }
    var v44 uint8 = 0
    if true {
        if true {
            var v45 bool = true
            for fi46 := 0; fi46 < 19; fi46++ {
                if (448 == 65535) {
                    for fi47 := 0; fi47 < 19; fi47++ {
                        switch 157 {
                        case -2147483648:
                            v48 := GenS6{F0: 92}
                            _ = v48
                        case 128:
                            if (4294967295 >= 31) {
                                var v49 int8 = v16
                                _ = v49
                            }
                        }
                        _ = fi47
                    }
                }
                v50 := GenS5{F0: 1}
                _ = fi46
                _ = v50
            }
            switch v15 {
            case "foo bar":
                for fi51 := 0; fi51 < 11; fi51++ {
                    var v52 bool = (!true)
                    switch v15 {
                    case "\n\t":
                        switch v15 {
                        case "\n\t":
                            if len(v1) > 0 {
                                v1[0] = v17
                            }
                        case "hello":
                            if ((!v45) || ((((v15 + v15) + (v15 + (v15 + v15))) + v15) >= "")) {
                                if (203 < uint32(v44)) {
                                    for fi53 := 0; fi53 < 8; fi53++ {
                                        var v54 float64 = 0.0
                                        switch (v17 * (v17 << ((18446744073709551615) & 63))) {
                                        case 32768:
                                            switch "\n\t" {
                                            case "a":
                                                v55 := GenS3{F0: 18446744073709551615, F1: 0}
                                                var v56 uint = 65535
                                                v57 := GenS6{F0: v16}
                                                _ = v55
                                                _ = v56
                                                _ = v57
                                            case "\n\t":
                                                v52 = false
                                            default:
                                                if (!false) {
                                                    _ = len(v3)
                                                } else {
                                                    if (v2 != v2) {
                                                        switch v15 {
                                                        case "\n\t":
                                                            switch "\n\t" {
                                                            case "":
                                                                _ = len(v3)
                                                            }
                                                        case "foo bar":
                                                            var v58 float32 = v2
                                                            _ = v58
                                                        case "hello":
                                                            v2 = v2
                                                        }
                                                        fmt.Println(v54, v15, v45)
                                                    } else {
                                                        if (v2 == -2.25) {
                                                            v15 = (v15 + "\n\t")
                                                            var v59 uint16 = 901
                                                            _ = v59
                                                        } else {
                                                            v15 = v15
                                                            v45 = v45
                                                            v60 := GenS5{F0: 0}
                                                            _ = v60
                                                        }
                                                        if v45 {
                                                            fmt.Println(v44, v54, v17)
                                                            var v61 float32 = (v2 - v2)
                                                            if (42 <= 1) {
                                                                v61 += (v2 - (((((float32(0.1) / v2) * (v61 - float32(float64(float32(float64(1e10)))))) / v2) * v2) - v2))
                                                                v2 -= (float32((v54 / 3.14159)) / (v2 - v61))
                                                                v62 := GenS1{F0: 299, F1: 1, F2: 65535}
                                                                _ = v62
                                                            }
                                                            _ = v61
                                                        }
                                                        v63 := GenS7{F0: uint(v44), F1: 255}
                                                        v2 = 0.0
                                                        v52 = (!true)
                                                        fmt.Println(v17)
                                                        v64 := GenS6{F0: v16}
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                }
                                            }
                                        case 0:
                                            if ((1.5 < (-2.25 * v2)) || (!((v15 > "hello") && v45))) {
                                                v65 := GenS3{F0: 971, F1: 1}
                                                _ = v65
                                            }
                                        case -968:
                                            var v66 int8 = int8(fi53)
                                            _ = v66
                                        default:
                                            _ = len(v3)
                                        }
                                        var v67 float64 = 0.1
                                        _ = fi53
                                        _ = v54
                                        _ = v67
                                    }
                                    var v68 uint = 255
                                    _ = v68
                                }
                            } else {
                                for fi69 := 0; fi69 < 18; fi69++ {
                                    v70 := GenS8{F0: 986, F1: 1000, F2: 4294967295}
                                    _ = fi69
                                    _ = v70
                                }
                                switch (v15 + "foo bar") {
                                case "foo bar":
                                    switch fi51 {
                                    case 51:
                                        var v71 uint64 = genFn0(-32768)
                                        _ = v71
                                    case 0:
                                        v2 += ((v2 + (v2 * (v2 / v2))) + v2)
                                    case 134:
                                        v72 := GenS4{F0: "hello", F1: 42, F2: (^(fi51))}
                                        _ = v72
                                    default:
                                        v17 = (v17 << ((530) & 63))
                                    }
                                case "a":
                                    for fi73 := 0; fi73 < 7; fi73++ {
                                        v74 := GenS5{F0: 403}
                                        _ = fi73
                                        _ = v74
                                    }
                                }
                            }
                        default:
                            for fi75 := 0; fi75 < 16; fi75++ {
                                var v76 uint64 = 0
                                var v77 uint32 = 2147483648
                                fmt.Println(v17)
                                _ = fi75
                                _ = v76
                                _ = v77
                            }
                        }
                    case "hello":
                        if ((!v45) && true) {
                            if ((!(uint16(v44) == 709)) || v52) {
                                for fi78 := 0; fi78 < 15; fi78++ {
                                    fmt.Println(v52, v45, v2)
                                    var v79 bool = (true && (v45 || v45))
                                    _ = fi78
                                    _ = v79
                                }
                            }
                            var v80 uint32 = 4294967295
                            switch "\n\t" {
                            case "foo bar":
                                _ = len(v1)
                            case "":
                                v52 = (v15 == v15)
                            }
                            _ = v80
                        }
                    default:
                        v81 := v3[0]
                        _ = v81
                    }
                    _ = fi51
                    _ = v52
                }
            }
            _ = v45
        } else {
            for fi82 := 0; fi82 < 3; fi82++ {
                _ = len(v3)
                _ = fi82
            }
        }
        _ = len(v1)
        if false {
            if ((!(float64(v2) >= ((float64((-2.25 - v2)) / 1.0) * (3.14159 - 3.14159)))) || (190 >= 489)) {
                var v83 uint16 = 255
                for fi84 := 0; fi84 < 12; fi84++ {
                    v2 -= ((((1.5 + ((0.0 / (1.5 * v2)) / v2)) + (v2 * (v2 - v2))) / ((((1e10 / (float32(1.5) - 3.14159)) / float32((-2.25 / (float64(v2) - (float64(3.14159) * (0.1 * -2.25)))))) / 3.14159) / (v2 - v2))) / float32(3.14159))
                    _ = fi84
                }
                _ = v83
            } else {
                switch (v15 + v15) {
                case "":
                    if ((-992 <= 9223372036854775807) && (v44 <= 2)) {
                        for fi85 := 0; fi85 < 7; fi85++ {
                            if true {
                                switch v17 {
                                case 0:
                                    v86 := GenS7{F0: 255, F1: 550}
                                    _ = v86
                                case 42:
                                    fmt.Println(v2)
                                default:
                                    v15 = v15
                                }
                            } else {
                                if (true && false) {
                                    var v87 float64 = (((-2.25 + 1e10) * -2.25) / 1.0)
                                    var v88 int32 = -1
                                    v89 := GenS9{F0: 958, F1: "", F2: 1}
                                    _ = v87
                                    _ = v88
                                    _ = v89
                                } else {
                                    v90 := GenS7{F0: 209, F1: 570}
                                    if false {
                                        for fi91 := 0; fi91 < 14; fi91++ {
                                            fmt.Println(v2, v16, fi85)
                                            if len(v1) > 0 {
                                                v1[0] = 32768
                                            }
                                            var v92 uint16 = 32767
                                            _ = fi91
                                            _ = v92
                                        }
                                        for fi93 := 0; fi93 < 2; fi93++ {
                                            v94 := GenS10{F0: (^(v44)), F1: 215}
                                            switch (v15 + (v15 + (v15 + v15))) {
                                            case "a":
                                                v2 -= (float32(float64((v2 / (1.5 * (0.0 + v2))))) + 0.0)
                                            case "":
                                                if ((^(v16)) > 1) {
                                                    var v95 uint64 = 18446744073709551615
                                                    _ = v95
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = (fi93 | v17)
                                                }
                                                var v96 uint64 = 1
                                                v15 = (v15 + v15)
                                                _ = v96
                                            case "\n\t":
                                                fmt.Println(fi85)
                                                switch "" {
                                                case "a":
                                                    switch v17 {
                                                    case 719:
                                                        v97 := GenS11{F0: (!false), F1: true, F2: 42}
                                                        _ = v97
                                                    case -2147483648:
                                                        v98 := GenS12{F0: false, F1: fi85, F2: 2147483648}
                                                        _ = v98
                                                    case 4294967295:
                                                        v99 := v3[2]
                                                        _ = v99
                                                    default:
                                                        v100 := GenS2{F0: 65535}
                                                        _ = v100
                                                    }
                                                default:
                                                    var v101 int32 = -2147483648
                                                    _ = v101
                                                }
                                                fmt.Println(v15, v2, v17)
                                            }
                                            _ = fi93
                                            _ = v94
                                        }
                                    }
                                    v102 := GenS11{F0: (true && true), F1: ((v2 + 3.14159) > v2), F2: 2147483648}
                                    _ = v102
                                    _ = v90
                                }
                                switch v15 {
                                case "hello":
                                    for fi103 := 0; fi103 < 9; fi103++ {
                                        var v104 string = v15
                                        var v105 bool = false
                                        fmt.Println(v44, fi85)
                                        _ = fi103
                                        _ = v104
                                        _ = v105
                                    }
                                case "foo bar":
                                    var v106 int16 = 1000
                                    _ = v106
                                case "":
                                    v15 = v15
                                default:
                                    fmt.Println(v2, v16)
                                }
                            }
                            _ = fi85
                        }
                        for fi107 := 0; fi107 < 20; fi107++ {
                            switch (fi107 / (((v17 << ((750) & 63))) | 1)) {
                            case -1:
                                switch ((v15 + v15) + v15) {
                                case "":
                                    var v108 float64 = (((0.0 / 1.0) / -2.25) * -2.25)
                                    _ = v108
                                case "a":
                                    v109 := v3[2147483647]
                                    _ = v109
                                default:
                                    var v110 uint = 0
                                    _ = v110
                                }
                            }
                            v111 := GenS9{F0: 2, F1: (v15 + "hello"), F2: 0}
                            var v112 int16 = 1000
                            _ = fi107
                            _ = v111
                            _ = v112
                        }
                    }
                case "foo bar":
                    var v113 float32 = 0.0
                    _ = v113
                }
                var v114 uint64 = 4294967295
                if false {
                    v115 := GenS13{F0: int64(v17)}
                    _ = v115
                } else {
                    for fi116 := 0; fi116 < 18; fi116++ {
                        switch (fi116 ^ v17) {
                        case -1:
                            switch int(v44) {
                            case -2147483648:
                                for fi117 := 0; fi117 < 20; fi117++ {
                                    v1 = append(v1, 9223372036854775807)
                                    var v118 uint64 = (v114 ^ (^(v114)))
                                    _ = fi117
                                    _ = v118
                                }
                            case 2147483647:
                                for fi119 := 0; fi119 < 9; fi119++ {
                                    v2 = (v2 * (float32(3.14159) * float32(1e10)))
                                    var v120 uint16 = 65535
                                    _ = fi119
                                    _ = v120
                                }
                            default:
                                var v121 int32 = 2147483647
                                _ = v121
                            }
                        case -256:
                            var v122 float64 = (float64(v2) / 1e10)
                            _ = v122
                        case -2147483648:
                            var v123 bool = false
                            _ = v123
                        }
                        v124 := GenS14{F0: 1, F1: v114}
                        var v125 bool = false
                        _ = fi116
                        _ = v124
                        _ = v125
                    }
                    for fi126 := 0; fi126 < 17; fi126++ {
                        v127 := GenS7{F0: 2, F1: 42}
                        _ = len(v1)
                        switch fi126 {
                        case 32767:
                            switch "a" {
                            case "":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                fmt.Println(fi126, v16)
                            }
                        case 981:
                            fmt.Println(v114, fi126, fi126)
                        }
                        _ = fi126
                        _ = v127
                    }
                    fmt.Println(v15)
                }
                _ = v114
            }
            v3[2147483648] = 991
        } else {
            if (-774 != 2147483647) {
                fmt.Println(v44)
                if (1 != uint32(v44)) {
                    var v128 int8 = v16
                    var v129 uint64 = genFn0(-1)
                    _ = v128
                    _ = v129
                } else {
                    if false {
                        if (false || ((1.5 + -2.25) >= 1e10)) {
                            _ = len(v1)
                            var v130 float32 = v2
                            var v131 int64 = -1
                            _ = v130
                            _ = v131
                        }
                        v2 -= 0.0
                    }
                    for fi132 := 0; fi132 < 5; fi132++ {
                        v2 = (1.5 / (v2 - (3.14159 - 1.5)))
                        v133 := GenS11{F0: true, F1: (106 <= 0), F2: 0}
                        _ = fi132
                        _ = v133
                    }
                }
                v134 := GenS2{F0: 1}
                _ = v134
            } else {
                var v135 string = v15
                var v136 uint32 = 42
                _ = v135
                _ = v136
            }
        }
    } else {
        var v137 uint8 = (^(v44))
        _ = v137
    }
    fmt.Println(v2, v15, v44)
    switch v15 {
    case "a":
        v3[176] = 2147483648
    default:
        var v138 uint8 = (v44 / ((0) | 1))
        _ = v138
    }
    for fi139 := 0; fi139 < 6; fi139++ {
        switch ("\n\t" + v15) {
        case "\n\t":
            fmt.Println(v15, v44)
        case "foo bar":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            v140 := GenS15{F0: -852}
            _ = v140
        }
        v16 += (v16 - 100)
        _ = fi139
    }
    var v141 uint = 2
    for fi142 := 0; fi142 < 12; fi142++ {
        for fi143 := 0; fi143 < 1; fi143++ {
            _ = len(v3)
            var v144 int = fi142
            v145 := GenS5{F0: v141}
            _ = fi143
            _ = v144
            _ = v145
        }
        switch 32767 {
        case 1:
            v141 -= (v141 * (v141 | (^(v141))))
        case 2147483647:
            switch v15 {
            case "\n\t":
                if ((!true) || (!("a" >= v15))) {
                    switch fi142 {
                    case -663:
                        var v146 uint8 = v44
                        _ = v146
                    case 42:
                        _ = len(v1)
                    case 128:
                        var v147 uint64 = 4294967295
                        _ = v147
                    default:
                        switch "foo bar" {
                        case "\n\t":
                            if false {
                                fmt.Println(fi142)
                            } else {
                                var v148 uint64 = 2
                                fmt.Println(v2, v2, v44)
                                _ = v148
                            }
                        }
                    }
                    switch v17 {
                    case 42:
                        var v149 int = (fi142 * -9223372036854775808)
                        _ = v149
                    case 32767:
                        v141 = v141
                    case 2147483647:
                        v44 -= 1
                    }
                    var v150 int8 = 42
                    _ = v150
                } else {
                    fmt.Println(v15, v2)
                }
            case "a":
                if (!false) {
                    if (32767 <= 42) {
                        for fi151 := 0; fi151 < 10; fi151++ {
                            v17 = v17
                            v152 := GenS16{F0: (^(v17))}
                            v153 := GenS14{F0: -2147483648, F1: 18446744073709551615}
                            var v154 float32 = v2
                            var v155 uint32 = 2
                            _ = fi151
                            _ = v152
                            _ = v153
                            _ = v154
                            _ = v155
                        }
                    } else {
                        var v156 int8 = v16
                        _ = v156
                    }
                    v15 = "\n\t"
                } else {
                    for fi157 := 0; fi157 < 14; fi157++ {
                        var v158 int8 = v16
                        switch v15 {
                        case "":
                            var v159 int16 = 1000
                            _ = v159
                        case "foo bar":
                            v160 := GenS8{F0: int32(v44), F1: -975, F2: 1}
                            _ = v160
                        case "\n\t":
                            var v161 int8 = (v16 << (((v141 ^ (v141 / (((^(v141))) | 1)))) & 63))
                            _ = v161
                        default:
                            switch v15 {
                            case "\n\t":
                                fmt.Println(fi157, fi142, v17)
                            default:
                                switch fi142 {
                                case 42:
                                    if len(v1) > 0 {
                                        v1[0] = (fi157 / ((fi157) | 1))
                                    }
                                case 128:
                                    var v162 int16 = 42
                                    _ = v162
                                case 4294967295:
                                    if (1 != -943) {
                                        var v163 uint8 = v44
                                        _ = v163
                                    }
                                default:
                                    switch fi157 {
                                    case -319:
                                        v2 += (v2 * v2)
                                    case 4294967295:
                                        switch "\n\t" {
                                        case "":
                                            v2 = 3.14159
                                        case "\n\t":
                                            if true {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v164 uint = v141
                                                v158 = 0
                                                _ = v164
                                            } else {
                                                var v165 int32 = 42
                                                _ = v165
                                            }
                                        case "hello":
                                            v166 := GenS3{F0: 0, F1: 0}
                                            _ = v166
                                        }
                                    case -483:
                                        var v167 uint = 65535
                                        v15 = (v15 + v15)
                                        _ = v167
                                    default:
                                        v168 := GenS17{F0: 65535, F1: -1}
                                        _ = v168
                                    }
                                }
                                if false {
                                    v169 := v3[2]
                                    v44 = 0
                                    fmt.Println(v141)
                                    _ = v169
                                }
                            }
                        }
                        _ = fi157
                        _ = v158
                    }
                    v3[0] = 42
                }
            case "":
                var v170 uint = 2
                _ = v170
            default:
                for fi171 := 0; fi171 < 16; fi171++ {
                    v1 = append(v1, (fi171 - 32767))
                    delete(v3, 644)
                    switch (((v15 + ((((v15 + (v15 + v15)) + v15) + v15) + v15)) + v15) + (("a" + "hello") + "\n\t")) {
                    case "":
                        v15 = v15
                    default:
                        var v172 float32 = v2
                        _ = v172
                    }
                    _ = fi171
                }
            }
        }
        var v173 uint8 = v44
        _ = fi142
        _ = v173
    }
    v16 = v16
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v141:", v141)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    _ = v1
    _ = v14
    _ = v141
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v44
}
