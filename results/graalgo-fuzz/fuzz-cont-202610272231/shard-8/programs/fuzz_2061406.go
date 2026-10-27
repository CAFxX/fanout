package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 uint; F1 uint32; F2 uint }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int; F1 uint }
type GenS11 struct { F0 uint8; F1 uint64 }
type GenS12 struct { F0 int16; F1 bool }
type GenS13 struct { F0 int32; F1 int64 }
type GenS14 struct { F0 uint; F1 bool }
type GenS15 struct { F0 int64; F1 uint }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 string }
type GenS18 struct { F0 int; F1 int64 }
type GenS19 struct { F0 uint; F1 uint8 }
type GenS20 struct { F0 uint8; F1 string }
type GenS21 struct { F0 string; F1 bool; F2 int32 }
type GenS22 struct { F0 uint8; F1 uint64 }
type GenS23 struct { F0 int8 }
type GenS24 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS25 struct { F0 uint; F1 uint32; F2 int }
type GenS26 struct { F0 int64 }
type GenS27 struct { F0 uint; F1 uint16 }
type GenS28 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS29 struct { F0 uint8; F1 int8 }
type GenS30 struct { F0 string; F1 uint8; F2 uint32 }
type GenS31 struct { F0 int16; F1 bool }

func main() {
    var v1 uint = 668
    var v2 int = 42
    var v3 uint = v1
    v4 := GenS1{F0: -1, F1: 225}
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        v5 := []int{v2, (v2 / ((v2) | 1))}
        _ = v5
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        fmt.Println(fi6, fi6, fi6)
        switch (^(v2)) {
        case 4294967295:
            if true {
                var v7 uint64 = 18446744073709551615
                if false {
                    var v8 int32 = -26
                    fmt.Println(v2)
                    _ = v8
                }
                v9 := []int{-349, (v2 + (-(fi6))), v2, (fi6 | -905)}
                _ = v7
                _ = v9
            } else {
                v10 := GenS2{F0: 134, F1: 0, F2: 42}
                v11 := []int{(v2 | v2)}
                _ = v10
                _ = v11
            }
        case -176:
            v12 := []int{1, (-(v2)), 0, v2}
            _ = v12
        }
        v13 := []int{fi6, v2}
        _ = fi6
        _ = v13
    }
    v14 := []int{v2, 32768, (v2 >> ((v1) & 63)), (-(v2))}
    for fi15 := 0; fi15 < 12; fi15++ {
        if false {
            v1 += v1
        }
        if (!true) {
            v1 = (v1 << ((v3) & 63))
        }
        _ = fi15
    }
    v2 = (v2 + v2)
    if (42 > 65535) {
        if (int(v3) >= v2) {
            v3 = (^(v1))
            var v16 uint = 1
            _ = v16
        } else {
            if (255 > 32767) {
                var v17 uint = (v3 % ((0) | 1))
                for fi18 := 0; fi18 < 16; fi18++ {
                    if len(v14) > 0 {
                        v14[0] = v2
                    }
                    _ = len(v14)
                    _ = fi18
                }
                for fi19 := 0; fi19 < 9; fi19++ {
                    if ((false || true) && (((("foo bar" + "a") + "\n\t") + ("a" + "a")) > "\n\t")) {
                        switch (v2 & v2) {
                        case 1:
                            var v20 float32 = (-2.25 - (1.5 + -2.25))
                            _ = v20
                        case 8:
                            for fi21 := 0; fi21 < 6; fi21++ {
                                switch (fi19 * (v2 | -2147483648)) {
                                case 778:
                                    switch fi21 {
                                    case 205:
                                        _ = len(v14)
                                    case 127:
                                        var v22 int8 = -38
                                        _ = v22
                                    default:
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                    }
                                case 967:
                                    switch (fi19 + 2147483647) {
                                    case -541:
                                        if false {
                                            var v23 float32 = -2.25
                                            _ = v23
                                        }
                                        v2 -= (fi19 % ((v2) | 1))
                                    case 387:
                                        if (true || ("foo bar" < "hello")) {
                                            v17 = 42
                                        }
                                    }
                                default:
                                    v14 = append(v14, 42)
                                }
                                v2 = fi19
                                fmt.Println(v2)
                                if true {
                                    switch "\n\t" {
                                    case "":
                                        switch v2 {
                                        case 4294967295:
                                            var v24 uint8 = 27
                                            _ = v24
                                        default:
                                            v2 -= (v2 * fi19)
                                        }
                                    case "foo bar":
                                        fmt.Println(v2, fi21, fi21)
                                    case "a":
                                        v2 = 1
                                    }
                                    v2 += 42
                                    switch v2 {
                                    case 32767:
                                        v25 := GenS2{F0: 255, F1: 42, F2: 197}
                                        _ = v25
                                    case 128:
                                        v2 = (fi19 & fi21)
                                    default:
                                        var v26 int64 = -829
                                        _ = v26
                                    }
                                    v1 += v3
                                    var v27 uint = v17
                                    _ = v27
                                }
                                if (!((2 != 0) || false)) {
                                    v1 = v17
                                    if (false || true) {
                                        if (((3.14159 * 1e10) / 0.1) > 1e10) {
                                            v28 := GenS1{F0: -95, F1: 2}
                                            var v29 string = "\n\t"
                                            v30 := GenS3{F0: 806, F1: 255}
                                            _ = len(v14)
                                            if (1 <= -352) {
                                                var v31 int16 = -459
                                                var v32 float32 = 0.0
                                                v29 = v29
                                                v1 += v1
                                                _ = v31
                                                _ = v32
                                            } else {
                                                fmt.Println(v17)
                                            }
                                            v17 = 925
                                            _ = v28
                                            _ = v29
                                            _ = v30
                                        } else {
                                            if (false && false) {
                                                var v33 float32 = (((3.14159 * 3.14159) - ((1e10 / 3.14159) / -2.25)) + float32(((3.14159 / 1.0) + 0.1)))
                                                var v34 bool = ("hello" <= ("foo bar" + "\n\t"))
                                                v34 = ((2147483647 <= 65535) && false)
                                                _ = v33
                                                _ = v34
                                            }
                                            v3 += (v17 << ((255) & 63))
                                        }
                                        v1 -= (v1 | 2)
                                        if false {
                                            if true {
                                                v3 = v1
                                            }
                                            fmt.Println(fi19, v2, v2)
                                        }
                                        if false {
                                            v2 += (fi19 + v2)
                                        } else {
                                            v3 -= v3
                                            if len(v14) > 0 {
                                                v14[0] = 4294967295
                                            }
                                            if len(v14) > 0 {
                                                v14[0] = (-(fi19))
                                            }
                                        }
                                        _ = len(v14)
                                        fmt.Println(fi19, v3)
                                        if (("hello" + "hello") >= ("\n\t" + "hello")) {
                                            v17 = v17
                                            var v35 uint8 = 1
                                            _ = v35
                                        } else {
                                            v2 -= 42
                                            v3 = (v1 - 18446744073709551615)
                                            if false {
                                                v1 -= 2
                                                fmt.Println(v1)
                                            } else {
                                                var v36 float64 = (1.5 * (-2.25 / 3.14159))
                                                _ = v36
                                            }
                                        }
                                        if true {
                                            v37 := GenS4{F0: -1, F1: 417, F2: 112}
                                            _ = v37
                                        } else {
                                            if true {
                                                v38 := GenS5{F0: -109}
                                                _ = v38
                                            } else {
                                                var v39 float64 = (float64(-2.25) + 1.5)
                                                if true {
                                                    if (!(431 >= v3)) {
                                                        if (273 < 1000) {
                                                            v40 := GenS3{F0: 189, F1: 137}
                                                            switch 32767 {
                                                            case 42:
                                                                v2 = fi19
                                                            case 127:
                                                                if false {
                                                                    if true {
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            v14 = append(v14, 266)
                                                                        case "a":
                                                                            _ = len(v14)
                                                                        }
                                                                        var v41 float64 = (v39 + 3.14159)
                                                                        _ = v41
                                                                    }
                                                                    var v42 uint8 = 162
                                                                    v17 -= 1
                                                                    _ = v42
                                                                }
                                                            case -9223372036854775808:
                                                                v39 -= v39
                                                            default:
                                                                v2 -= (^(fi21))
                                                                v43 := GenS6{F0: 1, F1: 4294967295}
                                                                var v44 int32 = -528
                                                                _ = v43
                                                                _ = v44
                                                            }
                                                            fmt.Println(v3, fi21, fi21)
                                                            _ = v40
                                                        }
                                                        fmt.Println(v1, v17)
                                                        switch (fi21 << (((v1 % (((v1 / ((v3) | 1))) | 1))) & 63)) {
                                                        case 9223372036854775807:
                                                            if true {
                                                                if ((v39 == v39) && false) {
                                                                    var v45 bool = (!true)
                                                                    var v46 uint = (^(v17))
                                                                    var v47 string = ("\n\t" + ("" + ("a" + "")))
                                                                    _ = v45
                                                                    _ = v46
                                                                    _ = v47
                                                                } else {
                                                                    var v48 float64 = (v39 / 1.0)
                                                                    _ = v48
                                                                }
                                                                var v49 uint32 = 114
                                                                _ = v49
                                                            }
                                                        default:
                                                            v50 := GenS4{F0: -746, F1: 0, F2: 123}
                                                            _ = v50
                                                        }
                                                    } else {
                                                        var v51 bool = true
                                                        _ = v51
                                                    }
                                                    if (true || true) {
                                                        v52 := GenS5{F0: 95}
                                                        _ = v52
                                                    } else {
                                                        if (0 < 2) {
                                                            fmt.Println(v39, v17)
                                                            if len(v14) > 0 {
                                                                v14[0] = (fi19 / ((128) | 1))
                                                            }
                                                        } else {
                                                            v1 += (v3 << (((v3 & 65535)) & 63))
                                                        }
                                                    }
                                                    if true {
                                                        if (42 >= 42) {
                                                            v14 = append(v14, (v2 & fi19))
                                                            if true {
                                                                v3 -= (v17 | (v17 & (v3 << (((v1 << ((291) & 63))) & 63))))
                                                                v39 = (v39 * v39)
                                                            }
                                                            if len(v14) >= 2 {
                                                                v14 = v14[1:len(v14)]
                                                            }
                                                        }
                                                    } else {
                                                        v17 = (^(v3))
                                                    }
                                                } else {
                                                    v53 := GenS7{F0: "", F1: -719}
                                                    if false {
                                                        var v54 uint16 = 32767
                                                        _ = v54
                                                    } else {
                                                        if true {
                                                            switch v2 {
                                                            case 846:
                                                                var v55 int8 = 92
                                                                _ = v55
                                                            case -9223372036854775808:
                                                                fmt.Println(v2, v17, v1)
                                                            default:
                                                                v56 := GenS2{F0: 208, F1: -1, F2: 42}
                                                                _ = v56
                                                            }
                                                            var v57 uint16 = 65535
                                                            _ = v57
                                                        } else {
                                                            v58 := GenS8{F0: v17, F1: 42, F2: v1}
                                                            if false {
                                                                if (true || ((false && false) && (("" != ("\n\t" + "\n\t")) && false))) {
                                                                    var v59 float64 = v39
                                                                    var v60 string = "hello"
                                                                    var v61 int32 = 1
                                                                    _ = v59
                                                                    _ = v60
                                                                    _ = v61
                                                                } else {
                                                                    switch fi19 {
                                                                    case 1:
                                                                        fmt.Println(v1)
                                                                    }
                                                                }
                                                            }
                                                            var v62 uint8 = 0
                                                            _ = v58
                                                            _ = v62
                                                        }
                                                        switch 42 {
                                                        case 32768:
                                                            var v63 int64 = 430
                                                            _ = v63
                                                        case 32767:
                                                            switch ("a" + "a") {
                                                            case "\n\t":
                                                                if (((v39 - (v39 + v39)) <= (-2.25 + (v39 * float64(0.0)))) && (true && (fi19 != 4294967295))) {
                                                                    var v64 int16 = -137
                                                                    switch (fi21 ^ fi19) {
                                                                    case 32768:
                                                                        switch "a" {
                                                                        case "\n\t":
                                                                            var v65 float64 = (0.0 * 3.14159)
                                                                            _ = v65
                                                                        case "hello":
                                                                            var v66 float32 = float32(v39)
                                                                            _ = v66
                                                                        case "foo bar":
                                                                            v3 += 0
                                                                        }
                                                                    case -1:
                                                                        if (false || (false && (3.14159 < 1.5))) {
                                                                            if (((4294967295 >= uint32(v1)) && (1 != fi19)) || false) {
                                                                                var v67 bool = false
                                                                                switch (fi19 ^ 637) {
                                                                                case 0:
                                                                                    fmt.Println(v67, v17)
                                                                                case 9223372036854775807:
                                                                                    var v68 uint = (v17 / ((v17) | 1))
                                                                                    _ = v68
                                                                                case 32768:
                                                                                    var v69 uint8 = 187
                                                                                    _ = v69
                                                                                }
                                                                                _ = v67
                                                                            }
                                                                        } else {
                                                                            v1 -= (v17 - v17)
                                                                        }
                                                                    default:
                                                                        var v70 uint16 = 242
                                                                        _ = v70
                                                                    }
                                                                    _ = v64
                                                                }
                                                            case "":
                                                                var v71 int16 = 5
                                                                _ = v71
                                                            default:
                                                                if false {
                                                                    if false {
                                                                        var v72 uint8 = 255
                                                                        _ = v72
                                                                    } else {
                                                                        if len(v14) > 0 {
                                                                            v14[0] = -9223372036854775808
                                                                        }
                                                                        if len(v14) > 0 {
                                                                            v14[0] = (fi19 & (-(v2)))
                                                                        }
                                                                    }
                                                                    switch -927 {
                                                                    case -121:
                                                                        if (fi21 <= fi21) {
                                                                            var v73 int32 = 957
                                                                            _ = v73
                                                                        } else {
                                                                            var v74 float64 = v39
                                                                            if true {
                                                                                var v75 int = v2
                                                                                v76 := GenS3{F0: 0, F1: 127}
                                                                                _ = v75
                                                                                _ = v76
                                                                            }
                                                                            switch -1 {
                                                                            case -967:
                                                                                v74 = 3.14159
                                                                            case 9223372036854775807:
                                                                                if true {
                                                                                    if (!false) {
                                                                                        var v77 string = "hello"
                                                                                        var v78 uint = v17
                                                                                        v14 = append(v14, (v2 << (((v3 - (v17 - (v1 << ((375) & 63))))) & 63)))
                                                                                        if true {
                                                                                            var v79 float64 = ((v74 * (float64(float32((((((v39 / v74) * float64(((0.0 * 1e10) - 1e10))) - (v39 * (v39 + 3.14159))) / (v74 + 1e10)) * 0.0))) / 1.0)) - (v39 / ((v74 * v74) / v74)))
                                                                                            _ = v79
                                                                                        }
                                                                                        v2 = fi19
                                                                                        v80 := GenS6{F0: 4294967295, F1: 814}
                                                                                        _ = v77
                                                                                        _ = v78
                                                                                        _ = v80
                                                                                    } else {
                                                                                        if false {
                                                                                            fmt.Println(fi21, v39, v1)
                                                                                            var v81 uint32 = 4294967295
                                                                                            v74 -= (3.14159 * v39)
                                                                                            v2 -= fi21
                                                                                            v1 = (v3 << (((v17 & v3)) & 63))
                                                                                            v82 := GenS1{F0: 42, F1: v81}
                                                                                            _ = v81
                                                                                            _ = v82
                                                                                        } else {
                                                                                            v83 := GenS5{F0: -1}
                                                                                            if len(v14) > 0 {
                                                                                                v14[0] = fi21
                                                                                            }
                                                                                            switch 949 {
                                                                                            case 210:
                                                                                                var v84 uint64 = 244
                                                                                                _ = v84
                                                                                            case 0:
                                                                                                var v85 uint = (v17 - v3)
                                                                                                if false {
                                                                                                    v14 = append(v14, fi21)
                                                                                                }
                                                                                                switch ((("foo bar" + "hello") + "") + "a") {
                                                                                                case "":
                                                                                                    v3 -= v1
                                                                                                default:
                                                                                                    var v86 int32 = 604
                                                                                                    if false {
                                                                                                        _ = len(v14)
                                                                                                    } else {
                                                                                                        var v87 uint = (v17 / (((v3 - (^(v3)))) | 1))
                                                                                                        _ = v87
                                                                                                    }
                                                                                                    var v88 int = v2
                                                                                                    _ = v86
                                                                                                    _ = v88
                                                                                                }
                                                                                                _ = v85
                                                                                            case -77:
                                                                                                if len(v14) >= 2 {
                                                                                                    v14 = v14[1:len(v14)]
                                                                                                }
                                                                                            }
                                                                                            _ = v83
                                                                                        }
                                                                                        switch "\n\t" {
                                                                                        case "hello":
                                                                                            switch ("" + "foo bar") {
                                                                                            case "hello":
                                                                                                var v89 int32 = 0
                                                                                                _ = v89
                                                                                            default:
                                                                                                if true {
                                                                                                    if true {
                                                                                                        v3 = 2
                                                                                                    } else {
                                                                                                        v90 := GenS3{F0: 255, F1: 235}
                                                                                                        _ = v90
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        case "\n\t":
                                                                                            switch ("\n\t" + "a") {
                                                                                            case "foo bar":
                                                                                                v91 := GenS9{F0: 0}
                                                                                                _ = v91
                                                                                            case "\n\t":
                                                                                                fmt.Println(fi19, fi21)
                                                                                            case "":
                                                                                                if false {
                                                                                                    var v92 uint = 2
                                                                                                    switch "hello" {
                                                                                                    case "\n\t":
                                                                                                        v93 := GenS10{F0: -520, F1: 255}
                                                                                                        _ = v93
                                                                                                    }
                                                                                                    _ = v92
                                                                                                }
                                                                                            default:
                                                                                                var v94 uint64 = 857
                                                                                                _ = v94
                                                                                            }
                                                                                        case "":
                                                                                            v95 := GenS8{F0: 255, F1: 0, F2: v17}
                                                                                            _ = v95
                                                                                        default:
                                                                                            fmt.Println(v74, v2, v74)
                                                                                            v96 := GenS9{F0: 127}
                                                                                            v74 -= ((((0.0 / (v74 / v39)) * v74) + v39) / v74)
                                                                                            _ = v96
                                                                                        }
                                                                                        var v97 uint = (v1 ^ v1)
                                                                                        if len(v14) >= 2 {
                                                                                            v14 = v14[1:len(v14)]
                                                                                        }
                                                                                        v98 := GenS2{F0: 73, F1: -1, F2: 1}
                                                                                        _ = v97
                                                                                        _ = v98
                                                                                    }
                                                                                }
                                                                            case 567:
                                                                                var v99 int16 = int16(fi21)
                                                                                _ = v99
                                                                            }
                                                                            _ = v74
                                                                        }
                                                                    default:
                                                                        var v100 int = (fi21 ^ v2)
                                                                        _ = v100
                                                                    }
                                                                } else {
                                                                    switch 0 {
                                                                    case 475:
                                                                        var v101 uint8 = 42
                                                                        var v102 uint32 = 65535
                                                                        if (!true) {
                                                                            var v103 int64 = -1
                                                                            _ = v103
                                                                        }
                                                                        _ = v101
                                                                        _ = v102
                                                                    case -2147483648:
                                                                        v2 = 9223372036854775807
                                                                    case 9223372036854775807:
                                                                        var v104 uint16 = 569
                                                                        _ = v104
                                                                    default:
                                                                        var v105 int8 = 1
                                                                        _ = v105
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(fi19, fi19, v3)
                                                    _ = v53
                                                }
                                                if ((v39 <= (1e10 + float64(float32(1.5)))) && (!true)) {
                                                    if len(v14) >= 2 {
                                                        v14 = v14[1:len(v14)]
                                                    }
                                                    var v106 uint64 = 2
                                                    v39 += (v39 / (((1.5 + (0.0 * 1.5)) - (3.14159 * 3.14159)) - (0.1 + v39)))
                                                    _ = v106
                                                } else {
                                                    var v107 int64 = 1
                                                    _ = v107
                                                }
                                                _ = v39
                                            }
                                            switch (fi19 / (((-(fi21))) | 1)) {
                                            case -2147483648:
                                                v1 -= 312
                                            case 32767:
                                                var v108 uint = (v1 & 255)
                                                _ = v108
                                            case 32768:
                                                v17 += (v1 % ((4294967295) | 1))
                                            }
                                        }
                                    } else {
                                        v109 := GenS11{F0: 128, F1: 4294967295}
                                        switch "a" {
                                        case "\n\t":
                                            if (!true) {
                                                switch 127 {
                                                case 32767:
                                                    v1 -= v1
                                                case -148:
                                                    v110 := GenS12{F0: 0, F1: (true || true)}
                                                    _ = v110
                                                case 4294967295:
                                                    if len(v14) >= 2 {
                                                        v14 = v14[1:len(v14)]
                                                    }
                                                }
                                                fmt.Println(fi19, v17, fi19)
                                                v17 -= (v3 - v3)
                                            }
                                        default:
                                            v3 = v17
                                        }
                                        _ = v109
                                    }
                                } else {
                                    _ = len(v14)
                                    switch (-(fi19)) {
                                    case 32768:
                                        var v111 uint8 = 143
                                        _ = v111
                                    case -329:
                                        if (false || true) {
                                            switch ("hello" + "hello") {
                                            case "foo bar":
                                                var v112 string = "a"
                                                _ = v112
                                            case "":
                                                v113 := GenS2{F0: 255, F1: 1, F2: 0}
                                                _ = v113
                                            case "hello":
                                                if true {
                                                    fmt.Println(v17, v3)
                                                    if (1e10 >= 1e10) {
                                                        if (true || false) {
                                                            v1 -= 4294967295
                                                        } else {
                                                            fmt.Println(v3, fi21, v2)
                                                        }
                                                        v2 = 1
                                                    } else {
                                                        _ = len(v14)
                                                        if true {
                                                            v3 += v17
                                                            v114 := GenS10{F0: fi21, F1: 2}
                                                            if true {
                                                                fmt.Println(v1, v1)
                                                                _ = len(v14)
                                                                v3 = (v3 * (v3 & 255))
                                                                if len(v14) > 0 {
                                                                    v14[0] = -318
                                                                }
                                                            } else {
                                                                v2 = fi19
                                                                v115 := GenS13{F0: 349, F1: 1}
                                                                _ = v115
                                                            }
                                                            _ = v114
                                                        } else {
                                                            if false {
                                                                v1 -= v17
                                                                v116 := GenS10{F0: fi21, F1: v3}
                                                                _ = v116
                                                            } else {
                                                                _ = len(v14)
                                                                v2 -= -412
                                                            }
                                                            v3 += (v1 & 752)
                                                        }
                                                    }
                                                }
                                            }
                                            if true {
                                                v2 = fi21
                                            } else {
                                                _ = len(v14)
                                                switch "\n\t" {
                                                case "hello":
                                                    v2 -= (fi21 / (((fi19 | 1)) | 1))
                                                default:
                                                    if ((v3 * v1) != v1) {
                                                        var v117 int32 = 0
                                                        v17 -= (^(v3))
                                                        if len(v14) >= 2 {
                                                            v14 = v14[1:len(v14)]
                                                        }
                                                        _ = v117
                                                    } else {
                                                        v1 = (v3 & v17)
                                                        switch "a" {
                                                        case "":
                                                            var v118 bool = true
                                                            _ = v118
                                                        case "\n\t":
                                                            v119 := GenS7{F0: ("" + ""), F1: 0}
                                                            v120 := GenS11{F0: 2, F1: 345}
                                                            fmt.Println(v2, fi21, v3)
                                                            _ = v119
                                                            _ = v120
                                                        default:
                                                            v121 := GenS14{F0: 4294967295, F1: true}
                                                            _ = v121
                                                        }
                                                        switch (v2 << ((v17) & 63)) {
                                                        case -47:
                                                            if len(v14) >= 2 {
                                                                v14 = v14[1:len(v14)]
                                                            }
                                                        case 946:
                                                            if true {
                                                                v122 := GenS15{F0: 42, F1: v17}
                                                                var v123 bool = false
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    var v124 float64 = ((3.14159 * 1.5) - (1e10 - (1.5 + (0.1 + ((1e10 * (1e10 - (0.1 + 1.5))) / 1.5)))))
                                                                    _ = v124
                                                                case "foo bar":
                                                                    if (1e10 != -2.25) {
                                                                        v1 += 42
                                                                    }
                                                                }
                                                                _ = v122
                                                                _ = v123
                                                            } else {
                                                                v1 -= v3
                                                                v2 = fi19
                                                            }
                                                            switch fi21 {
                                                            case 1:
                                                                v2 = (fi21 << (((v3 | (v17 * (v3 + 65535)))) & 63))
                                                            default:
                                                                v17 = v1
                                                            }
                                                            v2 = (v2 ^ fi21)
                                                            v14 = append(v14, (fi21 | fi21))
                                                        }
                                                    }
                                                }
                                                if len(v14) >= 2 {
                                                    v14 = v14[1:len(v14)]
                                                }
                                                v2 = fi19
                                                var v125 int64 = 0
                                                var v126 float32 = 0.1
                                                _ = v125
                                                _ = v126
                                            }
                                        }
                                        if (32767 < 32767) {
                                            v127 := GenS16{F0: 65535}
                                            _ = v127
                                        } else {
                                            v128 := GenS17{F0: "hello"}
                                            var v129 int64 = 42
                                            fmt.Println(fi19, v129, v1)
                                            if (true || false) {
                                                v130 := GenS18{F0: fi21, F1: v129}
                                                switch 0 {
                                                case -9223372036854775808:
                                                    v131 := GenS19{F0: v3, F1: 255}
                                                    _ = v131
                                                case 9223372036854775807:
                                                    if true {
                                                        fmt.Println(v2, v1, v1)
                                                        switch ("a" + "\n\t") {
                                                        case "hello":
                                                            v132 := GenS20{F0: 42, F1: "a"}
                                                            _ = v132
                                                        default:
                                                            v14 = append(v14, (fi19 % (((-(fi19))) | 1)))
                                                        }
                                                        fmt.Println(v17, v17, v17)
                                                    } else {
                                                        var v133 int64 = (v129 | 9223372036854775807)
                                                        _ = v133
                                                    }
                                                }
                                                if (v2 > v2) {
                                                    v129 = v129
                                                    if len(v14) > 0 {
                                                        v14[0] = (fi21 ^ (fi21 % (((v2 | (v2 + v2))) | 1)))
                                                    }
                                                }
                                                _ = v130
                                            }
                                            _ = v128
                                            _ = v129
                                        }
                                    case 0:
                                        var v134 float32 = -2.25
                                        _ = v134
                                    }
                                    var v135 uint = (v3 | 65535)
                                    _ = len(v14)
                                    _ = v135
                                }
                                v17 += 42
                                v136 := GenS3{F0: 65535, F1: 255}
                                var v137 uint32 = 2147483648
                                _ = fi21
                                _ = v136
                                _ = v137
                            }
                        }
                        if (!(!(42 >= 42))) {
                            for fi138 := 0; fi138 < 6; fi138++ {
                                v17 -= v3
                                if false {
                                    switch v2 {
                                    case 128:
                                        fmt.Println(v17)
                                    case 2147483647:
                                        if len(v14) > 0 {
                                            v14[0] = 32767
                                        }
                                    case 42:
                                        v1 += (v1 >> (((v1 % ((v1) | 1))) & 63))
                                    default:
                                        fmt.Println(fi19)
                                    }
                                    if ((float32((-2.25 + float64((0.1 + 0.0)))) <= (-2.25 + 0.1)) || false) {
                                        v3 += 18446744073709551615
                                        v17 += 18446744073709551615
                                        v139 := GenS18{F0: -2147483648, F1: -1}
                                        _ = v139
                                    } else {
                                        if (v2 >= fi138) {
                                            if (-59 == 42) {
                                                v140 := GenS12{F0: -497, F1: false}
                                                v2 = -660
                                                v141 := GenS21{F0: "", F1: true, F2: 0}
                                                _ = v140
                                                _ = v141
                                            }
                                        }
                                    }
                                    var v142 uint16 = 255
                                    _ = v142
                                } else {
                                    v17 += 42
                                    if len(v14) > 0 {
                                        v14[0] = (-(fi19))
                                    }
                                }
                                _ = fi138
                            }
                            v143 := GenS22{F0: 205, F1: 4294967295}
                            switch ((("foo bar" + "\n\t") + (("a" + "\n\t") + "foo bar")) + "hello") {
                            case "a":
                                v17 = (^(v1))
                            case "":
                                for fi144 := 0; fi144 < 15; fi144++ {
                                    var v145 float64 = (float64(0.0) - -2.25)
                                    var v146 uint64 = 2
                                    v147 := GenS23{F0: 68}
                                    switch "hello" {
                                    case "":
                                        var v148 int16 = 1
                                        _ = v148
                                    case "hello":
                                        var v149 int32 = -955
                                        _ = v149
                                    default:
                                        v150 := GenS24{F0: 589, F1: 0, F2: 327}
                                        _ = v150
                                    }
                                    _ = fi144
                                    _ = v145
                                    _ = v146
                                    _ = v147
                                }
                            }
                            _ = v143
                        } else {
                            v151 := GenS25{F0: v1, F1: 2147483647, F2: 4294967295}
                            _ = v151
                        }
                        if false {
                            var v152 int64 = -463
                            var v153 uint64 = 2
                            var v154 uint = v3
                            _ = v152
                            _ = v153
                            _ = v154
                        } else {
                            fmt.Println(v1, v17)
                            switch "\n\t" {
                            case "foo bar":
                                switch ("foo bar" + (((("" + "hello") + (("foo bar" + "a") + "\n\t")) + "hello") + "hello")) {
                                case "hello":
                                    fmt.Println(fi19, v17)
                                default:
                                    v3 += (^(v1))
                                }
                            case "hello":
                                if len(v14) > 0 {
                                    v14[0] = v2
                                }
                            case "a":
                                for fi155 := 0; fi155 < 19; fi155++ {
                                    v2 = 128
                                    v14 = append(v14, fi19)
                                    var v156 uint = (v3 << ((v17) & 63))
                                    _ = fi155
                                    _ = v156
                                }
                            default:
                                v1 = v1
                            }
                            var v157 int16 = 32767
                            _ = v157
                        }
                    } else {
                        v3 += (v3 ^ 1)
                        for fi158 := 0; fi158 < 18; fi158++ {
                            var v159 uint = v1
                            var v160 uint = v17
                            v161 := GenS26{F0: -641}
                            v159 = v17
                            _ = fi158
                            _ = v159
                            _ = v160
                            _ = v161
                        }
                        switch "\n\t" {
                        case "":
                            v17 = (v1 + v17)
                        case "a":
                            var v162 int16 = 1000
                            _ = v162
                        }
                    }
                    v3 -= v1
                    _ = fi19
                }
                _ = v17
            }
        }
        for fi163 := 0; fi163 < 7; fi163++ {
            var v164 uint16 = 255
            fmt.Println(v164)
            v14 = append(v14, (v2 - fi163))
            _ = fi163
            _ = v164
        }
        var v165 uint = (v3 / (((v1 + (v1 >> (((^(v3))) & 63)))) | 1))
        _ = v165
    } else {
        if true {
            v166 := GenS12{F0: -141, F1: ((560 != 255) && false)}
            v167 := GenS13{F0: 42, F1: -99}
            _ = v166
            _ = v167
        } else {
            for fi168 := 0; fi168 < 10; fi168++ {
                var v169 uint8 = 0
                if len(v14) > 0 {
                    v14[0] = 127
                }
                var v170 uint = v3
                _ = fi168
                _ = v169
                _ = v170
            }
            v171 := GenS27{F0: 0, F1: 73}
            if len(v14) > 0 {
                v14[0] = v2
            }
            _ = v171
        }
        for fi172 := 0; fi172 < 13; fi172++ {
            if false {
                for fi173 := 0; fi173 < 11; fi173++ {
                    v1 = v3
                    if false {
                        switch "\n\t" {
                        case "hello":
                            v174 := GenS28{F0: 42, F1: 109, F2: 127}
                            _ = v174
                        case "":
                            var v175 int32 = 1
                            _ = v175
                        case "\n\t":
                            if true {
                                v176 := GenS29{F0: 255, F1: 42}
                                v3 = (v3 % ((65535) | 1))
                                _ = v176
                            } else {
                                if false {
                                    v1 += (^(v3))
                                } else {
                                    var v177 int16 = 90
                                    v178 := GenS7{F0: (("\n\t" + "hello") + (("" + (("\n\t" + "") + "\n\t")) + "\n\t")), F1: 42}
                                    _ = v177
                                    _ = v178
                                }
                                if len(v14) > 0 {
                                    v14[0] = v2
                                }
                            }
                        default:
                            v2 = v2
                        }
                        if len(v14) > 0 {
                            v14[0] = fi173
                        }
                        switch v2 {
                        case -2147483648:
                            var v179 int64 = 1
                            _ = v179
                        case 127:
                            var v180 int64 = -110
                            _ = v180
                        case 338:
                            switch "hello" {
                            case "hello":
                                var v181 float64 = 3.14159
                                _ = v181
                            case "\n\t":
                                v182 := GenS30{F0: ("\n\t" + ""), F1: 1, F2: 4294967295}
                                _ = v182
                            default:
                                v3 -= v3
                            }
                        }
                    } else {
                        if true {
                            switch fi173 {
                            case -1:
                                v3 = (v3 % (((v3 ^ v1)) | 1))
                            case -9223372036854775808:
                                if true {
                                    fmt.Println(v3, fi172)
                                } else {
                                    var v183 int64 = -9223372036854775808
                                    var v184 uint8 = 178
                                    _ = v183
                                    _ = v184
                                }
                            }
                            fmt.Println(fi172, v3, fi172)
                        }
                        v3 -= v1
                        if len(v14) > 0 {
                            v14[0] = (v2 % ((v2) | 1))
                        }
                        switch "a" {
                        case "foo bar":
                            if true {
                                v185 := GenS24{F0: v3, F1: 1, F2: 2}
                                _ = v185
                            } else {
                                var v186 int8 = -1
                                _ = v186
                            }
                            switch ("foo bar" + "a") {
                            case "":
                                if true {
                                    fmt.Println(v1)
                                    fmt.Println(fi173, fi173, fi172)
                                    v2 -= (fi173 % ((128) | 1))
                                }
                            }
                            v2 -= (fi173 << ((255) & 63))
                            var v187 string = "foo bar"
                            _ = v187
                        }
                        v2 = (^(fi173))
                    }
                    _ = fi173
                }
            }
            _ = fi172
        }
    }
    v14 = append(v14, (v2 - (v2 * v2)))
    var v188 int16 = -1
    v189 := GenS31{F0: v188, F1: (!false)}
    var v190 uint = v1
    v188 = (v188 << ((2) & 63))
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    v3 = v190
    var v191 int8 = int8(v2)
    v192 := GenS20{F0: 255, F1: ""}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v188:", v188)
    fmt.Println("v190:", v190)
    fmt.Println("v191:", v191)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v188
    _ = v189
    _ = v190
    _ = v191
    _ = v192
    _ = v2
    _ = v3
    _ = v4
}
