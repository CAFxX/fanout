package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 string; F2 int }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 uint16; F1 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int8; F1 int8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 bool }
func genFn0() uint {
    switch 42 {
    case -1:
        v5 := GenS1{F0: 518}
        _ = v5
    case -485:
        switch (("hello" + ("foo bar" + ("hello" + "foo bar"))) + ("\n\t" + "\n\t")) {
        case "\n\t":
            if true {
                var v6 string = "a"
                _ = v6
            }
        case "a":
            var v7 int = 128
            _ = v7
        }
    default:
        var v8 uint16 = 239
        _ = v8
    }
    switch -9223372036854775808 {
    case -35:
        var v9 uint16 = 665
        _ = v9
    case 1:
        if true {
            if false {
                var v10 float32 = (0.1 * 1e10)
                _ = v10
            } else {
                var v11 int16 = -1
                v12 := []int{32768, 9223372036854775807, 844, 4294967295}
                _ = v11
                _ = v12
            }
        }
    case 128:
        var v13 int64 = 9223372036854775807
        _ = v13
    default:
        if true {
            v14 := GenS1{F0: 1}
            _ = v14
        } else {
            v15 := []int{-2147483648, -686}
            var v16 string = "foo bar"
            for fi17 := 0; fi17 < 17; fi17++ {
                switch (fi17 + 4294967295) {
                case -777:
                    var v18 bool = false
                    _ = v18
                }
                v19 := GenS1{F0: 4294967295}
                _ = fi17
                _ = v19
            }
            _ = v15
            _ = v16
        }
    }
    v20 := []int{9223372036854775807, 32768}
    for fi21 := 0; fi21 < 16; fi21++ {
        for fi22 := 0; fi22 < 8; fi22++ {
            switch (fi21 * (^(fi21))) {
            case 989:
                var v23 int32 = -2147483648
                _ = v23
            case -9223372036854775808:
                switch (("\n\t" + ("" + "\n\t")) + (("a" + "foo bar") + ("\n\t" + ("\n\t" + "foo bar")))) {
                case "":
                    if (!true) {
                        v24 := GenS1{F0: 1}
                        _ = len(v20)
                        _ = len(v20)
                        v25 := GenS2{F0: -128}
                        if ((fi21 >= fi22) || (1 == 463)) {
                            fmt.Println(fi22, fi21)
                            var v26 int64 = 517
                            _ = v26
                        } else {
                            var v27 uint64 = 0
                            if len(v20) > 0 {
                                v20[0] = fi21
                            }
                            _ = v27
                        }
                        _ = v24
                        _ = v25
                    }
                    var v28 int8 = 99
                    if (true && (0 != 9223372036854775807)) {
                        var v29 float32 = -2.25
                        if (fi22 <= (fi21 ^ (fi22 % (((^(fi21))) | 1)))) {
                            var v30 float64 = 1.5
                            v29 += v29
                            fmt.Println(fi22, fi21)
                            _ = v30
                        }
                        _ = v29
                    } else {
                        var v31 float32 = (float32(((1e10 - 1.5) * 1e10)) / (0.0 - 1e10))
                        switch 32767 {
                        case 4294967295:
                            fmt.Println(v28, fi22, fi21)
                            switch (fi21 ^ 4294967295) {
                            case 2147483647:
                                switch "" {
                                case "\n\t":
                                    fmt.Println(fi21, v28)
                                    v32 := GenS2{F0: (^(v28))}
                                    _ = v32
                                default:
                                    fmt.Println(v31, v28, v28)
                                }
                            case 926:
                                v28 = v28
                            case 4294967295:
                                var v33 uint64 = 4294967295
                                _ = v33
                            default:
                                switch "hello" {
                                case "":
                                    v34 := GenS2{F0: v28}
                                    _ = v34
                                case "a":
                                    switch "\n\t" {
                                    case "hello":
                                        v35 := GenS2{F0: v28}
                                        _ = v35
                                    case "\n\t":
                                        fmt.Println(v31)
                                    default:
                                        v36 := GenS3{F0: 32767, F1: "hello", F2: fi22}
                                        _ = v36
                                    }
                                case "foo bar":
                                    v28 -= v28
                                }
                            }
                        case -204:
                            fmt.Println(v31)
                        default:
                            v31 += v31
                        }
                        switch ("hello" + "hello") {
                        case "":
                            if (false && (31 < 2147483647)) {
                                v31 += v31
                                fmt.Println(fi22, fi21, fi22)
                                switch (fi22 / ((fi22) | 1)) {
                                case -9223372036854775808:
                                    v31 = (float32(1e10) + 1.5)
                                default:
                                    v28 = v28
                                }
                                var v37 float64 = (3.14159 * (1.5 - (1.5 - 0.1)))
                                if len(v20) > 0 {
                                    v20[0] = fi21
                                }
                                if ((0.0 == v37) && false) {
                                    var v38 int8 = (v28 | v28)
                                    v39 := GenS1{F0: 1}
                                    _ = v38
                                    _ = v39
                                } else {
                                    if false {
                                        v28 += v28
                                    } else {
                                        if len(v20) >= 2 {
                                            v20 = v20[1:len(v20)]
                                        }
                                        v40 := GenS2{F0: (^(v28))}
                                        if true {
                                            v20 = append(v20, 127)
                                            v41 := GenS3{F0: 65535, F1: "\n\t", F2: fi21}
                                            v20 = append(v20, fi21)
                                            _ = v41
                                        }
                                        _ = v40
                                    }
                                }
                                var v42 int8 = (v28 / ((v28) | 1))
                                _ = v37
                                _ = v42
                            }
                        default:
                            if (!((("" + "\n\t") + "a") > "a")) {
                                fmt.Println(v31, v28)
                                v31 += v31
                                v28 -= (v28 & (v28 | v28))
                            } else {
                                _ = len(v20)
                                if (1000 != 32767) {
                                    fmt.Println(v31, fi22)
                                    var v43 uint8 = 0
                                    _ = v43
                                } else {
                                    switch (fi22 / ((fi21) | 1)) {
                                    case -234:
                                        var v44 int64 = 0
                                        _ = v44
                                    case 42:
                                        fmt.Println(fi22, v31, v31)
                                    default:
                                        v28 = (v28 - (^(v28)))
                                    }
                                    if true {
                                        var v45 uint8 = 127
                                        _ = len(v20)
                                        _ = v45
                                    }
                                    v28 += v28
                                }
                                if len(v20) > 0 {
                                    v20[0] = fi21
                                }
                            }
                        }
                        _ = v31
                    }
                    _ = v28
                case "\n\t":
                    fmt.Println(fi22, fi21)
                default:
                    v46 := GenS4{F0: false, F1: (true && false)}
                    _ = v46
                }
            case 128:
                v47 := GenS5{F0: 32767, F1: -350}
                _ = v47
            }
            _ = fi22
        }
        v20 = append(v20, fi21)
        for fi48 := 0; fi48 < 2; fi48++ {
            v49 := GenS6{F0: 1}
            if (9223372036854775807 < 1) {
                switch "\n\t" {
                case "\n\t":
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(fi21, fi21)
                    case "\n\t":
                        var v50 uint8 = 2
                        _ = v50
                    case "hello":
                        var v51 uint32 = 2
                        _ = v51
                    }
                case "a":
                    if false {
                        if false {
                            var v52 int8 = -1
                            _ = v52
                        } else {
                            fmt.Println(fi21, fi48)
                            switch "a" {
                            case "foo bar":
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                            case "\n\t":
                                var v53 bool = true
                                _ = v53
                            }
                        }
                    } else {
                        var v54 string = ("foo bar" + "a")
                        var v55 float64 = 0.0
                        _ = v54
                        _ = v55
                    }
                case "foo bar":
                    var v56 int32 = 1
                    _ = v56
                default:
                    if false {
                        if true {
                            var v57 bool = true
                            fmt.Println(fi48, v57, v57)
                            v57 = (-1 < 0)
                            _ = v57
                        }
                        var v58 int8 = -128
                        v58 = int8(fi21)
                        _ = v58
                    } else {
                        var v59 float32 = 3.14159
                        switch (fi48 + fi21) {
                        case 1:
                            if (0.0 >= 0.0) {
                                if (2147483647 != 814) {
                                    var v60 int8 = 0
                                    switch "foo bar" {
                                    case "hello":
                                        switch "foo bar" {
                                        case "":
                                            v59 -= (1e10 - float32(3.14159))
                                        case "a":
                                            var v61 int = (-(fi21))
                                            _ = v61
                                        case "\n\t":
                                            fmt.Println(v59, v59, v60)
                                        }
                                    }
                                    var v62 uint16 = 0
                                    _ = v60
                                    _ = v62
                                } else {
                                    fmt.Println(fi48)
                                }
                                v59 = v59
                            }
                        default:
                            switch "a" {
                            case "":
                                v59 -= 0.1
                            default:
                                var v63 uint16 = 853
                                _ = v63
                            }
                        }
                        _ = v59
                    }
                }
                v64 := GenS7{F0: -4, F1: -1}
                _ = v64
            } else {
                v65 := GenS7{F0: 100, F1: -128}
                _ = v65
            }
            _ = fi48
            _ = v49
        }
        _ = fi21
    }
    _ = v20
    return 42
}

func main() {
    var v1 float64 = (1.5 / 1.0)
    v2 := []int{}
    var v3 string = "\n\t"
    var v4 int64 = 723
    v66 := GenS8{F0: 127}
    v1 -= 3.14159
    v67 := GenS8{F0: 2}
    v68 := GenS8{F0: 1}
    v3 = "foo bar"
    v3 = "hello"
    for fi69 := 0; fi69 < 9; fi69++ {
        var v70 int16 = 1
        fmt.Println(v70, fi69, fi69)
        _ = fi69
        _ = v70
    }
    for fi71 := 0; fi71 < 9; fi71++ {
        v72 := GenS9{F0: ((-2147483648 >= -2147483648) && (1 > 42))}
        v1 = v1
        _ = fi71
        _ = v72
    }
    v4 = v4
    for fi73 := 0; fi73 < 2; fi73++ {
        for fi74 := 0; fi74 < 14; fi74++ {
            var v75 int32 = 1
            _ = fi74
            _ = v75
        }
        _ = fi73
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v66
    _ = v67
    _ = v68
}
