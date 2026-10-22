package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint64 }
type GenS2 struct { F0 int8; F1 uint8; F2 string }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 uint; F2 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 int32 = 2147483647
    var v2 int64 = 42
    v3 := []string{"a", "\n\t", (("foo bar" + "a") + "a"), "a"}
    var v4 uint16 = 494
    for fi5 := 0; fi5 < 6; fi5++ {
        if true {
            v2 -= v2
            v1 -= v1
        }
        _ = fi5
    }
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    var v6 uint = 18446744073709551615
    var v7 int16 = 41
    for fi8 := 0; fi8 < 6; fi8++ {
        if (v6 == (v6 % ((v6) | 1))) {
            if len(v3) > 0 {
                v3[0] = ""
            }
        }
        var v9 int32 = -616
        for fi10 := 0; fi10 < 20; fi10++ {
            switch "foo bar" {
            case "hello":
                fmt.Println(v1)
            default:
                switch 695 {
                case 4294967295:
                    var v11 int64 = v2
                    _ = v11
                case 32767:
                    v9 += -2147483648
                default:
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                }
                v2 -= (v2 - v2)
            }
            var v12 int64 = v2
            v7 = v7
            _ = fi10
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    v13 := GenS1{F0: (v2 >> ((552) & 63)), F1: v2, F2: 979}
    for fi14 := 0; fi14 < 11; fi14++ {
        for fi15 := 0; fi15 < 9; fi15++ {
            v16 := GenS1{F0: int64(v4), F1: 42, F2: 301}
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 6; fi17++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v18 uint64 = 1
            if false {
                v6 = v6
                _ = len(v3)
            }
            _ = fi17
            _ = v18
        }
        var v19 string = ""
        _ = fi14
        _ = v19
    }
    if false {
        if (("hello" + (("a" + "hello") + (("a" + "\n\t") + ""))) < "a") {
            fmt.Println(v1)
        }
    } else {
        v6 -= (^(v6))
        var v20 uint8 = 1
        if false {
            v7 += v7
        }
        _ = v20
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        if ((3.14159 * -2.25) <= float64(1.5)) {
            v1 += (v1 * v1)
            if (0.1 < (1e10 - ((1e10 / 1e10) + float64(-2.25)))) {
                if false {
                    var v22 int = (fi21 ^ fi21)
                    v3 = append(v3, ("hello" + "\n\t"))
                    _ = v22
                }
            }
            for fi23 := 0; fi23 < 12; fi23++ {
                fmt.Println(v4, v2)
                v3 = append(v3, "hello")
                switch 127 {
                case 32767:
                    v2 += v2
                default:
                    v3 = append(v3, (("hello" + "foo bar") + ("a" + "")))
                }
                _ = fi23
            }
        } else {
            fmt.Println(v4, v6, v2)
            for fi24 := 0; fi24 < 3; fi24++ {
                var v25 int = 42
                if ((false && ((v4 / ((v4) | 1)) != (v4 & (v4 * 380)))) || ((fi21 + v25) > fi24)) {
                    var v26 int64 = 42
                    var v27 uint32 = 2
                    _ = v26
                    _ = v27
                } else {
                    var v28 int = fi24
                    v28 -= v25
                    v28 -= 9223372036854775807
                    _ = v28
                }
                v2 = v2
                _ = fi24
                _ = v25
            }
            for fi29 := 0; fi29 < 1; fi29++ {
                v6 += v6
                if true {
                    v6 += v6
                    if (!(true && false)) {
                        if (!false) {
                            var v30 uint32 = 908
                            var v31 int8 = -5
                            var v32 bool = true
                            _ = v30
                            _ = v31
                            _ = v32
                        } else {
                            if (uint16(v6) <= (v4 + v4)) {
                                fmt.Println(fi21)
                                if false {
                                    if true {
                                        fmt.Println(v2)
                                    }
                                    if (!false) {
                                        if false {
                                            if (true || false) {
                                                var v33 int64 = -1
                                                var v34 int = 538
                                                v6 = v6
                                                _ = v33
                                                _ = v34
                                            }
                                            var v35 uint8 = 1
                                            _ = v35
                                        }
                                        var v36 bool = true
                                        _ = v36
                                    }
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v37 int64 = 1
                                    _ = v37
                                }
                                switch 127 {
                                case 1:
                                    switch "hello" {
                                    case "\n\t":
                                        v2 = v2
                                    case "foo bar":
                                        v38 := GenS2{F0: 83, F1: 128, F2: "\n\t"}
                                        _ = v38
                                    }
                                case 197:
                                    switch "" {
                                    case "hello":
                                        v2 += (v2 * (v2 * v2))
                                    case "a":
                                        if (!(((!true) || (4294967295 != 0)) || (214 != 65535))) {
                                            if (-2.25 > float64(float32(((((3.14159 / float64(0.1)) + 0.1) / (((0.1 - (-2.25 + 1e10)) - 1e10) / ((float64(float32(1.5)) * 1e10) + (-2.25 / 1e10)))) - 3.14159)))) {
                                                fmt.Println(v7, v4)
                                                fmt.Println(v1, v6, v6)
                                                var v39 uint64 = 432
                                                _ = v39
                                            }
                                        }
                                    case "foo bar":
                                        if (1 > -1) {
                                            fmt.Println(v1, fi29, v2)
                                            var v40 float64 = 1e10
                                            var v41 int32 = v1
                                            _ = v40
                                            _ = v41
                                        }
                                    }
                                case 760:
                                    switch 4294967295 {
                                    case -934:
                                        var v42 uint = 2
                                        var v43 uint32 = 2
                                        var v44 int32 = (v1 * v1)
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    case -9223372036854775808:
                                        fmt.Println(v4)
                                    default:
                                        switch fi29 {
                                        case -467:
                                            _ = len(v3)
                                        case 2147483647:
                                            if true {
                                                switch fi21 {
                                                case 32768:
                                                    if len(v3) > 0 {
                                                        v3[0] = "\n\t"
                                                    }
                                                case 42:
                                                    var v45 int64 = -954
                                                    if true {
                                                        v7 -= (v7 % ((v7) | 1))
                                                        switch (fi21 & (fi29 * 326)) {
                                                        case 42:
                                                            fmt.Println(v6)
                                                        case -9223372036854775808:
                                                            if (v2 < 42) {
                                                                v3 = append(v3, "hello")
                                                                if (441 != (v6 % ((uint(v4)) | 1))) {
                                                                    if (false && ((!(((32768 == int(v4)) || false) && true)) && false)) {
                                                                        var v46 uint16 = v4
                                                                        var v47 float32 = (0.0 / 0.1)
                                                                        _ = v46
                                                                        _ = v47
                                                                    } else {
                                                                        v48 := GenS3{F0: -87, F1: uint8(v4)}
                                                                        fmt.Println(v6, fi29)
                                                                        fmt.Println(v6)
                                                                        _ = v48
                                                                    }
                                                                }
                                                            } else {
                                                                fmt.Println(fi29)
                                                                v2 += v2
                                                                v4 = v4
                                                                v49 := GenS3{F0: 127, F1: 121}
                                                                _ = v49
                                                            }
                                                        case 32768:
                                                            fmt.Println(v45, v7, v4)
                                                        default:
                                                            _ = len(v3)
                                                        }
                                                    } else {
                                                        if false {
                                                            if len(v3) > 0 {
                                                                v3[0] = (("foo bar" + "") + "\n\t")
                                                            }
                                                            if len(v3) > 0 {
                                                                v3[0] = ""
                                                            }
                                                        }
                                                    }
                                                    if (!(("hello" + "hello") == (("foo bar" + ("foo bar" + "a")) + ("hello" + "hello")))) {
                                                        switch (fi29 / (((-(fi21))) | 1)) {
                                                        case 436:
                                                            if false {
                                                                v50 := GenS4{F0: v7}
                                                                v6 += (v6 + (v6 | (^(v6))))
                                                                var v51 uint16 = v4
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                        case -1:
                                                            fmt.Println(v2, v6)
                                                        default:
                                                            v52 := GenS5{F0: 584, F1: (v6 ^ 2), F2: 667}
                                                            _ = v52
                                                        }
                                                    }
                                                    _ = v45
                                                case 1:
                                                    if (!false) {
                                                        var v53 int = (fi29 + 2147483647)
                                                        v54 := GenS6{F0: 154}
                                                        switch ("" + "") {
                                                        case "\n\t":
                                                            v3 = append(v3, "a")
                                                        case "hello":
                                                            v55 := GenS7{F0: (v4 + (v4 & (v4 << ((v6) & 63))))}
                                                            _ = len(v3)
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            var v56 int32 = int32(fi29)
                                                            v57 := GenS5{F0: (fi21 / (((v53 >> (((v6 / ((2) | 1))) & 63))) | 1)), F1: 65535, F2: uint64(v6)}
                                                            _ = v55
                                                            _ = v56
                                                            _ = v57
                                                        default:
                                                            v2 = v2
                                                            v1 = 42
                                                        }
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                default:
                                                    if (int64(fi29) != -967) {
                                                        fmt.Println(fi21, v4, v7)
                                                    }
                                                }
                                                var v58 string = "\n\t"
                                                _ = v58
                                            }
                                        }
                                    }
                                }
                            }
                        }
                        var v59 uint64 = 732
                        if (("" + "hello") <= (("a" + "\n\t") + "hello")) {
                            fmt.Println(v6, v6, v59)
                            var v60 float32 = (0.1 / 1.0)
                            _ = v60
                        } else {
                            v6 = (v6 - (v6 % ((v6) | 1)))
                            v4 -= v4
                            var v61 uint16 = (^(v4))
                            _ = v61
                        }
                        _ = v59
                    }
                }
                v62 := GenS5{F0: (fi29 - fi21), F1: (v6 - v6), F2: 1}
                v1 -= v1
                switch fi21 {
                case 4294967295:
                    var v63 uint = (v6 >> ((v6) & 63))
                    _ = v63
                default:
                    fmt.Println(fi21)
                }
                _ = fi29
                _ = v62
            }
        }
        v6 += v6
        _ = fi21
    }
    var v64 uint16 = v4
    fmt.Println(v2, v2, v6)
    for fi65 := 0; fi65 < 17; fi65++ {
        var v66 int16 = -32768
        var v67 float64 = 0.1
        _ = fi65
        _ = v66
        _ = v67
    }
    var v68 uint64 = 2
    v69 := GenS1{F0: int64(v1), F1: (v2 ^ v2), F2: (v68 | v68)}
    v64 -= v4
    v70 := GenS8{F0: (v4 % (((v4 << ((v6) & 63))) | 1)), F1: v68}
    var v71 uint = v6
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    fmt.Println("v68:", v68)
    fmt.Println("v7:", v7)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v64
    _ = v68
    _ = v69
    _ = v7
    _ = v70
    _ = v71
}
