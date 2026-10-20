package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64; F1 int16 }
type GenS8 struct { F0 int8; F1 int16; F2 string }
type GenS9 struct { F0 uint; F1 uint; F2 uint }
type GenS10 struct { F0 bool; F1 int32 }
type GenS11 struct { F0 bool; F1 uint }
type GenS12 struct { F0 int32; F1 int8 }
type GenS13 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS14 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS15 struct { F0 int16; F1 int16 }
type GenS16 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS17 struct { F0 int8; F1 uint; F2 int }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 int8; F1 bool; F2 bool }
type GenS20 struct { F0 string; F1 uint32 }
func genFn0() uint16 {
    if (!false) {
        var v5 int = 4294967295
        fmt.Println(v5)
        var v6 int = -341
        _ = v5
        _ = v6
    } else {
        var v7 string = "\n\t"
        _ = v7
    }
    var v8 int = 32768
    fmt.Println(v8)
    if ((42 >= 225) && true) {
        v9 := []int{42, v8}
        fmt.Println(v8)
        if false {
            v10 := GenS1{F0: -1, F1: 127, F2: 2}
            _ = v10
        }
        _ = v9
    } else {
        var v11 uint16 = 65535
        for fi12 := 0; fi12 < 19; fi12++ {
            for fi13 := 0; fi13 < 2; fi13++ {
                var v14 int64 = 42
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
        _ = v11
    }
    if true {
        v8 = v8
        var v15 uint16 = 213
        var v16 int64 = -9223372036854775808
        _ = v15
        _ = v16
    }
    _ = v8
    return 1
}
func genFn1(p0 uint8, p1 uint) uint32 {
    if true {
        v17 := []int{int(p0)}
        fmt.Println(p1)
        var v18 int16 = 42
        _ = v17
        _ = v18
    }
    p0 -= 127
    if true {
        v19 := []int{-898, -2147483648, 2147483647, 42}
        if len(v19) > 0 {
            v19[0] = 180
        }
        _ = v19
    }
    v20 := GenS1{F0: -32768, F1: -50, F2: p0}
    for fi21 := 0; fi21 < 5; fi21++ {
        for fi22 := 0; fi22 < 6; fi22++ {
            if false {
                if ((42 > 16) && (p0 == p0)) {
                    var v23 float64 = 0.1
                    _ = v23
                } else {
                    v24 := GenS2{F0: p1, F1: -37}
                    _ = v24
                }
                switch -2147483648 {
                case -95:
                    var v25 uint16 = genFn0()
                    _ = v25
                default:
                    if false {
                        var v26 uint16 = 403
                        v27 := []int{fi21, -1, fi21, -1}
                        var v28 int = fi21
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                }
            }
            var v29 uint64 = 42
            _ = fi22
            _ = v29
        }
        var v30 uint = (p1 ^ p1)
        if (217 < 0) {
            fmt.Println(p0)
            v31 := []int{fi21, int(p0), (^(fi21))}
            _ = v31
        }
        _ = fi21
        _ = v30
    }
    _ = p0
    _ = p1
    _ = v20
    return 1
}

func main() {
    var v1 uint16 = 42
    v2 := []int{2147483647, 4294967295, 1, 32768}
    var v3 uint32 = 42
    v4 := []uint32{(^(v3))}
    v1 += v1
    if ((!false) && (42 < 1)) {
        switch ("hello" + "\n\t") {
        case "hello":
            for fi32 := 0; fi32 < 4; fi32++ {
                v33 := GenS3{F0: int8(fi32), F1: -32768}
                switch "hello" {
                case "hello":
                    v3 = uint32(v1)
                case "a":
                    v34 := GenS4{F0: (^(v3)), F1: false}
                    _ = v34
                case "foo bar":
                    v3 += 880
                default:
                    if (true || (uint8(v1) < 127)) {
                        fmt.Println(v3, fi32)
                    }
                }
                _ = fi32
                _ = v33
            }
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    } else {
        v35 := GenS5{F0: 0}
        if (699 < 2) {
            for fi36 := 0; fi36 < 15; fi36++ {
                var v37 uint = 42
                v38 := GenS6{F0: 12}
                _ = fi36
                _ = v37
                _ = v38
            }
            v39 := GenS1{F0: -1, F1: -128, F2: 128}
            _ = v39
        } else {
            v1 -= genFn0()
            var v40 uint64 = 2
            _ = v40
        }
        if (true || false) {
            switch "foo bar" {
            case "hello":
                switch "a" {
                case "":
                    if (-32768 != 0) {
                        if true {
                            if (!(genFn0() != 784)) {
                                v3 -= (v3 | (^(v3)))
                            } else {
                                var v41 float64 = (0.1 + (-2.25 + (0.1 - 3.14159)))
                                v42 := GenS5{F0: v3}
                                switch "hello" {
                                case "a":
                                    var v43 string = "\n\t"
                                    _ = v43
                                }
                                _ = v41
                                _ = v42
                            }
                            v1 = v1
                        } else {
                            if (255 < 128) {
                                if false {
                                    for fi44 := 0; fi44 < 18; fi44++ {
                                        for fi45 := 0; fi45 < 7; fi45++ {
                                            v1 -= v1
                                            switch fi45 {
                                            case 127:
                                                if (!(!(4294967295 <= 1))) {
                                                    var v46 string = "foo bar"
                                                    v47 := GenS4{F0: v3, F1: true}
                                                    _ = v46
                                                    _ = v47
                                                } else {
                                                    fmt.Println(fi44, fi44, fi44)
                                                    fmt.Println(v3, fi45)
                                                }
                                            case 2147483647:
                                                v3 = genFn1(127, 18446744073709551615)
                                                var v48 uint = 42
                                                var v49 uint = (^(v48))
                                                _ = v48
                                                _ = v49
                                            default:
                                                if (!false) {
                                                    if true {
                                                        _ = len(v2)
                                                    } else {
                                                        fmt.Println(fi45)
                                                        fmt.Println(fi44, fi44)
                                                    }
                                                    switch 127 {
                                                    case 1:
                                                        var v50 int32 = -2147483648
                                                        _ = v50
                                                    case 319:
                                                        switch fi44 {
                                                        case 427:
                                                            v1 = genFn0()
                                                        case -2147483648:
                                                            switch fi44 {
                                                            case -632:
                                                                switch 2147483647 {
                                                                case 1:
                                                                    if false {
                                                                        var v51 uint = 255
                                                                        v52 := GenS7{F0: int64(fi45), F1: 42}
                                                                        switch "\n\t" {
                                                                        case "":
                                                                            v1 = 2
                                                                        case "foo bar":
                                                                            switch fi45 {
                                                                            case 2147483647:
                                                                                if (32767 < 42) {
                                                                                    v3 = v3
                                                                                }
                                                                            case 9223372036854775807:
                                                                                var v53 int64 = 233
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                var v54 uint = v51
                                                                                if false {
                                                                                    fmt.Println(fi44, fi44)
                                                                                    v1 = (v1 / (((v1 >> ((v54) & 63))) | 1))
                                                                                    v53 += (v53 - (v53 | v53))
                                                                                } else {
                                                                                    if false {
                                                                                        var v55 int32 = 156
                                                                                        _ = v55
                                                                                    }
                                                                                }
                                                                                _ = v53
                                                                                _ = v54
                                                                            case 42:
                                                                                var v56 bool = true
                                                                                _ = v56
                                                                            default:
                                                                                var v57 int8 = 40
                                                                                _ = v57
                                                                            }
                                                                        case "hello":
                                                                            var v58 int8 = 42
                                                                            _ = v58
                                                                        }
                                                                        _ = v51
                                                                        _ = v52
                                                                    } else {
                                                                        var v59 int8 = -52
                                                                        _ = v59
                                                                    }
                                                                default:
                                                                    if (false && true) {
                                                                        fmt.Println(v1, v3, fi44)
                                                                    } else {
                                                                        if (128 < (fi44 & fi44)) {
                                                                            var v60 bool = (true || true)
                                                                            v3 += 4294967295
                                                                            if v60 {
                                                                                v61 := GenS4{F0: v3, F1: v60}
                                                                                switch 127 {
                                                                                case 498:
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = v3
                                                                                    }
                                                                                case 898:
                                                                                    if ((!v60) && v60) {
                                                                                        fmt.Println(fi45, v60, fi45)
                                                                                        var v62 uint32 = (v3 << ((950) & 63))
                                                                                        v3 = (^(v3))
                                                                                        _ = v62
                                                                                    } else {
                                                                                        var v63 int32 = 42
                                                                                        switch "hello" {
                                                                                        case "a":
                                                                                            if v60 {
                                                                                                switch fi44 {
                                                                                                case 9223372036854775807:
                                                                                                    v2 = append(v2, 128)
                                                                                                case 4294967295:
                                                                                                    var v64 int = fi44
                                                                                                    _ = v64
                                                                                                }
                                                                                                var v65 int64 = -9223372036854775808
                                                                                                _ = v65
                                                                                            } else {
                                                                                                v66 := GenS1{F0: 0, F1: -128, F2: 42}
                                                                                                _ = v66
                                                                                            }
                                                                                        case "hello":
                                                                                            var v67 float64 = 1.5
                                                                                            _ = v67
                                                                                        default:
                                                                                            switch ("foo bar" + "\n\t") {
                                                                                            case "":
                                                                                                v68 := GenS8{F0: -1, F1: -1, F2: ((("\n\t" + "a") + "hello") + "a")}
                                                                                                _ = v68
                                                                                            default:
                                                                                                if true {
                                                                                                    v69 := GenS4{F0: (v3 >> ((1) & 63)), F1: false}
                                                                                                    _ = v69
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        v4 = append(v4, v3)
                                                                                        _ = v63
                                                                                    }
                                                                                default:
                                                                                    fmt.Println(v60, fi44, v3)
                                                                                }
                                                                                var v70 int16 = 320
                                                                                _ = v61
                                                                                _ = v70
                                                                            }
                                                                            _ = v60
                                                                        } else {
                                                                            var v71 uint = 1
                                                                            v72 := GenS7{F0: -9223372036854775808, F1: -16}
                                                                            v73 := GenS3{F0: 0, F1: 1000}
                                                                            _ = v71
                                                                            _ = v72
                                                                            _ = v73
                                                                        }
                                                                        v74 := GenS8{F0: -128, F1: -677, F2: ""}
                                                                        _ = v74
                                                                    }
                                                                }
                                                            }
                                                        default:
                                                            v75 := GenS9{F0: 18446744073709551615, F1: 1, F2: 62}
                                                            _ = v75
                                                        }
                                                    case -1:
                                                        v76 := GenS10{F0: (!true), F1: 42}
                                                        _ = v76
                                                    }
                                                    var v77 bool = true
                                                    _ = v77
                                                }
                                            }
                                            v4 = append(v4, genFn1(127, 0))
                                            _ = fi45
                                        }
                                        _ = fi44
                                    }
                                    if true {
                                        var v78 int8 = -31
                                        var v79 float64 = 1e10
                                        _ = v78
                                        _ = v79
                                    } else {
                                        v3 += 42
                                        var v80 int8 = -1
                                        _ = v80
                                    }
                                    if true {
                                        v2 = append(v2, 565)
                                        for fi81 := 0; fi81 < 19; fi81++ {
                                            v3 = genFn1(1, 626)
                                            var v82 uint8 = 1
                                            var v83 int8 = 0
                                            _ = fi81
                                            _ = v82
                                            _ = v83
                                        }
                                        switch "a" {
                                        case "hello":
                                            _ = len(v2)
                                        case "\n\t":
                                            v1 += (v1 | 482)
                                        case "foo bar":
                                            if true {
                                                var v84 float64 = 1.5
                                                v85 := GenS11{F0: (false && true), F1: 1}
                                                fmt.Println(v3)
                                                _ = v84
                                                _ = v85
                                            } else {
                                                v2 = append(v2, 0)
                                                v3 += 31
                                            }
                                        }
                                    } else {
                                        if true {
                                            fmt.Println(v1)
                                            for fi86 := 0; fi86 < 3; fi86++ {
                                                v3 = genFn1(255, 4294967295)
                                                _ = fi86
                                            }
                                        }
                                        if (!false) {
                                            switch 2147483647 {
                                            case 128:
                                                switch 9223372036854775807 {
                                                case 128:
                                                    v87 := GenS12{F0: 42, F1: 0}
                                                    _ = v87
                                                }
                                            }
                                            v3 = v3
                                        } else {
                                            var v88 string = "a"
                                            var v89 int64 = -1
                                            _ = v88
                                            _ = v89
                                        }
                                    }
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(v3, v1)
                                    v90 := GenS12{F0: 972, F1: 100}
                                    _ = v90
                                }
                                fmt.Println(v1)
                            } else {
                                var v91 int32 = 155
                                if (0.1 < 1.5) {
                                    if len(v2) > 0 {
                                        v2[0] = 0
                                    }
                                    for fi92 := 0; fi92 < 12; fi92++ {
                                        var v93 int16 = 16
                                        var v94 uint8 = 2
                                        _ = fi92
                                        _ = v93
                                        _ = v94
                                    }
                                    v4 = append(v4, v3)
                                }
                                _ = v91
                            }
                            if (-13 != -1) {
                                fmt.Println(v1)
                                if ((1 > -1) || false) {
                                    var v95 uint8 = 1
                                    v96 := GenS13{F0: 127, F1: 734, F2: 252}
                                    v1 = 1
                                    _ = v95
                                    _ = v96
                                } else {
                                    v3 = 2147483648
                                    fmt.Println(v3, v3)
                                }
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = 1
                                }
                            }
                            var v97 uint64 = 440
                            _ = v97
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        switch -363 {
                        case 0:
                            var v98 string = (("" + "") + "\n\t")
                            _ = v98
                        case 32767:
                            switch "" {
                            case "":
                                var v99 uint64 = 70
                                _ = v99
                            case "hello":
                                fmt.Println(v3)
                            case "a":
                                var v100 int16 = 621
                                _ = v100
                            }
                        case 128:
                            for fi101 := 0; fi101 < 5; fi101++ {
                                switch (^(fi101)) {
                                case 4294967295:
                                    if ((272 == 1) || (4294967295 != 451)) {
                                        if len(v2) > 0 {
                                            v2[0] = (fi101 >> ((42) & 63))
                                        }
                                    } else {
                                        switch 128 {
                                        case 416:
                                            for fi102 := 0; fi102 < 15; fi102++ {
                                                v1 += (v1 | genFn0())
                                                fmt.Println(v3, fi101, v1)
                                                var v103 uint = 42
                                                _ = fi102
                                                _ = v103
                                            }
                                        case 32768:
                                            if (18446744073709551615 == 4294967295) {
                                                fmt.Println(fi101, fi101, v1)
                                                _ = len(v4)
                                                for fi104 := 0; fi104 < 11; fi104++ {
                                                    if len(v4) > 0 {
                                                        v4[0] = (v3 & (v3 / ((v3) | 1)))
                                                    }
                                                    if true {
                                                        if (39 != -128) {
                                                            v2 = append(v2, fi104)
                                                        } else {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            fmt.Println(fi104)
                                                            if (32767 >= -32768) {
                                                                v1 += v1
                                                                v3 -= v3
                                                            }
                                                        }
                                                        v105 := GenS14{F0: (v3 + 2), F1: -815, F2: v3}
                                                        _ = v105
                                                    }
                                                    v106 := GenS2{F0: 0, F1: -85}
                                                    _ = fi104
                                                    _ = v106
                                                }
                                            }
                                        case 0:
                                            fmt.Println(fi101, fi101, v1)
                                        default:
                                            if false {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                fmt.Println(v1, v1, fi101)
                                            } else {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                v107 := GenS15{F0: 0, F1: 1000}
                                                v3 -= uint32(v1)
                                                _ = v107
                                            }
                                        }
                                        fmt.Println(v3, v1)
                                    }
                                case -2147483648:
                                    if (!false) {
                                        if (!true) {
                                            if true {
                                                if false {
                                                    for fi108 := 0; fi108 < 4; fi108++ {
                                                        switch fi101 {
                                                        case 42:
                                                            v3 -= (^(v3))
                                                            var v109 uint64 = 4294967295
                                                            _ = v109
                                                        case 487:
                                                            if false {
                                                                fmt.Println(v3)
                                                                _ = len(v4)
                                                            }
                                                        default:
                                                            var v110 uint8 = 77
                                                            _ = v110
                                                        }
                                                        _ = fi108
                                                    }
                                                    var v111 uint32 = 2147483648
                                                    _ = v111
                                                } else {
                                                    v3 = 2
                                                    if (true || (!false)) {
                                                        switch (fi101 / (((fi101 % ((int(v1)) | 1))) | 1)) {
                                                        case -123:
                                                            var v112 float32 = 1e10
                                                            _ = v112
                                                        default:
                                                            for fi113 := 0; fi113 < 16; fi113++ {
                                                                var v114 bool = true
                                                                var v115 string = "hello"
                                                                _ = fi113
                                                                _ = v114
                                                                _ = v115
                                                            }
                                                        }
                                                        if true {
                                                            for fi116 := 0; fi116 < 18; fi116++ {
                                                                fmt.Println(v1)
                                                                v117 := GenS16{F0: genFn0(), F1: 18446744073709551615, F2: -1}
                                                                var v118 bool = ((^(fi116)) <= 968)
                                                                _ = fi116
                                                                _ = v117
                                                                _ = v118
                                                            }
                                                            v3 += v3
                                                        } else {
                                                            var v119 uint16 = (v1 << ((65535) & 63))
                                                            _ = v119
                                                        }
                                                    } else {
                                                        var v120 float64 = 1.5
                                                        if true {
                                                            if ((0.0 * (float32(3.14159) + 0.0)) < (0.0 - -2.25)) {
                                                                if (-15 != 1000) {
                                                                    switch (fi101 % (((fi101 >> ((762) & 63))) | 1)) {
                                                                    case 1:
                                                                        v3 += 1
                                                                    case 610:
                                                                        for fi121 := 0; fi121 < 10; fi121++ {
                                                                            var v122 uint16 = 42
                                                                            _ = fi121
                                                                            _ = v122
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            v1 -= (^(v1))
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        } else {
                                                            fmt.Println(fi101, v120, v3)
                                                            if (128 != 184) {
                                                                for fi123 := 0; fi123 < 2; fi123++ {
                                                                    switch -483 {
                                                                    case -75:
                                                                        fmt.Println(v3, fi101)
                                                                    }
                                                                    _ = fi123
                                                                }
                                                            }
                                                            switch ("\n\t" + ("" + "hello")) {
                                                            case "hello":
                                                                v124 := GenS5{F0: (v3 / ((0) | 1))}
                                                                _ = v124
                                                            default:
                                                                if false {
                                                                    for fi125 := 0; fi125 < 16; fi125++ {
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        fmt.Println(v1)
                                                                        var v126 float32 = 1e10
                                                                        _ = fi125
                                                                        _ = v126
                                                                    }
                                                                    var v127 int = -280
                                                                    for fi128 := 0; fi128 < 12; fi128++ {
                                                                        fmt.Println(v127)
                                                                        _ = fi128
                                                                    }
                                                                    _ = v127
                                                                } else {
                                                                    if (!(-1 != -128)) {
                                                                        var v129 uint8 = 1
                                                                        _ = v129
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = v120
                                                    }
                                                }
                                            } else {
                                                var v130 uint = 338
                                                v1 += (v1 / (((v1 | v1)) | 1))
                                                if ((v1 - genFn0()) != v1) {
                                                    v3 += (^(v3))
                                                } else {
                                                    v131 := GenS17{F0: -92, F1: uint(v1), F2: fi101}
                                                    v130 += (v130 % ((v130) | 1))
                                                    _ = v131
                                                }
                                                _ = v130
                                            }
                                            if true {
                                                var v132 uint = 2
                                                v2 = append(v2, int(v1))
                                                var v133 int64 = -9223372036854775808
                                                _ = v132
                                                _ = v133
                                            } else {
                                                if true {
                                                    switch (^(fi101)) {
                                                    case -2147483648:
                                                        var v134 int16 = 244
                                                        _ = v134
                                                    case 0:
                                                        var v135 float64 = (((1.5 / 1e10) / 3.14159) * (0.0 - (0.0 * (-2.25 + 1e10))))
                                                        _ = v135
                                                    default:
                                                        fmt.Println(fi101)
                                                    }
                                                    v3 -= (^(v3))
                                                    if (622 >= 18446744073709551615) {
                                                        v2 = append(v2, fi101)
                                                    } else {
                                                        if len(v2) > 0 {
                                                            v2[0] = fi101
                                                        }
                                                    }
                                                } else {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        fmt.Println(v1, fi101)
                                                    default:
                                                        fmt.Println(fi101, fi101, v1)
                                                    }
                                                    var v136 int64 = 0
                                                    _ = v136
                                                }
                                            }
                                            v3 = genFn1(1, 193)
                                        } else {
                                            var v137 uint8 = 1
                                            _ = v137
                                        }
                                        switch fi101 {
                                        case -930:
                                            var v138 bool = true
                                            _ = v138
                                        }
                                    }
                                case 32767:
                                    switch "foo bar" {
                                    case "":
                                        v139 := GenS18{F0: 100}
                                        _ = v139
                                    default:
                                        v1 = (^(v1))
                                    }
                                default:
                                    fmt.Println(v1, v1, v3)
                                }
                                _ = fi101
                            }
                        }
                    } else {
                        var v140 int16 = -415
                        if false {
                            v140 = -534
                            var v141 float64 = (0.0 - 1.5)
                            if (0 != 2) {
                                v2 = append(v2, 1)
                                v1 = 32767
                            }
                            _ = v141
                        } else {
                            if ((v140 <= (v140 + (v140 >> ((4294967295) & 63)))) && true) {
                                v142 := GenS8{F0: 1, F1: (v140 ^ (v140 * v140)), F2: ((("a" + ("a" + ("hello" + ("" + "hello")))) + "foo bar") + ("" + "hello"))}
                                v143 := GenS8{F0: 10, F1: (^(v140)), F2: ("\n\t" + ("foo bar" + "a"))}
                                _ = v142
                                _ = v143
                            } else {
                                var v144 uint32 = genFn1(128, 2)
                                if false {
                                    switch 992 {
                                    case -9223372036854775808:
                                        for fi145 := 0; fi145 < 18; fi145++ {
                                            for fi146 := 0; fi146 < 17; fi146++ {
                                                if false {
                                                    if len(v2) > 0 {
                                                        v2[0] = (-(fi146))
                                                    }
                                                    v140 = 1
                                                    var v147 uint8 = 111
                                                    _ = v147
                                                } else {
                                                    var v148 uint = 460
                                                    v3 = 0
                                                    v144 += (v144 % ((v144) | 1))
                                                    _ = v148
                                                }
                                                fmt.Println(v140, v3, v3)
                                                _ = fi146
                                            }
                                            var v149 int64 = 42
                                            _ = fi145
                                            _ = v149
                                        }
                                    case 248:
                                        switch 0 {
                                        case 128:
                                            switch "\n\t" {
                                            case "\n\t":
                                                fmt.Println(v1)
                                            case "foo bar":
                                                var v150 float32 = (0.1 / (-2.25 + float32((1e10 * float64(((1.5 + 3.14159) - -2.25))))))
                                                _ = v150
                                            default:
                                                var v151 int32 = 2147483647
                                                _ = v151
                                            }
                                        case 99:
                                            fmt.Println(v3, v144, v144)
                                        }
                                    case 32768:
                                        switch -415 {
                                        case -630:
                                            v152 := GenS6{F0: 66}
                                            _ = v152
                                        case 32767:
                                            if len(v4) > 0 {
                                                v4[0] = 42
                                            }
                                        }
                                    default:
                                        if (127 > 255) {
                                            for fi153 := 0; fi153 < 18; fi153++ {
                                                if len(v2) > 0 {
                                                    v2[0] = (fi153 & 127)
                                                }
                                                var v154 int64 = 1
                                                _ = fi153
                                                _ = v154
                                            }
                                        }
                                    }
                                    v140 -= int16(v1)
                                }
                                _ = v144
                            }
                        }
                        if (!(!true)) {
                            v140 = v140
                            var v155 int32 = 2147483647
                            var v156 int8 = 1
                            _ = v155
                            _ = v156
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            for fi157 := 0; fi157 < 8; fi157++ {
                                for fi158 := 0; fi158 < 1; fi158++ {
                                    v159 := GenS18{F0: -109}
                                    _ = fi158
                                    _ = v159
                                }
                                _ = fi157
                            }
                            v1 = 2
                        }
                        _ = v140
                    }
                case "\n\t":
                    v1 -= 0
                case "hello":
                    fmt.Println(v1, v1)
                default:
                    for fi160 := 0; fi160 < 14; fi160++ {
                        v161 := GenS19{F0: 42, F1: true, F2: (0 >= 42)}
                        var v162 string = ""
                        _ = fi160
                        _ = v161
                        _ = v162
                    }
                }
            }
            v163 := GenS12{F0: -919, F1: 127}
            _ = v163
        } else {
            var v164 uint8 = 42
            _ = v164
        }
        _ = v35
    }
    v1 = 595
    var v165 uint32 = (v3 + v3)
    v4 = append(v4, 2147483648)
    var v166 int64 = -845
    v3 = v165
    v166 = (v166 & v166)
    v166 -= -1
    if false {
        v167 := GenS5{F0: (v3 & v165)}
        _ = v167
    }
    v168 := GenS20{F0: "\n\t", F1: genFn1(42, 2)}
    var v169 uint32 = (v3 & v165)
    v170 := GenS7{F0: (v166 >> ((255) & 63)), F1: -503}
    if false {
        v1 += 703
    } else {
        var v171 int = 496
        _ = v171
    }
    fmt.Println(v166, v166)
    v1 = (v1 % ((uint16(v3)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v165:", v165)
    fmt.Println("v166:", v166)
    fmt.Println("v169:", v169)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v165
    _ = v166
    _ = v168
    _ = v169
    _ = v170
    _ = v2
    _ = v3
    _ = v4
}
