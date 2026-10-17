package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint64; F2 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int; F1 uint8; F2 bool }
type GenS8 struct { F0 bool; F1 int }
type GenS9 struct { F0 int64; F1 bool }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint64; F1 int32; F2 int16 }
type GenS13 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int32) uint16 {
    var v6 bool = (!false)
    if v6 {
        fmt.Println(v6, v6)
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            v9 := GenS1{F0: (-(fi7)), F1: "\n\t", F2: (p0 / ((-2147483648) | 1))}
            v10 := []int{-2147483648, (fi8 / (((fi8 - fi7)) | 1))}
            _ = fi8
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 16; fi11++ {
            v12 := GenS1{F0: (^(fi7)), F1: "", F2: (p0 ^ (^(p0)))}
            if v6 {
                var v13 uint = 405
                _ = v13
            } else {
                if (1 == 207) {
                    p0 -= (p0 / ((p0) | 1))
                    var v14 uint16 = 351
                    _ = v14
                }
                switch "foo bar" {
                case "hello":
                    v15 := []int{}
                    _ = v15
                case "a":
                    fmt.Println(p0)
                }
            }
            v16 := []int{fi11, 813}
            _ = fi11
            _ = v12
            _ = v16
        }
        _ = fi7
    }
    switch "hello" {
    case "a":
        var v17 float32 = (0.1 / 1e10)
        _ = v17
    case "":
        v18 := GenS1{F0: 2147483647, F1: "", F2: (p0 + p0)}
        _ = v18
    default:
        v6 = v6
    }
    v6 = true
    _ = p0
    _ = v6
    return 255
}

func main() {
    v1 := []int{4294967295, 32767, -2147483648, 127}
    v2 := []int{}
    v3 := []string{("a" + (("" + ("a" + ("" + ("hello" + ("foo bar" + "a"))))) + "foo bar")), "\n\t", "\n\t"}
    var v4 int8 = -52
    var v5 int32 = -2147483648
    v19 := GenS1{F0: 127, F1: "\n\t", F2: (v5 * v5)}
    if false {
        if true {
            switch 32768 {
            case 4294967295:
                fmt.Println(v4)
            case -2147483648:
                switch ("\n\t" + ("\n\t" + ("foo bar" + "\n\t"))) {
                case "foo bar":
                    v20 := GenS2{F0: -901}
                    _ = v20
                case "\n\t":
                    if true {
                        var v21 uint16 = 42
                        v22 := GenS3{F0: 0, F1: 1, F2: ""}
                        _ = v21
                        _ = v22
                    } else {
                        if ("a" < "hello") {
                            v5 = v5
                        }
                        for fi23 := 0; fi23 < 2; fi23++ {
                            v24 := GenS4{F0: 103}
                            for fi25 := 0; fi25 < 5; fi25++ {
                                v5 += (v5 - (v5 * int32(v4)))
                                v26 := GenS5{F0: 18446744073709551615}
                                fmt.Println(fi25)
                                _ = fi25
                                _ = v26
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi23
                            _ = v24
                        }
                    }
                case "hello":
                    fmt.Println(v4, v5)
                default:
                    switch "\n\t" {
                    case "":
                        v5 = -1
                    case "\n\t":
                        for fi27 := 0; fi27 < 17; fi27++ {
                            var v28 bool = false
                            v29 := GenS3{F0: 712, F1: 2, F2: ((("a" + "foo bar") + ("a" + ("" + ("a" + ((("\n\t" + ("hello" + "hello")) + "") + ("foo bar" + "")))))) + "foo bar")}
                            v4 -= (v4 % ((-1) | 1))
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    default:
                        v5 = (v5 >> ((0) & 63))
                    }
                }
            }
            if true {
                var v30 bool = false
                _ = v30
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v31 := GenS5{F0: 1}
                _ = v31
            }
            if (42 < 127) {
                v32 := GenS3{F0: 661, F1: 18446744073709551615, F2: "hello"}
                for fi33 := 0; fi33 < 10; fi33++ {
                    fmt.Println(v5, v5, fi33)
                    if (24 == 18446744073709551615) {
                        var v34 float32 = (0.0 - -2.25)
                        _ = v34
                    }
                    _ = len(v1)
                    _ = fi33
                }
                _ = v32
            }
        } else {
            var v35 uint8 = 1
            var v36 int16 = 0
            if false {
                var v37 float32 = 0.0
                _ = v37
            }
            _ = v35
            _ = v36
        }
    }
    if (!false) {
        for fi38 := 0; fi38 < 6; fi38++ {
            var v39 int8 = 42
            for fi40 := 0; fi40 < 20; fi40++ {
                var v41 int64 = -9223372036854775808
                _ = len(v3)
                _ = fi40
                _ = v41
            }
            _ = fi38
            _ = v39
        }
        v42 := GenS5{F0: 2}
        if false {
            v4 = v4
            v4 += (v4 >> ((778) & 63))
        } else {
            fmt.Println(v4, v4)
            fmt.Println(v4)
        }
        _ = v42
    } else {
        v5 = 235
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v43 float32 = (1e10 + (1.5 + (float32(((0.0 + (1.5 - 1e10)) - ((0.1 - (3.14159 * (-2.25 + 3.14159))) + float64(float32(3.14159))))) - 0.0)))
    for fi44 := 0; fi44 < 2; fi44++ {
        if false {
            var v45 uint64 = 0
            if true {
                for fi46 := 0; fi46 < 10; fi46++ {
                    var v47 string = "hello"
                    _ = fi46
                    _ = v47
                }
                fmt.Println(v4)
            }
            if true {
                for fi48 := 0; fi48 < 19; fi48++ {
                    v43 += v43
                    if false {
                        var v49 int64 = -9223372036854775808
                        var v50 uint16 = 32767
                        var v51 uint8 = 127
                        _ = v49
                        _ = v50
                        _ = v51
                    }
                    v52 := GenS6{F0: -709}
                    _ = fi48
                    _ = v52
                }
                var v53 int16 = -1
                _ = len(v1)
                _ = v53
            }
            _ = v45
        } else {
            v54 := GenS4{F0: 0}
            v55 := GenS2{F0: -32768}
            _ = v54
            _ = v55
        }
        if true {
            var v56 string = "hello"
            fmt.Println(fi44)
            _ = v56
        } else {
            for fi57 := 0; fi57 < 2; fi57++ {
                var v58 uint = 4294967295
                var v59 int = (fi57 - fi44)
                _ = fi57
                _ = v58
                _ = v59
            }
        }
        v43 = (v43 - (v43 - float32(0.0)))
        _ = fi44
    }
    if ((!(659 != 4294967295)) && (!false)) {
        var v60 float64 = 0.0
        _ = v60
    }
    v61 := GenS2{F0: 1000}
    for fi62 := 0; fi62 < 1; fi62++ {
        switch (fi62 - (-(fi62))) {
        case 127:
            if false {
                fmt.Println(fi62, v43)
                v4 = v4
            } else {
                v3 = append(v3, "hello")
                var v63 int = -1
                _ = v63
            }
        case -9223372036854775808:
            var v64 uint32 = 2
            _ = v64
        default:
            v65 := GenS7{F0: fi62, F1: 2, F2: true}
            _ = v65
        }
        v3 = append(v3, "hello")
        switch ("a" + "foo bar") {
        case "\n\t":
            for fi66 := 0; fi66 < 15; fi66++ {
                if (!(!((!false) && (!true)))) {
                    fmt.Println(fi66, v5)
                    var v67 bool = (0 < -971)
                    _ = v67
                }
                _ = fi66
            }
        case "a":
            if len(v1) > 0 {
                v1[0] = fi62
            }
        case "":
            if ((2147483648 < 151) || (67 == 32767)) {
                var v68 uint8 = 127
                switch "hello" {
                case "\n\t":
                    switch "foo bar" {
                    case "":
                        if false {
                            v43 = v43
                        }
                    }
                case "foo bar":
                    v69 := GenS8{F0: (65535 <= 2147483648), F1: fi62}
                    _ = v69
                default:
                    if (1 <= 255) {
                        var v70 string = "a"
                        fmt.Println(v70, v43, v68)
                        var v71 int32 = (v5 + v5)
                        _ = v70
                        _ = v71
                    }
                }
                _ = v68
            } else {
                switch (fi62 & (fi62 ^ fi62)) {
                case 0:
                    for fi72 := 0; fi72 < 2; fi72++ {
                        var v73 bool = ("hello" != "a")
                        v43 -= v43
                        switch ("\n\t" + ("\n\t" + "\n\t")) {
                        case "\n\t":
                            v5 += v5
                        case "foo bar":
                            if v73 {
                                v4 += int8(v5)
                            } else {
                                v74 := GenS9{F0: -9223372036854775808, F1: v73}
                                v75 := GenS10{F0: 94}
                                v76 := GenS10{F0: 32767}
                                _ = v74
                                _ = v75
                                _ = v76
                            }
                        case "hello":
                            fmt.Println(v43, fi62)
                            _ = len(v3)
                            switch int(v5) {
                            case 1:
                                var v77 float64 = 0.0
                                _ = v77
                            case 971:
                                _ = len(v1)
                            case 127:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            default:
                                switch "foo bar" {
                                case "\n\t":
                                    var v78 int = (fi72 * int(v5))
                                    _ = v78
                                case "foo bar":
                                    switch "" {
                                    case "a":
                                        v43 -= (v43 / (float32(float64(float32(float64((1.5 / 3.14159))))) + v43))
                                    case "hello":
                                        if v73 {
                                            if len(v3) > 0 {
                                                v3[0] = "foo bar"
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v79 uint32 = 2147483647
                                            if true {
                                                var v80 bool = true
                                                if (!true) {
                                                    var v81 int = (fi72 << ((2) & 63))
                                                    if false {
                                                        v5 -= v5
                                                    } else {
                                                        var v82 uint = 375
                                                        _ = v82
                                                    }
                                                    fmt.Println(fi62, v73, v80)
                                                    _ = v81
                                                }
                                                _ = v80
                                            }
                                            var v83 uint64 = 510
                                            v84 := GenS8{F0: (((v43 * v43) / 1e10) != 3.14159), F1: (fi72 % (((fi62 >> ((287) & 63))) | 1))}
                                            _ = v79
                                            _ = v83
                                            _ = v84
                                        }
                                        fmt.Println(v43, fi72)
                                        v43 = ((1e10 - float32(-2.25)) + v43)
                                    }
                                case "a":
                                    fmt.Println(fi62, fi72, fi72)
                                }
                                if (!v73) {
                                    if true {
                                        var v85 int16 = 736
                                        v86 := GenS11{F0: fi72}
                                        if true {
                                            if (-656 != (v85 >> ((141) & 63))) {
                                                v43 = v43
                                            } else {
                                                v43 -= float32(0.1)
                                                v87 := GenS12{F0: 2, F1: v5, F2: v85}
                                                _ = v87
                                            }
                                            fmt.Println(fi62, v43, v73)
                                        }
                                        _ = v85
                                        _ = v86
                                    }
                                    var v88 uint8 = 189
                                    _ = v88
                                }
                                v4 = (^(v4))
                                v43 = float32((-2.25 / 1.0))
                            }
                        default:
                            v73 = (-2.25 < v43)
                        }
                        _ = fi72
                        _ = v73
                    }
                case 942:
                    v5 = v5
                case 2147483647:
                    for fi89 := 0; fi89 < 13; fi89++ {
                        if false {
                            fmt.Println(v43, v4)
                            switch fi89 {
                            case 1:
                                v43 -= float32(1e10)
                            case 0:
                                if false {
                                    v5 -= (v5 + (v5 & v5))
                                    var v90 uint64 = 0
                                    _ = v90
                                } else {
                                    switch fi62 {
                                    case 4294967295:
                                        fmt.Println(v5, fi89, v4)
                                    }
                                }
                            default:
                                var v91 float32 = (v43 * (1e10 - 1e10))
                                _ = v91
                            }
                            v92 := GenS9{F0: 0, F1: true}
                            _ = v92
                        } else {
                            switch (fi62 >> ((817) & 63)) {
                            case -568:
                                var v93 uint = 782
                                _ = v93
                            case 0:
                                if true {
                                    v5 = v5
                                    _ = len(v1)
                                }
                            default:
                                if len(v2) > 0 {
                                    v2[0] = (-(fi89))
                                }
                            }
                            var v94 int = (fi89 | -2147483648)
                            _ = v94
                        }
                        switch "foo bar" {
                        case "\n\t":
                            var v95 bool = false
                            _ = v95
                        }
                        v96 := GenS3{F0: 4294967295, F1: 1, F2: "\n\t"}
                        _ = fi89
                        _ = v96
                    }
                default:
                    _ = len(v1)
                }
                v97 := GenS12{F0: 382, F1: (v5 * v5), F2: 1}
                if false {
                    v98 := GenS13{F0: 18446744073709551615, F1: 18446744073709551615}
                    _ = v98
                }
                _ = v97
            }
        }
        _ = fi62
    }
    var v99 int16 = 1
    var v100 int = 128
    if len(v1) > 0 {
        v1[0] = -9223372036854775808
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v100:", v100)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v61
    _ = v99
}
