package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 bool }
type GenS2 struct { F0 uint64; F1 uint64; F2 int }
type GenS3 struct { F0 bool; F1 bool; F2 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64; F1 int; F2 uint32 }
type GenS7 struct { F0 int8; F1 uint; F2 int16 }
type GenS8 struct { F0 int16; F1 uint32; F2 string }
type GenS9 struct { F0 int16; F1 int64 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 bool; F1 bool; F2 uint64 }
type GenS12 struct { F0 int; F1 int8; F2 bool }
type GenS13 struct { F0 uint; F1 string }
type GenS14 struct { F0 uint64 }
func genFn0() uint16 {
    v5 := []int{-331, 4294967295}
    var v6 float32 = 1.5
    _ = v5
    _ = v6
    return 211
}
func genFn1(p0 int8, p1 uint32, p2 uint8) string {
    if true {
        for fi7 := 0; fi7 < 15; fi7++ {
            v8 := []int{(fi7 + fi7), (fi7 << ((2) & 63)), fi7, (fi7 * fi7)}
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := GenS1{F0: "a", F1: -32768, F2: (-503 >= -1)}
            for fi11 := 0; fi11 < 13; fi11++ {
                switch (-(fi9)) {
                case 2147483647:
                    p0 = (p0 | (p0 ^ (p0 % ((p0) | 1))))
                case 177:
                    if (((-851 >= -108) || ((1e10 * -2.25) >= 3.14159)) && true) {
                        fmt.Println(p2, p1, p0)
                        if (!(p2 > p2)) {
                            v12 := GenS1{F0: "", F1: -1, F2: (4294967295 < (fi9 - fi9))}
                            var v13 uint16 = 2
                            _ = v12
                            _ = v13
                        }
                        fmt.Println(fi11)
                        p2 = (p2 ^ (p2 - 251))
                    } else {
                        var v14 float32 = float32(0.0)
                        if true {
                            var v15 uint16 = 239
                            if false {
                                fmt.Println(v14)
                                switch "foo bar" {
                                case "\n\t":
                                    switch 0 {
                                    case -263:
                                        var v16 uint16 = 1
                                        _ = v16
                                    case -1:
                                        v15 -= v15
                                    default:
                                        if false {
                                            v17 := GenS2{F0: 2, F1: uint64(v15), F2: 32768}
                                            var v18 float64 = float64((((v14 * v14) + ((v14 * (v14 + float32(3.14159))) / v14)) * v14))
                                            _ = v17
                                            _ = v18
                                        } else {
                                            p1 -= 711
                                        }
                                    }
                                default:
                                    switch ("a" + "a") {
                                    case "hello":
                                        var v19 int64 = -520
                                        _ = v19
                                    case "\n\t":
                                        v15 = 65535
                                    }
                                }
                                v20 := GenS1{F0: "", F1: 1000, F2: false}
                                v15 = genFn0()
                                _ = v20
                            } else {
                                v15 = v15
                                var v21 int16 = -1
                                _ = v21
                            }
                            var v22 uint32 = (p1 * (p1 + 65535))
                            _ = v15
                            _ = v22
                        }
                        _ = v14
                    }
                case 97:
                    var v23 uint8 = (p2 ^ 127)
                    _ = v23
                default:
                    p0 += (p0 * (p0 & (-(p0))))
                }
                var v24 float64 = (1e10 - (-2.25 * (1e10 - 3.14159)))
                var v25 uint8 = p2
                var v26 uint16 = genFn0()
                var v27 uint64 = 2
                _ = fi11
                _ = v24
                _ = v25
                _ = v26
                _ = v27
            }
            _ = fi9
            _ = v10
        }
    } else {
        fmt.Println(p1, p2, p1)
    }
    if (-468 == 0) {
        p2 -= 2
        p1 = 2
        if true {
            var v28 float64 = -2.25
            fmt.Println(p2, p0, p1)
            v28 -= (v28 * v28)
            _ = v28
        } else {
            p1 = p1
            fmt.Println(p2, p2)
        }
    }
    switch int(p1) {
    case -529:
        for fi29 := 0; fi29 < 9; fi29++ {
            p2 = (p2 % (((p2 + 178)) | 1))
            p1 = p1
            var v30 int = 335
            _ = fi29
            _ = v30
        }
    default:
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = p2
    return "hello"
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]string)
    v3 := make(map[string]int)
    var v4 int32 = 1
    for fi31 := 0; fi31 < 10; fi31++ {
        v4 += v4
        if false {
            v4 = (v4 & (^(v4)))
        }
        _ = fi31
    }
    if true {
        v32 := GenS1{F0: "\n\t", F1: -1, F2: (("a" + ("hello" + "")) >= "\n\t")}
        var v33 float32 = 3.14159
        _ = v32
        _ = v33
    } else {
        v34 := []int{-935, -242, 600, -1}
        var v35 int64 = -9223372036854775808
        _ = v34
        _ = v35
    }
    v36 := GenS3{F0: ("a" != "foo bar"), F1: ("foo bar" < "hello"), F2: (v4 % ((-2147483648) | 1))}
    var v37 int16 = 1
    v38 := []int{4294967295}
    if false {
        for fi39 := 0; fi39 < 11; fi39++ {
            var v40 int = 32767
            var v41 float32 = -2.25
            _ = fi39
            _ = v40
            _ = v41
        }
    } else {
        v42 := GenS4{F0: (^(v4))}
        for fi43 := 0; fi43 < 2; fi43++ {
            var v44 uint32 = 4294967295
            _ = fi43
            _ = v44
        }
        v45 := GenS5{F0: 1}
        _ = v42
        _ = v45
    }
    switch "hello" {
    case "hello":
        switch ((("\n\t" + ("foo bar" + "foo bar")) + "hello") + "hello") {
        case "hello":
            switch "" {
            case "a":
                var v46 uint32 = 685
                _ = v46
            case "foo bar":
                var v47 uint8 = 8
                _ = v47
            case "\n\t":
                v48 := GenS6{F0: 18446744073709551615, F1: 2147483647, F2: 2147483648}
                _ = v48
            default:
                var v49 float32 = 0.0
                _ = v49
            }
        case "foo bar":
            if (v37 >= (v37 | -647)) {
                var v50 int8 = 1
                for fi51 := 0; fi51 < 7; fi51++ {
                    var v52 uint32 = 2147483647
                    _ = len(v38)
                    _ = fi51
                    _ = v52
                }
                _ = v50
            }
        }
    case "\n\t":
        fmt.Println(v4, v37)
    case "a":
        v2[42] = ("" + "foo bar")
    default:
        var v53 int16 = (^(v37))
        _ = v53
    }
    if false {
        var v54 float64 = 0.0
        if (42 >= 1) {
            var v55 uint32 = 2
            for fi56 := 0; fi56 < 15; fi56++ {
                v38 = append(v38, fi56)
                switch "hello" {
                case "hello":
                    var v57 bool = true
                    _ = v57
                case "\n\t":
                    v58 := GenS1{F0: ("a" + (("\n\t" + "hello") + (("hello" + ("hello" + ("a" + "foo bar"))) + "foo bar"))), F1: (-(v37)), F2: true}
                    _ = v58
                case "foo bar":
                    for fi59 := 0; fi59 < 15; fi59++ {
                        var v60 uint16 = 1
                        _ = fi59
                        _ = v60
                    }
                default:
                    if len(v38) > 0 {
                        v38[0] = 317
                    }
                }
                switch fi56 {
                case 32767:
                    for fi61 := 0; fi61 < 4; fi61++ {
                        switch (fi61 ^ (^(fi61))) {
                        case 0:
                            v55 = (v55 - 1)
                        case 365:
                            switch fi56 {
                            case -537:
                                v55 -= v55
                            case 0:
                                delete(v1, (v55 & v55))
                            }
                        }
                        if true {
                            if true {
                                switch "foo bar" {
                                case "a":
                                    if false {
                                        if (fi56 <= (-(fi61))) {
                                            v55 = (v55 | 2147483647)
                                        } else {
                                            switch "\n\t" {
                                            case "hello":
                                                var v62 int32 = (v4 + (v4 + v4))
                                                _ = v62
                                            case "foo bar":
                                                v63 := GenS7{F0: 127, F1: 4294967295, F2: (v37 ^ (v37 & v37))}
                                                _ = v63
                                            }
                                        }
                                    }
                                case "\n\t":
                                    switch "foo bar" {
                                    case "":
                                        if len(v38) > 0 {
                                            v38[0] = (fi61 % (((fi56 & fi61)) | 1))
                                        }
                                    default:
                                        switch "" {
                                        case "foo bar":
                                            switch fi61 {
                                            case 4294967295:
                                                v55 = (v55 ^ v55)
                                            case 0:
                                                v4 = (v4 + v4)
                                                switch fi56 {
                                                case -2147483648:
                                                    switch fi56 {
                                                    case 0:
                                                        var v64 int = (fi56 | 0)
                                                        _ = v64
                                                    case 714:
                                                        v3[("foo bar" + (("hello" + "") + ("a" + ("a" + "a"))))] = 424
                                                        if len(v38) > 0 {
                                                            v38[0] = (fi61 - (fi61 ^ fi56))
                                                        }
                                                        var v65 uint16 = 1
                                                        _ = v65
                                                    case 4294967295:
                                                        var v66 string = "\n\t"
                                                        _ = v66
                                                    }
                                                case 4294967295:
                                                    switch (fi61 + 135) {
                                                    case -1:
                                                        if true {
                                                            var v67 bool = true
                                                            _ = v67
                                                        }
                                                    case 569:
                                                        var v68 uint = 581
                                                        _ = v68
                                                    case 991:
                                                        var v69 uint8 = 42
                                                        _ = v69
                                                    }
                                                default:
                                                    v4 = (v4 + v4)
                                                }
                                                fmt.Println(fi56, fi56, v37)
                                            case 42:
                                                _ = len(v38)
                                            default:
                                                if true {
                                                    v37 = int16(v4)
                                                }
                                            }
                                        case "\n\t":
                                            switch fi56 {
                                            case -1:
                                                v70 := GenS8{F0: (v37 | (v37 & (v37 << ((65535) & 63)))), F1: v55, F2: "\n\t"}
                                                _ = v70
                                            }
                                        case "a":
                                            var v71 uint64 = 42
                                            _ = v71
                                        default:
                                            fmt.Println(v4)
                                            v55 -= v55
                                            var v72 int64 = 42
                                            _ = v72
                                        }
                                    }
                                case "hello":
                                    if false {
                                        if (!false) {
                                            v4 -= -545
                                            v73 := GenS9{F0: 42, F1: -9223372036854775808}
                                            var v74 uint = 65535
                                            _ = v73
                                            _ = v74
                                        }
                                        v2[174] = "a"
                                    }
                                    v55 = v55
                                    if (42 > 42) {
                                        var v75 string = ""
                                        v76 := v3[(v75 + "foo bar")]
                                        _ = v75
                                        _ = v76
                                    }
                                }
                                fmt.Println(fi61, fi56)
                                v77 := GenS10{F0: 2}
                                v78 := GenS11{F0: false, F1: ((!(!(127 <= 0))) || (!(true || (fi56 != fi61)))), F2: 18446744073709551615}
                                var v79 int32 = 2147483647
                                _ = v77
                                _ = v78
                                _ = v79
                            }
                        }
                        _ = fi61
                    }
                case -2147483648:
                    for fi80 := 0; fi80 < 16; fi80++ {
                        var v81 float32 = (-2.25 - 3.14159)
                        var v82 bool = false
                        _ = fi80
                        _ = v81
                        _ = v82
                    }
                case 128:
                    v83 := GenS11{F0: false, F1: (((v4 > 42) || (true && true)) || true), F2: 386}
                    _ = v83
                }
                _ = fi56
            }
            for fi84 := 0; fi84 < 7; fi84++ {
                if len(v38) >= 2 {
                    v38 = v38[1:len(v38)]
                }
                v85 := GenS12{F0: fi84, F1: 0, F2: true}
                v55 -= v55
                _ = fi84
                _ = v85
            }
            _ = v55
        }
        for fi86 := 0; fi86 < 20; fi86++ {
            if (35 >= 595) {
                if false {
                    if (false && (false || (false && ((1 != 0) || true)))) {
                        var v87 float64 = 0.1
                        _ = v87
                    } else {
                        if (true || (2 < 1)) {
                            var v88 string = ("hello" + (((("hello" + "") + ("\n\t" + ("" + "hello"))) + "hello") + "\n\t"))
                            _ = v88
                        } else {
                            var v89 uint64 = 2
                            var v90 uint = 1
                            if len(v38) >= 2 {
                                v38 = v38[1:len(v38)]
                            }
                            _ = v89
                            _ = v90
                        }
                    }
                    switch fi86 {
                    case 128:
                        _ = len(v38)
                    case -436:
                        v54 += v54
                    default:
                        v38 = append(v38, fi86)
                    }
                } else {
                    switch ("a" + ("\n\t" + "")) {
                    case "\n\t":
                        v91 := GenS13{F0: 42, F1: ("foo bar" + "hello")}
                        _ = v91
                    case "foo bar":
                        var v92 int = fi86
                        _ = v92
                    case "a":
                        if (100 <= 107) {
                            var v93 string = "hello"
                            v94 := GenS11{F0: true, F1: (!(true && true)), F2: 2}
                            _ = v93
                            _ = v94
                        }
                    }
                    var v95 int64 = 1
                    v96 := GenS14{F0: 758}
                    _ = v95
                    _ = v96
                }
                if false {
                    v4 += v4
                } else {
                    var v97 float64 = v54
                    v54 -= v54
                    fmt.Println(v37)
                    _ = v97
                }
            } else {
                for fi98 := 0; fi98 < 19; fi98++ {
                    var v99 uint64 = 42
                    v54 = 1e10
                    _ = len(v38)
                    _ = fi98
                    _ = v99
                }
            }
            _ = fi86
        }
        _ = v54
    } else {
        v4 = (v4 & v4)
    }
    if (!true) {
        v2[-955] = "a"
        delete(v1, 4294967295)
        v37 -= v37
    } else {
        var v100 int = 673
        _ = v100
    }
    var v101 uint16 = 32767
    for fi102 := 0; fi102 < 2; fi102++ {
        var v103 uint64 = 2
        _ = fi102
        _ = v103
    }
    switch "foo bar" {
    case "hello":
        if false {
            v3["\n\t"] = 784
            v101 = 502
            var v104 bool = true
            _ = v104
        }
    case "":
        v4 = v4
    }
    var v105 uint = 485
    for fi106 := 0; fi106 < 9; fi106++ {
        switch ("a" + "hello") {
        case "foo bar":
            var v107 int32 = (v4 ^ (v4 % (((v4 ^ v4)) | 1)))
            _ = v107
        default:
            v108 := GenS12{F0: (fi106 + (fi106 % ((fi106) | 1))), F1: 127, F2: (("foo bar" + ("\n\t" + "a")) > "")}
            _ = v108
        }
        delete(v2, (fi106 / ((fi106) | 1)))
        _ = fi106
    }
    v105 += uint(v101)
    if true {
        v101 += (v101 & genFn0())
    }
    var v109 int8 = 112
    var v110 bool = (32768 < 32767)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v101:", v101)
    fmt.Println("v105:", v105)
    fmt.Println("v109:", v109)
    fmt.Println("v110:", v110)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v101
    _ = v105
    _ = v109
    _ = v110
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
