package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 uint16; F2 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int; F1 uint8 }
type GenS7 struct { F0 uint; F1 bool }
type GenS8 struct { F0 uint; F1 int32 }
type GenS9 struct { F0 bool; F1 string }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint64; F1 int32 }
type GenS12 struct { F0 uint32; F1 uint64; F2 bool }
type GenS13 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS14 struct { F0 int64; F1 int8; F2 uint16 }
type GenS15 struct { F0 bool; F1 int8; F2 bool }

func main() {
    v1 := []string{((("foo bar" + "foo bar") + ("hello" + "")) + ((("hello" + (("\n\t" + "a") + "a")) + "\n\t") + "hello")), "foo bar", ("hello" + "foo bar")}
    var v2 uint8 = 128
    v3 := []string{"a"}
    v4 := make(map[uint32]uint32)
    var v5 uint = 2
    var v6 uint = (v5 + (v5 % ((v5) | 1)))
    v7 := GenS1{F0: 50, F1: -32768, F2: 9223372036854775807}
    v8 := GenS1{F0: 255, F1: -1, F2: 42}
    v9 := GenS1{F0: 1, F1: 802, F2: -1}
    v5 = (v5 & 372)
    var v10 uint32 = 2
    var v11 uint8 = (v2 * 128)
    v11 = 1
    if (42 <= 292) {
        switch "foo bar" {
        case "\n\t":
            switch 127 {
            case 32768:
                switch 9223372036854775807 {
                case -111:
                    var v12 int8 = 100
                    _ = v12
                case 463:
                    v3 = append(v3, "foo bar")
                }
            case 128:
                var v13 int8 = 0
                _ = v13
            default:
                v5 = v5
            }
        case "a":
            switch "\n\t" {
            case "hello":
                delete(v4, (v10 << (((v6 - v5)) & 63)))
            case "a":
                for fi14 := 0; fi14 < 20; fi14++ {
                    if (0 < 302) {
                        v6 += (v5 ^ v6)
                    } else {
                        for fi15 := 0; fi15 < 1; fi15++ {
                            v5 += 42
                            if true {
                                fmt.Println(v2)
                                fmt.Println(v5)
                            } else {
                                var v16 int32 = -2147483648
                                var v17 uint32 = 352
                                _ = v16
                                _ = v17
                            }
                            _ = fi15
                        }
                    }
                    if (false || true) {
                        var v18 float64 = float64((-2.25 / float32(((1e10 - 1e10) - (((0.1 - 0.0) + float64((0.0 + ((0.1 * 0.1) + (0.0 / (0.0 + 0.1)))))) * (float64((3.14159 + 0.0)) * 0.1))))))
                        switch "" {
                        case "":
                            if true {
                                v2 += v2
                            } else {
                                var v19 int64 = -278
                                for fi20 := 0; fi20 < 15; fi20++ {
                                    if (!(!((0 <= 584) && false))) {
                                        v21 := GenS1{F0: 989, F1: -1, F2: -1}
                                        _ = v21
                                    }
                                    if (false || (!true)) {
                                        switch int(v10) {
                                        case -9223372036854775808:
                                            switch 2147483647 {
                                            case 128:
                                                switch fi20 {
                                                case -9223372036854775808:
                                                    v5 = (^(v5))
                                                case 54:
                                                    v22 := GenS1{F0: 473, F1: 1000, F2: v19}
                                                    _ = v22
                                                case 1:
                                                    v23 := GenS2{F0: (v10 % ((v10) | 1)), F1: v5}
                                                    _ = v23
                                                default:
                                                    if (("foo bar" > (("hello" + (("\n\t" + (("a" + "\n\t") + "foo bar")) + ("" + ("foo bar" + (("" + ("foo bar" + ("" + "\n\t"))) + "hello"))))) + "foo bar")) && true) {
                                                        v19 = (v19 | (v19 ^ v19))
                                                        var v24 int8 = 100
                                                        switch "a" {
                                                        case "a":
                                                            switch (-(fi14)) {
                                                            case 32768:
                                                                v18 = (v18 * ((((v18 + 1.5) + 0.1) / (v18 * 0.1)) - v18))
                                                            default:
                                                                fmt.Println(fi14, v19, v10)
                                                            }
                                                        case "hello":
                                                            switch -2147483648 {
                                                            case 127:
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    if false {
                                                                        var v25 uint32 = v10
                                                                        _ = v25
                                                                    }
                                                                case "a":
                                                                    var v26 int16 = -275
                                                                    _ = v26
                                                                case "":
                                                                    fmt.Println(v11, fi20, v24)
                                                                default:
                                                                    v27 := GenS3{F0: 1}
                                                                    _ = v27
                                                                }
                                                            }
                                                        case "foo bar":
                                                            if (true && false) {
                                                                if (0 >= fi20) {
                                                                    fmt.Println(v24)
                                                                }
                                                                fmt.Println(fi20)
                                                            } else {
                                                                fmt.Println(fi14)
                                                                _ = len(v3)
                                                                fmt.Println(v18)
                                                            }
                                                        default:
                                                            if (!(((1e10 * (v18 / 1e10)) < v18) || true)) {
                                                                v28 := GenS1{F0: 1, F1: 32767, F2: (v19 >> (((v5 & 255)) & 63))}
                                                                _ = v28
                                                            }
                                                        }
                                                        _ = v24
                                                    } else {
                                                        var v29 float32 = (1.5 + 3.14159)
                                                        fmt.Println(v2, v19, fi14)
                                                        _ = v29
                                                    }
                                                }
                                            default:
                                                if len(v3) > 0 {
                                                    v3[0] = ""
                                                }
                                                v6 -= v5
                                                v6 = 738
                                                v11 += 255
                                                var v30 uint64 = 0
                                                _ = v30
                                            }
                                        default:
                                            var v31 uint = v5
                                            _ = v31
                                        }
                                        var v32 float64 = (float64((1e10 * 1.5)) * v18)
                                        if (true && true) {
                                            if true {
                                                v33 := GenS4{F0: (v6 - v5), F1: 65535, F2: 65535}
                                                var v34 uint16 = 170
                                                _ = v33
                                                _ = v34
                                            }
                                            v18 = float64((0.1 + 1.5))
                                        }
                                        v18 = v32
                                        _ = v32
                                    } else {
                                        var v35 float32 = 0.1
                                        var v36 int32 = -2147483648
                                        _ = v35
                                        _ = v36
                                    }
                                    fmt.Println(v18, v11, v18)
                                    if true {
                                        switch "hello" {
                                        case "a":
                                            var v37 uint = 601
                                            _ = v37
                                        case "foo bar":
                                            fmt.Println(v19, v11)
                                        }
                                    }
                                    _ = fi20
                                }
                                v5 = v5
                                _ = v19
                            }
                        default:
                            v5 = (v5 | v5)
                        }
                        var v38 bool = true
                        _ = v18
                        _ = v38
                    } else {
                        v2 += (v2 * (v2 << ((v5) & 63)))
                        v6 += (v6 + v5)
                        v2 = v11
                    }
                    var v39 float64 = (-2.25 * 3.14159)
                    _ = fi14
                    _ = v39
                }
            case "foo bar":
                v5 += (v6 | (v5 - (v5 % ((42) | 1))))
            default:
                switch -1 {
                case -915:
                    fmt.Println(v10, v11, v6)
                default:
                    v40 := GenS5{F0: (v6 % ((4294967295) | 1))}
                    _ = v40
                }
            }
        case "hello":
            var v41 string = ""
            _ = v41
        }
        switch 4294967295 {
        case -178:
            fmt.Println(v11)
        default:
            for fi42 := 0; fi42 < 2; fi42++ {
                if false {
                    var v43 string = "hello"
                    var v44 int = (fi42 - fi42)
                    _ = v43
                    _ = v44
                } else {
                    v45 := GenS6{F0: (fi42 << ((uint(v10)) & 63)), F1: v11}
                    if (true && (2 <= 109)) {
                        switch 32768 {
                        case -102:
                            if true {
                                v46 := GenS3{F0: 1}
                                _ = v46
                            }
                        case 4294967295:
                            v47 := GenS7{F0: v5, F1: false}
                            _ = v47
                        case -567:
                            v6 = uint(v10)
                        default:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        for fi48 := 0; fi48 < 6; fi48++ {
                            var v49 uint32 = (v10 / ((2147483647) | 1))
                            _ = fi48
                            _ = v49
                        }
                        var v50 int32 = -1
                        _ = v50
                    } else {
                        fmt.Println(v6, v6, v6)
                        for fi51 := 0; fi51 < 2; fi51++ {
                            fmt.Println(v11, fi42, v2)
                            _ = fi51
                        }
                    }
                    var v52 bool = false
                    _ = v45
                    _ = v52
                }
                _ = fi42
            }
        }
    } else {
        fmt.Println(v2)
        var v53 float32 = 3.14159
        _ = v53
    }
    for fi54 := 0; fi54 < 10; fi54++ {
        if false {
            switch fi54 {
            case -342:
                if false {
                    switch (fi54 - fi54) {
                    case 0:
                        switch ("a" + "hello") {
                        case "\n\t":
                            var v55 int8 = 100
                            _ = v55
                        }
                    case -873:
                        for fi56 := 0; fi56 < 15; fi56++ {
                            switch ("a" + "\n\t") {
                            case "hello":
                                v10 = (v10 ^ v10)
                            case "":
                                var v57 float64 = float64(-2.25)
                                _ = v57
                            }
                            if true {
                                v58 := GenS7{F0: (^(v6)), F1: false}
                                v59 := GenS4{F0: 809, F1: 2, F2: v6}
                                var v60 float64 = ((1.5 * (-2.25 / 1.0)) - 1.5)
                                _ = v58
                                _ = v59
                                _ = v60
                            } else {
                                if false {
                                    v11 -= (v11 + 128)
                                    if true {
                                        if (100 <= 42) {
                                            var v61 uint = v6
                                            if (uint8(v6) >= v11) {
                                                if true {
                                                    var v62 bool = false
                                                    var v63 float32 = (float32((float64(1e10) * 1.5)) * (float32(float64(float32(((1.5 - ((float64((0.1 + -2.25)) + ((((1.5 + -2.25) - (-2.25 / 1.0)) - 3.14159) / 0.1)) + (float64(-2.25) * 1.5))) / -2.25)))) / (1.5 / (3.14159 + (((1e10 * 3.14159) + 0.0) * (3.14159 * 1e10))))))
                                                    _ = v62
                                                    _ = v63
                                                }
                                                v6 = (v6 & 4294967295)
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v4[v10] = (^(v10))
                                                switch (fi54 - fi54) {
                                                case 9223372036854775807:
                                                    switch (fi54 - -1) {
                                                    case 128:
                                                        switch fi56 {
                                                        case 0:
                                                            var v64 int = fi54
                                                            _ = v64
                                                        case 32767:
                                                            var v65 string = "a"
                                                            _ = v65
                                                        }
                                                    case 2147483647:
                                                        v2 -= (v2 | v11)
                                                    case 9223372036854775807:
                                                        _ = len(v4)
                                                    default:
                                                        var v66 int = fi56
                                                        _ = v66
                                                    }
                                                case -838:
                                                    var v67 uint64 = 0
                                                    var v68 int = (fi56 * fi54)
                                                    v69 := GenS3{F0: -405}
                                                    _ = v67
                                                    _ = v68
                                                    _ = v69
                                                case 2147483647:
                                                    if true {
                                                        fmt.Println(v10, v11)
                                                        v2 = 0
                                                        delete(v4, (^(v10)))
                                                    }
                                                default:
                                                    v6 += v6
                                                    var v70 uint16 = uint16(v10)
                                                    _ = v70
                                                }
                                            } else {
                                                delete(v4, v10)
                                            }
                                            _ = v61
                                        }
                                    }
                                    v11 = v2
                                    fmt.Println(fi56, fi54)
                                    fmt.Println(v5)
                                    _ = len(v1)
                                    var v71 int64 = 0
                                    _ = v71
                                } else {
                                    _ = len(v4)
                                    v5 += (^(v5))
                                    if true {
                                        switch "a" {
                                        case "hello":
                                            if true {
                                                fmt.Println(v6, fi54)
                                                var v72 int16 = 1000
                                                fmt.Println(v5, fi54)
                                                if len(v3) > 0 {
                                                    v3[0] = "a"
                                                }
                                                v72 = -562
                                                _ = v72
                                            }
                                        case "foo bar":
                                            if (true && true) {
                                                v5 -= v5
                                                v4[uint32(v2)] = (^(v10))
                                                v2 += v2
                                                v11 = 42
                                                switch (-(fi56)) {
                                                case 128:
                                                    v10 = v10
                                                }
                                            } else {
                                                v6 += (v5 >> ((v6) & 63))
                                                if false {
                                                    _ = len(v4)
                                                    delete(v4, v10)
                                                }
                                                var v73 bool = (true || false)
                                                switch fi56 {
                                                case 718:
                                                    var v74 bool = (324 == 42)
                                                    _ = v74
                                                }
                                                _ = v73
                                            }
                                        case "a":
                                            switch "" {
                                            case "\n\t":
                                                if true {
                                                    switch "\n\t" {
                                                    case "a":
                                                        v2 -= v2
                                                        if (float32((0.1 - (((-2.25 - 0.1) / 1.5) - (float64((0.0 / -2.25)) / (0.0 - 1e10))))) != float32(1e10)) {
                                                            switch fi54 {
                                                            case 32768:
                                                                if true {
                                                                    var v75 int64 = 0
                                                                    var v76 uint32 = (v10 << (((v5 & v6)) & 63))
                                                                    _ = v75
                                                                    _ = v76
                                                                }
                                                            case -666:
                                                                v77 := GenS8{F0: 14, F1: -255}
                                                                _ = v77
                                                            default:
                                                                var v78 float32 = float32(1.5)
                                                                _ = v78
                                                            }
                                                            if false {
                                                                if ((0.1 - 1e10) == -2.25) {
                                                                    v11 -= v2
                                                                    if true {
                                                                        if false {
                                                                            v11 = (v11 - v2)
                                                                        }
                                                                    } else {
                                                                        v11 -= v2
                                                                    }
                                                                    v79 := GenS7{F0: (^(v6)), F1: (((((!(-1 >= 42)) && true) || true) && false) || (2 < 138))}
                                                                    _ = v79
                                                                }
                                                            }
                                                        } else {
                                                            var v80 int64 = 830
                                                            if false {
                                                                var v81 int8 = -110
                                                                var v82 int64 = -9223372036854775808
                                                                v80 -= -9223372036854775808
                                                                var v83 bool = false
                                                                _ = v81
                                                                _ = v82
                                                                _ = v83
                                                            }
                                                            _ = v80
                                                        }
                                                        v84 := v4[(v10 & (v10 % (((v10 >> ((v5) & 63))) | 1)))]
                                                        var v85 uint16 = 42
                                                        v4[(v10 % ((uint32(v2)) | 1))] = (v10 << (((v6 + 4294967295)) & 63))
                                                        v86 := GenS9{F0: true, F1: "foo bar"}
                                                        _ = v84
                                                        _ = v85
                                                        _ = v86
                                                    case "hello":
                                                        v2 -= v2
                                                    default:
                                                        var v87 uint64 = 4294967295
                                                        _ = v87
                                                    }
                                                    v2 = (v11 - 42)
                                                } else {
                                                    v88 := GenS5{F0: 202}
                                                    var v89 int8 = -128
                                                    if false {
                                                        if len(v1) > 0 {
                                                            v1[0] = ("hello" + (("hello" + "\n\t") + ""))
                                                        }
                                                        v10 = v10
                                                        v2 -= (v2 >> ((v5) & 63))
                                                    } else {
                                                        if (!(0 < 0)) {
                                                            v4[v10] = 0
                                                        } else {
                                                            v6 = v5
                                                        }
                                                        v90 := v4[v10]
                                                        var v91 uint64 = 467
                                                        _ = v90
                                                        _ = v91
                                                    }
                                                    _ = v88
                                                    _ = v89
                                                }
                                            case "":
                                                var v92 int32 = 0
                                                _ = v92
                                            default:
                                                v3 = append(v3, "a")
                                            }
                                        default:
                                            if (!(true && (int32(v5) <= -2147483648))) {
                                                var v93 int8 = -69
                                                var v94 bool = ((v10 - 1) == (v10 * 4294967295))
                                                if false {
                                                    if (!(!((!false) && (v94 || v94)))) {
                                                        v2 = v2
                                                        v95 := GenS10{F0: (float32(float64(0.0)) > (1.5 / 1.0))}
                                                        var v96 int = fi54
                                                        if v94 {
                                                            delete(v4, v10)
                                                            v5 = (^(v5))
                                                            if false {
                                                                v6 += (^(v6))
                                                                v97 := GenS11{F0: 212, F1: -1}
                                                                if len(v1) > 0 {
                                                                    v1[0] = "\n\t"
                                                                }
                                                                _ = v97
                                                            }
                                                        }
                                                        if v94 {
                                                            var v98 uint32 = v10
                                                            _ = v98
                                                        }
                                                        _ = v95
                                                        _ = v96
                                                    } else {
                                                        v99 := GenS12{F0: v10, F1: 4294967295, F2: v94}
                                                        if len(v3) > 0 {
                                                            v3[0] = ""
                                                        }
                                                        if v94 {
                                                            v5 = v5
                                                            var v100 uint64 = 42
                                                            fmt.Println(v11)
                                                            _ = v100
                                                        }
                                                        _ = v99
                                                    }
                                                }
                                                if (814 < (v6 << (((v5 % ((65535) | 1))) & 63))) {
                                                    var v101 float64 = -2.25
                                                    v5 = (v6 - 0)
                                                    switch fi56 {
                                                    case 1:
                                                        if v94 {
                                                            v101 += ((3.14159 + float64(3.14159)) - v101)
                                                            v102 := GenS13{F0: v2, F1: 0, F2: 697}
                                                            fmt.Println(v11)
                                                            _ = v102
                                                        } else {
                                                            var v103 uint64 = 42
                                                            _ = v103
                                                        }
                                                    default:
                                                        v11 -= uint8(v10)
                                                    }
                                                    _ = v101
                                                } else {
                                                    switch "" {
                                                    case "a":
                                                        if (!v94) {
                                                            var v104 uint8 = v11
                                                            fmt.Println(v5, v11)
                                                            fmt.Println(v10, v2)
                                                            _ = v104
                                                        } else {
                                                            v2 = v11
                                                            v5 = (v5 / ((v6) | 1))
                                                            if v94 {
                                                                v10 -= 65535
                                                                v105 := GenS4{F0: v5, F1: 65535, F2: 837}
                                                                if true {
                                                                    fmt.Println(v94)
                                                                    v106 := GenS14{F0: 42, F1: (v93 << (((v6 << ((390) & 63))) & 63)), F2: 653}
                                                                    var v107 bool = (9223372036854775807 >= 9223372036854775807)
                                                                    v93 = -37
                                                                    switch "" {
                                                                    case "":
                                                                        v108 := v4[(v10 / ((v10) | 1))]
                                                                        _ = v108
                                                                    }
                                                                    _ = v106
                                                                    _ = v107
                                                                } else {
                                                                    var v109 uint16 = 255
                                                                    delete(v4, v10)
                                                                    _ = v109
                                                                }
                                                                _ = v105
                                                            }
                                                        }
                                                    case "foo bar":
                                                        delete(v4, v10)
                                                    default:
                                                        fmt.Println(v5, fi56)
                                                    }
                                                    if v94 {
                                                        v10 += v10
                                                    } else {
                                                        delete(v4, 42)
                                                        var v110 int64 = 1
                                                        _ = v110
                                                    }
                                                    v5 = (v6 ^ v5)
                                                }
                                                if (!v94) {
                                                    var v111 uint32 = 42
                                                    var v112 uint64 = uint64(v2)
                                                    if v94 {
                                                        var v113 int8 = (v93 | v93)
                                                        var v114 uint16 = 529
                                                        v114 -= (v114 + v114)
                                                        _ = v113
                                                        _ = v114
                                                    }
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = v111
                                                    _ = v112
                                                } else {
                                                    fmt.Println(fi54, v6)
                                                    fmt.Println(v2, v5)
                                                    var v115 int16 = 299
                                                    _ = v115
                                                }
                                                _ = v93
                                                _ = v94
                                            }
                                        }
                                    } else {
                                        v11 = 8
                                        if true {
                                            v6 = v6
                                            if (!(fi54 >= 2147483647)) {
                                                if (671 <= -2147483648) {
                                                    var v116 uint8 = v2
                                                    switch (-(fi56)) {
                                                    case 127:
                                                        var v117 int32 = 2147483647
                                                        _ = v117
                                                    case -1:
                                                        _ = len(v4)
                                                    }
                                                    switch (("\n\t" + "hello") + "foo bar") {
                                                    case "a":
                                                        v4[v10] = (v10 | uint32(v11))
                                                    case "\n\t":
                                                        v10 += 1
                                                    default:
                                                        switch fi56 {
                                                        case 1:
                                                            v116 += 13
                                                        case 197:
                                                            var v118 string = "\n\t"
                                                            _ = v118
                                                        }
                                                    }
                                                    _ = v116
                                                }
                                            }
                                        }
                                        var v119 int16 = 1000
                                        _ = v119
                                    }
                                }
                                var v120 uint64 = 0
                                _ = v120
                            }
                            v5 += (v6 & (v5 / ((v5) | 1)))
                            switch (fi54 >> ((18446744073709551615) & 63)) {
                            case 42:
                                var v121 bool = true
                                _ = v121
                            }
                            v10 += 136
                            if false {
                                v10 += 42
                            }
                            var v122 bool = true
                            v11 = (v11 * v11)
                            var v123 uint64 = uint64(v11)
                            var v124 uint8 = 128
                            if true {
                                v10 += (v10 + v10)
                                delete(v4, 65535)
                            }
                            _ = fi56
                            _ = v122
                            _ = v123
                            _ = v124
                        }
                    case -435:
                        for fi125 := 0; fi125 < 20; fi125++ {
                            if false {
                                v2 = 2
                                delete(v4, v10)
                                if (!false) {
                                    v4[312] = v10
                                }
                            }
                            _ = fi125
                        }
                    }
                    v126 := GenS9{F0: (-25 != 33), F1: "hello"}
                    for fi127 := 0; fi127 < 8; fi127++ {
                        var v128 uint16 = 1
                        _ = fi127
                        _ = v128
                    }
                    _ = v126
                }
            case 211:
                v2 = v2
            case 9223372036854775807:
                if true {
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                }
            }
        } else {
            v4[(v10 ^ v10)] = v10
            v4[(^(v10))] = v10
        }
        for fi129 := 0; fi129 < 18; fi129++ {
            v130 := GenS15{F0: ("\n\t" < ""), F1: -1, F2: ((!false) || true)}
            _ = fi129
            _ = v130
        }
        var v131 int16 = 1
        _ = fi54
        _ = v131
    }
    if false {
        if true {
            if false {
                if true {
                    for fi132 := 0; fi132 < 10; fi132++ {
                        var v133 uint8 = v11
                        var v134 uint64 = 1
                        _ = fi132
                        _ = v133
                        _ = v134
                    }
                }
            }
        }
    } else {
        var v135 uint64 = 18446744073709551615
        _ = v135
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
