package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int64; F2 int32 }
type GenS3 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64; F1 int32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int; F1 string; F2 int16 }
type GenS11 struct { F0 int8; F1 bool; F2 int8 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint32; F1 int32; F2 string }
type GenS14 struct { F0 int; F1 int }
type GenS15 struct { F0 int; F1 int16; F2 uint }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int; F1 uint32; F2 bool }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 bool; F1 uint8; F2 string }
func genFn0(p0 uint8, p1 bool) uint64 {
    var v7 uint = 0
    var v8 string = ""
    if (true || (42 > 42)) {
        var v9 uint32 = 720
        _ = v9
    } else {
        p1 = (32767 < -508)
        var v10 string = ""
        var v11 uint64 = 245
        _ = v10
        _ = v11
    }
    fmt.Println(p0, p1, p0)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int8 = -89
    v2 := []int{}
    var v3 string = "a"
    var v4 float64 = float64(1e10)
    v5 := make(map[int]int)
    v6 := make(map[uint32]string)
    v3 = ""
    var v12 uint = 255
    var v13 uint64 = genFn0(59, true)
    delete(v6, 0)
    v1 = v1
    v13 += v13
    v14 := GenS1{F0: -2147483648}
    if (65535 < 1) {
        if ((209 == 255) || false) {
            for fi15 := 0; fi15 < 9; fi15++ {
                fmt.Println(v13, v13, v12)
                var v16 int64 = 117
                v17 := GenS1{F0: fi15}
                _ = fi15
                _ = v16
                _ = v17
            }
            switch (("hello" + v3) + "hello") {
            case "\n\t":
                v18 := GenS2{F0: true, F1: 1, F2: 1}
                _ = v18
            case "":
                switch 1 {
                case -9223372036854775808:
                    var v19 float64 = v4
                    _ = v19
                case 1:
                    fmt.Println(v13)
                default:
                    switch v3 {
                    case "a":
                        switch v3 {
                        case "a":
                            for fi20 := 0; fi20 < 2; fi20++ {
                                fmt.Println(v3)
                                switch (v3 + "a") {
                                case "foo bar":
                                    for fi21 := 0; fi21 < 3; fi21++ {
                                        var v22 bool = ((v1 * (-(v1))) != -1)
                                        v23 := GenS3{F0: v1, F1: 222, F2: 65535}
                                        _ = fi21
                                        _ = v22
                                        _ = v23
                                    }
                                case "":
                                    var v24 bool = (1 >= 1)
                                    _ = v24
                                default:
                                    v4 = (v4 / float64(((float32(v4) - (1.5 * -2.25)) * (-2.25 * 0.1))))
                                }
                                v25 := GenS3{F0: v1, F1: 122, F2: 0}
                                _ = fi20
                                _ = v25
                            }
                        case "":
                            var v26 uint16 = 28
                            _ = v26
                        }
                    case "\n\t":
                        var v27 bool = (((2147483648 <= 44) || false) || false)
                        _ = v27
                    default:
                        for fi28 := 0; fi28 < 14; fi28++ {
                            if (false && true) {
                                if false {
                                    if false {
                                        for fi29 := 0; fi29 < 20; fi29++ {
                                            if true {
                                                if true {
                                                    var v30 int8 = (v1 / ((v1) | 1))
                                                    switch (fi28 * (^(fi28))) {
                                                    case 42:
                                                        switch (fi29 ^ fi28) {
                                                        case -2147483648:
                                                            var v31 uint8 = 2
                                                            _ = v31
                                                        case -554:
                                                            switch fi28 {
                                                            case 4294967295:
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            case -611:
                                                                if (fi29 < int(v13)) {
                                                                    v5[fi28] = (-(fi28))
                                                                    var v32 int64 = -66
                                                                    v2 = append(v2, (fi29 >> ((v12) & 63)))
                                                                    _ = v32
                                                                } else {
                                                                    v30 += (v30 % ((v30) | 1))
                                                                }
                                                                v13 = (v13 + 0)
                                                                if false {
                                                                    v12 -= (^(v12))
                                                                }
                                                            default:
                                                                var v33 int64 = -9223372036854775808
                                                                _ = v33
                                                            }
                                                        }
                                                    case 9223372036854775807:
                                                        fmt.Println(fi28, v1, v1)
                                                    default:
                                                        fmt.Println(fi29, fi29)
                                                        v34 := GenS1{F0: (fi29 * fi28)}
                                                        var v35 int16 = 42
                                                        switch v3 {
                                                        case "foo bar":
                                                            v36 := GenS3{F0: 24, F1: 1, F2: 0}
                                                            _ = v36
                                                        default:
                                                            switch (fi28 + fi28) {
                                                            case 128:
                                                                v3 = "a"
                                                            case 127:
                                                                var v37 int16 = 1
                                                                _ = v37
                                                            case 185:
                                                                if (!false) {
                                                                    v38 := GenS4{F0: (v30 >> ((v12) & 63))}
                                                                    _ = v38
                                                                } else {
                                                                    switch v3 {
                                                                    case "a":
                                                                        switch v3 {
                                                                        case "foo bar":
                                                                            var v39 int8 = (v1 + 100)
                                                                            _ = v39
                                                                        case "":
                                                                            switch (fi28 & (fi29 - fi29)) {
                                                                            case 69:
                                                                                v12 = 4294967295
                                                                            case -742:
                                                                                fmt.Println(v1, fi29)
                                                                            }
                                                                            fmt.Println(v30)
                                                                        case "a":
                                                                            v40 := GenS5{F0: 2147483648}
                                                                            _ = v40
                                                                        default:
                                                                            var v41 uint8 = 1
                                                                            _ = v41
                                                                        }
                                                                    default:
                                                                        var v42 int32 = -826
                                                                        _ = v42
                                                                    }
                                                                    switch v3 {
                                                                    case "\n\t":
                                                                        var v43 uint32 = 2147483648
                                                                        _ = v43
                                                                    }
                                                                    v44 := GenS2{F0: true, F1: -209, F2: 743}
                                                                    _ = v44
                                                                }
                                                            }
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            v30 += v30
                                                        }
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                    if false {
                                                        var v45 bool = ((^(v12)) > uint(v13))
                                                        v3 = ("\n\t" + "\n\t")
                                                        v4 = -2.25
                                                        _ = v45
                                                    }
                                                    _ = v30
                                                } else {
                                                    if ((!(!(24 == uint8(v13)))) && true) {
                                                        switch (fi29 % (((-(fi29))) | 1)) {
                                                        case 940:
                                                            fmt.Println(fi29, v1, fi29)
                                                        case 127:
                                                            var v46 uint64 = (v13 - (v13 + (^(v13))))
                                                            _ = v46
                                                        case -125:
                                                            if len(v2) > 0 {
                                                                v2[0] = (-(fi28))
                                                            }
                                                        }
                                                        fmt.Println(v1, v3, v3)
                                                        fmt.Println(v4, v13, fi29)
                                                    }
                                                    var v47 int64 = int64(fi29)
                                                    _ = v47
                                                }
                                                var v48 float64 = float64((-2.25 + ((((3.14159 / 0.1) + 1e10) * 1e10) - ((0.0 - ((0.1 - 0.0) + 1.5)) * 3.14159))))
                                                fmt.Println(v13, fi28, v12)
                                                _ = v48
                                            }
                                            _ = fi29
                                        }
                                    }
                                }
                            }
                            _ = fi28
                        }
                    }
                }
            default:
                switch ("a" + v3) {
                case "a":
                    v13 = 4294967295
                case "hello":
                    if true {
                        if (4294967295 == (^(v13))) {
                            v49 := v5[1]
                            if (((v13 / ((v13) | 1)) == 1) && false) {
                                for fi50 := 0; fi50 < 7; fi50++ {
                                    v51 := GenS6{F0: v1, F1: 0}
                                    var v52 int16 = 0
                                    _ = fi50
                                    _ = v51
                                    _ = v52
                                }
                                if len(v2) > 0 {
                                    v2[0] = v49
                                }
                            } else {
                                var v53 uint32 = 2147483647
                                _ = v53
                            }
                            v4 = float64((0.1 / 1e10))
                            _ = v49
                        }
                        v54 := GenS7{F0: (v13 * v13)}
                        if false {
                            var v55 uint64 = (v13 | 4294967295)
                            fmt.Println(v3)
                            _ = v55
                        } else {
                            switch "\n\t" {
                            case "a":
                                fmt.Println(v1, v12, v4)
                            }
                        }
                        _ = v54
                    } else {
                        var v56 int16 = 42
                        for fi57 := 0; fi57 < 5; fi57++ {
                            v1 = (^(v1))
                            v58 := GenS1{F0: fi57}
                            _ = fi57
                            _ = v58
                        }
                        _ = v56
                    }
                }
            }
            if (v13 > (v13 ^ v13)) {
                if true {
                    fmt.Println(v12, v12)
                    v2 = append(v2, -77)
                    var v59 int16 = 0
                    _ = v59
                } else {
                    for fi60 := 0; fi60 < 10; fi60++ {
                        _ = len(v6)
                        for fi61 := 0; fi61 < 7; fi61++ {
                            var v62 float32 = (0.0 + (0.0 + (float32(3.14159) * 1e10)))
                            var v63 bool = true
                            if v63 {
                                var v64 int = fi61
                                switch fi60 {
                                case -503:
                                    v65 := GenS1{F0: v64}
                                    _ = v65
                                case -690:
                                    v66 := GenS3{F0: 0, F1: 2, F2: 521}
                                    _ = v66
                                default:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                if v63 {
                                    if (v63 && (false && v63)) {
                                        var v67 int32 = -112
                                        v68 := GenS7{F0: genFn0(127, v63)}
                                        _ = v67
                                        _ = v68
                                    }
                                    v64 = (v64 % (((v64 * v64)) | 1))
                                    v63 = (!(v12 < 141))
                                    var v69 uint16 = 929
                                    var v70 bool = false
                                    var v71 int8 = v1
                                    _ = v69
                                    _ = v70
                                    _ = v71
                                }
                                _ = v64
                            } else {
                                v72 := v6[35]
                                _ = v72
                            }
                            if v63 {
                                if (!(128 == 2)) {
                                    var v73 uint = (^(v12))
                                    _ = v73
                                }
                                v4 += (0.0 - (float64(-2.25) + v4))
                            } else {
                                switch v3 {
                                case "":
                                    var v74 int32 = 1
                                    _ = v74
                                default:
                                    var v75 int32 = 42
                                    _ = v75
                                }
                                v13 -= 1
                            }
                            fmt.Println(v4, v62)
                            if len(v2) > 0 {
                                v2[0] = 99
                            }
                            if len(v2) > 0 {
                                v2[0] = (fi61 + fi61)
                            }
                            if true {
                                v76 := GenS8{F0: v13, F1: 0}
                                switch v3 {
                                case "a":
                                    if v63 {
                                        v4 -= v4
                                    } else {
                                        v1 = 42
                                    }
                                case "hello":
                                    v77 := GenS2{F0: false, F1: 262, F2: 2147483647}
                                    _ = v77
                                default:
                                    v63 = false
                                }
                                fmt.Println(v4, v12)
                                _ = v76
                            }
                            var v78 uint32 = 4294967295
                            if v63 {
                                var v79 uint = (v12 / ((v12) | 1))
                                v2 = append(v2, fi61)
                                if (-965 > -2147483648) {
                                    var v80 uint64 = (v13 % ((v13) | 1))
                                    if v63 {
                                        v81 := v5[fi60]
                                        switch "hello" {
                                        case "a":
                                            v82 := GenS2{F0: false, F1: 9223372036854775807, F2: -525}
                                            _ = v82
                                        case "foo bar":
                                            if (177 <= 128) {
                                                var v83 float32 = (((float32(3.14159) * float32(1.5)) - v62) + v62)
                                                if ((^(v12)) > uint(v78)) {
                                                    v78 -= (^(v78))
                                                    v83 = 0.0
                                                    var v84 int32 = 1
                                                    fmt.Println(v1, v13, v84)
                                                    _ = v84
                                                }
                                                var v85 uint64 = (v80 + 42)
                                                _ = v83
                                                _ = v85
                                            } else {
                                                fmt.Println(v79)
                                                if v63 {
                                                    delete(v6, (^(v78)))
                                                } else {
                                                    var v86 float32 = v62
                                                    _ = v86
                                                }
                                            }
                                        case "hello":
                                            var v87 int32 = 0
                                            _ = v87
                                        }
                                        _ = v81
                                    } else {
                                        switch "a" {
                                        case "a":
                                            if v63 {
                                                var v88 uint16 = 65535
                                                _ = len(v5)
                                                var v89 uint32 = (v78 | (v78 / (((v78 | 997)) | 1)))
                                                _ = v88
                                                _ = v89
                                            }
                                        case "hello":
                                            var v90 bool = true
                                            _ = v90
                                        case "foo bar":
                                            switch "" {
                                            case "a":
                                                var v91 int16 = -901
                                                _ = v91
                                            case "hello":
                                                if false {
                                                    v2 = append(v2, -488)
                                                }
                                            default:
                                                if v63 {
                                                    if (v12 == (v12 >> ((v79) & 63))) {
                                                        var v92 int8 = v1
                                                        fmt.Println(v1)
                                                        _ = v92
                                                    }
                                                    if ((3.14159 / v62) < (3.14159 + (v62 * v62))) {
                                                        v13 += (^(v13))
                                                        var v93 uint16 = uint16(v79)
                                                        var v94 int = (fi61 >> ((0) & 63))
                                                        _ = v93
                                                        _ = v94
                                                    }
                                                    switch "\n\t" {
                                                    case "":
                                                        var v95 uint64 = v80
                                                        if true {
                                                            v96 := GenS9{F0: 127}
                                                            _ = v96
                                                        }
                                                        _ = v95
                                                    default:
                                                        var v97 string = ("" + v3)
                                                        _ = v97
                                                    }
                                                } else {
                                                    var v98 uint = v12
                                                    if (-1 != 1000) {
                                                        _ = len(v2)
                                                    } else {
                                                        var v99 uint16 = 65535
                                                        var v100 uint64 = (^(v13))
                                                        if (!v63) {
                                                            if true {
                                                                v99 = (v99 + uint16(v98))
                                                                var v101 int64 = 0
                                                                v102 := v6[(v78 % ((0) | 1))]
                                                                if (!true) {
                                                                    v4 -= v4
                                                                    v103 := v5[(fi60 * (fi61 + int(v79)))]
                                                                    fmt.Println(v62, fi60, fi61)
                                                                    _ = v103
                                                                }
                                                                if (!v63) {
                                                                    v3 = (v102 + ((v3 + v102) + ""))
                                                                } else {
                                                                    if false {
                                                                        switch (((("" + (v3 + ((v102 + ("hello" + v102)) + v3))) + v3) + v3) + v3) {
                                                                        case "":
                                                                            if v63 {
                                                                                var v104 string = "foo bar"
                                                                                fmt.Println(v62)
                                                                                if true {
                                                                                    v105 := v5[int(v101)]
                                                                                    switch fi60 {
                                                                                    case 42:
                                                                                        fmt.Println(v105, v1, v3)
                                                                                    case -542:
                                                                                        if (1000 < 0) {
                                                                                            fmt.Println(v3, v101)
                                                                                            if (-1 >= 481) {
                                                                                                fmt.Println(v101, v105)
                                                                                                switch 9223372036854775807 {
                                                                                                case 42:
                                                                                                    if true {
                                                                                                        var v106 uint64 = (^(v80))
                                                                                                        _ = v106
                                                                                                    } else {
                                                                                                        if v63 {
                                                                                                            if v63 {
                                                                                                                var v107 uint64 = (v80 / ((v100) | 1))
                                                                                                                switch ((("\n\t" + v104) + "") + (v102 + "")) {
                                                                                                                case "hello":
                                                                                                                    var v108 int64 = v101
                                                                                                                    _ = v108
                                                                                                                case "":
                                                                                                                    v109 := GenS8{F0: genFn0(2, (!(float32((((float64(1.5) - v4) - 0.0) / 1e10)) < v62))), F1: -2147483648}
                                                                                                                    _ = v109
                                                                                                                }
                                                                                                                _ = v107
                                                                                                            } else {
                                                                                                                var v110 float32 = v62
                                                                                                                v80 = 666
                                                                                                                v105 -= fi61
                                                                                                                _ = v110
                                                                                                            }
                                                                                                        } else {
                                                                                                            v62 += v62
                                                                                                            if true {
                                                                                                                var v111 uint32 = v78
                                                                                                                v98 = v12
                                                                                                                v112 := GenS2{F0: (((v100 >> (((v12 ^ (^(v79)))) & 63)) != uint64(v99)) || false), F1: v101, F2: 1}
                                                                                                                _ = v111
                                                                                                                _ = v112
                                                                                                            } else {
                                                                                                                delete(v6, 273)
                                                                                                            }
                                                                                                        }
                                                                                                        v6[v78] = ("" + "a")
                                                                                                        v113 := GenS2{F0: (fi61 == int(v98)), F1: (v101 << (((v98 >> ((709) & 63))) & 63)), F2: -859}
                                                                                                        v12 += 1
                                                                                                        _ = v113
                                                                                                    }
                                                                                                case 1:
                                                                                                    var v114 string = ""
                                                                                                    _ = v114
                                                                                                case 4294967295:
                                                                                                    v1 -= (v1 % ((1) | 1))
                                                                                                }
                                                                                            } else {
                                                                                                if v63 {
                                                                                                    if v63 {
                                                                                                        var v115 uint16 = v99
                                                                                                        _ = v115
                                                                                                    }
                                                                                                    if v63 {
                                                                                                        v98 = uint(v13)
                                                                                                        switch fi60 {
                                                                                                        case 157:
                                                                                                            v116 := GenS8{F0: v80, F1: 1}
                                                                                                            _ = v116
                                                                                                        }
                                                                                                        var v117 uint = (v79 >> ((v98) & 63))
                                                                                                        fmt.Println(v98, v1)
                                                                                                        _ = v117
                                                                                                    } else {
                                                                                                        switch (fi60 / ((fi61) | 1)) {
                                                                                                        case 128:
                                                                                                            v5[(fi61 / (((^(fi60))) | 1))] = 352
                                                                                                        case 42:
                                                                                                            v101 = v101
                                                                                                        case 4294967295:
                                                                                                            v80 -= 407
                                                                                                        default:
                                                                                                            var v118 int8 = 105
                                                                                                            _ = v118
                                                                                                        }
                                                                                                    }
                                                                                                    var v119 int8 = 36
                                                                                                    _ = v119
                                                                                                } else {
                                                                                                    if v63 {
                                                                                                        v13 -= (v100 + 42)
                                                                                                    }
                                                                                                    v120 := GenS10{F0: int(v98), F1: v102, F2: 734}
                                                                                                    _ = v120
                                                                                                }
                                                                                                var v121 uint64 = uint64(v79)
                                                                                                v1 -= int8(v98)
                                                                                                _ = v121
                                                                                            }
                                                                                            v122 := GenS9{F0: 2}
                                                                                            _ = v122
                                                                                        } else {
                                                                                            v104 = (v102 + (v3 + v3))
                                                                                        }
                                                                                    }
                                                                                    _ = v105
                                                                                } else {
                                                                                    var v123 int16 = 643
                                                                                    v124 := GenS5{F0: (v78 + (v78 ^ v78))}
                                                                                    _ = v123
                                                                                    _ = v124
                                                                                }
                                                                                _ = v104
                                                                            } else {
                                                                                v101 += v101
                                                                                var v125 bool = true
                                                                                if false {
                                                                                    var v126 uint8 = 165
                                                                                    _ = v126
                                                                                }
                                                                                _ = v125
                                                                            }
                                                                        case "hello":
                                                                            v127 := GenS1{F0: -9223372036854775808}
                                                                            _ = v127
                                                                        case "\n\t":
                                                                            v79 -= 18446744073709551615
                                                                        default:
                                                                            var v128 float64 = v4
                                                                            _ = v128
                                                                        }
                                                                        v100 -= v100
                                                                    } else {
                                                                        v129 := GenS7{F0: uint64(v99)}
                                                                        var v130 int = (fi61 + (fi61 - fi60))
                                                                        _ = v129
                                                                        _ = v130
                                                                    }
                                                                    v131 := GenS5{F0: v78}
                                                                    v101 += (v101 * int64(v78))
                                                                    v101 -= v101
                                                                    _ = v131
                                                                }
                                                                _ = v101
                                                                _ = v102
                                                            } else {
                                                                if ("\n\t" <= "\n\t") {
                                                                    var v132 int8 = (-(v1))
                                                                    switch (v3 + v3) {
                                                                    case "a":
                                                                        v133 := GenS8{F0: (v80 + v80), F1: 42}
                                                                        _ = v133
                                                                    case "foo bar":
                                                                        v78 += 65535
                                                                    case "\n\t":
                                                                        var v134 float64 = 0.1
                                                                        _ = v134
                                                                    }
                                                                    v135 := GenS11{F0: (-(v132)), F1: (false && v63), F2: (^(v1))}
                                                                    _ = v132
                                                                    _ = v135
                                                                }
                                                                v2 = append(v2, (fi61 - fi60))
                                                                if (!(!(fi60 > (fi60 % ((fi61) | 1))))) {
                                                                    var v136 int8 = (v1 | v1)
                                                                    _ = v136
                                                                } else {
                                                                    v137 := GenS12{F0: -987}
                                                                    _ = v137
                                                                }
                                                                if ((float32(0.1) * v62) != v62) {
                                                                    v4 = (v4 * 1e10)
                                                                }
                                                                v4 += ((1.5 + v4) + v4)
                                                            }
                                                            var v138 int32 = -92
                                                            if false {
                                                                if len(v2) > 0 {
                                                                    v2[0] = fi61
                                                                }
                                                                v139 := GenS7{F0: (v13 / ((v13) | 1))}
                                                                _ = v139
                                                            } else {
                                                                var v140 uint16 = v99
                                                                var v141 uint32 = (v78 >> ((v79) & 63))
                                                                _ = v140
                                                                _ = v141
                                                            }
                                                            _ = v138
                                                        } else {
                                                            v79 -= (v79 | v79)
                                                        }
                                                        _ = v100
                                                        _ = v99
                                                    }
                                                    _ = len(v5)
                                                    _ = v98
                                                }
                                            }
                                        }
                                        fmt.Println(v79, v12)
                                        v142 := GenS13{F0: (v78 & 4294967295), F1: 2147483647, F2: "a"}
                                        _ = v142
                                    }
                                    _ = v80
                                }
                                _ = v79
                            }
                            v4 = 3.14159
                            _ = fi61
                            _ = v62
                            _ = v63
                            _ = v78
                        }
                        _ = fi60
                    }
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                for fi143 := 0; fi143 < 1; fi143++ {
                    var v144 bool = (!false)
                    fmt.Println(v13)
                    if (42 > 2147483647) {
                        fmt.Println(v12, fi143)
                        v145 := GenS14{F0: int(v1), F1: fi143}
                        _ = v145
                    }
                    _ = fi143
                    _ = v144
                }
            }
        } else {
            v4 -= 0.1
            v146 := GenS9{F0: 128}
            _ = v146
        }
        var v147 string = ((v3 + "") + ("foo bar" + v3))
        v148 := GenS2{F0: false, F1: 388, F2: 60}
        _ = v147
        _ = v148
    } else {
        var v149 float32 = 1e10
        if false {
            if ((686 != 1000) || (((false || false) && (((v149 / (1.5 * 1.5)) * float32((0.0 - (((float64((((float32(float64((v149 / float32(v4)))) / v149) - ((v149 - float32(((0.1 * 1e10) - v4))) * (float32(v4) * (0.1 - float32(v4))))) + -2.25)) - (((3.14159 / ((v4 - v4) / (v4 - 0.1))) * ((0.0 / v4) - ((v4 / v4) + v4))) - ((0.1 * ((v4 + (1.5 - 0.1)) / v4)) / v4))) - 0.1) * float64((v149 - 0.1)))))) >= (0.0 - v149))) && false)) {
                v149 = 3.14159
                if (true || (false || true)) {
                    for fi150 := 0; fi150 < 10; fi150++ {
                        fmt.Println(v12, v149, v3)
                        for fi151 := 0; fi151 < 8; fi151++ {
                            var v152 uint8 = 0
                            _ = fi151
                            _ = v152
                        }
                        var v153 int = 42
                        _ = fi150
                        _ = v153
                    }
                    switch v3 {
                    case "hello":
                        var v154 string = "\n\t"
                        _ = v154
                    case "a":
                        var v155 uint = v12
                        _ = v155
                    default:
                        var v156 int = -233
                        _ = v156
                    }
                    v149 += v149
                } else {
                    if false {
                        var v157 float32 = ((v149 * (v149 / 1.0)) + (0.0 * 0.0))
                        for fi158 := 0; fi158 < 12; fi158++ {
                            v12 = (v12 << ((2) & 63))
                            _ = fi158
                        }
                        if (1.5 <= (v149 / float32((((v4 * (0.1 - -2.25)) + float64(1e10)) / 1.5)))) {
                            v3 = v3
                            var v159 uint8 = 1
                            fmt.Println(v1, v1)
                            _ = v159
                        }
                        _ = v157
                    } else {
                        for fi160 := 0; fi160 < 11; fi160++ {
                            v161 := GenS1{F0: fi160}
                            for fi162 := 0; fi162 < 5; fi162++ {
                                if (false || (!(!true))) {
                                    if false {
                                        v163 := v6[65535]
                                        switch v163 {
                                        case "a":
                                            v1 += v1
                                        case "hello":
                                            var v164 uint = (v12 & (v12 * (v12 | v12)))
                                            _ = v164
                                        case "":
                                            if false {
                                                var v165 int16 = 639
                                                v166 := GenS6{F0: (v1 * v1), F1: 191}
                                                _ = len(v2)
                                                _ = v165
                                                _ = v166
                                            }
                                        }
                                        v167 := GenS15{F0: (fi160 % (((fi162 + 128)) | 1)), F1: 1, F2: 0}
                                        _ = v163
                                        _ = v167
                                    }
                                }
                                if ((v4 < float64(float32(1e10))) && true) {
                                    var v168 bool = ((v3 + "foo bar") > ("hello" + "\n\t"))
                                    if (v168 || (v4 <= v4)) {
                                        var v169 float64 = v4
                                        fmt.Println(v3)
                                        _ = v169
                                    } else {
                                        v168 = (1 >= 0)
                                        var v170 uint = v12
                                        var v171 int = -2147483648
                                        if len(v2) > 0 {
                                            v2[0] = fi162
                                        }
                                        _ = v170
                                        _ = v171
                                    }
                                    _ = v168
                                } else {
                                    v172 := GenS8{F0: (v13 ^ genFn0(127, (true || false))), F1: int32(v13)}
                                    var v173 int8 = -1
                                    _ = v172
                                    _ = v173
                                }
                                _ = fi162
                            }
                            _ = fi160
                            _ = v161
                        }
                        switch ("a" + "hello") {
                        case "a":
                            v174 := GenS9{F0: 2}
                            _ = v174
                        default:
                            var v175 uint8 = 127
                            _ = v175
                        }
                        var v176 int16 = 1
                        _ = v176
                    }
                    if true {
                        if true {
                            if len(v2) > 0 {
                                v2[0] = 42
                            }
                            for fi177 := 0; fi177 < 16; fi177++ {
                                fmt.Println(v13, fi177)
                                var v178 int16 = -597
                                v1 -= 37
                                _ = fi177
                                _ = v178
                            }
                            if false {
                                v179 := GenS16{F0: (!(true || false))}
                                _ = v179
                            }
                        } else {
                            if true {
                                if false {
                                    var v180 int8 = (-(v1))
                                    _ = v180
                                }
                                for fi181 := 0; fi181 < 14; fi181++ {
                                    v12 += 777
                                    var v182 string = (v3 + (v3 + v3))
                                    var v183 int64 = -277
                                    _ = fi181
                                    _ = v182
                                    _ = v183
                                }
                            }
                        }
                        var v184 uint32 = 65535
                        for fi185 := 0; fi185 < 3; fi185++ {
                            delete(v6, v184)
                            var v186 uint32 = v184
                            for fi187 := 0; fi187 < 4; fi187++ {
                                v12 = v12
                                if (32768 >= (fi185 % ((-248) | 1))) {
                                    v186 += v186
                                    v186 = (v184 & v184)
                                }
                                _ = fi187
                            }
                            _ = fi185
                            _ = v186
                        }
                        _ = v184
                    } else {
                        if (1 < 255) {
                            v149 -= v149
                            v188 := GenS6{F0: (v1 % ((v1) | 1)), F1: 0}
                            _ = v188
                        }
                        v1 = v1
                        v189 := GenS17{F0: 0, F1: 65535, F2: (9223372036854775807 != 32767)}
                        _ = v189
                    }
                    v190 := GenS11{F0: 1, F1: ((255 == 32767) || (0 != 65)), F2: v1}
                    _ = v190
                }
                for fi191 := 0; fi191 < 18; fi191++ {
                    v149 -= v149
                    switch (v3 + v3) {
                    case "hello":
                        v149 = 0.0
                    case "\n\t":
                        var v192 string = (v3 + v3)
                        _ = v192
                    case "foo bar":
                        fmt.Println(v149, v149, v149)
                    }
                    _ = fi191
                }
            } else {
                for fi193 := 0; fi193 < 11; fi193++ {
                    var v194 int64 = 0
                    v5[fi193] = (fi193 & (^(fi193)))
                    v195 := GenS18{F0: (false || false)}
                    _ = fi193
                    _ = v194
                    _ = v195
                }
            }
            fmt.Println(v3, v149, v3)
            v2 = append(v2, -1)
        } else {
            v149 -= v149
            switch v3 {
            case "foo bar":
                if true {
                    var v196 uint = 2
                    _ = v196
                } else {
                    if false {
                        switch "foo bar" {
                        case "foo bar":
                            var v197 uint8 = 157
                            _ = v197
                        default:
                            v13 += (v13 + 42)
                        }
                    }
                }
            case "a":
                v1 = v1
            case "":
                var v198 int8 = v1
                _ = v198
            }
            v199 := GenS19{F0: (!false), F1: 1, F2: "hello"}
            _ = v199
        }
        fmt.Println(v12, v1, v1)
        _ = v149
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v200 uint = v12
    for fi201 := 0; fi201 < 11; fi201++ {
        _ = len(v2)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if false {
            var v202 bool = false
            _ = v202
        } else {
            v4 = v4
            var v203 int = fi201
            _ = v203
        }
        _ = fi201
    }
    v3 = v3
    v204 := GenS15{F0: 127, F1: int16(v13), F2: 18446744073709551615}
    var v205 string = ("a" + ((v3 + ((v3 + (("hello" + v3) + v3)) + (("foo bar" + "\n\t") + (v3 + "foo bar")))) + "\n\t"))
    v200 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v200:", v200)
    fmt.Println("v205:", v205)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v200
    _ = v204
    _ = v205
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
