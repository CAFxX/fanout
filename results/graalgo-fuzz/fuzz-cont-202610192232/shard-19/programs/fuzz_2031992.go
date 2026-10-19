package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 string; F1 uint }
type GenS8 struct { F0 uint32; F1 uint16 }
type GenS9 struct { F0 int16; F1 bool; F2 int8 }
type GenS10 struct { F0 int32; F1 int8 }
type GenS11 struct { F0 int; F1 int; F2 int }
type GenS12 struct { F0 int; F1 int8 }
type GenS13 struct { F0 bool; F1 uint64 }
type GenS14 struct { F0 int; F1 int8; F2 int16 }
type GenS15 struct { F0 int; F1 uint16 }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 int16; F1 string }

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]bool)
    var v3 float32 = 1e10
    var v4 float64 = 0.0
    var v5 uint32 = 0
    if (42 >= 1) {
        fmt.Println(v5, v3, v1)
        for fi6 := 0; fi6 < 12; fi6++ {
            switch (fi6 + (fi6 << ((53) & 63))) {
            case 101:
                var v7 int32 = 2147483647
                _ = v7
            case 9223372036854775807:
                if false {
                    if (v1 <= -107) {
                        if false {
                            v8 := []int{}
                            if (v4 <= -2.25) {
                                switch "a" {
                                case "foo bar":
                                    v9 := GenS1{F0: 42, F1: 0}
                                    _ = v9
                                case "hello":
                                    var v10 uint = 65535
                                    _ = v10
                                }
                                var v11 int64 = 42
                                v11 -= 578
                                _ = v11
                            } else {
                                if true {
                                    switch fi6 {
                                    case -1:
                                        v1 -= 1
                                    case 42:
                                        if false {
                                            if true {
                                                v12 := GenS1{F0: 390, F1: 18446744073709551615}
                                                var v13 uint64 = 852
                                                v8 = append(v8, -2147483648)
                                                _ = v12
                                                _ = v13
                                            }
                                            var v14 uint = 2
                                            _ = v14
                                        }
                                    case 661:
                                        var v15 uint16 = 32767
                                        _ = v15
                                    }
                                    switch (^(fi6)) {
                                    case 47:
                                        v4 = 1.5
                                    default:
                                        v16 := GenS1{F0: -455, F1: 573}
                                        _ = v16
                                    }
                                } else {
                                    v8 = append(v8, 128)
                                    for fi17 := 0; fi17 < 10; fi17++ {
                                        if len(v8) > 0 {
                                            v8[0] = fi17
                                        }
                                        _ = fi17
                                    }
                                    for fi18 := 0; fi18 < 14; fi18++ {
                                        _ = len(v2)
                                        var v19 int8 = 42
                                        _ = fi18
                                        _ = v19
                                    }
                                }
                                v1 = 42
                                v1 = 127
                            }
                            _ = v8
                        } else {
                            for fi20 := 0; fi20 < 18; fi20++ {
                                if (("foo bar" + "hello") > "") {
                                    v5 = 400
                                    switch -441 {
                                    case 42:
                                        var v21 int16 = 150
                                        _ = v21
                                    default:
                                        v3 = 0.0
                                    }
                                    v22 := []int{74}
                                    _ = v22
                                } else {
                                    if (-2147483648 >= fi6) {
                                        switch fi20 {
                                        case 447:
                                            if (857 < 9223372036854775807) {
                                                v23 := GenS1{F0: 60, F1: 0}
                                                fmt.Println(v3, v3, v3)
                                                v24 := GenS2{F0: 42, F1: (v5 / (((v5 - (v5 / ((v5) | 1)))) | 1))}
                                                _ = v23
                                                _ = v24
                                            } else {
                                                var v25 int64 = 0
                                                if (true && (!true)) {
                                                    var v26 uint8 = 0
                                                    v27 := GenS3{F0: v1}
                                                    _ = v26
                                                    _ = v27
                                                } else {
                                                    if false {
                                                        var v28 int16 = 32767
                                                        var v29 uint8 = 0
                                                        var v30 int32 = 183
                                                        _ = v28
                                                        _ = v29
                                                        _ = v30
                                                    }
                                                }
                                                _ = v25
                                            }
                                        default:
                                            var v31 int8 = v1
                                            if true {
                                                v32 := GenS3{F0: (v1 << ((187) & 63))}
                                                if true {
                                                    var v33 string = ""
                                                    v4 = 1e10
                                                    v31 = (v31 >> ((2) & 63))
                                                    _ = v33
                                                }
                                                if true {
                                                    var v34 uint64 = 18446744073709551615
                                                    switch ("foo bar" + ("a" + "a")) {
                                                    case "":
                                                        if (float32((v4 * 1e10)) != float32(-2.25)) {
                                                            var v35 uint32 = (v5 / ((4294967295) | 1))
                                                            v34 = (v34 / ((v34) | 1))
                                                            var v36 string = "hello"
                                                            _ = v35
                                                            _ = v36
                                                        }
                                                    }
                                                    fmt.Println(fi6)
                                                    _ = v34
                                                } else {
                                                    v3 += float32(v4)
                                                    v5 -= v5
                                                }
                                                _ = v32
                                            }
                                            _ = v31
                                        }
                                        v37 := GenS2{F0: 65535, F1: (v5 - v5)}
                                        fmt.Println(fi20, v5)
                                        _ = v37
                                    }
                                    var v38 int = fi6
                                    if (true || false) {
                                        var v39 float64 = v4
                                        v4 += 1e10
                                        v40 := GenS4{F0: v1}
                                        _ = v39
                                        _ = v40
                                    } else {
                                        v2[v5] = true
                                        v38 -= fi6
                                        v5 += v5
                                    }
                                    _ = v38
                                }
                                var v41 uint8 = 37
                                _ = fi20
                                _ = v41
                            }
                            if true {
                                _ = len(v2)
                                switch fi6 {
                                case -969:
                                    fmt.Println(v3)
                                case -151:
                                    var v42 bool = false
                                    _ = v42
                                default:
                                    v43 := GenS1{F0: -1, F1: 255}
                                    _ = v43
                                }
                            }
                            if ((1 <= 42) || false) {
                                fmt.Println(v3, fi6)
                                v44 := GenS5{F0: 926}
                                switch (fi6 ^ (fi6 * 0)) {
                                case 0:
                                    fmt.Println(v5, v4, v4)
                                }
                                _ = v44
                            }
                        }
                        if (!(!false)) {
                            for fi45 := 0; fi45 < 2; fi45++ {
                                if ((10 <= 661) && (!(fi6 != (fi6 ^ fi6)))) {
                                    fmt.Println(v1, v3, v1)
                                } else {
                                    fmt.Println(v5)
                                    fmt.Println(v4, v4, fi6)
                                    if (808 >= 48) {
                                        v3 = 1.5
                                        v5 -= v5
                                    } else {
                                        var v46 uint32 = (v5 * (v5 * 42))
                                        v47 := GenS6{F0: 42, F1: 18446744073709551615}
                                        v48 := GenS5{F0: 9223372036854775807}
                                        _ = v46
                                        _ = v47
                                        _ = v48
                                    }
                                    var v49 int8 = (v1 * v1)
                                    v50 := GenS7{F0: "a", F1: 15}
                                    _ = v49
                                    _ = v50
                                }
                                v51 := GenS2{F0: 255, F1: (v5 / ((1) | 1))}
                                v1 = (v1 / (((v1 % ((v1) | 1))) | 1))
                                var v52 uint64 = 0
                                _ = fi45
                                _ = v51
                                _ = v52
                            }
                            v53 := GenS1{F0: -647, F1: 212}
                            _ = v53
                        }
                    }
                    v1 -= (v1 / ((v1) | 1))
                    switch (^(fi6)) {
                    case -159:
                        v3 += v3
                    }
                }
            case 464:
                for fi54 := 0; fi54 < 19; fi54++ {
                    v1 = -128
                    if (false || true) {
                        var v55 int16 = -47
                        if (1 == 0) {
                            v56 := []int{127, 120}
                            var v57 int = fi54
                            v55 -= v55
                            _ = v56
                            _ = v57
                        } else {
                            v58 := GenS7{F0: "\n\t", F1: 255}
                            v1 -= (v1 + (v1 * (v1 >> ((0) & 63))))
                            switch -1 {
                            case 127:
                                switch 756 {
                                case -2147483648:
                                    fmt.Println(v4, fi6, v3)
                                case 0:
                                    v59 := []int{242, -2147483648}
                                    _ = v59
                                case -1:
                                    v5 += 4294967295
                                default:
                                    v5 -= v5
                                }
                            case 1:
                                v60 := GenS8{F0: 654, F1: 255}
                                _ = v60
                            case -2147483648:
                                _ = len(v2)
                            default:
                                v2[v5] = (42 != 0)
                            }
                            _ = v58
                        }
                        _ = v55
                    } else {
                        fmt.Println(v5)
                    }
                    _ = fi54
                }
            default:
                switch 42 {
                case -661:
                    var v61 float32 = v3
                    _ = v61
                case 9223372036854775807:
                    for fi62 := 0; fi62 < 20; fi62++ {
                        var v63 int32 = 64
                        delete(v2, v5)
                        _ = fi62
                        _ = v63
                    }
                case 4294967295:
                    if true {
                        for fi64 := 0; fi64 < 12; fi64++ {
                            fmt.Println(fi64, fi64, v3)
                            v65 := v2[(v5 & v5)]
                            var v66 uint16 = 2
                            _ = fi64
                            _ = v65
                            _ = v66
                        }
                    } else {
                        v67 := GenS9{F0: -32768, F1: ("foo bar" >= (("a" + "hello") + "hello")), F2: v1}
                        fmt.Println(v4)
                        v4 -= (v4 * v4)
                        _ = v67
                    }
                default:
                    switch "\n\t" {
                    case "\n\t":
                        v68 := GenS10{F0: 709, F1: (v1 / (((v1 | v1)) | 1))}
                        _ = v68
                    }
                }
            }
            v69 := GenS1{F0: 178, F1: 1}
            v70 := GenS6{F0: 120, F1: 1}
            _ = fi6
            _ = v69
            _ = v70
        }
    }
    switch "hello" {
    case "hello":
        _ = len(v2)
    default:
        fmt.Println(v1, v4, v4)
    }
    fmt.Println(v3)
    fmt.Println(v4)
    _ = len(v2)
    var v71 int16 = 259
    if (4294967295 == 4294967295) {
        v72 := []int{-2147483648}
        _ = v72
    }
    var v73 int16 = (v71 % ((v71) | 1))
    for fi74 := 0; fi74 < 9; fi74++ {
        v4 = (1e10 / v4)
        if true {
            if (-1 >= 42) {
                if true {
                    v75 := GenS9{F0: v73, F1: (0 >= 18), F2: v1}
                    var v76 uint32 = (v5 | v5)
                    _ = v75
                    _ = v76
                }
            }
        }
        _ = fi74
    }
    if false {
        if false {
            v77 := GenS11{F0: 416, F1: -9223372036854775808, F2: 128}
            v2[4294967295] = true
            _ = v77
        }
        fmt.Println(v1, v4)
        for fi78 := 0; fi78 < 4; fi78++ {
            fmt.Println(v71, fi78)
            v79 := GenS12{F0: (^(fi78)), F1: v1}
            _ = fi78
            _ = v79
        }
    }
    v80 := []int{32767, 2147483647}
    var v81 int16 = -346
    switch "a" {
    case "foo bar":
        if true {
            v82 := GenS13{F0: (true && true), F1: 18446744073709551615}
            if ((v81 % (((v81 ^ v81)) | 1)) != 42) {
                v83 := GenS14{F0: 752, F1: v1, F2: (v73 | v73)}
                for fi84 := 0; fi84 < 10; fi84++ {
                    switch (fi84 * fi84) {
                    case 207:
                        if ((^(v5)) > (^(v5))) {
                            if ((!(v3 != 1.5)) || (uint(v5) == 4294967295)) {
                                v85 := GenS10{F0: -1, F1: v1}
                                if false {
                                    if (false && false) {
                                        var v86 uint64 = 609
                                        _ = v86
                                    } else {
                                        fmt.Println(v5)
                                        var v87 int16 = v71
                                        v88 := GenS15{F0: 2147483647, F1: 2}
                                        _ = v87
                                        _ = v88
                                    }
                                    v89 := GenS3{F0: (v1 + v1)}
                                    _ = v89
                                }
                                var v90 int64 = -9223372036854775808
                                _ = v85
                                _ = v90
                            } else {
                                var v91 int = int(v5)
                                fmt.Println(v91, v81)
                                _ = v91
                            }
                        }
                    }
                    _ = fi84
                }
                _ = v83
            } else {
                var v92 float32 = float32((v4 / v4))
                if false {
                    v80 = append(v80, 836)
                    v93 := GenS16{F0: v73}
                    var v94 int8 = (v1 ^ (v1 + (v1 | -101)))
                    _ = v93
                    _ = v94
                } else {
                    var v95 string = ("foo bar" + ("foo bar" + ((("hello" + "a") + "") + ("foo bar" + "hello"))))
                    _ = v95
                }
                _ = v92
            }
            _ = v82
        }
    case "\n\t":
        switch "\n\t" {
        case "":
            for fi96 := 0; fi96 < 17; fi96++ {
                v2[(^(v5))] = (65535 < 255)
                var v97 int = 2147483647
                for fi98 := 0; fi98 < 20; fi98++ {
                    var v99 uint32 = 534
                    switch (fi96 & 128) {
                    case 674:
                        var v100 uint16 = 68
                        _ = v100
                    default:
                        var v101 float64 = 1.5
                        _ = v101
                    }
                    _ = fi98
                    _ = v99
                }
                _ = fi96
                _ = v97
            }
        case "\n\t":
            var v102 int64 = 42
            _ = v102
        case "hello":
            var v103 int32 = -431
            _ = v103
        default:
            var v104 uint = 158
            _ = v104
        }
    case "":
        v105 := GenS9{F0: (v73 >> ((2) & 63)), F1: (!(4294967295 != 9223372036854775807)), F2: v1}
        _ = v105
    default:
        for fi106 := 0; fi106 < 18; fi106++ {
            fmt.Println(v81, v71, v81)
            var v107 int32 = 42
            _ = fi106
            _ = v107
        }
    }
    v108 := GenS3{F0: (^(v1))}
    for fi109 := 0; fi109 < 9; fi109++ {
        v110 := GenS4{F0: (v1 + (v1 << ((258) & 63)))}
        v111 := GenS9{F0: v81, F1: true, F2: int8(v81)}
        if (((true && (255 == 1)) && (0 != 614)) || true) {
            v112 := GenS17{F0: -742, F1: ("" + ("" + ""))}
            v5 = (v5 ^ 2147483647)
            _ = v112
        } else {
            var v113 int = fi109
            _ = v113
        }
        _ = fi109
        _ = v110
        _ = v111
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v71:", v71)
    fmt.Println("v73:", v73)
    fmt.Println("slicelen v80:", len(v80), v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v108
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v71
    _ = v73
    _ = v80
    _ = v81
}
