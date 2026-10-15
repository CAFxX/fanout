package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint; F1 uint8; F2 string }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64; F1 string; F2 uint }
type GenS7 struct { F0 int64; F1 int64; F2 uint32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint64; F1 uint }
func genFn0() uint64 {
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 uint8 = 142
        _ = fi6
        _ = v7
    }
    var v8 int32 = 1
    if false {
        switch 128 {
        case 1:
            v8 = v8
        }
        v8 -= (v8 << ((2) & 63))
        v9 := []int{2147483647, 42, -9223372036854775808, 0}
        _ = v9
    }
    v10 := GenS1{F0: true, F1: 141}
    v11 := []int{128, -1}
    _ = v10
    _ = v11
    _ = v8
    return 1
}
func genFn1(p0 uint32) uint64 {
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 uint = 42
        for fi14 := 0; fi14 < 6; fi14++ {
            v15 := []int{(fi12 * fi14), 1, -9223372036854775808, fi14}
            var v16 int8 = -128
            v13 -= v13
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    var v17 uint8 = 255
    switch "foo bar" {
    case "\n\t":
        var v18 uint = 686
        _ = v18
    case "":
        for fi19 := 0; fi19 < 11; fi19++ {
            p0 = (p0 - 856)
            _ = fi19
        }
    }
    v20 := []int{-922, 2147483647, -2147483648, 0}
    _ = p0
    _ = v17
    _ = v20
    return uint64(v17)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 42
    var v3 int32 = 2147483647
    var v4 uint64 = 42
    v5 := []string{("\n\t" + ("foo bar" + "foo bar")), "a", "hello"}
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := GenS1{F0: ((false || false) && false), F1: 255}
        v5 = append(v5, "a")
        _ = fi21
        _ = v22
    }
    fmt.Println(v2)
    for fi23 := 0; fi23 < 7; fi23++ {
        if true {
            fmt.Println(v4)
            fmt.Println(v3, v4, v3)
            var v24 uint = 255
            _ = v24
        }
        v25 := GenS1{F0: (!true), F1: 128}
        if true {
            if (408 != 255) {
                for fi26 := 0; fi26 < 10; fi26++ {
                    switch ("" + "foo bar") {
                    case "\n\t":
                        v1[fi23] = uint32(v4)
                    case "":
                        v3 = -2147483648
                    case "hello":
                        var v27 uint = 292
                        _ = v27
                    default:
                        fmt.Println(v3, v3, v2)
                        var v28 int16 = 110
                        _ = v28
                    }
                    _ = fi26
                }
                for fi29 := 0; fi29 < 12; fi29++ {
                    v30 := v1[(-(v2))]
                    var v31 uint64 = (v4 + v4)
                    v32 := GenS1{F0: true, F1: 76}
                    _ = fi29
                    _ = v30
                    _ = v31
                    _ = v32
                }
            }
        }
        _ = fi23
        _ = v25
    }
    if true {
        if true {
            _ = len(v1)
        }
    }
    switch (v2 - v2) {
    case -9223372036854775808:
        var v33 int16 = 42
        _ = v33
    case -132:
        v34 := GenS2{F0: 42, F1: 255, F2: "foo bar"}
        _ = v34
    case 1:
        fmt.Println(v3, v3, v3)
    default:
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
    }
    for fi35 := 0; fi35 < 9; fi35++ {
        v3 = (v3 / ((-2147483648) | 1))
        v36 := GenS2{F0: 0, F1: 128, F2: ""}
        for fi37 := 0; fi37 < 17; fi37++ {
            var v38 int = (fi35 * fi35)
            v38 = fi35
            _ = fi37
            _ = v38
        }
        _ = fi35
        _ = v36
    }
    var v39 string = ("" + "foo bar")
    var v40 uint64 = (v4 ^ v4)
    for fi41 := 0; fi41 < 20; fi41++ {
        v42 := GenS2{F0: 57, F1: 0, F2: v39}
        if false {
            var v43 uint32 = 42
            _ = v43
        } else {
            v40 = (^(v40))
        }
        _ = fi41
        _ = v42
    }
    switch v39 {
    case "\n\t":
        if false {
            if (2 <= 2) {
                for fi44 := 0; fi44 < 17; fi44++ {
                    for fi45 := 0; fi45 < 17; fi45++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        if false {
                            if true {
                                switch (fi44 * -9223372036854775808) {
                                case -823:
                                    var v46 string = ((v39 + (("foo bar" + v39) + "")) + v39)
                                    _ = v46
                                case 267:
                                    v39 = v39
                                    if (!false) {
                                        if true {
                                            v47 := GenS2{F0: 2, F1: 128, F2: "hello"}
                                            if (((922 < 1000) && true) || false) {
                                                switch (fi45 - (fi45 ^ v2)) {
                                                case -9223372036854775808:
                                                    v4 = v40
                                                case 2147483647:
                                                    v48 := GenS1{F0: (391 >= 0), F1: 41}
                                                    _ = v48
                                                }
                                            }
                                            v49 := GenS1{F0: true, F1: 2}
                                            _ = v47
                                            _ = v49
                                        }
                                        v3 = -778
                                        var v50 int64 = 968
                                        v51 := GenS1{F0: (368 != 649), F1: 1}
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = v50
                                        _ = v51
                                    }
                                case 378:
                                    if true {
                                        fmt.Println(v4, v4)
                                        switch ("" + v39) {
                                        case "foo bar":
                                            var v52 uint16 = 32767
                                            _ = v52
                                        case "a":
                                            var v53 bool = true
                                            _ = v53
                                        default:
                                            v54 := GenS1{F0: (true && true), F1: 42}
                                            _ = v54
                                        }
                                        v55 := GenS1{F0: (true && (uint16(v40) < 333)), F1: 42}
                                        _ = v55
                                    }
                                    v40 += (^(v4))
                                    var v56 int64 = 9223372036854775807
                                    _ = v56
                                default:
                                    switch fi45 {
                                    case 32768:
                                        v40 = (v40 * v40)
                                    case -971:
                                        fmt.Println(v2, v39, v39)
                                    }
                                }
                                v4 = (v40 ^ (^(v4)))
                            } else {
                                var v57 int8 = -128
                                v40 = v40
                                _ = v57
                            }
                        } else {
                            _ = len(v5)
                        }
                        fmt.Println(v3, v3, v4)
                        v3 = v3
                        v58 := GenS1{F0: (v39 <= (v39 + "\n\t")), F1: 42}
                        switch v39 {
                        case "a":
                            if true {
                                _ = len(v5)
                                var v59 uint32 = 319
                                switch fi45 {
                                case 32767:
                                    fmt.Println(v2, fi45)
                                }
                                _ = v59
                            } else {
                                v60 := GenS1{F0: true, F1: 96}
                                if (65535 <= 1) {
                                    v2 += (v2 | fi45)
                                    if (1 >= 255) {
                                        v2 += int(v3)
                                        if true {
                                            v61 := GenS3{F0: 255, F1: 754}
                                            if (!false) {
                                                v62 := GenS4{F0: 0, F1: -535}
                                                switch fi44 {
                                                case -9223372036854775808:
                                                    var v63 string = "\n\t"
                                                    _ = v63
                                                case -1:
                                                    if false {
                                                        switch v39 {
                                                        case "":
                                                            v64 := GenS5{F0: true}
                                                            _ = v64
                                                        }
                                                        switch v39 {
                                                        case "":
                                                            v4 = (v4 / (((^(v40))) | 1))
                                                        default:
                                                            var v65 uint32 = 65535
                                                            _ = v65
                                                        }
                                                        delete(v1, v2)
                                                    }
                                                default:
                                                    if (-1 != 0) {
                                                        _ = len(v5)
                                                        v4 += v40
                                                    }
                                                }
                                                _ = v62
                                            }
                                            v66 := GenS5{F0: false}
                                            _ = v61
                                            _ = v66
                                        }
                                    } else {
                                        switch (-(v2)) {
                                        case 32767:
                                            switch -906 {
                                            case -2147483648:
                                                switch (v39 + ("a" + v39)) {
                                                case "foo bar":
                                                    _ = len(v1)
                                                default:
                                                    fmt.Println(v40)
                                                }
                                            case 576:
                                                v4 += genFn1(4294967295)
                                            }
                                        case 2147483647:
                                            if (!(true || false)) {
                                                if (fi44 >= (^(v2))) {
                                                    switch v39 {
                                                    case "foo bar":
                                                        var v67 float64 = 1e10
                                                        _ = v67
                                                    default:
                                                        v4 = (v4 % ((117) | 1))
                                                    }
                                                } else {
                                                    v39 = ("foo bar" + "")
                                                    v4 = 18446744073709551615
                                                    v3 = 0
                                                }
                                                _ = len(v1)
                                                var v68 uint64 = v40
                                                _ = v68
                                            } else {
                                                v69 := GenS4{F0: v2, F1: -32768}
                                                v1[(fi44 | (fi44 | fi44))] = 0
                                                _ = v69
                                            }
                                        case 32768:
                                            if false {
                                                if (-128 < 1) {
                                                    fmt.Println(v2)
                                                    v70 := GenS6{F0: (v4 + (v4 / ((v40) | 1))), F1: v39, F2: 940}
                                                    v1[fi45] = 2
                                                    fmt.Println(v3, v2)
                                                    var v71 uint32 = 2147483647
                                                    _ = len(v1)
                                                    var v72 float64 = (0.0 / 1.5)
                                                    v4 += (^(v4))
                                                    if ("foo bar" == (v39 + "hello")) {
                                                        switch fi44 {
                                                        case 9223372036854775807:
                                                            var v73 string = v39
                                                            _ = v73
                                                        case 358:
                                                            v40 -= (v40 << ((42) & 63))
                                                        case 724:
                                                            v1[fi45] = 4294967295
                                                        }
                                                    } else {
                                                        if ((((v72 * (v72 - v72)) - v72) / 0.1) > v72) {
                                                            if len(v5) > 0 {
                                                                v5[0] = (v39 + v39)
                                                            }
                                                            v4 -= 0
                                                        } else {
                                                            if false {
                                                                var v74 uint32 = v71
                                                                _ = v74
                                                            } else {
                                                                v3 = v3
                                                                _ = len(v5)
                                                                switch v2 {
                                                                case -9223372036854775808:
                                                                    var v75 uint8 = 128
                                                                    _ = v75
                                                                case -896:
                                                                    if (409 > 86) {
                                                                        v76 := GenS7{F0: 920, F1: -9223372036854775808, F2: v71}
                                                                        _ = v76
                                                                    } else {
                                                                        var v77 float64 = (v72 * 0.1)
                                                                        fmt.Println(v71)
                                                                        _ = v77
                                                                    }
                                                                case 127:
                                                                    v40 = 18446744073709551615
                                                                }
                                                            }
                                                            v78 := GenS6{F0: v40, F1: "\n\t", F2: 255}
                                                            var v79 int32 = v3
                                                            _ = v78
                                                            _ = v79
                                                        }
                                                        var v80 int64 = -114
                                                        v81 := GenS3{F0: 42, F1: (^(v80))}
                                                        _ = v80
                                                        _ = v81
                                                    }
                                                    _ = v70
                                                    _ = v71
                                                    _ = v72
                                                }
                                                var v82 string = ("foo bar" + v39)
                                                _ = v82
                                            } else {
                                                fmt.Println(v39, v2, v39)
                                            }
                                            v4 = (v4 << ((1) & 63))
                                            var v83 uint64 = v40
                                            _ = v83
                                        default:
                                            v4 -= 403
                                        }
                                    }
                                    if (v39 <= "a") {
                                        fmt.Println(fi45, v39, fi44)
                                        var v84 bool = (v3 >= 1)
                                        _ = v84
                                    } else {
                                        if false {
                                            if false {
                                                v85 := GenS8{F0: 65535}
                                                v1[(v2 >> ((0) & 63))] = 42
                                                _ = v85
                                            } else {
                                                var v86 float32 = 1.5
                                                v3 -= (-(v3))
                                                _ = v86
                                            }
                                            v2 = (^(fi44))
                                            switch 101 {
                                            case 4294967295:
                                                v40 += (^(v4))
                                            default:
                                                if (false || (1 != 957)) {
                                                    var v87 float64 = 3.14159
                                                    var v88 float64 = v87
                                                    _ = v87
                                                    _ = v88
                                                }
                                            }
                                        } else {
                                            switch v39 {
                                            case "":
                                                v89 := v1[745]
                                                _ = v89
                                            case "\n\t":
                                                if (true && (-1 > -1)) {
                                                    if false {
                                                        switch fi44 {
                                                        case -810:
                                                            v90 := v1[fi45]
                                                            _ = v90
                                                        case -2147483648:
                                                            var v91 float64 = float64(-2.25)
                                                            _ = v91
                                                        default:
                                                            if true {
                                                                v2 = (fi45 << ((255) & 63))
                                                            } else {
                                                                var v92 float64 = (0.1 + ((3.14159 + 0.1) * 0.0))
                                                                fmt.Println(fi44)
                                                                v93 := GenS9{F0: true}
                                                                _ = v92
                                                                _ = v93
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    var v94 int32 = (^(v3))
                                                    if (true && true) {
                                                        switch v2 {
                                                        case 128:
                                                            if (false || true) {
                                                                v4 = v40
                                                                _ = len(v1)
                                                            }
                                                        default:
                                                            if false {
                                                                var v95 int = fi44
                                                                switch -366 {
                                                                case 527:
                                                                    v40 = 42
                                                                case 226:
                                                                    v4 += v40
                                                                case 9223372036854775807:
                                                                    var v96 uint8 = 169
                                                                    _ = v96
                                                                default:
                                                                    if false {
                                                                        delete(v1, (-(fi45)))
                                                                    }
                                                                }
                                                                v4 -= v4
                                                                _ = v95
                                                            } else {
                                                                var v97 int8 = -128
                                                                v40 -= genFn0()
                                                                v1[-726] = 2147483647
                                                                _ = v97
                                                            }
                                                        }
                                                        v98 := GenS10{F0: true}
                                                        _ = v98
                                                    }
                                                    var v99 int = (fi45 << ((589) & 63))
                                                    var v100 int32 = v3
                                                    v2 -= v2
                                                    var v101 int32 = (-(v100))
                                                    if len(v5) > 0 {
                                                        v5[0] = (v39 + "hello")
                                                    }
                                                    _ = v100
                                                    _ = v101
                                                    _ = v94
                                                    _ = v99
                                                }
                                            }
                                        }
                                    }
                                }
                                _ = v60
                            }
                        case "":
                            v102 := GenS4{F0: -9223372036854775808, F1: 307}
                            _ = v102
                        }
                        _ = fi45
                        _ = v58
                    }
                    var v103 uint64 = (v4 | v40)
                    fmt.Println(v4)
                    _ = fi44
                    _ = v103
                }
            }
            v4 = 2
        } else {
            for fi104 := 0; fi104 < 1; fi104++ {
                var v105 string = v39
                var v106 int = v2
                _ = fi104
                _ = v105
                _ = v106
            }
            v40 = v40
            v4 += genFn1(964)
        }
    case "foo bar":
        for fi107 := 0; fi107 < 16; fi107++ {
            if true {
                for fi108 := 0; fi108 < 4; fi108++ {
                    v39 = "a"
                    _ = fi108
                }
                v2 = v2
            } else {
                v109 := v1[(v2 >> ((42) & 63))]
                var v110 int32 = v3
                _ = v109
                _ = v110
            }
            _ = fi107
        }
    }
    v2 += (-(v2))
    var v111 uint = 201
    for fi112 := 0; fi112 < 8; fi112++ {
        v113 := GenS11{F0: fi112}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi112
        _ = v113
    }
    v40 -= v40
    switch (v2 ^ 437) {
    case 1:
        v114 := GenS12{F0: 42, F1: v111}
        _ = v114
    case -707:
        fmt.Println(v3, v40, v4)
    }
    for fi115 := 0; fi115 < 15; fi115++ {
        var v116 float64 = 0.0
        _ = fi115
        _ = v116
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v111:", v111)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v111
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
