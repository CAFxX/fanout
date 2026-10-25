package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint; F1 string; F2 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 uint8; F1 uint }
type GenS7 struct { F0 int8; F1 int32 }
type GenS8 struct { F0 uint; F1 int8; F2 string }
type GenS9 struct { F0 int8; F1 uint; F2 uint32 }
type GenS10 struct { F0 int; F1 uint64 }
type GenS11 struct { F0 int32; F1 int; F2 string }
type GenS12 struct { F0 bool; F1 int64 }
func genFn0(p0 bool, p1 bool, p2 int32) int32 {
    var v4 uint64 = 74
    if (!p0) {
        fmt.Println(v4, p2, p0)
    }
    var v5 int64 = -1
    if p1 {
        var v6 int = 439
        _ = v6
    } else {
        switch 1 {
        case -8:
            v7 := []int{128, 32768, 18}
            _ = v7
        case 349:
            var v8 uint64 = v4
            _ = v8
        default:
            v5 = -9223372036854775808
        }
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS1{F0: 133, F1: true, F2: 0}
        v4 += (v4 << ((2) & 63))
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (-(p2))
}

func main() {
    var v1 string = "\n\t"
    var v2 int16 = 621
    v3 := make(map[uint32]string)
    var v11 int16 = 1
    var v12 bool = true
    var v13 string = v1
    switch (v1 + (v13 + v13)) {
    case "\n\t":
        if true {
            switch 2147483647 {
            case 1:
                v1 = ("" + v13)
            }
            v14 := []int{}
            v12 = (-25 <= -218)
            _ = v14
        } else {
            if false {
                v15 := GenS1{F0: 701, F1: v12, F2: 529}
                v1 = v1
                _ = len(v3)
                _ = v15
            } else {
                if true {
                    fmt.Println(v12, v11, v12)
                } else {
                    v16 := GenS2{F0: 905, F1: "", F2: 838}
                    v1 = v13
                    _ = v16
                }
                fmt.Println(v13, v12)
                var v17 uint64 = 4294967295
                _ = v17
            }
        }
    case "foo bar":
        var v18 uint8 = 255
        _ = v18
    case "a":
        var v19 int32 = -261
        _ = v19
    default:
        var v20 int = 127
        _ = v20
    }
    v1 = ""
    if (!v12) {
        for fi21 := 0; fi21 < 6; fi21++ {
            fmt.Println(v12, v1, v1)
            v22 := GenS2{F0: 1, F1: v13, F2: 886}
            _ = fi21
            _ = v22
        }
        delete(v3, 4294967295)
        v13 = "a"
    } else {
        delete(v3, 0)
        v23 := []int{}
        switch 498 {
        case 151:
            if false {
                v12 = v12
            }
        case 228:
            if false {
                v1 = "foo bar"
                var v24 uint8 = 255
                var v25 float32 = (0.0 / 3.14159)
                _ = v24
                _ = v25
            } else {
                switch ((v13 + v13) + ((v13 + v13) + "hello")) {
                case "a":
                    v23 = append(v23, -877)
                case "":
                    if (true && (v12 && v12)) {
                        for fi26 := 0; fi26 < 20; fi26++ {
                            var v27 int64 = 0
                            switch fi26 {
                            case -2147483648:
                                v28 := GenS3{F0: v1}
                                _ = v28
                            case 42:
                                if (0 != 42) {
                                    v11 += 1
                                }
                            case 128:
                                fmt.Println(v2)
                            }
                            if len(v23) >= 2 {
                                v23 = v23[1:len(v23)]
                            }
                            _ = fi26
                            _ = v27
                        }
                        if ((1.5 < float32(1.5)) && (v12 && v12)) {
                            switch -848 {
                            case -9223372036854775808:
                                v11 = (v11 << ((255) & 63))
                            case 42:
                                var v29 uint32 = 958
                                _ = v29
                            }
                            for fi30 := 0; fi30 < 11; fi30++ {
                                v31 := GenS4{F0: 279}
                                delete(v3, 574)
                                _ = fi30
                                _ = v31
                            }
                            var v32 float32 = -2.25
                            _ = v32
                        } else {
                            var v33 int8 = 0
                            _ = v33
                        }
                        var v34 uint = 4294967295
                        _ = v34
                    }
                default:
                    switch v13 {
                    case "foo bar":
                        if v12 {
                            v35 := GenS4{F0: 598}
                            _ = v35
                        }
                    case "hello":
                        var v36 uint32 = 42
                        _ = v36
                    case "\n\t":
                        _ = len(v23)
                    }
                }
                delete(v3, 2147483647)
            }
        }
        _ = v23
    }
    for fi37 := 0; fi37 < 3; fi37++ {
        if ((fi37 <= -402) && (!v12)) {
            for fi38 := 0; fi38 < 1; fi38++ {
                switch "a" {
                case "hello":
                    switch fi37 {
                    case 42:
                        fmt.Println(v12)
                    case 786:
                        var v39 uint = 0
                        _ = v39
                    case 423:
                        fmt.Println(v1, fi37)
                    default:
                        if v12 {
                            if (-227 == -705) {
                                if (42 >= 4294967295) {
                                    var v40 int = (fi37 & fi38)
                                    _ = v40
                                }
                                v11 = 0
                                v11 = v2
                            }
                            var v41 uint16 = 32767
                            if (fi38 != int(v41)) {
                                v41 = v41
                            }
                            _ = v41
                        } else {
                            var v42 uint8 = 91
                            fmt.Println(v2, v1)
                            var v43 int = fi37
                            _ = v42
                            _ = v43
                        }
                    }
                case "\n\t":
                    if false {
                        if (v12 && (v12 && (!(!(0 < -840))))) {
                            v44 := []int{fi38}
                            fmt.Println(v2, v12, v1)
                            _ = v44
                        }
                        var v45 int8 = -128
                        _ = v45
                    }
                default:
                    var v46 uint64 = 18446744073709551615
                    _ = v46
                }
                var v47 uint8 = 150
                _ = fi38
                _ = v47
            }
            v11 += (v11 << ((42) & 63))
            for fi48 := 0; fi48 < 12; fi48++ {
                v11 = (v2 - (^(v2)))
                _ = fi48
            }
        }
        if (!v12) {
            for fi49 := 0; fi49 < 3; fi49++ {
                v2 = 32767
                _ = fi49
            }
        } else {
            var v50 bool = true
            v51 := GenS5{F0: -1, F1: "\n\t"}
            _ = v50
            _ = v51
        }
        v11 = 0
        _ = fi37
    }
    var v52 uint16 = 255
    if (-1 > -128) {
        switch "" {
        case "\n\t":
            v53 := v3[0]
            _ = v53
        case "foo bar":
            v52 = v52
        case "a":
            if (!v12) {
                switch 4294967295 {
                case 870:
                    v11 = (v2 & 25)
                case 0:
                    v54 := []int{-2147483648, -501, -2147483648, 4294967295}
                    _ = v54
                default:
                    v55 := []int{}
                    _ = v55
                }
                var v56 float64 = float64((0.1 - float32(0.0)))
                if (140 == 4294967295) {
                    var v57 uint64 = 118
                    v58 := v3[42]
                    v59 := []int{2147483647}
                    _ = v57
                    _ = v58
                    _ = v59
                } else {
                    v60 := GenS6{F0: 127, F1: 981}
                    if v12 {
                        var v61 uint = 18446744073709551615
                        switch 9223372036854775807 {
                        case 127:
                            switch v1 {
                            case "a":
                                switch ((v1 + "a") + "") {
                                case "foo bar":
                                    for fi62 := 0; fi62 < 12; fi62++ {
                                        fmt.Println(v52)
                                        v2 += int16(v52)
                                        fmt.Println(v11, fi62)
                                        _ = fi62
                                    }
                                case "\n\t":
                                    v63 := []int{145, -2147483648, 4294967295, 747}
                                    _ = v63
                                }
                            case "\n\t":
                                v11 = 1000
                            }
                        case 2147483647:
                            v64 := GenS7{F0: 0, F1: 2147483647}
                            _ = v64
                        }
                        _ = v61
                    } else {
                        v56 -= 1e10
                        v1 = v1
                    }
                    _ = v60
                }
                _ = v56
            }
        default:
            if true {
                v12 = (!(false || true))
            }
        }
        var v65 uint32 = 2147483647
        v66 := GenS1{F0: (v52 & (v52 & (v52 * (v52 * (^(v52)))))), F1: false, F2: (v65 % (((v65 & v65)) | 1))}
        _ = v65
        _ = v66
    }
    for fi67 := 0; fi67 < 9; fi67++ {
        if (-1 > 2147483647) {
            for fi68 := 0; fi68 < 6; fi68++ {
                fmt.Println(fi67, v2, v12)
                if (-210 != -2147483648) {
                    _ = len(v3)
                    v12 = v12
                } else {
                    v69 := []int{}
                    v70 := v3[2147483648]
                    _ = v69
                    _ = v70
                }
                _ = fi68
            }
            v71 := GenS6{F0: 12, F1: 65535}
            _ = v71
        }
        v11 = 54
        _ = fi67
    }
    for fi72 := 0; fi72 < 19; fi72++ {
        fmt.Println(v2, v12)
        v73 := GenS7{F0: 1, F1: -495}
        for fi74 := 0; fi74 < 3; fi74++ {
            var v75 bool = true
            if ((194 == 0) && ((-9223372036854775808 > 42) || v75)) {
                fmt.Println(fi72, fi74)
                if ((((v2 >> ((577) & 63)) != v11) || ((^(fi72)) > (fi74 + (^(fi74))))) && (((3.14159 * 3.14159) + 0.0) != -2.25)) {
                    var v76 bool = (!v75)
                    var v77 uint8 = 0
                    v77 -= v77
                    _ = v76
                    _ = v77
                }
            } else {
                v1 = "\n\t"
                v11 = v11
            }
            if v75 {
                v78 := []int{}
                v79 := GenS7{F0: 95, F1: 0}
                if ((-2.25 / 1.5) <= (float64(1e10) * -2.25)) {
                    if (!v12) {
                        switch "foo bar" {
                        case "":
                            v80 := GenS5{F0: 975, F1: v1}
                            _ = v80
                        }
                    } else {
                        switch fi72 {
                        case -9223372036854775808:
                            if (18446744073709551615 < 2) {
                                var v81 float64 = (float64(-2.25) / 0.1)
                                fmt.Println(v1, v13, v13)
                                v82 := GenS3{F0: v1}
                                v83 := GenS3{F0: v13}
                                switch fi74 {
                                case -2147483648:
                                    if (-128 >= -128) {
                                        if v12 {
                                            if (127 < 222) {
                                                var v84 int16 = 32767
                                                if len(v78) > 0 {
                                                    v78[0] = fi74
                                                }
                                                v12 = (-1 < -126)
                                                var v85 uint32 = 42
                                                _ = v84
                                                _ = v85
                                            }
                                            switch (fi72 | -9223372036854775808) {
                                            case 4294967295:
                                                switch v13 {
                                                case "hello":
                                                    if v12 {
                                                        v86 := GenS8{F0: 65535, F1: 81, F2: (v13 + ("hello" + v1))}
                                                        v87 := GenS9{F0: 100, F1: 255, F2: 58}
                                                        var v88 float64 = ((v81 + v81) + (v81 - v81))
                                                        _ = v86
                                                        _ = v87
                                                        _ = v88
                                                    } else {
                                                        v81 += 0.1
                                                    }
                                                case "":
                                                    var v89 uint32 = 42
                                                    _ = v89
                                                }
                                            case 32768:
                                                v11 = 0
                                            case 32767:
                                                if ((false || true) || ("\n\t" != (v13 + v13))) {
                                                    fmt.Println(fi72)
                                                    v11 = 32767
                                                }
                                            default:
                                                if (!true) {
                                                    fmt.Println(v75, v12, fi72)
                                                    switch (fi72 | 127) {
                                                    case -2147483648:
                                                        v12 = true
                                                    }
                                                    v1 = ""
                                                }
                                            }
                                        }
                                        fmt.Println(v75, v12, v81)
                                    } else {
                                        v90 := GenS10{F0: (fi72 >> ((4294967295) & 63)), F1: 4294967295}
                                        if ((v12 || false) && (-37 != 42)) {
                                            v3[789] = v1
                                            var v91 float64 = float64((1.5 / (3.14159 / float32(v81))))
                                            _ = v91
                                        }
                                        _ = v90
                                    }
                                default:
                                    v75 = ((1.5 * 3.14159) < (0.0 * -2.25))
                                }
                                _ = v81
                                _ = v82
                                _ = v83
                            }
                        default:
                            v92 := GenS9{F0: 119, F1: 739, F2: 277}
                            _ = v92
                        }
                        v93 := GenS11{F0: genFn0(v12, ((^(v11)) != v2), -2147483648), F1: fi72, F2: (v13 + "a")}
                        v11 = int16(fi72)
                        _ = v93
                    }
                    var v94 float32 = 1.5
                    fmt.Println(v2, fi72)
                    _ = v94
                } else {
                    v95 := v3[2]
                    switch (fi72 * (fi74 & -9223372036854775808)) {
                    case 0:
                        if (227 > 42) {
                            var v96 int = fi72
                            fmt.Println(v95, fi74)
                            var v97 uint64 = 0
                            _ = v96
                            _ = v97
                        } else {
                            var v98 uint8 = 255
                            if v12 {
                                fmt.Println(v1, v11)
                                if v12 {
                                    v12 = v12
                                    v75 = v12
                                }
                                var v99 int64 = -245
                                v13 = ((("hello" + v1) + ("hello" + v95)) + ("a" + ("\n\t" + v95)))
                                _ = v99
                            }
                            if len(v78) >= 2 {
                                v78 = v78[1:len(v78)]
                            }
                            _ = v98
                        }
                    case 42:
                        var v100 int = (fi74 + fi74)
                        _ = v100
                    default:
                        var v101 uint8 = 2
                        _ = v101
                    }
                    var v102 uint16 = v52
                    var v103 int32 = 1
                    var v104 bool = v75
                    _ = v102
                    _ = v103
                    _ = v104
                    _ = v95
                }
                _ = v78
                _ = v79
            } else {
                var v105 int16 = (v11 + (^(v2)))
                fmt.Println(v75, v52, v105)
                fmt.Println(v2, v12, v75)
                _ = v105
            }
            v106 := []int{(fi72 / ((-256) | 1))}
            v52 = v52
            _ = fi74
            _ = v106
            _ = v75
        }
        _ = fi72
        _ = v73
    }
    for fi107 := 0; fi107 < 11; fi107++ {
        var v108 int = 2147483647
        fmt.Println(v12, v108)
        _ = fi107
        _ = v108
    }
    var v109 string = "a"
    v12 = ((!true) && v12)
    fmt.Println(v52, v12, v2)
    if (0 <= 400) {
        v52 += v52
        if (!true) {
            var v110 uint64 = 263
            if (v12 && false) {
                var v111 uint16 = (v52 - 255)
                v112 := []int{-801, -2147483648, 4294967295}
                if len(v112) > 0 {
                    v112[0] = 13
                }
                _ = v111
                _ = v112
            }
            var v113 uint16 = 1
            _ = v110
            _ = v113
        }
        v11 += (v11 * int16(v52))
    }
    v12 = v12
    switch "\n\t" {
    case "":
        for fi114 := 0; fi114 < 9; fi114++ {
            var v115 uint64 = 4294967295
            switch ("\n\t" + (v1 + (((v13 + (v13 + v109)) + v13) + (v13 + v13)))) {
            case "a":
                fmt.Println(v109, v13, v52)
            case "hello":
                if v12 {
                    var v116 int64 = 9223372036854775807
                    switch fi114 {
                    case 32767:
                        for fi117 := 0; fi117 < 2; fi117++ {
                            var v118 uint32 = 2147483648
                            _ = fi117
                            _ = v118
                        }
                    }
                    _ = v116
                } else {
                    v52 = 1
                    v52 += (v52 % ((v52) | 1))
                }
            default:
                var v119 float32 = 0.1
                _ = v119
            }
            _ = fi114
            _ = v115
        }
    case "hello":
        v120 := GenS12{F0: (316 >= 65535), F1: 414}
        _ = v120
    }
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v52:", v52)
    _ = v1
    _ = v109
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v52
}
