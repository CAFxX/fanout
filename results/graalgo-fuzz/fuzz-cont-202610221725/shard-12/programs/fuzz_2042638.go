package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 int8; F1 int16; F2 bool }
type GenS6 struct { F0 uint8; F1 int16; F2 string }
type GenS7 struct { F0 uint64; F1 int16 }
type GenS8 struct { F0 int16; F1 uint }
type GenS9 struct { F0 uint64; F1 int32 }
type GenS10 struct { F0 int32; F1 int; F2 int8 }
type GenS11 struct { F0 int8; F1 bool; F2 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 float32 = (float32(0.0) - 0.0)
    var v3 uint = 2
    v4 := make(map[uint32]uint32)
    var v5 int64 = 9223372036854775807
    var v6 string = ""
    var v7 uint64 = 4294967295
    if false {
        if false {
            if (v7 != v7) {
                fmt.Println(v3, v7)
                v8 := []int{2147483647, 2147483647}
                v2 += 3.14159
                _ = v8
            }
        }
    } else {
        delete(v4, 2147483647)
        v9 := []int{42, -9223372036854775808, 1}
        fmt.Println(v3, v2)
        _ = v9
    }
    if (!true) {
        fmt.Println(v5, v7, v3)
        v10 := []int{42, 9223372036854775807, 1, 2147483647}
        _ = v10
    }
    fmt.Println(v5, v6, v7)
    if (true && (!false)) {
        if ((-246 <= -1) && (!(((^(v3)) > v3) || ((!true) && true)))) {
            var v11 float64 = (0.1 * (1e10 - 1.5))
            if ((!(true || (-688 >= -1))) && ((762 <= 2147483647) && false)) {
                if (!(279 >= 193)) {
                    var v12 uint16 = 752
                    switch "foo bar" {
                    case "\n\t":
                        v13 := []int{}
                        _ = v13
                    case "hello":
                        var v14 string = v6
                        _ = v14
                    }
                    fmt.Println(v11)
                    _ = v12
                } else {
                    var v15 uint8 = 255
                    delete(v1, 32767)
                    for fi16 := 0; fi16 < 9; fi16++ {
                        fmt.Println(v2, v2)
                        v17 := []int{fi16, 855, fi16, fi16}
                        if (!(!true)) {
                            if len(v17) > 0 {
                                v17[0] = (fi16 % ((fi16) | 1))
                            }
                            switch 32767 {
                            case 2147483647:
                                if ((42 == 2) || true) {
                                    if ((42 <= -1) || false) {
                                        var v18 int = 128
                                        _ = v18
                                    } else {
                                        if ((0.1 == 0.0) || ("" > v6)) {
                                            for fi19 := 0; fi19 < 12; fi19++ {
                                                if len(v17) > 0 {
                                                    v17[0] = 4294967295
                                                }
                                                v20 := GenS1{F0: (v15 >> ((uint(v7)) & 63))}
                                                _ = fi19
                                                _ = v20
                                            }
                                            v17 = append(v17, fi16)
                                        }
                                        v21 := GenS1{F0: (v15 % ((v15) | 1))}
                                        if (65535 < 537) {
                                            var v22 uint32 = 668
                                            var v23 int8 = -116
                                            _ = v22
                                            _ = v23
                                        } else {
                                            switch (fi16 | (-(fi16))) {
                                            case 32768:
                                                var v24 int8 = 72
                                                _ = v24
                                            case 260:
                                                for fi25 := 0; fi25 < 8; fi25++ {
                                                    switch ("a" + v6) {
                                                    case "":
                                                        if false {
                                                            v5 = (v5 ^ v5)
                                                            v1[(^(fi25))] = fi16
                                                            v3 = v3
                                                        } else {
                                                            if (!true) {
                                                                v26 := GenS2{F0: 1, F1: v6}
                                                                var v27 int16 = 32767
                                                                fmt.Println(v27, v15)
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    var v28 uint32 = 4294967295
                                                                    _ = v28
                                                                case "":
                                                                    fmt.Println(v15, v5)
                                                                    switch v6 {
                                                                    case "foo bar":
                                                                        var v29 uint32 = 2
                                                                        _ = v29
                                                                    case "hello":
                                                                        v7 += (v7 | (^(v7)))
                                                                    }
                                                                case "hello":
                                                                    v27 += v27
                                                                default:
                                                                    if len(v17) >= 2 {
                                                                        v17 = v17[1:len(v17)]
                                                                    }
                                                                }
                                                                _ = v26
                                                                _ = v27
                                                            }
                                                            _ = len(v1)
                                                            switch (v6 + v6) {
                                                            case "a":
                                                                v1[(fi25 << ((v3) & 63))] = fi16
                                                            case "\n\t":
                                                                fmt.Println(v6, v2, v3)
                                                            }
                                                            v3 -= (v3 - (v3 - v3))
                                                        }
                                                    case "foo bar":
                                                        var v30 float64 = (0.1 / 0.1)
                                                        _ = v30
                                                    }
                                                    fmt.Println(v5, v7, v3)
                                                    switch (fi16 / (((fi25 >> (((v3 & 971)) & 63))) | 1)) {
                                                    case 32768:
                                                        _ = len(v17)
                                                    case 0:
                                                        v6 = (v6 + v6)
                                                        v11 = v11
                                                        var v31 uint64 = (^(v7))
                                                        if ((v15 < 251) && (42 != uint16(v31))) {
                                                            delete(v1, int(v3))
                                                            var v32 int32 = -1
                                                            var v33 int64 = v5
                                                            _ = v32
                                                            _ = v33
                                                        } else {
                                                            if false {
                                                                v1[(fi16 * -2147483648)] = fi25
                                                                v34 := GenS1{F0: (v15 ^ (v15 * (v15 - 42)))}
                                                                var v35 int8 = 118
                                                                _ = v34
                                                                _ = v35
                                                            } else {
                                                                if false {
                                                                    fmt.Println(v2, v2, fi25)
                                                                    var v36 int8 = 100
                                                                    switch (fi16 >> (((v3 / (((v3 << (((v3 << ((v3) & 63))) & 63))) | 1))) & 63)) {
                                                                    case 127:
                                                                        if true {
                                                                            if true {
                                                                                if (!false) {
                                                                                    if true {
                                                                                        v7 = v31
                                                                                    }
                                                                                    var v37 float64 = float64(((3.14159 / float32(v11)) - float32((1.5 * float64(v2)))))
                                                                                    fmt.Println(v31, v5, v3)
                                                                                    _ = v37
                                                                                }
                                                                                if true {
                                                                                    v1[fi25] = 32767
                                                                                    v38 := v4[568]
                                                                                    switch (fi25 >> (((v3 - 18446744073709551615)) & 63)) {
                                                                                    case 2147483647:
                                                                                        _ = len(v17)
                                                                                    case -1:
                                                                                        v39 := GenS2{F0: v38, F1: ((v6 + "\n\t") + ("foo bar" + "hello"))}
                                                                                        var v40 int32 = 2147483647
                                                                                        _ = v39
                                                                                        _ = v40
                                                                                    }
                                                                                    _ = v38
                                                                                } else {
                                                                                    var v41 float64 = ((v11 - (v11 - (v11 / 1e10))) - v11)
                                                                                    if true {
                                                                                        v41 = v41
                                                                                    }
                                                                                    var v42 float32 = -2.25
                                                                                    _ = v41
                                                                                    _ = v42
                                                                                }
                                                                            }
                                                                            v3 -= (^(v3))
                                                                            fmt.Println(v3, v36)
                                                                        }
                                                                    default:
                                                                        if true {
                                                                            if (!false) {
                                                                                v1[(fi16 ^ fi25)] = fi16
                                                                                var v43 bool = false
                                                                                _ = v43
                                                                            }
                                                                        } else {
                                                                            if len(v17) > 0 {
                                                                                v17[0] = fi16
                                                                            }
                                                                            if ((42 <= 42) && false) {
                                                                                var v44 int16 = 408
                                                                                var v45 string = (("" + "\n\t") + v6)
                                                                                fmt.Println(v15, fi25, v11)
                                                                                if len(v17) > 0 {
                                                                                    v17[0] = (fi25 - (fi16 / (((fi25 ^ (fi16 - (fi25 * 2147483647)))) | 1)))
                                                                                }
                                                                                _ = v44
                                                                                _ = v45
                                                                            } else {
                                                                                fmt.Println(v36, v6)
                                                                                v31 = (v31 * 18446744073709551615)
                                                                                v6 = ("" + (v6 + v6))
                                                                            }
                                                                            fmt.Println(v5, v3, fi25)
                                                                        }
                                                                    }
                                                                    _ = v36
                                                                }
                                                                v46 := GenS1{F0: v15}
                                                                fmt.Println(fi25, v3)
                                                                _ = v46
                                                            }
                                                            var v47 int8 = -2
                                                            if (v3 >= (v3 ^ (^(v3)))) {
                                                                if (!(true || (-1 == -1))) {
                                                                    _ = len(v4)
                                                                    v2 = (v2 - v2)
                                                                    v48 := GenS3{F0: (^(fi25))}
                                                                    if (((^(v7)) < (v7 << ((v3) & 63))) || true) {
                                                                        var v49 int32 = 1
                                                                        v50 := GenS2{F0: 2147483647, F1: v6}
                                                                        v11 = (v11 / v11)
                                                                        _ = v49
                                                                        _ = v50
                                                                    } else {
                                                                        fmt.Println(v15, v7)
                                                                    }
                                                                    _ = v48
                                                                }
                                                                v3 = (v3 >> ((2) & 63))
                                                                v2 = v2
                                                            } else {
                                                                v11 += 1e10
                                                                v51 := GenS4{F0: 65535, F1: 74}
                                                                var v52 bool = (!(-2147483648 != -1))
                                                                _ = v51
                                                                _ = v52
                                                            }
                                                            switch fi16 {
                                                            case 0:
                                                                v53 := v1[9223372036854775807]
                                                                _ = v53
                                                            case -882:
                                                                fmt.Println(v31, v11, v7)
                                                            }
                                                            if len(v17) > 0 {
                                                                v17[0] = fi16
                                                            }
                                                            _ = v47
                                                        }
                                                        _ = v31
                                                    case 128:
                                                        if true {
                                                            var v54 int64 = v5
                                                            v54 -= v5
                                                            var v55 uint64 = (v7 << ((v3) & 63))
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                    default:
                                                        v15 -= uint8(v7)
                                                    }
                                                    _ = fi25
                                                }
                                            case 4294967295:
                                                fmt.Println(v5, v2, v3)
                                            }
                                            v17 = append(v17, fi16)
                                            if ((^(v3)) < 255) {
                                                v5 = (-(v5))
                                                if true {
                                                    for fi56 := 0; fi56 < 6; fi56++ {
                                                        v57 := GenS4{F0: 0, F1: -984}
                                                        fmt.Println(v7, v6)
                                                        v2 = v2
                                                        _ = fi56
                                                        _ = v57
                                                    }
                                                    v5 -= -974
                                                }
                                                v58 := GenS2{F0: 2147483647, F1: "\n\t"}
                                                _ = v58
                                            } else {
                                                if false {
                                                    for fi59 := 0; fi59 < 8; fi59++ {
                                                        fmt.Println(v5, v6, fi16)
                                                        var v60 int8 = 42
                                                        delete(v4, 4294967295)
                                                        _ = fi59
                                                        _ = v60
                                                    }
                                                    v11 = -2.25
                                                    v61 := GenS5{F0: -68, F1: int16(v3), F2: ((v3 > (v3 / (((v3 | 1)) | 1))) && true)}
                                                    _ = v61
                                                }
                                                v62 := GenS3{F0: fi16}
                                                _ = v62
                                            }
                                        }
                                        _ = v21
                                    }
                                    if len(v17) > 0 {
                                        v17[0] = (-(fi16))
                                    }
                                }
                            default:
                                v63 := GenS5{F0: -128, F1: 42, F2: false}
                                _ = v63
                            }
                        } else {
                            var v64 uint32 = 0
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                            for fi65 := 0; fi65 < 9; fi65++ {
                                var v66 uint = (v3 >> ((v3) & 63))
                                _ = fi65
                                _ = v66
                            }
                            _ = v64
                        }
                        _ = fi16
                        _ = v17
                    }
                    _ = v15
                }
                v67 := GenS5{F0: 1, F1: -32768, F2: true}
                _ = v67
            } else {
                switch (v6 + "a") {
                case "hello":
                    var v68 int8 = -110
                    _ = v68
                default:
                    var v69 uint16 = 42
                    _ = v69
                }
            }
            _ = v11
        } else {
            v70 := []int{4294967295}
            var v71 int16 = 42
            v7 = (v7 ^ (^(v7)))
            _ = v70
            _ = v71
        }
    } else {
        for fi72 := 0; fi72 < 11; fi72++ {
            v3 = 255
            _ = fi72
        }
    }
    v7 = (v7 & (v7 / ((42) | 1)))
    v3 = (v3 >> ((v3) & 63))
    if false {
        for fi73 := 0; fi73 < 15; fi73++ {
            for fi74 := 0; fi74 < 7; fi74++ {
                switch fi74 {
                case -2147483648:
                    v5 -= (v5 * (v5 >> ((892) & 63)))
                case 2147483647:
                    v6 = v6
                case 42:
                    v75 := []int{}
                    _ = v75
                }
                if false {
                    var v76 uint32 = 343
                    v2 = ((3.14159 / v2) + 3.14159)
                    _ = v76
                }
                delete(v1, 198)
                _ = fi74
            }
            _ = fi73
        }
        fmt.Println(v3)
    }
    if false {
        switch "hello" {
        case "\n\t":
            if (!(-2.25 > 0.0)) {
                fmt.Println(v5)
                v6 = ((v6 + (v6 + v6)) + (v6 + v6))
                v5 = -892
            }
        default:
            var v77 int = 513
            _ = v77
        }
        switch -9223372036854775808 {
        case -2147483648:
            v78 := GenS1{F0: 41}
            _ = v78
        case 837:
            v79 := GenS6{F0: 1, F1: 1000, F2: v6}
            _ = v79
        case 128:
            if true {
                v2 += float32((-2.25 + (-2.25 + 3.14159)))
                v80 := []int{127}
                for fi81 := 0; fi81 < 17; fi81++ {
                    var v82 float32 = v2
                    _ = fi81
                    _ = v82
                }
                _ = v80
            }
        }
        v83 := GenS2{F0: 452, F1: v6}
        _ = v83
    } else {
        v2 -= -2.25
        _ = len(v1)
    }
    v6 = v6
    if false {
        v84 := GenS4{F0: 1, F1: -2147483648}
        v85 := []int{-185, int(v3), -533}
        if ((2147483647 != 2147483648) || false) {
            switch (v6 + v6) {
            case "\n\t":
                for fi86 := 0; fi86 < 8; fi86++ {
                    _ = len(v85)
                    v3 = uint(v7)
                    v6 = "hello"
                    _ = fi86
                }
            case "foo bar":
                for fi87 := 0; fi87 < 12; fi87++ {
                    fmt.Println(v7, v7, v2)
                    if false {
                        v88 := GenS5{F0: -128, F1: 1, F2: false}
                        _ = v88
                    }
                    v7 = 517
                    _ = fi87
                }
            }
            switch -1 {
            case 127:
                var v89 int32 = -2147483648
                _ = v89
            case -936:
                for fi90 := 0; fi90 < 8; fi90++ {
                    var v91 int8 = 127
                    _ = fi90
                    _ = v91
                }
            }
        } else {
            if (!(!false)) {
                switch 1 {
                case 128:
                    var v92 uint64 = (^(v7))
                    _ = v92
                default:
                    v93 := GenS2{F0: 2, F1: v6}
                    _ = v93
                }
                switch v6 {
                case "a":
                    for fi94 := 0; fi94 < 4; fi94++ {
                        var v95 uint8 = 0
                        var v96 float32 = float32((-2.25 * (float64((0.0 / v2)) * 1.5)))
                        if true {
                            if false {
                                fmt.Println(v5)
                            } else {
                                fmt.Println(v6, v6, v2)
                            }
                            for fi97 := 0; fi97 < 18; fi97++ {
                                var v98 float32 = (v2 + v2)
                                fmt.Println(fi97)
                                if (!(1000 >= 317)) {
                                    if false {
                                        switch (fi94 % ((2147483647) | 1)) {
                                        case 127:
                                            if (9223372036854775807 > (v5 / (((v5 / (((v5 - (v5 ^ (v5 & 9223372036854775807)))) | 1))) | 1))) {
                                                switch (fi97 ^ (fi94 - (fi94 ^ fi97))) {
                                                case 127:
                                                    v6 = v6
                                                default:
                                                    fmt.Println(v7, fi94)
                                                }
                                                var v99 uint16 = 169
                                                _ = v99
                                            }
                                        default:
                                            _ = len(v4)
                                        }
                                        var v100 uint8 = (v95 / (((v95 ^ v95)) | 1))
                                        v101 := GenS7{F0: (v7 >> ((v3) & 63)), F1: 172}
                                        _ = v100
                                        _ = v101
                                    }
                                    if (-343 > 834) {
                                        var v102 uint8 = v95
                                        _ = v102
                                    }
                                    v103 := v1[(fi94 - (^(fi94)))]
                                    _ = v103
                                }
                                _ = fi97
                                _ = v98
                            }
                            v104 := v4[1]
                            _ = v104
                        } else {
                            if false {
                                var v105 float32 = float32(1.5)
                                switch fi94 {
                                case 2147483647:
                                    var v106 int8 = 0
                                    _ = v106
                                default:
                                    v85 = append(v85, (fi94 >> ((255) & 63)))
                                }
                                _ = v105
                            } else {
                                if (-2.25 < 1.5) {
                                    if len(v85) > 0 {
                                        v85[0] = (-(fi94))
                                    }
                                    v5 = v5
                                    switch fi94 {
                                    case 2147483647:
                                        var v107 int = 4294967295
                                        _ = v107
                                    case 1:
                                        v2 += 1e10
                                    case 32768:
                                        v96 += v2
                                    }
                                }
                                if true {
                                    fmt.Println(v2)
                                    var v108 float64 = (3.14159 * (0.1 + (((1e10 - 3.14159) / 1.0) - ((1.5 / 1.0) - -2.25))))
                                    v2 = v96
                                    _ = v108
                                } else {
                                    fmt.Println(fi94, v5)
                                    v7 = v7
                                }
                            }
                        }
                        _ = fi94
                        _ = v95
                        _ = v96
                    }
                case "foo bar":
                    _ = len(v85)
                default:
                    switch ((("hello" + ("" + v6)) + v6) + ((((v6 + "\n\t") + (((v6 + ((v6 + "") + v6)) + ("" + v6)) + v6)) + (("\n\t" + v6) + "")) + ((("hello" + v6) + v6) + ""))) {
                    case "foo bar":
                        v5 -= (v5 ^ v5)
                    }
                }
                var v109 int32 = 42
                _ = v109
            } else {
                if true {
                    if (v5 > -9223372036854775808) {
                        if (993 > 0) {
                            for fi110 := 0; fi110 < 12; fi110++ {
                                v3 = (^(v3))
                                v7 = (v7 * (v7 | 0))
                                _ = fi110
                            }
                        } else {
                            v7 = 50
                            var v111 int8 = 100
                            for fi112 := 0; fi112 < 3; fi112++ {
                                v113 := GenS1{F0: 76}
                                switch v6 {
                                case "hello":
                                    var v114 uint32 = 272
                                    _ = v114
                                case "":
                                    v115 := GenS8{F0: 668, F1: v3}
                                    _ = v115
                                default:
                                    var v116 uint32 = 616
                                    _ = v116
                                }
                                _ = fi112
                                _ = v113
                            }
                            _ = v111
                        }
                    }
                    v117 := v1[4294967295]
                    _ = v117
                } else {
                    for fi118 := 0; fi118 < 2; fi118++ {
                        switch "\n\t" {
                        case "":
                            switch v6 {
                            case "":
                                if len(v85) >= 2 {
                                    v85 = v85[1:len(v85)]
                                }
                            case "a":
                                if (!true) {
                                    var v119 uint = v3
                                    v6 = ""
                                    _ = v119
                                }
                            default:
                                fmt.Println(fi118, v3, v2)
                            }
                        default:
                            var v120 int32 = 812
                            _ = v120
                        }
                        if false {
                            for fi121 := 0; fi121 < 1; fi121++ {
                                if (!true) {
                                    v122 := GenS1{F0: 128}
                                    fmt.Println(v7)
                                    if len(v85) >= 2 {
                                        v85 = v85[1:len(v85)]
                                    }
                                    _ = v122
                                }
                                if false {
                                    v5 = v5
                                    v5 = v5
                                    fmt.Println(v2)
                                    v6 = v6
                                } else {
                                    var v123 int16 = 42
                                    var v124 int16 = v123
                                    _ = v123
                                    _ = v124
                                }
                                var v125 string = v6
                                v125 = v125
                                _ = fi121
                                _ = v125
                            }
                            v126 := GenS9{F0: v7, F1: -2147483648}
                            v3 -= v3
                            _ = v126
                        }
                        _ = len(v85)
                        _ = fi118
                    }
                    var v127 bool = (!true)
                    _ = v127
                }
                switch (("foo bar" + v6) + "\n\t") {
                case "":
                    v85 = append(v85, -1)
                case "foo bar":
                    v6 = v6
                }
                for fi128 := 0; fi128 < 4; fi128++ {
                    delete(v4, 2)
                    _ = fi128
                }
            }
        }
        _ = v84
        _ = v85
    } else {
        var v129 float64 = (0.1 * (1.5 / 1e10))
        v130 := GenS10{F0: -2147483648, F1: -1, F2: int8(v3)}
        var v131 uint8 = 2
        _ = v129
        _ = v130
        _ = v131
    }
    if true {
        v132 := GenS5{F0: 70, F1: 32767, F2: ((v5 * (v5 - (v5 << ((255) & 63)))) > (v5 + 1))}
        v3 -= (v3 + v3)
        v133 := GenS11{F0: 127, F1: false, F2: 485}
        _ = v132
        _ = v133
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
