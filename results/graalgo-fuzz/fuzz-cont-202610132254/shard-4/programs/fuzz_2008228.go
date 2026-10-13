package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 bool; F1 uint16; F2 bool }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 uint; F1 int; F2 uint32 }
type GenS9 struct { F0 int32; F1 int32 }
type GenS10 struct { F0 int32; F1 uint32; F2 int16 }
type GenS11 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint, p1 uint) uint64 {
    v5 := []int{9223372036854775807, 903}
    var v6 int64 = 42
    var v7 uint32 = 65535
    switch 140 {
    case -319:
        for fi8 := 0; fi8 < 19; fi8++ {
            for fi9 := 0; fi9 < 6; fi9++ {
                p1 = 2
                p0 += (p1 >> ((p0) & 63))
                if (!true) {
                    if false {
                        if (0 != -128) {
                            v6 += (v6 - (v6 ^ v6))
                        }
                        fmt.Println(v6)
                        v10 := GenS1{F0: v6, F1: 2}
                        _ = v10
                    }
                    v7 -= 2
                    if (int(v6) < fi9) {
                        v6 = (v6 ^ v6)
                    }
                }
                switch "foo bar" {
                case "":
                    if true {
                        var v11 bool = (0 <= 1)
                        _ = v11
                    }
                default:
                    fmt.Println(p0, v6, p0)
                }
                v7 = (^(v7))
                _ = fi9
            }
            _ = fi8
        }
    case -1:
        if ((-1 <= 42) && (true || (!true))) {
            if true {
                var v12 int16 = -742
                switch 32768 {
                case 4294967295:
                    p0 = 537
                case -1:
                    for fi13 := 0; fi13 < 9; fi13++ {
                        v12 = (^(v12))
                        _ = fi13
                    }
                default:
                    if false {
                        var v14 uint = 803
                        v15 := GenS2{F0: -1}
                        _ = v14
                        _ = v15
                    }
                }
                _ = v12
            } else {
                for fi16 := 0; fi16 < 20; fi16++ {
                    for fi17 := 0; fi17 < 15; fi17++ {
                        _ = len(v5)
                        v18 := GenS2{F0: -128}
                        _ = fi17
                        _ = v18
                    }
                    _ = fi16
                }
                v19 := GenS2{F0: 42}
                _ = v19
            }
            var v20 string = ("\n\t" + "a")
            _ = v20
        } else {
            v6 = -1
            var v21 int16 = 1000
            switch 9223372036854775807 {
            case -2147483648:
                v22 := GenS1{F0: (v6 << (((p0 % ((p0) | 1))) & 63)), F1: 1}
                _ = v22
            }
            _ = v21
        }
    case 4294967295:
        if false {
            if (4294967295 <= 0) {
                p1 += 4294967295
                if ((3.14159 < 1e10) || ((("" + "foo bar") + ((("a" + ("" + ("" + ("" + "a")))) + "") + "")) >= "hello")) {
                    switch -716 {
                    case -9223372036854775808:
                        if len(v5) > 0 {
                            v5[0] = 128
                        }
                    default:
                        var v23 uint8 = 128
                        _ = v23
                    }
                    p0 = (p1 * (^(p0)))
                    fmt.Println(p0, v6, v6)
                }
            }
            switch 42 {
            case 236:
                if len(v5) > 0 {
                    v5[0] = 128
                }
            default:
                if len(v5) > 0 {
                    v5[0] = 128
                }
            }
        }
    default:
        fmt.Println(v6, p1, v7)
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[int]bool)
    var v3 int8 = 1
    var v4 int = 42
    var v24 string = "a"
    if (18446744073709551615 < genFn0(42, 1)) {
        var v25 int16 = 92
        v26 := GenS3{F0: 1, F1: 1}
        switch (((v24 + "") + (v24 + v24)) + (v24 + "\n\t")) {
        case "a":
            v4 -= -9223372036854775808
        case "":
            var v27 float64 = 0.0
            _ = v27
        }
        _ = v25
        _ = v26
    } else {
        if (128 != 128) {
            var v28 uint16 = (v1 & (v1 - v1))
            switch -9223372036854775808 {
            case 2147483647:
                switch -813 {
                case 127:
                    switch 42 {
                    case -616:
                        if false {
                            switch v4 {
                            case 4294967295:
                                var v29 int32 = 42
                                _ = v29
                            case 127:
                                var v30 uint16 = v28
                                _ = v30
                            case -1:
                                fmt.Println(v1, v3)
                            default:
                                fmt.Println(v28, v1, v24)
                            }
                            switch "hello" {
                            case "foo bar":
                                var v31 float32 = 3.14159
                                _ = v31
                            default:
                                var v32 int = v4
                                _ = v32
                            }
                        }
                    case 497:
                        v1 -= 32767
                    case -2147483648:
                        var v33 int = (v4 * int(v1))
                        _ = v33
                    default:
                        if true {
                            v3 += v3
                        } else {
                            v1 = (v1 % (((^(v28))) | 1))
                            var v34 int64 = -1
                            for fi35 := 0; fi35 < 7; fi35++ {
                                switch v4 {
                                case -2147483648:
                                    v24 = ((v24 + "\n\t") + (v24 + v24))
                                case 1000:
                                    var v36 uint32 = 2147483647
                                    _ = v36
                                case 128:
                                    for fi37 := 0; fi37 < 10; fi37++ {
                                        if true {
                                            var v38 uint = 42
                                            switch v4 {
                                            case 32768:
                                                switch v4 {
                                                case 127:
                                                    if false {
                                                        delete(v2, fi37)
                                                    } else {
                                                        v2[v4] = (v4 <= (v4 << ((1) & 63)))
                                                    }
                                                case 32768:
                                                    switch "" {
                                                    case "\n\t":
                                                        if (v4 < v4) {
                                                            v3 = v3
                                                            var v39 float32 = 3.14159
                                                            var v40 uint32 = 2
                                                            var v41 uint = 703
                                                            _ = len(v2)
                                                            var v42 int32 = -2147483648
                                                            _ = v39
                                                            _ = v40
                                                            _ = v41
                                                            _ = v42
                                                        } else {
                                                            fmt.Println(fi35, v3)
                                                            v43 := GenS2{F0: (v3 - (v3 ^ 0))}
                                                            _ = v43
                                                        }
                                                    case "foo bar":
                                                        delete(v2, (fi35 - v4))
                                                    }
                                                case -240:
                                                    var v44 uint32 = 822
                                                    _ = v44
                                                }
                                            case -9223372036854775808:
                                                v45 := []int{4294967295, (^(v4))}
                                                _ = v45
                                            case 0:
                                                v46 := GenS1{F0: v34, F1: 1}
                                                _ = v46
                                            default:
                                                v47 := GenS4{F0: false, F1: (v28 + v28), F2: true}
                                                _ = v47
                                            }
                                            if (42 > 1) {
                                                if (0 != 127) {
                                                    _ = len(v2)
                                                    var v48 uint32 = 970
                                                    if (75 > (v3 / (((v3 * 1)) | 1))) {
                                                        var v49 uint64 = genFn0((v38 & 4294967295), (v38 / ((v38) | 1)))
                                                        _ = v49
                                                    }
                                                    _ = v48
                                                } else {
                                                    fmt.Println(v4, v1)
                                                }
                                                var v50 bool = (!(false && true))
                                                delete(v2, fi35)
                                                _ = v50
                                            }
                                            _ = v38
                                        }
                                        _ = fi37
                                    }
                                }
                                v51 := GenS5{F0: 2, F1: 797}
                                var v52 int32 = -1
                                _ = fi35
                                _ = v51
                                _ = v52
                            }
                            _ = v34
                        }
                    }
                case 42:
                    if (42 == 398) {
                        if (!false) {
                            if (32767 != 0) {
                                fmt.Println(v4)
                                v53 := GenS6{F0: 1}
                                _ = v53
                            } else {
                                for fi54 := 0; fi54 < 17; fi54++ {
                                    var v55 int8 = v3
                                    var v56 bool = true
                                    _ = fi54
                                    _ = v55
                                    _ = v56
                                }
                                if (v1 == v28) {
                                    v57 := GenS7{F0: false, F1: 866}
                                    var v58 uint8 = uint8(v28)
                                    if false {
                                        v59 := []int{(v4 % ((-1) | 1)), 232, -345}
                                        switch v4 {
                                        case -871:
                                            v58 = 98
                                        case -151:
                                            v60 := GenS8{F0: 1, F1: (v4 * 42), F2: 708}
                                            _ = v60
                                        case -2147483648:
                                            var v61 int = v4
                                            _ = v61
                                        }
                                        v24 = (v24 + v24)
                                        _ = v59
                                    }
                                    _ = v57
                                    _ = v58
                                }
                            }
                            v62 := []int{}
                            _ = v62
                        } else {
                            if true {
                                var v63 uint8 = 234
                                _ = v63
                            }
                        }
                        for fi64 := 0; fi64 < 14; fi64++ {
                            v65 := v2[(v4 | (fi64 + (-(fi64))))]
                            if (!(!(!(uint(v1) == 42)))) {
                                v66 := GenS7{F0: false, F1: -1}
                                fmt.Println(v65, v24)
                                _ = v66
                            } else {
                                if (v28 == v1) {
                                    v28 = 0
                                    v67 := v2[v4]
                                    v65 = v65
                                    _ = v67
                                } else {
                                    switch v4 {
                                    case -665:
                                        v68 := GenS9{F0: -1, F1: -1}
                                        _ = v68
                                    case -374:
                                        switch v24 {
                                        case "a":
                                            for fi69 := 0; fi69 < 18; fi69++ {
                                                v3 += 42
                                                if v65 {
                                                    if (!(1000 == int16(fi64))) {
                                                        var v70 uint8 = 128
                                                        _ = len(v2)
                                                        switch "" {
                                                        case "\n\t":
                                                            v2[-1] = (true || false)
                                                        }
                                                        v28 += v1
                                                        _ = v70
                                                    }
                                                    v71 := []int{32767, v4, v4, (fi69 | (-(v4)))}
                                                    if v65 {
                                                        var v72 int64 = -287
                                                        v73 := GenS10{F0: -1, F1: 816, F2: 1000}
                                                        _ = v72
                                                        _ = v73
                                                    } else {
                                                        delete(v2, (^(fi69)))
                                                        var v74 uint8 = 42
                                                        if true {
                                                            var v75 int64 = 42
                                                            switch (fi64 << ((0) & 63)) {
                                                            case -1:
                                                                v76 := v2[(^(v4))]
                                                                _ = v76
                                                            case 612:
                                                                var v77 uint8 = v74
                                                                _ = v77
                                                            default:
                                                                fmt.Println(fi64, v74, fi69)
                                                            }
                                                            _ = v75
                                                        } else {
                                                            v24 = v24
                                                            var v78 int = 2147483647
                                                            _ = v78
                                                        }
                                                        _ = len(v71)
                                                        _ = v74
                                                    }
                                                    v79 := GenS2{F0: v3}
                                                    _ = v71
                                                    _ = v79
                                                }
                                                _ = fi69
                                            }
                                        case "":
                                            for fi80 := 0; fi80 < 15; fi80++ {
                                                v81 := []int{0, (v4 * fi80), (fi80 - 128), (-(fi80))}
                                                v28 -= v1
                                                fmt.Println(v1)
                                                _ = fi80
                                                _ = v81
                                            }
                                        case "\n\t":
                                            if false {
                                                switch ((v24 + v24) + v24) {
                                                case "hello":
                                                    var v82 float64 = -2.25
                                                    _ = v82
                                                case "foo bar":
                                                    v83 := []int{}
                                                    _ = v83
                                                default:
                                                    if (65535 >= 2147483648) {
                                                        for fi84 := 0; fi84 < 19; fi84++ {
                                                            var v85 bool = ((!(79 < 128)) || (350 >= v4))
                                                            switch v4 {
                                                            case -426:
                                                                v28 = v1
                                                            case -9223372036854775808:
                                                                if true {
                                                                    v65 = false
                                                                    var v86 uint = 348
                                                                    _ = v86
                                                                } else {
                                                                    var v87 uint = 0
                                                                    _ = v87
                                                                }
                                                            }
                                                            _ = fi84
                                                            _ = v85
                                                        }
                                                    }
                                                }
                                                v88 := GenS8{F0: 2, F1: fi64, F2: 503}
                                                switch (-(v4)) {
                                                case -2147483648:
                                                    v89 := []int{(v4 >> ((931) & 63)), -185, (fi64 & (v4 / ((v4) | 1))), (fi64 | (v4 / (((v4 * 9223372036854775807)) | 1)))}
                                                    _ = v89
                                                case 737:
                                                    v90 := []int{v4, (v4 - (fi64 | 128))}
                                                    _ = v90
                                                default:
                                                    v65 = ((3.14159 > 1.5) || (1.5 > (1.5 / (0.0 - (float64(-2.25) / 0.1)))))
                                                }
                                                _ = v88
                                            }
                                        default:
                                            v4 += 42
                                        }
                                    }
                                    v65 = true
                                }
                            }
                            _ = fi64
                            _ = v65
                        }
                        var v91 uint = 484
                        _ = v91
                    }
                }
            default:
                fmt.Println(v3)
            }
            var v92 int8 = (v3 | v3)
            _ = v28
            _ = v92
        }
        v4 = v4
    }
    if (v1 == (v1 - v1)) {
        delete(v2, v4)
    }
    if false {
        for fi93 := 0; fi93 < 2; fi93++ {
            var v94 int8 = v3
            _ = fi93
            _ = v94
        }
        for fi95 := 0; fi95 < 19; fi95++ {
            if (206 == -9223372036854775808) {
                v96 := []int{354, fi95, 1, -871}
                if true {
                    var v97 float64 = 0.1
                    v1 = v1
                    for fi98 := 0; fi98 < 2; fi98++ {
                        var v99 float64 = (-2.25 - (3.14159 * float64(3.14159)))
                        switch (v4 ^ (fi95 | (v4 | -845))) {
                        case 0:
                            if false {
                                v24 = (v24 + "foo bar")
                                v97 = (float64(0.0) / (-2.25 * ((1e10 * -2.25) / v97)))
                            } else {
                                var v100 uint32 = 0
                                var v101 int32 = -158
                                _ = v100
                                _ = v101
                            }
                        case 32767:
                            if ((v4 | 9223372036854775807) > fi98) {
                                fmt.Println(v4, v3, fi95)
                            } else {
                                var v102 string = v24
                                v103 := v2[fi98]
                                _ = v102
                                _ = v103
                            }
                        }
                        v3 = (v3 ^ v3)
                        var v104 int16 = -268
                        var v105 float64 = float64(0.1)
                        _ = fi98
                        _ = v104
                        _ = v105
                        _ = v99
                    }
                    _ = v97
                }
                _ = v96
            } else {
                if ((!(0.1 != (0.0 + 1e10))) || (true || true)) {
                    var v106 float64 = (((1e10 / (3.14159 * 1e10)) * ((1.5 / 1.0) / -2.25)) - (-2.25 + (0.1 + 0.1)))
                    v1 += v1
                    v107 := []int{(fi95 % ((fi95) | 1))}
                    _ = v106
                    _ = v107
                }
            }
            v108 := GenS11{F0: 127, F1: 0}
            _ = fi95
            _ = v108
        }
    }
    var v109 float32 = 0.1
    for fi110 := 0; fi110 < 18; fi110++ {
        if ((v1 * (v1 / ((v1) | 1))) == v1) {
            var v111 int16 = 32767
            v112 := GenS10{F0: 0, F1: 225, F2: (^(v111))}
            _ = v111
            _ = v112
        } else {
            fmt.Println(v4)
            if (4294967295 != 0) {
                switch v24 {
                case "\n\t":
                    var v113 string = "\n\t"
                    _ = v113
                case "hello":
                    fmt.Println(v1, v4, v3)
                default:
                    for fi114 := 0; fi114 < 9; fi114++ {
                        switch v24 {
                        case "hello":
                            v1 -= (v1 - v1)
                        case "a":
                            v1 -= (v1 * (v1 * 605))
                        case "":
                            var v115 int32 = -332
                            _ = v115
                        }
                        _ = fi114
                    }
                }
                v109 -= v109
                v3 += 127
            }
        }
        _ = fi110
    }
    for fi116 := 0; fi116 < 5; fi116++ {
        var v117 bool = true
        _ = fi116
        _ = v117
    }
    for fi118 := 0; fi118 < 19; fi118++ {
        v4 += -129
        v119 := GenS10{F0: 134, F1: 2, F2: 823}
        _ = fi118
        _ = v119
    }
    fmt.Println(v4, v3, v24)
    var v120 int64 = 1
    var v121 uint8 = 146
    v109 += v109
    var v122 uint16 = (v1 & v1)
    delete(v2, (v4 - -2147483648))
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v120:", v120)
    fmt.Println("v121:", v121)
    fmt.Println("v122:", v122)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v109
    _ = v120
    _ = v121
    _ = v122
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
