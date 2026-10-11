package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint; F2 int }
type GenS3 struct { F0 uint64; F1 string; F2 uint64 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 int8; F1 int64 }
type GenS7 struct { F0 bool; F1 int64; F2 bool }
type GenS8 struct { F0 int; F1 bool; F2 int64 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int64; F1 bool }
type GenS11 struct { F0 int; F1 uint16; F2 int8 }
type GenS12 struct { F0 uint64; F1 int; F2 int8 }
type GenS13 struct { F0 uint64; F1 int8 }
type GenS14 struct { F0 int64; F1 int8; F2 int }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 int64; F1 int8 }
type GenS17 struct { F0 bool }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 0
    v3 := []string{}
    var v4 string = "foo bar"
    v5 := make(map[int]uint32)
    var v6 uint = 644
    v6 = uint(v2)
    if true {
        var v7 string = v4
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := GenS1{F0: (false && true)}
            v6 -= v6
            _ = fi8
            _ = v9
        }
        v6 = 2
        _ = v7
    }
    var v10 int64 = -670
    switch "a" {
    case "a":
        switch "a" {
        case "a":
            if ((-536 >= -2147483648) && false) {
                if (0 != 42) {
                    _ = len(v3)
                } else {
                    v10 = v10
                    v11 := GenS2{F0: 65535, F1: 483, F2: 32767}
                    _ = v11
                }
                if false {
                    var v12 uint = (v6 | v6)
                    for fi13 := 0; fi13 < 11; fi13++ {
                        for fi14 := 0; fi14 < 4; fi14++ {
                            fmt.Println(v10, fi13)
                            v15 := GenS3{F0: v2, F1: (v4 + ""), F2: (^(v2))}
                            _ = fi14
                            _ = v15
                        }
                        for fi16 := 0; fi16 < 7; fi16++ {
                            v12 += v12
                            _ = fi16
                        }
                        var v17 uint64 = (^(v2))
                        _ = fi13
                        _ = v17
                    }
                    _ = v12
                }
                for fi18 := 0; fi18 < 10; fi18++ {
                    switch v4 {
                    case "a":
                        v19 := v5[fi18]
                        _ = v19
                    case "foo bar":
                        v6 -= v6
                    }
                    var v20 uint16 = 65535
                    v6 -= (v6 & v6)
                    _ = fi18
                    _ = v20
                }
            }
        case "\n\t":
            for fi21 := 0; fi21 < 9; fi21++ {
                for fi22 := 0; fi22 < 19; fi22++ {
                    if (false || false) {
                        var v23 uint64 = v2
                        v10 += (v10 ^ v10)
                        v4 = (("\n\t" + v4) + "\n\t")
                        v24 := GenS3{F0: v23, F1: ((v4 + (("hello" + ((((v4 + ("a" + v4)) + "a") + (v4 + v4)) + "hello")) + v4)) + ("" + v4)), F2: (v23 + 0)}
                        fmt.Println(fi21, fi22)
                        _ = v23
                        _ = v24
                    }
                    _ = fi22
                }
                _ = fi21
            }
        case "foo bar":
            if (true || false) {
                if true {
                    fmt.Println(v6)
                } else {
                    v6 -= (v6 ^ v6)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                v4 = v4
                v25 := GenS2{F0: 0, F1: v6, F2: 127}
                _ = v25
            } else {
                _ = len(v3)
                var v26 int = 4294967295
                v10 = 42
                _ = v26
            }
        default:
            for fi27 := 0; fi27 < 19; fi27++ {
                v28 := GenS3{F0: 18446744073709551615, F1: v4, F2: 0}
                if (false && true) {
                    var v29 int = -2147483648
                    v4 = v4
                    v4 = v4
                    _ = v29
                } else {
                    if (!((fi27 % (((fi27 >> ((v6) & 63))) | 1)) != 2147483647)) {
                        var v30 uint8 = 49
                        _ = v30
                    }
                    v31 := GenS3{F0: 875, F1: v4, F2: 42}
                    _ = v31
                }
                v10 -= (v10 >> ((v6) & 63))
                _ = fi27
                _ = v28
            }
        }
    case "\n\t":
        fmt.Println(v10, v10)
    default:
        if true {
            v32 := v1[v4]
            _ = v32
        }
    }
    v3 = append(v3, v4)
    fmt.Println(v4, v10)
    for fi33 := 0; fi33 < 20; fi33++ {
        delete(v5, 128)
        v6 = 822
        _ = fi33
    }
    fmt.Println(v2)
    fmt.Println(v6)
    switch 361 {
    case 446:
        var v34 int = 228
        _ = v34
    case -1:
        v3 = append(v3, v4)
    case 0:
        if true {
            fmt.Println(v4, v4)
            var v35 uint16 = 32767
            if len(v3) > 0 {
                v3[0] = "a"
            }
            _ = v35
        }
    default:
        var v36 uint64 = 662
        _ = v36
    }
    for fi37 := 0; fi37 < 13; fi37++ {
        for fi38 := 0; fi38 < 13; fi38++ {
            if true {
                v2 -= (^(v2))
                var v39 uint8 = 0
                v40 := GenS4{F0: 2, F1: ((v4 + v4) + "a")}
                _ = v39
                _ = v40
            } else {
                if ((-927 <= fi37) && ((v6 >> ((910) & 63)) == (^(v6)))) {
                    fmt.Println(v2, fi37)
                    fmt.Println(fi38)
                }
                switch fi38 {
                case -2147483648:
                    var v41 int32 = -307
                    _ = v41
                case 15:
                    var v42 bool = true
                    _ = v42
                case -1:
                    v10 -= -860
                default:
                    v4 = "a"
                }
                if len(v3) > 0 {
                    v3[0] = v4
                }
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = len(v5)
            _ = fi38
        }
        var v43 int8 = 102
        v2 += v2
        _ = fi37
        _ = v43
    }
    v2 = (^(v2))
    switch v4 {
    case "hello":
        switch v4 {
        case "a":
            for fi44 := 0; fi44 < 19; fi44++ {
                v45 := GenS3{F0: (v2 & (v2 / (((v2 + (v2 >> ((v6) & 63)))) | 1))), F1: v4, F2: v2}
                if true {
                    v2 += (^(v2))
                    var v46 string = v4
                    _ = v46
                } else {
                    v6 -= v6
                    fmt.Println(v6)
                    v47 := GenS5{F0: (v2 + v2), F1: 504}
                    _ = v47
                }
                _ = fi44
                _ = v45
            }
        case "hello":
            var v48 int = -1
            _ = v48
        }
    case "\n\t":
        fmt.Println(v10, v6, v10)
    case "":
        v6 = 42
    default:
        for fi49 := 0; fi49 < 2; fi49++ {
            if true {
                if true {
                    var v50 bool = true
                    v2 = (v2 >> (((v6 & 538)) & 63))
                    switch (("hello" + v4) + "") {
                    case "":
                        var v51 uint64 = v2
                        _ = v51
                    case "a":
                        v4 = (v4 + "foo bar")
                    case "foo bar":
                        v52 := GenS5{F0: v2, F1: 787}
                        _ = v52
                    default:
                        switch (^(fi49)) {
                        case 32768:
                            if v50 {
                                var v53 float64 = 0.1
                                fmt.Println(v53)
                                _ = v53
                            } else {
                                v2 += v2
                                for fi54 := 0; fi54 < 4; fi54++ {
                                    fmt.Println(fi49, v50, v10)
                                    var v55 float64 = 0.0
                                    if v50 {
                                        fmt.Println(v10)
                                        fmt.Println(v2, v2)
                                    }
                                    _ = fi54
                                    _ = v55
                                }
                                v4 = v4
                            }
                        case -9223372036854775808:
                            v6 = v6
                        }
                    }
                    _ = v50
                } else {
                    for fi56 := 0; fi56 < 6; fi56++ {
                        v4 = v4
                        if (!((false && (1.5 < (1.5 / 1.0))) && true)) {
                            v4 = (v4 + "\n\t")
                            v6 += 942
                            v57 := GenS6{F0: 1, F1: v10}
                            _ = v57
                        } else {
                            if ((!false) && true) {
                                var v58 string = v4
                                fmt.Println(v58)
                                _ = v58
                            }
                        }
                        _ = fi56
                    }
                    for fi59 := 0; fi59 < 4; fi59++ {
                        _ = len(v1)
                        var v60 int = fi49
                        _ = fi59
                        _ = v60
                    }
                }
                v10 = int64(v2)
                switch v4 {
                case "foo bar":
                    if (false || false) {
                        if (v2 == (v2 << (((v6 - 255)) & 63))) {
                            switch v4 {
                            case "":
                                for fi61 := 0; fi61 < 15; fi61++ {
                                    if false {
                                        v62 := GenS6{F0: 117, F1: (v10 + (v10 | v10))}
                                        delete(v1, v4)
                                        switch (("\n\t" + "hello") + ((v4 + "hello") + v4)) {
                                        case "a":
                                            var v63 float32 = (-2.25 + ((0.1 - 1e10) * ((3.14159 - 1.5) + 1.5)))
                                            _ = v63
                                        }
                                        if ((v10 ^ v10) >= 1) {
                                            if true {
                                                v10 -= v10
                                                switch v4 {
                                                case "":
                                                    if true {
                                                        var v64 float64 = (-2.25 / 1e10)
                                                        _ = v64
                                                    } else {
                                                        switch -159 {
                                                        case 127:
                                                            v65 := GenS7{F0: ((3.14159 >= ((3.14159 * ((((3.14159 + (3.14159 + (-2.25 / ((3.14159 - -2.25) / float64(0.1))))) / 1.0) - -2.25) + float64((0.1 * (float32(3.14159) + 0.1))))) + ((0.0 - 0.1) - 0.0))) && ((fi49 / (((fi61 >> ((v6) & 63))) | 1)) <= (fi49 << ((v6) & 63)))), F1: 42, F2: (!true)}
                                                            _ = v65
                                                        case 358:
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        case -9223372036854775808:
                                                            v2 = v2
                                                        default:
                                                            v3 = append(v3, "a")
                                                        }
                                                        _ = len(v1)
                                                        var v66 uint16 = uint16(v6)
                                                        _ = v66
                                                    }
                                                case "\n\t":
                                                    v2 = 1
                                                default:
                                                    v10 += v10
                                                }
                                                if (v4 > "foo bar") {
                                                    v67 := GenS4{F0: 42, F1: v4}
                                                    v68 := GenS8{F0: fi49, F1: false, F2: -1}
                                                    _ = v67
                                                    _ = v68
                                                } else {
                                                    v4 = ((v4 + "") + v4)
                                                    v10 += -562
                                                    v10 = (^(v10))
                                                    v10 = 1
                                                    v3 = append(v3, "\n\t")
                                                    var v69 int64 = v10
                                                    switch (v4 + ("" + v4)) {
                                                    case "a":
                                                        if (v4 == v4) {
                                                            switch ("\n\t" + "a") {
                                                            case "a":
                                                                var v70 uint16 = 1
                                                                _ = v70
                                                            case "hello":
                                                                if false {
                                                                    var v71 int8 = 100
                                                                    var v72 bool = false
                                                                    fmt.Println(fi61, v6, fi61)
                                                                    if (v72 || (629 == 65535)) {
                                                                        v2 -= (v2 & 0)
                                                                        _ = len(v1)
                                                                        v73 := GenS9{F0: 1}
                                                                        _ = v73
                                                                    }
                                                                    _ = v71
                                                                    _ = v72
                                                                }
                                                                v74 := GenS2{F0: uint16(v6), F1: (v6 % (((v6 | v6)) | 1)), F2: fi49}
                                                                if false {
                                                                    var v75 uint64 = v2
                                                                    switch fi61 {
                                                                    case 0:
                                                                        var v76 uint = 18446744073709551615
                                                                        _ = v76
                                                                    case -2147483648:
                                                                        var v77 uint64 = v2
                                                                        v2 = v75
                                                                        switch fi49 {
                                                                        case 0:
                                                                            switch v4 {
                                                                            case "":
                                                                                var v78 float32 = 1.5
                                                                                _ = v78
                                                                            case "hello":
                                                                                var v79 uint = v6
                                                                                _ = v79
                                                                            case "\n\t":
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = v4
                                                                                }
                                                                                v10 -= 9223372036854775807
                                                                            default:
                                                                                v4 = ("foo bar" + "\n\t")
                                                                            }
                                                                        case 32768:
                                                                            if true {
                                                                                if (true || ((^(v6)) <= 0)) {
                                                                                    v80 := GenS1{F0: (2 != 42)}
                                                                                    _ = v80
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v77, v69)
                                                                                var v81 float64 = 0.1
                                                                                fmt.Println(v4, v6, v81)
                                                                                _ = v81
                                                                            }
                                                                        case -251:
                                                                            var v82 uint = v6
                                                                            _ = v82
                                                                        default:
                                                                            var v83 int16 = -254
                                                                            _ = v83
                                                                        }
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (((v4 + v4) + (v4 + (v4 + (v4 + v4)))) + "")
                                                                        }
                                                                        v6 -= (v6 >> ((v6) & 63))
                                                                        var v84 uint8 = 186
                                                                        _ = v77
                                                                        _ = v84
                                                                    }
                                                                    v85 := GenS8{F0: 127, F1: (true || (0.1 > (((float64((((-2.25 + -2.25) * 1e10) / ((((0.0 / (0.0 + (0.0 - (1e10 / -2.25)))) * 0.0) + 1.5) / (((((1.5 - -2.25) + 3.14159) + (((1e10 - 3.14159) + 0.1) + -2.25)) / ((1e10 + 1e10) / 3.14159)) * 0.1)))) - (-2.25 / float64((1e10 / 1.0)))) / 1e10) * 3.14159))), F2: 9223372036854775807}
                                                                    switch v4 {
                                                                    case "\n\t":
                                                                        v86 := GenS10{F0: (v69 | v10), F1: (968 != 575)}
                                                                        _ = v86
                                                                    case "hello":
                                                                        if false {
                                                                            v3 = append(v3, v4)
                                                                            v87 := GenS9{F0: 42}
                                                                            if len(v3) > 0 {
                                                                                v3[0] = (((v4 + (v4 + (v4 + (("foo bar" + v4) + (("" + "a") + v4))))) + v4) + v4)
                                                                            }
                                                                            var v88 float64 = (-2.25 + 1.5)
                                                                            if ((!(!true)) || (228 != 127)) {
                                                                                var v89 int16 = -32768
                                                                                v90 := GenS4{F0: 2, F1: ((v4 + v4) + "hello")}
                                                                                _ = v89
                                                                                _ = v90
                                                                            } else {
                                                                                var v91 uint8 = 1
                                                                                fmt.Println(v75, v91)
                                                                                v2 = uint64(v6)
                                                                                _ = v91
                                                                            }
                                                                            fmt.Println(v6)
                                                                            delete(v5, fi49)
                                                                            switch (^(fi61)) {
                                                                            case 4294967295:
                                                                                v6 -= v6
                                                                            case 32768:
                                                                                v3 = append(v3, (v4 + v4))
                                                                            }
                                                                            fmt.Println(fi61, v6)
                                                                            _ = v87
                                                                            _ = v88
                                                                        } else {
                                                                            if (-169 >= 11) {
                                                                                var v92 int = (fi61 >> ((206) & 63))
                                                                                _ = v92
                                                                            }
                                                                            var v93 uint8 = 1
                                                                            _ = v93
                                                                        }
                                                                    }
                                                                    if (!false) {
                                                                        v94 := GenS11{F0: (fi61 + (fi49 % (((fi61 | (fi61 << ((v6) & 63)))) | 1))), F1: 0, F2: 42}
                                                                        v6 += v6
                                                                        var v95 uint8 = 0
                                                                        var v96 float32 = 1.5
                                                                        v97 := GenS12{F0: v2, F1: -1, F2: -114}
                                                                        v98 := GenS13{F0: v75, F1: 127}
                                                                        _ = v94
                                                                        _ = v95
                                                                        _ = v96
                                                                        _ = v97
                                                                        _ = v98
                                                                    }
                                                                    _ = v75
                                                                    _ = v85
                                                                } else {
                                                                    v99 := GenS9{F0: 32767}
                                                                    var v100 bool = true
                                                                    _ = v100
                                                                    _ = v99
                                                                }
                                                                _ = v74
                                                            }
                                                            v101 := GenS14{F0: 0, F1: -127, F2: fi61}
                                                            v2 -= v2
                                                            _ = v101
                                                        } else {
                                                            var v102 string = v4
                                                            v4 = v102
                                                            _ = v102
                                                        }
                                                    default:
                                                        var v103 float64 = -2.25
                                                        _ = v103
                                                    }
                                                    v10 = 9223372036854775807
                                                    v104 := GenS14{F0: v10, F1: 15, F2: (fi49 & (fi61 - (-(fi61))))}
                                                    _ = v104
                                                    _ = v69
                                                }
                                                var v105 bool = (1 != -1)
                                                if ((((v4 < ("foo bar" + ("a" + v4))) || false) && false) && v105) {
                                                    if v105 {
                                                        var v106 int8 = 0
                                                        fmt.Println(v10, v6, v106)
                                                        if len(v3) > 0 {
                                                            v3[0] = "foo bar"
                                                        }
                                                        v107 := GenS8{F0: 32767, F1: (!true), F2: v10}
                                                        _ = v106
                                                        _ = v107
                                                    } else {
                                                        switch 1 {
                                                        case 32767:
                                                            delete(v1, (v4 + (v4 + "hello")))
                                                        }
                                                        fmt.Println(v105, v6, v10)
                                                    }
                                                    var v108 int8 = 100
                                                    v105 = v105
                                                    _ = v108
                                                } else {
                                                    if v105 {
                                                        v109 := GenS15{F0: (v10 << (((v6 / (((v6 & v6)) | 1))) & 63))}
                                                        _ = len(v3)
                                                        var v110 string = "a"
                                                        _ = v109
                                                        _ = v110
                                                    } else {
                                                        switch (v4 + "hello") {
                                                        case "":
                                                            _ = len(v1)
                                                        }
                                                    }
                                                    v1[("a" + "foo bar")] = 1
                                                    var v111 uint8 = 242
                                                    fmt.Println(v111)
                                                    _ = v111
                                                }
                                                _ = v105
                                            }
                                        }
                                        v6 = (v6 << ((65535) & 63))
                                        _ = v62
                                    } else {
                                        if (1.5 == (1.5 / ((-2.25 * 1e10) / ((-2.25 / 1e10) - 1e10)))) {
                                            v10 = v10
                                            v3 = append(v3, ("" + v4))
                                            var v112 uint = 255
                                            _ = v112
                                        }
                                        fmt.Println(fi61, fi61)
                                        fmt.Println(fi49, v4, fi61)
                                    }
                                    fmt.Println(v10)
                                    _ = fi61
                                }
                            case "hello":
                                v3 = append(v3, "a")
                            case "a":
                                v6 -= (^(v6))
                            }
                            if false {
                                var v113 uint = (v6 / (((^(v6))) | 1))
                                _ = v113
                            }
                        }
                        v3 = append(v3, "\n\t")
                    } else {
                        for fi114 := 0; fi114 < 2; fi114++ {
                            v115 := GenS15{F0: (v10 / (((v10 << (((v6 ^ 2)) & 63))) | 1))}
                            _ = fi114
                            _ = v115
                        }
                        switch fi49 {
                        case 2147483647:
                            var v116 uint = 65535
                            _ = v116
                        case 4294967295:
                            v117 := GenS5{F0: (^(v2)), F1: 32767}
                            _ = v117
                        case 787:
                            v118 := GenS16{F0: 263, F1: -40}
                            _ = v118
                        default:
                            var v119 int8 = -123
                            _ = v119
                        }
                    }
                default:
                    if true {
                        var v120 uint = v6
                        if false {
                            v121 := GenS11{F0: 1, F1: 42, F2: 0}
                            _ = v121
                        }
                        var v122 uint8 = 255
                        _ = v120
                        _ = v122
                    } else {
                        if (!false) {
                            var v123 uint64 = (v2 ^ uint64(v6))
                            if false {
                                v1[(v4 + (v4 + ("hello" + "\n\t")))] = 2147483647
                                var v124 uint8 = 229
                                _ = v124
                            }
                            _ = v123
                        } else {
                            for fi125 := 0; fi125 < 8; fi125++ {
                                v4 = v4
                                _ = len(v1)
                                _ = fi125
                            }
                            v2 += 18446744073709551615
                            v126 := GenS17{F0: (-2147483648 != 603)}
                            _ = v126
                        }
                        v5[4294967295] = 65535
                    }
                }
            }
            var v127 uint32 = 936
            var v128 float64 = 0.0
            _ = fi49
            _ = v127
            _ = v128
        }
    }
    v4 = ((v4 + "a") + v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
