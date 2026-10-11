package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 string; F1 uint64; F2 uint8 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint32; F1 uint8 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 bool; F1 uint8 }
type GenS15 struct { F0 uint8; F1 uint8; F2 bool }
type GenS16 struct { F0 uint64; F1 int8 }
type GenS17 struct { F0 int16; F1 uint8 }
type GenS18 struct { F0 int16; F1 int16 }
func genFn0(p0 int32) int8 {
    for fi5 := 0; fi5 < 13; fi5++ {
        fmt.Println(p0, fi5)
        p0 = -813
        p0 -= p0
        _ = fi5
    }
    p0 = 1
    fmt.Println(p0)
    for fi6 := 0; fi6 < 7; fi6++ {
        p0 = p0
        if (0 >= 42) {
            v7 := GenS1{F0: 4294967295, F1: 206, F2: fi6}
            for fi8 := 0; fi8 < 16; fi8++ {
                v9 := GenS1{F0: 1, F1: 65535, F2: (fi8 & (fi6 >> ((18446744073709551615) & 63)))}
                var v10 int32 = -2147483648
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = v7
        } else {
            for fi11 := 0; fi11 < 8; fi11++ {
                if ((p0 >> ((265) & 63)) <= p0) {
                    p0 = (p0 >> ((1) & 63))
                    if false {
                        var v12 uint8 = 127
                        var v13 uint = 65535
                        _ = v12
                        _ = v13
                    } else {
                        fmt.Println(p0, fi11, fi11)
                        var v14 uint32 = 4294967295
                        _ = v14
                    }
                    v15 := []int{4294967295, 1}
                    _ = v15
                } else {
                    v16 := GenS1{F0: 65535, F1: 42, F2: (^(fi11))}
                    v17 := []int{(fi6 - fi6), fi11, fi6}
                    p0 = p0
                    _ = v16
                    _ = v17
                }
                var v18 int8 = 52
                _ = fi11
                _ = v18
            }
            for fi19 := 0; fi19 < 3; fi19++ {
                if (42 >= 108) {
                    p0 += p0
                    var v20 string = ""
                    v21 := GenS2{F0: -1}
                    p0 += 2147483647
                    v20 = ((v20 + v20) + ((v20 + v20) + ("\n\t" + ("a" + v20))))
                    _ = v20
                    _ = v21
                } else {
                    fmt.Println(fi6)
                }
                var v22 uint = 42
                var v23 int = fi6
                v24 := GenS3{F0: 4294967295}
                var v25 uint32 = 42
                p0 = (p0 / ((p0) | 1))
                if (true && true) {
                    v26 := []int{fi6}
                    v23 = (fi19 / (((v23 * v23)) | 1))
                    p0 = 0
                    if false {
                        var v27 string = "a"
                        _ = v27
                    } else {
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                    }
                    _ = v26
                } else {
                    p0 = (p0 + (p0 >> ((v22) & 63)))
                    if (0 != -737) {
                        var v28 uint16 = 527
                        var v29 float32 = -2.25
                        var v30 string = "a"
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    if (!false) {
                        v31 := []int{-2147483648}
                        if len(v31) >= 2 {
                            v31 = v31[1:len(v31)]
                        }
                        fmt.Println(p0, v25, v23)
                        _ = v31
                    } else {
                        v32 := []int{(^(fi6))}
                        v23 += -188
                        var v33 int32 = 1
                        _ = v32
                        _ = v33
                    }
                }
                v22 += (v22 / (((v22 + 255)) | 1))
                _ = fi19
                _ = v22
                _ = v23
                _ = v24
                _ = v25
            }
            switch fi6 {
            case 127:
                if ((fi6 & 9223372036854775807) <= 32768) {
                    switch fi6 {
                    case -760:
                        switch "" {
                        case "hello":
                            for fi34 := 0; fi34 < 17; fi34++ {
                                v35 := GenS4{F0: 722}
                                _ = fi34
                                _ = v35
                            }
                        case "a":
                            if true {
                                var v36 uint32 = 65535
                                for fi37 := 0; fi37 < 10; fi37++ {
                                    v38 := []int{42, (fi37 ^ 0), (fi6 >> ((1) & 63))}
                                    v39 := GenS1{F0: 573, F1: v36, F2: 0}
                                    if false {
                                        v40 := GenS4{F0: 589}
                                        v41 := GenS5{F0: 1}
                                        v36 = (v36 | (v36 | v36))
                                        var v42 float64 = 1.5
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                    }
                                    _ = fi37
                                    _ = v38
                                    _ = v39
                                }
                                v36 = (^(v36))
                                _ = v36
                            }
                        case "":
                            fmt.Println(fi6, p0)
                        default:
                            if (float64((1.5 + float32(3.14159))) < (-2.25 - (-2.25 - float64(3.14159)))) {
                                var v43 uint64 = 49
                                _ = v43
                            }
                        }
                    default:
                        var v44 int64 = -1
                        _ = v44
                    }
                    for fi45 := 0; fi45 < 20; fi45++ {
                        if true {
                            p0 -= (p0 / ((p0) | 1))
                            var v46 uint16 = 1
                            _ = v46
                        } else {
                            var v47 int16 = 1000
                            if false {
                                switch "hello" {
                                case "\n\t":
                                    var v48 int = 0
                                    _ = v48
                                case "hello":
                                    v49 := GenS3{F0: 65535}
                                    _ = v49
                                }
                                fmt.Println(fi6, fi6)
                                v47 = v47
                            }
                            v50 := GenS6{F0: ("hello" + "")}
                            _ = v47
                            _ = v50
                        }
                        if true {
                            var v51 uint8 = 22
                            switch 0 {
                            case 482:
                                p0 += p0
                            }
                            switch (fi45 + 32768) {
                            case -1:
                                v52 := GenS4{F0: uint64(v51)}
                                _ = v52
                            case -730:
                                if true {
                                    v53 := GenS2{F0: -957}
                                    var v54 uint16 = 32767
                                    if (!(!("" < "hello"))) {
                                        v55 := GenS6{F0: ("hello" + "")}
                                        switch 849 {
                                        case -332:
                                            v51 -= v51
                                        }
                                        v56 := []int{fi45}
                                        _ = v55
                                        _ = v56
                                    }
                                    _ = v53
                                    _ = v54
                                } else {
                                    fmt.Println(p0, fi45, fi6)
                                }
                            default:
                                switch (^(fi45)) {
                                case -416:
                                    v57 := GenS7{F0: 2}
                                    _ = v57
                                case 2147483647:
                                    switch "a" {
                                    case "foo bar":
                                        var v58 float64 = (-2.25 * 0.0)
                                        _ = v58
                                    case "hello":
                                        var v59 uint16 = 32767
                                        _ = v59
                                    }
                                }
                            }
                            v60 := []int{}
                            _ = len(v60)
                            _ = v51
                            _ = v60
                        }
                        _ = fi45
                    }
                    v61 := GenS7{F0: 1}
                    _ = v61
                }
            case -65:
                if true {
                    p0 = p0
                    switch fi6 {
                    case 9223372036854775807:
                        var v62 int16 = 42
                        _ = v62
                    case 1:
                        v63 := []int{fi6, fi6, fi6}
                        _ = v63
                    case -646:
                        var v64 int64 = 933
                        _ = v64
                    }
                }
            default:
                fmt.Println(p0, fi6)
            }
        }
        _ = fi6
    }
    _ = p0
    return 127
}

func main() {
    var v1 int = 41
    v2 := []string{}
    var v3 int32 = 0
    v4 := make(map[uint32]string)
    switch -9223372036854775808 {
    case 449:
        v65 := GenS1{F0: 0, F1: 2, F2: 128}
        _ = v65
    }
    var v66 string = "a"
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    v1 = v1
    v3 = v3
    if true {
        fmt.Println(v3, v1, v3)
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v67 int8 = -60
        if (false || (!false)) {
            var v68 uint = 4294967295
            switch v1 {
            case -119:
                if false {
                    if (!(!false)) {
                        _ = len(v2)
                        for fi69 := 0; fi69 < 20; fi69++ {
                            var v70 float64 = -2.25
                            _ = fi69
                            _ = v70
                        }
                    } else {
                        v71 := GenS7{F0: 415}
                        for fi72 := 0; fi72 < 11; fi72++ {
                            v73 := GenS8{F0: 738}
                            for fi74 := 0; fi74 < 12; fi74++ {
                                switch (-(fi74)) {
                                case 42:
                                    switch (fi74 - (fi74 & (v1 / ((-2147483648) | 1)))) {
                                    case -9223372036854775808:
                                        if (false || true) {
                                            v75 := GenS7{F0: 65535}
                                            if ((((v1 <= 128) && true) && true) || (!false)) {
                                                var v76 float64 = float64(float32(float64((float32(1.5) * 3.14159))))
                                                v2 = append(v2, v66)
                                                v76 += v76
                                                _ = v76
                                            } else {
                                                if false {
                                                    v66 = ((v66 + ("foo bar" + v66)) + "foo bar")
                                                }
                                                var v77 uint16 = 43
                                                _ = v77
                                            }
                                            switch v66 {
                                            case "foo bar":
                                                v1 = (fi74 - (v1 & (fi72 ^ 9223372036854775807)))
                                            case "\n\t":
                                                var v78 float32 = (1.5 / (1e10 + 1.5))
                                                _ = v78
                                            }
                                            _ = v75
                                        } else {
                                            delete(v4, 627)
                                            var v79 int64 = 53
                                            _ = v79
                                        }
                                    case 127:
                                        if true {
                                            fmt.Println(v66)
                                            v68 = v68
                                        } else {
                                            v80 := v4[4294967295]
                                            _ = v80
                                        }
                                    default:
                                        v67 = 42
                                    }
                                case 1:
                                    var v81 uint16 = 255
                                    _ = v81
                                default:
                                    var v82 int16 = 42
                                    var v83 int32 = v3
                                    if true {
                                        _ = len(v2)
                                        var v84 string = ""
                                        v68 -= (v68 >> ((65535) & 63))
                                        _ = v84
                                    }
                                    _ = v82
                                    _ = v83
                                }
                                v85 := GenS9{F0: 53}
                                _ = fi74
                                _ = v85
                            }
                            fmt.Println(v66, fi72)
                            _ = fi72
                            _ = v73
                        }
                        v86 := GenS10{F0: ("\n\t" + v66), F1: 4294967295, F2: 81}
                        _ = v71
                        _ = v86
                    }
                    var v87 uint32 = 65535
                    _ = v87
                } else {
                    for fi88 := 0; fi88 < 20; fi88++ {
                        var v89 uint32 = 2147483647
                        if true {
                            v68 = (v68 - (v68 << (((v68 ^ v68)) & 63)))
                            var v90 string = v66
                            _ = v90
                        }
                        var v91 float64 = 1e10
                        _ = fi88
                        _ = v89
                        _ = v91
                    }
                    v92 := GenS4{F0: 4294967295}
                    _ = v92
                }
            case 996:
                switch int(v3) {
                case -696:
                    var v93 uint64 = 2
                    _ = v93
                case 127:
                    v68 -= v68
                default:
                    _ = len(v4)
                }
            default:
                v3 += (v3 ^ v3)
            }
            if ((v3 % ((2147483647) | 1)) != (v3 | v3)) {
                v68 = v68
                for fi94 := 0; fi94 < 11; fi94++ {
                    v68 -= (v68 | (^(v68)))
                    v95 := GenS11{F0: 1}
                    var v96 uint = v68
                    _ = fi94
                    _ = v95
                    _ = v96
                }
            } else {
                fmt.Println(v3)
                delete(v4, 2)
            }
            _ = v68
        }
        _ = v67
    }
    v1 = 0
    v66 = (v66 + v66)
    v97 := GenS3{F0: 65535}
    if true {
        switch v1 {
        case 32768:
            fmt.Println(v1, v3)
        case 9223372036854775807:
            switch v66 {
            case "a":
                fmt.Println(v3)
            default:
                v3 -= (v3 + (v3 << ((0) & 63)))
            }
        }
        v66 = (v66 + v66)
        var v98 uint8 = 2
        _ = v98
    } else {
        for fi99 := 0; fi99 < 10; fi99++ {
            var v100 int8 = 1
            v2 = append(v2, (v66 + ("" + "")))
            fmt.Println(v100)
            _ = fi99
            _ = v100
        }
    }
    v101 := GenS12{F0: 2147483647, F1: 125}
    if false {
        switch v1 {
        case 128:
            v66 = v66
        case -9223372036854775808:
            var v102 int = (v1 - (v1 + (v1 | (v1 - v1))))
            _ = v102
        }
        var v103 int32 = (v3 | -1)
        var v104 uint32 = 21
        _ = v103
        _ = v104
    } else {
        v105 := v4[65535]
        if true {
            var v106 bool = (!(!(!false)))
            if false {
                v1 = 9223372036854775807
                fmt.Println(v106, v106)
                var v107 float32 = -2.25
                _ = v107
            } else {
                v108 := GenS9{F0: 42}
                _ = v108
            }
            _ = v106
        }
        if (v3 > v3) {
            v109 := GenS13{F0: 992}
            v66 = (v66 + (v105 + (v66 + v105)))
            _ = v109
        }
        _ = v105
    }
    fmt.Println(v3)
    if true {
        var v110 uint = 4294967295
        if true {
            v111 := GenS6{F0: "\n\t"}
            var v112 string = v66
            v4[898] = "a"
            _ = v111
            _ = v112
        }
        for fi113 := 0; fi113 < 14; fi113++ {
            switch (v1 - 32767) {
            case 152:
                v114 := GenS2{F0: (v3 >> ((v110) & 63))}
                _ = v114
            case 284:
                switch v66 {
                case "foo bar":
                    switch (v66 + "") {
                    case "foo bar":
                        v3 += -1
                    case "a":
                        var v115 int32 = (v3 | 42)
                        _ = v115
                    default:
                        v110 += (v110 + (v110 % (((v110 * v110)) | 1)))
                    }
                case "a":
                    fmt.Println(fi113, v1)
                default:
                    for fi116 := 0; fi116 < 18; fi116++ {
                        v117 := GenS14{F0: true, F1: 220}
                        _ = fi116
                        _ = v117
                    }
                }
            case 4294967295:
                var v118 int8 = 42
                _ = v118
            }
            fmt.Println(v1, v3)
            if (true && ((fi113 <= -273) && ((-58 >= 24) && false))) {
                for fi119 := 0; fi119 < 19; fi119++ {
                    v110 += (v110 + (v110 | (v110 / (((v110 & v110)) | 1))))
                    v3 += v3
                    _ = fi119
                }
                v66 = ((v66 + "hello") + ("foo bar" + (v66 + ("foo bar" + ("foo bar" + "foo bar")))))
                switch v66 {
                case "\n\t":
                    if (true || (!true)) {
                        v110 = (v110 - 204)
                        var v120 float32 = 1.5
                        _ = v120
                    }
                default:
                    if true {
                        switch (fi113 * (v1 - (fi113 + -9223372036854775808))) {
                        case 128:
                            _ = len(v4)
                        case 0:
                            fmt.Println(v110, v110, v110)
                        default:
                            var v121 int = v1
                            _ = v121
                        }
                        var v122 float64 = float64(3.14159)
                        var v123 int16 = 0
                        _ = v122
                        _ = v123
                    }
                }
            }
            _ = fi113
        }
        _ = v110
    }
    for fi124 := 0; fi124 < 3; fi124++ {
        var v125 float32 = 0.1
        if true {
            switch ("a" + "a") {
            case "\n\t":
                if (255 <= 42) {
                    if false {
                        v126 := GenS13{F0: 414}
                        switch "a" {
                        case "hello":
                            v3 += 765
                        case "foo bar":
                            if (false && (40 >= 941)) {
                                v127 := GenS11{F0: -1}
                                v128 := GenS15{F0: 42, F1: 1, F2: (1 < 0)}
                                _ = v127
                                _ = v128
                            } else {
                                v129 := GenS16{F0: 466, F1: 22}
                                v2 = append(v2, "a")
                                _ = v129
                            }
                        }
                        fmt.Println(v125)
                        _ = v126
                    }
                } else {
                    v4[1] = v66
                    fmt.Println(v125)
                    if false {
                        var v130 uint = 2
                        if false {
                            fmt.Println(v125, v130, fi124)
                        }
                        switch v1 {
                        case 42:
                            if true {
                                if true {
                                    v125 = 1e10
                                    if (-1 > 42) {
                                        _ = len(v2)
                                        v3 -= (v3 + v3)
                                        v66 = v66
                                    } else {
                                        v131 := GenS6{F0: (v66 + "\n\t")}
                                        _ = v131
                                    }
                                    for fi132 := 0; fi132 < 19; fi132++ {
                                        if (!(!(!false))) {
                                            if true {
                                                var v133 int = v1
                                                _ = v133
                                            } else {
                                                var v134 uint32 = 42
                                                _ = v134
                                            }
                                            fmt.Println(v1, v125, v1)
                                            if (42 != 351) {
                                                var v135 float32 = v125
                                                v66 = v66
                                                v136 := v4[383]
                                                _ = v135
                                                _ = v136
                                            } else {
                                                if true {
                                                    var v137 bool = (true && (!(!(!(true || (4294967295 >= 347))))))
                                                    var v138 uint16 = 681
                                                    _ = v137
                                                    _ = v138
                                                }
                                                v1 = v1
                                            }
                                        }
                                        v125 -= (v125 - float32((0.0 - (1e10 + -2.25))))
                                        v3 += (v3 >> (((^(v130))) & 63))
                                        var v139 string = v66
                                        fmt.Println(v130)
                                        _ = fi132
                                        _ = v139
                                    }
                                }
                                var v140 int16 = -32768
                                _ = v140
                            } else {
                                var v141 int64 = 9223372036854775807
                                fmt.Println(v141)
                                _ = v141
                            }
                        case -1:
                            switch (fi124 % (((v1 + 9223372036854775807)) | 1)) {
                            case 128:
                                var v142 int32 = -291
                                _ = v142
                            case -151:
                                var v143 int8 = 42
                                _ = v143
                            }
                        default:
                            for fi144 := 0; fi144 < 12; fi144++ {
                                switch fi124 {
                                case 32767:
                                    if true {
                                        v1 -= v1
                                        if true {
                                            var v145 int64 = -9223372036854775808
                                            v4[4294967295] = v66
                                            v1 = v1
                                            v145 += (v145 & v145)
                                            _ = v145
                                        }
                                    } else {
                                        v125 += v125
                                        if true {
                                            v130 += 1
                                        }
                                        v146 := GenS5{F0: -128}
                                        _ = v146
                                    }
                                }
                                var v147 uint64 = 108
                                if (-9223372036854775808 >= 832) {
                                    var v148 uint32 = 2147483647
                                    _ = v148
                                }
                                _ = fi144
                                _ = v147
                            }
                        }
                        _ = v130
                    }
                }
            case "":
                switch "foo bar" {
                case "a":
                    var v149 int8 = 34
                    _ = v149
                case "foo bar":
                    for fi150 := 0; fi150 < 13; fi150++ {
                        fmt.Println(v1, v3, fi124)
                        if (false || (2 <= 1)) {
                            var v151 int64 = 1
                            var v152 int = -772
                            if true {
                                v153 := GenS5{F0: 0}
                                _ = v153
                            } else {
                                var v154 int8 = 42
                                v155 := GenS17{F0: 32767, F1: 1}
                                if true {
                                    v151 = v151
                                    var v156 int16 = 1000
                                    v157 := v4[42]
                                    _ = v156
                                    _ = v157
                                } else {
                                    fmt.Println(fi150, fi124, fi150)
                                }
                                _ = v154
                                _ = v155
                            }
                            _ = v151
                            _ = v152
                        } else {
                            fmt.Println(v125, v1, v3)
                            var v158 uint = 303
                            _ = v158
                        }
                        _ = len(v4)
                        _ = len(v2)
                        _ = fi150
                    }
                }
            default:
                for fi159 := 0; fi159 < 1; fi159++ {
                    v125 = 1.5
                    var v160 uint16 = 0
                    _ = fi159
                    _ = v160
                }
            }
            _ = len(v4)
            v66 = (v66 + "foo bar")
        }
        _ = fi124
        _ = v125
    }
    switch (v1 * 337) {
    case 9223372036854775807:
        v161 := GenS18{F0: 42, F1: -1}
        _ = v161
    }
    v1 += v1
    var v162 uint16 = 202
    fmt.Println("v1:", v1)
    fmt.Println("v162:", v162)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v66:", v66)
    _ = v1
    _ = v101
    _ = v162
    _ = v2
    _ = v3
    _ = v4
    _ = v66
    _ = v97
}
