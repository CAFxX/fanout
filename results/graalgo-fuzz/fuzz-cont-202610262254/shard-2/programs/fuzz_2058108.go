package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int32; F1 uint64 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int; F1 uint }
type GenS11 struct { F0 uint; F1 uint32; F2 int8 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 int8; F1 string; F2 uint8 }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 int32; F1 int64; F2 int8 }
type GenS18 struct { F0 bool; F1 int32; F2 uint32 }
type GenS19 struct { F0 uint8; F1 int; F2 int32 }
type GenS20 struct { F0 bool; F1 int16 }
type GenS21 struct { F0 uint }
type GenS22 struct { F0 int32; F1 bool }
type GenS23 struct { F0 int32 }
func genFn0(p0 int64, p1 int) bool {
    if false {
        var v5 int32 = -2147483648
        if (0 > 585) {
            v5 -= (^(v5))
        } else {
            for fi6 := 0; fi6 < 15; fi6++ {
                v7 := GenS1{F0: p0, F1: 875, F2: "\n\t"}
                fmt.Println(p0, p1)
                for fi8 := 0; fi8 < 11; fi8++ {
                    p1 = fi6
                    _ = fi8
                }
                _ = fi6
                _ = v7
            }
            var v9 uint = 634
            for fi10 := 0; fi10 < 19; fi10++ {
                v11 := GenS1{F0: (p0 << ((v9) & 63)), F1: 65535, F2: ""}
                v5 += (v5 << ((v9) & 63))
                fmt.Println(v5)
                _ = fi10
                _ = v11
            }
            _ = v9
        }
        _ = v5
    }
    fmt.Println(p1, p0)
    switch (^(p1)) {
    case -9223372036854775808:
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 int32 = 2147483647
            for fi14 := 0; fi14 < 9; fi14++ {
                v15 := []int{}
                v16 := GenS1{F0: (p0 & (p0 << ((2) & 63))), F1: 758, F2: "foo bar"}
                var v17 uint64 = 2
                v18 := GenS1{F0: 0, F1: 217, F2: ("a" + "a")}
                if ((true || (0.1 == ((1.5 + -2.25) / (((-2.25 - (3.14159 * -2.25)) + float64((0.1 + -2.25))) - 3.14159)))) || true) {
                    var v19 float64 = 1e10
                    if false {
                        switch (fi12 >> ((255) & 63)) {
                        case 2147483647:
                            v15 = append(v15, (fi12 + int(p0)))
                        case 0:
                            switch p1 {
                            case -2147483648:
                                var v20 uint32 = 1
                                _ = v20
                            case 9223372036854775807:
                                var v21 uint64 = (^(v17))
                                _ = v21
                            case 0:
                                if ((fi12 ^ fi14) == (fi12 & fi14)) {
                                    var v22 uint8 = 42
                                    switch "foo bar" {
                                    case "\n\t":
                                        v15 = append(v15, 2147483647)
                                        switch "" {
                                        case "foo bar":
                                            _ = len(v15)
                                        default:
                                            switch (fi12 & p1) {
                                            case -352:
                                                var v23 float32 = 1e10
                                                _ = v23
                                            case 4294967295:
                                                v22 = v22
                                            case 128:
                                                v17 = v17
                                            }
                                        }
                                        if (1e10 != -2.25) {
                                            v13 = 2147483647
                                            v15 = append(v15, (^(p1)))
                                            if ((false || true) && false) {
                                                var v24 float64 = 1e10
                                                fmt.Println(fi14)
                                                _ = v24
                                            } else {
                                                p0 = 550
                                                var v25 uint8 = (v22 >> ((18446744073709551615) & 63))
                                                fmt.Println(p0)
                                                _ = v25
                                            }
                                            var v26 float32 = 3.14159
                                            var v27 uint = 2
                                            _ = v26
                                            _ = v27
                                        } else {
                                            _ = len(v15)
                                            if (!((((v19 * 0.1) + ((3.14159 * ((((v19 + (((v19 - v19) / ((((1.5 + (3.14159 - float64(1e10))) + ((3.14159 + v19) / 1.5)) * (float64(1.5) - (((v19 - -2.25) + ((0.0 * 1.5) + (v19 + -2.25))) - -2.25))) * v19)) / v19)) / 0.1) / float64(-2.25)) - ((v19 - float64((1.5 * (1e10 - (float32(v19) / (float32(v19) - (0.0 * (((((1.5 - (0.0 - -2.25)) + -2.25) - 1e10) - 0.0) - (0.0 - (((1e10 * 1e10) / 0.1) + ((3.14159 - ((1e10 + -2.25) - 0.1)) + (((-2.25 * float32(v19)) + (float32(3.14159) * (1e10 / 3.14159))) * (-2.25 - 1e10))))))))))))) / v19))) + 1e10)) / v19) <= v19)) {
                                                v13 = v13
                                                v28 := GenS1{F0: (-(p0)), F1: 283, F2: "a"}
                                                _ = v28
                                            } else {
                                                v15 = append(v15, p1)
                                                if ((^(fi12)) < fi12) {
                                                    v29 := GenS2{F0: 9223372036854775807}
                                                    v30 := GenS2{F0: (p0 >> ((128) & 63))}
                                                    var v31 bool = false
                                                    _ = v29
                                                    _ = v30
                                                    _ = v31
                                                } else {
                                                    v32 := GenS2{F0: (p0 % (((p0 - (p0 * -838))) | 1))}
                                                    if ((0.1 * float64((1e10 + 0.1))) != (v19 * v19)) {
                                                        p0 = p0
                                                    } else {
                                                        switch (("\n\t" + "hello") + ("" + ("hello" + ("" + ("hello" + (("a" + ("\n\t" + ("a" + "a"))) + "foo bar")))))) {
                                                        case "a":
                                                            v33 := GenS1{F0: (p0 >> ((255) & 63)), F1: 1, F2: "foo bar"}
                                                            _ = v33
                                                        default:
                                                            switch "\n\t" {
                                                            case "hello":
                                                                if true {
                                                                    p1 += (fi14 | (fi12 ^ 359))
                                                                    v13 = v13
                                                                    var v34 int64 = (p0 - -9223372036854775808)
                                                                    _ = v34
                                                                }
                                                            case "\n\t":
                                                                switch "a" {
                                                                case "":
                                                                    v35 := GenS3{F0: 32767}
                                                                    _ = v35
                                                                }
                                                            }
                                                        }
                                                    }
                                                    v36 := GenS4{F0: 32768}
                                                    var v37 uint64 = v17
                                                    _ = v32
                                                    _ = v36
                                                    _ = v37
                                                }
                                                v19 = v19
                                            }
                                        }
                                        var v38 bool = (0 > 850)
                                        fmt.Println(p1)
                                        v39 := GenS4{F0: p1}
                                        _ = v38
                                        _ = v39
                                    case "foo bar":
                                        var v40 uint8 = 1
                                        _ = v40
                                    }
                                    _ = v22
                                }
                                if ((!false) || false) {
                                    if false {
                                        var v41 int16 = 0
                                        if true {
                                            v42 := GenS5{F0: v41, F1: 204}
                                            v43 := GenS6{F0: 332}
                                            v15 = append(v15, (fi14 >> ((18446744073709551615) & 63)))
                                            _ = v42
                                            _ = v43
                                        }
                                        if false {
                                            if false {
                                                switch (fi12 >> ((1) & 63)) {
                                                case 42:
                                                    v44 := GenS7{F0: (v41 * 0)}
                                                    _ = v44
                                                case 127:
                                                    switch "\n\t" {
                                                    case "a":
                                                        p1 = fi14
                                                    case "foo bar":
                                                        if len(v15) >= 2 {
                                                            v15 = v15[1:len(v15)]
                                                        }
                                                    case "":
                                                        var v45 int8 = 42
                                                        _ = v45
                                                    }
                                                default:
                                                    v46 := GenS4{F0: p1}
                                                    var v47 uint16 = 0
                                                    _ = v46
                                                    _ = v47
                                                }
                                                if len(v15) >= 2 {
                                                    v15 = v15[1:len(v15)]
                                                }
                                                var v48 string = "foo bar"
                                                v49 := GenS8{F0: v13, F1: 389}
                                                v50 := GenS2{F0: 9223372036854775807}
                                                v19 -= float64((1.5 + (1.5 / -2.25)))
                                                _ = v48
                                                _ = v49
                                                _ = v50
                                            } else {
                                                if (false || true) {
                                                    var v51 float64 = v19
                                                    var v52 float64 = v51
                                                    var v53 uint16 = 411
                                                    _ = v51
                                                    _ = v52
                                                    _ = v53
                                                }
                                                v19 = (v19 - (v19 * v19))
                                                if len(v15) > 0 {
                                                    v15[0] = 32768
                                                }
                                            }
                                            fmt.Println(p1)
                                            if (true || true) {
                                                v54 := GenS9{F0: 42}
                                                _ = v54
                                            } else {
                                                v17 = 1
                                                if false {
                                                    if false {
                                                        v13 -= 2147483647
                                                        var v55 int32 = v13
                                                        _ = v55
                                                    } else {
                                                        v56 := GenS4{F0: (p1 | (p1 >> ((1) & 63)))}
                                                        _ = v56
                                                    }
                                                    if true {
                                                        var v57 int64 = (p0 & 0)
                                                        var v58 int8 = -112
                                                        p0 += (p0 % ((v57) | 1))
                                                        _ = v57
                                                        _ = v58
                                                    } else {
                                                        switch fi12 {
                                                        case 4294967295:
                                                            switch fi12 {
                                                            case 42:
                                                                var v59 float64 = 0.1
                                                                _ = v59
                                                            case 9223372036854775807:
                                                                p0 = 42
                                                            case -44:
                                                                if len(v15) >= 2 {
                                                                    v15 = v15[1:len(v15)]
                                                                }
                                                            default:
                                                                _ = len(v15)
                                                            }
                                                        }
                                                        if (!((v41 + (v41 * v41)) > 32767)) {
                                                            p1 -= (p1 | (-(fi12)))
                                                            switch fi12 {
                                                            case 2147483647:
                                                                v19 -= 3.14159
                                                            default:
                                                                v60 := GenS1{F0: -9223372036854775808, F1: 0, F2: (("hello" + (("" + ("" + ("" + "foo bar"))) + ("\n\t" + "\n\t"))) + (("hello" + "a") + "hello"))}
                                                                _ = v60
                                                            }
                                                        }
                                                        var v61 string = ("\n\t" + "a")
                                                        _ = v61
                                                    }
                                                    var v62 uint32 = 554
                                                    _ = v62
                                                } else {
                                                    var v63 int64 = (p0 ^ -268)
                                                    _ = v63
                                                }
                                            }
                                            if ((!(!true)) || (false && (!(!((v41 + -966) < 1000))))) {
                                                if false {
                                                    switch ((((("\n\t" + "hello") + "\n\t") + "foo bar") + "") + "hello") {
                                                    case "\n\t":
                                                        v64 := GenS5{F0: v41, F1: 0}
                                                        _ = v64
                                                    }
                                                    if true {
                                                        fmt.Println(p1)
                                                        if (255 != 42) {
                                                            v41 -= v41
                                                        }
                                                    }
                                                }
                                            }
                                            p1 = fi12
                                            v17 -= 380
                                        } else {
                                            switch "a" {
                                            case "":
                                                switch 788 {
                                                case 32768:
                                                    v13 = (v13 & (v13 % (((v13 << ((2) & 63))) | 1)))
                                                case -858:
                                                    v17 -= 4294967295
                                                case 349:
                                                    v41 = v41
                                                default:
                                                    if true {
                                                        switch 2147483647 {
                                                        case 2147483647:
                                                            p1 += fi14
                                                        case 548:
                                                            v17 += (v17 >> ((42) & 63))
                                                        }
                                                        if true {
                                                            fmt.Println(p0)
                                                            if false {
                                                                v65 := GenS10{F0: fi14, F1: 853}
                                                                v66 := GenS11{F0: 223, F1: 2147483648, F2: 42}
                                                                _ = v65
                                                                _ = v66
                                                            }
                                                            v67 := GenS12{F0: int8(p1)}
                                                            fmt.Println(v13, v13)
                                                            _ = v67
                                                        } else {
                                                            if len(v15) > 0 {
                                                                v15[0] = (p1 % (((fi12 << ((4294967295) & 63))) | 1))
                                                            }
                                                            v19 = v19
                                                        }
                                                    } else {
                                                        _ = len(v15)
                                                        if ((v41 << ((255) & 63)) != v41) {
                                                            var v68 int = 527
                                                            _ = v68
                                                        }
                                                    }
                                                }
                                                fmt.Println(p1, v19)
                                                fmt.Println(v19, p1, v17)
                                            case "a":
                                                p1 = p1
                                            case "foo bar":
                                                switch (p1 ^ fi12) {
                                                case 2147483647:
                                                    v69 := GenS4{F0: fi12}
                                                    _ = v69
                                                }
                                                if len(v15) > 0 {
                                                    v15[0] = -1
                                                }
                                                var v70 float64 = (v19 + float64((0.1 + (0.0 * 1.5))))
                                                fmt.Println(p0, v17, p0)
                                                _ = v70
                                            default:
                                                p1 = p1
                                            }
                                            v41 += 1
                                            fmt.Println(p0)
                                        }
                                        var v71 uint32 = 2
                                        _ = v41
                                        _ = v71
                                    } else {
                                        v72 := GenS9{F0: 387}
                                        if (false && ("\n\t" <= "foo bar")) {
                                            var v73 int8 = 42
                                            fmt.Println(p0, p1)
                                            fmt.Println(p1, p0)
                                            if (616 != 507) {
                                                var v74 uint16 = 65535
                                                if ((!(!true)) || (-1 != (p0 * (p0 & (p0 + (p0 * 42)))))) {
                                                    switch (("" + "\n\t") + "hello") {
                                                    case "foo bar":
                                                        if len(v15) >= 2 {
                                                            v15 = v15[1:len(v15)]
                                                        }
                                                    case "\n\t":
                                                        var v75 int = int(v13)
                                                        _ = v75
                                                    case "a":
                                                        v15 = append(v15, fi12)
                                                        var v76 uint64 = 2
                                                        _ = v76
                                                    }
                                                } else {
                                                    if len(v15) >= 2 {
                                                        v15 = v15[1:len(v15)]
                                                    }
                                                    switch 1 {
                                                    case -333:
                                                        v77 := GenS12{F0: -128}
                                                        _ = v77
                                                    case 1:
                                                        if false {
                                                            if false {
                                                                v78 := GenS5{F0: 32767, F1: 255}
                                                                var v79 uint8 = 28
                                                                fmt.Println(fi12, v74)
                                                                switch -914 {
                                                                case -152:
                                                                    if len(v15) > 0 {
                                                                        v15[0] = (fi14 ^ (p1 & 470))
                                                                    }
                                                                }
                                                                v17 -= uint64(v79)
                                                                var v80 int8 = v73
                                                                _ = v78
                                                                _ = v79
                                                                _ = v80
                                                            } else {
                                                                _ = len(v15)
                                                            }
                                                        } else {
                                                            p1 = 2147483647
                                                            switch (p1 + (p1 << ((uint(v17)) & 63))) {
                                                            case 128:
                                                                if true {
                                                                    v81 := GenS13{F0: 188, F1: 69, F2: 4294967295}
                                                                    var v82 int16 = -32768
                                                                    _ = v81
                                                                    _ = v82
                                                                }
                                                            default:
                                                                v17 -= (v17 ^ (v17 << ((0) & 63)))
                                                            }
                                                            if len(v15) >= 2 {
                                                                v15 = v15[1:len(v15)]
                                                            }
                                                        }
                                                    case 2147483647:
                                                        p0 += (p0 << ((4294967295) & 63))
                                                    }
                                                    v83 := GenS9{F0: 2}
                                                    _ = v83
                                                }
                                                if (-9223372036854775808 < p1) {
                                                    v84 := GenS12{F0: (v73 & (v73 / (((v73 / ((v73) | 1))) | 1)))}
                                                    _ = v84
                                                } else {
                                                    if true {
                                                        v17 = (v17 | v17)
                                                        v85 := GenS14{F0: 4294967295}
                                                        fmt.Println(v13, p1)
                                                        _ = v85
                                                    } else {
                                                        switch 47 {
                                                        case 584:
                                                            v86 := GenS14{F0: v17}
                                                            _ = v86
                                                        case 900:
                                                            v87 := GenS15{F0: 127, F1: (("a" + "hello") + "\n\t"), F2: 127}
                                                            _ = v87
                                                        default:
                                                            if len(v15) >= 2 {
                                                                v15 = v15[1:len(v15)]
                                                            }
                                                        }
                                                        switch (fi12 - p1) {
                                                        case -968:
                                                            v88 := GenS5{F0: 1000, F1: 42}
                                                            _ = v88
                                                        }
                                                        var v89 uint32 = 42
                                                        _ = v89
                                                    }
                                                    switch fi14 {
                                                    case -580:
                                                        v13 += (v13 | v13)
                                                    default:
                                                        p0 = (p0 * p0)
                                                    }
                                                    v90 := GenS9{F0: 1}
                                                    _ = v90
                                                }
                                                var v91 uint = 18446744073709551615
                                                var v92 float64 = 0.1
                                                var v93 uint8 = 122
                                                _ = v74
                                                _ = v91
                                                _ = v92
                                                _ = v93
                                            }
                                            v94 := GenS16{F0: 1}
                                            _ = v73
                                            _ = v94
                                        } else {
                                            v17 = (^(v17))
                                        }
                                        v19 = float64((1e10 * -2.25))
                                        v95 := GenS17{F0: (v13 - (-(v13))), F1: -814, F2: 87}
                                        _ = v72
                                        _ = v95
                                    }
                                    switch (fi12 ^ int(p0)) {
                                    case -9223372036854775808:
                                        if (false || false) {
                                            v13 = 42
                                            if false {
                                                if (true || false) {
                                                    if false {
                                                        v15 = append(v15, (fi12 % (((fi14 % ((fi12) | 1))) | 1)))
                                                    }
                                                    var v96 int64 = p0
                                                    var v97 string = ("a" + "foo bar")
                                                    _ = v96
                                                    _ = v97
                                                }
                                                if true {
                                                    var v98 uint64 = v17
                                                    v15 = append(v15, (-(fi14)))
                                                    var v99 uint8 = 255
                                                    _ = v98
                                                    _ = v99
                                                } else {
                                                    v100 := GenS9{F0: 255}
                                                    fmt.Println(v19, p0)
                                                    _ = v100
                                                }
                                            } else {
                                                fmt.Println(p0)
                                                var v101 uint16 = 334
                                                v101 = v101
                                                v13 -= 549
                                                switch "a" {
                                                case "hello":
                                                    p0 = 787
                                                case "\n\t":
                                                    _ = len(v15)
                                                case "":
                                                    if len(v15) > 0 {
                                                        v15[0] = fi14
                                                    }
                                                }
                                                _ = v101
                                            }
                                            if (v19 <= (1e10 + v19)) {
                                                v19 = -2.25
                                            } else {
                                                if (!(false || true)) {
                                                    v17 -= 42
                                                    var v102 uint32 = 638
                                                    switch (fi14 >> ((65535) & 63)) {
                                                    case -306:
                                                        var v103 uint64 = 18446744073709551615
                                                        _ = v103
                                                    case -9223372036854775808:
                                                        if false {
                                                            var v104 uint32 = 4294967295
                                                            v102 -= (v104 % ((v104) | 1))
                                                            var v105 int64 = (p0 << ((18446744073709551615) & 63))
                                                            _ = v104
                                                            _ = v105
                                                        } else {
                                                            _ = len(v15)
                                                        }
                                                    }
                                                    var v106 int32 = (^(v13))
                                                    _ = v102
                                                    _ = v106
                                                } else {
                                                    var v107 string = "a"
                                                    if true {
                                                        switch "foo bar" {
                                                        case "a":
                                                            var v108 int8 = -128
                                                            _ = v108
                                                        case "\n\t":
                                                            var v109 int32 = (v13 / ((v13) | 1))
                                                            _ = v109
                                                        }
                                                    }
                                                    _ = v107
                                                }
                                                if true {
                                                    v110 := GenS9{F0: 255}
                                                    _ = v110
                                                }
                                            }
                                        } else {
                                            p1 -= (fi14 * (fi12 >> ((2) & 63)))
                                            fmt.Println(p1)
                                        }
                                    default:
                                        fmt.Println(v19, v13, fi12)
                                    }
                                    if (float32((float64(float32(v19)) + 1.5)) >= ((-2.25 + (0.0 / float32(float64(float32(v19))))) / ((-2.25 * (float32(v19) - (3.14159 * 0.0))) * 3.14159))) {
                                        var v111 float32 = -2.25
                                        v112 := GenS18{F0: false, F1: v13, F2: 2147483647}
                                        _ = v111
                                        _ = v112
                                    } else {
                                        if len(v15) > 0 {
                                            v15[0] = p1
                                        }
                                        var v113 uint8 = 14
                                        fmt.Println(v17, v113)
                                        _ = v113
                                    }
                                } else {
                                    v17 -= (v17 | v17)
                                }
                            default:
                                var v114 float32 = ((1.5 + ((float32(float64(0.1)) * float32((v19 - (v19 + ((v19 / -2.25) - (0.1 * v19)))))) - (1.5 / 1e10))) + 1.5)
                                _ = v114
                            }
                        default:
                            var v115 int32 = (v13 & -896)
                            _ = v115
                        }
                        v15 = append(v15, (^(p1)))
                        if ((9223372036854775807 == p0) || true) {
                            v116 := GenS14{F0: (v17 << ((4294967295) & 63))}
                            switch int(v17) {
                            case 32767:
                                switch (fi12 >> ((42) & 63)) {
                                case -658:
                                    p0 = (p0 & 832)
                                    v15 = append(v15, -13)
                                case -1:
                                    if len(v15) > 0 {
                                        v15[0] = fi12
                                    }
                                case 9223372036854775807:
                                    if (-579 > 0) {
                                        var v117 uint64 = 18446744073709551615
                                        if len(v15) > 0 {
                                            v15[0] = (fi12 + 112)
                                        }
                                        v13 -= (-(v13))
                                        _ = v117
                                    }
                                default:
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                }
                            case -40:
                                v13 = 42
                            case 128:
                                switch (^(fi12)) {
                                case 128:
                                    v118 := GenS19{F0: 255, F1: fi12, F2: v13}
                                    _ = v118
                                case 42:
                                    if false {
                                        v19 += ((v19 * v19) - 3.14159)
                                        var v119 float32 = 0.0
                                        _ = v119
                                    }
                                case 32768:
                                    var v120 int64 = 1
                                    _ = v120
                                }
                            default:
                                var v121 int16 = -980
                                _ = v121
                            }
                            _ = v116
                        }
                    }
                    var v122 int64 = p0
                    _ = v122
                    _ = v19
                } else {
                    v123 := GenS6{F0: 894}
                    switch 0 {
                    case 32767:
                        if (-32768 < 0) {
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                            var v124 float32 = 1e10
                            if false {
                                fmt.Println(p1, v13, p1)
                                v15 = append(v15, fi14)
                                if false {
                                    var v125 uint = 2
                                    var v126 uint64 = v17
                                    v127 := GenS20{F0: (2 < (v125 * (v125 / ((v125) | 1)))), F1: 1}
                                    _ = v125
                                    _ = v126
                                    _ = v127
                                }
                            }
                            var v128 float32 = v124
                            _ = v124
                            _ = v128
                        }
                    default:
                        fmt.Println(fi12, p0)
                    }
                    _ = v123
                }
                var v129 float32 = -2.25
                _ = fi14
                _ = v129
                _ = v15
                _ = v16
                _ = v17
                _ = v18
            }
            _ = fi12
            _ = v13
        }
    case 32768:
        p1 += -877
    case -627:
        p1 = p1
    default:
        for fi130 := 0; fi130 < 6; fi130++ {
            v131 := []int{}
            switch -587 {
            case 1:
                fmt.Println(fi130)
            default:
                switch "" {
                case "a":
                    p1 = (fi130 - p1)
                case "":
                    p0 = (^(p0))
                }
            }
            _ = fi130
            _ = v131
        }
    }
    p1 -= (-(p1))
    _ = p0
    _ = p1
    return ((p0 & (p0 << ((13) & 63))) > p0)
}
func genFn1() bool {
    var v132 float32 = 0.1
    v132 -= v132
    if ((2147483648 > 479) || false) {
        v132 += v132
        for fi133 := 0; fi133 < 16; fi133++ {
            switch "a" {
            case "a":
                var v134 float32 = v132
                _ = v134
            default:
                v132 = 0.1
            }
            _ = fi133
        }
    } else {
        v135 := GenS21{F0: 0}
        v132 += float32(((1e10 * (float64(v132) * -2.25)) * -2.25))
        _ = v135
    }
    v136 := GenS10{F0: 1, F1: 4294967295}
    _ = v132
    _ = v136
    return false
}

func main() {
    var v1 uint8 = 0
    var v2 string = "hello"
    var v3 string = v2
    var v4 uint64 = 1
    v137 := GenS8{F0: -1, F1: 414}
    v138 := GenS22{F0: 42, F1: (!false)}
    v139 := GenS23{F0: 918}
    var v140 int32 = 658
    for fi141 := 0; fi141 < 5; fi141++ {
        v142 := GenS8{F0: (v140 >> ((255) & 63)), F1: (v4 >> ((uint(v4)) & 63))}
        if false {
            for fi143 := 0; fi143 < 8; fi143++ {
                var v144 int16 = 0
                switch fi143 {
                case -9223372036854775808:
                    v3 = "\n\t"
                }
                v140 -= v140
                _ = fi143
                _ = v144
            }
        }
        _ = fi141
        _ = v142
    }
    var v145 uint16 = 895
    v146 := []int{int(v145), 1, 48}
    var v147 uint32 = 2
    switch -2147483648 {
    case 766:
        var v148 uint32 = uint32(v145)
        _ = v148
    default:
        if len(v146) > 0 {
            v146[0] = -678
        }
    }
    var v149 uint = 267
    var v150 uint8 = v1
    for fi151 := 0; fi151 < 5; fi151++ {
        for fi152 := 0; fi152 < 3; fi152++ {
            v146 = append(v146, fi152)
            if true {
                if false {
                    v140 -= (v140 % (((v140 + int32(v149))) | 1))
                }
            } else {
                fmt.Println(v150, v145)
                if len(v146) >= 2 {
                    v146 = v146[1:len(v146)]
                }
            }
            var v153 int32 = (v140 ^ (v140 - v140))
            v3 = "foo bar"
            v2 = ((("" + (v3 + "")) + ("a" + v3)) + v2)
            var v154 uint16 = (v145 * v145)
            _ = fi152
            _ = v153
            _ = v154
        }
        _ = len(v146)
        for fi155 := 0; fi155 < 1; fi155++ {
            var v156 uint32 = uint32(v1)
            _ = fi155
            _ = v156
        }
        _ = fi151
    }
    _ = len(v146)
    _ = len(v146)
    fmt.Println(v2, v3, v145)
    fmt.Println("v1:", v1)
    fmt.Println("v140:", v140)
    fmt.Println("v145:", v145)
    fmt.Println("slicelen v146:", len(v146), v146)
    fmt.Println("v147:", v147)
    fmt.Println("v149:", v149)
    fmt.Println("v150:", v150)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v137
    _ = v138
    _ = v139
    _ = v140
    _ = v145
    _ = v146
    _ = v147
    _ = v149
    _ = v150
    _ = v2
    _ = v3
    _ = v4
}
