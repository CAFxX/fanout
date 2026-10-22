package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int16 }
type GenS2 struct { F0 bool; F1 int; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 uint16; F1 uint; F2 uint }
type GenS7 struct { F0 uint; F1 int32; F2 uint32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 string; F1 uint64 }
type GenS12 struct { F0 int; F1 uint64; F2 uint16 }
type GenS13 struct { F0 uint32; F1 int16; F2 uint }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 int }

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"hello", "foo bar", ((("a" + "hello") + ("\n\t" + "a")) + ("" + ("" + (("a" + "hello") + "a"))))}
    v3 := make(map[string]uint32)
    var v4 int32 = -1
    for fi5 := 0; fi5 < 17; fi5++ {
        v4 -= -731
        var v6 int8 = 42
        for fi7 := 0; fi7 < 9; fi7++ {
            v4 = v4
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    v8 := v1[65535]
    v4 += -2147483648
    v4 -= (v4 - 1)
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 uint64 = 4294967295
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 int8 = 16
        var v13 int = 127
        switch v13 {
        case -2147483648:
            if true {
                switch v13 {
                case 2147483647:
                    fmt.Println(v8)
                case -181:
                    if (255 > 0) {
                        switch "\n\t" {
                        case "foo bar":
                            for fi14 := 0; fi14 < 3; fi14++ {
                                v13 = (-(fi14))
                                fmt.Println(v4, v4)
                                if true {
                                    v15 := GenS1{F0: 908, F1: false, F2: -289}
                                    v12 += v12
                                    var v16 uint8 = 1
                                    _ = v15
                                    _ = v16
                                }
                                _ = fi14
                            }
                        }
                    } else {
                        for fi17 := 0; fi17 < 7; fi17++ {
                            if false {
                                var v18 int32 = v4
                                var v19 uint32 = 668
                                v18 = -2147483648
                                _ = v18
                                _ = v19
                            } else {
                                fmt.Println(v8, fi11)
                            }
                            var v20 string = ((v8 + v8) + (v8 + v8))
                            switch ("" + v20) {
                            case "\n\t":
                                v8 = ("" + (v20 + v8))
                            case "a":
                                v13 += (v13 & 9223372036854775807)
                            case "":
                                v3[v20] = 65535
                                var v21 bool = (!true)
                                v13 -= (^(v13))
                                _ = v21
                            default:
                                if ((0 != 802) && (("foo bar" < v8) && false)) {
                                    if (-32768 != -32768) {
                                        v22 := GenS1{F0: 700, F1: false, F2: 0}
                                        _ = v22
                                    }
                                    var v23 int = fi17
                                    var v24 float64 = (0.1 + -2.25)
                                    _ = v23
                                    _ = v24
                                }
                            }
                            _ = fi17
                            _ = v20
                        }
                        var v25 bool = false
                        _ = v25
                    }
                }
                if (219 == 0) {
                    var v26 string = (v8 + (("\n\t" + v8) + v8))
                    switch v26 {
                    case "foo bar":
                        var v27 string = (v8 + (v26 + ""))
                        _ = v27
                    case "hello":
                        v8 = "hello"
                    case "":
                        v28 := GenS1{F0: 752, F1: ((2 != 0) && ((!false) || false)), F2: 324}
                        _ = v28
                    }
                    switch (fi11 + fi11) {
                    case 32768:
                        if len(v2) > 0 {
                            v2[0] = (v26 + "a")
                        }
                    }
                    _ = v26
                }
                v29 := GenS2{F0: false, F1: v13, F2: -114}
                _ = v29
            } else {
                v4 += 42
                v30 := GenS2{F0: false, F1: fi11, F2: -32768}
                switch v13 {
                case 1:
                    if true {
                        switch "a" {
                        case "":
                            var v31 int8 = 82
                            _ = v31
                        }
                        var v32 float64 = 0.0
                        _ = v32
                    } else {
                        if false {
                            switch v8 {
                            case "hello":
                                v13 += (v13 % ((fi11) | 1))
                            case "a":
                                for fi33 := 0; fi33 < 3; fi33++ {
                                    v12 -= v12
                                    _ = fi33
                                }
                            case "\n\t":
                                for fi34 := 0; fi34 < 20; fi34++ {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(v13)
                                    _ = fi34
                                }
                            default:
                                for fi35 := 0; fi35 < 1; fi35++ {
                                    var v36 float32 = 0.0
                                    _ = fi35
                                    _ = v36
                                }
                            }
                            for fi37 := 0; fi37 < 1; fi37++ {
                                v38 := GenS1{F0: 13, F1: true, F2: -32768}
                                v13 += fi11
                                v4 += (v4 & v4)
                                _ = fi37
                                _ = v38
                            }
                            v3[v8] = 323
                        }
                        if (-32768 != 0) {
                            fmt.Println(v8, v8, v4)
                            if false {
                                if len(v2) > 0 {
                                    v2[0] = "hello"
                                }
                                for fi39 := 0; fi39 < 5; fi39++ {
                                    v8 = ("" + "\n\t")
                                    v40 := GenS3{F0: (-(v4))}
                                    _ = fi39
                                    _ = v40
                                }
                            } else {
                                for fi41 := 0; fi41 < 3; fi41++ {
                                    fmt.Println(fi11, v4, v4)
                                    v42 := GenS4{F0: 65535, F1: (v12 & v12)}
                                    v2 = append(v2, v8)
                                    _ = fi41
                                    _ = v42
                                }
                            }
                        }
                    }
                case -2147483648:
                    v12 = (v12 / (((v12 - -86)) | 1))
                default:
                    if (((("a" + "hello") + (("a" + v8) + ((v8 + ("" + v8)) + v8))) + v8) < v8) {
                        if false {
                            v43 := v3[((v8 + ((v8 + "a") + v8)) + "")]
                            v13 += (fi11 | (fi11 + (fi11 / (((v13 + (-(v13)))) | 1))))
                            _ = v43
                        }
                        for fi44 := 0; fi44 < 17; fi44++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v45 := GenS5{F0: 2, F1: 223}
                            var v46 int32 = (v4 | -2147483648)
                            if (42 > 496) {
                                v47 := GenS5{F0: 1, F1: 4294967295}
                                _ = v47
                            } else {
                                switch "\n\t" {
                                case "hello":
                                    if false {
                                        _ = len(v1)
                                    }
                                case "foo bar":
                                    var v48 bool = (4294967295 <= 65535)
                                    _ = v48
                                case "":
                                    v12 = (-(v12))
                                }
                                v49 := v3["foo bar"]
                                v49 = v49
                                if false {
                                    if false {
                                        var v50 uint32 = 4294967295
                                        _ = v50
                                    } else {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = len(v1)
                                        fmt.Println(v49)
                                    }
                                }
                                _ = v49
                            }
                            fmt.Println(v46, v4, v12)
                            _ = fi44
                            _ = v45
                            _ = v46
                        }
                    } else {
                        v51 := GenS1{F0: 2, F1: false, F2: 558}
                        switch v8 {
                        case "a":
                            v52 := GenS2{F0: true, F1: v13, F2: 1}
                            _ = v52
                        case "foo bar":
                            v8 = v8
                        case "hello":
                            v53 := v3[("a" + v8)]
                            _ = v53
                        }
                        _ = v51
                    }
                }
                _ = v30
            }
        default:
            if false {
                v4 = v4
                if true {
                    switch "" {
                    case "foo bar":
                        var v54 uint32 = 42
                        _ = v54
                    case "\n\t":
                        delete(v3, "a")
                    }
                    switch (v8 + (v8 + "\n\t")) {
                    case "hello":
                        var v55 string = ("foo bar" + ((v8 + ("\n\t" + v8)) + v8))
                        _ = v55
                    case "\n\t":
                        for fi56 := 0; fi56 < 20; fi56++ {
                            switch v8 {
                            case "foo bar":
                                var v57 uint8 = 128
                                _ = v57
                            case "a":
                                var v58 uint16 = 59
                                _ = v58
                            case "hello":
                                var v59 int = (v13 ^ fi56)
                                _ = v59
                            default:
                                var v60 string = v8
                                _ = v60
                            }
                            v8 = v8
                            _ = fi56
                        }
                    default:
                        v61 := GenS6{F0: 32767, F1: 190, F2: 42}
                        _ = v61
                    }
                    v62 := GenS4{F0: 13, F1: (v12 | (v12 ^ v12))}
                    _ = v62
                } else {
                    if true {
                        if false {
                            var v63 bool = false
                            if (0.1 == ((3.14159 / 0.1) - (1.5 + 1.5))) {
                                v63 = (-9223372036854775808 == v13)
                                switch (("hello" + "\n\t") + "foo bar") {
                                case "hello":
                                    v12 = 100
                                }
                                switch "\n\t" {
                                case "hello":
                                    var v64 uint = 0
                                    _ = v64
                                default:
                                    for fi65 := 0; fi65 < 17; fi65++ {
                                        v4 -= v4
                                        if v63 {
                                            v66 := GenS6{F0: 42, F1: 871, F2: 674}
                                            var v67 int8 = v12
                                            v67 = v67
                                            v12 = v67
                                            if false {
                                                if len(v2) > 0 {
                                                    v2[0] = ""
                                                }
                                            }
                                            _ = v66
                                            _ = v67
                                        } else {
                                            var v68 int8 = v12
                                            var v69 string = "\n\t"
                                            _ = v68
                                            _ = v69
                                        }
                                        v70 := GenS4{F0: 0, F1: v12}
                                        if ((100 < v12) && ((("" + v8) + v8) <= (((v8 + "a") + "a") + ((v8 + v8) + v8)))) {
                                            if (32767 >= 2) {
                                                if (2 < 255) {
                                                    v71 := GenS3{F0: (v4 & v4)}
                                                    v8 = v8
                                                    v8 = v8
                                                    _ = v71
                                                }
                                            }
                                            switch (v13 % ((1) | 1)) {
                                            case -443:
                                                if ((-489 > 0) || (18446744073709551615 <= 2)) {
                                                    fmt.Println(fi65, fi65)
                                                    v72 := GenS6{F0: 2, F1: 18446744073709551615, F2: 0}
                                                    if (484 != 2147483648) {
                                                        var v73 uint64 = 4294967295
                                                        var v74 float32 = (0.1 / (3.14159 + -2.25))
                                                        v75 := GenS2{F0: ((v73 << ((18446744073709551615) & 63)) < (^(v73))), F1: 127, F2: 42}
                                                        _ = v73
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                    var v76 int = 732
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    _ = v72
                                                    _ = v76
                                                }
                                            default:
                                                var v77 uint32 = 71
                                                _ = v77
                                            }
                                            if false {
                                                v3[v8] = 42
                                                if v63 {
                                                    if false {
                                                        switch ((v8 + v8) + v8) {
                                                        case "foo bar":
                                                            var v78 int = -210
                                                            switch (fi65 >> ((4294967295) & 63)) {
                                                            case 2147483647:
                                                                v4 = v4
                                                            case 963:
                                                                v79 := GenS1{F0: 1, F1: ((!true) && (1e10 <= -2.25)), F2: 42}
                                                                _ = v79
                                                            case 32767:
                                                                var v80 uint64 = 0
                                                                _ = v80
                                                            default:
                                                                if (1 < 799) {
                                                                    v8 = v8
                                                                } else {
                                                                    v81 := GenS7{F0: 18446744073709551615, F1: (v4 / (((v4 << ((4294967295) & 63))) | 1)), F2: 1}
                                                                    delete(v3, v8)
                                                                    if (false && (-32768 >= 1000)) {
                                                                        if (!v63) {
                                                                            v78 += 9223372036854775807
                                                                        }
                                                                    }
                                                                    _ = v81
                                                                }
                                                            }
                                                            _ = v78
                                                        case "a":
                                                            var v82 uint = 0
                                                            _ = v82
                                                        case "hello":
                                                            v83 := GenS2{F0: false, F1: 649, F2: 42}
                                                            _ = v83
                                                        default:
                                                            switch fi65 {
                                                            case -591:
                                                                if false {
                                                                    v84 := GenS8{F0: -193}
                                                                    _ = v84
                                                                } else {
                                                                    if v63 {
                                                                        v12 += v12
                                                                        v12 = (-(v12))
                                                                        var v85 uint64 = 4294967295
                                                                        fmt.Println(v63)
                                                                        var v86 float32 = ((((-2.25 / (0.1 / float32(0.1))) + 0.0) / (1.5 + float32(1.5))) + 0.1)
                                                                        _ = v85
                                                                        _ = v86
                                                                    } else {
                                                                        v87 := GenS6{F0: 42, F1: 0, F2: 65535}
                                                                        _ = v87
                                                                    }
                                                                    v13 -= 4294967295
                                                                    v88 := GenS9{F0: 18446744073709551615, F1: -899, F2: 61}
                                                                    _ = v88
                                                                }
                                                                var v89 bool = (411 <= 799)
                                                                if len(v2) > 0 {
                                                                    v2[0] = v8
                                                                }
                                                                _ = v89
                                                            case 0:
                                                                if ((v12 | 1) < (v12 * (v12 - (^(v12))))) {
                                                                    _ = len(v2)
                                                                }
                                                            case 2147483647:
                                                                if (882 == 274) {
                                                                    v90 := GenS10{F0: 18446744073709551615}
                                                                    var v91 string = v8
                                                                    _ = v90
                                                                    _ = v91
                                                                } else {
                                                                    var v92 uint8 = 252
                                                                    v63 = v63
                                                                    _ = v92
                                                                }
                                                                var v93 uint8 = 255
                                                                switch v8 {
                                                                case "foo bar":
                                                                    var v94 uint32 = 2147483648
                                                                    _ = v94
                                                                }
                                                                _ = v93
                                                            }
                                                        }
                                                    }
                                                    if v63 {
                                                        if (!v63) {
                                                            fmt.Println(v4, fi65)
                                                            v95 := GenS10{F0: 42}
                                                            _ = v95
                                                        }
                                                    }
                                                } else {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    fmt.Println(fi11, v8, v12)
                                                }
                                            } else {
                                                v13 = int(v12)
                                                var v96 int = (fi11 * 471)
                                                fmt.Println(fi11)
                                                switch v8 {
                                                case "":
                                                    var v97 bool = true
                                                    _ = v97
                                                case "foo bar":
                                                    v13 += fi11
                                                case "a":
                                                    v98 := v1[912]
                                                    _ = v98
                                                default:
                                                    var v99 uint8 = 42
                                                    _ = v99
                                                }
                                                v8 = v8
                                                var v100 float32 = float32(1.5)
                                                _ = v100
                                                _ = v96
                                            }
                                            if v63 {
                                                if true {
                                                    var v101 float32 = (3.14159 - 1.5)
                                                    if v63 {
                                                        v8 = (v8 + (v8 + v8))
                                                    } else {
                                                        v102 := GenS3{F0: v4}
                                                        fmt.Println(v4)
                                                        fmt.Println(v101, v8, v8)
                                                        _ = v102
                                                    }
                                                    fmt.Println(v13)
                                                    _ = v101
                                                }
                                            }
                                            switch v8 {
                                            case "a":
                                                v63 = false
                                            default:
                                                v3[v8] = 374
                                            }
                                            var v103 int64 = 42
                                            var v104 int64 = (-(v103))
                                            v8 = ""
                                            _ = v103
                                            _ = v104
                                        } else {
                                            switch (-(fi11)) {
                                            case 128:
                                                var v105 bool = false
                                                _ = v105
                                            case 0:
                                                if v63 {
                                                    if true {
                                                        fmt.Println(fi11)
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            var v106 int32 = (v4 * (v4 ^ 2147483647))
                                                            _ = v106
                                                        }
                                                        v12 = v12
                                                    }
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            case 2147483647:
                                                var v107 int16 = 1000
                                                _ = v107
                                            default:
                                                if true {
                                                    v12 = (v12 / (((^(v12))) | 1))
                                                    v8 = "a"
                                                } else {
                                                    v13 = (-(v13))
                                                }
                                            }
                                            if (float64(1.5) <= 3.14159) {
                                                switch (^(v13)) {
                                                case 167:
                                                    v12 = 42
                                                }
                                            } else {
                                                v4 += 585
                                                var v108 int = v13
                                                _ = v108
                                            }
                                            var v109 uint16 = 920
                                            _ = v109
                                        }
                                        _ = fi65
                                        _ = v70
                                    }
                                }
                            } else {
                                if (v8 != ("foo bar" + (v8 + "a"))) {
                                    v110 := GenS11{F0: "foo bar", F1: 0}
                                    v13 = v13
                                    delete(v1, 2147483648)
                                    _ = v110
                                }
                                var v111 uint = 647
                                _ = v111
                            }
                            _ = v63
                        } else {
                            v2 = append(v2, v8)
                            v13 += v13
                            v8 = v8
                        }
                        v12 = v12
                        v12 += v12
                    } else {
                        var v112 int32 = 42
                        v8 = ""
                        _ = v112
                    }
                    var v113 float32 = (3.14159 + 1.5)
                    v4 = (v4 & (-(v4)))
                    _ = v113
                }
            }
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi114 := 0; fi114 < 11; fi114++ {
        fmt.Println(fi114, fi114)
        _ = fi114
    }
    v115 := GenS12{F0: 1, F1: 1, F2: 718}
    for fi116 := 0; fi116 < 16; fi116++ {
        var v117 int64 = -1
        var v118 int16 = -207
        fmt.Println(v4, v118, v117)
        _ = fi116
        _ = v117
        _ = v118
    }
    var v119 uint32 = 42
    switch -734 {
    case -1:
        switch (v8 + "foo bar") {
        case "a":
            var v120 uint64 = 4294967295
            _ = v120
        }
    }
    if (v119 > v119) {
        var v121 int8 = 1
        v4 += 42
        _ = v121
    } else {
        if false {
            v119 = (v119 & (v119 & (v119 & v119)))
            v122 := GenS13{F0: 1, F1: 32767, F2: 0}
            _ = v122
        } else {
            for fi123 := 0; fi123 < 13; fi123++ {
                for fi124 := 0; fi124 < 10; fi124++ {
                    v125 := GenS13{F0: (^(v119)), F1: -1, F2: 65535}
                    v126 := GenS11{F0: "a", F1: 1}
                    var v127 int = 2147483647
                    _ = fi124
                    _ = v125
                    _ = v126
                    _ = v127
                }
                v128 := GenS14{F0: -1}
                _ = fi123
                _ = v128
            }
            var v129 uint64 = 42
            fmt.Println(v4)
            _ = v129
        }
        for fi130 := 0; fi130 < 8; fi130++ {
            v119 += 42
            v131 := GenS15{F0: (fi130 << ((42) & 63))}
            if len(v2) > 0 {
                v2[0] = v8
            }
            _ = fi130
            _ = v131
        }
        fmt.Println(v119, v119)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v119:", v119)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v115
    _ = v119
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
