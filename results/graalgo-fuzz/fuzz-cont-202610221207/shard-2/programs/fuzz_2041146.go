package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS2 struct { F0 int32; F1 string; F2 uint64 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int; F1 uint64; F2 bool }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int16; F1 int32; F2 uint }
type GenS11 struct { F0 uint16; F1 int32 }
type GenS12 struct { F0 uint; F1 int }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int32; F1 uint64; F2 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 515
    var v3 float32 = (1.5 + ((float32((0.1 * -2.25)) - 3.14159) / 1.0))
    v4 := []string{("\n\t" + "a"), "foo bar", "foo bar"}
    var v5 uint32 = 65535
    if ((0.1 * 0.0) == 0.0) {
        _ = len(v4)
        for fi6 := 0; fi6 < 5; fi6++ {
            v5 += v5
            switch "hello" {
            case "\n\t":
                for fi7 := 0; fi7 < 14; fi7++ {
                    if (float64(v3) != (((0.1 / -2.25) - 3.14159) * 0.0)) {
                        v8 := GenS1{F0: 26, F1: v5, F2: 13}
                        delete(v1, "foo bar")
                        fmt.Println(v5, fi7)
                        _ = v8
                    } else {
                        v3 -= 1.5
                        var v9 float32 = (v3 - 0.1)
                        _ = v9
                    }
                    _ = fi7
                }
            case "a":
                fmt.Println(v5)
            case "":
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v5, fi6)
                case "a":
                    v2 = 2
                case "":
                    if (false && false) {
                        for fi10 := 0; fi10 < 1; fi10++ {
                            v3 += (v3 + v3)
                            switch (fi10 / (((fi10 | -2147483648)) | 1)) {
                            case 9223372036854775807:
                                v11 := GenS2{F0: int32(v5), F1: "\n\t", F2: 565}
                                _ = v11
                            default:
                                var v12 uint8 = 128
                                _ = v12
                            }
                            fmt.Println(v5, fi10, v5)
                            _ = fi10
                        }
                        _ = len(v4)
                        if false {
                            switch "foo bar" {
                            case "":
                                v13 := GenS3{F0: -9223372036854775808, F1: -790}
                                _ = v13
                            case "a":
                                for fi14 := 0; fi14 < 5; fi14++ {
                                    v5 += 2
                                    v15 := GenS2{F0: -2147483648, F1: ("" + "hello"), F2: 504}
                                    v16 := GenS2{F0: -1, F1: (("foo bar" + (("\n\t" + "hello") + "hello")) + ("foo bar" + "hello")), F2: 469}
                                    _ = fi14
                                    _ = v15
                                    _ = v16
                                }
                            default:
                                v5 += 574
                            }
                            v17 := GenS3{F0: 1, F1: 42}
                            _ = v17
                        }
                    } else {
                        switch "foo bar" {
                        case "":
                            v2 -= (v2 / (((v2 | (v2 * v2))) | 1))
                        case "a":
                            if false {
                                var v18 uint = 970
                                if false {
                                    v3 -= v3
                                    var v19 uint64 = 2
                                    if len(v4) > 0 {
                                        v4[0] = "\n\t"
                                    }
                                    _ = v19
                                }
                                v20 := GenS1{F0: 42, F1: v5, F2: 92}
                                _ = v18
                                _ = v20
                            }
                        case "hello":
                            var v21 int32 = -543
                            _ = v21
                        }
                        if len(v4) > 0 {
                            v4[0] = (("a" + "") + ("foo bar" + (("" + "a") + ("\n\t" + ""))))
                        }
                        var v22 bool = true
                        _ = v22
                    }
                }
            }
            v23 := GenS1{F0: 42, F1: (^(v5)), F2: 37}
            _ = fi6
            _ = v23
        }
        if true {
            var v24 int8 = int8(v2)
            fmt.Println(v2, v3)
            var v25 uint = 65535
            _ = v24
            _ = v25
        }
    } else {
        v2 -= v2
    }
    if (2 > 42) {
        v26 := GenS4{F0: 0}
        var v27 int16 = 32767
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 5; fi28++ {
        v2 += v2
        if false {
            if (343 == 2147483647) {
                var v29 int64 = 746
                _ = v29
            }
            v30 := GenS3{F0: -9223372036854775808, F1: 0}
            _ = v30
        } else {
            switch 317 {
            case 9223372036854775807:
                if (uint64(v5) != (^(v2))) {
                    var v31 uint32 = v5
                    v31 = 0
                    _ = v31
                }
            default:
                v5 -= v5
            }
            for fi32 := 0; fi32 < 11; fi32++ {
                v1[("\n\t" + "foo bar")] = (653 != int64(fi32))
                _ = fi32
            }
            _ = len(v1)
        }
        _ = fi28
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        v5 += v5
        _ = fi33
    }
    v5 = (v5 - 4294967295)
    _ = len(v1)
    if false {
        if (0 <= 213) {
            if false {
                var v34 int = 4294967295
                _ = v34
            }
            v3 = ((0.0 / v3) - (((v3 * v3) + 3.14159) * v3))
            for fi35 := 0; fi35 < 20; fi35++ {
                var v36 int8 = 0
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v37 := GenS4{F0: 774}
                _ = fi35
                _ = v36
                _ = v37
            }
        } else {
            var v38 string = ""
            _ = v38
        }
        if true {
            _ = len(v4)
        } else {
            v39 := GenS2{F0: 0, F1: "foo bar", F2: 642}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if true {
                for fi40 := 0; fi40 < 15; fi40++ {
                    var v41 uint64 = v2
                    if true {
                        _ = len(v4)
                        if false {
                            fmt.Println(v5, v41)
                            if (761 < -36) {
                                v42 := GenS4{F0: 0}
                                _ = v42
                            }
                            var v43 int = (fi40 >> ((4294967295) & 63))
                            _ = v43
                        } else {
                            var v44 uint32 = 741
                            var v45 int = 128
                            for fi46 := 0; fi46 < 13; fi46++ {
                                switch fi40 {
                                case 230:
                                    var v47 int64 = 1
                                    _ = v47
                                case 9223372036854775807:
                                    v4 = append(v4, "hello")
                                }
                                var v48 string = (("a" + "foo bar") + "foo bar")
                                var v49 float32 = float32(-2.25)
                                _ = fi46
                                _ = v48
                                _ = v49
                            }
                            _ = v44
                            _ = v45
                        }
                    } else {
                        _ = len(v4)
                        var v50 float64 = 3.14159
                        var v51 int64 = 0
                        _ = v50
                        _ = v51
                    }
                    v3 = v3
                    _ = fi40
                    _ = v41
                }
            } else {
                _ = len(v1)
            }
            _ = v39
        }
        fmt.Println(v2)
    } else {
        v52 := GenS5{F0: 0, F1: v2}
        if true {
            v5 = (v5 + v5)
            fmt.Println(v5, v2)
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v52
    }
    v2 += (v2 - uint64(v5))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 -= v3
    switch "hello" {
    case "hello":
        switch 42 {
        case 4294967295:
            v53 := GenS6{F0: "foo bar"}
            _ = v53
        case 582:
            for fi54 := 0; fi54 < 3; fi54++ {
                for fi55 := 0; fi55 < 4; fi55++ {
                    fmt.Println(fi54)
                    v3 = v3
                    if ((fi55 - 9223372036854775807) < 404) {
                        switch "" {
                        case "hello":
                            if false {
                                v56 := GenS7{F0: 1}
                                fmt.Println(fi54, v3)
                                if true {
                                    switch 1 {
                                    case -1:
                                        var v57 string = "\n\t"
                                        _ = v57
                                    case 624:
                                        var v58 int = fi54
                                        _ = v58
                                    default:
                                        v5 = (v5 & (^(v5)))
                                    }
                                }
                                v5 += (v5 % ((2147483648) | 1))
                                v3 += v3
                                switch 4294967295 {
                                case 0:
                                    var v59 uint = 138
                                    _ = v59
                                default:
                                    var v60 int64 = 0
                                    _ = v60
                                }
                                v61 := GenS3{F0: -939, F1: int32(v2)}
                                _ = v56
                                _ = v61
                            } else {
                                fmt.Println(v2, v3, fi55)
                            }
                            v3 = -2.25
                        case "foo bar":
                            var v62 int32 = -753
                            _ = v62
                        }
                        if true {
                            v63 := GenS7{F0: 1}
                            fmt.Println(fi54, v5)
                            var v64 float32 = 1e10
                            var v65 uint = 523
                            _ = len(v4)
                            _ = v63
                            _ = v64
                            _ = v65
                        }
                    } else {
                        var v66 int32 = 1
                        var v67 uint16 = 462
                        var v68 int = fi55
                        _ = v66
                        _ = v67
                        _ = v68
                    }
                    _ = fi55
                }
                _ = fi54
            }
        case -9223372036854775808:
            for fi69 := 0; fi69 < 13; fi69++ {
                v70 := GenS1{F0: 110, F1: v5, F2: 113}
                for fi71 := 0; fi71 < 11; fi71++ {
                    if false {
                        if true {
                            _ = len(v1)
                        } else {
                            v72 := v1[("foo bar" + ("hello" + "\n\t"))]
                            v3 = (v3 + (((v3 - v3) - v3) - (v3 * (((-2.25 - ((v3 - v3) - v3)) - v3) - v3))))
                            v73 := v1["foo bar"]
                            _ = v72
                            _ = v73
                        }
                        v74 := GenS8{F0: -1, F1: v2, F2: false}
                        var v75 uint16 = 255
                        _ = v74
                        _ = v75
                    } else {
                        fmt.Println(v2)
                    }
                    _ = fi71
                }
                var v76 uint32 = (v5 - v5)
                _ = fi69
                _ = v70
                _ = v76
            }
        default:
            if (-9223372036854775808 < 0) {
                for fi77 := 0; fi77 < 1; fi77++ {
                    for fi78 := 0; fi78 < 2; fi78++ {
                        v2 += (v2 ^ (v2 / ((1) | 1)))
                        if (!((^(v5)) >= 647)) {
                            v3 = float32(float64(v3))
                            v5 = v5
                        }
                        _ = fi78
                    }
                    v3 = (1e10 / (v3 / v3))
                    _ = fi77
                }
                switch 4294967295 {
                case 128:
                    v1[("\n\t" + "foo bar")] = (1 >= 2)
                case 621:
                    v5 -= v5
                default:
                    v5 = 0
                }
            }
        }
    case "":
        v3 += ((((v3 - ((v3 - (v3 + (v3 / ((v3 * (v3 + v3)) * 0.1)))) + ((((float32(-2.25) + float32(1.5)) + (-2.25 / (float32(1.5) * (v3 + v3)))) - (float32(3.14159) - v3)) + (-2.25 * 0.0)))) + float32(0.0)) / 1e10) - v3)
    case "a":
        if true {
            v79 := GenS9{F0: 42}
            _ = len(v4)
            _ = v79
        } else {
            _ = len(v4)
            v80 := GenS1{F0: 128, F1: v5, F2: 100}
            v3 -= v3
            _ = v80
        }
    }
    v81 := GenS6{F0: ("" + ("a" + "\n\t"))}
    v82 := GenS1{F0: 171, F1: 878, F2: -128}
    var v83 int = -1
    for fi84 := 0; fi84 < 17; fi84++ {
        fmt.Println(v5, v2)
        for fi85 := 0; fi85 < 12; fi85++ {
            v86 := GenS2{F0: 48, F1: ((((("foo bar" + ("a" + "foo bar")) + "foo bar") + "a") + ((("a" + "a") + "foo bar") + "a")) + ("\n\t" + ("hello" + ("foo bar" + "hello")))), F2: v2}
            v87 := GenS10{F0: 42, F1: -2147483648, F2: 796}
            if true {
                v88 := GenS5{F0: -659, F1: v2}
                v83 -= (v83 << ((18446744073709551615) & 63))
                _ = len(v4)
                _ = v88
            } else {
                if false {
                    v89 := GenS11{F0: 215, F1: 627}
                    switch fi85 {
                    case 42:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case -1:
                        if false {
                            if (!(729 >= v2)) {
                                if false {
                                    if true {
                                        _ = len(v4)
                                    }
                                    var v90 uint = 954
                                    _ = v90
                                }
                            }
                        }
                    }
                    v91 := GenS9{F0: 1}
                    _ = v89
                    _ = v91
                } else {
                    var v92 float64 = ((0.1 * 1.5) / 3.14159)
                    _ = v92
                }
                if (-252 != -2147483648) {
                    v93 := GenS12{F0: 522, F1: 1}
                    if false {
                        _ = len(v4)
                        fmt.Println(v2, v2)
                        v94 := GenS13{F0: 9223372036854775807}
                        _ = v94
                    }
                    if len(v4) > 0 {
                        v4[0] = ("hello" + "a")
                    }
                    _ = v93
                } else {
                    var v95 float64 = ((1.5 * 1e10) / 1.0)
                    _ = v95
                }
                var v96 float32 = 1e10
                _ = v96
            }
            _ = fi85
            _ = v86
            _ = v87
        }
        v97 := GenS14{F0: 42, F1: 311, F2: ((42 > 881) && true)}
        _ = fi84
        _ = v97
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v81
    _ = v82
    _ = v83
}
