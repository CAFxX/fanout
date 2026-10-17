package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int32 }
type GenS2 struct { F0 uint8; F1 int8; F2 int32 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8; F1 string }
func genFn0(p0 bool) int16 {
    var v7 uint64 = 915
    v7 = (^(v7))
    p0 = (1 > 42)
    fmt.Println(p0)
    var v8 int64 = -9223372036854775808
    _ = p0
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 float32 = 1.5
    var v2 float64 = -2.25
    v3 := []int{-2147483648, -519, 67}
    v4 := []string{("\n\t" + ("" + "a")), "foo bar", ((("hello" + "a") + "foo bar") + "hello")}
    v5 := []int{193, -1}
    var v6 uint32 = 796
    fmt.Println(v1)
    switch "" {
    case "foo bar":
        switch ("\n\t" + ("hello" + "foo bar")) {
        case "\n\t":
            switch -2147483648 {
            case -905:
                fmt.Println(v1, v2)
            }
        }
    case "hello":
        v9 := GenS1{F0: 1, F1: 1, F2: -1}
        _ = v9
    case "a":
        var v10 int16 = 32767
        _ = v10
    default:
        switch "hello" {
        case "hello":
            v1 += v1
        default:
            if (255 < 116) {
                if false {
                    for fi11 := 0; fi11 < 4; fi11++ {
                        fmt.Println(v1)
                        if true {
                            v2 = v2
                            var v12 float64 = (v2 + (0.0 / v2))
                            _ = v12
                        }
                        v13 := GenS1{F0: -1, F1: 1, F2: 750}
                        _ = fi11
                        _ = v13
                    }
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v14 string = "a"
                    if true {
                        switch (v14 + "a") {
                        case "foo bar":
                            v6 += v6
                        case "":
                            for fi15 := 0; fi15 < 17; fi15++ {
                                var v16 int64 = 1
                                v17 := GenS1{F0: 103, F1: v16, F2: int32(v6)}
                                v6 = v6
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                        case "hello":
                            for fi18 := 0; fi18 < 9; fi18++ {
                                for fi19 := 0; fi19 < 10; fi19++ {
                                    v20 := GenS1{F0: -94, F1: 214, F2: 42}
                                    var v21 int32 = -755
                                    fmt.Println(v14, fi19)
                                    _ = fi19
                                    _ = v20
                                    _ = v21
                                }
                                _ = fi18
                            }
                        default:
                            if (v14 > v14) {
                                fmt.Println(v2, v1, v14)
                                if false {
                                    var v22 int32 = -258
                                    v2 = v2
                                    if len(v5) > 0 {
                                        v5[0] = 133
                                    }
                                    _ = v22
                                } else {
                                    v2 = v2
                                    switch 2147483647 {
                                    case 4294967295:
                                        if false {
                                            var v23 uint64 = 193
                                            _ = v23
                                        }
                                    }
                                    var v24 uint32 = v6
                                    _ = v24
                                }
                                if false {
                                    fmt.Println(v2)
                                    if true {
                                        fmt.Println(v14)
                                    } else {
                                        if false {
                                            v25 := GenS1{F0: 42, F1: 42, F2: 218}
                                            var v26 int = 378
                                            _ = v25
                                            _ = v26
                                        } else {
                                            v27 := GenS1{F0: 127, F1: -235, F2: 1}
                                            switch ((((("\n\t" + (v14 + v14)) + (("a" + "") + v14)) + "foo bar") + v14) + v14) {
                                            case "foo bar":
                                                fmt.Println(v6, v6, v2)
                                            case "\n\t":
                                                _ = len(v3)
                                            }
                                            for fi28 := 0; fi28 < 8; fi28++ {
                                                v14 = "a"
                                                _ = fi28
                                            }
                                            _ = v27
                                        }
                                        fmt.Println(v1)
                                    }
                                } else {
                                    var v29 uint = uint(v6)
                                    _ = v29
                                }
                            } else {
                                if true {
                                    v1 = ((0.0 + v1) * v1)
                                    if true {
                                        var v30 bool = (-128 == -128)
                                        fmt.Println(v14, v30)
                                        var v31 float32 = v1
                                        _ = v30
                                        _ = v31
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 32767
                                    }
                                } else {
                                    switch v14 {
                                    case "hello":
                                        for fi32 := 0; fi32 < 13; fi32++ {
                                            switch v14 {
                                            case "":
                                                if true {
                                                    v33 := GenS1{F0: 105, F1: 117, F2: -1}
                                                    _ = len(v5)
                                                    var v34 int16 = 1000
                                                    _ = v33
                                                    _ = v34
                                                } else {
                                                    v2 += 0.0
                                                }
                                            case "hello":
                                                for fi35 := 0; fi35 < 3; fi35++ {
                                                    var v36 int8 = -46
                                                    var v37 float64 = (((-2.25 * (((v2 * 0.0) * 1.5) / (v2 - v2))) + ((1e10 * v2) - v2)) * ((v2 - v2) / v2))
                                                    fmt.Println(v6)
                                                    var v38 float64 = (0.1 - v37)
                                                    v2 = v38
                                                    _ = fi35
                                                    _ = v36
                                                    _ = v37
                                                    _ = v38
                                                }
                                            }
                                            _ = fi32
                                        }
                                    case "":
                                        v14 = v14
                                    case "foo bar":
                                        if false {
                                            var v39 uint64 = 1
                                            var v40 float32 = v1
                                            switch -172 {
                                            case -166:
                                                if false {
                                                    v41 := GenS1{F0: 1, F1: 42, F2: -84}
                                                    for fi42 := 0; fi42 < 17; fi42++ {
                                                        var v43 float64 = ((v2 / 1e10) - v2)
                                                        fmt.Println(v43, v40, fi42)
                                                        var v44 int64 = 1
                                                        _ = fi42
                                                        _ = v43
                                                        _ = v44
                                                    }
                                                    _ = v41
                                                }
                                            case 42:
                                                fmt.Println(v39, v6)
                                            case -151:
                                                fmt.Println(v1, v1)
                                            }
                                            _ = v39
                                            _ = v40
                                        } else {
                                            switch -2147483648 {
                                            case -370:
                                                switch 128 {
                                                case 184:
                                                    switch 2147483647 {
                                                    case 740:
                                                        if len(v5) > 0 {
                                                            v5[0] = 32767
                                                        }
                                                    case 4294967295:
                                                        _ = len(v5)
                                                    case 1:
                                                        if len(v3) > 0 {
                                                            v3[0] = 32767
                                                        }
                                                    default:
                                                        for fi45 := 0; fi45 < 13; fi45++ {
                                                            switch fi45 {
                                                            case 297:
                                                                for fi46 := 0; fi46 < 11; fi46++ {
                                                                    if (!true) {
                                                                        if false {
                                                                            var v47 int64 = -1
                                                                            v48 := GenS2{F0: 220, F1: -128, F2: -966}
                                                                            v49 := GenS3{F0: fi45, F1: 0}
                                                                            _ = v47
                                                                            _ = v48
                                                                            _ = v49
                                                                        } else {
                                                                            v6 += (^(v6))
                                                                            var v50 float32 = v1
                                                                            _ = v50
                                                                        }
                                                                        var v51 int = 924
                                                                        _ = v51
                                                                    } else {
                                                                        if false {
                                                                            if ("a" > v14) {
                                                                                if false {
                                                                                    v14 = v14
                                                                                } else {
                                                                                    v52 := GenS4{F0: 0}
                                                                                    var v53 string = "foo bar"
                                                                                    v2 = (1e10 * (v2 + 3.14159))
                                                                                    _ = v52
                                                                                    _ = v53
                                                                                }
                                                                                var v54 string = ("" + (("hello" + v14) + ("" + "")))
                                                                                _ = v54
                                                                            } else {
                                                                                var v55 uint8 = 2
                                                                                switch (^(fi46)) {
                                                                                case 130:
                                                                                    v56 := GenS4{F0: 908}
                                                                                    _ = v56
                                                                                default:
                                                                                    if (364 > 665) {
                                                                                        v57 := GenS1{F0: -128, F1: 9223372036854775807, F2: -2147483648}
                                                                                        var v58 uint8 = uint8(v6)
                                                                                        _ = v57
                                                                                        _ = v58
                                                                                    }
                                                                                }
                                                                                v55 -= (^(v55))
                                                                                _ = v55
                                                                            }
                                                                            var v59 string = (v14 + (("" + v14) + v14))
                                                                            _ = len(v4)
                                                                            _ = v59
                                                                        }
                                                                    }
                                                                    v60 := GenS3{F0: (fi46 >> ((255) & 63)), F1: 0}
                                                                    fmt.Println(v1, v6)
                                                                    v4 = append(v4, (v14 + (v14 + v14)))
                                                                    _ = fi46
                                                                    _ = v60
                                                                }
                                                            }
                                                            v2 = 1e10
                                                            _ = fi45
                                                        }
                                                    }
                                                }
                                            case -605:
                                                for fi61 := 0; fi61 < 15; fi61++ {
                                                    fmt.Println(v2, v1)
                                                    v6 = (v6 + v6)
                                                    _ = fi61
                                                }
                                            }
                                            var v62 float32 = (-2.25 + v1)
                                            _ = v62
                                        }
                                    }
                                }
                            }
                        }
                        v14 = "foo bar"
                    } else {
                        switch 4294967295 {
                        case -762:
                            v63 := GenS5{F0: false}
                            _ = v63
                        default:
                            var v64 uint32 = v6
                            _ = v64
                        }
                    }
                    _ = v14
                }
                v5 = append(v5, 2147483647)
            }
        }
    }
    for fi65 := 0; fi65 < 17; fi65++ {
        for fi66 := 0; fi66 < 12; fi66++ {
            v67 := GenS5{F0: (true && false)}
            _ = fi66
            _ = v67
        }
        switch "hello" {
        case "hello":
            for fi68 := 0; fi68 < 10; fi68++ {
                if ((127 >= 73) || false) {
                    switch (fi68 & fi65) {
                    case -2147483648:
                        var v69 uint = 0
                        _ = v69
                    case 1:
                        fmt.Println(v1, v1, fi65)
                        switch fi65 {
                        case 32768:
                            v2 += 1e10
                        case 504:
                            fmt.Println(v1, fi68, fi68)
                        }
                        var v70 bool = (!true)
                        _ = v70
                    case 4294967295:
                        v71 := GenS4{F0: 995}
                        _ = v71
                    default:
                        fmt.Println(fi68, fi68, v2)
                    }
                }
                v1 -= float32(0.1)
                if false {
                    v6 -= (^(v6))
                    if (false && true) {
                        v6 = 227
                    }
                }
                v72 := GenS3{F0: fi65, F1: v6}
                v6 += v6
                v2 = v2
                _ = fi68
                _ = v72
            }
        case "a":
            switch "\n\t" {
            case "foo bar":
                v1 = float32((v2 / v2))
            default:
                fmt.Println(fi65)
            }
        default:
            v2 = 1.5
        }
        _ = fi65
    }
    for fi73 := 0; fi73 < 8; fi73++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v74 int16 = -1
        _ = fi73
        _ = v74
    }
    if (-37 <= 0) {
        for fi75 := 0; fi75 < 8; fi75++ {
            var v76 uint64 = 0
            v6 += (v6 ^ v6)
            _ = fi75
            _ = v76
        }
        for fi77 := 0; fi77 < 13; fi77++ {
            switch fi77 {
            case 32768:
                for fi78 := 0; fi78 < 20; fi78++ {
                    v4 = append(v4, "hello")
                    _ = fi78
                }
            case 32767:
                v79 := GenS4{F0: 214}
                _ = v79
            default:
                fmt.Println(v6, v2)
            }
            v5 = append(v5, (fi77 + (fi77 | fi77)))
            var v80 uint32 = (^(v6))
            _ = fi77
            _ = v80
        }
    }
    var v81 uint16 = 0
    v1 = float32(float64(v1))
    fmt.Println(v2)
    var v82 int16 = genFn0(true)
    _ = len(v4)
    fmt.Println(v1, v2, v2)
    v82 -= -635
    for fi83 := 0; fi83 < 18; fi83++ {
        fmt.Println(v82)
        v6 -= uint32(v81)
        v3 = append(v3, (fi83 << ((0) & 63)))
        _ = fi83
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v84 int8 = 42
    _ = len(v5)
    for fi85 := 0; fi85 < 15; fi85++ {
        if (false || (v6 <= (v6 % (((v6 | v6)) | 1)))) {
            switch "a" {
            case "\n\t":
                if ((float64(0.1) / v2) < 0.1) {
                    for fi86 := 0; fi86 < 16; fi86++ {
                        switch int(v6) {
                        case -9223372036854775808:
                            if false {
                                switch -796 {
                                case -393:
                                    v6 -= (^(v6))
                                case 127:
                                    v81 = (^(v81))
                                case 0:
                                    switch "hello" {
                                    case "hello":
                                        if (false || (v1 == 0.0)) {
                                            var v87 string = "a"
                                            _ = v87
                                        }
                                    case "foo bar":
                                        v82 = 0
                                    default:
                                        if (2 < uint(v6)) {
                                            v81 = (v81 % ((980) | 1))
                                        }
                                    }
                                }
                            }
                        default:
                            if true {
                                var v88 float64 = 1e10
                                _ = v88
                            } else {
                                v89 := GenS5{F0: (128 != 255)}
                                v84 = v84
                                _ = v89
                            }
                        }
                        var v90 int8 = v84
                        _ = fi86
                        _ = v90
                    }
                    for fi91 := 0; fi91 < 17; fi91++ {
                        v2 += v2
                        _ = fi91
                    }
                    v6 += (^(v6))
                }
            case "foo bar":
                for fi92 := 0; fi92 < 11; fi92++ {
                    var v93 bool = (-9223372036854775808 < -743)
                    if len(v5) > 0 {
                        v5[0] = (fi92 - fi85)
                    }
                    _ = fi92
                    _ = v93
                }
            case "":
                var v94 int32 = -646
                _ = v94
            }
            if ((v6 & 65535) == v6) {
                for fi95 := 0; fi95 < 15; fi95++ {
                    if ((-(fi85)) <= -9223372036854775808) {
                        if ((0 != 1) && true) {
                            var v96 uint64 = 42
                            v81 -= uint16(v96)
                            v82 += 198
                            var v97 float32 = 1.5
                            if false {
                                v81 += uint16(v6)
                            }
                            _ = v96
                            _ = v97
                        } else {
                            v84 += 111
                            var v98 float32 = v1
                            _ = v98
                        }
                        if false {
                            v1 = float32((v2 / (v2 / v2)))
                        } else {
                            var v99 uint16 = v81
                            _ = v99
                        }
                    } else {
                        if len(v4) > 0 {
                            v4[0] = ("hello" + "hello")
                        }
                        fmt.Println(fi95, fi85, v6)
                        var v100 float64 = (3.14159 + (v2 / ((v2 * v2) / v2)))
                        _ = v100
                    }
                    _ = fi95
                }
                v81 = v81
                v82 = (-(v82))
            }
            var v101 int64 = 1
            _ = v101
        } else {
            if (2147483647 <= (v6 & v6)) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v102 := GenS6{F0: "a"}
                _ = v102
            }
        }
        if true {
            v81 = v81
            v81 = v81
        } else {
            switch "\n\t" {
            case "hello":
                var v103 bool = true
                _ = v103
            case "foo bar":
                fmt.Println(v81, v81, v1)
            }
            fmt.Println(v2, v6)
            for fi104 := 0; fi104 < 16; fi104++ {
                if true {
                    _ = len(v4)
                } else {
                    if len(v4) > 0 {
                        v4[0] = "a"
                    }
                    switch fi104 {
                    case 0:
                        if len(v4) > 0 {
                            v4[0] = ("\n\t" + "foo bar")
                        }
                    case -1:
                        _ = len(v5)
                    }
                }
                v1 = 3.14159
                v105 := GenS7{F0: (v84 >> ((255) & 63)), F1: "a"}
                _ = fi104
                _ = v105
            }
        }
        v2 -= v2
        _ = fi85
    }
    if false {
        var v106 string = "hello"
        _ = v106
    } else {
        fmt.Println(v82)
        for fi107 := 0; fi107 < 19; fi107++ {
            switch int(v6) {
            case 2147483647:
                if (v6 >= v6) {
                    v82 += -523
                    if (!(int(v81) != -9223372036854775808)) {
                        v81 = 392
                    } else {
                        for fi108 := 0; fi108 < 12; fi108++ {
                            v5 = append(v5, (fi107 / ((20) | 1)))
                            if ((0.0 < 0.0) && (false && (false || (!true)))) {
                                v109 := GenS4{F0: 0}
                                _ = v109
                            }
                            v1 -= v1
                            _ = fi108
                        }
                        if len(v5) > 0 {
                            v5[0] = fi107
                        }
                        _ = len(v5)
                    }
                    switch fi107 {
                    case 0:
                        fmt.Println(v1)
                    default:
                        if ((v81 * 1) == v81) {
                            if true {
                                var v110 int64 = 42
                                _ = v110
                            } else {
                                if (!((v2 + float64(((((float32(0.0) + v1) + v1) - float32((float64(float32(1.5)) + float64((3.14159 / (((-2.25 - ((float32(v2) / -2.25) + (v1 + v1))) * (v1 * v1)) - float32(1.5))))))) - (v1 / 1.0)))) >= float64((v1 / v1)))) {
                                    v81 = v81
                                    fmt.Println(v82)
                                    v2 += (v2 + v2)
                                }
                                v2 += 1e10
                                fmt.Println(v6)
                            }
                            var v111 string = ("a" + ("a" + "hello"))
                            var v112 uint8 = 2
                            _ = v111
                            _ = v112
                        } else {
                            if len(v3) > 0 {
                                v3[0] = (fi107 + (fi107 % (((fi107 + fi107)) | 1)))
                            }
                            for fi113 := 0; fi113 < 4; fi113++ {
                                if len(v5) > 0 {
                                    v5[0] = -479
                                }
                                if len(v4) > 0 {
                                    v4[0] = "hello"
                                }
                                _ = fi113
                            }
                        }
                    }
                }
            case 0:
                v81 += v81
            case 1000:
                if false {
                    if len(v4) > 0 {
                        v4[0] = ((((("" + ("a" + ("a" + "foo bar"))) + (("hello" + "hello") + "a")) + "a") + "hello") + "")
                    }
                }
            default:
                for fi114 := 0; fi114 < 17; fi114++ {
                    fmt.Println(fi107)
                    var v115 float32 = (1.5 + 3.14159)
                    _ = fi114
                    _ = v115
                }
            }
            var v116 uint = 1
            _ = fi107
            _ = v116
        }
        var v117 string = (("\n\t" + "foo bar") + "hello")
        _ = v117
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v81
    _ = v82
    _ = v84
}
