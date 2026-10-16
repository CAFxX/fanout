package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int64 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 int8; F1 string; F2 uint32 }
type GenS6 struct { F0 bool; F1 uint16; F2 int64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int64; F1 uint8; F2 bool }
type GenS10 struct { F0 string; F1 uint64; F2 uint32 }
type GenS11 struct { F0 int16; F1 uint }
type GenS12 struct { F0 int64; F1 string }

func main() {
    var v1 int16 = 978
    var v2 int8 = 49
    var v3 float64 = 1e10
    v4 := []string{((("" + "hello") + ("foo bar" + "a")) + (("foo bar" + ("\n\t" + ("foo bar" + "\n\t"))) + "\n\t")), "\n\t"}
    if len(v4) > 0 {
        v4[0] = ""
    }
    if true {
        v1 -= v1
    } else {
        var v5 uint16 = 315
        _ = v5
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (618 != 32768) {
        v1 += (v1 * -670)
        v1 -= (v1 << ((932) & 63))
    }
    var v6 int8 = v2
    var v7 int16 = 0
    v8 := GenS1{F0: 255, F1: 2, F2: -1}
    v6 = 0
    v3 = v3
    v6 -= (v2 % ((v2) | 1))
    _ = len(v4)
    v6 += v2
    if (0 > 4294967295) {
        v6 = (v6 - (v6 >> ((1) & 63)))
        if true {
            if (!true) {
                if ((!false) && true) {
                    v9 := GenS2{F0: -957, F1: 42}
                    for fi10 := 0; fi10 < 18; fi10++ {
                        var v11 uint32 = 65535
                        _ = fi10
                        _ = v11
                    }
                    _ = v9
                } else {
                    v6 = (v6 / ((v6) | 1))
                    switch -309 {
                    case 127:
                        switch -920 {
                        case 4294967295:
                            v2 += v2
                        }
                    case -16:
                        for fi12 := 0; fi12 < 20; fi12++ {
                            var v13 int = 9223372036854775807
                            for fi14 := 0; fi14 < 1; fi14++ {
                                switch fi12 {
                                case -486:
                                    var v15 uint = 255
                                    _ = v15
                                case -2147483648:
                                    fmt.Println(v1, fi12, fi12)
                                }
                                _ = len(v4)
                                _ = fi14
                            }
                            _ = fi12
                            _ = v13
                        }
                    case 1:
                        v1 -= (v1 >> ((4294967295) & 63))
                    default:
                        if true {
                            switch 4294967295 {
                            case -73:
                                if (-2147483648 >= 2147483647) {
                                    for fi16 := 0; fi16 < 5; fi16++ {
                                        if ((("" + "a") < ("" + ("foo bar" + "hello"))) && (717 < 354)) {
                                            var v17 uint32 = 4294967295
                                            if true {
                                                switch int(v7) {
                                                case 128:
                                                    if (!(!((-86 >= 9223372036854775807) && false))) {
                                                        var v18 float64 = ((v3 * v3) * (((((float64(3.14159) / float64(float32((-2.25 / (((v3 / -2.25) - v3) + v3))))) - v3) * v3) / v3) * ((float64(3.14159) + v3) + 1e10)))
                                                        _ = v18
                                                    } else {
                                                        var v19 uint8 = 75
                                                        switch "a" {
                                                        case "foo bar":
                                                            v2 += v6
                                                        case "a":
                                                            if false {
                                                                switch fi16 {
                                                                case -100:
                                                                    var v20 int = fi16
                                                                    _ = v20
                                                                default:
                                                                    var v21 int = (fi16 | fi16)
                                                                    _ = v21
                                                                }
                                                                switch (fi16 ^ fi16) {
                                                                case 1:
                                                                    if len(v4) > 0 {
                                                                        v4[0] = "hello"
                                                                    }
                                                                case -9223372036854775808:
                                                                    v22 := GenS1{F0: 42, F1: 984, F2: -598}
                                                                    _ = v22
                                                                case 787:
                                                                    var v23 uint64 = 637
                                                                    _ = v23
                                                                }
                                                            }
                                                        default:
                                                            v2 = v6
                                                        }
                                                        _ = v19
                                                    }
                                                case -2147483648:
                                                    v24 := GenS3{F0: 1}
                                                    _ = v24
                                                case 9223372036854775807:
                                                    if (v3 != (0.0 / (v3 * 1e10))) {
                                                        v25 := GenS4{F0: 0, F1: 255}
                                                        for fi26 := 0; fi26 < 7; fi26++ {
                                                            v27 := GenS3{F0: 42}
                                                            _ = fi26
                                                            _ = v27
                                                        }
                                                        _ = v25
                                                    }
                                                }
                                                v28 := GenS5{F0: (v6 >> ((2) & 63)), F1: "hello", F2: v17}
                                                v3 += (float64((0.0 / ((float32(v3) / (0.1 * 1e10)) - float32((v3 / 3.14159))))) / ((1e10 + (3.14159 * v3)) * v3))
                                                _ = v28
                                            } else {
                                                v2 += (v6 * v6)
                                            }
                                            v29 := GenS1{F0: 884, F1: 2, F2: -9223372036854775808}
                                            _ = v17
                                            _ = v29
                                        } else {
                                            fmt.Println(v2)
                                        }
                                        for fi30 := 0; fi30 < 6; fi30++ {
                                            if (false || true) {
                                                if true {
                                                    fmt.Println(v1, fi30, fi16)
                                                    v1 = (v1 | (v7 ^ 810))
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        var v31 int = (fi16 * fi30)
                                                        _ = v31
                                                    case "a":
                                                        fmt.Println(v1)
                                                        var v32 int32 = 0
                                                        _ = v32
                                                    case "":
                                                        var v33 uint8 = 127
                                                        var v34 int8 = 127
                                                        _ = v33
                                                        _ = v34
                                                    }
                                                    if (v3 == (v3 * v3)) {
                                                        if true {
                                                            var v35 bool = true
                                                            _ = v35
                                                        }
                                                        v4 = append(v4, "hello")
                                                        v6 += -13
                                                        var v36 int64 = -864
                                                        _ = v36
                                                    }
                                                } else {
                                                    v4 = append(v4, (("hello" + "\n\t") + "hello"))
                                                    if false {
                                                        v2 += (v6 % ((-1) | 1))
                                                        v6 = (v6 - (v2 << ((13) & 63)))
                                                        var v37 uint8 = 116
                                                        _ = v37
                                                    } else {
                                                        switch ("hello" + "a") {
                                                        case "foo bar":
                                                            switch (fi30 | fi16) {
                                                            case 32767:
                                                                fmt.Println(v3, v6)
                                                            case 398:
                                                                v3 = v3
                                                            }
                                                        case "hello":
                                                            var v38 int = (fi30 >> ((4294967295) & 63))
                                                            _ = v38
                                                        case "\n\t":
                                                            if len(v4) > 0 {
                                                                v4[0] = "\n\t"
                                                            }
                                                        default:
                                                            v7 += v1
                                                        }
                                                    }
                                                    var v39 uint32 = 2147483647
                                                    _ = v39
                                                }
                                            }
                                            v1 = v7
                                            v40 := GenS6{F0: (-1 > 715), F1: 32767, F2: 0}
                                            v41 := GenS2{F0: -1, F1: 0}
                                            switch fi16 {
                                            case 759:
                                                var v42 float32 = (float32((3.14159 + v3)) / (0.1 * 0.1))
                                                _ = v42
                                            }
                                            v6 = v2
                                            _ = fi30
                                            _ = v40
                                            _ = v41
                                        }
                                        if (true && true) {
                                            v43 := GenS5{F0: 42, F1: "\n\t", F2: 878}
                                            if (v7 > (-(v1))) {
                                                var v44 int64 = -9223372036854775808
                                                switch "hello" {
                                                case "\n\t":
                                                    fmt.Println(v6, v1, v3)
                                                case "":
                                                    fmt.Println(v6, v7, v6)
                                                case "foo bar":
                                                    var v45 bool = (true || true)
                                                    _ = v45
                                                }
                                                _ = v44
                                            } else {
                                                for fi46 := 0; fi46 < 19; fi46++ {
                                                    v47 := GenS2{F0: -107, F1: 4294967295}
                                                    _ = fi46
                                                    _ = v47
                                                }
                                                if (0 >= 0) {
                                                    v7 -= (v1 % ((v1) | 1))
                                                }
                                                switch 128 {
                                                case 0:
                                                    var v48 uint = 4294967295
                                                    _ = v48
                                                case 128:
                                                    v6 -= (v2 & v2)
                                                case -255:
                                                    if len(v4) > 0 {
                                                        v4[0] = "hello"
                                                    }
                                                }
                                            }
                                            for fi49 := 0; fi49 < 15; fi49++ {
                                                if (2147483648 > 886) {
                                                    if (true && (1 != v6)) {
                                                        if (!((0.0 > 0.0) && (v1 <= (v7 - (v1 % (((-(v7))) | 1)))))) {
                                                            var v50 int16 = (v1 >> ((289) & 63))
                                                            if (float64(-2.25) == ((v3 + -2.25) - (v3 + (((v3 * (float64(0.0) + v3)) * 1e10) + (v3 * float64(-2.25)))))) {
                                                                var v51 int = -334
                                                                v51 -= v51
                                                                var v52 int = (fi16 | (fi16 | fi16))
                                                                _ = v51
                                                                _ = v52
                                                            } else {
                                                                var v53 uint32 = 0
                                                                _ = v53
                                                            }
                                                            _ = v50
                                                        } else {
                                                            var v54 uint8 = 235
                                                            var v55 int8 = (v2 << ((255) & 63))
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                        if len(v4) > 0 {
                                                            v4[0] = (("foo bar" + (("hello" + "\n\t") + ("\n\t" + (("\n\t" + ("" + ("hello" + ((("a" + "hello") + ("a" + ("hello" + ((("" + "") + "") + (("" + "hello") + ("" + "")))))) + ("" + "foo bar"))))) + "foo bar")))) + "hello")
                                                        }
                                                    }
                                                    v7 -= v1
                                                    v1 += -32768
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v56 := GenS6{F0: false, F1: 255, F2: -781}
                                                    v7 = v7
                                                    _ = v56
                                                }
                                                _ = fi49
                                            }
                                            _ = v43
                                        }
                                        _ = fi16
                                    }
                                    v57 := GenS2{F0: -2147483648, F1: 0}
                                    _ = v57
                                }
                            }
                            if false {
                                _ = len(v4)
                                v1 += v7
                                if true {
                                    var v58 uint16 = 42
                                    var v59 int16 = (v1 + v7)
                                    _ = v58
                                    _ = v59
                                } else {
                                    var v60 float64 = (v3 / v3)
                                    if false {
                                        switch ("" + (("foo bar" + "") + "\n\t")) {
                                        case "foo bar":
                                            fmt.Println(v7, v6, v6)
                                        case "":
                                            for fi61 := 0; fi61 < 3; fi61++ {
                                                v1 += (v7 % (((v1 + (^(v7)))) | 1))
                                                _ = len(v4)
                                                _ = fi61
                                            }
                                        case "\n\t":
                                            fmt.Println(v3, v2)
                                        default:
                                            v6 -= v6
                                        }
                                        if true {
                                            v62 := GenS2{F0: 42, F1: 743}
                                            for fi63 := 0; fi63 < 17; fi63++ {
                                                for fi64 := 0; fi64 < 9; fi64++ {
                                                    v4 = append(v4, "\n\t")
                                                    _ = fi64
                                                }
                                                _ = fi63
                                            }
                                            if (!true) {
                                                if (false && ((3.14159 + 0.1) == ((1e10 - ((0.1 / -2.25) / float32(v3))) * 0.0))) {
                                                    var v65 int64 = -1
                                                    _ = v65
                                                } else {
                                                    switch 514 {
                                                    case 32768:
                                                        for fi66 := 0; fi66 < 20; fi66++ {
                                                            v67 := GenS2{F0: 514, F1: 1}
                                                            fmt.Println(fi66, fi66, v3)
                                                            var v68 uint32 = 2147483647
                                                            _ = fi66
                                                            _ = v67
                                                            _ = v68
                                                        }
                                                    default:
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    }
                                                    var v69 int = -2147483648
                                                    if true {
                                                        if (v69 > 0) {
                                                            fmt.Println(v2, v69, v1)
                                                        } else {
                                                            var v70 string = ("a" + ("a" + ("a" + "")))
                                                            _ = v70
                                                        }
                                                    }
                                                    _ = v69
                                                }
                                                var v71 uint8 = 42
                                                _ = v71
                                            }
                                            _ = v62
                                        } else {
                                            switch "" {
                                            case "a":
                                                var v72 int = 1
                                                _ = v72
                                            case "foo bar":
                                                for fi73 := 0; fi73 < 18; fi73++ {
                                                    var v74 float64 = ((((-2.25 - ((v60 + (0.1 + v3)) * (v3 + (1e10 - v60)))) - (v60 / v60)) - (float64(3.14159) / -2.25)) + v60)
                                                    v2 -= 35
                                                    _ = fi73
                                                    _ = v74
                                                }
                                            default:
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                            switch 874 {
                                            case 4294967295:
                                                v75 := GenS3{F0: -9223372036854775808}
                                                _ = v75
                                            }
                                        }
                                        var v76 bool = false
                                        _ = v76
                                    } else {
                                        v3 += v60
                                        if len(v4) > 0 {
                                            v4[0] = ("hello" + "a")
                                        }
                                        v7 += (v1 + v7)
                                    }
                                    _ = v60
                                }
                            }
                            if false {
                                v77 := GenS7{F0: 523}
                                _ = v77
                            } else {
                                if true {
                                    for fi78 := 0; fi78 < 14; fi78++ {
                                        if true {
                                            for fi79 := 0; fi79 < 14; fi79++ {
                                                var v80 float64 = float64((float32(v3) * (1.5 - ((((1e10 / ((((0.0 + 0.1) / 1.0) - 1.5) / 3.14159)) + float32((-2.25 - float64((((1e10 * 1e10) + float32(v3)) * (0.0 + float32(v3))))))) / (float32(0.1) + ((0.1 / (0.1 + (0.1 * 0.1))) * 1e10))) / (((1e10 * (3.14159 + 3.14159)) / 0.1) - 0.0)))))
                                                var v81 uint64 = 1
                                                _ = fi79
                                                _ = v80
                                                _ = v81
                                            }
                                            switch "foo bar" {
                                            case "a":
                                                v82 := GenS2{F0: 2147483647, F1: 65535}
                                                _ = v82
                                            case "":
                                                switch ("a" + "") {
                                                case "foo bar":
                                                    for fi83 := 0; fi83 < 19; fi83++ {
                                                        var v84 float32 = 0.1
                                                        v3 += float64(v84)
                                                        _ = fi83
                                                        _ = v84
                                                    }
                                                case "\n\t":
                                                    var v85 float64 = (v3 + -2.25)
                                                    _ = v85
                                                case "":
                                                    var v86 uint8 = 255
                                                    _ = v86
                                                }
                                            case "hello":
                                                var v87 uint32 = 2147483647
                                                _ = v87
                                            }
                                        }
                                        var v88 bool = (-1 == int32(fi78))
                                        _ = len(v4)
                                        _ = fi78
                                        _ = v88
                                    }
                                } else {
                                    fmt.Println(v7)
                                }
                            }
                        } else {
                            _ = len(v4)
                            for fi89 := 0; fi89 < 1; fi89++ {
                                var v90 int32 = -2147483648
                                _ = fi89
                                _ = v90
                            }
                            for fi91 := 0; fi91 < 14; fi91++ {
                                v92 := GenS8{F0: v1}
                                fmt.Println(v7, v6, v6)
                                if false {
                                    if false {
                                        for fi93 := 0; fi93 < 18; fi93++ {
                                            switch 618 {
                                            case -761:
                                                v6 = v2
                                            }
                                            _ = fi93
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = "a"
                                        }
                                        switch (fi91 >> ((689) & 63)) {
                                        case -9223372036854775808:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        case 0:
                                            var v94 int8 = (v2 ^ v2)
                                            _ = v94
                                        case 1:
                                            if false {
                                                v95 := GenS9{F0: 0, F1: 0, F2: true}
                                                if (!false) {
                                                    var v96 float64 = 1e10
                                                    if true {
                                                        v6 += (v2 & (v2 << ((480) & 63)))
                                                        if false {
                                                            v97 := GenS2{F0: -2147483648, F1: 255}
                                                            for fi98 := 0; fi98 < 8; fi98++ {
                                                                if (v7 >= v1) {
                                                                    v7 = v1
                                                                }
                                                                v6 -= v2
                                                                _ = fi98
                                                            }
                                                            _ = v97
                                                        } else {
                                                            fmt.Println(v96)
                                                            for fi99 := 0; fi99 < 9; fi99++ {
                                                                fmt.Println(v1, v1, v6)
                                                                var v100 float32 = ((0.0 + ((0.0 * float32(v96)) - (1.5 + -2.25))) + 0.1)
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                _ = fi99
                                                                _ = v100
                                                            }
                                                        }
                                                        var v101 int8 = v2
                                                        _ = v101
                                                    } else {
                                                        v102 := GenS8{F0: -32768}
                                                        var v103 bool = (true || true)
                                                        _ = v102
                                                        _ = v103
                                                    }
                                                    _ = v96
                                                }
                                                v1 = (v7 / ((v1) | 1))
                                                _ = v95
                                            } else {
                                                for fi104 := 0; fi104 < 6; fi104++ {
                                                    if (94 > 65535) {
                                                        if true {
                                                            var v105 uint16 = 805
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            var v106 uint = 255
                                                            _ = v105
                                                            _ = v106
                                                        }
                                                        if (-2.25 <= (0.1 / 1.0)) {
                                                            v107 := GenS6{F0: false, F1: 742, F2: -9223372036854775808}
                                                            fmt.Println(fi91, v1, v7)
                                                            v3 = v3
                                                            _ = v107
                                                        }
                                                    } else {
                                                        _ = len(v4)
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        v108 := GenS3{F0: int64(fi91)}
                                                        _ = v108
                                                    }
                                                    _ = fi104
                                                }
                                                switch (fi91 | fi91) {
                                                case 0:
                                                    for fi109 := 0; fi109 < 14; fi109++ {
                                                        var v110 bool = (128 >= 219)
                                                        if (32767 == 857) {
                                                            fmt.Println(fi109, v6, fi91)
                                                            v2 = v6
                                                            v111 := GenS10{F0: "", F1: 697, F2: 758}
                                                            _ = v111
                                                        } else {
                                                            v2 = -86
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                        var v112 uint64 = 704
                                                        v3 += (0.1 - v3)
                                                        switch "\n\t" {
                                                        case "a":
                                                            if (128 >= int(v2)) {
                                                                var v113 bool = true
                                                                _ = v113
                                                            } else {
                                                                v1 += (v1 - (v1 ^ 42))
                                                                if len(v4) > 0 {
                                                                    v4[0] = ("" + "\n\t")
                                                                }
                                                            }
                                                        }
                                                        _ = fi109
                                                        _ = v110
                                                        _ = v112
                                                    }
                                                default:
                                                    if false {
                                                        if (-947 <= 427) {
                                                            var v114 int16 = v7
                                                            if (!((fi91 << ((0) & 63)) == (fi91 | 128))) {
                                                                if (-128 > (v2 | (v2 / ((v6) | 1)))) {
                                                                    v4 = append(v4, ("foo bar" + ""))
                                                                } else {
                                                                    v115 := GenS11{F0: (-(v7)), F1: 18446744073709551615}
                                                                    _ = v115
                                                                }
                                                                v3 -= (v3 + 3.14159)
                                                                v2 -= (-(v6))
                                                            }
                                                            for fi116 := 0; fi116 < 14; fi116++ {
                                                                var v117 int16 = (v7 & v1)
                                                                v3 += (0.1 / (float64(0.0) * v3))
                                                                _ = fi116
                                                                _ = v117
                                                            }
                                                            _ = v114
                                                        } else {
                                                            v1 = v1
                                                            v6 -= v2
                                                        }
                                                        for fi118 := 0; fi118 < 9; fi118++ {
                                                            var v119 int64 = -468
                                                            _ = fi118
                                                            _ = v119
                                                        }
                                                        var v120 int16 = v1
                                                        _ = v120
                                                    }
                                                }
                                                var v121 uint8 = 184
                                                _ = v121
                                            }
                                        }
                                    } else {
                                        v6 = v2
                                        v122 := GenS7{F0: 107}
                                        if (201 >= 970) {
                                            if (!true) {
                                                var v123 bool = ((2147483647 > 894) || (true || (!(int(v7) >= 32767))))
                                                var v124 int32 = int32(v1)
                                                _ = v123
                                                _ = v124
                                            }
                                            var v125 int = fi91
                                            _ = v125
                                        } else {
                                            for fi126 := 0; fi126 < 14; fi126++ {
                                                v4 = append(v4, (("foo bar" + "foo bar") + ("hello" + "a")))
                                                v127 := GenS2{F0: 733, F1: 65535}
                                                _ = fi126
                                                _ = v127
                                            }
                                        }
                                        _ = v122
                                    }
                                    var v128 int16 = -35
                                    var v129 int32 = -159
                                    _ = v128
                                    _ = v129
                                }
                                _ = fi91
                                _ = v92
                            }
                        }
                    }
                }
            } else {
                v130 := GenS6{F0: true, F1: 0, F2: 852}
                v3 += 0.1
                _ = v130
            }
            if (0 <= 93) {
                var v131 string = ""
                var v132 string = v131
                var v133 int64 = 1
                _ = v131
                _ = v132
                _ = v133
            }
        } else {
            v7 -= (-(v7))
        }
    }
    v134 := GenS12{F0: 0, F1: "a"}
    _ = len(v4)
    v6 = v2
    for fi135 := 0; fi135 < 14; fi135++ {
        v2 -= 100
        v136 := GenS7{F0: 32767}
        _ = fi135
        _ = v136
    }
    var v137 float64 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v137:", v137)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v134
    _ = v137
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
