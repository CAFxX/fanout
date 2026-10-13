package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 uint8; F2 int16 }
type GenS7 struct { F0 uint; F1 int32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int64; F1 string; F2 string }

func main() {
    var v1 int64 = -1
    var v2 uint8 = 128
    v3 := make(map[uint32]int)
    v4 := make(map[int]uint32)
    if ((0 != 2147483647) || false) {
        var v5 uint64 = 18446744073709551615
        for fi6 := 0; fi6 < 11; fi6++ {
            v7 := []int{(fi6 + (fi6 & 127)), fi6, 32767}
            v5 = v5
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 12; fi8++ {
            switch ("a" + "a") {
            case "a":
                if (!((-193 == -582) && true)) {
                    for fi9 := 0; fi9 < 9; fi9++ {
                        v1 += (v1 % ((v1) | 1))
                        _ = fi9
                    }
                } else {
                    v2 -= v2
                    if (1 >= -128) {
                        var v10 int16 = -13
                        var v11 uint = 42
                        v2 += 1
                        _ = v10
                        _ = v11
                    } else {
                        v12 := GenS1{F0: 1}
                        var v13 string = ("" + "hello")
                        var v14 bool = (-54 <= -70)
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    fmt.Println(v1, v5, v1)
                }
            case "\n\t":
                switch "hello" {
                case "foo bar":
                    var v15 int64 = (v1 >> ((65535) & 63))
                    _ = v15
                default:
                    if false {
                        v16 := []int{(fi8 - (fi8 % ((fi8) | 1))), (fi8 << ((255) & 63)), (fi8 % ((fi8) | 1)), -1}
                        _ = v16
                    }
                }
            case "":
                switch 9223372036854775807 {
                case 1:
                    v17 := v3[282]
                    _ = v17
                case 0:
                    for fi18 := 0; fi18 < 10; fi18++ {
                        switch ("foo bar" + "") {
                        case "hello":
                            var v19 float64 = 3.14159
                            _ = v19
                        case "\n\t":
                            v20 := v3[2147483647]
                            _ = v20
                        }
                        var v21 uint64 = 323
                        v5 = v5
                        v22 := GenS2{F0: fi8, F1: 1, F2: 2147483647}
                        fmt.Println(v21)
                        var v23 float32 = ((0.1 - ((1.5 * float32(((0.0 + 1e10) / -2.25))) / 1.5)) * 0.0)
                        _ = fi18
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                default:
                    if (true || (!true)) {
                        switch (fi8 & -1) {
                        case 127:
                            switch "\n\t" {
                            case "foo bar":
                                var v24 int16 = 1
                                _ = v24
                            case "hello":
                                for fi25 := 0; fi25 < 1; fi25++ {
                                    var v26 int = -9223372036854775808
                                    var v27 uint = 42
                                    var v28 bool = true
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                            case "a":
                                fmt.Println(v5, v2)
                            default:
                                var v29 int16 = 42
                                _ = v29
                            }
                        case -273:
                            _ = len(v4)
                        }
                    } else {
                        v5 = (v5 / ((v5) | 1))
                        if false {
                            if false {
                                v5 = v5
                                v1 = (v1 << ((1) & 63))
                                var v30 uint8 = uint8(v5)
                                _ = v30
                            } else {
                                v1 = (v1 / (((^(v1))) | 1))
                            }
                        }
                        if true {
                            for fi31 := 0; fi31 < 14; fi31++ {
                                switch "hello" {
                                case "a":
                                    v32 := []int{(fi31 << ((0) & 63)), fi8}
                                    _ = v32
                                default:
                                    var v33 uint = 959
                                    _ = v33
                                }
                                var v34 int64 = (v1 | (v1 | 9223372036854775807))
                                _ = fi31
                                _ = v34
                            }
                            v1 -= (v1 - 203)
                            v4[fi8] = 1
                        } else {
                            var v35 float64 = (0.1 * 3.14159)
                            fmt.Println(v1)
                            if false {
                                fmt.Println(v35, v2, v2)
                                fmt.Println(v1, v2, v1)
                                v35 -= 0.1
                            } else {
                                var v36 string = ("a" + ("foo bar" + "\n\t"))
                                if false {
                                    fmt.Println(v1)
                                    v2 += v2
                                    switch "\n\t" {
                                    case "foo bar":
                                        v37 := v3[71]
                                        _ = v37
                                    default:
                                        v2 -= 2
                                    }
                                } else {
                                    v35 = (((v35 * (v35 - 0.1)) + v35) / v35)
                                }
                                v38 := []int{(fi8 >> ((1) & 63))}
                                _ = v36
                                _ = v38
                            }
                            _ = v35
                        }
                    }
                }
            }
            _ = fi8
        }
        _ = v5
    } else {
        var v39 int16 = 1000
        if true {
            v3[576] = 4294967295
            for fi40 := 0; fi40 < 7; fi40++ {
                v41 := v4[(fi40 % (((fi40 % (((fi40 * (fi40 & 1))) | 1))) | 1))]
                var v42 int16 = v39
                v43 := []int{fi40, (fi40 - (-(fi40)))}
                _ = fi40
                _ = v41
                _ = v42
                _ = v43
            }
        }
        if (((42 != 2) || true) && (96 < 0)) {
            var v44 int64 = 595
            fmt.Println(v2, v2, v2)
            v4[-397] = 4294967295
            _ = v44
        }
        _ = v39
    }
    switch 0 {
    case 127:
        var v45 uint16 = 65535
        _ = v45
    case 707:
        delete(v3, 461)
    }
    v2 -= v2
    v1 = (v1 | -1)
    var v46 int = int(v1)
    var v47 int = (v46 ^ v46)
    for fi48 := 0; fi48 < 9; fi48++ {
        v49 := GenS3{F0: 4294967295, F1: 253, F2: 920}
        switch (("hello" + "hello") + ("foo bar" + ("" + "a"))) {
        case "hello":
            for fi50 := 0; fi50 < 2; fi50++ {
                fmt.Println(v1)
                _ = fi50
            }
        default:
            v51 := GenS3{F0: 73, F1: 2, F2: 2}
            _ = v51
        }
        switch "foo bar" {
        case "\n\t":
            switch "a" {
            case "":
                v52 := GenS2{F0: (fi48 % ((9223372036854775807) | 1)), F1: (v1 & v1), F2: 4294967295}
                _ = v52
            default:
                v46 += v46
            }
        case "a":
            v47 += v47
        case "hello":
            v4[127] = 1
        default:
            var v53 int16 = 1000
            _ = v53
        }
        _ = fi48
        _ = v49
    }
    var v54 int16 = 42
    var v55 uint64 = 2
    v2 -= v2
    for fi56 := 0; fi56 < 1; fi56++ {
        fmt.Println(v1, v46, v47)
        fmt.Println(v55)
        var v57 uint16 = uint16(v2)
        _ = fi56
        _ = v57
    }
    for fi58 := 0; fi58 < 11; fi58++ {
        v59 := GenS4{F0: 0, F1: fi58}
        fmt.Println(v47, v1, fi58)
        v46 = (v46 % ((v46) | 1))
        _ = fi58
        _ = v59
    }
    if false {
        v55 -= (v55 % (((v55 | 4294967295)) | 1))
    }
    switch v46 {
    case -9223372036854775808:
        v60 := GenS3{F0: v55, F1: (v55 / ((v55) | 1)), F2: 65535}
        _ = v60
    case -501:
        v2 += 0
    case 127:
        switch "" {
        case "hello":
            v54 += v54
        case "a":
            switch "\n\t" {
            case "hello":
                if ((v47 | 32767) != (v46 ^ (-(v46)))) {
                    v55 -= v55
                } else {
                    for fi61 := 0; fi61 < 7; fi61++ {
                        for fi62 := 0; fi62 < 18; fi62++ {
                            v63 := []int{}
                            _ = fi62
                            _ = v63
                        }
                        fmt.Println(v47, v46, v55)
                        var v64 int64 = -428
                        _ = fi61
                        _ = v64
                    }
                    v55 = v55
                    fmt.Println(v2)
                }
            case "a":
                fmt.Println(v54, v1, v55)
            case "foo bar":
                for fi65 := 0; fi65 < 9; fi65++ {
                    if false {
                        fmt.Println(v55)
                    } else {
                        var v66 int32 = 0
                        var v67 int64 = (v1 & int64(v2))
                        _ = v66
                        _ = v67
                    }
                    for fi68 := 0; fi68 < 13; fi68++ {
                        var v69 uint8 = (^(v2))
                        v2 -= 255
                        v70 := GenS4{F0: (v1 << ((464) & 63)), F1: (v46 * -160)}
                        _ = fi68
                        _ = v69
                        _ = v70
                    }
                    _ = fi65
                }
            }
        }
    default:
        switch 0 {
        case 0:
            v55 = 849
        case 127:
            var v71 int = v46
            _ = v71
        }
    }
    for fi72 := 0; fi72 < 19; fi72++ {
        v54 = (v54 ^ 21)
        var v73 uint8 = (^(v2))
        _ = fi72
        _ = v73
    }
    var v74 int64 = v1
    for fi75 := 0; fi75 < 19; fi75++ {
        var v76 uint16 = 42
        v55 = v55
        switch "" {
        case "a":
            if (0 < (^(v76))) {
                for fi77 := 0; fi77 < 6; fi77++ {
                    v78 := GenS5{F0: 548}
                    if ((v2 % ((v2) | 1)) >= 54) {
                        fmt.Println(v46, v47, v76)
                        v79 := GenS2{F0: 4294967295, F1: 42, F2: 357}
                        v80 := GenS6{F0: -2147483648, F1: v2, F2: 42}
                        _ = v79
                        _ = v80
                    } else {
                        if (!false) {
                            if (true || (0 < 0)) {
                                v4[32768] = 2147483647
                                v81 := GenS2{F0: int(v2), F1: -1, F2: 65535}
                                _ = v81
                            } else {
                                if (!("\n\t" > "foo bar")) {
                                    switch "foo bar" {
                                    case "":
                                        if false {
                                            var v82 int16 = v54
                                            v83 := []int{}
                                            _ = v82
                                            _ = v83
                                        }
                                    default:
                                        v84 := v4[127]
                                        _ = v84
                                    }
                                    _ = len(v4)
                                    var v85 float32 = 0.0
                                    delete(v3, 4294967295)
                                    _ = v85
                                } else {
                                    delete(v3, 860)
                                    v4[(fi75 ^ (fi75 + (-(v46))))] = 2147483648
                                    v86 := []int{}
                                    _ = v86
                                }
                            }
                            switch (("hello" + "foo bar") + "") {
                            case "foo bar":
                                var v87 int64 = 1
                                _ = v87
                            case "":
                                v88 := GenS7{F0: 585, F1: 1}
                                _ = v88
                            }
                            v89 := []int{int(v74), (-(fi75)), -850}
                            _ = v89
                        } else {
                            delete(v4, (fi75 * fi75))
                            v90 := GenS4{F0: v1, F1: -9223372036854775808}
                            _ = len(v4)
                            fmt.Println(v2, v54)
                            var v91 uint16 = uint16(v55)
                            v76 = (v91 - (^(v91)))
                            v2 = 50
                            v92 := []int{128, -1, (v47 % (((v47 | v47)) | 1))}
                            v93 := GenS8{F0: "\n\t"}
                            _ = v90
                            _ = v91
                            _ = v92
                            _ = v93
                        }
                    }
                    _ = fi77
                    _ = v78
                }
                v94 := GenS2{F0: (v46 % ((v46) | 1)), F1: (v74 / (((v1 ^ 1)) | 1)), F2: 1}
                fmt.Println(v54, v54, fi75)
                _ = v94
            } else {
                for fi95 := 0; fi95 < 13; fi95++ {
                    v96 := GenS9{F0: -2147483648}
                    if (!(1 < 100)) {
                        _ = len(v3)
                        v97 := GenS3{F0: v55, F1: 2, F2: 2147483647}
                        _ = v97
                    } else {
                        v74 = 9223372036854775807
                        if false {
                            var v98 int = (-(fi75))
                            _ = v98
                        }
                        v99 := []int{v47}
                        v54 = v54
                        _ = v99
                    }
                    v47 = (fi75 * (v46 - (v46 / (((-(fi95))) | 1))))
                    v54 = v54
                    v55 -= v55
                    fmt.Println(v76, v47)
                    var v100 uint16 = 539
                    var v101 int32 = -204
                    var v102 bool = (0 > 526)
                    _ = fi95
                    _ = v100
                    _ = v101
                    _ = v102
                    _ = v96
                }
                for fi103 := 0; fi103 < 12; fi103++ {
                    fmt.Println(v54)
                    if ("\n\t" != "foo bar") {
                        fmt.Println(v46, v46, v74)
                        if (!false) {
                            _ = len(v4)
                            fmt.Println(fi103, fi103)
                            v104 := GenS10{F0: int64(v55), F1: "a", F2: "a"}
                            _ = len(v3)
                            switch "hello" {
                            case "\n\t":
                                var v105 int = (^(v46))
                                _ = v105
                            }
                            _ = v104
                        }
                        v106 := GenS7{F0: 65535, F1: int32(v54)}
                        _ = v106
                    }
                    if true {
                        v55 = uint64(v76)
                        if true {
                            if (32767 <= v76) {
                                var v107 float64 = 1.5
                                _ = v107
                            } else {
                                v108 := []int{}
                                v109 := v4[2147483647]
                                _ = v108
                                _ = v109
                            }
                            var v110 int32 = 1
                            var v111 int16 = (v54 % ((-687) | 1))
                            if (!true) {
                                v112 := GenS6{F0: (v110 << ((42) & 63)), F1: 166, F2: (v54 / (((v54 ^ int16(v76))) | 1))}
                                var v113 int = (v46 - fi103)
                                fmt.Println(v110)
                                _ = v112
                                _ = v113
                            }
                            _ = v110
                            _ = v111
                        }
                    } else {
                        v47 += v47
                        v46 = (fi103 + (-(fi75)))
                        if true {
                            if true {
                                var v114 int64 = v74
                                var v115 uint64 = (v55 + 901)
                                _ = v114
                                _ = v115
                            } else {
                                v2 = (v2 << ((4294967295) & 63))
                            }
                            var v116 uint32 = 357
                            v47 += (v46 & 127)
                            _ = v116
                        } else {
                            if (862 < 148) {
                                var v117 uint64 = 42
                                var v118 string = (("\n\t" + "a") + "foo bar")
                                _ = len(v3)
                                _ = v117
                                _ = v118
                            }
                        }
                    }
                    var v119 float32 = 0.1
                    _ = fi103
                    _ = v119
                }
            }
        case "\n\t":
            for fi120 := 0; fi120 < 11; fi120++ {
                fmt.Println(fi75, v2)
                _ = fi120
            }
        default:
            for fi121 := 0; fi121 < 1; fi121++ {
                v122 := GenS1{F0: 65535}
                var v123 int = (v46 >> ((460) & 63))
                _ = fi121
                _ = v122
                _ = v123
            }
        }
        _ = fi75
        _ = v76
    }
    switch "" {
    case "\n\t":
        fmt.Println(v47)
    case "foo bar":
        v1 = (v1 | (-(v1)))
    case "":
        for fi124 := 0; fi124 < 13; fi124++ {
            v47 = v47
            _ = fi124
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v54
    _ = v55
    _ = v74
}
