package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int64 }
type GenS2 struct { F0 uint64; F1 int64; F2 int }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 int16; F1 uint; F2 int64 }
type GenS7 struct { F0 uint8; F1 string }
type GenS8 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 uint8 = 127
    var v2 float32 = 1e10
    v3 := make(map[uint32]string)
    v4 := make(map[string]string)
    v5 := []int{}
    v4["foo bar"] = "hello"
    var v6 int64 = 1
    if len(v5) > 0 {
        v5[0] = 128
    }
    if (true || true) {
        v7 := GenS1{F0: (v6 + (v6 % ((v6) | 1))), F1: 134, F2: (v6 - 920)}
        var v8 bool = true
        _ = v7
        _ = v8
    }
    _ = len(v3)
    v9 := v4["hello"]
    v10 := GenS1{F0: 1, F1: (v6 ^ (v6 >> ((2) & 63))), F2: (v6 / ((-9223372036854775808) | 1))}
    v11 := GenS1{F0: -9223372036854775808, F1: v6, F2: (^(v6))}
    v1 -= v1
    if (!(42 < 65535)) {
        if ((v6 / ((9223372036854775807) | 1)) <= (v6 * 42)) {
            if (4 < 127) {
                fmt.Println(v1, v1, v6)
            } else {
                switch 32768 {
                case 2147483647:
                    if (true || false) {
                        _ = len(v5)
                        var v12 int32 = -1
                        _ = v12
                    } else {
                        v13 := GenS1{F0: -1, F1: 72, F2: v6}
                        v2 -= float32(3.14159)
                        _ = v13
                    }
                default:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                switch 32767 {
                case 702:
                    switch 32768 {
                    case 703:
                        v1 = v1
                    case -1:
                        var v14 int64 = (v6 / ((-1) | 1))
                        _ = v14
                    case -9223372036854775808:
                        v15 := v3[1]
                        _ = v15
                    default:
                        v16 := GenS1{F0: 42, F1: (v6 << ((26) & 63)), F2: 42}
                        _ = v16
                    }
                }
            }
            _ = len(v5)
            v9 = ""
        } else {
            v1 = v1
            v5 = append(v5, -2147483648)
            if true {
                v2 = (v2 / v2)
            } else {
                if false {
                    var v17 bool = (false || (v9 != v9))
                    _ = v17
                }
                if (!(true || (0 >= 18446744073709551615))) {
                    v5 = append(v5, -2147483648)
                }
            }
        }
        var v18 uint32 = 0
        _ = v18
    } else {
        v6 = (v6 * (v6 * 0))
        var v19 float64 = (0.0 * -2.25)
        _ = v19
    }
    switch -9223372036854775808 {
    case -1:
        v1 += (^(v1))
    case -9223372036854775808:
        switch (v9 + "") {
        case "\n\t":
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 int32 = -2147483648
                _ = fi20
                _ = v21
            }
        case "foo bar":
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 int = fi22
                v9 = "hello"
                _ = fi22
                _ = v23
            }
        }
    case -997:
        v24 := GenS1{F0: (-(v6)), F1: (^(v6)), F2: v6}
        _ = v24
    }
    if (65535 < 936) {
        v2 += ((v2 - ((float32((1e10 + 0.1)) * v2) * v2)) * v2)
        v5 = append(v5, -561)
        if false {
            fmt.Println(v1, v2)
            v25 := GenS2{F0: 4294967295, F1: v6, F2: -578}
            _ = v25
        }
    }
    fmt.Println(v6, v9, v9)
    switch (("a" + "foo bar") + v9) {
    case "a":
        _ = len(v5)
    case "hello":
        if true {
            for fi26 := 0; fi26 < 5; fi26++ {
                var v27 uint64 = 0
                var v28 float64 = 0.0
                _ = fi26
                _ = v27
                _ = v28
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi29 := 0; fi29 < 3; fi29++ {
                if true {
                    var v30 float64 = ((-2.25 - 1.5) - (1.5 / ((0.0 * float64(v2)) - ((-2.25 - (0.0 - 0.0)) + 1e10))))
                    v31 := GenS2{F0: 277, F1: (v6 | (-(v6))), F2: fi29}
                    v9 = v9
                    _ = v30
                    _ = v31
                } else {
                    var v32 uint32 = 2
                    var v33 int = int(v1)
                    _ = v32
                    _ = v33
                }
                var v34 int8 = 1
                var v35 uint64 = 4294967295
                _ = fi29
                _ = v34
                _ = v35
            }
        } else {
            var v36 int64 = v6
            if (390 != 2) {
                v6 -= -811
            } else {
                var v37 int = 719
                var v38 uint32 = 65535
                _ = v37
                _ = v38
            }
            var v39 uint = 491
            _ = v36
            _ = v39
        }
    default:
        if ((0.1 - 3.14159) == 0.1) {
            var v40 uint64 = 154
            _ = v40
        } else {
            switch ((v9 + "foo bar") + (v9 + ("hello" + v9))) {
            case "":
                if false {
                    var v41 uint64 = uint64(v1)
                    _ = v41
                } else {
                    for fi42 := 0; fi42 < 15; fi42++ {
                        if (!(v9 > ("foo bar" + (("foo bar" + v9) + (v9 + ((("a" + ((v9 + (v9 + "")) + v9)) + "\n\t") + v9)))))) {
                            if ((false && (true || (1e10 < ((float64(((v2 - (1e10 - 0.1)) + 0.0)) * 0.1) * (float64(0.1) + -2.25))))) || false) {
                                if (false || false) {
                                    switch v9 {
                                    case "hello":
                                        var v43 float64 = 1e10
                                        _ = v43
                                    case "\n\t":
                                        var v44 bool = ((true || false) || true)
                                        _ = v44
                                    case "a":
                                        for fi45 := 0; fi45 < 19; fi45++ {
                                            var v46 uint64 = 0
                                            v47 := v4[v9]
                                            v47 = v9
                                            if true {
                                                var v48 float64 = 3.14159
                                                var v49 uint64 = (v46 << ((939) & 63))
                                                if (12 <= 0) {
                                                    v6 += (v6 * -9223372036854775808)
                                                    v5 = append(v5, fi42)
                                                    v50 := GenS2{F0: (v46 | v46), F1: 42, F2: (fi42 >> ((18446744073709551615) & 63))}
                                                    _ = v50
                                                } else {
                                                    v1 += v1
                                                    var v51 uint8 = v1
                                                    _ = v51
                                                }
                                                _ = v48
                                                _ = v49
                                            }
                                            if (false || (-2.25 < 1e10)) {
                                                if len(v5) > 0 {
                                                    v5[0] = (fi45 & fi42)
                                                }
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                switch "a" {
                                                case "hello":
                                                    v6 -= (v6 << ((1) & 63))
                                                default:
                                                    v52 := GenS3{F0: 2147483647, F1: (-(fi42))}
                                                    _ = v52
                                                }
                                            } else {
                                                v9 = (v47 + v47)
                                            }
                                            _ = fi45
                                            _ = v46
                                            _ = v47
                                        }
                                    default:
                                        switch -2147483648 {
                                        case 42:
                                            v2 -= (v2 / 1.5)
                                        case -9223372036854775808:
                                            if (false || false) {
                                                var v53 int8 = 127
                                                if true {
                                                    v5 = append(v5, fi42)
                                                    fmt.Println(v2, v6)
                                                } else {
                                                    if (-32768 <= -823) {
                                                        v1 = v1
                                                        switch (v9 + "\n\t") {
                                                        case "\n\t":
                                                            for fi54 := 0; fi54 < 7; fi54++ {
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                _ = len(v5)
                                                                _ = fi54
                                                            }
                                                        }
                                                        v53 += 1
                                                    }
                                                    v55 := GenS1{F0: 1, F1: v6, F2: -665}
                                                    v56 := GenS3{F0: (fi42 * fi42), F1: 1}
                                                    _ = v55
                                                    _ = v56
                                                }
                                                _ = v53
                                            } else {
                                                switch (fi42 & (fi42 * (-(fi42)))) {
                                                case 4294967295:
                                                    v6 = (v6 >> ((42) & 63))
                                                case 407:
                                                    v1 = 77
                                                }
                                                for fi57 := 0; fi57 < 5; fi57++ {
                                                    v9 = ("a" + "hello")
                                                    _ = fi57
                                                }
                                                delete(v4, ("a" + v9))
                                            }
                                        case 0:
                                            fmt.Println(v1, v6)
                                        default:
                                            var v58 bool = (126 == 127)
                                            _ = v58
                                        }
                                    }
                                    var v59 float32 = (v2 + (v2 * float32(-2.25)))
                                    _ = v59
                                } else {
                                    fmt.Println(v2, v9)
                                    var v60 uint = 2
                                    _ = v60
                                }
                                var v61 uint32 = 933
                                _ = v61
                            } else {
                                var v62 string = v9
                                _ = v62
                            }
                        } else {
                            var v63 bool = true
                            _ = v63
                        }
                        switch (fi42 << ((255) & 63)) {
                        case 32767:
                            fmt.Println(v1, fi42, fi42)
                        case -1:
                            var v64 int64 = (v6 ^ -9223372036854775808)
                            _ = v64
                        case 1:
                            for fi65 := 0; fi65 < 13; fi65++ {
                                v66 := v3[2]
                                _ = fi65
                                _ = v66
                            }
                        default:
                            v67 := GenS4{F0: 427}
                            _ = v67
                        }
                        _ = fi42
                    }
                    v9 = v9
                    fmt.Println(v1)
                }
            default:
                if (!true) {
                    var v68 bool = (7 <= 42)
                    delete(v3, 2)
                    v1 = (^(v1))
                    _ = v68
                } else {
                    switch (("a" + (v9 + (("a" + v9) + v9))) + v9) {
                    case "a":
                        if false {
                            for fi69 := 0; fi69 < 14; fi69++ {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v1 = v1
                                _ = fi69
                            }
                            v1 -= 249
                            for fi70 := 0; fi70 < 9; fi70++ {
                                v71 := GenS3{F0: (fi70 - (fi70 << ((4294967295) & 63))), F1: (-(fi70))}
                                fmt.Println(v9, v2, v1)
                                switch v9 {
                                case "a":
                                    if (!((fi70 + -1) < (-(fi70)))) {
                                        switch 105 {
                                        case -305:
                                            var v72 bool = (!(false && (false || true)))
                                            _ = v72
                                        case 0:
                                            var v73 string = (v9 + "a")
                                            _ = v73
                                        default:
                                            var v74 uint8 = 148
                                            _ = v74
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v75 bool = true
                                            _ = v75
                                        case "\n\t":
                                            switch v9 {
                                            case "":
                                                var v76 bool = ((714 != 420) && true)
                                                _ = v76
                                            case "\n\t":
                                                delete(v4, v9)
                                            default:
                                                v77 := GenS4{F0: 255}
                                                _ = v77
                                            }
                                        }
                                    } else {
                                        fmt.Println(v6, v9)
                                        if ((true && true) && false) {
                                            switch fi70 {
                                            case -9223372036854775808:
                                                var v78 string = (("foo bar" + v9) + ((v9 + ((v9 + v9) + v9)) + v9))
                                                _ = v78
                                            case -2147483648:
                                                switch (fi70 * 4294967295) {
                                                case 42:
                                                    for fi79 := 0; fi79 < 9; fi79++ {
                                                        switch (fi70 + fi70) {
                                                        case 9223372036854775807:
                                                            if false {
                                                                var v80 bool = ((-32768 < 32767) || (!true))
                                                                if v80 {
                                                                    if v80 {
                                                                        var v81 int8 = -114
                                                                        v80 = v80
                                                                        _ = v81
                                                                    }
                                                                    v6 = (v6 + 42)
                                                                    if (!v80) {
                                                                        v9 = v9
                                                                        var v82 uint = 323
                                                                        if v80 {
                                                                            fmt.Println(v9)
                                                                        }
                                                                        _ = v82
                                                                    } else {
                                                                        switch v9 {
                                                                        case "\n\t":
                                                                            delete(v4, "a")
                                                                        default:
                                                                            v83 := GenS1{F0: (-(v6)), F1: v6, F2: v6}
                                                                            _ = v83
                                                                        }
                                                                    }
                                                                    var v84 int16 = 632
                                                                    v5 = append(v5, 127)
                                                                    v85 := GenS5{F0: 1, F1: v84}
                                                                    _ = v84
                                                                    _ = v85
                                                                }
                                                                _ = v80
                                                            }
                                                        }
                                                        var v86 uint32 = 4294967295
                                                        _ = fi79
                                                        _ = v86
                                                    }
                                                case 2147483647:
                                                    var v87 string = v9
                                                    _ = v87
                                                }
                                            default:
                                                var v88 float64 = 0.1
                                                _ = v88
                                            }
                                            if (!true) {
                                                var v89 uint16 = 199
                                                if len(v5) > 0 {
                                                    v5[0] = int(v89)
                                                }
                                                switch "hello" {
                                                case "\n\t":
                                                    switch fi70 {
                                                    case -2147483648:
                                                        v89 += (v89 % ((v89) | 1))
                                                    case 32767:
                                                        for fi90 := 0; fi90 < 1; fi90++ {
                                                            var v91 uint32 = 2
                                                            _ = fi90
                                                            _ = v91
                                                        }
                                                    }
                                                case "a":
                                                    v92 := GenS6{F0: int16(fi70), F1: 2, F2: v6}
                                                    _ = v92
                                                default:
                                                    var v93 string = "foo bar"
                                                    _ = v93
                                                }
                                                _ = v89
                                            } else {
                                                v1 -= v1
                                                v1 += v1
                                            }
                                            switch (^(fi70)) {
                                            case 128:
                                                if (!true) {
                                                    for fi94 := 0; fi94 < 11; fi94++ {
                                                        v95 := GenS7{F0: (^(v1)), F1: ("hello" + v9)}
                                                        v4[v9] = "\n\t"
                                                        _ = fi94
                                                        _ = v95
                                                    }
                                                }
                                            case 32767:
                                                v6 -= 358
                                            case -666:
                                                var v96 uint16 = 0
                                                _ = v96
                                            default:
                                                _ = len(v4)
                                            }
                                        }
                                    }
                                case "hello":
                                    var v97 int32 = -2147483648
                                    _ = v97
                                default:
                                    switch ((v9 + (v9 + v9)) + v9) {
                                    case "a":
                                        v1 -= v1
                                    }
                                }
                                _ = fi70
                                _ = v71
                            }
                        } else {
                            switch v9 {
                            case "foo bar":
                                var v98 uint = 1
                                _ = v98
                            case "hello":
                                var v99 uint = 65535
                                _ = v99
                            case "a":
                                v100 := GenS4{F0: 65535}
                                _ = v100
                            default:
                                var v101 float32 = float32(-2.25)
                                _ = v101
                            }
                            v1 -= v1
                        }
                    }
                    for fi102 := 0; fi102 < 7; fi102++ {
                        switch -1 {
                        case 2147483647:
                            var v103 uint32 = 2147483648
                            _ = v103
                        case 1:
                            switch 0 {
                            case -151:
                                var v104 float32 = float32((0.1 / 1e10))
                                _ = v104
                            case 4294967295:
                                var v105 float32 = -2.25
                                _ = v105
                            case -371:
                                for fi106 := 0; fi106 < 7; fi106++ {
                                    switch "hello" {
                                    case "\n\t":
                                        var v107 int = (fi102 >> ((873) & 63))
                                        _ = v107
                                    case "":
                                        var v108 uint = 4294967295
                                        _ = v108
                                    default:
                                        v109 := GenS2{F0: uint64(v1), F1: v6, F2: (fi102 / ((9223372036854775807) | 1))}
                                        _ = v109
                                    }
                                    var v110 float32 = (v2 - v2)
                                    _ = fi106
                                    _ = v110
                                }
                            default:
                                for fi111 := 0; fi111 < 13; fi111++ {
                                    fmt.Println(v9, fi102)
                                    _ = fi111
                                }
                            }
                        }
                        v112 := GenS8{F0: 0, F1: 226}
                        _ = fi102
                        _ = v112
                    }
                }
            }
            var v113 string = v9
            _ = v113
        }
    }
    for fi114 := 0; fi114 < 1; fi114++ {
        if ((1.5 * 3.14159) >= float64((v2 - 0.1))) {
            var v115 float32 = (v2 + float32((1e10 * (0.1 / 1.0))))
            if (32767 > 1) {
                v115 = v115
            }
            _ = v115
        }
        for fi116 := 0; fi116 < 7; fi116++ {
            fmt.Println(v1, v9, v9)
            fmt.Println(v1)
            _ = fi116
        }
        for fi117 := 0; fi117 < 9; fi117++ {
            var v118 uint = 0
            _ = fi117
            _ = v118
        }
        _ = fi114
    }
    var v119 int = 42
    fmt.Println("v1:", v1)
    fmt.Println("v119:", v119)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v119
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
