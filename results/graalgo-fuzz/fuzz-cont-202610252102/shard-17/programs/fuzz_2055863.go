package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint; F2 int16 }
type GenS4 struct { F0 int64; F1 bool; F2 int64 }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 int32; F1 int64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint; F1 uint32; F2 int32 }
type GenS11 struct { F0 int8; F1 int64 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int; F1 string; F2 uint32 }
type GenS15 struct { F0 int32; F1 string }
type GenS16 struct { F0 uint; F1 string; F2 int }
type GenS17 struct { F0 int8; F1 string; F2 bool }
func genFn0(p0 uint32, p1 int64) uint {
    fmt.Println(p0)
    var v5 int16 = -1
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            if (("foo bar" + ("\n\t" + ("hello" + "\n\t"))) > ("foo bar" + "a")) {
                if (!((1.5 < -2.25) || (v5 == (v5 / ((v5) | 1))))) {
                    switch "" {
                    case "\n\t":
                        v8 := GenS1{F0: 1, F1: 191}
                        _ = v8
                    case "a":
                        switch fi7 {
                        case 32768:
                            var v9 uint16 = 255
                            if true {
                                if ((true || ("foo bar" < "foo bar")) || true) {
                                    v10 := GenS2{F0: 244}
                                    _ = v10
                                }
                                if true {
                                    if true {
                                        v11 := []int{(fi6 << ((uint(p0)) & 63))}
                                        _ = v11
                                    } else {
                                        var v12 int64 = p1
                                        _ = v12
                                    }
                                    switch "a" {
                                    case "":
                                        v9 -= (^(v9))
                                    case "hello":
                                        v13 := GenS3{F0: 0, F1: 65535, F2: (v5 >> ((2) & 63))}
                                        _ = v13
                                    case "foo bar":
                                        p0 += 42
                                    }
                                    var v14 string = ("hello" + "foo bar")
                                    _ = v14
                                } else {
                                    if false {
                                        var v15 int64 = 42
                                        _ = v15
                                    } else {
                                        var v16 int32 = 1
                                        _ = v16
                                    }
                                    var v17 uint = 0
                                    var v18 float64 = (1e10 / 1.0)
                                    _ = v17
                                    _ = v18
                                }
                                v5 -= 163
                            }
                            _ = v9
                        default:
                            var v19 string = ("a" + "")
                            _ = v19
                        }
                    case "foo bar":
                        p1 -= p1
                    default:
                        p1 += (p1 & p1)
                    }
                    var v20 int = 516
                    var v21 string = ("hello" + (("foo bar" + "foo bar") + ""))
                    if ((v5 % ((v5) | 1)) == v5) {
                        v22 := GenS1{F0: 148, F1: 18446744073709551615}
                        if true {
                            v23 := GenS1{F0: 1, F1: 18446744073709551615}
                            var v24 uint64 = 4294967295
                            v20 = v20
                            _ = v23
                            _ = v24
                        }
                        _ = v22
                    }
                    p1 += (^(p1))
                    _ = v20
                    _ = v21
                } else {
                    if true {
                        p1 -= 599
                        v25 := GenS3{F0: -1, F1: 4294967295, F2: 1}
                        _ = v25
                    }
                    var v26 uint8 = 92
                    _ = v26
                }
                v27 := []int{1}
                _ = v27
            }
            fmt.Println(fi6, p0, p0)
            _ = fi7
        }
        for fi28 := 0; fi28 < 3; fi28++ {
            v29 := GenS4{F0: (p1 % (((p1 + p1)) | 1)), F1: (0.0 >= (0.1 / ((-2.25 * 0.1) / 1e10))), F2: p1}
            var v30 int32 = int32(v5)
            if (v30 != 1) {
                p1 = int64(p0)
            }
            _ = fi28
            _ = v29
            _ = v30
        }
        var v31 float32 = float32(0.1)
        _ = fi6
        _ = v31
    }
    v32 := GenS5{F0: 385, F1: false}
    v33 := []int{128, -963, -854}
    _ = p0
    _ = p1
    _ = v32
    _ = v33
    _ = v5
    return 18446744073709551615
}

func main() {
    var v1 float32 = -2.25
    var v2 uint16 = 255
    var v3 uint16 = v2
    var v4 bool = (!(-1 != 441))
    v34 := []int{32768, 42, 937}
    var v35 string = ("a" + (("a" + "a") + ""))
    switch (v35 + "a") {
    case "hello":
        fmt.Println(v3)
    case "\n\t":
        for fi36 := 0; fi36 < 2; fi36++ {
            v2 = v3
            v37 := GenS6{F0: -332, F1: -581}
            fmt.Println(v1, v1)
            _ = fi36
            _ = v37
        }
    }
    v2 = v3
    for fi38 := 0; fi38 < 10; fi38++ {
        if (!v4) {
            switch v35 {
            case "a":
                v39 := GenS7{F0: 611}
                _ = v39
            case "hello":
                var v40 string = v35
                _ = v40
            case "":
                v41 := GenS2{F0: 2}
                _ = v41
            }
            for fi42 := 0; fi42 < 17; fi42++ {
                if v4 {
                    v43 := GenS4{F0: -24, F1: (false && (3.14159 >= 3.14159)), F2: 9223372036854775807}
                    var v44 float32 = v1
                    v3 = v2
                    _ = v43
                    _ = v44
                }
                if (-6 != -123) {
                    v2 += v3
                    if (!(((v2 >> ((306) & 63)) >= v2) && (true && false))) {
                        v4 = (!true)
                        _ = len(v34)
                    }
                    if (!(!(true || v4))) {
                        v45 := GenS8{F0: v4}
                        if (v4 && false) {
                            var v46 int16 = -32768
                            _ = v46
                        }
                        switch v35 {
                        case "foo bar":
                            v35 = v35
                        case "a":
                            v47 := GenS5{F0: v2, F1: v4}
                            if len(v34) >= 2 {
                                v34 = v34[1:len(v34)]
                            }
                            v48 := GenS4{F0: -9223372036854775808, F1: ((fi42 * (fi42 >> ((912) & 63))) > (fi42 >> ((117) & 63))), F2: 649}
                            v49 := GenS1{F0: 178, F1: 1}
                            _ = v47
                            _ = v48
                            _ = v49
                        default:
                            switch v35 {
                            case "":
                                v50 := GenS2{F0: 0}
                                _ = v50
                            case "\n\t":
                                switch (fi42 / (((fi42 >> ((751) & 63))) | 1)) {
                                case -2147483648:
                                    var v51 uint32 = 1
                                    _ = v51
                                default:
                                    v2 = 731
                                }
                            case "a":
                                v1 -= v1
                            default:
                                v35 = "\n\t"
                            }
                        }
                        switch (^(fi42)) {
                        case 32767:
                            fmt.Println(v2, v4, v4)
                        }
                        _ = v45
                    }
                    if (2 != 1) {
                        if false {
                            v3 -= (v2 << ((87) & 63))
                            v3 = (v3 * v2)
                            v4 = v4
                        } else {
                            v3 = 77
                            _ = len(v34)
                            v4 = (2 >= 18446744073709551615)
                        }
                        if v4 {
                            fmt.Println(fi42, fi38, v35)
                            var v52 int8 = 1
                            if len(v34) >= 2 {
                                v34 = v34[1:len(v34)]
                            }
                            _ = v52
                        } else {
                            v53 := GenS7{F0: genFn0(2147483647, 9223372036854775807)}
                            if (-2.25 >= 0.0) {
                                if v4 {
                                    switch (fi42 & 233) {
                                    case 32768:
                                        v54 := GenS6{F0: -1, F1: 0}
                                        _ = v54
                                    case -182:
                                        if len(v34) > 0 {
                                            v34[0] = (^(fi38))
                                        }
                                    default:
                                        v2 -= v2
                                    }
                                    var v55 uint = 255
                                    if (false || v4) {
                                        v3 = v3
                                        v4 = v4
                                    }
                                    _ = v55
                                } else {
                                    v2 -= (v2 ^ v3)
                                    v56 := GenS6{F0: 2147483647, F1: 0}
                                    if (!v4) {
                                        if len(v34) > 0 {
                                            v34[0] = fi42
                                        }
                                    } else {
                                        v3 += 255
                                        if v4 {
                                            var v57 float32 = 0.1
                                            var v58 int64 = 1
                                            _ = v57
                                            _ = v58
                                        }
                                        var v59 int16 = -313
                                        _ = v59
                                    }
                                    _ = v56
                                }
                                var v60 int16 = 46
                                v61 := GenS5{F0: (v2 ^ v2), F1: (!((v1 - 1.5) == 3.14159))}
                                var v62 int16 = (^(v60))
                                v1 = (v1 * 0.1)
                                _ = v60
                                _ = v61
                                _ = v62
                            }
                            switch "hello" {
                            case "hello":
                                var v63 int8 = 27
                                _ = v63
                            case "a":
                                v4 = (true && ((0.0 > 1.5) && (-2.25 == -2.25)))
                            }
                            _ = v53
                        }
                    } else {
                        v4 = (209 == 0)
                        v34 = append(v34, 0)
                    }
                    var v64 bool = (v4 || v4)
                    var v65 uint16 = (v2 % ((v2) | 1))
                    _ = v64
                    _ = v65
                } else {
                    fmt.Println(v1, fi38)
                    if true {
                        fmt.Println(v1, fi38)
                        var v66 int32 = 87
                        var v67 uint32 = 0
                        v68 := GenS9{F0: v35}
                        fmt.Println(v1)
                        v4 = v4
                        _ = v66
                        _ = v67
                        _ = v68
                    }
                    if len(v34) > 0 {
                        v34[0] = (fi42 * fi42)
                    }
                }
                if len(v34) >= 2 {
                    v34 = v34[1:len(v34)]
                }
                _ = fi42
            }
            if (!(42 < 335)) {
                var v69 int16 = -359
                switch 0 {
                case 9223372036854775807:
                    var v70 float64 = 1.5
                    _ = v70
                default:
                    var v71 string = "\n\t"
                    _ = v71
                }
                for fi72 := 0; fi72 < 2; fi72++ {
                    var v73 int64 = 0
                    var v74 uint = 2
                    var v75 int32 = -1
                    _ = fi72
                    _ = v73
                    _ = v74
                    _ = v75
                }
                _ = v69
            } else {
                switch ("" + "foo bar") {
                case "":
                    v76 := GenS3{F0: 1, F1: 4294967295, F2: 32767}
                    _ = v76
                case "hello":
                    for fi77 := 0; fi77 < 12; fi77++ {
                        switch fi38 {
                        case 127:
                            var v78 int64 = -1
                            var v79 uint32 = 65535
                            fmt.Println(fi77, fi77)
                            switch v35 {
                            case "hello":
                                switch "foo bar" {
                                case "\n\t":
                                    _ = len(v34)
                                case "foo bar":
                                    if v4 {
                                        if true {
                                            if len(v34) > 0 {
                                                v34[0] = (fi38 | -4)
                                            }
                                            v78 -= v78
                                            v80 := GenS8{F0: (-991 > 2147483647)}
                                            _ = v80
                                        }
                                        v35 = v35
                                        v35 = "\n\t"
                                    }
                                case "hello":
                                    fmt.Println(v1, fi77)
                                }
                            case "foo bar":
                                fmt.Println(v3, v79, v4)
                                v34 = append(v34, (fi38 / ((fi77) | 1)))
                                v81 := GenS10{F0: 4294967295, F1: 42, F2: 0}
                                _ = v81
                            default:
                                if (((v35 + v35) + "foo bar") <= (((v35 + v35) + v35) + "")) {
                                    if true {
                                        switch ((v35 + v35) + (v35 + v35)) {
                                        case "a":
                                            if (((v35 < v35) && v4) || v4) {
                                                var v82 int16 = 1
                                                v35 = (v35 + "\n\t")
                                                _ = v82
                                            } else {
                                                v83 := GenS11{F0: -21, F1: v78}
                                                v78 -= (v78 + (v78 | 1))
                                                _ = v83
                                            }
                                        }
                                    } else {
                                        var v84 int64 = v78
                                        _ = v84
                                    }
                                } else {
                                    var v85 float64 = (float64(v1) + 0.1)
                                    if (v35 == v35) {
                                        if (uint8(v2) < 255) {
                                            var v86 uint32 = 2147483648
                                            v35 = (v35 + ("" + v35))
                                            fmt.Println(fi77)
                                            _ = v86
                                        } else {
                                            if (true || (v4 || ((4294967295 < (fi38 / ((-10) | 1))) || false))) {
                                                v79 = v79
                                                v85 = (((float64(v1) / (0.1 / 0.1)) * (v85 + 0.1)) * 3.14159)
                                                v78 += v78
                                                var v87 int64 = (v78 | (v78 % ((9223372036854775807) | 1)))
                                                _ = v87
                                            } else {
                                                v3 = 0
                                                if v4 {
                                                    var v88 float64 = v85
                                                    switch fi38 {
                                                    case 32767:
                                                        switch fi77 {
                                                        case -2147483648:
                                                            if v4 {
                                                                var v89 uint32 = v79
                                                                _ = v89
                                                            }
                                                        case 0:
                                                            v4 = (1e10 <= 1.5)
                                                        default:
                                                            if (!v4) {
                                                                switch 567 {
                                                                case -916:
                                                                    var v90 int8 = 103
                                                                    _ = v90
                                                                case 32767:
                                                                    var v91 bool = v4
                                                                    _ = v91
                                                                case 0:
                                                                    if (9223372036854775807 > fi77) {
                                                                        var v92 int16 = 32767
                                                                        _ = v92
                                                                    } else {
                                                                        var v93 uint = 2
                                                                        var v94 int8 = -1
                                                                        fmt.Println(v93, v93)
                                                                        var v95 int16 = int16(v94)
                                                                        _ = v93
                                                                        _ = v94
                                                                        _ = v95
                                                                    }
                                                                }
                                                                _ = len(v34)
                                                                v78 += 0
                                                            }
                                                        }
                                                    case -857:
                                                        if len(v34) > 0 {
                                                            v34[0] = fi77
                                                        }
                                                    case -982:
                                                        v34 = append(v34, fi77)
                                                        switch ("a" + "foo bar") {
                                                        case "\n\t":
                                                            v96 := GenS4{F0: v78, F1: true, F2: (-(v78))}
                                                            if v4 {
                                                                var v97 uint16 = 896
                                                                v98 := GenS5{F0: (v97 - v3), F1: true}
                                                                _ = v97
                                                                _ = v98
                                                            }
                                                            v88 += v88
                                                            _ = v96
                                                        case "":
                                                            v34 = append(v34, (fi77 ^ fi38))
                                                        }
                                                    }
                                                    v3 -= (v2 & 0)
                                                    var v99 bool = true
                                                    if ((-1 >= -1) || (v99 && v4)) {
                                                        _ = len(v34)
                                                        var v100 int16 = -765
                                                        switch fi77 {
                                                        case 1:
                                                            var v101 int16 = (v100 + v100)
                                                            _ = v101
                                                        case -50:
                                                            v100 += 1000
                                                        case 71:
                                                            _ = len(v34)
                                                        default:
                                                            var v102 int = (fi77 | (fi38 * (^(fi77))))
                                                            _ = v102
                                                        }
                                                        fmt.Println(v85, v88)
                                                        v103 := GenS12{F0: 0}
                                                        v104 := GenS13{F0: 214}
                                                        _ = v100
                                                        _ = v103
                                                        _ = v104
                                                    }
                                                    _ = v88
                                                    _ = v99
                                                }
                                            }
                                            var v105 float64 = (v85 / ((v85 / (v85 / (((0.1 / (float64((v1 * (-2.25 - (float32(((3.14159 - (((1e10 - v85) - v85) / (v85 / v85))) + 0.1)) / v1)))) / (((v85 + 0.1) * v85) + v85))) * v85) * (((v85 * (3.14159 * (1.5 - 1.5))) + (float64(((-2.25 - ((v1 / v1) - v1)) * 0.1)) - v85)) / (((v85 / v85) * (v85 * float64(0.0))) * v85))))) / v85))
                                            v34 = append(v34, 4294967295)
                                            if v4 {
                                                v2 = v3
                                                var v106 uint64 = 52
                                                v78 = v78
                                                _ = v106
                                            }
                                            _ = v105
                                        }
                                        if len(v34) >= 2 {
                                            v34 = v34[1:len(v34)]
                                        }
                                        if (float64(0.0) != (float64(0.1) + v85)) {
                                            if v4 {
                                                var v107 bool = (2147483648 != 282)
                                                v108 := GenS8{F0: (951 < -1)}
                                                _ = v107
                                                _ = v108
                                            }
                                            v78 = v78
                                        }
                                        if v4 {
                                            var v109 bool = (v1 == (0.0 - ((v1 / v1) / (-2.25 + v1))))
                                            var v110 string = v35
                                            var v111 uint8 = uint8(v79)
                                            v3 -= v2
                                            v4 = (!false)
                                            _ = v109
                                            _ = v110
                                            _ = v111
                                        } else {
                                            v35 = v35
                                            v2 += (v3 - (^(v2)))
                                            switch ("" + "foo bar") {
                                            case "hello":
                                                v1 += 1e10
                                            }
                                        }
                                        switch (v35 + ("hello" + (v35 + v35))) {
                                        case "":
                                            var v112 int32 = 0
                                            _ = v112
                                        case "a":
                                            v4 = v4
                                        case "foo bar":
                                            switch fi38 {
                                            case 32768:
                                                switch fi77 {
                                                case 4294967295:
                                                    v113 := GenS14{F0: 0, F1: v35, F2: v79}
                                                    v114 := GenS2{F0: 358}
                                                    v2 = (^(v2))
                                                    _ = v113
                                                    _ = v114
                                                case -524:
                                                    var v115 uint16 = v3
                                                    _ = v115
                                                }
                                            }
                                        }
                                        switch (fi38 + (fi38 & fi77)) {
                                        case 128:
                                            if false {
                                                fmt.Println(fi38, v35)
                                                v116 := GenS11{F0: 127, F1: 0}
                                                if v4 {
                                                    if (!(-32768 <= -32768)) {
                                                        if true {
                                                            if false {
                                                                var v117 int16 = 1000
                                                                v34 = append(v34, 1)
                                                                _ = v117
                                                            }
                                                            v78 -= v78
                                                            v118 := GenS14{F0: fi77, F1: (v35 + v35), F2: v79}
                                                            _ = v118
                                                        }
                                                    } else {
                                                        fmt.Println(v3, v3, v3)
                                                        v119 := GenS13{F0: uint8(v3)}
                                                        _ = v119
                                                    }
                                                    var v120 int16 = 0
                                                    _ = v120
                                                }
                                                _ = v116
                                            }
                                        case -598:
                                            fmt.Println(v2, v85)
                                        }
                                    }
                                    _ = v85
                                }
                            }
                            switch ((v35 + "foo bar") + v35) {
                            case "hello":
                                v121 := GenS13{F0: 127}
                                _ = len(v34)
                                var v122 int8 = 42
                                v123 := GenS7{F0: 922}
                                v35 = (v35 + v35)
                                _ = v121
                                _ = v122
                                _ = v123
                            }
                            var v124 bool = v4
                            _ = v124
                            _ = v78
                            _ = v79
                        case 52:
                            var v125 int64 = -1
                            _ = len(v34)
                            fmt.Println(fi77, v1, v4)
                            _ = v125
                        case 32768:
                            fmt.Println(fi77, fi38)
                            v4 = true
                            v3 = 0
                            if false {
                                v1 = (1.5 / 1.0)
                                var v126 bool = (-1 <= 42)
                                if len(v34) > 0 {
                                    v34[0] = fi38
                                }
                                var v127 bool = (false && v126)
                                v126 = v127
                                var v128 string = ("hello" + ((("a" + (v35 + "a")) + "\n\t") + "\n\t"))
                                var v129 int32 = 123
                                var v130 uint64 = 2
                                var v131 int32 = v129
                                switch fi77 {
                                case 32767:
                                    v3 -= v2
                                case 9223372036854775807:
                                    if v126 {
                                        var v132 uint32 = uint32(v2)
                                        _ = len(v34)
                                        switch "hello" {
                                        case "":
                                            if false {
                                                fmt.Println(v127)
                                                switch v128 {
                                                case "a":
                                                    if v126 {
                                                        var v133 float64 = 0.0
                                                        v34 = append(v34, (fi38 + 2147483647))
                                                        var v134 uint64 = 18446744073709551615
                                                        _ = v133
                                                        _ = v134
                                                    } else {
                                                        if len(v34) >= 2 {
                                                            v34 = v34[1:len(v34)]
                                                        }
                                                        var v135 string = ("\n\t" + "hello")
                                                        v3 += v2
                                                        _ = v135
                                                    }
                                                    var v136 int32 = v131
                                                    fmt.Println(v2, v127)
                                                    var v137 string = v128
                                                    _ = v136
                                                    _ = v137
                                                case "\n\t":
                                                    var v138 int32 = 85
                                                    _ = v138
                                                default:
                                                    if ((false && ((v132 >> ((0) & 63)) < v132)) || v126) {
                                                        var v139 int = fi77
                                                        switch (v128 + (v128 + ((v128 + (v128 + v35)) + ("a" + v128)))) {
                                                        case "\n\t":
                                                            if len(v34) > 0 {
                                                                v34[0] = (fi77 >> ((65535) & 63))
                                                            }
                                                        case "":
                                                            if len(v34) > 0 {
                                                                v34[0] = (fi38 + fi38)
                                                            }
                                                        }
                                                        _ = v139
                                                    } else {
                                                        if v126 {
                                                            switch ("foo bar" + (((v128 + v35) + v35) + (v35 + (v35 + v35)))) {
                                                            case "hello":
                                                                v140 := GenS6{F0: 55, F1: -1}
                                                                _ = v140
                                                            default:
                                                                _ = len(v34)
                                                            }
                                                        } else {
                                                            switch "" {
                                                            case "":
                                                                var v141 int8 = 42
                                                                _ = v141
                                                            case "hello":
                                                                switch -1 {
                                                                case 9223372036854775807:
                                                                    var v142 uint64 = v130
                                                                    _ = v142
                                                                case 99:
                                                                    v1 += 3.14159
                                                                default:
                                                                    fmt.Println(v2, fi77, v129)
                                                                }
                                                            case "foo bar":
                                                                if true {
                                                                    v3 += (v2 * 0)
                                                                    if (!(2 == 1)) {
                                                                        var v143 int16 = 32767
                                                                        fmt.Println(v129, v129, v35)
                                                                        if len(v34) >= 2 {
                                                                            v34 = v34[1:len(v34)]
                                                                        }
                                                                        _ = v143
                                                                    }
                                                                    if v126 {
                                                                        var v144 uint16 = v2
                                                                        _ = v144
                                                                    }
                                                                }
                                                            default:
                                                                v128 = "\n\t"
                                                            }
                                                            switch (fi38 >> ((42) & 63)) {
                                                            case -1:
                                                                var v145 string = v128
                                                                _ = v145
                                                            case 992:
                                                                var v146 int32 = v129
                                                                _ = v146
                                                            case -769:
                                                                v147 := GenS15{F0: -2147483648, F1: v128}
                                                                _ = v147
                                                            }
                                                            v34 = append(v34, -240)
                                                            var v148 uint8 = 255
                                                            var v149 int64 = 9223372036854775807
                                                            _ = v148
                                                            _ = v149
                                                        }
                                                    }
                                                }
                                                if (1 != 1) {
                                                    switch (fi77 | fi77) {
                                                    case -930:
                                                        v150 := GenS6{F0: (v131 + -346), F1: -540}
                                                        _ = v150
                                                    default:
                                                        var v151 bool = (v126 && ((942 < 2) && v127))
                                                        _ = v151
                                                    }
                                                    v126 = (!false)
                                                } else {
                                                    switch v35 {
                                                    case "hello":
                                                        v152 := GenS11{F0: -20, F1: 42}
                                                        _ = v152
                                                    case "foo bar":
                                                        v130 = (v130 / ((uint64(v132)) | 1))
                                                    case "a":
                                                        var v153 float32 = (0.1 + ((-2.25 + v1) * v1))
                                                        _ = v153
                                                    }
                                                }
                                            } else {
                                                if len(v34) > 0 {
                                                    v34[0] = -759
                                                }
                                                var v154 int8 = -116
                                                _ = v154
                                            }
                                        case "foo bar":
                                            v34 = append(v34, (fi38 - fi77))
                                        }
                                        _ = v132
                                    } else {
                                        if (!v4) {
                                            v2 += 0
                                        } else {
                                            fmt.Println(v130)
                                            var v155 float32 = 0.0
                                            _ = v155
                                        }
                                        if v4 {
                                            if len(v34) > 0 {
                                                v34[0] = fi38
                                            }
                                            v4 = false
                                            v156 := GenS10{F0: 81, F1: uint32(v130), F2: v131}
                                            _ = v156
                                        } else {
                                            if ((v131 * -1) <= (v131 >> ((900) & 63))) {
                                                var v157 float32 = v1
                                                if (0.0 > float64(v1)) {
                                                    var v158 int16 = 501
                                                    v34 = append(v34, fi38)
                                                    _ = v158
                                                } else {
                                                    var v159 string = "a"
                                                    var v160 uint32 = 1
                                                    _ = v159
                                                    _ = v160
                                                }
                                                _ = v157
                                            } else {
                                                var v161 uint64 = v130
                                                if v126 {
                                                    var v162 string = ((v128 + (v35 + (("hello" + ("a" + "foo bar")) + "hello"))) + (v35 + (v128 + v35)))
                                                    var v163 uint32 = 1
                                                    _ = v162
                                                    _ = v163
                                                } else {
                                                    var v164 uint = uint(v3)
                                                    _ = v164
                                                }
                                                _ = v161
                                            }
                                        }
                                    }
                                }
                                var v165 string = (("hello" + (((v35 + v35) + "hello") + (v35 + v128))) + "\n\t")
                                _ = v126
                                _ = v127
                                _ = v128
                                _ = v129
                                _ = v130
                                _ = v131
                                _ = v165
                            } else {
                                if len(v34) > 0 {
                                    v34[0] = (fi38 | fi77)
                                }
                                var v166 int32 = 42
                                _ = v166
                            }
                        default:
                            switch ("foo bar" + (v35 + (v35 + v35))) {
                            case "hello":
                                var v167 bool = v4
                                _ = v167
                            case "\n\t":
                                var v168 uint32 = 65535
                                _ = v168
                            case "a":
                                if true {
                                    if len(v34) > 0 {
                                        v34[0] = fi77
                                    }
                                } else {
                                    if (0 >= 0) {
                                        v169 := GenS5{F0: v3, F1: false}
                                        _ = v169
                                    } else {
                                        v4 = (((!(float32((0.1 * float64(v1))) < float32(1e10))) && false) && false)
                                        v4 = v4
                                    }
                                }
                            default:
                                var v170 float64 = 1.5
                                _ = v170
                            }
                        }
                        v1 = (v1 * 0.1)
                        if ((false || true) || (!(1.5 < float32(0.1)))) {
                            v2 = (v3 ^ v2)
                            var v171 bool = ((-1 != -1) || (v4 || v4))
                            _ = v171
                        }
                        _ = fi77
                    }
                case "a":
                    if (v4 || v4) {
                        v2 = (v2 | (^(v3)))
                        if len(v34) >= 2 {
                            v34 = v34[1:len(v34)]
                        }
                    }
                }
                for fi172 := 0; fi172 < 16; fi172++ {
                    v1 = v1
                    if ((0.1 - 1e10) <= 0.1) {
                        v173 := GenS11{F0: -114, F1: 9223372036854775807}
                        _ = v173
                    }
                    v4 = false
                    _ = fi172
                }
                var v174 float64 = ((0.0 + 1e10) + -2.25)
                _ = v174
            }
        }
        var v175 int16 = -32768
        v2 = v3
        _ = fi38
        _ = v175
    }
    v4 = false
    _ = len(v34)
    var v176 int16 = 645
    for fi177 := 0; fi177 < 4; fi177++ {
        if (0 != -109) {
            var v178 int32 = 0
            _ = v178
        } else {
            fmt.Println(v176, v2, fi177)
            fmt.Println(v2, v2)
        }
        _ = fi177
    }
    var v179 string = (v35 + "")
    var v180 string = ("foo bar" + v35)
    if false {
        if (564 >= 702) {
            for fi181 := 0; fi181 < 11; fi181++ {
                var v182 bool = ((v179 + "") > ("\n\t" + ((("hello" + ((v35 + v35) + "foo bar")) + v179) + v180)))
                v183 := GenS16{F0: 0, F1: v35, F2: -395}
                _ = fi181
                _ = v182
                _ = v183
            }
            v176 += -717
        }
        v35 = v179
        v184 := GenS11{F0: 127, F1: 0}
        _ = v184
    } else {
        _ = len(v34)
    }
    v185 := GenS17{F0: -128, F1: (("hello" + ((("foo bar" + v179) + v180) + v179)) + ""), F2: (1 == 54)}
    v176 = (v176 >> ((272) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v176:", v176)
    fmt.Println("v179:", v179)
    fmt.Println("v180:", v180)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v176
    _ = v179
    _ = v180
    _ = v185
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
