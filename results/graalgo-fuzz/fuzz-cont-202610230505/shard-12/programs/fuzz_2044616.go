package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint8; F1 int64; F2 int16 }
type GenS10 struct { F0 uint8; F1 uint; F2 bool }
type GenS11 struct { F0 int16; F1 string; F2 uint32 }
type GenS12 struct { F0 int64 }
func genFn0() uint16 {
    v7 := []int{-2147483648, 1, 2147483647, 2147483647}
    for fi8 := 0; fi8 < 18; fi8++ {
        if true {
            _ = len(v7)
            var v9 int8 = 1
            var v10 string = "hello"
            _ = v10
            _ = v9
        } else {
            var v11 bool = true
            for fi12 := 0; fi12 < 20; fi12++ {
                v11 = true
                v13 := GenS1{F0: 883, F1: -1}
                v11 = false
                v7 = append(v7, fi8)
                if false {
                    var v14 float32 = 1.5
                    _ = v14
                }
                _ = fi12
                _ = v13
            }
            _ = v11
        }
        _ = fi8
    }
    _ = v7
    return 42
}
func genFn1(p0 uint32) int8 {
    var v15 float64 = 0.1
    var v16 uint16 = 32767
    v16 -= genFn0()
    v15 -= (v15 / ((v15 / (0.0 + v15)) + float64(-2.25)))
    p0 += 2
    _ = p0
    _ = v15
    _ = v16
    return 127
}

func main() {
    v1 := []string{"", "foo bar", ("\n\t" + ""), ""}
    var v2 uint = 998
    v3 := make(map[int]string)
    var v4 string = (("\n\t" + "foo bar") + "\n\t")
    v5 := make(map[string]bool)
    var v6 uint32 = 2147483647
    v2 = v2
    switch 4294967295 {
    case 32767:
        if (-1 < -128) {
            if (!true) {
                var v17 int64 = 1
                if false {
                    v2 = (v2 * (v2 + 446))
                    for fi18 := 0; fi18 < 8; fi18++ {
                        v4 = v4
                        v5[(v4 + (v4 + ((v4 + (v4 + "a")) + v4)))] = false
                        _ = fi18
                    }
                }
                var v19 float32 = ((float32(((1e10 - ((1e10 - -2.25) + 3.14159)) + -2.25)) * -2.25) - 1e10)
                _ = v17
                _ = v19
            } else {
                var v20 string = "hello"
                _ = v20
            }
            var v21 int32 = 294
            _ = v21
        } else {
            v6 = (v6 - v6)
            switch 127 {
            case 2147483647:
                v6 = v6
            case 0:
                for fi22 := 0; fi22 < 17; fi22++ {
                    for fi23 := 0; fi23 < 1; fi23++ {
                        v24 := GenS2{F0: (genFn1(v6) <= 42), F1: 65535}
                        _ = fi23
                        _ = v24
                    }
                    var v25 float64 = 0.1
                    var v26 int8 = -97
                    _ = fi22
                    _ = v25
                    _ = v26
                }
            case 42:
                var v27 int64 = -265
                _ = v27
            default:
                if ((127 != 14) && (1e10 > 0.1)) {
                    var v28 uint32 = (v6 / ((v6) | 1))
                    var v29 float64 = 0.0
                    v29 = (0.0 / v29)
                    _ = v28
                    _ = v29
                }
            }
        }
    case 257:
        switch v4 {
        case "\n\t":
            v4 = v4
        case "foo bar":
            if false {
                switch -9223372036854775808 {
                case -681:
                    for fi30 := 0; fi30 < 14; fi30++ {
                        v4 = v4
                        switch 128 {
                        case 2147483647:
                            var v31 uint64 = 2
                            _ = v31
                        default:
                            v1 = append(v1, "hello")
                        }
                        _ = fi30
                    }
                }
                fmt.Println(v2)
            } else {
                switch -2147483648 {
                case -427:
                    switch -1 {
                    case -324:
                        fmt.Println(v4, v2, v6)
                    case -9223372036854775808:
                        for fi32 := 0; fi32 < 16; fi32++ {
                            var v33 uint32 = 2
                            _ = fi32
                            _ = v33
                        }
                    }
                default:
                    if (-475 >= -1) {
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                        var v34 int16 = -32768
                        switch "\n\t" {
                        case "hello":
                            var v35 float32 = 1e10
                            _ = v35
                        case "foo bar":
                            v36 := GenS3{F0: v34}
                            _ = v36
                        }
                        _ = v34
                    } else {
                        switch -9223372036854775808 {
                        case 127:
                            if true {
                                if false {
                                    v37 := GenS4{F0: v6, F1: v6}
                                    _ = v37
                                }
                            } else {
                                fmt.Println(v6, v4, v6)
                            }
                        }
                    }
                }
                v4 = "a"
                v2 = (^(v2))
            }
        case "":
            switch (v4 + (v4 + "a")) {
            case "\n\t":
                if (-9223372036854775808 < 127) {
                    fmt.Println(v6, v6, v6)
                    if true {
                        switch "a" {
                        case "":
                            for fi38 := 0; fi38 < 5; fi38++ {
                                for fi39 := 0; fi39 < 18; fi39++ {
                                    fmt.Println(v4, v4)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v2 += (v2 >> ((v2) & 63))
                                    if true {
                                        v4 = "hello"
                                        v40 := GenS5{F0: 0}
                                        _ = v40
                                    } else {
                                        if (4294967295 >= 0) {
                                            if true {
                                                var v41 uint64 = 0
                                                _ = v41
                                            }
                                            switch v4 {
                                            case "\n\t":
                                                v6 = v6
                                            case "foo bar":
                                                switch v4 {
                                                case "a":
                                                    if (true || (4294967295 == 18446744073709551615)) {
                                                        var v42 uint = 18446744073709551615
                                                        v6 += 4294967295
                                                        _ = v42
                                                    } else {
                                                        fmt.Println(v2, v6, fi38)
                                                        v4 = (("\n\t" + v4) + "")
                                                    }
                                                    if (0 == 0) {
                                                        var v43 uint16 = 32767
                                                        if false {
                                                            var v44 uint32 = v6
                                                            fmt.Println(v4, fi38)
                                                            if (("hello" + v4) <= v4) {
                                                                fmt.Println(fi39)
                                                                v2 -= 65535
                                                                var v45 uint = v2
                                                                _ = v45
                                                            }
                                                            _ = v44
                                                        } else {
                                                            v46 := GenS5{F0: 42}
                                                            v1 = append(v1, "")
                                                            var v47 int16 = 1000
                                                            _ = v46
                                                            _ = v47
                                                        }
                                                        v4 = ((((((v4 + "a") + ("\n\t" + "\n\t")) + ("foo bar" + v4)) + "") + v4) + "\n\t")
                                                        _ = v43
                                                    }
                                                    fmt.Println(fi38, v2, fi39)
                                                    v48 := GenS1{F0: v6, F1: 847}
                                                    var v49 int8 = 91
                                                    if false {
                                                        v50 := GenS6{F0: (((v4 + "") + ("hello" + "foo bar")) + (v4 + (v4 + (v4 + ((v4 + "hello") + ("foo bar" + v4))))))}
                                                        switch ((v4 + ((v4 + "a") + v4)) + v4) {
                                                        case "":
                                                            var v51 int = fi39
                                                            _ = v51
                                                        case "hello":
                                                            v6 += v6
                                                        }
                                                        _ = v50
                                                    }
                                                    v2 = v2
                                                    if len(v1) > 0 {
                                                        v1[0] = v4
                                                    }
                                                    _ = v48
                                                    _ = v49
                                                }
                                            case "":
                                                var v52 uint = v2
                                                var v53 int64 = -9223372036854775808
                                                switch ("a" + "hello") {
                                                case "a":
                                                    switch "a" {
                                                    case "foo bar":
                                                        v2 = v52
                                                    case "\n\t":
                                                        var v54 int = fi39
                                                        _ = v54
                                                    }
                                                case "":
                                                    var v55 int = (fi38 & fi38)
                                                    _ = v55
                                                default:
                                                    v56 := GenS5{F0: int16(v6)}
                                                    _ = v56
                                                }
                                                var v57 uint32 = v6
                                                _ = v52
                                                _ = v53
                                                _ = v57
                                            }
                                            if true {
                                                if true {
                                                    v2 += 4294967295
                                                    switch ("" + (v4 + (((v4 + (v4 + (v4 + ("a" + v4)))) + (v4 + v4)) + v4))) {
                                                    case "":
                                                        v6 -= 65535
                                                    case "foo bar":
                                                        v4 = ("hello" + v4)
                                                    case "a":
                                                        if (-2.25 != (-2.25 + (3.14159 + (1e10 * ((1e10 / 3.14159) + (1.5 - float64(1e10))))))) {
                                                            var v58 int64 = 42
                                                            v59 := GenS1{F0: uint32(v2), F1: -122}
                                                            v60 := GenS7{F0: false}
                                                            _ = v58
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                    }
                                                    v4 = ("\n\t" + v4)
                                                }
                                            }
                                            v61 := GenS4{F0: (^(v6)), F1: (^(v6))}
                                            _ = v61
                                        }
                                    }
                                    var v62 uint32 = (v6 - (v6 - (^(v6))))
                                    _ = fi39
                                    _ = v62
                                }
                                var v63 string = ("\n\t" + ("" + v4))
                                _ = fi38
                                _ = v63
                            }
                        case "foo bar":
                            switch 32767 {
                            case -1:
                                var v64 string = "\n\t"
                                _ = v64
                            case 32767:
                                for fi65 := 0; fi65 < 4; fi65++ {
                                    v6 = (v6 - 0)
                                    v66 := GenS5{F0: 42}
                                    delete(v3, 518)
                                    _ = fi65
                                    _ = v66
                                }
                            }
                        case "hello":
                            switch 32768 {
                            case 32767:
                                var v67 uint64 = 18446744073709551615
                                _ = v67
                            }
                        }
                    }
                    if (2 == 0) {
                        v68 := GenS5{F0: -121}
                        _ = v68
                    } else {
                        v69 := v5[v4]
                        _ = v69
                    }
                } else {
                    var v70 int8 = 100
                    v70 = 10
                    _ = v70
                }
            default:
                v6 -= 423
            }
        }
    case 290:
        switch v4 {
        case "":
            for fi71 := 0; fi71 < 12; fi71++ {
                fmt.Println(v4, v4)
                v4 = "\n\t"
                _ = fi71
            }
        case "foo bar":
            switch 42 {
            case 42:
                switch (("a" + "\n\t") + ((v4 + v4) + v4)) {
                case "hello":
                    var v72 float32 = float32(((1.5 - 3.14159) / 3.14159))
                    _ = v72
                case "a":
                    _ = len(v3)
                default:
                    if ((v2 + v2) < (v2 - 42)) {
                        v4 = ((v4 + (v4 + v4)) + "")
                        if false {
                            v4 = v4
                            fmt.Println(v4, v6)
                        }
                        var v73 int32 = 920
                        _ = v73
                    }
                }
            }
        }
    default:
        var v74 uint64 = 1000
        _ = v74
    }
    fmt.Println(v6, v4, v2)
    v4 = v4
    if true {
        v75 := GenS6{F0: v4}
        _ = v75
    } else {
        switch v4 {
        case "a":
            for fi76 := 0; fi76 < 1; fi76++ {
                for fi77 := 0; fi77 < 1; fi77++ {
                    v6 -= (v6 ^ (v6 % ((65535) | 1)))
                    v78 := GenS5{F0: 1}
                    v4 = ("" + "a")
                    _ = fi77
                    _ = v78
                }
                _ = fi76
            }
        }
    }
    if ((!true) || (!true)) {
        if false {
            if (!true) {
                v6 = 65535
                delete(v5, (v4 + v4))
                v4 = (v4 + "a")
            }
        }
        if (2 > 255) {
            var v79 float64 = float64(3.14159)
            var v80 bool = (-128 != 0)
            _ = v79
            _ = v80
        }
        var v81 uint8 = 127
        _ = v81
    }
    for fi82 := 0; fi82 < 1; fi82++ {
        if false {
            if true {
                var v83 int16 = -79
                if false {
                    switch v4 {
                    case "hello":
                        var v84 bool = false
                        _ = v84
                    case "":
                        v6 += 2
                    case "\n\t":
                        v83 -= 1000
                    }
                } else {
                    v85 := GenS4{F0: 1, F1: (v6 - (^(v6)))}
                    var v86 uint32 = (v6 << ((v2) & 63))
                    if (!false) {
                        v87 := GenS8{F0: (v4 + (((("hello" + v4) + v4) + ("\n\t" + (v4 + v4))) + v4))}
                        _ = v87
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v88 := GenS9{F0: 1, F1: 9223372036854775807, F2: -326}
                        for fi89 := 0; fi89 < 19; fi89++ {
                            switch (((v4 + ((v4 + "") + ("" + "foo bar"))) + v4) + v4) {
                            case "\n\t":
                                fmt.Println(fi89, fi82)
                            case "foo bar":
                                var v90 int8 = 0
                                _ = len(v3)
                                _ = v90
                            default:
                                var v91 string = "\n\t"
                                _ = v91
                            }
                            switch (fi89 % ((475) | 1)) {
                            case 128:
                                v2 = (v2 + v2)
                            case -9223372036854775808:
                                v86 = v6
                            }
                            _ = fi89
                        }
                        _ = v88
                    }
                    _ = v85
                    _ = v86
                }
                _ = v83
            } else {
                for fi92 := 0; fi92 < 10; fi92++ {
                    if true {
                        switch v4 {
                        case "\n\t":
                            v93 := GenS10{F0: 105, F1: (^(v2)), F2: false}
                            _ = v93
                        default:
                            v94 := GenS11{F0: -1, F1: (v4 + ""), F2: v6}
                            _ = v94
                        }
                        if (false || true) {
                            v4 = v4
                            var v95 uint16 = 32767
                            v96 := GenS12{F0: -186}
                            _ = v95
                            _ = v96
                        } else {
                            if len(v1) > 0 {
                                v1[0] = ((v4 + (v4 + v4)) + "")
                            }
                        }
                    }
                    if len(v1) > 0 {
                        v1[0] = (v4 + "")
                    }
                    if (v2 <= 65535) {
                        fmt.Println(v4, v6)
                        v1 = append(v1, "\n\t")
                        if false {
                            switch (fi92 / (((fi92 ^ fi92)) | 1)) {
                            case 32768:
                                if true {
                                    var v97 float32 = float32(((1.5 / 3.14159) - float64(0.0)))
                                    fmt.Println(v4)
                                    switch v4 {
                                    case "a":
                                        var v98 uint8 = 2
                                        _ = v98
                                    default:
                                        var v99 uint32 = (v6 & v6)
                                        v4 = "a"
                                        _ = v99
                                    }
                                    _ = v97
                                }
                            }
                        } else {
                            v100 := GenS11{F0: 1, F1: ("hello" + v4), F2: (v6 | 52)}
                            _ = v100
                        }
                    }
                    delete(v3, (-(fi82)))
                    _ = fi92
                }
                for fi101 := 0; fi101 < 12; fi101++ {
                    if false {
                        var v102 uint16 = 65535
                        _ = v102
                    }
                    _ = fi101
                }
                v4 = v4
            }
        }
        for fi103 := 0; fi103 < 20; fi103++ {
            var v104 int = fi103
            v104 += fi103
            _ = fi103
            _ = v104
        }
        _ = fi82
    }
    v6 = (v6 * (v6 + v6))
    if (true && false) {
        var v105 uint32 = 2
        _ = v105
    }
    var v106 uint64 = 18446744073709551615
    if ((255 <= 964) && (-1 != -534)) {
        if (v106 > uint64(v6)) {
            switch int(v6) {
            case 4294967295:
                fmt.Println(v6, v2)
            default:
                delete(v3, -9223372036854775808)
            }
        } else {
            fmt.Println(v2, v106, v2)
            v107 := GenS3{F0: 198}
            _ = v107
        }
        for fi108 := 0; fi108 < 17; fi108++ {
            var v109 int8 = -10
            var v110 uint64 = (v106 & uint64(v6))
            switch fi108 {
            case -9223372036854775808:
                if (!false) {
                    for fi111 := 0; fi111 < 6; fi111++ {
                        if (genFn0() < 515) {
                            var v112 int = (fi111 & (fi108 - fi108))
                            v112 += (fi111 / ((fi108) | 1))
                            _ = v112
                        } else {
                            fmt.Println(v109, fi111, fi108)
                            var v113 int = fi108
                            _ = v113
                        }
                        var v114 float32 = 3.14159
                        _ = fi111
                        _ = v114
                    }
                    v115 := GenS5{F0: 1000}
                    _ = v115
                }
            case 169:
                if true {
                    v110 = v106
                }
            }
            _ = fi108
            _ = v109
            _ = v110
        }
        for fi116 := 0; fi116 < 6; fi116++ {
            fmt.Println(v6, v4, v2)
            v117 := GenS7{F0: true}
            if (true && true) {
                fmt.Println(v2)
            }
            _ = fi116
            _ = v117
        }
    }
    v4 = ((("\n\t" + ("a" + ("" + v4))) + v4) + ("foo bar" + ((v4 + v4) + v4)))
    _ = len(v3)
    var v118 uint8 = 244
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v106:", v106)
    fmt.Println("v118:", v118)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v106
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
