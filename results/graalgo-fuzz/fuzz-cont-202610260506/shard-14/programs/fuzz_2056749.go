package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint32; F2 bool }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 uint8 }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 int16; F1 uint16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int16; F1 int16 }
type GenS11 struct { F0 int64; F1 uint32 }
type GenS12 struct { F0 int16; F1 int }
type GenS13 struct { F0 uint64; F1 int8; F2 int64 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 int }
type GenS16 struct { F0 int32; F1 int32; F2 uint8 }
type GenS17 struct { F0 int64; F1 uint }
type GenS18 struct { F0 int; F1 int32 }
type GenS19 struct { F0 uint8 }
type GenS20 struct { F0 uint16; F1 int }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{237, 42, 787}
    v3 := []int{42}
    var v4 float64 = ((-2.25 + (0.0 + 0.0)) + 1e10)
    v5 := make(map[uint32]bool)
    var v6 int64 = 42
    v6 -= v6
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            switch "" {
            case "foo bar":
                var v9 string = "foo bar"
                _ = v9
            default:
                var v10 uint32 = 1
                _ = v10
            }
            _ = fi8
        }
        if false {
            if (!(false || ((false && false) || ((false || (!false)) || false)))) {
                switch "hello" {
                case "foo bar":
                    switch ("\n\t" + ("\n\t" + ("\n\t" + ("\n\t" + "")))) {
                    case "":
                        if ("hello" != "foo bar") {
                            var v11 uint64 = 18446744073709551615
                            fmt.Println(v4, v11, v4)
                            var v12 int8 = -1
                            _ = v11
                            _ = v12
                        }
                    case "a":
                        v13 := GenS1{F0: 65535, F1: 989, F2: 235}
                        _ = v13
                    case "\n\t":
                        v4 = 1.5
                    default:
                        v14 := GenS1{F0: 0, F1: 2147483647, F2: 1}
                        _ = v14
                    }
                case "a":
                    if true {
                        fmt.Println(v4, fi7)
                        switch "a" {
                        case "":
                            if false {
                                switch -627 {
                                case 2147483647:
                                    for fi15 := 0; fi15 < 13; fi15++ {
                                        if false {
                                            var v16 uint64 = 4294967295
                                            _ = v16
                                        } else {
                                            v17 := GenS1{F0: 32767, F1: 42, F2: 610}
                                            if (0 == 255) {
                                                switch ("foo bar" + (((("a" + "a") + "hello") + ("" + "")) + "hello")) {
                                                case "":
                                                    var v18 float64 = 1.5
                                                    _ = v18
                                                case "hello":
                                                    v6 -= (-(v6))
                                                    var v19 uint16 = 403
                                                    if true {
                                                        v4 += (-2.25 + 0.0)
                                                        _ = len(v5)
                                                    } else {
                                                        if (fi15 <= (fi15 * fi15)) {
                                                            var v20 int = fi7
                                                            v21 := GenS1{F0: (v19 * 1), F1: -634, F2: 634}
                                                            var v22 uint64 = 18446744073709551615
                                                            fmt.Println(v4)
                                                            _ = len(v1)
                                                            _ = v20
                                                            _ = v21
                                                            _ = v22
                                                        } else {
                                                            v19 = v19
                                                        }
                                                        v19 += v19
                                                        if (!false) {
                                                            if (true && false) {
                                                                fmt.Println(fi7, fi7, fi7)
                                                            } else {
                                                                v4 -= float64(0.0)
                                                                var v23 uint = uint(v19)
                                                                var v24 int64 = v6
                                                                _ = v23
                                                                _ = v24
                                                            }
                                                            fmt.Println(fi15, fi7, v4)
                                                            v25 := GenS1{F0: 42, F1: -282, F2: 0}
                                                            var v26 uint32 = 849
                                                            _ = v25
                                                            _ = v26
                                                        }
                                                    }
                                                    _ = v19
                                                case "foo bar":
                                                    if (0 >= 1) {
                                                        v27 := GenS2{F0: 255}
                                                        _ = v27
                                                    }
                                                default:
                                                    v4 += (((v4 - v4) - float64(0.0)) + ((0.0 * 1.5) + v4))
                                                }
                                                v6 = v6
                                            } else {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    switch fi15 {
                                                    case -2147483648:
                                                        v28 := GenS3{F0: false, F1: 42, F2: false}
                                                        _ = v28
                                                    case 32767:
                                                        fmt.Println(fi15, v4)
                                                    case 757:
                                                        if false {
                                                            v29 := GenS1{F0: 351, F1: -769, F2: 2}
                                                            v6 = -1
                                                            _ = v29
                                                        }
                                                    default:
                                                        v3 = append(v3, 32767)
                                                    }
                                                case "":
                                                    if len(v3) > 0 {
                                                        v3[0] = (fi7 % (((fi15 % ((fi7) | 1))) | 1))
                                                    }
                                                case "hello":
                                                    v30 := GenS3{F0: (!false), F1: 0, F2: ((!true) && true)}
                                                    _ = v30
                                                default:
                                                    _ = len(v2)
                                                }
                                                v6 += (^(v6))
                                                var v31 int8 = 1
                                                var v32 uint64 = 926
                                                _ = v31
                                                _ = v32
                                            }
                                            v4 = (0.0 * 0.1)
                                            _ = v17
                                        }
                                        v6 += -740
                                        var v33 uint = 42
                                        v33 = v33
                                        var v34 uint16 = 42
                                        _ = fi15
                                        _ = v33
                                        _ = v34
                                    }
                                }
                                if (!false) {
                                    v35 := GenS4{F0: v6, F1: 0}
                                    for fi36 := 0; fi36 < 17; fi36++ {
                                        if true {
                                            delete(v5, 42)
                                        } else {
                                            var v37 uint8 = 146
                                            _ = v37
                                        }
                                        switch (fi7 << ((1) & 63)) {
                                        case 32767:
                                            fmt.Println(v6, v4, fi36)
                                        case 4294967295:
                                            if (-32768 < -960) {
                                                var v38 int = fi36
                                                v39 := GenS5{F0: (v38 << ((2) & 63))}
                                                _ = v38
                                                _ = v39
                                            } else {
                                                var v40 int = fi36
                                                v6 += (v6 << ((335) & 63))
                                                _ = v40
                                            }
                                        case 1:
                                            v4 = (v4 * (v4 - 1.5))
                                        default:
                                            fmt.Println(v6)
                                        }
                                        _ = fi36
                                    }
                                    var v41 float64 = v4
                                    _ = v35
                                    _ = v41
                                }
                                for fi42 := 0; fi42 < 11; fi42++ {
                                    v43 := GenS4{F0: (^(v6)), F1: 42}
                                    var v44 float32 = 3.14159
                                    v2 = append(v2, 4294967295)
                                    _ = fi42
                                    _ = v43
                                    _ = v44
                                }
                            } else {
                                v4 = v4
                                for fi45 := 0; fi45 < 10; fi45++ {
                                    switch ("hello" + "a") {
                                    case "a":
                                        var v46 string = "a"
                                        _ = v46
                                    case "hello":
                                        var v47 uint16 = 480
                                        _ = v47
                                    default:
                                        var v48 uint = 18446744073709551615
                                        _ = v48
                                    }
                                    v49 := v5[525]
                                    v1[42] = (100 == 127)
                                    _ = fi45
                                    _ = v49
                                }
                                if (!(2147483647 != 0)) {
                                    _ = len(v3)
                                    var v50 int64 = (v6 << ((852) & 63))
                                    if true {
                                        fmt.Println(fi7, v4)
                                        var v51 uint32 = 65535
                                        var v52 int = fi7
                                        _ = v51
                                        _ = v52
                                    }
                                    _ = v50
                                } else {
                                    if (436 == 1) {
                                        v4 = v4
                                        v4 = 0.1
                                    }
                                }
                            }
                        case "foo bar":
                            v53 := GenS6{F0: -1, F1: 42}
                            _ = v53
                        }
                    } else {
                        if ("" == (("foo bar" + ("foo bar" + "\n\t")) + "\n\t")) {
                            v2 = append(v2, 0)
                        } else {
                            _ = len(v5)
                            if false {
                                v6 -= (v6 ^ v6)
                                if (-713 != 1) {
                                    v6 = (v6 >> ((4294967295) & 63))
                                    var v54 float32 = (3.14159 * 1e10)
                                    _ = v54
                                } else {
                                    var v55 int64 = (v6 >> ((2) & 63))
                                    v4 += v4
                                    var v56 uint8 = 43
                                    _ = v55
                                    _ = v56
                                }
                            }
                            v57 := GenS4{F0: -9223372036854775808, F1: 42}
                            _ = v57
                        }
                        v58 := GenS7{F0: 811, F1: ""}
                        if true {
                            fmt.Println(v6, v4, fi7)
                        }
                        _ = v58
                    }
                }
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    v59 := GenS6{F0: 42, F1: 255}
                    _ = v59
                }
            }
            var v60 int64 = 1
            _ = v60
        }
        _ = fi7
    }
    var v61 uint = 4294967295
    switch 0 {
    case 93:
        v62 := GenS8{F0: -1, F1: 42}
        _ = v62
    case 9223372036854775807:
        for fi63 := 0; fi63 < 4; fi63++ {
            if true {
                var v64 uint8 = 1
                switch fi63 {
                case -682:
                    v61 -= (v61 ^ 744)
                case 1:
                    if false {
                        v64 = (v64 / ((v64) | 1))
                    } else {
                        var v65 int64 = (v6 ^ v6)
                        if ((fi63 << (((v61 * 2)) & 63)) != 621) {
                            v66 := GenS8{F0: 1, F1: 0}
                            _ = v66
                        } else {
                            v67 := GenS9{F0: 18446744073709551615}
                            var v68 string = ""
                            _ = v67
                            _ = v68
                        }
                        _ = v65
                    }
                }
                switch (fi63 * 42) {
                case 689:
                    var v69 uint64 = 18446744073709551615
                    _ = v69
                case 4294967295:
                    v70 := GenS10{F0: int16(v61), F1: 42}
                    _ = v70
                default:
                    if false {
                        v4 += v4
                    } else {
                        v71 := GenS11{F0: (v6 - (v6 | v6)), F1: 65535}
                        switch "" {
                        case "\n\t":
                            var v72 int16 = 69
                            _ = v72
                        case "a":
                            switch ("\n\t" + "a") {
                            case "":
                                v73 := GenS12{F0: 0, F1: 32768}
                                _ = v73
                            case "foo bar":
                                var v74 bool = true
                                _ = v74
                            case "hello":
                                v75 := GenS13{F0: 1, F1: 100, F2: (v6 << (((v61 - (v61 & v61))) & 63))}
                                _ = v75
                            default:
                                if false {
                                    if (v6 != -9223372036854775808) {
                                        for fi76 := 0; fi76 < 12; fi76++ {
                                            switch -250 {
                                            case 4294967295:
                                                var v77 float64 = (float64(1e10) * -2.25)
                                                _ = v77
                                            case -60:
                                                v64 += (v64 / ((v64) | 1))
                                            }
                                            if (0.1 != (v4 - ((v4 - float64(-2.25)) * (((((v4 * -2.25) - v4) - v4) - ((((float64(1e10) - ((v4 + v4) * ((0.0 * v4) + ((v4 * 3.14159) - v4)))) + 1.5) + v4) / v4)) / v4)))) {
                                                if (true || (false || (42 <= 65535))) {
                                                    if ((false || false) || (290 <= 1)) {
                                                        v78 := GenS14{F0: 0}
                                                        _ = v78
                                                    }
                                                    v79 := GenS3{F0: true, F1: 464, F2: ((v64 + v64) < v64)}
                                                    _ = v79
                                                } else {
                                                    var v80 uint8 = v64
                                                    v5[1] = false
                                                    v80 = 2
                                                    _ = v80
                                                }
                                                v81 := GenS9{F0: 42}
                                                _ = v81
                                            }
                                            v82 := GenS9{F0: 18446744073709551615}
                                            _ = len(v2)
                                            v83 := GenS15{F0: fi63}
                                            _ = fi76
                                            _ = v82
                                            _ = v83
                                        }
                                        var v84 uint64 = 42
                                        _ = v84
                                    }
                                    for fi85 := 0; fi85 < 8; fi85++ {
                                        if (uint(v64) == (v61 / ((v61) | 1))) {
                                            v4 += float64((float32(-2.25) + 0.0))
                                            var v86 uint = (v61 + (v61 >> (((v61 & v61)) & 63)))
                                            _ = v86
                                        }
                                        v64 += (v64 / ((v64) | 1))
                                        _ = fi85
                                    }
                                    v6 += -939
                                }
                            }
                        case "":
                            for fi87 := 0; fi87 < 6; fi87++ {
                                var v88 string = "foo bar"
                                if true {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if true {
                                        var v89 float32 = -2.25
                                        var v90 uint32 = 2147483647
                                        _ = v89
                                        _ = v90
                                    } else {
                                        v91 := GenS2{F0: v64}
                                        _ = v91
                                    }
                                    var v92 int64 = (v6 | (-(v6)))
                                    _ = v92
                                } else {
                                    v4 = (((v4 + v4) - float64(1e10)) * v4)
                                    fmt.Println(v88, v61)
                                }
                                _ = fi87
                                _ = v88
                            }
                        default:
                            v64 -= (v64 | uint8(v61))
                        }
                        _ = v71
                    }
                }
                _ = v64
            } else {
                v61 = (v61 - v61)
                v6 -= v6
            }
            for fi93 := 0; fi93 < 12; fi93++ {
                v61 = (^(v61))
                var v94 string = ""
                var v95 uint32 = 2147483648
                _ = fi93
                _ = v94
                _ = v95
            }
            var v96 uint32 = 2
            _ = fi63
            _ = v96
        }
    case 911:
        v97 := GenS16{F0: 568, F1: -2147483648, F2: 128}
        _ = v97
    default:
        if true {
            if ((0 <= -1) || false) {
                v6 += -657
                v6 = (v6 & (v6 % ((v6) | 1)))
            }
            v98 := GenS17{F0: -9223372036854775808, F1: v61}
            v5[2147483647] = (false && true)
            _ = v98
        }
    }
    for fi99 := 0; fi99 < 1; fi99++ {
        v3 = append(v3, (^(fi99)))
        _ = fi99
    }
    for fi100 := 0; fi100 < 20; fi100++ {
        if (((-2.25 * v4) / float64(1e10)) == 1.5) {
            fmt.Println(v4, v61, v4)
            if true {
                fmt.Println(v61, v4)
                v101 := GenS18{F0: 4294967295, F1: -722}
                _ = v101
            }
        } else {
            if (-1 >= -1) {
                if ((!true) && false) {
                    var v102 int32 = 2147483647
                    switch "foo bar" {
                    case "hello":
                        v4 = (v4 * -2.25)
                    case "":
                        var v103 uint16 = 61
                        _ = v103
                    case "a":
                        var v104 bool = false
                        _ = v104
                    }
                    _ = v102
                }
                var v105 uint = (v61 & (v61 / (((v61 - 2)) | 1)))
                _ = v105
            } else {
                for fi106 := 0; fi106 < 14; fi106++ {
                    switch "hello" {
                    case "foo bar":
                        var v107 uint64 = 4294967295
                        _ = v107
                    case "hello":
                        var v108 float64 = (3.14159 - 1e10)
                        _ = v108
                    }
                    v61 = 2
                    if (-53 >= 42) {
                        var v109 uint8 = uint8(v61)
                        v61 = (v61 | (^(v61)))
                        _ = v109
                    }
                    var v110 uint16 = 967
                    if (int32(fi100) <= 1) {
                        var v111 int32 = -177
                        v61 -= v61
                        fmt.Println(fi106, fi100, v110)
                        _ = v111
                    }
                    _ = fi106
                    _ = v110
                }
                if true {
                    var v112 uint = 358
                    delete(v1, 234)
                    _ = v112
                }
            }
            var v113 uint32 = 641
            _ = v113
        }
        v114 := GenS2{F0: 2}
        switch (fi100 & fi100) {
        case 2147483647:
            var v115 float64 = float64(0.1)
            _ = v115
        case 1:
            switch 127 {
            case 295:
                switch "foo bar" {
                case "":
                    var v116 float32 = 1e10
                    _ = v116
                case "foo bar":
                    if (!true) {
                        v3 = append(v3, fi100)
                        v117 := GenS14{F0: 612}
                        if ((4294967295 == 4294967295) || ((1e10 / 0.1) != 3.14159)) {
                            v6 = (v6 * v6)
                            v118 := GenS2{F0: 0}
                            _ = v118
                        } else {
                            switch "\n\t" {
                            case "hello":
                                var v119 float64 = v4
                                _ = v119
                            case "foo bar":
                                var v120 uint16 = 0
                                _ = v120
                            case "\n\t":
                                switch fi100 {
                                case 4294967295:
                                    v61 -= 764
                                case 598:
                                    _ = len(v1)
                                case -445:
                                    v6 += 0
                                }
                            default:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                        _ = v117
                    }
                }
            case -9223372036854775808:
                var v121 uint64 = 4294967295
                _ = v121
            default:
                v122 := GenS19{F0: 255}
                _ = v122
            }
        case 128:
            var v123 uint64 = 4294967295
            _ = v123
        default:
            v124 := GenS20{F0: 577, F1: fi100}
            _ = v124
        }
        _ = fi100
        _ = v114
    }
    v4 += 1.5
    v125 := GenS1{F0: 796, F1: 874, F2: 731}
    v126 := GenS18{F0: -363, F1: -2147483648}
    v6 -= (v6 - v6)
    _ = len(v2)
    switch (("hello" + ("\n\t" + "")) + "a") {
    case "a":
        v4 = float64(1e10)
    case "hello":
        if true {
            if true {
                if (!(v4 <= (-2.25 / v4))) {
                    switch "hello" {
                    case "a":
                        fmt.Println(v4, v61)
                    case "foo bar":
                        var v127 uint8 = 128
                        _ = v127
                    case "\n\t":
                        fmt.Println(v6, v4, v61)
                    }
                    v6 = v6
                } else {
                    v128 := v5[464]
                    v129 := GenS10{F0: -32768, F1: 0}
                    _ = v128
                    _ = v129
                }
                v61 = v61
                if (false || false) {
                    v6 = v6
                    for fi130 := 0; fi130 < 5; fi130++ {
                        var v131 int32 = 247
                        _ = fi130
                        _ = v131
                    }
                    var v132 bool = (!true)
                    _ = v132
                }
            }
        }
    case "\n\t":
        switch 4294967295 {
        case 32768:
            var v133 uint = v61
            _ = v133
        case 264:
            var v134 string = ((("hello" + "\n\t") + (((((((("\n\t" + (("a" + "") + "")) + ("hello" + ("a" + ("" + ("" + ("\n\t" + "foo bar")))))) + "\n\t") + "hello") + (("\n\t" + "foo bar") + (((("hello" + "hello") + "hello") + (("\n\t" + "hello") + "")) + (("foo bar" + "") + (("a" + ("hello" + "")) + ("hello" + "")))))) + "foo bar") + ("hello" + "")) + "\n\t")) + "foo bar")
            _ = v134
        default:
            v6 += (v6 | v6)
        }
    }
    var v135 uint32 = 48
    v6 = v6
    v61 -= v61
    v136 := GenS20{F0: 463, F1: 32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v135:", v135)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v125
    _ = v126
    _ = v135
    _ = v136
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
}
