package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 bool; F1 uint; F2 string }
type GenS6 struct { F0 int; F1 int64; F2 uint16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint8; F1 uint64; F2 uint }
type GenS10 struct { F0 bool; F1 int8 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 bool; F1 string }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = -56
    v3 := make(map[int]bool)
    var v4 string = ("foo bar" + (("" + ("foo bar" + "\n\t")) + "hello"))
    if true {
        var v5 int64 = 9223372036854775807
        if true {
            for fi6 := 0; fi6 < 11; fi6++ {
                switch fi6 {
                case -497:
                    if (2147483647 < 210) {
                        v7 := GenS1{F0: 42}
                        _ = v7
                    } else {
                        for fi8 := 0; fi8 < 6; fi8++ {
                            fmt.Println(fi8, fi8, fi6)
                            _ = len(v1)
                            var v9 string = ""
                            fmt.Println(fi8, v4, v9)
                            v10 := GenS1{F0: 128}
                            v11 := GenS2{F0: 42, F1: v2}
                            switch ((v9 + v4) + ("\n\t" + "a")) {
                            case "foo bar":
                                var v12 uint16 = 65535
                                _ = v12
                            case "a":
                                fmt.Println(v2, v4, fi6)
                            }
                            v9 = v4
                            var v13 int32 = 42
                            _ = fi8
                            _ = v10
                            _ = v11
                            _ = v13
                            _ = v9
                        }
                        if (-1 != 1000) {
                            if (42 > 4294967295) {
                                switch -565 {
                                case 127:
                                    for fi14 := 0; fi14 < 10; fi14++ {
                                        if (-669 > 1) {
                                            switch (-(fi6)) {
                                            case 42:
                                                v15 := []int{fi6, (fi6 >> ((1) & 63))}
                                                _ = v15
                                            case 418:
                                                switch ((v4 + v4) + (v4 + ((v4 + "foo bar") + "hello"))) {
                                                case "foo bar":
                                                    var v16 uint16 = 2
                                                    _ = v16
                                                case "":
                                                    v2 = v2
                                                default:
                                                    switch (v4 + (v4 + "")) {
                                                    case "":
                                                        switch -9223372036854775808 {
                                                        case 49:
                                                            switch v4 {
                                                            case "hello":
                                                                var v17 int32 = 1
                                                                _ = v17
                                                            default:
                                                                switch fi14 {
                                                                case 0:
                                                                    v2 = -1
                                                                case 103:
                                                                    var v18 int32 = int32(fi14)
                                                                    _ = v18
                                                                }
                                                            }
                                                        case 107:
                                                            v5 += v5
                                                        case -9223372036854775808:
                                                            v19 := []int{(fi14 << ((73) & 63)), 128, (fi6 >> ((739) & 63))}
                                                            _ = v19
                                                        }
                                                    case "foo bar":
                                                        switch v4 {
                                                        case "":
                                                            var v20 bool = ((65535 == 65535) || ((v4 + "a") == (("\n\t" + v4) + ("a" + ((v4 + ("\n\t" + (v4 + v4))) + v4)))))
                                                            _ = v20
                                                        case "\n\t":
                                                            if (615 > -1) {
                                                                if false {
                                                                    var v21 int = fi6
                                                                    _ = v21
                                                                } else {
                                                                    v5 += (v5 >> ((761) & 63))
                                                                    switch (fi14 % ((fi14) | 1)) {
                                                                    case -5:
                                                                        var v22 int = (fi14 << ((4294967295) & 63))
                                                                        if false {
                                                                            v23 := GenS3{F0: 42, F1: 65535}
                                                                            var v24 uint64 = 440
                                                                            if (0 == 280) {
                                                                                switch (v4 + (v4 + v4)) {
                                                                                case "a":
                                                                                    if ((0.0 * (1e10 * 1.5)) < -2.25) {
                                                                                        fmt.Println(v22, fi14, fi14)
                                                                                        fmt.Println(v24, v5, fi6)
                                                                                        v25 := []int{(fi14 + (v22 * v22)), -2147483648, v22, 128}
                                                                                        if len(v25) > 0 {
                                                                                            v25[0] = fi14
                                                                                        }
                                                                                        v26 := GenS2{F0: 0, F1: (v2 % ((v2) | 1))}
                                                                                        _ = v25
                                                                                        _ = v26
                                                                                    } else {
                                                                                        fmt.Println(fi14, v5)
                                                                                        fmt.Println(v5, v24, fi6)
                                                                                        v27 := GenS1{F0: 2}
                                                                                        _ = v27
                                                                                    }
                                                                                default:
                                                                                    v28 := GenS4{F0: (v22 & fi6), F1: 0}
                                                                                    _ = v28
                                                                                }
                                                                                v29 := GenS2{F0: -1, F1: int8(fi6)}
                                                                                if true {
                                                                                    var v30 int = 24
                                                                                    var v31 int32 = 42
                                                                                    var v32 int64 = (^(v5))
                                                                                    _ = v30
                                                                                    _ = v31
                                                                                    _ = v32
                                                                                } else {
                                                                                    v33 := []int{v22, fi14, -2147483648, (fi6 - (v22 + (-(fi14))))}
                                                                                    if false {
                                                                                        if false {
                                                                                            v5 += (v5 - (v5 + (v5 + v5)))
                                                                                        }
                                                                                        v34 := GenS2{F0: 0, F1: (v2 | v2)}
                                                                                        _ = v34
                                                                                    } else {
                                                                                        var v35 float32 = (1.5 - 1.5)
                                                                                        if true {
                                                                                            v36 := GenS5{F0: false, F1: 460, F2: v4}
                                                                                            var v37 string = "a"
                                                                                            if (float64(-2.25) != -2.25) {
                                                                                                fmt.Println(v4)
                                                                                                var v38 uint8 = 1
                                                                                                _ = v38
                                                                                            }
                                                                                            _ = v36
                                                                                            _ = v37
                                                                                        } else {
                                                                                            switch 128 {
                                                                                            case 4294967295:
                                                                                                switch (^(fi6)) {
                                                                                                case 811:
                                                                                                    v39 := GenS3{F0: 2147483648, F1: 65535}
                                                                                                    _ = v39
                                                                                                }
                                                                                                var v40 int16 = -32768
                                                                                                _ = v40
                                                                                            }
                                                                                        }
                                                                                        v41 := GenS2{F0: 1000, F1: (v2 | -1)}
                                                                                        _ = v35
                                                                                        _ = v41
                                                                                    }
                                                                                    var v42 bool = false
                                                                                    v2 = (v2 * v2)
                                                                                    switch v4 {
                                                                                    case "":
                                                                                        var v43 uint8 = 255
                                                                                        _ = v43
                                                                                    }
                                                                                    v5 -= v5
                                                                                    _ = v33
                                                                                    _ = v42
                                                                                }
                                                                                _ = v29
                                                                            }
                                                                            if true {
                                                                                var v44 float32 = (((1.5 + 0.0) - float32(3.14159)) + float32((3.14159 * 1e10)))
                                                                                switch v4 {
                                                                                case "\n\t":
                                                                                    var v45 bool = false
                                                                                    var v46 uint8 = 255
                                                                                    v47 := GenS6{F0: fi6, F1: -918, F2: 255}
                                                                                    var v48 uint = 18446744073709551615
                                                                                    v2 += int8(v24)
                                                                                    _ = v45
                                                                                    _ = v46
                                                                                    _ = v47
                                                                                    _ = v48
                                                                                case "":
                                                                                    if false {
                                                                                        var v49 int64 = -911
                                                                                        var v50 int64 = v5
                                                                                        _ = v49
                                                                                        _ = v50
                                                                                    } else {
                                                                                        var v51 bool = false
                                                                                        var v52 int16 = 201
                                                                                        _ = v51
                                                                                        _ = v52
                                                                                    }
                                                                                }
                                                                                if ((!(2 > 18446744073709551615)) || (0 >= 42)) {
                                                                                    v2 -= (v2 >> ((1) & 63))
                                                                                    v22 = -877
                                                                                    fmt.Println(v44, v44, v4)
                                                                                }
                                                                                var v53 int8 = 127
                                                                                _ = v44
                                                                                _ = v53
                                                                            } else {
                                                                                v2 = (v2 * v2)
                                                                                delete(v3, fi6)
                                                                                v54 := []int{127, v22, (fi14 + (fi6 & fi6))}
                                                                                _ = v54
                                                                            }
                                                                            v4 = (((v4 + v4) + v4) + (v4 + v4))
                                                                            _ = v23
                                                                            _ = v24
                                                                        }
                                                                        _ = v22
                                                                    case 4294967295:
                                                                        var v55 int8 = -55
                                                                        _ = v55
                                                                    case -9223372036854775808:
                                                                        fmt.Println(v5)
                                                                    }
                                                                }
                                                            } else {
                                                                v56 := GenS4{F0: fi6, F1: (v2 / (((-(v2))) | 1))}
                                                                _ = v56
                                                            }
                                                        default:
                                                            v5 = (v5 & v5)
                                                        }
                                                        if true {
                                                            var v57 int64 = v5
                                                            v57 = (v5 << ((2) & 63))
                                                            var v58 uint64 = 437
                                                            _ = v57
                                                            _ = v58
                                                        }
                                                    case "\n\t":
                                                        if (false && (false && false)) {
                                                            v4 = ""
                                                        }
                                                    default:
                                                        v4 = (v4 + v4)
                                                    }
                                                }
                                                switch ("a" + (v4 + (v4 + ((v4 + (((v4 + v4) + ("\n\t" + v4)) + (v4 + v4))) + (v4 + ((("a" + v4) + (v4 + ((v4 + ((v4 + "") + (("a" + ("foo bar" + "a")) + "\n\t"))) + v4))) + "")))))) {
                                                case "a":
                                                    var v59 bool = true
                                                    _ = v59
                                                }
                                            }
                                            v60 := GenS7{F0: -9223372036854775808}
                                            var v61 int8 = 127
                                            _ = v60
                                            _ = v61
                                        }
                                        v62 := GenS8{F0: 1}
                                        _ = fi14
                                        _ = v62
                                    }
                                case -9223372036854775808:
                                    var v63 uint16 = 32767
                                    _ = v63
                                default:
                                    switch (v4 + v4) {
                                    case "\n\t":
                                        v2 = (v2 - 100)
                                    case "foo bar":
                                        v3[(fi6 >> ((105) & 63))] = (1000 == -593)
                                    default:
                                        v4 = v4
                                    }
                                }
                                switch (fi6 / (((fi6 >> ((370) & 63))) | 1)) {
                                case -2147483648:
                                    delete(v3, (fi6 | fi6))
                                case 396:
                                    var v64 uint64 = 556
                                    _ = v64
                                case 4294967295:
                                    if ((0 > v2) || ((v2 - (v2 / ((v2) | 1))) < (v2 ^ v2))) {
                                        if true {
                                            if (!false) {
                                                delete(v3, fi6)
                                            } else {
                                                var v65 int64 = v5
                                                if true {
                                                    v66 := GenS9{F0: 49, F1: 581, F2: 388}
                                                    _ = v66
                                                } else {
                                                    v2 -= 100
                                                    var v67 string = v4
                                                    if true {
                                                        v4 = v4
                                                    } else {
                                                        switch (-(fi6)) {
                                                        case 9223372036854775807:
                                                            v68 := v1[65535]
                                                            _ = v68
                                                        case 682:
                                                            for fi69 := 0; fi69 < 17; fi69++ {
                                                                switch ((v67 + v4) + (v4 + v67)) {
                                                                case "foo bar":
                                                                    switch "" {
                                                                    case "":
                                                                        v70 := []int{fi6, fi6}
                                                                        _ = v70
                                                                    case "\n\t":
                                                                        if (-2147483648 == 1) {
                                                                            v71 := GenS9{F0: 191, F1: 0, F2: 353}
                                                                            v3[(fi69 / (((fi69 ^ int(v65))) | 1))] = false
                                                                            if true {
                                                                                delete(v1, 4294967295)
                                                                            }
                                                                            _ = v71
                                                                        }
                                                                        var v72 uint32 = 4294967295
                                                                        _ = v72
                                                                    }
                                                                default:
                                                                    v73 := v3[fi69]
                                                                    _ = v73
                                                                }
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    if false {
                                                                        v74 := GenS10{F0: (-1 != 2147483647), F1: (v2 % (((v2 | (v2 / ((42) | 1)))) | 1))}
                                                                        _ = v74
                                                                    }
                                                                default:
                                                                    switch 32768 {
                                                                    case 127:
                                                                        v75 := GenS10{F0: false, F1: -1}
                                                                        _ = v75
                                                                    case -123:
                                                                        switch ((v4 + (v4 + v67)) + ("hello" + v4)) {
                                                                        case "\n\t":
                                                                            v76 := GenS4{F0: fi6, F1: v2}
                                                                            _ = v76
                                                                        case "foo bar":
                                                                            v77 := GenS11{F0: 2}
                                                                            _ = v77
                                                                        }
                                                                    case 268:
                                                                        if (!(!(1.5 <= 3.14159))) {
                                                                            var v78 bool = (0 < 32767)
                                                                            var v79 int8 = v2
                                                                            _ = v78
                                                                            _ = v79
                                                                        }
                                                                    default:
                                                                        switch "foo bar" {
                                                                        case "a":
                                                                            switch (fi69 & fi69) {
                                                                            case -1:
                                                                                var v80 int32 = -2147483648
                                                                                _ = v80
                                                                            default:
                                                                                _ = len(v3)
                                                                                switch (((((v67 + (v4 + "a")) + v67) + (v4 + "hello")) + v67) + v4) {
                                                                                case "":
                                                                                    v81 := []int{(fi6 ^ fi6), -9223372036854775808}
                                                                                    _ = v81
                                                                                case "a":
                                                                                    switch 42 {
                                                                                    case 32767:
                                                                                        v82 := GenS10{F0: false, F1: 42}
                                                                                        _ = v82
                                                                                    case 234:
                                                                                        v67 = "\n\t"
                                                                                    case 337:
                                                                                        v4 = "a"
                                                                                    }
                                                                                case "hello":
                                                                                    delete(v3, (-(fi6)))
                                                                                default:
                                                                                    switch v67 {
                                                                                    case "":
                                                                                        fmt.Println(v67)
                                                                                    case "foo bar":
                                                                                        v67 = ""
                                                                                    case "a":
                                                                                        if (0 >= (v2 % (((v2 | (^(v2)))) | 1))) {
                                                                                            var v83 uint16 = 342
                                                                                            v83 += (v83 - 65535)
                                                                                            switch -1 {
                                                                                            case 2147483647:
                                                                                                v1[65535] = (v67 + "\n\t")
                                                                                                delete(v1, 372)
                                                                                                v65 += 767
                                                                                            }
                                                                                            v84 := GenS9{F0: 42, F1: 1, F2: 1}
                                                                                            _ = v83
                                                                                            _ = v84
                                                                                        }
                                                                                    }
                                                                                }
                                                                                v65 -= (-(v5))
                                                                                var v85 uint = 65535
                                                                                _ = v85
                                                                            }
                                                                        case "\n\t":
                                                                            v2 = v2
                                                                        case "hello":
                                                                            v86 := GenS9{F0: 128, F1: 2, F2: 632}
                                                                            _ = v86
                                                                        }
                                                                    }
                                                                }
                                                                if false {
                                                                    var v87 float32 = -2.25
                                                                    v88 := []int{fi69}
                                                                    _ = v87
                                                                    _ = v88
                                                                }
                                                                _ = fi69
                                                            }
                                                        default:
                                                            if false {
                                                                fmt.Println(v67, v2, v2)
                                                                if true {
                                                                    v4 = v67
                                                                    var v89 bool = ((!false) && true)
                                                                    var v90 int64 = (v65 / ((-242) | 1))
                                                                    _ = v89
                                                                    _ = v90
                                                                }
                                                            } else {
                                                                _ = len(v3)
                                                            }
                                                        }
                                                        v91 := GenS7{F0: (v65 * 1)}
                                                        _ = v91
                                                    }
                                                    _ = v67
                                                }
                                                var v92 int16 = -1
                                                _ = v65
                                                _ = v92
                                            }
                                        }
                                        for fi93 := 0; fi93 < 16; fi93++ {
                                            _ = len(v3)
                                            _ = fi93
                                        }
                                        var v94 int64 = 0
                                        _ = v94
                                    } else {
                                        v95 := v3[fi6]
                                        _ = v95
                                    }
                                default:
                                    var v96 uint16 = 0
                                    _ = v96
                                }
                            } else {
                                for fi97 := 0; fi97 < 6; fi97++ {
                                    v4 = v4
                                    _ = fi97
                                }
                                var v98 float32 = -2.25
                                _ = v98
                            }
                            if (true || (!false)) {
                                switch "a" {
                                case "":
                                    var v99 uint16 = 255
                                    _ = v99
                                case "foo bar":
                                    v100 := GenS12{F0: (!false), F1: "foo bar"}
                                    _ = v100
                                case "\n\t":
                                    v2 -= v2
                                }
                                fmt.Println(fi6, fi6)
                            }
                            v101 := []int{fi6, fi6, (-(fi6)), fi6}
                            _ = v101
                        }
                    }
                }
                var v102 uint64 = 42
                _ = fi6
                _ = v102
            }
        } else {
            for fi103 := 0; fi103 < 13; fi103++ {
                if true {
                    for fi104 := 0; fi104 < 6; fi104++ {
                        var v105 bool = true
                        _ = fi104
                        _ = v105
                    }
                } else {
                    var v106 uint = 0
                    _ = v106
                }
                v107 := GenS4{F0: fi103, F1: 127}
                _ = fi103
                _ = v107
            }
        }
        _ = v5
    }
    for fi108 := 0; fi108 < 16; fi108++ {
        if true {
            if false {
                fmt.Println(v2)
                v109 := GenS11{F0: 18446744073709551615}
                fmt.Println(v4, v2)
                _ = v109
            }
        }
        fmt.Println(v2, fi108, v2)
        v3[437] = false
        _ = fi108
    }
    v110 := []int{}
    fmt.Println(v2, v2)
    if ((42 >= 71) && true) {
        var v111 float32 = 3.14159
        var v112 uint32 = 0
        for fi113 := 0; fi113 < 11; fi113++ {
            if true {
                fmt.Println(v112, fi113)
                var v114 uint8 = 127
                _ = v114
            } else {
                v111 += v111
                var v115 uint64 = 18446744073709551615
                _ = v115
            }
            delete(v3, 42)
            _ = fi113
        }
        _ = v111
        _ = v112
    }
    var v116 int32 = 373
    v2 = (v2 - v2)
    v117 := GenS4{F0: 964, F1: v2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v110:", len(v110), v110)
    fmt.Println("v116:", v116)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v110
    _ = v116
    _ = v117
    _ = v2
    _ = v3
    _ = v4
}
