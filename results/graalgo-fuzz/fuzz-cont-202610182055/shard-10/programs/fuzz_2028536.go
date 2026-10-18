package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int8; F2 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 uint; F1 uint16 }
type GenS6 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS7 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS8 struct { F0 uint64; F1 int8 }
type GenS9 struct { F0 int16; F1 int32; F2 int }
type GenS10 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS11 struct { F0 uint; F1 bool; F2 string }
type GenS12 struct { F0 uint32; F1 string; F2 bool }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int8; F1 int64; F2 uint }
type GenS15 struct { F0 int; F1 uint32; F2 string }
type GenS16 struct { F0 int; F1 uint16 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 int32; F1 uint16; F2 string }
type GenS19 struct { F0 int64 }
type GenS20 struct { F0 uint64; F1 int32 }
type GenS21 struct { F0 int16; F1 uint32 }
type GenS22 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS23 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS24 struct { F0 uint64 }
type GenS25 struct { F0 int }
type GenS26 struct { F0 int16 }
type GenS27 struct { F0 int16; F1 string }
type GenS28 struct { F0 int32; F1 uint; F2 bool }
type GenS29 struct { F0 uint32 }
type GenS30 struct { F0 uint8; F1 int32 }
type GenS31 struct { F0 int64; F1 uint32; F2 bool }

func main() {
    var v1 uint = 765
    var v2 int64 = -1
    v3 := []uint32{743, 590, 2147483648}
    var v4 uint = 1
    var v5 float32 = ((((-2.25 * float32((1e10 / 3.14159))) * -2.25) + 0.0) - 1e10)
    if true {
        var v6 float64 = float64(v5)
        _ = v6
    } else {
        if (!false) {
            fmt.Println(v5)
            var v7 uint = 42
            if (32767 > 255) {
                v2 = (v2 ^ (v2 >> (((v4 - v4)) & 63)))
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
                if false {
                    if false {
                        for fi8 := 0; fi8 < 6; fi8++ {
                            v9 := GenS1{F0: 65535}
                            _ = fi8
                            _ = v9
                        }
                        v7 -= v7
                        if (!(!(1 != -917))) {
                            v1 = (v7 * (v7 * (v4 << ((680) & 63))))
                            var v10 uint16 = 65535
                            _ = v10
                        } else {
                            _ = len(v3)
                            v11 := GenS2{F0: 379, F1: -128, F2: -1}
                            v5 = (v5 * (v5 - 1e10))
                            _ = v11
                        }
                    }
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
                if false {
                    v12 := GenS1{F0: 1}
                    _ = v12
                } else {
                    var v13 int64 = (-(v2))
                    var v14 int8 = 42
                    v13 -= -54
                    _ = v13
                    _ = v14
                }
                var v15 uint = v7
                _ = v15
            }
            _ = v7
        } else {
            v16 := GenS3{F0: -1}
            var v17 uint16 = 65535
            var v18 int32 = -1
            _ = v16
            _ = v17
            _ = v18
        }
        switch "\n\t" {
        case "":
            for fi19 := 0; fi19 < 7; fi19++ {
                fmt.Println(v1)
                _ = fi19
            }
        case "hello":
            var v20 int = 815
            _ = v20
        default:
            v3 = append(v3, 2147483648)
        }
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 float32 = 0.0
            if (fi21 <= 9223372036854775807) {
                switch (fi21 & fi21) {
                case 127:
                    if true {
                        fmt.Println(v1, v4)
                    } else {
                        fmt.Println(v5, v5, v22)
                        var v23 string = (("a" + "\n\t") + ("a" + (("a" + "foo bar") + "\n\t")))
                        _ = v23
                    }
                case -1:
                    var v24 int8 = -1
                    _ = v24
                }
                v1 = (v4 | v1)
            } else {
                var v25 uint32 = 1
                var v26 uint8 = 255
                fmt.Println(v22, v22, v2)
                _ = v25
                _ = v26
            }
            v27 := GenS4{F0: (("foo bar" + "a") + ""), F1: 0}
            _ = fi21
            _ = v22
            _ = v27
        }
    }
    if (!true) {
        switch "foo bar" {
        case "\n\t":
            var v28 uint16 = 694
            _ = v28
        case "hello":
            if (42 < 1) {
                v29 := GenS5{F0: (v1 / (((v1 ^ 232)) | 1)), F1: 1}
                var v30 string = ("" + "")
                if true {
                    v31 := GenS5{F0: v4, F1: 42}
                    _ = v31
                } else {
                    v32 := GenS6{F0: 515, F1: 1, F2: 4294967295}
                    if (!true) {
                        v5 += (v5 / v5)
                        fmt.Println(v4)
                    } else {
                        v2 -= v2
                        v33 := GenS7{F0: 4294967295, F1: 128, F2: -2147483648}
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v33
                    }
                    v30 = v30
                    _ = v32
                }
                _ = v29
                _ = v30
            } else {
                var v34 int = 0
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = len(v3)
                _ = v34
            }
        default:
            v5 = v5
        }
        var v35 float64 = (1e10 - -2.25)
        _ = v35
    } else {
        v36 := GenS3{F0: -197}
        v2 -= (v2 / (((-(v2))) | 1))
        switch (("foo bar" + (("hello" + ("foo bar" + "")) + (("a" + "a") + "\n\t"))) + "foo bar") {
        case "hello":
            if false {
                v37 := GenS4{F0: "hello", F1: 1000}
                if len(v3) > 0 {
                    v3[0] = 963
                }
                if false {
                    switch -302 {
                    case -1:
                        v2 -= v2
                    case -876:
                        var v38 string = ("" + ("" + ("hello" + "\n\t")))
                        _ = v38
                    case -499:
                        var v39 bool = true
                        _ = v39
                    default:
                        fmt.Println(v5, v2, v4)
                    }
                    switch 32768 {
                    case 32768:
                        switch ("\n\t" + "foo bar") {
                        case "foo bar":
                            v2 += 1
                        case "hello":
                            switch "\n\t" {
                            case "foo bar":
                                switch 127 {
                                case 32768:
                                    v1 -= (v4 / (((^(v4))) | 1))
                                case 0:
                                    fmt.Println(v2, v1)
                                case 496:
                                    var v40 int8 = 0
                                    _ = v40
                                default:
                                    _ = len(v3)
                                }
                            case "hello":
                                for fi41 := 0; fi41 < 11; fi41++ {
                                    switch "foo bar" {
                                    case "\n\t":
                                        fmt.Println(v1, v2, v5)
                                    case "foo bar":
                                        for fi42 := 0; fi42 < 12; fi42++ {
                                            var v43 float32 = 1.5
                                            _ = fi42
                                            _ = v43
                                        }
                                    case "a":
                                        var v44 bool = (0.1 < v5)
                                        _ = v44
                                    default:
                                        v45 := GenS6{F0: 42, F1: 1, F2: 1}
                                        _ = v45
                                    }
                                    var v46 uint32 = 688
                                    _ = fi41
                                    _ = v46
                                }
                            case "":
                                switch 0 {
                                case 4294967295:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case -2147483648:
                                    var v47 bool = true
                                    _ = v47
                                case 383:
                                    v48 := GenS7{F0: 42, F1: 128, F2: 42}
                                    _ = v48
                                default:
                                    if ((!(false || (true && (!true)))) && false) {
                                        if (true || false) {
                                            var v49 uint = (^(v4))
                                            fmt.Println(v2, v2)
                                            v50 := GenS1{F0: 32767}
                                            _ = v49
                                            _ = v50
                                        } else {
                                            switch 2147483647 {
                                            case 447:
                                                v4 -= v1
                                            case 2147483647:
                                                v5 = (v5 * 0.0)
                                            case 0:
                                                for fi51 := 0; fi51 < 8; fi51++ {
                                                    var v52 uint64 = 922
                                                    switch "" {
                                                    case "\n\t":
                                                        fmt.Println(v5, v52)
                                                    case "hello":
                                                        for fi53 := 0; fi53 < 7; fi53++ {
                                                            if true {
                                                                if (v2 <= -1) {
                                                                    switch ("a" + (("" + "hello") + ("\n\t" + ("foo bar" + ((("foo bar" + "hello") + (("\n\t" + "\n\t") + "foo bar")) + "a"))))) {
                                                                    case "\n\t":
                                                                        var v54 int = 0
                                                                        fmt.Println(v4, v52)
                                                                        switch ("a" + "") {
                                                                        case "":
                                                                            v55 := GenS5{F0: v4, F1: 42}
                                                                            v2 -= (v2 & v2)
                                                                            fmt.Println(v5)
                                                                            _ = v55
                                                                        case "foo bar":
                                                                            var v56 uint64 = v52
                                                                            _ = v56
                                                                        case "\n\t":
                                                                            if (!((!false) && (!(!(false && true))))) {
                                                                                var v57 uint8 = 1
                                                                                v58 := GenS8{F0: (v52 ^ (v52 | (v52 >> (((v1 & (^(v4)))) & 63)))), F1: 90}
                                                                                _ = v57
                                                                                _ = v58
                                                                            } else {
                                                                                v52 -= 264
                                                                                v59 := GenS4{F0: "", F1: 762}
                                                                                _ = v59
                                                                            }
                                                                        }
                                                                        _ = v54
                                                                    case "foo bar":
                                                                        var v60 uint32 = 65535
                                                                        _ = v60
                                                                    }
                                                                    v5 = v5
                                                                    v61 := GenS7{F0: 42, F1: 8, F2: 2147483647}
                                                                    _ = v61
                                                                }
                                                                var v62 int32 = -2147483648
                                                                _ = len(v3)
                                                                _ = v62
                                                            } else {
                                                                var v63 uint16 = 255
                                                                fmt.Println(fi53, v2, fi51)
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                _ = v63
                                                            }
                                                            if true {
                                                                v2 = (v2 / (((v2 + (-(v2)))) | 1))
                                                                v1 = 0
                                                            } else {
                                                                switch "hello" {
                                                                case "hello":
                                                                    var v64 uint16 = 255
                                                                    _ = v64
                                                                case "a":
                                                                    if false {
                                                                        var v65 int = fi51
                                                                        _ = len(v3)
                                                                        v66 := GenS9{F0: -1, F1: int32(v2), F2: (^(fi53))}
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            if true {
                                                                                if (!(false || true)) {
                                                                                    if false {
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                switch "" {
                                                                                case "hello":
                                                                                    v65 = (v65 % ((int(v52)) | 1))
                                                                                case "foo bar":
                                                                                    switch "a" {
                                                                                    case "hello":
                                                                                        v67 := GenS4{F0: "hello", F1: 0}
                                                                                        _ = v67
                                                                                    case "\n\t":
                                                                                        v1 = (v4 & v1)
                                                                                    case "foo bar":
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                    }
                                                                                case "a":
                                                                                    v52 += v52
                                                                                default:
                                                                                    var v68 uint8 = 1
                                                                                    _ = v68
                                                                                }
                                                                                if false {
                                                                                    var v69 uint8 = 220
                                                                                    if (-2.25 < (3.14159 - v5)) {
                                                                                        if (v69 <= v69) {
                                                                                            switch "" {
                                                                                            case "hello":
                                                                                                if (uint(v52) < 0) {
                                                                                                    switch (fi53 >> (((v4 - (v4 - (v1 >> ((v4) & 63))))) & 63)) {
                                                                                                    case 127:
                                                                                                        switch "a" {
                                                                                                        case "a":
                                                                                                            v3 = append(v3, 1)
                                                                                                        case "foo bar":
                                                                                                            if len(v3) > 0 {
                                                                                                                v3[0] = 2147483648
                                                                                                            }
                                                                                                        case "hello":
                                                                                                            v1 = v1
                                                                                                        }
                                                                                                    case 1:
                                                                                                        switch "a" {
                                                                                                        case "\n\t":
                                                                                                            _ = len(v3)
                                                                                                        case "a":
                                                                                                            fmt.Println(fi51)
                                                                                                        }
                                                                                                        var v70 bool = (!true)
                                                                                                        _ = v70
                                                                                                    case -715:
                                                                                                        v2 += int64(fi51)
                                                                                                    default:
                                                                                                        v52 = 42
                                                                                                    }
                                                                                                    var v71 uint = (v4 * v1)
                                                                                                    _ = v71
                                                                                                } else {
                                                                                                    fmt.Println(v5)
                                                                                                    v4 = (v4 - 65535)
                                                                                                    var v72 int16 = 0
                                                                                                    _ = v72
                                                                                                }
                                                                                            case "foo bar":
                                                                                                fmt.Println(v65, v65)
                                                                                            case "\n\t":
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = 4294967295
                                                                                                }
                                                                                            default:
                                                                                                v69 = 2
                                                                                            }
                                                                                        } else {
                                                                                            if (v52 == v52) {
                                                                                                if (!false) {
                                                                                                    v1 -= v4
                                                                                                    switch fi53 {
                                                                                                    case 42:
                                                                                                        switch fi53 {
                                                                                                        case 206:
                                                                                                            v73 := GenS1{F0: 65535}
                                                                                                            _ = v73
                                                                                                        case 809:
                                                                                                            if false {
                                                                                                                v1 -= 65535
                                                                                                            }
                                                                                                        default:
                                                                                                            switch "a" {
                                                                                                            case "hello":
                                                                                                                var v74 int64 = (-(v2))
                                                                                                                _ = v74
                                                                                                            case "":
                                                                                                                var v75 int8 = 0
                                                                                                                _ = v75
                                                                                                            }
                                                                                                        }
                                                                                                    case 128:
                                                                                                        var v76 float32 = ((v5 - float32(1e10)) - v5)
                                                                                                        _ = v76
                                                                                                    default:
                                                                                                        var v77 float32 = (((float32(1.5) + (3.14159 + v5)) + v5) * v5)
                                                                                                        _ = v77
                                                                                                    }
                                                                                                }
                                                                                                switch fi53 {
                                                                                                case 127:
                                                                                                    var v78 float64 = 1e10
                                                                                                    _ = v78
                                                                                                case 683:
                                                                                                    switch "\n\t" {
                                                                                                    case "hello":
                                                                                                        fmt.Println(v5, v2, fi51)
                                                                                                    case "foo bar":
                                                                                                        v79 := GenS6{F0: 18446744073709551615, F1: (^(v52)), F2: v52}
                                                                                                        _ = v79
                                                                                                    case "a":
                                                                                                        var v80 uint8 = v69
                                                                                                        _ = v80
                                                                                                    default:
                                                                                                        var v81 uint64 = v52
                                                                                                        _ = v81
                                                                                                    }
                                                                                                default:
                                                                                                    fmt.Println(v2, v4)
                                                                                                }
                                                                                            } else {
                                                                                                v82 := GenS10{F0: (v69 << ((v4) & 63)), F1: 65535, F2: 0}
                                                                                                var v83 string = ("\n\t" + "\n\t")
                                                                                                _ = v82
                                                                                                _ = v83
                                                                                            }
                                                                                            v84 := GenS8{F0: v52, F1: -34}
                                                                                            v85 := GenS11{F0: v1, F1: (v1 >= v4), F2: (((((("a" + "\n\t") + "hello") + "foo bar") + "") + "") + "hello")}
                                                                                            _ = v84
                                                                                            _ = v85
                                                                                        }
                                                                                        fmt.Println(v69, fi53)
                                                                                        if true {
                                                                                            v65 += 9223372036854775807
                                                                                            var v86 int16 = -1
                                                                                            var v87 uint = 4294967295
                                                                                            _ = v86
                                                                                            _ = v87
                                                                                        }
                                                                                    }
                                                                                    v88 := GenS2{F0: 65535, F1: 100, F2: 32767}
                                                                                    _ = v69
                                                                                    _ = v88
                                                                                } else {
                                                                                    if false {
                                                                                        v1 = (v4 & (v4 >> ((v4) & 63)))
                                                                                        switch v65 {
                                                                                        case 128:
                                                                                            v89 := GenS12{F0: 777, F1: ("foo bar" + "foo bar"), F2: ((v5 - (v5 / v5)) >= (v5 / v5))}
                                                                                            _ = v89
                                                                                        case 32767:
                                                                                            if false {
                                                                                                _ = len(v3)
                                                                                                var v90 uint = v4
                                                                                                if len(v3) >= 2 {
                                                                                                    v3 = v3[1:len(v3)]
                                                                                                }
                                                                                                _ = v90
                                                                                            } else {
                                                                                                v65 = fi53
                                                                                                fmt.Println(fi53, v52)
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        v52 = 537
                                                                                        if false {
                                                                                            if true {
                                                                                                if len(v3) >= 2 {
                                                                                                    v3 = v3[1:len(v3)]
                                                                                                }
                                                                                                v5 = -2.25
                                                                                                var v91 float32 = v5
                                                                                                _ = v91
                                                                                            } else {
                                                                                                if true {
                                                                                                    var v92 uint32 = 129
                                                                                                    v65 -= (v65 % (((fi53 & (v65 | (v65 | fi51)))) | 1))
                                                                                                    switch "a" {
                                                                                                    case "":
                                                                                                        v2 = -887
                                                                                                    case "foo bar":
                                                                                                        switch 4294967295 {
                                                                                                        case -9223372036854775808:
                                                                                                            v4 = (v4 % ((v4) | 1))
                                                                                                            v92 -= (^(v92))
                                                                                                            switch -2147483648 {
                                                                                                            case 9223372036854775807:
                                                                                                                var v93 uint8 = 127
                                                                                                                _ = v93
                                                                                                            case 2147483647:
                                                                                                                fmt.Println(v92, v2, fi51)
                                                                                                            case -729:
                                                                                                                if false {
                                                                                                                    v94 := GenS13{F0: 0}
                                                                                                                    v3 = append(v3, (v92 << (((v1 >> (((^(v1))) & 63))) & 63)))
                                                                                                                    _ = v94
                                                                                                                }
                                                                                                                fmt.Println(fi51)
                                                                                                            default:
                                                                                                                v95 := GenS14{F0: -1, F1: v2, F2: v4}
                                                                                                                _ = v95
                                                                                                            }
                                                                                                        default:
                                                                                                            var v96 int64 = (v2 + (v2 - 0))
                                                                                                            _ = v96
                                                                                                        }
                                                                                                    default:
                                                                                                        if len(v3) >= 2 {
                                                                                                            v3 = v3[1:len(v3)]
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v92
                                                                                                } else {
                                                                                                    _ = len(v3)
                                                                                                    switch "" {
                                                                                                    case "\n\t":
                                                                                                        v52 = 42
                                                                                                    default:
                                                                                                        fmt.Println(v2, fi53)
                                                                                                    }
                                                                                                    _ = len(v3)
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            if true {
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = 1
                                                                                                }
                                                                                                v2 -= v2
                                                                                                fmt.Println(fi53, v65, fi51)
                                                                                                var v97 uint8 = 175
                                                                                                v2 -= (v2 - v2)
                                                                                                _ = v97
                                                                                            } else {
                                                                                                var v98 bool = (128 == 0)
                                                                                                _ = v98
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v99 := GenS14{F0: -112, F1: 483, F2: (v4 / (((v4 - 61)) | 1))}
                                                                                    v2 = (v2 ^ -636)
                                                                                    _ = v99
                                                                                }
                                                                                if (!true) {
                                                                                    switch ("a" + "hello") {
                                                                                    case "hello":
                                                                                        _ = len(v3)
                                                                                    case "a":
                                                                                        switch -192 {
                                                                                        case 42:
                                                                                            var v100 bool = false
                                                                                            _ = v100
                                                                                        default:
                                                                                            v101 := GenS15{F0: v65, F1: 42, F2: ("\n\t" + "\n\t")}
                                                                                            _ = v101
                                                                                        }
                                                                                    }
                                                                                    if (2147483647 <= -2147483648) {
                                                                                        v102 := GenS16{F0: (v65 >> ((561) & 63)), F1: 549}
                                                                                        v52 = 0
                                                                                        _ = v102
                                                                                    }
                                                                                    var v103 bool = (767 != 392)
                                                                                    v103 = true
                                                                                    _ = v103
                                                                                } else {
                                                                                    v65 += 2147483647
                                                                                    if (0 == 32767) {
                                                                                        switch (^(fi51)) {
                                                                                        case 849:
                                                                                            v65 = fi51
                                                                                            if false {
                                                                                                v1 -= (v4 + 4294967295)
                                                                                                switch 0 {
                                                                                                case -1:
                                                                                                    if false {
                                                                                                        switch fi53 {
                                                                                                        case -639:
                                                                                                            v104 := GenS14{F0: -1, F1: v2, F2: 595}
                                                                                                            _ = v104
                                                                                                        default:
                                                                                                            var v105 uint = 170
                                                                                                            _ = v105
                                                                                                        }
                                                                                                    } else {
                                                                                                        if ((v65 >> ((42) & 63)) >= (v65 * fi51)) {
                                                                                                            v106 := GenS5{F0: 4294967295, F1: 389}
                                                                                                            _ = len(v3)
                                                                                                            v107 := GenS17{F0: 0}
                                                                                                            switch -462 {
                                                                                                            case 812:
                                                                                                                if true {
                                                                                                                    v4 = (v4 >> ((v1) & 63))
                                                                                                                    v108 := GenS5{F0: (v1 & v4), F1: 32767}
                                                                                                                    _ = v108
                                                                                                                } else {
                                                                                                                    fmt.Println(v4)
                                                                                                                }
                                                                                                            case 42:
                                                                                                                if len(v3) >= 2 {
                                                                                                                    v3 = v3[1:len(v3)]
                                                                                                                }
                                                                                                            case 585:
                                                                                                                v109 := GenS8{F0: 119, F1: -106}
                                                                                                                _ = v109
                                                                                                            default:
                                                                                                                if false {
                                                                                                                    switch "" {
                                                                                                                    case "foo bar":
                                                                                                                        v2 -= v2
                                                                                                                    case "\n\t":
                                                                                                                        if len(v3) > 0 {
                                                                                                                            v3[0] = 365
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v110 int = (fi51 % ((v65) | 1))
                                                                                                                    switch ("" + "a") {
                                                                                                                    case "foo bar":
                                                                                                                        v65 -= v65
                                                                                                                    case "\n\t":
                                                                                                                        v1 = v4
                                                                                                                    }
                                                                                                                    var v111 uint = (v1 << (((v1 - 487)) & 63))
                                                                                                                    var v112 uint64 = v52
                                                                                                                    _ = v110
                                                                                                                    _ = v111
                                                                                                                    _ = v112
                                                                                                                }
                                                                                                            }
                                                                                                            var v113 int8 = 42
                                                                                                            _ = v106
                                                                                                            _ = v107
                                                                                                            _ = v113
                                                                                                        }
                                                                                                        _ = len(v3)
                                                                                                    }
                                                                                                case 32768:
                                                                                                    switch "a" {
                                                                                                    case "hello":
                                                                                                        v52 -= (v52 * v52)
                                                                                                    case "a":
                                                                                                        _ = len(v3)
                                                                                                    case "":
                                                                                                        var v114 bool = true
                                                                                                        _ = v114
                                                                                                    }
                                                                                                default:
                                                                                                    v115 := GenS18{F0: 637, F1: 42, F2: ""}
                                                                                                    _ = v115
                                                                                                }
                                                                                            }
                                                                                            v116 := GenS19{F0: (^(v2))}
                                                                                            _ = v116
                                                                                        case 0:
                                                                                            fmt.Println(v65, v1)
                                                                                        default:
                                                                                            v65 -= (fi53 | (fi51 >> (((v1 * (^(v4)))) & 63)))
                                                                                        }
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = 0
                                                                                        }
                                                                                    } else {
                                                                                        if false {
                                                                                            v65 = fi51
                                                                                            switch (fi53 * v65) {
                                                                                            case 32767:
                                                                                                v4 += (v1 | (v1 % (((v1 ^ (^(v1)))) | 1)))
                                                                                            case 461:
                                                                                                if len(v3) >= 2 {
                                                                                                    v3 = v3[1:len(v3)]
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            v4 += v4
                                                                                            if (!(4294967295 != 0)) {
                                                                                                var v117 uint = 4294967295
                                                                                                v2 += v2
                                                                                                v118 := GenS14{F0: 100, F1: v2, F2: 4294967295}
                                                                                                fmt.Println(fi51)
                                                                                                v4 += (v117 & v1)
                                                                                                if (-32 <= 55) {
                                                                                                    switch "foo bar" {
                                                                                                    case "":
                                                                                                        var v119 int8 = 20
                                                                                                        _ = v119
                                                                                                    case "hello":
                                                                                                        v2 -= int64(v117)
                                                                                                    case "a":
                                                                                                        if (!true) {
                                                                                                            var v120 int = (fi53 / (((fi51 / ((fi51) | 1))) | 1))
                                                                                                            if false {
                                                                                                                fmt.Println(v117, v52, v1)
                                                                                                            }
                                                                                                            v2 += 9223372036854775807
                                                                                                            _ = v120
                                                                                                        } else {
                                                                                                            if len(v3) >= 2 {
                                                                                                                v3 = v3[1:len(v3)]
                                                                                                            }
                                                                                                            if (v52 != v52) {
                                                                                                                switch "" {
                                                                                                                case "foo bar":
                                                                                                                    switch (v65 | 127) {
                                                                                                                    case 1:
                                                                                                                        v121 := GenS9{F0: 0, F1: 1, F2: (fi53 - fi53)}
                                                                                                                        _ = v121
                                                                                                                    case 783:
                                                                                                                        if (255 == 42) {
                                                                                                                            switch "foo bar" {
                                                                                                                            case "a":
                                                                                                                                var v122 int8 = 1
                                                                                                                                _ = v122
                                                                                                                            case "":
                                                                                                                                fmt.Println(fi51, v65)
                                                                                                                            default:
                                                                                                                                if ((v1 ^ (v117 * 4294967295)) < v117) {
                                                                                                                                    if (!((^(v2)) != v2)) {
                                                                                                                                        var v123 string = "a"
                                                                                                                                        switch 0 {
                                                                                                                                        case 0:
                                                                                                                                            if len(v3) >= 2 {
                                                                                                                                                v3 = v3[1:len(v3)]
                                                                                                                                            }
                                                                                                                                            if false {
                                                                                                                                                if len(v3) >= 2 {
                                                                                                                                                    v3 = v3[1:len(v3)]
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            var v124 int64 = -9223372036854775808
                                                                                                                                            _ = v124
                                                                                                                                        }
                                                                                                                                        if len(v3) > 0 {
                                                                                                                                            v3[0] = 65535
                                                                                                                                        }
                                                                                                                                        _ = v123
                                                                                                                                    } else {
                                                                                                                                        var v125 int8 = 127
                                                                                                                                        v117 += v1
                                                                                                                                        _ = v125
                                                                                                                                    }
                                                                                                                                    v4 = v1
                                                                                                                                    var v126 uint64 = 179
                                                                                                                                    _ = v126
                                                                                                                                } else {
                                                                                                                                    var v127 uint8 = 0
                                                                                                                                    _ = v127
                                                                                                                                }
                                                                                                                            }
                                                                                                                            _ = len(v3)
                                                                                                                            v65 = v65
                                                                                                                        }
                                                                                                                    }
                                                                                                                default:
                                                                                                                    v52 -= (^(v52))
                                                                                                                }
                                                                                                                v128 := GenS20{F0: v52, F1: -191}
                                                                                                                v3 = append(v3, 2147483648)
                                                                                                                v129 := GenS21{F0: 32767, F1: 715}
                                                                                                                _ = v128
                                                                                                                _ = v129
                                                                                                            }
                                                                                                        }
                                                                                                    default:
                                                                                                        v130 := GenS22{F0: (v52 / ((uint64(v117)) | 1)), F1: true, F2: 2}
                                                                                                        _ = v130
                                                                                                    }
                                                                                                    switch (^(fi53)) {
                                                                                                    case -2147483648:
                                                                                                        if true {
                                                                                                            v131 := GenS5{F0: (^(v117)), F1: 100}
                                                                                                            if (false && (1.5 < float64(((v5 - float32((float64((1e10 - (1.5 - v5))) * 1e10))) + ((v5 - (v5 - ((v5 - v5) / 0.1))) + v5))))) {
                                                                                                                v4 -= 1
                                                                                                                v132 := GenS3{F0: 110}
                                                                                                                v133 := GenS23{F0: 65535, F1: v1, F2: 4294967295}
                                                                                                                _ = v132
                                                                                                                _ = v133
                                                                                                            }
                                                                                                            _ = v131
                                                                                                        }
                                                                                                    }
                                                                                                    var v134 string = "\n\t"
                                                                                                    _ = len(v3)
                                                                                                    _ = v134
                                                                                                }
                                                                                                if (true || (-32768 == 766)) {
                                                                                                    var v135 uint8 = 152
                                                                                                    var v136 string = ""
                                                                                                    if ((v52 ^ (^(v52))) == (^(v52))) {
                                                                                                        var v137 string = ((v136 + (v136 + v136)) + "foo bar")
                                                                                                        switch 127 {
                                                                                                        case 1:
                                                                                                            v4 -= v117
                                                                                                        default:
                                                                                                            fmt.Println(v2, fi51)
                                                                                                            var v138 float64 = 1.5
                                                                                                            var v139 int32 = 42
                                                                                                            _ = v138
                                                                                                            _ = v139
                                                                                                        }
                                                                                                        _ = v137
                                                                                                    }
                                                                                                    _ = v135
                                                                                                    _ = v136
                                                                                                } else {
                                                                                                    var v140 uint64 = (v52 + v52)
                                                                                                    _ = v140
                                                                                                }
                                                                                                var v141 float64 = 3.14159
                                                                                                v2 = (-(v2))
                                                                                                _ = len(v3)
                                                                                                _ = v117
                                                                                                _ = v118
                                                                                                _ = v141
                                                                                            } else {
                                                                                                v142 := GenS24{F0: (v52 >> ((v4) & 63))}
                                                                                                _ = v142
                                                                                            }
                                                                                            if ((v4 % (((v1 + (^(v4)))) | 1)) >= v1) {
                                                                                                if false {
                                                                                                    v143 := GenS18{F0: -2147483648, F1: 22, F2: "foo bar"}
                                                                                                    v2 = 187
                                                                                                    v144 := GenS4{F0: "foo bar", F1: 639}
                                                                                                    _ = v143
                                                                                                    _ = v144
                                                                                                }
                                                                                                v145 := GenS1{F0: 42}
                                                                                                if (!(true && false)) {
                                                                                                    if false {
                                                                                                        fmt.Println(fi53)
                                                                                                        fmt.Println(v1)
                                                                                                    }
                                                                                                    _ = len(v3)
                                                                                                    switch "" {
                                                                                                    case "foo bar":
                                                                                                        v5 += 3.14159
                                                                                                    case "\n\t":
                                                                                                        if len(v3) >= 2 {
                                                                                                            v3 = v3[1:len(v3)]
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                _ = v145
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    switch -361 {
                                                                                    case 0:
                                                                                        var v146 uint8 = 2
                                                                                        _ = v146
                                                                                    }
                                                                                }
                                                                            }
                                                                        default:
                                                                            fmt.Println(v1, v5)
                                                                        }
                                                                        switch ("foo bar" + ("hello" + "")) {
                                                                        case "":
                                                                            v4 += v1
                                                                        case "foo bar":
                                                                            v5 += 0.0
                                                                        case "\n\t":
                                                                            v5 += (float32(-2.25) * (v5 + v5))
                                                                        default:
                                                                            if (1 == 255) {
                                                                                fmt.Println(fi51)
                                                                            } else {
                                                                                var v147 uint8 = 42
                                                                                fmt.Println(v65, v4)
                                                                                _ = v147
                                                                            }
                                                                        }
                                                                        _ = v65
                                                                        _ = v66
                                                                    }
                                                                }
                                                            }
                                                            fmt.Println(v4, v2, fi53)
                                                            _ = fi53
                                                        }
                                                    }
                                                    switch fi51 {
                                                    case 0:
                                                        for fi148 := 0; fi148 < 18; fi148++ {
                                                            var v149 uint16 = uint16(v52)
                                                            if len(v3) > 0 {
                                                                v3[0] = 501
                                                            }
                                                            _ = fi148
                                                            _ = v149
                                                        }
                                                    case 337:
                                                        var v150 int16 = -954
                                                        _ = v150
                                                    }
                                                    _ = fi51
                                                    _ = v52
                                                }
                                            default:
                                                switch 32768 {
                                                case 127:
                                                    var v151 float32 = (-2.25 + (v5 - v5))
                                                    _ = v151
                                                case 701:
                                                    var v152 uint32 = 4294967295
                                                    _ = v152
                                                case 0:
                                                    _ = len(v3)
                                                }
                                            }
                                        }
                                        v153 := GenS25{F0: 2147483647}
                                        v154 := GenS26{F0: -107}
                                        _ = v153
                                        _ = v154
                                    }
                                }
                            }
                        case "\n\t":
                            for fi155 := 0; fi155 < 10; fi155++ {
                                switch fi155 {
                                case 127:
                                    var v156 float32 = (0.0 * 1e10)
                                    _ = v156
                                case 9223372036854775807:
                                    switch fi155 {
                                    case 0:
                                        var v157 int32 = 2147483647
                                        _ = v157
                                    case -857:
                                        if (false || true) {
                                            var v158 int64 = -1
                                            _ = v158
                                        }
                                    case 52:
                                        v4 = (^(v1))
                                    }
                                }
                                var v159 string = ""
                                var v160 float32 = 1e10
                                _ = fi155
                                _ = v159
                                _ = v160
                            }
                        default:
                            var v161 float64 = 0.1
                            _ = v161
                        }
                    default:
                        switch "foo bar" {
                        case "foo bar":
                            v162 := GenS27{F0: 244, F1: "hello"}
                            _ = v162
                        case "":
                            fmt.Println(v2, v2, v2)
                        }
                    }
                    var v163 uint32 = 676
                    _ = v163
                } else {
                    switch 32767 {
                    case 32767:
                        v164 := GenS26{F0: -435}
                        _ = v164
                    default:
                        switch -1 {
                        case 0:
                            v5 = ((0.1 - 0.0) + v5)
                        default:
                            if true {
                                v165 := GenS28{F0: 0, F1: (^(v1)), F2: true}
                                _ = v165
                            } else {
                                for fi166 := 0; fi166 < 18; fi166++ {
                                    var v167 int64 = v2
                                    _ = fi166
                                    _ = v167
                                }
                                v3 = append(v3, 42)
                                for fi168 := 0; fi168 < 15; fi168++ {
                                    v169 := GenS22{F0: 18446744073709551615, F1: (664 >= 127), F2: 1}
                                    for fi170 := 0; fi170 < 2; fi170++ {
                                        if (!(!(851 > 480))) {
                                            v171 := GenS19{F0: (^(v2))}
                                            var v172 uint32 = 803
                                            _ = v171
                                            _ = v172
                                        }
                                        var v173 float64 = 1e10
                                        _ = fi170
                                        _ = v173
                                    }
                                    v2 -= v2
                                    _ = fi168
                                    _ = v169
                                }
                            }
                        }
                    }
                    fmt.Println(v2, v5)
                }
                _ = v37
            } else {
                fmt.Println(v4, v4, v4)
            }
        case "a":
            v174 := GenS29{F0: 0}
            _ = v174
        case "foo bar":
            var v175 uint32 = 1
            _ = v175
        }
        _ = v36
    }
    for fi176 := 0; fi176 < 12; fi176++ {
        v1 = v1
        _ = fi176
    }
    var v177 int64 = 42
    for fi178 := 0; fi178 < 4; fi178++ {
        if len(v3) > 0 {
            v3[0] = 42
        }
        _ = fi178
    }
    v179 := GenS4{F0: "a", F1: 32767}
    fmt.Println(v5)
    v2 -= (v177 - 9223372036854775807)
    switch "" {
    case "a":
        switch -9223372036854775808 {
        case 931:
            v180 := GenS25{F0: -2147483648}
            _ = v180
        default:
            fmt.Println(v2)
        }
    default:
        v3 = append(v3, 990)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch 0 {
    case 9223372036854775807:
        var v181 int8 = 42
        _ = v181
    }
    v182 := GenS30{F0: 36, F1: 2147483647}
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    var v183 int16 = -649
    v184 := GenS31{F0: (v2 | 0), F1: 707, F2: (255 >= 65535)}
    v185 := GenS17{F0: 42}
    v4 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v177:", v177)
    fmt.Println("v183:", v183)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v177
    _ = v179
    _ = v182
    _ = v183
    _ = v184
    _ = v185
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
