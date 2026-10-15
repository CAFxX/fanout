package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 uint; F1 bool; F2 string }
type GenS5 struct { F0 int64; F1 uint; F2 int16 }
type GenS6 struct { F0 int64; F1 uint16 }
type GenS7 struct { F0 int32; F1 bool }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int; F1 bool; F2 int }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint8 }

func main() {
    var v1 int = 4294967295
    v2 := make(map[string]uint32)
    v3 := make(map[int]bool)
    v4 := make(map[string]uint32)
    v5 := make(map[int]bool)
    var v6 int8 = 0
    var v7 string = "\n\t"
    v7 = "a"
    v7 = "foo bar"
    v1 = v1
    v7 = "foo bar"
    v4["hello"] = 120
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 uint8 = 2
        var v10 int16 = 0
        if false {
            fmt.Println(v6, v9, v7)
            for fi11 := 0; fi11 < 17; fi11++ {
                v1 -= (fi8 - v1)
                _ = fi11
            }
        } else {
            v12 := GenS1{F0: uint16(v9), F1: -1, F2: (v7 + "foo bar")}
            v13 := []int{(v1 | fi8)}
            v14 := GenS2{F0: 661}
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        v16 := GenS1{F0: 995, F1: 904, F2: ("" + (v7 + v7))}
        if (42 > 0) {
            if false {
                var v17 int64 = -9223372036854775808
                if (!false) {
                    for fi18 := 0; fi18 < 19; fi18++ {
                        v6 -= (^(v6))
                        var v19 uint64 = 42
                        _ = fi18
                        _ = v19
                    }
                    v7 = "foo bar"
                } else {
                    var v20 int = (v1 & -2147483648)
                    var v21 uint = 150
                    _ = v20
                    _ = v21
                }
                _ = v17
            } else {
                var v22 int = 127
                _ = v22
            }
            fmt.Println(fi15, v1, fi15)
            switch 529 {
            case -479:
                for fi23 := 0; fi23 < 10; fi23++ {
                    var v24 float32 = float32(3.14159)
                    if (v1 <= (v1 % ((v1) | 1))) {
                        v25 := []int{-9223372036854775808, (fi23 - 32767), -135}
                        fmt.Println(fi23, v1, v7)
                        v7 = "hello"
                        _ = v25
                    } else {
                        fmt.Println(v24)
                        var v26 string = v7
                        var v27 bool = true
                        v7 = "foo bar"
                        fmt.Println(v1)
                        if false {
                            fmt.Println(v26)
                        } else {
                            v28 := GenS2{F0: -537}
                            _ = v28
                        }
                        _ = v26
                        _ = v27
                    }
                    fmt.Println(fi23)
                    _ = fi23
                    _ = v24
                }
            case 32768:
                for fi29 := 0; fi29 < 13; fi29++ {
                    var v30 bool = false
                    switch "\n\t" {
                    case "a":
                        switch (-(v1)) {
                        case -935:
                            v6 = (-(v6))
                        }
                    case "foo bar":
                        if (v1 > fi15) {
                            switch (v7 + v7) {
                            case "\n\t":
                                var v31 uint16 = 2
                                _ = v31
                            case "":
                                switch "hello" {
                                case "\n\t":
                                    v32 := GenS1{F0: 65535, F1: 0, F2: v7}
                                    if (!(299 >= 42)) {
                                        v1 += (-(fi29))
                                        v7 = "foo bar"
                                    }
                                    if (((v30 && ((42 > 1000) && true)) && (32767 <= 0)) || ((v30 || v30) && v30)) {
                                        var v33 int32 = 42
                                        v1 = fi29
                                        _ = v33
                                    }
                                    _ = v32
                                default:
                                    var v34 string = v7
                                    _ = v34
                                }
                            default:
                                v35 := GenS2{F0: -1}
                                _ = v35
                            }
                        } else {
                            v6 = v6
                            fmt.Println(v1, v1, v7)
                        }
                    case "":
                        switch "\n\t" {
                        case "\n\t":
                            v1 = 9223372036854775807
                        default:
                            v36 := GenS2{F0: -9223372036854775808}
                            _ = v36
                        }
                    }
                    switch (v1 - (fi15 << ((65535) & 63))) {
                    case -371:
                        v30 = v30
                        var v37 int64 = -9223372036854775808
                        v38 := []int{2147483647, v1}
                        if true {
                            if ((("" + v7) + v7) <= (v7 + (v7 + v7))) {
                                switch v7 {
                                case "\n\t":
                                    var v39 int16 = int16(v37)
                                    _ = v39
                                case "":
                                    v6 = v6
                                case "hello":
                                    fmt.Println(v37, v6)
                                default:
                                    var v40 int8 = v6
                                    var v41 int16 = 201
                                    switch ("a" + v7) {
                                    case "\n\t":
                                        var v42 int64 = (v37 ^ v37)
                                        _ = v42
                                    case "foo bar":
                                        var v43 uint16 = 42
                                        _ = v43
                                    case "a":
                                        var v44 string = ("hello" + v7)
                                        _ = v44
                                    default:
                                        if (65535 >= 336) {
                                            var v45 int8 = -32
                                            var v46 int32 = 437
                                            _ = v45
                                            _ = v46
                                        }
                                    }
                                    v40 = 63
                                    _ = v40
                                    _ = v41
                                }
                            } else {
                                if v30 {
                                    switch ((v7 + "foo bar") + ((v7 + v7) + ((((v7 + (("\n\t" + (v7 + v7)) + (v7 + "foo bar"))) + v7) + (v7 + ("\n\t" + v7))) + "\n\t"))) {
                                    case "\n\t":
                                        v7 = "a"
                                    default:
                                        v1 += 1
                                    }
                                }
                                if (0.1 < 0.1) {
                                    switch fi29 {
                                    case 128:
                                        v47 := GenS3{F0: 65535, F1: 410}
                                        _ = v47
                                    case -2147483648:
                                        v1 = fi15
                                    case 2147483647:
                                        var v48 bool = false
                                        _ = v48
                                    default:
                                        var v49 int = (fi15 ^ 32767)
                                        _ = v49
                                    }
                                    if true {
                                        var v50 uint64 = 392
                                        _ = v50
                                    }
                                    if false {
                                        switch (fi29 & (fi29 * (fi15 & (fi29 * fi29)))) {
                                        case -164:
                                            v51 := GenS4{F0: 0, F1: v30, F2: v7}
                                            _ = v51
                                        default:
                                            v1 = 32768
                                        }
                                    } else {
                                        if (0.0 > (0.1 / 1.5)) {
                                            var v52 uint32 = 0
                                            _ = v52
                                        }
                                        switch "a" {
                                        case "hello":
                                            _ = len(v3)
                                        case "a":
                                            if v30 {
                                                if (!("\n\t" >= (v7 + v7))) {
                                                    v6 -= 127
                                                    _ = len(v3)
                                                } else {
                                                    if false {
                                                        v30 = true
                                                    } else {
                                                        v53 := GenS4{F0: 228, F1: true, F2: v7}
                                                        _ = v53
                                                    }
                                                }
                                                fmt.Println(fi15)
                                                var v54 uint8 = 2
                                                _ = v54
                                            } else {
                                                v7 = ""
                                                if (2147483647 == 2147483648) {
                                                    var v55 uint = 573
                                                    var v56 uint64 = 4294967295
                                                    v57 := GenS1{F0: 32767, F1: v37, F2: v7}
                                                    _ = v55
                                                    _ = v56
                                                    _ = v57
                                                }
                                                v6 = v6
                                            }
                                            fmt.Println(v6, fi29)
                                            if ((-(v1)) == fi29) {
                                                if (v30 && v30) {
                                                    var v58 int32 = 642
                                                    v59 := GenS2{F0: v37}
                                                    v60 := GenS5{F0: 42, F1: 1, F2: -1}
                                                    v6 += 1
                                                    _ = v58
                                                    _ = v59
                                                    _ = v60
                                                } else {
                                                    v61 := GenS4{F0: 42, F1: v30, F2: "foo bar"}
                                                    if (!(v30 && (42 >= 4294967295))) {
                                                        v37 = v37
                                                        if len(v38) > 0 {
                                                            v38[0] = 222
                                                        }
                                                    }
                                                    _ = v61
                                                }
                                            }
                                        }
                                        if len(v38) > 0 {
                                            v38[0] = fi15
                                        }
                                    }
                                    v62 := GenS6{F0: v37, F1: 0}
                                    v6 = (v6 >> ((801) & 63))
                                    _ = v62
                                } else {
                                    _ = len(v3)
                                    var v63 string = (("" + "a") + "")
                                    _ = v63
                                }
                                fmt.Println(v37)
                            }
                            if len(v38) > 0 {
                                v38[0] = v1
                            }
                            if (v30 && v30) {
                                _ = len(v38)
                                var v64 string = v7
                                v7 = v7
                                _ = v64
                            } else {
                                v1 = fi15
                                var v65 float32 = 0.0
                                v37 -= (v37 / (((v37 << ((758) & 63))) | 1))
                                v6 -= v6
                                if false {
                                    v66 := GenS5{F0: (v37 | v37), F1: 1, F2: -610}
                                    v37 += (-(v37))
                                    var v67 int16 = -671
                                    _ = v66
                                    _ = v67
                                } else {
                                    v30 = (v1 <= 949)
                                    v6 = (v6 & v6)
                                    v6 += v6
                                }
                                _ = v65
                            }
                        }
                        v7 = ""
                        _ = v37
                        _ = v38
                    case -919:
                        if ((v7 + ("" + v7)) > ((v7 + "") + v7)) {
                            if v30 {
                                switch v7 {
                                case "hello":
                                    if true {
                                        var v68 float32 = -2.25
                                        v69 := GenS7{F0: 2147483647, F1: v30}
                                        _ = v68
                                        _ = v69
                                    } else {
                                        var v70 float64 = 0.0
                                        var v71 string = v7
                                        _ = v70
                                        _ = v71
                                    }
                                }
                            }
                        } else {
                            switch fi15 {
                            case 127:
                                fmt.Println(fi15, v6, fi15)
                            case 32767:
                                var v72 int32 = 42
                                _ = v72
                            default:
                                v73 := GenS7{F0: -1, F1: v30}
                                _ = v73
                            }
                            v74 := GenS1{F0: 0, F1: 9223372036854775807, F2: "foo bar"}
                            switch (v7 + ((v7 + ("a" + v7)) + "foo bar")) {
                            case "":
                                if false {
                                    var v75 int8 = 1
                                    v30 = (false || true)
                                    _ = v75
                                } else {
                                    var v76 int = v1
                                    fmt.Println(v7, fi29)
                                    _ = v76
                                }
                            case "foo bar":
                                delete(v4, (((("foo bar" + "") + ("\n\t" + v7)) + v7) + v7))
                            case "\n\t":
                                v77 := GenS6{F0: -9223372036854775808, F1: 0}
                                _ = v77
                            }
                            _ = v74
                        }
                    default:
                        if (!(!false)) {
                            if (32767 <= 14) {
                                switch v7 {
                                case "a":
                                    if ((fi29 / ((v1) | 1)) >= 1) {
                                        var v78 int32 = 0
                                        var v79 uint32 = 4294967295
                                        _ = v78
                                        _ = v79
                                    }
                                    fmt.Println(fi29, fi15, v6)
                                    var v80 float64 = ((1e10 * (1.5 + 0.0)) * 0.1)
                                    if v30 {
                                        v6 = (v6 % (((^(v6))) | 1))
                                        var v81 string = v7
                                        v82 := v2[v81]
                                        _ = v81
                                        _ = v82
                                    }
                                    v7 = ((v7 + "\n\t") + ("\n\t" + v7))
                                    v83 := v2["foo bar"]
                                    _ = v80
                                    _ = v83
                                case "\n\t":
                                    var v84 int16 = -371
                                    _ = v84
                                default:
                                    v30 = false
                                    var v85 int8 = v6
                                    if (((((float64(1e10) - 1e10) * 1e10) / (0.1 / 3.14159)) - 1e10) >= (1.5 * (3.14159 / 1e10))) {
                                        switch fi29 {
                                        case 631:
                                            _ = len(v5)
                                        default:
                                            v7 = "a"
                                        }
                                        v2[("foo bar" + ((v7 + v7) + "foo bar"))] = 213
                                    } else {
                                        if (!(v7 > v7)) {
                                            var v86 uint8 = 1
                                            var v87 uint16 = 32767
                                            _ = v86
                                            _ = v87
                                        }
                                        v6 = (v6 >> ((0) & 63))
                                        if (127 != 42) {
                                            v88 := GenS1{F0: 369, F1: 1, F2: v7}
                                            fmt.Println(v30, fi15, v85)
                                            v89 := GenS8{F0: (float64(-2.25) == (1.5 * 1e10))}
                                            v90 := v3[-914]
                                            _ = v88
                                            _ = v89
                                            _ = v90
                                        }
                                        switch "a" {
                                        case "\n\t":
                                            var v91 int32 = 0
                                            _ = v91
                                        case "a":
                                            v3[fi29] = true
                                        default:
                                            v7 = v7
                                            var v92 uint16 = 42
                                            _ = v92
                                        }
                                        v93 := GenS1{F0: 393, F1: 42, F2: ""}
                                        _ = v93
                                    }
                                    var v94 uint = 65535
                                    _ = v85
                                    _ = v94
                                }
                                fmt.Println(v1)
                                var v95 int8 = (v6 + v6)
                                v30 = (0 != 812)
                                _ = v95
                            } else {
                                v96 := []int{127, fi29, (fi15 & -1)}
                                v1 += 32768
                                v5[-2147483648] = v30
                                if ((((!(v7 < v7)) || (1 < (v6 & (v6 & v6)))) && (!v30)) && false) {
                                    v97 := GenS9{F0: ((0.0 * (((1e10 / 3.14159) - (1e10 + 3.14159)) / (float32(((((0.1 / 1e10) - 1.5) * ((0.0 * 1e10) - 1.5)) * 0.1)) - (3.14159 / 3.14159)))) > (float32((0.0 - (0.0 * 0.0))) / float32(0.1)))}
                                    if len(v96) > 0 {
                                        v96[0] = (-(fi15))
                                    }
                                    _ = v97
                                }
                                v1 = -737
                                fmt.Println(fi29, v6)
                                fmt.Println(v6)
                                _ = v96
                            }
                            v1 = v1
                            if v30 {
                                if v30 {
                                    var v98 int = v1
                                    var v99 int64 = 9223372036854775807
                                    v100 := GenS6{F0: (v99 + v99), F1: 32767}
                                    _ = v100
                                    _ = v98
                                    _ = v99
                                }
                                switch (-(v1)) {
                                case 9223372036854775807:
                                    if v30 {
                                        v1 = fi15
                                        fmt.Println(fi15, v1)
                                    }
                                case 32768:
                                    v7 = ((((v7 + "a") + v7) + ("foo bar" + v7)) + "")
                                case -672:
                                    var v101 int16 = 1
                                    _ = v101
                                default:
                                    v102 := GenS4{F0: 0, F1: (!v30), F2: ""}
                                    _ = v102
                                }
                                var v103 string = "a"
                                _ = v103
                            } else {
                                switch v7 {
                                case "foo bar":
                                    var v104 float32 = 1e10
                                    _ = v104
                                case "a":
                                    fmt.Println(v7, v7)
                                default:
                                    var v105 int64 = -1
                                    _ = v105
                                }
                                switch 4294967295 {
                                case -490:
                                    delete(v2, (v7 + "a"))
                                case -9223372036854775808:
                                    _ = len(v3)
                                }
                            }
                        } else {
                            var v106 uint8 = 127
                            if v30 {
                                switch 2147483647 {
                                case -1:
                                    var v107 bool = (2 == 18446744073709551615)
                                    _ = v107
                                default:
                                    var v108 int32 = 2147483647
                                    _ = v108
                                }
                                v109 := GenS10{F0: fi15, F1: v30, F2: (^(fi29))}
                                v1 += v1
                                _ = v109
                            }
                            switch 128 {
                            case -252:
                                v106 += (v106 ^ v106)
                            case 382:
                                v7 = "hello"
                            case 0:
                                v1 -= v1
                                v110 := []int{fi15, (fi15 + -9223372036854775808), -377}
                                _ = v110
                            default:
                                switch (fi29 / (((fi29 | (v1 & fi29))) | 1)) {
                                case 0:
                                    v7 = ("" + v7)
                                }
                            }
                            var v111 uint16 = 255
                            if (!((0.1 - 3.14159) <= ((((1.5 - (0.0 - 3.14159)) - ((0.0 + -2.25) / -2.25)) * (3.14159 / 1.5)) / 1.0))) {
                                switch (fi29 % ((9223372036854775807) | 1)) {
                                case 4294967295:
                                    fmt.Println(v30, fi29)
                                }
                            }
                            if false {
                                var v112 int64 = 9223372036854775807
                                if (!((18446744073709551615 >= 370) || ((42 < 39) && v30))) {
                                    var v113 int64 = 417
                                    switch "hello" {
                                    case "a":
                                        switch (v7 + (v7 + "\n\t")) {
                                        case "foo bar":
                                            v114 := GenS3{F0: (v111 % (((v111 / ((v111) | 1))) | 1)), F1: 4294967295}
                                            _ = v114
                                        }
                                        v115 := []int{-2147483648, v1, -509, (-(fi29))}
                                        _ = v115
                                    case "hello":
                                        var v116 uint32 = 412
                                        _ = v116
                                    default:
                                        var v117 int = 713
                                        _ = v117
                                    }
                                    _ = v113
                                }
                                _ = v112
                            } else {
                                fmt.Println(v6, v7, fi29)
                                switch "\n\t" {
                                case "\n\t":
                                    var v118 bool = ((v30 || v30) && true)
                                    _ = v118
                                case "":
                                    v106 += v106
                                case "hello":
                                    v119 := []int{}
                                    _ = v119
                                default:
                                    fmt.Println(v1)
                                }
                                v1 -= int(v6)
                            }
                            _ = v106
                            _ = v111
                        }
                    }
                    _ = fi29
                    _ = v30
                }
            case -1:
                var v120 float32 = 0.0
                _ = v120
            }
        } else {
            fmt.Println(v7)
            v121 := []int{v1, 1}
            _ = v121
        }
        _ = fi15
        _ = v16
    }
    if ((-1 >= -32768) || true) {
        v122 := GenS11{F0: v1}
        _ = len(v4)
        _ = v122
    }
    var v123 int8 = -1
    if true {
        v124 := GenS12{F0: 128}
        _ = v124
    }
    if false {
        _ = len(v3)
    } else {
        switch v7 {
        case "hello":
            v6 -= (v6 << ((2) & 63))
        case "":
            var v125 uint8 = 65
            _ = v125
        }
        var v126 uint8 = 255
        _ = v126
    }
    var v127 int8 = -85
    for fi128 := 0; fi128 < 10; fi128++ {
        v123 = 127
        var v129 float32 = (-2.25 / -2.25)
        _ = fi128
        _ = v129
    }
    v130 := []int{(v1 | v1), (v1 ^ (v1 >> ((60) & 63))), 32768}
    for fi131 := 0; fi131 < 13; fi131++ {
        fmt.Println(v6)
        v132 := GenS7{F0: -559, F1: false}
        _ = fi131
        _ = v132
    }
    v127 += (v127 - v127)
    fmt.Println("v1:", v1)
    fmt.Println("v123:", v123)
    fmt.Println("v127:", v127)
    fmt.Println("slicelen v130:", len(v130), v130)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v123
    _ = v127
    _ = v130
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
