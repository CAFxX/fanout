package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int32; F1 bool; F2 uint16 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 uint; F1 string; F2 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool; F1 bool; F2 int64 }
type GenS8 struct { F0 bool; F1 int8; F2 bool }
type GenS9 struct { F0 uint64; F1 int }
type GenS10 struct { F0 uint64; F1 int; F2 int16 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint32 }

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[uint32]uint32)
    v3 := []uint32{2, 0, 65535}
    var v4 uint16 = 231
    fmt.Println(v1)
    if false {
        switch ("foo bar" + ("hello" + ("foo bar" + "a"))) {
        case "hello":
            switch "a" {
            case "a":
                v4 += 654
            case "":
                var v5 bool = (42 >= 2)
                _ = v5
            case "hello":
                switch int(v4) {
                case 1:
                    if len(v3) > 0 {
                        v3[0] = 2147483648
                    }
                case 32767:
                    v4 = v4
                default:
                    v6 := GenS1{F0: 2, F1: (true || (!(true && (false || true)))), F2: 18446744073709551615}
                    _ = v6
                }
            }
        case "a":
            v1 += -1
        case "foo bar":
            for fi7 := 0; fi7 < 11; fi7++ {
                var v8 bool = true
                switch (("foo bar" + "a") + "hello") {
                case "foo bar":
                    fmt.Println(v4, v1)
                case "a":
                    if (1e10 >= 1e10) {
                        switch -637 {
                        case 2147483647:
                            var v9 int32 = 2147483647
                            _ = v9
                        case -936:
                            for fi10 := 0; fi10 < 12; fi10++ {
                                fmt.Println(v8)
                                v8 = v8
                                _ = fi10
                            }
                        case 0:
                            var v11 int16 = -73
                            _ = v11
                        default:
                            delete(v2, 695)
                        }
                        v4 = v4
                        v12 := GenS1{F0: (v4 >> ((469) & 63)), F1: (!v8), F2: 255}
                        _ = v12
                    }
                }
                _ = len(v3)
                _ = fi7
                _ = v8
            }
        default:
            var v13 uint64 = 4294967295
            _ = v13
        }
        var v14 int8 = 0
        _ = v14
    } else {
        switch 128 {
        case -1:
            v15 := GenS1{F0: (v4 * (v4 >> ((2) & 63))), F1: false, F2: 498}
            _ = v15
        case 2147483647:
            for fi16 := 0; fi16 < 14; fi16++ {
                var v17 uint32 = 2147483647
                var v18 bool = (!(255 <= 641))
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        v19 := v2[65535]
        var v20 int64 = -9223372036854775808
        _ = v19
        _ = v20
    }
    switch "foo bar" {
    case "\n\t":
        _ = len(v3)
    }
    var v21 uint = 42
    switch ("" + "a") {
    case "foo bar":
        v4 = 2
    case "":
        var v22 uint8 = 0
        _ = v22
    }
    if (!true) {
        for fi23 := 0; fi23 < 18; fi23++ {
            v24 := GenS1{F0: (^(v4)), F1: true, F2: v21}
            v25 := GenS1{F0: v4, F1: ("foo bar" <= ((("a" + "a") + "hello") + "\n\t")), F2: v21}
            _ = fi23
            _ = v24
            _ = v25
        }
        var v26 uint32 = 4294967295
        _ = v26
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        fmt.Println(v4, v1, v4)
        for fi28 := 0; fi28 < 7; fi28++ {
            v29 := GenS1{F0: v4, F1: (2 >= 42), F2: v21}
            v21 = (v21 | (^(v21)))
            v1 = -1
            _ = fi28
            _ = v29
        }
        if false {
            _ = len(v3)
        } else {
            switch ("hello" + "") {
            case "a":
                switch -2147483648 {
                case 32768:
                    var v30 int8 = 1
                    _ = v30
                case 1:
                    if (285 == 65535) {
                        v31 := GenS2{F0: ("" + ((("" + "") + "hello") + ("" + (("" + "foo bar") + "\n\t")))), F1: -102}
                        fmt.Println(v4, v21, v1)
                        _ = v31
                    }
                case 878:
                    for fi32 := 0; fi32 < 13; fi32++ {
                        fmt.Println(v1)
                        var v33 int16 = 0
                        _ = fi32
                        _ = v33
                    }
                default:
                    var v34 int32 = (v1 - (v1 * v1))
                    _ = v34
                }
            case "foo bar":
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v4, fi27, v4)
                }
            case "hello":
                v1 += -1
            default:
                v35 := GenS2{F0: ("\n\t" + ("a" + ("foo bar" + "\n\t"))), F1: -62}
                _ = v35
            }
        }
        _ = fi27
    }
    if false {
        for fi36 := 0; fi36 < 7; fi36++ {
            var v37 int16 = -32768
            if (fi36 != fi36) {
                for fi38 := 0; fi38 < 17; fi38++ {
                    var v39 uint8 = 42
                    _ = fi38
                    _ = v39
                }
            } else {
                switch 32768 {
                case 32768:
                    var v40 uint64 = 874
                    _ = v40
                default:
                    var v41 float32 = -2.25
                    _ = v41
                }
                switch (fi36 - (fi36 & (fi36 + (fi36 - fi36)))) {
                case 0:
                    for fi42 := 0; fi42 < 9; fi42++ {
                        var v43 uint32 = 918
                        if (12 != 122) {
                            v37 += v37
                        }
                        _ = fi42
                        _ = v43
                    }
                case 1:
                    fmt.Println(fi36)
                case 127:
                    var v44 bool = (!(4294967295 == 780))
                    _ = v44
                default:
                    if false {
                        var v45 int32 = (^(v1))
                        v46 := GenS3{F0: (^(v45)), F1: ((!false) || (v45 > 42)), F2: uint16(v21)}
                        if true {
                            if false {
                                v21 += v21
                                switch "\n\t" {
                                case "foo bar":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case "a":
                                    v47 := GenS2{F0: "foo bar", F1: 1}
                                    _ = v47
                                case "hello":
                                    v48 := GenS4{F0: v4, F1: 255}
                                    _ = v48
                                }
                            }
                            var v49 bool = (!(false || (v21 <= 367)))
                            _ = v49
                        } else {
                            for fi50 := 0; fi50 < 18; fi50++ {
                                v51 := GenS2{F0: (("foo bar" + "a") + "a"), F1: -1}
                                v1 -= -2147483648
                                _ = fi50
                                _ = v51
                            }
                            v52 := GenS5{F0: (^(v21)), F1: ("\n\t" + "hello"), F2: (v37 | v37)}
                            _ = v52
                        }
                        _ = v45
                        _ = v46
                    } else {
                        for fi53 := 0; fi53 < 15; fi53++ {
                            if ((!true) && ((^(v37)) == (v37 + v37))) {
                                v2[2147483648] = 4294967295
                            } else {
                                v21 = (v21 / ((4294967295) | 1))
                                v54 := GenS5{F0: 0, F1: "hello", F2: (v37 / (((v37 + 1000)) | 1))}
                                switch 2147483647 {
                                case -324:
                                    switch (fi36 & (fi53 + fi53)) {
                                    case 1:
                                        v55 := GenS1{F0: 32767, F1: false, F2: v21}
                                        _ = v55
                                    case 127:
                                        if (1 == (v4 / (((^(v4))) | 1))) {
                                            fmt.Println(fi36)
                                            var v56 string = "\n\t"
                                            if (1.5 < 1e10) {
                                                var v57 int32 = v1
                                                _ = v57
                                            } else {
                                                v3 = append(v3, 2147483647)
                                                var v58 float64 = (1.5 * ((1.5 / float64(-2.25)) * 3.14159))
                                                _ = v58
                                            }
                                            _ = v56
                                        }
                                    default:
                                        var v59 uint8 = 2
                                        _ = v59
                                    }
                                }
                                delete(v2, 0)
                                _ = v54
                            }
                            switch (fi36 | (fi53 * (fi36 >> ((v21) & 63)))) {
                            case 32768:
                                fmt.Println(v4, v37, v1)
                            }
                            v2[0] = 952
                            _ = fi53
                        }
                        v60 := GenS6{F0: 65535}
                        _ = v60
                    }
                }
                var v61 uint16 = v4
                _ = v61
            }
            _ = fi36
            _ = v37
        }
    } else {
        v4 -= (v4 | v4)
        v62 := GenS5{F0: v21, F1: "hello", F2: 1000}
        if (2147483648 > 1) {
            v4 += v4
        } else {
            v4 = v4
            v21 = v21
            fmt.Println(v1)
        }
        _ = v62
    }
    switch 952 {
    case 4294967295:
        for fi63 := 0; fi63 < 5; fi63++ {
            for fi64 := 0; fi64 < 3; fi64++ {
                delete(v2, 42)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi64
            }
            if (!false) {
                if false {
                    fmt.Println(v21, v4, v4)
                    fmt.Println(v21)
                    if false {
                        var v65 uint16 = (v4 + v4)
                        _ = v65
                    }
                } else {
                    _ = len(v2)
                    var v66 float32 = (((1.5 - ((0.1 - 1e10) / (-2.25 * 3.14159))) + (float32(0.0) * 1.5)) / -2.25)
                    _ = v66
                }
            }
            _ = fi63
        }
    case -1:
        fmt.Println(v4, v1, v4)
    case -2147483648:
        v21 -= (v21 | (v21 ^ (v21 ^ 1)))
    default:
        for fi67 := 0; fi67 < 11; fi67++ {
            if (0.0 < 1e10) {
                v21 += 2
                if false {
                    for fi68 := 0; fi68 < 9; fi68++ {
                        var v69 float32 = 0.1
                        switch "foo bar" {
                        case "hello":
                            v4 = (v4 / ((v4) | 1))
                        }
                        v70 := GenS3{F0: -2147483648, F1: (-128 > int8(v1)), F2: (v4 - v4)}
                        _ = fi68
                        _ = v69
                        _ = v70
                    }
                    v21 = (v21 + v21)
                    v71 := GenS7{F0: (!(!((^(v4)) == (v4 - (v4 | v4))))), F1: false, F2: 0}
                    _ = v71
                }
                for fi72 := 0; fi72 < 3; fi72++ {
                    v1 = (v1 & v1)
                    _ = fi72
                }
            } else {
                for fi73 := 0; fi73 < 11; fi73++ {
                    switch fi73 {
                    case -776:
                        v2[42] = 0
                    case 128:
                        v4 = 0
                    case 32768:
                        switch fi73 {
                        case 32768:
                            if (!(9223372036854775807 <= 9223372036854775807)) {
                                switch 2147483647 {
                                case -37:
                                    v74 := GenS2{F0: "hello", F1: 100}
                                    _ = v74
                                default:
                                    var v75 uint = v21
                                    _ = v75
                                }
                                fmt.Println(v21)
                            } else {
                                v1 = v1
                                if true {
                                    if ((55 <= 0) || true) {
                                        v76 := GenS8{F0: false, F1: -53, F2: true}
                                        v77 := GenS9{F0: 2, F1: 128}
                                        v78 := GenS10{F0: 42, F1: 42, F2: 1}
                                        _ = v76
                                        _ = v77
                                        _ = v78
                                    } else {
                                        v79 := GenS11{F0: 194}
                                        _ = v79
                                    }
                                    v80 := GenS1{F0: (v4 - v4), F1: ((1 >= 100) || true), F2: (v21 << (((v21 % ((uint(v4)) | 1))) & 63))}
                                    _ = len(v3)
                                    switch (fi67 / (((fi67 / (((fi73 + 42)) | 1))) | 1)) {
                                    case 388:
                                        v81 := v2[42]
                                        _ = v81
                                    default:
                                        var v82 uint16 = 32767
                                        v83 := GenS4{F0: v4, F1: 128}
                                        fmt.Println(v1, v21)
                                        _ = v82
                                        _ = v83
                                    }
                                    _ = v80
                                }
                                v21 = 255
                            }
                        }
                    default:
                        var v84 int64 = 9223372036854775807
                        _ = v84
                    }
                    _ = fi73
                }
            }
            if len(v3) > 0 {
                v3[0] = 2147483648
            }
            _ = fi67
        }
    }
    if (0 != -1) {
        v85 := v2[2147483647]
        v4 = 255
        if (42 < 42) {
            if (false || ((-32768 >= 0) && true)) {
                var v86 int64 = -9223372036854775808
                var v87 string = ("hello" + "")
                _ = v86
                _ = v87
            }
        } else {
            var v88 uint16 = v4
            fmt.Println(v88, v85)
            v89 := GenS12{F0: v85}
            _ = v88
            _ = v89
        }
        _ = v85
    }
    var v90 float64 = (-2.25 - 0.0)
    for fi91 := 0; fi91 < 17; fi91++ {
        v3 = append(v3, 693)
        _ = len(v3)
        delete(v2, 639)
        _ = fi91
    }
    v92 := v2[2147483647]
    var v93 uint16 = (v4 ^ 2)
    v2[65535] = v92
    var v94 int8 = 127
    v95 := GenS9{F0: 551, F1: 9223372036854775807}
    var v96 int = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v90:", v90)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v90
    _ = v92
    _ = v93
    _ = v94
    _ = v95
    _ = v96
}
