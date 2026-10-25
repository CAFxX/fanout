package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 int16; F2 uint }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 int16; F1 uint; F2 int }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int16; F1 int16; F2 uint64 }
type GenS12 struct { F0 uint64; F1 int; F2 uint64 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 uint64; F1 uint; F2 int32 }

func main() {
    var v1 uint32 = 2147483647
    var v2 uint = 4294967295
    v3 := []uint32{(v1 ^ (^(v1))), 1, 1, v1}
    v4 := []int{32768, 32767, 0, -2147483648}
    var v5 int16 = 42
    v6 := []uint32{v1}
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            v5 -= (^(v5))
            _ = fi8
        }
        if len(v6) > 0 {
            v6[0] = (v1 | v1)
        }
        for fi9 := 0; fi9 < 4; fi9++ {
            switch (-(fi9)) {
            case 42:
                v10 := GenS1{F0: 32767}
                _ = len(v4)
                var v11 float64 = float64(0.1)
                v12 := GenS1{F0: v5}
                v11 -= float64(-2.25)
                _ = v10
                _ = v11
                _ = v12
            case 9223372036854775807:
                fmt.Println(v5, v5)
            case 1:
                fmt.Println(fi7)
            }
            v13 := GenS2{F0: 0}
            _ = fi9
            _ = v13
        }
        _ = fi7
    }
    v14 := GenS3{F0: 42, F1: (v5 >> ((v2) & 63)), F2: v2}
    v2 = (v2 | v2)
    var v15 float64 = 0.1
    switch "" {
    case "a":
        switch "a" {
        case "foo bar":
            v15 = (0.0 - (float64(0.1) + 3.14159))
        }
    }
    switch ("\n\t" + ("hello" + "hello")) {
    case "foo bar":
        var v16 uint64 = 1
        _ = v16
    case "a":
        switch -822 {
        case 4294967295:
            var v17 string = (("" + ("hello" + ("\n\t" + "foo bar"))) + ("\n\t" + ("\n\t" + ("\n\t" + "hello"))))
            _ = v17
        }
    default:
        if (570 > 266) {
            v18 := GenS3{F0: -650, F1: (v5 & (v5 & 32767)), F2: (v2 >> (((^(v2))) & 63))}
            _ = v18
        } else {
            var v19 int64 = -9223372036854775808
            switch 770 {
            case -2147483648:
                fmt.Println(v5, v1, v2)
            }
            _ = v19
        }
    }
    v20 := GenS4{F0: 1, F1: 1}
    if (true && false) {
        v3 = append(v3, v1)
        if ((!false) && false) {
            v2 = 913
            v21 := GenS2{F0: -1}
            _ = v21
        } else {
            switch "foo bar" {
            case "hello":
                v22 := GenS5{F0: 4294967295, F1: 1}
                _ = v22
            case "":
                if true {
                    switch "a" {
                    case "":
                        if true {
                            if (((100 > -99) && (true && false)) || false) {
                                for fi23 := 0; fi23 < 17; fi23++ {
                                    if false {
                                        v3 = append(v3, (v1 + v1))
                                    } else {
                                        v1 = v1
                                    }
                                    _ = fi23
                                }
                                switch ((("a" + "foo bar") + "") + "hello") {
                                case "\n\t":
                                    if (-2147483648 >= 2147483647) {
                                        var v24 string = "\n\t"
                                        if true {
                                            v24 = "foo bar"
                                            _ = len(v4)
                                            if true {
                                                v2 -= v2
                                            }
                                        } else {
                                            if (!(uint8(v1) != 42)) {
                                                if ((v24 + "\n\t") > v24) {
                                                    var v25 string = "\n\t"
                                                    _ = v25
                                                }
                                                if len(v6) > 0 {
                                                    v6[0] = v1
                                                }
                                            }
                                            if true {
                                                var v26 uint = v2
                                                switch "hello" {
                                                case "hello":
                                                    var v27 string = "a"
                                                    _ = v27
                                                case "foo bar":
                                                    fmt.Println(v5)
                                                case "a":
                                                    var v28 uint = 1
                                                    _ = v28
                                                default:
                                                    fmt.Println(v5)
                                                }
                                                _ = v26
                                            } else {
                                                for fi29 := 0; fi29 < 2; fi29++ {
                                                    var v30 int = 390
                                                    _ = fi29
                                                    _ = v30
                                                }
                                            }
                                        }
                                        var v31 int16 = 32767
                                        _ = v24
                                        _ = v31
                                    }
                                case "a":
                                    switch 42 {
                                    case 4294967295:
                                        if true {
                                            for fi32 := 0; fi32 < 11; fi32++ {
                                                switch 9223372036854775807 {
                                                case 32768:
                                                    v5 = -32768
                                                case 42:
                                                    var v33 int8 = 100
                                                    _ = v33
                                                }
                                                _ = fi32
                                            }
                                            var v34 int8 = 0
                                            _ = v34
                                        }
                                    }
                                case "foo bar":
                                    fmt.Println(v2, v1)
                                }
                                var v35 int16 = v5
                                _ = v35
                            } else {
                                for fi36 := 0; fi36 < 8; fi36++ {
                                    var v37 int = (fi36 + fi36)
                                    if true {
                                        v38 := GenS6{F0: 1, F1: ((1.5 / 1.0) != -2.25)}
                                        for fi39 := 0; fi39 < 2; fi39++ {
                                            var v40 string = ""
                                            _ = fi39
                                            _ = v40
                                        }
                                        for fi41 := 0; fi41 < 10; fi41++ {
                                            v1 = 651
                                            _ = fi41
                                        }
                                        _ = v38
                                    } else {
                                        v1 -= 2147483647
                                        var v42 uint16 = 166
                                        _ = v42
                                    }
                                    var v43 int8 = 0
                                    _ = fi36
                                    _ = v37
                                    _ = v43
                                }
                                switch -22 {
                                case 127:
                                    if true {
                                        var v44 int64 = -1
                                        for fi45 := 0; fi45 < 1; fi45++ {
                                            var v46 float64 = (v15 - 1.5)
                                            var v47 uint16 = 32767
                                            _ = fi45
                                            _ = v46
                                            _ = v47
                                        }
                                        _ = v44
                                    } else {
                                        var v48 bool = true
                                        _ = v48
                                    }
                                default:
                                    v49 := GenS7{F0: v5, F1: (v2 % ((65535) | 1)), F2: 32768}
                                    _ = v49
                                }
                            }
                        } else {
                            v15 -= v15
                        }
                    case "hello":
                        var v50 uint8 = 128
                        _ = v50
                    }
                    for fi51 := 0; fi51 < 13; fi51++ {
                        var v52 uint32 = (v1 >> (((v2 & v2)) & 63))
                        _ = fi51
                        _ = v52
                    }
                    for fi53 := 0; fi53 < 15; fi53++ {
                        v1 -= 0
                        v15 = -2.25
                        if true {
                            v1 = (v1 << ((v2) & 63))
                            v2 = v2
                            if true {
                                var v54 uint32 = 2147483648
                                _ = len(v6)
                                fmt.Println(v1, v2)
                                _ = v54
                            } else {
                                var v55 uint32 = (v1 >> (((v2 * v2)) & 63))
                                v1 -= v55
                                for fi56 := 0; fi56 < 6; fi56++ {
                                    var v57 uint = v2
                                    fmt.Println(v57)
                                    if true {
                                        v5 += (v5 % ((198) | 1))
                                    }
                                    _ = fi56
                                    _ = v57
                                }
                                _ = v55
                            }
                        } else {
                            var v58 uint8 = 53
                            for fi59 := 0; fi59 < 15; fi59++ {
                                v60 := GenS8{F0: (-2.25 != float32((v15 * (float64((-2.25 - 1e10)) + v15))))}
                                v58 += v58
                                if ((false && (v15 == 1e10)) || ((true || true) || (!(!((941 <= 0) && (v1 == uint32(v2))))))) {
                                    _ = len(v3)
                                    if len(v3) > 0 {
                                        v3[0] = v1
                                    }
                                } else {
                                    var v61 uint16 = 32767
                                    _ = v61
                                }
                                _ = fi59
                                _ = v60
                            }
                            v62 := GenS9{F0: (fi53 - 0)}
                            _ = v58
                            _ = v62
                        }
                        _ = fi53
                    }
                }
            }
            if (901 <= 42) {
                v5 = (-(v5))
                var v63 uint64 = 1
                _ = v63
            } else {
                var v64 float32 = -2.25
                v1 -= v1
                _ = v64
            }
            var v65 string = "a"
            _ = v65
        }
        v1 = (v1 * (v1 - v1))
    } else {
        fmt.Println(v15, v2, v5)
    }
    for fi66 := 0; fi66 < 5; fi66++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi66
    }
    switch -9223372036854775808 {
    case 625:
        if len(v3) > 0 {
            v3[0] = v1
        }
    case -848:
        if (!(0 < 42)) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v67 := GenS10{F0: -2147483648}
            v68 := GenS2{F0: -933}
            _ = v67
            _ = v68
        }
    case 2147483647:
        v5 = (v5 % (((-(v5))) | 1))
    default:
        for fi69 := 0; fi69 < 20; fi69++ {
            var v70 int16 = v5
            for fi71 := 0; fi71 < 5; fi71++ {
                if (true && true) {
                    if true {
                        v1 = 4294967295
                    } else {
                        fmt.Println(fi71, v1)
                        v6 = append(v6, (v1 | (^(v1))))
                        v5 = v5
                    }
                    v72 := GenS11{F0: v5, F1: 32767, F2: 1}
                    _ = v72
                }
                if (42 <= 100) {
                    v73 := GenS2{F0: 1}
                    v74 := GenS7{F0: (v70 + (v5 | v5)), F1: 0, F2: (fi71 & (fi69 - (fi69 >> ((v2) & 63))))}
                    switch (-(fi69)) {
                    case 9223372036854775807:
                        if len(v3) > 0 {
                            v3[0] = (v1 / (((v1 - (v1 + 2))) | 1))
                        }
                    case 127:
                        v70 += (v5 << (((v2 | (v2 >> ((v2) & 63)))) & 63))
                    default:
                        var v75 bool = (88 != 1)
                        _ = v75
                    }
                    var v76 int64 = -9223372036854775808
                    switch "a" {
                    case "hello":
                        v77 := GenS12{F0: 360, F1: fi69, F2: 4294967295}
                        _ = v77
                    case "":
                        var v78 uint8 = 1
                        _ = v78
                    case "\n\t":
                        if (0 >= (fi69 + (fi71 << ((v2) & 63)))) {
                            v2 = v2
                            var v79 int = -119
                            _ = v79
                        }
                    default:
                        if len(v3) > 0 {
                            v3[0] = (^(v1))
                        }
                    }
                    if (false || false) {
                        if false {
                            switch "\n\t" {
                            case "hello":
                                v70 = -1
                            case "\n\t":
                                v80 := GenS13{F0: 0}
                                _ = v80
                            default:
                                fmt.Println(v2, fi69, fi71)
                            }
                            var v81 int16 = v70
                            _ = v81
                        }
                        if false {
                            if len(v4) > 0 {
                                v4[0] = (-(fi69))
                            }
                            if (543 <= 751) {
                                fmt.Println(v76, fi69, fi69)
                            }
                        }
                    }
                    v82 := GenS8{F0: false}
                    v83 := GenS5{F0: 31, F1: 4294967295}
                    _ = v73
                    _ = v74
                    _ = v76
                    _ = v82
                    _ = v83
                }
                _ = fi71
            }
            _ = fi69
            _ = v70
        }
    }
    fmt.Println(v1, v2)
    if (-1 <= 0) {
        var v84 bool = (471 >= 4294967295)
        _ = v84
    } else {
        v5 += v5
        v1 = v1
    }
    v15 = 0.1
    if false {
        if ((^(v2)) > 217) {
            var v85 int = -619
            _ = v85
        } else {
            if true {
                for fi86 := 0; fi86 < 18; fi86++ {
                    fmt.Println(v5)
                    v87 := GenS2{F0: -943}
                    switch "a" {
                    case "hello":
                        if false {
                            fmt.Println(v15, fi86, v15)
                            switch ("\n\t" + "foo bar") {
                            case "":
                                switch fi86 {
                                case -282:
                                    if (417 < 1) {
                                        var v88 uint8 = 149
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v88
                                    }
                                case -878:
                                    if false {
                                        _ = len(v3)
                                        fmt.Println(v15, v15, v5)
                                        v5 = v5
                                    } else {
                                        var v89 uint32 = (v1 + (v1 * v1))
                                        _ = v89
                                    }
                                case 42:
                                    fmt.Println(v2)
                                default:
                                    v2 = v2
                                }
                            case "a":
                                for fi90 := 0; fi90 < 11; fi90++ {
                                    if (true && true) {
                                        v15 = v15
                                    }
                                    _ = fi90
                                }
                            case "foo bar":
                                v5 = v5
                            default:
                                if true {
                                    if false {
                                        var v91 int16 = 0
                                        _ = v91
                                    }
                                    switch fi86 {
                                    case 2147483647:
                                        v92 := GenS9{F0: (^(fi86))}
                                        _ = v92
                                    case -711:
                                        if false {
                                            fmt.Println(v5, v15)
                                        } else {
                                            var v93 uint32 = 0
                                            _ = v93
                                        }
                                    }
                                }
                            }
                        }
                    default:
                        var v94 uint16 = 261
                        _ = v94
                    }
                    _ = fi86
                    _ = v87
                }
            } else {
                _ = len(v6)
                if true {
                    var v95 float32 = float32((float64(1.5) + (v15 * (v15 + (1.5 / v15)))))
                    if (2 < v2) {
                        v4 = append(v4, 4294967295)
                    } else {
                        v5 -= (v5 / ((v5) | 1))
                    }
                    v2 += 504
                    _ = v95
                }
            }
        }
    } else {
        if false {
            fmt.Println(v5, v5)
            v1 = 2
        } else {
            var v96 string = "\n\t"
            var v97 bool = true
            _ = v96
            _ = v97
        }
        _ = len(v6)
        switch "hello" {
        case "":
            v98 := GenS8{F0: (!true)}
            _ = v98
        case "hello":
            v99 := GenS14{F0: v2}
            _ = v99
        case "\n\t":
            for fi100 := 0; fi100 < 12; fi100++ {
                for fi101 := 0; fi101 < 1; fi101++ {
                    var v102 int16 = v5
                    var v103 float64 = v15
                    if false {
                        var v104 string = ""
                        v1 += (v1 << ((v2) & 63))
                        _ = v104
                    } else {
                        v3 = append(v3, 0)
                        v105 := GenS15{F0: v5}
                        _ = v105
                    }
                    _ = fi101
                    _ = v102
                    _ = v103
                }
                fmt.Println(fi100, fi100)
                for fi106 := 0; fi106 < 18; fi106++ {
                    if len(v4) > 0 {
                        v4[0] = fi100
                    }
                    switch (fi100 >> (((v2 | v2)) & 63)) {
                    case 32767:
                        var v107 string = "a"
                        _ = v107
                    default:
                        var v108 uint64 = 4294967295
                        _ = v108
                    }
                    if false {
                        var v109 uint = (^(v2))
                        v4 = append(v4, fi106)
                        if (false && (-2147483648 != 42)) {
                            var v110 uint64 = 2
                            v15 = ((v15 + v15) + v15)
                            _ = v110
                        } else {
                            if true {
                                v111 := GenS3{F0: 9223372036854775807, F1: v5, F2: (v109 / ((0) | 1))}
                                var v112 uint32 = (^(v1))
                                var v113 uint = (v2 - 65535)
                                _ = v111
                                _ = v112
                                _ = v113
                            } else {
                                v114 := GenS16{F0: 0, F1: (v2 / ((v109) | 1)), F2: -994}
                                v4 = append(v4, (-(fi106)))
                                var v115 uint16 = 372
                                _ = v114
                                _ = v115
                            }
                            v2 = (v2 ^ v109)
                            fmt.Println(fi106, v2)
                        }
                        _ = v109
                    } else {
                        v116 := GenS5{F0: fi106, F1: v2}
                        var v117 int = (^(fi106))
                        _ = v116
                        _ = v117
                    }
                    fmt.Println(v5, fi106, v15)
                    _ = fi106
                }
                _ = fi100
            }
        default:
            fmt.Println(v2, v2, v5)
        }
    }
    fmt.Println(v15)
    v1 -= 519
    fmt.Println(v2, v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
