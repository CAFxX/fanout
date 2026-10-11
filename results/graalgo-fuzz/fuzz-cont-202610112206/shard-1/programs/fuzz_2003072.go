package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 int16; F1 int64; F2 uint64 }
type GenS7 struct { F0 uint8; F1 uint; F2 int16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int; F1 uint8; F2 uint16 }
type GenS11 struct { F0 uint; F1 uint }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int; F1 bool; F2 uint8 }
type GenS14 struct { F0 bool; F1 string }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 int; F1 bool; F2 uint32 }
type GenS18 struct { F0 uint64 }
type GenS19 struct { F0 uint; F1 int }
type GenS20 struct { F0 bool; F1 bool }
type GenS21 struct { F0 int; F1 int64; F2 uint8 }
type GenS22 struct { F0 int; F1 uint8; F2 int32 }
type GenS23 struct { F0 uint16 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    var v4 uint32 = 645
    if (true || (false && (!false))) {
        v4 += (^(v4))
        var v5 int16 = 1000
        _ = v5
    } else {
        v6 := GenS1{F0: 259}
        var v7 uint16 = 1
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        if false {
            v9 := GenS2{F0: 42, F1: "foo bar"}
            fmt.Println(v4)
            switch "" {
            case "a":
                var v10 int16 = 953
                _ = v10
            case "":
                if true {
                    for fi11 := 0; fi11 < 3; fi11++ {
                        var v12 float64 = -2.25
                        if (!(!(false || true))) {
                            var v13 uint32 = v4
                            _ = v13
                        } else {
                            var v14 uint = 4294967295
                            _ = v14
                        }
                        _ = fi11
                        _ = v12
                    }
                    var v15 uint8 = 196
                    _ = v15
                } else {
                    v4 = v4
                    for fi16 := 0; fi16 < 20; fi16++ {
                        var v17 int16 = 1
                        if (0.0 == 1.5) {
                            if true {
                                fmt.Println(v17, v17, fi16)
                            }
                        }
                        var v18 int16 = 42
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                }
            }
            _ = v9
        }
        _ = fi8
    }
    switch "\n\t" {
    case "a":
        v19 := []int{1}
        _ = v19
    case "foo bar":
        if false {
            var v20 uint = 42
            _ = v20
        }
    default:
        var v21 int64 = 1
        _ = v21
    }
    switch -438 {
    case -9223372036854775808:
        v4 -= (v4 ^ 1)
    case -261:
        var v22 uint8 = 116
        _ = v22
    }
    var v23 int16 = 0
    if (477 < 564) {
        switch "a" {
        case "a":
            if (!true) {
                if (833 <= 79) {
                    v23 = -688
                    if false {
                        var v24 int16 = 32767
                        _ = v24
                    } else {
                        fmt.Println(v23, v4)
                    }
                    var v25 uint8 = 164
                    _ = v25
                } else {
                    fmt.Println(v4, v23)
                    v26 := []int{-457, 1}
                    for fi27 := 0; fi27 < 11; fi27++ {
                        v4 = v4
                        fmt.Println(v23)
                        fmt.Println(fi27, v4, v23)
                        _ = fi27
                    }
                    _ = v26
                }
                if (!false) {
                    v28 := GenS3{F0: 153, F1: 1}
                    _ = v28
                }
                for fi29 := 0; fi29 < 10; fi29++ {
                    var v30 float32 = ((-2.25 - 1e10) - (0.1 - 3.14159))
                    v23 -= (^(v23))
                    _ = fi29
                    _ = v30
                }
            }
        case "hello":
            v31 := GenS2{F0: -610, F1: "a"}
            _ = v31
        default:
            v32 := GenS1{F0: 1}
            _ = v32
        }
        fmt.Println(v4, v23)
        v33 := GenS3{F0: 127, F1: uint(v4)}
        _ = v33
    } else {
        v4 = v4
        for fi34 := 0; fi34 < 6; fi34++ {
            var v35 float64 = 0.1
            _ = fi34
            _ = v35
        }
        v4 = (v4 << ((392) & 63))
    }
    if true {
        v36 := GenS4{F0: (("\n\t" + ("foo bar" + ("foo bar" + "a"))) + "hello"), F1: true}
        v4 -= (v4 & 4294967295)
        var v37 int8 = 127
        _ = v36
        _ = v37
    } else {
        v38 := GenS5{F0: "hello", F1: 127}
        v39 := GenS2{F0: -238, F1: ""}
        v40 := GenS6{F0: v23, F1: -782, F2: 0}
        _ = v38
        _ = v39
        _ = v40
    }
    fmt.Println(v4)
    v41 := GenS7{F0: 1, F1: 983, F2: (v23 % ((v23) | 1))}
    for fi42 := 0; fi42 < 14; fi42++ {
        v23 = (v23 | v23)
        v23 += -712
        var v43 float64 = (0.1 * 3.14159)
        _ = fi42
        _ = v43
    }
    switch -403 {
    case -2147483648:
        v4 += (v4 * v4)
    case 42:
        switch "hello" {
        case "a":
            v44 := GenS3{F0: 199, F1: 2}
            _ = v44
        case "foo bar":
            for fi45 := 0; fi45 < 7; fi45++ {
                var v46 uint32 = v4
                for fi47 := 0; fi47 < 8; fi47++ {
                    v23 = 32767
                    var v48 uint16 = 255
                    if (58 != 92) {
                        var v49 int64 = 0
                        _ = v49
                    }
                    _ = fi47
                    _ = v48
                }
                _ = fi45
                _ = v46
            }
        case "hello":
            if true {
                var v50 float64 = ((0.0 - -2.25) * 0.1)
                v51 := v1[42]
                for fi52 := 0; fi52 < 13; fi52++ {
                    v53 := v2[fi52]
                    if (v51 <= v51) {
                        if true {
                            if false {
                                var v54 float32 = 1e10
                                var v55 int64 = 42
                                for fi56 := 0; fi56 < 6; fi56++ {
                                    fmt.Println(fi56)
                                    var v57 uint16 = 20
                                    switch fi56 {
                                    case 0:
                                        var v58 int = (fi52 - 2147483647)
                                        v55 = (v55 + (-(v55)))
                                        _ = v58
                                    default:
                                        delete(v2, fi56)
                                    }
                                    _ = fi56
                                    _ = v57
                                }
                                _ = v54
                                _ = v55
                            } else {
                                fmt.Println(v4, v23, v4)
                                if false {
                                    delete(v1, fi52)
                                    v59 := GenS8{F0: 255}
                                    v23 = v23
                                    _ = v59
                                } else {
                                    switch (fi52 & fi52) {
                                    case 625:
                                        var v60 uint32 = 617
                                        _ = v60
                                    case 0:
                                        fmt.Println(fi52, v51, v53)
                                    case 402:
                                        fmt.Println(v51)
                                    }
                                    if true {
                                        var v61 uint64 = 4294967295
                                        fmt.Println(v4)
                                        _ = v61
                                    }
                                    delete(v3, (v51 + (v51 + v51)))
                                }
                                if (true && true) {
                                    for fi62 := 0; fi62 < 6; fi62++ {
                                        if (v50 >= (1.5 / 1.0)) {
                                            var v63 string = "hello"
                                            v64 := GenS6{F0: v23, F1: -1, F2: 42}
                                            fmt.Println(v51)
                                            _ = v63
                                            _ = v64
                                        }
                                        if ((2147483648 == v4) && (!false)) {
                                            v50 -= (0.0 * ((((v50 / float64(3.14159)) / (((v50 / v50) + float64(3.14159)) * ((v50 * (((v50 * v50) * v50) - (v50 + (1.5 / (v50 + float64(3.14159)))))) - v50))) * (float64(1e10) / v50)) - -2.25))
                                            fmt.Println(v4, v51)
                                            var v65 int = (fi62 & 614)
                                            _ = v65
                                        }
                                        _ = fi62
                                    }
                                    if true {
                                        fmt.Println(v53, v53)
                                        if ((2 == 1) && (!false)) {
                                            v66 := []int{1, fi52}
                                            _ = v66
                                        } else {
                                            if false {
                                                v53 = (v4 / (((v4 | v53)) | 1))
                                            }
                                            v1[fi52] = ("\n\t" + ("\n\t" + "a"))
                                            var v67 uint64 = 4294967295
                                            _ = v67
                                        }
                                    } else {
                                        v53 = 2
                                        for fi68 := 0; fi68 < 19; fi68++ {
                                            switch fi68 {
                                            case 9223372036854775807:
                                                switch "hello" {
                                                case "hello":
                                                    v23 -= (^(v23))
                                                default:
                                                    v69 := GenS9{F0: (fi52 | (fi68 ^ fi52))}
                                                    _ = v69
                                                }
                                            case 127:
                                                _ = len(v1)
                                            default:
                                                var v70 int32 = -2147483648
                                                _ = v70
                                            }
                                            switch (((v51 + v51) + v51) + "") {
                                            case "":
                                                var v71 bool = (!(0.1 > ((1e10 * 0.1) * 0.1)))
                                                _ = v71
                                            default:
                                                var v72 uint64 = 0
                                                _ = v72
                                            }
                                            _ = fi68
                                        }
                                    }
                                    var v73 int64 = -1
                                    _ = v73
                                } else {
                                    if (v53 == 552) {
                                        if (1 == 255) {
                                            fmt.Println(v53, v4, fi52)
                                            var v74 int32 = 837
                                            for fi75 := 0; fi75 < 16; fi75++ {
                                                v23 = v23
                                                var v76 int8 = 42
                                                var v77 int = (fi75 >> ((669) & 63))
                                                _ = fi75
                                                _ = v76
                                                _ = v77
                                            }
                                            _ = v74
                                        }
                                    } else {
                                        switch 128 {
                                        case 4294967295:
                                            var v78 float64 = v50
                                            _ = v78
                                        case -1:
                                            v79 := GenS9{F0: -613}
                                            _ = v79
                                        case 42:
                                            _ = len(v2)
                                        default:
                                            var v80 uint8 = 128
                                            _ = v80
                                        }
                                    }
                                    v81 := GenS10{F0: fi52, F1: 2, F2: 1}
                                    v82 := v2[42]
                                    _ = v81
                                    _ = v82
                                }
                            }
                        } else {
                            fmt.Println(v4)
                            switch v51 {
                            case "a":
                                v83 := GenS11{F0: 1, F1: 597}
                                _ = v83
                            default:
                                switch fi52 {
                                case -1:
                                    v53 = v53
                                case 9223372036854775807:
                                    var v84 int8 = 101
                                    _ = v84
                                default:
                                    var v85 uint = 739
                                    _ = v85
                                }
                            }
                            v86 := GenS12{F0: 42}
                            _ = v86
                        }
                        v50 = float64((-2.25 - (3.14159 * 1e10)))
                        for fi87 := 0; fi87 < 12; fi87++ {
                            switch v51 {
                            case "":
                                v23 += (v23 << ((65535) & 63))
                            case "hello":
                                if false {
                                    if (785 < 38) {
                                        if false {
                                            if (false || false) {
                                                delete(v2, (-(fi87)))
                                                v88 := GenS9{F0: 32768}
                                                _ = v88
                                            } else {
                                                v89 := v3[(("hello" + v51) + ("hello" + v51))]
                                                v50 = v50
                                                v90 := GenS8{F0: 42}
                                                v51 = ("foo bar" + v89)
                                                _ = v89
                                                _ = v90
                                            }
                                            if false {
                                                var v91 uint16 = 255
                                                _ = v91
                                            }
                                        }
                                        v53 = (v4 | v53)
                                    } else {
                                        v92 := []int{-9223372036854775808, fi52, fi52, (^(fi87))}
                                        _ = v92
                                    }
                                    v50 = (1.5 / ((0.1 * float64(3.14159)) * v50))
                                    v93 := GenS13{F0: fi52, F1: (18446744073709551615 == 648), F2: 1}
                                    delete(v3, v51)
                                    switch fi52 {
                                    case 348:
                                        v94 := GenS14{F0: (false || (864 < 838)), F1: v51}
                                        _ = v94
                                    case -256:
                                        v4 += (v53 >> ((0) & 63))
                                    }
                                    _ = v93
                                }
                            }
                            if false {
                                switch (fi52 & fi87) {
                                case -555:
                                    v4 = (v53 >> ((18446744073709551615) & 63))
                                case -1:
                                    delete(v1, (-(fi87)))
                                default:
                                    fmt.Println(v23)
                                }
                            } else {
                                v95 := GenS6{F0: (v23 ^ (v23 / (((v23 << ((784) & 63))) | 1))), F1: -9223372036854775808, F2: 0}
                                _ = v95
                            }
                            if (-9223372036854775808 == 42) {
                                if false {
                                    v96 := GenS15{F0: 32767}
                                    _ = v96
                                } else {
                                    v50 += (((v50 / v50) * v50) * ((v50 - v50) / 1.5))
                                    var v97 uint = 1
                                    if true {
                                        if false {
                                            v98 := GenS16{F0: fi87}
                                            _ = v98
                                        }
                                        if (263 > 0) {
                                            switch v51 {
                                            case "a":
                                                if (!(!true)) {
                                                    if true {
                                                        fmt.Println(v23)
                                                        switch v51 {
                                                        case "\n\t":
                                                            v99 := GenS4{F0: (v51 + v51), F1: true}
                                                            _ = v99
                                                        case "hello":
                                                            v100 := GenS17{F0: fi52, F1: true, F2: v4}
                                                            _ = v100
                                                        case "a":
                                                            v1[fi87] = v51
                                                        }
                                                        if true {
                                                            fmt.Println(v53, fi87)
                                                            var v101 float32 = (1e10 + 0.0)
                                                            v101 = v101
                                                            _ = v101
                                                        } else {
                                                            switch (v51 + v51) {
                                                            case "":
                                                                switch fi52 {
                                                                case 178:
                                                                    v102 := GenS18{F0: 18446744073709551615}
                                                                    _ = v102
                                                                }
                                                            case "hello":
                                                                v53 = v53
                                                            case "foo bar":
                                                                fmt.Println(fi52, v97, fi87)
                                                            }
                                                            fmt.Println(v97, v50, v97)
                                                            fmt.Println(v53, v97)
                                                        }
                                                    }
                                                }
                                            default:
                                                switch (("foo bar" + "foo bar") + "foo bar") {
                                                case "hello":
                                                    var v103 uint8 = 254
                                                    v3[((v51 + v51) + v51)] = v51
                                                    _ = v103
                                                case "foo bar":
                                                    v104 := GenS12{F0: 10}
                                                    _ = v104
                                                case "\n\t":
                                                    v105 := GenS19{F0: v97, F1: (fi52 % (((-(fi52))) | 1))}
                                                    _ = v105
                                                default:
                                                    v4 = 2
                                                }
                                            }
                                            switch v51 {
                                            case "":
                                                v106 := []int{32768, fi52}
                                                _ = v106
                                            case "a":
                                                if true {
                                                    v107 := GenS18{F0: 987}
                                                    v51 = v51
                                                    _ = v107
                                                } else {
                                                    var v108 int32 = 71
                                                    v109 := GenS16{F0: fi52}
                                                    _ = v108
                                                    _ = v109
                                                }
                                            case "hello":
                                                v110 := GenS19{F0: (v97 & v97), F1: fi52}
                                                _ = v110
                                            }
                                        }
                                    }
                                    _ = v97
                                }
                                var v111 float32 = 1e10
                                _ = v111
                            } else {
                                var v112 uint8 = 127
                                v113 := GenS5{F0: ("\n\t" + v51), F1: (-(fi52))}
                                v114 := []int{92}
                                var v115 int32 = int32(v53)
                                var v116 int = (fi87 << ((4294967295) & 63))
                                _ = v112
                                _ = v113
                                _ = v114
                                _ = v115
                                _ = v116
                            }
                            _ = fi87
                        }
                    } else {
                        for fi117 := 0; fi117 < 12; fi117++ {
                            if false {
                                v51 = v51
                                var v118 uint = 3
                                var v119 int8 = 59
                                _ = v118
                                _ = v119
                            } else {
                                if (((!false) && false) && true) {
                                    v50 = ((1.5 / v50) * 1.5)
                                } else {
                                    v120 := GenS20{F0: false, F1: (((^(v53)) <= v53) && (true || (200 != 210)))}
                                    switch v51 {
                                    case "foo bar":
                                        if false {
                                            v121 := GenS21{F0: 520, F1: int64(fi117), F2: 81}
                                            _ = v121
                                        }
                                    case "\n\t":
                                        fmt.Println(fi117)
                                    }
                                    _ = v120
                                }
                                if true {
                                    v23 = v23
                                    v122 := v1[(fi117 >> ((42) & 63))]
                                    v123 := GenS22{F0: (fi117 * (-(fi117))), F1: 249, F2: 42}
                                    var v124 float32 = 0.0
                                    switch v51 {
                                    case "\n\t":
                                        var v125 bool = false
                                        _ = v125
                                    case "":
                                        var v126 int8 = 42
                                        _ = v126
                                    case "foo bar":
                                        v127 := GenS15{F0: 65535}
                                        _ = v127
                                    default:
                                        v50 = v50
                                    }
                                    _ = v122
                                    _ = v123
                                    _ = v124
                                } else {
                                    fmt.Println(fi117, v23, v23)
                                }
                                if true {
                                    v53 -= v4
                                } else {
                                    var v128 uint = 361
                                    var v129 string = (v51 + (v51 + "a"))
                                    v23 = (v23 * 42)
                                    _ = v128
                                    _ = v129
                                }
                                fmt.Println(v51, v4)
                                if true {
                                    var v130 uint64 = 129
                                    _ = v130
                                } else {
                                    v131 := GenS6{F0: (v23 >> ((1) & 63)), F1: 1, F2: 0}
                                    v23 = (^(v23))
                                    v132 := GenS23{F0: 255}
                                    _ = v131
                                    _ = v132
                                }
                            }
                            v51 = (("foo bar" + (v51 + v51)) + v51)
                            _ = fi117
                        }
                        v133 := v2[(fi52 - -194)]
                        _ = v133
                    }
                    _ = fi52
                    _ = v53
                }
                _ = v50
                _ = v51
            } else {
                v134 := GenS18{F0: 2}
                v135 := []int{0, -753, 1, 2147483647}
                for fi136 := 0; fi136 < 6; fi136++ {
                    for fi137 := 0; fi137 < 13; fi137++ {
                        var v138 int = fi136
                        var v139 uint16 = 65535
                        v23 = (v23 % ((v23) | 1))
                        _ = fi137
                        _ = v138
                        _ = v139
                    }
                    _ = fi136
                }
                _ = v134
                _ = v135
            }
        }
    case -9223372036854775808:
        var v140 string = ("foo bar" + "hello")
        _ = v140
    }
    v23 += -92
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v41
}
