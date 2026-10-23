package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 int32; F1 int }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int8 }
func genFn0(p0 uint32, p1 bool) uint32 {
    if (true || p1) {
        if (false && true) {
            v4 := GenS1{F0: 4}
            if false {
                v5 := []int{121, -786, 475, -1}
                _ = v5
            }
            fmt.Println(p1, p0)
            _ = v4
        } else {
            var v6 string = "\n\t"
            _ = v6
        }
        switch "" {
        case "\n\t":
            p0 = (^(p0))
        case "":
            v7 := []int{32767}
            _ = v7
        case "foo bar":
            v8 := GenS1{F0: 0}
            _ = v8
        }
    } else {
        fmt.Println(p1, p0)
    }
    if false {
        for fi9 := 0; fi9 < 5; fi9++ {
            for fi10 := 0; fi10 < 10; fi10++ {
                if true {
                    fmt.Println(fi9)
                    fmt.Println(p0)
                }
                if p1 {
                    p0 += (p0 % ((p0) | 1))
                    if (fi10 >= fi9) {
                        switch ("a" + "\n\t") {
                        case "a":
                            p0 = (p0 / ((719) | 1))
                        }
                        switch (fi10 ^ 4294967295) {
                        case -2147483648:
                            fmt.Println(fi10)
                        case -9223372036854775808:
                            v11 := GenS2{F0: "\n\t", F1: uint16(p0)}
                            _ = v11
                        case -1:
                            p0 = p0
                        default:
                            v12 := []int{942, fi10, 42}
                            _ = v12
                        }
                        v13 := []int{}
                        _ = v13
                    } else {
                        var v14 uint64 = 18446744073709551615
                        _ = v14
                    }
                    p1 = (2 > 63)
                    p0 = p0
                }
                switch "foo bar" {
                case "foo bar":
                    var v15 uint64 = 18446744073709551615
                    _ = v15
                case "hello":
                    var v16 string = "a"
                    p1 = (!p1)
                    if p1 {
                        v16 = v16
                        if p1 {
                            var v17 uint64 = uint64(p0)
                            v18 := GenS3{F0: v17}
                            _ = v17
                            _ = v18
                        } else {
                            var v19 int8 = 11
                            fmt.Println(fi10)
                            _ = v19
                        }
                    }
                    _ = v16
                }
                _ = fi10
            }
            var v20 int = fi9
            _ = fi9
            _ = v20
        }
    }
    _ = p0
    _ = p1
    return 758
}
func genFn1(p0 int64, p1 bool, p2 bool) int {
    fmt.Println(p2)
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]string)
    var v3 float64 = 0.0
    v3 = 3.14159
    v3 -= float64((float32((v3 - v3)) / 1.5))
    switch 42 {
    case -917:
        v3 = v3
    case -425:
        if false {
            if (1000 == 0) {
                v21 := GenS2{F0: (("hello" + "foo bar") + "hello"), F1: 32767}
                v3 = v3
                _ = v21
            }
            v3 += (v3 + (v3 - 3.14159))
            for fi22 := 0; fi22 < 13; fi22++ {
                var v23 int32 = 161
                v23 += (v23 / ((v23) | 1))
                var v24 float32 = float32(v3)
                _ = fi22
                _ = v23
                _ = v24
            }
        } else {
            var v25 uint64 = 106
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        v3 = v3
        _ = fi26
    }
    var v27 int64 = -339
    for fi28 := 0; fi28 < 18; fi28++ {
        v29 := []int{}
        if true {
            v2[fi28] = ("a" + ("\n\t" + ("foo bar" + "\n\t")))
            if (((0.1 * ((3.14159 / 3.14159) / (0.1 + v3))) * 3.14159) > (v3 + float64((-2.25 + 1.5)))) {
                if false {
                    var v30 int64 = 268
                    if (!false) {
                        var v31 int8 = 0
                        _ = v31
                    }
                    v32 := GenS2{F0: (("a" + "a") + "a"), F1: 255}
                    _ = v30
                    _ = v32
                }
                fmt.Println(v27, v3, v3)
                switch "" {
                case "a":
                    var v33 float32 = 0.1
                    _ = v33
                case "hello":
                    if false {
                        var v34 bool = true
                        v2[fi28] = "hello"
                        _ = v34
                    }
                case "":
                    v35 := GenS4{F0: (!(32767 <= 42)), F1: 42}
                    _ = v35
                default:
                    v2[-1] = (("a" + ("\n\t" + "")) + "")
                }
            } else {
                v36 := GenS1{F0: 0}
                switch "\n\t" {
                case "foo bar":
                    var v37 uint32 = 2
                    _ = v37
                case "a":
                    for fi38 := 0; fi38 < 14; fi38++ {
                        v39 := GenS2{F0: ("hello" + ""), F1: 1}
                        v40 := GenS3{F0: 895}
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                default:
                    var v41 int = genFn1(9223372036854775807, (!(!(100 != 127))), false)
                    _ = v41
                }
                switch genFn1(-816, false, false) {
                case 127:
                    for fi42 := 0; fi42 < 9; fi42++ {
                        switch (fi28 << ((65535) & 63)) {
                        case 299:
                            if true {
                                _ = len(v1)
                                var v43 uint = 2
                                v27 += -9223372036854775808
                                _ = v43
                            }
                        }
                        _ = fi42
                    }
                default:
                    var v44 int8 = -128
                    _ = v44
                }
                _ = v36
            }
            v29 = append(v29, (fi28 % (((fi28 | fi28)) | 1)))
        }
        _ = fi28
        _ = v29
    }
    switch "a" {
    case "foo bar":
        var v45 int16 = 42
        _ = v45
    case "":
        for fi46 := 0; fi46 < 4; fi46++ {
            var v47 int16 = 1
            if (!(true && true)) {
                if (0.0 >= float32(3.14159)) {
                    if true {
                        fmt.Println(v47)
                        switch (("foo bar" + "\n\t") + "a") {
                        case "\n\t":
                            for fi48 := 0; fi48 < 2; fi48++ {
                                var v49 float32 = -2.25
                                if (638 != 65535) {
                                    var v50 float32 = v49
                                    fmt.Println(fi48)
                                    _ = v50
                                }
                                var v51 int = fi46
                                _ = fi48
                                _ = v49
                                _ = v51
                            }
                        case "a":
                            for fi52 := 0; fi52 < 9; fi52++ {
                                var v53 int8 = -1
                                fmt.Println(v53)
                                var v54 uint8 = 1
                                _ = fi52
                                _ = v53
                                _ = v54
                            }
                        default:
                            var v55 bool = (v27 <= 1)
                            _ = v55
                        }
                    }
                    for fi56 := 0; fi56 < 19; fi56++ {
                        v57 := GenS5{F0: 42, F1: 127}
                        if false {
                            v58 := GenS5{F0: -2147483648, F1: genFn1(v27, (93 > 42), false)}
                            _ = v58
                        }
                        _ = fi56
                        _ = v57
                    }
                    if false {
                        var v59 uint = 18
                        _ = v59
                    } else {
                        for fi60 := 0; fi60 < 16; fi60++ {
                            v61 := GenS1{F0: -1}
                            var v62 uint = 159
                            _ = fi60
                            _ = v61
                            _ = v62
                        }
                        switch (fi46 % (((fi46 ^ -9223372036854775808)) | 1)) {
                        case 127:
                            if (!false) {
                                for fi63 := 0; fi63 < 13; fi63++ {
                                    if false {
                                        var v64 uint64 = 912
                                        _ = v64
                                    }
                                    v27 -= (v27 - v27)
                                    _ = fi63
                                }
                            }
                        }
                        var v65 int8 = 42
                        _ = v65
                    }
                }
                v66 := []int{(fi46 >> ((735) & 63))}
                v67 := GenS2{F0: ((("a" + "a") + ((((("" + "foo bar") + "") + ("a" + "a")) + "\n\t") + ("" + "\n\t"))) + "a"), F1: 895}
                _ = v66
                _ = v67
            } else {
                switch ("\n\t" + (("\n\t" + "a") + (("hello" + (("\n\t" + ("a" + "hello")) + ((("\n\t" + "") + "a") + "hello"))) + "a"))) {
                case "a":
                    var v68 int64 = 0
                    _ = v68
                case "foo bar":
                    v47 += (v47 % ((1) | 1))
                case "\n\t":
                    var v69 int16 = v47
                    _ = v69
                }
                v47 = v47
                var v70 string = "a"
                _ = v70
            }
            _ = fi46
            _ = v47
        }
    default:
        for fi71 := 0; fi71 < 6; fi71++ {
            v3 = v3
            v27 = 0
            for fi72 := 0; fi72 < 17; fi72++ {
                fmt.Println(v27, v27)
                var v73 int64 = v27
                _ = fi72
                _ = v73
            }
            _ = fi71
        }
    }
    if (true || ("hello" < "a")) {
        _ = len(v2)
    }
    fmt.Println(v3, v27)
    var v74 int64 = (v27 * 9223372036854775807)
    switch -9223372036854775808 {
    case 42:
        if (false || false) {
            if true {
                var v75 uint8 = 121
                v76 := GenS6{F0: -9223372036854775808, F1: 42}
                v27 -= -1
                _ = v75
                _ = v76
            }
            var v77 int16 = 277
            _ = v77
        } else {
            if (1 >= 0) {
                v78 := []int{32768, 2147483647, 128}
                _ = len(v78)
                if (228 <= 29) {
                    var v79 int32 = 42
                    v78 = append(v78, -180)
                    _ = v79
                }
                _ = v78
            } else {
                if true {
                    fmt.Println(v74)
                } else {
                    if (((v27 >> ((0) & 63)) != (v74 % (((v27 >> ((1) & 63))) | 1))) || (false || ((true && false) || true))) {
                        var v80 uint8 = 197
                        _ = v80
                    }
                }
            }
            v3 -= 3.14159
            if (true || (1.5 > ((3.14159 * (1e10 + 0.0)) + -2.25))) {
                if (1 <= -120) {
                    if false {
                        if (false || ((false && (true || true)) || (true && (255 <= 65535)))) {
                            v74 += 1
                            var v81 int16 = -453
                            var v82 int8 = 0
                            _ = v81
                            _ = v82
                        }
                    } else {
                        fmt.Println(v3, v3, v3)
                    }
                    var v83 string = ("a" + "a")
                    _ = v83
                } else {
                    var v84 uint32 = 189
                    if (27 != 42) {
                        var v85 int16 = -1
                        if (18446744073709551615 < 0) {
                            var v86 int = 218
                            _ = v86
                        }
                        _ = v85
                    } else {
                        v87 := GenS7{F0: 0}
                        var v88 uint64 = 0
                        v84 += v84
                        _ = v87
                        _ = v88
                    }
                    _ = v84
                }
                delete(v1, 128)
                v74 = 1
            } else {
                for fi89 := 0; fi89 < 5; fi89++ {
                    v27 -= v27
                    var v90 int32 = 1
                    fmt.Println(v74)
                    _ = fi89
                    _ = v90
                }
                v27 = (-(v27))
            }
        }
    }
    v91 := GenS8{F0: -49}
    for fi92 := 0; fi92 < 13; fi92++ {
        v27 = -161
        for fi93 := 0; fi93 < 18; fi93++ {
            delete(v1, (fi93 / (((^(fi93))) | 1)))
            var v94 int32 = 1
            fmt.Println(v94)
            _ = fi93
            _ = v94
        }
        _ = fi92
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v74
    _ = v91
}
