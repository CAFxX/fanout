package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 string; F2 uint32 }
type GenS4 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS5 struct { F0 int64; F1 uint; F2 uint8 }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 uint; F1 string }
type GenS9 struct { F0 int; F1 int32; F2 uint }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int32; F1 bool; F2 bool }
type GenS12 struct { F0 int8; F1 bool }
type GenS13 struct { F0 int16; F1 int64 }
func genFn0() int16 {
    var v6 uint32 = 4294967295
    switch 42 {
    case -1:
        v7 := GenS1{F0: 0, F1: -440, F2: 2147483648}
        _ = v7
    case 2147483647:
        v6 = 355
    case -770:
        if true {
            for fi8 := 0; fi8 < 16; fi8++ {
                if ((157 >= 128) && (!true)) {
                    v6 -= v6
                }
                _ = fi8
            }
        }
    }
    var v9 int8 = -91
    var v10 int16 = -110
    fmt.Println(v6, v10, v6)
    _ = v10
    _ = v6
    _ = v9
    return int16(v9)
}

func main() {
    var v1 uint8 = 42
    var v2 float32 = (3.14159 / (1e10 + 0.1))
    var v3 uint64 = 4294967295
    v4 := []string{"hello"}
    v5 := make(map[uint32]uint32)
    v11 := GenS2{F0: -1}
    for fi12 := 0; fi12 < 6; fi12++ {
        fmt.Println(v3, v1)
        if true {
            v13 := GenS2{F0: -1}
            var v14 int = 32767
            _ = len(v4)
            _ = v13
            _ = v14
        } else {
            var v15 uint8 = v1
            v15 -= v1
            switch 1 {
            case 0:
                var v16 uint64 = v3
                _ = v16
            case -499:
                v2 -= 3.14159
            default:
                var v17 string = (("foo bar" + "a") + "hello")
                _ = v17
            }
            _ = v15
        }
        _ = fi12
    }
    v2 = v2
    var v18 int8 = -1
    if false {
        if (false || (!true)) {
            v19 := GenS1{F0: (v18 << ((1) & 63)), F1: 9223372036854775807, F2: 4294967295}
            if (!false) {
                for fi20 := 0; fi20 < 10; fi20++ {
                    for fi21 := 0; fi21 < 20; fi21++ {
                        fmt.Println(fi20)
                        delete(v5, 0)
                        switch "a" {
                        case "hello":
                            var v22 bool = false
                            _ = v22
                        }
                        v4 = append(v4, "")
                        v18 = v18
                        _ = fi21
                    }
                    _ = fi20
                }
                v2 = v2
            }
            _ = v19
        } else {
            var v23 int64 = 1
            fmt.Println(v3, v2)
            for fi24 := 0; fi24 < 17; fi24++ {
                var v25 int16 = -32768
                _ = fi24
                _ = v25
            }
            _ = v23
        }
        for fi26 := 0; fi26 < 2; fi26++ {
            for fi27 := 0; fi27 < 9; fi27++ {
                fmt.Println(fi27, fi26)
                v28 := GenS2{F0: 0}
                fmt.Println(fi26, v18, fi27)
                _ = fi27
                _ = v28
            }
            if false {
                v29 := GenS1{F0: v18, F1: fi26, F2: 1}
                v1 = 100
                _ = v29
            } else {
                switch "\n\t" {
                case "a":
                    _ = len(v4)
                case "hello":
                    v1 = v1
                case "foo bar":
                    if (true || false) {
                        for fi30 := 0; fi30 < 11; fi30++ {
                            v3 = (v3 | v3)
                            _ = fi30
                        }
                        v31 := GenS1{F0: (v18 << ((2) & 63)), F1: (fi26 / (((fi26 * (fi26 ^ fi26))) | 1)), F2: 42}
                        switch "a" {
                        case "\n\t":
                            v1 = v1
                        case "hello":
                            if (!false) {
                                var v32 int16 = 1000
                                fmt.Println(v32, v2, v3)
                                switch "\n\t" {
                                case "\n\t":
                                    var v33 int8 = (v18 & (v18 % (((^(v18))) | 1)))
                                    _ = v33
                                default:
                                    v18 += 49
                                }
                                _ = v32
                            }
                        case "a":
                            if (true || false) {
                                if true {
                                    fmt.Println(v3, fi26)
                                } else {
                                    v2 -= v2
                                }
                                var v34 uint16 = 436
                                v35 := GenS3{F0: 32767, F1: "\n\t", F2: 4294967295}
                                _ = v34
                                _ = v35
                            } else {
                                for fi36 := 0; fi36 < 14; fi36++ {
                                    if true {
                                        _ = len(v5)
                                    } else {
                                        if false {
                                            v1 += 128
                                        }
                                    }
                                    v2 += (v2 - float32(0.0))
                                    v37 := GenS3{F0: -597, F1: "\n\t", F2: 166}
                                    switch "\n\t" {
                                    case "":
                                        fmt.Println(fi26, fi26)
                                    case "foo bar":
                                        delete(v5, 2147483648)
                                    case "\n\t":
                                        var v38 int32 = 0
                                        _ = v38
                                    }
                                    _ = fi36
                                    _ = v37
                                }
                            }
                        }
                        _ = v31
                    }
                }
            }
            _ = fi26
        }
        switch 2147483647 {
        case 127:
            v3 += (v3 - (v3 / (((v3 / (((^(v3))) | 1))) | 1)))
        default:
            fmt.Println(v18)
        }
    } else {
        var v39 uint16 = 42
        _ = v39
    }
    v4 = append(v4, ("foo bar" + "a"))
    switch ("foo bar" + ("" + "foo bar")) {
    case "a":
        var v40 int8 = v18
        _ = v40
    case "hello":
        for fi41 := 0; fi41 < 18; fi41++ {
            if false {
                for fi42 := 0; fi42 < 9; fi42++ {
                    v43 := GenS1{F0: (^(v18)), F1: (fi42 % (((fi42 & 797)) | 1)), F2: 105}
                    if (v18 >= 127) {
                        switch (("\n\t" + (("hello" + "foo bar") + "hello")) + "\n\t") {
                        case "hello":
                            if (true || (((("hello" + ("" + "foo bar")) > "hello") && ((true || true) || true)) || ("" >= ("\n\t" + "hello")))) {
                                v44 := GenS4{F0: v1, F1: false, F2: uint16(v3)}
                                fmt.Println(v18, fi41)
                                v45 := GenS1{F0: (-(v18)), F1: (^(fi41)), F2: 2147483648}
                                _ = v44
                                _ = v45
                            }
                        case "foo bar":
                            v46 := GenS5{F0: 0, F1: 2, F2: (v1 | 42)}
                            _ = v46
                        case "a":
                            switch (("\n\t" + "hello") + ("" + ("foo bar" + "a"))) {
                            case "a":
                                var v47 int32 = 42
                                _ = v47
                            }
                        }
                    } else {
                        if false {
                            if (-1 == 42) {
                                if (618 < 255) {
                                    switch fi42 {
                                    case -9223372036854775808:
                                        if (!false) {
                                            switch 4294967295 {
                                            case 1:
                                                _ = len(v5)
                                            case -435:
                                                switch fi42 {
                                                case 42:
                                                    var v48 int = (^(fi42))
                                                    _ = v48
                                                case 44:
                                                    fmt.Println(v18)
                                                case 32768:
                                                    v4 = append(v4, (("" + ("hello" + "foo bar")) + "hello"))
                                                }
                                            default:
                                                v2 = float32(float64(3.14159))
                                            }
                                            if true {
                                                switch "a" {
                                                case "hello":
                                                    if false {
                                                        v18 += v18
                                                        if (485 < 0) {
                                                            v18 = v18
                                                            v49 := GenS4{F0: (v1 * 2), F1: (170 >= 65535), F2: 2}
                                                            _ = v49
                                                        }
                                                    } else {
                                                        if false {
                                                            if false {
                                                                if false {
                                                                    v50 := GenS4{F0: v1, F1: (((((v2 / (float32((3.14159 * 3.14159)) * v2)) / v2) * 3.14159) / 3.14159) >= (v2 * (-2.25 - v2))), F2: 2}
                                                                    _ = v50
                                                                }
                                                                var v51 string = ("hello" + "\n\t")
                                                                _ = len(v5)
                                                                _ = v51
                                                            }
                                                            if (("hello" >= "\n\t") && (!(!(v3 <= v3)))) {
                                                                var v52 bool = ((32768 == (^(fi42))) || (false && false))
                                                                var v53 float32 = 1e10
                                                                _ = v52
                                                                _ = v53
                                                            }
                                                            var v54 uint = 4294967295
                                                            _ = v54
                                                        } else {
                                                            var v55 uint8 = v1
                                                            var v56 uint = 1
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                    }
                                                case "foo bar":
                                                    v1 = (v1 - (v1 & (v1 ^ (v1 + v1))))
                                                }
                                            } else {
                                                v4 = append(v4, "hello")
                                                switch ("hello" + ("" + ("hello" + "hello"))) {
                                                case "":
                                                    var v57 uint = 2
                                                    _ = v57
                                                case "foo bar":
                                                    v58 := GenS3{F0: genFn0(), F1: ("a" + "hello"), F2: 691}
                                                    _ = v58
                                                case "a":
                                                    v1 = (v1 / ((v1) | 1))
                                                default:
                                                    v59 := GenS6{F0: -9223372036854775808, F1: true}
                                                    _ = v59
                                                }
                                                switch fi42 {
                                                case -978:
                                                    var v60 float64 = 0.1
                                                    _ = v60
                                                case 32768:
                                                    if (1 < 0) {
                                                        fmt.Println(v2)
                                                        var v61 uint = 0
                                                        if ("" >= "\n\t") {
                                                            v2 = (v2 - (v2 * 0.1))
                                                        } else {
                                                            v3 += 201
                                                            if (!true) {
                                                                if ((^(v1)) == v1) {
                                                                    var v62 uint = v61
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        switch fi42 {
                                                                        case 9223372036854775807:
                                                                            v63 := GenS3{F0: 911, F1: "foo bar", F2: 42}
                                                                            _ = v63
                                                                        case -1:
                                                                            v64 := GenS6{F0: 1, F1: true}
                                                                            _ = v64
                                                                        }
                                                                    }
                                                                    v65 := GenS1{F0: 11, F1: fi41, F2: 329}
                                                                    _ = v62
                                                                    _ = v65
                                                                }
                                                                v66 := GenS6{F0: 0, F1: true}
                                                                var v67 int8 = (v18 * (-(v18)))
                                                                _ = v66
                                                                _ = v67
                                                            } else {
                                                                var v68 int64 = -368
                                                                _ = v68
                                                            }
                                                            if len(v4) > 0 {
                                                                v4[0] = ("\n\t" + "foo bar")
                                                            }
                                                        }
                                                        _ = v61
                                                    }
                                                case 4294967295:
                                                    v69 := GenS7{F0: -2147483648, F1: "\n\t"}
                                                    _ = v69
                                                }
                                            }
                                            if false {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    if (v2 <= float32(3.14159)) {
                                                        var v70 uint = 2
                                                        fmt.Println(fi41, v2, v18)
                                                        _ = len(v4)
                                                        _ = v70
                                                    } else {
                                                        v71 := GenS8{F0: 65535, F1: "hello"}
                                                        _ = v71
                                                    }
                                                case "hello":
                                                    v2 += float32(3.14159)
                                                case "foo bar":
                                                    v72 := GenS9{F0: fi41, F1: -1, F2: 776}
                                                    _ = v72
                                                default:
                                                    switch (("" + "a") + "a") {
                                                    case "hello":
                                                        v73 := GenS10{F0: fi41}
                                                        _ = v73
                                                    case "":
                                                        v1 += (v1 ^ v1)
                                                    }
                                                }
                                                v3 -= v3
                                                v74 := GenS6{F0: int64(v3), F1: true}
                                                _ = v74
                                            }
                                            delete(v5, 42)
                                            var v75 uint = 42
                                            _ = v75
                                        } else {
                                            v4 = append(v4, "\n\t")
                                        }
                                    case -1:
                                        switch "a" {
                                        case "foo bar":
                                            var v76 int32 = 0
                                            _ = v76
                                        case "hello":
                                            if (((v1 / ((255) | 1)) != (v1 / ((v1) | 1))) && true) {
                                                if (true && ((v1 + v1) < v1)) {
                                                    var v77 uint16 = 32767
                                                    if false {
                                                        v78 := GenS1{F0: -1, F1: -1, F2: 2147483648}
                                                        v79 := GenS3{F0: 1, F1: (("hello" + ("hello" + "\n\t")) + ""), F2: 759}
                                                        _ = v78
                                                        _ = v79
                                                    } else {
                                                        var v80 int64 = 42
                                                        var v81 int64 = v80
                                                        fmt.Println(v3, v77)
                                                        _ = v80
                                                        _ = v81
                                                    }
                                                    if (!false) {
                                                        var v82 int32 = 2147483647
                                                        _ = v82
                                                    }
                                                    v3 = (^(v3))
                                                    _ = v77
                                                }
                                            }
                                        default:
                                            v83 := GenS11{F0: 1, F1: (38 == 241), F2: true}
                                            var v84 uint16 = 32767
                                            v2 -= v2
                                            _ = v83
                                            _ = v84
                                        }
                                    case 304:
                                        fmt.Println(v18, fi41, v3)
                                    }
                                    v3 -= v3
                                } else {
                                    if (42 == 9223372036854775807) {
                                        switch (fi41 & 172) {
                                        case -2147483648:
                                            if len(v4) > 0 {
                                                v4[0] = ("\n\t" + ("hello" + ""))
                                            }
                                        default:
                                            var v85 float64 = 0.0
                                            _ = v85
                                        }
                                        var v86 uint64 = v3
                                        if ((((32767 > 1) || false) && (-212 <= 42)) || (v2 > 1e10)) {
                                            v87 := v5[1]
                                            fmt.Println(v86, v2)
                                            v2 -= (float32(0.0) * v2)
                                            _ = v87
                                        } else {
                                            var v88 uint16 = 65535
                                            v89 := GenS9{F0: (fi41 >> ((18446744073709551615) & 63)), F1: -194, F2: 65535}
                                            v2 = v2
                                            var v90 bool = (fi41 > 186)
                                            _ = v88
                                            _ = v89
                                            _ = v90
                                        }
                                        _ = v86
                                    } else {
                                        v91 := GenS12{F0: (^(v18)), F1: (-264 != 9223372036854775807)}
                                        var v92 uint = 612
                                        _ = v91
                                        _ = v92
                                    }
                                    v3 += v3
                                    v93 := v5[408]
                                    v94 := v5[v93]
                                    _ = v93
                                    _ = v94
                                }
                            } else {
                                v95 := GenS13{F0: 200, F1: -1}
                                var v96 int16 = -1
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v95
                                _ = v96
                            }
                            v2 -= (v2 + v2)
                        }
                        fmt.Println(fi41, fi41, fi41)
                        var v97 int16 = 1
                        _ = v97
                    }
                    fmt.Println(v1, v18, v18)
                    _ = fi42
                    _ = v43
                }
                v1 = 255
                var v98 uint16 = 42
                _ = v98
            } else {
                for fi99 := 0; fi99 < 8; fi99++ {
                    v5[0] = 848
                    _ = fi99
                }
                fmt.Println(v18, fi41)
            }
            _ = fi41
        }
    case "foo bar":
        v1 = v1
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v18 = -17
    v1 += v1
    switch "a" {
    case "hello":
        fmt.Println(v3)
    case "foo bar":
        v3 = 126
    case "\n\t":
        _ = len(v5)
    default:
        var v100 string = "foo bar"
        _ = v100
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 = (^(v3))
    v3 = (v3 - (v3 << ((436) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
