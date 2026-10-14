package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64; F1 int8; F2 int16 }
type GenS6 struct { F0 uint; F1 uint64; F2 int32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool; F1 uint16 }
type GenS9 struct { F0 int32; F1 uint32 }
type GenS10 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 int = 95
    var v2 uint16 = 32767
    var v3 uint32 = 858
    var v4 string = "foo bar"
    var v5 int = -538
    if true {
        v6 := []int{(v1 & 1)}
        v2 = 65535
        _ = v6
    }
    var v7 string = "foo bar"
    var v8 int32 = -1
    var v9 int16 = -32768
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 int8 = 100
        _ = fi10
        _ = v11
    }
    if (true && true) {
        fmt.Println(v1, v3)
    } else {
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 int8 = -128
            var v14 uint = 326
            for fi15 := 0; fi15 < 12; fi15++ {
                v16 := []int{(v1 % ((9223372036854775807) | 1))}
                v14 = v14
                v17 := GenS1{F0: 128}
                v1 -= (-(fi15))
                if false {
                    v8 += v8
                } else {
                    v16 = append(v16, (fi15 / ((9223372036854775807) | 1)))
                    switch 9223372036854775807 {
                    case 4294967295:
                        v5 = (fi15 - (^(fi15)))
                        var v18 uint32 = (v3 + (v3 - 65535))
                        if len(v16) > 0 {
                            v16[0] = (^(v1))
                        }
                        _ = v18
                    case 1:
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                    }
                    v4 = v7
                    v19 := GenS2{F0: 136, F1: v13}
                    _ = v19
                }
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        var v20 uint32 = (v3 % ((936) | 1))
        _ = v20
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        switch v1 {
        case 32767:
            v22 := GenS3{F0: 0, F1: 1}
            _ = v22
        case -1:
            var v23 bool = (4294967295 != 0)
            _ = v23
        default:
            if (1 == 1) {
                var v24 int = v5
                if (0.0 == (((1e10 * 0.1) / 1.5) * 0.1)) {
                    for fi25 := 0; fi25 < 8; fi25++ {
                        switch fi21 {
                        case -1:
                            var v26 string = v7
                            _ = v26
                        case 642:
                            var v27 float32 = float32(1e10)
                            _ = v27
                        case 156:
                            v4 = v7
                        default:
                            switch (v24 >> ((0) & 63)) {
                            case -786:
                                fmt.Println(v7, v9, v9)
                            case 933:
                                v28 := GenS2{F0: 127, F1: 42}
                                _ = v28
                            default:
                                v8 = v8
                                if (4294967295 != 2) {
                                    var v29 bool = (!true)
                                    _ = v29
                                } else {
                                    if (1e10 <= -2.25) {
                                        var v30 int = fi21
                                        v31 := []int{4294967295, 2147483647, (v24 >> ((2) & 63)), (v24 | 1)}
                                        switch 0 {
                                        case 252:
                                            v7 = v4
                                        }
                                        _ = v30
                                        _ = v31
                                    }
                                    fmt.Println(fi25)
                                    v32 := GenS2{F0: 221, F1: 93}
                                    var v33 string = ""
                                    v34 := []int{}
                                    var v35 float32 = 0.1
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                                v1 = 293
                            }
                        }
                        v36 := []int{v1, (v1 + -615)}
                        _ = fi25
                        _ = v36
                    }
                    for fi37 := 0; fi37 < 20; fi37++ {
                        switch (v5 << ((269) & 63)) {
                        case 4294967295:
                            v38 := GenS2{F0: 128, F1: 5}
                            _ = v38
                        }
                        v5 += 1
                        switch (((v7 + "\n\t") + v7) + v4) {
                        case "hello":
                            var v39 bool = (0 != 815)
                            _ = v39
                        case "\n\t":
                            switch v7 {
                            case "foo bar":
                                switch "\n\t" {
                                case "a":
                                    fmt.Println(v3, v3)
                                    switch (-(fi21)) {
                                    case 0:
                                        var v40 int8 = 81
                                        _ = v40
                                    default:
                                        var v41 int64 = 42
                                        _ = v41
                                    }
                                case "\n\t":
                                    v42 := GenS3{F0: 42, F1: 42}
                                    _ = v42
                                case "foo bar":
                                    switch "hello" {
                                    case "\n\t":
                                        v43 := GenS2{F0: 1, F1: 127}
                                        _ = v43
                                    default:
                                        v44 := GenS1{F0: 53}
                                        v1 = (^(fi37))
                                        _ = v44
                                    }
                                }
                            case "hello":
                                var v45 uint64 = 0
                                _ = v45
                            }
                        case "foo bar":
                            fmt.Println(v9, v5, v24)
                        }
                        _ = fi37
                    }
                    var v46 uint32 = v3
                    _ = v46
                }
                var v47 float64 = float64(0.0)
                _ = v24
                _ = v47
            }
        }
        _ = fi21
    }
    for fi48 := 0; fi48 < 12; fi48++ {
        var v49 int16 = (v9 | (^(v9)))
        _ = fi48
        _ = v49
    }
    if true {
        v50 := GenS4{F0: 2}
        v51 := []int{v1, 394, v5}
        v9 -= 1000
        _ = v50
        _ = v51
    } else {
        if true {
            for fi52 := 0; fi52 < 18; fi52++ {
                var v53 uint16 = (v2 >> ((65535) & 63))
                for fi54 := 0; fi54 < 13; fi54++ {
                    v1 += -1
                    if true {
                        v55 := GenS5{F0: 1, F1: 0, F2: v9}
                        _ = v55
                    } else {
                        if (9223372036854775807 != 42) {
                            v56 := GenS6{F0: 126, F1: uint64(v3), F2: (v8 / ((v8) | 1))}
                            v57 := GenS6{F0: 65535, F1: 18446744073709551615, F2: (v8 & v8)}
                            _ = v56
                            _ = v57
                        }
                        v58 := GenS3{F0: 0, F1: 42}
                        switch "\n\t" {
                        case "":
                            var v59 int16 = v9
                            v60 := []int{fi52, 356, (^(fi52))}
                            _ = v59
                            _ = v60
                        case "hello":
                            fmt.Println(fi52, v2)
                        case "foo bar":
                            v61 := []int{1, (-(fi52)), (v1 * (^(v1))), -1}
                            _ = v61
                        }
                        switch v5 {
                        case 127:
                            var v62 int16 = v9
                            _ = v62
                        case -2147483648:
                            v9 += v9
                        case 1:
                            var v63 uint32 = 42
                            _ = v63
                        default:
                            if true {
                                v64 := GenS7{F0: "foo bar"}
                                _ = v64
                            }
                        }
                        var v65 int64 = 1
                        _ = v58
                        _ = v65
                    }
                    v66 := GenS2{F0: 1, F1: -128}
                    _ = fi54
                    _ = v66
                }
                v3 = v3
                _ = fi52
                _ = v53
            }
            var v67 float64 = -2.25
            _ = v67
        } else {
            for fi68 := 0; fi68 < 18; fi68++ {
                var v69 int16 = (v9 * 0)
                if false {
                    var v70 float32 = float32((((float64((float32(3.14159) * 0.0)) / 0.1) + 1.5) - 0.1))
                    if ((0 == 963) || (!false)) {
                        if (-2147483648 <= (v8 * (v8 * v8))) {
                            var v71 float32 = 0.1
                            v72 := []int{623}
                            var v73 int64 = 504
                            _ = v71
                            _ = v72
                            _ = v73
                        } else {
                            if ((1.5 + v70) == v70) {
                                v74 := []int{(fi68 | (fi68 & (fi68 % (((fi68 / ((v5) | 1))) | 1))))}
                                if (-99 < 0) {
                                    for fi75 := 0; fi75 < 9; fi75++ {
                                        v2 = v2
                                        fmt.Println(fi68)
                                        fmt.Println(v69, v2)
                                        _ = fi75
                                    }
                                } else {
                                    var v76 int16 = v9
                                    v3 = v3
                                    v69 -= v69
                                    _ = v76
                                }
                                _ = v74
                            } else {
                                if (42 != int32(v3)) {
                                    v77 := GenS7{F0: "a"}
                                    _ = v77
                                } else {
                                    fmt.Println(v7, v3)
                                    var v78 float32 = v70
                                    _ = v78
                                }
                                v79 := []int{v1, (fi68 | int(v8)), (fi68 - -2147483648)}
                                _ = v79
                            }
                            if true {
                                v5 -= 127
                            } else {
                                v69 = -32768
                                var v80 uint8 = 128
                                v8 -= v8
                                _ = v80
                            }
                            switch v5 {
                            case 4294967295:
                                for fi81 := 0; fi81 < 4; fi81++ {
                                    fmt.Println(v2, fi81)
                                    var v82 uint = 1
                                    _ = fi81
                                    _ = v82
                                }
                            case -1:
                                var v83 int32 = 1
                                _ = v83
                            }
                        }
                        v9 += 1
                    } else {
                        if (!((322 > v8) || true)) {
                            if false {
                                if (9223372036854775807 < 9223372036854775807) {
                                    v1 = v5
                                    fmt.Println(v9, v1, v69)
                                    if ("a" > "") {
                                        v3 += v3
                                        for fi84 := 0; fi84 < 3; fi84++ {
                                            v4 = (v4 + "a")
                                            if (42 <= 18446744073709551615) {
                                                v85 := []int{(fi68 / ((fi68) | 1)), 0, (fi84 % ((fi68) | 1)), 4294967295}
                                                _ = v85
                                            }
                                            _ = fi84
                                        }
                                        v86 := []int{}
                                        _ = v86
                                    } else {
                                        var v87 string = "hello"
                                        var v88 uint32 = v3
                                        _ = v87
                                        _ = v88
                                    }
                                } else {
                                    switch (v4 + ("\n\t" + (v7 + "hello"))) {
                                    case "a":
                                        v1 = 1
                                    }
                                }
                            } else {
                                switch int(v2) {
                                case 127:
                                    if (!(0 == -9223372036854775808)) {
                                        switch v7 {
                                        case "a":
                                            v89 := GenS3{F0: -1, F1: 42}
                                            _ = v89
                                        case "foo bar":
                                            v90 := GenS1{F0: 127}
                                            _ = v90
                                        }
                                        v91 := GenS3{F0: 418, F1: uint64(v3)}
                                        _ = v91
                                    }
                                case 699:
                                    if (v2 <= v2) {
                                        v69 += v69
                                        fmt.Println(v8, v7)
                                    }
                                case 32767:
                                    var v92 int8 = 67
                                    _ = v92
                                }
                            }
                            v93 := []int{v1}
                            _ = v93
                        }
                    }
                    if (0 > 0) {
                        var v94 int16 = int16(v3)
                        for fi95 := 0; fi95 < 17; fi95++ {
                            if (false || (!((!false) && false))) {
                                v96 := GenS2{F0: 255, F1: 100}
                                _ = v96
                            }
                            _ = fi95
                        }
                        _ = v94
                    }
                    _ = v70
                } else {
                    v4 = v7
                }
                _ = fi68
                _ = v69
            }
            v97 := GenS8{F0: (true && (((v9 != 727) && true) && (false || ((v5 & 0) <= int(v3))))), F1: (v2 | v2)}
            _ = v97
        }
        v7 = v7
        var v98 uint16 = 255
        _ = v98
    }
    v99 := GenS5{F0: 602, F1: 0, F2: (v9 | v9)}
    for fi100 := 0; fi100 < 9; fi100++ {
        var v101 int16 = (v9 % ((590) | 1))
        var v102 string = (v4 + (v4 + (v7 + v7)))
        _ = fi100
        _ = v101
        _ = v102
    }
    if ((^(v2)) >= v2) {
        fmt.Println(v2, v7)
        for fi103 := 0; fi103 < 8; fi103++ {
            if false {
                fmt.Println(v4, v5, v3)
                v104 := []int{-1, v5, 32767, (fi103 / ((v1) | 1))}
                var v105 float64 = -2.25
                _ = v104
                _ = v105
            } else {
                if (false || true) {
                    v1 = v5
                    v106 := GenS4{F0: 65535}
                    var v107 int8 = -128
                    _ = v106
                    _ = v107
                } else {
                    v108 := GenS2{F0: 0, F1: 60}
                    _ = v108
                }
                for fi109 := 0; fi109 < 16; fi109++ {
                    v2 -= 979
                    v110 := []int{(v5 >> ((255) & 63)), int(v2), v1}
                    v7 = ""
                    _ = fi109
                    _ = v110
                }
                var v111 int = (fi103 % (((fi103 / (((v1 << ((1) & 63))) | 1))) | 1))
                _ = v111
            }
            if false {
                var v112 bool = false
                _ = v112
            } else {
                switch v4 {
                case "foo bar":
                    for fi113 := 0; fi113 < 18; fi113++ {
                        v114 := GenS1{F0: 1}
                        _ = fi113
                        _ = v114
                    }
                case "":
                    var v115 int64 = -9223372036854775808
                    _ = v115
                case "a":
                    if true {
                        if (42 >= 1) {
                            fmt.Println(v1, v3)
                            v116 := GenS9{F0: v8, F1: v3}
                            for fi117 := 0; fi117 < 7; fi117++ {
                                switch v7 {
                                case "":
                                    v118 := []int{(v5 + (v5 / (((v5 / ((-1) | 1))) | 1)))}
                                    _ = v118
                                case "\n\t":
                                    v119 := GenS8{F0: (21 >= 4294967295), F1: (v2 - 2)}
                                    _ = v119
                                default:
                                    var v120 int16 = -1
                                    _ = v120
                                }
                                v121 := []int{(v5 ^ 114)}
                                if (v5 < fi103) {
                                    var v122 int64 = -1
                                    v123 := GenS5{F0: 548, F1: -22, F2: (v9 << ((479) & 63))}
                                    _ = v122
                                    _ = v123
                                } else {
                                    _ = len(v121)
                                }
                                _ = fi117
                                _ = v121
                            }
                            _ = v116
                        }
                        for fi124 := 0; fi124 < 7; fi124++ {
                            v125 := GenS10{F0: (v2 & 517), F1: 2}
                            _ = fi124
                            _ = v125
                        }
                        v126 := GenS5{F0: 257, F1: 42, F2: v9}
                        _ = v126
                    }
                default:
                    var v127 uint = 4294967295
                    _ = v127
                }
            }
            v128 := []int{}
            _ = fi103
            _ = v128
        }
        v129 := []int{-546, v5}
        _ = v129
    }
    for fi130 := 0; fi130 < 18; fi130++ {
        fmt.Println(v8)
        v131 := GenS1{F0: 1}
        _ = fi130
        _ = v131
    }
    switch (v1 % ((v1) | 1)) {
    case 243:
        var v132 uint32 = v3
        _ = v132
    case 128:
        var v133 uint32 = v3
        _ = v133
    case -534:
        v134 := []int{v1}
        _ = v134
    }
    v3 = (^(v3))
    v135 := GenS3{F0: 1, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v135
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
    _ = v99
}
