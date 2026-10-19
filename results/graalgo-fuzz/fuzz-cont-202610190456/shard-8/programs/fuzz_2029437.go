package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint }
type GenS4 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 int; F1 uint16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 string; F1 uint16; F2 int16 }
type GenS10 struct { F0 int16; F1 int; F2 uint64 }
type GenS11 struct { F0 int8; F1 uint64 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint32; F1 int }

func main() {
    v1 := []string{"a", "foo bar", "foo bar"}
    v2 := []int{423, 401, -9223372036854775808, 32767}
    var v3 int = 127
    v4 := []uint32{1, 1, 462, 2147483648}
    var v5 int8 = 1
    var v6 int = v3
    v6 -= (v6 * (v3 - v3))
    var v7 uint16 = 862
    for fi8 := 0; fi8 < 1; fi8++ {
        v7 = v7
        v9 := GenS1{F0: v5, F1: 42}
        v10 := GenS2{F0: -2147483648, F1: (v7 % ((v7) | 1)), F2: 42}
        _ = fi8
        _ = v10
        _ = v9
    }
    v1 = append(v1, ("" + "foo bar"))
    v3 += 32767
    if true {
        switch ("" + "hello") {
        case "a":
            var v11 float64 = ((0.1 + 0.0) - 3.14159)
            _ = v11
        case "":
            switch "hello" {
            case "foo bar":
                v5 += -128
            }
        }
        for fi12 := 0; fi12 < 17; fi12++ {
            for fi13 := 0; fi13 < 10; fi13++ {
                v2 = append(v2, (^(v3)))
                _ = fi13
            }
            switch ("a" + "hello") {
            case "hello":
                for fi14 := 0; fi14 < 20; fi14++ {
                    fmt.Println(v5, fi12)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v3 = fi12
                    _ = fi14
                }
            default:
                fmt.Println(v6)
            }
            _ = fi12
        }
    }
    var v15 int16 = 134
    if false {
        v16 := GenS3{F0: 1, F1: 2, F2: 2}
        _ = v16
    } else {
        var v17 int = (-(v3))
        if len(v2) > 0 {
            v2[0] = v17
        }
        _ = v17
    }
    if ((!(0 < -255)) || ((0.0 + float64(0.1)) > float64(0.1))) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v5 = v5
    fmt.Println(v5, v7)
    switch "hello" {
    case "foo bar":
        switch v6 {
        case -1:
            switch -945 {
            case 1:
                var v18 int8 = 0
                _ = v18
            case -2147483648:
                switch "" {
                case "":
                    for fi19 := 0; fi19 < 17; fi19++ {
                        if (!(!true)) {
                            switch v3 {
                            case -9223372036854775808:
                                fmt.Println(v3, v15, v15)
                            }
                            fmt.Println(fi19, v5)
                            v7 = (v7 >> ((705) & 63))
                        } else {
                            if (true || true) {
                                if ((2147483647 == 65535) && false) {
                                    v20 := GenS2{F0: (v6 << ((65535) & 63)), F1: v7, F2: 0}
                                    var v21 string = "a"
                                    if false {
                                        fmt.Println(v21, v7)
                                        fmt.Println(v3)
                                    } else {
                                        for fi22 := 0; fi22 < 7; fi22++ {
                                            var v23 int8 = v5
                                            fmt.Println(v3, v3, v3)
                                            _ = fi22
                                            _ = v23
                                        }
                                        if true {
                                            v21 = "\n\t"
                                            var v24 float64 = 0.0
                                            v15 += (^(v15))
                                            _ = v24
                                        } else {
                                            v3 -= -777
                                            v15 += 0
                                            v21 = v21
                                        }
                                        switch v6 {
                                        case 128:
                                            v21 = v21
                                        case 297:
                                            if (-10 != 0) {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v25 float64 = ((0.1 + float64(((0.0 * (1.5 / 1e10)) - float32((((-2.25 / 0.1) * -2.25) - (0.1 / (1e10 - 0.1))))))) + 0.1)
                                                _ = v25
                                            }
                                        case -725:
                                            v3 = (v3 << ((18446744073709551615) & 63))
                                        default:
                                            switch v21 {
                                            case "":
                                                switch v6 {
                                                case 2147483647:
                                                    if false {
                                                        var v26 uint8 = 197
                                                        if (-9223372036854775808 <= int64(v6)) {
                                                            if true {
                                                                var v27 int64 = -871
                                                                _ = v27
                                                            }
                                                            var v28 bool = true
                                                            for fi29 := 0; fi29 < 10; fi29++ {
                                                                if v28 {
                                                                    if (v15 <= int16(v7)) {
                                                                        var v30 int16 = (v15 >> ((65535) & 63))
                                                                        _ = v30
                                                                    } else {
                                                                        v21 = ("a" + "")
                                                                        v3 = 2147483647
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                    }
                                                                }
                                                                if v28 {
                                                                    if (v28 && (-2.25 >= 0.0)) {
                                                                        fmt.Println(fi29)
                                                                        v31 := GenS3{F0: (^(v7)), F1: (v7 / ((v7) | 1)), F2: 65535}
                                                                        v28 = false
                                                                        _ = v31
                                                                    } else {
                                                                        fmt.Println(v5, fi29, v21)
                                                                        _ = len(v1)
                                                                        switch 42 {
                                                                        case -1:
                                                                            v7 = v7
                                                                        case 42:
                                                                            v32 := GenS3{F0: v7, F1: (v7 * (v7 * (^(v7)))), F2: 65535}
                                                                            _ = v32
                                                                        case 0:
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                        }
                                                                    }
                                                                    fmt.Println(v3, v6, fi19)
                                                                }
                                                                _ = fi29
                                                            }
                                                            _ = v28
                                                        } else {
                                                            for fi33 := 0; fi33 < 16; fi33++ {
                                                                v34 := GenS4{F0: 0, F1: 42, F2: (^(v7))}
                                                                if (false || (!true)) {
                                                                    v26 = (^(v26))
                                                                } else {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                }
                                                                v6 = (^(v6))
                                                                _ = fi33
                                                                _ = v34
                                                            }
                                                        }
                                                        v35 := GenS3{F0: v7, F1: v7, F2: 2}
                                                        _ = v26
                                                        _ = v35
                                                    }
                                                case 61:
                                                    switch (v3 % ((-1) | 1)) {
                                                    case 309:
                                                        if (2 > 2) {
                                                            v36 := GenS3{F0: 42, F1: v7, F2: 255}
                                                            _ = v36
                                                        }
                                                    case 128:
                                                        var v37 string = v21
                                                        _ = v37
                                                    }
                                                default:
                                                    _ = len(v4)
                                                }
                                            default:
                                                var v38 float32 = (0.1 + 3.14159)
                                                _ = v38
                                            }
                                        }
                                    }
                                    _ = v20
                                    _ = v21
                                }
                                switch "a" {
                                case "":
                                    v39 := GenS1{F0: 42, F1: 2147483647}
                                    _ = v39
                                case "foo bar":
                                    switch "a" {
                                    case "\n\t":
                                        if (522 < 1) {
                                            v5 = v5
                                            if ((3.14159 == 1e10) || true) {
                                                for fi40 := 0; fi40 < 4; fi40++ {
                                                    fmt.Println(v3, v15)
                                                    if len(v4) > 0 {
                                                        v4[0] = 14
                                                    }
                                                    if false {
                                                        var v41 int64 = 9223372036854775807
                                                        if true {
                                                            var v42 float64 = 0.0
                                                            v43 := GenS5{F0: 2, F1: ""}
                                                            _ = v42
                                                            _ = v43
                                                        }
                                                        v6 += (v6 & (v6 - -950))
                                                        _ = v41
                                                    }
                                                    _ = fi40
                                                }
                                            }
                                            v7 -= (v7 & (v7 + (v7 / (((v7 >> ((201) & 63))) | 1))))
                                        }
                                    case "":
                                        v44 := GenS5{F0: 42, F1: ""}
                                        _ = v44
                                    case "hello":
                                        var v45 bool = ((v15 << ((0) & 63)) == (v15 - (v15 / ((v15) | 1))))
                                        _ = v45
                                    }
                                case "hello":
                                    v46 := GenS6{F0: v3, F1: (v7 / (((v7 >> ((170) & 63))) | 1))}
                                    _ = v46
                                default:
                                    var v47 int8 = 127
                                    _ = v47
                                }
                            } else {
                                switch "foo bar" {
                                case "hello":
                                    v2 = append(v2, (v3 ^ -997))
                                }
                                v2 = append(v2, v6)
                            }
                            v48 := GenS7{F0: (^(v7))}
                            if true {
                                v5 = 1
                            } else {
                                if true {
                                    v49 := GenS8{F0: 42}
                                    for fi50 := 0; fi50 < 20; fi50++ {
                                        v5 = v5
                                        _ = fi50
                                    }
                                    for fi51 := 0; fi51 < 10; fi51++ {
                                        var v52 int16 = v15
                                        if (true || (!(false && false))) {
                                            if (251 == 666) {
                                                var v53 int64 = -1
                                                if false {
                                                    var v54 int = 1
                                                    _ = v54
                                                } else {
                                                    _ = len(v2)
                                                    v52 = 1000
                                                }
                                                switch (v3 - v3) {
                                                case -2147483648:
                                                    v2 = append(v2, 504)
                                                case 393:
                                                    var v55 uint16 = v7
                                                    _ = v55
                                                }
                                                var v56 int = int(v7)
                                                _ = v53
                                                _ = v56
                                            } else {
                                                switch -235 {
                                                case -304:
                                                    v57 := GenS6{F0: v6, F1: 255}
                                                    if false {
                                                        fmt.Println(v15, v15)
                                                    }
                                                    _ = v57
                                                case -808:
                                                    v3 = fi19
                                                case 9223372036854775807:
                                                    var v58 int = fi51
                                                    var v59 int32 = 1
                                                    if len(v1) > 0 {
                                                        v1[0] = ""
                                                    }
                                                    _ = v58
                                                    _ = v59
                                                default:
                                                    v7 -= v7
                                                }
                                                var v60 string = "\n\t"
                                                if false {
                                                    var v61 int = (fi19 * 42)
                                                    _ = v61
                                                }
                                                _ = v60
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = 2147483647
                                            }
                                            v15 = (v52 & -1)
                                        } else {
                                            v5 = v5
                                        }
                                        _ = fi51
                                        _ = v52
                                    }
                                    _ = v49
                                }
                                _ = len(v4)
                            }
                            _ = v48
                        }
                        _ = fi19
                    }
                case "hello":
                    v6 = 128
                }
            case 32768:
                for fi62 := 0; fi62 < 12; fi62++ {
                    switch -9223372036854775808 {
                    case -2147483648:
                        var v63 int8 = (^(v5))
                        _ = v63
                    default:
                        for fi64 := 0; fi64 < 14; fi64++ {
                            _ = len(v1)
                            _ = fi64
                        }
                    }
                    _ = fi62
                }
            default:
                v65 := GenS2{F0: -1, F1: (v7 + (v7 ^ 922)), F2: 147}
                _ = v65
            }
        }
    default:
        v66 := GenS9{F0: ("\n\t" + ("a" + (((("" + ("foo bar" + "")) + "a") + "a") + "a"))), F1: v7, F2: int16(v3)}
        _ = v66
    }
    var v67 uint64 = 232
    fmt.Println(v67, v5, v7)
    for fi68 := 0; fi68 < 18; fi68++ {
        v67 = 29
        switch -2147483648 {
        case 32767:
            for fi69 := 0; fi69 < 3; fi69++ {
                var v70 int16 = 1
                if (2 <= 204) {
                    _ = len(v1)
                    v4 = append(v4, 0)
                }
                switch "" {
                case "foo bar":
                    var v71 string = ((("foo bar" + "\n\t") + (("\n\t" + "hello") + (("" + (("" + "") + "\n\t")) + ("foo bar" + "")))) + "foo bar")
                    _ = v71
                case "hello":
                    v72 := GenS6{F0: 2147483647, F1: v7}
                    _ = v72
                }
                _ = fi69
                _ = v70
            }
        case 9223372036854775807:
            v3 = 0
        case 1:
            for fi73 := 0; fi73 < 17; fi73++ {
                if true {
                    var v74 bool = ((-2147483648 > -1) && (!false))
                    _ = v74
                }
                if (v67 != v67) {
                    v67 -= v67
                    if false {
                        v3 += 1
                        fmt.Println(fi73)
                    } else {
                        if false {
                            v7 = 32767
                        } else {
                            switch "hello" {
                            case "hello":
                                switch fi68 {
                                case 32767:
                                    var v75 uint64 = v67
                                    _ = v75
                                default:
                                    switch ("\n\t" + ("a" + "a")) {
                                    case "foo bar":
                                        v15 += v15
                                    }
                                }
                            case "":
                                switch (v6 % ((32767) | 1)) {
                                case 42:
                                    v67 = 1
                                }
                            default:
                                if len(v1) > 0 {
                                    v1[0] = ("foo bar" + "\n\t")
                                }
                            }
                            fmt.Println(fi68, v7)
                            fmt.Println(v67, fi73, fi68)
                        }
                    }
                } else {
                    var v76 uint8 = 127
                    var v77 uint32 = 2
                    _ = v76
                    _ = v77
                }
                if (-2.25 > (1e10 * 0.1)) {
                    switch "\n\t" {
                    case "a":
                        if (0.1 != ((-2.25 * float32((1e10 * 1e10))) * -2.25)) {
                            if (-803 <= 593) {
                                v78 := GenS8{F0: 182}
                                var v79 uint = 255
                                v80 := GenS1{F0: (v5 << ((174) & 63)), F1: -1}
                                _ = v78
                                _ = v79
                                _ = v80
                            }
                            if (false && ((("\n\t" + "a") + "a") > ("\n\t" + "\n\t"))) {
                                if true {
                                    if true {
                                        var v81 bool = ((-(v5)) != (v5 + -1))
                                        if (v81 || (42 <= v67)) {
                                            var v82 int32 = -1
                                            var v83 int32 = -1
                                            _ = v82
                                            _ = v83
                                        } else {
                                            var v84 bool = (!true)
                                            if true {
                                                var v85 uint32 = uint32(v7)
                                                v86 := GenS10{F0: v15, F1: (fi68 + (^(fi68))), F2: v67}
                                                _ = v85
                                                _ = v86
                                            }
                                            if v84 {
                                                v81 = (v81 || ((v7 >> ((2) & 63)) < 144))
                                                var v87 uint16 = v7
                                                fmt.Println(v81, fi73, v5)
                                                if v81 {
                                                    if (v67 >= (v67 ^ v67)) {
                                                        switch (("hello" + "foo bar") + "") {
                                                        case "hello":
                                                            v2 = append(v2, v6)
                                                        case "foo bar":
                                                            if v84 {
                                                                var v88 int16 = (v15 & (-(v15)))
                                                                _ = v88
                                                            } else {
                                                                if len(v4) > 0 {
                                                                    v4[0] = 4294967295
                                                                }
                                                            }
                                                        case "a":
                                                            if (-421 < 2147483647) {
                                                                var v89 bool = false
                                                                var v90 int32 = -310
                                                                _ = v89
                                                                _ = v90
                                                            }
                                                        default:
                                                            switch "" {
                                                            case "hello":
                                                                v91 := GenS11{F0: 124, F1: (v67 >> ((4294967295) & 63))}
                                                                if ((v67 << ((535) & 63)) >= v67) {
                                                                    if (1 < v15) {
                                                                        var v92 uint16 = (v87 ^ 255)
                                                                        if v81 {
                                                                            var v93 uint32 = 2
                                                                            if len(v2) > 0 {
                                                                                v2[0] = 42
                                                                            }
                                                                            _ = v93
                                                                        }
                                                                        v84 = v81
                                                                        _ = v92
                                                                    }
                                                                    fmt.Println(v7, v84, fi73)
                                                                } else {
                                                                    v7 = (v7 * 32767)
                                                                }
                                                                v94 := GenS10{F0: v15, F1: v6, F2: (v67 & (v67 << ((315) & 63)))}
                                                                _ = v91
                                                                _ = v94
                                                            default:
                                                                v5 = v5
                                                            }
                                                        }
                                                        var v95 uint16 = (v7 * (v7 - 65535))
                                                        v67 = (^(v67))
                                                        _ = v95
                                                    }
                                                } else {
                                                    var v96 uint = 42
                                                    if v81 {
                                                        if (!v81) {
                                                            v97 := GenS6{F0: 1, F1: 0}
                                                            var v98 uint = (v96 - v96)
                                                            _ = v97
                                                            _ = v98
                                                        } else {
                                                            v5 += 42
                                                            v87 -= 0
                                                            if v84 {
                                                                var v99 int16 = v15
                                                                if len(v4) > 0 {
                                                                    v4[0] = 518
                                                                }
                                                                _ = v99
                                                            }
                                                        }
                                                        v87 -= 65535
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                    _ = v96
                                                }
                                                var v100 float32 = float32(1.5)
                                                _ = v100
                                                _ = v87
                                            }
                                            v6 = (fi73 - (v6 / ((v3) | 1)))
                                            var v101 uint64 = v67
                                            _ = v101
                                            _ = v84
                                        }
                                        _ = v81
                                    } else {
                                        v102 := GenS10{F0: (v15 * (v15 + int16(v7))), F1: v6, F2: v67}
                                        if false {
                                            if len(v2) > 0 {
                                                v2[0] = (fi68 + (v6 + fi68))
                                            }
                                            fmt.Println(v15, fi73, v3)
                                        }
                                        v103 := GenS12{F0: 2147483647}
                                        _ = v102
                                        _ = v103
                                    }
                                }
                            }
                        } else {
                            switch "a" {
                            case "a":
                                if (v67 > (v67 << ((2) & 63))) {
                                    v67 += (v67 | 18446744073709551615)
                                    if (!false) {
                                        v67 -= (v67 / ((v67) | 1))
                                    }
                                    v104 := GenS13{F0: 0, F1: (fi73 << ((65535) & 63))}
                                    _ = v104
                                } else {
                                    fmt.Println(v67)
                                }
                            case "\n\t":
                                var v105 uint16 = 688
                                var v106 uint = 802
                                v105 = 803
                                if len(v1) > 0 {
                                    v1[0] = "hello"
                                }
                                v4 = append(v4, 1)
                                _ = v105
                                _ = v106
                            case "hello":
                                if len(v1) > 0 {
                                    v1[0] = (("a" + "") + "a")
                                }
                            default:
                                fmt.Println(fi73, v7, v6)
                            }
                        }
                    case "hello":
                        var v107 int16 = v15
                        _ = v107
                    default:
                        v108 := GenS3{F0: (v7 * (^(v7))), F1: 962, F2: 0}
                        _ = v108
                    }
                } else {
                    v67 += (^(v67))
                    v7 += (^(v7))
                    v6 = (^(fi73))
                }
                _ = fi73
            }
        }
        _ = fi68
    }
    var v109 uint = 92
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v109:", v109)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v109
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v7
}
