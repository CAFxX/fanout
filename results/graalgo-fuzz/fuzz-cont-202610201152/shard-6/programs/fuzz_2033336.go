package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 bool; F1 int8 }
type GenS10 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int32; F1 uint; F2 int }
type GenS13 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS14 struct { F0 string; F1 int64; F2 int }
type GenS15 struct { F0 uint; F1 uint8; F2 uint16 }

func main() {
    var v1 int = 0
    var v2 int64 = int64(v1)
    var v3 uint = 2
    var v4 float32 = -2.25
    var v5 uint8 = 12
    v1 += -407
    v3 += v3
    var v6 string = ("\n\t" + "\n\t")
    v7 := []int{}
    v4 -= (v4 - 1e10)
    v2 -= v2
    switch "foo bar" {
    case "":
        var v8 string = v6
        _ = v8
    case "foo bar":
        v9 := GenS1{F0: v2, F1: -740, F2: (((!(!((v2 | 0) > v2))) && true) || (!false))}
        _ = v9
    default:
        var v10 string = (v6 + "hello")
        _ = v10
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if (579 > 235) {
        v11 := GenS2{F0: 788}
        switch v6 {
        case "foo bar":
            fmt.Println(v1, v3, v4)
        }
        if len(v7) > 0 {
            v7[0] = v1
        }
        _ = v11
    }
    if (65535 <= 2) {
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 int = (fi12 ^ (v1 | 4294967295))
            for fi14 := 0; fi14 < 6; fi14++ {
                switch (v13 << (((^(v3))) & 63)) {
                case 127:
                    var v15 uint16 = 277
                    _ = v15
                }
                if ((!false) || (!true)) {
                    var v16 float64 = ((1e10 / (((((1.5 * 1e10) - (-2.25 / 3.14159)) / 3.14159) * float64(v4)) * float64(v4))) + 1.5)
                    var v17 int = 4294967295
                    if true {
                        v16 += v16
                    } else {
                        v1 = (fi14 >> ((791) & 63))
                    }
                    v13 = -963
                    v18 := GenS3{F0: 216, F1: -2147483648}
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    fmt.Println(v1)
                    v2 -= 42
                    fmt.Println(fi14)
                }
                var v19 int32 = 42
                v1 -= (-(fi14))
                _ = fi14
                _ = v19
            }
            _ = fi12
            _ = v13
        }
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        _ = len(v7)
        for fi21 := 0; fi21 < 14; fi21++ {
            switch (((v6 + v6) + v6) + "hello") {
            case "foo bar":
                if (v4 > 3.14159) {
                    v22 := GenS1{F0: 9223372036854775807, F1: -1, F2: true}
                    if false {
                        var v23 string = "a"
                        if (v6 != "a") {
                            if (true && false) {
                                _ = len(v7)
                                var v24 uint64 = 42
                                if (58 == 65535) {
                                    v3 += 902
                                    switch (fi20 * fi21) {
                                    case 32768:
                                        var v25 int64 = (v2 >> ((834) & 63))
                                        _ = v25
                                    case 2147483647:
                                        if false {
                                            if (false || false) {
                                                if (638 >= 0) {
                                                    if len(v7) > 0 {
                                                        v7[0] = (fi21 | (-(fi21)))
                                                    }
                                                    switch (v23 + ((v6 + "\n\t") + (("foo bar" + v23) + "a"))) {
                                                    case "":
                                                        fmt.Println(v5, v4)
                                                    case "hello":
                                                        v26 := GenS4{F0: (v3 ^ (v3 - uint(v24))), F1: 0}
                                                        _ = v26
                                                    default:
                                                        var v27 uint64 = (^(v24))
                                                        _ = v27
                                                    }
                                                } else {
                                                    v28 := GenS1{F0: v2, F1: -1, F2: false}
                                                    if false {
                                                        if true {
                                                            v3 = 602
                                                        } else {
                                                            if true {
                                                                var v29 int = (fi20 | int(v2))
                                                                var v30 uint64 = v24
                                                                _ = v29
                                                                _ = v30
                                                            } else {
                                                                var v31 uint32 = 587
                                                                _ = v31
                                                            }
                                                            var v32 int32 = -2147483648
                                                            _ = v32
                                                        }
                                                        switch fi20 {
                                                        case 1:
                                                            v4 = -2.25
                                                        case 127:
                                                            if true {
                                                                v33 := GenS1{F0: v2, F1: 156, F2: true}
                                                                _ = v33
                                                            } else {
                                                                var v34 uint32 = 4294967295
                                                                if true {
                                                                    v6 = ("hello" + ("hello" + v23))
                                                                    var v35 uint32 = v34
                                                                    _ = v35
                                                                }
                                                                var v36 int = (fi21 | 1)
                                                                _ = v34
                                                                _ = v36
                                                            }
                                                        case 650:
                                                            fmt.Println(fi21, v2, v6)
                                                        default:
                                                            v5 += 42
                                                        }
                                                        v37 := GenS5{F0: 1, F1: 446}
                                                        _ = v37
                                                    } else {
                                                        v24 = (v24 - v24)
                                                        fmt.Println(v24)
                                                        fmt.Println(v23)
                                                        var v38 uint8 = v5
                                                        if false {
                                                            v7 = append(v7, 897)
                                                        } else {
                                                            if true {
                                                                v39 := GenS6{F0: v24, F1: (v3 << ((255) & 63))}
                                                                _ = v39
                                                            } else {
                                                                v6 = v23
                                                                var v40 int64 = v2
                                                                _ = v40
                                                            }
                                                            v1 = fi20
                                                            v41 := GenS2{F0: v24}
                                                            _ = v41
                                                        }
                                                        _ = v38
                                                    }
                                                    _ = v28
                                                }
                                                fmt.Println(v2)
                                            } else {
                                                v5 = 127
                                                var v42 float32 = (1e10 + 0.0)
                                                _ = v42
                                            }
                                            switch fi20 {
                                            case -9223372036854775808:
                                                switch (-(v1)) {
                                                case 523:
                                                    v4 += v4
                                                    switch v6 {
                                                    case "\n\t":
                                                        var v43 int64 = 1
                                                        _ = v43
                                                    case "hello":
                                                        fmt.Println(v24, v3)
                                                    }
                                                case 42:
                                                    if len(v7) >= 2 {
                                                        v7 = v7[1:len(v7)]
                                                    }
                                                case -408:
                                                    switch -9223372036854775808 {
                                                    case -324:
                                                        v44 := GenS7{F0: 1, F1: v6}
                                                        _ = v44
                                                    case 502:
                                                        var v45 int64 = v2
                                                        _ = v45
                                                    case 586:
                                                        var v46 float32 = 0.0
                                                        _ = v46
                                                    }
                                                }
                                            case 2147483647:
                                                switch "" {
                                                case "":
                                                    var v47 bool = ((v6 <= "\n\t") && true)
                                                    _ = v47
                                                case "hello":
                                                    v2 += (v2 ^ -763)
                                                case "\n\t":
                                                    v6 = ((v6 + v23) + ("a" + v23))
                                                    v2 = (^(v2))
                                                }
                                                if true {
                                                    fmt.Println(v5, v24)
                                                    var v48 bool = ((65535 == 1) || false)
                                                    _ = v48
                                                } else {
                                                    var v49 float32 = v4
                                                    _ = v49
                                                }
                                                var v50 int32 = 6
                                                _ = v50
                                            case 949:
                                                if ((v24 >> (((^(v3))) & 63)) >= 0) {
                                                    var v51 float32 = (v4 * (((v4 * 1e10) * v4) + (1e10 / 1.0)))
                                                    _ = v51
                                                }
                                            default:
                                                if (!false) {
                                                    switch v6 {
                                                    case "":
                                                        fmt.Println(v5, fi21)
                                                    case "foo bar":
                                                        var v52 float32 = (0.0 / ((float32(((1.5 * (0.0 + float64(0.0))) - 0.1)) - (v4 * (((3.14159 - 1.5) + 0.0) + v4))) / 1e10))
                                                        _ = v52
                                                    }
                                                }
                                            }
                                            v7 = append(v7, 42)
                                            _ = len(v7)
                                            if false {
                                                if true {
                                                    switch (fi21 & 1) {
                                                    case -1:
                                                        _ = len(v7)
                                                    case 207:
                                                        if (1 != 2147483647) {
                                                            var v53 bool = (2 > 255)
                                                            _ = v53
                                                        } else {
                                                            v54 := GenS3{F0: 840, F1: -2147483648}
                                                            v7 = append(v7, 32768)
                                                            _ = v54
                                                        }
                                                    case -9223372036854775808:
                                                        fmt.Println(v4)
                                                    }
                                                    fmt.Println(fi21, v1, v5)
                                                }
                                                if len(v7) > 0 {
                                                    v7[0] = 347
                                                }
                                                if true {
                                                    v7 = append(v7, (fi21 & (fi21 << (((v3 - v3)) & 63))))
                                                    var v55 bool = (((((v4 * 1e10) <= 1.5) || (255 < (v3 >> ((1) & 63)))) && false) || false)
                                                    var v56 bool = ((v1 << (((^(v3))) & 63)) > (v1 - 782))
                                                    _ = v55
                                                    _ = v56
                                                } else {
                                                    v5 = 0
                                                    var v57 int32 = 0
                                                    _ = v57
                                                }
                                            } else {
                                                var v58 float64 = 0.1
                                                v7 = append(v7, -2147483648)
                                                switch 32767 {
                                                case -1:
                                                    fmt.Println(v24)
                                                }
                                                _ = v58
                                            }
                                        }
                                    }
                                    var v59 bool = false
                                    _ = v59
                                }
                                _ = v24
                            }
                        }
                        v60 := GenS2{F0: 4294967295}
                        _ = v23
                        _ = v60
                    }
                    switch (^(fi21)) {
                    case 679:
                        v5 += 2
                    case 1:
                        if false {
                            v7 = append(v7, v1)
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        } else {
                            v61 := GenS8{F0: v3}
                            _ = v61
                        }
                    default:
                        fmt.Println(v5, v2)
                        _ = len(v7)
                        var v62 uint32 = 0
                        _ = v62
                    }
                    _ = v22
                } else {
                    if true {
                        switch 32767 {
                        case 696:
                            var v63 int64 = (v2 << ((1) & 63))
                            _ = v63
                        case 4294967295:
                            if false {
                                v64 := GenS5{F0: 2147483647, F1: 2}
                                switch -1 {
                                case 9223372036854775807:
                                    var v65 uint64 = 670
                                    _ = v65
                                case 0:
                                    var v66 string = ""
                                    _ = v66
                                case -2147483648:
                                    switch fi20 {
                                    case 32768:
                                        var v67 uint8 = (v5 & (v5 << ((v3) & 63)))
                                        _ = v67
                                    case -717:
                                        var v68 int16 = 42
                                        _ = v68
                                    case -9223372036854775808:
                                        var v69 uint = v3
                                        v70 := GenS5{F0: 254, F1: 2}
                                        _ = v69
                                        _ = v70
                                    default:
                                        v1 = 766
                                    }
                                }
                                v5 = (^(v5))
                                _ = v64
                            } else {
                                var v71 uint16 = 96
                                _ = v71
                            }
                        }
                        if ((42 <= uint32(v3)) && false) {
                            _ = len(v7)
                        } else {
                            var v72 float64 = (float64((v4 * v4)) + 0.1)
                            v73 := GenS3{F0: 0, F1: -2147483648}
                            if (-131 == 1000) {
                                var v74 uint32 = 2
                                var v75 uint16 = 2
                                v6 = v6
                                _ = v74
                                _ = v75
                            }
                            switch (fi20 << (((v3 - v3)) & 63)) {
                            case -210:
                                if false {
                                    v72 = 0.1
                                    var v76 float32 = 0.1
                                    _ = v76
                                }
                            case 32767:
                                switch 128 {
                                case 696:
                                    var v77 int8 = 115
                                    _ = v77
                                case 2147483647:
                                    if true {
                                        fmt.Println(v72, fi21, v5)
                                        fmt.Println(v4, v1, v6)
                                        switch v6 {
                                        case "a":
                                            if false {
                                                if len(v7) >= 2 {
                                                    v7 = v7[1:len(v7)]
                                                }
                                                var v78 string = (v6 + "\n\t")
                                                var v79 int8 = 127
                                                _ = v78
                                                _ = v79
                                            }
                                        case "hello":
                                            if (!true) {
                                                var v80 uint = v3
                                                _ = v80
                                            } else {
                                                switch 128 {
                                                case 32767:
                                                    v6 = ("a" + v6)
                                                case 9223372036854775807:
                                                    fmt.Println(v1, fi20, v72)
                                                }
                                                v81 := GenS9{F0: (!true), F1: -66}
                                                _ = v81
                                            }
                                        case "foo bar":
                                            fmt.Println(fi21, v5)
                                        }
                                    } else {
                                        v6 = (v6 + v6)
                                        var v82 int = (^(v1))
                                        if (65535 <= 32767) {
                                            var v83 bool = (false && true)
                                            if (0 > 18446744073709551615) {
                                                fmt.Println(v4)
                                                if v83 {
                                                    var v84 int64 = (v2 << (((v3 * v3)) & 63))
                                                    v3 = 18446744073709551615
                                                    var v85 uint8 = (v5 | v5)
                                                    _ = v84
                                                    _ = v85
                                                }
                                                if (v83 && (514 < 42)) {
                                                    var v86 float64 = v72
                                                    _ = v86
                                                }
                                                v2 -= -360
                                                v83 = v83
                                            } else {
                                                var v87 uint16 = 0
                                                fmt.Println(v1, v83, v82)
                                                _ = v87
                                            }
                                            _ = v83
                                        } else {
                                            var v88 uint16 = 66
                                            if true {
                                                v89 := GenS4{F0: (v3 - (v3 & (v3 + (v3 + (^(v3)))))), F1: 423}
                                                _ = v89
                                            }
                                            if (v4 < v4) {
                                                if true {
                                                    var v90 int8 = 0
                                                    _ = v90
                                                }
                                            } else {
                                                v1 -= v1
                                                var v91 uint32 = 2
                                                v7 = append(v7, -2147483648)
                                                _ = v91
                                            }
                                            _ = v88
                                        }
                                        _ = v82
                                    }
                                case 9223372036854775807:
                                    v6 = "a"
                                }
                            }
                            _ = v72
                            _ = v73
                        }
                        v92 := GenS8{F0: 2}
                        _ = v92
                    } else {
                        var v93 float32 = (v4 / ((0.0 * v4) * ((v4 - float32((1.5 / -2.25))) / v4)))
                        var v94 bool = false
                        _ = v93
                        _ = v94
                    }
                    v3 = 42
                }
            }
            v95 := GenS5{F0: 64, F1: 65535}
            if false {
                v96 := GenS7{F0: 274, F1: v6}
                var v97 int64 = v2
                v97 -= v97
                if (-1 <= 67) {
                    v1 += (v1 + fi21)
                    if (263 != (v3 << (((v3 << (((v3 - v3)) & 63))) & 63))) {
                        switch v1 {
                        case 2147483647:
                            fmt.Println(fi21, v97, v1)
                        case -1:
                            if (557 < (v3 >> ((v3) & 63))) {
                                var v98 int = fi21
                                v2 += (^(v97))
                                if true {
                                    if false {
                                        v7 = append(v7, -41)
                                        var v99 bool = (!((((fi21 - (v98 / (((fi20 / ((int(v2)) | 1))) | 1))) != (v1 << ((v3) & 63))) && true) && (v6 > (v6 + v6))))
                                        if false {
                                            switch -431 {
                                            case 32767:
                                                fmt.Println(v3)
                                            default:
                                                if false {
                                                    v6 = (v6 + ((v6 + (v6 + v6)) + v6))
                                                    if len(v7) > 0 {
                                                        v7[0] = v1
                                                    }
                                                    if (-268 <= -560) {
                                                        if v99 {
                                                            var v100 float64 = 0.1
                                                            _ = v100
                                                        }
                                                        if (!v99) {
                                                            switch -746 {
                                                            case 2147483647:
                                                                if ((v2 / ((v2) | 1)) <= -925) {
                                                                    var v101 uint16 = 537
                                                                    _ = v101
                                                                } else {
                                                                    var v102 bool = v99
                                                                    fmt.Println(fi20)
                                                                    _ = v102
                                                                }
                                                            case 0:
                                                                v103 := GenS10{F0: 65535, F1: v5, F2: 0}
                                                                _ = v103
                                                            }
                                                            v104 := GenS3{F0: -1, F1: int32(v3)}
                                                            _ = len(v7)
                                                            _ = v104
                                                        } else {
                                                            var v105 uint8 = v5
                                                            if len(v7) >= 2 {
                                                                v7 = v7[1:len(v7)]
                                                            }
                                                            _ = v105
                                                        }
                                                        v7 = append(v7, v1)
                                                    }
                                                }
                                            }
                                            v106 := GenS11{F0: 100}
                                            if len(v7) > 0 {
                                                v7[0] = (fi21 >> (((v3 % (((v3 - 179)) | 1))) & 63))
                                            }
                                            _ = v106
                                        } else {
                                            switch "a" {
                                            case "hello":
                                                fmt.Println(v4, v5, fi20)
                                            }
                                        }
                                        _ = v99
                                    }
                                } else {
                                    var v107 int8 = -110
                                    switch v6 {
                                    case "\n\t":
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                        switch v6 {
                                        case "foo bar":
                                            var v108 uint8 = v5
                                            _ = v108
                                        case "hello":
                                            v107 += -28
                                        default:
                                            var v109 int8 = v107
                                            _ = v109
                                        }
                                        v3 += (v3 / (((v3 | v3)) | 1))
                                    case "hello":
                                        if ((v3 / ((2) | 1)) >= (v3 >> ((v3) & 63))) {
                                            fmt.Println(v5, v5)
                                            if true {
                                                if len(v7) > 0 {
                                                    v7[0] = v98
                                                }
                                                var v110 uint8 = 159
                                                _ = v110
                                            }
                                        }
                                    default:
                                        v111 := GenS9{F0: (0 >= 1), F1: (v107 >> ((v3) & 63))}
                                        _ = v111
                                    }
                                    var v112 bool = true
                                    _ = v107
                                    _ = v112
                                }
                                _ = v98
                            }
                        default:
                            fmt.Println(fi20, v97, fi20)
                        }
                    }
                } else {
                    v113 := GenS2{F0: 74}
                    if len(v7) > 0 {
                        v7[0] = (fi21 & (^(v1)))
                    }
                    _ = v113
                }
                var v114 uint32 = 881
                _ = v114
                _ = v96
                _ = v97
            } else {
                v115 := GenS11{F0: 127}
                _ = v115
            }
            v116 := GenS4{F0: (v3 - v3), F1: 32767}
            v117 := GenS6{F0: 42, F1: v3}
            _ = fi21
            _ = v116
            _ = v117
            _ = v95
        }
        _ = fi20
    }
    v118 := GenS8{F0: 42}
    if (true && false) {
        v2 = -1
        switch v1 {
        case -92:
            for fi119 := 0; fi119 < 5; fi119++ {
                if false {
                    if (true || true) {
                        var v120 uint8 = v5
                        for fi121 := 0; fi121 < 5; fi121++ {
                            if (!((false && false) && true)) {
                                switch v1 {
                                case -2147483648:
                                    v2 = v2
                                case 894:
                                    v2 -= -637
                                default:
                                    v122 := GenS8{F0: (v3 & v3)}
                                    _ = v122
                                }
                                switch "foo bar" {
                                case "foo bar":
                                    v123 := GenS12{F0: 90, F1: v3, F2: (fi121 >> ((v3) & 63))}
                                    _ = v123
                                case "\n\t":
                                    if false {
                                        v7 = append(v7, v1)
                                        v124 := GenS7{F0: 0, F1: v6}
                                        _ = v124
                                    }
                                case "hello":
                                    switch (v6 + v6) {
                                    case "\n\t":
                                        v2 += (v2 | v2)
                                    case "":
                                        v125 := GenS13{F0: (v120 & v5), F1: 4294967295, F2: 0}
                                        _ = v125
                                    case "hello":
                                        var v126 float64 = -2.25
                                        v127 := GenS1{F0: (^(v2)), F1: 0, F2: false}
                                        _ = v126
                                        _ = v127
                                    default:
                                        v5 = (v5 * 0)
                                    }
                                default:
                                    _ = len(v7)
                                }
                                _ = len(v7)
                            } else {
                                switch 4294967295 {
                                case 1:
                                    v4 = 1.5
                                default:
                                    v4 = (v4 + 0.0)
                                    v128 := GenS13{F0: (v120 >> ((uint(v5)) & 63)), F1: 282, F2: 2}
                                    v129 := GenS14{F0: (((v6 + ("foo bar" + "a")) + ((v6 + ("" + v6)) + v6)) + ("foo bar" + "hello")), F1: (v2 ^ (v2 / ((v2) | 1))), F2: int(v2)}
                                    var v130 int16 = -32768
                                    _ = v128
                                    _ = v129
                                    _ = v130
                                }
                                v6 = (v6 + v6)
                            }
                            if ((v2 + v2) > v2) {
                                var v131 int16 = -1
                                switch (^(fi119)) {
                                case 9223372036854775807:
                                    var v132 bool = false
                                    _ = v132
                                case 2147483647:
                                    if len(v7) > 0 {
                                        v7[0] = -1
                                    }
                                }
                                var v133 int16 = v131
                                v1 -= (v1 * v1)
                                if true {
                                    v7 = append(v7, (fi119 ^ 127))
                                } else {
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                }
                                fmt.Println(v6, v133, v3)
                                _ = v131
                                _ = v133
                            } else {
                                if len(v7) > 0 {
                                    v7[0] = 128
                                }
                                switch (v6 + v6) {
                                case "foo bar":
                                    switch 9223372036854775807 {
                                    case 2147483647:
                                        var v134 int8 = 127
                                        _ = v134
                                    }
                                case "":
                                    if (v2 < 9223372036854775807) {
                                        v120 -= (^(v120))
                                    }
                                    switch (v1 - (^(fi119))) {
                                    case 32768:
                                        var v135 uint32 = 550
                                        _ = v135
                                    case 2147483647:
                                        v5 += (v120 | (^(v5)))
                                    default:
                                        fmt.Println(v2, fi119)
                                    }
                                    if (!(true && (v4 < v4))) {
                                        v2 -= v2
                                    } else {
                                        v1 += 127
                                    }
                                default:
                                    _ = len(v7)
                                }
                                var v136 uint8 = uint8(v3)
                                _ = v136
                            }
                            _ = fi121
                        }
                        v3 = 42
                        _ = v120
                    } else {
                        var v137 uint16 = 300
                        _ = v137
                    }
                    var v138 uint64 = 4294967295
                    v139 := GenS5{F0: 465, F1: 0}
                    _ = v138
                    _ = v139
                } else {
                    v140 := GenS15{F0: v3, F1: 4, F2: 401}
                    var v141 uint64 = 42
                    if len(v7) > 0 {
                        v7[0] = v1
                    }
                    _ = v140
                    _ = v141
                }
                v7 = append(v7, (v1 + (fi119 * (v1 << ((1) & 63)))))
                _ = fi119
            }
        case -335:
            v3 += v3
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
