package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 int; F1 uint8; F2 uint }
type GenS5 struct { F0 uint64; F1 string; F2 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int; F1 uint64 }
type GenS10 struct { F0 uint; F1 bool }
type GenS11 struct { F0 uint16; F1 uint16 }
type GenS12 struct { F0 int8; F1 uint8; F2 int }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int; F1 int32; F2 uint16 }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 uint32; F1 string; F2 int16 }
type GenS18 struct { F0 uint16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    var v4 uint32 = 392
    var v5 int = 538
    var v6 uint32 = v4
    if true {
        v6 = (v4 >> ((42) & 63))
        v6 = (v4 % (((v4 - (^(v6)))) | 1))
    } else {
        delete(v3, v5)
    }
    var v7 string = ((("foo bar" + "a") + "") + "foo bar")
    v8 := []int{}
    for fi9 := 0; fi9 < 18; fi9++ {
        fmt.Println(v6, v4)
        v10 := v3[fi9]
        fmt.Println(v4, fi9, v7)
        _ = fi9
        _ = v10
    }
    v8 = append(v8, (v5 ^ (v5 * v5)))
    for fi11 := 0; fi11 < 18; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            v6 -= 972
            _ = fi12
        }
        if (127 > 42) {
            var v13 int64 = -1
            var v14 float32 = (3.14159 + (1.5 - 1.5))
            _ = v13
            _ = v14
        } else {
            v15 := GenS1{F0: 119, F1: 0}
            var v16 float64 = -2.25
            switch v5 {
            case -9223372036854775808:
                if (((2 < 2147483648) || true) && (false || false)) {
                    var v17 int16 = -32768
                    _ = v17
                }
            case 246:
                fmt.Println(fi11)
            default:
                v18 := GenS1{F0: 1, F1: 65}
                _ = v18
            }
            _ = v15
            _ = v16
        }
        _ = fi11
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        switch v5 {
        case 127:
            if (-1 != 42) {
                v4 = (v6 / (((v4 - (v6 & (v6 * v6)))) | 1))
                v4 = v4
            }
        case 598:
            v5 = 127
        case 9223372036854775807:
            v4 -= 0
        }
        if true {
            switch fi19 {
            case 127:
                switch ((v7 + v7) + v7) {
                case "hello":
                    v20 := GenS1{F0: 32767, F1: 1}
                    _ = v20
                case "foo bar":
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                }
            case 385:
                var v21 int32 = -1
                _ = v21
            default:
                if (-2.25 != (0.0 * float32(0.1))) {
                    for fi22 := 0; fi22 < 20; fi22++ {
                        v4 -= 2
                        if ((v5 > (fi19 & fi19)) && false) {
                            switch v7 {
                            case "hello":
                                fmt.Println(fi19, fi22)
                            case "":
                                if false {
                                    fmt.Println(fi19, fi22, fi19)
                                }
                            case "a":
                                var v23 uint64 = 4294967295
                                _ = v23
                            }
                            v8 = append(v8, 0)
                        }
                        _ = fi22
                    }
                    for fi24 := 0; fi24 < 10; fi24++ {
                        v4 = 856
                        v4 = 4294967295
                        _ = fi24
                    }
                } else {
                    v25 := GenS1{F0: -706, F1: 127}
                    if len(v8) > 0 {
                        v8[0] = (fi19 | (fi19 >> ((857) & 63)))
                    }
                    _ = v25
                }
            }
            var v26 int16 = 0
            _ = v26
        } else {
            if (true && (!false)) {
                if len(v8) > 0 {
                    v8[0] = fi19
                }
                v3[(^(v5))] = true
                v7 = v7
            }
            v6 -= v6
            v4 = v4
        }
        _ = fi19
    }
    switch v7 {
    case "\n\t":
        if (825 != 408) {
            v5 = v5
            if (!(!(((v7 + v7) + v7) > (("hello" + v7) + "hello")))) {
                if (18446744073709551615 < 1) {
                    var v27 uint64 = 31
                    v4 = v4
                    var v28 int32 = 318
                    _ = v27
                    _ = v28
                }
            }
        }
    case "hello":
        var v29 int32 = 1
        _ = v29
    case "foo bar":
        for fi30 := 0; fi30 < 6; fi30++ {
            fmt.Println(v6, v7, fi30)
            _ = fi30
        }
    }
    v7 = ((("" + v7) + v7) + "\n\t")
    for fi31 := 0; fi31 < 17; fi31++ {
        for fi32 := 0; fi32 < 10; fi32++ {
            v5 -= (v5 - fi32)
            var v33 int64 = 1
            _ = fi32
            _ = v33
        }
        for fi34 := 0; fi34 < 17; fi34++ {
            switch v7 {
            case "":
                v4 -= 65535
            case "\n\t":
                fmt.Println(v6)
            case "foo bar":
                var v35 uint32 = v4
                if (!(false && (42 < 612))) {
                    var v36 string = ("" + v7)
                    var v37 uint16 = 42
                    fmt.Println(v5)
                    _ = v36
                    _ = v37
                } else {
                    var v38 int8 = -96
                    switch (v5 / ((fi34) | 1)) {
                    case 4294967295:
                        fmt.Println(v5, v35, fi34)
                    case -1:
                        fmt.Println(v6, v5)
                    case 32768:
                        v5 -= v5
                    default:
                        switch (^(fi34)) {
                        case -2147483648:
                            if (!(42 == 181)) {
                                _ = len(v8)
                                v39 := GenS1{F0: -186, F1: v38}
                                fmt.Println(v5, fi31, fi31)
                                v40 := GenS1{F0: 1, F1: v38}
                                _ = v39
                                _ = v40
                            }
                        case 42:
                            if true {
                                v41 := GenS2{F0: -15}
                                fmt.Println(v5, fi31)
                                _ = v41
                            }
                        case 2147483647:
                            if false {
                                if ((0.1 + 1.5) < (3.14159 / 3.14159)) {
                                    var v42 uint = 42
                                    delete(v1, (fi31 >> ((v42) & 63)))
                                    _ = v42
                                }
                                v35 = (v6 + v35)
                                v43 := GenS2{F0: (fi34 * (fi31 / ((fi34) | 1)))}
                                _ = v43
                            }
                            var v44 int16 = 0
                            v45 := v3[32767]
                            _ = v44
                            _ = v45
                        }
                    }
                    var v46 uint32 = 326
                    _ = v38
                    _ = v46
                }
                _ = v35
            default:
                switch 9223372036854775807 {
                case 32768:
                    var v47 uint = 138
                    _ = v47
                case 992:
                    fmt.Println(v5)
                case 9223372036854775807:
                    var v48 float32 = (((float32(3.14159) * 1e10) * float32((1e10 - 1.5))) + 1.5)
                    _ = v48
                default:
                    if true {
                        if (42 != -1) {
                            fmt.Println(v4, v5)
                            v5 = -1
                        }
                        v1[(^(fi31))] = v4
                        var v49 uint16 = 32767
                        var v50 float32 = (1.5 / 1.0)
                        switch (^(v5)) {
                        case 655:
                            switch -534 {
                            case 815:
                                var v51 uint16 = (v49 << ((18446744073709551615) & 63))
                                _ = v51
                            }
                            if true {
                                v52 := GenS3{F0: fi34, F1: 970}
                                _ = v52
                            } else {
                                switch v7 {
                                case "":
                                    if true {
                                        v49 += v49
                                        if (0.1 != -2.25) {
                                            v53 := GenS4{F0: 0, F1: 227, F2: 585}
                                            _ = v53
                                        } else {
                                            if false {
                                                switch (v7 + v7) {
                                                case "foo bar":
                                                    if false {
                                                        var v54 bool = true
                                                        v55 := GenS2{F0: (fi34 + -121)}
                                                        _ = v54
                                                        _ = v55
                                                    } else {
                                                        fmt.Println(v7, v4)
                                                        v5 = 259
                                                        var v56 uint64 = 0
                                                        _ = v56
                                                    }
                                                case "a":
                                                    var v57 uint8 = 255
                                                    _ = v57
                                                case "\n\t":
                                                    v49 -= (v49 << ((33) & 63))
                                                default:
                                                    v49 -= 255
                                                }
                                                switch "a" {
                                                case "foo bar":
                                                    v6 += v6
                                                    var v58 int8 = -29
                                                    _ = v58
                                                case "":
                                                    var v59 int16 = int16(fi34)
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                    v4 -= v6
                                                    _ = v59
                                                case "a":
                                                    v60 := GenS1{F0: -1, F1: 114}
                                                    _ = v60
                                                }
                                            } else {
                                                v5 -= fi31
                                                if false {
                                                    switch -9223372036854775808 {
                                                    case 4294967295:
                                                        var v61 string = v7
                                                        _ = v61
                                                    case 127:
                                                        v62 := GenS4{F0: (fi31 * (^(fi34))), F1: 255, F2: uint(v4)}
                                                        _ = v62
                                                    default:
                                                        var v63 uint8 = 227
                                                        _ = v63
                                                    }
                                                }
                                            }
                                        }
                                        var v64 uint8 = 2
                                        _ = v64
                                    } else {
                                        v7 = ("hello" + (("hello" + "") + v7))
                                        var v65 uint64 = 18446744073709551615
                                        fmt.Println(fi34, v49, v5)
                                        v66 := GenS5{F0: v65, F1: "foo bar", F2: 1000}
                                        var v67 uint32 = v4
                                        var v68 int8 = 0
                                        v6 += v4
                                        v4 = (v6 / (((v6 - 30)) | 1))
                                        _ = v65
                                        _ = v66
                                        _ = v67
                                        _ = v68
                                    }
                                    if false {
                                        var v69 uint = 413
                                        _ = v69
                                    } else {
                                        var v70 uint8 = 2
                                        v6 += (^(v4))
                                        v4 -= v6
                                        switch (fi34 * fi31) {
                                        case 32768:
                                            v71 := GenS6{F0: 2}
                                            _ = v71
                                        default:
                                            fmt.Println(v5, v7, v6)
                                        }
                                        switch fi31 {
                                        case 32768:
                                            fmt.Println(v50, fi34)
                                        case 127:
                                            if (v4 != v6) {
                                                v6 = (^(v4))
                                                switch fi31 {
                                                case 917:
                                                    if false {
                                                        _ = len(v3)
                                                        v72 := GenS7{F0: 1}
                                                        _ = v72
                                                    } else {
                                                        switch v7 {
                                                        case "":
                                                            v6 -= (v4 << ((18446744073709551615) & 63))
                                                        case "hello":
                                                            v1[(fi34 - fi31)] = (^(v6))
                                                        case "a":
                                                            if (-1 > (v5 % ((fi31) | 1))) {
                                                                var v73 string = v7
                                                                var v74 bool = true
                                                                v6 -= uint32(v70)
                                                                _ = v73
                                                                _ = v74
                                                            }
                                                        }
                                                    }
                                                }
                                                fmt.Println(fi31, v50, fi34)
                                            }
                                        case -9223372036854775808:
                                            fmt.Println(v6, v5)
                                        default:
                                            var v75 int8 = -17
                                            _ = v75
                                        }
                                        _ = v70
                                    }
                                    v8 = append(v8, fi34)
                                case "hello":
                                    var v76 int8 = -1
                                    _ = v76
                                default:
                                    fmt.Println(fi31)
                                }
                                v77 := GenS8{F0: -32768}
                                switch v7 {
                                case "":
                                    _ = len(v2)
                                case "hello":
                                    fmt.Println(v50, v4)
                                case "a":
                                    v78 := GenS9{F0: -2147483648, F1: 2}
                                    v49 += 42
                                    _ = v78
                                }
                                _ = v77
                            }
                        }
                        fmt.Println(v49)
                        _ = v49
                        _ = v50
                    } else {
                        var v79 float32 = ((float32(float64(0.1)) + 1e10) - 0.1)
                        if false {
                            v79 = v79
                            if (v79 != 3.14159) {
                                var v80 string = "hello"
                                _ = v80
                            } else {
                                v81 := GenS5{F0: 584, F1: (v7 + v7), F2: -370}
                                switch "" {
                                case "\n\t":
                                    if (128 >= 242) {
                                        v4 += (v6 & (^(v6)))
                                    }
                                    var v82 uint32 = v4
                                    _ = v82
                                case "foo bar":
                                    if false {
                                        v7 = v7
                                        switch (((("" + "") + (v7 + "")) + ("\n\t" + "foo bar")) + v7) {
                                        case "hello":
                                            switch v7 {
                                            case "foo bar":
                                                if ((!true) || (v79 >= 0.0)) {
                                                    v83 := GenS10{F0: 636, F1: true}
                                                    switch (v7 + "hello") {
                                                    case "":
                                                        var v84 int16 = -32768
                                                        _ = v84
                                                    case "hello":
                                                        v6 = (v4 * 42)
                                                    case "\n\t":
                                                        v85 := GenS11{F0: 42, F1: 1}
                                                        _ = v85
                                                    default:
                                                        switch fi34 {
                                                        case 9223372036854775807:
                                                            if ((-2.25 >= ((v79 - v79) + ((1e10 / 3.14159) - v79))) || false) {
                                                                var v86 uint8 = 128
                                                                v87 := GenS12{F0: 1, F1: 128, F2: (v5 & (fi34 - v5))}
                                                                _ = v86
                                                                _ = v87
                                                            }
                                                        case 501:
                                                            v8 = append(v8, v5)
                                                        }
                                                        v5 -= v5
                                                        var v88 uint = 1
                                                        v88 += 65535
                                                        _ = v88
                                                    }
                                                    switch "" {
                                                    case "":
                                                        fmt.Println(v7, fi31, v79)
                                                    }
                                                    fmt.Println(fi31, v5, v5)
                                                    v89 := GenS13{F0: fi34}
                                                    _ = v83
                                                    _ = v89
                                                }
                                            case "a":
                                                v90 := GenS14{F0: -915}
                                                _ = v90
                                            case "\n\t":
                                                v6 += 623
                                            }
                                        case "\n\t":
                                            var v91 int32 = -2147483648
                                            _ = v91
                                        case "":
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                        }
                                        fmt.Println(v6, fi31, v79)
                                    } else {
                                        v92 := GenS15{F0: (v5 - -860), F1: -105, F2: 255}
                                        _ = v92
                                    }
                                default:
                                    var v93 uint8 = 2
                                    _ = v93
                                }
                                v5 += (v5 % ((fi34) | 1))
                                _ = v81
                            }
                            if (100 >= -120) {
                                var v94 uint16 = 32767
                                if (18446744073709551615 > 1) {
                                    v6 += 1
                                    if true {
                                        if true {
                                            var v95 int64 = 793
                                            var v96 uint64 = 1
                                            v97 := GenS16{F0: 238}
                                            _ = v95
                                            _ = v96
                                            _ = v97
                                        }
                                        var v98 int8 = 111
                                        delete(v1, v5)
                                        var v99 float64 = -2.25
                                        _ = v98
                                        _ = v99
                                    }
                                    v5 += 0
                                    var v100 float32 = 0.1
                                    v101 := v2[(("hello" + "hello") + v7)]
                                    _ = v100
                                    _ = v101
                                } else {
                                    var v102 float32 = -2.25
                                    _ = v102
                                }
                                if true {
                                    v7 = v7
                                }
                                _ = v94
                            }
                            var v103 uint32 = 699
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            if (true && (0 > 65535)) {
                                if (true && (1 > 885)) {
                                    if (!true) {
                                        var v104 uint = 255
                                        _ = v104
                                    }
                                    fmt.Println(v5)
                                } else {
                                    v105 := GenS17{F0: 439, F1: v7, F2: 0}
                                    v5 = 0
                                    _ = v105
                                }
                                v3[1] = false
                            } else {
                                var v106 bool = ((0 > 1) && false)
                                v106 = (-1 != -199)
                                if (-636 <= fi34) {
                                    if (1.5 < v79) {
                                        v107 := GenS18{F0: 32767}
                                        switch (v5 << ((2) & 63)) {
                                        case 1:
                                            v5 = (v5 % ((21) | 1))
                                        case 9223372036854775807:
                                            fmt.Println(v106, v106)
                                        }
                                        fmt.Println(v5, v5, v79)
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                        _ = v107
                                    }
                                } else {
                                    var v108 uint = 255
                                    _ = v108
                                }
                                _ = v106
                            }
                            fmt.Println(fi34, v6, fi34)
                            var v109 uint16 = 65535
                            v6 = (v4 + v103)
                            _ = v103
                            _ = v109
                        } else {
                            delete(v3, v5)
                            fmt.Println(fi34)
                            v110 := GenS7{F0: 883}
                            _ = v110
                        }
                        _ = len(v8)
                        _ = v79
                    }
                }
            }
            v8 = append(v8, (^(v5)))
            fmt.Println(v7, fi31)
            if false {
                var v111 int = fi34
                var v112 int = v5
                v6 -= v4
                _ = v111
                _ = v112
            }
            if false {
                var v113 int16 = 1000
                v113 -= (v113 * v113)
                _ = v113
            }
            _ = fi34
        }
        fmt.Println(v5, v5, v4)
        _ = fi31
    }
    v5 = 127
    for fi114 := 0; fi114 < 15; fi114++ {
        v5 = (-(v5))
        v4 -= v6
        _ = fi114
    }
    v1[(v5 - 128)] = v4
    var v115 int32 = 171
    var v116 int = v5
    if false {
        fmt.Println(v7)
    } else {
        v4 += v4
    }
    for fi117 := 0; fi117 < 1; fi117++ {
        fmt.Println(v6, v116, v4)
        for fi118 := 0; fi118 < 9; fi118++ {
            var v119 int32 = -1
            fmt.Println(v7, fi118)
            _ = len(v1)
            _ = fi118
            _ = v119
        }
        if (true || false) {
            v5 = fi117
            switch v116 {
            case -771:
                v6 = (v4 + v6)
            }
        }
        _ = fi117
    }
    v1[(v5 * (v116 / (((v5 >> ((1) & 63))) | 1)))] = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v115:", v115)
    fmt.Println("v116:", v116)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v115
    _ = v116
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
