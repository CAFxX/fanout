package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint8; F1 uint; F2 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32; F1 int64; F2 uint32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint8; F1 bool }
type GenS8 struct { F0 int32; F1 int16; F2 int32 }
type GenS9 struct { F0 string; F1 uint32 }
type GenS10 struct { F0 int64; F1 int64 }
type GenS11 struct { F0 int; F1 bool }
type GenS12 struct { F0 uint32; F1 int16; F2 int8 }
type GenS13 struct { F0 string; F1 int16; F2 uint }
type GenS14 struct { F0 uint32; F1 int64 }
type GenS15 struct { F0 int16; F1 int32 }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 int64; F1 int64 }
type GenS18 struct { F0 uint32; F1 uint8 }
func genFn0(p0 int8) uint8 {
    v7 := GenS1{F0: 1, F1: -85}
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint8 = 2
        fmt.Println(fi8, v9)
        if (-2.25 > -2.25) {
            p0 = 42
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 int8 = p0
                var v12 float32 = 1e10
                var v13 float64 = 0.0
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
        }
        _ = fi8
        _ = v9
    }
    v14 := GenS1{F0: 0, F1: 1}
    var v15 int32 = 53
    if (false || (!(!true))) {
        var v16 string = ("\n\t" + ((("hello" + "") + "\n\t") + ("a" + "a")))
        fmt.Println(v16, v16, v15)
        v17 := GenS2{F0: 147, F1: 42, F2: 833}
        _ = v16
        _ = v17
    } else {
        if false {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(p0)
            default:
                v18 := []int{-9223372036854775808}
                _ = v18
            }
            if (1000 >= 508) {
                switch "foo bar" {
                case "\n\t":
                    v19 := []int{}
                    _ = v19
                case "hello":
                    fmt.Println(p0, p0)
                case "a":
                    for fi20 := 0; fi20 < 8; fi20++ {
                        v21 := []int{fi20, fi20}
                        _ = fi20
                        _ = v21
                    }
                }
                fmt.Println(v15)
                v22 := []int{32768, 32767, 132, -9223372036854775808}
                _ = v22
            } else {
                for fi23 := 0; fi23 < 13; fi23++ {
                    v24 := GenS3{F0: (p0 / (((p0 >> ((516) & 63))) | 1))}
                    if false {
                        v15 = v15
                    } else {
                        var v25 int8 = 0
                        p0 -= 0
                        _ = v25
                    }
                    _ = fi23
                    _ = v24
                }
                p0 = 0
                v26 := GenS1{F0: 127, F1: (v15 * v15)}
                _ = v26
            }
            if false {
                var v27 uint8 = 92
                for fi28 := 0; fi28 < 15; fi28++ {
                    var v29 int16 = -32768
                    v15 = v15
                    _ = fi28
                    _ = v29
                }
                _ = v27
            } else {
                for fi30 := 0; fi30 < 12; fi30++ {
                    v31 := GenS3{F0: (-(p0))}
                    _ = fi30
                    _ = v31
                }
                v15 = v15
                for fi32 := 0; fi32 < 19; fi32++ {
                    if false {
                        fmt.Println(p0, p0, p0)
                        if true {
                            v33 := []int{(fi32 & 2147483647), (fi32 + fi32)}
                            v15 += v15
                            _ = v33
                        } else {
                            p0 -= p0
                        }
                        v34 := GenS2{F0: 0, F1: 954, F2: v15}
                        _ = v34
                    }
                    _ = fi32
                }
            }
        }
        switch "\n\t" {
        case "foo bar":
            for fi35 := 0; fi35 < 18; fi35++ {
                v36 := []int{}
                _ = fi35
                _ = v36
            }
        case "\n\t":
            p0 += 68
        }
        p0 = p0
    }
    _ = p0
    _ = v14
    _ = v15
    _ = v7
    return 94
}

func main() {
    var v1 int = 812
    v2 := []uint32{2}
    var v3 uint = 584
    var v4 int64 = 42
    v5 := []uint32{}
    v6 := []uint32{65535}
    var v37 float64 = -2.25
    if ((0.0 + ((1.5 + v37) + 1.5)) >= v37) {
        if (v37 == 1.5) {
            if false {
                v1 -= v1
                switch v1 {
                case 2147483647:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            } else {
                v37 = v37
                var v38 uint = (v3 / ((v3) | 1))
                v37 += float64(1e10)
                _ = v38
            }
            v3 += (^(v3))
            v39 := GenS4{F0: 1}
            _ = v39
        } else {
            for fi40 := 0; fi40 < 4; fi40++ {
                if false {
                    v41 := GenS3{F0: 75}
                    v42 := GenS1{F0: 127, F1: -1}
                    _ = v41
                    _ = v42
                } else {
                    v43 := GenS5{F0: -2147483648, F1: v4, F2: 2}
                    _ = v43
                }
                _ = fi40
            }
            v44 := GenS6{F0: false}
            _ = v44
        }
        v37 -= v37
        for fi45 := 0; fi45 < 18; fi45++ {
            if false {
                for fi46 := 0; fi46 < 9; fi46++ {
                    switch "foo bar" {
                    case "foo bar":
                        if ((true && true) && false) {
                            if (50 > 31) {
                                var v47 float64 = -2.25
                                v1 = fi46
                                fmt.Println(fi46)
                                _ = v47
                            }
                            v48 := GenS4{F0: 76}
                            if (((-(v4)) > (v4 + (v4 | (v4 << (((^(v3))) & 63))))) && true) {
                                if false {
                                    v3 = (v3 & (^(v3)))
                                    v37 = (v37 + v37)
                                    var v49 uint32 = 2147483647
                                    v50 := GenS7{F0: genFn0(-128), F1: (1 < 199)}
                                    _ = v49
                                    _ = v50
                                } else {
                                    v51 := GenS6{F0: (false && false)}
                                    var v52 int64 = (v4 - v4)
                                    _ = v51
                                    _ = v52
                                }
                            }
                            switch (-(fi46)) {
                            case 32768:
                                switch -1 {
                                case -2147483648:
                                    if true {
                                        v1 -= -1
                                        var v53 string = ("a" + "foo bar")
                                        _ = v53
                                    } else {
                                        if false {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v54 := GenS7{F0: 74, F1: false}
                                            _ = v54
                                        }
                                        var v55 uint8 = genFn0(100)
                                        _ = v55
                                    }
                                }
                            case 212:
                                fmt.Println(v37)
                            case -943:
                                var v56 int64 = v4
                                if (!((("hello" + "a") != (("a" + ("a" + "foo bar")) + (("a" + ("foo bar" + (("foo bar" + "hello") + "hello"))) + ("hello" + ("foo bar" + "foo bar"))))) || ((!true) && ((v3 * v3) >= 781)))) {
                                    fmt.Println(fi46, fi46, fi45)
                                    fmt.Println(v56)
                                }
                                _ = v56
                            default:
                                var v57 float32 = 0.1
                                v37 = v37
                                _ = v57
                            }
                            _ = v48
                        }
                    case "\n\t":
                        if false {
                            var v58 uint32 = 42
                            _ = v58
                        } else {
                            if false {
                                v59 := GenS8{F0: 1, F1: 0, F2: -640}
                                v60 := GenS9{F0: ("foo bar" + ("foo bar" + "hello")), F1: 2147483647}
                                _ = len(v5)
                                _ = v59
                                _ = v60
                            }
                            fmt.Println(fi46)
                            v37 -= v37
                        }
                    case "hello":
                        v4 = v4
                    }
                    var v61 int16 = 42
                    _ = fi46
                    _ = v61
                }
                var v62 float32 = (0.0 - 0.0)
                switch fi45 {
                case 0:
                    if false {
                        if true {
                            if true {
                                var v63 uint8 = 255
                                v64 := GenS10{F0: v4, F1: -9223372036854775808}
                                switch "foo bar" {
                                case "hello":
                                    if (!(true || false)) {
                                        if (-568 < -1) {
                                            fmt.Println(v62, v63, v62)
                                        }
                                        v1 += (-(fi45))
                                    } else {
                                        v37 = ((float64(float32((v37 - (v37 - v37)))) * -2.25) * v37)
                                        if (("a" <= "hello") && ((v3 > v3) && (false || true))) {
                                            _ = len(v2)
                                        }
                                        v1 = fi45
                                    }
                                case "a":
                                    v65 := GenS11{F0: v1, F1: ((true || (!(v3 <= (v3 >> (((v3 / ((v3) | 1))) & 63))))) && (false || (true && (2 >= 0))))}
                                    _ = v65
                                }
                                _ = v63
                                _ = v64
                            } else {
                                var v66 string = "hello"
                                _ = v66
                            }
                            v1 -= fi45
                        } else {
                            for fi67 := 0; fi67 < 11; fi67++ {
                                v4 = v4
                                v68 := GenS12{F0: 4294967295, F1: -1, F2: 127}
                                _ = fi67
                                _ = v68
                            }
                            var v69 float64 = ((v37 / v37) * 1e10)
                            _ = v69
                        }
                        fmt.Println(v3, fi45)
                    }
                default:
                    v37 = ((((float64(((0.0 + v62) * 1e10)) / v37) - (3.14159 * v37)) * 0.1) * (v37 / v37))
                }
                _ = v62
            } else {
                if ((v4 % ((v4) | 1)) == v4) {
                    if len(v6) > 0 {
                        v6[0] = 42
                    }
                } else {
                    var v70 uint64 = 4294967295
                    _ = v70
                }
                v4 -= (v4 % ((v4) | 1))
            }
            if (!(((297 >= 528) && true) || (!false))) {
                switch "a" {
                case "a":
                    if (32767 != (fi45 % ((fi45) | 1))) {
                        for fi71 := 0; fi71 < 11; fi71++ {
                            v4 -= (v4 + -811)
                            v37 = ((0.0 * v37) + v37)
                            _ = fi71
                        }
                        if true {
                            var v72 uint32 = 1
                            v1 += int(v72)
                            _ = v72
                        } else {
                            if (1 != 42) {
                                _ = len(v6)
                                switch "" {
                                case "a":
                                    v73 := GenS13{F0: "a", F1: 700, F2: 2}
                                    _ = v73
                                case "\n\t":
                                    switch "a" {
                                    case "foo bar":
                                        v1 += 128
                                    case "":
                                        var v74 int16 = 32767
                                        _ = v74
                                    }
                                }
                            } else {
                                var v75 string = "\n\t"
                                var v76 float64 = (-2.25 / float64((((0.0 - 0.1) / -2.25) * ((((0.0 + (3.14159 / (1e10 - 3.14159))) - ((1e10 * float32(0.1)) + -2.25)) * 1.5) + 1.5))))
                                if ((-2.25 + -2.25) >= v37) {
                                    var v77 bool = ((v3 + v3) == (v3 * (v3 ^ v3)))
                                    if ((1e10 - 0.1) <= 3.14159) {
                                        switch (fi45 * (v1 >> (((^(v3))) & 63))) {
                                        case -1:
                                            v78 := GenS8{F0: 0, F1: 0, F2: 42}
                                            _ = v78
                                        default:
                                            fmt.Println(v77, v37, fi45)
                                        }
                                    } else {
                                        v4 = (v4 % ((v4) | 1))
                                        if v77 {
                                            v1 += 42
                                            switch v1 {
                                            case 128:
                                                for fi79 := 0; fi79 < 12; fi79++ {
                                                    if (v77 && v77) {
                                                        v77 = (v4 >= 343)
                                                    }
                                                    v80 := GenS14{F0: 140, F1: -9223372036854775808}
                                                    _ = fi79
                                                    _ = v80
                                                }
                                            case -2147483648:
                                                var v81 uint8 = 2
                                                _ = v81
                                            case 2147483647:
                                                if v77 {
                                                    if false {
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                        var v82 int64 = (^(v4))
                                                        v37 = (0.1 - ((v37 + v76) + v37))
                                                        _ = v82
                                                    } else {
                                                        for fi83 := 0; fi83 < 12; fi83++ {
                                                            v84 := GenS15{F0: -926, F1: -220}
                                                            v77 = ((1 < (fi45 % (((v1 & (fi83 | (^(fi45))))) | 1))) || v77)
                                                            if len(v2) > 0 {
                                                                v2[0] = 0
                                                            }
                                                            _ = fi83
                                                            _ = v84
                                                        }
                                                    }
                                                    v77 = v77
                                                    v37 = (1e10 + 0.0)
                                                } else {
                                                    v85 := GenS3{F0: int8(v3)}
                                                    var v86 int = 806
                                                    _ = v85
                                                    _ = v86
                                                }
                                            }
                                        }
                                    }
                                    var v87 uint16 = 2
                                    _ = v77
                                    _ = v87
                                } else {
                                    if (-128 >= 1) {
                                        v88 := GenS6{F0: true}
                                        switch fi45 {
                                        case -1:
                                            v37 = v76
                                        case -9223372036854775808:
                                            for fi89 := 0; fi89 < 11; fi89++ {
                                                var v90 uint64 = 128
                                                v91 := GenS5{F0: int32(v3), F1: (v4 & v4), F2: 65535}
                                                _ = fi89
                                                _ = v90
                                                _ = v91
                                            }
                                        }
                                        v92 := GenS11{F0: v1, F1: true}
                                        _ = v88
                                        _ = v92
                                    }
                                    v5 = append(v5, 0)
                                    var v93 int16 = -1
                                    _ = v93
                                }
                                _ = v75
                                _ = v76
                            }
                        }
                        for fi94 := 0; fi94 < 14; fi94++ {
                            v95 := GenS7{F0: 42, F1: true}
                            _ = fi94
                            _ = v95
                        }
                    }
                }
                v96 := GenS15{F0: 1, F1: -1}
                _ = v96
            } else {
                fmt.Println(v1, v4)
            }
            _ = fi45
        }
    } else {
        v3 = v3
        var v97 int16 = -333
        _ = v97
    }
    for fi98 := 0; fi98 < 16; fi98++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi98
    }
    v99 := GenS16{F0: 42}
    switch ("foo bar" + "") {
    case "hello":
        for fi100 := 0; fi100 < 8; fi100++ {
            for fi101 := 0; fi101 < 11; fi101++ {
                var v102 int8 = 100
                _ = fi101
                _ = v102
            }
            _ = fi100
        }
    }
    v103 := GenS17{F0: (v4 ^ -9223372036854775808), F1: (v4 >> (((v3 - v3)) & 63))}
    for fi104 := 0; fi104 < 13; fi104++ {
        v1 = int(v3)
        _ = fi104
    }
    var v105 uint = v3
    switch ("\n\t" + ("" + "\n\t")) {
    case "hello":
        v4 -= 32
    case "":
        fmt.Println(v37, v3)
    default:
        for fi106 := 0; fi106 < 12; fi106++ {
            if true {
                if true {
                    fmt.Println(v3, v4, v105)
                    v105 = v105
                    v4 = 1
                } else {
                    v107 := GenS1{F0: 2, F1: 2147483647}
                    if true {
                        var v108 string = "a"
                        switch ("" + v108) {
                        case "hello":
                            if (!((!true) || true)) {
                                v109 := GenS18{F0: 604, F1: genFn0(13)}
                                var v110 int16 = 1000
                                _ = v109
                                _ = v110
                            } else {
                                var v111 float64 = (0.0 + v37)
                                _ = v111
                            }
                        case "foo bar":
                            v112 := GenS8{F0: 1, F1: 359, F2: int32(v1)}
                            _ = v112
                        case "a":
                            var v113 int16 = 42
                            _ = v113
                        }
                        _ = v108
                    }
                    _ = v107
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                for fi114 := 0; fi114 < 17; fi114++ {
                    v3 = 1
                    _ = fi114
                }
            } else {
                var v115 float32 = float32((v37 - v37))
                v37 = v37
                _ = v115
            }
            _ = fi106
        }
    }
    fmt.Println(v4, v105)
    _ = len(v5)
    var v116 uint32 = 4294967295
    v4 = (v4 ^ 935)
    switch "\n\t" {
    case "":
        if true {
            if len(v6) > 0 {
                v6[0] = v116
            }
            var v117 uint = v3
            v5 = append(v5, (v116 % ((v116) | 1)))
            _ = v117
        } else {
            var v118 uint8 = uint8(v3)
            v118 -= 2
            _ = v118
        }
    case "a":
        switch "hello" {
        case "\n\t":
            var v119 uint8 = 255
            _ = v119
        case "":
            var v120 uint8 = 42
            _ = v120
        default:
            if true {
                switch ("hello" + "hello") {
                case "":
                    if (!((v116 | (v116 % ((v116) | 1))) == 4294967295)) {
                        var v121 string = "a"
                        switch "\n\t" {
                        case "foo bar":
                            v122 := GenS15{F0: 42, F1: 42}
                            _ = v122
                        case "a":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            var v123 uint8 = 0
                            _ = v123
                        }
                        _ = v121
                    } else {
                        if true {
                            v2 = append(v2, 4294967295)
                        } else {
                            v3 -= v105
                        }
                        var v124 int64 = 9223372036854775807
                        _ = v124
                    }
                case "hello":
                    if (!(42 > -1)) {
                        fmt.Println(v3, v105, v116)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if true {
                            for fi125 := 0; fi125 < 17; fi125++ {
                                switch v1 {
                                case 32768:
                                    v1 = (v1 & 128)
                                case -406:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                case -1:
                                    v2 = append(v2, (v116 & (v116 << (((v3 & 4294967295)) & 63))))
                                }
                                var v126 string = "a"
                                _ = fi125
                                _ = v126
                            }
                            if (true && (!(18446744073709551615 < 42))) {
                                v127 := GenS3{F0: 0}
                                v105 = (v3 | v105)
                                _ = v127
                            } else {
                                v37 = v37
                            }
                        } else {
                            for fi128 := 0; fi128 < 9; fi128++ {
                                for fi129 := 0; fi129 < 9; fi129++ {
                                    fmt.Println(fi128, v3)
                                    _ = fi129
                                }
                                _ = fi128
                            }
                        }
                    }
                case "\n\t":
                    if len(v6) > 0 {
                        v6[0] = 618
                    }
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("v116:", v116)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v103
    _ = v105
    _ = v116
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v99
}
