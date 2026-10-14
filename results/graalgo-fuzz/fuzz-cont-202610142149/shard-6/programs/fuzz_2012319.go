package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 int32; F2 bool }
type GenS7 struct { F0 int64; F1 int8; F2 int16 }
type GenS8 struct { F0 bool; F1 int }
type GenS9 struct { F0 uint32; F1 uint }
type GenS10 struct { F0 uint; F1 int32 }
type GenS11 struct { F0 bool; F1 int16; F2 int64 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int; F1 uint32 }
type GenS15 struct { F0 uint32; F1 uint }
type GenS16 struct { F0 uint; F1 uint64 }
type GenS17 struct { F0 int32; F1 uint16 }
type GenS18 struct { F0 uint8; F1 int32 }
type GenS19 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS20 struct { F0 uint8; F1 string; F2 int32 }
type GenS21 struct { F0 uint; F1 bool }
type GenS22 struct { F0 uint64 }
type GenS23 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS24 struct { F0 bool; F1 int64 }
type GenS25 struct { F0 uint16; F1 int; F2 int64 }
type GenS26 struct { F0 uint8 }
type GenS27 struct { F0 string; F1 int8; F2 string }
type GenS28 struct { F0 bool; F1 int8; F2 uint64 }
func genFn0(p0 string) int8 {
    p0 = ("a" + p0)
    var v7 uint32 = 2147483647
    var v8 int32 = 0
    fmt.Println(v8, v7)
    fmt.Println(v7, v7, v8)
    _ = p0
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int = -354
    var v2 float32 = 1.5
    var v3 int = v1
    var v4 uint32 = 2147483648
    var v5 int8 = 127
    v6 := make(map[int]uint32)
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 2; fi10++ {
            if (true || (32767 == -836)) {
                v4 += v4
                var v11 int8 = 1
                v12 := []int{}
                _ = v11
                _ = v12
            } else {
                v13 := GenS1{F0: 171, F1: 200, F2: 42}
                switch "\n\t" {
                case "\n\t":
                    var v14 uint64 = 42
                    _ = v14
                case "":
                    if false {
                        var v15 uint32 = v4
                        _ = v15
                    } else {
                        v5 = (v5 % (((v5 % ((genFn0("a")) | 1))) | 1))
                        v16 := GenS1{F0: 1, F1: 16, F2: (v4 >> ((417) & 63))}
                        v1 += fi9
                        _ = v16
                    }
                default:
                    v5 += genFn0(("" + "foo bar"))
                }
                fmt.Println(v4)
                _ = v13
            }
            v17 := []int{376, (^(fi10)), int(v4)}
            v3 += (v3 % ((233) | 1))
            var v18 int8 = genFn0((((("foo bar" + ("foo bar" + "a")) + "hello") + "hello") + "\n\t"))
            _ = fi10
            _ = v17
            _ = v18
        }
        _ = fi9
    }
    delete(v6, v3)
    var v19 int8 = genFn0((("foo bar" + ("" + "")) + ("a" + "a")))
    v19 = (v5 - -128)
    for fi20 := 0; fi20 < 9; fi20++ {
        if (true && (true && false)) {
            if true {
                if true {
                    switch "\n\t" {
                    case "":
                        for fi21 := 0; fi21 < 19; fi21++ {
                            var v22 float64 = float64((-2.25 * 1.5))
                            var v23 uint32 = v4
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                    case "\n\t":
                        var v24 int8 = v19
                        _ = v24
                    }
                    fmt.Println(fi20, v5)
                    var v25 uint64 = 4294967295
                    _ = v25
                } else {
                    v26 := GenS2{F0: (^(v4))}
                    v5 += (v5 ^ v19)
                    _ = v26
                }
                for fi27 := 0; fi27 < 8; fi27++ {
                    v28 := []int{4294967295, (v1 | fi20)}
                    if len(v28) >= 2 {
                        v28 = v28[1:len(v28)]
                    }
                    switch fi27 {
                    case -9223372036854775808:
                        var v29 int64 = 1
                        _ = v29
                    case 42:
                        v6[v3] = (v4 * (v4 * v4))
                    default:
                        v2 = 1.5
                    }
                    var v30 bool = false
                    v28 = append(v28, int(v19))
                    _ = fi27
                    _ = v28
                    _ = v30
                }
                var v31 int = fi20
                _ = v31
            } else {
                var v32 float64 = 0.0
                v3 -= (v1 * fi20)
                _ = v32
            }
        } else {
            if true {
                if true {
                    for fi33 := 0; fi33 < 4; fi33++ {
                        switch "foo bar" {
                        case "\n\t":
                            fmt.Println(v2, v2)
                        case "":
                            switch -1 {
                            case 32767:
                                if false {
                                    v34 := GenS2{F0: v4}
                                    fmt.Println(v5, v1, fi20)
                                    _ = v34
                                }
                            case 4294967295:
                                var v35 uint64 = 338
                                _ = v35
                            default:
                                v2 -= v2
                            }
                            var v36 int = fi20
                            _ = v36
                        default:
                            var v37 int = (^(fi33))
                            _ = v37
                        }
                        _ = fi33
                    }
                }
                v38 := GenS3{F0: -69, F1: false}
                v39 := v6[(^(v1))]
                _ = v38
                _ = v39
            }
            if true {
                for fi40 := 0; fi40 < 19; fi40++ {
                    var v41 int8 = 127
                    var v42 string = "foo bar"
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                v3 = (v1 >> ((4294967295) & 63))
            } else {
                var v43 string = "foo bar"
                for fi44 := 0; fi44 < 11; fi44++ {
                    v45 := v6[fi44]
                    fmt.Println(v2, v43)
                    switch ((v43 + "hello") + ("" + v43)) {
                    case "\n\t":
                        fmt.Println(v45, v45)
                    }
                    if ((^(v5)) < v5) {
                        v3 += 1
                        v46 := GenS4{F0: 831}
                        _ = v46
                    } else {
                        var v47 uint8 = 80
                        v48 := v6[(v3 % ((9223372036854775807) | 1))]
                        _ = v47
                        _ = v48
                    }
                    var v49 int64 = -56
                    _ = fi44
                    _ = v45
                    _ = v49
                }
                _ = v43
            }
            v50 := v6[0]
            _ = v50
        }
        for fi51 := 0; fi51 < 19; fi51++ {
            switch ("foo bar" + "hello") {
            case "":
                var v52 uint = 2
                _ = v52
            default:
                fmt.Println(v1, v3)
            }
            _ = fi51
        }
        _ = fi20
    }
    var v53 int = v3
    if ((127 != 43) || false) {
        v19 = (v5 ^ -123)
        v54 := []int{(v3 >> ((30) & 63)), v3, (v53 << ((256) & 63)), (v3 & v1)}
        if (float64(float32((-2.25 / 1e10))) >= 0.1) {
            for fi55 := 0; fi55 < 8; fi55++ {
                v56 := GenS1{F0: 255, F1: 222, F2: (v4 << ((2) & 63))}
                var v57 int64 = 1
                _ = fi55
                _ = v56
                _ = v57
            }
        } else {
            fmt.Println(v1, v19)
            if false {
                if false {
                    v58 := GenS4{F0: 42}
                    _ = v58
                } else {
                    v59 := GenS1{F0: 42, F1: 42, F2: v4}
                    if true {
                        for fi60 := 0; fi60 < 20; fi60++ {
                            v61 := GenS5{F0: -360}
                            var v62 int16 = -869
                            _ = fi60
                            _ = v61
                            _ = v62
                        }
                    }
                    _ = v59
                }
                v5 += 0
                switch "\n\t" {
                case "":
                    if false {
                        for fi63 := 0; fi63 < 9; fi63++ {
                            var v64 uint8 = 2
                            for fi65 := 0; fi65 < 20; fi65++ {
                                switch "hello" {
                                case "":
                                    switch (^(v1)) {
                                    case 2147483647:
                                        fmt.Println(v19, v3, v64)
                                    case 9223372036854775807:
                                        switch (fi63 + (fi65 | (v3 % ((fi65) | 1)))) {
                                        case 680:
                                            var v66 int64 = -889
                                            _ = v66
                                        case 1:
                                            if false {
                                                v67 := GenS2{F0: 2147483647}
                                                var v68 string = "a"
                                                v1 = fi63
                                                _ = v67
                                                _ = v68
                                            }
                                        }
                                    default:
                                        var v69 float64 = ((float64(v2) / 3.14159) + 0.0)
                                        if (!true) {
                                            if false {
                                                if true {
                                                    if (((v2 + (v2 / v2)) / v2) != v2) {
                                                        _ = len(v6)
                                                        _ = len(v54)
                                                    }
                                                    v70 := GenS3{F0: -727, F1: false}
                                                    var v71 int = fi65
                                                    _ = v70
                                                    _ = v71
                                                }
                                                var v72 int8 = v19
                                                _ = v72
                                            } else {
                                                v5 = int8(v64)
                                                v1 = 4294967295
                                                if true {
                                                    if (((4294967295 != (fi65 | fi65)) && true) && true) {
                                                        var v73 int8 = (v19 ^ 127)
                                                        var v74 float64 = 0.0
                                                        v5 -= v5
                                                        _ = v73
                                                        _ = v74
                                                    } else {
                                                        v53 += (v53 << ((1) & 63))
                                                    }
                                                    var v75 uint16 = 2
                                                    if ("foo bar" <= "") {
                                                        v76 := GenS3{F0: 42, F1: (v64 != v64)}
                                                        if (32767 != -1) {
                                                            v77 := GenS3{F0: -1, F1: false}
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                var v78 float32 = v2
                                                                _ = v78
                                                            default:
                                                                var v79 uint = 255
                                                                _ = v79
                                                            }
                                                            _ = v77
                                                        } else {
                                                            if false {
                                                                if false {
                                                                    fmt.Println(v2)
                                                                    var v80 uint8 = v64
                                                                    _ = v80
                                                                } else {
                                                                    v81 := GenS6{F0: true, F1: -21, F2: ((1e10 <= v69) || (v4 <= v4))}
                                                                    var v82 int16 = 1
                                                                    v83 := GenS5{F0: -470}
                                                                    _ = v81
                                                                    _ = v82
                                                                    _ = v83
                                                                }
                                                                fmt.Println(v64, v2)
                                                            }
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                var v84 int16 = -199
                                                                _ = v84
                                                            case "hello":
                                                                fmt.Println(v4, v75, v75)
                                                            default:
                                                                switch (fi63 ^ 2147483647) {
                                                                case 32768:
                                                                    var v85 uint64 = 4294967295
                                                                    _ = v85
                                                                case 1:
                                                                    v64 += v64
                                                                default:
                                                                    if (!true) {
                                                                        if (!false) {
                                                                            v53 = v53
                                                                        } else {
                                                                            switch "foo bar" {
                                                                            case "a":
                                                                                v2 -= (-2.25 - float32(float64(1.5)))
                                                                            case "\n\t":
                                                                                var v86 int = v1
                                                                                _ = v86
                                                                            case "hello":
                                                                                var v87 uint16 = v75
                                                                                _ = v87
                                                                            default:
                                                                                v2 -= 0.1
                                                                            }
                                                                            v88 := GenS7{F0: int64(v5), F1: genFn0((("hello" + "hello") + ("" + "a"))), F2: -816}
                                                                            _ = v88
                                                                        }
                                                                        fmt.Println(v2)
                                                                        var v89 int8 = 100
                                                                        _ = v89
                                                                    } else {
                                                                        var v90 int16 = 982
                                                                        if (!true) {
                                                                            if (38 <= (v19 % ((127) | 1))) {
                                                                                var v91 uint32 = v4
                                                                                v92 := v6[fi63]
                                                                                _ = v91
                                                                                _ = v92
                                                                            }
                                                                            switch fi63 {
                                                                            case -2147483648:
                                                                                if true {
                                                                                    v93 := GenS1{F0: 702, F1: v64, F2: v4}
                                                                                    _ = v93
                                                                                } else {
                                                                                    var v94 int = (^(v3))
                                                                                    v95 := GenS4{F0: v75}
                                                                                    _ = v94
                                                                                    _ = v95
                                                                                }
                                                                            }
                                                                            v19 = v5
                                                                        }
                                                                        var v96 uint32 = (v4 - (v4 % ((v4) | 1)))
                                                                        v4 -= (v4 >> ((42) & 63))
                                                                        fmt.Println(v5)
                                                                        switch (v53 >> ((225) & 63)) {
                                                                        case 4294967295:
                                                                            v6[(v53 % ((fi63) | 1))] = v96
                                                                        case 32767:
                                                                            fmt.Println(v4)
                                                                        default:
                                                                            var v97 int32 = -950
                                                                            _ = v97
                                                                        }
                                                                        _ = v90
                                                                        _ = v96
                                                                    }
                                                                }
                                                            }
                                                            var v98 int64 = 1
                                                            _ = v98
                                                        }
                                                        v53 = 128
                                                        _ = v76
                                                    } else {
                                                        var v99 uint64 = 1
                                                        if false {
                                                            var v100 uint = 42
                                                            if false {
                                                                var v101 string = ""
                                                                v100 = v100
                                                                _ = v101
                                                            }
                                                            var v102 uint = v100
                                                            _ = v100
                                                            _ = v102
                                                        }
                                                        v3 -= 32767
                                                        _ = v99
                                                    }
                                                    _ = v75
                                                } else {
                                                    v3 -= -112
                                                    switch fi63 {
                                                    case 0:
                                                        if (fi65 > (^(v1))) {
                                                            if (!(2147483647 >= -186)) {
                                                                switch "foo bar" {
                                                                case "":
                                                                    var v103 bool = (false || true)
                                                                    _ = v103
                                                                case "a":
                                                                    switch ("foo bar" + "foo bar") {
                                                                    case "\n\t":
                                                                        switch v1 {
                                                                        case 4294967295:
                                                                            switch (fi65 % ((fi63) | 1)) {
                                                                            case 1:
                                                                                v1 -= (fi65 - v53)
                                                                            case -246:
                                                                                var v104 uint = 943
                                                                                _ = v104
                                                                            default:
                                                                                if true {
                                                                                    var v105 uint = 18446744073709551615
                                                                                    _ = v105
                                                                                }
                                                                            }
                                                                        case 128:
                                                                            v106 := GenS1{F0: 42, F1: (v64 << ((0) & 63)), F2: 2}
                                                                            _ = v106
                                                                        default:
                                                                            if (1000 > 0) {
                                                                                v53 = v53
                                                                                v54 = append(v54, (fi63 ^ v53))
                                                                            } else {
                                                                                delete(v6, 4294967295)
                                                                                var v107 string = "hello"
                                                                                if true {
                                                                                    v1 = -1
                                                                                    v108 := v6[v53]
                                                                                    _ = v108
                                                                                } else {
                                                                                    v64 = 51
                                                                                    if true {
                                                                                        if (4294967295 != 18446744073709551615) {
                                                                                            v109 := GenS8{F0: (!(false && ((fi63 ^ (fi63 << ((42) & 63))) != fi65))), F1: (^(v1))}
                                                                                            _ = len(v6)
                                                                                            var v110 uint8 = 55
                                                                                            v111 := GenS7{F0: -1, F1: v19, F2: -1}
                                                                                            var v112 uint = 1
                                                                                            _ = v109
                                                                                            _ = v110
                                                                                            _ = v111
                                                                                            _ = v112
                                                                                        } else {
                                                                                            switch (-(v53)) {
                                                                                            case 2147483647:
                                                                                                fmt.Println(v5, v1)
                                                                                            case 32767:
                                                                                                v113 := GenS5{F0: -1}
                                                                                                _ = v113
                                                                                            }
                                                                                            fmt.Println(v2, v107, v107)
                                                                                            var v114 string = (v107 + ("\n\t" + ((v107 + v107) + v107)))
                                                                                            switch v114 {
                                                                                            case "":
                                                                                                var v115 uint16 = 2
                                                                                                var v116 int32 = 1
                                                                                                _ = v115
                                                                                                _ = v116
                                                                                            case "\n\t":
                                                                                                v117 := GenS9{F0: (v4 % (((v4 ^ v4)) | 1)), F1: 595}
                                                                                                var v118 uint8 = 0
                                                                                                _ = len(v6)
                                                                                                _ = v117
                                                                                                _ = v118
                                                                                            }
                                                                                            v119 := GenS1{F0: 1, F1: v64, F2: 2}
                                                                                            _ = v114
                                                                                            _ = v119
                                                                                        }
                                                                                        v3 = v3
                                                                                    } else {
                                                                                        if (4294967295 > 829) {
                                                                                            fmt.Println(v3)
                                                                                            var v120 int16 = 1
                                                                                            _ = v120
                                                                                        } else {
                                                                                            v5 -= v19
                                                                                        }
                                                                                        v2 -= ((1.5 * ((v2 / v2) + float32(((1.5 / float64((3.14159 - v2))) - float64(float32((float64(v2) / (v69 + 1.5)))))))) / (0.1 + (v2 + v2)))
                                                                                    }
                                                                                }
                                                                                v121 := GenS4{F0: 32767}
                                                                                _ = v107
                                                                                _ = v121
                                                                            }
                                                                        }
                                                                    case "":
                                                                        fmt.Println(v5)
                                                                    case "hello":
                                                                        var v122 float32 = -2.25
                                                                        _ = v122
                                                                    }
                                                                case "foo bar":
                                                                    v123 := v6[(fi63 + fi65)]
                                                                    _ = v123
                                                                default:
                                                                    fmt.Println(fi63)
                                                                }
                                                            } else {
                                                                v3 = v1
                                                                v19 = genFn0("hello")
                                                                var v124 uint8 = (v64 | 155)
                                                                _ = v124
                                                            }
                                                            v2 = 1.5
                                                            switch ("hello" + "hello") {
                                                            case "foo bar":
                                                                _ = len(v6)
                                                            case "\n\t":
                                                                v125 := GenS10{F0: 18446744073709551615, F1: 847}
                                                                var v126 bool = true
                                                                _ = v125
                                                                _ = v126
                                                            case "a":
                                                                fmt.Println(v1, fi65)
                                                            }
                                                        } else {
                                                            if (true || (523 <= 255)) {
                                                                v127 := GenS2{F0: (^(v4))}
                                                                _ = v127
                                                            } else {
                                                                v53 -= v1
                                                            }
                                                        }
                                                        if ((("" + "") + "") != "\n\t") {
                                                            var v128 bool = true
                                                            v3 = v1
                                                            _ = v128
                                                        }
                                                    case 32768:
                                                        v3 = v53
                                                    case 2147483647:
                                                        v129 := GenS1{F0: 255, F1: (v64 % (((v64 ^ 128)) | 1)), F2: v4}
                                                        _ = v129
                                                    default:
                                                        v2 = (float32(0.1) + (v2 - ((v2 * (v2 * (0.1 + (v2 * float32(1e10))))) - 0.0)))
                                                    }
                                                    v130 := v6[(v53 & -770)]
                                                    _ = v130
                                                }
                                            }
                                            var v131 int16 = 32767
                                            _ = v131
                                        }
                                        v132 := GenS1{F0: uint(v64), F1: (^(v64)), F2: 4294967295}
                                        _ = v132
                                        _ = v69
                                    }
                                case "a":
                                    switch (("hello" + "hello") + "") {
                                    case "hello":
                                        v6[-606] = 65535
                                    case "":
                                        var v133 uint = 489
                                        _ = v133
                                    default:
                                        v5 = v5
                                    }
                                }
                                _ = fi65
                            }
                            for fi134 := 0; fi134 < 10; fi134++ {
                                if (((true || (2 <= 18446744073709551615)) || true) && (v4 >= (^(v4)))) {
                                    v1 -= 544
                                    if true {
                                        if true {
                                            if (255 <= 230) {
                                                switch ("a" + ("\n\t" + "")) {
                                                case "\n\t":
                                                    fmt.Println(fi134)
                                                case "":
                                                    if (!false) {
                                                        var v135 uint64 = uint64(v64)
                                                        v136 := GenS9{F0: v4, F1: 42}
                                                        v4 -= v4
                                                        _ = v135
                                                        _ = v136
                                                    }
                                                case "foo bar":
                                                    switch v1 {
                                                    case 42:
                                                        if true {
                                                            v2 = v2
                                                        } else {
                                                            v19 -= v5
                                                            v19 = v19
                                                        }
                                                    case 9223372036854775807:
                                                        var v137 uint16 = 1
                                                        _ = v137
                                                    default:
                                                        switch (v3 >> ((740) & 63)) {
                                                        case 42:
                                                            switch (v3 | (v3 & 844)) {
                                                            case 1:
                                                                var v138 int8 = v19
                                                                _ = v138
                                                            case 128:
                                                                if (((-2147483648 > 732) && false) && true) {
                                                                    v1 = 1
                                                                    fmt.Println(v53, v1, v53)
                                                                    var v139 uint32 = 42
                                                                    var v140 uint = 2
                                                                    _ = v139
                                                                    _ = v140
                                                                } else {
                                                                    if false {
                                                                        fmt.Println(fi63, v19)
                                                                        v6[(^(fi63))] = v4
                                                                    }
                                                                }
                                                            case 32767:
                                                                var v141 int32 = 42
                                                                _ = v141
                                                            default:
                                                                v4 -= (v4 * v4)
                                                            }
                                                        default:
                                                            var v142 uint32 = (v4 | (v4 << ((uint(v64)) & 63)))
                                                            _ = v142
                                                        }
                                                    }
                                                }
                                                v5 += v19
                                            }
                                            var v143 int = (v3 << ((927) & 63))
                                            v64 += (^(v64))
                                            fmt.Println(v4, fi134, v4)
                                            var v144 int16 = 1000
                                            _ = v143
                                            _ = v144
                                        } else {
                                            v53 = 1
                                            var v145 int32 = 425
                                            _ = v145
                                        }
                                        switch "foo bar" {
                                        case "":
                                            var v146 string = "hello"
                                            _ = v146
                                        case "hello":
                                            fmt.Println(v64)
                                            if (((float32(-2.25) - (((-2.25 - (v2 / v2)) * 0.0) - v2)) <= v2) || true) {
                                                var v147 int = -2147483648
                                                _ = len(v6)
                                                v148 := GenS9{F0: v4, F1: 18446744073709551615}
                                                v149 := GenS7{F0: 870, F1: 100, F2: 32767}
                                                var v150 uint8 = v64
                                                _ = len(v54)
                                                fmt.Println(v147, v64, fi63)
                                                _ = v147
                                                _ = v148
                                                _ = v149
                                                _ = v150
                                            } else {
                                                if ((v2 / 0.1) <= 3.14159) {
                                                    var v151 uint32 = v4
                                                    v152 := GenS9{F0: (v4 << ((42) & 63)), F1: 18446744073709551615}
                                                    _ = v151
                                                    _ = v152
                                                } else {
                                                    v153 := GenS8{F0: (1 > 1), F1: (v3 << ((0) & 63))}
                                                    var v154 string = "a"
                                                    switch ((v154 + v154) + ("a" + ("" + v154))) {
                                                    case "\n\t":
                                                        switch v1 {
                                                        case 785:
                                                            if (!true) {
                                                                var v155 uint8 = (v64 * v64)
                                                                _ = len(v54)
                                                                v156 := GenS11{F0: (0 == 42), F1: 0, F2: 0}
                                                                _ = v155
                                                                _ = v156
                                                            }
                                                        case -10:
                                                            v157 := GenS4{F0: 255}
                                                            _ = v157
                                                        default:
                                                            if (false || false) {
                                                                if (475 < uint64(v64)) {
                                                                    v158 := GenS1{F0: 2, F1: (^(v64)), F2: v4}
                                                                    _ = v158
                                                                }
                                                                switch v154 {
                                                                case "\n\t":
                                                                    switch 32767 {
                                                                    case -886:
                                                                        var v159 int16 = 1000
                                                                        _ = v159
                                                                    default:
                                                                        v160 := GenS12{F0: 1}
                                                                        _ = v160
                                                                    }
                                                                case "foo bar":
                                                                    v161 := GenS6{F0: (v1 <= fi134), F1: 6, F2: true}
                                                                    _ = v161
                                                                default:
                                                                    if true {
                                                                        switch v3 {
                                                                        case -760:
                                                                            v5 -= 0
                                                                        case -558:
                                                                            fmt.Println(fi63, v4)
                                                                            v19 = (-(v5))
                                                                            var v162 float64 = 0.1
                                                                            v163 := GenS13{F0: (v19 >> ((255) & 63))}
                                                                            _ = v162
                                                                            _ = v163
                                                                        }
                                                                        v1 = (^(fi134))
                                                                        v4 = 2
                                                                    } else {
                                                                        v54 = append(v54, fi63)
                                                                        if (fi63 == 9223372036854775807) {
                                                                            v164 := GenS14{F0: v53, F1: v4}
                                                                            fmt.Println(v1, v3, v64)
                                                                            v165 := GenS15{F0: (v4 ^ 42), F1: 563}
                                                                            _ = v164
                                                                            _ = v165
                                                                        } else {
                                                                            delete(v6, -2147483648)
                                                                            if true {
                                                                                if false {
                                                                                    switch "" {
                                                                                    case "\n\t":
                                                                                        fmt.Println(v19, v2)
                                                                                    case "foo bar":
                                                                                        switch -494 {
                                                                                        case 218:
                                                                                            var v166 uint32 = uint32(v64)
                                                                                            if true {
                                                                                                if true {
                                                                                                    var v167 uint8 = v64
                                                                                                    fmt.Println(v53, v3, fi63)
                                                                                                    fmt.Println(v2, v4, fi63)
                                                                                                    _ = v167
                                                                                                } else {
                                                                                                    var v168 uint64 = 974
                                                                                                    var v169 uint = 42
                                                                                                    var v170 int32 = -2147483648
                                                                                                    _ = v168
                                                                                                    _ = v169
                                                                                                    _ = v170
                                                                                                }
                                                                                            }
                                                                                            _ = v166
                                                                                        case -751:
                                                                                            if false {
                                                                                                if len(v54) >= 2 {
                                                                                                    v54 = v54[1:len(v54)]
                                                                                                }
                                                                                            } else {
                                                                                                if (!false) {
                                                                                                    var v171 float32 = 0.1
                                                                                                    _ = v171
                                                                                                }
                                                                                                switch v53 {
                                                                                                case -1:
                                                                                                    v3 = (v1 % ((fi134) | 1))
                                                                                                case 128:
                                                                                                    fmt.Println(fi134, fi63)
                                                                                                }
                                                                                                delete(v6, v3)
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    switch "hello" {
                                                                                    case "":
                                                                                        v172 := GenS14{F0: v1, F1: (v4 % ((uint32(v64)) | 1))}
                                                                                        _ = v172
                                                                                    default:
                                                                                        if (255 <= 134) {
                                                                                            var v173 int8 = v19
                                                                                            switch 9223372036854775807 {
                                                                                            case 1:
                                                                                                var v174 int16 = -32768
                                                                                                _ = v174
                                                                                            default:
                                                                                                if true {
                                                                                                    delete(v6, (v53 + 32767))
                                                                                                    if len(v54) >= 2 {
                                                                                                        v54 = v54[1:len(v54)]
                                                                                                    }
                                                                                                    v175 := GenS3{F0: 1, F1: false}
                                                                                                    _ = v175
                                                                                                } else {
                                                                                                    var v176 uint32 = 42
                                                                                                    _ = v176
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(v64)
                                                                                            _ = len(v54)
                                                                                            v6[(fi134 / ((32767) | 1))] = v4
                                                                                            if ((v64 != (v64 << ((255) & 63))) && false) {
                                                                                                v177 := v6[fi63]
                                                                                                _ = v177
                                                                                            } else {
                                                                                                var v178 int8 = (v173 * int8(v4))
                                                                                                _ = v178
                                                                                            }
                                                                                            if false {
                                                                                                var v179 uint8 = 128
                                                                                                _ = v179
                                                                                            }
                                                                                            var v180 uint8 = v64
                                                                                            _ = v173
                                                                                            _ = v180
                                                                                        } else {
                                                                                            switch (v3 - (-(v1))) {
                                                                                            case -2147483648:
                                                                                                v19 = (v19 + (v5 + genFn0(((((v154 + v154) + (v154 + "hello")) + ("\n\t" + v154)) + (v154 + "foo bar")))))
                                                                                            }
                                                                                            v5 = (v5 >> ((18446744073709551615) & 63))
                                                                                        }
                                                                                    }
                                                                                    var v181 float64 = 0.1
                                                                                    v182 := GenS7{F0: -408, F1: v5, F2: 0}
                                                                                    if true {
                                                                                        var v183 uint64 = 18446744073709551615
                                                                                        var v184 int64 = -725
                                                                                        _ = v183
                                                                                        _ = v184
                                                                                    } else {
                                                                                        switch (fi63 % (((v53 - (fi63 & v1))) | 1)) {
                                                                                        case 0:
                                                                                            if false {
                                                                                                var v185 uint32 = 1000
                                                                                                v186 := GenS16{F0: 18446744073709551615, F1: 4294967295}
                                                                                                v5 = (^(v5))
                                                                                                v187 := GenS11{F0: ((-1 >= 0) && true), F1: 42, F2: -155}
                                                                                                v5 -= -63
                                                                                                _ = v185
                                                                                                _ = v186
                                                                                                _ = v187
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    if len(v54) > 0 {
                                                                                        v54[0] = fi134
                                                                                    }
                                                                                    if (!((!true) && (0.0 > v2))) {
                                                                                        v181 = (v181 / v181)
                                                                                        var v188 uint = 751
                                                                                        _ = v188
                                                                                    }
                                                                                    var v189 float32 = v2
                                                                                    _ = v181
                                                                                    _ = v182
                                                                                    _ = v189
                                                                                }
                                                                                var v190 uint64 = 42
                                                                                v5 += (v19 & (-(v19)))
                                                                                _ = v190
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                v53 = v1
                                                            } else {
                                                                var v191 float32 = -2.25
                                                                if len(v54) > 0 {
                                                                    v54[0] = 128
                                                                }
                                                                _ = v191
                                                            }
                                                        }
                                                    }
                                                    _ = v153
                                                    _ = v154
                                                }
                                                switch (v53 / ((v1) | 1)) {
                                                case -1:
                                                    v192 := GenS17{F0: 1, F1: 2}
                                                    fmt.Println(v64)
                                                    v193 := GenS18{F0: v64, F1: -2147483648}
                                                    _ = v192
                                                    _ = v193
                                                case 392:
                                                    var v194 float32 = v2
                                                    _ = v194
                                                default:
                                                    var v195 uint16 = 471
                                                    _ = v195
                                                }
                                                v1 += int(v19)
                                            }
                                            fmt.Println(v4, v1, v3)
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v196 int64 = 0
                                                _ = v196
                                            case "":
                                                if true {
                                                    var v197 int16 = 1
                                                    var v198 bool = true
                                                    var v199 int = (v3 | (v53 ^ (fi134 | -581)))
                                                    _ = v197
                                                    _ = v198
                                                    _ = v199
                                                } else {
                                                    if false {
                                                        if len(v54) > 0 {
                                                            v54[0] = (-(fi134))
                                                        }
                                                        if len(v54) >= 2 {
                                                            v54 = v54[1:len(v54)]
                                                        }
                                                    } else {
                                                        if (1 > 1) {
                                                            if (18446744073709551615 < 4294967295) {
                                                                fmt.Println(v3)
                                                                fmt.Println(fi134, fi63)
                                                                v200 := GenS1{F0: 1, F1: v64, F2: (v4 & (v4 & v4))}
                                                                _ = v200
                                                            } else {
                                                                v201 := GenS19{F0: 2147483647, F1: 18446744073709551615, F2: 1}
                                                                var v202 int = v53
                                                                switch (fi134 & 32768) {
                                                                case 4294967295:
                                                                    var v203 int16 = -191
                                                                    _ = v203
                                                                case 9223372036854775807:
                                                                    v54 = append(v54, (v3 + v53))
                                                                default:
                                                                    fmt.Println(fi134)
                                                                }
                                                                _ = v201
                                                                _ = v202
                                                            }
                                                            _ = len(v6)
                                                        } else {
                                                            var v204 int64 = 42
                                                            v205 := GenS17{F0: -1, F1: 0}
                                                            _ = v204
                                                            _ = v205
                                                        }
                                                        var v206 int32 = 2147483647
                                                        switch v1 {
                                                        case 1:
                                                            _ = len(v54)
                                                        default:
                                                            var v207 uint32 = (^(v4))
                                                            _ = v207
                                                        }
                                                        _ = v206
                                                    }
                                                    v64 = (^(v64))
                                                    var v208 int8 = -17
                                                    _ = v208
                                                }
                                                if true {
                                                    if ("" != "foo bar") {
                                                        if ((!false) || false) {
                                                            var v209 bool = true
                                                            var v210 float32 = 0.0
                                                            switch ("foo bar" + "\n\t") {
                                                            case "foo bar":
                                                                if false {
                                                                    v211 := GenS12{F0: 1}
                                                                    fmt.Println(v19, v4, v53)
                                                                    v212 := GenS5{F0: 42}
                                                                    _ = v211
                                                                    _ = v212
                                                                } else {
                                                                    var v213 uint = uint(v64)
                                                                    if (float32(-2.25) > 0.0) {
                                                                        v214 := GenS8{F0: v209, F1: (v3 << ((v213) & 63))}
                                                                        _ = v214
                                                                    } else {
                                                                        v215 := GenS16{F0: (v213 / (((v213 << ((v213) & 63))) | 1)), F1: 0}
                                                                        var v216 float32 = v2
                                                                        var v217 int16 = -32768
                                                                        _ = v215
                                                                        _ = v216
                                                                        _ = v217
                                                                    }
                                                                    fmt.Println(v210, v53, v1)
                                                                    _ = v213
                                                                }
                                                                v5 -= 100
                                                            }
                                                            v19 = int8(v64)
                                                            if len(v54) >= 2 {
                                                                v54 = v54[1:len(v54)]
                                                            }
                                                            v53 = v1
                                                            var v218 uint64 = 667
                                                            v5 -= (v19 << ((2) & 63))
                                                            _ = v209
                                                            _ = v210
                                                            _ = v218
                                                        }
                                                        var v219 uint32 = (v4 * uint32(v64))
                                                        var v220 string = ""
                                                        _ = v219
                                                        _ = v220
                                                    } else {
                                                        v3 += fi134
                                                        var v221 uint64 = 558
                                                        _ = v221
                                                    }
                                                    v4 = v4
                                                    v222 := GenS20{F0: v64, F1: "\n\t", F2: -1}
                                                    _ = v222
                                                } else {
                                                    if false {
                                                        v54 = append(v54, 127)
                                                        v64 += 1
                                                        v53 += (-(fi63))
                                                    }
                                                    delete(v6, (v53 + fi63))
                                                }
                                                if len(v54) > 0 {
                                                    v54[0] = v1
                                                }
                                            default:
                                                var v223 uint32 = (v4 << ((201) & 63))
                                                _ = v223
                                            }
                                            var v224 int8 = genFn0("hello")
                                            _ = v224
                                        }
                                        fmt.Println(fi63, v3, v2)
                                        if false {
                                            if len(v54) > 0 {
                                                v54[0] = v53
                                            }
                                            var v225 int = fi63
                                            var v226 uint16 = 769
                                            switch 32767 {
                                            case -9223372036854775808:
                                                if len(v54) >= 2 {
                                                    v54 = v54[1:len(v54)]
                                                }
                                                v1 -= (v53 + v3)
                                                switch ((("\n\t" + "") + "foo bar") + "foo bar") {
                                                case "a":
                                                    var v227 int16 = -1
                                                    _ = v227
                                                case "foo bar":
                                                    if (-128 != v19) {
                                                        if len(v54) >= 2 {
                                                            v54 = v54[1:len(v54)]
                                                        }
                                                        var v228 int16 = 0
                                                        var v229 string = ""
                                                        var v230 uint32 = v4
                                                        _ = v228
                                                        _ = v229
                                                        _ = v230
                                                    } else {
                                                        v231 := GenS21{F0: 18446744073709551615, F1: (3.14159 == float64(((1e10 + 0.0) * v2)))}
                                                        v232 := GenS22{F0: 1}
                                                        _ = v231
                                                        _ = v232
                                                    }
                                                case "hello":
                                                    var v233 uint32 = v4
                                                    var v234 uint = 997
                                                    _ = v233
                                                    _ = v234
                                                }
                                            }
                                            if false {
                                                v235 := GenS11{F0: ((^(v226)) >= uint16(v4)), F1: 1000, F2: -9223372036854775808}
                                                var v236 uint = 255
                                                if false {
                                                    if false {
                                                        v19 -= -55
                                                    } else {
                                                        var v237 uint16 = v226
                                                        var v238 float32 = (v2 - (0.0 * v2))
                                                        v239 := GenS3{F0: -9223372036854775808, F1: true}
                                                        v236 -= v236
                                                        switch "\n\t" {
                                                        case "a":
                                                            delete(v6, fi63)
                                                        case "foo bar":
                                                            var v240 float64 = -2.25
                                                            _ = v240
                                                        default:
                                                            v241 := GenS20{F0: v64, F1: (("foo bar" + "") + ("foo bar" + "hello")), F2: 1}
                                                            _ = v241
                                                        }
                                                        _ = v237
                                                        _ = v238
                                                        _ = v239
                                                    }
                                                    switch (-(fi134)) {
                                                    case -9223372036854775808:
                                                        v1 -= v53
                                                    default:
                                                        if true {
                                                            if (!(("\n\t" == "\n\t") && true)) {
                                                                _ = len(v6)
                                                                v4 = (v4 >> ((v236) & 63))
                                                            }
                                                            v242 := GenS13{F0: 127}
                                                            if false {
                                                                var v243 uint64 = 423
                                                                v236 -= 1
                                                                var v244 float64 = 1e10
                                                                _ = v243
                                                                _ = v244
                                                            }
                                                            v245 := GenS10{F0: 65535, F1: -436}
                                                            v246 := GenS8{F0: false, F1: (v53 % (((^(v225))) | 1))}
                                                            _ = v242
                                                            _ = v245
                                                            _ = v246
                                                        }
                                                    }
                                                    v247 := GenS5{F0: -204}
                                                    v54 = append(v54, (v225 + int(v4)))
                                                    v248 := GenS10{F0: 0, F1: 1}
                                                    _ = v247
                                                    _ = v248
                                                }
                                                _ = v235
                                                _ = v236
                                            }
                                            _ = v225
                                            _ = v226
                                        }
                                        switch -2147483648 {
                                        case 128:
                                            var v249 int64 = 42
                                            _ = v249
                                        case 54:
                                            switch (v3 << ((15) & 63)) {
                                            case -2147483648:
                                                if true {
                                                    fmt.Println(v4, fi134)
                                                } else {
                                                    fmt.Println(v4, v3, v19)
                                                }
                                            default:
                                                if len(v54) > 0 {
                                                    v54[0] = 42
                                                }
                                            }
                                        case 4294967295:
                                            switch (v3 << ((874) & 63)) {
                                            case 0:
                                                _ = len(v54)
                                            case -959:
                                                if false {
                                                    switch (v1 - (v53 | (v53 ^ 4294967295))) {
                                                    case 42:
                                                        fmt.Println(v2)
                                                    case 1:
                                                        v250 := v6[fi63]
                                                        _ = v250
                                                    default:
                                                        var v251 uint64 = 764
                                                        _ = v251
                                                    }
                                                }
                                            default:
                                                switch v53 {
                                                case 0:
                                                    _ = len(v6)
                                                case -665:
                                                    var v252 bool = (0 > -1)
                                                    _ = v252
                                                }
                                            }
                                        }
                                        var v253 uint = 306
                                        switch "\n\t" {
                                        case "":
                                            var v254 uint64 = 18446744073709551615
                                            _ = v254
                                        case "a":
                                            v1 = v3
                                        default:
                                            if len(v54) > 0 {
                                                v54[0] = (fi134 & -472)
                                            }
                                        }
                                        v255 := GenS23{F0: (v64 % (((v64 >> ((18446744073709551615) & 63))) | 1)), F1: 65535, F2: (^(v4))}
                                        _ = v253
                                        _ = v255
                                    } else {
                                        if false {
                                            if true {
                                                _ = len(v54)
                                                v256 := GenS17{F0: 5, F1: 235}
                                                v4 -= 2147483648
                                                _ = v256
                                            } else {
                                                var v257 bool = false
                                                if (!(2 > 18446744073709551615)) {
                                                    if len(v54) > 0 {
                                                        v54[0] = (-(fi63))
                                                    }
                                                    if ((!("\n\t" <= "\n\t")) || false) {
                                                        v258 := GenS16{F0: uint(v64), F1: 42}
                                                        v3 -= 128
                                                        if (!((-(v5)) != int8(v64))) {
                                                            v4 = (v4 - v4)
                                                            if false {
                                                                var v259 float64 = float64(float32((1e10 / 1e10)))
                                                                v260 := GenS24{F0: (v257 && v257), F1: -9223372036854775808}
                                                                var v261 int = v3
                                                                _ = v259
                                                                _ = v260
                                                                _ = v261
                                                            } else {
                                                                v64 += 0
                                                                if (!(2 > 18446744073709551615)) {
                                                                    var v262 int16 = 42
                                                                    _ = v262
                                                                } else {
                                                                    if false {
                                                                        if ((v19 | genFn0("hello")) <= (v19 * (v5 % ((genFn0("a")) | 1)))) {
                                                                            v263 := GenS25{F0: 65535, F1: fi134, F2: -99}
                                                                            _ = v263
                                                                        }
                                                                        var v264 int16 = int16(v64)
                                                                        _ = v264
                                                                    } else {
                                                                        var v265 int32 = 1
                                                                        _ = v265
                                                                    }
                                                                }
                                                                v266 := GenS26{F0: v64}
                                                                _ = v266
                                                            }
                                                            delete(v6, v3)
                                                        }
                                                        _ = v258
                                                    } else {
                                                        var v267 float32 = 1e10
                                                        _ = v267
                                                    }
                                                } else {
                                                    var v268 bool = (!(!((true && v257) || (v257 && v257))))
                                                    var v269 float32 = v2
                                                    _ = v268
                                                    _ = v269
                                                }
                                                v270 := GenS17{F0: 1, F1: 32767}
                                                _ = v257
                                                _ = v270
                                            }
                                        } else {
                                            if len(v54) > 0 {
                                                v54[0] = (^(v3))
                                            }
                                            if (552 >= 339) {
                                                fmt.Println(v3, v64, fi63)
                                                v19 = (v5 & v5)
                                                v19 += v19
                                            }
                                            fmt.Println(v19)
                                        }
                                        v64 = (^(v64))
                                        if false {
                                            if len(v54) > 0 {
                                                v54[0] = 516
                                            }
                                        }
                                        v54 = append(v54, (v1 << ((2) & 63)))
                                        v271 := GenS27{F0: ("a" + "a"), F1: (-(v19)), F2: ""}
                                        _ = v271
                                    }
                                }
                                v4 -= 2
                                if (32767 < 665) {
                                    fmt.Println(v5, v53, v2)
                                    if (!(1 == -429)) {
                                        v272 := GenS27{F0: "", F1: (v19 - v5), F2: "foo bar"}
                                        _ = v272
                                    } else {
                                        v273 := GenS28{F0: (42 != 937), F1: int8(v53), F2: 23}
                                        v3 = 0
                                        _ = v273
                                    }
                                    var v274 uint32 = 1
                                    _ = v274
                                } else {
                                    fmt.Println(v4, v5, v53)
                                    var v275 bool = true
                                    var v276 int32 = 304
                                    _ = v275
                                    _ = v276
                                }
                                _ = fi134
                            }
                            _ = fi63
                            _ = v64
                        }
                    } else {
                        if true {
                            fmt.Println(v3, v2)
                            v4 = 42
                            v277 := GenS26{F0: 0}
                            _ = v277
                        } else {
                            fmt.Println(v2)
                        }
                    }
                case "foo bar":
                    if false {
                        v4 = v4
                        switch "a" {
                        case "":
                            v6[(v53 ^ -9223372036854775808)] = v4
                        }
                        var v278 int32 = 1
                        _ = v278
                    } else {
                        fmt.Println(v53, v2, v3)
                        v2 = v2
                    }
                case "\n\t":
                    if false {
                        fmt.Println(v19)
                    }
                default:
                    fmt.Println(v3)
                }
            } else {
                for fi279 := 0; fi279 < 5; fi279++ {
                    var v280 uint8 = 42
                    v281 := v6[32768]
                    _ = fi279
                    _ = v280
                    _ = v281
                }
                _ = len(v54)
                for fi282 := 0; fi282 < 2; fi282++ {
                    v53 -= (v3 >> ((42) & 63))
                    _ = len(v54)
                    var v283 uint64 = 18446744073709551615
                    _ = fi282
                    _ = v283
                }
            }
        }
        _ = v54
    } else {
        var v284 uint = 65535
        v5 = (v19 | (v5 * v19))
        v5 += genFn0("foo bar")
        _ = v284
    }
    v19 = v5
    var v285 float32 = (float32(0.1) * 3.14159)
    v3 -= (v3 | (v1 >> ((18446744073709551615) & 63)))
    delete(v6, (-(v3)))
    v2 -= 0.0
    if ((-612 >= 1) || (18446744073709551615 == 65535)) {
        if (true && false) {
            v2 = v2
            var v286 float32 = 3.14159
            _ = v286
        }
        v4 -= (v4 ^ (v4 << ((1) & 63)))
    }
    switch "foo bar" {
    case "\n\t":
        switch "foo bar" {
        case "a":
            if (1 >= 53) {
                var v287 int16 = 32767
                for fi288 := 0; fi288 < 14; fi288++ {
                    fmt.Println(v5, v1, v1)
                    var v289 uint32 = (^(v4))
                    for fi290 := 0; fi290 < 1; fi290++ {
                        if false {
                            v291 := GenS28{F0: (1 <= (v287 - -325)), F1: -1, F2: 42}
                            var v292 float64 = ((0.0 / float64(float32(0.1))) / 1.0)
                            v289 = 42
                            _ = v291
                            _ = v292
                        } else {
                            v2 -= v285
                        }
                        v293 := GenS17{F0: 573, F1: 630}
                        var v294 uint64 = 42
                        _ = fi290
                        _ = v293
                        _ = v294
                    }
                    _ = fi288
                    _ = v289
                }
                _ = v287
            } else {
                v295 := GenS27{F0: (("\n\t" + "hello") + "\n\t"), F1: (v5 % (((v5 * genFn0("a"))) | 1)), F2: ""}
                v296 := []int{v3, (v1 + (v3 + (v53 << ((42) & 63))))}
                _ = v295
                _ = v296
            }
        case "foo bar":
            v297 := []int{int(v5)}
            _ = v297
        }
    case "a":
        if (!true) {
            for fi298 := 0; fi298 < 16; fi298++ {
                v4 -= (v4 << ((171) & 63))
                _ = fi298
            }
            v4 -= (v4 + v4)
            switch v3 {
            case 127:
                var v299 int16 = 1000
                _ = v299
            case 0:
                var v300 uint = 0
                _ = v300
            case 9223372036854775807:
                var v301 uint32 = (v4 - (v4 / ((v4) | 1)))
                _ = v301
            default:
                for fi302 := 0; fi302 < 12; fi302++ {
                    var v303 uint = 508
                    _ = fi302
                    _ = v303
                }
            }
        }
    case "":
        var v304 int32 = 2147483647
        _ = v304
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v285:", v285)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v285
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
}
