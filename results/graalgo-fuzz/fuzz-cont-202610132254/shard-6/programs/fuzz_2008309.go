package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint8; F1 int64; F2 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 uint; F1 int; F2 int16 }
type GenS6 struct { F0 int; F1 int; F2 int }
type GenS7 struct { F0 string; F1 uint16 }
type GenS8 struct { F0 string; F1 string; F2 uint64 }
type GenS9 struct { F0 uint8; F1 bool }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint32; F1 int; F2 int8 }

func main() {
    var v1 int16 = 1
    var v2 bool = false
    var v3 float64 = float64(0.0)
    v4 := make(map[uint32]string)
    v5 := make(map[int]int)
    var v6 int64 = -522
    for fi7 := 0; fi7 < 15; fi7++ {
        fmt.Println(v2, v1)
        _ = fi7
    }
    v8 := GenS1{F0: 124, F1: 4294967295}
    if true {
        fmt.Println(v2)
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 bool = (("a" + ("a" + ("foo bar" + "foo bar"))) < "hello")
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    if (!((2147483647 <= 667) && true)) {
        v12 := v5[-2147483648]
        _ = v12
    } else {
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 int32 = 2147483647
            v15 := GenS2{F0: 128, F1: (-(v6)), F2: ""}
            v16 := GenS2{F0: 255, F1: v6, F2: ("" + "a")}
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        if (v1 == (v1 & (v1 % (((v1 / (((^(v1))) | 1))) | 1)))) {
            fmt.Println(v3)
            switch "\n\t" {
            case "":
                if (v2 && true) {
                    if (!v2) {
                        var v17 string = "hello"
                        if v2 {
                            var v18 int64 = v6
                            var v19 uint8 = 2
                            _ = v18
                            _ = v19
                        }
                        _ = v17
                    }
                    switch int(v6) {
                    case 1:
                        v20 := GenS3{F0: 4294967295}
                        _ = v20
                    case -9223372036854775808:
                        v2 = false
                    }
                }
            }
            switch "\n\t" {
            case "hello":
                v6 += v6
            case "\n\t":
                delete(v4, 4294967295)
            default:
                if true {
                    for fi21 := 0; fi21 < 3; fi21++ {
                        fmt.Println(v6)
                        var v22 float32 = 0.0
                        _ = fi21
                        _ = v22
                    }
                } else {
                    v23 := GenS4{F0: -584, F1: (("\n\t" + "\n\t") + "\n\t")}
                    fmt.Println(v1)
                    switch 742 {
                    case 804:
                        v24 := []int{-630, 32768, 1, 1}
                        _ = v24
                    }
                    _ = v23
                }
            }
        } else {
            for fi25 := 0; fi25 < 7; fi25++ {
                var v26 int64 = (v6 << ((4294967295) & 63))
                for fi27 := 0; fi27 < 3; fi27++ {
                    v28 := []int{fi27, -2147483648}
                    if true {
                        v29 := GenS1{F0: 0, F1: 4294967295}
                        v30 := v4[27]
                        v26 += v6
                        _ = v29
                        _ = v30
                    } else {
                        var v31 int = fi27
                        var v32 bool = (!false)
                        var v33 uint8 = 0
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    _ = fi27
                    _ = v28
                }
                _ = fi25
                _ = v26
            }
        }
    }
    v34 := GenS4{F0: -9223372036854775808, F1: (("foo bar" + ("\n\t" + "foo bar")) + ("foo bar" + ""))}
    for fi35 := 0; fi35 < 12; fi35++ {
        for fi36 := 0; fi36 < 2; fi36++ {
            v2 = (!(141 < 1))
            _ = fi36
        }
        _ = fi35
    }
    v37 := v5[4294967295]
    v38 := GenS5{F0: 255, F1: v37, F2: (v1 | 32767)}
    if true {
        v1 += v1
        var v39 uint64 = 4294967295
        _ = v39
    } else {
        for fi40 := 0; fi40 < 11; fi40++ {
            v6 = -337
            if v2 {
                var v41 int = (fi40 / (((fi40 >> ((65535) & 63))) | 1))
                if false {
                    switch fi40 {
                    case 316:
                        v42 := []int{}
                        _ = v42
                    case 1:
                        fmt.Println(v3)
                    case -9223372036854775808:
                        v41 = 0
                    }
                    v3 = v3
                } else {
                    var v43 uint8 = 255
                    for fi44 := 0; fi44 < 6; fi44++ {
                        if (v2 || ((4294967295 == 0) && (100 > -87))) {
                            v41 -= 32767
                            if (558 > 0) {
                                var v45 int = fi40
                                _ = v45
                            } else {
                                var v46 bool = v2
                                var v47 float64 = 1e10
                                v5[fi44] = fi40
                                _ = v46
                                _ = v47
                            }
                            v48 := GenS4{F0: 42, F1: ""}
                            _ = v48
                        }
                        if true {
                            if v2 {
                                var v49 int64 = v6
                                v1 -= (v1 - (v1 + v1))
                                _ = v49
                            } else {
                                v50 := GenS6{F0: (^(v37)), F1: 363, F2: (v41 >> ((215) & 63))}
                                _ = v50
                            }
                            v1 = v1
                            if ((((0.1 / 1.5) * float32(-2.25)) * (0.0 * -2.25)) < 0.1) {
                                v43 -= 1
                                if ((42 == 884) || v2) {
                                    v2 = true
                                    v5[v37] = (v41 / ((fi40) | 1))
                                }
                                if (-1 <= -1) {
                                    v51 := []int{}
                                    v2 = true
                                    v43 -= v43
                                    _ = v51
                                } else {
                                    v52 := []int{(v41 * (-(fi40))), 1, -9223372036854775808, -894}
                                    if (4294967295 == 4294967295) {
                                        var v53 int16 = v1
                                        if false {
                                            v41 -= v37
                                        } else {
                                            if (230 <= 42) {
                                                var v54 int32 = -951
                                                switch (v41 - fi44) {
                                                case 2147483647:
                                                    v55 := GenS4{F0: (fi44 * (fi44 / ((fi44) | 1))), F1: ""}
                                                    _ = v55
                                                case -410:
                                                    v53 += v53
                                                case -2147483648:
                                                    v2 = (3.14159 > -2.25)
                                                default:
                                                    v56 := GenS7{F0: ("a" + "hello"), F1: 2}
                                                    _ = v56
                                                }
                                                _ = v54
                                            }
                                        }
                                        _ = v53
                                    } else {
                                        v57 := GenS7{F0: "a", F1: 764}
                                        _ = v57
                                    }
                                    _ = v52
                                }
                            }
                            if false {
                                switch 2147483647 {
                                case -2147483648:
                                    fmt.Println(v6)
                                case 2147483647:
                                    switch "a" {
                                    case "a":
                                        var v58 float64 = v3
                                        _ = v58
                                    case "\n\t":
                                        fmt.Println(v43, v6, v6)
                                    }
                                default:
                                    if (!v2) {
                                        fmt.Println(v2, v6, v6)
                                    } else {
                                        var v59 uint8 = v43
                                        if (2 < 333) {
                                            var v60 int = 2147483647
                                            _ = v60
                                        } else {
                                            var v61 float64 = 0.0
                                            _ = v61
                                        }
                                        _ = v59
                                    }
                                }
                                var v62 string = ("a" + "\n\t")
                                v63 := []int{}
                                if (v2 || false) {
                                    if v2 {
                                        switch (fi44 >> ((53) & 63)) {
                                        case 0:
                                            var v64 uint = 1
                                            _ = v64
                                        case 127:
                                            if true {
                                                if len(v63) > 0 {
                                                    v63[0] = v37
                                                }
                                            } else {
                                                var v65 float32 = 0.0
                                                var v66 bool = ((v6 & 0) == 0)
                                                _ = v65
                                                _ = v66
                                            }
                                        }
                                    }
                                    var v67 uint8 = v43
                                    var v68 string = (v62 + "\n\t")
                                    _ = v67
                                    _ = v68
                                }
                                switch int(v6) {
                                case 32767:
                                    v69 := GenS4{F0: -804, F1: v62}
                                    _ = v69
                                case 128:
                                    var v70 int64 = v6
                                    _ = v70
                                case 1:
                                    v62 = v62
                                }
                                _ = v62
                                _ = v63
                            } else {
                                v5[(v37 ^ (fi40 >> ((42) & 63)))] = (v41 | v41)
                            }
                            if (v2 && (2147483647 != 500)) {
                                fmt.Println(v6, v3)
                                var v71 string = ""
                                fmt.Println(v37, v3)
                                _ = v71
                            }
                        }
                        _ = fi44
                    }
                    _ = v43
                }
                _ = v41
            } else {
                var v72 uint = 42
                for fi73 := 0; fi73 < 15; fi73++ {
                    var v74 int8 = 0
                    v2 = (true && (!true))
                    switch fi73 {
                    case 128:
                        var v75 int16 = (v1 % ((-98) | 1))
                        _ = v75
                    case 32767:
                        if (255 > 42) {
                            _ = len(v4)
                            fmt.Println(v3, v37)
                        } else {
                            switch v37 {
                            case -128:
                                var v76 int = (fi73 >> (((^(v72))) & 63))
                                _ = v76
                            case 2147483647:
                                var v77 uint = (v72 | 65535)
                                _ = v77
                            default:
                                switch 198 {
                                case 42:
                                    var v78 uint8 = 2
                                    _ = v78
                                case 32768:
                                    v37 = 375
                                case 127:
                                    v79 := []int{(-(fi40)), 1, fi73}
                                    _ = v79
                                }
                            }
                            if ((!((2 >= 2) && v2)) && (!v2)) {
                                v1 = (^(v1))
                                v80 := []int{526}
                                _ = v80
                            } else {
                                var v81 bool = (!(-279 < -1))
                                _ = v81
                            }
                            var v82 int16 = (^(v1))
                            _ = v82
                        }
                        switch int(v1) {
                        case -1:
                            v74 += 0
                        }
                    case 9223372036854775807:
                        fmt.Println(v1)
                    default:
                        if (((225 > 1) || v2) || (v72 > (v72 << ((v72) & 63)))) {
                            v83 := GenS1{F0: (-(v74)), F1: 401}
                            switch fi73 {
                            case 1:
                                if (v2 || ((v74 << ((255) & 63)) >= (v74 ^ (v74 % (((v74 / (((v74 >> ((v72) & 63))) | 1))) | 1))))) {
                                    if false {
                                        v84 := GenS4{F0: (fi73 / ((v37) | 1)), F1: "a"}
                                        v3 = (((((float64(1.5) * v3) * ((float64(3.14159) - ((float64(-2.25) / 0.1) + ((float64((float32(((v3 + ((((((((v3 / 1.5) / 1.0) - v3) * (3.14159 / ((3.14159 / (float64(3.14159) * (1.5 / v3))) * (v3 + 1.5)))) / v3) - v3) * 1.5) + v3)) / 1.0)) * ((1e10 - (0.1 + 0.1)) * 3.14159))) * ((1e10 * -2.25) - (3.14159 * (-2.25 + (v3 / 3.14159))))) / v3))) * 1e10)) + v3) + ((1.5 / (0.1 * v3)) * v3)) - 0.1)
                                        if v2 {
                                            if false {
                                                v6 = (-(v6))
                                                switch "hello" {
                                                case "":
                                                    v2 = (255 == v72)
                                                case "\n\t":
                                                    if (0 < 384) {
                                                        var v85 string = ("hello" + ((("a" + "foo bar") + "a") + ("hello" + "\n\t")))
                                                        switch (fi73 - fi73) {
                                                        case 32768:
                                                            var v86 int = 631
                                                            _ = v86
                                                        default:
                                                            switch "hello" {
                                                            case "hello":
                                                                delete(v4, 1)
                                                            case "":
                                                                v37 = fi40
                                                            }
                                                        }
                                                        _ = v85
                                                    } else {
                                                        var v87 uint16 = 387
                                                        _ = v87
                                                    }
                                                case "hello":
                                                    fmt.Println(v72, v3, v6)
                                                }
                                                var v88 uint8 = 1
                                                _ = v88
                                            } else {
                                                v89 := []int{(fi73 * fi73)}
                                                v2 = false
                                                _ = v89
                                            }
                                            var v90 uint32 = 1
                                            v91 := GenS3{F0: 18446744073709551615}
                                            _ = v90
                                            _ = v91
                                        }
                                        _ = v84
                                    } else {
                                        _ = len(v5)
                                        var v92 int8 = v74
                                        switch (fi73 | 238) {
                                        case -9223372036854775808:
                                            switch ((("hello" + ((("foo bar" + "\n\t") + "hello") + "")) + ("hello" + "")) + "\n\t") {
                                            case "hello":
                                                v93 := GenS7{F0: "hello", F1: 0}
                                                _ = v93
                                            default:
                                                v94 := GenS8{F0: "", F1: ("\n\t" + ("hello" + "foo bar")), F2: 993}
                                                _ = v94
                                            }
                                        case -203:
                                            v92 = v74
                                        default:
                                            v95 := GenS1{F0: v92, F1: v72}
                                            _ = v95
                                        }
                                        _ = v92
                                    }
                                    v96 := []int{}
                                    v97 := GenS9{F0: 128, F1: (2 <= 278)}
                                    _ = v96
                                    _ = v97
                                }
                            default:
                                var v98 string = "\n\t"
                                _ = v98
                            }
                            var v99 uint8 = 42
                            _ = v83
                            _ = v99
                        }
                    }
                    _ = fi73
                    _ = v74
                }
                _ = v72
            }
            v100 := []int{}
            _ = fi40
            _ = v100
        }
        fmt.Println(v6, v3, v2)
    }
    if ((-288 > v6) && v2) {
        if (v2 || v2) {
            v37 += (v37 << ((137) & 63))
            v101 := GenS10{F0: (v6 << ((2) & 63))}
            fmt.Println(v3, v6, v37)
            _ = v101
        }
        v2 = (80 >= 100)
        var v102 bool = ((v1 + (v1 >> ((255) & 63))) == v1)
        _ = v102
    }
    v103 := GenS10{F0: (-(v6))}
    v104 := []int{v37, (v37 * v37), v37, (v37 / ((v37) | 1))}
    v105 := v5[(^(v37))]
    if (v2 || (18446744073709551615 <= 60)) {
        v6 = -9223372036854775808
        var v106 uint64 = 42
        switch "\n\t" {
        case "\n\t":
            if v2 {
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v37, v106, v37)
                default:
                    for fi107 := 0; fi107 < 8; fi107++ {
                        if ((v2 && ("foo bar" < ("\n\t" + ""))) && (v2 && true)) {
                            v3 += 1.5
                            v108 := GenS10{F0: (v6 % ((v6) | 1))}
                            for fi109 := 0; fi109 < 16; fi109++ {
                                switch ("" + ("\n\t" + "foo bar")) {
                                case "a":
                                    fmt.Println(fi109, fi107)
                                case "hello":
                                    switch fi109 {
                                    case -171:
                                        if true {
                                            var v110 uint = 65535
                                            var v111 uint64 = 4294967295
                                            _ = v110
                                            _ = v111
                                        } else {
                                            v104 = append(v104, fi107)
                                            switch "" {
                                            case "foo bar":
                                                var v112 bool = v2
                                                _ = v112
                                            case "a":
                                                v113 := GenS11{F0: 2147483647, F1: 32768, F2: 96}
                                                _ = v113
                                            case "":
                                                if true {
                                                    v3 += (v3 - 1e10)
                                                    var v114 int64 = 1
                                                    _ = v114
                                                } else {
                                                    fmt.Println(v6, v1)
                                                }
                                            default:
                                                if ((v2 || (v3 <= float64(0.0))) && (!v2)) {
                                                    if ((v2 || v2) || ((v6 - (v6 ^ -1)) > int64(v1))) {
                                                        delete(v4, 2147483648)
                                                        fmt.Println(fi109, v2)
                                                    } else {
                                                        v105 = -1
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                                _ = fi109
                            }
                            _ = v108
                        }
                        var v115 int64 = 9223372036854775807
                        _ = fi107
                        _ = v115
                    }
                }
            } else {
                var v116 uint32 = 1
                _ = v116
            }
        case "a":
            switch "" {
            case "foo bar":
                if v2 {
                    for fi117 := 0; fi117 < 16; fi117++ {
                        v37 = v105
                        v118 := GenS2{F0: 255, F1: 9223372036854775807, F2: ""}
                        _ = fi117
                        _ = v118
                    }
                    v119 := GenS8{F0: "", F1: (((("\n\t" + "hello") + "\n\t") + "") + ""), F2: v106}
                    _ = v119
                } else {
                    var v120 bool = ((-2.25 - 0.1) > ((3.14159 + -2.25) - 0.1))
                    if len(v104) >= 2 {
                        v104 = v104[1:len(v104)]
                    }
                    _ = v120
                }
            case "hello":
                var v121 uint = uint(v106)
                _ = v121
            case "a":
                v37 = v105
            }
        case "hello":
            if len(v104) > 0 {
                v104[0] = (v37 >> ((42) & 63))
            }
        }
        _ = v106
    } else {
        fmt.Println(v37, v37, v2)
        var v122 float64 = (v3 / (1.5 * ((float64(0.0) / (float64((1e10 - 1.5)) - v3)) * v3)))
        v37 = (v105 % ((250) | 1))
        _ = v122
    }
    v37 += v37
    _ = len(v104)
    v6 -= v6
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v104:", len(v104), v104)
    fmt.Println("v105:", v105)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v103
    _ = v104
    _ = v105
    _ = v2
    _ = v3
    _ = v34
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
