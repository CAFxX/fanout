package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 uint16; F2 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 uint; F2 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int64; F1 uint }
func genFn0(p0 uint16, p1 int16) int16 {
    var v6 bool = false
    fmt.Println(v6, p0, p1)
    for fi7 := 0; fi7 < 15; fi7++ {
        p0 = 255
        _ = fi7
    }
    v6 = false
    p1 -= p1
    _ = p0
    _ = p1
    _ = v6
    return (p1 + 32767)
}

func main() {
    var v1 int64 = -1
    var v2 uint32 = 65535
    var v3 int8 = 127
    var v4 int8 = v3
    v5 := make(map[string]bool)
    v8 := []int{-2147483648, 42}
    fmt.Println(v3)
    if ((1.5 - 0.0) >= 3.14159) {
        v4 += 32
    } else {
        v3 -= (v3 % ((v3) | 1))
        v3 = (v4 + (v4 >> ((65535) & 63)))
        v1 += -1
    }
    if (!(v2 > v2)) {
        switch 127 {
        case -1:
            v3 = (v4 & (v3 + v4))
        case 32768:
            v9 := GenS1{F0: 1}
            _ = v9
        }
        v2 += 2
        var v10 int16 = 42
        _ = v10
    }
    v1 = -1
    switch ("hello" + (("foo bar" + "a") + "\n\t")) {
    case "foo bar":
        fmt.Println(v3, v2)
    case "":
        _ = len(v5)
    case "\n\t":
        v4 += v3
    default:
        for fi11 := 0; fi11 < 17; fi11++ {
            v12 := GenS1{F0: 2}
            switch (fi11 << ((4294967295) & 63)) {
            case -653:
                fmt.Println(v3, v4, v3)
            case 127:
                switch "foo bar" {
                case "\n\t":
                    switch "\n\t" {
                    case "hello":
                        if false {
                            switch fi11 {
                            case 172:
                                var v13 int16 = 42
                                _ = v13
                            }
                            if (false && (413 > 950)) {
                                for fi14 := 0; fi14 < 12; fi14++ {
                                    var v15 string = ""
                                    v4 -= (^(v4))
                                    _ = fi14
                                    _ = v15
                                }
                            } else {
                                switch fi11 {
                                case 127:
                                    var v16 uint = 1
                                    _ = v16
                                default:
                                    for fi17 := 0; fi17 < 16; fi17++ {
                                        if ((!(9223372036854775807 > 9223372036854775807)) || true) {
                                            fmt.Println(v4)
                                        } else {
                                            fmt.Println(v3)
                                            v1 -= 169
                                            v3 = v3
                                        }
                                        v2 -= 2147483647
                                        _ = fi17
                                    }
                                }
                                var v18 string = ("\n\t" + "foo bar")
                                _ = v18
                            }
                            switch "hello" {
                            case "foo bar":
                                var v19 int32 = -2147483648
                                _ = v19
                            case "a":
                                for fi20 := 0; fi20 < 12; fi20++ {
                                    v21 := GenS1{F0: 560}
                                    _ = fi20
                                    _ = v21
                                }
                            case "\n\t":
                                if true {
                                    var v22 int32 = 0
                                    for fi23 := 0; fi23 < 3; fi23++ {
                                        if (("a" != "\n\t") && true) {
                                            v3 += 42
                                            switch fi23 {
                                            case 801:
                                                var v24 int16 = 42
                                                _ = v24
                                            case -4:
                                                v2 += v2
                                                fmt.Println(v3, v4)
                                            case 0:
                                                v3 += -22
                                            }
                                        }
                                        _ = fi23
                                    }
                                    v4 = -1
                                    _ = v22
                                }
                            }
                        }
                    default:
                        v2 = v2
                    }
                case "a":
                    for fi25 := 0; fi25 < 9; fi25++ {
                        v26 := GenS1{F0: 65535}
                        v3 = -128
                        var v27 int64 = (v1 << ((4294967295) & 63))
                        switch (fi25 % ((fi11) | 1)) {
                        case -920:
                            _ = len(v8)
                        case 1:
                            switch "foo bar" {
                            case "foo bar":
                                switch "" {
                                case "a":
                                    v28 := GenS1{F0: 535}
                                    _ = v28
                                case "\n\t":
                                    v29 := GenS1{F0: 544}
                                    _ = v29
                                case "hello":
                                    v30 := GenS1{F0: 58}
                                    _ = v30
                                default:
                                    fmt.Println(v4, v4)
                                    v3 = (v3 * v4)
                                }
                            case "a":
                                var v31 int64 = (v27 * v1)
                                _ = v31
                            case "":
                                switch "a" {
                                case "foo bar":
                                    var v32 float64 = ((((-2.25 * (1e10 + ((1e10 + ((1e10 * float64(1e10)) - 0.0)) - -2.25))) + (1.5 / 1.0)) - 3.14159) + 0.0)
                                    _ = v32
                                case "":
                                    fmt.Println(fi11)
                                    var v33 int64 = (v27 % ((v27) | 1))
                                    _ = v33
                                }
                            }
                        default:
                            var v34 float64 = (-2.25 + (1.5 + 1.5))
                            _ = v34
                        }
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                case "foo bar":
                    if true {
                        for fi35 := 0; fi35 < 14; fi35++ {
                            var v36 int32 = -493
                            _ = fi35
                            _ = v36
                        }
                        var v37 int16 = 997
                        _ = v37
                    }
                }
            }
            _ = fi11
            _ = v12
        }
    }
    if true {
        fmt.Println(v3, v4, v3)
        if (!(65535 > 2)) {
            for fi38 := 0; fi38 < 19; fi38++ {
                var v39 float32 = ((0.0 * 1e10) + 0.1)
                var v40 int8 = v4
                _ = fi38
                _ = v39
                _ = v40
            }
            v3 += 1
            var v41 float32 = 0.1
            _ = v41
        }
    } else {
        if (4294967295 < (v2 + (v2 * v2))) {
            var v42 bool = false
            fmt.Println(v42, v3, v42)
            v43 := v5[("foo bar" + "a")]
            _ = v42
            _ = v43
        } else {
            v4 += (^(v3))
            v1 -= v1
        }
        var v44 bool = ((^(v3)) < (v4 | (v3 >> ((0) & 63))))
        switch "hello" {
        case "a":
            if true {
                for fi45 := 0; fi45 < 10; fi45++ {
                    v3 -= v3
                    for fi46 := 0; fi46 < 2; fi46++ {
                        if (42 < (v4 + -33)) {
                            v5[("" + "a")] = (38 > 561)
                            if (2 > v2) {
                                v44 = v44
                                fmt.Println(fi46)
                                v2 = v2
                            } else {
                                var v47 string = "\n\t"
                                var v48 float64 = 0.0
                                _ = v47
                                _ = v48
                            }
                        }
                        _ = fi46
                    }
                    if v44 {
                        fmt.Println(v4, v2, v3)
                    } else {
                        v49 := v5["hello"]
                        if v44 {
                            if (247 < -813) {
                                v50 := GenS1{F0: 141}
                                _ = v50
                            } else {
                                switch 4294967295 {
                                case 32768:
                                    var v51 uint = 65535
                                    _ = v51
                                case -2147483648:
                                    var v52 int8 = 127
                                    _ = v52
                                case 32767:
                                    var v53 int8 = (v4 - v4)
                                    _ = v53
                                default:
                                    var v54 uint16 = 42
                                    _ = v54
                                }
                                var v55 string = "hello"
                                v56 := GenS1{F0: 65535}
                                _ = v55
                                _ = v56
                            }
                            for fi57 := 0; fi57 < 9; fi57++ {
                                v4 = (v4 | v3)
                                _ = fi57
                            }
                            if ("foo bar" != "foo bar") {
                                var v58 uint = 42
                                var v59 uint32 = v2
                                var v60 int16 = 1000
                                _ = v58
                                _ = v59
                                _ = v60
                            } else {
                                v49 = ((-(fi45)) >= (^(fi45)))
                                for fi61 := 0; fi61 < 7; fi61++ {
                                    v3 += 1
                                    _ = fi61
                                }
                            }
                        }
                        switch 32768 {
                        case -1:
                            for fi62 := 0; fi62 < 2; fi62++ {
                                if (!(0.1 >= (-2.25 - -2.25))) {
                                    _ = len(v5)
                                    var v63 uint32 = 2147483648
                                    _ = v63
                                } else {
                                    v4 = (v3 / ((100) | 1))
                                    var v64 int16 = 711
                                    var v65 int64 = (v1 * v1)
                                    _ = v64
                                    _ = v65
                                }
                                v66 := GenS2{F0: (fi62 << ((1) & 63)), F1: 32767, F2: (v4 / ((v4) | 1))}
                                if false {
                                    v1 = 1
                                    if ("a" < ("\n\t" + (((("hello" + "") + "") + "hello") + "\n\t"))) {
                                        v8 = append(v8, fi45)
                                        if true {
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                            if len(v8) > 0 {
                                                v8[0] = int(v1)
                                            }
                                            if false {
                                                var v67 int64 = -1
                                                _ = v67
                                            } else {
                                                if (4294967295 < 2) {
                                                    var v68 uint8 = 255
                                                    v44 = (v1 >= 0)
                                                    v3 += 1
                                                    _ = v68
                                                }
                                                _ = len(v8)
                                            }
                                        } else {
                                            var v69 uint = 0
                                            v70 := GenS1{F0: v69}
                                            _ = v69
                                            _ = v70
                                        }
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                    }
                                    var v71 int16 = 1000
                                    _ = v71
                                } else {
                                    _ = len(v8)
                                }
                                _ = fi62
                                _ = v66
                            }
                        case -2147483648:
                            if ((v4 != v3) || v49) {
                                for fi72 := 0; fi72 < 2; fi72++ {
                                    var v73 int16 = 42
                                    switch "" {
                                    case "hello":
                                        v74 := GenS3{F0: v44}
                                        _ = v74
                                    default:
                                        var v75 int64 = (v1 - v1)
                                        _ = v75
                                    }
                                    _ = fi72
                                    _ = v73
                                }
                                switch (fi45 & fi45) {
                                case 32767:
                                    var v76 uint32 = (v2 * (^(v2)))
                                    _ = v76
                                default:
                                    v3 = v4
                                }
                            }
                        default:
                            var v77 float32 = (1.5 / (((float32((1e10 / (0.1 * 1.5))) + 0.1) / float32(((-2.25 / 1e10) - 3.14159))) + (0.0 / 1.5)))
                            _ = v77
                        }
                        _ = v49
                    }
                    _ = fi45
                }
            }
        case "\n\t":
            var v78 int16 = -1
            _ = v78
        case "foo bar":
            switch "" {
            case "":
                v79 := GenS2{F0: -9223372036854775808, F1: 32767, F2: v3}
                _ = v79
            case "foo bar":
                if len(v8) > 0 {
                    v8[0] = 9223372036854775807
                }
            default:
                if (true && (127 != 188)) {
                    for fi80 := 0; fi80 < 3; fi80++ {
                        if (42 < 245) {
                            v2 += v2
                        } else {
                            var v81 int64 = (^(v1))
                            var v82 int16 = 0
                            _ = v81
                            _ = v82
                        }
                        for fi83 := 0; fi83 < 15; fi83++ {
                            _ = len(v5)
                            var v84 int64 = v1
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            _ = fi83
                            _ = v84
                        }
                        _ = fi80
                    }
                    v2 = v2
                }
            }
        }
        _ = v44
    }
    v4 += v4
    v3 = (^(v3))
    v85 := GenS4{F0: -715, F1: 255, F2: -2147483648}
    v3 = v4
    switch 141 {
    case -776:
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    default:
        if true {
            if (false || (!true)) {
                if (1 > -427) {
                    var v86 int32 = -591
                    _ = v86
                } else {
                    v2 = (v2 & (^(v2)))
                }
                var v87 float64 = 1e10
                _ = v87
            }
            var v88 float32 = -2.25
            _ = v88
        } else {
            switch -195 {
            case 0:
                switch 75 {
                case 1:
                    if ((((1.5 * (1.5 - 1.5)) - (3.14159 / float32((float64(-2.25) / float64((-2.25 + 0.1)))))) * float32((float64(1e10) / 1.0))) == (float32(0.0) + (float32((0.0 / (((3.14159 - (((-2.25 - ((3.14159 / 3.14159) + 0.0)) + 1.5) / -2.25)) - 3.14159) + 3.14159))) / (3.14159 + (3.14159 * float32(float64(0.0))))))) {
                        v4 = v4
                    } else {
                        for fi89 := 0; fi89 < 11; fi89++ {
                            if false {
                                var v90 int8 = 41
                                var v91 int16 = 32767
                                var v92 bool = (!false)
                                _ = v90
                                _ = v91
                                _ = v92
                            } else {
                                for fi93 := 0; fi93 < 9; fi93++ {
                                    var v94 int32 = 613
                                    var v95 float64 = ((3.14159 - 1.5) - float64(1e10))
                                    v4 = -128
                                    _ = fi93
                                    _ = v94
                                    _ = v95
                                }
                                switch "\n\t" {
                                case "foo bar":
                                    var v96 uint32 = (v2 - 4294967295)
                                    _ = v96
                                case "":
                                    v3 = (v3 - -128)
                                case "hello":
                                    for fi97 := 0; fi97 < 5; fi97++ {
                                        if false {
                                            if true {
                                                if true {
                                                    var v98 float32 = (((0.0 - -2.25) * (1e10 + 1.5)) - ((1.5 + (1e10 + 1e10)) * -2.25))
                                                    if false {
                                                        v3 = v4
                                                        var v99 bool = false
                                                        _ = v99
                                                    }
                                                    _ = v98
                                                } else {
                                                    v1 = 0
                                                    fmt.Println(fi89)
                                                    v100 := GenS5{F0: 1}
                                                    v101 := GenS5{F0: 389}
                                                    fmt.Println(v2)
                                                    _ = v100
                                                    _ = v101
                                                }
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v3 += 42
                                                case "hello":
                                                    v102 := GenS6{F0: 904}
                                                    _ = v102
                                                }
                                                v103 := GenS2{F0: (-(fi97)), F1: 0, F2: -105}
                                                _ = v103
                                            }
                                            var v104 uint64 = 0
                                            _ = v104
                                        }
                                        v3 = 0
                                        _ = fi97
                                    }
                                }
                            }
                            if ((-(v1)) > (v1 >> ((488) & 63))) {
                                v4 -= 127
                            } else {
                                switch (^(fi89)) {
                                case 9223372036854775807:
                                    var v105 uint = 1
                                    _ = v105
                                default:
                                    if (804 == 65535) {
                                        _ = len(v8)
                                        v2 = (v2 >> ((2) & 63))
                                        fmt.Println(v1, v3)
                                    }
                                }
                                if true {
                                    v4 += (^(v3))
                                    var v106 int = (-(fi89))
                                    for fi107 := 0; fi107 < 1; fi107++ {
                                        v108 := GenS7{F0: (v2 - (^(v2)))}
                                        v3 = v4
                                        fmt.Println(v4, fi107)
                                        fmt.Println(v1, fi107)
                                        _ = fi107
                                        _ = v108
                                    }
                                    _ = v106
                                }
                            }
                            _ = fi89
                        }
                        for fi109 := 0; fi109 < 10; fi109++ {
                            var v110 int8 = v4
                            _ = fi109
                            _ = v110
                        }
                        v2 -= v2
                    }
                default:
                    if (977 > 0) {
                        var v111 bool = (-944 >= 32767)
                        v112 := GenS8{F0: 0, F1: 4294967295}
                        var v113 int8 = v4
                        _ = v111
                        _ = v112
                        _ = v113
                    }
                }
            case 128:
                switch "" {
                case "a":
                    switch 770 {
                    case 879:
                        for fi114 := 0; fi114 < 14; fi114++ {
                            v3 = 1
                            _ = fi114
                        }
                    }
                case "\n\t":
                    v115 := GenS1{F0: 65535}
                    _ = v115
                }
            }
            v4 -= v4
        }
    }
    v4 -= 108
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v85
}
