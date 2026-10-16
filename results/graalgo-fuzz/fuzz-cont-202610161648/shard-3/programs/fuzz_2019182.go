package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS6 struct { F0 string; F1 uint32; F2 uint64 }
type GenS7 struct { F0 string; F1 uint64; F2 uint32 }
type GenS8 struct { F0 int32; F1 uint8; F2 int64 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 string; F1 int16 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint64; F1 uint32 }
type GenS14 struct { F0 uint16; F1 uint16 }
type GenS15 struct { F0 int }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 string; F1 int64 }
type GenS18 struct { F0 bool; F1 uint8 }
type GenS19 struct { F0 bool; F1 int8 }
type GenS20 struct { F0 string; F1 uint32 }
type GenS21 struct { F0 int8 }
type GenS22 struct { F0 uint32; F1 uint64; F2 uint }

func main() {
    var v1 int = 32767
    var v2 float64 = 3.14159
    var v3 int = (v1 ^ (v1 ^ (v1 + v1)))
    v4 := []int{274, -599, v3}
    var v5 int = v3
    switch "a" {
    case "a":
        for fi6 := 0; fi6 < 15; fi6++ {
            fmt.Println(v1, v2)
            v7 := GenS1{F0: 65535, F1: 42}
            if ((v2 > 3.14159) && (!(("hello" < "\n\t") || true))) {
                _ = len(v4)
                v2 -= v2
            } else {
                for fi8 := 0; fi8 < 5; fi8++ {
                    var v9 int16 = 1000
                    switch "" {
                    case "":
                        if true {
                            v9 = 0
                            v1 = (v3 >> ((65535) & 63))
                            var v10 bool = false
                            _ = v10
                        }
                    default:
                        switch "a" {
                        case "foo bar":
                            var v11 uint = 18446744073709551615
                            _ = v11
                        case "":
                            v5 -= (fi6 ^ v3)
                        case "hello":
                            if (false && true) {
                                fmt.Println(v9, fi6, v5)
                                if (((fi8 << ((42) & 63)) == v3) || false) {
                                    if true {
                                        var v12 uint8 = 128
                                        v13 := GenS2{F0: 1}
                                        _ = v12
                                        _ = v13
                                    } else {
                                        v14 := GenS1{F0: 32767, F1: 32767}
                                        _ = v14
                                    }
                                    v15 := GenS1{F0: 0, F1: 32767}
                                    v3 = v3
                                    _ = v15
                                }
                                v16 := GenS2{F0: 1}
                                _ = v16
                            } else {
                                var v17 int64 = 543
                                if ((false && (457 != 2)) || true) {
                                    v18 := GenS1{F0: 1, F1: 42}
                                    fmt.Println(v3, v17)
                                    _ = v18
                                } else {
                                    v19 := GenS2{F0: 42}
                                    v20 := GenS3{F0: -128, F1: "\n\t"}
                                    _ = v19
                                    _ = v20
                                }
                                var v21 float64 = (v2 * 3.14159)
                                _ = v17
                                _ = v21
                            }
                        default:
                            v1 = (v1 & v3)
                        }
                    }
                    if true {
                        if (-2147483648 >= 1) {
                            if ((856 < 18446744073709551615) || true) {
                                switch (fi8 / (((v1 ^ (fi6 & v5))) | 1)) {
                                case 815:
                                    fmt.Println(v3, fi6, v9)
                                case 1:
                                    v1 += v3
                                default:
                                    fmt.Println(v5, v5)
                                }
                                v22 := GenS4{F0: 1, F1: ("a" + "")}
                                _ = v22
                            } else {
                                v23 := GenS5{F0: 4294967295, F1: (v9 >> ((1) & 63)), F2: 2147483648}
                                if len(v4) > 0 {
                                    v4[0] = (-(v1))
                                }
                                _ = v23
                            }
                            v24 := GenS6{F0: "", F1: 817, F2: 4294967295}
                            var v25 int32 = 1
                            v25 = (^(v25))
                            if (255 <= 255) {
                                switch ("" + "hello") {
                                case "":
                                    v26 := GenS5{F0: 18446744073709551615, F1: (v9 - 32767), F2: 551}
                                    _ = v26
                                }
                                v3 += 127
                                v25 = (v25 ^ (v25 * (v25 / (((v25 - -770)) | 1))))
                            } else {
                                var v27 bool = (128 != 255)
                                switch "" {
                                case "":
                                    v9 -= v9
                                case "\n\t":
                                    v3 += (fi8 * (v5 | v5))
                                case "foo bar":
                                    v28 := GenS4{F0: 0, F1: ("\n\t" + "a")}
                                    _ = v28
                                default:
                                    if v27 {
                                        var v29 uint64 = 42
                                        switch 93 {
                                        case 2147483647:
                                            var v30 string = "a"
                                            fmt.Println(v3, v25, v29)
                                            _ = v30
                                        case 233:
                                            fmt.Println(v9)
                                            switch "a" {
                                            case "foo bar":
                                                if len(v4) > 0 {
                                                    v4[0] = (v1 >> ((1) & 63))
                                                }
                                            default:
                                                v25 -= v25
                                            }
                                        case -233:
                                            fmt.Println(v27, fi8, fi8)
                                        default:
                                            v27 = (2 >= 42)
                                            var v31 int64 = int64(v9)
                                            _ = v31
                                        }
                                        v32 := GenS4{F0: 42, F1: (("hello" + "hello") + "\n\t")}
                                        _ = v29
                                        _ = v32
                                    }
                                }
                                _ = v27
                            }
                            _ = v24
                            _ = v25
                        } else {
                            if true {
                                v1 += -784
                                if (!true) {
                                    if true {
                                        var v33 bool = true
                                        var v34 float32 = 3.14159
                                        var v35 int64 = -9
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    } else {
                                        if false {
                                            v1 = (-(fi8))
                                            fmt.Println(v2, v9, fi6)
                                            var v36 string = ("" + "foo bar")
                                            _ = v36
                                        } else {
                                            v2 -= v2
                                            v37 := GenS7{F0: ("foo bar" + "a"), F1: 0, F2: 4294967295}
                                            _ = v37
                                        }
                                    }
                                    switch (v1 + (^(v3))) {
                                    case 32768:
                                        var v38 uint16 = 42
                                        _ = v38
                                    case 42:
                                        v39 := GenS8{F0: 2147483647, F1: 42, F2: 501}
                                        _ = v39
                                    case 0:
                                        v40 := GenS9{F0: 32767}
                                        _ = v40
                                    default:
                                        var v41 int32 = 317
                                        _ = v41
                                    }
                                    var v42 uint32 = 4294967295
                                    _ = v42
                                }
                                v5 -= (v5 - 1)
                            } else {
                                var v43 uint8 = 1
                                _ = v43
                            }
                            v3 += (^(v3))
                            var v44 string = "a"
                            var v45 int = (fi8 + 27)
                            var v46 uint = 837
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                    } else {
                        if (42 == 294) {
                            if false {
                                var v47 uint64 = 1
                                switch v3 {
                                case -468:
                                    var v48 int32 = 42
                                    _ = v48
                                default:
                                    var v49 uint16 = 499
                                    _ = v49
                                }
                                if false {
                                    if true {
                                        v2 -= v2
                                    }
                                } else {
                                    if true {
                                        v4 = append(v4, int(v47))
                                        v50 := GenS5{F0: v47, F1: (v9 % ((1000) | 1)), F2: 42}
                                        switch "\n\t" {
                                        case "hello":
                                            if true {
                                                var v51 float32 = 3.14159
                                                _ = v51
                                            } else {
                                                var v52 uint8 = 127
                                                if true {
                                                    var v53 bool = false
                                                    _ = v53
                                                } else {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = v52
                                            }
                                        default:
                                            v54 := GenS10{F0: (("foo bar" + "\n\t") + ("hello" + "")), F1: v9}
                                            _ = v54
                                        }
                                        _ = v50
                                    }
                                    if (-329 >= -582) {
                                        if true {
                                            v55 := GenS4{F0: 49, F1: (("" + ("\n\t" + "\n\t")) + ("hello" + "a"))}
                                            var v56 uint16 = 42
                                            _ = v55
                                            _ = v56
                                        } else {
                                            if (!false) {
                                                _ = len(v4)
                                                _ = len(v4)
                                                var v57 uint = 42
                                                _ = v57
                                            }
                                            v9 += (v9 >> ((319) & 63))
                                            v3 = fi6
                                        }
                                        switch "hello" {
                                        case "a":
                                            if true {
                                                v9 = v9
                                            }
                                        }
                                        if false {
                                            var v58 int = v5
                                            v2 += 1e10
                                            _ = v58
                                        }
                                    } else {
                                        v2 = v2
                                        v5 += fi8
                                        if (894 <= 0) {
                                            v5 = 522
                                            if true {
                                                switch "foo bar" {
                                                case "":
                                                    fmt.Println(v2, fi8, v9)
                                                case "a":
                                                    v59 := GenS11{F0: 2}
                                                    _ = v59
                                                default:
                                                    if ((!(42 < -1)) || (2 <= 172)) {
                                                        v47 -= 1
                                                        _ = len(v4)
                                                        if len(v4) > 0 {
                                                            v4[0] = fi6
                                                        }
                                                    }
                                                }
                                                switch ("a" + "a") {
                                                case "a":
                                                    v9 -= (^(v9))
                                                case "foo bar":
                                                    switch (fi8 >> ((0) & 63)) {
                                                    case 4294967295:
                                                        v3 += (^(fi8))
                                                    }
                                                case "":
                                                    var v60 int = fi8
                                                    _ = v60
                                                }
                                                var v61 bool = true
                                                var v62 float32 = -2.25
                                                v63 := GenS12{F0: -71}
                                                _ = v61
                                                _ = v62
                                                _ = v63
                                            }
                                        }
                                    }
                                }
                                _ = len(v4)
                                _ = v47
                            }
                            v64 := GenS4{F0: 255, F1: ("hello" + "")}
                            if ("\n\t" >= "") {
                                if ((v9 * (v9 >> ((803) & 63))) <= (v9 >> ((65535) & 63))) {
                                    v65 := GenS3{F0: 1, F1: "hello"}
                                    _ = v65
                                }
                            } else {
                                v5 = (^(fi8))
                                switch ("" + ("hello" + "foo bar")) {
                                case "":
                                    if true {
                                        var v66 float32 = 0.1
                                        _ = v66
                                    }
                                default:
                                    var v67 uint64 = 18446744073709551615
                                    _ = v67
                                }
                                v2 += ((v2 + (-2.25 * (0.1 * ((3.14159 * v2) * float64(0.1))))) + float64((((3.14159 * 1e10) - (3.14159 / 1.0)) / 1.5)))
                            }
                            v5 += v1
                            var v68 string = "a"
                            v69 := GenS3{F0: 93, F1: v68}
                            _ = v64
                            _ = v68
                            _ = v69
                        } else {
                            if (v9 == (v9 | v9)) {
                                var v70 float32 = float32(v2)
                                if false {
                                    switch fi6 {
                                    case 42:
                                        v9 -= (v9 & -1)
                                    }
                                    v71 := GenS13{F0: 1, F1: 2}
                                    _ = v71
                                } else {
                                    fmt.Println(fi6)
                                    fmt.Println(v1)
                                }
                                _ = v70
                            }
                        }
                        if (127 <= 81) {
                            if true {
                                v5 -= fi6
                            }
                            v5 = (v3 << ((65535) & 63))
                        } else {
                            switch "hello" {
                            case "a":
                                v9 = 42
                            case "foo bar":
                                var v72 float64 = (v2 + (v2 * 0.1))
                                _ = v72
                            default:
                                switch "a" {
                                case "hello":
                                    if len(v4) > 0 {
                                        v4[0] = (v5 >> ((18446744073709551615) & 63))
                                    }
                                case "foo bar":
                                    if false {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v4 = append(v4, v5)
                                        _ = len(v4)
                                    } else {
                                        if (758 <= 2147483647) {
                                            switch "a" {
                                            case "hello":
                                                v2 -= (v2 * -2.25)
                                            }
                                            var v73 int8 = 126
                                            v74 := GenS1{F0: 255, F1: 65535}
                                            if false {
                                                v9 = v9
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                if (127 < 25) {
                                                    _ = len(v4)
                                                    _ = len(v4)
                                                } else {
                                                    if (-1 == v73) {
                                                        v73 -= v73
                                                        v4 = append(v4, (v1 | (-(fi8))))
                                                        v4 = append(v4, 0)
                                                    } else {
                                                        var v75 float32 = 1.5
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        var v76 int32 = 2147483647
                                                        _ = v75
                                                        _ = v76
                                                    }
                                                    if true {
                                                        if (471 >= 65535) {
                                                            var v77 uint16 = 255
                                                            v4 = append(v4, (v5 & (fi8 / ((-1) | 1))))
                                                            v2 = ((((v2 - v2) * v2) - v2) + (float64(3.14159) + 0.0))
                                                            _ = v77
                                                        }
                                                        var v78 uint8 = 2
                                                        if (0.0 >= 1.5) {
                                                            fmt.Println(v1)
                                                            if false {
                                                                fmt.Println(v73)
                                                                var v79 uint32 = 2147483647
                                                                _ = v79
                                                            }
                                                        }
                                                        v80 := GenS14{F0: 0, F1: 65535}
                                                        _ = v78
                                                        _ = v80
                                                    } else {
                                                        if true {
                                                            _ = len(v4)
                                                            var v81 uint32 = 2147483648
                                                            if len(v4) > 0 {
                                                                v4[0] = (fi6 + 4294967295)
                                                            }
                                                            _ = v81
                                                        } else {
                                                            if true {
                                                                v9 += 1000
                                                                v82 := GenS4{F0: 255, F1: "foo bar"}
                                                                v83 := GenS15{F0: v1}
                                                                _ = v82
                                                                _ = v83
                                                            }
                                                        }
                                                        v84 := GenS16{F0: 1}
                                                        _ = v84
                                                    }
                                                    v4 = append(v4, v5)
                                                    _ = len(v4)
                                                    _ = len(v4)
                                                }
                                            }
                                            v85 := GenS4{F0: 0, F1: "foo bar"}
                                            var v86 int32 = 1
                                            fmt.Println(fi6)
                                            v87 := GenS15{F0: (fi8 / ((fi6) | 1))}
                                            switch ("a" + "hello") {
                                            case "a":
                                                if false {
                                                    fmt.Println(v9)
                                                }
                                            case "hello":
                                                v73 = v73
                                            case "foo bar":
                                                switch (v1 - fi8) {
                                                case -895:
                                                    if true {
                                                        var v88 bool = false
                                                        if len(v4) > 0 {
                                                            v4[0] = 9223372036854775807
                                                        }
                                                        _ = v88
                                                    } else {
                                                        v89 := GenS17{F0: "a", F1: 42}
                                                        if (("\n\t" + "\n\t") != "hello") {
                                                            v1 = (fi6 / ((int(v9)) | 1))
                                                            fmt.Println(fi8)
                                                        } else {
                                                            switch (^(v5)) {
                                                            case 4294967295:
                                                                v90 := GenS18{F0: true, F1: 0}
                                                                _ = v90
                                                            default:
                                                                var v91 int32 = -2147483648
                                                                _ = v91
                                                            }
                                                            fmt.Println(v2)
                                                            var v92 uint16 = 255
                                                            _ = v92
                                                        }
                                                        var v93 uint8 = 128
                                                        _ = v89
                                                        _ = v93
                                                    }
                                                default:
                                                    switch fi8 {
                                                    case 9223372036854775807:
                                                        v86 = v86
                                                    }
                                                }
                                            default:
                                                if true {
                                                    v4 = append(v4, (v1 & -985))
                                                }
                                            }
                                            _ = v73
                                            _ = v74
                                            _ = v85
                                            _ = v86
                                            _ = v87
                                        } else {
                                            switch (^(v3)) {
                                            case -729:
                                                v5 = -1
                                            }
                                            fmt.Println(v1, v3)
                                        }
                                        var v94 uint8 = 1
                                        switch (("hello" + (("\n\t" + ("foo bar" + "a")) + "hello")) + ("\n\t" + "foo bar")) {
                                        case "\n\t":
                                            if (true && (v2 < (v2 - (0.1 + (-2.25 + 1e10))))) {
                                                var v95 int8 = 1
                                                switch v5 {
                                                case 0:
                                                    v2 = v2
                                                }
                                                var v96 uint32 = 703
                                                _ = v95
                                                _ = v96
                                            }
                                        }
                                        fmt.Println(fi8, v2)
                                        var v97 int16 = (v9 + v9)
                                        _ = v94
                                        _ = v97
                                    }
                                case "\n\t":
                                    fmt.Println(fi8)
                                }
                            }
                        }
                    }
                    _ = fi8
                    _ = v9
                }
            }
            _ = fi6
            _ = v7
        }
    }
    var v98 uint32 = 1
    switch v1 {
    case -950:
        if true {
            if (65535 < (v98 + (^(v98)))) {
                var v99 int8 = 0
                v4 = append(v4, (v1 - v1))
                _ = v99
            } else {
                v1 = (v1 % ((902) | 1))
                var v100 int16 = 0
                _ = v100
            }
            fmt.Println(v2, v5, v2)
            var v101 string = "foo bar"
            _ = v101
        }
    case 91:
        var v102 int = 32767
        _ = v102
    }
    switch 362 {
    case 591:
        switch v5 {
        case 2147483647:
            v103 := GenS7{F0: "\n\t", F1: 18446744073709551615, F2: v98}
            _ = v103
        case 247:
            for fi104 := 0; fi104 < 17; fi104++ {
                if false {
                    fmt.Println(v2, v5)
                    var v105 uint64 = 1
                    _ = v105
                } else {
                    v5 -= (v1 | v3)
                    for fi106 := 0; fi106 < 7; fi106++ {
                        var v107 bool = (18446744073709551615 < 18446744073709551615)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if v107 {
                            v4 = append(v4, 128)
                        } else {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            fmt.Println(v3, fi104, fi104)
                            v108 := GenS19{F0: false, F1: 1}
                            var v109 uint8 = 255
                            _ = v108
                            _ = v109
                        }
                        _ = fi106
                        _ = v107
                    }
                    if (!((!true) && (1.5 < (((1e10 - (0.1 + (1.5 + 3.14159))) / 3.14159) * 0.0)))) {
                        var v110 int64 = 1
                        v1 = 128
                        v111 := GenS11{F0: 0}
                        _ = v110
                        _ = v111
                    } else {
                        if ((v98 / (((v98 / ((2147483647) | 1))) | 1)) <= v98) {
                            for fi112 := 0; fi112 < 20; fi112++ {
                                switch "a" {
                                case "a":
                                    v5 = -2147483648
                                case "foo bar":
                                    v3 += fi112
                                case "\n\t":
                                    v5 = (^(v1))
                                }
                                v113 := GenS12{F0: -102}
                                v3 -= fi104
                                _ = fi112
                                _ = v113
                            }
                        } else {
                            switch (v3 - -911) {
                            case 32768:
                                var v114 int = v5
                                _ = v114
                            case 63:
                                v5 += v1
                            case -405:
                                var v115 int8 = 50
                                _ = v115
                            }
                            switch "\n\t" {
                            case "":
                                switch ("a" + "hello") {
                                case "\n\t":
                                    var v116 uint64 = 42
                                    _ = v116
                                case "a":
                                    if true {
                                        if true {
                                            for fi117 := 0; fi117 < 15; fi117++ {
                                                var v118 uint16 = 0
                                                _ = fi117
                                                _ = v118
                                            }
                                            if true {
                                                _ = len(v4)
                                            } else {
                                                fmt.Println(v3)
                                            }
                                            var v119 int16 = 23
                                            _ = v119
                                        } else {
                                            v98 = (^(v98))
                                            if (672 < 0) {
                                                for fi120 := 0; fi120 < 15; fi120++ {
                                                    var v121 float64 = (v2 / 1.5)
                                                    v2 = float64(float32((((-2.25 / 3.14159) / v2) * float64(0.1))))
                                                    _ = fi120
                                                    _ = v121
                                                }
                                            }
                                            var v122 uint64 = 556
                                            _ = v122
                                        }
                                        v98 = v98
                                        _ = len(v4)
                                    } else {
                                        switch fi104 {
                                        case 944:
                                            for fi123 := 0; fi123 < 3; fi123++ {
                                                var v124 uint16 = 92
                                                v1 += (fi123 | v1)
                                                _ = len(v4)
                                                if (v124 != (^(v124))) {
                                                    var v125 uint8 = 190
                                                    v4 = append(v4, v3)
                                                    _ = v125
                                                }
                                                _ = fi123
                                                _ = v124
                                            }
                                        case 584:
                                            var v126 string = ""
                                            _ = v126
                                        case 127:
                                            var v127 uint64 = 269
                                            _ = v127
                                        default:
                                            for fi128 := 0; fi128 < 16; fi128++ {
                                                if (!true) {
                                                    var v129 uint8 = 127
                                                    _ = v129
                                                } else {
                                                    var v130 int = v5
                                                    _ = v130
                                                }
                                                switch "foo bar" {
                                                case "foo bar":
                                                    fmt.Println(v3)
                                                default:
                                                    if len(v4) > 0 {
                                                        v4[0] = v3
                                                    }
                                                }
                                                _ = fi128
                                            }
                                        }
                                        var v131 uint64 = 18446744073709551615
                                        if (0 != 423) {
                                            var v132 float32 = -2.25
                                            var v133 int8 = 115
                                            _ = v132
                                            _ = v133
                                        } else {
                                            v3 = 0
                                        }
                                        _ = v131
                                    }
                                case "foo bar":
                                    if (((3.14159 - 1e10) - float32(((v2 + ((v2 * (3.14159 - 0.0)) + ((v2 + 3.14159) / v2))) - v2))) < 1.5) {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        switch (v5 + (v5 & v1)) {
                                        case -9223372036854775808:
                                            _ = len(v4)
                                        case 786:
                                            v5 = 32767
                                        case 42:
                                            v2 = float64(((((1e10 + (-2.25 / ((1.5 + float32(v2)) - ((3.14159 / 1.5) / -2.25)))) + 1e10) * (0.0 - float32(3.14159))) / 0.1))
                                        }
                                    } else {
                                        v2 -= v2
                                        if (42 >= 42) {
                                            if false {
                                                var v134 bool = true
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = v134
                                            } else {
                                                v1 -= (^(v5))
                                            }
                                        }
                                        v2 = ((1.5 + v2) / -2.25)
                                    }
                                default:
                                    var v135 float32 = (((-2.25 / 1e10) + 0.0) + (1e10 * 3.14159))
                                    _ = v135
                                }
                            case "foo bar":
                                var v136 int16 = 42
                                _ = v136
                            default:
                                v137 := GenS4{F0: 255, F1: "foo bar"}
                                _ = v137
                            }
                            fmt.Println(v98)
                        }
                        if len(v4) > 0 {
                            v4[0] = (fi104 + -1)
                        }
                        v98 += v98
                    }
                }
                _ = len(v4)
                _ = fi104
            }
        default:
            if false {
                switch "a" {
                case "hello":
                    if true {
                        var v138 string = ""
                        var v139 string = v138
                        _ = v138
                        _ = v139
                    } else {
                        v98 = 2147483647
                    }
                case "a":
                    if true {
                        switch 128 {
                        case 1:
                            var v140 uint = 0
                            _ = v140
                        case 671:
                            v2 += (float64((3.14159 * (float32(float64(((1.5 * -2.25) + 1.5))) * (-2.25 * float32((1e10 - (0.1 * ((v2 * (v2 + v2)) + v2)))))))) * 1e10)
                        default:
                            for fi141 := 0; fi141 < 20; fi141++ {
                                for fi142 := 0; fi142 < 10; fi142++ {
                                    var v143 float32 = (1.5 + -2.25)
                                    v3 += 0
                                    v144 := GenS13{F0: 1, F1: (v98 - v98)}
                                    _ = fi142
                                    _ = v143
                                    _ = v144
                                }
                                var v145 int64 = 0
                                v146 := GenS20{F0: "foo bar", F1: (^(v98))}
                                _ = fi141
                                _ = v145
                                _ = v146
                            }
                        }
                    } else {
                        if false {
                            if false {
                                fmt.Println(v1, v3)
                                if true {
                                    switch v1 {
                                    case 32768:
                                        v147 := GenS21{F0: 100}
                                        _ = v147
                                    case -2147483648:
                                        var v148 string = ""
                                        _ = v148
                                    default:
                                        var v149 uint8 = 167
                                        _ = v149
                                    }
                                    for fi150 := 0; fi150 < 8; fi150++ {
                                        if false {
                                            v2 = -2.25
                                            v5 = 32768
                                            v98 = 104
                                        } else {
                                            switch "hello" {
                                            case "":
                                                var v151 uint16 = 201
                                                _ = v151
                                            case "a":
                                                fmt.Println(v3, v5, v2)
                                            default:
                                                for fi152 := 0; fi152 < 8; fi152++ {
                                                    switch fi150 {
                                                    case -2147483648:
                                                        v3 = (v1 % (((v3 - v3)) | 1))
                                                        var v153 int = (-(v5))
                                                        _ = v153
                                                    case -766:
                                                        v3 = int(v98)
                                                    }
                                                    var v154 uint64 = uint64(v98)
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    _ = fi152
                                                    _ = v154
                                                }
                                            }
                                            var v155 uint32 = (v98 % ((0) | 1))
                                            _ = v155
                                        }
                                        _ = fi150
                                    }
                                    var v156 float32 = 0.0
                                    _ = v156
                                }
                            } else {
                                v2 -= 1.5
                            }
                            v1 = (v1 | 127)
                        } else {
                            var v157 bool = (true || true)
                            fmt.Println(v157)
                            var v158 int16 = 12
                            _ = v157
                            _ = v158
                        }
                        v1 -= 42
                    }
                default:
                    var v159 int16 = 739
                    _ = v159
                }
                var v160 uint64 = 0
                _ = len(v4)
                _ = v160
            } else {
                v1 = 1
                for fi161 := 0; fi161 < 3; fi161++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    switch (-(v1)) {
                    case 32767:
                        v2 += 3.14159
                    case -1:
                        v98 -= 0
                    default:
                        var v162 uint8 = 251
                        _ = v162
                    }
                    if (4294967295 != v98) {
                        v98 += v98
                    }
                    _ = fi161
                }
            }
        }
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v4 = append(v4, v1)
    if false {
        var v163 int32 = 438
        _ = len(v4)
        switch ("" + "a") {
        case "foo bar":
            v4 = append(v4, (v1 / ((v3) | 1)))
        case "":
            for fi164 := 0; fi164 < 19; fi164++ {
                v2 += 1.5
                _ = fi164
            }
        }
        _ = v163
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v3 = v5
    switch "hello" {
    case "":
        var v165 int64 = 0
        _ = v165
    case "foo bar":
        if false {
            v166 := GenS16{F0: 65535}
            _ = v166
        }
    case "\n\t":
        fmt.Println(v5, v2)
    default:
        for fi167 := 0; fi167 < 6; fi167++ {
            v168 := GenS5{F0: 214, F1: 364, F2: v98}
            v169 := GenS22{F0: v98, F1: 0, F2: 761}
            switch "a" {
            case "hello":
                if false {
                    for fi170 := 0; fi170 < 13; fi170++ {
                        var v171 int16 = 0
                        fmt.Println(v98)
                        _ = fi170
                        _ = v171
                    }
                    var v172 string = ("\n\t" + ("\n\t" + "\n\t"))
                    var v173 uint8 = 128
                    _ = v172
                    _ = v173
                }
            case "foo bar":
                var v174 float64 = float64(-2.25)
                _ = v174
            }
            _ = fi167
            _ = v168
            _ = v169
        }
    }
    fmt.Println(v2, v1)
    v175 := GenS12{F0: 9}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v175
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v98
}
