package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 bool; F1 uint; F2 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 int8; F1 string; F2 string }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint64; F1 uint; F2 int }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int64; F1 int8; F2 uint16 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint16; F1 uint16; F2 uint8 }
func genFn0() uint32 {
    switch 128 {
    case 127:
        if (42 >= 42) {
            v5 := []int{32767}
            for fi6 := 0; fi6 < 3; fi6++ {
                var v7 int16 = 1
                for fi8 := 0; fi8 < 6; fi8++ {
                    if len(v5) > 0 {
                        v5[0] = (^(fi8))
                    }
                    v9 := GenS1{F0: 0, F1: fi8}
                    _ = fi8
                    _ = v9
                }
                var v10 int32 = 2147483647
                _ = fi6
                _ = v10
                _ = v7
            }
            _ = v5
        } else {
            if (200 <= 127) {
                var v11 uint64 = 1
                _ = v11
            } else {
                for fi12 := 0; fi12 < 4; fi12++ {
                    if (-1 != -1) {
                        v13 := GenS1{F0: -1, F1: (fi12 & fi12)}
                        _ = v13
                    }
                    _ = fi12
                }
                if true {
                    var v14 int64 = -185
                    switch -1 {
                    case 42:
                        var v15 int16 = 1
                        _ = v15
                    case 32767:
                        fmt.Println(v14)
                    }
                    v16 := []int{-314, int(v14), 450}
                    _ = v14
                    _ = v16
                } else {
                    var v17 int = 128
                    _ = v17
                }
            }
            for fi18 := 0; fi18 < 20; fi18++ {
                fmt.Println(fi18)
                v19 := GenS1{F0: 984, F1: (fi18 | (fi18 ^ (fi18 << ((255) & 63))))}
                _ = fi18
                _ = v19
            }
        }
    case -271:
        if true {
            v20 := []int{}
            _ = v20
        }
    case -702:
        var v21 int64 = -540
        _ = v21
    }
    var v22 int8 = -128
    for fi23 := 0; fi23 < 5; fi23++ {
        if (!false) {
            for fi24 := 0; fi24 < 8; fi24++ {
                v25 := GenS1{F0: -804, F1: fi23}
                fmt.Println(fi24)
                _ = fi24
                _ = v25
            }
            v26 := []int{fi23}
            _ = v26
        }
        _ = fi23
    }
    switch 1 {
    case -467:
        if false {
            v27 := GenS1{F0: 32767, F1: -92}
            var v28 int32 = -2147483648
            v22 = v22
            _ = v27
            _ = v28
        } else {
            for fi29 := 0; fi29 < 6; fi29++ {
                v22 += (v22 << ((846) & 63))
                if (!((("" + "foo bar") != ("" + (("" + "") + ("" + "")))) && false)) {
                    var v30 int32 = -552
                    v31 := []int{(^(fi29))}
                    for fi32 := 0; fi32 < 11; fi32++ {
                        if (!true) {
                            var v33 string = "hello"
                            _ = v33
                        }
                        _ = fi32
                    }
                    _ = v30
                    _ = v31
                }
                _ = fi29
            }
        }
    case 128:
        v34 := GenS1{F0: 814, F1: 128}
        _ = v34
    }
    for fi35 := 0; fi35 < 17; fi35++ {
        if (1 == v22) {
            v22 = v22
            if (false && true) {
                for fi36 := 0; fi36 < 6; fi36++ {
                    if (false && true) {
                        v22 += (-(v22))
                        switch "" {
                        case "foo bar":
                            v37 := GenS2{F0: false, F1: 622, F2: -330}
                            _ = v37
                        default:
                            v38 := []int{0}
                            _ = v38
                        }
                        if true {
                            if (true || true) {
                                fmt.Println(v22)
                                switch "\n\t" {
                                case "foo bar":
                                    var v39 int64 = 0
                                    _ = v39
                                case "":
                                    if false {
                                        fmt.Println(fi36, fi36)
                                        var v40 int16 = -930
                                        v40 -= v40
                                        _ = v40
                                    } else {
                                        v41 := []int{(fi36 / ((fi36) | 1))}
                                        v42 := GenS2{F0: ("foo bar" < "\n\t"), F1: 1, F2: (fi35 << ((255) & 63))}
                                        _ = v41
                                        _ = v42
                                    }
                                case "a":
                                    if true {
                                        fmt.Println(fi35)
                                        v43 := GenS3{F0: 2}
                                        var v44 string = ""
                                        _ = v43
                                        _ = v44
                                    } else {
                                        if ((0.0 < 0.0) && ((126 != 1) && (("" + "") >= "hello"))) {
                                            v45 := GenS4{F0: fi36, F1: (!true)}
                                            var v46 string = ""
                                            var v47 uint = 255
                                            _ = v45
                                            _ = v46
                                            _ = v47
                                        } else {
                                            v22 += v22
                                            v48 := GenS2{F0: (!false), F1: 914, F2: fi35}
                                            _ = v48
                                        }
                                        if true {
                                            v49 := GenS5{F0: (^(v22)), F1: "\n\t", F2: ("hello" + "hello")}
                                            _ = v49
                                        }
                                        v50 := []int{32767, (fi36 ^ 2147483647), -2147483648}
                                        _ = v50
                                    }
                                }
                            } else {
                                v22 -= (v22 + v22)
                            }
                            switch (fi35 % (((fi35 * 42)) | 1)) {
                            case 0:
                                v51 := GenS4{F0: 0, F1: true}
                                _ = v51
                            }
                            v52 := GenS4{F0: 32767, F1: (v22 != v22)}
                            v53 := GenS3{F0: 2}
                            _ = v52
                            _ = v53
                        } else {
                            fmt.Println(fi35, fi36)
                            v22 += (v22 / ((v22) | 1))
                        }
                    }
                    if (-32768 >= -32768) {
                        var v54 int = fi35
                        switch v54 {
                        case -729:
                            if (-9223372036854775808 <= 9223372036854775807) {
                                var v55 bool = ((39 <= 127) || false)
                                v56 := GenS3{F0: 42}
                                var v57 int = -77
                                _ = v55
                                _ = v56
                                _ = v57
                            }
                        case 592:
                            switch 32768 {
                            case -548:
                                v22 += 127
                            default:
                                var v58 int8 = v22
                                _ = v58
                            }
                        case 4294967295:
                            v54 -= (v54 - (^(v54)))
                        default:
                            var v59 uint = 120
                            _ = v59
                        }
                        _ = v54
                    }
                    var v60 uint = 255
                    v60 -= (v60 | v60)
                    _ = fi36
                    _ = v60
                }
            } else {
                v61 := GenS6{F0: 65535}
                var v62 int64 = -97
                v63 := GenS7{F0: 2, F1: 2, F2: fi35}
                _ = v61
                _ = v62
                _ = v63
            }
        } else {
            fmt.Println(fi35, v22)
            var v64 int8 = v22
            _ = v64
        }
        v22 += v22
        _ = fi35
    }
    _ = v22
    return 65535
}
func genFn1(p0 int, p1 int64, p2 uint32) uint32 {
    switch "a" {
    case "\n\t":
        switch "\n\t" {
        case "foo bar":
            v65 := []int{(p0 + 2147483647), 128, p0}
            _ = v65
        }
    case "foo bar":
        v66 := []int{(^(p0)), 32767, p0}
        _ = v66
    case "a":
        switch p0 {
        case -483:
            var v67 int = 1
            _ = v67
        default:
            p2 -= p2
        }
    }
    var v68 uint16 = 42
    for fi69 := 0; fi69 < 12; fi69++ {
        switch ("foo bar" + "") {
        case "\n\t":
            var v70 int16 = 0
            _ = v70
        case "a":
            for fi71 := 0; fi71 < 6; fi71++ {
                p0 = (^(p0))
                var v72 float64 = 1.5
                p2 += genFn0()
                _ = fi71
                _ = v72
            }
        default:
            var v73 uint64 = 164
            _ = v73
        }
        _ = fi69
    }
    var v74 string = ("" + ("hello" + "a"))
    _ = p0
    _ = p1
    _ = p2
    _ = v68
    _ = v74
    return 42
}

func main() {
    v1 := []uint32{0, 1, 65535, 527}
    var v2 uint64 = 2
    var v3 int32 = -317
    var v4 uint64 = (v2 | v2)
    if (-1 >= 1) {
        _ = len(v1)
        _ = len(v1)
    } else {
        if false {
            for fi75 := 0; fi75 < 3; fi75++ {
                v2 = 746
                _ = fi75
            }
        }
        var v76 int = -2147483648
        if false {
            var v77 int8 = -103
            v2 = 42
            v78 := GenS8{F0: 2}
            _ = v77
            _ = v78
        } else {
            if false {
                v76 = 32768
                var v79 float64 = 1e10
                _ = v79
            }
        }
        _ = v76
    }
    var v80 float32 = 0.0
    for fi81 := 0; fi81 < 1; fi81++ {
        v1 = append(v1, 481)
        v82 := GenS6{F0: 313}
        _ = fi81
        _ = v82
    }
    v3 += v3
    v80 -= v80
    if (32767 == -32768) {
        for fi83 := 0; fi83 < 1; fi83++ {
            switch ("" + "") {
            case "hello":
                var v84 int = fi83
                _ = v84
            case "":
                for fi85 := 0; fi85 < 4; fi85++ {
                    v3 -= (v3 << ((18446744073709551615) & 63))
                    if false {
                        v80 = v80
                        if true {
                            if ((("foo bar" != "hello") && (128 > 127)) && (!false)) {
                                v86 := GenS2{F0: true, F1: 65535, F2: (fi85 / ((fi85) | 1))}
                                _ = v86
                            }
                            switch int(v4) {
                            case 128:
                                if false {
                                    v2 -= 42
                                    var v87 uint32 = 4294967295
                                    _ = len(v1)
                                    v88 := GenS2{F0: false, F1: 660, F2: (fi83 + 207)}
                                    _ = v87
                                    _ = v88
                                } else {
                                    v89 := GenS9{F0: 810, F1: -128, F2: 42}
                                    v90 := GenS10{F0: 2147483647}
                                    v3 += (v3 * (v3 - v3))
                                    _ = v89
                                    _ = v90
                                }
                            case 32768:
                                if (!false) {
                                    var v91 uint32 = 65535
                                    v4 = 971
                                    _ = v91
                                }
                            default:
                                if false {
                                    var v92 int8 = 1
                                    _ = v92
                                }
                            }
                            switch ("a" + "\n\t") {
                            case "\n\t":
                                var v93 bool = ((((fi85 << ((338) & 63)) < fi85) || (true || true)) || true)
                                _ = v93
                            case "":
                                fmt.Println(fi83, v80)
                            case "foo bar":
                                if (fi85 == (fi85 % ((fi83) | 1))) {
                                    if false {
                                        v2 = 42
                                        if true {
                                            if ((v2 / (((v4 | v4)) | 1)) >= 4294967295) {
                                                switch "" {
                                                case "foo bar":
                                                    v4 = 1
                                                }
                                                var v94 int8 = 42
                                                _ = v94
                                            }
                                            var v95 uint = 0
                                            _ = v95
                                        }
                                        switch "foo bar" {
                                        case "foo bar":
                                            v80 = v80
                                        }
                                    }
                                } else {
                                    var v96 int8 = 0
                                    v96 = v96
                                    if false {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v4 = (^(v4))
                                        fmt.Println(fi85, v3)
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v80 += ((v80 * v80) * (v80 / v80))
                                    }
                                    _ = v96
                                }
                            default:
                                v97 := GenS7{F0: 989, F1: 1, F2: 127}
                                v3 -= v3
                                fmt.Println(v2, v2, v4)
                                _ = v97
                            }
                        } else {
                            switch "a" {
                            case "":
                                if true {
                                    var v98 int32 = v3
                                    switch (fi83 >> ((411) & 63)) {
                                    case 0:
                                        switch "\n\t" {
                                        case "foo bar":
                                            v99 := GenS2{F0: false, F1: 0, F2: 32767}
                                            _ = v99
                                        case "a":
                                            v100 := GenS4{F0: fi85, F1: false}
                                            _ = v100
                                        case "hello":
                                            var v101 int64 = 1
                                            _ = v101
                                        default:
                                            v2 = v2
                                        }
                                    case 2147483647:
                                        var v102 int8 = -126
                                        _ = v102
                                    case 42:
                                        fmt.Println(v98, v80)
                                        var v103 uint32 = genFn0()
                                        v103 = (v103 ^ (v103 / (((v103 - v103)) | 1)))
                                        var v104 int32 = 332
                                        _ = v103
                                        _ = v104
                                    }
                                    var v105 string = "\n\t"
                                    _ = v105
                                    _ = v98
                                }
                            case "\n\t":
                                v2 += v2
                            case "a":
                                if ((v2 >= 18446744073709551615) && (true && true)) {
                                    fmt.Println(v4, v4)
                                    if len(v1) > 0 {
                                        v1[0] = 0
                                    }
                                } else {
                                    var v106 uint32 = 154
                                    v107 := GenS4{F0: 4294967295, F1: false}
                                    _ = v106
                                    _ = v107
                                }
                            }
                        }
                        v108 := GenS7{F0: (^(v2)), F1: 4294967295, F2: (-(fi85))}
                        _ = v108
                    }
                    switch "hello" {
                    case "\n\t":
                        v109 := GenS11{F0: 8}
                        _ = v109
                    case "foo bar":
                        var v110 int64 = 0
                        _ = v110
                    case "":
                        if false {
                            v4 = v2
                        }
                    default:
                        var v111 int32 = (v3 - (v3 * (v3 + v3)))
                        _ = v111
                    }
                    _ = fi85
                }
            case "a":
                var v112 uint = 362
                _ = v112
            }
            v2 += v4
            var v113 uint16 = 0
            _ = fi83
            _ = v113
        }
    } else {
        switch "hello" {
        case "foo bar":
            for fi114 := 0; fi114 < 3; fi114++ {
                for fi115 := 0; fi115 < 13; fi115++ {
                    switch fi114 {
                    case 980:
                        v116 := GenS12{F0: -9223372036854775808}
                        _ = v116
                    case -1:
                        v1 = append(v1, 219)
                    }
                    var v117 int = fi115
                    _ = fi115
                    _ = v117
                }
                var v118 int = -419
                fmt.Println(v80)
                _ = fi114
                _ = v118
            }
        case "a":
            v2 += 4294967295
        case "\n\t":
            for fi119 := 0; fi119 < 10; fi119++ {
                v1 = append(v1, genFn1(32767, 1, genFn0()))
                v120 := GenS13{F0: (-158 < 1)}
                if len(v1) > 0 {
                    v1[0] = 12
                }
                _ = fi119
                _ = v120
            }
        }
        switch "hello" {
        case "\n\t":
            for fi121 := 0; fi121 < 11; fi121++ {
                if (817 > 4294967295) {
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                    if true {
                        for fi122 := 0; fi122 < 9; fi122++ {
                            var v123 uint8 = 127
                            v4 -= 18446744073709551615
                            v123 -= v123
                            switch "foo bar" {
                            case "a":
                                if ((1.5 + 3.14159) > 0.0) {
                                    fmt.Println(fi122, v4)
                                    switch "foo bar" {
                                    case "a":
                                        v3 += (^(v3))
                                    case "hello":
                                        if (100 > 0) {
                                            switch "foo bar" {
                                            case "a":
                                                v123 += 1
                                            default:
                                                _ = len(v1)
                                            }
                                            v124 := GenS3{F0: v4}
                                            _ = v124
                                        } else {
                                            var v125 int = 513
                                            fmt.Println(v2)
                                            if false {
                                                var v126 uint = 1
                                                v3 -= 2147483647
                                                _ = v126
                                            }
                                            var v127 uint = 0
                                            _ = v125
                                            _ = v127
                                        }
                                    case "":
                                        _ = len(v1)
                                    }
                                    fmt.Println(fi121, v2)
                                }
                            }
                            _ = fi122
                            _ = v123
                        }
                    } else {
                        for fi128 := 0; fi128 < 12; fi128++ {
                            v4 = (v4 * 979)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi128
                        }
                        v129 := GenS5{F0: 1, F1: ("hello" + ""), F2: "a"}
                        _ = v129
                    }
                    for fi130 := 0; fi130 < 12; fi130++ {
                        var v131 uint16 = 0
                        var v132 int64 = 291
                        _ = fi130
                        _ = v131
                        _ = v132
                    }
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                for fi133 := 0; fi133 < 5; fi133++ {
                    v3 += v3
                    if ((255 <= 65535) && false) {
                        v4 += v2
                        v134 := GenS14{F0: 32767, F1: 65535, F2: 42}
                        _ = v134
                    }
                    _ = fi133
                }
                _ = fi121
            }
        default:
            _ = len(v1)
        }
    }
    fmt.Println(v2)
    var v135 int = -23
    fmt.Println(v3, v80)
    var v136 uint16 = uint16(v4)
    if false {
        fmt.Println(v2, v4, v136)
        var v137 int = -632
        _ = v137
    }
    v2 -= (v4 / (((^(v2))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v135:", v135)
    fmt.Println("v136:", v136)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v135
    _ = v136
    _ = v2
    _ = v3
    _ = v4
    _ = v80
}
