package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 uint64; F1 int64; F2 uint }
type GenS6 struct { F0 int32; F1 string; F2 uint32 }
type GenS7 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS8 struct { F0 int64; F1 int }
type GenS9 struct { F0 int32; F1 uint16 }

func main() {
    var v1 float64 = 1e10
    var v2 int = 32767
    v3 := make(map[int]string)
    var v4 float64 = v1
    v5 := make(map[string]uint32)
    var v6 float32 = 1.5
    v7 := GenS1{F0: 673, F1: 1, F2: 0}
    fmt.Println(v1, v1, v6)
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 int16 = 0
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v11 := []int{v2}
    if len(v11) > 0 {
        v11[0] = 32767
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        v1 += v4
        _ = fi12
    }
    v3[v2] = ("hello" + "")
    for fi13 := 0; fi13 < 6; fi13++ {
        if (!(2147483647 != 168)) {
            var v14 float32 = float32(float64(v6))
            v15 := GenS1{F0: 4294967295, F1: 0, F2: -537}
            v14 = (float32(v4) + 3.14159)
            _ = v14
            _ = v15
        }
        _ = fi13
    }
    v6 += (v6 * v6)
    if (!(!(32767 == 32767))) {
        var v16 float64 = (v1 * v4)
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        v6 += float32((v1 - float64(float32((0.0 / (float64(1.5) * (v4 * v1)))))))
        _ = v16
    } else {
        for fi17 := 0; fi17 < 1; fi17++ {
            v6 -= ((0.0 + v6) + (((((1e10 + ((v6 / 1.5) / v6)) * 3.14159) - v6) * v6) - float32(3.14159)))
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            delete(v3, v2)
            _ = fi17
        }
        v6 = ((float32(0.0) + v6) - (0.1 / v6))
        switch (v2 - -2147483648) {
        case 2147483647:
            _ = len(v11)
        case 4294967295:
            var v18 int32 = 1
            _ = v18
        case 32768:
            if false {
                switch (v2 * 491) {
                case -1:
                    fmt.Println(v1, v4)
                case 9223372036854775807:
                    if false {
                        v19 := GenS1{F0: 649, F1: 42, F2: 1}
                        var v20 bool = (-9223372036854775808 >= -1)
                        v20 = v20
                        _ = v19
                        _ = v20
                    } else {
                        v4 = float64(-2.25)
                        v6 -= v6
                        fmt.Println(v6)
                    }
                case 32767:
                    for fi21 := 0; fi21 < 5; fi21++ {
                        var v22 float64 = (v4 + (0.1 * (v4 / float64((1.5 * 0.1)))))
                        _ = fi21
                        _ = v22
                    }
                default:
                    fmt.Println(v4)
                }
                var v23 int16 = int16(v2)
                var v24 float64 = (3.14159 + v4)
                _ = v23
                _ = v24
            } else {
                v2 -= -616
                v5[("" + "")] = 843
            }
        default:
            for fi25 := 0; fi25 < 3; fi25++ {
                var v26 uint32 = 2
                var v27 uint32 = v26
                _ = fi25
                _ = v26
                _ = v27
            }
        }
    }
    switch (^(v2)) {
    case 32767:
        _ = len(v11)
    }
    switch "" {
    case "a":
        v1 -= v1
    case "foo bar":
        switch v2 {
        case 9223372036854775807:
            for fi28 := 0; fi28 < 12; fi28++ {
                for fi29 := 0; fi29 < 11; fi29++ {
                    v2 -= fi29
                    _ = fi29
                }
                _ = fi28
            }
        }
    case "\n\t":
        var v30 uint16 = 993
        _ = v30
    default:
        switch -1 {
        case -1:
            if false {
                for fi31 := 0; fi31 < 2; fi31++ {
                    if (!false) {
                        switch (^(fi31)) {
                        case 1:
                            v32 := v5[(("" + "\n\t") + ("hello" + ("\n\t" + "hello")))]
                            _ = v32
                        }
                        switch (-(fi31)) {
                        case -818:
                            v5["foo bar"] = 2147483647
                        default:
                            v1 = (float64(v6) + v1)
                        }
                    } else {
                        if false {
                            if ("a" != "") {
                                v6 += (v6 / (v6 / v6))
                                fmt.Println(fi31, v4)
                                var v33 float64 = v1
                                _ = v33
                            }
                            if (false || (false && true)) {
                                if false {
                                    for fi34 := 0; fi34 < 4; fi34++ {
                                        v2 -= v2
                                        fmt.Println(fi31)
                                        switch ("a" + "\n\t") {
                                        case "":
                                            switch fi31 {
                                            case -453:
                                                v35 := GenS2{F0: 2147483648, F1: 0, F2: 1}
                                                _ = v35
                                            }
                                        case "\n\t":
                                            fmt.Println(fi31, v4, v2)
                                        case "a":
                                            var v36 int32 = -1
                                            _ = v36
                                        }
                                        _ = fi34
                                    }
                                } else {
                                    var v37 int64 = -351
                                    _ = v37
                                }
                            } else {
                                for fi38 := 0; fi38 < 17; fi38++ {
                                    var v39 uint = 4294967295
                                    if (uint32(v39) > 419) {
                                        v4 = v1
                                        if len(v11) > 0 {
                                            v11[0] = (fi31 ^ 0)
                                        }
                                    } else {
                                        v5["foo bar"] = 2
                                        var v40 int32 = 1
                                        _ = v40
                                    }
                                    v41 := v3[fi31]
                                    _ = fi38
                                    _ = v39
                                    _ = v41
                                }
                                if true {
                                    v6 = 1e10
                                    for fi42 := 0; fi42 < 5; fi42++ {
                                        fmt.Println(fi42, v6)
                                        if true {
                                            _ = len(v11)
                                        }
                                        if true {
                                            var v43 string = "foo bar"
                                            var v44 int16 = 0
                                            _ = v43
                                            _ = v44
                                        }
                                        var v45 uint32 = 677
                                        _ = fi42
                                        _ = v45
                                    }
                                }
                                v2 -= (v2 ^ fi31)
                            }
                        } else {
                            if (-403 < 9223372036854775807) {
                                if true {
                                    v46 := GenS1{F0: 249, F1: 127, F2: -544}
                                    v5[("" + "")] = 2147483648
                                    var v47 uint8 = 117
                                    _ = v46
                                    _ = v47
                                }
                                v48 := GenS1{F0: 2, F1: -1, F2: 139}
                                for fi49 := 0; fi49 < 1; fi49++ {
                                    fmt.Println(v1, v1, v1)
                                    var v50 int = -9223372036854775808
                                    fmt.Println(v4)
                                    _ = fi49
                                    _ = v50
                                }
                                _ = v48
                            } else {
                                if ((v2 / (((v2 % ((v2) | 1))) | 1)) == (fi31 | v2)) {
                                    v51 := v3[fi31]
                                    fmt.Println(v2)
                                    _ = v51
                                }
                            }
                            var v52 string = "foo bar"
                            _ = v52
                        }
                    }
                    _ = len(v11)
                    var v53 uint = 42
                    _ = fi31
                    _ = v53
                }
                for fi54 := 0; fi54 < 2; fi54++ {
                    switch (v2 * v2) {
                    case 32767:
                        switch "\n\t" {
                        case "hello":
                            v55 := GenS2{F0: 816, F1: 127, F2: 0}
                            _ = v55
                        case "a":
                            if true {
                                for fi56 := 0; fi56 < 16; fi56++ {
                                    if len(v11) > 0 {
                                        v11[0] = (fi54 % (((v2 + fi54)) | 1))
                                    }
                                    var v57 float64 = float64(1.5)
                                    v57 = ((v1 / -2.25) + 1.5)
                                    v58 := GenS3{F0: -2147483648, F1: -9223372036854775808}
                                    var v59 int = -646
                                    _ = fi56
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                }
                                _ = len(v3)
                                for fi60 := 0; fi60 < 10; fi60++ {
                                    if true {
                                        if true {
                                            if true {
                                                switch fi60 {
                                                case 712:
                                                    _ = len(v11)
                                                case 1:
                                                    switch fi54 {
                                                    case -869:
                                                        fmt.Println(v1, fi54, v6)
                                                    case 127:
                                                        switch "\n\t" {
                                                        case "hello":
                                                            switch "a" {
                                                            case "a":
                                                                v1 = v1
                                                            }
                                                        case "a":
                                                            if (502 >= 0) {
                                                                delete(v3, (fi54 >> ((1) & 63)))
                                                                v1 = 0.1
                                                                var v61 uint32 = 2
                                                                _ = v61
                                                            }
                                                        default:
                                                            v3[(fi60 | v2)] = "\n\t"
                                                        }
                                                    case 9:
                                                        if false {
                                                            fmt.Println(v1, v6)
                                                            _ = len(v11)
                                                            var v62 uint = 382
                                                            fmt.Println(v4, v4)
                                                            _ = v62
                                                        } else {
                                                            var v63 uint8 = 127
                                                            _ = v63
                                                        }
                                                    default:
                                                        var v64 int16 = -1
                                                        _ = v64
                                                    }
                                                    var v65 float32 = float32(((float64((v6 / float32(float64(v6)))) - 3.14159) * 1.5))
                                                    _ = v65
                                                case -9223372036854775808:
                                                    fmt.Println(v2, v4)
                                                default:
                                                    var v66 int = 32767
                                                    _ = v66
                                                }
                                                v1 = float64(1.5)
                                            } else {
                                                v4 -= v1
                                                v6 = v6
                                            }
                                        } else {
                                            fmt.Println(fi60, v1, v6)
                                        }
                                        v67 := GenS4{F0: 42, F1: false}
                                        v2 = fi54
                                        _ = v67
                                    }
                                    fmt.Println(v2, v2)
                                    if true {
                                        v4 = v1
                                    } else {
                                        v2 -= (fi60 / ((9223372036854775807) | 1))
                                        if (2 != 42) {
                                            v4 += v1
                                        } else {
                                            v68 := GenS5{F0: 849, F1: -9223372036854775808, F2: 42}
                                            _ = v68
                                        }
                                    }
                                    _ = fi60
                                }
                            }
                        }
                    case -9223372036854775808:
                        if (-2.25 >= ((v4 - float64(v6)) - 0.0)) {
                            if true {
                                v69 := GenS6{F0: -748, F1: "\n\t", F2: 298}
                                for fi70 := 0; fi70 < 12; fi70++ {
                                    switch "a" {
                                    case "hello":
                                        var v71 uint32 = 42
                                        _ = v71
                                    case "":
                                        v1 = 3.14159
                                    }
                                    var v72 int8 = 1
                                    v73 := GenS6{F0: -1, F1: "hello", F2: 906}
                                    _ = fi70
                                    _ = v72
                                    _ = v73
                                }
                                _ = v69
                            }
                            if true {
                                fmt.Println(v6)
                            } else {
                                if len(v11) > 0 {
                                    v11[0] = -9223372036854775808
                                }
                                var v74 uint16 = 78
                                v6 -= v6
                                _ = v74
                            }
                            v75 := GenS2{F0: 950, F1: 2, F2: -1}
                            _ = v75
                        } else {
                            v4 = -2.25
                            for fi76 := 0; fi76 < 9; fi76++ {
                                v4 += (v1 + 0.0)
                                if (false && true) {
                                    var v77 string = ""
                                    _ = v77
                                } else {
                                    if true {
                                        v78 := GenS2{F0: 4294967295, F1: 128, F2: 1}
                                        if true {
                                            v79 := GenS5{F0: 2, F1: 9223372036854775807, F2: 4294967295}
                                            var v80 int32 = 676
                                            v80 = -981
                                            _ = v79
                                            _ = v80
                                        }
                                        _ = v78
                                    }
                                    var v81 int = -9223372036854775808
                                    switch (v81 >> ((18446744073709551615) & 63)) {
                                    case 865:
                                        if (!(0 <= 1)) {
                                            var v82 uint = 42
                                            _ = v82
                                        }
                                    case 4294967295:
                                        v2 -= (v2 + fi54)
                                    case 9223372036854775807:
                                        var v83 string = "\n\t"
                                        _ = v83
                                    default:
                                        switch ("foo bar" + "a") {
                                        case "a":
                                            v84 := v3[9223372036854775807]
                                            _ = v84
                                        case "":
                                            v85 := GenS7{F0: 42, F1: 42, F2: 584}
                                            _ = v85
                                        case "\n\t":
                                            v4 += (1.5 / v1)
                                        }
                                    }
                                    v86 := GenS7{F0: 42, F1: 127, F2: 297}
                                    _ = v81
                                    _ = v86
                                }
                                var v87 bool = true
                                v4 = v1
                                _ = fi76
                                _ = v87
                            }
                        }
                    case 127:
                        if false {
                            for fi88 := 0; fi88 < 17; fi88++ {
                                var v89 uint64 = 774
                                v2 = v2
                                _ = fi88
                                _ = v89
                            }
                        } else {
                            var v90 uint8 = 127
                            var v91 float64 = 1e10
                            switch ("a" + "") {
                            case "":
                                var v92 uint8 = 42
                                _ = v92
                            case "foo bar":
                                v93 := GenS5{F0: 829, F1: 42, F2: 807}
                                _ = v93
                            case "hello":
                                _ = len(v5)
                            default:
                                switch "\n\t" {
                                case "":
                                    if len(v11) > 0 {
                                        v11[0] = 911
                                    }
                                case "a":
                                    v91 = (((1e10 / v1) + 1.5) + (v1 + v91))
                                case "hello":
                                    var v94 bool = true
                                    _ = v94
                                }
                            }
                            _ = v90
                            _ = v91
                        }
                    default:
                        var v95 int16 = -1
                        _ = v95
                    }
                    v11 = append(v11, (fi54 << ((674) & 63)))
                    _ = fi54
                }
            }
        case -454:
            switch 127 {
            case 1:
                for fi96 := 0; fi96 < 4; fi96++ {
                    switch "hello" {
                    case "a":
                        v1 -= (v1 / ((v1 + v1) / (3.14159 - v4)))
                    case "":
                        if (v4 != 0.0) {
                            var v97 int32 = 0
                            v97 += v97
                            if false {
                                var v98 string = (("" + "a") + ("\n\t" + ("" + "")))
                                var v99 int = (fi96 & 0)
                                v100 := GenS5{F0: 353, F1: 1, F2: 4294967295}
                                _ = v100
                                _ = v98
                                _ = v99
                            }
                            _ = v97
                        } else {
                            v2 = v2
                            v2 = fi96
                            if false {
                                for fi101 := 0; fi101 < 9; fi101++ {
                                    var v102 int8 = 1
                                    if false {
                                        v6 = v6
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v103 uint = 1
                                            _ = v103
                                        default:
                                            v4 -= v1
                                        }
                                        v104 := GenS3{F0: -321, F1: 0}
                                        v105 := GenS2{F0: 2147483647, F1: 128, F2: (v102 << ((0) & 63))}
                                        _ = v104
                                        _ = v105
                                    } else {
                                        var v106 float64 = v1
                                        _ = v106
                                    }
                                    _ = fi101
                                    _ = v102
                                }
                                v4 = 0.1
                                v6 = v6
                            } else {
                                v5["\n\t"] = 441
                                v107 := GenS4{F0: 2, F1: (false && false)}
                                v108 := GenS7{F0: 213, F1: 0, F2: 0}
                                _ = v107
                                _ = v108
                            }
                        }
                    }
                    _ = fi96
                }
            case -356:
                v1 += ((1e10 / v1) - ((v4 * ((v4 + v1) * (v4 / 1.0))) / ((0.1 - 0.1) + v4)))
            }
        default:
            var v109 bool = (-9223372036854775808 > 9223372036854775807)
            _ = v109
        }
    }
    v110 := GenS1{F0: 875, F1: -128, F2: 9223372036854775807}
    if len(v11) > 0 {
        v11[0] = -115
    }
    v2 = (-(v2))
    switch ("a" + "hello") {
    case "a":
        switch (v2 * -639) {
        case -9223372036854775808:
            v2 = (-(v2))
        case 987:
            switch (v2 ^ (v2 - (v2 * v2))) {
            case -164:
                v111 := GenS3{F0: -2147483648, F1: 1}
                _ = v111
            default:
                v4 = float64((0.0 / (v6 - 3.14159)))
            }
        }
    }
    switch v2 {
    case 0:
        if ((2 < 170) || ((-1 <= 42) && (65535 < 255))) {
            var v112 int = (v2 & v2)
            if false {
                var v113 uint = 1
                fmt.Println(v4)
                _ = v113
            }
            for fi114 := 0; fi114 < 4; fi114++ {
                fmt.Println(v1, fi114, v112)
                switch fi114 {
                case 4294967295:
                    v115 := GenS8{F0: -1, F1: (v112 | (v2 / (((fi114 - fi114)) | 1)))}
                    _ = v115
                case 127:
                    v11 = append(v11, v2)
                }
                fmt.Println(v2)
                _ = fi114
            }
            _ = v112
        }
    case 28:
        fmt.Println(v4)
    default:
        if (false && (false && (!(((220 < 1000) && (!false)) && false)))) {
            for fi116 := 0; fi116 < 19; fi116++ {
                _ = len(v3)
                _ = fi116
            }
            var v117 uint64 = 755
            _ = v117
        } else {
            delete(v3, (^(v2)))
        }
    }
    v118 := GenS9{F0: 1, F1: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v110
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
