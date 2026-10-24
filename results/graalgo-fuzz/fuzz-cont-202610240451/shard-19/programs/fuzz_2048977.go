package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16; F1 bool; F2 uint }
type GenS6 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS7 struct { F0 int; F1 int64 }
type GenS8 struct { F0 string; F1 string }
type GenS9 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS10 struct { F0 int32; F1 uint64; F2 bool }
type GenS11 struct { F0 bool; F1 uint32 }
type GenS12 struct { F0 int64; F1 string; F2 uint8 }
type GenS13 struct { F0 uint32 }

func main() {
    var v1 uint16 = 255
    var v2 float32 = float32(((0.1 / 1.0) * -2.25))
    v3 := []uint32{4294967295, 2147483648}
    var v4 int16 = -32768
    v5 := []int{-2147483648, 32767}
    v6 := GenS1{F0: 65535}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 -= (v1 ^ (v1 + (v1 * v1)))
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v9 int8 = 42
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi8
            _ = v9
        }
        fmt.Println(v2, v2)
        _ = fi7
    }
    if false {
        if false {
            var v10 int8 = -128
            _ = v10
        } else {
            var v11 float32 = v2
            if true {
                var v12 float32 = 3.14159
                fmt.Println(v2, v11, v11)
                if true {
                    if false {
                        var v13 uint64 = 0
                        if len(v5) > 0 {
                            v5[0] = 2147483647
                        }
                        var v14 int64 = 616
                        _ = v13
                        _ = v14
                    }
                } else {
                    for fi15 := 0; fi15 < 7; fi15++ {
                        switch "hello" {
                        case "\n\t":
                            switch 9223372036854775807 {
                            case -55:
                                var v16 float32 = v2
                                _ = v16
                            case 972:
                                switch "hello" {
                                case "hello":
                                    if false {
                                        v12 += v11
                                        if (!(!(42 <= 2))) {
                                            for fi17 := 0; fi17 < 15; fi17++ {
                                                var v18 int16 = (v4 / ((v4) | 1))
                                                var v19 int32 = 251
                                                v20 := GenS2{F0: 100}
                                                switch (fi17 & fi17) {
                                                case 42:
                                                    v21 := GenS3{F0: 2, F1: -128}
                                                    _ = v21
                                                case 0:
                                                    var v22 uint8 = 127
                                                    _ = v22
                                                case 2147483647:
                                                    if true {
                                                        if false {
                                                            if false {
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                var v23 uint32 = 65535
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                _ = v23
                                                            }
                                                            if ((true || (((^(v1)) > v1) || (true || ((3.14159 * (((1.5 - (((0.1 - 0.0) + (3.14159 * 3.14159)) / (0.1 + (1.5 - 0.0)))) - -2.25) - 0.0)) < 3.14159)))) && (((false && false) || ("a" <= (("a" + (("hello" + "a") + "\n\t")) + "foo bar"))) && false)) {
                                                                var v24 uint = 929
                                                                v25 := GenS2{F0: 127}
                                                                _ = v24
                                                                _ = v25
                                                            }
                                                            fmt.Println(v19)
                                                            switch 1 {
                                                            case -9223372036854775808:
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    v19 = 328
                                                                case "a":
                                                                    switch "a" {
                                                                    case "a":
                                                                        var v26 uint32 = 42
                                                                        _ = v26
                                                                    case "foo bar":
                                                                        fmt.Println(fi15, v19)
                                                                    }
                                                                    v4 = (v4 * (v18 * (v18 / (((v4 / ((v4) | 1))) | 1))))
                                                                    var v27 uint8 = 168
                                                                    _ = v27
                                                                default:
                                                                    var v28 int8 = -57
                                                                    _ = v28
                                                                }
                                                            case -879:
                                                                fmt.Println(v18, v12, v18)
                                                            case 42:
                                                                if (v18 != (^(v18))) {
                                                                    var v29 uint32 = 2
                                                                    v30 := GenS3{F0: 0, F1: 127}
                                                                    var v31 bool = false
                                                                    fmt.Println(v31, v1)
                                                                    var v32 uint32 = (^(v29))
                                                                    _ = v29
                                                                    _ = v30
                                                                    _ = v31
                                                                    _ = v32
                                                                } else {
                                                                    v5 = append(v5, fi17)
                                                                    v12 += -2.25
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 0
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                default:
                                                    _ = len(v3)
                                                }
                                                _ = fi17
                                                _ = v18
                                                _ = v19
                                                _ = v20
                                            }
                                            var v33 float32 = (v2 + v11)
                                            _ = v33
                                        }
                                        _ = len(v5)
                                    } else {
                                        if false {
                                            fmt.Println(v11, v2, v4)
                                            v12 = 1e10
                                        }
                                        var v34 int = 1
                                        v4 -= 42
                                        _ = v34
                                    }
                                case "":
                                    fmt.Println(v4, v2, v11)
                                case "\n\t":
                                    v11 = (v11 / v12)
                                }
                            }
                        case "a":
                            for fi35 := 0; fi35 < 15; fi35++ {
                                v11 -= v2
                                v36 := GenS4{F0: 1}
                                v37 := GenS4{F0: 18446744073709551615}
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                        default:
                            v5 = append(v5, (fi15 << ((65535) & 63)))
                        }
                        v38 := GenS5{F0: v1, F1: true, F2: 698}
                        var v39 uint64 = 666
                        _ = fi15
                        _ = v38
                        _ = v39
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v40 := GenS6{F0: (v1 * (v1 % (((v1 >> ((292) & 63))) | 1))), F1: v1, F2: -4}
                    _ = v40
                }
                _ = v12
            } else {
                v1 -= v1
                var v41 int32 = 929
                if (2147483647 <= 2147483648) {
                    fmt.Println(v1, v2)
                    for fi42 := 0; fi42 < 5; fi42++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v43 int32 = v41
                        if true {
                            v11 = v2
                            _ = len(v5)
                            v44 := GenS7{F0: (fi42 * (fi42 + 0)), F1: 42}
                            _ = v44
                        }
                        _ = fi42
                        _ = v43
                    }
                    if (false || ("" == ("hello" + "hello"))) {
                        if (126 != v1) {
                            for fi45 := 0; fi45 < 14; fi45++ {
                                var v46 uint = 18446744073709551615
                                _ = fi45
                                _ = v46
                            }
                            _ = len(v3)
                            fmt.Println(v1)
                        } else {
                            if true {
                                if (((255 == 4294967295) || false) && false) {
                                    switch -447 {
                                    case 2147483647:
                                        if ("hello" < "\n\t") {
                                            v47 := GenS4{F0: 18446744073709551615}
                                            var v48 uint16 = (v1 + 65535)
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = v47
                                            _ = v48
                                        }
                                    case -9223372036854775808:
                                        if len(v5) > 0 {
                                            v5[0] = 32767
                                        }
                                    }
                                    if true {
                                        v4 += 0
                                    } else {
                                        var v49 string = ""
                                        _ = v49
                                    }
                                    var v50 int = 32768
                                    _ = v50
                                } else {
                                    for fi51 := 0; fi51 < 10; fi51++ {
                                        switch ("a" + "\n\t") {
                                        case "hello":
                                            for fi52 := 0; fi52 < 10; fi52++ {
                                                fmt.Println(v11, v41)
                                                _ = fi52
                                            }
                                        case "":
                                            switch -2147483648 {
                                            case -815:
                                                var v53 int32 = v41
                                                _ = v53
                                            case -9223372036854775808:
                                                for fi54 := 0; fi54 < 1; fi54++ {
                                                    fmt.Println(v11, v11)
                                                    if true {
                                                        var v55 int32 = (v41 + (v41 << ((255) & 63)))
                                                        var v56 uint16 = 2
                                                        var v57 bool = (!(18446744073709551615 < 42))
                                                        if true {
                                                            v41 += 1
                                                            if (!((v55 + v55) != v55)) {
                                                                if (-108 <= -1) {
                                                                    v58 := GenS4{F0: 0}
                                                                    _ = v58
                                                                }
                                                                var v59 uint16 = 1
                                                                _ = v59
                                                            } else {
                                                                fmt.Println(v11)
                                                            }
                                                        } else {
                                                            var v60 string = "hello"
                                                            _ = v60
                                                        }
                                                        _ = v55
                                                        _ = v56
                                                        _ = v57
                                                    }
                                                    _ = fi54
                                                }
                                            case -2147483648:
                                                for fi61 := 0; fi61 < 4; fi61++ {
                                                    v2 = 1e10
                                                    _ = fi61
                                                }
                                            }
                                        }
                                        for fi62 := 0; fi62 < 16; fi62++ {
                                            var v63 uint8 = uint8(v1)
                                            switch (fi62 + (^(fi51))) {
                                            case 127:
                                                if (!true) {
                                                    var v64 int = (^(fi51))
                                                    v65 := GenS8{F0: (("\n\t" + "\n\t") + "\n\t"), F1: ""}
                                                    _ = v64
                                                    _ = v65
                                                } else {
                                                    v66 := GenS1{F0: 65535}
                                                    var v67 int = (fi62 + fi51)
                                                    fmt.Println(v67, v1)
                                                    _ = v66
                                                    _ = v67
                                                }
                                                v11 += v11
                                            case -1:
                                                _ = len(v3)
                                            default:
                                                v68 := GenS9{F0: -278, F1: (v1 << ((65535) & 63)), F2: 2147483647}
                                                _ = v68
                                            }
                                            var v69 uint16 = v1
                                            v69 += v69
                                            _ = fi62
                                            _ = v63
                                            _ = v69
                                        }
                                        _ = fi51
                                    }
                                    for fi70 := 0; fi70 < 12; fi70++ {
                                        for fi71 := 0; fi71 < 13; fi71++ {
                                            v11 += (v2 * v11)
                                            if (18446744073709551615 != 0) {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                switch fi71 {
                                                case -1:
                                                    v3 = append(v3, 65535)
                                                case -342:
                                                    fmt.Println(v4, fi71)
                                                default:
                                                    var v72 float32 = 1.5
                                                    _ = v72
                                                }
                                                switch fi70 {
                                                case 591:
                                                    if (0.0 < (((v11 / float32(0.1)) / v2) + 1.5)) {
                                                        fmt.Println(v1, v41, v11)
                                                    }
                                                case 9223372036854775807:
                                                    switch (fi71 - fi71) {
                                                    case 1:
                                                        var v73 float32 = v2
                                                        _ = v73
                                                    case -464:
                                                        v74 := GenS2{F0: 20}
                                                        v75 := GenS8{F0: "", F1: ""}
                                                        var v76 uint16 = v1
                                                        _ = v74
                                                        _ = v75
                                                        _ = v76
                                                    }
                                                }
                                            } else {
                                                _ = len(v3)
                                                if len(v3) > 0 {
                                                    v3[0] = 952
                                                }
                                                v3 = append(v3, 65535)
                                            }
                                            var v77 int8 = 0
                                            _ = fi71
                                            _ = v77
                                        }
                                        _ = fi70
                                    }
                                }
                                v41 = 42
                                if (false && true) {
                                    for fi78 := 0; fi78 < 16; fi78++ {
                                        _ = len(v5)
                                        var v79 int16 = v4
                                        switch "" {
                                        case "a":
                                            var v80 uint8 = 42
                                            _ = v80
                                        case "\n\t":
                                            if false {
                                                v81 := GenS6{F0: 604, F1: 593, F2: -58}
                                                switch ("\n\t" + ("foo bar" + ("" + ("foo bar" + "a")))) {
                                                case "foo bar":
                                                    v1 += v1
                                                case "hello":
                                                    if ((v1 - (v1 >> ((539) & 63))) != (v1 & 42)) {
                                                        v79 = (v79 + (-(v4)))
                                                        _ = len(v3)
                                                    } else {
                                                        var v82 bool = true
                                                        var v83 int32 = v41
                                                        v84 := GenS9{F0: (v4 << ((42) & 63)), F1: (^(v1)), F2: 1}
                                                        _ = v82
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                default:
                                                    v85 := GenS6{F0: (^(v1)), F1: (^(v1)), F2: -16}
                                                    _ = v85
                                                }
                                                _ = v81
                                            }
                                        default:
                                            if (int32(v1) < v41) {
                                                var v86 float32 = 0.0
                                                _ = v86
                                            } else {
                                                switch "a" {
                                                case "":
                                                    v87 := GenS10{F0: (-(v41)), F1: 905, F2: false}
                                                    _ = v87
                                                case "a":
                                                    if len(v3) > 0 {
                                                        v3[0] = 65535
                                                    }
                                                case "foo bar":
                                                    v88 := GenS2{F0: -73}
                                                    _ = v88
                                                }
                                                for fi89 := 0; fi89 < 9; fi89++ {
                                                    if true {
                                                        switch (^(fi78)) {
                                                        case -578:
                                                            switch fi89 {
                                                            case 42:
                                                                v90 := GenS10{F0: 2147483647, F1: uint64(v1), F2: true}
                                                                _ = v90
                                                            case 816:
                                                                v5 = append(v5, (^(fi89)))
                                                            }
                                                        case 9223372036854775807:
                                                            _ = len(v3)
                                                        case 42:
                                                            v79 -= 283
                                                        default:
                                                            var v91 int64 = -1
                                                            _ = v91
                                                        }
                                                    }
                                                    var v92 uint32 = 2147483647
                                                    _ = fi89
                                                    _ = v92
                                                }
                                            }
                                        }
                                        _ = fi78
                                        _ = v79
                                    }
                                    if ("\n\t" <= "hello") {
                                        v2 += v11
                                        var v93 int16 = (v4 ^ v4)
                                        _ = v93
                                    } else {
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v94 string = "foo bar"
                                            _ = v94
                                        default:
                                            for fi95 := 0; fi95 < 13; fi95++ {
                                                var v96 uint16 = 0
                                                _ = fi95
                                                _ = v96
                                            }
                                        }
                                    }
                                    var v97 float64 = -2.25
                                    _ = v97
                                } else {
                                    fmt.Println(v1, v41)
                                }
                            }
                            v98 := GenS11{F0: (false || true), F1: 2}
                            var v99 string = ("hello" + "foo bar")
                            _ = v98
                            _ = v99
                        }
                        v4 -= (v4 * 740)
                        for fi100 := 0; fi100 < 13; fi100++ {
                            var v101 float64 = ((1e10 * 3.14159) / -2.25)
                            _ = fi100
                            _ = v101
                        }
                    }
                } else {
                    if true {
                        var v102 int32 = 2147483647
                        _ = v102
                    }
                    fmt.Println(v11, v1, v4)
                }
                _ = v41
            }
            var v103 uint = 622
            _ = v103
            _ = v11
        }
    } else {
        v2 -= (v2 / v2)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v2 = ((3.14159 + v2) / ((v2 * (v2 * 1.5)) + float32((1e10 * 1.5))))
    }
    switch 42 {
    case 0:
        v104 := GenS12{F0: 0, F1: "", F2: 255}
        _ = v104
    case 2147483647:
        var v105 uint8 = 128
        _ = v105
    }
    fmt.Println(v1, v1, v2)
    switch 32767 {
    case 32768:
        v1 = (v1 - (v1 << ((0) & 63)))
    default:
        var v106 string = "foo bar"
        _ = v106
    }
    for fi107 := 0; fi107 < 15; fi107++ {
        var v108 uint64 = 42
        v1 = v1
        var v109 int32 = int32(v4)
        _ = fi107
        _ = v108
        _ = v109
    }
    switch ("foo bar" + "") {
    case "":
        v110 := GenS13{F0: 0}
        _ = v110
    case "foo bar":
        var v111 uint = 255
        _ = v111
    default:
        if true {
            v5 = append(v5, 175)
        }
    }
    v1 = (v1 ^ v1)
    var v112 int64 = 9223372036854775807
    if true {
        var v113 int = 42
        v1 += v1
        _ = v113
    }
    var v114 bool = true
    for fi115 := 0; fi115 < 6; fi115++ {
        v2 = ((float32(1.5) + (((1.5 * 1e10) + float32(-2.25)) + float32(1e10))) * v2)
        fmt.Println(v112, v114, v2)
        v2 -= (float32((1e10 * 1.5)) / (1.5 / v2))
        _ = fi115
    }
    fmt.Println("v1:", v1)
    fmt.Println("v112:", v112)
    fmt.Println("v114:", v114)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v112
    _ = v114
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
