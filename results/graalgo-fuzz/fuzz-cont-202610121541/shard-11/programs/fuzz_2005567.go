package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 int64; F1 uint32; F2 int16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 bool; F1 uint; F2 int32 }
type GenS9 struct { F0 int16; F1 uint32; F2 int8 }
type GenS10 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS11 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS12 struct { F0 uint32; F1 int16 }
type GenS13 struct { F0 int32; F1 uint }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 string }
func genFn0(p0 uint32, p1 int, p2 int8) int8 {
    v4 := []int{(p1 * -2147483648), (p1 & (p1 ^ (p1 + p1))), (p1 | (p1 / (((p1 | p1)) | 1))), (p1 % ((p1) | 1))}
    var v5 int = 9223372036854775807
    var v6 uint32 = p0
    _ = len(v4)
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return p2
}

func main() {
    v1 := []string{"foo bar", ""}
    var v2 int8 = 100
    var v3 int32 = 1
    v7 := GenS1{F0: 1, F1: 65535}
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 uint16 = 655
            fmt.Println(v2)
            v3 = (v3 << ((349) & 63))
            _ = fi9
            _ = v10
        }
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            if (!(1 > 65535)) {
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                var v11 uint = 0
                _ = v11
            } else {
                switch (fi8 % ((-751) | 1)) {
                case -2147483648:
                    var v12 uint8 = 255
                    _ = v12
                case 4294967295:
                    v13 := GenS2{F0: (fi8 / ((367) | 1)), F1: 65535, F2: 255}
                    _ = v13
                case 32768:
                    var v14 string = "foo bar"
                    _ = v14
                }
                var v15 float32 = 0.1
                var v16 float32 = v15
                _ = v15
                _ = v16
            }
            if (!((843 < 0) || (int16(v3) > 0))) {
                for fi17 := 0; fi17 < 3; fi17++ {
                    var v18 bool = false
                    v19 := GenS1{F0: 32767, F1: 4294967295}
                    var v20 int8 = -1
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v20
                }
                var v21 int8 = 100
                if (0 <= 0) {
                    _ = len(v1)
                }
                _ = v21
            } else {
                v2 = (^(v2))
                var v22 string = (("a" + "") + "\n\t")
                _ = v22
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        for fi23 := 0; fi23 < 16; fi23++ {
            if false {
                if false {
                    v24 := GenS1{F0: 42, F1: 255}
                    switch ("foo bar" + (("a" + ("hello" + "hello")) + "a")) {
                    case "a":
                        switch "foo bar" {
                        case "":
                            var v25 string = ("foo bar" + "\n\t")
                            _ = v25
                        case "a":
                            switch ("foo bar" + "foo bar") {
                            case "hello":
                                var v26 int16 = -560
                                _ = v26
                            case "a":
                                fmt.Println(fi23, v3)
                            default:
                                fmt.Println(v3, fi8)
                            }
                        case "\n\t":
                            if false {
                                var v27 int16 = 0
                                var v28 float64 = ((0.0 - -2.25) + (-2.25 * 0.1))
                                v1 = append(v1, ("hello" + "hello"))
                                _ = v27
                                _ = v28
                            }
                        default:
                            fmt.Println(v2, fi8, v3)
                        }
                    case "hello":
                        if (-586 > -1) {
                            if true {
                                _ = len(v1)
                                v3 = 42
                            }
                            var v29 int16 = -32768
                            _ = v29
                        } else {
                            v1 = append(v1, "hello")
                            v3 = (^(v3))
                            fmt.Println(fi23)
                            switch "\n\t" {
                            case "foo bar":
                                v30 := GenS3{F0: (v3 ^ (v3 ^ -1)), F1: 0}
                                var v31 bool = (((-9223372036854775808 == 1) || (!false)) || false)
                                _ = v30
                                _ = v31
                            case "":
                                if len(v1) > 0 {
                                    v1[0] = "foo bar"
                                }
                            }
                        }
                    }
                    switch (("\n\t" + "a") + ("a" + "\n\t")) {
                    case "":
                        switch fi23 {
                        case -19:
                            v2 = 127
                        case -1:
                            var v32 string = ((("\n\t" + "foo bar") + "\n\t") + "foo bar")
                            _ = v32
                        case 1:
                            v1 = append(v1, "hello")
                        default:
                            var v33 float64 = 0.1
                            _ = v33
                        }
                        if (42 <= 18446744073709551615) {
                            v34 := GenS3{F0: v3, F1: 569}
                            fmt.Println(v2)
                            _ = v34
                        }
                    default:
                        v35 := GenS2{F0: fi8, F1: 793, F2: 127}
                        _ = v35
                    }
                    switch (fi23 ^ (^(fi8))) {
                    case 242:
                        if false {
                            v36 := GenS1{F0: 675, F1: 735}
                            _ = v36
                        }
                    }
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(v3)
                    case "hello":
                        v37 := GenS4{F0: "hello", F1: 334}
                        _ = v37
                    default:
                        var v38 uint = 2
                        _ = v38
                    }
                    _ = v24
                } else {
                    switch "a" {
                    case "foo bar":
                        fmt.Println(fi8, v2)
                    case "":
                        if (427 >= 0) {
                            if true {
                                if len(v1) > 0 {
                                    v1[0] = "hello"
                                }
                            }
                            var v39 uint16 = 256
                            if true {
                                switch ("foo bar" + (("\n\t" + ((("" + "\n\t") + "\n\t") + "a")) + "\n\t")) {
                                case "hello":
                                    v40 := GenS4{F0: "\n\t", F1: (v3 / ((1) | 1))}
                                    _ = v40
                                }
                                v1 = append(v1, (("" + "\n\t") + ("hello" + "")))
                                fmt.Println(v2, v39, fi23)
                                v1 = append(v1, (("" + ("a" + ("" + "a"))) + "hello"))
                            }
                            _ = v39
                        }
                    default:
                        switch ("a" + (("a" + "hello") + (("a" + ("\n\t" + "hello")) + "hello"))) {
                        case "\n\t":
                            fmt.Println(v2, v2)
                        case "":
                            if (!false) {
                                if len(v1) > 0 {
                                    v1[0] = "\n\t"
                                }
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                            } else {
                                v1 = append(v1, "")
                                if true {
                                    v41 := GenS5{F0: 0, F1: fi8}
                                    var v42 uint64 = 185
                                    _ = v41
                                    _ = v42
                                } else {
                                    var v43 uint64 = 18446744073709551615
                                    if ((0 == 0) && true) {
                                        var v44 bool = true
                                        if (("foo bar" + "") >= (("" + "hello") + "\n\t")) {
                                            if false {
                                                if len(v1) > 0 {
                                                    v1[0] = ((("foo bar" + "\n\t") + "a") + "")
                                                }
                                            } else {
                                                v2 = v2
                                                var v45 string = ("" + "foo bar")
                                                var v46 int = 4294967295
                                                _ = v45
                                                _ = v46
                                            }
                                            fmt.Println(v43, fi23)
                                        } else {
                                            if (1 <= -9223372036854775808) {
                                                var v47 uint16 = 340
                                                var v48 bool = true
                                                _ = v47
                                                _ = v48
                                            }
                                            fmt.Println(v3, fi8, v44)
                                            var v49 int64 = 0
                                            fmt.Println(v2, v44, v49)
                                            _ = v49
                                        }
                                        if false {
                                            var v50 string = ""
                                            _ = v50
                                        } else {
                                            var v51 int32 = v3
                                            _ = v51
                                        }
                                        _ = v44
                                    } else {
                                        v2 = v2
                                        var v52 int16 = 42
                                        var v53 int16 = v52
                                        _ = v52
                                        _ = v53
                                    }
                                    _ = v43
                                }
                                v54 := GenS6{F0: -9223372036854775808, F1: 42, F2: -32768}
                                _ = v54
                            }
                        default:
                            var v55 float64 = 3.14159
                            v55 = ((v55 + float64(1.5)) + float64(0.0))
                            _ = v55
                        }
                    }
                    var v56 uint32 = 1
                    _ = v56
                }
                var v57 bool = false
                if (!false) {
                    v3 = v3
                    v1 = append(v1, "")
                    fmt.Println(v57, v2)
                    if v57 {
                        fmt.Println(v2, v3, fi23)
                        if false {
                            v1 = append(v1, "\n\t")
                            var v58 uint16 = 0
                            _ = v58
                        }
                        var v59 uint16 = 65535
                        _ = v59
                    } else {
                        if (v57 || ((4294967295 >= 35) && v57)) {
                            v2 = genFn0(2147483648, 983, v2)
                        }
                        var v60 int16 = 42
                        _ = v60
                    }
                } else {
                    v57 = v57
                    var v61 bool = v57
                    v3 -= v3
                    _ = v61
                }
                switch "a" {
                case "a":
                    var v62 string = "\n\t"
                    _ = v62
                case "\n\t":
                    switch 1 {
                    case 0:
                        v3 = (v3 - -843)
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                default:
                    var v63 float64 = (((1e10 * (1e10 - 1e10)) + float64(-2.25)) / (1.5 * 0.1))
                    _ = v63
                }
                _ = v57
            } else {
                v3 = -1
                v3 = (v3 / ((v3) | 1))
                v64 := GenS5{F0: 255, F1: fi23}
                _ = v64
            }
            v65 := GenS1{F0: 526, F1: 4294967295}
            _ = fi23
            _ = v65
        }
        _ = fi8
    }
    v3 = v3
    switch -628 {
    case 397:
        for fi66 := 0; fi66 < 13; fi66++ {
            v2 -= -24
            v3 -= (^(v3))
            if (!false) {
                for fi67 := 0; fi67 < 1; fi67++ {
                    if (1e10 <= (0.0 * 1.5)) {
                        v2 = (-(v2))
                        var v68 int16 = 32767
                        v69 := GenS4{F0: (((("foo bar" + ("\n\t" + "hello")) + "") + "a") + "\n\t"), F1: (^(v3))}
                        _ = v68
                        _ = v69
                    }
                    _ = fi67
                }
                if ((1.5 - (1e10 - 0.0)) <= 1e10) {
                    if false {
                        fmt.Println(fi66, v2, fi66)
                    }
                    fmt.Println(fi66, fi66, v3)
                }
                v2 = genFn0(2147483648, -2147483648, (v2 + genFn0(2147483648, fi66, (v2 + v2))))
            } else {
                v70 := GenS2{F0: (^(fi66)), F1: 65535, F2: 20}
                switch ((("\n\t" + "") + (("\n\t" + "foo bar") + "\n\t")) + "foo bar") {
                case "a":
                    switch ("\n\t" + "a") {
                    case "foo bar":
                        v71 := GenS7{F0: (fi66 % (((fi66 & -2147483648)) | 1))}
                        _ = v71
                    case "hello":
                        var v72 uint8 = 128
                        _ = v72
                    default:
                        v73 := GenS1{F0: -76, F1: 604}
                        _ = v73
                    }
                default:
                    if (128 >= 128) {
                        if (589 != -654) {
                            var v74 int32 = (v3 % (((v3 + (v3 / ((v3) | 1)))) | 1))
                            if (true && ("a" == "a")) {
                                switch ("hello" + "\n\t") {
                                case "foo bar":
                                    switch fi66 {
                                    case 365:
                                        v75 := GenS5{F0: 371, F1: 128}
                                        _ = v75
                                    case -409:
                                        v3 = v3
                                    case 42:
                                        v1 = append(v1, "a")
                                    default:
                                        var v76 bool = false
                                        _ = v76
                                    }
                                case "\n\t":
                                    v77 := GenS8{F0: (42 < 4294967295), F1: 492, F2: (v3 * v3)}
                                    _ = v77
                                }
                                var v78 uint = 425
                                _ = v78
                            }
                            v79 := GenS9{F0: 1, F1: 300, F2: -128}
                            _ = v74
                            _ = v79
                        } else {
                            v80 := GenS5{F0: 65535, F1: (fi66 & (fi66 / ((-842) | 1)))}
                            fmt.Println(fi66)
                            _ = v80
                        }
                    } else {
                        if false {
                            v2 += (v2 * (^(v2)))
                        }
                        switch "hello" {
                        case "hello":
                            var v81 int64 = -1
                            _ = v81
                        default:
                            for fi82 := 0; fi82 < 3; fi82++ {
                                v2 += v2
                                var v83 bool = true
                                var v84 string = "a"
                                _ = fi82
                                _ = v83
                                _ = v84
                            }
                        }
                    }
                }
                switch "\n\t" {
                case "hello":
                    if (245 < 692) {
                        v3 -= (v3 / ((804) | 1))
                        v85 := GenS9{F0: 32767, F1: 0, F2: (v2 << ((2) & 63))}
                        switch "a" {
                        case "hello":
                            v2 = (v2 | genFn0(2, fi66, (v2 % ((v2) | 1))))
                        }
                        _ = v85
                    }
                default:
                    if true {
                        var v86 int8 = v2
                        v87 := GenS4{F0: ("" + "a"), F1: v3}
                        _ = v86
                        _ = v87
                    }
                }
                _ = v70
            }
            _ = fi66
        }
    }
    if (0.1 == 1e10) {
        var v88 float32 = ((0.1 - (3.14159 - (((-2.25 / (1.5 + (1e10 * (1.5 * 0.0)))) * 0.0) + 0.0))) - 0.0)
        v3 -= v3
        _ = v88
    } else {
        if false {
            switch "\n\t" {
            case "foo bar":
                var v89 uint32 = 2147483648
                _ = v89
            }
            v3 = 2147483647
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        var v90 int = 42
        _ = v90
    }
    switch int(v3) {
    case -1:
        switch -628 {
        case 4294967295:
            switch "\n\t" {
            case "a":
                v91 := GenS1{F0: 1, F1: 0}
                _ = v91
            case "":
                if (32767 > 32767) {
                    var v92 int64 = 474
                    _ = v92
                } else {
                    _ = len(v1)
                }
            }
        case 2147483647:
            if ((1.5 == 0.0) || ((3.14159 * 3.14159) >= 1e10)) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                v3 -= (v3 * v3)
                if false {
                    v1 = append(v1, "foo bar")
                    switch (("a" + "\n\t") + "foo bar") {
                    case "hello":
                        var v93 uint64 = 1
                        _ = v93
                    case "a":
                        fmt.Println(v3)
                    default:
                        v94 := GenS10{F0: v2, F1: 1, F2: 2}
                        _ = v94
                    }
                    v2 = int8(v3)
                }
                v3 -= 2147483647
            }
        case 1:
            if (42 == 65535) {
                v2 += v2
            } else {
                v3 -= 42
                if (((-2.25 + (((1e10 + 0.0) * (1e10 * (1e10 * float32(1e10)))) + 1e10)) / 1.5) != float32((1e10 - (1.5 / (-2.25 + float64(1.5)))))) {
                    for fi95 := 0; fi95 < 14; fi95++ {
                        var v96 uint32 = 657
                        for fi97 := 0; fi97 < 1; fi97++ {
                            if false {
                                var v98 int8 = (^(v2))
                                v99 := GenS5{F0: 283, F1: fi95}
                                _ = v98
                                _ = v99
                            } else {
                                v3 = v3
                                var v100 int64 = -1
                                v101 := GenS8{F0: true, F1: 4294967295, F2: v3}
                                _ = v100
                                _ = v101
                            }
                            _ = fi97
                        }
                        v3 = v3
                        _ = fi95
                        _ = v96
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    v2 -= v2
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
        }
    case -9223372036854775808:
        v3 = (^(v3))
    default:
        fmt.Println(v3, v2)
    }
    var v102 uint64 = 42
    _ = len(v1)
    if (v102 == v102) {
        v103 := GenS6{F0: 1, F1: 42, F2: 0}
        for fi104 := 0; fi104 < 7; fi104++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v102 = 1
            var v105 uint = 153
            _ = fi104
            _ = v105
        }
        var v106 int8 = v2
        _ = v103
        _ = v106
    } else {
        if false {
            v107 := GenS6{F0: -9223372036854775808, F1: 65535, F2: 491}
            switch ((("foo bar" + "a") + "a") + "foo bar") {
            case "\n\t":
                fmt.Println(v2, v2, v2)
            case "foo bar":
                if false {
                    switch "" {
                    case "\n\t":
                        var v108 float32 = 0.0
                        _ = v108
                    case "":
                        v1 = append(v1, "foo bar")
                    case "hello":
                        var v109 uint32 = 65535
                        _ = v109
                    default:
                        var v110 int32 = v3
                        _ = v110
                    }
                } else {
                    v2 += v2
                    fmt.Println(v2, v2)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            case "":
                var v111 int64 = 9223372036854775807
                _ = v111
            default:
                for fi112 := 0; fi112 < 11; fi112++ {
                    var v113 uint32 = 818
                    if ((v3 - 42) != (-(v3))) {
                        v102 = v102
                        var v114 int64 = 9223372036854775807
                        v1 = append(v1, "a")
                        _ = v114
                    } else {
                        v2 -= v2
                    }
                    fmt.Println(v3)
                    _ = fi112
                    _ = v113
                }
            }
            if true {
                v1 = append(v1, (("" + "\n\t") + ""))
                v102 -= (^(v102))
            }
            _ = v107
        } else {
            if false {
                for fi115 := 0; fi115 < 15; fi115++ {
                    v2 = (v2 >> ((1) & 63))
                    if ("a" > ("\n\t" + "hello")) {
                        for fi116 := 0; fi116 < 17; fi116++ {
                            v117 := GenS11{F0: 137, F1: 32767, F2: 776}
                            _ = fi116
                            _ = v117
                        }
                    }
                    v118 := GenS2{F0: fi115, F1: 65535, F2: 15}
                    _ = fi115
                    _ = v118
                }
                if len(v1) > 0 {
                    v1[0] = ("" + "")
                }
            } else {
                var v119 uint16 = 42
                _ = v119
            }
            if true {
                fmt.Println(v102, v102)
                for fi120 := 0; fi120 < 16; fi120++ {
                    fmt.Println(fi120, v102)
                    if len(v1) > 0 {
                        v1[0] = ("a" + (("foo bar" + "foo bar") + (("\n\t" + "hello") + (((("hello" + "") + ("" + (("foo bar" + "") + (("a" + "hello") + "a")))) + "a") + ""))))
                    }
                    _ = fi120
                }
                var v121 int64 = -545
                _ = v121
            } else {
                if (true || (!false)) {
                    var v122 uint = 161
                    _ = v122
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v123 := GenS12{F0: 0, F1: 42}
                _ = v123
            }
            if (77 > v2) {
                var v124 int32 = v3
                _ = v124
            } else {
                var v125 int64 = 9223372036854775807
                v126 := GenS1{F0: 939, F1: 717}
                if true {
                    _ = len(v1)
                    var v127 uint8 = 255
                    for fi128 := 0; fi128 < 6; fi128++ {
                        if true {
                            for fi129 := 0; fi129 < 8; fi129++ {
                                if false {
                                    v3 += v3
                                }
                                var v130 float32 = (3.14159 - float32(3.14159))
                                _ = fi129
                                _ = v130
                            }
                            for fi131 := 0; fi131 < 18; fi131++ {
                                v127 = v127
                                switch "\n\t" {
                                case "hello":
                                    v132 := GenS8{F0: ((v102 | 2) <= 173), F1: 65535, F2: (v3 % (((v3 - v3)) | 1))}
                                    _ = v132
                                case "\n\t":
                                    if true {
                                        v3 = (v3 | v3)
                                        if false {
                                            switch fi128 {
                                            case 859:
                                                fmt.Println(fi128, v102)
                                            default:
                                                var v133 int8 = (v2 & -124)
                                                _ = v133
                                            }
                                            v134 := GenS9{F0: 917, F1: 2147483648, F2: v2}
                                            _ = v134
                                        } else {
                                            v135 := GenS8{F0: (false || ((false || true) && true)), F1: 18446744073709551615, F2: (v3 >> ((2) & 63))}
                                            if (65535 < 908) {
                                                switch "" {
                                                case "foo bar":
                                                    v136 := GenS12{F0: 42, F1: 1000}
                                                    if true {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    } else {
                                                        var v137 float64 = 3.14159
                                                        _ = v137
                                                    }
                                                    switch "" {
                                                    case "foo bar":
                                                        var v138 int8 = int8(v102)
                                                        _ = v138
                                                    case "\n\t":
                                                        fmt.Println(v127, fi131, v127)
                                                    default:
                                                        v102 = 4294967295
                                                    }
                                                    _ = v136
                                                case "a":
                                                    v3 += -315
                                                }
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            } else {
                                                v3 = (v3 >> ((18446744073709551615) & 63))
                                            }
                                            if (18446744073709551615 <= 2) {
                                                fmt.Println(v102)
                                            }
                                            _ = v135
                                        }
                                        _ = len(v1)
                                    } else {
                                        v2 = (v2 + 1)
                                    }
                                case "a":
                                    fmt.Println(fi131)
                                default:
                                    if (1.5 != ((-2.25 / -2.25) + -2.25)) {
                                        switch 0 {
                                        case -9223372036854775808:
                                            if false {
                                                if false {
                                                    v102 += (v102 >> ((1) & 63))
                                                    if len(v1) > 0 {
                                                        v1[0] = "foo bar"
                                                    }
                                                    var v139 uint = 4294967295
                                                    var v140 uint32 = 900
                                                    if ((v140 / ((v140) | 1)) > (v140 << ((v139) & 63))) {
                                                        v102 = v102
                                                        fmt.Println(v127)
                                                    } else {
                                                        switch "a" {
                                                        case "a":
                                                            var v141 float64 = ((1e10 / 3.14159) + -2.25)
                                                            _ = v141
                                                        case "hello":
                                                            v142 := GenS7{F0: fi128}
                                                            _ = v142
                                                        case "":
                                                            if ((^(v102)) < (v102 % ((0) | 1))) {
                                                                switch int(v2) {
                                                                case 32767:
                                                                    fmt.Println(v102, v140, v127)
                                                                }
                                                                switch "a" {
                                                                case "\n\t":
                                                                    switch "" {
                                                                    case "a":
                                                                        var v143 float32 = 0.0
                                                                        _ = v143
                                                                    case "":
                                                                        v3 += v3
                                                                    case "\n\t":
                                                                        switch "" {
                                                                        case "hello":
                                                                            if (true || true) {
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                v144 := GenS13{F0: (v3 >> (((v139 % ((v139) | 1))) & 63)), F1: (v139 >> ((v139) & 63))}
                                                                                v127 += 0
                                                                                _ = v144
                                                                            }
                                                                        default:
                                                                            fmt.Println(v139)
                                                                        }
                                                                    default:
                                                                        v145 := GenS14{F0: (v2 - v2)}
                                                                        _ = v145
                                                                    }
                                                                default:
                                                                    v127 += (^(v127))
                                                                }
                                                                fmt.Println(v102, fi131, v140)
                                                            }
                                                        }
                                                    }
                                                    var v146 int32 = 0
                                                    _ = v139
                                                    _ = v140
                                                    _ = v146
                                                }
                                                if (255 > 65535) {
                                                    _ = len(v1)
                                                }
                                                var v147 bool = ((^(v102)) != v102)
                                                _ = v147
                                            }
                                        case 2147483647:
                                            v127 -= (v127 % (((v127 | (v127 / (((^(v127))) | 1)))) | 1))
                                        case 1:
                                            var v148 float32 = 0.1
                                            _ = v148
                                        default:
                                            v149 := GenS8{F0: ((v125 >> ((42) & 63)) >= v125), F1: uint(v127), F2: int32(v102)}
                                            _ = v149
                                        }
                                        _ = len(v1)
                                        v1 = append(v1, "")
                                    } else {
                                        v125 -= int64(v102)
                                    }
                                }
                                var v150 int16 = 0
                                _ = fi131
                                _ = v150
                            }
                        } else {
                            fmt.Println(v127, v127)
                            var v151 uint32 = 42
                            _ = v151
                        }
                        _ = fi128
                    }
                    _ = v127
                } else {
                    v125 = v125
                    if false {
                        v152 := GenS11{F0: -682, F1: 255, F2: 0}
                        _ = v152
                    } else {
                        for fi153 := 0; fi153 < 16; fi153++ {
                            _ = len(v1)
                            v125 += v125
                            var v154 bool = ((^(v3)) > -1)
                            _ = fi153
                            _ = v154
                        }
                        fmt.Println(v125, v125, v102)
                    }
                    if ((-2.25 / 1.0) <= 0.0) {
                        for fi155 := 0; fi155 < 18; fi155++ {
                            v156 := GenS15{F0: 676}
                            if false {
                                v157 := GenS16{F0: ("a" + ("" + "hello"))}
                                var v158 int8 = v2
                                v159 := GenS12{F0: 42, F1: 32767}
                                _ = v157
                                _ = v158
                                _ = v159
                            }
                            _ = fi155
                            _ = v156
                        }
                        switch -2147483648 {
                        case 4294967295:
                            var v160 uint16 = 32767
                            _ = v160
                        default:
                            if len(v1) > 0 {
                                v1[0] = "\n\t"
                            }
                        }
                        if len(v1) > 0 {
                            v1[0] = ("hello" + "foo bar")
                        }
                    }
                }
                _ = v125
                _ = v126
            }
        }
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            switch "" {
            case "hello":
                var v161 int32 = 1
                _ = v161
            case "":
                fmt.Println(v102, v2)
            default:
                _ = len(v1)
            }
        }
    }
    switch 128 {
    case 128:
        fmt.Println(v2, v102, v3)
    case 540:
        v102 = 1
    default:
        switch -1 {
        case 0:
            for fi162 := 0; fi162 < 1; fi162++ {
                var v163 string = ("hello" + "a")
                _ = fi162
                _ = v163
            }
        }
    }
    v164 := GenS14{F0: 20}
    for fi165 := 0; fi165 < 1; fi165++ {
        switch ("\n\t" + ("a" + "foo bar")) {
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            var v166 int8 = 100
            _ = v166
        }
        if ((255 <= 18446744073709551615) && (489 < -907)) {
            for fi167 := 0; fi167 < 17; fi167++ {
                var v168 int16 = 1000
                var v169 uint8 = 0
                v102 = (v102 >> ((250) & 63))
                _ = fi167
                _ = v168
                _ = v169
            }
            switch ("hello" + "") {
            case "hello":
                for fi170 := 0; fi170 < 5; fi170++ {
                    if true {
                        var v171 int8 = v2
                        _ = v171
                    }
                    _ = fi170
                }
            case "a":
                v172 := GenS6{F0: 514, F1: 4294967295, F2: -294}
                _ = v172
            default:
                fmt.Println(v3, fi165)
            }
        } else {
            var v173 bool = ((fi165 - (fi165 << ((379) & 63))) == fi165)
            for fi174 := 0; fi174 < 8; fi174++ {
                v175 := GenS13{F0: v3, F1: 916}
                v102 = 4294967295
                _ = fi174
                _ = v175
            }
            fmt.Println(v2, v102)
            _ = v173
        }
        v3 -= int32(fi165)
        _ = fi165
    }
    var v176 int = -695
    v2 -= (v2 << ((171) & 63))
    v3 -= v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v102:", v102)
    fmt.Println("v176:", v176)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v102
    _ = v164
    _ = v176
    _ = v2
    _ = v3
    _ = v7
}
