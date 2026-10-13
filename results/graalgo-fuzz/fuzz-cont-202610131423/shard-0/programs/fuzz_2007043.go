package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int32 }
type GenS2 struct { F0 string; F1 uint8; F2 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 uint; F2 bool }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 string; F1 uint; F2 int }
type GenS9 struct { F0 bool }

func main() {
    var v1 float32 = (1e10 * 1.5)
    var v2 bool = true
    var v3 string = "foo bar"
    v4 := make(map[string]string)
    var v5 float32 = 1e10
    var v6 bool = v2
    v7 := GenS1{F0: "\n\t", F1: 189, F2: 1}
    v8 := []int{}
    if v2 {
        if (42 == 0) {
            v9 := GenS1{F0: "foo bar", F1: 18446744073709551615, F2: -1}
            _ = v9
        } else {
            fmt.Println(v6)
            for fi10 := 0; fi10 < 18; fi10++ {
                v11 := GenS1{F0: v3, F1: 65535, F2: 0}
                var v12 int32 = 42
                _ = fi10
                _ = v11
                _ = v12
            }
        }
    }
    var v13 int = -9223372036854775808
    var v14 int8 = 71
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := GenS1{F0: "hello", F1: 42, F2: -1}
        var v17 float64 = 0.1
        v4[v3] = ("hello" + v3)
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 float32 = v5
    v19 := GenS1{F0: (("foo bar" + "a") + v3), F1: 706, F2: 1}
    for fi20 := 0; fi20 < 17; fi20++ {
        fmt.Println(v18, v6, v6)
        _ = fi20
    }
    fmt.Println(v13)
    fmt.Println(v5, v6, v14)
    v14 += 17
    v14 = v14
    if (true && v6) {
        if (1 == 4294967295) {
            _ = len(v4)
        } else {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
        v21 := GenS1{F0: "foo bar", F1: 1, F2: 2147483647}
        for fi22 := 0; fi22 < 9; fi22++ {
            for fi23 := 0; fi23 < 17; fi23++ {
                switch "\n\t" {
                case "":
                    var v24 int32 = 199
                    _ = v24
                case "hello":
                    v18 = (v18 / (-2.25 * v18))
                default:
                    v25 := GenS1{F0: (("" + v3) + (v3 + v3)), F1: 4294967295, F2: -536}
                    _ = v25
                }
                switch v3 {
                case "a":
                    v4[v3] = v3
                case "hello":
                    v5 += 0.1
                case "\n\t":
                    var v26 uint8 = 1
                    _ = v26
                default:
                    if v6 {
                        var v27 uint32 = 42
                        _ = v27
                    } else {
                        _ = len(v4)
                        var v28 int8 = v14
                        if v2 {
                            switch (-(fi23)) {
                            case 4294967295:
                                v8 = append(v8, fi23)
                            case 9223372036854775807:
                                v13 = (v13 % (((fi23 - 127)) | 1))
                            case 67:
                                v14 = (v28 * v28)
                            }
                        } else {
                            var v29 int8 = -128
                            var v30 int16 = 1000
                            if v2 {
                                var v31 bool = (!v6)
                                var v32 bool = (608 >= 0)
                                _ = v31
                                _ = v32
                            } else {
                                var v33 string = ("a" + v3)
                                _ = v33
                            }
                            _ = v29
                            _ = v30
                        }
                        _ = v28
                    }
                }
                v1 = (v5 * -2.25)
                _ = fi23
            }
            for fi34 := 0; fi34 < 1; fi34++ {
                v8 = append(v8, 32767)
                _ = fi34
            }
            if v6 {
                v35 := GenS2{F0: v3, F1: 1, F2: 594}
                v18 = v18
                if v2 {
                    v36 := GenS3{F0: ""}
                    if false {
                        v37 := GenS4{F0: -9223372036854775808}
                        _ = v37
                    } else {
                        var v38 int = v13
                        v13 = (v13 / ((-349) | 1))
                        v39 := GenS4{F0: -292}
                        _ = v38
                        _ = v39
                    }
                    v40 := GenS5{F0: 127, F1: 255, F2: v2}
                    _ = v36
                    _ = v40
                } else {
                    if len(v8) > 0 {
                        v8[0] = -866
                    }
                    if (!v6) {
                        delete(v4, (((v3 + (v3 + v3)) + v3) + (v3 + ((v3 + "\n\t") + v3))))
                    } else {
                        fmt.Println(v13, v13)
                        fmt.Println(v6)
                        var v41 int = v13
                        _ = v41
                    }
                    if (1 <= 2) {
                        _ = len(v8)
                    }
                }
                _ = v35
            } else {
                fmt.Println(fi22, v1)
                if v2 {
                    if ((0.1 - 3.14159) <= 0.1) {
                        for fi42 := 0; fi42 < 3; fi42++ {
                            if (!(255 >= 65535)) {
                                fmt.Println(v18, fi22)
                                v8 = append(v8, (fi42 >> ((535) & 63)))
                                var v43 float32 = v1
                                _ = v43
                            }
                            switch (v13 % (((fi42 << ((4294967295) & 63))) | 1)) {
                            case 32767:
                                var v44 int16 = -1
                                _ = v44
                            case -138:
                                if true {
                                    v13 = (v13 ^ 127)
                                    switch v3 {
                                    case "":
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                    case "a":
                                        v2 = v2
                                    case "hello":
                                        v45 := GenS3{F0: v3}
                                        _ = v45
                                    default:
                                        v46 := GenS2{F0: "", F1: 1, F2: 18446744073709551615}
                                        _ = v46
                                    }
                                }
                            }
                            _ = fi42
                        }
                        if ((-128 < v14) && v2) {
                            v1 = v5
                            switch -1 {
                            case 4294967295:
                                var v47 float32 = (float32(1e10) + 3.14159)
                                _ = v47
                            case 32768:
                                for fi48 := 0; fi48 < 7; fi48++ {
                                    if ((v3 + "hello") >= ("a" + v3)) {
                                        if ((42 <= 42) && (780 <= 170)) {
                                            v14 = (^(v14))
                                        }
                                        v14 += 0
                                    }
                                    v5 += 3.14159
                                    _ = fi48
                                }
                            case -9223372036854775808:
                                var v49 uint32 = 65535
                                _ = v49
                            default:
                                fmt.Println(v18)
                            }
                        }
                        var v50 float32 = v1
                        _ = v50
                    } else {
                        for fi51 := 0; fi51 < 4; fi51++ {
                            var v52 float64 = 1.5
                            fmt.Println(v52, v2)
                            var v53 int64 = -1
                            _ = fi51
                            _ = v52
                            _ = v53
                        }
                        delete(v4, v3)
                    }
                    var v54 float64 = (1e10 + (float64(v1) + (-2.25 * (0.1 * (-2.25 / 1.0)))))
                    _ = v54
                } else {
                    if ((v13 & (fi22 * 128)) > fi22) {
                        if (v6 || (v3 <= "foo bar")) {
                            for fi55 := 0; fi55 < 3; fi55++ {
                                fmt.Println(fi22)
                                switch (v3 + v3) {
                                case "":
                                    v56 := GenS6{F0: -9223372036854775808}
                                    _ = v56
                                case "foo bar":
                                    var v57 int16 = 42
                                    _ = v57
                                case "a":
                                    v3 = "\n\t"
                                default:
                                    v58 := v4[v3]
                                    _ = v58
                                }
                                _ = fi55
                            }
                            switch 32768 {
                            case -831:
                                var v59 int16 = -411
                                _ = v59
                            case 214:
                                switch (fi22 / (((^(fi22))) | 1)) {
                                case 4294967295:
                                    for fi60 := 0; fi60 < 19; fi60++ {
                                        v61 := GenS4{F0: 42}
                                        switch v13 {
                                        case 32768:
                                            v13 = 127
                                        case 378:
                                            fmt.Println(fi60, v3, fi22)
                                        case 2147483647:
                                            delete(v4, (("foo bar" + v3) + "\n\t"))
                                            if v6 {
                                                v62 := GenS3{F0: v3}
                                                v63 := GenS2{F0: v3, F1: 2, F2: 641}
                                                _ = v62
                                                _ = v63
                                            }
                                        default:
                                            var v64 bool = false
                                            _ = v64
                                        }
                                        _ = fi60
                                        _ = v61
                                    }
                                default:
                                    var v65 bool = v2
                                    _ = v65
                                }
                            default:
                                switch (-(v13)) {
                                case -336:
                                    if true {
                                        for fi66 := 0; fi66 < 16; fi66++ {
                                            v18 = (((v5 * v5) * ((v5 / (v1 * v1)) * ((v18 - -2.25) * (v5 * (v5 * (((v1 * (-2.25 / v1)) / ((v1 + ((0.0 / v18) - v5)) - (0.1 - (v1 / -2.25)))) + float32((-2.25 * ((0.1 - 3.14159) * -2.25))))))))) / (float32(1e10) - v18))
                                            _ = fi66
                                        }
                                    }
                                }
                            }
                        }
                        switch v13 {
                        case 2147483647:
                            v1 = -2.25
                        case -537:
                            switch (v13 / (((fi22 | -2147483648)) | 1)) {
                            case -2147483648:
                                if (v14 != v14) {
                                    v67 := GenS1{F0: v3, F1: 111, F2: -88}
                                    var v68 uint32 = 454
                                    _ = v67
                                    _ = v68
                                } else {
                                    v6 = (1 < 2)
                                }
                            case -1:
                                v3 = ""
                            }
                        default:
                            switch v3 {
                            case "hello":
                                switch ((("foo bar" + "") + v3) + v3) {
                                case "\n\t":
                                    switch (fi22 & fi22) {
                                    case 42:
                                        for fi69 := 0; fi69 < 7; fi69++ {
                                            v70 := GenS1{F0: (("\n\t" + (v3 + (("foo bar" + v3) + v3))) + v3), F1: 459, F2: -408}
                                            v71 := GenS4{F0: -702}
                                            if true {
                                                if (!true) {
                                                    if ("foo bar" == v3) {
                                                        v1 = 1.5
                                                        delete(v4, v3)
                                                    } else {
                                                        v72 := GenS2{F0: "foo bar", F1: 127, F2: 1}
                                                        _ = v72
                                                    }
                                                }
                                                v73 := GenS1{F0: v3, F1: 42, F2: 2147483647}
                                                _ = v73
                                            }
                                            _ = fi69
                                            _ = v70
                                            _ = v71
                                        }
                                    case 128:
                                        if (592 > 379) {
                                            switch (^(v13)) {
                                            case 32767:
                                                v8 = append(v8, fi22)
                                            case 1:
                                                if v2 {
                                                    v13 += (fi22 & fi22)
                                                } else {
                                                    var v74 int64 = 736
                                                    _ = v74
                                                }
                                            case 0:
                                                switch ("a" + v3) {
                                                case "\n\t":
                                                    if v6 {
                                                        var v75 int64 = -1
                                                        if v6 {
                                                            fmt.Println(v18)
                                                            fmt.Println(v13, v2, v3)
                                                        } else {
                                                            var v76 uint8 = 2
                                                            _ = v76
                                                        }
                                                        for fi77 := 0; fi77 < 3; fi77++ {
                                                            fmt.Println(v1, v1, v5)
                                                            var v78 float64 = 1.5
                                                            v75 -= 9223372036854775807
                                                            switch v13 {
                                                            case -9223372036854775808:
                                                                v8 = append(v8, (v13 >> ((815) & 63)))
                                                            case -895:
                                                                if true {
                                                                    fmt.Println(v6, fi77)
                                                                    if false {
                                                                        var v79 uint = 322
                                                                        if (((v79 / (((v79 & v79)) | 1)) < (v79 % (((^(v79))) | 1))) && ((v3 + v3) <= (("" + "a") + "\n\t"))) {
                                                                            _ = len(v4)
                                                                            v14 += 100
                                                                        } else {
                                                                            _ = len(v8)
                                                                            if (int16(v14) != 0) {
                                                                                fmt.Println(v18)
                                                                                v75 -= 580
                                                                                v79 = (v79 >> (((v79 >> (((^(v79))) & 63))) & 63))
                                                                            } else {
                                                                                var v80 string = (("\n\t" + v3) + (v3 + v3))
                                                                                if true {
                                                                                    v81 := GenS1{F0: v80, F1: 4294967295, F2: 42}
                                                                                    _ = v81
                                                                                }
                                                                                _ = v80
                                                                            }
                                                                        }
                                                                        _ = v79
                                                                    } else {
                                                                        if v6 {
                                                                            v1 = ((-2.25 + (v5 + 1e10)) - (v18 - 0.0))
                                                                        } else {
                                                                            delete(v4, "\n\t")
                                                                            var v82 int = fi22
                                                                            _ = v82
                                                                        }
                                                                        if (v13 > -9223372036854775808) {
                                                                            v83 := GenS1{F0: (v3 + v3), F1: 800, F2: -2147483648}
                                                                            var v84 int16 = 539
                                                                            _ = v83
                                                                            _ = v84
                                                                        } else {
                                                                            v8 = append(v8, (v13 << ((2) & 63)))
                                                                            var v85 float32 = v18
                                                                            v8 = append(v8, (-(v13)))
                                                                            _ = v85
                                                                        }
                                                                        v86 := GenS2{F0: v3, F1: 255, F2: 2}
                                                                        _ = v86
                                                                    }
                                                                }
                                                            default:
                                                                var v87 int = (^(fi77))
                                                                _ = v87
                                                            }
                                                            _ = fi77
                                                            _ = v78
                                                        }
                                                        _ = v75
                                                    }
                                                case "foo bar":
                                                    if v2 {
                                                        v88 := GenS4{F0: 9223372036854775807}
                                                        _ = v88
                                                    }
                                                default:
                                                    if ((476 > 75) || (1 > 18446744073709551615)) {
                                                        for fi89 := 0; fi89 < 3; fi89++ {
                                                            v6 = (!(false || v2))
                                                            v8 = append(v8, (fi22 - fi22))
                                                            _ = len(v4)
                                                            if (192 <= 120) {
                                                                v13 += (-(fi22))
                                                                var v90 string = v3
                                                                v6 = v2
                                                                _ = v90
                                                            }
                                                            v91 := GenS1{F0: (v3 + v3), F1: 641, F2: 42}
                                                            _ = fi89
                                                            _ = v91
                                                        }
                                                    }
                                                }
                                            default:
                                                var v92 int16 = 32767
                                                _ = v92
                                            }
                                        }
                                    case 127:
                                        v3 = "\n\t"
                                    }
                                case "a":
                                    var v93 int = (^(fi22))
                                    _ = v93
                                case "hello":
                                    var v94 bool = v2
                                    _ = v94
                                }
                            case "a":
                                fmt.Println(v1)
                            default:
                                if len(v8) > 0 {
                                    v8[0] = -1
                                }
                            }
                        }
                        v18 += v5
                    }
                    if (293 >= 762) {
                        for fi95 := 0; fi95 < 2; fi95++ {
                            switch (((v3 + v3) + v3) + (v3 + "a")) {
                            case "foo bar":
                                var v96 uint32 = 65535
                                _ = v96
                            default:
                                var v97 string = v3
                                _ = v97
                            }
                            _ = fi95
                        }
                    }
                }
            }
            _ = fi22
        }
        _ = v21
    } else {
        for fi98 := 0; fi98 < 18; fi98++ {
            var v99 int64 = -9223372036854775808
            fmt.Println(v3)
            switch (-(fi98)) {
            case -1:
                var v100 uint8 = 255
                _ = v100
            default:
                var v101 uint = 4294967295
                _ = v101
            }
            _ = fi98
            _ = v99
        }
    }
    if v6 {
        v102 := GenS5{F0: 1, F1: 4294967295, F2: (127 < 42)}
        _ = v102
    }
    for fi103 := 0; fi103 < 18; fi103++ {
        for fi104 := 0; fi104 < 11; fi104++ {
            var v105 uint32 = 539
            if (v6 || (0 >= -32768)) {
                switch 127 {
                case 0:
                    var v106 int = (fi103 << ((846) & 63))
                    _ = v106
                }
                if ((1.5 + (0.0 / (1.5 * (1e10 / 3.14159)))) != -2.25) {
                    v107 := GenS7{F0: (v105 + 789)}
                    v108 := GenS2{F0: v3, F1: 0, F2: 18446744073709551615}
                    v6 = false
                    _ = v107
                    _ = v108
                }
                var v109 int8 = (v14 | 1)
                fmt.Println(v3, v1)
                v110 := GenS8{F0: v3, F1: 42, F2: fi103}
                var v111 int64 = 0
                _ = v109
                _ = v110
                _ = v111
            }
            if v6 {
                v112 := GenS3{F0: ("" + (v3 + v3))}
                if len(v8) > 0 {
                    v8[0] = (fi104 | (-(fi104)))
                }
                _ = v112
            } else {
                v13 = (v13 / (((-(fi104))) | 1))
            }
            _ = fi104
            _ = v105
        }
        if v2 {
            v18 -= float32((0.0 / ((-2.25 + 0.1) + (0.1 * 1e10))))
            if ((fi103 | v13) < (^(fi103))) {
                v6 = (v6 || true)
                switch "" {
                case "":
                    v14 = v14
                case "hello":
                    if (!v6) {
                        for fi113 := 0; fi113 < 11; fi113++ {
                            v5 += float32((float64(-2.25) + (3.14159 / 1.5)))
                            v5 -= (((v18 - (v1 * (((3.14159 - v5) + v18) - (v5 * v5)))) + v1) / v1)
                            _ = fi113
                        }
                    }
                case "a":
                    var v114 uint16 = 898
                    _ = v114
                }
                switch (v13 & v13) {
                case -280:
                    switch ("a" + v3) {
                    case "\n\t":
                        for fi115 := 0; fi115 < 16; fi115++ {
                            var v116 string = (v3 + v3)
                            v117 := GenS1{F0: "foo bar", F1: 626, F2: 2147483647}
                            var v118 float64 = 1.5
                            _ = fi115
                            _ = v116
                            _ = v117
                            _ = v118
                        }
                    }
                }
            }
            v18 = v1
        }
        v119 := GenS9{F0: (!v6)}
        _ = fi103
        _ = v119
    }
    switch -685 {
    case 128:
        if len(v8) > 0 {
            v8[0] = v13
        }
    case 1:
        v13 += (v13 | v13)
    }
    v120 := GenS7{F0: 742}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v120
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
