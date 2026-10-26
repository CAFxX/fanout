package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 uint32; F2 uint }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 uint8; F1 int16; F2 int64 }
type GenS6 struct { F0 int32; F1 string; F2 uint16 }
type GenS7 struct { F0 int; F1 int64; F2 int64 }

func main() {
    var v1 float64 = 0.0
    v2 := make(map[uint32]uint32)
    v3 := []uint32{}
    var v4 float32 = 0.0
    var v5 bool = ("" <= "\n\t")
    v6 := []string{"", ("\n\t" + ("foo bar" + "\n\t")), ("a" + "hello"), "hello"}
    switch (("hello" + "a") + "foo bar") {
    case "a":
        v4 -= (((v4 - (v4 + v4)) / v4) - (v4 - 3.14159))
    case "":
        for fi7 := 0; fi7 < 8; fi7++ {
            v4 = v4
            if true {
                v1 += 1e10
                v8 := GenS1{F0: 167, F1: 1, F2: -486}
                v6 = append(v6, "foo bar")
                _ = v8
            } else {
                v9 := GenS2{F0: v5}
                _ = v9
            }
            _ = fi7
        }
    }
    var v10 int16 = 42
    v1 += 0.0
    fmt.Println(v1)
    v1 -= float64(v4)
    v11 := GenS2{F0: v5}
    for fi12 := 0; fi12 < 4; fi12++ {
        v1 += v1
        _ = fi12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        switch "hello" {
        case "\n\t":
            v4 = 3.14159
        case "":
            var v14 int64 = 9223372036854775807
            _ = v14
        default:
            var v15 int16 = v10
            _ = v15
        }
        for fi16 := 0; fi16 < 8; fi16++ {
            v5 = v5
            _ = fi16
        }
        v1 -= v1
        _ = fi13
    }
    switch -482 {
    case 252:
        switch "foo bar" {
        case "\n\t":
            v4 -= v4
        case "hello":
            switch 4294967295 {
            case -2147483648:
                if ((0 != 1) && ("" < "hello")) {
                    var v17 uint32 = 2
                    _ = v17
                } else {
                    switch (("hello" + "hello") + "hello") {
                    case "foo bar":
                        v18 := GenS2{F0: (42 <= -432)}
                        _ = v18
                    case "":
                        v4 -= (v4 - v4)
                    default:
                        v10 -= v10
                    }
                    for fi19 := 0; fi19 < 19; fi19++ {
                        switch "a" {
                        case "a":
                            for fi20 := 0; fi20 < 12; fi20++ {
                                var v21 uint = 63
                                v21 += v21
                                _ = fi20
                                _ = v21
                            }
                        default:
                            if false {
                                v22 := GenS3{F0: ("foo bar" + ("" + "foo bar")), F1: 4294967295, F2: 625}
                                for fi23 := 0; fi23 < 13; fi23++ {
                                    v1 = (3.14159 - v1)
                                    var v24 float64 = v1
                                    var v25 float64 = v1
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                }
                                v1 -= (v1 - (v1 * (0.0 + (float64(v4) * 3.14159))))
                                _ = v22
                            }
                        }
                        var v26 uint16 = 0
                        v4 = 3.14159
                        _ = fi19
                        _ = v26
                    }
                }
            }
        default:
            v10 = 1000
        }
    case -1:
        v5 = v5
    default:
        switch ("\n\t" + "") {
        case "\n\t":
            var v27 float32 = ((v4 / float32((v1 * (v1 / 3.14159)))) * float32((v1 + 1e10)))
            _ = v27
        case "hello":
            var v28 uint8 = 1
            _ = v28
        case "":
            for fi29 := 0; fi29 < 7; fi29++ {
                var v30 int16 = (v10 | (v10 >> ((65535) & 63)))
                _ = fi29
                _ = v30
            }
        }
    }
    v31 := GenS1{F0: 127, F1: -929, F2: 603}
    if v5 {
        v10 -= v10
    }
    v10 -= (^(v10))
    for fi32 := 0; fi32 < 10; fi32++ {
        switch fi32 {
        case -2147483648:
            v33 := GenS4{F0: 401, F1: v5}
            _ = v33
        default:
            switch fi32 {
            case -9223372036854775808:
                v10 -= v10
            case 305:
                if false {
                    for fi34 := 0; fi34 < 5; fi34++ {
                        v5 = v5
                        var v35 int32 = 1
                        _ = fi34
                        _ = v35
                    }
                    if ((0 == -123) || true) {
                        var v36 uint = 4294967295
                        var v37 int16 = -32768
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = v36
                        _ = v37
                    }
                    _ = len(v2)
                }
            case 105:
                for fi38 := 0; fi38 < 18; fi38++ {
                    if (475 <= 1) {
                        if v5 {
                            if (65535 > 4294967295) {
                                v39 := GenS2{F0: ((v10 << ((255) & 63)) == (^(v10)))}
                                _ = v39
                            } else {
                                v40 := GenS3{F0: "\n\t", F1: 1, F2: 774}
                                if v5 {
                                    if ((v4 < ((3.14159 / v4) + (((0.0 + (-2.25 * v4)) + (v4 / (v4 * float32((0.0 / ((0.0 + v1) * 1.5)))))) + (((-2.25 / 3.14159) * (float32(v1) * v4)) + v4)))) || (!false)) {
                                        delete(v2, 43)
                                        v3 = append(v3, 940)
                                    } else {
                                        var v41 int64 = -9223372036854775808
                                        if (((^(fi38)) == fi38) || (4294967295 > 0)) {
                                            if (!("\n\t" == ("foo bar" + "foo bar"))) {
                                                var v42 int64 = 42
                                                var v43 float32 = (0.1 + (0.0 * (1e10 / ((v4 * v4) * v4))))
                                                v42 += v42
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v42
                                                _ = v43
                                            }
                                            var v44 uint8 = 24
                                            var v45 bool = ((v41 % ((-1) | 1)) == -9223372036854775808)
                                            var v46 uint8 = v44
                                            fmt.Println(v5, v41, v10)
                                            var v47 float64 = -2.25
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                            _ = v47
                                        }
                                        v48 := GenS1{F0: 2, F1: 1, F2: 42}
                                        _ = v41
                                        _ = v48
                                    }
                                    if (42 > 115) {
                                        fmt.Println(v4, v10)
                                        switch "a" {
                                        case "hello":
                                            v5 = v5
                                        case "\n\t":
                                            v1 = (((v1 * v1) * v1) / v1)
                                        case "":
                                            v10 -= int16(fi32)
                                        }
                                        v49 := GenS3{F0: "\n\t", F1: 65535, F2: 0}
                                        _ = v49
                                    } else {
                                        var v50 int8 = -106
                                        var v51 int64 = -9223372036854775808
                                        v10 = (v10 * (v10 | (v10 >> ((2) & 63))))
                                        switch fi38 {
                                        case -1:
                                            v5 = v5
                                        case -182:
                                            var v52 float64 = (-2.25 + v1)
                                            _ = v52
                                        case 32768:
                                            var v53 uint = 42
                                            _ = v53
                                        }
                                        _ = v50
                                        _ = v51
                                    }
                                    if false {
                                        v54 := GenS2{F0: v5}
                                        delete(v2, 34)
                                        v10 = (v10 >> ((2) & 63))
                                        _ = v54
                                    } else {
                                        switch "a" {
                                        case "":
                                            v55 := GenS1{F0: 42, F1: 1, F2: 299}
                                            _ = v55
                                        case "hello":
                                            v1 += (float64(3.14159) * v1)
                                        case "a":
                                            var v56 bool = v5
                                            switch fi32 {
                                            case 4294967295:
                                                fmt.Println(v10, v10)
                                            default:
                                                var v57 uint8 = 224
                                                _ = v57
                                            }
                                            var v58 uint32 = 854
                                            if (110 != 42) {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    fmt.Println(v58)
                                                    _ = len(v2)
                                                    if true {
                                                        v59 := GenS5{F0: 42, F1: (v10 >> ((65535) & 63)), F2: 1}
                                                        var v60 int16 = (v10 + v10)
                                                        _ = v59
                                                        _ = v60
                                                    } else {
                                                        fmt.Println(v10, v4, v58)
                                                        v5 = (!((!(-643 > 42)) && (1 != 18446744073709551615)))
                                                        var v61 int64 = -13
                                                        _ = v61
                                                    }
                                                case "hello":
                                                    v62 := GenS6{F0: -2147483648, F1: "\n\t", F2: 32767}
                                                    _ = v62
                                                case "a":
                                                    var v63 uint16 = 42
                                                    _ = v63
                                                default:
                                                    var v64 int16 = (v10 >> ((225) & 63))
                                                    _ = v64
                                                }
                                                switch "a" {
                                                case "":
                                                    if v56 {
                                                        if (v10 < (v10 & (v10 / (((^(v10))) | 1)))) {
                                                            var v65 string = "hello"
                                                            switch (((("hello" + v65) + (v65 + v65)) + "") + v65) {
                                                            case "a":
                                                                v4 = float32(v1)
                                                            case "hello":
                                                                v4 = v4
                                                            }
                                                            switch (fi32 % (((fi38 | 70)) | 1)) {
                                                            case 859:
                                                                v65 = "\n\t"
                                                            }
                                                            _ = v65
                                                        }
                                                    } else {
                                                        var v66 uint = 2
                                                        var v67 float32 = 0.0
                                                        _ = v66
                                                        _ = v67
                                                    }
                                                default:
                                                    var v68 float32 = 0.1
                                                    _ = v68
                                                }
                                                v56 = v5
                                            }
                                            _ = v56
                                            _ = v58
                                        default:
                                            var v69 uint = 255
                                            _ = v69
                                        }
                                        switch fi32 {
                                        case -9223372036854775808:
                                            switch -2147483648 {
                                            case 128:
                                                var v70 uint8 = 82
                                                _ = v70
                                            case 32767:
                                                v71 := GenS7{F0: (fi38 | fi32), F1: 42, F2: 0}
                                                _ = v71
                                            case -1:
                                                fmt.Println(v10)
                                            }
                                        }
                                    }
                                } else {
                                    v1 += v1
                                    if (!false) {
                                        fmt.Println(v5, v10)
                                    } else {
                                        v1 = v1
                                        v1 = (((v1 / -2.25) + (float64(float32(float64(0.1))) + float64(v4))) / 1e10)
                                        if (42 < -1) {
                                            v72 := GenS1{F0: 127, F1: -1, F2: 308}
                                            _ = v72
                                        } else {
                                            fmt.Println(fi32, v1, fi32)
                                            if v5 {
                                                v10 += (v10 / ((-181) | 1))
                                                fmt.Println(v10, v4, fi32)
                                            }
                                        }
                                    }
                                    if (v5 || v5) {
                                        var v73 bool = (false && v5)
                                        var v74 bool = v73
                                        _ = v73
                                        _ = v74
                                    }
                                }
                                _ = v40
                            }
                            fmt.Println(v10)
                            v5 = v5
                        } else {
                            _ = len(v6)
                            _ = len(v2)
                            if (2 == 127) {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                _ = len(v6)
                                v10 -= v10
                            }
                        }
                    }
                    var v75 int = fi38
                    _ = fi38
                    _ = v75
                }
            default:
                var v76 int64 = 1
                _ = v76
            }
        }
        v77 := GenS7{F0: 4294967295, F1: -1, F2: 9223372036854775807}
        _ = fi32
        _ = v77
    }
    var v78 int8 = 80
    fmt.Println(v4, v78, v1)
    fmt.Println(v1, v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v5 = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v78
}
