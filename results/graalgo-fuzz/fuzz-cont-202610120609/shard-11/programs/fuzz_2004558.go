package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int64; F1 uint8; F2 uint16 }
type GenS8 struct { F0 bool; F1 int32; F2 int32 }
type GenS9 struct { F0 uint64; F1 int64; F2 int8 }
type GenS10 struct { F0 uint64; F1 uint64 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint; F1 int16 }
type GenS14 struct { F0 int64; F1 int32 }
type GenS15 struct { F0 uint32; F1 string }
type GenS16 struct { F0 int16; F1 int; F2 int }
type GenS17 struct { F0 int8; F1 uint }
type GenS18 struct { F0 uint32 }
func genFn0() bool {
    var v5 int8 = 101
    v6 := GenS1{F0: true, F1: 0, F2: 4294967295}
    var v7 int16 = -776
    v8 := []int{-2147483648, 784, -9223372036854775808}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (!false)
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]int)
    v3 := []int{-2147483648, -2147483648}
    var v4 string = "a"
    var v9 int16 = 1
    v3 = append(v3, 127)
    var v10 int32 = 2147483647
    var v11 uint = 65535
    for fi12 := 0; fi12 < 12; fi12++ {
        v9 += v9
        fmt.Println(v4, v4)
        v13 := GenS1{F0: (186 < 243), F1: 9223372036854775807, F2: 4294967295}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        switch 42 {
        case 2147483647:
            switch v4 {
            case "":
                switch (v4 + v4) {
                case "a":
                    for fi15 := 0; fi15 < 14; fi15++ {
                        v9 -= 1
                        var v16 bool = false
                        var v17 int32 = v10
                        if false {
                            _ = len(v3)
                            var v18 float64 = float64(1e10)
                            v19 := GenS1{F0: true, F1: 0, F2: 0}
                            _ = v18
                            _ = v19
                        } else {
                            var v20 uint = (v11 ^ 18446744073709551615)
                            v21 := GenS1{F0: (v16 && v16), F1: -9223372036854775808, F2: 1}
                            v22 := GenS1{F0: false, F1: -732, F2: 449}
                            var v23 uint32 = 2
                            switch v4 {
                            case "foo bar":
                                fmt.Println(v4, v23)
                            case "":
                                v1 = 0.1
                            case "hello":
                                if v16 {
                                    var v24 int16 = -213
                                    _ = v24
                                }
                            default:
                                var v25 int8 = 117
                                _ = v25
                            }
                            var v26 int32 = v17
                            fmt.Println(v16, v4)
                            v27 := v2[v4]
                            _ = v20
                            _ = v21
                            _ = v22
                            _ = v23
                            _ = v26
                            _ = v27
                        }
                        var v28 int32 = v17
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v28
                    }
                default:
                    v4 = (v4 + (((v4 + (("a" + v4) + "hello")) + v4) + (("foo bar" + "") + "")))
                }
            case "a":
                for fi29 := 0; fi29 < 20; fi29++ {
                    if len(v3) > 0 {
                        v3[0] = fi14
                    }
                    var v30 uint16 = 32767
                    var v31 float64 = (1e10 * 1e10)
                    if ((float64(1e10) >= float64(0.1)) && false) {
                        var v32 int8 = 100
                        v33 := GenS1{F0: true, F1: int64(v9), F2: 4294967295}
                        _ = v32
                        _ = v33
                    } else {
                        var v34 int8 = 1
                        switch 127 {
                        case 42:
                            switch (fi14 + (fi14 / ((fi14) | 1))) {
                            case 184:
                                if false {
                                    v3 = append(v3, -367)
                                    if (false || (1e10 != 1.5)) {
                                        switch int(v11) {
                                        case -2147483648:
                                            switch v4 {
                                            case "foo bar":
                                                v9 += (v9 | (v9 / (((v9 + v9)) | 1)))
                                            case "\n\t":
                                                var v35 int = fi14
                                                _ = v35
                                            case "":
                                                var v36 uint = 2
                                                _ = v36
                                            }
                                        default:
                                            fmt.Println(fi14)
                                        }
                                        fmt.Println(v9)
                                    }
                                }
                            case -9223372036854775808:
                                v37 := GenS1{F0: (!true), F1: int64(v10), F2: 4294967295}
                                _ = v37
                            default:
                                v38 := GenS2{F0: (v30 <= 928), F1: 2, F2: 1}
                                _ = v38
                            }
                        default:
                            v3 = append(v3, 2147483647)
                        }
                        v9 += (v9 | -287)
                        _ = v34
                    }
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            case "\n\t":
                for fi39 := 0; fi39 < 16; fi39++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v40 := GenS3{F0: 42, F1: 2}
                    if ((!(!(!(int8(v10) <= 1)))) && true) {
                        var v41 float32 = (((1.5 * 1.5) - 1.5) * 0.1)
                        switch v4 {
                        case "foo bar":
                            switch v4 {
                            case "foo bar":
                                fmt.Println(v41, v11)
                            case "a":
                                if (!false) {
                                    var v42 uint32 = 1
                                    _ = v42
                                } else {
                                    switch fi14 {
                                    case -2147483648:
                                        if (float64(v41) == (((1.5 - (v1 + v1)) - (v1 / v1)) * 0.1)) {
                                            v43 := v2[v4]
                                            switch 9223372036854775807 {
                                            case -1:
                                                if false {
                                                    var v44 uint8 = 2
                                                    if false {
                                                        v11 = v11
                                                    } else {
                                                        if ((4294967295 < 4294967295) || ((!true) && true)) {
                                                            if true {
                                                                v9 = v9
                                                            }
                                                            v44 = v44
                                                            var v45 uint32 = 42
                                                            _ = v45
                                                        } else {
                                                            if (true || true) {
                                                                v1 = v1
                                                            }
                                                            v43 = 469
                                                            if (int(v44) > (fi14 >> ((18446744073709551615) & 63))) {
                                                                if ((fi39 >> ((v11) & 63)) != (v43 % ((32768) | 1))) {
                                                                    v1 = float64(0.1)
                                                                    delete(v2, v4)
                                                                }
                                                                switch (v43 + v43) {
                                                                case 0:
                                                                    if true {
                                                                        v46 := GenS4{F0: 811, F1: (v41 > 3.14159), F2: 4294967295}
                                                                        if (false && (!true)) {
                                                                            v47 := GenS1{F0: false, F1: -660, F2: 4294967295}
                                                                            _ = v47
                                                                        } else {
                                                                            if true {
                                                                                if (fi14 > v43) {
                                                                                    var v48 float64 = v1
                                                                                    var v49 uint64 = 42
                                                                                    v2["\n\t"] = fi14
                                                                                    v48 += v48
                                                                                    var v50 uint64 = (^(v49))
                                                                                    _ = v48
                                                                                    _ = v49
                                                                                    _ = v50
                                                                                }
                                                                                var v51 int = fi39
                                                                                fmt.Println(v4, v41, v11)
                                                                                _ = v51
                                                                            } else {
                                                                                switch (v4 + (v4 + v4)) {
                                                                                case "":
                                                                                    v3 = append(v3, (v43 - -9223372036854775808))
                                                                                default:
                                                                                    v52 := GenS5{F0: 498, F1: 18446744073709551615}
                                                                                    _ = v52
                                                                                }
                                                                                var v53 uint = v11
                                                                                var v54 uint8 = (v44 - v44)
                                                                                _ = v53
                                                                                _ = v54
                                                                            }
                                                                            v41 += v41
                                                                            v55 := GenS5{F0: 693, F1: 0}
                                                                            _ = v55
                                                                        }
                                                                        v41 = v41
                                                                        _ = v46
                                                                    }
                                                                case -169:
                                                                    v56 := GenS2{F0: (42 < 65535), F1: 16, F2: 0}
                                                                    _ = v56
                                                                case -610:
                                                                    v11 -= 0
                                                                    v57 := GenS6{F0: v11}
                                                                    _ = v57
                                                                }
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    if (0 == 2) {
                                                                        v9 += v9
                                                                    }
                                                                case "a":
                                                                    switch v4 {
                                                                    case "a":
                                                                        fmt.Println(v44)
                                                                    }
                                                                case "foo bar":
                                                                    fmt.Println(v41)
                                                                default:
                                                                    switch v4 {
                                                                    case "a":
                                                                        v9 += v9
                                                                    case "\n\t":
                                                                        v4 = (((v4 + "foo bar") + "\n\t") + (v4 + (((v4 + "a") + ("hello" + v4)) + v4)))
                                                                    case "hello":
                                                                        fmt.Println(fi14)
                                                                    }
                                                                    switch v4 {
                                                                    case "hello":
                                                                        fmt.Println(v10)
                                                                    case "\n\t":
                                                                        fmt.Println(v4, v44)
                                                                    case "foo bar":
                                                                        if false {
                                                                            var v58 bool = true
                                                                            switch fi39 {
                                                                            case -576:
                                                                                v58 = v58
                                                                            case 4294967295:
                                                                                v59 := GenS7{F0: 42, F1: 2, F2: 65535}
                                                                                _ = v59
                                                                            default:
                                                                                v41 = (v41 - 3.14159)
                                                                            }
                                                                            _ = len(v2)
                                                                            fmt.Println(v9)
                                                                            _ = v58
                                                                        } else {
                                                                            if false {
                                                                                v60 := GenS8{F0: true, F1: v10, F2: 710}
                                                                                var v61 int64 = -1
                                                                                fmt.Println(v61, v1, v10)
                                                                                _ = v60
                                                                                _ = v61
                                                                            } else {
                                                                                var v62 uint16 = 478
                                                                                var v63 int16 = v9
                                                                                v62 += (^(v62))
                                                                                _ = v62
                                                                                _ = v63
                                                                            }
                                                                            var v64 uint = 192
                                                                            var v65 int32 = v10
                                                                            switch (v4 + "hello") {
                                                                            case "":
                                                                                if (2 > 1) {
                                                                                    if true {
                                                                                        var v66 float64 = (v1 - -2.25)
                                                                                        switch v4 {
                                                                                        case "\n\t":
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                        case "foo bar":
                                                                                            v41 += float32(float64(v41))
                                                                                        }
                                                                                        var v67 int64 = -319
                                                                                        _ = v66
                                                                                        _ = v67
                                                                                    }
                                                                                    switch 1 {
                                                                                    case -1:
                                                                                        if false {
                                                                                            var v68 uint8 = (v44 / ((v44) | 1))
                                                                                            if ((1 != -128) && false) {
                                                                                                delete(v2, "hello")
                                                                                                v69 := GenS9{F0: 498, F1: 0, F2: -62}
                                                                                                var v70 uint16 = 1
                                                                                                _ = v69
                                                                                                _ = v70
                                                                                            }
                                                                                            switch fi14 {
                                                                                            case 2147483647:
                                                                                                var v71 int64 = 0
                                                                                                _ = v71
                                                                                            case 519:
                                                                                                var v72 int64 = 1
                                                                                                _ = v72
                                                                                            case 127:
                                                                                                var v73 float32 = 1.5
                                                                                                _ = v73
                                                                                            default:
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = (^(fi39))
                                                                                                }
                                                                                            }
                                                                                            _ = v68
                                                                                        }
                                                                                    default:
                                                                                        v65 -= 2147483647
                                                                                    }
                                                                                    v41 -= float32(v1)
                                                                                }
                                                                            default:
                                                                                v74 := GenS10{F0: 0, F1: 2}
                                                                                _ = v74
                                                                            }
                                                                            _ = v64
                                                                            _ = v65
                                                                        }
                                                                    default:
                                                                        v2[("a" + "")] = (v43 + (fi39 + fi39))
                                                                    }
                                                                    v44 = v44
                                                                }
                                                                var v75 int = (-(v43))
                                                                var v76 int32 = v10
                                                                fmt.Println(v1)
                                                                v43 = (v43 << (((v11 & v11)) & 63))
                                                                v75 = 127
                                                                _ = v75
                                                                _ = v76
                                                            }
                                                            v77 := GenS11{F0: (v11 % (((v11 ^ (^(v11)))) | 1))}
                                                            _ = v77
                                                        }
                                                    }
                                                    _ = v44
                                                }
                                            case -384:
                                                switch "hello" {
                                                case "hello":
                                                    if true {
                                                        v1 = (v1 - (float64((v41 - (3.14159 / v41))) + v1))
                                                        v1 += float64(1e10)
                                                        if (false && (true && (((0 != 18446744073709551615) && false) && true))) {
                                                            fmt.Println(v41, v4, v43)
                                                            var v78 uint8 = 113
                                                            v43 = fi14
                                                            _ = v78
                                                        }
                                                    }
                                                case "a":
                                                    v41 = (1.5 + ((((v41 * (float32((-2.25 + v1)) - (float32(v1) * 0.0))) + 1e10) * (3.14159 + v41)) / (0.1 * (-2.25 + (1e10 * v41)))))
                                                }
                                            }
                                            v79 := GenS12{F0: 42}
                                            v80 := GenS13{F0: v11, F1: v9}
                                            _ = len(v2)
                                            _ = v43
                                            _ = v79
                                            _ = v80
                                        } else {
                                            if (false || ((-32768 > 32767) || (true || false))) {
                                                var v81 int = fi39
                                                switch v4 {
                                                case "":
                                                    var v82 uint = v11
                                                    _ = v82
                                                case "hello":
                                                    if (fi39 > (v81 | 1)) {
                                                        _ = len(v2)
                                                        switch fi39 {
                                                        case 42:
                                                            v10 = (-(v10))
                                                        case -9223372036854775808:
                                                            var v83 uint16 = 255
                                                            _ = v83
                                                        default:
                                                            if true {
                                                                v81 -= 127
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            } else {
                                                                var v84 bool = (1 > 0)
                                                                _ = v84
                                                            }
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            v3 = append(v3, fi14)
                                                            v4 = "foo bar"
                                                        }
                                                        v41 -= v41
                                                    } else {
                                                        v81 += fi39
                                                    }
                                                case "\n\t":
                                                    fmt.Println(v10, v11)
                                                default:
                                                    v85 := GenS14{F0: 601, F1: v10}
                                                    _ = v85
                                                }
                                                var v86 uint32 = 2147483647
                                                _ = v81
                                                _ = v86
                                            }
                                        }
                                        var v87 uint16 = 1
                                        var v88 bool = true
                                        fmt.Println(v10, fi14)
                                        var v89 int64 = -1
                                        if (((v89 ^ (^(v89))) <= 9223372036854775807) || (v41 < v41)) {
                                            _ = len(v2)
                                            var v90 int32 = 483
                                            _ = v90
                                        }
                                        _ = v87
                                        _ = v88
                                        _ = v89
                                    }
                                }
                            default:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        case "":
                            if (v11 <= v11) {
                                if len(v3) > 0 {
                                    v3[0] = 32767
                                }
                                if ((fi39 | fi14) > fi14) {
                                    var v91 uint32 = 2147483648
                                    v4 = v4
                                    switch ("hello" + ((v4 + (((v4 + v4) + v4) + (v4 + "foo bar"))) + "foo bar")) {
                                    case "":
                                        if len(v3) > 0 {
                                            v3[0] = 9223372036854775807
                                        }
                                    }
                                    _ = v91
                                }
                            }
                        case "a":
                            v92 := GenS10{F0: 1, F1: 2}
                            _ = v92
                        }
                        switch ((v4 + (v4 + (v4 + "\n\t"))) + v4) {
                        case "\n\t":
                            v1 = (v1 * v1)
                        case "hello":
                            switch fi14 {
                            case 9223372036854775807:
                                if (v4 > v4) {
                                    if (v1 < float64(v41)) {
                                        fmt.Println(v4, fi39, fi39)
                                        if (true && (!false)) {
                                            v11 += v11
                                            fmt.Println(v1, fi14, v11)
                                        } else {
                                            v1 += v1
                                            if false {
                                                var v93 uint = v11
                                                v94 := GenS9{F0: 240, F1: 1, F2: -12}
                                                _ = v93
                                                _ = v94
                                            } else {
                                                var v95 float64 = 1.5
                                                var v96 int8 = -1
                                                fmt.Println(v1)
                                                if false {
                                                    v96 = 100
                                                    if true {
                                                        v97 := v2[v4]
                                                        _ = v97
                                                    }
                                                }
                                                _ = v95
                                                _ = v96
                                            }
                                            switch "" {
                                            case "\n\t":
                                                v4 = v4
                                            default:
                                                if (951 >= 4294967295) {
                                                    if (v4 != v4) {
                                                        switch "a" {
                                                        case "":
                                                            v98 := GenS11{F0: v11}
                                                            _ = v98
                                                        case "a":
                                                            v99 := GenS10{F0: 191, F1: 745}
                                                            _ = v99
                                                        case "hello":
                                                            var v100 uint16 = 243
                                                            _ = v100
                                                        }
                                                        v101 := GenS12{F0: -23}
                                                        _ = v101
                                                    } else {
                                                        if (0 < 100) {
                                                            var v102 string = v4
                                                            _ = v102
                                                        }
                                                        fmt.Println(v4, fi14, v9)
                                                        var v103 uint16 = 104
                                                        v104 := GenS6{F0: v11}
                                                        _ = v103
                                                        _ = v104
                                                    }
                                                } else {
                                                    if false {
                                                        fmt.Println(v1, fi39)
                                                    } else {
                                                        v41 -= v41
                                                        if true {
                                                            var v105 int8 = 80
                                                            var v106 float64 = -2.25
                                                            _ = v105
                                                            _ = v106
                                                        }
                                                    }
                                                }
                                            }
                                            v1 -= ((v1 + (0.1 / v1)) - 1e10)
                                            v107 := GenS15{F0: 776, F1: ("hello" + "foo bar")}
                                            _ = v107
                                        }
                                        v108 := GenS6{F0: (^(v11))}
                                        _ = v108
                                    }
                                    fmt.Println(v4, v11)
                                    _ = len(v2)
                                    if false {
                                        if true {
                                            v109 := GenS16{F0: v9, F1: (fi14 ^ (fi14 | -370)), F2: fi14}
                                            delete(v2, "")
                                            v1 -= -2.25
                                            _ = v109
                                        } else {
                                            _ = len(v3)
                                        }
                                        var v110 uint = (v11 + v11)
                                        var v111 int = (fi39 - fi14)
                                        _ = v110
                                        _ = v111
                                    } else {
                                        var v112 string = v4
                                        _ = v112
                                    }
                                    fmt.Println(fi14, v11)
                                }
                            case 0:
                                var v113 uint = v11
                                _ = v113
                            }
                        case "":
                            var v114 uint64 = 42
                            _ = v114
                        }
                        _ = v41
                    } else {
                        v115 := GenS8{F0: false, F1: (^(v10)), F2: (v10 & (v10 * 0))}
                        fmt.Println(v11, v11)
                        delete(v2, "")
                        v4 = v4
                        var v116 float64 = (v1 - 1.5)
                        if false {
                            v1 = float64(1.5)
                        }
                        var v117 uint8 = 1
                        _ = v115
                        _ = v116
                        _ = v117
                    }
                    _ = fi39
                    _ = v40
                }
            default:
                for fi118 := 0; fi118 < 13; fi118++ {
                    switch v4 {
                    case "a":
                        v119 := v2[v4]
                        _ = v119
                    case "":
                        var v120 int16 = -32768
                        _ = v120
                    }
                    v9 = v9
                    v11 -= 724
                    _ = fi118
                }
            }
        }
        for fi121 := 0; fi121 < 4; fi121++ {
            if true {
                var v122 uint32 = 689
                if (0.0 <= 0.1) {
                    v123 := GenS5{F0: 864, F1: 1}
                    _ = v123
                } else {
                    if false {
                        v4 = v4
                        if false {
                            var v124 float64 = 0.0
                            _ = v124
                        }
                        if false {
                            v11 = 300
                            switch 260 {
                            case -801:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case 573:
                                var v125 int8 = 1
                                _ = v125
                            default:
                                v122 -= (v122 & (v122 * (v122 | v122)))
                            }
                            v122 = v122
                            if false {
                                _ = len(v2)
                            }
                        } else {
                            fmt.Println(v10, fi14, v9)
                            var v126 uint64 = 0
                            v127 := GenS11{F0: v11}
                            var v128 bool = false
                            _ = v126
                            _ = v127
                            _ = v128
                        }
                        v129 := GenS2{F0: false, F1: 603, F2: 442}
                        v4 = ((v4 + (v4 + "\n\t")) + ((v4 + "foo bar") + v4))
                        v130 := v2[v4]
                        v131 := GenS6{F0: 431}
                        v130 = fi121
                        _ = v129
                        _ = v130
                        _ = v131
                    } else {
                        switch v4 {
                        case "foo bar":
                            var v132 int8 = -128
                            _ = v132
                        default:
                            v3 = append(v3, (fi14 / ((1) | 1)))
                        }
                        v122 = v122
                    }
                    switch fi121 {
                    case 32768:
                        var v133 int64 = 0
                        _ = v133
                    default:
                        if true {
                            switch fi121 {
                            case 32768:
                                v134 := GenS17{F0: 0, F1: (v11 / (((^(v11))) | 1))}
                                _ = v134
                            case 704:
                                var v135 uint = (v11 + v11)
                                _ = v135
                            case 0:
                                v122 = 2147483647
                            }
                            if true {
                                var v136 float32 = float32(0.0)
                                if false {
                                    v137 := GenS16{F0: (v9 / ((v9) | 1)), F1: fi121, F2: (fi121 & (fi14 >> ((v11) & 63)))}
                                    v4 = (("foo bar" + v4) + (("foo bar" + ((((("\n\t" + "a") + v4) + v4) + ((v4 + (("\n\t" + ("a" + v4)) + v4)) + v4)) + v4)) + "a"))
                                    v136 -= 3.14159
                                    _ = v137
                                } else {
                                    v136 = v136
                                    if (130 > 2) {
                                        v3 = append(v3, (fi121 * fi14))
                                        var v138 bool = (2 >= 2)
                                        v139 := GenS16{F0: -1, F1: 4294967295, F2: 127}
                                        if (!(!v138)) {
                                            var v140 int8 = -1
                                            _ = v140
                                        }
                                        _ = v138
                                        _ = v139
                                    } else {
                                        switch 2147483647 {
                                        case 1:
                                            v141 := GenS1{F0: false, F1: -1, F2: 38}
                                            _ = v141
                                        default:
                                            fmt.Println(v122)
                                        }
                                    }
                                }
                                v10 -= 1
                                _ = v136
                            }
                            v142 := GenS18{F0: v122}
                            _ = v142
                        } else {
                            if (true || (!true)) {
                                if false {
                                    if false {
                                        v122 = (v122 | 4294967295)
                                    }
                                } else {
                                    v4 = v4
                                    switch v4 {
                                    case "foo bar":
                                        v122 -= v122
                                    }
                                }
                                if ((v11 ^ v11) < (^(v11))) {
                                    v122 = (v122 % (((v122 / (((v122 | (v122 << ((101) & 63)))) | 1))) | 1))
                                    var v143 uint32 = v122
                                    _ = v143
                                }
                                if false {
                                    var v144 int16 = v9
                                    v9 = 1
                                    v144 -= (v144 ^ 1)
                                    if len(v3) > 0 {
                                        v3[0] = fi121
                                    }
                                    delete(v2, ("" + v4))
                                    _ = v144
                                } else {
                                    if true {
                                        fmt.Println(v11)
                                    }
                                    var v145 int16 = (v9 / (((v9 >> ((711) & 63))) | 1))
                                    fmt.Println(fi121, v11, fi121)
                                    fmt.Println(v10, v122, v4)
                                    _ = v145
                                }
                            }
                            var v146 int16 = (v9 | (v9 * (v9 ^ v9)))
                            var v147 int64 = -9223372036854775808
                            var v148 uint16 = 153
                            _ = v146
                            _ = v147
                            _ = v148
                        }
                    }
                }
                switch (^(fi14)) {
                case 32767:
                    v2[((v4 + v4) + (((v4 + v4) + "\n\t") + v4))] = (fi121 ^ fi121)
                case -256:
                    v149 := GenS6{F0: (^(v11))}
                    _ = v149
                case 18:
                    v1 -= 1e10
                }
                switch fi121 {
                case 128:
                    var v150 float64 = (-2.25 * 0.1)
                    _ = v150
                case 323:
                    _ = len(v3)
                default:
                    delete(v2, ("" + (v4 + (((v4 + v4) + v4) + "foo bar"))))
                }
                if (v9 > (v9 - (v9 % ((1) | 1)))) {
                    var v151 bool = false
                    _ = v151
                } else {
                    v4 = (((("" + v4) + (v4 + v4)) + "") + v4)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = v122
            }
            _ = fi121
        }
        var v152 int32 = 1
        _ = fi14
        _ = v152
    }
    delete(v2, (v4 + v4))
    fmt.Println(v4, v11, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
