package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 int32; F1 int8; F2 uint }
type GenS6 struct { F0 bool; F1 int32; F2 uint64 }
type GenS7 struct { F0 string; F1 int }
type GenS8 struct { F0 uint32; F1 int; F2 int64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint; F1 uint32 }
type GenS11 struct { F0 int16; F1 uint32 }
type GenS12 struct { F0 int; F1 int16; F2 int8 }
func genFn0(p0 uint32, p1 uint8, p2 int8) uint8 {
    v5 := []int{}
    if (41 != 1) {
        var v6 int64 = 1
        var v7 uint64 = 852
        _ = len(v5)
        _ = v6
        _ = v7
    } else {
        if false {
            if false {
                for fi8 := 0; fi8 < 7; fi8++ {
                    p0 = 65535
                    _ = fi8
                }
                if (true || (!((-1 == -32768) || ((!true) && (true && (false && false)))))) {
                    if ((830 <= 0) && (((-(p2)) <= p2) && (-2147483648 != -2147483648))) {
                        var v9 int64 = 256
                        for fi10 := 0; fi10 < 14; fi10++ {
                            var v11 string = ""
                            _ = fi10
                            _ = v11
                        }
                        var v12 bool = (0 >= 0)
                        _ = v12
                        _ = v9
                    } else {
                        p1 = p1
                    }
                    for fi13 := 0; fi13 < 17; fi13++ {
                        v14 := GenS1{F0: 561, F1: -117}
                        p1 -= 1
                        _ = len(v5)
                        _ = fi13
                        _ = v14
                    }
                }
                var v15 uint32 = (p0 << ((42) & 63))
                _ = v15
            } else {
                if (!false) {
                    var v16 string = "foo bar"
                    var v17 int16 = -1
                    _ = v16
                    _ = v17
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            p0 = p0
        } else {
            var v18 int8 = 0
            _ = v18
        }
        var v19 int8 = p2
        if ((2 >= 42) || ((42 != p0) && true)) {
            v5 = append(v5, -315)
        } else {
            switch "" {
            case "foo bar":
                if len(v5) > 0 {
                    v5[0] = -9223372036854775808
                }
            case "hello":
                for fi20 := 0; fi20 < 15; fi20++ {
                    var v21 uint32 = 4294967295
                    if false {
                        var v22 uint64 = 0
                        _ = v22
                    }
                    v23 := GenS2{F0: int32(fi20), F1: 4294967295, F2: 901}
                    _ = fi20
                    _ = v21
                    _ = v23
                }
            case "\n\t":
                if true {
                    var v24 string = (("" + "hello") + "\n\t")
                    for fi25 := 0; fi25 < 3; fi25++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi25
                    }
                    _ = v24
                }
            }
            fmt.Println(p1, p2)
        }
        _ = v19
    }
    var v26 float64 = (float64(0.1) - 1e10)
    var v27 int = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    _ = v5
    return p1
}

func main() {
    var v1 uint = 2
    var v2 int8 = 127
    var v3 uint = (v1 >> (((v1 - v1)) & 63))
    v4 := make(map[string]int)
    switch 179 {
    case 830:
        v1 += (v3 / (((v1 & (^(v3)))) | 1))
    case 65:
        _ = len(v4)
    default:
        switch 4294967295 {
        case -597:
            for fi28 := 0; fi28 < 20; fi28++ {
                for fi29 := 0; fi29 < 11; fi29++ {
                    v2 -= v2
                    v30 := []int{int(v1), fi29, 2147483647}
                    _ = fi29
                    _ = v30
                }
                for fi31 := 0; fi31 < 4; fi31++ {
                    v32 := []int{fi31, (fi28 / ((2147483647) | 1)), 1, 4294967295}
                    var v33 uint8 = 184
                    v2 = 127
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                _ = fi28
            }
        }
    }
    v3 -= (^(v3))
    var v34 int64 = 9223372036854775807
    v2 -= -12
    for fi35 := 0; fi35 < 7; fi35++ {
        switch (("a" + "hello") + "a") {
        case "foo bar":
            v36 := GenS2{F0: -1, F1: 495, F2: 1}
            _ = v36
        case "":
            for fi37 := 0; fi37 < 16; fi37++ {
                if false {
                    v38 := GenS3{F0: 128, F1: 128}
                    v39 := GenS4{F0: 42, F1: -361}
                    v2 = v2
                    fmt.Println(fi37, fi37, v34)
                    var v40 int = fi35
                    _ = v38
                    _ = v39
                    _ = v40
                }
                v41 := GenS2{F0: -1, F1: 0, F2: v3}
                _ = fi37
                _ = v41
            }
        case "a":
            for fi42 := 0; fi42 < 16; fi42++ {
                v43 := GenS1{F0: 2, F1: (^(v34))}
                _ = fi42
                _ = v43
            }
        }
        v3 -= 255
        for fi44 := 0; fi44 < 10; fi44++ {
            v45 := GenS1{F0: 644, F1: 0}
            _ = fi44
            _ = v45
        }
        _ = fi35
    }
    v46 := GenS4{F0: 42, F1: -767}
    if false {
        var v47 uint64 = 683
        _ = v47
    } else {
        v48 := []int{32768, -648}
        v3 -= v3
        if (("" + "hello") >= ("hello" + ("" + ((("foo bar" + "hello") + "") + "")))) {
            v1 = 0
            for fi49 := 0; fi49 < 12; fi49++ {
                v34 -= v34
                var v50 string = "hello"
                var v51 int8 = (v2 % ((v2) | 1))
                _ = fi49
                _ = v50
                _ = v51
            }
            var v52 int = 127
            _ = v52
        }
        _ = v48
    }
    switch -32 {
    case 2147483647:
        v34 += v34
    }
    if (true && true) {
        if true {
            for fi53 := 0; fi53 < 10; fi53++ {
                for fi54 := 0; fi54 < 20; fi54++ {
                    var v55 float32 = 0.1
                    if true {
                        v34 = 0
                        v56 := GenS5{F0: 0, F1: v2, F2: 255}
                        _ = v56
                    } else {
                        if false {
                            v2 = (v2 ^ (v2 >> ((v1) & 63)))
                        } else {
                            v57 := GenS6{F0: (false || (v55 <= 0.1)), F1: 897, F2: 910}
                            v58 := GenS5{F0: 42, F1: v2, F2: (v3 & v3)}
                            if (!false) {
                                if (1 >= 18446744073709551615) {
                                    v55 = (v55 / v55)
                                } else {
                                    v1 = (v1 + (v1 * v3))
                                }
                            }
                            _ = v57
                            _ = v58
                        }
                        v34 = int64(v2)
                        fmt.Println(v3)
                    }
                    v59 := GenS7{F0: ("\n\t" + "\n\t"), F1: (-(fi54))}
                    if false {
                        switch "" {
                        case "\n\t":
                            var v60 int16 = -32768
                            _ = v60
                        case "":
                            v34 = (^(v34))
                            switch (fi53 & (fi53 << ((2) & 63))) {
                            case 671:
                                fmt.Println(fi53, fi53)
                            case 32768:
                                v2 += v2
                            case -94:
                                fmt.Println(v55, fi54)
                            }
                            v61 := GenS5{F0: 42, F1: (v2 & v2), F2: v3}
                            if true {
                                v62 := []int{fi54}
                                _ = v62
                            }
                            var v63 string = ("foo bar" + ("a" + (("" + ("a" + "\n\t")) + ("" + ("" + "")))))
                            if true {
                                v64 := GenS8{F0: uint32(v1), F1: (fi54 ^ fi54), F2: (v34 * v34)}
                                v34 -= (-(v34))
                                v2 = (v2 << ((4294967295) & 63))
                                _ = v64
                            } else {
                                var v65 int32 = 1
                                var v66 bool = true
                                v67 := GenS3{F0: 0, F1: 1}
                                _ = v65
                                _ = v66
                                _ = v67
                            }
                            _ = v61
                            _ = v63
                        }
                        v1 -= (v3 + (v1 - 65535))
                        v2 = (v2 ^ (v2 / ((v2) | 1)))
                    }
                    v2 = v2
                    _ = fi54
                    _ = v55
                    _ = v59
                }
                _ = fi53
            }
            var v68 float64 = (0.1 + 0.1)
            _ = v68
        } else {
            v69 := GenS8{F0: 4294967295, F1: -9223372036854775808, F2: (v34 * (v34 % ((v34) | 1)))}
            v2 = (v2 % ((v2) | 1))
            for fi70 := 0; fi70 < 20; fi70++ {
                v3 -= v3
                for fi71 := 0; fi71 < 9; fi71++ {
                    fmt.Println(v2, fi70)
                    v1 = 4294967295
                    switch (fi71 | 4294967295) {
                    case 0:
                        var v72 float32 = (-2.25 / ((-2.25 / 1e10) / 0.1))
                        _ = v72
                    }
                    switch "foo bar" {
                    case "\n\t":
                        var v73 bool = (false && (v2 < (v2 + (v2 / (((v2 ^ v2)) | 1)))))
                        _ = v73
                    }
                    _ = fi71
                }
                _ = fi70
            }
            _ = v69
        }
        switch 42 {
        case 9223372036854775807:
            v4["hello"] = 447
        }
    }
    v74 := GenS7{F0: ("\n\t" + (("\n\t" + "hello") + (("\n\t" + "foo bar") + ("hello" + "foo bar")))), F1: -1}
    for fi75 := 0; fi75 < 2; fi75++ {
        var v76 int8 = v2
        var v77 uint16 = 42
        _ = fi75
        _ = v76
        _ = v77
    }
    if true {
        delete(v4, ("" + "foo bar"))
    } else {
        if (v34 < (^(v34))) {
            if (v34 != v34) {
                if false {
                    if true {
                        var v78 float64 = 1e10
                        for fi79 := 0; fi79 < 5; fi79++ {
                            fmt.Println(v1, v78, v1)
                            var v80 int16 = 1
                            _ = fi79
                            _ = v80
                        }
                        if (324 != 2) {
                            v81 := []int{127, 32768}
                            delete(v4, ("\n\t" + "foo bar"))
                            for fi82 := 0; fi82 < 11; fi82++ {
                                v2 = v2
                                _ = fi82
                            }
                            _ = v81
                        } else {
                            v2 += -1
                            if false {
                                v83 := GenS2{F0: -743, F1: 42, F2: (v3 % ((255) | 1))}
                                _ = v83
                            }
                            switch ("foo bar" + ("hello" + "a")) {
                            case "a":
                                for fi84 := 0; fi84 < 4; fi84++ {
                                    v85 := []int{(^(fi84)), (fi84 - fi84)}
                                    for fi86 := 0; fi86 < 20; fi86++ {
                                        if ((v34 % (((v34 >> ((1) & 63))) | 1)) == (v34 | (v34 >> (((v3 >> ((v1) & 63))) & 63)))) {
                                            if true {
                                                fmt.Println(v1, fi86)
                                            } else {
                                                if len(v85) > 0 {
                                                    v85[0] = fi84
                                                }
                                                var v87 int8 = v2
                                                switch (("foo bar" + "") + "a") {
                                                case "foo bar":
                                                    switch "foo bar" {
                                                    case "hello":
                                                        fmt.Println(v1, v87, v3)
                                                    }
                                                case "\n\t":
                                                    v1 = 255
                                                case "hello":
                                                    var v88 int = (^(fi86))
                                                    _ = v88
                                                default:
                                                    var v89 int64 = v34
                                                    _ = v89
                                                }
                                                _ = v87
                                            }
                                        }
                                        _ = fi86
                                    }
                                    for fi90 := 0; fi90 < 2; fi90++ {
                                        var v91 uint64 = 965
                                        _ = fi90
                                        _ = v91
                                    }
                                    _ = fi84
                                    _ = v85
                                }
                            case "foo bar":
                                fmt.Println(v2)
                            default:
                                var v92 int16 = 1000
                                _ = v92
                            }
                        }
                        _ = v78
                    }
                }
            } else {
                v34 = v34
                if true {
                    var v93 uint8 = 127
                    for fi94 := 0; fi94 < 5; fi94++ {
                        var v95 uint64 = 18446744073709551615
                        _ = fi94
                        _ = v95
                    }
                    _ = v93
                }
                v96 := GenS7{F0: "a", F1: 127}
                _ = v96
            }
            for fi97 := 0; fi97 < 11; fi97++ {
                for fi98 := 0; fi98 < 2; fi98++ {
                    if (!false) {
                        var v99 uint32 = 0
                        switch "" {
                        case "foo bar":
                            if (!((v99 % ((v99) | 1)) != (v99 >> (((v1 * (v1 - v3))) & 63)))) {
                                v100 := GenS9{F0: 2}
                                switch fi97 {
                                case 9223372036854775807:
                                    var v101 uint32 = v99
                                    _ = v101
                                default:
                                    fmt.Println(v2)
                                }
                                v102 := GenS1{F0: 831, F1: (-(v34))}
                                switch (-(fi98)) {
                                case -1:
                                    var v103 int64 = -9223372036854775808
                                    _ = v103
                                case 379:
                                    v3 -= 2
                                case 42:
                                    v2 -= (v2 * v2)
                                default:
                                    fmt.Println(v99)
                                    v104 := GenS1{F0: 2, F1: (v34 / (((v34 ^ -9223372036854775808)) | 1))}
                                    _ = v104
                                }
                                _ = v100
                                _ = v102
                            } else {
                                var v105 int8 = 64
                                var v106 uint8 = 128
                                _ = v105
                                _ = v106
                            }
                        default:
                            v99 = v99
                        }
                        var v107 uint16 = 32767
                        _ = v107
                        _ = v99
                    }
                    if true {
                        fmt.Println(v1)
                    } else {
                        fmt.Println(v34, v2, fi97)
                        if true {
                            if false {
                                if true {
                                    v108 := []int{}
                                    _ = v108
                                } else {
                                    if (float32(-2.25) > 0.0) {
                                        fmt.Println(v34)
                                        var v109 uint8 = 2
                                        _ = v109
                                    }
                                    v34 = (-(v34))
                                    var v110 int32 = 2147483647
                                    _ = v110
                                }
                                var v111 int = (fi97 / (((fi97 / ((fi98) | 1))) | 1))
                                _ = v111
                            } else {
                                var v112 int32 = 42
                                var v113 int32 = (-(v112))
                                _ = v112
                                _ = v113
                            }
                            if (genFn0(42, 98, v2) != 1) {
                                var v114 int32 = -2147483648
                                fmt.Println(fi97, v3)
                                _ = v114
                            } else {
                                switch "a" {
                                case "foo bar":
                                    var v115 float32 = 1.5
                                    _ = v115
                                case "hello":
                                    if false {
                                        if (65535 < 0) {
                                            if false {
                                                v116 := GenS10{F0: (^(v1)), F1: 2147483648}
                                                var v117 uint64 = 4294967295
                                                _ = v116
                                                _ = v117
                                            } else {
                                                var v118 float32 = 0.0
                                                if (2 <= (^(v1))) {
                                                    v34 = -9223372036854775808
                                                    if ((!true) || ((!true) || false)) {
                                                        var v119 int = (fi98 ^ 9223372036854775807)
                                                        if false {
                                                            v120 := GenS8{F0: 38, F1: fi98, F2: 42}
                                                            v34 += (v34 + v34)
                                                            var v121 int8 = (v2 % ((-128) | 1))
                                                            _ = v120
                                                            _ = v121
                                                        } else {
                                                            switch 886 {
                                                            case 42:
                                                                if false {
                                                                    v122 := []int{}
                                                                    _ = len(v122)
                                                                    var v123 uint8 = 1
                                                                    if true {
                                                                        switch fi98 {
                                                                        case 32767:
                                                                            var v124 int64 = v34
                                                                            _ = v124
                                                                        case 9223372036854775807:
                                                                            if len(v122) >= 2 {
                                                                                v122 = v122[1:len(v122)]
                                                                            }
                                                                        default:
                                                                            if true {
                                                                                var v125 uint16 = 32767
                                                                                v126 := GenS1{F0: (v125 & v125), F1: v34}
                                                                                if (4294967295 >= 42) {
                                                                                    if len(v122) > 0 {
                                                                                        v122[0] = fi98
                                                                                    }
                                                                                } else {
                                                                                    v127 := GenS2{F0: 2147483647, F1: 577, F2: v3}
                                                                                    _ = v127
                                                                                }
                                                                                _ = v125
                                                                                _ = v126
                                                                            } else {
                                                                                v1 -= (v3 | v1)
                                                                                delete(v4, "a")
                                                                                v119 = (fi97 + (v119 >> ((v3) & 63)))
                                                                            }
                                                                        }
                                                                        fmt.Println(v1, v119)
                                                                        var v128 string = "hello"
                                                                        _ = v128
                                                                    }
                                                                    _ = v122
                                                                    _ = v123
                                                                }
                                                            case 1:
                                                                var v129 uint32 = 53
                                                                _ = v129
                                                            }
                                                        }
                                                        _ = v119
                                                    } else {
                                                        v130 := []int{2147483647, fi98, fi97, (fi97 / ((fi97) | 1))}
                                                        _ = len(v4)
                                                        var v131 bool = ((v2 & -128) > (v2 | (v2 >> (((v1 - (v3 | (v3 & 2)))) & 63))))
                                                        if ((v2 | (v2 | v2)) <= v2) {
                                                            v4[(("hello" + ("foo bar" + ("\n\t" + "foo bar"))) + (("hello" + ("foo bar" + ((((("\n\t" + "foo bar") + "") + ("foo bar" + ("foo bar" + "foo bar"))) + "\n\t") + "a"))) + "hello"))] = (-(fi97))
                                                        } else {
                                                            switch "a" {
                                                            case "hello":
                                                                v2 = (^(v2))
                                                            default:
                                                                if v131 {
                                                                    v130 = append(v130, (fi97 - fi98))
                                                                    fmt.Println(fi97, v34, v131)
                                                                    var v132 uint64 = 0
                                                                    _ = v132
                                                                } else {
                                                                    fmt.Println(v3, v34, v1)
                                                                }
                                                            }
                                                        }
                                                        var v133 int = (fi97 & (fi98 / (((fi98 + (fi97 >> ((v1) & 63)))) | 1)))
                                                        _ = v130
                                                        _ = v131
                                                        _ = v133
                                                    }
                                                }
                                                _ = v118
                                            }
                                        } else {
                                            v134 := GenS11{F0: 42, F1: 2147483648}
                                            var v135 uint = v3
                                            _ = v134
                                            _ = v135
                                        }
                                    }
                                }
                            }
                            _ = len(v4)
                        } else {
                            v1 = (^(v3))
                            v136 := []int{(fi98 + (fi98 >> ((0) & 63))), fi98, fi97}
                            _ = v136
                        }
                    }
                    switch (fi98 % (((-(fi97))) | 1)) {
                    case 42:
                        if (0 != genFn0(2147483647, 47, v2)) {
                            fmt.Println(v3)
                            v137 := []int{fi97, fi97}
                            v34 = (v34 * v34)
                            _ = v137
                        }
                    }
                    _ = fi98
                }
                _ = fi97
            }
        } else {
            v3 = 18446744073709551615
            var v138 uint8 = uint8(v3)
            var v139 float64 = (1.5 + (-2.25 / (1.5 / 1.5)))
            _ = v138
            _ = v139
        }
        if (42 > 4294967295) {
            switch ("foo bar" + "a") {
            case "a":
                if true {
                    fmt.Println(v2, v3)
                    v140 := GenS7{F0: "a", F1: 4294967295}
                    _ = v140
                } else {
                    v3 += v1
                }
            case "foo bar":
                for fi141 := 0; fi141 < 4; fi141++ {
                    var v142 int = fi141
                    _ = len(v4)
                    var v143 int8 = (v2 ^ (-(v2)))
                    _ = fi141
                    _ = v142
                    _ = v143
                }
            case "":
                v34 = int64(v1)
            }
            for fi144 := 0; fi144 < 13; fi144++ {
                fmt.Println(fi144)
                _ = fi144
            }
        }
    }
    switch -44 {
    case -1:
        if ((!(!false)) || true) {
            fmt.Println(v34, v1)
        } else {
            if true {
                if (32768 >= 32767) {
                    if (733 == v34) {
                        v145 := GenS12{F0: 1, F1: -715, F2: 127}
                        _ = v145
                    } else {
                        var v146 float64 = 1e10
                        var v147 string = ("foo bar" + ("hello" + "hello"))
                        _ = v146
                        _ = v147
                    }
                    if ((^(v2)) != v2) {
                        var v148 uint32 = 791
                        var v149 float64 = 1e10
                        switch -9223372036854775808 {
                        case 4294967295:
                            var v150 int8 = -99
                            _ = v150
                        default:
                            var v151 bool = (223 < 225)
                            _ = v151
                        }
                        _ = v148
                        _ = v149
                    }
                    fmt.Println(v34, v2, v2)
                } else {
                    var v152 int64 = (v34 / (((-(v34))) | 1))
                    v153 := GenS11{F0: -32768, F1: 42}
                    _ = v152
                    _ = v153
                }
                v154 := v4["\n\t"]
                _ = v154
            } else {
                v2 -= v2
            }
            var v155 uint8 = 128
            _ = v155
        }
    case 32768:
        v2 = (v2 - -53)
    default:
        v156 := []int{2147483647}
        _ = v156
    }
    fmt.Println(v2, v34)
    for fi157 := 0; fi157 < 14; fi157++ {
        if true {
            v158 := []int{(fi157 + fi157), int(v2), (fi157 | (-(fi157)))}
            v1 = v3
            if true {
                fmt.Println(v1, v2)
            }
            _ = v158
        } else {
            v3 += (^(v3))
        }
        var v159 int8 = v2
        _ = fi157
        _ = v159
    }
    var v160 uint = v1
    fmt.Println("v1:", v1)
    fmt.Println("v160:", v160)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v160
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v46
    _ = v74
}
