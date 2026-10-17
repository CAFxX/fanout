package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 bool; F1 uint64 }
type GenS9 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS10 struct { F0 uint16; F1 int64 }
type GenS11 struct { F0 uint; F1 int16; F2 string }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS14 struct { F0 int8; F1 uint64; F2 int8 }
func genFn0() uint {
    var v6 uint = 1
    if (90 == 18446744073709551615) {
        v7 := []int{}
        if len(v7) > 0 {
            v7[0] = 32767
        }
        _ = v7
    } else {
        for fi8 := 0; fi8 < 13; fi8++ {
            v9 := GenS1{F0: -811, F1: ""}
            var v10 int32 = 2147483647
            switch fi8 {
            case -9223372036854775808:
                for fi11 := 0; fi11 < 15; fi11++ {
                    var v12 float32 = 3.14159
                    if false {
                        v6 = (v6 % ((v6) | 1))
                        v13 := GenS1{F0: 1, F1: "foo bar"}
                        v12 = v12
                        _ = v13
                    } else {
                        v6 = (v6 * v6)
                    }
                    v6 += (v6 % ((v6) | 1))
                    _ = fi11
                    _ = v12
                }
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        v14 := []int{0, -9223372036854775808}
        _ = v14
    }
    var v15 uint64 = 756
    v6 += 65535
    v16 := GenS1{F0: 42, F1: "hello"}
    _ = v15
    _ = v16
    _ = v6
    return v6
}
func genFn1(p0 uint8) uint16 {
    if (18446744073709551615 != 4294967295) {
        if (!false) {
            v17 := GenS1{F0: 498, F1: "hello"}
            var v18 float64 = 1.5
            var v19 uint8 = p0
            _ = v17
            _ = v18
            _ = v19
        }
        switch 127 {
        case 9223372036854775807:
            for fi20 := 0; fi20 < 18; fi20++ {
                fmt.Println(p0)
                _ = fi20
            }
        case -606:
            if (-97 < 1) {
                for fi21 := 0; fi21 < 2; fi21++ {
                    var v22 bool = (0 > 126)
                    fmt.Println(v22)
                    var v23 int8 = -40
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            }
        case 804:
            switch "\n\t" {
            case "hello":
                var v24 int8 = 127
                _ = v24
            }
        }
        v25 := GenS1{F0: 181, F1: "a"}
        _ = v25
    } else {
        switch 0 {
        case -991:
            fmt.Println(p0)
        case 202:
            var v26 uint16 = 255
            _ = v26
        case -9223372036854775808:
            var v27 bool = (127 > -1)
            _ = v27
        }
        fmt.Println(p0)
        for fi28 := 0; fi28 < 7; fi28++ {
            var v29 uint32 = 2147483648
            _ = fi28
            _ = v29
        }
    }
    var v30 uint = 4294967295
    v30 -= (v30 | genFn0())
    _ = p0
    _ = v30
    return 522
}

func main() {
    var v1 bool = false
    var v2 float64 = ((0.0 + 1.5) + ((3.14159 / (1e10 + 3.14159)) / -2.25))
    var v3 float32 = 3.14159
    var v4 uint32 = 1
    var v5 int32 = 0
    if false {
        switch 128 {
        case 127:
            for fi31 := 0; fi31 < 19; fi31++ {
                fmt.Println(v1, fi31, fi31)
                var v32 int64 = 9223372036854775807
                for fi33 := 0; fi33 < 5; fi33++ {
                    if (127 <= 54) {
                        v34 := []int{42, (fi33 - int(v5))}
                        if v1 {
                            if (true || v1) {
                                v35 := GenS1{F0: v32, F1: "hello"}
                                if (v1 || v1) {
                                    if len(v34) > 0 {
                                        v34[0] = fi33
                                    }
                                }
                                var v36 int32 = -2147483648
                                _ = v35
                                _ = v36
                            }
                            v37 := GenS1{F0: (v32 * v32), F1: ("foo bar" + "")}
                            var v38 int16 = 652
                            _ = v37
                            _ = v38
                        }
                        switch fi33 {
                        case 2147483647:
                            fmt.Println(v2, v2, v2)
                        default:
                            if true {
                                var v39 int8 = -1
                                _ = v39
                            } else {
                                v40 := GenS1{F0: v32, F1: "hello"}
                                _ = v40
                            }
                        }
                        _ = v34
                    } else {
                        switch (fi33 + fi33) {
                        case 32767:
                            var v41 uint32 = v4
                            _ = v41
                        case -2147483648:
                            switch (fi33 % ((-2147483648) | 1)) {
                            case 9223372036854775807:
                                fmt.Println(v4)
                            default:
                                if ((v32 * -9223372036854775808) > v32) {
                                    fmt.Println(v3, v4)
                                } else {
                                    switch 1 {
                                    case 0:
                                        switch fi31 {
                                        case 166:
                                            var v42 bool = false
                                            _ = v42
                                        case 127:
                                            var v43 int = fi33
                                            _ = v43
                                        case 1:
                                            switch (fi31 ^ int(v5)) {
                                            case 127:
                                                var v44 uint64 = uint64(v4)
                                                _ = v44
                                            }
                                        default:
                                            v3 -= 1.5
                                        }
                                    case 9223372036854775807:
                                        if v1 {
                                            fmt.Println(v2, v3, v5)
                                        } else {
                                            if (100 < 42) {
                                                v3 = ((-2.25 / -2.25) * (((((v3 / v3) * v3) / -2.25) * 3.14159) - 1.5))
                                                v45 := []int{}
                                                v2 = ((v2 / v2) - v2)
                                                v46 := GenS2{F0: (v32 / ((0) | 1)), F1: (fi31 + (fi33 >> ((2) & 63)))}
                                                if (!(v2 != (float64(v3) / v2))) {
                                                    if len(v45) > 0 {
                                                        v45[0] = (fi33 << ((188) & 63))
                                                    }
                                                    v45 = append(v45, (fi31 + 127))
                                                    var v47 int16 = 1
                                                    v48 := GenS3{F0: (v5 & (v5 | v5)), F1: (v4 ^ (v4 + (^(v4)))), F2: 2}
                                                    v49 := GenS4{F0: v32, F1: 756}
                                                    fmt.Println(v2, v32, v5)
                                                    _ = v47
                                                    _ = v48
                                                    _ = v49
                                                }
                                                _ = v45
                                                _ = v46
                                            } else {
                                                switch ("\n\t" + ("hello" + "a")) {
                                                case "foo bar":
                                                    var v50 float64 = float64(3.14159)
                                                    _ = v50
                                                case "a":
                                                    fmt.Println(fi31, fi31)
                                                case "\n\t":
                                                    if v1 {
                                                        if v1 {
                                                            switch "" {
                                                            case "hello":
                                                                v51 := GenS5{F0: 128, F1: v4, F2: (^(v32))}
                                                                _ = v51
                                                            case "\n\t":
                                                                v4 = 606
                                                            }
                                                        }
                                                    } else {
                                                        v52 := GenS6{F0: (v32 | 0), F1: ("hello" + ("a" + "hello"))}
                                                        _ = v52
                                                    }
                                                default:
                                                    v2 += v2
                                                }
                                            }
                                            v53 := []int{(fi31 * fi31), (fi31 + fi33), (fi33 + fi31)}
                                            _ = v53
                                        }
                                    default:
                                        var v54 string = "\n\t"
                                        _ = v54
                                    }
                                    if v1 {
                                        v3 = 3.14159
                                        v3 += v3
                                        v55 := []int{fi31}
                                        var v56 int64 = v32
                                        _ = v55
                                        _ = v56
                                    }
                                    if (!("" > ((("" + ("a" + "hello")) + "a") + ""))) {
                                        var v57 int64 = (v32 % ((9223372036854775807) | 1))
                                        fmt.Println(v57, v5, v32)
                                        if ("" == "foo bar") {
                                            var v58 int16 = -1
                                            v59 := GenS1{F0: 0, F1: ""}
                                            if (v1 || true) {
                                                v60 := []int{fi31, fi33, 128}
                                                _ = v60
                                            } else {
                                                switch fi33 {
                                                case 2147483647:
                                                    v61 := []int{fi33, (fi33 & 0), fi33}
                                                    _ = v61
                                                case 128:
                                                    switch -2147483648 {
                                                    case 4294967295:
                                                        v62 := GenS4{F0: -9223372036854775808, F1: 242}
                                                        _ = v62
                                                    default:
                                                        if ((-2147483648 != (fi33 / (((fi33 ^ -1)) | 1))) || (509 <= -9223372036854775808)) {
                                                            v5 -= 42
                                                            if (!(v1 && ((!(!(-9223372036854775808 == (fi31 ^ 2147483647)))) || (((v58 % ((-1) | 1)) <= (-(v58))) && (("foo bar" + "hello") <= ("" + "\n\t")))))) {
                                                                fmt.Println(v32, v2, v3)
                                                            } else {
                                                                fmt.Println(fi31)
                                                                var v63 int = (fi31 + (fi33 >> ((genFn0()) & 63)))
                                                                _ = v63
                                                            }
                                                            if ((v58 << ((347) & 63)) >= 32767) {
                                                                v4 = v4
                                                            } else {
                                                                fmt.Println(v32, v4, v57)
                                                            }
                                                        }
                                                    }
                                                case 4294967295:
                                                    fmt.Println(v4, v32)
                                                default:
                                                    var v64 uint16 = 960
                                                    _ = v64
                                                }
                                                v57 -= v57
                                                v57 += (v32 >> ((719) & 63))
                                            }
                                            _ = v58
                                            _ = v59
                                        } else {
                                            var v65 int = (fi31 * fi31)
                                            switch "hello" {
                                            case "foo bar":
                                                var v66 uint16 = 42
                                                _ = v66
                                            case "a":
                                                if v1 {
                                                    if v1 {
                                                        v67 := []int{fi31}
                                                        var v68 int8 = 1
                                                        _ = v67
                                                        _ = v68
                                                    } else {
                                                        var v69 uint32 = v4
                                                        var v70 int32 = -1
                                                        switch v65 {
                                                        case 264:
                                                            v71 := GenS4{F0: (v32 & v57), F1: 18446744073709551615}
                                                            _ = v71
                                                        }
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                } else {
                                                    fmt.Println(v3, v65)
                                                    fmt.Println(v4, v2, v2)
                                                    if (2 == 0) {
                                                        v72 := GenS3{F0: (v5 + (v5 + (v5 ^ 535))), F1: (v4 ^ (v4 / ((841) | 1))), F2: 0}
                                                        v73 := GenS4{F0: v57, F1: 105}
                                                        v4 = (^(v4))
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                    if v1 {
                                                        var v74 int32 = (v5 & 714)
                                                        v1 = (true && v1)
                                                        _ = v74
                                                    } else {
                                                        fmt.Println(v3, v32, v4)
                                                        if v1 {
                                                            v75 := GenS5{F0: 35, F1: (v4 >> ((255) & 63)), F2: v32}
                                                            _ = v75
                                                        }
                                                    }
                                                    var v76 int8 = 41
                                                    _ = v76
                                                }
                                            }
                                            _ = v65
                                        }
                                        _ = v57
                                    } else {
                                        v77 := GenS4{F0: 0, F1: 2}
                                        if v1 {
                                            var v78 uint16 = genFn1(255)
                                            var v79 int16 = 0
                                            if v1 {
                                                v80 := []int{}
                                                if ((v1 || (!(true && ((v4 >= 4294967295) || v1)))) && v1) {
                                                    v78 = 65535
                                                    var v81 string = ""
                                                    v5 = 1
                                                    _ = v81
                                                }
                                                _ = v80
                                            }
                                            if v1 {
                                                switch (fi33 / (((fi33 >> ((65535) & 63))) | 1)) {
                                                case 2147483647:
                                                    fmt.Println(v5, v78, fi31)
                                                default:
                                                    if v1 {
                                                        v82 := GenS1{F0: (v32 >> ((255) & 63)), F1: "\n\t"}
                                                        v83 := GenS1{F0: v32, F1: "foo bar"}
                                                        v84 := []int{(fi33 >> ((18446744073709551615) & 63)), fi31}
                                                        _ = v82
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                }
                                                if (int8(v4) <= 0) {
                                                    v85 := []int{(fi33 - 128), 93}
                                                    _ = v85
                                                } else {
                                                    if (!(100 >= 127)) {
                                                        v3 -= ((((v3 + float32((v2 / v2))) + v3) - (-2.25 + v3)) - v3)
                                                    }
                                                }
                                                if v1 {
                                                    var v86 float32 = v3
                                                    _ = v86
                                                } else {
                                                    v87 := []int{}
                                                    v88 := GenS7{F0: (v32 + -9223372036854775808)}
                                                    switch ("foo bar" + ("foo bar" + "a")) {
                                                    case "\n\t":
                                                        v89 := GenS2{F0: (-(v32)), F1: fi33}
                                                        _ = v89
                                                    case "":
                                                        var v90 int64 = (^(v32))
                                                        _ = v90
                                                    }
                                                    _ = v87
                                                    _ = v88
                                                }
                                            } else {
                                                if v1 {
                                                    var v91 int32 = (^(v5))
                                                    v2 -= (v2 + ((0.1 - v2) / v2))
                                                    v92 := GenS3{F0: v5, F1: v4, F2: 0}
                                                    _ = v91
                                                    _ = v92
                                                } else {
                                                    v4 -= 2147483647
                                                    if (1 > 1) {
                                                        v93 := GenS8{F0: v1, F1: 1}
                                                        var v94 int32 = (v5 & v5)
                                                        v2 = float64(v3)
                                                        _ = v93
                                                        _ = v94
                                                    } else {
                                                        var v95 int8 = -128
                                                        var v96 uint16 = (v78 % (((v78 | v78)) | 1))
                                                        fmt.Println(v96, v3)
                                                        _ = v95
                                                        _ = v96
                                                    }
                                                    var v97 uint32 = v4
                                                    _ = v97
                                                }
                                                var v98 uint8 = 42
                                                _ = v98
                                            }
                                            var v99 int16 = 521
                                            _ = v78
                                            _ = v79
                                            _ = v99
                                        } else {
                                            v100 := []int{}
                                            _ = v100
                                        }
                                        _ = v77
                                    }
                                }
                            }
                        }
                        var v101 float64 = (v2 / v2)
                        _ = v101
                    }
                    _ = fi33
                }
                _ = fi31
                _ = v32
            }
        default:
            var v102 uint8 = 1
            _ = v102
        }
        v5 = v5
    }
    for fi103 := 0; fi103 < 10; fi103++ {
        var v104 int32 = 2147483647
        switch ("foo bar" + (((((("" + "foo bar") + "\n\t") + (("" + ("foo bar" + "hello")) + "foo bar")) + (("" + "a") + "\n\t")) + "a") + "hello")) {
        case "foo bar":
            v3 += v3
        case "hello":
            if ((false && v1) && false) {
                switch fi103 {
                case -843:
                    v105 := []int{fi103}
                    _ = v105
                case 0:
                    switch "hello" {
                    case "":
                        v106 := []int{fi103, fi103}
                        _ = v106
                    case "a":
                        v1 = v1
                    default:
                        for fi107 := 0; fi107 < 10; fi107++ {
                            v108 := GenS1{F0: 0, F1: ("foo bar" + ("foo bar" + "foo bar"))}
                            var v109 float32 = 0.0
                            fmt.Println(fi103, v3, v3)
                            _ = fi107
                            _ = v108
                            _ = v109
                        }
                    }
                case -1:
                    v110 := []int{(fi103 ^ fi103)}
                    _ = v110
                default:
                    if v1 {
                        for fi111 := 0; fi111 < 2; fi111++ {
                            switch 2147483647 {
                            case -1:
                                v112 := GenS8{F0: (51 > 42), F1: 18446744073709551615}
                                _ = v112
                            case 42:
                                if false {
                                    v113 := []int{fi103, fi103, fi103}
                                    v114 := GenS8{F0: false, F1: 4294967295}
                                    _ = v113
                                    _ = v114
                                }
                            case -103:
                                fmt.Println(v2, v4)
                                switch (fi111 | 42) {
                                case -432:
                                    v115 := GenS9{F0: -128, F1: 255, F2: 18446744073709551615}
                                    _ = v115
                                case 127:
                                    v116 := []int{-9223372036854775808}
                                    _ = v116
                                case 9223372036854775807:
                                    if v1 {
                                        v4 = (v4 ^ 65535)
                                        if v1 {
                                            if true {
                                                v117 := GenS3{F0: 42, F1: v4, F2: 2}
                                                v118 := []int{fi103, 0, (fi111 / ((fi103) | 1))}
                                                switch 2147483647 {
                                                case 180:
                                                    v118 = append(v118, int(v5))
                                                case -2147483648:
                                                    v2 = float64((-2.25 + (1.5 + v3)))
                                                    _ = len(v118)
                                                case 1:
                                                    switch ("" + "foo bar") {
                                                    case "foo bar":
                                                        if (genFn1(127) > 909) {
                                                            var v119 uint64 = 4294967295
                                                            _ = v119
                                                        }
                                                    case "a":
                                                        v5 = v5
                                                    case "hello":
                                                        switch ("hello" + "\n\t") {
                                                        case "":
                                                            switch (^(fi111)) {
                                                            case 32767:
                                                                _ = len(v118)
                                                                v4 = (v4 + v4)
                                                            case 32768:
                                                                fmt.Println(fi111, v4, v4)
                                                            case 1:
                                                                if v1 {
                                                                    v2 -= v2
                                                                    var v120 int8 = 1
                                                                    _ = v120
                                                                } else {
                                                                    if (("" + "a") == ("" + ("foo bar" + ((("foo bar" + "hello") + "foo bar") + "foo bar")))) {
                                                                        _ = len(v118)
                                                                        v1 = v1
                                                                        var v121 uint64 = 18446744073709551615
                                                                        _ = v121
                                                                    } else {
                                                                        if len(v118) >= 2 {
                                                                            v118 = v118[1:len(v118)]
                                                                        }
                                                                        switch "a" {
                                                                        case "a":
                                                                            v122 := GenS1{F0: -1, F1: "a"}
                                                                            _ = v122
                                                                        }
                                                                        fmt.Println(v4)
                                                                    }
                                                                    if true {
                                                                        if len(v118) >= 2 {
                                                                            v118 = v118[1:len(v118)]
                                                                        }
                                                                        if (43 >= 0) {
                                                                            v104 += (^(v5))
                                                                        } else {
                                                                            v1 = ((((("\n\t" + ("" + "foo bar")) + "\n\t") + "a") + "\n\t") >= "foo bar")
                                                                        }
                                                                    } else {
                                                                        var v123 uint16 = genFn1(255)
                                                                        _ = len(v118)
                                                                        _ = v123
                                                                    }
                                                                }
                                                            }
                                                            var v124 int32 = 788
                                                            if true {
                                                                v4 = (v4 + v4)
                                                            } else {
                                                                v1 = true
                                                                v125 := GenS2{F0: 11, F1: 42}
                                                                v126 := GenS8{F0: (18446744073709551615 >= 632), F1: 1}
                                                                _ = v125
                                                                _ = v126
                                                            }
                                                            _ = v124
                                                        case "hello":
                                                            if (!v1) {
                                                                var v127 bool = (128 < 170)
                                                                _ = v127
                                                            } else {
                                                                if len(v118) > 0 {
                                                                    v118[0] = (fi103 | fi111)
                                                                }
                                                                if len(v118) >= 2 {
                                                                    v118 = v118[1:len(v118)]
                                                                }
                                                                v128 := GenS3{F0: (^(v5)), F1: v4, F2: 263}
                                                                var v129 string = (((("hello" + "") + "hello") + "") + ("a" + "hello"))
                                                                _ = v128
                                                                _ = v129
                                                            }
                                                        }
                                                    }
                                                default:
                                                    v5 -= 0
                                                }
                                                if v1 {
                                                    var v130 string = ("a" + "hello")
                                                    var v131 float64 = float64(float32(3.14159))
                                                    _ = v130
                                                    _ = v131
                                                } else {
                                                    fmt.Println(v1)
                                                }
                                                _ = v117
                                                _ = v118
                                            } else {
                                                switch 4294967295 {
                                                case -833:
                                                    switch "hello" {
                                                    case "foo bar":
                                                        v132 := GenS2{F0: -911, F1: fi103}
                                                        _ = v132
                                                    }
                                                case 9223372036854775807:
                                                    if (genFn0() != 343) {
                                                        v133 := GenS8{F0: (560 < 813), F1: 2}
                                                        var v134 string = "hello"
                                                        var v135 uint32 = v4
                                                        var v136 uint64 = 333
                                                        fmt.Println(v135)
                                                        var v137 uint = 4294967295
                                                        _ = v133
                                                        _ = v134
                                                        _ = v135
                                                        _ = v136
                                                        _ = v137
                                                    } else {
                                                        if v1 {
                                                            fmt.Println(v104)
                                                            if false {
                                                                if false {
                                                                    switch fi103 {
                                                                    case -690:
                                                                        v138 := []int{-750, fi103, fi103}
                                                                        _ = v138
                                                                    case 9223372036854775807:
                                                                        fmt.Println(fi103, v4)
                                                                        fmt.Println(v1)
                                                                        v104 = (v5 ^ (v104 & 42))
                                                                    case 127:
                                                                        if ((!(("" >= "a") && (!((((v3 / 1.5) + 0.1) - 0.1) >= 1.5)))) && (-1 <= 42)) {
                                                                            if (!v1) {
                                                                                var v139 float32 = v3
                                                                                _ = v139
                                                                            } else {
                                                                                var v140 int8 = 1
                                                                                _ = v140
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            } else {
                                                                var v141 uint8 = 1
                                                                var v142 bool = (v1 && ((v4 << ((4294967295) & 63)) > v4))
                                                                if true {
                                                                    var v143 int = 9223372036854775807
                                                                    _ = v143
                                                                }
                                                                _ = v141
                                                                _ = v142
                                                            }
                                                            v144 := []int{(^(fi103)), 2147483647, (fi111 ^ (fi103 / ((1) | 1))), (fi103 + (^(fi111)))}
                                                            _ = v144
                                                        } else {
                                                            var v145 bool = v1
                                                            _ = v145
                                                        }
                                                    }
                                                default:
                                                    switch (fi103 >> ((18446744073709551615) & 63)) {
                                                    case -122:
                                                        v3 = v3
                                                    case -9223372036854775808:
                                                        v146 := GenS3{F0: (v104 << ((189) & 63)), F1: (^(v4)), F2: 42}
                                                        _ = v146
                                                    default:
                                                        switch "a" {
                                                        case "hello":
                                                            v147 := []int{-2147483648, -1, 4294967295, fi103}
                                                            _ = v147
                                                        case "a":
                                                            var v148 uint8 = 215
                                                            _ = v148
                                                        default:
                                                            v149 := GenS9{F0: -128, F1: 732, F2: 369}
                                                            switch fi111 {
                                                            case -1:
                                                                fmt.Println(v4, v2, fi111)
                                                            case 42:
                                                                v150 := []int{(fi111 << ((genFn0()) & 63))}
                                                                _ = v150
                                                            case 1:
                                                                v151 := GenS5{F0: 127, F1: (v4 | 805), F2: 0}
                                                                _ = v151
                                                            default:
                                                                var v152 int64 = 42
                                                                _ = v152
                                                            }
                                                            v2 = ((v2 - v2) * v2)
                                                            _ = v149
                                                        }
                                                    }
                                                }
                                                v1 = (210 <= 95)
                                                var v153 uint64 = 906
                                                _ = v153
                                            }
                                            v154 := []int{fi111, fi103, fi103}
                                            if v1 {
                                                v3 -= (((-2.25 + v3) - v3) / v3)
                                                var v155 bool = (1 == 0)
                                                _ = v155
                                            } else {
                                                var v156 int32 = (v104 - v104)
                                                v3 = ((float32((float64((3.14159 - (v3 - (0.1 + v3)))) - -2.25)) * v3) / v3)
                                                v2 = (1.5 / v2)
                                                _ = len(v154)
                                                _ = v156
                                            }
                                            _ = v154
                                        }
                                    } else {
                                        v157 := []int{fi103, fi111, int(v5), fi103}
                                        _ = v157
                                    }
                                }
                            }
                            _ = fi111
                        }
                        v158 := GenS10{F0: 808, F1: 0}
                        v5 = v5
                        _ = v158
                    }
                }
                for fi159 := 0; fi159 < 1; fi159++ {
                    v104 -= v5
                    if (-32768 != -674) {
                        v1 = v1
                        var v160 int32 = v5
                        switch (("a" + "") + (("a" + "a") + "a")) {
                        case "hello":
                            v161 := GenS11{F0: 719, F1: -32768, F2: "foo bar"}
                            var v162 uint32 = 4294967295
                            v163 := GenS4{F0: -2, F1: 2}
                            _ = v161
                            _ = v162
                            _ = v163
                        }
                        _ = v160
                    }
                    switch 32767 {
                    case 127:
                        if ((4294967295 >= v4) && v1) {
                            if (false && (genFn0() > 625)) {
                                v164 := []int{fi159}
                                _ = v164
                            }
                            var v165 int16 = 1000
                            v166 := []int{}
                            v1 = false
                            if v1 {
                                if true {
                                    if (42 >= 1) {
                                        v2 = v2
                                        var v167 uint8 = 0
                                        var v168 uint8 = v167
                                        v169 := GenS8{F0: v1, F1: 2}
                                        _ = v167
                                        _ = v168
                                        _ = v169
                                    }
                                    v2 = ((v2 + (v2 * v2)) * v2)
                                } else {
                                    switch "hello" {
                                    case "foo bar":
                                        v2 = v2
                                    case "hello":
                                        if len(v166) >= 2 {
                                            v166 = v166[1:len(v166)]
                                        }
                                    default:
                                        var v170 int32 = v5
                                        _ = v170
                                    }
                                }
                                fmt.Println(fi159, v104, v165)
                                var v171 int8 = 42
                                if (((!(!v1)) && false) && (705 > 255)) {
                                    v171 = (v171 + (v171 << ((42) & 63)))
                                    if ((!((^(v5)) < 724)) && (((((3.14159 - (float32(v2) / float32(v2))) * (v3 + v3)) * 1.5) <= 1e10) || (v1 && (v1 && v1)))) {
                                        if len(v166) >= 2 {
                                            v166 = v166[1:len(v166)]
                                        }
                                    } else {
                                        var v172 string = ("" + "")
                                        if true {
                                            if (fi159 >= (fi159 | fi103)) {
                                                v3 = 3.14159
                                                v173 := GenS2{F0: -9223372036854775808, F1: fi159}
                                                _ = v173
                                            }
                                            switch -72 {
                                            case -9223372036854775808:
                                                var v174 uint8 = 127
                                                _ = v174
                                            case 422:
                                                switch 42 {
                                                case -320:
                                                    var v175 float32 = (1.5 * (float32((-2.25 / v2)) * v3))
                                                    _ = v175
                                                case 32767:
                                                    v2 = 0.1
                                                case -2147483648:
                                                    v1 = (!false)
                                                default:
                                                    v166 = append(v166, (-(fi159)))
                                                }
                                            case 664:
                                                _ = len(v166)
                                            }
                                            v3 = (v3 / 0.1)
                                        }
                                        _ = v172
                                    }
                                }
                                _ = v171
                            } else {
                                switch fi159 {
                                case -608:
                                    if (42 != 1) {
                                        if (v2 >= (-2.25 - (1e10 - (((((float64((0.0 / (float32(1e10) * (0.0 / ((v3 * v3) / (v3 / v3)))))) / (1.5 - float64((float32((v2 + (((-2.25 - (3.14159 - 1.5)) + (v2 - (0.0 / v2))) / 3.14159))) * v3)))) * ((-2.25 - 1.5) - ((v2 + v2) - v2))) + v2) - (v2 + ((0.0 - -2.25) + -2.25))) * v2)))) {
                                            if (v4 != v4) {
                                                if false {
                                                    if true {
                                                        var v176 uint8 = 128
                                                        v176 += (v176 * 1)
                                                        var v177 int32 = (v5 + v5)
                                                        var v178 uint64 = 42
                                                        fmt.Println(v176, v1, v165)
                                                        _ = v176
                                                        _ = v177
                                                        _ = v178
                                                    }
                                                    var v179 bool = v1
                                                    if (("" < ("a" + ("a" + ""))) && v179) {
                                                        var v180 uint16 = 301
                                                        fmt.Println(v5, v179)
                                                        _ = v180
                                                    }
                                                    v181 := GenS12{F0: 42}
                                                    _ = v179
                                                    _ = v181
                                                } else {
                                                    var v182 uint32 = (v4 / (((^(v4))) | 1))
                                                    v4 += (^(v182))
                                                    _ = v182
                                                }
                                                fmt.Println(v2)
                                            }
                                            var v183 string = ("" + "")
                                            _ = v183
                                        } else {
                                            switch ("hello" + ((("\n\t" + "\n\t") + ("\n\t" + (((("\n\t" + ("hello" + "a")) + "a") + "foo bar") + ("foo bar" + "foo bar")))) + ("" + ("\n\t" + "hello")))) {
                                            case "hello":
                                                v2 += v2
                                            case "a":
                                                if v1 {
                                                    v2 = v2
                                                    if ("a" == "\n\t") {
                                                        if v1 {
                                                            var v184 string = "\n\t"
                                                            var v185 uint8 = 2
                                                            _ = v184
                                                            _ = v185
                                                        }
                                                    }
                                                }
                                            }
                                            v5 = -474
                                            v186 := GenS11{F0: 1, F1: (v165 * -1), F2: ("foo bar" + "a")}
                                            var v187 bool = (v3 != v3)
                                            var v188 int8 = 127
                                            v166 = append(v166, fi103)
                                            _ = v186
                                            _ = v187
                                            _ = v188
                                        }
                                    } else {
                                        if len(v166) > 0 {
                                            v166[0] = fi159
                                        }
                                        v2 = ((float64((0.0 / v3)) + (0.1 * (v2 - v2))) - v2)
                                    }
                                case 127:
                                    if (0 <= 128) {
                                        if len(v166) >= 2 {
                                            v166 = v166[1:len(v166)]
                                        }
                                    }
                                case 8:
                                    var v189 uint8 = 127
                                    _ = v189
                                }
                                fmt.Println(fi159)
                                if v1 {
                                    v104 = v5
                                    switch ("" + "") {
                                    case "":
                                        v165 = -32768
                                    case "\n\t":
                                        var v190 int64 = 1
                                        _ = v190
                                    case "hello":
                                        switch "" {
                                        case "hello":
                                            v166 = append(v166, 4294967295)
                                        case "a":
                                            v3 += 3.14159
                                            fmt.Println(v1)
                                        }
                                    }
                                    var v191 float64 = v2
                                    _ = v191
                                } else {
                                    v104 += 1
                                    v165 += (v165 * v165)
                                    var v192 uint32 = (v4 / ((v4) | 1))
                                    _ = v192
                                }
                            }
                            switch -1 {
                            case 0:
                                _ = len(v166)
                            case 9223372036854775807:
                                var v193 uint32 = v4
                                _ = v193
                            default:
                                v194 := GenS12{F0: 1}
                                _ = v194
                            }
                            _ = v165
                            _ = v166
                        } else {
                            var v195 uint8 = 2
                            var v196 uint = uint(v4)
                            _ = v195
                            _ = v196
                        }
                    default:
                        var v197 uint8 = 224
                        _ = v197
                    }
                    _ = fi159
                }
                if (!false) {
                    for fi198 := 0; fi198 < 2; fi198++ {
                        v199 := []int{-90}
                        v3 = v3
                        v1 = (genFn0() >= 4294967295)
                        switch 4294967295 {
                        case 128:
                            if false {
                                fmt.Println(fi198, v4)
                                var v200 uint16 = uint16(v4)
                                switch (fi103 % ((fi103) | 1)) {
                                case 2147483647:
                                    v4 += 4294967295
                                }
                                _ = v200
                            }
                        case 2147483647:
                            v201 := GenS8{F0: (534 >= 4294967295), F1: 1}
                            _ = v201
                        }
                        var v202 int8 = -128
                        _ = fi198
                        _ = v199
                        _ = v202
                    }
                    v203 := []int{(fi103 & fi103), (fi103 + -2147483648), fi103, (fi103 & 32768)}
                    if false {
                        fmt.Println(v2)
                        fmt.Println(fi103, v5)
                    } else {
                        for fi204 := 0; fi204 < 3; fi204++ {
                            v203 = append(v203, 0)
                            var v205 float32 = ((((1e10 / v3) / v3) * v3) - v3)
                            if v1 {
                                v5 += v104
                            } else {
                                var v206 int16 = -852
                                if ((v3 / v205) == v205) {
                                    if (uint(v4) <= 2) {
                                        var v207 uint8 = 199
                                        _ = v207
                                    }
                                    v203 = append(v203, (fi204 - (-(fi103))))
                                    v203 = append(v203, 42)
                                } else {
                                    v208 := GenS13{F0: 753, F1: 66, F2: 65535}
                                    v206 = (v206 + 42)
                                    var v209 uint16 = genFn1(128)
                                    v104 = (v104 + v104)
                                    if len(v203) > 0 {
                                        v203[0] = fi103
                                    }
                                    v209 = (v209 >> ((1) & 63))
                                    _ = v208
                                    _ = v209
                                }
                                fmt.Println(v2)
                                _ = v206
                            }
                            _ = fi204
                            _ = v205
                        }
                    }
                    _ = v203
                } else {
                    v4 = (v4 | v4)
                    var v210 int64 = 9223372036854775807
                    _ = v210
                }
            } else {
                for fi211 := 0; fi211 < 15; fi211++ {
                    if true {
                        v212 := []int{(fi103 | -1), 2147483647, fi103}
                        if v1 {
                            v3 += 3.14159
                        }
                        _ = v212
                    }
                    var v213 uint16 = genFn1(1)
                    var v214 uint = 2
                    _ = fi211
                    _ = v213
                    _ = v214
                }
                v104 -= v5
            }
        case "":
            for fi215 := 0; fi215 < 13; fi215++ {
                v216 := []int{int(v4), (fi215 | (fi103 | 42)), fi215}
                switch ((("\n\t" + "foo bar") + ("\n\t" + "")) + "foo bar") {
                case "hello":
                    switch "" {
                    case "a":
                        if len(v216) > 0 {
                            v216[0] = (fi215 & fi103)
                        }
                    }
                case "\n\t":
                    v4 = (v4 + v4)
                    _ = len(v216)
                    v1 = (!v1)
                }
                _ = fi215
                _ = v216
            }
        default:
            v217 := GenS7{F0: int64(fi103)}
            _ = v217
        }
        v1 = v1
        _ = fi103
        _ = v104
    }
    var v218 uint8 = 218
    fmt.Println(v4)
    v219 := []int{42, 128}
    v1 = (0 >= -1)
    if (!(v5 >= v5)) {
        v220 := GenS3{F0: v5, F1: (v4 / (((v4 % ((v4) | 1))) | 1)), F2: 4294967295}
        _ = v220
    }
    if v1 {
        var v221 uint32 = (v4 - (^(v4)))
        if v1 {
            v222 := GenS14{F0: -66, F1: 1, F2: 127}
            _ = v222
        }
        v219 = append(v219, -1)
        _ = v221
    } else {
        if (("\n\t" + ("foo bar" + "hello")) != "a") {
            fmt.Println(v1, v5, v3)
        } else {
            if len(v219) >= 2 {
                v219 = v219[1:len(v219)]
            }
            var v223 int16 = 1
            v1 = (((v3 * (0.0 * float32(1e10))) + v3) >= v3)
            _ = v223
        }
    }
    switch 262 {
    case 251:
        v3 += -2.25
    case 546:
        v5 = -1
    }
    v3 = (((0.0 + v3) - 1e10) / float32(3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v218:", v218)
    fmt.Println("slicelen v219:", len(v219), v219)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v218
    _ = v219
    _ = v3
    _ = v4
    _ = v5
}
