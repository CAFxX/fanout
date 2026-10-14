package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint; F1 int; F2 uint16 }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 int; F1 uint32; F2 int64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int; F1 uint32; F2 int16 }
type GenS8 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int32; F1 uint32 }
type GenS11 struct { F0 uint64; F1 int }
type GenS12 struct { F0 int16; F1 uint16 }
func genFn0() uint64 {
    if true {
        var v7 int16 = 1
        if (!(4294967295 >= 680)) {
            var v8 float32 = float32(-2.25)
            if (!true) {
                var v9 uint32 = 2
                var v10 uint32 = (v9 | v9)
                v11 := GenS1{F0: false}
                _ = v10
                _ = v11
                _ = v9
            }
            _ = v8
        } else {
            v12 := []int{32767, 4294967295, 881}
            for fi13 := 0; fi13 < 18; fi13++ {
                var v14 int16 = 1
                if (true || (false || true)) {
                    v15 := GenS1{F0: (((0.0 == 1.5) && false) || (18446744073709551615 == 42))}
                    _ = v15
                }
                _ = fi13
                _ = v14
            }
            _ = v12
        }
        _ = v7
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 int16 = 622
        v18 := []int{1, (^(fi16)), fi16}
        _ = fi16
        _ = v17
        _ = v18
    }
    return 1
}
func genFn1() uint8 {
    if true {
        var v19 int64 = 0
        v20 := GenS2{F0: (v19 / (((v19 % ((-9223372036854775808) | 1))) | 1)), F1: 42}
        _ = v19
        _ = v20
    } else {
        var v21 uint8 = 108
        for fi22 := 0; fi22 < 6; fi22++ {
            switch ("" + "a") {
            case "foo bar":
                switch "\n\t" {
                case "foo bar":
                    for fi23 := 0; fi23 < 15; fi23++ {
                        var v24 int16 = -1
                        v25 := []int{fi22, fi22, (^(fi23))}
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                case "":
                    v26 := []int{(^(fi22)), fi22, 127, (fi22 - 42)}
                    _ = v26
                }
            case "":
                switch "hello" {
                case "foo bar":
                    v27 := []int{fi22, (fi22 << ((18446744073709551615) & 63))}
                    _ = v27
                default:
                    var v28 string = (("a" + (("\n\t" + "") + "a")) + ("a" + ""))
                    _ = v28
                }
            case "\n\t":
                var v29 uint = 1
                _ = v29
            default:
                for fi30 := 0; fi30 < 14; fi30++ {
                    var v31 int16 = -577
                    _ = fi30
                    _ = v31
                }
            }
            _ = fi22
        }
        for fi32 := 0; fi32 < 20; fi32++ {
            v33 := []int{(fi32 + -942)}
            _ = fi32
            _ = v33
        }
        _ = v21
    }
    for fi34 := 0; fi34 < 16; fi34++ {
        fmt.Println(fi34)
        _ = fi34
    }
    var v35 int16 = -985
    v35 -= v35
    _ = v35
    return 230
}

func main() {
    var v1 uint8 = 128
    var v2 uint64 = 829
    var v3 float64 = 0.0
    var v4 int8 = 49
    var v5 string = "foo bar"
    v6 := make(map[string]bool)
    if (v4 > -128) {
        delete(v6, v5)
        var v36 string = v5
        switch (v5 + v5) {
        case "a":
            v37 := GenS2{F0: -1, F1: -1}
            _ = v37
        default:
            if true {
                v3 = v3
            }
        }
        _ = v36
    } else {
        if true {
            switch 128 {
            case 9223372036854775807:
                if true {
                    v4 -= 85
                    fmt.Println(v4, v4, v1)
                } else {
                    v5 = (v5 + ("a" + "hello"))
                }
            case 128:
                switch v5 {
                case "foo bar":
                    var v38 float64 = 1.5
                    _ = v38
                case "":
                    v3 = float64(0.1)
                case "hello":
                    _ = len(v6)
                }
            case -357:
                var v39 int8 = (^(v4))
                _ = v39
            default:
                for fi40 := 0; fi40 < 8; fi40++ {
                    switch fi40 {
                    case -220:
                        v41 := GenS1{F0: (1000 < -32768)}
                        _ = v41
                    case 42:
                        if (false && true) {
                            v42 := GenS3{F0: 18446744073709551615, F1: (fi40 / ((128) | 1)), F2: 582}
                            _ = v42
                        } else {
                            switch fi40 {
                            case 127:
                                var v43 int32 = 2147483647
                                _ = v43
                            case 804:
                                v44 := GenS4{F0: v5, F1: 735}
                                _ = v44
                            }
                            v3 = (0.1 * 0.0)
                            switch ("\n\t" + "") {
                            case "hello":
                                var v45 float64 = v3
                                _ = v45
                            case "foo bar":
                                if false {
                                    var v46 int16 = 32767
                                    _ = v46
                                } else {
                                    var v47 float32 = float32(1.5)
                                    v3 += v3
                                    switch v5 {
                                    case "a":
                                        if (0 >= v4) {
                                            if false {
                                                fmt.Println(v4, v4, v4)
                                                v4 += v4
                                                v47 += v47
                                            } else {
                                                v48 := []int{}
                                                _ = v48
                                            }
                                            var v49 int8 = (v4 - (v4 | (v4 / ((v4) | 1))))
                                            _ = v49
                                        }
                                    default:
                                        fmt.Println(v2)
                                    }
                                    _ = v47
                                }
                            }
                        }
                    case -1:
                        v3 += (1.5 / float64((float32(-2.25) / 0.1)))
                    default:
                        if (false && true) {
                            var v50 int32 = 250
                            _ = v50
                        }
                    }
                    switch v5 {
                    case "":
                        v2 = (v2 * (v2 & (v2 ^ (v2 * v2))))
                    default:
                        fmt.Println(v5, v5, v3)
                    }
                    _ = fi40
                }
            }
            for fi51 := 0; fi51 < 3; fi51++ {
                switch (v5 + v5) {
                case "a":
                    var v52 float64 = float64((float32(float64(0.0)) * (0.1 / -2.25)))
                    _ = v52
                case "":
                    switch ("foo bar" + (v5 + "hello")) {
                    case "a":
                        var v53 string = ""
                        _ = v53
                    case "\n\t":
                        switch "foo bar" {
                        case "foo bar":
                            fmt.Println(v3, fi51)
                        default:
                            if true {
                                var v54 float32 = 1.5
                                for fi55 := 0; fi55 < 15; fi55++ {
                                    delete(v6, (v5 + v5))
                                    v56 := []int{9223372036854775807, (fi51 | (fi55 >> ((0) & 63))), 128}
                                    v54 += v54
                                    _ = fi55
                                    _ = v56
                                }
                                v57 := []int{}
                                _ = v54
                                _ = v57
                            }
                        }
                    default:
                        v58 := GenS5{F0: 9223372036854775807, F1: 152, F2: 1}
                        _ = v58
                    }
                }
                var v59 int16 = 511
                for fi60 := 0; fi60 < 3; fi60++ {
                    fmt.Println(v1, v5, v3)
                    var v61 uint32 = 2147483648
                    _ = fi60
                    _ = v61
                }
                _ = fi51
                _ = v59
            }
        }
        fmt.Println(v3, v4, v1)
    }
    v62 := GenS3{F0: 255, F1: 127, F2: 42}
    v63 := GenS6{F0: 480}
    if (42 <= 42) {
        switch "hello" {
        case "foo bar":
            var v64 int8 = 0
            _ = v64
        default:
            if ((!false) || (!(!(v3 != (0.1 * v3))))) {
                for fi65 := 0; fi65 < 15; fi65++ {
                    if false {
                        fmt.Println(fi65)
                    }
                    _ = fi65
                }
                fmt.Println(v3, v5, v2)
            } else {
                v5 = (v5 + v5)
                if true {
                    switch v5 {
                    case "":
                        for fi66 := 0; fi66 < 20; fi66++ {
                            var v67 bool = true
                            v68 := v6[("foo bar" + v5)]
                            v69 := []int{(fi66 + fi66), fi66, 612, (fi66 >> ((uint(v2)) & 63))}
                            _ = fi66
                            _ = v67
                            _ = v68
                            _ = v69
                        }
                    case "\n\t":
                        for fi70 := 0; fi70 < 9; fi70++ {
                            v1 -= v1
                            var v71 uint32 = 1
                            _ = fi70
                            _ = v71
                        }
                    case "hello":
                        v72 := GenS4{F0: v5, F1: 4294967295}
                        _ = v72
                    }
                } else {
                    var v73 int8 = (v4 << ((4294967295) & 63))
                    _ = v73
                }
            }
        }
    }
    var v74 int8 = v4
    for fi75 := 0; fi75 < 17; fi75++ {
        if (2147483648 >= 0) {
            fmt.Println(fi75, v2, v2)
            var v76 string = v5
            _ = v76
        } else {
            _ = len(v6)
            v77 := GenS4{F0: "hello", F1: 24}
            v78 := GenS1{F0: (42 == 1)}
            _ = v77
            _ = v78
        }
        switch "foo bar" {
        case "foo bar":
            var v79 int64 = 42
            _ = v79
        case "":
            var v80 string = (v5 + "hello")
            _ = v80
        case "hello":
            v74 = v74
        default:
            var v81 string = v5
            _ = v81
        }
        if true {
            v74 = v4
            var v82 int = fi75
            _ = v82
        } else {
            if true {
                if ((1 < 0) || true) {
                    if false {
                        for fi83 := 0; fi83 < 8; fi83++ {
                            switch (((v5 + "") + "a") + "foo bar") {
                            case "hello":
                                switch "foo bar" {
                                case "":
                                    var v84 int = fi75
                                    _ = v84
                                }
                            case "\n\t":
                                v85 := []int{int(v1), 2147483647}
                                _ = v85
                            case "a":
                                var v86 int = fi83
                                _ = v86
                            default:
                                v3 = v3
                            }
                            fmt.Println(v2)
                            _ = fi83
                        }
                        var v87 uint8 = (v1 & (v1 | v1))
                        if (-1 < 1) {
                            var v88 int32 = 1
                            var v89 int = fi75
                            for fi90 := 0; fi90 < 11; fi90++ {
                                v74 = (v74 & -128)
                                var v91 int16 = -32768
                                _ = fi90
                                _ = v91
                            }
                            _ = v88
                            _ = v89
                        }
                        _ = v87
                    } else {
                        for fi92 := 0; fi92 < 9; fi92++ {
                            switch fi75 {
                            case 128:
                                var v93 int32 = -722
                                _ = v93
                            }
                            _ = fi92
                        }
                    }
                    for fi94 := 0; fi94 < 8; fi94++ {
                        fmt.Println(v4, v4)
                        if true {
                            fmt.Println(v3, v2)
                            var v95 int = 0
                            switch (v5 + v5) {
                            case "":
                                var v96 int8 = v4
                                _ = v96
                            default:
                                v97 := GenS2{F0: 9223372036854775807, F1: 42}
                                _ = v97
                            }
                            _ = v95
                        } else {
                            v98 := GenS1{F0: false}
                            _ = v98
                        }
                        if (108 <= 0) {
                            v3 = 0.0
                        }
                        _ = fi94
                    }
                }
            } else {
                if ((float64(3.14159) / v3) == v3) {
                    switch (fi75 >> ((0) & 63)) {
                    case 905:
                        v99 := v6[((v5 + "") + v5)]
                        _ = v99
                    case 2147483647:
                        if true {
                            v2 = (v2 ^ (v2 / ((v2) | 1)))
                            for fi100 := 0; fi100 < 1; fi100++ {
                                if false {
                                    v101 := GenS5{F0: fi75, F1: 1, F2: -1}
                                    fmt.Println(v5, v2)
                                    if (v5 != (("a" + v5) + (((v5 + "a") + "") + ""))) {
                                        if true {
                                            fmt.Println(v3, v2)
                                        }
                                        fmt.Println(v3)
                                    }
                                    _ = v101
                                } else {
                                    fmt.Println(v5, v74, v5)
                                    v2 = (v2 / (((v2 * 4294967295)) | 1))
                                }
                                if false {
                                    fmt.Println(v3)
                                    fmt.Println(v3)
                                    if true {
                                        v74 = 0
                                        v102 := GenS7{F0: -2147483648, F1: 2, F2: 1000}
                                        fmt.Println(fi100, v74, v1)
                                        _ = v102
                                    } else {
                                        if ((!true) && true) {
                                            if true {
                                                v103 := GenS1{F0: (fi75 >= (^(fi100)))}
                                                v4 -= (v4 * v74)
                                                _ = v103
                                            } else {
                                                v3 -= 0.1
                                                fmt.Println(v2, v1, v1)
                                            }
                                            var v104 int32 = 1
                                            fmt.Println(v1)
                                            _ = v104
                                        }
                                    }
                                } else {
                                    v3 = (v3 - v3)
                                }
                                _ = fi100
                            }
                            v105 := GenS4{F0: ((("" + v5) + ("hello" + "")) + v5), F1: 255}
                            _ = v105
                        } else {
                            v106 := GenS8{F0: -1, F1: (^(v1)), F2: 255}
                            _ = v106
                        }
                    }
                    _ = len(v6)
                } else {
                    if (false || ((3.14159 * v3) != 0.0)) {
                        var v107 int = (fi75 ^ 2147483647)
                        var v108 int32 = 265
                        for fi109 := 0; fi109 < 1; fi109++ {
                            if (1 > uint32(v1)) {
                                v5 = v5
                                if (18446744073709551615 < 65535) {
                                    v110 := GenS9{F0: v108}
                                    fmt.Println(v107, v107)
                                    v107 = (fi109 / ((-2147483648) | 1))
                                    _ = v110
                                }
                                if (1000 <= 42) {
                                    fmt.Println(v3, v3)
                                    v107 += (fi109 >> ((4294967295) & 63))
                                    var v111 int = v107
                                    _ = v111
                                } else {
                                    switch ("hello" + "a") {
                                    case "":
                                        var v112 float32 = 1e10
                                        _ = v112
                                    case "\n\t":
                                        v113 := v6[v5]
                                        _ = v113
                                    case "foo bar":
                                        var v114 float32 = ((1e10 - -2.25) + 3.14159)
                                        _ = v114
                                    default:
                                        var v115 int32 = (v108 + v108)
                                        _ = v115
                                    }
                                    var v116 int64 = 42
                                    _ = v116
                                }
                            } else {
                                v117 := GenS10{F0: -336, F1: 141}
                                v118 := GenS2{F0: 0, F1: -1}
                                _ = v117
                                _ = v118
                            }
                            _ = fi109
                        }
                        _ = v107
                        _ = v108
                    } else {
                        fmt.Println(v5, v5, v4)
                        for fi119 := 0; fi119 < 8; fi119++ {
                            if false {
                                if true {
                                    if (42 != -362) {
                                        _ = len(v6)
                                        if (v3 >= 1e10) {
                                            delete(v6, "a")
                                            v4 -= (v4 + (v4 - 100))
                                            v4 += 97
                                        }
                                        var v120 int8 = (^(v74))
                                        _ = v120
                                    } else {
                                        if (true || (1 != -1)) {
                                            switch (-(fi119)) {
                                            case -1:
                                                var v121 uint32 = 2
                                                _ = v121
                                            case 2147483647:
                                                v122 := GenS9{F0: 880}
                                                _ = v122
                                            default:
                                                if true {
                                                    v123 := []int{int(v2), fi75}
                                                    _ = v123
                                                } else {
                                                    v5 = (v5 + "a")
                                                }
                                            }
                                            switch "foo bar" {
                                            case "hello":
                                                var v124 int64 = 9223372036854775807
                                                _ = v124
                                            case "foo bar":
                                                v125 := GenS6{F0: 0}
                                                _ = v125
                                            }
                                        }
                                    }
                                    v4 = 127
                                }
                                var v126 uint32 = 2
                                _ = v126
                            } else {
                                v74 -= (v4 ^ (v4 - v74))
                            }
                            _ = fi119
                        }
                        fmt.Println(v74, v4)
                    }
                }
                for fi127 := 0; fi127 < 4; fi127++ {
                    var v128 float32 = -2.25
                    v4 = (v4 & (^(v74)))
                    switch "" {
                    case "hello":
                        fmt.Println(v128, v3, v4)
                    case "":
                        fmt.Println(fi75, v74)
                    }
                    var v129 uint32 = 831
                    _ = fi127
                    _ = v128
                    _ = v129
                }
            }
            fmt.Println(fi75, v4)
        }
        _ = fi75
    }
    if ((true && (128 != 587)) || (!(uint16(v1) > 983))) {
        v130 := GenS11{F0: v2, F1: 127}
        var v131 int32 = 571
        v132 := []int{32767, -11}
        _ = v130
        _ = v131
        _ = v132
    }
    v133 := GenS7{F0: 0, F1: 4294967295, F2: 32767}
    v134 := GenS7{F0: -2147483648, F1: 642, F2: 42}
    v135 := GenS12{F0: 32767, F1: 32767}
    fmt.Println(v74, v74)
    for fi136 := 0; fi136 < 5; fi136++ {
        v1 -= 0
        fmt.Println(v1, v2, v3)
        _ = fi136
    }
    var v137 int16 = -32768
    var v138 uint32 = 42
    for fi139 := 0; fi139 < 14; fi139++ {
        v3 = (v3 * 1e10)
        var v140 uint64 = (v2 | v2)
        v2 += 651
        _ = fi139
        _ = v140
    }
    v138 = (v138 * (v138 + (v138 + (v138 << ((2) & 63)))))
    var v141 uint8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v137:", v137)
    fmt.Println("v138:", v138)
    fmt.Println("v141:", v141)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v74:", v74)
    _ = v1
    _ = v133
    _ = v134
    _ = v135
    _ = v137
    _ = v138
    _ = v141
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v63
    _ = v74
}
