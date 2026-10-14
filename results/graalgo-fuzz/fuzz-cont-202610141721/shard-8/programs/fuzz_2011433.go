package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 uint8; F1 int16 }
type GenS7 struct { F0 int8; F1 int64 }
type GenS8 struct { F0 int8; F1 string }
type GenS9 struct { F0 uint8; F1 uint64; F2 bool }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int16; F1 string; F2 int }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint8; F1 int64 }
type GenS14 struct { F0 int64; F1 uint }
func genFn0(p0 int16, p1 int) uint64 {
    if (!false) {
        fmt.Println(p1, p1)
        var v7 int = (p1 | p1)
        _ = v7
    }
    p1 = p1
    var v8 uint32 = 1
    v9 := GenS1{F0: 4294967295, F1: 884, F2: 30}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 437
}

func main() {
    v1 := []int{0, 128}
    var v2 uint32 = 332
    v3 := []uint32{(v2 << ((998) & 63)), (v2 >> ((984) & 63))}
    v4 := make(map[int]uint32)
    var v5 int64 = -9223372036854775808
    var v6 int = -150
    if true {
        v2 += 635
        var v10 uint8 = 2
        _ = v10
    } else {
        var v11 uint8 = 127
        _ = v11
    }
    if (!true) {
        _ = len(v4)
    } else {
        v12 := GenS2{F0: int64(v2)}
        _ = v12
    }
    var v13 uint16 = 186
    if (true || false) {
        if len(v3) > 0 {
            v3[0] = (v2 & (^(v2)))
        }
        v4[(v6 % ((128) | 1))] = v2
    } else {
        for fi14 := 0; fi14 < 2; fi14++ {
            v5 += v5
            for fi15 := 0; fi15 < 10; fi15++ {
                switch v6 {
                case -2147483648:
                    _ = len(v4)
                case 0:
                    var v16 int16 = -1
                    _ = v16
                }
                v17 := GenS3{F0: 2147483647}
                _ = fi15
                _ = v17
            }
            switch int(v13) {
            case 472:
                v18 := GenS4{F0: 2}
                _ = v18
            default:
                if false {
                    var v19 uint = 765
                    _ = v19
                }
            }
            _ = fi14
        }
    }
    if true {
        if ((false && false) || (42 <= -89)) {
            var v20 float64 = 3.14159
            var v21 uint8 = 1
            var v22 uint16 = v13
            _ = v20
            _ = v21
            _ = v22
        }
        if (!false) {
            for fi23 := 0; fi23 < 10; fi23++ {
                if ((v5 + (^(v5))) != (v5 / (((v5 * (v5 | (v5 * v5)))) | 1))) {
                    var v24 int8 = 76
                    if (!(!false)) {
                        switch "hello" {
                        case "a":
                            if false {
                                for fi25 := 0; fi25 < 12; fi25++ {
                                    var v26 int = v6
                                    _ = fi25
                                    _ = v26
                                }
                                v27 := GenS2{F0: v5}
                                v28 := GenS5{F0: 1, F1: -602}
                                _ = v27
                                _ = v28
                            }
                        case "":
                            var v29 bool = (0 > 65535)
                            _ = v29
                        case "\n\t":
                            if true {
                                v13 -= v13
                                for fi30 := 0; fi30 < 14; fi30++ {
                                    v31 := GenS6{F0: 72, F1: -33}
                                    if false {
                                        delete(v4, fi30)
                                    }
                                    if true {
                                        var v32 uint16 = 0
                                        v6 = -1
                                        if ("foo bar" != ("hello" + (("foo bar" + (("" + ("hello" + "a")) + ("foo bar" + "a"))) + ""))) {
                                            switch -1 {
                                            case 1:
                                                v3 = append(v3, v2)
                                            case 42:
                                                v4[(-(fi23))] = (v2 / ((215) | 1))
                                                var v33 string = "foo bar"
                                                var v34 int64 = v5
                                                _ = v33
                                                _ = v34
                                            }
                                            v35 := GenS6{F0: 42, F1: 1000}
                                            v32 = v13
                                            _ = v35
                                        }
                                        _ = v32
                                    }
                                    _ = fi30
                                    _ = v31
                                }
                                var v36 string = "a"
                                _ = v36
                            }
                        }
                    } else {
                        for fi37 := 0; fi37 < 7; fi37++ {
                            if false {
                                var v38 int8 = (v24 | 112)
                                if (255 != (v13 + uint16(v2))) {
                                    fmt.Println(v2, v24)
                                    _ = len(v1)
                                    var v39 uint8 = 0
                                    _ = v39
                                }
                                if (false || true) {
                                    var v40 string = "foo bar"
                                    v2 = (v2 >> ((255) & 63))
                                    v24 = -87
                                    _ = v40
                                }
                                fmt.Println(v38)
                                var v41 string = (("\n\t" + "") + "foo bar")
                                _ = v38
                                _ = v41
                            }
                            var v42 uint16 = v13
                            v3 = append(v3, (v2 / ((v2) | 1)))
                            _ = fi37
                            _ = v42
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = v24
                } else {
                    if (255 != 737) {
                        v13 = (v13 / ((v13) | 1))
                        if len(v1) > 0 {
                            v1[0] = (fi23 / ((fi23) | 1))
                        }
                        if true {
                            if ((596 > 42) || true) {
                                switch 32767 {
                                case -938:
                                    fmt.Println(v5)
                                case 0:
                                    v43 := GenS4{F0: 4294967295}
                                    _ = v43
                                case 127:
                                    var v44 bool = (!false)
                                    _ = v44
                                }
                            }
                        }
                    } else {
                        for fi45 := 0; fi45 < 19; fi45++ {
                            v13 += (v13 >> ((uint(v13)) & 63))
                            var v46 uint = 851
                            if true {
                                v5 = (v5 - v5)
                                v47 := GenS4{F0: 18446744073709551615}
                                v5 += (-(v5))
                                _ = v47
                            } else {
                                var v48 uint16 = 32767
                                var v49 string = "a"
                                var v50 int = 127
                                _ = v48
                                _ = v49
                                _ = v50
                            }
                            _ = fi45
                            _ = v46
                        }
                        var v51 uint64 = 1
                        var v52 uint32 = v2
                        _ = v51
                        _ = v52
                    }
                }
                var v53 int16 = -1
                fmt.Println(v53, v13, v6)
                _ = fi23
                _ = v53
            }
        } else {
            if false {
                if ("a" > ("" + ((("\n\t" + ("hello" + ("" + (((("foo bar" + "hello") + "hello") + ("\n\t" + "a")) + "")))) + ("a" + ("foo bar" + "\n\t"))) + "foo bar"))) {
                    switch "hello" {
                    case "\n\t":
                        v6 -= -383
                    case "foo bar":
                        switch -9223372036854775808 {
                        case 977:
                            if false {
                                v54 := GenS1{F0: 2, F1: (v2 + (^(v2))), F2: uint8(v13)}
                                if true {
                                    v2 = v2
                                } else {
                                    for fi55 := 0; fi55 < 15; fi55++ {
                                        v13 += (^(v13))
                                        _ = fi55
                                    }
                                    v56 := v4[(v6 | 4294967295)]
                                    fmt.Println(v2)
                                    _ = v56
                                }
                                _ = v54
                            }
                        case -575:
                            if true {
                                switch v6 {
                                case 42:
                                    for fi57 := 0; fi57 < 3; fi57++ {
                                        var v58 int16 = -32768
                                        v58 += int16(v13)
                                        _ = fi57
                                        _ = v58
                                    }
                                case -1:
                                    _ = len(v4)
                                }
                                for fi59 := 0; fi59 < 12; fi59++ {
                                    var v60 uint16 = 0
                                    _ = fi59
                                    _ = v60
                                }
                            } else {
                                switch "\n\t" {
                                case "hello":
                                    var v61 uint32 = v2
                                    _ = v61
                                default:
                                    v5 += v5
                                }
                                v3 = append(v3, v2)
                                if true {
                                    v6 += (v6 << ((226) & 63))
                                    if true {
                                        _ = len(v1)
                                    }
                                    if false {
                                        var v62 uint32 = (v2 >> ((4294967295) & 63))
                                        _ = v62
                                    }
                                }
                            }
                        }
                    }
                    if (((("hello" + (("foo bar" + "\n\t") + "\n\t")) + "") + "a") <= "hello") {
                        switch v6 {
                        case -229:
                            v6 = (^(v6))
                        case 1:
                            var v63 int32 = 42
                            _ = v63
                        default:
                            for fi64 := 0; fi64 < 1; fi64++ {
                                var v65 uint = 953
                                v66 := GenS4{F0: 18446744073709551615}
                                switch (fi64 * (fi64 & -155)) {
                                case 25:
                                    v67 := GenS5{F0: 2, F1: int32(v13)}
                                    _ = v67
                                case 127:
                                    v6 -= 42
                                case 32768:
                                    v68 := GenS5{F0: 42, F1: -1}
                                    _ = v68
                                default:
                                    for fi69 := 0; fi69 < 18; fi69++ {
                                        if (((!true) && false) || true) {
                                            v70 := GenS7{F0: 100, F1: (-(v5))}
                                            fmt.Println(fi64, v65)
                                            v4[(-(v6))] = v2
                                            _ = v70
                                        } else {
                                            v5 -= (v5 ^ (-(v5)))
                                            var v71 uint = v65
                                            _ = v71
                                        }
                                        v1 = append(v1, int(v13))
                                        _ = fi69
                                    }
                                }
                                _ = fi64
                                _ = v65
                                _ = v66
                            }
                        }
                        v3 = append(v3, 0)
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    if (((^(v5)) == v5) && (526 > 899)) {
                        v72 := GenS7{F0: 1, F1: v5}
                        for fi73 := 0; fi73 < 13; fi73++ {
                            if false {
                                switch "foo bar" {
                                case "foo bar":
                                    v74 := GenS2{F0: (v5 % ((v5) | 1))}
                                    _ = v74
                                case "":
                                    var v75 int = (v6 >> ((0) & 63))
                                    _ = v75
                                case "\n\t":
                                    switch "a" {
                                    case "\n\t":
                                        var v76 float32 = 3.14159
                                        _ = v76
                                    case "foo bar":
                                        v6 += fi73
                                    case "a":
                                        if false {
                                            var v77 int64 = (v5 | -167)
                                            _ = v77
                                        }
                                    }
                                }
                                fmt.Println(v5, v2, v2)
                                var v78 int = fi73
                                _ = v78
                            }
                            v79 := GenS4{F0: 42}
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi73
                            _ = v79
                        }
                        _ = v72
                    } else {
                        if false {
                            var v80 uint16 = v13
                            fmt.Println(v2, v2, v80)
                            fmt.Println(v6)
                            _ = v80
                        } else {
                            var v81 uint32 = 2147483647
                            v6 += (v6 & v6)
                            _ = len(v4)
                            _ = v81
                        }
                        v6 -= (^(v6))
                        if true {
                            v3 = append(v3, (^(v2)))
                            var v82 uint = 65535
                            _ = v82
                        }
                    }
                    v2 -= v2
                    for fi83 := 0; fi83 < 14; fi83++ {
                        var v84 float64 = ((1e10 * (0.0 + 0.0)) - ((1.5 * (1e10 + (1e10 - -2.25))) - ((float64(1e10) * float64(1e10)) * ((((-2.25 / 1.5) + 0.1) / 1.0) - (3.14159 - 1.5)))))
                        if true {
                            _ = len(v3)
                            v84 += v84
                        } else {
                            for fi85 := 0; fi85 < 12; fi85++ {
                                v86 := GenS2{F0: v5}
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = fi85
                                _ = v86
                            }
                            fmt.Println(v84)
                            _ = len(v1)
                        }
                        _ = fi83
                        _ = v84
                    }
                }
                var v87 float32 = (1e10 + -2.25)
                switch "foo bar" {
                case "":
                    if (uint64(v2) <= 1) {
                        switch 0 {
                        case -816:
                            if (184 < 149) {
                                switch "a" {
                                case "hello":
                                    fmt.Println(v87, v5)
                                case "\n\t":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case "":
                                    if true {
                                        v87 = 3.14159
                                        v3 = append(v3, (v2 + v2))
                                    } else {
                                        if ((0.0 + (float64(v87) + ((0.1 / (0.0 - -2.25)) + 1e10))) == 3.14159) {
                                            v6 += 128
                                            for fi88 := 0; fi88 < 20; fi88++ {
                                                var v89 string = "a"
                                                v87 -= v87
                                                _ = fi88
                                                _ = v89
                                            }
                                            for fi90 := 0; fi90 < 18; fi90++ {
                                                v13 = v13
                                                if ((("a" + "a") > ("hello" + "\n\t")) || (3.14159 == 0.1)) {
                                                    if false {
                                                        fmt.Println(v13, v2)
                                                        fmt.Println(v2, fi90)
                                                    }
                                                    v5 = v5
                                                } else {
                                                    for fi91 := 0; fi91 < 17; fi91++ {
                                                        fmt.Println(v5)
                                                        v3 = append(v3, v2)
                                                        v92 := GenS8{F0: 100, F1: ("" + "hello")}
                                                        _ = fi91
                                                        _ = v92
                                                    }
                                                    if (42 >= genFn0(822, 65)) {
                                                        v2 = (v2 / ((v2) | 1))
                                                    } else {
                                                        if ((v2 & (v2 >> ((65535) & 63))) > 42) {
                                                            var v93 int16 = -1
                                                            v1 = append(v1, (v6 >> ((1) & 63)))
                                                            var v94 string = "hello"
                                                            _ = v93
                                                            _ = v94
                                                        }
                                                        fmt.Println(v87, v6, v13)
                                                    }
                                                }
                                                v6 += -423
                                                _ = fi90
                                            }
                                        } else {
                                            v13 -= v13
                                            if (((v5 - -9223372036854775808) >= v5) && false) {
                                                v95 := GenS9{F0: 128, F1: 4294967295, F2: false}
                                                _ = v95
                                            }
                                            v96 := GenS10{F0: v6}
                                            _ = v96
                                        }
                                        v4[(v6 / ((128) | 1))] = v2
                                        var v97 int = (v6 % ((851) | 1))
                                        _ = v97
                                    }
                                }
                            }
                        case 0:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case 1:
                            for fi98 := 0; fi98 < 7; fi98++ {
                                v3 = append(v3, (v2 % ((516) | 1)))
                                for fi99 := 0; fi99 < 8; fi99++ {
                                    v87 = 1e10
                                    if true {
                                        if (!false) {
                                            fmt.Println(v6, fi98)
                                            v100 := GenS11{F0: 1000, F1: (("hello" + "") + "hello"), F2: v6}
                                            fmt.Println(v2)
                                            if true {
                                                if (!(!true)) {
                                                    v101 := GenS11{F0: 0, F1: "foo bar", F2: (fi98 * fi98)}
                                                    var v102 int32 = 2147483647
                                                    v87 = ((float32(0.0) / 1.5) * (v87 * v87))
                                                    _ = v101
                                                    _ = v102
                                                } else {
                                                    v87 -= (float32(1e10) / v87)
                                                    v103 := v4[(v6 | 465)]
                                                    if (((("" <= "hello") && (!(42 <= 0))) || true) && false) {
                                                        var v104 float32 = (v87 + v87)
                                                        _ = len(v4)
                                                        _ = len(v4)
                                                        _ = v104
                                                    }
                                                    _ = v103
                                                }
                                                v4[(fi98 >> ((839) & 63))] = 1
                                            }
                                            v5 -= 42
                                            if (!false) {
                                                var v105 uint = 210
                                                if true {
                                                    switch "hello" {
                                                    case "foo bar":
                                                        v13 += 255
                                                    }
                                                    var v106 int16 = 42
                                                    v107 := GenS1{F0: 703, F1: (v2 | (v2 ^ (^(v2)))), F2: 2}
                                                    v4[fi98] = uint32(v13)
                                                    v87 += (v87 - v87)
                                                    _ = v106
                                                    _ = v107
                                                }
                                                _ = v105
                                            }
                                            _ = v100
                                        } else {
                                            delete(v4, v6)
                                            v108 := GenS12{F0: -781}
                                            v109 := GenS13{F0: 86, F1: 42}
                                            _ = v108
                                            _ = v109
                                        }
                                    } else {
                                        switch (fi98 << ((65535) & 63)) {
                                        case 2147483647:
                                            if true {
                                                var v110 uint64 = 4294967295
                                                var v111 int16 = -32768
                                                var v112 bool = (!(true && true))
                                                _ = v110
                                                _ = v111
                                                _ = v112
                                            }
                                        case 127:
                                            v6 = (v6 - 858)
                                        case 32768:
                                            var v113 uint8 = 42
                                            _ = v113
                                        default:
                                            delete(v4, 2147483647)
                                        }
                                        delete(v4, -9223372036854775808)
                                    }
                                    _ = len(v4)
                                    _ = fi99
                                }
                                for fi114 := 0; fi114 < 4; fi114++ {
                                    fmt.Println(v87, v6, v87)
                                    var v115 uint = 480
                                    var v116 int16 = 32767
                                    v2 = v2
                                    v117 := GenS14{F0: 42, F1: v115}
                                    _ = fi114
                                    _ = v115
                                    _ = v116
                                    _ = v117
                                }
                                _ = fi98
                            }
                        }
                        var v118 uint = 65535
                        _ = v118
                    } else {
                        v2 = (v2 >> ((65535) & 63))
                        v5 -= (^(v5))
                    }
                case "\n\t":
                    v6 = v6
                }
                _ = v87
            }
        }
    }
    var v119 uint64 = 42
    var v120 int16 = 42
    var v121 uint32 = 2
    if (!true) {
        switch "" {
        case "a":
            if (42 < 32767) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if (true || false) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    for fi122 := 0; fi122 < 1; fi122++ {
                        v123 := GenS14{F0: v5, F1: 42}
                        _ = fi122
                        _ = v123
                    }
                }
                var v124 int8 = 92
                _ = v124
            }
        case "hello":
            v120 += v120
        case "":
            if (!((!(0.0 >= (((1.5 * 1e10) + (float64(0.0) / 1e10)) + 1e10))) || (-716 > v120))) {
                v5 += -9223372036854775808
            } else {
                v5 = v5
            }
        default:
            var v125 int32 = 42
            _ = v125
        }
        _ = len(v3)
        for fi126 := 0; fi126 < 14; fi126++ {
            v121 = (v121 & v121)
            v121 -= v2
            v119 -= (v119 / (((v119 & 708)) | 1))
            _ = fi126
        }
    }
    fmt.Println(v2, v13, v119)
    for fi127 := 0; fi127 < 14; fi127++ {
        v120 -= (v120 - (^(v120)))
        _ = fi127
    }
    v13 = v13
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v119:", v119)
    fmt.Println("v120:", v120)
    fmt.Println("v121:", v121)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v119
    _ = v120
    _ = v121
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
