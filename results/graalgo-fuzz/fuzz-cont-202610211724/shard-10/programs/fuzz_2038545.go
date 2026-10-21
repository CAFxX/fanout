package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 int64; F1 uint; F2 int32 }
type GenS8 struct { F0 uint8; F1 uint32 }
type GenS9 struct { F0 int32; F1 int64; F2 uint }
type GenS10 struct { F0 int8; F1 uint64; F2 uint }
type GenS11 struct { F0 uint; F1 uint64 }
type GenS12 struct { F0 uint64; F1 int64 }
type GenS13 struct { F0 int16; F1 bool }
type GenS14 struct { F0 uint32; F1 int16 }
type GenS15 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS16 struct { F0 uint8; F1 int }
type GenS17 struct { F0 string }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 int16; F1 int }
type GenS20 struct { F0 int32 }
type GenS21 struct { F0 int8; F1 int32 }
type GenS22 struct { F0 int8; F1 int32 }
type GenS23 struct { F0 bool }
type GenS24 struct { F0 int16 }
func genFn0(p0 uint64) string {
    var v4 uint = 589
    switch "hello" {
    case "":
        fmt.Println(v4)
    case "a":
        v4 -= (v4 % ((540) | 1))
    }
    var v5 int = 127
    v5 = (v5 * -2147483648)
    _ = p0
    _ = v4
    _ = v5
    return "hello"
}
func genFn1(p0 uint) int64 {
    fmt.Println(p0)
    if true {
        if (1 > 2) {
            for fi6 := 0; fi6 < 16; fi6++ {
                var v7 uint = (p0 | (p0 >> (((p0 >> ((p0) & 63))) & 63)))
                var v8 uint32 = 680
                if false {
                    for fi9 := 0; fi9 < 17; fi9++ {
                        if false {
                            if false {
                                if ((false && ("foo bar" <= "")) || true) {
                                    v10 := GenS1{F0: 1, F1: -2147483648, F2: 346}
                                    p0 = v7
                                    if false {
                                        v7 = (p0 >> ((417) & 63))
                                        fmt.Println(v8, v7, v7)
                                        v11 := []int{fi9, (^(fi9))}
                                        _ = v11
                                    }
                                    _ = v10
                                } else {
                                    fmt.Println(v7)
                                    if true {
                                        var v12 int8 = 100
                                        v13 := []int{(fi6 - (fi9 - fi6)), (fi6 & (fi6 | fi9))}
                                        var v14 uint = (p0 ^ (p0 << (((p0 << ((v7) & 63))) & 63)))
                                        _ = v12
                                        _ = v13
                                        _ = v14
                                    }
                                    if true {
                                        if true {
                                            if (((!(false || true)) || ((!false) && (1 < 0))) && (-52 >= -1)) {
                                                var v15 uint16 = 292
                                                if (fi6 < fi9) {
                                                    var v16 int16 = -258
                                                    if ((p0 << ((65535) & 63)) != (v7 * 1)) {
                                                        v17 := GenS1{F0: 128, F1: 42, F2: v7}
                                                        if false {
                                                            if (-2.25 < 0.0) {
                                                                switch "a" {
                                                                case "":
                                                                    fmt.Println(v16, v15)
                                                                default:
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        var v18 bool = true
                                                                        _ = v18
                                                                    case "":
                                                                        v7 += (p0 ^ 1)
                                                                    }
                                                                }
                                                            } else {
                                                                if false {
                                                                    v19 := []int{fi9, 0, fi9}
                                                                    v7 -= v7
                                                                    switch fi6 {
                                                                    case -38:
                                                                        v15 -= v15
                                                                    case 964:
                                                                        p0 += uint(v8)
                                                                        var v20 int64 = -571
                                                                        _ = v20
                                                                    default:
                                                                        if true {
                                                                            var v21 float32 = 1e10
                                                                            _ = v21
                                                                        } else {
                                                                            switch (-(fi6)) {
                                                                            case 1:
                                                                                switch ("" + "") {
                                                                                case "":
                                                                                    var v22 uint32 = v8
                                                                                    _ = v22
                                                                                case "a":
                                                                                    switch 1 {
                                                                                    case 4294967295:
                                                                                        switch ("foo bar" + "\n\t") {
                                                                                        case "a":
                                                                                            v7 -= (^(v7))
                                                                                        case "\n\t":
                                                                                            if true {
                                                                                                v7 += v7
                                                                                            } else {
                                                                                                _ = len(v19)
                                                                                                if len(v19) > 0 {
                                                                                                    v19[0] = (fi9 | (-(fi6)))
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            fmt.Println(fi9, v15, fi6)
                                                                                        }
                                                                                    case -787:
                                                                                        v16 = 1
                                                                                    case 702:
                                                                                        v23 := GenS2{F0: fi6, F1: -597}
                                                                                        _ = v23
                                                                                    }
                                                                                case "hello":
                                                                                    v19 = append(v19, fi9)
                                                                                default:
                                                                                    if ((((v8 << ((v7) & 63)) >= 0) || (v7 == (v7 % (((v7 << ((p0) & 63))) | 1)))) && false) {
                                                                                        if (!(42 >= 35)) {
                                                                                            v24 := GenS2{F0: (fi9 & (fi6 ^ fi9)), F1: -2147483648}
                                                                                            _ = v24
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case 4294967295:
                                                                                switch "hello" {
                                                                                case "a":
                                                                                    v8 -= v8
                                                                                case "hello":
                                                                                    if true {
                                                                                        var v25 int16 = (v16 - (v16 ^ v16))
                                                                                        _ = v25
                                                                                    }
                                                                                    switch -9223372036854775808 {
                                                                                    case -1:
                                                                                        var v26 uint8 = uint8(v15)
                                                                                        _ = v26
                                                                                    case -641:
                                                                                        var v27 uint32 = v8
                                                                                        _ = v27
                                                                                    case -9223372036854775808:
                                                                                        v28 := GenS2{F0: (^(fi6)), F1: -2147483648}
                                                                                        _ = v28
                                                                                    default:
                                                                                        v8 -= v8
                                                                                    }
                                                                                    var v29 int = (fi9 + (fi9 << ((p0) & 63)))
                                                                                    v29 -= (fi6 | (fi9 % ((fi9) | 1)))
                                                                                    if len(v19) > 0 {
                                                                                        v19[0] = v29
                                                                                    }
                                                                                    _ = v29
                                                                                }
                                                                            }
                                                                            if true {
                                                                                var v30 int16 = v16
                                                                                v8 += v8
                                                                                _ = v30
                                                                            }
                                                                            switch "" {
                                                                            case "a":
                                                                                v31 := GenS1{F0: 255, F1: -2147483648, F2: 4294967295}
                                                                                _ = v31
                                                                            case "":
                                                                                switch fi9 {
                                                                                case 128:
                                                                                    if false {
                                                                                        v7 -= p0
                                                                                        fmt.Println(v15)
                                                                                        v16 -= v16
                                                                                    } else {
                                                                                        switch ("hello" + "hello") {
                                                                                        case "hello":
                                                                                            v7 = 0
                                                                                        case "":
                                                                                            v15 = 1
                                                                                        case "a":
                                                                                            if len(v19) > 0 {
                                                                                                v19[0] = (fi6 + -516)
                                                                                            }
                                                                                        default:
                                                                                            switch fi9 {
                                                                                            case -244:
                                                                                                v16 = (v16 + v16)
                                                                                            case 128:
                                                                                                if (0 >= -1) {
                                                                                                    v32 := GenS3{F0: 209, F1: uint8(p0)}
                                                                                                    if (1e10 == 0.0) {
                                                                                                        var v33 uint32 = v8
                                                                                                        var v34 int32 = 959
                                                                                                        _ = v33
                                                                                                        _ = v34
                                                                                                    }
                                                                                                    _ = v32
                                                                                                } else {
                                                                                                    v16 = v16
                                                                                                    if len(v19) >= 2 {
                                                                                                        v19 = v19[1:len(v19)]
                                                                                                    }
                                                                                                    var v35 int8 = -128
                                                                                                    fmt.Println(v8)
                                                                                                    v16 = (v16 / ((v16) | 1))
                                                                                                    _ = v35
                                                                                                }
                                                                                            default:
                                                                                                if (!((127 == 190) && false)) {
                                                                                                    v19 = append(v19, fi9)
                                                                                                    var v36 uint32 = 15
                                                                                                    _ = v36
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                case -669:
                                                                                    switch fi9 {
                                                                                    case 128:
                                                                                        v15 -= (v15 | (v15 % (((v15 >> (((^(p0))) & 63))) | 1)))
                                                                                    case 290:
                                                                                        fmt.Println(v16, p0)
                                                                                    default:
                                                                                        var v37 int32 = 42
                                                                                        _ = v37
                                                                                    }
                                                                                default:
                                                                                    var v38 int32 = -237
                                                                                    _ = v38
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v19
                                                                } else {
                                                                    var v39 bool = false
                                                                    if (!false) {
                                                                        v40 := GenS2{F0: fi6, F1: 2147483647}
                                                                        _ = v40
                                                                    }
                                                                    switch (fi9 ^ -2147483648) {
                                                                    case 128:
                                                                        v41 := []int{fi6, (fi9 | fi9)}
                                                                        _ = v41
                                                                    }
                                                                    if v39 {
                                                                        v42 := []int{(fi6 >> ((2) & 63))}
                                                                        switch fi6 {
                                                                        case 9223372036854775807:
                                                                            switch "a" {
                                                                            case "\n\t":
                                                                                fmt.Println(v7, v39, fi6)
                                                                            case "foo bar":
                                                                                fmt.Println(v15, v7)
                                                                                var v43 uint32 = v8
                                                                                _ = v43
                                                                            }
                                                                        case -9223372036854775808:
                                                                            switch 42 {
                                                                            case -1:
                                                                                var v44 int16 = v16
                                                                                _ = v44
                                                                            case 4294967295:
                                                                                if len(v42) >= 2 {
                                                                                    v42 = v42[1:len(v42)]
                                                                                }
                                                                            case 42:
                                                                                if len(v42) >= 2 {
                                                                                    v42 = v42[1:len(v42)]
                                                                                }
                                                                            default:
                                                                                v7 = v7
                                                                            }
                                                                        default:
                                                                            var v45 uint16 = v15
                                                                            _ = v45
                                                                        }
                                                                        v46 := GenS4{F0: 255}
                                                                        _ = v42
                                                                        _ = v46
                                                                    }
                                                                    _ = v39
                                                                }
                                                            }
                                                        } else {
                                                            p0 += p0
                                                            v47 := GenS5{F0: 255}
                                                            v48 := []int{(fi6 << (((v7 - 42)) & 63)), (-(fi9))}
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                        _ = v17
                                                    } else {
                                                        v49 := GenS3{F0: 9, F1: 224}
                                                        v7 = (v7 + uint(v15))
                                                        _ = v49
                                                    }
                                                    p0 = (p0 * v7)
                                                    v50 := GenS6{F0: (((("\n\t" + "\n\t") + ("" + "hello")) + "") + "a"), F1: v16}
                                                    v51 := GenS3{F0: 2, F1: 1}
                                                    var v52 int = (fi9 / ((fi9) | 1))
                                                    p0 = 42
                                                    _ = v16
                                                    _ = v50
                                                    _ = v51
                                                    _ = v52
                                                }
                                                v53 := GenS7{F0: 9223372036854775807, F1: (^(p0)), F2: -2147483648}
                                                _ = v15
                                                _ = v53
                                            } else {
                                                if (312 != 9223372036854775807) {
                                                    v54 := GenS8{F0: 2, F1: 42}
                                                    v7 -= (v7 - (^(v7)))
                                                    _ = v54
                                                } else {
                                                    var v55 float64 = 0.0
                                                    var v56 uint = p0
                                                    _ = v55
                                                    _ = v56
                                                }
                                                if true {
                                                    var v57 float32 = 1.5
                                                    v57 = 0.0
                                                    _ = v57
                                                }
                                            }
                                            if true {
                                                if true {
                                                    v58 := []int{(fi6 % (((fi9 & fi6)) | 1)), 2147483647}
                                                    _ = v58
                                                }
                                            } else {
                                                v7 += (v7 << (((p0 % ((p0) | 1))) & 63))
                                                v59 := []int{(fi9 << ((255) & 63))}
                                                _ = v59
                                            }
                                            if false {
                                                v7 = (p0 | 228)
                                                p0 -= (^(p0))
                                                v7 = v7
                                                v60 := []int{}
                                                _ = v60
                                            } else {
                                                if false {
                                                    switch "foo bar" {
                                                    case "":
                                                        switch (("" + "\n\t") + "foo bar") {
                                                        case "hello":
                                                            fmt.Println(v7, v7)
                                                        case "\n\t":
                                                            switch (fi6 | -1) {
                                                            case 157:
                                                                switch int(p0) {
                                                                case -829:
                                                                    v61 := GenS9{F0: -2147483648, F1: 0, F2: (v7 & p0)}
                                                                    _ = v61
                                                                case -1:
                                                                    v8 = (v8 + 308)
                                                                case 42:
                                                                    p0 = (v7 - v7)
                                                                default:
                                                                    if (100 <= 1) {
                                                                        if ((1.5 + 1e10) <= 1.5) {
                                                                            fmt.Println(p0)
                                                                        }
                                                                        v7 = v7
                                                                        if true {
                                                                            var v62 int16 = -1
                                                                            var v63 float64 = (float64((0.0 + (3.14159 - 1.5))) / 1.0)
                                                                            var v64 int = -865
                                                                            _ = v62
                                                                            _ = v63
                                                                            _ = v64
                                                                        } else {
                                                                            fmt.Println(v8)
                                                                            p0 -= v7
                                                                            v65 := []int{(fi6 ^ (fi9 % ((fi6) | 1))), 2147483647, fi6, fi6}
                                                                            if true {
                                                                                if len(v65) > 0 {
                                                                                    v65[0] = fi9
                                                                                }
                                                                                var v66 uint8 = 119
                                                                                v7 = (v7 % (((p0 ^ (v7 / ((65535) | 1)))) | 1))
                                                                                _ = v66
                                                                            } else {
                                                                                if ((v7 << (((p0 % (((^(v7))) | 1))) & 63)) < (^(v7))) {
                                                                                    if false {
                                                                                        v7 -= (^(p0))
                                                                                        if len(v65) >= 2 {
                                                                                            v65 = v65[1:len(v65)]
                                                                                        }
                                                                                        var v67 uint8 = 42
                                                                                        _ = v67
                                                                                    } else {
                                                                                        if true {
                                                                                            var v68 uint8 = 206
                                                                                            var v69 string = ("a" + "a")
                                                                                            _ = v68
                                                                                            _ = v69
                                                                                        }
                                                                                        p0 -= 295
                                                                                        var v70 int8 = -1
                                                                                        var v71 int64 = 903
                                                                                        _ = v70
                                                                                        _ = v71
                                                                                    }
                                                                                    var v72 uint16 = 65535
                                                                                    var v73 float32 = 1.5
                                                                                    v74 := GenS10{F0: 74, F1: 488, F2: 785}
                                                                                    if true {
                                                                                        v75 := GenS11{F0: (p0 >> (((v7 / ((42) | 1))) & 63)), F1: uint64(v8)}
                                                                                        var v76 string = ("a" + (("foo bar" + "a") + "hello"))
                                                                                        switch v76 {
                                                                                        case "a":
                                                                                            v77 := GenS6{F0: v76, F1: -32768}
                                                                                            _ = v77
                                                                                        case "hello":
                                                                                            _ = len(v65)
                                                                                        }
                                                                                        _ = v75
                                                                                        _ = v76
                                                                                    } else {
                                                                                        v78 := GenS4{F0: 42}
                                                                                        var v79 uint32 = (v8 * v8)
                                                                                        _ = v78
                                                                                        _ = v79
                                                                                    }
                                                                                    _ = v72
                                                                                    _ = v73
                                                                                    _ = v74
                                                                                } else {
                                                                                    switch fi6 {
                                                                                    case -739:
                                                                                        v7 += v7
                                                                                    case 611:
                                                                                        v80 := GenS12{F0: 42, F1: 550}
                                                                                        _ = v80
                                                                                    case -640:
                                                                                        var v81 int64 = 0
                                                                                        _ = v81
                                                                                    default:
                                                                                        var v82 float64 = -2.25
                                                                                        _ = v82
                                                                                    }
                                                                                    v65 = append(v65, 4294967295)
                                                                                    p0 += (v7 * (v7 ^ p0))
                                                                                    if true {
                                                                                        var v83 uint32 = (v8 << (((p0 + 65535)) & 63))
                                                                                        var v84 uint32 = 42
                                                                                        var v85 uint8 = 255
                                                                                        _ = v83
                                                                                        _ = v84
                                                                                        _ = v85
                                                                                    } else {
                                                                                        var v86 uint64 = 0
                                                                                        v65 = append(v65, (^(fi9)))
                                                                                        fmt.Println(v7, fi9)
                                                                                        _ = v86
                                                                                    }
                                                                                }
                                                                                var v87 int32 = 914
                                                                                _ = v87
                                                                            }
                                                                            if false {
                                                                                v88 := GenS4{F0: 229}
                                                                                _ = v88
                                                                            }
                                                                            _ = v65
                                                                        }
                                                                    }
                                                                }
                                                                v7 = (p0 % ((v7) | 1))
                                                                v89 := GenS4{F0: 160}
                                                                _ = v89
                                                            default:
                                                                fmt.Println(v7)
                                                            }
                                                        }
                                                        v90 := GenS6{F0: (("foo bar" + (("foo bar" + "hello") + "a")) + ""), F1: -1}
                                                        if true {
                                                            v91 := GenS13{F0: 42, F1: true}
                                                            v92 := GenS14{F0: v8, F1: 42}
                                                            var v93 string = "a"
                                                            _ = v91
                                                            _ = v92
                                                            _ = v93
                                                        }
                                                        _ = v90
                                                    case "hello":
                                                        v94 := []int{}
                                                        _ = v94
                                                    }
                                                    fmt.Println(v8, fi9, v8)
                                                    if false {
                                                        var v95 int64 = -1
                                                        _ = v95
                                                    } else {
                                                        p0 = (p0 & (p0 / ((v7) | 1)))
                                                        switch fi9 {
                                                        case 127:
                                                            v96 := GenS15{F0: 65535, F1: 1, F2: 0}
                                                            _ = v96
                                                        case 32768:
                                                            v97 := GenS16{F0: 118, F1: (fi6 - 128)}
                                                            _ = v97
                                                        default:
                                                            v98 := GenS17{F0: ((("" + ("\n\t" + "foo bar")) + ("foo bar" + "\n\t")) + "a")}
                                                            _ = v98
                                                        }
                                                        if (0 < 127) {
                                                            v99 := []int{127}
                                                            _ = v99
                                                        } else {
                                                            var v100 uint8 = 188
                                                            fmt.Println(v8)
                                                            _ = v100
                                                        }
                                                    }
                                                } else {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        var v101 int = -9223372036854775808
                                                        _ = v101
                                                    default:
                                                        if (42 >= -32768) {
                                                            v8 -= v8
                                                        }
                                                    }
                                                    v7 = (p0 | 2)
                                                }
                                            }
                                        } else {
                                            var v102 int8 = 127
                                            var v103 uint32 = v8
                                            _ = v102
                                            _ = v103
                                        }
                                    } else {
                                        v104 := GenS18{F0: (("\n\t" > "foo bar") && false)}
                                        _ = v104
                                    }
                                    var v105 uint16 = 2
                                    var v106 int8 = -128
                                    var v107 bool = false
                                    _ = v105
                                    _ = v106
                                    _ = v107
                                }
                            }
                        } else {
                            v7 = 255
                            if true {
                                var v108 int8 = -128
                                switch (fi9 >> ((1) & 63)) {
                                case 2147483647:
                                    var v109 int32 = 0
                                    _ = v109
                                case 32768:
                                    v110 := GenS1{F0: 1, F1: -1, F2: v7}
                                    _ = v110
                                default:
                                    switch (fi9 * fi9) {
                                    case -9223372036854775808:
                                        if (4294967295 < 42) {
                                            p0 = p0
                                            fmt.Println(v108, v8)
                                        } else {
                                            if true {
                                                v111 := []int{(fi6 - -169), (^(fi6)), (fi9 | fi9), (fi6 ^ -199)}
                                                if true {
                                                    switch "a" {
                                                    case "":
                                                        var v112 float32 = float32(1e10)
                                                        _ = v112
                                                    case "a":
                                                        var v113 uint = v7
                                                        _ = v113
                                                    }
                                                    var v114 int8 = v108
                                                    _ = v114
                                                }
                                                _ = v111
                                            }
                                            var v115 float64 = -2.25
                                            _ = v115
                                        }
                                    case -744:
                                        v116 := []int{fi6}
                                        _ = v116
                                    case 32768:
                                        v117 := GenS5{F0: (^(p0))}
                                        _ = v117
                                    }
                                }
                                _ = v108
                            } else {
                                v8 -= (^(v8))
                                var v118 bool = (-2147483648 == 528)
                                v8 = v8
                                _ = v118
                            }
                            if (false && ((^(v8)) < (v8 << ((4294967295) & 63)))) {
                                if false {
                                    v8 -= 664
                                    var v119 float64 = float64(1.5)
                                    _ = v119
                                } else {
                                    if true {
                                        var v120 int = fi6
                                        _ = v120
                                    } else {
                                        v7 += p0
                                        var v121 uint64 = 0
                                        v122 := []int{(fi9 ^ fi9)}
                                        _ = v121
                                        _ = v122
                                    }
                                }
                                var v123 int32 = 1
                                v124 := []int{fi9, fi9, -909}
                                _ = v123
                                _ = v124
                            } else {
                                v8 += 42
                                var v125 float64 = (3.14159 + 3.14159)
                                _ = v125
                            }
                        }
                        var v126 int8 = 1
                        _ = fi9
                        _ = v126
                    }
                } else {
                    for fi127 := 0; fi127 < 1; fi127++ {
                        v128 := GenS7{F0: 9223372036854775807, F1: v7, F2: -543}
                        var v129 int32 = -288
                        _ = fi127
                        _ = v128
                        _ = v129
                    }
                }
                _ = fi6
                _ = v7
                _ = v8
            }
            p0 = p0
            v130 := GenS19{F0: 40, F1: 591}
            _ = v130
        } else {
            if (false && (p0 <= (p0 - p0))) {
                p0 = (p0 << ((p0) & 63))
                switch ("\n\t" + "foo bar") {
                case "hello":
                    p0 -= p0
                }
                v131 := GenS18{F0: false}
                _ = v131
            }
            var v132 uint64 = 2
            v133 := []int{1, 127, 725}
            _ = v132
            _ = v133
        }
        v134 := []int{0}
        _ = v134
    } else {
        p0 -= (^(p0))
        for fi135 := 0; fi135 < 1; fi135++ {
            p0 = (^(p0))
            var v136 uint8 = 12
            _ = fi135
            _ = v136
        }
    }
    _ = p0
    return 412
}

func main() {
    var v1 string = "foo bar"
    var v2 string = v1
    v3 := make(map[int]uint32)
    if (32767 != 42) {
        if false {
            v3[9223372036854775807] = 4294967295
            v2 = v2
        }
        if (18446744073709551615 <= 2) {
            for fi137 := 0; fi137 < 5; fi137++ {
                if ((("" + (((v2 + "\n\t") + ("a" + "\n\t")) + (v1 + v2))) + (("foo bar" + "hello") + v2)) < (("foo bar" + "\n\t") + v1)) {
                    switch "" {
                    case "":
                        fmt.Println(v1)
                    case "hello":
                        if false {
                            var v138 uint16 = 2
                            switch (fi137 << ((933) & 63)) {
                            case 1:
                                var v139 uint = 205
                                _ = v139
                            case -2147483648:
                                if (-2.25 <= 1e10) {
                                    v140 := []int{fi137, (fi137 & (fi137 % ((fi137) | 1))), 4294967295}
                                    _ = v140
                                }
                            default:
                                v141 := GenS4{F0: 1}
                                _ = v141
                            }
                            for fi142 := 0; fi142 < 10; fi142++ {
                                fmt.Println(v138, v1)
                                _ = fi142
                            }
                            _ = v138
                        } else {
                            v2 = (v1 + "hello")
                            for fi143 := 0; fi143 < 1; fi143++ {
                                var v144 int8 = 0
                                fmt.Println(fi143)
                                v145 := GenS20{F0: 1}
                                _ = fi143
                                _ = v144
                                _ = v145
                            }
                            var v146 float64 = (3.14159 - 0.0)
                            _ = v146
                        }
                    case "foo bar":
                        v1 = (v1 + "")
                    default:
                        fmt.Println(fi137, v1, v1)
                    }
                } else {
                    switch (fi137 << ((710) & 63)) {
                    case 32767:
                        var v147 int16 = 1
                        _ = v147
                    default:
                        v148 := GenS21{F0: 1, F1: 663}
                        _ = v148
                    }
                }
                _ = fi137
            }
        }
        switch v1 {
        case "\n\t":
            v3[330] = 65535
        case "":
            v149 := GenS5{F0: 4294967295}
            _ = v149
        }
    } else {
        var v150 uint64 = 894
        _ = v150
    }
    fmt.Println(v1)
    if false {
        v3[32768] = 352
        fmt.Println(v2)
    }
    switch (((v1 + v1) + v1) + v1) {
    case "a":
        for fi151 := 0; fi151 < 7; fi151++ {
            _ = len(v3)
            var v152 uint8 = 127
            for fi153 := 0; fi153 < 15; fi153++ {
                var v154 uint64 = 42
                switch v2 {
                case "hello":
                    if true {
                        var v155 uint64 = (v154 + (v154 % ((911) | 1)))
                        if false {
                            var v156 string = v2
                            _ = v156
                        } else {
                            var v157 uint64 = (v155 / (((v155 % (((^(v155))) | 1))) | 1))
                            v158 := GenS11{F0: 860, F1: (v154 / (((v157 / (((v157 >> ((1) & 63))) | 1))) | 1))}
                            _ = v157
                            _ = v158
                        }
                        _ = v155
                    }
                case "a":
                    v2 = v1
                }
                v159 := GenS22{F0: 1, F1: 2147483647}
                delete(v3, fi151)
                switch 9223372036854775807 {
                case 4294967295:
                    v152 -= v152
                default:
                    switch v2 {
                    case "a":
                        if (-128 <= -128) {
                            v160 := GenS9{F0: 581, F1: 42, F2: 2}
                            _ = v160
                        } else {
                            if (32767 == 42) {
                                switch (fi151 % (((^(fi151))) | 1)) {
                                case 540:
                                    var v161 uint = 201
                                    _ = v161
                                case 1:
                                    var v162 int8 = -128
                                    _ = v162
                                }
                            } else {
                                var v163 int32 = 2147483647
                                v164 := v3[fi151]
                                if true {
                                    var v165 float32 = 1.5
                                    _ = v165
                                }
                                _ = v163
                                _ = v164
                            }
                            v166 := GenS23{F0: false}
                            _ = v166
                        }
                    case "":
                        if false {
                            var v167 string = v1
                            var v168 string = ""
                            v154 -= (v154 % ((v154) | 1))
                            _ = v167
                            _ = v168
                        } else {
                            if true {
                                v169 := []int{fi153, fi151}
                                v1 = "foo bar"
                                v169 = append(v169, (fi151 / ((-960) | 1)))
                                _ = v169
                            } else {
                                if true {
                                    if false {
                                        if (3.14159 != 3.14159) {
                                            var v170 int64 = 0
                                            var v171 int8 = -128
                                            var v172 uint = 255
                                            if true {
                                                var v173 bool = true
                                                _ = v173
                                            } else {
                                                var v174 bool = true
                                                v175 := []int{(^(fi153)), fi153, (fi151 ^ fi151)}
                                                var v176 float64 = 1e10
                                                _ = v174
                                                _ = v175
                                                _ = v176
                                            }
                                            _ = v170
                                            _ = v171
                                            _ = v172
                                        } else {
                                            fmt.Println(fi153, v152, v1)
                                            var v177 uint32 = 1
                                            var v178 int32 = 42
                                            _ = v177
                                            _ = v178
                                        }
                                        v179 := GenS9{F0: 42, F1: 9223372036854775807, F2: 255}
                                        fmt.Println(fi153, v154)
                                        _ = v179
                                    } else {
                                        switch v1 {
                                        case "foo bar":
                                            switch (("hello" + "\n\t") + v2) {
                                            case "hello":
                                                v152 = (v152 + v152)
                                            case "\n\t":
                                                if false {
                                                    if (-733 < 9223372036854775807) {
                                                        fmt.Println(fi153, fi151)
                                                    } else {
                                                        var v180 uint64 = (v154 * 42)
                                                        v181 := GenS6{F0: (v2 + (v2 + (v1 + ("hello" + ((((v2 + v1) + "a") + v2) + v2))))), F1: 1000}
                                                        var v182 uint = 2
                                                        _ = v180
                                                        _ = v181
                                                        _ = v182
                                                    }
                                                    v183 := GenS21{F0: 0, F1: -672}
                                                    _ = v183
                                                } else {
                                                    v184 := []int{fi153, (fi153 | (fi153 << ((4294967295) & 63))), (fi153 >> ((538) & 63))}
                                                    switch -68 {
                                                    case 739:
                                                        v1 = (v1 + v2)
                                                    case -1:
                                                        var v185 string = "a"
                                                        _ = v185
                                                    case 127:
                                                        if true {
                                                            if true {
                                                                v152 -= (v152 + v152)
                                                                fmt.Println(fi153, fi153, v152)
                                                            } else {
                                                                delete(v3, (fi151 / ((fi151) | 1)))
                                                                v1 = v1
                                                            }
                                                            switch (-(fi151)) {
                                                            case -1:
                                                                var v186 uint16 = 32767
                                                                _ = v186
                                                            case 59:
                                                                switch (fi151 + fi151) {
                                                                case 32767:
                                                                    var v187 float64 = 3.14159
                                                                    _ = v187
                                                                case -372:
                                                                    if len(v184) > 0 {
                                                                        v184[0] = (fi151 + (^(fi153)))
                                                                    }
                                                                case 9223372036854775807:
                                                                    v184 = append(v184, fi151)
                                                                    var v188 int32 = -2147483648
                                                                    v189 := GenS24{F0: 0}
                                                                    _ = v188
                                                                    _ = v189
                                                                }
                                                            default:
                                                                var v190 float32 = 0.0
                                                                _ = v190
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(v152, v154)
                                                    var v191 bool = false
                                                    _ = v184
                                                    _ = v191
                                                }
                                            default:
                                                v154 = 4294967295
                                            }
                                        case "\n\t":
                                            var v192 uint = 4294967295
                                            _ = v192
                                        case "hello":
                                            v193 := v3[863]
                                            _ = v193
                                        default:
                                            v2 = v1
                                        }
                                    }
                                    v194 := []int{}
                                    _ = v194
                                }
                                if false {
                                    var v195 float32 = 1e10
                                    var v196 int64 = 216
                                    _ = v195
                                    _ = v196
                                } else {
                                    fmt.Println(fi153, v2)
                                    v2 = v1
                                    fmt.Println(fi153)
                                }
                                var v197 int16 = 481
                                _ = v197
                            }
                        }
                    case "hello":
                        var v198 int32 = 2147483647
                        _ = v198
                    default:
                        if true {
                            var v199 uint32 = 1
                            _ = v199
                        } else {
                            v200 := GenS11{F0: 580, F1: 115}
                            var v201 float64 = ((1.5 / (-2.25 - 1.5)) * ((1e10 * 1.5) - float64(0.1)))
                            _ = v200
                            _ = v201
                        }
                    }
                    v202 := []int{fi151}
                    v152 = (^(v152))
                    _ = v202
                }
                _ = fi153
                _ = v154
                _ = v159
            }
            _ = fi151
            _ = v152
        }
    case "hello":
        for fi203 := 0; fi203 < 8; fi203++ {
            for fi204 := 0; fi204 < 18; fi204++ {
                v205 := []int{}
                switch "" {
                case "\n\t":
                    var v206 float32 = 0.1
                    _ = v206
                case "hello":
                    _ = len(v3)
                case "foo bar":
                    var v207 bool = false
                    _ = v207
                default:
                    v1 = v2
                    v2 = ("\n\t" + ("\n\t" + "\n\t"))
                    var v208 uint8 = 0
                    _ = v208
                }
                var v209 int16 = 32767
                _ = fi204
                _ = v205
                _ = v209
            }
            v2 = (v2 + v2)
            v210 := []int{32768, 896, fi203}
            _ = fi203
            _ = v210
        }
    }
    delete(v3, 32768)
    var v211 int64 = genFn1(255)
    var v212 uint16 = 0
    var v213 uint = 633
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v211:", v211)
    fmt.Println("v212:", v212)
    fmt.Println("v213:", v213)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v211
    _ = v212
    _ = v213
    _ = v3
}
