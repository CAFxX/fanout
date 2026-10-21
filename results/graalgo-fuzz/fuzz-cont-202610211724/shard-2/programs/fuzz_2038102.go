package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint32; F2 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 int; F1 int32; F2 int }
type GenS7 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS8 struct { F0 uint16; F1 int; F2 bool }
type GenS9 struct { F0 string; F1 uint32; F2 int16 }
type GenS10 struct { F0 string; F1 uint; F2 uint8 }
type GenS11 struct { F0 uint32; F1 uint32 }
type GenS12 struct { F0 bool; F1 string; F2 int16 }
type GenS13 struct { F0 int64; F1 bool }
type GenS14 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS15 struct { F0 bool; F1 int64 }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 bool; F1 int32 }
type GenS18 struct { F0 int64; F1 string; F2 int64 }
type GenS19 struct { F0 uint8 }
func genFn0(p0 int8, p1 uint) uint64 {
    if (true || true) {
        if (65535 > 1) {
            v7 := []int{4294967295, 0, 371, 42}
            _ = v7
        } else {
            var v8 int8 = 42
            switch "\n\t" {
            case "a":
                var v9 int = -890
                _ = v9
            default:
                v8 -= -128
            }
            for fi10 := 0; fi10 < 20; fi10++ {
                fmt.Println(v8, p1)
                v8 += (v8 - 100)
                _ = fi10
            }
            _ = v8
        }
        p1 = (p1 << ((p1) & 63))
    }
    if false {
        switch -120 {
        case 32768:
            if true {
                p0 = 83
                var v11 string = "a"
                for fi12 := 0; fi12 < 16; fi12++ {
                    p1 = p1
                    var v13 uint8 = 0
                    if (1000 < 457) {
                        switch fi12 {
                        case -2147483648:
                            var v14 uint16 = 255
                            _ = v14
                        default:
                            v13 += (^(v13))
                        }
                    } else {
                        v13 = v13
                        var v15 string = v11
                        var v16 bool = false
                        _ = v15
                        _ = v16
                    }
                    _ = fi12
                    _ = v13
                }
                _ = v11
            }
        }
        var v17 int16 = 42
        _ = v17
    }
    _ = p0
    _ = p1
    return 955
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 4294967295
    v3 := make(map[string]int)
    var v4 uint8 = 21
    var v5 bool = false
    var v6 uint8 = v4
    v18 := GenS1{F0: 309, F1: false, F2: 2147483648}
    switch -9223372036854775808 {
    case 854:
        if false {
            switch ("hello" + "foo bar") {
            case "":
                if false {
                    v2 = v2
                }
            case "\n\t":
                var v19 int = 119
                _ = v19
            }
        }
    case 127:
        v2 = genFn0(-128, 1)
    default:
        var v20 int16 = 42
        _ = v20
    }
    v4 = 0
    var v21 int32 = 276
    var v22 uint8 = 255
    fmt.Println(v6, v6)
    for fi23 := 0; fi23 < 12; fi23++ {
        v24 := []int{(fi23 >> ((42) & 63)), (fi23 * fi23)}
        _ = fi23
        _ = v24
    }
    var v25 uint16 = 1
    v4 -= 127
    v26 := GenS2{F0: int64(v4), F1: 1, F2: -895}
    var v27 int8 = 66
    if v5 {
        if v5 {
            var v28 uint32 = 647
            _ = v28
        } else {
            v3["hello"] = 127
            v5 = (!(-136 > 128))
        }
    } else {
        v29 := GenS2{F0: 0, F1: 676, F2: 32768}
        _ = v29
    }
    fmt.Println(v2, v22)
    v30 := []int{int(v4), -869, 880}
    if (!(!v5)) {
        v31 := GenS2{F0: 1, F1: 2147483647, F2: 4294967295}
        if (!false) {
            v30 = append(v30, 0)
        } else {
            v6 += v6
            if v5 {
                v27 = v27
                for fi32 := 0; fi32 < 11; fi32++ {
                    for fi33 := 0; fi33 < 5; fi33++ {
                        v34 := GenS3{F0: v27}
                        switch (fi32 | fi32) {
                        case -2147483648:
                            _ = len(v1)
                        case 42:
                            switch -601 {
                            case 42:
                                v35 := GenS1{F0: 265, F1: v5, F2: 2}
                                _ = v35
                            case 32767:
                                v36 := GenS3{F0: v27}
                                var v37 int16 = 1000
                                var v38 float64 = (-2.25 + 3.14159)
                                var v39 string = ("a" + "\n\t")
                                var v40 string = v39
                                if v5 {
                                    var v41 int8 = 42
                                    v5 = v5
                                    var v42 float64 = (((v38 - (v38 * (v38 / 1.5))) - 1e10) * float64(0.0))
                                    _ = v41
                                    _ = v42
                                } else {
                                    var v43 int32 = int32(v25)
                                    var v44 int64 = -9223372036854775808
                                    _ = v43
                                    _ = v44
                                }
                                _ = v36
                                _ = v37
                                _ = v38
                                _ = v39
                                _ = v40
                            default:
                                v4 = v6
                                var v45 int8 = (v27 % (((v27 << ((514) & 63))) | 1))
                                v25 -= 857
                                if v5 {
                                    v3[("foo bar" + ("hello" + "\n\t"))] = fi33
                                    var v46 bool = ("" >= ("a" + "a"))
                                    v21 = (v21 >> ((uint(v4)) & 63))
                                    _ = v46
                                }
                                _ = v45
                            }
                        case -452:
                            v47 := v3[""]
                            v48 := GenS4{F0: 138}
                            _ = v47
                            _ = v48
                        default:
                            switch fi32 {
                            case 0:
                                switch ("" + "hello") {
                                case "foo bar":
                                    switch "\n\t" {
                                    case "a":
                                        if v5 {
                                            if len(v30) > 0 {
                                                v30[0] = (-(fi32))
                                            }
                                        } else {
                                            var v49 uint16 = v25
                                            _ = v49
                                        }
                                    case "hello":
                                        var v50 string = ""
                                        _ = v50
                                    default:
                                        if v5 {
                                            if v5 {
                                                if len(v30) >= 2 {
                                                    v30 = v30[1:len(v30)]
                                                }
                                            } else {
                                                if true {
                                                    var v51 uint64 = v2
                                                    _ = v51
                                                }
                                                v52 := v3["foo bar"]
                                                _ = v52
                                            }
                                            switch (fi32 << ((1) & 63)) {
                                            case -9223372036854775808:
                                                switch "hello" {
                                                case "foo bar":
                                                    var v53 uint8 = (v22 & 51)
                                                    _ = v53
                                                default:
                                                    v1[65535] = "hello"
                                                }
                                            case 32768:
                                                v54 := GenS5{F0: genFn0(v27, 255), F1: genFn0(v27, 255)}
                                                _ = v54
                                            case 9223372036854775807:
                                                var v55 int8 = (v27 * v27)
                                                _ = v55
                                            }
                                            v56 := v1[2]
                                            _ = len(v1)
                                            _ = v56
                                        } else {
                                            if true {
                                                v4 -= (v4 / (((v22 >> ((0) & 63))) | 1))
                                            } else {
                                                v27 -= v27
                                                switch "hello" {
                                                case "foo bar":
                                                    v57 := v1[4294967295]
                                                    _ = v57
                                                case "hello":
                                                    v58 := GenS1{F0: 2, F1: (42 <= (v21 & -930)), F2: 674}
                                                    _ = v58
                                                case "\n\t":
                                                    v2 = v2
                                                default:
                                                    switch "" {
                                                    case "\n\t":
                                                        if true {
                                                            v6 = (v6 >> ((65535) & 63))
                                                        }
                                                        v22 -= v4
                                                        v59 := GenS2{F0: 339, F1: 47, F2: fi32}
                                                        switch (fi33 + fi33) {
                                                        case -1:
                                                            fmt.Println(v5)
                                                        case 1:
                                                            if ((v21 ^ v21) == 42) {
                                                                v25 -= uint16(v2)
                                                                switch "" {
                                                                case "a":
                                                                    if len(v30) > 0 {
                                                                        v30[0] = 4294967295
                                                                    }
                                                                    var v60 float64 = (1.5 - 0.0)
                                                                    _ = v60
                                                                case "\n\t":
                                                                    var v61 int16 = 42
                                                                    _ = v61
                                                                }
                                                                fmt.Println(v21)
                                                                if ((v5 || true) || v5) {
                                                                    if (!v5) {
                                                                        if (-913 != 275) {
                                                                            var v62 uint = 714
                                                                            v63 := GenS6{F0: fi33, F1: v21, F2: 9223372036854775807}
                                                                            _ = v62
                                                                            _ = v63
                                                                        }
                                                                        v64 := GenS7{F0: -1, F1: uint32(v25), F2: genFn0(v27, 288)}
                                                                        v65 := GenS2{F0: 9223372036854775807, F1: 42, F2: 32767}
                                                                        var v66 int = fi32
                                                                        _ = v64
                                                                        _ = v65
                                                                        _ = v66
                                                                    } else {
                                                                        v67 := GenS6{F0: fi32, F1: -2147483648, F2: fi33}
                                                                        _ = v67
                                                                    }
                                                                    var v68 uint16 = v25
                                                                    v25 = 873
                                                                    _ = v68
                                                                } else {
                                                                    if (-214 == 234) {
                                                                        var v69 float32 = float32(((1.5 + 0.0) + 0.1))
                                                                        if v5 {
                                                                            var v70 int16 = 1
                                                                            var v71 float64 = (1e10 - float64(0.0))
                                                                            var v72 uint16 = v25
                                                                            _ = v70
                                                                            _ = v71
                                                                            _ = v72
                                                                        }
                                                                        if (!(1e10 == (float64((1e10 * v69)) + (0.1 / 1.0)))) {
                                                                            fmt.Println(v69)
                                                                            v21 = int32(v2)
                                                                        } else {
                                                                            var v73 float32 = (((1.5 - (v69 - 1e10)) / v69) + v69)
                                                                            var v74 uint8 = v22
                                                                            _ = v73
                                                                            _ = v74
                                                                        }
                                                                        _ = v69
                                                                    } else {
                                                                        v75 := GenS6{F0: (fi33 & 128), F1: (v21 ^ -185), F2: fi32}
                                                                        _ = v75
                                                                    }
                                                                    if v5 {
                                                                        v76 := GenS2{F0: 9223372036854775807, F1: 42, F2: int(v27)}
                                                                        _ = v76
                                                                    }
                                                                    if (!v5) {
                                                                        _ = len(v3)
                                                                        if false {
                                                                            var v77 int = (fi32 - (fi32 + fi32))
                                                                            if ((float32(float64(((1.5 + 1e10) - (1.5 * 0.1)))) * 0.0) <= 1e10) {
                                                                                if (!(42 <= (v6 + 127))) {
                                                                                    if v5 {
                                                                                        var v78 int = fi32
                                                                                        _ = v78
                                                                                    } else {
                                                                                        var v79 float64 = (0.0 * ((0.1 + 3.14159) + 3.14159))
                                                                                        _ = v79
                                                                                    }
                                                                                } else {
                                                                                    var v80 uint16 = (v25 + (v25 % (((v25 << ((42) & 63))) | 1)))
                                                                                    fmt.Println(v4)
                                                                                    _ = v80
                                                                                }
                                                                                if len(v30) >= 2 {
                                                                                    v30 = v30[1:len(v30)]
                                                                                }
                                                                                var v81 uint16 = (v25 - v25)
                                                                                var v82 uint = 4294967295
                                                                                _ = v81
                                                                                _ = v82
                                                                            }
                                                                            if ((v25 % (((v25 / (((^(v25))) | 1))) | 1)) < v25) {
                                                                                v83 := GenS6{F0: (fi33 / ((fi32) | 1)), F1: v21, F2: fi33}
                                                                                switch (fi33 >> ((742) & 63)) {
                                                                                case -1:
                                                                                    var v84 bool = ((fi32 / (((v77 ^ (v77 << ((uint(v25)) & 63)))) | 1)) < (v77 / ((int(v21)) | 1)))
                                                                                    _ = v84
                                                                                default:
                                                                                    _ = len(v1)
                                                                                }
                                                                                _ = v83
                                                                            } else {
                                                                                if ((0.1 > -2.25) || false) {
                                                                                    _ = len(v30)
                                                                                }
                                                                                if (v5 || ((v27 << ((503) & 63)) > -1)) {
                                                                                    switch "" {
                                                                                    case "foo bar":
                                                                                        switch 9223372036854775807 {
                                                                                        case 127:
                                                                                            if false {
                                                                                                v21 = (v21 >> ((255) & 63))
                                                                                                switch ((((("\n\t" + "hello") + "") + "a") + "a") + "foo bar") {
                                                                                                case "":
                                                                                                    fmt.Println(v5, v21, v2)
                                                                                                case "hello":
                                                                                                    v22 += 0
                                                                                                default:
                                                                                                    var v85 uint16 = (v25 - (v25 - v25))
                                                                                                    _ = v85
                                                                                                }
                                                                                                v86 := GenS3{F0: (^(v27))}
                                                                                                if v5 {
                                                                                                    v87 := GenS8{F0: 295, F1: fi33, F2: (713 < v25)}
                                                                                                    var v88 uint16 = (v25 / (((v25 << ((288) & 63))) | 1))
                                                                                                    switch 128 {
                                                                                                    case -9223372036854775808:
                                                                                                        fmt.Println(v88, v27, v21)
                                                                                                    case 9223372036854775807:
                                                                                                        if false {
                                                                                                            v89 := GenS5{F0: v2, F1: (v2 % ((v2) | 1))}
                                                                                                            _ = v89
                                                                                                        } else {
                                                                                                            if (!v5) {
                                                                                                                var v90 float64 = 1.5
                                                                                                                var v91 uint8 = 167
                                                                                                                var v92 uint = 255
                                                                                                                v93 := GenS9{F0: "\n\t", F1: 2147483648, F2: -32768}
                                                                                                                v21 = v21
                                                                                                                _ = v90
                                                                                                                _ = v91
                                                                                                                _ = v92
                                                                                                                _ = v93
                                                                                                            }
                                                                                                            if ((-(v27)) >= v27) {
                                                                                                                var v94 bool = (!v5)
                                                                                                                _ = v94
                                                                                                            } else {
                                                                                                                var v95 int64 = -9223372036854775808
                                                                                                                v77 = fi32
                                                                                                                v88 -= v25
                                                                                                                _ = v95
                                                                                                            }
                                                                                                            switch (fi33 - -9223372036854775808) {
                                                                                                            case 32768:
                                                                                                                var v96 int16 = 0
                                                                                                                _ = v96
                                                                                                            case 1:
                                                                                                                if (false || true) {
                                                                                                                    var v97 int16 = 1
                                                                                                                    fmt.Println(v77, fi33)
                                                                                                                    _ = v97
                                                                                                                }
                                                                                                            case -255:
                                                                                                                if true {
                                                                                                                    if ((v21 + -1) == (v21 + (v21 * -1))) {
                                                                                                                        fmt.Println(v6)
                                                                                                                        if false {
                                                                                                                            if v5 {
                                                                                                                                v27 -= 1
                                                                                                                            }
                                                                                                                            fmt.Println(v5)
                                                                                                                        }
                                                                                                                        v88 = v88
                                                                                                                    }
                                                                                                                    switch "foo bar" {
                                                                                                                    case "":
                                                                                                                        var v98 float32 = float32((3.14159 - 0.0))
                                                                                                                        _ = v98
                                                                                                                    }
                                                                                                                }
                                                                                                            default:
                                                                                                                var v99 uint8 = v22
                                                                                                                _ = v99
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v87
                                                                                                    _ = v88
                                                                                                }
                                                                                                v2 += v2
                                                                                                fmt.Println(v4)
                                                                                                v100 := GenS2{F0: -1, F1: 1, F2: int(v25)}
                                                                                                var v101 int8 = (v27 / ((42) | 1))
                                                                                                _ = v100
                                                                                                _ = v101
                                                                                                _ = v86
                                                                                            } else {
                                                                                                if v5 {
                                                                                                    var v102 uint8 = (^(v22))
                                                                                                    _ = len(v30)
                                                                                                    _ = v102
                                                                                                }
                                                                                            }
                                                                                        case 32768:
                                                                                            var v103 uint64 = (^(v2))
                                                                                            _ = v103
                                                                                        default:
                                                                                            v104 := GenS9{F0: "foo bar", F1: 2147483647, F2: 0}
                                                                                            _ = v104
                                                                                        }
                                                                                    case "a":
                                                                                        v2 -= (^(v2))
                                                                                    case "":
                                                                                        v30 = append(v30, fi32)
                                                                                    }
                                                                                    fmt.Println(v77, fi33)
                                                                                }
                                                                                if v5 {
                                                                                    v21 = v21
                                                                                    if true {
                                                                                        if (4294967295 < 701) {
                                                                                            fmt.Println(v5)
                                                                                            v5 = false
                                                                                            var v105 float32 = float32(1e10)
                                                                                            _ = v105
                                                                                        } else {
                                                                                            v21 += v21
                                                                                            var v106 int16 = 0
                                                                                            fmt.Println(fi32, v4)
                                                                                            _ = v106
                                                                                        }
                                                                                        var v107 bool = true
                                                                                        _ = v107
                                                                                    } else {
                                                                                        v6 = (^(v4))
                                                                                    }
                                                                                } else {
                                                                                    var v108 int = fi33
                                                                                    v109 := GenS10{F0: "hello", F1: 0, F2: v22}
                                                                                    if len(v30) > 0 {
                                                                                        v30[0] = fi33
                                                                                    }
                                                                                    var v110 int = 878
                                                                                    _ = v108
                                                                                    _ = v109
                                                                                    _ = v110
                                                                                }
                                                                            }
                                                                            v25 = (v25 / ((uint16(v6)) | 1))
                                                                            if false {
                                                                                var v111 int = -9223372036854775808
                                                                                switch (v111 >> ((535) & 63)) {
                                                                                case 654:
                                                                                    var v112 uint8 = v6
                                                                                    v113 := GenS10{F0: "hello", F1: 1, F2: (^(v112))}
                                                                                    _ = v112
                                                                                    _ = v113
                                                                                default:
                                                                                    var v114 bool = (!v5)
                                                                                    _ = v114
                                                                                }
                                                                                _ = v111
                                                                            } else {
                                                                                v115 := GenS11{F0: 2147483648, F1: 42}
                                                                                _ = v115
                                                                            }
                                                                            if v5 {
                                                                                if ((!true) || v5) {
                                                                                    var v116 int8 = 1
                                                                                    v27 -= (v116 + (v27 ^ 127))
                                                                                    _ = v116
                                                                                }
                                                                            } else {
                                                                                v2 = v2
                                                                                switch 32767 {
                                                                                case 2147483647:
                                                                                    if (9223372036854775807 < 40) {
                                                                                        _ = len(v30)
                                                                                        var v117 uint16 = v25
                                                                                        if len(v30) > 0 {
                                                                                            v30[0] = fi32
                                                                                        }
                                                                                        _ = v117
                                                                                    }
                                                                                case -2147483648:
                                                                                    _ = len(v30)
                                                                                case 42:
                                                                                    v118 := GenS12{F0: false, F1: ("" + "foo bar"), F2: -518}
                                                                                    _ = v118
                                                                                }
                                                                                switch ((("\n\t" + "\n\t") + "\n\t") + "foo bar") {
                                                                                case "":
                                                                                    v119 := GenS3{F0: (^(v27))}
                                                                                    if len(v30) > 0 {
                                                                                        v30[0] = (^(v77))
                                                                                    }
                                                                                    _ = v119
                                                                                case "foo bar":
                                                                                    v1[2147483647] = "hello"
                                                                                }
                                                                            }
                                                                            if ((v4 - 0) > (v6 | 2)) {
                                                                                if true {
                                                                                    if (!((genFn0(-128, 18446744073709551615) <= 2) && true)) {
                                                                                        var v120 float32 = 1.5
                                                                                        if v5 {
                                                                                            if (!((((v120 / ((((-2.25 + v120) - -2.25) - (-2.25 / 1.0)) + 1.5)) + (((v120 - float32(1.5)) / float32(1.5)) + v120)) + v120) <= ((((v120 / v120) * v120) + (0.0 / v120)) / 3.14159))) {
                                                                                                if (fi32 == (fi33 >> ((65535) & 63))) {
                                                                                                    var v121 float32 = 3.14159
                                                                                                    if ((v77 + (v77 - (fi32 << ((2) & 63)))) != 128) {
                                                                                                        _ = len(v3)
                                                                                                        switch v77 {
                                                                                                        case 2147483647:
                                                                                                            var v122 float32 = v121
                                                                                                            _ = v122
                                                                                                        }
                                                                                                        v123 := GenS11{F0: 2147483647, F1: 65535}
                                                                                                        _ = v123
                                                                                                    }
                                                                                                    var v124 uint8 = 127
                                                                                                    v5 = v5
                                                                                                    var v125 bool = (v5 || false)
                                                                                                    _ = v121
                                                                                                    _ = v124
                                                                                                    _ = v125
                                                                                                }
                                                                                                if true {
                                                                                                    if v5 {
                                                                                                        switch ((("\n\t" + ("foo bar" + "foo bar")) + "a") + "\n\t") {
                                                                                                        case "\n\t":
                                                                                                            fmt.Println(v5, v4)
                                                                                                        default:
                                                                                                            fmt.Println(fi33, v77)
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v126 uint8 = v4
                                                                                                        _ = v126
                                                                                                    }
                                                                                                    if v5 {
                                                                                                        switch 4294967295 {
                                                                                                        case 827:
                                                                                                            _ = len(v30)
                                                                                                        case -2147483648:
                                                                                                            v120 += (v120 + v120)
                                                                                                        case 4294967295:
                                                                                                            var v127 int32 = (v21 + v21)
                                                                                                            _ = v127
                                                                                                        }
                                                                                                        fmt.Println(v6)
                                                                                                        var v128 uint32 = 65535
                                                                                                        _ = v128
                                                                                                    } else {
                                                                                                        _ = len(v3)
                                                                                                        var v129 int8 = v27
                                                                                                        _ = v129
                                                                                                    }
                                                                                                    switch fi33 {
                                                                                                    case 9223372036854775807:
                                                                                                        if (v5 || true) {
                                                                                                            v130 := GenS8{F0: v25, F1: (v77 + v77), F2: v5}
                                                                                                            var v131 int = fi32
                                                                                                            v132 := GenS13{F0: 805, F1: false}
                                                                                                            _ = v130
                                                                                                            _ = v131
                                                                                                            _ = v132
                                                                                                        } else {
                                                                                                            v77 -= (fi32 | (fi32 - -370))
                                                                                                            switch "a" {
                                                                                                            case "\n\t":
                                                                                                                var v133 uint64 = genFn0((v27 >> ((4294967295) & 63)), 206)
                                                                                                                _ = v133
                                                                                                            case "hello":
                                                                                                                v134 := GenS14{F0: uint16(v6), F1: 2147483648, F2: (v27 % ((v27) | 1))}
                                                                                                                _ = v134
                                                                                                            case "a":
                                                                                                                fmt.Println(v2, fi33)
                                                                                                            default:
                                                                                                                v6 -= 255
                                                                                                            }
                                                                                                            var v135 uint64 = (v2 + 2)
                                                                                                            _ = v135
                                                                                                        }
                                                                                                        if false {
                                                                                                            v136 := GenS15{F0: false, F1: -113}
                                                                                                            var v137 bool = (52 >= 32767)
                                                                                                            v30 = append(v30, v77)
                                                                                                            _ = v136
                                                                                                            _ = v137
                                                                                                        } else {
                                                                                                            var v138 int16 = -32768
                                                                                                            _ = v138
                                                                                                        }
                                                                                                    case 2147483647:
                                                                                                        switch fi33 {
                                                                                                        case 2147483647:
                                                                                                            var v139 int32 = (v21 & -1)
                                                                                                            _ = v139
                                                                                                        case 127:
                                                                                                            fmt.Println(v21, v21)
                                                                                                            if (!v5) {
                                                                                                                var v140 int32 = v21
                                                                                                                if len(v30) >= 2 {
                                                                                                                    v30 = v30[1:len(v30)]
                                                                                                                }
                                                                                                                var v141 uint = 18446744073709551615
                                                                                                                _ = v140
                                                                                                                _ = v141
                                                                                                            }
                                                                                                            v142 := GenS16{F0: genFn0(127, 303)}
                                                                                                            _ = v142
                                                                                                        case -2147483648:
                                                                                                            var v143 float32 = 1.5
                                                                                                            _ = v143
                                                                                                        }
                                                                                                    case -9223372036854775808:
                                                                                                        switch (fi33 & (^(v77))) {
                                                                                                        case 46:
                                                                                                            switch "" {
                                                                                                            case "":
                                                                                                                v77 = -715
                                                                                                            case "a":
                                                                                                                var v144 bool = (!true)
                                                                                                                _ = v144
                                                                                                            default:
                                                                                                                v22 -= v6
                                                                                                            }
                                                                                                        case -38:
                                                                                                            v145 := v3["\n\t"]
                                                                                                            _ = v145
                                                                                                        default:
                                                                                                            v77 = fi32
                                                                                                        }
                                                                                                    }
                                                                                                    v146 := GenS17{F0: v5, F1: (v21 & 2147483647)}
                                                                                                    _ = v146
                                                                                                } else {
                                                                                                    var v147 int16 = -1
                                                                                                    _ = v147
                                                                                                }
                                                                                            } else {
                                                                                                fmt.Println(v21, v21, v6)
                                                                                            }
                                                                                            v148 := GenS8{F0: 0, F1: 0, F2: (-1 <= 1)}
                                                                                            v149 := GenS16{F0: (v2 ^ (^(v2)))}
                                                                                            _ = v148
                                                                                            _ = v149
                                                                                        }
                                                                                        var v150 uint8 = (v22 & 1)
                                                                                        var v151 uint8 = v6
                                                                                        _ = len(v1)
                                                                                        _ = v120
                                                                                        _ = v150
                                                                                        _ = v151
                                                                                    } else {
                                                                                        switch fi33 {
                                                                                        case 1:
                                                                                            v152 := GenS18{F0: -9223372036854775808, F1: "\n\t", F2: -926}
                                                                                            _ = v152
                                                                                        case -2147483648:
                                                                                            fmt.Println(v6, v6, fi32)
                                                                                        default:
                                                                                            if (v21 != (v21 - v21)) {
                                                                                                if len(v30) >= 2 {
                                                                                                    v30 = v30[1:len(v30)]
                                                                                                }
                                                                                                v4 += v22
                                                                                                var v153 uint64 = (v2 % (((v2 * v2)) | 1))
                                                                                                v4 += (v4 + (^(v6)))
                                                                                                var v154 int = (v77 + v77)
                                                                                                _ = v153
                                                                                                _ = v154
                                                                                            } else {
                                                                                                var v155 int16 = 13
                                                                                                if ((v21 % (((v21 * v21)) | 1)) > (v21 * 42)) {
                                                                                                    _ = len(v3)
                                                                                                }
                                                                                                var v156 int64 = -1
                                                                                                _ = v155
                                                                                                _ = v156
                                                                                            }
                                                                                        }
                                                                                        var v157 int = 672
                                                                                        _ = v157
                                                                                    }
                                                                                    var v158 int = fi33
                                                                                    v22 -= 43
                                                                                    _ = v158
                                                                                }
                                                                                v159 := GenS7{F0: 42, F1: uint32(v2), F2: (v2 * genFn0((^(v27)), 765))}
                                                                                v160 := GenS16{F0: 593}
                                                                                _ = v159
                                                                                _ = v160
                                                                            }
                                                                            fmt.Println(fi33, v21, fi32)
                                                                            v77 += 1
                                                                            _ = v77
                                                                        }
                                                                        v6 = (v4 * v6)
                                                                        if len(v30) > 0 {
                                                                            v30[0] = (-(fi32))
                                                                        }
                                                                    }
                                                                    fmt.Println(v5, v25, fi32)
                                                                }
                                                            }
                                                        }
                                                        var v161 float32 = 3.14159
                                                        v21 -= (v21 / (((v21 % (((v21 | -2147483648)) | 1))) | 1))
                                                        _ = v161
                                                        _ = v59
                                                    case "foo bar":
                                                        var v162 bool = (v5 && v5)
                                                        _ = v162
                                                    default:
                                                        var v163 float64 = (float64(1.5) + (0.0 + -2.25))
                                                        _ = v163
                                                    }
                                                }
                                                v4 = (^(v4))
                                            }
                                        }
                                    }
                                case "":
                                    if ((v5 && v5) && ((!v5) && v5)) {
                                        v164 := GenS19{F0: v22}
                                        v6 = (v6 % ((163) | 1))
                                        fmt.Println(v4, v2, fi32)
                                        _ = v164
                                    } else {
                                        var v165 uint64 = (v2 ^ (^(v2)))
                                        _ = v165
                                    }
                                case "a":
                                    v166 := GenS2{F0: -625, F1: 177, F2: (fi33 << ((42) & 63))}
                                    _ = v166
                                }
                            }
                        }
                        _ = fi33
                        _ = v34
                    }
                    var v167 int32 = 0
                    _ = fi32
                    _ = v167
                }
                var v168 int32 = (v21 << ((4294967295) & 63))
                _ = v168
            } else {
                fmt.Println(v4, v2)
            }
        }
        v21 -= v21
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
