package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int32 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint; F1 int; F2 uint16 }
type GenS4 struct { F0 uint64; F1 int; F2 uint32 }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint; F1 uint32 }
type GenS10 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS11 struct { F0 uint32; F1 uint8; F2 bool }
type GenS12 struct { F0 bool; F1 int32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 1
    v3 := []int{}
    var v4 bool = (!(81 >= 42))
    if v4 {
        for fi5 := 0; fi5 < 9; fi5++ {
            v4 = true
            _ = fi5
        }
        v6 := GenS1{F0: -1, F1: 625, F2: 2147483647}
        v4 = true
        _ = v6
    } else {
        v4 = ((("\n\t" + ("" + ("" + ""))) >= "foo bar") && ((v4 || v4) && v4))
        if len(v3) > 0 {
            v3[0] = 1
        }
        var v7 int = -10
        _ = v7
    }
    v4 = (v4 || (954 >= 683))
    v2 = (v2 ^ v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v4)
    v8 := GenS2{F0: (v2 * v2), F1: 117}
    var v9 int32 = 0
    if (true || (!false)) {
        v10 := GenS1{F0: 0, F1: 440, F2: (v9 | 1)}
        switch "a" {
        case "a":
            if v4 {
                switch -200 {
                case 1:
                    if v4 {
                        for fi11 := 0; fi11 < 13; fi11++ {
                            v12 := GenS3{F0: 65535, F1: fi11, F2: 65535}
                            switch "" {
                            case "":
                                v13 := GenS2{F0: 1000, F1: 0}
                                _ = v13
                            case "foo bar":
                                fmt.Println(v4, v9, v4)
                            case "\n\t":
                                v2 -= v2
                            default:
                                var v14 int32 = v9
                                _ = v14
                            }
                            _ = fi11
                            _ = v12
                        }
                    } else {
                        if (!(128 == 42)) {
                            v3 = append(v3, 2)
                        } else {
                            var v15 bool = (42 < 127)
                            v16 := GenS4{F0: 0, F1: -844, F2: 2}
                            _ = v15
                            _ = v16
                        }
                        switch 1 {
                        case 9223372036854775807:
                            var v17 uint64 = 18446744073709551615
                            _ = v17
                        case -538:
                            v18 := GenS5{F0: 128, F1: 0, F2: 1}
                            _ = v18
                        default:
                            for fi19 := 0; fi19 < 17; fi19++ {
                                var v20 string = "foo bar"
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = fi19
                                _ = v20
                            }
                        }
                    }
                case -1:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case -422:
                    if v4 {
                        if v4 {
                            fmt.Println(v4, v9)
                        } else {
                            v9 = 1
                            for fi21 := 0; fi21 < 12; fi21++ {
                                if (-2.25 >= float64((-2.25 / 0.1))) {
                                    fmt.Println(v9, v4, v2)
                                    var v22 int32 = v9
                                    _ = v22
                                } else {
                                    v23 := GenS2{F0: -676, F1: 2}
                                    for fi24 := 0; fi24 < 5; fi24++ {
                                        var v25 int8 = 0
                                        var v26 uint64 = 18446744073709551615
                                        _ = fi24
                                        _ = v25
                                        _ = v26
                                    }
                                    _ = v23
                                }
                                switch fi21 {
                                case -9223372036854775808:
                                    v2 -= v2
                                default:
                                    var v27 uint8 = 255
                                    _ = v27
                                }
                                _ = fi21
                            }
                            v28 := GenS1{F0: 1, F1: 42, F2: (-(v9))}
                            _ = v28
                        }
                        var v29 uint64 = 1
                        _ = v29
                    } else {
                        if true {
                            for fi30 := 0; fi30 < 14; fi30++ {
                                var v31 string = ""
                                fmt.Println(v4, v9, fi30)
                                _ = fi30
                                _ = v31
                            }
                            v2 = (v2 << ((807) & 63))
                        } else {
                            fmt.Println(v4, v2, v2)
                            var v32 int16 = (-(v2))
                            var v33 uint8 = 128
                            _ = v32
                            _ = v33
                        }
                        if v4 {
                            v34 := GenS2{F0: v2, F1: 158}
                            if (true || (149 > 42)) {
                                for fi35 := 0; fi35 < 8; fi35++ {
                                    v9 = (v9 - 2147483647)
                                    v36 := GenS5{F0: 194, F1: 1, F2: 32767}
                                    if (v4 && v4) {
                                        var v37 float64 = 0.0
                                        v9 = (v9 & (v9 + (v9 ^ 42)))
                                        v4 = (65535 >= 1)
                                        _ = v37
                                    }
                                    _ = fi35
                                    _ = v36
                                }
                                if v4 {
                                    switch -9223372036854775808 {
                                    case 326:
                                        if (!v4) {
                                            var v38 int64 = 200
                                            fmt.Println(v38)
                                            if v4 {
                                                v38 = (-(v38))
                                                if (!v4) {
                                                    var v39 uint8 = 237
                                                    for fi40 := 0; fi40 < 19; fi40++ {
                                                        if v4 {
                                                            fmt.Println(v39, fi40, fi40)
                                                        } else {
                                                            for fi41 := 0; fi41 < 19; fi41++ {
                                                                switch 130 {
                                                                case 48:
                                                                    fmt.Println(v38, v2, v2)
                                                                case -765:
                                                                    if (1 >= 65535) {
                                                                        v4 = (v2 > v2)
                                                                        delete(v1, "a")
                                                                        v42 := GenS6{F0: 2147483648, F1: 1}
                                                                        v2 = (v2 * v2)
                                                                        switch fi41 {
                                                                        case -868:
                                                                            var v43 int64 = v38
                                                                            _ = v43
                                                                        default:
                                                                            var v44 int8 = 100
                                                                            _ = v44
                                                                        }
                                                                        var v45 uint16 = 42
                                                                        var v46 uint16 = v45
                                                                        _ = v42
                                                                        _ = v45
                                                                        _ = v46
                                                                    }
                                                                case 42:
                                                                    v2 = (v2 - 32767)
                                                                }
                                                                v47 := GenS1{F0: fi40, F1: 569, F2: v9}
                                                                if v4 {
                                                                    v9 += (v9 + (v9 | (v9 * v9)))
                                                                    v4 = true
                                                                    var v48 string = "hello"
                                                                    _ = v48
                                                                }
                                                                _ = fi41
                                                                _ = v47
                                                            }
                                                            for fi49 := 0; fi49 < 11; fi49++ {
                                                                var v50 uint16 = 1
                                                                var v51 int = -366
                                                                fmt.Println(v4, v2, v39)
                                                                _ = fi49
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                        }
                                                        v52 := GenS7{F0: (("foo bar" + "a") + "a"), F1: v9}
                                                        _ = fi40
                                                        _ = v52
                                                    }
                                                    var v53 uint8 = v39
                                                    _ = v39
                                                    _ = v53
                                                }
                                            }
                                            _ = v38
                                        } else {
                                            _ = len(v3)
                                            fmt.Println(v2, v9)
                                            v2 = v2
                                        }
                                    case 127:
                                        for fi54 := 0; fi54 < 14; fi54++ {
                                            if false {
                                                v4 = v4
                                                var v55 int32 = 2147483647
                                                _ = v55
                                            } else {
                                                v4 = (0.0 == 0.1)
                                                fmt.Println(v9)
                                            }
                                            _ = fi54
                                        }
                                    case 32768:
                                        if v4 {
                                            v9 = v9
                                            fmt.Println(v4, v2)
                                            if ((18446744073709551615 >= 2) && v4) {
                                                for fi56 := 0; fi56 < 11; fi56++ {
                                                    fmt.Println(v9, v2, v4)
                                                    fmt.Println(v2)
                                                    for fi57 := 0; fi57 < 4; fi57++ {
                                                        if v4 {
                                                            delete(v1, "\n\t")
                                                            v58 := GenS8{F0: (true && (((-2.25 - (float32(float64((1e10 * (1e10 + 1.5)))) / (1.5 / 0.1))) - (1.5 + (3.14159 * (1.5 / (3.14159 / 1.5))))) >= float32(-2.25)))}
                                                            if (650 == 4294967295) {
                                                                v9 = (v9 & v9)
                                                                fmt.Println(v9, fi57)
                                                            }
                                                            _ = len(v3)
                                                            _ = v58
                                                        } else {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            v59 := GenS3{F0: 65535, F1: (fi57 * (fi57 & fi57)), F2: 106}
                                                            var v60 float32 = 1e10
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                        _ = fi57
                                                    }
                                                    _ = fi56
                                                }
                                            }
                                        }
                                    }
                                } else {
                                    v9 = -961
                                }
                                var v61 float32 = -2.25
                                _ = v61
                            } else {
                                var v62 int8 = 1
                                _ = v62
                            }
                            _ = v34
                        }
                    }
                }
                switch 127 {
                case -85:
                    v63 := GenS3{F0: 0, F1: 32768, F2: 255}
                    _ = v63
                case 0:
                    fmt.Println(v9)
                }
            }
        case "foo bar":
            v64 := GenS4{F0: 0, F1: -9223372036854775808, F2: 65535}
            _ = v64
        case "":
            if ((65 < 1) || false) {
                v65 := GenS3{F0: 782, F1: 1, F2: 2}
                v66 := GenS9{F0: 0, F1: 42}
                fmt.Println(v2)
                _ = v65
                _ = v66
            } else {
                v2 = v2
                switch -571 {
                case 127:
                    switch 185 {
                    case -707:
                        fmt.Println(v9, v2, v2)
                    case 2147483647:
                        for fi67 := 0; fi67 < 7; fi67++ {
                            if v4 {
                                var v68 int16 = v2
                                var v69 uint32 = 14
                                _ = v68
                                _ = v69
                            } else {
                                v2 = (v2 & (v2 - -32768))
                            }
                            var v70 uint16 = 754
                            _ = fi67
                            _ = v70
                        }
                    case 0:
                        if ((32767 >= 2) && v4) {
                            switch -225 {
                            case -1:
                                for fi71 := 0; fi71 < 12; fi71++ {
                                    v72 := GenS4{F0: 1, F1: fi71, F2: 2147483648}
                                    _ = fi71
                                    _ = v72
                                }
                            case -2147483648:
                                for fi73 := 0; fi73 < 1; fi73++ {
                                    switch "hello" {
                                    case "foo bar":
                                        if v4 {
                                            var v74 int = fi73
                                            var v75 float32 = (1.5 - (1.5 + (-2.25 / 0.1)))
                                            _ = v74
                                            _ = v75
                                        }
                                    }
                                    _ = fi73
                                }
                            }
                            v9 += -2147483648
                            fmt.Println(v4)
                        } else {
                            if (128 < 118) {
                                switch "" {
                                case "a":
                                    if (true && (32767 <= 32767)) {
                                        var v76 bool = true
                                        v2 += 1000
                                        var v77 int64 = 480
                                        _ = v76
                                        _ = v77
                                    } else {
                                        switch -9223372036854775808 {
                                        case -696:
                                            switch 128 {
                                            case -2147483648:
                                                var v78 int64 = 9223372036854775807
                                                _ = v78
                                            case 32768:
                                                for fi79 := 0; fi79 < 18; fi79++ {
                                                    v4 = false
                                                    for fi80 := 0; fi80 < 6; fi80++ {
                                                        fmt.Println(v4, v9, v4)
                                                        v2 -= (v2 ^ v2)
                                                        v81 := GenS8{F0: ((0.0 != 0.1) || false)}
                                                        _ = fi80
                                                        _ = v81
                                                    }
                                                    _ = fi79
                                                }
                                            default:
                                                v82 := GenS9{F0: 65535, F1: 4294967295}
                                                _ = v82
                                            }
                                        case 32767:
                                            if v4 {
                                                if len(v3) > 0 {
                                                    v3[0] = -9223372036854775808
                                                }
                                            } else {
                                                var v83 float32 = 0.0
                                                for fi84 := 0; fi84 < 2; fi84++ {
                                                    v85 := GenS3{F0: 203, F1: -2147483648, F2: 433}
                                                    v83 = (3.14159 - v83)
                                                    v83 += ((0.1 / 1.0) + ((((v83 * v83) * (((float32(float64((v83 / (3.14159 * v83)))) - v83) + v83) / (v83 / v83))) / (1e10 - (1e10 * v83))) - v83))
                                                    _ = fi84
                                                    _ = v85
                                                }
                                                var v86 int16 = (v2 % ((v2) | 1))
                                                _ = v83
                                                _ = v86
                                            }
                                        case 1:
                                            v87 := GenS4{F0: 18446744073709551615, F1: 981, F2: 65535}
                                            _ = v87
                                        default:
                                            for fi88 := 0; fi88 < 9; fi88++ {
                                                for fi89 := 0; fi89 < 1; fi89++ {
                                                    v2 = 1
                                                    _ = fi89
                                                }
                                                _ = fi88
                                            }
                                        }
                                        v2 += v2
                                        for fi90 := 0; fi90 < 6; fi90++ {
                                            for fi91 := 0; fi91 < 6; fi91++ {
                                                var v92 uint64 = 42
                                                var v93 int = fi91
                                                fmt.Println(v4, fi91, v92)
                                                var v94 uint64 = 0
                                                var v95 float64 = ((((((0.1 - 3.14159) + (1.5 / 3.14159)) + (-2.25 * 0.0)) * 1.5) - float64(((3.14159 + 1.5) + float32(float64(0.0))))) + (1e10 + ((1.5 - 1e10) / 1.0)))
                                                _ = fi91
                                                _ = v92
                                                _ = v93
                                                _ = v94
                                                _ = v95
                                            }
                                            var v96 uint = 18446744073709551615
                                            v96 = 746
                                            _ = fi90
                                            _ = v96
                                        }
                                    }
                                case "foo bar":
                                    fmt.Println(v4, v4, v2)
                                }
                                var v97 string = "hello"
                                v98 := GenS10{F0: 2, F1: 42, F2: 2}
                                _ = v97
                                _ = v98
                            } else {
                                v99 := GenS11{F0: 616, F1: 1, F2: v4}
                                v9 = v9
                                if true {
                                    var v100 uint16 = 42
                                    _ = v100
                                }
                                _ = v99
                            }
                        }
                    default:
                        v101 := GenS12{F0: (100 <= 0), F1: v9}
                        _ = v101
                    }
                case 32767:
                    switch "\n\t" {
                    case "\n\t":
                        var v102 int = 32767
                        _ = v102
                    case "hello":
                        var v103 float32 = ((0.1 + (((float32((-2.25 / (1e10 - 1.5))) / 1e10) * float32(0.0)) * float32((3.14159 - 0.0)))) - ((float32(0.1) / 0.1) - float32(((float64((3.14159 + float32((0.1 * (0.1 - (3.14159 + 0.1)))))) * 0.0) * ((1e10 - 0.1) - -2.25)))))
                        _ = v103
                    case "foo bar":
                        v9 -= v9
                    default:
                        for fi104 := 0; fi104 < 8; fi104++ {
                            for fi105 := 0; fi105 < 7; fi105++ {
                                v1["foo bar"] = 258
                                switch (^(fi104)) {
                                case 1:
                                    fmt.Println(fi105, fi104)
                                default:
                                    v9 = 0
                                }
                                _ = fi105
                            }
                            _ = fi104
                        }
                    }
                default:
                    if false {
                        if (!false) {
                            var v106 string = "foo bar"
                            switch ("hello" + v106) {
                            case "\n\t":
                                v9 = v9
                            case "hello":
                                var v107 int = 4294967295
                                _ = v107
                            case "foo bar":
                                if (9223372036854775807 == 9223372036854775807) {
                                    switch 2147483647 {
                                    case 32767:
                                        v9 = (^(v9))
                                    case -1:
                                        switch 302 {
                                        case -9223372036854775808:
                                            v9 += (v9 & (-(v9)))
                                        case 0:
                                            for fi108 := 0; fi108 < 2; fi108++ {
                                                v9 -= v9
                                                _ = fi108
                                            }
                                        case -2147483648:
                                            switch (("foo bar" + (v106 + v106)) + "") {
                                            case "a":
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            case "\n\t":
                                                var v109 uint32 = 42
                                                _ = v109
                                            case "foo bar":
                                                if false {
                                                    var v110 string = ("foo bar" + v106)
                                                    _ = v110
                                                } else {
                                                    v111 := GenS8{F0: v4}
                                                    _ = v111
                                                }
                                            }
                                        default:
                                            var v112 uint32 = 334
                                            _ = v112
                                        }
                                    default:
                                        var v113 uint = 18446744073709551615
                                        _ = v113
                                    }
                                    var v114 int16 = v2
                                    v1["hello"] = 200
                                    _ = v114
                                }
                            }
                            v115 := GenS4{F0: 18446744073709551615, F1: -9223372036854775808, F2: 42}
                            _ = v106
                            _ = v115
                        } else {
                            fmt.Println(v9, v4)
                            for fi116 := 0; fi116 < 19; fi116++ {
                                if len(v3) > 0 {
                                    v3[0] = fi116
                                }
                                var v117 string = "hello"
                                _ = fi116
                                _ = v117
                            }
                        }
                        var v118 uint32 = 967
                        _ = v118
                    }
                }
            }
        default:
            var v119 uint8 = 128
            _ = v119
        }
        if (20 != 549) {
            if v4 {
                var v120 int16 = v2
                for fi121 := 0; fi121 < 4; fi121++ {
                    v2 = (v120 >> ((133) & 63))
                    _ = len(v1)
                    _ = fi121
                }
                _ = v120
            }
        } else {
            var v122 uint32 = 83
            _ = v122
        }
        _ = v10
    } else {
        fmt.Println(v2)
        fmt.Println(v9)
    }
    for fi123 := 0; fi123 < 8; fi123++ {
        v1[("" + "foo bar")] = 714
        if v4 {
            var v124 string = (("foo bar" + "hello") + "a")
            _ = v124
        } else {
            if false {
                v2 -= v2
                v4 = ((2 > 735) && (!((float32(float64(0.0)) * (0.0 / (1.5 * float32(1e10)))) > -2.25)))
            }
        }
        _ = fi123
    }
    switch 0 {
    case -121:
        var v125 float64 = 1.5
        _ = v125
    case 2147483647:
        var v126 int8 = 119
        _ = v126
    case -97:
        if v4 {
            v2 = (-(v2))
        } else {
            v9 = int32(v2)
        }
    default:
        fmt.Println(v9, v9, v4)
    }
    for fi127 := 0; fi127 < 8; fi127++ {
        if (18446744073709551615 >= 561) {
            var v128 uint64 = 810
            _ = v128
        }
        switch 42 {
        case 159:
            var v129 string = ("" + ("" + "hello"))
            _ = v129
        default:
            fmt.Println(fi127, v2, v4)
        }
        _ = fi127
    }
    v4 = v4
    v9 = (v9 | (v9 ^ v9))
    var v130 int32 = v9
    if (42 != 933) {
        if len(v3) > 0 {
            v3[0] = -2147483648
        }
        var v131 int64 = -984
        _ = v131
    } else {
        v4 = v4
        for fi132 := 0; fi132 < 10; fi132++ {
            var v133 int8 = -71
            _ = fi132
            _ = v133
        }
    }
    v9 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v130:", v130)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v130
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
