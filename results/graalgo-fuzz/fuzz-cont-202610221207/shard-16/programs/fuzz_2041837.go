package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 uint64; F2 uint }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 bool; F1 uint32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint; F1 uint16 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 int; F1 int64; F2 int64 }
type GenS16 struct { F0 uint8; F1 uint; F2 uint }
type GenS17 struct { F0 int; F1 uint32; F2 int64 }
type GenS18 struct { F0 int32; F1 int8 }
type GenS19 struct { F0 uint64; F1 int32 }
type GenS20 struct { F0 uint; F1 bool }
func genFn0() uint32 {
    var v5 int8 = 127
    var v6 bool = (-1 > -887)
    switch -1 {
    case 0:
        var v7 float32 = (1.5 + (3.14159 / 1e10))
        _ = v7
    case -1:
        v8 := []int{-331}
        _ = v8
    default:
        switch "hello" {
        case "":
            fmt.Println(v5, v5)
        }
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        for fi10 := 0; fi10 < 6; fi10++ {
            v11 := []int{(fi9 * fi10), fi9, 298}
            var v12 uint8 = 173
            if (0.0 != 1.5) {
                var v13 uint = 18446744073709551615
                v5 += -1
                if v6 {
                    _ = len(v11)
                    if ((v12 ^ 128) != (v12 - 42)) {
                        var v14 int64 = -9223372036854775808
                        _ = v14
                    } else {
                        v6 = true
                    }
                    var v15 bool = (!(v6 && (!(!v6))))
                    _ = v15
                } else {
                    var v16 int64 = 9223372036854775807
                    v11 = append(v11, int(v5))
                    _ = v16
                }
                _ = v13
            } else {
                var v17 float64 = -2.25
                _ = v17
            }
            v18 := GenS1{F0: 1, F1: 0}
            var v19 string = "hello"
            _ = fi10
            _ = v11
            _ = v12
            _ = v18
            _ = v19
        }
        v20 := []int{fi9, (fi9 >> ((647) & 63)), fi9, fi9}
        _ = fi9
        _ = v20
    }
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 int8, p1 string) uint16 {
    v21 := []int{}
    var v22 int16 = 256
    v22 += v22
    var v23 int64 = -9223372036854775808
    var v24 uint = 2
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    return 175
}

func main() {
    v1 := []int{42, -9223372036854775808, 0, 9223372036854775807}
    var v2 float64 = 1.5
    var v3 int = -9223372036854775808
    var v4 uint16 = 1
    for fi25 := 0; fi25 < 15; fi25++ {
        if true {
            for fi26 := 0; fi26 < 13; fi26++ {
                var v27 int64 = 9223372036854775807
                v27 = v27
                if len(v1) > 0 {
                    v1[0] = (fi25 << ((1) & 63))
                }
                _ = fi26
                _ = v27
            }
            fmt.Println(v3)
        } else {
            fmt.Println(v3)
            v3 -= v3
        }
        _ = fi25
    }
    switch ("a" + ("\n\t" + "\n\t")) {
    case "foo bar":
        switch (^(v3)) {
        case 665:
            v28 := GenS2{F0: 42}
            _ = v28
        case 0:
            for fi29 := 0; fi29 < 17; fi29++ {
                v30 := GenS1{F0: 1, F1: -974}
                v2 -= v2
                _ = fi29
                _ = v30
            }
        default:
            for fi31 := 0; fi31 < 17; fi31++ {
                switch fi31 {
                case 9223372036854775807:
                    v1 = append(v1, fi31)
                case 883:
                    var v32 uint16 = 32767
                    _ = v32
                case 6:
                    for fi33 := 0; fi33 < 1; fi33++ {
                        if true {
                            var v34 uint = 255
                            v35 := GenS3{F0: (4294967295 != 42)}
                            _ = len(v1)
                            var v36 float32 = (-2.25 * ((1e10 * 1.5) / (3.14159 * -2.25)))
                            _ = v34
                            _ = v35
                            _ = v36
                        } else {
                            if (216 == 2) {
                                var v37 uint64 = 2
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v37
                            }
                        }
                        switch (fi33 & (v3 - v3)) {
                        case 128:
                            v2 = v2
                        case -733:
                            v38 := GenS1{F0: 255, F1: -569}
                            _ = v38
                        case 899:
                            v39 := GenS3{F0: ((v4 & genFn1(-1, "foo bar")) == v4)}
                            _ = v39
                        default:
                            v40 := GenS2{F0: 42}
                            _ = v40
                        }
                        _ = fi33
                    }
                }
                _ = fi31
            }
        }
    case "":
        for fi41 := 0; fi41 < 3; fi41++ {
            v3 += fi41
            v2 -= v2
            _ = fi41
        }
    case "a":
        v4 += v4
    default:
        _ = len(v1)
    }
    if true {
        if (0.1 >= (1e10 + 1e10)) {
            v3 += (v3 % ((v3) | 1))
            if (32767 <= 32767) {
                fmt.Println(v4, v4)
                v4 -= 255
                switch "\n\t" {
                case "a":
                    for fi42 := 0; fi42 < 16; fi42++ {
                        if false {
                            switch "foo bar" {
                            case "":
                                v3 += (-(fi42))
                            case "foo bar":
                                for fi43 := 0; fi43 < 3; fi43++ {
                                    var v44 int64 = 9223372036854775807
                                    _ = fi43
                                    _ = v44
                                }
                            case "a":
                                var v45 float32 = 1e10
                                _ = v45
                            default:
                                var v46 int32 = int32(fi42)
                                _ = v46
                            }
                            _ = len(v1)
                            v47 := GenS4{F0: 2, F1: 624, F2: 65535}
                            _ = v47
                        } else {
                            fmt.Println(fi42, fi42, v2)
                            v3 += fi42
                            if len(v1) > 0 {
                                v1[0] = (^(fi42))
                            }
                        }
                        _ = fi42
                    }
                default:
                    switch (v3 << ((2) & 63)) {
                    case 9223372036854775807:
                        var v48 bool = true
                        _ = v48
                    case 127:
                        if (568 >= 0) {
                            v49 := GenS3{F0: true}
                            for fi50 := 0; fi50 < 7; fi50++ {
                                if len(v1) > 0 {
                                    v1[0] = -1
                                }
                                v51 := GenS4{F0: 0, F1: 0, F2: uint(v4)}
                                fmt.Println(v2, v2)
                                _ = fi50
                                _ = v51
                            }
                            for fi52 := 0; fi52 < 18; fi52++ {
                                var v53 uint = 4294967295
                                v4 = v4
                                _ = fi52
                                _ = v53
                            }
                            _ = v49
                        } else {
                            if (2 <= 255) {
                                for fi54 := 0; fi54 < 7; fi54++ {
                                    fmt.Println(fi54)
                                    if true {
                                        var v55 float32 = 0.0
                                        if true {
                                            var v56 int = fi54
                                            if len(v1) > 0 {
                                                v1[0] = (-(v56))
                                            }
                                            _ = v56
                                        } else {
                                            v57 := GenS5{F0: 65535, F1: 2}
                                            _ = v57
                                        }
                                        var v58 bool = false
                                        _ = v55
                                        _ = v58
                                    }
                                    _ = fi54
                                }
                                for fi59 := 0; fi59 < 10; fi59++ {
                                    fmt.Println(v3)
                                    _ = fi59
                                }
                            } else {
                                var v60 int32 = 1
                                _ = v60
                            }
                            v2 += float64(-2.25)
                            v2 = (v2 + v2)
                        }
                    }
                }
            } else {
                var v61 int64 = 123
                _ = v61
            }
            v3 = (v3 + (v3 ^ v3))
        } else {
            for fi62 := 0; fi62 < 5; fi62++ {
                var v63 bool = false
                _ = fi62
                _ = v63
            }
            v64 := GenS6{F0: 1}
            var v65 uint16 = (v4 | v4)
            _ = v64
            _ = v65
        }
        for fi66 := 0; fi66 < 4; fi66++ {
            var v67 int = v3
            var v68 int32 = -2147483648
            _ = fi66
            _ = v67
            _ = v68
        }
    }
    var v69 int32 = -323
    v4 = genFn1(-1, "\n\t")
    for fi70 := 0; fi70 < 13; fi70++ {
        for fi71 := 0; fi71 < 14; fi71++ {
            if (("a" + "") > "hello") {
                v1 = append(v1, v3)
            } else {
                switch "\n\t" {
                case "foo bar":
                    if (genFn1(-1, "") <= 255) {
                        fmt.Println(fi70)
                        v1 = append(v1, fi70)
                        if (!true) {
                            v72 := GenS7{F0: (^(fi71))}
                            _ = v72
                        } else {
                            v1 = append(v1, (fi71 * (^(fi70))))
                        }
                    } else {
                        var v73 uint32 = 2
                        fmt.Println(v2, v2)
                        _ = v73
                    }
                case "hello":
                    v69 -= v69
                default:
                    v74 := GenS7{F0: v3}
                    _ = v74
                }
            }
            if true {
                v75 := GenS5{F0: 4294967295, F1: 2147483647}
                _ = v75
            }
            switch "hello" {
            case "hello":
                v2 = 3.14159
            case "\n\t":
                v76 := GenS8{F0: (0 < 588), F1: genFn0()}
                _ = v76
            }
            v77 := GenS7{F0: fi70}
            v78 := GenS9{F0: (fi71 - fi70)}
            v1 = append(v1, -490)
            _ = fi71
            _ = v77
            _ = v78
        }
        _ = fi70
    }
    v3 += (v3 * v3)
    for fi79 := 0; fi79 < 4; fi79++ {
        v1 = append(v1, 32767)
        if true {
            var v80 float64 = v2
            _ = v80
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi79
    }
    if (18446744073709551615 < 42) {
        v81 := GenS2{F0: 0}
        _ = v81
    } else {
        var v82 int64 = int64(v3)
        _ = v82
    }
    var v83 int32 = 42
    var v84 uint64 = 1
    v69 -= 2147483647
    if true {
        for fi85 := 0; fi85 < 1; fi85++ {
            var v86 int64 = -1
            v3 -= int(v84)
            _ = fi85
            _ = v86
        }
        v2 += (v2 / 0.1)
        var v87 uint8 = 155
        _ = v87
    } else {
        for fi88 := 0; fi88 < 8; fi88++ {
            if true {
                var v89 uint8 = 1
                fmt.Println(v89, v69)
                _ = v89
            }
            _ = fi88
        }
        switch v3 {
        case 4294967295:
            if (genFn0() > 36) {
                var v90 uint32 = 65535
                _ = v90
            } else {
                if false {
                    if (genFn0() > 1) {
                        switch v3 {
                        case -764:
                            fmt.Println(v2, v4)
                        case 562:
                            if false {
                                fmt.Println(v83)
                                switch v3 {
                                case -2147483648:
                                    var v91 int16 = 32767
                                    _ = v91
                                case 0:
                                    v92 := GenS5{F0: 255, F1: 642}
                                    _ = v92
                                case -492:
                                    var v93 int8 = 1
                                    _ = v93
                                }
                                for fi94 := 0; fi94 < 3; fi94++ {
                                    v95 := GenS4{F0: 4294967295, F1: v84, F2: 0}
                                    if true {
                                        for fi96 := 0; fi96 < 13; fi96++ {
                                            if (0 != (^(v69))) {
                                                v97 := GenS10{F0: 123}
                                                _ = v97
                                            } else {
                                                if true {
                                                    v83 = (v69 | (v69 | (-(v83))))
                                                    var v98 float64 = -2.25
                                                    fmt.Println(v98, v98)
                                                    switch ("" + "") {
                                                    case "foo bar":
                                                        var v99 string = ("\n\t" + "foo bar")
                                                        _ = v99
                                                    case "hello":
                                                        v2 = 0.1
                                                    }
                                                    _ = v98
                                                }
                                            }
                                            var v100 int32 = -1
                                            v100 -= (-(v100))
                                            v101 := GenS7{F0: v3}
                                            _ = fi96
                                            _ = v100
                                            _ = v101
                                        }
                                        v2 -= 1e10
                                    }
                                    v102 := GenS5{F0: 18446744073709551615, F1: 2147483647}
                                    _ = fi94
                                    _ = v102
                                    _ = v95
                                }
                            }
                        case 990:
                            for fi103 := 0; fi103 < 7; fi103++ {
                                if false {
                                    if (!(!false)) {
                                        var v104 float64 = float64((1e10 / (-2.25 + (((float32(0.1) * (-2.25 / (3.14159 + float32(v2)))) / -2.25) - 1e10))))
                                        var v105 uint32 = 1
                                        switch (("\n\t" + "hello") + ("foo bar" + "foo bar")) {
                                        case "\n\t":
                                            if true {
                                                var v106 uint8 = 2
                                                if true {
                                                    v107 := GenS7{F0: (fi103 & (^(fi103)))}
                                                    var v108 uint16 = v4
                                                    _ = v107
                                                    _ = v108
                                                }
                                                _ = v106
                                            }
                                        }
                                        _ = v104
                                        _ = v105
                                    }
                                    var v109 int16 = 133
                                    _ = v109
                                } else {
                                    if true {
                                        v110 := GenS3{F0: true}
                                        if len(v1) > 0 {
                                            v1[0] = fi103
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = (fi103 << ((18446744073709551615) & 63))
                                        }
                                        _ = v110
                                    }
                                    v4 = (v4 >> ((65535) & 63))
                                    v84 += (v84 ^ 42)
                                }
                                if false {
                                    var v111 uint32 = 2147483647
                                    _ = len(v1)
                                    _ = v111
                                }
                                _ = fi103
                            }
                        }
                        var v112 int8 = -19
                        _ = v112
                    } else {
                        if false {
                            var v113 int64 = 1
                            _ = v113
                        }
                        v69 += (v83 & (^(v83)))
                        if ((69 != 215) && ((!true) || true)) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch 32767 {
                            case -1:
                                var v114 uint64 = 0
                                _ = v114
                            case 222:
                                if (71 == 0) {
                                    if (v4 < v4) {
                                        v4 += (v4 >> ((486) & 63))
                                        v115 := GenS5{F0: 1, F1: genFn0()}
                                        _ = v115
                                    } else {
                                        for fi116 := 0; fi116 < 5; fi116++ {
                                            if (165 != 128) {
                                                if ((v84 % ((230) | 1)) != v84) {
                                                    for fi117 := 0; fi117 < 2; fi117++ {
                                                        fmt.Println(v84)
                                                        var v118 uint32 = 360
                                                        _ = fi117
                                                        _ = v118
                                                    }
                                                } else {
                                                    v4 += v4
                                                }
                                                for fi119 := 0; fi119 < 4; fi119++ {
                                                    v84 += (^(v84))
                                                    _ = fi119
                                                }
                                                for fi120 := 0; fi120 < 19; fi120++ {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    if ((v84 - 0) > 645) {
                                                        var v121 int = int(v4)
                                                        v69 = (v83 ^ int32(v4))
                                                        if true {
                                                            var v122 int32 = v83
                                                            _ = v122
                                                        }
                                                        _ = v121
                                                    }
                                                    _ = fi120
                                                }
                                            } else {
                                                fmt.Println(v2, v3, v3)
                                                _ = len(v1)
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                            fmt.Println(fi116, v83, v4)
                                            _ = fi116
                                        }
                                        if false {
                                            fmt.Println(v2)
                                            var v123 bool = false
                                            var v124 uint64 = v84
                                            _ = v123
                                            _ = v124
                                        } else {
                                            var v125 int64 = 1
                                            _ = v125
                                        }
                                    }
                                    switch "a" {
                                    case "a":
                                        var v126 float32 = 3.14159
                                        _ = v126
                                    case "hello":
                                        switch 127 {
                                        case 9223372036854775807:
                                            v127 := GenS11{F0: 457, F1: 32767}
                                            _ = v127
                                        case 5:
                                            if true {
                                                v83 = -988
                                                v2 = 0.1
                                            }
                                        case -2147483648:
                                            switch "" {
                                            case "foo bar":
                                                var v128 float32 = -2.25
                                                _ = v128
                                            case "hello":
                                                switch v3 {
                                                case -512:
                                                    var v129 bool = (65535 <= 977)
                                                    _ = v129
                                                }
                                            case "\n\t":
                                                v130 := GenS12{F0: (((-2.25 + v2) < (3.14159 * 0.1)) && false)}
                                                _ = v130
                                            default:
                                                v84 = (v84 - (^(v84)))
                                            }
                                        }
                                    case "\n\t":
                                        if (true && true) {
                                            if (!((v69 - (v69 * v69)) <= int32(v4))) {
                                                if (-1 != 42) {
                                                    if false {
                                                        v2 = (1e10 + (((((3.14159 - 1e10) + (v2 - v2)) + v2) * (v2 * ((v2 * (float64(-2.25) - ((0.0 / v2) + (v2 * ((float64(3.14159) - v2) - v2))))) + 0.1))) / 3.14159))
                                                        var v131 uint32 = uint32(v84)
                                                        _ = len(v1)
                                                        _ = v131
                                                    } else {
                                                        fmt.Println(v4, v4)
                                                    }
                                                    v132 := GenS7{F0: (v3 & (-(v3)))}
                                                    _ = v132
                                                }
                                            }
                                        } else {
                                            v133 := GenS3{F0: (v84 == v84)}
                                            _ = v133
                                        }
                                    }
                                }
                            default:
                                fmt.Println(v69)
                            }
                        }
                    }
                    for fi134 := 0; fi134 < 6; fi134++ {
                        if false {
                            if false {
                                if true {
                                    fmt.Println(v69)
                                }
                                if len(v1) > 0 {
                                    v1[0] = fi134
                                }
                                fmt.Println(fi134)
                            }
                            var v135 string = "a"
                            fmt.Println(v4, v135)
                            _ = v135
                        }
                        _ = fi134
                    }
                    v136 := GenS13{F0: 144, F1: 1, F2: 422}
                    _ = v136
                } else {
                    v4 += 239
                }
                var v137 string = ""
                if len(v1) > 0 {
                    v1[0] = v3
                }
                _ = v137
            }
        case 124:
            if (!true) {
                fmt.Println(v3)
            }
        case -2147483648:
            if true {
                switch "\n\t" {
                case "\n\t":
                    v84 += 955
                default:
                    v2 -= (((v2 * 0.1) * float64(3.14159)) + (v2 - 0.0))
                }
                v2 = v2
            } else {
                if false {
                    switch ("a" + "a") {
                    case "":
                        var v138 int64 = 42
                        _ = v138
                    case "\n\t":
                        if true {
                            if ("\n\t" > "a") {
                                v83 += v83
                                var v139 uint16 = (v4 - v4)
                                v140 := GenS14{F0: 0}
                                _ = v139
                                _ = v140
                            }
                            if (-413 > 0) {
                                for fi141 := 0; fi141 < 11; fi141++ {
                                    v69 -= v83
                                    _ = fi141
                                }
                            } else {
                                if (true && (94 >= 127)) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if (((true && true) || (((0.0 * 0.1) / -2.25) < 1e10)) && true) {
                                        switch (v3 >> ((1) & 63)) {
                                        case 270:
                                            v83 += v69
                                        default:
                                            if false {
                                                switch (^(v3)) {
                                                case -259:
                                                    var v142 int64 = 42
                                                    _ = v142
                                                case -100:
                                                    switch 1 {
                                                    case 42:
                                                        var v143 int16 = 32767
                                                        _ = v143
                                                    case -789:
                                                        var v144 uint32 = 990
                                                        _ = v144
                                                    case 32767:
                                                        v84 = v84
                                                    default:
                                                        var v145 int32 = v69
                                                        _ = v145
                                                    }
                                                default:
                                                    v84 = uint64(v4)
                                                }
                                                fmt.Println(v83, v83, v2)
                                            } else {
                                                switch "" {
                                                case "foo bar":
                                                    for fi146 := 0; fi146 < 7; fi146++ {
                                                        switch "hello" {
                                                        case "hello":
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        case "\n\t":
                                                            v147 := GenS13{F0: 127, F1: 463, F2: -9223372036854775808}
                                                            _ = v147
                                                        default:
                                                            v84 -= (v84 >> ((0) & 63))
                                                        }
                                                        v148 := GenS1{F0: 469, F1: 9223372036854775807}
                                                        v4 -= genFn1(1, ("foo bar" + "foo bar"))
                                                        _ = fi146
                                                        _ = v148
                                                    }
                                                case "\n\t":
                                                    var v149 int8 = 42
                                                    _ = v149
                                                case "a":
                                                    var v150 uint = 18446744073709551615
                                                    _ = v150
                                                default:
                                                    _ = len(v1)
                                                }
                                                fmt.Println(v69)
                                            }
                                        }
                                        if true {
                                            var v151 uint32 = 2147483648
                                            var v152 uint = 483
                                            fmt.Println(v4)
                                            _ = v151
                                            _ = v152
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            if false {
                                                fmt.Println(v69, v83)
                                                v1 = append(v1, v3)
                                            } else {
                                                v153 := GenS13{F0: 42, F1: 306, F2: 1}
                                                for fi154 := 0; fi154 < 3; fi154++ {
                                                    var v155 uint16 = v4
                                                    var v156 int16 = 861
                                                    v157 := GenS11{F0: 255, F1: (v155 ^ v155)}
                                                    _ = fi154
                                                    _ = v155
                                                    _ = v156
                                                    _ = v157
                                                }
                                                for fi158 := 0; fi158 < 6; fi158++ {
                                                    switch v3 {
                                                    case 1:
                                                        for fi159 := 0; fi159 < 1; fi159++ {
                                                            if true {
                                                                if ((v84 / ((v84) | 1)) != v84) {
                                                                    v84 -= v84
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        switch (fi159 * (fi159 | v3)) {
                                                                        case 9:
                                                                            switch ("\n\t" + "a") {
                                                                            case "foo bar":
                                                                                v2 = float64(-2.25)
                                                                            case "hello":
                                                                                var v160 float32 = 0.1
                                                                                _ = v160
                                                                            case "":
                                                                                switch ("foo bar" + ("" + "\n\t")) {
                                                                                case "foo bar":
                                                                                    v83 = v69
                                                                                }
                                                                            }
                                                                        case 2147483647:
                                                                            v69 = v69
                                                                        default:
                                                                            v2 = v2
                                                                        }
                                                                    case "\n\t":
                                                                        if (775 <= 2) {
                                                                            if len(v1) > 0 {
                                                                                v1[0] = -9223372036854775808
                                                                            }
                                                                            var v161 int64 = -1
                                                                            _ = v161
                                                                        } else {
                                                                            switch 1 {
                                                                            case 32768:
                                                                                var v162 int32 = -288
                                                                                _ = v162
                                                                            default:
                                                                                var v163 int64 = 0
                                                                                _ = v163
                                                                            }
                                                                            var v164 float32 = (1.5 * ((1.5 + 1.5) + 1.5))
                                                                            _ = v164
                                                                        }
                                                                    case "a":
                                                                        if len(v1) > 0 {
                                                                            v1[0] = fi158
                                                                        }
                                                                    default:
                                                                        if len(v1) > 0 {
                                                                            v1[0] = (fi159 * fi158)
                                                                        }
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        v2 = 3.14159
                                                                        var v165 int64 = -9223372036854775808
                                                                        var v166 float64 = (0.1 - v2)
                                                                        _ = v165
                                                                        _ = v166
                                                                    } else {
                                                                        var v167 bool = false
                                                                        if v167 {
                                                                            v83 -= (v69 ^ v69)
                                                                        }
                                                                        var v168 int16 = 1000
                                                                        _ = v167
                                                                        _ = v168
                                                                    }
                                                                }
                                                                switch (fi158 - int(v4)) {
                                                                case 128:
                                                                    var v169 uint64 = v84
                                                                    _ = v169
                                                                case 32767:
                                                                    v170 := GenS7{F0: 32767}
                                                                    _ = v170
                                                                }
                                                                if true {
                                                                    fmt.Println(v83, v4, v3)
                                                                    var v171 int8 = -1
                                                                    var v172 uint16 = (v4 & v4)
                                                                    v171 = -128
                                                                    var v173 string = "hello"
                                                                    fmt.Println(v172)
                                                                    if (v84 != (v84 & v84)) {
                                                                        var v174 uint = uint(v84)
                                                                        fmt.Println(v4, v3)
                                                                        var v175 bool = (!true)
                                                                        _ = v174
                                                                        _ = v175
                                                                    } else {
                                                                        var v176 uint8 = 127
                                                                        fmt.Println(v176, v4)
                                                                        _ = v176
                                                                    }
                                                                    _ = v171
                                                                    _ = v172
                                                                    _ = v173
                                                                } else {
                                                                    var v177 int = (v3 & (v3 ^ v3))
                                                                    _ = v177
                                                                }
                                                                _ = len(v1)
                                                            }
                                                            _ = fi159
                                                        }
                                                    case -9223372036854775808:
                                                        var v178 uint16 = (v4 - (v4 >> ((2) & 63)))
                                                        _ = v178
                                                    }
                                                    var v179 int32 = (v69 - (v83 * (-(v69))))
                                                    _ = fi158
                                                    _ = v179
                                                }
                                                _ = v153
                                            }
                                            if (!(!true)) {
                                                v2 += 1e10
                                            } else {
                                                for fi180 := 0; fi180 < 15; fi180++ {
                                                    if false {
                                                        for fi181 := 0; fi181 < 10; fi181++ {
                                                            v83 = 516
                                                            v182 := GenS4{F0: uint(v84), F1: v84, F2: 2}
                                                            _ = fi181
                                                            _ = v182
                                                        }
                                                        for fi183 := 0; fi183 < 10; fi183++ {
                                                            switch (fi183 & fi183) {
                                                            case -884:
                                                                var v184 int = fi180
                                                                _ = v184
                                                            }
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            if ((!(-58 > -30)) || false) {
                                                                var v185 string = ("foo bar" + "")
                                                                v186 := GenS15{F0: int(v69), F1: -685, F2: 0}
                                                                fmt.Println(fi180, v3)
                                                                _ = v185
                                                                _ = v186
                                                            }
                                                            _ = fi183
                                                        }
                                                        switch (v3 * (fi180 * v3)) {
                                                        case 9223372036854775807:
                                                            v3 = v3
                                                        case 42:
                                                            v4 += genFn1(100, "foo bar")
                                                        case 436:
                                                            var v187 int8 = 0
                                                            _ = v187
                                                        default:
                                                            v69 = (v83 % ((2147483647) | 1))
                                                        }
                                                    }
                                                    _ = fi180
                                                }
                                            }
                                        }
                                    } else {
                                        for fi188 := 0; fi188 < 5; fi188++ {
                                            fmt.Println(fi188, v3)
                                            switch ("a" + "") {
                                            case "":
                                                var v189 int64 = 846
                                                _ = v189
                                            default:
                                                _ = len(v1)
                                            }
                                            _ = fi188
                                        }
                                        switch ("a" + "hello") {
                                        case "foo bar":
                                            v2 = v2
                                        }
                                    }
                                } else {
                                    var v190 uint = 65535
                                    _ = v190
                                }
                                var v191 int64 = 1
                                v4 = (^(v4))
                                _ = v191
                            }
                            var v192 int16 = 1000
                            _ = v192
                        }
                    case "a":
                        fmt.Println(v83, v84, v4)
                    default:
                        if false {
                            if ((v3 % ((v3) | 1)) >= (v3 - 32768)) {
                                fmt.Println(v3, v84, v2)
                                var v193 uint16 = 0
                                _ = v193
                            } else {
                                v194 := GenS16{F0: 255, F1: 950, F2: 42}
                                v83 += 42
                                var v195 int16 = -1
                                _ = v194
                                _ = v195
                            }
                            v84 = 388
                        } else {
                            var v196 int32 = v69
                            v197 := GenS17{F0: -524, F1: 2, F2: 42}
                            v198 := GenS15{F0: -2147483648, F1: -9223372036854775808, F2: 204}
                            _ = v196
                            _ = v197
                            _ = v198
                        }
                    }
                    v1 = append(v1, -1)
                }
                switch "foo bar" {
                case "":
                    if (genFn0() < 171) {
                        if false {
                            var v199 int64 = 0
                            _ = v199
                        }
                        for fi200 := 0; fi200 < 2; fi200++ {
                            v2 = (v2 * (0.1 / float64(0.1)))
                            _ = fi200
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                case "\n\t":
                    _ = len(v1)
                case "hello":
                    fmt.Println(v2, v83, v83)
                default:
                    fmt.Println(v84, v3)
                }
            }
        default:
            for fi201 := 0; fi201 < 7; fi201++ {
                if false {
                    v83 -= (v83 - (^(v83)))
                    v4 = 32767
                } else {
                    if (false && ((!false) || false)) {
                        var v202 int32 = (v83 * (v69 + (-(v69))))
                        v69 += v83
                        switch "foo bar" {
                        case "foo bar":
                            v2 += v2
                        case "\n\t":
                            for fi203 := 0; fi203 < 14; fi203++ {
                                v4 = v4
                                fmt.Println(v69)
                                _ = fi203
                            }
                        case "a":
                            var v204 uint16 = (v4 & 42)
                            _ = v204
                        default:
                            v1 = append(v1, 32767)
                        }
                        _ = v202
                    }
                    v2 = 1e10
                }
                var v205 uint = uint(v84)
                v83 = v83
                _ = fi201
                _ = v205
            }
        }
        for fi206 := 0; fi206 < 19; fi206++ {
            v4 = v4
            if true {
                if ((v83 * int32(v3)) >= (v83 ^ (v69 ^ 0))) {
                    if true {
                        var v207 uint8 = 0
                        if true {
                            v69 += (v69 - v83)
                        } else {
                            if ((v3 + 9223372036854775807) == v3) {
                                for fi208 := 0; fi208 < 19; fi208++ {
                                    var v209 uint16 = uint16(v207)
                                    _ = fi208
                                    _ = v209
                                }
                                _ = len(v1)
                                v210 := GenS18{F0: v83, F1: -42}
                                _ = v210
                            }
                        }
                        if true {
                            v211 := GenS8{F0: (true || true), F1: 851}
                            v4 -= (v4 << ((278) & 63))
                            var v212 int = 32767
                            _ = v211
                            _ = v212
                        } else {
                            v1 = append(v1, v3)
                        }
                        _ = v207
                    }
                    var v213 uint16 = genFn1(-107, ("hello" + (("a" + "hello") + "a")))
                    if ((2 < 2147483647) || false) {
                        v213 = (v213 / ((v4) | 1))
                        v214 := GenS19{F0: 512, F1: (v69 >> ((42) & 63))}
                        _ = v214
                    } else {
                        var v215 int = (v3 * (v3 >> ((65535) & 63)))
                        _ = len(v1)
                        if true {
                            v216 := GenS1{F0: 255, F1: -484}
                            fmt.Println(v215)
                            var v217 int16 = 42
                            _ = v216
                            _ = v217
                        } else {
                            for fi218 := 0; fi218 < 4; fi218++ {
                                if len(v1) > 0 {
                                    v1[0] = 2147483647
                                }
                                var v219 int32 = v83
                                _ = fi218
                                _ = v219
                            }
                            var v220 uint16 = 526
                            _ = len(v1)
                            _ = v220
                        }
                        _ = v215
                    }
                    _ = v213
                }
                v221 := GenS11{F0: 1, F1: genFn1(-1, "")}
                v222 := GenS20{F0: 2, F1: false}
                _ = v221
                _ = v222
            }
            if len(v1) > 0 {
                v1[0] = (fi206 % (((fi206 % ((-9223372036854775808) | 1))) | 1))
            }
            _ = fi206
        }
    }
    v223 := GenS13{F0: 1, F1: 865, F2: -1}
    fmt.Println(v2)
    switch v3 {
    case -779:
        var v224 int8 = -1
        _ = v224
    case 9223372036854775807:
        for fi225 := 0; fi225 < 11; fi225++ {
            for fi226 := 0; fi226 < 9; fi226++ {
                v227 := GenS12{F0: (!true)}
                var v228 uint64 = v84
                fmt.Println(v84, v228)
                var v229 int64 = -69
                _ = fi226
                _ = v227
                _ = v228
                _ = v229
            }
            var v230 uint8 = 2
            v231 := GenS18{F0: (v83 % (((v83 | v69)) | 1)), F1: -128}
            _ = fi225
            _ = v230
            _ = v231
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v69:", v69)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v223
    _ = v3
    _ = v4
    _ = v69
    _ = v83
    _ = v84
}
