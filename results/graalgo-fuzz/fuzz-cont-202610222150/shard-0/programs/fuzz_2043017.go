package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 int; F1 uint64; F2 uint8 }
type GenS6 struct { F0 bool; F1 string; F2 int }
type GenS7 struct { F0 int; F1 uint16; F2 int16 }
type GenS8 struct { F0 uint32; F1 string; F2 uint16 }
type GenS9 struct { F0 int32; F1 uint64 }
type GenS10 struct { F0 uint64; F1 uint32 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int32; F1 int; F2 uint }
type GenS13 struct { F0 int }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int; F1 int }
type GenS17 struct { F0 int8; F1 uint8; F2 int8 }
type GenS18 struct { F0 uint32; F1 uint16; F2 uint64 }
func genFn0(p0 int16, p1 int8, p2 bool) string {
    v7 := GenS1{F0: 255, F1: (true && (!false)), F2: 70}
    v8 := GenS2{F0: 2, F1: -1}
    v9 := GenS1{F0: 42, F1: p2, F2: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return ((("hello" + ("a" + ("" + ""))) + "a") + "a")
}
func genFn1() string {
    var v10 bool = ((0 == 42) && false)
    v11 := GenS1{F0: 917, F1: v10, F2: 4294967295}
    v12 := []int{128, 464}
    _ = v10
    _ = v11
    _ = v12
    return ""
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 776
    v3 := make(map[int]string)
    var v4 uint = 65535
    v5 := make(map[uint32]string)
    var v6 int64 = 42
    if (false && false) {
        switch 0 {
        case 128:
            v13 := []int{-399, 0, int(v2), -1}
            _ = v13
        default:
            v14 := GenS2{F0: (v2 & (v2 / ((v2) | 1))), F1: 1}
            _ = v14
        }
        v6 = 1
        switch (("hello" + "\n\t") + ("" + "")) {
        case "foo bar":
            v15 := GenS3{F0: 3, F1: 32767}
            _ = v15
        case "hello":
            v1["hello"] = false
        case "\n\t":
            v16 := []int{}
            _ = v16
        }
    } else {
        v17 := []int{32767, 317, -976}
        if false {
            v1[""] = false
            v4 = (v2 & (v2 % ((538) | 1)))
            delete(v1, (("\n\t" + "hello") + "hello"))
        } else {
            if false {
                fmt.Println(v6)
                for fi18 := 0; fi18 < 14; fi18++ {
                    if len(v17) > 0 {
                        v17[0] = (fi18 | (fi18 << (((v4 - v2)) & 63)))
                    }
                    var v19 uint64 = 1
                    _ = fi18
                    _ = v19
                }
                v6 += -1
            } else {
                fmt.Println(v4, v6, v4)
                v4 += v4
            }
            v4 = (v4 >> (((v2 + (v4 % ((v4) | 1)))) & 63))
        }
        fmt.Println(v2)
        _ = v17
    }
    v20 := []int{9223372036854775807, 32768}
    v2 = v2
    fmt.Println(v2, v2)
    switch "foo bar" {
    case "hello":
        var v21 int32 = -2147483648
        _ = v21
    }
    v4 = (v4 - (v4 | 65535))
    if true {
        var v22 int = 89
        var v23 int = (v22 / (((v22 ^ (-(v22)))) | 1))
        _ = v22
        _ = v23
    } else {
        if (!true) {
            var v24 uint64 = 673
            if true {
                for fi25 := 0; fi25 < 19; fi25++ {
                    v4 -= (v2 + (v4 >> ((0) & 63)))
                    v6 -= 0
                    for fi26 := 0; fi26 < 18; fi26++ {
                        var v27 uint8 = 255
                        fmt.Println(fi26)
                        v27 -= v27
                        _ = fi26
                        _ = v27
                    }
                    _ = fi25
                }
                _ = len(v5)
                fmt.Println(v6, v24)
            } else {
                _ = len(v20)
            }
            if true {
                for fi28 := 0; fi28 < 7; fi28++ {
                    v20 = append(v20, fi28)
                    if (-1 > 0) {
                        if ((!(((float32(1e10) * (-2.25 * 1.5)) - 1e10) >= ((1.5 - 1e10) * -2.25))) && true) {
                            for fi29 := 0; fi29 < 8; fi29++ {
                                v4 = v2
                                _ = fi29
                            }
                            fmt.Println(v24, v6, v24)
                        } else {
                            var v30 uint8 = 69
                            _ = v30
                        }
                        var v31 uint16 = 32767
                        v20 = append(v20, fi28)
                        _ = v31
                    } else {
                        v2 += v4
                        v24 = 192
                        fmt.Println(fi28, v6)
                    }
                    _ = fi28
                }
                for fi32 := 0; fi32 < 12; fi32++ {
                    if false {
                        switch ("foo bar" + "") {
                        case "":
                            v4 -= (v2 % ((uint(v24)) | 1))
                        case "foo bar":
                            if (v24 < v24) {
                                v33 := GenS1{F0: v4, F1: true, F2: uint64(v4)}
                                _ = v33
                            } else {
                                for fi34 := 0; fi34 < 17; fi34++ {
                                    if true {
                                        var v35 string = "a"
                                        _ = v35
                                    } else {
                                        v4 = (^(v4))
                                        v6 = (v6 << ((18446744073709551615) & 63))
                                    }
                                    v2 -= (v2 - v4)
                                    _ = fi34
                                }
                                v4 += 255
                                for fi36 := 0; fi36 < 3; fi36++ {
                                    v37 := GenS3{F0: 647, F1: 1}
                                    v38 := GenS4{F0: 65535, F1: 32767}
                                    _ = fi36
                                    _ = v37
                                    _ = v38
                                }
                            }
                        default:
                            switch "" {
                            case "foo bar":
                                for fi39 := 0; fi39 < 1; fi39++ {
                                    switch fi39 {
                                    case 569:
                                        var v40 int = fi32
                                        _ = v40
                                    case 4294967295:
                                        var v41 float64 = 1e10
                                        _ = v41
                                    case 617:
                                        delete(v5, 42)
                                    default:
                                        var v42 bool = true
                                        _ = v42
                                    }
                                    var v43 uint8 = 127
                                    var v44 uint64 = 435
                                    _ = fi39
                                    _ = v43
                                    _ = v44
                                }
                            default:
                                v4 += 4294967295
                            }
                        }
                        fmt.Println(v2, v6, v24)
                    }
                    _ = fi32
                }
                fmt.Println(v4)
            } else {
                fmt.Println(v6, v24)
                if true {
                    if false {
                        v45 := GenS5{F0: 2147483647, F1: v24, F2: 2}
                        v46 := GenS6{F0: (1.5 >= float32(-2.25)), F1: "", F2: 2147483647}
                        _ = v45
                        _ = v46
                    }
                } else {
                    for fi47 := 0; fi47 < 13; fi47++ {
                        v48 := GenS7{F0: 127, F1: 213, F2: 32767}
                        _ = fi47
                        _ = v48
                    }
                    v6 -= v6
                }
                _ = len(v5)
            }
            _ = v24
        } else {
            _ = len(v3)
        }
        v2 -= v4
        v49 := GenS2{F0: (v4 * (v2 | v2)), F1: -32768}
        _ = v49
    }
    fmt.Println(v4, v2)
    if (!false) {
        var v50 uint8 = 255
        _ = v50
    } else {
        v4 -= v4
        delete(v5, 2147483647)
        v6 += v6
    }
    var v51 int8 = -128
    if false {
        var v52 uint32 = 903
        var v53 float64 = (1e10 - 1.5)
        switch 32767 {
        case 1:
            var v54 float32 = 1e10
            _ = v54
        case 32768:
            if true {
                fmt.Println(v51)
                for fi55 := 0; fi55 < 18; fi55++ {
                    var v56 string = "hello"
                    fmt.Println(v51, fi55)
                    _ = fi55
                    _ = v56
                }
                if (!(("a" >= "") || true)) {
                    switch "hello" {
                    case "a":
                        var v57 uint32 = 1
                        _ = v57
                    case "":
                        if (false || (true || (42 >= 128))) {
                            fmt.Println(v6, v2)
                            var v58 bool = (32767 <= 2147483647)
                            for fi59 := 0; fi59 < 19; fi59++ {
                                if (v58 && v58) {
                                    v51 += 0
                                    if v58 {
                                        var v60 float32 = 1.5
                                        _ = v60
                                    } else {
                                        switch fi59 {
                                        case -510:
                                            if v58 {
                                                for fi61 := 0; fi61 < 9; fi61++ {
                                                    var v62 uint = (v2 & v4)
                                                    _ = fi61
                                                    _ = v62
                                                }
                                                v20 = append(v20, 325)
                                            } else {
                                                v52 = (^(v52))
                                                if (1 > 2) {
                                                    fmt.Println(v52, v2)
                                                    var v63 string = ""
                                                    _ = v63
                                                }
                                            }
                                        default:
                                            if (!(2 == 65535)) {
                                                v6 = v6
                                                switch "" {
                                                case "a":
                                                    v4 += (v2 - (v2 % (((^(v2))) | 1)))
                                                default:
                                                    var v64 int64 = -1
                                                    _ = v64
                                                }
                                            }
                                        }
                                    }
                                }
                                _ = fi59
                            }
                            _ = v58
                        }
                    case "\n\t":
                        if len(v20) > 0 {
                            v20[0] = 32767
                        }
                    }
                    fmt.Println(v6)
                    _ = len(v3)
                } else {
                    var v65 int = -517
                    if (1000 == -1) {
                        _ = len(v1)
                        for fi66 := 0; fi66 < 7; fi66++ {
                            v52 = 42
                            for fi67 := 0; fi67 < 12; fi67++ {
                                switch "foo bar" {
                                case "a":
                                    if len(v20) > 0 {
                                        v20[0] = fi66
                                    }
                                case "foo bar":
                                    v52 -= (v52 * 1)
                                default:
                                    if false {
                                        switch "\n\t" {
                                        case "\n\t":
                                            if (347 != 4294967295) {
                                                v68 := v5[v52]
                                                _ = v68
                                            } else {
                                                var v69 uint64 = 0
                                                if true {
                                                    var v70 int8 = 100
                                                    delete(v3, v65)
                                                    _ = v70
                                                } else {
                                                    fmt.Println(v52, v6)
                                                    if (!true) {
                                                        switch fi66 {
                                                        case 146:
                                                            switch ("hello" + "foo bar") {
                                                            case "foo bar":
                                                                fmt.Println(fi67, v69)
                                                            case "":
                                                                var v71 int16 = 0
                                                                _ = v71
                                                            }
                                                        case -17:
                                                            v65 = fi66
                                                        }
                                                    } else {
                                                        v53 -= v53
                                                        var v72 string = ("" + "hello")
                                                        var v73 int32 = 42
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                    v2 -= v4
                                                }
                                                var v74 int32 = 1
                                                _ = v69
                                                _ = v74
                                            }
                                        }
                                        if false {
                                            v75 := GenS8{F0: 2147483647, F1: "\n\t", F2: 398}
                                            _ = v75
                                        } else {
                                            var v76 int64 = 9223372036854775807
                                            v3[(fi66 << (((v4 | 65535)) & 63))] = ""
                                            if false {
                                                v51 = 74
                                                if len(v20) > 0 {
                                                    v20[0] = (fi66 << ((v2) & 63))
                                                }
                                            } else {
                                                switch "\n\t" {
                                                case "a":
                                                    var v77 int16 = -787
                                                    _ = v77
                                                case "hello":
                                                    switch "\n\t" {
                                                    case "a":
                                                        v78 := GenS9{F0: -2147483648, F1: 18446744073709551615}
                                                        var v79 int8 = v51
                                                        fmt.Println(fi67, v51)
                                                        _ = v78
                                                        _ = v79
                                                    case "":
                                                        if true {
                                                            var v80 uint8 = 42
                                                            v81 := GenS10{F0: 0, F1: uint32(v2)}
                                                            var v82 float32 = 0.0
                                                            var v83 string = "foo bar"
                                                            v5[(^(v52))] = ("a" + (v83 + v83))
                                                            _ = v80
                                                            _ = v81
                                                            _ = v82
                                                            _ = v83
                                                        }
                                                    default:
                                                        if (!false) {
                                                            v2 = v4
                                                            var v84 uint32 = 1
                                                            if (("" + ("" + ("a" + ("a" + (("a" + "hello") + "a"))))) != "") {
                                                                switch (fi66 >> (((^(v4))) & 63)) {
                                                                case -9223372036854775808:
                                                                    switch fi66 {
                                                                    case 32767:
                                                                        delete(v3, (fi67 * (fi67 & (fi67 % (((fi66 | v65)) | 1)))))
                                                                    case -604:
                                                                        v85 := GenS11{F0: 149}
                                                                        _ = v85
                                                                    }
                                                                case 42:
                                                                    if (!true) {
                                                                        if (42 > 1) {
                                                                            var v86 uint64 = 873
                                                                            var v87 int64 = v76
                                                                            v2 = (v4 >> (((v2 << (((v2 / (((^(v4))) | 1))) & 63))) & 63))
                                                                            _ = v86
                                                                            _ = v87
                                                                        }
                                                                        fmt.Println(v76, v53, fi66)
                                                                        v88 := GenS12{F0: 196, F1: fi67, F2: (^(v4))}
                                                                        _ = v88
                                                                    } else {
                                                                        fmt.Println(v76)
                                                                        fmt.Println(fi66)
                                                                    }
                                                                }
                                                                switch v65 {
                                                                case -33:
                                                                    if true {
                                                                        if false {
                                                                            var v89 float64 = v53
                                                                            _ = v89
                                                                        } else {
                                                                            if (false && true) {
                                                                                switch 9223372036854775807 {
                                                                                case -853:
                                                                                    var v90 int = (v65 ^ 185)
                                                                                    _ = v90
                                                                                case 4294967295:
                                                                                    if len(v20) >= 2 {
                                                                                        v20 = v20[1:len(v20)]
                                                                                    }
                                                                                    _ = len(v20)
                                                                                    v20 = append(v20, v65)
                                                                                case -9223372036854775808:
                                                                                    v51 -= (v51 & v51)
                                                                                }
                                                                                fmt.Println(v76, v53, v4)
                                                                            } else {
                                                                                _ = len(v1)
                                                                                if (v53 > 1.5) {
                                                                                    v4 = (v2 * (v4 + v2))
                                                                                    fmt.Println(v6)
                                                                                    var v91 string = "a"
                                                                                    var v92 string = (v91 + (v91 + "a"))
                                                                                    fmt.Println(v2)
                                                                                    _ = v91
                                                                                    _ = v92
                                                                                }
                                                                                var v93 int64 = v76
                                                                                switch (fi66 * (fi66 - (v65 / (((fi67 / ((v65) | 1))) | 1)))) {
                                                                                case 2147483647:
                                                                                    if (v6 <= v76) {
                                                                                        v84 = 1
                                                                                        var v94 int32 = 1
                                                                                        var v95 int16 = -1
                                                                                        switch "\n\t" {
                                                                                        case "hello":
                                                                                            fmt.Println(v6, v76, v2)
                                                                                        case "foo bar":
                                                                                            switch 32768 {
                                                                                            case 0:
                                                                                                switch (("hello" + "a") + ("foo bar" + ("" + ""))) {
                                                                                                case "hello":
                                                                                                    var v96 uint64 = 4294967295
                                                                                                    _ = v96
                                                                                                case "":
                                                                                                    if len(v20) > 0 {
                                                                                                        v20[0] = fi67
                                                                                                    }
                                                                                                case "a":
                                                                                                    if false {
                                                                                                        switch "" {
                                                                                                        case "hello":
                                                                                                            var v97 int64 = (v6 ^ (v93 + 71))
                                                                                                            _ = v97
                                                                                                        case "":
                                                                                                            var v98 string = (("a" + (("a" + "") + ("hello" + "a"))) + "")
                                                                                                            _ = v98
                                                                                                        }
                                                                                                        var v99 uint8 = 218
                                                                                                        fmt.Println(v51, v2, v65)
                                                                                                        _ = v99
                                                                                                    } else {
                                                                                                        switch "hello" {
                                                                                                        case "":
                                                                                                            v94 = 42
                                                                                                        case "\n\t":
                                                                                                            fmt.Println(v53, v6)
                                                                                                        case "hello":
                                                                                                            fmt.Println(v65, v84, v94)
                                                                                                        }
                                                                                                        v6 = (v6 * (v6 + v76))
                                                                                                        switch "foo bar" {
                                                                                                        case "":
                                                                                                            if len(v20) >= 2 {
                                                                                                                v20 = v20[1:len(v20)]
                                                                                                            }
                                                                                                        }
                                                                                                        v100 := GenS2{F0: v4, F1: v95}
                                                                                                        v5[4294967295] = "\n\t"
                                                                                                        _ = v100
                                                                                                    }
                                                                                                }
                                                                                            case 32767:
                                                                                                v52 -= 987
                                                                                            default:
                                                                                                fmt.Println(v2, v65)
                                                                                            }
                                                                                        case "\n\t":
                                                                                            if false {
                                                                                                if false {
                                                                                                    _ = len(v20)
                                                                                                    v101 := v1["a"]
                                                                                                    switch (fi66 * int(v95)) {
                                                                                                    case -34:
                                                                                                        var v102 uint = (v4 + v2)
                                                                                                        _ = v102
                                                                                                    case 127:
                                                                                                        fmt.Println(v53)
                                                                                                    case 32767:
                                                                                                        switch v65 {
                                                                                                        case 448:
                                                                                                            v1["hello"] = true
                                                                                                        case 998:
                                                                                                            if v101 {
                                                                                                                if len(v20) >= 2 {
                                                                                                                    v20 = v20[1:len(v20)]
                                                                                                                }
                                                                                                                var v103 uint32 = (v84 >> ((v4) & 63))
                                                                                                                _ = v103
                                                                                                            }
                                                                                                        default:
                                                                                                            v104 := GenS7{F0: 127, F1: 42, F2: v95}
                                                                                                            var v105 int32 = (v94 | -1)
                                                                                                            _ = len(v5)
                                                                                                            _ = v104
                                                                                                            _ = v105
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v101
                                                                                                }
                                                                                            } else {
                                                                                                var v106 uint8 = 225
                                                                                                var v107 uint = 65535
                                                                                                v76 -= (^(v6))
                                                                                                _ = v106
                                                                                                _ = v107
                                                                                            }
                                                                                        }
                                                                                        v108 := GenS11{F0: 235}
                                                                                        _ = v108
                                                                                        _ = v94
                                                                                        _ = v95
                                                                                    }
                                                                                case 379:
                                                                                    if (0 < 0) {
                                                                                        v93 -= (v76 << ((v4) & 63))
                                                                                        v109 := GenS13{F0: (fi67 - 9223372036854775807)}
                                                                                        _ = v109
                                                                                    }
                                                                                case 456:
                                                                                    switch "\n\t" {
                                                                                    case "a":
                                                                                        v110 := GenS5{F0: fi67, F1: 18446744073709551615, F2: uint8(v52)}
                                                                                        _ = v110
                                                                                    case "":
                                                                                        var v111 string = ""
                                                                                        _ = v111
                                                                                    case "hello":
                                                                                        var v112 uint16 = 1
                                                                                        _ = v112
                                                                                    }
                                                                                default:
                                                                                    v52 -= (^(v52))
                                                                                }
                                                                                fmt.Println(v4, v2, v84)
                                                                                _ = v93
                                                                            }
                                                                            v20 = append(v20, fi67)
                                                                            var v113 uint64 = 42
                                                                            _ = v113
                                                                        }
                                                                    } else {
                                                                        v53 = v53
                                                                        delete(v3, 127)
                                                                        v76 -= (v6 * v6)
                                                                        fmt.Println(v84, fi67, fi66)
                                                                        if true {
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                switch "foo bar" {
                                                                                case "":
                                                                                    if (0 <= 530) {
                                                                                        var v114 int32 = 0
                                                                                        fmt.Println(v84, v51)
                                                                                        _ = v114
                                                                                    } else {
                                                                                        _ = len(v1)
                                                                                        switch -263 {
                                                                                        case -426:
                                                                                            var v115 int64 = v6
                                                                                            _ = v115
                                                                                        case 127:
                                                                                            var v116 uint = v4
                                                                                            _ = v116
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    if (127 < 25) {
                                                                                        _ = len(v20)
                                                                                        var v117 float64 = (float64(3.14159) + ((v53 / v53) * ((float64(1e10) * v53) - v53)))
                                                                                        fmt.Println(v52)
                                                                                        _ = v117
                                                                                    } else {
                                                                                        if len(v20) >= 2 {
                                                                                            v20 = v20[1:len(v20)]
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case "":
                                                                                var v118 int = 9223372036854775807
                                                                                _ = v118
                                                                            default:
                                                                                var v119 float32 = float32(3.14159)
                                                                                var v120 int8 = v51
                                                                                fmt.Println(v76)
                                                                                _ = v119
                                                                                _ = v120
                                                                            }
                                                                            v76 = v76
                                                                            var v121 float32 = 1e10
                                                                            _ = v121
                                                                        } else {
                                                                            v4 += v2
                                                                        }
                                                                        v20 = append(v20, 852)
                                                                    }
                                                                }
                                                                var v122 int8 = (^(v51))
                                                                _ = len(v20)
                                                                _ = v122
                                                            } else {
                                                                v6 -= (v76 + (v6 * v6))
                                                                var v123 uint32 = (v84 << ((v4) & 63))
                                                                _ = v123
                                                            }
                                                            fmt.Println(v2)
                                                            _ = v84
                                                        } else {
                                                            v4 = 255
                                                            var v124 int64 = v6
                                                            _ = v124
                                                        }
                                                    }
                                                    _ = len(v20)
                                                    var v125 int8 = -128
                                                    if len(v20) > 0 {
                                                        v20[0] = (fi67 % (((fi66 % ((fi67) | 1))) | 1))
                                                    }
                                                    _ = v125
                                                case "\n\t":
                                                    v126 := GenS1{F0: v2, F1: false, F2: 160}
                                                    _ = v126
                                                }
                                            }
                                            var v127 uint64 = 18446744073709551615
                                            _ = v127
                                            _ = v76
                                        }
                                    } else {
                                        v52 = v52
                                    }
                                }
                                fmt.Println(v6, v6)
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                                var v128 float32 = ((0.0 / (((3.14159 + (1e10 * ((float32(v53) / 1.5) - (1e10 * 0.1)))) + -2.25) / (0.1 / (((((1.5 + (-2.25 * (-2.25 + 0.1))) * 3.14159) - -2.25) / (3.14159 * (0.1 - (3.14159 / 0.1)))) * 1.5)))) / 1.0)
                                _ = fi67
                                _ = v128
                            }
                            delete(v1, (("" + ((("foo bar" + ("a" + "\n\t")) + "hello") + "hello")) + ("a" + "a")))
                            _ = fi66
                        }
                    } else {
                        v52 = uint32(v4)
                        v65 = (v65 % ((128) | 1))
                        fmt.Println(v2, v53)
                    }
                    if ((v6 * v6) < (v6 / ((v6) | 1))) {
                        for fi129 := 0; fi129 < 11; fi129++ {
                            if (!(true || true)) {
                                switch (v65 >> (((v2 << ((v2) & 63))) & 63)) {
                                case 127:
                                    fmt.Println(v6, v65, v53)
                                case 128:
                                    v3[4294967295] = (("a" + "hello") + "")
                                default:
                                    switch "\n\t" {
                                    case "a":
                                        var v130 uint8 = 255
                                        _ = v130
                                    case "hello":
                                        if true {
                                            fmt.Println(v65, fi129, v51)
                                            v65 += fi129
                                            v2 -= (v4 | v2)
                                        } else {
                                            var v131 uint64 = 2
                                            _ = v131
                                        }
                                    case "foo bar":
                                        var v132 float32 = (0.1 / 0.1)
                                        _ = v132
                                    default:
                                        v5[v52] = ((("hello" + ("" + ((("foo bar" + "") + "\n\t") + "hello"))) + "\n\t") + "\n\t")
                                    }
                                }
                                for fi133 := 0; fi133 < 17; fi133++ {
                                    var v134 uint16 = 905
                                    _ = fi133
                                    _ = v134
                                }
                                switch "" {
                                case "a":
                                    var v135 int16 = 32767
                                    _ = v135
                                }
                            } else {
                                v136 := GenS14{F0: 42}
                                v137 := GenS15{F0: 42}
                                _ = v136
                                _ = v137
                            }
                            v2 -= (v2 - (v2 + v4))
                            fmt.Println(v52)
                            _ = fi129
                        }
                    }
                    _ = v65
                }
            }
        }
        _ = v52
        _ = v53
    } else {
        if false {
            for fi138 := 0; fi138 < 19; fi138++ {
                v2 = (^(v2))
                v51 = (v51 * 68)
                _ = fi138
            }
            for fi139 := 0; fi139 < 12; fi139++ {
                if (((v51 <= v51) && false) || true) {
                    v4 = 1
                } else {
                    for fi140 := 0; fi140 < 1; fi140++ {
                        v6 = v6
                        _ = fi140
                    }
                    delete(v5, 106)
                }
                switch ("a" + "foo bar") {
                case "\n\t":
                    if false {
                        v51 = 100
                    }
                case "hello":
                    switch fi139 {
                    case 32768:
                        for fi141 := 0; fi141 < 12; fi141++ {
                            if (v6 < v6) {
                                var v142 uint64 = 4294967295
                                v1[""] = (-1 <= 2147483647)
                                v143 := GenS7{F0: (fi139 - fi139), F1: 65535, F2: -32768}
                                if (2147483647 < 2147483647) {
                                    switch "hello" {
                                    case "a":
                                        fmt.Println(v2, fi141, v4)
                                    case "foo bar":
                                        switch (^(fi139)) {
                                        case -2147483648:
                                            v142 -= (v142 | (v142 | (v142 << ((v4) & 63))))
                                        default:
                                            var v144 bool = true
                                            _ = v144
                                        }
                                    case "\n\t":
                                        var v145 bool = false
                                        _ = v145
                                    }
                                }
                                _ = v142
                                _ = v143
                            }
                            v51 -= v51
                            fmt.Println(v2)
                            _ = fi141
                        }
                    case 2147483647:
                        var v146 float32 = 0.1
                        _ = v146
                    default:
                        switch fi139 {
                        case 879:
                            for fi147 := 0; fi147 < 1; fi147++ {
                                if true {
                                    v2 = (v4 << (((v2 << ((v4) & 63))) & 63))
                                    delete(v5, 0)
                                } else {
                                    if false {
                                        v4 = 48
                                        if ((0 < 552) && true) {
                                            switch fi147 {
                                            case 128:
                                                v148 := GenS14{F0: 4294967295}
                                                _ = v148
                                            case -925:
                                                var v149 bool = true
                                                _ = v149
                                            default:
                                                v3[32767] = "\n\t"
                                            }
                                        }
                                    }
                                }
                                var v150 uint8 = 102
                                _ = fi147
                                _ = v150
                            }
                        case 9223372036854775807:
                            v151 := GenS16{F0: fi139, F1: (fi139 - fi139)}
                            _ = v151
                        default:
                            _ = len(v20)
                        }
                    }
                case "foo bar":
                    var v152 uint8 = 15
                    _ = v152
                }
                v6 -= v6
                _ = fi139
            }
            v2 -= v4
        } else {
            if (!((326 > 32767) && true)) {
                var v153 int64 = (v6 ^ (-(v6)))
                var v154 uint8 = uint8(v2)
                v155 := GenS15{F0: 2}
                _ = v153
                _ = v154
                _ = v155
            } else {
                for fi156 := 0; fi156 < 14; fi156++ {
                    var v157 uint16 = 32767
                    v51 += 34
                    if ((v4 >> (((v4 * (v2 - (^(v2))))) & 63)) <= v4) {
                        if false {
                            v158 := GenS10{F0: 4294967295, F1: 4294967295}
                            var v159 int32 = 153
                            _ = len(v3)
                            _ = v158
                            _ = v159
                        }
                        v160 := GenS4{F0: (v157 | v157), F1: 1}
                        _ = v160
                    } else {
                        v2 += (v2 >> (((v2 & (v2 % ((16) | 1)))) & 63))
                        if (true && false) {
                            v4 -= (^(v2))
                        } else {
                            var v161 int8 = -121
                            for fi162 := 0; fi162 < 9; fi162++ {
                                v51 = 100
                                v157 += 65535
                                var v163 int = 32768
                                _ = fi162
                                _ = v163
                            }
                            _ = v161
                        }
                    }
                    _ = fi156
                    _ = v157
                }
            }
        }
        for fi164 := 0; fi164 < 14; fi164++ {
            v3[128] = ""
            switch "\n\t" {
            case "foo bar":
                if ("\n\t" > ("\n\t" + "a")) {
                    for fi165 := 0; fi165 < 16; fi165++ {
                        v2 = v4
                        v166 := GenS13{F0: (fi165 << ((v4) & 63))}
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                        v2 -= 255
                        var v167 string = "a"
                        v168 := v3[fi165]
                        v51 -= int8(v6)
                        if len(v20) > 0 {
                            v20[0] = (fi165 | (-(fi165)))
                        }
                        var v169 int64 = 42
                        var v170 string = ("a" + "foo bar")
                        var v171 int16 = -32768
                        _ = fi165
                        _ = v166
                        _ = v167
                        _ = v168
                        _ = v169
                        _ = v170
                        _ = v171
                    }
                }
            case "":
                v172 := GenS17{F0: (v51 - (v51 % ((v51) | 1))), F1: 82, F2: v51}
                _ = v172
            case "a":
                v6 -= v6
            }
            var v173 float32 = (0.1 * float32((1e10 / 1.0)))
            _ = fi164
            _ = v173
        }
        v174 := GenS18{F0: 2147483647, F1: 42, F2: 0}
        _ = v174
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
}
