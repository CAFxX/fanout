package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 bool; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int16; F1 int16; F2 int32 }
type GenS6 struct { F0 int8; F1 bool; F2 bool }
type GenS7 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint64; F1 int8; F2 string }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint64; F1 string }
func genFn0(p0 uint8, p1 int, p2 string) int {
    switch ("" + "\n\t") {
    case "foo bar":
        p1 -= p1
    case "":
        v5 := GenS1{F0: 255}
        _ = v5
    case "a":
        var v6 uint64 = 643
        _ = v6
    default:
        p2 = p2
    }
    if false {
        switch p2 {
        case "hello":
            for fi7 := 0; fi7 < 5; fi7++ {
                for fi8 := 0; fi8 < 10; fi8++ {
                    v9 := []int{fi7, (fi8 >> ((65535) & 63))}
                    p2 = p2
                    v10 := GenS2{F0: (fi8 | (^(fi7))), F1: (false && true), F2: 1000}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                p0 = 127
                for fi11 := 0; fi11 < 8; fi11++ {
                    var v12 float64 = 1e10
                    _ = fi11
                    _ = v12
                }
                _ = fi7
            }
        default:
            switch ("hello" + p2) {
            case "hello":
                var v13 uint = 2
                _ = v13
            case "foo bar":
                v14 := []int{42, 2147483647, -1}
                _ = v14
            case "\n\t":
                switch (p1 << ((0) & 63)) {
                case 127:
                    p1 -= (p1 - p1)
                }
            default:
                if true {
                    if (!(true && false)) {
                        if false {
                            switch p2 {
                            case "\n\t":
                                if true {
                                    for fi15 := 0; fi15 < 8; fi15++ {
                                        if (-1 >= 180) {
                                            var v16 uint32 = 13
                                            if false {
                                                switch p2 {
                                                case "":
                                                    v17 := GenS3{F0: 1}
                                                    _ = v17
                                                case "foo bar":
                                                    var v18 int = p1
                                                    _ = v18
                                                case "hello":
                                                    v19 := GenS4{F0: 1000, F1: -9223372036854775808}
                                                    _ = v19
                                                default:
                                                    p2 = "\n\t"
                                                }
                                            } else {
                                                var v20 int = 4294967295
                                                fmt.Println(p0, v16, fi15)
                                                for fi21 := 0; fi21 < 5; fi21++ {
                                                    v20 = (fi15 ^ p1)
                                                    var v22 int64 = -9223372036854775808
                                                    var v23 float32 = 0.1
                                                    p1 = (-(fi15))
                                                    _ = fi21
                                                    _ = v22
                                                    _ = v23
                                                }
                                                _ = v20
                                            }
                                            v24 := GenS2{F0: (p1 >> ((1) & 63)), F1: (-9223372036854775808 >= fi15), F2: -1}
                                            _ = v16
                                            _ = v24
                                        }
                                        var v25 uint16 = 1
                                        p2 = p2
                                        _ = fi15
                                        _ = v25
                                    }
                                    switch "hello" {
                                    case "foo bar":
                                        for fi26 := 0; fi26 < 9; fi26++ {
                                            var v27 bool = (((p0 / ((p0) | 1)) <= p0) || (-760 <= 672))
                                            _ = fi26
                                            _ = v27
                                        }
                                    case "\n\t":
                                        switch p1 {
                                        case -367:
                                            v28 := GenS3{F0: -1}
                                            _ = v28
                                        }
                                    default:
                                        for fi29 := 0; fi29 < 10; fi29++ {
                                            p0 = p0
                                            _ = fi29
                                        }
                                    }
                                } else {
                                    fmt.Println(p0, p0, p0)
                                }
                            case "foo bar":
                                for fi30 := 0; fi30 < 10; fi30++ {
                                    v31 := []int{(p1 - (p1 + 9223372036854775807)), fi30, (p1 % ((fi30) | 1)), p1}
                                    if false {
                                        var v32 uint32 = 2147483647
                                        _ = v32
                                    }
                                    if len(v31) >= 2 {
                                        v31 = v31[1:len(v31)]
                                    }
                                    _ = fi30
                                    _ = v31
                                }
                            case "a":
                                v33 := []int{(p1 | p1)}
                                _ = v33
                            }
                        }
                    }
                } else {
                    v34 := GenS2{F0: p1, F1: false, F2: 138}
                    v35 := GenS3{F0: 42}
                    _ = v34
                    _ = v35
                }
            }
        }
        for fi36 := 0; fi36 < 1; fi36++ {
            for fi37 := 0; fi37 < 14; fi37++ {
                p2 = "a"
                v38 := GenS3{F0: 875}
                var v39 bool = true
                _ = fi37
                _ = v38
                _ = v39
            }
            _ = fi36
        }
    }
    p2 = p2
    var v40 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v40
    return (p1 | (p1 << (((v40 / ((v40) | 1))) & 63)))
}

func main() {
    var v1 uint16 = 255
    var v2 string = "hello"
    var v3 float32 = ((1e10 / (1.5 * float32(-2.25))) / 1.5)
    var v4 int16 = -32768
    v41 := []int{-991}
    var v42 uint32 = 993
    v1 += (v1 >> ((18446744073709551615) & 63))
    fmt.Println(v42)
    var v43 uint8 = 42
    switch genFn0(v43, 32768, v2) {
    case 32768:
        if (!false) {
            v44 := GenS5{F0: v4, F1: (v4 | v4), F2: -745}
            if (329 <= v42) {
                switch ((v2 + "a") + "hello") {
                case "foo bar":
                    if len(v41) > 0 {
                        v41[0] = 32768
                    }
                default:
                    v45 := GenS6{F0: -128, F1: (v2 != v2), F2: (uint64(v42) == 0)}
                    _ = v45
                }
                fmt.Println(v3)
                v4 -= v4
            }
            v42 += v42
            _ = v44
        } else {
            switch v2 {
            case "foo bar":
                if true {
                    for fi46 := 0; fi46 < 7; fi46++ {
                        v47 := GenS4{F0: (^(v4)), F1: 1}
                        if false {
                            for fi48 := 0; fi48 < 19; fi48++ {
                                v42 -= (v42 >> ((607) & 63))
                                var v49 uint64 = 42
                                _ = fi48
                                _ = v49
                            }
                            v42 = v42
                            var v50 uint = uint(v1)
                            _ = v50
                        }
                        var v51 float32 = (v3 / 1.5)
                        _ = fi46
                        _ = v47
                        _ = v51
                    }
                    if (!true) {
                        for fi52 := 0; fi52 < 10; fi52++ {
                            for fi53 := 0; fi53 < 11; fi53++ {
                                var v54 uint32 = (^(v42))
                                _ = fi53
                                _ = v54
                            }
                            _ = fi52
                        }
                    } else {
                        v55 := GenS2{F0: 0, F1: (genFn0(0, 1, v2) > 9223372036854775807), F2: v4}
                        v56 := GenS5{F0: (v4 - v4), F1: (v4 - -1), F2: -1}
                        _ = v55
                        _ = v56
                    }
                    for fi57 := 0; fi57 < 6; fi57++ {
                        for fi58 := 0; fi58 < 15; fi58++ {
                            v59 := GenS3{F0: 42}
                            var v60 int16 = 1000
                            v61 := GenS5{F0: (v60 - v4), F1: v60, F2: 927}
                            _ = fi58
                            _ = v59
                            _ = v60
                            _ = v61
                        }
                        v2 = v2
                        _ = fi57
                    }
                } else {
                    for fi62 := 0; fi62 < 2; fi62++ {
                        v43 -= 0
                        if ((v43 - (v43 % ((v43) | 1))) < v43) {
                            var v63 int16 = (^(v4))
                            var v64 bool = true
                            v3 = float32(0.1)
                            _ = v63
                            _ = v64
                        }
                        for fi65 := 0; fi65 < 18; fi65++ {
                            switch "foo bar" {
                            case "a":
                                fmt.Println(fi62, v2, v2)
                            case "hello":
                                if false {
                                    v42 += uint32(v1)
                                } else {
                                    switch (fi62 - (-(fi62))) {
                                    case -2147483648:
                                        v43 += v43
                                        v66 := GenS4{F0: v4, F1: 1}
                                        switch (^(fi65)) {
                                        case 32768:
                                            v67 := GenS7{F0: (v42 & (v42 >> ((0) & 63))), F1: -128, F2: 2}
                                            _ = v67
                                        default:
                                            fmt.Println(v4, v4)
                                        }
                                        v68 := GenS2{F0: genFn0((v43 | 78), fi62, "foo bar"), F1: false, F2: -1}
                                        v41 = append(v41, (^(fi62)))
                                        _ = v66
                                        _ = v68
                                    case 9223372036854775807:
                                        fmt.Println(v3)
                                    default:
                                        v41 = append(v41, 9223372036854775807)
                                    }
                                    v1 = (v1 >> ((255) & 63))
                                }
                            }
                            switch (^(fi65)) {
                            case 820:
                                v2 = ""
                            case -1:
                                var v69 uint8 = uint8(v1)
                                _ = v69
                            case 9223372036854775807:
                                v3 += v3
                            }
                            if false {
                                v70 := GenS3{F0: 42}
                                v43 = 2
                                v71 := GenS3{F0: -2147483648}
                                _ = v70
                                _ = v71
                            }
                            _ = fi65
                        }
                        _ = fi62
                    }
                    fmt.Println(v3)
                    switch 42 {
                    case 42:
                        v41 = append(v41, -9223372036854775808)
                    default:
                        v41 = append(v41, 128)
                    }
                }
            case "\n\t":
                fmt.Println(v1)
            }
            if true {
                if false {
                    switch (v2 + v2) {
                    case "":
                        for fi72 := 0; fi72 < 13; fi72++ {
                            switch ("" + v2) {
                            case "hello":
                                if false {
                                    if ((!((!(-50 != 42)) || (!false))) && (251 <= v1)) {
                                        if (v1 <= 2) {
                                            v4 = 4
                                        } else {
                                            var v73 uint32 = uint32(v43)
                                            fmt.Println(v4, v73)
                                            var v74 float32 = (0.0 + (1.5 / 1.5))
                                            _ = v73
                                            _ = v74
                                        }
                                        fmt.Println(v43)
                                        if ((fi72 / ((-166) | 1)) < (-(fi72))) {
                                            var v75 float64 = 3.14159
                                            if false {
                                                if len(v41) > 0 {
                                                    v41[0] = fi72
                                                }
                                                fmt.Println(v43)
                                                var v76 int8 = 1
                                                _ = v76
                                            }
                                            v3 += ((3.14159 / v3) + 0.1)
                                            _ = v75
                                        }
                                    }
                                    var v77 int16 = v4
                                    _ = v77
                                }
                            case "":
                                fmt.Println(v2, v43)
                            default:
                                if true {
                                    if true {
                                        for fi78 := 0; fi78 < 17; fi78++ {
                                            v79 := GenS2{F0: 128, F1: true, F2: v4}
                                            if (-2.25 != 3.14159) {
                                                switch fi72 {
                                                case 975:
                                                    var v80 int = fi78
                                                    v43 = 128
                                                    v81 := GenS6{F0: 1, F1: (18446744073709551615 == 577), F2: true}
                                                    _ = v80
                                                    _ = v81
                                                }
                                                v43 = (v43 | v43)
                                            }
                                            _ = fi78
                                            _ = v79
                                        }
                                    } else {
                                        var v82 float32 = v3
                                        if (v42 != 2) {
                                            if false {
                                                for fi83 := 0; fi83 < 20; fi83++ {
                                                    var v84 float64 = 1.5
                                                    if (true || (v1 > 42)) {
                                                        if false {
                                                            v82 -= float32((v84 + ((((3.14159 + v84) - 1.5) * v84) + (((0.0 * float64(float32(float64(v82)))) - v84) - v84))))
                                                        } else {
                                                            if len(v41) >= 2 {
                                                                v41 = v41[1:len(v41)]
                                                            }
                                                        }
                                                    } else {
                                                        v3 = v3
                                                    }
                                                    _ = fi83
                                                    _ = v84
                                                }
                                                v4 -= (v4 - v4)
                                            }
                                            if false {
                                                v42 += (v42 + 17)
                                                if (!((v4 % (((v4 - -797)) | 1)) == v4)) {
                                                    v85 := GenS8{F0: v42}
                                                    if (4294967295 > 42) {
                                                        fmt.Println(v82, v4, v42)
                                                        for fi86 := 0; fi86 < 19; fi86++ {
                                                            var v87 uint32 = (^(v42))
                                                            if false {
                                                                var v88 int8 = -128
                                                                switch v2 {
                                                                case "hello":
                                                                    v4 = (v4 - (v4 | v4))
                                                                case "":
                                                                    if len(v41) >= 2 {
                                                                        v41 = v41[1:len(v41)]
                                                                    }
                                                                default:
                                                                    var v89 uint64 = 450
                                                                    _ = v89
                                                                }
                                                                var v90 uint64 = uint64(v43)
                                                                _ = v88
                                                                _ = v90
                                                            } else {
                                                                var v91 uint = 0
                                                                _ = v91
                                                            }
                                                            _ = fi86
                                                            _ = v87
                                                        }
                                                        if len(v41) >= 2 {
                                                            v41 = v41[1:len(v41)]
                                                        }
                                                    }
                                                    v41 = append(v41, fi72)
                                                    _ = v85
                                                }
                                                var v92 uint32 = (v42 & v42)
                                                _ = v92
                                            }
                                            var v93 float32 = float32(((1e10 + -2.25) / 0.1))
                                            _ = v93
                                        }
                                        _ = v82
                                    }
                                }
                            }
                            switch (fi72 % ((42) | 1)) {
                            case 105:
                                if true {
                                    v3 = (float32(1e10) + (0.1 * (v3 + 1.5)))
                                    for fi94 := 0; fi94 < 14; fi94++ {
                                        if false {
                                            v95 := GenS8{F0: (v42 & (v42 * v42))}
                                            var v96 float32 = v3
                                            v4 = 42
                                            _ = v95
                                            _ = v96
                                        } else {
                                            if ((^(v1)) >= (v1 / (((v1 >> ((65535) & 63))) | 1))) {
                                                var v97 uint32 = v42
                                                _ = v97
                                            }
                                            var v98 int8 = -30
                                            _ = v98
                                        }
                                        var v99 int8 = -77
                                        _ = fi94
                                        _ = v99
                                    }
                                    if len(v41) >= 2 {
                                        v41 = v41[1:len(v41)]
                                    }
                                }
                            default:
                                v43 += v43
                            }
                            _ = fi72
                        }
                    case "\n\t":
                        for fi100 := 0; fi100 < 16; fi100++ {
                            v1 = v1
                            var v101 uint32 = (v42 & (v42 % ((v42) | 1)))
                            if true {
                                var v102 bool = false
                                fmt.Println(fi100)
                                _ = v102
                            } else {
                                v2 = "foo bar"
                            }
                            _ = fi100
                            _ = v101
                        }
                    default:
                        switch genFn0((v43 - 128), -389, v2) {
                        case 0:
                            if false {
                                if (true && true) {
                                    if (false || (!(false || true))) {
                                        for fi103 := 0; fi103 < 19; fi103++ {
                                            switch v2 {
                                            case "hello":
                                                for fi104 := 0; fi104 < 17; fi104++ {
                                                    v105 := GenS3{F0: 0}
                                                    _ = fi104
                                                    _ = v105
                                                }
                                            }
                                            for fi106 := 0; fi106 < 11; fi106++ {
                                                v43 += 255
                                                if len(v41) > 0 {
                                                    v41[0] = 0
                                                }
                                                _ = fi106
                                            }
                                            if true {
                                                v107 := GenS8{F0: v42}
                                                switch v2 {
                                                case "a":
                                                    var v108 int32 = -2147483648
                                                    _ = v108
                                                case "foo bar":
                                                    for fi109 := 0; fi109 < 18; fi109++ {
                                                        fmt.Println(fi109, fi103)
                                                        fmt.Println(v42)
                                                        v1 = 810
                                                        _ = fi109
                                                    }
                                                default:
                                                    v3 += v3
                                                }
                                                v110 := GenS9{F0: 18446744073709551615, F1: int8(v1), F2: v2}
                                                _ = v107
                                                _ = v110
                                            } else {
                                                _ = len(v41)
                                                v42 = v42
                                            }
                                            _ = fi103
                                        }
                                        if false {
                                            var v111 float64 = (0.1 * 0.1)
                                            if len(v41) >= 2 {
                                                v41 = v41[1:len(v41)]
                                            }
                                            var v112 uint8 = 128
                                            _ = v111
                                            _ = v112
                                        } else {
                                            for fi113 := 0; fi113 < 6; fi113++ {
                                                v41 = append(v41, -414)
                                                for fi114 := 0; fi114 < 5; fi114++ {
                                                    var v115 string = (v2 + ("" + v2))
                                                    if len(v41) >= 2 {
                                                        v41 = v41[1:len(v41)]
                                                    }
                                                    var v116 bool = false
                                                    var v117 int8 = 42
                                                    fmt.Println(v1, v42, v4)
                                                    _ = fi114
                                                    _ = v115
                                                    _ = v116
                                                    _ = v117
                                                }
                                                _ = fi113
                                            }
                                            v43 -= (v43 * (v43 ^ (v43 % ((v43) | 1))))
                                            if true {
                                                v42 = 620
                                                if false {
                                                    var v118 int = -9223372036854775808
                                                    for fi119 := 0; fi119 < 10; fi119++ {
                                                        var v120 int64 = -45
                                                        _ = fi119
                                                        _ = v120
                                                    }
                                                    _ = v118
                                                }
                                                var v121 float32 = ((float32(0.1) / float32((0.0 - (1e10 + 3.14159)))) * (1e10 * (float32(((-2.25 - (float64(v3) / 1e10)) - 1.5)) - v3)))
                                                _ = v121
                                            }
                                        }
                                    } else {
                                        if true {
                                            v122 := GenS10{F0: false}
                                            var v123 uint64 = 561
                                            _ = v122
                                            _ = v123
                                        }
                                        if len(v41) > 0 {
                                            v41[0] = 1
                                        }
                                    }
                                    var v124 int32 = 622
                                    _ = v124
                                }
                                v125 := GenS1{F0: 255}
                                v126 := GenS11{F0: 2, F1: v2}
                                _ = v125
                                _ = v126
                            } else {
                                if len(v41) >= 2 {
                                    v41 = v41[1:len(v41)]
                                }
                                v1 += 1
                            }
                        case 128:
                            if false {
                                for fi127 := 0; fi127 < 15; fi127++ {
                                    fmt.Println(v3)
                                    v4 = (v4 - (^(v4)))
                                    for fi128 := 0; fi128 < 19; fi128++ {
                                        v129 := GenS11{F0: 2, F1: v2}
                                        fmt.Println(v3)
                                        fmt.Println(fi127)
                                        _ = fi128
                                        _ = v129
                                    }
                                    _ = fi127
                                }
                                if len(v41) >= 2 {
                                    v41 = v41[1:len(v41)]
                                }
                                var v130 int32 = 2147483647
                                _ = v130
                            }
                        default:
                            v42 = v42
                        }
                    }
                    v1 -= (v1 << ((291) & 63))
                } else {
                    v4 = v4
                }
            } else {
                v43 -= 42
                v131 := GenS9{F0: 42, F1: -128, F2: "hello"}
                _ = v131
            }
        }
    }
    v3 = (v3 * (v3 / (v3 - v3)))
    var v132 uint64 = 115
    var v133 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v132:", v132)
    fmt.Println("v133:", v133)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v132
    _ = v133
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
}
