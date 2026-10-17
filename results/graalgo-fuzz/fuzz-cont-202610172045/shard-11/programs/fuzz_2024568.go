package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 uint; F2 int8 }
type GenS5 struct { F0 int16; F1 uint16; F2 int32 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint32; F1 int8 }
type GenS9 struct { F0 int16; F1 uint32 }
type GenS10 struct { F0 int16; F1 uint32; F2 int }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int64; F1 string; F2 uint64 }
type GenS14 struct { F0 string; F1 int16; F2 int64 }
type GenS15 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 int64; F1 uint; F2 uint16 }
type GenS18 struct { F0 uint16; F1 uint16 }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 string; F1 int64 }
type GenS21 struct { F0 uint8; F1 bool }
type GenS22 struct { F0 int }
type GenS23 struct { F0 int8; F1 int16 }
type GenS24 struct { F0 uint32 }
func genFn0(p0 int16, p1 int, p2 uint64) uint {
    if false {
        v4 := []int{}
        v4 = append(v4, (p1 * (p1 << ((246) & 63))))
        switch -919 {
        case -2147483648:
            if (true || false) {
                if (-9223372036854775808 != 1) {
                    var v5 float64 = ((3.14159 + float64(0.1)) / 0.1)
                    var v6 string = "\n\t"
                    _ = v5
                    _ = v6
                } else {
                    v7 := GenS1{F0: p2}
                    v8 := GenS1{F0: p2}
                    _ = v7
                    _ = v8
                }
                for fi9 := 0; fi9 < 12; fi9++ {
                    for fi10 := 0; fi10 < 13; fi10++ {
                        switch "foo bar" {
                        case "hello":
                            if false {
                                if true {
                                    v11 := GenS2{F0: (p2 << ((1) & 63)), F1: 255}
                                    switch ("a" + "a") {
                                    case "foo bar":
                                        var v12 uint32 = 0
                                        _ = v12
                                    case "\n\t":
                                        p1 += fi9
                                    case "":
                                        var v13 float64 = 0.1
                                        _ = v13
                                    default:
                                        v14 := GenS1{F0: p2}
                                        _ = v14
                                    }
                                    _ = v11
                                } else {
                                    var v15 int16 = (^(p0))
                                    _ = v15
                                }
                                _ = len(v4)
                            }
                        default:
                            var v16 uint = 106
                            _ = v16
                        }
                        _ = fi10
                    }
                    v17 := GenS3{F0: -1}
                    for fi18 := 0; fi18 < 14; fi18++ {
                        if false {
                            var v19 uint8 = 19
                            if len(v4) > 0 {
                                v4[0] = 0
                            }
                            var v20 int = (fi18 | (fi18 << ((uint(p2)) & 63)))
                            var v21 uint = 255
                            _ = v19
                            _ = v20
                            _ = v21
                        } else {
                            p0 = 32767
                            if false {
                                v22 := GenS3{F0: 1}
                                _ = v22
                            }
                            switch "\n\t" {
                            case "a":
                                switch "\n\t" {
                                case "foo bar":
                                    if len(v4) > 0 {
                                        v4[0] = (fi9 + fi9)
                                    }
                                case "":
                                    fmt.Println(fi18, p2)
                                case "\n\t":
                                    _ = len(v4)
                                    if false {
                                        v23 := GenS1{F0: p2}
                                        fmt.Println(p0, p0)
                                        switch (("\n\t" + "hello") + "a") {
                                        case "":
                                            v24 := GenS4{F0: 213, F1: 42, F2: 127}
                                            _ = v24
                                        }
                                        _ = v23
                                    } else {
                                        p1 += (^(fi9))
                                    }
                                    fmt.Println(p0)
                                default:
                                    if (true || true) {
                                        v25 := GenS2{F0: 2, F1: 255}
                                        v26 := GenS4{F0: 151, F1: 42, F2: 42}
                                        if false {
                                            p2 = (p2 * p2)
                                            p0 -= (p0 / ((p0) | 1))
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        } else {
                                            v27 := GenS2{F0: 399, F1: 876}
                                            var v28 int = fi18
                                            _ = v27
                                            _ = v28
                                        }
                                        _ = v25
                                        _ = v26
                                    }
                                }
                            case "\n\t":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case "":
                                var v29 int = (p1 + (p1 & (fi18 + int(p2))))
                                _ = v29
                            }
                        }
                        p1 -= fi9
                        if false {
                            switch fi18 {
                            case 9223372036854775807:
                                v30 := GenS5{F0: 1, F1: 798, F2: -1}
                                _ = v30
                            case 699:
                                var v31 uint16 = 65535
                                _ = v31
                            case 2147483647:
                                if true {
                                    switch (fi18 >> ((18446744073709551615) & 63)) {
                                    case 2147483647:
                                        var v32 uint32 = 2147483648
                                        _ = v32
                                    default:
                                        v33 := GenS4{F0: 406, F1: 18446744073709551615, F2: 0}
                                        _ = v33
                                    }
                                    if true {
                                        if (true && false) {
                                            fmt.Println(fi18, p1, fi9)
                                            v34 := GenS3{F0: -128}
                                            _ = v34
                                        } else {
                                            p0 += (p0 % ((42) | 1))
                                            switch (fi18 | 128) {
                                            case 2147483647:
                                                _ = len(v4)
                                            case -9223372036854775808:
                                                if len(v4) > 0 {
                                                    v4[0] = fi9
                                                }
                                            case 0:
                                                p1 += fi9
                                            default:
                                                switch -319 {
                                                case -9223372036854775808:
                                                    if ((734 < p0) || false) {
                                                        v35 := GenS2{F0: p2, F1: 476}
                                                        _ = len(v4)
                                                        _ = len(v4)
                                                        _ = v35
                                                    } else {
                                                        var v36 uint32 = 0
                                                        _ = v36
                                                    }
                                                case -2147483648:
                                                    p1 -= p1
                                                case 495:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                            }
                                        }
                                        switch 0 {
                                        case -1:
                                            if len(v4) > 0 {
                                                v4[0] = fi9
                                            }
                                        case 42:
                                            if true {
                                                var v37 uint64 = p2
                                                _ = v37
                                            } else {
                                                v38 := GenS2{F0: 4294967295, F1: 18446744073709551615}
                                                fmt.Println(p0, fi18, p2)
                                                v39 := GenS6{F0: 127, F1: 457}
                                                var v40 int64 = -456
                                                if (false || (!(-128 != 85))) {
                                                    if (((true || true) && (127 > 1)) && false) {
                                                        var v41 string = "\n\t"
                                                        v41 = v41
                                                        if true {
                                                            if (uint(p2) == 65535) {
                                                                if (false || (((1.5 > (1.5 - 1.5)) || (!(!true))) && ((p2 ^ p2) <= p2))) {
                                                                    _ = len(v4)
                                                                }
                                                            }
                                                            fmt.Println(fi18, p0)
                                                        }
                                                        _ = v41
                                                    } else {
                                                        if false {
                                                            var v42 int = fi9
                                                            _ = v42
                                                        }
                                                        p0 += (p0 * 506)
                                                        v43 := GenS2{F0: (p2 + p2), F1: 42}
                                                        _ = v43
                                                    }
                                                    v44 := GenS7{F0: 255}
                                                    _ = v44
                                                } else {
                                                    var v45 uint8 = 69
                                                    if len(v4) > 0 {
                                                        v4[0] = p1
                                                    }
                                                    _ = v45
                                                }
                                                _ = v38
                                                _ = v39
                                                _ = v40
                                            }
                                        case -9223372036854775808:
                                            if len(v4) > 0 {
                                                v4[0] = fi18
                                            }
                                        default:
                                            if true {
                                                v4 = append(v4, (fi9 - -208))
                                                v46 := GenS5{F0: 93, F1: 42, F2: int32(p1)}
                                                var v47 float64 = ((1e10 - 1e10) + (float64(1e10) / (1.5 * float64((-2.25 / 1.0)))))
                                                _ = v46
                                                _ = v47
                                            } else {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                switch (^(fi9)) {
                                                case 436:
                                                    v48 := GenS5{F0: p0, F1: 986, F2: -2147483648}
                                                    _ = v48
                                                case 285:
                                                    fmt.Println(fi18)
                                                default:
                                                    fmt.Println(p1)
                                                }
                                                var v49 uint32 = 169
                                                _ = v49
                                            }
                                        }
                                        if (!true) {
                                            switch "" {
                                            case "hello":
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            case "":
                                                switch "" {
                                                case "\n\t":
                                                    var v50 int8 = -1
                                                    var v51 uint32 = 4294967295
                                                    _ = v50
                                                    _ = v51
                                                case "a":
                                                    v52 := GenS4{F0: 2, F1: 800, F2: 44}
                                                    _ = v52
                                                default:
                                                    switch 0 {
                                                    case -260:
                                                        if false {
                                                            v53 := GenS8{F0: 4294967295, F1: -1}
                                                            _ = v53
                                                        } else {
                                                            var v54 int64 = 0
                                                            _ = v54
                                                        }
                                                    case 4294967295:
                                                        if false {
                                                            v4 = append(v4, (fi9 & int(p0)))
                                                            fmt.Println(fi18, fi9)
                                                            switch ("hello" + "foo bar") {
                                                            case "":
                                                                if true {
                                                                    v4 = append(v4, fi9)
                                                                    if ((3.14159 > 1.5) || true) {
                                                                        var v55 string = ("hello" + (("" + "\n\t") + "a"))
                                                                        if true {
                                                                            var v56 uint64 = (^(p2))
                                                                            v56 += (p2 % (((p2 & 237)) | 1))
                                                                            _ = v56
                                                                        }
                                                                        v57 := GenS9{F0: p0, F1: 0}
                                                                        p2 -= p2
                                                                        if (v55 == (v55 + "\n\t")) {
                                                                            var v58 bool = (!false)
                                                                            _ = v58
                                                                        }
                                                                        _ = v55
                                                                        _ = v57
                                                                    } else {
                                                                        var v59 int32 = 346
                                                                        if (!(0 != 42)) {
                                                                            p2 += (p2 & p2)
                                                                            if (true || false) {
                                                                                if (p2 > 532) {
                                                                                    var v60 uint8 = 229
                                                                                    v61 := GenS9{F0: p0, F1: 42}
                                                                                    p1 += fi18
                                                                                    var v62 int = (fi18 + fi18)
                                                                                    _ = v60
                                                                                    _ = v61
                                                                                    _ = v62
                                                                                }
                                                                            }
                                                                            v63 := GenS6{F0: 174, F1: p2}
                                                                            _ = v63
                                                                        } else {
                                                                            v64 := GenS10{F0: (p0 / (((-(p0))) | 1)), F1: 4294967295, F2: fi18}
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                            _ = v64
                                                                        }
                                                                        _ = v59
                                                                    }
                                                                    if true {
                                                                        p0 += (p0 & p0)
                                                                    } else {
                                                                        v65 := GenS11{F0: 699}
                                                                        p2 -= 2
                                                                        _ = len(v4)
                                                                        _ = v65
                                                                    }
                                                                } else {
                                                                    var v66 int16 = p0
                                                                    switch fi9 {
                                                                    case 9223372036854775807:
                                                                        var v67 bool = false
                                                                        _ = v67
                                                                    case -9223372036854775808:
                                                                        v68 := GenS12{F0: fi9}
                                                                        _ = v68
                                                                    case -81:
                                                                        v69 := GenS10{F0: v66, F1: 1, F2: (p1 ^ (p1 - 32767))}
                                                                        v70 := GenS3{F0: -128}
                                                                        _ = v69
                                                                        _ = v70
                                                                    default:
                                                                        v71 := GenS1{F0: (p2 + p2)}
                                                                        _ = v71
                                                                    }
                                                                    _ = v66
                                                                }
                                                            default:
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            }
                                                        } else {
                                                            p0 = p0
                                                            if false {
                                                                var v72 bool = (!false)
                                                                _ = v72
                                                            } else {
                                                                p2 = p2
                                                                if (!true) {
                                                                    p1 = -249
                                                                }
                                                            }
                                                        }
                                                    default:
                                                        if false {
                                                            p2 += (p2 ^ (p2 | (p2 >> ((0) & 63))))
                                                        }
                                                    }
                                                }
                                            default:
                                                p1 = 9223372036854775807
                                            }
                                        } else {
                                            p1 += (fi18 - (^(p1)))
                                            p1 += (p1 >> ((0) & 63))
                                            if (!(!false)) {
                                                switch ("hello" + ("" + ("a" + "\n\t"))) {
                                                case "":
                                                    var v73 float32 = 0.0
                                                    _ = v73
                                                }
                                            }
                                            switch 2147483647 {
                                            case 281:
                                                var v74 float32 = 1e10
                                                _ = v74
                                            case 2147483647:
                                                switch (^(fi18)) {
                                                case 42:
                                                    _ = len(v4)
                                                default:
                                                    if (195 == 255) {
                                                        fmt.Println(p0, fi18, p0)
                                                        if false {
                                                            switch "foo bar" {
                                                            case "a":
                                                                if ((p1 / ((0) | 1)) <= p1) {
                                                                    fmt.Println(fi18, p1, p1)
                                                                    var v75 uint = 4294967295
                                                                    if len(v4) > 0 {
                                                                        v4[0] = (p1 & fi18)
                                                                    }
                                                                    _ = v75
                                                                }
                                                            case "hello":
                                                                p2 = p2
                                                            case "\n\t":
                                                                p1 += p1
                                                            }
                                                            var v76 int8 = -37
                                                            if (4294967295 <= 0) {
                                                                var v77 uint = 130
                                                                p0 -= 658
                                                                fmt.Println(fi9, p0, p0)
                                                                fmt.Println(p2, p2)
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    if len(v4) > 0 {
                                                                        v4[0] = p1
                                                                    }
                                                                default:
                                                                    v77 += v77
                                                                }
                                                                _ = v77
                                                            } else {
                                                                fmt.Println(p0, p2)
                                                                switch (fi9 | -9223372036854775808) {
                                                                case 32767:
                                                                    if (!true) {
                                                                        v78 := GenS13{F0: 382, F1: "", F2: (p2 / (((p2 + p2)) | 1))}
                                                                        _ = v78
                                                                    } else {
                                                                        _ = len(v4)
                                                                        fmt.Println(fi18)
                                                                    }
                                                                case 654:
                                                                    if true {
                                                                        var v79 int64 = 17
                                                                        if (false || (0 <= 447)) {
                                                                            if len(v4) > 0 {
                                                                                v4[0] = fi18
                                                                            }
                                                                        } else {
                                                                            if true {
                                                                                var v80 int16 = int16(p1)
                                                                                switch 127 {
                                                                                case 127:
                                                                                    var v81 uint8 = 42
                                                                                    _ = v81
                                                                                case -2147483648:
                                                                                    if (!true) {
                                                                                        var v82 int8 = -56
                                                                                        _ = v82
                                                                                    } else {
                                                                                        v76 += v76
                                                                                        var v83 int8 = 0
                                                                                        v84 := GenS8{F0: 1, F1: v76}
                                                                                        _ = v83
                                                                                        _ = v84
                                                                                    }
                                                                                }
                                                                                fmt.Println(p0, v76, fi9)
                                                                                _ = v80
                                                                            }
                                                                        }
                                                                        if len(v4) > 0 {
                                                                            v4[0] = (fi9 - 9223372036854775807)
                                                                        }
                                                                        _ = v79
                                                                    }
                                                                default:
                                                                    v4 = append(v4, fi9)
                                                                }
                                                                var v85 bool = true
                                                                if len(v4) > 0 {
                                                                    v4[0] = (p1 % ((231) | 1))
                                                                }
                                                                v85 = (false || v85)
                                                                p0 = p0
                                                                var v86 float32 = 1.5
                                                                p1 = 42
                                                                v86 += float32(float64(v86))
                                                                _ = v85
                                                                _ = v86
                                                            }
                                                            _ = v76
                                                        } else {
                                                            fmt.Println(p2, fi18, fi9)
                                                        }
                                                    }
                                                }
                                                var v87 bool = (p0 >= p0)
                                                _ = v87
                                            case 1:
                                                fmt.Println(p0, p1, p2)
                                            default:
                                                switch ((("a" + ("" + "")) + "hello") + "hello") {
                                                case "":
                                                    p0 += 42
                                                default:
                                                    if false {
                                                        v88 := GenS7{F0: 1}
                                                        fmt.Println(p1, p2, fi9)
                                                        if len(v4) > 0 {
                                                            v4[0] = 127
                                                        }
                                                        _ = v88
                                                    }
                                                }
                                            }
                                        }
                                        switch "foo bar" {
                                        case "hello":
                                            switch "" {
                                            case "":
                                                var v89 float64 = ((1e10 - ((float64(0.0) * 3.14159) / 1e10)) / 1.0)
                                                _ = v89
                                            }
                                        case "a":
                                            p1 = p1
                                        case "":
                                            if false {
                                                var v90 string = "hello"
                                                p1 = (fi9 | -1)
                                                switch ((v90 + ("\n\t" + (("\n\t" + v90) + (v90 + v90)))) + v90) {
                                                case "\n\t":
                                                    fmt.Println(v90)
                                                case "hello":
                                                    p0 = p0
                                                case "a":
                                                    fmt.Println(p0, p2)
                                                default:
                                                    var v91 int32 = 2147483647
                                                    _ = v91
                                                }
                                                _ = v90
                                            } else {
                                                switch "\n\t" {
                                                case "a":
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        fmt.Println(p0, fi9, p1)
                                                        p2 = (p2 * 4294967295)
                                                        v4 = append(v4, fi9)
                                                    case "a":
                                                        v4 = append(v4, 332)
                                                    case "hello":
                                                        fmt.Println(p1)
                                                    default:
                                                        var v92 string = ("hello" + "hello")
                                                        _ = v92
                                                    }
                                                case "hello":
                                                    fmt.Println(p0, p0)
                                                default:
                                                    var v93 float64 = -2.25
                                                    _ = v93
                                                }
                                                v4 = append(v4, fi9)
                                                var v94 uint = 18446744073709551615
                                                v4 = append(v4, fi18)
                                                _ = v94
                                            }
                                        }
                                    }
                                    fmt.Println(fi9, fi18, p0)
                                    var v95 uint8 = 2
                                    _ = v95
                                } else {
                                    if false {
                                        var v96 int16 = p0
                                        p1 = 42
                                        _ = v96
                                    }
                                    var v97 uint = 18446744073709551615
                                    _ = v97
                                }
                            default:
                                v98 := GenS11{F0: 42}
                                _ = v98
                            }
                            if true {
                                var v99 uint16 = 676
                                _ = v99
                            } else {
                                if false {
                                    v100 := GenS6{F0: 37, F1: (p2 / ((1) | 1))}
                                    v101 := GenS3{F0: 42}
                                    switch "a" {
                                    case "foo bar":
                                        var v102 uint32 = 2147483647
                                        _ = v102
                                    case "":
                                        fmt.Println(fi9)
                                    }
                                    _ = v100
                                    _ = v101
                                } else {
                                    p1 = p1
                                    var v103 uint64 = (^(p2))
                                    _ = v103
                                }
                                fmt.Println(p1)
                            }
                        }
                        _ = fi18
                    }
                    _ = fi9
                    _ = v17
                }
                if len(v4) > 0 {
                    v4[0] = p1
                }
            }
        default:
            if true {
                switch 9223372036854775807 {
                case 2147483647:
                    fmt.Println(p2, p0)
                case 0:
                    fmt.Println(p1, p0)
                case -2147483648:
                    switch p1 {
                    case 9223372036854775807:
                        p2 = (^(p2))
                    default:
                        if (!(((1e10 != (((1.5 / 3.14159) - (1e10 + 1e10)) * 0.0)) || ("foo bar" >= "foo bar")) || true)) {
                            v104 := GenS3{F0: 0}
                            for fi105 := 0; fi105 < 7; fi105++ {
                                p2 = 2
                                if true {
                                    switch "\n\t" {
                                    case "\n\t":
                                        if true {
                                            fmt.Println(p2, p1, fi105)
                                            v106 := GenS5{F0: p0, F1: 2, F2: -63}
                                            p1 = fi105
                                            _ = v106
                                        } else {
                                            for fi107 := 0; fi107 < 15; fi107++ {
                                                if (127 != 194) {
                                                    _ = len(v4)
                                                    fmt.Println(fi107, fi107, p0)
                                                }
                                                if (1 == 1) {
                                                    v108 := GenS6{F0: 42, F1: (p2 / ((42) | 1))}
                                                    _ = v108
                                                } else {
                                                    fmt.Println(p0, fi105)
                                                    _ = len(v4)
                                                    p1 = -583
                                                    if len(v4) > 0 {
                                                        v4[0] = 413
                                                    }
                                                }
                                                switch "\n\t" {
                                                case "a":
                                                    p1 += (fi105 * 32768)
                                                case "hello":
                                                    if false {
                                                        p2 += p2
                                                        fmt.Println(fi107, p0, p2)
                                                    }
                                                }
                                                _ = fi107
                                            }
                                            if true {
                                                p0 -= (p0 | p0)
                                                var v109 float32 = 1e10
                                                _ = v109
                                            }
                                            var v110 uint16 = 524
                                            _ = v110
                                        }
                                    }
                                }
                                var v111 int64 = 1
                                _ = fi105
                                _ = v111
                            }
                            _ = v104
                        }
                    }
                }
                p1 = p1
            }
        }
        _ = v4
    } else {
        var v112 string = "a"
        p1 += 685
        switch ((((v112 + "\n\t") + "\n\t") + "") + (v112 + v112)) {
        case "":
            switch v112 {
            case "foo bar":
                fmt.Println(p1)
            case "":
                v113 := []int{p1, p1}
                _ = v113
            case "hello":
                switch 556 {
                case 4294967295:
                    switch v112 {
                    case "a":
                        v114 := GenS14{F0: "a", F1: int16(p2), F2: 9223372036854775807}
                        _ = v114
                    case "hello":
                        switch ((v112 + v112) + "") {
                        case "":
                            switch (p1 & -684) {
                            case 9223372036854775807:
                                var v115 int32 = -2147483648
                                _ = v115
                            case 288:
                                switch p1 {
                                case 0:
                                    switch 127 {
                                    case 127:
                                        var v116 int = (-(p1))
                                        _ = v116
                                    case -812:
                                        var v117 uint = 466
                                        _ = v117
                                    case -839:
                                        for fi118 := 0; fi118 < 14; fi118++ {
                                            p0 -= p0
                                            _ = fi118
                                        }
                                    default:
                                        for fi119 := 0; fi119 < 20; fi119++ {
                                            var v120 string = ("a" + v112)
                                            v121 := GenS6{F0: 127, F1: p2}
                                            _ = fi119
                                            _ = v120
                                            _ = v121
                                        }
                                    }
                                default:
                                    v122 := GenS4{F0: 0, F1: 623, F2: 100}
                                    _ = v122
                                }
                            case 281:
                                var v123 uint16 = 32767
                                _ = v123
                            }
                        }
                    case "foo bar":
                        fmt.Println(p0, v112)
                    }
                case 127:
                    if true {
                        v112 = (("" + ("\n\t" + "\n\t")) + "hello")
                        fmt.Println(p1)
                        switch ("" + (((v112 + ("foo bar" + v112)) + v112) + v112)) {
                        case "hello":
                            v124 := GenS15{F0: false, F1: 42, F2: p2}
                            _ = v124
                        }
                    }
                case 0:
                    for fi125 := 0; fi125 < 14; fi125++ {
                        for fi126 := 0; fi126 < 10; fi126++ {
                            v127 := []int{(-(p1))}
                            var v128 uint16 = 42
                            _ = fi126
                            _ = v127
                            _ = v128
                        }
                        fmt.Println(v112)
                        _ = fi125
                    }
                }
            default:
                var v129 float64 = -2.25
                _ = v129
            }
        case "\n\t":
            for fi130 := 0; fi130 < 7; fi130++ {
                fmt.Println(p1, p2)
                _ = fi130
            }
        case "a":
            v131 := GenS16{F0: 1}
            _ = v131
        }
        _ = v112
    }
    if (!(2 >= 622)) {
        v132 := GenS17{F0: 0, F1: 992, F2: uint16(p2)}
        var v133 int = 903
        _ = v132
        _ = v133
    } else {
        for fi134 := 0; fi134 < 15; fi134++ {
            v135 := GenS18{F0: 42, F1: 42}
            if (true && false) {
                v136 := []int{(fi134 - (fi134 | 772)), p1, 32768}
                if (!(42 == 595)) {
                    v137 := GenS19{F0: ((false && false) || false)}
                    var v138 uint = 4294967295
                    if false {
                        v136 = append(v136, p1)
                        if len(v136) > 0 {
                            v136[0] = p1
                        }
                    } else {
                        switch "a" {
                        case "foo bar":
                            p1 += fi134
                        case "hello":
                            if (-559 == 2147483647) {
                                switch (p1 >> ((42) & 63)) {
                                case -929:
                                    for fi139 := 0; fi139 < 10; fi139++ {
                                        p2 -= (p2 * 18446744073709551615)
                                        if len(v136) > 0 {
                                            v136[0] = fi139
                                        }
                                        _ = fi139
                                    }
                                case 32768:
                                    v140 := GenS20{F0: "hello", F1: -9223372036854775808}
                                    _ = v140
                                case 127:
                                    switch (^(p1)) {
                                    case 1:
                                        for fi141 := 0; fi141 < 5; fi141++ {
                                            v138 += 1
                                            _ = fi141
                                        }
                                    case 32767:
                                        var v142 string = ("a" + ("a" + ""))
                                        _ = v142
                                    case -767:
                                        var v143 uint = 580
                                        _ = v143
                                    }
                                }
                            } else {
                                v144 := GenS13{F0: 0, F1: ("foo bar" + "\n\t"), F2: (^(p2))}
                                v145 := GenS19{F0: (18446744073709551615 >= v138)}
                                if len(v136) > 0 {
                                    v136[0] = (fi134 ^ 40)
                                }
                                _ = v144
                                _ = v145
                            }
                        case "":
                            switch ("\n\t" + ("\n\t" + ("hello" + "a"))) {
                            case "\n\t":
                                var v146 int64 = 1
                                _ = v146
                            default:
                                fmt.Println(fi134, p0, fi134)
                            }
                        default:
                            var v147 string = ("\n\t" + ("a" + "\n\t"))
                            _ = v147
                        }
                        v138 = (v138 - 255)
                        var v148 int8 = 37
                        _ = v148
                    }
                    _ = v137
                    _ = v138
                }
                _ = v136
            }
            _ = fi134
            _ = v135
        }
    }
    var v149 int64 = 536
    fmt.Println(p2, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v149
    return 1
}
func genFn1(p0 int32, p1 int8, p2 int32) uint16 {
    var v150 float64 = (0.1 * ((0.0 / (1.5 * 1e10)) * 3.14159))
    var v151 float64 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v150
    _ = v151
    return 15
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[int]string)
    v3 := []uint32{}
    fmt.Println(v1)
    var v152 bool = false
    v1 = v1
    for fi153 := 0; fi153 < 6; fi153++ {
        var v154 float64 = 0.0
        _ = fi153
        _ = v154
    }
    if true {
        if (1 <= 261) {
            var v155 float64 = 0.0
            switch "\n\t" {
            case "\n\t":
                fmt.Println(v1, v155, v152)
            }
            for fi156 := 0; fi156 < 17; fi156++ {
                for fi157 := 0; fi157 < 8; fi157++ {
                    var v158 uint64 = 1
                    _ = fi157
                    _ = v158
                }
                for fi159 := 0; fi159 < 6; fi159++ {
                    v160 := GenS16{F0: 0}
                    if (0 != 179) {
                        v152 = v152
                        if v152 {
                            v155 = float64((float32(v155) * 3.14159))
                        } else {
                            _ = len(v2)
                            delete(v2, 4294967295)
                        }
                        v161 := GenS7{F0: 127}
                        v162 := GenS14{F0: "hello", F1: 689, F2: 337}
                        _ = v161
                        _ = v162
                    }
                    v163 := GenS2{F0: 4294967295, F1: 1}
                    _ = fi159
                    _ = v160
                    _ = v163
                }
                v1 = (v1 - 2)
                _ = fi156
            }
            _ = v155
        }
        v164 := GenS17{F0: -527, F1: 257, F2: genFn1(1, -91, 42)}
        _ = v164
    } else {
        v152 = false
        var v165 string = ""
        _ = v165
    }
    if (!(true && (18446744073709551615 >= 0))) {
        if v152 {
            if (!true) {
                if false {
                    for fi166 := 0; fi166 < 7; fi166++ {
                        if v152 {
                            var v167 bool = false
                            _ = v167
                        } else {
                            _ = len(v2)
                            var v168 int32 = 0
                            _ = v168
                        }
                        var v169 int = 32767
                        _ = fi166
                        _ = v169
                    }
                    v170 := GenS21{F0: 148, F1: (v152 && false)}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v170
                }
            }
            var v171 int32 = 42
            v172 := GenS22{F0: -483}
            _ = v171
            _ = v172
        } else {
            var v173 uint = 4294967295
            _ = v173
        }
    } else {
        var v174 uint16 = 255
        _ = v174
    }
    if (v152 || v152) {
        v175 := GenS23{F0: 42, F1: 0}
        _ = v175
    } else {
        var v176 uint32 = v1
        v1 -= 0
        fmt.Println(v176)
        _ = v176
    }
    v177 := GenS14{F0: "\n\t", F1: 1000, F2: -791}
    delete(v2, 42)
    for fi178 := 0; fi178 < 12; fi178++ {
        fmt.Println(v152, fi178)
        _ = fi178
    }
    v152 = v152
    v179 := GenS24{F0: (v1 & v1)}
    switch 4294967295 {
    case 364:
        _ = len(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v152:", v152)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v152
    _ = v177
    _ = v179
    _ = v2
    _ = v3
}
