package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int16; F1 string }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 bool; F1 bool }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int8; F1 int64 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS15 struct { F0 bool; F1 uint16 }
type GenS16 struct { F0 uint; F1 uint8; F2 uint }
type GenS17 struct { F0 uint; F1 uint16; F2 int8 }
func genFn0(p0 string) int64 {
    var v7 int = 789
    for fi8 := 0; fi8 < 20; fi8++ {
        v7 = -976
        _ = fi8
    }
    var v9 int16 = 813
    if false {
        var v10 float32 = -2.25
        v10 = (1e10 / v10)
        p0 = p0
        _ = v10
    }
    var v11 uint32 = 1
    _ = p0
    _ = v11
    _ = v7
    _ = v9
    return -1
}
func genFn1(p0 uint16, p1 uint32) uint32 {
    fmt.Println(p0, p0)
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 int = 128
    v2 := make(map[string]string)
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    v5 := make(map[int]string)
    var v6 int64 = 1
    for fi12 := 0; fi12 < 3; fi12++ {
        v6 = -1
        var v13 string = ("hello" + "hello")
        _ = fi12
        _ = v13
    }
    if false {
        v1 += 2147483647
        var v14 int8 = -128
        _ = v14
    } else {
        var v15 uint8 = 42
        v1 -= 0
        v6 = v6
        _ = v15
    }
    if true {
        v16 := []int{v1, v1, (^(v1)), (v1 << ((0) & 63))}
        _ = v16
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 float64 = 1.5
        if false {
            v19 := GenS1{F0: 42, F1: v1}
            var v20 uint8 = 201
            if false {
                switch (v1 & 4294967295) {
                case 749:
                    v21 := GenS2{F0: 42, F1: v6, F2: (v20 - v20)}
                    _ = v21
                }
            }
            _ = v19
            _ = v20
        } else {
            if (255 == 32767) {
                for fi22 := 0; fi22 < 10; fi22++ {
                    var v23 int32 = 259
                    v24 := GenS2{F0: 366, F1: (v6 % (((v6 * v6)) | 1)), F2: 127}
                    var v25 int64 = 0
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
            } else {
                for fi26 := 0; fi26 < 5; fi26++ {
                    v6 = (v6 ^ (v6 << ((4294967295) & 63)))
                    _ = fi26
                }
                for fi27 := 0; fi27 < 5; fi27++ {
                    v2[""] = ("hello" + "\n\t")
                    v6 = 9223372036854775807
                    if false {
                        if (((!(false || (false && false))) || (v6 == int64(v1))) || false) {
                            fmt.Println(fi17, v18)
                            v28 := v4[2]
                            _ = v28
                        } else {
                            if (255 > 489) {
                                fmt.Println(v1, v18, v1)
                                fmt.Println(fi27, v1, fi27)
                                v6 -= -374
                            }
                        }
                        fmt.Println(v18, v18)
                    } else {
                        v2["hello"] = ("hello" + "hello")
                    }
                    _ = fi27
                }
                v1 = 32767
            }
            _ = len(v3)
            v29 := GenS2{F0: 2147483648, F1: (v6 / ((v6) | 1)), F2: 211}
            _ = v29
        }
        _ = fi17
        _ = v18
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        if (1 <= 167) {
            switch -325 {
            case -9223372036854775808:
                v6 -= (v6 << ((1) & 63))
            case 9223372036854775807:
                var v31 uint16 = 646
                _ = v31
            case 1:
                var v32 uint8 = 2
                _ = v32
            }
            if false {
                var v33 uint64 = 18446744073709551615
                _ = v33
            } else {
                var v34 string = ("" + "foo bar")
                v35 := []int{}
                v36 := GenS1{F0: 18446744073709551615, F1: (v1 - (fi30 - v1))}
                _ = v34
                _ = v35
                _ = v36
            }
        } else {
            switch "hello" {
            case "":
                v6 += (v6 * 42)
            case "\n\t":
                v1 = v1
            default:
                if false {
                    if (65535 != genFn1(0, 858)) {
                        if true {
                            var v37 int8 = 42
                            switch "hello" {
                            case "hello":
                                switch v1 {
                                case 2147483647:
                                    for fi38 := 0; fi38 < 5; fi38++ {
                                        v1 += (fi30 << ((0) & 63))
                                        var v39 int16 = int16(fi30)
                                        _ = fi38
                                        _ = v39
                                    }
                                case 9223372036854775807:
                                    v6 = -9223372036854775808
                                default:
                                    fmt.Println(v37, v1)
                                }
                            default:
                                for fi40 := 0; fi40 < 14; fi40++ {
                                    if (2 != (v37 & (v37 * -127))) {
                                        switch "foo bar" {
                                        case "":
                                            _ = len(v5)
                                        }
                                    }
                                    var v41 int16 = 0
                                    var v42 uint64 = 18446744073709551615
                                    _ = fi40
                                    _ = v41
                                    _ = v42
                                }
                            }
                            if ((v37 / (((v37 - (v37 / (((^(v37))) | 1)))) | 1)) >= v37) {
                                for fi43 := 0; fi43 < 6; fi43++ {
                                    if (-1 <= 1000) {
                                        v44 := GenS2{F0: 517, F1: v6, F2: 2}
                                        v45 := []int{fi30}
                                        v1 += (fi43 << ((221) & 63))
                                        _ = v44
                                        _ = v45
                                    } else {
                                        if false {
                                            if ((0.0 / -2.25) < 1e10) {
                                                v46 := GenS2{F0: 2147483647, F1: (v6 % ((v6) | 1)), F2: 181}
                                                var v47 uint32 = 35
                                                fmt.Println(v47, v37, v1)
                                                _ = v46
                                                _ = v47
                                            }
                                            if (!((false && true) && false)) {
                                                fmt.Println(v37, v1)
                                                var v48 int8 = (^(v37))
                                                _ = v48
                                            } else {
                                                v37 -= (v37 / (((v37 * (v37 + v37))) | 1))
                                                if (0 > -1) {
                                                    v6 = (-(v6))
                                                    v4[25] = (fi30 ^ (fi43 & (v1 + 127)))
                                                    var v49 int64 = -839
                                                    _ = v49
                                                } else {
                                                    v50 := []int{}
                                                    v51 := GenS3{F0: 0}
                                                    _ = v50
                                                    _ = v51
                                                }
                                            }
                                        } else {
                                            v52 := v2[""]
                                            if (65535 == 1) {
                                                if (369 < -37) {
                                                    v6 -= (v6 + v6)
                                                    v1 = fi43
                                                }
                                            } else {
                                                v53 := []int{(fi30 ^ (fi43 / (((fi43 ^ fi30)) | 1)))}
                                                switch (v1 - fi30) {
                                                case -169:
                                                    if false {
                                                        var v54 int64 = v6
                                                        _ = v54
                                                    }
                                                case 0:
                                                    v6 = v6
                                                default:
                                                    var v55 float32 = 0.1
                                                    _ = v55
                                                }
                                                var v56 int8 = (v37 ^ 42)
                                                _ = v53
                                                _ = v56
                                            }
                                            switch v52 {
                                            case "foo bar":
                                                fmt.Println(fi30, v37)
                                            case "":
                                                var v57 int32 = -568
                                                _ = v57
                                            default:
                                                switch 0 {
                                                case 2147483647:
                                                    v6 = 42
                                                default:
                                                    var v58 uint16 = 32767
                                                    _ = v58
                                                }
                                            }
                                            var v59 uint = 2
                                            v60 := []int{(-(v1))}
                                            if (-1 != 127) {
                                                fmt.Println(v52, v52, v52)
                                                var v61 string = v52
                                                _ = v61
                                            } else {
                                                var v62 uint = (v59 >> ((4294967295) & 63))
                                                if (!(!false)) {
                                                    v63 := GenS1{F0: 227, F1: v1}
                                                    if (42 > 796) {
                                                        var v64 int = (v1 - v1)
                                                        v1 -= v1
                                                        var v65 uint64 = 1
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                    _ = v63
                                                }
                                                _ = v62
                                            }
                                            fmt.Println(fi43, v6)
                                            var v66 string = v52
                                            v60 = append(v60, fi30)
                                            _ = v52
                                            _ = v59
                                            _ = v60
                                            _ = v66
                                        }
                                    }
                                    if (!(!(!false))) {
                                        v37 += (^(v37))
                                    }
                                    if true {
                                        var v67 int8 = (v37 + 110)
                                        _ = len(v2)
                                        _ = v67
                                    } else {
                                        v68 := []int{972, (fi30 - 920)}
                                        _ = v68
                                    }
                                    _ = fi43
                                }
                                v69 := GenS3{F0: 255}
                                _ = v69
                            }
                            _ = v37
                        } else {
                            v6 += v6
                            for fi70 := 0; fi70 < 7; fi70++ {
                                if false {
                                    var v71 int8 = -101
                                    fmt.Println(v6, fi30, v6)
                                    if (false && true) {
                                        v6 += (v6 << ((1) & 63))
                                    } else {
                                        var v72 int = (v1 & v1)
                                        var v73 uint8 = 1
                                        switch (v1 >> ((346) & 63)) {
                                        case 0:
                                            v74 := GenS4{F0: false, F1: (v73 << ((255) & 63))}
                                            var v75 int = 1
                                            _ = v74
                                            _ = v75
                                        default:
                                            if (2147483647 == 2147483647) {
                                                var v76 string = "a"
                                                var v77 int16 = int16(v71)
                                                v78 := GenS2{F0: 61, F1: 9223372036854775807, F2: v73}
                                                v79 := GenS3{F0: 255}
                                                var v80 int16 = v77
                                                if false {
                                                    var v81 uint16 = 75
                                                    if (1 >= 18446744073709551615) {
                                                        switch v1 {
                                                        case 42:
                                                            v82 := []int{-528, v1, (^(fi30)), (v1 * (v72 << ((2) & 63)))}
                                                            _ = v82
                                                        }
                                                        v83 := GenS1{F0: 1, F1: int(v81)}
                                                        switch v76 {
                                                        case "a":
                                                            v72 = v72
                                                        case "hello":
                                                            if (!(true && (!true))) {
                                                                switch (fi30 & (-(fi70))) {
                                                                case -950:
                                                                    switch v1 {
                                                                    case -1:
                                                                        v84 := []int{}
                                                                        _ = v84
                                                                    case 699:
                                                                        v85 := GenS5{F0: (fi70 ^ fi30)}
                                                                        _ = v85
                                                                    case 2147483647:
                                                                        v86 := GenS5{F0: (fi70 + (v1 - v1))}
                                                                        _ = v86
                                                                    }
                                                                }
                                                            }
                                                        case "\n\t":
                                                            var v87 float64 = -2.25
                                                            _ = v87
                                                        default:
                                                            var v88 uint16 = (v81 / ((v81) | 1))
                                                            _ = v88
                                                        }
                                                        _ = v83
                                                    } else {
                                                        switch "\n\t" {
                                                        case "a":
                                                            v89 := GenS4{F0: (v76 < v76), F1: v73}
                                                            _ = v89
                                                        case "":
                                                            v90 := []int{fi30, (v72 << ((18446744073709551615) & 63)), fi30, int(v77)}
                                                            _ = v90
                                                        default:
                                                            v91 := GenS3{F0: v81}
                                                            _ = v91
                                                        }
                                                        v92 := []int{(v72 | fi30), (v1 + fi30), v72, -9223372036854775808}
                                                        v77 -= 77
                                                        _ = v92
                                                    }
                                                    _ = v81
                                                }
                                                _ = v76
                                                _ = v77
                                                _ = v78
                                                _ = v79
                                                _ = v80
                                            }
                                        }
                                        v93 := GenS1{F0: 0, F1: 2147483647}
                                        v6 -= (^(v6))
                                        _ = v72
                                        _ = v73
                                        _ = v93
                                    }
                                    if ((1 > 0) && (!false)) {
                                        var v94 string = "\n\t"
                                        fmt.Println(fi70)
                                        switch 2147483647 {
                                        case 9223372036854775807:
                                            if true {
                                                v95 := GenS6{F0: -341}
                                                v94 = "\n\t"
                                                switch 2147483647 {
                                                case 1:
                                                    var v96 string = (v94 + v94)
                                                    _ = v96
                                                case 32768:
                                                    v97 := GenS7{F0: 255}
                                                    _ = v97
                                                case 42:
                                                    switch v94 {
                                                    case "foo bar":
                                                        switch "\n\t" {
                                                        case "":
                                                            fmt.Println(fi70, v1)
                                                        }
                                                    case "\n\t":
                                                        var v98 uint = 65535
                                                        _ = v98
                                                    default:
                                                        delete(v3, "foo bar")
                                                    }
                                                }
                                                if ((float32((float64(0.1) / 1.0)) < (0.1 * float32((3.14159 * -2.25)))) && true) {
                                                    if true {
                                                        v99 := GenS1{F0: 334, F1: fi70}
                                                        _ = v99
                                                    }
                                                    switch v94 {
                                                    case "":
                                                        v94 = v94
                                                    default:
                                                        var v100 uint64 = 18446744073709551615
                                                        _ = v100
                                                    }
                                                    var v101 float32 = 0.1
                                                    _ = v101
                                                } else {
                                                    switch v94 {
                                                    case "foo bar":
                                                        v102 := GenS1{F0: 389, F1: 518}
                                                        _ = v102
                                                    case "hello":
                                                        var v103 float32 = (0.0 / 1.0)
                                                        _ = v103
                                                    case "":
                                                        v104 := GenS8{F0: 1, F1: v94}
                                                        _ = v104
                                                    default:
                                                        v105 := []int{}
                                                        _ = v105
                                                    }
                                                    v71 = -120
                                                }
                                                _ = v95
                                            } else {
                                                switch v94 {
                                                case "\n\t":
                                                    switch (fi70 / ((fi70) | 1)) {
                                                    case 127:
                                                        fmt.Println(fi30, fi70)
                                                    case 32768:
                                                        if false {
                                                            v106 := []int{fi70, (fi30 / (((v1 % (((fi30 << ((732) & 63))) | 1))) | 1))}
                                                            _ = v106
                                                        } else {
                                                            fmt.Println(v1, fi70)
                                                            v6 -= genFn0(v94)
                                                        }
                                                    default:
                                                        v1 = 127
                                                    }
                                                default:
                                                    var v107 int = (v1 | fi30)
                                                    _ = v107
                                                }
                                            }
                                        default:
                                            v1 += (fi70 & -1)
                                        }
                                        _ = v94
                                    }
                                    _ = v71
                                } else {
                                    if false {
                                        var v108 float64 = (0.0 / 1.0)
                                        delete(v4, 42)
                                        _ = v108
                                    } else {
                                        var v109 string = "\n\t"
                                        if (false && false) {
                                            v110 := GenS1{F0: 348, F1: 2147483647}
                                            var v111 int64 = (v6 >> ((42) & 63))
                                            fmt.Println(v6, v6)
                                            v109 = v109
                                            _ = v110
                                            _ = v111
                                        } else {
                                            v6 = v6
                                            v109 = "foo bar"
                                        }
                                        _ = v109
                                    }
                                    if (251 < 2) {
                                        v6 = v6
                                        switch 127 {
                                        case 127:
                                            v112 := []int{v1, (v1 * fi70), (fi30 | (-(fi70))), (fi30 - fi70)}
                                            _ = v112
                                        case -245:
                                            v1 = 4294967295
                                            fmt.Println(fi30)
                                        }
                                    } else {
                                        var v113 int16 = 1000
                                        v6 = (v6 % ((v6) | 1))
                                        switch (fi70 - fi30) {
                                        case -2147483648:
                                            v6 -= 42
                                        case 128:
                                            if false {
                                                switch (-(fi70)) {
                                                case 127:
                                                    v1 = (fi70 * 204)
                                                case 32767:
                                                    v1 -= (v1 | fi70)
                                                    v114 := []int{}
                                                    _ = v114
                                                default:
                                                    if (7 != 128) {
                                                        switch "" {
                                                        case "hello":
                                                            if ((!false) || ((v6 * -1) == (-(v6)))) {
                                                                var v115 uint = 866
                                                                if (((!(!true)) && false) && ((^(v6)) > (v6 + v6))) {
                                                                    if true {
                                                                        v1 = (fi30 >> ((v115) & 63))
                                                                    } else {
                                                                        v116 := []int{v1, v1}
                                                                        _ = v116
                                                                    }
                                                                    if true {
                                                                        var v117 uint32 = 4294967295
                                                                        v118 := GenS8{F0: 1000, F1: ("" + "foo bar")}
                                                                        fmt.Println(v113)
                                                                        _ = v117
                                                                        _ = v118
                                                                    } else {
                                                                        var v119 int = fi70
                                                                        _ = v119
                                                                    }
                                                                } else {
                                                                    var v120 int64 = (v6 ^ (^(v6)))
                                                                    v1 = (v1 << (((v115 >> (((v115 % ((v115) | 1))) & 63))) & 63))
                                                                    v1 = (v1 >> (((^(v115))) & 63))
                                                                    v121 := GenS4{F0: true, F1: 50}
                                                                    _ = v120
                                                                    _ = v121
                                                                }
                                                                _ = v115
                                                            }
                                                        case "foo bar":
                                                            var v122 uint64 = 42
                                                            _ = v122
                                                        }
                                                        if true {
                                                            v113 += (v113 & v113)
                                                            v123 := GenS8{F0: v113, F1: (("" + "\n\t") + ("" + "\n\t"))}
                                                            if true {
                                                                _ = len(v3)
                                                                fmt.Println(v113, v113, v1)
                                                                if (!(!((((("\n\t" + "hello") + "foo bar") + "a") + "") > ""))) {
                                                                    v1 = 1
                                                                    var v124 int16 = 42
                                                                    v125 := GenS3{F0: 255}
                                                                    _ = v124
                                                                    _ = v125
                                                                } else {
                                                                    v126 := GenS6{F0: (v113 | v113)}
                                                                    _ = v126
                                                                }
                                                            } else {
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    v113 = v113
                                                                case "":
                                                                    fmt.Println(v113)
                                                                default:
                                                                    switch v1 {
                                                                    case 128:
                                                                        v6 += genFn0((("\n\t" + "hello") + "hello"))
                                                                    case 9223372036854775807:
                                                                        v5[(^(fi70))] = "foo bar"
                                                                        if false {
                                                                            v113 += v113
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            var v127 uint16 = 223
                                                            v128 := []int{(fi70 >> ((2) & 63)), fi70, 9223372036854775807, (fi30 % (((v1 % ((-2147483648) | 1))) | 1))}
                                                            _ = v123
                                                            _ = v127
                                                            _ = v128
                                                        } else {
                                                            var v129 int64 = v6
                                                            v1 += fi70
                                                            _ = v129
                                                        }
                                                        switch 32767 {
                                                        case 4294967295:
                                                            v130 := []int{(fi30 | -2147483648), 0, (^(fi70))}
                                                            _ = v130
                                                        case 32767:
                                                            v131 := GenS9{F0: v113}
                                                            _ = v131
                                                        }
                                                        v1 = -347
                                                    } else {
                                                        v132 := []int{v1, (fi30 << ((4294967295) & 63)), (v1 - (^(fi30)))}
                                                        _ = v132
                                                    }
                                                }
                                                if false {
                                                    fmt.Println(fi30, fi30, fi30)
                                                } else {
                                                    var v133 uint = 626
                                                    _ = v133
                                                }
                                                var v134 uint16 = 1
                                                _ = v134
                                            } else {
                                                v6 = 0
                                                v6 = (-(v6))
                                                switch "\n\t" {
                                                case "hello":
                                                    v135 := GenS3{F0: 1}
                                                    _ = v135
                                                case "":
                                                    var v136 uint8 = 128
                                                    _ = v136
                                                case "a":
                                                    var v137 uint64 = 9
                                                    _ = v137
                                                default:
                                                    switch "a" {
                                                    case "":
                                                        delete(v3, "foo bar")
                                                    case "a":
                                                        if ((-2.25 <= (3.14159 - 1e10)) || true) {
                                                            if true {
                                                                v138 := []int{(-(fi30)), 32767}
                                                                v139 := GenS4{F0: false, F1: 128}
                                                                var v140 uint64 = 217
                                                                _ = v138
                                                                _ = v139
                                                                _ = v140
                                                            }
                                                            if (-128 != 0) {
                                                                v113 = 42
                                                            } else {
                                                                v141 := []int{(^(v1))}
                                                                fmt.Println(fi70, fi70)
                                                                var v142 int = (v1 + fi70)
                                                                v5[-1] = "a"
                                                                var v143 float64 = float64((1e10 * 0.1))
                                                                v1 = -495
                                                                if false {
                                                                    var v144 string = "\n\t"
                                                                    v1 = (fi30 % ((-613) | 1))
                                                                    _ = v144
                                                                } else {
                                                                    var v145 uint16 = 1
                                                                    v3["hello"] = v1
                                                                    _ = v145
                                                                }
                                                                if false {
                                                                    fmt.Println(v113)
                                                                    fmt.Println(v143)
                                                                    v143 += v143
                                                                } else {
                                                                    v146 := GenS7{F0: 42}
                                                                    var v147 int64 = -1
                                                                    v147 -= v147
                                                                    if (323 >= 255) {
                                                                        var v148 int8 = 0
                                                                        _ = v148
                                                                    } else {
                                                                        v149 := GenS3{F0: 0}
                                                                        var v150 uint16 = 65535
                                                                        v151 := GenS10{F0: (!(1.5 < 1.5)), F1: true}
                                                                        fmt.Println(fi70)
                                                                        _ = v149
                                                                        _ = v150
                                                                        _ = v151
                                                                    }
                                                                    _ = v146
                                                                    _ = v147
                                                                }
                                                                var v152 uint64 = 4294967295
                                                                _ = v141
                                                                _ = v142
                                                                _ = v143
                                                                _ = v152
                                                            }
                                                        } else {
                                                            switch "a" {
                                                            case "\n\t":
                                                                var v153 uint32 = 2147483647
                                                                _ = v153
                                                            }
                                                        }
                                                    case "\n\t":
                                                        var v154 float32 = (0.1 / float32((float64(0.0) + 3.14159)))
                                                        _ = v154
                                                    }
                                                }
                                                if true {
                                                    var v155 int8 = 99
                                                    v156 := GenS11{F0: 65535}
                                                    v157 := GenS12{F0: -1, F1: int64(v1)}
                                                    _ = v155
                                                    _ = v156
                                                    _ = v157
                                                } else {
                                                    if true {
                                                        v158 := GenS13{F0: 4294967295}
                                                        _ = v158
                                                    }
                                                    v6 = (v6 * -9223372036854775808)
                                                    fmt.Println(v6, fi30)
                                                }
                                                v6 -= v6
                                                v113 += (v113 * v113)
                                            }
                                        default:
                                            v6 += -1
                                        }
                                        var v159 string = ((("" + ("foo bar" + (("a" + (((("hello" + (("foo bar" + "hello") + ("a" + (("foo bar" + ("a" + ("a" + ("" + "")))) + ("hello" + "foo bar"))))) + ("hello" + (("" + "\n\t") + (("\n\t" + ("" + (("\n\t" + (("\n\t" + "hello") + "\n\t")) + "\n\t"))) + "")))) + "foo bar") + "a")) + "\n\t"))) + ("hello" + ("" + "foo bar"))) + "\n\t")
                                        _ = v113
                                        _ = v159
                                    }
                                    if true {
                                        var v160 bool = false
                                        var v161 uint16 = 0
                                        fmt.Println(v161, v6, fi70)
                                        _ = v160
                                        _ = v161
                                    }
                                }
                                _ = fi70
                            }
                        }
                    }
                } else {
                    var v162 bool = ((1e10 / (((0.1 + float32((1e10 / 3.14159))) * 0.0) + ((1e10 - -2.25) / 1.5))) == float32(0.1))
                    _ = v162
                }
            }
        }
        v1 -= v1
        _ = fi30
    }
    v1 -= v1
    v6 += v6
    v163 := GenS3{F0: 635}
    switch ("foo bar" + "") {
    case "a":
        var v164 int32 = 1
        _ = v164
    case "\n\t":
        switch "foo bar" {
        case "\n\t":
            var v165 uint16 = 0
            _ = v165
        case "":
            for fi166 := 0; fi166 < 18; fi166++ {
                for fi167 := 0; fi167 < 3; fi167++ {
                    if (282 >= 65535) {
                        v168 := GenS14{F0: 502, F1: 4294967295, F2: 11}
                        var v169 int64 = (v6 % ((v6) | 1))
                        _ = v168
                        _ = v169
                    } else {
                        var v170 int64 = v6
                        var v171 int8 = 1
                        _ = v170
                        _ = v171
                    }
                    delete(v4, 65535)
                    v172 := GenS2{F0: 65535, F1: v6, F2: 255}
                    _ = fi167
                    _ = v172
                }
                v173 := GenS15{F0: false, F1: 0}
                _ = fi166
                _ = v173
            }
        default:
            for fi174 := 0; fi174 < 5; fi174++ {
                for fi175 := 0; fi175 < 3; fi175++ {
                    v6 -= v6
                    v6 = 0
                    if (574 > 986) {
                        v1 = (^(v1))
                        var v176 int64 = (v6 - (v6 * v6))
                        _ = v176
                    }
                    v177 := GenS16{F0: 0, F1: 255, F2: 65535}
                    v1 = (v1 + -839)
                    _ = fi175
                    _ = v177
                }
                _ = fi174
            }
        }
    default:
        for fi178 := 0; fi178 < 20; fi178++ {
            v179 := GenS17{F0: 722, F1: 164, F2: -1}
            if true {
                var v180 float64 = ((3.14159 + 1e10) / 1.0)
                v1 -= (v1 + v1)
                _ = v180
            }
            var v181 int64 = 182
            _ = fi178
            _ = v179
            _ = v181
        }
    }
    for fi182 := 0; fi182 < 7; fi182++ {
        var v183 string = "a"
        v184 := v4[genFn1(638, 2)]
        _ = fi182
        _ = v183
        _ = v184
    }
    v185 := GenS6{F0: -1}
    var v186 int64 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v186:", v186)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v163
    _ = v185
    _ = v186
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
