package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16; F1 uint; F2 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint8; F1 int }
type GenS8 struct { F0 uint64; F1 int }
type GenS9 struct { F0 uint8; F1 uint16 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint8; F1 string }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 int; F1 uint32; F2 uint32 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint32; F1 uint }
type GenS16 struct { F0 bool; F1 bool; F2 uint16 }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 uint32; F1 uint16; F2 int }
type GenS19 struct { F0 int8; F1 string; F2 uint8 }
type GenS20 struct { F0 int32 }
type GenS21 struct { F0 string; F1 int }
type GenS22 struct { F0 uint32; F1 int }
type GenS23 struct { F0 uint32; F1 string }
type GenS24 struct { F0 int; F1 int32; F2 int32 }
type GenS25 struct { F0 string; F1 int64 }

func main() {
    var v1 uint32 = 2147483647
    v2 := []string{("a" + ""), "hello", ""}
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]uint32)
    var v5 float64 = (1e10 * 1e10)
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        if false {
            for fi6 := 0; fi6 < 19; fi6++ {
                switch "a" {
                case "\n\t":
                    var v7 uint8 = 2
                    _ = v7
                default:
                    fmt.Println(v5)
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi6
            }
            for fi8 := 0; fi8 < 15; fi8++ {
                for fi9 := 0; fi9 < 4; fi9++ {
                    v1 = v1
                    _ = fi9
                }
                _ = fi8
            }
        }
    }
    if false {
        v1 = (v1 >> ((894) & 63))
        v5 += (v5 * ((0.0 / ((1.5 / 1.0) * v5)) + -2.25))
    } else {
        var v10 uint32 = 26
        if true {
            switch -84 {
            case 961:
                v2 = append(v2, "")
            }
        } else {
            v11 := GenS1{F0: 9223372036854775807}
            v12 := GenS2{F0: 497}
            _ = v11
            _ = v12
        }
        var v13 int8 = -40
        _ = v10
        _ = v13
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        for fi15 := 0; fi15 < 10; fi15++ {
            if true {
                if false {
                    var v16 uint32 = 2147483648
                    _ = v16
                }
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            v3 -= (v3 + (v3 * (^(v3))))
            _ = fi15
        }
        for fi17 := 0; fi17 < 3; fi17++ {
            v18 := GenS3{F0: -1, F1: -1}
            var v19 int8 = -103
            fmt.Println(fi17, v19, v1)
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi14
    }
    if false {
        var v20 float64 = (float64(0.0) / v5)
        _ = v20
    } else {
        for fi21 := 0; fi21 < 18; fi21++ {
            v5 = v5
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
            v22 := GenS2{F0: 55}
            _ = fi21
            _ = v22
        }
        var v23 int64 = (v3 | 42)
        for fi24 := 0; fi24 < 15; fi24++ {
            switch -827 {
            case 236:
                var v25 int16 = 0
                _ = v25
            default:
                v5 -= (v5 + v5)
            }
            var v26 bool = false
            if ((499 == 65535) && (-691 > -805)) {
                v27 := GenS1{F0: 444}
                _ = v27
            } else {
                switch 42 {
                case -1:
                    var v28 float64 = 0.1
                    _ = v28
                case 0:
                    var v29 int8 = 79
                    _ = v29
                default:
                    var v30 uint16 = 42
                    _ = v30
                }
            }
            _ = fi24
            _ = v26
        }
        _ = v23
    }
    v31 := GenS1{F0: (v3 & (^(v3)))}
    var v32 uint64 = 18446744073709551615
    v33 := GenS2{F0: v32}
    if false {
        for fi34 := 0; fi34 < 18; fi34++ {
            switch fi34 {
            case 1:
                if (((1e10 + 0.0) * 1e10) > 0.1) {
                    for fi35 := 0; fi35 < 1; fi35++ {
                        var v36 uint8 = 255
                        if ((false || true) && false) {
                            if false {
                                v37 := GenS2{F0: v32}
                                var v38 uint32 = (^(v1))
                                v4[-1] = 4294967295
                                var v39 float32 = (1.5 - -2.25)
                                _ = len(v2)
                                _ = v37
                                _ = v38
                                _ = v39
                            } else {
                                var v40 uint32 = (v1 * v1)
                                if false {
                                    v41 := GenS4{F0: (fi35 / ((fi34) | 1))}
                                    switch fi34 {
                                    case -67:
                                        var v42 uint32 = uint32(v36)
                                        _ = v42
                                    case 1:
                                        v43 := GenS5{F0: 255, F1: 18446744073709551615, F2: v3}
                                        _ = v43
                                    }
                                    v44 := GenS3{F0: -1, F1: int16(v3)}
                                    if (false || true) {
                                        v36 -= 42
                                        switch (fi35 * 2147483647) {
                                        case 128:
                                            var v45 float64 = 3.14159
                                            _ = v45
                                        case -337:
                                            var v46 uint16 = 65535
                                            _ = v46
                                        case 42:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            switch "hello" {
                                            case "":
                                                v47 := v4[fi35]
                                                _ = v47
                                            }
                                            var v48 uint64 = (v32 - 568)
                                            _ = v48
                                        }
                                        v40 = v40
                                    } else {
                                        if (((32767 < 1) && false) || (79 == 42)) {
                                            fmt.Println(v40)
                                            if (0.1 != 3.14159) {
                                                fmt.Println(v1)
                                            } else {
                                                if (!((("foo bar" + "") + "hello") != "hello")) {
                                                    v5 = (v5 / 3.14159)
                                                    v49 := GenS3{F0: -87, F1: 1}
                                                    v32 = v32
                                                    var v50 int32 = 2147483647
                                                    v2 = append(v2, "hello")
                                                    _ = v49
                                                    _ = v50
                                                } else {
                                                    switch fi34 {
                                                    case -1:
                                                        var v51 uint64 = 0
                                                        _ = v51
                                                    case 128:
                                                        if false {
                                                            v32 = 4294967295
                                                            fmt.Println(v1, v40)
                                                            switch fi35 {
                                                            case 666:
                                                                if true {
                                                                    switch (^(fi35)) {
                                                                    case 9223372036854775807:
                                                                        delete(v4, (fi34 & fi35))
                                                                    case -940:
                                                                        if false {
                                                                            switch "" {
                                                                            case "foo bar":
                                                                                if (fi34 >= 328) {
                                                                                    _ = len(v2)
                                                                                    v52 := GenS3{F0: -514, F1: 1}
                                                                                    if true {
                                                                                        if ("a" <= ("\n\t" + "a")) {
                                                                                            if false {
                                                                                                var v53 uint64 = (v32 ^ v32)
                                                                                                v54 := GenS2{F0: 18446744073709551615}
                                                                                                v55 := GenS5{F0: 42, F1: 42, F2: (v3 + -9223372036854775808)}
                                                                                                if (!(901 <= 2)) {
                                                                                                    v36 = v36
                                                                                                }
                                                                                                v2 = append(v2, "a")
                                                                                                _ = v53
                                                                                                _ = v54
                                                                                                _ = v55
                                                                                            } else {
                                                                                                v56 := GenS1{F0: 897}
                                                                                                _ = v56
                                                                                            }
                                                                                            if false {
                                                                                                delete(v4, 42)
                                                                                            } else {
                                                                                                v5 += 3.14159
                                                                                                fmt.Println(fi34, v1, v40)
                                                                                            }
                                                                                            fmt.Println(fi35, v1, v1)
                                                                                            if (true || (!false)) {
                                                                                                if ((!(!false)) || (((v3 >> ((0) & 63)) > v3) || ((!((611 <= 242) || false)) || ((^(v32)) >= v32)))) {
                                                                                                    v57 := GenS6{F0: v36}
                                                                                                    v36 -= 1
                                                                                                    _ = v57
                                                                                                } else {
                                                                                                    v3 -= 0
                                                                                                    v58 := v4[128]
                                                                                                    _ = v58
                                                                                                }
                                                                                                v59 := GenS2{F0: (v32 ^ (v32 & 0))}
                                                                                                _ = v59
                                                                                            } else {
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            v60 := GenS7{F0: (v36 ^ 0), F1: fi35}
                                                                                            _ = v60
                                                                                        }
                                                                                        v61 := v4[fi35]
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                        _ = v61
                                                                                    }
                                                                                    _ = v52
                                                                                }
                                                                                var v62 int32 = 1
                                                                                var v63 int = -2147483648
                                                                                v32 = 4294967295
                                                                                _ = v62
                                                                                _ = v63
                                                                            case "\n\t":
                                                                                if false {
                                                                                    var v64 string = ("\n\t" + "\n\t")
                                                                                    _ = v64
                                                                                } else {
                                                                                    var v65 float64 = (v5 - float64(3.14159))
                                                                                    _ = v65
                                                                                }
                                                                            case "":
                                                                                v5 = v5
                                                                            }
                                                                        }
                                                                    }
                                                                    v2 = append(v2, "\n\t")
                                                                    var v66 int = fi35
                                                                    _ = v66
                                                                }
                                                            case 0:
                                                                fmt.Println(v40)
                                                            case -302:
                                                                v4[367] = v40
                                                            }
                                                        }
                                                        v67 := GenS8{F0: v32, F1: fi34}
                                                        v36 -= (v36 & 26)
                                                        _ = v67
                                                    case 127:
                                                        _ = len(v2)
                                                    }
                                                    v4[int(v1)] = v40
                                                    v32 = v32
                                                    if len(v2) > 0 {
                                                        v2[0] = "hello"
                                                    }
                                                    if false {
                                                        fmt.Println(v3, v32, v3)
                                                        if (127 < 127) {
                                                            v40 += v40
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                    }
                                                }
                                                fmt.Println(v5)
                                                var v68 uint64 = v32
                                                v69 := GenS1{F0: v3}
                                                _ = v68
                                                _ = v69
                                            }
                                        }
                                        v70 := GenS5{F0: 65535, F1: 2, F2: (v3 & v3)}
                                        _ = v70
                                    }
                                    _ = v41
                                    _ = v44
                                } else {
                                    v71 := v4[fi35]
                                    switch 32767 {
                                    case -2147483648:
                                        v3 = (v3 % ((0) | 1))
                                        if false {
                                            v36 = v36
                                        } else {
                                            v40 = (v40 * v40)
                                            v71 = v71
                                        }
                                        v71 = (v71 | (v40 * uint32(v32)))
                                        if false {
                                            v3 += 142
                                            if false {
                                                v32 -= v32
                                            } else {
                                                if (true || false) {
                                                    var v72 uint64 = (v32 + (^(v32)))
                                                    v73 := GenS3{F0: 1, F1: 32767}
                                                    if (1 >= 175) {
                                                        fmt.Println(v5, v1)
                                                        if (false && true) {
                                                            fmt.Println(v71)
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        if (v36 <= (v36 % ((18) | 1))) {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        v74 := v4[fi35]
                                                        v75 := GenS4{F0: fi35}
                                                        _ = v74
                                                        _ = v75
                                                    } else {
                                                        var v76 uint32 = 699
                                                        _ = len(v2)
                                                        var v77 int32 = 42
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                    switch (fi35 & int(v3)) {
                                                    case 4294967295:
                                                        var v78 float32 = -2.25
                                                        _ = v78
                                                    default:
                                                        var v79 int32 = -1
                                                        _ = v79
                                                    }
                                                    if false {
                                                        var v80 int = (^(fi35))
                                                        fmt.Println(v71, v5, v1)
                                                        v81 := GenS9{F0: v36, F1: 0}
                                                        _ = v80
                                                        _ = v81
                                                    } else {
                                                        if true {
                                                            v71 += 0
                                                        } else {
                                                            v82 := GenS7{F0: 1, F1: (fi35 + fi35)}
                                                            if (false && false) {
                                                                v36 = v36
                                                                var v83 string = "a"
                                                                _ = v83
                                                            } else {
                                                                delete(v4, (fi34 * 128))
                                                                var v84 int8 = int8(v72)
                                                                if (!false) {
                                                                    fmt.Println(v5, v84, fi34)
                                                                    delete(v4, int(v84))
                                                                    delete(v4, 9223372036854775807)
                                                                }
                                                                _ = v84
                                                            }
                                                            _ = v82
                                                        }
                                                        switch (^(fi35)) {
                                                        case 4294967295:
                                                            if (1 == 0) {
                                                                if (("" <= ("\n\t" + ("foo bar" + ("\n\t" + ("" + ("" + "a")))))) || true) {
                                                                    var v85 uint64 = (v72 / ((v32) | 1))
                                                                    switch 9223372036854775807 {
                                                                    case -2147483648:
                                                                        var v86 int8 = 1
                                                                        _ = v86
                                                                    case 1:
                                                                        v87 := GenS10{F0: (v71 << ((330) & 63))}
                                                                        _ = v87
                                                                    }
                                                                    v32 = 560
                                                                    if true {
                                                                        fmt.Println(v40, v40, v71)
                                                                        v36 += v36
                                                                    }
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        var v88 float64 = (3.14159 + 1.5)
                                                                        _ = v88
                                                                    case "a":
                                                                        var v89 float64 = float64(0.1)
                                                                        _ = v89
                                                                    case "hello":
                                                                        v90 := v4[(fi34 | fi34)]
                                                                        _ = v90
                                                                    }
                                                                    _ = v85
                                                                } else {
                                                                    var v91 int = fi34
                                                                    delete(v4, (^(fi35)))
                                                                    _ = v91
                                                                }
                                                                var v92 int8 = -128
                                                                v32 -= v32
                                                                fmt.Println(v3, v72, v5)
                                                                _ = v92
                                                            }
                                                        default:
                                                            if true {
                                                                var v93 int = int(v72)
                                                                v36 -= uint8(v71)
                                                                _ = v93
                                                            } else {
                                                                v71 = (v40 | v71)
                                                            }
                                                        }
                                                        if (v3 == (v3 + (-(v3)))) {
                                                            if false {
                                                                v94 := GenS10{F0: (v40 << ((0) & 63))}
                                                                _ = v94
                                                            }
                                                            fmt.Println(v72, v1)
                                                            v95 := GenS11{F0: v36, F1: ("" + ("a" + "a"))}
                                                            _ = v95
                                                        }
                                                    }
                                                    _ = v72
                                                    _ = v73
                                                }
                                            }
                                        }
                                        if (979 >= 42) {
                                            v71 = 2147483647
                                            if (!(!true)) {
                                                switch fi34 {
                                                case -86:
                                                    delete(v4, fi35)
                                                    v2 = append(v2, "a")
                                                    if ((!true) && (865 >= 215)) {
                                                        var v96 int32 = -1
                                                        _ = v96
                                                    }
                                                case 1:
                                                    v32 = v32
                                                }
                                            } else {
                                                if true {
                                                    if ((1e10 != 0.1) || (65535 >= 255)) {
                                                        var v97 float64 = v5
                                                        v3 += (-(v3))
                                                        v5 = v5
                                                        _ = v97
                                                    }
                                                } else {
                                                    v98 := GenS10{F0: 65535}
                                                    fmt.Println(v3)
                                                    _ = v98
                                                }
                                                if false {
                                                    switch "" {
                                                    case "\n\t":
                                                        var v99 float32 = 3.14159
                                                        _ = v99
                                                    case "a":
                                                        v3 = (^(v3))
                                                    case "hello":
                                                        _ = len(v2)
                                                    }
                                                }
                                                if false {
                                                    v100 := GenS12{F0: (-(v3))}
                                                    v71 = (^(v71))
                                                    _ = v100
                                                } else {
                                                    v36 = v36
                                                }
                                                if false {
                                                    switch fi35 {
                                                    case 12:
                                                        v32 = v32
                                                    case -2147483648:
                                                        v3 = (v3 / ((v3) | 1))
                                                    default:
                                                        if false {
                                                            v36 -= (^(v36))
                                                        }
                                                    }
                                                } else {
                                                    if (0 > 179) {
                                                        var v101 int = (fi34 / (((fi34 & 1)) | 1))
                                                        v102 := GenS9{F0: 0, F1: 32767}
                                                        v4[(fi34 ^ 128)] = (^(v71))
                                                        v103 := GenS8{F0: 18446744073709551615, F1: (fi35 >> ((42) & 63))}
                                                        if (!false) {
                                                            if true {
                                                                if (42 != 42) {
                                                                    switch (fi35 ^ -1) {
                                                                    case -1:
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            if true {
                                                                                _ = len(v2)
                                                                                v104 := GenS13{F0: v101, F1: 1, F2: v71}
                                                                                if false {
                                                                                    v5 -= 0.1
                                                                                }
                                                                                if ((4294967295 != 18446744073709551615) || (!false)) {
                                                                                    v105 := GenS14{F0: -32768}
                                                                                    _ = v105
                                                                                }
                                                                                _ = len(v2)
                                                                                _ = v104
                                                                            } else {
                                                                                v106 := GenS9{F0: 128, F1: 32767}
                                                                                fmt.Println(fi35, v36)
                                                                                fmt.Println(v32, v32)
                                                                                var v107 int = (fi34 + fi35)
                                                                                _ = v106
                                                                                _ = v107
                                                                            }
                                                                        default:
                                                                            fmt.Println(fi35, v36)
                                                                        }
                                                                    case 32768:
                                                                        var v108 int = 1
                                                                        _ = v108
                                                                    default:
                                                                        v109 := GenS8{F0: v32, F1: fi35}
                                                                        _ = v109
                                                                    }
                                                                    v2 = append(v2, "a")
                                                                    switch 2147483647 {
                                                                    case -1:
                                                                        var v110 int16 = 42
                                                                        _ = v110
                                                                    default:
                                                                        if (int16(v32) > 0) {
                                                                            var v111 int32 = int32(fi34)
                                                                            if false {
                                                                                var v112 uint64 = v32
                                                                                _ = v112
                                                                            } else {
                                                                                var v113 int32 = v111
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = "\n\t"
                                                                                }
                                                                                v114 := GenS2{F0: v32}
                                                                                var v115 uint32 = (v40 ^ v40)
                                                                                if false {
                                                                                    v116 := GenS4{F0: fi35}
                                                                                    v3 = v3
                                                                                    var v117 uint8 = v36
                                                                                    _ = v116
                                                                                    _ = v117
                                                                                }
                                                                                _ = v113
                                                                                _ = v114
                                                                                _ = v115
                                                                            }
                                                                            _ = v111
                                                                        } else {
                                                                            v101 += fi34
                                                                        }
                                                                    }
                                                                } else {
                                                                    switch v101 {
                                                                    case 4294967295:
                                                                        if true {
                                                                            switch (v101 * (^(fi35))) {
                                                                            case 651:
                                                                                v4[fi35] = v40
                                                                            default:
                                                                                var v118 float64 = v5
                                                                                v32 -= (^(v32))
                                                                                var v119 uint32 = 65535
                                                                                _ = v118
                                                                                _ = v119
                                                                            }
                                                                            switch "foo bar" {
                                                                            case "a":
                                                                                v1 -= v1
                                                                            case "\n\t":
                                                                                v120 := v4[fi35]
                                                                                _ = v120
                                                                            }
                                                                        }
                                                                    case -2147483648:
                                                                        _ = len(v4)
                                                                    case -855:
                                                                        v121 := v4[9223372036854775807]
                                                                        v122 := GenS15{F0: v121, F1: 65535}
                                                                        if ((fi35 % (((fi35 >> ((688) & 63))) | 1)) < fi35) {
                                                                            var v123 bool = (!false)
                                                                            if v123 {
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                v124 := GenS1{F0: v3}
                                                                                v125 := GenS8{F0: (^(v32)), F1: 32767}
                                                                                _ = v124
                                                                                _ = v125
                                                                            } else {
                                                                                v126 := GenS1{F0: 490}
                                                                                _ = v126
                                                                            }
                                                                            var v127 int16 = 1000
                                                                            if v123 {
                                                                                var v128 float64 = ((v5 / ((((v5 - (((3.14159 * 1.5) + v5) * v5)) * v5) - -2.25) - 0.1)) * (v5 * 3.14159))
                                                                                v1 += v71
                                                                                _ = v128
                                                                            }
                                                                            if v123 {
                                                                                v129 := v4[fi35]
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    v129 -= v1
                                                                                default:
                                                                                    var v130 bool = (!v123)
                                                                                    v131 := GenS14{F0: (v127 % ((v127) | 1))}
                                                                                    v40 = v1
                                                                                    _ = v130
                                                                                    _ = v131
                                                                                }
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                var v132 int32 = 913
                                                                                _ = v129
                                                                                _ = v132
                                                                            }
                                                                            switch "a" {
                                                                            case "hello":
                                                                                if v123 {
                                                                                    v2 = append(v2, "")
                                                                                }
                                                                            default:
                                                                                var v133 int32 = 0
                                                                                _ = v133
                                                                            }
                                                                            _ = v123
                                                                            _ = v127
                                                                        }
                                                                        _ = v121
                                                                        _ = v122
                                                                    default:
                                                                        var v134 uint32 = uint32(v32)
                                                                        _ = v134
                                                                    }
                                                                    switch ("a" + "a") {
                                                                    case "hello":
                                                                        fmt.Println(v1, v32)
                                                                    case "foo bar":
                                                                        switch v101 {
                                                                        case -1:
                                                                            fmt.Println(v1, v40, v71)
                                                                        case -2147483648:
                                                                            if true {
                                                                                var v135 uint64 = 2
                                                                                switch ("a" + "foo bar") {
                                                                                case "":
                                                                                    var v136 uint16 = 2
                                                                                    _ = v136
                                                                                default:
                                                                                    var v137 int64 = v3
                                                                                    _ = v137
                                                                                }
                                                                                if false {
                                                                                    v2 = append(v2, "hello")
                                                                                }
                                                                                _ = v135
                                                                            }
                                                                        }
                                                                    case "":
                                                                        v40 += v1
                                                                        v138 := GenS16{F0: (!false), F1: false, F2: 255}
                                                                        var v139 float64 = v5
                                                                        _ = v138
                                                                        _ = v139
                                                                    }
                                                                    var v140 int64 = v3
                                                                    _ = v140
                                                                }
                                                                v3 = 9223372036854775807
                                                            } else {
                                                                if (((true || ((fi35 + (v101 + fi34)) == v101)) && false) || false) {
                                                                    fmt.Println(v3, v1)
                                                                }
                                                            }
                                                            v2 = append(v2, "hello")
                                                        }
                                                        switch (-(fi34)) {
                                                        case -2147483648:
                                                            if false {
                                                                fmt.Println(v40)
                                                                v141 := GenS17{F0: 84}
                                                                _ = v141
                                                            }
                                                        case 2147483647:
                                                            fmt.Println(v36)
                                                        case 504:
                                                            var v142 int8 = 42
                                                            _ = v142
                                                        }
                                                        _ = v101
                                                        _ = v102
                                                        _ = v103
                                                    }
                                                    if ((^(v1)) >= v71) {
                                                        v36 = (v36 / ((v36) | 1))
                                                        v32 = (v32 / (((v32 & v32)) | 1))
                                                        v5 = v5
                                                        v143 := GenS18{F0: 65535, F1: 1, F2: -474}
                                                        v1 = (^(v1))
                                                        v3 = v3
                                                        v144 := GenS19{F0: 42, F1: ("\n\t" + "hello"), F2: v36}
                                                        var v145 string = ("a" + "\n\t")
                                                        if (555 != (v3 ^ v3)) {
                                                            if true {
                                                                var v146 uint16 = 255
                                                                switch (fi35 & fi35) {
                                                                case 128:
                                                                    if ((v36 != uint8(v71)) && (1000 < -418)) {
                                                                        if (true || true) {
                                                                            v5 -= 1.5
                                                                            v71 = 1
                                                                        }
                                                                        _ = len(v4)
                                                                        var v147 int8 = 1
                                                                        _ = v147
                                                                    }
                                                                case 32767:
                                                                    if (false || (v3 < v3)) {
                                                                        switch (^(fi35)) {
                                                                        case 32768:
                                                                            if len(v2) > 0 {
                                                                                v2[0] = v145
                                                                            }
                                                                        case -408:
                                                                            if (true || (true || true)) {
                                                                                v71 = 0
                                                                            }
                                                                        case -1:
                                                                            var v148 uint32 = (v1 + 42)
                                                                            v146 = v146
                                                                            var v149 int = fi34
                                                                            _ = v148
                                                                            _ = v149
                                                                        }
                                                                    } else {
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        var v150 uint = 1
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        _ = v150
                                                                    }
                                                                case 127:
                                                                    v151 := GenS11{F0: 2, F1: "foo bar"}
                                                                    _ = v151
                                                                default:
                                                                    var v152 float32 = 0.0
                                                                    _ = v152
                                                                }
                                                                v3 -= v3
                                                                _ = v146
                                                            } else {
                                                                fmt.Println(v145)
                                                                if (true && (3.14159 > 1.5)) {
                                                                    var v153 int64 = -1
                                                                    _ = v153
                                                                } else {
                                                                    fmt.Println(v1, fi34, v40)
                                                                    v154 := GenS7{F0: (v36 ^ v36), F1: (fi35 >> ((2) & 63))}
                                                                    if false {
                                                                        delete(v4, -1)
                                                                        var v155 int8 = 15
                                                                        fmt.Println(v3, v155, v32)
                                                                        if (284 >= -369) {
                                                                            switch v145 {
                                                                            case "hello":
                                                                                fmt.Println(v36)
                                                                            default:
                                                                                var v156 uint8 = v36
                                                                                v3 = (v3 % (((^(v3))) | 1))
                                                                                _ = v156
                                                                            }
                                                                            switch fi35 {
                                                                            case 32768:
                                                                                var v157 uint32 = (v40 - 42)
                                                                                var v158 int16 = 1
                                                                                v157 -= v1
                                                                                _ = v157
                                                                                _ = v158
                                                                            case 128:
                                                                                if false {
                                                                                    v159 := GenS15{F0: (v1 << ((267) & 63)), F1: 577}
                                                                                    _ = v159
                                                                                } else {
                                                                                    var v160 uint16 = 34
                                                                                    v32 = v32
                                                                                    var v161 uint8 = v36
                                                                                    _ = v160
                                                                                    _ = v161
                                                                                }
                                                                            }
                                                                            fmt.Println(v32, v32)
                                                                        } else {
                                                                            v40 += (v1 / (((v40 * v71)) | 1))
                                                                            v155 = (v155 | (^(v155)))
                                                                            var v162 bool = true
                                                                            _ = v162
                                                                        }
                                                                        _ = v155
                                                                    }
                                                                    _ = v154
                                                                }
                                                                fmt.Println(fi34, v36, v145)
                                                                v32 = v32
                                                            }
                                                        }
                                                        _ = v143
                                                        _ = v144
                                                        _ = v145
                                                    }
                                                }
                                            }
                                            fmt.Println(v32, v32, v3)
                                        } else {
                                            v36 -= v36
                                            switch -1 {
                                            case -852:
                                                v2 = append(v2, ("hello" + "\n\t"))
                                            }
                                            if true {
                                                v32 += 42
                                                switch "\n\t" {
                                                case "\n\t":
                                                    var v163 float32 = (float32(float64(3.14159)) / 1.0)
                                                    _ = v163
                                                }
                                                var v164 bool = true
                                                v3 -= (v3 / ((42) | 1))
                                                v165 := GenS20{F0: -446}
                                                fmt.Println(v40)
                                                var v166 uint64 = (v32 * v32)
                                                _ = v164
                                                _ = v165
                                                _ = v166
                                            }
                                        }
                                    case 0:
                                        v71 = v71
                                    case 4294967295:
                                        if true {
                                            v36 += (^(v36))
                                            v5 += 3.14159
                                            v1 += 2147483648
                                            fmt.Println(v32)
                                            fmt.Println(v3, v36, v32)
                                            switch -265 {
                                            case 32767:
                                                if true {
                                                    v1 -= v1
                                                } else {
                                                    v1 += (v1 + (v71 % ((v71) | 1)))
                                                }
                                            }
                                        } else {
                                            v4[(fi34 * (fi34 << ((65535) & 63)))] = v40
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v167 := GenS21{F0: "foo bar", F1: fi35}
                                            _ = v167
                                        }
                                    default:
                                        var v168 uint = 2
                                        _ = v168
                                    }
                                    fmt.Println(v3)
                                    switch (-(fi35)) {
                                    case -2147483648:
                                        v169 := v4[fi34]
                                        _ = v169
                                    default:
                                        var v170 uint16 = 65535
                                        _ = v170
                                    }
                                    fmt.Println(v32)
                                    _ = v71
                                }
                                _ = v40
                            }
                        } else {
                            var v171 int32 = 42
                            _ = v171
                        }
                        _ = fi35
                        _ = v36
                    }
                    var v172 uint = 42
                    _ = v172
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            default:
                v2 = append(v2, "hello")
            }
            _ = fi34
        }
    } else {
        var v173 uint = 4294967295
        v173 += (^(v173))
        _ = v173
    }
    switch 0 {
    case -497:
        v1 = v1
    case 1:
        var v174 int = -2147483648
        _ = v174
    case 32767:
        var v175 int8 = -1
        _ = v175
    }
    switch -817 {
    case 574:
        fmt.Println(v3, v32)
    case 128:
        v176 := GenS15{F0: v1, F1: 65535}
        _ = v176
    case -849:
        switch 221 {
        case -859:
            var v177 int64 = (-(v3))
            _ = v177
        case -990:
            v178 := GenS1{F0: 0}
            _ = v178
        default:
            fmt.Println(v32, v32)
        }
    default:
        v32 += (v32 % (((v32 | 863)) | 1))
    }
    v179 := GenS22{F0: (v1 & (^(v1))), F1: 42}
    v180 := v4[int(v1)]
    v181 := GenS2{F0: uint64(v180)}
    var v182 bool = (0 > -549)
    if v182 {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi183 := 0; fi183 < 15; fi183++ {
            if (255 != 255) {
                v184 := GenS12{F0: v3}
                if v182 {
                    if false {
                        switch ("foo bar" + ("a" + "\n\t")) {
                        case "foo bar":
                            v185 := GenS23{F0: 99, F1: "foo bar"}
                            _ = v185
                        }
                        v3 -= (v3 & v3)
                        for fi186 := 0; fi186 < 16; fi186++ {
                            fmt.Println(v182)
                            _ = fi186
                        }
                    } else {
                        v32 += v32
                        v32 += 0
                    }
                    v1 = v1
                    switch fi183 {
                    case 127:
                        for fi187 := 0; fi187 < 15; fi187++ {
                            if ((((0 < 0) || (!v182)) && true) && v182) {
                                fmt.Println(v180, v180)
                            } else {
                                v188 := GenS12{F0: -497}
                                var v189 int32 = -956
                                var v190 string = ""
                                _ = v188
                                _ = v189
                                _ = v190
                            }
                            var v191 int = (fi187 * (fi183 - fi183))
                            _ = fi187
                            _ = v191
                        }
                    }
                }
                switch (-(fi183)) {
                case -2147483648:
                    fmt.Println(fi183, v182)
                default:
                    v192 := GenS17{F0: 127}
                    _ = v192
                }
                _ = v184
            } else {
                _ = len(v2)
                for fi193 := 0; fi193 < 19; fi193++ {
                    v5 = 1e10
                    var v194 uint = 751
                    v32 = (v32 / ((uint64(v180)) | 1))
                    if len(v2) > 0 {
                        v2[0] = ("hello" + (("foo bar" + (("hello" + "a") + ("a" + "\n\t"))) + ""))
                    }
                    var v195 uint8 = 128
                    _ = fi193
                    _ = v194
                    _ = v195
                }
                if false {
                    for fi196 := 0; fi196 < 9; fi196++ {
                        v197 := GenS15{F0: (v180 - (^(v1))), F1: 0}
                        v198 := v4[0]
                        _ = fi196
                        _ = v197
                        _ = v198
                    }
                }
            }
            _ = fi183
        }
    } else {
        if (!true) {
            if len(v2) > 0 {
                v2[0] = "a"
            }
        }
        switch "hello" {
        case "\n\t":
            for fi199 := 0; fi199 < 13; fi199++ {
                v200 := GenS7{F0: 2, F1: 32767}
                var v201 uint32 = (v1 - 471)
                _ = fi199
                _ = v200
                _ = v201
            }
        case "":
            v202 := GenS24{F0: 0, F1: 42, F2: 2147483647}
            _ = v202
        default:
            v203 := GenS25{F0: "hello", F1: (v3 >> ((42) & 63))}
            _ = v203
        }
        v3 = -469
    }
    v182 = (v32 >= v32)
    for fi204 := 0; fi204 < 19; fi204++ {
        v3 = v3
        _ = fi204
    }
    var v205 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v180:", v180)
    fmt.Println("v182:", v182)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v205:", v205)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v179
    _ = v180
    _ = v181
    _ = v182
    _ = v2
    _ = v205
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
