package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS2 struct { F0 bool; F1 uint32; F2 int }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint16; F1 int64 }
type GenS9 struct { F0 string; F1 uint32; F2 bool }
func genFn0() string {
    for fi6 := 0; fi6 < 10; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := []int{(-(fi6)), 32767, (fi7 ^ (fi6 - 843)), (-(fi6))}
            v9 := GenS1{F0: 223, F1: 0, F2: 9223372036854775807}
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    var v10 string = "\n\t"
    v10 = (v10 + "hello")
    _ = v10
    return ("hello" + v10)
}
func genFn1(p0 uint, p1 int64) uint8 {
    p0 = 1
    p1 = (^(p1))
    switch "foo bar" {
    case "hello":
        p0 = 18446744073709551615
    case "a":
        p1 -= p1
    default:
        var v11 uint = (p0 - p0)
        _ = v11
    }
    var v12 uint16 = 255
    v13 := []int{1, -2147483648, 128}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 255
}

func main() {
    var v1 uint16 = 42
    v2 := make(map[string]int)
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]uint32)
    var v5 int32 = 292
    fmt.Println(v5)
    _ = len(v3)
    for fi14 := 0; fi14 < 18; fi14++ {
        if (!false) {
            _ = len(v2)
        }
        v15 := GenS1{F0: 72, F1: 42, F2: 359}
        _ = fi14
        _ = v15
    }
    v5 = (v5 % (((^(v5))) | 1))
    var v16 uint32 = 865
    for fi17 := 0; fi17 < 14; fi17++ {
        if (true || (false || true)) {
            var v18 int32 = -1
            v19 := []int{-225, (^(fi17)), 597, (^(fi17))}
            var v20 float64 = 3.14159
            _ = v18
            _ = v19
            _ = v20
        }
        v5 -= v5
        _ = fi17
    }
    var v21 int32 = 42
    v22 := []int{-737}
    switch 127 {
    case 217:
        for fi23 := 0; fi23 < 20; fi23++ {
            if true {
                switch "hello" {
                case "":
                    var v24 bool = ((v1 << ((530) & 63)) < (^(v1)))
                    _ = v24
                case "a":
                    v25 := GenS1{F0: 0, F1: v16, F2: -1}
                    _ = v25
                default:
                    v1 += v1
                }
                if (-2.25 <= 0.1) {
                    switch "a" {
                    case "hello":
                        switch ("foo bar" + ("\n\t" + "a")) {
                        case "":
                            for fi26 := 0; fi26 < 12; fi26++ {
                                var v27 string = "a"
                                v27 = v27
                                if (true || false) {
                                    v28 := v4[v16]
                                    v21 += v21
                                    switch ((v27 + v27) + (((v27 + (v27 + "foo bar")) + v27) + (v27 + v27))) {
                                    case "foo bar":
                                        var v29 int = (^(fi26))
                                        _ = v29
                                    case "":
                                        if false {
                                            v30 := GenS1{F0: 106, F1: v16, F2: -573}
                                            _ = v30
                                        } else {
                                            _ = len(v22)
                                            if true {
                                                delete(v2, v27)
                                                v22 = append(v22, (fi23 / (((fi26 >> ((541) & 63))) | 1)))
                                                if true {
                                                    fmt.Println(v5, v1)
                                                    v28 = v16
                                                    var v31 float32 = float32(-2.25)
                                                    _ = v31
                                                } else {
                                                    var v32 int = fi23
                                                    _ = v32
                                                }
                                                var v33 int32 = (v21 ^ v5)
                                                if ((1.5 <= 0.0) && (uint8(v16) < 2)) {
                                                    if ((-2147483648 > (v5 - (v21 - v5))) || (false && false)) {
                                                        delete(v2, "\n\t")
                                                    }
                                                    v33 = (v33 / ((v33) | 1))
                                                    var v34 int = fi23
                                                    _ = v34
                                                }
                                                v22 = append(v22, fi23)
                                                fmt.Println(v27)
                                                _ = len(v22)
                                                v3[(v28 * v28)] = (v28 << ((641) & 63))
                                                if false {
                                                    v35 := GenS1{F0: 255, F1: v16, F2: 43}
                                                    _ = v35
                                                }
                                                if len(v22) > 0 {
                                                    v22[0] = fi26
                                                }
                                                _ = v33
                                            } else {
                                                switch v27 {
                                                case "hello":
                                                    var v36 float32 = 1.5
                                                    _ = v36
                                                }
                                                v1 -= uint16(v16)
                                                v37 := GenS1{F0: 1, F1: v16, F2: -1}
                                                _ = v37
                                            }
                                            _ = len(v22)
                                        }
                                    case "a":
                                        if (-133 < -9223372036854775808) {
                                            var v38 float64 = float64((0.0 + float32(3.14159)))
                                            if false {
                                                if len(v22) >= 2 {
                                                    v22 = v22[1:len(v22)]
                                                }
                                            }
                                            _ = v38
                                        } else {
                                            var v39 int64 = 43
                                            v21 -= (v21 % (((v21 >> ((65535) & 63))) | 1))
                                            if (108 > 100) {
                                                v21 += v5
                                                var v40 int64 = 951
                                                _ = v40
                                            }
                                            _ = len(v22)
                                            _ = v39
                                        }
                                    }
                                    _ = v28
                                } else {
                                    if true {
                                        var v41 int = 793
                                        v5 = 42
                                        _ = v41
                                    }
                                    if (127 != -125) {
                                        v22 = append(v22, 521)
                                        if false {
                                            if true {
                                                var v42 uint = 2
                                                var v43 bool = (201 < -9223372036854775808)
                                                v16 = 1
                                                _ = v42
                                                _ = v43
                                            } else {
                                                v21 = 1
                                            }
                                            v44 := GenS2{F0: true, F1: uint32(v1), F2: fi23}
                                            v45 := GenS3{F0: -123, F1: (("" + (v27 + v27)) + (v27 + v27))}
                                            _ = v44
                                            _ = v45
                                        } else {
                                            v21 -= -500
                                            var v46 float64 = 0.1
                                            v47 := GenS4{F0: (v5 + -575), F1: v1}
                                            _ = v46
                                            _ = v47
                                        }
                                        var v48 int = fi23
                                        _ = v48
                                    }
                                    v5 = v21
                                    switch v27 {
                                    case "\n\t":
                                        var v49 uint64 = 580
                                        _ = v49
                                    default:
                                        if (1 != 484) {
                                            v27 = ((v27 + "a") + "\n\t")
                                            var v50 bool = (true && false)
                                            _ = v50
                                        }
                                    }
                                    if true {
                                        if (42 >= 777) {
                                            var v51 int16 = -1
                                            if ((fi23 >> ((uint(v1)) & 63)) <= fi26) {
                                                var v52 int = fi26
                                                if (false || (!true)) {
                                                    if true {
                                                        fmt.Println(v21, v52)
                                                    }
                                                } else {
                                                    switch (fi23 ^ fi26) {
                                                    case -2147483648:
                                                        delete(v3, v16)
                                                    case -257:
                                                        var v53 uint = 4294967295
                                                        _ = v53
                                                    }
                                                    var v54 int16 = v51
                                                    var v55 uint32 = (v16 / ((1) | 1))
                                                    _ = v54
                                                    _ = v55
                                                }
                                                _ = v52
                                            }
                                            _ = v51
                                        }
                                        var v56 uint8 = 2
                                        v2["\n\t"] = (fi26 | fi23)
                                        _ = v56
                                    } else {
                                        v57 := GenS2{F0: true, F1: (v16 | (v16 % ((v16) | 1))), F2: fi23}
                                        v1 = v1
                                        _ = v57
                                    }
                                    v5 = v5
                                }
                                _ = fi26
                                _ = v27
                            }
                        default:
                            v5 -= v5
                        }
                    case "\n\t":
                        v16 = (v16 | (v16 + (v16 / ((743) | 1))))
                    case "a":
                        var v58 uint32 = 2147483647
                        _ = v58
                    }
                }
                if ((42 <= 127) || (!true)) {
                    v16 = 1
                } else {
                    for fi59 := 0; fi59 < 13; fi59++ {
                        var v60 string = (("" + "") + "\n\t")
                        var v61 uint8 = 255
                        var v62 int16 = 32767
                        if len(v22) > 0 {
                            v22[0] = fi59
                        }
                        _ = fi59
                        _ = v60
                        _ = v61
                        _ = v62
                    }
                }
            } else {
                v16 += v16
                var v63 float32 = 1e10
                fmt.Println(fi23, v5, fi23)
                _ = v63
            }
            if false {
                if false {
                    v5 += v5
                    var v64 float32 = 1.5
                    switch (fi23 / (((^(fi23))) | 1)) {
                    case 32768:
                        if (255 != 255) {
                            var v65 int64 = -1
                            _ = v65
                        }
                    case 842:
                        var v66 uint64 = 42
                        _ = v66
                    case 990:
                        if true {
                            var v67 uint8 = 212
                            _ = v67
                        } else {
                            var v68 bool = ((v1 * 1) == 32767)
                            _ = v68
                        }
                    }
                    _ = v64
                } else {
                    var v69 uint32 = v16
                    _ = v69
                }
            } else {
                v21 = (^(v21))
            }
            v70 := GenS3{F0: 1, F1: "hello"}
            _ = fi23
            _ = v70
        }
    case 32768:
        for fi71 := 0; fi71 < 15; fi71++ {
            fmt.Println(v16)
            _ = fi71
        }
    case 128:
        var v72 int64 = -9223372036854775808
        _ = v72
    }
    var v73 uint8 = 127
    v74 := GenS5{F0: 255}
    v75 := GenS2{F0: true, F1: (v16 % ((v16) | 1)), F2: 32768}
    v76 := GenS4{F0: v21, F1: 42}
    if true {
        if (!true) {
            v1 -= (v1 | v1)
            if ((0.0 + (1.5 + (0.0 + 3.14159))) != 1e10) {
                for fi77 := 0; fi77 < 11; fi77++ {
                    v21 -= 2147483647
                    if false {
                        if len(v22) > 0 {
                            v22[0] = (^(fi77))
                        }
                        v78 := GenS5{F0: 65}
                        _ = v78
                    }
                    _ = fi77
                }
            }
        } else {
            if false {
                if (2147483647 <= (v5 % ((v21) | 1))) {
                    v22 = append(v22, 4294967295)
                } else {
                    if true {
                        v22 = append(v22, 127)
                        v79 := GenS4{F0: 0, F1: (^(v1))}
                        _ = v79
                    }
                    for fi80 := 0; fi80 < 1; fi80++ {
                        if (!true) {
                            switch "hello" {
                            case "\n\t":
                                v73 -= v73
                            case "foo bar":
                                if false {
                                    if true {
                                        var v81 float32 = 1.5
                                        _ = v81
                                    }
                                }
                            default:
                                switch "" {
                                case "hello":
                                    if (!((true || ("foo bar" != "a")) && ("hello" != ("\n\t" + (("hello" + "hello") + "\n\t"))))) {
                                        for fi82 := 0; fi82 < 7; fi82++ {
                                            var v83 float64 = 1.5
                                            v84 := GenS6{F0: false, F1: (true || false)}
                                            _ = fi82
                                            _ = v83
                                            _ = v84
                                        }
                                        if (true || true) {
                                            if len(v22) > 0 {
                                                v22[0] = (^(fi80))
                                            }
                                        } else {
                                            v85 := GenS3{F0: int16(v1), F1: ""}
                                            fmt.Println(v73, v73, v1)
                                            _ = v85
                                        }
                                        v21 -= (v21 - v5)
                                    }
                                case "":
                                    for fi86 := 0; fi86 < 2; fi86++ {
                                        v21 = (v21 >> ((744) & 63))
                                        fmt.Println(v1, fi86, v73)
                                        _ = len(v22)
                                        _ = fi86
                                    }
                                case "foo bar":
                                    var v87 float64 = 3.14159
                                    _ = v87
                                default:
                                    v88 := GenS7{F0: ""}
                                    _ = v88
                                }
                            }
                            var v89 int16 = 0
                            _ = v89
                        }
                        _ = fi80
                    }
                }
            } else {
                if true {
                    if false {
                        if false {
                            for fi90 := 0; fi90 < 17; fi90++ {
                                var v91 float32 = 0.0
                                _ = fi90
                                _ = v91
                            }
                        } else {
                            if (!(!(true && (255 > 65535)))) {
                                var v92 int8 = 0
                                _ = v92
                            }
                            v73 -= genFn1(18446744073709551615, 0)
                        }
                        if (v73 > (v73 - v73)) {
                            var v93 float64 = (-2.25 * 3.14159)
                            v16 = (v16 << ((250) & 63))
                            _ = v93
                        } else {
                            var v94 string = ""
                            var v95 float32 = ((-2.25 + 0.0) + 0.1)
                            switch "a" {
                            case "foo bar":
                                if (-1 <= 1) {
                                    if (776 <= 708) {
                                        _ = len(v22)
                                        if (0 > -225) {
                                            switch (v94 + v94) {
                                            case "hello":
                                                for fi96 := 0; fi96 < 14; fi96++ {
                                                    v97 := GenS8{F0: (^(v1)), F1: 1}
                                                    v16 += (v16 * (v16 >> ((42) & 63)))
                                                    v98 := GenS6{F0: true, F1: ("hello" > (v94 + v94))}
                                                    _ = fi96
                                                    _ = v97
                                                    _ = v98
                                                }
                                            case "foo bar":
                                                var v99 uint8 = genFn1(255, -9223372036854775808)
                                                _ = v99
                                            case "a":
                                                for fi100 := 0; fi100 < 15; fi100++ {
                                                    var v101 float32 = v95
                                                    if false {
                                                        if true {
                                                            v102 := GenS7{F0: v94}
                                                            var v103 uint16 = v1
                                                            _ = v102
                                                            _ = v103
                                                        }
                                                        var v104 int64 = -39
                                                        _ = v104
                                                    }
                                                    var v105 int32 = int32(fi100)
                                                    _ = fi100
                                                    _ = v101
                                                    _ = v105
                                                }
                                            }
                                            for fi106 := 0; fi106 < 2; fi106++ {
                                                var v107 int64 = -182
                                                for fi108 := 0; fi108 < 19; fi108++ {
                                                    var v109 string = (v94 + "a")
                                                    v73 = (v73 * 13)
                                                    if (!(0 < 448)) {
                                                        v110 := GenS2{F0: true, F1: (v16 * (^(v16))), F2: (fi108 % ((2147483647) | 1))}
                                                        _ = v110
                                                    } else {
                                                        v107 = v107
                                                    }
                                                    var v111 float64 = 0.1
                                                    v109 = v94
                                                    _ = fi108
                                                    _ = v109
                                                    _ = v111
                                                }
                                                fmt.Println(v21, fi106, v94)
                                                _ = fi106
                                                _ = v107
                                            }
                                        } else {
                                            switch "a" {
                                            case "\n\t":
                                                var v112 int8 = 127
                                                _ = v112
                                            default:
                                                if len(v22) >= 2 {
                                                    v22 = v22[1:len(v22)]
                                                }
                                            }
                                            v113 := GenS5{F0: 255}
                                            v114 := GenS5{F0: 2}
                                            _ = v113
                                            _ = v114
                                        }
                                    } else {
                                        v115 := GenS6{F0: (!true), F1: true}
                                        _ = len(v2)
                                        switch 42 {
                                        case 2147483647:
                                            v94 = v94
                                        case 32767:
                                            var v116 int64 = int64(v73)
                                            _ = v116
                                        case 760:
                                            v94 = v94
                                        }
                                        _ = v115
                                    }
                                    var v117 int = 2147483647
                                    _ = v117
                                }
                            default:
                                var v118 string = v94
                                _ = v118
                            }
                            _ = v94
                            _ = v95
                        }
                    }
                    v1 -= 0
                }
                switch 42 {
                case -9223372036854775808:
                    if false {
                        v2["hello"] = -1
                    } else {
                        v119 := GenS5{F0: 65535}
                        v21 = int32(v73)
                        fmt.Println(v73)
                        _ = v119
                    }
                case 45:
                    var v120 uint64 = 1
                    _ = v120
                }
            }
        }
        v5 += (v21 & (v21 << ((0) & 63)))
    } else {
        var v121 uint8 = (v73 ^ (v73 ^ genFn1(440, 886)))
        fmt.Println(v21, v1)
        _ = v121
    }
    for fi122 := 0; fi122 < 8; fi122++ {
        for fi123 := 0; fi123 < 20; fi123++ {
            v124 := GenS3{F0: 0, F1: (("foo bar" + "foo bar") + "")}
            _ = fi123
            _ = v124
        }
        for fi125 := 0; fi125 < 6; fi125++ {
            v126 := GenS2{F0: (!true), F1: v16, F2: fi122}
            v2[("" + "foo bar")] = 9223372036854775807
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            _ = fi125
            _ = v126
        }
        _ = fi122
    }
    for fi127 := 0; fi127 < 7; fi127++ {
        v73 += (v73 << ((228) & 63))
        _ = fi127
    }
    switch "hello" {
    case "":
        var v128 uint32 = (v16 % ((65535) | 1))
        _ = v128
    case "\n\t":
        if (-13 > 5) {
            if true {
                v16 = 2
                v21 = (v5 ^ v21)
            }
            if (!(1 < 859)) {
                v4[(^(v16))] = v16
                v129 := GenS9{F0: "a", F1: v16, F2: true}
                _ = v129
            }
        } else {
            v5 = v5
        }
    case "a":
        v22 = append(v22, -1)
    }
    fmt.Println(v21)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v73
    _ = v74
    _ = v75
    _ = v76
}
