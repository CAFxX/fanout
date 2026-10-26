package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8; F1 int32; F2 int8 }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 uint16; F1 uint; F2 bool }
type GenS8 struct { F0 int8; F1 uint }
type GenS9 struct { F0 uint32; F1 uint8 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint; F1 uint32 }
type GenS12 struct { F0 int8; F1 string; F2 int32 }
type GenS13 struct { F0 int64; F1 uint64; F2 string }
type GenS14 struct { F0 int; F1 int32; F2 uint8 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 uint16; F1 uint }
type GenS17 struct { F0 int8; F1 uint32 }
type GenS18 struct { F0 int64; F1 int8; F2 uint16 }
type GenS19 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS20 struct { F0 bool }
type GenS21 struct { F0 bool }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 int }
type GenS24 struct { F0 int32 }
type GenS25 struct { F0 uint; F1 uint8 }
type GenS26 struct { F0 uint16; F1 int32 }
type GenS27 struct { F0 int; F1 int; F2 uint16 }
type GenS28 struct { F0 uint32 }
type GenS29 struct { F0 uint64; F1 int; F2 int32 }
type GenS30 struct { F0 int32; F1 bool; F2 uint16 }
type GenS31 struct { F0 int; F1 uint16; F2 int }
type GenS32 struct { F0 uint32; F1 uint16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]bool)
    var v3 string = ((((("hello" + "") + "\n\t") + ("" + "a")) + "") + ("foo bar" + ("hello" + (("a" + ("hello" + ("foo bar" + "hello"))) + ("" + "\n\t")))))
    v4 := []int{-2147483648, -1}
    var v5 uint8 = 206
    v6 := []uint32{791, 2147483647, 65535, 1}
    var v7 int = 127
    for fi8 := 0; fi8 < 18; fi8++ {
        if false {
            var v9 uint8 = v5
            if ((v5 >> ((1) & 63)) < (v9 % ((v9) | 1))) {
                v4 = append(v4, fi8)
                if (2 != 506) {
                    v9 = v5
                    if (110 < 1) {
                        if len(v6) > 0 {
                            v6[0] = 42
                        }
                        _ = len(v1)
                    } else {
                        var v10 int8 = -128
                        v9 += v5
                        _ = v10
                    }
                }
            } else {
                v3 = "foo bar"
            }
            _ = v9
        }
        var v11 int = v7
        switch (v7 * (fi8 >> ((18446744073709551615) & 63))) {
        case -379:
            v12 := GenS1{F0: v3, F1: true, F2: 676}
            _ = v12
        default:
            var v13 uint16 = 893
            _ = v13
        }
        _ = fi8
        _ = v11
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        switch (^(fi14)) {
        case 0:
            var v15 int32 = 214
            _ = v15
        case 676:
            for fi16 := 0; fi16 < 20; fi16++ {
                v3 = (v3 + (v3 + v3))
                var v17 uint64 = 42
                v5 = (v5 >> ((18446744073709551615) & 63))
                _ = fi16
                _ = v17
            }
        case 42:
            if true {
                v7 = int(v5)
                for fi18 := 0; fi18 < 10; fi18++ {
                    if true {
                        if false {
                            v19 := GenS2{F0: 206}
                            var v20 int16 = 32767
                            _ = v19
                            _ = v20
                        } else {
                            v21 := GenS3{F0: 4294967295, F1: (v5 << ((511) & 63))}
                            var v22 int64 = -386
                            switch (v3 + v3) {
                            case "\n\t":
                                switch ("\n\t" + v3) {
                                case "foo bar":
                                    var v23 bool = false
                                    _ = v23
                                case "":
                                    switch fi14 {
                                    case -1:
                                        if len(v4) > 0 {
                                            v4[0] = 14
                                        }
                                    case 9223372036854775807:
                                        v1[(v7 - 128)] = 576
                                    }
                                case "hello":
                                    var v24 int64 = v22
                                    _ = v24
                                }
                            case "a":
                                if (!(true && ((0 > 1) || false))) {
                                    fmt.Println(v7, fi18, v3)
                                    fmt.Println(v22, v5)
                                    switch ((v3 + (v3 + v3)) + v3) {
                                    case "":
                                        switch "a" {
                                        case "hello":
                                            var v25 int32 = 42
                                            fmt.Println(v25, fi18, fi14)
                                            v25 -= (v25 % (((v25 >> ((457) & 63))) | 1))
                                            _ = v25
                                        default:
                                            var v26 uint8 = 69
                                            _ = v26
                                        }
                                    case "\n\t":
                                        var v27 uint = 279
                                        _ = v27
                                    default:
                                        v7 += -2147483648
                                    }
                                } else {
                                    fmt.Println(fi18, fi14, v3)
                                    var v28 int64 = (^(v22))
                                    _ = v28
                                }
                            case "hello":
                                v29 := GenS3{F0: 923, F1: (v5 << ((353) & 63))}
                                _ = v29
                            }
                            _ = v21
                            _ = v22
                        }
                        v30 := GenS2{F0: -2147483648}
                        _ = v30
                    }
                    switch "a" {
                    case "a":
                        if (!false) {
                            var v31 int = -1
                            if ((4294967295 >= 0) || ((0.1 * 1.5) != -2.25)) {
                                switch fi14 {
                                case 4294967295:
                                    switch v7 {
                                    case 4294967295:
                                        var v32 float64 = 0.1
                                        _ = v32
                                    }
                                case 377:
                                    if ((v5 - (v5 / ((v5) | 1))) >= (v5 | 0)) {
                                        switch v31 {
                                        case 716:
                                            if (9223372036854775807 > 0) {
                                                delete(v2, v3)
                                                v3 = (v3 + "a")
                                                v5 += 128
                                            } else {
                                                v3 = v3
                                                if (false && false) {
                                                    if (0 != 204) {
                                                        if true {
                                                            v33 := GenS4{F0: 0}
                                                            _ = v33
                                                        } else {
                                                            v31 += 32768
                                                            v34 := GenS5{F0: 42, F1: -2147483648, F2: -128}
                                                            switch 0 {
                                                            case 128:
                                                                if (32767 >= 32767) {
                                                                    fmt.Println(v31, v31)
                                                                    var v35 int = -766
                                                                    var v36 uint32 = 65535
                                                                    switch (v3 + v3) {
                                                                    case "foo bar":
                                                                        v36 -= 0
                                                                    default:
                                                                        v31 += fi18
                                                                    }
                                                                    switch v35 {
                                                                    case 4294967295:
                                                                        _ = len(v2)
                                                                    case -220:
                                                                        var v37 uint64 = 4294967295
                                                                        _ = v37
                                                                    case 127:
                                                                        var v38 float32 = (1e10 - 0.1)
                                                                        _ = v38
                                                                    }
                                                                    v39 := GenS4{F0: 0}
                                                                    _ = v35
                                                                    _ = v36
                                                                    _ = v39
                                                                } else {
                                                                    v40 := GenS4{F0: -1}
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                    fmt.Println(v7, v5)
                                                                    var v41 uint = 42
                                                                    if true {
                                                                        v42 := GenS6{F0: fi18, F1: v41}
                                                                        if (false || true) {
                                                                            v43 := GenS7{F0: 42, F1: (v41 * 2), F2: (100 <= 127)}
                                                                            _ = v43
                                                                        } else {
                                                                            v4 = append(v4, (v31 - v31))
                                                                            fmt.Println(v41)
                                                                            v41 -= 2
                                                                            _ = len(v4)
                                                                        }
                                                                        var v44 uint = v41
                                                                        switch v3 {
                                                                        case "":
                                                                            v1[(v31 | v31)] = 42
                                                                        case "hello":
                                                                            v45 := GenS5{F0: 100, F1: 766, F2: 127}
                                                                            _ = v45
                                                                        default:
                                                                            v46 := GenS1{F0: v3, F1: (false && (!(42 <= 255))), F2: v41}
                                                                            _ = v46
                                                                        }
                                                                        _ = v42
                                                                        _ = v44
                                                                    } else {
                                                                        fmt.Println(fi14, v41, fi14)
                                                                        var v47 float64 = float64(0.1)
                                                                        v48 := GenS8{F0: 100, F1: v41}
                                                                        _ = v47
                                                                        _ = v48
                                                                    }
                                                                    _ = v40
                                                                    _ = v41
                                                                }
                                                            case 460:
                                                                if false {
                                                                    v49 := GenS9{F0: 42, F1: 128}
                                                                    _ = v49
                                                                }
                                                            }
                                                            if false {
                                                                fmt.Println(v7)
                                                                v5 -= v5
                                                                fmt.Println(v31, fi18, v5)
                                                            } else {
                                                                v50 := GenS10{F0: v5}
                                                                _ = v50
                                                            }
                                                            _ = v34
                                                        }
                                                    } else {
                                                        if (1000 > -870) {
                                                            v7 = fi14
                                                        } else {
                                                            var v51 int16 = 32767
                                                            _ = v51
                                                        }
                                                    }
                                                    var v52 string = "foo bar"
                                                    v3 = v52
                                                    _ = v52
                                                }
                                            }
                                        case 42:
                                            switch "hello" {
                                            case "foo bar":
                                                var v53 uint16 = 1
                                                _ = v53
                                            case "hello":
                                                switch v31 {
                                                case -1:
                                                    var v54 uint32 = 4294967295
                                                    _ = v54
                                                case -86:
                                                    switch (fi18 + v7) {
                                                    case 9223372036854775807:
                                                        v3 = "\n\t"
                                                    case 2147483647:
                                                        v55 := GenS11{F0: 2, F1: 0}
                                                        switch v3 {
                                                        case "":
                                                            if (1 >= 9223372036854775807) {
                                                                if len(v6) > 0 {
                                                                    v6[0] = 158
                                                                }
                                                                switch v7 {
                                                                case 9223372036854775807:
                                                                    var v56 int8 = 2
                                                                    _ = v56
                                                                case 42:
                                                                    switch -1 {
                                                                    case 4294967295:
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            v31 -= 9223372036854775807
                                                                        default:
                                                                            _ = len(v6)
                                                                        }
                                                                    case -2147483648:
                                                                        if ("foo bar" <= ("\n\t" + (v3 + v3))) {
                                                                            v5 += v5
                                                                            if ((!true) && ((-2.25 / ((0.1 + 1.5) + 0.0)) == 0.0)) {
                                                                                var v57 uint16 = 136
                                                                                switch (fi14 * (v7 - 4294967295)) {
                                                                                case -1:
                                                                                    delete(v2, v3)
                                                                                case 4294967295:
                                                                                    v57 = (v57 % ((243) | 1))
                                                                                case -210:
                                                                                    v57 = (v57 / ((352) | 1))
                                                                                }
                                                                                _ = v57
                                                                            } else {
                                                                                v58 := GenS12{F0: 1, F1: v3, F2: 268}
                                                                                fmt.Println(fi14)
                                                                                fmt.Println(v7)
                                                                                _ = v58
                                                                            }
                                                                            fmt.Println(fi18, v7)
                                                                            switch (v31 % (((v7 - v7)) | 1)) {
                                                                            case 128:
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = (fi14 << ((1) & 63))
                                                                                }
                                                                            case 32768:
                                                                                v31 += fi18
                                                                                v3 = "foo bar"
                                                                                var v59 string = (v3 + v3)
                                                                                _ = v59
                                                                            case -1:
                                                                                delete(v2, v3)
                                                                            }
                                                                            var v60 float64 = 0.1
                                                                            _ = v60
                                                                        }
                                                                    default:
                                                                        if (0 < 32767) {
                                                                            v4 = append(v4, v31)
                                                                            var v61 float32 = ((1.5 / 3.14159) - 3.14159)
                                                                            _ = v61
                                                                        }
                                                                        v7 += v7
                                                                        _ = len(v6)
                                                                    }
                                                                default:
                                                                    switch 9223372036854775807 {
                                                                    case 128:
                                                                        v62 := GenS4{F0: -128}
                                                                        switch 0 {
                                                                        case 4294967295:
                                                                            v3 = ("" + v3)
                                                                        case 32767:
                                                                            if len(v4) > 0 {
                                                                                v4[0] = 32767
                                                                            }
                                                                        case -2147483648:
                                                                            var v63 float64 = 0.1
                                                                            _ = v63
                                                                        }
                                                                        _ = v62
                                                                    case -2147483648:
                                                                        if ((4294967295 < 42) || (v5 >= v5)) {
                                                                            v64 := GenS10{F0: 0}
                                                                            _ = len(v2)
                                                                            var v65 int64 = -9223372036854775808
                                                                            _ = v64
                                                                            _ = v65
                                                                        } else {
                                                                            fmt.Println(v31, v31, v31)
                                                                            v66 := GenS13{F0: 42, F1: 18446744073709551615, F2: v3}
                                                                            var v67 uint = 65535
                                                                            var v68 int16 = 42
                                                                            _ = v66
                                                                            _ = v67
                                                                            _ = v68
                                                                        }
                                                                    case 839:
                                                                        if false {
                                                                            v5 -= (^(v5))
                                                                            switch (^(v7)) {
                                                                            case -9223372036854775808:
                                                                                var v69 uint = 255
                                                                                _ = v69
                                                                            }
                                                                            if false {
                                                                                v5 -= v5
                                                                                var v70 int = 32768
                                                                                _ = v70
                                                                            }
                                                                        }
                                                                    default:
                                                                        v2[""] = true
                                                                    }
                                                                }
                                                                v3 = v3
                                                            } else {
                                                                switch "a" {
                                                                case "hello":
                                                                    var v71 float64 = 1.5
                                                                    _ = v71
                                                                case "a":
                                                                    v5 += v5
                                                                case "":
                                                                    v72 := GenS14{F0: fi18, F1: -1, F2: (v5 << ((4294967295) & 63))}
                                                                    _ = v72
                                                                }
                                                                if (!(true && false)) {
                                                                    if (1 != 0) {
                                                                        v6 = append(v6, 2147483648)
                                                                        fmt.Println(fi18, v7, v3)
                                                                    }
                                                                    v3 = "foo bar"
                                                                    if false {
                                                                        var v73 float32 = 0.0
                                                                        var v74 uint = 38
                                                                        if (0.1 == (((1.5 / (1.5 - (-2.25 + 0.0))) + 1.5) / 1e10)) {
                                                                            var v75 float64 = (1.5 - float64(float32(-2.25)))
                                                                            fmt.Println(v73, v3)
                                                                            _ = v75
                                                                        }
                                                                        var v76 int16 = 1
                                                                        _ = v73
                                                                        _ = v74
                                                                        _ = v76
                                                                    }
                                                                } else {
                                                                    var v77 float64 = (-2.25 / 0.1)
                                                                    v78 := GenS15{F0: 722}
                                                                    var v79 int = 2147483647
                                                                    v31 += 32767
                                                                    if true {
                                                                        v5 = (v5 | v5)
                                                                    }
                                                                    _ = v77
                                                                    _ = v78
                                                                    _ = v79
                                                                }
                                                                delete(v2, (("a" + v3) + "a"))
                                                            }
                                                            var v80 uint32 = 1
                                                            var v81 float64 = 0.0
                                                            _ = v80
                                                            _ = v81
                                                        case "\n\t":
                                                            delete(v1, fi14)
                                                        }
                                                        switch v31 {
                                                        case 78:
                                                            v5 = v5
                                                        case 0:
                                                            if (-2.25 <= -2.25) {
                                                                v82 := GenS11{F0: 4294967295, F1: 4294967295}
                                                                switch fi14 {
                                                                case -9223372036854775808:
                                                                    switch ((v3 + (v3 + ("\n\t" + "hello"))) + "hello") {
                                                                    case "a":
                                                                        if (65535 > 4294967295) {
                                                                            var v83 uint16 = 65535
                                                                            _ = v83
                                                                        } else {
                                                                            if (1 == 828) {
                                                                                if false {
                                                                                    v5 = v5
                                                                                    var v84 float64 = 0.1
                                                                                    var v85 int = v31
                                                                                    _ = v84
                                                                                    _ = v85
                                                                                }
                                                                                var v86 int16 = 0
                                                                                _ = len(v2)
                                                                                _ = v86
                                                                            } else {
                                                                                var v87 uint16 = 0
                                                                                _ = v87
                                                                            }
                                                                            switch (("" + v3) + (((v3 + (v3 + ((v3 + v3) + (v3 + (v3 + "foo bar"))))) + (("hello" + (v3 + "hello")) + "hello")) + v3)) {
                                                                            case "foo bar":
                                                                                var v88 int64 = -1
                                                                                _ = v88
                                                                            case "hello":
                                                                                switch v7 {
                                                                                case 42:
                                                                                    var v89 int32 = int32(v31)
                                                                                    _ = v89
                                                                                case -9223372036854775808:
                                                                                    v31 -= v31
                                                                                case 128:
                                                                                    if ((true && (false || false)) || true) {
                                                                                        v5 = (^(v5))
                                                                                        v4 = append(v4, v7)
                                                                                        var v90 float64 = (3.14159 + (float64(1.5) / 1.0))
                                                                                        v3 = v3
                                                                                        switch v7 {
                                                                                        case -1:
                                                                                            switch -929 {
                                                                                            case -9223372036854775808:
                                                                                                v3 = v3
                                                                                            case 128:
                                                                                                fmt.Println(v31, v5)
                                                                                            case -10:
                                                                                                var v91 uint32 = 0
                                                                                                _ = v91
                                                                                            }
                                                                                        case 0:
                                                                                            v6 = append(v6, 2147483648)
                                                                                        default:
                                                                                            var v92 string = "a"
                                                                                            _ = v92
                                                                                        }
                                                                                        _ = v90
                                                                                    }
                                                                                default:
                                                                                    v93 := GenS16{F0: 965, F1: 159}
                                                                                    _ = v93
                                                                                }
                                                                            }
                                                                        }
                                                                    case "hello":
                                                                        v3 = "a"
                                                                        v5 -= (v5 | v5)
                                                                    case "":
                                                                        v5 = 2
                                                                    }
                                                                }
                                                                _ = v82
                                                            } else {
                                                                v7 = 32767
                                                                switch (fi18 - fi14) {
                                                                case 988:
                                                                    v31 += (^(v31))
                                                                case -727:
                                                                    v5 = 42
                                                                case -702:
                                                                    var v94 uint32 = 224
                                                                    _ = v94
                                                                default:
                                                                    var v95 int32 = 368
                                                                    _ = v95
                                                                }
                                                                switch -1 {
                                                                case -708:
                                                                    v96 := GenS3{F0: 0, F1: v5}
                                                                    _ = v96
                                                                case -1:
                                                                    v97 := GenS13{F0: 1, F1: 0, F2: ""}
                                                                    _ = v97
                                                                case 874:
                                                                    var v98 uint16 = 1
                                                                    _ = v98
                                                                }
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    var v99 float32 = 1.5
                                                                    _ = v99
                                                                default:
                                                                    if false {
                                                                        fmt.Println(fi14)
                                                                        v100 := GenS17{F0: 100, F1: 1}
                                                                        v101 := GenS17{F0: 42, F1: 65535}
                                                                        _ = v100
                                                                        _ = v101
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = v55
                                                    default:
                                                        if false {
                                                            if (-128 > -128) {
                                                                var v102 int16 = 32767
                                                                var v103 int64 = -1
                                                                var v104 int8 = 100
                                                                _ = v102
                                                                _ = v103
                                                                _ = v104
                                                            } else {
                                                                if (v5 < (^(v5))) {
                                                                    var v105 int8 = 127
                                                                    _ = v105
                                                                }
                                                                v31 -= 32768
                                                                v7 = v31
                                                            }
                                                            if (true && true) {
                                                                v106 := GenS2{F0: -942}
                                                                fmt.Println(v31, v7)
                                                                var v107 int16 = 42
                                                                var v108 int64 = -9223372036854775808
                                                                _ = v106
                                                                _ = v107
                                                                _ = v108
                                                            }
                                                            switch v3 {
                                                            case "":
                                                                v3 = v3
                                                            case "hello":
                                                                if (839 != 255) {
                                                                    var v109 uint8 = v5
                                                                    v31 = (fi18 % ((32768) | 1))
                                                                    v110 := GenS2{F0: -1}
                                                                    _ = v109
                                                                    _ = v110
                                                                } else {
                                                                    v5 -= 130
                                                                    var v111 int64 = 42
                                                                    var v112 bool = (0 != 15)
                                                                    _ = v111
                                                                    _ = v112
                                                                }
                                                            case "a":
                                                                switch "a" {
                                                                case "":
                                                                    if (true || ((-32768 != 42) && false)) {
                                                                        v113 := GenS9{F0: 993, F1: 2}
                                                                        _ = v113
                                                                    }
                                                                case "foo bar":
                                                                    var v114 uint8 = (^(v5))
                                                                    _ = v114
                                                                case "\n\t":
                                                                    delete(v2, v3)
                                                                default:
                                                                    fmt.Println(v31, v3, fi14)
                                                                }
                                                            default:
                                                                if (219 == 4294967295) {
                                                                    fmt.Println(v5)
                                                                    switch 19 {
                                                                    case 4294967295:
                                                                        v4 = append(v4, (v31 | -2147483648))
                                                                    case -416:
                                                                        var v115 bool = true
                                                                        _ = v115
                                                                    }
                                                                    v116 := v1[v31]
                                                                    _ = v116
                                                                } else {
                                                                    if false {
                                                                        var v117 int = fi18
                                                                        if true {
                                                                            switch (v117 & fi18) {
                                                                            case -1:
                                                                                v118 := GenS11{F0: 347, F1: 2147483647}
                                                                                _ = v118
                                                                            }
                                                                            var v119 uint64 = 2
                                                                            v120 := GenS18{F0: 1, F1: 0, F2: 255}
                                                                            _ = v119
                                                                            _ = v120
                                                                        } else {
                                                                            v121 := GenS10{F0: (^(v5))}
                                                                            _ = v121
                                                                        }
                                                                        _ = v117
                                                                    }
                                                                }
                                                            }
                                                            v5 = v5
                                                        } else {
                                                            var v122 int = fi14
                                                            var v123 uint = 18446744073709551615
                                                            var v124 int8 = -1
                                                            _ = v122
                                                            _ = v123
                                                            _ = v124
                                                        }
                                                    }
                                                }
                                            case "a":
                                                if true {
                                                    v3 = v3
                                                }
                                            default:
                                                var v125 uint16 = 42
                                                _ = v125
                                            }
                                        }
                                    } else {
                                        var v126 int = v7
                                        _ = v126
                                    }
                                    fmt.Println(v7, v31, fi18)
                                    v5 += 1
                                case 440:
                                    v31 = -9223372036854775808
                                default:
                                    fmt.Println(v7, v5)
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v127 := GenS19{F0: (v3 <= v3), F1: 18446744073709551615, F2: uint32(v5)}
                                _ = v127
                            } else {
                                var v128 uint64 = 2
                                v129 := GenS20{F0: (uint16(v5) >= 65535)}
                                var v130 int8 = 0
                                var v131 uint8 = v5
                                switch (fi14 << ((uint(v131)) & 63)) {
                                case -9223372036854775808:
                                    fmt.Println(fi18)
                                }
                                v132 := GenS4{F0: v130}
                                var v133 uint32 = 2147483648
                                switch (v7 & (v7 / ((v7) | 1))) {
                                case -2147483648:
                                    _ = len(v4)
                                case 0:
                                    v133 += 65535
                                case 4294967295:
                                    var v134 bool = true
                                    if false {
                                        if (!(((1 < 42) && v134) && false)) {
                                            if v134 {
                                                var v135 uint = 1
                                                v135 = (v135 << (((^(v135))) & 63))
                                                _ = v135
                                            } else {
                                                _ = len(v6)
                                                _ = len(v4)
                                                var v136 uint = 795
                                                _ = v136
                                            }
                                        } else {
                                            var v137 uint = 65535
                                            if ((v130 >> ((4294967295) & 63)) >= v130) {
                                                switch v7 {
                                                case -9223372036854775808:
                                                    fmt.Println(v131, v137)
                                                }
                                            }
                                            fmt.Println(v134, v134, v137)
                                            v1[fi14] = v133
                                            v138 := GenS11{F0: 13, F1: (v133 << (((v137 * (^(v137)))) & 63))}
                                            _ = v137
                                            _ = v138
                                        }
                                    }
                                    v139 := GenS17{F0: v130, F1: 1}
                                    _ = v134
                                    _ = v139
                                }
                                v140 := GenS21{F0: false}
                                var v141 float64 = (-2.25 + 1e10)
                                delete(v2, (("\n\t" + "a") + ("" + "foo bar")))
                                _ = v128
                                _ = v129
                                _ = v130
                                _ = v131
                                _ = v132
                                _ = v133
                                _ = v140
                                _ = v141
                            }
                            _ = v31
                        }
                    case "hello":
                        switch fi18 {
                        case 436:
                            v142 := GenS22{F0: false}
                            _ = v142
                        case 2147483647:
                            v143 := GenS5{F0: -128, F1: 1, F2: 96}
                            _ = v143
                        case 9223372036854775807:
                            v144 := GenS23{F0: fi18}
                            _ = v144
                        default:
                            v145 := GenS24{F0: 0}
                            _ = v145
                        }
                    }
                    _ = fi18
                }
                var v146 int = (fi14 - v7)
                _ = v146
            } else {
                if true {
                    switch (("a" + v3) + "foo bar") {
                    case "foo bar":
                        v147 := GenS25{F0: 1, F1: 42}
                        _ = v147
                    default:
                        if (false || ((((-2.25 * 1e10) / (1.5 + 1.5)) - 3.14159) > float64((3.14159 / 1.5)))) {
                            v3 = (v3 + "a")
                            v7 = fi14
                        } else {
                            if (18446744073709551615 < 0) {
                                v3 = v3
                            }
                            v5 -= v5
                        }
                    }
                }
                for fi148 := 0; fi148 < 7; fi148++ {
                    if (((1000 != 417) || false) || (42 == -1)) {
                        if len(v6) > 0 {
                            v6[0] = 2147483647
                        }
                        v5 -= 3
                        v7 += 32768
                    }
                    switch -1 {
                    case 9223372036854775807:
                        v3 = v3
                    default:
                        fmt.Println(fi148)
                    }
                    var v149 uint16 = 42
                    _ = fi148
                    _ = v149
                }
            }
        default:
            var v150 bool = (!(-941 > 95))
            _ = v150
        }
        switch v3 {
        case "a":
            v151 := GenS26{F0: 255, F1: 2147483647}
            _ = v151
        case "foo bar":
            switch (v3 + "hello") {
            case "a":
                v152 := GenS27{F0: (v7 | v7), F1: v7, F2: 42}
                _ = v152
            default:
                var v153 uint32 = 2
                _ = v153
            }
        case "hello":
            var v154 uint = 2
            _ = v154
        }
        _ = fi14
    }
    switch "foo bar" {
    case "\n\t":
        delete(v2, v3)
    case "hello":
        fmt.Println(v7, v3, v7)
    }
    v155 := GenS28{F0: 2147483647}
    v7 += v7
    var v156 int = (v7 - 657)
    for fi157 := 0; fi157 < 11; fi157++ {
        fmt.Println(v5)
        for fi158 := 0; fi158 < 9; fi158++ {
            var v159 int16 = -1
            v160 := GenS29{F0: 937, F1: fi158, F2: -1}
            _ = fi158
            _ = v159
            _ = v160
        }
        _ = fi157
    }
    v3 = v3
    v5 += (^(v5))
    for fi161 := 0; fi161 < 17; fi161++ {
        var v162 string = "foo bar"
        for fi163 := 0; fi163 < 13; fi163++ {
            v164 := GenS14{F0: (fi161 * -9223372036854775808), F1: -1, F2: v5}
            fmt.Println(fi161)
            switch (fi161 * -1) {
            case -2147483648:
                if (2 >= 42) {
                    if true {
                        v165 := GenS28{F0: 2147483648}
                        _ = v165
                    } else {
                        var v166 uint64 = 4294967295
                        if (!true) {
                            switch (v156 & v7) {
                            case -58:
                                v167 := GenS3{F0: 269, F1: 255}
                                _ = v167
                            case 2147483647:
                                fmt.Println(fi163, fi161, v5)
                                var v168 float64 = (0.1 * 1e10)
                                var v169 uint8 = v5
                                _ = v168
                                _ = v169
                            case 9223372036854775807:
                                v156 -= 517
                            }
                        } else {
                            if (!(!(!((true || ((v5 >= (v5 - v5)) && (-448 == 2147483647))) || true)))) {
                                switch v3 {
                                case "a":
                                    _ = len(v4)
                                case "\n\t":
                                    switch v162 {
                                    case "foo bar":
                                        var v170 int64 = 9223372036854775807
                                        _ = v170
                                    case "hello":
                                        fmt.Println(v166, v5)
                                    case "":
                                        var v171 int = (fi161 & fi163)
                                        _ = v171
                                    default:
                                        if (1 < (^(v5))) {
                                            fmt.Println(v5, v7, fi161)
                                            var v172 float32 = float32((1.5 + 1e10))
                                            if true {
                                                v166 = (v166 | 207)
                                                var v173 bool = false
                                                v156 = (v156 % (((fi163 >> ((2) & 63))) | 1))
                                                if true {
                                                    v174 := GenS30{F0: 1, F1: v173, F2: 592}
                                                    _ = v174
                                                }
                                                if v173 {
                                                    v162 = ("hello" + (v162 + (v162 + v162)))
                                                } else {
                                                    v166 -= v166
                                                    _ = len(v2)
                                                }
                                                v5 = 2
                                                _ = v173
                                            } else {
                                                fmt.Println(v172)
                                                switch (fi161 & -376) {
                                                case 128:
                                                    fmt.Println(v156)
                                                case -50:
                                                    if len(v6) > 0 {
                                                        v6[0] = 2
                                                    }
                                                }
                                            }
                                            if ((-128 >= 61) || (v5 > 1)) {
                                                var v175 uint8 = (v5 + (v5 * (v5 * 119)))
                                                var v176 int32 = 2147483647
                                                if false {
                                                    if len(v6) > 0 {
                                                        v6[0] = 2147483647
                                                    }
                                                    v5 -= (v175 | v5)
                                                    fmt.Println(v3, v176)
                                                } else {
                                                    v5 -= v175
                                                    fmt.Println(fi161, v176)
                                                    fmt.Println(v166, fi161)
                                                }
                                                _ = v175
                                                _ = v176
                                            } else {
                                                fmt.Println(fi163, v172)
                                            }
                                            _ = v172
                                        }
                                    }
                                default:
                                    _ = len(v2)
                                }
                            }
                            if (-1 == 93) {
                                v5 = (v5 | v5)
                                v162 = v3
                            } else {
                                var v177 uint64 = v166
                                var v178 uint16 = 255
                                v3 = v3
                                _ = v177
                                _ = v178
                            }
                            if (1 <= v166) {
                                v179 := GenS13{F0: 42, F1: (^(v166)), F2: v162}
                                var v180 int16 = 42
                                switch 42 {
                                case 916:
                                    v181 := GenS19{F0: (((((v7 != v7) || true) && true) && ((-(fi161)) == v7)) || (uint16(v166) <= 0)), F1: v166, F2: 2147483647}
                                    _ = v181
                                case 32768:
                                    if ((v5 & (v5 & (v5 + v5))) >= v5) {
                                        var v182 uint = 99
                                        _ = v182
                                    } else {
                                        var v183 uint = 861
                                        if true {
                                            v183 = 1
                                        }
                                        var v184 uint16 = 45
                                        _ = v183
                                        _ = v184
                                    }
                                default:
                                    fmt.Println(v166)
                                }
                                _ = v179
                                _ = v180
                            }
                            var v185 int32 = -1
                            fmt.Println(fi161, fi161, v5)
                            delete(v2, ("a" + v162))
                            switch v156 {
                            case 128:
                                fmt.Println(v3)
                            default:
                                if len(v4) > 0 {
                                    v4[0] = (^(v156))
                                }
                                var v186 uint16 = 65535
                                v5 = (v5 * 127)
                                _ = v186
                            }
                            _ = v185
                        }
                        _ = v166
                    }
                }
            case 1:
                v187 := GenS20{F0: (int32(v5) == 902)}
                _ = v187
            case 32767:
                var v188 uint16 = 2
                _ = v188
            }
            v5 = 3
            fmt.Println(v5, v162)
            _ = fi163
            _ = v164
        }
        fmt.Println(v156, fi161)
        _ = fi161
        _ = v162
    }
    var v189 uint8 = 127
    v156 = v7
    v189 = v189
    v190 := v1[(v156 * v7)]
    for fi191 := 0; fi191 < 14; fi191++ {
        for fi192 := 0; fi192 < 14; fi192++ {
            var v193 float32 = (3.14159 - (0.1 / 1.5))
            _ = fi192
            _ = v193
        }
        if true {
            if len(v6) > 0 {
                v6[0] = 2147483647
            }
        } else {
            for fi194 := 0; fi194 < 15; fi194++ {
                v156 -= (-(fi194))
                v156 += v7
                if ((1e10 <= -2.25) && (true && (0.0 >= 0.0))) {
                    v195 := v1[(-(fi191))]
                    switch v3 {
                    case "foo bar":
                        v196 := GenS31{F0: (fi191 << ((1) & 63)), F1: 323, F2: fi194}
                        _ = v196
                    case "\n\t":
                        var v197 uint16 = 606
                        _ = v197
                    default:
                        switch v3 {
                        case "\n\t":
                            if (!(v3 != ((v3 + "") + v3))) {
                                v198 := GenS32{F0: v195, F1: 2}
                                var v199 int32 = 264
                                v199 = v199
                                _ = v198
                                _ = v199
                            } else {
                                var v200 int16 = int16(v7)
                                var v201 int = v7
                                _ = v200
                                _ = v201
                            }
                        case "hello":
                            switch (-(fi194)) {
                            case -836:
                                v5 = v189
                            default:
                                var v202 bool = false
                                _ = v202
                            }
                            var v203 uint64 = 592
                            _ = v203
                        }
                    }
                    _ = v195
                }
                _ = fi194
            }
            if (1.5 == -2.25) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = len(v4)
            }
            if (-405 >= -580) {
                if (-1 == -1) {
                    v5 = (v5 / ((v189) | 1))
                    if true {
                        var v204 int64 = -811
                        var v205 int32 = -178
                        _ = v204
                        _ = v205
                    } else {
                        var v206 float64 = 1e10
                        for fi207 := 0; fi207 < 3; fi207++ {
                            v208 := GenS25{F0: 4294967295, F1: 132}
                            _ = fi207
                            _ = v208
                        }
                        switch v3 {
                        case "hello":
                            if (v206 == float64(0.1)) {
                                fmt.Println(v189)
                                delete(v1, v156)
                            }
                        case "":
                            for fi209 := 0; fi209 < 9; fi209++ {
                                fmt.Println(fi191, v7, v156)
                                v7 = (v7 + v156)
                                _ = fi209
                            }
                        default:
                            if false {
                                fmt.Println(v206)
                                var v210 uint64 = 18446744073709551615
                                for fi211 := 0; fi211 < 1; fi211++ {
                                    fmt.Println(v189)
                                    v212 := GenS12{F0: 1, F1: "foo bar", F2: 0}
                                    switch 127 {
                                    case 198:
                                        v210 += 4294967295
                                    default:
                                        v210 += v210
                                    }
                                    _ = fi211
                                    _ = v212
                                }
                                _ = v210
                            }
                        }
                        _ = v206
                    }
                } else {
                    if true {
                        for fi213 := 0; fi213 < 12; fi213++ {
                            v214 := GenS25{F0: 948, F1: 42}
                            v189 = v189
                            var v215 bool = true
                            var v216 string = v3
                            _ = fi213
                            _ = v214
                            _ = v215
                            _ = v216
                        }
                    }
                    v5 = (^(v189))
                    if len(v6) > 0 {
                        v6[0] = (v190 | v190)
                    }
                }
            }
        }
        switch (v3 + v3) {
        case "hello":
            if ((428 == 975) && true) {
                fmt.Println(v189, v156)
                v217 := GenS13{F0: 0, F1: 42, F2: v3}
                _ = v217
            }
        case "a":
            for fi218 := 0; fi218 < 14; fi218++ {
                fmt.Println(v5, fi191)
                _ = fi218
            }
        default:
            if (-2147483648 != v7) {
                v219 := GenS24{F0: -164}
                _ = v219
            }
        }
        _ = fi191
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v156:", v156)
    fmt.Println("v189:", v189)
    fmt.Println("v190:", v190)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v155
    _ = v156
    _ = v189
    _ = v190
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
