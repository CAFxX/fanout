package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 bool }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 bool; F1 uint }
type GenS8 struct { F0 uint8; F1 int64; F2 int }
type GenS9 struct { F0 bool; F1 uint16 }
type GenS10 struct { F0 uint; F1 string; F2 uint8 }
type GenS11 struct { F0 int32 }

func main() {
    v1 := []string{"\n\t", "hello", "", ("\n\t" + "a")}
    var v2 int8 = 127
    var v3 float32 = -2.25
    var v4 float64 = 3.14159
    for fi5 := 0; fi5 < 20; fi5++ {
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := GenS1{F0: (v2 / ((-1) | 1)), F1: 65535, F2: false}
            v3 = float32(v4)
            switch "" {
            case "foo bar":
                if len(v1) > 0 {
                    v1[0] = ""
                }
            case "hello":
                var v8 int16 = 1
                _ = v8
            }
            _ = fi6
            _ = v7
        }
        var v9 uint32 = 0
        _ = fi5
        _ = v9
    }
    switch (("hello" + "foo bar") + (("foo bar" + (((("a" + ("a" + "foo bar")) + "a") + ("foo bar" + "\n\t")) + ("\n\t" + "foo bar"))) + "a")) {
    case "a":
        var v10 int32 = 2147483647
        _ = v10
    case "foo bar":
        fmt.Println(v2)
    case "\n\t":
        switch "a" {
        case "foo bar":
            switch -1 {
            case 32768:
                fmt.Println(v4, v3, v4)
            case 402:
                v11 := GenS2{F0: 9223372036854775807, F1: -776}
                _ = v11
            default:
                for fi12 := 0; fi12 < 6; fi12++ {
                    v13 := GenS1{F0: 100, F1: 18446744073709551615, F2: ((v4 * v4) != v4)}
                    _ = fi12
                    _ = v13
                }
            }
        default:
            if true {
                switch 1 {
                case -711:
                    for fi14 := 0; fi14 < 3; fi14++ {
                        v15 := GenS3{F0: false}
                        var v16 uint16 = 255
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                case 128:
                    switch "foo bar" {
                    case "\n\t":
                        var v17 float32 = v3
                        _ = v17
                    case "foo bar":
                        _ = len(v1)
                    default:
                        v1 = append(v1, "hello")
                    }
                case -1:
                    v4 -= float64((1.5 / (((0.0 - 0.1) / v3) + 0.0)))
                default:
                    v18 := GenS1{F0: v2, F1: 255, F2: (((v4 + v4) < (float64((v3 / 1.0)) + v4)) && true)}
                    _ = v18
                }
                v2 = v2
            } else {
                for fi19 := 0; fi19 < 2; fi19++ {
                    for fi20 := 0; fi20 < 2; fi20++ {
                        if (4294967295 <= 983) {
                            var v21 uint = 4294967295
                            v3 = (0.0 - v3)
                            if (true && (94 < -2147483648)) {
                                var v22 uint16 = 255
                                _ = v22
                            } else {
                                var v23 bool = false
                                fmt.Println(v21)
                                var v24 float32 = v3
                                _ = v23
                                _ = v24
                            }
                            fmt.Println(v4, v2, fi19)
                            v2 = (v2 + 58)
                            _ = v21
                        }
                        var v25 int16 = -32768
                        _ = fi20
                        _ = v25
                    }
                    _ = fi19
                }
            }
        }
    default:
        var v26 bool = false
        _ = v26
    }
    if (!(false || (2147483647 > 9223372036854775807))) {
        v27 := GenS2{F0: 1, F1: 1}
        v2 -= v2
        if false {
            var v28 int16 = -32768
            v29 := GenS4{F0: 725}
            if (-335 < -9223372036854775808) {
                v28 = v28
            } else {
                var v30 bool = (!false)
                _ = v30
            }
            _ = v28
            _ = v29
        }
        _ = v27
    }
    v4 += (float64((v3 + 0.1)) * float64(-2.25))
    for fi31 := 0; fi31 < 11; fi31++ {
        switch (fi31 / ((-9223372036854775808) | 1)) {
        case 127:
            if true {
                if false {
                    v32 := GenS5{F0: 1}
                    _ = v32
                } else {
                    switch "foo bar" {
                    case "":
                        v3 = (v3 - 3.14159)
                    default:
                        v2 = 1
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v33 uint64 = 42
                    _ = v33
                }
            }
        case 2147483647:
            v34 := GenS4{F0: 2147483647}
            _ = v34
        case -9223372036854775808:
            if (!(("hello" != "hello") && true)) {
                fmt.Println(fi31)
                switch "foo bar" {
                case "":
                    v2 -= v2
                default:
                    for fi35 := 0; fi35 < 9; fi35++ {
                        if (-1 <= -2147483648) {
                            v3 += v3
                            if true {
                                var v36 string = ("a" + "a")
                                v2 += (v2 ^ v2)
                                if (65535 >= 42) {
                                    v37 := GenS1{F0: (v2 / ((v2) | 1)), F1: 2, F2: (true && true)}
                                    if len(v1) > 0 {
                                        v1[0] = "foo bar"
                                    }
                                    if true {
                                        _ = len(v1)
                                        fmt.Println(fi35, v4, fi35)
                                        v38 := GenS3{F0: (false || (1 == 33))}
                                        _ = v38
                                    } else {
                                        var v39 uint32 = 732
                                        _ = v39
                                    }
                                    _ = v37
                                }
                                _ = v36
                            } else {
                                if (1 <= 128) {
                                    var v40 int16 = 0
                                    v40 = v40
                                    fmt.Println(fi31, fi31)
                                    switch 127 {
                                    case 128:
                                        v3 = (3.14159 / -2.25)
                                    case 32767:
                                        var v41 float32 = float32(1.5)
                                        _ = v41
                                    case -45:
                                        v3 = (1e10 / v3)
                                    default:
                                        switch "\n\t" {
                                        case "":
                                            if len(v1) > 0 {
                                                v1[0] = ("a" + "a")
                                            }
                                            v2 = (v2 & (v2 / (((v2 * v2)) | 1)))
                                        case "foo bar":
                                            fmt.Println(v40, fi35)
                                        default:
                                            v42 := GenS6{F0: 0, F1: 1}
                                            _ = v42
                                        }
                                    }
                                    _ = v40
                                }
                            }
                        }
                        v43 := GenS7{F0: (125 >= 1), F1: 2}
                        _ = fi35
                        _ = v43
                    }
                }
            } else {
                var v44 int32 = -2147483648
                _ = v44
            }
        default:
            for fi45 := 0; fi45 < 7; fi45++ {
                v3 -= float32(1e10)
                if true {
                    if true {
                        v46 := GenS3{F0: true}
                        _ = v46
                    } else {
                        fmt.Println(fi45)
                        switch "foo bar" {
                        case "\n\t":
                            v4 = float64(v3)
                        case "":
                            v4 = v4
                        case "foo bar":
                            var v47 float64 = 1.5
                            _ = v47
                        }
                        fmt.Println(v4, v3)
                    }
                } else {
                    var v48 int32 = 42
                    _ = len(v1)
                    _ = v48
                }
                if len(v1) > 0 {
                    v1[0] = ""
                }
                _ = fi45
            }
        }
        _ = fi31
    }
    if false {
        if (v2 < v2) {
            if (((("a" + ("\n\t" + "\n\t")) + "\n\t") + "foo bar") >= ("foo bar" + ("foo bar" + "hello"))) {
                v2 = v2
            } else {
                v4 -= float64((((3.14159 + v3) - float32(v4)) - (v3 / ((3.14159 / v3) * (float32(v4) + float32(v4))))))
                v2 = 116
            }
            v49 := GenS6{F0: 112, F1: 255}
            var v50 float64 = (((v4 / v4) - 0.1) / ((0.0 + float64(v3)) + 0.0))
            _ = v49
            _ = v50
        } else {
            if (505 >= 65535) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v3 -= ((v3 + v3) * ((float32((3.14159 * ((v4 / ((v4 + (v4 - v4)) / (0.1 + (v4 + v4)))) + -2.25))) / v3) - v3))
                }
            }
            switch 32767 {
            case 27:
                var v51 string = ("foo bar" + ("hello" + ((((("" + "a") + "") + (("a" + (("\n\t" + "\n\t") + "")) + "\n\t")) + (("a" + "") + (("\n\t" + ("hello" + "")) + "\n\t"))) + ("a" + ""))))
                _ = v51
            case 168:
                v52 := GenS8{F0: 255, F1: 733, F2: 128}
                _ = v52
            }
            switch -2147483648 {
            case -924:
                if false {
                    v2 += (-(v2))
                    v4 = 0.0
                } else {
                    fmt.Println(v4, v3, v2)
                }
            case 128:
                var v53 int = 42
                _ = v53
            }
        }
    }
    v54 := GenS6{F0: 2147483648, F1: 0}
    v3 = (v3 / v3)
    if len(v1) > 0 {
        v1[0] = "a"
    }
    for fi55 := 0; fi55 < 20; fi55++ {
        switch ("a" + "hello") {
        case "hello":
            v56 := GenS9{F0: false, F1: 0}
            _ = v56
        case "\n\t":
            fmt.Println(v3, v3)
        case "a":
            if (!(!(1.5 > 1e10))) {
                v57 := GenS10{F0: 42, F1: "hello", F2: 2}
                for fi58 := 0; fi58 < 3; fi58++ {
                    if ((18446744073709551615 > 42) || ((fi55 != -300) && true)) {
                        v59 := GenS2{F0: fi55, F1: -381}
                        fmt.Println(v2, v2, v2)
                        if ((2 < 377) || false) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v2 += v2
                        } else {
                            v4 -= (((v4 * (v4 - (v4 + v4))) / ((float64(-2.25) / float64(-2.25)) * v4)) / v4)
                            v2 += v2
                        }
                        _ = v59
                    } else {
                        if (!(!true)) {
                            v4 = v4
                        } else {
                            v4 = (float64(3.14159) - float64((float32((3.14159 - float64((v3 + 0.0)))) + -2.25)))
                            v60 := GenS9{F0: false, F1: 32767}
                            switch 9223372036854775807 {
                            case -9223372036854775808:
                                if ((((!false) || false) || true) && (1 == 2)) {
                                    var v61 uint16 = 858
                                    var v62 int16 = 32767
                                    _ = v61
                                    _ = v62
                                }
                            default:
                                fmt.Println(fi58, v2)
                            }
                            _ = v60
                        }
                        switch ("a" + ("" + "hello")) {
                        case "a":
                            v63 := GenS11{F0: 870}
                            _ = v63
                        case "foo bar":
                            v3 += float32(float64(-2.25))
                        case "hello":
                            v64 := GenS8{F0: 0, F1: 516, F2: -461}
                            _ = v64
                        default:
                            if true {
                                switch ("\n\t" + "hello") {
                                case "\n\t":
                                    fmt.Println(v4, v4)
                                }
                            }
                        }
                        var v65 int16 = 0
                        _ = v65
                    }
                    _ = fi58
                }
                _ = v57
            }
        default:
            for fi66 := 0; fi66 < 14; fi66++ {
                switch "\n\t" {
                case "foo bar":
                    v3 = 1.5
                }
                var v67 uint32 = 4294967295
                v3 -= -2.25
                fmt.Println(fi55)
                _ = fi66
                _ = v67
            }
        }
        for fi68 := 0; fi68 < 15; fi68++ {
            v2 += (v2 | v2)
            v3 = ((3.14159 / 0.1) + 0.0)
            _ = len(v1)
            _ = fi68
        }
        if (2147483647 <= 0) {
            _ = len(v1)
            for fi69 := 0; fi69 < 13; fi69++ {
                v1 = append(v1, (((((("foo bar" + "") + "") + "a") + "\n\t") + "hello") + "\n\t"))
                var v70 float64 = (v4 - v4)
                switch fi55 {
                case 127:
                    var v71 uint = 951
                    _ = v71
                }
                _ = fi69
                _ = v70
            }
            switch "hello" {
            case "foo bar":
                var v72 bool = true
                _ = v72
            default:
                v2 += 1
            }
        }
        _ = fi55
    }
    if len(v1) > 0 {
        v1[0] = ""
    }
    if len(v1) > 0 {
        v1[0] = ("hello" + ("foo bar" + "foo bar"))
    }
    for fi73 := 0; fi73 < 3; fi73++ {
        if (2147483647 < 4294967295) {
            switch 32767 {
            case -902:
                v4 = (v4 - float64(v3))
            case 4294967295:
                for fi74 := 0; fi74 < 20; fi74++ {
                    switch "hello" {
                    case "":
                        var v75 int16 = -994
                        _ = v75
                    }
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                    _ = fi74
                }
            }
            v2 -= -104
        }
        _ = fi73
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v54
}
