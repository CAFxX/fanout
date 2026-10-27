package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int16 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint16; F1 uint; F2 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 int8; F2 uint32 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 bool; F1 int; F2 int16 }
type GenS8 struct { F0 uint16; F1 int16 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint8; F1 bool }
type GenS11 struct { F0 uint64; F1 int64; F2 uint }
type GenS12 struct { F0 bool; F1 int }
type GenS13 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS14 struct { F0 int16; F1 uint64; F2 string }
type GenS15 struct { F0 int8; F1 int64; F2 uint }
type GenS16 struct { F0 uint8; F1 int32 }
type GenS17 struct { F0 string; F1 int; F2 int64 }
type GenS18 struct { F0 string; F1 int32 }
type GenS19 struct { F0 int8; F1 uint8 }
type GenS20 struct { F0 uint16 }
type GenS21 struct { F0 uint32 }
type GenS22 struct { F0 uint64; F1 int16 }
type GenS23 struct { F0 int32; F1 uint }
type GenS24 struct { F0 int32; F1 uint64 }
type GenS25 struct { F0 int32 }
type GenS26 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS27 struct { F0 uint32; F1 int64; F2 string }
type GenS28 struct { F0 int; F1 uint }
type GenS29 struct { F0 uint32; F1 int64; F2 int64 }
type GenS30 struct { F0 int64 }

func main() {
    var v1 uint64 = 966
    v2 := make(map[int]uint32)
    v3 := []string{}
    var v4 uint8 = 0
    v5 := make(map[int]uint32)
    v6 := GenS1{F0: 42, F1: 0, F2: 42}
    if false {
        if true {
            v7 := GenS2{F0: -691, F1: true}
            if len(v3) > 0 {
                v3[0] = ""
            }
            for fi8 := 0; fi8 < 1; fi8++ {
                if (750 > 42) {
                    var v9 int64 = 9223372036854775807
                    _ = v9
                } else {
                    switch fi8 {
                    case 0:
                        v10 := GenS3{F0: 0, F1: 1, F2: ("a" + "hello")}
                        _ = v10
                    case -2147483648:
                        for fi11 := 0; fi11 < 16; fi11++ {
                            if true {
                                var v12 int32 = -2147483648
                                _ = v12
                            } else {
                                v13 := GenS4{F0: 528}
                                var v14 int32 = -1
                                var v15 string = "hello"
                                _ = v13
                                _ = v14
                                _ = v15
                            }
                            v16 := GenS1{F0: 295, F1: 42, F2: 32767}
                            var v17 bool = true
                            _ = fi11
                            _ = v16
                            _ = v17
                        }
                    default:
                        v1 = 175
                    }
                    var v18 float32 = 1e10
                    switch ("a" + "") {
                    case "a":
                        var v19 uint16 = 1
                        _ = v19
                    default:
                        v4 = v4
                    }
                    _ = v18
                }
                if false {
                    switch 603 {
                    case 1:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case -129:
                        v20 := v2[(^(fi8))]
                        _ = v20
                    case 42:
                        var v21 int32 = 2147483647
                        _ = v21
                    default:
                        var v22 uint8 = (v4 >> ((42) & 63))
                        _ = v22
                    }
                } else {
                    var v23 bool = false
                    _ = v23
                }
                v1 -= v1
                _ = fi8
            }
            _ = v7
        } else {
            v4 = (v4 & (v4 % (((v4 / (((v4 - 97)) | 1))) | 1)))
        }
        if true {
            for fi24 := 0; fi24 < 1; fi24++ {
                for fi25 := 0; fi25 < 8; fi25++ {
                    v26 := GenS5{F0: 9223372036854775807, F1: 1, F2: 98}
                    _ = len(v5)
                    switch (fi24 >> ((42) & 63)) {
                    case 4294967295:
                        v4 -= 0
                    }
                    _ = fi25
                    _ = v26
                }
                _ = fi24
            }
            if true {
                var v27 int16 = 0
                _ = v27
            } else {
                v4 = 116
                fmt.Println(v1, v4)
            }
        } else {
            var v28 uint16 = 2
            switch 127 {
            case -851:
                var v29 float32 = (1.5 / (float32((0.1 + ((0.0 + float64((1.5 - float32(3.14159)))) / 0.1))) + (float32((0.1 * 1e10)) * 0.0)))
                _ = v29
            case 128:
                if ((v4 == v4) || false) {
                    for fi30 := 0; fi30 < 11; fi30++ {
                        if ((v4 * v4) < v4) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v3 = append(v3, "hello")
                        } else {
                            switch (("\n\t" + "a") + "\n\t") {
                            case "a":
                                v4 -= 0
                            default:
                                var v31 int = (fi30 - (fi30 * (fi30 << ((255) & 63))))
                                _ = v31
                            }
                        }
                        v4 -= 127
                        _ = fi30
                    }
                    for fi32 := 0; fi32 < 6; fi32++ {
                        switch "" {
                        case "":
                            v1 += v1
                        case "a":
                            v1 += v1
                        }
                        _ = fi32
                    }
                    v3 = append(v3, ("a" + "a"))
                }
            default:
                v1 -= 2
            }
            var v33 int32 = 321
            _ = v28
            _ = v33
        }
        fmt.Println(v4, v4)
    } else {
        var v34 uint8 = (^(v4))
        _ = v34
    }
    switch "a" {
    case "":
        v2[42] = 1
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "a":
        v3 = append(v3, (("a" + "a") + "a"))
    default:
        _ = len(v3)
    }
    v3 = append(v3, "foo bar")
    v1 -= v1
    if true {
        var v35 int = -2147483648
        _ = v35
    } else {
        if ((1e10 + (-2.25 * float32(3.14159))) == 3.14159) {
            for fi36 := 0; fi36 < 3; fi36++ {
                v37 := v2[(fi36 + (fi36 - (fi36 - (fi36 + fi36))))]
                if (-609 == 9223372036854775807) {
                    v37 = v37
                } else {
                    fmt.Println(v37, fi36)
                    switch 2147483647 {
                    case 14:
                        var v38 uint = 157
                        _ = v38
                    case 0:
                        _ = len(v2)
                    case 127:
                        _ = len(v3)
                    default:
                        v1 = (v1 << ((0) & 63))
                    }
                }
                _ = fi36
                _ = v37
            }
        } else {
            for fi39 := 0; fi39 < 16; fi39++ {
                v4 += (v4 + (^(v4)))
                if true {
                    v4 = v4
                }
                _ = fi39
            }
            for fi40 := 0; fi40 < 20; fi40++ {
                v4 = uint8(v1)
                v41 := GenS6{F0: 333, F1: (^(v4))}
                _ = fi40
                _ = v41
            }
            for fi42 := 0; fi42 < 8; fi42++ {
                var v43 int16 = 1000
                v44 := GenS6{F0: 307, F1: 42}
                _ = fi42
                _ = v43
                _ = v44
            }
        }
        for fi45 := 0; fi45 < 19; fi45++ {
            v1 -= (v1 / ((v1) | 1))
            v46 := GenS4{F0: -32768}
            _ = fi45
            _ = v46
        }
        if true {
            switch 575 {
            case 42:
                if true {
                    var v47 float64 = (0.0 * (-2.25 * (-2.25 + (3.14159 + 0.1))))
                    _ = v47
                } else {
                    fmt.Println(v1, v4)
                    for fi48 := 0; fi48 < 17; fi48++ {
                        for fi49 := 0; fi49 < 2; fi49++ {
                            if (!true) {
                                if false {
                                    delete(v5, 0)
                                    v4 += v4
                                    if true {
                                        var v50 int64 = 279
                                        if (!(!((v4 << ((2) & 63)) != (v4 / (((v4 * v4)) | 1))))) {
                                            switch ("foo bar" + "hello") {
                                            case "":
                                                switch 32768 {
                                                case 127:
                                                    v51 := v2[fi49]
                                                    _ = v51
                                                case 2147483647:
                                                    if (946 == 255) {
                                                        if ((1.5 + 0.1) < 0.0) {
                                                            v4 += v4
                                                            var v52 uint8 = 255
                                                            var v53 float32 = 3.14159
                                                            _ = v52
                                                            _ = v53
                                                        } else {
                                                            var v54 uint8 = (^(v4))
                                                            v55 := GenS7{F0: (!(((!false) || false) && false)), F1: fi48, F2: -688}
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                        switch "a" {
                                                        case "":
                                                            fmt.Println(fi49)
                                                        default:
                                                            if (true && true) {
                                                                var v56 uint8 = 1
                                                                v50 -= v50
                                                                _ = v56
                                                            }
                                                        }
                                                        if false {
                                                            v1 -= 476
                                                            v57 := GenS1{F0: 2, F1: -1, F2: -607}
                                                            _ = v57
                                                        } else {
                                                            fmt.Println(v4, fi49, fi48)
                                                            if len(v3) > 0 {
                                                                v3[0] = "hello"
                                                            }
                                                            fmt.Println(v1)
                                                        }
                                                        switch fi49 {
                                                        case 127:
                                                            fmt.Println(fi49, fi48, v4)
                                                        case 2147483647:
                                                            v58 := GenS7{F0: true, F1: (-(fi49)), F2: 0}
                                                            _ = v58
                                                        case 23:
                                                            _ = len(v5)
                                                        default:
                                                            if (!(v4 > (^(v4)))) {
                                                                var v59 float64 = (0.1 - (0.1 - 1.5))
                                                                var v60 uint8 = 68
                                                                _ = v59
                                                                _ = v60
                                                            } else {
                                                                var v61 uint32 = 2147483648
                                                                _ = v61
                                                            }
                                                        }
                                                        v4 = (v4 / ((v4) | 1))
                                                        if (0 != 4294967295) {
                                                            v50 -= v50
                                                            fmt.Println(v1)
                                                            v1 = (^(v1))
                                                            v4 -= v4
                                                        } else {
                                                            if true {
                                                                switch fi49 {
                                                                case 42:
                                                                    v62 := GenS8{F0: 2, F1: 32767}
                                                                    _ = v62
                                                                case 127:
                                                                    var v63 uint16 = 676
                                                                    _ = v63
                                                                default:
                                                                    v4 = 2
                                                                }
                                                                v4 = v4
                                                                _ = len(v2)
                                                            } else {
                                                                var v64 uint8 = 203
                                                                _ = v64
                                                            }
                                                        }
                                                        switch "a" {
                                                        case "":
                                                            var v65 float32 = 0.0
                                                            _ = v65
                                                        case "hello":
                                                            v50 -= (v50 << ((65535) & 63))
                                                        case "\n\t":
                                                            if (false && false) {
                                                                var v66 int64 = v50
                                                                var v67 int8 = -113
                                                                var v68 int8 = (v67 + v67)
                                                                fmt.Println(v68)
                                                                _ = v66
                                                                _ = v67
                                                                _ = v68
                                                            }
                                                        }
                                                    } else {
                                                        v1 -= (v1 ^ v1)
                                                        if false {
                                                            v69 := GenS3{F0: 42, F1: 18446744073709551615, F2: "\n\t"}
                                                            v4 -= v4
                                                            v1 = v1
                                                            _ = v69
                                                        }
                                                    }
                                                }
                                            case "\n\t":
                                                v70 := GenS9{F0: -2147483648}
                                                _ = v70
                                            case "foo bar":
                                                v71 := GenS10{F0: v4, F1: (!true)}
                                                _ = v71
                                            }
                                        }
                                        if (v1 < 1) {
                                            switch "\n\t" {
                                            case "\n\t":
                                                v4 = 42
                                            default:
                                                if true {
                                                    if ((65535 <= 577) || (!((42 >= uint16(v4)) && (-345 > 0)))) {
                                                        v72 := GenS11{F0: v1, F1: (v50 << ((2) & 63)), F2: 1}
                                                        _ = v72
                                                    } else {
                                                        var v73 uint64 = (v1 ^ (v1 * 2))
                                                        var v74 bool = ((((^(v4)) >= v4) || (true && (false || true))) && ((("\n\t" + "\n\t") >= "foo bar") && true))
                                                        _ = v73
                                                        _ = v74
                                                    }
                                                    var v75 float32 = 1e10
                                                    _ = v75
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = ""
                                                    }
                                                }
                                            }
                                            if (1 == 0) {
                                                if ((v4 << ((4294967295) & 63)) < (v4 << ((18446744073709551615) & 63))) {
                                                    var v76 int = fi48
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = v76
                                                }
                                            }
                                        }
                                        switch fi49 {
                                        case 4294967295:
                                            if false {
                                                var v77 int16 = 1
                                                _ = v77
                                            } else {
                                                if false {
                                                    var v78 int = -213
                                                    _ = v78
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    switch (fi48 + 9223372036854775807) {
                                                    case -458:
                                                        var v79 int64 = v50
                                                        _ = v79
                                                    case 9223372036854775807:
                                                        switch fi49 {
                                                        case 450:
                                                            v1 += (v1 + 494)
                                                        case -769:
                                                            if true {
                                                                if true {
                                                                    if true {
                                                                        switch (fi49 | 2147483647) {
                                                                        case 2147483647:
                                                                            var v80 uint32 = 4294967295
                                                                            _ = v80
                                                                        case -725:
                                                                            if (!(true || (-113 >= -32768))) {
                                                                                v1 = 97
                                                                                if true {
                                                                                    switch 32767 {
                                                                                    case 42:
                                                                                        switch 128 {
                                                                                        case -1:
                                                                                            var v81 int8 = 127
                                                                                            _ = v81
                                                                                        case -801:
                                                                                            var v82 float64 = (1e10 / 1e10)
                                                                                            _ = v82
                                                                                        default:
                                                                                            switch (fi49 + fi48) {
                                                                                            case -2147483648:
                                                                                                v1 -= (^(v1))
                                                                                                switch fi49 {
                                                                                                case 1:
                                                                                                    switch (-(fi48)) {
                                                                                                    case 4294967295:
                                                                                                        v83 := GenS12{F0: (true && true), F1: fi49}
                                                                                                        _ = v83
                                                                                                    case -2147483648:
                                                                                                        if true {
                                                                                                            v84 := GenS1{F0: 42, F1: 42, F2: -1}
                                                                                                            _ = v84
                                                                                                        } else {
                                                                                                            var v85 uint32 = 0
                                                                                                            _ = v85
                                                                                                        }
                                                                                                        var v86 float64 = 1e10
                                                                                                        v87 := GenS11{F0: (v1 >> ((255) & 63)), F1: 1, F2: 18446744073709551615}
                                                                                                        _ = v86
                                                                                                        _ = v87
                                                                                                    case 32767:
                                                                                                        var v88 int64 = (v50 - (v50 + (v50 - -9223372036854775808)))
                                                                                                        _ = v88
                                                                                                    default:
                                                                                                        switch fi49 {
                                                                                                        case -453:
                                                                                                            v1 = v1
                                                                                                        case 4294967295:
                                                                                                            if (fi48 != -2147483648) {
                                                                                                                switch "\n\t" {
                                                                                                                case "":
                                                                                                                    if (false && (!(true || (359 < fi48)))) {
                                                                                                                        var v89 float64 = 0.1
                                                                                                                        v90 := GenS13{F0: (v1 & 0), F1: 4294967295, F2: -85}
                                                                                                                        fmt.Println(v89, v89, fi48)
                                                                                                                        if false {
                                                                                                                            switch -367 {
                                                                                                                            case 9223372036854775807:
                                                                                                                                var v91 uint32 = 4294967295
                                                                                                                                _ = v91
                                                                                                                            case 1:
                                                                                                                                v92 := GenS6{F0: 65535, F1: 150}
                                                                                                                                var v93 bool = true
                                                                                                                                v4 -= 46
                                                                                                                                _ = v92
                                                                                                                                _ = v93
                                                                                                                            case 32767:
                                                                                                                                fmt.Println(fi49, v4, fi49)
                                                                                                                                var v94 uint8 = v4
                                                                                                                                _ = v94
                                                                                                                            }
                                                                                                                            var v95 float64 = float64(3.14159)
                                                                                                                            var v96 int = (fi48 - 32767)
                                                                                                                            switch 273 {
                                                                                                                            case 4294967295:
                                                                                                                                var v97 int64 = -9223372036854775808
                                                                                                                                _ = v97
                                                                                                                            case 9223372036854775807:
                                                                                                                                v98 := GenS5{F0: v50, F1: 42, F2: 2147483647}
                                                                                                                                _ = v98
                                                                                                                            case 1:
                                                                                                                                v95 = float64((0.1 + -2.25))
                                                                                                                            }
                                                                                                                            _ = v95
                                                                                                                            _ = v96
                                                                                                                        } else {
                                                                                                                            switch ("\n\t" + "a") {
                                                                                                                            case "foo bar":
                                                                                                                                v4 += v4
                                                                                                                            case "hello":
                                                                                                                                if len(v3) >= 2 {
                                                                                                                                    v3 = v3[1:len(v3)]
                                                                                                                                }
                                                                                                                                var v99 uint8 = 47
                                                                                                                                _ = v99
                                                                                                                            default:
                                                                                                                                v100 := GenS3{F0: 32767, F1: 4294967295, F2: ("" + "a")}
                                                                                                                                _ = v100
                                                                                                                            }
                                                                                                                        }
                                                                                                                        _ = v89
                                                                                                                        _ = v90
                                                                                                                    } else {
                                                                                                                        var v101 int32 = 390
                                                                                                                        _ = v101
                                                                                                                    }
                                                                                                                default:
                                                                                                                    if len(v3) >= 2 {
                                                                                                                        v3 = v3[1:len(v3)]
                                                                                                                    }
                                                                                                                }
                                                                                                                fmt.Println(v4, fi48)
                                                                                                                v4 -= (v4 / ((42) | 1))
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    switch "hello" {
                                                                                                    case "foo bar":
                                                                                                        if false {
                                                                                                            v50 += 1
                                                                                                        } else {
                                                                                                            v4 = v4
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        if ("\n\t" <= "") {
                                                                                                            var v102 uint8 = 42
                                                                                                            var v103 int16 = 1
                                                                                                            if true {
                                                                                                                fmt.Println(v50, v50, v102)
                                                                                                                v104 := GenS7{F0: true, F1: 975, F2: 42}
                                                                                                                var v105 uint16 = 1
                                                                                                                _ = v104
                                                                                                                _ = v105
                                                                                                            }
                                                                                                            _ = v102
                                                                                                            _ = v103
                                                                                                        } else {
                                                                                                            switch (fi49 << ((42) & 63)) {
                                                                                                            case 440:
                                                                                                                if (false && true) {
                                                                                                                    var v106 float64 = 1.5
                                                                                                                    _ = v106
                                                                                                                }
                                                                                                                v4 += 139
                                                                                                                fmt.Println(fi49, v1, v50)
                                                                                                                var v107 int32 = -903
                                                                                                                _ = v107
                                                                                                            case 32767:
                                                                                                                v108 := GenS10{F0: v4, F1: false}
                                                                                                                _ = v108
                                                                                                            case 9223372036854775807:
                                                                                                                v4 = (v4 & v4)
                                                                                                            default:
                                                                                                                switch (-(fi48)) {
                                                                                                                case 0:
                                                                                                                    var v109 float64 = 1e10
                                                                                                                    _ = v109
                                                                                                                case 127:
                                                                                                                    v1 = (v1 - (v1 % ((uint64(v4)) | 1)))
                                                                                                                default:
                                                                                                                    v4 = (v4 + 2)
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    case "\n\t":
                                                                                                        v50 = (v50 * v50)
                                                                                                    }
                                                                                                }
                                                                                                var v110 uint64 = (v1 % ((0) | 1))
                                                                                                _ = v110
                                                                                            case -840:
                                                                                                fmt.Println(v50)
                                                                                            case 42:
                                                                                                v111 := GenS11{F0: (v1 >> ((255) & 63)), F1: (v50 - 944), F2: 42}
                                                                                                _ = v111
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        if true {
                                                                                            if false {
                                                                                                if (((fi49 + (fi49 / (((-(fi48))) | 1))) > fi48) || false) {
                                                                                                    var v112 int = 860
                                                                                                    if false {
                                                                                                        v50 = v50
                                                                                                        switch 0 {
                                                                                                        case 622:
                                                                                                            if (false || true) {
                                                                                                                var v113 uint32 = 4294967295
                                                                                                                delete(v2, fi48)
                                                                                                                var v114 int8 = 15
                                                                                                                var v115 float32 = 1.5
                                                                                                                _ = v113
                                                                                                                _ = v114
                                                                                                                _ = v115
                                                                                                            } else {
                                                                                                                var v116 uint8 = 39
                                                                                                                var v117 float32 = -2.25
                                                                                                                _ = v116
                                                                                                                _ = v117
                                                                                                            }
                                                                                                        default:
                                                                                                            if len(v3) >= 2 {
                                                                                                                v3 = v3[1:len(v3)]
                                                                                                            }
                                                                                                        }
                                                                                                        var v118 uint64 = v1
                                                                                                        _ = v118
                                                                                                    } else {
                                                                                                        v112 = (fi49 >> ((598) & 63))
                                                                                                        if (false || (!(255 != v4))) {
                                                                                                            v112 += fi48
                                                                                                            if true {
                                                                                                                v50 -= (v50 & 9223372036854775807)
                                                                                                                v119 := GenS2{F0: 0, F1: false}
                                                                                                                var v120 int = (fi48 ^ -447)
                                                                                                                v1 = (^(v1))
                                                                                                                if true {
                                                                                                                    v4 = v4
                                                                                                                } else {
                                                                                                                    var v121 uint = 644
                                                                                                                    if true {
                                                                                                                        v122 := GenS7{F0: (!false), F1: 42, F2: 42}
                                                                                                                        _ = v122
                                                                                                                    } else {
                                                                                                                        v1 += v1
                                                                                                                        _ = len(v3)
                                                                                                                        switch "" {
                                                                                                                        case "\n\t":
                                                                                                                            var v123 float64 = float64(0.0)
                                                                                                                            _ = v123
                                                                                                                        case "hello":
                                                                                                                            v3 = append(v3, "hello")
                                                                                                                        }
                                                                                                                        v124 := GenS5{F0: 0, F1: -117, F2: 1}
                                                                                                                        _ = v124
                                                                                                                    }
                                                                                                                    if (2147483647 >= -767) {
                                                                                                                        switch ("foo bar" + "hello") {
                                                                                                                        case "foo bar":
                                                                                                                            var v125 int8 = -128
                                                                                                                            _ = v125
                                                                                                                        case "":
                                                                                                                            v126 := GenS11{F0: 499, F1: v50, F2: v121}
                                                                                                                            _ = v126
                                                                                                                        }
                                                                                                                        delete(v5, fi48)
                                                                                                                    }
                                                                                                                    v3 = append(v3, ("" + "foo bar"))
                                                                                                                    v121 = 202
                                                                                                                    _ = v121
                                                                                                                }
                                                                                                                _ = v119
                                                                                                                _ = v120
                                                                                                            } else {
                                                                                                                var v127 int32 = -1
                                                                                                                v128 := GenS12{F0: (((fi49 - 269) <= -1) && false), F1: -604}
                                                                                                                v112 -= v112
                                                                                                                _ = v127
                                                                                                                _ = v128
                                                                                                            }
                                                                                                            if false {
                                                                                                                var v129 bool = false
                                                                                                                _ = v129
                                                                                                            } else {
                                                                                                                var v130 int = v112
                                                                                                                _ = v130
                                                                                                            }
                                                                                                            var v131 int32 = 42
                                                                                                            _ = v131
                                                                                                        } else {
                                                                                                            switch ("hello" + "foo bar") {
                                                                                                            case "a":
                                                                                                                v132 := GenS5{F0: v50, F1: int8(v1), F2: 2}
                                                                                                                _ = v132
                                                                                                            case "foo bar":
                                                                                                                v1 = 29
                                                                                                                switch (fi48 >> ((42) & 63)) {
                                                                                                                case 432:
                                                                                                                    if true {
                                                                                                                        fmt.Println(fi49)
                                                                                                                        _ = len(v2)
                                                                                                                        v133 := GenS14{F0: 1000, F1: 1, F2: "\n\t"}
                                                                                                                        _ = v133
                                                                                                                    }
                                                                                                                default:
                                                                                                                    v134 := GenS15{F0: 0, F1: (v50 ^ v50), F2: 1}
                                                                                                                    _ = v134
                                                                                                                }
                                                                                                            default:
                                                                                                                v135 := v2[fi49]
                                                                                                                _ = v135
                                                                                                            }
                                                                                                            fmt.Println(v50, v50)
                                                                                                        }
                                                                                                        if true {
                                                                                                            v136 := GenS5{F0: (-(v50)), F1: 100, F2: 2147483648}
                                                                                                            var v137 int8 = int8(fi49)
                                                                                                            v1 = v1
                                                                                                            v138 := GenS16{F0: (^(v4)), F1: 42}
                                                                                                            switch "" {
                                                                                                            case "a":
                                                                                                                v139 := GenS14{F0: 1, F1: v1, F2: "foo bar"}
                                                                                                                _ = v139
                                                                                                            case "hello":
                                                                                                                v140 := GenS15{F0: (v137 | v137), F1: int64(v137), F2: 2}
                                                                                                                _ = v140
                                                                                                            case "foo bar":
                                                                                                                var v141 bool = (42 < 255)
                                                                                                                _ = v141
                                                                                                            default:
                                                                                                                var v142 string = "foo bar"
                                                                                                                _ = v142
                                                                                                            }
                                                                                                            _ = v136
                                                                                                            _ = v137
                                                                                                            _ = v138
                                                                                                        } else {
                                                                                                            var v143 uint = 636
                                                                                                            _ = v143
                                                                                                        }
                                                                                                    }
                                                                                                    v144 := GenS17{F0: "\n\t", F1: (-(v112)), F2: 263}
                                                                                                    _ = v112
                                                                                                    _ = v144
                                                                                                }
                                                                                            } else {
                                                                                                if (821 < (v50 - -1)) {
                                                                                                    v4 -= 42
                                                                                                    var v145 uint8 = v4
                                                                                                    v50 -= v50
                                                                                                    _ = v145
                                                                                                }
                                                                                                if ("hello" < "foo bar") {
                                                                                                    v1 -= v1
                                                                                                    if false {
                                                                                                        if true {
                                                                                                            var v146 uint64 = (v1 + (^(v1)))
                                                                                                            switch (("a" + ("" + "hello")) + ("\n\t" + "a")) {
                                                                                                            case "hello":
                                                                                                                v147 := GenS14{F0: 42, F1: v1, F2: "hello"}
                                                                                                                _ = v147
                                                                                                            default:
                                                                                                                switch (("hello" + ("" + "a")) + ("foo bar" + "hello")) {
                                                                                                                case "hello":
                                                                                                                    if false {
                                                                                                                        delete(v2, -939)
                                                                                                                    } else {
                                                                                                                        v1 = 0
                                                                                                                        v4 = (v4 & v4)
                                                                                                                    }
                                                                                                                case "":
                                                                                                                    if ((false || false) || (-639 > -1)) {
                                                                                                                        var v148 int64 = (v50 >> ((0) & 63))
                                                                                                                        v50 += (v50 | 355)
                                                                                                                        var v149 int16 = 32767
                                                                                                                        _ = v148
                                                                                                                        _ = v149
                                                                                                                    } else {
                                                                                                                        v150 := v5[32768]
                                                                                                                        _ = v150
                                                                                                                    }
                                                                                                                case "\n\t":
                                                                                                                    v4 = (^(v4))
                                                                                                                default:
                                                                                                                    if false {
                                                                                                                        v1 = v146
                                                                                                                        v50 = v50
                                                                                                                    } else {
                                                                                                                        v1 = 18446744073709551615
                                                                                                                        v151 := GenS18{F0: "foo bar", F1: 0}
                                                                                                                        v146 -= (v146 ^ 0)
                                                                                                                        _ = v151
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v146
                                                                                                        } else {
                                                                                                            var v152 uint8 = v4
                                                                                                            _ = v152
                                                                                                        }
                                                                                                    }
                                                                                                    var v153 uint32 = 0
                                                                                                    _ = v153
                                                                                                }
                                                                                            }
                                                                                            if (uint32(v4) > 65535) {
                                                                                                if (!true) {
                                                                                                    v1 += v1
                                                                                                    v3 = append(v3, "a")
                                                                                                } else {
                                                                                                    switch "\n\t" {
                                                                                                    case "\n\t":
                                                                                                        fmt.Println(fi48, fi48, v50)
                                                                                                    case "foo bar":
                                                                                                        switch fi48 {
                                                                                                        case 428:
                                                                                                            switch fi48 {
                                                                                                            case 32767:
                                                                                                                var v154 int32 = 0
                                                                                                                _ = v154
                                                                                                            default:
                                                                                                                var v155 int16 = 32767
                                                                                                                _ = v155
                                                                                                            }
                                                                                                        }
                                                                                                    default:
                                                                                                        switch (fi49 << ((218) & 63)) {
                                                                                                        case -681:
                                                                                                            var v156 string = "\n\t"
                                                                                                            _ = v156
                                                                                                        case 486:
                                                                                                            if len(v3) >= 2 {
                                                                                                                v3 = v3[1:len(v3)]
                                                                                                            }
                                                                                                        default:
                                                                                                            switch (fi48 ^ 127) {
                                                                                                            case 4294967295:
                                                                                                                switch fi48 {
                                                                                                                case -2147483648:
                                                                                                                    var v157 uint = 18446744073709551615
                                                                                                                    _ = v157
                                                                                                                case 42:
                                                                                                                    var v158 int64 = v50
                                                                                                                    _ = v158
                                                                                                                default:
                                                                                                                    v50 = 815
                                                                                                                    v1 = 1
                                                                                                                    v2[fi48] = 0
                                                                                                                    switch "foo bar" {
                                                                                                                    case "":
                                                                                                                        v50 += (v50 >> ((355) & 63))
                                                                                                                    case "foo bar":
                                                                                                                        v159 := GenS19{F0: -50, F1: (v4 * (v4 | (v4 | v4)))}
                                                                                                                        _ = v159
                                                                                                                    }
                                                                                                                }
                                                                                                            case 73:
                                                                                                                var v160 float32 = 0.1
                                                                                                                _ = v160
                                                                                                            default:
                                                                                                                switch ("foo bar" + (("" + "") + "")) {
                                                                                                                case "foo bar":
                                                                                                                    var v161 bool = false
                                                                                                                    _ = v161
                                                                                                                default:
                                                                                                                    fmt.Println(v4, v4)
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    var v162 int = (fi48 << ((18446744073709551615) & 63))
                                                                                                    _ = v162
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    if false {
                                                                                        v163 := GenS4{F0: -32768}
                                                                                        _ = v163
                                                                                    }
                                                                                }
                                                                                if (v4 <= 0) {
                                                                                    if true {
                                                                                        if (-1 == 2147483647) {
                                                                                            if ((float32(0.1) * 0.0) == float32(-2.25)) {
                                                                                                v164 := GenS10{F0: v4, F1: false}
                                                                                                var v165 uint = 1
                                                                                                _ = v164
                                                                                                _ = v165
                                                                                            }
                                                                                            v1 = 0
                                                                                        } else {
                                                                                            fmt.Println(v4, v50, v1)
                                                                                            var v166 uint16 = 42
                                                                                            v2[657] = 0
                                                                                            _ = v166
                                                                                        }
                                                                                    }
                                                                                    if true {
                                                                                        fmt.Println(v50, v1)
                                                                                    }
                                                                                } else {
                                                                                    v167 := GenS20{F0: 255}
                                                                                    v168 := GenS7{F0: (((v4 ^ v4) > (v4 - 0)) && (1 == (^(fi49)))), F1: (fi48 | fi48), F2: -32768}
                                                                                    fmt.Println(v1, fi49, v4)
                                                                                    _ = v167
                                                                                    _ = v168
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    var v169 int64 = int64(v4)
                                                                    _ = v169
                                                                }
                                                            } else {
                                                                switch "" {
                                                                case "foo bar":
                                                                    v170 := v5[fi48]
                                                                    _ = v170
                                                                case "hello":
                                                                    v50 = (v50 & 0)
                                                                default:
                                                                    v4 += (v4 + (v4 % ((0) | 1)))
                                                                }
                                                            }
                                                        }
                                                        v5[(fi49 * (fi49 + int(v4)))] = 4294967295
                                                        switch (fi48 >> ((1) & 63)) {
                                                        case 32768:
                                                            var v171 uint64 = 1
                                                            _ = v171
                                                        case 170:
                                                            var v172 uint16 = 2
                                                            _ = v172
                                                        default:
                                                            if (true && true) {
                                                                switch "a" {
                                                                case "a":
                                                                    if (!false) {
                                                                        var v173 int8 = 1
                                                                        var v174 string = ("" + "\n\t")
                                                                        fmt.Println(v173, fi49)
                                                                        _ = v173
                                                                        _ = v174
                                                                    } else {
                                                                        v175 := GenS14{F0: 42, F1: 712, F2: "a"}
                                                                        _ = v175
                                                                    }
                                                                case "foo bar":
                                                                    switch ("" + "") {
                                                                    case "\n\t":
                                                                        switch ("hello" + ("" + ((("" + "hello") + "a") + "hello"))) {
                                                                        case "foo bar":
                                                                            if false {
                                                                                if false {
                                                                                    v1 = v1
                                                                                    v50 = v50
                                                                                    var v176 int = (^(fi48))
                                                                                    _ = len(v3)
                                                                                    switch "a" {
                                                                                    case "foo bar":
                                                                                        v4 = v4
                                                                                    case "hello":
                                                                                        if (-2147483648 == 664) {
                                                                                            var v177 int64 = v50
                                                                                            v178 := GenS21{F0: 800}
                                                                                            _ = v177
                                                                                            _ = v178
                                                                                        } else {
                                                                                            v179 := v2[fi49]
                                                                                            v180 := GenS4{F0: -995}
                                                                                            v4 += v4
                                                                                            _ = v179
                                                                                            _ = v180
                                                                                        }
                                                                                    case "a":
                                                                                        _ = len(v5)
                                                                                    }
                                                                                    _ = v176
                                                                                } else {
                                                                                    v181 := GenS10{F0: 0, F1: false}
                                                                                    var v182 float32 = 1e10
                                                                                    _ = v181
                                                                                    _ = v182
                                                                                }
                                                                                v4 = (v4 / (((^(v4))) | 1))
                                                                                switch "" {
                                                                                case "foo bar":
                                                                                    switch 9223372036854775807 {
                                                                                    case 9223372036854775807:
                                                                                        var v183 float32 = ((3.14159 / float32((1.5 + -2.25))) * ((1e10 + (1e10 * 3.14159)) / (-2.25 + 0.1)))
                                                                                        _ = v183
                                                                                    default:
                                                                                        if (true && true) {
                                                                                            v50 = v50
                                                                                            if (127 >= -115) {
                                                                                                v4 -= (v4 & v4)
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        case "":
                                                                            var v184 uint32 = 42
                                                                            _ = v184
                                                                        default:
                                                                            var v185 uint16 = uint16(v1)
                                                                            _ = v185
                                                                        }
                                                                        var v186 string = "foo bar"
                                                                        switch fi48 {
                                                                        case -663:
                                                                            var v187 uint8 = v4
                                                                            _ = v187
                                                                        case -9223372036854775808:
                                                                            v50 += v50
                                                                        case 918:
                                                                            if false {
                                                                                v3 = append(v3, ("foo bar" + v186))
                                                                                v1 = (^(v1))
                                                                                _ = len(v3)
                                                                                var v188 float64 = ((3.14159 + float64(1e10)) + 0.1)
                                                                                _ = v188
                                                                            } else {
                                                                                switch fi49 {
                                                                                case 47:
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                case 42:
                                                                                    v189 := GenS19{F0: -128, F1: 127}
                                                                                    _ = v189
                                                                                case 2147483647:
                                                                                    fmt.Println(fi48, fi49)
                                                                                    v190 := GenS13{F0: v1, F1: v1, F2: 101}
                                                                                    _ = v190
                                                                                default:
                                                                                    var v191 uint16 = 32767
                                                                                    _ = v191
                                                                                }
                                                                                v186 = (("" + v186) + ("hello" + (v186 + ((((v186 + v186) + v186) + v186) + "\n\t"))))
                                                                            }
                                                                            var v192 uint32 = 209
                                                                            fmt.Println(v1, v186, v192)
                                                                            _ = v192
                                                                        }
                                                                        _ = v186
                                                                    default:
                                                                        v50 = 1
                                                                    }
                                                                case "\n\t":
                                                                    var v193 bool = false
                                                                    _ = v193
                                                                default:
                                                                    var v194 float64 = (float64(-2.25) * float64(((float32(((-2.25 / 1.0) / 1.0)) / (((0.1 * (0.1 * 1e10)) * (-2.25 * 1.5)) - (0.1 * ((float32(0.1) / (3.14159 + 3.14159)) / (float32(float64(float32(1e10))) + float32(1.5)))))) * 3.14159)))
                                                                    _ = v194
                                                                }
                                                                v50 -= (v50 << ((835) & 63))
                                                            }
                                                        }
                                                    default:
                                                        v4 = (v4 ^ v4)
                                                    }
                                                    var v195 uint32 = 936
                                                    var v196 uint32 = v195
                                                    v5[(fi49 >> ((676) & 63))] = v196
                                                    _ = v195
                                                    _ = v196
                                                }
                                                v1 -= (v1 % ((v1) | 1))
                                            }
                                        case -1:
                                            if false {
                                                v197 := GenS18{F0: "foo bar", F1: -220}
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v197
                                            } else {
                                                switch (("hello" + "hello") + "\n\t") {
                                                case "a":
                                                    v50 -= v50
                                                default:
                                                    v198 := GenS22{F0: v1, F1: 42}
                                                    _ = v198
                                                }
                                            }
                                        }
                                        v50 = v50
                                        var v199 int64 = 9223372036854775807
                                        v3 = append(v3, ("hello" + ""))
                                        _ = v199
                                        _ = v50
                                    } else {
                                        if false {
                                            switch (fi48 % (((fi48 + 2147483647)) | 1)) {
                                            case 128:
                                                if true {
                                                    v4 = (v4 | v4)
                                                    v1 = v1
                                                } else {
                                                    delete(v2, fi48)
                                                    if false {
                                                        v1 += 18446744073709551615
                                                        if (2 >= 4294967295) {
                                                            if false {
                                                                v200 := v2[-2147483648]
                                                                delete(v5, fi49)
                                                                switch (fi48 % ((fi49) | 1)) {
                                                                case 128:
                                                                    v201 := GenS15{F0: 1, F1: 934, F2: 42}
                                                                    _ = v201
                                                                case -917:
                                                                    v200 -= 344
                                                                case -890:
                                                                    v4 -= (v4 % (((v4 - 51)) | 1))
                                                                }
                                                                v200 += v200
                                                                _ = v200
                                                            } else {
                                                                var v202 string = ""
                                                                _ = v202
                                                            }
                                                        } else {
                                                            switch fi49 {
                                                            case 32767:
                                                                if true {
                                                                    if true {
                                                                        _ = len(v3)
                                                                        var v203 uint32 = 2147483648
                                                                        _ = v203
                                                                    } else {
                                                                        switch fi48 {
                                                                        case 32768:
                                                                            switch "" {
                                                                            case "foo bar":
                                                                                v1 = (v1 | (v1 * (^(v1))))
                                                                            case "a":
                                                                                v204 := GenS1{F0: 297, F1: 45, F2: 42}
                                                                                _ = v204
                                                                            default:
                                                                                v205 := v5[9223372036854775807]
                                                                                _ = v205
                                                                            }
                                                                        case -9223372036854775808:
                                                                            switch fi48 {
                                                                            case 32767:
                                                                                v1 += (v1 ^ (v1 | 4294967295))
                                                                            case -1:
                                                                                var v206 uint16 = 318
                                                                                _ = v206
                                                                            default:
                                                                                v1 += v1
                                                                            }
                                                                        case 127:
                                                                            v207 := v5[(^(fi48))]
                                                                            _ = v207
                                                                        }
                                                                    }
                                                                    var v208 uint64 = v1
                                                                    var v209 uint64 = v1
                                                                    _ = v208
                                                                    _ = v209
                                                                }
                                                            case 42:
                                                                fmt.Println(v4, v1)
                                                            case -2147483648:
                                                                var v210 uint = 0
                                                                _ = v210
                                                            default:
                                                                v4 -= (v4 / (((v4 | (v4 % ((1) | 1)))) | 1))
                                                            }
                                                        }
                                                        v4 = v4
                                                        v4 = v4
                                                        v4 = 90
                                                    }
                                                    fmt.Println(v1, v1)
                                                }
                                            default:
                                                v3 = append(v3, ((("" + "\n\t") + (("" + "a") + ("foo bar" + "a"))) + "foo bar"))
                                            }
                                            v211 := GenS10{F0: (^(v4)), F1: true}
                                            _ = v211
                                        }
                                        _ = len(v3)
                                    }
                                } else {
                                    var v212 string = ("hello" + "foo bar")
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v212
                                }
                                var v213 string = ""
                                v214 := GenS18{F0: (v213 + v213), F1: 664}
                                _ = v213
                                _ = v214
                            } else {
                                var v215 uint32 = 4294967295
                                var v216 int16 = 0
                                v1 = v1
                                _ = len(v5)
                                if false {
                                    if (v216 == (v216 << ((0) & 63))) {
                                        v216 -= (v216 % ((0) | 1))
                                    }
                                    switch ("a" + "\n\t") {
                                    case "\n\t":
                                        v215 = 857
                                    case "a":
                                        v1 -= v1
                                    case "foo bar":
                                        if true {
                                            v217 := GenS20{F0: 0}
                                            if ((2147483647 != 727) || true) {
                                                v218 := GenS17{F0: "foo bar", F1: fi48, F2: 9223372036854775807}
                                                _ = v218
                                            } else {
                                                var v219 uint = 65535
                                                _ = v219
                                            }
                                            v5[-2147483648] = 518
                                            _ = v217
                                        } else {
                                            v3 = append(v3, "foo bar")
                                        }
                                    default:
                                        v1 += (v1 * (v1 * v1))
                                    }
                                } else {
                                    var v220 uint64 = v1
                                    var v221 int8 = 127
                                    v222 := GenS23{F0: 2147483647, F1: 127}
                                    _ = v220
                                    _ = v221
                                    _ = v222
                                }
                                v223 := GenS24{F0: 0, F1: (v1 % ((42) | 1))}
                                v224 := GenS21{F0: 2147483648}
                                _ = v215
                                _ = v216
                                _ = v223
                                _ = v224
                            }
                            var v225 float32 = 1.5
                            v225 += (v225 + -2.25)
                            var v226 uint16 = 839
                            _ = fi49
                            _ = v225
                            _ = v226
                        }
                        for fi227 := 0; fi227 < 17; fi227++ {
                            if (-616 >= 1) {
                                v228 := GenS23{F0: 42, F1: 18446744073709551615}
                                v4 = v4
                                var v229 string = "a"
                                _ = v228
                                _ = v229
                            } else {
                                var v230 int64 = 9223372036854775807
                                _ = v230
                            }
                            v231 := GenS19{F0: 0, F1: 65}
                            if (false || true) {
                                var v232 float64 = 0.1
                                var v233 string = "\n\t"
                                if (!true) {
                                    v4 = v4
                                    delete(v5, (fi48 >> ((115) & 63)))
                                    var v234 uint8 = v4
                                    _ = v234
                                }
                                _ = v232
                                _ = v233
                            }
                            if (!(v1 <= 18446744073709551615)) {
                                v235 := GenS25{F0: 2147483647}
                                _ = v235
                            } else {
                                var v236 int = fi227
                                var v237 bool = (-1 == -1)
                                v238 := GenS8{F0: 346, F1: -132}
                                _ = v236
                                _ = v237
                                _ = v238
                            }
                            _ = fi227
                            _ = v231
                        }
                        _ = fi48
                    }
                    v1 += 42
                }
            case 9223372036854775807:
                for fi239 := 0; fi239 < 6; fi239++ {
                    for fi240 := 0; fi240 < 2; fi240++ {
                        var v241 uint32 = 0
                        _ = fi240
                        _ = v241
                    }
                    if (1 < -1) {
                        var v242 float64 = (((1e10 - (((((0.0 * (0.1 + (0.0 / 0.1))) + (3.14159 / 0.1)) / -2.25) + 0.1) / -2.25)) + 0.1) - -2.25)
                        v242 = 1.5
                        _ = v242
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi239
                }
            case -573:
                if true {
                    fmt.Println(v1, v4)
                } else {
                    switch ("hello" + "foo bar") {
                    case "":
                        v4 = (v4 >> ((4294967295) & 63))
                    case "a":
                        v3 = append(v3, ("\n\t" + ""))
                    case "hello":
                        v4 = v4
                    }
                    v1 = (v1 * (v1 / ((v1) | 1)))
                }
            default:
                var v243 float32 = (((0.1 + 0.0) * float32(0.1)) - (float32((-2.25 + (0.0 / 3.14159))) / (((0.0 + 1.5) + ((3.14159 + 1e10) - (float32((3.14159 + 0.0)) / 1.5))) - (0.0 - 0.1))))
                _ = v243
            }
            var v244 uint32 = 579
            _ = v244
        } else {
            fmt.Println(v1)
        }
    }
    for fi245 := 0; fi245 < 19; fi245++ {
        delete(v5, 0)
        switch "hello" {
        case "":
            v246 := GenS26{F0: 255, F1: (^(v4)), F2: 48}
            _ = v246
        case "a":
            v1 += (v1 / ((v1) | 1))
        }
        fmt.Println(v1, v4)
        _ = fi245
    }
    switch -677 {
    case -207:
        for fi247 := 0; fi247 < 15; fi247++ {
            var v248 int32 = -1
            var v249 uint8 = 93
            _ = fi247
            _ = v248
            _ = v249
        }
    case 128:
        v250 := GenS27{F0: 1, F1: -9223372036854775808, F2: ("" + "\n\t")}
        _ = v250
    }
    if (!true) {
        v5[-1] = 42
    } else {
        v251 := GenS28{F0: int(v1), F1: 4294967295}
        var v252 bool = false
        v253 := GenS15{F0: 1, F1: 1, F2: 239}
        _ = v251
        _ = v252
        _ = v253
    }
    if (((true || false) || (-128 >= 1)) && false) {
        if (-1 == 2147483647) {
            v4 = 91
            switch 551 {
            case 493:
                for fi254 := 0; fi254 < 1; fi254++ {
                    for fi255 := 0; fi255 < 6; fi255++ {
                        var v256 int32 = 1
                        _ = fi255
                        _ = v256
                    }
                    _ = fi254
                }
            case -2147483648:
                if false {
                    var v257 int = 0
                    v1 += (v1 | 4294967295)
                    _ = v257
                }
            default:
                v1 += 1
            }
        } else {
            var v258 uint = 2
            for fi259 := 0; fi259 < 11; fi259++ {
                for fi260 := 0; fi260 < 3; fi260++ {
                    var v261 int8 = 93
                    _ = fi260
                    _ = v261
                }
                v262 := GenS13{F0: 18446744073709551615, F1: (^(v1)), F2: 0}
                _ = fi259
                _ = v262
            }
            _ = v258
        }
        v1 += v1
        if false {
            if (1 <= 995) {
                for fi263 := 0; fi263 < 7; fi263++ {
                    var v264 uint8 = v4
                    var v265 bool = (true && true)
                    v266 := GenS17{F0: "\n\t", F1: fi263, F2: 133}
                    _ = fi263
                    _ = v264
                    _ = v265
                    _ = v266
                }
            } else {
                v267 := GenS12{F0: true, F1: 943}
                _ = v267
            }
            v268 := GenS29{F0: 65535, F1: 42, F2: -879}
            _ = v268
        }
    } else {
        var v269 uint64 = v1
        v270 := GenS30{F0: -39}
        _ = v269
        _ = v270
    }
    v1 = (v1 >> ((42) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
