package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint64; F2 uint }
type GenS4 struct { F0 int32; F1 uint32; F2 int32 }
type GenS5 struct { F0 int16; F1 uint32 }
type GenS6 struct { F0 int64; F1 int; F2 string }
type GenS7 struct { F0 int8; F1 uint32 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int8; F1 int32 }
type GenS10 struct { F0 uint64; F1 int; F2 int8 }
type GenS11 struct { F0 uint16; F1 bool }
type GenS12 struct { F0 string; F1 int8; F2 uint16 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int32; F1 uint; F2 int64 }
func genFn0() int16 {
    var v7 string = "\n\t"
    var v8 int64 = 42
    if false {
        var v9 uint8 = 128
        _ = v9
    } else {
        fmt.Println(v8, v7)
    }
    v10 := []int{-9223372036854775808, 2147483647, -2147483648, -2147483648}
    _ = v10
    _ = v7
    _ = v8
    return 592
}
func genFn1(p0 uint32) int {
    v11 := []int{}
    var v12 string = "foo bar"
    for fi13 := 0; fi13 < 7; fi13++ {
        switch "" {
        case "hello":
            v14 := GenS1{F0: (fi13 + fi13), F1: 127}
            _ = v14
        case "a":
            for fi15 := 0; fi15 < 17; fi15++ {
                v11 = append(v11, fi15)
                _ = fi15
            }
        case "":
            v11 = append(v11, -1)
        default:
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
        }
        if (32767 != fi13) {
            v16 := GenS2{F0: 687}
            if len(v11) > 0 {
                v11[0] = 9223372036854775807
            }
            if true {
                var v17 int8 = -1
                switch (^(fi13)) {
                case 1:
                    v18 := GenS2{F0: (p0 | p0)}
                    _ = v18
                case 9223372036854775807:
                    p0 = (^(p0))
                case 150:
                    var v19 uint64 = 4294967295
                    _ = v19
                default:
                    v12 = ((v12 + (v12 + ((((((("a" + v12) + v12) + v12) + (v12 + "foo bar")) + v12) + v12) + "a"))) + (((((v12 + "a") + v12) + v12) + (((v12 + v12) + v12) + "a")) + v12))
                }
                v12 = v12
                _ = v17
            }
            _ = v16
        } else {
            var v20 uint = 304
            _ = v20
        }
        var v21 float32 = (float32(float64(3.14159)) / 1.5)
        _ = fi13
        _ = v21
    }
    _ = p0
    _ = v11
    _ = v12
    return 190
}

func main() {
    var v1 int16 = 0
    var v2 float32 = (((1.5 - 0.1) / -2.25) / 0.1)
    v3 := make(map[uint32]int)
    var v4 float32 = v2
    v5 := []uint32{}
    var v6 int64 = -1
    fmt.Println(v2)
    for fi22 := 0; fi22 < 10; fi22++ {
        v23 := v3[187]
        if ((32767 != 26) && true) {
            if (false || false) {
                v6 = (-(v6))
            } else {
                fmt.Println(v1, fi22, fi22)
                v6 += v6
                if false {
                    if ((v1 << ((255) & 63)) >= genFn0()) {
                        fmt.Println(v1)
                        if ((855 != 65535) || (!true)) {
                            if true {
                                switch "hello" {
                                case "hello":
                                    v24 := GenS2{F0: 2147483647}
                                    _ = v24
                                }
                                if (!false) {
                                    v4 = (v4 * ((v2 - (v4 - (((float32(1e10) / v2) / v4) * v4))) / 1.5))
                                    for fi25 := 0; fi25 < 16; fi25++ {
                                        v5 = append(v5, 65535)
                                        var v26 int = (fi25 ^ -314)
                                        _ = fi25
                                        _ = v26
                                    }
                                    if true {
                                        for fi27 := 0; fi27 < 8; fi27++ {
                                            var v28 int = (^(v23))
                                            v29 := v3[2147483647]
                                            if (int(v1) > (fi22 * fi22)) {
                                                v23 = genFn1(841)
                                                _ = len(v5)
                                                switch 0 {
                                                case 4294967295:
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                case 32767:
                                                    switch (v28 % ((genFn1(2147483647)) | 1)) {
                                                    case 42:
                                                        var v30 int = genFn1(65535)
                                                        _ = v30
                                                    case -547:
                                                        v2 += 1.5
                                                        fmt.Println(v1, v6, fi22)
                                                        var v31 float32 = v2
                                                        _ = v31
                                                    case 32767:
                                                        if (false && (v1 > (v1 + (v1 / (((v1 % (((^(v1))) | 1))) | 1))))) {
                                                            v32 := GenS2{F0: 166}
                                                            v4 += float32(((1e10 - (((((float64(v2) * (0.1 - -2.25)) - 0.0) + (3.14159 / (float64(v4) / 1.0))) / 1.0) - 0.1)) + -2.25))
                                                            if ((255 <= 0) || false) {
                                                                v33 := GenS1{F0: v28, F1: 1}
                                                                switch "" {
                                                                case "hello":
                                                                    var v34 uint32 = 1
                                                                    _ = v34
                                                                case "foo bar":
                                                                    v35 := GenS3{F0: 42, F1: 42, F2: 0}
                                                                    _ = v35
                                                                default:
                                                                    var v36 int = (fi27 / ((v29) | 1))
                                                                    _ = v36
                                                                }
                                                                switch "hello" {
                                                                case "a":
                                                                    var v37 bool = (4294967295 < 2)
                                                                    _ = v37
                                                                case "":
                                                                    if true {
                                                                        fmt.Println(v28)
                                                                        switch (v29 % (((^(v28))) | 1)) {
                                                                        case 32768:
                                                                            var v38 uint = 4294967295
                                                                            _ = v38
                                                                        case 128:
                                                                            v39 := GenS4{F0: -260, F1: 0, F2: 1}
                                                                            switch "a" {
                                                                            case "\n\t":
                                                                                switch v28 {
                                                                                case 4294967295:
                                                                                    var v40 uint64 = 957
                                                                                    _ = v40
                                                                                case -1:
                                                                                    var v41 float64 = 0.1
                                                                                    _ = v41
                                                                                case 445:
                                                                                    var v42 bool = (((-9223372036854775808 <= -589) || ((-(v1)) != genFn0())) || (((v6 + -1) == v6) && false))
                                                                                    _ = v42
                                                                                default:
                                                                                    v43 := v3[2]
                                                                                    _ = v43
                                                                                }
                                                                            case "":
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                            case "foo bar":
                                                                                switch "\n\t" {
                                                                                case "a":
                                                                                    if true {
                                                                                        var v44 uint = 42
                                                                                        _ = v44
                                                                                    }
                                                                                case "foo bar":
                                                                                    v1 -= (-(v1))
                                                                                case "":
                                                                                    fmt.Println(v23, v28, fi22)
                                                                                }
                                                                            }
                                                                            var v45 int = (v29 & v23)
                                                                            _ = v39
                                                                            _ = v45
                                                                        }
                                                                    } else {
                                                                        if true {
                                                                            var v46 bool = true
                                                                            v47 := GenS4{F0: 441, F1: 65535, F2: 1}
                                                                            var v48 uint64 = 4294967295
                                                                            _ = v46
                                                                            _ = v47
                                                                            _ = v48
                                                                        } else {
                                                                            fmt.Println(v23, v6, v1)
                                                                            fmt.Println(v4, v2, v2)
                                                                            if ((1 > (v6 << ((0) & 63))) || (1 < 4294967295)) {
                                                                                v6 -= (^(v6))
                                                                            } else {
                                                                                v49 := GenS5{F0: v1, F1: 610}
                                                                                _ = v49
                                                                            }
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    switch 0 {
                                                                    case 127:
                                                                        if true {
                                                                            if false {
                                                                                var v50 bool = true
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                                var v51 int = (v28 - fi22)
                                                                                _ = v50
                                                                                _ = v51
                                                                            }
                                                                        } else {
                                                                            var v52 int64 = (v6 - 1)
                                                                            _ = v52
                                                                        }
                                                                    default:
                                                                        var v53 string = ""
                                                                        _ = v53
                                                                    }
                                                                default:
                                                                    v1 = (v1 << ((723) & 63))
                                                                }
                                                                _ = v33
                                                            }
                                                            switch 32767 {
                                                            case 128:
                                                                v2 -= v2
                                                            default:
                                                                v54 := GenS6{F0: v6, F1: genFn1(831), F2: "a"}
                                                                _ = v54
                                                            }
                                                            _ = v32
                                                        } else {
                                                            var v55 float32 = v4
                                                            var v56 int64 = v6
                                                            switch (v29 * (-(v28))) {
                                                            case -1:
                                                                v5 = append(v5, 2147483648)
                                                            default:
                                                                fmt.Println(v55, v56)
                                                            }
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                    default:
                                                        if true {
                                                            delete(v3, 65535)
                                                            v57 := GenS4{F0: -1, F1: 2, F2: -9}
                                                            v2 -= float32((1.5 + 0.1))
                                                            var v58 int64 = 42
                                                            _ = v57
                                                            _ = v58
                                                        }
                                                    }
                                                default:
                                                    v28 -= v23
                                                }
                                            } else {
                                                switch "hello" {
                                                case "hello":
                                                    if (18446744073709551615 > 2) {
                                                        switch v23 {
                                                        case -2147483648:
                                                            if (!true) {
                                                                v3[2147483647] = (v28 | (fi22 << ((108) & 63)))
                                                                var v59 string = "a"
                                                                _ = v59
                                                            } else {
                                                                v4 = ((v4 / (v2 + 0.0)) + ((v4 * (-2.25 - v4)) / float32(-2.25)))
                                                                v29 += (v29 & v23)
                                                                var v60 float64 = ((-2.25 + float64(((v2 - v4) + v4))) + 3.14159)
                                                                _ = len(v5)
                                                                _ = v60
                                                            }
                                                        default:
                                                            v29 = (v29 ^ 128)
                                                        }
                                                        if ("\n\t" >= ((("foo bar" + ("hello" + ((("hello" + "a") + "") + ("" + ("\n\t" + ("\n\t" + ("\n\t" + "\n\t"))))))) + ("\n\t" + "foo bar")) + ("" + ("\n\t" + "hello")))) {
                                                            v28 = genFn1(2147483648)
                                                            v3[714] = (^(fi27))
                                                        } else {
                                                            v6 -= v6
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                v5 = append(v5, 562)
                                                            case "a":
                                                                _ = len(v5)
                                                            case "hello":
                                                                v1 -= v1
                                                            default:
                                                                var v61 uint64 = 2
                                                                _ = v61
                                                            }
                                                        }
                                                        if false {
                                                            var v62 float32 = 0.0
                                                            var v63 uint16 = 255
                                                            v23 -= genFn1(42)
                                                            _ = v62
                                                            _ = v63
                                                        } else {
                                                            v28 = (fi22 >> ((282) & 63))
                                                        }
                                                    } else {
                                                        v2 = 1.5
                                                    }
                                                case "\n\t":
                                                    v3[42] = genFn1(1)
                                                case "a":
                                                    if true {
                                                        v29 = (-(v23))
                                                        if len(v5) > 0 {
                                                            v5[0] = 0
                                                        }
                                                        var v64 int64 = (v6 & v6)
                                                        if false {
                                                            var v65 uint8 = 95
                                                            _ = len(v3)
                                                            var v66 float32 = ((v2 + v4) / 1.5)
                                                            _ = v65
                                                            _ = v66
                                                        } else {
                                                            var v67 uint = 0
                                                            var v68 bool = true
                                                            switch "a" {
                                                            case "\n\t":
                                                                var v69 uint16 = 255
                                                                _ = v69
                                                            case "a":
                                                                var v70 int16 = (v1 << (((^(v67))) & 63))
                                                                _ = v70
                                                            case "":
                                                                v29 = (fi27 >> ((v67) & 63))
                                                            default:
                                                                if (((!(!v68)) && (v68 && (!(-2147483648 >= 684)))) && v68) {
                                                                    var v71 float64 = 1.5
                                                                    _ = v71
                                                                } else {
                                                                    if (v68 && (v68 || (!(v29 <= 1)))) {
                                                                        if (!(2147483647 != 905)) {
                                                                            var v72 int8 = 100
                                                                            v73 := GenS5{F0: v1, F1: 2147483647}
                                                                            if v68 {
                                                                                switch fi27 {
                                                                                case 4294967295:
                                                                                    var v74 int16 = v1
                                                                                    _ = v74
                                                                                default:
                                                                                    var v75 float32 = (((1.5 * 1.5) * 0.0) * v4)
                                                                                    _ = v75
                                                                                }
                                                                                fmt.Println(v28, v68)
                                                                                var v76 string = ""
                                                                                _ = v76
                                                                            } else {
                                                                                var v77 int32 = 670
                                                                                _ = v77
                                                                            }
                                                                            _ = v72
                                                                            _ = v73
                                                                        } else {
                                                                            if v68 {
                                                                                var v78 int32 = 2147483647
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                                var v79 float32 = v4
                                                                                _ = v78
                                                                                _ = v79
                                                                            } else {
                                                                                var v80 float64 = 0.1
                                                                                _ = len(v3)
                                                                                var v81 uint8 = 2
                                                                                _ = v80
                                                                                _ = v81
                                                                            }
                                                                            v5 = append(v5, 2147483648)
                                                                            if true {
                                                                                v82 := v3[198]
                                                                                _ = v82
                                                                            } else {
                                                                                v83 := GenS1{F0: v29, F1: 128}
                                                                                var v84 string = ""
                                                                                v29 += (fi27 << ((2) & 63))
                                                                                _ = v83
                                                                                _ = v84
                                                                            }
                                                                        }
                                                                        v85 := GenS6{F0: (-(v64)), F1: 300, F2: "foo bar"}
                                                                        if false {
                                                                            var v86 float32 = (v4 + (float32(-2.25) * v2))
                                                                            if v68 {
                                                                                switch 42 {
                                                                                case 32768:
                                                                                    v86 -= v2
                                                                                case 483:
                                                                                    fmt.Println(v1, v64, v2)
                                                                                case -208:
                                                                                    v64 = 1
                                                                                }
                                                                                var v87 uint64 = 1
                                                                                switch (fi22 & 9223372036854775807) {
                                                                                case -2147483648:
                                                                                    v88 := GenS7{F0: 100, F1: 182}
                                                                                    _ = v88
                                                                                default:
                                                                                    v89 := GenS3{F0: -1, F1: v87, F2: 606}
                                                                                    _ = v89
                                                                                }
                                                                                _ = v87
                                                                            } else {
                                                                                var v90 float64 = 1.5
                                                                                _ = v90
                                                                            }
                                                                            v5 = append(v5, 2147483648)
                                                                            _ = v86
                                                                        }
                                                                        if v68 {
                                                                            if (-222 < 0) {
                                                                                var v91 int16 = v1
                                                                                switch ("" + ("a" + ("a" + "\n\t"))) {
                                                                                case "a":
                                                                                    if v68 {
                                                                                        if false {
                                                                                            v92 := GenS3{F0: 42, F1: 2, F2: (v67 - (v67 % (((v67 ^ v67)) | 1)))}
                                                                                            _ = v92
                                                                                        }
                                                                                    } else {
                                                                                        var v93 uint8 = 127
                                                                                        if (!(true || v68)) {
                                                                                            v3[4294967295] = (v28 - v29)
                                                                                            _ = len(v5)
                                                                                        }
                                                                                        v2 = float32((3.14159 * 1e10))
                                                                                        _ = v93
                                                                                    }
                                                                                }
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    var v94 bool = v68
                                                                                    _ = v94
                                                                                }
                                                                                v95 := GenS7{F0: 42, F1: 1}
                                                                                v23 += (v29 >> ((65535) & 63))
                                                                                if v68 {
                                                                                    fmt.Println(v23, v2, fi22)
                                                                                }
                                                                                var v96 uint8 = 2
                                                                                _ = v91
                                                                                _ = v95
                                                                                _ = v96
                                                                            }
                                                                            if (1 > 65535) {
                                                                                v4 -= (v4 * -2.25)
                                                                            }
                                                                            if v68 {
                                                                                _ = len(v5)
                                                                                if (!v68) {
                                                                                    if v68 {
                                                                                        v97 := GenS8{F0: 255}
                                                                                        v98 := GenS9{F0: 32, F1: 2147483647}
                                                                                        if true {
                                                                                            v23 = v28
                                                                                            v99 := v3[4294967295]
                                                                                            var v100 int32 = 2147483647
                                                                                            _ = v100
                                                                                            _ = v99
                                                                                        }
                                                                                        _ = v97
                                                                                        _ = v98
                                                                                    } else {
                                                                                        v23 -= 128
                                                                                        if (400 == uint16(v67)) {
                                                                                            switch -9223372036854775808 {
                                                                                            case 32767:
                                                                                                fmt.Println(v23)
                                                                                            case 4294967295:
                                                                                                v101 := GenS1{F0: (v29 - fi22), F1: 255}
                                                                                                _ = v101
                                                                                            case 416:
                                                                                                v28 += (fi27 - (v28 / ((genFn1(1)) | 1)))
                                                                                            default:
                                                                                                fmt.Println(v4, fi22)
                                                                                                fmt.Println(v1)
                                                                                                if v68 {
                                                                                                    if (v1 != (v1 | 32767)) {
                                                                                                        v64 += 0
                                                                                                    }
                                                                                                    v28 += (fi27 / ((v23) | 1))
                                                                                                    v67 = v67
                                                                                                    var v102 uint8 = 0
                                                                                                    var v103 uint32 = 1
                                                                                                    _ = v102
                                                                                                    _ = v103
                                                                                                }
                                                                                                fmt.Println(v64)
                                                                                                if (-104 < -1) {
                                                                                                    fmt.Println(fi27)
                                                                                                }
                                                                                                v67 += v67
                                                                                                var v104 uint8 = 42
                                                                                                _ = v104
                                                                                            }
                                                                                        } else {
                                                                                            if (int(v64) < genFn1(65535)) {
                                                                                                var v105 int8 = 127
                                                                                                var v106 uint8 = 163
                                                                                                v64 = v64
                                                                                                _ = v105
                                                                                                _ = v106
                                                                                            } else {
                                                                                                var v107 uint16 = 42
                                                                                                var v108 bool = (150 != 0)
                                                                                                if ((^(v107)) >= (v107 & (^(v107)))) {
                                                                                                    v107 = v107
                                                                                                    switch -902 {
                                                                                                    case -2147483648:
                                                                                                        if (1 < 0) {
                                                                                                            var v109 int16 = (v1 * -760)
                                                                                                            fmt.Println(v1)
                                                                                                            v4 = ((3.14159 - v2) + 0.1)
                                                                                                            if v108 {
                                                                                                                var v110 int16 = genFn0()
                                                                                                                var v111 float64 = ((float64(float32((1e10 * (-2.25 - 0.1)))) + 0.0) - 1e10)
                                                                                                                _ = v110
                                                                                                                _ = v111
                                                                                                            }
                                                                                                            _ = v109
                                                                                                        }
                                                                                                    case 1:
                                                                                                        var v112 uint = v67
                                                                                                        _ = v112
                                                                                                    case 967:
                                                                                                        v4 = v2
                                                                                                    default:
                                                                                                        v3[2147483648] = (fi22 >> (((v67 << ((v67) & 63))) & 63))
                                                                                                    }
                                                                                                } else {
                                                                                                    v113 := GenS3{F0: -75, F1: 174, F2: (v67 >> (((v67 >> ((v67) & 63))) & 63))}
                                                                                                    var v114 float64 = -2.25
                                                                                                    v29 = (-(fi27))
                                                                                                    _ = v113
                                                                                                    _ = v114
                                                                                                }
                                                                                                switch "foo bar" {
                                                                                                case "hello":
                                                                                                    v115 := GenS4{F0: -814, F1: 376, F2: 0}
                                                                                                    _ = v115
                                                                                                case "a":
                                                                                                    _ = len(v5)
                                                                                                case "\n\t":
                                                                                                    switch v29 {
                                                                                                    case 128:
                                                                                                        if (v68 || v68) {
                                                                                                            v64 -= -89
                                                                                                        } else {
                                                                                                            v67 += (v67 * v67)
                                                                                                            var v116 bool = false
                                                                                                            var v117 uint32 = 2147483647
                                                                                                            v118 := GenS3{F0: 67, F1: 42, F2: (v67 % ((v67) | 1))}
                                                                                                            v119 := GenS3{F0: 107, F1: 0, F2: (^(v67))}
                                                                                                            v64 += v6
                                                                                                            _ = v116
                                                                                                            _ = v117
                                                                                                            _ = v118
                                                                                                            _ = v119
                                                                                                        }
                                                                                                        v120 := GenS10{F0: 444, F1: (fi22 + (v23 * (fi22 * genFn1(2)))), F2: -25}
                                                                                                        fmt.Println(v107, v2)
                                                                                                        _ = v120
                                                                                                    case -176:
                                                                                                        var v121 uint16 = (v107 >> ((uint(v107)) & 63))
                                                                                                        _ = v121
                                                                                                    case -1:
                                                                                                        v67 -= v67
                                                                                                    default:
                                                                                                        var v122 uint32 = 4294967295
                                                                                                        _ = v122
                                                                                                    }
                                                                                                }
                                                                                                _ = v107
                                                                                                _ = v108
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    v23 += (v23 * fi22)
                                                                                } else {
                                                                                    v2 += ((v4 / ((3.14159 + (v2 * v4)) - v4)) + 0.0)
                                                                                    fmt.Println(v64, fi27)
                                                                                    if (13 <= -1) {
                                                                                        switch (-(v23)) {
                                                                                        case -397:
                                                                                            v123 := GenS9{F0: -26, F1: 0}
                                                                                            _ = v123
                                                                                        }
                                                                                        var v124 float32 = v4
                                                                                        _ = v124
                                                                                    }
                                                                                    var v125 int = v23
                                                                                    _ = v125
                                                                                }
                                                                            }
                                                                        } else {
                                                                            var v126 uint32 = 2147483647
                                                                            fmt.Println(fi22)
                                                                            _ = v126
                                                                        }
                                                                        _ = v85
                                                                    }
                                                                }
                                                            }
                                                            if len(v5) > 0 {
                                                                v5[0] = 2147483648
                                                            }
                                                            delete(v3, 2147483647)
                                                            _ = v67
                                                            _ = v68
                                                        }
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        _ = v64
                                                    }
                                                }
                                            }
                                            if (((0.1 / v4) < -2.25) || ((((!false) || false) || true) && ((("\n\t" + "a") == ("\n\t" + "")) || false))) {
                                                if (!((int32(v29) != -154) && (42 >= 0))) {
                                                    v127 := GenS11{F0: 60, F1: (!true)}
                                                    var v128 uint8 = 162
                                                    _ = len(v5)
                                                    _ = v127
                                                    _ = v128
                                                }
                                            }
                                            _ = fi27
                                            _ = v28
                                            _ = v29
                                        }
                                        if (!true) {
                                            switch (v23 ^ (fi22 & fi22)) {
                                            case 9223372036854775807:
                                                v6 = (v6 - v6)
                                            default:
                                                if true {
                                                    delete(v3, 2147483647)
                                                }
                                            }
                                            v2 -= (-2.25 + (float32(1e10) * v2))
                                        } else {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            v1 += (v1 | 42)
                                        }
                                        v1 = (v1 * (v1 << ((314) & 63)))
                                    } else {
                                        delete(v3, 2)
                                        switch 128 {
                                        case 128:
                                            v6 = (v6 | (v6 | 9223372036854775807))
                                        }
                                    }
                                }
                            } else {
                                var v129 uint64 = 4294967295
                                _ = v129
                            }
                            _ = len(v5)
                            var v130 bool = true
                            _ = v130
                        } else {
                            v3[2] = v23
                            fmt.Println(v1)
                            fmt.Println(v4, v6, v23)
                        }
                    } else {
                        var v131 int16 = v1
                        v2 = (((0.0 - (-2.25 * float32(float64(v2)))) - v4) * v2)
                        _ = v131
                    }
                    v132 := GenS12{F0: ("a" + "foo bar"), F1: 127, F2: 661}
                    _ = v132
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            }
            v6 -= v6
        }
        if (-1 >= 914) {
            for fi133 := 0; fi133 < 1; fi133++ {
                var v134 bool = true
                fmt.Println(v134, v134)
                _ = fi133
                _ = v134
            }
            fmt.Println(v6, v4, v23)
        }
        _ = fi22
        _ = v23
    }
    fmt.Println(v1, v4)
    if (!(-403 < (v1 >> ((208) & 63)))) {
        v6 = v6
    }
    var v135 float64 = ((((3.14159 / float64(((v4 / 1e10) + ((v2 - (v4 * v2)) - (v2 - float32((0.0 + 3.14159))))))) + ((1e10 + (3.14159 * (((0.0 - 3.14159) + 3.14159) + 1.5))) / 1.5)) * 3.14159) - 1.5)
    if (!false) {
        switch "\n\t" {
        case "":
            var v136 uint16 = 210
            _ = v136
        case "foo bar":
            var v137 float32 = v4
            _ = v137
        case "a":
            var v138 float64 = 1e10
            _ = v138
        }
    } else {
        fmt.Println(v1, v135)
    }
    v1 -= (^(v1))
    if ("\n\t" < "\n\t") {
        if (!false) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            delete(v3, 0)
        }
    } else {
        if true {
            v135 += (((v135 + v135) / ((-2.25 / (1e10 * 3.14159)) * -2.25)) / float64(1e10))
            var v139 int32 = 1
            v140 := GenS5{F0: v1, F1: 146}
            _ = v139
            _ = v140
        } else {
            var v141 bool = (!true)
            _ = v141
        }
        if true {
            v1 = genFn0()
            if (v1 >= (v1 ^ v1)) {
                delete(v3, 2147483647)
                var v142 int = 128
                _ = v142
            }
        } else {
            for fi143 := 0; fi143 < 5; fi143++ {
                _ = len(v3)
                v144 := GenS13{F0: -36}
                _ = fi143
                _ = v144
            }
            var v145 uint16 = 255
            _ = v145
        }
    }
    _ = len(v5)
    for fi146 := 0; fi146 < 6; fi146++ {
        var v147 float32 = (v4 + v2)
        _ = fi146
        _ = v147
    }
    v148 := GenS14{F0: 2147483647, F1: 1, F2: (v6 << ((65535) & 63))}
    fmt.Println(v4)
    var v149 float32 = v2
    fmt.Println(v6, v149, v149)
    var v150 string = ("hello" + "\n\t")
    fmt.Println(v4)
    if (false && (false || (0 <= 42))) {
        var v151 uint = 255
        switch -9223372036854775808 {
        case 836:
            var v152 int32 = -1
            _ = v152
        case 0:
            fmt.Println(v6, v149, v135)
        case -461:
            var v153 int16 = (v1 | (^(v1)))
            _ = v153
        default:
            v5 = append(v5, 184)
        }
        for fi154 := 0; fi154 < 11; fi154++ {
            for fi155 := 0; fi155 < 2; fi155++ {
                var v156 int8 = -38
                _ = fi155
                _ = v156
            }
            _ = fi154
        }
        _ = v151
    }
    v5 = append(v5, 1)
    fmt.Println("v1:", v1)
    fmt.Println("v135:", v135)
    fmt.Println("v149:", v149)
    fmt.Println("v150:", v150)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v135
    _ = v148
    _ = v149
    _ = v150
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
