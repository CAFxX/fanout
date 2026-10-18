package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int64; F2 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 int8; F2 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint; F1 int32; F2 int64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 299
    v3 := make(map[uint32]string)
    v4 := make(map[string]int)
    v5 := make(map[int]bool)
    v6 := []string{"hello", "", "", "foo bar"}
    v2 = (v2 << ((18446744073709551615) & 63))
    var v7 int16 = 42
    for fi8 := 0; fi8 < 13; fi8++ {
        switch "hello" {
        case "foo bar":
            var v9 float64 = (((1.5 * (1.5 / 1.5)) * 1e10) + (3.14159 / 1.5))
            _ = v9
        case "a":
            if len(v6) > 0 {
                v6[0] = "hello"
            }
        case "":
            var v10 int16 = (v7 * 1)
            _ = v10
        default:
            delete(v1, (fi8 | (^(fi8))))
        }
        var v11 int8 = 46
        v12 := GenS1{F0: -561, F1: 255, F2: 262}
        _ = fi8
        _ = v11
        _ = v12
    }
    var v13 int64 = 1
    var v14 uint8 = 46
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 uint32 = 0
        for fi17 := 0; fi17 < 4; fi17++ {
            var v18 float32 = 0.1
            _ = fi17
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    v19 := GenS1{F0: (v7 >> ((531) & 63)), F1: 2, F2: v2}
    var v20 float64 = (0.1 + ((((0.0 / 3.14159) * ((((0.0 - (3.14159 + 0.1)) - 0.0) - 1.5) + 0.1)) + float64(0.0)) - 0.1))
    delete(v1, 4294967295)
    var v21 int16 = v7
    switch 1 {
    case 268:
        v22 := GenS2{F0: 127}
        _ = v22
    case -9223372036854775808:
        var v23 string = ("foo bar" + ("hello" + "foo bar"))
        _ = v23
    case -536:
        if (!false) {
            fmt.Println(v13, v2, v20)
            v24 := GenS3{F0: -131, F1: v13, F2: 0}
            v13 = -9223372036854775808
            _ = v24
        }
    }
    var v25 int = 4294967295
    var v26 uint32 = v2
    for fi27 := 0; fi27 < 17; fi27++ {
        if true {
            switch "" {
            case "hello":
                v14 -= v14
            case "\n\t":
                var v28 uint16 = 32767
                _ = v28
            default:
                if (18446744073709551615 <= uint64(v2)) {
                    v21 = v21
                    var v29 uint32 = (v26 >> ((65535) & 63))
                    _ = v29
                } else {
                    v30 := GenS4{F0: (true && true)}
                    for fi31 := 0; fi31 < 10; fi31++ {
                        switch ("a" + (("" + "\n\t") + ("\n\t" + "foo bar"))) {
                        case "foo bar":
                            fmt.Println(v2)
                        case "":
                            fmt.Println(v25)
                        case "hello":
                            v32 := GenS5{F0: v13, F1: -1, F2: 0}
                            _ = v32
                        }
                        if true {
                            switch "\n\t" {
                            case "a":
                                var v33 float32 = ((3.14159 / (1.5 * 1.5)) - 1e10)
                                _ = v33
                            case "\n\t":
                                delete(v3, v26)
                            case "foo bar":
                                fmt.Println(v14, v26, fi27)
                            }
                            var v34 string = "\n\t"
                            var v35 int16 = v7
                            _ = v34
                            _ = v35
                        } else {
                            fmt.Println(v26)
                        }
                        v6 = append(v6, "foo bar")
                        _ = fi31
                    }
                    if true {
                        _ = len(v1)
                        if (v7 <= (v21 - v7)) {
                            v36 := GenS6{F0: "foo bar"}
                            _ = v36
                        }
                        if (-397 <= -578) {
                            v37 := GenS2{F0: (v14 >> ((2) & 63))}
                            _ = v37
                        } else {
                            switch v25 {
                            case 128:
                                if false {
                                    v38 := GenS2{F0: (v14 & (^(v14)))}
                                    var v39 uint32 = 0
                                    v7 = v21
                                    _ = v38
                                    _ = v39
                                } else {
                                    for fi40 := 0; fi40 < 13; fi40++ {
                                        var v41 uint8 = v14
                                        v21 -= -1
                                        _ = fi40
                                        _ = v41
                                    }
                                    if (("\n\t" >= "hello") && true) {
                                        switch ("\n\t" + ("hello" + ("" + ""))) {
                                        case "hello":
                                            if (!false) {
                                                for fi42 := 0; fi42 < 2; fi42++ {
                                                    v43 := GenS3{F0: -2147483648, F1: (v13 * v13), F2: (v14 << ((uint(v14)) & 63))}
                                                    _ = fi42
                                                    _ = v43
                                                }
                                            } else {
                                                v44 := GenS6{F0: "foo bar"}
                                                for fi45 := 0; fi45 < 13; fi45++ {
                                                    if true {
                                                        v46 := GenS5{F0: v13, F1: 1, F2: 777}
                                                        switch ("" + "") {
                                                        case "a":
                                                            v25 = fi45
                                                        default:
                                                            v47 := GenS7{F0: 18446744073709551615, F1: -653, F2: (v13 * (v13 + v13))}
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            _ = len(v4)
                                                            _ = v47
                                                        }
                                                        var v48 uint64 = 107
                                                        _ = v46
                                                        _ = v48
                                                    } else {
                                                        fmt.Println(v26, v14, v13)
                                                        var v49 uint8 = (v14 | v14)
                                                        _ = len(v3)
                                                        _ = v49
                                                    }
                                                    if false {
                                                        if len(v6) > 0 {
                                                            v6[0] = "a"
                                                        }
                                                    } else {
                                                        v20 = (v20 / (v20 / v20))
                                                        v6 = append(v6, "hello")
                                                        fmt.Println(v21, fi45)
                                                        fmt.Println(fi27, v7)
                                                    }
                                                    var v50 uint8 = 61
                                                    _ = fi45
                                                    _ = v50
                                                }
                                                _ = v44
                                            }
                                        case "a":
                                            var v51 int = -9223372036854775808
                                            _ = v51
                                        case "foo bar":
                                            var v52 uint = 49
                                            _ = v52
                                        default:
                                            if true {
                                                v2 = v2
                                                _ = len(v5)
                                            } else {
                                                switch (v25 ^ fi27) {
                                                case 2147483647:
                                                    var v53 int16 = (-(v21))
                                                    _ = v53
                                                default:
                                                    v54 := GenS8{F0: 1}
                                                    _ = v54
                                                }
                                                v55 := GenS9{F0: v13}
                                                v20 += (v20 + v20)
                                                _ = v55
                                            }
                                        }
                                        v7 = v21
                                    } else {
                                        v26 = v26
                                        v20 = (v20 * v20)
                                    }
                                    _ = len(v6)
                                }
                            case -577:
                                var v56 float64 = 1.5
                                _ = v56
                            default:
                                for fi57 := 0; fi57 < 4; fi57++ {
                                    if false {
                                        var v58 uint16 = 425
                                        if (v58 == 1) {
                                            if false {
                                                if false {
                                                    var v59 int = (v25 ^ fi27)
                                                    _ = v59
                                                } else {
                                                    if (1 <= 42) {
                                                        var v60 uint8 = v14
                                                        fmt.Println(v26, v14)
                                                        _ = v60
                                                    }
                                                    var v61 uint = 0
                                                    var v62 float32 = ((3.14159 - 1.5) + -2.25)
                                                    _ = v61
                                                    _ = v62
                                                }
                                                var v63 int8 = 0
                                                _ = v63
                                            }
                                            fmt.Println(v14)
                                        } else {
                                            fmt.Println(v25)
                                            var v64 int32 = 1
                                            _ = v64
                                        }
                                        var v65 float32 = float32((v20 * ((0.1 / 0.1) + v20)))
                                        _ = v58
                                        _ = v65
                                    }
                                    _ = fi57
                                }
                            }
                            delete(v1, (fi27 << ((2) & 63)))
                            switch (v25 / ((-336) | 1)) {
                            case 127:
                                var v66 uint64 = 42
                                _ = v66
                            case -286:
                                for fi67 := 0; fi67 < 10; fi67++ {
                                    v68 := GenS7{F0: 0, F1: -2147483648, F2: (v13 << ((440) & 63))}
                                    switch "" {
                                    case "hello":
                                        fmt.Println(v20, fi27)
                                    default:
                                        v13 += v13
                                    }
                                    switch (("foo bar" + "\n\t") + "\n\t") {
                                    case "a":
                                        var v69 int8 = -1
                                        v70 := GenS3{F0: -76, F1: v13, F2: 100}
                                        _ = v69
                                        _ = v70
                                    case "foo bar":
                                        if false {
                                            v13 = v13
                                        } else {
                                            switch "hello" {
                                            case "foo bar":
                                                var v71 string = (((("hello" + (("a" + (("hello" + "") + ("" + ("a" + (("a" + "") + ("\n\t" + "foo bar")))))) + "hello")) + "") + "foo bar") + "a")
                                                _ = v71
                                            case "hello":
                                                var v72 float32 = float32(1.5)
                                                _ = v72
                                            }
                                        }
                                    default:
                                        v25 += 2147483647
                                    }
                                    _ = fi67
                                    _ = v68
                                }
                            case 32768:
                                fmt.Println(fi27, v14)
                            default:
                                for fi73 := 0; fi73 < 5; fi73++ {
                                    if (!false) {
                                        v74 := GenS2{F0: v14}
                                        fmt.Println(v26, fi73)
                                        var v75 uint64 = 148
                                        _ = v74
                                        _ = v75
                                    } else {
                                        v14 = (^(v14))
                                        fmt.Println(v26, v21, v2)
                                    }
                                    _ = fi73
                                }
                            }
                        }
                    }
                    _ = v30
                }
            }
            v76 := GenS8{F0: 2}
            var v77 int16 = v21
            _ = v76
            _ = v77
        } else {
            v78 := GenS8{F0: uint64(v2)}
            v79 := GenS1{F0: 1, F1: 255, F2: (v26 % ((v2) | 1))}
            v80 := GenS3{F0: 1, F1: v13, F2: (v14 ^ 127)}
            _ = v78
            _ = v79
            _ = v80
        }
        _ = fi27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
