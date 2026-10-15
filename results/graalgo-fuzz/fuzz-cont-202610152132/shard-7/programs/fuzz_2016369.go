package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint8; F2 string }
type GenS3 struct { F0 uint8; F1 int8; F2 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 int8; F2 uint8 }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 uint16; F1 bool; F2 int32 }
type GenS8 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS9 struct { F0 uint64; F1 int8 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 string; F1 uint }
type GenS12 struct { F0 int32; F1 int8 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int; F1 int64; F2 uint }
func genFn0(p0 uint, p1 uint8) int {
    var v4 uint8 = 206
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := []int{-509, (fi5 * fi5), fi5, (fi5 / ((fi5) | 1))}
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return 2147483647
}
func genFn1() int64 {
    var v7 uint = 65535
    v7 = (v7 & (v7 >> ((v7) & 63)))
    if (false || (0.1 >= 3.14159)) {
        var v8 uint16 = 943
        v9 := []int{831, 2147483647}
        if ((145 <= 2147483647) || ((!false) || (!true))) {
            switch "hello" {
            case "foo bar":
                var v10 uint32 = 4294967295
                _ = v10
            case "":
                for fi11 := 0; fi11 < 18; fi11++ {
                    if (true || (false || true)) {
                        var v12 uint = (v7 & 18446744073709551615)
                        var v13 uint64 = 234
                        var v14 float64 = (1e10 / float64((0.1 - (3.14159 + 0.0))))
                        _ = v12
                        _ = v13
                        _ = v14
                    } else {
                        fmt.Println(fi11)
                        v15 := GenS1{F0: (fi11 / ((-1) | 1))}
                        _ = v15
                    }
                    _ = fi11
                }
            }
            for fi16 := 0; fi16 < 9; fi16++ {
                for fi17 := 0; fi17 < 10; fi17++ {
                    fmt.Println(fi17, v8)
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    _ = fi17
                }
                var v18 uint16 = v8
                _ = fi16
                _ = v18
            }
            if false {
                fmt.Println(v8, v7)
                v7 += 4294967295
                switch 554 {
                case 42:
                    switch -2147483648 {
                    case 81:
                        if true {
                            v19 := GenS1{F0: 32767}
                            fmt.Println(v7)
                            v9 = append(v9, -1)
                            _ = v19
                        }
                    case 291:
                        v9 = append(v9, -9223372036854775808)
                    }
                case 0:
                    switch 9223372036854775807 {
                    case -9223372036854775808:
                        v9 = append(v9, -1)
                    case 42:
                        for fi20 := 0; fi20 < 5; fi20++ {
                            switch fi20 {
                            case 431:
                                fmt.Println(v8, v8)
                            case 2147483647:
                                fmt.Println(fi20)
                            case -2147483648:
                                var v21 string = "\n\t"
                                _ = v21
                            default:
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                            }
                            _ = fi20
                        }
                    default:
                        for fi22 := 0; fi22 < 7; fi22++ {
                            v7 = 195
                            switch "\n\t" {
                            case "":
                                v7 = 127
                            case "foo bar":
                                v8 -= (v8 % ((65535) | 1))
                            }
                            for fi23 := 0; fi23 < 12; fi23++ {
                                switch "a" {
                                case "\n\t":
                                    if true {
                                        v24 := GenS2{F0: "a", F1: 35, F2: "foo bar"}
                                        v7 -= v7
                                        v8 -= 2
                                        _ = v24
                                    }
                                case "foo bar":
                                    if false {
                                        v8 = (v8 - 627)
                                    }
                                default:
                                    if (!false) {
                                        var v25 float32 = (((1e10 * float32((1.5 * (3.14159 + (-2.25 + 1.5))))) * (-2.25 - 0.0)) * (0.1 - float32(0.1)))
                                        _ = v25
                                    } else {
                                        v26 := GenS3{F0: 127, F1: 100, F2: 1000}
                                        var v27 uint = (v7 - 541)
                                        if (4294967295 <= 849) {
                                            if (!false) {
                                                var v28 uint32 = 42
                                                _ = v28
                                            } else {
                                                if (false || (false && true)) {
                                                    if false {
                                                        v29 := GenS4{F0: "\n\t"}
                                                        var v30 uint32 = 0
                                                        _ = v29
                                                        _ = v30
                                                    } else {
                                                        fmt.Println(v27)
                                                    }
                                                    if (!true) {
                                                        var v31 uint = v27
                                                        v8 -= (v8 ^ (v8 & (v8 << ((v7) & 63))))
                                                        _ = v31
                                                    } else {
                                                        if len(v9) > 0 {
                                                            v9[0] = 42
                                                        }
                                                        v7 -= v27
                                                    }
                                                    v27 += v27
                                                } else {
                                                    if true {
                                                        v8 += (v8 / ((0) | 1))
                                                        var v32 int32 = 2147483647
                                                        switch genFn0(221, 127) {
                                                        case 4294967295:
                                                            if false {
                                                                v7 = (v27 << ((4294967295) & 63))
                                                                var v33 float32 = 1e10
                                                                _ = v33
                                                            } else {
                                                                switch 234 {
                                                                case -9223372036854775808:
                                                                    if (0 < 493) {
                                                                        v7 = (v7 ^ (^(v27)))
                                                                        if true {
                                                                            fmt.Println(fi22, fi23)
                                                                            if ((42 != (v32 & (v32 | (v32 ^ v32)))) && true) {
                                                                                var v34 int16 = -32768
                                                                                var v35 int32 = v32
                                                                                var v36 uint32 = 0
                                                                                _ = v34
                                                                                _ = v35
                                                                                _ = v36
                                                                            } else {
                                                                                v37 := GenS5{F0: 0, F1: 0, F2: 128}
                                                                                v7 = (v27 & (v27 & v7))
                                                                                var v38 int32 = v32
                                                                                _ = v37
                                                                                _ = v38
                                                                            }
                                                                        }
                                                                    }
                                                                case -961:
                                                                    _ = len(v9)
                                                                default:
                                                                    v39 := GenS6{F0: -391, F1: ((-1 <= v32) || (0 > 1))}
                                                                    _ = v39
                                                                }
                                                                v40 := GenS7{F0: (v8 | (v8 ^ (v8 + (v8 / (((v8 ^ 521)) | 1))))), F1: false, F2: (v32 + -1)}
                                                                _ = v40
                                                            }
                                                        }
                                                        v8 -= (v8 ^ v8)
                                                        v41 := GenS8{F0: 18446744073709551615, F1: v8, F2: 0}
                                                        v42 := GenS4{F0: "hello"}
                                                        if ("a" < "\n\t") {
                                                            if (!(550 < -769)) {
                                                                if false {
                                                                    if false {
                                                                        var v43 uint16 = 65535
                                                                        if false {
                                                                            switch ((("" + "") + "hello") + "hello") {
                                                                            case "hello":
                                                                                v7 -= v7
                                                                            case "a":
                                                                                fmt.Println(fi22, fi23)
                                                                            case "":
                                                                                v32 += v32
                                                                            default:
                                                                                switch ("foo bar" + "a") {
                                                                                case "hello":
                                                                                    var v44 float32 = float32(-2.25)
                                                                                    _ = v44
                                                                                case "foo bar":
                                                                                    if false {
                                                                                        v32 = (v32 | 0)
                                                                                        fmt.Println(fi23)
                                                                                    } else {
                                                                                        v32 = (v32 * v32)
                                                                                        var v45 int8 = 67
                                                                                        if ((0.0 != 0.0) && (!((v32 % ((v32) | 1)) != (v32 ^ (v32 >> ((42) & 63)))))) {
                                                                                            fmt.Println(fi22)
                                                                                        } else {
                                                                                            if (1000 <= -9223372036854775808) {
                                                                                                var v46 int8 = -34
                                                                                                v47 := GenS8{F0: 18446744073709551615, F1: (v8 / ((138) | 1)), F2: -1}
                                                                                                _ = v46
                                                                                                _ = v47
                                                                                            } else {
                                                                                                var v48 int32 = -1
                                                                                                v49 := GenS9{F0: 0, F1: -128}
                                                                                                _ = v48
                                                                                                _ = v49
                                                                                            }
                                                                                            switch fi23 {
                                                                                            case 305:
                                                                                                fmt.Println(v7)
                                                                                            case 2147483647:
                                                                                                switch "a" {
                                                                                                case "":
                                                                                                    v8 = v43
                                                                                                }
                                                                                            case 42:
                                                                                                v27 += (v27 | v27)
                                                                                            default:
                                                                                                var v50 int32 = (v32 << ((v7) & 63))
                                                                                                _ = v50
                                                                                            }
                                                                                        }
                                                                                        if len(v9) >= 2 {
                                                                                            v9 = v9[1:len(v9)]
                                                                                        }
                                                                                        _ = v45
                                                                                    }
                                                                                default:
                                                                                    v51 := GenS10{F0: 2147483648}
                                                                                    _ = v51
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v27 = v7
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                v27 = v7
                                                                            case "a":
                                                                                fmt.Println(v27, fi23, fi22)
                                                                            }
                                                                        }
                                                                        _ = v43
                                                                    }
                                                                    v52 := GenS8{F0: 42, F1: 1, F2: 42}
                                                                    _ = v52
                                                                } else {
                                                                    fmt.Println(v7, v8)
                                                                }
                                                            }
                                                            v27 -= v7
                                                        } else {
                                                            v32 -= (v32 << ((4294967295) & 63))
                                                        }
                                                        _ = v32
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                    v27 -= (v27 & (v27 % ((373) | 1)))
                                                    v8 = 994
                                                    switch (fi23 ^ fi22) {
                                                    case 9223372036854775807:
                                                        var v53 uint8 = 0
                                                        _ = v53
                                                    }
                                                }
                                                v7 -= 255
                                                var v54 uint32 = 2147483647
                                                _ = v54
                                            }
                                            switch fi23 {
                                            case 966:
                                                v55 := GenS10{F0: 493}
                                                _ = v55
                                            case -9223372036854775808:
                                                v56 := GenS9{F0: 709, F1: 127}
                                                _ = v56
                                            case 127:
                                                var v57 uint64 = 982
                                                _ = v57
                                            default:
                                                var v58 int = genFn0(100, 255)
                                                _ = v58
                                            }
                                            v59 := GenS1{F0: (fi22 >> ((v27) & 63))}
                                            _ = v59
                                        }
                                        _ = v26
                                        _ = v27
                                    }
                                }
                                _ = fi23
                            }
                            _ = fi22
                        }
                    }
                case 32768:
                    v60 := GenS11{F0: ("\n\t" + "hello"), F1: v7}
                    _ = v60
                default:
                    for fi61 := 0; fi61 < 9; fi61++ {
                        v7 += v7
                        switch "\n\t" {
                        case "":
                            v7 -= (v7 ^ v7)
                        }
                        _ = fi61
                    }
                }
            } else {
                if (!true) {
                    fmt.Println(v7)
                }
                if (!(v8 == (v8 ^ 0))) {
                    v62 := GenS5{F0: 1, F1: 0, F2: 128}
                    _ = v62
                }
            }
        } else {
            var v63 int16 = -32768
            v8 += v8
            var v64 bool = false
            _ = v63
            _ = v64
        }
        _ = v8
        _ = v9
    }
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 2147483647
    var v3 uint = 0
    for fi65 := 0; fi65 < 20; fi65++ {
        if true {
            for fi66 := 0; fi66 < 16; fi66++ {
                if true {
                    var v67 bool = (fi65 < 72)
                    if v67 {
                        v68 := GenS5{F0: 0, F1: 127, F2: 43}
                        _ = v68
                    } else {
                        if v67 {
                            v2 = (v2 ^ 42)
                            switch "a" {
                            case "\n\t":
                                var v69 int64 = 42
                                _ = v69
                            }
                        }
                        v70 := GenS11{F0: "a", F1: (^(v3))}
                        _ = v70
                    }
                    v3 -= 627
                    var v71 uint8 = 127
                    if (!(!(v67 && v67))) {
                        switch (fi66 >> (((v3 >> ((v3) & 63))) & 63)) {
                        case 203:
                            if ((v3 << ((65535) & 63)) == (v3 % ((v3) | 1))) {
                                fmt.Println(v2)
                                v3 = v3
                                if v67 {
                                    if v67 {
                                        var v72 int64 = -1
                                        if (!v67) {
                                            fmt.Println(v3, v2, fi65)
                                            v67 = (v72 > (-(v72)))
                                            if true {
                                                switch ("hello" + (("a" + ("hello" + ("hello" + "foo bar"))) + "")) {
                                                case "a":
                                                    var v73 bool = (false && (!("\n\t" >= "\n\t")))
                                                    _ = v73
                                                }
                                                v74 := []int{127, (^(fi66)), (fi65 | (fi65 % (((fi65 / (((fi65 << ((4294967295) & 63))) | 1))) | 1)))}
                                                var v75 uint = (v3 << (((^(v3))) & 63))
                                                _ = v74
                                                _ = v75
                                            }
                                        } else {
                                            var v76 int32 = 2147483647
                                            _ = v76
                                        }
                                        v77 := GenS6{F0: v72, F1: v67}
                                        _ = v72
                                        _ = v77
                                    } else {
                                        var v78 uint = (v3 ^ (v3 << (((^(v3))) & 63)))
                                        _ = v78
                                    }
                                } else {
                                    v79 := GenS11{F0: "foo bar", F1: 236}
                                    var v80 uint8 = (v71 + v71)
                                    _ = v79
                                    _ = v80
                                }
                                var v81 uint = 65535
                                switch "foo bar" {
                                case "foo bar":
                                    v82 := GenS12{F0: (-(v2)), F1: 42}
                                    _ = v82
                                }
                                _ = v81
                            } else {
                                v83 := []int{}
                                v2 = 1
                                v3 -= 319
                                _ = v83
                            }
                        }
                        if v67 {
                            v84 := []int{-199, (^(fi66)), 42}
                            var v85 int16 = -446
                            fmt.Println(fi65, v85, v3)
                            _ = v84
                            _ = v85
                        } else {
                            v86 := []int{(fi66 % ((fi65) | 1))}
                            _ = v86
                        }
                        switch (fi65 + genFn0((v3 * v3), 255)) {
                        case 741:
                            switch ("\n\t" + "a") {
                            case "":
                                if false {
                                    v87 := v1[42]
                                    _ = v87
                                } else {
                                    fmt.Println(v67)
                                    fmt.Println(fi65, v71)
                                }
                            default:
                                var v88 bool = (!(!false))
                                _ = v88
                            }
                        default:
                            var v89 int8 = 13
                            _ = v89
                        }
                        v67 = true
                    } else {
                        if false {
                            fmt.Println(v2)
                            var v90 float64 = 0.1
                            var v91 int8 = 42
                            _ = v90
                            _ = v91
                        } else {
                            switch ("" + ("\n\t" + "hello")) {
                            case "hello":
                                fmt.Println(fi65, v2, fi66)
                            }
                            v92 := []int{(fi66 / (((fi65 % ((9223372036854775807) | 1))) | 1))}
                            _ = v92
                        }
                    }
                    var v93 uint64 = 1
                    _ = v67
                    _ = v71
                    _ = v93
                } else {
                    v2 = (^(v2))
                }
                v3 -= 255
                _ = fi66
            }
        }
        _ = fi65
    }
    switch 127 {
    case 128:
        v2 += (v2 >> (((v3 << (((v3 >> ((v3) & 63))) & 63))) & 63))
    case 127:
        v2 = 2147483647
    case 9223372036854775807:
        delete(v1, 472)
    default:
        v94 := []int{42, 127, 0, -720}
        _ = v94
    }
    fmt.Println(v3)
    v2 = -725
    v95 := GenS9{F0: 4294967295, F1: 127}
    var v96 float64 = 0.1
    switch "a" {
    case "foo bar":
        if (genFn0((v3 - 18446744073709551615), 127) == 42) {
            v97 := GenS10{F0: 2147483648}
            if (true && true) {
                if false {
                    var v98 uint8 = 1
                    _ = v98
                }
            } else {
                if (true || (1e10 == ((1.5 * (1e10 / float32(v96))) * 0.0))) {
                    var v99 int64 = 1
                    v100 := GenS3{F0: 2, F1: int8(v2), F2: -1}
                    switch 127 {
                    case 740:
                        v1[573] = ""
                    }
                    _ = v100
                    _ = v99
                } else {
                    var v101 float32 = (1e10 * 1.5)
                    if (true || (!true)) {
                        v3 += (v3 >> ((v3) & 63))
                        switch (("\n\t" + "\n\t") + (((("hello" + ("a" + ("a" + "foo bar"))) + ("" + "a")) + "") + "foo bar")) {
                        case "\n\t":
                            var v102 int64 = genFn1()
                            _ = v102
                        case "":
                            switch -1 {
                            case 2147483647:
                                if true {
                                    var v103 uint64 = 321
                                    _ = v103
                                } else {
                                    fmt.Println(v96, v3)
                                    if true {
                                        var v104 uint8 = 88
                                        _ = v104
                                    }
                                }
                            case 865:
                                for fi105 := 0; fi105 < 4; fi105++ {
                                    var v106 uint8 = 128
                                    var v107 int32 = v2
                                    _ = fi105
                                    _ = v106
                                    _ = v107
                                }
                            }
                        }
                        v108 := GenS13{F0: (^(v2))}
                        _ = v108
                    }
                    _ = v101
                }
            }
            _ = v97
        } else {
            if ((!false) && ((153 == 18446744073709551615) && false)) {
                v1[0] = "a"
                v2 += (v2 & v2)
            }
            delete(v1, 326)
        }
    }
    switch 4294967295 {
    case 520:
        var v109 uint64 = 1
        _ = v109
    case 42:
        if true {
            var v110 uint8 = 42
            _ = v110
        }
    default:
        if (32767 >= 204) {
            v111 := []int{-198, 0, 127}
            v96 = 0.0
            for fi112 := 0; fi112 < 5; fi112++ {
                _ = len(v111)
                v3 = (^(v3))
                _ = fi112
            }
            _ = v111
        }
    }
    v113 := GenS14{F0: 4294967295, F1: 555, F2: (v3 ^ v3)}
    v114 := GenS2{F0: ("" + ""), F1: 1, F2: ("" + ("foo bar" + "hello"))}
    switch genFn0(v3, 30) {
    case 4294967295:
        for fi115 := 0; fi115 < 9; fi115++ {
            var v116 int = 2147483647
            var v117 float64 = v96
            _ = fi115
            _ = v116
            _ = v117
        }
    case 32767:
        fmt.Println(v96, v2, v2)
    }
    for fi118 := 0; fi118 < 20; fi118++ {
        v119 := GenS8{F0: 0, F1: 65535, F2: 1}
        v3 = v3
        _ = fi118
        _ = v119
    }
    fmt.Println(v96, v3)
    if (3 >= 2) {
        v96 = float64((((3.14159 - (float32(v96) * (3.14159 - float32((v96 + v96))))) + (float32(v96) - (-2.25 - 1e10))) - float32(v96)))
        if false {
            for fi120 := 0; fi120 < 18; fi120++ {
                v1[0] = ""
                _ = fi120
            }
        } else {
            var v121 int32 = 2147483647
            _ = v121
        }
    } else {
        v122 := []int{32768, 127, 0, 1}
        _ = v122
    }
    var v123 float32 = (((float32(float64(-2.25)) / -2.25) + ((1.5 / (0.0 + -2.25)) / (float32((((v96 / (v96 / v96)) * 3.14159) - 0.1)) - (1e10 + 1e10)))) + ((1e10 + 1e10) + (-2.25 / float32(v96))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v123:", v123)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v113
    _ = v114
    _ = v123
    _ = v2
    _ = v3
    _ = v95
    _ = v96
}
