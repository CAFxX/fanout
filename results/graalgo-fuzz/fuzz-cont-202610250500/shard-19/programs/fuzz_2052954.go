package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int16 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS5 struct { F0 int; F1 int16; F2 uint8 }
type GenS6 struct { F0 int; F1 int }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS9 struct { F0 uint16; F1 uint64 }
type GenS10 struct { F0 int32; F1 uint16 }
type GenS11 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS12 struct { F0 uint64; F1 bool }
type GenS13 struct { F0 uint16 }
func genFn0() uint32 {
    if (189 != 128) {
        switch 898 {
        case 0:
            var v7 bool = (42 != 261)
            _ = v7
        case -2147483648:
            if (42 != -9223372036854775808) {
                var v8 uint = 18446744073709551615
                v8 += v8
                v9 := GenS1{F0: v8, F1: 9223372036854775807, F2: -1}
                _ = v8
                _ = v9
            }
        }
    } else {
        var v10 int8 = 127
        v11 := GenS1{F0: 542, F1: 42, F2: -639}
        _ = v10
        _ = v11
    }
    var v12 float64 = (1e10 / 3.14159)
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 uint = 2
        _ = fi13
        _ = v14
    }
    v12 = (v12 + v12)
    _ = v12
    return 4294967295
}
func genFn1() uint32 {
    var v15 int32 = 818
    for fi16 := 0; fi16 < 8; fi16++ {
        switch fi16 {
        case 4294967295:
            for fi17 := 0; fi17 < 12; fi17++ {
                v15 += v15
                v18 := []int{}
                var v19 int32 = v15
                _ = fi17
                _ = v18
                _ = v19
            }
        case 128:
            fmt.Println(fi16, v15)
        default:
            for fi20 := 0; fi20 < 10; fi20++ {
                if true {
                    var v21 float64 = (1.5 - (0.1 * 3.14159))
                    v22 := []int{(fi16 - fi16), 0, fi16, -815}
                    _ = v21
                    _ = v22
                }
                _ = fi20
            }
        }
        fmt.Println(fi16, v15)
        fmt.Println(fi16)
        _ = fi16
    }
    v23 := GenS1{F0: 1, F1: 9223372036854775807, F2: 42}
    for fi24 := 0; fi24 < 10; fi24++ {
        if (1e10 != 0.1) {
            for fi25 := 0; fi25 < 11; fi25++ {
                if true {
                    v26 := GenS1{F0: 65535, F1: -9223372036854775808, F2: -701}
                    _ = v26
                }
                if true {
                    if (float64((float32(3.14159) / 1.5)) == ((3.14159 - 0.0) / 3.14159)) {
                        v27 := []int{fi25}
                        if (0 != 42) {
                            _ = len(v27)
                            if true {
                                var v28 int64 = -1
                                _ = v28
                            } else {
                                v29 := GenS1{F0: 570, F1: -1, F2: -566}
                                _ = v29
                            }
                        }
                        v30 := GenS1{F0: 204, F1: -9223372036854775808, F2: -1}
                        _ = v27
                        _ = v30
                    } else {
                        var v31 uint8 = 220
                        var v32 uint16 = 774
                        v32 = 0
                        v31 = v31
                        _ = v31
                        _ = v32
                    }
                    var v33 uint64 = 18446744073709551615
                    v15 = (v15 >> ((18446744073709551615) & 63))
                    _ = v33
                }
                _ = fi25
            }
            v34 := GenS2{F0: 219, F1: fi24}
            v35 := GenS3{F0: (fi24 % (((^(fi24))) | 1)), F1: -128}
            _ = v34
            _ = v35
        }
        v36 := GenS3{F0: (fi24 & (-(fi24))), F1: 127}
        var v37 uint = 4294967295
        _ = fi24
        _ = v36
        _ = v37
    }
    _ = v15
    _ = v23
    return genFn0()
}

func main() {
    v1 := []uint32{2, 463, 1, 1}
    var v2 int8 = 100
    v3 := make(map[int]string)
    var v4 uint8 = 42
    var v5 string = ("" + ("hello" + "foo bar"))
    var v6 int32 = 2147483647
    for fi38 := 0; fi38 < 2; fi38++ {
        for fi39 := 0; fi39 < 9; fi39++ {
            v4 = 128
            var v40 uint = 255
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    if false {
        v41 := GenS4{F0: 18446744073709551615, F1: 2, F2: -1}
        v42 := GenS5{F0: 2147483647, F1: -233, F2: 42}
        _ = v41
        _ = v42
    } else {
        if (0.1 != (float32(3.14159) / 1e10)) {
            var v43 int32 = v6
            for fi44 := 0; fi44 < 7; fi44++ {
                if (255 <= 32767) {
                    var v45 uint8 = (v4 % ((v4) | 1))
                    _ = v45
                } else {
                    v6 = (v43 - (-(v43)))
                }
                if ((!(!true)) && true) {
                    switch (fi44 / ((fi44) | 1)) {
                    case -2147483648:
                        var v46 int32 = (^(v6))
                        _ = v46
                    case 128:
                        fmt.Println(v4)
                    case 32767:
                        fmt.Println(fi44)
                    default:
                        var v47 float64 = 0.1
                        _ = v47
                    }
                    v48 := GenS5{F0: (fi44 & fi44), F1: -32768, F2: 85}
                    fmt.Println(v6)
                    _ = v48
                }
                var v49 uint16 = 2
                _ = fi44
                _ = v49
            }
            v50 := v3[128]
            _ = v43
            _ = v50
        }
    }
    for fi51 := 0; fi51 < 13; fi51++ {
        if true {
            switch (v5 + ((((("a" + ("a" + "a")) + "\n\t") + ((((v5 + "\n\t") + (v5 + ("a" + "hello"))) + (((((((v5 + (v5 + v5)) + "hello") + "") + (v5 + v5)) + v5) + (v5 + (("hello" + (v5 + v5)) + v5))) + v5)) + v5)) + "\n\t") + v5)) {
            case "a":
                if (204 == 133) {
                    if true {
                        v52 := GenS6{F0: -1, F1: 116}
                        if false {
                            var v53 float64 = float64(0.0)
                            if (2147483647 <= v6) {
                                if (!(v4 <= (^(v4)))) {
                                    v6 -= (-(v6))
                                } else {
                                    for fi54 := 0; fi54 < 15; fi54++ {
                                        v6 += (v6 | 0)
                                        var v55 uint32 = 65535
                                        _ = fi54
                                        _ = v55
                                    }
                                    v2 -= v2
                                    v56 := GenS1{F0: 68, F1: -9223372036854775808, F2: 0}
                                    _ = v56
                                }
                                v2 += (-(v2))
                                switch (("\n\t" + (v5 + ("" + ""))) + "a") {
                                case "hello":
                                    v57 := GenS4{F0: 42, F1: 42, F2: 42}
                                    _ = v57
                                case "\n\t":
                                    v6 -= (v6 | 42)
                                default:
                                    var v58 float32 = 0.1
                                    _ = v58
                                }
                            } else {
                                for fi59 := 0; fi59 < 11; fi59++ {
                                    v60 := GenS4{F0: 0, F1: 655, F2: 575}
                                    _ = fi59
                                    _ = v60
                                }
                                if (-607 <= int16(v2)) {
                                    var v61 uint64 = 107
                                    v53 = v53
                                    for fi62 := 0; fi62 < 17; fi62++ {
                                        fmt.Println(v6, v2)
                                        _ = fi62
                                    }
                                    _ = v61
                                }
                            }
                            fmt.Println(v6, v53, v5)
                            _ = v53
                        } else {
                            fmt.Println(fi51)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch fi51 {
                            case 128:
                                for fi63 := 0; fi63 < 11; fi63++ {
                                    _ = len(v1)
                                    v64 := GenS7{F0: v5}
                                    v2 = v2
                                    _ = fi63
                                    _ = v64
                                }
                            case 794:
                                for fi65 := 0; fi65 < 16; fi65++ {
                                    _ = len(v1)
                                    fmt.Println(v2, fi51, v2)
                                    _ = fi65
                                }
                            case 863:
                                for fi66 := 0; fi66 < 17; fi66++ {
                                    fmt.Println(fi51)
                                    v4 -= (^(v4))
                                    fmt.Println(v4, fi66)
                                    _ = fi66
                                }
                            default:
                                v67 := GenS8{F0: v4, F1: 541, F2: 918}
                                _ = v67
                            }
                        }
                        _ = len(v1)
                        _ = v52
                    }
                    if (917 == 18446744073709551615) {
                        fmt.Println(v5, v4, v6)
                        if true {
                            switch v5 {
                            case "foo bar":
                                v5 = (("hello" + "") + ((v5 + v5) + (v5 + v5)))
                            case "\n\t":
                                switch fi51 {
                                case 1:
                                    v68 := GenS8{F0: (v4 - v4), F1: 1, F2: 1}
                                    _ = v68
                                case 128:
                                    v2 -= 1
                                case 0:
                                    fmt.Println(v5, v6, v6)
                                }
                            case "":
                                _ = len(v3)
                            default:
                                if false {
                                    v69 := GenS5{F0: (fi51 & fi51), F1: int16(v2), F2: v4}
                                    _ = v69
                                } else {
                                    fmt.Println(v4)
                                    v1 = append(v1, 0)
                                }
                            }
                        } else {
                            if len(v1) > 0 {
                                v1[0] = genFn1()
                            }
                            if len(v1) > 0 {
                                v1[0] = 832
                            }
                            v5 = "\n\t"
                        }
                        for fi70 := 0; fi70 < 10; fi70++ {
                            var v71 int8 = (v2 / ((127) | 1))
                            _ = fi70
                            _ = v71
                        }
                    } else {
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v1 = append(v1, genFn0())
                            v72 := GenS8{F0: (^(v4)), F1: 4294967295, F2: 1}
                            _ = v72
                        }
                    }
                } else {
                    v5 = ("hello" + (((v5 + v5) + v5) + "a"))
                }
            case "\n\t":
                fmt.Println(fi51, v6)
            default:
                if true {
                    fmt.Println(v5)
                } else {
                    fmt.Println(v5, v6)
                }
            }
            delete(v3, -9223372036854775808)
        } else {
            fmt.Println(v6, fi51, v4)
            for fi73 := 0; fi73 < 16; fi73++ {
                v3[fi73] = "a"
                if true {
                    v74 := GenS9{F0: 32767, F1: 688}
                    var v75 uint = 4294967295
                    _ = v74
                    _ = v75
                } else {
                    v76 := GenS10{F0: v6, F1: 243}
                    var v77 bool = true
                    v77 = (!(v4 < v4))
                    _ = v76
                    _ = v77
                }
                fmt.Println(v4)
                v5 = (v5 + v5)
                delete(v3, fi51)
                _ = fi73
            }
        }
        v78 := GenS2{F0: 0, F1: fi51}
        _ = fi51
        _ = v78
    }
    if false {
        v79 := GenS7{F0: ((v5 + "foo bar") + v5)}
        v80 := v3[148]
        for fi81 := 0; fi81 < 8; fi81++ {
            v5 = v80
            if (2 > 65535) {
                v3[(fi81 - fi81)] = v80
                if (9223372036854775807 <= 1) {
                    switch (fi81 - (fi81 & fi81)) {
                    case 542:
                        if len(v1) > 0 {
                            v1[0] = 0
                        }
                    case -479:
                        v82 := GenS11{F0: (v6 ^ v6), F1: 42, F2: 1}
                        _ = v82
                    case -549:
                        if false {
                            if len(v1) > 0 {
                                v1[0] = genFn1()
                            }
                            var v83 bool = false
                            v84 := GenS5{F0: (fi81 & (fi81 >> ((330) & 63))), F1: 973, F2: (v4 & v4)}
                            _ = v83
                            _ = v84
                        }
                    }
                }
                if ((false || (-1 <= -697)) || false) {
                    for fi85 := 0; fi85 < 9; fi85++ {
                        fmt.Println(v6)
                        _ = fi85
                    }
                    if false {
                        v86 := GenS3{F0: (fi81 ^ (fi81 * 9223372036854775807)), F1: 77}
                        v87 := GenS5{F0: fi81, F1: -399, F2: v4}
                        _ = v86
                        _ = v87
                    }
                    var v88 bool = true
                    _ = v88
                } else {
                    v89 := GenS10{F0: (v6 / (((v6 ^ 0)) | 1)), F1: 820}
                    _ = v89
                }
            } else {
                v1 = append(v1, 505)
                for fi90 := 0; fi90 < 6; fi90++ {
                    v80 = v5
                    v6 = 650
                    switch fi81 {
                    case 127:
                        var v91 uint64 = 288
                        _ = v91
                    case 128:
                        if true {
                            v6 -= (v6 / (((-(v6))) | 1))
                        }
                    }
                    v92 := v3[fi90]
                    v5 = ("a" + v80)
                    _ = fi90
                    _ = v92
                }
            }
            if true {
                fmt.Println(v2, v80)
                if (true || ((-32768 > 32767) && (1e10 > 3.14159))) {
                    v2 = -1
                    var v93 bool = false
                    _ = v93
                } else {
                    if (42 > 666) {
                        v4 = (^(v4))
                    }
                }
            }
            _ = fi81
        }
        _ = v79
        _ = v80
    }
    v2 -= (v2 >> ((1) & 63))
    switch 1 {
    case -2147483648:
        if true {
            switch 1 {
            case 94:
                v4 += (v4 % (((v4 | (v4 | v4))) | 1))
            default:
                var v94 uint32 = 42
                _ = v94
            }
            for fi95 := 0; fi95 < 14; fi95++ {
                if true {
                    v4 = v4
                    var v96 float64 = 1.5
                    _ = v96
                }
                _ = fi95
            }
        } else {
            v97 := GenS7{F0: v5}
            v98 := GenS4{F0: 0, F1: 2, F2: -9223372036854775808}
            v4 = (v4 & 162)
            _ = v97
            _ = v98
        }
    case -1:
        v99 := GenS12{F0: 644, F1: false}
        _ = v99
    default:
        switch 9223372036854775807 {
        case -1:
            _ = len(v1)
        default:
            v100 := GenS12{F0: 4294967295, F1: (42 == -2147483648)}
            _ = v100
        }
    }
    for fi101 := 0; fi101 < 14; fi101++ {
        var v102 uint8 = (^(v4))
        if ((fi101 - 959) == fi101) {
            if (!true) {
                fmt.Println(v2)
            }
            if (((-2.25 * (1.5 - (((1.5 - ((1e10 / 1.5) / -2.25)) / 3.14159) + 1e10))) > (float32(1.5) / 0.1)) && (!false)) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            for fi103 := 0; fi103 < 2; fi103++ {
                var v104 uint = 723
                _ = fi103
                _ = v104
            }
        } else {
            v102 = (^(v102))
            var v105 uint32 = 1
            switch 32767 {
            case 2147483647:
                _ = len(v1)
            case 127:
                switch (fi101 >> ((255) & 63)) {
                case 4294967295:
                    var v106 int16 = 1000
                    _ = v106
                case 1:
                    v102 -= (v4 >> ((0) & 63))
                }
            case -1:
                if true {
                    v107 := GenS10{F0: (v6 >> ((0) & 63)), F1: 907}
                    if len(v1) > 0 {
                        v1[0] = 2147483648
                    }
                    _ = v107
                }
            default:
                if ((v4 < 1) && (42 < 18446744073709551615)) {
                    var v108 uint16 = 255
                    var v109 int32 = v6
                    for fi110 := 0; fi110 < 6; fi110++ {
                        var v111 int32 = (v109 << ((4294967295) & 63))
                        _ = fi110
                        _ = v111
                    }
                    _ = v108
                    _ = v109
                }
            }
            _ = v105
        }
        for fi112 := 0; fi112 < 11; fi112++ {
            if ((v4 != v4) || true) {
                switch v5 {
                case "hello":
                    v2 -= v2
                }
                if true {
                    if true {
                        var v113 int = fi112
                        _ = v113
                    }
                    var v114 int64 = -798
                    var v115 int16 = -175
                    _ = v114
                    _ = v115
                }
            } else {
                fmt.Println(v102, v5, v2)
            }
            var v116 int8 = (-(v2))
            if ((v6 - (v6 + (v6 - v6))) < (-(v6))) {
                v4 = (v4 / ((0) | 1))
            } else {
                v117 := GenS7{F0: v5}
                v118 := GenS13{F0: 969}
                _ = len(v1)
                if len(v1) > 0 {
                    v1[0] = 925
                }
                _ = v117
                _ = v118
            }
            _ = fi112
            _ = v116
        }
        _ = fi101
        _ = v102
    }
    var v119 int32 = v6
    v120 := v3[42]
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v119:", v119)
    fmt.Println("v120:", v120)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v119
    _ = v120
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
