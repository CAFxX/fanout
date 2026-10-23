package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int8; F1 int8; F2 uint }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int16; F1 int; F2 int64 }
type GenS9 struct { F0 uint16; F1 int64; F2 int16 }
type GenS10 struct { F0 uint; F1 int8; F2 string }
type GenS11 struct { F0 int32; F1 int16; F2 int }
type GenS12 struct { F0 uint16; F1 bool; F2 int }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int16; F1 int16 }
type GenS16 struct { F0 uint; F1 uint64 }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 int16; F1 bool }
type GenS20 struct { F0 int8; F1 string }
type GenS21 struct { F0 uint32; F1 uint8; F2 uint }
type GenS22 struct { F0 uint64 }
type GenS23 struct { F0 uint8 }
type GenS24 struct { F0 bool; F1 int32; F2 bool }
type GenS25 struct { F0 uint8; F1 bool }
type GenS26 struct { F0 int; F1 uint32; F2 bool }
type GenS27 struct { F0 int }

func main() {
    v1 := make(map[int]int)
    var v2 int64 = -93
    v3 := []uint32{575, 731, 42}
    var v4 uint64 = 4294967295
    var v5 int64 = v2
    var v6 int64 = (v5 << ((4294967295) & 63))
    var v7 float32 = (1e10 + ((0.1 / 3.14159) - (0.0 / 1.0)))
    delete(v1, -749)
    v3 = append(v3, 4294967295)
    if false {
        v8 := GenS1{F0: -569, F1: 65535}
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: 1}
    var v10 int8 = 100
    v3 = append(v3, 0)
    fmt.Println(v2, v7)
    v7 += float32((3.14159 - 1e10))
    switch (("" + "") + "foo bar") {
    case "":
        var v11 uint16 = 422
        _ = v11
    case "a":
        var v12 int8 = (v10 << ((427) & 63))
        _ = v12
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        if false {
            switch "a" {
            case "a":
                var v14 float64 = 0.0
                _ = v14
            case "\n\t":
                v15 := GenS2{F0: 2147483648, F1: 0, F2: 42}
                _ = v15
            case "hello":
                if false {
                    switch fi13 {
                    case 32767:
                        _ = len(v1)
                    case 9223372036854775807:
                        v4 += v4
                    }
                    for fi16 := 0; fi16 < 1; fi16++ {
                        v7 = -2.25
                        switch fi16 {
                        case -1:
                            if true {
                                if ((fi13 >> ((1) & 63)) < fi16) {
                                    v17 := v1[fi13]
                                    _ = v17
                                }
                                if false {
                                    v18 := GenS1{F0: 601, F1: 227}
                                    _ = v18
                                } else {
                                    var v19 bool = (v10 < int8(fi16))
                                    _ = v19
                                }
                            }
                        case -612:
                            switch 2147483647 {
                            case 128:
                                if true {
                                    var v20 bool = (((574 > 18446744073709551615) || (42 > 65535)) && false)
                                    var v21 string = "foo bar"
                                    if (true && (0 > uint8(v4))) {
                                        if (false || (!((!(!(v20 && (1000 == -1)))) || (v21 <= v21)))) {
                                            v5 = (v5 << ((255) & 63))
                                            if v20 {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                v5 -= v5
                                                if len(v3) > 0 {
                                                    v3[0] = 4294967295
                                                }
                                            }
                                        } else {
                                            if (!(255 != 2)) {
                                                switch fi13 {
                                                case 32767:
                                                    v1[fi16] = (fi16 % (((fi13 / (((^(fi13))) | 1))) | 1))
                                                default:
                                                    var v22 float64 = 0.0
                                                    _ = v22
                                                }
                                                if v20 {
                                                    v2 = v2
                                                }
                                            } else {
                                                var v23 float64 = 0.0
                                                _ = v23
                                            }
                                            v24 := GenS3{F0: (v10 | 127), F1: v10, F2: 434}
                                            if true {
                                                _ = len(v1)
                                                _ = len(v3)
                                                var v25 uint32 = 101
                                                switch fi16 {
                                                case 580:
                                                    _ = len(v1)
                                                    var v26 int32 = -246
                                                    _ = v26
                                                }
                                                _ = v25
                                            } else {
                                                if v20 {
                                                    v2 -= (v6 << ((255) & 63))
                                                    v27 := GenS4{F0: 1, F1: v20}
                                                    _ = v27
                                                } else {
                                                    if (272 <= 0) {
                                                        if (42 != 32767) {
                                                            v28 := GenS5{F0: 42}
                                                            v1[fi16] = fi13
                                                            fmt.Println(v10, v4)
                                                            _ = v28
                                                        }
                                                        if (v20 && true) {
                                                            switch v21 {
                                                            case "\n\t":
                                                                fmt.Println(v5)
                                                            case "foo bar":
                                                                v29 := GenS5{F0: 42}
                                                                var v30 bool = v20
                                                                if true {
                                                                    fmt.Println(fi13)
                                                                }
                                                                _ = v29
                                                                _ = v30
                                                            default:
                                                                var v31 uint32 = 2147483648
                                                                _ = v31
                                                            }
                                                        }
                                                        var v32 uint64 = (v4 % (((v4 | 0)) | 1))
                                                        _ = v32
                                                    }
                                                }
                                            }
                                            _ = v24
                                        }
                                        if (!(446 <= int16(v10))) {
                                            _ = len(v3)
                                        } else {
                                            var v33 string = (("" + (v21 + "\n\t")) + "a")
                                            _ = v33
                                        }
                                        if v20 {
                                            v2 -= (v2 >> ((65535) & 63))
                                            var v34 string = v21
                                            switch (fi16 ^ (fi13 ^ fi16)) {
                                            case 42:
                                                var v35 int64 = v2
                                                _ = v35
                                            case 9223372036854775807:
                                                switch (fi13 / ((-2147483648) | 1)) {
                                                case 460:
                                                    var v36 uint8 = 78
                                                    _ = v36
                                                }
                                            }
                                            _ = v34
                                        } else {
                                            v21 = "a"
                                            var v37 uint64 = 42
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = v37
                                        }
                                        v38 := GenS6{F0: v6, F1: int64(v4)}
                                        _ = v38
                                    } else {
                                        var v39 uint8 = 42
                                        var v40 uint8 = v39
                                        _ = v39
                                        _ = v40
                                    }
                                    _ = v20
                                    _ = v21
                                }
                            case 9223372036854775807:
                                var v41 float64 = (-2.25 + (float64(1.5) + (1.5 - ((0.0 * float64(v7)) * -2.25))))
                                _ = v41
                            default:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        case -566:
                            v4 = v4
                        }
                        if (true && false) {
                            v42 := GenS1{F0: -2147483648, F1: 32767}
                            v1[fi16] = (fi13 * 42)
                            if (-1 <= 0) {
                                switch (fi13 & fi16) {
                                case 9223372036854775807:
                                    v2 -= (v5 - -9223372036854775808)
                                default:
                                    fmt.Println(v7)
                                }
                                var v43 float64 = 0.0
                                _ = v43
                            }
                            v44 := GenS3{F0: 58, F1: (v10 << ((4294967295) & 63)), F2: 961}
                            _ = len(v1)
                            _ = v42
                            _ = v44
                        } else {
                            if false {
                                var v45 bool = false
                                var v46 uint32 = 619
                                fmt.Println(v45, v4)
                                _ = v45
                                _ = v46
                            } else {
                                var v47 float32 = float32((0.1 * (0.1 / 1.0)))
                                _ = len(v1)
                                v48 := v1[(fi16 + 9223372036854775807)]
                                _ = v47
                                _ = v48
                            }
                        }
                        _ = fi16
                    }
                    switch (^(fi13)) {
                    case 235:
                        var v49 int8 = v10
                        _ = v49
                    }
                }
            }
            if (v7 != v7) {
                for fi50 := 0; fi50 < 12; fi50++ {
                    v1[(fi13 - (fi50 >> ((18446744073709551615) & 63)))] = (fi50 << ((670) & 63))
                    if (2 >= 42) {
                        var v51 int = (-(fi50))
                        switch "foo bar" {
                        case "foo bar":
                            v2 = 647
                        case "a":
                            fmt.Println(v7)
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        if (true && false) {
                            fmt.Println(v51)
                            switch "hello" {
                            case "":
                                var v52 int16 = 1000
                                _ = v52
                            case "a":
                                var v53 float64 = (3.14159 + ((0.1 * 3.14159) + 3.14159))
                                _ = v53
                            default:
                                v7 -= (float32(-2.25) / float32((0.1 - (0.1 / (0.1 + float64((0.0 - float32(-2.25))))))))
                            }
                        }
                        _ = v51
                    } else {
                        fmt.Println(v5)
                        v10 -= v10
                        _ = len(v1)
                    }
                    _ = fi50
                }
                v1[fi13] = 1
                if false {
                    if (3.14159 > v7) {
                        var v54 int8 = (^(v10))
                        _ = v54
                    }
                    if true {
                        v55 := GenS7{F0: (fi13 * (fi13 & fi13))}
                        if len(v3) > 0 {
                            v3[0] = 2147483647
                        }
                        switch fi13 {
                        case 42:
                            var v56 int64 = v5
                            _ = v56
                        }
                        _ = v55
                    }
                }
            } else {
                v7 = (v7 - 3.14159)
                if len(v3) > 0 {
                    v3[0] = 830
                }
                if (true && (v5 != v2)) {
                    var v57 string = "hello"
                    switch v57 {
                    case "":
                        var v58 string = ("a" + ("foo bar" + v57))
                        _ = v58
                    case "hello":
                        var v59 int32 = int32(fi13)
                        _ = v59
                    case "foo bar":
                        fmt.Println(v2)
                    }
                    fmt.Println(v10)
                    _ = v57
                }
            }
            delete(v1, fi13)
        } else {
            if (-688 >= -1) {
                var v60 int = (fi13 * (^(fi13)))
                var v61 int8 = v10
                _ = v60
                _ = v61
            } else {
                for fi62 := 0; fi62 < 9; fi62++ {
                    v3 = append(v3, 2147483648)
                    if false {
                        if false {
                            fmt.Println(v6)
                        } else {
                            v5 -= int64(fi13)
                            if false {
                                v7 += (v7 / 1.0)
                                fmt.Println(v4, fi62, v7)
                                v2 = (^(v6))
                            } else {
                                var v63 float32 = 1e10
                                _ = v63
                            }
                            var v64 int8 = v10
                            _ = v64
                        }
                    }
                    _ = fi62
                }
            }
            var v65 float32 = v7
            _ = v65
        }
        _ = fi13
    }
    switch ("foo bar" + (("a" + ("" + "a")) + ("foo bar" + "\n\t"))) {
    case "hello":
        var v66 bool = ("\n\t" <= ("hello" + (((("a" + (("foo bar" + (("hello" + "foo bar") + "foo bar")) + "foo bar")) + ("a" + "")) + "") + "hello")))
        _ = v66
    case "":
        if true {
            v67 := GenS8{F0: 1, F1: 128, F2: v2}
            switch 1 {
            case -99:
                fmt.Println(v7, v5, v4)
            default:
                for fi68 := 0; fi68 < 3; fi68++ {
                    for fi69 := 0; fi69 < 5; fi69++ {
                        switch "foo bar" {
                        case "a":
                            var v70 int8 = 100
                            _ = v70
                        case "foo bar":
                            v6 -= (-(v2))
                        }
                        v5 -= 1
                        v7 += 1e10
                        _ = fi69
                    }
                    v7 += v7
                    _ = fi68
                }
            }
            _ = v67
        } else {
            for fi71 := 0; fi71 < 5; fi71++ {
                switch (^(fi71)) {
                case 839:
                    fmt.Println(v2)
                case 343:
                    if (false && true) {
                        v5 = v5
                        if ((3.14159 <= float32(1.5)) && ((("foo bar" + ("hello" + ("a" + ""))) + "foo bar") == ((("\n\t" + "a") + "foo bar") + (("foo bar" + ("" + ("hello" + "hello"))) + (((("hello" + ("\n\t" + (("a" + "") + ""))) + ("hello" + "hello")) + ("" + "foo bar")) + "a"))))) {
                            switch "" {
                            case "a":
                                if false {
                                    for fi72 := 0; fi72 < 9; fi72++ {
                                        if (false || (v6 != (v6 - v2))) {
                                            v4 += (v4 % (((v4 >> ((2) & 63))) | 1))
                                        } else {
                                            v4 = v4
                                            switch fi72 {
                                            case 9223372036854775807:
                                                v73 := GenS2{F0: 668, F1: 1, F2: 42}
                                                _ = v73
                                            case 577:
                                                switch (fi72 / ((fi72) | 1)) {
                                                case 2147483647:
                                                    var v74 float64 = 3.14159
                                                    _ = v74
                                                case 618:
                                                    if false {
                                                        switch fi71 {
                                                        case -9223372036854775808:
                                                            v75 := GenS6{F0: v5, F1: 9223372036854775807}
                                                            _ = v75
                                                        case 1:
                                                            var v76 uint64 = 0
                                                            _ = v76
                                                        }
                                                    }
                                                    switch (fi71 | 9223372036854775807) {
                                                    case 4294967295:
                                                        var v77 int64 = v2
                                                        _ = v77
                                                    case 42:
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            v78 := GenS9{F0: 2, F1: v2, F2: 0}
                                                            var v79 int64 = (-(v2))
                                                            _ = v78
                                                            _ = v79
                                                        case "foo bar":
                                                            fmt.Println(v10, v7, fi71)
                                                            if false {
                                                                fmt.Println(v7, v4, v2)
                                                                var v80 int32 = -1
                                                                v81 := GenS10{F0: 0, F1: (v10 & v10), F2: (("\n\t" + "foo bar") + "\n\t")}
                                                                _ = v80
                                                                _ = v81
                                                            } else {
                                                                if false {
                                                                    if (v7 > (0.1 + v7)) {
                                                                        if (42 <= 42) {
                                                                            v82 := GenS11{F0: -2147483648, F1: 1000, F2: fi72}
                                                                            switch (fi72 | fi71) {
                                                                            case 765:
                                                                                var v83 uint = 18446744073709551615
                                                                                _ = v83
                                                                            case -481:
                                                                                var v84 int8 = (-(v10))
                                                                                _ = v84
                                                                            case 4294967295:
                                                                                switch "" {
                                                                                case "hello":
                                                                                    var v85 uint32 = 585
                                                                                    v10 += (v10 * v10)
                                                                                    v10 += v10
                                                                                    fmt.Println(v5, v4, v10)
                                                                                    v6 = 42
                                                                                    _ = v85
                                                                                case "":
                                                                                    _ = len(v1)
                                                                                case "\n\t":
                                                                                    var v86 string = ("hello" + "hello")
                                                                                    _ = v86
                                                                                }
                                                                            default:
                                                                                v87 := GenS2{F0: 4294967295, F1: 766, F2: 65535}
                                                                                _ = v87
                                                                            }
                                                                            _ = v82
                                                                        } else {
                                                                            var v88 int64 = (v6 / (((v5 + (v6 * (v6 / (((v2 << ((884) & 63))) | 1))))) | 1))
                                                                            v6 = v6
                                                                            if false {
                                                                                v5 = v88
                                                                                var v89 int64 = (-(v88))
                                                                                _ = v89
                                                                            } else {
                                                                                if true {
                                                                                    v10 = (v10 % (((v10 & (v10 + (v10 * v10)))) | 1))
                                                                                    v2 -= (v6 / ((v2) | 1))
                                                                                } else {
                                                                                    var v90 uint32 = 213
                                                                                    v6 = v2
                                                                                    var v91 int32 = -1
                                                                                    _ = v90
                                                                                    _ = v91
                                                                                }
                                                                                var v92 string = "hello"
                                                                                if false {
                                                                                    v6 = v2
                                                                                    if (((0.1 * 1e10) < (-2.25 * 0.1)) && (1 == 0)) {
                                                                                        if false {
                                                                                            switch (fi71 - fi71) {
                                                                                            case 1:
                                                                                                var v93 int16 = 755
                                                                                                _ = v93
                                                                                            case 32768:
                                                                                                v94 := v1[-9223372036854775808]
                                                                                                _ = v94
                                                                                            }
                                                                                            if true {
                                                                                                if len(v3) >= 2 {
                                                                                                    v3 = v3[1:len(v3)]
                                                                                                }
                                                                                            } else {
                                                                                                var v95 uint32 = 42
                                                                                                switch 0 {
                                                                                                case 32767:
                                                                                                    v96 := GenS12{F0: 65535, F1: ((false || false) || (!true)), F2: fi71}
                                                                                                    _ = v96
                                                                                                case -211:
                                                                                                    v7 -= v7
                                                                                                default:
                                                                                                    fmt.Println(v88)
                                                                                                }
                                                                                                var v97 string = (v92 + v92)
                                                                                                fmt.Println(fi72, v5)
                                                                                                if true {
                                                                                                    var v98 uint16 = 65535
                                                                                                    var v99 float64 = 0.1
                                                                                                    _ = v98
                                                                                                    _ = v99
                                                                                                } else {
                                                                                                    switch fi72 {
                                                                                                    case -837:
                                                                                                        if len(v3) >= 2 {
                                                                                                            v3 = v3[1:len(v3)]
                                                                                                        }
                                                                                                    case -620:
                                                                                                        var v100 uint64 = (v4 | (v4 - (^(v4))))
                                                                                                        _ = v100
                                                                                                    case 32768:
                                                                                                        var v101 uint32 = (v95 & 2)
                                                                                                        _ = v101
                                                                                                    default:
                                                                                                        var v102 float32 = (-2.25 / -2.25)
                                                                                                        _ = v102
                                                                                                    }
                                                                                                    v1[(fi71 % ((fi71) | 1))] = (fi72 & (fi72 % (((-(fi72))) | 1)))
                                                                                                }
                                                                                                if true {
                                                                                                    if (218 > 2) {
                                                                                                        if len(v3) > 0 {
                                                                                                            v3[0] = v95
                                                                                                        }
                                                                                                        var v103 bool = false
                                                                                                        v6 = v2
                                                                                                        v5 = (v88 & (v6 + (v2 * 42)))
                                                                                                        v104 := GenS13{F0: 65535}
                                                                                                        _ = v103
                                                                                                        _ = v104
                                                                                                    }
                                                                                                    switch (("hello" + (v92 + v92)) + ((v97 + "") + v97)) {
                                                                                                    case "hello":
                                                                                                        if (!(v10 != (-(v10)))) {
                                                                                                            v97 = (("a" + ("" + "a")) + "")
                                                                                                        } else {
                                                                                                            v4 = (^(v4))
                                                                                                            v105 := GenS14{F0: 2}
                                                                                                            v106 := GenS7{F0: fi71}
                                                                                                            v10 = v10
                                                                                                            _ = v105
                                                                                                            _ = v106
                                                                                                        }
                                                                                                    case "a":
                                                                                                        var v107 string = v97
                                                                                                        _ = v107
                                                                                                    case "\n\t":
                                                                                                        var v108 uint = 18446744073709551615
                                                                                                        _ = v108
                                                                                                    default:
                                                                                                        v109 := GenS5{F0: 171}
                                                                                                        _ = v109
                                                                                                    }
                                                                                                } else {
                                                                                                    var v110 bool = (0 == 32767)
                                                                                                    if (0 != int16(v10)) {
                                                                                                        switch (-(fi72)) {
                                                                                                        case 32768:
                                                                                                            fmt.Println(v88)
                                                                                                        case -682:
                                                                                                            if (0 != 319) {
                                                                                                                v3 = append(v3, (^(v95)))
                                                                                                            }
                                                                                                        }
                                                                                                        v5 = (v6 ^ v6)
                                                                                                        var v111 int8 = 42
                                                                                                        _ = v111
                                                                                                    } else {
                                                                                                        fmt.Println(v6)
                                                                                                        v1[int(v95)] = (fi72 ^ -2147483648)
                                                                                                        v3 = append(v3, (v95 - 0))
                                                                                                    }
                                                                                                    _ = v110
                                                                                                }
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = 2
                                                                                                }
                                                                                                _ = v95
                                                                                                _ = v97
                                                                                            }
                                                                                        } else {
                                                                                            v6 += (v5 - 1)
                                                                                            var v112 uint32 = 144
                                                                                            v10 = (-(v10))
                                                                                            _ = v112
                                                                                        }
                                                                                        var v113 float64 = 1.5
                                                                                        _ = v113
                                                                                    }
                                                                                    if false {
                                                                                        fmt.Println(v6, v5)
                                                                                    }
                                                                                } else {
                                                                                    var v114 int = (fi71 | (^(fi71)))
                                                                                    _ = v114
                                                                                }
                                                                                _ = v92
                                                                            }
                                                                            _ = v88
                                                                        }
                                                                        switch "\n\t" {
                                                                        case "":
                                                                            var v115 int32 = -1
                                                                            _ = v115
                                                                        default:
                                                                            if (0 >= 255) {
                                                                                v116 := v1[fi72]
                                                                                fmt.Println(fi72)
                                                                                _ = v116
                                                                            } else {
                                                                                if (0.1 <= (3.14159 / (3.14159 + 0.1))) {
                                                                                    if (42 == 32768) {
                                                                                        if false {
                                                                                            v7 = v7
                                                                                            v117 := GenS15{F0: -271, F1: 42}
                                                                                            v4 += (^(v4))
                                                                                            _ = v117
                                                                                        } else {
                                                                                            v118 := GenS13{F0: 410}
                                                                                            v10 += 100
                                                                                            _ = v118
                                                                                        }
                                                                                        if false {
                                                                                            v1[0] = fi71
                                                                                        } else {
                                                                                            if (v2 <= v6) {
                                                                                                v4 = v4
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = 925
                                                                                                }
                                                                                                fmt.Println(fi71)
                                                                                            }
                                                                                            var v119 int32 = 1
                                                                                            switch (("\n\t" + "a") + "\n\t") {
                                                                                            case "\n\t":
                                                                                                var v120 uint32 = 2147483647
                                                                                                _ = v120
                                                                                            case "foo bar":
                                                                                                if ((v4 & 2) == 1) {
                                                                                                    var v121 uint64 = v4
                                                                                                    if len(v3) > 0 {
                                                                                                        v3[0] = 2147483647
                                                                                                    }
                                                                                                    _ = v121
                                                                                                }
                                                                                            }
                                                                                            _ = v119
                                                                                        }
                                                                                        v122 := GenS16{F0: 65535, F1: v4}
                                                                                        var v123 uint8 = 128
                                                                                        switch (("" + "a") + "") {
                                                                                        case "foo bar":
                                                                                            v2 += 1
                                                                                        case "":
                                                                                            var v124 int64 = v5
                                                                                            _ = v124
                                                                                        default:
                                                                                            v125 := GenS17{F0: 23}
                                                                                            _ = v125
                                                                                        }
                                                                                        _ = v122
                                                                                        _ = v123
                                                                                    }
                                                                                    var v126 int16 = 164
                                                                                    _ = v126
                                                                                }
                                                                                switch fi71 {
                                                                                case 0:
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                }
                                                                                _ = len(v1)
                                                                            }
                                                                        }
                                                                    }
                                                                    v127 := GenS8{F0: 1000, F1: fi72, F2: 0}
                                                                    fmt.Println(v10)
                                                                    _ = v127
                                                                } else {
                                                                    v4 = (^(v4))
                                                                    _ = len(v1)
                                                                    if true {
                                                                        v128 := GenS14{F0: 0}
                                                                        _ = v128
                                                                    }
                                                                }
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                switch (("a" + "hello") + (("a" + "\n\t") + "foo bar")) {
                                                                case "a":
                                                                    v7 = v7
                                                                case "foo bar":
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        v3 = append(v3, 4294967295)
                                                                    case "a":
                                                                        var v129 string = ((("" + "hello") + ("a" + "hello")) + "foo bar")
                                                                        _ = v129
                                                                    case "":
                                                                        if true {
                                                                            var v130 uint64 = v4
                                                                            v6 -= v2
                                                                            fmt.Println(v10, v6, fi71)
                                                                            _ = v130
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            if true {
                                                                v1[1] = 32768
                                                                v6 = (v2 << ((0) & 63))
                                                            } else {
                                                                fmt.Println(v5)
                                                                var v131 uint32 = 936
                                                                _ = v131
                                                            }
                                                        case "hello":
                                                            _ = len(v3)
                                                        }
                                                    case -706:
                                                        v132 := GenS8{F0: 766, F1: (fi71 << ((42) & 63)), F2: (v2 % ((839) | 1))}
                                                        _ = v132
                                                    }
                                                }
                                            case 32767:
                                                v133 := GenS6{F0: v5, F1: v2}
                                                var v134 int32 = 42
                                                _ = v133
                                                _ = v134
                                            }
                                        }
                                        switch "" {
                                        case "":
                                            v135 := GenS3{F0: (v10 % ((127) | 1)), F1: v10, F2: 1}
                                            _ = v135
                                        case "hello":
                                            var v136 float32 = (v7 / (float32(3.14159) + float32(1.5)))
                                            _ = v136
                                        }
                                        _ = fi72
                                    }
                                    for fi137 := 0; fi137 < 3; fi137++ {
                                        v138 := GenS9{F0: 181, F1: v2, F2: -1}
                                        _ = fi137
                                        _ = v138
                                    }
                                    var v139 uint64 = 4294967295
                                    _ = v139
                                }
                            }
                            v7 += ((float32(float64(v7)) * ((v7 / 1.5) - v7)) - 3.14159)
                        } else {
                            switch "a" {
                            case "\n\t":
                                fmt.Println(v10, v4, v6)
                            default:
                                fmt.Println(fi71)
                            }
                            if (fi71 <= (fi71 % ((fi71) | 1))) {
                                var v140 string = "foo bar"
                                v2 = -9223372036854775808
                                fmt.Println(v6, fi71)
                                _ = v140
                            }
                            for fi141 := 0; fi141 < 14; fi141++ {
                                if (0 == 924) {
                                    var v142 bool = false
                                    if v142 {
                                        fmt.Println(v4, v142)
                                        switch -401 {
                                        case 127:
                                            fmt.Println(v6, fi71)
                                        case 32768:
                                            if v142 {
                                                v2 += (v6 % ((v5) | 1))
                                                switch (fi71 / (((fi71 - fi141)) | 1)) {
                                                case 4294967295:
                                                    if v142 {
                                                        v3 = append(v3, 65535)
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            switch (fi141 ^ int(v5)) {
                                                            case 231:
                                                                var v143 bool = true
                                                                _ = v143
                                                            case -89:
                                                                fmt.Println(fi141, v7, v6)
                                                            }
                                                        case "foo bar":
                                                            switch (fi141 + 112) {
                                                            case -150:
                                                                v7 = float32(1e10)
                                                            case 9223372036854775807:
                                                                var v144 int = fi141
                                                                _ = v144
                                                            case 127:
                                                                if v142 {
                                                                    _ = len(v3)
                                                                    v145 := GenS18{F0: -253}
                                                                    _ = v145
                                                                } else {
                                                                    var v146 uint32 = 42
                                                                    _ = v146
                                                                }
                                                            default:
                                                                v5 = int64(v10)
                                                            }
                                                        }
                                                        if len(v3) > 0 {
                                                            v3[0] = 114
                                                        }
                                                    }
                                                case 2147483647:
                                                    v7 = (v7 + v7)
                                                    var v147 string = ""
                                                    _ = v147
                                                case -871:
                                                    v148 := GenS19{F0: -1, F1: true}
                                                    _ = v148
                                                }
                                            } else {
                                                var v149 int64 = 9223372036854775807
                                                _ = v149
                                            }
                                        }
                                        v150 := GenS20{F0: v10, F1: ""}
                                        fmt.Println(v10, v2)
                                        fmt.Println(v142, v6, v6)
                                        _ = v150
                                    }
                                    v151 := GenS21{F0: 714, F1: 255, F2: 230}
                                    _ = v142
                                    _ = v151
                                } else {
                                    if true {
                                        v4 = (v4 - (v4 & 461))
                                    }
                                    var v152 uint = 18446744073709551615
                                    _ = v152
                                }
                                _ = fi141
                            }
                        }
                        if (0 > 42) {
                            for fi153 := 0; fi153 < 18; fi153++ {
                                if true {
                                    if (!false) {
                                        fmt.Println(fi153, fi153)
                                    }
                                }
                                if (v7 >= v7) {
                                    var v154 uint32 = 42
                                    _ = v154
                                } else {
                                    if (v4 != 357) {
                                        v155 := GenS20{F0: v10, F1: ("a" + "hello")}
                                        _ = v155
                                    }
                                }
                                var v156 int8 = (v10 << ((65535) & 63))
                                _ = fi153
                                _ = v156
                            }
                        }
                    }
                case 32767:
                    fmt.Println(v4)
                }
                _ = fi71
            }
        }
    }
    if false {
        var v157 int32 = -109
        v158 := GenS14{F0: 373}
        _ = v157
        _ = v158
    } else {
        if ((!(24 < 255)) || (!(v4 < v4))) {
            var v159 int8 = -1
            if (386 >= 886) {
                var v160 uint8 = 120
                var v161 int32 = 42
                _ = v160
                _ = v161
            } else {
                var v162 float64 = -2.25
                if false {
                    v162 += v162
                    for fi163 := 0; fi163 < 17; fi163++ {
                        for fi164 := 0; fi164 < 6; fi164++ {
                            v4 += (v4 ^ (v4 | 4294967295))
                            if false {
                                var v165 int32 = 1
                                _ = v165
                            } else {
                                v166 := GenS22{F0: (v4 & v4)}
                                v2 += (^(v2))
                                _ = v166
                            }
                            _ = fi164
                        }
                        _ = fi163
                    }
                } else {
                    for fi167 := 0; fi167 < 14; fi167++ {
                        var v168 int8 = (v159 % ((127) | 1))
                        _ = fi167
                        _ = v168
                    }
                    v169 := GenS17{F0: 2147483648}
                    _ = v169
                }
                _ = v162
            }
            _ = v159
        } else {
            if (!true) {
                v5 -= v6
                v170 := GenS12{F0: 32767, F1: (835 == 42), F2: -1}
                if (0 > 42) {
                    if (32767 != -1) {
                        fmt.Println(v7, v6)
                        var v171 uint16 = 335
                        v171 += (v171 >> ((1) & 63))
                        _ = v171
                    } else {
                        if (true || (true && true)) {
                            _ = len(v3)
                            if false {
                                var v172 int16 = 1
                                _ = v172
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        } else {
                            for fi173 := 0; fi173 < 9; fi173++ {
                                if (!(true && true)) {
                                    var v174 float64 = 0.1
                                    v175 := GenS19{F0: 32767, F1: ((1 < 2) || (127 <= 255))}
                                    _ = v174
                                    _ = v175
                                } else {
                                    v2 = (v2 * v5)
                                }
                                _ = fi173
                            }
                        }
                    }
                    _ = len(v1)
                    v176 := GenS10{F0: 1, F1: 0, F2: "hello"}
                    _ = v176
                }
                _ = v170
            } else {
                fmt.Println(v7)
                if true {
                    if true {
                        _ = len(v1)
                    } else {
                        var v177 int64 = (v2 >> ((255) & 63))
                        _ = len(v3)
                        _ = v177
                    }
                    v10 -= v10
                    var v178 float32 = ((0.1 * 1e10) / float32(0.1))
                    _ = v178
                }
                if true {
                    v179 := GenS19{F0: -64, F1: false}
                    if false {
                        for fi180 := 0; fi180 < 9; fi180++ {
                            var v181 int = 778
                            _ = fi180
                            _ = v181
                        }
                        var v182 uint32 = 2147483648
                        var v183 uint = uint(v182)
                        _ = v182
                        _ = v183
                    } else {
                        if false {
                            v5 += 1
                            v184 := GenS13{F0: 0}
                            _ = v184
                        } else {
                            if true {
                                if false {
                                    v185 := GenS4{F0: v6, F1: (!false)}
                                    var v186 float32 = v7
                                    var v187 int16 = -32768
                                    _ = v185
                                    _ = v186
                                    _ = v187
                                } else {
                                    var v188 int = 42
                                    fmt.Println(v7, v4)
                                    _ = v188
                                }
                                v6 = v5
                                var v189 bool = (0.0 == (3.14159 - 0.0))
                                _ = v189
                            }
                            switch 127 {
                            case 127:
                                for fi190 := 0; fi190 < 3; fi190++ {
                                    if (!true) {
                                        for fi191 := 0; fi191 < 17; fi191++ {
                                            v192 := GenS23{F0: 2}
                                            _ = fi191
                                            _ = v192
                                        }
                                        v6 = v5
                                    } else {
                                        if (v2 == v5) {
                                            v5 = (v2 * (-(v2)))
                                            v193 := GenS24{F0: false, F1: 1, F2: true}
                                            _ = v193
                                        }
                                    }
                                    _ = fi190
                                }
                            }
                        }
                        _ = len(v1)
                    }
                    switch "" {
                    case "a":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "hello":
                        fmt.Println(v2, v10)
                    default:
                        if true {
                            if false {
                                if (42 <= -2147483648) {
                                    v7 = (v7 * v7)
                                    for fi194 := 0; fi194 < 20; fi194++ {
                                        v195 := GenS23{F0: 28}
                                        if (1 == 594) {
                                            if true {
                                                var v196 int32 = -2147483648
                                                _ = v196
                                            }
                                            for fi197 := 0; fi197 < 3; fi197++ {
                                                v7 = v7
                                                switch (fi197 >> ((65535) & 63)) {
                                                case 2147483647:
                                                    v198 := GenS25{F0: 1, F1: false}
                                                    _ = v198
                                                default:
                                                    v199 := v1[55]
                                                    _ = v199
                                                }
                                                fmt.Println(v2, v7)
                                                if len(v3) > 0 {
                                                    v3[0] = 2
                                                }
                                                v2 = (-(v2))
                                                _ = fi197
                                            }
                                            var v200 uint32 = 2
                                            _ = v200
                                        } else {
                                            var v201 float32 = (3.14159 * ((float32(0.1) - v7) * v7))
                                            v202 := GenS5{F0: 255}
                                            _ = v201
                                            _ = v202
                                        }
                                        _ = fi194
                                        _ = v195
                                    }
                                    delete(v1, 32767)
                                }
                                v203 := GenS23{F0: 220}
                                _ = v203
                            }
                        }
                    }
                    _ = v179
                }
            }
            for fi204 := 0; fi204 < 16; fi204++ {
                var v205 int16 = 1
                var v206 float64 = -2.25
                switch (fi204 % (((-(fi204))) | 1)) {
                case 4294967295:
                    fmt.Println(v6)
                case 2147483647:
                    for fi207 := 0; fi207 < 1; fi207++ {
                        var v208 uint32 = 2147483648
                        v2 += 9223372036854775807
                        switch fi204 {
                        case 32767:
                            switch (fi207 / (((fi204 / (((fi207 & (fi207 % (((-(fi204))) | 1)))) | 1))) | 1)) {
                            case 240:
                                var v209 int64 = 9223372036854775807
                                _ = v209
                            default:
                                switch "hello" {
                                case "hello":
                                    var v210 float32 = v7
                                    _ = v210
                                case "a":
                                    v211 := GenS9{F0: 381, F1: (v6 << ((uint(v208)) & 63)), F2: (v205 ^ v205)}
                                    _ = v211
                                default:
                                    delete(v1, 4294967295)
                                }
                            }
                        case 9223372036854775807:
                            if false {
                                var v212 int16 = (v205 & (v205 & v205))
                                _ = v212
                            } else {
                                switch 32768 {
                                case 0:
                                    var v213 int16 = int16(v10)
                                    if false {
                                        var v214 bool = true
                                        var v215 int8 = (v10 - 0)
                                        fmt.Println(v205)
                                        switch (fi204 + fi207) {
                                        case 42:
                                            v216 := GenS9{F0: 1, F1: v6, F2: 42}
                                            _ = v216
                                        case 350:
                                            var v217 int64 = v6
                                            _ = v217
                                        case 128:
                                            if len(v3) > 0 {
                                                v3[0] = v208
                                            }
                                        default:
                                            var v218 int16 = 1
                                            _ = v218
                                        }
                                        v10 -= -1
                                        _ = v214
                                        _ = v215
                                    } else {
                                        if ((326 == 65535) && (65535 > 18446744073709551615)) {
                                            v219 := GenS26{F0: 1, F1: v208, F2: true}
                                            var v220 float64 = 0.0
                                            var v221 string = "a"
                                            _ = v219
                                            _ = v220
                                            _ = v221
                                        } else {
                                            _ = len(v3)
                                        }
                                        var v222 float64 = (v206 + 1e10)
                                        _ = v222
                                    }
                                    _ = v213
                                case -1:
                                    v5 -= 0
                                default:
                                    v6 += (v2 + (v5 & -1))
                                }
                                fmt.Println(v208, v5)
                            }
                        case 127:
                            fmt.Println(v208)
                        default:
                            v223 := GenS27{F0: 0}
                            _ = v223
                        }
                        _ = fi207
                        _ = v208
                    }
                default:
                    v224 := GenS7{F0: fi204}
                    _ = v224
                }
                _ = fi204
                _ = v205
                _ = v206
            }
            v4 += v4
        }
        for fi225 := 0; fi225 < 8; fi225++ {
            fmt.Println(v7)
            _ = fi225
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
