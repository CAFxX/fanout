package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 string; F1 uint64; F2 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS6 struct { F0 uint64; F1 int16; F2 int32 }
type GenS7 struct { F0 int32; F1 int8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 string; F1 uint64 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int32; F1 uint; F2 uint64 }
type GenS12 struct { F0 string; F1 bool }
type GenS13 struct { F0 uint8; F1 int32; F2 uint32 }
func genFn0() int8 {
    var v7 uint16 = 2
    if ((1 != 0) && false) {
        for fi8 := 0; fi8 < 14; fi8++ {
            v7 += (^(v7))
            for fi9 := 0; fi9 < 5; fi9++ {
                var v10 int64 = -263
                fmt.Println(fi8, v10, fi9)
                if (1 >= v10) {
                    v10 = (v10 / (((-(v10))) | 1))
                }
                v10 = (v10 % (((v10 % (((^(v10))) | 1))) | 1))
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
        var v11 int = 42
        for fi12 := 0; fi12 < 17; fi12++ {
            v7 = v7
            _ = fi12
        }
        _ = v11
    } else {
        for fi13 := 0; fi13 < 13; fi13++ {
            v14 := GenS1{F0: 103}
            fmt.Println(fi13, v7)
            _ = fi13
            _ = v14
        }
        fmt.Println(v7)
        switch 0 {
        case 127:
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 uint8 = 0
                switch ("a" + "hello") {
                case "foo bar":
                    var v17 float64 = 0.0
                    _ = v17
                case "a":
                    var v18 bool = ((!false) && (!false))
                    _ = v18
                case "":
                    var v19 float32 = 1.5
                    _ = v19
                }
                _ = fi15
                _ = v16
            }
        default:
            var v20 uint16 = v7
            _ = v20
        }
    }
    var v21 int = 9223372036854775807
    if false {
        v22 := GenS1{F0: 1}
        switch 32768 {
        case 608:
            switch (v21 | v21) {
            case 4294967295:
                switch v21 {
                case 2147483647:
                    fmt.Println(v7, v7)
                case 799:
                    v23 := GenS1{F0: 181}
                    _ = v23
                case 850:
                    v24 := GenS1{F0: 1}
                    _ = v24
                }
            case -437:
                var v25 uint8 = 156
                _ = v25
            default:
                v26 := []int{v21, v21, (-(v21))}
                _ = v26
            }
        case -260:
            fmt.Println(v21)
        }
        for fi27 := 0; fi27 < 16; fi27++ {
            for fi28 := 0; fi28 < 14; fi28++ {
                v29 := GenS1{F0: 255}
                _ = fi28
                _ = v29
            }
            var v30 uint64 = 0
            _ = fi27
            _ = v30
        }
        _ = v22
    } else {
        v21 -= v21
        var v31 uint8 = 0
        var v32 int = v21
        _ = v31
        _ = v32
    }
    v21 -= v21
    _ = v21
    _ = v7
    return 1
}

func main() {
    v1 := []int{32768}
    var v2 int = 0
    var v3 float32 = ((0.1 - 1e10) * float32((((1.5 / ((-2.25 / float64((((((3.14159 - (1e10 * 0.1)) - (1.5 * 0.1)) + -2.25) - (0.0 * 3.14159)) / 0.1))) * 0.1)) / 1.0) + ((1.5 / float64(1e10)) + float64((3.14159 / (-2.25 / (1.5 * 1e10))))))))
    v4 := []int{v2, (v2 * 32767), 4294967295, v2}
    var v5 string = ("hello" + "")
    var v6 uint8 = 169
    fmt.Println(v6, v6)
    _ = len(v1)
    for fi33 := 0; fi33 < 1; fi33++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (0 != 18446744073709551615) {
            if true {
                v2 = -22
                for fi34 := 0; fi34 < 20; fi34++ {
                    fmt.Println(v6, v2)
                    v3 = (v3 * ((float32(float64(-2.25)) - -2.25) * v3))
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(v2, v5)
                    _ = fi34
                }
                if true {
                    var v35 int16 = 42
                    _ = v35
                } else {
                    if false {
                        var v36 int = (v2 + 2147483647)
                        v5 = ""
                        _ = v36
                    }
                    switch 0 {
                    case 1:
                        if (false || (-123 <= genFn0())) {
                            for fi37 := 0; fi37 < 7; fi37++ {
                                v38 := GenS1{F0: 18446744073709551615}
                                var v39 float32 = (-2.25 + (v3 - v3))
                                var v40 int8 = 96
                                if true {
                                    _ = len(v4)
                                    if (42 < 1) {
                                        var v41 bool = (!false)
                                        switch "\n\t" {
                                        case "hello":
                                            v6 -= (v6 >> ((1) & 63))
                                        default:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        fmt.Println(v6)
                                        _ = v41
                                    }
                                    var v42 uint16 = 1
                                    _ = v42
                                }
                                _ = fi37
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                        }
                    case 9223372036854775807:
                        fmt.Println(v2, v5)
                    }
                    if (float64((float32(float64(v3)) * v3)) <= -2.25) {
                        for fi43 := 0; fi43 < 17; fi43++ {
                            var v44 uint64 = 680
                            v45 := GenS1{F0: 0}
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                        v6 = v6
                    } else {
                        var v46 int64 = 0
                        if ((v6 ^ (v6 >> ((2) & 63))) > (v6 | v6)) {
                            for fi47 := 0; fi47 < 1; fi47++ {
                                v5 = v5
                                _ = fi47
                            }
                            if (false || true) {
                                v3 = 0.1
                                if false {
                                    var v48 uint16 = 65535
                                    _ = v48
                                }
                                if len(v4) > 0 {
                                    v4[0] = (fi33 * (fi33 & 0))
                                }
                            } else {
                                for fi49 := 0; fi49 < 10; fi49++ {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    fmt.Println(v6, v3)
                                    var v50 int64 = 42
                                    _ = fi49
                                    _ = v50
                                }
                            }
                            var v51 float64 = 3.14159
                            _ = v51
                        }
                        for fi52 := 0; fi52 < 13; fi52++ {
                            var v53 uint32 = 528
                            var v54 string = "hello"
                            switch "a" {
                            case "hello":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            _ = fi52
                            _ = v53
                            _ = v54
                        }
                        _ = v46
                    }
                }
            } else {
                v55 := GenS2{F0: 32767, F1: 1}
                if false {
                    v5 = "foo bar"
                }
                for fi56 := 0; fi56 < 20; fi56++ {
                    v57 := GenS1{F0: 18446744073709551615}
                    v1 = append(v1, fi33)
                    _ = fi56
                    _ = v57
                }
                _ = v55
            }
            switch (((v5 + (v5 + (((((v5 + v5) + ("foo bar" + v5)) + v5) + (v5 + "a")) + "a"))) + (v5 + "hello")) + v5) {
            case "foo bar":
                fmt.Println(v6, fi33)
            case "a":
                v58 := GenS1{F0: 42}
                _ = v58
            case "":
                v6 -= 2
            }
        }
        switch v5 {
        case "":
            v2 = fi33
        case "a":
            fmt.Println(fi33, v2, fi33)
        case "hello":
            if (!(2147483648 != 987)) {
                var v59 int = (v2 + 4294967295)
                v2 = (fi33 & -1)
                if true {
                    v2 -= (v59 ^ (v2 * (v59 / ((v2) | 1))))
                } else {
                    fmt.Println(v6)
                    _ = len(v4)
                }
                _ = v59
            } else {
                if (v2 == (fi33 + 4294967295)) {
                    var v60 int16 = -1
                    _ = v60
                } else {
                    v2 -= fi33
                }
            }
        default:
            v1 = append(v1, 127)
        }
        _ = fi33
    }
    switch ("foo bar" + (v5 + "")) {
    case "foo bar":
        for fi61 := 0; fi61 < 7; fi61++ {
            var v62 int16 = 1000
            var v63 float32 = float32(0.1)
            if (135 < 2) {
                var v64 uint64 = 718
                v65 := GenS3{F0: "", F1: 42, F2: fi61}
                _ = v64
                _ = v65
            } else {
                v66 := GenS4{F0: 65535}
                var v67 int64 = 0
                if len(v4) > 0 {
                    v4[0] = (fi61 - v2)
                }
                _ = v66
                _ = v67
            }
            _ = fi61
            _ = v62
            _ = v63
        }
    default:
        if (false && (false && (true && false))) {
            _ = len(v4)
            var v68 int = v2
            v5 = v5
            _ = v68
        }
    }
    for fi69 := 0; fi69 < 7; fi69++ {
        v6 -= (^(v6))
        v4 = append(v4, (v2 + v2))
        var v70 int = (v2 / (((fi69 * -1)) | 1))
        _ = fi69
        _ = v70
    }
    v6 = 58
    if len(v4) > 0 {
        v4[0] = (v2 - 9223372036854775807)
    }
    var v71 uint16 = 255
    for fi72 := 0; fi72 < 13; fi72++ {
        for fi73 := 0; fi73 < 2; fi73++ {
            var v74 uint16 = v71
            _ = fi73
            _ = v74
        }
        fmt.Println(v6, v6, v5)
        _ = fi72
    }
    var v75 uint64 = 2
    if (785 >= -311) {
        if (498 != 255) {
            fmt.Println(v75, v75)
            switch "a" {
            case "a":
                var v76 uint16 = v71
                _ = v76
            default:
                var v77 int = (v2 >> ((847) & 63))
                _ = v77
            }
        }
        if ((v71 ^ uint16(v6)) >= (v71 - v71)) {
            v75 -= 0
            v75 -= 0
        }
        if ((^(v6)) >= 255) {
            for fi78 := 0; fi78 < 4; fi78++ {
                v79 := GenS2{F0: -32768, F1: 739}
                v80 := GenS5{F0: 837, F1: -128, F2: v6}
                if true {
                    for fi81 := 0; fi81 < 6; fi81++ {
                        v82 := GenS2{F0: 477, F1: 1000}
                        _ = fi81
                        _ = v82
                    }
                }
                _ = fi78
                _ = v79
                _ = v80
            }
        }
    } else {
        var v83 uint = 65535
        if false {
            var v84 int8 = 24
            switch v5 {
            case "foo bar":
                for fi85 := 0; fi85 < 16; fi85++ {
                    var v86 uint32 = 0
                    v3 = v3
                    var v87 bool = true
                    _ = fi85
                    _ = v86
                    _ = v87
                }
            case "hello":
                if len(v1) > 0 {
                    v1[0] = v2
                }
            case "\n\t":
                if false {
                    switch ((("foo bar" + "\n\t") + "foo bar") + v5) {
                    case "hello":
                        for fi88 := 0; fi88 < 18; fi88++ {
                            if false {
                                for fi89 := 0; fi89 < 6; fi89++ {
                                    v2 = fi88
                                    v83 -= 42
                                    _ = fi89
                                }
                                var v90 uint8 = (v6 * (v6 | v6))
                                switch v2 {
                                case 128:
                                    switch v5 {
                                    case "hello":
                                        _ = len(v1)
                                    }
                                case 2147483647:
                                    fmt.Println(v3, v83, v5)
                                }
                                _ = v90
                            }
                            _ = fi88
                        }
                    case "\n\t":
                        if false {
                            if false {
                                v83 -= 18446744073709551615
                                for fi91 := 0; fi91 < 2; fi91++ {
                                    var v92 float32 = (v3 * v3)
                                    var v93 bool = false
                                    for fi94 := 0; fi94 < 15; fi94++ {
                                        var v95 uint = v83
                                        var v96 float32 = float32(0.1)
                                        _ = fi94
                                        _ = v95
                                        _ = v96
                                    }
                                    _ = fi91
                                    _ = v92
                                    _ = v93
                                }
                                if len(v4) > 0 {
                                    v4[0] = (v2 & (v2 << ((v83) & 63)))
                                }
                            } else {
                                switch (v5 + v5) {
                                case "foo bar":
                                    fmt.Println(v71, v75)
                                case "\n\t":
                                    if len(v1) > 0 {
                                        v1[0] = (v2 & (-(v2)))
                                    }
                                default:
                                    v84 += (v84 | v84)
                                }
                                v4 = append(v4, (v2 + v2))
                            }
                            if (3.14159 >= (0.1 / 1e10)) {
                                v4 = append(v4, v2)
                                for fi97 := 0; fi97 < 12; fi97++ {
                                    v1 = append(v1, v2)
                                    _ = fi97
                                }
                                for fi98 := 0; fi98 < 15; fi98++ {
                                    v2 -= (fi98 * 65)
                                    for fi99 := 0; fi99 < 4; fi99++ {
                                        if len(v4) > 0 {
                                            v4[0] = 52
                                        }
                                        v75 = (v75 % ((v75) | 1))
                                        var v100 float64 = 1.5
                                        _ = fi99
                                        _ = v100
                                    }
                                    _ = fi98
                                }
                            } else {
                                fmt.Println(v84)
                                switch v5 {
                                case "":
                                    var v101 float32 = v3
                                    _ = v101
                                default:
                                    var v102 uint8 = v6
                                    _ = v102
                                }
                            }
                        }
                    }
                    _ = len(v4)
                }
            }
            _ = v84
        } else {
            if false {
                v75 -= v75
            }
        }
        _ = v83
    }
    if (v6 < v6) {
        var v103 int16 = 32767
        v71 = (^(v71))
        v104 := GenS5{F0: (v71 << ((1) & 63)), F1: int8(v71), F2: 128}
        _ = v103
        _ = v104
    } else {
        if ("foo bar" >= v5) {
            v5 = v5
            v75 = (v75 * v75)
        } else {
            switch v5 {
            case "hello":
                v105 := GenS6{F0: v75, F1: -32768, F2: 42}
                _ = v105
            case "a":
                switch 0 {
                case -815:
                    for fi106 := 0; fi106 < 1; fi106++ {
                        v3 += (1.5 * -2.25)
                        if len(v1) > 0 {
                            v1[0] = fi106
                        }
                        _ = fi106
                    }
                }
            case "\n\t":
                for fi107 := 0; fi107 < 3; fi107++ {
                    if len(v4) > 0 {
                        v4[0] = (v2 & v2)
                    }
                    v3 -= v3
                    _ = fi107
                }
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            switch v2 {
            case -1:
                var v108 int64 = 9223372036854775807
                _ = v108
            case 42:
                switch -137 {
                case 127:
                    v2 = (^(v2))
                case 4294967295:
                    switch v5 {
                    case "a":
                        for fi109 := 0; fi109 < 17; fi109++ {
                            var v110 int = fi109
                            _ = fi109
                            _ = v110
                        }
                    case "hello":
                        v1 = append(v1, v2)
                    case "\n\t":
                        if (v75 <= 0) {
                            v75 = 42
                            v5 = v5
                        }
                    default:
                        v71 = (v71 * uint16(v6))
                    }
                default:
                    v6 += 143
                }
            case -9223372036854775808:
                _ = len(v1)
            default:
                v111 := GenS3{F0: v5, F1: 4294967295, F2: (v2 / ((v2) | 1))}
                _ = v111
            }
            v5 = v5
        }
        var v112 int64 = 0
        _ = v112
    }
    var v113 float32 = (1e10 / v3)
    for fi114 := 0; fi114 < 3; fi114++ {
        fmt.Println(v71)
        _ = fi114
    }
    v6 = (v6 * (^(v6)))
    switch v2 {
    case -364:
        switch v2 {
        case 914:
            switch v5 {
            case "a":
                switch v5 {
                case "":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    if true {
                        var v115 int = (v2 ^ -9223372036854775808)
                        _ = v115
                    }
                }
            case "foo bar":
                for fi116 := 0; fi116 < 1; fi116++ {
                    var v117 uint64 = v75
                    if (!false) {
                        if false {
                            fmt.Println(fi116, v3, v3)
                            var v118 bool = (!(4294967295 <= 18446744073709551615))
                            _ = v118
                        }
                        var v119 int8 = 110
                        _ = v119
                    } else {
                        var v120 int = (fi116 << ((1) & 63))
                        v5 = v5
                        for fi121 := 0; fi121 < 7; fi121++ {
                            if true {
                                switch "\n\t" {
                                case "hello":
                                    v122 := GenS7{F0: -443, F1: 0}
                                    _ = v122
                                case "":
                                    var v123 int8 = 100
                                    _ = v123
                                case "foo bar":
                                    var v124 float32 = float32(1.5)
                                    _ = v124
                                }
                                v1 = append(v1, (fi116 + int(v71)))
                                v75 -= v117
                                switch v5 {
                                case "a":
                                    if false {
                                        switch v120 {
                                        case 375:
                                            if len(v4) > 0 {
                                                v4[0] = (v120 / (((v2 % (((fi121 & (-(fi116)))) | 1))) | 1))
                                            }
                                        default:
                                            if false {
                                                v125 := GenS8{F0: 2}
                                                _ = v125
                                            } else {
                                                var v126 string = v5
                                                v127 := GenS9{F0: v126, F1: 1}
                                                _ = v126
                                                _ = v127
                                            }
                                        }
                                    }
                                default:
                                    if true {
                                        if (false && true) {
                                            switch 2147483647 {
                                            case -985:
                                                v113 = 3.14159
                                                var v128 uint32 = 4294967295
                                                if true {
                                                    _ = len(v4)
                                                } else {
                                                    v129 := GenS8{F0: 1}
                                                    _ = v129
                                                }
                                                if false {
                                                    _ = len(v1)
                                                    v3 = 1.5
                                                } else {
                                                    v130 := GenS10{F0: 62}
                                                    _ = len(v1)
                                                    _ = v130
                                                }
                                                v2 += (v120 << ((65535) & 63))
                                                if ((-(v2)) == (v2 ^ (fi116 & fi121))) {
                                                    _ = len(v4)
                                                } else {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                v131 := GenS6{F0: (v75 | 641), F1: -463, F2: 878}
                                                _ = v128
                                                _ = v131
                                            case 32768:
                                                v3 = 1e10
                                            }
                                        } else {
                                            switch "foo bar" {
                                            case "\n\t":
                                                v1 = append(v1, (v2 << ((855) & 63)))
                                            default:
                                                v71 = (^(v71))
                                            }
                                            var v132 int8 = 102
                                            fmt.Println(v6, v5, v5)
                                            _ = v132
                                        }
                                        var v133 int32 = int32(v71)
                                        v120 -= (fi116 / ((32767) | 1))
                                        _ = v133
                                    }
                                }
                                v71 -= (v71 % (((v71 * 2)) | 1))
                            } else {
                                v6 -= 1
                            }
                            fmt.Println(v5)
                            if false {
                                v134 := GenS11{F0: -513, F1: 923, F2: v117}
                                _ = v134
                            } else {
                                v1 = append(v1, (fi121 ^ v120))
                                if len(v1) > 0 {
                                    v1[0] = 4294967295
                                }
                            }
                            _ = fi121
                        }
                        _ = v120
                    }
                    _ = fi116
                    _ = v117
                }
            default:
                var v135 int = (v2 + (v2 * (v2 / (((v2 ^ (v2 - v2))) | 1))))
                _ = v135
            }
        case -2147483648:
            var v136 float32 = (((0.0 + 1e10) + 1e10) - v3)
            _ = v136
        default:
            var v137 int16 = -1
            _ = v137
        }
    case 2147483647:
        for fi138 := 0; fi138 < 20; fi138++ {
            v139 := GenS2{F0: 42, F1: 191}
            if (true || false) {
                if false {
                    var v140 int32 = -723
                    v6 = v6
                    v113 += (0.0 / float32(0.1))
                    _ = v140
                } else {
                    for fi141 := 0; fi141 < 19; fi141++ {
                        switch (("hello" + v5) + "a") {
                        case "hello":
                            if (!(!true)) {
                                v5 = v5
                                v113 += (1e10 + 3.14159)
                            } else {
                                v71 = v71
                                v142 := GenS3{F0: v5, F1: v75, F2: fi138}
                                v143 := GenS12{F0: "foo bar", F1: false}
                                _ = v142
                                _ = v143
                            }
                        }
                        v3 = v3
                        v144 := GenS13{F0: (v6 + v6), F1: 1, F2: 613}
                        _ = fi141
                        _ = v144
                    }
                    var v145 uint8 = (^(v6))
                    switch v5 {
                    case "hello":
                        if (!(133 >= 18446744073709551615)) {
                            v146 := GenS2{F0: 32767, F1: 516}
                            var v147 uint8 = v145
                            _ = v146
                            _ = v147
                        } else {
                            var v148 int32 = -1
                            v2 = v2
                            v148 = v148
                            _ = v148
                        }
                    case "foo bar":
                        for fi149 := 0; fi149 < 17; fi149++ {
                            var v150 uint32 = 2147483647
                            var v151 float32 = (0.1 * v113)
                            _ = fi149
                            _ = v150
                            _ = v151
                        }
                    case "":
                        fmt.Println(v5)
                    }
                    _ = v145
                }
                v5 = v5
                if len(v4) > 0 {
                    v4[0] = fi138
                }
            }
            _ = fi138
            _ = v139
        }
    }
    v152 := GenS5{F0: (v71 >> ((1) & 63)), F1: 100, F2: v6}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v113:", v113)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v71:", v71)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v113
    _ = v152
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v71
    _ = v75
}
