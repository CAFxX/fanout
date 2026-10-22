package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 string }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16; F1 int }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 bool; F1 uint }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint; F1 uint8; F2 int32 }
type GenS11 struct { F0 uint8; F1 bool }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]int)
    var v3 string = ("\n\t" + "hello")
    v4 := []string{}
    switch v3 {
    case "foo bar":
        v3 = "a"
    case "a":
        switch "a" {
        case "hello":
            if len(v4) > 0 {
                v4[0] = v3
            }
        case "":
            var v5 int16 = 0
            _ = v5
        }
    }
    switch 32767 {
    case 127:
        for fi6 := 0; fi6 < 15; fi6++ {
            if false {
                fmt.Println(v3, v3)
                fmt.Println(v3)
                var v7 int8 = 1
                _ = v7
            } else {
                v8 := GenS1{F0: ((v3 + (((v3 + ("foo bar" + v3)) + "") + v3)) + v3), F1: -15, F2: (((v3 + v3) + v3) + v3)}
                if false {
                    if (false && (true || (true && (!(0.1 > 1.5))))) {
                        switch fi6 {
                        case 2147483647:
                            switch ("" + "\n\t") {
                            case "foo bar":
                                var v9 int64 = -586
                                _ = v9
                            case "":
                                v3 = "foo bar"
                            }
                        case -1:
                            v3 = v3
                        default:
                            var v10 int32 = 0
                            _ = v10
                        }
                    } else {
                        for fi11 := 0; fi11 < 8; fi11++ {
                            if ((fi11 | (fi11 ^ fi11)) != (fi6 >> ((0) & 63))) {
                                switch fi6 {
                                case 0:
                                    if false {
                                        v12 := GenS2{F0: -568, F1: (9223372036854775807 < 42)}
                                        switch fi6 {
                                        case 0:
                                            fmt.Println(fi11)
                                            if ((true && (42 >= 42)) || false) {
                                                var v13 int = (fi11 + fi6)
                                                fmt.Println(fi11, fi6)
                                                _ = v13
                                            } else {
                                                if false {
                                                    var v14 float64 = (1.5 - (1e10 * float64((3.14159 + (0.1 / 1.0)))))
                                                    var v15 int64 = -1
                                                    fmt.Println(fi11, fi6)
                                                    _ = v14
                                                    _ = v15
                                                } else {
                                                    switch fi11 {
                                                    case 42:
                                                        v2[(^(fi11))] = (fi6 ^ fi6)
                                                    case 414:
                                                        if false {
                                                            v16 := GenS1{F0: v3, F1: -9223372036854775808, F2: "hello"}
                                                            var v17 uint64 = 121
                                                            _ = v16
                                                            _ = v17
                                                        }
                                                    }
                                                    delete(v2, (^(fi11)))
                                                    v2[0] = fi6
                                                }
                                                v3 = ((v3 + v3) + v3)
                                            }
                                        case -891:
                                            fmt.Println(v3)
                                        default:
                                            var v18 int32 = 2147483647
                                            switch "" {
                                            case "hello":
                                                var v19 string = "\n\t"
                                                _ = v19
                                            }
                                            _ = v18
                                        }
                                        _ = v12
                                    }
                                case 422:
                                    var v20 uint8 = 49
                                    _ = v20
                                case 9223372036854775807:
                                    v3 = v3
                                }
                            }
                            switch (-(fi11)) {
                            case 278:
                                if true {
                                    if (0 == 0) {
                                        if false {
                                            _ = len(v2)
                                            v3 = v3
                                        } else {
                                            v3 = "a"
                                            switch v3 {
                                            case "":
                                                v3 = v3
                                            default:
                                                var v21 uint = 868
                                                _ = v21
                                            }
                                        }
                                        var v22 uint = 254
                                        var v23 int16 = 0
                                        _ = v22
                                        _ = v23
                                    } else {
                                        if true {
                                            v3 = v3
                                        } else {
                                            fmt.Println(v3)
                                        }
                                        v24 := GenS3{F0: (fi11 << ((298) & 63)), F1: -18}
                                        fmt.Println(fi6)
                                        _ = v24
                                    }
                                    v25 := GenS1{F0: v3, F1: -327, F2: v3}
                                    _ = v25
                                }
                            case -2147483648:
                                v26 := GenS4{F0: -186}
                                _ = v26
                            case 116:
                                fmt.Println(v3)
                            default:
                                v3 = "foo bar"
                            }
                            _ = fi11
                        }
                        v3 = ((("" + (v3 + (v3 + ""))) + "hello") + "\n\t")
                        switch fi6 {
                        case 127:
                            for fi27 := 0; fi27 < 9; fi27++ {
                                v28 := GenS5{F0: 0}
                                var v29 int32 = -1
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                        default:
                            fmt.Println(fi6)
                        }
                    }
                    v3 = "a"
                }
                if false {
                    if (1.5 != (0.1 + ((0.1 / 1.5) - ((-2.25 * 3.14159) + ((0.0 + 1.5) + -2.25))))) {
                        var v30 bool = true
                        v2[(fi6 >> ((4294967295) & 63))] = fi6
                        v30 = v30
                        _ = v30
                    }
                } else {
                    if (200 == 42) {
                        v3 = "\n\t"
                    } else {
                        if ((854 != 0) || (((-2.25 + ((1.5 + 1.5) + 1e10)) < 0.1) || false)) {
                            if false {
                                v3 = ("" + (((("foo bar" + (((v3 + v3) + ((v3 + "hello") + "foo bar")) + v3)) + v3) + (v3 + v3)) + v3))
                                fmt.Println(fi6, v3)
                            } else {
                                if (true && (1 >= 42)) {
                                    fmt.Println(fi6, v3)
                                    v3 = v3
                                }
                                v4 = append(v4, v3)
                                switch v3 {
                                case "":
                                    v31 := GenS6{F0: 255, F1: fi6}
                                    _ = v31
                                case "\n\t":
                                    fmt.Println(v3, v3)
                                case "foo bar":
                                    if true {
                                        switch fi6 {
                                        case 128:
                                            v2[(fi6 & (fi6 / (((-(fi6))) | 1)))] = fi6
                                        default:
                                            v3 = v3
                                        }
                                    }
                                }
                            }
                            v32 := GenS2{F0: -9223372036854775808, F1: (0 <= 2147483647)}
                            if (127 >= 1) {
                                if (!false) {
                                    switch (fi6 / (((^(fi6))) | 1)) {
                                    case -617:
                                        for fi33 := 0; fi33 < 2; fi33++ {
                                            v34 := GenS4{F0: fi6}
                                            switch v3 {
                                            case "\n\t":
                                                var v35 bool = (("\n\t" >= v3) && (0 != -52))
                                                _ = v35
                                            case "hello":
                                                v36 := GenS4{F0: (fi6 + (fi33 ^ fi6))}
                                                _ = v36
                                            case "":
                                                fmt.Println(fi33, fi6)
                                            }
                                            _ = len(v4)
                                            var v37 uint16 = 914
                                            if (-1 >= -1) {
                                                v37 += (v37 | (^(v37)))
                                                v37 = 711
                                                if (!true) {
                                                    var v38 float64 = 1.5
                                                    _ = v38
                                                }
                                            }
                                            var v39 int32 = 2147483647
                                            _ = fi33
                                            _ = v34
                                            _ = v37
                                            _ = v39
                                        }
                                    case 127:
                                        v3 = (v3 + "")
                                    case -9223372036854775808:
                                        if true {
                                            v3 = (v3 + "\n\t")
                                            v40 := GenS6{F0: 65535, F1: fi6}
                                            switch (fi6 ^ fi6) {
                                            case 32767:
                                                fmt.Println(v3, fi6)
                                            case 127:
                                                if (2147483647 >= -306) {
                                                    if false {
                                                        var v41 uint16 = 2
                                                        _ = v41
                                                    }
                                                }
                                            case 488:
                                                fmt.Println(v3, fi6)
                                            default:
                                                v3 = "hello"
                                            }
                                            _ = v40
                                        } else {
                                            switch (^(fi6)) {
                                            case 2147483647:
                                                for fi42 := 0; fi42 < 5; fi42++ {
                                                    v43 := GenS3{F0: fi42, F1: -1}
                                                    var v44 uint32 = 962
                                                    _ = fi42
                                                    _ = v43
                                                    _ = v44
                                                }
                                            default:
                                                var v45 uint32 = 4294967295
                                                _ = v45
                                            }
                                            v1[42] = (false || true)
                                            if (!true) {
                                                var v46 uint32 = 4294967295
                                                _ = v46
                                            } else {
                                                var v47 int16 = 1
                                                _ = v47
                                            }
                                        }
                                    default:
                                        switch v3 {
                                        case "a":
                                            for fi48 := 0; fi48 < 20; fi48++ {
                                                var v49 uint32 = 2147483647
                                                _ = fi48
                                                _ = v49
                                            }
                                        case "hello":
                                            switch fi6 {
                                            case -2147483648:
                                                fmt.Println(fi6, v3)
                                            }
                                        case "foo bar":
                                            for fi50 := 0; fi50 < 7; fi50++ {
                                                var v51 uint8 = 255
                                                _ = fi50
                                                _ = v51
                                            }
                                        default:
                                            fmt.Println(fi6, fi6)
                                        }
                                    }
                                } else {
                                    var v52 float32 = (0.1 / 1.0)
                                    v53 := GenS7{F0: 127}
                                    var v54 uint8 = 127
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                }
                                switch (v3 + (v3 + (v3 + v3))) {
                                case "hello":
                                    var v55 int64 = -9223372036854775808
                                    _ = v55
                                }
                            } else {
                                var v56 bool = true
                                if (!v56) {
                                    var v57 float64 = 0.1
                                    _ = v57
                                }
                                _ = v56
                            }
                            _ = v32
                        }
                        for fi58 := 0; fi58 < 15; fi58++ {
                            v59 := GenS2{F0: 42, F1: (165 > 0)}
                            v60 := GenS3{F0: (fi58 | 127), F1: 1}
                            _ = fi58
                            _ = v59
                            _ = v60
                        }
                    }
                    var v61 int16 = -1
                    _ = v61
                }
                _ = v8
            }
            _ = fi6
        }
    default:
        if (false && true) {
            v3 = "\n\t"
            v3 = ""
            switch -450 {
            case 2147483647:
                var v62 int = 351
                _ = v62
            default:
                for fi63 := 0; fi63 < 2; fi63++ {
                    v64 := GenS6{F0: 853, F1: fi63}
                    if true {
                        for fi65 := 0; fi65 < 2; fi65++ {
                            if true {
                                v66 := GenS5{F0: -509}
                                _ = v66
                            } else {
                                v67 := GenS7{F0: -128}
                                _ = v67
                            }
                            _ = fi65
                        }
                        var v68 int8 = 0
                        _ = v68
                    } else {
                        switch (("\n\t" + "") + v3) {
                        case "":
                            var v69 uint8 = 2
                            _ = v69
                        }
                    }
                    switch v3 {
                    case "":
                        v1[-279] = false
                    case "hello":
                        v3 = v3
                    case "a":
                        fmt.Println(v3)
                    default:
                        v3 = v3
                    }
                    _ = fi63
                    _ = v64
                }
            }
        } else {
            if false {
                if len(v4) > 0 {
                    v4[0] = v3
                }
            } else {
                for fi70 := 0; fi70 < 1; fi70++ {
                    v4 = append(v4, "a")
                    if (3.14159 <= (0.0 / 0.1)) {
                        switch 127 {
                        case 232:
                            v3 = v3
                        case 1:
                            var v71 string = v3
                            _ = v71
                        case 32767:
                            v72 := GenS7{F0: 42}
                            _ = v72
                        }
                        var v73 float64 = float64(float32((((-2.25 - float64(3.14159)) - 3.14159) / float64(float32(-2.25)))))
                        var v74 bool = true
                        _ = v73
                        _ = v74
                    } else {
                        fmt.Println(v3, fi70)
                    }
                    v3 = v3
                    _ = fi70
                }
                for fi75 := 0; fi75 < 4; fi75++ {
                    var v76 int32 = 1
                    if len(v4) > 0 {
                        v4[0] = "hello"
                    }
                    v77 := GenS6{F0: 65535, F1: fi75}
                    _ = fi75
                    _ = v76
                    _ = v77
                }
            }
            if false {
                var v78 string = v3
                fmt.Println(v3)
                _ = v78
            } else {
                var v79 int = -2147483648
                _ = v79
            }
        }
    }
    if (!((1 == 1) && false)) {
        v3 = "hello"
    } else {
        fmt.Println(v3)
        var v80 float32 = 1.5
        v81 := v2[9223372036854775807]
        _ = v80
        _ = v81
    }
    v3 = (((v3 + ((v3 + (("a" + ("hello" + (v3 + v3))) + (((("\n\t" + (v3 + v3)) + (((((("hello" + (v3 + (("\n\t" + (("\n\t" + (v3 + "")) + ((v3 + (("hello" + ("foo bar" + ("a" + v3))) + "a")) + ((v3 + v3) + "")))) + v3))) + v3) + "\n\t") + "a") + v3) + "foo bar")) + v3) + "hello"))) + v3)) + (v3 + (("\n\t" + "hello") + v3))) + ((v3 + (v3 + (((v3 + (v3 + (v3 + v3))) + v3) + "hello"))) + (("foo bar" + "foo bar") + ((v3 + (("\n\t" + "") + (v3 + v3))) + "hello"))))
    v3 = (((v3 + v3) + ("" + "\n\t")) + "hello")
    v82 := GenS8{F0: false, F1: 0}
    v3 = v3
    v3 = v3
    switch 32767 {
    case 193:
        fmt.Println(v3)
    }
    if (!(!false)) {
        fmt.Println(v3)
        var v83 int = 127
        v3 = v3
        _ = v83
    }
    switch "foo bar" {
    case "a":
        if true {
            v4 = append(v4, "\n\t")
            for fi84 := 0; fi84 < 5; fi84++ {
                if true {
                    v85 := GenS5{F0: 42}
                    if false {
                        v4 = append(v4, v3)
                        v3 = v3
                        v1[(^(fi84))] = (((1 == 0) || true) && true)
                    }
                    for fi86 := 0; fi86 < 11; fi86++ {
                        if false {
                            var v87 uint16 = 65535
                            if (!false) {
                                var v88 bool = (((("hello" + (("foo bar" + "a") + v3)) + (v3 + ("foo bar" + (v3 + (v3 + (v3 + v3)))))) + (v3 + ((v3 + "") + (((v3 + "\n\t") + (("hello" + "") + "\n\t")) + v3)))) != v3)
                                _ = v88
                            }
                            _ = v87
                        } else {
                            var v89 int32 = -2147483648
                            var v90 uint32 = 361
                            _ = v89
                            _ = v90
                        }
                        if (true || (!true)) {
                            v91 := GenS8{F0: false, F1: 18446744073709551615}
                            v2[fi84] = 42
                            _ = v91
                        }
                        _ = fi86
                    }
                    _ = v85
                } else {
                    fmt.Println(fi84)
                    var v92 uint = 65535
                    fmt.Println(v3)
                    _ = v92
                }
                for fi93 := 0; fi93 < 7; fi93++ {
                    v3 = (("hello" + v3) + "hello")
                    _ = fi93
                }
                _ = fi84
            }
        } else {
            v94 := v2[-2147483648]
            if ((("\n\t" + v3) + "\n\t") != (v3 + (((v3 + ("" + v3)) + "\n\t") + v3))) {
                for fi95 := 0; fi95 < 4; fi95++ {
                    v2[(v94 << ((594) & 63))] = (^(fi95))
                    v94 -= fi95
                    _ = fi95
                }
            } else {
                v3 = v3
                for fi96 := 0; fi96 < 20; fi96++ {
                    for fi97 := 0; fi97 < 5; fi97++ {
                        switch fi96 {
                        case 128:
                            delete(v2, (fi97 & 582))
                        case -2147483648:
                            fmt.Println(v3, fi96)
                        }
                        v94 = 32768
                        v3 = (v3 + v3)
                        if false {
                            v98 := GenS5{F0: 0}
                            _ = v98
                        }
                        _ = fi97
                    }
                    _ = fi96
                }
            }
            for fi99 := 0; fi99 < 16; fi99++ {
                fmt.Println(fi99)
                _ = fi99
            }
            _ = v94
        }
    case "\n\t":
        for fi100 := 0; fi100 < 18; fi100++ {
            if false {
                v1[32768] = (365 > 2147483647)
            } else {
                _ = len(v4)
                switch (fi100 >> ((255) & 63)) {
                case 32767:
                    if (false && ((float32(1.5) > 0.0) && false)) {
                        v3 = (("hello" + (v3 + "foo bar")) + v3)
                    } else {
                        var v101 int32 = 2147483647
                        _ = v101
                    }
                case 32768:
                    if ((float32(((-2.25 / 1.0) * 0.1)) / float32(1e10)) < 0.1) {
                        v102 := GenS5{F0: 1}
                        _ = v102
                    } else {
                        var v103 float32 = 1e10
                        _ = v103
                    }
                }
                v2[(fi100 | fi100)] = 91
            }
            v3 = v3
            switch fi100 {
            case -9223372036854775808:
                v3 = v3
            default:
                v3 = v3
            }
            _ = fi100
        }
    default:
        if false {
            var v104 uint32 = 65535
            switch (v3 + ("hello" + v3)) {
            case "a":
                switch 2147483647 {
                case 32767:
                    switch (((("\n\t" + "hello") + v3) + v3) + ("a" + "\n\t")) {
                    case "foo bar":
                        if (false || true) {
                            v104 -= 65535
                            v105 := GenS1{F0: ("a" + v3), F1: 42, F2: v3}
                            _ = v105
                        }
                    case "a":
                        fmt.Println(v104, v104)
                    default:
                        fmt.Println(v3, v104)
                    }
                case 448:
                    v3 = v3
                case 0:
                    v3 = v3
                }
            case "hello":
                switch 496 {
                case 1:
                    var v106 float32 = -2.25
                    _ = v106
                default:
                    for fi107 := 0; fi107 < 16; fi107++ {
                        v3 = ("hello" + v3)
                        fmt.Println(fi107, fi107, v104)
                        var v108 float64 = 3.14159
                        _ = fi107
                        _ = v108
                    }
                }
            default:
                v109 := GenS9{F0: 42}
                _ = v109
            }
            _ = v104
        }
    }
    fmt.Println(v3)
    v110 := GenS10{F0: 4294967295, F1: 2, F2: 749}
    v3 = (v3 + ("" + ("" + "\n\t")))
    if false {
        if false {
            for fi111 := 0; fi111 < 9; fi111++ {
                fmt.Println(v3)
                v3 = ("\n\t" + ((v3 + (((v3 + v3) + ("hello" + v3)) + v3)) + v3))
                _ = fi111
            }
        }
    } else {
        fmt.Println(v3)
        v112 := GenS11{F0: 127, F1: (!false)}
        _ = v112
    }
    var v113 int64 = 42
    fmt.Println(v113, v113)
    switch v3 {
    case "a":
        if true {
            var v114 float64 = float64(-2.25)
            _ = v114
        }
    case "foo bar":
        v3 = v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v113:", v113)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v110
    _ = v113
    _ = v2
    _ = v3
    _ = v4
    _ = v82
}
