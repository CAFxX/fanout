package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint; F2 uint64 }
type GenS3 struct { F0 bool; F1 uint; F2 string }
type GenS4 struct { F0 int; F1 int; F2 uint32 }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int64; F1 int32 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint16; F1 int16; F2 int32 }
type GenS14 struct { F0 int16; F1 int }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 int16; F1 string; F2 bool }
func genFn0(p0 uint64, p1 int16) int64 {
    var v6 int8 = int8(p1)
    var v7 string = "foo bar"
    if (p0 < 538) {
        fmt.Println(p0, p1)
    }
    var v8 string = v7
    v9 := []int{}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -9223372036854775808
}
func genFn1(p0 string) uint64 {
    switch "hello" {
    case "":
        v10 := []int{759, 840}
        _ = v10
    case "\n\t":
        v11 := []int{32767, 1, 32768}
        _ = v11
    }
    p0 = p0
    v12 := []int{127, -2147483648, 4294967295}
    p0 = "foo bar"
    if (3.14159 > 3.14159) {
        p0 = ("" + (((p0 + (("\n\t" + p0) + p0)) + p0) + "a"))
        var v13 uint16 = 216
        _ = v13
    }
    _ = p0
    _ = v12
    return 18446744073709551615
}

func main() {
    v1 := []string{"foo bar", (("foo bar" + "") + (("\n\t" + "foo bar") + "a")), ((("" + "hello") + ("\n\t" + ("" + (("\n\t" + ("foo bar" + (("" + (("foo bar" + (("a" + ("\n\t" + "foo bar")) + "\n\t")) + "hello")) + "hello"))) + "hello")))) + "a"), ("foo bar" + "a")}
    var v2 int64 = -1
    v3 := []int{-1}
    var v4 uint8 = 255
    var v5 uint = 382
    v3 = append(v3, 4294967295)
    for fi14 := 0; fi14 < 8; fi14++ {
        v2 -= (v2 & (v2 + -1))
        if false {
            v4 -= (^(v4))
            if false {
                if len(v1) > 0 {
                    v1[0] = ("hello" + "hello")
                }
                for fi15 := 0; fi15 < 7; fi15++ {
                    v4 = 62
                    _ = fi15
                }
                v4 += 128
            } else {
                if ((!false) || (-473 > -10)) {
                    if len(v1) > 0 {
                        v1[0] = ("a" + (("hello" + ("" + "a")) + ""))
                    }
                    v16 := GenS1{F0: 42}
                    _ = v16
                }
                v4 = 42
                var v17 int8 = 0
                _ = v17
            }
        } else {
            var v18 int32 = 2147483647
            switch 379 {
            case 42:
                for fi19 := 0; fi19 < 16; fi19++ {
                    switch fi19 {
                    case -198:
                        var v20 uint16 = 984
                        _ = v20
                    }
                    v5 = 255
                    v21 := GenS2{F0: 1, F1: v5, F2: 0}
                    _ = fi19
                    _ = v21
                }
            case 4294967295:
                var v22 string = "foo bar"
                _ = v22
            case 736:
                v23 := GenS1{F0: uint8(v5)}
                _ = v23
            default:
                var v24 float32 = 0.0
                _ = v24
            }
            _ = v18
        }
        _ = fi14
    }
    v25 := GenS3{F0: false, F1: 255, F2: ("\n\t" + "\n\t")}
    switch "a" {
    case "":
        for fi26 := 0; fi26 < 6; fi26++ {
            for fi27 := 0; fi27 < 17; fi27++ {
                var v28 int32 = -1
                v4 = v4
                v29 := GenS4{F0: (fi27 & 720), F1: 127, F2: 4294967295}
                _ = fi27
                _ = v28
                _ = v29
            }
            if len(v1) > 0 {
                v1[0] = ("a" + ("\n\t" + "a"))
            }
            v2 = (v2 + 0)
            _ = fi26
        }
    case "foo bar":
        v4 = 127
    case "\n\t":
        var v30 uint8 = (v4 ^ v4)
        _ = v30
    default:
        var v31 int64 = v2
        _ = v31
    }
    for fi32 := 0; fi32 < 15; fi32++ {
        var v33 int64 = (v2 << ((v5) & 63))
        fmt.Println(v5, v5, fi32)
        v34 := GenS5{F0: v4, F1: 2}
        _ = fi32
        _ = v33
        _ = v34
    }
    v35 := GenS3{F0: false, F1: v5, F2: "hello"}
    v3 = append(v3, 9223372036854775807)
    v1 = append(v1, "hello")
    var v36 int64 = (-(v2))
    v2 -= (^(v36))
    if (!true) {
        switch 43 {
        case -653:
            _ = len(v3)
        default:
            v36 = (v36 | v36)
        }
        var v37 int16 = 320
        _ = v37
    } else {
        if false {
            switch -2147483648 {
            case -692:
                if (false || true) {
                    var v38 uint = v5
                    v2 = v36
                    var v39 int64 = (v36 >> (((v5 ^ (v5 - v38))) & 63))
                    _ = v38
                    _ = v39
                }
            case -272:
                var v40 int = 32768
                _ = v40
            default:
                v5 += (v5 + (v5 / ((v5) | 1)))
            }
            for fi41 := 0; fi41 < 9; fi41++ {
                v4 = v4
                fmt.Println(v2, v2, v5)
                var v42 int8 = 1
                _ = fi41
                _ = v42
            }
            if (!false) {
                if len(v3) > 0 {
                    v3[0] = -984
                }
            } else {
                var v43 uint = (^(v5))
                v43 += 865
                _ = v43
            }
        } else {
            var v44 uint8 = v4
            v1 = append(v1, ("" + "foo bar"))
            _ = v44
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v45 int32 = 42
        _ = v45
    }
    v4 -= v4
    var v46 float32 = ((1.5 * 3.14159) - 1e10)
    for fi47 := 0; fi47 < 1; fi47++ {
        switch "" {
        case "foo bar":
            if false {
                v46 = v46
                switch (("foo bar" + "a") + "\n\t") {
                case "":
                    v48 := GenS6{F0: (fi47 << ((0) & 63)), F1: (v5 - 555)}
                    _ = v48
                case "hello":
                    if true {
                        for fi49 := 0; fi49 < 8; fi49++ {
                            var v50 int16 = -32768
                            if len(v1) > 0 {
                                v1[0] = "hello"
                            }
                            _ = fi49
                            _ = v50
                        }
                    }
                case "foo bar":
                    switch 42 {
                    case -9223372036854775808:
                        v1 = append(v1, ("\n\t" + ""))
                    case 380:
                        var v51 int = (fi47 * fi47)
                        _ = v51
                    default:
                        for fi52 := 0; fi52 < 19; fi52++ {
                            v36 = genFn0(142, -1)
                            _ = fi52
                        }
                    }
                }
                v53 := GenS3{F0: (v36 != v2), F1: v5, F2: "foo bar"}
                _ = v53
            } else {
                fmt.Println(v2, v5)
                for fi54 := 0; fi54 < 16; fi54++ {
                    var v55 int64 = v2
                    switch "" {
                    case "hello":
                        v46 = (v46 + 1.5)
                    }
                    if true {
                        if false {
                            v56 := GenS2{F0: v36, F1: v5, F2: 47}
                            v5 -= v5
                            _ = v56
                        } else {
                            if ((v4 / (((v4 << ((v5) & 63))) | 1)) != v4) {
                                v57 := GenS4{F0: (fi54 | 2147483647), F1: fi54, F2: 200}
                                var v58 int8 = 100
                                v1 = append(v1, (((((("a" + "\n\t") + "foo bar") + "\n\t") + ("\n\t" + ("\n\t" + "hello"))) + ("hello" + "foo bar")) + (("" + (("\n\t" + "foo bar") + "a")) + "a")))
                                _ = v57
                                _ = v58
                            }
                            var v59 float32 = v46
                            fmt.Println(v5, fi54, v2)
                            _ = v59
                        }
                        v60 := GenS1{F0: (v4 << ((v5) & 63))}
                        fmt.Println(fi54, v55)
                        _ = v60
                    } else {
                        v2 = v55
                    }
                    var v61 uint16 = 42
                    v61 = (v61 ^ v61)
                    _ = fi54
                    _ = v55
                    _ = v61
                }
                var v62 int8 = -11
                _ = v62
            }
        default:
            var v63 int32 = 42
            _ = v63
        }
        _ = fi47
    }
    for fi64 := 0; fi64 < 18; fi64++ {
        var v65 int = (fi64 / ((fi64) | 1))
        for fi66 := 0; fi66 < 7; fi66++ {
            switch (fi66 << ((18446744073709551615) & 63)) {
            case -985:
                var v67 uint8 = v4
                _ = v67
            default:
                fmt.Println(v5)
            }
            v2 = genFn0(179, -32768)
            _ = fi66
        }
        _ = fi64
        _ = v65
    }
    var v68 uint = 126
    for fi69 := 0; fi69 < 11; fi69++ {
        for fi70 := 0; fi70 < 12; fi70++ {
            if true {
                v71 := GenS4{F0: (fi70 >> ((v68) & 63)), F1: fi70, F2: 42}
                switch (("hello" + ("a" + "a")) + (((("\n\t" + ("a" + "")) + "a") + "\n\t") + ("a" + (("\n\t" + "hello") + "foo bar")))) {
                case "hello":
                    _ = len(v1)
                case "\n\t":
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                case "a":
                    fmt.Println(v36, v5)
                }
                if true {
                    fmt.Println(v5, v2, v5)
                    var v72 uint = uint(v4)
                    var v73 int32 = int32(v4)
                    var v74 uint64 = 2
                    _ = v72
                    _ = v73
                    _ = v74
                }
                _ = v71
            }
            var v75 uint16 = uint16(v68)
            if (true || false) {
                var v76 string = ""
                _ = v76
            } else {
                switch 223 {
                case 32767:
                    if (v4 != v4) {
                        if (((!false) || true) && true) {
                            v2 += (v36 & (v36 & (v36 + 0)))
                            var v77 uint64 = 42
                            _ = v77
                        } else {
                            if (true || (0.0 < 1e10)) {
                                v3 = append(v3, fi69)
                                var v78 float64 = 1e10
                                if false {
                                    if false {
                                        var v79 int8 = 1
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v79
                                    } else {
                                        v46 += (0.0 / float32(v78))
                                    }
                                    fmt.Println(v2)
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if false {
                                        var v80 float32 = (float32(v78) + (v46 / -2.25))
                                        fmt.Println(v68, fi70)
                                        v46 += ((0.1 / (v46 + v46)) + (((v46 - (0.0 + v46)) - v80) / v46))
                                        _ = v80
                                    } else {
                                        switch (("hello" + ((((("\n\t" + "a") + "a") + "\n\t") + "a") + "foo bar")) + ("hello" + "a")) {
                                        case "hello":
                                            var v81 bool = (645 != 449)
                                            _ = v81
                                        default:
                                            v75 = (^(v75))
                                        }
                                        if ((v78 * (v78 - ((v78 / 1e10) / ((v78 * (0.1 * 0.1)) - v78)))) < 0.1) {
                                            var v82 uint32 = 2147483647
                                            fmt.Println(v46, v68)
                                            var v83 float32 = v46
                                            if (false && ("hello" == "foo bar")) {
                                                v68 = v5
                                                var v84 uint16 = uint16(v82)
                                                _ = v84
                                            }
                                            _ = v82
                                            _ = v83
                                        } else {
                                            v1 = append(v1, "hello")
                                            v85 := GenS7{F0: "\n\t"}
                                            _ = v85
                                        }
                                        v86 := GenS8{F0: (^(v5))}
                                        _ = v86
                                    }
                                    var v87 int = -713
                                    switch (fi69 ^ 9223372036854775807) {
                                    case -107:
                                        var v88 bool = false
                                        _ = v88
                                    }
                                    _ = v87
                                }
                                _ = v78
                            }
                            var v89 int32 = 0
                            v75 += v75
                            _ = v89
                        }
                        v68 += 42
                        switch fi69 {
                        case 42:
                            v5 += 4294967295
                        case -988:
                            if (-128 >= -1) {
                                if true {
                                    if false {
                                        if (80 > v4) {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = "foo bar"
                                            }
                                            var v90 float32 = ((1.5 - v46) + float32(0.1))
                                            var v91 int32 = 42
                                            var v92 int16 = 901
                                            _ = v90
                                            _ = v91
                                            _ = v92
                                        }
                                        var v93 uint8 = (v4 / ((v4) | 1))
                                        _ = v93
                                    } else {
                                        v4 = v4
                                        var v94 uint16 = (^(v75))
                                        fmt.Println(v68, v5)
                                        _ = v94
                                    }
                                    fmt.Println(v36)
                                }
                                var v95 bool = (1e10 < v46)
                                v46 += v46
                                _ = v95
                            }
                        default:
                            var v96 string = "foo bar"
                            _ = v96
                        }
                    }
                case -421:
                    var v97 uint8 = (v4 & (v4 ^ v4))
                    _ = v97
                case 0:
                    v98 := GenS9{F0: v2, F1: 463}
                    _ = v98
                }
                v99 := GenS10{F0: (v75 & 2)}
                if true {
                    if ((-1 > 42) && false) {
                        v3 = append(v3, 42)
                        var v100 int64 = -1
                        v101 := GenS11{F0: -2147483648}
                        fmt.Println(v4, v2, v5)
                        v68 -= (v68 % ((v5) | 1))
                        v102 := GenS4{F0: fi69, F1: 9223372036854775807, F2: 2}
                        switch 9223372036854775807 {
                        case 2147483647:
                            var v103 int64 = v100
                            _ = v103
                        case 963:
                            if true {
                                if (42 == -710) {
                                    var v104 int32 = 2147483647
                                    v5 -= v68
                                    _ = len(v1)
                                    _ = v104
                                }
                            } else {
                                v46 = float32((3.14159 / 1.0))
                                v105 := GenS10{F0: v75}
                                _ = v105
                            }
                        case -88:
                            if (0 != 1) {
                                var v106 uint = v5
                                var v107 int8 = 0
                                switch "a" {
                                case "a":
                                    switch (fi69 << ((948) & 63)) {
                                    case -969:
                                        v46 -= 3.14159
                                    case -494:
                                        if (v46 <= v46) {
                                            var v108 int32 = int32(fi69)
                                            v68 -= 0
                                            var v109 uint8 = (v4 / (((^(v4))) | 1))
                                            if (112 < 4294967295) {
                                                var v110 int64 = (v36 - 217)
                                                _ = v110
                                            }
                                            _ = v108
                                            _ = v109
                                        } else {
                                            var v111 uint16 = (v75 * v75)
                                            _ = v111
                                        }
                                    case 457:
                                        v106 -= v68
                                    default:
                                        fmt.Println(v75)
                                    }
                                case "\n\t":
                                    if false {
                                        fmt.Println(v106, v68)
                                        var v112 float64 = (3.14159 * 0.0)
                                        var v113 uint32 = 731
                                        _ = v112
                                        _ = v113
                                    } else {
                                        fmt.Println(v100, fi70)
                                        var v114 uint32 = 51
                                        _ = v114
                                    }
                                case "foo bar":
                                    if false {
                                        v115 := GenS2{F0: (v36 << (((v68 + (v68 >> ((uint(v4)) & 63)))) & 63)), F1: v5, F2: 18446744073709551615}
                                        if true {
                                            var v116 int64 = genFn0(2, 32767)
                                            if false {
                                                switch (("foo bar" + (("" + ("\n\t" + ("hello" + "foo bar"))) + ("" + ("foo bar" + "hello")))) + "hello") {
                                                case "foo bar":
                                                    var v117 int32 = 2147483647
                                                    _ = v117
                                                case "hello":
                                                    fmt.Println(v106, v46)
                                                case "a":
                                                    v118 := GenS12{F0: 4294967295}
                                                    _ = v118
                                                }
                                                v107 -= int8(v106)
                                                v106 = v5
                                            } else {
                                                var v119 uint64 = 2
                                                _ = v119
                                            }
                                            v120 := GenS10{F0: 42}
                                            if true {
                                                if false {
                                                    fmt.Println(v2)
                                                    v121 := GenS4{F0: fi70, F1: 0, F2: 2}
                                                    v122 := GenS13{F0: 42, F1: 465, F2: -680}
                                                    _ = v121
                                                    _ = v122
                                                } else {
                                                    fmt.Println(v2)
                                                    switch "hello" {
                                                    case "hello":
                                                        fmt.Println(v100)
                                                    case "a":
                                                        v123 := GenS13{F0: uint16(v5), F1: -32768, F2: 1}
                                                        _ = v123
                                                    case "\n\t":
                                                        v124 := GenS14{F0: -1, F1: 1}
                                                        _ = v124
                                                    default:
                                                        v125 := GenS1{F0: v4}
                                                        _ = v125
                                                    }
                                                    fmt.Println(v116)
                                                }
                                            }
                                            _ = v116
                                            _ = v120
                                        } else {
                                            var v126 int = fi69
                                            _ = len(v3)
                                            _ = v126
                                        }
                                        var v127 uint32 = 0
                                        _ = v115
                                        _ = v127
                                    } else {
                                        v36 = -9223372036854775808
                                        fmt.Println(v4)
                                        if (true || (!((!false) || true))) {
                                            if (v5 == 264) {
                                                var v128 bool = (!false)
                                                v4 += (v4 % (((v4 << ((v106) & 63))) | 1))
                                                _ = v128
                                            } else {
                                                v129 := GenS1{F0: v4}
                                                _ = v129
                                            }
                                            switch ("\n\t" + "") {
                                            case "":
                                                fmt.Println(v5)
                                            }
                                            var v130 float64 = 3.14159
                                            _ = v130
                                        }
                                    }
                                default:
                                    var v131 float32 = (v46 + float32(1e10))
                                    _ = v131
                                }
                                if (1 <= -359) {
                                    var v132 uint16 = v75
                                    v5 -= v68
                                    v2 = v100
                                    v133 := GenS15{F0: v75}
                                    var v134 int = fi69
                                    _ = v132
                                    _ = v133
                                    _ = v134
                                }
                                _ = v106
                                _ = v107
                            }
                        default:
                            var v135 int64 = (v2 - (v36 ^ -969))
                            _ = v135
                        }
                        v1 = append(v1, "")
                        _ = v100
                        _ = v101
                        _ = v102
                    } else {
                        v36 = (^(v2))
                        v136 := GenS3{F0: true, F1: 255, F2: "foo bar"}
                        fmt.Println(v5, v5)
                        _ = v136
                    }
                    var v137 int32 = 1
                    v5 = 4294967295
                    _ = v137
                }
                _ = v99
            }
            _ = fi70
            _ = v75
        }
        v138 := GenS16{F0: 42, F1: ((("a" + ("a" + (("foo bar" + "") + ("hello" + "hello")))) + ("a" + "foo bar")) + ((("" + "\n\t") + "\n\t") + ("" + "a"))), F2: (4294967295 >= 42)}
        _ = fi69
        _ = v138
    }
    switch 9223372036854775807 {
    case 4294967295:
        var v139 uint16 = 42
        _ = v139
    case 32767:
        if false {
            v68 = v5
        }
    case -1:
        v68 = uint(v4)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v46
    _ = v5
    _ = v68
}
