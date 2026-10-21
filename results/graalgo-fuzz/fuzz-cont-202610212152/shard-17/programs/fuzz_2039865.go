package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint32; F1 int16; F2 int8 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int; F1 uint16 }
type GenS9 struct { F0 int; F1 int32 }
type GenS10 struct { F0 int; F1 uint16 }
type GenS11 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int8; F1 int }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 87
    var v3 bool = (true && false)
    v2 = v2
    if v3 {
        for fi4 := 0; fi4 < 9; fi4++ {
            v5 := GenS1{F0: "hello", F1: 348, F2: 0}
            var v6 float32 = 0.1
            _ = fi4
            _ = v5
            _ = v6
        }
    }
    v7 := []int{127, -1, 32767, -379}
    _ = len(v1)
    v8 := GenS1{F0: "foo bar", F1: 2, F2: 42}
    if v3 {
        v9 := GenS1{F0: "", F1: 536, F2: 42}
        if false {
            if v3 {
                v10 := v1[-842]
                _ = v10
            } else {
                if (((0.0 / (0.0 + 3.14159)) > 1.5) || (425 != 104)) {
                    var v11 int = 9223372036854775807
                    fmt.Println(v3, v3)
                    fmt.Println(v3, v2, v11)
                    _ = v11
                } else {
                    if true {
                        if false {
                            v3 = true
                            for fi12 := 0; fi12 < 9; fi12++ {
                                if v3 {
                                    v3 = (127 > 551)
                                    switch "foo bar" {
                                    case "foo bar":
                                        for fi13 := 0; fi13 < 20; fi13++ {
                                            var v14 int32 = -2147483648
                                            _ = fi13
                                            _ = v14
                                        }
                                    case "hello":
                                        if v3 {
                                            var v15 int64 = -9223372036854775808
                                            if (v3 && (-722 > 1)) {
                                                switch (fi12 + -1) {
                                                case 42:
                                                    v2 += 187
                                                case -192:
                                                    if v3 {
                                                        var v16 uint32 = 4294967295
                                                        _ = v16
                                                    } else {
                                                        var v17 float64 = float64((-2.25 - (float32(0.0) * 0.0)))
                                                        var v18 int8 = 0
                                                        v19 := GenS2{F0: 127, F1: v3}
                                                        _ = v17
                                                        _ = v18
                                                        _ = v19
                                                    }
                                                case 44:
                                                    switch 32768 {
                                                    case -1:
                                                        if v3 {
                                                            v20 := GenS3{F0: 42, F1: 248, F2: 38}
                                                            if false {
                                                                v21 := GenS1{F0: "foo bar", F1: 4294967295, F2: 42}
                                                                if (!v3) {
                                                                    if v3 {
                                                                        var v22 uint16 = 24
                                                                        var v23 string = (("hello" + "hello") + "\n\t")
                                                                        v1[(fi12 << ((385) & 63))] = 2147483648
                                                                        _ = v22
                                                                        _ = v23
                                                                    } else {
                                                                        for fi24 := 0; fi24 < 16; fi24++ {
                                                                            v3 = true
                                                                            var v25 int64 = -164
                                                                            _ = fi24
                                                                            _ = v25
                                                                        }
                                                                        if v3 {
                                                                            v26 := GenS1{F0: "foo bar", F1: 4294967295, F2: 4294967295}
                                                                            _ = v26
                                                                        } else {
                                                                            v27 := GenS1{F0: "\n\t", F1: 1, F2: 1}
                                                                            var v28 string = (("" + ("a" + "hello")) + "")
                                                                            _ = v27
                                                                            _ = v28
                                                                        }
                                                                    }
                                                                    for fi29 := 0; fi29 < 18; fi29++ {
                                                                        if len(v7) >= 2 {
                                                                            v7 = v7[1:len(v7)]
                                                                        }
                                                                        var v30 uint16 = 90
                                                                        _ = fi29
                                                                        _ = v30
                                                                    }
                                                                }
                                                                switch ("a" + "a") {
                                                                case "hello":
                                                                    if false {
                                                                        if v3 {
                                                                            fmt.Println(v2)
                                                                            _ = len(v1)
                                                                            v31 := GenS2{F0: (v2 << ((65535) & 63)), F1: false}
                                                                            _ = v31
                                                                        }
                                                                        var v32 uint64 = 260
                                                                        v2 += (v2 + 128)
                                                                        _ = v32
                                                                    } else {
                                                                        v15 -= (^(v15))
                                                                        v33 := GenS1{F0: "", F1: 4294967295, F2: 4294967295}
                                                                        _ = v33
                                                                    }
                                                                }
                                                                _ = v21
                                                            } else {
                                                                var v34 bool = v3
                                                                v15 = int64(fi12)
                                                                _ = v34
                                                            }
                                                            _ = v20
                                                        } else {
                                                            v15 += (^(v15))
                                                        }
                                                    default:
                                                        v15 = (v15 >> ((65535) & 63))
                                                    }
                                                }
                                                v15 += -78
                                                _ = len(v7)
                                            } else {
                                                v35 := GenS4{F0: 2, F1: 128}
                                                v15 = v15
                                                _ = v35
                                            }
                                            _ = v15
                                        } else {
                                            v7 = append(v7, 4294967295)
                                        }
                                    case "\n\t":
                                        for fi36 := 0; fi36 < 11; fi36++ {
                                            v37 := GenS5{F0: 295}
                                            var v38 uint64 = 1
                                            v39 := GenS3{F0: 2147483648, F1: 32767, F2: -128}
                                            _ = fi36
                                            _ = v37
                                            _ = v38
                                            _ = v39
                                        }
                                    default:
                                        for fi40 := 0; fi40 < 10; fi40++ {
                                            switch fi12 {
                                            case 4294967295:
                                                if (fi12 >= 32767) {
                                                    v2 = (^(v2))
                                                    var v41 uint64 = 393
                                                    if (!v3) {
                                                        fmt.Println(v3)
                                                        v42 := GenS3{F0: 0, F1: -784, F2: 100}
                                                        if v3 {
                                                            fmt.Println(fi12)
                                                        } else {
                                                            if v3 {
                                                                v43 := GenS1{F0: "", F1: 641, F2: (v41 >> ((18446744073709551615) & 63))}
                                                                if (v3 || (0.1 <= 0.0)) {
                                                                    v44 := GenS1{F0: "foo bar", F1: (v41 / ((0) | 1)), F2: v41}
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        v2 += (v2 & v2)
                                                                    case "\n\t":
                                                                        switch "hello" {
                                                                        case "\n\t":
                                                                            var v45 uint64 = 2
                                                                            _ = v45
                                                                        case "":
                                                                            v1[(fi12 + (fi40 ^ (fi40 | fi12)))] = 42
                                                                        case "foo bar":
                                                                            fmt.Println(v41)
                                                                        }
                                                                    case "":
                                                                        if (uint(v41) < 42) {
                                                                            fmt.Println(fi12)
                                                                            fmt.Println(v41)
                                                                        }
                                                                    }
                                                                    v41 = v41
                                                                    _ = v44
                                                                }
                                                                v3 = v3
                                                                _ = v43
                                                            } else {
                                                                if v3 {
                                                                    v46 := GenS1{F0: ("\n\t" + "\n\t"), F1: 914, F2: v41}
                                                                    _ = v46
                                                                } else {
                                                                    switch ("hello" + "\n\t") {
                                                                    case "a":
                                                                        v47 := GenS6{F0: 2, F1: 707}
                                                                        _ = v47
                                                                    default:
                                                                        var v48 int16 = -32768
                                                                        _ = v48
                                                                    }
                                                                    if (!false) {
                                                                        fmt.Println(v3)
                                                                        v2 = (v2 % (((v2 - v2)) | 1))
                                                                    }
                                                                    v2 += (^(v2))
                                                                    v2 += v2
                                                                    v2 = v2
                                                                }
                                                                switch fi40 {
                                                                case 724:
                                                                    var v49 int = -911
                                                                    if (v2 < v2) {
                                                                        var v50 int32 = 1
                                                                        if (-972 < 0) {
                                                                            var v51 float64 = 0.1
                                                                            fmt.Println(fi12)
                                                                            _ = v51
                                                                        }
                                                                        _ = v50
                                                                    }
                                                                    if (v3 && (0 != 100)) {
                                                                        if v3 {
                                                                            _ = len(v1)
                                                                            if v3 {
                                                                                var v52 int64 = 42
                                                                                fmt.Println(fi40)
                                                                                _ = v52
                                                                            }
                                                                        }
                                                                        v41 += (^(v41))
                                                                    } else {
                                                                        if ((v41 >> ((255) & 63)) > v41) {
                                                                            var v53 uint16 = 665
                                                                            switch fi12 {
                                                                            case -153:
                                                                                v1[fi40] = 42
                                                                            case 32768:
                                                                                fmt.Println(v53, v41)
                                                                            default:
                                                                                v49 = 999
                                                                            }
                                                                            _ = v53
                                                                        } else {
                                                                            var v54 uint64 = (^(v41))
                                                                            _ = v54
                                                                        }
                                                                        v3 = (41 == (v2 ^ (v2 << ((4294967295) & 63))))
                                                                    }
                                                                    _ = v49
                                                                case 42:
                                                                    var v55 string = "\n\t"
                                                                    if len(v7) > 0 {
                                                                        v7[0] = (fi12 + fi12)
                                                                    }
                                                                    if (127 >= 127) {
                                                                        if v3 {
                                                                            fmt.Println(v55, v41)
                                                                            fmt.Println(fi12, v2, v41)
                                                                        } else {
                                                                            if ((0.1 * 1.5) >= (-2.25 - 3.14159)) {
                                                                                v56 := GenS2{F0: v2, F1: true}
                                                                                if v3 {
                                                                                    if false {
                                                                                        var v57 int = 32768
                                                                                        if (1e10 != (3.14159 + 3.14159)) {
                                                                                            v3 = v3
                                                                                            var v58 float64 = 0.0
                                                                                            _ = v58
                                                                                        } else {
                                                                                            var v59 int = fi12
                                                                                            switch v55 {
                                                                                            case "":
                                                                                                if v3 {
                                                                                                    var v60 bool = true
                                                                                                    if v60 {
                                                                                                        v3 = false
                                                                                                        var v61 float32 = -2.25
                                                                                                        v62 := GenS1{F0: "", F1: 18446744073709551615, F2: (v41 - 162)}
                                                                                                        _ = v61
                                                                                                        _ = v62
                                                                                                    }
                                                                                                    if v3 {
                                                                                                        var v63 uint = 895
                                                                                                        v57 += 42
                                                                                                        _ = v63
                                                                                                    } else {
                                                                                                        if v3 {
                                                                                                            var v64 bool = (v59 > 42)
                                                                                                            v2 += (v2 | 24)
                                                                                                            v65 := GenS7{F0: 171}
                                                                                                            _ = v64
                                                                                                            _ = v65
                                                                                                        }
                                                                                                        delete(v1, fi40)
                                                                                                    }
                                                                                                    _ = v60
                                                                                                }
                                                                                            case "foo bar":
                                                                                                switch "" {
                                                                                                case "hello":
                                                                                                    if v3 {
                                                                                                        var v66 bool = false
                                                                                                        _ = v66
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    v3 = false
                                                                                                case "":
                                                                                                    var v67 uint32 = uint32(v41)
                                                                                                    _ = v67
                                                                                                }
                                                                                            case "\n\t":
                                                                                                v68 := GenS7{F0: 273}
                                                                                                _ = v68
                                                                                            default:
                                                                                                if false {
                                                                                                    var v69 int8 = 127
                                                                                                    _ = v69
                                                                                                } else {
                                                                                                    var v70 int64 = 9223372036854775807
                                                                                                    _ = v70
                                                                                                }
                                                                                            }
                                                                                            _ = v59
                                                                                        }
                                                                                        _ = v57
                                                                                    }
                                                                                }
                                                                                _ = v56
                                                                            } else {
                                                                                v55 = v55
                                                                            }
                                                                            if v3 {
                                                                                v41 -= (v41 >> ((2) & 63))
                                                                                if len(v7) > 0 {
                                                                                    v7[0] = 0
                                                                                }
                                                                            }
                                                                            var v71 int32 = 613
                                                                            var v72 uint8 = v2
                                                                            _ = v71
                                                                            _ = v72
                                                                        }
                                                                    }
                                                                    _ = v55
                                                                case 32768:
                                                                    var v73 uint32 = 1
                                                                    _ = v73
                                                                default:
                                                                    var v74 int32 = -2147483648
                                                                    delete(v1, fi12)
                                                                    var v75 string = ""
                                                                    _ = v74
                                                                    _ = v75
                                                                }
                                                                var v76 float64 = ((float64(0.0) + 1e10) * 0.1)
                                                                _ = v76
                                                            }
                                                        }
                                                        if (0 < uint16(v2)) {
                                                            var v77 int8 = 100
                                                            if (!true) {
                                                                v78 := GenS8{F0: fi12, F1: 2}
                                                                if (!v3) {
                                                                    if (2 < 65535) {
                                                                        var v79 uint64 = v41
                                                                        v80 := v1[fi12]
                                                                        var v81 float32 = (1.5 / 1e10)
                                                                        _ = v79
                                                                        _ = v80
                                                                        _ = v81
                                                                    } else {
                                                                        switch 127 {
                                                                        case -801:
                                                                            switch (fi40 >> ((728) & 63)) {
                                                                            case 2147483647:
                                                                                if (18446744073709551615 < (v41 / (((v41 & v41)) | 1))) {
                                                                                    v77 = 1
                                                                                }
                                                                            case -707:
                                                                                if v3 {
                                                                                    var v82 uint32 = 533
                                                                                    _ = v82
                                                                                } else {
                                                                                    fmt.Println(v2, fi40)
                                                                                    delete(v1, (fi12 * (fi12 + (fi40 ^ fi40))))
                                                                                }
                                                                            }
                                                                        case 2147483647:
                                                                            v2 -= (v2 | 0)
                                                                            var v83 int32 = 0
                                                                            _ = v83
                                                                        default:
                                                                            if true {
                                                                                switch 9223372036854775807 {
                                                                                case 0:
                                                                                    v77 += (v77 << ((65535) & 63))
                                                                                case 1:
                                                                                    switch 32767 {
                                                                                    case 4294967295:
                                                                                        var v84 uint64 = v41
                                                                                        _ = v84
                                                                                    default:
                                                                                        v2 -= (v2 << ((uint(v2)) & 63))
                                                                                    }
                                                                                default:
                                                                                    var v85 float64 = (float64(-2.25) / 1.0)
                                                                                    _ = v85
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    if v3 {
                                                                        v41 = (^(v41))
                                                                        v3 = (2 == uint32(v41))
                                                                        var v86 int32 = 2147483647
                                                                        var v87 uint16 = 946
                                                                        _ = v86
                                                                        _ = v87
                                                                    }
                                                                    v3 = (true && (v2 > (v2 - (v2 | v2))))
                                                                    switch "hello" {
                                                                    case "":
                                                                        var v88 string = "a"
                                                                        _ = v88
                                                                    case "\n\t":
                                                                        var v89 float64 = float64(3.14159)
                                                                        _ = v89
                                                                    case "hello":
                                                                        v90 := GenS2{F0: v2, F1: true}
                                                                        _ = v90
                                                                    }
                                                                } else {
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        _ = len(v7)
                                                                    case "hello":
                                                                        var v91 int32 = -277
                                                                        _ = v91
                                                                    case "a":
                                                                        var v92 float64 = ((0.0 * 1.5) * (0.0 / (1.5 * -2.25)))
                                                                        _ = v92
                                                                    default:
                                                                        var v93 uint16 = 32767
                                                                        _ = v93
                                                                    }
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        v41 = (^(v41))
                                                                    }
                                                                }
                                                                _ = v78
                                                            }
                                                            switch 967 {
                                                            case -1:
                                                                v3 = false
                                                            case 2147483647:
                                                                if v3 {
                                                                    v94 := v1[2147483647]
                                                                    var v95 uint64 = v41
                                                                    v95 = (v41 - 4294967295)
                                                                    v7 = append(v7, fi12)
                                                                    v77 += v77
                                                                    if len(v7) >= 2 {
                                                                        v7 = v7[1:len(v7)]
                                                                    }
                                                                    _ = v94
                                                                    _ = v95
                                                                } else {
                                                                    var v96 float64 = (0.1 - float64(0.0))
                                                                    if true {
                                                                        v41 = (v41 >> ((628) & 63))
                                                                    }
                                                                    var v97 float32 = 0.0
                                                                    fmt.Println(v3, v2, v97)
                                                                    var v98 uint64 = (v41 | (^(v41)))
                                                                    _ = v96
                                                                    _ = v97
                                                                    _ = v98
                                                                }
                                                                var v99 float32 = (0.0 / (1e10 / (3.14159 - (3.14159 / 1.5))))
                                                                _ = v99
                                                            default:
                                                                switch fi40 {
                                                                case -1:
                                                                    var v100 int = (fi40 & 32767)
                                                                    _ = v100
                                                                default:
                                                                    if (v77 != (v77 * (v77 | (v77 % ((int8(fi12)) | 1))))) {
                                                                        v101 := GenS9{F0: 127, F1: 0}
                                                                        var v102 float64 = 0.1
                                                                        switch "" {
                                                                        case "hello":
                                                                            if (((v41 * v41) < v41) && v3) {
                                                                                if (!(735 < 42)) {
                                                                                    var v103 float32 = (((float32(v102) + -2.25) * -2.25) / (((1e10 + (-2.25 * 0.1)) - 0.0) / 1.0))
                                                                                    if v3 {
                                                                                        if (!(!(!true))) {
                                                                                            switch "foo bar" {
                                                                                            case "\n\t":
                                                                                                v104 := GenS7{F0: 2}
                                                                                                _ = v104
                                                                                            case "foo bar":
                                                                                                switch "a" {
                                                                                                case "hello":
                                                                                                    var v105 string = ""
                                                                                                    _ = v105
                                                                                                default:
                                                                                                    v2 += (^(v2))
                                                                                                }
                                                                                                var v106 float64 = v102
                                                                                                switch fi40 {
                                                                                                case 4294967295:
                                                                                                    var v107 int8 = v77
                                                                                                    _ = v107
                                                                                                }
                                                                                                switch (fi12 / (((fi12 - (fi40 ^ fi40))) | 1)) {
                                                                                                case 9223372036854775807:
                                                                                                    v3 = true
                                                                                                default:
                                                                                                    if ((^(v2)) < 246) {
                                                                                                        var v108 int64 = 9223372036854775807
                                                                                                        if v3 {
                                                                                                            v3 = (56 != 811)
                                                                                                            switch fi40 {
                                                                                                            case 32768:
                                                                                                                var v109 uint16 = 32767
                                                                                                                _ = v109
                                                                                                            default:
                                                                                                                var v110 uint8 = 255
                                                                                                                _ = v110
                                                                                                            }
                                                                                                            v111 := GenS6{F0: 2, F1: 278}
                                                                                                            _ = v111
                                                                                                        }
                                                                                                        if true {
                                                                                                            v7 = append(v7, 0)
                                                                                                            switch (fi12 << ((904) & 63)) {
                                                                                                            case -9223372036854775808:
                                                                                                                _ = len(v7)
                                                                                                                if false {
                                                                                                                    v112 := GenS9{F0: 0, F1: 42}
                                                                                                                    v113 := GenS5{F0: (^(v41))}
                                                                                                                    var v114 int8 = (^(v77))
                                                                                                                    v115 := GenS10{F0: 548, F1: 65535}
                                                                                                                    var v116 uint = 497
                                                                                                                    _ = v112
                                                                                                                    _ = v113
                                                                                                                    _ = v114
                                                                                                                    _ = v115
                                                                                                                    _ = v116
                                                                                                                }
                                                                                                                if (v3 || (v3 && (v106 > (float64(v103) * v106)))) {
                                                                                                                    switch fi40 {
                                                                                                                    case 386:
                                                                                                                        if v3 {
                                                                                                                            v41 = (^(v41))
                                                                                                                            var v117 int32 = 0
                                                                                                                            if len(v7) > 0 {
                                                                                                                                v7[0] = -1
                                                                                                                            }
                                                                                                                            _ = v117
                                                                                                                        }
                                                                                                                    case -105:
                                                                                                                        v108 = v108
                                                                                                                    case 128:
                                                                                                                        delete(v1, (fi12 - 4294967295))
                                                                                                                    }
                                                                                                                    fmt.Println(v2)
                                                                                                                }
                                                                                                                if (v2 <= v2) {
                                                                                                                    v7 = append(v7, (-(fi12)))
                                                                                                                }
                                                                                                                v41 = 2
                                                                                                            case 4294967295:
                                                                                                                if (fi12 != 4294967295) {
                                                                                                                    var v118 uint8 = (v2 + 249)
                                                                                                                    _ = v118
                                                                                                                }
                                                                                                            default:
                                                                                                                var v119 int16 = 32767
                                                                                                                _ = v119
                                                                                                            }
                                                                                                            var v120 int64 = v108
                                                                                                            _ = v120
                                                                                                        }
                                                                                                        _ = v108
                                                                                                    }
                                                                                                }
                                                                                                if len(v7) >= 2 {
                                                                                                    v7 = v7[1:len(v7)]
                                                                                                }
                                                                                                _ = v106
                                                                                            }
                                                                                        } else {
                                                                                            if (9223372036854775807 == 350) {
                                                                                                v2 -= 42
                                                                                            } else {
                                                                                                v3 = (2147483647 <= 1)
                                                                                            }
                                                                                            var v121 int32 = 0
                                                                                            if false {
                                                                                                v1[2147483647] = 378
                                                                                            }
                                                                                            _ = v121
                                                                                        }
                                                                                    } else {
                                                                                        switch "\n\t" {
                                                                                        case "a":
                                                                                            if true {
                                                                                                var v122 int16 = 1
                                                                                                _ = v122
                                                                                            } else {
                                                                                                if v3 {
                                                                                                    v103 = float32(v102)
                                                                                                } else {
                                                                                                    if ((!(1 == 2)) || v3) {
                                                                                                        var v123 float32 = float32(0.0)
                                                                                                        if v3 {
                                                                                                            if (!true) {
                                                                                                                v7 = append(v7, (fi12 + fi40))
                                                                                                            }
                                                                                                            v124 := GenS4{F0: uint64(v2), F1: 774}
                                                                                                            switch int(v41) {
                                                                                                            case -1:
                                                                                                                var v125 int8 = 51
                                                                                                                _ = v125
                                                                                                            }
                                                                                                            _ = v124
                                                                                                        } else {
                                                                                                            var v126 int32 = 2147483647
                                                                                                            _ = v126
                                                                                                        }
                                                                                                        var v127 uint16 = 255
                                                                                                        _ = v123
                                                                                                        _ = v127
                                                                                                    } else {
                                                                                                        var v128 int32 = 100
                                                                                                        _ = v128
                                                                                                    }
                                                                                                    var v129 uint16 = 2
                                                                                                    var v130 uint64 = (v41 >> ((65535) & 63))
                                                                                                    _ = v129
                                                                                                    _ = v130
                                                                                                }
                                                                                                if (v3 && (!(792 != 2147483647))) {
                                                                                                    switch fi12 {
                                                                                                    case 127:
                                                                                                        var v131 int32 = -1
                                                                                                        _ = v131
                                                                                                    }
                                                                                                    var v132 int8 = v77
                                                                                                    _ = v132
                                                                                                }
                                                                                                var v133 int16 = -32768
                                                                                                v134 := GenS9{F0: fi12, F1: -293}
                                                                                                _ = v133
                                                                                                _ = v134
                                                                                            }
                                                                                            fmt.Println(v2, fi40)
                                                                                            v7 = append(v7, (fi40 - 4294967295))
                                                                                        default:
                                                                                            v102 = v102
                                                                                        }
                                                                                    }
                                                                                    _ = v103
                                                                                }
                                                                            }
                                                                        case "\n\t":
                                                                            v135 := GenS3{F0: 4294967295, F1: 35, F2: (v77 >> ((339) & 63))}
                                                                            _ = v135
                                                                        }
                                                                        _ = v101
                                                                        _ = v102
                                                                    } else {
                                                                        if false {
                                                                            if (!(!(true || (true && false)))) {
                                                                                if (-42 != -1) {
                                                                                    v41 = v41
                                                                                    switch ("\n\t" + ("\n\t" + "foo bar")) {
                                                                                    case "\n\t":
                                                                                        var v136 string = ("foo bar" + "\n\t")
                                                                                        _ = v136
                                                                                    case "":
                                                                                        v137 := GenS11{F0: (-1 < -1), F1: (v2 >> ((1) & 63)), F2: 914}
                                                                                        _ = v137
                                                                                    case "hello":
                                                                                        v138 := GenS1{F0: "\n\t", F1: (v41 | 42), F2: (v41 % (((v41 / ((v41) | 1))) | 1))}
                                                                                        _ = v138
                                                                                    default:
                                                                                        var v139 uint64 = (^(v41))
                                                                                        _ = v139
                                                                                    }
                                                                                    switch "hello" {
                                                                                    case "foo bar":
                                                                                        fmt.Println(v2, v3)
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                v2 = v2
                                                                            }
                                                                        } else {
                                                                            v3 = (255 == 0)
                                                                            v7 = append(v7, fi40)
                                                                        }
                                                                        v3 = ((v2 & uint8(v41)) <= 2)
                                                                        if len(v7) > 0 {
                                                                            v7[0] = fi40
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            _ = v77
                                                        } else {
                                                            if v3 {
                                                                v3 = (true || (897 == 759))
                                                            }
                                                        }
                                                        v2 = 1
                                                        v2 += v2
                                                        _ = len(v7)
                                                        _ = v42
                                                    }
                                                    v41 = 176
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        var v140 uint16 = 420
                                                        _ = v140
                                                    case "":
                                                        var v141 uint16 = 268
                                                        _ = v141
                                                    }
                                                    if len(v7) >= 2 {
                                                        v7 = v7[1:len(v7)]
                                                    }
                                                    switch (fi12 % (((^(fi40))) | 1)) {
                                                    case 493:
                                                        switch "a" {
                                                        case "a":
                                                            fmt.Println(v41, v3, v2)
                                                        default:
                                                            fmt.Println(fi12, v41, v41)
                                                        }
                                                    default:
                                                        v3 = false
                                                    }
                                                    _ = v41
                                                }
                                            case 194:
                                                var v142 uint8 = (v2 << ((255) & 63))
                                                if v3 {
                                                    v142 = (v142 / ((0) | 1))
                                                    v7 = append(v7, (fi12 / (((fi40 & (fi12 >> ((858) & 63)))) | 1)))
                                                } else {
                                                    var v143 string = "foo bar"
                                                    v142 = 128
                                                    _ = v143
                                                }
                                                _ = v142
                                            }
                                            _ = fi40
                                        }
                                    }
                                    if v3 {
                                        for fi144 := 0; fi144 < 10; fi144++ {
                                            delete(v1, (fi12 % ((fi12) | 1)))
                                            _ = fi144
                                        }
                                        delete(v1, 128)
                                    } else {
                                        fmt.Println(v2)
                                        if v3 {
                                            v3 = true
                                            v145 := GenS6{F0: 255, F1: 330}
                                            _ = v145
                                        } else {
                                            switch (fi12 % (((^(fi12))) | 1)) {
                                            case 32767:
                                                v146 := GenS12{F0: 0}
                                                _ = v146
                                            default:
                                                switch "foo bar" {
                                                case "hello":
                                                    _ = len(v1)
                                                default:
                                                    v147 := GenS3{F0: 4294967295, F1: 204, F2: -68}
                                                    _ = v147
                                                }
                                            }
                                        }
                                        for fi148 := 0; fi148 < 17; fi148++ {
                                            if (!v3) {
                                                v149 := GenS7{F0: 857}
                                                fmt.Println(v3)
                                                _ = v149
                                            }
                                            v150 := v1[-1]
                                            if (v3 || v3) {
                                                if false {
                                                    v150 += (v150 | 4294967295)
                                                    v2 = (v2 << ((18446744073709551615) & 63))
                                                    fmt.Println(v3)
                                                    fmt.Println(fi148, fi12, v3)
                                                    v150 = v150
                                                }
                                                var v151 string = ""
                                                _ = v151
                                            } else {
                                                switch 0 {
                                                case 32768:
                                                    _ = len(v1)
                                                case 198:
                                                    fmt.Println(v3)
                                                }
                                                fmt.Println(v3, fi148)
                                            }
                                            var v152 uint = 0
                                            var v153 int64 = 932
                                            _ = fi148
                                            _ = v150
                                            _ = v152
                                            _ = v153
                                        }
                                    }
                                }
                                _ = fi12
                            }
                        }
                    }
                }
            }
            if v3 {
                var v154 int32 = 1
                v2 = 0
                for fi155 := 0; fi155 < 3; fi155++ {
                    for fi156 := 0; fi156 < 8; fi156++ {
                        var v157 int8 = 42
                        var v158 uint16 = 400
                        if false {
                            var v159 uint64 = 233
                            _ = v159
                        }
                        _ = fi156
                        _ = v157
                        _ = v158
                    }
                    fmt.Println(fi155, v3)
                    _ = fi155
                }
                _ = v154
            } else {
                fmt.Println(v2)
            }
            fmt.Println(v3)
        }
        v3 = false
        _ = v9
    } else {
        for fi160 := 0; fi160 < 17; fi160++ {
            var v161 int64 = 0
            _ = fi160
            _ = v161
        }
        v162 := GenS13{F0: 1, F1: -35}
        _ = v162
    }
    if v3 {
        v1[-9223372036854775808] = 1
        if v3 {
            v3 = (118 < 1)
            if ((3.14159 + (3.14159 - float32(((0.0 - (1.5 * -2.25)) / 1e10)))) < 3.14159) {
                if (!false) {
                    var v163 string = ""
                    switch -1 {
                    case -9223372036854775808:
                        v163 = "foo bar"
                    }
                    if true {
                        switch v163 {
                        case "foo bar":
                            var v164 int32 = 1
                            _ = v164
                        }
                        var v165 uint = 255
                        for fi166 := 0; fi166 < 6; fi166++ {
                            for fi167 := 0; fi167 < 19; fi167++ {
                                switch v163 {
                                case "hello":
                                    var v168 float64 = 0.1
                                    _ = v168
                                }
                                v165 = v165
                                _ = fi167
                            }
                            if v3 {
                                for fi169 := 0; fi169 < 20; fi169++ {
                                    v2 -= (^(v2))
                                    _ = fi169
                                }
                            }
                            var v170 float64 = -2.25
                            _ = fi166
                            _ = v170
                        }
                        _ = v165
                    }
                    _ = v163
                }
            } else {
                v2 = (^(v2))
                v3 = false
                v171 := GenS13{F0: 127, F1: -201}
                _ = v171
            }
        }
    }
    switch "foo bar" {
    case "foo bar":
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    case "hello":
        var v172 int = -564
        _ = v172
    case "a":
        v173 := GenS11{F0: (-237 >= -639), F1: 1, F2: 652}
        _ = v173
    }
    for fi174 := 0; fi174 < 20; fi174++ {
        var v175 string = "hello"
        _ = fi174
        _ = v175
    }
    var v176 float32 = ((-2.25 + 3.14159) + 3.14159)
    var v177 int64 = -1
    v1[4294967295] = 175
    _ = len(v1)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v176:", v176)
    fmt.Println("v177:", v177)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v176
    _ = v177
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
