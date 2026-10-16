package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 string }
type GenS2 struct { F0 int8; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS7 struct { F0 uint16; F1 uint32; F2 uint }
type GenS8 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint64; F1 uint8 }
type GenS13 struct { F0 int16; F1 int64 }
type GenS14 struct { F0 uint32; F1 uint }
type GenS15 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS16 struct { F0 int64; F1 uint; F2 uint }
type GenS17 struct { F0 int64; F1 uint }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 int32; F1 int16; F2 uint8 }

func main() {
    v1 := []int{90, 4294967295}
    var v2 uint = 100
    v3 := make(map[uint32]bool)
    if (false || (0 < 1)) {
        v1 = append(v1, -285)
        v2 = (^(v2))
        switch 9223372036854775807 {
        case -215:
            var v4 uint8 = 255
            _ = v4
        default:
            v5 := GenS1{F0: 255, F1: 42, F2: "\n\t"}
            _ = v5
        }
    } else {
        var v6 string = "a"
        if (-25 >= -1) {
            var v7 int32 = 976
            v8 := GenS1{F0: 2, F1: -43, F2: v6}
            for fi9 := 0; fi9 < 15; fi9++ {
                switch v6 {
                case "a":
                    var v10 int = 634
                    _ = v10
                case "hello":
                    if false {
                        v6 = "\n\t"
                        for fi11 := 0; fi11 < 18; fi11++ {
                            v2 -= (v2 % (((v2 - v2)) | 1))
                            _ = fi11
                        }
                        for fi12 := 0; fi12 < 17; fi12++ {
                            var v13 string = v6
                            if ((1 != 255) || false) {
                                v14 := GenS1{F0: 194, F1: fi9, F2: "\n\t"}
                                var v15 uint32 = 407
                                if false {
                                    switch int(v2) {
                                    case 341:
                                        v16 := GenS2{F0: 0, F1: 42, F2: (-(v7))}
                                        _ = v16
                                    default:
                                        v1 = append(v1, fi9)
                                    }
                                    v7 = int32(fi12)
                                    if (false && true) {
                                        v7 += 1
                                        fmt.Println(fi9)
                                    }
                                    var v17 uint64 = 0
                                    var v18 int64 = 381
                                    _ = v17
                                    _ = v18
                                } else {
                                    if false {
                                        v19 := GenS1{F0: 2, F1: 32768, F2: "foo bar"}
                                        var v20 uint16 = 255
                                        v21 := GenS2{F0: 100, F1: (v15 ^ v15), F2: (v7 * 0)}
                                        fmt.Println(v15, v15, v15)
                                        v22 := GenS3{F0: 0, F1: uint(v20)}
                                        switch "a" {
                                        case "hello":
                                            v23 := GenS3{F0: 42, F1: (v2 & 771)}
                                            _ = v23
                                        case "":
                                            v6 = v6
                                        }
                                        if true {
                                            if (false || true) {
                                                var v24 uint32 = v15
                                                _ = v24
                                            }
                                            if false {
                                                v6 = ((v6 + (v6 + "")) + v13)
                                                v6 = ("hello" + v13)
                                                if false {
                                                    v7 = v7
                                                }
                                            }
                                        } else {
                                            fmt.Println(fi9, v20, fi9)
                                            fmt.Println(fi9, v2)
                                            switch (v13 + "hello") {
                                            case "\n\t":
                                                _ = len(v1)
                                                switch (fi12 % (((fi12 - 32767)) | 1)) {
                                                case 32768:
                                                    switch (fi9 % ((0) | 1)) {
                                                    case 32767:
                                                        v25 := GenS1{F0: 125, F1: (fi9 | (fi12 * int(v7))), F2: v6}
                                                        _ = v25
                                                    case 2147483647:
                                                        v26 := GenS4{F0: 127}
                                                        _ = v26
                                                    case 127:
                                                        switch (fi9 | (fi9 * (fi12 / (((-(fi9))) | 1)))) {
                                                        case -997:
                                                            v6 = v13
                                                        case -851:
                                                            var v27 bool = true
                                                            switch ((("hello" + (v6 + (("\n\t" + ((("\n\t" + "") + v6) + v6)) + (v13 + "foo bar")))) + v13) + v6) {
                                                            case "a":
                                                                v20 += v20
                                                                var v28 uint32 = 38
                                                                _ = v28
                                                            case "foo bar":
                                                                fmt.Println(v2)
                                                            case "":
                                                                if (127 == 1) {
                                                                    v20 = v20
                                                                    v29 := GenS5{F0: fi9, F1: 127}
                                                                    var v30 float32 = 0.0
                                                                    v13 = v6
                                                                    v31 := v3[(v15 * v15)]
                                                                    fmt.Println(v6, v2)
                                                                    switch -2147483648 {
                                                                    case 1:
                                                                        var v32 int8 = -118
                                                                        _ = v32
                                                                    }
                                                                    _ = v29
                                                                    _ = v30
                                                                    _ = v31
                                                                } else {
                                                                    var v33 int8 = 100
                                                                    v20 -= 2
                                                                    _ = v33
                                                                }
                                                            default:
                                                                var v34 int64 = 9223372036854775807
                                                                _ = v34
                                                            }
                                                            v6 = "hello"
                                                            v27 = v27
                                                            _ = v27
                                                        case 2147483647:
                                                            fmt.Println(v15)
                                                        default:
                                                            v13 = v13
                                                        }
                                                    default:
                                                        var v35 string = v6
                                                        if true {
                                                            v36 := GenS2{F0: 127, F1: (^(v15)), F2: v7}
                                                            _ = v36
                                                        }
                                                        delete(v3, v15)
                                                        _ = v35
                                                    }
                                                }
                                                switch v13 {
                                                case "a":
                                                    _ = len(v3)
                                                case "foo bar":
                                                    var v37 int64 = -664
                                                    _ = v37
                                                }
                                            case "":
                                                var v38 float32 = 0.0
                                                _ = v38
                                            case "foo bar":
                                                var v39 int32 = v7
                                                _ = v39
                                            default:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                            var v40 string = v13
                                            _ = v40
                                        }
                                        switch v6 {
                                        case "foo bar":
                                            if len(v1) > 0 {
                                                v1[0] = fi9
                                            }
                                        case "":
                                            var v41 uint32 = 42
                                            _ = v41
                                        }
                                        _ = v19
                                        _ = v20
                                        _ = v21
                                        _ = v22
                                    }
                                    if (v7 != v7) {
                                        v3[v15] = false
                                        var v42 string = ((v6 + v6) + v13)
                                        var v43 int = int(v15)
                                        var v44 uint32 = (v15 ^ (v15 % ((v15) | 1)))
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    }
                                    switch 128 {
                                    case 0:
                                        v45 := GenS6{F0: 42, F1: 255, F2: 0}
                                        _ = v45
                                    case 9223372036854775807:
                                        _ = len(v3)
                                    default:
                                        if ((fi12 & 0) <= 446) {
                                            if false {
                                                if ((!(!false)) || ((true && false) && true)) {
                                                    v46 := GenS7{F0: 1, F1: 2147483647, F2: v2}
                                                    _ = v46
                                                } else {
                                                    v47 := GenS3{F0: 135, F1: v2}
                                                    v6 = "foo bar"
                                                    if ((fi12 + -667) < (fi9 | fi12)) {
                                                        if (true && (0 == 255)) {
                                                            fmt.Println(fi9, v7)
                                                        }
                                                        v2 -= 2
                                                        if false {
                                                            switch "\n\t" {
                                                            case "hello":
                                                                var v48 float64 = 0.0
                                                                _ = v48
                                                            case "foo bar":
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    fmt.Println(v7)
                                                                case "hello":
                                                                    v15 -= v15
                                                                    if false {
                                                                        var v49 uint = (v2 >> (((v2 - v2)) & 63))
                                                                        _ = v49
                                                                    } else {
                                                                        v50 := GenS3{F0: 128, F1: v2}
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        _ = v50
                                                                    }
                                                                case "foo bar":
                                                                    _ = len(v1)
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        v7 = 0
                                                    }
                                                    _ = v47
                                                }
                                                fmt.Println(v2)
                                            } else {
                                                if false {
                                                    v3[2] = true
                                                    var v51 uint64 = 666
                                                    _ = v51
                                                } else {
                                                    var v52 uint64 = 1
                                                    var v53 uint = v2
                                                    _ = v52
                                                    _ = v53
                                                }
                                                v1 = append(v1, fi9)
                                                v13 = v6
                                            }
                                            if true {
                                                v7 += (-(v7))
                                            }
                                            _ = len(v3)
                                        } else {
                                            v2 = (v2 & v2)
                                            v54 := v3[(v15 / ((737) | 1))]
                                            delete(v3, (v15 & v15))
                                            fmt.Println(v13, fi12)
                                            v54 = v54
                                            var v55 int = fi12
                                            _ = v54
                                            _ = v55
                                        }
                                    }
                                }
                                _ = v14
                                _ = v15
                            }
                            _ = fi12
                            _ = v13
                        }
                    } else {
                        for fi56 := 0; fi56 < 9; fi56++ {
                            fmt.Println(fi9)
                            v57 := GenS5{F0: (fi56 % ((32767) | 1)), F1: 1}
                            var v58 uint64 = 18446744073709551615
                            v59 := GenS6{F0: 9223372036854775807, F1: 32767, F2: 811}
                            if true {
                                var v60 float32 = 1e10
                                _ = len(v3)
                                _ = v60
                            }
                            var v61 uint64 = v58
                            v62 := GenS1{F0: 78, F1: fi9, F2: "a"}
                            var v63 float64 = float64(-2.25)
                            _ = fi56
                            _ = v57
                            _ = v58
                            _ = v59
                            _ = v61
                            _ = v62
                            _ = v63
                        }
                        if (32767 < 158) {
                            var v64 int64 = -9223372036854775808
                            for fi65 := 0; fi65 < 10; fi65++ {
                                v64 += 1
                                var v66 uint16 = 32767
                                _ = fi65
                                _ = v66
                            }
                            _ = v64
                        }
                    }
                case "foo bar":
                    if false {
                        v67 := GenS7{F0: 32767, F1: 2, F2: 0}
                        if (-1 <= 0) {
                            for fi68 := 0; fi68 < 2; fi68++ {
                                if (!false) {
                                    var v69 int16 = 621
                                    fmt.Println(v2)
                                    _ = v69
                                }
                                v70 := v3[2147483647]
                                v7 = v7
                                if v70 {
                                    var v71 int16 = -996
                                    _ = v71
                                }
                                _ = fi68
                                _ = v70
                            }
                        } else {
                            var v72 string = v6
                            _ = v72
                        }
                        _ = len(v1)
                        _ = v67
                    }
                default:
                    if true {
                        v2 = (v2 >> (((v2 ^ (v2 ^ v2))) & 63))
                        var v73 float32 = -2.25
                        fmt.Println(v2, fi9, v73)
                        _ = v73
                    } else {
                        if (uint16(v2) < 1) {
                            var v74 int = (fi9 | fi9)
                            _ = v74
                        } else {
                            if (!(false || ((v7 <= int32(fi9)) && false))) {
                                _ = len(v1)
                                var v75 uint8 = 255
                                var v76 uint16 = 1
                                _ = v75
                                _ = v76
                            }
                            v7 += v7
                        }
                        var v77 bool = (127 >= 127)
                        var v78 uint32 = 2147483647
                        _ = v77
                        _ = v78
                    }
                }
                fmt.Println(v6, v2, v7)
                var v79 uint = 65535
                _ = fi9
                _ = v79
            }
            _ = v7
            _ = v8
        }
        v6 = (("a" + v6) + v6)
        _ = v6
    }
    if false {
        switch "" {
        case "a":
            var v80 bool = (!false)
            _ = v80
        case "":
            v2 -= (v2 + (v2 << (((v2 & (^(v2)))) & 63)))
        case "foo bar":
            fmt.Println(v2)
        default:
            var v81 int8 = 1
            _ = v81
        }
        for fi82 := 0; fi82 < 11; fi82++ {
            switch ((("a" + "hello") + "") + "\n\t") {
            case "hello":
                for fi83 := 0; fi83 < 4; fi83++ {
                    fmt.Println(fi82)
                    _ = fi83
                }
            case "foo bar":
                for fi84 := 0; fi84 < 11; fi84++ {
                    switch (^(fi82)) {
                    case -9223372036854775808:
                        var v85 int16 = 42
                        _ = v85
                    case 4294967295:
                        fmt.Println(fi84)
                    case 41:
                        var v86 float64 = 1.5
                        _ = v86
                    default:
                        if (-888 >= 0) {
                            if (116 > 45) {
                                switch "\n\t" {
                                case "hello":
                                    v87 := GenS8{F0: 1, F1: 255, F2: 53}
                                    _ = v87
                                case "a":
                                    fmt.Println(fi84)
                                case "\n\t":
                                    fmt.Println(fi84, fi84)
                                default:
                                    if (!true) {
                                        v88 := GenS9{F0: 0}
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v88
                                    }
                                }
                            }
                            var v89 uint32 = 18
                            switch ("hello" + (("a" + "a") + "\n\t")) {
                            case "foo bar":
                                _ = len(v1)
                            default:
                                v2 = 65535
                            }
                            _ = v89
                        } else {
                            v90 := GenS8{F0: -958, F1: 128, F2: 42}
                            v2 -= 42
                            v2 += v2
                            _ = v90
                        }
                    }
                    v2 += (v2 / ((v2) | 1))
                    if ((((((-2.25 - 1.5) - -2.25) + -2.25) * -2.25) <= (1e10 * 0.0)) && (!true)) {
                        var v91 string = (("foo bar" + "hello") + "\n\t")
                        _ = v91
                    }
                    var v92 int32 = 512
                    if true {
                        v93 := GenS6{F0: 1, F1: 0, F2: 136}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if len(v1) > 0 {
                            v1[0] = (fi82 / (((fi84 >> (((v2 ^ v2)) & 63))) | 1))
                        }
                        _ = v93
                    } else {
                        v2 += v2
                    }
                    _ = fi84
                    _ = v92
                }
            case "":
                for fi94 := 0; fi94 < 12; fi94++ {
                    v2 = v2
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi94
                }
            }
            for fi95 := 0; fi95 < 12; fi95++ {
                switch "" {
                case "\n\t":
                    if (false && ((float64(1.5) > (3.14159 - 1.5)) || (-580 <= -556))) {
                        var v96 int64 = 42
                        v2 = (^(v2))
                        _ = v96
                    } else {
                        v2 -= 357
                        switch "hello" {
                        case "\n\t":
                            _ = len(v3)
                        }
                        var v97 uint = 812
                        v98 := GenS10{F0: 65535}
                        _ = v97
                        _ = v98
                    }
                case "":
                    var v99 uint = v2
                    _ = v99
                default:
                    switch ("" + ("a" + "foo bar")) {
                    case "a":
                        v1 = append(v1, fi82)
                    case "\n\t":
                        v100 := GenS4{F0: 42}
                        _ = v100
                    case "hello":
                        var v101 int16 = -80
                        _ = v101
                    }
                }
                if true {
                    v2 = 0
                    fmt.Println(fi82)
                    switch "a" {
                    case "hello":
                        if (639 >= 1) {
                            var v102 uint64 = 42
                            if (true || (0 == 2147483647)) {
                                var v103 string = "\n\t"
                                switch (v103 + v103) {
                                case "":
                                    if true {
                                        v102 -= (v102 ^ 1)
                                        _ = len(v1)
                                    } else {
                                        v1 = append(v1, (fi82 | (fi82 << ((v2) & 63))))
                                        v2 += v2
                                    }
                                }
                                _ = v103
                            }
                            switch "\n\t" {
                            case "":
                                if ((true || (float64(((1.5 - 0.1) / 3.14159)) < 0.0)) || true) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    switch fi82 {
                                    case -1:
                                        v104 := GenS9{F0: 0}
                                        if (0.0 != -2.25) {
                                            var v105 int8 = -1
                                            var v106 uint16 = 180
                                            if true {
                                                var v107 float32 = 0.1
                                                v3[206] = (!true)
                                                _ = v107
                                            } else {
                                                if false {
                                                    v102 -= v102
                                                    v2 += 4294967295
                                                    v106 = 851
                                                    if false {
                                                        v108 := v3[680]
                                                        var v109 int32 = 42
                                                        if v108 {
                                                            if v108 {
                                                                var v110 uint16 = (v106 >> (((v2 & v2)) & 63))
                                                                v102 += v102
                                                                _ = v110
                                                            }
                                                        } else {
                                                            if (-91 == -1) {
                                                                var v111 string = "foo bar"
                                                                var v112 uint32 = 4294967295
                                                                _ = v111
                                                                _ = v112
                                                            }
                                                            var v113 int64 = 283
                                                            v106 -= v106
                                                            _ = v113
                                                        }
                                                        v108 = true
                                                        if (fi95 != fi82) {
                                                            delete(v3, 214)
                                                        }
                                                        v2 = (v2 - (^(v2)))
                                                        _ = v108
                                                        _ = v109
                                                    }
                                                    _ = len(v1)
                                                }
                                                fmt.Println(v102)
                                            }
                                            _ = v105
                                            _ = v106
                                        } else {
                                            switch "foo bar" {
                                            case "hello":
                                                if (42 < 904) {
                                                    if false {
                                                        v2 -= (v2 - 454)
                                                    }
                                                    v102 = v102
                                                } else {
                                                    switch fi95 {
                                                    case 32768:
                                                        switch (("hello" + "a") + ("" + ("" + "a"))) {
                                                        case "":
                                                            v102 = 410
                                                        case "foo bar":
                                                            v114 := GenS9{F0: -32768}
                                                            _ = v114
                                                        default:
                                                            var v115 float64 = 0.0
                                                            _ = v115
                                                        }
                                                    default:
                                                        if (-2.25 == 0.0) {
                                                            var v116 int64 = 0
                                                            if true {
                                                                v102 += v102
                                                                fmt.Println(fi82, v116, v2)
                                                                fmt.Println(v2, v2, v102)
                                                            } else {
                                                                delete(v3, 87)
                                                                v3[4294967295] = (61 >= 127)
                                                            }
                                                            if (float64(1.5) >= 0.1) {
                                                                switch ("a" + "") {
                                                                case "hello":
                                                                    v116 += (-(v116))
                                                                    v117 := v3[0]
                                                                    v1 = append(v1, (fi82 % ((fi82) | 1)))
                                                                    _ = v117
                                                                case "a":
                                                                    fmt.Println(v2, v102, v102)
                                                                case "":
                                                                    v2 = 217
                                                                }
                                                            } else {
                                                                var v118 bool = (!false)
                                                                if (1000 == 32767) {
                                                                    _ = len(v1)
                                                                    var v119 int = fi95
                                                                    _ = v119
                                                                } else {
                                                                    v2 = v2
                                                                }
                                                                _ = v118
                                                            }
                                                            _ = v116
                                                        } else {
                                                            var v120 uint64 = 18446744073709551615
                                                            if true {
                                                                if true {
                                                                    v102 -= 4294967295
                                                                    v3[1] = (!(!(-888 > -1)))
                                                                    if ((v120 - 18446744073709551615) >= 2) {
                                                                        if (v2 <= v2) {
                                                                            v121 := GenS8{F0: 1, F1: 0, F2: (^(v102))}
                                                                            var v122 int8 = -128
                                                                            v123 := GenS11{F0: fi82}
                                                                            _ = v121
                                                                            _ = v122
                                                                            _ = v123
                                                                        } else {
                                                                            v3[2] = ((v2 << ((0) & 63)) == 1)
                                                                        }
                                                                        fmt.Println(v102)
                                                                        v124 := GenS12{F0: 188, F1: 1}
                                                                        _ = v124
                                                                    } else {
                                                                        _ = len(v1)
                                                                        v125 := GenS13{F0: 0, F1: 9223372036854775807}
                                                                        if len(v1) > 0 {
                                                                            v1[0] = (fi95 << (((^(v2))) & 63))
                                                                        }
                                                                        _ = v125
                                                                    }
                                                                } else {
                                                                    var v126 int16 = 0
                                                                    v1 = append(v1, (fi95 & 32768))
                                                                    _ = v126
                                                                }
                                                                v2 = (v2 / ((42) | 1))
                                                                var v127 float32 = (float32(-2.25) / float32(1.5))
                                                                _ = v127
                                                            }
                                                            var v128 int16 = 0
                                                            _ = v120
                                                            _ = v128
                                                        }
                                                    }
                                                    v129 := GenS14{F0: 42, F1: 65535}
                                                    switch (fi82 << ((42) & 63)) {
                                                    case 2147483647:
                                                        v2 -= v2
                                                    case -258:
                                                        v130 := GenS5{F0: fi95, F1: 118}
                                                        _ = v130
                                                    case 0:
                                                        v102 = (v102 & (v102 >> ((v2) & 63)))
                                                    }
                                                    _ = v129
                                                }
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = len(v3)
                                        }
                                        _ = v104
                                    case 490:
                                        fmt.Println(v2, v102)
                                    case 469:
                                        fmt.Println(fi82, fi95, v2)
                                    default:
                                        var v131 int8 = 0
                                        _ = v131
                                    }
                                }
                            default:
                                fmt.Println(fi82, v102)
                            }
                            _ = v102
                        }
                    case "a":
                        switch (("hello" + (("foo bar" + "hello") + "foo bar")) + ("\n\t" + "a")) {
                        case "hello":
                            switch -564 {
                            case -1:
                                fmt.Println(fi82)
                            default:
                                if (true || ((-738 < 1) && (!(-2.25 >= ((0.1 - ((1.5 / (float64(-2.25) + 3.14159)) * 0.1)) * (float64(-2.25) + (3.14159 / 1.5))))))) {
                                    switch int(v2) {
                                    case -2147483648:
                                        var v132 int32 = 0
                                        v133 := GenS15{F0: 42, F1: (-(v132)), F2: uint32(v2)}
                                        _ = v132
                                        _ = v133
                                    case 145:
                                        switch "foo bar" {
                                        case "foo bar":
                                            v134 := GenS16{F0: -849, F1: v2, F2: (v2 + (v2 >> ((v2) & 63)))}
                                            _ = v134
                                        default:
                                            switch (-(fi82)) {
                                            case 991:
                                                if (true || true) {
                                                    v135 := GenS8{F0: 42, F1: 255, F2: 42}
                                                    var v136 float64 = (((0.0 * ((1e10 - float64(1e10)) - (-2.25 / (float64(1.5) * 0.1)))) / -2.25) + -2.25)
                                                    if len(v1) > 0 {
                                                        v1[0] = (fi95 & (fi95 / ((716) | 1)))
                                                    }
                                                    _ = v135
                                                    _ = v136
                                                } else {
                                                    v2 = (v2 << (((v2 + 1)) & 63))
                                                }
                                            }
                                        }
                                    case 4294967295:
                                        var v137 int8 = 121
                                        _ = v137
                                    default:
                                        v2 -= 1
                                    }
                                    v138 := GenS3{F0: 255, F1: (v2 * 42)}
                                    var v139 string = ("hello" + ("\n\t" + "\n\t"))
                                    _ = v138
                                    _ = v139
                                }
                            }
                        default:
                            if true {
                                if (2147483647 != 544) {
                                    switch "\n\t" {
                                    case "a":
                                        var v140 int64 = 0
                                        if len(v1) > 0 {
                                            v1[0] = (-(fi95))
                                        }
                                        var v141 bool = ((fi95 + -407) >= (fi95 ^ (fi95 ^ fi82)))
                                        _ = v140
                                        _ = v141
                                    case "hello":
                                        if false {
                                            v2 = (v2 ^ v2)
                                        }
                                    default:
                                        var v142 uint32 = 485
                                        _ = v142
                                    }
                                }
                            }
                        }
                    }
                }
                v2 = (v2 / ((931) | 1))
                _ = fi95
            }
            _ = fi82
        }
        v143 := GenS17{F0: 1, F1: (v2 ^ (v2 * (v2 / (((v2 << ((v2) & 63))) | 1))))}
        _ = v143
    }
    for fi144 := 0; fi144 < 11; fi144++ {
        v2 = 340
        _ = fi144
    }
    v2 += 0
    v2 += 4294967295
    v145 := GenS2{F0: 1, F1: 2, F2: -2147483648}
    for fi146 := 0; fi146 < 4; fi146++ {
        var v147 uint = 42
        _ = len(v1)
        v147 += v147
        _ = fi146
        _ = v147
    }
    if true {
        switch "a" {
        case "a":
            v148 := GenS7{F0: 1, F1: 65535, F2: v2}
            _ = v148
        case "\n\t":
            v149 := GenS1{F0: 0, F1: -9223372036854775808, F2: "foo bar"}
            _ = v149
        case "foo bar":
            v2 = v2
        default:
            switch (("\n\t" + (((("" + ("a" + (((("\n\t" + ("a" + ("" + "\n\t"))) + "\n\t") + "a") + ("hello" + "")))) + "a") + "") + (("hello" + "\n\t") + "hello"))) + "foo bar") {
            case "hello":
                if false {
                    var v150 float32 = -2.25
                    v151 := GenS4{F0: -88}
                    v152 := GenS18{F0: false}
                    _ = v150
                    _ = v151
                    _ = v152
                } else {
                    v2 = v2
                    var v153 int = -9223372036854775808
                    v2 = 727
                    _ = v153
                }
            }
        }
        for fi154 := 0; fi154 < 7; fi154++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v1 = append(v1, (-(fi154)))
            if false {
                for fi155 := 0; fi155 < 13; fi155++ {
                    if (v2 > v2) {
                        v3[2147483648] = ("foo bar" < ("" + ""))
                        if (((("\n\t" + "\n\t") + "hello") + "\n\t") <= "a") {
                            var v156 uint8 = 1
                            v157 := GenS16{F0: 1, F1: (v2 >> ((v2) & 63)), F2: v2}
                            _ = v156
                            _ = v157
                        }
                        if true {
                            var v158 uint = 42
                            if len(v1) > 0 {
                                v1[0] = (fi155 % ((fi155) | 1))
                            }
                            _ = v158
                        }
                    }
                    var v159 float64 = ((((float64(((0.0 / 1.5) / -2.25)) / 1.5) - 1e10) + float64((1.5 * 0.0))) + -2.25)
                    _ = fi155
                    _ = v159
                }
            }
            _ = fi154
        }
    } else {
        fmt.Println(v2)
        if false {
            var v160 uint16 = 42
            v161 := GenS16{F0: 0, F1: (v2 + v2), F2: (v2 & 212)}
            delete(v3, 4294967295)
            _ = v160
            _ = v161
        } else {
            var v162 float64 = 1e10
            _ = v162
        }
    }
    fmt.Println(v2)
    var v163 uint64 = 42
    v2 = v2
    fmt.Println(v2, v2)
    v164 := v3[291]
    v165 := GenS19{F0: 0, F1: -1, F2: 0}
    fmt.Println(v2, v164, v163)
    var v166 float64 = 1.5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v163:", v163)
    fmt.Println("v164:", v164)
    fmt.Println("v166:", v166)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v145
    _ = v163
    _ = v164
    _ = v165
    _ = v166
    _ = v2
    _ = v3
}
