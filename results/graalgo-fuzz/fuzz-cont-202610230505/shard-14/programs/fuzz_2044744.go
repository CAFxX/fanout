package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 int8; F2 bool }
type GenS6 struct { F0 int16; F1 uint32; F2 bool }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int8; F1 int32 }

func main() {
    v1 := []string{"a", ("hello" + ("\n\t" + "hello")), ("\n\t" + ("" + "a")), "foo bar"}
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    v4 := []string{}
    v5 := make(map[string]int)
    v6 := make(map[string]uint32)
    var v7 bool = (false || (2 != 158))
    var v8 uint8 = 2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v9 := GenS1{F0: 9223372036854775807, F1: 1}
    v10 := GenS1{F0: 1, F1: 0}
    if v7 {
        v11 := GenS2{F0: 2147483647}
        _ = v11
    }
    v7 = v7
    for fi12 := 0; fi12 < 9; fi12++ {
        if v7 {
            v4 = append(v4, ("" + "\n\t"))
        }
        for fi13 := 0; fi13 < 3; fi13++ {
            switch ("" + "foo bar") {
            case "foo bar":
                switch "" {
                case "a":
                    switch "a" {
                    case "hello":
                        fmt.Println(v7)
                    default:
                        var v14 uint8 = v8
                        _ = v14
                    }
                case "\n\t":
                    v15 := GenS1{F0: 1, F1: 107}
                    v16 := GenS1{F0: -9223372036854775808, F1: 42}
                    v7 = (!(float32(1.5) > 0.0))
                    _ = v15
                    _ = v16
                }
            case "a":
                _ = len(v5)
            case "":
                if true {
                    switch (-(fi13)) {
                    case 339:
                        var v17 float32 = 1e10
                        _ = v17
                    default:
                        fmt.Println(v8, v8)
                        if (-1 != 127) {
                            var v18 uint = 138
                            _ = v18
                        }
                    }
                } else {
                    v8 -= (^(v8))
                    v8 = (v8 / ((v8) | 1))
                    if (576 <= 42) {
                        v7 = true
                        if (((3.14159 - 0.0) != 1e10) || true) {
                            fmt.Println(fi13, fi12, v7)
                        } else {
                            v1 = append(v1, "\n\t")
                        }
                        if (2 <= 32767) {
                            var v19 uint16 = 65535
                            var v20 int64 = 0
                            _ = v19
                            _ = v20
                        } else {
                            v21 := GenS3{F0: v7, F1: 32767}
                            if v7 {
                                v8 += 0
                                v7 = v7
                                v7 = (v7 || v7)
                            }
                            var v22 uint = 319
                            var v23 int64 = 42
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                    } else {
                        v8 -= v8
                        switch ("hello" + (("\n\t" + "hello") + ("foo bar" + "a"))) {
                        case "foo bar":
                            if (!false) {
                                v24 := GenS3{F0: v7, F1: 32767}
                                var v25 uint32 = 1
                                fmt.Println(v25)
                                _ = v24
                                _ = v25
                            } else {
                                v26 := GenS3{F0: ("hello" < "foo bar"), F1: 42}
                                var v27 int = (fi13 << ((4294967295) & 63))
                                fmt.Println(v7, v27)
                                _ = v26
                                _ = v27
                            }
                            var v28 uint8 = v8
                            var v29 uint32 = 553
                            _ = v28
                            _ = v29
                        case "\n\t":
                            switch (fi13 - (fi12 >> ((1) & 63))) {
                            case 581:
                                var v30 int32 = 1
                                _ = v30
                            }
                        case "a":
                            var v31 int = (fi12 * fi13)
                            _ = v31
                        }
                        v32 := GenS4{F0: uint(v8)}
                        _ = v32
                    }
                }
            }
            v7 = v7
            _ = len(v1)
            _ = fi13
        }
        v7 = false
        _ = fi12
    }
    _ = len(v4)
    var v33 int64 = 0
    var v34 float64 = 1.5
    if ((^(v8)) < v8) {
        _ = len(v5)
    }
    if (!v7) {
        for fi35 := 0; fi35 < 15; fi35++ {
            v36 := v6["foo bar"]
            var v37 int32 = 2147483647
            switch "a" {
            case "\n\t":
                for fi38 := 0; fi38 < 18; fi38++ {
                    switch 907 {
                    case 776:
                        v39 := GenS5{F0: "a", F1: 0, F2: ((v8 != v8) || true)}
                        _ = v39
                    }
                    if (!false) {
                        _ = len(v4)
                    } else {
                        v34 += v34
                        v40 := v3["\n\t"]
                        switch "hello" {
                        case "a":
                            v37 += 2147483647
                        case "foo bar":
                            var v41 string = "hello"
                            _ = v41
                        }
                        _ = v40
                    }
                    if true {
                        switch ((((("\n\t" + ("\n\t" + "")) + "") + "\n\t") + ("foo bar" + "a")) + "hello") {
                        case "":
                            v42 := GenS1{F0: (^(v33)), F1: 18446744073709551615}
                            _ = v42
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v6[("" + ("a" + "hello"))] = (v36 % ((v36) | 1))
                    } else {
                        v34 += 0.0
                        switch (fi38 - (fi38 | (fi38 << ((2) & 63)))) {
                        case 127:
                            v43 := GenS4{F0: 180}
                            _ = v43
                        case -1:
                            var v44 int32 = v37
                            _ = v44
                        }
                    }
                    _ = fi38
                }
            default:
                var v45 uint16 = 359
                _ = v45
            }
            _ = fi35
            _ = v36
            _ = v37
        }
        var v46 int16 = 1000
        _ = v46
    } else {
        for fi47 := 0; fi47 < 9; fi47++ {
            for fi48 := 0; fi48 < 16; fi48++ {
                v49 := GenS4{F0: 500}
                v50 := GenS4{F0: 65535}
                var v51 int8 = 71
                _ = fi48
                _ = v49
                _ = v50
                _ = v51
            }
            switch -235 {
            case 623:
                if (!true) {
                    v7 = false
                    fmt.Println(v7, fi47, v33)
                    switch 812 {
                    case -9223372036854775808:
                        _ = len(v2)
                    case -744:
                        v52 := GenS6{F0: -36, F1: uint32(v8), F2: v7}
                        _ = v52
                    }
                }
            case 32768:
                for fi53 := 0; fi53 < 1; fi53++ {
                    var v54 int32 = 2147483647
                    switch "hello" {
                    case "":
                        fmt.Println(v33)
                    default:
                        v54 += 2147483647
                    }
                    if (42 <= 596) {
                        if false {
                            v1 = append(v1, "foo bar")
                        }
                        switch "foo bar" {
                        case "foo bar":
                            _ = len(v2)
                        default:
                            switch ("hello" + ("foo bar" + "hello")) {
                            case "hello":
                                switch ("hello" + "\n\t") {
                                case "\n\t":
                                    v55 := GenS3{F0: v7, F1: 42}
                                    _ = v55
                                case "":
                                    var v56 int16 = -257
                                    _ = v56
                                case "foo bar":
                                    v57 := GenS4{F0: 1}
                                    _ = v57
                                }
                            case "\n\t":
                                if (v7 || (!true)) {
                                    fmt.Println(v8, v33)
                                } else {
                                    if v7 {
                                        v33 -= (-(v33))
                                        var v58 int8 = -8
                                        v8 = v8
                                        _ = v58
                                    }
                                }
                            default:
                                var v59 uint8 = (v8 % (((v8 & 148)) | 1))
                                _ = v59
                            }
                        }
                        v54 = 1
                    } else {
                        v4 = append(v4, "a")
                        fmt.Println(v8, v54, v54)
                        fmt.Println(fi47, v8)
                    }
                    _ = fi53
                    _ = v54
                }
            default:
                v60 := GenS7{F0: 42}
                _ = v60
            }
            v61 := GenS4{F0: 18446744073709551615}
            _ = fi47
            _ = v61
        }
        if ((0.0 + float64(((0.0 * 1.5) / (float32(0.0) + (3.14159 / 1.0))))) == v34) {
            for fi62 := 0; fi62 < 1; fi62++ {
                v63 := GenS8{F0: 42}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi62
                _ = v63
            }
        } else {
            var v64 int = 42
            _ = len(v6)
            v65 := GenS6{F0: 139, F1: 2, F2: (!v7)}
            _ = v64
            _ = v65
        }
    }
    for fi66 := 0; fi66 < 13; fi66++ {
        v67 := GenS8{F0: 0}
        if (v7 || false) {
            if true {
                var v68 float32 = (3.14159 / 1.0)
                v69 := GenS5{F0: "", F1: -128, F2: (280 >= 4294967295)}
                _ = v68
                _ = v69
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            for fi70 := 0; fi70 < 12; fi70++ {
                v71 := GenS9{F0: (("" + ("\n\t" + (((("foo bar" + ("\n\t" + "foo bar")) + ("a" + (("hello" + "a") + ""))) + ("" + "\n\t")) + (("" + "") + "\n\t")))) + ("foo bar" + ("hello" + (("hello" + (("a" + "hello") + (("hello" + "") + "\n\t"))) + "hello"))))}
                _ = fi70
                _ = v71
            }
            switch fi66 {
            case 884:
                fmt.Println(fi66, v7, fi66)
            case -18:
                fmt.Println(v7, v33)
            case 587:
                switch fi66 {
                case -9223372036854775808:
                    v34 = (v34 / 3.14159)
                case 247:
                    fmt.Println(v8, v34, fi66)
                default:
                    switch fi66 {
                    case 42:
                        v7 = (128 > 0)
                    default:
                        var v72 float32 = 1e10
                        _ = v72
                    }
                }
            default:
                v34 -= 1.5
            }
        }
        if ((2147483647 > 2147483648) || (2147483647 >= -2147483648)) {
            switch "\n\t" {
            case "":
                v7 = (!true)
            default:
                if (4294967295 < 18446744073709551615) {
                    v73 := GenS9{F0: (("hello" + ("foo bar" + "\n\t")) + "a")}
                    v34 += -2.25
                    var v74 uint64 = 1
                    _ = v73
                    _ = v74
                }
            }
            v75 := GenS10{F0: 1000}
            _ = v75
        } else {
            var v76 uint = 42
            _ = v76
        }
        _ = fi66
        _ = v67
    }
    if (!v7) {
        switch "" {
        case "":
            v77 := v3[("" + "hello")]
            _ = v77
        case "\n\t":
            var v78 bool = v7
            _ = v78
        default:
            v33 = v33
        }
        v79 := GenS11{F0: 42, F1: -1}
        _ = v79
    } else {
        var v80 bool = v7
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v34, v8)
        case "\n\t":
            var v81 int = 2147483647
            _ = v81
        case "hello":
            v82 := GenS7{F0: 18446744073709551615}
            _ = v82
        }
        var v83 int8 = -128
        _ = v80
        _ = v83
    }
    fmt.Println(v34, v34, v33)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
