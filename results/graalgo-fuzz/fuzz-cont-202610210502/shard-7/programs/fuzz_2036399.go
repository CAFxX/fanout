package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 int8; F1 int; F2 uint }
type GenS6 struct { F0 bool; F1 int32; F2 uint32 }
type GenS7 struct { F0 uint64; F1 uint64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS10 struct { F0 int8; F1 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{}
    var v3 int8 = -121
    v4 := make(map[int]string)
    var v5 int = -915
    var v6 uint = 65535
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 9223372036854775807, F1: 65535}
        var v9 int16 = 0
        if len(v2) > 0 {
            v2[0] = ((((("hello" + "a") + ("\n\t" + ((("a" + "hello") + ("hello" + "")) + "foo bar"))) + ("" + ("\n\t" + ("\n\t" + "foo bar")))) + "a") + (((("foo bar" + "") + ((("foo bar" + "hello") + ("\n\t" + ("a" + ("foo bar" + "")))) + "hello")) + ("hello" + "")) + ""))
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v5)
    var v10 int64 = 739
    v3 = v3
    switch v5 {
    case 128:
        var v11 uint16 = 121
        _ = v11
    case 1:
        for fi12 := 0; fi12 < 20; fi12++ {
            v6 += v6
            _ = fi12
        }
    case 4294967295:
        for fi13 := 0; fi13 < 6; fi13++ {
            for fi14 := 0; fi14 < 2; fi14++ {
                v15 := v4[(v5 | (v5 * 128))]
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
    }
    fmt.Println(v3)
    v10 += v10
    if (!false) {
        v6 -= v6
    } else {
        v16 := GenS2{F0: 1, F1: 127, F2: 151}
        _ = v16
    }
    v17 := GenS1{F0: 42, F1: v6}
    for fi18 := 0; fi18 < 4; fi18++ {
        _ = len(v4)
        _ = fi18
    }
    _ = len(v4)
    var v19 float32 = 1.5
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 uint16 = 1
        _ = len(v2)
        switch "\n\t" {
        case "hello":
            var v22 uint32 = 2147483648
            _ = v22
        case "a":
            for fi23 := 0; fi23 < 16; fi23++ {
                var v24 uint = uint(v21)
                fmt.Println(v10, v24, v24)
                var v25 int32 = 42
                v21 += 32767
                v26 := GenS3{F0: v5}
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
        case "\n\t":
            if (2 < 127) {
                v3 -= 16
                if (((v21 > (^(v21))) || (182 == 1)) || false) {
                    v19 = (v19 / float32(1.5))
                    v27 := GenS4{F0: 127, F1: v21}
                    fmt.Println(v6, v5, fi20)
                    _ = v27
                }
            }
        }
        _ = fi20
        _ = v21
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        if true {
            for fi29 := 0; fi29 < 20; fi29++ {
                if false {
                    fmt.Println(fi29, v5)
                    switch "\n\t" {
                    case "a":
                        v30 := GenS4{F0: -128, F1: 0}
                        _ = v30
                    default:
                        var v31 uint16 = 32767
                        _ = v31
                    }
                }
                v10 = v10
                _ = len(v2)
                v32 := GenS4{F0: 127, F1: 32767}
                v33 := GenS2{F0: 337, F1: 2, F2: 0}
                v34 := GenS4{F0: v3, F1: 661}
                _ = fi29
                _ = v32
                _ = v33
                _ = v34
            }
            switch "hello" {
            case "\n\t":
                if true {
                    for fi35 := 0; fi35 < 12; fi35++ {
                        v5 = 128
                        var v36 uint = 255
                        if (767 != 18446744073709551615) {
                            if (!true) {
                                v37 := GenS5{F0: (v3 ^ (v3 + v3)), F1: (v5 / ((807) | 1)), F2: 65535}
                                _ = v37
                            }
                            switch fi28 {
                            case -968:
                                switch "" {
                                case "":
                                    var v38 float64 = ((1e10 * 0.1) + -2.25)
                                    _ = v38
                                default:
                                    switch "\n\t" {
                                    case "a":
                                        if (!false) {
                                            switch (fi35 - -1) {
                                            case -1:
                                                var v39 float32 = v19
                                                _ = v39
                                            case 128:
                                                var v40 uint16 = 2
                                                _ = v40
                                            case 543:
                                                var v41 float32 = v19
                                                _ = v41
                                            default:
                                                var v42 float32 = (v19 + (1e10 - v19))
                                                if (127 <= 0) {
                                                    switch "a" {
                                                    case "foo bar":
                                                        fmt.Println(v10, fi35, v19)
                                                    }
                                                }
                                                var v43 bool = false
                                                _ = v42
                                                _ = v43
                                            }
                                        } else {
                                            v44 := GenS4{F0: (v3 ^ int8(v5)), F1: 42}
                                            var v45 int64 = v10
                                            _ = v44
                                            _ = v45
                                        }
                                    case "\n\t":
                                        switch "foo bar" {
                                        case "a":
                                            var v46 int64 = 42
                                            _ = v46
                                        case "hello":
                                            fmt.Println(v3, fi28, v36)
                                        default:
                                            var v47 uint16 = 255
                                            _ = v47
                                        }
                                    case "foo bar":
                                        v3 = v3
                                    }
                                }
                            default:
                                v48 := GenS6{F0: (true || (!true)), F1: -2147483648, F2: 1}
                                _ = v48
                            }
                        }
                        _ = fi35
                        _ = v36
                    }
                }
            case "foo bar":
                v49 := GenS5{F0: -128, F1: (fi28 ^ (fi28 << ((v6) & 63))), F2: (v6 | (v6 + (v6 & v6)))}
                _ = v49
            default:
                var v50 int32 = -83
                _ = v50
            }
        } else {
            v51 := GenS3{F0: v5}
            delete(v1, 2147483648)
            if (((v19 == 1e10) || false) || true) {
                switch -389 {
                case 1:
                    if false {
                        var v52 int64 = 787
                        v52 -= 1
                        v53 := GenS5{F0: (v3 + v3), F1: fi28, F2: 255}
                        _ = v52
                        _ = v53
                    }
                case 4294967295:
                    switch -1 {
                    case -2147483648:
                        fmt.Println(v5, v5, v19)
                    }
                case 455:
                    v6 = v6
                }
                switch (fi28 << ((18446744073709551615) & 63)) {
                case 1:
                    if (245 == 1) {
                        v19 = (3.14159 - (float32((((0.1 + ((-2.25 + (0.0 + 1.5)) - 3.14159)) - 0.0) / 0.1)) * float32(1e10)))
                    }
                case 448:
                    if ((994 > v6) && false) {
                        v54 := GenS7{F0: 50, F1: 655}
                        v55 := GenS4{F0: -75, F1: 271}
                        _ = v54
                        _ = v55
                    } else {
                        fmt.Println(v10)
                        for fi56 := 0; fi56 < 6; fi56++ {
                            v3 = (^(v3))
                            v10 = 610
                            var v57 uint64 = 0
                            _ = fi56
                            _ = v57
                        }
                    }
                case 324:
                    v4[v5] = "\n\t"
                default:
                    if false {
                        v58 := GenS4{F0: (v3 << ((v6) & 63)), F1: 1}
                        var v59 bool = true
                        v19 = 1e10
                        _ = v58
                        _ = v59
                    } else {
                        switch "hello" {
                        case "":
                            v60 := v4[v5]
                            _ = v60
                        case "\n\t":
                            switch "foo bar" {
                            case "a":
                                for fi61 := 0; fi61 < 18; fi61++ {
                                    v62 := GenS5{F0: 100, F1: v5, F2: v6}
                                    fmt.Println(v6)
                                    fmt.Println(v3, v19)
                                    switch (fi28 % ((32768) | 1)) {
                                    case 32768:
                                        fmt.Println(v5, v19, fi61)
                                    default:
                                        fmt.Println(v5, fi28)
                                    }
                                    _ = len(v4)
                                    _ = fi61
                                    _ = v62
                                }
                            case "foo bar":
                                for fi63 := 0; fi63 < 12; fi63++ {
                                    v3 -= v3
                                    fmt.Println(v3, v10, v5)
                                    v3 = (-(v3))
                                    v5 -= v5
                                    v10 = (v10 / (((v10 ^ v10)) | 1))
                                    if (2 == 31) {
                                        var v64 int = -2147483648
                                        fmt.Println(v6, fi63, v6)
                                        _ = v64
                                    }
                                    _ = fi63
                                }
                            default:
                                var v65 int8 = -17
                                _ = v65
                            }
                        default:
                            switch 128 {
                            case 42:
                                v6 = (^(v6))
                            case 32768:
                                v19 = v19
                            default:
                                if (!false) {
                                    v1[1] = 2
                                    for fi66 := 0; fi66 < 18; fi66++ {
                                        v5 -= fi66
                                        var v67 string = ("hello" + "")
                                        _ = fi66
                                        _ = v67
                                    }
                                } else {
                                    if len(v2) > 0 {
                                        v2[0] = "hello"
                                    }
                                    for fi68 := 0; fi68 < 20; fi68++ {
                                        fmt.Println(fi68)
                                        var v69 float32 = (0.1 - float32(1e10))
                                        _ = fi68
                                        _ = v69
                                    }
                                    v6 += 18446744073709551615
                                }
                            }
                        }
                        if false {
                            v10 = 0
                            v19 = (v19 / v19)
                            v6 = (v6 << ((v6) & 63))
                        }
                        for fi70 := 0; fi70 < 5; fi70++ {
                            v10 -= v10
                            v71 := GenS8{F0: "foo bar"}
                            _ = fi70
                            _ = v71
                        }
                    }
                }
                switch "a" {
                case "a":
                    v19 = v19
                default:
                    fmt.Println(v6, v6)
                }
            } else {
                v3 = v3
                if true {
                    v72 := GenS9{F0: 1, F1: 4294967295, F2: 2}
                    _ = v72
                }
                v73 := GenS10{F0: 1, F1: 0}
                _ = v73
            }
            _ = v51
        }
        _ = len(v1)
        _ = fi28
    }
    if false {
        var v74 uint8 = 199
        _ = v74
    }
    v19 = float32((1e10 / 1e10))
    switch "\n\t" {
    case "a":
        var v75 string = (("\n\t" + "foo bar") + ("a" + "hello"))
        _ = v75
    case "\n\t":
        var v76 string = "hello"
        _ = v76
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
