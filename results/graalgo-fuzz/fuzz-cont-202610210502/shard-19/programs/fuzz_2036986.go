package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int32; F1 bool }
type GenS8 struct { F0 uint64; F1 int8 }
type GenS9 struct { F0 int64; F1 uint8 }
type GenS10 struct { F0 int8; F1 uint32; F2 int32 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 bool; F1 int; F2 uint32 }
type GenS13 struct { F0 int8; F1 int8; F2 uint8 }
type GenS14 struct { F0 uint; F1 uint8; F2 uint }
type GenS15 struct { F0 int }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 int; F1 uint64 }
type GenS18 struct { F0 int16 }
type GenS19 struct { F0 int }
type GenS20 struct { F0 uint32; F1 bool; F2 int }
type GenS21 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS22 struct { F0 string; F1 int }
type GenS23 struct { F0 bool }

func main() {
    var v1 float64 = float64(((-2.25 + (3.14159 * -2.25)) / -2.25))
    var v2 float64 = v1
    var v3 int16 = -1
    var v4 int16 = 32767
    v5 := make(map[int]int)
    v6 := []int{-2147483648, -500}
    var v7 string = ""
    v8 := GenS1{F0: 692, F1: 127, F2: -1}
    v3 -= v4
    fmt.Println(v7)
    v6 = append(v6, 4294967295)
    var v9 string = ("hello" + v7)
    if true {
        fmt.Println(v3, v7)
        fmt.Println(v1, v2, v7)
        fmt.Println(v1)
    } else {
        var v10 float64 = v1
        var v11 bool = false
        if v11 {
            if (!v11) {
                switch (v9 + v7) {
                case "":
                    if v11 {
                        v6 = append(v6, -1)
                    } else {
                        v4 += (v4 * (v3 >> ((278) & 63)))
                        var v12 float64 = 1.5
                        v13 := GenS1{F0: 2147483648, F1: -66, F2: 121}
                        _ = v12
                        _ = v13
                    }
                case "hello":
                    var v14 uint64 = 2
                    _ = v14
                }
            }
        }
        _ = v10
        _ = v11
    }
    fmt.Println(v7, v3)
    v1 -= v1
    var v15 int = -876
    for fi16 := 0; fi16 < 1; fi16++ {
        v9 = (("\n\t" + (v7 + "a")) + v9)
        v17 := GenS2{F0: 9223372036854775807}
        var v18 int16 = v4
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 uint = 914
    v20 := v5[(v15 - v15)]
    if (363 != 4294967295) {
        var v21 float32 = (1e10 * 3.14159)
        for fi22 := 0; fi22 < 7; fi22++ {
            if (true && true) {
                var v23 uint16 = 250
                v9 = v9
                if len(v6) > 0 {
                    v6[0] = (v20 / ((v15) | 1))
                }
                _ = v23
            } else {
                switch v9 {
                case "a":
                    for fi24 := 0; fi24 < 8; fi24++ {
                        v9 = "foo bar"
                        if true {
                            if (!true) {
                                switch fi24 {
                                case 0:
                                    switch fi22 {
                                    case -275:
                                        v4 = v4
                                    case -145:
                                        var v25 int8 = -22
                                        _ = v25
                                    case 32768:
                                        fmt.Println(v19)
                                    }
                                default:
                                    fmt.Println(v19, v3)
                                }
                                fmt.Println(v9, v15)
                                v26 := GenS3{F0: 2147483647}
                                _ = v26
                            } else {
                                switch v20 {
                                case -896:
                                    switch v7 {
                                    case "":
                                        fmt.Println(fi22)
                                    case "hello":
                                        v4 -= v4
                                    default:
                                        var v27 int = 9223372036854775807
                                        _ = v27
                                    }
                                case 32768:
                                    if (0 < 85) {
                                        v20 = (^(v15))
                                    } else {
                                        var v28 bool = ((v1 - v2) <= (v1 / v1))
                                        v4 += (v4 << ((0) & 63))
                                        _ = v28
                                    }
                                case -723:
                                    switch v20 {
                                    case 128:
                                        switch 93 {
                                        case 101:
                                            var v29 uint16 = 125
                                            _ = v29
                                        case -595:
                                            var v30 uint64 = 4294967295
                                            fmt.Println(v4, v15)
                                            v31 := GenS4{F0: "\n\t", F1: 333}
                                            v32 := GenS2{F0: 0}
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            var v33 string = "a"
                                            switch (v15 ^ -879) {
                                            case 32768:
                                                if (true && false) {
                                                    fmt.Println(v30, fi22, v20)
                                                    switch v7 {
                                                    case "":
                                                        switch v7 {
                                                        case "hello":
                                                            v34 := v5[fi24]
                                                            _ = v34
                                                        case "\n\t":
                                                            fmt.Println(v7, fi22)
                                                        case "a":
                                                            var v35 uint8 = 255
                                                            _ = v35
                                                        default:
                                                            fmt.Println(v7)
                                                        }
                                                    case "\n\t":
                                                        v2 = v1
                                                    default:
                                                        switch v7 {
                                                        case "":
                                                            var v36 int = (fi22 ^ fi24)
                                                            _ = v36
                                                        case "hello":
                                                            fmt.Println(v9, v20, fi24)
                                                        default:
                                                            if (fi22 >= v20) {
                                                                fmt.Println(v15, v19)
                                                                v37 := GenS5{F0: int16(fi22)}
                                                                switch v7 {
                                                                case "\n\t":
                                                                    var v38 uint32 = 1
                                                                    _ = v38
                                                                }
                                                                _ = v37
                                                            } else {
                                                                v7 = v7
                                                                switch v20 {
                                                                case 9223372036854775807:
                                                                    switch (fi24 << ((v19) & 63)) {
                                                                    case -2147483648:
                                                                        v39 := GenS4{F0: v33, F1: 42}
                                                                        _ = v39
                                                                    case -451:
                                                                        var v40 int8 = -128
                                                                        v40 = int8(v19)
                                                                        _ = v40
                                                                    case -257:
                                                                        _ = len(v6)
                                                                    }
                                                                case -9223372036854775808:
                                                                    if false {
                                                                        v5[v20] = (fi24 + v15)
                                                                        if (-2.25 > v2) {
                                                                            v5[v20] = v15
                                                                            if (1 > 1) {
                                                                                v6 = append(v6, (v15 - (v15 << ((v19) & 63))))
                                                                            } else {
                                                                                _ = len(v5)
                                                                                switch (fi22 & (fi24 % (((v20 + 0)) | 1))) {
                                                                                case 128:
                                                                                    if false {
                                                                                        if (0 != -9223372036854775808) {
                                                                                            if false {
                                                                                                v19 = v19
                                                                                            }
                                                                                        }
                                                                                        v6 = append(v6, (^(v20)))
                                                                                        var v41 uint32 = 65535
                                                                                        if true {
                                                                                            if (v2 > v1) {
                                                                                                _ = len(v5)
                                                                                                v42 := GenS1{F0: 0, F1: int8(v41), F2: -128}
                                                                                                _ = v42
                                                                                            } else {
                                                                                                v43 := GenS2{F0: 9223372036854775807}
                                                                                                _ = v43
                                                                                            }
                                                                                        }
                                                                                        v44 := GenS1{F0: (^(v41)), F1: -108, F2: 0}
                                                                                        v4 += (v4 | (^(v4)))
                                                                                        _ = v41
                                                                                        _ = v44
                                                                                    }
                                                                                    var v45 int64 = -980
                                                                                    _ = v45
                                                                                case 9223372036854775807:
                                                                                    fmt.Println(v2, v15, v1)
                                                                                case -120:
                                                                                    var v46 string = (v7 + (v9 + v33))
                                                                                    var v47 uint = v19
                                                                                    var v48 float32 = (1.5 / (float32(v2) + float32(v2)))
                                                                                    if false {
                                                                                        if false {
                                                                                            switch "foo bar" {
                                                                                            case "hello":
                                                                                                var v49 int32 = -11
                                                                                                _ = v49
                                                                                            default:
                                                                                                v33 = v46
                                                                                                if ((false || false) || (true || (2147483647 < 2147483647))) {
                                                                                                    if false {
                                                                                                        fmt.Println(fi22)
                                                                                                        if false {
                                                                                                            v5[(-(fi22))] = fi22
                                                                                                            var v50 string = ""
                                                                                                            var v51 float32 = float32((1e10 + (v1 / float64(v21))))
                                                                                                            delete(v5, fi22)
                                                                                                            _ = v50
                                                                                                            _ = v51
                                                                                                        } else {
                                                                                                            fmt.Println(fi24, v46, v4)
                                                                                                        }
                                                                                                        v21 = v21
                                                                                                    } else {
                                                                                                        fmt.Println(v3, v2, v20)
                                                                                                        if false {
                                                                                                            v52 := GenS1{F0: 2147483648, F1: -1, F2: 42}
                                                                                                            v53 := GenS6{F0: 0}
                                                                                                            switch (("hello" + v7) + (v46 + "a")) {
                                                                                                            case "":
                                                                                                                if true {
                                                                                                                    if (int(v3) != 1) {
                                                                                                                        v9 = ("a" + ("a" + v46))
                                                                                                                        if true {
                                                                                                                            var v54 bool = true
                                                                                                                            v47 += v47
                                                                                                                            v55 := GenS5{F0: (v4 * v3)}
                                                                                                                            _ = v54
                                                                                                                            _ = v55
                                                                                                                        }
                                                                                                                        if false {
                                                                                                                            var v56 float64 = v2
                                                                                                                            switch "a" {
                                                                                                                            case "\n\t":
                                                                                                                                v20 = fi22
                                                                                                                            case "foo bar":
                                                                                                                                if ((^(v30)) > (v30 ^ (v30 / ((18446744073709551615) | 1)))) {
                                                                                                                                    if true {
                                                                                                                                        fmt.Println(v33, v2)
                                                                                                                                        v1 = float64(-2.25)
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            case "hello":
                                                                                                                                v57 := GenS7{F0: -2147483648, F1: (2147483647 != -2147483648)}
                                                                                                                                _ = v57
                                                                                                                            }
                                                                                                                            _ = v56
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        v46 = v33
                                                                                                                        switch v46 {
                                                                                                                        case "hello":
                                                                                                                            var v58 int = fi22
                                                                                                                            _ = v58
                                                                                                                        default:
                                                                                                                            v59 := GenS4{F0: "a", F1: 2147483647}
                                                                                                                            _ = v59
                                                                                                                        }
                                                                                                                        v3 -= 1000
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    switch v7 {
                                                                                                                    case "foo bar":
                                                                                                                        v5[0] = fi22
                                                                                                                    case "a":
                                                                                                                        switch ((v7 + "\n\t") + v46) {
                                                                                                                        case "":
                                                                                                                            v60 := GenS8{F0: v30, F1: 0}
                                                                                                                            _ = v60
                                                                                                                        case "foo bar":
                                                                                                                            v3 = (v3 ^ v3)
                                                                                                                        case "\n\t":
                                                                                                                            v19 -= (^(v47))
                                                                                                                        default:
                                                                                                                            if false {
                                                                                                                                var v61 int64 = 42
                                                                                                                                _ = v61
                                                                                                                            }
                                                                                                                        }
                                                                                                                    case "":
                                                                                                                        v7 = "a"
                                                                                                                    default:
                                                                                                                        v3 += v3
                                                                                                                    }
                                                                                                                    v2 = -2.25
                                                                                                                    if false {
                                                                                                                        v62 := GenS2{F0: 0}
                                                                                                                        if (!(true && (42 <= 42))) {
                                                                                                                            v63 := GenS9{F0: 1, F1: 169}
                                                                                                                            _ = v63
                                                                                                                        } else {
                                                                                                                            switch (v15 * fi22) {
                                                                                                                            case 779:
                                                                                                                                fmt.Println(v2, v1, v30)
                                                                                                                            default:
                                                                                                                                if false {
                                                                                                                                    var v64 uint64 = (v30 * 47)
                                                                                                                                    v2 = 3.14159
                                                                                                                                    if (false && true) {
                                                                                                                                        switch v46 {
                                                                                                                                        case "hello":
                                                                                                                                            v4 += v3
                                                                                                                                        case "":
                                                                                                                                            v64 += (^(v64))
                                                                                                                                        }
                                                                                                                                        var v65 int32 = -1
                                                                                                                                        if false {
                                                                                                                                            var v66 int32 = (v65 >> ((v19) & 63))
                                                                                                                                            v67 := GenS5{F0: -412}
                                                                                                                                            _ = v66
                                                                                                                                            _ = v67
                                                                                                                                        }
                                                                                                                                        _ = v65
                                                                                                                                    }
                                                                                                                                    _ = v64
                                                                                                                                } else {
                                                                                                                                    if (((v9 + (v7 + "hello")) + "") < ("hello" + (v46 + (v9 + "a")))) {
                                                                                                                                        var v68 int16 = (^(v3))
                                                                                                                                        v69 := GenS6{F0: 443}
                                                                                                                                        switch v15 {
                                                                                                                                        case -9223372036854775808:
                                                                                                                                            fmt.Println(v19, v46, v33)
                                                                                                                                        case 2147483647:
                                                                                                                                            fmt.Println(v21)
                                                                                                                                        }
                                                                                                                                        _ = v68
                                                                                                                                        _ = v69
                                                                                                                                    } else {
                                                                                                                                        var v70 uint16 = 2
                                                                                                                                        fmt.Println(v30, v46, fi22)
                                                                                                                                        _ = v70
                                                                                                                                    }
                                                                                                                                    v71 := GenS7{F0: 0, F1: true}
                                                                                                                                    _ = v71
                                                                                                                                }
                                                                                                                            }
                                                                                                                            v1 = ((1e10 + v1) * (v2 / (v1 * v2)))
                                                                                                                            if len(v6) > 0 {
                                                                                                                                v6[0] = 128
                                                                                                                            }
                                                                                                                            if false {
                                                                                                                                switch v15 {
                                                                                                                                case -2147483648:
                                                                                                                                    if true {
                                                                                                                                        fmt.Println(v48, v2, v48)
                                                                                                                                        if (2 == 0) {
                                                                                                                                            v1 += v2
                                                                                                                                        } else {
                                                                                                                                            v4 += (v3 % (((v3 << ((uint(v30)) & 63))) | 1))
                                                                                                                                            delete(v5, (-(v15)))
                                                                                                                                        }
                                                                                                                                        var v72 bool = false
                                                                                                                                        _ = v72
                                                                                                                                    } else {
                                                                                                                                        switch (fi22 ^ v15) {
                                                                                                                                        case 1:
                                                                                                                                            var v73 int = (v20 ^ v20)
                                                                                                                                            _ = v73
                                                                                                                                        case 63:
                                                                                                                                            v74 := GenS10{F0: 42, F1: 4294967295, F2: 2147483647}
                                                                                                                                            _ = v74
                                                                                                                                        case 857:
                                                                                                                                            var v75 uint8 = 42
                                                                                                                                            _ = v75
                                                                                                                                        default:
                                                                                                                                            var v76 int = v20
                                                                                                                                            _ = v76
                                                                                                                                        }
                                                                                                                                        v47 += (v47 / (((^(v19))) | 1))
                                                                                                                                        if (2 < 281) {
                                                                                                                                            v77 := GenS10{F0: -128, F1: 490, F2: 0}
                                                                                                                                            _ = v77
                                                                                                                                        } else {
                                                                                                                                            fmt.Println(v30, v4, v33)
                                                                                                                                            v78 := GenS5{F0: v3}
                                                                                                                                            v79 := GenS1{F0: uint32(v19), F1: 0, F2: 0}
                                                                                                                                            _ = v78
                                                                                                                                            _ = v79
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                case 691:
                                                                                                                                    fmt.Println(v20)
                                                                                                                                case 42:
                                                                                                                                    v80 := GenS11{F0: int32(v4)}
                                                                                                                                    _ = v80
                                                                                                                                default:
                                                                                                                                    fmt.Println(v21, v2, v3)
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                var v81 int32 = -223
                                                                                                                                var v82 float32 = float32(3.14159)
                                                                                                                                if (v48 >= ((v82 + (v48 + 0.0)) - float32((v1 * v1)))) {
                                                                                                                                    var v83 string = "a"
                                                                                                                                    _ = v83
                                                                                                                                }
                                                                                                                                _ = v81
                                                                                                                                _ = v82
                                                                                                                            }
                                                                                                                            switch "" {
                                                                                                                            case "hello":
                                                                                                                                v9 = v33
                                                                                                                            case "":
                                                                                                                                fmt.Println(v7, v47, v9)
                                                                                                                            case "foo bar":
                                                                                                                                v15 += v15
                                                                                                                            }
                                                                                                                            v84 := GenS12{F0: ((v9 < v46) && (!true)), F1: fi22, F2: uint32(v19)}
                                                                                                                            _ = v84
                                                                                                                        }
                                                                                                                        v6 = append(v6, int(v3))
                                                                                                                        _ = v62
                                                                                                                    } else {
                                                                                                                        v85 := GenS1{F0: 741, F1: -128, F2: 42}
                                                                                                                        var v86 uint = 459
                                                                                                                        _ = v85
                                                                                                                        _ = v86
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            fmt.Println(v2)
                                                                                                            var v87 uint = v19
                                                                                                            _ = v52
                                                                                                            _ = v53
                                                                                                            _ = v87
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    v88 := GenS5{F0: v3}
                                                                                                    if (false && (false || false)) {
                                                                                                        v89 := GenS13{F0: -1, F1: -101, F2: 0}
                                                                                                        fmt.Println(v2)
                                                                                                        if true {
                                                                                                            if (((1e10 * v48) != v21) && (!true)) {
                                                                                                                var v90 int = (v20 % ((128) | 1))
                                                                                                                _ = v90
                                                                                                            }
                                                                                                            var v91 uint = v19
                                                                                                            v92 := GenS14{F0: v91, F1: 255, F2: v47}
                                                                                                            _ = v91
                                                                                                            _ = v92
                                                                                                        } else {
                                                                                                            switch "" {
                                                                                                            case "hello":
                                                                                                                if (2 > (v30 >> ((v19) & 63))) {
                                                                                                                    if false {
                                                                                                                        switch (v15 << ((440) & 63)) {
                                                                                                                        case 32767:
                                                                                                                            v47 -= 1
                                                                                                                        }
                                                                                                                        var v93 int8 = -57
                                                                                                                        var v94 int64 = 1
                                                                                                                        v3 -= (v4 | (v3 ^ 1))
                                                                                                                        fmt.Println(v30)
                                                                                                                        _ = v93
                                                                                                                        _ = v94
                                                                                                                    } else {
                                                                                                                        switch v9 {
                                                                                                                        case "a":
                                                                                                                            v48 += 1e10
                                                                                                                            var v95 int8 = int8(fi22)
                                                                                                                            if true {
                                                                                                                                fmt.Println(v15, v3, v15)
                                                                                                                                fmt.Println(v9, v48, v30)
                                                                                                                                var v96 bool = false
                                                                                                                                _ = v96
                                                                                                                            }
                                                                                                                            _ = v95
                                                                                                                        case "hello":
                                                                                                                            switch fi22 {
                                                                                                                            case 4294967295:
                                                                                                                                switch v15 {
                                                                                                                                case -247:
                                                                                                                                    v20 = (fi24 | (v15 ^ v15))
                                                                                                                                case 32768:
                                                                                                                                    v97 := GenS6{F0: 2}
                                                                                                                                    v6 = append(v6, (v20 * fi22))
                                                                                                                                    _ = v97
                                                                                                                                default:
                                                                                                                                    if true {
                                                                                                                                        v98 := GenS6{F0: uint16(v47)}
                                                                                                                                        var v99 int16 = 42
                                                                                                                                        _ = v98
                                                                                                                                        _ = v99
                                                                                                                                    } else {
                                                                                                                                        fmt.Println(v7, v7, fi22)
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            case -9223372036854775808:
                                                                                                                                v100 := GenS12{F0: ((true || false) || false), F1: v20, F2: 42}
                                                                                                                                _ = v100
                                                                                                                            default:
                                                                                                                                fmt.Println(v7, fi22, v48)
                                                                                                                            }
                                                                                                                        }
                                                                                                                        var v101 int16 = (v4 ^ v3)
                                                                                                                        switch v46 {
                                                                                                                        case "":
                                                                                                                            v102 := GenS5{F0: 1}
                                                                                                                            _ = v102
                                                                                                                        }
                                                                                                                        _ = v101
                                                                                                                    }
                                                                                                                    var v103 uint32 = 42
                                                                                                                    _ = v103
                                                                                                                } else {
                                                                                                                    if true {
                                                                                                                        if ((((!false) || (835 == 152)) && ((false || (128 == (v20 - 128))) || (!false))) && ((-(v20)) == (fi22 << ((137) & 63)))) {
                                                                                                                            var v104 uint8 = 0
                                                                                                                            _ = v104
                                                                                                                        }
                                                                                                                        var v105 int32 = -1
                                                                                                                        _ = v105
                                                                                                                    }
                                                                                                                    switch (fi24 % ((32767) | 1)) {
                                                                                                                    case 0:
                                                                                                                        v48 = float32((v1 - v1))
                                                                                                                    case 1:
                                                                                                                        if false {
                                                                                                                            v6 = append(v6, v20)
                                                                                                                            v106 := GenS15{F0: v20}
                                                                                                                            _ = v106
                                                                                                                        }
                                                                                                                    }
                                                                                                                    fmt.Println(v4, v1, v4)
                                                                                                                }
                                                                                                            }
                                                                                                            var v107 float32 = ((3.14159 + 3.14159) + 0.1)
                                                                                                            var v108 int32 = -289
                                                                                                            _ = v107
                                                                                                            _ = v108
                                                                                                        }
                                                                                                        _ = v89
                                                                                                    }
                                                                                                    _ = v88
                                                                                                }
                                                                                                var v109 uint64 = 4294967295
                                                                                                _ = v109
                                                                                            }
                                                                                            var v110 int64 = 1
                                                                                            if (0 == 1) {
                                                                                                var v111 uint32 = 2147483648
                                                                                                if false {
                                                                                                    v112 := GenS1{F0: v111, F1: 42, F2: 12}
                                                                                                    v113 := GenS16{F0: (^(v110))}
                                                                                                    switch fi24 {
                                                                                                    case 506:
                                                                                                        v114 := GenS17{F0: -1, F1: (^(v30))}
                                                                                                        _ = v114
                                                                                                    case -274:
                                                                                                        switch int(v30) {
                                                                                                        case -9223372036854775808:
                                                                                                            v3 -= v3
                                                                                                        case 4294967295:
                                                                                                            v6 = append(v6, v20)
                                                                                                            var v115 int16 = (v3 + 1)
                                                                                                            if false {
                                                                                                                v116 := GenS12{F0: false, F1: (^(fi24)), F2: 4294967295}
                                                                                                                fmt.Println(v110, v115, v9)
                                                                                                                if (1e10 >= v1) {
                                                                                                                    var v117 float64 = v2
                                                                                                                    v7 = ("" + v9)
                                                                                                                    _ = v117
                                                                                                                }
                                                                                                                v115 = (v3 - (v4 & v4))
                                                                                                                var v118 int32 = 229
                                                                                                                var v119 float64 = (v2 / (v1 + v1))
                                                                                                                _ = v116
                                                                                                                _ = v118
                                                                                                                _ = v119
                                                                                                            }
                                                                                                            _ = v115
                                                                                                        default:
                                                                                                            v120 := GenS4{F0: "", F1: (v111 / ((176) | 1))}
                                                                                                            _ = v120
                                                                                                        }
                                                                                                    case 2147483647:
                                                                                                        var v121 int64 = (v110 * v110)
                                                                                                        _ = v121
                                                                                                    }
                                                                                                    _ = v112
                                                                                                    _ = v113
                                                                                                }
                                                                                                _ = v111
                                                                                            } else {
                                                                                                var v122 uint64 = v30
                                                                                                var v123 uint = 18446744073709551615
                                                                                                delete(v5, int(v47))
                                                                                                _ = v122
                                                                                                _ = v123
                                                                                            }
                                                                                            _ = v110
                                                                                        } else {
                                                                                            var v124 int8 = 100
                                                                                            if ((v124 / ((v124) | 1)) >= (v124 + (v124 | v124))) {
                                                                                                var v125 uint32 = 4294967295
                                                                                                var v126 float64 = -2.25
                                                                                                v127 := GenS18{F0: -32768}
                                                                                                var v128 int8 = (^(v124))
                                                                                                switch ((v7 + "\n\t") + v33) {
                                                                                                case "":
                                                                                                    v129 := GenS19{F0: 9223372036854775807}
                                                                                                    switch (v46 + v9) {
                                                                                                    case "foo bar":
                                                                                                        fmt.Println(v1, v15, v126)
                                                                                                    default:
                                                                                                        if ("foo bar" != v9) {
                                                                                                            v21 = v48
                                                                                                            var v130 int16 = v3
                                                                                                            v5[(v20 * 0)] = (v20 % ((fi22) | 1))
                                                                                                            _ = v130
                                                                                                        } else {
                                                                                                            v131 := GenS20{F0: v125, F1: false, F2: -802}
                                                                                                            _ = v131
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v129
                                                                                                case "\n\t":
                                                                                                    if (!false) {
                                                                                                        if true {
                                                                                                            v125 += (v125 - 65535)
                                                                                                            if (155 != 127) {
                                                                                                                fmt.Println(v4, v125, v128)
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                case "foo bar":
                                                                                                    v132 := GenS1{F0: v125, F1: v124, F2: 1}
                                                                                                    _ = v132
                                                                                                default:
                                                                                                    if (!(0 == -2147483648)) {
                                                                                                        var v133 int64 = 1
                                                                                                        var v134 int64 = v133
                                                                                                        var v135 bool = true
                                                                                                        _ = v133
                                                                                                        _ = v134
                                                                                                        _ = v135
                                                                                                    } else {
                                                                                                        v15 = (-(v15))
                                                                                                    }
                                                                                                    if true {
                                                                                                        var v136 uint64 = 601
                                                                                                        var v137 uint = (v47 >> ((v47) & 63))
                                                                                                        _ = v136
                                                                                                        _ = v137
                                                                                                    } else {
                                                                                                        switch (v20 / ((fi24) | 1)) {
                                                                                                        case -9223372036854775808:
                                                                                                            v5[fi24] = v15
                                                                                                        case 2147483647:
                                                                                                            if true {
                                                                                                                switch ((((("a" + v46) + ("\n\t" + "\n\t")) + v9) + (("hello" + ((v46 + ((v9 + v33) + v9)) + v7)) + (v46 + (v7 + (v9 + (v46 + (v7 + (v7 + (v9 + ""))))))))) + "a") {
                                                                                                                case "hello":
                                                                                                                    v138 := v5[v15]
                                                                                                                    _ = v138
                                                                                                                case "foo bar":
                                                                                                                    var v139 uint32 = (v125 | (v125 << (((v47 << ((v19) & 63))) & 63)))
                                                                                                                    _ = v139
                                                                                                                case "\n\t":
                                                                                                                    switch v9 {
                                                                                                                    case "foo bar":
                                                                                                                        v140 := GenS10{F0: 42, F1: v125, F2: 2147483647}
                                                                                                                        _ = v140
                                                                                                                    case "hello":
                                                                                                                        fmt.Println(fi22, v125, v125)
                                                                                                                    }
                                                                                                                }
                                                                                                                var v141 int16 = v4
                                                                                                                _ = v141
                                                                                                            } else {
                                                                                                                var v142 int32 = 2147483647
                                                                                                                v30 = 18446744073709551615
                                                                                                                _ = v142
                                                                                                            }
                                                                                                        case 4294967295:
                                                                                                            var v143 float64 = v1
                                                                                                            _ = v143
                                                                                                        default:
                                                                                                            fmt.Println(v48)
                                                                                                        }
                                                                                                        switch (v15 ^ -1) {
                                                                                                        case -1:
                                                                                                            if (v4 >= (v4 + v4)) {
                                                                                                                var v144 float32 = float32(0.0)
                                                                                                                if (42 != 42) {
                                                                                                                    v46 = v46
                                                                                                                    v145 := GenS21{F0: 4294967295, F1: 9223372036854775807, F2: (^(v30))}
                                                                                                                    _ = v145
                                                                                                                } else {
                                                                                                                    _ = len(v6)
                                                                                                                    if false {
                                                                                                                        v144 = (v21 + float32(-2.25))
                                                                                                                    }
                                                                                                                    if true {
                                                                                                                        if len(v6) >= 2 {
                                                                                                                            v6 = v6[1:len(v6)]
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                if (!(!(!(!true)))) {
                                                                                                                    var v146 int64 = 0
                                                                                                                    var v147 uint32 = v125
                                                                                                                    _ = v146
                                                                                                                    _ = v147
                                                                                                                } else {
                                                                                                                    v5[fi24] = (^(fi22))
                                                                                                                    if true {
                                                                                                                        var v148 bool = true
                                                                                                                        if (65535 <= uint32(v47)) {
                                                                                                                            var v149 float32 = (1.5 * 1.5)
                                                                                                                            _ = v149
                                                                                                                        }
                                                                                                                        if len(v6) > 0 {
                                                                                                                            v6[0] = 32768
                                                                                                                        }
                                                                                                                        v4 = (v4 % (((v3 & v3)) | 1))
                                                                                                                        _ = v148
                                                                                                                    }
                                                                                                                    if true {
                                                                                                                        v125 = v125
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v144
                                                                                                            }
                                                                                                        default:
                                                                                                            if true {
                                                                                                                fmt.Println(fi24, v20, v125)
                                                                                                                v150 := v5[(v15 % ((int(v124)) | 1))]
                                                                                                                _ = len(v5)
                                                                                                                if (false && (165 >= v4)) {
                                                                                                                    var v151 float32 = 1e10
                                                                                                                    if false {
                                                                                                                        var v152 float64 = 0.1
                                                                                                                        _ = v152
                                                                                                                    } else {
                                                                                                                        var v153 float64 = float64(v151)
                                                                                                                        var v154 bool = (1 < v3)
                                                                                                                        fmt.Println(v30, v2, v150)
                                                                                                                        _ = v153
                                                                                                                        _ = v154
                                                                                                                    }
                                                                                                                    _ = v151
                                                                                                                }
                                                                                                                v155 := GenS12{F0: false, F1: v150, F2: v125}
                                                                                                                var v156 int8 = v128
                                                                                                                v157 := GenS4{F0: (v46 + (((v9 + v9) + ((("a" + ("hello" + "a")) + (v33 + (v9 + v7))) + (v9 + v9))) + ((("hello" + "") + v9) + "foo bar"))), F1: v125}
                                                                                                                _ = v150
                                                                                                                _ = v155
                                                                                                                _ = v156
                                                                                                                _ = v157
                                                                                                            }
                                                                                                        }
                                                                                                        fmt.Println(v3)
                                                                                                    }
                                                                                                    var v158 float32 = v21
                                                                                                    var v159 int = (^(v15))
                                                                                                    _ = v158
                                                                                                    _ = v159
                                                                                                }
                                                                                                v6 = append(v6, fi22)
                                                                                                _ = v125
                                                                                                _ = v126
                                                                                                _ = v127
                                                                                                _ = v128
                                                                                            } else {
                                                                                                if ((v1 > v2) && false) {
                                                                                                    var v160 int = -538
                                                                                                    delete(v5, 836)
                                                                                                    _ = v160
                                                                                                }
                                                                                                var v161 uint8 = 42
                                                                                                v21 = 0.0
                                                                                                v9 = v46
                                                                                                switch v7 {
                                                                                                case "\n\t":
                                                                                                    if true {
                                                                                                        var v162 uint32 = 143
                                                                                                        _ = v162
                                                                                                    } else {
                                                                                                        v21 -= ((v48 + 0.1) + (v21 / ((((v21 - 3.14159) - v21) - v21) / v48)))
                                                                                                        v5[(fi24 >> ((v47) & 63))] = fi22
                                                                                                        v15 = (v15 * (fi22 + (v15 & (fi22 ^ 128))))
                                                                                                    }
                                                                                                case "":
                                                                                                    var v163 uint32 = 2
                                                                                                    _ = v163
                                                                                                default:
                                                                                                    switch ((v9 + v33) + (("foo bar" + v46) + "foo bar")) {
                                                                                                    case "\n\t":
                                                                                                        var v164 bool = (!false)
                                                                                                        _ = v164
                                                                                                    case "":
                                                                                                        v21 += -2.25
                                                                                                    case "foo bar":
                                                                                                        v21 = float32((v2 + ((-2.25 / (((1e10 * v2) * (0.1 * 1.5)) / 1e10)) * 0.0)))
                                                                                                    default:
                                                                                                        var v165 float32 = 0.0
                                                                                                        _ = v165
                                                                                                    }
                                                                                                    v166 := GenS10{F0: v124, F1: 2147483648, F2: 851}
                                                                                                    v48 = 3.14159
                                                                                                    _ = v166
                                                                                                }
                                                                                                _ = v161
                                                                                            }
                                                                                            v167 := GenS5{F0: 42}
                                                                                            _ = v124
                                                                                            _ = v167
                                                                                        }
                                                                                        var v168 int = fi22
                                                                                        _ = v168
                                                                                    } else {
                                                                                        if len(v6) >= 2 {
                                                                                            v6 = v6[1:len(v6)]
                                                                                        }
                                                                                    }
                                                                                    _ = v46
                                                                                    _ = v47
                                                                                    _ = v48
                                                                                default:
                                                                                    fmt.Println(v33, v15, v19)
                                                                                }
                                                                                switch v9 {
                                                                                case "\n\t":
                                                                                    var v169 int16 = (v4 - v4)
                                                                                    _ = v169
                                                                                case "hello":
                                                                                    v5[(v15 % (((fi24 | (v15 - (fi24 - (fi24 & (fi24 ^ 127)))))) | 1))] = (fi24 % (((fi22 | (-(fi22)))) | 1))
                                                                                }
                                                                            }
                                                                        }
                                                                        if (false || false) {
                                                                            v170 := GenS20{F0: 2147483648, F1: (true && true), F2: (fi24 % (((v15 % ((-650) | 1))) | 1))}
                                                                            v7 = (v9 + "hello")
                                                                            _ = v170
                                                                        } else {
                                                                            v2 -= 1.5
                                                                            v15 += 42
                                                                            _ = len(v6)
                                                                        }
                                                                        v171 := GenS17{F0: fi22, F1: (^(v30))}
                                                                        v9 = ""
                                                                        _ = len(v5)
                                                                        var v172 uint8 = 128
                                                                        v5[v20] = (fi24 << ((42) & 63))
                                                                        _ = v171
                                                                        _ = v172
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                            case -764:
                                                var v173 int32 = 0
                                                _ = v173
                                            default:
                                                var v174 int16 = (v4 * 1000)
                                                _ = v174
                                            }
                                            v30 -= (^(v30))
                                            _ = v30
                                            _ = v31
                                            _ = v32
                                            _ = v33
                                        }
                                    case 0:
                                        v7 = v7
                                    case -419:
                                        v7 = v7
                                    }
                                }
                                fmt.Println(v4, v4)
                            }
                        } else {
                            v175 := GenS22{F0: v7, F1: fi24}
                            _ = v175
                        }
                        _ = fi24
                    }
                case "hello":
                    if true {
                        v5[(fi22 * (v15 & (v15 << ((v19) & 63))))] = (v20 % ((-2147483648) | 1))
                        _ = len(v5)
                        v176 := GenS4{F0: "\n\t", F1: 4294967295}
                        _ = v176
                    }
                }
                switch (v15 + v15) {
                case 451:
                    var v177 int = v20
                    _ = v177
                case 617:
                    v15 = int(v4)
                case 820:
                    var v178 int64 = -228
                    _ = v178
                }
                for fi179 := 0; fi179 < 10; fi179++ {
                    switch 737 {
                    case -1:
                        var v180 float32 = ((0.0 - v21) / (v21 - v21))
                        _ = v180
                    default:
                        if (-68 == 127) {
                            var v181 int64 = 0
                            v181 += v181
                            v7 = ("foo bar" + "a")
                            _ = v181
                        } else {
                            _ = len(v5)
                            if (true && ((-574 < -9223372036854775808) && (954 < 0))) {
                                switch 9223372036854775807 {
                                case 2147483647:
                                    var v182 int8 = 52
                                    _ = v182
                                }
                            }
                            if ((false && (true && true)) && (false && ("hello" == v9))) {
                                v183 := GenS23{F0: (-887 < -1)}
                                _ = v183
                            } else {
                                v2 = float64(v21)
                                if true {
                                    v184 := GenS7{F0: -1, F1: false}
                                    v185 := GenS2{F0: 1}
                                    _ = v184
                                    _ = v185
                                } else {
                                    var v186 uint32 = 1
                                    _ = v186
                                }
                            }
                            v187 := GenS4{F0: v9, F1: 0}
                            _ = v187
                        }
                    }
                    var v188 uint16 = 1
                    v4 = v4
                    _ = fi179
                    _ = v188
                }
            }
            if (((1.5 + (float32(v1) + v21)) != v21) || false) {
                for fi189 := 0; fi189 < 6; fi189++ {
                    fmt.Println(v3, v2, v9)
                    _ = fi189
                }
            }
            fmt.Println(v3)
            _ = fi22
        }
        var v190 int32 = -1
        _ = v190
        _ = v21
    }
    switch "" {
    case "\n\t":
        var v191 uint8 = 127
        _ = v191
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
