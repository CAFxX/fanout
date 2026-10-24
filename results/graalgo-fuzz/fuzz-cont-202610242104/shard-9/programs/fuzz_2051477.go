package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 bool; F1 string; F2 int16 }
type GenS4 struct { F0 uint64; F1 int; F2 uint32 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 int; F1 int }
type GenS7 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint16; F1 int }
type GenS12 struct { F0 int32; F1 uint32 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint32; F1 int32 }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 uint; F1 uint32; F2 string }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 int; F1 int64; F2 uint32 }
type GenS19 struct { F0 int64; F1 int16 }
type GenS20 struct { F0 bool }
type GenS21 struct { F0 int; F1 uint32 }
type GenS22 struct { F0 int64; F1 uint16; F2 int64 }
type GenS23 struct { F0 string; F1 bool }
type GenS24 struct { F0 int64 }
type GenS25 struct { F0 int; F1 int64; F2 int8 }
type GenS26 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS27 struct { F0 uint64 }
type GenS28 struct { F0 int16; F1 string }
type GenS29 struct { F0 uint }
type GenS30 struct { F0 uint32; F1 uint8 }
type GenS31 struct { F0 int16 }
type GenS32 struct { F0 bool; F1 string }
type GenS33 struct { F0 uint }
type GenS34 struct { F0 uint16; F1 int8 }
type GenS35 struct { F0 uint8; F1 uint16 }
type GenS36 struct { F0 uint }
type GenS37 struct { F0 int }
func genFn0(p0 string, p1 uint8) uint8 {
    v7 := GenS1{F0: -128}
    fmt.Println(p1, p0)
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v7
    return 255
}
func genFn1(p0 string, p1 int32, p2 string) int8 {
    var v8 float64 = 0.0
    v9 := []int{32767, -423, -2147483648, 4294967295}
    var v10 int = 42
    if (0 != 42) {
        p1 = (p1 / (((-(p1))) | 1))
        v11 := GenS1{F0: -20}
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 127
}

func main() {
    var v1 uint8 = 81
    v2 := []int{2147483647, 42, -9223372036854775808}
    v3 := make(map[int]string)
    var v4 int8 = -76
    var v5 bool = (!true)
    var v6 int64 = 1
    v12 := GenS1{F0: genFn1("foo bar", 0, "foo bar")}
    for fi13 := 0; fi13 < 20; fi13++ {
        switch (fi13 / ((-1) | 1)) {
        case 1:
            v4 = (v4 % ((v4) | 1))
        default:
            for fi14 := 0; fi14 < 14; fi14++ {
                fmt.Println(fi13, v4)
                v5 = false
                switch "foo bar" {
                case "foo bar":
                    switch fi13 {
                    case 12:
                        switch ("a" + "a") {
                        case "\n\t":
                            switch (fi14 % (((fi13 ^ (fi14 + (fi14 / (((fi13 ^ (-(fi13)))) | 1))))) | 1)) {
                            case 42:
                                if v5 {
                                    var v15 uint = 2
                                    var v16 float32 = (3.14159 - (0.0 / -2.25))
                                    _ = v15
                                    _ = v16
                                }
                            case 653:
                                v17 := GenS1{F0: v4}
                                _ = v17
                            default:
                                if true {
                                    var v18 string = "\n\t"
                                    _ = v18
                                } else {
                                    v19 := GenS1{F0: (v4 ^ int8(v1))}
                                    switch -196 {
                                    case -641:
                                        if (42 == -1) {
                                            if false {
                                                v20 := GenS1{F0: -55}
                                                switch fi13 {
                                                case 879:
                                                    fmt.Println(v5)
                                                default:
                                                    delete(v3, -9223372036854775808)
                                                }
                                                if (1000 > 42) {
                                                    switch fi13 {
                                                    case -9223372036854775808:
                                                        var v21 int8 = genFn1("hello", -2147483648, "\n\t")
                                                        _ = v21
                                                    case 328:
                                                        fmt.Println(v1)
                                                        if v5 {
                                                            v3[(fi13 - 705)] = ("foo bar" + ("hello" + "hello"))
                                                        }
                                                    default:
                                                        switch "a" {
                                                        case "a":
                                                            fmt.Println(v4, fi13, fi13)
                                                        case "hello":
                                                            v5 = ((v4 ^ (v4 & v4)) < 1)
                                                        }
                                                    }
                                                    switch fi13 {
                                                    case 769:
                                                        fmt.Println(v4, fi14, fi13)
                                                    default:
                                                        v6 = v6
                                                    }
                                                    v22 := GenS1{F0: v4}
                                                    fmt.Println(v5, v1)
                                                    _ = v22
                                                }
                                                var v23 string = "hello"
                                                if (v5 && v5) {
                                                    var v24 string = "\n\t"
                                                    var v25 int16 = 1000
                                                    fmt.Println(v1, fi14, v5)
                                                    _ = v24
                                                    _ = v25
                                                } else {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    var v26 float32 = (0.0 - 0.0)
                                                    v23 = ""
                                                    _ = v26
                                                }
                                                _ = v20
                                                _ = v23
                                            }
                                            var v27 int32 = 0
                                            var v28 float32 = 3.14159
                                            _ = v27
                                            _ = v28
                                        } else {
                                            v4 += (v4 << ((65535) & 63))
                                            v29 := GenS1{F0: v4}
                                            if (!((v6 & (v6 / (((v6 >> ((280) & 63))) | 1))) != (v6 - (v6 >> ((18446744073709551615) & 63))))) {
                                                v6 -= v6
                                                var v30 uint = 317
                                                _ = v30
                                            }
                                            _ = v29
                                        }
                                    case 4294967295:
                                        var v31 int = fi13
                                        _ = v31
                                    default:
                                        v32 := GenS1{F0: genFn1(((("\n\t" + "foo bar") + "") + "\n\t"), 1, ("hello" + "foo bar"))}
                                        _ = v32
                                    }
                                    delete(v3, fi14)
                                    _ = v19
                                }
                            }
                        default:
                            switch (fi14 * 32768) {
                            case 2147483647:
                                var v33 float64 = 3.14159
                                _ = v33
                            case -876:
                                if ((-683 == -32768) || v5) {
                                    v34 := GenS1{F0: v4}
                                    var v35 bool = true
                                    _ = v34
                                    _ = v35
                                } else {
                                    if true {
                                        v4 -= v4
                                        var v36 int8 = v4
                                        _ = v36
                                    }
                                    switch "" {
                                    case "\n\t":
                                        var v37 uint32 = 367
                                        _ = v37
                                    }
                                    v6 -= -9223372036854775808
                                    if (v5 && false) {
                                        v5 = ((4294967295 > 207) || ("" == "foo bar"))
                                        if ((v5 || (117 < v4)) || true) {
                                            v6 = v6
                                            v5 = ((fi14 >= (^(fi13))) && v5)
                                            var v38 uint = 961
                                            _ = v38
                                        }
                                    } else {
                                        var v39 float64 = 0.1
                                        if (4294967295 <= 42) {
                                            switch "hello" {
                                            case "a":
                                                var v40 float64 = v39
                                                _ = v40
                                            case "foo bar":
                                                var v41 float64 = v39
                                                _ = v41
                                            case "":
                                                v42 := GenS1{F0: v4}
                                                _ = v42
                                            default:
                                                _ = len(v3)
                                            }
                                        } else {
                                            v43 := GenS1{F0: (v4 ^ v4)}
                                            var v44 uint32 = 1
                                            v4 = (v4 % (((v4 | genFn1("\n\t", 0, (("a" + "") + "foo bar")))) | 1))
                                            _ = v43
                                            _ = v44
                                        }
                                        var v45 float64 = (1.5 * v39)
                                        fmt.Println(v39, v4, fi14)
                                        fmt.Println(fi14, v5)
                                        _ = v39
                                        _ = v45
                                    }
                                    fmt.Println(v1, fi14, v1)
                                    if (v5 && (v5 || ((1e10 < 0.0) && v5))) {
                                        v4 -= (v4 & (v4 - 1))
                                        v6 += 643
                                        v46 := GenS1{F0: (^(v4))}
                                        _ = v46
                                    } else {
                                        if (v5 && v5) {
                                            v47 := GenS1{F0: genFn1("a", 42, "\n\t")}
                                            switch fi13 {
                                            case 9223372036854775807:
                                                switch 765 {
                                                case 128:
                                                    if true {
                                                        if v5 {
                                                            v1 = 146
                                                            switch "" {
                                                            case "":
                                                                if (float32(3.14159) >= -2.25) {
                                                                    var v48 int16 = 1
                                                                    v49 := GenS1{F0: v4}
                                                                    fmt.Println(v6, v1)
                                                                    v4 = v4
                                                                    _ = v48
                                                                    _ = v49
                                                                } else {
                                                                    v5 = v5
                                                                    switch fi13 {
                                                                    case -2147483648:
                                                                        v4 = genFn1("hello", 0, "")
                                                                    default:
                                                                        if len(v2) > 0 {
                                                                            v2[0] = (fi14 << ((92) & 63))
                                                                        }
                                                                    }
                                                                    if true {
                                                                        switch "" {
                                                                        case "hello":
                                                                            var v50 uint64 = 1
                                                                            _ = v50
                                                                        case "\n\t":
                                                                            var v51 bool = false
                                                                            _ = v51
                                                                        }
                                                                        delete(v3, -2147483648)
                                                                        if v5 {
                                                                            if (237 == 42) {
                                                                                v52 := GenS2{F0: 42, F1: v6}
                                                                                var v53 uint = 2
                                                                                _ = v52
                                                                                _ = v53
                                                                            }
                                                                        }
                                                                    } else {
                                                                        var v54 int = (^(fi14))
                                                                        _ = v54
                                                                    }
                                                                    v5 = true
                                                                    v55 := GenS1{F0: genFn1(("a" + "hello"), -560, ("hello" + ("foo bar" + "\n\t")))}
                                                                    _ = v55
                                                                }
                                                                var v56 uint64 = 154
                                                                _ = v56
                                                            case "a":
                                                                switch "" {
                                                                case "hello":
                                                                    switch (fi13 / ((fi13) | 1)) {
                                                                    case -1:
                                                                        var v57 string = "a"
                                                                        _ = v57
                                                                    case 9223372036854775807:
                                                                        v4 = v4
                                                                    }
                                                                default:
                                                                    if (805 <= 398) {
                                                                        switch "a" {
                                                                        case "hello":
                                                                            fmt.Println(v5, fi14, fi13)
                                                                            fmt.Println(v1, v6, v6)
                                                                            v1 = genFn0("", (v1 + (v1 * 42)))
                                                                            var v58 float32 = (1e10 * (1e10 * ((3.14159 + (-2.25 / 1e10)) - 1.5)))
                                                                            if (0 == (v4 / ((genFn1(("" + "foo bar"), 2147483647, "hello")) | 1))) {
                                                                                var v59 float64 = (0.1 * 1e10)
                                                                                v1 -= (^(v1))
                                                                                v2 = append(v2, 658)
                                                                                _ = v59
                                                                            }
                                                                            _ = v58
                                                                        case "a":
                                                                            switch (fi14 << ((86) & 63)) {
                                                                            case 1:
                                                                                switch (("" + ("a" + "hello")) + "a") {
                                                                                case "\n\t":
                                                                                    if (42 == 1) {
                                                                                        delete(v3, fi14)
                                                                                        var v60 uint8 = 128
                                                                                        var v61 int32 = -2147483648
                                                                                        _ = v60
                                                                                        _ = v61
                                                                                    } else {
                                                                                        v1 += v1
                                                                                        fmt.Println(v1, fi14, v1)
                                                                                    }
                                                                                case "a":
                                                                                    fmt.Println(fi13, v1)
                                                                                default:
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                }
                                                                            case 371:
                                                                                if (fi13 > (fi14 | fi14)) {
                                                                                    delete(v3, fi13)
                                                                                    switch "\n\t" {
                                                                                    case "hello":
                                                                                        v62 := GenS3{F0: ((!(!v5)) || (int64(fi13) == (v6 % ((v6) | 1)))), F1: "a", F2: 1000}
                                                                                        _ = v62
                                                                                    case "a":
                                                                                        _ = len(v2)
                                                                                    case "":
                                                                                        var v63 uint16 = 1
                                                                                        _ = v63
                                                                                    default:
                                                                                        if v5 {
                                                                                            v1 += 1
                                                                                            if (v1 == genFn0("foo bar", v1)) {
                                                                                                if v5 {
                                                                                                    v1 = (v1 / ((0) | 1))
                                                                                                    if v5 {
                                                                                                        v64 := GenS4{F0: 4294967295, F1: (fi13 * 32767), F2: 4294967295}
                                                                                                        if (v5 && false) {
                                                                                                            if false {
                                                                                                                switch (fi14 / (((fi14 - (fi14 * (fi13 | -2147483648)))) | 1)) {
                                                                                                                case 42:
                                                                                                                    var v65 uint8 = (v1 - (v1 >> ((579) & 63)))
                                                                                                                    _ = v65
                                                                                                                case -714:
                                                                                                                    v1 += (v1 & v1)
                                                                                                                case -2147483648:
                                                                                                                    if ((255 > 0) && v5) {
                                                                                                                        var v66 int64 = 9223372036854775807
                                                                                                                        _ = v66
                                                                                                                    } else {
                                                                                                                        var v67 uint64 = 0
                                                                                                                        var v68 int32 = 42
                                                                                                                        v69 := GenS2{F0: 122, F1: -9223372036854775808}
                                                                                                                        if true {
                                                                                                                            switch fi14 {
                                                                                                                            case -1:
                                                                                                                                v5 = v5
                                                                                                                            case 32768:
                                                                                                                                fmt.Println(v68)
                                                                                                                            case 263:
                                                                                                                                var v70 int64 = int64(v68)
                                                                                                                                _ = v70
                                                                                                                            }
                                                                                                                            v71 := GenS5{F0: v68, F1: v68}
                                                                                                                            v72 := GenS6{F0: (fi14 ^ (fi13 >> ((0) & 63))), F1: -4}
                                                                                                                            _ = v71
                                                                                                                            _ = v72
                                                                                                                        } else {
                                                                                                                            v73 := GenS7{F0: 913, F1: (v1 + (v1 & v1)), F2: 65535}
                                                                                                                            v2 = append(v2, -2147483648)
                                                                                                                            v74 := GenS5{F0: v68, F1: v68}
                                                                                                                            _ = v73
                                                                                                                            _ = v74
                                                                                                                        }
                                                                                                                        if (0 <= 533) {
                                                                                                                            if (fi13 == fi14) {
                                                                                                                                var v75 uint32 = 2147483647
                                                                                                                                _ = v75
                                                                                                                            }
                                                                                                                            v4 = genFn1("a", (v68 & 2147483647), "a")
                                                                                                                            switch 1 {
                                                                                                                            case -1:
                                                                                                                                v76 := GenS8{F0: 4294967295}
                                                                                                                                _ = v76
                                                                                                                            case -152:
                                                                                                                                _ = len(v2)
                                                                                                                            default:
                                                                                                                                v68 -= (v68 | v68)
                                                                                                                            }
                                                                                                                            v6 -= v6
                                                                                                                            _ = len(v3)
                                                                                                                        }
                                                                                                                        _ = v67
                                                                                                                        _ = v68
                                                                                                                        _ = v69
                                                                                                                    }
                                                                                                                default:
                                                                                                                    if (255 >= 32767) {
                                                                                                                        if (!v5) {
                                                                                                                            v77 := GenS8{F0: 65535}
                                                                                                                            var v78 float32 = 0.1
                                                                                                                            v4 -= -13
                                                                                                                            _ = v77
                                                                                                                            _ = v78
                                                                                                                        } else {
                                                                                                                            var v79 float32 = (-2.25 * (((-2.25 * (-2.25 * (3.14159 / 1.0))) * (-2.25 / -2.25)) * -2.25))
                                                                                                                            _ = v79
                                                                                                                        }
                                                                                                                        fmt.Println(v4)
                                                                                                                        v4 -= v4
                                                                                                                        switch -1 {
                                                                                                                        case -32:
                                                                                                                            switch -786 {
                                                                                                                            case 0:
                                                                                                                                v4 -= (v4 >> ((255) & 63))
                                                                                                                            case 516:
                                                                                                                                var v80 int8 = (-(v4))
                                                                                                                                _ = v80
                                                                                                                            case 42:
                                                                                                                                v4 -= 1
                                                                                                                            }
                                                                                                                        case -2147483648:
                                                                                                                            v81 := GenS3{F0: v5, F1: (("" + (("a" + "a") + ("foo bar" + ("\n\t" + "a")))) + ("\n\t" + "hello")), F2: 1000}
                                                                                                                            _ = v81
                                                                                                                        }
                                                                                                                        if v5 {
                                                                                                                            if v5 {
                                                                                                                                delete(v3, 847)
                                                                                                                                v3[763] = "foo bar"
                                                                                                                            } else {
                                                                                                                                var v82 uint8 = 42
                                                                                                                                _ = v82
                                                                                                                            }
                                                                                                                            if (v1 < 127) {
                                                                                                                                var v83 bool = v5
                                                                                                                                switch "foo bar" {
                                                                                                                                case "foo bar":
                                                                                                                                    fmt.Println(v4, fi14)
                                                                                                                                    if len(v2) >= 2 {
                                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                                    }
                                                                                                                                case "a":
                                                                                                                                    if (((!true) || true) && (1e10 == 1.5)) {
                                                                                                                                        v84 := GenS2{F0: 2, F1: (v6 | v6)}
                                                                                                                                        v1 += (v1 << ((2) & 63))
                                                                                                                                        _ = v84
                                                                                                                                    }
                                                                                                                                case "":
                                                                                                                                    if v5 {
                                                                                                                                        switch (fi14 >> ((255) & 63)) {
                                                                                                                                        case 4294967295:
                                                                                                                                            v6 = -9223372036854775808
                                                                                                                                        case -2147483648:
                                                                                                                                            var v85 int8 = 42
                                                                                                                                            _ = v85
                                                                                                                                        default:
                                                                                                                                            v86 := GenS7{F0: 294, F1: (v1 >> ((255) & 63)), F2: 255}
                                                                                                                                            _ = v86
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        v1 -= v1
                                                                                                                                        _ = len(v2)
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                v5 = v83
                                                                                                                                var v87 uint8 = v1
                                                                                                                                if v5 {
                                                                                                                                    if (42 == 4294967295) {
                                                                                                                                        if v5 {
                                                                                                                                            var v88 int32 = 2147483647
                                                                                                                                            if (false && v83) {
                                                                                                                                                var v89 int64 = -976
                                                                                                                                                var v90 float64 = 0.1
                                                                                                                                                var v91 int8 = v4
                                                                                                                                                switch "foo bar" {
                                                                                                                                                case "\n\t":
                                                                                                                                                    v1 -= v87
                                                                                                                                                }
                                                                                                                                                var v92 int = fi14
                                                                                                                                                _ = v89
                                                                                                                                                _ = v90
                                                                                                                                                _ = v91
                                                                                                                                                _ = v92
                                                                                                                                            } else {
                                                                                                                                                v2 = append(v2, (^(fi14)))
                                                                                                                                            }
                                                                                                                                            if len(v2) > 0 {
                                                                                                                                                v2[0] = 4294967295
                                                                                                                                            }
                                                                                                                                            _ = v88
                                                                                                                                        } else {
                                                                                                                                            v93 := GenS9{F0: fi13}
                                                                                                                                            v1 += 42
                                                                                                                                            if len(v2) >= 2 {
                                                                                                                                                v2 = v2[1:len(v2)]
                                                                                                                                            }
                                                                                                                                            fmt.Println(v83, v83)
                                                                                                                                            var v94 int16 = -46
                                                                                                                                            _ = v93
                                                                                                                                            _ = v94
                                                                                                                                        }
                                                                                                                                        switch "a" {
                                                                                                                                        case "":
                                                                                                                                            v83 = v83
                                                                                                                                        default:
                                                                                                                                            delete(v3, -48)
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        v95 := GenS8{F0: 65535}
                                                                                                                                        _ = v95
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    if (1 < 2) {
                                                                                                                                        _ = len(v3)
                                                                                                                                        v96 := GenS2{F0: uint16(v1), F1: (v6 + 0)}
                                                                                                                                        switch "\n\t" {
                                                                                                                                        case "\n\t":
                                                                                                                                            v6 = 530
                                                                                                                                        case "":
                                                                                                                                            v3[(fi14 & fi13)] = "a"
                                                                                                                                        case "a":
                                                                                                                                            var v97 string = (("" + "\n\t") + "foo bar")
                                                                                                                                            _ = v97
                                                                                                                                        }
                                                                                                                                        _ = v96
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                _ = v83
                                                                                                                                _ = v87
                                                                                                                            } else {
                                                                                                                                v98 := v3[32767]
                                                                                                                                _ = v98
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                v99 := GenS10{F0: 65535}
                                                                                                                if v5 {
                                                                                                                    v4 -= v4
                                                                                                                } else {
                                                                                                                    v6 = (v6 | 728)
                                                                                                                    v6 -= v6
                                                                                                                }
                                                                                                                _ = v99
                                                                                                            }
                                                                                                            if true {
                                                                                                                switch fi13 {
                                                                                                                case -235:
                                                                                                                    v3[(fi13 << ((127) & 63))] = "\n\t"
                                                                                                                case -717:
                                                                                                                    if v5 {
                                                                                                                        var v100 uint8 = v1
                                                                                                                        v100 -= 127
                                                                                                                        _ = v100
                                                                                                                    }
                                                                                                                case 658:
                                                                                                                    v1 = v1
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v101 uint64 = 4294967295
                                                                                                                v102 := v3[-2147483648]
                                                                                                                var v103 int8 = (-(v4))
                                                                                                                _ = v101
                                                                                                                _ = v102
                                                                                                                _ = v103
                                                                                                            }
                                                                                                        }
                                                                                                        fmt.Println(fi14, fi13, v4)
                                                                                                        v5 = v5
                                                                                                        v104 := v3[(fi13 * fi14)]
                                                                                                        var v105 float64 = 3.14159
                                                                                                        if false {
                                                                                                            v105 -= float64((0.1 * (-2.25 * (0.0 / (1.5 - (3.14159 / float32(((((v105 * (float64(0.1) / float64(0.1))) - v105) + 1e10) * 3.14159))))))))
                                                                                                        }
                                                                                                        _ = v104
                                                                                                        _ = v105
                                                                                                        _ = v64
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            v2 = append(v2, 32768)
                                                                                                            v106 := GenS9{F0: -919}
                                                                                                            fmt.Println(fi13, v1, v1)
                                                                                                            var v107 uint8 = (^(v1))
                                                                                                            switch -1 {
                                                                                                            case -1:
                                                                                                                fmt.Println(fi13, v107)
                                                                                                            }
                                                                                                            _ = v106
                                                                                                            _ = v107
                                                                                                        } else {
                                                                                                            if len(v2) >= 2 {
                                                                                                                v2 = v2[1:len(v2)]
                                                                                                            }
                                                                                                            if false {
                                                                                                                if true {
                                                                                                                    var v108 uint32 = 2147483648
                                                                                                                    _ = v108
                                                                                                                } else {
                                                                                                                    v1 -= v1
                                                                                                                    if ((-(fi14)) != int(v6)) {
                                                                                                                        if len(v2) >= 2 {
                                                                                                                            v2 = v2[1:len(v2)]
                                                                                                                        }
                                                                                                                        if (1 > int16(v4)) {
                                                                                                                            if (!true) {
                                                                                                                                if v5 {
                                                                                                                                    if (42 < 1) {
                                                                                                                                        v109 := GenS2{F0: 65535, F1: v6}
                                                                                                                                        v110 := GenS11{F0: 255, F1: fi14}
                                                                                                                                        var v111 uint64 = 430
                                                                                                                                        _ = v109
                                                                                                                                        _ = v110
                                                                                                                                        _ = v111
                                                                                                                                    } else {
                                                                                                                                        var v112 uint = 206
                                                                                                                                        _ = v112
                                                                                                                                    }
                                                                                                                                    fmt.Println(v6, fi13, v4)
                                                                                                                                } else {
                                                                                                                                    v113 := GenS12{F0: -2147483648, F1: 2}
                                                                                                                                    v114 := GenS6{F0: (fi13 - (fi14 - fi14)), F1: 128}
                                                                                                                                    _ = v113
                                                                                                                                    _ = v114
                                                                                                                                }
                                                                                                                                if (!false) {
                                                                                                                                    var v115 uint = 65535
                                                                                                                                    _ = v115
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                if (-603 == 0) {
                                                                                                                                    var v116 uint16 = 879
                                                                                                                                    switch "foo bar" {
                                                                                                                                    case "foo bar":
                                                                                                                                        switch "" {
                                                                                                                                        case "\n\t":
                                                                                                                                            var v117 bool = (((v4 - genFn1("hello", 185, "a")) > (v4 >> ((42) & 63))) || v5)
                                                                                                                                            var v118 uint64 = 381
                                                                                                                                            if (0 != 448) {
                                                                                                                                                if ((v4 - v4) == v4) {
                                                                                                                                                    if (("hello" + "hello") >= (("hello" + ("foo bar" + "a")) + "\n\t")) {
                                                                                                                                                        switch "hello" {
                                                                                                                                                        case "":
                                                                                                                                                            var v119 int8 = 10
                                                                                                                                                            _ = v119
                                                                                                                                                        case "a":
                                                                                                                                                            v116 = v116
                                                                                                                                                        default:
                                                                                                                                                            if (v6 <= v6) {
                                                                                                                                                                if (float64(0.1) >= 1.5) {
                                                                                                                                                                    v1 -= (v1 ^ genFn0("\n\t", (v1 << ((972) & 63))))
                                                                                                                                                                    v116 -= uint16(v1)
                                                                                                                                                                }
                                                                                                                                                                var v120 string = ("" + "hello")
                                                                                                                                                                v2 = append(v2, fi14)
                                                                                                                                                                _ = v120
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        var v121 uint8 = (v1 >> ((69) & 63))
                                                                                                                                                        _ = v121
                                                                                                                                                    } else {
                                                                                                                                                        v117 = (1 >= (^(v116)))
                                                                                                                                                    }
                                                                                                                                                    _ = len(v3)
                                                                                                                                                    if v117 {
                                                                                                                                                        v122 := v3[fi13]
                                                                                                                                                        fmt.Println(v6, v1)
                                                                                                                                                        _ = v122
                                                                                                                                                    } else {
                                                                                                                                                        if true {
                                                                                                                                                            var v123 float64 = 0.0
                                                                                                                                                            if (v123 >= 1.5) {
                                                                                                                                                                var v124 string = "foo bar"
                                                                                                                                                                var v125 float32 = -2.25
                                                                                                                                                                var v126 int = fi14
                                                                                                                                                                v3[(^(fi13))] = ("" + ("a" + ((v124 + v124) + (v124 + (((((("hello" + (v124 + "\n\t")) + v124) + (((v124 + v124) + v124) + v124)) + "foo bar") + "\n\t") + ("" + ((((v124 + ("a" + v124)) + v124) + "foo bar") + v124)))))))
                                                                                                                                                                v117 = true
                                                                                                                                                                _ = v124
                                                                                                                                                                _ = v125
                                                                                                                                                                _ = v126
                                                                                                                                                            }
                                                                                                                                                            _ = v123
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            _ = v117
                                                                                                                                            _ = v118
                                                                                                                                        case "a":
                                                                                                                                            var v127 string = ""
                                                                                                                                            if v5 {
                                                                                                                                                v127 = v127
                                                                                                                                                var v128 int16 = 1000
                                                                                                                                                if ((!v5) && (!(v5 && v5))) {
                                                                                                                                                    if (!v5) {
                                                                                                                                                        var v129 uint64 = 42
                                                                                                                                                        v1 = (v1 << ((18446744073709551615) & 63))
                                                                                                                                                        var v130 uint = 1
                                                                                                                                                        _ = v129
                                                                                                                                                        _ = v130
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                                _ = v128
                                                                                                                                            }
                                                                                                                                            fmt.Println(fi13, fi13, fi13)
                                                                                                                                            _ = v127
                                                                                                                                        default:
                                                                                                                                            v4 = genFn1("\n\t", 42, ("" + "a"))
                                                                                                                                        }
                                                                                                                                    case "hello":
                                                                                                                                        v2 = append(v2, (fi14 & (-(fi14))))
                                                                                                                                    case "a":
                                                                                                                                        if len(v2) > 0 {
                                                                                                                                            v2[0] = fi13
                                                                                                                                        }
                                                                                                                                        v6 = v6
                                                                                                                                        var v131 int = (fi13 ^ (^(fi13)))
                                                                                                                                        _ = v131
                                                                                                                                    default:
                                                                                                                                        var v132 float32 = ((0.1 - (1e10 + 1e10)) * 0.1)
                                                                                                                                        _ = v132
                                                                                                                                    }
                                                                                                                                    v2 = append(v2, (fi13 + (-(fi14))))
                                                                                                                                    if v5 {
                                                                                                                                        if true {
                                                                                                                                            switch (fi13 - (fi14 & 1)) {
                                                                                                                                            case -9223372036854775808:
                                                                                                                                                v6 = v6
                                                                                                                                            case 175:
                                                                                                                                                v1 += (v1 ^ (v1 << ((1) & 63)))
                                                                                                                                            case 42:
                                                                                                                                                var v133 uint64 = 42
                                                                                                                                                _ = v133
                                                                                                                                            default:
                                                                                                                                                if v5 {
                                                                                                                                                    v134 := GenS4{F0: 432, F1: fi14, F2: 42}
                                                                                                                                                    v135 := GenS5{F0: -729, F1: 0}
                                                                                                                                                    _ = v134
                                                                                                                                                    _ = v135
                                                                                                                                                } else {
                                                                                                                                                    switch "foo bar" {
                                                                                                                                                    case "foo bar":
                                                                                                                                                        fmt.Println(v116, v1, v5)
                                                                                                                                                    }
                                                                                                                                                    v136 := GenS13{F0: v6}
                                                                                                                                                    _ = v136
                                                                                                                                                }
                                                                                                                                                v3[(fi14 | fi14)] = "a"
                                                                                                                                                if len(v2) > 0 {
                                                                                                                                                    v2[0] = (fi13 % ((fi13) | 1))
                                                                                                                                                }
                                                                                                                                                var v137 float64 = (1e10 * 3.14159)
                                                                                                                                                _ = v137
                                                                                                                                            }
                                                                                                                                            v6 = (v6 & v6)
                                                                                                                                            v6 += v6
                                                                                                                                        }
                                                                                                                                        v1 -= (v1 + 0)
                                                                                                                                    }
                                                                                                                                    fmt.Println(v5)
                                                                                                                                    _ = v116
                                                                                                                                }
                                                                                                                            }
                                                                                                                            if true {
                                                                                                                                v138 := GenS9{F0: (fi14 + (fi14 - fi13))}
                                                                                                                                var v139 int16 = 0
                                                                                                                                _ = v138
                                                                                                                                _ = v139
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            fmt.Println(v1, v1)
                                                                                                                            if ((0.1 * 3.14159) > (1.5 + -2.25)) {
                                                                                                                                v2 = append(v2, 9223372036854775807)
                                                                                                                                v3[fi14] = ""
                                                                                                                                switch fi13 {
                                                                                                                                case 32768:
                                                                                                                                    v140 := GenS14{F0: uint32(v1), F1: -914}
                                                                                                                                    _ = v140
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                    fmt.Println(fi14, v6, fi13)
                                                                                                                }
                                                                                                                var v141 uint64 = 779
                                                                                                                v4 -= v4
                                                                                                                _ = v141
                                                                                                            } else {
                                                                                                                var v142 uint32 = 2147483648
                                                                                                                if v5 {
                                                                                                                    fmt.Println(fi13, v5)
                                                                                                                    v143 := GenS5{F0: 2147483647, F1: 625}
                                                                                                                    switch fi13 {
                                                                                                                    case 0:
                                                                                                                        fmt.Println(v4)
                                                                                                                    case 802:
                                                                                                                        switch "a" {
                                                                                                                        case "foo bar":
                                                                                                                            v144 := GenS4{F0: 1, F1: fi13, F2: v142}
                                                                                                                            _ = v144
                                                                                                                        case "":
                                                                                                                            var v145 string = "a"
                                                                                                                            _ = v145
                                                                                                                        default:
                                                                                                                            if len(v2) >= 2 {
                                                                                                                                v2 = v2[1:len(v2)]
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v142 = v142
                                                                                                                        if false {
                                                                                                                            var v146 bool = (!(!true))
                                                                                                                            _ = v146
                                                                                                                        } else {
                                                                                                                            v6 += (v6 / ((-1) | 1))
                                                                                                                        }
                                                                                                                        if len(v2) > 0 {
                                                                                                                            v2[0] = 32768
                                                                                                                        }
                                                                                                                    case -9223372036854775808:
                                                                                                                        if (false || (65535 <= 0)) {
                                                                                                                            switch "\n\t" {
                                                                                                                            case "\n\t":
                                                                                                                                var v147 uint = 832
                                                                                                                                _ = v147
                                                                                                                            case "foo bar":
                                                                                                                                v148 := GenS15{F0: -1}
                                                                                                                                _ = v148
                                                                                                                            case "":
                                                                                                                                delete(v3, (fi14 % ((fi13) | 1)))
                                                                                                                                if (3.14159 == ((0.1 + (0.1 - 3.14159)) + 0.1)) {
                                                                                                                                    switch "hello" {
                                                                                                                                    case "hello":
                                                                                                                                        v1 += (^(v1))
                                                                                                                                    case "foo bar":
                                                                                                                                        v2 = append(v2, -1)
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    var v149 uint16 = 32767
                                                                                                                                    if v5 {
                                                                                                                                        switch (fi14 / (((fi14 | (fi13 << ((4294967295) & 63)))) | 1)) {
                                                                                                                                        case -409:
                                                                                                                                            v150 := GenS8{F0: 413}
                                                                                                                                            _ = v150
                                                                                                                                        default:
                                                                                                                                            v5 = ((!(false || (1 >= 42))) && true)
                                                                                                                                            v151 := GenS16{F0: 42, F1: (v142 - v142), F2: ""}
                                                                                                                                            _ = v151
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    _ = v149
                                                                                                                                }
                                                                                                                                v2 = append(v2, 4294967295)
                                                                                                                            default:
                                                                                                                                v6 += v6
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            var v152 bool = (!v5)
                                                                                                                            if (((v6 >> ((540) & 63)) <= v6) && v5) {
                                                                                                                                var v153 bool = ((-2147483648 >= 42) || v152)
                                                                                                                                if (0.0 != 0.0) {
                                                                                                                                    if (909 > 0) {
                                                                                                                                        if (v153 || v152) {
                                                                                                                                            if v153 {
                                                                                                                                                if len(v2) > 0 {
                                                                                                                                                    v2[0] = (fi14 >> ((0) & 63))
                                                                                                                                                }
                                                                                                                                                switch (fi13 | 269) {
                                                                                                                                                case 42:
                                                                                                                                                    if (42 != 1) {
                                                                                                                                                        v154 := GenS5{F0: -711, F1: -978}
                                                                                                                                                        _ = v154
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                                v152 = v5
                                                                                                                                            }
                                                                                                                                        } else {
                                                                                                                                            v4 = 123
                                                                                                                                            fmt.Println(v5)
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        switch "\n\t" {
                                                                                                                                        case "foo bar":
                                                                                                                                            v155 := GenS9{F0: (fi14 % ((fi13) | 1))}
                                                                                                                                            _ = v155
                                                                                                                                        }
                                                                                                                                        switch (fi13 * (fi13 * 32767)) {
                                                                                                                                        case 128:
                                                                                                                                            delete(v3, (fi13 >> ((255) & 63)))
                                                                                                                                        case 127:
                                                                                                                                            var v156 uint = 787
                                                                                                                                            _ = v156
                                                                                                                                        default:
                                                                                                                                            var v157 uint64 = 556
                                                                                                                                            _ = v157
                                                                                                                                        }
                                                                                                                                        v5 = (v5 && v152)
                                                                                                                                    }
                                                                                                                                    fmt.Println(v153, v152)
                                                                                                                                    var v158 uint64 = 830
                                                                                                                                    _ = v158
                                                                                                                                }
                                                                                                                                if (v153 || false) {
                                                                                                                                    v1 = v1
                                                                                                                                    var v159 int = (fi13 >> ((2) & 63))
                                                                                                                                    _ = len(v3)
                                                                                                                                    _ = v159
                                                                                                                                } else {
                                                                                                                                    v160 := GenS17{F0: (v142 - uint32(v1))}
                                                                                                                                    _ = v160
                                                                                                                                }
                                                                                                                                _ = v153
                                                                                                                            } else {
                                                                                                                                var v161 string = "hello"
                                                                                                                                if v152 {
                                                                                                                                    fmt.Println(fi13)
                                                                                                                                }
                                                                                                                                var v162 string = v161
                                                                                                                                _ = v161
                                                                                                                                _ = v162
                                                                                                                            }
                                                                                                                            _ = v152
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v143
                                                                                                                } else {
                                                                                                                    v4 = v4
                                                                                                                    if (((v5 || v5) || ((v6 + (v6 % ((v6) | 1))) != int64(v4))) && ((!true) && v5)) {
                                                                                                                        var v163 int = fi13
                                                                                                                        var v164 int8 = (v4 / ((v4) | 1))
                                                                                                                        _ = len(v3)
                                                                                                                        _ = v163
                                                                                                                        _ = v164
                                                                                                                    } else {
                                                                                                                        var v165 uint32 = v142
                                                                                                                        var v166 uint32 = 34
                                                                                                                        _ = v165
                                                                                                                        _ = v166
                                                                                                                    }
                                                                                                                    var v167 uint64 = 0
                                                                                                                    _ = v167
                                                                                                                }
                                                                                                                v168 := GenS3{F0: (float64(-2.25) != 3.14159), F1: "a", F2: 32767}
                                                                                                                _ = v142
                                                                                                                _ = v168
                                                                                                            }
                                                                                                            v169 := GenS15{F0: 1}
                                                                                                            _ = v169
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    var v170 int8 = v4
                                                                                                    _ = v170
                                                                                                }
                                                                                                v5 = (!(v5 && false))
                                                                                                if v5 {
                                                                                                    v2 = append(v2, -684)
                                                                                                    v3[(fi13 / ((fi14) | 1))] = ("\n\t" + "a")
                                                                                                    if v5 {
                                                                                                        var v171 uint64 = 852
                                                                                                        var v172 int8 = -1
                                                                                                        switch (fi13 + (fi13 / (((fi13 & fi13)) | 1))) {
                                                                                                        case 32768:
                                                                                                            var v173 bool = v5
                                                                                                            _ = v173
                                                                                                        case -1:
                                                                                                            _ = len(v3)
                                                                                                        }
                                                                                                        _ = v171
                                                                                                        _ = v172
                                                                                                    } else {
                                                                                                        var v174 bool = (!(1 == 798))
                                                                                                        _ = v174
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                if true {
                                                                                                    var v175 bool = v5
                                                                                                    var v176 bool = true
                                                                                                    if (!((fi14 << ((uint(v1)) & 63)) == fi13)) {
                                                                                                        v5 = v175
                                                                                                        v175 = false
                                                                                                        v177 := GenS18{F0: 346, F1: -9223372036854775808, F2: 2147483647}
                                                                                                        _ = v177
                                                                                                    } else {
                                                                                                        var v178 int64 = 940
                                                                                                        fmt.Println(v1)
                                                                                                        switch (("" + "foo bar") + "foo bar") {
                                                                                                        case "":
                                                                                                            v1 = v1
                                                                                                        case "\n\t":
                                                                                                            v178 = (v178 - -844)
                                                                                                        }
                                                                                                        _ = v178
                                                                                                    }
                                                                                                    _ = v175
                                                                                                    _ = v176
                                                                                                } else {
                                                                                                    fmt.Println(fi13, v5, fi13)
                                                                                                    var v179 uint64 = 42
                                                                                                    v180 := GenS2{F0: 789, F1: v6}
                                                                                                    _ = v179
                                                                                                    _ = v180
                                                                                                }
                                                                                            }
                                                                                            var v181 int = (fi13 | (fi14 << ((672) & 63)))
                                                                                            _ = v181
                                                                                        }
                                                                                    }
                                                                                    switch 32768 {
                                                                                    case -1:
                                                                                        if false {
                                                                                            v182 := GenS15{F0: 1}
                                                                                            if v5 {
                                                                                                switch "a" {
                                                                                                case "foo bar":
                                                                                                    if true {
                                                                                                        v6 += (-(v6))
                                                                                                        var v183 int64 = (^(v6))
                                                                                                        if (1 == 4294967295) {
                                                                                                            v1 = 2
                                                                                                            if (0 == -32768) {
                                                                                                                switch (("\n\t" + "") + ("\n\t" + ("hello" + (("foo bar" + "hello") + ("" + "foo bar"))))) {
                                                                                                                case "\n\t":
                                                                                                                    var v184 uint16 = 0
                                                                                                                    _ = v184
                                                                                                                }
                                                                                                                v2 = append(v2, -1)
                                                                                                                var v185 int16 = -32768
                                                                                                                _ = v185
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v183
                                                                                                    } else {
                                                                                                        v5 = ((-99 <= -1) || (float32(((3.14159 + ((float64(((1e10 * ((1e10 / 1e10) * 0.1)) - (3.14159 / float32(0.1)))) + 0.1) / -2.25)) + 0.0)) >= float32(-2.25)))
                                                                                                        var v186 int32 = 0
                                                                                                        _ = v186
                                                                                                    }
                                                                                                case "a":
                                                                                                    fmt.Println(v1)
                                                                                                default:
                                                                                                    v6 -= v6
                                                                                                }
                                                                                            }
                                                                                            _ = v182
                                                                                        } else {
                                                                                            var v187 bool = false
                                                                                            if v187 {
                                                                                                var v188 string = ("hello" + "")
                                                                                                _ = v188
                                                                                            }
                                                                                            var v189 uint = 914
                                                                                            if true {
                                                                                                var v190 float32 = 3.14159
                                                                                                _ = v190
                                                                                            }
                                                                                            switch fi13 {
                                                                                            case 127:
                                                                                                v5 = (!((v4 - v4) < (v4 ^ 42)))
                                                                                            case 32768:
                                                                                                if ((388 > 65535) || (!v5)) {
                                                                                                    fmt.Println(v5, fi13, v189)
                                                                                                } else {
                                                                                                    if false {
                                                                                                        var v191 float32 = -2.25
                                                                                                        v1 = v1
                                                                                                        _ = v191
                                                                                                    } else {
                                                                                                        switch "a" {
                                                                                                        case "":
                                                                                                            var v192 string = ("foo bar" + "\n\t")
                                                                                                            _ = v192
                                                                                                        default:
                                                                                                            v193 := GenS15{F0: 42}
                                                                                                            _ = v193
                                                                                                        }
                                                                                                        if (false && true) {
                                                                                                            fmt.Println(v187, v1, v1)
                                                                                                        } else {
                                                                                                            v1 = 12
                                                                                                        }
                                                                                                        v3[-831] = (("foo bar" + ((("a" + "foo bar") + "foo bar") + ("\n\t" + ("foo bar" + "a")))) + "hello")
                                                                                                    }
                                                                                                    v194 := GenS19{F0: (^(v6)), F1: -32768}
                                                                                                    _ = v194
                                                                                                }
                                                                                            default:
                                                                                                v195 := GenS20{F0: (true && (1 > 2))}
                                                                                                _ = v195
                                                                                            }
                                                                                            _ = v187
                                                                                            _ = v189
                                                                                        }
                                                                                    case 391:
                                                                                        var v196 uint = 4294967295
                                                                                        _ = v196
                                                                                    case 498:
                                                                                        switch (fi13 >> ((360) & 63)) {
                                                                                        case -880:
                                                                                            v3[fi14] = ("a" + "foo bar")
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        case "foo bar":
                                                                            if v5 {
                                                                                var v197 uint8 = (v1 >> ((463) & 63))
                                                                                var v198 uint8 = (v1 - 128)
                                                                                if v5 {
                                                                                    if true {
                                                                                        fmt.Println(v5, v1)
                                                                                    }
                                                                                    if v5 {
                                                                                        v199 := GenS21{F0: fi13, F1: 2147483647}
                                                                                        if (42 > 1) {
                                                                                            v2 = append(v2, fi13)
                                                                                        }
                                                                                        switch ("hello" + "\n\t") {
                                                                                        case "\n\t":
                                                                                            v6 -= (-(v6))
                                                                                        case "a":
                                                                                            var v200 int8 = (v4 + v4)
                                                                                            _ = v200
                                                                                        }
                                                                                        fmt.Println(v6, v198)
                                                                                        v201 := GenS22{F0: v6, F1: 334, F2: -496}
                                                                                        _ = v199
                                                                                        _ = v201
                                                                                    }
                                                                                }
                                                                                _ = v197
                                                                                _ = v198
                                                                            } else {
                                                                                if (1 > 2147483648) {
                                                                                    var v202 bool = (!v5)
                                                                                    v202 = (v5 || ((-2.25 / ((-2.25 + 1e10) * -2.25)) != -2.25))
                                                                                    _ = v202
                                                                                } else {
                                                                                    var v203 int16 = 0
                                                                                    if (448 == 449) {
                                                                                        v203 = -507
                                                                                        v4 -= (v4 ^ v4)
                                                                                        fmt.Println(v5)
                                                                                    } else {
                                                                                        v204 := GenS13{F0: -1}
                                                                                        _ = v204
                                                                                    }
                                                                                    _ = v203
                                                                                }
                                                                                if (!true) {
                                                                                    v205 := GenS23{F0: ("" + "\n\t"), F1: true}
                                                                                    v6 = (v6 - v6)
                                                                                    _ = v205
                                                                                } else {
                                                                                    if true {
                                                                                        var v206 string = ((("\n\t" + ("" + "\n\t")) + "\n\t") + "foo bar")
                                                                                        _ = v206
                                                                                    }
                                                                                    if v5 {
                                                                                        if ((fi14 - (fi14 ^ (^(fi13)))) != fi14) {
                                                                                            fmt.Println(fi13, fi13, fi14)
                                                                                        }
                                                                                        var v207 uint8 = 1
                                                                                        _ = v207
                                                                                    }
                                                                                    v208 := GenS24{F0: (v6 * (^(v6)))}
                                                                                    _ = v208
                                                                                }
                                                                                v1 = 255
                                                                                if (fi14 >= fi14) {
                                                                                    if ((^(v4)) != (v4 ^ (v4 - genFn1("hello", 2147483647, "")))) {
                                                                                        var v209 string = "foo bar"
                                                                                        _ = v209
                                                                                    }
                                                                                    _ = len(v2)
                                                                                    v210 := GenS9{F0: (fi14 >> ((42) & 63))}
                                                                                    if (-1 > fi13) {
                                                                                        v211 := v3[fi13]
                                                                                        if ((1.5 + float64(-2.25)) > 1.5) {
                                                                                            v6 -= 1
                                                                                            v211 = ""
                                                                                            v4 -= (v4 ^ (v4 | 1))
                                                                                            v4 -= -1
                                                                                        }
                                                                                        var v212 uint8 = 117
                                                                                        _ = v211
                                                                                        _ = v212
                                                                                    } else {
                                                                                        switch "\n\t" {
                                                                                        case "":
                                                                                            v3[int(v4)] = "\n\t"
                                                                                        case "foo bar":
                                                                                            var v213 int8 = v4
                                                                                            _ = v213
                                                                                        case "a":
                                                                                            var v214 uint32 = 1
                                                                                            _ = v214
                                                                                        }
                                                                                        var v215 uint8 = (v1 % ((genFn0("foo bar", 255)) | 1))
                                                                                        var v216 int16 = -1
                                                                                        _ = v215
                                                                                        _ = v216
                                                                                    }
                                                                                    _ = v210
                                                                                }
                                                                                v6 = (-(v6))
                                                                                if (!((v5 && v5) || v5)) {
                                                                                    v6 = v6
                                                                                    if ((!(float32((1.5 + float64((float32(((float64(3.14159) * (((float64(3.14159) + 1.5) * float64(0.1)) + 1.5)) + (1e10 / (3.14159 * 3.14159)))) / (0.0 + ((float32((-2.25 - -2.25)) - (float32(-2.25) - 0.0)) / (1e10 * 1e10))))))) <= 0.0)) && ((0.1 / 0.1) <= 3.14159)) {
                                                                                        var v217 string = "a"
                                                                                        var v218 float64 = 0.1
                                                                                        v6 = (v6 / ((v6) | 1))
                                                                                        _ = v217
                                                                                        _ = v218
                                                                                    } else {
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                        fmt.Println(v4)
                                                                                        if true {
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                            if (!(true && true)) {
                                                                                                v5 = false
                                                                                                v5 = (!v5)
                                                                                            } else {
                                                                                                var v219 uint8 = (v1 * (v1 & genFn0(("foo bar" + "hello"), v1)))
                                                                                                _ = v219
                                                                                            }
                                                                                            if v5 {
                                                                                                if v5 {
                                                                                                    var v220 float64 = (((((1.5 + 0.0) / 3.14159) + 0.1) * (1e10 - -2.25)) * 1.5)
                                                                                                    fmt.Println(v1, v4, fi14)
                                                                                                    v221 := GenS22{F0: -9223372036854775808, F1: 210, F2: v6}
                                                                                                    _ = v220
                                                                                                    _ = v221
                                                                                                } else {
                                                                                                    fmt.Println(fi14, v5)
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        if true {
                                                                                            v1 -= (^(v1))
                                                                                            if v5 {
                                                                                                v222 := GenS9{F0: (fi14 | fi14)}
                                                                                                var v223 int64 = v6
                                                                                                v224 := GenS25{F0: (fi14 + fi14), F1: v223, F2: 9}
                                                                                                _ = v222
                                                                                                _ = v223
                                                                                                _ = v224
                                                                                            } else {
                                                                                                var v225 uint8 = v1
                                                                                                var v226 int64 = v6
                                                                                                if (1e10 == float32(0.1)) {
                                                                                                    if true {
                                                                                                        var v227 uint8 = v225
                                                                                                        var v228 string = "a"
                                                                                                        _ = v227
                                                                                                        _ = v228
                                                                                                    } else {
                                                                                                        v229 := GenS7{F0: 1, F1: 1, F2: 782}
                                                                                                        v226 = (^(v226))
                                                                                                        v230 := GenS13{F0: v6}
                                                                                                        _ = v229
                                                                                                        _ = v230
                                                                                                    }
                                                                                                    v4 = v4
                                                                                                    v231 := GenS18{F0: (fi13 | (fi14 - 189)), F1: 42, F2: 65535}
                                                                                                    _ = v231
                                                                                                } else {
                                                                                                    if true {
                                                                                                        if v5 {
                                                                                                            v4 -= (v4 % ((v4) | 1))
                                                                                                            v232 := GenS26{F0: 639, F1: 18446744073709551615, F2: (^(v1))}
                                                                                                            _ = v232
                                                                                                        }
                                                                                                        v226 = v226
                                                                                                        fmt.Println(fi14, v225)
                                                                                                    } else {
                                                                                                        v233 := GenS27{F0: 343}
                                                                                                        v234 := GenS19{F0: (v6 / ((v6) | 1)), F1: 0}
                                                                                                        _ = v233
                                                                                                        _ = v234
                                                                                                    }
                                                                                                    fmt.Println(v225)
                                                                                                    if len(v2) >= 2 {
                                                                                                        v2 = v2[1:len(v2)]
                                                                                                    }
                                                                                                }
                                                                                                v4 -= v4
                                                                                                _ = len(v3)
                                                                                                v225 -= 128
                                                                                                _ = v225
                                                                                                _ = v226
                                                                                            }
                                                                                            fmt.Println(fi14)
                                                                                        } else {
                                                                                            switch ("hello" + "\n\t") {
                                                                                            case "\n\t":
                                                                                                if v5 {
                                                                                                    var v235 uint8 = (v1 ^ (v1 + v1))
                                                                                                    var v236 float64 = 0.0
                                                                                                    v4 = v4
                                                                                                    _ = v235
                                                                                                    _ = v236
                                                                                                }
                                                                                            case "hello":
                                                                                                if (840 <= 255) {
                                                                                                    v4 -= (^(v4))
                                                                                                    v237 := v3[(fi13 + int(v6))]
                                                                                                    _ = len(v3)
                                                                                                    _ = v237
                                                                                                }
                                                                                            default:
                                                                                                switch (fi14 - (fi14 - 1)) {
                                                                                                case -1:
                                                                                                    v6 -= (v6 % (((v6 - v6)) | 1))
                                                                                                }
                                                                                                if (false && ((v1 / (((v1 - 2)) | 1)) < 127)) {
                                                                                                    v238 := GenS28{F0: -525, F1: ("\n\t" + ("a" + "a"))}
                                                                                                    _ = v238
                                                                                                }
                                                                                                var v239 int32 = 0
                                                                                                _ = v239
                                                                                            }
                                                                                        }
                                                                                        v240 := GenS8{F0: 943}
                                                                                        _ = v240
                                                                                    }
                                                                                    v241 := GenS12{F0: -1, F1: 2}
                                                                                    var v242 uint64 = 0
                                                                                    v242 -= (v242 | 4294967295)
                                                                                    _ = v241
                                                                                    _ = v242
                                                                                }
                                                                                if true {
                                                                                    v5 = v5
                                                                                    v5 = (65535 <= 390)
                                                                                }
                                                                                v243 := GenS29{F0: 65535}
                                                                                if v5 {
                                                                                    v244 := GenS30{F0: 823, F1: (v1 << ((0) & 63))}
                                                                                    if (fi14 <= fi13) {
                                                                                        switch 4294967295 {
                                                                                        case 32768:
                                                                                            switch "foo bar" {
                                                                                            case "a":
                                                                                                fmt.Println(fi14, v5)
                                                                                            case "":
                                                                                                var v245 uint8 = 255
                                                                                                _ = v245
                                                                                            case "foo bar":
                                                                                                var v246 int = fi13
                                                                                                v6 -= v6
                                                                                                if false {
                                                                                                    fmt.Println(v4)
                                                                                                } else {
                                                                                                    var v247 uint16 = 2
                                                                                                    _ = v247
                                                                                                }
                                                                                                _ = v246
                                                                                            default:
                                                                                                if ((v6 | v6) == (v6 << ((145) & 63))) {
                                                                                                    var v248 bool = (32767 >= 65535)
                                                                                                    _ = v248
                                                                                                } else {
                                                                                                    if v5 {
                                                                                                        v6 -= v6
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        case -9223372036854775808:
                                                                                            fmt.Println(v1)
                                                                                        case 42:
                                                                                            v6 -= 42
                                                                                        }
                                                                                        if (int8(fi14) >= genFn1("\n\t", 0, "")) {
                                                                                            fmt.Println(v6, fi13, v5)
                                                                                            var v249 int8 = v4
                                                                                            _ = v249
                                                                                        }
                                                                                        var v250 float32 = 0.0
                                                                                        _ = v250
                                                                                    } else {
                                                                                        fmt.Println(v5, v5, fi14)
                                                                                    }
                                                                                    if (1 == 2) {
                                                                                        var v251 int32 = -2147483648
                                                                                        v252 := GenS15{F0: v251}
                                                                                        _ = v251
                                                                                        _ = v252
                                                                                    } else {
                                                                                        v5 = (((!v5) && v5) && false)
                                                                                        if (v4 <= (v4 << ((619) & 63))) {
                                                                                            v3[4294967295] = ("a" + ("" + ("hello" + "a")))
                                                                                            v253 := v3[(^(fi13))]
                                                                                            _ = len(v3)
                                                                                            _ = v253
                                                                                        } else {
                                                                                            var v254 float64 = 0.0
                                                                                            if v5 {
                                                                                                _ = len(v3)
                                                                                                switch ((("" + "") + ("hello" + ("" + (("foo bar" + "hello") + ("foo bar" + "hello"))))) + "foo bar") {
                                                                                                case "\n\t":
                                                                                                    var v255 bool = false
                                                                                                    var v256 int16 = 32767
                                                                                                    _ = v255
                                                                                                    _ = v256
                                                                                                case "a":
                                                                                                    v257 := GenS31{F0: 100}
                                                                                                    _ = v257
                                                                                                default:
                                                                                                    fmt.Println(v1, v254)
                                                                                                }
                                                                                            } else {
                                                                                                if (false && (4294967295 > 2147483648)) {
                                                                                                    fmt.Println(fi13, v254, v6)
                                                                                                } else {
                                                                                                    if len(v2) > 0 {
                                                                                                        v2[0] = (-(fi13))
                                                                                                    }
                                                                                                    if true {
                                                                                                        v1 = v1
                                                                                                    } else {
                                                                                                        v258 := GenS19{F0: v6, F1: 1000}
                                                                                                        v259 := GenS6{F0: fi13, F1: fi13}
                                                                                                        v5 = false
                                                                                                        _ = v258
                                                                                                        _ = v259
                                                                                                    }
                                                                                                    var v260 string = "a"
                                                                                                    _ = v260
                                                                                                }
                                                                                            }
                                                                                            if (fi13 >= (fi14 | fi13)) {
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = fi14
                                                                                                }
                                                                                            }
                                                                                            var v261 int = (-(fi13))
                                                                                            _ = v254
                                                                                            _ = v261
                                                                                        }
                                                                                        var v262 bool = (!(42 >= -32768))
                                                                                        var v263 int = (fi14 % ((fi13) | 1))
                                                                                        var v264 uint16 = 0
                                                                                        var v265 float32 = (0.1 * 0.0)
                                                                                        var v266 int64 = (-(v6))
                                                                                        _ = v262
                                                                                        _ = v263
                                                                                        _ = v264
                                                                                        _ = v265
                                                                                        _ = v266
                                                                                    }
                                                                                    _ = v244
                                                                                } else {
                                                                                    delete(v3, (fi14 >> ((65535) & 63)))
                                                                                    v267 := GenS32{F0: (!((v5 || false) || (v5 && true))), F1: "hello"}
                                                                                    var v268 string = "a"
                                                                                    _ = v267
                                                                                    _ = v268
                                                                                }
                                                                                _ = v243
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            default:
                                                                v269 := GenS13{F0: (^(v6))}
                                                                _ = v269
                                                            }
                                                            v1 -= genFn0("", v1)
                                                        }
                                                    } else {
                                                        var v270 uint = 137
                                                        _ = v270
                                                    }
                                                case -910:
                                                    v6 -= -1
                                                }
                                            }
                                            _ = v47
                                        } else {
                                            fmt.Println(fi14, v1, fi14)
                                        }
                                        fmt.Println(fi14, v6, v4)
                                        v5 = ((-1 > 0) || (v5 || false))
                                        if len(v2) > 0 {
                                            v2[0] = (fi14 | (fi14 ^ (fi14 | (fi14 + fi14))))
                                        }
                                        var v271 string = ""
                                        _ = v271
                                    }
                                    fmt.Println(v6, fi14)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v1 = v1
                                }
                            case 0:
                                v4 -= (v4 | (v4 - 42))
                            default:
                                if len(v2) > 0 {
                                    v2[0] = (-(fi14))
                                }
                            }
                        }
                    case 218:
                        v272 := GenS33{F0: 563}
                        _ = v272
                    case 0:
                        var v273 uint8 = v1
                        _ = v273
                    default:
                        var v274 float32 = 1.5
                        _ = v274
                    }
                }
                _ = fi14
            }
        }
        for fi275 := 0; fi275 < 2; fi275++ {
            var v276 uint32 = 1
            v5 = v5
            v6 = 875
            _ = fi275
            _ = v276
        }
        fmt.Println(v4, fi13, v1)
        _ = fi13
    }
    var v277 float64 = 1.5
    v278 := GenS34{F0: 730, F1: 127}
    v277 += (float64((1.5 - -2.25)) - v277)
    v279 := GenS35{F0: 17, F1: 32767}
    switch "a" {
    case "\n\t":
        if false {
            for fi280 := 0; fi280 < 4; fi280++ {
                v277 -= v277
                v277 = ((1.5 * 0.0) - ((((0.0 + v277) - (1.5 + 1.5)) / 0.1) / v277))
                _ = fi280
            }
            var v281 uint8 = (^(v1))
            fmt.Println(v277)
            _ = v281
        }
    case "foo bar":
        v282 := GenS36{F0: 948}
        _ = v282
    }
    delete(v3, 42)
    for fi283 := 0; fi283 < 19; fi283++ {
        v4 = (^(v4))
        for fi284 := 0; fi284 < 6; fi284++ {
            _ = len(v3)
            v1 += 0
            _ = fi284
        }
        var v285 string = ""
        _ = fi283
        _ = v285
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch -16 {
    case 0:
        for fi286 := 0; fi286 < 11; fi286++ {
            v1 += v1
            v6 = -1
            var v287 float64 = v277
            _ = fi286
            _ = v287
        }
    case 4294967295:
        v288 := v3[4294967295]
        _ = v288
    case -9223372036854775808:
        for fi289 := 0; fi289 < 14; fi289++ {
            switch (-(fi289)) {
            case -2147483648:
                var v290 float64 = v277
                _ = v290
            }
            v291 := v3[(-(fi289))]
            v292 := GenS37{F0: (^(fi289))}
            _ = fi289
            _ = v291
            _ = v292
        }
    }
    v6 = (v6 * v6)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v277:", v277)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v277
    _ = v278
    _ = v279
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
