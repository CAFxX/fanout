package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int8; F1 int32; F2 uint }
type GenS5 struct { F0 int16; F1 int64; F2 uint }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 uint8; F1 bool }
type GenS8 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS9 struct { F0 bool; F1 uint16 }
type GenS10 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS11 struct { F0 int16; F1 uint8 }
type GenS12 struct { F0 int; F1 bool; F2 int8 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 bool; F1 int; F2 string }
type GenS15 struct { F0 uint8; F1 string; F2 int32 }
type GenS16 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 uint8; F1 int64 }
type GenS20 struct { F0 string }

func main() {
    var v1 int8 = -128
    var v2 uint16 = 65535
    v3 := []uint32{2, 2, 99}
    var v4 uint32 = 42
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (4294967295 != 1) {
        for fi5 := 0; fi5 < 19; fi5++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi5
        }
        v6 := GenS1{F0: -874}
        _ = v6
    } else {
        fmt.Println(v2, v4)
        if (!false) {
            var v7 string = ("\n\t" + ("a" + (("hello" + "foo bar") + "")))
            for fi8 := 0; fi8 < 14; fi8++ {
                for fi9 := 0; fi9 < 17; fi9++ {
                    switch "" {
                    case "a":
                        if true {
                            v10 := GenS2{F0: 124, F1: 128, F2: 0}
                            _ = v10
                        }
                    case "":
                        switch v7 {
                        case "foo bar":
                            v4 = v4
                        default:
                            v3 = append(v3, 4294967295)
                        }
                    default:
                        var v11 int = (fi8 % ((fi8) | 1))
                        _ = v11
                    }
                    _ = fi9
                }
                v4 += 2
                _ = fi8
            }
            _ = v7
        } else {
            v2 = v2
            var v12 float32 = 1.5
            _ = v12
        }
        var v13 uint32 = uint32(v2)
        _ = v13
    }
    _ = len(v3)
    for fi14 := 0; fi14 < 4; fi14++ {
        v2 = 409
        switch "\n\t" {
        case "\n\t":
            var v15 float32 = (0.0 + 3.14159)
            _ = v15
        case "hello":
            v1 += 121
        default:
            var v16 bool = true
            _ = v16
        }
        var v17 uint16 = v2
        _ = fi14
        _ = v17
    }
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v2, v4, v2)
    case "a":
        if (2 > 1) {
            if (!false) {
                v4 += (v4 / ((65535) | 1))
            }
        } else {
            v4 = (v4 - (v4 << ((42) & 63)))
            if false {
                if true {
                    fmt.Println(v1, v4, v2)
                } else {
                    var v18 bool = (255 < 127)
                    for fi19 := 0; fi19 < 13; fi19++ {
                        if v18 {
                            var v20 bool = v18
                            v3 = append(v3, 2147483648)
                            if len(v3) > 0 {
                                v3[0] = v4
                            }
                            _ = v20
                        } else {
                            if (fi19 == (fi19 >> ((366) & 63))) {
                                v4 += uint32(v2)
                                v2 += v2
                            }
                            if (((v18 || v18) && (65535 >= (v2 * (v2 & (v2 | 146))))) && v18) {
                                _ = len(v3)
                            }
                        }
                        fmt.Println(v2, v4, v1)
                        _ = fi19
                    }
                    _ = v18
                }
                switch (("foo bar" + ("" + "foo bar")) + "") {
                case "":
                    for fi21 := 0; fi21 < 18; fi21++ {
                        v2 -= (^(v2))
                        _ = fi21
                    }
                default:
                    for fi22 := 0; fi22 < 1; fi22++ {
                        v3 = append(v3, (v4 << ((255) & 63)))
                        _ = fi22
                    }
                }
            }
        }
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if true {
        if true {
            var v23 int64 = 0
            for fi24 := 0; fi24 < 18; fi24++ {
                v23 -= (v23 % ((42) | 1))
                if (!false) {
                    var v25 uint16 = 2
                    if false {
                        fmt.Println(fi24)
                        fmt.Println(fi24, v25)
                    } else {
                        if false {
                            v25 = v25
                            switch ("foo bar" + "\n\t") {
                            case "\n\t":
                                var v26 int8 = (v1 ^ v1)
                                _ = v26
                            case "foo bar":
                                if true {
                                    v27 := GenS3{F0: fi24, F1: 42}
                                    v1 -= 0
                                    _ = v27
                                }
                            case "":
                                var v28 int64 = v23
                                _ = v28
                            default:
                                v2 += v25
                            }
                        } else {
                            for fi29 := 0; fi29 < 17; fi29++ {
                                v30 := GenS4{F0: (v1 | (v1 + 127)), F1: 42, F2: 825}
                                if false {
                                    var v31 float32 = 0.1
                                    if (-32768 == -1) {
                                        v1 = (v1 & v1)
                                        if ((false || ((fi29 - fi29) < fi29)) || true) {
                                            switch "foo bar" {
                                            case "foo bar":
                                                fmt.Println(v4, v2)
                                            default:
                                                switch int(v25) {
                                                case 9223372036854775807:
                                                    fmt.Println(v31)
                                                }
                                                switch (fi29 | fi29) {
                                                case 128:
                                                    v32 := GenS2{F0: 65535, F1: 1, F2: 0}
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        var v33 uint16 = v25
                                                        _ = v33
                                                    case "":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    case "\n\t":
                                                        var v34 uint32 = v4
                                                        _ = v34
                                                    }
                                                    var v35 uint8 = 180
                                                    v36 := GenS5{F0: 0, F1: (v23 & v23), F2: 18446744073709551615}
                                                    _ = v32
                                                    _ = v35
                                                    _ = v36
                                                }
                                            }
                                        }
                                        if false {
                                            if (!true) {
                                                v37 := GenS3{F0: (fi24 % (((fi24 << ((1) & 63))) | 1)), F1: 572}
                                                var v38 int16 = -837
                                                v39 := GenS4{F0: v1, F1: int32(fi29), F2: 1}
                                                _ = v37
                                                _ = v38
                                                _ = v39
                                            } else {
                                                v40 := GenS5{F0: 1000, F1: (v23 << ((42) & 63)), F2: 4294967295}
                                                _ = v40
                                            }
                                        }
                                    } else {
                                        if (84 != uint8(v2)) {
                                            if ((v23 - v23) != v23) {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            } else {
                                                v4 += v4
                                                if false {
                                                    if (true || true) {
                                                        v41 := GenS6{F0: 351, F1: 32767}
                                                        var v42 int8 = (-(v1))
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                    var v43 int = fi24
                                                    v44 := GenS7{F0: 117, F1: (4294967295 != 1)}
                                                    _ = v43
                                                    _ = v44
                                                } else {
                                                    fmt.Println(fi29)
                                                }
                                                fmt.Println(v2, v31, fi29)
                                                v2 = (v25 * v25)
                                            }
                                            if ((true || ((v23 ^ v23) < (-(v23)))) && (!false)) {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                v31 = (v31 / (((v31 + float32(1.5)) / v31) + (1e10 - v31)))
                                                if len(v3) > 0 {
                                                    v3[0] = (^(v4))
                                                }
                                                switch (fi24 / ((fi24) | 1)) {
                                                case 32768:
                                                    v3 = append(v3, v4)
                                                case 4294967295:
                                                    var v45 int = fi24
                                                    _ = v45
                                                }
                                            }
                                            fmt.Println(fi24, fi24, v23)
                                            v46 := GenS3{F0: 0, F1: 1}
                                            _ = v46
                                        } else {
                                            v47 := GenS8{F0: 42, F1: ((4294967295 >= 931) && (!true)), F2: v4}
                                            _ = v47
                                        }
                                    }
                                    var v48 float32 = v31
                                    _ = v31
                                    _ = v48
                                } else {
                                    var v49 uint32 = v4
                                    var v50 int64 = v23
                                    var v51 uint16 = v25
                                    _ = v49
                                    _ = v50
                                    _ = v51
                                }
                                v4 -= (^(v4))
                                _ = fi29
                                _ = v30
                            }
                            if (false && true) {
                                var v52 float32 = 3.14159
                                _ = v52
                            }
                        }
                    }
                    _ = v25
                } else {
                    v2 += (v2 ^ v2)
                    switch (fi24 + (fi24 >> ((932) & 63))) {
                    case 1:
                        switch ("hello" + "") {
                        case "foo bar":
                            var v53 uint64 = 2
                            _ = v53
                        }
                    case 42:
                        if ((fi24 + (fi24 - fi24)) >= (fi24 / (((fi24 * (fi24 / ((fi24) | 1)))) | 1))) {
                            fmt.Println(v1)
                            for fi54 := 0; fi54 < 1; fi54++ {
                                v3 = append(v3, (v4 % (((v4 >> ((0) & 63))) | 1)))
                                switch "" {
                                case "\n\t":
                                    var v55 string = "foo bar"
                                    _ = v55
                                case "hello":
                                    if len(v3) > 0 {
                                        v3[0] = 82
                                    }
                                default:
                                    var v56 int16 = 0
                                    _ = v56
                                }
                                var v57 float32 = 0.1
                                _ = fi54
                                _ = v57
                            }
                            var v58 int = (fi24 / ((fi24) | 1))
                            _ = v58
                        }
                    }
                    _ = len(v3)
                }
                v4 = 267
                _ = fi24
            }
            for fi59 := 0; fi59 < 13; fi59++ {
                if ((v1 * (v1 & (v1 * v1))) != (v1 / ((103) | 1))) {
                    for fi60 := 0; fi60 < 3; fi60++ {
                        fmt.Println(v4)
                        v1 = (v1 % ((0) | 1))
                        var v61 string = ""
                        _ = fi60
                        _ = v61
                    }
                    fmt.Println(v1, v1, fi59)
                }
                for fi62 := 0; fi62 < 19; fi62++ {
                    v23 += (v23 * (v23 | v23))
                    _ = fi62
                }
                _ = fi59
            }
            _ = v23
        } else {
            v3 = append(v3, v4)
            for fi63 := 0; fi63 < 13; fi63++ {
                switch "hello" {
                case "\n\t":
                    if (!(true || ((!true) || ((true || (42 <= 1)) && ((v1 - v1) > v1))))) {
                        if false {
                            for fi64 := 0; fi64 < 5; fi64++ {
                                var v65 int64 = 9223372036854775807
                                if ((true || (("foo bar" >= (("a" + "\n\t") + "")) && (false || false))) && ("hello" == "")) {
                                    v4 = v4
                                    switch fi64 {
                                    case 32767:
                                        fmt.Println(fi64, v1, v2)
                                    case -426:
                                        if false {
                                            var v66 bool = (4294967295 != 846)
                                            fmt.Println(fi64, v4, v2)
                                            var v67 bool = (!((1 <= 769) || (!v66)))
                                            _ = v66
                                            _ = v67
                                        } else {
                                            var v68 int64 = (v65 % ((v65) | 1))
                                            _ = v68
                                        }
                                        v65 = (-(v65))
                                        if len(v3) > 0 {
                                            v3[0] = 2147483648
                                        }
                                        if true {
                                            var v69 int8 = v1
                                            _ = v69
                                        } else {
                                            v1 = (v1 << ((587) & 63))
                                        }
                                        switch 0 {
                                        case 127:
                                            if true {
                                                v4 = (v4 - (v4 << ((42) & 63)))
                                                fmt.Println(v4, v1)
                                                v70 := GenS8{F0: 27, F1: false, F2: 2147483647}
                                                _ = v70
                                            }
                                        case 0:
                                            if false {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    var v71 uint64 = 18446744073709551615
                                                    _ = v71
                                                case "":
                                                    if false {
                                                        fmt.Println(fi64, v65)
                                                        if (true || false) {
                                                            if (!false) {
                                                                _ = len(v3)
                                                            } else {
                                                                if len(v3) > 0 {
                                                                    v3[0] = (v4 % ((v4) | 1))
                                                                }
                                                                if len(v3) > 0 {
                                                                    v3[0] = (v4 & (^(v4)))
                                                                }
                                                                var v72 uint16 = (v2 - 435)
                                                                _ = v72
                                                            }
                                                            v73 := GenS3{F0: (fi63 | (^(fi64))), F1: 18446744073709551615}
                                                            var v74 int = int(v65)
                                                            switch (fi63 & fi64) {
                                                            case -9223372036854775808:
                                                                fmt.Println(fi64, fi63, fi64)
                                                            case -1:
                                                                switch 42 {
                                                                case 0:
                                                                    v75 := GenS9{F0: false, F1: v2}
                                                                    _ = v75
                                                                }
                                                            case 868:
                                                                _ = len(v3)
                                                            default:
                                                                v76 := GenS9{F0: (false && false), F1: (v2 ^ v2)}
                                                                _ = v76
                                                            }
                                                            v77 := GenS10{F0: v2, F1: 2, F2: v2}
                                                            _ = v73
                                                            _ = v74
                                                            _ = v77
                                                        } else {
                                                            if ((1 > 106) || ((v1 % ((v1) | 1)) <= v1)) {
                                                                v2 += 2
                                                            }
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            var v78 uint16 = 0
                                                            v79 := GenS2{F0: 32767, F1: 151, F2: 2147483647}
                                                            fmt.Println(v65)
                                                            _ = v78
                                                            _ = v79
                                                        }
                                                    } else {
                                                        var v80 uint8 = 46
                                                        var v81 float64 = 3.14159
                                                        _ = v80
                                                        _ = v81
                                                    }
                                                case "hello":
                                                    v82 := GenS10{F0: 255, F1: 18446744073709551615, F2: 29}
                                                    _ = v82
                                                }
                                                fmt.Println(v65, fi63, v1)
                                                if true {
                                                    if true {
                                                        if ((v2 % ((2) | 1)) > v2) {
                                                            var v83 int32 = 654
                                                            if false {
                                                                v65 = v65
                                                                v4 = v4
                                                                var v84 uint16 = v2
                                                                _ = v84
                                                            }
                                                            var v85 int64 = v65
                                                            v1 = v1
                                                            v86 := GenS11{F0: 1, F1: 0}
                                                            _ = v83
                                                            _ = v85
                                                            _ = v86
                                                        }
                                                        v2 = (^(v2))
                                                    }
                                                    var v87 int16 = 0
                                                    fmt.Println(v87, v87, v4)
                                                    v87 += v87
                                                    var v88 uint8 = 107
                                                    _ = v87
                                                    _ = v88
                                                } else {
                                                    var v89 bool = true
                                                    v90 := GenS12{F0: (fi64 + (fi64 ^ -383)), F1: v89, F2: int8(v2)}
                                                    _ = v89
                                                    _ = v90
                                                }
                                            } else {
                                                switch ("\n\t" + "foo bar") {
                                                case "hello":
                                                    var v91 uint8 = 20
                                                    _ = v91
                                                default:
                                                    if (!false) {
                                                        v2 -= (v2 ^ 255)
                                                    } else {
                                                        v92 := GenS3{F0: (fi63 * (fi63 | (-(fi64)))), F1: 721}
                                                        v2 = v2
                                                        if (!(!true)) {
                                                            switch fi63 {
                                                            case 2147483647:
                                                                fmt.Println(fi63, v65, v2)
                                                                _ = len(v3)
                                                                var v93 float32 = (((1e10 * 3.14159) * -2.25) * -2.25)
                                                                fmt.Println(v2)
                                                                var v94 float32 = v93
                                                                _ = v93
                                                                _ = v94
                                                            }
                                                            var v95 uint32 = (v4 & v4)
                                                            if ((v95 % (((v4 >> ((1) & 63))) | 1)) >= v4) {
                                                                if true {
                                                                    v65 -= (^(v65))
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    v4 += 532
                                                                    v2 -= (v2 * (v2 & (v2 % ((0) | 1))))
                                                                } else {
                                                                    v1 += v1
                                                                    var v96 uint = 0
                                                                    _ = v96
                                                                }
                                                            }
                                                            _ = v95
                                                        }
                                                        _ = v92
                                                    }
                                                }
                                                v2 = (^(v2))
                                            }
                                        case 389:
                                            v97 := GenS13{F0: -265}
                                            var v98 int = fi63
                                            _ = v97
                                            _ = v98
                                        }
                                    case 2147483647:
                                        v99 := GenS14{F0: false, F1: fi63, F2: "foo bar"}
                                        _ = v99
                                    }
                                }
                                if ((-2.25 / ((1.5 * (((((0.0 / 1.0) + -2.25) + (float32(3.14159) + 0.1)) - 0.1) * 1e10)) * (-2.25 * 3.14159))) > ((0.0 + 0.1) * 0.0)) {
                                    if false {
                                        var v100 uint32 = v4
                                        _ = v100
                                    } else {
                                        v2 -= (v2 + (v2 ^ (v2 ^ (v2 / ((v2) | 1)))))
                                        v65 = (v65 ^ (v65 / ((v65) | 1)))
                                    }
                                }
                                if len(v3) > 0 {
                                    v3[0] = v4
                                }
                                v4 = (v4 * 0)
                                _ = fi64
                                _ = v65
                            }
                        } else {
                            fmt.Println(v1)
                        }
                    }
                }
                _ = fi63
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        switch "a" {
        case "":
            v2 = 32767
        case "hello":
            var v101 string = ("a" + "a")
            _ = v101
        default:
            fmt.Println(v1, v4)
        }
    }
    if true {
        if true {
            var v102 float32 = 1.5
            _ = v102
        } else {
            for fi103 := 0; fi103 < 4; fi103++ {
                fmt.Println(v2)
                switch (fi103 + fi103) {
                case 127:
                    if len(v3) > 0 {
                        v3[0] = uint32(v2)
                    }
                case 922:
                    v104 := GenS1{F0: -626}
                    _ = v104
                case 42:
                    if false {
                        for fi105 := 0; fi105 < 18; fi105++ {
                            v106 := GenS4{F0: v1, F1: 1, F2: 18446744073709551615}
                            var v107 uint32 = (v4 ^ v4)
                            v108 := GenS15{F0: 128, F1: "", F2: -436}
                            _ = fi105
                            _ = v106
                            _ = v107
                            _ = v108
                        }
                    } else {
                        for fi109 := 0; fi109 < 5; fi109++ {
                            v2 -= (v2 + 42)
                            fmt.Println(v1, v4)
                            _ = fi109
                        }
                        var v110 bool = ((!(!false)) || (true && true))
                        for fi111 := 0; fi111 < 2; fi111++ {
                            if (855 > 4294967295) {
                                fmt.Println(v1, v1, fi111)
                            }
                            v112 := GenS2{F0: v2, F1: 1, F2: -713}
                            var v113 uint64 = 79
                            _ = fi111
                            _ = v112
                            _ = v113
                        }
                        _ = v110
                    }
                default:
                    for fi114 := 0; fi114 < 13; fi114++ {
                        v115 := GenS16{F0: 283, F1: (!false), F2: 1}
                        if false {
                            switch "hello" {
                            case "a":
                                var v116 bool = (false && true)
                                if (v116 || ((v4 >> ((65535) & 63)) == (^(v4)))) {
                                    var v117 bool = (false || v116)
                                    switch fi103 {
                                    case -388:
                                        switch ("\n\t" + "hello") {
                                        case "hello":
                                            var v118 float32 = 3.14159
                                            _ = v118
                                        case "":
                                            if ((352 <= 332) || v117) {
                                                var v119 float64 = float64(-2.25)
                                                fmt.Println(v1, v119)
                                                _ = v119
                                            } else {
                                                var v120 uint = 65535
                                                if false {
                                                    v120 += v120
                                                    var v121 uint32 = (v4 & (v4 >> ((18446744073709551615) & 63)))
                                                    if true {
                                                        switch "hello" {
                                                        case "a":
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        case "":
                                                            var v122 float64 = 3.14159
                                                            _ = v122
                                                        default:
                                                            v123 := GenS16{F0: 983, F1: true, F2: 1}
                                                            _ = v123
                                                        }
                                                        switch ("" + "") {
                                                        case "":
                                                            v124 := GenS10{F0: uint16(v4), F1: (v120 & (v120 & (v120 >> (((v120 & v120)) & 63)))), F2: v2}
                                                            _ = v124
                                                        }
                                                        if (0 < 797) {
                                                            var v125 int32 = -2147483648
                                                            var v126 bool = v116
                                                            _ = v125
                                                            _ = v126
                                                        }
                                                    } else {
                                                        if ((!((984 < v2) && v116)) && (32767 <= 740)) {
                                                            var v127 int16 = 1000
                                                            if (1 >= 2147483648) {
                                                                v128 := GenS17{F0: false}
                                                                _ = v128
                                                            }
                                                            if len(v3) > 0 {
                                                                v3[0] = v121
                                                            }
                                                            _ = v127
                                                        }
                                                        v129 := GenS7{F0: 2, F1: (!(!(990 <= (v4 << ((65535) & 63)))))}
                                                        switch (fi114 >> ((v120) & 63)) {
                                                        case 687:
                                                            if true {
                                                                var v130 int64 = -718
                                                                _ = v130
                                                            } else {
                                                                if (!(255 < 0)) {
                                                                    switch -713 {
                                                                    case 2147483647:
                                                                        var v131 int64 = -558
                                                                        _ = v131
                                                                    case 4294967295:
                                                                        var v132 int16 = 0
                                                                        _ = v132
                                                                    }
                                                                } else {
                                                                    v2 = (v2 | v2)
                                                                    v133 := GenS18{F0: -729}
                                                                    _ = v133
                                                                }
                                                                var v134 uint16 = 65535
                                                                var v135 bool = (!v117)
                                                                _ = v134
                                                                _ = v135
                                                            }
                                                        case -2147483648:
                                                            v136 := GenS5{F0: 1000, F1: 11, F2: v120}
                                                            _ = v136
                                                        case 32767:
                                                            var v137 int32 = 432
                                                            _ = v137
                                                        default:
                                                            v138 := GenS7{F0: 1, F1: v116}
                                                            _ = v138
                                                        }
                                                        _ = v129
                                                    }
                                                    _ = v121
                                                }
                                                _ = v120
                                            }
                                        default:
                                            fmt.Println(v1, v2)
                                        }
                                    }
                                    v116 = false
                                    _ = v117
                                } else {
                                    _ = len(v3)
                                    var v139 float64 = (1e10 + 0.0)
                                    v2 -= 65535
                                    _ = v139
                                }
                                _ = v116
                            default:
                                v140 := GenS9{F0: true, F1: 2}
                                _ = v140
                            }
                        }
                        _ = fi114
                        _ = v115
                    }
                }
                _ = fi103
            }
            v1 -= (^(v1))
            var v141 uint16 = 252
            _ = v141
        }
        for fi142 := 0; fi142 < 5; fi142++ {
            if false {
                switch fi142 {
                case -2147483648:
                    v3 = append(v3, v4)
                case 1:
                    v2 -= v2
                case 716:
                    for fi143 := 0; fi143 < 3; fi143++ {
                        var v144 int8 = v1
                        if true {
                            v2 = (^(v2))
                            switch fi142 {
                            case 128:
                                switch fi142 {
                                case 9223372036854775807:
                                    fmt.Println(fi142)
                                }
                                var v145 int64 = 42
                                var v146 float64 = -2.25
                                _ = v145
                                _ = v146
                            }
                        } else {
                            var v147 int32 = 48
                            v1 += -1
                            _ = v147
                        }
                        _ = fi143
                        _ = v144
                    }
                }
                if true {
                    for fi148 := 0; fi148 < 1; fi148++ {
                        v149 := GenS19{F0: 128, F1: 0}
                        var v150 int16 = 0
                        var v151 bool = true
                        _ = fi148
                        _ = v149
                        _ = v150
                        _ = v151
                    }
                    if (991 != 287) {
                        v1 -= (v1 & 32)
                        v1 = (v1 | v1)
                    }
                    if len(v3) > 0 {
                        v3[0] = 2
                    }
                } else {
                    var v152 float32 = 0.1
                    _ = v152
                }
            }
            _ = fi142
        }
    } else {
        for fi153 := 0; fi153 < 19; fi153++ {
            if (int(v4) != (fi153 & 9223372036854775807)) {
                var v154 int64 = 1
                if true {
                    v154 += 81
                    var v155 int8 = (v1 * v1)
                    _ = v155
                }
                _ = v154
            } else {
                if (-32768 > 32767) {
                    var v156 uint32 = 656
                    switch (fi153 >> ((uint(v4)) & 63)) {
                    case -25:
                        if true {
                            var v157 int32 = 0
                            var v158 float32 = float32(1.5)
                            var v159 uint32 = (^(v4))
                            _ = v157
                            _ = v158
                            _ = v159
                        } else {
                            v160 := GenS4{F0: (v1 & (v1 << ((65535) & 63))), F1: 0, F2: 4294967295}
                            var v161 uint32 = v156
                            if ((v156 | (v156 << ((240) & 63))) < (v4 % ((42) | 1))) {
                                v162 := GenS6{F0: v2, F1: v2}
                                v161 = 694
                                var v163 bool = false
                                _ = v162
                                _ = v163
                            } else {
                                fmt.Println(fi153, v161)
                            }
                            _ = v160
                            _ = v161
                        }
                    case 127:
                        v164 := GenS15{F0: 1, F1: "a", F2: 2147483647}
                        _ = v164
                    case 2147483647:
                        for fi165 := 0; fi165 < 4; fi165++ {
                            _ = len(v3)
                            _ = len(v3)
                            _ = fi165
                        }
                    }
                    _ = v156
                } else {
                    var v166 int8 = v1
                    _ = v166
                }
                for fi167 := 0; fi167 < 11; fi167++ {
                    fmt.Println(v2, v4)
                    v2 += v2
                    switch (fi167 >> ((0) & 63)) {
                    case -741:
                        var v168 string = "\n\t"
                        _ = v168
                    }
                    _ = fi167
                }
                v169 := GenS20{F0: "foo bar"}
                _ = v169
            }
            _ = fi153
        }
    }
    v170 := GenS14{F0: ((2 < 2) && true), F1: -1, F2: "foo bar"}
    v171 := GenS12{F0: -130, F1: false, F2: -44}
    var v172 bool = (!true)
    fmt.Println("v1:", v1)
    fmt.Println("v172:", v172)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v170
    _ = v171
    _ = v172
    _ = v2
    _ = v3
    _ = v4
}
