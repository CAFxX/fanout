package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int32; F2 int64 }
type GenS4 struct { F0 bool; F1 uint64; F2 int32 }
type GenS5 struct { F0 bool; F1 uint; F2 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int16; F1 int; F2 int32 }
type GenS8 struct { F0 uint32; F1 int }
type GenS9 struct { F0 uint64; F1 int }
type GenS10 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS11 struct { F0 bool; F1 uint; F2 int32 }
type GenS12 struct { F0 uint64; F1 uint; F2 bool }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    var v4 int32 = -256
    v5 := []uint32{2147483648, 1, 1, 0}
    if (18446744073709551615 <= 646) {
        fmt.Println(v4)
        var v6 uint32 = 42
        _ = len(v2)
        _ = v6
    }
    if true {
        v7 := GenS1{F0: true, F1: -9223372036854775808}
        switch "" {
        case "foo bar":
            var v8 int8 = 127
            _ = v8
        }
        _ = v7
    } else {
        var v9 uint = 18446744073709551615
        _ = v9
    }
    var v10 string = (("\n\t" + "foo bar") + ("hello" + "hello"))
    v11 := GenS1{F0: true, F1: 9223372036854775807}
    switch -983 {
    case -2147483648:
        if (9223372036854775807 >= -2147483648) {
            for fi12 := 0; fi12 < 3; fi12++ {
                var v13 uint16 = 255
                _ = fi12
                _ = v13
            }
            var v14 int64 = -9223372036854775808
            _ = v14
        }
    case -1:
        switch 752 {
        case 32768:
            v4 += v4
        case 703:
            var v15 uint16 = 255
            _ = v15
        case -2147483648:
            if len(v5) > 0 {
                v5[0] = 0
            }
        }
    case 128:
        v4 = -1
    default:
        v4 = 1
    }
    v16 := GenS2{F0: 2147483647, F1: 2, F2: 128}
    if (0 == 9223372036854775807) {
        v10 = "a"
        var v17 float32 = 1e10
        _ = v17
    } else {
        var v18 string = v10
        for fi19 := 0; fi19 < 17; fi19++ {
            fmt.Println(v10)
            v10 = (v10 + "foo bar")
            for fi20 := 0; fi20 < 7; fi20++ {
                fmt.Println(fi20)
                _ = fi20
            }
            _ = fi19
        }
        _ = v18
    }
    v4 = (v4 - v4)
    v21 := v1[262]
    var v22 float32 = -2.25
    switch ((("hello" + "hello") + "hello") + "\n\t") {
    case "foo bar":
        v23 := GenS2{F0: 0, F1: 65535, F2: 255}
        _ = v23
    case "":
        for fi24 := 0; fi24 < 15; fi24++ {
            switch (v10 + v10) {
            case "a":
                switch (("hello" + ((v10 + ((v10 + v10) + "a")) + v10)) + v10) {
                case "foo bar":
                    if true {
                        switch ((((v10 + (v10 + "")) + "a") + "a") + v10) {
                        case "hello":
                            var v25 uint32 = 2
                            _ = v25
                        case "":
                            if ((65535 < 18446744073709551615) && true) {
                                fmt.Println(v4, v10)
                            } else {
                                if true {
                                    v2[fi24] = v21
                                    var v26 int16 = 0
                                    fmt.Println(v22, v26)
                                    _ = v26
                                } else {
                                    var v27 float64 = 3.14159
                                    switch "foo bar" {
                                    case "foo bar":
                                        for fi28 := 0; fi28 < 20; fi28++ {
                                            v29 := GenS3{F0: 65535, F1: 1, F2: 42}
                                            if ((!(!false)) && true) {
                                                v5 = append(v5, 2)
                                                if (0 == 42) {
                                                    fmt.Println(v21, fi24, fi24)
                                                } else {
                                                    fmt.Println(fi28, v4, v22)
                                                    v4 = -309
                                                    v22 += (-2.25 - v22)
                                                }
                                            } else {
                                                v3[657] = "\n\t"
                                                v4 -= 0
                                                fmt.Println(fi24)
                                                v2[(fi24 % (((fi24 % ((fi24) | 1))) | 1))] = (^(fi28))
                                                if len(v5) > 0 {
                                                    v5[0] = 2147483647
                                                }
                                            }
                                            _ = fi28
                                            _ = v29
                                        }
                                    case "\n\t":
                                        v4 = v4
                                    default:
                                        var v30 uint64 = 1
                                        _ = v30
                                    }
                                    _ = v27
                                }
                                var v31 uint64 = 176
                                _ = v31
                            }
                        }
                        for fi32 := 0; fi32 < 11; fi32++ {
                            fmt.Println(fi24, v4, fi32)
                            _ = fi32
                        }
                        var v33 uint64 = 737
                        _ = v33
                    }
                case "\n\t":
                    if false {
                        switch 32767 {
                        case 128:
                            for fi34 := 0; fi34 < 14; fi34++ {
                                if true {
                                    if (false || (!false)) {
                                        if (2 == 1) {
                                            v22 = ((v22 + 0.0) / ((float32((0.0 + (3.14159 + 1.5))) * 0.1) / v22))
                                            if (v21 > v21) {
                                                if (v21 > (fi24 << ((255) & 63))) {
                                                    _ = len(v5)
                                                }
                                                if (32767 < fi24) {
                                                    v35 := GenS2{F0: (-(v21)), F1: 255, F2: 127}
                                                    _ = v35
                                                } else {
                                                    v4 = v4
                                                    var v36 uint64 = 42
                                                    var v37 uint64 = 1
                                                    _ = v36
                                                    _ = v37
                                                }
                                            } else {
                                                if true {
                                                    v10 = "\n\t"
                                                } else {
                                                    var v38 float32 = (v22 * (3.14159 * (3.14159 + (v22 * (float32(((1e10 + 1.5) / float64((v22 / (((v22 / 0.1) / v22) / (0.1 / 1.0)))))) * (-2.25 + v22))))))
                                                    if (255 <= 1) {
                                                        v38 = float32(((1.5 / (((1e10 * float64((v22 + ((v22 / -2.25) + (1.5 + float32(float64(v38))))))) + 0.1) / 3.14159)) * 1.5))
                                                    } else {
                                                        switch "" {
                                                        case "foo bar":
                                                            switch "" {
                                                            case "foo bar":
                                                                v39 := GenS4{F0: ((^(fi24)) == (fi34 + v21)), F1: 834, F2: 2147483647}
                                                                _ = v39
                                                            case "\n\t":
                                                                v40 := GenS5{F0: false, F1: 2, F2: 120}
                                                                _ = v40
                                                            }
                                                        case "":
                                                            v4 -= (v4 << ((65535) & 63))
                                                        case "hello":
                                                            v22 += 3.14159
                                                        default:
                                                            v1[352] = (fi34 % ((v21) | 1))
                                                            fmt.Println(fi34, fi34)
                                                            switch (v21 | (fi34 * v21)) {
                                                            case -2147483648:
                                                                v41 := v2[fi24]
                                                                _ = v41
                                                            case 751:
                                                                switch v10 {
                                                                case "":
                                                                    if false {
                                                                        v21 = -366
                                                                        if (!false) {
                                                                            _ = len(v2)
                                                                            switch -2147483648 {
                                                                            case 32768:
                                                                                switch (fi34 % ((1) | 1)) {
                                                                                case 32767:
                                                                                    v4 -= v4
                                                                                case 42:
                                                                                    v4 = -849
                                                                                default:
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = 377
                                                                                    }
                                                                                }
                                                                            default:
                                                                                v42 := GenS3{F0: 65535, F1: v4, F2: 1}
                                                                                _ = v42
                                                                            }
                                                                            _ = len(v5)
                                                                            if (997 >= 2) {
                                                                                v22 = ((v22 - float32(1e10)) + (v38 + -2.25))
                                                                                if false {
                                                                                    _ = len(v5)
                                                                                    fmt.Println(v38)
                                                                                } else {
                                                                                    _ = len(v5)
                                                                                    switch (-(fi24)) {
                                                                                    case 1:
                                                                                        var v43 int32 = 2147483647
                                                                                        _ = v43
                                                                                    case 2147483647:
                                                                                        if len(v5) >= 2 {
                                                                                            v5 = v5[1:len(v5)]
                                                                                        }
                                                                                    default:
                                                                                        v44 := GenS2{F0: (fi24 << ((255) & 63)), F1: 301, F2: 0}
                                                                                        _ = v44
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                v38 -= v22
                                                                                v45 := GenS6{F0: 381}
                                                                                _ = v45
                                                                            }
                                                                            v46 := GenS3{F0: 570, F1: (v4 | (v4 + (v4 - v4))), F2: -1}
                                                                            _ = v46
                                                                        }
                                                                        _ = len(v1)
                                                                    } else {
                                                                        fmt.Println(v10, v38)
                                                                    }
                                                                case "a":
                                                                    v21 = fi34
                                                                }
                                                            }
                                                        }
                                                    }
                                                    var v47 int64 = 9223372036854775807
                                                    _ = v38
                                                    _ = v47
                                                }
                                                var v48 int64 = 0
                                                v5 = append(v5, 129)
                                                _ = v48
                                            }
                                            v49 := GenS7{F0: 1, F1: (fi34 - (v21 * (-(fi24)))), F2: int32(fi34)}
                                            _ = v49
                                        } else {
                                            var v50 int8 = 81
                                            fmt.Println(v22)
                                            v4 = (v4 / (((v4 % ((v4) | 1))) | 1))
                                            _ = v50
                                        }
                                    }
                                    if ("foo bar" < (((v10 + (v10 + ("hello" + v10))) + v10) + v10)) {
                                        v22 -= float32(3.14159)
                                        var v51 int32 = v4
                                        if (42 < -1) {
                                            v22 -= float32(-2.25)
                                        } else {
                                            var v52 int16 = 1000
                                            fmt.Println(v22, v51, v52)
                                            _ = v52
                                        }
                                        _ = v51
                                    } else {
                                        switch (v21 / ((fi24) | 1)) {
                                        case -2147483648:
                                            v53 := GenS7{F0: 0, F1: (fi34 ^ (fi34 & 127)), F2: 42}
                                            _ = v53
                                        case -327:
                                            var v54 bool = false
                                            _ = v54
                                        default:
                                            var v55 float32 = v22
                                            _ = v55
                                        }
                                    }
                                } else {
                                    v21 += fi34
                                }
                                var v56 uint = 255
                                _ = fi34
                                _ = v56
                            }
                        case 42:
                            v21 = 4294967295
                        }
                        v57 := v1[4294967295]
                        v58 := GenS6{F0: 2}
                        _ = v57
                        _ = v58
                    } else {
                        v59 := GenS6{F0: 18446744073709551615}
                        if false {
                            if false {
                                if true {
                                    v10 = "foo bar"
                                    if (577 != 446) {
                                        if false {
                                            v4 = v4
                                            var v60 int8 = -128
                                            for fi61 := 0; fi61 < 12; fi61++ {
                                                v2[(v21 >> ((4294967295) & 63))] = fi61
                                                _ = fi61
                                            }
                                            _ = v60
                                        } else {
                                            if true {
                                                if (!(65535 == 1)) {
                                                    for fi62 := 0; fi62 < 3; fi62++ {
                                                        v4 = 1
                                                        if false {
                                                            v10 = "foo bar"
                                                            if false {
                                                                var v63 uint8 = 0
                                                                var v64 bool = false
                                                                v21 += (fi24 >> ((0) & 63))
                                                                v63 += (v63 << ((42) & 63))
                                                                var v65 uint = 811
                                                                _ = v63
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                            switch "a" {
                                                            case "foo bar":
                                                                _ = len(v5)
                                                            case "a":
                                                                v4 -= (v4 & (v4 + 1))
                                                                v66 := v1[390]
                                                                switch 948 {
                                                                case 0:
                                                                    if false {
                                                                        var v67 float64 = 0.0
                                                                        if (false || (100 != 42)) {
                                                                            delete(v1, 1)
                                                                            v68 := GenS6{F0: 18446744073709551615}
                                                                            _ = v68
                                                                        }
                                                                        var v69 int = (fi24 - (fi62 >> ((0) & 63)))
                                                                        _ = v67
                                                                        _ = v69
                                                                    }
                                                                case 32768:
                                                                    if len(v5) > 0 {
                                                                        v5[0] = 874
                                                                    }
                                                                default:
                                                                    if (!(!((v4 >> ((18446744073709551615) & 63)) <= v4))) {
                                                                        v70 := GenS7{F0: 0, F1: fi62, F2: 2147483647}
                                                                        v66 -= (-(v21))
                                                                        v22 += ((0.1 + (((0.0 + -2.25) - 3.14159) - v22)) + (v22 * -2.25))
                                                                        _ = v70
                                                                    } else {
                                                                        if false {
                                                                            switch (fi62 >> ((42) & 63)) {
                                                                            case 795:
                                                                                var v71 uint32 = 42
                                                                                if ((!true) && ((v4 > (v4 >> ((0) & 63))) || (-2.25 != (1.5 * -2.25)))) {
                                                                                    switch "\n\t" {
                                                                                    case "":
                                                                                        var v72 uint16 = 2
                                                                                        _ = v72
                                                                                    case "\n\t":
                                                                                        fmt.Println(fi62, v71)
                                                                                    case "a":
                                                                                        var v73 string = v10
                                                                                        _ = v73
                                                                                    }
                                                                                } else {
                                                                                    v22 -= v22
                                                                                    switch "" {
                                                                                    case "a":
                                                                                        fmt.Println(fi62, v21)
                                                                                    case "\n\t":
                                                                                        var v74 int32 = v4
                                                                                        _ = v74
                                                                                    }
                                                                                }
                                                                                _ = v71
                                                                            case -2147483648:
                                                                                v21 -= fi24
                                                                            }
                                                                            if true {
                                                                                switch v10 {
                                                                                case "a":
                                                                                    v10 = (v10 + "a")
                                                                                default:
                                                                                    var v75 float64 = float64((float32(float64(v22)) - v22))
                                                                                    _ = v75
                                                                                }
                                                                                var v76 uint8 = 1
                                                                                v22 += 0.0
                                                                                fmt.Println(v66, v66)
                                                                                v76 -= (^(v76))
                                                                                _ = v76
                                                                            }
                                                                            if (false && false) {
                                                                                fmt.Println(v66)
                                                                                if (0 < -1) {
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = 42
                                                                                    }
                                                                                }
                                                                                v77 := GenS8{F0: 1, F1: (v66 / ((15) | 1))}
                                                                                _ = v77
                                                                            } else {
                                                                                fmt.Println(fi62, fi24)
                                                                                _ = len(v3)
                                                                            }
                                                                            var v78 uint16 = 255
                                                                            fmt.Println(fi24, v78, v10)
                                                                            switch (fi62 + fi62) {
                                                                            case 42:
                                                                                var v79 uint8 = 0
                                                                                _ = v79
                                                                            case -2147483648:
                                                                                switch v21 {
                                                                                case 32767:
                                                                                    if false {
                                                                                        var v80 uint64 = 4294967295
                                                                                        if true {
                                                                                            if true {
                                                                                                switch v21 {
                                                                                                case 4294967295:
                                                                                                    v81 := GenS3{F0: 65535, F1: 116, F2: 0}
                                                                                                    _ = v81
                                                                                                case 128:
                                                                                                    v21 -= (^(fi62))
                                                                                                case 2147483647:
                                                                                                    if true {
                                                                                                        if (119 != 75) {
                                                                                                            v78 += 2
                                                                                                            if len(v5) >= 2 {
                                                                                                                v5 = v5[1:len(v5)]
                                                                                                            }
                                                                                                            if ((true || true) && (v4 <= 42)) {
                                                                                                                delete(v2, fi62)
                                                                                                            } else {
                                                                                                                fmt.Println(v21, v10)
                                                                                                                fmt.Println(v80)
                                                                                                                if len(v5) >= 2 {
                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        switch "a" {
                                                                                                        case "\n\t":
                                                                                                            if (!(2 != 1)) {
                                                                                                                if true {
                                                                                                                    if true {
                                                                                                                        v82 := GenS4{F0: true, F1: (v80 / (((v80 >> ((2) & 63))) | 1)), F2: (v4 ^ v4)}
                                                                                                                        _ = v82
                                                                                                                    } else {
                                                                                                                        v78 = v78
                                                                                                                        fmt.Println(v78, v80)
                                                                                                                        v78 = 738
                                                                                                                    }
                                                                                                                    if ("\n\t" != (v10 + v10)) {
                                                                                                                        var v83 uint8 = 18
                                                                                                                        if len(v5) > 0 {
                                                                                                                            v5[0] = 0
                                                                                                                        }
                                                                                                                        var v84 float64 = -2.25
                                                                                                                        _ = v83
                                                                                                                        _ = v84
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    switch (v21 << ((255) & 63)) {
                                                                                                                    case 32768:
                                                                                                                        v66 = fi62
                                                                                                                        var v85 bool = (v80 >= (v80 - v80))
                                                                                                                        var v86 uint16 = 0
                                                                                                                        _ = v85
                                                                                                                        _ = v86
                                                                                                                    case 127:
                                                                                                                        if len(v5) > 0 {
                                                                                                                            v5[0] = 498
                                                                                                                        }
                                                                                                                    }
                                                                                                                    fmt.Println(v78)
                                                                                                                    fmt.Println(v10)
                                                                                                                    fmt.Println(v22, v66)
                                                                                                                    if (!(!(18446744073709551615 == 986))) {
                                                                                                                        var v87 uint64 = 42
                                                                                                                        _ = v87
                                                                                                                    } else {
                                                                                                                        v88 := GenS7{F0: 602, F1: v66, F2: (v4 - (v4 % ((1) | 1)))}
                                                                                                                        var v89 uint8 = 0
                                                                                                                        v21 -= (v66 << ((1) & 63))
                                                                                                                        _ = v88
                                                                                                                        _ = v89
                                                                                                                    }
                                                                                                                    var v90 bool = (((true || false) && (0 == 4294967295)) || false)
                                                                                                                    _ = v90
                                                                                                                }
                                                                                                                var v91 int64 = 1
                                                                                                                v80 = (v80 - v80)
                                                                                                                _ = v91
                                                                                                            } else {
                                                                                                                v21 += fi62
                                                                                                                fmt.Println(v78)
                                                                                                                v66 -= (v21 / (((fi62 << ((255) & 63))) | 1))
                                                                                                            }
                                                                                                        case "a":
                                                                                                            var v92 uint8 = 255
                                                                                                            _ = v92
                                                                                                        default:
                                                                                                            if (v4 <= v4) {
                                                                                                                var v93 float32 = v22
                                                                                                                fmt.Println(v78, fi62)
                                                                                                                _ = v93
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    v94 := v2[v21]
                                                                                                    _ = v94
                                                                                                }
                                                                                            } else {
                                                                                                var v95 uint64 = v80
                                                                                                if false {
                                                                                                    v4 += (v4 * (v4 % (((v4 / ((v4) | 1))) | 1)))
                                                                                                    v3[1] = v10
                                                                                                    if true {
                                                                                                        v96 := GenS5{F0: (((v4 + (v4 ^ 0)) > 0) && (true && (!true))), F1: 756, F2: 606}
                                                                                                        var v97 int16 = 0
                                                                                                        _ = v96
                                                                                                        _ = v97
                                                                                                    }
                                                                                                } else {
                                                                                                    fmt.Println(fi62)
                                                                                                }
                                                                                                v66 += v66
                                                                                                _ = v95
                                                                                            }
                                                                                        } else {
                                                                                            var v98 uint64 = 151
                                                                                            var v99 string = v10
                                                                                            var v100 uint16 = v78
                                                                                            _ = v100
                                                                                            _ = v98
                                                                                            _ = v99
                                                                                        }
                                                                                        var v101 string = "hello"
                                                                                        _ = v101
                                                                                        _ = v80
                                                                                    }
                                                                                case 490:
                                                                                    v21 = 4294967295
                                                                                case 992:
                                                                                    if (1000 >= 32767) {
                                                                                        fmt.Println(v10, v10)
                                                                                        var v102 int64 = -302
                                                                                        var v103 string = ""
                                                                                        if true {
                                                                                            if true {
                                                                                                if true {
                                                                                                    v104 := v2[fi24]
                                                                                                    v66 -= (v104 + -51)
                                                                                                    _ = v104
                                                                                                }
                                                                                            } else {
                                                                                                fmt.Println(v103)
                                                                                                switch ((((v10 + ((v103 + "\n\t") + v103)) + "\n\t") + v10) + "\n\t") {
                                                                                                case "foo bar":
                                                                                                    if (2147483648 < 2147483648) {
                                                                                                        v105 := GenS4{F0: false, F1: 0, F2: v4}
                                                                                                        if ((0.0 - (0.0 * 3.14159)) < 0.0) {
                                                                                                            v78 += (v78 & (v78 * v78))
                                                                                                            if true {
                                                                                                                var v106 int16 = -32768
                                                                                                                if true {
                                                                                                                    var v107 int16 = (v106 / ((v106) | 1))
                                                                                                                    v108 := GenS7{F0: (v107 & v107), F1: v66, F2: v4}
                                                                                                                    switch v103 {
                                                                                                                    case "\n\t":
                                                                                                                        v107 = (v106 + (v107 + (v106 + v106)))
                                                                                                                    case "foo bar":
                                                                                                                        var v109 int8 = -1
                                                                                                                        _ = v109
                                                                                                                    case "":
                                                                                                                        switch fi24 {
                                                                                                                        case 356:
                                                                                                                            _ = len(v5)
                                                                                                                        case 898:
                                                                                                                            v21 -= v21
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v107
                                                                                                                    _ = v108
                                                                                                                } else {
                                                                                                                    v103 = "foo bar"
                                                                                                                    fmt.Println(v106)
                                                                                                                }
                                                                                                                var v110 int64 = v102
                                                                                                                _ = v106
                                                                                                                _ = v110
                                                                                                            } else {
                                                                                                                v103 = (("a" + v103) + (v10 + "foo bar"))
                                                                                                                if len(v5) > 0 {
                                                                                                                    v5[0] = 42
                                                                                                                }
                                                                                                                var v111 float64 = 3.14159
                                                                                                                _ = v111
                                                                                                            }
                                                                                                            var v112 uint = 285
                                                                                                            v10 = "hello"
                                                                                                            _ = v112
                                                                                                        }
                                                                                                        _ = v105
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    var v113 float64 = (((0.1 / ((1.5 / 1.5) / float64(((v22 + float32(((1e10 - 1e10) * (float64(v22) * 0.1)))) * (v22 * v22))))) * float64(1e10)) / -2.25)
                                                                                                    _ = v113
                                                                                                case "":
                                                                                                    v114 := GenS9{F0: uint64(v78), F1: v21}
                                                                                                    _ = v114
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = v102
                                                                                        _ = v103
                                                                                    } else {
                                                                                        v115 := GenS4{F0: (!(!false)), F1: 1, F2: v4}
                                                                                        _ = v115
                                                                                    }
                                                                                }
                                                                            case 32767:
                                                                                var v116 uint8 = 127
                                                                                _ = v116
                                                                            }
                                                                            if len(v5) > 0 {
                                                                                v5[0] = 42
                                                                            }
                                                                            if true {
                                                                                if (-2147483648 > v4) {
                                                                                    v117 := v2[(v21 & 2147483647)]
                                                                                    var v118 int64 = 884
                                                                                    v66 -= 1
                                                                                    _ = v117
                                                                                    _ = v118
                                                                                }
                                                                                switch v66 {
                                                                                case 2147483647:
                                                                                    var v119 bool = (!false)
                                                                                    _ = v119
                                                                                case 441:
                                                                                    v120 := GenS10{F0: 1, F1: (v78 / ((32767) | 1)), F2: 1}
                                                                                    _ = v120
                                                                                case 9223372036854775807:
                                                                                    v10 = ((v10 + "a") + "a")
                                                                                default:
                                                                                    v121 := GenS7{F0: 1000, F1: (v66 * fi62), F2: v4}
                                                                                    _ = v121
                                                                                }
                                                                            }
                                                                            var v122 float64 = (((-2.25 / (1e10 * (-2.25 + float64((float32(0.1) / 3.14159))))) / 3.14159) - -2.25)
                                                                            var v123 bool = false
                                                                            _ = v122
                                                                            _ = v123
                                                                            _ = v78
                                                                        }
                                                                    }
                                                                    fmt.Println(fi62, v21)
                                                                }
                                                                _ = v66
                                                            case "\n\t":
                                                                v21 = 42
                                                            }
                                                        } else {
                                                            fmt.Println(fi62)
                                                            v124 := GenS11{F0: (42 >= 18446744073709551615), F1: 569, F2: v4}
                                                            _ = v124
                                                        }
                                                        if true {
                                                            v21 += v21
                                                        }
                                                        if (!false) {
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            var v125 uint64 = 18446744073709551615
                                                            fmt.Println(v4)
                                                            fmt.Println(fi24)
                                                            v125 -= v125
                                                            v22 = v22
                                                            fmt.Println(fi24, fi62, v4)
                                                            _ = v125
                                                        }
                                                        _ = fi62
                                                    }
                                                    for fi126 := 0; fi126 < 18; fi126++ {
                                                        v5 = append(v5, 900)
                                                        switch (v21 * (-(fi126))) {
                                                        case 1:
                                                            v22 = v22
                                                        case -1:
                                                            var v127 uint = 255
                                                            var v128 int32 = 2147483647
                                                            var v129 float32 = v22
                                                            _ = len(v2)
                                                            var v130 int64 = 9223372036854775807
                                                            _ = v127
                                                            _ = v128
                                                            _ = v129
                                                            _ = v130
                                                        case 122:
                                                            var v131 float32 = v22
                                                            _ = v131
                                                        default:
                                                            v132 := GenS7{F0: 0, F1: v21, F2: (v4 >> ((315) & 63))}
                                                            _ = v132
                                                        }
                                                        var v133 float64 = -2.25
                                                        v4 = v4
                                                        _ = fi126
                                                        _ = v133
                                                    }
                                                }
                                            } else {
                                                switch 780 {
                                                case 32767:
                                                    if len(v5) > 0 {
                                                        v5[0] = 42
                                                    }
                                                case 9223372036854775807:
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                default:
                                                    for fi134 := 0; fi134 < 8; fi134++ {
                                                        fmt.Println(fi24)
                                                        _ = fi134
                                                    }
                                                }
                                                var v135 int = 32767
                                                _ = v135
                                            }
                                            v136 := GenS10{F0: 813, F1: 656, F2: 127}
                                            var v137 float32 = v22
                                            _ = v136
                                            _ = v137
                                        }
                                    }
                                } else {
                                    delete(v2, v21)
                                }
                                v22 = v22
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            switch (v21 - 9223372036854775807) {
                            case -2147483648:
                                switch (v21 & 2147483647) {
                                case 32767:
                                    var v138 bool = true
                                    _ = v138
                                case 9223372036854775807:
                                    v3[0] = v10
                                default:
                                    v139 := GenS12{F0: 2, F1: 4294967295, F2: false}
                                    _ = v139
                                }
                            }
                        } else {
                            for fi140 := 0; fi140 < 10; fi140++ {
                                var v141 uint16 = 255
                                v10 = v10
                                var v142 uint8 = 0
                                _ = fi140
                                _ = v141
                                _ = v142
                            }
                            v3[65535] = v10
                        }
                        _ = v59
                    }
                }
            case "":
                fmt.Println(v10, v22)
            case "foo bar":
                fmt.Println(fi24, v4, fi24)
            default:
                var v143 int32 = (v4 - (-(v4)))
                _ = v143
            }
            fmt.Println(fi24, fi24, v10)
            if ((!(255 < 255)) || ((v22 * v22) > v22)) {
                v21 -= (^(fi24))
                if len(v5) > 0 {
                    v5[0] = 1
                }
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            _ = fi24
        }
    case "a":
        var v144 uint = 255
        _ = v144
    }
    v145 := GenS4{F0: ((2 <= 128) || (1.5 != 1.5)), F1: 0, F2: v4}
    var v146 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v146:", v146)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v145
    _ = v146
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
