package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 uint32; F2 string }
type GenS5 struct { F0 int16; F1 int64; F2 bool }
type GenS6 struct { F0 int; F1 bool }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint32; F1 int8 }
type GenS11 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS12 struct { F0 int16; F1 uint32 }
type GenS13 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 bool; F1 int; F2 uint }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 int64; F1 uint8 }
type GenS18 struct { F0 string; F1 int32 }
type GenS19 struct { F0 int64; F1 string; F2 int }
type GenS20 struct { F0 uint; F1 int16; F2 uint64 }
func genFn0(p0 uint64) uint8 {
    v6 := GenS1{F0: p0, F1: 4294967295, F2: p0}
    p0 = p0
    v7 := []int{32767, -214}
    var v8 int8 = 42
    _ = len(v7)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 126
}
func genFn1() uint16 {
    v9 := []int{42, -1, 500}
    if len(v9) > 0 {
        v9[0] = 32767
    }
    var v10 uint32 = 798
    var v11 uint32 = v10
    _ = v10
    _ = v11
    _ = v9
    return 42
}

func main() {
    var v1 int8 = -1
    var v2 uint64 = 161
    var v3 uint8 = 1
    var v4 int32 = 132
    var v5 string = ""
    var v12 int64 = 625
    var v13 int = 339
    v4 = (v4 << ((1) & 63))
    switch "" {
    case "\n\t":
        for fi14 := 0; fi14 < 18; fi14++ {
            if (v12 > v12) {
                var v15 float64 = 0.0
                v4 += 2147483647
                for fi16 := 0; fi16 < 12; fi16++ {
                    var v17 int16 = -32768
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
            if ((0.1 / 1.0) <= 1e10) {
                for fi18 := 0; fi18 < 12; fi18++ {
                    if (988 <= 42) {
                        v3 += (v3 + genFn0((v2 << ((0) & 63))))
                        var v19 int16 = -32768
                        var v20 int64 = v12
                        _ = v19
                        _ = v20
                    }
                    if false {
                        fmt.Println(fi14, v12, fi18)
                        var v21 float32 = (1.5 - float32(1e10))
                        switch v5 {
                        case "a":
                            var v22 int16 = -1
                            _ = v22
                        case "hello":
                            switch 2147483647 {
                            case 128:
                                if true {
                                    v23 := GenS1{F0: v2, F1: 2, F2: v2}
                                    var v24 uint32 = 42
                                    v25 := GenS1{F0: (v2 >> ((255) & 63)), F1: 2147483648, F2: (v2 >> ((65535) & 63))}
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                            default:
                                if true {
                                    var v26 float64 = (((1.5 / 1.0) + (float64(((v21 * float32(float64((1e10 - v21)))) - 1e10)) * 1e10)) - (((-2.25 + 3.14159) / 1e10) * float64(v21)))
                                    _ = v26
                                }
                                var v27 float32 = v21
                                var v28 int = v13
                                _ = v27
                                _ = v28
                            }
                        case "foo bar":
                            switch v5 {
                            case "":
                                v4 += (v4 + 2147483647)
                            case "a":
                                if ((((!false) || true) && true) && true) {
                                    if true {
                                        v29 := GenS2{F0: fi18, F1: ("hello" + "")}
                                        var v30 uint = uint(v2)
                                        v4 = int32(v2)
                                        _ = v29
                                        _ = v30
                                    }
                                }
                            }
                        }
                        _ = v21
                    } else {
                        var v31 uint64 = v2
                        v1 = 127
                        _ = v31
                    }
                    v1 -= (v1 << ((1) & 63))
                    _ = fi18
                }
                switch -1 {
                case 2147483647:
                    for fi32 := 0; fi32 < 8; fi32++ {
                        switch int(v3) {
                        case 32768:
                            var v33 bool = ((!(487 >= int16(v12))) || false)
                            _ = v33
                        case -647:
                            if ((0.1 - 0.1) < 0.1) {
                                v2 += 18446744073709551615
                            } else {
                                var v34 uint64 = 2
                                v35 := GenS2{F0: (fi32 / ((-602) | 1)), F1: ((("hello" + (("\n\t" + v5) + (v5 + "foo bar"))) + "hello") + "")}
                                _ = v34
                                _ = v35
                            }
                        case 927:
                            switch "hello" {
                            case "a":
                                switch 2147483647 {
                                case 150:
                                    switch (v13 - v13) {
                                    case 32768:
                                        v36 := []int{v13, v13, (-(fi14)), (v13 | 1)}
                                        _ = v36
                                    case 32767:
                                        if true {
                                            v12 += (v12 << ((609) & 63))
                                        }
                                    default:
                                        v37 := []int{9223372036854775807, 128, (-(fi32))}
                                        _ = v37
                                    }
                                case 2147483647:
                                    v4 -= (v4 << ((42) & 63))
                                case 128:
                                    v38 := []int{}
                                    _ = v38
                                }
                            case "hello":
                                v39 := GenS3{F0: (v2 | v2)}
                                _ = v39
                            default:
                                var v40 int16 = -1
                                _ = v40
                            }
                        }
                        switch ("hello" + v5) {
                        case "\n\t":
                            v13 = (-(v13))
                        }
                        _ = fi32
                    }
                case 20:
                    v41 := GenS4{F0: (v2 << ((771) & 63)), F1: 2147483648, F2: v5}
                    _ = v41
                case -9223372036854775808:
                    switch v5 {
                    case "a":
                        v42 := []int{953, (fi14 + v13), (v13 << ((65535) & 63)), v13}
                        _ = v42
                    case "foo bar":
                        fmt.Println(v3, v5, v1)
                    }
                }
            }
            _ = fi14
        }
    default:
        v43 := GenS5{F0: -377, F1: v12, F2: false}
        _ = v43
    }
    switch (v5 + (v5 + (v5 + ("foo bar" + v5)))) {
    case "a":
        var v44 int32 = -782
        _ = v44
    case "\n\t":
        fmt.Println(v2)
    case "hello":
        switch 9223372036854775807 {
        case 32768:
            v45 := GenS1{F0: 296, F1: uint32(v2), F2: (v2 / ((v2) | 1))}
            _ = v45
        case -117:
            if (0.1 > 0.1) {
                if (29 > 48) {
                    switch (v13 * v13) {
                    case 476:
                        for fi46 := 0; fi46 < 15; fi46++ {
                            for fi47 := 0; fi47 < 6; fi47++ {
                                var v48 int64 = v12
                                switch (fi47 % (((fi46 & 9223372036854775807)) | 1)) {
                                case -59:
                                    var v49 float32 = 0.1
                                    _ = v49
                                case 1:
                                    v4 = int32(v3)
                                default:
                                    v50 := []int{(fi46 % ((-9223372036854775808) | 1)), (fi47 >> ((0) & 63))}
                                    _ = v50
                                }
                                _ = fi47
                                _ = v48
                            }
                            v4 = (v4 ^ v4)
                            _ = fi46
                        }
                    default:
                        if (0.1 == (-2.25 - float32((float64((-2.25 + 0.0)) / -2.25)))) {
                            var v51 uint64 = 735
                            if (!(v13 <= int(v1))) {
                                v1 -= -128
                                v1 = (v1 + v1)
                            }
                            if (!((v1 - v1) < (v1 % ((96) | 1)))) {
                                for fi52 := 0; fi52 < 2; fi52++ {
                                    if false {
                                        for fi53 := 0; fi53 < 9; fi53++ {
                                            v3 = (v3 + (^(v3)))
                                            if (32767 != 194) {
                                                v54 := GenS2{F0: 4294967295, F1: ""}
                                                v55 := []int{v13, v13, (fi53 * 127), v13}
                                                _ = v54
                                                _ = v55
                                            } else {
                                                v13 = 0
                                            }
                                            _ = fi53
                                        }
                                        switch 32767 {
                                        case -1:
                                            for fi56 := 0; fi56 < 7; fi56++ {
                                                v13 -= 0
                                                _ = fi56
                                            }
                                        case 507:
                                            var v57 bool = false
                                            _ = v57
                                        case 128:
                                            switch v5 {
                                            case "a":
                                                v58 := GenS3{F0: v2}
                                                _ = v58
                                            case "hello":
                                                v59 := GenS6{F0: (fi52 * v13), F1: false}
                                                _ = v59
                                            default:
                                                v51 -= (v2 + 2)
                                            }
                                        }
                                    }
                                    _ = fi52
                                }
                                fmt.Println(v51)
                                if true {
                                    var v60 bool = (-32768 <= -1)
                                    if ((!(v1 > 90)) || v60) {
                                        var v61 string = (v5 + ((v5 + v5) + v5))
                                        v60 = v60
                                        fmt.Println(v1)
                                        _ = v61
                                    }
                                    for fi62 := 0; fi62 < 3; fi62++ {
                                        v60 = false
                                        if v60 {
                                            switch fi62 {
                                            case 42:
                                                if (v3 != 139) {
                                                    v63 := []int{int(v4), fi62, v13}
                                                    fmt.Println(v13)
                                                    _ = v63
                                                }
                                            }
                                            if v60 {
                                                if ((fi62 - 42) >= v13) {
                                                    v64 := []int{}
                                                    v13 = fi62
                                                    _ = v64
                                                }
                                                if ((v3 << ((2) & 63)) != 255) {
                                                    var v65 uint32 = 4294967295
                                                    v2 += 4294967295
                                                    _ = v65
                                                } else {
                                                    fmt.Println(v4, v2, v4)
                                                }
                                                v66 := GenS7{F0: (v2 << ((373) & 63)), F1: v4}
                                                _ = v66
                                            } else {
                                                for fi67 := 0; fi67 < 20; fi67++ {
                                                    v68 := []int{fi62, int(v12), (fi62 << ((458) & 63)), v13}
                                                    _ = fi67
                                                    _ = v68
                                                }
                                                fmt.Println(v60)
                                                v69 := GenS7{F0: 1, F1: v4}
                                                _ = v69
                                            }
                                        } else {
                                            fmt.Println(v5)
                                            switch fi62 {
                                            case -859:
                                                v3 -= v3
                                            case 18:
                                                v70 := []int{(v13 & 1), 128}
                                                _ = v70
                                            default:
                                                fmt.Println(v5, v4)
                                            }
                                            v4 += -2147483648
                                        }
                                        _ = fi62
                                    }
                                    _ = v60
                                } else {
                                    if true {
                                        v71 := GenS7{F0: (v2 << ((65535) & 63)), F1: v4}
                                        switch -922 {
                                        case 4294967295:
                                            var v72 float32 = (((0.1 / 3.14159) * float32(1.5)) + -2.25)
                                            _ = v72
                                        default:
                                            var v73 int16 = 1000
                                            _ = v73
                                        }
                                        _ = v71
                                    } else {
                                        v74 := []int{32768, v13}
                                        for fi75 := 0; fi75 < 8; fi75++ {
                                            if len(v74) >= 2 {
                                                v74 = v74[1:len(v74)]
                                            }
                                            var v76 uint = 65535
                                            _ = fi75
                                            _ = v76
                                        }
                                        _ = v74
                                    }
                                }
                            }
                            _ = v51
                        } else {
                            for fi77 := 0; fi77 < 16; fi77++ {
                                fmt.Println(fi77, v1, v1)
                                v5 = v5
                                if false {
                                    var v78 bool = (0 == 42)
                                    var v79 bool = false
                                    _ = v78
                                    _ = v79
                                }
                                _ = fi77
                            }
                            switch ("hello" + v5) {
                            case "\n\t":
                                for fi80 := 0; fi80 < 12; fi80++ {
                                    switch (v5 + v5) {
                                    case "hello":
                                        var v81 float32 = 3.14159
                                        _ = v81
                                    }
                                    var v82 int64 = -555
                                    _ = fi80
                                    _ = v82
                                }
                            }
                        }
                    }
                    v83 := GenS8{F0: -32768}
                    _ = v83
                } else {
                    v84 := GenS9{F0: v12}
                    _ = v84
                }
            }
        case 4294967295:
            v12 -= v12
        }
    default:
        if true {
            v13 = v13
        } else {
            var v85 uint16 = 1
            fmt.Println(v85, v2)
            _ = v85
        }
    }
    switch (^(v13)) {
    case -220:
        for fi86 := 0; fi86 < 5; fi86++ {
            v13 = fi86
            v87 := []int{}
            _ = fi86
            _ = v87
        }
    case 42:
        var v88 uint16 = 566
        _ = v88
    }
    if false {
        v2 += v2
        v89 := GenS10{F0: 42, F1: (v1 / (((v1 / ((v1) | 1))) | 1))}
        _ = v89
    }
    v1 = (^(v1))
    var v90 int8 = (v1 << ((uint(v2)) & 63))
    for fi91 := 0; fi91 < 11; fi91++ {
        if ((v3 - (v3 + 255)) == v3) {
            var v92 int32 = 1
            var v93 int8 = v90
            if (v92 == -419) {
                switch (v13 ^ fi91) {
                case 32768:
                    if ((v2 ^ v2) < 0) {
                        v2 = (v2 & v2)
                        fmt.Println(v5, fi91, v13)
                    } else {
                        if (true || true) {
                            var v94 float64 = -2.25
                            if ((v4 >> ((2) & 63)) == v92) {
                                switch ("" + "") {
                                case "\n\t":
                                    var v95 uint8 = 196
                                    _ = v95
                                case "":
                                    v96 := GenS5{F0: -609, F1: (v12 & v12), F2: false}
                                    _ = v96
                                default:
                                    var v97 uint = uint(v2)
                                    _ = v97
                                }
                                v98 := GenS11{F0: 0, F1: v3, F2: -145}
                                for fi99 := 0; fi99 < 11; fi99++ {
                                    if (0.1 > -2.25) {
                                        if true {
                                            var v100 string = "hello"
                                            switch v13 {
                                            case 127:
                                                switch ("a" + "") {
                                                case "foo bar":
                                                    var v101 int8 = (v93 & v93)
                                                    _ = v101
                                                case "a":
                                                    var v102 bool = true
                                                    _ = v102
                                                default:
                                                    v103 := GenS12{F0: int16(v12), F1: 0}
                                                    _ = v103
                                                }
                                            case -1:
                                                v12 -= (v12 >> ((18446744073709551615) & 63))
                                            }
                                            _ = v100
                                        } else {
                                            v90 = v90
                                        }
                                        v12 = v12
                                        v104 := GenS10{F0: 2147483647, F1: (^(v93))}
                                        _ = v104
                                    } else {
                                        v105 := GenS13{F0: (v2 + v2), F1: v12, F2: v2}
                                        fmt.Println(v12, v3, v92)
                                        _ = v105
                                    }
                                    var v106 uint16 = 42
                                    v107 := GenS14{F0: -126}
                                    _ = fi99
                                    _ = v106
                                    _ = v107
                                }
                                _ = v98
                            } else {
                                var v108 uint64 = v2
                                _ = v108
                            }
                            _ = v94
                        } else {
                            switch (v13 ^ v13) {
                            case -9223372036854775808:
                                if (false || ((3.14159 * -2.25) > (0.0 + 0.0))) {
                                    v109 := []int{670}
                                    _ = v109
                                } else {
                                    v13 = (v13 << ((1) & 63))
                                    var v110 uint32 = 42
                                    v2 = (v2 % ((uint64(v3)) | 1))
                                    _ = v110
                                }
                            default:
                                for fi111 := 0; fi111 < 5; fi111++ {
                                    var v112 int8 = v90
                                    v113 := []int{}
                                    fmt.Println(v90)
                                    if false {
                                        if (!(533 > 2)) {
                                            fmt.Println(v3, v2)
                                            v93 = (v93 | (v112 & -128))
                                        }
                                    }
                                    var v114 uint8 = (v3 * v3)
                                    v12 += (v12 & (v12 - (v12 - (v12 % ((-9223372036854775808) | 1)))))
                                    switch "hello" {
                                    case "foo bar":
                                        fmt.Println(fi91)
                                    case "":
                                        v115 := GenS2{F0: (fi111 << ((4294967295) & 63)), F1: "a"}
                                        _ = v115
                                    case "hello":
                                        var v116 uint8 = (v3 / ((v114) | 1))
                                        _ = v116
                                    default:
                                        var v117 int8 = (v90 | v1)
                                        _ = v117
                                    }
                                    switch v5 {
                                    case "foo bar":
                                        if false {
                                            if (!false) {
                                                switch 171 {
                                                case 32768:
                                                    if false {
                                                        if true {
                                                            var v118 int64 = 42
                                                            fmt.Println(fi111, v4, fi91)
                                                            _ = v118
                                                        }
                                                    } else {
                                                        var v119 int64 = 470
                                                        _ = v119
                                                    }
                                                }
                                            }
                                            var v120 uint = 0
                                            var v121 int32 = v4
                                            _ = v120
                                            _ = v121
                                        }
                                    case "":
                                        v122 := GenS15{F0: (!false), F1: (v13 + 0), F2: 2}
                                        _ = v122
                                    default:
                                        v90 -= (v112 - v93)
                                    }
                                    _ = fi111
                                    _ = v112
                                    _ = v113
                                    _ = v114
                                }
                            }
                            v4 -= (v92 & -231)
                            v123 := []int{(fi91 << ((431) & 63)), v13, (v13 / ((128) | 1)), -2147483648}
                            _ = v123
                        }
                        switch fi91 {
                        case 0:
                            var v124 uint64 = (v2 & 4294967295)
                            _ = v124
                        case -657:
                            v125 := []int{fi91}
                            _ = v125
                        case 9223372036854775807:
                            for fi126 := 0; fi126 < 6; fi126++ {
                                if true {
                                    v4 = (v92 & int32(v2))
                                }
                                fmt.Println(v1)
                                v127 := GenS16{F0: 0}
                                var v128 int8 = (^(v93))
                                switch v5 {
                                case "hello":
                                    var v129 int32 = (v92 << ((42) & 63))
                                    _ = v129
                                case "\n\t":
                                    v130 := GenS11{F0: (v3 % ((0) | 1)), F1: genFn0(v2), F2: int16(v3)}
                                    _ = v130
                                case "foo bar":
                                    switch "\n\t" {
                                    case "foo bar":
                                        v131 := GenS17{F0: 0, F1: v3}
                                        _ = v131
                                    }
                                default:
                                    v132 := []int{fi126, (v13 << ((119) & 63)), v13}
                                    _ = v132
                                }
                                _ = fi126
                                _ = v127
                                _ = v128
                            }
                        }
                        var v133 int32 = int32(v3)
                        _ = v133
                    }
                case 55:
                    var v134 uint = 121
                    _ = v134
                }
                v135 := []int{int(v1), (fi91 ^ fi91), (v13 >> ((1) & 63))}
                _ = v135
            } else {
                for fi136 := 0; fi136 < 19; fi136++ {
                    v137 := GenS18{F0: "hello", F1: v92}
                    _ = fi136
                    _ = v137
                }
                switch (fi91 & -9223372036854775808) {
                case -2147483648:
                    v4 = (v92 % ((v92) | 1))
                case 322:
                    var v138 int16 = 1000
                    _ = v138
                case -997:
                    if true {
                        v139 := GenS12{F0: 32767, F1: 1}
                        _ = v139
                    } else {
                        for fi140 := 0; fi140 < 19; fi140++ {
                            if (42 == 332) {
                                var v141 int16 = -32768
                                _ = v141
                            } else {
                                if (256 >= 42) {
                                    v3 = (v3 << ((0) & 63))
                                } else {
                                    v13 = (v13 >> ((1) & 63))
                                    var v142 uint16 = 0
                                    _ = v142
                                }
                                v143 := []int{}
                                _ = len(v143)
                                var v144 uint8 = v3
                                _ = v143
                                _ = v144
                            }
                            var v145 uint = 65535
                            _ = fi140
                            _ = v145
                        }
                        v146 := GenS19{F0: v12, F1: ("hello" + ((v5 + "\n\t") + (v5 + (((((v5 + v5) + v5) + "foo bar") + "a") + v5)))), F2: fi91}
                        _ = v146
                    }
                }
            }
            _ = v92
            _ = v93
        }
        v1 = (v1 - (v1 | (v90 * (v90 << ((2) & 63)))))
        _ = fi91
    }
    var v147 int64 = (v12 % (((v12 << ((18446744073709551615) & 63))) | 1))
    var v148 int64 = (^(v12))
    v147 += -9223372036854775808
    switch (v13 + 9223372036854775807) {
    case 1:
        v149 := GenS1{F0: v2, F1: 1, F2: 13}
        _ = v149
    case 42:
        if (4294967295 == 0) {
            for fi150 := 0; fi150 < 8; fi150++ {
                v151 := []int{42, v13, fi150}
                if false {
                    var v152 int = (-(v13))
                    v12 += (^(v147))
                    _ = v152
                } else {
                    for fi153 := 0; fi153 < 13; fi153++ {
                        switch (-(fi150)) {
                        case 42:
                            var v154 uint32 = 988
                            _ = v154
                        case 67:
                            var v155 uint32 = 2147483648
                            _ = v155
                        }
                        _ = fi153
                    }
                    var v156 int = (v13 - (fi150 & (v13 / ((1) | 1))))
                    _ = v156
                }
                var v157 string = (v5 + v5)
                _ = fi150
                _ = v151
                _ = v157
            }
            if true {
                var v158 uint32 = 42
                v4 = (v4 % ((v4) | 1))
                _ = v158
            }
            var v159 uint16 = 255
            _ = v159
        }
    case 854:
        if (!false) {
            if (true || (true && true)) {
                var v160 float32 = 0.1
                _ = v160
            } else {
                var v161 string = "hello"
                switch v5 {
                case "hello":
                    for fi162 := 0; fi162 < 12; fi162++ {
                        v163 := []int{(v13 | fi162), (v13 - fi162)}
                        _ = fi162
                        _ = v163
                    }
                case "\n\t":
                    switch 128 {
                    case -2147483648:
                        if false {
                            var v164 bool = (!(((float64(3.14159) >= 0.1) || false) || false))
                            _ = v164
                        }
                    case 32767:
                        if (65535 <= 647) {
                            v13 -= (v13 << ((42) & 63))
                        } else {
                            var v165 uint = 2
                            _ = v165
                        }
                    }
                }
                _ = v161
            }
            v5 = "hello"
        } else {
            v5 = v5
            v166 := []int{(v13 + 128)}
            _ = v166
        }
    default:
        v167 := GenS20{F0: 1, F1: 1000, F2: (v2 % (((v2 >> ((65535) & 63))) | 1))}
        _ = v167
    }
    for fi168 := 0; fi168 < 8; fi168++ {
        fmt.Println(v4, fi168)
        v169 := GenS1{F0: (v2 / (((v2 * (v2 ^ (v2 - v2)))) | 1)), F1: 2147483648, F2: (^(v2))}
        _ = fi168
        _ = v169
    }
    switch "a" {
    case "hello":
        v12 += (v148 + (v12 % (((v147 ^ (-(v147)))) | 1)))
    case "":
        if ((v5 + v5) <= "foo bar") {
            v170 := []int{806}
            fmt.Println(v148, v1)
            var v171 int8 = 0
            _ = v170
            _ = v171
        }
    default:
        v172 := []int{int(v2), 2147483647}
        _ = v172
    }
    var v173 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v147:", v147)
    fmt.Println("v148:", v148)
    fmt.Println("v173:", v173)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v12
    _ = v13
    _ = v147
    _ = v148
    _ = v173
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v90
}
