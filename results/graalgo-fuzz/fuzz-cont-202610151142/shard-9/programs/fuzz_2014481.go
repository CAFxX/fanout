package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int32; F1 int8; F2 uint64 }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 uint; F1 bool }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 int16; F1 int16 }
type GenS8 struct { F0 int64; F1 uint; F2 int8 }
type GenS9 struct { F0 uint16; F1 int; F2 string }
type GenS10 struct { F0 int; F1 uint8 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 bool; F1 uint16; F2 int8 }
type GenS14 struct { F0 uint16; F1 int16 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 uint8; F1 bool; F2 int64 }
type GenS17 struct { F0 uint32; F1 bool }
type GenS18 struct { F0 int32; F1 int8 }
type GenS19 struct { F0 uint16; F1 int64 }
type GenS20 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS21 struct { F0 int8 }
type GenS22 struct { F0 int64 }
type GenS23 struct { F0 uint16; F1 int32 }
type GenS24 struct { F0 int32 }
type GenS25 struct { F0 uint64 }

func main() {
    v1 := make(map[string]bool)
    v2 := []string{("a" + ""), "a"}
    var v3 int32 = 0
    v4 := make(map[uint32]string)
    v5 := make(map[string]int)
    v3 = (-(v3))
    _ = len(v1)
    if (((-2.25 - float64(3.14159)) <= 0.1) || false) {
        for fi6 := 0; fi6 < 17; fi6++ {
            v7 := GenS1{F0: 4294967295}
            var v8 uint64 = 0
            _ = len(v2)
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    fmt.Println(v3)
    if (true || (1.5 != 3.14159)) {
        fmt.Println(v3)
        v3 += v3
        v3 += (v3 & v3)
    } else {
        v9 := GenS1{F0: 42}
        v10 := GenS1{F0: 0}
        if true {
            var v11 int8 = -1
            var v12 float32 = (float32(0.1) + (3.14159 * 3.14159))
            _ = v11
            _ = v12
        }
        _ = v10
        _ = v9
    }
    var v13 int8 = 0
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 string = ""
        if false {
            if false {
                for fi16 := 0; fi16 < 20; fi16++ {
                    if (!(int(v13) > (fi16 / ((fi14) | 1)))) {
                        v15 = "foo bar"
                        if false {
                            v15 = v15
                            if false {
                                var v17 int32 = (v3 / (((v3 * v3)) | 1))
                                fmt.Println(v3, fi16)
                                _ = v17
                            } else {
                                switch 32768 {
                                case -186:
                                    var v18 int32 = v3
                                    _ = v18
                                case 42:
                                    if true {
                                        switch fi16 {
                                        case 128:
                                            v19 := GenS2{F0: 128, F1: 0}
                                            _ = v19
                                        default:
                                            switch v15 {
                                            case "hello":
                                                if false {
                                                    v15 = ("foo bar" + ("\n\t" + v15))
                                                    v15 = (("\n\t" + v15) + "\n\t")
                                                }
                                            default:
                                                var v20 float32 = float32(0.0)
                                                _ = v20
                                            }
                                        }
                                        v21 := GenS3{F0: -2147483648, F1: v13, F2: 192}
                                        _ = v21
                                    } else {
                                        delete(v1, v15)
                                    }
                                case -289:
                                    var v22 uint8 = 126
                                    _ = v22
                                default:
                                    var v23 int8 = (^(v13))
                                    _ = v23
                                }
                            }
                            var v24 int32 = (^(v3))
                            _ = v24
                        } else {
                            var v25 int = int(v3)
                            if false {
                                if true {
                                    var v26 uint = 637
                                    v25 = v25
                                    _ = v26
                                } else {
                                    var v27 float64 = 1.5
                                    v27 = v27
                                    switch 32767 {
                                    case 1:
                                        switch fi14 {
                                        case 32768:
                                            v27 -= v27
                                        }
                                    case 32767:
                                        v15 = "a"
                                    case -132:
                                        fmt.Println(v25)
                                    default:
                                        var v28 float32 = (-2.25 / 1.5)
                                        _ = v28
                                    }
                                    switch "hello" {
                                    case "foo bar":
                                        switch v15 {
                                        case "\n\t":
                                            var v29 uint64 = 4294967295
                                            _ = v29
                                        default:
                                            fmt.Println(fi16)
                                            switch v25 {
                                            case -264:
                                                var v30 bool = (255 == 42)
                                                _ = v30
                                            case -1:
                                                v1[("" + "foo bar")] = true
                                            case 128:
                                                v31 := GenS2{F0: 0, F1: v3}
                                                _ = v31
                                            default:
                                                v2 = append(v2, ("" + "\n\t"))
                                            }
                                            if true {
                                                switch (fi14 | (fi16 << ((4294967295) & 63))) {
                                                case -452:
                                                    _ = len(v2)
                                                }
                                            } else {
                                                v13 -= v13
                                                if (!true) {
                                                    var v32 int32 = 156
                                                    fmt.Println(fi14, fi16, fi16)
                                                    fmt.Println(v15)
                                                    if true {
                                                        fmt.Println(v27, v32)
                                                        if ((!true) || false) {
                                                            v33 := GenS4{F0: -1, F1: false}
                                                            if (-657 >= 1000) {
                                                                fmt.Println(v15, v13, v32)
                                                            }
                                                            var v34 uint = 1
                                                            _ = v33
                                                            _ = v34
                                                        }
                                                    } else {
                                                        if (!((1 <= -1) && true)) {
                                                            if (1 >= 0) {
                                                                if true {
                                                                    v25 = (-(v25))
                                                                } else {
                                                                    v3 = 1
                                                                }
                                                                var v35 int8 = -128
                                                                _ = v35
                                                            }
                                                            v15 = ((((("" + "a") + v15) + ("hello" + ("a" + v15))) + v15) + "foo bar")
                                                            var v36 uint8 = 255
                                                            _ = v36
                                                        } else {
                                                            _ = len(v4)
                                                            if (32767 != -554) {
                                                                switch v15 {
                                                                case "foo bar":
                                                                    fmt.Println(v27, v25, v25)
                                                                }
                                                                if (v3 < (v32 % (((v3 - -233)) | 1))) {
                                                                    v2 = append(v2, (v15 + "\n\t"))
                                                                    if true {
                                                                        v37 := GenS5{F0: 42, F1: false}
                                                                        if false {
                                                                            switch (fi16 << ((255) & 63)) {
                                                                            case -2147483648:
                                                                                if (1e10 <= 3.14159) {
                                                                                    var v38 int8 = (v13 ^ v13)
                                                                                    _ = len(v1)
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                    v38 = v38
                                                                                    _ = v38
                                                                                } else {
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = "foo bar"
                                                                                    }
                                                                                    v1[v15] = ((true && true) || false)
                                                                                    v25 += v25
                                                                                    var v39 float64 = v27
                                                                                    v40 := v4[788]
                                                                                    _ = v39
                                                                                    _ = v40
                                                                                }
                                                                            case 127:
                                                                                if true {
                                                                                    var v41 bool = (1 <= 34)
                                                                                    _ = v41
                                                                                }
                                                                            default:
                                                                                v42 := GenS2{F0: 2, F1: 1}
                                                                                _ = v42
                                                                            }
                                                                            v43 := GenS1{F0: 510}
                                                                            _ = v43
                                                                        } else {
                                                                            var v44 float64 = float64(3.14159)
                                                                            switch (fi16 * 1) {
                                                                            case 9223372036854775807:
                                                                                v27 = float64((-2.25 / float32(1e10)))
                                                                            case 1:
                                                                                var v45 int16 = -1
                                                                                _ = v45
                                                                            case -1:
                                                                                var v46 bool = false
                                                                                _ = v46
                                                                            default:
                                                                                var v47 int32 = 42
                                                                                _ = v47
                                                                            }
                                                                            _ = v44
                                                                        }
                                                                        v15 = (v15 + ("\n\t" + (((v15 + v15) + v15) + v15)))
                                                                        _ = v37
                                                                    }
                                                                    v27 -= -2.25
                                                                    fmt.Println(v15)
                                                                }
                                                                if (0 < -1) {
                                                                    var v48 int = fi16
                                                                    var v49 uint32 = 4294967295
                                                                    v50 := GenS6{F0: -32768, F1: v3}
                                                                    _ = v48
                                                                    _ = v49
                                                                    _ = v50
                                                                }
                                                            } else {
                                                                v3 = v32
                                                            }
                                                            var v51 int32 = v32
                                                            switch (fi14 * fi16) {
                                                            case 32767:
                                                                fmt.Println(v51, v3, v25)
                                                            default:
                                                                v52 := GenS7{F0: -32768, F1: 798}
                                                                _ = v52
                                                            }
                                                            var v53 int32 = (v3 >> ((255) & 63))
                                                            _ = v51
                                                            _ = v53
                                                        }
                                                        v54 := GenS4{F0: -1, F1: false}
                                                        v25 -= int(v13)
                                                        _ = v54
                                                    }
                                                    _ = v32
                                                } else {
                                                    if (1 != 2147483648) {
                                                        v1["hello"] = (v3 <= v3)
                                                        if true {
                                                            v15 = "foo bar"
                                                        } else {
                                                            v55 := GenS8{F0: 1, F1: 49, F2: (v13 >> ((1) & 63))}
                                                            _ = v55
                                                        }
                                                        if true {
                                                            if (128 <= 62) {
                                                                switch (("" + v15) + ((v15 + (v15 + (v15 + (v15 + (v15 + ((((((v15 + (("hello" + v15) + (v15 + ((v15 + ("foo bar" + "foo bar")) + "a")))) + "\n\t") + "\n\t") + (v15 + "hello")) + v15) + v15)))))) + v15)) {
                                                                case "a":
                                                                    v5[v15] = v25
                                                                default:
                                                                    var v56 float64 = v27
                                                                    _ = v56
                                                                }
                                                            } else {
                                                                switch "\n\t" {
                                                                case "":
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        if false {
                                                                            v13 = v13
                                                                            switch v25 {
                                                                            case 9223372036854775807:
                                                                                if true {
                                                                                    var v57 float32 = -2.25
                                                                                    v13 -= v13
                                                                                    var v58 float32 = float32(1e10)
                                                                                    _ = v57
                                                                                    _ = v58
                                                                                } else {
                                                                                    var v59 float32 = 1e10
                                                                                    _ = v59
                                                                                }
                                                                            case 42:
                                                                                if true {
                                                                                    var v60 float32 = (-2.25 * (1e10 * (0.0 + 0.1)))
                                                                                    _ = v60
                                                                                }
                                                                            }
                                                                            v61 := GenS2{F0: 127, F1: (v3 + (v3 ^ int32(v13)))}
                                                                            _ = v61
                                                                        } else {
                                                                            var v62 int = (fi16 & v25)
                                                                            _ = v62
                                                                        }
                                                                    case "a":
                                                                        v25 += (-(fi14))
                                                                    default:
                                                                        var v63 float64 = 1.5
                                                                        _ = v63
                                                                    }
                                                                    v64 := GenS1{F0: 175}
                                                                    _ = v64
                                                                case "\n\t":
                                                                    switch v25 {
                                                                    case -347:
                                                                        _ = len(v5)
                                                                    }
                                                                default:
                                                                    var v65 uint8 = 194
                                                                    _ = v65
                                                                }
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                v15 = ((("a" + (v15 + v15)) + (("" + ("a" + v15)) + "foo bar")) + v15)
                                                                if ((32767 < 32767) || true) {
                                                                    fmt.Println(fi16, v15, v15)
                                                                    fmt.Println(v3)
                                                                }
                                                            }
                                                            var v66 uint32 = 11
                                                            switch "hello" {
                                                            case "\n\t":
                                                                if true {
                                                                    if ((381 < 1) || false) {
                                                                        v27 += v27
                                                                        v67 := GenS9{F0: 889, F1: fi14, F2: "a"}
                                                                        if false {
                                                                            v68 := GenS4{F0: 753, F1: (!(!true))}
                                                                            _ = v68
                                                                        }
                                                                        _ = v67
                                                                    } else {
                                                                        fmt.Println(v27, fi16, v66)
                                                                    }
                                                                    var v69 int16 = -32768
                                                                    _ = v69
                                                                }
                                                                var v70 int = fi16
                                                                _ = v70
                                                            case "a":
                                                                var v71 float64 = (v27 / 1.5)
                                                                _ = v71
                                                            case "hello":
                                                                v72 := GenS10{F0: v25, F1: 159}
                                                                _ = v72
                                                            }
                                                            _ = v66
                                                        }
                                                    } else {
                                                        switch (v15 + "foo bar") {
                                                        case "hello":
                                                            var v73 string = v15
                                                            _ = v73
                                                        case "\n\t":
                                                            v74 := GenS11{F0: v13}
                                                            _ = v74
                                                        case "":
                                                            fmt.Println(fi14, v27, fi16)
                                                        default:
                                                            v13 -= v13
                                                        }
                                                        switch (fi14 + (fi14 / (((fi14 - fi16)) | 1))) {
                                                        case 791:
                                                            var v75 int = 127
                                                            _ = v75
                                                        }
                                                    }
                                                    v13 += -128
                                                }
                                                if true {
                                                    var v76 int32 = v3
                                                    _ = v76
                                                } else {
                                                    if (255 < 32767) {
                                                        var v77 uint16 = 65535
                                                        _ = v77
                                                    } else {
                                                        var v78 string = "foo bar"
                                                        v79 := GenS5{F0: 42, F1: false}
                                                        _ = v78
                                                        _ = v79
                                                    }
                                                    if true {
                                                        v27 += float64(0.0)
                                                        var v80 uint64 = 1
                                                        _ = v80
                                                    }
                                                    switch (fi14 + 103) {
                                                    case -198:
                                                        var v81 int8 = v13
                                                        switch (v25 % ((fi14) | 1)) {
                                                        case 262:
                                                            var v82 float64 = v27
                                                            _ = v82
                                                        case 9223372036854775807:
                                                            if (!false) {
                                                                v83 := GenS12{F0: 2147483647}
                                                                if len(v2) > 0 {
                                                                    v2[0] = "foo bar"
                                                                }
                                                                _ = v83
                                                            } else {
                                                                var v84 uint64 = 4294967295
                                                                v84 += (^(v84))
                                                                v2 = append(v2, "\n\t")
                                                                _ = v84
                                                            }
                                                        case -9223372036854775808:
                                                            v15 = v15
                                                        default:
                                                            var v85 uint64 = 540
                                                            v3 = v3
                                                            _ = v85
                                                        }
                                                        _ = v81
                                                    default:
                                                        v15 = v15
                                                    }
                                                    var v86 int = (-(fi14))
                                                    _ = v86
                                                }
                                            }
                                        }
                                    case "":
                                        v87 := GenS7{F0: -1, F1: 766}
                                        _ = v87
                                    default:
                                        v88 := GenS13{F0: (2147483647 == 1), F1: 366, F2: v13}
                                        _ = v88
                                    }
                                    fmt.Println(fi16, fi16, v13)
                                    v13 = (v13 / (((v13 + v13)) | 1))
                                    _ = v27
                                }
                            }
                            _ = v25
                        }
                        switch fi16 {
                        case 42:
                            if true {
                                v3 += v3
                                var v89 int32 = (v3 << ((2) & 63))
                                v89 = -825
                                _ = v89
                            } else {
                                var v90 int = (fi14 * fi16)
                                v13 = 127
                                v90 -= 922
                                _ = v90
                            }
                            v91 := GenS12{F0: v3}
                            _ = v91
                        case 9223372036854775807:
                            v92 := GenS12{F0: (v3 + 0)}
                            _ = v92
                        case 2147483647:
                            fmt.Println(fi14, v3)
                        default:
                            v13 = v13
                        }
                        v93 := GenS7{F0: 898, F1: -32768}
                        var v94 int16 = 1
                        _ = v93
                        _ = v94
                    } else {
                        if (!(((v3 * v3) <= 2147483647) || false)) {
                            if true {
                                v95 := GenS6{F0: -447, F1: (v3 & (v3 & (v3 % (((v3 & v3)) | 1))))}
                                if (fi16 != fi14) {
                                    switch v15 {
                                    case "a":
                                        v15 = (v15 + "\n\t")
                                    default:
                                        v3 = v3
                                    }
                                    var v96 uint = 2
                                    v97 := GenS2{F0: 128, F1: v3}
                                    _ = v96
                                    _ = v97
                                } else {
                                    switch v15 {
                                    case "foo bar":
                                        fmt.Println(v3, fi14, v15)
                                    default:
                                        fmt.Println(fi16, v3, fi14)
                                        switch (fi14 ^ fi16) {
                                        case -9223372036854775808:
                                            fmt.Println(fi16, v13, v13)
                                        default:
                                            v5[((("\n\t" + v15) + (v15 + "")) + (("\n\t" + v15) + "hello"))] = (fi14 + (fi16 / ((fi14) | 1)))
                                        }
                                        if true {
                                            v2 = append(v2, v15)
                                            if true {
                                                v3 = (-(v3))
                                                v98 := GenS6{F0: 1000, F1: v3}
                                                switch int(v13) {
                                                case -2147483648:
                                                    fmt.Println(v3)
                                                case 4294967295:
                                                    fmt.Println(v13, fi16, v13)
                                                case 595:
                                                    if false {
                                                        v2 = append(v2, v15)
                                                        v99 := GenS6{F0: -1, F1: int32(v13)}
                                                        v100 := v5["\n\t"]
                                                        _ = v100
                                                        _ = v99
                                                    } else {
                                                        var v101 int = 84
                                                        _ = v101
                                                    }
                                                }
                                                _ = v98
                                            }
                                            fmt.Println(v15)
                                        }
                                    }
                                    v3 = v3
                                    switch (fi14 & fi14) {
                                    case -1:
                                        switch (("" + v15) + (v15 + v15)) {
                                        case "a":
                                            if false {
                                                v102 := v4[2147483648]
                                                _ = v102
                                            } else {
                                                _ = len(v5)
                                                var v103 int16 = -14
                                                var v104 bool = (v3 != (v3 * v3))
                                                _ = v103
                                                _ = v104
                                            }
                                        case "hello":
                                            var v105 int32 = v3
                                            _ = v105
                                        case "\n\t":
                                            if true {
                                                v106 := GenS14{F0: 872, F1: -32768}
                                                var v107 int32 = (v3 ^ v3)
                                                _ = v106
                                                _ = v107
                                            } else {
                                                switch "foo bar" {
                                                case "hello":
                                                    v15 = ""
                                                default:
                                                    fmt.Println(v15)
                                                    var v108 uint8 = 69
                                                    _ = v108
                                                }
                                            }
                                        default:
                                            fmt.Println(fi16, fi14)
                                        }
                                    case 4294967295:
                                        var v109 int32 = -2147483648
                                        _ = v109
                                    default:
                                        fmt.Println(fi16, v15)
                                    }
                                }
                                v3 -= 0
                                _ = v95
                            } else {
                                if true {
                                    switch (fi14 - fi14) {
                                    case 4294967295:
                                        if true {
                                            fmt.Println(v3, fi14, v3)
                                        } else {
                                            if (281 > 0) {
                                                fmt.Println(fi14)
                                                var v110 int16 = 453
                                                _ = v110
                                            }
                                            var v111 uint32 = 4294967295
                                            v112 := GenS2{F0: 10, F1: v3}
                                            _ = v111
                                            _ = v112
                                        }
                                    case 609:
                                        if false {
                                            v5[v15] = fi16
                                        } else {
                                            v15 = (v15 + "foo bar")
                                            if false {
                                                var v113 uint16 = 828
                                                _ = v113
                                            } else {
                                                var v114 bool = true
                                                _ = v114
                                            }
                                        }
                                    case 128:
                                        var v115 int64 = 9223372036854775807
                                        _ = v115
                                    default:
                                        switch fi14 {
                                        case 1:
                                            v13 = (^(v13))
                                        case -743:
                                            var v116 int16 = -490
                                            _ = v116
                                        }
                                    }
                                } else {
                                    switch 450 {
                                    case -1:
                                        v117 := GenS13{F0: (false || false), F1: 8, F2: (v13 ^ v13)}
                                        _ = v117
                                    case 42:
                                        v118 := GenS7{F0: int16(v3), F1: 0}
                                        _ = v118
                                    }
                                }
                                if false {
                                    v3 += v3
                                }
                                switch v15 {
                                case "foo bar":
                                    if false {
                                        v13 += (^(v13))
                                        if false {
                                            fmt.Println(v3)
                                            fmt.Println(v15, fi16)
                                            var v119 int64 = 1
                                            v119 -= v119
                                            var v120 float32 = 0.0
                                            _ = v119
                                            _ = v120
                                        }
                                        v121 := GenS15{F0: 831}
                                        _ = v121
                                    } else {
                                        switch fi14 {
                                        case 42:
                                            v15 = (("hello" + v15) + v15)
                                        case 2147483647:
                                            fmt.Println(v15, fi16)
                                        case 127:
                                            var v122 int = fi16
                                            _ = v122
                                        }
                                        v15 = v15
                                        var v123 int16 = -1
                                        _ = v123
                                    }
                                case "\n\t":
                                    var v124 int = (fi14 - fi16)
                                    _ = v124
                                default:
                                    var v125 int32 = (v3 % ((2147483647) | 1))
                                    _ = v125
                                }
                            }
                        } else {
                            switch (fi16 & int(v13)) {
                            case 1:
                                var v126 int64 = -1
                                _ = v126
                            case 533:
                                switch v15 {
                                case "":
                                    if false {
                                        fmt.Println(v13)
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v127 uint = 18446744073709551615
                                        _ = v127
                                    } else {
                                        switch 186 {
                                        case 128:
                                            if ((((3.14159 + 0.1) / 0.1) - 0.0) >= 1e10) {
                                                switch v15 {
                                                case "hello":
                                                    v128 := GenS8{F0: -1, F1: 22, F2: v13}
                                                    _ = v128
                                                }
                                                v13 -= v13
                                                var v129 bool = (true || (!true))
                                                if (int32(v13) == (v3 + (v3 % (((v3 / ((-1) | 1))) | 1)))) {
                                                    v13 -= v13
                                                }
                                                if false {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    switch v15 {
                                                    case "\n\t":
                                                        switch fi16 {
                                                        case 42:
                                                            if (!(!v129)) {
                                                                if v129 {
                                                                    var v130 float32 = 0.0
                                                                    v131 := GenS4{F0: -32768, F1: v129}
                                                                    _ = v130
                                                                    _ = v131
                                                                } else {
                                                                    v129 = ((!(2147483648 < 3)) || true)
                                                                    var v132 int8 = v13
                                                                    _ = v132
                                                                }
                                                                var v133 bool = true
                                                                _ = v133
                                                            }
                                                        case 2147483647:
                                                            if (1 != 128) {
                                                                v134 := GenS12{F0: v3}
                                                                if (-267 >= 32767) {
                                                                    var v135 int32 = v3
                                                                    v13 = -86
                                                                    switch -1 {
                                                                    case 32:
                                                                        var v136 string = v15
                                                                        _ = v136
                                                                    default:
                                                                        var v137 bool = v129
                                                                        _ = v137
                                                                    }
                                                                    v135 = (^(v3))
                                                                    var v138 bool = ((false && v129) && v129)
                                                                    _ = v135
                                                                    _ = v138
                                                                }
                                                                if v129 {
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        v139 := GenS16{F0: 134, F1: (v129 || v129), F2: 9223372036854775807}
                                                                        _ = v139
                                                                    case "foo bar":
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    }
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    if (!(!(127 < 0))) {
                                                                        v3 += (^(v3))
                                                                        if false {
                                                                            v140 := GenS11{F0: (v13 & v13)}
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            var v141 int8 = v13
                                                                            fmt.Println(v15)
                                                                            v15 = v15
                                                                            fmt.Println(v129)
                                                                            _ = v140
                                                                            _ = v141
                                                                        }
                                                                    } else {
                                                                        if v129 {
                                                                            v142 := GenS17{F0: 962, F1: true}
                                                                            _ = v142
                                                                        } else {
                                                                            fmt.Println(v3)
                                                                        }
                                                                        fmt.Println(v13, fi14, v129)
                                                                        v13 = v13
                                                                        v5[v15] = (fi14 - -2147483648)
                                                                    }
                                                                }
                                                                _ = v134
                                                            } else {
                                                                if (128 <= 140) {
                                                                    if false {
                                                                        v143 := GenS18{F0: (v3 ^ (v3 / (((v3 >> ((813) & 63))) | 1))), F1: (v13 + v13)}
                                                                        v144 := GenS1{F0: 65535}
                                                                        _ = v143
                                                                        _ = v144
                                                                    }
                                                                    switch fi16 {
                                                                    case -261:
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            v145 := GenS19{F0: 2, F1: 42}
                                                                            _ = v145
                                                                        default:
                                                                            if (false && (18446744073709551615 > 65535)) {
                                                                                v13 -= (v13 & 127)
                                                                                v129 = true
                                                                                if (9223372036854775807 > -128) {
                                                                                    var v146 int8 = (v13 | v13)
                                                                                    if true {
                                                                                        v147 := GenS11{F0: v146}
                                                                                        var v148 uint8 = 128
                                                                                        v13 -= v13
                                                                                        _ = v147
                                                                                        _ = v148
                                                                                    } else {
                                                                                        fmt.Println(v15)
                                                                                        v129 = (9223372036854775807 > 42)
                                                                                    }
                                                                                    _ = v146
                                                                                }
                                                                                _ = len(v2)
                                                                                v3 -= (v3 & v3)
                                                                            } else {
                                                                                var v149 int = fi14
                                                                                _ = v149
                                                                            }
                                                                            fmt.Println(v129, v15)
                                                                        }
                                                                    case 42:
                                                                        v150 := GenS8{F0: -723, F1: 65535, F2: v13}
                                                                        var v151 int8 = (v13 ^ v13)
                                                                        if ((fi14 + fi16) != (fi16 | (fi16 ^ fi16))) {
                                                                            var v152 uint8 = 2
                                                                            v151 = (v151 | v151)
                                                                            v153 := GenS20{F0: 42, F1: -1, F2: 1}
                                                                            fmt.Println(v13, v13)
                                                                            v13 = -128
                                                                            _ = v152
                                                                            _ = v153
                                                                        }
                                                                        _ = v150
                                                                        _ = v151
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    case "foo bar":
                                                        v13 -= (v13 & (v13 << ((255) & 63)))
                                                    }
                                                    v4[2147483648] = "hello"
                                                } else {
                                                    v15 = "foo bar"
                                                }
                                                _ = v129
                                            } else {
                                                var v154 uint32 = 2147483647
                                                if (42 != -32768) {
                                                    v2 = append(v2, v15)
                                                    v154 += (v154 & v154)
                                                    if false {
                                                        var v155 float32 = 0.1
                                                        var v156 int32 = 0
                                                        _ = v155
                                                        _ = v156
                                                    }
                                                }
                                                v154 += v154
                                                _ = v154
                                            }
                                        }
                                        v15 = "\n\t"
                                    }
                                case "a":
                                    var v157 int32 = -2147483648
                                    _ = v157
                                }
                            }
                        }
                        delete(v5, (v15 + ((((v15 + (((("hello" + v15) + v15) + v15) + "\n\t")) + v15) + "\n\t") + v15)))
                        if (!false) {
                            v15 = ("foo bar" + v15)
                            if true {
                                if (false && false) {
                                    v3 = 42
                                    delete(v5, (v15 + v15))
                                    v13 += (-(v13))
                                    v158 := GenS19{F0: 65535, F1: -1}
                                    v15 = (v15 + v15)
                                    _ = v158
                                }
                                if (!((true || true) || true)) {
                                    switch (-(fi16)) {
                                    case -695:
                                        var v159 string = v15
                                        _ = v159
                                    case 0:
                                        fmt.Println(fi16, v3)
                                    default:
                                        var v160 float32 = ((1e10 * 1.5) / ((3.14159 - -2.25) / (((1e10 * float32(float64(1e10))) + 1.5) - ((1e10 / 1.0) + ((0.0 + 1.5) * (0.1 / ((1.5 - (-2.25 - 1e10)) / (((3.14159 * 0.0) + -2.25) / ((0.0 + 1e10) / 1e10)))))))))
                                        _ = v160
                                    }
                                    v1["\n\t"] = false
                                    if ((735 == 32767) && (((0.1 * (0.0 - 1.5)) / 1e10) <= -2.25)) {
                                        v161 := GenS3{F0: v3, F1: v13, F2: 42}
                                        _ = v161
                                    } else {
                                        var v162 int16 = 1000
                                        var v163 uint8 = 128
                                        var v164 int = fi16
                                        _ = v162
                                        _ = v163
                                        _ = v164
                                    }
                                }
                                v165 := GenS10{F0: 1, F1: 1}
                                v4[1] = "hello"
                                _ = v165
                            }
                            v15 = v15
                        } else {
                            var v166 bool = (((v15 + "\n\t") + v15) >= ((v15 + v15) + "foo bar"))
                            if v166 {
                                if true {
                                    var v167 int32 = -980
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if ((fi14 * 0) != (-(fi14))) {
                                        v168 := GenS21{F0: v13}
                                        _ = v168
                                    }
                                    _ = v167
                                }
                                if (2 > 42) {
                                    if true {
                                        v15 = ""
                                        v5[v15] = (-(fi14))
                                    }
                                    var v169 int16 = 42
                                    _ = v169
                                }
                            } else {
                                var v170 float64 = (1e10 - 1.5)
                                _ = v170
                            }
                            if (!(214 >= 1)) {
                                if len(v2) > 0 {
                                    v2[0] = (v15 + "foo bar")
                                }
                            } else {
                                var v171 bool = (0 != 0)
                                fmt.Println(v166)
                                _ = v171
                            }
                            if v166 {
                                v15 = v15
                                v3 += (v3 | v3)
                            }
                            _ = v166
                        }
                    }
                    v13 += v13
                    _ = fi16
                }
                if false {
                    for fi172 := 0; fi172 < 20; fi172++ {
                        v173 := GenS22{F0: -1}
                        _ = fi172
                        _ = v173
                    }
                    v5[v15] = fi14
                } else {
                    switch (v15 + v15) {
                    case "foo bar":
                        v13 += (v13 << ((2) & 63))
                    case "hello":
                        v174 := GenS7{F0: 0, F1: 1000}
                        _ = v174
                    case "a":
                        if len(v2) > 0 {
                            v2[0] = v15
                        }
                    default:
                        if (!true) {
                            switch (v15 + "hello") {
                            case "\n\t":
                                v13 = 1
                            case "foo bar":
                                for fi175 := 0; fi175 < 5; fi175++ {
                                    if (1 == 893) {
                                        switch v15 {
                                        case "\n\t":
                                            fmt.Println(v3, fi14, fi175)
                                        case "a":
                                            if (0.0 == 0.1) {
                                                var v176 uint64 = 624
                                                var v177 int = 9223372036854775807
                                                v15 = (v15 + ((("hello" + (v15 + ((v15 + v15) + v15))) + "\n\t") + "a"))
                                                _ = v176
                                                _ = v177
                                            }
                                        }
                                    } else {
                                        switch (v15 + v15) {
                                        case "foo bar":
                                            var v178 int = (^(fi175))
                                            _ = v178
                                        default:
                                            var v179 uint64 = 4294967295
                                            _ = v179
                                        }
                                        if ((v3 * (v3 ^ (v3 | (v3 + 2147483647)))) != (v3 << ((255) & 63))) {
                                            fmt.Println(fi14, v13, v13)
                                            var v180 uint32 = 0
                                            _ = v180
                                        }
                                    }
                                    _ = fi175
                                }
                            }
                            var v181 uint32 = 1
                            switch v15 {
                            case "a":
                                _ = len(v2)
                            case "foo bar":
                                for fi182 := 0; fi182 < 1; fi182++ {
                                    var v183 bool = false
                                    v183 = true
                                    v184 := GenS15{F0: 613}
                                    v1[(v15 + "a")] = (!v183)
                                    v185 := GenS14{F0: 0, F1: 0}
                                    _ = fi182
                                    _ = v183
                                    _ = v184
                                    _ = v185
                                }
                            }
                            _ = v181
                        } else {
                            delete(v4, 2)
                        }
                    }
                }
                v15 = v15
            }
            v15 = v15
            v186 := GenS23{F0: 511, F1: v3}
            _ = v186
        } else {
            var v187 int8 = v13
            if false {
                switch (fi14 / (((fi14 ^ 127)) | 1)) {
                case 32768:
                    switch fi14 {
                    case 9223372036854775807:
                        v188 := GenS24{F0: (-(v3))}
                        _ = v188
                    case -1:
                        v189 := GenS25{F0: 42}
                        _ = v189
                    }
                }
                var v190 int32 = (v3 % ((v3) | 1))
                if (2 < 555) {
                    fmt.Println(v3, v190, v13)
                    v191 := GenS4{F0: 1, F1: true}
                    v13 = v13
                    _ = v191
                }
                _ = v190
            } else {
                fmt.Println(v187, v3, fi14)
                fmt.Println(v3, v3)
                var v192 float64 = float64(0.0)
                _ = v192
            }
            _ = v187
        }
        _ = fi14
        _ = v15
    }
    var v193 int32 = (v3 * (v3 << ((689) & 63)))
    switch -9223372036854775808 {
    case 1:
        fmt.Println(v3, v193, v193)
    case 2147483647:
        v194 := GenS25{F0: 4294967295}
        _ = v194
    case 9223372036854775807:
        v195 := v4[0]
        _ = v195
    }
    v4[2147483647] = ("\n\t" + "")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v193:", v193)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v193
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
