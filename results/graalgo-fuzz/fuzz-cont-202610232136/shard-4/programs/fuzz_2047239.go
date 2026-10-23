package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool; F1 int; F2 int32 }
type GenS8 struct { F0 uint8; F1 int8; F2 int64 }
type GenS9 struct { F0 uint16; F1 uint32 }
type GenS10 struct { F0 uint64; F1 uint32 }
type GenS11 struct { F0 string }
func genFn0(p0 int8, p1 uint32) uint16 {
    p1 = p1
    var v4 int16 = 0
    switch "\n\t" {
    case "\n\t":
        v5 := []int{1, 2147483647}
        _ = v5
    case "foo bar":
        fmt.Println(p1, p0)
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v4
    return 42
}
func genFn1(p0 int8) uint {
    if true {
        var v6 int8 = p0
        _ = v6
    } else {
        fmt.Println(p0)
        p0 += 127
        switch 1 {
        case 0:
            if (!false) {
                switch 32767 {
                case 2147483647:
                    v7 := GenS1{F0: 1, F1: 2}
                    _ = v7
                default:
                    v8 := GenS1{F0: 8, F1: 2}
                    _ = v8
                }
            }
        case -91:
            p0 += p0
        }
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 uint = 1
            v11 -= 0
            var v12 float64 = -2.25
            v13 := []int{fi9, (fi10 & (fi10 << ((v11) & 63))), (fi10 - fi10)}
            fmt.Println(v11)
            var v14 uint64 = 4294967295
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi9
    }
    _ = p0
    return 1
}

func main() {
    v1 := []int{128, 32767}
    var v2 uint16 = 1
    v3 := make(map[uint32]string)
    v15 := GenS2{F0: 65535, F1: 255, F2: 42}
    fmt.Println(v2)
    var v16 uint8 = 128
    _ = len(v1)
    fmt.Println(v16, v2)
    var v17 int8 = -3
    var v18 int = 729
    v19 := GenS3{F0: -613, F1: (!false)}
    v2 = 0
    v18 = v18
    for fi20 := 0; fi20 < 1; fi20++ {
        for fi21 := 0; fi21 < 1; fi21++ {
            fmt.Println(fi20, fi21)
            var v22 int = (fi21 >> ((42) & 63))
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 12; fi23++ {
            switch "hello" {
            case "\n\t":
                v2 -= (v2 + genFn0(v17, 49))
            case "":
                var v24 bool = (32767 < -32768)
                _ = v24
            case "hello":
                v16 = uint8(v2)
            default:
                v3[64] = "a"
            }
            if (v17 == v17) {
                var v25 int32 = int32(v2)
                fmt.Println(fi20)
                var v26 uint64 = 4294967295
                _ = v25
                _ = v26
            } else {
                var v27 float32 = float32(0.0)
                var v28 int8 = 100
                _ = v27
                _ = v28
            }
            var v29 int16 = -32768
            _ = fi23
            _ = v29
        }
        _ = fi20
    }
    v30 := GenS2{F0: 42, F1: 42, F2: v16}
    switch v18 {
    case -484:
        for fi31 := 0; fi31 < 17; fi31++ {
            switch "" {
            case "\n\t":
                v2 += 887
            case "a":
                var v32 string = "hello"
                _ = v32
            }
            switch "hello" {
            case "hello":
                var v33 uint16 = genFn0(v17, 242)
                _ = v33
            case "":
                if (-2.25 > ((((0.1 / 1.0) / (0.1 * (-2.25 + ((1e10 - ((3.14159 * 0.1) + 1.5)) + 3.14159)))) + 3.14159) / (-2.25 / 1.5))) {
                    _ = len(v1)
                } else {
                    for fi34 := 0; fi34 < 20; fi34++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v35 := GenS1{F0: (v17 / (((v17 >> ((18446744073709551615) & 63))) | 1)), F1: 255}
                        _ = fi34
                        _ = v35
                    }
                    if (!false) {
                        if true {
                            v36 := GenS4{F0: v18, F1: (v17 | v17)}
                            var v37 string = "hello"
                            _ = v36
                            _ = v37
                        }
                        v38 := GenS5{F0: (v17 | v17)}
                        var v39 float64 = 1e10
                        _ = v38
                        _ = v39
                    }
                    if (265 == 2147483647) {
                        if false {
                            for fi40 := 0; fi40 < 6; fi40++ {
                                if (false && false) {
                                    v16 += v16
                                    v17 += v17
                                    fmt.Println(fi31, fi31)
                                } else {
                                    var v41 int16 = 32767
                                    _ = v41
                                }
                                var v42 uint32 = 362
                                _ = fi40
                                _ = v42
                            }
                        }
                        var v43 uint64 = 674
                        _ = v43
                    }
                }
            case "foo bar":
                var v44 int32 = 0
                _ = v44
            }
            _ = fi31
        }
    case 0:
        if false {
            if (v16 <= 2) {
                for fi45 := 0; fi45 < 14; fi45++ {
                    switch fi45 {
                    case 2147483647:
                        v17 -= v17
                    case 1:
                        v16 += (^(v16))
                    case 32768:
                        switch fi45 {
                        case 446:
                            fmt.Println(v18, v18, v16)
                        }
                    default:
                        if (0 == 1) {
                            v16 += 42
                        }
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi45
                }
                v46 := GenS4{F0: v18, F1: -1}
                v47 := GenS6{F0: v2}
                _ = v46
                _ = v47
            } else {
                var v48 float64 = -2.25
                var v49 uint64 = 18446744073709551615
                _ = v48
                _ = v49
            }
        } else {
            for fi50 := 0; fi50 < 3; fi50++ {
                switch "hello" {
                case "\n\t":
                    switch "a" {
                    case "foo bar":
                        _ = len(v1)
                    case "\n\t":
                        if ((v2 << ((1) & 63)) >= (^(v2))) {
                            if true {
                                switch v18 {
                                case -1:
                                    if true {
                                        if len(v1) > 0 {
                                            v1[0] = (v18 >> ((18446744073709551615) & 63))
                                        }
                                    }
                                }
                                if true {
                                    var v51 float64 = 1.5
                                    _ = v51
                                }
                                switch "\n\t" {
                                case "foo bar":
                                    v52 := GenS4{F0: (-(v18)), F1: v17}
                                    _ = v52
                                default:
                                    var v53 uint8 = (^(v16))
                                    _ = v53
                                }
                            }
                            for fi54 := 0; fi54 < 2; fi54++ {
                                v55 := GenS1{F0: v17, F1: genFn0((v17 + v17), 38)}
                                _ = fi54
                                _ = v55
                            }
                        }
                    case "hello":
                        for fi56 := 0; fi56 < 12; fi56++ {
                            var v57 uint64 = 42
                            var v58 uint = 65535
                            fmt.Println(v57)
                            _ = fi56
                            _ = v57
                            _ = v58
                        }
                    }
                case "a":
                    var v59 int32 = 1
                    _ = v59
                case "":
                    if false {
                        if false {
                            v60 := GenS6{F0: (^(v2))}
                            _ = v60
                        } else {
                            for fi61 := 0; fi61 < 2; fi61++ {
                                if len(v1) > 0 {
                                    v1[0] = 155
                                }
                                _ = fi61
                            }
                        }
                        v62 := v3[75]
                        _ = v62
                    } else {
                        switch fi50 {
                        case -66:
                            v2 = (^(v2))
                        case 252:
                            var v63 uint = 223
                            _ = v63
                        case 695:
                            for fi64 := 0; fi64 < 14; fi64++ {
                                fmt.Println(v2, fi64)
                                if (!(-1 > 161)) {
                                    v3[133] = "hello"
                                } else {
                                    var v65 int32 = -127
                                    if true {
                                        fmt.Println(v18)
                                        var v66 int8 = (v17 * v17)
                                        _ = v66
                                    } else {
                                        v2 = (v2 * (v2 | 733))
                                        var v67 float32 = ((-2.25 - 0.0) / ((1.5 * 1e10) * float32(3.14159)))
                                        v67 -= v67
                                        var v68 int32 = (v65 << ((553) & 63))
                                        var v69 uint64 = 18446744073709551615
                                        _ = v67
                                        _ = v68
                                        _ = v69
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v65
                                }
                                var v70 uint8 = 2
                                _ = fi64
                                _ = v70
                            }
                        default:
                            if len(v1) > 0 {
                                v1[0] = v18
                            }
                        }
                        if true {
                            for fi71 := 0; fi71 < 11; fi71++ {
                                v72 := GenS7{F0: (42 < -9223372036854775808), F1: fi71, F2: -61}
                                if (float32(0.1) > 1e10) {
                                    if false {
                                        if len(v1) > 0 {
                                            v1[0] = (fi71 ^ fi50)
                                        }
                                        fmt.Println(v2)
                                        switch "a" {
                                        case "foo bar":
                                            _ = len(v1)
                                        default:
                                            fmt.Println(v2, v16)
                                        }
                                        var v73 float64 = (1.5 + (1.5 / 1e10))
                                        v74 := GenS2{F0: 42, F1: (v16 - v16), F2: 0}
                                        _ = v73
                                        _ = v74
                                    } else {
                                        var v75 int8 = 127
                                        if false {
                                            if true {
                                                switch fi50 {
                                                case 4294967295:
                                                    fmt.Println(v17, v2)
                                                    var v76 bool = (false && true)
                                                    var v77 int = (fi50 / ((fi71) | 1))
                                                    _ = v76
                                                    _ = v77
                                                default:
                                                    v78 := GenS8{F0: v16, F1: -103, F2: 835}
                                                    _ = v78
                                                }
                                                v79 := GenS2{F0: 880, F1: 158, F2: v16}
                                                var v80 bool = (74 > -9223372036854775808)
                                                _ = v79
                                                _ = v80
                                            }
                                            var v81 uint64 = 42
                                            var v82 uint16 = genFn0((v17 << ((18446744073709551615) & 63)), 4294967295)
                                            _ = v81
                                            _ = v82
                                        }
                                        _ = v75
                                    }
                                    var v83 float32 = (((1e10 + float32(-2.25)) + 0.1) / (3.14159 - -2.25))
                                    v16 = 2
                                    fmt.Println(v2, v18, v16)
                                    _ = v83
                                }
                                _ = fi71
                                _ = v72
                            }
                            if true {
                                v16 = v16
                                for fi84 := 0; fi84 < 11; fi84++ {
                                    if (!(genFn1(v17) > 600)) {
                                        var v85 uint = 2
                                        if (!false) {
                                            if false {
                                                v18 -= v18
                                                switch 4294967295 {
                                                case 947:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                case -594:
                                                    var v86 uint32 = 151
                                                    _ = v86
                                                case 4294967295:
                                                    v17 += v17
                                                }
                                            }
                                            v85 -= 0
                                            if false {
                                                fmt.Println(fi84, v85, v2)
                                            } else {
                                                v87 := GenS9{F0: v2, F1: 343}
                                                var v88 uint16 = 32767
                                                _ = v87
                                                _ = v88
                                            }
                                            v89 := GenS2{F0: 42, F1: 239, F2: 1}
                                            v3[789] = (("a" + "foo bar") + "\n\t")
                                            _ = v89
                                        } else {
                                            if true {
                                                delete(v3, 42)
                                                if false {
                                                    v3[484] = "foo bar"
                                                    fmt.Println(v2, v16, v17)
                                                    var v90 uint = v85
                                                    v1 = append(v1, (^(fi50)))
                                                    if ((v16 * 210) < uint8(v2)) {
                                                        v17 -= (v17 << ((v90) & 63))
                                                        var v91 int16 = -464
                                                        _ = v91
                                                    }
                                                    _ = v90
                                                } else {
                                                    var v92 uint16 = (v2 % (((v2 * v2)) | 1))
                                                    fmt.Println(v18, fi50)
                                                    v93 := GenS9{F0: (v92 & (v92 - (^(v2)))), F1: 270}
                                                    _ = v92
                                                    _ = v93
                                                }
                                            }
                                            v85 = 0
                                        }
                                        _ = v85
                                    }
                                    if false {
                                        if true {
                                            var v94 float64 = 1.5
                                            var v95 uint64 = 627
                                            switch fi50 {
                                            case 0:
                                                var v96 int8 = 127
                                                _ = v96
                                            case -1:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            default:
                                                fmt.Println(fi50, v95)
                                            }
                                            v95 = v95
                                            _ = v94
                                            _ = v95
                                        }
                                    }
                                    v16 = (v16 << ((18446744073709551615) & 63))
                                    v97 := GenS10{F0: 2, F1: 4294967295}
                                    _ = fi84
                                    _ = v97
                                }
                                var v98 string = "foo bar"
                                _ = v98
                            } else {
                                v18 -= int(v2)
                            }
                        }
                    }
                }
                switch "\n\t" {
                case "a":
                    v2 -= (v2 | 0)
                case "foo bar":
                    var v99 string = "a"
                    _ = v99
                case "hello":
                    v17 += int8(v16)
                }
                _ = fi50
            }
            var v100 uint32 = 2147483647
            _ = v100
        }
    case -165:
        var v101 uint8 = v16
        _ = v101
    default:
        var v102 int32 = -826
        _ = v102
    }
    if ((18446744073709551615 < 244) && true) {
        if false {
            for fi103 := 0; fi103 < 9; fi103++ {
                fmt.Println(v18, v17, v17)
                _ = fi103
            }
            for fi104 := 0; fi104 < 16; fi104++ {
                fmt.Println(v17, fi104, v16)
                var v105 float64 = 0.0
                _ = fi104
                _ = v105
            }
            var v106 int16 = 42
            _ = v106
        }
        if true {
            for fi107 := 0; fi107 < 5; fi107++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v108 := GenS11{F0: "a"}
                v3[2147483647] = ("\n\t" + "\n\t")
                _ = fi107
                _ = v108
            }
            var v109 uint = 255
            _ = v109
        } else {
            var v110 int16 = -1
            var v111 float64 = 0.0
            _ = v110
            _ = v111
        }
        var v112 int32 = 1
        _ = v112
    }
    v2 -= v2
    var v113 int8 = v17
    if ((genFn0(42, 1) >= (v2 & (v2 << ((genFn1(v17)) & 63)))) && false) {
        switch "hello" {
        case "hello":
            fmt.Println(v16, v17)
        }
        v114 := GenS3{F0: int64(v18), F1: false}
        v115 := GenS9{F0: (v2 - (v2 & 1)), F1: 1}
        _ = v114
        _ = v115
    } else {
        if (42 != 1) {
            for fi116 := 0; fi116 < 1; fi116++ {
                v16 -= 0
                v113 = v17
                _ = fi116
            }
            if true {
                v117 := GenS1{F0: (-(v113)), F1: v2}
                for fi118 := 0; fi118 < 19; fi118++ {
                    switch (fi118 * (v18 * (^(v18)))) {
                    case 2147483647:
                        var v119 float64 = 1e10
                        _ = v119
                    case -2147483648:
                        switch "foo bar" {
                        case "":
                            switch "a" {
                            case "\n\t":
                                var v120 float32 = 3.14159
                                _ = v120
                            }
                        }
                    case -220:
                        v2 = v2
                    }
                    var v121 float32 = 1e10
                    var v122 string = "\n\t"
                    _ = fi118
                    _ = v121
                    _ = v122
                }
                _ = v117
            } else {
                var v123 uint = 65535
                _ = v123
            }
        } else {
            for fi124 := 0; fi124 < 11; fi124++ {
                v17 = (v113 & 42)
                v18 = (v18 << ((201) & 63))
                _ = fi124
            }
            switch int(v2) {
            case 2147483647:
                var v125 uint8 = (^(v16))
                _ = v125
            case -2147483648:
                if (!false) {
                    v18 += (v18 << ((2) & 63))
                }
            case 231:
                v3[2] = ("a" + (((("" + "a") + (("\n\t" + "a") + "foo bar")) + "") + ("a" + "foo bar")))
            default:
                _ = len(v1)
            }
            for fi126 := 0; fi126 < 12; fi126++ {
                if false {
                    for fi127 := 0; fi127 < 19; fi127++ {
                        var v128 uint8 = (v16 * v16)
                        _ = fi127
                        _ = v128
                    }
                    fmt.Println(fi126, v18)
                } else {
                    var v129 int32 = -2147483648
                    var v130 uint8 = 1
                    _ = v129
                    _ = v130
                }
                var v131 int = (fi126 * v18)
                _ = fi126
                _ = v131
            }
        }
    }
    v2 -= v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v113:", v113)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v113
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
}
