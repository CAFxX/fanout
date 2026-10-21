package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint }
type GenS4 struct { F0 int; F1 int16; F2 uint16 }
type GenS5 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int32; F1 bool; F2 uint64 }
func genFn0(p0 int64, p1 int64, p2 int8) uint32 {
    var v5 uint64 = 765
    switch ("foo bar" + "hello") {
    case "hello":
        fmt.Println(p0)
    case "a":
        switch 32768 {
        case 0:
            switch "\n\t" {
            case "\n\t":
                for fi6 := 0; fi6 < 6; fi6++ {
                    if true {
                        v7 := GenS1{F0: (v5 % ((v5) | 1)), F1: v5}
                        p1 = (p0 + 9223372036854775807)
                        p1 = p0
                        _ = v7
                    }
                    var v8 float32 = (1e10 - 3.14159)
                    switch (("" + "a") + "") {
                    case "foo bar":
                        p2 = p2
                    case "\n\t":
                        p1 = p1
                    case "":
                        if false {
                            v9 := []int{(fi6 ^ fi6), 349, (fi6 % (((fi6 / (((^(fi6))) | 1))) | 1)), fi6}
                            fmt.Println(fi6)
                            _ = v9
                        }
                    default:
                        p2 = p2
                    }
                    _ = fi6
                    _ = v8
                }
            case "a":
                fmt.Println(v5, v5, v5)
            case "":
                switch 127 {
                case 9223372036854775807:
                    if (0.1 != 0.1) {
                        if true {
                            v10 := []int{-782, 4294967295, 400, 128}
                            _ = v10
                        } else {
                            var v11 uint = 4294967295
                            _ = v11
                        }
                        if false {
                            if ((p1 + (p0 & p0)) > (p0 & (^(p1)))) {
                                p2 += -128
                                var v12 uint = 542
                                _ = v12
                            } else {
                                v5 = (v5 >> ((0) & 63))
                                if false {
                                    for fi13 := 0; fi13 < 13; fi13++ {
                                        for fi14 := 0; fi14 < 20; fi14++ {
                                            switch 0 {
                                            case 127:
                                                fmt.Println(p2, p2)
                                            default:
                                                fmt.Println(fi14)
                                            }
                                            fmt.Println(fi14, fi13, v5)
                                            v15 := []int{}
                                            v5 += (v5 / ((v5) | 1))
                                            _ = fi14
                                            _ = v15
                                        }
                                        _ = fi13
                                    }
                                    v16 := GenS2{F0: v5, F1: p1}
                                    _ = v16
                                } else {
                                    p0 = (^(p0))
                                    var v17 string = ("\n\t" + "hello")
                                    v18 := []int{128, 0, 127, -2147483648}
                                    _ = v17
                                    _ = v18
                                }
                                p1 = p0
                            }
                            v19 := GenS1{F0: 4294967295, F1: v5}
                            var v20 int64 = (p0 - (p0 - (^(p0))))
                            _ = v19
                            _ = v20
                        } else {
                            var v21 bool = true
                            v22 := []int{0, 4294967295, 2147483647, 42}
                            _ = v21
                            _ = v22
                        }
                    }
                }
            }
        default:
            v23 := []int{-2147483648}
            _ = v23
        }
    case "foo bar":
        v24 := GenS3{F0: 1, F1: 2147483648, F2: 2}
        _ = v24
    default:
        if true {
            v25 := GenS1{F0: v5, F1: 0}
            _ = v25
        } else {
            if true {
                var v26 int16 = 0
                if ((p1 + (p0 & 820)) != 0) {
                    fmt.Println(v26, p0, v26)
                    fmt.Println(p2, v26, v5)
                    var v27 float64 = (3.14159 - (1e10 / 1.0))
                    _ = v27
                }
                _ = v26
            } else {
                v28 := GenS4{F0: 2147483647, F1: 0, F2: 42}
                fmt.Println(p0)
                _ = v28
            }
        }
    }
    for fi29 := 0; fi29 < 20; fi29++ {
        switch fi29 {
        case 127:
            v30 := []int{-922}
            _ = v30
        case 832:
            switch "" {
            case "hello":
                for fi31 := 0; fi31 < 15; fi31++ {
                    var v32 int16 = 549
                    if (!true) {
                        fmt.Println(v5, v32)
                    }
                    v33 := GenS3{F0: 2, F1: 2147483648, F2: 18446744073709551615}
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            default:
                if (!(4294967295 != 2)) {
                    fmt.Println(p2)
                    p0 = p1
                }
            }
        }
        v34 := GenS5{F0: 127, F1: v5, F2: (p2 + -1)}
        var v35 uint64 = 4294967295
        _ = fi29
        _ = v34
        _ = v35
    }
    v36 := GenS5{F0: 188, F1: 0, F2: p2}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v36
    _ = v5
    return 2
}

func main() {
    var v1 string = ("a" + "hello")
    v2 := make(map[int]bool)
    var v3 uint32 = 2
    var v4 int16 = -1
    v3 -= (v3 ^ (v3 % (((v3 % ((0) | 1))) | 1)))
    if true {
        var v37 int = 42
        _ = v37
    }
    v38 := GenS6{F0: 32767}
    var v39 uint16 = 315
    var v40 uint32 = (^(v3))
    v4 = v4
    if (65535 != 4294967295) {
        switch -1 {
        case -504:
            v39 = 835
        }
    }
    v40 += genFn0(-305, 42, 100)
    var v41 float32 = (1.5 + (3.14159 / 0.1))
    v2[127] = true
    if (2 != 32767) {
        fmt.Println(v40)
        if true {
            if true {
                if false {
                    if false {
                        var v42 int64 = -513
                        v43 := []int{42, 486}
                        if true {
                            if len(v43) >= 2 {
                                v43 = v43[1:len(v43)]
                            }
                        }
                        _ = v42
                        _ = v43
                    } else {
                        v4 += (v4 + 1000)
                        switch "foo bar" {
                        case "a":
                            if true {
                                if (985 >= 18446744073709551615) {
                                    v1 = ("foo bar" + "hello")
                                } else {
                                    v40 -= 65535
                                    if false {
                                        fmt.Println(v39, v3)
                                        _ = len(v2)
                                        var v44 uint16 = (v39 * (v39 | (v39 % ((v39) | 1))))
                                        _ = v44
                                    }
                                }
                                v40 = (v40 << ((1) & 63))
                            } else {
                                delete(v2, 368)
                                v39 = (^(v39))
                            }
                        case "hello":
                            v40 += (v3 >> ((2) & 63))
                        case "\n\t":
                            if (0 == v39) {
                                fmt.Println(v41)
                            }
                        }
                        switch 128 {
                        case 9223372036854775807:
                            switch (v1 + v1) {
                            case "hello":
                                v3 = (v3 - (v3 & 0))
                            case "":
                                if (v39 == (v39 & (v39 * (^(v39))))) {
                                    for fi45 := 0; fi45 < 5; fi45++ {
                                        v46 := GenS1{F0: 42, F1: 0}
                                        if (2147483647 <= 862) {
                                            delete(v2, (fi45 << ((471) & 63)))
                                            _ = len(v2)
                                            fmt.Println(v4, fi45)
                                        } else {
                                            if true {
                                                var v47 int8 = -74
                                                if (128 == 1) {
                                                    for fi48 := 0; fi48 < 11; fi48++ {
                                                        if false {
                                                            var v49 uint64 = 1
                                                            v40 = (v40 << ((276) & 63))
                                                            var v50 float64 = (1e10 / 0.1)
                                                            _ = v49
                                                            _ = v50
                                                        } else {
                                                            fmt.Println(fi45)
                                                            v2[(fi45 >> ((0) & 63))] = ((v47 % (((v47 ^ (v47 >> ((4294967295) & 63)))) | 1)) <= v47)
                                                        }
                                                        v41 = (v41 / 1.5)
                                                        _ = fi48
                                                    }
                                                } else {
                                                    v51 := []int{128, int(v40), 4294967295}
                                                    _ = v51
                                                }
                                                var v52 uint = 559
                                                _ = v47
                                                _ = v52
                                            }
                                            var v53 string = v1
                                            for fi54 := 0; fi54 < 20; fi54++ {
                                                switch (v1 + "foo bar") {
                                                case "a":
                                                    v55 := GenS2{F0: 0, F1: 0}
                                                    if false {
                                                        if (!(100 > -70)) {
                                                            var v56 int64 = 1
                                                            fmt.Println(v53, v4)
                                                            v57 := []int{(fi45 | fi45), 1}
                                                            var v58 int16 = v4
                                                            v2[fi54] = true
                                                            _ = v56
                                                            _ = v57
                                                            _ = v58
                                                        }
                                                        fmt.Println(v4, v39, v1)
                                                        if false {
                                                            fmt.Println(v41, v39, fi54)
                                                            v59 := GenS7{F0: 127}
                                                            var v60 int8 = 127
                                                            switch int(v39) {
                                                            case -9223372036854775808:
                                                                v61 := []int{fi54, fi45, (fi45 | (fi54 * -2147483648))}
                                                                _ = v61
                                                            default:
                                                                v1 = v53
                                                            }
                                                            switch (fi45 >> ((42) & 63)) {
                                                            case -163:
                                                                v4 = 550
                                                            default:
                                                                switch "hello" {
                                                                case "hello":
                                                                    v53 = v53
                                                                }
                                                            }
                                                            _ = v59
                                                            _ = v60
                                                        } else {
                                                            v62 := GenS3{F0: 127, F1: v40, F2: 122}
                                                            v63 := []int{(fi45 >> ((0) & 63)), 4294967295}
                                                            _ = v62
                                                            _ = v63
                                                        }
                                                    }
                                                    if false {
                                                        var v64 int16 = v4
                                                        if false {
                                                            var v65 float64 = 1e10
                                                            _ = v65
                                                        }
                                                        v53 = (v53 + (v1 + ("hello" + v1)))
                                                        if false {
                                                            if ((true || true) && false) {
                                                                var v66 uint64 = 690
                                                                switch fi45 {
                                                                case 381:
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        v67 := GenS8{F0: 7}
                                                                        fmt.Println(v66)
                                                                        _ = v67
                                                                    }
                                                                case -9223372036854775808:
                                                                    var v68 uint64 = v66
                                                                    _ = v68
                                                                case 577:
                                                                    v66 += (v66 << ((65535) & 63))
                                                                }
                                                                _ = v66
                                                            }
                                                        } else {
                                                            v69 := []int{(-(fi54)), fi54}
                                                            fmt.Println(v40, v3)
                                                            v41 += (-2.25 + v41)
                                                            _ = v69
                                                        }
                                                        _ = v64
                                                    }
                                                    _ = v55
                                                case "\n\t":
                                                    var v70 uint = 1
                                                    _ = v70
                                                case "foo bar":
                                                    var v71 int16 = (v4 | v4)
                                                    _ = v71
                                                }
                                                v4 = v4
                                                if (689 != 4294967295) {
                                                    v72 := []int{fi54, 32767, fi54}
                                                    delete(v2, 4294967295)
                                                    _ = v72
                                                } else {
                                                    v73 := GenS8{F0: 4294967295}
                                                    v74 := GenS9{F0: 127}
                                                    _ = v73
                                                    _ = v74
                                                }
                                                _ = fi54
                                            }
                                            _ = v53
                                        }
                                        _ = fi45
                                        _ = v46
                                    }
                                    if (false && (-9223372036854775808 <= 1)) {
                                        v75 := GenS10{F0: 0}
                                        _ = v75
                                    }
                                    var v76 uint8 = 0
                                    _ = v76
                                } else {
                                    var v77 uint64 = 4294967295
                                    _ = v77
                                }
                            case "\n\t":
                                if ((float32(((0.1 / 1.5) + (((1.5 / (3.14159 - (1.5 * (float64(v41) / (1.5 - ((1.5 * (3.14159 + 0.0)) / 0.1)))))) / 0.1) + 0.0))) - v41) > v41) {
                                    for fi78 := 0; fi78 < 5; fi78++ {
                                        v79 := GenS10{F0: 127}
                                        _ = fi78
                                        _ = v79
                                    }
                                } else {
                                    if true {
                                        switch "hello" {
                                        case "\n\t":
                                            switch 2147483647 {
                                            case -9223372036854775808:
                                                switch v1 {
                                                case "a":
                                                    fmt.Println(v40, v3, v41)
                                                case "":
                                                    for fi80 := 0; fi80 < 14; fi80++ {
                                                        for fi81 := 0; fi81 < 20; fi81++ {
                                                            var v82 uint8 = 255
                                                            if ((v41 + (float32(1e10) + (v41 - ((v41 / (3.14159 + float32(0.1))) + 1e10)))) == v41) {
                                                                var v83 bool = (true && true)
                                                                _ = len(v2)
                                                                var v84 bool = (!v83)
                                                                _ = v83
                                                                _ = v84
                                                            }
                                                            _ = fi81
                                                            _ = v82
                                                        }
                                                        _ = fi80
                                                    }
                                                }
                                            case 569:
                                                v85 := GenS11{F0: 305, F1: false, F2: 1}
                                                _ = v85
                                            case -465:
                                                v2[126] = (1 < 127)
                                            }
                                        default:
                                            v3 -= (^(v40))
                                        }
                                        var v86 int8 = 0
                                        _ = v86
                                    }
                                }
                            default:
                                v87 := []int{-502, 9223372036854775807, 127}
                                _ = v87
                            }
                        }
                    }
                    v4 = v4
                }
                for fi88 := 0; fi88 < 16; fi88++ {
                    switch -458 {
                    case -9223372036854775808:
                        v41 = 1e10
                    default:
                        switch int(v4) {
                        case 1:
                            v40 = (v40 ^ v3)
                        default:
                            if true {
                                v40 -= (v40 ^ 0)
                                for fi89 := 0; fi89 < 15; fi89++ {
                                    var v90 int = 9223372036854775807
                                    var v91 uint8 = 127
                                    var v92 float64 = (0.1 * (((3.14159 * 0.1) * float64(-2.25)) * 1.5))
                                    _ = fi89
                                    _ = v90
                                    _ = v91
                                    _ = v92
                                }
                                v93 := GenS5{F0: 255, F1: 2, F2: -1}
                                _ = v93
                            } else {
                                var v94 uint32 = v40
                                _ = v94
                            }
                        }
                    }
                    _ = fi88
                }
            } else {
                var v95 int64 = -1
                _ = v95
            }
        }
    } else {
        v39 += v39
    }
    var v96 uint8 = 128
    v39 -= v39
    fmt.Println(v40, v3)
    _ = len(v2)
    fmt.Println(v39, v4, v4)
    var v97 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v96
    _ = v97
}
