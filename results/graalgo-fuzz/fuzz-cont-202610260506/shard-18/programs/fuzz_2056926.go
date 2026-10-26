package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int16; F2 uint8 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 uint32; F2 uint16 }
type GenS7 struct { F0 uint32; F1 int64; F2 int64 }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS10 struct { F0 uint16; F1 int8; F2 int16 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint64; F1 uint }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int16; F1 int16 }
type GenS15 struct { F0 int8; F1 bool; F2 string }
type GenS16 struct { F0 int64; F1 int16 }
func genFn0() uint16 {
    var v5 uint8 = 127
    v5 -= (v5 >> ((1) & 63))
    _ = v5
    return 0
}
func genFn1(p0 int32, p1 int32, p2 uint64) uint {
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint = 26
        _ = fi6
        _ = v7
    }
    var v8 int = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 4294967295
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]string)
    var v3 int16 = 42
    var v4 float32 = 0.1
    if ((0.1 / (-2.25 - 0.1)) > 3.14159) {
        if false {
            v4 = ((float32((float64((1e10 + v4)) - 3.14159)) * (v4 / 1e10)) / v4)
            switch 2147483647 {
            case 1:
                for fi9 := 0; fi9 < 20; fi9++ {
                    v10 := GenS1{F0: 128, F1: 2147483647}
                    v11 := GenS1{F0: 2, F1: 42}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            default:
                var v12 int64 = -1
                _ = v12
            }
        } else {
            v2[42] = ("foo bar" + (("foo bar" + "a") + ""))
            v13 := GenS1{F0: 127, F1: -998}
            _ = v13
        }
    }
    fmt.Println(v4, v3)
    v14 := GenS1{F0: 200, F1: -2147483648}
    if true {
        var v15 uint64 = 18446744073709551615
        if true {
            fmt.Println(v15)
        }
        var v16 int64 = 9223372036854775807
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 13; fi17++ {
            delete(v2, fi17)
            switch "" {
            case "\n\t":
                fmt.Println(v3, fi17, fi17)
            }
            _ = fi17
        }
        delete(v2, 110)
        v18 := GenS2{F0: 90}
        _ = v18
    }
    v19 := []int{-783, 42}
    if false {
        v3 = v3
        for fi20 := 0; fi20 < 4; fi20++ {
            if false {
                for fi21 := 0; fi21 < 7; fi21++ {
                    v3 -= 1000
                    if false {
                        v3 -= v3
                        v4 = 1e10
                        fmt.Println(fi20, fi21)
                    } else {
                        v4 -= 1.5
                        v22 := GenS2{F0: 255}
                        var v23 uint16 = 0
                        _ = v22
                        _ = v23
                    }
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                    _ = fi21
                }
                fmt.Println(v4, fi20)
                if (-1 == -234) {
                    switch 2147483647 {
                    case 496:
                        if false {
                            _ = len(v19)
                            v3 += (-(v3))
                        } else {
                            v4 += float32((3.14159 + 3.14159))
                            fmt.Println(fi20, v3, v3)
                            if (100 <= -50) {
                                v3 -= 768
                                if (65535 != 65535) {
                                    v4 -= 3.14159
                                }
                            }
                        }
                    case -97:
                        v4 -= 0.1
                    case 838:
                        v24 := GenS1{F0: 5, F1: 42}
                        _ = v24
                    default:
                        if len(v19) > 0 {
                            v19[0] = 579
                        }
                    }
                } else {
                    switch fi20 {
                    case 2147483647:
                        if (!(124 < genFn1(int32(fi20), -1, 1))) {
                            fmt.Println(v3)
                            var v25 uint8 = 42
                            _ = v25
                        } else {
                            switch (("foo bar" + "") + "\n\t") {
                            case "\n\t":
                                v3 -= v3
                            case "hello":
                                for fi26 := 0; fi26 < 18; fi26++ {
                                    var v27 int64 = 42
                                    _ = fi26
                                    _ = v27
                                }
                            case "a":
                                if true {
                                    var v28 int64 = -9223372036854775808
                                    v3 = v3
                                    _ = v28
                                } else {
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                    v19 = append(v19, (fi20 / (((fi20 << ((genFn1(2147483647, -2147483648, 0)) & 63))) | 1)))
                                }
                            }
                        }
                    case 0:
                        switch ("hello" + ("foo bar" + "hello")) {
                        case "a":
                            if (!true) {
                                if (!false) {
                                    v29 := v2[128]
                                    _ = v29
                                }
                            }
                        case "foo bar":
                            v3 += v3
                        }
                    case 1:
                        for fi30 := 0; fi30 < 12; fi30++ {
                            fmt.Println(fi20, v3, fi20)
                            _ = fi30
                        }
                    default:
                        delete(v2, fi20)
                    }
                }
            }
            _ = fi20
        }
    } else {
        if (!(true && true)) {
            var v31 uint64 = 0
            _ = v31
        } else {
            var v32 uint = 2
            _ = v32
        }
        for fi33 := 0; fi33 < 9; fi33++ {
            v4 = (v4 + (v4 + (v4 / 1e10)))
            v3 = (v3 | v3)
            _ = fi33
        }
        switch -355 {
        case 2147483647:
            for fi34 := 0; fi34 < 4; fi34++ {
                if false {
                    if false {
                        for fi35 := 0; fi35 < 11; fi35++ {
                            if (v4 > (-2.25 + ((v4 / ((((1e10 + v4) / 1e10) - (v4 / 0.1)) * float32(0.0))) / ((v4 + (1e10 / 1e10)) / v4)))) {
                                v36 := GenS2{F0: 130}
                                if (2 <= 2) {
                                    var v37 int = 128
                                    _ = v37
                                } else {
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                    v38 := GenS3{F0: 2, F1: (v3 % ((v3) | 1)), F2: 2}
                                    var v39 uint = 1
                                    _ = v38
                                    _ = v39
                                }
                                v40 := GenS4{F0: 128, F1: -36}
                                _ = v36
                                _ = v40
                            }
                            _ = fi35
                        }
                        if (1.5 < 1e10) {
                            var v41 int64 = 473
                            v41 += v41
                            v3 -= (v3 ^ (-(v3)))
                            _ = v41
                        } else {
                            var v42 uint = genFn1(1, 996, 42)
                            if (true || true) {
                                v43 := GenS5{F0: 42}
                                v4 = (v4 + (((v4 * v4) + 0.0) / ((1e10 - (v4 - v4)) / v4)))
                                _ = v43
                            }
                            for fi44 := 0; fi44 < 10; fi44++ {
                                switch "" {
                                case "\n\t":
                                    var v45 int32 = 0
                                    _ = v45
                                case "a":
                                    var v46 uint = (^(v42))
                                    _ = v46
                                case "hello":
                                    if true {
                                        delete(v2, (fi44 * (fi34 >> ((v42) & 63))))
                                        if false {
                                            var v47 float32 = float32(0.1)
                                            v48 := GenS6{F0: (v42 + 65535), F1: 153, F2: genFn0()}
                                            _ = v47
                                            _ = v48
                                        } else {
                                            if true {
                                                if (!false) {
                                                    if true {
                                                        if true {
                                                            v2[fi44] = "a"
                                                        } else {
                                                            var v49 uint32 = 934
                                                            v3 = (^(v3))
                                                            _ = v49
                                                        }
                                                        var v50 bool = true
                                                        var v51 uint32 = 42
                                                        _ = v50
                                                        _ = v51
                                                    } else {
                                                        var v52 uint8 = 42
                                                        fmt.Println(v3)
                                                        _ = v52
                                                    }
                                                    v1[4294967295] = ((0 >= -10) || true)
                                                } else {
                                                    v53 := v2[(^(fi34))]
                                                    if (1 > 42) {
                                                        if (!(244 != 255)) {
                                                            if (true && (-168 >= fi34)) {
                                                                v54 := GenS5{F0: 127}
                                                                _ = v54
                                                            }
                                                            if ((false && true) || true) {
                                                                var v55 bool = true
                                                                v3 = (v3 * 42)
                                                                var v56 uint8 = 1
                                                                _ = v55
                                                                _ = v56
                                                            }
                                                            v4 = v4
                                                            v4 += (v4 - v4)
                                                        }
                                                        if (!(v53 >= (v53 + ("\n\t" + "a")))) {
                                                            _ = len(v19)
                                                            v53 = (v53 + (v53 + ("\n\t" + v53)))
                                                            _ = len(v1)
                                                        } else {
                                                            var v57 int16 = 32767
                                                            switch fi44 {
                                                            case 42:
                                                                var v58 int = -13
                                                                var v59 int8 = 127
                                                                _ = v58
                                                                _ = v59
                                                            }
                                                            if false {
                                                                switch (fi34 & (fi44 ^ (fi44 >> ((1) & 63)))) {
                                                                case 0:
                                                                    _ = len(v19)
                                                                default:
                                                                    switch v53 {
                                                                    case "":
                                                                        v60 := GenS4{F0: (^(fi44)), F1: 1}
                                                                        _ = v60
                                                                    case "hello":
                                                                        v61 := v1[911]
                                                                        _ = v61
                                                                    default:
                                                                        v42 = (^(v42))
                                                                    }
                                                                }
                                                                v62 := GenS5{F0: 119}
                                                                _ = v62
                                                            } else {
                                                                v63 := GenS7{F0: 1, F1: -390, F2: -9223372036854775808}
                                                                _ = v63
                                                            }
                                                            _ = v57
                                                        }
                                                        v64 := GenS8{F0: 449, F1: fi44}
                                                        _ = v64
                                                    } else {
                                                        if true {
                                                            var v65 uint = (v42 | (v42 & (^(v42))))
                                                            _ = v65
                                                        } else {
                                                            var v66 int32 = -804
                                                            fmt.Println(fi44, v42, v53)
                                                            if (!(-128 == -39)) {
                                                                v4 += (float32((float64(((v4 / (v4 / float32(1.5))) * (-2.25 + v4))) / 1e10)) * v4)
                                                            } else {
                                                                if true {
                                                                    if (v66 > 2147483647) {
                                                                        v4 = v4
                                                                        v67 := GenS5{F0: 0}
                                                                        var v68 uint = v42
                                                                        var v69 bool = true
                                                                        _ = v67
                                                                        _ = v68
                                                                        _ = v69
                                                                    }
                                                                }
                                                                var v70 uint16 = 65535
                                                                fmt.Println(fi34, v3, v42)
                                                                _ = v70
                                                            }
                                                            v4 -= v4
                                                            var v71 uint8 = 51
                                                            _ = v66
                                                            _ = v71
                                                        }
                                                        v3 = v3
                                                        v72 := GenS3{F0: (v42 / ((622) | 1)), F1: (^(v3)), F2: 2}
                                                        _ = v72
                                                    }
                                                    if false {
                                                        var v73 int32 = -836
                                                        _ = len(v19)
                                                        _ = v73
                                                    }
                                                    var v74 float64 = 0.0
                                                    v75 := GenS9{F0: -32768, F1: genFn0(), F2: 0}
                                                    if false {
                                                        v3 += (v3 >> ((65535) & 63))
                                                        if (!true) {
                                                            if true {
                                                                v76 := GenS10{F0: 42, F1: 1, F2: v3}
                                                                _ = v76
                                                            }
                                                            if (!(v42 > v42)) {
                                                                var v77 int64 = -1
                                                                if false {
                                                                    if (true && true) {
                                                                        v78 := GenS11{F0: 2}
                                                                        v53 = ""
                                                                        if true {
                                                                            switch fi34 {
                                                                            case 42:
                                                                                var v79 float32 = 1.5
                                                                                _ = v79
                                                                            case 544:
                                                                                if (false && false) {
                                                                                    v53 = v53
                                                                                } else {
                                                                                    var v80 string = v53
                                                                                    _ = v80
                                                                                }
                                                                            }
                                                                        } else {
                                                                            var v81 string = ""
                                                                            v2[(fi44 | -396)] = ((v53 + "\n\t") + v53)
                                                                            _ = v81
                                                                        }
                                                                        _ = v78
                                                                    } else {
                                                                        var v82 int8 = 81
                                                                        _ = v82
                                                                    }
                                                                    if len(v19) >= 2 {
                                                                        v19 = v19[1:len(v19)]
                                                                    }
                                                                    if false {
                                                                        v83 := GenS12{F0: 42, F1: v42}
                                                                        _ = v83
                                                                    } else {
                                                                        switch (^(fi34)) {
                                                                        case -414:
                                                                            switch (fi44 >> (((v42 / (((^(v42))) | 1))) & 63)) {
                                                                            case -2147483648:
                                                                                delete(v2, (fi34 & 32767))
                                                                            case 44:
                                                                                var v84 float64 = 1.5
                                                                                _ = v84
                                                                            }
                                                                        case 42:
                                                                            v77 = (-(v77))
                                                                        default:
                                                                            v85 := GenS12{F0: 18446744073709551615, F1: 694}
                                                                            _ = v85
                                                                        }
                                                                        var v86 uint = (v42 ^ 2)
                                                                        _ = v86
                                                                    }
                                                                } else {
                                                                    v74 -= 1.5
                                                                    if len(v19) >= 2 {
                                                                        v19 = v19[1:len(v19)]
                                                                    }
                                                                    v87 := GenS13{F0: 0}
                                                                    _ = v87
                                                                }
                                                                if len(v19) > 0 {
                                                                    v19[0] = 0
                                                                }
                                                                _ = v77
                                                            } else {
                                                                v3 = v3
                                                                v88 := GenS2{F0: 42}
                                                                _ = v88
                                                            }
                                                            if (2147483647 >= 0) {
                                                                var v89 int8 = -8
                                                                v90 := v1[715]
                                                                _ = v89
                                                                _ = v90
                                                            }
                                                        }
                                                        var v91 string = v53
                                                        _ = v91
                                                    }
                                                    v3 = (v3 + (v3 % ((-1) | 1)))
                                                    switch v53 {
                                                    case "\n\t":
                                                        v53 = (v53 + v53)
                                                    case "hello":
                                                        switch (fi44 >> (((v42 << ((v42) & 63))) & 63)) {
                                                        case -389:
                                                            switch (fi44 + fi34) {
                                                            case 42:
                                                                v19 = append(v19, 73)
                                                                if (false && false) {
                                                                    v1[2147483647] = (uint64(v42) > 18446744073709551615)
                                                                } else {
                                                                    v3 -= (v3 << (((v42 << ((0) & 63))) & 63))
                                                                    v3 += (v3 & (v3 / ((int16(fi44)) | 1)))
                                                                    if (100 >= -14) {
                                                                        v42 -= 226
                                                                        var v92 uint8 = 223
                                                                        _ = v92
                                                                    } else {
                                                                        var v93 int = (fi44 / (((fi44 << ((genFn1(-683, 1, 42)) & 63))) | 1))
                                                                        if (false || ((false && (!(128 <= 255))) && false)) {
                                                                            switch fi44 {
                                                                            case 9223372036854775807:
                                                                                var v94 float32 = ((float32((v74 - 3.14159)) / ((0.1 + float32((v74 / v74))) - (1e10 - float32(1.5)))) / v4)
                                                                                _ = v94
                                                                            case 638:
                                                                                v93 = fi34
                                                                            case 603:
                                                                                switch (fi44 | fi34) {
                                                                                case 42:
                                                                                    v19 = append(v19, (fi34 | (v93 << ((v42) & 63))))
                                                                                case -9223372036854775808:
                                                                                    v95 := v1[2147483648]
                                                                                    _ = v95
                                                                                default:
                                                                                    fmt.Println(v53, v3)
                                                                                }
                                                                            default:
                                                                                fmt.Println(v74, v93, v4)
                                                                            }
                                                                            v96 := GenS13{F0: 4294967295}
                                                                            if ((false && (true && true)) || ((2147483647 >= 4294967295) && true)) {
                                                                                if ((v93 + (fi34 % ((32768) | 1))) != fi34) {
                                                                                    var v97 float64 = (-2.25 * v74)
                                                                                    v4 += (((((v4 / v4) - 1e10) / v4) + (v4 * 1.5)) + 3.14159)
                                                                                    if (128 == 255) {
                                                                                        fmt.Println(v74, v97, fi34)
                                                                                        v4 = (v4 * (3.14159 + (((v4 - v4) * (v4 - (v4 / 1e10))) + v4)))
                                                                                    }
                                                                                    if (-9223372036854775808 >= 1) {
                                                                                        var v98 uint64 = 18446744073709551615
                                                                                        _ = v98
                                                                                    } else {
                                                                                        v74 = v74
                                                                                    }
                                                                                    v99 := GenS14{F0: (^(v3)), F1: v3}
                                                                                    _ = v97
                                                                                    _ = v99
                                                                                }
                                                                                v19 = append(v19, (v93 ^ (fi44 * int(v3))))
                                                                                fmt.Println(v74, v74, v42)
                                                                            }
                                                                            _ = v96
                                                                        } else {
                                                                            var v100 bool = false
                                                                            _ = v100
                                                                        }
                                                                        var v101 int8 = 100
                                                                        _ = v101
                                                                        _ = v93
                                                                    }
                                                                    if len(v19) > 0 {
                                                                        v19[0] = (fi34 - fi44)
                                                                    }
                                                                }
                                                                if true {
                                                                    v42 = (v42 / ((357) | 1))
                                                                    v4 += v4
                                                                    var v102 int = fi34
                                                                    _ = v102
                                                                }
                                                            case 128:
                                                                v74 -= (float64(float32(v74)) - ((v74 * ((1e10 / (v74 + float64(v4))) + -2.25)) * ((v74 - v74) * (v74 * (v74 / v74)))))
                                                            case 127:
                                                                v103 := GenS2{F0: 128}
                                                                _ = v103
                                                            }
                                                        }
                                                    case "a":
                                                        var v104 string = v53
                                                        _ = v104
                                                    default:
                                                        v105 := v1[4294967295]
                                                        _ = v105
                                                    }
                                                    switch v53 {
                                                    case "hello":
                                                        v53 = "foo bar"
                                                        fmt.Println(fi34, v3)
                                                        switch fi34 {
                                                        case 42:
                                                            if ((611 > 1) && true) {
                                                                if true {
                                                                    switch 32767 {
                                                                    case 794:
                                                                        v106 := GenS6{F0: genFn1(-101, 1, 18446744073709551615), F1: 42, F2: 32767}
                                                                        _ = v106
                                                                    }
                                                                } else {
                                                                    var v107 string = v53
                                                                    _ = v107
                                                                }
                                                            }
                                                            if ((77 == 1) && false) {
                                                                v74 = v74
                                                                var v108 int = 477
                                                                _ = v108
                                                            }
                                                            switch fi44 {
                                                            case 128:
                                                                v3 += v3
                                                            case 32768:
                                                                var v109 uint = (^(v42))
                                                                if (false || true) {
                                                                    v110 := GenS11{F0: 0}
                                                                    _ = v110
                                                                } else {
                                                                    fmt.Println(v74, v74)
                                                                    fmt.Println(v74, fi44)
                                                                }
                                                                _ = v109
                                                            }
                                                        case -1:
                                                            delete(v2, fi44)
                                                        case 728:
                                                            if (2147483648 != 2) {
                                                                v3 += (v3 ^ (^(v3)))
                                                            }
                                                        default:
                                                            if (!(true && true)) {
                                                                v53 = ""
                                                                var v111 int = fi44
                                                                _ = v111
                                                            }
                                                        }
                                                    case "\n\t":
                                                        v4 = v4
                                                    }
                                                    v4 -= 3.14159
                                                    _ = v53
                                                    _ = v74
                                                    _ = v75
                                                }
                                            } else {
                                                v4 = (float32(3.14159) / 1.0)
                                            }
                                        }
                                        _ = len(v2)
                                    }
                                default:
                                    var v112 int = (fi44 & fi44)
                                    _ = v112
                                }
                                v113 := GenS13{F0: (v42 * 42)}
                                _ = fi44
                                _ = v113
                            }
                            _ = v42
                        }
                    } else {
                        v114 := GenS15{F0: -128, F1: (!false), F2: "foo bar"}
                        var v115 float64 = (((0.1 / 0.1) / 0.1) / 1e10)
                        _ = v114
                        _ = v115
                    }
                    v116 := GenS16{F0: -1, F1: (v3 - v3)}
                    _ = v116
                }
                switch -491 {
                case 32768:
                    for fi117 := 0; fi117 < 2; fi117++ {
                        var v118 uint16 = genFn0()
                        _ = fi117
                        _ = v118
                    }
                case 2147483647:
                    switch "\n\t" {
                    case "foo bar":
                        var v119 bool = (804 >= 739)
                        _ = v119
                    case "a":
                        v120 := GenS7{F0: 2147483647, F1: -488, F2: 0}
                        _ = v120
                    case "":
                        for fi121 := 0; fi121 < 8; fi121++ {
                            v3 += 1000
                            _ = fi121
                        }
                    default:
                        fmt.Println(v3)
                    }
                case 32767:
                    var v122 int16 = v3
                    _ = v122
                }
                _ = fi34
            }
        case 4294967295:
            if false {
                v4 -= 1e10
                var v123 uint8 = 0
                v4 = v4
                _ = v123
            }
        default:
            if ((false && (!(false || (-2.25 <= -2.25)))) || true) {
                switch 2147483647 {
                case 17:
                    v4 += (float32(-2.25) + v4)
                default:
                    if (255 != 224) {
                        delete(v1, 42)
                        var v124 int = 32767
                        _ = v124
                    }
                }
                var v125 uint32 = 42
                var v126 int8 = 100
                _ = v125
                _ = v126
            }
        }
    }
    v127 := v1[2147483648]
    switch 884 {
    case -89:
        v4 -= 0.0
    case 9223372036854775807:
        switch -838 {
        case 42:
            var v128 int8 = -101
            _ = v128
        case -33:
            switch "a" {
            case "\n\t":
                for fi129 := 0; fi129 < 2; fi129++ {
                    for fi130 := 0; fi130 < 18; fi130++ {
                        var v131 int16 = (^(v3))
                        fmt.Println(v4)
                        _ = fi130
                        _ = v131
                    }
                    v4 = 0.0
                    _ = fi129
                }
            default:
                if (("" <= "hello") && true) {
                    v3 += v3
                    fmt.Println(v4, v4)
                }
            }
        default:
            v132 := GenS14{F0: (v3 | v3), F1: v3}
            _ = v132
        }
    case 128:
        v4 -= float32(1e10)
    }
    _ = len(v19)
    _ = len(v19)
    v3 = -1
    if (244 <= 0) {
        if v127 {
            v4 -= v4
            var v133 int64 = 9223372036854775807
            _ = v133
        }
        var v134 string = "a"
        _ = v134
    }
    fmt.Println(v127, v127, v4)
    var v135 int16 = v3
    v135 -= -32768
    switch "foo bar" {
    case "a":
        if (0 > 42) {
            v127 = false
        } else {
            v3 -= (v135 % ((1000) | 1))
            fmt.Println(v127)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v127:", v127)
    fmt.Println("v135:", v135)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v127
    _ = v135
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
