package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int8; F2 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 int8; F2 int8 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint; F1 uint16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int64; F1 string }
type GenS11 struct { F0 uint; F1 string; F2 int8 }
type GenS12 struct { F0 int8; F1 int }
type GenS13 struct { F0 bool; F1 uint8; F2 bool }
type GenS14 struct { F0 bool; F1 int32; F2 uint }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 uint64; F1 int64 }
type GenS19 struct { F0 uint64; F1 int32; F2 int16 }
func genFn0(p0 uint16) int {
    var v4 uint16 = (^(p0))
    fmt.Println(v4, p0)
    var v5 uint16 = (p0 ^ (^(v4)))
    _ = p0
    _ = v4
    _ = v5
    return -2147483648
}
func genFn1(p0 uint64) uint {
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 bool = (((p0 + p0) != (p0 ^ (p0 ^ (p0 >> ((4294967295) & 63))))) && ("foo bar" < ""))
            v9 := []int{(fi6 | (^(fi6)))}
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 bool = true
        _ = fi6
        _ = v10
    }
    var v11 uint16 = 32767
    _ = p0
    _ = v11
    return 667
}

func main() {
    var v1 int16 = 958
    var v2 uint16 = 2
    var v3 uint64 = 959
    if (true || (128 != 0)) {
        fmt.Println(v1, v1)
        var v12 float64 = 0.1
        v3 += v3
        _ = v12
    } else {
        v13 := GenS1{F0: ("a" == "a")}
        v1 = (-(v1))
        if (!false) {
            fmt.Println(v3)
        } else {
            var v14 float64 = ((1e10 * 3.14159) + 1e10)
            _ = v14
        }
        _ = v13
    }
    v1 -= (v1 << ((uint(v2)) & 63))
    v3 += v3
    var v15 uint32 = 461
    switch -1 {
    case 128:
        var v16 uint8 = 2
        _ = v16
    case -243:
        v17 := []int{}
        _ = v17
    }
    var v18 int64 = 644
    if (true || (v3 >= v3)) {
        v3 += 549
        v19 := GenS2{F0: v2, F1: -46, F2: -1}
        v20 := []int{-223}
        _ = v19
        _ = v20
    } else {
        if (1e10 < (float64(1e10) + (0.1 - -2.25))) {
            v3 -= (v3 ^ (v3 ^ (v3 + 18446744073709551615)))
        } else {
            v21 := []int{32767, 32768}
            switch "foo bar" {
            case "foo bar":
                var v22 uint = uint(v15)
                _ = v22
            default:
                if (float32((((0.0 / 0.1) * (((-2.25 * float64(-2.25)) - (3.14159 / (0.1 - ((0.1 - float64((-2.25 * 1.5))) - 0.0)))) - (float64(0.1) * ((0.0 + (0.1 - float64(-2.25))) - 1.5)))) / (float64(1e10) / float64(1e10)))) != ((0.1 + 0.1) + (-2.25 + 0.1))) {
                    var v23 int64 = (^(v18))
                    for fi24 := 0; fi24 < 18; fi24++ {
                        switch (fi24 << ((4294967295) & 63)) {
                        case 9223372036854775807:
                            v1 += -1
                        }
                        _ = fi24
                    }
                    v25 := GenS2{F0: v2, F1: 127, F2: 127}
                    _ = v23
                    _ = v25
                }
            }
            _ = v21
        }
    }
    for fi26 := 0; fi26 < 14; fi26++ {
        switch fi26 {
        case 32767:
            v27 := []int{fi26, (-(fi26)), fi26}
            _ = v27
        case -1:
            fmt.Println(v1)
        case -211:
            v28 := GenS3{F0: (v1 / ((v1) | 1))}
            _ = v28
        default:
            v18 += (v18 & v18)
        }
        _ = fi26
    }
    switch -1 {
    case 4294967295:
        switch 2147483647 {
        case 2147483647:
            var v29 float32 = 0.0
            _ = v29
        case -749:
            if true {
                if false {
                    if (-939 == -2147483648) {
                        if (!false) {
                            v3 = 1
                            v2 = 2
                            for fi30 := 0; fi30 < 10; fi30++ {
                                v1 += v1
                                for fi31 := 0; fi31 < 5; fi31++ {
                                    v2 += (v2 % (((v2 ^ v2)) | 1))
                                    _ = fi31
                                }
                                _ = fi30
                            }
                        } else {
                            for fi32 := 0; fi32 < 4; fi32++ {
                                v18 -= -1
                                _ = fi32
                            }
                        }
                        v33 := []int{-338, 569, 934, -909}
                        _ = v33
                    } else {
                        if false {
                            if true {
                                var v34 string = "\n\t"
                                _ = v34
                            } else {
                                if false {
                                    switch 32767 {
                                    case 9223372036854775807:
                                        v35 := []int{0, 414, -1, -574}
                                        _ = v35
                                    case 32768:
                                        for fi36 := 0; fi36 < 3; fi36++ {
                                            fmt.Println(v2, fi36)
                                            for fi37 := 0; fi37 < 14; fi37++ {
                                                if (-1 < -128) {
                                                    v38 := []int{int(v1)}
                                                    if (false && true) {
                                                        switch "a" {
                                                        case "":
                                                            v2 -= uint16(v3)
                                                        case "foo bar":
                                                            v2 += (v2 + uint16(v3))
                                                        }
                                                        v39 := GenS3{F0: v1}
                                                        var v40 int32 = 42
                                                        _ = v39
                                                        _ = v40
                                                    } else {
                                                        if false {
                                                            var v41 bool = false
                                                            _ = v41
                                                        }
                                                        if (((127 == 127) || (true && true)) && ((!(1 < uint32(v2))) || ((924 <= 255) || true))) {
                                                            v2 -= v2
                                                            if ((!(!((2 <= 457) || false))) || ((false && ((^(v2)) >= (v2 % (((^(v2))) | 1)))) && (!false))) {
                                                                if len(v38) >= 2 {
                                                                    v38 = v38[1:len(v38)]
                                                                }
                                                            } else {
                                                                fmt.Println(v3, fi36, fi36)
                                                                v42 := GenS4{F0: 2147483647, F1: 100, F2: 93}
                                                                _ = v42
                                                            }
                                                            if (2147483647 <= -1) {
                                                                var v43 float32 = 3.14159
                                                                v1 += v1
                                                                _ = v43
                                                            } else {
                                                                switch -1 {
                                                                case -757:
                                                                    switch "a" {
                                                                    case "":
                                                                        fmt.Println(v3)
                                                                    case "a":
                                                                        v15 += v15
                                                                    default:
                                                                        switch ((("hello" + "") + "foo bar") + "a") {
                                                                        case "foo bar":
                                                                            if (false || true) {
                                                                                var v44 string = ""
                                                                                switch fi36 {
                                                                                case -173:
                                                                                    if true {
                                                                                        var v45 uint64 = 18446744073709551615
                                                                                        _ = v45
                                                                                    } else {
                                                                                        v44 = v44
                                                                                        var v46 uint64 = v3
                                                                                        _ = v46
                                                                                    }
                                                                                case 9223372036854775807:
                                                                                    var v47 uint32 = v15
                                                                                    _ = v47
                                                                                default:
                                                                                    if true {
                                                                                        fmt.Println(v3)
                                                                                        var v48 float32 = 1.5
                                                                                        if len(v38) >= 2 {
                                                                                            v38 = v38[1:len(v38)]
                                                                                        }
                                                                                        _ = v48
                                                                                    } else {
                                                                                        var v49 uint = 4294967295
                                                                                        switch (fi36 << ((v49) & 63)) {
                                                                                        case 579:
                                                                                            fmt.Println(fi36, v1, fi37)
                                                                                        }
                                                                                        _ = v49
                                                                                    }
                                                                                }
                                                                                var v50 int = (fi37 - genFn0(v2))
                                                                                _ = v44
                                                                                _ = v50
                                                                            }
                                                                        case "":
                                                                            v38 = append(v38, genFn0(v2))
                                                                        default:
                                                                            v51 := GenS2{F0: 2, F1: 0, F2: 42}
                                                                            _ = v51
                                                                        }
                                                                    }
                                                                case -1:
                                                                    fmt.Println(v1, v18, v1)
                                                                }
                                                            }
                                                            v15 -= v15
                                                        }
                                                    }
                                                    _ = v38
                                                } else {
                                                    switch ("a" + ("" + (("a" + "\n\t") + ("a" + "\n\t")))) {
                                                    case "hello":
                                                        v52 := GenS1{F0: ((^(v3)) < 18446744073709551615)}
                                                        _ = v52
                                                    case "a":
                                                        v53 := GenS2{F0: (^(v2)), F1: 1, F2: -99}
                                                        _ = v53
                                                    case "\n\t":
                                                        var v54 float64 = 1.5
                                                        _ = v54
                                                    default:
                                                        v55 := GenS5{F0: (v15 * v15), F1: (v3 - (^(v3)))}
                                                        _ = v55
                                                    }
                                                    if true {
                                                        var v56 int8 = 100
                                                        var v57 uint64 = (^(v3))
                                                        v57 -= 340
                                                        _ = v56
                                                        _ = v57
                                                    }
                                                    v58 := GenS3{F0: v1}
                                                    _ = v58
                                                }
                                                if (false && true) {
                                                    v59 := []int{int(v2), (fi37 | -2147483648), (^(fi36)), 127}
                                                    _ = v59
                                                } else {
                                                    v60 := []int{int(v15), fi36}
                                                    v61 := GenS6{F0: (v2 * 32767)}
                                                    v62 := GenS4{F0: 1, F1: -1, F2: -73}
                                                    v63 := GenS4{F0: 2147483647, F1: -1, F2: 0}
                                                    _ = v60
                                                    _ = v61
                                                    _ = v62
                                                    _ = v63
                                                }
                                                switch ("hello" + "foo bar") {
                                                case "":
                                                    v1 += v1
                                                case "\n\t":
                                                    v18 = int64(v3)
                                                }
                                                _ = fi37
                                            }
                                            _ = fi36
                                        }
                                    }
                                }
                                var v64 float64 = 1.5
                                _ = v64
                            }
                            v65 := []int{0}
                            _ = v65
                        } else {
                            v3 -= (v3 - 4294967295)
                            switch -9223372036854775808 {
                            case 9223372036854775807:
                                var v66 uint = 2
                                _ = v66
                            case -823:
                                var v67 int64 = -445
                                _ = v67
                            case 127:
                                if true {
                                    v68 := []int{9223372036854775807, genFn0(0)}
                                    for fi69 := 0; fi69 < 13; fi69++ {
                                        var v70 uint8 = 42
                                        _ = len(v68)
                                        var v71 int64 = (v18 + v18)
                                        _ = fi69
                                        _ = v70
                                        _ = v71
                                    }
                                    switch 32768 {
                                    case -1:
                                        var v72 uint = 42
                                        _ = v72
                                    default:
                                        switch 0 {
                                        case -9223372036854775808:
                                            v73 := GenS4{F0: -36, F1: 100, F2: -1}
                                            _ = v73
                                        case 1:
                                            for fi74 := 0; fi74 < 10; fi74++ {
                                                _ = len(v68)
                                                var v75 int32 = -1
                                                v1 = (v1 >> ((0) & 63))
                                                _ = fi74
                                                _ = v75
                                            }
                                        default:
                                            var v76 float64 = (-2.25 - float64(3.14159))
                                            _ = v76
                                        }
                                    }
                                    _ = v68
                                } else {
                                    for fi77 := 0; fi77 < 5; fi77++ {
                                        v2 = (v2 * 590)
                                        for fi78 := 0; fi78 < 5; fi78++ {
                                            switch fi78 {
                                            case 0:
                                                var v79 int8 = 120
                                                _ = v79
                                            }
                                            var v80 string = ""
                                            if (false || false) {
                                                if (!true) {
                                                    v81 := []int{int(v1)}
                                                    _ = v81
                                                } else {
                                                    if true {
                                                        switch ("" + "a") {
                                                        case "foo bar":
                                                            v82 := GenS2{F0: (v2 | (v2 ^ (v2 | (^(v2))))), F1: 0, F2: 100}
                                                            _ = v82
                                                        case "":
                                                            if false {
                                                                v83 := GenS3{F0: v1}
                                                                if (false || (!(false || true))) {
                                                                    if ((true && (!(!(true || (v2 >= v2))))) && (v80 <= "hello")) {
                                                                        v84 := []int{(fi77 << ((uint(v15)) & 63)), (fi78 >> ((42) & 63)), fi77, 861}
                                                                        fmt.Println(fi77, v15)
                                                                        fmt.Println(v80, v15)
                                                                        var v85 uint8 = 248
                                                                        var v86 bool = (float64(3.14159) < -2.25)
                                                                        _ = v84
                                                                        _ = v85
                                                                        _ = v86
                                                                    } else {
                                                                        var v87 float64 = 0.0
                                                                        fmt.Println(v15, v1)
                                                                        var v88 uint16 = (v2 ^ 32767)
                                                                        _ = v87
                                                                        _ = v88
                                                                    }
                                                                    switch "a" {
                                                                    case "hello":
                                                                        v80 = v80
                                                                    case "a":
                                                                        var v89 uint32 = v15
                                                                        _ = v89
                                                                    case "foo bar":
                                                                        var v90 int16 = 1000
                                                                        _ = v90
                                                                    default:
                                                                        switch "a" {
                                                                        case "hello":
                                                                            v18 -= (v18 & 42)
                                                                        case "foo bar":
                                                                            v91 := GenS4{F0: 233, F1: -75, F2: 100}
                                                                            _ = v91
                                                                        }
                                                                    }
                                                                    v92 := []int{(^(fi77))}
                                                                    _ = v92
                                                                } else {
                                                                    var v93 uint32 = 2
                                                                    var v94 uint32 = (^(v15))
                                                                    v15 = (^(v93))
                                                                    _ = v93
                                                                    _ = v94
                                                                }
                                                                v95 := GenS3{F0: v1}
                                                                _ = v83
                                                                _ = v95
                                                            } else {
                                                                var v96 int64 = v18
                                                                fmt.Println(v80, v15, v2)
                                                                v3 = 658
                                                                _ = v96
                                                            }
                                                        case "hello":
                                                            var v97 int32 = 1
                                                            _ = v97
                                                        default:
                                                            if true {
                                                                v3 += 42
                                                                v1 -= (v1 ^ v1)
                                                            }
                                                        }
                                                    } else {
                                                        if false {
                                                            v1 = 1000
                                                            v98 := []int{4294967295, 32767, fi78}
                                                            _ = v98
                                                        } else {
                                                            fmt.Println(v18, v2, v18)
                                                            v2 -= (v2 - 65535)
                                                        }
                                                        fmt.Println(v3, fi77)
                                                        v99 := GenS5{F0: uint32(v2), F1: 0}
                                                        _ = v99
                                                    }
                                                    v1 += 300
                                                    v100 := GenS7{F0: 349}
                                                    var v101 float32 = (3.14159 / 1.0)
                                                    _ = v100
                                                    _ = v101
                                                }
                                            }
                                            _ = fi78
                                            _ = v80
                                        }
                                        if true {
                                            v102 := GenS8{F0: 969, F1: (^(v2))}
                                            fmt.Println(v18, v1, v3)
                                            var v103 uint8 = 255
                                            _ = v102
                                            _ = v103
                                        } else {
                                            v104 := []int{(^(fi77)), genFn0(v2)}
                                            v18 += v18
                                            _ = v104
                                        }
                                        _ = fi77
                                    }
                                    switch -485 {
                                    case 241:
                                        var v105 uint64 = v3
                                        _ = v105
                                    default:
                                        fmt.Println(v15, v18, v2)
                                    }
                                }
                            }
                            var v106 string = ("" + "foo bar")
                            _ = v106
                        }
                        for fi107 := 0; fi107 < 12; fi107++ {
                            for fi108 := 0; fi108 < 18; fi108++ {
                                switch "a" {
                                case "":
                                    v109 := GenS9{F0: (v2 - 65535)}
                                    _ = v109
                                case "\n\t":
                                    fmt.Println(v18)
                                case "foo bar":
                                    v110 := []int{(fi107 | (-(fi107)))}
                                    if (160 <= 1) {
                                        if false {
                                            if len(v110) >= 2 {
                                                v110 = v110[1:len(v110)]
                                            }
                                        }
                                    }
                                    _ = v110
                                }
                                _ = fi108
                            }
                            _ = fi107
                        }
                        var v111 string = ("hello" + ("" + ""))
                        _ = v111
                    }
                }
                switch "\n\t" {
                case "\n\t":
                    if (42 <= 42) {
                        v112 := []int{0, 2147483647}
                        v18 = -1
                        if len(v112) > 0 {
                            v112[0] = 761
                        }
                        _ = v112
                    }
                case "":
                    v18 -= v18
                }
                fmt.Println(v3)
            } else {
                if (v2 != 785) {
                    if (true && (v1 == v1)) {
                        if true {
                            v113 := []int{}
                            _ = v113
                        } else {
                            v114 := []int{-9223372036854775808, -879}
                            _ = v114
                        }
                        for fi115 := 0; fi115 < 13; fi115++ {
                            v116 := []int{fi115}
                            fmt.Println(fi115)
                            _ = fi115
                            _ = v116
                        }
                        for fi117 := 0; fi117 < 9; fi117++ {
                            v118 := GenS5{F0: (^(v15)), F1: (v3 + (v3 ^ (v3 ^ 0)))}
                            _ = fi117
                            _ = v118
                        }
                    }
                    for fi119 := 0; fi119 < 10; fi119++ {
                        if ((false && false) || false) {
                            var v120 int16 = 220
                            v121 := GenS9{F0: (v2 & 2)}
                            if false {
                                var v122 uint = 1
                                fmt.Println(v3)
                                if (false || false) {
                                    var v123 float64 = (3.14159 + 0.1)
                                    for fi124 := 0; fi124 < 3; fi124++ {
                                        if (true || (396 != v2)) {
                                            if false {
                                                if true {
                                                    var v125 uint = (v122 * v122)
                                                    _ = v125
                                                }
                                                var v126 uint = (v122 << ((v122) & 63))
                                                var v127 uint64 = (v3 * v3)
                                                _ = v126
                                                _ = v127
                                            }
                                            var v128 int64 = (-(v18))
                                            _ = v128
                                        }
                                        fmt.Println(v2, v15)
                                        _ = fi124
                                    }
                                    _ = v123
                                } else {
                                    v120 = v1
                                    v18 = v18
                                    v122 += v122
                                }
                                _ = v122
                            } else {
                                v129 := []int{}
                                if (v3 <= (^(v3))) {
                                    v130 := GenS10{F0: int64(fi119), F1: (((("" + ("a" + "foo bar")) + ("foo bar" + "")) + ("a" + (("\n\t" + "hello") + "foo bar"))) + (((("a" + "foo bar") + "a") + "foo bar") + "a"))}
                                    _ = v130
                                } else {
                                    var v131 bool = true
                                    _ = v131
                                }
                                _ = v129
                            }
                            _ = v120
                            _ = v121
                        }
                        for fi132 := 0; fi132 < 13; fi132++ {
                            var v133 uint = 42
                            switch ("hello" + "a") {
                            case "hello":
                                switch ("foo bar" + "a") {
                                case "\n\t":
                                    if (42 <= 128) {
                                        fmt.Println(v1, fi132, fi119)
                                        if (-128 != -1) {
                                            var v134 int64 = v18
                                            v135 := GenS11{F0: genFn1(v3), F1: ("\n\t" + "\n\t"), F2: -25}
                                            _ = v134
                                            _ = v135
                                        } else {
                                            if (false || true) {
                                                var v136 int64 = v18
                                                if false {
                                                    v137 := GenS12{F0: -128, F1: genFn0((v2 ^ (v2 % (((^(v2))) | 1))))}
                                                    v138 := GenS13{F0: false, F1: 155, F2: true}
                                                    if (((831 >= 1) && true) && (-99 != 127)) {
                                                        var v139 uint = v133
                                                        var v140 uint64 = v3
                                                        var v141 int16 = v1
                                                        _ = v139
                                                        _ = v140
                                                        _ = v141
                                                    } else {
                                                        if (false || false) {
                                                            if false {
                                                                var v142 float32 = -2.25
                                                                var v143 int16 = int16(v133)
                                                                v144 := []int{fi119}
                                                                var v145 uint16 = v2
                                                                _ = v142
                                                                _ = v143
                                                                _ = v144
                                                                _ = v145
                                                            } else {
                                                                var v146 int16 = (v1 + -32768)
                                                                switch (fi132 + 2147483647) {
                                                                case 9223372036854775807:
                                                                    var v147 uint = v133
                                                                    _ = v147
                                                                }
                                                                if true {
                                                                    var v148 uint16 = (v2 >> (((v133 / (((^(v133))) | 1))) & 63))
                                                                    _ = v148
                                                                } else {
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        v149 := GenS12{F0: 100, F1: fi132}
                                                                        _ = v149
                                                                    case "":
                                                                        switch (fi119 << ((v133) & 63)) {
                                                                        case 915:
                                                                            v150 := []int{fi132, fi119}
                                                                            var v151 float32 = (1e10 / 3.14159)
                                                                            switch (fi119 % (((fi132 * (-(fi119)))) | 1)) {
                                                                            case 0:
                                                                                switch ("foo bar" + "\n\t") {
                                                                                case "a":
                                                                                    v152 := GenS14{F0: (0 < uint8(v15)), F1: -720, F2: (^(v133))}
                                                                                    _ = v152
                                                                                default:
                                                                                    switch "a" {
                                                                                    case "":
                                                                                        fmt.Println(v2, fi119, fi132)
                                                                                    case "\n\t":
                                                                                        fmt.Println(v2, v136, v146)
                                                                                    default:
                                                                                        switch (^(fi119)) {
                                                                                        case 419:
                                                                                            switch (fi119 % (((fi132 ^ fi132)) | 1)) {
                                                                                            case 127:
                                                                                                var v153 int32 = 2147483647
                                                                                                _ = v153
                                                                                            case 894:
                                                                                                _ = len(v150)
                                                                                            }
                                                                                        default:
                                                                                            var v154 int = (fi119 * fi132)
                                                                                            _ = v154
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case 970:
                                                                                var v155 int16 = -32768
                                                                                _ = v155
                                                                            }
                                                                            _ = v150
                                                                            _ = v151
                                                                        case -2147483648:
                                                                            fmt.Println(v2)
                                                                        case 32767:
                                                                            if (v2 >= 0) {
                                                                                var v156 int8 = -128
                                                                                v157 := GenS15{F0: true}
                                                                                _ = v156
                                                                                _ = v157
                                                                            } else {
                                                                                v158 := GenS16{F0: -32768}
                                                                                var v159 float64 = (1e10 - 3.14159)
                                                                                var v160 uint32 = (v15 * 2147483647)
                                                                                _ = v158
                                                                                _ = v159
                                                                                _ = v160
                                                                            }
                                                                        }
                                                                    case "a":
                                                                        if false {
                                                                            if false {
                                                                                if true {
                                                                                    switch (fi132 ^ fi119) {
                                                                                    case 4294967295:
                                                                                        v161 := GenS17{F0: (1 != uint32(v2))}
                                                                                        _ = v161
                                                                                    default:
                                                                                        var v162 uint8 = 178
                                                                                        _ = v162
                                                                                    }
                                                                                    var v163 float32 = 1e10
                                                                                    v15 -= 372
                                                                                    _ = v163
                                                                                }
                                                                                var v164 int32 = -678
                                                                                _ = v164
                                                                            }
                                                                            if false {
                                                                                if false {
                                                                                    if true {
                                                                                        if false {
                                                                                            v165 := GenS2{F0: 2, F1: 100, F2: 1}
                                                                                            _ = v165
                                                                                        } else {
                                                                                            v15 = 2147483647
                                                                                            fmt.Println(v3)
                                                                                            v166 := []int{(fi132 / ((genFn0(v2)) | 1)), fi132, fi132}
                                                                                            _ = v166
                                                                                        }
                                                                                    }
                                                                                }
                                                                                v167 := GenS18{F0: (^(v3)), F1: v136}
                                                                                var v168 float32 = 1e10
                                                                                v169 := []int{-2147483648, genFn0((v2 % ((v2) | 1))), -732, (fi132 - genFn0(2))}
                                                                                _ = v167
                                                                                _ = v168
                                                                                _ = v169
                                                                            } else {
                                                                                var v170 int = fi119
                                                                                _ = v170
                                                                            }
                                                                            var v171 int = (fi119 + (fi119 >> ((uint(v3)) & 63)))
                                                                            _ = v171
                                                                        }
                                                                    }
                                                                    if (!true) {
                                                                        v15 -= v15
                                                                        v3 = v3
                                                                        v133 = v133
                                                                    } else {
                                                                        v18 -= -9223372036854775808
                                                                        if true {
                                                                            var v172 uint32 = (^(v15))
                                                                            _ = v172
                                                                        }
                                                                    }
                                                                    v173 := []int{535, fi119, (fi119 ^ 843), fi132}
                                                                    v136 -= v136
                                                                    _ = v173
                                                                }
                                                                fmt.Println(v1, v2)
                                                                v1 = (v146 + (-(v146)))
                                                                if false {
                                                                    var v174 int8 = -74
                                                                    _ = v174
                                                                } else {
                                                                    v2 = (v2 & 945)
                                                                    var v175 float32 = 3.14159
                                                                    v2 = 967
                                                                    _ = v175
                                                                }
                                                                _ = v146
                                                            }
                                                            var v176 int32 = int32(v3)
                                                            if false {
                                                                if false {
                                                                    v136 = int64(v133)
                                                                }
                                                                var v177 int16 = v1
                                                                var v178 uint64 = 774
                                                                _ = v177
                                                                _ = v178
                                                            } else {
                                                                v179 := []int{(fi132 - fi132), (fi132 | (fi119 | (^(fi132)))), fi119, fi132}
                                                                _ = v179
                                                            }
                                                            _ = v176
                                                        } else {
                                                            var v180 uint32 = (v15 + 2)
                                                            _ = v180
                                                        }
                                                        if (false && true) {
                                                            var v181 uint32 = 255
                                                            _ = v181
                                                        } else {
                                                            var v182 uint16 = (v2 * uint16(v3))
                                                            _ = v182
                                                        }
                                                        switch "foo bar" {
                                                        case "":
                                                            var v183 int32 = -2147483648
                                                            _ = v183
                                                        case "foo bar":
                                                            v2 = v2
                                                            v184 := GenS14{F0: (!(v2 >= (v2 >> ((4294967295) & 63)))), F1: int32(v133), F2: 553}
                                                            v185 := []int{(fi132 & (fi132 - (fi132 & 501))), (fi132 - fi119), genFn0(v2)}
                                                            if true {
                                                                fmt.Println(v18, fi132)
                                                                v185 = append(v185, fi119)
                                                                v186 := GenS13{F0: ((fi119 + (^(fi132))) <= int(v15)), F1: 42, F2: ((v18 & v18) < (^(v18)))}
                                                                _ = v186
                                                            }
                                                            v18 += -1
                                                            _ = v184
                                                            _ = v185
                                                        case "\n\t":
                                                            v187 := GenS3{F0: (v1 << ((v133) & 63))}
                                                            _ = v187
                                                        }
                                                    }
                                                    v1 += 42
                                                    var v188 uint16 = 288
                                                    _ = v137
                                                    _ = v138
                                                    _ = v188
                                                } else {
                                                    var v189 string = "foo bar"
                                                    _ = v189
                                                }
                                                _ = v136
                                            }
                                            switch ("" + "a") {
                                            case "a":
                                                var v190 string = (("\n\t" + "\n\t") + "")
                                                _ = v190
                                            case "hello":
                                                var v191 string = "foo bar"
                                                _ = v191
                                            }
                                        }
                                        switch -1 {
                                        case 2147483647:
                                            switch "" {
                                            case "foo bar":
                                                v192 := []int{-532, 390}
                                                _ = v192
                                            case "a":
                                                var v193 uint64 = v3
                                                _ = v193
                                            }
                                        case 1:
                                            v1 += v1
                                        case -962:
                                            if (!(("foo bar" + "") <= (("" + ("a" + "")) + ((("hello" + ("hello" + "foo bar")) + ("a" + "hello")) + "hello")))) {
                                                v133 -= v133
                                                if false {
                                                    v194 := []int{42, -726}
                                                    _ = v194
                                                } else {
                                                    v195 := GenS16{F0: v1}
                                                    _ = v195
                                                }
                                            }
                                        default:
                                            switch (fi132 / ((genFn0((v2 % ((32767) | 1)))) | 1)) {
                                            case -744:
                                                if true {
                                                    v18 = v18
                                                    switch (fi132 ^ (fi132 & (fi119 >> ((v133) & 63)))) {
                                                    case 4294967295:
                                                        switch fi119 {
                                                        case 1:
                                                            switch genFn0(323) {
                                                            case 128:
                                                                v196 := GenS2{F0: 42, F1: int8(v3), F2: -9}
                                                                _ = v196
                                                            default:
                                                                switch "a" {
                                                                case "foo bar":
                                                                    var v197 float64 = 3.14159
                                                                    _ = v197
                                                                case "a":
                                                                    v198 := GenS19{F0: 42, F1: 2147483647, F2: (v1 * (^(v1)))}
                                                                    _ = v198
                                                                case "\n\t":
                                                                    v199 := []int{(^(fi132)), genFn0((v2 << (((v133 / ((v133) | 1))) & 63)))}
                                                                    _ = v199
                                                                default:
                                                                    v15 = v15
                                                                }
                                                            }
                                                        case 9223372036854775807:
                                                            if (fi119 >= 741) {
                                                                v15 += 4294967295
                                                                v200 := []int{4294967295, fi132, fi119, (fi132 & 775)}
                                                                if false {
                                                                    var v201 bool = (true && true)
                                                                    if v201 {
                                                                        if len(v200) >= 2 {
                                                                            v200 = v200[1:len(v200)]
                                                                        }
                                                                        var v202 int8 = 100
                                                                        fmt.Println(v2, v3)
                                                                        _ = v202
                                                                    } else {
                                                                        if ((!v201) && v201) {
                                                                            v133 = (v133 << ((65535) & 63))
                                                                            var v203 int8 = -128
                                                                            fmt.Println(v133)
                                                                            _ = v203
                                                                        }
                                                                        v18 = 0
                                                                        if (!v201) {
                                                                            _ = len(v200)
                                                                            if len(v200) >= 2 {
                                                                                v200 = v200[1:len(v200)]
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = len(v200)
                                                                    _ = v201
                                                                }
                                                                _ = v200
                                                            } else {
                                                                if false {
                                                                    v204 := []int{fi119, (fi119 & genFn0(v2)), (fi132 << ((genFn1(2)) & 63)), (fi132 * genFn0((v2 & v2)))}
                                                                    _ = v204
                                                                } else {
                                                                    var v205 uint16 = (v2 * 32767)
                                                                    _ = v205
                                                                }
                                                            }
                                                        }
                                                    }
                                                    v15 = (v15 & (v15 ^ 445))
                                                }
                                            default:
                                                v206 := []int{(fi132 & genFn0(v2)), int(v3), (fi119 * -512), fi119}
                                                v15 -= 1
                                                v18 = v18
                                                _ = v206
                                            }
                                        }
                                    }
                                case "hello":
                                    fmt.Println(v15, v2)
                                }
                            }
                            v15 -= v15
                            _ = fi132
                            _ = v133
                        }
                        _ = fi119
                    }
                    if true {
                        switch genFn0((v2 << ((2) & 63))) {
                        case 2147483647:
                            v207 := []int{genFn0(v2), int(v2), 127, 32768}
                            _ = v207
                        case 42:
                            if (v2 < (v2 << ((4294967295) & 63))) {
                                var v208 int64 = v18
                                fmt.Println(v1)
                                _ = v208
                            } else {
                                switch ("foo bar" + "") {
                                case "a":
                                    switch 127 {
                                    case 32768:
                                        v18 -= 9223372036854775807
                                    case -2147483648:
                                        v209 := []int{338}
                                        _ = v209
                                    }
                                case "foo bar":
                                    if (249 <= 2) {
                                        for fi210 := 0; fi210 < 20; fi210++ {
                                            fmt.Println(v2, v1, v2)
                                            for fi211 := 0; fi211 < 17; fi211++ {
                                                v212 := []int{127}
                                                _ = fi211
                                                _ = v212
                                            }
                                            v3 = (v3 << ((65535) & 63))
                                            _ = fi210
                                        }
                                    }
                                default:
                                    v213 := []int{32767}
                                    _ = v213
                                }
                                var v214 int16 = (v1 & v1)
                                switch 42 {
                                case -9223372036854775808:
                                    var v215 int64 = (v18 % ((int64(v3)) | 1))
                                    _ = v215
                                case 2147483647:
                                    fmt.Println(v2, v18, v15)
                                }
                                _ = v214
                            }
                        }
                    }
                } else {
                    var v216 int8 = -128
                    fmt.Println(v216)
                    switch "foo bar" {
                    case "foo bar":
                        var v217 uint16 = v2
                        _ = v217
                    case "":
                        var v218 int = int(v3)
                        _ = v218
                    }
                    _ = v216
                }
                for fi219 := 0; fi219 < 9; fi219++ {
                    fmt.Println(v3, v2)
                    v18 += 1
                    for fi220 := 0; fi220 < 13; fi220++ {
                        var v221 bool = true
                        if false {
                            var v222 uint8 = 1
                            if v221 {
                                v223 := []int{genFn0((v2 ^ (v2 << ((genFn1((v3 - 4294967295))) & 63))))}
                                if (v221 && v221) {
                                    var v224 float32 = (3.14159 - 0.0)
                                    _ = v224
                                }
                                if len(v223) > 0 {
                                    v223[0] = fi219
                                }
                                _ = v223
                            } else {
                                var v225 uint = 18446744073709551615
                                _ = v225
                            }
                            switch 127 {
                            case 0:
                                if (fi219 <= genFn0(v2)) {
                                    fmt.Println(v15)
                                    v221 = ((v1 != 862) || (3.14159 > (1.5 / (((float32(3.14159) - (0.1 - (((0.1 * (1.5 + 3.14159)) / float32(3.14159)) * 1.5))) - ((-2.25 + (3.14159 * 0.0)) - float32(1e10))) * 0.1))))
                                }
                            case 235:
                                switch fi220 {
                                case 128:
                                    var v226 uint = 805
                                    _ = v226
                                default:
                                    var v227 int64 = v18
                                    _ = v227
                                }
                            }
                            _ = v222
                        }
                        switch ("a" + "\n\t") {
                        case "foo bar":
                            if (v221 || v221) {
                                v18 = 1
                                v1 = (v1 / ((v1) | 1))
                                var v228 uint = genFn1(v3)
                                _ = v228
                            }
                        case "a":
                            v229 := GenS6{F0: v2}
                            _ = v229
                        case "\n\t":
                            v230 := GenS8{F0: 1, F1: (^(v2))}
                            _ = v230
                        }
                        _ = fi220
                        _ = v221
                    }
                    _ = fi219
                }
            }
        case 32767:
            var v231 int32 = -1
            _ = v231
        }
    case 32768:
        v232 := []int{-29}
        _ = v232
    case -1:
        v15 = 2147483648
    default:
        v2 -= (^(v2))
    }
    for fi233 := 0; fi233 < 18; fi233++ {
        v234 := []int{fi233}
        var v235 uint8 = 1
        _ = fi233
        _ = v234
        _ = v235
    }
    v18 -= v18
    v2 = uint16(v15)
    switch 128 {
    case 4294967295:
        for fi236 := 0; fi236 < 3; fi236++ {
            v237 := GenS10{F0: -9223372036854775808, F1: ""}
            _ = fi236
            _ = v237
        }
    case -897:
        for fi238 := 0; fi238 < 10; fi238++ {
            switch (("\n\t" + "a") + ((("\n\t" + "hello") + "hello") + "\n\t")) {
            case "a":
                fmt.Println(v15, v3)
            }
            var v239 uint16 = 720
            var v240 uint32 = (v15 << ((genFn1((v3 & (^(v3))))) & 63))
            _ = fi238
            _ = v239
            _ = v240
        }
    default:
        if ("foo bar" != "\n\t") {
            var v241 int16 = v1
            v1 = (^(v241))
            _ = v241
        }
    }
    if false {
        v242 := []int{128, genFn0(v2)}
        var v243 bool = true
        _ = v242
        _ = v243
    } else {
        for fi244 := 0; fi244 < 19; fi244++ {
            fmt.Println(v18)
            v1 = (v1 + v1)
            _ = fi244
        }
        var v245 string = (("foo bar" + "") + ("" + ""))
        _ = v245
    }
    fmt.Println(v15, v2)
    v246 := []int{-2147483648, -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v246:", len(v246), v246)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v246
    _ = v3
}
