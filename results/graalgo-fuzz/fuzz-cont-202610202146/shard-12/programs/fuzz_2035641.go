package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32; F1 uint8 }
type GenS8 struct { F0 string; F1 uint8 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int64; F1 uint64 }
type GenS11 struct { F0 uint; F1 uint }
type GenS12 struct { F0 uint64; F1 uint16 }

func main() {
    var v1 uint8 = 128
    v2 := []uint32{0, 912, 2, 4294967295}
    var v3 int16 = 42
    var v4 int = 366
    var v5 string = (("hello" + (("foo bar" + "foo bar") + "")) + "hello")
    v3 -= (^(v3))
    v1 += 164
    v6 := GenS1{F0: v5, F1: 42, F2: 42}
    var v7 uint64 = 0
    var v8 string = v5
    v9 := GenS1{F0: v5, F1: 417, F2: 65535}
    v8 = v8
    var v10 uint64 = (v7 * 42)
    v11 := GenS1{F0: v8, F1: -1, F2: 569}
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v3 = (^(v3))
        switch "foo bar" {
        case "a":
            var v12 uint64 = v7
            _ = v12
        case "foo bar":
            switch (v4 * (v4 + -333)) {
            case 256:
                for fi13 := 0; fi13 < 14; fi13++ {
                    var v14 float32 = 0.1
                    if (18446744073709551615 == 2) {
                        switch "foo bar" {
                        case "":
                            _ = len(v2)
                        case "a":
                            if len(v2) > 0 {
                                v2[0] = 2147483648
                            }
                        case "\n\t":
                            v10 -= v10
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v1 += (v1 + 128)
                    } else {
                        if (1 > 0) {
                            var v15 uint64 = 4294967295
                            var v16 float64 = 1e10
                            v14 = v14
                            _ = v15
                            _ = v16
                        }
                        var v17 uint16 = 65535
                        _ = v17
                    }
                    if true {
                        v2 = append(v2, 2147483648)
                        var v18 float64 = (3.14159 + 0.1)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v18
                    }
                    _ = fi13
                    _ = v14
                }
            case 42:
                _ = len(v2)
            case -1:
                if len(v2) > 0 {
                    v2[0] = 346
                }
            }
        case "hello":
            v3 -= (v3 * v3)
        default:
            if true {
                switch (v8 + v8) {
                case "\n\t":
                    fmt.Println(v8, v10, v10)
                }
                for fi19 := 0; fi19 < 11; fi19++ {
                    v2 = append(v2, 15)
                    _ = fi19
                }
            } else {
                v4 += v4
                v8 = ("\n\t" + v5)
            }
        }
    }
    switch 1 {
    case 127:
        _ = len(v2)
    case 592:
        var v20 uint32 = 0
        _ = v20
    }
    if true {
        fmt.Println(v10, v4)
        var v21 int32 = 42
        _ = v21
    }
    switch (v4 * v4) {
    case 2147483647:
        var v22 bool = (true && (v3 < 32767))
        _ = v22
    }
    if true {
        switch 1 {
        case -9223372036854775808:
            if (!(v1 <= v1)) {
                var v23 uint8 = v1
                for fi24 := 0; fi24 < 18; fi24++ {
                    for fi25 := 0; fi25 < 11; fi25++ {
                        var v26 uint16 = uint16(v23)
                        v27 := GenS1{F0: "a", F1: -2147483648, F2: (^(v26))}
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    if len(v2) > 0 {
                        v2[0] = 2
                    }
                    var v28 int32 = -743
                    _ = fi24
                    _ = v28
                }
                _ = v23
            } else {
                switch (v8 + ("foo bar" + ("\n\t" + v8))) {
                case "\n\t":
                    switch "hello" {
                    case "foo bar":
                        v2 = append(v2, 2147483647)
                    case "hello":
                        v29 := GenS1{F0: "foo bar", F1: int32(v3), F2: 924}
                        _ = v29
                    }
                case "a":
                    v3 = v3
                default:
                    fmt.Println(v5, v5)
                }
                if true {
                    switch v8 {
                    case "\n\t":
                        switch v5 {
                        case "\n\t":
                            switch v5 {
                            case "":
                                if len(v2) > 0 {
                                    v2[0] = 42
                                }
                            case "foo bar":
                                switch v4 {
                                case 1:
                                    if (v5 > v5) {
                                        if false {
                                            var v30 int = 32767
                                            fmt.Println(v1)
                                            var v31 uint64 = v7
                                            _ = v30
                                            _ = v31
                                        } else {
                                            fmt.Println(v4, v3)
                                            for fi32 := 0; fi32 < 8; fi32++ {
                                                v10 -= (v7 * v10)
                                                var v33 int64 = 313
                                                for fi34 := 0; fi34 < 18; fi34++ {
                                                    _ = len(v2)
                                                    _ = fi34
                                                }
                                                _ = fi32
                                                _ = v33
                                            }
                                            if false {
                                                v3 = 1
                                                if true {
                                                    var v35 int32 = -33
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    _ = v35
                                                }
                                            }
                                        }
                                        if true {
                                            switch v4 {
                                            case 42:
                                                v4 = -1
                                            case 404:
                                                v2 = append(v2, 141)
                                            default:
                                                for fi36 := 0; fi36 < 18; fi36++ {
                                                    v10 = (^(v7))
                                                    v5 = v5
                                                    if (65535 > 1) {
                                                        for fi37 := 0; fi37 < 10; fi37++ {
                                                            if false {
                                                                switch "hello" {
                                                                case "a":
                                                                    if (-6 != 8) {
                                                                        var v38 int32 = -59
                                                                        if ((!true) || (1 == 65535)) {
                                                                            if (v10 <= (^(v7))) {
                                                                                if (1 >= 42) {
                                                                                    if false {
                                                                                        if (757 < v38) {
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                            v2 = append(v2, 2)
                                                                                            if ((370 != 573) && false) {
                                                                                                v4 -= 2
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                                v39 := GenS1{F0: (("" + v5) + "a"), F1: (v38 << ((435) & 63)), F2: 689}
                                                                                                _ = v39
                                                                                            }
                                                                                            v40 := GenS2{F0: (v1 >> ((905) & 63)), F1: 855, F2: 92}
                                                                                            fmt.Println(v7, v38)
                                                                                            _ = v40
                                                                                        } else {
                                                                                            v41 := GenS3{F0: 326, F1: (v38 + (^(v38)))}
                                                                                            v3 = (v3 / (((v3 >> ((806) & 63))) | 1))
                                                                                            fmt.Println(v7)
                                                                                            v1 = (^(v1))
                                                                                            var v42 string = v8
                                                                                            _ = v41
                                                                                            _ = v42
                                                                                        }
                                                                                        v38 = v38
                                                                                        var v43 uint64 = v7
                                                                                        _ = v43
                                                                                    }
                                                                                    if (v8 < (v8 + (v8 + v8))) {
                                                                                        var v44 uint16 = 76
                                                                                        v45 := GenS3{F0: 2, F1: v38}
                                                                                        v46 := GenS4{F0: -53}
                                                                                        _ = v44
                                                                                        _ = v45
                                                                                        _ = v46
                                                                                    }
                                                                                    fmt.Println(v38, fi37, v7)
                                                                                }
                                                                            } else {
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                var v47 float32 = float32(float64((0.1 * (-2.25 * 0.1))))
                                                                                _ = v47
                                                                            }
                                                                            var v48 int16 = v3
                                                                            if ((true || (844 <= -1)) || (0.1 < 1.5)) {
                                                                                if false {
                                                                                    if ((-1 == 1) || true) {
                                                                                        v1 = v1
                                                                                        fmt.Println(fi37, v10)
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                var v49 uint = 986
                                                                                v7 += (v10 + uint64(v1))
                                                                                _ = v49
                                                                            }
                                                                            _ = v48
                                                                        }
                                                                        _ = v38
                                                                    } else {
                                                                        fmt.Println(v1, v7, v5)
                                                                    }
                                                                case "foo bar":
                                                                    v3 = 0
                                                                case "":
                                                                    _ = len(v2)
                                                                default:
                                                                    var v50 int = (-(fi37))
                                                                    _ = v50
                                                                }
                                                            } else {
                                                                v3 = (v3 + (v3 ^ v3))
                                                                _ = len(v2)
                                                                if false {
                                                                    var v51 bool = true
                                                                    _ = v51
                                                                }
                                                            }
                                                            _ = fi37
                                                        }
                                                        fmt.Println(v10, v3, v3)
                                                    }
                                                    _ = fi36
                                                }
                                            }
                                            v10 -= v7
                                        } else {
                                            var v52 int64 = 42
                                            switch (((("\n\t" + v5) + v8) + v5) + v5) {
                                            case "a":
                                                if len(v2) > 0 {
                                                    v2[0] = 815
                                                }
                                            case "hello":
                                                if (!((v3 ^ v3) < (v3 >> ((1) & 63)))) {
                                                    if true {
                                                        v4 = (v4 - 9223372036854775807)
                                                    } else {
                                                        v53 := GenS5{F0: (v10 + 2), F1: 42}
                                                        var v54 int = (^(v4))
                                                        var v55 int16 = v3
                                                        _ = v53
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                    for fi56 := 0; fi56 < 3; fi56++ {
                                                        v8 = v5
                                                        _ = fi56
                                                    }
                                                    var v57 float32 = ((3.14159 / (3.14159 * float32(1.5))) / (float32(-2.25) / ((3.14159 / 3.14159) - 1e10)))
                                                    _ = v57
                                                } else {
                                                    v2 = append(v2, 4294967295)
                                                    var v58 float64 = (((1e10 / 1.5) - float64(0.1)) * -2.25)
                                                    fmt.Println(v52)
                                                    _ = v58
                                                }
                                            }
                                            _ = v52
                                        }
                                        var v59 int16 = v3
                                        _ = v59
                                    } else {
                                        v2 = append(v2, 0)
                                        if false {
                                            for fi60 := 0; fi60 < 7; fi60++ {
                                                v2 = append(v2, 2147483648)
                                                _ = fi60
                                            }
                                            var v61 string = ("foo bar" + "hello")
                                            _ = v61
                                        } else {
                                            for fi62 := 0; fi62 < 8; fi62++ {
                                                v63 := GenS3{F0: 332, F1: 2147483647}
                                                _ = fi62
                                                _ = v63
                                            }
                                        }
                                        switch "hello" {
                                        case "":
                                            v8 = ("foo bar" + (((v8 + (v8 + v8)) + "hello") + v8))
                                        case "foo bar":
                                            var v64 string = ("\n\t" + v5)
                                            _ = v64
                                        case "hello":
                                            var v65 int64 = 0
                                            _ = v65
                                        }
                                    }
                                default:
                                    v7 += v7
                                }
                            case "\n\t":
                                v66 := GenS6{F0: -108}
                                _ = v66
                            }
                        case "a":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            fmt.Println(v8, v10, v1)
                        }
                    case "":
                        if ((0.1 + 1.5) >= ((0.0 + 0.1) / float64(1e10))) {
                            v2 = append(v2, 944)
                            for fi67 := 0; fi67 < 16; fi67++ {
                                v68 := GenS6{F0: 0}
                                v69 := GenS5{F0: (v7 | v10), F1: 100}
                                var v70 int32 = -441
                                _ = fi67
                                _ = v68
                                _ = v69
                                _ = v70
                            }
                        }
                    case "foo bar":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        switch "foo bar" {
                        case "a":
                            switch 339 {
                            case 304:
                                v2 = append(v2, 4294967295)
                            case 4294967295:
                                fmt.Println(v10, v4, v1)
                            case 1:
                                fmt.Println(v10)
                            }
                        case "":
                            fmt.Println(v5)
                        case "hello":
                            var v71 uint8 = v1
                            _ = v71
                        }
                    }
                    v2 = append(v2, 42)
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                }
            }
        case -1:
            v4 -= v4
        case -779:
            for fi72 := 0; fi72 < 17; fi72++ {
                for fi73 := 0; fi73 < 8; fi73++ {
                    v2 = append(v2, 42)
                    fmt.Println(v10, v1)
                    _ = fi73
                }
                _ = fi72
            }
        default:
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
        }
        if true {
            for fi74 := 0; fi74 < 16; fi74++ {
                v75 := GenS6{F0: 0}
                v3 -= (^(v3))
                for fi76 := 0; fi76 < 19; fi76++ {
                    v77 := GenS1{F0: (v8 + ("foo bar" + (v5 + v8))), F1: -679, F2: uint16(v7)}
                    switch fi74 {
                    case 32767:
                        var v78 float64 = float64((-2.25 * 1.5))
                        _ = v78
                    default:
                        v8 = "\n\t"
                    }
                    var v79 int32 = 881
                    switch fi76 {
                    case -2147483648:
                        v80 := GenS7{F0: 745, F1: (v1 & (v1 * v1))}
                        _ = v80
                    case -9223372036854775808:
                        v7 += (^(v10))
                    case 1:
                        fmt.Println(v3, v79, v3)
                    default:
                        switch (("\n\t" + v5) + v5) {
                        case "":
                            v3 -= v3
                        default:
                            v81 := GenS8{F0: v8, F1: (v1 * (v1 / ((v1) | 1)))}
                            _ = v81
                        }
                    }
                    _ = fi76
                    _ = v77
                    _ = v79
                }
                _ = fi74
                _ = v75
            }
            v2 = append(v2, 4294967295)
        } else {
            if (-128 > 127) {
                var v82 bool = (false || true)
                _ = v82
            } else {
                var v83 uint16 = 255
                var v84 float32 = (0.1 + 0.1)
                v85 := GenS7{F0: 42, F1: 0}
                _ = v83
                _ = v84
                _ = v85
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        for fi86 := 0; fi86 < 5; fi86++ {
            for fi87 := 0; fi87 < 10; fi87++ {
                v88 := GenS4{F0: -128}
                if false {
                    var v89 float32 = (1e10 + ((0.1 / 1e10) + 0.1))
                    switch v4 {
                    case 876:
                        var v90 int16 = 1000
                        _ = v90
                    case -9223372036854775808:
                        var v91 uint = 42
                        _ = v91
                    default:
                        var v92 uint8 = (v1 << ((65535) & 63))
                        _ = v92
                    }
                    _ = v89
                } else {
                    var v93 uint32 = 595
                    if len(v2) > 0 {
                        v2[0] = (v93 + (v93 + v93))
                    }
                    _ = v93
                }
                switch "foo bar" {
                case "\n\t":
                    var v94 float64 = 1e10
                    _ = v94
                case "hello":
                    fmt.Println(v4)
                case "foo bar":
                    v3 = (v3 - 32767)
                }
                _ = fi87
                _ = v88
            }
            for fi95 := 0; fi95 < 3; fi95++ {
                fmt.Println(v7)
                if false {
                    v96 := GenS7{F0: 0, F1: v1}
                    var v97 int = v4
                    switch 42 {
                    case 4294967295:
                        v98 := GenS1{F0: "hello", F1: -588, F2: 255}
                        v8 = v8
                        if true {
                            if len(v2) > 0 {
                                v2[0] = 2
                            }
                            v2 = append(v2, 42)
                        }
                        _ = v98
                    case -238:
                        v1 -= v1
                    case 32768:
                        v99 := GenS4{F0: 11}
                        _ = v99
                    }
                    _ = v96
                    _ = v97
                } else {
                    v100 := GenS9{F0: "\n\t"}
                    v101 := GenS8{F0: (v8 + v8), F1: v1}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v100
                    _ = v101
                }
                if false {
                    v4 = fi86
                }
                _ = fi95
            }
            _ = fi86
        }
    } else {
        for fi102 := 0; fi102 < 5; fi102++ {
            switch fi102 {
            case 9223372036854775807:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 2147483647:
                for fi103 := 0; fi103 < 7; fi103++ {
                    v4 += v4
                    _ = fi103
                }
            case 0:
                switch v4 {
                case 292:
                    for fi104 := 0; fi104 < 1; fi104++ {
                        var v105 uint32 = 2
                        _ = len(v2)
                        v106 := GenS2{F0: (v1 << ((780) & 63)), F1: 1, F2: v1}
                        v4 = (^(fi102))
                        if (!true) {
                            switch fi104 {
                            case 1:
                                if true {
                                    var v107 uint = 1
                                    v108 := GenS6{F0: 625}
                                    _ = v107
                                    _ = v108
                                } else {
                                    var v109 int = fi102
                                    if (false || ((float32(1.5) >= 1.5) && ("foo bar" == v5))) {
                                        if (true && (0 < v3)) {
                                            if false {
                                                v110 := GenS2{F0: v1, F1: 591, F2: v1}
                                                var v111 uint64 = v7
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v109 = 632
                                                var v112 int8 = 1
                                                _ = v110
                                                _ = v111
                                                _ = v112
                                            } else {
                                                v109 += (v109 * v4)
                                                var v113 int = fi102
                                                var v114 float64 = 1e10
                                                _ = v113
                                                _ = v114
                                            }
                                            _ = len(v2)
                                        }
                                        var v115 int32 = -1
                                        var v116 float64 = 3.14159
                                        _ = v115
                                        _ = v116
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = (v105 ^ 924)
                                        }
                                        _ = len(v2)
                                        var v117 uint64 = v7
                                        _ = v117
                                    }
                                    _ = v109
                                }
                            case 944:
                                v118 := GenS10{F0: 1, F1: (v7 % (((^(v10))) | 1))}
                                _ = v118
                            case -9223372036854775808:
                                if len(v2) > 0 {
                                    v2[0] = (v105 + 4294967295)
                                }
                            }
                        } else {
                            v119 := GenS11{F0: 42, F1: 18446744073709551615}
                            _ = v119
                        }
                        _ = fi104
                        _ = v105
                        _ = v106
                    }
                default:
                    if ((32767 > 2) || false) {
                        v5 = "a"
                        _ = len(v2)
                        _ = len(v2)
                    }
                }
            }
            switch "" {
            case "foo bar":
                v120 := GenS8{F0: (v8 + ("" + "a")), F1: (v1 ^ (v1 - v1))}
                _ = v120
            case "\n\t":
                v121 := GenS12{F0: v10, F1: 32767}
                _ = v121
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi102
        }
    }
    v122 := GenS4{F0: 52}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v122
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
