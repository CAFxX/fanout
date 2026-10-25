package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 uint8; F2 int16 }
type GenS7 struct { F0 int8; F1 uint32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int32; F1 uint; F2 uint32 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int32; F1 uint8 }
type GenS13 struct { F0 uint32; F1 string; F2 int16 }
type GenS14 struct { F0 uint }
func genFn0(p0 uint32, p1 int64) uint32 {
    if (1 <= 128) {
        p1 += 9223372036854775807
        p1 = (p1 % ((p1) | 1))
        var v5 int = 127
        _ = v5
    }
    var v6 bool = true
    for fi7 := 0; fi7 < 9; fi7++ {
        if ((!false) && (499 <= 2147483647)) {
            v8 := []int{}
            var v9 int64 = (-(p1))
            _ = v8
            _ = v9
        } else {
            v10 := GenS1{F0: 15}
            var v11 int = 0
            _ = v10
            _ = v11
        }
        var v12 int16 = 1
        p0 = p0
        _ = fi7
        _ = v12
    }
    v13 := []int{}
    _ = p0
    _ = p1
    _ = v13
    _ = v6
    return p0
}
func genFn1(p0 uint16, p1 int64) int8 {
    for fi14 := 0; fi14 < 6; fi14++ {
        fmt.Println(p1)
        var v15 string = ""
        _ = fi14
        _ = v15
    }
    switch "hello" {
    case "hello":
        if (0.1 <= (0.1 * 0.1)) {
            var v16 int8 = 100
            v16 = v16
            p1 -= p1
            _ = v16
        }
    default:
        switch 128 {
        case 1:
            fmt.Println(p1, p0)
        case 9223372036854775807:
            fmt.Println(p1)
        case 398:
            var v17 int8 = 100
            _ = v17
        default:
            v18 := []int{127, 1}
            _ = v18
        }
    }
    p1 += 0
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 int16 = -959
    var v2 uint16 = 977
    var v3 int16 = v1
    var v4 float32 = ((0.0 * 1e10) * 3.14159)
    v19 := []int{127, 9223372036854775807, 128}
    if (0.1 > 0.1) {
        v19 = append(v19, -114)
        if false {
            var v20 uint8 = 2
            v2 = v2
            var v21 int64 = 0
            _ = v20
            _ = v21
        }
    } else {
        var v22 float32 = v4
        _ = v22
    }
    if (!false) {
        var v23 int16 = -32768
        for fi24 := 0; fi24 < 3; fi24++ {
            for fi25 := 0; fi25 < 14; fi25++ {
                v4 += 0.1
                var v26 int16 = (v1 - 1)
                v26 += v3
                _ = fi25
                _ = v26
            }
            _ = fi24
        }
        _ = v23
    } else {
        v27 := GenS1{F0: -28}
        v28 := GenS1{F0: 14}
        var v29 int8 = genFn1((v2 % ((886) | 1)), -1)
        _ = v27
        _ = v28
        _ = v29
    }
    v4 += v4
    if len(v19) > 0 {
        v19[0] = 42
    }
    fmt.Println(v3, v2, v2)
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    v4 -= (v4 + ((v4 * float32(float64(v4))) - (v4 / 1.5)))
    v4 = v4
    fmt.Println(v1)
    if (-1 < 9223372036854775807) {
        _ = len(v19)
        switch 127 {
        case 42:
            if ((!true) && (79 > 0)) {
                v4 = (1.5 / 0.1)
            }
        case -2147483648:
            for fi30 := 0; fi30 < 8; fi30++ {
                fmt.Println(v3)
                var v31 uint64 = 1
                v1 += v3
                _ = fi30
                _ = v31
            }
        case 32767:
            var v32 float64 = (((3.14159 - 0.0) / 0.1) / ((3.14159 + (((0.0 * ((float64(v4) * 0.1) / 1e10)) / 1.5) / (1.5 + 0.1))) / 0.1))
            _ = v32
        default:
            v4 -= v4
        }
        switch ((("a" + "hello") + ("foo bar" + (("foo bar" + "foo bar") + "hello"))) + ("\n\t" + "")) {
        case "":
            if false {
                var v33 float32 = (v4 - 0.0)
                v3 -= (-(v1))
                _ = v33
            } else {
                if (2 == 18446744073709551615) {
                    switch 2147483647 {
                    case -9223372036854775808:
                        v3 = v1
                    }
                } else {
                    for fi34 := 0; fi34 < 19; fi34++ {
                        for fi35 := 0; fi35 < 17; fi35++ {
                            v36 := GenS1{F0: 100}
                            _ = len(v19)
                            if false {
                                if true {
                                    v4 = v4
                                    fmt.Println(v2)
                                    var v37 uint16 = (v2 / (((^(v2))) | 1))
                                    if len(v19) > 0 {
                                        v19[0] = fi35
                                    }
                                    v38 := GenS1{F0: genFn1((v2 % ((42) | 1)), 42)}
                                    v19 = append(v19, 62)
                                    _ = v37
                                    _ = v38
                                }
                                v3 = -32768
                                v3 = (v1 + v1)
                            } else {
                                if false {
                                    v1 = v1
                                    var v39 uint32 = 869
                                    _ = v39
                                } else {
                                    v40 := GenS1{F0: -65}
                                    if (-407 <= 603) {
                                        v1 += (v1 >> ((65535) & 63))
                                        switch ("" + "foo bar") {
                                        case "hello":
                                            switch "a" {
                                            case "foo bar":
                                                var v41 bool = (-280 < 1)
                                                _ = v41
                                            case "hello":
                                                if true {
                                                    v2 += (v2 % ((0) | 1))
                                                    var v42 uint8 = 1
                                                    _ = v42
                                                }
                                            case "":
                                                var v43 float32 = v4
                                                _ = v43
                                            }
                                        case "foo bar":
                                            _ = len(v19)
                                        case "a":
                                            switch (fi34 * 551) {
                                            case 9223372036854775807:
                                                fmt.Println(fi34)
                                            case 2147483647:
                                                if false {
                                                    if len(v19) >= 2 {
                                                        v19 = v19[1:len(v19)]
                                                    }
                                                    fmt.Println(v2, fi35)
                                                }
                                            }
                                        }
                                        if (true || (1 >= 73)) {
                                            if true {
                                                var v44 float64 = -2.25
                                                _ = v44
                                            } else {
                                                if (858 >= 2) {
                                                    if true {
                                                        v3 -= 1000
                                                        v45 := GenS1{F0: genFn1(575, 42)}
                                                        if true {
                                                            var v46 int32 = 0
                                                            var v47 int64 = 1
                                                            _ = v46
                                                            _ = v47
                                                        }
                                                        _ = v45
                                                    }
                                                } else {
                                                    fmt.Println(v4, v2)
                                                    if (1 <= 2147483647) {
                                                        v48 := GenS1{F0: 85}
                                                        switch fi35 {
                                                        case 2147483647:
                                                            v3 = -234
                                                        case 1:
                                                            fmt.Println(fi35)
                                                        case -273:
                                                            if len(v19) >= 2 {
                                                                v19 = v19[1:len(v19)]
                                                            }
                                                        }
                                                        if (1 == 73) {
                                                            var v49 int8 = 20
                                                            var v50 float32 = (v4 - v4)
                                                            switch (fi34 % ((32767) | 1)) {
                                                            case 531:
                                                                switch fi35 {
                                                                case 128:
                                                                    var v51 int8 = v49
                                                                    _ = v51
                                                                case 535:
                                                                    v19 = append(v19, (fi34 | 2147483647))
                                                                case 2147483647:
                                                                    v3 = v3
                                                                }
                                                                var v52 int64 = 1
                                                                _ = v52
                                                            case -949:
                                                                if len(v19) >= 2 {
                                                                    v19 = v19[1:len(v19)]
                                                                }
                                                            }
                                                            if ((v50 >= v50) && (37 >= 667)) {
                                                                if (!true) {
                                                                    v53 := GenS2{F0: (fi35 ^ fi34), F1: genFn0(42, -872), F2: 720}
                                                                    if ((!(v49 <= 127)) && true) {
                                                                        v54 := GenS2{F0: (fi35 % (((fi35 / (((fi34 - fi34)) | 1))) | 1)), F1: 65535, F2: 4294967295}
                                                                        if (false || true) {
                                                                            v4 += (v4 + (v50 - (v50 + (v50 - v50))))
                                                                            v3 = v3
                                                                            if (2147483647 < -952) {
                                                                                if len(v19) > 0 {
                                                                                    v19[0] = fi34
                                                                                }
                                                                                fmt.Println(v50, v1)
                                                                                if (!(0 < (-(v3)))) {
                                                                                    v55 := GenS2{F0: -9223372036854775808, F1: 368, F2: 267}
                                                                                    v1 += v3
                                                                                    var v56 int64 = 9223372036854775807
                                                                                    _ = v55
                                                                                    _ = v56
                                                                                }
                                                                            } else {
                                                                                v57 := GenS1{F0: v49}
                                                                                _ = v57
                                                                            }
                                                                        }
                                                                        if false {
                                                                            v50 = 0.0
                                                                            var v58 uint16 = v2
                                                                            if len(v19) > 0 {
                                                                                v19[0] = 1
                                                                            }
                                                                            if false {
                                                                                if (true && (true || ((v2 ^ (^(v58))) == v58))) {
                                                                                    if (!(0 < -2147483648)) {
                                                                                        var v59 float64 = (-2.25 - -2.25)
                                                                                        var v60 uint = 255
                                                                                        var v61 bool = ((v1 == 1) || true)
                                                                                        v62 := GenS1{F0: genFn1(v2, 1)}
                                                                                        var v63 uint8 = 42
                                                                                        var v64 int = (fi34 - fi34)
                                                                                        if v61 {
                                                                                            v63 = v63
                                                                                            v19 = append(v19, fi34)
                                                                                            switch (fi35 - fi35) {
                                                                                            case 0:
                                                                                                switch v64 {
                                                                                                case 0:
                                                                                                    _ = len(v19)
                                                                                                case 266:
                                                                                                    v59 -= (float64(((((v4 + v4) - v4) - v50) + (1.5 - (v4 / v50)))) / float64(v50))
                                                                                                case -670:
                                                                                                    v65 := GenS1{F0: (v49 % ((v49) | 1))}
                                                                                                    _ = v65
                                                                                                }
                                                                                            case 4294967295:
                                                                                                var v66 string = "foo bar"
                                                                                                _ = v66
                                                                                            default:
                                                                                                var v67 int = (v64 + -959)
                                                                                                _ = v67
                                                                                            }
                                                                                        } else {
                                                                                            v68 := GenS3{F0: 1, F1: (v49 / (((^(v49))) | 1))}
                                                                                            if len(v19) > 0 {
                                                                                                v19[0] = (-(v64))
                                                                                            }
                                                                                            _ = v68
                                                                                        }
                                                                                        if (!true) {
                                                                                            v69 := GenS1{F0: genFn1(v58, -9223372036854775808)}
                                                                                            v60 -= (^(v60))
                                                                                            _ = v69
                                                                                        }
                                                                                        v19 = append(v19, 654)
                                                                                        _ = v59
                                                                                        _ = v60
                                                                                        _ = v61
                                                                                        _ = v62
                                                                                        _ = v63
                                                                                        _ = v64
                                                                                    } else {
                                                                                        v70 := GenS4{F0: -9223372036854775808}
                                                                                        if ((true || false) || (("" + ("" + "")) > "hello")) {
                                                                                            v1 = (v3 << ((1) & 63))
                                                                                        }
                                                                                        _ = len(v19)
                                                                                        _ = v70
                                                                                    }
                                                                                    if len(v19) > 0 {
                                                                                        v19[0] = 826
                                                                                    }
                                                                                } else {
                                                                                    _ = len(v19)
                                                                                }
                                                                                var v71 int64 = 0
                                                                                _ = v71
                                                                            }
                                                                            _ = v58
                                                                        }
                                                                        _ = v54
                                                                    } else {
                                                                        fmt.Println(v2, v1)
                                                                        v72 := GenS4{F0: 0}
                                                                        if (0 >= -215) {
                                                                            if true {
                                                                                if len(v19) >= 2 {
                                                                                    v19 = v19[1:len(v19)]
                                                                                }
                                                                                v1 += v3
                                                                            }
                                                                            switch "\n\t" {
                                                                            case "hello":
                                                                                v3 -= v1
                                                                            case "\n\t":
                                                                                if false {
                                                                                    var v73 int64 = -660
                                                                                    _ = v73
                                                                                }
                                                                            default:
                                                                                var v74 uint64 = 2
                                                                                _ = v74
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v1)
                                                                        }
                                                                        _ = v72
                                                                    }
                                                                    _ = len(v19)
                                                                    _ = v53
                                                                }
                                                                if (42 < genFn1(v2, -1)) {
                                                                    fmt.Println(v1, v4)
                                                                } else {
                                                                    var v75 string = "foo bar"
                                                                    v75 = v75
                                                                    v19 = append(v19, (fi35 & 2147483647))
                                                                    v75 = (v75 + v75)
                                                                    fmt.Println(v2, v4)
                                                                    _ = v75
                                                                }
                                                                fmt.Println(fi34, fi34, v49)
                                                            }
                                                            _ = v49
                                                            _ = v50
                                                        } else {
                                                            v2 += (v2 & v2)
                                                            var v76 int64 = -1
                                                            _ = v76
                                                        }
                                                        _ = v48
                                                    }
                                                    if (!(true || false)) {
                                                        if ((-945 > -1) || false) {
                                                            if len(v19) > 0 {
                                                                v19[0] = fi35
                                                            }
                                                            var v77 float32 = v4
                                                            _ = v77
                                                        } else {
                                                            fmt.Println(v4, v4)
                                                            v2 += 65535
                                                        }
                                                    }
                                                }
                                            }
                                            fmt.Println(v3, fi34)
                                        }
                                    } else {
                                        switch fi34 {
                                        case 0:
                                            switch fi35 {
                                            case 4294967295:
                                                switch "\n\t" {
                                                case "hello":
                                                    v3 = (v3 ^ (v3 | v3))
                                                }
                                            case -953:
                                                var v78 uint32 = 2147483648
                                                if true {
                                                    if len(v19) >= 2 {
                                                        v19 = v19[1:len(v19)]
                                                    }
                                                }
                                                switch fi34 {
                                                case -1:
                                                    fmt.Println(v78, v78)
                                                case 42:
                                                    var v79 uint64 = 1
                                                    _ = v79
                                                default:
                                                    switch (^(fi34)) {
                                                    case 42:
                                                        fmt.Println(v1, v2, v2)
                                                        var v80 int16 = v3
                                                        _ = v80
                                                    default:
                                                        v3 = v3
                                                    }
                                                }
                                                _ = v78
                                            case 42:
                                                if true {
                                                    if false {
                                                        if (2147483647 == -186) {
                                                            var v81 string = "a"
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                var v82 uint = 259
                                                                _ = v82
                                                            case "":
                                                                v2 = (v2 + v2)
                                                            case "\n\t":
                                                                fmt.Println(v1, v2)
                                                            default:
                                                                var v83 int32 = -1
                                                                _ = v83
                                                            }
                                                            _ = v81
                                                        }
                                                        if len(v19) > 0 {
                                                            v19[0] = fi35
                                                        }
                                                    } else {
                                                        switch (fi34 >> ((65535) & 63)) {
                                                        case 9223372036854775807:
                                                            v84 := GenS5{F0: 42}
                                                            _ = v84
                                                        case -801:
                                                            if false {
                                                                v1 = 42
                                                            }
                                                        case -2147483648:
                                                            fmt.Println(v2)
                                                        default:
                                                            var v85 bool = true
                                                            _ = v85
                                                        }
                                                    }
                                                } else {
                                                    var v86 string = "a"
                                                    var v87 uint64 = 342
                                                    if false {
                                                        if (4294967295 > 18446744073709551615) {
                                                            switch "" {
                                                            case "":
                                                                v86 = (v86 + v86)
                                                            case "\n\t":
                                                                fmt.Println(v4, v2)
                                                            }
                                                            v4 = v4
                                                            v2 -= (v2 + 255)
                                                            v88 := GenS1{F0: -112}
                                                            _ = v88
                                                        } else {
                                                            var v89 float64 = 1.5
                                                            var v90 bool = false
                                                            _ = v89
                                                            _ = v90
                                                        }
                                                        _ = len(v19)
                                                        v87 += v87
                                                    }
                                                    if (false || false) {
                                                        v91 := GenS4{F0: -1}
                                                        _ = v91
                                                    }
                                                    fmt.Println(fi35, v86, fi34)
                                                    _ = v86
                                                    _ = v87
                                                }
                                            }
                                        case -2147483648:
                                            switch fi35 {
                                            case 0:
                                                v2 += v2
                                            case -689:
                                                switch (^(fi35)) {
                                                case -2147483648:
                                                    switch "a" {
                                                    case "hello":
                                                        v4 -= v4
                                                    case "\n\t":
                                                        if len(v19) >= 2 {
                                                            v19 = v19[1:len(v19)]
                                                        }
                                                    case "a":
                                                        var v92 int8 = 127
                                                        _ = v92
                                                    }
                                                case 128:
                                                    fmt.Println(v1, v1)
                                                default:
                                                    if false {
                                                        fmt.Println(v4, v2)
                                                    } else {
                                                        if false {
                                                            fmt.Println(fi34, fi34, v3)
                                                        } else {
                                                            switch (-(fi34)) {
                                                            case 127:
                                                                v93 := GenS6{F0: 716, F1: 56, F2: v3}
                                                                _ = v93
                                                            case 9223372036854775807:
                                                                fmt.Println(v1)
                                                            case 2147483647:
                                                                switch -1 {
                                                                case 1:
                                                                    v94 := GenS7{F0: genFn1((v2 & (v2 - v2)), -9223372036854775808), F1: 2}
                                                                    v95 := GenS4{F0: -470}
                                                                    fmt.Println(fi34, fi34)
                                                                    switch fi34 {
                                                                    case 42:
                                                                        v19 = append(v19, (fi34 << ((42) & 63)))
                                                                    case -1:
                                                                        var v96 int8 = genFn1(v2, 9223372036854775807)
                                                                        _ = v96
                                                                    }
                                                                    _ = v94
                                                                    _ = v95
                                                                }
                                                            default:
                                                                _ = len(v19)
                                                            }
                                                        }
                                                        v1 = (v3 * v1)
                                                        v97 := GenS8{F0: 128}
                                                        if (255 == 150) {
                                                            if (128 < 230) {
                                                                switch (fi34 / ((4294967295) | 1)) {
                                                                case 32767:
                                                                    fmt.Println(v2, fi34)
                                                                case -2147483648:
                                                                    v3 -= v1
                                                                case 230:
                                                                    fmt.Println(v3, v4, fi34)
                                                                }
                                                            } else {
                                                                if false {
                                                                    v1 += v1
                                                                    v98 := GenS9{F0: 65535}
                                                                    if ("hello" <= "hello") {
                                                                        var v99 float32 = float32(0.0)
                                                                        v99 = 3.14159
                                                                        _ = v99
                                                                    }
                                                                    _ = v98
                                                                }
                                                                v4 = (v4 - 3.14159)
                                                            }
                                                            v100 := GenS8{F0: (fi34 * (fi35 << ((0) & 63)))}
                                                            _ = v100
                                                        } else {
                                                            v101 := GenS2{F0: 4294967295, F1: 690, F2: 103}
                                                            _ = v101
                                                        }
                                                        if true {
                                                            v1 -= (v1 * v1)
                                                            if true {
                                                                if len(v19) >= 2 {
                                                                    v19 = v19[1:len(v19)]
                                                                }
                                                            }
                                                            var v102 int64 = 19
                                                            v103 := GenS2{F0: fi35, F1: 2147483648, F2: 2}
                                                            _ = v102
                                                            _ = v103
                                                        }
                                                        fmt.Println(v4, v3)
                                                        if len(v19) >= 2 {
                                                            v19 = v19[1:len(v19)]
                                                        }
                                                        _ = v97
                                                    }
                                                }
                                            case 839:
                                                v3 = (v1 % (((v3 << ((1) & 63))) | 1))
                                                var v104 float32 = (v4 + 1e10)
                                                _ = v104
                                            default:
                                                switch 0 {
                                                case 128:
                                                    v4 -= ((v4 / v4) - 0.1)
                                                case -1:
                                                    _ = len(v19)
                                                    var v105 int = 32768
                                                    _ = v105
                                                }
                                            }
                                        case -1:
                                            v106 := GenS10{F0: 42, F1: 0, F2: 2147483647}
                                            _ = v106
                                        default:
                                            v4 -= (v4 - float32(-2.25))
                                        }
                                        fmt.Println(fi35, v3, fi34)
                                        var v107 int = (fi34 | fi35)
                                        _ = v107
                                    }
                                    _ = v40
                                }
                                fmt.Println(v4)
                                v1 -= v1
                            }
                            _ = fi35
                            _ = v36
                        }
                        if true {
                            v2 += (^(v2))
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                            var v108 float32 = (0.0 - (float32((1e10 - 0.0)) - (float32(float64(0.0)) / (3.14159 - 0.0))))
                            _ = v108
                        }
                        _ = fi34
                    }
                    for fi109 := 0; fi109 < 9; fi109++ {
                        if len(v19) > 0 {
                            v19[0] = 42
                        }
                        _ = fi109
                    }
                    if true {
                        fmt.Println(v3)
                        fmt.Println(v3)
                    } else {
                        v110 := GenS1{F0: 108}
                        _ = v110
                    }
                }
                v1 = v1
            }
        default:
            var v111 int32 = 42
            _ = v111
        }
    } else {
        switch 1 {
        case 9223372036854775807:
            _ = len(v19)
        case 602:
            if (504 >= -9223372036854775808) {
                if true {
                    v112 := GenS5{F0: -1}
                    v3 -= 0
                    _ = v112
                } else {
                    fmt.Println(v3)
                    v19 = append(v19, 0)
                    v2 = (v2 * (^(v2)))
                }
            }
        case -289:
            _ = len(v19)
        }
        if false {
            v113 := GenS11{F0: uint(v2)}
            _ = v113
        }
    }
    switch "" {
    case "foo bar":
        var v114 int8 = genFn1(v2, -1)
        _ = v114
    case "hello":
        if (0 > 65535) {
            var v115 int16 = (v3 * v1)
            _ = v115
        } else {
            v1 += v3
        }
    case "":
        switch 128 {
        case 4294967295:
            v116 := GenS11{F0: 768}
            _ = v116
        case 931:
            if (42 != 2147483647) {
                v19 = append(v19, 1)
                _ = len(v19)
                switch "foo bar" {
                case "hello":
                    v4 = v4
                case "\n\t":
                    v3 -= 42
                }
            } else {
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                v3 += 32767
            }
        case -65:
            if (true && (true && true)) {
                v3 += v3
                var v117 int32 = 1
                _ = v117
            } else {
                switch 4294967295 {
                case 128:
                    v1 += (v1 % ((v1) | 1))
                case -199:
                    var v118 int16 = 42
                    _ = v118
                case -2147483648:
                    for fi119 := 0; fi119 < 9; fi119++ {
                        var v120 uint16 = 1
                        for fi121 := 0; fi121 < 9; fi121++ {
                            if true {
                                v19 = append(v19, fi119)
                                var v122 bool = false
                                v123 := GenS9{F0: 971}
                                _ = v122
                                _ = v123
                            } else {
                                fmt.Println(fi119, v2)
                                var v124 int16 = v1
                                if false {
                                    switch fi121 {
                                    case 0:
                                        v125 := GenS12{F0: 0, F1: 255}
                                        _ = v125
                                    case -19:
                                        fmt.Println(v1, v4, v1)
                                    case 42:
                                        v120 -= v2
                                    }
                                    fmt.Println(v1, fi121)
                                    fmt.Println(v3, v1, fi121)
                                    var v126 string = ("" + "hello")
                                    fmt.Println(fi121, v120)
                                    _ = v126
                                }
                                _ = v124
                            }
                            var v127 float32 = (float32((3.14159 * (1.5 + 0.1))) / v4)
                            if false {
                                _ = len(v19)
                                if (42 < -1) {
                                    fmt.Println(fi119, v1, v127)
                                    var v128 int16 = v3
                                    _ = v128
                                }
                                v127 = v4
                            }
                            _ = fi121
                            _ = v127
                        }
                        _ = fi119
                        _ = v120
                    }
                default:
                    fmt.Println(v3, v3, v3)
                }
            }
        }
    default:
        var v129 float32 = (float32((1e10 / 1.0)) * v4)
        _ = v129
    }
    if false {
        fmt.Println(v4, v2)
    } else {
        var v130 uint64 = 1
        v131 := GenS8{F0: 9223372036854775807}
        _ = v130
        _ = v131
    }
    v132 := GenS7{F0: 127, F1: 680}
    for fi133 := 0; fi133 < 14; fi133++ {
        v2 = v2
        switch int(v2) {
        case 1:
            v1 += 793
        }
        _ = fi133
    }
    if ((v2 << ((42) & 63)) >= (v2 << ((18446744073709551615) & 63))) {
        for fi134 := 0; fi134 < 4; fi134++ {
            fmt.Println(v1, fi134)
            var v135 bool = true
            var v136 uint8 = 127
            _ = fi134
            _ = v135
            _ = v136
        }
        if (!true) {
            _ = len(v19)
            if (-52 != 0) {
                var v137 uint32 = 42
                _ = v137
            } else {
                var v138 uint64 = 733
                fmt.Println(v4, v4)
                _ = v138
            }
            if true {
                for fi139 := 0; fi139 < 9; fi139++ {
                    if ("\n\t" < "\n\t") {
                        v3 += 1
                        v2 -= (v2 / ((42) | 1))
                        v19 = append(v19, (^(fi139)))
                    }
                    _ = fi139
                }
            }
        } else {
            switch "a" {
            case "foo bar":
                v2 += (^(v2))
            case "":
                switch "a" {
                case "\n\t":
                    if true {
                        var v140 string = "hello"
                        _ = v140
                    }
                case "a":
                    if (!((v3 ^ (v3 % ((v3) | 1))) == v1)) {
                        switch ("hello" + "\n\t") {
                        case "":
                            for fi141 := 0; fi141 < 17; fi141++ {
                                v142 := GenS7{F0: -128, F1: 2147483647}
                                fmt.Println(v2, v3)
                                fmt.Println(v3, v1)
                                _ = fi141
                                _ = v142
                            }
                        case "\n\t":
                            v143 := GenS6{F0: 4294967295, F1: 1, F2: v3}
                            _ = v143
                        default:
                            for fi144 := 0; fi144 < 20; fi144++ {
                                var v145 string = ((("\n\t" + "\n\t") + (("foo bar" + "") + ("\n\t" + "\n\t"))) + ("a" + (("hello" + "hello") + "a")))
                                _ = fi144
                                _ = v145
                            }
                        }
                    } else {
                        for fi146 := 0; fi146 < 16; fi146++ {
                            for fi147 := 0; fi147 < 15; fi147++ {
                                v1 = v3
                                v148 := GenS13{F0: 4294967295, F1: "foo bar", F2: v1}
                                _ = fi147
                                _ = v148
                            }
                            if true {
                                if (true && ((3.14159 - (-2.25 / 1.0)) >= 0.1)) {
                                    switch 127 {
                                    case 0:
                                        fmt.Println(v3, v2, v1)
                                    case 2147483647:
                                        for fi149 := 0; fi149 < 20; fi149++ {
                                            fmt.Println(fi146, fi149, v3)
                                            _ = fi149
                                        }
                                    }
                                    switch "" {
                                    case "":
                                        var v150 string = (("hello" + ("a" + ("hello" + (("hello" + "a") + "foo bar")))) + "hello")
                                        _ = v150
                                    case "hello":
                                        if (!false) {
                                            v2 = 2
                                            if true {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    if (!(("\n\t" + "foo bar") > "")) {
                                                        fmt.Println(fi146)
                                                    } else {
                                                        var v151 int = (-(fi146))
                                                        var v152 int32 = -1
                                                        _ = v151
                                                        _ = v152
                                                    }
                                                case "\n\t":
                                                    v4 = v4
                                                }
                                            }
                                            for fi153 := 0; fi153 < 7; fi153++ {
                                                v154 := GenS12{F0: 0, F1: 127}
                                                _ = fi153
                                                _ = v154
                                            }
                                        }
                                    case "\n\t":
                                        v1 = v3
                                    default:
                                        v3 = (v1 ^ int16(fi146))
                                    }
                                    v1 = 439
                                }
                                v155 := GenS14{F0: 65535}
                                _ = v155
                            } else {
                                if true {
                                    v4 = v4
                                }
                            }
                            v4 = -2.25
                            _ = fi146
                        }
                        if len(v19) > 0 {
                            v19[0] = 0
                        }
                    }
                }
            case "\n\t":
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
            default:
                if (145 > -894) {
                    var v156 uint8 = 140
                    _ = v156
                } else {
                    var v157 uint64 = 0
                    var v158 int32 = -266
                    _ = v157
                    _ = v158
                }
            }
        }
    } else {
        for fi159 := 0; fi159 < 13; fi159++ {
            fmt.Println(v2, v1)
            v2 -= 2
            if (1 >= 2147483648) {
                v2 += (v2 ^ (v2 | (v2 - (v2 - 0))))
                fmt.Println(fi159, v2)
            } else {
                if ((fi159 >> ((18446744073709551615) & 63)) > (fi159 | fi159)) {
                    switch 32768 {
                    case 933:
                        v1 = -32768
                    case 9223372036854775807:
                        switch "a" {
                        case "":
                            if len(v19) > 0 {
                                v19[0] = (-(fi159))
                            }
                        case "hello":
                            v4 += -2.25
                        case "foo bar":
                            switch 2147483647 {
                            case 4294967295:
                                if true {
                                    v160 := GenS7{F0: -1, F1: 65535}
                                    for fi161 := 0; fi161 < 11; fi161++ {
                                        v162 := GenS3{F0: (fi161 >> ((18446744073709551615) & 63)), F1: 42}
                                        if (v1 <= -97) {
                                            v19 = append(v19, (fi161 % ((fi159) | 1)))
                                            fmt.Println(v3, v3, v4)
                                            fmt.Println(v4, v3)
                                        } else {
                                            _ = len(v19)
                                            v2 = v2
                                        }
                                        var v163 float32 = v4
                                        _ = fi161
                                        _ = v162
                                        _ = v163
                                    }
                                    _ = v160
                                } else {
                                    v2 = (v2 / ((v2) | 1))
                                    var v164 uint8 = 2
                                    v19 = append(v19, -689)
                                    _ = v164
                                }
                            case 0:
                                var v165 int32 = 0
                                _ = v165
                            case 128:
                                if ((1000 > 109) || (!((fi159 >> ((2) & 63)) < fi159))) {
                                    if len(v19) > 0 {
                                        v19[0] = (fi159 >> ((2) & 63))
                                    }
                                    for fi166 := 0; fi166 < 3; fi166++ {
                                        v1 -= (v1 / (((v1 << ((255) & 63))) | 1))
                                        var v167 int8 = genFn1(79, -1)
                                        _ = fi166
                                        _ = v167
                                    }
                                    v2 = 2
                                }
                            }
                        }
                    case -1:
                        var v168 uint32 = 258
                        _ = v168
                    default:
                        var v169 int64 = -724
                        _ = v169
                    }
                    v1 -= (v3 >> ((1) & 63))
                }
            }
            _ = fi159
        }
        var v170 float64 = 3.14159
        _ = v170
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v132
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
