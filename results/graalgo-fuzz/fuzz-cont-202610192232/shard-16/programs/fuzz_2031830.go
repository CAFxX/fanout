package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32; F1 string; F2 bool }
type GenS6 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS7 struct { F0 int8; F1 uint8; F2 string }
type GenS8 struct { F0 int16; F1 int; F2 int16 }
type GenS9 struct { F0 int; F1 bool; F2 int32 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint64; F1 string; F2 uint64 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint8; F1 uint; F2 uint }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 int; F1 int }
type GenS16 struct { F0 string; F1 int }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 uint }
type GenS19 struct { F0 uint; F1 string }
type GenS20 struct { F0 uint16; F1 uint64 }
type GenS21 struct { F0 string }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    var v3 uint16 = 65535
    v4 := make(map[uint32]int)
    v5 := GenS1{F0: 255, F1: -652, F2: 0}
    if true {
        v3 = v3
        var v6 float64 = -2.25
        _ = v6
    }
    v3 -= (v3 / (((^(v3))) | 1))
    switch ("hello" + ("a" + "\n\t")) {
    case "hello":
        var v7 float32 = 0.0
        _ = v7
    case "foo bar":
        v8 := []int{}
        _ = v8
    case "a":
        v3 = v3
    }
    fmt.Println(v3)
    if (!true) {
        var v9 int8 = 54
        _ = v9
    }
    switch 127 {
    case 1:
        v10 := GenS2{F0: 11, F1: -616, F2: 1}
        _ = v10
    case -244:
        if (-2147483648 > 0) {
            v11 := GenS2{F0: 42, F1: 42, F2: 18446744073709551615}
            if true {
                v3 -= v3
                v12 := []int{128}
                v3 = v3
                _ = v12
            }
            _ = v11
        }
    case 9223372036854775807:
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 uint8 = 49
            var v15 bool = (("foo bar" == "\n\t") && false)
            if ("\n\t" != ("foo bar" + "foo bar")) {
                var v16 uint64 = 0
                var v17 bool = (float32(1.5) != (1.5 - 0.1))
                _ = v16
                _ = v17
            } else {
                for fi18 := 0; fi18 < 4; fi18++ {
                    var v19 float32 = (1e10 / 1.0)
                    v4[0] = (-(fi13))
                    if v15 {
                        v20 := []int{(fi18 ^ (fi13 * -1)), fi18}
                        if (v3 == 32767) {
                            if (636 == uint64(v14)) {
                                v3 += (v3 * v3)
                                fmt.Println(v3, v3, v15)
                                v21 := GenS1{F0: uint(v3), F1: fi13, F2: 4}
                                _ = v21
                            } else {
                                switch ("\n\t" + "foo bar") {
                                case "":
                                    if v15 {
                                        if (("\n\t" + ("" + "hello")) < ("hello" + "")) {
                                            fmt.Println(fi13, v3, v19)
                                            if false {
                                                if (9223372036854775807 < 42) {
                                                    if v15 {
                                                        var v22 int32 = 0
                                                        var v23 string = "\n\t"
                                                        if v15 {
                                                            var v24 int64 = 42
                                                            switch v23 {
                                                            case "":
                                                                var v25 uint64 = 276
                                                                _ = v25
                                                            case "a":
                                                                var v26 int8 = 0
                                                                _ = v26
                                                            case "hello":
                                                                fmt.Println(v23, v14)
                                                            default:
                                                                v27 := GenS3{F0: -1}
                                                                _ = v27
                                                            }
                                                            v23 = v23
                                                            _ = v24
                                                        }
                                                        _ = v22
                                                        _ = v23
                                                    }
                                                }
                                                if false {
                                                    v14 = 255
                                                    v19 = v19
                                                }
                                                v28 := GenS1{F0: 2, F1: (fi13 << ((25) & 63)), F2: 0}
                                                _ = v28
                                            }
                                            switch ("foo bar" + "foo bar") {
                                            case "a":
                                                v29 := GenS3{F0: 1000}
                                                _ = v29
                                            case "":
                                                fmt.Println(v15, v14, fi13)
                                            }
                                        } else {
                                            v30 := GenS1{F0: 770, F1: fi18, F2: 394}
                                            _ = v30
                                        }
                                        if (0.1 >= float32((-2.25 + 3.14159))) {
                                            if (!v15) {
                                                v31 := GenS3{F0: 1000}
                                                _ = v31
                                            } else {
                                                switch (^(fi18)) {
                                                case 127:
                                                    v32 := GenS3{F0: 32767}
                                                    var v33 uint64 = 0
                                                    _ = v32
                                                    _ = v33
                                                case -402:
                                                    v20 = append(v20, fi18)
                                                default:
                                                    switch ("\n\t" + "foo bar") {
                                                    case "foo bar":
                                                        var v34 int64 = 1
                                                        _ = v34
                                                    default:
                                                        v35 := GenS4{F0: v3}
                                                        _ = v35
                                                    }
                                                }
                                                v15 = (!(!(v15 || v15)))
                                                if false {
                                                    v19 = v19
                                                    v3 += v3
                                                    var v36 int32 = 1
                                                    if (-96 != 127) {
                                                        switch 2147483647 {
                                                        case 0:
                                                            var v37 int8 = -84
                                                            _ = v37
                                                        case 2147483647:
                                                            switch fi13 {
                                                            case -77:
                                                                if v15 {
                                                                    v36 = 42
                                                                    if (v14 >= (^(v14))) {
                                                                        var v38 uint8 = v14
                                                                        var v39 uint64 = 75
                                                                        var v40 bool = (32767 == 0)
                                                                        _ = v38
                                                                        _ = v39
                                                                        _ = v40
                                                                    }
                                                                    var v41 uint = 255
                                                                    _ = v41
                                                                } else {
                                                                    v42 := GenS2{F0: 2, F1: -443, F2: 0}
                                                                    v19 -= ((v19 * v19) * float32(3.14159))
                                                                    if len(v20) >= 2 {
                                                                        v20 = v20[1:len(v20)]
                                                                    }
                                                                    var v43 uint16 = v3
                                                                    if v15 {
                                                                        if len(v20) >= 2 {
                                                                            v20 = v20[1:len(v20)]
                                                                        }
                                                                        var v44 int64 = 42
                                                                        _ = v44
                                                                    }
                                                                    _ = v42
                                                                    _ = v43
                                                                }
                                                            case 2147483647:
                                                                var v45 float32 = 3.14159
                                                                _ = v45
                                                            case -2147483648:
                                                                if (v3 < 32767) {
                                                                    v46 := GenS1{F0: 255, F1: fi18, F2: 2147483648}
                                                                    var v47 int8 = 1
                                                                    v48 := GenS5{F0: 495, F1: (("\n\t" + (("" + "a") + ("\n\t" + ("a" + "")))) + ("foo bar" + ("foo bar" + "foo bar"))), F2: v15}
                                                                    _ = v46
                                                                    _ = v47
                                                                    _ = v48
                                                                } else {
                                                                    fmt.Println(v3, fi13)
                                                                    switch 32768 {
                                                                    case -669:
                                                                        delete(v1, 9223372036854775807)
                                                                    case 2147483647:
                                                                        fmt.Println(v36, v3, fi18)
                                                                    }
                                                                    if v15 {
                                                                        v19 = 0.1
                                                                        v15 = false
                                                                    }
                                                                }
                                                            }
                                                        case -977:
                                                            if len(v20) > 0 {
                                                                v20[0] = (fi18 * int(v14))
                                                            }
                                                        default:
                                                            fmt.Println(fi18, v19, v15)
                                                        }
                                                        delete(v4, 2147483648)
                                                        switch (fi18 + fi18) {
                                                        case -996:
                                                            v19 += 0.0
                                                        case 2147483647:
                                                            v3 = v3
                                                        }
                                                    } else {
                                                        v3 = (v3 - (^(v3)))
                                                        fmt.Println(fi13, v15)
                                                        var v49 int = 2147483647
                                                        _ = v49
                                                    }
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        switch fi18 {
                                                        case 491:
                                                            var v50 uint = 255
                                                            _ = v50
                                                        case 99:
                                                            switch "" {
                                                            case "foo bar":
                                                                switch (^(fi18)) {
                                                                case -1:
                                                                    fmt.Println(v36, v15)
                                                                default:
                                                                    switch (("hello" + "foo bar") + (("\n\t" + ("" + (("" + (("\n\t" + ("foo bar" + "hello")) + ("" + "a"))) + "hello"))) + "\n\t")) {
                                                                    case "\n\t":
                                                                        v20 = append(v20, (^(fi18)))
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                v51 := GenS2{F0: 4294967295, F1: 1, F2: 42}
                                                                if (-1 <= -160) {
                                                                    v52 := GenS2{F0: 535, F1: -9223372036854775808, F2: 852}
                                                                    var v53 uint16 = v3
                                                                    _ = v52
                                                                    _ = v53
                                                                }
                                                                v3 = v3
                                                                _ = v51
                                                            default:
                                                                v15 = (false && (594 < 697))
                                                            }
                                                        case 32767:
                                                            var v54 int = fi13
                                                            _ = v54
                                                        }
                                                    }
                                                    v19 += v19
                                                    v14 = v14
                                                    _ = v36
                                                }
                                            }
                                            v55 := GenS3{F0: -4}
                                            _ = v55
                                        } else {
                                            var v56 float64 = 0.0
                                            if v15 {
                                                v19 -= -2.25
                                            }
                                            _ = v56
                                        }
                                    } else {
                                        switch ("\n\t" + ("hello" + ("a" + ("a" + ("\n\t" + ("" + ("foo bar" + ("\n\t" + "hello")))))))) {
                                        case "foo bar":
                                            v57 := GenS2{F0: 0, F1: 42, F2: 0}
                                            _ = v57
                                        case "":
                                            v58 := GenS6{F0: (v3 * v3), F1: 0, F2: v14}
                                            _ = v58
                                        case "hello":
                                            if len(v20) >= 2 {
                                                v20 = v20[1:len(v20)]
                                            }
                                        }
                                        v59 := GenS3{F0: -32768}
                                        switch (fi13 ^ fi18) {
                                        case 9223372036854775807:
                                            switch fi18 {
                                            case 32768:
                                                v3 += (^(v3))
                                            case -9223372036854775808:
                                                delete(v2, -1)
                                            default:
                                                if v15 {
                                                    fmt.Println(v19, v15)
                                                    var v60 int64 = 42
                                                    _ = v60
                                                }
                                            }
                                        case 370:
                                            if false {
                                                v3 += v3
                                                var v61 uint16 = v3
                                                var v62 uint = 1
                                                _ = v61
                                                _ = v62
                                            }
                                        case 0:
                                            var v63 bool = ((!(v3 == v3)) || false)
                                            _ = v63
                                        }
                                        _ = len(v20)
                                        v64 := GenS7{F0: -54, F1: 194, F2: "foo bar"}
                                        _ = v59
                                        _ = v64
                                    }
                                case "foo bar":
                                    switch "a" {
                                    case "\n\t":
                                        v3 += 515
                                    case "hello":
                                        switch "\n\t" {
                                        case "":
                                            if false {
                                                fmt.Println(v3)
                                                var v65 int64 = -212
                                                var v66 float32 = v19
                                                _ = v65
                                                _ = v66
                                            }
                                        case "hello":
                                            v15 = v15
                                        case "\n\t":
                                            fmt.Println(v19, v15, v19)
                                        default:
                                            switch (("a" + "") + "\n\t") {
                                            case "\n\t":
                                                switch (("" + "") + ("a" + ("hello" + ""))) {
                                                case "\n\t":
                                                    fmt.Println(v19, v19, v3)
                                                    v15 = (42 != 4294967295)
                                                    var v67 uint16 = v3
                                                    _ = v67
                                                }
                                            case "foo bar":
                                                v68 := GenS8{F0: 19, F1: fi13, F2: -32768}
                                                _ = v68
                                            default:
                                                v69 := v2[(fi13 >> ((255) & 63))]
                                                _ = v69
                                            }
                                        }
                                    }
                                }
                                if (3.14159 < (3.14159 * -2.25)) {
                                    if v15 {
                                        v70 := GenS6{F0: (v3 + v3), F1: 156, F2: v14}
                                        fmt.Println(fi13, v15)
                                        _ = v70
                                    } else {
                                        fmt.Println(v15)
                                        fmt.Println(v19, fi13)
                                    }
                                }
                                v71 := GenS2{F0: 4294967295, F1: -442, F2: 1}
                                switch fi18 {
                                case -622:
                                    v14 += v14
                                case -2147483648:
                                    var v72 uint8 = (v14 - 128)
                                    switch "hello" {
                                    case "":
                                        v3 += uint16(v14)
                                    }
                                    if true {
                                        if ((!v15) || (2 != v72)) {
                                            if true {
                                                switch "hello" {
                                                case "":
                                                    switch "a" {
                                                    case "\n\t":
                                                        var v73 float64 = 0.0
                                                        _ = v73
                                                    }
                                                case "hello":
                                                    v19 = v19
                                                case "a":
                                                    switch (("a" + "foo bar") + "a") {
                                                    case "a":
                                                        v74 := GenS9{F0: (fi13 ^ (^(fi18))), F1: (v15 && v15), F2: 0}
                                                        _ = v74
                                                    case "":
                                                        v15 = v15
                                                    case "hello":
                                                        if (!v15) {
                                                            var v75 uint = 4294967295
                                                            v76 := GenS1{F0: (^(v75)), F1: -889, F2: 758}
                                                            switch "hello" {
                                                            case "\n\t":
                                                                v15 = v15
                                                            case "":
                                                                fmt.Println(v3, fi13, v15)
                                                            default:
                                                                var v77 uint8 = (^(v72))
                                                                _ = v77
                                                            }
                                                            var v78 float32 = (((v19 + float32(float64((((v19 * 1.5) - v19) / v19)))) + v19) * v19)
                                                            var v79 uint16 = v3
                                                            _ = v75
                                                            _ = v76
                                                            _ = v78
                                                            _ = v79
                                                        } else {
                                                            v80 := GenS10{F0: -747}
                                                            v81 := GenS4{F0: 32767}
                                                            v82 := GenS11{F0: 4294967295, F1: "a", F2: 478}
                                                            _ = v80
                                                            _ = v81
                                                            _ = v82
                                                        }
                                                    default:
                                                        fmt.Println(fi18, v72)
                                                    }
                                                }
                                                var v83 uint8 = 128
                                                _ = v83
                                            } else {
                                                switch 924 {
                                                case 9223372036854775807:
                                                    v14 = v14
                                                }
                                                v84 := GenS12{F0: 0}
                                                _ = v84
                                            }
                                        } else {
                                            v85 := GenS6{F0: v3, F1: 154, F2: 1}
                                            _ = v85
                                        }
                                    } else {
                                        var v86 uint16 = (v3 ^ v3)
                                        if false {
                                            if v15 {
                                                if (-2 < -120) {
                                                    var v87 uint64 = 1
                                                    switch fi18 {
                                                    case 32767:
                                                        v87 = (v87 ^ v87)
                                                    case -6:
                                                        var v88 uint64 = v87
                                                        _ = v88
                                                    default:
                                                        v89 := GenS7{F0: 0, F1: 145, F2: "a"}
                                                        _ = v89
                                                    }
                                                    v86 -= 255
                                                    _ = v87
                                                }
                                                v3 = v86
                                            }
                                            switch "a" {
                                            case "":
                                                v90 := GenS13{F0: 150, F1: 0, F2: 4294967295}
                                                _ = v90
                                            case "foo bar":
                                                var v91 uint32 = 288
                                                _ = v91
                                            case "hello":
                                                v19 += 3.14159
                                            default:
                                                v92 := GenS10{F0: fi18}
                                                _ = v92
                                            }
                                            switch 760 {
                                            case 2147483647:
                                                delete(v4, 570)
                                            case 0:
                                                var v93 bool = v15
                                                v94 := GenS10{F0: (-(fi18))}
                                                fmt.Println(fi18, v72)
                                                _ = v93
                                                _ = v94
                                            case 1:
                                                switch (((("foo bar" + (((("hello" + "foo bar") + (("a" + "") + "\n\t")) + "foo bar") + ("foo bar" + ("" + (("a" + ("a" + ("\n\t" + "foo bar"))) + "foo bar"))))) + "foo bar") + "a") + ("" + "")) {
                                                case "\n\t":
                                                    fmt.Println(v86, fi13)
                                                default:
                                                    v95 := GenS14{F0: 50}
                                                    _ = v95
                                                }
                                            }
                                            fmt.Println(v14)
                                        }
                                        _ = v86
                                    }
                                    if (v14 <= v72) {
                                        fmt.Println(v3, v72)
                                        if (-1 <= int8(v14)) {
                                            if (0 == 32767) {
                                                if v15 {
                                                    v15 = (2 >= 483)
                                                }
                                                v96 := GenS15{F0: 42, F1: 32768}
                                                var v97 uint16 = 255
                                                _ = v96
                                                _ = v97
                                            } else {
                                                if (!(true && v15)) {
                                                    var v98 uint8 = (^(v14))
                                                    v99 := GenS4{F0: v3}
                                                    _ = v98
                                                    _ = v99
                                                }
                                                if len(v20) > 0 {
                                                    v20[0] = (^(fi18))
                                                }
                                                var v100 uint8 = (v14 & 127)
                                                _ = v100
                                            }
                                            v101 := GenS16{F0: ("a" + "\n\t"), F1: (fi18 / ((870) | 1))}
                                            _ = len(v1)
                                            _ = v101
                                        } else {
                                            if v15 {
                                                fmt.Println(v19)
                                                if false {
                                                    var v102 uint16 = v3
                                                    v3 = v3
                                                    v19 -= 0.1
                                                    _ = v102
                                                } else {
                                                    var v103 bool = (!v15)
                                                    _ = v103
                                                }
                                                if v15 {
                                                    v104 := GenS17{F0: 2}
                                                    v105 := GenS18{F0: 74}
                                                    _ = v104
                                                    _ = v105
                                                }
                                            }
                                            switch fi13 {
                                            case 0:
                                                var v106 uint16 = 1
                                                _ = v106
                                            case 885:
                                                v3 = v3
                                            case 2147483647:
                                                v107 := v1[fi18]
                                                _ = v107
                                            default:
                                                v14 += 0
                                            }
                                            v14 += (^(v72))
                                        }
                                        switch "foo bar" {
                                        case "foo bar":
                                            v108 := GenS9{F0: (fi18 + fi18), F1: v15, F2: 0}
                                            _ = v108
                                        case "a":
                                            var v109 uint64 = 762
                                            var v110 int8 = 42
                                            if ("hello" <= "hello") {
                                                v111 := GenS19{F0: 0, F1: ""}
                                                var v112 uint = 18446744073709551615
                                                v110 = v110
                                                _ = v111
                                                _ = v112
                                            }
                                            var v113 bool = false
                                            delete(v4, 1)
                                            v113 = v113
                                            _ = v109
                                            _ = v110
                                            _ = v113
                                        }
                                        v14 = v14
                                        v72 -= (v72 | (^(v72)))
                                    } else {
                                        fmt.Println(v19, fi13)
                                        v3 += v3
                                        switch 688 {
                                        case 9223372036854775807:
                                            if v15 {
                                                v72 -= 2
                                                var v114 uint16 = v3
                                                if v15 {
                                                    fmt.Println(v3, v72)
                                                    var v115 bool = v15
                                                    _ = v115
                                                } else {
                                                    v19 = v19
                                                    v116 := GenS20{F0: 65535, F1: 4294967295}
                                                    var v117 uint64 = 1
                                                    _ = v116
                                                    _ = v117
                                                }
                                                _ = v114
                                            } else {
                                                var v118 bool = (438 == 861)
                                                v118 = (v118 && (-128 == 90))
                                                switch (fi13 + fi13) {
                                                case 127:
                                                    var v119 int8 = -15
                                                    _ = v119
                                                default:
                                                    switch ("a" + "foo bar") {
                                                    case "a":
                                                        if (!v15) {
                                                            _ = len(v1)
                                                            v118 = (false || v118)
                                                            if (!v15) {
                                                                v72 = v72
                                                            } else {
                                                                var v120 uint64 = 4294967295
                                                                if v15 {
                                                                    if len(v20) > 0 {
                                                                        v20[0] = fi18
                                                                    }
                                                                    if (1.5 > (3.14159 / float64((v19 / v19)))) {
                                                                        v15 = (fi18 != 127)
                                                                    } else {
                                                                        var v121 uint32 = 65535
                                                                        _ = v121
                                                                    }
                                                                    v122 := GenS8{F0: 571, F1: fi18, F2: -974}
                                                                    _ = v122
                                                                }
                                                                v15 = v15
                                                                _ = v120
                                                            }
                                                            var v123 uint64 = uint64(v3)
                                                            if len(v20) > 0 {
                                                                v20[0] = fi18
                                                            }
                                                            if (954 > (v123 << ((255) & 63))) {
                                                                _ = len(v20)
                                                                var v124 int = fi13
                                                                _ = len(v20)
                                                                _ = v124
                                                            }
                                                            _ = v123
                                                        } else {
                                                            var v125 uint8 = v72
                                                            v14 = (v14 >> ((0) & 63))
                                                            _ = v125
                                                        }
                                                    case "\n\t":
                                                        var v126 uint8 = v72
                                                        _ = v126
                                                    case "foo bar":
                                                        v2[(fi13 & (-(fi13)))] = (fi18 / ((2147483647) | 1))
                                                    }
                                                }
                                                _ = v118
                                            }
                                        case 4294967295:
                                            if (9223372036854775807 != 0) {
                                                v127 := GenS15{F0: (fi13 + -252), F1: (fi18 << ((0) & 63))}
                                                v15 = (1 != int16(v3))
                                                var v128 int = 9223372036854775807
                                                _ = v127
                                                _ = v128
                                            }
                                        }
                                    }
                                    _ = v72
                                case 1:
                                    switch ("foo bar" + ("\n\t" + "")) {
                                    case "\n\t":
                                        v19 = v19
                                    case "":
                                        var v129 int16 = 365
                                        _ = v129
                                    default:
                                        v1[fi18] = 42
                                    }
                                }
                                _ = v71
                            }
                            v15 = (!v15)
                            var v130 bool = v15
                            switch (fi18 + 0) {
                            case 850:
                                if v130 {
                                    delete(v4, 4294967295)
                                    switch (fi18 - 2147483647) {
                                    case -9223372036854775808:
                                        if ((!v130) && true) {
                                            var v131 uint16 = v3
                                            fmt.Println(fi18)
                                            v15 = (-672 >= -529)
                                            _ = v131
                                        }
                                        v132 := GenS9{F0: fi18, F1: false, F2: 613}
                                        var v133 float64 = ((float64((float32(-2.25) / (0.0 - float32(float64((v19 + v19)))))) + float64((float32(-2.25) + v19))) / 1.5)
                                        _ = v132
                                        _ = v133
                                    case 0:
                                        var v134 int = fi13
                                        _ = v134
                                    case -2147483648:
                                        if len(v20) > 0 {
                                            v20[0] = fi13
                                        }
                                    default:
                                        var v135 uint64 = 1
                                        _ = v135
                                    }
                                    v19 = float32((3.14159 - 0.0))
                                } else {
                                    v14 = (v14 % (((v14 >> ((60) & 63))) | 1))
                                    v15 = v15
                                    switch fi13 {
                                    case 2147483647:
                                        var v136 int64 = -940
                                        _ = v136
                                    case 9223372036854775807:
                                        if v130 {
                                            if ((("\n\t" + "a") + "a") >= (("a" + "a") + "hello")) {
                                                v19 += v19
                                                if ((int16(v3) < 1000) || ((v14 - 42) < (v14 - 128))) {
                                                    var v137 uint16 = v3
                                                    v19 = ((3.14159 * v19) - v19)
                                                    _ = v137
                                                }
                                            }
                                        }
                                    default:
                                        v15 = ((fi18 == fi18) && ((1e10 - 0.1) != -2.25))
                                    }
                                    _ = len(v20)
                                }
                            case 401:
                                if true {
                                    v138 := GenS21{F0: ((("a" + "a") + "hello") + "")}
                                    _ = v138
                                }
                                fmt.Println(v3)
                                if len(v20) > 0 {
                                    v20[0] = -2147483648
                                }
                                fmt.Println(v3, v130)
                                fmt.Println(v15, v19)
                            case -1:
                                fmt.Println(v3, fi18)
                            default:
                                var v139 string = ""
                                var v140 float32 = ((v19 * (v19 + (1e10 - v19))) * (1.5 / v19))
                                v140 += 3.14159
                                _ = v139
                                _ = v140
                            }
                            v141 := GenS6{F0: (v3 << ((822) & 63)), F1: -1, F2: (^(v14))}
                            _ = v130
                            _ = v141
                        } else {
                            v14 = v14
                            fmt.Println(fi18, v15)
                        }
                        fmt.Println(v3)
                        v14 = v14
                        switch "\n\t" {
                        case "hello":
                            v3 = (^(v3))
                        }
                        _ = v20
                    }
                    _ = fi18
                    _ = v19
                }
                switch "a" {
                case "a":
                    var v142 uint32 = 42
                    _ = v142
                default:
                    v15 = (v15 && true)
                }
                v15 = (215 != 1)
            }
            _ = fi13
            _ = v14
            _ = v15
        }
    default:
        fmt.Println(v3)
    }
    if (42 < 408) {
        var v143 uint64 = 507
        v144 := GenS13{F0: 0, F1: 255, F2: 659}
        _ = v143
        _ = v144
    } else {
        fmt.Println(v3)
        if (!((-2.25 > 1e10) && (0.0 >= 0.0))) {
            v145 := GenS20{F0: (v3 >> ((42) & 63)), F1: 1}
            if (!(501 == 478)) {
                v4[348] = 9223372036854775807
                var v146 int16 = 580
                v146 -= v146
                _ = v146
            }
            var v147 uint = 255
            _ = v145
            _ = v147
        } else {
            v148 := []int{int(v3), 32767, 32768}
            _ = v148
        }
    }
    for fi149 := 0; fi149 < 16; fi149++ {
        fmt.Println(fi149)
        var v150 uint64 = 2
        _ = fi149
        _ = v150
    }
    if (true && (636 >= -1)) {
        switch 0 {
        case 4294967295:
            for fi151 := 0; fi151 < 10; fi151++ {
                v3 = v3
                _ = fi151
            }
        }
        v3 -= (v3 - v3)
    }
    v152 := GenS19{F0: 255, F1: ("foo bar" + ("foo bar" + ("a" + ("hello" + ""))))}
    var v153 uint = 275
    if true {
        v153 -= 436
    }
    fmt.Println(v153)
    if true {
        v3 = (v3 | v3)
        if (-54 > 54) {
            v3 = (^(v3))
        } else {
            v154 := v2[32767]
            v153 = v153
            _ = v154
        }
        v153 = v153
    }
    var v155 float32 = (0.1 + 0.0)
    v156 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v153:", v153)
    fmt.Println("v155:", v155)
    fmt.Println("slicelen v156:", len(v156), v156)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v152
    _ = v153
    _ = v155
    _ = v156
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
