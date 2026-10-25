package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 bool; F1 uint8; F2 int32 }
type GenS10 struct { F0 int16; F1 int16; F2 uint32 }
type GenS11 struct { F0 uint; F1 bool }
type GenS12 struct { F0 int }
func genFn0(p0 uint32, p1 int64, p2 int) int64 {
    v4 := GenS1{F0: 1000}
    p2 += (p2 + 1)
    switch (("foo bar" + "hello") + "") {
    case "":
        if (0 >= 42) {
            p1 -= 42
        }
    case "hello":
        if (42 >= 100) {
            p0 = (p0 + p0)
            var v5 string = "\n\t"
            _ = v5
        } else {
            v6 := GenS2{F0: 4294967295}
            for fi7 := 0; fi7 < 2; fi7++ {
                v8 := []int{}
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 19; fi9++ {
                p2 = p2
                _ = fi9
            }
            _ = v6
        }
    case "a":
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := []int{(^(fi10)), (fi10 / ((-982) | 1)), fi10, (fi10 ^ fi10)}
            _ = fi10
            _ = v11
        }
    default:
        p1 -= p1
    }
    var v12 float64 = 0.0
    p2 -= (p2 / ((p2) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v4
    return 0
}
func genFn1(p0 uint, p1 int) int32 {
    p1 += p1
    var v13 float64 = (((0.0 * float64(1.5)) + 3.14159) / float64((3.14159 + (float32((-2.25 * (float64((((1.5 / -2.25) + 3.14159) - 3.14159)) - 3.14159))) - 1.5))))
    _ = p0
    _ = p1
    _ = v13
    return 42
}

func main() {
    var v1 int8 = 63
    var v2 bool = (false && true)
    var v3 string = ("foo bar" + ("" + "foo bar"))
    switch "foo bar" {
    case "":
        if ((3.14159 + 0.0) != 0.0) {
            if v2 {
                v14 := []int{-1, 127, 42, 32768}
                _ = v14
            } else {
                if (578 != 4294967295) {
                    fmt.Println(v1, v2, v3)
                } else {
                    switch (v3 + (v3 + (((("a" + v3) + "\n\t") + (v3 + (v3 + ("foo bar" + (((v3 + (("\n\t" + v3) + ("\n\t" + v3))) + (v3 + "hello")) + (v3 + (v3 + "foo bar"))))))) + v3))) {
                    case "foo bar":
                        var v15 uint8 = 199
                        _ = v15
                    case "a":
                        switch 128 {
                        case 9223372036854775807:
                            v16 := GenS2{F0: 376}
                            _ = v16
                        case -96:
                            for fi17 := 0; fi17 < 17; fi17++ {
                                v18 := []int{fi17, (fi17 >> ((4294967295) & 63)), fi17, (fi17 >> ((4294967295) & 63))}
                                if ("foo bar" >= (("a" + "hello") + v3)) {
                                    if (-2.25 != 0.0) {
                                        if true {
                                            v1 -= 98
                                            if v2 {
                                                fmt.Println(v2)
                                                for fi19 := 0; fi19 < 9; fi19++ {
                                                    var v20 uint16 = 42
                                                    var v21 uint = 296
                                                    _ = fi19
                                                    _ = v20
                                                    _ = v21
                                                }
                                                v3 = (v3 + (v3 + "\n\t"))
                                            } else {
                                                switch -2147483648 {
                                                case 42:
                                                    v1 = (-(v1))
                                                case 128:
                                                    v2 = true
                                                case 0:
                                                    if v2 {
                                                        if (18446744073709551615 == 2) {
                                                            if true {
                                                                v22 := GenS3{F0: 2147483647, F1: (9223372036854775807 < 42)}
                                                                fmt.Println(v1, fi17)
                                                                _ = v22
                                                            } else {
                                                                v2 = true
                                                            }
                                                            switch v3 {
                                                            case "\n\t":
                                                                fmt.Println(v1)
                                                            case "foo bar":
                                                                if (!(2147483647 >= -738)) {
                                                                    v2 = (3.14159 <= 0.0)
                                                                }
                                                            case "a":
                                                                if false {
                                                                    _ = len(v18)
                                                                }
                                                            default:
                                                                for fi23 := 0; fi23 < 20; fi23++ {
                                                                    v2 = false
                                                                    var v24 float64 = 3.14159
                                                                    if (0.0 != 0.0) {
                                                                        switch (-(fi23)) {
                                                                        case 128:
                                                                            v25 := GenS4{F0: fi23}
                                                                            _ = v25
                                                                        case -9223372036854775808:
                                                                            if true {
                                                                                v1 = (v1 ^ (v1 - (^(v1))))
                                                                            } else {
                                                                                var v26 bool = true
                                                                                _ = v26
                                                                            }
                                                                        default:
                                                                            fmt.Println(v3, v1, fi17)
                                                                        }
                                                                        if ((65535 <= 1) || (v2 && v2)) {
                                                                            var v27 uint32 = 4294967295
                                                                            fmt.Println(v3)
                                                                            v1 = v1
                                                                            _ = v27
                                                                        }
                                                                    }
                                                                    _ = fi23
                                                                    _ = v24
                                                                }
                                                            }
                                                            if (fi17 < 127) {
                                                                if (false && (0.0 >= float32((float64((1.5 - 1e10)) / (((3.14159 - (float64(0.0) - 1e10)) + 1.5) - 0.0))))) {
                                                                    switch fi17 {
                                                                    case 42:
                                                                        for fi28 := 0; fi28 < 19; fi28++ {
                                                                            v2 = (!false)
                                                                            var v29 uint = 527
                                                                            var v30 float64 = 0.0
                                                                            _ = fi28
                                                                            _ = v29
                                                                            _ = v30
                                                                        }
                                                                    }
                                                                    v2 = false
                                                                    v2 = (!false)
                                                                } else {
                                                                    v31 := GenS5{F0: v1, F1: v2}
                                                                    v1 = (v1 + 127)
                                                                    _ = v31
                                                                }
                                                                for fi32 := 0; fi32 < 7; fi32++ {
                                                                    v2 = v2
                                                                    if v2 {
                                                                        var v33 uint = 857
                                                                        v34 := GenS1{F0: int16(fi32)}
                                                                        if v2 {
                                                                            fmt.Println(fi32, v3)
                                                                            v35 := GenS4{F0: fi32}
                                                                            var v36 int = (fi17 % ((407) | 1))
                                                                            if (255 != 42) {
                                                                                v37 := GenS6{F0: true}
                                                                                _ = v37
                                                                            }
                                                                            _ = v35
                                                                            _ = v36
                                                                        } else {
                                                                            v38 := GenS3{F0: 2147483648, F1: v2}
                                                                            if len(v18) > 0 {
                                                                                v18[0] = (fi32 << (((v33 % (((v33 << ((2) & 63))) | 1))) & 63))
                                                                            }
                                                                            v39 := GenS7{F0: -32}
                                                                            _ = v38
                                                                            _ = v39
                                                                        }
                                                                        _ = v33
                                                                        _ = v34
                                                                    }
                                                                    _ = fi32
                                                                }
                                                            }
                                                        } else {
                                                            v40 := GenS2{F0: 2}
                                                            _ = len(v18)
                                                            switch fi17 {
                                                            case -9223372036854775808:
                                                                var v41 uint8 = 1
                                                                _ = v41
                                                            default:
                                                                switch fi17 {
                                                                case 32768:
                                                                    if ((((1 <= 562) || false) && false) && true) {
                                                                        if true {
                                                                            fmt.Println(v2, v2, v1)
                                                                            var v42 int16 = -1
                                                                            _ = v42
                                                                        } else {
                                                                            if true {
                                                                                var v43 uint = 156
                                                                                if len(v18) >= 2 {
                                                                                    v18 = v18[1:len(v18)]
                                                                                }
                                                                                _ = v43
                                                                            }
                                                                            for fi44 := 0; fi44 < 8; fi44++ {
                                                                                v45 := GenS4{F0: fi17}
                                                                                v2 = (!(502 == 42))
                                                                                if len(v18) >= 2 {
                                                                                    v18 = v18[1:len(v18)]
                                                                                }
                                                                                _ = len(v18)
                                                                                _ = fi44
                                                                                _ = v45
                                                                            }
                                                                            v1 += (v1 * (v1 - int8(fi17)))
                                                                        }
                                                                        for fi46 := 0; fi46 < 9; fi46++ {
                                                                            v1 += v1
                                                                            if v2 {
                                                                                var v47 int16 = -107
                                                                                fmt.Println(v47, v47, v2)
                                                                                _ = v47
                                                                            } else {
                                                                                if true {
                                                                                    fmt.Println(v1)
                                                                                    if v2 {
                                                                                        if v2 {
                                                                                            if v2 {
                                                                                                var v48 float64 = 0.1
                                                                                                var v49 uint = 65535
                                                                                                _ = v48
                                                                                                _ = v49
                                                                                            } else {
                                                                                                v3 = v3
                                                                                            }
                                                                                            if len(v18) >= 2 {
                                                                                                v18 = v18[1:len(v18)]
                                                                                            }
                                                                                        }
                                                                                        v50 := GenS5{F0: (v1 - 1), F1: (genFn1(1, 32767) != 2147483647)}
                                                                                        _ = v50
                                                                                    }
                                                                                    if false {
                                                                                        if len(v18) >= 2 {
                                                                                            v18 = v18[1:len(v18)]
                                                                                        }
                                                                                        var v51 bool = (42 < -209)
                                                                                        _ = v51
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(fi17)
                                                                                    if ((v1 >> ((520) & 63)) > v1) {
                                                                                        if len(v18) > 0 {
                                                                                            v18[0] = fi17
                                                                                        }
                                                                                    }
                                                                                    if len(v18) > 0 {
                                                                                        v18[0] = fi17
                                                                                    }
                                                                                }
                                                                                var v52 float32 = -2.25
                                                                                v52 += ((0.0 / (3.14159 + (v52 * (v52 - -2.25)))) * float32(3.14159))
                                                                                var v53 int16 = 716
                                                                                _ = v52
                                                                                _ = v53
                                                                            }
                                                                            v2 = ((fi17 << ((255) & 63)) < -9223372036854775808)
                                                                            v1 += (v1 >> ((255) & 63))
                                                                            _ = fi46
                                                                        }
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            v3 = v3
                                                                        default:
                                                                            var v54 uint = 940
                                                                            _ = v54
                                                                        }
                                                                    } else {
                                                                        fmt.Println(fi17, v3, v1)
                                                                        for fi55 := 0; fi55 < 2; fi55++ {
                                                                            fmt.Println(v1, v1)
                                                                            _ = fi55
                                                                        }
                                                                    }
                                                                case -845:
                                                                    fmt.Println(fi17, v2, v2)
                                                                case 0:
                                                                    if len(v18) > 0 {
                                                                        v18[0] = (fi17 | (fi17 % (((fi17 >> ((255) & 63))) | 1)))
                                                                    }
                                                                default:
                                                                    fmt.Println(v3, v3, v1)
                                                                }
                                                            }
                                                            _ = v40
                                                        }
                                                    } else {
                                                        v1 = (v1 << ((1) & 63))
                                                    }
                                                }
                                                if len(v18) >= 2 {
                                                    v18 = v18[1:len(v18)]
                                                }
                                            }
                                            var v56 uint64 = 42
                                            _ = v56
                                        } else {
                                            var v57 uint32 = 42
                                            if true {
                                                v57 = v57
                                                v58 := GenS1{F0: 303}
                                                v59 := GenS5{F0: 42, F1: (v2 && false)}
                                                _ = v58
                                                _ = v59
                                            } else {
                                                v57 = v57
                                                v60 := GenS8{F0: v3}
                                                v57 -= v57
                                                _ = v60
                                            }
                                            v2 = (!v2)
                                            _ = v57
                                        }
                                        if (104 < 822) {
                                            for fi61 := 0; fi61 < 3; fi61++ {
                                                _ = len(v18)
                                                if (!v2) {
                                                    v18 = append(v18, (fi17 & fi61))
                                                }
                                                _ = fi61
                                            }
                                        } else {
                                            v62 := GenS3{F0: 2147483647, F1: v2}
                                            v2 = (false && false)
                                            _ = v62
                                        }
                                        v3 = v3
                                    } else {
                                        var v63 uint16 = 807
                                        switch 0 {
                                        case 2147483647:
                                            v3 = ((v3 + "foo bar") + v3)
                                        case 42:
                                            if v2 {
                                                for fi64 := 0; fi64 < 10; fi64++ {
                                                    v63 = (^(v63))
                                                    if false {
                                                        var v65 string = "\n\t"
                                                        fmt.Println(v65)
                                                        var v66 uint64 = 0
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                    _ = fi64
                                                }
                                            } else {
                                                switch v3 {
                                                case "\n\t":
                                                    fmt.Println(v3)
                                                case "hello":
                                                    var v67 int8 = v1
                                                    _ = v67
                                                case "foo bar":
                                                    var v68 int = fi17
                                                    _ = v68
                                                default:
                                                    if false {
                                                        if len(v18) > 0 {
                                                            v18[0] = 518
                                                        }
                                                        for fi69 := 0; fi69 < 8; fi69++ {
                                                            var v70 uint32 = 2
                                                            if len(v18) > 0 {
                                                                v18[0] = fi69
                                                            }
                                                            _ = fi69
                                                            _ = v70
                                                        }
                                                    } else {
                                                        var v71 int32 = -645
                                                        _ = v71
                                                    }
                                                }
                                            }
                                        case 198:
                                            for fi72 := 0; fi72 < 12; fi72++ {
                                                var v73 uint16 = (v63 / ((32767) | 1))
                                                _ = fi72
                                                _ = v73
                                            }
                                        }
                                        if (("" + "\n\t") != v3) {
                                            if v2 {
                                                if (v63 > v63) {
                                                    v74 := GenS8{F0: ""}
                                                    _ = v74
                                                }
                                                if v2 {
                                                    v2 = true
                                                    var v75 int8 = (^(v1))
                                                    _ = v75
                                                } else {
                                                    v76 := GenS4{F0: fi17}
                                                    var v77 uint32 = 0
                                                    _ = v76
                                                    _ = v77
                                                }
                                                switch v3 {
                                                case "\n\t":
                                                    v78 := GenS4{F0: fi17}
                                                    _ = v78
                                                case "hello":
                                                    v18 = append(v18, -982)
                                                default:
                                                    if (((5 < 223) && false) && v2) {
                                                        v63 = (v63 >> ((454) & 63))
                                                    } else {
                                                        if len(v18) >= 2 {
                                                            v18 = v18[1:len(v18)]
                                                        }
                                                        fmt.Println(v1, v2, v3)
                                                    }
                                                }
                                            }
                                            var v79 uint16 = v63
                                            _ = v79
                                        }
                                        _ = v63
                                    }
                                    var v80 int64 = -651
                                    _ = v80
                                } else {
                                    var v81 float64 = 1e10
                                    var v82 float64 = (v81 - 0.0)
                                    for fi83 := 0; fi83 < 15; fi83++ {
                                        switch ("hello" + v3) {
                                        case "\n\t":
                                            v3 = ""
                                        case "foo bar":
                                            switch fi83 {
                                            case 1:
                                                v81 = v82
                                            case 127:
                                                v3 = "hello"
                                            default:
                                                switch fi83 {
                                                case -528:
                                                    switch "a" {
                                                    case "":
                                                        var v84 int32 = 2147483647
                                                        _ = v84
                                                    default:
                                                        var v85 uint32 = 888
                                                        _ = v85
                                                    }
                                                default:
                                                    var v86 int8 = (v1 << ((0) & 63))
                                                    _ = v86
                                                }
                                            }
                                        default:
                                            var v87 float64 = (-2.25 * (0.1 + (0.0 + v82)))
                                            _ = v87
                                        }
                                        _ = fi83
                                    }
                                    _ = v81
                                    _ = v82
                                }
                                _ = fi17
                                _ = v18
                            }
                        }
                    default:
                        var v88 int64 = 1
                        _ = v88
                    }
                    v3 = ("\n\t" + ("foo bar" + "a"))
                    v89 := GenS2{F0: 1}
                    _ = v89
                }
                fmt.Println(v2)
                v3 = "\n\t"
            }
            for fi90 := 0; fi90 < 4; fi90++ {
                v3 = "hello"
                for fi91 := 0; fi91 < 1; fi91++ {
                    v92 := GenS9{F0: (-48 != v1), F1: 0, F2: -820}
                    _ = fi91
                    _ = v92
                }
                _ = fi90
            }
        } else {
            var v93 uint8 = 127
            if false {
                for fi94 := 0; fi94 < 16; fi94++ {
                    var v95 uint8 = v93
                    fmt.Println(fi94)
                    _ = fi94
                    _ = v95
                }
                v3 = (((v3 + v3) + ("" + "a")) + "foo bar")
                v96 := GenS10{F0: -1, F1: 42, F2: 23}
                _ = v96
            } else {
                var v97 int16 = -828
                _ = v97
            }
            if (!((v93 ^ (v93 - v93)) != (v93 % ((v93) | 1)))) {
                v98 := GenS11{F0: 343, F1: (2147483647 == 1)}
                for fi99 := 0; fi99 < 19; fi99++ {
                    var v100 string = "hello"
                    if (18446744073709551615 > 255) {
                        v101 := []int{fi99, fi99, (^(fi99))}
                        var v102 int = 0
                        fmt.Println(v102, v3)
                        _ = v101
                        _ = v102
                    }
                    v103 := []int{(fi99 ^ fi99), fi99}
                    _ = fi99
                    _ = v100
                    _ = v103
                }
                _ = v98
            } else {
                var v104 uint = 42
                v104 = v104
                var v105 int32 = -465
                _ = v104
                _ = v105
            }
            _ = v93
        }
    case "foo bar":
        if (!(255 >= 2)) {
            var v106 uint64 = 865
            if (uint16(v106) <= 379) {
                if (false && v2) {
                    for fi107 := 0; fi107 < 11; fi107++ {
                        switch (fi107 | 2147483647) {
                        case -9223372036854775808:
                            switch v3 {
                            case "a":
                                var v108 int32 = -1
                                _ = v108
                            default:
                                switch (-(fi107)) {
                                case 1:
                                    for fi109 := 0; fi109 < 15; fi109++ {
                                        v106 = v106
                                        v110 := []int{fi109}
                                        _ = fi109
                                        _ = v110
                                    }
                                }
                            }
                        case 2147483647:
                            for fi111 := 0; fi111 < 18; fi111++ {
                                v3 = ""
                                fmt.Println(fi111)
                                v106 = (v106 - v106)
                                _ = fi111
                            }
                        case 9223372036854775807:
                            var v112 uint64 = 4294967295
                            _ = v112
                        }
                        _ = fi107
                    }
                    switch "foo bar" {
                    case "a":
                        fmt.Println(v3, v3, v1)
                    case "":
                        for fi113 := 0; fi113 < 4; fi113++ {
                            v114 := []int{}
                            v115 := GenS11{F0: 687, F1: v2}
                            _ = fi113
                            _ = v114
                            _ = v115
                        }
                    case "foo bar":
                        v106 -= 0
                    }
                    v1 -= v1
                }
                switch v3 {
                case "foo bar":
                    var v116 int64 = -1
                    _ = v116
                }
                switch v3 {
                case "foo bar":
                    var v117 int32 = 42
                    _ = v117
                }
            }
            v118 := GenS3{F0: 65535, F1: (!v2)}
            _ = v106
            _ = v118
        } else {
            var v119 uint32 = 42
            _ = v119
        }
    }
    v2 = v2
    var v120 int16 = -549
    v2 = ((!v2) && (2 >= 43))
    v121 := []int{128}
    var v122 int64 = genFn0(883, genFn0(1, 42, -1), 1)
    v121 = append(v121, -2147483648)
    var v123 string = "a"
    fmt.Println(v120, v123, v120)
    if v2 {
        v124 := GenS2{F0: 859}
        var v125 uint64 = 42
        _ = len(v121)
        _ = v124
        _ = v125
    } else {
        fmt.Println(v122)
        switch 32768 {
        case 1:
            _ = len(v121)
        case -1:
            v122 = (^(v122))
        case 32767:
            if (42 == 2) {
                v123 = (v123 + "a")
                for fi126 := 0; fi126 < 5; fi126++ {
                    if true {
                        v122 = genFn0(42, genFn0(1, (v122 + -9223372036854775808), 2147483647), (fi126 - -2147483648))
                    } else {
                        v127 := GenS1{F0: v120}
                        v122 += (v122 | (v122 % ((-1) | 1)))
                        _ = v127
                    }
                    if (-32768 == 625) {
                        for fi128 := 0; fi128 < 2; fi128++ {
                            var v129 float64 = 0.1
                            _ = fi128
                            _ = v129
                        }
                    } else {
                        if (fi126 > (fi126 % (((-(fi126))) | 1))) {
                            v1 = -1
                        } else {
                            if (0.1 > 1.5) {
                                if (255 != 255) {
                                    if (v2 || (v2 && ((v1 ^ (v1 | (v1 ^ int8(fi126)))) != -128))) {
                                        _ = len(v121)
                                    }
                                    v123 = v3
                                } else {
                                    v120 -= -1
                                    v122 = (v122 + genFn0(343, v122, fi126))
                                }
                                var v130 int16 = 0
                                for fi131 := 0; fi131 < 7; fi131++ {
                                    v1 -= (v1 + -115)
                                    switch (fi126 >> ((371) & 63)) {
                                    case 128:
                                        var v132 uint64 = 2
                                        _ = v132
                                    }
                                    var v133 float64 = 0.0
                                    v120 -= 1000
                                    _ = fi131
                                    _ = v133
                                }
                                _ = v130
                            }
                            for fi134 := 0; fi134 < 17; fi134++ {
                                v135 := GenS12{F0: fi134}
                                v3 = v3
                                v121 = append(v121, (^(fi134)))
                                v123 = ("a" + ("" + (v3 + v3)))
                                _ = fi134
                                _ = v135
                            }
                            fmt.Println(v3, v2, fi126)
                        }
                        var v136 uint32 = 691
                        if v2 {
                            switch v3 {
                            case "hello":
                                v137 := GenS4{F0: 9223372036854775807}
                                _ = v137
                            default:
                                v3 = v123
                            }
                            var v138 int64 = 1
                            v120 = v120
                            _ = v138
                        } else {
                            _ = len(v121)
                            v139 := GenS11{F0: 453, F1: v2}
                            _ = v139
                        }
                        _ = v136
                    }
                    _ = fi126
                }
                v120 = int16(v122)
            } else {
                v140 := GenS8{F0: ((v3 + v3) + (((v123 + "\n\t") + "hello") + ((v3 + v123) + v3)))}
                _ = v140
            }
        default:
            var v141 float32 = 1e10
            _ = v141
        }
    }
    var v142 string = v3
    var v143 int32 = 1
    v144 := GenS10{F0: (v120 & v120), F1: 1, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v120:", v120)
    fmt.Println("slicelen v121:", len(v121), v121)
    fmt.Println("v122:", v122)
    fmt.Println("v123:", v123)
    fmt.Println("v142:", v142)
    fmt.Println("v143:", v143)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v120
    _ = v121
    _ = v122
    _ = v123
    _ = v142
    _ = v143
    _ = v144
    _ = v2
    _ = v3
}
