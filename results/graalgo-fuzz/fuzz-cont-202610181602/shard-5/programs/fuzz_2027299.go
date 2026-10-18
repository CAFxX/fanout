package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 int64; F2 int }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 int64; F1 int8; F2 int16 }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 uint16; F1 int8; F2 string }
type GenS8 struct { F0 bool; F1 int32 }
type GenS9 struct { F0 uint16; F1 string; F2 uint }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS13 struct { F0 uint64; F1 uint }
type GenS14 struct { F0 int8; F1 int }
type GenS15 struct { F0 uint8; F1 int; F2 uint }
type GenS16 struct { F0 uint32; F1 string }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 uint; F1 int32; F2 int8 }
func genFn0(p0 int16, p1 int, p2 bool) uint8 {
    var v7 string = ("\n\t" + "foo bar")
    v8 := GenS1{F0: (p1 * 127)}
    fmt.Println(p2, p1, p0)
    switch "\n\t" {
    case "":
        var v9 uint = 65535
        _ = v9
    default:
        if p2 {
            p0 += (p0 / ((p0) | 1))
            v10 := GenS2{F0: (^(p0))}
            _ = v10
        }
    }
    var v11 float64 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint8, p1 uint32, p2 uint64) uint32 {
    for fi12 := 0; fi12 < 3; fi12++ {
        fmt.Println(p0, fi12, fi12)
        _ = fi12
    }
    var v13 int16 = 32767
    switch 442 {
    case -783:
        if (!false) {
            p2 = p2
            if (!true) {
                var v14 int64 = 201
                var v15 int64 = v14
                if false {
                    for fi16 := 0; fi16 < 12; fi16++ {
                        var v17 float32 = 3.14159
                        for fi18 := 0; fi18 < 13; fi18++ {
                            fmt.Println(p0, p2)
                            var v19 uint64 = 0
                            switch "hello" {
                            case "hello":
                                if false {
                                    var v20 int16 = v13
                                    v21 := []int{(-(fi16))}
                                    var v22 int8 = -1
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                }
                            case "foo bar":
                                var v23 uint = 1
                                _ = v23
                            case "\n\t":
                                var v24 int32 = 1
                                var v25 uint64 = v19
                                var v26 uint = 18446744073709551615
                                fmt.Println(v24, fi18, p2)
                                _ = v24
                                _ = v25
                                _ = v26
                            default:
                                v27 := GenS1{F0: (^(fi16))}
                                _ = v27
                            }
                            _ = fi18
                            _ = v19
                        }
                        _ = fi16
                        _ = v17
                    }
                }
                _ = v14
                _ = v15
            } else {
                v28 := []int{42, 128, -2147483648}
                var v29 uint16 = uint16(p1)
                _ = len(v28)
                _ = v28
                _ = v29
            }
        }
    case 4294967295:
        switch 32768 {
        case 107:
            v30 := []int{-2147483648, -833, 42, 584}
            _ = v30
        case 32768:
            p1 += uint32(p2)
        default:
            for fi31 := 0; fi31 < 16; fi31++ {
                v32 := GenS2{F0: -1}
                if true {
                    switch "" {
                    case "":
                        v33 := GenS1{F0: fi31}
                        _ = v33
                    case "a":
                        var v34 int64 = 1
                        _ = v34
                    default:
                        switch "foo bar" {
                        case "\n\t":
                            for fi35 := 0; fi35 < 1; fi35++ {
                                if (!((^(p2)) >= 2)) {
                                    v13 = 0
                                    p2 = uint64(p0)
                                    fmt.Println(fi31, v13)
                                }
                                var v36 float32 = 0.0
                                if (v36 < 0.1) {
                                    var v37 int = 1
                                    _ = v37
                                } else {
                                    fmt.Println(p1, v13)
                                    v38 := []int{fi35}
                                    _ = v38
                                }
                                _ = fi35
                                _ = v36
                            }
                        default:
                            for fi39 := 0; fi39 < 3; fi39++ {
                                v40 := GenS2{F0: v13}
                                v41 := GenS1{F0: int(v13)}
                                v42 := []int{fi31, (fi31 * 9223372036854775807), fi39, fi31}
                                v13 -= (v13 | v13)
                                _ = fi39
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                        }
                    }
                } else {
                    if (!true) {
                        if (1 > 0) {
                            if (!(2147483647 == -1)) {
                                var v43 bool = true
                                if false {
                                    var v44 uint16 = 381
                                    _ = v44
                                } else {
                                    if ((18446744073709551615 > 42) && ((v43 || v43) && (false && v43))) {
                                        p2 = 42
                                        switch "hello" {
                                        case "foo bar":
                                            switch fi31 {
                                            case 4294967295:
                                                if (65535 > 123) {
                                                    fmt.Println(p0)
                                                }
                                            case 32768:
                                                var v45 string = "a"
                                                _ = v45
                                            }
                                        case "hello":
                                            for fi46 := 0; fi46 < 2; fi46++ {
                                                if (!(false || false)) {
                                                    var v47 float64 = (-2.25 * 1e10)
                                                    fmt.Println(fi31)
                                                    _ = v47
                                                } else {
                                                    v48 := []int{fi31, int(p2)}
                                                    _ = v48
                                                }
                                                _ = fi46
                                            }
                                        case "":
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v49 int64 = 344
                                                _ = v49
                                            case "hello":
                                                if (((1e10 == 1.5) && (-689 == -9223372036854775808)) && (p1 < (^(p1)))) {
                                                    v50 := []int{1}
                                                    for fi51 := 0; fi51 < 9; fi51++ {
                                                        var v52 int64 = -1
                                                        _ = fi51
                                                        _ = v52
                                                    }
                                                    _ = v50
                                                } else {
                                                    if (!((fi31 + fi31) > fi31)) {
                                                        for fi53 := 0; fi53 < 11; fi53++ {
                                                            p2 = (p2 ^ (p2 + p2))
                                                            _ = fi53
                                                        }
                                                        fmt.Println(v43)
                                                        fmt.Println(fi31, p1)
                                                    } else {
                                                        var v54 int16 = (-(v13))
                                                        _ = v54
                                                    }
                                                }
                                            case "\n\t":
                                                var v55 uint8 = p0
                                                _ = v55
                                            }
                                        default:
                                            v13 = (v13 + v13)
                                        }
                                    }
                                    var v56 int16 = (v13 % ((-1) | 1))
                                    v57 := []int{}
                                    _ = v56
                                    _ = v57
                                }
                                v58 := GenS1{F0: fi31}
                                _ = v43
                                _ = v58
                            }
                        }
                    } else {
                        fmt.Println(p0, p2)
                        switch fi31 {
                        case -452:
                            var v59 float64 = 0.1
                            _ = v59
                        case -181:
                            v13 += 0
                        case 2147483647:
                            var v60 int = (-(fi31))
                            _ = v60
                        }
                    }
                }
                switch "a" {
                case "foo bar":
                    if (((v13 > (v13 % (((v13 * v13)) | 1))) || false) && false) {
                        var v61 uint8 = p0
                        if false {
                            v62 := GenS3{F0: 0, F1: -130, F2: (fi31 + 142)}
                            _ = v62
                        }
                        var v63 float32 = 1e10
                        _ = v61
                        _ = v63
                    }
                default:
                    p2 = (p2 + (p2 >> ((4294967295) & 63)))
                }
                _ = fi31
                _ = v32
            }
        }
    }
    p2 = (p2 / ((p2) | 1))
    var v64 int = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v64
    return (^(p1))
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]bool)
    v3 := make(map[int]uint32)
    var v4 string = ("" + "hello")
    var v5 int32 = -575
    v6 := []string{v4, v4, v4, "\n\t"}
    if false {
        var v65 int = 0
        if true {
            var v66 int16 = -32768
            fmt.Println(v65)
            if false {
                var v67 int32 = v5
                switch v65 {
                case 1:
                    for fi68 := 0; fi68 < 6; fi68++ {
                        for fi69 := 0; fi69 < 3; fi69++ {
                            v5 += 2147483647
                            var v70 uint32 = 4294967295
                            switch v4 {
                            case "":
                                v71 := GenS4{F0: v5, F1: 127}
                                _ = v71
                            }
                            _ = fi69
                            _ = v70
                        }
                        _ = fi68
                    }
                case -9223372036854775808:
                    var v72 uint16 = 0
                    _ = v72
                default:
                    v4 = v4
                }
                _ = v67
            }
            _ = v66
        } else {
            if true {
                var v73 int64 = 1
                _ = v73
            } else {
                switch (v4 + (v4 + v4)) {
                case "\n\t":
                    for fi74 := 0; fi74 < 10; fi74++ {
                        if false {
                            v6 = append(v6, v4)
                            var v75 string = v4
                            for fi76 := 0; fi76 < 3; fi76++ {
                                delete(v3, -9223372036854775808)
                                _ = fi76
                            }
                            _ = v75
                        } else {
                            if (fi74 >= -314) {
                                switch v4 {
                                case "\n\t":
                                    delete(v3, -1)
                                case "a":
                                    fmt.Println(v5)
                                case "":
                                    _ = len(v2)
                                }
                                if ((!(genFn0(1000, 128, ((fi74 | (fi74 ^ v65)) <= fi74)) != 178)) && true) {
                                    v4 = v4
                                    v77 := GenS5{F0: 40, F1: 0, F2: 476}
                                    _ = v77
                                }
                            } else {
                                v5 += (-(v5))
                                for fi78 := 0; fi78 < 5; fi78++ {
                                    v79 := GenS4{F0: (-(v5)), F1: 255}
                                    _ = len(v2)
                                    switch v4 {
                                    case "a":
                                        if (!true) {
                                            if (false || ((((v4 + "hello") + v4) >= (v4 + ("\n\t" + "\n\t"))) && false)) {
                                                var v80 float64 = 1e10
                                                v5 = (-(v5))
                                                fmt.Println(v4, fi74)
                                                switch "hello" {
                                                case "hello":
                                                    fmt.Println(v4)
                                                }
                                                _ = v80
                                            }
                                        } else {
                                            v65 = fi78
                                        }
                                    case "hello":
                                        if false {
                                            v65 -= (v65 - 2147483647)
                                            v81 := GenS6{F0: "\n\t", F1: -83}
                                            _ = v81
                                        } else {
                                            var v82 int32 = (^(v5))
                                            _ = v82
                                        }
                                    case "\n\t":
                                        if (((-2.25 / 1.5) / 1.0) > -2.25) {
                                            if (255 < 2) {
                                                fmt.Println(v65)
                                                v65 -= 32767
                                            } else {
                                                v83 := GenS4{F0: v5, F1: 255}
                                                _ = v83
                                            }
                                            v5 = v5
                                            if true {
                                                var v84 uint32 = genFn1(1, genFn1(128, 0, 18446744073709551615), 18446744073709551615)
                                                _ = v84
                                            }
                                        }
                                    default:
                                        v65 -= (fi78 >> ((65535) & 63))
                                    }
                                    _ = fi78
                                    _ = v79
                                }
                                var v85 float32 = -2.25
                                _ = v85
                            }
                            fmt.Println(v5, fi74, fi74)
                            if (!(int32(fi74) != (v5 & v5))) {
                                switch 365 {
                                case -426:
                                    if (0 <= -230) {
                                        var v86 int8 = -128
                                        v87 := v2[v4]
                                        if (4294967295 < 448) {
                                            var v88 int32 = (v5 >> ((255) & 63))
                                            if true {
                                                v89 := GenS4{F0: v88, F1: 128}
                                                v5 = (v5 - (v88 * 42))
                                                for fi90 := 0; fi90 < 1; fi90++ {
                                                    fmt.Println(v65, v4, v5)
                                                    var v91 uint64 = 2
                                                    v92 := GenS4{F0: (v88 + int32(v86)), F1: uint8(v91)}
                                                    var v93 int = fi90
                                                    _ = fi90
                                                    _ = v91
                                                    _ = v92
                                                    _ = v93
                                                }
                                                _ = v89
                                            }
                                            _ = v88
                                        } else {
                                            for fi94 := 0; fi94 < 10; fi94++ {
                                                v3[(v65 >> ((591) & 63))] = 565
                                                if false {
                                                    fmt.Println(fi94, v87, fi94)
                                                    _ = len(v6)
                                                    v95 := GenS5{F0: 0, F1: (^(v86)), F2: int16(v86)}
                                                    _ = v95
                                                } else {
                                                    v5 = (-(v5))
                                                }
                                                _ = fi94
                                            }
                                            _ = len(v6)
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        }
                                        _ = v86
                                        _ = v87
                                    } else {
                                        var v96 uint32 = 2147483647
                                        _ = v96
                                    }
                                case -9223372036854775808:
                                    if true {
                                        var v97 uint8 = 2
                                        _ = v97
                                    }
                                case 822:
                                    for fi98 := 0; fi98 < 15; fi98++ {
                                        if true {
                                            if (-848 != (fi98 * 635)) {
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                _ = len(v1)
                                            }
                                        } else {
                                            v99 := GenS7{F0: 32767, F1: -1, F2: ("a" + v4)}
                                            _ = v99
                                        }
                                        delete(v3, (fi98 + -91))
                                        _ = fi98
                                    }
                                }
                                if (false && false) {
                                    for fi100 := 0; fi100 < 2; fi100++ {
                                        fmt.Println(fi74, v65, v4)
                                        var v101 int8 = 42
                                        fmt.Println(v101, v101, v101)
                                        _ = fi100
                                        _ = v101
                                    }
                                } else {
                                    var v102 int16 = 0
                                    _ = v102
                                }
                                v103 := GenS2{F0: 1}
                                _ = v103
                            } else {
                                var v104 int64 = -284
                                _ = v104
                            }
                        }
                        var v105 uint16 = 129
                        _ = fi74
                        _ = v105
                    }
                case "a":
                    var v106 int8 = 100
                    _ = v106
                case "hello":
                    v107 := GenS3{F0: 103, F1: 42, F2: (v65 & (v65 & (v65 ^ v65)))}
                    _ = v107
                default:
                    var v108 int8 = 6
                    _ = v108
                }
                if false {
                    for fi109 := 0; fi109 < 20; fi109++ {
                        fmt.Println(v4, v65)
                        if true {
                            for fi110 := 0; fi110 < 10; fi110++ {
                                v111 := GenS8{F0: false, F1: v5}
                                _ = fi110
                                _ = v111
                            }
                            if true {
                                for fi112 := 0; fi112 < 13; fi112++ {
                                    fmt.Println(v5, v4)
                                    v2[(v4 + v4)] = (-171 <= 42)
                                    _ = fi112
                                }
                                fmt.Println(fi109, fi109, v4)
                            }
                        } else {
                            var v113 uint = 255
                            _ = v113
                        }
                        _ = fi109
                    }
                }
                var v114 int = v65
                _ = v114
            }
            if false {
                var v115 uint = 481
                fmt.Println(v115, v5)
                _ = v115
            } else {
                switch 579 {
                case 2147483647:
                    if true {
                        switch v65 {
                        case -377:
                            v65 += (v65 | 127)
                        case 9223372036854775807:
                            for fi116 := 0; fi116 < 11; fi116++ {
                                v4 = "hello"
                                fmt.Println(v65, v65)
                                fmt.Println(v5, v5, v4)
                                _ = fi116
                            }
                        }
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                }
                if false {
                    v5 -= (v5 % (((v5 & (v5 >> ((271) & 63)))) | 1))
                    var v117 float32 = 1.5
                    _ = v117
                } else {
                    for fi118 := 0; fi118 < 20; fi118++ {
                        v65 -= 9223372036854775807
                        if false {
                            var v119 int64 = -9223372036854775808
                            _ = v119
                        }
                        _ = fi118
                    }
                    for fi120 := 0; fi120 < 4; fi120++ {
                        delete(v2, ("" + ("\n\t" + v4)))
                        v121 := GenS9{F0: 42, F1: v4, F2: 18446744073709551615}
                        _ = fi120
                        _ = v121
                    }
                    v122 := GenS10{F0: (v5 >> ((1) & 63))}
                    _ = v122
                }
            }
        }
        _ = v65
    } else {
        if false {
            var v123 int = 32767
            for fi124 := 0; fi124 < 6; fi124++ {
                v6 = append(v6, ("a" + v4))
                _ = fi124
            }
            _ = v123
        }
    }
    switch v4 {
    case "\n\t":
        switch -801 {
        case 167:
            v125 := v3[665]
            _ = v125
        case -624:
            for fi126 := 0; fi126 < 4; fi126++ {
                var v127 uint64 = 747
                _ = fi126
                _ = v127
            }
        }
    case "foo bar":
        v3[-9223372036854775808] = 2
    default:
        var v128 float32 = -2.25
        _ = v128
    }
    v5 = v5
    switch v4 {
    case "foo bar":
        v2[((v4 + ((("\n\t" + ("hello" + v4)) + ("\n\t" + "hello")) + "foo bar")) + "a")] = (v4 < (v4 + v4))
    case "hello":
        v4 = v4
    default:
        v129 := GenS1{F0: 32767}
        _ = v129
    }
    v5 = v5
    for fi130 := 0; fi130 < 19; fi130++ {
        if (!(!true)) {
            var v131 bool = true
            _ = v131
        } else {
            for fi132 := 0; fi132 < 12; fi132++ {
                var v133 float64 = (0.1 - 0.1)
                _ = fi132
                _ = v133
            }
            v5 = v5
        }
        var v134 uint16 = 42
        _ = fi130
        _ = v134
    }
    if false {
        fmt.Println(v5, v4)
        v135 := GenS5{F0: -190, F1: 91, F2: -767}
        _ = v135
    } else {
        switch 32767 {
        case 42:
            v5 = (v5 ^ v5)
        case 2147483647:
            for fi136 := 0; fi136 < 13; fi136++ {
                fmt.Println(v5, fi136)
                var v137 string = ((v4 + (("hello" + ("foo bar" + "hello")) + "\n\t")) + (v4 + "\n\t"))
                _ = fi136
                _ = v137
            }
        case 857:
            for fi138 := 0; fi138 < 2; fi138++ {
                fmt.Println(v4, v4)
                _ = fi138
            }
        }
        switch v4 {
        case "\n\t":
            _ = len(v1)
        case "a":
            var v139 float64 = 1e10
            _ = v139
        case "hello":
            switch -9223372036854775808 {
            case -469:
                v5 -= (v5 ^ v5)
            case -144:
                switch v4 {
                case "foo bar":
                    if true {
                        var v140 float32 = ((3.14159 - (0.0 / 1.5)) / -2.25)
                        fmt.Println(v5, v4, v5)
                        var v141 bool = (686 > 2147483647)
                        _ = v140
                        _ = v141
                    } else {
                        switch "" {
                        case "\n\t":
                            fmt.Println(v5)
                        }
                    }
                }
            case -92:
                for fi142 := 0; fi142 < 15; fi142++ {
                    v143 := GenS8{F0: true, F1: (^(v5))}
                    _ = fi142
                    _ = v143
                }
            default:
                v144 := v3[32768]
                _ = v144
            }
        }
        v145 := GenS7{F0: 2, F1: -1, F2: ""}
        _ = v145
    }
    fmt.Println(v5, v4)
    switch "\n\t" {
    case "foo bar":
        if false {
            v4 = v4
            for fi146 := 0; fi146 < 5; fi146++ {
                for fi147 := 0; fi147 < 12; fi147++ {
                    v148 := GenS4{F0: v5, F1: 1}
                    _ = fi147
                    _ = v148
                }
                v4 = ("a" + "\n\t")
                v149 := GenS1{F0: (fi146 - 128)}
                _ = fi146
                _ = v149
            }
            v150 := GenS10{F0: (^(v5))}
            _ = v150
        } else {
            for fi151 := 0; fi151 < 15; fi151++ {
                v4 = v4
                for fi152 := 0; fi152 < 8; fi152++ {
                    v153 := GenS1{F0: (fi152 >> ((820) & 63))}
                    var v154 int = fi152
                    _ = fi152
                    _ = v153
                    _ = v154
                }
                _ = fi151
            }
            for fi155 := 0; fi155 < 20; fi155++ {
                if len(v6) > 0 {
                    v6[0] = (("" + v4) + ((v4 + "\n\t") + v4))
                }
                v156 := GenS8{F0: true, F1: v5}
                v157 := GenS7{F0: 0, F1: -96, F2: v4}
                _ = fi155
                _ = v156
                _ = v157
            }
            if (!((v5 | v5) <= v5)) {
                if ((!(128 > 1)) && true) {
                    var v158 uint = 2
                    var v159 uint8 = 1
                    _ = v158
                    _ = v159
                } else {
                    fmt.Println(v4, v4)
                    for fi160 := 0; fi160 < 12; fi160++ {
                        switch "a" {
                        case "":
                            if false {
                                if len(v6) > 0 {
                                    v6[0] = v4
                                }
                                var v161 string = (("\n\t" + (v4 + v4)) + (v4 + ("foo bar" + "")))
                                v5 += 2147483647
                                _ = v161
                            }
                        case "a":
                            for fi162 := 0; fi162 < 8; fi162++ {
                                v4 = v4
                                var v163 int8 = 100
                                var v164 uint16 = 32767
                                if false {
                                    var v165 int8 = -1
                                    if (-112 < 9223372036854775807) {
                                        var v166 int8 = v163
                                        v164 = (v164 / ((v164) | 1))
                                        _ = v166
                                    }
                                    var v167 float32 = (1.5 + (float32(float64(1e10)) * 0.1))
                                    _ = v165
                                    _ = v167
                                } else {
                                    v6 = append(v6, "")
                                    if (0.1 < -2.25) {
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        var v168 uint64 = 519
                                        delete(v1, genFn1(137, 620, 330))
                                        _ = v168
                                    }
                                }
                                var v169 uint8 = 127
                                _ = fi162
                                _ = v163
                                _ = v164
                                _ = v169
                            }
                        case "foo bar":
                            v4 = "foo bar"
                        default:
                            if true {
                                if false {
                                    fmt.Println(v5, v5)
                                    if (-586 != 1) {
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        var v170 bool = false
                                        _ = v170
                                    }
                                } else {
                                    v171 := GenS9{F0: 42, F1: "a", F2: 42}
                                    for fi172 := 0; fi172 < 18; fi172++ {
                                        delete(v1, 143)
                                        if false {
                                            var v173 int8 = 104
                                            var v174 bool = false
                                            switch fi172 {
                                            case 4294967295:
                                                v4 = v4
                                            default:
                                                var v175 uint64 = 1
                                                _ = v175
                                            }
                                            _ = v173
                                            _ = v174
                                        } else {
                                            fmt.Println(v4, v4, fi160)
                                            v176 := GenS11{F0: v5}
                                            if (!(4294967295 >= 42)) {
                                                v177 := GenS12{F0: 65535, F1: 1, F2: 42}
                                                _ = v177
                                            } else {
                                                if false {
                                                    v178 := GenS13{F0: 669, F1: 578}
                                                    v6 = append(v6, "a")
                                                    _ = v178
                                                }
                                            }
                                            if len(v6) > 0 {
                                                v6[0] = ((v4 + v4) + (v4 + v4))
                                            }
                                            v179 := GenS5{F0: 0, F1: -128, F2: 32767}
                                            v6 = append(v6, v4)
                                            v180 := GenS14{F0: 1, F1: fi160}
                                            _ = v176
                                            _ = v179
                                            _ = v180
                                        }
                                        if (false || true) {
                                            var v181 float64 = (1e10 + ((1.5 - (1e10 / -2.25)) / 1e10))
                                            _ = v181
                                        }
                                        _ = fi172
                                    }
                                    _ = v171
                                }
                                v1[966] = v4
                            } else {
                                switch "\n\t" {
                                case "a":
                                    if true {
                                        v4 = (v4 + ("a" + v4))
                                        switch (fi160 << ((65535) & 63)) {
                                        case 911:
                                            var v182 int8 = 23
                                            _ = v182
                                        case 32767:
                                            var v183 uint32 = 2
                                            _ = v183
                                        }
                                    } else {
                                        for fi184 := 0; fi184 < 9; fi184++ {
                                            var v185 bool = false
                                            v185 = false
                                            _ = fi184
                                            _ = v185
                                        }
                                    }
                                case "foo bar":
                                    if true {
                                        if (false || (genFn0(-305, (^(fi160)), true) > 42)) {
                                            var v186 uint16 = 170
                                            v4 = (v4 + v4)
                                            _ = v186
                                        } else {
                                            v4 = v4
                                            switch fi160 {
                                            case -517:
                                                var v187 uint64 = 4294967295
                                                _ = v187
                                            case 32767:
                                                v188 := GenS4{F0: v5, F1: 2}
                                                _ = v188
                                            default:
                                                fmt.Println(fi160)
                                            }
                                            for fi189 := 0; fi189 < 18; fi189++ {
                                                var v190 int16 = 0
                                                switch fi189 {
                                                case 32767:
                                                    _ = len(v1)
                                                case -9223372036854775808:
                                                    v190 += -19
                                                case 716:
                                                    var v191 bool = false
                                                    _ = v191
                                                default:
                                                    _ = len(v6)
                                                    v190 += v190
                                                    if (false && false) {
                                                        var v192 uint16 = 922
                                                        v190 += (v190 << ((218) & 63))
                                                        fmt.Println(v4)
                                                        _ = v192
                                                    } else {
                                                        if true {
                                                            v190 -= (v190 ^ (v190 ^ v190))
                                                            if (127 >= -37) {
                                                                v193 := GenS6{F0: v4, F1: -1}
                                                                switch fi160 {
                                                                case -2147483648:
                                                                    var v194 string = v4
                                                                    fmt.Println(v190, v5, v194)
                                                                    _ = v194
                                                                default:
                                                                    if false {
                                                                        v195 := GenS12{F0: 42, F1: 2, F2: 255}
                                                                        v5 += v5
                                                                        _ = v195
                                                                    } else {
                                                                        if (2147483648 > 2147483647) {
                                                                            v196 := GenS10{F0: (v5 + -2147483648)}
                                                                            var v197 int64 = -9223372036854775808
                                                                            _ = len(v3)
                                                                            _ = v196
                                                                            _ = v197
                                                                        } else {
                                                                            var v198 int = -358
                                                                            switch "" {
                                                                            case "foo bar":
                                                                                v5 -= (v5 * v5)
                                                                                if false {
                                                                                    var v199 string = ("foo bar" + "foo bar")
                                                                                    var v200 uint32 = 603
                                                                                    v3[v198] = (v200 + v200)
                                                                                    v201 := GenS2{F0: (v190 << ((42) & 63))}
                                                                                    _ = v199
                                                                                    _ = v200
                                                                                    _ = v201
                                                                                } else {
                                                                                    fmt.Println(v5, v190)
                                                                                    v198 += (fi189 / (((v198 | fi189)) | 1))
                                                                                }
                                                                                if len(v6) >= 2 {
                                                                                    v6 = v6[1:len(v6)]
                                                                                }
                                                                            case "\n\t":
                                                                                var v202 bool = false
                                                                                _ = v202
                                                                            default:
                                                                                v190 -= (v190 + (v190 | (v190 | 1000)))
                                                                            }
                                                                            _ = v198
                                                                        }
                                                                    }
                                                                }
                                                                if true {
                                                                    fmt.Println(v190, v5, v4)
                                                                    if (v5 == (v5 << ((567) & 63))) {
                                                                        switch "" {
                                                                        case "":
                                                                            if len(v6) > 0 {
                                                                                v6[0] = v4
                                                                            }
                                                                        case "a":
                                                                            v5 += (v5 - v5)
                                                                        default:
                                                                            var v203 uint8 = 1
                                                                            _ = v203
                                                                        }
                                                                    }
                                                                    v190 = (v190 << ((205) & 63))
                                                                }
                                                                _ = v193
                                                            }
                                                            _ = len(v6)
                                                        }
                                                        var v204 float64 = ((0.1 - float64((float32(0.1) * 3.14159))) * 1e10)
                                                        v204 += v204
                                                        _ = v204
                                                    }
                                                    v205 := GenS15{F0: 1, F1: (-(fi189)), F2: 4294967295}
                                                    _ = v205
                                                }
                                                v206 := GenS16{F0: 187, F1: v4}
                                                switch (fi160 & fi189) {
                                                case 0:
                                                    var v207 bool = (382 <= 955)
                                                    _ = v207
                                                default:
                                                    fmt.Println(fi160, fi189, fi189)
                                                }
                                                _ = fi189
                                                _ = v190
                                                _ = v206
                                            }
                                        }
                                        fmt.Println(fi160, v4, fi160)
                                    }
                                default:
                                    v5 = v5
                                }
                            }
                        }
                        _ = fi160
                    }
                }
                fmt.Println(v4, v4)
                for fi208 := 0; fi208 < 9; fi208++ {
                    var v209 int16 = -1
                    switch (fi208 | (fi208 & (fi208 / (((^(fi208))) | 1)))) {
                    case 128:
                        v5 = -553
                    default:
                        var v210 int8 = 100
                        _ = v210
                    }
                    _ = fi208
                    _ = v209
                }
            }
        }
    case "\n\t":
        fmt.Println(v4, v5)
    case "hello":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    default:
        var v211 uint8 = 36
        _ = v211
    }
    v4 = "\n\t"
    var v212 string = "\n\t"
    if false {
        for fi213 := 0; fi213 < 7; fi213++ {
            for fi214 := 0; fi214 < 20; fi214++ {
                v4 = "hello"
                var v215 float32 = (1.5 + (float32(-2.25) - -2.25))
                var v216 string = "\n\t"
                _ = len(v6)
                _ = fi214
                _ = v215
                _ = v216
            }
            v217 := GenS17{F0: 1}
            _ = fi213
            _ = v217
        }
        switch 929 {
        case 127:
            if false {
                if false {
                    if true {
                        for fi218 := 0; fi218 < 19; fi218++ {
                            v219 := GenS18{F0: 18446744073709551615, F1: (^(v5)), F2: 0}
                            _ = fi218
                            _ = v219
                        }
                    } else {
                        v220 := GenS12{F0: 4294967295, F1: 2147483647, F2: 1}
                        for fi221 := 0; fi221 < 4; fi221++ {
                            v5 += (v5 >> ((42) & 63))
                            _ = fi221
                        }
                        if (false && (1 >= 1)) {
                            fmt.Println(v212, v5, v5)
                            v222 := GenS13{F0: 559, F1: 0}
                            var v223 uint16 = 0
                            _ = v222
                            _ = v223
                        } else {
                            fmt.Println(v212, v4, v4)
                        }
                        _ = v220
                    }
                    switch "" {
                    case "foo bar":
                        v224 := v1[2147483648]
                        _ = v224
                    case "a":
                        if (3.14159 == 0.0) {
                            var v225 uint16 = 0
                            var v226 uint8 = 1
                            _ = v225
                            _ = v226
                        } else {
                            var v227 string = "foo bar"
                            switch ("" + v4) {
                            case "hello":
                                var v228 int32 = v5
                                _ = v228
                            default:
                                v3[-9223372036854775808] = 221
                            }
                            _ = v227
                        }
                    case "\n\t":
                        var v229 bool = false
                        _ = v229
                    }
                }
                fmt.Println(v4)
                var v230 uint64 = 0
                _ = v230
            } else {
                v4 = v4
            }
        case 675:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        default:
            if true {
                var v231 int16 = -126
                _ = v231
            } else {
                v212 = v4
                var v232 string = v212
                v233 := v2[(((v4 + v212) + (("foo bar" + v232) + "")) + v212)]
                _ = v232
                _ = v233
            }
        }
    } else {
        v3[981] = 490
        var v234 uint = 18446744073709551615
        _ = v234
    }
    v235 := GenS15{F0: 128, F1: 9223372036854775807, F2: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v212:", v212)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v212
    _ = v235
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
