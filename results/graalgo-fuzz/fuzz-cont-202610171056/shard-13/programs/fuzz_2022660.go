package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64; F1 int32 }
func genFn0(p0 string, p1 int) string {
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: "a", F1: 42}
        v9 := []int{fi7, (fi7 % ((128) | 1)), (p1 + -9223372036854775808), fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(p1)
    fmt.Println(p0, p0)
    switch p1 {
    case 9223372036854775807:
        var v10 float32 = (3.14159 * 1.5)
        _ = v10
    default:
        switch p0 {
        case "\n\t":
            v11 := GenS1{F0: ("a" + ((p0 + (("foo bar" + (p0 + p0)) + p0)) + ("\n\t" + "\n\t"))), F1: 32767}
            _ = v11
        default:
            var v12 uint = 31
            _ = v12
        }
    }
    var v13 bool = (p0 != "foo bar")
    _ = p0
    _ = p1
    _ = v13
    return ("foo bar" + ("foo bar" + "foo bar"))
}

func main() {
    var v1 uint = 42
    v2 := []int{1}
    var v3 float64 = 0.1
    var v4 int = 516
    v5 := []int{(v4 / (((v4 >> (((v1 * 4294967295)) & 63))) | 1)), v4}
    var v6 int16 = 42
    v4 += -743
    if true {
        v2 = append(v2, 4294967295)
        var v14 bool = false
        for fi15 := 0; fi15 < 12; fi15++ {
            v6 -= (v6 * (v6 / ((v6) | 1)))
            if len(v2) > 0 {
                v2[0] = v4
            }
            for fi16 := 0; fi16 < 17; fi16++ {
                switch (fi16 * (fi15 & fi15)) {
                case 4294967295:
                    if (42 >= 9223372036854775807) {
                        fmt.Println(fi15, v6)
                        var v17 bool = (!((v1 | 255) < v1))
                        var v18 int64 = 1
                        _ = v17
                        _ = v18
                    }
                case -818:
                    var v19 int = (fi16 >> (((v1 << (((v1 + v1)) & 63))) & 63))
                    _ = v19
                default:
                    if len(v5) > 0 {
                        v5[0] = (fi16 << (((v1 % ((898) | 1))) & 63))
                    }
                    v4 += v4
                }
                v20 := GenS1{F0: "hello", F1: v6}
                switch ("" + (("" + "\n\t") + "hello")) {
                case "":
                    switch "" {
                    case "hello":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                case "hello":
                    v1 = v1
                case "\n\t":
                    v21 := GenS1{F0: "", F1: v6}
                    _ = v21
                default:
                    fmt.Println(v3, v3)
                    v22 := GenS2{F0: 0, F1: 2147483647, F2: 136}
                    var v23 bool = v14
                    _ = v22
                    _ = v23
                }
                _ = fi16
                _ = v20
            }
            _ = fi15
        }
        _ = v14
    }
    v24 := GenS2{F0: 770, F1: 2147483648, F2: v1}
    if (730 <= 1) {
        for fi25 := 0; fi25 < 3; fi25++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi25
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v26 := GenS1{F0: ("\n\t" + ("" + ("\n\t" + (("hello" + "foo bar") + "\n\t")))), F1: 42}
    for fi27 := 0; fi27 < 2; fi27++ {
        v28 := GenS3{F0: 100, F1: ""}
        fmt.Println(v1)
        if false {
            var v29 int16 = 0
            var v30 string = ("hello" + "a")
            v5 = append(v5, -742)
            _ = v29
            _ = v30
        } else {
            v6 -= (v6 + 42)
            var v31 uint32 = 1
            for fi32 := 0; fi32 < 4; fi32++ {
                v33 := GenS3{F0: -1, F1: ("foo bar" + "")}
                v4 = (fi32 * v4)
                var v34 float32 = 0.0
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = v31
        }
        _ = fi27
        _ = v28
    }
    switch "a" {
    case "":
        var v35 uint64 = 18446744073709551615
        _ = v35
    case "hello":
        switch "a" {
        case "a":
            v36 := GenS4{F0: -2147483648}
            _ = v36
        case "\n\t":
            switch "a" {
            case "foo bar":
                var v37 int = -944
                _ = v37
            case "hello":
                switch "\n\t" {
                case "hello":
                    v6 -= v6
                case "\n\t":
                    var v38 string = ("foo bar" + (("hello" + "foo bar") + ""))
                    _ = v38
                case "foo bar":
                    if (42 > 2147483647) {
                        v4 = (v4 * 0)
                        if true {
                            v39 := GenS2{F0: 0, F1: 4294967295, F2: v1}
                            var v40 float32 = (1e10 / ((1e10 / 1e10) + 0.1))
                            _ = len(v2)
                            _ = v39
                            _ = v40
                        } else {
                            var v41 uint32 = 794
                            var v42 uint64 = 497
                            v4 = (v4 >> (((v1 & v1)) & 63))
                            _ = v41
                            _ = v42
                        }
                        v1 -= v1
                    } else {
                        for fi43 := 0; fi43 < 7; fi43++ {
                            var v44 uint32 = 42
                            v1 += 4294967295
                            var v45 int8 = 63
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                    }
                }
            case "":
                if ((v1 & v1) <= v1) {
                    if false {
                        if true {
                            if (42 <= 138) {
                                for fi46 := 0; fi46 < 12; fi46++ {
                                    switch "\n\t" {
                                    case "foo bar":
                                        if ((1 < 100) || true) {
                                            switch "" {
                                            case "\n\t":
                                                var v47 uint16 = 0
                                                _ = v47
                                            case "hello":
                                                v48 := GenS4{F0: 2147483647}
                                                _ = v48
                                            default:
                                                switch (fi46 + 32767) {
                                                case 607:
                                                    var v49 uint32 = 917
                                                    _ = v49
                                                default:
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                }
                                            }
                                        } else {
                                            var v50 bool = false
                                            v51 := GenS4{F0: 42}
                                            v3 = (v3 - (((v3 * (-2.25 * v3)) + v3) + ((v3 + v3) + (v3 - v3))))
                                            _ = v50
                                            _ = v51
                                        }
                                    default:
                                        for fi52 := 0; fi52 < 15; fi52++ {
                                            v53 := GenS4{F0: -2147483648}
                                            var v54 uint64 = 0
                                            if false {
                                                switch (fi46 >> ((42) & 63)) {
                                                case 705:
                                                    v6 = (v6 ^ (v6 + (v6 % (((v6 | 1)) | 1))))
                                                case -691:
                                                    if ((!true) || (!true)) {
                                                        switch int(v1) {
                                                        case 515:
                                                            v54 = 330
                                                        case -312:
                                                            if false {
                                                                var v55 int32 = -622
                                                                _ = v55
                                                            }
                                                        default:
                                                            var v56 uint64 = (v54 / ((v54) | 1))
                                                            _ = v56
                                                        }
                                                    } else {
                                                        v5 = append(v5, -1)
                                                        if true {
                                                            fmt.Println(fi46, v54, fi52)
                                                        } else {
                                                            v1 += (v1 >> (((^(v1))) & 63))
                                                            v57 := GenS2{F0: 2147483647, F1: 401, F2: v1}
                                                            _ = v57
                                                        }
                                                    }
                                                case -2147483648:
                                                    var v58 uint64 = v54
                                                    _ = v58
                                                }
                                                fmt.Println(v54)
                                                switch fi46 {
                                                case -642:
                                                    fmt.Println(v3)
                                                case 4294967295:
                                                    if (true || true) {
                                                        var v59 uint16 = 65535
                                                        switch fi52 {
                                                        case -2147483648:
                                                            if ("hello" < "hello") {
                                                                switch (fi52 / ((v4) | 1)) {
                                                                case 2147483647:
                                                                    v6 = (v6 << ((v1) & 63))
                                                                case 40:
                                                                    v1 = v1
                                                                    v54 += (v54 << (((v1 / (((^(v1))) | 1))) & 63))
                                                                case 127:
                                                                    var v60 uint = v1
                                                                    _ = v60
                                                                }
                                                                var v61 uint8 = 186
                                                                if ((v6 == int16(v59)) && (219 <= 18446744073709551615)) {
                                                                    var v62 int64 = 1
                                                                    _ = v62
                                                                } else {
                                                                    fmt.Println(v59)
                                                                    v4 = (v4 ^ (fi46 % ((128) | 1)))
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                }
                                                                _ = v61
                                                            } else {
                                                                var v63 bool = ((0.0 * v3) > (v3 + 1e10))
                                                                if len(v2) > 0 {
                                                                    v2[0] = -1
                                                                }
                                                                if (!((0.1 <= 0.0) || ((true && false) || ((v1 ^ v1) < 0)))) {
                                                                    if (v54 != v54) {
                                                                        var v64 int32 = 2147483647
                                                                        _ = v64
                                                                    }
                                                                    if true {
                                                                        var v65 string = "\n\t"
                                                                        v63 = (105 > 42)
                                                                        _ = v65
                                                                    } else {
                                                                        v2 = append(v2, (fi52 & (^(v4))))
                                                                        var v66 float64 = (3.14159 - v3)
                                                                        v6 = (-(v6))
                                                                        _ = v66
                                                                    }
                                                                    fmt.Println(v1, fi52, v63)
                                                                }
                                                                _ = v63
                                                            }
                                                        }
                                                        v54 = (v54 & v54)
                                                        _ = v59
                                                    }
                                                default:
                                                    var v67 int32 = 613
                                                    _ = v67
                                                }
                                            }
                                            _ = fi52
                                            _ = v53
                                            _ = v54
                                        }
                                    }
                                    for fi68 := 0; fi68 < 9; fi68++ {
                                        v3 = 0.0
                                        v6 = v6
                                        _ = fi68
                                    }
                                    _ = fi46
                                }
                                switch "\n\t" {
                                case "\n\t":
                                    var v69 uint32 = 0
                                    _ = v69
                                case "a":
                                    for fi70 := 0; fi70 < 8; fi70++ {
                                        v2 = append(v2, -9223372036854775808)
                                        _ = fi70
                                    }
                                default:
                                    for fi71 := 0; fi71 < 17; fi71++ {
                                        if (((821 > 42) || true) || false) {
                                            if (!((v6 | v6) == (v6 >> ((v1) & 63)))) {
                                                _ = len(v2)
                                                for fi72 := 0; fi72 < 14; fi72++ {
                                                    var v73 uint16 = 2
                                                    _ = fi72
                                                    _ = v73
                                                }
                                            } else {
                                                v74 := GenS5{F0: 1}
                                                var v75 float64 = (0.0 - v3)
                                                _ = v74
                                                _ = v75
                                            }
                                        }
                                        var v76 uint64 = 541
                                        v1 -= (v1 * (^(v1)))
                                        _ = fi71
                                        _ = v76
                                    }
                                }
                            }
                        }
                        _ = len(v5)
                        var v77 int64 = -1
                        _ = v77
                    }
                    if false {
                        var v78 int = (v4 << ((2) & 63))
                        _ = v78
                    } else {
                        var v79 int = v4
                        _ = v79
                    }
                } else {
                    v6 = v6
                    v2 = append(v2, (v4 & -2147483648))
                    if true {
                        switch v4 {
                        case 436:
                            var v80 float32 = 0.1
                            _ = v80
                        case 42:
                            var v81 uint = 1
                            _ = v81
                        }
                        v82 := GenS6{F0: 349, F1: 582}
                        _ = v82
                    }
                }
            default:
                if (2147483648 < 2147483647) {
                    _ = len(v2)
                }
            }
        case "hello":
            v6 = (v6 << ((v1) & 63))
        default:
            fmt.Println(v1, v1)
        }
    }
    v3 -= (v3 + (v3 - (v3 + v3)))
    if (!false) {
        switch (v4 | (^(v4))) {
        case 701:
            var v83 int64 = 0
            _ = v83
        case -561:
            var v84 int16 = -400
            _ = v84
        case 32767:
            v1 = (v1 + v1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
