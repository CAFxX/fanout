package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 int8; F1 int32; F2 uint8 }
type GenS7 struct { F0 uint64; F1 int8 }
type GenS8 struct { F0 int32; F1 bool; F2 int64 }
type GenS9 struct { F0 int16; F1 string }
type GenS10 struct { F0 int32; F1 uint16; F2 bool }
type GenS11 struct { F0 int32; F1 int16; F2 uint64 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int8; F1 int16 }
type GenS14 struct { F0 bool; F1 int8; F2 int32 }
type GenS15 struct { F0 uint16; F1 int16 }
type GenS16 struct { F0 string; F1 int }
func genFn0(p0 int32, p1 uint, p2 uint) int8 {
    v5 := GenS1{F0: 1, F1: 1000}
    v6 := GenS1{F0: 42, F1: 375}
    p1 = p2
    if true {
        p2 -= (p2 << (((p1 - p1)) & 63))
    }
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 int32) int32 {
    fmt.Println(p0)
    v7 := GenS2{F0: 128}
    var v8 int16 = 0
    var v9 uint64 = 161
    v10 := GenS3{F0: 123, F1: 1}
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return p0
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 2
    v3 := make(map[uint32]int)
    var v4 uint = 0
    fmt.Println(v2, v4)
    v11 := GenS3{F0: 0, F1: 0}
    var v12 int16 = -401
    v4 = v4
    switch -2147483648 {
    case 9223372036854775807:
        if (v12 <= v12) {
            v13 := GenS2{F0: 1}
            if false {
                switch "hello" {
                case "foo bar":
                    v2 = 0
                case "a":
                    if true {
                        switch 615 {
                        case -1:
                            v14 := v3[2147483648]
                            _ = v14
                        case 128:
                            if (0 == 32768) {
                                for fi15 := 0; fi15 < 12; fi15++ {
                                    var v16 int32 = 231
                                    switch "foo bar" {
                                    case "a":
                                        if false {
                                            fmt.Println(fi15)
                                            for fi17 := 0; fi17 < 12; fi17++ {
                                                v4 += v4
                                                fmt.Println(v2)
                                                _ = fi17
                                            }
                                        } else {
                                            var v18 int16 = v12
                                            fmt.Println(v18)
                                            var v19 uint32 = 1
                                            _ = v18
                                            _ = v19
                                        }
                                    case "hello":
                                        var v20 int8 = 127
                                        _ = v20
                                    default:
                                        var v21 string = ""
                                        _ = v21
                                    }
                                    _ = fi15
                                    _ = v16
                                }
                                var v22 uint16 = v2
                                _ = v22
                            } else {
                                var v23 bool = true
                                var v24 float32 = 0.1
                                v4 = v4
                                _ = v23
                                _ = v24
                            }
                        default:
                            v25 := v1[-652]
                            _ = v25
                        }
                        v26 := GenS3{F0: 240, F1: (v4 ^ (v4 - v4))}
                        _ = v26
                    }
                case "":
                    var v27 int64 = -1
                    _ = v27
                }
                switch "" {
                case "foo bar":
                    for fi28 := 0; fi28 < 11; fi28++ {
                        var v29 float32 = (3.14159 / 1e10)
                        for fi30 := 0; fi30 < 9; fi30++ {
                            v31 := GenS4{F0: v2}
                            _ = fi30
                            _ = v31
                        }
                        _ = fi28
                        _ = v29
                    }
                case "":
                    switch 32768 {
                    case 2147483647:
                        if false {
                            v3[2147483647] = -364
                            v4 = v4
                        } else {
                            for fi32 := 0; fi32 < 14; fi32++ {
                                if true {
                                    if true {
                                        var v33 bool = (false && ("" == "\n\t"))
                                        v12 -= v12
                                        v34 := []int{fi32, fi32}
                                        _ = v33
                                        _ = v34
                                    }
                                    v2 = v2
                                }
                                for fi35 := 0; fi35 < 7; fi35++ {
                                    if ((!(1.5 <= 3.14159)) && (!true)) {
                                        if false {
                                            v4 = v4
                                        }
                                    }
                                    v36 := []int{(fi32 | (fi32 & (fi32 >> ((uint(v2)) & 63)))), (fi35 + fi32), fi32}
                                    if (12 < 42) {
                                        var v37 uint = 216
                                        if true {
                                            var v38 float32 = -2.25
                                            var v39 uint = (^(v4))
                                            var v40 uint16 = v2
                                            _ = v38
                                            _ = v39
                                            _ = v40
                                        }
                                        fmt.Println(v4, fi32, fi32)
                                        _ = v37
                                    }
                                    _ = fi35
                                    _ = v36
                                }
                                _ = fi32
                            }
                        }
                    case 4294967295:
                        switch (("" + "\n\t") + "\n\t") {
                        case "a":
                            delete(v1, -2147483648)
                        }
                    default:
                        _ = len(v1)
                    }
                case "\n\t":
                    v41 := []int{127, int(v12)}
                    _ = v41
                }
                var v42 uint16 = (v2 % ((v2) | 1))
                _ = v42
            }
            v43 := GenS3{F0: 2, F1: 1}
            _ = v13
            _ = v43
        }
    }
    if (42 <= 18446744073709551615) {
        v12 = v12
        v4 += v4
        fmt.Println(v2, v12)
    } else {
        switch ("hello" + ("a" + "hello")) {
        case "":
            var v44 int64 = -1
            _ = v44
        case "hello":
            var v45 uint32 = 2147483648
            _ = v45
        }
        switch ("\n\t" + "a") {
        case "foo bar":
            if false {
                var v46 int = 4294967295
                for fi47 := 0; fi47 < 3; fi47++ {
                    v48 := []int{409, (-(fi47))}
                    _ = fi47
                    _ = v48
                }
                _ = v46
            }
        case "a":
            if false {
                fmt.Println(v4, v2)
                for fi49 := 0; fi49 < 20; fi49++ {
                    if true {
                        if ((v12 ^ (v12 & (^(v12)))) == (v12 | (v12 / ((v12) | 1)))) {
                            v4 += 618
                            var v50 int64 = 0
                            _ = v50
                        } else {
                            if (float64(1.5) >= 1.5) {
                                for fi51 := 0; fi51 < 7; fi51++ {
                                    switch 42 {
                                    case 42:
                                        fmt.Println(fi51, v4, v12)
                                    case 0:
                                        switch "\n\t" {
                                        case "":
                                            delete(v1, fi49)
                                        case "\n\t":
                                            fmt.Println(v2)
                                        default:
                                            if (true && false) {
                                                var v52 int8 = int8(v4)
                                                v52 -= (v52 * (v52 ^ v52))
                                                v2 += v2
                                                var v53 bool = true
                                                _ = v52
                                                _ = v53
                                            }
                                        }
                                    }
                                    switch (fi49 % (((fi51 | (-(fi51)))) | 1)) {
                                    case 4294967295:
                                        var v54 int64 = 188
                                        _ = v54
                                    case -410:
                                        if (!((65535 >= (v4 - (v4 * v4))) || true)) {
                                            switch (fi51 - (fi49 * (-(fi49)))) {
                                            case 32767:
                                                if false {
                                                    var v55 uint64 = 941
                                                    v56 := GenS5{F0: 2147483647, F1: false}
                                                    _ = v55
                                                    _ = v56
                                                } else {
                                                    var v57 string = ("foo bar" + "")
                                                    _ = v57
                                                }
                                                v58 := []int{(^(fi49)), fi49}
                                                _ = v58
                                            case 13:
                                                v12 = (v12 * 708)
                                            case 2147483647:
                                                switch fi49 {
                                                case -2147483648:
                                                    switch (("foo bar" + (("\n\t" + "") + "\n\t")) + "hello") {
                                                    case "foo bar":
                                                        v59 := []int{127, (fi51 ^ fi49), fi51, fi49}
                                                        _ = v59
                                                    case "hello":
                                                        if (!false) {
                                                            var v60 int32 = -1
                                                            _ = v60
                                                        }
                                                    case "":
                                                        if false {
                                                            v61 := GenS3{F0: 134, F1: (^(v4))}
                                                            var v62 uint64 = 4294967295
                                                            _ = v61
                                                            _ = v62
                                                        }
                                                    default:
                                                        var v63 bool = false
                                                        _ = v63
                                                    }
                                                    v12 = v12
                                                case 1:
                                                    var v64 bool = false
                                                    _ = v64
                                                case 2147483647:
                                                    v12 -= (-(v12))
                                                default:
                                                    delete(v1, (-(fi51)))
                                                }
                                            default:
                                                var v65 string = "\n\t"
                                                _ = v65
                                            }
                                            var v66 bool = true
                                            _ = v66
                                        }
                                    case 496:
                                        if false {
                                            delete(v1, (fi49 ^ (fi51 / (((fi49 >> ((65535) & 63))) | 1))))
                                            v12 = v12
                                            var v67 int8 = -14
                                            _ = v67
                                        } else {
                                            fmt.Println(v4)
                                            if true {
                                                v4 -= 463
                                            } else {
                                                v68 := GenS3{F0: 7, F1: v4}
                                                fmt.Println(v4, fi49, v12)
                                                _ = v68
                                            }
                                        }
                                    default:
                                        var v69 uint32 = 42
                                        _ = v69
                                    }
                                    _ = len(v3)
                                    _ = fi51
                                }
                                if (977 >= 4294967295) {
                                    if true {
                                        var v70 int16 = v12
                                        switch ("" + "a") {
                                        case "":
                                            switch "" {
                                            case "a":
                                                fmt.Println(v2, v70)
                                            case "hello":
                                                for fi71 := 0; fi71 < 15; fi71++ {
                                                    v12 -= v12
                                                    var v72 int64 = -9223372036854775808
                                                    _ = fi71
                                                    _ = v72
                                                }
                                            }
                                        }
                                        if (!false) {
                                            v73 := GenS6{F0: -1, F1: 2147483647, F2: 2}
                                            _ = v73
                                        } else {
                                            switch fi49 {
                                            case 2147483647:
                                                v74 := GenS7{F0: 0, F1: 1}
                                                _ = v74
                                            default:
                                                if false {
                                                    if false {
                                                        switch int(v2) {
                                                        case 1:
                                                            delete(v3, 2)
                                                        default:
                                                            v75 := GenS8{F0: -658, F1: (true || true), F2: 0}
                                                            _ = v75
                                                        }
                                                        var v76 int32 = 0
                                                        _ = v76
                                                    }
                                                    var v77 int32 = -2147483648
                                                    _ = v77
                                                } else {
                                                    if true {
                                                        for fi78 := 0; fi78 < 8; fi78++ {
                                                            if (!((!false) && (v12 == -938))) {
                                                                v79 := GenS5{F0: -1, F1: true}
                                                                v80 := GenS5{F0: fi49, F1: true}
                                                                _ = v79
                                                                _ = v80
                                                            } else {
                                                                v4 -= (^(v4))
                                                                if false {
                                                                    var v81 uint16 = 719
                                                                    _ = v81
                                                                } else {
                                                                    fmt.Println(fi49, v12, fi78)
                                                                    var v82 uint64 = 106
                                                                    var v83 int64 = 1
                                                                    v4 = v4
                                                                    v84 := []int{(fi49 / ((fi78) | 1)), fi78, (fi49 - 42)}
                                                                    _ = v82
                                                                    _ = v83
                                                                    _ = v84
                                                                }
                                                            }
                                                            fmt.Println(v2)
                                                            _ = fi78
                                                        }
                                                        delete(v1, (fi49 ^ (fi49 / (((fi49 ^ -197)) | 1))))
                                                    } else {
                                                        v70 += (^(v70))
                                                        v4 += 155
                                                    }
                                                    for fi85 := 0; fi85 < 7; fi85++ {
                                                        fmt.Println(fi49)
                                                        _ = fi85
                                                    }
                                                    v70 += v12
                                                }
                                            }
                                            v70 = (v70 + v12)
                                            if ((v12 << ((uint(v2)) & 63)) == v12) {
                                                fmt.Println(fi49, v4, v4)
                                                v86 := GenS9{F0: (v12 | v12), F1: "foo bar"}
                                                _ = v86
                                            } else {
                                                for fi87 := 0; fi87 < 14; fi87++ {
                                                    if true {
                                                        v4 = (v4 / ((42) | 1))
                                                        if true {
                                                            var v88 uint8 = 152
                                                            _ = v88
                                                        }
                                                        v89 := GenS10{F0: 2147483647, F1: v2, F2: (false && true)}
                                                        _ = v89
                                                    }
                                                    if false {
                                                        v2 = (^(v2))
                                                    }
                                                    v2 -= (^(v2))
                                                    v90 := []int{fi87, fi87, 32768, fi49}
                                                    _ = fi87
                                                    _ = v90
                                                }
                                                v91 := []int{-676, fi49, (fi49 * (^(fi49))), (fi49 / ((128) | 1))}
                                                for fi92 := 0; fi92 < 6; fi92++ {
                                                    if (v2 != 2) {
                                                        var v93 int16 = (v70 >> (((^(v4))) & 63))
                                                        _ = v93
                                                    } else {
                                                        v94 := GenS5{F0: fi92, F1: (-2147483648 > 127)}
                                                        v95 := GenS11{F0: -954, F1: 1000, F2: 4294967295}
                                                        _ = v94
                                                        _ = v95
                                                    }
                                                    _ = fi92
                                                }
                                                _ = v91
                                            }
                                        }
                                        _ = v70
                                    } else {
                                        if false {
                                            for fi96 := 0; fi96 < 16; fi96++ {
                                                fmt.Println(fi49, v12)
                                                fmt.Println(v12, v4)
                                                v2 -= 2
                                                _ = fi96
                                            }
                                            var v97 uint8 = 127
                                            _ = v97
                                        }
                                    }
                                    if (!false) {
                                        v4 = (v4 % (((v4 + v4)) | 1))
                                    } else {
                                        v98 := GenS6{F0: 51, F1: 677, F2: 0}
                                        _ = v98
                                    }
                                    v2 = 1
                                }
                                v99 := v1[(fi49 + fi49)]
                                _ = v99
                            } else {
                                v4 = (v4 * (v4 >> ((v4) & 63)))
                                v4 += 65535
                            }
                            var v100 uint64 = 4294967295
                            _ = v100
                        }
                    }
                    v2 = 565
                    _ = fi49
                }
            }
        default:
            var v101 string = (("hello" + "a") + "a")
            _ = v101
        }
        v102 := GenS2{F0: 0}
        _ = v102
    }
    for fi103 := 0; fi103 < 13; fi103++ {
        var v104 int8 = 71
        for fi105 := 0; fi105 < 9; fi105++ {
            v106 := GenS2{F0: 126}
            v107 := GenS2{F0: uint8(v4)}
            _ = fi105
            _ = v106
            _ = v107
        }
        switch (fi103 / ((fi103) | 1)) {
        case -9223372036854775808:
            if false {
                v104 = (v104 << ((v4) & 63))
                var v108 uint16 = 2
                _ = v108
            } else {
                var v109 uint32 = 2147483647
                _ = v109
            }
        case 32768:
            v110 := GenS12{F0: 42}
            _ = v110
        }
        _ = fi103
        _ = v104
    }
    for fi111 := 0; fi111 < 1; fi111++ {
        if true {
            var v112 uint16 = 32767
            v112 += 306
            v4 -= v4
            _ = v112
        }
        v4 = v4
        switch "\n\t" {
        case "hello":
            if false {
                v113 := []int{}
                switch ("a" + ("a" + "foo bar")) {
                case "hello":
                    if (-359 < 1) {
                        if false {
                            fmt.Println(fi111, v2, v2)
                        } else {
                            if (3.14159 != -2.25) {
                                if (105 <= 0) {
                                    if (!false) {
                                        if (0 < 0) {
                                            for fi114 := 0; fi114 < 5; fi114++ {
                                                var v115 uint16 = (^(v2))
                                                _ = fi114
                                                _ = v115
                                            }
                                        } else {
                                            for fi116 := 0; fi116 < 11; fi116++ {
                                                if true {
                                                    if len(v113) > 0 {
                                                        v113[0] = -9223372036854775808
                                                    }
                                                }
                                                switch (fi116 >> ((42) & 63)) {
                                                case -697:
                                                    v117 := GenS13{F0: 126, F1: 1}
                                                    _ = len(v113)
                                                    _ = v117
                                                case 9223372036854775807:
                                                    var v118 int16 = v12
                                                    _ = v118
                                                case 2147483647:
                                                    var v119 int8 = 100
                                                    _ = v119
                                                default:
                                                    fmt.Println(v12, fi111, fi116)
                                                    v120 := GenS10{F0: -9, F1: (^(v2)), F2: (!false)}
                                                    _ = v120
                                                }
                                                v12 -= (v12 + 0)
                                                _ = fi116
                                            }
                                        }
                                        var v121 float64 = -2.25
                                        fmt.Println(fi111, v12)
                                        _ = v121
                                    }
                                    v122 := GenS8{F0: -2147483648, F1: false, F2: -9223372036854775808}
                                    for fi123 := 0; fi123 < 11; fi123++ {
                                        v2 = (v2 * v2)
                                        if (!false) {
                                            _ = len(v1)
                                            var v124 uint32 = uint32(v4)
                                            fmt.Println(fi111)
                                            switch ("" + "foo bar") {
                                            case "hello":
                                                var v125 uint8 = 97
                                                _ = v125
                                            case "":
                                                var v126 uint64 = 42
                                                _ = v126
                                            }
                                            var v127 uint16 = 561
                                            _ = v124
                                            _ = v127
                                        } else {
                                            if (false || true) {
                                                var v128 int32 = 42
                                                v129 := v3[0]
                                                _ = v128
                                                _ = v129
                                            }
                                        }
                                        var v130 float32 = (0.0 / float32(((((0.1 * -2.25) - -2.25) + 3.14159) + 1.5)))
                                        _ = fi123
                                        _ = v130
                                    }
                                    _ = v122
                                } else {
                                    v2 += v2
                                }
                            } else {
                                v3[1] = fi111
                                v3[65535] = fi111
                                var v131 uint8 = 237
                                _ = v131
                            }
                            switch "a" {
                            case "foo bar":
                                if len(v113) >= 2 {
                                    v113 = v113[1:len(v113)]
                                }
                            default:
                                v4 -= (v4 - v4)
                            }
                            var v132 float64 = -2.25
                            _ = v132
                        }
                    } else {
                        var v133 float32 = (1.5 / 0.1)
                        if (25 == int8(v12)) {
                            fmt.Println(v12)
                        }
                        for fi134 := 0; fi134 < 20; fi134++ {
                            var v135 bool = ((v4 % ((v4) | 1)) < v4)
                            _ = fi134
                            _ = v135
                        }
                        _ = v133
                    }
                case "foo bar":
                    for fi136 := 0; fi136 < 12; fi136++ {
                        var v137 uint32 = 2
                        _ = fi136
                        _ = v137
                    }
                case "a":
                    if false {
                        for fi138 := 0; fi138 < 5; fi138++ {
                            v139 := GenS8{F0: genFn1(-2147483648), F1: false, F2: 1}
                            var v140 float32 = 3.14159
                            _ = fi138
                            _ = v139
                            _ = v140
                        }
                        for fi141 := 0; fi141 < 9; fi141++ {
                            v142 := GenS13{F0: 0, F1: v12}
                            var v143 int64 = 1
                            var v144 int64 = 1
                            _ = fi141
                            _ = v142
                            _ = v143
                            _ = v144
                        }
                        _ = len(v113)
                    } else {
                        if len(v113) > 0 {
                            v113[0] = -1
                        }
                        for fi145 := 0; fi145 < 13; fi145++ {
                            v113 = append(v113, -1)
                            v146 := GenS12{F0: -940}
                            v4 -= (v4 - (^(v4)))
                            _ = fi145
                            _ = v146
                        }
                    }
                default:
                    var v147 int = fi111
                    _ = v147
                }
                _ = v113
            }
        case "":
            if false {
                switch -350 {
                case -2147483648:
                    var v148 string = ""
                    _ = v148
                case -975:
                    switch "" {
                    case "\n\t":
                        var v149 bool = ((-2147483648 <= 0) || ((-9223372036854775808 != 580) && false))
                        _ = v149
                    case "a":
                        if true {
                            for fi150 := 0; fi150 < 3; fi150++ {
                                if true {
                                    delete(v3, 2147483648)
                                }
                                fmt.Println(fi150)
                                _ = fi150
                            }
                            if (-2.25 <= 0.1) {
                                if ("hello" > ((("foo bar" + "") + "\n\t") + "")) {
                                    switch fi111 {
                                    case 9223372036854775807:
                                        var v151 int = (fi111 - (fi111 | fi111))
                                        _ = v151
                                    case 4294967295:
                                        var v152 uint = (^(v4))
                                        _ = v152
                                    default:
                                        for fi153 := 0; fi153 < 9; fi153++ {
                                            v154 := GenS6{F0: -66, F1: -1, F2: uint8(v4)}
                                            delete(v1, -609)
                                            v2 = (v2 * v2)
                                            _ = fi153
                                            _ = v154
                                        }
                                    }
                                } else {
                                    v155 := GenS14{F0: true, F1: 0, F2: genFn1(-1)}
                                    v156 := GenS15{F0: (v2 >> ((v4) & 63)), F1: v12}
                                    if ("a" < "hello") {
                                        v2 += v2
                                        var v157 int = (fi111 >> (((v4 / (((^(v4))) | 1))) & 63))
                                        v12 += v12
                                        _ = v157
                                    } else {
                                        v12 = v12
                                        v4 -= (v4 - v4)
                                    }
                                    _ = v155
                                    _ = v156
                                }
                            } else {
                                _ = len(v1)
                                switch fi111 {
                                case 4294967295:
                                    v1[-2147483648] = "hello"
                                case 501:
                                    if false {
                                        switch "a" {
                                        case "":
                                            if false {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v158 int32 = -1
                                                    _ = v158
                                                case "":
                                                    if true {
                                                        v4 += (v4 & v4)
                                                    }
                                                case "a":
                                                    switch (fi111 | fi111) {
                                                    case 336:
                                                        v4 = 4294967295
                                                    case 570:
                                                        var v159 int32 = -312
                                                        _ = v159
                                                    case 32768:
                                                        switch "hello" {
                                                        case "foo bar":
                                                            var v160 int16 = v12
                                                            _ = v160
                                                        case "\n\t":
                                                            v161 := GenS4{F0: 0}
                                                            _ = v161
                                                        default:
                                                            var v162 uint32 = 4294967295
                                                            _ = v162
                                                        }
                                                    default:
                                                        for fi163 := 0; fi163 < 3; fi163++ {
                                                            v164 := v1[fi111]
                                                            v164 = v164
                                                            _ = fi163
                                                            _ = v164
                                                        }
                                                    }
                                                }
                                            }
                                        case "a":
                                            switch "a" {
                                            case "a":
                                                v165 := GenS16{F0: "foo bar", F1: 25}
                                                _ = v165
                                            case "hello":
                                                for fi166 := 0; fi166 < 4; fi166++ {
                                                    v12 -= 0
                                                    _ = fi166
                                                }
                                            default:
                                                fmt.Println(fi111)
                                            }
                                        }
                                    } else {
                                        switch fi111 {
                                        case 42:
                                            v12 += -451
                                        }
                                        fmt.Println(v2)
                                    }
                                case 42:
                                    var v167 uint = v4
                                    _ = v167
                                }
                            }
                        } else {
                            delete(v1, fi111)
                            v12 -= v12
                        }
                    default:
                        v4 -= (v4 & v4)
                    }
                }
                fmt.Println(v12, v2)
            }
        case "\n\t":
            var v168 bool = (!(!(-916 < 411)))
            _ = v168
        }
        _ = fi111
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
