package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint; F1 string; F2 uint16 }
type GenS6 struct { F0 string; F1 int16; F2 int32 }
type GenS7 struct { F0 uint8; F1 uint32 }
type GenS8 struct { F0 int64; F1 int8; F2 int8 }
type GenS9 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint; F1 uint32; F2 uint }
type GenS12 struct { F0 uint; F1 int32; F2 int16 }
type GenS13 struct { F0 bool }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]uint32)
    var v3 uint8 = 127
    v4 := make(map[int]bool)
    v5 := []string{"a", "foo bar", ("a" + ("hello" + ((("\n\t" + "a") + "foo bar") + "a")))}
    _ = len(v1)
    if ((32767 >= 425) || false) {
        var v6 uint8 = 2
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
        _ = v6
    } else {
        for fi7 := 0; fi7 < 14; fi7++ {
            if (165 < 1) {
                if (65535 == 42) {
                    v8 := GenS1{F0: 0, F1: -1}
                    _ = len(v4)
                    _ = v8
                }
                if (!false) {
                    for fi9 := 0; fi9 < 18; fi9++ {
                        if false {
                            fmt.Println(fi7, v3)
                            var v10 bool = (-1 < -1)
                            v11 := GenS2{F0: 2, F1: 4294967295}
                            _ = v10
                            _ = v11
                        }
                        _ = fi9
                    }
                    fmt.Println(fi7, fi7)
                } else {
                    v3 = v3
                }
                var v12 int32 = 1
                _ = v12
            } else {
                v3 = (v3 - (v3 << ((255) & 63)))
                if false {
                    var v13 int32 = 1
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 16; fi14++ {
                        v3 = 255
                        _ = fi14
                    }
                    if false {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = len(v2)
                    }
                }
            }
            _ = fi7
        }
        var v15 int32 = 0
        var v16 int8 = 1
        _ = v15
        _ = v16
    }
    switch ("a" + (("a" + "hello") + ("" + ("foo bar" + "foo bar")))) {
    case "":
        fmt.Println(v3)
    case "foo bar":
        if false {
            fmt.Println(v3)
            _ = len(v5)
        } else {
            v17 := GenS2{F0: (v3 / ((127) | 1)), F1: 255}
            var v18 uint8 = 1
            _ = v17
            _ = v18
        }
    }
    v3 += 1
    var v19 int32 = 515
    v5 = append(v5, ("a" + "foo bar"))
    for fi20 := 0; fi20 < 4; fi20++ {
        switch fi20 {
        case 42:
            switch "a" {
            case "a":
                v19 = (v19 & (v19 ^ v19))
            case "foo bar":
                var v21 int8 = -56
                _ = v21
            case "hello":
                for fi22 := 0; fi22 < 5; fi22++ {
                    var v23 string = "\n\t"
                    if len(v5) > 0 {
                        v5[0] = v23
                    }
                    fmt.Println(fi22)
                    _ = fi22
                    _ = v23
                }
            }
        case -992:
            v3 += v3
        case 2147483647:
            v19 = v19
        default:
            v24 := GenS1{F0: 4294967295, F1: fi20}
            _ = v24
        }
        _ = fi20
    }
    switch -988 {
    case -709:
        var v25 uint64 = 42
        _ = v25
    case 32767:
        fmt.Println(v19, v3)
    default:
        if (1 == 100) {
            var v26 int16 = 32767
            _ = v26
        } else {
            if false {
                var v27 uint64 = 4294967295
                _ = v27
            } else {
                var v28 int = 4294967295
                var v29 uint64 = 0
                switch (v28 + (v28 & (v28 ^ v28))) {
                case -9223372036854775808:
                    var v30 int = (v28 + v28)
                    _ = v30
                case -618:
                    var v31 uint32 = uint32(v29)
                    _ = v31
                }
                _ = v28
                _ = v29
            }
            v32 := GenS2{F0: v3, F1: 1}
            v19 = 2147483647
            _ = v32
        }
    }
    if (!true) {
        v3 = 1
    }
    if false {
        v19 = v19
        if false {
            var v33 uint = 65535
            v4[0] = (false || true)
            var v34 bool = ((true && true) && (552 != 4294967295))
            _ = v33
            _ = v34
        } else {
            switch ("" + ((("hello" + ((("\n\t" + (("a" + (("" + "a") + ("hello" + "hello"))) + ("foo bar" + "a"))) + ("hello" + "a")) + (((("" + ("\n\t" + ("\n\t" + ""))) + "") + "") + ("foo bar" + ("a" + "hello"))))) + "\n\t") + ("" + ("\n\t" + (("hello" + ("\n\t" + "foo bar")) + "foo bar"))))) {
            case "\n\t":
                fmt.Println(v19, v19)
            case "":
                fmt.Println(v3, v19)
            default:
                switch ("" + "foo bar") {
                case "a":
                    var v35 uint32 = 831
                    _ = v35
                }
            }
            for fi36 := 0; fi36 < 4; fi36++ {
                v1["a"] = (true || (32767 == 42))
                switch ("foo bar" + ("\n\t" + "foo bar")) {
                case "a":
                    v5 = append(v5, "hello")
                case "hello":
                    if false {
                        var v37 uint32 = 2
                        delete(v1, "\n\t")
                        if len(v5) > 0 {
                            v5[0] = ""
                        }
                        _ = v37
                    } else {
                        for fi38 := 0; fi38 < 8; fi38++ {
                            switch fi36 {
                            case 32768:
                                switch (fi36 ^ 332) {
                                case -917:
                                    var v39 int16 = 459
                                    _ = v39
                                case 223:
                                    v3 -= 154
                                case -484:
                                    v5 = append(v5, "hello")
                                    fmt.Println(v3, fi36)
                                    if (((!(!false)) && (true || (fi38 <= fi36))) || true) {
                                        var v40 int32 = v19
                                        if false {
                                            switch fi38 {
                                            case 430:
                                                v3 += 2
                                            case 32767:
                                                v40 -= (v19 << ((460) & 63))
                                            }
                                            var v41 int32 = (-(v40))
                                            fmt.Println(v3)
                                            switch ((("\n\t" + ("" + "hello")) + "a") + ("a" + ((("\n\t" + ("foo bar" + "foo bar")) + "foo bar") + ("" + "foo bar")))) {
                                            case "\n\t":
                                                _ = len(v5)
                                            }
                                            var v42 uint8 = 40
                                            v43 := GenS1{F0: 885, F1: (fi36 * (fi38 ^ fi36))}
                                            _ = v41
                                            _ = v42
                                            _ = v43
                                        } else {
                                            v44 := GenS3{F0: 18446744073709551615, F1: ((!false) || true), F2: 2}
                                            if (((3.14159 + ((1.5 / 3.14159) - 0.1)) <= (((1e10 + 1.5) - 3.14159) * 1.5)) && true) {
                                                if true {
                                                    if true {
                                                        if false {
                                                            v19 = v40
                                                            switch fi36 {
                                                            case -1:
                                                                var v45 int64 = 454
                                                                _ = v45
                                                            case 128:
                                                                v46 := GenS1{F0: 54, F1: (fi38 >> ((160) & 63))}
                                                                _ = v46
                                                            case -9223372036854775808:
                                                                var v47 float32 = ((float32((0.0 * 3.14159)) / ((1e10 * 3.14159) / 1.0)) / 1.0)
                                                                _ = v47
                                                            default:
                                                                v48 := GenS1{F0: 2147483647, F1: fi36}
                                                                _ = v48
                                                            }
                                                            fmt.Println(v3, v3)
                                                            v4[fi38] = true
                                                            fmt.Println(fi36)
                                                            v49 := GenS1{F0: 65535, F1: 32767}
                                                            v4[(-(fi38))] = (2 == 220)
                                                            _ = v49
                                                        } else {
                                                            if false {
                                                                switch (fi36 / ((fi36) | 1)) {
                                                                case -9223372036854775808:
                                                                    var v50 float64 = 1.5
                                                                    _ = v50
                                                                case 127:
                                                                    var v51 uint64 = 4294967295
                                                                    if true {
                                                                        v5 = append(v5, "a")
                                                                        v40 -= -2147483648
                                                                        v40 = (v40 % ((-2147483648) | 1))
                                                                    }
                                                                    _ = v51
                                                                case -130:
                                                                    if (756 <= 42) {
                                                                        if (160 <= 601) {
                                                                            fmt.Println(v3, v19)
                                                                            v52 := GenS4{F0: 912}
                                                                            if (false || (!(!(!(uint16(v3) < 1))))) {
                                                                                if false {
                                                                                    v40 = v19
                                                                                }
                                                                                fmt.Println(v3, fi36, fi36)
                                                                                var v53 int8 = 100
                                                                                _ = v53
                                                                            }
                                                                            v3 = v3
                                                                            _ = v52
                                                                        }
                                                                    }
                                                                default:
                                                                    var v54 int = int(v3)
                                                                    _ = v54
                                                                }
                                                                if true {
                                                                    if ((false && (v3 < (v3 + v3))) && (!false)) {
                                                                        v55 := GenS5{F0: 65535, F1: ("a" + "a"), F2: 873}
                                                                        _ = v55
                                                                    }
                                                                    var v56 float64 = 1e10
                                                                    _ = v56
                                                                } else {
                                                                    var v57 float32 = -2.25
                                                                    var v58 uint = 18446744073709551615
                                                                    var v59 uint16 = 2
                                                                    v60 := GenS6{F0: "hello", F1: 32767, F2: 0}
                                                                    _ = v57
                                                                    _ = v58
                                                                    _ = v59
                                                                    _ = v60
                                                                }
                                                                var v61 uint8 = 2
                                                                _ = v61
                                                            } else {
                                                                switch (fi36 - 4294967295) {
                                                                case 128:
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                case -648:
                                                                    v40 += v19
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                }
                                                            }
                                                            fmt.Println(v40, v19)
                                                            if (true || true) {
                                                                v62 := GenS2{F0: (v3 - 2), F1: 4294967295}
                                                                var v63 uint8 = (v3 + v3)
                                                                _ = v62
                                                                _ = v63
                                                            }
                                                            if (v19 == v40) {
                                                                if (1 < 255) {
                                                                    v64 := GenS7{F0: v3, F1: 42}
                                                                    if (338 < 970) {
                                                                        v5 = append(v5, "foo bar")
                                                                        fmt.Println(v40, v40, v40)
                                                                        v19 -= (-(v19))
                                                                    } else {
                                                                        v3 += 103
                                                                        v5 = append(v5, "foo bar")
                                                                    }
                                                                    var v65 bool = (65535 > 255)
                                                                    _ = v64
                                                                    _ = v65
                                                                } else {
                                                                    var v66 uint8 = (v3 | (v3 ^ v3))
                                                                    var v67 bool = (v66 <= (v66 << ((65535) & 63)))
                                                                    _ = v66
                                                                    _ = v67
                                                                }
                                                                if false {
                                                                    var v68 uint8 = 90
                                                                    _ = len(v5)
                                                                    _ = v68
                                                                }
                                                                v3 = v3
                                                            }
                                                            var v69 int16 = -867
                                                            _ = v69
                                                        }
                                                        _ = len(v5)
                                                    } else {
                                                        if true {
                                                            if false {
                                                                v70 := GenS6{F0: "\n\t", F1: 32767, F2: v40}
                                                                _ = v70
                                                            }
                                                            var v71 uint8 = 37
                                                            fmt.Println(v71, fi36, v3)
                                                            fmt.Println(fi36)
                                                            var v72 uint32 = 2147483648
                                                            fmt.Println(v72)
                                                            switch 2147483647 {
                                                            case 128:
                                                                switch -2147483648 {
                                                                case 0:
                                                                    v19 -= v40
                                                                case -988:
                                                                    fmt.Println(fi36)
                                                                case -987:
                                                                    v73 := v1[((("\n\t" + "foo bar") + "") + "a")]
                                                                    _ = v73
                                                                default:
                                                                    if false {
                                                                        var v74 int16 = 1
                                                                        _ = v74
                                                                    }
                                                                }
                                                            default:
                                                                v75 := GenS8{F0: 0, F1: 1, F2: 42}
                                                                _ = v75
                                                            }
                                                            _ = v71
                                                            _ = v72
                                                        } else {
                                                            fmt.Println(v19, v3)
                                                            switch (fi38 % ((fi36) | 1)) {
                                                            case -9223372036854775808:
                                                                switch ("" + "\n\t") {
                                                                case "hello":
                                                                    var v76 string = "\n\t"
                                                                    _ = v76
                                                                case "\n\t":
                                                                    var v77 bool = (((1.5 * 3.14159) < float64(-2.25)) && ((1 > 175) && false))
                                                                    _ = v77
                                                                case "a":
                                                                    v2["a"] = 42
                                                                }
                                                            case -1:
                                                                var v78 float32 = 1e10
                                                                _ = v78
                                                            }
                                                            v79 := GenS9{F0: 55, F1: 32767, F2: -1}
                                                            _ = v79
                                                        }
                                                    }
                                                    v19 -= v40
                                                    v19 += -885
                                                }
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v19 = -693
                                                case "":
                                                    v2[(("\n\t" + "a") + "a")] = 90
                                                case "\n\t":
                                                    fmt.Println(fi36, fi36)
                                                }
                                                _ = len(v4)
                                                var v80 uint64 = 18446744073709551615
                                                v19 += v19
                                                if false {
                                                    var v81 uint32 = 398
                                                    v5 = append(v5, "hello")
                                                    _ = v81
                                                }
                                                _ = v80
                                            } else {
                                                fmt.Println(v19)
                                            }
                                            v3 -= v3
                                            _ = v44
                                        }
                                        _ = v40
                                    }
                                default:
                                    v5 = append(v5, "a")
                                }
                            case -349:
                                v82 := GenS10{F0: 772}
                                _ = v82
                            }
                            v83 := GenS4{F0: 535}
                            var v84 int8 = -128
                            var v85 uint = 255
                            var v86 bool = false
                            _ = fi38
                            _ = v83
                            _ = v84
                            _ = v85
                            _ = v86
                        }
                    }
                case "foo bar":
                    switch "hello" {
                    case "a":
                        fmt.Println(v3)
                    case "foo bar":
                        for fi87 := 0; fi87 < 20; fi87++ {
                            fmt.Println(v19, v19, fi36)
                            v3 = (v3 ^ 128)
                            _ = fi87
                        }
                    default:
                        v88 := GenS9{F0: 23, F1: 32767, F2: 0}
                        _ = v88
                    }
                default:
                    if (false || false) {
                        v19 = (v19 - v19)
                        var v89 float32 = (float32(((3.14159 / 3.14159) / -2.25)) + (-2.25 * 1e10))
                        var v90 int16 = -659
                        _ = v89
                        _ = v90
                    } else {
                        v91 := GenS11{F0: 964, F1: 2147483647, F2: 1}
                        _ = v91
                    }
                }
                for fi92 := 0; fi92 < 2; fi92++ {
                    _ = len(v4)
                    v5 = append(v5, ("foo bar" + "\n\t"))
                    _ = len(v4)
                    _ = fi92
                }
                _ = fi36
            }
            var v93 int16 = 799
            _ = v93
        }
        v5 = append(v5, "foo bar")
    }
    delete(v1, "foo bar")
    if true {
        switch 127 {
        case -500:
            v94 := GenS12{F0: 10, F1: (v19 ^ (v19 & v19)), F2: 32767}
            _ = v94
        case -559:
            var v95 int64 = 0
            _ = v95
        case 261:
            for fi96 := 0; fi96 < 3; fi96++ {
                v3 += v3
                _ = fi96
            }
        }
        for fi97 := 0; fi97 < 12; fi97++ {
            v2[("hello" + "foo bar")] = 2147483647
            for fi98 := 0; fi98 < 16; fi98++ {
                var v99 int64 = 0
                _ = fi98
                _ = v99
            }
            _ = fi97
        }
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v3 -= (v3 & (^(v3)))
        for fi100 := 0; fi100 < 4; fi100++ {
            fmt.Println(fi100)
            _ = fi100
        }
    }
    if (164 >= 0) {
        var v101 bool = (9223372036854775807 <= 0)
        if true {
            for fi102 := 0; fi102 < 20; fi102++ {
                if true {
                    if (!(!((v19 & v19) != (v19 << ((1) & 63))))) {
                        var v103 int = fi102
                        fmt.Println(fi102)
                        v101 = (v101 && true)
                        _ = v103
                    }
                    v19 = v19
                    v3 = v3
                } else {
                    v19 += int32(v3)
                    if (!v101) {
                        for fi104 := 0; fi104 < 11; fi104++ {
                            v3 += (v3 % ((v3) | 1))
                            var v105 float64 = 1.5
                            switch ("\n\t" + ("hello" + "a")) {
                            case "foo bar":
                                var v106 uint8 = 0
                                _ = v106
                            default:
                                var v107 uint64 = 4294967295
                                _ = v107
                            }
                            _ = fi104
                            _ = v105
                        }
                        var v108 bool = v101
                        _ = v108
                    }
                }
                _ = fi102
            }
            var v109 int = 128
            _ = v109
        } else {
            if (!(!v101)) {
                if false {
                    if v101 {
                        if v101 {
                            v3 += v3
                        }
                        v101 = v101
                    } else {
                        var v110 uint64 = 410
                        for fi111 := 0; fi111 < 13; fi111++ {
                            v112 := GenS13{F0: (true && (((0.1 + (3.14159 + (((0.1 / float32(0.1)) - float32(1.5)) - 3.14159))) - float32(3.14159)) != 3.14159))}
                            fmt.Println(v101, v110)
                            var v113 int = fi111
                            _ = fi111
                            _ = v112
                            _ = v113
                        }
                        v110 -= (v110 | 0)
                        _ = v110
                    }
                    fmt.Println(v101)
                    if true {
                        var v114 int8 = 42
                        for fi115 := 0; fi115 < 9; fi115++ {
                            v114 = v114
                            if v101 {
                                fmt.Println(fi115, v19, v114)
                                v114 = v114
                            }
                            _ = fi115
                        }
                        if false {
                            for fi116 := 0; fi116 < 14; fi116++ {
                                switch "\n\t" {
                                case "":
                                    v114 -= -10
                                case "\n\t":
                                    v19 = v19
                                case "hello":
                                    for fi117 := 0; fi117 < 16; fi117++ {
                                        var v118 int8 = (v114 + 46)
                                        v19 += (v19 % ((v19) | 1))
                                        if (!v101) {
                                            if (!(1 != 0)) {
                                                v119 := GenS11{F0: 341, F1: 1, F2: 255}
                                                var v120 float32 = 1e10
                                                fmt.Println(v120, v19)
                                                _ = v119
                                                _ = v120
                                            } else {
                                                if v101 {
                                                    delete(v1, "a")
                                                    if len(v5) > 0 {
                                                        v5[0] = "\n\t"
                                                    }
                                                }
                                                v114 -= v118
                                                v19 += (v19 >> ((2) & 63))
                                            }
                                        }
                                        v3 = v3
                                        _ = fi117
                                        _ = v118
                                    }
                                default:
                                    switch (((((("hello" + ("\n\t" + (("foo bar" + "\n\t") + ""))) + ("" + ("foo bar" + "a"))) + "hello") + "a") + "foo bar") + ("hello" + "foo bar")) {
                                    case "hello":
                                        v101 = v101
                                    case "foo bar":
                                        delete(v4, fi116)
                                    default:
                                        var v121 int8 = (v114 % ((127) | 1))
                                        _ = v121
                                    }
                                }
                                for fi122 := 0; fi122 < 2; fi122++ {
                                    v101 = v101
                                    _ = fi122
                                }
                                _ = fi116
                            }
                            v123 := GenS4{F0: 1}
                            _ = v123
                        }
                        _ = v114
                    }
                } else {
                    var v124 int32 = (v19 * -1)
                    _ = v124
                }
                for fi125 := 0; fi125 < 5; fi125++ {
                    _ = len(v1)
                    _ = fi125
                }
                v126 := GenS6{F0: ("foo bar" + (("hello" + "a") + "foo bar")), F1: 32767, F2: 0}
                _ = v126
            } else {
                if v101 {
                    delete(v2, "hello")
                }
                var v127 int32 = v19
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v127
            }
            var v128 uint32 = 2147483648
            fmt.Println(v3, v19)
            _ = v128
        }
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v3, v101)
        default:
            var v129 int64 = int64(v19)
            _ = v129
        }
        _ = v101
    }
    var v130 int64 = -9223372036854775808
    v19 = -259
    for fi131 := 0; fi131 < 12; fi131++ {
        v132 := GenS1{F0: 42, F1: -1}
        v19 -= (v19 + -286)
        v130 = (v130 >> ((2) & 63))
        _ = fi131
        _ = v132
    }
    var v133 uint = 4294967295
    if false {
        var v134 int8 = 42
        var v135 int64 = 42
        _ = v134
        _ = v135
    } else {
        for fi136 := 0; fi136 < 11; fi136++ {
            fmt.Println(v133, v3, fi136)
            var v137 float32 = float32(float64(-2.25))
            v5 = append(v5, ((((("hello" + "a") + "foo bar") + "hello") + "\n\t") + ("\n\t" + "a")))
            _ = fi136
            _ = v137
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v130:", v130)
    fmt.Println("v133:", v133)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v130
    _ = v133
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
