package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 int16; F1 bool }
type GenS9 struct { F0 bool; F1 uint }
type GenS10 struct { F0 uint16; F1 uint64; F2 uint16 }

func main() {
    var v1 uint8 = 255
    var v2 uint64 = 1
    var v3 bool = false
    switch 32767 {
    case 2147483647:
        var v4 int = 42
        _ = v4
    case -1:
        v5 := []int{2147483647, -1, -1, -667}
        _ = v5
    }
    var v6 bool = v3
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 bool = (2147483647 < -1)
        var v9 uint32 = 953
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v3, v6, v1)
    var v10 int64 = int64(v1)
    if (v3 && v3) {
        if ((32767 <= 1) || v6) {
            fmt.Println(v10, v2)
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 int16 = -32768
                _ = fi11
                _ = v12
            }
        }
        v1 = v1
        for fi13 := 0; fi13 < 15; fi13++ {
            v14 := GenS1{F0: (^(v1)), F1: 805}
            var v15 uint8 = 255
            if v6 {
                switch "" {
                case "":
                    switch -987 {
                    case 4294967295:
                        var v16 int32 = -2147483648
                        _ = v16
                    case 128:
                        fmt.Println(v3, v3)
                    case 127:
                        v2 = uint64(v15)
                    default:
                        v3 = ((!false) && (!("a" != (("" + "") + "hello"))))
                    }
                case "\n\t":
                    for fi17 := 0; fi17 < 11; fi17++ {
                        if false {
                            var v18 uint8 = v15
                            _ = v18
                        }
                        var v19 int32 = -412
                        _ = fi17
                        _ = v19
                    }
                case "hello":
                    v6 = v6
                default:
                    switch fi13 {
                    case 0:
                        var v20 int64 = v10
                        _ = v20
                    case -569:
                        if (951 < 2147483647) {
                            v3 = (((-2.25 / 0.1) / (-2.25 * 3.14159)) > (-2.25 - (float64((1e10 - 0.1)) * 1e10)))
                            v3 = (v6 && (1e10 == 1e10))
                        }
                    case 32767:
                        if v3 {
                            for fi21 := 0; fi21 < 9; fi21++ {
                                fmt.Println(v3)
                                switch 4294967295 {
                                case 42:
                                    var v22 int16 = 1000
                                    _ = v22
                                case 127:
                                    switch 9223372036854775807 {
                                    case 98:
                                        switch "foo bar" {
                                        case "hello":
                                            fmt.Println(v10)
                                        default:
                                            if v6 {
                                                v3 = false
                                                if ((1 != 783) || v3) {
                                                    var v23 float64 = 3.14159
                                                    v23 = float64(-2.25)
                                                    _ = v23
                                                } else {
                                                    if (1 < 579) {
                                                        switch fi13 {
                                                        case -2147483648:
                                                            v6 = v3
                                                        case 32767:
                                                            v3 = v3
                                                        case 9223372036854775807:
                                                            v6 = v3
                                                        }
                                                    }
                                                    if v3 {
                                                        var v24 uint8 = (v1 % ((42) | 1))
                                                        fmt.Println(v15)
                                                        _ = v24
                                                    } else {
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v1 -= v1
                                                        case "hello":
                                                            switch "a" {
                                                            case "foo bar":
                                                                if v3 {
                                                                    v25 := []int{int(v2), 959}
                                                                    switch 127 {
                                                                    case 4294967295:
                                                                        if v6 {
                                                                            _ = len(v25)
                                                                            if len(v25) > 0 {
                                                                                v25[0] = 127
                                                                            }
                                                                        }
                                                                    case 0:
                                                                        if v3 {
                                                                            if len(v25) > 0 {
                                                                                v25[0] = 635
                                                                            }
                                                                        } else {
                                                                            if len(v25) >= 2 {
                                                                                v25 = v25[1:len(v25)]
                                                                            }
                                                                        }
                                                                    case 149:
                                                                        v3 = (-611 == 0)
                                                                        var v26 uint32 = 2147483647
                                                                        _ = v26
                                                                    default:
                                                                        v6 = (v6 && v3)
                                                                    }
                                                                    var v27 uint64 = 42
                                                                    _ = v25
                                                                    _ = v27
                                                                } else {
                                                                    fmt.Println(v15)
                                                                    if v3 {
                                                                        var v28 uint16 = 32767
                                                                        _ = v28
                                                                    }
                                                                    v2 -= uint64(v15)
                                                                    v29 := []int{fi21}
                                                                    switch (fi13 / ((fi13) | 1)) {
                                                                    case 92:
                                                                        v1 += v15
                                                                    case 42:
                                                                        if (!(v3 && (-91 > 748))) {
                                                                            var v30 float32 = 0.0
                                                                            var v31 int32 = 2147483647
                                                                            _ = v30
                                                                            _ = v31
                                                                        } else {
                                                                            if v3 {
                                                                                var v32 uint32 = 2147483647
                                                                                fmt.Println(v32, fi21)
                                                                                var v33 int = fi21
                                                                                _ = v32
                                                                                _ = v33
                                                                            } else {
                                                                                if true {
                                                                                    switch (("hello" + "a") + "hello") {
                                                                                    case "foo bar":
                                                                                        switch (fi13 / ((32768) | 1)) {
                                                                                        case 4294967295:
                                                                                            fmt.Println(fi13, v10, v10)
                                                                                        case 174:
                                                                                            var v34 int32 = int32(v10)
                                                                                            _ = v34
                                                                                        case 32767:
                                                                                            var v35 int16 = -200
                                                                                            _ = v35
                                                                                        }
                                                                                    case "\n\t":
                                                                                        switch "foo bar" {
                                                                                        case "hello":
                                                                                            fmt.Println(v3, v1)
                                                                                        case "":
                                                                                            switch 4294967295 {
                                                                                            case -187:
                                                                                                v3 = (v6 || v6)
                                                                                            case 2147483647:
                                                                                                fmt.Println(v1, v10)
                                                                                                var v36 uint32 = 2
                                                                                                v1 += v15
                                                                                                if ((float32(-2.25) / 0.1) == (float32((0.0 / -2.25)) * 3.14159)) {
                                                                                                    switch 474 {
                                                                                                    case 42:
                                                                                                        v6 = ((0.0 + (-2.25 / 1.5)) > 1.5)
                                                                                                    default:
                                                                                                        var v37 int16 = -1
                                                                                                        _ = v37
                                                                                                    }
                                                                                                    if (!v6) {
                                                                                                        if v6 {
                                                                                                            v38 := GenS1{F0: 128, F1: 2}
                                                                                                            _ = v38
                                                                                                        } else {
                                                                                                            var v39 string = "a"
                                                                                                            if (0.1 == 0.1) {
                                                                                                                v36 = (v36 & v36)
                                                                                                            } else {
                                                                                                                var v40 uint64 = 1
                                                                                                                v1 = 128
                                                                                                                _ = v40
                                                                                                            }
                                                                                                            _ = v39
                                                                                                        }
                                                                                                        _ = len(v29)
                                                                                                    } else {
                                                                                                        var v41 uint8 = 1
                                                                                                        v29 = append(v29, fi21)
                                                                                                        v42 := GenS1{F0: v41, F1: 729}
                                                                                                        _ = v41
                                                                                                        _ = v42
                                                                                                    }
                                                                                                    if ((fi13 & (fi13 / ((fi13) | 1))) > -1) {
                                                                                                        var v43 float64 = -2.25
                                                                                                        v6 = ((3.14159 * (v43 - v43)) != v43)
                                                                                                        v44 := GenS1{F0: v1, F1: 42}
                                                                                                        _ = v43
                                                                                                        _ = v44
                                                                                                    } else {
                                                                                                        var v45 uint32 = v36
                                                                                                        v46 := GenS2{F0: 255, F1: 1000}
                                                                                                        _ = v45
                                                                                                        _ = v46
                                                                                                    }
                                                                                                } else {
                                                                                                    switch "hello" {
                                                                                                    case "foo bar":
                                                                                                        _ = len(v29)
                                                                                                    case "":
                                                                                                        var v47 bool = v6
                                                                                                        switch 2147483647 {
                                                                                                        case 127:
                                                                                                            v3 = (!(!false))
                                                                                                        case 32768:
                                                                                                            v48 := GenS1{F0: (^(v1)), F1: 1}
                                                                                                            _ = v48
                                                                                                        case 9223372036854775807:
                                                                                                            switch fi13 {
                                                                                                            case 726:
                                                                                                                var v49 string = ("" + "a")
                                                                                                                _ = v49
                                                                                                            case 4294967295:
                                                                                                                var v50 int16 = 1
                                                                                                                _ = v50
                                                                                                            }
                                                                                                            var v51 bool = ((356 < (v36 * (v36 | (v36 + v36)))) && v47)
                                                                                                            var v52 uint64 = (v2 * 18446744073709551615)
                                                                                                            if len(v29) > 0 {
                                                                                                                v29[0] = 0
                                                                                                            }
                                                                                                            v52 += (v52 - (v52 >> ((1) & 63)))
                                                                                                            v53 := GenS2{F0: 4294967295, F1: 1}
                                                                                                            _ = v51
                                                                                                            _ = v52
                                                                                                            _ = v53
                                                                                                        default:
                                                                                                            if (v3 && (!(42 >= 42))) {
                                                                                                                v1 = (v1 | v15)
                                                                                                            }
                                                                                                            v6 = v3
                                                                                                        }
                                                                                                        _ = v47
                                                                                                    default:
                                                                                                        v54 := GenS3{F0: 1, F1: "a"}
                                                                                                        _ = v54
                                                                                                    }
                                                                                                    v55 := GenS2{F0: 42, F1: 1000}
                                                                                                    if (!false) {
                                                                                                        var v56 uint = 30
                                                                                                        if v6 {
                                                                                                            v36 = uint32(v2)
                                                                                                            v36 = v36
                                                                                                        }
                                                                                                        if (1.5 == 3.14159) {
                                                                                                            var v57 bool = v3
                                                                                                            v29 = append(v29, 2147483647)
                                                                                                            fmt.Println(fi21, v15)
                                                                                                            _ = v57
                                                                                                        }
                                                                                                        _ = v56
                                                                                                    } else {
                                                                                                        v6 = (0.0 >= (1.5 * 0.0))
                                                                                                        fmt.Println(fi21)
                                                                                                    }
                                                                                                    _ = v55
                                                                                                }
                                                                                                v3 = false
                                                                                                if (((1000 != 45) || (42 >= 42)) || (-1 == -32768)) {
                                                                                                    v6 = ((fi13 * fi21) <= fi21)
                                                                                                    v1 += (v1 + 0)
                                                                                                }
                                                                                                _ = v36
                                                                                            case 389:
                                                                                                fmt.Println(v2, fi21, v15)
                                                                                            default:
                                                                                                if len(v29) >= 2 {
                                                                                                    v29 = v29[1:len(v29)]
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    if (((^(v10)) >= v10) || ((fi21 * fi13) >= fi13)) {
                                                                                        var v58 int32 = 1
                                                                                        switch "" {
                                                                                        case "hello":
                                                                                            v1 = v15
                                                                                        case "foo bar":
                                                                                            v10 = (v10 & 1)
                                                                                        case "\n\t":
                                                                                            fmt.Println(v6, v2)
                                                                                        }
                                                                                        v3 = false
                                                                                        v10 -= v10
                                                                                        if len(v29) > 0 {
                                                                                            v29[0] = 405
                                                                                        }
                                                                                        _ = v58
                                                                                    } else {
                                                                                        v6 = (664 == 1)
                                                                                        var v59 int16 = -1
                                                                                        v60 := GenS1{F0: (v1 % ((uint8(v2)) | 1)), F1: 926}
                                                                                        v61 := GenS2{F0: 4294967295, F1: v59}
                                                                                        _ = v59
                                                                                        _ = v60
                                                                                        _ = v61
                                                                                    }
                                                                                } else {
                                                                                    v6 = (32767 != 120)
                                                                                }
                                                                                var v62 uint16 = 1
                                                                                var v63 string = ("hello" + "")
                                                                                _ = v62
                                                                                _ = v63
                                                                            }
                                                                            v1 = v1
                                                                            var v64 int8 = 1
                                                                            var v65 int16 = 377
                                                                            v15 -= (v15 & v15)
                                                                            v1 = uint8(v2)
                                                                            _ = v64
                                                                            _ = v65
                                                                        }
                                                                    case 4294967295:
                                                                        v29 = append(v29, fi21)
                                                                    }
                                                                    _ = v29
                                                                }
                                                            case "a":
                                                                if true {
                                                                    v10 = v10
                                                                }
                                                            case "hello":
                                                                switch ((("foo bar" + "\n\t") + ("\n\t" + "")) + "foo bar") {
                                                                case "\n\t":
                                                                    if v3 {
                                                                        var v66 uint8 = 1
                                                                        if v6 {
                                                                            v67 := GenS1{F0: (v1 ^ (v15 | uint8(v2))), F1: 1}
                                                                            var v68 uint32 = 1
                                                                            var v69 uint32 = (v68 * v68)
                                                                            var v70 string = ("" + "")
                                                                            _ = v67
                                                                            _ = v68
                                                                            _ = v69
                                                                            _ = v70
                                                                        } else {
                                                                            var v71 uint64 = (^(v2))
                                                                            v72 := GenS4{F0: (v1 / ((v1) | 1))}
                                                                            if (334 >= (^(v2))) {
                                                                                v73 := []int{(fi13 | fi13), (fi21 + 2147483647), 0, (fi13 >> ((255) & 63))}
                                                                                switch -9223372036854775808 {
                                                                                case 1:
                                                                                    switch "hello" {
                                                                                    case "a":
                                                                                        switch "hello" {
                                                                                        case "foo bar":
                                                                                            switch (^(fi13)) {
                                                                                            case -2147483648:
                                                                                                v71 = 4294967295
                                                                                            }
                                                                                        case "hello":
                                                                                            var v74 bool = v3
                                                                                            _ = v74
                                                                                        }
                                                                                    case "hello":
                                                                                        v15 -= (^(v1))
                                                                                    default:
                                                                                        switch -876 {
                                                                                        case -837:
                                                                                            switch 127 {
                                                                                            case 42:
                                                                                                fmt.Println(v6, v66)
                                                                                            case 0:
                                                                                                var v75 string = "foo bar"
                                                                                                _ = v75
                                                                                            case 4294967295:
                                                                                                var v76 int = (fi13 | fi13)
                                                                                                _ = v76
                                                                                            default:
                                                                                                v73 = append(v73, -1)
                                                                                            }
                                                                                        default:
                                                                                            if true {
                                                                                                switch "foo bar" {
                                                                                                case "":
                                                                                                    switch "foo bar" {
                                                                                                    case "hello":
                                                                                                        v77 := GenS5{F0: "a", F1: 0}
                                                                                                        _ = v77
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    var v78 uint16 = 32767
                                                                                                    _ = v78
                                                                                                case "hello":
                                                                                                    if v6 {
                                                                                                        if len(v73) >= 2 {
                                                                                                            v73 = v73[1:len(v73)]
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                var v79 float64 = 1.5
                                                                                                _ = v79
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v73
                                                                            }
                                                                            _ = v71
                                                                            _ = v72
                                                                        }
                                                                        _ = v66
                                                                    } else {
                                                                        v80 := []int{-9223372036854775808, (fi13 / ((0) | 1))}
                                                                        if v3 {
                                                                            var v81 uint16 = 294
                                                                            _ = v81
                                                                        }
                                                                        _ = v80
                                                                    }
                                                                }
                                                            default:
                                                                var v82 float64 = (0.0 - ((1e10 - 0.0) / (1.5 - (-2.25 / -2.25))))
                                                                _ = v82
                                                            }
                                                        }
                                                        fmt.Println(v3, fi21)
                                                    }
                                                    var v83 float32 = 0.0
                                                    _ = v83
                                                }
                                                var v84 float64 = 3.14159
                                                if v6 {
                                                    v85 := []int{fi21, -1, fi13}
                                                    switch fi21 {
                                                    case 203:
                                                        v2 = 587
                                                    case 1:
                                                        v86 := GenS6{F0: 48}
                                                        _ = v86
                                                    case 127:
                                                        v6 = false
                                                    }
                                                    _ = v85
                                                } else {
                                                    var v87 int16 = 0
                                                    switch 2147483647 {
                                                    case 128:
                                                        if v3 {
                                                            v88 := []int{fi13}
                                                            _ = v88
                                                        }
                                                    case -2147483648:
                                                        if v3 {
                                                            v10 -= v10
                                                            v89 := GenS5{F0: ("foo bar" + "\n\t"), F1: 1000}
                                                            _ = v89
                                                        }
                                                    case 963:
                                                        var v90 int = (fi13 >> ((2) & 63))
                                                        _ = v90
                                                    default:
                                                        if (!v6) {
                                                            fmt.Println(v6, fi21, v1)
                                                            var v91 int32 = -1
                                                            v92 := []int{(fi13 | (^(fi13))), int(v10), -2147483648}
                                                            _ = v91
                                                            _ = v92
                                                        } else {
                                                            var v93 uint8 = 82
                                                            _ = v93
                                                        }
                                                    }
                                                    _ = v87
                                                }
                                                var v94 int64 = v10
                                                v95 := []int{}
                                                _ = v84
                                                _ = v94
                                                _ = v95
                                            } else {
                                                v15 = v1
                                            }
                                        }
                                    case -1:
                                        v96 := GenS1{F0: v15, F1: 42}
                                        _ = v96
                                    case 952:
                                        var v97 int16 = 42
                                        _ = v97
                                    default:
                                        var v98 int = fi13
                                        _ = v98
                                    }
                                default:
                                    if (164 == 42) {
                                        v10 -= (v10 * v10)
                                        if v3 {
                                            switch (("\n\t" + "") + "") {
                                            case "\n\t":
                                                v2 = v2
                                            case "a":
                                                var v99 uint = 4294967295
                                                _ = v99
                                            case "":
                                                if v3 {
                                                    fmt.Println(v10)
                                                } else {
                                                    v100 := GenS3{F0: (v10 | v10), F1: "hello"}
                                                    v101 := GenS2{F0: 941, F1: 0}
                                                    v102 := GenS6{F0: 65535}
                                                    _ = v100
                                                    _ = v101
                                                    _ = v102
                                                }
                                            }
                                            var v103 int16 = 751
                                            if ((1e10 / 1.5) != 0.0) {
                                                var v104 string = "hello"
                                                v2 = (v2 * v2)
                                                _ = v104
                                            } else {
                                                v105 := GenS7{F0: (!v6), F1: 2147483647}
                                                var v106 float32 = ((3.14159 * 0.1) / 1.5)
                                                fmt.Println(v106)
                                                _ = v105
                                                _ = v106
                                            }
                                            _ = v103
                                        }
                                    }
                                }
                                _ = fi21
                            }
                        }
                    default:
                        v107 := GenS2{F0: 4294967295, F1: 0}
                        _ = v107
                    }
                }
                if (!false) {
                    v108 := []int{}
                    _ = v108
                }
            }
            _ = fi13
            _ = v14
            _ = v15
        }
    } else {
        if v3 {
            var v109 int64 = (v10 << ((892) & 63))
            var v110 float32 = 1.5
            switch 32767 {
            case 1:
                for fi111 := 0; fi111 < 2; fi111++ {
                    v112 := GenS8{F0: 712, F1: v3}
                    if (!(v3 && (!(v3 && ((-9223372036854775808 < v109) && (255 <= 901)))))) {
                        switch fi111 {
                        case 9223372036854775807:
                            switch ("a" + ("\n\t" + "\n\t")) {
                            case "hello":
                                fmt.Println(fi111, v10)
                            case "foo bar":
                                for fi113 := 0; fi113 < 4; fi113++ {
                                    var v114 int8 = 0
                                    switch 811 {
                                    case -954:
                                        var v115 uint = 644
                                        _ = v115
                                    case -2147483648:
                                        var v116 int = (fi111 + -879)
                                        _ = v116
                                    }
                                    _ = fi113
                                    _ = v114
                                }
                            default:
                                fmt.Println(v1, v6, v3)
                            }
                        case -1:
                            if v3 {
                                switch "\n\t" {
                                case "hello":
                                    switch "\n\t" {
                                    case "":
                                        v109 = v10
                                    case "hello":
                                        fmt.Println(v109, v10, v10)
                                    }
                                case "foo bar":
                                    var v117 int8 = 42
                                    _ = v117
                                case "":
                                    var v118 uint32 = 906
                                    _ = v118
                                }
                            }
                        case -14:
                            var v119 uint8 = 154
                            _ = v119
                        }
                    } else {
                        for fi120 := 0; fi120 < 9; fi120++ {
                            switch "" {
                            case "\n\t":
                                v121 := []int{(fi120 | (fi111 & (^(fi120)))), fi111}
                                _ = v121
                            }
                            _ = fi120
                        }
                        v10 -= (v109 | -1)
                        v122 := []int{}
                        _ = v122
                    }
                    fmt.Println(v10)
                    _ = fi111
                    _ = v112
                }
            }
            _ = v109
            _ = v110
        } else {
            if true {
                var v123 float64 = (3.14159 * (0.1 - 1e10))
                v1 = (v1 << ((18446744073709551615) & 63))
                _ = v123
            }
            v124 := GenS6{F0: 610}
            _ = v124
        }
    }
    var v125 float64 = -2.25
    v2 = (v2 & (^(v2)))
    v6 = v3
    switch 127 {
    case 4294967295:
        v126 := GenS9{F0: false, F1: 1}
        _ = v126
    case 32767:
        v127 := GenS3{F0: v10, F1: "a"}
        _ = v127
    }
    for fi128 := 0; fi128 < 2; fi128++ {
        v2 += 18446744073709551615
        _ = fi128
    }
    switch "foo bar" {
    case "foo bar":
        switch ("" + "hello") {
        case "\n\t":
            for fi129 := 0; fi129 < 7; fi129++ {
                for fi130 := 0; fi130 < 11; fi130++ {
                    fmt.Println(v10)
                    _ = fi130
                }
                v125 -= v125
                fmt.Println(v6)
                _ = fi129
            }
        case "a":
            v2 = (v2 * 465)
        case "foo bar":
            for fi131 := 0; fi131 < 2; fi131++ {
                v6 = (0 < -1)
                v132 := GenS8{F0: 0, F1: v3}
                _ = fi131
                _ = v132
            }
        default:
            var v133 int64 = v10
            _ = v133
        }
    case "hello":
        v134 := GenS7{F0: v6, F1: 1}
        _ = v134
    case "a":
        if true {
            var v135 uint8 = (^(v1))
            v135 = (v135 | (v1 - (^(v135))))
            _ = v135
        } else {
            var v136 uint64 = v2
            v137 := []int{9223372036854775807, 2147483647, 0}
            _ = v136
            _ = v137
        }
    default:
        fmt.Println(v10)
    }
    for fi138 := 0; fi138 < 4; fi138++ {
        if true {
            if (false && true) {
                v6 = ((-1 > 1000) || v3)
            }
            fmt.Println(fi138)
            switch (("hello" + ("hello" + "hello")) + "\n\t") {
            case "\n\t":
                var v139 int = 4294967295
                _ = v139
            case "a":
                v140 := GenS10{F0: 345, F1: v2, F2: 973}
                _ = v140
            case "hello":
                v3 = (4294967295 >= 2147483647)
            }
        }
        _ = fi138
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v125:", v125)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v125
    _ = v2
    _ = v3
    _ = v6
}
