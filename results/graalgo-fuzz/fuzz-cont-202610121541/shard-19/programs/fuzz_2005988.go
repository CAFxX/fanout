package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int32; F1 uint64; F2 int64 }
type GenS4 struct { F0 int; F1 uint64; F2 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 int8; F2 uint16 }
type GenS7 struct { F0 int8; F1 bool; F2 uint }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint16; F1 uint16; F2 uint }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int8; F1 bool; F2 int8 }
type GenS12 struct { F0 int32; F1 int32; F2 int8 }
type GenS13 struct { F0 uint; F1 uint8; F2 uint32 }
func genFn0(p0 uint8, p1 int) int16 {
    var v6 int64 = 1
    if (42 != 55) {
        var v7 float64 = 1.5
        v6 = 1
        _ = v7
    }
    switch "a" {
    case "hello":
        p0 = (p0 ^ (p0 | (p0 + p0)))
    }
    fmt.Println(v6, v6, p1)
    _ = p0
    _ = p1
    _ = v6
    return 0
}
func genFn1(p0 uint16, p1 uint64) uint32 {
    v8 := GenS1{F0: "", F1: 2}
    v9 := GenS1{F0: "\n\t", F1: 0}
    fmt.Println(p1, p1)
    p1 -= p1
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    var v1 uint64 = 142
    var v2 int16 = 1000
    var v3 float32 = (1e10 * 1e10)
    var v4 int = 516
    v5 := []uint32{2}
    switch "foo bar" {
    case "\n\t":
        if true {
            var v10 int32 = -2147483648
            var v11 int32 = (v10 * v10)
            _ = v10
            _ = v11
        }
    case "hello":
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 uint16 = 630
            if false {
                var v14 float64 = 0.0
                v14 = (v14 - 0.0)
                if ((851 == (v2 << ((2) & 63))) && (((v13 <= (v13 >> ((2) & 63))) && false) && true)) {
                    v14 = (v14 / (3.14159 - (v14 - v14)))
                    var v15 int = v4
                    _ = v15
                } else {
                    v3 -= float32(-2.25)
                }
                _ = v14
            }
            _ = fi12
            _ = v13
        }
    }
    var v16 int64 = -344
    var v17 int64 = (v16 + (v16 >> ((65535) & 63)))
    v18 := GenS2{F0: true, F1: 843}
    if true {
        fmt.Println(v17, v16)
    } else {
        v5 = append(v5, 2147483647)
        fmt.Println(v2, v1)
        v19 := GenS3{F0: 2147483647, F1: 0, F2: 42}
        _ = v19
    }
    var v20 bool = true
    for fi21 := 0; fi21 < 7; fi21++ {
        for fi22 := 0; fi22 < 11; fi22++ {
            if ((v4 - fi22) != v4) {
                var v23 int64 = (^(v16))
                v1 -= 333
                _ = v23
            }
            v24 := GenS1{F0: "\n\t", F1: 65535}
            _ = fi22
            _ = v24
        }
        if (172 <= (^(v4))) {
            v3 += (v3 * (v3 * v3))
            if len(v5) > 0 {
                v5[0] = 2147483647
            }
        }
        v16 = (v16 >> ((48) & 63))
        _ = fi21
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        fmt.Println(v16, v17)
        v17 = (v16 << ((0) & 63))
        _ = fi25
    }
    var v26 uint64 = (v1 / (((^(v1))) | 1))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v27 := GenS3{F0: -597, F1: v26, F2: 9223372036854775807}
    for fi28 := 0; fi28 < 6; fi28++ {
        v16 = (v16 - v16)
        _ = fi28
    }
    if (!(32767 < 930)) {
        if v20 {
            fmt.Println(v20, v17)
        } else {
            if v20 {
                var v29 int16 = genFn0(4, (v4 << ((4294967295) & 63)))
                var v30 int32 = 1
                _ = v29
                _ = v30
            } else {
                if false {
                    switch -9223372036854775808 {
                    case -1:
                        v31 := GenS4{F0: v4, F1: v26, F2: (v2 ^ v2)}
                        _ = v31
                    case 2147483647:
                        if v20 {
                            if v20 {
                                if ((!v20) && false) {
                                    switch (v4 | (v4 * (v4 & v4))) {
                                    case -2147483648:
                                        if (!(-947 >= -9223372036854775808)) {
                                            if len(v5) > 0 {
                                                v5[0] = 42
                                            }
                                            fmt.Println(v20, v20)
                                            v20 = (false || v20)
                                        }
                                    case 128:
                                        if false {
                                            var v32 int16 = genFn0(252, 127)
                                            var v33 float32 = v3
                                            _ = v32
                                            _ = v33
                                        } else {
                                            v20 = true
                                        }
                                    case -395:
                                        v20 = (genFn1(65535, v26) < 0)
                                    default:
                                        v34 := GenS4{F0: 128, F1: v1, F2: v2}
                                        _ = v34
                                    }
                                    if len(v5) > 0 {
                                        v5[0] = 935
                                    }
                                }
                                if v20 {
                                    if (!v20) {
                                        var v35 int = v4
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        v36 := GenS3{F0: 1, F1: 0, F2: v16}
                                        _ = v35
                                        _ = v36
                                    } else {
                                        v37 := GenS5{F0: (v2 * v2)}
                                        switch "" {
                                        case "\n\t":
                                            v26 = v26
                                        case "hello":
                                            v1 += (^(v26))
                                        case "foo bar":
                                            for fi38 := 0; fi38 < 2; fi38++ {
                                                v17 += (v16 * (v17 % (((v17 | (v17 | v16))) | 1)))
                                                for fi39 := 0; fi39 < 6; fi39++ {
                                                    if (true && v20) {
                                                        var v40 int8 = 42
                                                        if v20 {
                                                            _ = len(v5)
                                                            if false {
                                                                fmt.Println(v2)
                                                                var v41 uint16 = 1
                                                                _ = v41
                                                            }
                                                            if (1 > 122) {
                                                                v1 += v1
                                                                var v42 uint8 = 160
                                                                v43 := GenS1{F0: "", F1: 18446744073709551615}
                                                                _ = v42
                                                                _ = v43
                                                            }
                                                        }
                                                        _ = v40
                                                    } else {
                                                        fmt.Println(v20)
                                                        if len(v5) > 0 {
                                                            v5[0] = 241
                                                        }
                                                        if (!v20) {
                                                            var v44 uint64 = v1
                                                            v5 = append(v5, genFn1(731, v26))
                                                            var v45 string = "\n\t"
                                                            _ = v44
                                                            _ = v45
                                                        }
                                                    }
                                                    _ = fi39
                                                }
                                                if (genFn0(180, (v4 ^ v4)) != (v2 * 1)) {
                                                    if (!(42 == 74)) {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            if false {
                                                                var v46 uint64 = v26
                                                                _ = v46
                                                            } else {
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                v5 = append(v5, 2147483647)
                                                            }
                                                        case "foo bar":
                                                            v17 -= (^(v16))
                                                        default:
                                                            v16 -= (v16 >> ((255) & 63))
                                                        }
                                                        switch "" {
                                                        case "foo bar":
                                                            var v47 bool = (0.1 > 0.0)
                                                            _ = v47
                                                        default:
                                                            fmt.Println(fi38, v4)
                                                        }
                                                    } else {
                                                        v26 = v1
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    }
                                                    var v48 uint8 = 127
                                                    _ = v48
                                                }
                                                _ = fi38
                                            }
                                        }
                                        for fi49 := 0; fi49 < 13; fi49++ {
                                            var v50 uint32 = 436
                                            _ = fi49
                                            _ = v50
                                        }
                                        _ = v37
                                    }
                                    v1 = (v26 % ((2) | 1))
                                }
                            } else {
                                v51 := GenS4{F0: v4, F1: v26, F2: (v2 ^ genFn0(255, (^(v4))))}
                                var v52 float32 = float32((float64(float32((-2.25 - (0.1 - float64(v3))))) + 0.0))
                                v53 := GenS6{F0: (v2 ^ -803), F1: 20, F2: 0}
                                _ = v51
                                _ = v52
                                _ = v53
                            }
                        } else {
                            v17 = v17
                            if (127 >= uint8(v26)) {
                                v16 = (v17 + 1)
                                var v54 int = (v4 | (v4 << ((2) & 63)))
                                for fi55 := 0; fi55 < 13; fi55++ {
                                    switch 2147483647 {
                                    case 9223372036854775807:
                                        v20 = v20
                                    case 128:
                                        var v56 bool = false
                                        _ = v56
                                    default:
                                        for fi57 := 0; fi57 < 9; fi57++ {
                                            v20 = (((v54 - 32768) != int(v16)) && (1 >= 127))
                                            switch (fi55 >> ((112) & 63)) {
                                            case -203:
                                                v58 := GenS5{F0: 921}
                                                _ = v58
                                            }
                                            _ = fi57
                                        }
                                    }
                                    var v59 float32 = v3
                                    var v60 int = -9223372036854775808
                                    _ = fi55
                                    _ = v59
                                    _ = v60
                                }
                                _ = v54
                            } else {
                                for fi61 := 0; fi61 < 7; fi61++ {
                                    if true {
                                        if len(v5) > 0 {
                                            v5[0] = 4294967295
                                        }
                                        if v20 {
                                            if v20 {
                                                if (328 != 42) {
                                                    if (!(!false)) {
                                                        for fi62 := 0; fi62 < 6; fi62++ {
                                                            v2 = v2
                                                            var v63 int8 = 127
                                                            _ = fi62
                                                            _ = v63
                                                        }
                                                        switch -9223372036854775808 {
                                                        case 32768:
                                                            var v64 uint8 = 201
                                                            _ = v64
                                                        case -1:
                                                            fmt.Println(v2)
                                                        case 4294967295:
                                                            v65 := GenS7{F0: 42, F1: v20, F2: 856}
                                                            _ = v65
                                                        }
                                                    } else {
                                                        for fi66 := 0; fi66 < 20; fi66++ {
                                                            var v67 int8 = 42
                                                            _ = fi66
                                                            _ = v67
                                                        }
                                                        switch (fi61 / (((-(fi61))) | 1)) {
                                                        case -573:
                                                            var v68 int8 = -121
                                                            _ = v68
                                                        }
                                                        switch fi61 {
                                                        case 1:
                                                            fmt.Println(v26, v26)
                                                        case -9223372036854775808:
                                                            switch v4 {
                                                            case -832:
                                                                var v69 int = (fi61 | -143)
                                                                _ = v69
                                                            case 42:
                                                                switch v4 {
                                                                case 127:
                                                                    v70 := GenS2{F0: v20, F1: 175}
                                                                    _ = v70
                                                                }
                                                            case 127:
                                                                var v71 int = (v4 / ((v4) | 1))
                                                                _ = v71
                                                            }
                                                        default:
                                                            for fi72 := 0; fi72 < 4; fi72++ {
                                                                v16 -= v16
                                                                _ = fi72
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    if (-2.25 == -2.25) {
                                                        v16 += v16
                                                        if true {
                                                            if v20 {
                                                                switch 964 {
                                                                case 2147483647:
                                                                    switch -617 {
                                                                    case 128:
                                                                        v73 := GenS7{F0: 42, F1: ((v2 % ((v2) | 1)) <= genFn0(128, 42)), F2: 0}
                                                                        _ = v73
                                                                    case 1:
                                                                        var v74 uint16 = 32767
                                                                        _ = v74
                                                                    default:
                                                                        var v75 int64 = (v17 & 513)
                                                                        _ = v75
                                                                    }
                                                                case 829:
                                                                    if (255 != 2) {
                                                                        for fi76 := 0; fi76 < 10; fi76++ {
                                                                            switch "a" {
                                                                            case "\n\t":
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    var v77 int = (v4 ^ (fi61 << ((18446744073709551615) & 63)))
                                                                                    _ = v77
                                                                                default:
                                                                                    v78 := GenS8{F0: 312}
                                                                                    _ = v78
                                                                                }
                                                                            }
                                                                            _ = fi76
                                                                        }
                                                                        if v20 {
                                                                            var v79 int32 = 2147483647
                                                                            _ = v79
                                                                        }
                                                                        v80 := GenS9{F0: 524, F1: 255, F2: 0}
                                                                        _ = v80
                                                                    } else {
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            if v20 {
                                                                                v81 := GenS8{F0: 65535}
                                                                                var v82 bool = true
                                                                                _ = v81
                                                                                _ = v82
                                                                            }
                                                                        case "foo bar":
                                                                            switch v4 {
                                                                            case -1:
                                                                                var v83 int = (v4 + v4)
                                                                                _ = v83
                                                                            case -604:
                                                                                fmt.Println(v3)
                                                                            case -9223372036854775808:
                                                                                var v84 uint64 = 260
                                                                                _ = v84
                                                                            }
                                                                        default:
                                                                            _ = len(v5)
                                                                        }
                                                                        v20 = v20
                                                                    }
                                                                case -9223372036854775808:
                                                                    if (v1 > (v26 + v26)) {
                                                                        for fi85 := 0; fi85 < 8; fi85++ {
                                                                            var v86 uint = 29
                                                                            v3 = 3.14159
                                                                            v2 -= int16(v16)
                                                                            if (0.1 != (1.5 * 0.0)) {
                                                                                if v20 {
                                                                                    var v87 float64 = 3.14159
                                                                                    _ = v87
                                                                                } else {
                                                                                    v2 = genFn0(2, (fi85 | (fi85 & 42)))
                                                                                    v20 = v20
                                                                                }
                                                                                var v88 float32 = 1e10
                                                                                _ = v88
                                                                            }
                                                                            v26 -= 42
                                                                            _ = fi85
                                                                            _ = v86
                                                                        }
                                                                        v4 -= v4
                                                                        var v89 uint64 = 18446744073709551615
                                                                        _ = v89
                                                                    } else {
                                                                        fmt.Println(fi61)
                                                                        if v20 {
                                                                            var v90 int64 = (v16 >> ((42) & 63))
                                                                            if v20 {
                                                                                var v91 int32 = -723
                                                                                _ = v91
                                                                            }
                                                                            _ = v90
                                                                        }
                                                                        v17 = (v17 ^ v16)
                                                                    }
                                                                default:
                                                                    if len(v5) > 0 {
                                                                        v5[0] = 13
                                                                    }
                                                                }
                                                            }
                                                            var v92 float64 = 1.5
                                                            for fi93 := 0; fi93 < 5; fi93++ {
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                v94 := GenS7{F0: 0, F1: ((v20 && (1 != -1)) && v20), F2: 433}
                                                                _ = fi93
                                                                _ = v94
                                                            }
                                                            _ = v92
                                                        } else {
                                                            var v95 int64 = v16
                                                            fmt.Println(v4, v1)
                                                            v3 -= -2.25
                                                            _ = v95
                                                        }
                                                    } else {
                                                        fmt.Println(v16, v4, v26)
                                                    }
                                                    for fi96 := 0; fi96 < 11; fi96++ {
                                                        v1 = (v26 / (((v26 % ((573) | 1))) | 1))
                                                        _ = fi96
                                                    }
                                                }
                                                for fi97 := 0; fi97 < 16; fi97++ {
                                                    var v98 int16 = int16(v16)
                                                    v2 -= v98
                                                    _ = fi97
                                                    _ = v98
                                                }
                                                v99 := GenS2{F0: v20, F1: 2147483647}
                                                _ = v99
                                            }
                                            var v100 int = v4
                                            _ = v100
                                        }
                                        v101 := GenS10{F0: (v16 + (-(v17)))}
                                        _ = v101
                                    } else {
                                        v102 := GenS7{F0: int8(v16), F1: false, F2: 637}
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = v102
                                    }
                                    _ = fi61
                                }
                            }
                        }
                    case 32768:
                        var v103 string = "foo bar"
                        _ = v103
                    default:
                        v16 = (v16 << ((248) & 63))
                    }
                    v20 = v20
                }
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        v2 = (v2 & v2)
        for fi104 := 0; fi104 < 20; fi104++ {
            var v105 uint64 = (v26 & v26)
            for fi106 := 0; fi106 < 12; fi106++ {
                var v107 uint8 = 0
                _ = fi106
                _ = v107
            }
            _ = fi104
            _ = v105
        }
    } else {
        v108 := GenS11{F0: 1, F1: v20, F2: -58}
        _ = v108
    }
    for fi109 := 0; fi109 < 3; fi109++ {
        if false {
            v20 = true
            v2 = -32768
            v17 -= (-(v16))
        }
        for fi110 := 0; fi110 < 7; fi110++ {
            v17 -= (^(v16))
            _ = fi110
        }
        _ = fi109
    }
    for fi111 := 0; fi111 < 11; fi111++ {
        v17 += 465
        fmt.Println(v20, v2, v17)
        v17 = (v17 % ((-490) | 1))
        _ = fi111
    }
    if v20 {
        _ = len(v5)
        v112 := GenS12{F0: 0, F1: -2147483648, F2: 0}
        if v20 {
            for fi113 := 0; fi113 < 8; fi113++ {
                if v20 {
                    v3 -= v3
                    v3 = v3
                    v5 = append(v5, 42)
                } else {
                    var v114 float32 = v3
                    if v20 {
                        var v115 int16 = (v2 * v2)
                        v3 += v3
                        if len(v5) > 0 {
                            v5[0] = 65535
                        }
                        _ = v115
                    } else {
                        v20 = (42 <= 1)
                        switch ("hello" + "\n\t") {
                        case "":
                            if (!(0 == 6)) {
                                var v116 int32 = 2147483647
                                var v117 uint32 = genFn1(65535, v1)
                                _ = v116
                                _ = v117
                            }
                        default:
                            var v118 float64 = (3.14159 - (0.1 / (1.5 - (1.5 + -2.25))))
                            _ = v118
                        }
                        var v119 bool = (!(v2 < genFn0(1, (fi113 + (fi113 >> ((0) & 63))))))
                        _ = v119
                    }
                    var v120 int64 = -9223372036854775808
                    _ = v114
                    _ = v120
                }
                _ = len(v5)
                if v20 {
                    switch -874 {
                    case 0:
                        if v20 {
                            if ((true && false) || false) {
                                var v121 int = v4
                                _ = v121
                            } else {
                                for fi122 := 0; fi122 < 4; fi122++ {
                                    if len(v5) > 0 {
                                        v5[0] = 752
                                    }
                                    fmt.Println(v17, v3, v17)
                                    _ = fi122
                                }
                                if len(v5) > 0 {
                                    v5[0] = 156
                                }
                                switch fi113 {
                                case -78:
                                    var v123 uint = 4294967295
                                    _ = v123
                                }
                            }
                            var v124 int8 = -21
                            _ = v124
                        }
                    case 571:
                        switch "a" {
                        case "foo bar":
                            var v125 float32 = v3
                            _ = v125
                        case "\n\t":
                            var v126 uint32 = 42
                            _ = v126
                        case "a":
                            v17 = (v17 << ((1) & 63))
                        }
                    }
                    v127 := GenS13{F0: 65535, F1: 42, F2: uint32(v1)}
                    _ = v127
                }
                _ = fi113
            }
        } else {
            v3 = (-2.25 + 3.14159)
            v4 -= (v4 ^ (v4 - 117))
        }
        _ = v112
    } else {
        fmt.Println(v17)
    }
    v128 := GenS3{F0: -1, F1: v1, F2: -1}
    var v129 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v129:", v129)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v128
    _ = v129
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
