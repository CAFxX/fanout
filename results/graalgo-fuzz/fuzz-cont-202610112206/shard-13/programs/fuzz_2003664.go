package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 string; F2 int32 }
type GenS3 struct { F0 bool; F1 int32; F2 uint }
type GenS4 struct { F0 int32; F1 int64; F2 int16 }
type GenS5 struct { F0 int16; F1 uint; F2 int16 }
type GenS6 struct { F0 int8; F1 int16; F2 int16 }
type GenS7 struct { F0 int8; F1 uint64 }
type GenS8 struct { F0 int32; F1 bool }
type GenS9 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS10 struct { F0 uint; F1 uint16 }
type GenS11 struct { F0 string; F1 bool }
type GenS12 struct { F0 uint64 }
func genFn0(p0 uint32, p1 int64, p2 uint32) int32 {
    var v5 int16 = -32768
    v6 := []int{-901, -108, -2147483648, 985}
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -557
}

func main() {
    var v1 uint8 = 70
    var v2 bool = (42 == 0)
    var v3 string = "foo bar"
    v4 := []int{9223372036854775807, 545, 128}
    fmt.Println(v2, v1)
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 uint64 = 18446744073709551615
        _ = fi7
        _ = v8
    }
    if (2 < uint16(v1)) {
        if (-255 != -1) {
            switch v3 {
            case "":
                v9 := GenS1{F0: 0}
                _ = v9
            case "foo bar":
                v10 := GenS1{F0: 2}
                _ = v10
            case "hello":
                var v11 bool = (-1 < genFn0(2, 0, 2147483647))
                _ = v11
            default:
                var v12 uint16 = 32767
                _ = v12
            }
        } else {
            switch (v3 + v3) {
            case "foo bar":
                for fi13 := 0; fi13 < 3; fi13++ {
                    v2 = (32767 == 593)
                    _ = fi13
                }
            case "a":
                fmt.Println(v1)
            case "hello":
                if (!(2 < 283)) {
                    if v2 {
                        for fi14 := 0; fi14 < 12; fi14++ {
                            for fi15 := 0; fi15 < 15; fi15++ {
                                if (!false) {
                                    if v2 {
                                        var v16 uint16 = 42
                                        var v17 uint = 0
                                        var v18 float32 = 0.0
                                        var v19 int32 = genFn0(4294967295, -1, 495)
                                        _ = v16
                                        _ = v17
                                        _ = v18
                                        _ = v19
                                    } else {
                                        var v20 float32 = ((1.5 + 1.5) + 0.0)
                                        _ = len(v4)
                                        v21 := GenS2{F0: 206, F1: "\n\t", F2: -1}
                                        _ = v20
                                        _ = v21
                                    }
                                    var v22 int = 9223372036854775807
                                    _ = v22
                                }
                                if v2 {
                                    fmt.Println(v1, fi14, v2)
                                } else {
                                    fmt.Println(fi15, v3)
                                    var v23 int16 = -32768
                                    v24 := GenS2{F0: uint16(v1), F1: "hello", F2: -1}
                                    _ = v23
                                    _ = v24
                                }
                                _ = fi15
                            }
                            if (127 >= (v1 << ((42) & 63))) {
                                var v25 int8 = 127
                                for fi26 := 0; fi26 < 16; fi26++ {
                                    if len(v4) > 0 {
                                        v4[0] = -9223372036854775808
                                    }
                                    v25 = v25
                                    var v27 int32 = 2147483647
                                    _ = fi26
                                    _ = v27
                                }
                                _ = v25
                            }
                            v28 := GenS1{F0: 4294967295}
                            _ = fi14
                            _ = v28
                        }
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    if (741 > 340) {
                        v1 += v1
                        for fi29 := 0; fi29 < 3; fi29++ {
                            v3 = (v3 + (v3 + v3))
                            _ = fi29
                        }
                        switch 0 {
                        case -9223372036854775808:
                            for fi30 := 0; fi30 < 10; fi30++ {
                                fmt.Println(v3)
                                _ = fi30
                            }
                        default:
                            if true {
                                for fi31 := 0; fi31 < 11; fi31++ {
                                    var v32 uint32 = 0
                                    switch fi31 {
                                    case -449:
                                        for fi33 := 0; fi33 < 13; fi33++ {
                                            var v34 float64 = -2.25
                                            _ = fi33
                                            _ = v34
                                        }
                                    case 9223372036854775807:
                                        fmt.Println(v3, v1, v1)
                                    default:
                                        if true {
                                            for fi35 := 0; fi35 < 20; fi35++ {
                                                v36 := GenS3{F0: false, F1: 706, F2: 42}
                                                v32 -= (^(v32))
                                                _ = len(v4)
                                                if v2 {
                                                    var v37 uint64 = uint64(v1)
                                                    var v38 int32 = -212
                                                    v38 -= (v38 - (v38 << ((uint(v1)) & 63)))
                                                    v39 := GenS2{F0: 1, F1: ("" + "\n\t"), F2: v38}
                                                    _ = v37
                                                    _ = v38
                                                    _ = v39
                                                } else {
                                                    _ = len(v4)
                                                    switch (fi35 - fi35) {
                                                    case 2147483647:
                                                        if true {
                                                            switch (v3 + ("\n\t" + v3)) {
                                                            case "hello":
                                                                var v40 string = v3
                                                                _ = v40
                                                            }
                                                            v1 = 53
                                                        }
                                                    case -1:
                                                        v2 = ((!(false || v2)) || v2)
                                                        if (!(3.14159 > 0.1)) {
                                                            v2 = (1 < 42)
                                                            var v41 uint = 4294967295
                                                            _ = v41
                                                        }
                                                    case -2147483648:
                                                        fmt.Println(fi35, v32)
                                                    }
                                                }
                                                _ = fi35
                                                _ = v36
                                            }
                                            v3 = v3
                                        } else {
                                            for fi42 := 0; fi42 < 13; fi42++ {
                                                var v43 int = (fi31 | (fi42 | (fi31 + (-(fi42)))))
                                                v44 := GenS2{F0: 292, F1: "", F2: 1}
                                                _ = fi42
                                                _ = v43
                                                _ = v44
                                            }
                                            for fi45 := 0; fi45 < 11; fi45++ {
                                                v3 = (("foo bar" + (v3 + v3)) + (v3 + v3))
                                                var v46 uint32 = v32
                                                switch (fi31 + (fi31 % ((fi31) | 1))) {
                                                case 127:
                                                    fmt.Println(fi45)
                                                case 32767:
                                                    var v47 string = v3
                                                    _ = v47
                                                case -1:
                                                    v2 = v2
                                                }
                                                if (!(-128 != 42)) {
                                                    var v48 uint64 = 2
                                                    v49 := GenS4{F0: 420, F1: -9223372036854775808, F2: 42}
                                                    switch "foo bar" {
                                                    case "hello":
                                                        fmt.Println(v2, v2, v48)
                                                    }
                                                    _ = v48
                                                    _ = v49
                                                }
                                                switch -431 {
                                                case 1:
                                                    var v50 int16 = 0
                                                    var v51 int64 = -9223372036854775808
                                                    _ = v50
                                                    _ = v51
                                                case 32768:
                                                    switch (v3 + v3) {
                                                    case "\n\t":
                                                        v52 := GenS1{F0: 132}
                                                        _ = v52
                                                    case "hello":
                                                        if (-2.25 == -2.25) {
                                                            v53 := GenS5{F0: 712, F1: 0, F2: 0}
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            var v54 uint32 = (^(v46))
                                                            var v55 bool = (1 >= 9223372036854775807)
                                                            v55 = (true && false)
                                                            _ = v53
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                    default:
                                                        var v56 int64 = -9223372036854775808
                                                        _ = v56
                                                    }
                                                }
                                                _ = fi45
                                                _ = v46
                                            }
                                        }
                                    }
                                    _ = fi31
                                    _ = v32
                                }
                            } else {
                                fmt.Println(v2)
                                switch v3 {
                                case "\n\t":
                                    switch -960 {
                                    case 127:
                                        var v57 uint = 4294967295
                                        _ = v57
                                    case 9223372036854775807:
                                        if (v2 || (v2 || (!(false || (32767 != 302))))) {
                                            v2 = v2
                                            if (-9223372036854775808 >= 533) {
                                                for fi58 := 0; fi58 < 18; fi58++ {
                                                    v2 = (v2 || false)
                                                    for fi59 := 0; fi59 < 11; fi59++ {
                                                        v3 = "\n\t"
                                                        _ = fi59
                                                    }
                                                    for fi60 := 0; fi60 < 13; fi60++ {
                                                        v61 := GenS6{F0: 41, F1: -1, F2: -1}
                                                        var v62 uint16 = 2
                                                        _ = fi60
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                    _ = fi58
                                                }
                                            } else {
                                                if false {
                                                    fmt.Println(v3, v3, v1)
                                                    fmt.Println(v2, v2)
                                                    v3 = (("\n\t" + (v3 + v3)) + v3)
                                                }
                                            }
                                        } else {
                                            for fi63 := 0; fi63 < 20; fi63++ {
                                                v2 = (((1e10 + 1.5) / -2.25) == -2.25)
                                                var v64 int = fi63
                                                _ = fi63
                                                _ = v64
                                            }
                                            v3 = v3
                                        }
                                    }
                                case "a":
                                    v65 := GenS4{F0: -713, F1: -277, F2: 1000}
                                    _ = v65
                                case "hello":
                                    v66 := GenS5{F0: 1, F1: 899, F2: 873}
                                    _ = v66
                                default:
                                    var v67 int16 = 199
                                    _ = v67
                                }
                            }
                        }
                    } else {
                        v68 := GenS6{F0: -1, F1: 237, F2: -244}
                        _ = v68
                    }
                }
            }
            v69 := GenS7{F0: 0, F1: 1}
            _ = v69
        }
        var v70 uint32 = 2
        var v71 bool = (681 <= 42)
        _ = v70
        _ = v71
    }
    var v72 float32 = 1.5
    var v73 uint16 = 255
    if v2 {
        switch 32768 {
        case 128:
            if len(v4) > 0 {
                v4[0] = 268
            }
        case 32767:
            if v2 {
                v72 += v72
                v73 += uint16(v1)
            } else {
                var v74 bool = v2
                if v74 {
                    fmt.Println(v73)
                } else {
                    if v74 {
                        fmt.Println(v72)
                        switch 9223372036854775807 {
                        case 128:
                            v3 = "foo bar"
                        case -136:
                            fmt.Println(v3, v74)
                        case 1:
                            v75 := GenS8{F0: 1, F1: true}
                            _ = v75
                        }
                        if v2 {
                            v76 := GenS9{F0: 2147483647, F1: 42, F2: 0}
                            v77 := GenS4{F0: 1, F1: 459, F2: 1000}
                            v78 := GenS3{F0: (true || true), F1: 1, F2: 146}
                            _ = v76
                            _ = v77
                            _ = v78
                        } else {
                            v74 = false
                            v73 = (^(v73))
                        }
                    }
                    var v79 int16 = -208
                    _ = v79
                }
                if v2 {
                    switch v3 {
                    case "foo bar":
                        v80 := GenS1{F0: 2}
                        _ = v80
                    case "hello":
                        for fi81 := 0; fi81 < 3; fi81++ {
                            for fi82 := 0; fi82 < 18; fi82++ {
                                v73 = (^(v73))
                                var v83 int = fi82
                                _ = fi82
                                _ = v83
                            }
                            _ = fi81
                        }
                    case "":
                        if v74 {
                            for fi84 := 0; fi84 < 17; fi84++ {
                                var v85 int8 = -1
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = fi84
                                _ = v85
                            }
                            v86 := GenS7{F0: -60, F1: 42}
                            _ = v86
                        } else {
                            fmt.Println(v2)
                            v74 = (((v3 + ("\n\t" + v3)) <= (v3 + v3)) && ((v3 + "foo bar") <= v3))
                        }
                    default:
                        _ = len(v4)
                    }
                    if len(v4) > 0 {
                        v4[0] = 42
                    }
                    v72 -= v72
                }
                _ = v74
            }
        case 580:
            v2 = (4294967295 != 41)
        default:
            fmt.Println(v2, v3)
        }
        v72 = v72
    }
    for fi87 := 0; fi87 < 16; fi87++ {
        for fi88 := 0; fi88 < 18; fi88++ {
            v89 := GenS3{F0: v2, F1: 230, F2: 255}
            if (v2 && false) {
                _ = len(v4)
            }
            v90 := GenS7{F0: -1, F1: 4294967295}
            _ = fi88
            _ = v89
            _ = v90
        }
        _ = fi87
    }
    v1 += v1
    for fi91 := 0; fi91 < 5; fi91++ {
        var v92 bool = false
        v1 = 128
        for fi93 := 0; fi93 < 6; fi93++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            fmt.Println(v92, v72)
            v73 = v73
            _ = len(v4)
            _ = len(v4)
            _ = fi93
        }
        _ = fi91
        _ = v92
    }
    if ((v1 * v1) <= (v1 * (v1 * v1))) {
        v94 := GenS10{F0: 255, F1: v73}
        v2 = (2147483647 < 42)
        v72 -= (float32((1.5 + -2.25)) - float32(3.14159))
        _ = v94
    }
    if ((32767 > (v73 & v73)) && v2) {
        v1 += (v1 ^ v1)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch "hello" {
        case "\n\t":
            for fi95 := 0; fi95 < 13; fi95++ {
                var v96 uint32 = 230
                v3 = ("hello" + v3)
                v72 += 0.1
                _ = fi95
                _ = v96
            }
        case "":
            fmt.Println(v2, v3)
        }
    } else {
        v97 := GenS11{F0: "foo bar", F1: (!(((v73 % ((v73) | 1)) >= 42) && v2))}
        for fi98 := 0; fi98 < 2; fi98++ {
            if v2 {
                v2 = (!(!true))
                var v99 int64 = -66
                v100 := GenS11{F0: ("foo bar" + ("\n\t" + "")), F1: ((0 >= -1) || true)}
                _ = v100
                _ = v99
            }
            v101 := GenS12{F0: 4294967295}
            if v2 {
                switch ((v3 + "foo bar") + "a") {
                case "":
                    for fi102 := 0; fi102 < 9; fi102++ {
                        v72 = ((((v72 * v72) * (v72 - v72)) / v72) + (((v72 - ((v72 - ((v72 - 3.14159) - (3.14159 + (v72 * (0.0 * (v72 / v72)))))) / float32(-2.25))) / v72) * 3.14159))
                        if (!(0 >= -692)) {
                            fmt.Println(fi102, v72, v1)
                        }
                        _ = fi102
                    }
                case "hello":
                    if v2 {
                        fmt.Println(fi98, v2)
                    } else {
                        for fi103 := 0; fi103 < 3; fi103++ {
                            fmt.Println(v3)
                            v72 = v72
                            _ = fi103
                        }
                    }
                }
            }
            _ = fi98
            _ = v101
        }
        _ = v97
    }
    var v104 uint64 = 18446744073709551615
    var v105 string = "foo bar"
    _ = len(v4)
    if ((-1 > 127) || v2) {
        v1 -= v1
    } else {
        v72 += (v72 * v72)
        var v106 bool = (255 >= 2)
        _ = v106
    }
    v104 -= v104
    v3 = ""
    var v107 uint16 = (v73 + v73)
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v107:", v107)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v104
    _ = v105
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v72
    _ = v73
}
