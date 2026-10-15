package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int; F1 uint16; F2 uint }
type GenS9 struct { F0 int; F1 int16; F2 uint8 }
type GenS10 struct { F0 string; F1 uint16; F2 string }
func genFn0(p0 int, p1 int64) uint {
    switch "\n\t" {
    case "\n\t":
        p1 = (p1 ^ p1)
    default:
        for fi5 := 0; fi5 < 2; fi5++ {
            for fi6 := 0; fi6 < 12; fi6++ {
                if true {
                    var v7 float32 = 3.14159
                    v7 += v7
                    _ = v7
                }
                var v8 int16 = -134
                v8 -= 0
                v9 := []int{}
                v10 := GenS1{F0: 42, F1: 135}
                _ = fi6
                _ = v10
                _ = v8
                _ = v9
            }
            _ = fi5
        }
    }
    if (!true) {
        var v11 uint32 = 2
        for fi12 := 0; fi12 < 13; fi12++ {
            v11 -= (v11 / (((v11 - v11)) | 1))
            if ((^(v11)) > 2147483647) {
                var v13 int64 = (p1 - p1)
                for fi14 := 0; fi14 < 8; fi14++ {
                    var v15 string = "foo bar"
                    v16 := GenS1{F0: 255, F1: 0}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                if false {
                    var v17 int8 = 0
                    switch -421 {
                    case -2147483648:
                        v18 := GenS2{F0: 665, F1: 1}
                        _ = v18
                    case 0:
                        var v19 int = -9223372036854775808
                        _ = v19
                    case 4294967295:
                        if (((1e10 - 0.0) * (0.0 / -2.25)) >= (0.0 + (3.14159 / (1e10 * (1e10 + -2.25))))) {
                            var v20 int16 = 1000
                            _ = v20
                        } else {
                            if (156 >= 2) {
                                if false {
                                    if (0 == 42) {
                                        if true {
                                            var v21 uint32 = 2147483648
                                            switch (fi12 | 42) {
                                            case -1:
                                                v22 := []int{}
                                                _ = v22
                                            case 2147483647:
                                                v21 = 351
                                            default:
                                                var v23 int64 = (p1 * v13)
                                                _ = v23
                                            }
                                            _ = v21
                                        }
                                        var v24 float64 = (3.14159 * 1e10)
                                        _ = v24
                                    }
                                    v25 := GenS2{F0: 873, F1: v17}
                                    _ = v25
                                }
                                switch "hello" {
                                case "a":
                                    p1 = (v13 % (((p1 * (v13 ^ (^(p1))))) | 1))
                                case "":
                                    v26 := []int{int(v11), p0, (fi12 & (^(fi12))), (fi12 << ((392) & 63))}
                                    _ = v26
                                default:
                                    switch p0 {
                                    case 9223372036854775807:
                                        v27 := GenS1{F0: 42, F1: 65535}
                                        _ = v27
                                    default:
                                        fmt.Println(v13, v13)
                                    }
                                }
                                v28 := GenS3{F0: 117}
                                _ = v28
                            } else {
                                v29 := GenS4{F0: 18446744073709551615}
                                for fi30 := 0; fi30 < 15; fi30++ {
                                    v31 := []int{}
                                    _ = fi30
                                    _ = v31
                                }
                                fmt.Println(v17, fi12, fi12)
                                _ = v29
                            }
                            v32 := []int{(^(p0)), p0, fi12, -345}
                            v13 -= (-(v13))
                            _ = v32
                        }
                    default:
                        switch -9223372036854775808 {
                        case 1:
                            if false {
                                var v33 int32 = 2147483647
                                p0 = fi12
                                _ = v33
                            } else {
                                var v34 uint = 18446744073709551615
                                _ = v34
                            }
                        case -9223372036854775808:
                            switch fi12 {
                            case -374:
                                v35 := []int{}
                                _ = v35
                            default:
                                switch "hello" {
                                case "\n\t":
                                    p1 = p1
                                case "a":
                                    v36 := []int{(p0 + (p0 - p0))}
                                    _ = v36
                                }
                            }
                        case 758:
                            fmt.Println(p0, fi12)
                        default:
                            var v37 int8 = v17
                            _ = v37
                        }
                    }
                    var v38 int32 = 1
                    _ = v17
                    _ = v38
                }
                _ = v13
            }
            _ = fi12
        }
        _ = v11
    }
    v39 := GenS1{F0: 32767, F1: 255}
    for fi40 := 0; fi40 < 18; fi40++ {
        fmt.Println(fi40, p0)
        p1 += 42
        var v41 uint = 4294967295
        _ = fi40
        _ = v41
    }
    var v42 uint = 65535
    _ = p0
    _ = p1
    _ = v39
    _ = v42
    return (^(v42))
}

func main() {
    var v1 uint16 = 0
    v2 := []string{("\n\t" + "foo bar"), "\n\t", (((("hello" + "\n\t") + "\n\t") + "a") + "\n\t")}
    v3 := make(map[string]int)
    var v4 int32 = 1
    delete(v3, "")
    fmt.Println(v4)
    switch ((("hello" + "foo bar") + ("foo bar" + "a")) + "foo bar") {
    case "\n\t":
        switch -2147483648 {
        case -2147483648:
            if false {
                v43 := GenS5{F0: 0, F1: 32767}
                v4 = 452
                _ = v43
            } else {
                v1 += 0
                v44 := GenS6{F0: ((!true) && (!true)), F1: 735}
                var v45 uint32 = 0
                _ = v44
                _ = v45
            }
        case 4294967295:
            switch 2147483647 {
            case 32768:
                switch ("" + ("" + "hello")) {
                case "foo bar":
                    for fi46 := 0; fi46 < 16; fi46++ {
                        if false {
                            for fi47 := 0; fi47 < 6; fi47++ {
                                var v48 int8 = 42
                                if false {
                                    fmt.Println(v1, v1, fi46)
                                    v49 := GenS2{F0: 2, F1: (v48 ^ v48)}
                                    _ = v49
                                }
                                _ = fi47
                                _ = v48
                            }
                        } else {
                            if len(v2) > 0 {
                                v2[0] = "a"
                            }
                            var v50 int32 = 42
                            var v51 float32 = 0.1
                            _ = v50
                            _ = v51
                        }
                        _ = fi46
                    }
                case "a":
                    v4 += (v4 % ((v4) | 1))
                case "":
                    if false {
                        for fi52 := 0; fi52 < 7; fi52++ {
                            for fi53 := 0; fi53 < 3; fi53++ {
                                v1 -= (^(v1))
                                _ = fi53
                            }
                            v1 -= v1
                            _ = fi52
                        }
                        v4 += v4
                    }
                default:
                    var v54 uint8 = 49
                    _ = v54
                }
            case 9223372036854775807:
                if (-9223372036854775808 >= 42) {
                    switch (("hello" + "foo bar") + "hello") {
                    case "hello":
                        switch "\n\t" {
                        case "\n\t":
                            var v55 float64 = 0.0
                            _ = v55
                        case "a":
                            var v56 uint16 = (^(v1))
                            _ = v56
                        case "hello":
                            v57 := GenS5{F0: -32768, F1: -32768}
                            _ = v57
                        }
                    case "a":
                        v4 += 42
                    case "foo bar":
                        v4 = (v4 / ((v4) | 1))
                    default:
                        var v58 uint = 807
                        _ = v58
                    }
                } else {
                    fmt.Println(v4)
                    if true {
                        v59 := GenS6{F0: (false && true), F1: 18446744073709551615}
                        _ = v59
                    }
                    var v60 int8 = 127
                    _ = v60
                }
            case 444:
                v4 += v4
            }
        case -15:
            fmt.Println(v4, v1)
        }
    case "":
        if ((-77 <= -389) && true) {
            delete(v3, ("a" + "a"))
            if true {
                for fi61 := 0; fi61 < 20; fi61++ {
                    for fi62 := 0; fi62 < 7; fi62++ {
                        v4 += v4
                        v1 += (v1 - (v1 - (v1 - (v1 - 662))))
                        switch (fi62 & fi61) {
                        case -9223372036854775808:
                            if false {
                                var v63 uint16 = v1
                                _ = v63
                            }
                        case -1:
                            switch fi62 {
                            case -636:
                                var v64 uint = 255
                                _ = v64
                            default:
                                var v65 uint = 924
                                _ = v65
                            }
                        case 127:
                            _ = len(v2)
                        default:
                            _ = len(v3)
                        }
                        v1 = (v1 % (((^(v1))) | 1))
                        _ = fi62
                    }
                    _ = fi61
                }
                if (169 != uint64(v1)) {
                    v1 = (v1 + v1)
                } else {
                    var v66 int = 773
                    _ = v66
                }
                for fi67 := 0; fi67 < 5; fi67++ {
                    var v68 uint64 = 263
                    var v69 float32 = float32(-2.25)
                    _ = fi67
                    _ = v68
                    _ = v69
                }
            }
        } else {
            for fi70 := 0; fi70 < 6; fi70++ {
                fmt.Println(fi70, v4)
                _ = fi70
            }
            for fi71 := 0; fi71 < 16; fi71++ {
                var v72 int = (-(fi71))
                if (738 != 18446744073709551615) {
                    v1 = 42
                    v1 += v1
                }
                for fi73 := 0; fi73 < 6; fi73++ {
                    fmt.Println(fi73, fi71, v1)
                    if true {
                        switch (("\n\t" + (("\n\t" + ("hello" + "foo bar")) + ((("hello" + "\n\t") + "hello") + "a"))) + ((("foo bar" + ("foo bar" + ("\n\t" + ""))) + (("" + ("hello" + "a")) + "foo bar")) + "foo bar")) {
                        case "foo bar":
                            var v74 int32 = 0
                            _ = v74
                        case "a":
                            var v75 float32 = 1e10
                            _ = v75
                        default:
                            v76 := GenS3{F0: genFn0((v72 + (fi71 * fi73)), 0)}
                            _ = v76
                        }
                    }
                    if true {
                        var v77 uint8 = 2
                        var v78 float32 = ((float32(-2.25) * -2.25) - (0.0 + 1e10))
                        _ = v77
                        _ = v78
                    }
                    _ = fi73
                }
                _ = fi71
                _ = v72
            }
        }
    case "foo bar":
        for fi79 := 0; fi79 < 4; fi79++ {
            v1 = v1
            _ = fi79
        }
    default:
        v1 += (v1 | (v1 * (v1 >> ((4294967295) & 63))))
    }
    if ((true || false) || (int16(v1) < -540)) {
        if (100 == -13) {
            switch 42 {
            case 127:
                v80 := GenS5{F0: 42, F1: 279}
                _ = v80
            case 0:
                v4 = (v4 | -625)
            case 128:
                for fi81 := 0; fi81 < 19; fi81++ {
                    v1 = v1
                    v4 = -1
                    v3["\n\t"] = 42
                    _ = fi81
                }
            }
            var v82 int8 = -1
            v83 := GenS4{F0: 472}
            _ = v82
            _ = v83
        } else {
            if (!false) {
                var v84 float64 = 0.0
                var v85 uint64 = 716
                _ = v84
                _ = v85
            }
            if false {
                v86 := GenS7{F0: 2147483648}
                var v87 uint = 18446744073709551615
                _ = v86
                _ = v87
            }
            switch 4294967295 {
            case 245:
                for fi88 := 0; fi88 < 7; fi88++ {
                    var v89 uint16 = (v1 ^ (^(v1)))
                    _ = fi88
                    _ = v89
                }
            case 4294967295:
                v1 = v1
            default:
                for fi90 := 0; fi90 < 13; fi90++ {
                    fmt.Println(fi90, v4)
                    _ = fi90
                }
            }
        }
    }
    var v91 float64 = 3.14159
    switch 1 {
    case 32768:
        for fi92 := 0; fi92 < 5; fi92++ {
            v91 += (v91 + (v91 / v91))
            _ = fi92
        }
    }
    v4 = (-(v4))
    var v93 uint16 = v1
    v4 -= v4
    switch "foo bar" {
    case "\n\t":
        v91 -= v91
    case "hello":
        v94 := GenS8{F0: 32767, F1: v1, F2: 255}
        _ = v94
    case "foo bar":
        switch "a" {
        case "a":
            v2 = append(v2, "a")
        case "":
            switch "\n\t" {
            case "foo bar":
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
            case "hello":
                if true {
                    v95 := GenS9{F0: -1, F1: 32767, F2: 197}
                    for fi96 := 0; fi96 < 9; fi96++ {
                        v93 -= (^(v1))
                        fmt.Println(v91, fi96, v91)
                        for fi97 := 0; fi97 < 14; fi97++ {
                            if len(v2) > 0 {
                                v2[0] = ("" + "")
                            }
                            _ = fi97
                        }
                        _ = fi96
                    }
                    _ = v95
                }
            }
        case "\n\t":
            if (277 < 32768) {
                v91 -= 1e10
            }
        }
    }
    if ("" > ("" + "foo bar")) {
        var v98 uint32 = 42
        _ = v98
    } else {
        v91 = (1e10 * v91)
        v91 += v91
    }
    if false {
        v99 := GenS10{F0: "", F1: (v93 - (v93 - 255)), F2: "\n\t"}
        _ = v99
    }
    for fi100 := 0; fi100 < 15; fi100++ {
        _ = len(v2)
        _ = fi100
    }
    fmt.Println(v93, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v91:", v91)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v91
    _ = v93
}
