package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 string; F2 uint16 }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint16; F1 int }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 42
    var v3 string = ""
    v4 := make(map[int]string)
    v5 := make(map[string]bool)
    if (1.5 > 1e10) {
        fmt.Println(v3)
        var v6 bool = true
        v1[0] = v6
        _ = v6
    }
    v3 = v3
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: v3, F1: 67}
        var v9 int8 = -1
        v9 += (v9 / ((42) | 1))
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        _ = len(v1)
        _ = fi10
    }
    var v11 float64 = 3.14159
    if true {
        if ((32767 <= 65535) && true) {
            v2 = v2
            if false {
                v2 += (v2 << ((4294967295) & 63))
            } else {
                switch 1 {
                case -2147483648:
                    for fi12 := 0; fi12 < 12; fi12++ {
                        var v13 string = v3
                        var v14 uint8 = 188
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                case 9223372036854775807:
                    for fi15 := 0; fi15 < 10; fi15++ {
                        for fi16 := 0; fi16 < 1; fi16++ {
                            if false {
                                switch (("foo bar" + v3) + "") {
                                case "hello":
                                    var v17 uint32 = 1
                                    _ = v17
                                case "foo bar":
                                    var v18 int32 = 42
                                    _ = v18
                                case "":
                                    switch fi16 {
                                    case 128:
                                        var v19 uint32 = 65535
                                        _ = v19
                                    case 2147483647:
                                        var v20 int8 = (v2 & (v2 | 100))
                                        v21 := GenS2{F0: (fi16 << ((838) & 63)), F1: true}
                                        _ = v20
                                        _ = v21
                                    default:
                                        var v22 uint8 = 1
                                        _ = v22
                                    }
                                default:
                                    v23 := []int{(fi15 + 4294967295), fi16, (fi15 & (fi15 & 304))}
                                    _ = v23
                                }
                                v3 = "foo bar"
                                fmt.Println(fi15)
                            }
                            switch v3 {
                            case "hello":
                                v24 := []int{fi15}
                                _ = v24
                            }
                            if ("" >= "a") {
                                var v25 int32 = -1
                                v26 := []int{(fi16 & (fi15 % ((4294967295) | 1))), (^(fi16)), 32767}
                                v27 := GenS3{F0: (v25 - (v25 & -275))}
                                fmt.Println(v3, v2, fi15)
                                _ = v25
                                _ = v26
                                _ = v27
                            } else {
                                switch v3 {
                                case "\n\t":
                                    var v28 uint32 = 752
                                    _ = v28
                                case "":
                                    if true {
                                        v29 := []int{fi15, 2147483647}
                                        var v30 bool = true
                                        _ = v29
                                        _ = v30
                                    }
                                case "hello":
                                    switch v3 {
                                    case "hello":
                                        var v31 int64 = 0
                                        v32 := GenS4{F0: 42, F1: "a", F2: 1}
                                        _ = v31
                                        _ = v32
                                    }
                                }
                                var v33 uint16 = 1
                                if (fi15 != (fi16 ^ fi16)) {
                                    switch (fi15 - fi15) {
                                    case -9223372036854775808:
                                        v33 -= (^(v33))
                                    default:
                                        var v34 string = ("\n\t" + "a")
                                        _ = v34
                                    }
                                    var v35 float64 = v11
                                    _ = len(v4)
                                    _ = v35
                                }
                                _ = v33
                            }
                            _ = fi16
                        }
                        var v36 int64 = 1
                        _ = fi15
                        _ = v36
                    }
                case 273:
                    for fi37 := 0; fi37 < 14; fi37++ {
                        fmt.Println(v11)
                        _ = fi37
                    }
                default:
                    switch 42 {
                    case -1:
                        v38 := GenS1{F0: "foo bar", F1: 271}
                        _ = v38
                    }
                }
                v11 -= -2.25
            }
            if (116 < 277) {
                switch -2147483648 {
                case 127:
                    v39 := []int{9223372036854775807, -584, 534, 558}
                    _ = v39
                case 306:
                    for fi40 := 0; fi40 < 13; fi40++ {
                        var v41 float64 = ((v11 - float64(((1.5 + 1e10) / (float32(float64((0.0 * (-2.25 / (1e10 + ((((1e10 * ((float32(0.0) * ((0.1 * float32(float64(3.14159))) + ((1e10 / -2.25) - (1e10 - ((1e10 / ((((-2.25 - 0.1) - (3.14159 + 3.14159)) - ((1e10 * float32(v11)) + 1.5)) / 1.5)) + 3.14159))))) * 1.5)) + 0.0) / (1.5 / ((float32((float64(-2.25) - 1.5)) + 1.5) - (float32(((float64(0.0) + v11) * ((v11 + (v11 - -2.25)) - float64((-2.25 - float32((1e10 * v11))))))) / (3.14159 * (1e10 - 0.0)))))) * 3.14159)))))) - (1e10 - 1.5))))) + ((0.0 + -2.25) / ((1.5 * v11) / float64(3.14159))))
                        var v42 uint = 569
                        v43 := []int{fi40, -657, (fi40 / (((fi40 | (-(fi40)))) | 1))}
                        _ = fi40
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                }
            } else {
                if false {
                    fmt.Println(v2, v2)
                    switch 4294967295 {
                    case -817:
                        switch 216 {
                        case 360:
                            fmt.Println(v11, v11)
                        case 2147483647:
                            if (0 < 384) {
                                v44 := GenS4{F0: 234, F1: "", F2: 32767}
                                for fi45 := 0; fi45 < 13; fi45++ {
                                    v2 = 1
                                    _ = fi45
                                }
                                _ = v44
                            } else {
                                v46 := GenS2{F0: -2147483648, F1: (!(0 <= 42))}
                                v2 = 0
                                _ = v46
                            }
                        case 513:
                            var v47 uint = 187
                            _ = v47
                        default:
                            v11 += v11
                        }
                    }
                    v11 += (v11 + float64(float32((0.0 + v11))))
                } else {
                    fmt.Println(v11, v11)
                }
                if (true && true) {
                    v48 := []int{128}
                    _ = v48
                } else {
                    switch v3 {
                    case "foo bar":
                        v2 = 0
                    case "":
                        if ((127 != 0) || true) {
                            switch v3 {
                            case "hello":
                                fmt.Println(v11, v3, v2)
                            }
                            v2 = 1
                            v11 += v11
                        }
                    case "\n\t":
                        v49 := []int{1, 32767, 546}
                        _ = v49
                    default:
                        if ("foo bar" < v3) {
                            var v50 int16 = -1
                            var v51 float64 = (((float64((float32(float64((0.1 - 0.0))) + ((-2.25 / float32((v11 / v11))) / 1e10))) + (v11 * 0.1)) + (v11 / -2.25)) + v11)
                            var v52 float64 = 1e10
                            _ = v50
                            _ = v51
                            _ = v52
                        }
                    }
                    var v53 uint = 102
                    var v54 int = 1
                    _ = v53
                    _ = v54
                }
            }
        } else {
            if (!false) {
                v2 -= v2
            }
        }
        for fi55 := 0; fi55 < 18; fi55++ {
            v3 = v3
            _ = fi55
        }
        v56 := GenS4{F0: 2147483648, F1: v3, F2: 42}
        _ = v56
    } else {
        v2 += (v2 | v2)
        v2 += v2
    }
    if (false && true) {
        var v57 uint = 647
        for fi58 := 0; fi58 < 5; fi58++ {
            fmt.Println(v57, v57, v3)
            for fi59 := 0; fi59 < 7; fi59++ {
                var v60 int64 = 0
                var v61 uint8 = 0
                var v62 float32 = 1e10
                _ = fi59
                _ = v60
                _ = v61
                _ = v62
            }
            for fi63 := 0; fi63 < 3; fi63++ {
                v3 = v3
                if (1 != 9223372036854775807) {
                    if false {
                        var v64 uint64 = 367
                        _ = v64
                    } else {
                        if false {
                            v3 = v3
                            var v65 uint = 869
                            v66 := []int{398, (fi63 / ((504) | 1)), fi63, (fi63 - fi63)}
                            _ = v65
                            _ = v66
                        }
                        var v67 uint64 = 2
                        var v68 uint = (^(v57))
                        _ = v67
                        _ = v68
                    }
                    v69 := []int{286, fi63, (fi63 % ((fi58) | 1)), fi63}
                    v57 -= (v57 / ((65535) | 1))
                    switch "hello" {
                    case "hello":
                        switch fi58 {
                        case -9223372036854775808:
                            var v70 int = fi58
                            _ = v70
                        case 127:
                            var v71 uint8 = 255
                            _ = v71
                        case 4294967295:
                            v57 = (v57 * v57)
                        }
                    default:
                        if false {
                            delete(v4, fi58)
                            v2 -= v2
                            v2 = (v2 + (v2 / ((v2) | 1)))
                            fmt.Println(v57)
                        } else {
                            var v72 uint = v57
                            var v73 float32 = 0.0
                            _ = v72
                            _ = v73
                        }
                    }
                    _ = v69
                }
                _ = fi63
            }
            _ = fi58
        }
        _ = v57
    } else {
        var v74 int32 = -1
        _ = v74
    }
    v3 = "foo bar"
    switch v3 {
    case "hello":
        var v75 float64 = (v11 * v11)
        _ = v75
    case "foo bar":
        v76 := []int{970, -1}
        _ = v76
    case "\n\t":
        var v77 int64 = -1
        _ = v77
    }
    v11 -= v11
    if (42 > -2147483648) {
        var v78 uint16 = 1
        var v79 uint = 2
        if false {
            v11 -= v11
            v79 += uint(v78)
        } else {
            switch 4294967295 {
            case 9223372036854775807:
                for fi80 := 0; fi80 < 10; fi80++ {
                    v81 := v4[-313]
                    var v82 uint16 = (v78 | 255)
                    _ = fi80
                    _ = v81
                    _ = v82
                }
            case 128:
                v83 := []int{-2147483648, -1, 4294967295, 4294967295}
                _ = v83
            case 2147483647:
                v2 += (^(v2))
            default:
                switch -9223372036854775808 {
                case -9223372036854775808:
                    if false {
                        var v84 uint = uint(v78)
                        for fi85 := 0; fi85 < 18; fi85++ {
                            fmt.Println(v78, v3, v78)
                            v86 := []int{}
                            v87 := GenS1{F0: v3, F1: v78}
                            _ = fi85
                            _ = v86
                            _ = v87
                        }
                        v88 := GenS2{F0: 127, F1: (466 >= -3)}
                        _ = v84
                        _ = v88
                    }
                }
            }
        }
        _ = v78
        _ = v79
    } else {
        for fi89 := 0; fi89 < 8; fi89++ {
            v3 = (v3 + v3)
            _ = fi89
        }
        var v90 string = (v3 + v3)
        switch v90 {
        case "":
            fmt.Println(v11, v2, v3)
        }
        _ = v90
    }
    var v91 int64 = 1
    for fi92 := 0; fi92 < 17; fi92++ {
        switch v3 {
        case "":
            v2 = v2
        case "hello":
            var v93 float64 = v11
            _ = v93
        default:
            v11 = ((v11 * float64(3.14159)) / float64((float32(v11) - float32(0.0))))
        }
        fmt.Println(v11)
        _ = fi92
    }
    var v94 float32 = -2.25
    if false {
        switch "\n\t" {
        case "foo bar":
            for fi95 := 0; fi95 < 17; fi95++ {
                v96 := GenS5{F0: 128, F1: 42, F2: 1}
                var v97 int32 = 1
                _ = fi95
                _ = v96
                _ = v97
            }
        case "a":
            fmt.Println(v11)
        case "\n\t":
            v98 := GenS6{F0: 65535, F1: 128}
            _ = v98
        default:
            var v99 int16 = 0
            _ = v99
        }
        switch 2147483647 {
        case 1:
            var v100 int16 = 0
            _ = v100
        }
        v101 := GenS5{F0: 1, F1: 127, F2: 4294967295}
        _ = v101
    }
    if true {
        switch 4294967295 {
        case -1:
            if false {
                for fi102 := 0; fi102 < 15; fi102++ {
                    for fi103 := 0; fi103 < 11; fi103++ {
                        v3 = v3
                        if (!(("hello" + (v3 + v3)) != v3)) {
                            v3 = "hello"
                        } else {
                            var v104 int = fi103
                            var v105 uint8 = 1
                            switch (fi102 % ((fi102) | 1)) {
                            case 865:
                                var v106 uint32 = 2147483648
                                _ = v106
                            }
                            _ = v104
                            _ = v105
                        }
                        _ = fi103
                    }
                    v4[(-(fi102))] = v3
                    if (v2 < -8) {
                        var v107 int64 = (-(v91))
                        v108 := []int{}
                        _ = v107
                        _ = v108
                    } else {
                        for fi109 := 0; fi109 < 18; fi109++ {
                            var v110 int = 128
                            fmt.Println(fi109)
                            v94 = (v94 + (1.5 / v94))
                            var v111 bool = false
                            if (fi109 == v110) {
                                switch 9223372036854775807 {
                                case 601:
                                    v112 := []int{}
                                    _ = v112
                                case -1:
                                    v2 = v2
                                default:
                                    v113 := v1[42]
                                    _ = v113
                                }
                                if (v111 || false) {
                                    v114 := []int{}
                                    v111 = (-1 != 42)
                                    _ = v114
                                }
                                v115 := GenS7{F0: 458}
                                v116 := []int{(fi102 >> ((255) & 63)), (v110 * fi102), fi109}
                                if (3.14159 > v94) {
                                    switch v110 {
                                    case 1:
                                        fmt.Println(v3)
                                    case 2147483647:
                                        if v111 {
                                            var v117 uint = 65535
                                            _ = v117
                                        }
                                        if len(v116) >= 2 {
                                            v116 = v116[1:len(v116)]
                                        }
                                    }
                                } else {
                                    var v118 string = ("" + ("" + ("" + "\n\t")))
                                    v119 := GenS7{F0: 18446744073709551615}
                                    v120 := GenS6{F0: 42, F1: 128}
                                    _ = v118
                                    _ = v119
                                    _ = v120
                                }
                                _ = v115
                                _ = v116
                            } else {
                                var v121 int32 = 0
                                var v122 uint32 = 2147483648
                                if ((v111 || ((^(v122)) < (v122 * v122))) || v111) {
                                    var v123 float64 = v11
                                    var v124 int32 = v121
                                    v11 = (((-2.25 - -2.25) * v11) / v123)
                                    _ = v123
                                    _ = v124
                                }
                                _ = v121
                                _ = v122
                            }
                            _ = fi109
                            _ = v110
                            _ = v111
                        }
                        v91 -= 42
                    }
                    _ = fi102
                }
            }
        default:
            if (v91 >= v91) {
                fmt.Println(v2)
                v11 = 3.14159
                _ = len(v5)
            }
        }
        if (0.0 == ((v94 / ((v94 * v94) - v94)) + float32((v11 - (0.1 + (float64(0.0) - ((v11 / v11) * (((((1e10 + (0.1 / v11)) * v11) * 0.0) / (v11 * v11)) * v11)))))))) {
            v125 := GenS8{F0: 255, F1: 32767}
            var v126 float32 = v94
            var v127 int16 = -32768
            _ = v125
            _ = v126
            _ = v127
        } else {
            v91 -= 1
        }
        v2 -= (v2 % (((v2 | (-(v2)))) | 1))
    }
    v128 := GenS8{F0: 2, F1: -9223372036854775808}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v91:", v91)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v11
    _ = v128
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v91
    _ = v94
}
