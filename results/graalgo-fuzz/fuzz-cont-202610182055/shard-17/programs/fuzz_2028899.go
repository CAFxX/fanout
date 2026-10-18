package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 string; F1 uint; F2 int32 }
type GenS4 struct { F0 string; F1 uint64; F2 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8; F1 uint16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16; F1 uint32 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int8; F1 uint; F2 bool }
type GenS13 struct { F0 int32; F1 int8; F2 int64 }
type GenS14 struct { F0 int32; F1 int64 }
type GenS15 struct { F0 uint32; F1 uint64 }
type GenS16 struct { F0 int8; F1 uint }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 839
    var v3 int64 = 1
    var v4 int64 = (v3 + (v3 >> ((493) & 63)))
    fmt.Println(v4, v4, v2)
    if false {
        delete(v1, 42)
        if true {
            v5 := []int{}
            v6 := GenS1{F0: v2, F1: v4, F2: 392}
            var v7 int32 = -86
            _ = v5
            _ = v6
            _ = v7
        }
    }
    fmt.Println(v4, v4)
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := GenS2{F0: (fi8 << ((342) & 63)), F1: 2}
        _ = fi8
        _ = v9
    }
    var v10 uint8 = 0
    var v11 bool = false
    v12 := GenS2{F0: 127, F1: 578}
    v13 := GenS3{F0: (("foo bar" + "hello") + "foo bar"), F1: 4294967295, F2: 1}
    var v14 int16 = 0
    fmt.Println(v11, v3, v11)
    fmt.Println(v11, v2)
    v11 = v11
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 7; fi16++ {
            if true {
                v14 += (v14 & v14)
                v4 -= -477
            } else {
                switch (("a" + "\n\t") + "") {
                case "\n\t":
                    var v17 uint16 = v2
                    _ = v17
                default:
                    var v18 float32 = ((0.1 / ((float32((1e10 - 3.14159)) * (1e10 + -2.25)) / ((0.1 / 1.0) + ((1.5 * 0.1) + 0.0)))) / 1.0)
                    _ = v18
                }
            }
            _ = fi16
        }
        _ = fi15
    }
    v4 += (v3 % (((v3 + v3)) | 1))
    if (4294967295 <= 58) {
        var v19 uint16 = v2
        _ = v19
    }
    if (-2.25 > 1e10) {
        if ("" != ("hello" + "a")) {
            for fi20 := 0; fi20 < 19; fi20++ {
                var v21 string = (("a" + "foo bar") + "a")
                for fi22 := 0; fi22 < 1; fi22++ {
                    var v23 int16 = -793
                    fmt.Println(v23, fi22)
                    v24 := GenS4{F0: v21, F1: 753, F2: -44}
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                if (v11 || (fi20 > fi20)) {
                    if (v11 && (v11 && (((-1 != 127) && true) || (!(v11 && (v14 < v14)))))) {
                        if v11 {
                            v2 = (v2 << ((9) & 63))
                        }
                        if (v11 || (4294967295 >= 65535)) {
                            v4 = (v3 ^ (v3 + (v4 | v4)))
                            switch (fi20 << ((1) & 63)) {
                            case -709:
                                for fi25 := 0; fi25 < 2; fi25++ {
                                    fmt.Println(v10)
                                    var v26 uint8 = (v10 >> ((42) & 63))
                                    var v27 int16 = v14
                                    switch 32767 {
                                    case 766:
                                        switch (("\n\t" + v21) + "foo bar") {
                                        case "hello":
                                            fmt.Println(fi25, v3)
                                        }
                                    case 42:
                                        if (!false) {
                                            var v28 bool = v11
                                            v11 = false
                                            var v29 int = fi20
                                            v26 = (v10 << ((65535) & 63))
                                            _ = v28
                                            _ = v29
                                        } else {
                                            var v30 int8 = 100
                                            if v11 {
                                                switch "a" {
                                                case "\n\t":
                                                    v3 = v4
                                                default:
                                                    if false {
                                                        v11 = (v27 <= 42)
                                                        v31 := []int{fi20, (fi20 % (((fi20 >> ((42) & 63))) | 1))}
                                                        _ = v31
                                                    } else {
                                                        var v32 string = v21
                                                        fmt.Println(v27, v21)
                                                        _ = v32
                                                    }
                                                }
                                                delete(v1, 2147483648)
                                                fmt.Println(v26, v2)
                                                if (!(int8(v14) < (v30 << ((841) & 63)))) {
                                                    if (fi25 < (fi25 % ((fi25) | 1))) {
                                                        if v11 {
                                                            v21 = v21
                                                            v1[2147483648] = true
                                                        } else {
                                                            v33 := v1[2147483648]
                                                            if ((^(v14)) > (v14 - v27)) {
                                                                if v33 {
                                                                    var v34 float64 = 1.5
                                                                    delete(v1, 975)
                                                                    switch (fi20 / (((fi20 | (fi20 | fi20))) | 1)) {
                                                                    case 4294967295:
                                                                        v14 = 1
                                                                    case 32767:
                                                                        if true {
                                                                            v35 := GenS2{F0: -9223372036854775808, F1: 4294967295}
                                                                            v27 = 293
                                                                            v11 = v11
                                                                            _ = v35
                                                                        }
                                                                    case 323:
                                                                        switch (v21 + (((v21 + ((v21 + ((("a" + "a") + v21) + ((((v21 + "hello") + v21) + v21) + v21))) + (v21 + (v21 + v21)))) + v21) + "hello")) {
                                                                        case "a":
                                                                            var v36 int32 = 2147483647
                                                                            _ = v36
                                                                        }
                                                                    }
                                                                    _ = v34
                                                                } else {
                                                                    switch (^(fi25)) {
                                                                    case -9223372036854775808:
                                                                        var v37 uint32 = 0
                                                                        _ = v37
                                                                    }
                                                                    var v38 uint64 = 42
                                                                    v2 += (v2 * v2)
                                                                    _ = v38
                                                                }
                                                            } else {
                                                                v39 := GenS2{F0: 9223372036854775807, F1: 0}
                                                                if v11 {
                                                                    switch v21 {
                                                                    case "hello":
                                                                        delete(v1, 65535)
                                                                    case "foo bar":
                                                                        switch (fi25 >> ((1) & 63)) {
                                                                        case 0:
                                                                            if v11 {
                                                                                if ((v2 | 126) < 805) {
                                                                                    var v40 uint32 = 0
                                                                                    var v41 int = (fi25 - 128)
                                                                                    v41 = fi20
                                                                                    _ = v40
                                                                                    _ = v41
                                                                                } else {
                                                                                    var v42 uint8 = v26
                                                                                    _ = v42
                                                                                }
                                                                                v43 := []int{fi20, (fi25 - (-(fi20))), (^(fi20))}
                                                                                var v44 int8 = 31
                                                                                _ = v43
                                                                                _ = v44
                                                                            } else {
                                                                                v45 := []int{(fi20 / (((fi25 ^ 128)) | 1))}
                                                                                fmt.Println(v2, v10)
                                                                                _ = v45
                                                                            }
                                                                        case 9223372036854775807:
                                                                            if v33 {
                                                                                if (1e10 == 3.14159) {
                                                                                    var v46 float32 = 3.14159
                                                                                    v47 := GenS5{F0: 42}
                                                                                    _ = v46
                                                                                    _ = v47
                                                                                }
                                                                            } else {
                                                                                v10 -= 8
                                                                                if ((((0.1 - (1.5 * 1e10)) != 3.14159) || ((v4 ^ (v3 * (v4 / ((v4) | 1)))) != (-(v4)))) || false) {
                                                                                    fmt.Println(v21, v33)
                                                                                    fmt.Println(v3, v14)
                                                                                } else {
                                                                                    v21 = "a"
                                                                                    v33 = (-2147483648 < -48)
                                                                                }
                                                                            }
                                                                        default:
                                                                            v48 := []int{}
                                                                            _ = v48
                                                                        }
                                                                    }
                                                                    switch ("a" + v21) {
                                                                    case "foo bar":
                                                                        v49 := GenS5{F0: 697}
                                                                        fmt.Println(v2, v11, v33)
                                                                        if (v30 > 42) {
                                                                            v4 += v3
                                                                            if v11 {
                                                                                if true {
                                                                                    var v50 int64 = 9223372036854775807
                                                                                    _ = v50
                                                                                }
                                                                                var v51 string = (v21 + (((v21 + v21) + v21) + (v21 + "")))
                                                                                fmt.Println(v26)
                                                                                _ = v51
                                                                            }
                                                                            var v52 int16 = -1
                                                                            _ = v52
                                                                        } else {
                                                                            var v53 int32 = 2147483647
                                                                            v26 -= v26
                                                                            _ = v53
                                                                        }
                                                                        _ = v49
                                                                    case "a":
                                                                        var v54 float32 = 1e10
                                                                        _ = v54
                                                                    case "hello":
                                                                        var v55 uint64 = 807
                                                                        _ = v55
                                                                    default:
                                                                        fmt.Println(v21, fi25)
                                                                    }
                                                                    if v33 {
                                                                        v56 := []int{}
                                                                        _ = v56
                                                                    } else {
                                                                        var v57 uint16 = v2
                                                                        _ = v57
                                                                    }
                                                                }
                                                                _ = v39
                                                            }
                                                            v4 = v3
                                                            fmt.Println(v27, v26)
                                                            if true {
                                                                if false {
                                                                    switch fi20 {
                                                                    case -9223372036854775808:
                                                                        v58 := GenS2{F0: (fi25 * fi20), F1: 118}
                                                                        _ = v58
                                                                    case 66:
                                                                        if v33 {
                                                                            if true {
                                                                                switch "hello" {
                                                                                case "\n\t":
                                                                                    var v59 bool = v11
                                                                                    _ = v59
                                                                                case "a":
                                                                                    v60 := GenS6{F0: (v26 << ((42) & 63)), F1: (v2 | (v2 + 779))}
                                                                                    _ = v60
                                                                                case "":
                                                                                    v26 = 1
                                                                                default:
                                                                                    var v61 float32 = 0.1
                                                                                    _ = v61
                                                                                }
                                                                                var v62 uint32 = 450
                                                                                _ = v62
                                                                            }
                                                                            var v63 int32 = -1
                                                                            _ = v63
                                                                        }
                                                                    default:
                                                                        if v11 {
                                                                            var v64 int64 = 9223372036854775807
                                                                            var v65 int64 = (v3 << ((65535) & 63))
                                                                            _ = v64
                                                                            _ = v65
                                                                        }
                                                                    }
                                                                    var v66 int32 = 0
                                                                    _ = v66
                                                                } else {
                                                                    var v67 int16 = (v27 + (v27 ^ -752))
                                                                    v11 = v33
                                                                    _ = v67
                                                                }
                                                                v11 = (v30 < v30)
                                                                var v68 uint = 813
                                                                _ = v68
                                                            } else {
                                                                v69 := GenS5{F0: 65535}
                                                                _ = v69
                                                            }
                                                            if (1.5 == 1e10) {
                                                                v70 := GenS3{F0: ((v21 + ((v21 + "") + v21)) + v21), F1: 359, F2: -1}
                                                                v71 := GenS1{F0: v2, F1: (-(v4)), F2: 1}
                                                                var v72 int64 = v3
                                                                _ = v70
                                                                _ = v71
                                                                _ = v72
                                                            }
                                                            switch (fi20 & fi20) {
                                                            case 32768:
                                                                v73 := GenS1{F0: v2, F1: (v3 - (-(v4))), F2: 2}
                                                                _ = v73
                                                            case 4294967295:
                                                                var v74 uint64 = 437
                                                                _ = v74
                                                            case 276:
                                                                v75 := GenS1{F0: v2, F1: -1, F2: 957}
                                                                if (18446744073709551615 >= 422) {
                                                                    var v76 bool = ((1e10 * 1.5) <= ((0.0 - (((0.0 * 1.5) * 3.14159) / 1e10)) * 1.5))
                                                                    _ = v76
                                                                } else {
                                                                    if (true || v33) {
                                                                        v1[617] = false
                                                                        fmt.Println(fi25)
                                                                        if v11 {
                                                                            v10 -= 127
                                                                            var v77 uint64 = 2
                                                                            v11 = ((-246 < 1) && (!v33))
                                                                            var v78 float32 = (float32((0.0 * 1.5)) * ((float32((((((-2.25 * 3.14159) - ((0.0 / 1.5) + 1.5)) * 0.0) + 1.5) / 1e10)) * -2.25) * (1e10 * 3.14159)))
                                                                            _ = v77
                                                                            _ = v78
                                                                        }
                                                                    }
                                                                    var v79 int64 = (v3 % ((-1) | 1))
                                                                    _ = v79
                                                                }
                                                                v3 += v3
                                                                _ = v75
                                                            }
                                                            _ = v33
                                                        }
                                                        v80 := GenS7{F0: -198}
                                                        var v81 int64 = v3
                                                        _ = v80
                                                        _ = v81
                                                    } else {
                                                        switch 2147483647 {
                                                        case 42:
                                                            v2 += (^(v2))
                                                        case 221:
                                                            var v82 uint64 = 2
                                                            _ = v82
                                                        case 942:
                                                            var v83 uint16 = (v2 - v2)
                                                            v84 := []int{}
                                                            _ = len(v84)
                                                            switch "hello" {
                                                            case "":
                                                                if true {
                                                                    var v85 bool = (!(1 < -1))
                                                                    _ = len(v1)
                                                                    _ = v85
                                                                }
                                                            default:
                                                                if v11 {
                                                                    switch (v21 + (v21 + v21)) {
                                                                    case "foo bar":
                                                                        var v86 int8 = 1
                                                                        _ = v86
                                                                    case "":
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            var v87 uint = 29
                                                                            _ = v87
                                                                        case "":
                                                                            switch (-(fi25)) {
                                                                            case -2147483648:
                                                                                if true {
                                                                                    switch v21 {
                                                                                    case "":
                                                                                        var v88 uint32 = 2147483648
                                                                                        _ = v88
                                                                                    case "\n\t":
                                                                                        if (0.0 <= 1.5) {
                                                                                            fmt.Println(v30, v11, v26)
                                                                                        } else {
                                                                                            switch v21 {
                                                                                            case "":
                                                                                                if ((0.1 == 0.0) && (v30 == v30)) {
                                                                                                    v89 := GenS8{F0: v27, F1: 117}
                                                                                                    _ = v89
                                                                                                }
                                                                                            case "\n\t":
                                                                                                v90 := GenS5{F0: 0}
                                                                                                _ = v90
                                                                                            case "foo bar":
                                                                                                switch "foo bar" {
                                                                                                case "":
                                                                                                    var v91 uint8 = 255
                                                                                                    _ = v91
                                                                                                case "\n\t":
                                                                                                    var v92 uint16 = 2
                                                                                                    _ = v92
                                                                                                }
                                                                                            default:
                                                                                                v30 = -1
                                                                                            }
                                                                                            delete(v1, 2)
                                                                                        }
                                                                                    case "foo bar":
                                                                                        fmt.Println(v4, fi20, v26)
                                                                                    default:
                                                                                        var v93 float64 = 1e10
                                                                                        _ = v93
                                                                                    }
                                                                                    v14 -= v27
                                                                                    var v94 float32 = (0.1 / 1.0)
                                                                                    _ = v94
                                                                                } else {
                                                                                    v95 := GenS5{F0: 701}
                                                                                    v84 = append(v84, fi25)
                                                                                    _ = v95
                                                                                }
                                                                            }
                                                                        }
                                                                    default:
                                                                        v30 = (^(v30))
                                                                    }
                                                                    v96 := GenS9{F0: 0}
                                                                    _ = v96
                                                                } else {
                                                                    if ((v4 | 38) > (^(v3))) {
                                                                        v11 = (!(!(!(v11 && v11))))
                                                                    }
                                                                    var v97 string = v21
                                                                    if v11 {
                                                                        fmt.Println(v11, v14, v2)
                                                                        var v98 int = fi20
                                                                        _ = v98
                                                                    } else {
                                                                        if (42 != 0) {
                                                                            fmt.Println(v97, v26)
                                                                            if v11 {
                                                                                switch (((("foo bar" + "a") + "") + "a") + "a") {
                                                                                case "hello":
                                                                                    if false {
                                                                                        var v99 float32 = 0.1
                                                                                        v1[2] = true
                                                                                        switch fi25 {
                                                                                        case 570:
                                                                                            fmt.Println(v14, v99, v27)
                                                                                        default:
                                                                                            if len(v84) > 0 {
                                                                                                v84[0] = (fi20 + 501)
                                                                                            }
                                                                                        }
                                                                                        if v11 {
                                                                                            var v100 uint64 = 2
                                                                                            var v101 string = v97
                                                                                            v27 = (v27 & v14)
                                                                                            _ = len(v1)
                                                                                            v101 = v21
                                                                                            _ = v100
                                                                                            _ = v101
                                                                                        } else {
                                                                                            var v102 uint16 = 0
                                                                                            var v103 string = "a"
                                                                                            _ = v102
                                                                                            _ = v103
                                                                                        }
                                                                                        _ = len(v84)
                                                                                        fmt.Println(v2)
                                                                                        v30 += v30
                                                                                        v104 := GenS10{F0: v83}
                                                                                        _ = v104
                                                                                        _ = v99
                                                                                    } else {
                                                                                        switch v97 {
                                                                                        case "hello":
                                                                                            v105 := GenS11{F0: uint64(v83)}
                                                                                            _ = v105
                                                                                        }
                                                                                        _ = len(v84)
                                                                                        if v11 {
                                                                                            var v106 int32 = 399
                                                                                            var v107 float32 = 1.5
                                                                                            v108 := GenS8{F0: v27, F1: 944}
                                                                                            _ = v106
                                                                                            _ = v107
                                                                                            _ = v108
                                                                                        }
                                                                                        var v109 uint16 = v83
                                                                                        v110 := GenS12{F0: -103, F1: 255, F2: v11}
                                                                                        _ = v109
                                                                                        _ = v110
                                                                                    }
                                                                                case "":
                                                                                    if true {
                                                                                        var v111 float32 = 1.5
                                                                                        v3 = (v4 / (((v4 >> ((0) & 63))) | 1))
                                                                                        _ = v111
                                                                                    }
                                                                                case "\n\t":
                                                                                    v112 := GenS4{F0: "a", F1: 1, F2: v30}
                                                                                    _ = v112
                                                                                default:
                                                                                    fmt.Println(v10)
                                                                                }
                                                                                var v113 float32 = 3.14159
                                                                                v114 := GenS13{F0: 2147483647, F1: v30, F2: v4}
                                                                                _ = v113
                                                                                _ = v114
                                                                            }
                                                                            if v11 {
                                                                                if v11 {
                                                                                    v115 := GenS14{F0: 963, F1: (^(v4))}
                                                                                    _ = v115
                                                                                } else {
                                                                                    if (2 < 643) {
                                                                                        switch (fi25 % (((fi20 + fi25)) | 1)) {
                                                                                        case 2147483647:
                                                                                            v83 -= (v2 & v2)
                                                                                        default:
                                                                                            fmt.Println(fi20, v10, v2)
                                                                                        }
                                                                                        var v116 uint32 = 2147483647
                                                                                        switch v97 {
                                                                                        case "a":
                                                                                            v11 = (v27 != (^(v27)))
                                                                                        case "hello":
                                                                                            v26 = v26
                                                                                        case "foo bar":
                                                                                            var v117 bool = true
                                                                                            _ = v117
                                                                                        default:
                                                                                            if v11 {
                                                                                                v21 = v21
                                                                                                v27 += 42
                                                                                            } else {
                                                                                                fmt.Println(v3, fi25)
                                                                                            }
                                                                                        }
                                                                                        _ = v116
                                                                                    } else {
                                                                                        fmt.Println(v14)
                                                                                        if v11 {
                                                                                            delete(v1, 44)
                                                                                            var v118 int8 = v30
                                                                                            fmt.Println(v30, v4)
                                                                                            _ = v118
                                                                                        }
                                                                                    }
                                                                                    _ = len(v84)
                                                                                    delete(v1, 42)
                                                                                }
                                                                                v83 = v2
                                                                            }
                                                                        }
                                                                        fmt.Println(v83, v4, v97)
                                                                    }
                                                                    if v11 {
                                                                        _ = len(v1)
                                                                    } else {
                                                                        v119 := GenS15{F0: 1, F1: 1}
                                                                        if v11 {
                                                                            delete(v1, 1)
                                                                            v120 := GenS11{F0: 377}
                                                                            _ = v120
                                                                        } else {
                                                                            if (!(((v26 | v26) <= 2) || v11)) {
                                                                                var v121 int = fi25
                                                                                fmt.Println(v26)
                                                                                v122 := GenS13{F0: -1, F1: v30, F2: v4}
                                                                                _ = v121
                                                                                _ = v122
                                                                            } else {
                                                                                var v123 uint8 = 42
                                                                                var v124 int = (fi20 / ((563) | 1))
                                                                                var v125 uint64 = 875
                                                                                _ = v123
                                                                                _ = v124
                                                                                _ = v125
                                                                            }
                                                                            var v126 int64 = v3
                                                                            if (v11 && false) {
                                                                                if v11 {
                                                                                    _ = len(v1)
                                                                                    var v127 int8 = (v30 >> ((0) & 63))
                                                                                    v128 := GenS6{F0: (v10 | v10), F1: v83}
                                                                                    _ = v127
                                                                                    _ = v128
                                                                                } else {
                                                                                    v14 = -32768
                                                                                    v30 += (v30 & v30)
                                                                                    if (-1 != 42) {
                                                                                        v129 := GenS5{F0: 699}
                                                                                        v11 = ((-2.25 / (1e10 / -2.25)) == float32(float64(1.5)))
                                                                                        _ = v129
                                                                                    } else {
                                                                                        v130 := GenS16{F0: -74, F1: 42}
                                                                                        var v131 uint32 = 2147483647
                                                                                        switch ((v21 + "hello") + "foo bar") {
                                                                                        case "\n\t":
                                                                                            var v132 float64 = 3.14159
                                                                                            _ = v132
                                                                                        case "":
                                                                                            v133 := GenS5{F0: 2}
                                                                                            _ = v133
                                                                                        }
                                                                                        fmt.Println(v21)
                                                                                        var v134 uint8 = (v10 + (^(v10)))
                                                                                        v135 := GenS4{F0: v21, F1: 2, F2: v30}
                                                                                        _ = v130
                                                                                        _ = v131
                                                                                        _ = v134
                                                                                        _ = v135
                                                                                    }
                                                                                }
                                                                                var v136 string = v97
                                                                                v26 += (v10 - v10)
                                                                                fmt.Println(v4)
                                                                                fmt.Println(v4)
                                                                                if (v3 != (v126 >> ((18446744073709551615) & 63))) {
                                                                                    if true {
                                                                                        v1[uint32(v26)] = (2 == 2)
                                                                                        var v137 uint64 = 0
                                                                                        v138 := GenS5{F0: 2}
                                                                                        _ = v137
                                                                                        _ = v138
                                                                                    } else {
                                                                                        v139 := GenS1{F0: (v83 % ((v2) | 1)), F1: (v126 >> ((2) & 63)), F2: 0}
                                                                                        _ = len(v84)
                                                                                        fmt.Println(v10)
                                                                                        _ = v139
                                                                                    }
                                                                                    fmt.Println(v136, fi20, v97)
                                                                                    if v11 {
                                                                                        v10 += 42
                                                                                    }
                                                                                    if v11 {
                                                                                        var v140 int = (fi25 & (fi25 >> ((12) & 63)))
                                                                                        _ = v140
                                                                                    }
                                                                                    var v141 uint8 = (v26 - (v26 - v10))
                                                                                    _ = v141
                                                                                } else {
                                                                                    v2 = 1
                                                                                }
                                                                                var v142 int16 = (v14 + 0)
                                                                                _ = v136
                                                                                _ = v142
                                                                            } else {
                                                                                var v143 int8 = v30
                                                                                var v144 string = v21
                                                                                v83 = v2
                                                                                _ = v143
                                                                                _ = v144
                                                                            }
                                                                            _ = v126
                                                                        }
                                                                        _ = v119
                                                                    }
                                                                    _ = v97
                                                                }
                                                            }
                                                            _ = v83
                                                            _ = v84
                                                        }
                                                        fmt.Println(v4)
                                                    }
                                                }
                                                fmt.Println(v11, v21, fi20)
                                            }
                                            _ = v30
                                        }
                                    default:
                                        fmt.Println(v21, v26, v14)
                                    }
                                    fmt.Println(v21, v4, fi20)
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                }
                            case -1:
                                switch fi20 {
                                case 127:
                                    var v145 int = (-(fi20))
                                    _ = v145
                                case -543:
                                    for fi146 := 0; fi146 < 7; fi146++ {
                                        v1[0] = (uint8(v2) > (v10 % (((^(v10))) | 1)))
                                        delete(v1, 0)
                                        v1[909] = false
                                        var v147 int64 = v3
                                        v147 -= v3
                                        _ = fi146
                                        _ = v147
                                    }
                                case 11:
                                    fmt.Println(v2)
                                }
                            case 32767:
                                v21 = (("a" + (v21 + "a")) + v21)
                            }
                        }
                    } else {
                        fmt.Println(v14)
                        fmt.Println(v4, v10, v3)
                    }
                    var v148 bool = ((^(v10)) == v10)
                    v3 = v4
                    _ = v148
                } else {
                    fmt.Println(v3, v21, v10)
                }
                _ = fi20
                _ = v21
            }
            if v11 {
                v11 = (65535 > 320)
                var v149 float64 = 1.5
                _ = v149
            }
        }
    } else {
        for fi150 := 0; fi150 < 11; fi150++ {
            var v151 int = fi150
            for fi152 := 0; fi152 < 15; fi152++ {
                switch 32767 {
                case 127:
                    var v153 float64 = (1e10 + float64(0.1))
                    _ = v153
                case 4294967295:
                    v154 := GenS13{F0: 1, F1: 42, F2: 9223372036854775807}
                    _ = v154
                case 32767:
                    v3 += v3
                }
                if v11 {
                    _ = len(v1)
                    v151 += (v151 & (fi152 << ((65535) & 63)))
                }
                fmt.Println(v14, v14, v14)
                _ = fi152
            }
            v155 := []int{}
            _ = fi150
            _ = v151
            _ = v155
        }
    }
    v156 := GenS2{F0: 2147483647, F1: 39}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v156
    _ = v2
    _ = v3
    _ = v4
}
