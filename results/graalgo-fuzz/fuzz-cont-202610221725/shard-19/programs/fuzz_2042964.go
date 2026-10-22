package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 int64; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 uint; F1 int16; F2 uint64 }
type GenS6 struct { F0 int8; F1 uint16; F2 string }
type GenS7 struct { F0 bool; F1 int; F2 uint }
type GenS8 struct { F0 uint32; F1 int8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int; F1 int16; F2 int32 }
type GenS11 struct { F0 uint; F1 string }
type GenS12 struct { F0 uint16; F1 int8; F2 int64 }
type GenS13 struct { F0 uint64; F1 uint64 }
type GenS14 struct { F0 uint16; F1 string }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 uint32; F1 bool; F2 int }
func genFn0(p0 int32, p1 uint, p2 uint8) int {
    var v4 uint32 = 65535
    v5 := GenS1{F0: 118}
    if true {
        var v6 int64 = -9223372036854775808
        _ = v6
    } else {
        for fi7 := 0; fi7 < 9; fi7++ {
            switch "foo bar" {
            case "a":
                for fi8 := 0; fi8 < 13; fi8++ {
                    v9 := GenS1{F0: 42}
                    var v10 int8 = 0
                    if true {
                        if (true && true) {
                            switch (fi8 / (((^(fi7))) | 1)) {
                            case 369:
                                p2 = 1
                            case 4294967295:
                                var v11 uint32 = 2147483648
                                v10 = 42
                                _ = v11
                            case 431:
                                var v12 int32 = (p0 + 2147483647)
                                _ = v12
                            default:
                                v13 := GenS2{F0: 990, F1: 42, F2: (v4 ^ 2147483647)}
                                var v14 int32 = 0
                                _ = v13
                                _ = v14
                            }
                        }
                    } else {
                        var v15 uint = (^(p1))
                        v10 -= int8(fi8)
                        _ = v15
                    }
                    p1 -= p1
                    v16 := GenS1{F0: 0}
                    _ = fi8
                    _ = v10
                    _ = v16
                    _ = v9
                }
            case "foo bar":
                if true {
                    if true {
                        if false {
                            var v17 float64 = 1.5
                            if false {
                                v18 := GenS1{F0: 619}
                                var v19 uint16 = 631
                                if false {
                                    switch fi7 {
                                    case 127:
                                        if (p1 != 255) {
                                            v20 := GenS3{F0: 1, F1: 1}
                                            p0 = p0
                                            for fi21 := 0; fi21 < 5; fi21++ {
                                                var v22 float64 = float64((1e10 + (((-2.25 / 3.14159) / 0.1) / 3.14159)))
                                                if false {
                                                    var v23 uint = (p1 * p1)
                                                    v4 = (v4 + 2147483647)
                                                    var v24 uint32 = 885
                                                    _ = v23
                                                    _ = v24
                                                }
                                                _ = fi21
                                                _ = v22
                                            }
                                            _ = v20
                                        } else {
                                            var v25 uint32 = 2147483648
                                            if (((!false) || true) && true) {
                                                var v26 bool = ("foo bar" < "\n\t")
                                                v19 -= 42
                                                v27 := GenS4{F0: (-(fi7)), F1: p2}
                                                _ = v26
                                                _ = v27
                                            } else {
                                                var v28 uint8 = 1
                                                fmt.Println(p0, p0, p2)
                                                _ = v28
                                            }
                                            _ = v25
                                        }
                                    case -344:
                                        v19 = (v19 / (((v19 & (^(v19)))) | 1))
                                    case 42:
                                        if false {
                                            switch ("" + ("a" + "hello")) {
                                            case "hello":
                                                v29 := []int{int(p1)}
                                                _ = v29
                                            case "foo bar":
                                                for fi30 := 0; fi30 < 14; fi30++ {
                                                    v31 := []int{fi7, (fi7 << ((0) & 63)), fi30}
                                                    _ = fi30
                                                    _ = v31
                                                }
                                            case "a":
                                                for fi32 := 0; fi32 < 5; fi32++ {
                                                    var v33 int64 = 9223372036854775807
                                                    switch (fi7 - (fi32 - fi32)) {
                                                    case 0:
                                                        var v34 uint64 = 4294967295
                                                        _ = v34
                                                    case -676:
                                                        switch "" {
                                                        case "a":
                                                            fmt.Println(v4, v4)
                                                        case "foo bar":
                                                            v35 := []int{fi32}
                                                            _ = v35
                                                        }
                                                    case -1:
                                                        fmt.Println(p1, v17)
                                                    }
                                                    v36 := GenS1{F0: uint64(p1)}
                                                    _ = fi32
                                                    _ = v33
                                                    _ = v36
                                                }
                                            }
                                        }
                                    default:
                                        p1 -= p1
                                    }
                                } else {
                                    var v37 int32 = p0
                                    _ = v37
                                }
                                _ = v18
                                _ = v19
                            } else {
                                p2 = (p2 % (((p2 << ((p1) & 63))) | 1))
                            }
                            _ = v17
                        } else {
                            p2 += 42
                            if true {
                                p1 += 255
                                var v38 uint8 = p2
                                _ = v38
                            } else {
                                v39 := []int{fi7, fi7, (fi7 | 128)}
                                v4 = 222
                                var v40 uint32 = (v4 ^ 28)
                                _ = v39
                                _ = v40
                            }
                            v41 := []int{-9223372036854775808, (fi7 | (-(fi7))), fi7}
                            _ = v41
                        }
                        if false {
                            var v42 uint32 = 527
                            v43 := GenS5{F0: p1, F1: -763, F2: 467}
                            _ = v42
                            _ = v43
                        }
                        for fi44 := 0; fi44 < 1; fi44++ {
                            switch (fi44 + fi7) {
                            case 36:
                                if false {
                                    var v45 float32 = ((-2.25 * 3.14159) - 1e10)
                                    p0 = -1
                                    _ = v45
                                }
                                var v46 int64 = 1
                                _ = v46
                            case 4294967295:
                                fmt.Println(fi44)
                            }
                            v47 := GenS1{F0: 922}
                            fmt.Println(p2)
                            v48 := []int{(fi7 * fi7), (fi44 & (fi7 >> ((p1) & 63))), fi44}
                            _ = fi44
                            _ = v47
                            _ = v48
                        }
                    }
                    switch (((("a" + ("\n\t" + "\n\t")) + ((((("hello" + ("" + ((("foo bar" + ("\n\t" + "hello")) + "hello") + ((("" + (((((((((("" + ("\n\t" + "hello")) + "a") + ("foo bar" + "")) + ("hello" + "hello")) + ("hello" + "\n\t")) + "") + "hello") + (("\n\t" + "foo bar") + (("hello" + (("\n\t" + "") + "hello")) + ("a" + "")))) + "\n\t") + "")) + "\n\t") + "\n\t")))) + "hello") + "a") + "a") + "foo bar")) + "hello") + "") {
                    case "\n\t":
                        v4 += (v4 & v4)
                    case "foo bar":
                        if (int(p1) != 2147483647) {
                            var v49 uint16 = uint16(v4)
                            _ = v49
                        } else {
                            for fi50 := 0; fi50 < 4; fi50++ {
                                if false {
                                    var v51 uint8 = p2
                                    var v52 uint16 = 996
                                    if false {
                                        fmt.Println(p2)
                                        v53 := GenS6{F0: -128, F1: 1, F2: "hello"}
                                        v54 := []int{(fi50 * 32768), 4294967295, (^(fi7))}
                                        _ = v53
                                        _ = v54
                                    } else {
                                        fmt.Println(p0, v51, v51)
                                    }
                                    v51 -= 128
                                    v55 := []int{int(v52)}
                                    if (v52 != v52) {
                                        var v56 int32 = p0
                                        if false {
                                            if false {
                                                if false {
                                                    v51 = uint8(p1)
                                                    p2 = 42
                                                    switch (fi7 & (fi7 | 1)) {
                                                    case 600:
                                                        switch ("a" + "") {
                                                        case "":
                                                            var v57 uint64 = 967
                                                            v58 := GenS2{F0: p1, F1: 42, F2: (v4 / (((v4 >> ((2) & 63))) | 1))}
                                                            var v59 bool = true
                                                            _ = v57
                                                            _ = v58
                                                            _ = v59
                                                        case "a":
                                                            switch fi50 {
                                                            case 1:
                                                                switch 128 {
                                                                case 128:
                                                                    v4 = (v4 + uint32(v52))
                                                                default:
                                                                    switch fi7 {
                                                                    case -9223372036854775808:
                                                                        switch 4294967295 {
                                                                        case 142:
                                                                            var v60 float64 = (((0.1 / 1.0) * ((3.14159 - 1.5) * 3.14159)) + 0.1)
                                                                            _ = v60
                                                                        default:
                                                                            fmt.Println(p2)
                                                                        }
                                                                    case 2147483647:
                                                                        if len(v55) > 0 {
                                                                            v55[0] = (-(fi50))
                                                                        }
                                                                    case 595:
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            switch -9223372036854775808 {
                                                                            case 128:
                                                                                v51 += v51
                                                                            case 2147483647:
                                                                                var v61 float64 = 1e10
                                                                                _ = v61
                                                                            default:
                                                                                switch "" {
                                                                                case "":
                                                                                    if true {
                                                                                        var v62 float64 = (0.0 + float64((((1.5 - ((-2.25 - -2.25) + -2.25)) + -2.25) / float32(1e10))))
                                                                                        _ = v62
                                                                                    }
                                                                                }
                                                                            }
                                                                        case "a":
                                                                            p1 = (p1 % ((p1) | 1))
                                                                        }
                                                                    }
                                                                }
                                                            case 42:
                                                                if (!(!(("hello" >= "hello") && (1.5 != 3.14159)))) {
                                                                    if len(v55) > 0 {
                                                                        v55[0] = fi7
                                                                    }
                                                                    p1 = (^(p1))
                                                                    if true {
                                                                        var v63 bool = true
                                                                        if v63 {
                                                                            v63 = (!false)
                                                                            if v63 {
                                                                                var v64 uint = 603
                                                                                var v65 int = fi50
                                                                                _ = v64
                                                                                _ = v65
                                                                            } else {
                                                                                p1 = 42
                                                                                v66 := GenS7{F0: (272 == 42), F1: 2147483647, F2: uint(v52)}
                                                                                switch fi7 {
                                                                                case 1:
                                                                                    var v67 uint8 = p2
                                                                                    _ = v67
                                                                                case 42:
                                                                                    if len(v55) > 0 {
                                                                                        v55[0] = fi7
                                                                                    }
                                                                                case -963:
                                                                                    v56 = int32(fi50)
                                                                                    fmt.Println(fi50, v51, p1)
                                                                                default:
                                                                                    if v63 {
                                                                                        _ = len(v55)
                                                                                        v4 -= (v4 + (^(v4)))
                                                                                        var v68 float32 = 0.0
                                                                                        _ = v68
                                                                                    }
                                                                                }
                                                                                _ = v66
                                                                            }
                                                                            v52 = (v52 * v52)
                                                                            p2 -= p2
                                                                        } else {
                                                                            if false {
                                                                                var v69 int = fi7
                                                                                _ = v69
                                                                            } else {
                                                                                if true {
                                                                                    fmt.Println(v52, p1)
                                                                                    if (v52 >= 0) {
                                                                                        v63 = false
                                                                                        v52 = (v52 ^ 42)
                                                                                    } else {
                                                                                        var v70 int16 = -534
                                                                                        if len(v55) > 0 {
                                                                                            v55[0] = (fi50 / ((fi50) | 1))
                                                                                        }
                                                                                        _ = v70
                                                                                    }
                                                                                } else {
                                                                                    if len(v55) >= 2 {
                                                                                        v55 = v55[1:len(v55)]
                                                                                    }
                                                                                    if len(v55) > 0 {
                                                                                        v55[0] = 32767
                                                                                    }
                                                                                    v71 := GenS8{F0: v4, F1: -128}
                                                                                    v52 = (v52 % ((v52) | 1))
                                                                                    _ = v71
                                                                                }
                                                                                v55 = append(v55, (fi50 * fi50))
                                                                                switch ("a" + ("\n\t" + "a")) {
                                                                                case "a":
                                                                                    _ = len(v55)
                                                                                case "hello":
                                                                                    if (2 <= v51) {
                                                                                        if v63 {
                                                                                            p0 = v56
                                                                                            if len(v55) > 0 {
                                                                                                v55[0] = fi50
                                                                                            }
                                                                                            v51 += (v51 ^ (v51 % ((255) | 1)))
                                                                                            _ = len(v55)
                                                                                        }
                                                                                        var v72 uint64 = 18446744073709551615
                                                                                        _ = v72
                                                                                    }
                                                                                    p2 += 1
                                                                                case "foo bar":
                                                                                    v55 = append(v55, (fi7 * 0))
                                                                                    v56 += v56
                                                                                    var v73 uint8 = p2
                                                                                    switch fi50 {
                                                                                    case 460:
                                                                                        if v63 {
                                                                                            if (2 <= 2) {
                                                                                                fmt.Println(fi7, v4)
                                                                                                v74 := GenS9{F0: 929}
                                                                                                _ = v74
                                                                                            }
                                                                                        } else {
                                                                                            if v63 {
                                                                                                var v75 float32 = (0.0 / 1e10)
                                                                                                v76 := GenS7{F0: v63, F1: 42, F2: p1}
                                                                                                v73 += (v73 % ((v51) | 1))
                                                                                                _ = v75
                                                                                                _ = v76
                                                                                            }
                                                                                            switch (fi50 & fi50) {
                                                                                            case -220:
                                                                                                var v77 uint64 = 42
                                                                                                _ = v77
                                                                                            default:
                                                                                                v78 := GenS9{F0: (-(fi7))}
                                                                                                var v79 string = "\n\t"
                                                                                                _ = v78
                                                                                                _ = v79
                                                                                            }
                                                                                            if len(v55) >= 2 {
                                                                                                v55 = v55[1:len(v55)]
                                                                                            }
                                                                                        }
                                                                                    case -832:
                                                                                        p1 += (p1 / (((p1 << ((1) & 63))) | 1))
                                                                                        if (v4 > (v4 - v4)) {
                                                                                            if len(v55) >= 2 {
                                                                                                v55 = v55[1:len(v55)]
                                                                                            }
                                                                                            v80 := GenS4{F0: (fi50 >> ((541) & 63)), F1: v51}
                                                                                            v63 = false
                                                                                            fmt.Println(p0, v73, v4)
                                                                                            p2 = v51
                                                                                            _ = v80
                                                                                        } else {
                                                                                            switch fi50 {
                                                                                            case -9223372036854775808:
                                                                                                v81 := GenS2{F0: p1, F1: 0, F2: uint32(v73)}
                                                                                                _ = v81
                                                                                            default:
                                                                                                _ = len(v55)
                                                                                            }
                                                                                            v52 = (^(v52))
                                                                                        }
                                                                                    default:
                                                                                        p2 = v73
                                                                                    }
                                                                                    _ = v73
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v63
                                                                    }
                                                                    var v82 uint8 = p2
                                                                    _ = v82
                                                                } else {
                                                                    var v83 uint8 = p2
                                                                    _ = v83
                                                                }
                                                            }
                                                        case "\n\t":
                                                            fmt.Println(p1, v56, p1)
                                                            var v84 int8 = 1
                                                            var v85 string = ("foo bar" + "")
                                                            _ = v84
                                                            _ = v85
                                                        default:
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                if false {
                                                                    fmt.Println(p1)
                                                                    v86 := GenS10{F0: fi7, F1: 42, F2: p0}
                                                                    _ = v86
                                                                }
                                                            }
                                                        }
                                                    case 455:
                                                        if ((289 <= -667) && false) {
                                                            switch 334 {
                                                            case -9223372036854775808:
                                                                if true {
                                                                    v55 = append(v55, (fi7 >> (((p1 - 956)) & 63)))
                                                                    p1 = (p1 << (((p1 | p1)) & 63))
                                                                }
                                                            case 9223372036854775807:
                                                                if len(v55) > 0 {
                                                                    v55[0] = 9223372036854775807
                                                                }
                                                            case 179:
                                                                v87 := GenS11{F0: p1, F1: "\n\t"}
                                                                _ = v87
                                                            default:
                                                                v88 := GenS5{F0: uint(v51), F1: 0, F2: 1}
                                                                _ = v88
                                                            }
                                                            if (307 > 285) {
                                                                var v89 int64 = -1
                                                                var v90 string = ""
                                                                if false {
                                                                    p0 -= (p0 / ((p0) | 1))
                                                                } else {
                                                                    v91 := GenS7{F0: true, F1: -358, F2: 65535}
                                                                    v90 = ("a" + ("foo bar" + (("foo bar" + "hello") + "hello")))
                                                                    var v92 int16 = 39
                                                                    _ = len(v55)
                                                                    _ = v91
                                                                    _ = v92
                                                                }
                                                                _ = v89
                                                                _ = v90
                                                            } else {
                                                                fmt.Println(fi50, p2)
                                                            }
                                                            v93 := GenS2{F0: p1, F1: 1, F2: (v4 & (^(v4)))}
                                                            _ = v93
                                                        } else {
                                                            switch "a" {
                                                            case "a":
                                                                v51 = (^(v51))
                                                            case "\n\t":
                                                                switch (-(fi50)) {
                                                                case -900:
                                                                    var v94 uint = (p1 * uint(v52))
                                                                    _ = v94
                                                                case 4294967295:
                                                                    v51 -= 1
                                                                default:
                                                                    if false {
                                                                        v95 := GenS1{F0: 0}
                                                                        _ = v95
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    v55 = append(v55, fi7)
                                                }
                                                var v96 float64 = 3.14159
                                                _ = v96
                                            } else {
                                                if true {
                                                    if ((("a" + ("hello" + (("foo bar" + ("\n\t" + "foo bar")) + "a"))) < "") || (!(!false))) {
                                                        var v97 string = ("foo bar" + "foo bar")
                                                        v51 += (p2 | v51)
                                                        _ = v97
                                                    }
                                                    var v98 int32 = p0
                                                    if false {
                                                        if true {
                                                            var v99 int32 = (p0 | v56)
                                                            _ = v99
                                                        }
                                                        fmt.Println(v98, p0, p1)
                                                        var v100 int8 = -128
                                                        _ = v100
                                                    } else {
                                                        v51 -= (p2 & uint8(v52))
                                                        v98 += v98
                                                    }
                                                    if len(v55) >= 2 {
                                                        v55 = v55[1:len(v55)]
                                                    }
                                                    var v101 int16 = 42
                                                    fmt.Println(p2, fi7, v56)
                                                    _ = v101
                                                    _ = v98
                                                }
                                            }
                                            fmt.Println(p1, v51)
                                            var v102 string = "a"
                                            _ = v102
                                        }
                                        switch "hello" {
                                        case "hello":
                                            fmt.Println(p2, v51)
                                        case "a":
                                            v56 = (v56 >> ((p1) & 63))
                                        case "foo bar":
                                            p0 -= (v56 & 0)
                                        }
                                        fmt.Println(v56, p1, v52)
                                        _ = v56
                                    }
                                    var v103 uint64 = 0
                                    switch fi50 {
                                    case 32767:
                                        var v104 int64 = -1
                                        _ = v104
                                    default:
                                        var v105 int64 = -9223372036854775808
                                        _ = v105
                                    }
                                    _ = v103
                                    _ = v51
                                    _ = v52
                                    _ = v55
                                }
                                v106 := GenS12{F0: 0, F1: -112, F2: -590}
                                _ = fi50
                                _ = v106
                            }
                        }
                    }
                    var v107 uint = p1
                    _ = v107
                }
            }
            _ = fi7
        }
        var v108 string = (("foo bar" + "hello") + "a")
        _ = v108
    }
    v4 -= (^(v4))
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 625
}

func main() {
    var v1 int = 127
    var v2 uint16 = 2
    v3 := make(map[uint32]bool)
    for fi109 := 0; fi109 < 14; fi109++ {
        if false {
            delete(v3, 168)
            v110 := GenS13{F0: 18446744073709551615, F1: 679}
            v2 += (v2 << ((2) & 63))
            _ = v110
        } else {
            var v111 int64 = -78
            v112 := []int{(-(fi109)), v1, (fi109 / ((v1) | 1))}
            _ = v111
            _ = v112
        }
        v1 += 32767
        _ = fi109
    }
    switch (v1 | genFn0(254, 255, 128)) {
    case 1:
        if (false && false) {
            var v113 int = genFn0(0, 4294967295, 125)
            v114 := GenS14{F0: 585, F1: "hello"}
            _ = v113
            _ = v114
        } else {
            v1 += (v1 / (((v1 / (((v1 & (v1 >> ((2) & 63)))) | 1))) | 1))
            var v115 float32 = -2.25
            fmt.Println(v1, v1, v1)
            _ = v115
        }
    case 127:
        var v116 int64 = 42
        _ = v116
    default:
        fmt.Println(v2)
    }
    for fi117 := 0; fi117 < 10; fi117++ {
        v118 := []int{}
        _ = len(v118)
        for fi119 := 0; fi119 < 2; fi119++ {
            v120 := GenS14{F0: v2, F1: "foo bar"}
            v121 := GenS15{F0: 1}
            _ = fi119
            _ = v120
            _ = v121
        }
        _ = fi117
        _ = v118
    }
    v2 = v2
    v2 = v2
    for fi122 := 0; fi122 < 3; fi122++ {
        v1 += v1
        _ = fi122
    }
    switch -2147483648 {
    case -821:
        var v123 float64 = -2.25
        _ = v123
    case 2147483647:
        v124 := GenS4{F0: genFn0(1, 759, 157), F1: 128}
        _ = v124
    default:
        fmt.Println(v1, v1)
    }
    v125 := v3[2]
    var v126 bool = v125
    v127 := GenS16{F0: 1, F1: (!v126), F2: (^(v1))}
    fmt.Println(v126)
    if v125 {
        v2 = v2
        var v128 int16 = 32767
        var v129 uint16 = v2
        _ = v128
        _ = v129
    }
    v130 := []int{v1, -886, v1, v1}
    var v131 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v125:", v125)
    fmt.Println("v126:", v126)
    fmt.Println("slicelen v130:", len(v130), v130)
    fmt.Println("v131:", v131)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v125
    _ = v126
    _ = v127
    _ = v130
    _ = v131
    _ = v2
    _ = v3
}
