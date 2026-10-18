package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int16; F2 uint8 }
type GenS3 struct { F0 string; F1 int; F2 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 int8; F2 int }
type GenS6 struct { F0 int16; F1 uint16; F2 bool }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int8; F1 uint16 }
type GenS9 struct { F0 int; F1 bool }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 bool; F1 int16 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 string; F1 int; F2 uint16 }
func genFn0() uint32 {
    switch 12 {
    case -19:
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 int32 = -1
            _ = fi7
            _ = v8
        }
    case -823:
        var v9 uint8 = 36
        _ = v9
    default:
        v10 := []int{-511, 979, 1, 32767}
        _ = v10
    }
    var v11 int8 = -1
    _ = v11
    return 2147483648
}
func genFn1() int {
    var v12 bool = false
    var v13 string = "\n\t"
    _ = v12
    _ = v13
    return 2147483647
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{("foo bar" + "hello"), "foo bar"}
    var v3 uint = 42
    var v4 float32 = float32(1e10)
    var v5 bool = ("a" != ("a" + "a"))
    var v6 int = 1
    switch 32767 {
    case 4294967295:
        if v5 {
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 int16 = int16(v3)
                _ = fi14
                _ = v15
            }
            if (v5 || v5) {
                var v16 int32 = 0
                _ = v16
            } else {
                fmt.Println(v6)
                v3 -= v3
                for fi17 := 0; fi17 < 6; fi17++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi18 := 0; fi18 < 2; fi18++ {
                        v5 = true
                        delete(v1, "\n\t")
                        v19 := GenS1{F0: 32767}
                        _ = fi18
                        _ = v19
                    }
                    _ = fi17
                }
            }
        }
    case 42:
        v20 := GenS2{F0: (v3 * 2), F1: -32768, F2: 208}
        _ = v20
    case 291:
        v21 := GenS3{F0: "hello", F1: (v6 % ((v6) | 1)), F2: 35}
        _ = v21
    default:
        v22 := GenS3{F0: "a", F1: -463, F2: 2}
        _ = v22
    }
    switch v6 {
    case 32767:
        for fi23 := 0; fi23 < 6; fi23++ {
            v5 = true
            _ = fi23
        }
    case 9223372036854775807:
        v24 := GenS3{F0: "\n\t", F1: genFn1(), F2: 127}
        _ = v24
    }
    if true {
        switch "" {
        case "\n\t":
            var v25 float32 = (((((float32(0.1) - (1e10 + v4)) + ((1.5 * float32((0.1 - ((float64(v4) + (1e10 / (0.1 - (1e10 * 1e10)))) - 1e10)))) - (3.14159 + v4))) * v4) * v4) + float32(((1.5 / (0.0 + float64(v4))) + 0.0)))
            _ = v25
        case "foo bar":
            switch (v6 << (((v3 + (v3 - (v3 | (^(v3)))))) & 63)) {
            case 200:
                for fi26 := 0; fi26 < 9; fi26++ {
                    v6 = (-(fi26))
                    v27 := GenS4{F0: 1}
                    _ = fi26
                    _ = v27
                }
            case -613:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 0:
                if v5 {
                    var v28 int32 = -1
                    _ = v28
                } else {
                    v29 := GenS1{F0: int16(v3)}
                    _ = v29
                }
            default:
                for fi30 := 0; fi30 < 6; fi30++ {
                    for fi31 := 0; fi31 < 7; fi31++ {
                        _ = len(v2)
                        v4 = 3.14159
                        _ = fi31
                    }
                    var v32 uint16 = 270
                    _ = fi30
                    _ = v32
                }
            }
        default:
            switch v6 {
            case 42:
                for fi33 := 0; fi33 < 14; fi33++ {
                    v34 := GenS5{F0: 559, F1: -77, F2: (v6 - v6)}
                    _ = fi33
                    _ = v34
                }
            case 128:
                v35 := GenS6{F0: 32767, F1: 2, F2: v5}
                _ = v35
            default:
                v36 := GenS7{F0: 1}
                _ = v36
            }
        }
        switch "foo bar" {
        case "a":
            switch (v6 % ((v6) | 1)) {
            case 472:
                v37 := GenS8{F0: 0, F1: 1}
                _ = v37
            }
        case "hello":
            v38 := GenS9{F0: (v6 % ((genFn1()) | 1)), F1: true}
            _ = v38
        default:
            if true {
                fmt.Println(v6, v3, v5)
            } else {
                for fi39 := 0; fi39 < 11; fi39++ {
                    delete(v1, "foo bar")
                    _ = fi39
                }
                v4 -= v4
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
        for fi40 := 0; fi40 < 18; fi40++ {
            switch (v6 | 42) {
            case -636:
                var v41 int32 = -1
                _ = v41
            case -452:
                switch "\n\t" {
                case "hello":
                    var v42 uint8 = 128
                    _ = v42
                case "\n\t":
                    for fi43 := 0; fi43 < 4; fi43++ {
                        v44 := GenS8{F0: -92, F1: 2}
                        _ = fi43
                        _ = v44
                    }
                case "a":
                    if (v5 && true) {
                        for fi45 := 0; fi45 < 7; fi45++ {
                            if v5 {
                                v3 += v3
                            }
                            _ = fi45
                        }
                    } else {
                        if true {
                            switch (v6 % ((v6) | 1)) {
                            case 127:
                                v5 = true
                            case 2147483647:
                                if v5 {
                                    switch (v6 & fi40) {
                                    case 365:
                                        fmt.Println(v6)
                                    case 0:
                                        for fi46 := 0; fi46 < 4; fi46++ {
                                            fmt.Println(v6, fi40)
                                            delete(v1, "")
                                            delete(v1, "foo bar")
                                            _ = fi46
                                        }
                                    default:
                                        if v5 {
                                            var v47 uint64 = 1
                                            if (2 != 1) {
                                                switch "\n\t" {
                                                case "":
                                                    if ((v6 + 127) == 127) {
                                                        var v48 int32 = 794
                                                        v6 -= (-(v6))
                                                        if ((v48 & 2147483647) <= 0) {
                                                            var v49 int32 = (v48 | (v48 % ((947) | 1)))
                                                            if v5 {
                                                                for fi50 := 0; fi50 < 5; fi50++ {
                                                                    if (!v5) {
                                                                        v6 = (fi40 << (((v3 % ((v3) | 1))) & 63))
                                                                        var v51 int64 = 1
                                                                        var v52 int8 = -1
                                                                        _ = v51
                                                                        _ = v52
                                                                    } else {
                                                                        var v53 float32 = (float32(0.1) * ((v4 / (v4 / 1.0)) * v4))
                                                                        _ = v53
                                                                    }
                                                                    var v54 uint64 = (v47 ^ (v47 * 1))
                                                                    fmt.Println(v49)
                                                                    _ = fi50
                                                                    _ = v54
                                                                }
                                                                v2 = append(v2, (("" + "a") + ""))
                                                            } else {
                                                                v55 := GenS10{F0: "a"}
                                                                fmt.Println(v48)
                                                                _ = v55
                                                            }
                                                            _ = v49
                                                        } else {
                                                            v1["foo bar"] = (1 < v3)
                                                            for fi56 := 0; fi56 < 3; fi56++ {
                                                                if v5 {
                                                                    switch (("hello" + ("\n\t" + "a")) + "a") {
                                                                    case "":
                                                                        v3 += v3
                                                                    }
                                                                    fmt.Println(fi40, fi56, fi56)
                                                                } else {
                                                                    var v57 uint32 = 0
                                                                    v6 += (v6 << ((v3) & 63))
                                                                    v5 = v5
                                                                    _ = v57
                                                                }
                                                                _ = fi56
                                                            }
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        _ = v48
                                                    } else {
                                                        fmt.Println(v47)
                                                        var v58 string = "a"
                                                        _ = v58
                                                    }
                                                }
                                                v59 := GenS8{F0: 100, F1: 1}
                                                _ = v59
                                            } else {
                                                for fi60 := 0; fi60 < 11; fi60++ {
                                                    if len(v2) > 0 {
                                                        v2[0] = ("\n\t" + (("a" + ("\n\t" + ("a" + ((("\n\t" + "\n\t") + "") + "foo bar")))) + (("foo bar" + "foo bar") + ("foo bar" + "foo bar"))))
                                                    }
                                                    v61 := GenS3{F0: "foo bar", F1: genFn1(), F2: 255}
                                                    if (v5 && (!(8 != 1))) {
                                                        fmt.Println(v6, v3)
                                                        switch genFn1() {
                                                        case 0:
                                                            if true {
                                                                v5 = (1 < v47)
                                                            } else {
                                                                var v62 uint64 = (^(v47))
                                                                v6 -= (fi60 * -9223372036854775808)
                                                                _ = v62
                                                            }
                                                            switch "foo bar" {
                                                            case "hello":
                                                                switch fi40 {
                                                                case 127:
                                                                    switch fi40 {
                                                                    case -23:
                                                                        var v63 int = 42
                                                                        _ = v63
                                                                    case 583:
                                                                        var v64 int16 = -262
                                                                        _ = v64
                                                                    case 0:
                                                                        v6 = 128
                                                                    default:
                                                                        var v65 int64 = -81
                                                                        _ = v65
                                                                    }
                                                                }
                                                            default:
                                                                var v66 bool = ((65535 <= genFn0()) || false)
                                                                _ = v66
                                                            }
                                                            _ = len(v1)
                                                        case 677:
                                                            var v67 float64 = -2.25
                                                            _ = v67
                                                        case -2147483648:
                                                            v3 += 42
                                                        default:
                                                            v68 := v1[("hello" + ("foo bar" + "hello"))]
                                                            _ = v68
                                                        }
                                                    }
                                                    _ = fi60
                                                    _ = v61
                                                }
                                            }
                                            _ = v47
                                        } else {
                                            switch ("a" + (("hello" + ("\n\t" + (("hello" + "\n\t") + "a"))) + ("" + ""))) {
                                            case "a":
                                                var v69 bool = (!(!(genFn0() != genFn0())))
                                                _ = v69
                                            case "\n\t":
                                                var v70 int8 = 42
                                                _ = v70
                                            }
                                        }
                                    }
                                }
                            case 0:
                                v4 += (v4 - (v4 - 3.14159))
                            default:
                                if (v4 != v4) {
                                    v71 := GenS11{F0: (v3 % (((v3 & v3)) | 1))}
                                    _ = v71
                                }
                            }
                            v1["\n\t"] = v5
                            if false {
                                v72 := GenS12{F0: (18446744073709551615 > 42), F1: 124}
                                delete(v1, (("hello" + "\n\t") + "\n\t"))
                                v3 -= v3
                                _ = v72
                            }
                        }
                        _ = len(v1)
                    }
                default:
                    fmt.Println(fi40, v3)
                }
            default:
                v73 := GenS9{F0: genFn1(), F1: ((v3 + 124) != 997)}
                _ = v73
            }
            _ = fi40
        }
    } else {
        for fi74 := 0; fi74 < 10; fi74++ {
            var v75 int = (-(fi74))
            _ = fi74
            _ = v75
        }
        for fi76 := 0; fi76 < 1; fi76++ {
            var v77 uint = (v3 / ((v3) | 1))
            var v78 int16 = 42
            _ = fi76
            _ = v77
            _ = v78
        }
    }
    if v5 {
        var v79 bool = (true || v5)
        switch genFn1() {
        case -1:
            var v80 float32 = (v4 - v4)
            _ = v80
        default:
            v79 = (100 > 100)
        }
        _ = v79
    } else {
        fmt.Println(v4)
        switch (v6 % ((v6) | 1)) {
        case 556:
            var v81 float64 = 0.1
            _ = v81
        }
        v1["foo bar"] = true
    }
    v4 = (((0.1 - (1.5 + v4)) - (v4 / v4)) + v4)
    var v82 string = ("\n\t" + "foo bar")
    _ = len(v1)
    var v83 int64 = 1
    v84 := GenS13{F0: int(v3)}
    v85 := GenS4{F0: 82}
    if len(v2) > 0 {
        v2[0] = ""
    }
    if (v5 || ((v83 % (((v83 - (v83 + 1))) | 1)) < 9223372036854775807)) {
        if true {
            if v5 {
                var v86 uint16 = 606
                switch v82 {
                case "\n\t":
                    fmt.Println(v5, v6, v86)
                case "hello":
                    for fi87 := 0; fi87 < 10; fi87++ {
                        v6 = genFn1()
                        switch fi87 {
                        case 128:
                            if true {
                                fmt.Println(v82, v86)
                            }
                        case -313:
                            for fi88 := 0; fi88 < 17; fi88++ {
                                v89 := GenS14{F0: v82, F1: fi87, F2: v86}
                                _ = fi88
                                _ = v89
                            }
                        case -1:
                            fmt.Println(v83, fi87)
                        default:
                            if (v5 && v5) {
                                v5 = (v5 && v5)
                                switch fi87 {
                                case -1:
                                    v1["foo bar"] = (127 >= 127)
                                default:
                                    var v90 string = v82
                                    _ = v90
                                }
                                var v91 int16 = -429
                                _ = v91
                            } else {
                                var v92 int32 = -2147483648
                                v3 = v3
                                _ = v92
                            }
                        }
                        _ = fi87
                    }
                default:
                    for fi93 := 0; fi93 < 17; fi93++ {
                        for fi94 := 0; fi94 < 18; fi94++ {
                            var v95 float32 = (((v4 + float32(0.1)) + v4) * v4)
                            var v96 int8 = 1
                            _ = fi94
                            _ = v95
                            _ = v96
                        }
                        v3 += 1
                        _ = fi93
                    }
                }
                _ = v86
            }
        } else {
            var v97 float64 = (((0.1 / (float64(v4) / 3.14159)) + 1.5) / float64(v4))
            switch (v6 * (v6 | (v6 * 9223372036854775807))) {
            case 473:
                if (!v5) {
                    for fi98 := 0; fi98 < 7; fi98++ {
                        if (2147483647 < -953) {
                            v3 += (v3 >> ((1) & 63))
                        } else {
                            var v99 float64 = float64((3.14159 - (((float32(-2.25) * (float32(v97) + v4)) + 0.0) - v4)))
                            _ = v99
                        }
                        _ = fi98
                    }
                    v100 := v1[("" + v82)]
                    if v100 {
                        v2 = append(v2, v82)
                    }
                    _ = v100
                }
            case 4294967295:
                v101 := GenS3{F0: v82, F1: v6, F2: 127}
                _ = v101
            default:
                if v5 {
                    v82 = (v82 + v82)
                    if (0.1 < v4) {
                        v82 = v82
                        switch (v6 * (-(v6))) {
                        case 32768:
                            var v102 uint8 = 127
                            _ = v102
                        default:
                            if (127 == -97) {
                                var v103 uint = (^(v3))
                                switch "foo bar" {
                                case "":
                                    var v104 int64 = (v83 >> ((v103) & 63))
                                    _ = v104
                                default:
                                    for fi105 := 0; fi105 < 10; fi105++ {
                                        v4 -= v4
                                        _ = fi105
                                    }
                                }
                                _ = v103
                            } else {
                                var v106 uint16 = 42
                                _ = v106
                            }
                        }
                        fmt.Println(v97, v97, v97)
                    } else {
                        v3 -= (v3 % (((^(v3))) | 1))
                        fmt.Println(v3)
                    }
                    v2 = append(v2, "hello")
                }
            }
            switch (v6 / (((v6 ^ (-(v6)))) | 1)) {
            case 0:
                switch genFn1() {
                case 1:
                    for fi107 := 0; fi107 < 10; fi107++ {
                        v83 += 230
                        fmt.Println(v5, v6, v3)
                        _ = fi107
                    }
                case 32767:
                    v1[((v82 + v82) + v82)] = (42 >= 0)
                case 0:
                    var v108 float64 = (0.1 * v97)
                    _ = v108
                default:
                    for fi109 := 0; fi109 < 4; fi109++ {
                        var v110 int = 0
                        _ = fi109
                        _ = v110
                    }
                }
            case 9223372036854775807:
                var v111 uint = v3
                _ = v111
            default:
                v2 = append(v2, ("hello" + (v82 + "\n\t")))
            }
            _ = v97
        }
        switch (v6 | v6) {
        case -929:
            fmt.Println(v3, v5, v4)
        case 127:
            for fi112 := 0; fi112 < 1; fi112++ {
                var v113 uint8 = 135
                v5 = true
                v114 := GenS2{F0: 255, F1: 611, F2: uint8(v3)}
                _ = fi112
                _ = v113
                _ = v114
            }
        default:
            if (!false) {
                var v115 int8 = -28
                if (!v5) {
                    switch v6 {
                    case 9223372036854775807:
                        v116 := v1[v82]
                        _ = v116
                    case -500:
                        if (!true) {
                            for fi117 := 0; fi117 < 12; fi117++ {
                                switch (v82 + v82) {
                                case "hello":
                                    fmt.Println(v3, v82)
                                }
                                _ = fi117
                            }
                        }
                    default:
                        v118 := GenS4{F0: (v115 - v115)}
                        _ = v118
                    }
                }
                _ = v115
            }
        }
    } else {
        v82 = "foo bar"
        fmt.Println(v83, v4, v3)
    }
    v83 = -1
    var v119 int32 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v119:", v119)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v119
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v82
    _ = v83
    _ = v84
    _ = v85
}
