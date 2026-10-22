package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint8 }
type GenS2 struct { F0 int64; F1 int32; F2 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 uint64; F1 bool; F2 int64 }
type GenS9 struct { F0 int; F1 int16; F2 int16 }
type GenS10 struct { F0 uint64; F1 string; F2 int8 }
type GenS11 struct { F0 uint16; F1 int16 }
type GenS12 struct { F0 uint; F1 uint32 }
type GenS13 struct { F0 int64; F1 int8; F2 int }
type GenS14 struct { F0 int64; F1 bool; F2 uint16 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 uint32; F1 uint16 }
type GenS17 struct { F0 uint64; F1 uint32 }
type GenS18 struct { F0 int16; F1 string }
type GenS19 struct { F0 uint8; F1 int16; F2 int16 }
type GenS20 struct { F0 int32; F1 uint; F2 uint32 }

func main() {
    var v1 int16 = 1
    v2 := []int{127, 32767}
    var v3 uint64 = 0
    fmt.Println(v1, v3)
    var v4 string = "a"
    var v5 uint32 = 65535
    v4 = (v4 + "foo bar")
    var v6 uint16 = 42
    for fi7 := 0; fi7 < 20; fi7++ {
        if ((false && true) || (1.5 <= 3.14159)) {
            switch fi7 {
            case 42:
                for fi8 := 0; fi8 < 14; fi8++ {
                    fmt.Println(v5, v3)
                    _ = fi8
                }
            case -830:
                if ((("hello" + "foo bar") > ("a" + v4)) || false) {
                    for fi9 := 0; fi9 < 2; fi9++ {
                        if (v4 <= "foo bar") {
                            v6 += v6
                            var v10 string = v4
                            _ = len(v2)
                            _ = v10
                        }
                        if true {
                            _ = len(v2)
                        }
                        switch (("hello" + (v4 + "hello")) + v4) {
                        case "a":
                            v6 = (v6 % ((284) | 1))
                        case "":
                            v3 += (v3 & 1)
                            v1 += (v1 - v1)
                            v3 -= v3
                            var v11 uint = 2
                            v12 := GenS1{F0: (^(v11)), F1: v11, F2: 42}
                            _ = v11
                            _ = v12
                        case "hello":
                            if false {
                                v13 := GenS2{F0: 60, F1: -1, F2: 127}
                                _ = v13
                            } else {
                                v6 = 157
                                var v14 uint32 = (v5 | v5)
                                if false {
                                    switch fi7 {
                                    case 0:
                                        v15 := GenS3{F0: true}
                                        _ = v15
                                    case 89:
                                        var v16 uint = 875
                                        var v17 int = fi7
                                        var v18 int16 = v1
                                        _ = v16
                                        _ = v17
                                        _ = v18
                                    default:
                                        if true {
                                            var v19 float64 = 1.5
                                            v20 := GenS4{F0: 0}
                                            _ = v19
                                            _ = v20
                                        } else {
                                            if (744 != 11) {
                                                v6 = 329
                                                fmt.Println(fi9)
                                            }
                                        }
                                    }
                                } else {
                                    if true {
                                        switch fi9 {
                                        case -9223372036854775808:
                                            var v21 uint32 = (v14 / (((v14 & v14)) | 1))
                                            var v22 int8 = 0
                                            _ = v21
                                            _ = v22
                                        case 32767:
                                            v23 := GenS5{F0: true, F1: (v3 - (v3 << ((357) & 63))), F2: (v14 * (^(v14)))}
                                            var v24 float32 = 1e10
                                            _ = v23
                                            _ = v24
                                        case 0:
                                            switch v4 {
                                            case "":
                                                v4 = v4
                                            case "foo bar":
                                                v5 = v5
                                            case "hello":
                                                if (-1 < -1) {
                                                    var v25 float64 = 0.0
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        v14 = uint32(v3)
                                                    default:
                                                        var v26 int32 = 0
                                                        _ = v26
                                                    }
                                                    v1 += (v1 - v1)
                                                    if false {
                                                        var v27 int = 32768
                                                        _ = v27
                                                    }
                                                    _ = v25
                                                } else {
                                                    var v28 int16 = -32768
                                                    var v29 bool = false
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    switch ((v4 + v4) + v4) {
                                                    case "foo bar":
                                                        var v30 string = "a"
                                                        _ = v30
                                                    case "a":
                                                        v5 += 0
                                                    }
                                                    if v29 {
                                                        v5 = v5
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v2 = append(v2, fi9)
                                                    }
                                                    if (!(42 <= uint8(v14))) {
                                                        fmt.Println(fi9, v5)
                                                    } else {
                                                        var v31 int64 = -857
                                                        var v32 uint16 = 0
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                    _ = v28
                                                    _ = v29
                                                }
                                            }
                                        }
                                        v2 = append(v2, int(v3))
                                    }
                                    var v33 uint64 = (v3 / (((v3 & v3)) | 1))
                                    v4 = v4
                                    _ = v33
                                }
                                switch v4 {
                                case "a":
                                    v1 = 1
                                }
                                _ = v14
                            }
                        }
                        _ = fi9
                    }
                    switch v4 {
                    case "a":
                        for fi34 := 0; fi34 < 11; fi34++ {
                            if (80 != 0) {
                                var v35 string = v4
                                _ = v35
                            } else {
                                var v36 uint64 = (^(v3))
                                fmt.Println(fi7, fi7, fi7)
                                _ = v36
                            }
                            var v37 int = fi7
                            _ = fi34
                            _ = v37
                        }
                    default:
                        v38 := GenS6{F0: v3}
                        _ = v38
                    }
                }
            default:
                var v39 int32 = int32(v3)
                _ = v39
            }
            if false {
                var v40 uint64 = v3
                fmt.Println(v5)
                _ = v40
            }
            v2 = append(v2, (fi7 >> ((42) & 63)))
        } else {
            var v41 float64 = (-2.25 + 0.1)
            if (100 < 71) {
                if false {
                    if (fi7 <= 32767) {
                        for fi42 := 0; fi42 < 10; fi42++ {
                            switch (v4 + "") {
                            case "\n\t":
                                fmt.Println(v3)
                            case "foo bar":
                                if (127 < 127) {
                                    var v43 int16 = -165
                                    var v44 bool = ((v1 ^ v1) != (^(v43)))
                                    if len(v2) > 0 {
                                        v2[0] = fi42
                                    }
                                    _ = v43
                                    _ = v44
                                } else {
                                    var v45 uint64 = 691
                                    var v46 bool = (((0.0 - 1e10) / -2.25) == 0.1)
                                    _ = v45
                                    _ = v46
                                }
                            default:
                                var v47 int64 = 42
                                _ = v47
                            }
                            _ = fi42
                        }
                        v1 -= (v1 | v1)
                        if true {
                            v6 = (v6 << ((4294967295) & 63))
                            v48 := GenS3{F0: false}
                            _ = v48
                        } else {
                            for fi49 := 0; fi49 < 2; fi49++ {
                                v3 -= v3
                                var v50 bool = (1.5 > 0.1)
                                _ = fi49
                                _ = v50
                            }
                            if true {
                                if false {
                                    switch ((v4 + "\n\t") + (("hello" + (v4 + "foo bar")) + v4)) {
                                    case "foo bar":
                                        var v51 float64 = v41
                                        _ = v51
                                    case "\n\t":
                                        v52 := GenS1{F0: 102, F1: 1, F2: 151}
                                        _ = v52
                                    case "hello":
                                        if (2147483647 <= 648) {
                                            v5 = 2147483648
                                            if (-1 >= -128) {
                                                var v53 uint8 = 127
                                                _ = v53
                                            }
                                        } else {
                                            v5 += 390
                                            v54 := GenS7{F0: fi7, F1: 0}
                                            _ = v54
                                        }
                                    }
                                }
                                switch v4 {
                                case "\n\t":
                                    for fi55 := 0; fi55 < 3; fi55++ {
                                        v41 -= v41
                                        var v56 float32 = (float32(((1.5 + v41) * (v41 - v41))) + 3.14159)
                                        _ = fi55
                                        _ = v56
                                    }
                                case "a":
                                    for fi57 := 0; fi57 < 14; fi57++ {
                                        if false {
                                            switch v4 {
                                            case "hello":
                                                var v58 float32 = 1e10
                                                _ = v58
                                            }
                                            if ("\n\t" > ("foo bar" + v4)) {
                                                fmt.Println(fi7, v1, v3)
                                                if ((v4 + ((v4 + (v4 + ("a" + v4))) + "foo bar")) != v4) {
                                                    v1 = 42
                                                    v2 = append(v2, fi57)
                                                    var v59 int8 = -128
                                                    _ = v59
                                                }
                                            }
                                            var v60 float64 = v41
                                            var v61 int64 = -1
                                            var v62 string = v4
                                            _ = v60
                                            _ = v61
                                            _ = v62
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = (fi7 + fi7)
                                        }
                                        _ = fi57
                                    }
                                default:
                                    var v63 uint32 = (^(v5))
                                    _ = v63
                                }
                            }
                        }
                    }
                    if ((int64(fi7) >= -9223372036854775808) || (!false)) {
                        var v64 int64 = 0
                        switch fi7 {
                        case 9223372036854775807:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        v5 += 2147483648
                        _ = v64
                    }
                } else {
                    v65 := GenS3{F0: (true || (-9223372036854775808 != -1))}
                    _ = v65
                }
            } else {
                var v66 float32 = 0.1
                switch fi7 {
                case 42:
                    var v67 uint16 = (v6 & v6)
                    _ = v67
                }
                var v68 uint64 = 18446744073709551615
                _ = v66
                _ = v68
            }
            _ = v41
        }
        switch v4 {
        case "foo bar":
            if ((!(1 == 581)) || (0.1 <= 1e10)) {
                var v69 int16 = 910
                switch (v4 + v4) {
                case "\n\t":
                    fmt.Println(v3)
                case "":
                    if true {
                        fmt.Println(v6)
                    }
                case "hello":
                    for fi70 := 0; fi70 < 1; fi70++ {
                        _ = len(v2)
                        if ((v6 >= v6) && (!((1e10 + (1e10 / (((((float64((0.1 + 0.0)) / float64((((1.5 * 0.0) + -2.25) + 0.1))) * (-2.25 / (1.5 * 1e10))) + 1e10) - (1e10 * 1.5)) * (1e10 * (-2.25 * 1e10))))) == 1e10))) {
                            fmt.Println(v69)
                            var v71 uint64 = (^(v3))
                            var v72 int32 = 2147483647
                            v72 -= v72
                            v4 = "hello"
                            var v73 uint32 = v5
                            if (true && true) {
                                if (!((v72 << ((255) & 63)) <= 42)) {
                                    v2 = append(v2, fi7)
                                    var v74 uint = 110
                                    v71 = (v71 % (((v3 / (((v71 % ((v71) | 1))) | 1))) | 1))
                                    v74 = 409
                                    _ = v74
                                }
                            }
                            var v75 uint32 = 2147483648
                            _ = len(v2)
                            _ = v71
                            _ = v72
                            _ = v73
                            _ = v75
                        } else {
                            v4 = v4
                            v76 := GenS1{F0: uint(v5), F1: 862, F2: 128}
                            _ = v76
                        }
                        _ = fi70
                    }
                }
                _ = v69
            } else {
                v77 := GenS7{F0: fi7, F1: 1}
                v78 := GenS5{F0: false, F1: v3, F2: 511}
                _ = len(v2)
                _ = v77
                _ = v78
            }
        case "\n\t":
            v2 = append(v2, fi7)
        }
        fmt.Println(v3)
        _ = fi7
    }
    var v79 int32 = 42
    for fi80 := 0; fi80 < 17; fi80++ {
        var v81 string = (((v4 + v4) + "a") + v4)
        if false {
            _ = len(v2)
            for fi82 := 0; fi82 < 19; fi82++ {
                var v83 int32 = 1
                v81 = v4
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi82
                _ = v83
            }
        } else {
            v84 := GenS3{F0: (true || false)}
            var v85 int16 = v1
            v6 = (v6 / ((v6) | 1))
            _ = v84
            _ = v85
        }
        _ = fi80
        _ = v81
    }
    var v86 int = 42
    if true {
        v87 := GenS4{F0: 127}
        fmt.Println(v5, v86)
        _ = v87
    } else {
        var v88 bool = (v6 != 736)
        if v88 {
            var v89 int32 = (v79 * 0)
            v79 -= 42
            _ = v89
        } else {
            if ((v3 | v3) < v3) {
                v3 += 2
                var v90 int32 = (v79 / (((v79 - (v79 ^ 0))) | 1))
                _ = v90
            } else {
                switch ("a" + "\n\t") {
                case "hello":
                    var v91 float32 = 0.1
                    _ = v91
                case "foo bar":
                    switch (v4 + "foo bar") {
                    case "hello":
                        if v88 {
                            v92 := GenS6{F0: 4294967295}
                            _ = v92
                        } else {
                            var v93 uint64 = (v3 | (v3 - v3))
                            if (0.0 != (0.0 - ((((0.1 / float64(((((0.0 - 0.0) - 1e10) * (1e10 + 1e10)) + (3.14159 + (float32(3.14159) / 1.0))))) * 1.5) / (0.1 / 1.5)) * ((((1e10 / 1.5) / 1.0) / -2.25) - ((1e10 * 3.14159) * (1e10 * -2.25)))))) {
                                switch (v86 << ((614) & 63)) {
                                case 2147483647:
                                    v94 := GenS1{F0: 4294967295, F1: 255, F2: 158}
                                    _ = v94
                                case 4294967295:
                                    if (v4 > ((v4 + v4) + v4)) {
                                        if (v1 != -427) {
                                            v95 := GenS8{F0: v93, F1: false, F2: 419}
                                            if v88 {
                                                if v88 {
                                                    v96 := GenS9{F0: v86, F1: (-(v1)), F2: (v1 / ((v1) | 1))}
                                                    v97 := GenS9{F0: (v86 >> ((160) & 63)), F1: int16(v3), F2: v1}
                                                    _ = v96
                                                    _ = v97
                                                }
                                                switch (v86 | (v86 ^ (v86 + 0))) {
                                                case 1:
                                                    var v98 int = 4294967295
                                                    _ = v98
                                                default:
                                                    v99 := GenS10{F0: v93, F1: v4, F2: 127}
                                                    _ = v99
                                                }
                                                v100 := GenS11{F0: (v6 + v6), F1: (v1 >> ((334) & 63))}
                                                _ = v100
                                            }
                                            v86 = (v86 >> ((65535) & 63))
                                            _ = v95
                                        } else {
                                            v79 += v79
                                            v101 := GenS12{F0: 18446744073709551615, F1: 2147483648}
                                            _ = v101
                                        }
                                    } else {
                                        fmt.Println(v1, v1)
                                        var v102 uint32 = (^(v5))
                                        _ = v102
                                    }
                                }
                                v103 := GenS13{F0: int64(v79), F1: 112, F2: v86}
                                v79 -= -35
                                _ = v103
                            } else {
                                switch (v4 + v4) {
                                case "":
                                    v6 = v6
                                default:
                                    if v88 {
                                        fmt.Println(v5, v5)
                                        v93 = (v3 << ((18446744073709551615) & 63))
                                    }
                                }
                                _ = len(v2)
                                var v104 string = "foo bar"
                                _ = v104
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v93
                        }
                    case "foo bar":
                        if v88 {
                            fmt.Println(v1, v3)
                            if false {
                                v2 = append(v2, -1)
                                v86 = int(v79)
                                if (!true) {
                                    for fi105 := 0; fi105 < 3; fi105++ {
                                        var v106 int = -381
                                        _ = fi105
                                        _ = v106
                                    }
                                } else {
                                    if v88 {
                                        for fi107 := 0; fi107 < 6; fi107++ {
                                            if len(v2) > 0 {
                                                v2[0] = (v86 % ((-98) | 1))
                                            }
                                            _ = fi107
                                        }
                                    }
                                }
                            } else {
                                if (v4 < (v4 + "hello")) {
                                    v2 = append(v2, (v86 / (((v86 * (v86 % ((v86) | 1)))) | 1)))
                                    v108 := GenS4{F0: 42}
                                    if ((v88 && false) || true) {
                                        switch int(v3) {
                                        case 949:
                                            if (int64(v5) < 316) {
                                                for fi109 := 0; fi109 < 11; fi109++ {
                                                    v4 = (v4 + ("hello" + (v4 + (((v4 + (("\n\t" + (("hello" + (v4 + "hello")) + v4)) + v4)) + ("a" + (v4 + (v4 + (v4 + "\n\t"))))) + v4))))
                                                    v86 = (v86 * fi109)
                                                    _ = fi109
                                                }
                                            } else {
                                                fmt.Println(v4, v86, v79)
                                            }
                                        case 127:
                                            var v110 bool = (9223372036854775807 != 9223372036854775807)
                                            _ = v110
                                        }
                                    }
                                    _ = v108
                                }
                                if (2 > 4294967295) {
                                    var v111 string = "hello"
                                    var v112 int8 = -122
                                    _ = v111
                                    _ = v112
                                } else {
                                    var v113 int64 = -408
                                    var v114 float64 = 1e10
                                    _ = v113
                                    _ = v114
                                }
                                v79 += (v79 | v79)
                            }
                            if true {
                                var v115 float64 = float64((((1.5 - (0.1 / 1.5)) * -2.25) - (float32(1.5) - (1e10 * 3.14159))))
                                var v116 uint32 = 65535
                                v3 += v3
                                _ = v115
                                _ = v116
                            }
                        } else {
                            for fi117 := 0; fi117 < 3; fi117++ {
                                var v118 float64 = 0.0
                                _ = fi117
                                _ = v118
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    }
                case "\n\t":
                    v86 -= (v86 + v86)
                default:
                    if ((v3 << ((322) & 63)) >= v3) {
                        var v119 string = "hello"
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v119
                    } else {
                        v120 := GenS11{F0: v6, F1: (v1 + (v1 * -32768))}
                        if v88 {
                            if len(v2) > 0 {
                                v2[0] = v86
                            }
                        } else {
                            if (v5 >= (v5 & v5)) {
                                if (false && false) {
                                    if true {
                                        v1 = (v1 ^ -32768)
                                        v5 -= 857
                                    } else {
                                        fmt.Println(v3, v1)
                                        v79 = -1
                                    }
                                    _ = len(v2)
                                } else {
                                    for fi121 := 0; fi121 < 3; fi121++ {
                                        var v122 uint8 = uint8(v5)
                                        _ = fi121
                                        _ = v122
                                    }
                                    var v123 uint = 419
                                    if ((v3 * 15) > v3) {
                                        var v124 int = (-(v86))
                                        if (v88 && (-13 <= 127)) {
                                            fmt.Println(v124, v5)
                                            v6 = 65535
                                        } else {
                                            v125 := GenS13{F0: 42, F1: -128, F2: v86}
                                            if len(v2) > 0 {
                                                v2[0] = (v86 + int(v79))
                                            }
                                            v5 += (v5 | (v5 + v5))
                                            _ = v125
                                        }
                                        for fi126 := 0; fi126 < 3; fi126++ {
                                            var v127 uint8 = 250
                                            for fi128 := 0; fi128 < 2; fi128++ {
                                                switch ("a" + (v4 + ("hello" + "\n\t"))) {
                                                case "hello":
                                                    switch v4 {
                                                    case "a":
                                                        var v129 bool = v88
                                                        _ = v129
                                                    }
                                                case "a":
                                                    if (-128 < 52) {
                                                        switch (v86 % (((fi126 ^ 32768)) | 1)) {
                                                        case 4294967295:
                                                            switch v4 {
                                                            case "foo bar":
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            }
                                                        case 32767:
                                                            v123 += v123
                                                        }
                                                    } else {
                                                        if (!false) {
                                                            var v130 float64 = 0.1
                                                            _ = v130
                                                        } else {
                                                            if v88 {
                                                                v6 = 2
                                                                if v88 {
                                                                    v131 := GenS2{F0: 1, F1: (v79 << (((v123 ^ 121)) & 63)), F2: (^(v127))}
                                                                    if (1 <= 42) {
                                                                        var v132 int64 = 0
                                                                        fmt.Println(fi126)
                                                                        _ = v132
                                                                    } else {
                                                                        var v133 float32 = (-2.25 + (0.0 * 1.5))
                                                                        var v134 int16 = v1
                                                                        var v135 bool = (0.0 != (3.14159 + v133))
                                                                        _ = v133
                                                                        _ = v134
                                                                        _ = v135
                                                                    }
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        v136 := GenS13{F0: 632, F1: -1, F2: (fi126 & (v124 * (v124 >> ((65535) & 63))))}
                                                                        _ = v136
                                                                    case "foo bar":
                                                                        v137 := GenS10{F0: v3, F1: "hello", F2: 1}
                                                                        _ = v137
                                                                    case "":
                                                                        var v138 string = "a"
                                                                        _ = v138
                                                                    }
                                                                    _ = v131
                                                                } else {
                                                                    if v88 {
                                                                        v139 := GenS14{F0: 9223372036854775807, F1: (v123 >= (v123 ^ v123)), F2: 187}
                                                                        v140 := GenS9{F0: (-(v86)), F1: v1, F2: -1}
                                                                        _ = v139
                                                                        _ = v140
                                                                    } else {
                                                                        v141 := GenS15{F0: (v127 >> ((4294967295) & 63))}
                                                                        fmt.Println(v6)
                                                                        v6 += (v6 | v6)
                                                                        _ = v141
                                                                    }
                                                                }
                                                            } else {
                                                                fmt.Println(v127, v123, v6)
                                                            }
                                                            v127 -= (v127 / (((v127 >> (((v123 << (((v123 + (v123 / ((uint(v5)) | 1)))) & 63))) & 63))) | 1))
                                                        }
                                                        v142 := GenS16{F0: v5, F1: (v6 * v6)}
                                                        if (!false) {
                                                            v127 += v127
                                                            fmt.Println(v5, fi128, fi126)
                                                        }
                                                        _ = v142
                                                    }
                                                case "":
                                                    var v143 string = ("a" + v4)
                                                    if v88 {
                                                        var v144 int32 = (v79 ^ (v79 % ((v79) | 1)))
                                                        _ = v144
                                                    } else {
                                                        v145 := GenS17{F0: uint64(v127), F1: (^(v5))}
                                                        _ = v145
                                                    }
                                                    v1 -= v1
                                                    _ = v143
                                                default:
                                                    v5 = 2147483647
                                                }
                                                fmt.Println(v127, v4)
                                                v6 -= 255
                                                _ = fi128
                                            }
                                            switch v86 {
                                            case 42:
                                                fmt.Println(v4, fi126, v6)
                                            case 32767:
                                                fmt.Println(v123)
                                            default:
                                                if (v88 || v88) {
                                                    var v146 int8 = 100
                                                    switch v124 {
                                                    case 191:
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    default:
                                                        v147 := GenS15{F0: 1}
                                                        _ = v147
                                                    }
                                                    _ = v146
                                                }
                                            }
                                            _ = fi126
                                            _ = v127
                                        }
                                        _ = v124
                                    }
                                    _ = v123
                                }
                            }
                            var v148 bool = (v4 > (v4 + v4))
                            _ = v148
                        }
                        _ = v120
                    }
                }
                var v149 uint64 = uint64(v5)
                _ = v149
            }
            v5 = v5
            if v88 {
                if v88 {
                    v4 = ((v4 + v4) + v4)
                    var v150 uint = 255
                    v5 = v5
                    _ = v150
                }
            } else {
                v86 -= (-(v86))
                if ((v3 | (v3 / ((472) | 1))) < (v3 & 4294967295)) {
                    v5 = (v5 << ((18446744073709551615) & 63))
                }
                switch -362 {
                case 128:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
        }
        _ = v88
    }
    v5 = v5
    for fi151 := 0; fi151 < 4; fi151++ {
        if true {
            switch "a" {
            case "\n\t":
                fmt.Println(v6, v6, fi151)
            }
            var v152 int8 = -128
            v153 := GenS3{F0: (true || (42 <= 1))}
            _ = v152
            _ = v153
        }
        var v154 int64 = 42
        _ = fi151
        _ = v154
    }
    switch v86 {
    case 9223372036854775807:
        fmt.Println(v3)
    case 0:
        if false {
            v1 = -1
            var v155 int8 = -1
            var v156 int8 = (v155 & 42)
            _ = v155
            _ = v156
        } else {
            v157 := GenS11{F0: v6, F1: -248}
            v158 := GenS18{F0: v1, F1: ((("a" + "\n\t") + v4) + (v4 + "foo bar"))}
            _ = v157
            _ = v158
        }
    case -9223372036854775808:
        for fi159 := 0; fi159 < 19; fi159++ {
            v160 := GenS19{F0: 97, F1: (v1 & (v1 >> ((1) & 63))), F2: 0}
            if false {
                for fi161 := 0; fi161 < 20; fi161++ {
                    if (0.0 < ((1e10 + float32(-2.25)) - -2.25)) {
                        switch "" {
                        case "foo bar":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case "\n\t":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case "a":
                            switch "foo bar" {
                            case "":
                                fmt.Println(fi159, fi159, fi161)
                            case "\n\t":
                                switch (^(fi159)) {
                                case -420:
                                    var v162 uint32 = v5
                                    _ = v162
                                case -795:
                                    if (float32((0.0 + 0.0)) == (1.5 * (0.0 * float32(((-2.25 / 1.0) * 0.1))))) {
                                        var v163 float64 = (-2.25 - 1e10)
                                        switch v4 {
                                        case "hello":
                                            v3 = (v3 & 4294967295)
                                        case "\n\t":
                                            var v164 uint16 = 42
                                            var v165 bool = false
                                            if ((v165 && true) || true) {
                                                var v166 int8 = 42
                                                _ = v166
                                            } else {
                                                var v167 int = -20
                                                v165 = v165
                                                var v168 uint64 = (v3 / ((1) | 1))
                                                _ = v167
                                                _ = v168
                                            }
                                            fmt.Println(v5, v5)
                                            var v169 int16 = v1
                                            _ = v164
                                            _ = v165
                                            _ = v169
                                        case "a":
                                            switch ((v4 + v4) + "") {
                                            case "foo bar":
                                                var v170 uint = 2
                                                _ = v170
                                            }
                                        }
                                        _ = v163
                                    }
                                }
                            default:
                                fmt.Println(v1, v6, v6)
                            }
                        default:
                            switch "a" {
                            case "\n\t":
                                var v171 int = int(v1)
                                _ = v171
                            case "hello":
                                v172 := GenS20{F0: (v79 >> ((255) & 63)), F1: 65535, F2: (v5 << ((616) & 63))}
                                _ = v172
                            case "":
                                var v173 float64 = 1.5
                                _ = v173
                            default:
                                var v174 int16 = -172
                                _ = v174
                            }
                        }
                    } else {
                        var v175 uint16 = v6
                        _ = v175
                    }
                    _ = len(v2)
                    var v176 int8 = -1
                    _ = fi161
                    _ = v176
                }
                fmt.Println(fi159, v1)
                for fi177 := 0; fi177 < 3; fi177++ {
                    var v178 float32 = ((0.0 - 1e10) - (3.14159 + 1.5))
                    _ = len(v2)
                    _ = fi177
                    _ = v178
                }
            } else {
                var v179 int = (fi159 ^ v86)
                v5 = (^(v5))
                v2 = append(v2, int(v1))
                _ = v179
            }
            switch -176 {
            case -9223372036854775808:
                if (false && ((v79 << ((uint(v5)) & 63)) < (v79 + v79))) {
                    v180 := GenS4{F0: 1}
                    var v181 float32 = 0.0
                    _ = v180
                    _ = v181
                }
            case -2147483648:
                fmt.Println(v86, v1, v6)
            }
            _ = fi159
            _ = v160
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v79:", v79)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v79
    _ = v86
}
