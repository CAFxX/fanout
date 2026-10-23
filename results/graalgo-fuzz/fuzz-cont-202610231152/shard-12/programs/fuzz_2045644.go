package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int64; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS5 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 uint8; F1 uint64 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint64; F1 string; F2 uint64 }
type GenS10 struct { F0 uint16; F1 uint }
type GenS11 struct { F0 uint64; F1 int }
type GenS12 struct { F0 int; F1 string }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int16; F1 int64; F2 uint }
func genFn0() int8 {
    var v7 uint64 = 18446744073709551615
    v8 := GenS1{F0: (v7 & v7), F1: (v7 ^ v7), F2: 32767}
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"\n\t", "foo bar", (("a" + "") + ("" + "\n\t")), ("\n\t" + "\n\t")}
    var v3 int8 = 42
    var v4 int = 128
    var v5 bool = true
    var v6 int64 = int64(v3)
    switch "\n\t" {
    case "a":
        v9 := GenS1{F0: 4294967295, F1: 0, F2: 904}
        _ = v9
    case "hello":
        var v10 uint16 = 175
        _ = v10
    case "":
        for fi11 := 0; fi11 < 6; fi11++ {
            if (!v5) {
                v6 = (v6 / ((-146) | 1))
                v12 := GenS2{F0: 42, F1: (v6 & (v6 / ((v6) | 1))), F2: 2147483648}
                _ = v12
            }
            v5 = (false && v5)
            _ = fi11
        }
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    switch (v4 * (v4 >> ((887) & 63))) {
    case -455:
        v13 := GenS3{F0: 255}
        _ = v13
    default:
        switch v4 {
        case 128:
            var v14 string = ""
            _ = v14
        }
    }
    v4 = v4
    for fi15 := 0; fi15 < 2; fi15++ {
        fmt.Println(v6)
        _ = fi15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        if v5 {
            var v17 uint16 = 1
            switch (v4 / ((fi16) | 1)) {
            case 32767:
                v18 := GenS4{F0: 0, F1: v6, F2: uint64(v17)}
                _ = v18
            case -826:
                switch v4 {
                case -480:
                    var v19 int16 = -1
                    _ = v19
                case -221:
                    v5 = v5
                }
            case 1:
                switch "a" {
                case "a":
                    for fi20 := 0; fi20 < 11; fi20++ {
                        v21 := GenS1{F0: 18446744073709551615, F1: 705, F2: 32767}
                        _ = fi20
                        _ = v21
                    }
                case "hello":
                    var v22 int64 = (v6 | (v6 ^ 147))
                    _ = v22
                case "\n\t":
                    v2 = append(v2, (("\n\t" + ("hello" + (((("" + (("foo bar" + ("" + "\n\t")) + "a")) + ("a" + "")) + "a") + "foo bar"))) + (("foo bar" + ("\n\t" + (("hello" + "a") + "a"))) + "hello")))
                default:
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                }
            }
            v3 -= -1
            _ = v17
        } else {
            if ("hello" == "\n\t") {
                var v23 int = 127
                for fi24 := 0; fi24 < 5; fi24++ {
                    if (42 <= 78) {
                        fmt.Println(v23, v4)
                        switch 284 {
                        case 127:
                            var v25 bool = (-128 != genFn0())
                            _ = v25
                        case -104:
                            if v5 {
                                var v26 uint16 = 0
                                v27 := GenS3{F0: v26}
                                _ = v26
                                _ = v27
                            } else {
                                switch "" {
                                case "hello":
                                    v23 = (-(fi16))
                                case "":
                                    if false {
                                        v28 := GenS5{F0: 42, F1: -1, F2: 0}
                                        _ = v28
                                    } else {
                                        v29 := GenS6{F0: 42, F1: 0}
                                        _ = v29
                                    }
                                case "foo bar":
                                    var v30 uint16 = 0
                                    _ = v30
                                }
                            }
                        default:
                            if ((300 != 18446744073709551615) || (1 == 128)) {
                                v6 = int64(v3)
                                if v5 {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                var v31 int32 = 42
                                _ = v31
                            } else {
                                var v32 int = (v4 + v23)
                                switch "" {
                                case "":
                                    if v5 {
                                        v33 := GenS7{F0: 255, F1: 285}
                                        _ = v33
                                    }
                                case "hello":
                                    var v34 uint16 = 1
                                    if v5 {
                                        if len(v2) > 0 {
                                            v2[0] = ("hello" + "\n\t")
                                        }
                                        var v35 uint32 = 42
                                        _ = v35
                                    }
                                    if v5 {
                                        v6 += v6
                                    } else {
                                        fmt.Println(v23, v23, v4)
                                        v36 := GenS3{F0: v34}
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if ((float64(float32(1.5)) / 3.14159) == 0.0) {
                                            var v37 int = v23
                                            _ = v37
                                        } else {
                                            v6 += (v6 ^ (v6 + -553))
                                        }
                                        var v38 float64 = -2.25
                                        _ = v36
                                        _ = v38
                                    }
                                    _ = v34
                                default:
                                    var v39 float32 = 1.5
                                    _ = v39
                                }
                                _ = v32
                            }
                        }
                        _ = len(v2)
                        v40 := GenS5{F0: 2, F1: -1, F2: 178}
                        switch (("a" + "foo bar") + ("foo bar" + (("\n\t" + (("" + "\n\t") + "hello")) + "hello"))) {
                        case "foo bar":
                            switch "foo bar" {
                            case "a":
                                var v41 int8 = genFn0()
                                _ = v41
                            case "":
                                v42 := GenS8{F0: 9223372036854775807}
                                _ = v42
                            }
                        case "a":
                            v6 -= v6
                        case "":
                            var v43 float64 = 0.1
                            _ = v43
                        }
                        _ = v40
                    } else {
                        fmt.Println(fi16, fi24, v3)
                    }
                    if (!false) {
                        if v5 {
                            if len(v2) > 0 {
                                v2[0] = "\n\t"
                            }
                        } else {
                            fmt.Println(v3, v23, fi16)
                            v44 := GenS4{F0: 483, F1: (v6 ^ v6), F2: 656}
                            _ = v44
                        }
                    } else {
                        v23 -= 1
                        fmt.Println(v4, v3)
                    }
                    if v5 {
                        switch v4 {
                        case 32767:
                            var v45 int8 = v3
                            switch "a" {
                            case "\n\t":
                                switch (("hello" + (((("a" + ("\n\t" + "foo bar")) + ("\n\t" + ("\n\t" + ""))) + "a") + ("\n\t" + ((("foo bar" + ("hello" + "foo bar")) + "\n\t") + ("foo bar" + (("\n\t" + "") + "foo bar")))))) + "") {
                                case "a":
                                    var v46 uint16 = 2
                                    _ = v46
                                }
                            case "foo bar":
                                switch (fi24 + fi16) {
                                case 892:
                                    switch "foo bar" {
                                    case "\n\t":
                                        switch fi16 {
                                        case -323:
                                            var v47 int32 = 2147483647
                                            _ = v47
                                        }
                                    case "foo bar":
                                        v3 = (-(v45))
                                    case "a":
                                        if v5 {
                                            fmt.Println(fi16, v4)
                                        }
                                    default:
                                        if (("foo bar" + ("foo bar" + "foo bar")) == "\n\t") {
                                            v48 := GenS9{F0: 1, F1: ("\n\t" + "foo bar"), F2: 293}
                                            switch ("\n\t" + ("\n\t" + ("a" + "foo bar"))) {
                                            case "hello":
                                                switch (fi16 << ((1) & 63)) {
                                                case 42:
                                                    v5 = (!(!v5))
                                                default:
                                                    v5 = (("foo bar" + ("\n\t" + ("a" + "hello"))) < "hello")
                                                    fmt.Println(v3, v4, fi16)
                                                    var v49 float32 = float32(float64((((float32(0.1) / -2.25) / 1.0) * 1e10)))
                                                    _ = v49
                                                }
                                            case "\n\t":
                                                var v50 uint = 277
                                                _ = v50
                                            case "a":
                                                if len(v2) > 0 {
                                                    v2[0] = "\n\t"
                                                }
                                            default:
                                                if (460 == 65535) {
                                                    fmt.Println(v3, fi24, v4)
                                                    var v51 int64 = (v6 ^ v6)
                                                    fmt.Println(fi16)
                                                    _ = v51
                                                } else {
                                                    var v52 uint = 65535
                                                    if (33 >= 0) {
                                                        _ = len(v1)
                                                    }
                                                    var v53 string = (("\n\t" + "") + "\n\t")
                                                    _ = v52
                                                    _ = v53
                                                }
                                            }
                                            if v5 {
                                                v45 += (v3 | (v45 / (((v45 - int8(v6))) | 1)))
                                            }
                                            switch (v4 + (fi24 / (((^(v23))) | 1))) {
                                            case 9223372036854775807:
                                                v3 = (^(v3))
                                            case -2147483648:
                                                v5 = (146 < 127)
                                            case -1:
                                                v4 = v23
                                            }
                                            v23 += (fi16 ^ v23)
                                            var v54 float32 = (0.1 * 0.1)
                                            _ = v48
                                            _ = v54
                                        }
                                    }
                                case 32767:
                                    fmt.Println(fi24, v3, fi24)
                                }
                            default:
                                var v55 uint64 = 473
                                _ = v55
                            }
                            _ = v45
                        case 32768:
                            v5 = v5
                        }
                    } else {
                        var v56 string = "\n\t"
                        _ = v56
                    }
                    _ = fi24
                }
                _ = v23
            }
            if (0 != 1) {
                for fi57 := 0; fi57 < 17; fi57++ {
                    switch fi57 {
                    case 128:
                        v58 := GenS5{F0: 4294967295, F1: 42, F2: 255}
                        v59 := GenS7{F0: 128, F1: 4294967295}
                        _ = v58
                        _ = v59
                    case 20:
                        var v60 bool = v5
                        _ = v60
                    case -785:
                        if v5 {
                            if v5 {
                                if v5 {
                                    var v61 int = -2147483648
                                    _ = v61
                                }
                            }
                            if true {
                                if (false && ((42 < 639) && true)) {
                                    fmt.Println(v3)
                                } else {
                                    fmt.Println(v3, v4, fi16)
                                }
                                var v62 uint64 = 0
                                v63 := GenS10{F0: 65535, F1: 48}
                                _ = v62
                                _ = v63
                            } else {
                                v4 -= fi57
                                fmt.Println(v4, v4, v3)
                                fmt.Println(v6, v4, fi16)
                            }
                            if (174 == 255) {
                                var v64 uint = 42
                                if (!v5) {
                                    var v65 int8 = 127
                                    v5 = v5
                                    v4 = fi57
                                    _ = v65
                                } else {
                                    if false {
                                        var v66 uint32 = 1
                                        var v67 int = v4
                                        fmt.Println(v5, v4)
                                        _ = v66
                                        _ = v67
                                    } else {
                                        if (!(584 == 451)) {
                                            fmt.Println(fi57)
                                            _ = len(v1)
                                        } else {
                                            fmt.Println(v4)
                                            fmt.Println(fi16)
                                        }
                                        if v5 {
                                            fmt.Println(fi16, v3, fi16)
                                            var v68 uint32 = 352
                                            _ = v68
                                        } else {
                                            if false {
                                                v64 += v64
                                                switch 233 {
                                                case -31:
                                                    v69 := GenS10{F0: 1, F1: 188}
                                                    v1[fi16] = 4294967295
                                                    _ = v69
                                                case 32768:
                                                    var v70 int = (fi57 / ((v4) | 1))
                                                    _ = v70
                                                case -721:
                                                    if v5 {
                                                        if (!false) {
                                                            if (-32768 <= 32767) {
                                                                switch (v4 + fi16) {
                                                                case 463:
                                                                    v71 := GenS8{F0: (^(v6))}
                                                                    _ = v71
                                                                }
                                                                var v72 string = ""
                                                                v3 += 98
                                                                fmt.Println(fi57, v64, v6)
                                                                v3 = v3
                                                                _ = v72
                                                            }
                                                            fmt.Println(v3)
                                                            var v73 uint8 = 226
                                                            _ = v73
                                                        }
                                                        _ = len(v1)
                                                        switch ("a" + "\n\t") {
                                                        case "a":
                                                            v74 := GenS8{F0: v6}
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            switch "\n\t" {
                                                            case "":
                                                                fmt.Println(v6)
                                                            }
                                                            _ = v74
                                                        case "\n\t":
                                                            var v75 float32 = (1.5 + 1e10)
                                                            _ = v75
                                                        default:
                                                            var v76 bool = (0 >= 42)
                                                            _ = v76
                                                        }
                                                    } else {
                                                        switch 128 {
                                                        case 190:
                                                            if v5 {
                                                                var v77 string = "foo bar"
                                                                switch v77 {
                                                                case "hello":
                                                                    v78 := GenS11{F0: 140, F1: v4}
                                                                    if ((v77 + "a") < v77) {
                                                                        fmt.Println(v6)
                                                                    }
                                                                    _ = len(v2)
                                                                    _ = v78
                                                                }
                                                                v6 += (^(v6))
                                                                _ = v77
                                                            } else {
                                                                if (v5 && false) {
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        if v5 {
                                                                            var v79 int8 = genFn0()
                                                                            _ = v79
                                                                        }
                                                                    case "":
                                                                        var v80 int8 = v3
                                                                        _ = v80
                                                                    }
                                                                    v81 := GenS8{F0: (v6 + (v6 << ((v64) & 63)))}
                                                                    fmt.Println(fi16, v3, v4)
                                                                    _ = v81
                                                                }
                                                                v82 := GenS12{F0: 0, F1: ("foo bar" + ((((("a" + "") + ("\n\t" + "hello")) + ("" + ("foo bar" + ("" + "hello")))) + "a") + ("\n\t" + "a")))}
                                                                _ = v82
                                                            }
                                                        case -9223372036854775808:
                                                            v83 := GenS13{F0: 0}
                                                            _ = v83
                                                        default:
                                                            var v84 string = ("a" + "foo bar")
                                                            _ = v84
                                                        }
                                                        var v85 string = (("a" + "a") + ("hello" + "foo bar"))
                                                        _ = v85
                                                    }
                                                }
                                                v86 := GenS1{F0: 399, F1: 4294967295, F2: 2}
                                                _ = v86
                                            }
                                        }
                                    }
                                    v64 = (v64 << (((v64 ^ (v64 & v64))) & 63))
                                    v87 := v1[fi57]
                                    _ = v87
                                }
                                _ = v64
                            }
                        }
                    default:
                        v6 = (v6 | (^(v6)))
                    }
                    _ = fi57
                }
                if false {
                    v3 -= (-(v3))
                    _ = len(v2)
                }
                for fi88 := 0; fi88 < 4; fi88++ {
                    var v89 int = (fi88 % (((fi16 * (v4 / ((32767) | 1)))) | 1))
                    _ = fi88
                    _ = v89
                }
            } else {
                var v90 int = (fi16 << ((65535) & 63))
                v90 -= fi16
                if (420 > 113) {
                    for fi91 := 0; fi91 < 8; fi91++ {
                        switch "\n\t" {
                        case "":
                            var v92 uint8 = 1
                            _ = v92
                        case "hello":
                            var v93 uint = 255
                            _ = v93
                        case "foo bar":
                            switch fi91 {
                            case 127:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case 88:
                                fmt.Println(v6, v90)
                            case 32767:
                                _ = len(v2)
                            default:
                                v3 = 100
                            }
                        }
                        _ = fi91
                    }
                    v94 := GenS14{F0: -32768, F1: (v6 + v6), F2: 0}
                    v4 += v90
                    _ = v94
                } else {
                    v5 = v5
                    for fi95 := 0; fi95 < 14; fi95++ {
                        var v96 bool = v5
                        _ = fi95
                        _ = v96
                    }
                }
                _ = v90
            }
        }
        v2 = append(v2, "foo bar")
        var v97 uint16 = 42
        _ = fi16
        _ = v97
    }
    if (!((^(v4)) > v4)) {
        fmt.Println(v5, v4, v3)
    }
    if (v6 <= (v6 * v6)) {
        for fi98 := 0; fi98 < 3; fi98++ {
            v2 = append(v2, "\n\t")
            fmt.Println(v3, v4, v5)
            _ = fi98
        }
    }
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
