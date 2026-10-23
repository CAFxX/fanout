package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int64; F1 int64; F2 uint64 }
type GenS4 struct { F0 int8; F1 int16; F2 uint16 }
type GenS5 struct { F0 string; F1 uint8; F2 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint8; F1 int16 }
type GenS8 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS9 struct { F0 uint8; F1 int32 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int64; F1 bool }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint64; F1 uint16 }
func genFn0() uint {
    if false {
        var v6 float64 = 0.0
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS1{F0: (false || (9223372036854775807 <= -1)), F1: 255}
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        v9 := []int{-859, 0, -2147483648, 127}
        if len(v9) > 0 {
            v9[0] = 32768
        }
        _ = v9
    }
    v10 := []int{128}
    if (1000 == 1) {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 bool = (1 <= 42)
            v13 := GenS2{F0: 18446744073709551615, F1: 101}
            _ = fi11
            _ = v12
            _ = v13
        }
    } else {
        if true {
            var v14 uint16 = 32767
            _ = len(v10)
            v15 := GenS1{F0: true, F1: 48}
            _ = v14
            _ = v15
        } else {
            for fi16 := 0; fi16 < 20; fi16++ {
                var v17 uint = 2
                v18 := GenS3{F0: 895, F1: 0, F2: 42}
                _ = fi16
                _ = v17
                _ = v18
            }
            for fi19 := 0; fi19 < 15; fi19++ {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                _ = fi19
            }
            var v20 uint16 = 2
            _ = v20
        }
        var v21 uint8 = 127
        _ = v21
    }
    switch 1 {
    case 127:
        var v22 float64 = (((((1.5 * 1.5) - -2.25) * (1e10 / float64(3.14159))) + -2.25) + 3.14159)
        _ = v22
    case 9223372036854775807:
        var v23 int8 = 99
        _ = v23
    default:
        for fi24 := 0; fi24 < 14; fi24++ {
            var v25 float64 = 0.0
            switch (fi24 / ((fi24) | 1)) {
            case 42:
                for fi26 := 0; fi26 < 17; fi26++ {
                    if false {
                        if (!(!true)) {
                            switch "\n\t" {
                            case "\n\t":
                                v27 := GenS4{F0: -93, F1: 1, F2: 486}
                                _ = v27
                            default:
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                            }
                        }
                    } else {
                        fmt.Println(fi26, fi24, fi24)
                        v28 := GenS5{F0: "\n\t", F1: 0, F2: (("hello" + "\n\t") + (("a" + "hello") + "foo bar"))}
                        v25 += float64(((3.14159 / (1.5 / 1.0)) - 0.0))
                        _ = v28
                    }
                    _ = fi26
                }
            case 9223372036854775807:
                v29 := GenS1{F0: (v25 <= v25), F1: 1}
                _ = v29
            default:
                fmt.Println(v25, v25)
            }
            if false {
                v25 = 1e10
                v30 := GenS1{F0: (2 > 1), F1: 2}
                var v31 float32 = float32(v25)
                _ = v30
                _ = v31
            }
            _ = fi24
            _ = v25
        }
    }
    _ = v10
    return 65535
}
func genFn1(p0 int8) int {
    for fi32 := 0; fi32 < 14; fi32++ {
        switch ("" + "hello") {
        case "":
            var v33 string = "\n\t"
            _ = v33
        case "foo bar":
            p0 -= (p0 | (p0 ^ p0))
        case "hello":
            var v34 int64 = 0
            _ = v34
        default:
            switch ("hello" + "\n\t") {
            case "\n\t":
                switch "a" {
                case "a":
                    v35 := GenS4{F0: 127, F1: 1000, F2: 255}
                    _ = v35
                case "hello":
                    if false {
                        var v36 float64 = ((1e10 - 0.0) * (0.1 * 1.5))
                        var v37 int32 = 1
                        _ = v36
                        _ = v37
                    } else {
                        fmt.Println(fi32)
                    }
                case "":
                    v38 := GenS4{F0: (p0 + -1), F1: 1, F2: 32767}
                    _ = v38
                default:
                    p0 = p0
                }
            case "hello":
                var v39 float32 = 3.14159
                _ = v39
            case "":
                p0 += p0
            }
        }
        _ = fi32
    }
    v40 := []int{474, 4294967295}
    if true {
        p0 += p0
        v41 := GenS6{F0: -604}
        _ = v41
    }
    _ = p0
    _ = v40
    return 32767
}

func main() {
    var v1 int64 = 1
    v2 := []uint32{4294967295, 2, 42}
    v3 := []int{9223372036854775807}
    v4 := make(map[uint32]string)
    v5 := make(map[int]bool)
    var v42 uint64 = 18446744073709551615
    for fi43 := 0; fi43 < 20; fi43++ {
        var v44 int32 = -719
        var v45 int16 = 32767
        _ = fi43
        _ = v44
        _ = v45
    }
    if (0 < 42) {
        for fi46 := 0; fi46 < 1; fi46++ {
            v1 -= (^(v1))
            if true {
                var v47 int16 = 1
                switch fi46 {
                case 42:
                    delete(v4, 468)
                case 0:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    switch genFn1(-1) {
                    case -2147483648:
                        var v48 uint16 = 710
                        _ = v48
                    }
                }
                _ = v47
            } else {
                for fi49 := 0; fi49 < 17; fi49++ {
                    v50 := GenS6{F0: -217}
                    if (true && false) {
                        v4[932] = ""
                        switch "hello" {
                        case "a":
                            v2 = append(v2, 65535)
                        case "":
                            if ((v42 << ((99) & 63)) > v42) {
                                if (!("" > "hello")) {
                                    var v51 int16 = -1
                                    _ = v51
                                }
                                var v52 int = (fi49 - fi49)
                                if (42 >= 255) {
                                    if false {
                                        v53 := GenS7{F0: 128, F1: -620}
                                        _ = v53
                                    }
                                    v42 = (v42 - (^(v42)))
                                } else {
                                    v54 := GenS8{F0: 128, F1: 396, F2: 40}
                                    _ = v54
                                }
                                _ = v52
                            } else {
                                v3 = append(v3, 4294967295)
                            }
                        case "hello":
                            v55 := v4[71]
                            _ = v55
                        }
                        v56 := GenS9{F0: 0, F1: 1}
                        v1 = -9223372036854775808
                        v42 += v42
                        _ = v56
                    } else {
                        v1 = (v1 * (v1 / (((v1 / ((-323) | 1))) | 1)))
                    }
                    _ = fi49
                    _ = v50
                }
                switch ("hello" + "foo bar") {
                case "\n\t":
                    fmt.Println(v42, v42)
                case "a":
                    if (-257 != -32768) {
                        for fi57 := 0; fi57 < 19; fi57++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi57
                        }
                        v58 := GenS7{F0: 255, F1: -32768}
                        if ((false || (2 == 255)) || false) {
                            var v59 int64 = 1
                            var v60 float64 = 0.0
                            fmt.Println(v42, fi46)
                            _ = v59
                            _ = v60
                        } else {
                            for fi61 := 0; fi61 < 16; fi61++ {
                                switch fi46 {
                                case 42:
                                    v62 := GenS9{F0: 42, F1: 2147483647}
                                    v2 = append(v2, 4294967295)
                                    _ = len(v5)
                                    switch ("hello" + "hello") {
                                    case "\n\t":
                                        if (2 == 128) {
                                            if (!(-9223372036854775808 > (v1 / (((^(v1))) | 1)))) {
                                                if true {
                                                    fmt.Println(fi61, v1)
                                                }
                                            }
                                        } else {
                                            var v63 int16 = 0
                                            v64 := v5[fi46]
                                            v63 = (v63 | (^(v63)))
                                            if (2147483647 != 42) {
                                                if (127 > 42) {
                                                    v65 := v5[2147483647]
                                                    v63 = (-(v63))
                                                    if (!v64) {
                                                        switch (("a" + ("hello" + "")) + (("\n\t" + "foo bar") + (("hello" + "") + ""))) {
                                                        case "a":
                                                            v1 = int64(v63)
                                                        default:
                                                            v1 += v1
                                                        }
                                                        fmt.Println(v63, v64)
                                                        var v66 uint8 = 2
                                                        _ = v66
                                                    } else {
                                                        var v67 uint64 = 0
                                                        switch (fi46 % ((fi61) | 1)) {
                                                        case -746:
                                                            switch (^(fi61)) {
                                                            case 9223372036854775807:
                                                                v68 := GenS10{F0: -305}
                                                                _ = v68
                                                            }
                                                        case 616:
                                                            if true {
                                                                var v69 uint = 829
                                                                var v70 uint16 = 735
                                                                _ = v69
                                                                _ = v70
                                                            }
                                                        case 9223372036854775807:
                                                            switch "hello" {
                                                            case "\n\t":
                                                                var v71 uint64 = (v42 & (v67 | 366))
                                                                _ = v71
                                                            case "hello":
                                                                var v72 float64 = (3.14159 * 1.5)
                                                                _ = v72
                                                            default:
                                                                var v73 int8 = 100
                                                                _ = v73
                                                            }
                                                        default:
                                                            v3 = append(v3, genFn1(61))
                                                        }
                                                        _ = v67
                                                    }
                                                    switch (-(fi61)) {
                                                    case -307:
                                                        switch fi61 {
                                                        case 4294967295:
                                                            var v74 uint64 = (^(v42))
                                                            _ = v74
                                                        }
                                                    case 2147483647:
                                                        var v75 int = 32767
                                                        _ = v75
                                                    case 32768:
                                                        var v76 uint = 2
                                                        if true {
                                                            v5[fi46] = (-520 >= -407)
                                                        }
                                                        v65 = (false || v64)
                                                        _ = v76
                                                    default:
                                                        v64 = (!(42 > v42))
                                                    }
                                                    v65 = v64
                                                    fmt.Println(v63, v42, v42)
                                                    _ = len(v4)
                                                    if (!(-208 < 2147483647)) {
                                                        v63 -= -1
                                                        v77 := GenS3{F0: 121, F1: v1, F2: (v42 * v42)}
                                                        _ = v77
                                                    } else {
                                                        v78 := GenS8{F0: 1, F1: uint16(v42), F2: 2147483647}
                                                        v79 := GenS8{F0: 2, F1: 42, F2: 2147483647}
                                                        fmt.Println(v65, fi61)
                                                        _ = v78
                                                        _ = v79
                                                    }
                                                    _ = v65
                                                } else {
                                                    switch 1 {
                                                    case -236:
                                                        var v80 int8 = -128
                                                        _ = v80
                                                    }
                                                    v81 := GenS11{F0: (-(v1)), F1: v64}
                                                    _ = v81
                                                }
                                            } else {
                                                if (!((!(!v64)) || v64)) {
                                                    var v82 float32 = (0.0 - 1.5)
                                                    v42 = 4294967295
                                                    _ = v82
                                                } else {
                                                    switch (fi61 & fi61) {
                                                    case -9223372036854775808:
                                                        v42 += 374
                                                    case -1:
                                                        var v83 uint64 = (^(v42))
                                                        _ = v83
                                                    case 1:
                                                        switch ("hello" + "foo bar") {
                                                        case "foo bar":
                                                            v64 = (0.0 >= float64((3.14159 - 0.1)))
                                                        case "hello":
                                                            v64 = v64
                                                        }
                                                    default:
                                                        v84 := GenS6{F0: (v63 ^ (v63 & v63))}
                                                        _ = v84
                                                    }
                                                    v42 = v42
                                                    if (-2147483648 != 679) {
                                                        v64 = (v64 && (v64 && v64))
                                                        var v85 float64 = -2.25
                                                        _ = v85
                                                    }
                                                }
                                                var v86 int8 = 41
                                                _ = v86
                                            }
                                            _ = v63
                                            _ = v64
                                        }
                                    }
                                    _ = v62
                                case 4294967295:
                                    var v87 int32 = 69
                                    _ = v87
                                case -317:
                                    if true {
                                        var v88 bool = false
                                        _ = v88
                                    }
                                }
                                if len(v2) > 0 {
                                    v2[0] = 401
                                }
                                _ = fi61
                            }
                            if true {
                                switch (^(fi46)) {
                                case 10:
                                    var v89 float64 = (3.14159 - -2.25)
                                    _ = v89
                                case 0:
                                    v90 := GenS12{F0: "foo bar"}
                                    _ = v90
                                }
                                if false {
                                    var v91 uint8 = 42
                                    v92 := GenS13{F0: (v42 % (((v42 << ((uint(v91)) & 63))) | 1)), F1: 65535}
                                    _ = v91
                                    _ = v92
                                } else {
                                    var v93 uint16 = 42
                                    _ = v93
                                }
                                fmt.Println(fi46, v42)
                            } else {
                                if false {
                                    for fi94 := 0; fi94 < 11; fi94++ {
                                        switch (-(fi94)) {
                                        case 9223372036854775807:
                                            if false {
                                                v42 += (v42 >> ((18446744073709551615) & 63))
                                            }
                                        }
                                        v42 = v42
                                        _ = fi94
                                    }
                                    var v95 uint32 = uint32(v42)
                                    switch (fi46 << ((653) & 63)) {
                                    case -190:
                                        fmt.Println(v42)
                                    case 42:
                                        v96 := v5[fi46]
                                        _ = v96
                                    }
                                    _ = v95
                                }
                            }
                            fmt.Println(v42, v42, v1)
                        }
                        _ = v58
                    }
                case "foo bar":
                    if false {
                        var v97 string = ("" + "hello")
                        v97 = "hello"
                        _ = v97
                    }
                default:
                    switch fi46 {
                    case 32768:
                        v42 -= (v42 % (((v42 & 2)) | 1))
                    case -1:
                        var v98 int32 = 199
                        _ = v98
                    default:
                        fmt.Println(v42)
                    }
                }
                fmt.Println(v1, fi46)
            }
            _ = fi46
        }
    }
    v1 = 1
    v99 := GenS6{F0: -1}
    for fi100 := 0; fi100 < 19; fi100++ {
        for fi101 := 0; fi101 < 6; fi101++ {
            v102 := GenS9{F0: 128, F1: -909}
            v42 = 4294967295
            switch (^(fi100)) {
            case 128:
                v4[42] = (("" + ("hello" + (("\n\t" + "hello") + "\n\t"))) + "hello")
            case 32768:
                var v103 int = (fi101 / (((fi100 >> ((143) & 63))) | 1))
                _ = v103
            case 0:
                var v104 int8 = -101
                _ = v104
            default:
                var v105 float64 = 1e10
                _ = v105
            }
            var v106 int = (-(fi100))
            _ = fi101
            _ = v102
            _ = v106
        }
        var v107 uint32 = 4294967295
        fmt.Println(v107, v1, v42)
        _ = fi100
        _ = v107
    }
    v42 += (v42 >> ((4294967295) & 63))
    v1 = v1
    if len(v3) > 0 {
        v3[0] = genFn1(120)
    }
    for fi108 := 0; fi108 < 1; fi108++ {
        switch fi108 {
        case 1:
            var v109 int32 = 2147483647
            _ = v109
        }
        _ = fi108
    }
    v3 = append(v3, 537)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v99
}
