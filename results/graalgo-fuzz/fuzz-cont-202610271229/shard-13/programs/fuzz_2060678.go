package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint8; F1 string; F2 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int8; F1 uint64; F2 uint }
type GenS8 struct { F0 int64; F1 uint16; F2 int }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int8; F1 int8; F2 int64 }
type GenS13 struct { F0 int32; F1 uint; F2 uint32 }
func genFn0(p0 bool, p1 uint64) int8 {
    switch (("foo bar" + "hello") + ("\n\t" + "")) {
    case "hello":
        fmt.Println(p1)
    }
    v7 := []int{-9223372036854775808}
    var v8 int32 = 0
    for fi9 := 0; fi9 < 14; fi9++ {
        fmt.Println(p1)
        _ = len(v7)
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 int = (-(fi10))
            fmt.Println(p1, v11, p1)
            if true {
                if p0 {
                    if len(v7) > 0 {
                        v7[0] = fi10
                    }
                }
                v7 = append(v7, -2147483648)
                fmt.Println(v11, fi9)
            } else {
                v12 := GenS1{F0: -258, F1: 65535, F2: v8}
                _ = v12
            }
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -1
}
func genFn1() uint32 {
    if (!true) {
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 uint32 = 28
            fmt.Println(fi13, v14)
            _ = fi13
            _ = v14
        }
        if (!(255 > 65535)) {
            var v15 uint8 = 53
            var v16 uint8 = v15
            v16 = (v15 * (v15 << ((255) & 63)))
            _ = v15
            _ = v16
        } else {
            switch "hello" {
            case "":
                var v17 uint = 0
                _ = v17
            }
        }
    } else {
        for fi18 := 0; fi18 < 15; fi18++ {
            for fi19 := 0; fi19 < 12; fi19++ {
                switch 32767 {
                case 1:
                    var v20 uint64 = 859
                    _ = v20
                case -9223372036854775808:
                    switch (fi18 / ((fi19) | 1)) {
                    case -85:
                        var v21 uint8 = 132
                        _ = v21
                    case -2147483648:
                        v22 := []int{fi18, (fi18 - (fi19 % ((9223372036854775807) | 1)))}
                        _ = v22
                    }
                }
                v23 := GenS2{F0: 0, F1: "a", F2: -1}
                var v24 uint8 = 1
                _ = fi19
                _ = v23
                _ = v24
            }
            _ = fi18
        }
    }
    v25 := GenS1{F0: 42, F1: 2147483647, F2: 459}
    var v26 uint16 = 862
    v26 = (v26 >> ((18446744073709551615) & 63))
    _ = v25
    _ = v26
    return 257
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 308
    var v3 float32 = 1.5
    v4 := []uint32{130, v2}
    var v5 uint16 = 425
    v6 := make(map[int]uint32)
    v27 := v6[0]
    if true {
        v5 -= (v5 / (((v5 / ((744) | 1))) | 1))
        var v28 int64 = 1
        switch "" {
        case "\n\t":
            var v29 uint64 = 18446744073709551615
            _ = v29
        case "":
            var v30 uint16 = 42
            _ = v30
        case "a":
            switch "hello" {
            case "a":
                if (1 == 2147483647) {
                    var v31 uint32 = (v2 / (((v27 / ((genFn1()) | 1))) | 1))
                    for fi32 := 0; fi32 < 10; fi32++ {
                        v6[2147483647] = uint32(v5)
                        _ = fi32
                    }
                    if true {
                        for fi33 := 0; fi33 < 10; fi33++ {
                            if ((v31 - v31) != v31) {
                                if (4294967295 > (fi33 / ((fi33) | 1))) {
                                    v27 = (^(v31))
                                    v28 = v28
                                    switch "hello" {
                                    case "a":
                                        switch "\n\t" {
                                        case "":
                                            var v34 bool = false
                                            _ = v34
                                        }
                                    }
                                } else {
                                    var v35 string = ("a" + ("\n\t" + "foo bar"))
                                    fmt.Println(v3, v35)
                                    var v36 uint32 = (v27 - (v27 ^ v27))
                                    _ = v35
                                    _ = v36
                                }
                            }
                            v28 += v28
                            v3 = v3
                            _ = fi33
                        }
                    }
                    _ = v31
                } else {
                    v2 = genFn1()
                    v4 = append(v4, 383)
                    fmt.Println(v27)
                }
            default:
                for fi37 := 0; fi37 < 19; fi37++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    for fi38 := 0; fi38 < 7; fi38++ {
                        v3 = 1.5
                        var v39 string = ""
                        _ = fi38
                        _ = v39
                    }
                    _ = fi37
                }
            }
        }
        _ = v28
    } else {
        if ((127 != 128) && true) {
            fmt.Println(v27)
            fmt.Println(v5, v27)
        }
    }
    fmt.Println(v2)
    for fi40 := 0; fi40 < 5; fi40++ {
        if true {
            v1[""] = fi40
            if true {
                for fi41 := 0; fi41 < 8; fi41++ {
                    fmt.Println(fi40, v2, v5)
                    var v42 int64 = 9223372036854775807
                    v43 := GenS3{F0: 127}
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                switch "" {
                case "\n\t":
                    fmt.Println(v3, v3, v3)
                case "a":
                    var v44 int = (fi40 >> ((679) & 63))
                    _ = v44
                default:
                    fmt.Println(fi40, v5, v3)
                }
            } else {
                for fi45 := 0; fi45 < 3; fi45++ {
                    v3 = float32(0.0)
                    var v46 int8 = 18
                    switch "" {
                    case "a":
                        var v47 int = -261
                        _ = v47
                    case "":
                        _ = len(v4)
                    default:
                        v2 += (v2 % (((v2 << ((uint(v5)) & 63))) | 1))
                    }
                    if true {
                        fmt.Println(v2, v27)
                    }
                    switch "hello" {
                    case "hello":
                        v48 := GenS4{F0: "", F1: 141}
                        _ = v48
                    default:
                        v49 := GenS5{F0: 108, F1: 0}
                        fmt.Println(v27)
                        _ = v49
                    }
                    _ = fi45
                    _ = v46
                }
            }
        } else {
            for fi50 := 0; fi50 < 11; fi50++ {
                var v51 int16 = -32768
                _ = fi50
                _ = v51
            }
            var v52 string = "hello"
            _ = v52
        }
        v3 -= (v3 / 1.0)
        v53 := GenS5{F0: 127, F1: 18446744073709551615}
        _ = fi40
        _ = v53
    }
    _ = len(v1)
    if (-737 != -1) {
        var v54 float32 = (float32(1.5) - 3.14159)
        v55 := GenS2{F0: 127, F1: "hello", F2: 0}
        _ = v54
        _ = v55
    } else {
        if true {
            fmt.Println(v27)
            v27 -= (^(v2))
        } else {
            v56 := GenS6{F0: true}
            v57 := GenS6{F0: true}
            for fi58 := 0; fi58 < 6; fi58++ {
                if (!false) {
                    if false {
                        var v59 uint8 = 1
                        var v60 string = "\n\t"
                        _ = v59
                        _ = v60
                    }
                }
                _ = fi58
            }
            _ = v56
            _ = v57
        }
        v5 -= v5
        v2 += v2
    }
    for fi61 := 0; fi61 < 4; fi61++ {
        v62 := GenS1{F0: -775, F1: (v27 + v2), F2: -1}
        for fi63 := 0; fi63 < 17; fi63++ {
            var v64 int8 = 42
            v65 := GenS7{F0: 0, F1: 42, F2: uint(v5)}
            _ = fi63
            _ = v64
            _ = v65
        }
        for fi66 := 0; fi66 < 17; fi66++ {
            v27 += v27
            v4 = append(v4, genFn1())
            v3 = 0.1
            v27 = (^(v27))
            switch fi66 {
            case 54:
                v5 -= v5
            case -732:
                var v67 string = "\n\t"
                _ = v67
            default:
                var v68 float64 = float64(float32(0.1))
                _ = v68
            }
            _ = fi66
        }
        _ = fi61
        _ = v62
    }
    fmt.Println(v3)
    switch 684 {
    case 256:
        switch ("foo bar" + "hello") {
        case "\n\t":
            v69 := GenS3{F0: 14}
            _ = v69
        case "hello":
            if ((false && (!false)) || (-128 == -1)) {
                if false {
                    if false {
                        fmt.Println(v3)
                    } else {
                        switch 9223372036854775807 {
                        case 32768:
                            for fi70 := 0; fi70 < 18; fi70++ {
                                v5 = v5
                                switch "\n\t" {
                                case "":
                                    var v71 int = fi70
                                    _ = v71
                                case "foo bar":
                                    v27 -= v2
                                case "\n\t":
                                    if true {
                                        v72 := GenS2{F0: 2, F1: "", F2: 72}
                                        if true {
                                            v3 -= (0.1 * 3.14159)
                                        }
                                        _ = v72
                                    }
                                }
                                _ = fi70
                            }
                        }
                        v27 -= genFn1()
                        v73 := GenS7{F0: 127, F1: 42, F2: 274}
                        _ = v73
                    }
                    switch ("a" + ("foo bar" + (((("a" + "foo bar") + "hello") + "") + "a"))) {
                    case "\n\t":
                        var v74 int = 127
                        _ = v74
                    case "hello":
                        var v75 uint32 = v27
                        _ = v75
                    default:
                        if false {
                            var v76 uint32 = genFn1()
                            v5 += v5
                            var v77 float64 = 3.14159
                            _ = v76
                            _ = v77
                        } else {
                            v27 += 2147483647
                            var v78 int8 = 42
                            var v79 int16 = 1
                            _ = v78
                            _ = v79
                        }
                    }
                    v80 := GenS8{F0: -354, F1: 489, F2: 32768}
                    _ = v80
                } else {
                    fmt.Println(v27, v2)
                }
                delete(v6, 9223372036854775807)
            } else {
                for fi81 := 0; fi81 < 18; fi81++ {
                    v82 := GenS1{F0: 0, F1: 581, F2: 631}
                    if (!(float64(0.1) <= (((float64(v3) * 1e10) - -2.25) / (3.14159 + 0.1)))) {
                        v27 = (^(v2))
                        v2 += genFn1()
                        for fi83 := 0; fi83 < 8; fi83++ {
                            var v84 uint32 = 280
                            if ("hello" > ("\n\t" + "")) {
                                v4 = append(v4, 423)
                                v3 = 0.0
                                v85 := v1[(((("a" + "foo bar") + ("" + "a")) + ("foo bar" + "a")) + (("hello" + "\n\t") + ("hello" + (("foo bar" + "hello") + ((("a" + "hello") + "") + ("\n\t" + "\n\t"))))))]
                                fmt.Println(v5, fi83)
                                switch (v85 - v85) {
                                case 32768:
                                    v3 = v3
                                case 127:
                                    if (genFn1() > v27) {
                                        _ = len(v4)
                                    } else {
                                        v86 := GenS9{F0: -1}
                                        var v87 uint8 = 1
                                        _ = v86
                                        _ = v87
                                    }
                                case -9223372036854775808:
                                    if (2 > 461) {
                                        var v88 bool = false
                                        v84 = 2147483647
                                        switch fi83 {
                                        case 4294967295:
                                            var v89 bool = (!(!v88))
                                            _ = v89
                                        case 2147483647:
                                            if true {
                                                var v90 int = (fi83 * fi81)
                                                v91 := GenS10{F0: 42}
                                                _ = v90
                                                _ = v91
                                            } else {
                                                v84 -= v27
                                                if (!(2 == 132)) {
                                                    var v92 int16 = 0
                                                    _ = v92
                                                }
                                                var v93 string = "a"
                                                _ = v93
                                            }
                                        default:
                                            fmt.Println(v27)
                                        }
                                        _ = v88
                                    }
                                    if true {
                                        var v94 int8 = genFn0((true || (0.0 != 3.14159)), 18446744073709551615)
                                        v85 = (v85 | -825)
                                        switch ("" + (("foo bar" + ("hello" + "hello")) + "hello")) {
                                        case "hello":
                                            v84 -= (v84 * uint32(v5))
                                        case "":
                                            v4 = append(v4, (v84 % ((genFn1()) | 1)))
                                        }
                                        var v95 int64 = int64(v94)
                                        var v96 uint32 = genFn1()
                                        v27 -= (v2 >> ((42) & 63))
                                        _ = v94
                                        _ = v95
                                        _ = v96
                                    } else {
                                        fmt.Println(v2, v2)
                                    }
                                    fmt.Println(v3)
                                }
                                _ = v85
                            }
                            var v97 int8 = 73
                            v98 := v1["hello"]
                            _ = fi83
                            _ = v84
                            _ = v97
                            _ = v98
                        }
                    } else {
                        switch -477 {
                        case -9223372036854775808:
                            switch -9223372036854775808 {
                            case -755:
                                var v99 int8 = genFn0((481 < 490), 0)
                                _ = v99
                            case -9223372036854775808:
                                v5 = v5
                            default:
                                switch ("foo bar" + ("a" + ("\n\t" + ("" + "")))) {
                                case "":
                                    var v100 int = fi81
                                    _ = v100
                                case "hello":
                                    v101 := GenS9{F0: -115}
                                    _ = v101
                                }
                            }
                        default:
                            v102 := GenS9{F0: 1}
                            _ = v102
                        }
                        for fi103 := 0; fi103 < 17; fi103++ {
                            v104 := GenS7{F0: 0, F1: 956, F2: 4294967295}
                            v27 = v2
                            var v105 uint32 = (v27 * genFn1())
                            var v106 int = -2147483648
                            v1[("a" + "")] = (-(fi103))
                            _ = fi103
                            _ = v104
                            _ = v105
                            _ = v106
                        }
                    }
                    v107 := GenS11{F0: 65535}
                    _ = fi81
                    _ = v107
                    _ = v82
                }
            }
        }
    case -1:
        v4 = append(v4, v27)
    case 4294967295:
        v108 := GenS4{F0: "", F1: -787}
        _ = v108
    default:
        v109 := GenS12{F0: -1, F1: int8(v5), F2: 956}
        _ = v109
    }
    for fi110 := 0; fi110 < 10; fi110++ {
        v5 -= v5
        if ((0.0 + (3.14159 * ((float64(v3) / 1.5) / float64(v3)))) <= 1e10) {
            var v111 int32 = 42
            var v112 int32 = -1
            _ = v111
            _ = v112
        } else {
            v3 = 0.1
            v27 = (v2 + v2)
            for fi113 := 0; fi113 < 14; fi113++ {
                v27 = v2
                switch "a" {
                case "":
                    var v114 uint8 = 128
                    _ = v114
                default:
                    var v115 int = fi110
                    _ = v115
                }
                _ = fi113
            }
        }
        if (163 < 1) {
            v116 := GenS13{F0: -842, F1: 255, F2: (v27 + genFn1())}
            if (1 < 636) {
                if len(v4) > 0 {
                    v4[0] = v27
                }
            } else {
                fmt.Println(fi110, v27, v3)
                v117 := GenS10{F0: 2147483647}
                var v118 int8 = 127
                _ = v117
                _ = v118
            }
            _ = v116
        } else {
            v2 = (^(v27))
        }
        _ = fi110
    }
    var v119 int8 = 127
    var v120 int32 = -2147483648
    fmt.Println(v5, v27, v3)
    v120 += (v120 | (v120 >> ((42) & 63)))
    v3 += -2.25
    for fi121 := 0; fi121 < 13; fi121++ {
        if true {
            v3 = v3
        }
        _ = fi121
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v119:", v119)
    fmt.Println("v120:", v120)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v119
    _ = v120
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
