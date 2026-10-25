package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint16 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 string; F1 uint8; F2 uint32 }
type GenS5 struct { F0 string; F1 string; F2 int16 }
type GenS6 struct { F0 int64; F1 uint }
type GenS7 struct { F0 int; F1 uint8 }
type GenS8 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint32; F1 bool }
type GenS12 struct { F0 uint; F1 int32 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 int }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 bool; F1 bool }

func main() {
    var v1 int64 = 42
    v2 := make(map[int]int)
    var v3 int8 = 1
    v4 := make(map[int]bool)
    switch "foo bar" {
    case "\n\t":
        if false {
            switch 0 {
            case 737:
                fmt.Println(v3)
            case 0:
                fmt.Println(v3)
            }
            var v5 int64 = (v1 | (-(v1)))
            fmt.Println(v5)
            _ = v5
        } else {
            var v6 bool = false
            for fi7 := 0; fi7 < 10; fi7++ {
                var v8 uint16 = 42
                v1 += v1
                _ = fi7
                _ = v8
            }
            _ = v6
        }
    case "hello":
        var v9 int8 = 70
        _ = v9
    case "":
        v10 := GenS1{F0: 4294967295, F1: 42, F2: 255}
        _ = v10
    default:
        v1 = -9223372036854775808
    }
    var v11 float64 = 1.5
    fmt.Println(v1, v1, v11)
    v12 := []int{2147483647}
    var v13 int = 128
    fmt.Println(v1, v1)
    var v14 uint32 = 4294967295
    v15 := GenS1{F0: 0, F1: 1, F2: 871}
    v16 := GenS1{F0: 68, F1: 1, F2: 406}
    switch "\n\t" {
    case "a":
        v17 := v2[v13]
        _ = v17
    case "hello":
        _ = len(v12)
    case "\n\t":
        if true {
            var v18 float64 = float64(-2.25)
            var v19 int32 = int32(v14)
            _ = v18
            _ = v19
        } else {
            switch (v13 + (v13 << ((1) & 63))) {
            case 2147483647:
                switch (("a" + "") + "\n\t") {
                case "foo bar":
                    for fi20 := 0; fi20 < 13; fi20++ {
                        v11 -= v11
                        v14 = v14
                        if (225 <= 42) {
                            for fi21 := 0; fi21 < 13; fi21++ {
                                var v22 int8 = (^(v3))
                                v13 += -9223372036854775808
                                _ = fi21
                                _ = v22
                            }
                            fmt.Println(v11)
                        }
                        _ = fi20
                    }
                case "hello":
                    switch ((("foo bar" + ("" + (("hello" + "hello") + (((("hello" + "") + "a") + "") + "hello")))) + "") + ("\n\t" + "a")) {
                    case "a":
                        var v23 bool = ((true || false) && false)
                        _ = v23
                    case "":
                        for fi24 := 0; fi24 < 9; fi24++ {
                            for fi25 := 0; fi25 < 6; fi25++ {
                                var v26 uint16 = 65535
                                _ = fi25
                                _ = v26
                            }
                            switch (v13 % (((-(fi24))) | 1)) {
                            case 4294967295:
                                if false {
                                    var v27 bool = false
                                    _ = v27
                                } else {
                                    switch "foo bar" {
                                    case "foo bar":
                                        v1 = (v1 >> ((1) & 63))
                                    case "\n\t":
                                        if (4294967295 <= 271) {
                                            fmt.Println(v3)
                                            for fi28 := 0; fi28 < 17; fi28++ {
                                                v29 := GenS2{F0: v3, F1: v3}
                                                v30 := GenS1{F0: 255, F1: 564, F2: 1}
                                                v31 := GenS1{F0: 420, F1: 1, F2: 853}
                                                _ = fi28
                                                _ = v29
                                                _ = v30
                                                _ = v31
                                            }
                                        } else {
                                            for fi32 := 0; fi32 < 8; fi32++ {
                                                v33 := GenS3{F0: v3, F1: 65535}
                                                v11 = 0.0
                                                var v34 int = -748
                                                _ = fi32
                                                _ = v33
                                                _ = v34
                                            }
                                        }
                                    case "hello":
                                        v35 := GenS4{F0: "\n\t", F1: 255, F2: (^(v14))}
                                        _ = v35
                                    }
                                    if true {
                                        for fi36 := 0; fi36 < 1; fi36++ {
                                            v37 := GenS2{F0: v3, F1: -1}
                                            var v38 int16 = -80
                                            _ = fi36
                                            _ = v37
                                            _ = v38
                                        }
                                        if (!(true || ((false && true) && true))) {
                                            for fi39 := 0; fi39 < 13; fi39++ {
                                                switch (^(v13)) {
                                                case 127:
                                                    v11 -= v11
                                                case 546:
                                                    if false {
                                                        v13 = 42
                                                        var v40 int8 = (v3 ^ v3)
                                                        _ = len(v2)
                                                        _ = v40
                                                    } else {
                                                        fmt.Println(fi24)
                                                    }
                                                case 0:
                                                    fmt.Println(fi39, v13)
                                                }
                                                var v41 int16 = -32768
                                                if (v14 > v14) {
                                                    if (1 < (v41 | -32768)) {
                                                        fmt.Println(v41)
                                                    }
                                                    switch "" {
                                                    case "hello":
                                                        var v42 int = fi39
                                                        _ = v42
                                                    default:
                                                        _ = len(v12)
                                                    }
                                                    var v43 int = fi24
                                                    _ = v43
                                                } else {
                                                    if true {
                                                        if true {
                                                            var v44 int = v13
                                                            v45 := GenS5{F0: "foo bar", F1: "a", F2: v41}
                                                            _ = v44
                                                            _ = v45
                                                        }
                                                        if ((0 == 2147483647) || (32767 > -994)) {
                                                            v1 = (v1 & v1)
                                                        } else {
                                                            var v46 uint16 = 1
                                                            fmt.Println(v1, v41, v41)
                                                            _ = v46
                                                        }
                                                    }
                                                    if ("\n\t" != ("" + "foo bar")) {
                                                        delete(v4, (fi24 ^ fi39))
                                                    }
                                                    v11 += float64((((((float32(1.5) / float32((v11 - v11))) * 1.5) + 1.5) * 0.1) + float32(v11)))
                                                }
                                                _ = fi39
                                                _ = v41
                                            }
                                            for fi47 := 0; fi47 < 2; fi47++ {
                                                if (fi47 > (fi47 - fi24)) {
                                                    v1 = (v1 % ((v1) | 1))
                                                } else {
                                                    if true {
                                                        fmt.Println(v1, v3)
                                                        v48 := GenS1{F0: 2, F1: -2147483648, F2: 448}
                                                        var v49 uint32 = 0
                                                        v50 := GenS6{F0: 0, F1: 535}
                                                        v51 := GenS7{F0: -2147483648, F1: 42}
                                                        switch "a" {
                                                        case "a":
                                                            var v52 float32 = -2.25
                                                            _ = v52
                                                        case "\n\t":
                                                            v3 += v3
                                                        case "":
                                                            v14 += 42
                                                        }
                                                        if (!false) {
                                                            var v53 int64 = (^(v1))
                                                            _ = v53
                                                        }
                                                        _ = v48
                                                        _ = v49
                                                        _ = v50
                                                        _ = v51
                                                    } else {
                                                        v1 -= (v1 ^ (v1 << ((1) & 63)))
                                                        v1 += v1
                                                        v3 = v3
                                                    }
                                                    if len(v12) >= 2 {
                                                        v12 = v12[1:len(v12)]
                                                    }
                                                }
                                                fmt.Println(v14, fi24)
                                                v14 += (v14 - 4294967295)
                                                _ = fi47
                                            }
                                            var v54 string = "hello"
                                            _ = v54
                                        }
                                        v1 += (v1 - v1)
                                    } else {
                                        var v55 int16 = 32767
                                        v3 += 100
                                        _ = len(v12)
                                        _ = v55
                                    }
                                }
                            }
                            var v56 string = ("hello" + (("hello" + ("a" + ("" + (("foo bar" + "a") + ("a" + "\n\t"))))) + "\n\t"))
                            _ = fi24
                            _ = v56
                        }
                    case "hello":
                        _ = len(v2)
                    }
                default:
                    v57 := GenS8{F0: 42, F1: v14, F2: 65535}
                    _ = v57
                }
            }
            var v58 int32 = 0
            if false {
                for fi59 := 0; fi59 < 17; fi59++ {
                    for fi60 := 0; fi60 < 12; fi60++ {
                        v14 -= v14
                        v1 += v1
                        var v61 int16 = 42
                        _ = fi60
                        _ = v61
                    }
                    for fi62 := 0; fi62 < 10; fi62++ {
                        v63 := GenS2{F0: (v3 % ((v3) | 1)), F1: (-(v3))}
                        if false {
                            var v64 int = (fi62 / ((fi62) | 1))
                            var v65 int64 = v1
                            if (!false) {
                                if true {
                                    if ((!false) || ((!true) && (((767 == 255) && (-32768 > 32767)) && ((3.14159 / 3.14159) <= ((float64((-2.25 * ((3.14159 * 3.14159) / 1.0))) + 3.14159) - 3.14159))))) {
                                        fmt.Println(fi62, v11, v58)
                                        if (2 >= 128) {
                                            v3 -= 42
                                        } else {
                                            switch "" {
                                            case "":
                                                var v66 bool = false
                                                _ = v66
                                            case "foo bar":
                                                v67 := v2[v64]
                                                _ = v67
                                            case "\n\t":
                                                v68 := GenS2{F0: 127, F1: v3}
                                                _ = v68
                                            }
                                            if false {
                                                if true {
                                                    switch v64 {
                                                    case 2147483647:
                                                        if true {
                                                            var v69 uint = 65535
                                                            if false {
                                                                v1 = -9223372036854775808
                                                                var v70 int16 = 1000
                                                                _ = v70
                                                            } else {
                                                                fmt.Println(v14)
                                                                if (!(!(!true))) {
                                                                    if (true || true) {
                                                                        var v71 float64 = (v11 + v11)
                                                                        if true {
                                                                            v72 := GenS7{F0: (fi62 & 127), F1: uint8(v14)}
                                                                            _ = len(v2)
                                                                            v73 := GenS2{F0: v3, F1: v3}
                                                                            v11 = ((-2.25 / 1.5) * v11)
                                                                            v3 = -128
                                                                            _ = v72
                                                                            _ = v73
                                                                        } else {
                                                                            v12 = append(v12, (fi59 - 165))
                                                                            fmt.Println(v11)
                                                                        }
                                                                        _ = v71
                                                                    } else {
                                                                        var v74 string = "a"
                                                                        var v75 uint32 = (v14 | (^(v14)))
                                                                        var v76 uint16 = 405
                                                                        fmt.Println(v65)
                                                                        var v77 uint64 = 779
                                                                        _ = v74
                                                                        _ = v75
                                                                        _ = v76
                                                                        _ = v77
                                                                    }
                                                                    if false {
                                                                        v11 = 0.0
                                                                    } else {
                                                                        var v78 string = "foo bar"
                                                                        v4[v13] = (!true)
                                                                        var v79 string = (("foo bar" + "\n\t") + (("foo bar" + v78) + v78))
                                                                        _ = v78
                                                                        _ = v79
                                                                    }
                                                                    var v80 int8 = 100
                                                                    _ = v80
                                                                } else {
                                                                    v81 := GenS1{F0: v69, F1: (v58 ^ v58), F2: 32767}
                                                                    _ = v81
                                                                }
                                                                switch "hello" {
                                                                case "a":
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        if (!(true && false)) {
                                                                            var v82 uint16 = 65535
                                                                            _ = v82
                                                                        } else {
                                                                            switch (fi62 - (fi62 & fi62)) {
                                                                            case 128:
                                                                                if true {
                                                                                    if (2 < 925) {
                                                                                        if (true && true) {
                                                                                            v14 -= v14
                                                                                        } else {
                                                                                            v83 := GenS7{F0: fi59, F1: 2}
                                                                                            v69 = (v69 / ((v69) | 1))
                                                                                            var v84 bool = false
                                                                                            _ = v83
                                                                                            _ = v84
                                                                                        }
                                                                                        var v85 int64 = v65
                                                                                        var v86 bool = false
                                                                                        delete(v2, v64)
                                                                                        fmt.Println(v3, v14, v11)
                                                                                        _ = v85
                                                                                        _ = v86
                                                                                    } else {
                                                                                        v87 := GenS9{F0: 2}
                                                                                        _ = v87
                                                                                    }
                                                                                    var v88 uint8 = 1
                                                                                    var v89 uint64 = 530
                                                                                    _ = v88
                                                                                    _ = v89
                                                                                }
                                                                            default:
                                                                                fmt.Println(v58, v69, v69)
                                                                                fmt.Println(v58, fi59)
                                                                                v11 -= (float64(0.1) / 1.0)
                                                                            }
                                                                            switch (-(fi62)) {
                                                                            case 32767:
                                                                                fmt.Println(v13, v1, v3)
                                                                            case 42:
                                                                                if len(v12) > 0 {
                                                                                    v12[0] = (fi59 & fi62)
                                                                                }
                                                                            case -9223372036854775808:
                                                                                v90 := GenS10{F0: (v14 - 2147483648)}
                                                                                _ = v90
                                                                            default:
                                                                                var v91 int64 = (v65 >> ((v69) & 63))
                                                                                _ = v91
                                                                            }
                                                                            if ((!true) && (false || true)) {
                                                                                v14 -= 2147483647
                                                                            }
                                                                        }
                                                                    default:
                                                                        _ = len(v4)
                                                                    }
                                                                }
                                                            }
                                                            _ = v69
                                                        } else {
                                                            v11 = ((((3.14159 + v11) * 1e10) - ((v11 * v11) - v11)) - (3.14159 + v11))
                                                            var v92 uint = 230
                                                            switch 128 {
                                                            case 32767:
                                                                var v93 float32 = (float32((v11 - (v11 - float64(1e10)))) / 0.1)
                                                                _ = v93
                                                            case -9223372036854775808:
                                                                switch (fi59 ^ (fi59 - (fi59 & (^(v64))))) {
                                                                case 9223372036854775807:
                                                                    if true {
                                                                        if (!false) {
                                                                            var v94 int = (v13 >> ((v92) & 63))
                                                                            switch ("a" + "foo bar") {
                                                                            case "":
                                                                                switch fi62 {
                                                                                case 32767:
                                                                                    switch int(v1) {
                                                                                    case 4294967295:
                                                                                        if ((!true) && (((v11 - 0.1) * 3.14159) == -2.25)) {
                                                                                            var v95 string = "hello"
                                                                                            _ = v95
                                                                                        }
                                                                                    case 9223372036854775807:
                                                                                        if (1 != 42) {
                                                                                            v3 += 100
                                                                                        }
                                                                                    }
                                                                                case 128:
                                                                                    v96 := GenS11{F0: (^(v14)), F1: false}
                                                                                    _ = v96
                                                                                case 127:
                                                                                    if false {
                                                                                        var v97 uint8 = uint8(v92)
                                                                                        var v98 uint32 = uint32(v97)
                                                                                        v3 -= 100
                                                                                        _ = v97
                                                                                        _ = v98
                                                                                    } else {
                                                                                        if len(v12) >= 2 {
                                                                                            v12 = v12[1:len(v12)]
                                                                                        }
                                                                                        switch fi62 {
                                                                                        case 340:
                                                                                            v14 = v14
                                                                                        default:
                                                                                            var v99 int = (^(fi62))
                                                                                            _ = v99
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    switch -9223372036854775808 {
                                                                                    case 268:
                                                                                        v100 := GenS12{F0: (v92 >> (((v92 * (v92 + 255))) & 63)), F1: v58}
                                                                                        _ = v100
                                                                                    case -2147483648:
                                                                                        v4[649] = ((3.14159 != 0.1) && false)
                                                                                    default:
                                                                                        _ = len(v12)
                                                                                    }
                                                                                }
                                                                            case "\n\t":
                                                                                switch fi62 {
                                                                                case -398:
                                                                                    fmt.Println(fi59, v65, v64)
                                                                                case 127:
                                                                                    if ((v94 / (((v64 >> ((v92) & 63))) | 1)) <= (v64 % ((fi59) | 1))) {
                                                                                        if (!false) {
                                                                                            var v101 string = "hello"
                                                                                            if (true && (fi59 < 4294967295)) {
                                                                                                var v102 float64 = (v11 - float64(-2.25))
                                                                                                var v103 string = v101
                                                                                                _ = v102
                                                                                                _ = v103
                                                                                            } else {
                                                                                                v1 += v1
                                                                                                switch v101 {
                                                                                                case "foo bar":
                                                                                                    var v104 uint = v92
                                                                                                    _ = v104
                                                                                                case "hello":
                                                                                                    v2[v94] = fi59
                                                                                                    var v105 uint16 = 2
                                                                                                    _ = v105
                                                                                                case "a":
                                                                                                    if len(v12) > 0 {
                                                                                                        v12[0] = (^(v64))
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            v106 := GenS13{F0: ((((0.0 - 3.14159) * ((float32(v11) - 0.1) / 1.0)) * 3.14159) > (0.1 / -2.25))}
                                                                                            _ = v101
                                                                                            _ = v106
                                                                                        } else {
                                                                                            v1 = (v1 >> (((v92 - 812)) & 63))
                                                                                            var v107 int64 = (-(v1))
                                                                                            _ = v107
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v1 = (-(v1))
                                                                                }
                                                                            }
                                                                            v108 := GenS7{F0: (fi59 & (v94 % (((v13 + v13)) | 1))), F1: 2}
                                                                            _ = v108
                                                                            _ = v94
                                                                        } else {
                                                                            fmt.Println(fi62)
                                                                            if true {
                                                                                v109 := GenS7{F0: fi62, F1: 1}
                                                                                v110 := GenS13{F0: false}
                                                                                v64 += v64
                                                                                _ = v109
                                                                                _ = v110
                                                                            }
                                                                            _ = len(v4)
                                                                        }
                                                                    } else {
                                                                        v3 = (v3 >> ((v92) & 63))
                                                                        var v111 float64 = (float64(3.14159) * v11)
                                                                        fmt.Println(v1, v111, v111)
                                                                        v12 = append(v12, v64)
                                                                        if ((!true) && (-1 < -2147483648)) {
                                                                            fmt.Println(v3, v58, v64)
                                                                            var v112 uint64 = 0
                                                                            _ = v112
                                                                        } else {
                                                                            var v113 int8 = (v3 + int8(v1))
                                                                            _ = v113
                                                                        }
                                                                        _ = v111
                                                                    }
                                                                case 32767:
                                                                    if true {
                                                                        if ((fi62 >> (((v92 * v92)) & 63)) == (v64 >> ((v92) & 63))) {
                                                                            var v114 int8 = v3
                                                                            fmt.Println(fi59)
                                                                            var v115 uint64 = 0
                                                                            _ = v114
                                                                            _ = v115
                                                                        } else {
                                                                            v65 -= v1
                                                                        }
                                                                        var v116 string = ("" + "\n\t")
                                                                        var v117 bool = true
                                                                        _ = v116
                                                                        _ = v117
                                                                    } else {
                                                                        _ = len(v12)
                                                                        var v118 uint = 255
                                                                        if len(v12) > 0 {
                                                                            v12[0] = int(v58)
                                                                        }
                                                                        _ = v118
                                                                    }
                                                                case 128:
                                                                    _ = len(v12)
                                                                    v58 -= v58
                                                                }
                                                            default:
                                                                v119 := v4[(-(fi59))]
                                                                _ = v119
                                                            }
                                                            _ = v92
                                                        }
                                                    default:
                                                        delete(v4, (v13 % ((42) | 1)))
                                                    }
                                                } else {
                                                    if (!false) {
                                                        v1 += v65
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            fmt.Println(v58)
                                                        default:
                                                            fmt.Println(v1, v11)
                                                        }
                                                    } else {
                                                        var v120 string = ("\n\t" + "foo bar")
                                                        v65 = 0
                                                        _ = v120
                                                    }
                                                }
                                                var v121 uint = 18446744073709551615
                                                var v122 int32 = (v58 >> ((0) & 63))
                                                _ = v121
                                                _ = v122
                                            }
                                        }
                                        _ = len(v4)
                                    }
                                    v123 := GenS14{F0: 1}
                                    switch ("foo bar" + "") {
                                    case "a":
                                        v64 -= fi62
                                    case "foo bar":
                                        v3 = (v3 | (v3 / ((v3) | 1)))
                                    default:
                                        if (!false) {
                                            if false {
                                                fmt.Println(v14, fi59, v65)
                                                v124 := GenS12{F0: 846, F1: v58}
                                                _ = v124
                                            } else {
                                                if ((("" + ((((("hello" + "\n\t") + ("\n\t" + "a")) + "") + "foo bar") + "a")) < "hello") && (255 >= 249)) {
                                                    v125 := GenS15{F0: (^(fi62))}
                                                    v65 = v65
                                                    _ = v125
                                                }
                                            }
                                        }
                                    }
                                    var v126 string = "foo bar"
                                    _ = v123
                                    _ = v126
                                }
                                if (("foo bar" + "a") > "foo bar") {
                                    v127 := GenS2{F0: -1, F1: v3}
                                    var v128 uint16 = 32767
                                    _ = v127
                                    _ = v128
                                }
                            }
                            _ = v64
                            _ = v65
                        } else {
                            v58 = v58
                            v129 := GenS6{F0: (-(v1)), F1: 117}
                            v14 += v14
                            _ = v129
                        }
                        switch ("hello" + "foo bar") {
                        case "foo bar":
                            var v130 uint32 = (v14 << ((656) & 63))
                            _ = v130
                        case "hello":
                            fmt.Println(v1)
                        default:
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                            v131 := GenS16{F0: (v1 & -30)}
                            v13 += (v13 % ((v13) | 1))
                            _ = v131
                        }
                        _ = fi62
                        _ = v63
                    }
                    v132 := v4[(fi59 % ((v13) | 1))]
                    _ = fi59
                    _ = v132
                }
            }
            _ = v58
        }
    }
    v133 := GenS14{F0: 0}
    for fi134 := 0; fi134 < 4; fi134++ {
        var v135 float64 = 1.5
        _ = fi134
        _ = v135
    }
    for fi136 := 0; fi136 < 9; fi136++ {
        var v137 uint = 4294967295
        switch "" {
        case "foo bar":
            _ = len(v12)
        case "a":
            _ = len(v12)
        case "hello":
            v137 += v137
        default:
            switch "hello" {
            case "":
                switch fi136 {
                case 168:
                    fmt.Println(v137, fi136)
                case 0:
                    v2[394] = v13
                }
            case "a":
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            case "\n\t":
                v138 := GenS13{F0: false}
                _ = v138
            default:
                fmt.Println(v3, v3)
            }
        }
        _ = fi136
        _ = v137
    }
    fmt.Println(v11, v11)
    v139 := GenS17{F0: false, F1: false}
    v11 = (v11 * v11)
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    v1 -= int64(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v133
    _ = v139
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
