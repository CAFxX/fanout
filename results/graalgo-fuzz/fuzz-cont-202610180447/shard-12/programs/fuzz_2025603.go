package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int32 }
type GenS2 struct { F0 bool; F1 int; F2 bool }
type GenS3 struct { F0 int64; F1 bool; F2 bool }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 uint; F1 int16; F2 int8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 uint; F1 bool; F2 uint32 }
type GenS9 struct { F0 int32; F1 uint; F2 bool }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int16; F1 int64; F2 int8 }
type GenS13 struct { F0 uint; F1 int16 }
type GenS14 struct { F0 int64; F1 uint32; F2 int64 }
type GenS15 struct { F0 int; F1 uint; F2 int32 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 uint32; F1 string }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 int8; F1 int8 }
type GenS21 struct { F0 int; F1 string }
type GenS22 struct { F0 int }
type GenS23 struct { F0 uint64; F1 string }
type GenS24 struct { F0 int16; F1 uint64 }

func main() {
    var v1 uint = 255
    var v2 uint16 = 0
    var v3 uint16 = 571
    v4 := []string{"a", (("foo bar" + "a") + "hello"), "a", ("a" + "foo bar")}
    var v5 int8 = 0
    var v6 int = 996
    var v7 int16 = 0
    switch 667 {
    case -153:
        fmt.Println(v3, v6)
    case 2147483647:
        v8 := GenS1{F0: (v3 - (v3 * (v2 / (((v2 ^ 503)) | 1)))), F1: ((0.1 + (3.14159 - -2.25)) != 3.14159), F2: -9}
        _ = v8
    default:
        if true {
            v5 += int8(v6)
        } else {
            v9 := GenS2{F0: true, F1: (v6 - (^(v6))), F2: (846 <= uint64(v1))}
            var v10 uint = (v1 - v1)
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 float64 = -2.25
                switch v6 {
                case -29:
                    var v13 uint64 = 4294967295
                    _ = v13
                default:
                    for fi14 := 0; fi14 < 15; fi14++ {
                        v4 = append(v4, "")
                        if (!true) {
                            if true {
                                var v15 string = (("hello" + "\n\t") + "foo bar")
                                switch (v6 & -1) {
                                case 4294967295:
                                    v16 := GenS3{F0: -9223372036854775808, F1: false, F2: ((v7 + (v7 % ((32767) | 1))) != (v7 - v7))}
                                    _ = v16
                                default:
                                    switch (v6 | (fi14 + (^(fi14)))) {
                                    case 148:
                                        v17 := GenS4{F0: v3, F1: (v3 * 255)}
                                        _ = v17
                                    case 9223372036854775807:
                                        v5 += -128
                                        var v18 uint16 = (^(v2))
                                        v5 = 100
                                        _ = v18
                                    case 32768:
                                        v15 = v15
                                    default:
                                        var v19 uint64 = 2
                                        _ = v19
                                    }
                                }
                                _ = v15
                            }
                            var v20 int = (-(fi11))
                            if (3.14159 <= 0.1) {
                                v2 += (v3 >> ((v1) & 63))
                                var v21 uint64 = 147
                                _ = v21
                            }
                            _ = v20
                        }
                        fmt.Println(v2)
                        _ = fi14
                    }
                }
                for fi22 := 0; fi22 < 4; fi22++ {
                    v10 -= (v10 * 18446744073709551615)
                    switch "" {
                    case "\n\t":
                        if false {
                            var v23 uint8 = 42
                            var v24 uint16 = uint16(v1)
                            v25 := GenS1{F0: v3, F1: (1 != (v10 / (((v1 - (v1 ^ (v1 % ((0) | 1))))) | 1))), F2: int32(v7)}
                            _ = v23
                            _ = v24
                            _ = v25
                        } else {
                            v1 += v1
                            v26 := GenS5{F0: 255, F1: v7, F2: (v5 / ((v5) | 1))}
                            v4 = append(v4, ("hello" + ("a" + "\n\t")))
                            _ = v26
                        }
                    }
                    _ = fi22
                }
                _ = fi11
                _ = v12
            }
            _ = v10
            _ = v9
        }
    }
    v27 := GenS1{F0: v2, F1: (!true), F2: 846}
    fmt.Println(v6)
    var v28 uint16 = (v2 >> ((871) & 63))
    if true {
        switch (v6 ^ v6) {
        case 516:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            for fi29 := 0; fi29 < 6; fi29++ {
                switch "" {
                case "foo bar":
                    var v30 bool = false
                    _ = v30
                default:
                    v31 := GenS6{F0: fi29}
                    _ = v31
                }
                if true {
                    v2 -= 32767
                    v7 = (v7 - v7)
                    switch fi29 {
                    case -338:
                        switch 128 {
                        case 9223372036854775807:
                            if true {
                                fmt.Println(v6, v2)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                if false {
                                    v32 := GenS3{F0: -1, F1: (!true), F2: false}
                                    if len(v4) > 0 {
                                        v4[0] = "hello"
                                    }
                                    for fi33 := 0; fi33 < 10; fi33++ {
                                        var v34 int64 = 814
                                        fmt.Println(v7)
                                        _ = fi33
                                        _ = v34
                                    }
                                    _ = v32
                                } else {
                                    switch "" {
                                    case "foo bar":
                                        var v35 int32 = -952
                                        _ = v35
                                    case "a":
                                        var v36 int8 = v5
                                        _ = v36
                                    case "":
                                        v37 := GenS7{F0: 42, F1: 240}
                                        _ = v37
                                    }
                                }
                            } else {
                                v38 := GenS8{F0: (v1 ^ v1), F1: true, F2: uint32(v1)}
                                if true {
                                    switch "a" {
                                    case "\n\t":
                                        v28 = v3
                                    }
                                }
                                _ = v38
                            }
                        case 4294967295:
                            v1 = v1
                        case -1:
                            if (true && ((252 >= 4294967295) || (float64((1.5 / 1.5)) == 1.5))) {
                                switch -2147483648 {
                                case 42:
                                    v2 -= (v28 % (((^(v2))) | 1))
                                }
                                if (!(1e10 >= (0.1 + (0.0 * 3.14159)))) {
                                    v28 += 65535
                                }
                                var v39 uint8 = 0
                                _ = v39
                            }
                        }
                    case 128:
                        for fi40 := 0; fi40 < 3; fi40++ {
                            v3 = (v2 | (v3 + (v3 - (v2 / (((v28 + (^(v28)))) | 1)))))
                            switch ("\n\t" + (("foo bar" + "") + "")) {
                            case "foo bar":
                                fmt.Println(v3, v7, fi29)
                            default:
                                if true {
                                    switch int(v7) {
                                    case 4294967295:
                                        fmt.Println(v3, v5)
                                    case 770:
                                        v1 = (v1 << ((v1) & 63))
                                    default:
                                        if true {
                                            var v41 float32 = 3.14159
                                            if (42 == 2147483647) {
                                                switch (v6 + -1) {
                                                case 32768:
                                                    if (!(!false)) {
                                                        v41 += v41
                                                        v6 = (fi40 >> (((v1 % ((v1) | 1))) & 63))
                                                        var v42 string = "foo bar"
                                                        _ = v42
                                                    } else {
                                                        var v43 int8 = v5
                                                        v28 = 1
                                                        var v44 int16 = v7
                                                        _ = v43
                                                        _ = v44
                                                    }
                                                case 0:
                                                    var v45 float32 = ((v41 / v41) / v41)
                                                    _ = v45
                                                case -1:
                                                    switch (-(v6)) {
                                                    case 582:
                                                        var v46 uint16 = (^(v28))
                                                        _ = v46
                                                    case -345:
                                                        v6 = (v6 << ((4294967295) & 63))
                                                    }
                                                }
                                            }
                                            var v47 int64 = 42
                                            _ = v41
                                            _ = v47
                                        } else {
                                            if true {
                                                var v48 string = ("" + "hello")
                                                _ = v48
                                            } else {
                                                v28 += v2
                                            }
                                            v5 -= 42
                                            fmt.Println(v3)
                                        }
                                    }
                                    v5 += 42
                                } else {
                                    v6 += 95
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                            }
                            _ = fi40
                        }
                    case 32767:
                        var v49 int32 = -2147483648
                        _ = v49
                    default:
                        switch (v6 % ((32768) | 1)) {
                        case 1:
                            v5 -= v5
                        case -9223372036854775808:
                            if false {
                                for fi50 := 0; fi50 < 18; fi50++ {
                                    if ((int16(v6) > v7) && true) {
                                        if (true && true) {
                                            var v51 int64 = 0
                                            _ = v51
                                        } else {
                                            v4 = append(v4, "a")
                                            var v52 string = ("\n\t" + "")
                                            fmt.Println(v1, v52, v2)
                                            _ = v52
                                        }
                                    }
                                    v53 := GenS9{F0: -2147483648, F1: (v1 << (((v1 << (((v1 + (v1 / (((^(v1))) | 1)))) & 63))) & 63)), F2: true}
                                    _ = fi50
                                    _ = v53
                                }
                                switch v6 {
                                case -1:
                                    var v54 bool = false
                                    _ = v54
                                }
                                for fi55 := 0; fi55 < 14; fi55++ {
                                    switch 32767 {
                                    case 2147483647:
                                        var v56 float64 = (float64(-2.25) - 3.14159)
                                        _ = v56
                                    case -1:
                                        var v57 uint8 = 255
                                        _ = v57
                                    }
                                    var v58 uint64 = 18446744073709551615
                                    _ = fi55
                                    _ = v58
                                }
                            }
                        case -1:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        default:
                            if (0 >= (-(v7))) {
                                v59 := GenS8{F0: v1, F1: (0 < 713), F2: 0}
                                v6 = fi29
                                _ = v59
                            }
                        }
                    }
                } else {
                    switch fi29 {
                    case 1:
                        for fi60 := 0; fi60 < 11; fi60++ {
                            var v61 int = (^(v6))
                            _ = fi60
                            _ = v61
                        }
                    default:
                        if (true || ((1e10 * 1e10) > 1.5)) {
                            if (0 < v6) {
                                if (42 < -1) {
                                    switch (((("" + ("a" + "foo bar")) + "") + "\n\t") + (("foo bar" + ("foo bar" + ("foo bar" + (("" + ("hello" + "")) + "\n\t")))) + ((((((("hello" + "") + "foo bar") + "foo bar") + "a") + ("hello" + "a")) + "") + (("foo bar" + ("hello" + ((((("a" + "hello") + "a") + "foo bar") + ((("a" + "hello") + "foo bar") + (("\n\t" + (((("" + "") + "hello") + "foo bar") + "")) + (("\n\t" + "\n\t") + "\n\t")))) + "a"))) + "\n\t")))) {
                                    case "hello":
                                        v5 -= (v5 + v5)
                                    case "foo bar":
                                        _ = len(v4)
                                    }
                                }
                                switch "\n\t" {
                                case "foo bar":
                                    if false {
                                        switch ("foo bar" + "") {
                                        case "":
                                            fmt.Println(v28, v1, v5)
                                        case "hello":
                                            for fi62 := 0; fi62 < 3; fi62++ {
                                                var v63 uint32 = 2
                                                var v64 float64 = 1.5
                                                fmt.Println(fi62, fi62)
                                                if (1 <= 42) {
                                                    _ = len(v4)
                                                    fmt.Println(v7)
                                                } else {
                                                    if false {
                                                        switch "" {
                                                        case "a":
                                                            v64 = v64
                                                        case "\n\t":
                                                            var v65 bool = false
                                                            _ = v65
                                                        default:
                                                            var v66 uint8 = 128
                                                            _ = v66
                                                        }
                                                        _ = len(v4)
                                                    }
                                                }
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = fi62
                                                _ = v63
                                                _ = v64
                                            }
                                        case "foo bar":
                                            if ("hello" > "\n\t") {
                                                v5 += v5
                                                if len(v4) > 0 {
                                                    v4[0] = "foo bar"
                                                }
                                            }
                                        }
                                        var v67 uint = 18446744073709551615
                                        _ = v67
                                    } else {
                                        v28 += v2
                                        if true {
                                            var v68 float32 = (1e10 * -2.25)
                                            _ = len(v4)
                                            _ = v68
                                        } else {
                                            v69 := GenS9{F0: 2147483647, F1: 1, F2: ((^(v7)) == v7)}
                                            _ = v69
                                        }
                                        if true {
                                            v70 := GenS9{F0: -1, F1: v1, F2: (false && false)}
                                            fmt.Println(v5, v1, v2)
                                            _ = v70
                                        }
                                    }
                                case "hello":
                                    fmt.Println(v6, v7, fi29)
                                case "\n\t":
                                    for fi71 := 0; fi71 < 16; fi71++ {
                                        v2 = (v2 - (v28 % (((v2 & (v2 - v28))) | 1)))
                                        switch (^(v6)) {
                                        case -724:
                                            v72 := GenS10{F0: (v1 - (v1 + v1))}
                                            v73 := GenS11{F0: -2147483648}
                                            var v74 int32 = 42
                                            _ = v72
                                            _ = v73
                                            _ = v74
                                        case 42:
                                            fmt.Println(fi71)
                                        default:
                                            var v75 int = fi71
                                            _ = v75
                                        }
                                        fmt.Println(v1, v6, v5)
                                        var v76 int32 = -2147483648
                                        _ = fi71
                                        _ = v76
                                    }
                                default:
                                    var v77 int64 = -1
                                    _ = v77
                                }
                                for fi78 := 0; fi78 < 14; fi78++ {
                                    if ("" != ((("" + ("foo bar" + "a")) + "foo bar") + "")) {
                                        if ((uint64(v1) >= 0) || (1e10 != 1.5)) {
                                            if true {
                                                v4 = append(v4, "")
                                                var v79 int = (fi78 >> ((v1) & 63))
                                                if (42 == -9223372036854775808) {
                                                    if true {
                                                        var v80 uint16 = (^(v2))
                                                        fmt.Println(v6, v80, v3)
                                                        _ = v80
                                                    }
                                                    v81 := GenS5{F0: v1, F1: 32767, F2: v5}
                                                    fmt.Println(fi29, v1)
                                                    _ = v81
                                                }
                                                _ = v79
                                            } else {
                                                if false {
                                                    v4 = append(v4, (("hello" + "hello") + "a"))
                                                    v5 = (-(v5))
                                                }
                                            }
                                            v7 += v7
                                            v82 := GenS12{F0: -32768, F1: 9223372036854775807, F2: -33}
                                            if len(v4) > 0 {
                                                v4[0] = ("hello" + ("foo bar" + "\n\t"))
                                            }
                                            var v83 int8 = 109
                                            v2 = (^(v2))
                                            _ = v82
                                            _ = v83
                                        } else {
                                            var v84 float64 = 1e10
                                            fmt.Println(v3, v3)
                                            v6 = v6
                                            _ = v84
                                        }
                                    }
                                    var v85 int32 = -2147483648
                                    if true {
                                        if false {
                                            switch "hello" {
                                            case "hello":
                                                fmt.Println(v1, v3, v1)
                                            default:
                                                fmt.Println(v7, fi29)
                                            }
                                            var v86 float32 = 0.0
                                            _ = v86
                                        }
                                    }
                                    _ = fi78
                                    _ = v85
                                }
                            }
                        } else {
                            var v87 string = "foo bar"
                            _ = v87
                        }
                    }
                    fmt.Println(v3, v6, v2)
                }
                _ = fi29
            }
        }
        v5 = 72
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch v6 {
        case -9223372036854775808:
            if false {
                if (!true) {
                    if true {
                        switch (v6 >> ((255) & 63)) {
                        case 2147483647:
                            _ = len(v4)
                        }
                    } else {
                        if ("" == ("\n\t" + "foo bar")) {
                            _ = len(v4)
                            if (65535 >= 42) {
                                if len(v4) > 0 {
                                    v4[0] = ("foo bar" + "hello")
                                }
                            }
                        }
                        v28 = (^(v3))
                    }
                    v28 += (v3 & v28)
                } else {
                    v7 = (v7 | (v7 * v7))
                    v88 := GenS13{F0: v1, F1: -504}
                    var v89 uint16 = v2
                    _ = v88
                    _ = v89
                }
                if false {
                    v90 := GenS14{F0: 432, F1: 528, F2: 9223372036854775807}
                    _ = v90
                } else {
                    fmt.Println(v6, v7, v7)
                    fmt.Println(v7)
                }
                v2 -= (v28 % ((359) | 1))
            }
        case -316:
            v6 += -396
        case 1:
            for fi91 := 0; fi91 < 1; fi91++ {
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
                _ = fi91
            }
        default:
            if true {
                if len(v4) > 0 {
                    v4[0] = ("" + "foo bar")
                }
                v4 = append(v4, "a")
            }
        }
        switch 2147483647 {
        case -1:
            if true {
                if true {
                    v2 += 65535
                } else {
                    v6 += 32768
                    for fi92 := 0; fi92 < 18; fi92++ {
                        if true {
                            v5 = v5
                        }
                        fmt.Println(v1, v5)
                        v2 = (v28 & (v28 >> (((v1 / ((65535) | 1))) & 63)))
                        _ = fi92
                    }
                    v93 := GenS5{F0: v1, F1: v7, F2: v5}
                    _ = v93
                }
                v5 -= -49
                v6 -= v6
            }
        case 0:
            v3 -= 2
        default:
            for fi94 := 0; fi94 < 1; fi94++ {
                if false {
                    if (!(!(!false))) {
                        var v95 float32 = 1e10
                        fmt.Println(v28, v2, v28)
                        _ = v95
                    } else {
                        v4 = append(v4, ("" + ("foo bar" + "foo bar")))
                    }
                    fmt.Println(v2)
                    var v96 uint8 = 0
                    _ = v96
                }
                switch v6 {
                case -1:
                    var v97 int = v6
                    _ = v97
                case 1:
                    v98 := GenS5{F0: v1, F1: (v7 * (^(v7))), F2: -128}
                    _ = v98
                default:
                    if true {
                        v99 := GenS15{F0: 715, F1: 2, F2: -1}
                        for fi100 := 0; fi100 < 1; fi100++ {
                            if false {
                                var v101 uint = v1
                                _ = v101
                            }
                            if ((!true) && (1.5 >= 0.0)) {
                                var v102 string = "hello"
                                if (18446744073709551615 != 784) {
                                    if ((3.14159 + (1e10 + 1e10)) <= (-2.25 + (0.0 - 0.1))) {
                                        v103 := GenS16{F0: false}
                                        var v104 int64 = -9223372036854775808
                                        fmt.Println(v104, v6, v7)
                                        if false {
                                            v4 = append(v4, "foo bar")
                                            if (!false) {
                                                var v105 int64 = -1
                                                v3 += v3
                                                v105 += v104
                                                var v106 uint32 = 0
                                                var v107 int32 = 42
                                                _ = v105
                                                _ = v106
                                                _ = v107
                                            }
                                            var v108 int = fi100
                                            if (v102 == "a") {
                                                if false {
                                                    v28 -= 14
                                                    var v109 float64 = (((0.0 / 0.1) / ((1e10 / 0.1) + -2.25)) + 3.14159)
                                                    if false {
                                                        _ = len(v4)
                                                        if (!true) {
                                                            if true {
                                                                v110 := GenS2{F0: false, F1: (^(fi100)), F2: (true && false)}
                                                                v1 -= 1
                                                                var v111 int = v6
                                                                switch -80 {
                                                                case 32767:
                                                                    fmt.Println(v104, v28)
                                                                case -16:
                                                                    v112 := GenS17{F0: 103}
                                                                    _ = v112
                                                                }
                                                                _ = v110
                                                                _ = v111
                                                            } else {
                                                                v113 := GenS18{F0: 571, F1: ("\n\t" + ((("" + v102) + ("hello" + v102)) + ("hello" + (v102 + "foo bar"))))}
                                                                var v114 uint64 = 18446744073709551615
                                                                _ = v113
                                                                _ = v114
                                                            }
                                                            if (false && false) {
                                                                if (42 <= 107) {
                                                                    v115 := GenS18{F0: 65535, F1: ""}
                                                                    if ((("foo bar" + (("" + (v102 + "foo bar")) + v102)) <= (((v102 + ("" + ("\n\t" + ""))) + v102) + "")) || true) {
                                                                        if ((v7 << ((v1) & 63)) > -1) {
                                                                            var v116 int = -9223372036854775808
                                                                            if false {
                                                                                if true {
                                                                                    v5 -= (-(v5))
                                                                                    _ = len(v4)
                                                                                    v104 -= v104
                                                                                    var v117 float32 = (0.1 * 3.14159)
                                                                                    v109 -= float64((v117 * (v117 * v117)))
                                                                                    fmt.Println(fi94, v5, v104)
                                                                                    _ = v117
                                                                                }
                                                                                v118 := GenS10{F0: v1}
                                                                                _ = v118
                                                                            }
                                                                            _ = v116
                                                                        }
                                                                        var v119 uint64 = 1
                                                                        var v120 float32 = 0.1
                                                                        _ = v119
                                                                        _ = v120
                                                                    } else {
                                                                        switch ("" + v102) {
                                                                        case "hello":
                                                                            if len(v4) > 0 {
                                                                                v4[0] = ("a" + "foo bar")
                                                                            }
                                                                        case "foo bar":
                                                                            var v121 uint16 = (v2 ^ 2)
                                                                            _ = v121
                                                                        case "\n\t":
                                                                            var v122 int8 = v5
                                                                            _ = v122
                                                                        default:
                                                                            if (0.1 > 0.1) {
                                                                                if true {
                                                                                    switch -1 {
                                                                                    case 858:
                                                                                        v7 = v7
                                                                                    }
                                                                                    var v123 uint8 = 50
                                                                                    if (float64((1.5 / (((float32(v109) + (1.5 / (1e10 + (1e10 - (0.0 - 1e10))))) - -2.25) + -2.25))) != v109) {
                                                                                        if len(v4) > 0 {
                                                                                            v4[0] = v102
                                                                                        }
                                                                                        var v124 float64 = (v109 - v109)
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                        _ = v124
                                                                                    }
                                                                                    _ = v123
                                                                                }
                                                                                var v125 uint16 = 32767
                                                                                var v126 uint8 = 127
                                                                                _ = v125
                                                                                _ = v126
                                                                            } else {
                                                                                v127 := GenS19{F0: v3}
                                                                                var v128 float64 = (v109 - float64(((3.14159 + (0.0 - -2.25)) - -2.25)))
                                                                                if false {
                                                                                    if ((v3 ^ v2) >= v28) {
                                                                                        var v129 float64 = (float64(float32(0.0)) / (((v128 / 1e10) + 3.14159) - 0.0))
                                                                                        _ = v129
                                                                                    } else {
                                                                                        v1 -= v1
                                                                                        v1 = v1
                                                                                    }
                                                                                } else {
                                                                                    if true {
                                                                                        var v130 int16 = (-(v7))
                                                                                        if (v128 > ((0.1 + 1e10) + v109)) {
                                                                                            fmt.Println(v102, v5)
                                                                                            v131 := GenS20{F0: v5, F1: v5}
                                                                                            _ = v131
                                                                                        } else {
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = "foo bar"
                                                                                            }
                                                                                        }
                                                                                        var v132 string = ("" + (("a" + v102) + "a"))
                                                                                        _ = v130
                                                                                        _ = v132
                                                                                    }
                                                                                    _ = len(v4)
                                                                                }
                                                                                _ = v127
                                                                                _ = v128
                                                                            }
                                                                            v5 += v5
                                                                            if false {
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = v102
                                                                                }
                                                                            }
                                                                            var v133 float64 = 1e10
                                                                            _ = v133
                                                                        }
                                                                        v5 += (v5 * (v5 << (((v1 + (^(v1)))) & 63)))
                                                                        fmt.Println(v108)
                                                                        var v134 int64 = v104
                                                                        _ = v134
                                                                    }
                                                                    _ = v115
                                                                } else {
                                                                    v4 = append(v4, v102)
                                                                    if true {
                                                                        if (2 >= uint8(v2)) {
                                                                            v1 = (v1 ^ (v1 << ((1) & 63)))
                                                                            var v135 bool = false
                                                                            _ = v135
                                                                        } else {
                                                                            fmt.Println(v5, v5)
                                                                            var v136 string = "foo bar"
                                                                            if true {
                                                                                var v137 float64 = v109
                                                                                _ = v137
                                                                            }
                                                                            _ = v136
                                                                        }
                                                                        var v138 uint64 = 42
                                                                        v102 = (v102 + ("a" + v102))
                                                                        if len(v4) > 0 {
                                                                            v4[0] = v102
                                                                        }
                                                                        _ = v138
                                                                    } else {
                                                                        var v139 uint16 = (v28 << (((v1 * v1)) & 63))
                                                                        _ = len(v4)
                                                                        _ = v139
                                                                    }
                                                                    var v140 string = v102
                                                                    _ = v140
                                                                }
                                                                if false {
                                                                    v4 = append(v4, ("" + "hello"))
                                                                    v141 := GenS21{F0: -747, F1: "a"}
                                                                    v109 += v109
                                                                    _ = v141
                                                                }
                                                            }
                                                            if (!(!false)) {
                                                                v4 = append(v4, v102)
                                                                fmt.Println(fi94, v3, v3)
                                                            }
                                                            if (-1 > int32(v5)) {
                                                                fmt.Println(v2, v102)
                                                            } else {
                                                                var v142 string = v102
                                                                switch ((((v142 + v102) + "foo bar") + v142) + "") {
                                                                case "foo bar":
                                                                    if true {
                                                                        var v143 uint = v1
                                                                        _ = v143
                                                                    }
                                                                case "hello":
                                                                    v108 += 32767
                                                                default:
                                                                    if (!false) {
                                                                        var v144 int8 = 42
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                        var v145 uint16 = (v2 & (^(v2)))
                                                                        _ = v144
                                                                        _ = v145
                                                                    }
                                                                }
                                                                _ = v142
                                                            }
                                                            v104 += 9223372036854775807
                                                            v6 = (-(v108))
                                                        } else {
                                                            var v146 int = fi94
                                                            _ = v146
                                                        }
                                                        if false {
                                                            if true {
                                                                v147 := GenS22{F0: (fi94 | int(v104))}
                                                                v3 = v3
                                                                _ = v147
                                                            }
                                                        }
                                                    } else {
                                                        fmt.Println(v5, v28)
                                                    }
                                                    _ = v109
                                                }
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = v108
                                        }
                                        var v148 uint8 = 255
                                        _ = v103
                                        _ = v104
                                        _ = v148
                                    } else {
                                        _ = len(v4)
                                        if (!false) {
                                            switch v102 {
                                            case "a":
                                                if (42 <= v3) {
                                                    if (false && ((^(v1)) > (v1 ^ v1))) {
                                                        var v149 float64 = 3.14159
                                                        var v150 bool = true
                                                        _ = len(v4)
                                                        _ = v149
                                                        _ = v150
                                                    }
                                                    v151 := GenS1{F0: v2, F1: false, F2: int32(v1)}
                                                    var v152 uint32 = 2147483648
                                                    var v153 uint8 = 0
                                                    _ = v151
                                                    _ = v152
                                                    _ = v153
                                                }
                                            }
                                        }
                                        v5 = v5
                                    }
                                }
                                _ = v102
                            }
                            _ = fi100
                        }
                        _ = v99
                    } else {
                        v154 := GenS23{F0: 390, F1: ("" + "hello")}
                        _ = v154
                    }
                }
                for fi155 := 0; fi155 < 15; fi155++ {
                    v156 := GenS24{F0: 1, F1: 1}
                    v157 := GenS12{F0: (-(v7)), F1: -9223372036854775808, F2: (v5 | (-(v5)))}
                    if true {
                        if (false || (42 <= v1)) {
                            fmt.Println(v3, v1)
                            fmt.Println(v5)
                        } else {
                            switch v6 {
                            case 0:
                                var v158 int64 = 42
                                _ = v158
                            case -2147483648:
                                v3 = 32767
                            case 42:
                                var v159 bool = (0 != -396)
                                _ = v159
                            }
                        }
                        var v160 bool = false
                        _ = v160
                    }
                    _ = fi155
                    _ = v156
                    _ = v157
                }
                _ = fi94
            }
        }
    }
    fmt.Println(v7, v6, v5)
    switch (v6 ^ (v6 << (((v1 | (v1 >> ((910) & 63)))) & 63))) {
    case 969:
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
    case -1:
        var v161 int = (-(v6))
        _ = v161
    case 42:
        var v162 bool = false
        _ = v162
    default:
        v163 := GenS19{F0: (v3 >> (((v1 / (((^(v1))) | 1))) & 63))}
        _ = v163
    }
    if (!false) {
        var v164 uint = (v1 | (^(v1)))
        _ = v164
    } else {
        _ = len(v4)
    }
    fmt.Println(v7, v7, v3)
    var v165 int32 = 0
    fmt.Println(v6)
    for fi166 := 0; fi166 < 1; fi166++ {
        var v167 int64 = 42
        _ = fi166
        _ = v167
    }
    fmt.Println("v1:", v1)
    fmt.Println("v165:", v165)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v165
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
