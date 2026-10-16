package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int16; F1 int }
type GenS7 struct { F0 string; F1 uint; F2 int32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int8; F1 int64 }
type GenS11 struct { F0 int8; F1 int8 }
type GenS12 struct { F0 int; F1 uint64; F2 int8 }
func genFn0() int8 {
    v4 := GenS1{F0: 4294967295, F1: 1, F2: (false || false)}
    if ((((true || (false && (true || (((0 == 65535) || false) && false)))) && false) && true) || true) {
        var v5 uint64 = 2
        _ = v5
    }
    switch 127 {
    case 2147483647:
        var v6 uint16 = 32767
        _ = v6
    }
    switch "hello" {
    case "\n\t":
        v7 := []int{127, -486}
        _ = v7
    default:
        if (false || (!true)) {
            var v8 string = (((("a" + "foo bar") + ("a" + "\n\t")) + "hello") + "\n\t")
            switch 0 {
            case 9223372036854775807:
                v9 := []int{127}
                _ = v9
            case -668:
                v8 = ("hello" + v8)
            case -9223372036854775808:
                v8 = v8
            default:
                v10 := []int{-256, -101, 128}
                _ = v10
            }
            v8 = ("\n\t" + v8)
            _ = v8
        }
    }
    _ = v4
    return -106
}
func genFn1() uint16 {
    v11 := []int{}
    var v12 int64 = -635
    if (-2.25 > float64((-2.25 / -2.25))) {
        v13 := GenS1{F0: 42, F1: v12, F2: (-2.25 == 1.5)}
        _ = len(v11)
        for fi14 := 0; fi14 < 8; fi14++ {
            if true {
                switch (fi14 % (((fi14 | fi14)) | 1)) {
                case 128:
                    fmt.Println(v12, fi14)
                default:
                    v15 := GenS1{F0: 1, F1: (v12 - -849), F2: false}
                    _ = v15
                }
                v12 += (v12 % ((v12) | 1))
            } else {
                switch "a" {
                case "a":
                    for fi16 := 0; fi16 < 17; fi16++ {
                        if false {
                            var v17 uint64 = 0
                            var v18 string = ("foo bar" + ("a" + "foo bar"))
                            _ = v17
                            _ = v18
                        }
                        fmt.Println(fi14)
                        _ = len(v11)
                        _ = fi16
                    }
                case "foo bar":
                    switch ("hello" + "\n\t") {
                    case "":
                        v11 = append(v11, fi14)
                    }
                }
                fmt.Println(v12, v12)
            }
            _ = fi14
        }
        _ = v13
    } else {
        if (!true) {
            for fi19 := 0; fi19 < 20; fi19++ {
                for fi20 := 0; fi20 < 11; fi20++ {
                    v12 -= (v12 % ((0) | 1))
                    if true {
                        v12 = -1
                        var v21 int = fi20
                        _ = v21
                    }
                    _ = fi20
                }
                _ = fi19
            }
        }
        switch 9223372036854775807 {
        case 127:
            var v22 int = 128
            _ = v22
        case 32767:
            v12 = (v12 << ((255) & 63))
        }
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        fmt.Println(fi23, v12)
        v12 -= v12
        _ = fi23
    }
    _ = v11
    _ = v12
    return 42
}

func main() {
    v1 := []int{-637, 127}
    var v2 int32 = 859
    var v3 int16 = -32768
    v2 += v2
    if false {
        var v24 uint16 = 6
        if (!true) {
            v25 := GenS1{F0: 2, F1: 1, F2: true}
            _ = v25
        } else {
            if true {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi26 := 0; fi26 < 16; fi26++ {
                    if false {
                        if false {
                            fmt.Println(fi26, v3, fi26)
                            switch fi26 {
                            case -1:
                                v2 -= 42
                            case -2147483648:
                                var v27 uint32 = 301
                                _ = v27
                            default:
                                switch "" {
                                case "":
                                    fmt.Println(fi26, v2, v2)
                                case "a":
                                    v28 := GenS2{F0: 292}
                                    _ = v28
                                default:
                                    fmt.Println(v3, v3, fi26)
                                }
                            }
                            if ((("" + "foo bar") + "a") <= "") {
                                for fi29 := 0; fi29 < 16; fi29++ {
                                    v2 = (v2 % (((v2 / (((-(v2))) | 1))) | 1))
                                    var v30 uint8 = 127
                                    fmt.Println(fi29)
                                    _ = fi29
                                    _ = v30
                                }
                            }
                        }
                    } else {
                        var v31 float32 = float32(((0.0 - -2.25) * 1.5))
                        switch 9223372036854775807 {
                        case 32768:
                            v31 = v31
                        case -9223372036854775808:
                            v32 := GenS1{F0: uint64(v24), F1: 9223372036854775807, F2: (-128 <= -122)}
                            _ = v32
                        case 1:
                            if true {
                                v1 = append(v1, (fi26 % ((fi26) | 1)))
                                if (0.0 == -2.25) {
                                    if true {
                                        switch (fi26 / (((fi26 / ((-1) | 1))) | 1)) {
                                        case -9223372036854775808:
                                            v2 = (v2 ^ int32(v24))
                                        default:
                                            if true {
                                                switch (fi26 & fi26) {
                                                case -1:
                                                    var v33 int32 = v2
                                                    _ = v33
                                                default:
                                                    var v34 int8 = genFn0()
                                                    _ = v34
                                                }
                                            } else {
                                                v35 := GenS2{F0: 65535}
                                                fmt.Println(fi26, v2)
                                                _ = v35
                                            }
                                        }
                                        v31 += (v31 + ((((v31 * v31) - 1e10) - v31) + 0.0))
                                        var v36 uint32 = 2147483648
                                        _ = v36
                                    } else {
                                        for fi37 := 0; fi37 < 5; fi37++ {
                                            _ = len(v1)
                                            _ = fi37
                                        }
                                        switch "a" {
                                        case "\n\t":
                                            switch ("" + ("" + (("foo bar" + "") + (("hello" + ((("" + ("a" + ("hello" + "foo bar"))) + "\n\t") + ((("\n\t" + "foo bar") + "a") + "\n\t"))) + "a")))) {
                                            case "foo bar":
                                                v2 -= (-(v2))
                                            case "hello":
                                                v31 -= (0.1 + ((v31 / (1e10 * ((float32(0.0) / v31) + v31))) / v31))
                                            case "\n\t":
                                                if false {
                                                    if (4294967295 < 295) {
                                                        v2 += (v2 - v2)
                                                    } else {
                                                        if false {
                                                            v38 := GenS1{F0: 18446744073709551615, F1: 42, F2: (2 >= 127)}
                                                            v39 := GenS3{F0: -1, F1: (-(v2))}
                                                            _ = v38
                                                            _ = v39
                                                        }
                                                        if true {
                                                            fmt.Println(v2, v2)
                                                        }
                                                        v1 = append(v1, 2147483647)
                                                    }
                                                    var v40 uint = 255
                                                    _ = v40
                                                } else {
                                                    for fi41 := 0; fi41 < 14; fi41++ {
                                                        switch (("foo bar" + "a") + "") {
                                                        case "foo bar":
                                                            v42 := GenS2{F0: 42}
                                                            _ = v42
                                                        case "":
                                                            if (true && true) {
                                                                var v43 uint8 = 230
                                                                _ = v43
                                                            }
                                                        default:
                                                            var v44 uint32 = 37
                                                            _ = v44
                                                        }
                                                        _ = fi41
                                                    }
                                                    fmt.Println(v3, fi26, v31)
                                                    for fi45 := 0; fi45 < 14; fi45++ {
                                                        var v46 int64 = 9223372036854775807
                                                        _ = fi45
                                                        _ = v46
                                                    }
                                                }
                                            }
                                        case "a":
                                            v24 -= 1
                                        case "":
                                            var v47 uint64 = 42
                                            _ = v47
                                        default:
                                            _ = len(v1)
                                        }
                                    }
                                    v3 -= -32768
                                }
                                fmt.Println(v24, v31)
                            } else {
                                if (!true) {
                                    fmt.Println(v24)
                                } else {
                                    fmt.Println(v2, fi26)
                                    switch "" {
                                    case "\n\t":
                                        if false {
                                            v48 := GenS2{F0: 65535}
                                            _ = v48
                                        } else {
                                            var v49 uint16 = v24
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v49
                                        }
                                    case "a":
                                        var v50 int8 = genFn0()
                                        _ = v50
                                    }
                                }
                                fmt.Println(v3, v31)
                            }
                        }
                        _ = v31
                    }
                    if ("foo bar" > (("a" + ("hello" + (("foo bar" + ("foo bar" + ("\n\t" + "\n\t"))) + ("foo bar" + "\n\t")))) + "\n\t")) {
                        switch fi26 {
                        case -433:
                            for fi51 := 0; fi51 < 3; fi51++ {
                                if false {
                                    if ((-128 >= 0) && true) {
                                        if len(v1) > 0 {
                                            v1[0] = -982
                                        }
                                        if (!true) {
                                            v52 := GenS2{F0: 0}
                                            fmt.Println(fi51, v3, fi51)
                                            switch ("foo bar" + "foo bar") {
                                            case "":
                                                if len(v1) > 0 {
                                                    v1[0] = 42
                                                }
                                            case "\n\t":
                                                fmt.Println(v3, v24, v2)
                                            default:
                                                var v53 float32 = float32(3.14159)
                                                _ = v53
                                            }
                                            _ = v52
                                        }
                                        v54 := GenS4{F0: 1, F1: ("foo bar" + "foo bar")}
                                        _ = v54
                                    } else {
                                        var v55 string = "foo bar"
                                        var v56 int = 737
                                        if ((-2.25 * ((1e10 + 1e10) / ((float32((1e10 - 3.14159)) + ((1.5 / 1.0) - ((((-2.25 / 3.14159) - 1.5) * 0.1) * 1.5))) * 3.14159))) == 0.1) {
                                            v57 := GenS4{F0: 1, F1: v55}
                                            fmt.Println(fi26)
                                            fmt.Println(fi26, v3)
                                            _ = v57
                                        } else {
                                            v58 := GenS2{F0: 1}
                                            var v59 int32 = -2147483648
                                            v60 := GenS5{F0: 2}
                                            switch fi51 {
                                            case 788:
                                                var v61 float64 = 0.0
                                                _ = v61
                                            case 0:
                                                v24 = 42
                                            default:
                                                fmt.Println(v3, v3, fi26)
                                            }
                                            switch (((v55 + v55) + ((v55 + ("a" + ("\n\t" + (v55 + v55)))) + v55)) + "\n\t") {
                                            case "hello":
                                                switch v55 {
                                                case "foo bar":
                                                    if false {
                                                        if (!false) {
                                                            var v62 uint16 = (v24 & (v24 % (((v24 % ((v24) | 1))) | 1)))
                                                            _ = v62
                                                        } else {
                                                            if len(v1) > 0 {
                                                                v1[0] = v56
                                                            }
                                                            _ = len(v1)
                                                        }
                                                    }
                                                    v24 = (v24 << ((2) & 63))
                                                    var v63 int64 = 63
                                                    _ = v63
                                                case "a":
                                                    if false {
                                                        var v64 uint = 523
                                                        v1 = append(v1, (-(v56)))
                                                        _ = v64
                                                    }
                                                default:
                                                    fmt.Println(fi51, v2, v2)
                                                }
                                            case "a":
                                                if (((genFn1() < (^(v24))) || false) || (v2 != v59)) {
                                                    switch v55 {
                                                    case "a":
                                                        fmt.Println(v3, fi51, v3)
                                                    }
                                                    _ = len(v1)
                                                    switch fi51 {
                                                    case 2147483647:
                                                        var v65 bool = false
                                                        _ = v65
                                                    default:
                                                        switch fi26 {
                                                        case 9223372036854775807:
                                                            if true {
                                                                v66 := GenS6{F0: -215, F1: (fi51 >> ((1) & 63))}
                                                                if (true && false) {
                                                                    if ((v24 % (((^(v24))) | 1)) > v24) {
                                                                        if false {
                                                                            v56 += fi26
                                                                            v1 = append(v1, (-(fi26)))
                                                                        }
                                                                        fmt.Println(v56, v59)
                                                                        if true {
                                                                            v67 := GenS5{F0: 1}
                                                                            _ = v67
                                                                        }
                                                                        v56 = (fi26 | 0)
                                                                    }
                                                                    var v68 string = v55
                                                                    if (979 >= 445) {
                                                                        var v69 int = 32767
                                                                        _ = v69
                                                                    } else {
                                                                        v55 = ""
                                                                    }
                                                                    _ = v68
                                                                } else {
                                                                    if ((v3 * v3) != v3) {
                                                                        var v70 uint64 = 2
                                                                        v24 = (v24 - v24)
                                                                        v3 += v3
                                                                        _ = v70
                                                                    }
                                                                    var v71 int16 = (v3 * v3)
                                                                    _ = v71
                                                                }
                                                                _ = v66
                                                            }
                                                        case 2147483647:
                                                            v72 := GenS1{F0: 1, F1: -1, F2: true}
                                                            _ = v72
                                                        case 48:
                                                            if (true && false) {
                                                                if true {
                                                                    v73 := GenS3{F0: -525, F1: v59}
                                                                    _ = v73
                                                                } else {
                                                                    switch v55 {
                                                                    case "\n\t":
                                                                        _ = len(v1)
                                                                    case "hello":
                                                                        v2 += int32(fi51)
                                                                    default:
                                                                        _ = len(v1)
                                                                    }
                                                                }
                                                                if (true || false) {
                                                                    if false {
                                                                        v74 := GenS1{F0: 18446744073709551615, F1: -9223372036854775808, F2: false}
                                                                        _ = v74
                                                                    }
                                                                    fmt.Println(v55, v2, fi26)
                                                                } else {
                                                                    v75 := GenS7{F0: "a", F1: 1, F2: v2}
                                                                    if true {
                                                                        v76 := GenS3{F0: 42, F1: -1}
                                                                        _ = len(v1)
                                                                        switch -9223372036854775808 {
                                                                        case 0:
                                                                            var v77 uint = 0
                                                                            _ = v77
                                                                        case -2147483648:
                                                                            if len(v1) > 0 {
                                                                                v1[0] = fi26
                                                                            }
                                                                        default:
                                                                            if true {
                                                                                v59 += (v2 % ((v2) | 1))
                                                                                v24 = v24
                                                                            }
                                                                        }
                                                                        _ = v76
                                                                    } else {
                                                                        if (35 == 0) {
                                                                            if (false && true) {
                                                                                var v78 uint32 = 2147483648
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = 0
                                                                                }
                                                                                _ = v78
                                                                            }
                                                                            v59 = (v2 * 42)
                                                                        }
                                                                        fmt.Println(v56)
                                                                        var v79 int64 = 9223372036854775807
                                                                        _ = v79
                                                                    }
                                                                    var v80 float32 = float32((1.5 / -2.25))
                                                                    _ = v75
                                                                    _ = v80
                                                                }
                                                                switch 225 {
                                                                case -2147483648:
                                                                    var v81 float64 = 3.14159
                                                                    if (fi26 == -820) {
                                                                        v59 += (v2 << ((593) & 63))
                                                                        var v82 bool = false
                                                                        _ = v82
                                                                    } else {
                                                                        _ = len(v1)
                                                                    }
                                                                    v24 = (v24 ^ 255)
                                                                    _ = v81
                                                                case -438:
                                                                    v24 -= genFn1()
                                                                }
                                                            } else {
                                                                if false {
                                                                    v83 := GenS6{F0: -1, F1: int(v2)}
                                                                    _ = v83
                                                                } else {
                                                                    v24 -= 1
                                                                    switch v55 {
                                                                    case "a":
                                                                        if len(v1) > 0 {
                                                                            v1[0] = fi51
                                                                        }
                                                                    case "":
                                                                        if (693 > 20) {
                                                                            switch v56 {
                                                                            case 660:
                                                                                v84 := GenS8{F0: 527}
                                                                                _ = v84
                                                                            case 4294967295:
                                                                                v3 = 0
                                                                            }
                                                                            var v85 uint8 = 2
                                                                            if true {
                                                                                switch v55 {
                                                                                case "hello":
                                                                                    if (v55 <= v55) {
                                                                                        if false {
                                                                                            switch fi26 {
                                                                                            case 127:
                                                                                                switch (((("hello" + (((((v55 + v55) + v55) + "a") + (("\n\t" + (v55 + ("\n\t" + v55))) + (v55 + v55))) + (v55 + (("" + v55) + "\n\t")))) + (v55 + "\n\t")) + "\n\t") + v55) {
                                                                                                case "\n\t":
                                                                                                    var v86 uint32 = 65535
                                                                                                    _ = v86
                                                                                                case "a":
                                                                                                    if (false && (true && (-108 <= genFn0()))) {
                                                                                                        var v87 float32 = 3.14159
                                                                                                        _ = v87
                                                                                                    }
                                                                                                    if (v3 == 633) {
                                                                                                        switch (^(fi51)) {
                                                                                                        case 640:
                                                                                                            v3 += (v3 & -1)
                                                                                                        case 9223372036854775807:
                                                                                                            if (0 <= -1) {
                                                                                                                var v88 uint8 = 99
                                                                                                                switch -776 {
                                                                                                                case 500:
                                                                                                                    if len(v1) >= 2 {
                                                                                                                        v1 = v1[1:len(v1)]
                                                                                                                    }
                                                                                                                case 127:
                                                                                                                    v3 = (v3 | (v3 << ((255) & 63)))
                                                                                                                case -9223372036854775808:
                                                                                                                    var v89 uint32 = 2147483648
                                                                                                                    _ = v89
                                                                                                                }
                                                                                                                _ = v88
                                                                                                            }
                                                                                                        case -722:
                                                                                                            if (2 <= 302) {
                                                                                                                v24 -= v24
                                                                                                                v90 := GenS2{F0: 200}
                                                                                                                v24 -= (^(v24))
                                                                                                                _ = v90
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        switch v55 {
                                                                                                        case "\n\t":
                                                                                                            if true {
                                                                                                                switch (-(fi51)) {
                                                                                                                case 2147483647:
                                                                                                                    v91 := GenS6{F0: (v3 % ((v3) | 1)), F1: -2147483648}
                                                                                                                    _ = v91
                                                                                                                default:
                                                                                                                    var v92 bool = true
                                                                                                                    _ = v92
                                                                                                                }
                                                                                                                if ((v2 / ((0) | 1)) > 284) {
                                                                                                                    switch v56 {
                                                                                                                    case 0:
                                                                                                                        if (v85 < 135) {
                                                                                                                            _ = len(v1)
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                        default:
                                                                                                            if false {
                                                                                                                v85 -= (^(v85))
                                                                                                                if len(v1) > 0 {
                                                                                                                    v1[0] = fi26
                                                                                                                }
                                                                                                                v1 = append(v1, (fi26 >> ((2) & 63)))
                                                                                                            } else {
                                                                                                                var v93 string = (v55 + ("a" + "foo bar"))
                                                                                                                var v94 int64 = 207
                                                                                                                _ = v93
                                                                                                                _ = v94
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    if (1 >= 0) {
                                                                                                        v95 := GenS6{F0: (v3 % (((v3 - -32768)) | 1)), F1: (v56 >> ((255) & 63))}
                                                                                                        v2 = v59
                                                                                                        _ = v95
                                                                                                    }
                                                                                                }
                                                                                            case -2147483648:
                                                                                                var v96 uint16 = v24
                                                                                                _ = v96
                                                                                            case 2147483647:
                                                                                                fmt.Println(v59, v85)
                                                                                            }
                                                                                            var v97 int8 = 1
                                                                                            var v98 int32 = v59
                                                                                            _ = v97
                                                                                            _ = v98
                                                                                        } else {
                                                                                            var v99 int32 = -2147483648
                                                                                            _ = v99
                                                                                        }
                                                                                    } else {
                                                                                        v55 = "a"
                                                                                        fmt.Println(v55, v3)
                                                                                        if false {
                                                                                            v55 = "hello"
                                                                                            fmt.Println(v3, v55)
                                                                                            v55 = (v55 + ("" + v55))
                                                                                        } else {
                                                                                            v100 := GenS8{F0: 340}
                                                                                            _ = v100
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v2 -= 2147483647
                                                                                }
                                                                            } else {
                                                                                if true {
                                                                                    v56 = fi26
                                                                                    _ = len(v1)
                                                                                    v59 = v59
                                                                                    var v101 int32 = v59
                                                                                    if false {
                                                                                        var v102 int = 2147483647
                                                                                        v103 := GenS2{F0: 252}
                                                                                        var v104 int = int(v59)
                                                                                        _ = v102
                                                                                        _ = v103
                                                                                        _ = v104
                                                                                    } else {
                                                                                        var v105 uint8 = (v85 / (((v85 >> ((723) & 63))) | 1))
                                                                                        var v106 uint64 = 2
                                                                                        _ = v105
                                                                                        _ = v106
                                                                                    }
                                                                                    _ = v101
                                                                                } else {
                                                                                    var v107 string = v55
                                                                                    var v108 uint = 894
                                                                                    if (uint8(v24) <= v85) {
                                                                                        var v109 int64 = -926
                                                                                        _ = v109
                                                                                    }
                                                                                    _ = v107
                                                                                    _ = v108
                                                                                }
                                                                                var v110 uint8 = (v85 >> ((65535) & 63))
                                                                                fmt.Println(v3, v55)
                                                                                _ = v110
                                                                            }
                                                                            _ = v85
                                                                        }
                                                                        if (v56 != (fi26 % ((fi51) | 1))) {
                                                                            var v111 uint16 = 1
                                                                            fmt.Println(v2, v56, v3)
                                                                            fmt.Println(v3)
                                                                            _ = v111
                                                                        } else {
                                                                            v112 := GenS9{F0: 2}
                                                                            fmt.Println(v59)
                                                                            v56 -= 1
                                                                            _ = v112
                                                                        }
                                                                        var v113 int8 = 92
                                                                        var v114 int16 = v3
                                                                        _ = len(v1)
                                                                        _ = v113
                                                                        _ = v114
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                            case "foo bar":
                                                v3 -= v3
                                            default:
                                                v115 := GenS3{F0: 1, F1: (-(v59))}
                                                _ = v115
                                            }
                                            _ = v58
                                            _ = v59
                                            _ = v60
                                        }
                                        v56 += (fi26 * (fi51 & fi26))
                                        fmt.Println(fi26)
                                        var v116 uint64 = 2
                                        v55 = v55
                                        v2 -= v2
                                        v116 = (v116 ^ (v116 ^ v116))
                                        _ = v116
                                        _ = v55
                                        _ = v56
                                    }
                                    if false {
                                        v117 := GenS9{F0: 213}
                                        v3 = (v3 - (v3 & (v3 ^ (-(v3)))))
                                        var v118 bool = (!true)
                                        if false {
                                            if ((false && v118) && v118) {
                                                var v119 float64 = -2.25
                                                _ = v119
                                            }
                                        }
                                        _ = v117
                                        _ = v118
                                    }
                                    fmt.Println(v3, v24, fi51)
                                }
                                if (1 == 18446744073709551615) {
                                    _ = len(v1)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v120 int = 537
                                    _ = v120
                                }
                                if (!(true || false)) {
                                    switch (fi51 % ((fi51) | 1)) {
                                    case 127:
                                        v2 -= (-(v2))
                                    case 187:
                                        var v121 int64 = 42
                                        _ = v121
                                    case -11:
                                        switch fi26 {
                                        case 2147483647:
                                            var v122 uint16 = 42
                                            _ = v122
                                        case 315:
                                            var v123 uint16 = (^(v24))
                                            _ = v123
                                        }
                                    }
                                    var v124 uint8 = 128
                                    _ = v124
                                }
                                if ("\n\t" <= "a") {
                                    v2 += (v2 | -470)
                                }
                                _ = fi51
                            }
                        case 918:
                            for fi125 := 0; fi125 < 5; fi125++ {
                                var v126 uint32 = 370
                                if true {
                                    fmt.Println(v3)
                                    v2 -= (v2 / ((v2) | 1))
                                } else {
                                    v2 += -2147483648
                                    v24 = genFn1()
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = fi125
                                _ = v126
                            }
                        case 32768:
                            for fi127 := 0; fi127 < 9; fi127++ {
                                var v128 uint8 = 42
                                _ = fi127
                                _ = v128
                            }
                        }
                        for fi129 := 0; fi129 < 5; fi129++ {
                            v2 += v2
                            var v130 string = "foo bar"
                            switch 840 {
                            case -1:
                                fmt.Println(v24, v2)
                            case 32768:
                                if true {
                                    v131 := GenS6{F0: 238, F1: fi129}
                                    fmt.Println(v130)
                                    _ = v131
                                } else {
                                    var v132 int8 = genFn0()
                                    v130 = ("" + (v130 + v130))
                                    v2 += 42
                                    _ = v132
                                }
                                fmt.Println(fi26, v24)
                            }
                            _ = fi129
                            _ = v130
                        }
                        for fi133 := 0; fi133 < 14; fi133++ {
                            if (true || (!(-216 <= fi133))) {
                                if false {
                                    v134 := GenS10{F0: -1, F1: 275}
                                    _ = v134
                                }
                                switch "\n\t" {
                                case "hello":
                                    if true {
                                        v135 := GenS11{F0: 100, F1: 127}
                                        v2 = v2
                                        _ = v135
                                    }
                                    v2 = (v2 + (v2 * (^(v2))))
                                case "\n\t":
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v136 uint64 = 42
                                        _ = v136
                                    default:
                                        fmt.Println(v3, fi133)
                                    }
                                default:
                                    var v137 int32 = v2
                                    _ = v137
                                }
                                var v138 int8 = genFn0()
                                _ = v138
                            }
                            var v139 uint = 671
                            _ = fi133
                            _ = v139
                        }
                    } else {
                        var v140 float64 = 1e10
                        _ = v140
                    }
                    _ = fi26
                }
            } else {
                fmt.Println(v2, v2)
            }
            var v141 uint = 18446744073709551615
            _ = v141
        }
        v1 = append(v1, 42)
        _ = v24
    }
    v142 := GenS12{F0: 32768, F1: 2, F2: 100}
    var v143 int32 = (v2 * -832)
    if (false && (4294967295 >= 0)) {
        v2 -= -138
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v144 float64 = 1.5
    var v145 float32 = 1.5
    switch ("a" + "a") {
    case "":
        if (false && true) {
            switch 213 {
            case 128:
                if (v145 <= (1e10 + 0.1)) {
                    v145 = (v145 / v145)
                } else {
                    v2 -= v143
                    v144 += v144
                }
            case 0:
                if ((true || true) || false) {
                    switch 127 {
                    case 4294967295:
                        var v146 int64 = 471
                        _ = v146
                    case -491:
                        v144 -= v144
                    case 128:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    for fi147 := 0; fi147 < 18; fi147++ {
                        v2 -= v2
                        if (1 != 42) {
                            v148 := GenS1{F0: 0, F1: -9223372036854775808, F2: (!(0 != 2))}
                            _ = v148
                        }
                        if false {
                            if true {
                                fmt.Println(v2, v2)
                            } else {
                                fmt.Println(v2, v143, v3)
                            }
                            if ((v143 << ((4294967295) & 63)) <= (v143 / ((v143) | 1))) {
                                var v149 bool = (v145 != 0.1)
                                v150 := GenS3{F0: -17, F1: v143}
                                var v151 uint = 259
                                _ = v149
                                _ = v150
                                _ = v151
                            } else {
                                v3 += v3
                                var v152 uint8 = 38
                                _ = v152
                            }
                        }
                        _ = fi147
                    }
                }
            case -142:
                var v153 uint64 = 398
                _ = v153
            default:
                fmt.Println(v144, v145, v145)
            }
            var v154 bool = true
            if ((false || v154) && false) {
                _ = len(v1)
            } else {
                var v155 string = ("foo bar" + ("\n\t" + "foo bar"))
                _ = v155
            }
            _ = v154
        }
    }
    fmt.Println(v2)
    fmt.Println(v2, v145, v143)
    fmt.Println(v2, v143, v3)
    var v156 uint64 = 1
    var v157 bool = (v3 < v3)
    v3 += (^(v3))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v143:", v143)
    fmt.Println("v144:", v144)
    fmt.Println("v145:", v145)
    fmt.Println("v156:", v156)
    fmt.Println("v157:", v157)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v142
    _ = v143
    _ = v144
    _ = v145
    _ = v156
    _ = v157
    _ = v2
    _ = v3
}
