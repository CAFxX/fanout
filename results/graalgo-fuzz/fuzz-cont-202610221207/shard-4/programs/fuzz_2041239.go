package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int32; F2 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 string; F1 uint32 }
type GenS9 struct { F0 int16; F1 int }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint8; F1 int8; F2 string }
type GenS13 struct { F0 uint8 }

func main() {
    var v1 int32 = 42
    v2 := []int{-2147483648, 0}
    v3 := []uint32{1, 4294967295, 65535}
    var v4 uint = 2
    switch -2147483648 {
    case 128:
        v5 := GenS1{F0: 850}
        _ = v5
    case -1:
        for fi6 := 0; fi6 < 3; fi6++ {
            fmt.Println(fi6, fi6, v4)
            var v7 int = (fi6 % ((fi6) | 1))
            _ = fi6
            _ = v7
        }
    default:
        var v8 string = ("\n\t" + (("a" + "foo bar") + ("a" + "hello")))
        _ = v8
    }
    v9 := GenS2{F0: (-(v1)), F1: v1, F2: 464}
    v4 -= (v4 % ((v4) | 1))
    if ((v1 | -1) > 283) {
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 float32 = 1.5
            for fi12 := 0; fi12 < 8; fi12++ {
                v4 = v4
                if false {
                    if false {
                        if true {
                            v13 := GenS1{F0: 1}
                            if true {
                                if false {
                                    switch 42 {
                                    case -9223372036854775808:
                                        var v14 float64 = (((3.14159 + float64(v11)) + 0.1) / 3.14159)
                                        _ = v14
                                    case 9223372036854775807:
                                        v1 -= v1
                                    case 0:
                                        var v15 int32 = (^(v1))
                                        _ = v15
                                    }
                                    var v16 int16 = 1000
                                    switch "hello" {
                                    case "\n\t":
                                        v16 = (v16 ^ (v16 + (v16 + 280)))
                                    case "foo bar":
                                        var v17 int8 = -128
                                        _ = v17
                                    }
                                    _ = v16
                                }
                                v1 = (v1 << (((v4 ^ 943)) & 63))
                            } else {
                                if (fi10 > (fi12 * fi12)) {
                                    fmt.Println(v4, fi12, v11)
                                    var v18 uint32 = 4294967295
                                    v2 = append(v2, (fi10 | (^(fi10))))
                                    _ = v18
                                } else {
                                    switch ("hello" + "a") {
                                    case "foo bar":
                                        v1 = v1
                                    case "":
                                        var v19 bool = (!true)
                                        switch (fi12 % ((2147483647) | 1)) {
                                        case 32768:
                                            v1 = v1
                                            if (9223372036854775807 != -234) {
                                                switch ("hello" + ("hello" + "a")) {
                                                case "foo bar":
                                                    v20 := GenS3{F0: 1}
                                                    _ = v20
                                                case "hello":
                                                    if len(v2) > 0 {
                                                        v2[0] = -149
                                                    }
                                                case "a":
                                                    switch 793 {
                                                    case 0:
                                                        switch fi10 {
                                                        case 0:
                                                            var v21 float64 = 3.14159
                                                            _ = v21
                                                        default:
                                                            _ = len(v2)
                                                        }
                                                    case -1:
                                                        switch "a" {
                                                        case "":
                                                            fmt.Println(fi12, v19, fi10)
                                                        case "hello":
                                                            v3 = append(v3, 4294967295)
                                                        default:
                                                            v22 := GenS2{F0: (v1 - v1), F1: (v1 % (((v1 << ((1) & 63))) | 1)), F2: -9223372036854775808}
                                                            _ = v22
                                                        }
                                                    }
                                                }
                                                var v23 int32 = v1
                                                _ = v23
                                            } else {
                                                var v24 string = "hello"
                                                _ = v24
                                            }
                                            if false {
                                                var v25 int8 = -56
                                                var v26 uint16 = 255
                                                if (!(!false)) {
                                                    v27 := GenS2{F0: (v1 - -2147483648), F1: 194, F2: fi10}
                                                    _ = v27
                                                } else {
                                                    var v28 int32 = (-(v1))
                                                    if (42 <= 968) {
                                                        var v29 int = 1
                                                        if (-1 >= 0) {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        _ = v29
                                                    } else {
                                                        v25 = v25
                                                        fmt.Println(v4, v4, fi12)
                                                    }
                                                    var v30 int32 = (v28 * v28)
                                                    _ = v28
                                                    _ = v30
                                                }
                                                fmt.Println(fi12)
                                                _ = v25
                                                _ = v26
                                            } else {
                                                v3 = append(v3, 2147483647)
                                                if true {
                                                    if (241 < 0) {
                                                        switch 397 {
                                                        case 127:
                                                            v11 = v11
                                                        case -311:
                                                            var v31 int16 = 424
                                                            _ = v31
                                                        }
                                                        var v32 string = "\n\t"
                                                        _ = v32
                                                    }
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    v11 -= (v11 + ((1.5 * v11) / v11))
                                                }
                                            }
                                            var v33 int32 = v1
                                            _ = v33
                                        case -9223372036854775808:
                                            v34 := GenS2{F0: (v1 & v1), F1: 2147483647, F2: fi10}
                                            _ = v34
                                        case 290:
                                            var v35 int = 635
                                            _ = v35
                                        default:
                                            var v36 float64 = 0.0
                                            _ = v36
                                        }
                                        _ = v19
                                    }
                                    var v37 int8 = -95
                                    fmt.Println(v4, v4)
                                    _ = v37
                                }
                            }
                            _ = v13
                        } else {
                            fmt.Println(v11, fi10)
                            if true {
                                v11 = 1.5
                                var v38 float64 = 1e10
                                if (-710 >= 9223372036854775807) {
                                    v39 := GenS4{F0: 373}
                                    var v40 int = (fi10 / ((42) | 1))
                                    _ = v39
                                    _ = v40
                                } else {
                                    var v41 int16 = -817
                                    _ = v41
                                }
                                _ = v38
                            }
                            if false {
                                switch ("\n\t" + ((("\n\t" + "foo bar") + "hello") + "a")) {
                                case "\n\t":
                                    v42 := GenS5{F0: 2147483648}
                                    _ = v42
                                default:
                                    v11 += (v11 * v11)
                                    v11 -= (((-2.25 + ((1.5 - v11) / -2.25)) * v11) - (0.0 - v11))
                                    v4 -= (v4 % (((v4 << (((v4 >> ((0) & 63))) & 63))) | 1))
                                }
                                switch 318 {
                                case -397:
                                    var v43 bool = false
                                    _ = v43
                                }
                                if false {
                                    v4 = v4
                                } else {
                                    v4 = 900
                                }
                            }
                            var v44 uint = (v4 >> ((336) & 63))
                            _ = v44
                        }
                        if ((v11 * 1.5) > 3.14159) {
                            var v45 uint64 = 42
                            _ = v45
                        }
                    } else {
                        v46 := GenS6{F0: 0}
                        var v47 uint32 = 849
                        _ = v46
                        _ = v47
                    }
                    v48 := GenS7{F0: 0, F1: 0}
                    _ = len(v2)
                    _ = v48
                } else {
                    switch fi10 {
                    case 816:
                        var v49 int16 = 707
                        _ = v49
                    case -711:
                        v50 := GenS8{F0: "\n\t", F1: 2147483647}
                        switch fi12 {
                        case 9223372036854775807:
                            fmt.Println(fi12, v1)
                        case 32768:
                            v4 -= (v4 | v4)
                        }
                        var v51 bool = false
                        _ = v50
                        _ = v51
                    }
                }
                var v52 int32 = v1
                _ = fi12
                _ = v52
            }
            v53 := GenS9{F0: 1, F1: (fi10 | fi10)}
            _ = fi10
            _ = v11
            _ = v53
        }
        var v54 uint32 = 2
        _ = v54
    }
    for fi55 := 0; fi55 < 10; fi55++ {
        v56 := GenS8{F0: ("\n\t" + ""), F1: 0}
        for fi57 := 0; fi57 < 5; fi57++ {
            v58 := GenS4{F0: 2}
            _ = fi57
            _ = v58
        }
        switch 0 {
        case -437:
            var v59 bool = (!(!true))
            _ = v59
        case 32767:
            v60 := GenS10{F0: (v1 | v1)}
            _ = v60
        default:
            for fi61 := 0; fi61 < 2; fi61++ {
                _ = len(v3)
                if ("foo bar" > "foo bar") {
                    switch "hello" {
                    case "":
                        v62 := GenS3{F0: 127}
                        if false {
                            var v63 int64 = -821
                            v64 := GenS11{F0: 1}
                            _ = v63
                            _ = v64
                        }
                        v1 -= v1
                        _ = v62
                    case "hello":
                        var v65 uint = v4
                        _ = v65
                    }
                    v1 = 835
                }
                fmt.Println(v1, fi61, v1)
                var v66 string = "a"
                _ = fi61
                _ = v66
            }
        }
        _ = fi55
        _ = v56
    }
    var v67 float64 = float64((0.0 + ((-2.25 + (1e10 + (3.14159 + (-2.25 + 1e10)))) + (((((3.14159 / 1e10) / 3.14159) + (3.14159 / (((0.1 + 0.0) / 3.14159) + 3.14159))) * ((0.0 - 0.0) + float32(-2.25))) / 0.1))))
    for fi68 := 0; fi68 < 7; fi68++ {
        var v69 int = fi68
        var v70 bool = true
        v4 = v4
        _ = fi68
        _ = v69
        _ = v70
    }
    if true {
        v4 += (^(v4))
    }
    if true {
        switch 42 {
        case -1:
            var v71 uint64 = 1
            _ = v71
        case 29:
            if (true || (32767 >= 1000)) {
                v4 = v4
                for fi72 := 0; fi72 < 12; fi72++ {
                    for fi73 := 0; fi73 < 9; fi73++ {
                        if false {
                            var v74 float64 = (float64(float32((v67 - v67))) + (v67 - (0.0 - -2.25)))
                            _ = v74
                        }
                        v75 := GenS3{F0: 128}
                        _ = fi73
                        _ = v75
                    }
                    if true {
                        var v76 int64 = 0
                        if len(v3) > 0 {
                            v3[0] = 980
                        }
                        for fi77 := 0; fi77 < 3; fi77++ {
                            fmt.Println(v67, fi77)
                            if (42 == 0) {
                                if false {
                                    fmt.Println(v4, fi72, v1)
                                } else {
                                    v1 = (v1 + v1)
                                    switch ("foo bar" + "hello") {
                                    case "foo bar":
                                        var v78 uint = 65535
                                        _ = v78
                                    case "\n\t":
                                        v76 += v76
                                    }
                                    v79 := GenS12{F0: 2, F1: -128, F2: "\n\t"}
                                    _ = v79
                                }
                                var v80 uint32 = 51
                                fmt.Println(v67)
                                v4 += (v4 - 0)
                                _ = v80
                            }
                            _ = fi77
                        }
                        _ = v76
                    }
                    _ = fi72
                }
                var v81 uint8 = 172
                _ = v81
            } else {
                var v82 uint8 = 140
                var v83 uint64 = 4294967295
                for fi84 := 0; fi84 < 6; fi84++ {
                    var v85 int64 = 1
                    if (!true) {
                        switch "foo bar" {
                        case "hello":
                            v86 := GenS13{F0: 1}
                            _ = v86
                        case "":
                            fmt.Println(v1, v1, v85)
                        default:
                            if false {
                                if (float32(1.5) > 3.14159) {
                                    fmt.Println(v67)
                                    v67 = v67
                                }
                                var v87 uint32 = 65535
                                switch ("" + "\n\t") {
                                case "a":
                                    var v88 uint64 = v83
                                    _ = v88
                                default:
                                    var v89 int64 = (v85 - (-(v85)))
                                    _ = v89
                                }
                                _ = v87
                            } else {
                                var v90 int64 = (v85 & (-(v85)))
                                var v91 int32 = (-(v1))
                                switch "foo bar" {
                                case "a":
                                    v82 = (v82 >> ((462) & 63))
                                default:
                                    v85 -= (v90 + 0)
                                }
                                _ = v90
                                _ = v91
                            }
                        }
                        fmt.Println(v4, v85, fi84)
                        fmt.Println(v1, v85)
                    }
                    _ = fi84
                    _ = v85
                }
                _ = v82
                _ = v83
            }
        case -302:
            fmt.Println(v67, v1, v67)
        }
        switch 42 {
        case 32767:
            v92 := GenS1{F0: 9223372036854775807}
            _ = v92
        case 128:
            switch 912 {
            case 601:
                fmt.Println(v67)
            case -696:
                v2 = append(v2, -613)
            }
        default:
            v1 -= v1
        }
    }
    if (!(128 > 42)) {
        if false {
            var v93 bool = true
            var v94 int = -843
            _ = v93
            _ = v94
        }
    }
    fmt.Println(v1, v4)
    if (14 <= 181) {
        for fi95 := 0; fi95 < 16; fi95++ {
            var v96 uint32 = 1
            _ = fi95
            _ = v96
        }
    } else {
        var v97 int16 = 1000
        v2 = append(v2, 128)
        v97 = v97
        _ = v97
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v67
    _ = v9
}
