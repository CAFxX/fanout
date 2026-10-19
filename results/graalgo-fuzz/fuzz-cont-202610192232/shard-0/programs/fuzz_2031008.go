package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 string; F1 uint; F2 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int; F1 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int16; F1 int32; F2 uint8 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint }
func genFn0(p0 uint64) int {
    v7 := GenS1{F0: int16(p0)}
    p0 = p0
    p0 = (p0 + (p0 * p0))
    _ = p0
    _ = v7
    return 1
}

func main() {
    var v1 string = ("foo bar" + ((("hello" + ("foo bar" + "\n\t")) + (("" + "hello") + "a")) + ("hello" + "\n\t")))
    var v2 int16 = 32767
    var v3 int16 = 96
    v4 := []int{}
    var v5 string = v1
    v6 := []int{127, -829, 4294967295, 655}
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 float64 = -2.25
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 17; fi11++ {
            if false {
                if ((0.1 - (1e10 / 1.0)) != 0.1) {
                    v12 := GenS2{F0: 9223372036854775807, F1: ""}
                    _ = v12
                } else {
                    switch fi11 {
                    case 318:
                        v13 := GenS1{F0: v3}
                        _ = v13
                    case -1:
                        var v14 float32 = 1e10
                        var v15 string = v5
                        var v16 uint16 = 2
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    var v17 bool = (-1 > -81)
                    var v18 bool = v17
                    _ = v17
                    _ = v18
                }
            } else {
                switch (v5 + v1) {
                case "hello":
                    fmt.Println(fi11)
                case "":
                    if false {
                        v4 = append(v4, genFn0(1))
                    } else {
                        switch "" {
                        case "\n\t":
                            switch v1 {
                            case "hello":
                                if (!((!(127 >= 0)) && false)) {
                                    var v19 uint32 = 65535
                                    v20 := GenS1{F0: -32768}
                                    _ = v19
                                    _ = v20
                                }
                            case "\n\t":
                                if false {
                                    v3 = v2
                                    fmt.Println(v5, v3, fi8)
                                    if (8 >= 1) {
                                        var v21 uint16 = 442
                                        var v22 uint64 = 42
                                        v5 = v5
                                        _ = v21
                                        _ = v22
                                    } else {
                                        if (v1 != v5) {
                                            var v23 uint16 = 1
                                            v2 -= v2
                                            v23 = v23
                                            v24 := GenS3{F0: v5, F1: 700, F2: (2 >= 1)}
                                            _ = v23
                                            _ = v24
                                        }
                                        var v25 uint64 = 4294967295
                                        v1 = ("a" + v5)
                                        _ = v25
                                    }
                                    var v26 bool = (!(4294967295 > 0))
                                    fmt.Println(v2)
                                    v27 := GenS4{F0: (v3 >> ((42) & 63))}
                                    v28 := GenS1{F0: 42}
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                } else {
                                    fmt.Println(v3, v3)
                                    if len(v4) > 0 {
                                        v4[0] = fi11
                                    }
                                    v29 := GenS2{F0: 0, F1: "a"}
                                    _ = v29
                                }
                            case "a":
                                fmt.Println(v5)
                            }
                        case "hello":
                            var v30 float64 = float64(-2.25)
                            var v31 string = ("a" + v5)
                            _ = v30
                            _ = v31
                        case "foo bar":
                            if false {
                                if false {
                                    var v32 float64 = 3.14159
                                    v32 += 0.1
                                    if false {
                                        fmt.Println(fi11, v5, v1)
                                        if (2 <= 2) {
                                            v33 := GenS5{F0: 42}
                                            if (true || false) {
                                                v1 = ""
                                                switch fi8 {
                                                case 1:
                                                    _ = len(v4)
                                                case 0:
                                                    switch -9223372036854775808 {
                                                    case 2147483647:
                                                        if false {
                                                            v34 := GenS5{F0: 1}
                                                            v5 = "a"
                                                            switch 32767 {
                                                            case 118:
                                                                fmt.Println(v3, fi8, v1)
                                                            }
                                                            _ = v34
                                                        } else {
                                                            var v35 uint16 = 300
                                                            v4 = append(v4, 128)
                                                            _ = v35
                                                        }
                                                    case 0:
                                                        if len(v6) > 0 {
                                                            v6[0] = fi11
                                                        }
                                                    case 42:
                                                        if (-44 == -457) {
                                                            _ = len(v4)
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            v32 -= 3.14159
                                                        }
                                                    }
                                                }
                                                fmt.Println(fi8, v1)
                                            }
                                            _ = v33
                                        } else {
                                            v2 -= v3
                                        }
                                        v5 = ("a" + v1)
                                        v36 := GenS3{F0: v1, F1: 42, F2: false}
                                        if len(v6) > 0 {
                                            v6[0] = (fi11 * (fi11 - fi8))
                                        }
                                        if (2 == 2) {
                                            switch v1 {
                                            case "\n\t":
                                                if (1e10 == 1.5) {
                                                    v3 = v2
                                                    switch ("" + v5) {
                                                    case "hello":
                                                        if (-892 >= (^(v3))) {
                                                            var v37 uint = 255
                                                            switch v1 {
                                                            case "":
                                                                var v38 uint32 = 2
                                                                _ = v38
                                                            }
                                                            v37 = (v37 + (v37 | v37))
                                                            fmt.Println(fi11)
                                                            _ = v37
                                                        }
                                                    case "\n\t":
                                                        if ((0.1 != 3.14159) && (0.0 <= (((float32(v32) - 1.5) / 3.14159) - ((float32((0.1 + v32)) / 3.14159) * 1e10)))) {
                                                            v1 = v1
                                                            var v39 bool = false
                                                            _ = v39
                                                        }
                                                    }
                                                    v3 += v3
                                                    v3 = (v2 >> ((539) & 63))
                                                    if (1.5 == v32) {
                                                        v6 = append(v6, genFn0(42))
                                                    }
                                                }
                                            case "":
                                                _ = len(v4)
                                            case "foo bar":
                                                v40 := GenS6{F0: fi11}
                                                _ = v40
                                            default:
                                                var v41 bool = (9223372036854775807 > int64(v3))
                                                v5 = "a"
                                                _ = v41
                                            }
                                        } else {
                                            v2 = v2
                                        }
                                        _ = v36
                                    }
                                    var v42 float64 = v32
                                    v43 := GenS3{F0: (v1 + v5), F1: 65535, F2: true}
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v32 = (3.14159 - v32)
                                    v44 := GenS7{F0: 0, F1: (2 == 13)}
                                    _ = v32
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                } else {
                                    v2 -= v3
                                }
                            } else {
                                if (135 == 128) {
                                    v6 = append(v6, fi8)
                                } else {
                                    v5 = (v1 + v5)
                                    var v45 uint = 18446744073709551615
                                    _ = v45
                                }
                                var v46 uint = 255
                                var v47 uint32 = 2147483647
                                _ = v46
                                _ = v47
                            }
                            v48 := GenS1{F0: v2}
                            _ = v48
                        default:
                            switch genFn0(4294967295) {
                            case 0:
                                fmt.Println(fi8)
                            case 42:
                                var v49 int8 = -79
                                _ = v49
                            default:
                                if len(v6) > 0 {
                                    v6[0] = (fi8 - fi11)
                                }
                            }
                        }
                        v3 = v2
                        v2 -= (v2 << ((110) & 63))
                    }
                default:
                    var v50 float32 = float32(((-2.25 - 1e10) * 3.14159))
                    _ = v50
                }
                if true {
                    var v51 uint = 42
                    _ = v51
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            _ = fi11
        }
        _ = fi8
    }
    var v52 uint8 = 255
    v53 := GenS1{F0: v2}
    var v54 string = v1
    var v55 string = v1
    for fi56 := 0; fi56 < 16; fi56++ {
        fmt.Println(v2)
        v5 = v1
        v1 = v55
        _ = fi56
    }
    if (4294967295 <= 42) {
        fmt.Println(v2)
        switch (v55 + v55) {
        case "\n\t":
            if (false && true) {
                v57 := GenS8{F0: (^(v52))}
                if (!true) {
                    switch 2147483647 {
                    case -2147483648:
                        var v58 uint16 = 0
                        _ = v58
                    case 42:
                        v55 = ((v5 + "") + v1)
                    case 9223372036854775807:
                        switch "" {
                        case "foo bar":
                            if len(v6) > 0 {
                                v6[0] = 2147483647
                            }
                        case "a":
                            _ = len(v6)
                        }
                    default:
                        for fi59 := 0; fi59 < 15; fi59++ {
                            switch (v54 + ((v1 + v5) + (("hello" + v55) + v54))) {
                            case "\n\t":
                                var v60 int16 = 1
                                _ = v60
                            case "hello":
                                switch fi59 {
                                case -1:
                                    for fi61 := 0; fi61 < 13; fi61++ {
                                        switch v5 {
                                        case "\n\t":
                                            v2 = v2
                                        case "":
                                            var v62 int = genFn0(42)
                                            _ = v62
                                        }
                                        if false {
                                            fmt.Println(v5, v54)
                                            var v63 bool = (!(-40 >= -128))
                                            _ = v63
                                        } else {
                                            var v64 int32 = -993
                                            switch fi61 {
                                            case 42:
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            }
                                            v65 := GenS3{F0: "a", F1: 18446744073709551615, F2: (v64 < v64)}
                                            _ = v64
                                            _ = v65
                                        }
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        _ = fi61
                                    }
                                case 232:
                                    v52 = v52
                                case 127:
                                    v66 := GenS9{F0: 42}
                                    _ = v66
                                }
                            }
                            switch v55 {
                            case "":
                                if (2 <= 948) {
                                    if (9223372036854775807 != fi59) {
                                        v55 = v5
                                        switch v55 {
                                        case "":
                                            v5 = ("\n\t" + v5)
                                        case "a":
                                            for fi67 := 0; fi67 < 17; fi67++ {
                                                switch v1 {
                                                case "foo bar":
                                                    v1 = (v1 + "")
                                                case "hello":
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                case "\n\t":
                                                    if (605 > 18446744073709551615) {
                                                        v68 := GenS10{F0: 388}
                                                        switch fi67 {
                                                        case 42:
                                                            v4 = append(v4, (^(fi67)))
                                                        case -2147483648:
                                                            if true {
                                                                v2 = v2
                                                            }
                                                        case -223:
                                                            if true {
                                                                if true {
                                                                    v69 := GenS4{F0: v3}
                                                                    v1 = (v5 + v1)
                                                                    v70 := GenS8{F0: v52}
                                                                    _ = v69
                                                                    _ = v70
                                                                } else {
                                                                    if len(v6) > 0 {
                                                                        v6[0] = (fi59 % (((-(fi59))) | 1))
                                                                    }
                                                                }
                                                                if (-59 >= -128) {
                                                                    _ = len(v4)
                                                                    if true {
                                                                        fmt.Println(v2, v55)
                                                                        v55 = "\n\t"
                                                                        var v71 uint = 891
                                                                        _ = v71
                                                                    }
                                                                    var v72 float64 = ((0.0 * 1.5) * ((1.5 - 0.1) / 1e10))
                                                                    _ = v72
                                                                } else {
                                                                    var v73 int32 = 2147483647
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        v74 := GenS2{F0: int64(v52), F1: ""}
                                                                        _ = v74
                                                                    case "foo bar":
                                                                        v75 := GenS11{F0: (-(v3)), F1: v73, F2: 136}
                                                                        _ = v75
                                                                    }
                                                                    if true {
                                                                        v73 = v73
                                                                        switch fi59 {
                                                                        case 2147483647:
                                                                            v76 := GenS10{F0: 515}
                                                                            _ = v76
                                                                        case 32767:
                                                                            v77 := GenS7{F0: fi59, F1: true}
                                                                            _ = v77
                                                                        case 743:
                                                                            var v78 int16 = v2
                                                                            _ = v78
                                                                        default:
                                                                            fmt.Println(v1, v2)
                                                                        }
                                                                        var v79 uint64 = 18446744073709551615
                                                                        switch (v54 + v54) {
                                                                        case "\n\t":
                                                                            switch (v5 + v55) {
                                                                            case "a":
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                            }
                                                                        }
                                                                        v80 := GenS7{F0: fi67, F1: true}
                                                                        _ = v79
                                                                        _ = v80
                                                                    } else {
                                                                        v3 += -1
                                                                        if ((v73 * (v73 ^ v73)) != (v73 - (v73 << ((255) & 63)))) {
                                                                            fmt.Println(v1)
                                                                            var v81 int = (-(fi59))
                                                                            var v82 float32 = float32((0.0 - 0.0))
                                                                            v52 = (v52 >> ((2) & 63))
                                                                            _ = v81
                                                                            _ = v82
                                                                        } else {
                                                                            v83 := GenS3{F0: v1, F1: 2, F2: (2 < 316)}
                                                                            _ = v83
                                                                        }
                                                                        if false {
                                                                            fmt.Println(fi67, v52, fi67)
                                                                            var v84 string = v1
                                                                            _ = v84
                                                                        }
                                                                        if len(v6) >= 2 {
                                                                            v6 = v6[1:len(v6)]
                                                                        }
                                                                    }
                                                                    _ = v73
                                                                }
                                                                if false {
                                                                    var v85 int64 = -445
                                                                    _ = v85
                                                                }
                                                            }
                                                        }
                                                        if (true && false) {
                                                            var v86 float64 = (-2.25 - (0.0 + ((0.1 + (1.5 - float64(1e10))) / 1e10)))
                                                            _ = v86
                                                        }
                                                        _ = v68
                                                    } else {
                                                        _ = len(v4)
                                                        v87 := GenS8{F0: (v52 + (^(v52)))}
                                                        v88 := GenS12{F0: 1}
                                                        _ = v87
                                                        _ = v88
                                                    }
                                                default:
                                                    v5 = v1
                                                }
                                                _ = fi67
                                            }
                                        default:
                                            v55 = (("a" + "hello") + (("\n\t" + "") + "hello"))
                                        }
                                    }
                                } else {
                                    if (v3 == (v2 | v3)) {
                                        var v89 uint32 = 281
                                        _ = v89
                                    } else {
                                        var v90 int = (fi59 / ((-9223372036854775808) | 1))
                                        _ = v90
                                    }
                                    v2 = v2
                                }
                            default:
                                v91 := GenS7{F0: genFn0(872), F1: true}
                                _ = v91
                            }
                            if (!(18446744073709551615 <= 42)) {
                                var v92 int8 = -128
                                v2 -= v2
                                _ = v92
                            } else {
                                v6 = append(v6, genFn0(4294967295))
                                var v93 int32 = 42
                                _ = v93
                            }
                            _ = fi59
                        }
                    }
                }
                _ = v57
            } else {
                v5 = "\n\t"
                switch v1 {
                case "foo bar":
                    var v94 int8 = -128
                    _ = v94
                case "\n\t":
                    var v95 string = ("foo bar" + (v1 + ""))
                    _ = v95
                default:
                    switch 4294967295 {
                    case 127:
                        if false {
                            v3 = v3
                        } else {
                            var v96 uint = 255
                            _ = v96
                        }
                    case -700:
                        v2 = -88
                    }
                }
                fmt.Println(v1, v3, v2)
            }
        }
        if true {
            switch -9223372036854775808 {
            case 0:
                var v97 uint8 = 128
                _ = v97
            default:
                v98 := GenS7{F0: -1, F1: true}
                _ = v98
            }
            v3 -= -146
            var v99 uint = 124
            _ = v99
        }
    }
    if ((127 != 42) || (false && (742 != 336))) {
        for fi100 := 0; fi100 < 13; fi100++ {
            switch fi100 {
            case 42:
                v55 = "a"
            default:
                fmt.Println(fi100)
            }
            var v101 bool = (false || false)
            v5 = ("\n\t" + "a")
            _ = fi100
            _ = v101
        }
        var v102 uint8 = v52
        _ = v102
    } else {
        if false {
            for fi103 := 0; fi103 < 12; fi103++ {
                var v104 uint16 = 2
                fmt.Println(v104, fi103, v52)
                var v105 int8 = -1
                _ = fi103
                _ = v104
                _ = v105
            }
        } else {
            v52 += 127
        }
    }
    if false {
        for fi106 := 0; fi106 < 3; fi106++ {
            if (!(fi106 > (fi106 << ((152) & 63)))) {
                for fi107 := 0; fi107 < 9; fi107++ {
                    if false {
                        fmt.Println(v55, v54, v2)
                    }
                    v54 = "\n\t"
                    fmt.Println(fi106)
                    _ = fi107
                }
                var v108 float32 = ((1e10 - 3.14159) - 3.14159)
                for fi109 := 0; fi109 < 10; fi109++ {
                    switch genFn0(753) {
                    case -744:
                        v55 = v5
                    }
                    v1 = ("" + "a")
                    var v110 uint32 = 2
                    switch "" {
                    case "\n\t":
                        v3 += v3
                    default:
                        var v111 float64 = (-2.25 - ((0.1 / (-2.25 - (float64(float32(1e10)) / ((0.0 - 1e10) * 1e10)))) * 0.0))
                        _ = v111
                    }
                    v5 = ("foo bar" + v55)
                    _ = fi109
                    _ = v110
                }
                _ = v108
            }
            v2 += (v2 << ((472) & 63))
            var v112 uint32 = 956
            _ = fi106
            _ = v112
        }
    }
    if false {
        var v113 bool = ((!true) && ((!(-2147483648 > 1)) && (-1 != 127)))
        _ = v113
    } else {
        for fi114 := 0; fi114 < 14; fi114++ {
            for fi115 := 0; fi115 < 20; fi115++ {
                var v116 int32 = -2147483648
                if false {
                    v55 = (("foo bar" + (("hello" + "a") + v54)) + (v54 + "hello"))
                    var v117 uint = 164
                    var v118 uint64 = 1
                    v52 = 127
                    v117 -= 2
                    _ = v117
                    _ = v118
                }
                _ = len(v4)
                _ = fi115
                _ = v116
            }
            switch (v54 + (v55 + "\n\t")) {
            case "hello":
                v4 = append(v4, genFn0(2))
            case "":
                for fi119 := 0; fi119 < 9; fi119++ {
                    v54 = "a"
                    _ = fi119
                }
            default:
                v120 := GenS13{F0: 2}
                _ = v120
            }
            _ = fi114
        }
    }
    v121 := GenS6{F0: -274}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v121
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v6
}
