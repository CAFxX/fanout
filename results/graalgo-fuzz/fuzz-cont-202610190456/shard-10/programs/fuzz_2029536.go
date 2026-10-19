package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 int16; F1 int; F2 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint8; F1 uint16 }
type GenS9 struct { F0 int8; F1 uint64; F2 uint }
type GenS10 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS11 struct { F0 uint8 }
func genFn0(p0 int8, p1 uint8) uint {
    if false {
        for fi5 := 0; fi5 < 20; fi5++ {
            fmt.Println(fi5)
            _ = fi5
        }
    } else {
        for fi6 := 0; fi6 < 7; fi6++ {
            for fi7 := 0; fi7 < 7; fi7++ {
                p1 += p1
                var v8 float64 = 1.5
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 2; fi9++ {
                switch fi9 {
                case -2147483648:
                    v10 := GenS1{F0: 40, F1: 1, F2: 255}
                    _ = v10
                case 42:
                    fmt.Println(p1, p0)
                case 291:
                    p0 = (p0 * (^(p0)))
                }
                _ = fi9
            }
            _ = fi6
        }
    }
    fmt.Println(p1)
    switch 1 {
    case 4294967295:
        if true {
            for fi11 := 0; fi11 < 19; fi11++ {
                var v12 int64 = -1
                v13 := []int{9223372036854775807, (fi11 >> ((0) & 63)), fi11, -9223372036854775808}
                var v14 uint = 2
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            switch ("a" + "foo bar") {
            case "":
                p1 -= (p1 ^ 42)
            case "hello":
                fmt.Println(p1, p0)
            case "\n\t":
                var v15 bool = true
                _ = v15
            default:
                p1 -= (p1 / ((p1) | 1))
            }
            p1 = p1
        }
    case -912:
        var v16 uint64 = 42
        _ = v16
    case 2147483647:
        p0 = (p0 | (p0 << ((4294967295) & 63)))
    }
    var v17 int8 = (-(p0))
    v18 := GenS1{F0: 632, F1: 42, F2: 255}
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 4294967295
}

func main() {
    v1 := []uint32{4294967295}
    var v2 int64 = 0
    v3 := []int{-995, -34}
    var v4 uint32 = 546
    var v19 string = "a"
    for fi20 := 0; fi20 < 11; fi20++ {
        if (3.14159 != (0.0 + ((1e10 / 1.5) * (0.0 - (0.0 * 3.14159))))) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            var v21 uint8 = 2
            switch (fi20 / (((^(fi20))) | 1)) {
            case 919:
                fmt.Println(fi20)
            case 32767:
                if (65535 < 65535) {
                    var v22 uint32 = (v4 >> ((1) & 63))
                    _ = v22
                } else {
                    if false {
                        for fi23 := 0; fi23 < 18; fi23++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            switch v19 {
                            case "":
                                fmt.Println(fi23, v19)
                            case "hello":
                                v24 := GenS1{F0: v2, F1: 318, F2: 255}
                                _ = v24
                            case "foo bar":
                                switch ((("hello" + "hello") + v19) + (v19 + "")) {
                                case "a":
                                    v21 = (v21 & 127)
                                default:
                                    var v25 int16 = 0
                                    _ = v25
                                }
                            default:
                                v26 := GenS2{F0: (1.5 == float64(1e10))}
                                _ = v26
                            }
                            if len(v1) > 0 {
                                v1[0] = (v4 / (((^(v4))) | 1))
                            }
                            _ = fi23
                        }
                        switch fi20 {
                        case 2147483647:
                            if true {
                                v2 = (v2 << ((uint(v21)) & 63))
                            } else {
                                v2 -= v2
                                if false {
                                    if len(v1) > 0 {
                                        v1[0] = 42
                                    }
                                    v27 := GenS2{F0: false}
                                    _ = v27
                                }
                            }
                        default:
                            switch fi20 {
                            case 4294967295:
                                var v28 int8 = 0
                                _ = v28
                            default:
                                if true {
                                    v2 = v2
                                    for fi29 := 0; fi29 < 6; fi29++ {
                                        var v30 uint = genFn0(-1, (^(v21)))
                                        v30 += (v30 & 516)
                                        var v31 uint8 = v21
                                        switch "\n\t" {
                                        case "foo bar":
                                            v1 = append(v1, 4294967295)
                                        default:
                                            v3 = append(v3, 127)
                                        }
                                        switch "hello" {
                                        case "\n\t":
                                            var v32 uint = v30
                                            _ = v32
                                        default:
                                            v2 -= 763
                                        }
                                        switch (v19 + (v19 + "foo bar")) {
                                        case "a":
                                            v3 = append(v3, int(v31))
                                        case "\n\t":
                                            if len(v1) > 0 {
                                                v1[0] = (v4 - 810)
                                            }
                                        default:
                                            if len(v1) > 0 {
                                                v1[0] = v4
                                            }
                                        }
                                        v4 = 2147483647
                                        _ = fi29
                                        _ = v30
                                        _ = v31
                                    }
                                    for fi33 := 0; fi33 < 19; fi33++ {
                                        var v34 uint = genFn0(-107, (^(v21)))
                                        _ = fi33
                                        _ = v34
                                    }
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v35 float64 = 1.5
                                    if true {
                                        switch ("foo bar" + v19) {
                                        case "a":
                                            v2 = v2
                                        default:
                                            switch "hello" {
                                            case "\n\t":
                                                switch fi20 {
                                                case 2147483647:
                                                    fmt.Println(fi20, v21, v35)
                                                case 236:
                                                    var v36 float32 = 0.1
                                                    _ = v36
                                                case 497:
                                                    var v37 int = fi20
                                                    _ = v37
                                                default:
                                                    v38 := GenS3{F0: "\n\t", F1: 1}
                                                    _ = v38
                                                }
                                            case "hello":
                                                for fi39 := 0; fi39 < 10; fi39++ {
                                                    fmt.Println(v21, fi39, v21)
                                                    v40 := GenS2{F0: false}
                                                    _ = fi39
                                                    _ = v40
                                                }
                                            case "foo bar":
                                                var v41 bool = true
                                                _ = v41
                                            }
                                        }
                                    } else {
                                        fmt.Println(v35, v19)
                                        var v42 float64 = (0.0 + (v35 / v35))
                                        v19 = (v19 + "foo bar")
                                        _ = v42
                                    }
                                    _ = v35
                                }
                            }
                        }
                        for fi43 := 0; fi43 < 5; fi43++ {
                            if (1 == 265) {
                                v4 = 257
                                fmt.Println(v19, v21, v21)
                                switch fi43 {
                                case 260:
                                    var v44 string = (v19 + (((("foo bar" + (v19 + "hello")) + (v19 + v19)) + "foo bar") + (v19 + "")))
                                    _ = v44
                                }
                            } else {
                                fmt.Println(v19)
                                if (!false) {
                                    fmt.Println(fi43, fi43, fi43)
                                    v19 = v19
                                }
                                switch v19 {
                                case "":
                                    v4 -= 66
                                case "a":
                                    v3 = append(v3, (fi43 - (fi43 - 128)))
                                case "hello":
                                    if false {
                                        fmt.Println(v2, v2)
                                        if (false && false) {
                                            switch "" {
                                            case "foo bar":
                                                var v45 int32 = 0
                                                switch fi20 {
                                                case 1:
                                                    var v46 uint64 = 18446744073709551615
                                                    if (!true) {
                                                        if len(v1) > 0 {
                                                            v1[0] = v4
                                                        }
                                                        v19 = v19
                                                        v19 = v19
                                                    }
                                                    _ = v46
                                                case -9223372036854775808:
                                                    var v47 bool = (true || false)
                                                    _ = v47
                                                }
                                                _ = v45
                                            case "\n\t":
                                                _ = len(v1)
                                            case "hello":
                                                var v48 uint = 18446744073709551615
                                                _ = v48
                                            }
                                            if (false && (!true)) {
                                                if ((v4 | (v4 & 4294967295)) < (^(v4))) {
                                                    v19 = ("" + "a")
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v49 int16 = 32767
                                                    _ = v49
                                                } else {
                                                    v2 = -9223372036854775808
                                                }
                                                if false {
                                                    var v50 uint16 = 2
                                                    _ = v50
                                                } else {
                                                    v2 -= 732
                                                }
                                            } else {
                                                v2 -= (v2 << ((18446744073709551615) & 63))
                                                v21 -= v21
                                                v51 := GenS1{F0: int64(v4), F1: 558, F2: 42}
                                                var v52 uint = 2
                                                _ = v51
                                                _ = v52
                                            }
                                            var v53 bool = (255 == 706)
                                            _ = v53
                                        } else {
                                            v54 := GenS1{F0: v2, F1: 652, F2: 42}
                                            var v55 uint16 = 63
                                            if false {
                                                v1 = append(v1, 876)
                                                switch v19 {
                                                case "":
                                                    if true {
                                                        fmt.Println(fi20, fi43, fi43)
                                                    }
                                                case "\n\t":
                                                    v4 = 1
                                                default:
                                                    v19 = v19
                                                }
                                                switch "a" {
                                                case "\n\t":
                                                    switch fi20 {
                                                    case 32767:
                                                        var v56 string = (v19 + (("foo bar" + "foo bar") + v19))
                                                        _ = v56
                                                    default:
                                                        v57 := GenS1{F0: 42, F1: v55, F2: 65535}
                                                        _ = v57
                                                    }
                                                case "":
                                                    fmt.Println(fi20, v19, v2)
                                                default:
                                                    v19 = ("foo bar" + v19)
                                                }
                                            } else {
                                                v19 = v19
                                                if (true || (!false)) {
                                                    switch ("foo bar" + "a") {
                                                    case "foo bar":
                                                        if (-79 >= 41) {
                                                            v19 = v19
                                                            if true {
                                                                v58 := GenS4{F0: v55}
                                                                _ = v58
                                                            } else {
                                                                _ = len(v1)
                                                                var v59 int32 = 42
                                                                _ = v59
                                                            }
                                                            v60 := GenS1{F0: v2, F1: 42, F2: 18446744073709551615}
                                                            _ = v60
                                                        } else {
                                                            var v61 int = fi43
                                                            v62 := GenS5{F0: (fi43 / ((32767) | 1)), F1: 592}
                                                            if (0 < 1) {
                                                                switch "" {
                                                                case "":
                                                                    v63 := GenS6{F0: 32767, F1: fi20, F2: true}
                                                                    _ = v63
                                                                case "foo bar":
                                                                    fmt.Println(v2)
                                                                default:
                                                                    if false {
                                                                        if ((v4 | v4) == uint32(v55)) {
                                                                            var v64 uint32 = 2147483647
                                                                            if false {
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = int(v55)
                                                                                }
                                                                            }
                                                                            _ = v64
                                                                        } else {
                                                                            if true {
                                                                                if true {
                                                                                    v55 -= 0
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = int(v55)
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                switch ((v19 + (((("hello" + (v19 + ("\n\t" + v19))) + v19) + "hello") + ("a" + ""))) + (v19 + "")) {
                                                                                case "a":
                                                                                    var v65 int8 = 42
                                                                                    _ = v65
                                                                                }
                                                                            }
                                                                            if len(v1) > 0 {
                                                                                v1[0] = (v4 & v4)
                                                                            }
                                                                            var v66 int8 = 1
                                                                            _ = v66
                                                                        }
                                                                        v67 := GenS1{F0: -1, F1: (v55 + (v55 % (((v55 ^ v55)) | 1))), F2: 65535}
                                                                        fmt.Println(v4, v61, fi43)
                                                                        _ = v67
                                                                    }
                                                                }
                                                                switch ("a" + ("\n\t" + v19)) {
                                                                case "":
                                                                    var v68 uint = 2
                                                                    _ = v68
                                                                case "hello":
                                                                    v69 := GenS1{F0: v2, F1: (v55 ^ (v55 - 382)), F2: 1}
                                                                    _ = v69
                                                                case "foo bar":
                                                                    _ = len(v3)
                                                                }
                                                                v55 = 42
                                                            }
                                                            _ = v61
                                                            _ = v62
                                                        }
                                                    case "a":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    case "hello":
                                                        if false {
                                                            v3 = append(v3, (fi43 * (fi43 ^ (fi43 ^ fi20))))
                                                            v19 = (v19 + v19)
                                                            if (false || (!((!false) && (!false)))) {
                                                                var v70 int16 = -727
                                                                _ = v70
                                                            } else {
                                                                if false {
                                                                    v71 := GenS7{F0: (v55 % ((v55) | 1))}
                                                                    v4 += v4
                                                                    _ = v71
                                                                }
                                                                v19 = "a"
                                                                if (!false) {
                                                                    var v72 uint64 = 18446744073709551615
                                                                    _ = v72
                                                                }
                                                            }
                                                        } else {
                                                            v21 = (v21 - 58)
                                                            var v73 int16 = 1000
                                                            _ = v73
                                                        }
                                                    }
                                                    fmt.Println(v19, v2, fi43)
                                                } else {
                                                    fmt.Println(v55, v21)
                                                    v3 = append(v3, (-(fi20)))
                                                }
                                                fmt.Println(v21, fi43)
                                            }
                                            _ = v54
                                            _ = v55
                                        }
                                        var v74 uint16 = 42
                                        _ = v74
                                    } else {
                                        if (!false) {
                                            v75 := GenS4{F0: 951}
                                            v21 = 42
                                            v19 = "foo bar"
                                            _ = v75
                                        } else {
                                            v76 := GenS2{F0: true}
                                            if false {
                                                v4 -= (v4 & 552)
                                                var v77 uint = genFn0(-128, v21)
                                                _ = v77
                                            }
                                            _ = v76
                                        }
                                        v2 = -885
                                        v19 = v19
                                        var v78 uint32 = (^(v4))
                                        var v79 int64 = v2
                                        v80 := GenS8{F0: v21, F1: 663}
                                        if true {
                                            v81 := GenS9{F0: 100, F1: 166, F2: 1}
                                            _ = v81
                                        }
                                        _ = v78
                                        _ = v79
                                        _ = v80
                                    }
                                default:
                                    if false {
                                        v82 := GenS1{F0: v2, F1: 65535, F2: 42}
                                        if false {
                                            v2 -= -9223372036854775808
                                            v83 := GenS10{F0: 893, F1: 0, F2: v2}
                                            _ = v83
                                        } else {
                                            _ = len(v1)
                                            _ = len(v3)
                                            if len(v1) > 0 {
                                                v1[0] = 2147483647
                                            }
                                        }
                                        _ = v82
                                    }
                                }
                                if (!false) {
                                    if len(v1) > 0 {
                                        v1[0] = (v4 + (v4 & (v4 ^ 2147483647)))
                                    }
                                    v19 = ("" + "")
                                    var v84 int = 32768
                                    _ = v84
                                } else {
                                    _ = len(v3)
                                }
                                v1 = append(v1, v4)
                            }
                            var v85 int8 = -22
                            _ = fi43
                            _ = v85
                        }
                    } else {
                        fmt.Println(v21, v21, fi20)
                        var v86 uint32 = 501
                        _ = v86
                    }
                    if (false || ((255 >= genFn0(1, v21)) || (218 != 42))) {
                        var v87 float64 = (0.1 * 1e10)
                        var v88 int = -9223372036854775808
                        _ = v87
                        _ = v88
                    } else {
                        v21 -= v21
                        if false {
                            switch ("hello" + (("\n\t" + "foo bar") + "foo bar")) {
                            case "":
                                fmt.Println(v21, v21)
                            }
                            for fi89 := 0; fi89 < 17; fi89++ {
                                var v90 uint8 = v21
                                _ = fi89
                                _ = v90
                            }
                            var v91 uint8 = 195
                            _ = v91
                        } else {
                            v19 = (((v19 + v19) + "a") + "\n\t")
                            var v92 int8 = 0
                            _ = v92
                        }
                        v1 = append(v1, (v4 >> ((255) & 63)))
                    }
                    var v93 bool = (!(1e10 != 0.0))
                    _ = v93
                }
            }
            _ = v21
        }
        _ = fi20
    }
    switch 9223372036854775807 {
    case -155:
        fmt.Println(v2)
    case 106:
        v94 := GenS11{F0: 1}
        _ = v94
    default:
        var v95 int8 = -126
        _ = v95
    }
    for fi96 := 0; fi96 < 7; fi96++ {
        v4 = (v4 * v4)
        _ = fi96
    }
    v4 = (v4 % (((^(v4))) | 1))
    v19 = (v19 + v19)
    v1 = append(v1, (v4 & (v4 % (((^(v4))) | 1))))
    for fi97 := 0; fi97 < 20; fi97++ {
        for fi98 := 0; fi98 < 12; fi98++ {
            v19 = (v19 + ((v19 + v19) + v19))
            _ = fi98
        }
        fmt.Println(v19, v2)
        _ = fi97
    }
    v4 = (v4 ^ v4)
    fmt.Println(v2, v4, v19)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
