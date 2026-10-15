package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 int64; F1 uint64; F2 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 int64; F2 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int32; F1 int16 }
func genFn0(p0 int) uint64 {
    var v5 int16 = -203
    switch (p0 ^ 467) {
    case -9223372036854775808:
        if (true || false) {
            switch "hello" {
            case "\n\t":
                switch (p0 >> ((110) & 63)) {
                case -1:
                    switch int(v5) {
                    case 127:
                        switch "a" {
                        case "foo bar":
                            var v6 int32 = 42
                            _ = v6
                        case "\n\t":
                            v5 -= (v5 | -32768)
                        case "":
                            v7 := []int{1}
                            _ = v7
                        default:
                            var v8 float64 = 3.14159
                            _ = v8
                        }
                    case -42:
                        p0 = (p0 | 32768)
                    default:
                        p0 -= (p0 + (^(p0)))
                    }
                case 32767:
                    for fi9 := 0; fi9 < 2; fi9++ {
                        v10 := []int{42, fi9, fi9, fi9}
                        _ = len(v10)
                        _ = fi9
                        _ = v10
                    }
                case -582:
                    p0 = p0
                }
            case "hello":
                v11 := []int{p0, -2147483648}
                _ = v11
            default:
                if false {
                    if (!(182 < 0)) {
                        switch ("" + "hello") {
                        case "\n\t":
                            p0 = 487
                        case "":
                            if false {
                                for fi12 := 0; fi12 < 15; fi12++ {
                                    var v13 int8 = 28
                                    var v14 int32 = 496
                                    v13 = (v13 / ((v13) | 1))
                                    _ = fi12
                                    _ = v13
                                    _ = v14
                                }
                                fmt.Println(p0, p0)
                            } else {
                                var v15 uint32 = 1
                                _ = v15
                            }
                        case "foo bar":
                            fmt.Println(v5)
                        }
                        switch "" {
                        case "foo bar":
                            fmt.Println(v5, v5)
                        default:
                            var v16 uint64 = 18446744073709551615
                            _ = v16
                        }
                    } else {
                        p0 += p0
                        for fi17 := 0; fi17 < 9; fi17++ {
                            v5 = (v5 / ((1000) | 1))
                            v18 := []int{}
                            p0 = (^(fi17))
                            _ = fi17
                            _ = v18
                        }
                        p0 += p0
                    }
                    v19 := []int{p0, (p0 | p0)}
                    _ = v19
                } else {
                    v5 += (v5 | (v5 * v5))
                }
            }
            v5 = (v5 % ((-1) | 1))
            v20 := GenS1{F0: 1}
            _ = v20
        } else {
            for fi21 := 0; fi21 < 10; fi21++ {
                v22 := GenS2{F0: -128, F1: -44}
                _ = fi21
                _ = v22
            }
            v23 := GenS1{F0: 42}
            _ = v23
        }
    default:
        for fi24 := 0; fi24 < 14; fi24++ {
            if true {
                switch "" {
                case "foo bar":
                    if (-128 == -128) {
                        v25 := GenS2{F0: -15, F1: 20}
                        fmt.Println(v5)
                        _ = v25
                    }
                case "":
                    v5 += (v5 ^ 1)
                case "\n\t":
                    var v26 string = "\n\t"
                    _ = v26
                default:
                    v27 := []int{}
                    _ = v27
                }
                p0 -= int(v5)
            }
            var v28 float32 = 1.5
            _ = fi24
            _ = v28
        }
    }
    _ = p0
    _ = v5
    return 4
}
func genFn1(p0 int32, p1 int8, p2 uint) uint8 {
    v29 := GenS1{F0: -128}
    for fi30 := 0; fi30 < 15; fi30++ {
        v31 := GenS3{F0: int64(p0), F1: 18, F2: -607}
        v32 := []int{(fi30 ^ (fi30 / ((fi30) | 1))), fi30, (fi30 - (fi30 ^ fi30))}
        fmt.Println(fi30, fi30)
        _ = fi30
        _ = v31
        _ = v32
    }
    v33 := GenS3{F0: -831, F1: 1, F2: 1}
    if true {
        fmt.Println(p0, p0, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    _ = v33
    return 128
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    var v3 int = 159
    var v4 uint8 = 190
    v34 := GenS3{F0: -9223372036854775808, F1: genFn0(32767), F2: 81}
    fmt.Println(v3)
    for fi35 := 0; fi35 < 4; fi35++ {
        v36 := GenS3{F0: 0, F1: 42, F2: 431}
        switch (-(fi35)) {
        case 32768:
            var v37 int = (-(v3))
            _ = v37
        default:
            fmt.Println(v3, v4, fi35)
        }
        _ = fi35
        _ = v36
    }
    if (int(v4) < 9223372036854775807) {
        v38 := v1[(("a" + ("hello" + "hello")) + (((((("a" + (("a" + ("\n\t" + ("a" + "foo bar"))) + "a")) + "") + ((("hello" + ("a" + "foo bar")) + "hello") + "foo bar")) + "") + ("hello" + (("a" + "") + ("hello" + "")))) + ("a" + (("a" + ("hello" + "hello")) + "a"))))]
        v39 := []int{(v3 + 2147483647), (^(v3)), (^(v3))}
        _ = v38
        _ = v39
    } else {
        for fi40 := 0; fi40 < 4; fi40++ {
            v4 = genFn1(-429, -128, 18446744073709551615)
            if (888 <= 1) {
                switch "a" {
                case "a":
                    if (!false) {
                        if (874 == 42) {
                            v4 = 255
                            v41 := GenS2{F0: 60, F1: -124}
                            _ = v41
                        } else {
                            for fi42 := 0; fi42 < 2; fi42++ {
                                var v43 int = fi40
                                v3 += 532
                                v44 := GenS2{F0: -128, F1: int8(fi42)}
                                _ = fi42
                                _ = v43
                                _ = v44
                            }
                            if false {
                                fmt.Println(v4)
                                var v45 uint32 = 0
                                _ = v45
                            }
                        }
                        if true {
                            v46 := GenS4{F0: int64(v3)}
                            for fi47 := 0; fi47 < 5; fi47++ {
                                v4 = v4
                                _ = fi47
                            }
                            _ = v46
                        }
                        v48 := []int{42}
                        _ = v48
                    }
                case "":
                    v49 := []int{fi40, (v3 * -2147483648)}
                    _ = v49
                case "foo bar":
                    v3 = 253
                default:
                    switch v3 {
                    case -961:
                        var v50 bool = false
                        _ = v50
                    case 913:
                        if false {
                            v3 += (v3 * (^(v3)))
                        } else {
                            v51 := GenS5{F0: 100, F1: 666, F2: 2147483648}
                            _ = v51
                        }
                    case 128:
                        v1["a"] = "hello"
                    default:
                        if false {
                            var v52 uint8 = 128
                            for fi53 := 0; fi53 < 12; fi53++ {
                                switch v3 {
                                case 9223372036854775807:
                                    var v54 int64 = 1
                                    _ = v54
                                case 128:
                                    v55 := GenS1{F0: -128}
                                    _ = v55
                                default:
                                    var v56 uint32 = 4294967295
                                    _ = v56
                                }
                                var v57 uint16 = 971
                                v3 = v3
                                _ = fi53
                                _ = v57
                            }
                            v58 := []int{fi40, (-(fi40)), (^(v3))}
                            _ = v52
                            _ = v58
                        } else {
                            v59 := GenS1{F0: 49}
                            var v60 string = ("\n\t" + "")
                            for fi61 := 0; fi61 < 18; fi61++ {
                                v1[""] = v60
                                switch (fi40 * 877) {
                                case 0:
                                    fmt.Println(fi61, v3)
                                }
                                _ = fi61
                            }
                            _ = v59
                            _ = v60
                        }
                    }
                }
                var v62 float32 = 1e10
                var v63 uint64 = 18446744073709551615
                _ = v62
                _ = v63
            } else {
                if (127 <= -1) {
                    var v64 uint16 = 546
                    v65 := GenS3{F0: -1, F1: 0, F2: -749}
                    var v66 string = "foo bar"
                    _ = v64
                    _ = v65
                    _ = v66
                }
                v67 := []int{9223372036854775807, fi40, (v3 % ((fi40) | 1)), 2147483647}
                _ = v67
            }
            _ = fi40
        }
        v68 := GenS5{F0: -1, F1: -9223372036854775808, F2: 4294967295}
        _ = v68
    }
    switch v3 {
    case 42:
        switch 32767 {
        case 127:
            if false {
                if (uint32(v4) != 1) {
                    for fi69 := 0; fi69 < 11; fi69++ {
                        var v70 int64 = -9223372036854775808
                        v71 := GenS2{F0: 0, F1: 42}
                        _ = fi69
                        _ = v70
                        _ = v71
                    }
                    var v72 uint8 = 1
                    if (42 <= v4) {
                        var v73 int32 = 279
                        var v74 uint32 = 883
                        var v75 uint32 = (v74 & (^(v74)))
                        _ = v73
                        _ = v74
                        _ = v75
                    }
                    _ = v72
                }
            }
        case -941:
            fmt.Println(v4, v3)
        default:
            v4 += (v4 / ((0) | 1))
        }
    case 4294967295:
        if (!(65535 < 784)) {
            for fi76 := 0; fi76 < 14; fi76++ {
                var v77 string = "foo bar"
                for fi78 := 0; fi78 < 20; fi78++ {
                    v77 = (v77 + v77)
                    v79 := GenS6{F0: 0}
                    switch ("\n\t" + v77) {
                    case "":
                        v1[""] = v77
                    case "foo bar":
                        v80 := GenS5{F0: -23, F1: -9223372036854775808, F2: 752}
                        _ = v80
                    case "\n\t":
                        fmt.Println(fi78, v4)
                    default:
                        var v81 int32 = -1
                        _ = v81
                    }
                    _ = fi78
                    _ = v79
                }
                _ = fi76
                _ = v77
            }
        }
    case -9223372036854775808:
        v82 := GenS4{F0: 0}
        _ = v82
    default:
        v83 := v1[("" + "a")]
        _ = v83
    }
    v84 := GenS1{F0: -93}
    var v85 uint16 = 2
    v4 = v4
    v3 += v3
    for fi86 := 0; fi86 < 16; fi86++ {
        if (0 == -1) {
            if true {
                if (!false) {
                    if false {
                        var v87 uint8 = 42
                        v88 := v1[("" + "foo bar")]
                        _ = v87
                        _ = v88
                    }
                    v4 = (v4 | v4)
                }
            }
        } else {
            switch "foo bar" {
            case "\n\t":
                v3 -= fi86
            case "":
                if false {
                    for fi89 := 0; fi89 < 7; fi89++ {
                        var v90 float32 = 1e10
                        var v91 int16 = -32768
                        if true {
                            if (v85 <= (^(v85))) {
                                if true {
                                    v90 = v90
                                    v92 := GenS5{F0: -94, F1: 42, F2: 42}
                                    _ = v92
                                } else {
                                    v93 := GenS6{F0: 1}
                                    v94 := GenS7{F0: (v85 % ((367) | 1))}
                                    v95 := []int{fi86, 0, fi89}
                                    _ = v93
                                    _ = v94
                                    _ = v95
                                }
                                fmt.Println(v90, v85, v90)
                            }
                            if (!(("" + "hello") <= "\n\t")) {
                                switch (("a" + "foo bar") + ("a" + "hello")) {
                                case "\n\t":
                                    v96 := GenS6{F0: 42}
                                    _ = v96
                                case "foo bar":
                                    var v97 uint = uint(v85)
                                    _ = v97
                                case "":
                                    v98 := GenS8{F0: 42, F1: 0}
                                    _ = v98
                                default:
                                    v3 = fi86
                                }
                                v99 := v1[("foo bar" + "")]
                                _ = v99
                            }
                            switch ("" + ("foo bar" + "foo bar")) {
                            case "foo bar":
                                _ = len(v1)
                            case "":
                                fmt.Println(fi89, v3)
                            case "\n\t":
                                var v100 uint64 = 159
                                _ = v100
                            }
                        }
                        _ = fi89
                        _ = v90
                        _ = v91
                    }
                    for fi101 := 0; fi101 < 11; fi101++ {
                        switch "" {
                        case "foo bar":
                            v4 = 255
                        case "":
                            var v102 uint = 65535
                            _ = v102
                        case "\n\t":
                            switch "hello" {
                            case "a":
                                var v103 uint32 = 0
                                _ = v103
                            }
                        }
                        _ = fi101
                    }
                    v104 := []int{(fi86 >> ((942) & 63)), v3, 60}
                    _ = v104
                }
            case "foo bar":
                fmt.Println(fi86)
            }
        }
        _ = fi86
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v84
    _ = v85
}
