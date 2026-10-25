package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 uint; F1 uint; F2 uint64 }
type GenS8 struct { F0 int64; F1 uint64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int16; F1 int16; F2 int64 }
type GenS11 struct { F0 uint16; F1 string }
type GenS12 struct { F0 int64; F1 string }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int8; F1 int; F2 int8 }
func genFn0(p0 bool, p1 int64, p2 bool) uint16 {
    v6 := []int{2147483647, 1, 0}
    if len(v6) > 0 {
        v6[0] = 2147483647
    }
    var v7 string = "hello"
    if true {
        var v8 int64 = p1
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint16 = 67
    var v2 float32 = (0.0 - 0.1)
    v3 := make(map[int]string)
    var v4 uint32 = 2
    var v5 int = 839
    if true {
        var v9 int = v5
        v10 := GenS1{F0: 88}
        _ = v10
        _ = v9
    }
    v11 := []int{(^(v5)), v5, 462}
    var v12 uint = 2
    for fi13 := 0; fi13 < 14; fi13++ {
        if false {
            fmt.Println(v4, v5)
            for fi14 := 0; fi14 < 15; fi14++ {
                v15 := GenS1{F0: v1}
                if (84 < -1) {
                    var v16 int64 = 0
                    if false {
                        v2 -= float32(3.14159)
                        switch fi13 {
                        case 234:
                            switch "a" {
                            case "hello":
                                _ = len(v3)
                            default:
                                v12 = 18446744073709551615
                            }
                        case 32767:
                            switch "" {
                            case "a":
                                switch ("hello" + "foo bar") {
                                case "a":
                                    fmt.Println(v4, fi14)
                                default:
                                    v5 = (^(fi13))
                                }
                            case "foo bar":
                                var v17 int = -62
                                _ = v17
                            case "hello":
                                v4 += (^(v4))
                            }
                        case 42:
                            v12 += (^(v12))
                        }
                    }
                    _ = len(v3)
                    _ = v16
                } else {
                    var v18 uint = (v12 % ((v12) | 1))
                    switch "foo bar" {
                    case "a":
                        var v19 string = (("foo bar" + (("" + "") + (("\n\t" + "\n\t") + ""))) + ("" + ("foo bar" + "a")))
                        _ = v19
                    default:
                        if (true || (((!(1e10 != -2.25)) || false) || ((false || (false || true)) && true))) {
                            v20 := GenS1{F0: (^(v1))}
                            v21 := GenS1{F0: v1}
                            v2 += v2
                            _ = v20
                            _ = v21
                        } else {
                            var v22 int64 = int64(fi14)
                            var v23 int16 = -32768
                            _ = v22
                            _ = v23
                        }
                        var v24 int64 = 42
                        var v25 int8 = -16
                        _ = v24
                        _ = v25
                    }
                    _ = v18
                }
                _ = fi14
                _ = v15
            }
            fmt.Println(v1)
        }
        _ = len(v3)
        fmt.Println(v2, v1, v5)
        _ = fi13
    }
    var v26 uint32 = v4
    switch "a" {
    case "hello":
        var v27 uint = v12
        _ = v27
    case "":
        switch "hello" {
        case "a":
            var v28 uint = (v12 + (v12 / (((^(v12))) | 1)))
            _ = v28
        case "hello":
            var v29 int8 = 42
            _ = v29
        case "foo bar":
            fmt.Println(v4, v1)
        }
    default:
        v5 = 128
    }
    v11 = append(v11, (v5 >> (((v12 & (v12 >> (((v12 & 277)) & 63)))) & 63)))
    if (!true) {
        _ = len(v3)
    }
    if len(v11) > 0 {
        v11[0] = (^(v5))
    }
    v2 += 1.5
    if (false && true) {
        for fi30 := 0; fi30 < 13; fi30++ {
            if false {
                for fi31 := 0; fi31 < 10; fi31++ {
                    fmt.Println(v1, fi30)
                    _ = fi31
                }
                delete(v3, (v5 ^ (^(v5))))
            } else {
                fmt.Println(v4, v2, v12)
            }
            for fi32 := 0; fi32 < 5; fi32++ {
                switch fi32 {
                case 0:
                    var v33 uint64 = 0
                    _ = v33
                case 219:
                    if (255 == 107) {
                        if true {
                            if true {
                                v1 += 366
                            }
                            v34 := GenS2{F0: "\n\t"}
                            _ = v34
                        }
                        v35 := GenS2{F0: ("hello" + "\n\t")}
                        if false {
                            if len(v11) > 0 {
                                v11[0] = 0
                            }
                            delete(v3, fi32)
                        }
                        var v36 int = v5
                        v37 := GenS1{F0: (v1 & 1)}
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                }
                v3[fi32] = "hello"
                _ = fi32
            }
            for fi38 := 0; fi38 < 7; fi38++ {
                v2 -= (float32(1e10) - (v2 - v2))
                _ = fi38
            }
            _ = fi30
        }
    }
    switch "" {
    case "\n\t":
        if false {
            if (v2 < (v2 + 3.14159)) {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                var v39 int8 = 0
                _ = v39
            } else {
                var v40 int64 = -751
                v41 := v3[(v5 << ((v12) & 63))]
                _ = v40
                _ = v41
            }
            var v42 int8 = -128
            fmt.Println(v42, v42, v26)
            _ = v42
        }
    case "foo bar":
        if false {
            for fi43 := 0; fi43 < 13; fi43++ {
                for fi44 := 0; fi44 < 15; fi44++ {
                    var v45 int8 = 127
                    if true {
                        v4 = v26
                    } else {
                        switch (fi44 << (((v12 >> ((v12) & 63))) & 63)) {
                        case 1:
                            switch (fi43 + (fi44 / (((^(fi44))) | 1))) {
                            case 370:
                                var v46 uint16 = 42
                                _ = v46
                            case -1:
                                v45 += (v45 / (((v45 >> ((v12) & 63))) | 1))
                            }
                        case 735:
                            var v47 bool = false
                            _ = v47
                        case 4294967295:
                            v11 = append(v11, fi44)
                        default:
                            if (!false) {
                                fmt.Println(v4)
                            } else {
                                var v48 int8 = (v45 << (((v12 - (v12 % ((v12) | 1)))) & 63))
                                v49 := GenS2{F0: "\n\t"}
                                var v50 uint16 = v1
                                _ = v48
                                _ = v49
                                _ = v50
                            }
                        }
                        v1 = 42
                        v45 += 1
                        switch -2147483648 {
                        case -759:
                            v5 -= (fi43 + (fi44 - (fi43 | (fi44 >> ((v12) & 63)))))
                            if false {
                                if len(v11) > 0 {
                                    v11[0] = fi43
                                }
                            } else {
                                var v51 uint64 = 4294967295
                                v52 := GenS2{F0: "foo bar"}
                                if (841 > (^(fi43))) {
                                    if ((v1 % ((genFn0(false, -9223372036854775808, false)) | 1)) != (v1 >> (((v12 - v12)) & 63))) {
                                        v11 = append(v11, (-(fi44)))
                                    }
                                    v53 := v3[fi43]
                                    _ = v53
                                }
                                v54 := GenS1{F0: (^(v1))}
                                _ = v51
                                _ = v52
                                _ = v54
                            }
                        }
                    }
                    _ = fi44
                    _ = v45
                }
                if (65535 >= 18446744073709551615) {
                    v4 = (v26 & (v4 + v26))
                    switch ("foo bar" + ("a" + ("foo bar" + (((("\n\t" + "foo bar") + "\n\t") + "foo bar") + ("hello" + "\n\t"))))) {
                    case "\n\t":
                        var v55 uint = (v12 * (v12 / ((v12) | 1)))
                        _ = v55
                    case "hello":
                        v11 = append(v11, (-(fi43)))
                    case "foo bar":
                        var v56 uint = (v12 ^ uint(v4))
                        _ = v56
                    }
                    if len(v11) > 0 {
                        v11[0] = fi43
                    }
                }
                var v57 int32 = -1
                _ = fi43
                _ = v57
            }
        } else {
            var v58 uint64 = 4294967295
            _ = v58
        }
    case "a":
        for fi59 := 0; fi59 < 14; fi59++ {
            for fi60 := 0; fi60 < 6; fi60++ {
                v61 := GenS3{F0: 4294967295}
                if true {
                    v26 += 4294967295
                } else {
                    if (!(!((v5 >> ((65535) & 63)) <= (v5 >> (((v12 << ((0) & 63))) & 63))))) {
                        var v62 bool = false
                        v26 -= v26
                        _ = v62
                    } else {
                        var v63 uint16 = (v1 >> (((^(v12))) & 63))
                        if true {
                            var v64 bool = ((1e10 * 1.5) != 1e10)
                            var v65 string = ""
                            _ = v64
                            _ = v65
                        } else {
                            var v66 int16 = 844
                            var v67 int16 = v66
                            _ = v66
                            _ = v67
                        }
                        _ = v63
                    }
                    v68 := GenS4{F0: (fi60 - fi60), F1: 127}
                    if false {
                        v69 := GenS3{F0: 42}
                        v70 := GenS5{F0: 42}
                        v2 = v2
                        if ((!false) || true) {
                            v5 -= (fi60 ^ fi60)
                            if false {
                                v1 = (v1 % ((v1) | 1))
                            } else {
                                v71 := GenS6{F0: 4294967295, F1: v12}
                                _ = v71
                            }
                        }
                        if (!(((((255 == (v1 - v1)) && (-935 <= 1)) || false) && (!false)) && true)) {
                            switch ("hello" + "\n\t") {
                            case "":
                                if (!(!(!(0 >= 222)))) {
                                    if ((62 == 255) || false) {
                                        if true {
                                            switch 32767 {
                                            case -83:
                                                v12 = v12
                                            case -273:
                                                v12 -= 18446744073709551615
                                            }
                                            v2 = v2
                                            v72 := GenS5{F0: 42}
                                            _ = v72
                                        } else {
                                            var v73 int8 = 100
                                            _ = v73
                                        }
                                        if (18446744073709551615 >= 18446744073709551615) {
                                            v12 = v12
                                            var v74 uint = 2
                                            var v75 string = ""
                                            _ = v74
                                            _ = v75
                                        } else {
                                            _ = len(v11)
                                            v4 -= (v26 << ((4294967295) & 63))
                                        }
                                        _ = len(v11)
                                    }
                                }
                            }
                            var v76 int32 = -713
                            if ((!true) || (4294967295 > 18446744073709551615)) {
                                var v77 string = "a"
                                switch fi59 {
                                case -1:
                                    v11 = append(v11, (fi59 << ((v12) & 63)))
                                    v78 := v3[v5]
                                    _ = v78
                                default:
                                    switch 9223372036854775807 {
                                    case -1:
                                        if (false && (v2 <= 1e10)) {
                                            var v79 int8 = 100
                                            v80 := GenS7{F0: 4294967295, F1: (v12 << ((v12) & 63)), F2: 0}
                                            _ = v79
                                            _ = v80
                                        } else {
                                            fmt.Println(v2, v12, v76)
                                        }
                                    case 32767:
                                        v76 = (v76 - v76)
                                    default:
                                        if (!(-889 >= v76)) {
                                            var v81 uint64 = 18446744073709551615
                                            v82 := GenS1{F0: v1}
                                            var v83 string = "\n\t"
                                            _ = v81
                                            _ = v82
                                            _ = v83
                                        } else {
                                            if (!(0 < (v12 + 0))) {
                                                var v84 uint64 = 1
                                                v85 := GenS5{F0: 255}
                                                _ = v84
                                                _ = v85
                                            } else {
                                                fmt.Println(v12)
                                                v12 -= (v12 * 0)
                                                var v86 uint8 = 1
                                                _ = v86
                                            }
                                            v77 = v77
                                            v87 := GenS4{F0: (v5 / ((int(v76)) | 1)), F1: -50}
                                            _ = v87
                                        }
                                    }
                                }
                                _ = v77
                            }
                            _ = v76
                        } else {
                            v3[int(v1)] = "foo bar"
                            v12 -= (v12 / (((v12 / ((v12) | 1))) | 1))
                            if (v5 < (^(fi59))) {
                                if true {
                                    switch fi59 {
                                    case -1:
                                        var v88 int = 42
                                        v89 := GenS2{F0: "\n\t"}
                                        var v90 uint32 = 1
                                        _ = v88
                                        _ = v89
                                        _ = v90
                                    case 9223372036854775807:
                                        var v91 uint32 = (v26 & (v26 - v4))
                                        _ = v91
                                    case -890:
                                        fmt.Println(v12, fi60)
                                    }
                                    var v92 uint64 = 2
                                    fmt.Println(v5, v4, fi60)
                                    _ = v92
                                } else {
                                    if true {
                                        var v93 bool = true
                                        var v94 uint16 = (v1 | v1)
                                        v11 = append(v11, fi59)
                                        _ = v93
                                        _ = v94
                                    } else {
                                        switch "a" {
                                        case "":
                                            switch fi60 {
                                            case 127:
                                                if true {
                                                    if false {
                                                        if (false && false) {
                                                            var v95 float64 = 3.14159
                                                            if ((v1 < v1) || (float32((0.1 * (0.1 + (v95 * v95)))) <= -2.25)) {
                                                                var v96 int8 = 100
                                                                var v97 string = (("foo bar" + "\n\t") + "foo bar")
                                                                v98 := GenS8{F0: int64(v96), F1: 449}
                                                                _ = v96
                                                                _ = v97
                                                                _ = v98
                                                            } else {
                                                                switch v5 {
                                                                case 2147483647:
                                                                    v99 := GenS1{F0: (v1 * (v1 * genFn0(true, 9223372036854775807, true)))}
                                                                    _ = v99
                                                                default:
                                                                    v100 := GenS9{F0: v4}
                                                                    _ = v100
                                                                }
                                                                v101 := v3[(^(v5))]
                                                                v102 := GenS10{F0: 32767, F1: 0, F2: 9223372036854775807}
                                                                _ = v101
                                                                _ = v102
                                                            }
                                                            _ = v95
                                                        }
                                                        if ("foo bar" >= "\n\t") {
                                                            var v103 uint64 = 0
                                                            _ = v103
                                                        }
                                                        switch (v5 & 904) {
                                                        case 32767:
                                                            fmt.Println(v12, v2)
                                                        case 0:
                                                            var v104 float32 = (((float32(0.0) + v2) - (v2 * v2)) + float32(1.5))
                                                            var v105 uint8 = 1
                                                            if (!false) {
                                                                if ((!true) && (false || true)) {
                                                                    if (-1 >= int64(v12)) {
                                                                        switch fi60 {
                                                                        case 128:
                                                                            var v106 uint = v12
                                                                            _ = v106
                                                                        }
                                                                    }
                                                                    var v107 float32 = v2
                                                                    v108 := v3[v5]
                                                                    _ = v107
                                                                    _ = v108
                                                                }
                                                            } else {
                                                                v109 := GenS6{F0: 1, F1: v12}
                                                                _ = v109
                                                            }
                                                            _ = v104
                                                            _ = v105
                                                        case 127:
                                                            v3[(fi60 >> ((v12) & 63))] = "foo bar"
                                                        }
                                                    } else {
                                                        if true {
                                                            v26 -= v26
                                                            v110 := v3[(v5 + (v5 & fi60))]
                                                            if true {
                                                                v110 = v110
                                                                _ = len(v11)
                                                                _ = len(v3)
                                                            }
                                                            _ = v110
                                                        } else {
                                                            switch "" {
                                                            case "foo bar":
                                                                var v111 uint = (v12 ^ v12)
                                                                _ = v111
                                                            case "a":
                                                                var v112 uint64 = 0
                                                                _ = v112
                                                            case "hello":
                                                                v113 := GenS11{F0: 65535, F1: ("" + "")}
                                                                _ = v113
                                                            }
                                                            if true {
                                                                v2 -= v2
                                                                var v114 float64 = (float64(((0.0 / v2) / v2)) + 3.14159)
                                                                v115 := GenS10{F0: 0, F1: -32768, F2: 9223372036854775807}
                                                                _ = v114
                                                                _ = v115
                                                            }
                                                            if (!(4294967295 > 4294967295)) {
                                                                v4 = (v26 - 2147483647)
                                                            } else {
                                                                v116 := GenS12{F0: 42, F1: "hello"}
                                                                var v117 int = fi59
                                                                _ = v116
                                                                _ = v117
                                                            }
                                                        }
                                                        v26 += (v4 * 839)
                                                    }
                                                } else {
                                                    v118 := GenS13{F0: 200}
                                                    v119 := GenS14{F0: -47, F1: v5, F2: int8(v1)}
                                                    v120 := GenS8{F0: 1, F1: 0}
                                                    fmt.Println(v4, v26)
                                                    v1 = genFn0((-2147483648 == int32(v12)), 805, true)
                                                    if len(v11) >= 2 {
                                                        v11 = v11[1:len(v11)]
                                                    }
                                                    _ = v118
                                                    _ = v119
                                                    _ = v120
                                                }
                                                var v121 float32 = 3.14159
                                                _ = v121
                                            case 716:
                                                var v122 bool = ((^(v1)) <= (^(v1)))
                                                _ = v122
                                            case -551:
                                                v12 += v12
                                                var v123 bool = (!(-340 < -15))
                                                _ = v123
                                            }
                                        default:
                                            var v124 uint8 = 69
                                            _ = v124
                                        }
                                    }
                                }
                            } else {
                                var v125 uint16 = v1
                                if (-2.25 <= (float64(((v2 * float32((float64(0.1) / 1.0))) / 1.0)) - (1.5 / 1.0))) {
                                    var v126 int16 = -1
                                    _ = v126
                                } else {
                                    var v127 bool = true
                                    _ = v127
                                }
                                v128 := GenS7{F0: 682, F1: v12, F2: 4294967295}
                                switch (v5 - fi59) {
                                case -965:
                                    fmt.Println(v4)
                                case 127:
                                    var v129 uint8 = 42
                                    _ = v129
                                case 9223372036854775807:
                                    v26 = v26
                                default:
                                    var v130 float64 = (1.5 * (0.0 - 1.5))
                                    _ = v130
                                }
                                v131 := GenS7{F0: (v12 << ((v12) & 63)), F1: v12, F2: 4294967295}
                                _ = v125
                                _ = v128
                                _ = v131
                            }
                        }
                        _ = v69
                        _ = v70
                    }
                    _ = v68
                }
                fmt.Println(fi59, v1)
                _ = fi60
                _ = v61
            }
            var v132 int = 42
            v133 := GenS7{F0: (v12 / ((885) | 1)), F1: v12, F2: 1}
            _ = fi59
            _ = v132
            _ = v133
        }
    default:
        fmt.Println(v2, v4)
    }
    v134 := GenS1{F0: (^(v1))}
    v1 = (v1 << (((^(v12))) & 63))
    fmt.Println(v12)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v134
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
