package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 string; F1 string; F2 bool }
type GenS6 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 bool; F1 int }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int8; F1 bool }
type GenS12 struct { F0 uint8; F1 string; F2 int16 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint16; F1 int16; F2 uint }
type GenS15 struct { F0 int32; F1 int64; F2 int64 }
type GenS16 struct { F0 int32; F1 int32 }
func genFn0(p0 bool, p1 uint16) int16 {
    p0 = ((-1 <= 32768) && p0)
    v7 := GenS1{F0: 1}
    fmt.Println(p0)
    var v8 uint16 = (^(p1))
    var v9 uint32 = 701
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 32767
}

func main() {
    var v1 float32 = ((0.1 * ((((1.5 * 0.1) + 1e10) * 0.1) / 0.1)) + (((0.0 - -2.25) + (1e10 - 3.14159)) * 3.14159))
    v2 := make(map[uint32]uint32)
    v3 := []string{}
    v4 := []uint32{65535}
    v5 := []int{}
    var v6 string = ("" + "hello")
    v6 = (v6 + (v6 + ("hello" + ((v6 + "") + "\n\t"))))
    fmt.Println(v1, v1)
    v10 := GenS2{F0: (false && false), F1: 899}
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            fmt.Println(fi12, v1)
            _ = len(v5)
            v13 := GenS2{F0: (!false), F1: 144}
            _ = fi12
            _ = v13
        }
        var v14 float64 = 1e10
        delete(v2, 65535)
        _ = fi11
        _ = v14
    }
    if (!true) {
        var v15 string = (((v6 + (v6 + v6)) + (v6 + v6)) + v6)
        v1 = float32(0.0)
        if false {
            fmt.Println(v1, v15, v15)
            if false {
                v16 := GenS3{F0: 4294967295, F1: 42}
                _ = len(v5)
                _ = v16
            }
        }
        _ = v15
    } else {
        for fi17 := 0; fi17 < 18; fi17++ {
            switch "foo bar" {
            case "hello":
                var v18 uint64 = 42
                _ = v18
            case "foo bar":
                if (-2.25 <= ((0.1 + v1) * (3.14159 * 3.14159))) {
                    v19 := v2[2]
                    _ = v19
                } else {
                    v20 := GenS4{F0: 2, F1: -48}
                    for fi21 := 0; fi21 < 20; fi21++ {
                        switch v6 {
                        case "":
                            var v22 uint16 = 1
                            _ = v22
                        case "foo bar":
                            var v23 int16 = -1
                            _ = v23
                        case "a":
                            if false {
                                switch -318 {
                                case 1:
                                    v6 = v6
                                case 42:
                                    v24 := GenS5{F0: ((v6 + v6) + "a"), F1: "a", F2: true}
                                    _ = v24
                                case -1:
                                    fmt.Println(v1, v1)
                                default:
                                    fmt.Println(fi17)
                                }
                                v1 = v1
                                if (true && (0 != 128)) {
                                    switch fi21 {
                                    case 2147483647:
                                        fmt.Println(fi21, fi21)
                                    case 895:
                                        if ((false || ((("" + "hello") < "foo bar") && (4294967295 < 379))) && (152 == 188)) {
                                            v25 := GenS6{F0: 2, F1: 42, F2: 0}
                                            v2[2147483648] = 0
                                            _ = v25
                                        } else {
                                            switch v6 {
                                            case "\n\t":
                                                _ = len(v5)
                                            case "foo bar":
                                                switch (^(fi17)) {
                                                case 32767:
                                                    v26 := GenS1{F0: 127}
                                                    _ = v26
                                                case -139:
                                                    switch (fi17 / ((-2147483648) | 1)) {
                                                    case -386:
                                                        if (255 < 42) {
                                                            v2[909] = 722
                                                        } else {
                                                            if false {
                                                                v27 := GenS7{F0: 4294967295}
                                                                fmt.Println(v1, fi21, v6)
                                                                _ = v27
                                                            } else {
                                                                var v28 uint16 = 32767
                                                                if (48 < 42) {
                                                                    if true {
                                                                        fmt.Println(v1, v6)
                                                                        var v29 int = -135
                                                                        _ = v29
                                                                    }
                                                                    if true {
                                                                        delete(v2, 42)
                                                                    }
                                                                }
                                                                var v30 float64 = 1.5
                                                                _ = v28
                                                                _ = v30
                                                            }
                                                        }
                                                        switch (fi21 ^ fi17) {
                                                        case 127:
                                                            v1 = float32((float64(-2.25) + -2.25))
                                                        case 2147483647:
                                                            switch v6 {
                                                            case "a":
                                                                var v31 uint64 = 42
                                                                _ = v31
                                                            case "foo bar":
                                                                switch fi17 {
                                                                case 4294967295:
                                                                    v32 := GenS5{F0: ("foo bar" + (((v6 + "hello") + v6) + ((v6 + v6) + (v6 + v6)))), F1: "a", F2: (825 <= 0)}
                                                                    _ = v32
                                                                case 32767:
                                                                    if (660 < 42) {
                                                                        var v33 int8 = 1
                                                                        fmt.Println(fi21, fi17)
                                                                        fmt.Println(v33, fi21, v6)
                                                                        switch (fi21 / (((-(fi21))) | 1)) {
                                                                        case 9223372036854775807:
                                                                            if (664 >= 1) {
                                                                                if true {
                                                                                    var v34 int = -9223372036854775808
                                                                                    switch "" {
                                                                                    case "foo bar":
                                                                                        v35 := GenS5{F0: v6, F1: v6, F2: true}
                                                                                        _ = v35
                                                                                    case "":
                                                                                        fmt.Println(fi21, v1, fi17)
                                                                                    case "hello":
                                                                                        var v36 int64 = -1
                                                                                        _ = v36
                                                                                    default:
                                                                                        fmt.Println(v1, v33)
                                                                                    }
                                                                                    if (65535 > 0) {
                                                                                        _ = len(v5)
                                                                                        v1 = ((-2.25 / 1.5) / (1e10 * v1))
                                                                                        fmt.Println(v6)
                                                                                    }
                                                                                    _ = v34
                                                                                } else {
                                                                                    if false {
                                                                                        _ = len(v2)
                                                                                        if true {
                                                                                            var v37 int8 = (^(v33))
                                                                                            if true {
                                                                                                var v38 uint32 = 2147483648
                                                                                                v39 := GenS5{F0: v6, F1: v6, F2: (1 <= 1)}
                                                                                                var v40 int8 = v33
                                                                                                delete(v2, v38)
                                                                                                _ = v38
                                                                                                _ = v39
                                                                                                _ = v40
                                                                                            } else {
                                                                                                fmt.Println(v6)
                                                                                            }
                                                                                            _ = v37
                                                                                        }
                                                                                        v6 = "foo bar"
                                                                                        if false {
                                                                                            v1 -= 1e10
                                                                                        } else {
                                                                                            v6 = (v6 + "hello")
                                                                                            if len(v5) >= 2 {
                                                                                                v5 = v5[1:len(v5)]
                                                                                            }
                                                                                            v41 := GenS8{F0: v33}
                                                                                            _ = v41
                                                                                        }
                                                                                        v42 := v2[2147483648]
                                                                                        _ = v42
                                                                                    } else {
                                                                                        var v43 int8 = v33
                                                                                        _ = v43
                                                                                    }
                                                                                    if true {
                                                                                        v1 += v1
                                                                                        fmt.Println(v33)
                                                                                        v2[25] = 81
                                                                                    } else {
                                                                                        v44 := GenS9{F0: (true || (!false)), F1: -957}
                                                                                        _ = v44
                                                                                    }
                                                                                }
                                                                                var v45 float64 = 0.1
                                                                                var v46 int8 = 0
                                                                                _ = v45
                                                                                _ = v46
                                                                            } else {
                                                                                v6 = v6
                                                                                switch (^(fi17)) {
                                                                                case 2147483647:
                                                                                    v33 = v33
                                                                                default:
                                                                                    var v47 uint = 890
                                                                                    _ = v47
                                                                                }
                                                                                v48 := GenS9{F0: true, F1: 0}
                                                                                _ = v48
                                                                            }
                                                                        default:
                                                                            v49 := GenS10{F0: -1}
                                                                            v2[0] = 368
                                                                            v50 := GenS6{F0: 42, F1: 598, F2: 255}
                                                                            _ = v49
                                                                            _ = v50
                                                                        }
                                                                        v5 = append(v5, (^(fi17)))
                                                                        switch (fi17 & fi21) {
                                                                        case -940:
                                                                            var v51 uint = 4294967295
                                                                            v1 = 0.0
                                                                            var v52 float32 = (v1 * v1)
                                                                            var v53 int64 = 854
                                                                            _ = v51
                                                                            _ = v52
                                                                            _ = v53
                                                                        case 9223372036854775807:
                                                                            v54 := GenS7{F0: 617}
                                                                            _ = v54
                                                                        case 2147483647:
                                                                            delete(v2, 551)
                                                                        default:
                                                                            var v55 string = v6
                                                                            v55 = "foo bar"
                                                                            v1 = ((v1 - 1e10) + ((0.1 - v1) / v1))
                                                                            _ = v55
                                                                        }
                                                                        v56 := GenS11{F0: v33, F1: (4294967295 >= 1)}
                                                                        var v57 uint64 = 4294967295
                                                                        _ = v33
                                                                        _ = v56
                                                                        _ = v57
                                                                    } else {
                                                                        if (1 == 73) {
                                                                            switch fi17 {
                                                                            case 555:
                                                                                delete(v2, 773)
                                                                            case -683:
                                                                                fmt.Println(fi21, fi17)
                                                                            case 32768:
                                                                                v6 = ""
                                                                            default:
                                                                                fmt.Println(v6, fi17, fi21)
                                                                            }
                                                                            delete(v2, 1)
                                                                        } else {
                                                                            if true {
                                                                                var v58 uint8 = 42
                                                                                _ = v58
                                                                            }
                                                                            var v59 uint = 18446744073709551615
                                                                            switch (v6 + (v6 + ((v6 + (v6 + "\n\t")) + v6))) {
                                                                            case "":
                                                                                var v60 uint64 = 2
                                                                                _ = v60
                                                                            case "foo bar":
                                                                                var v61 float32 = (v1 - v1)
                                                                                _ = v61
                                                                            }
                                                                            v59 -= (v59 + v59)
                                                                            _ = v59
                                                                        }
                                                                    }
                                                                default:
                                                                    v62 := GenS12{F0: 1, F1: "\n\t", F2: 0}
                                                                    _ = v62
                                                                }
                                                            case "":
                                                                var v63 float32 = v1
                                                                _ = v63
                                                            default:
                                                                v1 += v1
                                                            }
                                                        case -9223372036854775808:
                                                            var v64 uint16 = 439
                                                            _ = v64
                                                        }
                                                    case 1:
                                                        v65 := GenS13{F0: 42}
                                                        _ = v65
                                                    case 535:
                                                        switch (v6 + v6) {
                                                        case "":
                                                            v66 := GenS3{F0: 1, F1: 2}
                                                            _ = v66
                                                        case "\n\t":
                                                            var v67 int = fi17
                                                            var v68 uint64 = 971
                                                            _ = v67
                                                            _ = v68
                                                        default:
                                                            switch "\n\t" {
                                                            case "a":
                                                                var v69 uint32 = 913
                                                                _ = v69
                                                            case "":
                                                                if false {
                                                                    var v70 uint64 = 912
                                                                    var v71 bool = false
                                                                    _ = v70
                                                                    _ = v71
                                                                } else {
                                                                    var v72 int16 = -1
                                                                    fmt.Println(v72, v6, v1)
                                                                    _ = v72
                                                                }
                                                            case "hello":
                                                                v1 = v1
                                                            }
                                                        }
                                                    }
                                                case 2147483647:
                                                    v6 = "hello"
                                                }
                                            case "hello":
                                                fmt.Println(fi17, fi21)
                                            }
                                            var v73 string = v6
                                            var v74 int64 = 0
                                            _ = v73
                                            _ = v74
                                        }
                                    }
                                    switch 859 {
                                    case -9223372036854775808:
                                        if (42 >= -128) {
                                            var v75 uint8 = 87
                                            var v76 int8 = -62
                                            v75 += v75
                                            _ = v75
                                            _ = v76
                                        }
                                        switch (^(fi21)) {
                                        case -370:
                                            fmt.Println(fi21, v6, v6)
                                        case -2147483648:
                                            var v77 int32 = 1
                                            switch v6 {
                                            case "foo bar":
                                                v6 = (v6 + v6)
                                            case "":
                                                var v78 int16 = genFn0(true, 255)
                                                _ = v78
                                            case "hello":
                                                if (9223372036854775807 >= -121) {
                                                    fmt.Println(fi21)
                                                }
                                            default:
                                                v79 := GenS14{F0: 32767, F1: 32767, F2: 0}
                                                _ = v79
                                            }
                                            var v80 uint = 255
                                            _ = v77
                                            _ = v80
                                        default:
                                            v81 := GenS2{F0: (333 != fi21), F1: 42}
                                            _ = v81
                                        }
                                        if true {
                                            if len(v5) > 0 {
                                                v5[0] = (fi17 ^ 9223372036854775807)
                                            }
                                        } else {
                                            var v82 uint16 = 375
                                            _ = v82
                                        }
                                        v1 -= (3.14159 - (((v1 + ((v1 + float32(((3.14159 - (3.14159 - 0.0)) / 0.1))) - float32(float64(1e10)))) - v1) * v1))
                                    case 1:
                                        switch v6 {
                                        case "hello":
                                            var v83 uint16 = 32767
                                            _ = v83
                                        }
                                    case 32768:
                                        v84 := GenS13{F0: 1}
                                        _ = v84
                                    default:
                                        if false {
                                            v1 = v1
                                            var v85 uint16 = 249
                                            _ = v85
                                        } else {
                                            if true {
                                                if len(v5) > 0 {
                                                    v5[0] = (-(fi17))
                                                }
                                                if true {
                                                    v1 -= 1e10
                                                    v6 = v6
                                                    v3 = append(v3, "hello")
                                                } else {
                                                    if false {
                                                        var v86 uint = 787
                                                        _ = v86
                                                    } else {
                                                        if ((919 < 9223372036854775807) || false) {
                                                            switch v6 {
                                                            case "a":
                                                                fmt.Println(fi21, fi21, fi21)
                                                            default:
                                                                switch v6 {
                                                                case "foo bar":
                                                                    v1 += (v1 * (v1 * (v1 + v1)))
                                                                }
                                                            }
                                                            var v87 uint16 = 1
                                                            v88 := GenS14{F0: (v87 * (v87 >> ((18446744073709551615) & 63))), F1: 42, F2: 428}
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                        v6 = v6
                                                    }
                                                }
                                            }
                                            v1 -= (v1 * ((v1 - (float32((-2.25 - (0.0 + -2.25))) + float32(3.14159))) * v1))
                                            v6 = v6
                                        }
                                    }
                                }
                            } else {
                                var v89 int8 = 0
                                v6 = v6
                                _ = v89
                            }
                            if false {
                                if (float32((1e10 * 1e10)) == v1) {
                                    v90 := GenS5{F0: "\n\t", F1: v6, F2: (!true)}
                                    var v91 int8 = -128
                                    _ = v90
                                    _ = v91
                                }
                            } else {
                                var v92 uint16 = 238
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v92
                            }
                        }
                        _ = fi21
                    }
                    v6 = (((v6 + (v6 + "a")) + v6) + v6)
                    _ = v20
                }
            case "\n\t":
                v1 -= v1
            }
            for fi93 := 0; fi93 < 17; fi93++ {
                if (true && ((334 < -9223372036854775808) && (-2147483648 == 0))) {
                    var v94 float64 = 0.1
                    _ = v94
                }
                if (!false) {
                    v1 = ((0.1 / v1) / v1)
                    if (0.1 < v1) {
                        fmt.Println(fi17)
                        var v95 uint64 = 795
                        _ = v95
                    }
                    v6 = v6
                }
                v96 := GenS4{F0: 2, F1: 100}
                _ = fi93
                _ = v96
            }
            _ = fi17
        }
        for fi97 := 0; fi97 < 14; fi97++ {
            _ = len(v5)
            v1 = 0.0
            _ = fi97
        }
        var v98 uint64 = 329
        _ = v98
    }
    delete(v2, 42)
    var v99 uint8 = 1
    var v100 string = ("a" + v6)
    switch 506 {
    case -211:
        v100 = (("hello" + v6) + "foo bar")
    case 128:
        v2[42] = 2147483648
    case 2147483647:
        for fi101 := 0; fi101 < 19; fi101++ {
            var v102 bool = false
            v103 := v2[830]
            _ = fi101
            _ = v102
            _ = v103
        }
    }
    if true {
        v104 := GenS15{F0: 332, F1: 9223372036854775807, F2: -9223372036854775808}
        if (false && false) {
            var v105 string = ""
            _ = v105
        } else {
            switch 42 {
            case -9223372036854775808:
                fmt.Println(v1, v99, v1)
            case 128:
                v1 += 0.0
            }
            var v106 string = v6
            _ = v106
        }
        for fi107 := 0; fi107 < 2; fi107++ {
            v108 := v2[1]
            if (1 != (v108 >> ((401) & 63))) {
                var v109 string = ((v6 + v6) + v6)
                var v110 int = fi107
                if true {
                    _ = len(v3)
                    for fi111 := 0; fi111 < 13; fi111++ {
                        var v112 uint16 = 42
                        _ = fi111
                        _ = v112
                    }
                }
                _ = v109
                _ = v110
            } else {
                for fi113 := 0; fi113 < 7; fi113++ {
                    v6 = (v6 + ("foo bar" + (((v6 + v6) + v100) + v100)))
                    _ = fi113
                }
                for fi114 := 0; fi114 < 17; fi114++ {
                    v6 = v6
                    if ((((!((!false) || true)) || (!false)) || false) && (100 <= -1)) {
                        v115 := GenS6{F0: 42, F1: (v108 + 357), F2: uint16(v99)}
                        v116 := v2[453]
                        _ = v115
                        _ = v116
                    } else {
                        if false {
                            var v117 int = (fi114 / ((fi114) | 1))
                            if true {
                                switch (fi114 >> ((151) & 63)) {
                                case 9223372036854775807:
                                    v118 := GenS7{F0: 18446744073709551615}
                                    _ = v118
                                case 127:
                                    v3 = append(v3, v6)
                                    v117 += v117
                                    var v119 int64 = -477
                                    _ = v119
                                }
                                delete(v2, 2147483648)
                                switch fi114 {
                                case 32767:
                                    fmt.Println(fi114)
                                case 9223372036854775807:
                                    fmt.Println(v108, v108, v108)
                                }
                            } else {
                                var v120 uint16 = uint16(v99)
                                _ = v120
                            }
                            switch fi114 {
                            case 42:
                                var v121 uint = 18446744073709551615
                                _ = v121
                            default:
                                v3 = append(v3, v6)
                            }
                            v1 += 1e10
                            _ = v117
                        }
                    }
                    var v122 float64 = (-2.25 * ((0.1 + 0.0) / -2.25))
                    _ = fi114
                    _ = v122
                }
            }
            var v123 uint16 = 65535
            _ = fi107
            _ = v108
            _ = v123
        }
        _ = v104
    } else {
        delete(v2, 2)
        v124 := GenS16{F0: 5, F1: 2147483647}
        fmt.Println(v1)
        _ = v124
    }
    var v125 uint = 1
    v125 = v125
    var v126 uint32 = 0
    var v127 uint64 = 602
    fmt.Println(v125, v100, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v125:", v125)
    fmt.Println("v126:", v126)
    fmt.Println("v127:", v127)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v10
    _ = v100
    _ = v125
    _ = v126
    _ = v127
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v99
}
