package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint; F1 uint64 }
type GenS7 struct { F0 int; F1 string; F2 int64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 bool; F1 int32 }
type GenS10 struct { F0 uint32; F1 uint8 }
func genFn0() int {
    v4 := []int{32767, 127, 975}
    if (-26 > 30) {
        var v5 string = ""
        for fi6 := 0; fi6 < 16; fi6++ {
            fmt.Println(fi6, v5)
            if false {
                if (!true) {
                    for fi7 := 0; fi7 < 19; fi7++ {
                        fmt.Println(fi7, fi6)
                        if (v5 <= "\n\t") {
                            var v8 uint32 = 304
                            _ = v8
                        } else {
                            fmt.Println(fi6, v5)
                            fmt.Println(v5, fi7, fi7)
                        }
                        _ = fi7
                    }
                }
                fmt.Println(v5, fi6)
                var v9 uint = 65535
                _ = v9
            }
            v10 := GenS1{F0: 65}
            _ = fi6
            _ = v10
        }
        fmt.Println(v5)
        _ = v5
    } else {
        if (1 > 42) {
            if true {
                if false {
                    var v11 string = "\n\t"
                    switch -209 {
                    case 32767:
                        if true {
                            fmt.Println(v11)
                            _ = len(v4)
                            if (884 <= -920) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            } else {
                                v11 = v11
                                v12 := GenS2{F0: 774}
                                _ = v12
                            }
                        }
                    case 32768:
                        v11 = v11
                    }
                    _ = v11
                }
                if (32767 <= 32767) {
                    var v13 uint16 = 94
                    var v14 string = "a"
                    _ = v13
                    _ = v14
                }
                for fi15 := 0; fi15 < 13; fi15++ {
                    var v16 uint = 65535
                    _ = fi15
                    _ = v16
                }
            }
            var v17 uint = 330
            switch 127 {
            case 128:
                v18 := GenS2{F0: 803}
                _ = v18
            case -9223372036854775808:
                v19 := GenS2{F0: 4294967295}
                _ = v19
            default:
                for fi20 := 0; fi20 < 15; fi20++ {
                    fmt.Println(v17)
                    for fi21 := 0; fi21 < 19; fi21++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v22 uint8 = 255
                        fmt.Println(v17)
                        _ = fi21
                        _ = v22
                    }
                    _ = fi20
                }
            }
            _ = v17
        }
        if (0 < 128) {
            var v23 uint8 = 64
            v23 = (v23 & 70)
            _ = v23
        } else {
            if (-1 > -9223372036854775808) {
                for fi24 := 0; fi24 < 19; fi24++ {
                    for fi25 := 0; fi25 < 17; fi25++ {
                        _ = len(v4)
                        _ = len(v4)
                        _ = fi25
                    }
                    v26 := GenS1{F0: 1}
                    var v27 int64 = 179
                    _ = fi24
                    _ = v26
                    _ = v27
                }
                for fi28 := 0; fi28 < 16; fi28++ {
                    for fi29 := 0; fi29 < 3; fi29++ {
                        var v30 float64 = ((float64(0.1) + ((0.0 / 1.0) + 1e10)) + -2.25)
                        fmt.Println(fi29, fi29)
                        _ = fi29
                        _ = v30
                    }
                    _ = fi28
                }
            } else {
                for fi31 := 0; fi31 < 15; fi31++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi31
                }
            }
        }
    }
    var v32 uint16 = 42
    var v33 int8 = -37
    if true {
        v33 += 42
        v33 -= (v33 - (v33 % ((127) | 1)))
        v33 += 53
    }
    _ = v32
    _ = v33
    _ = v4
    return 1
}
func genFn1() uint16 {
    var v34 string = ("a" + "a")
    if false {
        v35 := []int{32768, 0}
        if len(v35) > 0 {
            v35[0] = 756
        }
        _ = v35
    } else {
        var v36 int16 = 39
        _ = v36
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        if (3.14159 < 1e10) {
            var v38 uint32 = 16
            _ = v38
        }
        _ = fi37
    }
    switch 207 {
    case 32768:
        switch 127 {
        case 2147483647:
            v39 := GenS2{F0: 2}
            _ = v39
        case -2147483648:
            fmt.Println(v34)
        case 128:
            if false {
                for fi40 := 0; fi40 < 3; fi40++ {
                    var v41 int = genFn0()
                    _ = fi40
                    _ = v41
                }
            } else {
                switch v34 {
                case "a":
                    for fi42 := 0; fi42 < 11; fi42++ {
                        v43 := GenS3{F0: 42, F1: 0}
                        _ = fi42
                        _ = v43
                    }
                case "hello":
                    v34 = "hello"
                }
                v44 := GenS3{F0: 42, F1: 2147483647}
                _ = v44
            }
        }
    }
    v45 := []int{32767, 4294967295, 1, 4294967295}
    _ = v34
    _ = v45
    return 2
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{}
    v3 := make(map[int]uint32)
    var v46 int64 = -1
    var v47 float32 = 3.14159
    for fi48 := 0; fi48 < 17; fi48++ {
        switch fi48 {
        case 811:
            delete(v1, 2147483648)
        default:
            fmt.Println(v47, v47, v46)
        }
        if false {
            if false {
                var v49 bool = false
                var v50 uint32 = 1
                for fi51 := 0; fi51 < 10; fi51++ {
                    switch fi51 {
                    case 32767:
                        if (!false) {
                            switch genFn0() {
                            case 845:
                                v47 = v47
                            case 32768:
                                var v52 float32 = v47
                                _ = v52
                            case 0:
                                var v53 int16 = 502
                                _ = v53
                            default:
                                if v49 {
                                    switch 844 {
                                    case 32767:
                                        fmt.Println(v50, v50, v46)
                                    case 42:
                                        if v49 {
                                            if true {
                                                switch "\n\t" {
                                                case "":
                                                    v54 := GenS1{F0: -740}
                                                    v55 := GenS1{F0: v46}
                                                    _ = v54
                                                    _ = v55
                                                case "hello":
                                                    var v56 bool = (!((v46 >= v46) && false))
                                                    _ = v56
                                                case "\n\t":
                                                    v57 := GenS2{F0: 897}
                                                    _ = v57
                                                default:
                                                    if v49 {
                                                        fmt.Println(v47, fi48, v46)
                                                        fmt.Println(v47)
                                                        var v58 uint64 = 18446744073709551615
                                                        _ = v58
                                                    } else {
                                                        _ = len(v2)
                                                    }
                                                }
                                                if v49 {
                                                    var v59 float64 = 0.0
                                                    var v60 string = "\n\t"
                                                    _ = v59
                                                    _ = v60
                                                }
                                                v61 := GenS4{F0: 18446744073709551615, F1: 2}
                                                switch "\n\t" {
                                                case "\n\t":
                                                    var v62 float32 = v47
                                                    _ = v62
                                                }
                                                _ = v61
                                            }
                                            switch "a" {
                                            case "a":
                                                v3[genFn0()] = v50
                                            case "":
                                                v63 := GenS3{F0: 106, F1: -362}
                                                _ = v63
                                            default:
                                                if (v49 || true) {
                                                    delete(v1, (v50 ^ (v50 + v50)))
                                                    var v64 bool = (v49 || true)
                                                    _ = v64
                                                }
                                                switch "hello" {
                                                case "":
                                                    v47 = (v47 / v47)
                                                case "a":
                                                    var v65 uint32 = (v50 | (^(v50)))
                                                    _ = v65
                                                case "foo bar":
                                                    var v66 int32 = 42
                                                    _ = v66
                                                }
                                            }
                                        } else {
                                            fmt.Println(fi51, v49, v46)
                                            var v67 int64 = (v46 + (v46 >> ((0) & 63)))
                                            v46 = v46
                                            _ = v67
                                        }
                                    case -899:
                                        switch (fi48 * 42) {
                                        case 658:
                                            switch ("\n\t" + "a") {
                                            case "a":
                                                if (42 <= 255) {
                                                    if v49 {
                                                        var v68 uint64 = 4294967295
                                                        v69 := GenS1{F0: (v46 | v46)}
                                                        v68 += (v68 ^ (v68 ^ v68))
                                                        _ = v68
                                                        _ = v69
                                                    } else {
                                                        v49 = true
                                                    }
                                                    v46 -= v46
                                                    v50 = 2
                                                    v47 -= float32(1.5)
                                                    var v70 uint64 = 120
                                                    _ = v70
                                                }
                                            case "hello":
                                                v47 = v47
                                            }
                                        case 0:
                                            fmt.Println(fi51)
                                        default:
                                            v71 := GenS5{F0: 611}
                                            _ = v71
                                        }
                                    }
                                } else {
                                    delete(v1, 2147483647)
                                    if true {
                                        v72 := GenS5{F0: v50}
                                        _ = v72
                                    }
                                }
                                if v49 {
                                    var v73 int64 = (v46 >> ((2) & 63))
                                    if (((3.14159 / 1.0) >= ((1e10 / (0.1 / ((-2.25 + (1.5 * float64(float32(-2.25)))) + 3.14159))) - 1e10)) || v49) {
                                        var v74 int8 = -1
                                        _ = v74
                                    }
                                    _ = v73
                                }
                            }
                            var v75 int16 = 42
                            v50 = (^(v50))
                            _ = v75
                        } else {
                            switch "\n\t" {
                            case "hello":
                                fmt.Println(v49)
                            case "a":
                                v76 := GenS2{F0: 18446744073709551615}
                                _ = v76
                            case "":
                                var v77 bool = (!false)
                                switch 377 {
                                case 2147483647:
                                    v78 := v3[(^(fi51))]
                                    _ = v78
                                default:
                                    v1[65535] = fi51
                                }
                                _ = v77
                            default:
                                var v79 uint16 = 255
                                _ = v79
                            }
                            var v80 int16 = 1
                            _ = v80
                        }
                    case -1:
                        v50 = (v50 | (v50 << ((2) & 63)))
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    switch ("\n\t" + "foo bar") {
                    case "hello":
                        v81 := GenS6{F0: 189, F1: 4294967295}
                        _ = v81
                    case "\n\t":
                        if v49 {
                            v82 := GenS2{F0: 775}
                            _ = v82
                        }
                    }
                    _ = fi51
                }
                _ = v49
                _ = v50
            } else {
                for fi83 := 0; fi83 < 3; fi83++ {
                    var v84 string = "foo bar"
                    _ = fi83
                    _ = v84
                }
                fmt.Println(fi48, fi48, v47)
                v47 -= v47
            }
            switch (fi48 - genFn0()) {
            case 32768:
                if true {
                    var v85 int16 = 294
                    var v86 uint = 882
                    _ = v85
                    _ = v86
                }
            case 128:
                for fi87 := 0; fi87 < 12; fi87++ {
                    v88 := GenS6{F0: 1, F1: 1}
                    if true {
                        switch "\n\t" {
                        case "foo bar":
                            v46 += v46
                        case "a":
                            fmt.Println(fi48)
                        }
                        var v89 float64 = 0.0
                        var v90 uint16 = genFn1()
                        var v91 int8 = int8(v90)
                        _ = v89
                        _ = v90
                        _ = v91
                    } else {
                        v92 := GenS1{F0: (^(v46))}
                        _ = v92
                    }
                    var v93 uint64 = 4294967295
                    _ = fi87
                    _ = v88
                    _ = v93
                }
            }
            if (42 < -96) {
                var v94 int = (fi48 | genFn0())
                v95 := GenS4{F0: 4294967295, F1: 0}
                _ = v94
                _ = v95
            }
        }
        if false {
            switch fi48 {
            case 127:
                if false {
                    if false {
                        for fi96 := 0; fi96 < 3; fi96++ {
                            if false {
                                var v97 float64 = 0.0
                                fmt.Println(v47, fi96)
                                if ((1 == 0) && true) {
                                    if (true || (true || (!(286 >= 1)))) {
                                        var v98 uint64 = 4294967295
                                        if len(v2) > 0 {
                                            v2[0] = fi96
                                        }
                                        _ = v98
                                    }
                                    var v99 uint64 = 351
                                    v100 := GenS7{F0: fi96, F1: "", F2: v46}
                                    _ = v100
                                    _ = v99
                                }
                                _ = v97
                            } else {
                                fmt.Println(fi48)
                                v101 := GenS1{F0: (v46 / ((v46) | 1))}
                                _ = v101
                            }
                            var v102 uint32 = 2
                            fmt.Println(v102, v46)
                            _ = fi96
                            _ = v102
                        }
                        fmt.Println(fi48, v46, v47)
                        var v103 float32 = ((0.0 - v47) / v47)
                        _ = v103
                    } else {
                        v46 += 0
                        v3[556] = 1
                    }
                }
            case 9223372036854775807:
                v46 = v46
            case 32767:
                _ = len(v1)
            default:
                if (false && false) {
                    v104 := GenS7{F0: fi48, F1: "\n\t", F2: (v46 + (v46 >> ((637) & 63)))}
                    fmt.Println(v47)
                    v46 += v46
                    _ = v104
                } else {
                    fmt.Println(v46, v46, fi48)
                }
            }
            v105 := GenS7{F0: (fi48 / (((fi48 >> ((255) & 63))) | 1)), F1: "a", F2: v46}
            v106 := GenS8{F0: 1000}
            _ = v105
            _ = v106
        }
        _ = fi48
    }
    if (128 != 1) {
        switch "a" {
        case "a":
            fmt.Println(v46, v46)
        }
        var v107 uint16 = 255
        _ = v107
    }
    var v108 uint32 = 150
    if (100 <= -128) {
        v109 := GenS9{F0: (42 != 909), F1: 831}
        _ = v109
    }
    for fi110 := 0; fi110 < 3; fi110++ {
        v2 = append(v2, fi110)
        if true {
            for fi111 := 0; fi111 < 12; fi111++ {
                var v112 uint32 = 2147483647
                v113 := GenS5{F0: v108}
                _ = fi111
                _ = v112
                _ = v113
            }
            v114 := GenS1{F0: (-(v46))}
            _ = v114
        }
        for fi115 := 0; fi115 < 13; fi115++ {
            v47 += ((v47 + (0.1 / v47)) + (float32(0.0) * (float32(3.14159) / float32(0.1))))
            var v116 uint8 = 236
            v117 := GenS7{F0: fi110, F1: "\n\t", F2: int64(v108)}
            v116 += (v116 * v116)
            if false {
                switch (^(fi110)) {
                case 885:
                    var v118 int16 = 32767
                    _ = v118
                case 590:
                    fmt.Println(v46)
                }
                if true {
                    v46 = v46
                } else {
                    if (!("a" == "")) {
                        v119 := GenS9{F0: true, F1: -251}
                        var v120 uint = 557
                        var v121 float32 = v47
                        _ = v119
                        _ = v120
                        _ = v121
                    }
                    v116 = v116
                    var v122 int16 = 1000
                    _ = v122
                }
                _ = len(v1)
            }
            _ = fi115
            _ = v116
            _ = v117
        }
        _ = fi110
    }
    v108 = (v108 - (v108 ^ (v108 + (v108 * (v108 ^ 65535)))))
    for fi123 := 0; fi123 < 9; fi123++ {
        fmt.Println(v108, v108)
        for fi124 := 0; fi124 < 19; fi124++ {
            if true {
                var v125 int = (fi124 * fi124)
                _ = v125
            } else {
                var v126 uint64 = 2
                _ = v126
            }
            switch fi124 {
            case 0:
                var v127 int8 = 100
                _ = v127
            case 128:
                var v128 float64 = (0.0 + (3.14159 * 3.14159))
                _ = v128
            }
            _ = fi124
        }
        _ = fi123
    }
    fmt.Println(v108, v46, v108)
    switch genFn0() {
    case -908:
        v129 := GenS2{F0: 312}
        _ = v129
    case -253:
        if false {
            fmt.Println(v108, v108)
            v47 -= (v47 / (-2.25 * (v47 * float32(3.14159))))
            if false {
                fmt.Println(v46, v108)
            }
        } else {
            var v130 int16 = 0
            v47 -= 3.14159
            v130 = v130
            _ = v130
        }
    case 42:
        if false {
            v131 := GenS10{F0: (v108 ^ v108), F1: 229}
            v47 -= (v47 / v47)
            _ = v131
        } else {
            var v132 int8 = 9
            var v133 int16 = 0
            _ = v132
            _ = v133
        }
    default:
        var v134 int32 = -650
        _ = v134
    }
    var v135 int32 = -1
    _ = len(v2)
    v108 = 153
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v108:", v108)
    fmt.Println("v135:", v135)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v108
    _ = v135
    _ = v2
    _ = v3
    _ = v46
    _ = v47
}
