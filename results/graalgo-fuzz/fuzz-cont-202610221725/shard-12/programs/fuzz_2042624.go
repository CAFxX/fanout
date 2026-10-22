package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint8 }
type GenS2 struct { F0 int; F1 int; F2 bool }
type GenS3 struct { F0 uint; F1 bool; F2 int64 }
type GenS4 struct { F0 int8; F1 string; F2 bool }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 uint64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int16; F1 int }
type GenS10 struct { F0 uint64; F1 int32; F2 uint32 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int16; F1 int8; F2 int64 }
type GenS13 struct { F0 uint8; F1 int64; F2 int16 }
type GenS14 struct { F0 int8; F1 uint; F2 uint64 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 int32; F1 int8; F2 uint64 }
func genFn0(p0 int8, p1 string, p2 uint16) uint32 {
    for fi4 := 0; fi4 < 4; fi4++ {
        p1 = "foo bar"
        v5 := []int{}
        var v6 int8 = p0
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := []int{-9223372036854775808}
    if (!(-9223372036854775808 > -239)) {
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 int = 9223372036854775807
            var v10 int = (fi8 - v9)
            v10 = v9
            _ = fi8
            _ = v10
            _ = v9
        }
        if true {
            for fi11 := 0; fi11 < 19; fi11++ {
                if (("a" + p1) < (p1 + ("a" + "foo bar"))) {
                    var v12 int8 = (p0 & p0)
                    p1 = ((("a" + "") + p1) + p1)
                    var v13 uint64 = 738
                    _ = v12
                    _ = v13
                }
                _ = fi11
            }
            for fi14 := 0; fi14 < 19; fi14++ {
                if (629 != 4294967295) {
                    if false {
                        fmt.Println(p2, p1, fi14)
                        fmt.Println(p0, fi14, fi14)
                    } else {
                        p0 += p0
                        switch fi14 {
                        case -9223372036854775808:
                            for fi15 := 0; fi15 < 1; fi15++ {
                                if true {
                                    var v16 uint64 = 848
                                    p2 = (p2 & 1)
                                    var v17 bool = false
                                    _ = v16
                                    _ = v17
                                }
                                p0 = 1
                                _ = fi15
                            }
                        default:
                            switch (fi14 & 127) {
                            case 32767:
                                v7 = append(v7, fi14)
                            case 128:
                                for fi18 := 0; fi18 < 5; fi18++ {
                                    v19 := GenS1{F0: fi14, F1: 65535, F2: 42}
                                    v7 = append(v7, 4294967295)
                                    if len(v7) > 0 {
                                        v7[0] = -462
                                    }
                                    _ = fi18
                                    _ = v19
                                }
                            default:
                                for fi20 := 0; fi20 < 20; fi20++ {
                                    p1 = p1
                                    _ = fi20
                                }
                            }
                        }
                    }
                    for fi21 := 0; fi21 < 10; fi21++ {
                        var v22 uint8 = 42
                        switch fi21 {
                        case 0:
                            switch p1 {
                            case "hello":
                                switch p1 {
                                case "foo bar":
                                    if true {
                                        v23 := GenS1{F0: fi14, F1: 213, F2: v22}
                                        v24 := GenS1{F0: fi14, F1: 18446744073709551615, F2: (v22 - (v22 / ((2) | 1)))}
                                        fmt.Println(p2, v22)
                                        _ = v23
                                        _ = v24
                                    }
                                }
                            case "a":
                                fmt.Println(p2, p0, fi14)
                            case "\n\t":
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            }
                        default:
                            if false {
                                var v25 int = 608
                                var v26 int64 = 1
                                if ((3.14159 != -2.25) || false) {
                                    if len(v7) > 0 {
                                        v7[0] = (v25 | (v25 + fi21))
                                    }
                                    var v27 uint32 = 2
                                    p0 = p0
                                    var v28 int32 = 2147483647
                                    v29 := GenS1{F0: (-(fi21)), F1: 0, F2: (v22 | (v22 % ((v22) | 1)))}
                                    v25 += (fi14 - (fi21 & fi21))
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                } else {
                                    if (!true) {
                                        v30 := GenS1{F0: (^(v25)), F1: 713, F2: v22}
                                        v31 := GenS1{F0: (^(v25)), F1: 65535, F2: 255}
                                        if true {
                                            if (309 != 0) {
                                                var v32 bool = true
                                                _ = v32
                                            } else {
                                                v26 -= (v26 % ((v26) | 1))
                                                fmt.Println(p1)
                                            }
                                            v26 = v26
                                        }
                                        _ = v30
                                        _ = v31
                                    }
                                    var v33 uint16 = p2
                                    v26 = v26
                                    _ = v33
                                }
                                fmt.Println(v26)
                                var v34 bool = (644 == 730)
                                v7 = append(v7, fi14)
                                _ = v25
                                _ = v26
                                _ = v34
                            } else {
                                var v35 int64 = 1
                                p1 = p1
                                _ = v35
                            }
                        }
                        p2 -= 32767
                        var v36 int32 = -798
                        var v37 int32 = v36
                        p2 = p2
                        fmt.Println(p0)
                        _ = fi21
                        _ = v22
                        _ = v36
                        _ = v37
                    }
                    p2 = (p2 * p2)
                }
                _ = fi14
            }
            for fi38 := 0; fi38 < 11; fi38++ {
                switch (fi38 | (fi38 * fi38)) {
                case -569:
                    for fi39 := 0; fi39 < 1; fi39++ {
                        if ((18446744073709551615 < 65535) || true) {
                            var v40 int = fi39
                            if (44 < 2147483647) {
                                p2 += p2
                            } else {
                                fmt.Println(v40, p0)
                                p1 = p1
                                var v41 string = p1
                                if (-2.25 >= 0.1) {
                                    var v42 int8 = 100
                                    var v43 uint16 = 447
                                    var v44 uint16 = (^(v43))
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                }
                                _ = v41
                            }
                            _ = v40
                        } else {
                            if len(v7) > 0 {
                                v7[0] = -388
                            }
                            fmt.Println(p1, fi39)
                        }
                        _ = fi39
                    }
                case 2147483647:
                    var v45 int8 = p0
                    _ = v45
                default:
                    v7 = append(v7, (fi38 % (((fi38 << ((0) & 63))) | 1)))
                }
                p2 = p2
                for fi46 := 0; fi46 < 8; fi46++ {
                    if false {
                        if (false || (!((true || false) || (p2 > p2)))) {
                            switch 828 {
                            case -9223372036854775808:
                                var v47 int16 = 42
                                _ = v47
                            case 9223372036854775807:
                                v48 := GenS1{F0: (fi46 + fi38), F1: 65535, F2: 128}
                                _ = v48
                            default:
                                switch p1 {
                                case "hello":
                                    v49 := GenS2{F0: fi46, F1: fi46, F2: (0 == 2147483647)}
                                    v50 := GenS2{F0: (fi46 % (((fi46 * -619)) | 1)), F1: (fi46 | (fi38 - fi38)), F2: (127 >= 160)}
                                    if false {
                                        v51 := GenS3{F0: 2, F1: false, F2: -9223372036854775808}
                                        p1 = p1
                                        _ = v51
                                    } else {
                                        fmt.Println(p1, p0)
                                    }
                                    _ = v49
                                    _ = v50
                                }
                            }
                            v52 := GenS4{F0: p0, F1: (p1 + ""), F2: true}
                            _ = v52
                        }
                        v53 := GenS5{F0: -9223372036854775808, F1: (p0 | p0)}
                        _ = v53
                    }
                    p1 = "foo bar"
                    _ = fi46
                }
                _ = fi38
            }
        } else {
            v7 = append(v7, int(p2))
            if (682 >= 4294967295) {
                switch 9223372036854775807 {
                case -275:
                    for fi54 := 0; fi54 < 15; fi54++ {
                        var v55 uint8 = 2
                        _ = fi54
                        _ = v55
                    }
                }
            } else {
                for fi56 := 0; fi56 < 18; fi56++ {
                    if (true && false) {
                        switch p1 {
                        case "hello":
                            var v57 uint32 = 42
                            _ = v57
                        case "a":
                            if (0 > 745) {
                                v58 := GenS5{F0: -9223372036854775808, F1: 127}
                                fmt.Println(fi56)
                                for fi59 := 0; fi59 < 8; fi59++ {
                                    v60 := GenS6{F0: 42}
                                    _ = fi59
                                    _ = v60
                                }
                                _ = v58
                            } else {
                                v61 := GenS3{F0: 255, F1: true, F2: -843}
                                var v62 uint32 = 756
                                v63 := GenS5{F0: 1, F1: p0}
                                _ = v61
                                _ = v62
                                _ = v63
                            }
                        }
                    } else {
                        for fi64 := 0; fi64 < 9; fi64++ {
                            p2 += 0
                            v65 := GenS3{F0: 792, F1: true, F2: -310}
                            _ = fi64
                            _ = v65
                        }
                        _ = len(v7)
                        p1 = (("a" + ((("hello" + ((p1 + "\n\t") + "")) + "hello") + (p1 + "a"))) + ("foo bar" + (p1 + (("hello" + "a") + p1))))
                    }
                    if ((false || true) || true) {
                        if (!((false || (-9223372036854775808 == -311)) && false)) {
                            var v66 string = p1
                            v66 = ("" + ("foo bar" + "a"))
                            var v67 uint8 = 128
                            _ = v66
                            _ = v67
                        }
                        for fi68 := 0; fi68 < 5; fi68++ {
                            v69 := GenS1{F0: (fi68 * (fi56 * fi56)), F1: 42, F2: 0}
                            p2 = p2
                            fmt.Println(p1)
                            _ = fi68
                            _ = v69
                        }
                    } else {
                        if true {
                            if len(v7) > 0 {
                                v7[0] = (fi56 % ((fi56) | 1))
                            }
                        }
                        var v70 int64 = 1
                        if (fi56 >= 0) {
                            for fi71 := 0; fi71 < 15; fi71++ {
                                fmt.Println(v70)
                                switch "\n\t" {
                                case "\n\t":
                                    v72 := GenS2{F0: -47, F1: (fi71 | 4294967295), F2: (770 == (v70 / (((v70 | v70)) | 1)))}
                                    _ = v72
                                case "hello":
                                    if len(v7) > 0 {
                                        v7[0] = int(v70)
                                    }
                                    p2 += p2
                                }
                                switch -9223372036854775808 {
                                case 128:
                                    p1 = p1
                                case 401:
                                    var v73 uint8 = 127
                                    _ = v73
                                default:
                                    if (!false) {
                                        if true {
                                            p2 += p2
                                            var v74 bool = false
                                            fmt.Println(p1, v74, v74)
                                            _ = v74
                                        } else {
                                            v70 += v70
                                        }
                                    }
                                }
                                _ = fi71
                            }
                            v70 -= v70
                        }
                        _ = v70
                    }
                    _ = fi56
                }
                p2 = p2
                switch (p1 + (("hello" + p1) + (p1 + ("" + "hello")))) {
                case "foo bar":
                    switch 0 {
                    case -2147483648:
                        if false {
                            p2 = p2
                            var v75 uint = 2
                            _ = v75
                        }
                    case 2147483647:
                        if false {
                            if true {
                                p0 = (p0 % (((^(p0))) | 1))
                                switch ((p1 + "foo bar") + p1) {
                                case "":
                                    if len(v7) > 0 {
                                        v7[0] = 4294967295
                                    }
                                case "hello":
                                    var v76 int16 = -390
                                    _ = v76
                                default:
                                    if (false || (p2 <= p2)) {
                                        _ = len(v7)
                                        if (1 == 2147483648) {
                                            if (4294967295 > 269) {
                                                if (1000 > 32767) {
                                                    p0 = p0
                                                    if len(v7) >= 2 {
                                                        v7 = v7[1:len(v7)]
                                                    }
                                                    if len(v7) >= 2 {
                                                        v7 = v7[1:len(v7)]
                                                    }
                                                } else {
                                                    if true {
                                                        var v77 uint32 = 2147483647
                                                        for fi78 := 0; fi78 < 9; fi78++ {
                                                            switch fi78 {
                                                            case 1:
                                                                v79 := GenS2{F0: (-(fi78)), F1: (^(fi78)), F2: (159 < 247)}
                                                                _ = v79
                                                            default:
                                                                for fi80 := 0; fi80 < 16; fi80++ {
                                                                    switch (p1 + p1) {
                                                                    case "\n\t":
                                                                        var v81 int32 = 1
                                                                        _ = v81
                                                                    case "foo bar":
                                                                        v7 = append(v7, (fi80 + fi78))
                                                                    case "hello":
                                                                        v77 = v77
                                                                    }
                                                                    if false {
                                                                        var v82 uint16 = 32767
                                                                        p2 = v82
                                                                        _ = v82
                                                                    } else {
                                                                        v83 := GenS3{F0: 1, F1: (!false), F2: -9223372036854775808}
                                                                        p2 = 2
                                                                        p0 = p0
                                                                        _ = v83
                                                                    }
                                                                    _ = fi80
                                                                }
                                                            }
                                                            if (51 >= 18446744073709551615) {
                                                                if len(v7) > 0 {
                                                                    v7[0] = (fi78 & (fi78 * 32768))
                                                                }
                                                                var v84 uint32 = v77
                                                                _ = v84
                                                            } else {
                                                                var v85 uint16 = (p2 / (((p2 * p2)) | 1))
                                                                _ = v85
                                                            }
                                                            _ = fi78
                                                        }
                                                        _ = v77
                                                    }
                                                }
                                                v7 = append(v7, 4294967295)
                                            } else {
                                                v86 := GenS7{F0: p2, F1: 0}
                                                v7 = append(v7, 4294967295)
                                                _ = v86
                                            }
                                            var v87 bool = true
                                            _ = v87
                                        } else {
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                            for fi88 := 0; fi88 < 16; fi88++ {
                                                var v89 string = (("hello" + p1) + ((p1 + "a") + "\n\t"))
                                                v90 := GenS8{F0: (p1 + ("hello" + v89))}
                                                switch -881 {
                                                case -2147483648:
                                                    v89 = "foo bar"
                                                default:
                                                    var v91 int16 = 1
                                                    _ = v91
                                                }
                                                _ = fi88
                                                _ = v89
                                                _ = v90
                                            }
                                            switch 702 {
                                            case -701:
                                                p2 -= (p2 | p2)
                                            case 163:
                                                v7 = append(v7, -2147483648)
                                            case 985:
                                                if (!false) {
                                                    fmt.Println(p2, p1)
                                                } else {
                                                    if (4294967295 > 2) {
                                                        if (-9223372036854775808 == 638) {
                                                            p2 = p2
                                                        } else {
                                                            if false {
                                                                if false {
                                                                    var v92 uint8 = 128
                                                                    _ = v92
                                                                }
                                                            }
                                                            if false {
                                                                p2 = p2
                                                                if false {
                                                                    p0 += -80
                                                                    if true {
                                                                        fmt.Println(p2, p1)
                                                                        v93 := GenS6{F0: 453}
                                                                        p0 += (p0 << ((255) & 63))
                                                                        _ = v93
                                                                    } else {
                                                                        p1 = p1
                                                                        var v94 string = "foo bar"
                                                                        _ = v94
                                                                    }
                                                                } else {
                                                                    p2 -= (p2 + (p2 + p2))
                                                                    var v95 int32 = 919
                                                                    _ = v95
                                                                }
                                                                for fi96 := 0; fi96 < 5; fi96++ {
                                                                    for fi97 := 0; fi97 < 9; fi97++ {
                                                                        if ((-2.25 < 0.0) || ((1.5 + 0.0) < 0.1)) {
                                                                            fmt.Println(p1, p0, p1)
                                                                            v98 := GenS6{F0: 18446744073709551615}
                                                                            _ = v98
                                                                        }
                                                                        _ = fi97
                                                                    }
                                                                    if false {
                                                                        v99 := GenS9{F0: 0, F1: fi96}
                                                                        if ((1.5 / 3.14159) == 1e10) {
                                                                            if len(v7) > 0 {
                                                                                v7[0] = 9223372036854775807
                                                                            }
                                                                            var v100 int64 = -508
                                                                            var v101 int = (fi96 - fi96)
                                                                            _ = v100
                                                                            _ = v101
                                                                        }
                                                                        _ = v99
                                                                    } else {
                                                                        var v102 uint64 = 2
                                                                        _ = v102
                                                                    }
                                                                    if (!((!((((p0 * 100) >= 0) || true) || true)) || (("a" == p1) && false))) {
                                                                        for fi103 := 0; fi103 < 1; fi103++ {
                                                                            var v104 int16 = 32767
                                                                            var v105 int16 = v104
                                                                            v106 := GenS6{F0: 153}
                                                                            _ = fi103
                                                                            _ = v104
                                                                            _ = v105
                                                                            _ = v106
                                                                        }
                                                                        var v107 string = ((p1 + (p1 + p1)) + "")
                                                                        _ = v107
                                                                    }
                                                                    _ = fi96
                                                                }
                                                            } else {
                                                                switch "\n\t" {
                                                                case "a":
                                                                    if false {
                                                                        switch p1 {
                                                                        case "":
                                                                            fmt.Println(p2, p0)
                                                                        }
                                                                        var v108 uint8 = 128
                                                                        _ = len(v7)
                                                                        _ = v108
                                                                    }
                                                                case "\n\t":
                                                                    var v109 uint8 = 1
                                                                    _ = v109
                                                                case "foo bar":
                                                                    var v110 uint = 65535
                                                                    _ = v110
                                                                }
                                                                var v111 int = 127
                                                                _ = v111
                                                            }
                                                        }
                                                        for fi112 := 0; fi112 < 14; fi112++ {
                                                            fmt.Println(p1, p2, fi112)
                                                            _ = fi112
                                                        }
                                                    } else {
                                                        var v113 uint = 827
                                                        _ = v113
                                                    }
                                                    for fi114 := 0; fi114 < 2; fi114++ {
                                                        _ = len(v7)
                                                        _ = fi114
                                                    }
                                                }
                                            default:
                                                switch p1 {
                                                case "":
                                                    if true {
                                                        if len(v7) > 0 {
                                                            v7[0] = 1
                                                        }
                                                        for fi115 := 0; fi115 < 5; fi115++ {
                                                            var v116 bool = true
                                                            switch (p1 + "hello") {
                                                            case "":
                                                                switch p1 {
                                                                case "a":
                                                                    var v117 bool = (!(fi115 != fi115))
                                                                    _ = v117
                                                                case "hello":
                                                                    var v118 uint16 = 74
                                                                    _ = v118
                                                                }
                                                            case "\n\t":
                                                                v119 := GenS10{F0: 492, F1: 1, F2: uint32(p2)}
                                                                _ = v119
                                                            case "a":
                                                                p1 = (p1 + ((("" + "\n\t") + ((p1 + ((p1 + "foo bar") + p1)) + (p1 + ""))) + "\n\t"))
                                                            default:
                                                                for fi120 := 0; fi120 < 4; fi120++ {
                                                                    fmt.Println(p2)
                                                                    _ = fi120
                                                                }
                                                            }
                                                            var v121 uint16 = p2
                                                            _ = fi115
                                                            _ = v116
                                                            _ = v121
                                                        }
                                                    } else {
                                                        if (true && ((1e10 * 1.5) > (3.14159 * 0.1))) {
                                                            v7 = append(v7, -9223372036854775808)
                                                            p2 -= (p2 ^ (p2 << ((18446744073709551615) & 63)))
                                                        }
                                                        for fi122 := 0; fi122 < 2; fi122++ {
                                                            if len(v7) > 0 {
                                                                v7[0] = fi122
                                                            }
                                                            fmt.Println(fi122)
                                                            switch (^(fi122)) {
                                                            case -9223372036854775808:
                                                                fmt.Println(fi122, p1, p2)
                                                            default:
                                                                for fi123 := 0; fi123 < 3; fi123++ {
                                                                    var v124 uint32 = 1
                                                                    switch fi122 {
                                                                    case -102:
                                                                        p2 += 1
                                                                    }
                                                                    v124 = 2
                                                                    var v125 int32 = -928
                                                                    v126 := GenS4{F0: 83, F1: "foo bar", F2: (((p0 * (p0 + p0)) > p0) || ((true || true) || (!true)))}
                                                                    _ = fi123
                                                                    _ = v124
                                                                    _ = v125
                                                                    _ = v126
                                                                }
                                                            }
                                                            _ = fi122
                                                        }
                                                    }
                                                case "\n\t":
                                                    var v127 int = -1
                                                    _ = v127
                                                default:
                                                    if (231 >= 2) {
                                                        fmt.Println(p2, p1, p0)
                                                        v128 := GenS11{F0: true}
                                                        _ = v128
                                                    } else {
                                                        v129 := GenS12{F0: -178, F1: (p0 % ((p0) | 1)), F2: -867}
                                                        switch 0 {
                                                        case 128:
                                                            _ = len(v7)
                                                        case 42:
                                                            v7 = append(v7, 0)
                                                        case 32768:
                                                            var v130 int = 792
                                                            _ = v130
                                                        }
                                                        p1 = (p1 + ("foo bar" + p1))
                                                        _ = v129
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                            } else {
                                switch "\n\t" {
                                case "hello":
                                    for fi131 := 0; fi131 < 7; fi131++ {
                                        if true {
                                            var v132 float64 = 1.5
                                            if true {
                                                var v133 float64 = v132
                                                fmt.Println(p2, p0, v132)
                                                var v134 float64 = 0.0
                                                _ = v133
                                                _ = v134
                                            }
                                            _ = len(v7)
                                            _ = v132
                                        } else {
                                            if true {
                                                fmt.Println(p2)
                                                _ = len(v7)
                                            } else {
                                                v135 := GenS9{F0: 1, F1: fi131}
                                                if (true || (42 <= (fi131 % ((-109) | 1)))) {
                                                    for fi136 := 0; fi136 < 4; fi136++ {
                                                        p0 -= (p0 % (((p0 + (p0 * (p0 % ((int8(p2)) | 1))))) | 1))
                                                        _ = fi136
                                                    }
                                                }
                                                if len(v7) >= 2 {
                                                    v7 = v7[1:len(v7)]
                                                }
                                                _ = v135
                                            }
                                            for fi137 := 0; fi137 < 2; fi137++ {
                                                switch (p1 + ("a" + "")) {
                                                case "a":
                                                    fmt.Println(fi131, p1, p0)
                                                case "\n\t":
                                                    var v138 int = fi131
                                                    _ = v138
                                                }
                                                var v139 float64 = -2.25
                                                _ = fi137
                                                _ = v139
                                            }
                                        }
                                        p1 = ""
                                        _ = fi131
                                    }
                                case "\n\t":
                                    fmt.Println(p2)
                                default:
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                }
                                fmt.Println(p1, p0, p0)
                                if (!(!((!false) || (!(4294967295 > 386))))) {
                                    p1 = p1
                                    p0 -= -97
                                } else {
                                    p1 = p1
                                    p2 -= 255
                                    var v140 int = -9223372036854775808
                                    _ = v140
                                }
                            }
                            p2 -= (p2 * 32767)
                        } else {
                            p0 -= -20
                            if true {
                                p2 = p2
                                v7 = append(v7, -344)
                                if (float64(1.5) < 0.1) {
                                    p1 = "hello"
                                    if (((^(p2)) >= p2) && (false || true)) {
                                        if (3.14159 > 3.14159) {
                                            p0 += p0
                                            _ = len(v7)
                                        }
                                        for fi141 := 0; fi141 < 18; fi141++ {
                                            if (-1 < (-(p0))) {
                                                switch (((p1 + ("\n\t" + p1)) + p1) + p1) {
                                                case "hello":
                                                    p0 = p0
                                                case "foo bar":
                                                    for fi142 := 0; fi142 < 16; fi142++ {
                                                        p0 = 100
                                                        v143 := GenS13{F0: 210, F1: 1, F2: 42}
                                                        p2 = p2
                                                        _ = fi142
                                                        _ = v143
                                                    }
                                                case "":
                                                    var v144 float32 = -2.25
                                                    _ = v144
                                                }
                                                var v145 uint32 = 4294967295
                                                fmt.Println(p2, fi141, p0)
                                                _ = v145
                                            } else {
                                                for fi146 := 0; fi146 < 11; fi146++ {
                                                    v147 := GenS10{F0: 18446744073709551615, F1: 2147483647, F2: 2147483647}
                                                    _ = fi146
                                                    _ = v147
                                                }
                                                var v148 uint = 18446744073709551615
                                                fmt.Println(fi141, v148)
                                                _ = v148
                                            }
                                            _ = fi141
                                        }
                                        p2 += p2
                                    } else {
                                        v149 := GenS14{F0: p0, F1: 65535, F2: 4294967295}
                                        _ = v149
                                    }
                                }
                            } else {
                                if false {
                                    for fi150 := 0; fi150 < 2; fi150++ {
                                        fmt.Println(fi150, p0, p2)
                                        var v151 uint = 1
                                        _ = fi150
                                        _ = v151
                                    }
                                } else {
                                    v152 := GenS10{F0: 330, F1: 42, F2: 65535}
                                    _ = v152
                                }
                            }
                            _ = len(v7)
                        }
                    case -9223372036854775808:
                        v153 := GenS15{F0: 0}
                        _ = v153
                    }
                case "\n\t":
                    var v154 bool = false
                    _ = v154
                case "a":
                    switch 128 {
                    case 32768:
                        v155 := GenS3{F0: 2, F1: false, F2: 42}
                        _ = v155
                    }
                default:
                    fmt.Println(p1, p1)
                }
            }
        }
        fmt.Println(p2, p1)
    } else {
        if false {
            if (0 <= 42) {
                for fi156 := 0; fi156 < 13; fi156++ {
                    p1 = p1
                    _ = fi156
                }
            } else {
                var v157 uint = 2
                p2 = p2
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = v157
            }
        } else {
            for fi158 := 0; fi158 < 2; fi158++ {
                var v159 int64 = 357
                var v160 uint = 42
                _ = fi158
                _ = v159
                _ = v160
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1(p0 int64, p1 uint32) uint {
    p0 = p0
    var v161 uint16 = 32767
    p0 += 9223372036854775807
    var v162 uint32 = (p1 ^ (p1 & genFn0(100, "foo bar", v161)))
    _ = p0
    _ = p1
    _ = v161
    _ = v162
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 2
    v3 := make(map[string]bool)
    var v163 int8 = 100
    var v164 int8 = (^(v163))
    v165 := GenS4{F0: (v164 ^ v163), F1: "a", F2: true}
    var v166 int8 = (v163 & -125)
    fmt.Println(v2)
    if (18446744073709551615 <= 65535) {
        for fi167 := 0; fi167 < 18; fi167++ {
            v168 := v1[429]
            _ = fi167
            _ = v168
        }
        if true {
            switch 32767 {
            case 9223372036854775807:
                for fi169 := 0; fi169 < 5; fi169++ {
                    if true {
                        if false {
                            var v170 uint8 = 72
                            v171 := GenS7{F0: 4, F1: 4294967295}
                            _ = v170
                            _ = v171
                        } else {
                            fmt.Println(v2, v163, v2)
                        }
                    } else {
                        if false {
                            var v172 string = "a"
                            v173 := []int{fi169, (fi169 >> ((42) & 63)), (fi169 - (fi169 & 9223372036854775807))}
                            _ = v172
                            _ = v173
                        } else {
                            delete(v1, 2147483647)
                        }
                        v174 := GenS16{F0: 0, F1: (v166 % ((100) | 1)), F2: 18446744073709551615}
                        v2 -= v2
                        _ = v174
                    }
                    var v175 float32 = -2.25
                    v3["hello"] = (!((-609 <= 2147483647) || (!true)))
                    _ = fi169
                    _ = v175
                }
            case 740:
                var v176 uint32 = 2147483647
                _ = v176
            }
            v2 += (^(v2))
            v3["foo bar"] = true
        }
        v177 := GenS9{F0: 42, F1: -643}
        _ = v177
    }
    var v178 int16 = 42
    v166 += (v166 ^ v164)
    for fi179 := 0; fi179 < 16; fi179++ {
        switch "hello" {
        case "a":
            if (false || (true || (true && true))) {
                v180 := GenS12{F0: v178, F1: int8(fi179), F2: -1}
                var v181 float64 = 1.5
                _ = v180
                _ = v181
            }
        default:
            v163 = v166
        }
        switch "" {
        case "hello":
            for fi182 := 0; fi182 < 10; fi182++ {
                v164 = (v163 >> ((2) & 63))
                _ = fi182
            }
        }
        if true {
            fmt.Println(v164, v178)
        } else {
            var v183 uint16 = 1
            _ = v183
        }
        _ = fi179
    }
    v184 := GenS4{F0: (v163 % ((v166) | 1)), F1: ("foo bar" + (((("a" + "\n\t") + "") + ("\n\t" + "")) + ("\n\t" + ("\n\t" + "")))), F2: true}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v163:", v163)
    fmt.Println("v164:", v164)
    fmt.Println("v166:", v166)
    fmt.Println("v178:", v178)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v163
    _ = v164
    _ = v165
    _ = v166
    _ = v178
    _ = v184
    _ = v2
    _ = v3
}
