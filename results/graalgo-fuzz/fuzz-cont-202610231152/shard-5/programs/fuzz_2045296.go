package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 int16; F1 int16 }
type GenS7 struct { F0 uint16; F1 int64 }
type GenS8 struct { F0 int32; F1 int16; F2 int8 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int16; F1 int }
type GenS11 struct { F0 string; F1 int32; F2 int32 }
type GenS12 struct { F0 string; F1 int64 }
type GenS13 struct { F0 int8; F1 uint64 }

func main() {
    var v1 int = 128
    var v2 float64 = 0.1
    var v3 string = ("foo bar" + ("\n\t" + ("a" + "hello")))
    for fi4 := 0; fi4 < 6; fi4++ {
        v5 := GenS1{F0: 42, F1: -35}
        var v6 float32 = 1e10
        if (526 < 638) {
            v2 = v2
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 1, F1: 1}
    v8 := []int{(-(v1)), v1, (-(v1))}
    v9 := GenS1{F0: -1, F1: 100}
    if (2 >= 255) {
        switch "hello" {
        case "\n\t":
            if (0 != 999) {
                var v10 float32 = 1e10
                _ = v10
            }
        case "":
            switch 148 {
            case 42:
                var v11 int32 = 2147483647
                _ = v11
            case 2147483647:
                v12 := GenS2{F0: 4294967295, F1: 42}
                _ = v12
            }
        }
        v2 -= float64(1e10)
    } else {
        v3 = "hello"
    }
    var v13 float64 = (((float64(3.14159) / v2) * v2) + v2)
    v14 := GenS1{F0: -9223372036854775808, F1: 100}
    var v15 float32 = (-2.25 / -2.25)
    var v16 uint32 = 2147483647
    v13 -= (((v13 - ((0.1 + v13) + float64(float32(((v13 - v2) / (v2 * v13)))))) + v13) + v2)
    _ = len(v8)
    v17 := GenS3{F0: uint(v16), F1: 32767, F2: -1}
    v3 = v3
    for fi18 := 0; fi18 < 20; fi18++ {
        if (true && false) {
            switch (v1 / ((-749) | 1)) {
            case 645:
                switch ((v3 + ((v3 + v3) + "foo bar")) + "\n\t") {
                case "hello":
                    if false {
                        if len(v8) > 0 {
                            v8[0] = 966
                        }
                        var v19 string = v3
                        if true {
                            if true {
                                v1 += (fi18 & -1)
                                var v20 int = 1
                                _ = v20
                            } else {
                                if false {
                                    v2 -= (v2 * 3.14159)
                                    fmt.Println(v16, v19, v2)
                                } else {
                                    v3 = "hello"
                                    var v21 int32 = -1
                                    var v22 uint = 0
                                    _ = v21
                                    _ = v22
                                }
                                var v23 int64 = -9223372036854775808
                                v24 := GenS4{F0: 42, F1: 1}
                                _ = v23
                                _ = v24
                            }
                            if true {
                                v16 = v16
                                for fi25 := 0; fi25 < 8; fi25++ {
                                    var v26 int = (v1 * (fi18 / (((fi25 << ((18446744073709551615) & 63))) | 1)))
                                    _ = fi25
                                    _ = v26
                                }
                            }
                        } else {
                            if (0 <= 590) {
                                for fi27 := 0; fi27 < 6; fi27++ {
                                    v19 = (v19 + (v19 + (v19 + v3)))
                                    _ = fi27
                                }
                                if false {
                                    var v28 int64 = 9223372036854775807
                                    if (((-(v28)) != 1) && (v15 < v15)) {
                                        switch v1 {
                                        case 9223372036854775807:
                                            if false {
                                                v3 = ((v19 + (v3 + v19)) + v3)
                                                v29 := GenS5{F0: 42, F1: 42}
                                                if (-128 >= 100) {
                                                    switch fi18 {
                                                    case 9223372036854775807:
                                                        fmt.Println(v19, fi18, v2)
                                                    case 1:
                                                        switch (-(v1)) {
                                                        case -611:
                                                            for fi30 := 0; fi30 < 6; fi30++ {
                                                                if false {
                                                                    var v31 int8 = 12
                                                                    _ = v31
                                                                }
                                                                _ = fi30
                                                            }
                                                        case 1:
                                                            if len(v8) >= 2 {
                                                                v8 = v8[1:len(v8)]
                                                            }
                                                        }
                                                    }
                                                    v8 = append(v8, fi18)
                                                }
                                                _ = v29
                                            } else {
                                                v28 -= (v28 | 42)
                                                if true {
                                                    var v32 int16 = -362
                                                    switch 42 {
                                                    case 42:
                                                        for fi33 := 0; fi33 < 19; fi33++ {
                                                            if (v16 >= (v16 & v16)) {
                                                                v3 = v19
                                                            }
                                                            _ = fi33
                                                        }
                                                    case 32767:
                                                        for fi34 := 0; fi34 < 17; fi34++ {
                                                            var v35 bool = ((v28 & v28) < (^(v28)))
                                                            fmt.Println(fi34, v32)
                                                            switch v3 {
                                                            case "a":
                                                                var v36 float64 = v13
                                                                _ = v36
                                                            case "":
                                                                var v37 int8 = 1
                                                                _ = v37
                                                            }
                                                            _ = fi34
                                                            _ = v35
                                                        }
                                                    default:
                                                        if (2 != 818) {
                                                            for fi38 := 0; fi38 < 14; fi38++ {
                                                                v39 := GenS6{F0: (-(v32)), F1: (v32 << ((4294967295) & 63))}
                                                                _ = fi38
                                                                _ = v39
                                                            }
                                                            v28 = v28
                                                        }
                                                    }
                                                    switch -2147483648 {
                                                    case 2147483647:
                                                        v15 = float32((v2 - v2))
                                                    case 9223372036854775807:
                                                        var v40 int64 = (-(v28))
                                                        _ = v40
                                                    default:
                                                        switch "hello" {
                                                        case "":
                                                            var v41 uint = 831
                                                            _ = v41
                                                        }
                                                    }
                                                    _ = v32
                                                }
                                                if (int32(v28) <= 1) {
                                                    v42 := GenS1{F0: v28, F1: 42}
                                                    if false {
                                                        fmt.Println(v15, v15)
                                                        fmt.Println(v15)
                                                    }
                                                    if true {
                                                        var v43 float64 = 0.1
                                                        fmt.Println(v13, v15, v43)
                                                        var v44 float32 = 0.1
                                                        _ = v43
                                                        _ = v44
                                                    } else {
                                                        switch (fi18 ^ fi18) {
                                                        case 128:
                                                            v16 = v16
                                                        case 9223372036854775807:
                                                            var v45 int8 = -128
                                                            _ = v45
                                                        }
                                                    }
                                                    _ = v42
                                                }
                                            }
                                        case 241:
                                            fmt.Println(v16, v15)
                                        case -110:
                                            v8 = append(v8, (v1 | v1))
                                        default:
                                            var v46 uint32 = v16
                                            _ = v46
                                        }
                                    } else {
                                        var v47 uint32 = v16
                                        _ = v47
                                    }
                                    fmt.Println(v19, v28, v13)
                                    _ = v28
                                }
                            } else {
                                for fi48 := 0; fi48 < 4; fi48++ {
                                    if len(v8) > 0 {
                                        v8[0] = (fi18 - (v1 - (v1 >> ((2) & 63))))
                                    }
                                    _ = fi48
                                }
                            }
                            if (0 <= -445) {
                                v16 = 2147483648
                                v13 = ((v13 / -2.25) - v13)
                                if false {
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    if false {
                                        v3 = (("foo bar" + "\n\t") + "\n\t")
                                        for fi49 := 0; fi49 < 10; fi49++ {
                                            var v50 int32 = 0
                                            switch fi49 {
                                            case -2147483648:
                                                var v51 string = ""
                                                _ = v51
                                            case 32767:
                                                v52 := GenS7{F0: 32767, F1: 623}
                                                _ = v52
                                            default:
                                                v16 = (^(v16))
                                            }
                                            fmt.Println(v15, v13)
                                            _ = fi49
                                            _ = v50
                                        }
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                    } else {
                                        v53 := GenS7{F0: 406, F1: 0}
                                        _ = v53
                                    }
                                } else {
                                    if (893 != v16) {
                                        v13 -= (0.1 - v13)
                                        if true {
                                            switch "a" {
                                            case "a":
                                                if false {
                                                    for fi54 := 0; fi54 < 12; fi54++ {
                                                        v13 += (((((v2 - ((v2 - 1.5) * 0.1)) * ((float64(float32(v2)) - v2) * v2)) + (v13 - v13)) * v13) - v13)
                                                        var v55 int8 = 1
                                                        _ = fi54
                                                        _ = v55
                                                    }
                                                } else {
                                                    v15 = v15
                                                }
                                            case "":
                                                for fi56 := 0; fi56 < 7; fi56++ {
                                                    if false {
                                                        switch 128 {
                                                        case 813:
                                                            v8 = append(v8, fi56)
                                                        case 9223372036854775807:
                                                            v13 = v2
                                                        default:
                                                            if (42 == 798) {
                                                                fmt.Println(v15, fi56)
                                                                if false {
                                                                    var v57 int64 = -9223372036854775808
                                                                    v58 := GenS2{F0: fi56, F1: -580}
                                                                    v1 = (fi18 | v1)
                                                                    _ = v57
                                                                    _ = v58
                                                                } else {
                                                                    v19 = (("\n\t" + "") + v19)
                                                                    switch ((v3 + v3) + "a") {
                                                                    case "foo bar":
                                                                        v15 = 3.14159
                                                                    }
                                                                }
                                                                v59 := GenS5{F0: 2147483647, F1: 1}
                                                                v16 = 743
                                                                if len(v8) > 0 {
                                                                    v8[0] = fi56
                                                                }
                                                                v13 = (v13 * (v2 - v2))
                                                                switch (v1 & v1) {
                                                                case 2147483647:
                                                                    if (42 <= 128) {
                                                                        v60 := GenS8{F0: 1, F1: -1, F2: 8}
                                                                        v2 = v2
                                                                        _ = v60
                                                                    }
                                                                case 42:
                                                                    var v61 float64 = v13
                                                                    _ = v61
                                                                }
                                                                _ = v59
                                                            } else {
                                                                if (!(!(0 == 0))) {
                                                                    if false {
                                                                        fmt.Println(v3, v1, v16)
                                                                        v3 = "hello"
                                                                    } else {
                                                                        if len(v8) >= 2 {
                                                                            v8 = v8[1:len(v8)]
                                                                        }
                                                                        if (true && true) {
                                                                            var v62 uint8 = 59
                                                                            _ = v62
                                                                        } else {
                                                                            if false {
                                                                                fmt.Println(v15)
                                                                            } else {
                                                                                var v63 uint8 = 1
                                                                                var v64 uint16 = 42
                                                                                _ = v63
                                                                                _ = v64
                                                                            }
                                                                            v2 += v2
                                                                            v65 := GenS3{F0: 863, F1: 23, F2: 0}
                                                                            _ = v65
                                                                        }
                                                                    }
                                                                    v66 := GenS2{F0: (fi18 | v1), F1: -912}
                                                                    v16 = (^(v16))
                                                                    _ = v66
                                                                } else {
                                                                    v19 = v3
                                                                    if ((18446744073709551615 < 2) && (!false)) {
                                                                        v2 = 3.14159
                                                                        if len(v8) >= 2 {
                                                                            v8 = v8[1:len(v8)]
                                                                        }
                                                                        switch 2147483647 {
                                                                        case 1:
                                                                            fmt.Println(fi56, v2, v3)
                                                                            if (!false) {
                                                                                switch fi18 {
                                                                                case 9223372036854775807:
                                                                                    if len(v8) >= 2 {
                                                                                        v8 = v8[1:len(v8)]
                                                                                    }
                                                                                }
                                                                                var v67 uint8 = 127
                                                                                _ = v67
                                                                            }
                                                                            var v68 int16 = 1000
                                                                            _ = v68
                                                                        case 2147483647:
                                                                            v3 = v19
                                                                        case 42:
                                                                            v69 := GenS9{F0: 1}
                                                                            _ = v69
                                                                        default:
                                                                            var v70 int8 = -1
                                                                            _ = v70
                                                                        }
                                                                    } else {
                                                                        if len(v8) >= 2 {
                                                                            v8 = v8[1:len(v8)]
                                                                        }
                                                                        v19 = v3
                                                                        switch ("foo bar" + ((((v3 + v3) + "\n\t") + v3) + ("hello" + "\n\t"))) {
                                                                        case "a":
                                                                            v71 := GenS5{F0: 1, F1: 42}
                                                                            _ = v71
                                                                        case "":
                                                                            switch -2147483648 {
                                                                            case 0:
                                                                                v72 := GenS4{F0: 2, F1: 42}
                                                                                v73 := GenS5{F0: -986, F1: 1}
                                                                                _ = v72
                                                                                _ = v73
                                                                            }
                                                                        default:
                                                                            v74 := GenS6{F0: 0, F1: -32768}
                                                                            _ = v74
                                                                        }
                                                                    }
                                                                    if (-2147483648 == -1) {
                                                                        var v75 uint8 = 26
                                                                        if (!(!true)) {
                                                                            if len(v8) > 0 {
                                                                                v8[0] = 9223372036854775807
                                                                            }
                                                                            if (!false) {
                                                                                v75 = v75
                                                                            }
                                                                        } else {
                                                                            if len(v8) >= 2 {
                                                                                v8 = v8[1:len(v8)]
                                                                            }
                                                                            v3 = (v19 + ((v3 + "a") + v19))
                                                                            v76 := GenS7{F0: 596, F1: 699}
                                                                            var v77 uint64 = 767
                                                                            _ = v76
                                                                            _ = v77
                                                                        }
                                                                        _ = v75
                                                                    } else {
                                                                        v78 := GenS4{F0: 18446744073709551615, F1: -2147483648}
                                                                        _ = v78
                                                                    }
                                                                    if len(v8) >= 2 {
                                                                        v8 = v8[1:len(v8)]
                                                                    }
                                                                    var v79 float64 = v2
                                                                    _ = v79
                                                                }
                                                                v13 = 0.1
                                                                switch (v1 | -1) {
                                                                case -659:
                                                                    var v80 bool = ((v16 << ((65535) & 63)) != (v16 * v16))
                                                                    _ = v80
                                                                case -259:
                                                                    switch 32767 {
                                                                    case 811:
                                                                        v3 = v3
                                                                    case 42:
                                                                        if (float64((v15 + ((v15 + ((-2.25 * v15) + v15)) / v15))) > 1e10) {
                                                                            var v81 uint64 = 42
                                                                            if false {
                                                                                v13 -= (v2 / 1.0)
                                                                                fmt.Println(v3)
                                                                                v82 := GenS10{F0: -32768, F1: fi56}
                                                                                _ = v82
                                                                            } else {
                                                                                if false {
                                                                                    var v83 string = ("hello" + v3)
                                                                                    v84 := GenS11{F0: v83, F1: int32(v81), F2: -322}
                                                                                    _ = v83
                                                                                    _ = v84
                                                                                }
                                                                            }
                                                                            _ = v81
                                                                        }
                                                                    case -1:
                                                                        v15 += v15
                                                                    default:
                                                                        var v85 float64 = (((v13 - (0.0 - (0.1 / ((v13 * v2) + 0.1)))) + v2) - ((3.14159 * 0.1) * v2))
                                                                        _ = v85
                                                                    }
                                                                case 32768:
                                                                    v86 := GenS4{F0: 42, F1: -1}
                                                                    _ = v86
                                                                }
                                                                var v87 uint32 = 1
                                                                _ = v87
                                                            }
                                                            v88 := GenS4{F0: 134, F1: -2147483648}
                                                            _ = v88
                                                        }
                                                    } else {
                                                        v19 = v19
                                                    }
                                                    var v89 uint16 = 32767
                                                    var v90 int = -345
                                                    v15 = (((v15 - v15) / float32(v13)) - (((0.1 * (0.0 - (0.0 - (v15 * v15)))) * v15) * v15))
                                                    _ = fi56
                                                    _ = v89
                                                    _ = v90
                                                }
                                            }
                                            switch fi18 {
                                            case 42:
                                                var v91 bool = false
                                                _ = v91
                                            default:
                                                switch v1 {
                                                case -404:
                                                    switch -151 {
                                                    case 575:
                                                        v3 = "foo bar"
                                                    default:
                                                        switch v1 {
                                                        case 32768:
                                                            fmt.Println(v1)
                                                        case -9223372036854775808:
                                                            var v92 uint64 = 315
                                                            _ = v92
                                                        }
                                                    }
                                                case 0:
                                                    if len(v8) > 0 {
                                                        v8[0] = 1
                                                    }
                                                }
                                            }
                                        }
                                        if false {
                                            v1 = v1
                                            if len(v8) > 0 {
                                                v8[0] = fi18
                                            }
                                        }
                                    }
                                    if true {
                                        if (false || (32767 > 312)) {
                                            for fi93 := 0; fi93 < 12; fi93++ {
                                                if (v2 >= (float64((-2.25 / v15)) - 3.14159)) {
                                                    v94 := GenS10{F0: -1, F1: (fi93 ^ fi18)}
                                                    if len(v8) > 0 {
                                                        v8[0] = fi93
                                                    }
                                                    var v95 uint16 = 48
                                                    _ = v94
                                                    _ = v95
                                                }
                                                if (-760 == 9223372036854775807) {
                                                    v13 += -2.25
                                                    switch v19 {
                                                    case "hello":
                                                        v96 := GenS1{F0: -718, F1: -1}
                                                        _ = v96
                                                    }
                                                }
                                                var v97 uint8 = uint8(v16)
                                                _ = fi93
                                                _ = v97
                                            }
                                            fmt.Println(v16, v13)
                                        } else {
                                            var v98 int8 = -35
                                            _ = v98
                                        }
                                    } else {
                                        if (!(0 <= 0)) {
                                            var v99 int64 = 42
                                            if true {
                                                switch v3 {
                                                case "\n\t":
                                                    v100 := GenS5{F0: -1, F1: -47}
                                                    _ = v100
                                                case "hello":
                                                    v101 := GenS3{F0: 97, F1: 1, F2: 121}
                                                    _ = v101
                                                case "foo bar":
                                                    v19 = v3
                                                }
                                                for fi102 := 0; fi102 < 1; fi102++ {
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                    _ = fi102
                                                }
                                            }
                                            _ = v99
                                        }
                                    }
                                }
                            } else {
                                v16 = v16
                                for fi103 := 0; fi103 < 20; fi103++ {
                                    v104 := GenS9{F0: 621}
                                    _ = fi103
                                    _ = v104
                                }
                                switch ((v19 + v19) + v19) {
                                case "\n\t":
                                    fmt.Println(v16)
                                case "foo bar":
                                    for fi105 := 0; fi105 < 13; fi105++ {
                                        if (2147483648 < 42) {
                                            if true {
                                                if true {
                                                    if true {
                                                        v106 := GenS9{F0: 1}
                                                        if (true || (v1 <= fi18)) {
                                                            fmt.Println(fi105, v19, fi18)
                                                            var v107 uint16 = 32767
                                                            switch "" {
                                                            case "foo bar":
                                                                var v108 uint8 = 217
                                                                _ = v108
                                                            case "hello":
                                                                switch (v19 + (v19 + v3)) {
                                                                case "foo bar":
                                                                    fmt.Println(fi18, v107)
                                                                case "":
                                                                    fmt.Println(v13)
                                                                case "\n\t":
                                                                    if (!true) {
                                                                        if false {
                                                                            if (!true) {
                                                                                v109 := GenS6{F0: 1, F1: -32768}
                                                                                if true {
                                                                                    var v110 bool = false
                                                                                    _ = v110
                                                                                } else {
                                                                                    var v111 int32 = -1
                                                                                    v15 -= ((3.14159 / 1.0) * v15)
                                                                                    _ = v111
                                                                                }
                                                                                _ = v109
                                                                            } else {
                                                                                v112 := GenS4{F0: 1, F1: 2147483647}
                                                                                _ = v112
                                                                            }
                                                                            v113 := GenS8{F0: -648, F1: 1000, F2: -87}
                                                                            switch v19 {
                                                                            case "foo bar":
                                                                                var v114 float32 = float32(float64(0.0))
                                                                                _ = v114
                                                                            default:
                                                                                if (-676 <= -9223372036854775808) {
                                                                                    v16 = (^(v16))
                                                                                    if (true && true) {
                                                                                        if true {
                                                                                            fmt.Println(fi105)
                                                                                        }
                                                                                    } else {
                                                                                        v3 = v19
                                                                                    }
                                                                                    v115 := GenS12{F0: (v19 + "\n\t"), F1: 1}
                                                                                    _ = v115
                                                                                }
                                                                            }
                                                                            _ = v113
                                                                        }
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                fmt.Println(v107, fi105)
                                                            }
                                                            _ = v107
                                                        } else {
                                                            v116 := GenS3{F0: 42, F1: 0, F2: -1}
                                                            var v117 bool = true
                                                            _ = v116
                                                            _ = v117
                                                        }
                                                        _ = v106
                                                    }
                                                    var v118 bool = (-1 != -599)
                                                    v8 = append(v8, fi105)
                                                    _ = v118
                                                } else {
                                                    var v119 int = v1
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                    v1 = fi105
                                                    fmt.Println(v2, v19, v16)
                                                    _ = v119
                                                }
                                                _ = len(v8)
                                                v8 = append(v8, 32767)
                                                if (!(false && ((!(4294967295 != 0)) || (false || (false || (4294967295 != uint(v16))))))) {
                                                    fmt.Println(v2)
                                                }
                                            } else {
                                                fmt.Println(v15, v19)
                                            }
                                            v120 := GenS2{F0: v1, F1: -778}
                                            var v121 float64 = float64(0.1)
                                            if false {
                                                if len(v8) > 0 {
                                                    v8[0] = 829
                                                }
                                                switch fi18 {
                                                case 1:
                                                    v13 -= 0.0
                                                default:
                                                    switch (fi105 - fi105) {
                                                    case -786:
                                                        v16 = 2147483647
                                                    case 346:
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                    }
                                                }
                                            }
                                            var v122 int32 = 757
                                            _ = v120
                                            _ = v121
                                            _ = v122
                                        } else {
                                            var v123 uint8 = 234
                                            v19 = v19
                                            _ = v123
                                        }
                                        v124 := GenS2{F0: (fi18 - (fi105 & (fi18 * 9223372036854775807))), F1: -795}
                                        _ = fi105
                                        _ = v124
                                    }
                                }
                            }
                        }
                        _ = v19
                    } else {
                        for fi125 := 0; fi125 < 6; fi125++ {
                            if false {
                                v2 -= (float64((((0.1 / (v15 / -2.25)) - (v15 / v15)) + (v15 * (1e10 - 0.1)))) - v13)
                                var v126 uint64 = 108
                                _ = v126
                            }
                            _ = len(v8)
                            switch (v1 + (fi125 / ((fi125) | 1))) {
                            case 1:
                                var v127 float64 = (float64(v15) - v13)
                                var v128 int32 = 2147483647
                                fmt.Println(v16)
                                _ = v127
                                _ = v128
                            }
                            _ = fi125
                        }
                        v3 = ((v3 + "hello") + "hello")
                    }
                case "":
                    v129 := GenS2{F0: fi18, F1: 42}
                    _ = v129
                default:
                    v15 = v15
                }
            case 897:
                fmt.Println(v16, v1, v1)
            default:
                if ((false || false) || (1 == 1)) {
                    if (!(184 > -254)) {
                        fmt.Println(v16)
                    } else {
                        v13 -= v2
                        if (-107 == 0) {
                            v8 = append(v8, v1)
                            if len(v8) > 0 {
                                v8[0] = (fi18 >> ((2) & 63))
                            }
                        } else {
                            if true {
                                if (18446744073709551615 != 800) {
                                    if true {
                                        if true {
                                            fmt.Println(v13)
                                        }
                                        _ = len(v8)
                                        if true {
                                            var v130 bool = (644 >= 503)
                                            _ = v130
                                        } else {
                                            var v131 int32 = 240
                                            v3 = v3
                                            _ = v131
                                        }
                                    } else {
                                        var v132 uint32 = v16
                                        switch (v1 | v1) {
                                        case 32768:
                                            var v133 uint32 = v132
                                            _ = v133
                                        }
                                        _ = v132
                                    }
                                } else {
                                    v1 -= 127
                                }
                            } else {
                                var v134 float64 = v13
                                var v135 uint32 = 255
                                fmt.Println(v3, v16)
                                _ = v134
                                _ = v135
                            }
                            fmt.Println(v3, v2)
                            if (127 <= 1) {
                                v1 += -9223372036854775808
                                v8 = append(v8, fi18)
                            } else {
                                v136 := GenS13{F0: -1, F1: 18446744073709551615}
                                v15 = (v15 + 0.1)
                                _ = v136
                            }
                        }
                    }
                    fmt.Println(v16, v15, fi18)
                }
            }
            _ = len(v8)
        } else {
            fmt.Println(v13, v15, v15)
        }
        fmt.Println(v15)
        var v137 bool = true
        _ = fi18
        _ = v137
    }
    v15 = (v15 / (((v15 - ((0.0 + ((v15 * v15) / 1e10)) * v15)) / v15) + (v15 / ((v15 + v15) - v15))))
    for fi138 := 0; fi138 < 6; fi138++ {
        fmt.Println(v1, v15, fi138)
        switch 127 {
        case -96:
            var v139 int = (fi138 | v1)
            _ = v139
        case 312:
            fmt.Println(v16, v3)
        case 42:
            for fi140 := 0; fi140 < 5; fi140++ {
                v13 = v2
                v15 -= (v15 / v15)
                _ = fi140
            }
        default:
            switch fi138 {
            case 4294967295:
                fmt.Println(v15)
            case 155:
                if false {
                    for fi141 := 0; fi141 < 2; fi141++ {
                        v3 = (((v3 + v3) + v3) + "\n\t")
                        _ = fi141
                    }
                    _ = len(v8)
                }
            case -9223372036854775808:
                _ = len(v8)
            }
        }
        _ = fi138
    }
    v8 = append(v8, 0)
    v15 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
