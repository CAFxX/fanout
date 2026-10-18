package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 bool; F1 int16; F2 bool }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 uint64; F1 int8 }
type GenS8 struct { F0 int64; F1 int; F2 string }
type GenS9 struct { F0 uint16; F1 uint; F2 int16 }
type GenS10 struct { F0 int64; F1 bool }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int32; F1 int; F2 string }
type GenS15 struct { F0 uint16; F1 int8 }

func main() {
    v1 := []int{32768, 128, 128, 1}
    var v2 uint32 = 386
    v3 := make(map[string]uint32)
    v4 := []string{}
    var v5 bool = true
    var v6 uint64 = 42
    var v7 uint16 = 42
    var v8 bool = v5
    switch "foo bar" {
    case "hello":
        if (2 <= 607) {
            switch ("foo bar" + ("\n\t" + "hello")) {
            case "a":
                switch 32768 {
                case -2147483648:
                    fmt.Println(v2, v5)
                default:
                    v6 = v6
                }
            case "":
                if v5 {
                    v9 := GenS1{F0: 18, F1: v6}
                    var v10 uint = 18446744073709551615
                    _ = v10
                    _ = v9
                } else {
                    var v11 uint32 = (v2 >> ((0) & 63))
                    var v12 int16 = 42
                    _ = v11
                    _ = v12
                }
            case "foo bar":
                if (v7 < v7) {
                    for fi13 := 0; fi13 < 12; fi13++ {
                        fmt.Println(v2, v2, v7)
                        _ = fi13
                    }
                    var v14 uint16 = 255
                    _ = v14
                }
            }
        } else {
            v4 = append(v4, ("a" + ("foo bar" + "hello")))
        }
    case "a":
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 bool = (true && (true || (fi15 == fi15)))
            delete(v3, (("foo bar" + ("a" + (("foo bar" + "\n\t") + (("" + "\n\t") + "a")))) + "a"))
            _ = fi15
            _ = v16
        }
    case "\n\t":
        delete(v3, ("" + ("a" + ("hello" + "hello"))))
    default:
        var v17 string = "a"
        _ = v17
    }
    if (4294967295 != (v6 * (v6 >> ((0) & 63)))) {
        v18 := GenS2{F0: v2, F1: (^(v2))}
        if v5 {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v6 += v6
            v3[("\n\t" + "foo bar")] = uint32(v7)
        }
        _ = v18
    } else {
        if v5 {
            for fi19 := 0; fi19 < 19; fi19++ {
                for fi20 := 0; fi20 < 4; fi20++ {
                    v21 := GenS1{F0: 63, F1: v6}
                    v8 = (!(v5 || true))
                    _ = fi20
                    _ = v21
                }
                _ = fi19
            }
            for fi22 := 0; fi22 < 1; fi22++ {
                switch "foo bar" {
                case "":
                    v23 := GenS1{F0: 0, F1: 278}
                    _ = v23
                default:
                    v6 += (v6 << ((446) & 63))
                }
                _ = fi22
            }
            var v24 uint32 = v2
            _ = v24
        } else {
            var v25 uint = 0
            var v26 uint64 = v6
            _ = v25
            _ = v26
        }
        var v27 int64 = -1
        var v28 uint = 4294967295
        _ = v27
        _ = v28
    }
    switch -9223372036854775808 {
    case 2147483647:
        for fi29 := 0; fi29 < 10; fi29++ {
            var v30 uint64 = 42
            v4 = append(v4, "foo bar")
            v8 = v5
            _ = fi29
            _ = v30
        }
    case -287:
        switch -872 {
        case 128:
            v31 := GenS2{F0: (v2 >> ((65535) & 63)), F1: v2}
            _ = v31
        case 32768:
            for fi32 := 0; fi32 < 18; fi32++ {
                v5 = (32767 >= 1000)
                for fi33 := 0; fi33 < 6; fi33++ {
                    if v8 {
                        var v34 int8 = 1
                        v34 += -88
                        var v35 uint = 18446744073709551615
                        fmt.Println(v34)
                        _ = v34
                        _ = v35
                    }
                    if ((v6 | v6) == v6) {
                        switch "a" {
                        case "":
                            v36 := GenS3{F0: 65, F1: v8}
                            _ = v36
                        case "hello":
                            switch fi33 {
                            case 0:
                                fmt.Println(fi33, v5)
                            case 941:
                                v37 := v3["hello"]
                                _ = v37
                            case 2147483647:
                                if v5 {
                                    fmt.Println(v5)
                                    _ = len(v3)
                                    if (!false) {
                                        var v38 int = (fi32 - (fi32 | -2147483648))
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        fmt.Println(v38)
                                        v5 = false
                                        _ = v38
                                    }
                                } else {
                                    switch (fi32 >> ((932) & 63)) {
                                    case 386:
                                        switch ("foo bar" + ("foo bar" + ("foo bar" + "hello"))) {
                                        case "":
                                            v7 -= 32767
                                        case "\n\t":
                                            v6 -= 290
                                        }
                                    case 32768:
                                        if (!(2147483647 < (^(fi32)))) {
                                            fmt.Println(fi32, v6, fi32)
                                        }
                                    }
                                    v3["hello"] = v2
                                    v8 = (1.5 >= (-2.25 + 0.1))
                                }
                                if false {
                                    var v39 float32 = float32(float64(float32(((0.0 + 0.0) * -2.25))))
                                    delete(v3, ("\n\t" + "hello"))
                                    if (2 >= 255) {
                                        if v5 {
                                            v40 := GenS1{F0: -128, F1: (v6 & (v6 / (((v6 * (v6 / (((v6 - (^(v6)))) | 1)))) | 1)))}
                                            fmt.Println(v8, v8, v39)
                                            _ = v40
                                        } else {
                                            var v41 int = (fi32 / ((fi32) | 1))
                                            v2 += 4294967295
                                            _ = v41
                                        }
                                    } else {
                                        var v42 int32 = -1
                                        _ = v42
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = fi33
                                    }
                                    var v43 bool = (485 >= fi33)
                                    _ = v39
                                    _ = v43
                                } else {
                                    _ = len(v3)
                                }
                                if v8 {
                                    v8 = (635 <= 1)
                                    v5 = true
                                } else {
                                    v2 += v2
                                    var v44 float64 = 1.5
                                    _ = v44
                                }
                            }
                        default:
                            var v45 int32 = 213
                            _ = v45
                        }
                        _ = len(v4)
                        v6 = 4294967295
                        v6 = uint64(v2)
                        fmt.Println(v6)
                        var v46 int64 = 536
                        _ = v46
                    } else {
                        switch "hello" {
                        case "a":
                            switch "\n\t" {
                            case "foo bar":
                                if (v8 || v8) {
                                    if ((!true) && ((((v2 % ((v2) | 1)) <= 65535) || true) || true)) {
                                        v7 -= (v7 - 65535)
                                        v47 := GenS4{F0: v8, F1: -976, F2: v8}
                                        var v48 int32 = -321
                                        _ = v47
                                        _ = v48
                                    } else {
                                        if true {
                                            var v49 uint32 = (v2 + 4294967295)
                                            if (-821 >= 0) {
                                                if (!true) {
                                                    var v50 uint64 = (^(v6))
                                                    v51 := GenS5{F0: 1, F1: (v2 / (((^(v2))) | 1))}
                                                    v52 := GenS6{F0: v5, F1: 18446744073709551615}
                                                    _ = v50
                                                    _ = v51
                                                    _ = v52
                                                } else {
                                                    var v53 uint = uint(v7)
                                                    _ = v53
                                                }
                                                var v54 uint = 4294967295
                                                v1 = append(v1, (fi33 >> ((4294967295) & 63)))
                                                _ = v54
                                            }
                                            if v5 {
                                                var v55 float32 = 1.5
                                                switch (-(fi32)) {
                                                case 9223372036854775807:
                                                    var v56 int64 = -9223372036854775808
                                                    _ = v56
                                                default:
                                                    v55 -= 1.5
                                                }
                                                _ = v55
                                            }
                                            _ = v49
                                        } else {
                                            var v57 uint8 = 127
                                            switch ((("hello" + "foo bar") + "hello") + ("foo bar" + "hello")) {
                                            case "\n\t":
                                                if ((127 < v57) || v5) {
                                                    var v58 bool = v8
                                                    v59 := GenS7{F0: 4294967295, F1: 0}
                                                    v58 = v5
                                                    _ = v58
                                                    _ = v59
                                                } else {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v6 -= v6
                                                }
                                            case "foo bar":
                                                if ((-1 >= 927) && v8) {
                                                    v5 = (!(0 > 100))
                                                }
                                            }
                                            if (v57 <= (v57 & 200)) {
                                                var v60 int32 = -2147483648
                                                var v61 int8 = -35
                                                _ = v60
                                                _ = v61
                                            } else {
                                                fmt.Println(v2, fi33, fi32)
                                                v62 := GenS6{F0: (false && true), F1: 2}
                                                var v63 int64 = 42
                                                _ = v62
                                                _ = v63
                                            }
                                            _ = v57
                                        }
                                        switch ("foo bar" + "a") {
                                        case "a":
                                            v8 = false
                                        case "":
                                            if (9 == 2) {
                                                fmt.Println(fi33, v8)
                                                v3[("\n\t" + "\n\t")] = v2
                                            }
                                        case "\n\t":
                                            fmt.Println(v7)
                                        default:
                                            if true {
                                                v6 = 4294967295
                                            }
                                        }
                                    }
                                    switch "a" {
                                    case "hello":
                                        v64 := GenS4{F0: (v8 || false), F1: -794, F2: v8}
                                        _ = v64
                                    default:
                                        v1 = append(v1, fi32)
                                        v7 = 255
                                        _ = len(v4)
                                    }
                                    v5 = true
                                    if false {
                                        var v65 uint32 = 144
                                        _ = v65
                                    } else {
                                        v8 = (875 != -9223372036854775808)
                                        if false {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        } else {
                                            v2 -= v2
                                            fmt.Println(v6, v5, v2)
                                            v7 -= (v7 | v7)
                                        }
                                    }
                                    if v5 {
                                        v7 += v7
                                    }
                                }
                                v5 = v5
                                switch ("\n\t" + "\n\t") {
                                case "a":
                                    if false {
                                        if ((("foo bar" + "") + (("\n\t" + "\n\t") + "hello")) >= ("foo bar" + ("" + "a"))) {
                                            var v66 uint32 = v2
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = v66
                                        } else {
                                            v2 -= (v2 | v2)
                                            v67 := GenS8{F0: 0, F1: -9223372036854775808, F2: "a"}
                                            v5 = (v8 && v5)
                                            _ = v67
                                        }
                                        v2 += (^(v2))
                                        if (!false) {
                                            switch (fi33 << ((uint(v2)) & 63)) {
                                            case 32768:
                                                v68 := GenS4{F0: ((float32(float64(3.14159)) - 1e10) > (3.14159 + 1e10)), F1: -661, F2: (false && ((fi32 % (((^(fi32))) | 1)) >= (fi32 >> ((1) & 63))))}
                                                _ = v68
                                            case -722:
                                                fmt.Println(v8, fi32)
                                            case 128:
                                                v69 := GenS9{F0: v7, F1: 42, F2: 1000}
                                                _ = v69
                                            }
                                            if v5 {
                                                if v8 {
                                                    if v8 {
                                                        switch "" {
                                                        case "a":
                                                            v70 := GenS4{F0: (2147483647 >= 0), F1: 0, F2: true}
                                                            _ = v70
                                                        case "foo bar":
                                                            if v8 {
                                                                v8 = (!false)
                                                            } else {
                                                                v8 = (18446744073709551615 < 255)
                                                                fmt.Println(v7, fi32, v7)
                                                                v71 := GenS10{F0: -1, F1: true}
                                                                _ = v71
                                                            }
                                                        }
                                                        v72 := GenS7{F0: v6, F1: -104}
                                                        _ = v72
                                                    }
                                                    if v8 {
                                                        if v8 {
                                                            v73 := GenS11{F0: 110}
                                                            var v74 int = 137
                                                            _ = v73
                                                            _ = v74
                                                        }
                                                        fmt.Println(v2, v5, fi33)
                                                    } else {
                                                        switch fi33 {
                                                        case 476:
                                                            fmt.Println(fi32)
                                                        case 128:
                                                            fmt.Println(v8)
                                                        default:
                                                            v75 := GenS12{F0: 46, F1: 65535, F2: v6}
                                                            v5 = true
                                                            var v76 int16 = -32768
                                                            var v77 uint8 = 34
                                                            _ = v75
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                        v3[((("foo bar" + ("hello" + ("" + "a"))) + "foo bar") + "a")] = (v2 / ((1) | 1))
                                                        v5 = true
                                                        var v78 int16 = 42
                                                        var v79 uint = 560
                                                        v78 = v78
                                                        v5 = ((!false) || v5)
                                                        v8 = (!true)
                                                        _ = v78
                                                        _ = v79
                                                    }
                                                }
                                            }
                                            var v80 bool = v8
                                            fmt.Println(v7, v6, v8)
                                            v3[(("a" + ("a" + ("foo bar" + "hello"))) + "hello")] = v2
                                            _ = v80
                                        } else {
                                            v81 := GenS7{F0: v6, F1: 103}
                                            switch (fi32 - fi32) {
                                            case 32768:
                                                var v82 int32 = 1
                                                var v83 bool = ((771 > (v2 / ((v2) | 1))) || v5)
                                                v5 = v83
                                                _ = v82
                                                _ = v83
                                            case 1:
                                                v84 := GenS2{F0: v2, F1: 2}
                                                _ = v84
                                            case 4294967295:
                                                var v85 int8 = -1
                                                _ = v85
                                            default:
                                                switch (("" + ("hello" + "a")) + "foo bar") {
                                                case "hello":
                                                    if v5 {
                                                        var v86 string = (("\n\t" + "a") + ((("hello" + "a") + "foo bar") + ("a" + "")))
                                                        _ = v86
                                                    } else {
                                                        if v5 {
                                                            var v87 bool = ("hello" != "foo bar")
                                                            switch fi32 {
                                                            case -9223372036854775808:
                                                                var v88 int16 = int16(fi32)
                                                                _ = v88
                                                            case 4294967295:
                                                                var v89 uint8 = 0
                                                                _ = v89
                                                            case 0:
                                                                if v5 {
                                                                    var v90 uint16 = uint16(v6)
                                                                    _ = v90
                                                                } else {
                                                                    if v5 {
                                                                        if (v87 && ((1.5 / (float32(0.1) / 0.1)) < (((1.5 * -2.25) * 1e10) / 1.0))) {
                                                                            if v87 {
                                                                                delete(v3, "foo bar")
                                                                                v6 += v6
                                                                                v2 = (v2 / (((v2 / ((v2) | 1))) | 1))
                                                                            }
                                                                            v2 -= 597
                                                                        }
                                                                        var v91 int32 = 42
                                                                        _ = v91
                                                                    }
                                                                    if (!(127 < 42)) {
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                    } else {
                                                                        v92 := GenS1{F0: -87, F1: (v6 - 18446744073709551615)}
                                                                        fmt.Println(v2, fi32, fi32)
                                                                        fmt.Println(fi33, v8, v5)
                                                                        _ = v92
                                                                    }
                                                                    if ((v7 >> ((0) & 63)) == 32767) {
                                                                        v93 := GenS2{F0: 1, F1: v2}
                                                                        _ = v93
                                                                    }
                                                                }
                                                            }
                                                            fmt.Println(v6, v87, v87)
                                                            _ = v87
                                                        } else {
                                                            var v94 uint = 414
                                                            if (("\n\t" + "\n\t") != "hello") {
                                                                switch "\n\t" {
                                                                case "":
                                                                    fmt.Println(fi32)
                                                                case "hello":
                                                                    var v95 uint = (^(v94))
                                                                    _ = v95
                                                                }
                                                                v5 = (4294967295 > uint64(v7))
                                                                var v96 int64 = 1
                                                                _ = v96
                                                            }
                                                            _ = v94
                                                        }
                                                        var v97 int = (fi33 + fi32)
                                                        _ = v97
                                                    }
                                                }
                                            }
                                            var v98 int32 = 2147483647
                                            var v99 string = "a"
                                            v6 -= v6
                                            _ = v81
                                            _ = v98
                                            _ = v99
                                        }
                                        v5 = (float64(((3.14159 * (1e10 * (float32((0.0 - -2.25)) * -2.25))) + -2.25)) >= (3.14159 - (((1.5 / -2.25) / -2.25) - 0.0)))
                                        if (2147483647 == 42) {
                                            switch (fi33 * fi32) {
                                            case 1:
                                                v2 = 4294967295
                                                var v100 bool = v8
                                                _ = v100
                                            default:
                                                var v101 int = (fi33 | 0)
                                                _ = v101
                                            }
                                            var v102 uint32 = (v2 | (v2 << ((2) & 63)))
                                            var v103 uint16 = (v7 & 1)
                                            _ = v102
                                            _ = v103
                                        } else {
                                            v8 = v5
                                            v104 := GenS13{F0: -1}
                                            var v105 float64 = 1e10
                                            _ = v104
                                            _ = v105
                                        }
                                    }
                                case "hello":
                                    v106 := GenS11{F0: 379}
                                    _ = v106
                                }
                            case "\n\t":
                                var v107 uint64 = 4294967295
                                fmt.Println(v7)
                                _ = v107
                            case "hello":
                                v7 = (v7 + 808)
                            default:
                                v5 = false
                            }
                        }
                        var v108 bool = v8
                        _ = v108
                    }
                    if (v7 == 32767) {
                        var v109 float64 = 1.5
                        var v110 float32 = 0.0
                        switch "" {
                        case "foo bar":
                            v111 := GenS14{F0: 1, F1: fi33, F2: (((("a" + "a") + "\n\t") + "hello") + "a")}
                            _ = v111
                        }
                        _ = v109
                        _ = v110
                    } else {
                        v112 := GenS15{F0: 32767, F1: 42}
                        v6 = 0
                        var v113 int16 = -1
                        _ = v112
                        _ = v113
                    }
                    _ = fi33
                }
                _ = fi32
            }
        default:
            for fi114 := 0; fi114 < 7; fi114++ {
                _ = len(v3)
                _ = fi114
            }
        }
    default:
        var v115 bool = true
        _ = v115
    }
    fmt.Println(v8, v8)
    var v116 uint32 = 1
    var v117 uint = 65535
    if ((v117 % (((v117 / ((972) | 1))) | 1)) >= (v117 / (((v117 + (v117 / (((v117 | v117)) | 1)))) | 1))) {
        fmt.Println(v2)
        if v8 {
            for fi118 := 0; fi118 < 19; fi118++ {
                if (-2.25 >= (((((3.14159 * -2.25) + 1.5) / 1e10) + 0.1) * 3.14159)) {
                    if ((!(18446744073709551615 < (^(v117)))) && v5) {
                        if v5 {
                            v119 := GenS14{F0: -758, F1: fi118, F2: ""}
                            _ = v119
                        } else {
                            if (255 < v7) {
                                if ("\n\t" > "hello") {
                                    if (false && (280 != v117)) {
                                        v6 += (v6 ^ (v6 - (v6 | v6)))
                                        for fi120 := 0; fi120 < 6; fi120++ {
                                            fmt.Println(v8, v2)
                                            _ = fi120
                                        }
                                        if (v5 || (-1 <= 1)) {
                                            v121 := GenS8{F0: 1, F1: (fi118 / (((fi118 ^ (^(fi118)))) | 1)), F2: "hello"}
                                            _ = v121
                                        }
                                    }
                                }
                            } else {
                                fmt.Println(v117, v7, v117)
                            }
                        }
                    } else {
                        for fi122 := 0; fi122 < 18; fi122++ {
                            var v123 int = fi122
                            v6 += 0
                            _ = fi122
                            _ = v123
                        }
                    }
                } else {
                    if true {
                        fmt.Println(v5, v7)
                        fmt.Println(v6)
                    }
                    var v124 uint8 = 136
                    _ = v124
                }
                if (v8 || (uint8(v2) <= 42)) {
                    if true {
                        var v125 uint8 = 130
                        _ = len(v1)
                        _ = v125
                    } else {
                        if (v5 || (1 < int16(v117))) {
                            v126 := GenS13{F0: 1}
                            var v127 bool = true
                            if (2147483647 >= 42) {
                                v117 = (v117 | (^(v117)))
                                v8 = v127
                                v127 = false
                            }
                            _ = v126
                            _ = v127
                        }
                        v4 = append(v4, "a")
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    delete(v3, "\n\t")
                    v1 = append(v1, 42)
                    var v128 int32 = -1
                    _ = v128
                }
                _ = fi118
            }
        } else {
            v8 = false
            var v129 int16 = 32767
            _ = v129
        }
    } else {
        var v130 int32 = int32(v6)
        for fi131 := 0; fi131 < 3; fi131++ {
            var v132 int8 = -128
            if ((128 >= int(v130)) || true) {
                v2 += (v116 + 299)
                var v133 float32 = -2.25
                fmt.Println(v7)
                _ = v133
            } else {
                delete(v3, ("" + "foo bar"))
            }
            for fi134 := 0; fi134 < 13; fi134++ {
                var v135 int16 = 1
                _ = len(v1)
                _ = fi134
                _ = v135
            }
            _ = fi131
            _ = v132
        }
        _ = v130
    }
    fmt.Println(v6, v8, v8)
    if (-257 > 4294967295) {
        switch "a" {
        case "":
            if (v116 <= (^(v116))) {
                var v136 int16 = 0
                _ = v136
            }
        case "a":
            fmt.Println(v5, v117)
        default:
            if true {
                v117 = (v117 & 305)
                var v137 int64 = 1
                _ = v137
            } else {
                v116 = (v116 - v116)
                for fi138 := 0; fi138 < 6; fi138++ {
                    fmt.Println(v116, v117, v8)
                    _ = fi138
                }
                var v139 uint16 = 65535
                _ = v139
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v116:", v116)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v116
    _ = v117
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
