package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 int8; F1 bool; F2 uint }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS10 struct { F0 uint8; F1 uint16; F2 uint }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint; F1 uint32 }
type GenS13 struct { F0 uint8; F1 int8 }
type GenS14 struct { F0 uint8; F1 int64; F2 int16 }
type GenS15 struct { F0 uint8; F1 int; F2 uint8 }
type GenS16 struct { F0 string; F1 string }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 int64; F1 int64 }
type GenS19 struct { F0 int }
type GenS20 struct { F0 uint16 }
type GenS21 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 1
    var v3 uint32 = 2147483647
    v2 = (v2 & -2147483648)
    var v4 uint64 = 4294967295
    for fi5 := 0; fi5 < 16; fi5++ {
        v4 = (^(v4))
        var v6 uint64 = v4
        _ = fi5
        _ = v6
    }
    var v7 int = 9223372036854775807
    v8 := GenS1{F0: (v4 / ((247) | 1))}
    switch 2147483647 {
    case -247:
        if true {
            if true {
                switch "" {
                case "":
                    if ("a" != "\n\t") {
                        for fi9 := 0; fi9 < 9; fi9++ {
                            switch "foo bar" {
                            case "":
                                for fi10 := 0; fi10 < 20; fi10++ {
                                    switch fi9 {
                                    case 827:
                                        if (!true) {
                                            v11 := GenS2{F0: 814}
                                            switch "foo bar" {
                                            case "hello":
                                                if false {
                                                    v3 = v3
                                                    var v12 int = v7
                                                    switch 9223372036854775807 {
                                                    case 640:
                                                        _ = len(v1)
                                                    }
                                                    var v13 uint = 65535
                                                    _ = v12
                                                    _ = v13
                                                } else {
                                                    if false {
                                                        var v14 uint8 = 196
                                                        v15 := v1[v3]
                                                        _ = v14
                                                        _ = v15
                                                    }
                                                }
                                            case "a":
                                                v1[v3] = "\n\t"
                                            case "":
                                                if true {
                                                    if true {
                                                        _ = len(v1)
                                                        v16 := []int{(v7 ^ (^(fi9))), (-(fi9))}
                                                        _ = v16
                                                    }
                                                    fmt.Println(v4, v3)
                                                    if (fi9 > -9223372036854775808) {
                                                        delete(v1, (^(v3)))
                                                        var v17 int8 = 0
                                                        switch "foo bar" {
                                                        case "":
                                                            if false {
                                                                switch (^(v7)) {
                                                                case 32767:
                                                                    v2 += (v2 ^ (v2 & (v2 | (v2 | v2))))
                                                                case 1:
                                                                    v18 := []int{}
                                                                    _ = v18
                                                                }
                                                            }
                                                        case "\n\t":
                                                            if (1 > 4294967295) {
                                                                switch fi10 {
                                                                case -361:
                                                                    if ("" >= "foo bar") {
                                                                        fmt.Println(v2, v4)
                                                                        v19 := []int{95, 1, v7}
                                                                        if (937 <= 629) {
                                                                            v3 -= 964
                                                                            fmt.Println(v17, v2)
                                                                        }
                                                                        _ = v19
                                                                    } else {
                                                                        if (((false && false) && true) || true) {
                                                                            v20 := GenS3{F0: (v3 << ((1) & 63)), F1: -432}
                                                                            _ = v20
                                                                        } else {
                                                                            v3 += (v3 - v3)
                                                                            if (!false) {
                                                                                if true {
                                                                                    switch 1 {
                                                                                    case -642:
                                                                                        v7 += v7
                                                                                    case 2147483647:
                                                                                        if (((1e10 == ((0.1 / 1.0) - float64(((1e10 / (0.1 * (((((-2.25 * 0.1) / 3.14159) + (3.14159 - 0.0)) - 0.0) - 3.14159))) + (float32((0.0 + ((((-2.25 / 3.14159) / -2.25) / 1.0) / 0.1))) / 0.1))))) || (false && ((v3 >> ((1) & 63)) != v3))) || (!false)) {
                                                                                            if true {
                                                                                                fmt.Println(fi9, v17)
                                                                                            } else {
                                                                                                if true {
                                                                                                    v21 := []int{(v7 * fi9), int(v3)}
                                                                                                    v1[v3] = ""
                                                                                                    _ = v21
                                                                                                } else {
                                                                                                    var v22 float64 = ((0.0 + 1.5) - 1e10)
                                                                                                    var v23 int64 = 867
                                                                                                    fmt.Println(v7, fi10, v22)
                                                                                                    _ = v22
                                                                                                    _ = v23
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(v17)
                                                                                            v24 := []int{v7}
                                                                                            _ = v24
                                                                                        } else {
                                                                                            var v25 uint64 = v4
                                                                                            _ = v25
                                                                                        }
                                                                                    }
                                                                                    switch "hello" {
                                                                                    case "hello":
                                                                                        switch "a" {
                                                                                        case "hello":
                                                                                            v7 -= fi9
                                                                                        case "a":
                                                                                            switch ("" + "\n\t") {
                                                                                            case "":
                                                                                                fmt.Println(fi10, v2, v4)
                                                                                            case "a":
                                                                                                v26 := GenS2{F0: 1}
                                                                                                _ = v26
                                                                                            case "hello":
                                                                                                if ((true || false) || true) {
                                                                                                    v27 := GenS1{F0: v4}
                                                                                                    _ = v27
                                                                                                }
                                                                                            }
                                                                                        case "":
                                                                                            var v28 float64 = 0.1
                                                                                            _ = v28
                                                                                        default:
                                                                                            v29 := GenS4{F0: -1, F1: int64(v17)}
                                                                                            _ = v29
                                                                                        }
                                                                                    case "\n\t":
                                                                                        fmt.Println(fi9, v4)
                                                                                    case "":
                                                                                        if false {
                                                                                            v30 := GenS1{F0: (v4 ^ 1)}
                                                                                            var v31 int8 = 23
                                                                                            if true {
                                                                                                var v32 uint16 = 904
                                                                                                var v33 uint64 = v4
                                                                                                v34 := []int{(v7 | (fi10 >> ((411) & 63))), 2147483647}
                                                                                                var v35 uint = 255
                                                                                                v36 := v1[(v3 >> ((v35) & 63))]
                                                                                                _ = v32
                                                                                                _ = v33
                                                                                                _ = v34
                                                                                                _ = v35
                                                                                                _ = v36
                                                                                            } else {
                                                                                                v31 = v17
                                                                                                v17 += -1
                                                                                                if true {
                                                                                                    var v37 int8 = -128
                                                                                                    if true {
                                                                                                        var v38 uint8 = 7
                                                                                                        var v39 int32 = v2
                                                                                                        _ = v38
                                                                                                        _ = v39
                                                                                                    } else {
                                                                                                        v17 = v31
                                                                                                        var v40 int32 = v2
                                                                                                        if false {
                                                                                                            var v41 uint64 = v4
                                                                                                            _ = v41
                                                                                                        }
                                                                                                        if false {
                                                                                                            var v42 uint = 8
                                                                                                            _ = v42
                                                                                                        }
                                                                                                        _ = v40
                                                                                                    }
                                                                                                    _ = v37
                                                                                                } else {
                                                                                                    var v43 int32 = v2
                                                                                                    if (v4 == 1) {
                                                                                                        var v44 uint32 = (v3 ^ (v3 % (((v3 / ((v3) | 1))) | 1)))
                                                                                                        if (0 <= 0) {
                                                                                                            v45 := GenS4{F0: 0, F1: 9223372036854775807}
                                                                                                            switch "" {
                                                                                                            case "":
                                                                                                                v46 := []int{566}
                                                                                                                _ = v46
                                                                                                            }
                                                                                                            var v47 int = fi10
                                                                                                            _ = v45
                                                                                                            _ = v47
                                                                                                        } else {
                                                                                                            v48 := GenS4{F0: -815, F1: 894}
                                                                                                            _ = v48
                                                                                                        }
                                                                                                        if (false || true) {
                                                                                                            fmt.Println(v43, fi9, v7)
                                                                                                            var v49 string = "\n\t"
                                                                                                            _ = v49
                                                                                                        }
                                                                                                        _ = v44
                                                                                                    } else {
                                                                                                        fmt.Println(v7, fi10, fi9)
                                                                                                        if false {
                                                                                                            if false {
                                                                                                                v50 := GenS2{F0: 255}
                                                                                                                var v51 int16 = 1
                                                                                                                _ = v50
                                                                                                                _ = v51
                                                                                                            } else {
                                                                                                                var v52 string = "\n\t"
                                                                                                                if true {
                                                                                                                    switch fi10 {
                                                                                                                    case -196:
                                                                                                                        var v53 uint16 = 499
                                                                                                                        _ = v53
                                                                                                                    case 1:
                                                                                                                        var v54 uint = 65535
                                                                                                                        _ = v54
                                                                                                                    case -2147483648:
                                                                                                                        var v55 uint8 = 42
                                                                                                                        _ = v55
                                                                                                                    }
                                                                                                                    v56 := []int{(fi9 / ((fi10) | 1)), (v7 << ((839) & 63)), fi9}
                                                                                                                    if len(v56) > 0 {
                                                                                                                        v56[0] = 9223372036854775807
                                                                                                                    }
                                                                                                                    _ = v56
                                                                                                                }
                                                                                                                _ = v52
                                                                                                            }
                                                                                                            v57 := GenS4{F0: -782, F1: 865}
                                                                                                            fmt.Println(v17)
                                                                                                            _ = v57
                                                                                                        } else {
                                                                                                            var v58 uint8 = 199
                                                                                                            v17 = (v17 * v17)
                                                                                                            fmt.Println(v4, fi9, v43)
                                                                                                            v3 -= 2
                                                                                                            var v59 float64 = 0.1
                                                                                                            _ = v58
                                                                                                            _ = v59
                                                                                                        }
                                                                                                        v60 := v1[(v3 / (((v3 & (v3 * (v3 ^ (v3 / ((v3) | 1)))))) | 1))]
                                                                                                        _ = v60
                                                                                                    }
                                                                                                    _ = v43
                                                                                                }
                                                                                            }
                                                                                            v7 -= (fi10 * v7)
                                                                                            _ = v30
                                                                                            _ = v31
                                                                                        } else {
                                                                                            if true {
                                                                                                var v61 uint16 = 65535
                                                                                                if false {
                                                                                                    var v62 string = ("foo bar" + "")
                                                                                                    fmt.Println(v62)
                                                                                                    _ = v62
                                                                                                }
                                                                                                if ((18446744073709551615 >= 4294967295) && (1 == -1)) {
                                                                                                    switch "foo bar" {
                                                                                                    case "foo bar":
                                                                                                        fmt.Println(v17, fi10)
                                                                                                    case "a":
                                                                                                        var v63 int8 = (v17 - (v17 & (v17 / ((v17) | 1))))
                                                                                                        _ = v63
                                                                                                    case "hello":
                                                                                                        if (!((2 < 0) && true)) {
                                                                                                            var v64 int8 = (v17 / ((-1) | 1))
                                                                                                            _ = v64
                                                                                                        } else {
                                                                                                            v2 += v2
                                                                                                            v7 += 42
                                                                                                            var v65 uint32 = v3
                                                                                                            _ = v65
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                _ = v61
                                                                                            }
                                                                                            var v66 int16 = -553
                                                                                            var v67 int16 = (v66 | (v66 * v66))
                                                                                            var v68 uint8 = 128
                                                                                            _ = v66
                                                                                            _ = v67
                                                                                            _ = v68
                                                                                        }
                                                                                    default:
                                                                                        var v69 string = "hello"
                                                                                        _ = v69
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(fi9, v7)
                                                                                    v4 = (v4 / ((v4) | 1))
                                                                                }
                                                                                var v70 int64 = 1
                                                                                switch 1 {
                                                                                case 514:
                                                                                    switch -9223372036854775808 {
                                                                                    case 42:
                                                                                        fmt.Println(v7, fi9)
                                                                                    case -31:
                                                                                        v71 := GenS4{F0: v70, F1: (v70 | 1)}
                                                                                        _ = v71
                                                                                    case 4294967295:
                                                                                        var v72 uint = 2
                                                                                        _ = v72
                                                                                    default:
                                                                                        switch (v7 & 0) {
                                                                                        case -920:
                                                                                            v73 := GenS5{F0: uint32(v4), F1: 526, F2: (v70 | v70)}
                                                                                            _ = v73
                                                                                        default:
                                                                                            var v74 int16 = 833
                                                                                            _ = v74
                                                                                        }
                                                                                    }
                                                                                case -1:
                                                                                    v4 = 4294967295
                                                                                }
                                                                                v75 := GenS2{F0: 32767}
                                                                                _ = v70
                                                                                _ = v75
                                                                            }
                                                                            var v76 uint64 = v4
                                                                            var v77 int16 = -32768
                                                                            _ = v76
                                                                            _ = v77
                                                                        }
                                                                        v78 := GenS3{F0: v3, F1: (v2 & (v2 ^ 2147483647))}
                                                                        _ = v78
                                                                    }
                                                                case 9223372036854775807:
                                                                    var v79 string = "hello"
                                                                    _ = v79
                                                                default:
                                                                    var v80 uint = 755
                                                                    _ = v80
                                                                }
                                                                if (!((!false) && ((v17 | v17) <= v17))) {
                                                                    fmt.Println(fi9, v3, v2)
                                                                    v81 := []int{v7, (fi9 & int(v2)), v7, 2147483647}
                                                                    switch (v7 | v7) {
                                                                    case 9223372036854775807:
                                                                        var v82 int8 = 1
                                                                        _ = v82
                                                                    case -183:
                                                                        if (true || (v4 != v4)) {
                                                                            v4 += v4
                                                                            if (1.5 >= 1.5) {
                                                                                if true {
                                                                                    if (false || (!true)) {
                                                                                        if true {
                                                                                            switch fi9 {
                                                                                            case -728:
                                                                                                var v83 float32 = 0.0
                                                                                                _ = v83
                                                                                            }
                                                                                            v7 = v7
                                                                                            if (!(0 != 0)) {
                                                                                                v3 = v3
                                                                                            }
                                                                                        }
                                                                                        var v84 uint32 = (v3 & v3)
                                                                                        _ = v84
                                                                                    } else {
                                                                                        v85 := GenS6{F0: (^(v4)), F1: (^(v17))}
                                                                                        _ = v85
                                                                                    }
                                                                                    v4 = v4
                                                                                }
                                                                                var v86 int16 = -739
                                                                                _ = v86
                                                                            }
                                                                        } else {
                                                                            var v87 uint8 = 204
                                                                            if (1 < 2) {
                                                                                switch fi9 {
                                                                                case 127:
                                                                                    var v88 bool = (((v87 <= v87) && (float64(float32(1.5)) == float64(0.1))) && true)
                                                                                    _ = v88
                                                                                case 283:
                                                                                    if (int64(v87) >= -9223372036854775808) {
                                                                                        if true {
                                                                                            var v89 int64 = 1
                                                                                            _ = v89
                                                                                        }
                                                                                        var v90 int8 = v17
                                                                                        var v91 int32 = (v2 + v2)
                                                                                        if (0 == 265) {
                                                                                            switch "hello" {
                                                                                            case "a":
                                                                                                v4 -= v4
                                                                                            case "foo bar":
                                                                                                var v92 int16 = -649
                                                                                                var v93 uint = 42
                                                                                                _ = v92
                                                                                                _ = v93
                                                                                            case "hello":
                                                                                                switch "hello" {
                                                                                                case "hello":
                                                                                                    var v94 int8 = v90
                                                                                                    _ = v94
                                                                                                case "":
                                                                                                    switch "a" {
                                                                                                    case "foo bar":
                                                                                                        switch ("a" + ("" + "")) {
                                                                                                        case "a":
                                                                                                            fmt.Println(fi10)
                                                                                                        case "hello":
                                                                                                            var v95 uint64 = 18446744073709551615
                                                                                                            _ = v95
                                                                                                        case "":
                                                                                                            delete(v1, (v3 + (v3 | v3)))
                                                                                                        default:
                                                                                                            v96 := v1[v3]
                                                                                                            _ = v96
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    v91 = v91
                                                                                                }
                                                                                            }
                                                                                            if ((962 > 255) || false) {
                                                                                                var v97 int64 = -9223372036854775808
                                                                                                _ = v97
                                                                                            }
                                                                                        } else {
                                                                                            if true {
                                                                                                switch (("foo bar" + "\n\t") + "a") {
                                                                                                case "\n\t":
                                                                                                    v87 += 0
                                                                                                case "":
                                                                                                    var v98 uint = 18446744073709551615
                                                                                                    _ = v98
                                                                                                case "a":
                                                                                                    v99 := GenS2{F0: 1}
                                                                                                    _ = v99
                                                                                                default:
                                                                                                    v81 = append(v81, (^(v7)))
                                                                                                }
                                                                                                var v100 float64 = (((((((float64(0.0) + ((1.5 / float64((-2.25 + (0.1 - (-2.25 / 1e10))))) - 0.1)) + 0.0) - 0.0) * (1.5 * (3.14159 / 1e10))) - 1.5) - ((3.14159 * 1e10) / 1.0)) / 0.1)
                                                                                                var v101 int32 = v91
                                                                                                _ = v100
                                                                                                _ = v101
                                                                                            }
                                                                                            v102 := GenS7{F0: (v90 | (v90 | -128)), F1: (!true), F2: 18446744073709551615}
                                                                                            _ = v102
                                                                                        }
                                                                                        _ = v90
                                                                                        _ = v91
                                                                                    }
                                                                                }
                                                                                v3 -= (v3 - 1)
                                                                                var v103 float32 = float32((-2.25 / 0.1))
                                                                                _ = v103
                                                                            } else {
                                                                                if (0 != 1) {
                                                                                    v104 := GenS2{F0: 911}
                                                                                    switch "hello" {
                                                                                    case "\n\t":
                                                                                        if len(v81) >= 2 {
                                                                                            v81 = v81[1:len(v81)]
                                                                                        }
                                                                                    }
                                                                                    _ = v104
                                                                                } else {
                                                                                    fmt.Println(fi10)
                                                                                    v2 -= (v2 - 1)
                                                                                }
                                                                                var v105 int8 = (v17 >> ((494) & 63))
                                                                                var v106 int8 = v105
                                                                                _ = v105
                                                                                _ = v106
                                                                            }
                                                                            _ = v87
                                                                        }
                                                                    default:
                                                                        var v107 int = fi9
                                                                        _ = v107
                                                                    }
                                                                    _ = v81
                                                                } else {
                                                                    fmt.Println(v2)
                                                                    var v108 float64 = 1.5
                                                                    fmt.Println(v108, v4)
                                                                    if (!(!((-6 <= (-(v17))) && (0 == 331)))) {
                                                                        if false {
                                                                            v109 := []int{1, v7, (-(fi9)), v7}
                                                                            fmt.Println(v7)
                                                                            v110 := GenS2{F0: 32767}
                                                                            _ = v109
                                                                            _ = v110
                                                                        } else {
                                                                            var v111 uint32 = (v3 - v3)
                                                                            v2 += 1
                                                                            switch (v7 | (fi10 * (v7 << ((929) & 63)))) {
                                                                            case 982:
                                                                                var v112 uint16 = 65535
                                                                                _ = v112
                                                                            case 32768:
                                                                                v113 := GenS8{F0: (false && false)}
                                                                                _ = v113
                                                                            default:
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    if false {
                                                                                        var v114 uint8 = 2
                                                                                        fmt.Println(v111, fi9)
                                                                                        v115 := GenS4{F0: 404, F1: -1}
                                                                                        var v116 string = "a"
                                                                                        _ = v114
                                                                                        _ = v115
                                                                                        _ = v116
                                                                                    } else {
                                                                                        v117 := GenS9{F0: 255, F1: 4294967295, F2: 255}
                                                                                        if ((v7 ^ v7) != (fi9 * int(v3))) {
                                                                                            v118 := v1[65535]
                                                                                            v1[(v111 - v3)] = ("foo bar" + v118)
                                                                                            v4 = v4
                                                                                            _ = v118
                                                                                        } else {
                                                                                            fmt.Println(v108)
                                                                                        }
                                                                                        var v119 uint = 0
                                                                                        _ = v117
                                                                                        _ = v119
                                                                                    }
                                                                                case "":
                                                                                    v17 = v17
                                                                                case "a":
                                                                                    var v120 uint32 = (v3 >> ((534) & 63))
                                                                                    _ = v120
                                                                                default:
                                                                                    v7 += fi9
                                                                                }
                                                                            }
                                                                            _ = v111
                                                                        }
                                                                        v1[(^(v3))] = "a"
                                                                    } else {
                                                                        switch ("\n\t" + "a") {
                                                                        case "":
                                                                            v3 += (v3 - (v3 & (^(v3))))
                                                                        default:
                                                                            v121 := GenS6{F0: (v4 ^ v4), F1: (^(v17))}
                                                                            v122 := []int{922, fi9}
                                                                            if (false && true) {
                                                                                _ = len(v122)
                                                                                var v123 uint32 = (v3 << ((2) & 63))
                                                                                _ = v123
                                                                            }
                                                                            delete(v1, v3)
                                                                            _ = v121
                                                                            _ = v122
                                                                        }
                                                                        v7 = 697
                                                                    }
                                                                    v3 -= (^(v3))
                                                                    _ = v108
                                                                }
                                                            } else {
                                                                if false {
                                                                    if true {
                                                                        if (0.1 <= -2.25) {
                                                                            var v124 int8 = v17
                                                                            _ = v124
                                                                        } else {
                                                                            v125 := GenS10{F0: 211, F1: 42, F2: 423}
                                                                            if false {
                                                                                var v126 uint16 = 0
                                                                                if (false || ((854 >= (-(v2))) || ((!(v126 < 918)) && (!(false && true))))) {
                                                                                    v3 = (v3 / (((v3 | 2147483647)) | 1))
                                                                                } else {
                                                                                    var v127 float64 = (1e10 - 1.5)
                                                                                    v128 := GenS8{F0: (!(!true))}
                                                                                    _ = v127
                                                                                    _ = v128
                                                                                }
                                                                                v126 += v126
                                                                                _ = v126
                                                                            }
                                                                            v17 = (v17 + (v17 << ((65535) & 63)))
                                                                            var v129 int32 = int32(v17)
                                                                            v129 = (v129 % ((2147483647) | 1))
                                                                            v17 = (v17 & v17)
                                                                            _ = v125
                                                                            _ = v129
                                                                        }
                                                                        v17 -= (v17 * (v17 % ((17) | 1)))
                                                                    }
                                                                    v4 = v4
                                                                }
                                                            }
                                                        case "foo bar":
                                                            if false {
                                                                v17 -= 1
                                                            }
                                                        default:
                                                            var v130 string = ("\n\t" + "a")
                                                            switch fi9 {
                                                            case 32767:
                                                                v17 -= (v17 << ((723) & 63))
                                                            default:
                                                                v1[v3] = v130
                                                                delete(v1, v3)
                                                                if (-743 == 448) {
                                                                    v131 := []int{(v7 >> ((114) & 63)), int(v2), -492}
                                                                    if false {
                                                                        if (1 != 42) {
                                                                            var v132 bool = true
                                                                            _ = v132
                                                                        } else {
                                                                            var v133 int32 = v2
                                                                            _ = v133
                                                                        }
                                                                    } else {
                                                                        var v134 bool = ((v4 % ((v4) | 1)) >= v4)
                                                                        fmt.Println(v4, v130)
                                                                        _ = v134
                                                                    }
                                                                    _ = v131
                                                                } else {
                                                                    v17 = (v17 >> ((42) & 63))
                                                                }
                                                                if false {
                                                                    fmt.Println(v7, fi10)
                                                                    v135 := v1[v3]
                                                                    if (true || (32767 > 347)) {
                                                                        fmt.Println(v3, v7, v3)
                                                                        switch fi10 {
                                                                        case 42:
                                                                            v1[(v3 | 65535)] = "\n\t"
                                                                        case -2147483648:
                                                                            fmt.Println(v135)
                                                                        case -813:
                                                                            delete(v1, (^(v3)))
                                                                        default:
                                                                            var v136 string = (("" + (v130 + "hello")) + ((v135 + "foo bar") + v130))
                                                                            _ = v136
                                                                        }
                                                                    }
                                                                    _ = v135
                                                                }
                                                            }
                                                            v130 = ("\n\t" + "hello")
                                                            _ = v130
                                                        }
                                                        _ = v17
                                                    } else {
                                                        switch (-(v7)) {
                                                        case -1:
                                                            v137 := GenS7{F0: 1, F1: (0.1 <= (1e10 * ((3.14159 / 1.5) + 0.1))), F2: 255}
                                                            _ = v137
                                                        default:
                                                            fmt.Println(v7, v7, fi9)
                                                        }
                                                        if false {
                                                            v138 := GenS11{F0: 255}
                                                            if (0.0 > 0.0) {
                                                                if true {
                                                                    var v139 bool = true
                                                                    v140 := []int{270, -1, 2147483647}
                                                                    if v139 {
                                                                        if true {
                                                                            _ = len(v1)
                                                                            if true {
                                                                                v141 := GenS12{F0: uint(v3), F1: 987}
                                                                                _ = v141
                                                                            } else {
                                                                                var v142 float32 = 0.1
                                                                                switch fi10 {
                                                                                case 127:
                                                                                    v3 = v3
                                                                                case -476:
                                                                                    var v143 float64 = 3.14159
                                                                                    _ = v143
                                                                                case 1:
                                                                                    v144 := GenS1{F0: 2}
                                                                                    _ = v144
                                                                                default:
                                                                                    var v145 int16 = 680
                                                                                    _ = v145
                                                                                }
                                                                                if v139 {
                                                                                    v2 += 0
                                                                                } else {
                                                                                    var v146 string = "hello"
                                                                                    v147 := GenS13{F0: 120, F1: -1}
                                                                                    v146 = v146
                                                                                    _ = v146
                                                                                    _ = v147
                                                                                }
                                                                                _ = v142
                                                                            }
                                                                        } else {
                                                                            v2 -= (-(v2))
                                                                            if v139 {
                                                                                var v148 uint16 = 2
                                                                                var v149 int = 32767
                                                                                _ = v148
                                                                                _ = v149
                                                                            } else {
                                                                                v150 := GenS7{F0: -128, F1: false, F2: 4294967295}
                                                                                _ = v150
                                                                            }
                                                                            if v139 {
                                                                                var v151 bool = false
                                                                                v3 += v3
                                                                                _ = v151
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v139
                                                                    _ = v140
                                                                } else {
                                                                    fmt.Println(v3, v4)
                                                                    switch -2147483648 {
                                                                    case -608:
                                                                        if (false && ((false || false) && ((v3 / ((v3) | 1)) == (v3 | (v3 >> ((2) & 63)))))) {
                                                                            var v152 bool = true
                                                                            var v153 uint = 18446744073709551615
                                                                            _ = v152
                                                                            _ = v153
                                                                        }
                                                                    case 128:
                                                                        v4 += v4
                                                                    }
                                                                    if ((fi10 * v7) >= fi10) {
                                                                        var v154 int16 = 42
                                                                        v154 -= v154
                                                                        _ = v154
                                                                    }
                                                                }
                                                                var v155 uint32 = 65535
                                                                _ = v155
                                                            }
                                                            _ = v138
                                                        }
                                                        if true {
                                                            if true {
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    var v156 int8 = 17
                                                                    _ = v156
                                                                case "\n\t":
                                                                    if false {
                                                                        if false {
                                                                            v7 += 9223372036854775807
                                                                            v157 := GenS5{F0: v3, F1: 42, F2: -1}
                                                                            v158 := []int{(-(v7)), (^(v7)), (fi9 - (^(fi10)))}
                                                                            if (!(!true)) {
                                                                                if (42 > 9223372036854775807) {
                                                                                    if len(v158) >= 2 {
                                                                                        v158 = v158[1:len(v158)]
                                                                                    }
                                                                                }
                                                                                v159 := GenS14{F0: 128, F1: 122, F2: 42}
                                                                                if (!(!false)) {
                                                                                    v160 := GenS15{F0: 42, F1: (fi10 / ((32768) | 1)), F2: 175}
                                                                                    v2 += v2
                                                                                    _ = v160
                                                                                } else {
                                                                                    v2 = (v2 % ((-1) | 1))
                                                                                    var v161 int64 = 1
                                                                                    fmt.Println(v3)
                                                                                    _ = v161
                                                                                }
                                                                                _ = v159
                                                                            } else {
                                                                                v158 = append(v158, v7)
                                                                                var v162 bool = (92 >= 18446744073709551615)
                                                                                var v163 uint8 = 127
                                                                                switch fi10 {
                                                                                case 2147483647:
                                                                                    switch (^(fi9)) {
                                                                                    case 2147483647:
                                                                                        fmt.Println(v2)
                                                                                    }
                                                                                }
                                                                                _ = v162
                                                                                _ = v163
                                                                            }
                                                                            v164 := GenS2{F0: 960}
                                                                            _ = v157
                                                                            _ = v158
                                                                            _ = v164
                                                                        } else {
                                                                            var v165 uint16 = 322
                                                                            switch v7 {
                                                                            case -1:
                                                                                switch v7 {
                                                                                case 4294967295:
                                                                                    switch ("hello" + "hello") {
                                                                                    case "\n\t":
                                                                                        v166 := GenS4{F0: 42, F1: -1}
                                                                                        _ = v166
                                                                                    case "hello":
                                                                                        v1[v3] = (("hello" + "\n\t") + ((((("" + "hello") + "\n\t") + "\n\t") + "") + "foo bar"))
                                                                                    }
                                                                                case 9223372036854775807:
                                                                                    v167 := v1[v3]
                                                                                    _ = v167
                                                                                }
                                                                                v168 := GenS2{F0: (v165 & v165)}
                                                                                _ = v168
                                                                            case 64:
                                                                                var v169 uint64 = v4
                                                                                _ = v169
                                                                            default:
                                                                                v165 = 725
                                                                            }
                                                                            _ = v165
                                                                        }
                                                                        _ = len(v1)
                                                                        v170 := GenS5{F0: (^(v3)), F1: 32767, F2: 42}
                                                                        _ = v170
                                                                    } else {
                                                                        var v171 uint32 = (^(v3))
                                                                        if true {
                                                                            switch ("\n\t" + "foo bar") {
                                                                            case "a":
                                                                                var v172 int = (fi10 + fi9)
                                                                                _ = v172
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v2, v4)
                                                                        }
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            var v173 uint32 = v3
                                                                            _ = v173
                                                                        case "a":
                                                                            var v174 string = (("\n\t" + "a") + ("a" + ""))
                                                                            _ = v174
                                                                        }
                                                                        _ = v171
                                                                    }
                                                                    v175 := []int{-178, (fi9 << ((255) & 63)), (fi9 << ((662) & 63)), v7}
                                                                    var v176 uint = 327
                                                                    _ = v175
                                                                    _ = v176
                                                                case "":
                                                                    if false {
                                                                        if true {
                                                                            v177 := v1[(v3 * (v3 | (v3 ^ uint32(v4))))]
                                                                            v177 = ("\n\t" + ("\n\t" + ((v177 + (v177 + "foo bar")) + "foo bar")))
                                                                            var v178 uint = 802
                                                                            _ = v177
                                                                            _ = v178
                                                                        }
                                                                        var v179 uint = 4294967295
                                                                        v180 := []int{(fi9 - (fi10 - int(v179))), (fi9 >> ((v179) & 63)), (fi9 + (v7 << ((v179) & 63)))}
                                                                        _ = v179
                                                                        _ = v180
                                                                    } else {
                                                                        v3 -= v3
                                                                    }
                                                                }
                                                                v181 := []int{int(v4), v7, v7, (v7 & (fi10 >> ((627) & 63)))}
                                                                _ = v181
                                                            } else {
                                                                v182 := []int{-1, 9223372036854775807, fi9}
                                                                v4 = (^(v4))
                                                                v183 := GenS9{F0: 65535, F1: 4294967295, F2: 1}
                                                                _ = v182
                                                                _ = v183
                                                            }
                                                            v184 := []int{}
                                                            fmt.Println(v3)
                                                            _ = v184
                                                        }
                                                    }
                                                } else {
                                                    v2 = (v2 + (v2 ^ 1))
                                                }
                                            }
                                            fmt.Println(v2, v7, v4)
                                            _ = v11
                                        } else {
                                            v7 -= 4294967295
                                            if false {
                                                var v185 bool = (100 <= 100)
                                                var v186 float32 = 1e10
                                                _ = len(v1)
                                                _ = v185
                                                _ = v186
                                            }
                                            if ("" != "\n\t") {
                                                var v187 float32 = 1.5
                                                v188 := []int{(^(v7))}
                                                fmt.Println(v4, v187, v2)
                                                _ = v187
                                                _ = v188
                                            }
                                        }
                                    case 747:
                                        if false {
                                            fmt.Println(v2, v4)
                                        }
                                    case -1:
                                        v2 += (v2 / ((42) | 1))
                                    }
                                    v189 := GenS16{F0: "", F1: "a"}
                                    _ = fi10
                                    _ = v189
                                }
                            case "a":
                                if (!(!(false || (-2.25 < 0.0)))) {
                                    v3 = 800
                                    v4 -= uint64(v3)
                                } else {
                                    for fi190 := 0; fi190 < 4; fi190++ {
                                        if (273 >= 0) {
                                            v191 := GenS1{F0: v4}
                                            _ = v191
                                        } else {
                                            v192 := GenS17{F0: (v3 & 2147483648)}
                                            _ = v192
                                        }
                                        delete(v1, 255)
                                        var v193 int8 = 5
                                        _ = fi190
                                        _ = v193
                                    }
                                }
                            case "foo bar":
                                fmt.Println(v4, v7, v7)
                            }
                            for fi194 := 0; fi194 < 16; fi194++ {
                                v195 := GenS8{F0: (v7 >= v7)}
                                _ = fi194
                                _ = v195
                            }
                            _ = fi9
                        }
                        var v196 uint32 = v3
                        var v197 int32 = v2
                        _ = v196
                        _ = v197
                    }
                case "\n\t":
                    v3 = (v3 >> ((0) & 63))
                default:
                    v198 := GenS18{F0: 42, F1: -86}
                    _ = v198
                }
            } else {
                var v199 int8 = 1
                var v200 uint16 = 255
                var v201 bool = (("foo bar" + "hello") > ("" + "\n\t"))
                _ = v199
                _ = v200
                _ = v201
            }
        } else {
            for fi202 := 0; fi202 < 9; fi202++ {
                var v203 int8 = 100
                if (0.1 < float64((-2.25 + (0.0 - 1e10)))) {
                    fmt.Println(v4, v203, v3)
                } else {
                    v204 := v1[v3]
                    for fi205 := 0; fi205 < 19; fi205++ {
                        var v206 int16 = -536
                        v207 := GenS7{F0: 8, F1: false, F2: 42}
                        _ = fi205
                        _ = v206
                        _ = v207
                    }
                    if (!true) {
                        switch ("" + v204) {
                        case "foo bar":
                            for fi208 := 0; fi208 < 1; fi208++ {
                                var v209 bool = (!("" != v204))
                                _ = fi208
                                _ = v209
                            }
                        default:
                            var v210 uint8 = uint8(v4)
                            _ = v210
                        }
                        v203 -= 73
                        v203 = int8(fi202)
                    }
                    _ = v204
                }
                _ = fi202
                _ = v203
            }
        }
    case 32768:
        switch ("hello" + "\n\t") {
        case "hello":
            var v211 int64 = 9223372036854775807
            _ = v211
        case "\n\t":
            for fi212 := 0; fi212 < 17; fi212++ {
                var v213 int = fi212
                _ = fi212
                _ = v213
            }
        case "":
            for fi214 := 0; fi214 < 15; fi214++ {
                if false {
                    var v215 bool = false
                    _ = v215
                } else {
                    var v216 int16 = 1
                    v217 := GenS3{F0: v3, F1: v2}
                    _ = v216
                    _ = v217
                }
                if true {
                    if true {
                        if false {
                            var v218 int8 = 1
                            if false {
                                v3 += 2
                            }
                            v219 := []int{0, fi214, int(v4), (fi214 % ((9223372036854775807) | 1))}
                            _ = v218
                            _ = v219
                        }
                    }
                    v4 += (v4 | (^(v4)))
                    for fi220 := 0; fi220 < 11; fi220++ {
                        if ((true || (-9223372036854775808 == -1)) && (("hello" != "") || true)) {
                            if true {
                                var v221 uint16 = 191
                                var v222 float64 = ((3.14159 * (((3.14159 + 0.0) - 0.0) - (1e10 * 3.14159))) / 0.1)
                                _ = v221
                                _ = v222
                            } else {
                                v1[(v3 + (v3 >> ((760) & 63)))] = ""
                            }
                        } else {
                            v223 := GenS17{F0: (v3 * v3)}
                            v224 := []int{(-(v7)), (^(v7)), 680, fi214}
                            var v225 uint32 = v3
                            _ = v223
                            _ = v224
                            _ = v225
                        }
                        if true {
                            v226 := GenS19{F0: (fi214 - (v7 % ((fi220) | 1)))}
                            switch "a" {
                            case "hello":
                                v227 := v1[v3]
                                _ = v227
                            }
                            fmt.Println(fi220, v2, v7)
                            _ = v226
                        }
                        _ = fi220
                    }
                } else {
                    switch ("a" + "foo bar") {
                    case "":
                        delete(v1, (v3 + (v3 / ((v3) | 1))))
                    }
                }
                _ = fi214
            }
        }
    case 2147483647:
        fmt.Println(v2, v3, v7)
    }
    if true {
        v3 = v3
        fmt.Println(v7, v3, v3)
        for fi228 := 0; fi228 < 3; fi228++ {
            var v229 int16 = -951
            if false {
                fmt.Println(fi228, v229)
                for fi230 := 0; fi230 < 17; fi230++ {
                    v231 := []int{4294967295}
                    _ = fi230
                    _ = v231
                }
            } else {
                v232 := v1[1]
                v229 += (v229 + v229)
                _ = v232
            }
            _ = fi228
            _ = v229
        }
    }
    switch "a" {
    case "foo bar":
        v2 = int32(v4)
    case "":
        for fi233 := 0; fi233 < 14; fi233++ {
            v234 := []int{fi233, 127, v7, fi233}
            _ = fi233
            _ = v234
        }
    case "a":
        v7 = v7
    default:
        switch (v7 & (v7 % ((v7) | 1))) {
        case 703:
            v235 := GenS20{F0: 1}
            _ = v235
        case -9223372036854775808:
            v3 = 2147483648
        }
    }
    var v236 bool = (2147483647 > 2)
    var v237 uint32 = (v3 % ((1) | 1))
    v238 := GenS18{F0: -1, F1: 9223372036854775807}
    var v239 int8 = 42
    var v240 uint16 = 0
    var v241 uint16 = v240
    for fi242 := 0; fi242 < 17; fi242++ {
        v243 := GenS14{F0: 128, F1: 9223372036854775807, F2: 1000}
        if false {
            var v244 float64 = ((-2.25 - -2.25) * (0.0 / 1e10))
            for fi245 := 0; fi245 < 17; fi245++ {
                v246 := GenS21{F0: -9223372036854775808}
                _ = fi245
                _ = v246
            }
            switch "" {
            case "":
                v247 := GenS11{F0: (^(v241))}
                _ = v247
            default:
                var v248 bool = v236
                _ = v248
            }
            _ = v244
        }
        _ = fi242
        _ = v243
    }
    v3 -= v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v236:", v236)
    fmt.Println("v237:", v237)
    fmt.Println("v239:", v239)
    fmt.Println("v240:", v240)
    fmt.Println("v241:", v241)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v236
    _ = v237
    _ = v238
    _ = v239
    _ = v240
    _ = v241
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
