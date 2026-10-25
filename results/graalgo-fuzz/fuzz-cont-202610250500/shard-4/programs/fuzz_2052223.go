package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint8; F2 int64 }
type GenS4 struct { F0 int8; F1 uint; F2 int16 }
type GenS5 struct { F0 bool; F1 int; F2 bool }
type GenS6 struct { F0 uint; F1 bool; F2 int64 }
type GenS7 struct { F0 int; F1 uint16; F2 uint32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint32; F1 uint }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int16; F1 bool; F2 string }
type GenS12 struct { F0 bool; F1 string; F2 uint32 }
type GenS13 struct { F0 uint32; F1 uint64 }
type GenS14 struct { F0 string; F1 int8 }
type GenS15 struct { F0 uint; F1 int16; F2 int64 }
type GenS16 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS17 struct { F0 int; F1 int16; F2 int8 }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 int8; F1 int32; F2 int32 }
type GenS20 struct { F0 int16; F1 uint16; F2 string }
type GenS21 struct { F0 int8; F1 int16 }
type GenS22 struct { F0 int8; F1 bool; F2 uint8 }
type GenS23 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS24 struct { F0 int64 }
type GenS25 struct { F0 int }
type GenS26 struct { F0 uint64; F1 string; F2 string }
type GenS27 struct { F0 uint16; F1 int; F2 int16 }
type GenS28 struct { F0 uint32; F1 bool; F2 int8 }
type GenS29 struct { F0 int64 }
type GenS30 struct { F0 int16; F1 int32; F2 bool }
type GenS31 struct { F0 bool; F1 bool; F2 uint8 }

func main() {
    var v1 float64 = 0.0
    var v2 uint64 = 2
    var v3 uint = 577
    if ((((v1 - v1) <= v1) && (0 > 330)) && (73 == 42)) {
        v4 := []int{-9223372036854775808, 0, 32768}
        _ = v4
    }
    v5 := []int{128}
    fmt.Println(v2)
    v6 := GenS1{F0: v3, F1: 181, F2: 4294967295}
    fmt.Println(v3, v2, v3)
    v7 := GenS2{F0: 32767}
    v2 = 42
    if (true && (true && false)) {
        for fi8 := 0; fi8 < 19; fi8++ {
            v1 = float64((float32(v1) - 0.1))
            _ = fi8
        }
        var v9 float64 = float64(0.1)
        v1 += ((v1 + v9) / v1)
        _ = v9
    } else {
        for fi10 := 0; fi10 < 16; fi10++ {
            v11 := GenS2{F0: 262}
            v12 := GenS2{F0: 42}
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 int64 = 414
                if true {
                    if ((-2.25 - 1.5) <= 1.5) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v3 += uint(v2)
                        v14 = v14
                    }
                    _ = len(v5)
                    if (1 > 1) {
                        if len(v5) > 0 {
                            v5[0] = fi10
                        }
                        if len(v5) > 0 {
                            v5[0] = fi10
                        }
                    } else {
                        if false {
                            v15 := GenS3{F0: -1, F1: 127, F2: (v14 / ((-1) | 1))}
                            var v16 uint = (v3 & (v3 << ((255) & 63)))
                            _ = v15
                            _ = v16
                        }
                    }
                    var v17 int = (fi13 * 42)
                    v18 := GenS1{F0: (^(v3)), F1: 0, F2: 1}
                    _ = v17
                    _ = v18
                } else {
                    v1 = (1e10 * ((v1 + 0.1) * 3.14159))
                }
                switch (fi10 * 32768) {
                case 9223372036854775807:
                    v19 := GenS2{F0: 2}
                    _ = v19
                case -2147483648:
                    if (v2 >= (v2 / ((v2) | 1))) {
                        v3 += v3
                        v1 = v1
                        v5 = append(v5, (fi10 ^ 128))
                        if false {
                            v3 -= v3
                            v14 += v14
                        } else {
                            switch ("" + "a") {
                            case "foo bar":
                                var v20 float64 = float64((0.0 * 1e10))
                                _ = v20
                            case "\n\t":
                                if false {
                                    var v21 int32 = 0
                                    v3 = v3
                                    _ = v21
                                } else {
                                    if len(v5) > 0 {
                                        v5[0] = (fi13 << (((v3 << ((v3) & 63))) & 63))
                                    }
                                    v5 = append(v5, (fi10 | 127))
                                }
                            case "":
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            default:
                                if len(v5) > 0 {
                                    v5[0] = fi10
                                }
                            }
                        }
                    } else {
                        var v22 string = "hello"
                        v14 += -1
                        var v23 int64 = v14
                        _ = v22
                        _ = v23
                    }
                default:
                    v24 := GenS4{F0: -99, F1: (v3 ^ 42), F2: 1000}
                    _ = v24
                }
                _ = fi13
                _ = v14
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        v25 := GenS4{F0: 100, F1: 18446744073709551615, F2: 0}
        switch 128 {
        case 2147483647:
            if len(v5) > 0 {
                v5[0] = 9223372036854775807
            }
        case 609:
            if false {
                if false {
                    v26 := GenS5{F0: (-2147483648 <= 2147483647), F1: 42, F2: false}
                    _ = v26
                } else {
                    for fi27 := 0; fi27 < 20; fi27++ {
                        fmt.Println(v3, v2)
                        _ = fi27
                    }
                    fmt.Println(v1, v2, v2)
                }
                v28 := GenS5{F0: (2 >= 0), F1: 9223372036854775807, F2: (false || (false || true))}
                _ = v28
            } else {
                for fi29 := 0; fi29 < 12; fi29++ {
                    if true {
                        v30 := GenS4{F0: -49, F1: (v3 >> (((v3 * (v3 / ((v3) | 1)))) & 63)), F2: -32768}
                        for fi31 := 0; fi31 < 17; fi31++ {
                            v3 -= (v3 ^ v3)
                            var v32 float32 = float32((v1 / float64(3.14159)))
                            fmt.Println(fi29, fi31)
                            _ = fi31
                            _ = v32
                        }
                        _ = v30
                    }
                    _ = fi29
                }
                switch 517 {
                case -982:
                    v2 = v2
                }
                _ = len(v5)
            }
        }
        _ = v25
    }
    if false {
        if false {
            fmt.Println(v2, v3)
            if true {
                v5 = append(v5, 341)
                v33 := GenS2{F0: 509}
                v3 += v3
                _ = v33
            }
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if false {
                switch "foo bar" {
                case "":
                    if false {
                        for fi34 := 0; fi34 < 12; fi34++ {
                            var v35 float64 = 1.5
                            var v36 uint16 = 255
                            v5 = append(v5, -812)
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                        v5 = append(v5, 1)
                        var v37 bool = true
                        _ = v37
                    }
                case "hello":
                    v38 := GenS6{F0: v3, F1: false, F2: 42}
                    _ = v38
                default:
                    var v39 int32 = -2147483648
                    _ = v39
                }
                for fi40 := 0; fi40 < 2; fi40++ {
                    var v41 string = "\n\t"
                    v42 := GenS2{F0: 2}
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                var v43 int32 = -449
                _ = v43
            } else {
                var v44 float64 = ((float64(float32(v1)) * float64(1e10)) + (1.5 * v1))
                fmt.Println(v1, v1, v1)
                if false {
                    var v45 uint16 = 32767
                    var v46 uint = 356
                    _ = v45
                    _ = v46
                } else {
                    if (1 < 255) {
                        switch 127 {
                        case 42:
                            fmt.Println(v1, v1)
                        default:
                            if true {
                                _ = len(v5)
                                v1 += float64(((3.14159 * 0.0) - 0.1))
                            } else {
                                var v47 int8 = 127
                                _ = v47
                            }
                        }
                        switch "hello" {
                        case "hello":
                            var v48 float32 = float32((v1 / 0.1))
                            _ = v48
                        case "foo bar":
                            v49 := GenS3{F0: -393, F1: uint8(v2), F2: 9223372036854775807}
                            _ = v49
                        default:
                            v50 := GenS7{F0: -1, F1: 32767, F2: 2}
                            _ = v50
                        }
                    } else {
                        if false {
                            switch -9223372036854775808 {
                            case -2147483648:
                                v51 := GenS8{F0: 420}
                                _ = v51
                            case 127:
                                switch ((("" + "foo bar") + ("\n\t" + "a")) + (("hello" + "foo bar") + "hello")) {
                                case "":
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                default:
                                    if (-2147483648 <= -30) {
                                        var v52 uint64 = (v2 - 4294967295)
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        for fi53 := 0; fi53 < 13; fi53++ {
                                            var v54 int16 = 1
                                            _ = fi53
                                            _ = v54
                                        }
                                        _ = v52
                                    } else {
                                        if false {
                                            if (false && false) {
                                                for fi55 := 0; fi55 < 12; fi55++ {
                                                    var v56 uint8 = 0
                                                    _ = fi55
                                                    _ = v56
                                                }
                                                if (false || ("hello" == ("a" + ("a" + "\n\t")))) {
                                                    v57 := GenS9{F0: 2147483648, F1: (v3 + (^(v3)))}
                                                    _ = len(v5)
                                                    v44 -= (float64(1e10) - v44)
                                                    _ = v57
                                                } else {
                                                    for fi58 := 0; fi58 < 3; fi58++ {
                                                        for fi59 := 0; fi59 < 18; fi59++ {
                                                            if true {
                                                                v3 -= 0
                                                                switch (("a" + ("a" + "")) + ("a" + "foo bar")) {
                                                                case "a":
                                                                    v60 := GenS7{F0: 42, F1: 824, F2: uint32(v2)}
                                                                    _ = v60
                                                                case "hello":
                                                                    var v61 float32 = 3.14159
                                                                    _ = v61
                                                                }
                                                            }
                                                            if true {
                                                                switch (fi58 >> (((v3 << ((v3) & 63))) & 63)) {
                                                                case 32767:
                                                                    var v62 uint8 = 42
                                                                    _ = v62
                                                                default:
                                                                    fmt.Println(v44, v3, v44)
                                                                }
                                                            }
                                                            v44 -= float64(0.1)
                                                            _ = fi59
                                                        }
                                                        if (32767 >= -32768) {
                                                            if false {
                                                                var v63 uint8 = 2
                                                                for fi64 := 0; fi64 < 11; fi64++ {
                                                                    v2 += v2
                                                                    var v65 uint64 = 968
                                                                    fmt.Println(v1, v2)
                                                                    var v66 uint64 = (v65 >> ((v3) & 63))
                                                                    v3 = v3
                                                                    switch ("" + "hello") {
                                                                    case "a":
                                                                        v67 := GenS10{F0: 1000}
                                                                        _ = v67
                                                                    case "":
                                                                        v65 = v65
                                                                    case "hello":
                                                                        fmt.Println(v2)
                                                                    }
                                                                    _ = fi64
                                                                    _ = v65
                                                                    _ = v66
                                                                }
                                                                _ = v63
                                                            } else {
                                                                _ = len(v5)
                                                            }
                                                        } else {
                                                            v44 = 1.5
                                                            var v68 int64 = -1
                                                            for fi69 := 0; fi69 < 9; fi69++ {
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                var v70 float32 = -2.25
                                                                if (((-(fi58)) > fi69) && true) {
                                                                    v44 = v44
                                                                }
                                                                if ((int8(v3) >= -128) && (!(!(2147483647 == -2147483648)))) {
                                                                    var v71 float32 = v70
                                                                    _ = v71
                                                                } else {
                                                                    v72 := GenS7{F0: fi69, F1: 1, F2: 42}
                                                                    var v73 uint32 = 282
                                                                    fmt.Println(fi58, v1, v1)
                                                                    _ = v72
                                                                    _ = v73
                                                                }
                                                                if (!(!true)) {
                                                                    var v74 uint16 = 101
                                                                    var v75 uint32 = 659
                                                                    if false {
                                                                        var v76 string = ((("hello" + "\n\t") + "hello") + (("foo bar" + "") + ""))
                                                                        v77 := GenS11{F0: -1, F1: false, F2: "hello"}
                                                                        v1 -= (v1 + (v44 - v44))
                                                                        _ = v76
                                                                        _ = v77
                                                                    }
                                                                    _ = v74
                                                                    _ = v75
                                                                }
                                                                _ = fi69
                                                                _ = v70
                                                            }
                                                            _ = v68
                                                        }
                                                        var v78 uint32 = 837
                                                        _ = fi58
                                                        _ = v78
                                                    }
                                                    switch -822 {
                                                    case 4294967295:
                                                        if false {
                                                            var v79 int32 = 287
                                                            _ = v79
                                                        }
                                                    }
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                }
                                            } else {
                                                v1 += float64(((1e10 + 1e10) * 1e10))
                                            }
                                        }
                                        var v80 float64 = v1
                                        var v81 uint16 = 811
                                        _ = v80
                                        _ = v81
                                    }
                                }
                            default:
                                if false {
                                    _ = len(v5)
                                    var v82 int = 4294967295
                                    fmt.Println(v1, v3, v3)
                                    _ = v82
                                } else {
                                    v5 = append(v5, -444)
                                }
                            }
                            var v83 uint = v3
                            _ = v83
                        }
                        var v84 uint = (v3 ^ 18446744073709551615)
                        _ = v84
                    }
                    for fi85 := 0; fi85 < 20; fi85++ {
                        var v86 bool = false
                        for fi87 := 0; fi87 < 8; fi87++ {
                            var v88 int8 = 0
                            if (v86 || (false && v86)) {
                                switch "foo bar" {
                                case "a":
                                    v2 = v2
                                case "\n\t":
                                    if v86 {
                                        var v89 uint = v3
                                        var v90 uint8 = 0
                                        _ = v89
                                        _ = v90
                                    }
                                default:
                                    var v91 int32 = -318
                                    _ = v91
                                }
                            } else {
                                v44 = v1
                                fmt.Println(v86)
                                var v92 uint8 = 42
                                _ = v92
                            }
                            _ = fi87
                            _ = v88
                        }
                        for fi93 := 0; fi93 < 18; fi93++ {
                            v2 = v2
                            _ = fi93
                        }
                        _ = fi85
                        _ = v86
                    }
                }
                _ = v44
            }
            if false {
                v2 = 0
                if false {
                    if false {
                        for fi94 := 0; fi94 < 12; fi94++ {
                            v3 += v3
                            var v95 int64 = 0
                            _ = fi94
                            _ = v95
                        }
                        if (true || false) {
                            fmt.Println(v1)
                            v1 = 0.1
                            v1 -= (0.1 / v1)
                        }
                    } else {
                        switch "\n\t" {
                        case "":
                            v96 := GenS2{F0: 540}
                            _ = v96
                        case "hello":
                            if (!true) {
                                for fi97 := 0; fi97 < 6; fi97++ {
                                    if len(v5) > 0 {
                                        v5[0] = (fi97 + 0)
                                    }
                                    v3 = 812
                                    v98 := GenS5{F0: (false || (!((!(int(v2) == fi97)) || false))), F1: (fi97 << ((253) & 63)), F2: (246 >= 0)}
                                    _ = fi97
                                    _ = v98
                                }
                                switch 0 {
                                case 128:
                                    _ = len(v5)
                                case -9223372036854775808:
                                    for fi99 := 0; fi99 < 18; fi99++ {
                                        v5 = append(v5, (-(fi99)))
                                        for fi100 := 0; fi100 < 6; fi100++ {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            if true {
                                                _ = len(v5)
                                                v2 -= 4294967295
                                                switch (-(fi100)) {
                                                case -122:
                                                    switch "hello" {
                                                    case "\n\t":
                                                        v5 = append(v5, fi100)
                                                    case "foo bar":
                                                        var v101 uint64 = 151
                                                        _ = v101
                                                    case "":
                                                        v102 := GenS3{F0: -1, F1: 67, F2: -9223372036854775808}
                                                        _ = v102
                                                    default:
                                                        v2 = 1
                                                    }
                                                case -939:
                                                    v103 := GenS12{F0: false, F1: "a", F2: 426}
                                                    v104 := GenS8{F0: 949}
                                                    var v105 uint32 = 810
                                                    _ = v103
                                                    _ = v104
                                                    _ = v105
                                                }
                                            }
                                            v1 = (3.14159 / ((1e10 * v1) / v1))
                                            _ = fi100
                                        }
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = fi99
                                    }
                                }
                                if (42 < 413) {
                                    fmt.Println(v1)
                                    var v106 int32 = -73
                                    _ = v106
                                }
                            } else {
                                v3 = v3
                                switch "" {
                                case "hello":
                                    v3 = v3
                                case "foo bar":
                                    if (v1 != float64(1.5)) {
                                        if ((v1 != 1.5) && (!true)) {
                                            v1 = 0.1
                                            for fi107 := 0; fi107 < 16; fi107++ {
                                                var v108 uint = uint(v2)
                                                _ = fi107
                                                _ = v108
                                            }
                                            v109 := GenS2{F0: 42}
                                            _ = v109
                                        } else {
                                            if true {
                                                for fi110 := 0; fi110 < 18; fi110++ {
                                                    fmt.Println(v2, v2)
                                                    var v111 int8 = -1
                                                    if (1.5 < (0.0 + 0.0)) {
                                                        fmt.Println(v3, v111, v3)
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        v112 := GenS10{F0: 42}
                                                        _ = v112
                                                    } else {
                                                        for fi113 := 0; fi113 < 11; fi113++ {
                                                            var v114 uint64 = v2
                                                            switch "hello" {
                                                            case "foo bar":
                                                                v115 := GenS8{F0: -32768}
                                                                _ = v115
                                                            case "a":
                                                                if (-9223372036854775808 > 9223372036854775807) {
                                                                    if false {
                                                                        v111 = (v111 * (v111 << (((v3 + v3)) & 63)))
                                                                        if false {
                                                                            v116 := GenS6{F0: 0, F1: (true || true), F2: 42}
                                                                            var v117 int = -629
                                                                            var v118 uint64 = (^(v114))
                                                                            fmt.Println(v111)
                                                                            v117 = (fi110 >> ((v3) & 63))
                                                                            v3 -= 730
                                                                            if (v3 == 18446744073709551615) {
                                                                                if (false && false) {
                                                                                    v3 += (^(v3))
                                                                                }
                                                                                var v119 float64 = ((float64((3.14159 / -2.25)) - (v1 / 1e10)) - (-2.25 * ((v1 * float64(0.1)) - 0.1)))
                                                                                _ = v119
                                                                            } else {
                                                                                if true {
                                                                                    fmt.Println(fi110)
                                                                                    if ((v117 - (fi110 - fi110)) <= (v117 / ((-1) | 1))) {
                                                                                        v120 := GenS13{F0: 2147483648, F1: v2}
                                                                                        var v121 float64 = v1
                                                                                        _ = v120
                                                                                        _ = v121
                                                                                    } else {
                                                                                        var v122 string = "\n\t"
                                                                                        var v123 int32 = 1
                                                                                        fmt.Println(v117)
                                                                                        _ = v122
                                                                                        _ = v123
                                                                                    }
                                                                                    switch (^(fi110)) {
                                                                                    case -1:
                                                                                        var v124 int64 = -9223372036854775808
                                                                                        if true {
                                                                                            var v125 float64 = (0.0 + (v1 + v1))
                                                                                            _ = v125
                                                                                        } else {
                                                                                            v126 := GenS5{F0: true, F1: (fi113 - fi113), F2: false}
                                                                                            v2 -= (v2 & (v118 | v2))
                                                                                            if false {
                                                                                                var v127 int32 = 2147483647
                                                                                                _ = v127
                                                                                            }
                                                                                            _ = v126
                                                                                        }
                                                                                        _ = v124
                                                                                    case 128:
                                                                                        v118 += v2
                                                                                    case 1:
                                                                                        if true {
                                                                                            var v128 bool = true
                                                                                            _ = v128
                                                                                        }
                                                                                    default:
                                                                                        if true {
                                                                                            v118 = v2
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = len(v5)
                                                                                var v129 int = int(v114)
                                                                                _ = v129
                                                                            }
                                                                            v3 += v3
                                                                            var v130 uint8 = 42
                                                                            _ = v116
                                                                            _ = v117
                                                                            _ = v118
                                                                            _ = v130
                                                                        } else {
                                                                            if false {
                                                                                if false {
                                                                                    if ((-2.25 + v1) > float64(0.0)) {
                                                                                        v131 := GenS3{F0: 1, F1: 49, F2: 0}
                                                                                        v3 += 0
                                                                                        _ = v131
                                                                                    }
                                                                                    v2 += v2
                                                                                } else {
                                                                                    v114 -= (v114 << (((v3 - v3)) & 63))
                                                                                    v114 = 66
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = (fi113 & fi110)
                                                                                    }
                                                                                }
                                                                                var v132 uint16 = 1
                                                                                _ = v132
                                                                            }
                                                                            var v133 uint8 = 127
                                                                            v111 = (v111 << ((v3) & 63))
                                                                            v134 := GenS11{F0: -472, F1: (int8(v133) != -77), F2: ("" + "\n\t")}
                                                                            _ = v133
                                                                            _ = v134
                                                                        }
                                                                        if len(v5) > 0 {
                                                                            v5[0] = -2147483648
                                                                        }
                                                                    } else {
                                                                        if (!(false || false)) {
                                                                            v111 = v111
                                                                            if (v2 == 0) {
                                                                                if len(v5) > 0 {
                                                                                    v5[0] = (fi113 + (fi110 - fi110))
                                                                                }
                                                                            }
                                                                            v111 = (v111 - v111)
                                                                        } else {
                                                                            if (!true) {
                                                                                _ = len(v5)
                                                                            }
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                            switch "" {
                                                                            case "a":
                                                                                if true {
                                                                                    switch fi113 {
                                                                                    case -747:
                                                                                        var v135 int8 = v111
                                                                                        _ = v135
                                                                                    case 4294967295:
                                                                                        switch (("a" + "") + (("" + "") + ("" + ("a" + "")))) {
                                                                                        case "foo bar":
                                                                                            if (!true) {
                                                                                                v1 -= (v1 - v1)
                                                                                            }
                                                                                        case "\n\t":
                                                                                            v136 := GenS11{F0: 1000, F1: (true || (!((42 <= 9223372036854775807) && false))), F2: "hello"}
                                                                                            _ = v136
                                                                                        default:
                                                                                            v137 := GenS14{F0: "", F1: v111}
                                                                                            _ = v137
                                                                                        }
                                                                                    }
                                                                                    switch (fi113 * (fi113 - fi110)) {
                                                                                    case 127:
                                                                                        if len(v5) > 0 {
                                                                                            v5[0] = (fi113 >> ((v3) & 63))
                                                                                        }
                                                                                        if true {
                                                                                            var v138 bool = (float64(float32(v1)) == v1)
                                                                                            var v139 uint32 = 1
                                                                                            var v140 uint16 = 632
                                                                                            _ = v138
                                                                                            _ = v139
                                                                                            _ = v140
                                                                                        } else {
                                                                                            var v141 int = fi113
                                                                                            var v142 int = -9223372036854775808
                                                                                            _ = v141
                                                                                            _ = v142
                                                                                        }
                                                                                    default:
                                                                                        v2 += 473
                                                                                    }
                                                                                }
                                                                            default:
                                                                                if (true && ((v3 == 60) && false)) {
                                                                                    v3 -= v3
                                                                                    switch (fi113 * (fi110 + (fi110 - 830))) {
                                                                                    case 1:
                                                                                        v1 += -2.25
                                                                                    case 32767:
                                                                                        var v143 uint32 = 42
                                                                                        _ = v143
                                                                                    }
                                                                                    fmt.Println(v3, v3)
                                                                                } else {
                                                                                    v5 = append(v5, (fi113 << (((v3 << (((v3 + (v3 % (((^(v3))) | 1)))) & 63))) & 63)))
                                                                                    var v144 int8 = (v111 * (v111 ^ (v111 - -66)))
                                                                                    if (false || false) {
                                                                                        v2 -= 0
                                                                                        v2 -= v114
                                                                                    }
                                                                                    _ = v144
                                                                                }
                                                                            }
                                                                        }
                                                                        if (!true) {
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                            var v145 int = (fi110 << ((v3) & 63))
                                                                            v146 := GenS13{F0: 65535, F1: v114}
                                                                            _ = v145
                                                                            _ = v146
                                                                        } else {
                                                                            v147 := GenS15{F0: v3, F1: 42, F2: 42}
                                                                            _ = v147
                                                                        }
                                                                        v148 := GenS16{F0: 2147483647, F1: (^(v114)), F2: 127}
                                                                        var v149 uint16 = 65535
                                                                        if false {
                                                                            v114 += v114
                                                                            switch (("a" + (("foo bar" + (("\n\t" + ("a" + "hello")) + "a")) + ("\n\t" + ("hello" + "a")))) + (("foo bar" + "foo bar") + "hello")) {
                                                                            case "hello":
                                                                                var v150 int32 = -1
                                                                                _ = v150
                                                                            case "a":
                                                                                v151 := GenS5{F0: (false && true), F1: 988, F2: true}
                                                                                _ = v151
                                                                            }
                                                                        } else {
                                                                            v5 = append(v5, fi110)
                                                                        }
                                                                        _ = v148
                                                                        _ = v149
                                                                    }
                                                                } else {
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                    v2 = 18446744073709551615
                                                                }
                                                            default:
                                                                switch "" {
                                                                case "":
                                                                    var v152 float64 = (v1 / 0.1)
                                                                    _ = v152
                                                                case "\n\t":
                                                                    switch (^(fi113)) {
                                                                    case -1:
                                                                        _ = len(v5)
                                                                    case 0:
                                                                        v5 = append(v5, fi113)
                                                                    }
                                                                case "foo bar":
                                                                    fmt.Println(fi110, v111, fi110)
                                                                default:
                                                                    switch 127 {
                                                                    case -716:
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                    case -24:
                                                                        v153 := GenS5{F0: false, F1: (fi110 << ((v3) & 63)), F2: false}
                                                                        _ = v153
                                                                    case -9223372036854775808:
                                                                        if (-1 <= int32(fi110)) {
                                                                            switch "a" {
                                                                            case "\n\t":
                                                                                fmt.Println(fi110)
                                                                            }
                                                                            v154 := GenS17{F0: (fi113 ^ -481), F1: 986, F2: v111}
                                                                            _ = v154
                                                                        } else {
                                                                            if len(v5) > 0 {
                                                                                v5[0] = (fi110 << (((v3 | 65535)) & 63))
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            _ = fi113
                                                            _ = v114
                                                        }
                                                    }
                                                    _ = fi110
                                                    _ = v111
                                                }
                                            } else {
                                                fmt.Println(v3)
                                            }
                                            if ((!true) || false) {
                                                for fi155 := 0; fi155 < 15; fi155++ {
                                                    var v156 uint = (v3 & (v3 + v3))
                                                    if len(v5) > 0 {
                                                        v5[0] = (fi155 ^ 0)
                                                    }
                                                    _ = fi155
                                                    _ = v156
                                                }
                                            } else {
                                                for fi157 := 0; fi157 < 7; fi157++ {
                                                    v1 -= (0.0 - (((0.1 - float64(3.14159)) + (v1 - float64((float32(-2.25) + -2.25)))) - float64(-2.25)))
                                                    v158 := GenS18{F0: ((0 > 0) && (1.5 != 1e10))}
                                                    _ = fi157
                                                    _ = v158
                                                }
                                                switch (("\n\t" + "") + "") {
                                                case "hello":
                                                    _ = len(v5)
                                                case "\n\t":
                                                    var v159 float32 = (1.5 + 1e10)
                                                    _ = v159
                                                case "foo bar":
                                                    if true {
                                                        switch "a" {
                                                        case "":
                                                            v1 -= ((v1 * v1) - (float64((0.0 * float32(3.14159))) * (1e10 * float64(3.14159))))
                                                        }
                                                        if (2147483648 != 2147483647) {
                                                            var v160 string = "a"
                                                            _ = v160
                                                        } else {
                                                            var v161 uint = v3
                                                            _ = v161
                                                        }
                                                        if (2147483647 >= 0) {
                                                            for fi162 := 0; fi162 < 6; fi162++ {
                                                                v1 += v1
                                                                var v163 int16 = int16(v3)
                                                                fmt.Println(v163, v3)
                                                                _ = fi162
                                                                _ = v163
                                                            }
                                                            for fi164 := 0; fi164 < 13; fi164++ {
                                                                switch (fi164 - -1) {
                                                                case -1:
                                                                    var v165 uint16 = 42
                                                                    _ = v165
                                                                case 0:
                                                                    for fi166 := 0; fi166 < 14; fi166++ {
                                                                        fmt.Println(fi164, fi164)
                                                                        _ = fi166
                                                                    }
                                                                default:
                                                                    fmt.Println(fi164, fi164, fi164)
                                                                }
                                                                fmt.Println(v2, v2)
                                                                _ = fi164
                                                            }
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                fmt.Println(v2, v1, v2)
                                                            }
                                                        } else {
                                                            v167 := GenS18{F0: ((false && false) && true)}
                                                            if len(v5) > 0 {
                                                                v5[0] = -1
                                                            }
                                                            _ = v167
                                                        }
                                                    } else {
                                                        for fi168 := 0; fi168 < 12; fi168++ {
                                                            var v169 float64 = (3.14159 - (1e10 * ((1.5 / v1) + v1)))
                                                            if (!(v3 == (v3 * (v3 % (((v3 - (^(v3)))) | 1))))) {
                                                                v2 += 4294967295
                                                                if false {
                                                                    var v170 uint32 = 0
                                                                    var v171 uint8 = 1
                                                                    _ = v170
                                                                    _ = v171
                                                                }
                                                                var v172 int32 = 0
                                                                _ = v172
                                                            } else {
                                                                var v173 uint = (v3 / ((v3) | 1))
                                                                v174 := GenS12{F0: false, F1: "", F2: 4294967295}
                                                                _ = v173
                                                                _ = v174
                                                            }
                                                            _ = fi168
                                                            _ = v169
                                                        }
                                                        v175 := GenS6{F0: (v3 << (((v3 << ((v3) & 63))) & 63)), F1: (-2147483648 <= 992), F2: -9223372036854775808}
                                                        _ = v175
                                                    }
                                                }
                                            }
                                        }
                                    } else {
                                        v2 = v2
                                    }
                                }
                            }
                        case "foo bar":
                            var v176 int64 = 9223372036854775807
                            _ = v176
                        default:
                            v5 = append(v5, 2147483647)
                        }
                    }
                    if (255 > 2) {
                        switch ("" + "a") {
                        case "foo bar":
                            var v177 int64 = -1
                            _ = v177
                        }
                        fmt.Println(v1, v3)
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    var v178 int64 = 0
                    _ = v178
                } else {
                    switch "hello" {
                    case "hello":
                        if len(v5) > 0 {
                            v5[0] = -169
                        }
                    case "":
                        if (0 >= 1) {
                            _ = len(v5)
                            v3 = 42
                            for fi179 := 0; fi179 < 2; fi179++ {
                                if ((-1 == 9223372036854775807) && false) {
                                    v1 = v1
                                }
                                var v180 float64 = (1e10 - v1)
                                _ = fi179
                                _ = v180
                            }
                        } else {
                            v181 := GenS19{F0: 1, F1: -2147483648, F2: 1}
                            var v182 int64 = 42
                            for fi183 := 0; fi183 < 13; fi183++ {
                                for fi184 := 0; fi184 < 17; fi184++ {
                                    var v185 float32 = -2.25
                                    v186 := GenS20{F0: -255, F1: 1, F2: (("foo bar" + ("\n\t" + "foo bar")) + "\n\t")}
                                    v187 := GenS21{F0: 100, F1: int16(v182)}
                                    _ = fi184
                                    _ = v185
                                    _ = v186
                                    _ = v187
                                }
                                var v188 int16 = 42
                                v5 = append(v5, fi183)
                                _ = fi183
                                _ = v188
                            }
                            _ = v181
                            _ = v182
                        }
                    case "\n\t":
                        v3 = 0
                    }
                    v3 = v3
                }
                v1 = v1
            }
        }
        fmt.Println(v2, v3)
        var v189 int = 127
        _ = v189
    } else {
        v3 = (v3 + 386)
        var v190 int = 4294967295
        _ = v190
    }
    v191 := GenS15{F0: uint(v2), F1: -32768, F2: 1}
    v192 := GenS22{F0: 94, F1: false, F2: 127}
    fmt.Println(v1, v2, v1)
    v1 += (1.5 * v1)
    if (false && true) {
        switch "a" {
        case "\n\t":
            v1 -= 0.1
        case "":
            fmt.Println(v2)
        case "foo bar":
            if (float64(1e10) == v1) {
                for fi193 := 0; fi193 < 12; fi193++ {
                    if (!true) {
                        v1 = -2.25
                    }
                    var v194 uint = v3
                    _ = fi193
                    _ = v194
                }
                if false {
                    for fi195 := 0; fi195 < 16; fi195++ {
                        for fi196 := 0; fi196 < 19; fi196++ {
                            switch (fi196 & (fi195 ^ fi196)) {
                            case 1:
                                if true {
                                    if (false || true) {
                                        v2 += (^(v2))
                                    }
                                    v197 := GenS23{F0: -96, F1: 1, F2: 127}
                                    _ = v197
                                }
                            case -63:
                                if false {
                                    if (42 >= 2147483647) {
                                        var v198 float64 = 3.14159
                                        fmt.Println(v1)
                                        v1 -= 1.5
                                        if (true && (!(true || true))) {
                                            switch fi196 {
                                            case 42:
                                                v198 = -2.25
                                            case 70:
                                                var v199 uint32 = 4294967295
                                                _ = v199
                                            case 1:
                                                if (false || (-2147483648 > 0)) {
                                                    v2 -= (v2 - v2)
                                                }
                                            default:
                                                var v200 int = (fi196 - -2147483648)
                                                _ = v200
                                            }
                                            if (100 > 127) {
                                                var v201 uint32 = 65535
                                                switch int(v3) {
                                                case 2147483647:
                                                    if true {
                                                        v1 = 3.14159
                                                        v3 = 42
                                                        v202 := GenS24{F0: 906}
                                                        _ = v202
                                                    }
                                                case 128:
                                                    var v203 float32 = 1e10
                                                    _ = v203
                                                }
                                                fmt.Println(fi196)
                                                var v204 float32 = (-2.25 - 1.5)
                                                if len(v5) > 0 {
                                                    v5[0] = (fi196 - fi195)
                                                }
                                                _ = v201
                                                _ = v204
                                            }
                                        }
                                        _ = v198
                                    } else {
                                        v205 := GenS25{F0: fi196}
                                        _ = v205
                                    }
                                    var v206 uint = v3
                                    if (!false) {
                                        _ = len(v5)
                                        var v207 string = "a"
                                        var v208 uint16 = 1
                                        var v209 int16 = 42
                                        v3 = 0
                                        switch 127 {
                                        case -502:
                                            if (v207 == ("foo bar" + v207)) {
                                                if false {
                                                    v208 = (v208 * v208)
                                                    switch (fi195 * fi195) {
                                                    case -354:
                                                        v1 += (((0.0 - (float64(((-2.25 / 0.1) * (1e10 / (1e10 * 3.14159)))) * (v1 + (v1 / 3.14159)))) * v1) + float64((0.0 - 1.5)))
                                                        var v210 bool = (v206 <= v3)
                                                        var v211 string = v207
                                                        v209 = -642
                                                        fmt.Println(v211, v1, v1)
                                                        var v212 uint = (v3 | (v3 / (((v206 + (v3 & (v3 | (v3 & (v3 / (((v3 / ((v3) | 1))) | 1))))))) | 1)))
                                                        var v213 uint16 = v208
                                                        var v214 string = (("" + ("hello" + v211)) + "")
                                                        _ = v210
                                                        _ = v211
                                                        _ = v212
                                                        _ = v213
                                                        _ = v214
                                                    case 32767:
                                                        var v215 uint32 = 2147483648
                                                        _ = v215
                                                    case -229:
                                                        if false {
                                                            var v216 float64 = (-2.25 + v1)
                                                            v217 := GenS26{F0: v2, F1: v207, F2: ("a" + "")}
                                                            v218 := GenS27{F0: (v208 ^ (v208 * v208)), F1: fi195, F2: -1}
                                                            _ = v216
                                                            _ = v217
                                                            _ = v218
                                                        } else {
                                                            v1 += v1
                                                            v207 = ((((v207 + ("hello" + v207)) + "foo bar") + v207) + ((v207 + ("" + "")) + v207))
                                                            v2 = uint64(v206)
                                                            switch v207 {
                                                            case "\n\t":
                                                                fmt.Println(v1, v3, v209)
                                                            default:
                                                                var v219 int64 = -471
                                                                _ = v219
                                                            }
                                                            fmt.Println(v208, v208, v2)
                                                            var v220 string = (v207 + v207)
                                                            var v221 float32 = (1e10 + 0.1)
                                                            v5 = append(v5, fi195)
                                                            _ = v220
                                                            _ = v221
                                                        }
                                                    }
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                } else {
                                                    var v222 uint64 = v2
                                                    switch fi196 {
                                                    case -2147483648:
                                                        var v223 string = v207
                                                        _ = v223
                                                    case -307:
                                                        v208 = 32767
                                                    case 4294967295:
                                                        if (true && true) {
                                                            v207 = (v207 + "hello")
                                                        }
                                                    default:
                                                        var v224 uint = (v3 & v3)
                                                        _ = v224
                                                    }
                                                    _ = v222
                                                }
                                                _ = len(v5)
                                                var v225 uint16 = v208
                                                _ = v225
                                            } else {
                                                v226 := GenS4{F0: -1, F1: (v206 >> (((v206 % (((v206 << ((uint(v208)) & 63))) | 1))) & 63)), F2: (v209 << ((v3) & 63))}
                                                _ = v226
                                            }
                                        }
                                        _ = v207
                                        _ = v208
                                        _ = v209
                                    } else {
                                        v206 = v206
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                    _ = v206
                                } else {
                                    fmt.Println(fi195, fi196)
                                    fmt.Println(fi196)
                                    _ = len(v5)
                                }
                            }
                            var v227 uint32 = 1
                            _ = fi196
                            _ = v227
                        }
                        _ = fi195
                    }
                    switch -994 {
                    case -2147483648:
                        switch 32767 {
                        case 42:
                            if false {
                                v228 := GenS25{F0: 128}
                                v5 = append(v5, 0)
                                var v229 uint16 = 2
                                _ = v228
                                _ = v229
                            } else {
                                var v230 uint64 = (v2 - v2)
                                if true {
                                    var v231 uint8 = 2
                                    _ = v231
                                }
                                _ = v230
                            }
                        }
                    case 1:
                        for fi232 := 0; fi232 < 6; fi232++ {
                            fmt.Println(v3, fi232)
                            v233 := GenS28{F0: 194, F1: true, F2: -128}
                            var v234 uint16 = 65535
                            _ = fi232
                            _ = v233
                            _ = v234
                        }
                    }
                    var v235 int32 = int32(v3)
                    _ = v235
                } else {
                    var v236 string = "foo bar"
                    if (true || (v2 >= (v2 >> (((v3 + (v3 ^ (v3 + 0)))) & 63)))) {
                        var v237 uint64 = (^(v2))
                        if (128 == 248) {
                            v237 += (v2 ^ (v237 << ((v3) & 63)))
                            for fi238 := 0; fi238 < 15; fi238++ {
                                v1 = v1
                                _ = fi238
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        } else {
                            if ((v236 <= v236) || (1.5 != 1e10)) {
                                v239 := GenS10{F0: -279}
                                _ = v239
                            } else {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v237 = v2
                                for fi240 := 0; fi240 < 6; fi240++ {
                                    if len(v5) > 0 {
                                        v5[0] = (^(fi240))
                                    }
                                    _ = fi240
                                }
                            }
                            var v241 int8 = 53
                            _ = v241
                        }
                        _ = v237
                    } else {
                        for fi242 := 0; fi242 < 7; fi242++ {
                            var v243 int = 17
                            if true {
                                var v244 int32 = -323
                                _ = v244
                            } else {
                                v243 = fi242
                                v243 = v243
                                if (!false) {
                                    fmt.Println(v236)
                                    if (1 < 42) {
                                        fmt.Println(v2)
                                    } else {
                                        v245 := GenS29{F0: -9223372036854775808}
                                        _ = v245
                                    }
                                    v246 := GenS21{F0: 0, F1: -1}
                                    _ = v246
                                }
                            }
                            for fi247 := 0; fi247 < 4; fi247++ {
                                if false {
                                    var v248 uint16 = 1
                                    _ = v248
                                }
                                v249 := GenS9{F0: 1000, F1: v3}
                                var v250 float64 = v1
                                _ = fi247
                                _ = v249
                                _ = v250
                            }
                            _ = fi242
                            _ = v243
                        }
                        v251 := GenS30{F0: -32768, F1: -2147483648, F2: (42 <= -86)}
                        _ = v251
                    }
                    _ = v236
                }
            }
        default:
            switch 42 {
            case 2147483647:
                var v252 int16 = -32768
                _ = v252
            case 128:
                if false {
                    v253 := GenS21{F0: 1, F1: 1000}
                    _ = v253
                } else {
                    v3 += (v3 | v3)
                    if len(v5) > 0 {
                        v5[0] = -545
                    }
                    v1 = ((((v1 / v1) - (v1 - ((v1 / v1) - (v1 - 3.14159)))) + ((((0.1 * v1) / v1) * v1) / v1)) / 1.0)
                }
            case 4294967295:
                if true {
                    if true {
                        var v254 uint8 = 0
                        _ = v254
                    }
                } else {
                    fmt.Println(v2, v2)
                }
            default:
                var v255 int16 = 932
                _ = v255
            }
        }
    } else {
        var v256 int8 = 1
        if true {
            v2 += (v2 + v2)
        }
        for fi257 := 0; fi257 < 14; fi257++ {
            if false {
                for fi258 := 0; fi258 < 13; fi258++ {
                    var v259 float32 = 0.1
                    var v260 uint16 = 0
                    _ = fi258
                    _ = v259
                    _ = v260
                }
                for fi261 := 0; fi261 < 6; fi261++ {
                    v256 = -1
                    switch (fi257 / ((fi257) | 1)) {
                    case 2147483647:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case 1:
                        switch "a" {
                        case "":
                            var v262 uint16 = 65535
                            _ = v262
                        case "foo bar":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    case 777:
                        fmt.Println(v3, fi261, v1)
                    default:
                        fmt.Println(v2, v2)
                    }
                    if len(v5) > 0 {
                        v5[0] = fi257
                    }
                    _ = fi261
                }
                for fi263 := 0; fi263 < 20; fi263++ {
                    var v264 bool = true
                    var v265 uint = (v3 << ((42) & 63))
                    if (1 < 1) {
                        if v264 {
                            var v266 float64 = float64(-2.25)
                            _ = v266
                        }
                        v267 := GenS17{F0: int(v256), F1: 1, F2: v256}
                        _ = v267
                    } else {
                        v5 = append(v5, fi263)
                        switch int(v265) {
                        case 128:
                            v1 += (((1.5 / v1) - v1) * v1)
                        case 978:
                            var v268 int64 = 9223372036854775807
                            _ = v268
                        }
                    }
                    _ = fi263
                    _ = v264
                    _ = v265
                }
            } else {
                v3 -= v3
                if ((10 != 5) && (((v256 > v256) || (!true)) && (!(false && false)))) {
                    v269 := GenS31{F0: false, F1: true, F2: 127}
                    _ = len(v5)
                    var v270 int32 = -1
                    _ = v269
                    _ = v270
                } else {
                    var v271 uint32 = 535
                    if false {
                        switch (^(fi257)) {
                        case 0:
                            var v272 int8 = -1
                            _ = v272
                        case 720:
                            for fi273 := 0; fi273 < 5; fi273++ {
                                if ((float32(v1) != 1e10) || (!false)) {
                                    switch (fi273 ^ (fi257 | fi273)) {
                                    case 4294967295:
                                        var v274 uint16 = 32767
                                        _ = v274
                                    case 245:
                                        var v275 int = 32768
                                        _ = v275
                                    case 0:
                                        fmt.Println(fi273, v271, v1)
                                    }
                                }
                                fmt.Println(fi273, fi273, v2)
                                _ = fi273
                            }
                        case 324:
                            v256 += v256
                        }
                        switch "a" {
                        case "\n\t":
                            switch ("a" + "foo bar") {
                            case "foo bar":
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            case "a":
                                if true {
                                    if true {
                                        if false {
                                            fmt.Println(v2, v3)
                                        }
                                        fmt.Println(v2)
                                    }
                                    v276 := GenS25{F0: 913}
                                    _ = v276
                                } else {
                                    if (65535 == 1) {
                                        v277 := GenS22{F0: v256, F1: (!(!false)), F2: 42}
                                        _ = v277
                                    } else {
                                        v278 := GenS30{F0: -32768, F1: -1, F2: true}
                                        _ = v278
                                    }
                                }
                            default:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                        case "a":
                            var v279 bool = (!false)
                            _ = v279
                        default:
                            fmt.Println(v2, v3, v3)
                        }
                    } else {
                        for fi280 := 0; fi280 < 6; fi280++ {
                            if true {
                                v256 = int8(fi280)
                                fmt.Println(v3, v3, v3)
                                v3 = uint(v271)
                            }
                            var v281 uint16 = 687
                            v282 := GenS30{F0: int16(v271), F1: 0, F2: true}
                            _ = fi280
                            _ = v281
                            _ = v282
                        }
                        var v283 int64 = -9223372036854775808
                        _ = v283
                    }
                    _ = v271
                }
                fmt.Println(v1, v3, v1)
            }
            v3 = (v3 ^ (^(v3)))
            _ = fi257
        }
        _ = v256
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v191
    _ = v192
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
