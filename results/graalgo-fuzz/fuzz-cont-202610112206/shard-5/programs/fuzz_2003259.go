package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 bool }
type GenS2 struct { F0 int32; F1 bool; F2 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS6 struct { F0 string; F1 bool; F2 uint8 }
type GenS7 struct { F0 uint32; F1 int16; F2 int8 }
type GenS8 struct { F0 int8; F1 uint64; F2 int }
type GenS9 struct { F0 uint64; F1 int }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 int16; F1 uint }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 int32; F1 uint64; F2 int }
func genFn0(p0 int64, p1 bool) int8 {
    switch 9223372036854775807 {
    case 4294967295:
        var v7 string = "foo bar"
        _ = v7
    case -9223372036854775808:
        var v8 float32 = 0.1
        _ = v8
    default:
        v9 := []int{42, 127}
        _ = v9
    }
    if p1 {
        switch 31 {
        case 32767:
            v10 := []int{}
            _ = v10
        case 930:
            if p1 {
                var v11 uint8 = 148
                var v12 int32 = -2147483648
                _ = v11
                _ = v12
            } else {
                switch (("a" + (("\n\t" + "a") + "foo bar")) + "a") {
                case "hello":
                    switch 32768 {
                    case 224:
                        fmt.Println(p0, p0)
                    }
                }
                var v13 string = "hello"
                _ = v13
            }
        }
        v14 := GenS1{F0: 2, F1: -790, F2: p1}
        _ = v14
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        switch 0 {
        case 2147483647:
            if ((p1 && (!(-2147483648 > (fi15 + -331)))) || p1) {
                switch "hello" {
                case "":
                    v16 := GenS1{F0: 42, F1: 1, F2: false}
                    _ = v16
                case "a":
                    var v17 int16 = 1000
                    _ = v17
                case "foo bar":
                    v18 := []int{(fi15 >> ((633) & 63)), (fi15 & fi15), -691}
                    _ = v18
                default:
                    v19 := []int{(fi15 * fi15)}
                    _ = v19
                }
                var v20 float64 = 0.1
                p1 = (("" + "foo bar") != "")
                _ = v20
            }
        }
        v21 := []int{(fi15 * (fi15 + 128)), fi15, fi15, fi15}
        _ = fi15
        _ = v21
    }
    var v22 uint32 = 1
    v23 := []int{-1, 32768}
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return 62
}
func genFn1(p0 int64, p1 bool) int16 {
    p1 = true
    var v24 uint64 = 477
    if (0 != 42) {
        for fi25 := 0; fi25 < 2; fi25++ {
            fmt.Println(fi25)
            _ = fi25
        }
        if (false && true) {
            var v26 uint16 = 152
            p1 = (!false)
            _ = v26
        } else {
            switch 128 {
            case 128:
                p0 -= p0
            }
        }
    }
    _ = p0
    _ = p1
    _ = v24
    return 1000
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = ""
    v3 := make(map[uint32]uint32)
    var v4 int64 = 9223372036854775807
    var v5 float64 = 0.1
    var v6 uint8 = 190
    v27 := GenS1{F0: 0, F1: -785, F2: (v4 > -9223372036854775808)}
    if (!false) {
        if true {
            v28 := GenS1{F0: 32767, F1: 359, F2: true}
            _ = v28
        }
        for fi29 := 0; fi29 < 4; fi29++ {
            fmt.Println(fi29, v6)
            v5 += float64(1e10)
            _ = fi29
        }
        _ = len(v1)
    }
    for fi30 := 0; fi30 < 3; fi30++ {
        for fi31 := 0; fi31 < 16; fi31++ {
            v32 := GenS1{F0: 2, F1: -251, F2: (59 < 100)}
            v2 = "foo bar"
            _ = fi31
            _ = v32
        }
        v33 := GenS2{F0: 0, F1: ((!false) || (1 >= -128)), F2: genFn0((^(v4)), false)}
        for fi34 := 0; fi34 < 19; fi34++ {
            var v35 uint = 18446744073709551615
            v35 = (v35 - 61)
            var v36 bool = false
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi30
        _ = v33
    }
    if false {
        if (!false) {
            fmt.Println(v4, v4, v6)
            if (440 <= 2147483647) {
                v37 := GenS3{F0: 255}
                _ = v37
            } else {
                var v38 int = -9223372036854775808
                fmt.Println(v4)
                v39 := GenS1{F0: 851, F1: 103, F2: (-2.25 > 1.5)}
                _ = v38
                _ = v39
            }
            if ((!true) || false) {
                switch v2 {
                case "\n\t":
                    var v40 int16 = 1000
                    _ = v40
                case "hello":
                    switch ("\n\t" + "foo bar") {
                    case "hello":
                        v41 := GenS4{F0: 480, F1: true}
                        _ = v41
                    case "foo bar":
                        v4 = v4
                    }
                case "foo bar":
                    v6 -= (v6 * 12)
                }
            }
        }
        var v42 int = 2147483647
        _ = v42
    } else {
        fmt.Println(v6, v4)
    }
    switch -511 {
    case 0:
        if true {
            fmt.Println(v5, v4)
            var v43 int32 = -1
            _ = v43
        } else {
            switch 9223372036854775807 {
            case 127:
                var v44 bool = (true || (255 < 42))
                _ = v44
            default:
                v45 := GenS5{F0: false, F1: 1, F2: 127}
                _ = v45
            }
            if (!true) {
                var v46 uint = 1
                var v47 float64 = -2.25
                var v48 float32 = -2.25
                _ = v46
                _ = v47
                _ = v48
            }
        }
    default:
        var v49 int64 = (v4 / ((42) | 1))
        _ = v49
    }
    v50 := GenS2{F0: -2147483648, F1: false, F2: 1}
    var v51 uint16 = 662
    v2 = v2
    var v52 int64 = 1
    fmt.Println(v5, v6, v5)
    if (!false) {
        if false {
            var v53 bool = (!true)
            _ = v53
        } else {
            for fi54 := 0; fi54 < 3; fi54++ {
                if (v4 == (v52 ^ v4)) {
                    fmt.Println(v6, v52)
                    var v55 uint = 460
                    var v56 uint64 = 1
                    _ = v55
                    _ = v56
                }
                if true {
                    switch (v2 + "a") {
                    case "hello":
                        for fi57 := 0; fi57 < 5; fi57++ {
                            if true {
                                if true {
                                    v3[uint32(v6)] = 42
                                    switch ("hello" + "a") {
                                    case "\n\t":
                                        v58 := GenS3{F0: uint(v6)}
                                        _ = v58
                                    case "":
                                        if (-219 < -597) {
                                            v59 := GenS1{F0: (v51 ^ (v51 << ((2) & 63))), F1: 2147483647, F2: false}
                                            switch ((("" + v2) + v2) + "foo bar") {
                                            case "foo bar":
                                                if false {
                                                    var v60 int16 = 647
                                                    if true {
                                                        v6 -= v6
                                                        v61 := []int{}
                                                        var v62 uint = 4294967295
                                                        _ = v61
                                                        _ = v62
                                                    } else {
                                                        v5 += v5
                                                        v6 = (^(v6))
                                                        v60 += (^(v60))
                                                    }
                                                    v51 = (v51 ^ v51)
                                                    _ = v60
                                                } else {
                                                    if false {
                                                        v51 -= v51
                                                        var v63 int = fi54
                                                        _ = v63
                                                    }
                                                    var v64 bool = true
                                                    v4 = (-(v4))
                                                    _ = v64
                                                }
                                            case "\n\t":
                                                v65 := GenS4{F0: 228, F1: false}
                                                _ = v65
                                            case "":
                                                _ = len(v1)
                                            default:
                                                if (!(-65 >= v4)) {
                                                    var v66 uint64 = 904
                                                    v67 := []int{(^(fi54))}
                                                    _ = v66
                                                    _ = v67
                                                } else {
                                                    v68 := GenS3{F0: 4294967295}
                                                    var v69 bool = false
                                                    switch (^(fi54)) {
                                                    case 1:
                                                        fmt.Println(v52)
                                                    case 4294967295:
                                                        if (2147483648 > 841) {
                                                            switch v2 {
                                                            case "\n\t":
                                                                if v69 {
                                                                    if (1 > 944) {
                                                                        v70 := GenS5{F0: (!((v52 >= (v4 % ((42) | 1))) || v69)), F1: 18446744073709551615, F2: v6}
                                                                        fmt.Println(v6, fi54)
                                                                        v71 := GenS1{F0: 387, F1: 289, F2: (genFn0((v52 / ((-1) | 1)), (!(163 <= (fi57 + (fi57 - (fi54 % ((fi57) | 1))))))) >= 100)}
                                                                        _ = v70
                                                                        _ = v71
                                                                    } else {
                                                                        switch v2 {
                                                                        case "a":
                                                                            switch (fi54 >> ((255) & 63)) {
                                                                            case 77:
                                                                                if v69 {
                                                                                    v72 := []int{fi57}
                                                                                    var v73 int64 = (v52 / ((v52) | 1))
                                                                                    var v74 int32 = 0
                                                                                    v75 := GenS6{F0: "\n\t", F1: v69, F2: v6}
                                                                                    switch fi54 {
                                                                                    case 32767:
                                                                                        if v69 {
                                                                                            v76 := GenS7{F0: 999, F1: 0, F2: 52}
                                                                                            _ = v76
                                                                                        } else {
                                                                                            if v69 {
                                                                                                var v77 uint32 = 460
                                                                                                _ = v77
                                                                                            } else {
                                                                                                if v69 {
                                                                                                    _ = len(v72)
                                                                                                    v78 := GenS8{F0: -66, F1: 0, F2: fi54}
                                                                                                    var v79 int64 = v73
                                                                                                    v80 := GenS3{F0: 255}
                                                                                                    var v81 uint8 = v6
                                                                                                    v79 = (v4 & v4)
                                                                                                    v73 = (v52 ^ (v73 / ((v52) | 1)))
                                                                                                    _ = v78
                                                                                                    _ = v79
                                                                                                    _ = v80
                                                                                                    _ = v81
                                                                                                }
                                                                                                v6 += (^(v6))
                                                                                                _ = len(v1)
                                                                                            }
                                                                                            if false {
                                                                                                if false {
                                                                                                    v82 := GenS9{F0: 1, F1: int(v6)}
                                                                                                    var v83 int = (fi54 / ((fi57) | 1))
                                                                                                    _ = v82
                                                                                                    _ = v83
                                                                                                }
                                                                                                var v84 int = int(v74)
                                                                                                _ = v84
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v74 += (-(v74))
                                                                                    _ = v72
                                                                                    _ = v73
                                                                                    _ = v74
                                                                                    _ = v75
                                                                                }
                                                                            case -2147483648:
                                                                                v85 := GenS10{F0: v69}
                                                                                _ = v85
                                                                            case 42:
                                                                                var v86 uint16 = (v51 % ((v51) | 1))
                                                                                _ = v86
                                                                            }
                                                                        case "":
                                                                            fmt.Println(v51)
                                                                        default:
                                                                            v87 := GenS11{F0: v2}
                                                                            _ = v87
                                                                        }
                                                                        var v88 uint16 = (v51 ^ (v51 * 632))
                                                                        _ = v88
                                                                    }
                                                                    fmt.Println(v51, v6)
                                                                } else {
                                                                    switch fi54 {
                                                                    case 4294967295:
                                                                        v2 = v2
                                                                    default:
                                                                        var v89 uint8 = v6
                                                                        _ = v89
                                                                    }
                                                                    var v90 uint32 = 2
                                                                    if (((v2 + "") < v2) || false) {
                                                                        v69 = v69
                                                                        if (v4 != (-(v52))) {
                                                                            fmt.Println(v2, v6)
                                                                        }
                                                                        delete(v1, v90)
                                                                        v91 := GenS12{F0: 0}
                                                                        var v92 uint8 = v6
                                                                        _ = v91
                                                                        _ = v92
                                                                    } else {
                                                                        if v69 {
                                                                            v93 := GenS4{F0: 1, F1: (((v4 & (v52 | (v52 << ((4294967295) & 63)))) == 854) || v69)}
                                                                            v94 := []int{-1, -1}
                                                                            switch v2 {
                                                                            case "hello":
                                                                                var v95 uint32 = v90
                                                                                _ = v95
                                                                            case "":
                                                                                var v96 float64 = ((v5 * v5) * (v5 + 1.5))
                                                                                _ = v96
                                                                            case "\n\t":
                                                                                var v97 float32 = 0.0
                                                                                _ = v97
                                                                            }
                                                                            _ = v93
                                                                            _ = v94
                                                                        }
                                                                    }
                                                                    _ = v90
                                                                }
                                                            }
                                                            v98 := GenS13{F0: 42}
                                                            _ = v98
                                                        }
                                                    case 433:
                                                        fmt.Println(fi54)
                                                    }
                                                    switch ((v2 + (((((v2 + v2) + ("" + v2)) + v2) + "foo bar") + "hello")) + v2) {
                                                    case "":
                                                        var v99 int32 = 42
                                                        _ = v99
                                                    case "foo bar":
                                                        switch (v2 + v2) {
                                                        case "hello":
                                                            if true {
                                                                var v100 string = v2
                                                                _ = v100
                                                            } else {
                                                                fmt.Println(v5, v6, v5)
                                                            }
                                                        case "foo bar":
                                                            var v101 int32 = 0
                                                            _ = v101
                                                        case "a":
                                                            var v102 string = v2
                                                            _ = v102
                                                        }
                                                    case "\n\t":
                                                        fmt.Println(fi57)
                                                    }
                                                    _ = v68
                                                    _ = v69
                                                }
                                            }
                                            _ = v59
                                        } else {
                                            var v103 float32 = (1.5 + float32(float64((0.1 + 1.5))))
                                            v6 = v6
                                            var v104 int8 = -14
                                            _ = v103
                                            _ = v104
                                        }
                                    default:
                                        if ((^(v51)) > (v51 ^ (v51 << ((42) & 63)))) {
                                            if true {
                                                v3[65535] = 857
                                                switch "foo bar" {
                                                case "a":
                                                    v105 := []int{(fi54 - int(v6))}
                                                    v4 += 0
                                                    v6 -= v6
                                                    _ = v105
                                                default:
                                                    if true {
                                                        v52 = v4
                                                        var v106 bool = (42 != 571)
                                                        var v107 uint16 = 197
                                                        _ = v106
                                                        _ = v107
                                                    } else {
                                                        v108 := []int{-9223372036854775808, 708}
                                                        _ = v108
                                                    }
                                                }
                                            }
                                            if (!false) {
                                                v52 -= v52
                                            }
                                            v4 += (v4 | (v52 % ((167) | 1)))
                                        }
                                    }
                                } else {
                                    v109 := []int{}
                                    fmt.Println(v6, v51)
                                    _ = v109
                                }
                                if (866 != 65535) {
                                    if (-2.25 <= 0.1) {
                                        if false {
                                            v110 := GenS14{F0: (v6 / (((v6 ^ v6)) | 1))}
                                            if (!(!(false && false))) {
                                                fmt.Println(v52, v4)
                                                switch ((v2 + "\n\t") + (v2 + "")) {
                                                case "hello":
                                                    v111 := GenS15{F0: int16(fi54), F1: 707}
                                                    _ = v111
                                                case "foo bar":
                                                    v52 = (v4 | v4)
                                                }
                                                var v112 int = (fi57 << ((4294967295) & 63))
                                                _ = v112
                                            }
                                            if ((v52 % ((765) | 1)) <= 42) {
                                                var v113 uint64 = 42
                                                _ = v113
                                            }
                                            _ = v110
                                        }
                                        v51 += v51
                                    } else {
                                        v51 -= v51
                                        v51 = 65535
                                        v6 -= 42
                                        v114 := GenS14{F0: (^(v6))}
                                        _ = v114
                                    }
                                }
                            }
                            fmt.Println(v6, v6, fi54)
                            _ = fi57
                        }
                    case "a":
                        if (false || (42 > 623)) {
                            var v115 int64 = 42
                            _ = v115
                        }
                    }
                    for fi116 := 0; fi116 < 2; fi116++ {
                        var v117 uint8 = (v6 - v6)
                        v52 -= 42
                        v4 = v52
                        v118 := []int{}
                        var v119 int = (fi54 >> ((2) & 63))
                        switch "a" {
                        case "\n\t":
                            fmt.Println(v117, v51, v51)
                        case "foo bar":
                            v118 = append(v118, fi54)
                        default:
                            v120 := GenS4{F0: 9, F1: (true && (v5 < 0.1))}
                            _ = v120
                        }
                        v51 = 665
                        var v121 int16 = 1000
                        _ = fi116
                        _ = v117
                        _ = v118
                        _ = v119
                        _ = v121
                    }
                    v6 += (v6 >> ((42) & 63))
                }
                if false {
                    v122 := GenS16{F0: v4}
                    _ = len(v3)
                    v6 -= (v6 / ((128) | 1))
                    _ = v122
                } else {
                    var v123 int16 = -129
                    v124 := []int{(fi54 % ((0) | 1)), (fi54 % ((-2147483648) | 1)), fi54, 127}
                    switch fi54 {
                    case -567:
                        for fi125 := 0; fi125 < 4; fi125++ {
                            if false {
                                v126 := GenS17{F0: 2147483647, F1: 4294967295, F2: 4294967295}
                                switch -791 {
                                case 127:
                                    fmt.Println(v2, v2, v52)
                                default:
                                    v127 := v1[596]
                                    _ = v127
                                }
                                _ = v126
                            }
                            _ = len(v3)
                            switch (fi125 ^ (fi125 >> ((803) & 63))) {
                            case 2147483647:
                                fmt.Println(v2)
                            default:
                                var v128 int32 = -1
                                _ = v128
                            }
                            _ = fi125
                        }
                    default:
                        v4 = (v4 & (v4 << ((47) & 63)))
                    }
                    _ = v123
                    _ = v124
                }
                _ = fi54
            }
        }
        for fi129 := 0; fi129 < 11; fi129++ {
            v130 := GenS9{F0: 0, F1: 128}
            if (!true) {
                for fi131 := 0; fi131 < 12; fi131++ {
                    if false {
                        switch (("hello" + ("hello" + v2)) + v2) {
                        case "":
                            if (1.5 == v5) {
                                v51 -= v51
                                v4 = (v4 / ((v4) | 1))
                                v132 := []int{4294967295}
                                _ = v132
                            }
                            v133 := GenS6{F0: "foo bar", F1: false, F2: v6}
                            fmt.Println(fi129, v6)
                            _ = v133
                        case "hello":
                            switch "hello" {
                            case "a":
                                v6 = 128
                            default:
                                var v134 float64 = (v5 * v5)
                                _ = v134
                            }
                        case "foo bar":
                            delete(v1, 42)
                        }
                    }
                    v2 = "a"
                    var v135 string = ((v2 + v2) + (v2 + v2))
                    v6 = 255
                    _ = fi131
                    _ = v135
                }
                delete(v1, 65535)
            } else {
                for fi136 := 0; fi136 < 9; fi136++ {
                    v137 := GenS7{F0: 81, F1: 0, F2: genFn0((v52 >> ((4294967295) & 63)), true)}
                    v1[65535] = v2
                    _ = fi136
                    _ = v137
                }
                for fi138 := 0; fi138 < 16; fi138++ {
                    var v139 int = 1000
                    _ = fi138
                    _ = v139
                }
            }
            _ = fi129
            _ = v130
        }
    } else {
        switch 127 {
        case 927:
            v140 := GenS10{F0: false}
            _ = v140
        default:
            v4 = (v52 / (((v4 * (v52 - v52))) | 1))
        }
    }
    v4 += (v52 / (((^(v52))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
}
