package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 int8; F1 int32 }
type GenS6 struct { F0 int; F1 int16; F2 uint8 }
type GenS7 struct { F0 string; F1 uint8; F2 int8 }
type GenS8 struct { F0 uint; F1 uint }
type GenS9 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS10 struct { F0 int64; F1 string }
type GenS11 struct { F0 uint8; F1 uint8; F2 bool }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint32; F1 string; F2 uint }
type GenS14 struct { F0 int }
type GenS15 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS16 struct { F0 int32; F1 uint64 }
type GenS17 struct { F0 uint8; F1 uint16 }
type GenS18 struct { F0 bool; F1 uint64; F2 int64 }
type GenS19 struct { F0 int32; F1 string }
type GenS20 struct { F0 int64; F1 int8 }
func genFn0() int8 {
    var v4 int64 = -103
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 int64 = (^(v4))
        _ = fi5
        _ = v6
    }
    v4 = (v4 << ((42) & 63))
    _ = v4
    return 100
}

func main() {
    var v1 int32 = 185
    var v2 uint8 = 2
    var v3 uint8 = 118
    for fi7 := 0; fi7 < 6; fi7++ {
        if ((0.1 / -2.25) <= 0.1) {
            if ((!(!false)) && ((((1.5 * (0.1 * (0.1 * (-2.25 + float64(1.5))))) * 3.14159) * 1.5) > 0.0)) {
                for fi8 := 0; fi8 < 11; fi8++ {
                    if true {
                        v3 = v2
                        v1 = (-(v1))
                        fmt.Println(fi7)
                    }
                    v3 += (v2 >> ((65535) & 63))
                    v9 := []int{(-(fi7))}
                    _ = fi8
                    _ = v9
                }
                for fi10 := 0; fi10 < 9; fi10++ {
                    if (0 <= 410) {
                        switch ("a" + "hello") {
                        case "":
                            v11 := GenS1{F0: 1}
                            _ = v11
                        case "\n\t":
                            var v12 int64 = 198
                            _ = v12
                        default:
                            var v13 int8 = -128
                            _ = v13
                        }
                        v1 -= v1
                        v14 := []int{(fi10 | (fi7 + (fi10 / ((fi7) | 1)))), (-(fi7)), fi7}
                        _ = v14
                    } else {
                        v1 -= (^(v1))
                        v1 += v1
                        v1 = (v1 & (v1 % ((v1) | 1)))
                    }
                    if true {
                        if (128 <= 1) {
                            v1 = 863
                            v3 += v3
                        }
                        v15 := []int{(fi10 >> ((0) & 63)), (fi10 << ((65535) & 63)), (fi7 - 1)}
                        _ = v15
                    }
                    _ = fi10
                }
            }
            if true {
                if false {
                    if false {
                        v16 := GenS1{F0: 774}
                        _ = v16
                    }
                    var v17 uint32 = 2147483648
                    _ = v17
                }
            }
        }
        for fi18 := 0; fi18 < 9; fi18++ {
            if (!true) {
                var v19 uint8 = (v2 >> ((18446744073709551615) & 63))
                _ = v19
            }
            switch "\n\t" {
            case "foo bar":
                if (((("a" + ("" + "")) == "foo bar") || ("a" >= ("foo bar" + "a"))) || (!true)) {
                    switch "\n\t" {
                    case "a":
                        v1 = (^(v1))
                    case "hello":
                        v1 -= (v1 * v1)
                    default:
                        switch fi18 {
                        case 128:
                            var v20 string = "\n\t"
                            _ = v20
                        case 4294967295:
                            switch -9223372036854775808 {
                            case 128:
                                v21 := GenS2{F0: 255}
                                v2 = v3
                                switch (fi7 - (fi18 + 42)) {
                                case -783:
                                    fmt.Println(fi18, fi7)
                                }
                                _ = v21
                            default:
                                if (-2.25 > 3.14159) {
                                    v22 := GenS3{F0: -32768}
                                    if (9223372036854775807 <= 256) {
                                        fmt.Println(v3, v3, fi18)
                                        switch -983 {
                                        case 1:
                                            var v23 float32 = float32(-2.25)
                                            _ = v23
                                        case -1:
                                            v24 := GenS4{F0: 255, F1: fi7}
                                            _ = v24
                                        }
                                    }
                                    _ = v22
                                } else {
                                    switch (fi18 >> ((42) & 63)) {
                                    case -1:
                                        v25 := GenS3{F0: 42}
                                        _ = v25
                                    case 42:
                                        v1 = 207
                                    case 255:
                                        if (3.14159 >= (0.1 / 1.5)) {
                                            v1 -= v1
                                            v2 += v2
                                        } else {
                                            v26 := []int{-9223372036854775808, (fi18 * -812), fi7, -2147483648}
                                            var v27 string = "hello"
                                            _ = v26
                                            _ = v27
                                        }
                                    }
                                    v3 -= (v2 - v2)
                                }
                            }
                        default:
                            v28 := GenS2{F0: 65535}
                            _ = v28
                        }
                        v29 := []int{}
                        if len(v29) >= 2 {
                            v29 = v29[1:len(v29)]
                        }
                        v1 -= 42
                        fmt.Println(v2, fi18, fi18)
                        fmt.Println(v2, fi18, v3)
                        v30 := GenS2{F0: 961}
                        _ = v29
                        _ = v30
                    }
                    v31 := GenS3{F0: -32768}
                    v2 = v3
                    _ = v31
                }
            case "\n\t":
                v32 := GenS2{F0: 255}
                _ = v32
            case "a":
                v33 := GenS3{F0: -1}
                var v34 uint8 = v2
                var v35 uint32 = 927
                v36 := GenS1{F0: 153}
                _ = v33
                _ = v34
                _ = v35
                _ = v36
            }
            _ = fi18
        }
        _ = fi7
    }
    var v37 int32 = (v1 / ((42) | 1))
    switch ("hello" + "foo bar") {
    case "foo bar":
        fmt.Println(v1, v2, v37)
    case "\n\t":
        v37 = (-(v37))
    case "hello":
        v3 += (v3 & (^(v3)))
    }
    v38 := GenS5{F0: 127, F1: v37}
    if (((0.0 + 3.14159) / (3.14159 / 1.0)) > (0.0 - 3.14159)) {
        v37 = 0
        for fi39 := 0; fi39 < 1; fi39++ {
            if false {
                fmt.Println(v2, v3, v2)
            }
            fmt.Println(v1, fi39)
            if (((-2.25 / 1.0) - 1.5) > 0.1) {
                v3 = (v3 | (v3 | 0))
            } else {
                switch fi39 {
                case -1:
                    for fi40 := 0; fi40 < 5; fi40++ {
                        switch 128 {
                        case 32768:
                            v41 := GenS6{F0: fi40, F1: 42, F2: v3}
                            _ = v41
                        }
                        fmt.Println(v1, fi39, v3)
                        if (false || ((4294967295 > 270) || true)) {
                            v42 := GenS7{F0: ("hello" + ("a" + "")), F1: v3, F2: 37}
                            if true {
                                v37 = (-(v1))
                            }
                            _ = v42
                        }
                        var v43 int32 = (v37 ^ v1)
                        _ = fi40
                        _ = v43
                    }
                case -2147483648:
                    if ((v1 << ((65535) & 63)) < v1) {
                        for fi44 := 0; fi44 < 5; fi44++ {
                            if (true && true) {
                                if ((genFn0() == -128) || (1000 == 32767)) {
                                    if false {
                                        v2 = v3
                                        v45 := []int{962}
                                        fmt.Println(v3)
                                        _ = v45
                                    } else {
                                        v46 := []int{fi39, fi44}
                                        switch -291 {
                                        case -2147483648:
                                            _ = len(v46)
                                        }
                                        v47 := GenS8{F0: 1, F1: 2}
                                        _ = v46
                                        _ = v47
                                    }
                                }
                                v3 -= v2
                                switch "" {
                                case "":
                                    v48 := GenS7{F0: ("hello" + "\n\t"), F1: (v3 ^ (v2 / ((2) | 1))), F2: 55}
                                    _ = v48
                                case "\n\t":
                                    switch fi44 {
                                    case -9223372036854775808:
                                        v49 := []int{(fi44 ^ 9223372036854775807)}
                                        _ = v49
                                    default:
                                        if true {
                                            if true {
                                                var v50 int = fi39
                                                v50 = fi44
                                                var v51 int32 = 2147483647
                                                _ = v50
                                                _ = v51
                                            } else {
                                                v52 := []int{}
                                                _ = v52
                                            }
                                            v2 = (v2 / (((^(v2))) | 1))
                                        }
                                    }
                                    var v53 float64 = 1e10
                                    _ = v53
                                default:
                                    if ((((-582 == int16(fi44)) || true) || ((!true) || true)) && false) {
                                        if true {
                                            if false {
                                                v54 := []int{-446, fi44}
                                                _ = v54
                                            }
                                            v2 -= 225
                                            if true {
                                                var v55 uint = 65535
                                                _ = v55
                                            }
                                        }
                                        switch (fi39 + (fi44 / ((-883) | 1))) {
                                        case 0:
                                            v1 = v37
                                        case 2147483647:
                                            v56 := []int{(fi39 % ((-836) | 1)), 0, 1}
                                            _ = v56
                                        case -685:
                                            v57 := GenS4{F0: 786, F1: (fi44 - fi39)}
                                            _ = v57
                                        }
                                        switch "" {
                                        case "":
                                            v58 := GenS9{F0: (v2 >> ((65535) & 63)), F1: 0, F2: 343}
                                            _ = v58
                                        case "\n\t":
                                            var v59 uint = 1
                                            _ = v59
                                        default:
                                            v2 = v3
                                        }
                                    }
                                }
                            } else {
                                v3 = (v3 ^ v2)
                            }
                            switch "hello" {
                            case "":
                                fmt.Println(v2, v1)
                            case "a":
                                var v60 int16 = 1000
                                _ = v60
                            default:
                                fmt.Println(v37, v3, fi44)
                            }
                            if false {
                                v61 := []int{fi39}
                                var v62 uint64 = 42
                                fmt.Println(v37)
                                _ = v61
                                _ = v62
                            }
                            _ = fi44
                        }
                    }
                case 32767:
                    fmt.Println(fi39)
                default:
                    v2 += v2
                }
                v1 -= -919
                for fi63 := 0; fi63 < 19; fi63++ {
                    if false {
                        switch -2147483648 {
                        case -876:
                            v3 += v3
                        case -454:
                            if (4294967295 > 511) {
                                v64 := []int{(fi39 - 9223372036854775807), (^(fi63)), (fi39 * (fi39 ^ (fi63 << ((1) & 63))))}
                                v2 = v2
                                fmt.Println(v3)
                                _ = v64
                            } else {
                                if (!(!(1 > 18446744073709551615))) {
                                    switch ("foo bar" + ("a" + "hello")) {
                                    case "a":
                                        if true {
                                            v1 += (-(v1))
                                        } else {
                                            fmt.Println(v37)
                                            v65 := []int{}
                                            var v66 int64 = -1
                                            _ = v65
                                            _ = v66
                                        }
                                    case "foo bar":
                                        fmt.Println(v37)
                                    }
                                } else {
                                    if (!((("foo bar" + "") != ("\n\t" + "a")) || (-88 <= 100))) {
                                        v67 := GenS10{F0: 9223372036854775807, F1: "foo bar"}
                                        var v68 uint8 = (v3 / (((^(v2))) | 1))
                                        _ = v67
                                        _ = v68
                                    } else {
                                        if true {
                                            v2 -= 2
                                        }
                                        v37 = (v1 & -2147483648)
                                    }
                                }
                                var v69 int16 = 1000
                                var v70 int16 = (v69 + 32767)
                                _ = v69
                                _ = v70
                            }
                        }
                    } else {
                        v2 -= (^(v2))
                        var v71 int32 = 1
                        var v72 uint32 = 161
                        _ = v71
                        _ = v72
                    }
                    _ = fi63
                }
            }
            _ = fi39
        }
    } else {
        fmt.Println(v3, v1)
        fmt.Println(v3, v37, v37)
        switch 512 {
        case -1:
            v37 += (v37 | (v1 * (v37 - v1)))
        case -936:
            v73 := GenS5{F0: genFn0(), F1: (v1 + 2147483647)}
            _ = v73
        case 9223372036854775807:
            v74 := GenS7{F0: "foo bar", F1: (v3 + 1), F2: -125}
            _ = v74
        }
    }
    fmt.Println(v37, v3)
    v3 = 42
    fmt.Println(v37)
    switch "" {
    case "foo bar":
        v3 -= 117
    case "\n\t":
        v75 := GenS11{F0: v2, F1: v2, F2: (((-2.25 + 0.1) - 3.14159) >= 0.1)}
        _ = v75
    }
    v37 -= v1
    v76 := GenS10{F0: -540, F1: "\n\t"}
    fmt.Println(v1)
    v3 = v2
    switch 9223372036854775807 {
    case 127:
        var v77 int8 = 1
        _ = v77
    case 676:
        v78 := []int{-536, 0, -735, -1}
        _ = v78
    case 32768:
        fmt.Println(v2, v2, v37)
    }
    v79 := []int{441, 0, 32767, -9223372036854775808}
    if false {
        if true {
            fmt.Println(v3)
        } else {
            v80 := GenS12{F0: -5}
            switch 782 {
            case 438:
                for fi81 := 0; fi81 < 15; fi81++ {
                    v82 := GenS13{F0: 65535, F1: ("hello" + ""), F2: 4294967295}
                    for fi83 := 0; fi83 < 20; fi83++ {
                        v1 = -128
                        v3 -= v3
                        if true {
                            v79 = append(v79, (fi81 << ((18446744073709551615) & 63)))
                            switch "a" {
                            case "foo bar":
                                var v84 uint = 539
                                _ = v84
                            case "\n\t":
                                var v85 float32 = 3.14159
                                _ = v85
                            default:
                                v86 := GenS14{F0: 0}
                                _ = v86
                            }
                        } else {
                            if len(v79) >= 2 {
                                v79 = v79[1:len(v79)]
                            }
                        }
                        _ = fi83
                    }
                    _ = fi81
                    _ = v82
                }
            case 1:
                if len(v79) >= 2 {
                    v79 = v79[1:len(v79)]
                }
            default:
                var v87 bool = (("\n\t" + "hello") != "")
                _ = v87
            }
            v3 = v2
            _ = v80
        }
    } else {
        if true {
            if (4294967295 > 133) {
                if true {
                    for fi88 := 0; fi88 < 18; fi88++ {
                        v3 = 127
                        var v89 uint = 1
                        _ = fi88
                        _ = v89
                    }
                }
                if ((!false) && (false || ((v1 >= -1) || true))) {
                    v3 -= 0
                    v37 += 42
                    var v90 bool = false
                    _ = v90
                } else {
                    switch (("" + "\n\t") + "a") {
                    case "a":
                        for fi91 := 0; fi91 < 13; fi91++ {
                            if false {
                                for fi92 := 0; fi92 < 5; fi92++ {
                                    var v93 int8 = 0
                                    v3 = (v2 ^ 128)
                                    _ = fi92
                                    _ = v93
                                }
                            } else {
                                switch "hello" {
                                case "":
                                    v79 = append(v79, (fi91 / (((fi91 - (fi91 ^ 42))) | 1)))
                                case "foo bar":
                                    for fi94 := 0; fi94 < 3; fi94++ {
                                        v95 := GenS10{F0: 0, F1: ""}
                                        _ = fi94
                                        _ = v95
                                    }
                                }
                                fmt.Println(fi91, v2)
                            }
                            for fi96 := 0; fi96 < 3; fi96++ {
                                v97 := GenS15{F0: 2, F1: uint32(v2), F2: 1}
                                _ = fi96
                                _ = v97
                            }
                            v3 -= v2
                            _ = fi91
                        }
                    case "hello":
                        v3 += (^(v2))
                    case "\n\t":
                        var v98 uint32 = 586
                        _ = v98
                    }
                }
                switch "a" {
                case "":
                    fmt.Println(v3, v3, v1)
                case "foo bar":
                    switch 2147483647 {
                    case -226:
                        v99 := GenS2{F0: 1}
                        _ = v99
                    default:
                        v2 = (^(v3))
                    }
                case "\n\t":
                    v100 := GenS16{F0: v1, F1: 18446744073709551615}
                    _ = v100
                default:
                    if len(v79) >= 2 {
                        v79 = v79[1:len(v79)]
                    }
                }
            } else {
                fmt.Println(v3, v3, v1)
                if true {
                    if (1.5 == 0.1) {
                        v1 = v37
                        for fi101 := 0; fi101 < 9; fi101++ {
                            if (699 < 2) {
                                for fi102 := 0; fi102 < 5; fi102++ {
                                    v103 := GenS17{F0: 131, F1: 255}
                                    if len(v79) >= 2 {
                                        v79 = v79[1:len(v79)]
                                    }
                                    _ = fi102
                                    _ = v103
                                }
                                var v104 int64 = -1
                                _ = v104
                            }
                            var v105 string = "\n\t"
                            _ = fi101
                            _ = v105
                        }
                        var v106 int16 = 1
                        _ = v106
                    } else {
                        if len(v79) > 0 {
                            v79[0] = 0
                        }
                        for fi107 := 0; fi107 < 7; fi107++ {
                            var v108 uint16 = 571
                            if false {
                                var v109 int16 = 42
                                v110 := GenS18{F0: (81 <= 2), F1: 387, F2: 288}
                                _ = v109
                                _ = v110
                            } else {
                                v3 = v2
                                var v111 int64 = 340
                                _ = v111
                            }
                            for fi112 := 0; fi112 < 6; fi112++ {
                                if len(v79) >= 2 {
                                    v79 = v79[1:len(v79)]
                                }
                                var v113 uint16 = (v108 & (v108 - uint16(v3)))
                                _ = fi112
                                _ = v113
                            }
                            _ = fi107
                            _ = v108
                        }
                        switch (("\n\t" + "\n\t") + "a") {
                        case "foo bar":
                            if false {
                                for fi114 := 0; fi114 < 7; fi114++ {
                                    if true {
                                        if false {
                                            var v115 int64 = 495
                                            if len(v79) > 0 {
                                                v79[0] = (fi114 - 4294967295)
                                            }
                                            _ = v115
                                        } else {
                                            v116 := GenS19{F0: (v1 | (v1 ^ (v37 / ((0) | 1)))), F1: "a"}
                                            var v117 int = (^(fi114))
                                            _ = v116
                                            _ = v117
                                        }
                                        v37 = v37
                                    }
                                    _ = fi114
                                }
                            } else {
                                v79 = append(v79, 4294967295)
                            }
                        default:
                            if true {
                                v118 := GenS20{F0: -9223372036854775808, F1: -53}
                                var v119 string = "hello"
                                if len(v79) >= 2 {
                                    v79 = v79[1:len(v79)]
                                }
                                _ = v118
                                _ = v119
                            } else {
                                fmt.Println(v1, v3, v1)
                            }
                        }
                    }
                }
            }
            v2 = 128
        } else {
            var v120 uint64 = 835
            v37 -= v1
            _ = v120
        }
        v3 += (v3 >> ((65535) & 63))
        v1 += v1
    }
    fmt.Println(v37)
    switch -465 {
    case 819:
        v2 = 127
    case -932:
        for fi121 := 0; fi121 < 4; fi121++ {
            fmt.Println(fi121, fi121, fi121)
            v2 += 255
            _ = len(v79)
            _ = fi121
        }
    default:
        var v122 float32 = (float32(1.5) - 0.1)
        _ = v122
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v79:", len(v79), v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v76
    _ = v79
}
