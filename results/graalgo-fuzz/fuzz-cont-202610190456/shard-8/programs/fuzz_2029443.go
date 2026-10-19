package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint32; F1 int; F2 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint; F1 int32 }
type GenS7 struct { F0 bool; F1 uint; F2 int32 }
type GenS8 struct { F0 string; F1 int16; F2 uint64 }
type GenS9 struct { F0 uint8; F1 uint; F2 bool }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint16; F1 int64 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 string; F1 uint64 }
type GenS14 struct { F0 int8; F1 bool }
type GenS15 struct { F0 uint8; F1 uint }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int16 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]uint32)
    v3 := []string{"hello"}
    var v4 float32 = 1.5
    v5 := make(map[string]int)
    var v6 string = ""
    switch "a" {
    case "foo bar":
        for fi7 := 0; fi7 < 1; fi7++ {
            fmt.Println(fi7, v4, v6)
            _ = fi7
        }
    default:
        _ = len(v5)
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 int64 = -377
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi8
        _ = v9
    }
    v4 = v4
    v6 = v6
    var v10 uint16 = 42
    if (9223372036854775807 > 32767) {
        fmt.Println(v6, v6, v6)
    } else {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        var v11 int = -910
        _ = v11
    }
    if (false && (42 >= 0)) {
        v6 = "hello"
    } else {
        var v12 int64 = 42
        _ = v12
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        if true {
            var v14 int16 = 0
            var v15 int16 = (v14 / ((42) | 1))
            _ = v14
            _ = v15
        }
        fmt.Println(v4, v10, v6)
        delete(v1, fi13)
        _ = fi13
    }
    if len(v3) > 0 {
        v3[0] = v6
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 int64 = -1
            if ((v4 - v4) > v4) {
                fmt.Println(v18, fi17, v18)
                fmt.Println(v6, v18, fi17)
            } else {
                fmt.Println(fi17, v10)
            }
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    if (false && false) {
        v4 -= v4
        v19 := GenS1{F0: 42, F1: -1}
        v20 := GenS2{F0: 2147483647, F1: 0, F2: -2147483648}
        _ = v19
        _ = v20
    } else {
        switch -1 {
        case 42:
            var v21 uint16 = (v10 << ((1) & 63))
            _ = v21
        case 963:
            if ((!true) && (1 < 42)) {
                switch 1 {
                case -2147483648:
                    v22 := v1[32767]
                    _ = v22
                case 559:
                    if (true && true) {
                        var v23 int16 = 1
                        v6 = ("hello" + (("a" + v6) + "foo bar"))
                        _ = v23
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                case -9223372036854775808:
                    v6 = v6
                default:
                    for fi24 := 0; fi24 < 3; fi24++ {
                        if false {
                            var v25 int32 = 0
                            _ = v25
                        } else {
                            var v26 float64 = -2.25
                            if (0.0 != v26) {
                                if ((float64(0.0) - v26) > ((v26 * 1e10) + 0.1)) {
                                    v6 = ("\n\t" + v6)
                                    var v27 uint32 = 42
                                    v27 -= v27
                                    _ = v27
                                } else {
                                    for fi28 := 0; fi28 < 14; fi28++ {
                                        var v29 uint32 = 0
                                        _ = fi28
                                        _ = v29
                                    }
                                    switch (fi24 << ((0) & 63)) {
                                    case -729:
                                        v30 := GenS2{F0: 65535, F1: fi24, F2: 2147483647}
                                        _ = v30
                                    case 127:
                                        if true {
                                            var v31 uint32 = 116
                                            switch (-(fi24)) {
                                            case -1:
                                                switch int(v31) {
                                                case 32767:
                                                    var v32 int32 = 42
                                                    _ = v32
                                                case 42:
                                                    v4 -= -2.25
                                                default:
                                                    v33 := GenS2{F0: (v31 * v31), F1: (^(fi24)), F2: 2147483647}
                                                    _ = v33
                                                }
                                            case 1:
                                                switch fi24 {
                                                case -539:
                                                    for fi34 := 0; fi34 < 5; fi34++ {
                                                        v35 := GenS1{F0: 2, F1: 1}
                                                        _ = fi34
                                                        _ = v35
                                                    }
                                                default:
                                                    switch fi24 {
                                                    case 9223372036854775807:
                                                        switch 646 {
                                                        case -584:
                                                            if (v31 != 991) {
                                                                var v36 float64 = v26
                                                                switch (v6 + "\n\t") {
                                                                case "hello":
                                                                    if true {
                                                                        switch (fi24 % (((fi24 & fi24)) | 1)) {
                                                                        case 431:
                                                                            if (true && true) {
                                                                                v37 := GenS1{F0: 1, F1: 0}
                                                                                _ = v37
                                                                            } else {
                                                                                var v38 float32 = ((float32(1.5) / ((0.1 - ((v4 + v4) + ((v4 - (0.0 / float32(v36))) * 3.14159))) / float32(v36))) / v4)
                                                                                _ = v38
                                                                            }
                                                                        default:
                                                                            for fi39 := 0; fi39 < 17; fi39++ {
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                var v40 float32 = (float32(v36) - v4)
                                                                                _ = fi39
                                                                                _ = v40
                                                                            }
                                                                        }
                                                                        for fi41 := 0; fi41 < 18; fi41++ {
                                                                            if false {
                                                                                v36 += v26
                                                                                if (2 < 2) {
                                                                                    switch v6 {
                                                                                    case "foo bar":
                                                                                        fmt.Println(v10)
                                                                                    case "\n\t":
                                                                                        v4 += float32((v26 - 0.1))
                                                                                    case "a":
                                                                                        v10 -= v10
                                                                                    default:
                                                                                        v42 := GenS3{F0: 32767}
                                                                                        _ = v42
                                                                                    }
                                                                                    fmt.Println(fi41, v36)
                                                                                }
                                                                                var v43 uint32 = 2147483647
                                                                                v26 -= v26
                                                                                _ = v43
                                                                            } else {
                                                                                if true {
                                                                                    var v44 string = ((("foo bar" + "foo bar") + "") + "a")
                                                                                    delete(v1, int(v31))
                                                                                    _ = v44
                                                                                } else {
                                                                                    v4 -= v4
                                                                                    var v45 uint16 = v10
                                                                                    _ = v45
                                                                                }
                                                                                var v46 float32 = -2.25
                                                                                _ = len(v3)
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    if (529 < 32767) {
                                                                                        v2[-1] = uint32(v10)
                                                                                        if true {
                                                                                            var v47 uint8 = 127
                                                                                            switch (fi24 + 9223372036854775807) {
                                                                                            case 0:
                                                                                                var v48 int16 = 164
                                                                                                _ = v48
                                                                                            case 9223372036854775807:
                                                                                                if ((^(v10)) < v10) {
                                                                                                    var v49 uint8 = uint8(v31)
                                                                                                    _ = v49
                                                                                                } else {
                                                                                                    if true {
                                                                                                        v46 += 1.5
                                                                                                        var v50 uint16 = v10
                                                                                                        var v51 uint32 = (v31 << ((1) & 63))
                                                                                                        _ = v50
                                                                                                        _ = v51
                                                                                                    } else {
                                                                                                        var v52 int8 = -128
                                                                                                        v53 := GenS2{F0: v31, F1: 4294967295, F2: 138}
                                                                                                        v54 := GenS2{F0: (v31 % (((^(v31))) | 1)), F1: fi41, F2: -1}
                                                                                                        _ = v52
                                                                                                        _ = v53
                                                                                                        _ = v54
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                v6 = "\n\t"
                                                                                            }
                                                                                            fmt.Println(fi24, v4, v31)
                                                                                            _ = v47
                                                                                        } else {
                                                                                            var v55 string = "a"
                                                                                            _ = v55
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v46
                                                                            }
                                                                            if false {
                                                                                v10 = (v10 | v10)
                                                                                v6 = (v6 + v6)
                                                                            }
                                                                            var v56 uint8 = 182
                                                                            _ = fi41
                                                                            _ = v56
                                                                        }
                                                                        var v57 int64 = -1
                                                                        _ = v57
                                                                    }
                                                                }
                                                                _ = v36
                                                            } else {
                                                                v58 := GenS2{F0: v31, F1: fi24, F2: 42}
                                                                for fi59 := 0; fi59 < 2; fi59++ {
                                                                    fmt.Println(fi59, fi24, v6)
                                                                    if false {
                                                                        if true {
                                                                            var v60 string = "foo bar"
                                                                            if (!(875 != 425)) {
                                                                                var v61 int8 = -1
                                                                                _ = v61
                                                                            }
                                                                            _ = v60
                                                                        } else {
                                                                            if (100 < -128) {
                                                                                v62 := GenS4{F0: -1}
                                                                                v6 = v6
                                                                                if (v10 != 65535) {
                                                                                    _ = len(v2)
                                                                                    var v63 float64 = v26
                                                                                    _ = v63
                                                                                } else {
                                                                                    if (true || true) {
                                                                                        if (((v26 - ((float64(3.14159) - (-2.25 - 0.0)) * (3.14159 * v26))) - 3.14159) != v26) {
                                                                                            fmt.Println(v26, v31, fi59)
                                                                                            v4 = v4
                                                                                            switch "\n\t" {
                                                                                            case "foo bar":
                                                                                                fmt.Println(v31, fi24, v26)
                                                                                            case "hello":
                                                                                                v6 = ((v6 + ("\n\t" + v6)) + ("a" + v6))
                                                                                            }
                                                                                        }
                                                                                        v64 := v5[v6]
                                                                                        var v65 uint64 = 1
                                                                                        _ = v64
                                                                                        _ = v65
                                                                                    }
                                                                                    var v66 int8 = 46
                                                                                    v67 := GenS1{F0: 457, F1: (v66 ^ v66)}
                                                                                    _ = v66
                                                                                    _ = v67
                                                                                }
                                                                                switch v6 {
                                                                                case "":
                                                                                    v68 := v5[""]
                                                                                    fmt.Println(v4)
                                                                                    v69 := GenS5{F0: 100}
                                                                                    v4 = v4
                                                                                    _ = v68
                                                                                    _ = v69
                                                                                }
                                                                                _ = v62
                                                                            }
                                                                            if ((true && true) && (803 == 65535)) {
                                                                                switch (v6 + v6) {
                                                                                case "a":
                                                                                    var v70 string = v6
                                                                                    _ = v70
                                                                                default:
                                                                                    switch v6 {
                                                                                    case "":
                                                                                        switch "foo bar" {
                                                                                        case "a":
                                                                                            var v71 int32 = -809
                                                                                            _ = v71
                                                                                        default:
                                                                                            v3 = append(v3, "hello")
                                                                                        }
                                                                                    case "hello":
                                                                                        delete(v2, fi24)
                                                                                    default:
                                                                                        fmt.Println(fi59, fi24, v4)
                                                                                    }
                                                                                }
                                                                                if false {
                                                                                    if true {
                                                                                        switch -2147483648 {
                                                                                        case 357:
                                                                                            _ = len(v5)
                                                                                        case 239:
                                                                                            var v72 uint8 = 115
                                                                                            _ = v72
                                                                                        case 32767:
                                                                                            v26 += 3.14159
                                                                                        default:
                                                                                            var v73 float32 = v4
                                                                                            _ = v73
                                                                                        }
                                                                                        v74 := v5["hello"]
                                                                                        if false {
                                                                                            v74 += 134
                                                                                            fmt.Println(v74)
                                                                                            var v75 uint16 = (v10 + v10)
                                                                                            if false {
                                                                                                v76 := GenS2{F0: (v31 << ((406) & 63)), F1: (fi24 / ((128) | 1)), F2: 0}
                                                                                                if false {
                                                                                                    if (!true) {
                                                                                                        fmt.Println(v31, v26, fi59)
                                                                                                        var v77 bool = true
                                                                                                        _ = v77
                                                                                                    } else {
                                                                                                        var v78 float32 = float32(v26)
                                                                                                        _ = v78
                                                                                                    }
                                                                                                    var v79 string = (v6 + (v6 + (("\n\t" + (v6 + ((v6 + v6) + "\n\t"))) + ("foo bar" + "hello"))))
                                                                                                    _ = v79
                                                                                                }
                                                                                                _ = v76
                                                                                            }
                                                                                            _ = v75
                                                                                        }
                                                                                        v80 := GenS5{F0: 19}
                                                                                        fmt.Println(v10, v74)
                                                                                        _ = v74
                                                                                        _ = v80
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                v81 := GenS4{F0: -9223372036854775808}
                                                                                v10 += v10
                                                                                _ = v81
                                                                            }
                                                                            _ = len(v3)
                                                                            v82 := GenS6{F0: uint(v10), F1: -9}
                                                                            if ((128 != 181) && true) {
                                                                                v10 -= (v10 & 255)
                                                                                v26 -= 1.5
                                                                                var v83 uint = uint(v10)
                                                                                _ = v83
                                                                            }
                                                                            fmt.Println(v26)
                                                                            v31 -= v31
                                                                            _ = v82
                                                                        }
                                                                        v3 = append(v3, (v6 + v6))
                                                                    } else {
                                                                        var v84 int64 = 42
                                                                        var v85 string = ((v6 + "foo bar") + "")
                                                                        fmt.Println(v10, v26, v85)
                                                                        var v86 bool = true
                                                                        _ = v84
                                                                        _ = v85
                                                                        _ = v86
                                                                    }
                                                                    fmt.Println(fi59, fi59, v26)
                                                                    fmt.Println(fi24, fi59)
                                                                    v26 = (((float64(0.1) - (v26 + v26)) + v26) / v26)
                                                                    _ = fi59
                                                                }
                                                                if true {
                                                                    switch fi24 {
                                                                    case 9223372036854775807:
                                                                        if true {
                                                                            v87 := GenS7{F0: true, F1: 413, F2: 2147483647}
                                                                            var v88 int8 = -1
                                                                            _ = v87
                                                                            _ = v88
                                                                        } else {
                                                                            var v89 int = fi24
                                                                            switch v6 {
                                                                            case "hello":
                                                                                if (-79 <= 16) {
                                                                                    if (((410 <= 296) || (-2147483648 >= 2147483647)) || (42 <= 42)) {
                                                                                        v10 = 255
                                                                                        var v90 uint = 1
                                                                                        v91 := GenS1{F0: uint(v31), F1: 0}
                                                                                        _ = v90
                                                                                        _ = v91
                                                                                    }
                                                                                    for fi92 := 0; fi92 < 15; fi92++ {
                                                                                        fmt.Println(fi92, v10)
                                                                                        if (true && true) {
                                                                                            var v93 int = v89
                                                                                            v26 = float64(1.5)
                                                                                            _ = v93
                                                                                        }
                                                                                        if ((32767 == v10) || true) {
                                                                                            if (-9223372036854775808 == 596) {
                                                                                                switch "\n\t" {
                                                                                                case "":
                                                                                                    v3 = append(v3, v6)
                                                                                                case "\n\t":
                                                                                                    var v94 int16 = -529
                                                                                                    _ = v94
                                                                                                }
                                                                                                var v95 int64 = -205
                                                                                                var v96 bool = false
                                                                                                _ = v95
                                                                                                _ = v96
                                                                                            }
                                                                                            v3 = append(v3, ("foo bar" + (v6 + (v6 + v6))))
                                                                                            v6 = ("\n\t" + v6)
                                                                                        } else {
                                                                                            var v97 int16 = 0
                                                                                            v4 += (v4 - v4)
                                                                                            _ = v97
                                                                                        }
                                                                                        _ = fi92
                                                                                    }
                                                                                } else {
                                                                                    v98 := GenS2{F0: v31, F1: (v89 - fi24), F2: 1}
                                                                                    v89 -= (-(v89))
                                                                                    _ = v98
                                                                                }
                                                                            case "\n\t":
                                                                                v89 = v89
                                                                            case "a":
                                                                                var v99 uint32 = v31
                                                                                _ = v99
                                                                            }
                                                                            _ = v89
                                                                        }
                                                                    case 738:
                                                                        if (true || (!(!false))) {
                                                                            if true {
                                                                                var v100 uint8 = 2
                                                                                if (3.14159 >= 0.1) {
                                                                                    var v101 bool = ((fi24 != fi24) || true)
                                                                                    v6 = v6
                                                                                    v31 = (v31 ^ v31)
                                                                                    _ = v101
                                                                                }
                                                                                _ = v100
                                                                            } else {
                                                                                v5[v6] = -946
                                                                            }
                                                                            for fi102 := 0; fi102 < 11; fi102++ {
                                                                                v4 += v4
                                                                                _ = fi102
                                                                            }
                                                                        }
                                                                    case -1:
                                                                        if true {
                                                                            if true {
                                                                                if true {
                                                                                    v103 := v2[fi24]
                                                                                    var v104 uint16 = (v10 | v10)
                                                                                    _ = v103
                                                                                    _ = v104
                                                                                } else {
                                                                                    for fi105 := 0; fi105 < 11; fi105++ {
                                                                                        v106 := GenS4{F0: 0}
                                                                                        _ = fi105
                                                                                        _ = v106
                                                                                    }
                                                                                    if (fi24 != (^(fi24))) {
                                                                                        var v107 uint32 = v31
                                                                                        var v108 int64 = -9223372036854775808
                                                                                        _ = v107
                                                                                        _ = v108
                                                                                    }
                                                                                    switch fi24 {
                                                                                    case 32767:
                                                                                        v109 := GenS5{F0: -128}
                                                                                        _ = v109
                                                                                    case 42:
                                                                                        delete(v1, -469)
                                                                                    case 9223372036854775807:
                                                                                        for fi110 := 0; fi110 < 12; fi110++ {
                                                                                            var v111 uint32 = (v31 ^ v31)
                                                                                            if (505 >= 2147483647) {
                                                                                                if (!true) {
                                                                                                    var v112 int16 = 1000
                                                                                                    v3 = append(v3, ("foo bar" + v6))
                                                                                                    _ = v112
                                                                                                }
                                                                                            } else {
                                                                                                v4 = 3.14159
                                                                                                fmt.Println(v4, v6, v6)
                                                                                                var v113 uint = 560
                                                                                                _ = v113
                                                                                            }
                                                                                            if (1 >= 127) {
                                                                                                if true {
                                                                                                    v111 = v111
                                                                                                }
                                                                                                v114 := GenS1{F0: 1, F1: -56}
                                                                                                _ = v114
                                                                                            }
                                                                                            _ = fi110
                                                                                            _ = v111
                                                                                        }
                                                                                    default:
                                                                                        fmt.Println(v31, v6, v31)
                                                                                    }
                                                                                }
                                                                                fmt.Println(v31, v6, v10)
                                                                            }
                                                                            v10 -= (v10 >> ((0) & 63))
                                                                        }
                                                                    }
                                                                    v4 = v4
                                                                    var v115 bool = (true && false)
                                                                    _ = v115
                                                                }
                                                                _ = v58
                                                            }
                                                        case 0:
                                                            var v116 uint64 = 409
                                                            _ = v116
                                                        case -444:
                                                            var v117 int32 = -2147483648
                                                            _ = v117
                                                        default:
                                                            v118 := GenS2{F0: 2147483647, F1: (fi24 - -593), F2: -603}
                                                            _ = v118
                                                        }
                                                    case 42:
                                                        for fi119 := 0; fi119 < 9; fi119++ {
                                                            v31 -= (v31 * 2)
                                                            if false {
                                                                v5[(v6 + v6)] = fi119
                                                                var v120 int8 = 0
                                                                _ = v120
                                                            }
                                                            v10 -= 592
                                                            if true {
                                                                v121 := v1[(-(fi119))]
                                                                switch v6 {
                                                                case "foo bar":
                                                                    if false {
                                                                        switch v6 {
                                                                        case "\n\t":
                                                                            var v122 bool = (21 > 42)
                                                                            _ = v122
                                                                        case "hello":
                                                                            if (40 >= 154) {
                                                                                v123 := GenS8{F0: ("a" + ("" + v6)), F1: -32768, F2: 42}
                                                                                fmt.Println(fi24)
                                                                                var v124 int8 = -1
                                                                                _ = v123
                                                                                _ = v124
                                                                            }
                                                                            var v125 int16 = -32768
                                                                            if v121 {
                                                                                v121 = (v10 >= 2)
                                                                            } else {
                                                                                var v126 uint32 = (v31 & v31)
                                                                                v127 := GenS8{F0: (((v6 + (("hello" + ("" + v6)) + "hello")) + "") + v6), F1: 75, F2: 4294967295}
                                                                                _ = v126
                                                                                _ = v127
                                                                            }
                                                                            _ = v125
                                                                        default:
                                                                            v6 = v6
                                                                        }
                                                                        switch "foo bar" {
                                                                        case "a":
                                                                            if (fi119 != fi24) {
                                                                                if v121 {
                                                                                    var v128 int64 = 389
                                                                                    _ = v128
                                                                                }
                                                                                var v129 float64 = (1.5 / (1e10 * v26))
                                                                                _ = v129
                                                                            } else {
                                                                                v6 = v6
                                                                                fmt.Println(v4)
                                                                            }
                                                                        case "\n\t":
                                                                            var v130 uint8 = 2
                                                                            _ = v130
                                                                        case "":
                                                                            v10 = v10
                                                                        default:
                                                                            v131 := GenS4{F0: 996}
                                                                            _ = v131
                                                                        }
                                                                        var v132 int = int(v31)
                                                                        _ = v132
                                                                    } else {
                                                                        v133 := GenS7{F0: v121, F1: 42, F2: -2147483648}
                                                                        var v134 uint32 = (v31 ^ 2147483648)
                                                                        _ = v133
                                                                        _ = v134
                                                                    }
                                                                default:
                                                                    var v135 uint8 = 1
                                                                    _ = v135
                                                                }
                                                                _ = v121
                                                            }
                                                            v136 := GenS3{F0: 1000}
                                                            if (!(1.5 <= float64(v4))) {
                                                                if true {
                                                                    v137 := GenS4{F0: 9223372036854775807}
                                                                    if (v10 > (v10 << ((806) & 63))) {
                                                                        if (688 == v31) {
                                                                            v6 = v6
                                                                            var v138 float64 = ((v26 - v26) * v26)
                                                                            _ = v138
                                                                        }
                                                                    }
                                                                    var v139 string = (v6 + ("\n\t" + v6))
                                                                    _ = v137
                                                                    _ = v139
                                                                }
                                                                var v140 int32 = -2147483648
                                                                fmt.Println(v31, v26, fi24)
                                                                v10 = (^(v10))
                                                                _ = v140
                                                            }
                                                            _ = fi119
                                                            _ = v136
                                                        }
                                                    case -1:
                                                        var v141 int32 = -2147483648
                                                        _ = v141
                                                    default:
                                                        switch "a" {
                                                        case "foo bar":
                                                            v3 = append(v3, (v6 + (v6 + "\n\t")))
                                                        case "":
                                                            if ((fi24 ^ (fi24 & fi24)) <= (fi24 % ((2147483647) | 1))) {
                                                                v26 += -2.25
                                                            } else {
                                                                v142 := GenS9{F0: 120, F1: 42, F2: (false && (v31 != (^(v31))))}
                                                                if false {
                                                                    var v143 string = (v6 + v6)
                                                                    _ = v143
                                                                }
                                                                var v144 bool = (false && (!(false || (!(!((-160 < -650) && (!true)))))))
                                                                _ = v142
                                                                _ = v144
                                                            }
                                                        default:
                                                            if (2147483648 == (v31 ^ 895)) {
                                                                v10 = v10
                                                                var v145 uint64 = 4294967295
                                                                v146 := GenS6{F0: 0, F1: -1}
                                                                _ = v145
                                                                _ = v146
                                                            }
                                                        }
                                                    }
                                                }
                                            default:
                                                var v147 string = v6
                                                _ = v147
                                            }
                                            if (0 < v31) {
                                                v148 := GenS2{F0: v31, F1: (fi24 | int(v10)), F2: -1}
                                                _ = v148
                                            }
                                            _ = v31
                                        } else {
                                            if false {
                                                var v149 bool = false
                                                _ = v149
                                            } else {
                                                v10 = v10
                                                var v150 uint16 = (v10 >> ((42) & 63))
                                                var v151 int64 = 42
                                                _ = v150
                                                _ = v151
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = v6
                                            }
                                        }
                                    default:
                                        var v152 bool = true
                                        _ = v152
                                    }
                                    fmt.Println(v4, v4, v26)
                                }
                                switch v6 {
                                case "foo bar":
                                    var v153 int = fi24
                                    _ = v153
                                case "\n\t":
                                    delete(v2, fi24)
                                default:
                                    if true {
                                        v10 = 1
                                    }
                                }
                                var v154 uint8 = 255
                                _ = v154
                            }
                            _ = v26
                        }
                        _ = fi24
                    }
                }
                v155 := GenS10{F0: 4294967295}
                _ = v155
            }
        case -67:
            if false {
                var v156 int8 = 58
                v157 := GenS4{F0: -9223372036854775808}
                _ = v156
                _ = v157
            } else {
                switch -249 {
                case -254:
                    var v158 uint64 = 0
                    _ = v158
                default:
                    for fi159 := 0; fi159 < 11; fi159++ {
                        fmt.Println(v4, fi159, v4)
                        switch fi159 {
                        case -9223372036854775808:
                            switch (v6 + "\n\t") {
                            case "hello":
                                if true {
                                    fmt.Println(v10)
                                    if true {
                                        fmt.Println(fi159, fi159)
                                        var v160 int64 = -1
                                        if true {
                                            v4 -= 0.0
                                            v10 = v10
                                            var v161 bool = (18446744073709551615 >= 2)
                                            _ = v161
                                        }
                                        _ = v160
                                    } else {
                                        for fi162 := 0; fi162 < 10; fi162++ {
                                            fmt.Println(v6, fi162, v10)
                                            if (42 <= 1) {
                                                var v163 uint = 2
                                                var v164 int64 = int64(fi162)
                                                switch v6 {
                                                case "foo bar":
                                                    v164 += (v164 - v164)
                                                case "hello":
                                                    var v165 uint64 = 2
                                                    _ = v165
                                                case "a":
                                                    if (0 <= 42) {
                                                        v166 := v2[-786]
                                                        if (!false) {
                                                            switch "" {
                                                            case "foo bar":
                                                                v4 = v4
                                                            case "":
                                                                switch -719 {
                                                                case -387:
                                                                    fmt.Println(v6, v4)
                                                                case 4294967295:
                                                                    if (!(3.14159 <= (v4 * 1.5))) {
                                                                        switch (fi162 << ((v163) & 63)) {
                                                                        case 5:
                                                                            var v167 int = fi162
                                                                            v167 -= 99
                                                                            _ = v167
                                                                        }
                                                                        fmt.Println(v6)
                                                                        v6 = v6
                                                                    }
                                                                case 0:
                                                                    v164 -= (v164 & (v164 % ((v164) | 1)))
                                                                default:
                                                                    var v168 float32 = (v4 / v4)
                                                                    _ = v168
                                                                }
                                                            }
                                                            if false {
                                                                fmt.Println(v10)
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            } else {
                                                                v169 := GenS1{F0: (^(v163)), F1: 100}
                                                                _ = v169
                                                            }
                                                            var v170 int16 = -1
                                                            v171 := GenS7{F0: true, F1: (v163 * (v163 ^ (v163 << (((^(v163))) & 63)))), F2: -2147483648}
                                                            _ = v170
                                                            _ = v171
                                                        } else {
                                                            v164 += v164
                                                            var v172 bool = (v10 <= (v10 % ((1) | 1)))
                                                            _ = v172
                                                        }
                                                        v173 := GenS7{F0: false, F1: 467, F2: -1}
                                                        if false {
                                                            v163 = v163
                                                        } else {
                                                            v174 := GenS3{F0: 1}
                                                            v175 := GenS1{F0: (^(v163)), F1: int8(v164)}
                                                            fmt.Println(fi162, v163, v6)
                                                            fmt.Println(v166)
                                                            if (18446744073709551615 < 0) {
                                                                v10 = v10
                                                                fmt.Println(v6, v164)
                                                            }
                                                            _ = v174
                                                            _ = v175
                                                        }
                                                        _ = v166
                                                        _ = v173
                                                    } else {
                                                        v176 := GenS9{F0: 0, F1: (^(v163)), F2: (279 <= 60)}
                                                        if false {
                                                            _ = len(v1)
                                                            var v177 uint16 = 1
                                                            v10 -= v177
                                                            _ = v177
                                                        } else {
                                                            v178 := GenS3{F0: int16(v163)}
                                                            if ((1 == 127) || (0 >= 1)) {
                                                                var v179 uint64 = 83
                                                                v3 = append(v3, "hello")
                                                                switch fi159 {
                                                                case 4294967295:
                                                                    v10 = 325
                                                                }
                                                                _ = v179
                                                            }
                                                            _ = v178
                                                        }
                                                        _ = v176
                                                    }
                                                default:
                                                    v3 = append(v3, "hello")
                                                }
                                                _ = v163
                                                _ = v164
                                            }
                                            switch (fi159 % (((fi162 + fi162)) | 1)) {
                                            case -9223372036854775808:
                                                var v180 int16 = 1000
                                                _ = v180
                                            case 604:
                                                if len(v3) > 0 {
                                                    v3[0] = (v6 + ("foo bar" + "\n\t"))
                                                }
                                            case 0:
                                                v181 := GenS1{F0: 18446744073709551615, F1: 42}
                                                _ = v181
                                            default:
                                                switch (-(fi159)) {
                                                case -180:
                                                    v182 := GenS11{F0: v10, F1: 42}
                                                    _ = v182
                                                case 32767:
                                                    v183 := GenS12{F0: 42}
                                                    _ = v183
                                                case 171:
                                                    v184 := GenS1{F0: 0, F1: 42}
                                                    _ = v184
                                                }
                                            }
                                            _ = fi162
                                        }
                                        var v185 uint64 = 142
                                        v10 -= (v10 * v10)
                                        _ = v185
                                    }
                                    v186 := GenS6{F0: 775, F1: 0}
                                    _ = v186
                                } else {
                                    for fi187 := 0; fi187 < 1; fi187++ {
                                        v10 = (v10 % ((v10) | 1))
                                        if false {
                                            delete(v2, fi187)
                                            switch v6 {
                                            case "hello":
                                                v188 := GenS13{F0: ("foo bar" + "\n\t"), F1: 42}
                                                _ = v188
                                            case "":
                                                v6 = "\n\t"
                                            case "\n\t":
                                                v189 := GenS4{F0: 0}
                                                _ = v189
                                            default:
                                                switch (^(fi159)) {
                                                case 32768:
                                                    switch v6 {
                                                    case "\n\t":
                                                        if (32767 == fi187) {
                                                            switch fi159 {
                                                            case 42:
                                                                var v190 string = "\n\t"
                                                                _ = v190
                                                            default:
                                                                v191 := GenS14{F0: -1, F1: false}
                                                                _ = v191
                                                            }
                                                            var v192 uint8 = 42
                                                            if true {
                                                                fmt.Println(v4, v10)
                                                                v193 := GenS10{F0: 100}
                                                                fmt.Println(fi187, v192)
                                                                _ = v193
                                                            } else {
                                                                if true {
                                                                    var v194 bool = true
                                                                    v194 = (579 > 42)
                                                                    v195 := GenS15{F0: 242, F1: 18446744073709551615}
                                                                    _ = v194
                                                                    _ = v195
                                                                }
                                                                if false {
                                                                    v196 := GenS16{F0: true}
                                                                    _ = v196
                                                                } else {
                                                                    switch "a" {
                                                                    case "hello":
                                                                        if false {
                                                                            v197 := GenS9{F0: 238, F1: 755, F2: (42 < 1)}
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            v3 = append(v3, v6)
                                                                            _ = v197
                                                                        } else {
                                                                            if false {
                                                                                fmt.Println(fi159, v10)
                                                                            }
                                                                            var v198 uint32 = 1
                                                                            _ = v198
                                                                        }
                                                                    case "a":
                                                                        var v199 int8 = 1
                                                                        _ = v199
                                                                    case "":
                                                                        v200 := v2[(fi187 * (fi159 % ((fi159) | 1)))]
                                                                        _ = v200
                                                                    }
                                                                    switch v6 {
                                                                    case "":
                                                                        var v201 int16 = 0
                                                                        _ = v201
                                                                    default:
                                                                        v202 := GenS15{F0: 42, F1: 18446744073709551615}
                                                                        _ = len(v5)
                                                                        v203 := GenS5{F0: 1}
                                                                        _ = v202
                                                                        _ = v203
                                                                    }
                                                                }
                                                                v192 = 128
                                                                fmt.Println(v6, fi187, v192)
                                                            }
                                                            _ = v192
                                                        }
                                                    default:
                                                        var v204 string = v6
                                                        _ = v204
                                                    }
                                                case 91:
                                                    var v205 uint = 2
                                                    _ = v205
                                                case 1:
                                                    var v206 int32 = 42
                                                    _ = v206
                                                default:
                                                    var v207 uint8 = 255
                                                    _ = v207
                                                }
                                            }
                                        } else {
                                            switch ("" + ("\n\t" + v6)) {
                                            case "foo bar":
                                                var v208 float32 = (3.14159 + v4)
                                                _ = v208
                                            case "a":
                                                v5[((("" + v6) + "foo bar") + v6)] = 0
                                            case "":
                                                v6 = v6
                                            }
                                            switch (fi187 / ((fi159) | 1)) {
                                            case 1:
                                                var v209 int32 = 2147483647
                                                _ = v209
                                            case 13:
                                                v6 = v6
                                            case 128:
                                                var v210 bool = true
                                                v4 = 1e10
                                                if true {
                                                    if ("" < (v6 + "\n\t")) {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    } else {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                    if v210 {
                                                        fmt.Println(fi159, fi159, v6)
                                                        fmt.Println(fi187)
                                                    }
                                                    v4 -= v4
                                                } else {
                                                    var v211 uint32 = 869
                                                    var v212 uint8 = 255
                                                    if v210 {
                                                        var v213 uint32 = v211
                                                        _ = v213
                                                    } else {
                                                        v10 += uint16(v211)
                                                        var v214 float64 = 3.14159
                                                        var v215 uint32 = (^(v211))
                                                        _ = v214
                                                        _ = v215
                                                    }
                                                    _ = v211
                                                    _ = v212
                                                }
                                                var v216 int32 = 690
                                                v217 := GenS17{F0: -32768}
                                                _ = v210
                                                _ = v216
                                                _ = v217
                                            default:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                            v10 = (v10 / (((v10 | (v10 >> ((42) & 63)))) | 1))
                                        }
                                        _ = fi187
                                    }
                                }
                            case "\n\t":
                                var v218 bool = (0 <= 42)
                                _ = v218
                            default:
                                var v219 bool = (65535 > 2)
                                _ = v219
                            }
                        }
                        for fi220 := 0; fi220 < 2; fi220++ {
                            if ((4294967295 == 18446744073709551615) || false) {
                                v6 = (v6 + v6)
                                v4 = 3.14159
                            } else {
                                if ((115 <= v10) && (!false)) {
                                    if ((v10 * 987) != (v10 & 443)) {
                                        v10 = v10
                                        var v221 uint16 = (v10 - (v10 / ((v10) | 1)))
                                        fmt.Println(fi220)
                                        fmt.Println(fi220, v6, fi159)
                                        v222 := v1[fi220]
                                        v221 = v221
                                        switch 42 {
                                        case 2147483647:
                                            v223 := GenS7{F0: (!(115 > 2147483647)), F1: 65535, F2: -889}
                                            _ = v223
                                        default:
                                            fmt.Println(v6, v10)
                                        }
                                        _ = v221
                                        _ = v222
                                    } else {
                                        var v224 float64 = -2.25
                                        _ = v224
                                    }
                                    v10 -= (v10 << ((150) & 63))
                                    fmt.Println(fi159)
                                } else {
                                    var v225 int32 = -508
                                    v225 = 985
                                    _ = v225
                                }
                            }
                            var v226 uint16 = 2
                            _ = fi220
                            _ = v226
                        }
                        _ = fi159
                    }
                }
            }
        }
    }
    v227 := GenS16{F0: (false || true)}
    v4 += 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v227
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
